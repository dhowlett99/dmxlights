package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"syscall"

	"io/ioutil"
	"os/signal"
	"time"

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/patten"
	"github.com/rakyll/launchpad/mk2"
)

const (
	full = 3
)

var sequenceSpeed int
var fadeSpeed int
var presets map[string]bool
var savePreset bool
var flashButtons [][]bool

func savePresets(presets map[string]bool) {
	// Marshall the config into a json object.
	data, err := json.MarshalIndent(presets, "", " ")
	if err != nil {
		log.Fatalf("Error marshalling config: %v", err)
	}

	// Write to file
	err = ioutil.WriteFile("presets.json", data, 0644)
	if err != nil {
		log.Fatalf("Error writing config: %v to file:%s", err, "presets.json")
	}
}

func loadPresets() map[string]bool {

	presets := map[string]bool{}

	// Read the file.
	data, err := ioutil.ReadFile("presets.json")
	if err != nil {
		fmt.Printf("Error reading prests: %v from file:%s", err, "presets.json")
		return presets
	}

	err = json.Unmarshal(data, &presets)
	if err != nil {
		log.Fatalf("Error unmashalling presets: %v from file:%s", err, "presets.json")
	}

	return presets
}

// main thread is used to get commands from the lauchpad.
func main() {

	selectedSequence := 0
	presets = make(map[string]bool)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		savePresets(presets)
		os.Exit(1)
	}()

	fmt.Println("Derek common.Lighting")

	fmt.Println("Loading Presets")
	presets = loadPresets()
	fmt.Println("Loading Presets Done")

	// Setup a connection to the launchpad.
	pad, err := mk2.Open()
	if err != nil {
		log.Fatalf("Error initializing launchpad: %v", err)
	}
	defer pad.Close()

	fmt.Println("Set Programmers Mode")
	pad.Program()

	// Create a channel to send events to the launchpad.
	eventsForLauchpad := make(chan common.Light)

	// Now create a thread to handle those events.
	go listenAndSendToLaunchPad(eventsForLauchpad, pad)

	// Start off by turning off all of the Lights
	//pad.Clear()

	// Create a channel to listen for buttons being pressed.
	buttonChannel := pad.Listen()

	// Build the default set of Pattens.
	Pattens := patten.MakePatterns()

	// Make a channel to communicate with each sequence.
	// Create a channel for par cans.
	sequence1 := make(chan common.Sequence)
	// sequence2 := make(chan Sequence)
	// sequence3 := make(chan Sequence)
	// sequence4 := make(chan Sequence)

	sequences := []chan common.Sequence{}
	sequences = append(sequences, sequence1)
	// sequences = append(sequences, sequence2)
	// sequences = append(sequences, sequence3)
	// sequences = append(sequences, sequence4)

	// Make channels for each sequence to talk back to us on.
	readSequence1 := make(chan common.Sequence)
	// readSequence2 := make(chan Sequence)
	// readSequence3 := make(chan Sequence)
	// readSequence4 := make(chan Sequence)

	readSequences := []chan common.Sequence{}
	readSequences = append(readSequences, readSequence1)
	// readSequences = append(readSequences, readSequence2)
	// readSequences = append(readSequences, readSequence3)
	// readSequences = append(readSequences, readSequence4)

	// Start threads for each sequence.
	go CreateSequence(1, pad, eventsForLauchpad, sequence1, readSequence1, Pattens)
	// go CreateSequence(2, pad, eventsForLauchpad, sequence2, readSequence2, Pattens)
	// go CreateSequence(3, pad, eventsForLauchpad, sequence3, readSequence3, Pattens)
	// go CreateSequence(4, pad, eventsForLauchpad, sequence4, readSequence4, Pattens)

	// common.Light up any existing presets.
	initPresets(eventsForLauchpad, presets)

	fmt.Println("Setup Presets Done")

	// Initialize a ten length slice of empty slices
	flashButtons = make([][]bool, 9)

	// Initialize those 10 empty slices
	for i := 0; i < 9; i++ {
		flashButtons[i] = make([]bool, 9)
	}

	// common.Light the logo blue.
	pad.Light(8, -1, 79)
	// err = pad.Logo()
	// if err != nil {
	// 	fmt.Printf("Error setting logo %s\n", err.Error())
	// }

	for {
		select {

		case hit := <-buttonChannel:

			pad.Light(hit.X, hit.Y, 79)
			fmt.Printf("Pad X:%d Y:%d\n", hit.X, hit.Y)
			//time.Sleep(1 * time.Second)

			if hit.X == 0 && hit.Y == -1 {
				clearAll(pad, eventsForLauchpad, sequences)
				continue
			}

			// Save mode.
			if hit.X == 8 && hit.Y == 4 {
				if savePreset {
					savePreset = false
					flashButtons[8][4] = false
					continue
				}
				flashButtons[8][4] = true
				flashButton(pad, 8, 4, eventsForLauchpad, 1, 3, 0, 0)
				savePreset = true
			}

			// Ask all sequences for their current config and save in a file.
			if hit.X < 8 && (hit.Y > 3 && hit.Y < 7) {
				if savePreset {
					fmt.Printf("Write Config\n")
					presets[fmt.Sprint(hit.X)+","+fmt.Sprint(hit.Y)] = true
					LightOn(eventsForLauchpad, common.Light{
						X: hit.X, Y: hit.Y, Brightness: full, Red: 3, Green: 0, Blue: 0})
					fmt.Printf("Save Preset in X:%d Y:%d \n", hit.X, hit.Y)
					askToSaveConfig(sequences, readSequences, hit.X, hit.Y)
					savePreset = false
					flashButtons[8][4] = false
				} else {
					// Load config, but only if it exists in the presets map.
					if presets[fmt.Sprint(hit.X)+","+fmt.Sprint(hit.Y)] {
						fmt.Printf("Read Config:")
						fmt.Printf(" OK \n")
						clearAll(pad, eventsForLauchpad, sequences)
						LightOn(eventsForLauchpad, common.Light{X: hit.X, Y: hit.Y, Brightness: full, Red: 3, Green: 0, Blue: 0})
						askToLoadConfig(sequences, hit.X, hit.Y)
						// Reset flash buttons
						for x := 0; x < 9; x++ {
							for y := 0; y < 9; y++ {
								flashButtons[x][y] = false
							}
						}
						flashButtons[hit.X][hit.Y] = true
						flashButton(pad, hit.X, hit.Y, eventsForLauchpad, 1, 0, 3, 0)
					}
				}
			}

			// Select standard Patten.
			if hit.X == 2 && hit.Y == 7 {
				cmd := common.Sequence{
					UpdatePatten: true,
					Patten: common.Patten{
						Name: "rgbchase",
					},
				}
				if selectedSequence == 1 {
					sequence1 <- cmd
				}
				// if selectedSequence == 2 {
				// 	sequence2 <- cmd
				// }
				// if selectedSequence == 3 {
				// 	sequence3 <- cmd
				// }
				// if selectedSequence == 4 {
				// 	sequence4 <- cmd
				// }
				continue
			}
			// Select pairs Patten.
			if hit.X == 3 && hit.Y == 7 {
				cmd := common.Sequence{
					UpdatePatten: true,
					Patten: common.Patten{

						Name: "pairs",
					},
				}
				if selectedSequence == 1 {
					sequence1 <- cmd
				}
				// if selectedSequence == 2 {
				// 	sequence2 <- cmd
				// }
				// if selectedSequence == 3 {
				// 	sequence3 <- cmd
				// }
				// if selectedSequence == 4 {
				// 	sequence4 <- cmd
				// }
				continue
			}

			// Decrease speed of selected sequence.
			if hit.X == 0 && hit.Y == 7 {
				sequenceSpeed--
				if sequenceSpeed < 0 {
					sequenceSpeed = 1
				}
				fmt.Printf("Seq Speed %d\n", sequenceSpeed)
				cmd := common.Sequence{
					Speed:       sequenceSpeed,
					UpdateSpeed: true,
				}
				if selectedSequence == 1 {
					sequence1 <- cmd
				}
				// if selectedSequence == 2 {
				// 	sequence2 <- cmd
				// }
				// if selectedSequence == 3 {
				// 	sequence3 <- cmd
				// }
				// if selectedSequence == 4 {
				// 	sequence4 <- cmd
				// }

				continue
			}

			// Increase speed of selected sequence.
			if hit.X == 1 && hit.Y == 7 {
				sequenceSpeed++
				if sequenceSpeed > 12 {
					sequenceSpeed = 12
				}
				cmd := common.Sequence{
					Speed:       sequenceSpeed,
					UpdateSpeed: true,
				}
				if selectedSequence == 1 {
					sequence1 <- cmd
				}
				// if selectedSequence == 2 {
				// 	sequence2 <- cmd
				// }
				// if selectedSequence == 3 {
				// 	sequence3 <- cmd
				// }
				// if selectedSequence == 4 {
				// 	sequence4 <- cmd
				// }

				continue
			}

			// Select sequence 1.
			if hit.X == 8 && hit.Y == 0 {
				selectedSequence = 1
				LightOn(eventsForLauchpad, common.Light{X: 8, Y: 0, Brightness: full, Red: 3, Green: 3, Blue: 0})
				event := common.Light{X: 8, Y: 1, Brightness: full, Red: 0, Green: 0, Blue: 0}
				eventsForLauchpad <- event
				event = common.Light{X: 8, Y: 2, Brightness: full, Red: 0, Green: 0, Blue: 0}
				eventsForLauchpad <- event
				event = common.Light{X: 8, Y: 3, Brightness: full, Red: 0, Green: 0, Blue: 0}
				eventsForLauchpad <- event
				continue
			}

			// Select sequence 2.
			if hit.X == 8 && hit.Y == 1 {
				selectedSequence = 2
				event := common.Light{X: 8, Y: 1, Brightness: full, Red: 3, Green: 3, Blue: 0}
				eventsForLauchpad <- event

				event = common.Light{X: 8, Y: 0, Brightness: full, Red: 0, Green: 0, Blue: 0}
				eventsForLauchpad <- event
				event = common.Light{X: 8, Y: 2, Brightness: full, Red: 0, Green: 0, Blue: 0}
				eventsForLauchpad <- event
				event = common.Light{X: 8, Y: 3, Brightness: full, Red: 0, Green: 0, Blue: 0}
				eventsForLauchpad <- event
				continue
			}

			// Select sequence 3.
			if hit.X == 8 && hit.Y == 2 {
				selectedSequence = 3
				event := common.Light{X: 8, Y: 2, Brightness: full, Red: 3, Green: 3, Blue: 0}
				eventsForLauchpad <- event

				event = common.Light{X: 8, Y: 1, Brightness: full, Red: 0, Green: 0, Blue: 0}
				eventsForLauchpad <- event
				event = common.Light{X: 8, Y: 3, Brightness: full, Red: 0, Green: 0, Blue: 0}
				eventsForLauchpad <- event
				event = common.Light{X: 8, Y: 0, Brightness: full, Red: 0, Green: 0, Blue: 0}
				eventsForLauchpad <- event
				continue
			}

			// Select sequence 4.
			if hit.X == 8 && hit.Y == 3 {
				selectedSequence = 4
				event := common.Light{X: 8, Y: 3, Brightness: full, Red: 3, Green: 3, Blue: 0}
				eventsForLauchpad <- event

				event = common.Light{X: 8, Y: 1, Brightness: full, Red: 0, Green: 0, Blue: 0}
				eventsForLauchpad <- event
				event = common.Light{X: 8, Y: 2, Brightness: full, Red: 0, Green: 0, Blue: 0}
				eventsForLauchpad <- event
				event = common.Light{X: 8, Y: 4, Brightness: full, Red: 0, Green: 0, Blue: 0}
				eventsForLauchpad <- event
				continue
			}

			// Start sequence.
			if hit.X == 8 && hit.Y == 5 {
				cmd := common.Sequence{
					Start: true,
				}
				if selectedSequence == 1 {
					sequence1 <- cmd
				}
				// if selectedSequence == 2 {
				// 	sequence2 <- cmd
				// }
				// if selectedSequence == 3 {
				// 	sequence3 <- cmd
				// }
				// if selectedSequence == 4 {
				// 	sequence4 <- cmd
				// }
			}
			// Stop sequence.
			if hit.X == 8 && hit.Y == 6 {
				cmd := common.Sequence{
					Stop: true,
				}
				if selectedSequence == 1 {
					sequence1 <- cmd
				}
				// if selectedSequence == 2 {
				// 	sequence2 <- cmd
				// }
				// if selectedSequence == 3 {
				// 	sequence3 <- cmd
				// }
				// if selectedSequence == 4 {
				// 	sequence4 <- cmd
				// }
			}

			// Fade time decrease.
			if hit.X == 4 && hit.Y == 7 {
				fadeSpeed--
				if fadeSpeed < 0 {
					fadeSpeed = 0
				}
				fmt.Printf("Fade down speed:%d", fadeSpeed)
				fadeTime := setFade(fadeSpeed)
				cmd := common.Sequence{
					UpdateFade: true,
					FadeTime:   fadeTime,
					Number:     selectedSequence,
				}
				if selectedSequence == 1 {
					sequence1 <- cmd
				}
				// if selectedSequence == 2 {
				// 	sequence2 <- cmd
				// }
				// if selectedSequence == 3 {
				// 	sequence3 <- cmd
				// }
				// if selectedSequence == 4 {
				// 	sequence4 <- cmd
				// }

				continue
			}

			// Fade time increase.
			if hit.X == 5 && hit.Y == 7 {
				fadeSpeed++
				if fadeSpeed > 12 {
					fadeSpeed = 12
				}
				fmt.Printf("Fade up speed:%d", fadeSpeed)
				fadeTime := setSpeed(fadeSpeed)
				cmd := common.Sequence{
					UpdateFade: true,
					FadeTime:   fadeTime,
				}
				if selectedSequence == 1 {
					sequence1 <- cmd
				}
				// if selectedSequence == 2 {
				// 	sequence2 <- cmd
				// }
				// if selectedSequence == 3 {
				// 	sequence3 <- cmd
				// }
				// if selectedSequence == 4 {
				// 	sequence4 <- cmd
				// }

				continue
			}

			// // common.Light a button is pressed.
			// if !button[hit.X][hit.Y] {
			// 	event := common.Light{hit.X, hit.Y, 0, 3}
			// 	eventsForLauchpad <- event
			// 	button[hit.X][hit.Y] = true
			// } else {
			// 	event := common.Light{hit.X, hit.Y, 0, 0}
			// 	eventsForLauchpad <- event
			// 	button[hit.X][hit.Y] = false
			// }
		}

	}
}

