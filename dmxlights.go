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

type Light struct {
	X          int
	Y          int
	Brightness int
	Red        int
	Green      int
	Blue       int
}

type Color struct {
	R int
	G int
	B int
}

type Patten struct {
	Name     string
	Length   int // 8, 4 or 2
	Size     int
	Fixtures int // 8 Fixtures
	Chase    []int
	Steps    []Steps
}

type Sequence struct {
	// commands
	Start        bool
	Stop         bool
	ReadConfig   bool
	LoadConfig   bool
	UpdateSpeed  bool
	UpdatePatten bool
	UpdateFade   bool
	// parameters
	FadeTime     time.Duration
	Name         string
	Number       int
	Run          bool
	Patten       Patten
	Colors       []Color
	Speed        int
	CurrentSpeed time.Duration
	X            int
	Y            int
}

type Hit struct {
	X int
	Y int
}

type Steps struct {
	Fixtures []Fixture
}

type Fixture struct {
	Brightness int
	Colors     []Color
}

type ButtonPresets struct {
	X int
	Y int
}

type Event struct {
	Start   bool
	Fixture int
	Step    int
}

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

	fmt.Println("Derek Lighting")

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
	eventsForLauchpad := make(chan Light)

	// Now create a thread to handle those events.
	go listenAndSendToLaunchPad(eventsForLauchpad, pad)

	// Start off by turning off all of the lights
	//pad.Clear()

	// Create a channel to listen for buttons being pressed.
	buttonChannel := pad.Listen()

	// Build the default set of Pattens.
	Pattens := makePatterns()

	// Make a channel to communicate with each sequence.
	// Create a channel for par cans.
	sequence1 := make(chan Sequence)
	// sequence2 := make(chan Sequence)
	// sequence3 := make(chan Sequence)
	// sequence4 := make(chan Sequence)

	sequences := []chan Sequence{}
	sequences = append(sequences, sequence1)
	// sequences = append(sequences, sequence2)
	// sequences = append(sequences, sequence3)
	// sequences = append(sequences, sequence4)

	// Make channels for each sequence to talk back to us on.
	readSequence1 := make(chan Sequence)
	// readSequence2 := make(chan Sequence)
	// readSequence3 := make(chan Sequence)
	// readSequence4 := make(chan Sequence)

	readSequences := []chan Sequence{}
	readSequences = append(readSequences, readSequence1)
	// readSequences = append(readSequences, readSequence2)
	// readSequences = append(readSequences, readSequence3)
	// readSequences = append(readSequences, readSequence4)

	// Start threads for each sequence.
	go CreateSequence(1, pad, eventsForLauchpad, sequence1, readSequence1, Pattens)
	// go CreateSequence(2, pad, eventsForLauchpad, sequence2, readSequence2, Pattens)
	// go CreateSequence(3, pad, eventsForLauchpad, sequence3, readSequence3, Pattens)
	// go CreateSequence(4, pad, eventsForLauchpad, sequence4, readSequence4, Pattens)

	// Light up any existing presets.
	initPresets(eventsForLauchpad, presets)

	fmt.Println("Setup Presets Done")

	// Initialize a ten length slice of empty slices
	flashButtons = make([][]bool, 9)

	// Initialize those 10 empty slices
	for i := 0; i < 9; i++ {
		flashButtons[i] = make([]bool, 9)
	}

	// Light the logo blue.
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
					lightOn(eventsForLauchpad, Light{hit.X, hit.Y, full, 3, 0, 0})
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
						lightOn(eventsForLauchpad, Light{hit.X, hit.Y, full, 3, 0, 0})
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
				cmd := Sequence{
					UpdatePatten: true,
					Patten: Patten{
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
				cmd := Sequence{
					UpdatePatten: true,
					Patten: Patten{

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
				cmd := Sequence{
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
				cmd := Sequence{
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
				lightOn(eventsForLauchpad, Light{8, 0, full, 3, 3, 0})

				event := Light{8, 1, full, 0, 0, 0}
				eventsForLauchpad <- event
				event = Light{8, 2, full, 0, 0, 0}
				eventsForLauchpad <- event
				event = Light{8, 3, full, 0, 0, 0}
				eventsForLauchpad <- event
				continue
			}

			// Select sequence 2.
			if hit.X == 8 && hit.Y == 1 {
				selectedSequence = 2
				event := Light{8, 1, full, 3, 3, 0}
				eventsForLauchpad <- event

				event = Light{8, 0, full, 0, 0, 0}
				eventsForLauchpad <- event
				event = Light{8, 2, full, 0, 0, 0}
				eventsForLauchpad <- event
				event = Light{8, 3, full, 0, 0, 0}
				eventsForLauchpad <- event
				continue
			}

			// Select sequence 3.
			if hit.X == 8 && hit.Y == 2 {
				selectedSequence = 3
				event := Light{8, 2, full, 3, 3, 0}
				eventsForLauchpad <- event

				event = Light{8, 1, full, 0, 0, 0}
				eventsForLauchpad <- event
				event = Light{8, 3, full, 0, 0, 0}
				eventsForLauchpad <- event
				event = Light{8, 0, full, 0, 0, 0}
				eventsForLauchpad <- event
				continue
			}

			// Select sequence 4.
			if hit.X == 8 && hit.Y == 3 {
				selectedSequence = 4
				event := Light{8, 3, full, 3, 3, 0}
				eventsForLauchpad <- event

				event = Light{8, 1, full, 0, 0, 0}
				eventsForLauchpad <- event
				event = Light{8, 2, full, 0, 0, 0}
				eventsForLauchpad <- event
				event = Light{8, 4, full, 0, 0, 0}
				eventsForLauchpad <- event
				continue
			}

			// Start sequence.
			if hit.X == 8 && hit.Y == 5 {
				cmd := Sequence{
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
				cmd := Sequence{
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
				cmd := Sequence{
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
				cmd := Sequence{
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

			// // Light a button is pressed.
			// if !button[hit.X][hit.Y] {
			// 	event := Light{hit.X, hit.Y, 0, 3}
			// 	eventsForLauchpad <- event
			// 	button[hit.X][hit.Y] = true
			// } else {
			// 	event := Light{hit.X, hit.Y, 0, 0}
			// 	eventsForLauchpad <- event
			// 	button[hit.X][hit.Y] = false
			// }
		}

	}
}

func initPresets(eventsForLauchpad chan Light, presets map[string]bool) {
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			if presets[fmt.Sprint(x)+","+fmt.Sprint(y)] {
				lightOn(eventsForLauchpad, Light{x, y, full, 3, 0, 0})
			}
		}
	}
}

func clearAll(pad *mk2.Launchpad, eventsForLauchpad chan Light, sequences []chan Sequence) {
	fmt.Printf("C L E A R\n")
	pad.Reset()
	cmd := Sequence{
		Stop: true,
	}
	for _, seq := range sequences {
		seq <- cmd
	}

	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			if presets[fmt.Sprint(x)+","+fmt.Sprint(y)] {
				lightOn(eventsForLauchpad, Light{x, y, full, 3, 0, 0})
			}
		}
	}
}

func CreateSequence(mySequenceNumber int, pad *mk2.Launchpad, eventsForLauchpad chan Light, commandChannel chan Sequence, replyChannel chan Sequence, Pattens map[string]Patten) {

	fmt.Printf("Setup default command\n")
	// set default values.
	command := Sequence{
		Name:     "cans",
		Number:   mySequenceNumber,
		FadeTime: 0 * time.Millisecond,
		Run:      true,
		Patten: Patten{
			Name:     "standard",
			Length:   2,
			Size:     2,
			Fixtures: 8,
			Chase:    []int{1, 2, 3, 4, 5, 6, 7, 8},
			Steps:    Pattens["standard"].Steps,
		},
		CurrentSpeed: 500 * time.Millisecond,
		Colors: []Color{
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
	fixtureChannels := []chan Event{}
	for fixture := 0; fixture < command.Patten.Fixtures; fixture++ {
		channel := make(chan Event)
		fixtureChannels = append(fixtureChannels, channel)
	}

	// Now start the fixture threads listening.
	fmt.Printf("Now start the fixture threads listening.")
	for fixture, channel := range fixtureChannels {

		//fmt.Printf("Start a thread %d fixture.\n", fixture)
		//time.Sleep(1 * time.Second)

		go fixtureReceiver(channel, fixture, command, commandChannel, replyChannel, mySequenceNumber, Pattens, eventsForLauchpad)

	}

	event := Event{}
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
func fixtureReceiver(channel chan Event, fixture int, command Sequence, commandChannel chan Sequence, replyChannel chan Sequence, mySequenceNumber int, Pattens map[string]Patten, eventsForLauchpad chan Light) {

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
					event := Light{fixture, mySequenceNumber - 1, 3, R, green, B}
					eventsForLauchpad <- event
				}
				command = listenCommandChannelAndWait(command, commandChannel, replyChannel, (command.CurrentSpeed/4)/2, mySequenceNumber)
				for green := step[stepCount].Fixtures[fixture].Colors[0].G; green >= 0; green-- {
					command = listenCommandChannelAndWait(command, commandChannel, replyChannel, command.CurrentSpeed/4, mySequenceNumber)
					event := Light{fixture, mySequenceNumber - 1, 3, R, green, B}
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
func waitForConfig(replyChannel chan Sequence) Sequence {
	command := Sequence{}
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
func listenCommandChannelAndWait(command Sequence, commandChannel chan Sequence, replyChannel chan Sequence, CurrentSpeed time.Duration, mySequenceNumber int) Sequence {

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
func listenAndSendToLaunchPad(eventsForLauchpad chan Light, pad *mk2.Launchpad) {
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

// lightOn Turn on a light.
func lightOn(eventsForLauchpad chan Light, light Light) {
	event := Light{light.X, light.Y, full, light.Red, light.Green, light.Blue}
	eventsForLauchpad <- event
}

func makePatterns() map[string]Patten {

	Pattens := make(map[string]Patten)

	standard := Patten{
		Steps: []Steps{
			{
				Fixtures: []Fixture{
					{Brightness: full, Colors: []Color{{R: 0, G: 3, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []Fixture{
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 3, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []Fixture{
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 3, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []Fixture{
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 3, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []Fixture{
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 3, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []Fixture{
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 3, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []Fixture{
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 3, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []Fixture{
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 3, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
				},
			},
		},
	}

	rgbchase := Patten{
		Steps: []Steps{
			{
				Fixtures: []Fixture{
					{Brightness: full, Colors: []Color{{R: 3, G: 0, B: 0}, {R: 0, G: 3, B: 0}, {R: 0, G: 0, B: 3}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []Fixture{
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 3, G: 0, B: 0}, {R: 0, G: 3, B: 0}, {R: 0, G: 0, B: 3}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []Fixture{
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 3, G: 0, B: 0}, {R: 0, G: 3, B: 0}, {R: 0, G: 0, B: 3}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []Fixture{
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 3, G: 0, B: 0}, {R: 0, G: 3, B: 0}, {R: 0, G: 0, B: 3}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []Fixture{
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 3, G: 0, B: 0}, {R: 0, G: 3, B: 0}, {R: 0, G: 0, B: 3}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []Fixture{
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 3, G: 0, B: 0}, {R: 0, G: 3, B: 0}, {R: 0, G: 0, B: 3}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []Fixture{
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 3, G: 0, B: 0}, {R: 0, G: 3, B: 0}, {R: 0, G: 0, B: 3}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []Fixture{
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 3, G: 0, B: 0}, {R: 0, G: 3, B: 0}, {R: 0, G: 0, B: 3}}},
				},
			},
		},
	}

	pairs := Patten{
		Steps: []Steps{
			{
				Fixtures: []Fixture{
					{Brightness: full, Colors: []Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []Fixture{
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 3, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []Fixture{
					{Brightness: full, Colors: []Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []Fixture{
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 3, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []Fixture{
					{Brightness: full, Colors: []Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []Fixture{
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 3, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []Fixture{
					{Brightness: full, Colors: []Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []Fixture{
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []Color{{R: 3, G: 0, B: 0}}},
				},
			},
		},
	}

	Pattens["standard"] = standard
	Pattens["rgbchase"] = rgbchase
	Pattens["pairs"] = pairs

	return Pattens
}

func writeConfig(config []Sequence, filename string) {

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

func readConfig(filename string) []Sequence {

	config := []Sequence{}

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

func askToLoadConfig(sequences []chan Sequence, X int, Y int) {
	cmd := Sequence{
		LoadConfig: true,
		X:          X,
		Y:          Y,
	}
	for _, seq := range sequences {
		seq <- cmd
	}
}

func askToSaveConfig(sequences []chan Sequence, replyChannel []chan Sequence, X int, Y int) {

	fmt.Printf("askToSaveConfig: Save Preset in X:%d Y:%d \n", X, Y)
	config := []Sequence{}

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
	cmd := Sequence{
		ReadConfig: true,
		X:          X,
		Y:          Y,
	}
	for _, seq := range sequences {
		seq <- cmd
	}
}

func flashButton(pad *mk2.Launchpad, x int, y int, eventsForLauchpad chan Light, seqNumber int, green int, red int, blue int) {
	go func(pad *mk2.Launchpad, x int, y int) {

		for {
			// fmt.Printf("Flash X:%d Y:%d is %t\n", x, y, flashButtons[x][y])
			if !flashButtons[x][y] {
				break
			}
			event := Light{x, y, full, red, green, blue}
			eventsForLauchpad <- event

			time.Sleep(1 * time.Second)
			event = Light{x, y, full, 0, 0, 0}
			eventsForLauchpad <- event
			time.Sleep(1 * time.Second)
		}
	}(pad, x, y)
}