func initPresets(eventsForLauchpad chan common.Light, presets map[string]bool) {
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			if presets[fmt.Sprint(x)+","+fmt.Sprint(y)] {
				LightOn(eventsForLauchpad, common.Light{X: x, Y: y, Brightness: full, Red: 3, Green: 0, Blue: 0})
			}
		}
	}
}

func clearAll(pad *mk2.Launchpad, eventsForLauchpad chan common.Light, sequences []chan common.Sequence) {
	fmt.Printf("C L E A R\n")
	pad.Reset()
	cmd := common.Sequence{
		Stop: true,
	}
	for _, seq := range sequences {
		seq <- cmd
	}

	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			if presets[fmt.Sprint(x)+","+fmt.Sprint(y)] {
				LightOn(eventsForLauchpad, common.Light{X: x, Y: y, Brightness: full, Red: 3, Green: 0, Blue: 0})
			}
		}
	}
}

func CreateSequence(mySequenceNumber int, pad *mk2.Launchpad, eventsForLauchpad chan common.Light, commandChannel chan common.Sequence, replyChannel chan common.Sequence, Pattens map[string]common.Patten) {

	fmt.Printf("Setup default command\n")
	// set default values.
	command := common.Sequence{
		Name:     "cans",
		Number:   mySequenceNumber,
		FadeTime: 0 * time.Millisecond,
		Run:      true,
		Patten: common.Patten{
			Name:     "standard",
			Length:   2,
			Size:     2,
			Fixtures: 8,
			Chase:    []int{1, 2, 3, 4, 5, 6, 7, 8},
			Steps:    Pattens["standard"].Steps,
		},
		CurrentSpeed: 500 * time.Millisecond,
		Colors: []common.Color{
			{
				R: 0,
				G: 0,
				B: 0,
			},
		},
		Speed: 3,
	}

	// Create a channel for every fixture.
	fmt.Printf("Create a channel for every fixture.\n")
	fixtureChannels := []chan common.Event{}
	for fixture := 0; fixture < command.Patten.Fixtures; fixture++ {
		channel := make(chan common.Event)
		fixtureChannels = append(fixtureChannels, channel)
	}

	// Now start the fixture threads listening.
	fmt.Printf("Now start the fixture threads listening.")
	for fixture, channel := range fixtureChannels {

		//fmt.Printf("Start a thread %d fixture.\n", fixture)
		//time.Sleep(1 * time.Second)

		go fixtureReceiver(channel, fixture, command, commandChannel, replyChannel, mySequenceNumber, Pattens, eventsForLauchpad)

	}

	event := common.Event{}
	// Now start the fixture threads by sending an event.
	for {
		//fmt.Printf("Step to fixture loop %d fixtureChannels = %+v\n", event.Fixture, fixtureChannels)
		command = listenCommandChannelAndWait(command, commandChannel, replyChannel, command.CurrentSpeed, mySequenceNumber)
		for index, channel := range fixtureChannels {
			event.Fixture = index
			if command.Run {
				event.Start = true
			}
			channel <- event
		}
	}
}
func fixtureReceiver(channel chan common.Event, fixture int, command common.Sequence, commandChannel chan common.Sequence, replyChannel chan common.Sequence, mySequenceNumber int, Pattens map[string]common.Patten, eventsForLauchpad chan common.Light) {

	// Start the step counter so we know where we are in the sequence.
	stepCount := 0

	// Start the color counter.
	currentColor := 0

	fmt.Printf("Now Listening on channel %d\n", fixture)
	for {

		event := <-channel

		// Are we being asked to start.
		if event.Start {
			// Listen on this fixtures channel for the step events.
			step := Pattens[command.Patten.Name].Steps
			totalSteps := len(command.Patten.Steps)
			tolalColors := len(step[stepCount].Fixtures[fixture].Colors)

			R := step[stepCount].Fixtures[fixture].Colors[currentColor].R
			G := step[stepCount].Fixtures[fixture].Colors[currentColor].G
			B := step[stepCount].Fixtures[fixture].Colors[currentColor].B

			if currentColor <= tolalColors {
				currentColor++
			}
			// Fade up
			command = listenCommandChannelAndWait(command, commandChannel, replyChannel, (command.CurrentSpeed/4)/2, mySequenceNumber)
			if R > 0 || G > 0 || B > 0 {
				for green := 0; green <= step[stepCount].Fixtures[fixture].Colors[0].G; green++ {
					command = listenCommandChannelAndWait(command, commandChannel, replyChannel, command.CurrentSpeed/4, mySequenceNumber)
					event := common.Light{X: fixture, Y: mySequenceNumber - 1, Brightness: 3, Red: R, Green: green, Blue: B}
					eventsForLauchpad <- event
				}
				command = listenCommandChannelAndWait(command, commandChannel, replyChannel, (command.CurrentSpeed/4)/2, mySequenceNumber)
				for green := step[stepCount].Fixtures[fixture].Colors[0].G; green >= 0; green-- {
					command = listenCommandChannelAndWait(command, commandChannel, replyChannel, command.CurrentSpeed/4, mySequenceNumber)
					event := common.Light{X: fixture, Y: mySequenceNumber - 1, Brightness: 3, Red: R, Green: green, Blue: B}
					eventsForLauchpad <- event
				}
			}

			if currentColor == tolalColors {
				stepCount++
				currentColor = 0
			}

			if stepCount >= totalSteps {
				stepCount = 0
				currentColor = 0
			}
		}
	}
}

// waitForConfig
func waitForConfig(replyChannel chan common.Sequence) common.Sequence {
	command := common.Sequence{}
	select {
	case command = <-replyChannel:
		fmt.Printf("Config Received for seq: %s\n", command.Name)
		break
	case <-time.After(500 * time.Millisecond):
		fmt.Printf("Config TIMEOUT for seq: %s\n", command.Name)
		break
	}
	return command
}

// listenCommandChannelAndWait listens on channel for instructions or timeout and go to next step of sequence.
func listenCommandChannelAndWait(command common.Sequence, commandChannel chan common.Sequence, replyChannel chan common.Sequence, CurrentSpeed time.Duration, mySequenceNumber int) common.Sequence {

	currentCommand := command
	select {
	case command = <-commandChannel:
		//fmt.Printf("COMMAND\n")
		break
	case <-time.After(CurrentSpeed):
		//fmt.Printf("TIMEOUT\n")
		break
	}
	if command.UpdateSpeed {
		saveSpeed := command.Speed
		fmt.Printf("Received update speed %d\n", saveSpeed)
		CurrentSpeed = setSpeed(command.Speed)
		command = currentCommand
		command.CurrentSpeed = CurrentSpeed
		command.Speed = saveSpeed
	}

	if command.UpdatePatten {
		savePattenName := command.Patten.Name
		fmt.Printf("Received update pattten %s\n", savePattenName)
		command = currentCommand
		command.Patten.Name = savePattenName
		command.UpdatePatten = true
	}

	if command.UpdateFade {
		fadeTime := command.FadeTime
		fmt.Printf("Received new fade time of %v\n", fadeTime)
		command = currentCommand
		command.FadeTime = fadeTime
		command.UpdateFade = true
	}

	if command.Start {
		fmt.Printf("Received Start Seq \n")
		command = currentCommand
		command.Run = true
	}

	if command.Stop {
		fmt.Printf("Received Stop Seq \n")
		command = currentCommand
		command.Run = false
	}

	if command.ReadConfig {
		fmt.Printf("Sending Reply on %d\n", currentCommand.Number)
		currentCommand.X = command.X
		currentCommand.Y = command.Y
		replyChannel <- currentCommand
		command = currentCommand
	}

	if command.LoadConfig {
		X := command.X
		Y := command.Y
		config := readConfig(fmt.Sprintf("config%d.%d.json", X, Y))

		for _, seq := range config {
			if seq.Number == mySequenceNumber {
				command = seq
			}
		}
		command.LoadConfig = true
	}
	return command
}

// listenAndSendToLaunchPad is the thread that listens for events to send to
// the launch pad.  It is thread safe and is the only thread talking to the
// launch pad. A channel is used to queue the events to be sent.
func listenAndSendToLaunchPad(eventsForLauchpad chan common.Light, pad *mk2.Launchpad) {
	var green int
	var red int
	var blue int

	for {
		event := <-eventsForLauchpad
		switch event.Green {
		case 0:
			green = 0
		case 1:
			green = 19
		case 2:
			green = 22
		case 3:
			green = 21
		}

		switch event.Red {
		case 0:
			red = 0
		case 1:
			red = 7
		case 2:
			red = 6
		case 3:
			red = 5
		}

		switch event.Blue {
		case 0:
			blue = 0
		case 1:
			blue = 37
		case 2:
			blue = 38
		case 3:
			blue = 79
		}

		pad.Light(event.X, event.Y, green+red+blue)
	}
}

func setSpeed(commandSpeed int) (Speed time.Duration) {
	if commandSpeed == 0 {
		Speed = 3500
	}
	if commandSpeed == 1 {
		Speed = 3000
	}
	if commandSpeed == 2 {
		Speed = 2500
	}
	if commandSpeed == 3 {
		Speed = 1000
	}
	if commandSpeed == 4 {
		Speed = 1500
	}
	if commandSpeed == 5 {
		Speed = 1000
	}
	if commandSpeed == 6 {
		Speed = 750
	}
	if commandSpeed == 7 {
		Speed = 500
	}
	if commandSpeed == 8 {
		Speed = 250
	}
	if commandSpeed == 9 {
		Speed = 150
	}
	if commandSpeed == 10 {
		Speed = 100
	}
	if commandSpeed == 11 {
		Speed = 50
	}
	return Speed * time.Millisecond
}

func setFade(commandSpeed int) (Speed time.Duration) {
	if commandSpeed == 0 {
		Speed = 1000
	}
	if commandSpeed == 1 {
		Speed = 900
	}
	if commandSpeed == 2 {
		Speed = 800
	}
	if commandSpeed == 3 {
		Speed = 700
	}
	if commandSpeed == 4 {
		Speed = 600
	}
	if commandSpeed == 5 {
		Speed = 500
	}
	if commandSpeed == 6 {
		Speed = 400
	}
	if commandSpeed == 7 {
		Speed = 300
	}
	if commandSpeed == 8 {
		Speed = 200
	}
	if commandSpeed == 9 {
		Speed = 150
	}
	if commandSpeed == 10 {
		Speed = 100
	}
	if commandSpeed == 11 {
		Speed = 50
	}
	return Speed * time.Millisecond
}

// common.LightOn Turn on a common.Light.
func LightOn(eventsForLauchpad chan common.Light, Light common.Light) {
	event := common.Light{X: Light.X, Y: Light.Y, Brightness: full, Red: Light.Red, Green: Light.Green, Blue: Light.Blue}
	eventsForLauchpad <- event
}

func writeConfig(config []common.Sequence, filename string) {

	// Marshall the config into a json object.
	data, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		log.Fatalf("Error marshalling config: %v", err)
	}

	//fmt.Println(string(data))

	// Write to file
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		log.Fatalf("Error writing config: %v to file:%s", err, filename)
	}

}

func readConfig(filename string) []common.Sequence {

	config := []common.Sequence{}

	// Read the file.
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading config: %v from file:%s", err, filename)
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Error reading config: %v from file:%s", err, filename)
	}

	return config
}

func askToLoadConfig(sequences []chan common.Sequence, X int, Y int) {
	cmd := common.Sequence{
		LoadConfig: true,
		X:          X,
		Y:          Y,
	}
	for _, seq := range sequences {
		seq <- cmd
	}
}

func askToSaveConfig(sequences []chan common.Sequence, replyChannel []chan common.Sequence, X int, Y int) {

	fmt.Printf("askToSaveConfig: Save Preset in X:%d Y:%d \n", X, Y)
	config := []common.Sequence{}

	go func() {
		// wait for responses from sequences.
		time.Sleep(100 * time.Millisecond)
		for _, replyChannel := range replyChannel {
			config = append(config, waitForConfig(replyChannel))
		}
		// write to config file.
		writeConfig(config, fmt.Sprintf("config%d.%d.json", config[0].X, config[0].Y))
	}()

	// ask for all the sequencers for their config.
	cmd := common.Sequence{
		ReadConfig: true,
		X:          X,
		Y:          Y,
	}
	for _, seq := range sequences {
		seq <- cmd
	}
}

func flashButton(pad *mk2.Launchpad, x int, y int, eventsForLauchpad chan common.Light, seqNumber int, green int, red int, blue int) {
	go func(pad *mk2.Launchpad, x int, y int) {

		for {
			// fmt.Printf("Flash X:%d Y:%d is %t\n", x, y, flashButtons[x][y])
			if !flashButtons[x][y] {
				break
			}
			event := common.Light{X: x, Y: y, Brightness: full, Red: red, Green: green, Blue: blue}
			eventsForLauchpad <- event

			time.Sleep(1 * time.Second)
			event = common.Light{X: x, Y: y, Brightness: full, Red: 0, Green: 0, Blue: 0}
			eventsForLauchpad <- event
			time.Sleep(1 * time.Second)
		}
	}(pad, x, y)
}
