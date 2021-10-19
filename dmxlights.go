package main

import (
	"fmt"
	"log"
	"os"
	"syscall"

	"os/signal"
	"time"

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/config"
	"github.com/dhowlett99/dmxlights/pkg/dmx"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
	"github.com/oliread/usbdmx/ft232"

	"github.com/dhowlett99/dmxlights/pkg/launchpad"
	"github.com/dhowlett99/dmxlights/pkg/patten"
	"github.com/dhowlett99/dmxlights/pkg/presets"
	"github.com/dhowlett99/dmxlights/pkg/sequence"
	"github.com/dhowlett99/dmxlights/pkg/sound"
	"github.com/rakyll/launchpad/mk3"
)

const (
	full = 255
)

var sequenceSpeed int = 12
var fadeSpeed int
var size int
var masterBrightness int

//var color int
var savePreset bool
var selectedPatten = 0
var blackout bool = false

// main thread is used to get commands from the lauchpad.
func main() {

	var functionButtons [][]bool   // Function buttons.
	var functionSelectMode []bool  // Which sequence is in function selection mode.
	var selectButtonPressed []bool // Which sequence has its Select button pressed.
	var staticLamps [][]bool       // Static color lamps.
	var switchPositions [9][9]int  // Sorage for switch positions.
	var colorEditModeDone []bool   // This sequence is done color editing mode.
	fadeSpeed = 11                 // Default start at 50ms.
	masterBrightness = 255         // Affects all DMX fixtures and launchpad lamps.
	var lastStaticColorButtonX int // Which Static Color button did we change last.
	var lastStaticColorButtonY int // Which Static Color button did we change last.

	// Make an empty presets store.
	presetsStore := make(map[string]bool)

	// Save the presets on exit.
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		presets.SavePresets(presetsStore)
		os.Exit(1)
	}()

	fmt.Println("DMX Lighting")
	presetsStore = presets.LoadPresets()

	// Setup DMX interface.
	dmxController, err := dmx.NewDmXController()
	if err != nil {
		fmt.Printf("error initializing dmx interface: %v\n", err)
		os.Exit(1)
	}

	// Setup a connection to the launchpad.
	// Tested with a Novation Launchpad mini mk3.
	pad, err := mk3.Open()
	if err != nil {
		log.Fatalf("error initializing launchpad: %v", err)
	}
	defer pad.Close()

	// We need to be in programmers mode to use the launchpad.
	pad.Program()

	// Create all the channels I need.
	commandChannels := []chan common.Command{}
	replyChannels := []chan common.Sequence{}
	soundTriggerChannels := []chan common.Command{}
	updateChannels := []chan common.Sequence{}

	// Make channels for commands.
	for sequence := 0; sequence < 4; sequence++ {
		commandChannel := make(chan common.Command)
		commandChannels = append(commandChannels, commandChannel)
		replyChannel := make(chan common.Sequence)
		replyChannels = append(replyChannels, replyChannel)
		soundTriggerChannel := make(chan common.Command)
		soundTriggerChannels = append(soundTriggerChannels, soundTriggerChannel)
		updateChannel := make(chan common.Sequence)
		updateChannels = append(updateChannels, updateChannel)
	}

	// Now add them all to a handy channels struct.
	sequenceChannels := common.Channels{}
	sequenceChannels.CommmandChannels = commandChannels
	sequenceChannels.ReplyChannels = replyChannels
	sequenceChannels.SoundTriggerChannels = soundTriggerChannels
	sequenceChannels.UpdateChannels = updateChannels

	// Get a list of all the fixtures in the groups.
	fixturesConfig := fixture.LoadFixtures()

	// Create a channel to send events to the launchpad.
	eventsForLauchpad := make(chan common.ALight)

	// Now create a thread to handle those events.
	go launchpad.ListenAndSendToLaunchPad(eventsForLauchpad, pad)

	// Start off by turning off all of the Lights
	pad.Reset()

	// Create a channel to listen for buttons being pressed.
	buttonChannel := pad.Listen()

	// Build the default set of Pattens.
	pattens := patten.MakePatterns()
	sequence1 := sequence.CreateSequence("rgb", 0, pattens, fixturesConfig, sequenceChannels)
	sequence2 := sequence.CreateSequence("rgb", 1, pattens, fixturesConfig, sequenceChannels)
	sequence3 := sequence.CreateSequence("scanner", 2, pattens, fixturesConfig, sequenceChannels)
	sequence4 := sequence.CreateSequence("switch", 3, pattens, fixturesConfig, sequenceChannels)

	// Add Sequence to an array.
	sequences := []*common.Sequence{}
	sequences = append(sequences, &sequence1)
	sequences = append(sequences, &sequence2)
	sequences = append(sequences, &sequence3)
	sequences = append(sequences, &sequence4)

	// Create storage for the static color buttons.
	staticButton1 := common.StaticColorButton{}
	staticButton2 := common.StaticColorButton{}
	staticButton3 := common.StaticColorButton{}
	staticButton4 := common.StaticColorButton{}
	staticButton5 := common.StaticColorButton{}
	staticButton6 := common.StaticColorButton{}
	staticButton7 := common.StaticColorButton{}
	staticButton8 := common.StaticColorButton{}

	// Add the color buttons to an array.
	staticButtons := []common.StaticColorButton{}
	staticButtons = append(staticButtons, staticButton1)
	staticButtons = append(staticButtons, staticButton2)
	staticButtons = append(staticButtons, staticButton3)
	staticButtons = append(staticButtons, staticButton4)
	staticButtons = append(staticButtons, staticButton5)
	staticButtons = append(staticButtons, staticButton6)
	staticButtons = append(staticButtons, staticButton7)
	staticButtons = append(staticButtons, staticButton8)

	// soundTriggers is a an array of switches which control which sequence gets a music trigger.
	soundTriggers := []*common.Trigger{}
	soundTriggers = append(soundTriggers, &common.Trigger{SequenceNumber: 0, State: false})
	soundTriggers = append(soundTriggers, &common.Trigger{SequenceNumber: 1, State: false})
	soundTriggers = append(soundTriggers, &common.Trigger{SequenceNumber: 2, State: false})
	soundTriggers = append(soundTriggers, &common.Trigger{SequenceNumber: 3, State: false})

	// Create a sound trigger object and give it the sequences so it can access their configs.
	sound.NewSoundTrigger(soundTriggers, sequenceChannels)

	// Start threads for each sequence.
	go sequence.PlayNewSequence(sequence1, 0, pad, eventsForLauchpad, pattens, dmxController, fixturesConfig, sequenceChannels, soundTriggers)
	go sequence.PlayNewSequence(sequence2, 1, pad, eventsForLauchpad, pattens, dmxController, fixturesConfig, sequenceChannels, soundTriggers)
	go sequence.PlayNewSequence(sequence3, 2, pad, eventsForLauchpad, pattens, dmxController, fixturesConfig, sequenceChannels, soundTriggers)
	go sequence.PlayNewSequence(sequence4, 3, pad, eventsForLauchpad, pattens, dmxController, fixturesConfig, sequenceChannels, soundTriggers)

	// Light up any existing presets.
	presets.InitPresets(eventsForLauchpad, presetsStore)

	// Light the function buttons at the bottom.
	common.ShowFunctionButtons(sequence1, 7, eventsForLauchpad)

	// Initialize a ten length slice of empty slices for function buttons.
	functionButtons = make([][]bool, 9)
	// Initialize those 10 empty function button slices
	for i := 0; i < 9; i++ {
		functionButtons[i] = make([]bool, 9)
	}

	// Initialize a ten length slice of empty slices for select buttons.
	selectButtonPressed = make([]bool, 9)

	// Initialize a ten length slice of empty slices for function mode state.
	functionSelectMode = make([]bool, 9)

	// Initialize a ten length slice of empty slices for static lamps.
	staticLamps = make([][]bool, 9)
	// Initialize those 10 empty function button slices
	for i := 0; i < 9; i++ {
		staticLamps[i] = make([]bool, 9)
	}

	// Rememberwhen we've finished editing static colors.
	colorEditModeDone = make([]bool, 4)

	// Light the logo blue.
	pad.Light(8, -1, 0, 0, 255)

	// Light the clear button purple.
	common.LightOn(eventsForLauchpad, common.ALight{X: 0, Y: -1, Brightness: 255, Red: 200, Green: 0, Blue: 255})

	// Light the static color buttons.
	common.LightOn(eventsForLauchpad, common.ALight{X: 1, Y: -1, Brightness: 255, Red: 255, Green: 0, Blue: 0})
	common.LightOn(eventsForLauchpad, common.ALight{X: 2, Y: -1, Brightness: 255, Red: 0, Green: 255, Blue: 0})
	common.LightOn(eventsForLauchpad, common.ALight{X: 3, Y: -1, Brightness: 255, Red: 0, Green: 0, Blue: 255})

	// Light top functions.
	common.LightOn(eventsForLauchpad, common.ALight{X: 4, Y: -1, Brightness: 255, Red: 3, Green: 255, Blue: 255})
	common.LightOn(eventsForLauchpad, common.ALight{X: 5, Y: -1, Brightness: 255, Red: 3, Green: 255, Blue: 255})
	common.LightOn(eventsForLauchpad, common.ALight{X: 6, Y: -1, Brightness: 255, Red: 3, Green: 255, Blue: 255})
	common.LightOn(eventsForLauchpad, common.ALight{X: 7, Y: -1, Brightness: 255, Red: 3, Green: 255, Blue: 255})

	// Light the save, start, stop and blackout buttons.
	common.LightOn(eventsForLauchpad, common.ALight{X: 8, Y: 5, Brightness: full, Red: 255, Green: 255, Blue: 255})
	common.LightOn(eventsForLauchpad, common.ALight{X: 8, Y: 6, Brightness: full, Red: 255, Green: 255, Blue: 255})
	common.LightOn(eventsForLauchpad, common.ALight{X: 8, Y: 7, Brightness: full, Red: 255, Green: 255, Blue: 255})
	common.LightOn(eventsForLauchpad, common.ALight{X: 8, Y: 8, Brightness: full, Red: 255, Green: 255, Blue: 255})

	// Light the first sequence as the default selected.
	selectedSequence := 0
	common.SequenceSelect(eventsForLauchpad, selectedSequence)

	// Initialise the pattens.
	availablePattens := []string{}
	availablePattens = append(availablePattens, "standard")
	availablePattens = append(availablePattens, "inverted")
	availablePattens = append(availablePattens, "rgbchase")
	availablePattens = append(availablePattens, "pairs")
	availablePattens = append(availablePattens, "center")
	availablePattens = append(availablePattens, "colors")
	availablePattens = append(availablePattens, "fade")

	// Clear the pad.
	allFixturesOff(eventsForLauchpad, dmxController, fixturesConfig)

	// Main loop reading commands from the Novation Launchpad.
	for {

		hit := <-buttonChannel

		// Clear all the lights on the launchpad.
		if hit.X == 0 && hit.Y == -1 {
			launchpad.ClearAll(pad, presetsStore, eventsForLauchpad, commandChannels)
			allFixturesOff(eventsForLauchpad, dmxController, fixturesConfig)
			presets.ClearPresets(eventsForLauchpad, presetsStore)
			presets.InitPresets(eventsForLauchpad, presetsStore)

			// Make sure we stop all sequences.
			cmd := common.Command{
				Stop: true,
			}
			common.SendCommandToAllSequence(selectedSequence, cmd, commandChannels)

			// Swicth off any static colors.
			cmd = common.Command{
				UpdateStatic: false,
				Static:       false,
			}
			common.SendCommandToAllSequence(selectedSequence, cmd, commandChannels)

			// Clear the sequence colors.
			cmd = common.Command{
				ClearSequenceColor: true,
			}
			common.SendCommandToAllSequence(selectedSequence, cmd, commandChannels)

			continue
		}

		// Master brightness down.
		if hit.X == 6 && hit.Y == -1 {
			//fmt.Printf("Master Brightness Down %d\n", masterBrightness)
			masterBrightness = masterBrightness - 10
			if masterBrightness < 0 {
				masterBrightness = 0
			}
			cmd := common.Command{
				MasterBrightness: true,
				Master:           masterBrightness,
			}
			common.SendCommandToAllSequence(selectedSequence, cmd, commandChannels)
			continue
		}

		// Master brightness up.
		if hit.X == 7 && hit.Y == -1 {
			//fmt.Printf("Master Brightness Up %d\n", masterBrightness)
			masterBrightness = masterBrightness + 10
			if masterBrightness > 255 {
				masterBrightness = 255
			}
			cmd := common.Command{
				MasterBrightness: true,
				Master:           masterBrightness,
			}
			common.SendCommandToAllSequence(selectedSequence, cmd, commandChannels)
			continue
		}

		// Save mode.
		if hit.X == 8 && hit.Y == 4 {
			if savePreset {
				savePreset = false
				presets.InitPresets(eventsForLauchpad, presetsStore)
				common.LightOn(eventsForLauchpad, common.ALight{X: 8, Y: 4, Brightness: full, Red: 255, Green: 255, Blue: 255})
				continue
			}
			presets.InitPresets(eventsForLauchpad, presetsStore)
			launchpad.FlashLight(4, 8, 0x03, 0x5f, eventsForLauchpad)
			savePreset = true
			continue
		}

		// Ask all sequences for their current config and save in a file.
		//if hit.X < 8 && (hit.Y > 3 && hit.Y < 7) && !pad.IsBlocked() {
		if hit.X < 8 && (hit.Y > 3 && hit.Y < 7) {
			if savePreset {
				presetsStore[fmt.Sprint(hit.X)+","+fmt.Sprint(hit.Y)] = true
				common.LightOn(eventsForLauchpad, common.ALight{X: hit.X, Y: hit.Y, Brightness: full, Red: 255, Green: 0, Blue: 0})
				config.AskToSaveConfig(commandChannels, replyChannels, hit.X, hit.Y)
				savePreset = false
				launchpad.FlashLight(4, 8, 0, 0, eventsForLauchpad) // turn off the save button from flashing.
				presets.SavePresets(presetsStore)
				presets.ClearPresets(eventsForLauchpad, presetsStore)
				presets.InitPresets(eventsForLauchpad, presetsStore)
				launchpad.FlashLight(hit.Y, hit.X, 0x0d, 0x78, eventsForLauchpad)
			} else {
				// Load config, but only if it exists in the presets map.
				if presetsStore[fmt.Sprint(hit.X)+","+fmt.Sprint(hit.Y)] {

					// Stop all sequences, so we start in sync.
					cmd := common.Command{
						Stop: true,
					}
					common.SendCommandToAllSequence(selectedSequence, cmd, commandChannels)

					allFixturesOff(eventsForLauchpad, dmxController, fixturesConfig)
					time.Sleep(300 * time.Millisecond)

					// Load the config.
					config.AskToLoadConfig(commandChannels, hit.X, hit.Y)

					// Turn the selected preset light red.
					common.LightOn(eventsForLauchpad, common.ALight{X: hit.X, Y: hit.Y, Brightness: full, Red: 3, Green: 0, Blue: 0})
					presets.InitPresets(eventsForLauchpad, presetsStore)
					launchpad.FlashLight(hit.Y, hit.X, 0x0d, 0x78, eventsForLauchpad)

					// Preserve blackout.
					if !blackout {
						cmd := common.Command{
							Normal: true,
						}
						common.SendCommandToAllSequence(selectedSequence, cmd, commandChannels)
					}
				}
			}
			continue
		}

		// Decrement Patten.
		if hit.X == 2 && hit.Y == 7 {
			selectedPatten = selectedPatten - 1
			if selectedPatten < 0 {
				selectedPatten = 0
			}
			cmd := common.Command{
				UpdatePatten: true,
				Patten: common.Patten{
					Name: availablePattens[selectedPatten],
				},
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
			cmd = common.Command{
				Stop:  true,
				Speed: sequenceSpeed,
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

			cmd = common.Command{
				Start: true,
				Speed: sequenceSpeed,
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
			continue
		}

		// Increment Patten.
		if hit.X == 3 && hit.Y == 7 {
			selectedPatten = selectedPatten + 1
			if selectedPatten > len(availablePattens)-1 {
				selectedPatten = len(availablePattens) - 1
			}
			cmd := common.Command{
				Stop: true,
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

			cmd = common.Command{
				UpdatePatten: true,
				Patten: common.Patten{
					Name: availablePattens[selectedPatten],
				},
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

			cmd = common.Command{
				Stop:  true,
				Speed: sequenceSpeed,
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

			cmd = common.Command{
				Start: true,
				Speed: sequenceSpeed,
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
			continue
		}

		// Decrease speed of selected sequence.
		if hit.X == 0 && hit.Y == 7 {
			if !sequences[selectedSequence].MusicTrigger {
				sequenceSpeed--
				if sequenceSpeed < 0 {
					sequenceSpeed = 1
				}
				cmd := common.Command{
					Speed:       sequenceSpeed,
					UpdateSpeed: true,
				}
				common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
			}
			continue
		}

		// Increase speed of selected sequence.
		if hit.X == 1 && hit.Y == 7 {
			if !sequences[selectedSequence].MusicTrigger {
				sequenceSpeed++
				if sequenceSpeed > 21 {
					sequenceSpeed = 21
				}
				cmd := common.Command{
					Speed:       sequenceSpeed,
					UpdateSpeed: true,
				}
				common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
			}
			continue
		}

		// S E L E C T    S E Q U E N C E.
		// Select sequence 1.
		if hit.X == 8 && hit.Y == 0 {
			selectedSequence = 0
			HandleSelect(sequences, selectedSequence, eventsForLauchpad, selectButtonPressed, functionButtons,
				functionSelectMode, colorEditModeDone, commandChannels, sequenceChannels)
			setColorEditMode(selectedSequence, sequences, functionSelectMode, colorEditModeDone, selectButtonPressed, commandChannels, eventsForLauchpad)

			fmt.Printf("Seq %d  Static %t\n", selectedSequence, sequences[selectedSequence].Static)
			cmd := common.Command{
				PlayStaticOnce: true,
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

			continue
		}

		// Select sequence 2.
		if hit.X == 8 && hit.Y == 1 {
			selectedSequence = 1
			HandleSelect(sequences, selectedSequence, eventsForLauchpad, selectButtonPressed, functionButtons,
				functionSelectMode, colorEditModeDone, commandChannels, sequenceChannels)
			setColorEditMode(selectedSequence, sequences, functionSelectMode, colorEditModeDone, selectButtonPressed, commandChannels, eventsForLauchpad)

			cmd := common.Command{
				PlayStaticOnce: true,
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

			continue
		}

		// Select sequence 3.
		if hit.X == 8 && hit.Y == 2 {
			selectedSequence = 2
			HandleSelect(sequences, selectedSequence, eventsForLauchpad, selectButtonPressed, functionButtons,
				functionSelectMode, colorEditModeDone, commandChannels, sequenceChannels)
			setColorEditMode(selectedSequence, sequences, functionSelectMode, colorEditModeDone, selectButtonPressed, commandChannels, eventsForLauchpad)

			cmd := common.Command{
				PlayStaticOnce: true,
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

			continue
		}

		// Select sequence 4.
		if hit.X == 8 && hit.Y == 3 {
			selectedSequence = 3
			HandleSelect(sequences, selectedSequence, eventsForLauchpad, selectButtonPressed, functionButtons,
				functionSelectMode, colorEditModeDone, commandChannels, sequenceChannels)
			setColorEditMode(selectedSequence, sequences, functionSelectMode, colorEditModeDone, selectButtonPressed, commandChannels, eventsForLauchpad)

			cmd := common.Command{
				PlayStaticOnce: true,
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

			continue
		}

		// Start sequence.
		if hit.X == 8 && hit.Y == 5 {
			sequences[selectedSequence].MusicTrigger = false
			cmd := common.Command{
				Start: true,
				Speed: sequenceSpeed,
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
			common.LightOn(eventsForLauchpad, common.ALight{X: hit.X, Y: hit.Y, Brightness: full, Red: 255, Green: 0, Blue: 0})
			time.Sleep(100 * time.Millisecond)
			common.LightOn(eventsForLauchpad, common.ALight{X: hit.X, Y: hit.Y, Brightness: full, Red: 255, Green: 255, Blue: 255})
			continue
		}

		// Stop sequence.
		if hit.X == 8 && hit.Y == 6 {
			cmd := common.Command{
				Stop:  true,
				Speed: sequenceSpeed,
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
			common.LightOn(eventsForLauchpad, common.ALight{X: hit.X, Y: hit.Y, Brightness: full, Red: 255, Green: 0, Blue: 0})
			time.Sleep(100 * time.Millisecond)
			common.LightOn(eventsForLauchpad, common.ALight{X: hit.X, Y: hit.Y, Brightness: full, Red: 255, Green: 255, Blue: 255})
			continue
		}

		// Size decrease.
		if hit.X == 4 && hit.Y == 7 {
			size--
			if size < 0 {
				size = 0
			}
			// Send Update Fade Speed.
			cmd := common.Command{
				UpdateSize: true,
				Size:       size,
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
			continue
		}

		// Size increase.
		if hit.X == 5 && hit.Y == 7 {
			size++
			if size > 20 {
				size = 20
			}
			// Send size update.
			cmd := common.Command{
				UpdateSize: true,
				Size:       size,
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
			continue
		}

		// Fade time decrease.
		if hit.X == 6 && hit.Y == 7 {
			fadeSpeed--
			if fadeSpeed < 0 {
				fadeSpeed = 0
			}
			// Send fade update command.
			cmd := common.Command{
				DecreaseFade: true,
				FadeSpeed:    fadeSpeed,
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
			continue
		}

		// Fade time increase.
		if hit.X == 7 && hit.Y == 7 {
			fadeSpeed++
			if fadeSpeed > 20 {
				fadeSpeed = 20
			}
			// Send fade update command.
			cmd := common.Command{
				IncreaseFade: true,
				FadeSpeed:    fadeSpeed,
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
			continue
		}

		// F U N C T I O N  K E Y S
		if hit.X >= 0 && hit.X < 8 &&
			functionSelectMode[selectedSequence] &&
			!sequences[selectedSequence].Functions[common.Function5_Color].State {

			// Get an upto date copy of the sequence.
			cmd := common.Command{
				UpdateSequence: true,
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
			newSequence := <-updateChannels[selectedSequence]
			sequences[selectedSequence] = &newSequence

			// We clear first three function keys first. So that the toggle of the
			// chase modes will work.
			if hit.X < 3 {
				sequences[selectedSequence].Functions[common.Function1_Forward_Chase].State = false
				sequences[selectedSequence].Functions[common.Function2_Pairs_Chase].State = false
				sequences[selectedSequence].Functions[common.Function3_Inward_Chase].State = false
			}

			for _, functions := range sequences[selectedSequence].Functions {
				if hit.Y == functions.SequenceNumber {
					if !sequences[selectedSequence].Functions[hit.X].State {
						sequences[selectedSequence].Functions[hit.X].State = true
						break
					}
					if sequences[selectedSequence].Functions[hit.X].State {
						sequences[selectedSequence].Functions[hit.X].State = false
						break
					}
				}
			}

			// Send update functions command. This sets the temporary representation of
			// the function keys in the real sequence.
			cmd = common.Command{
				UpdateFunctions: true,
				Functions:       sequences[selectedSequence].Functions,
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

			// For the chase functions we only allow one at a time.
			common.SetFunctionKeys(sequences[selectedSequence].Functions, *sequences[selectedSequence])

			// Light the correct function key.
			common.ShowFunctionButtons(*sequences[selectedSequence], selectedSequence, eventsForLauchpad)

			continue
		}

		// S W I T C H   B U T T O N's Toggle State of switches for this sequence.
		if hit.X >= 0 && hit.X < 8 && !functionSelectMode[selectedSequence] &&
			hit.Y >= 0 &&
			hit.Y < 4 &&
			!sequences[selectedSequence].Functions[common.Function5_Color].State &&
			!sequences[selectedSequence].Functions[common.Function6_Static].State &&
			sequences[hit.Y].Type == "switch" {

			switchPositions[hit.Y][hit.X] = switchPositions[hit.Y][hit.X] + 1
			valuesLength := len(sequences[hit.Y].Switches[hit.X].Values)
			fmt.Printf("Length of Values %d\n", valuesLength)
			if switchPositions[hit.Y][hit.X] == valuesLength {
				switchPositions[hit.Y][hit.X] = 0
			}

			// Send a message to the sequence for it to toggle the selected switch.
			// hit.Y is the sequence.
			// hit.X is the switch.
			cmd := common.Command{
				UpdateSwitch:   true,
				SwitchNumber:   hit.X,
				SwitchPosition: switchPositions[hit.Y][hit.X],
			}
			common.SendCommandToSequence(hit.Y, cmd, commandChannels)

		}

		// F L A S H   B U T T O N S - Briefly light (flash) the fixtures based on current patten.
		if hit.X >= 0 && hit.X < 8 && !functionSelectMode[selectedSequence] &&
			hit.Y >= 0 &&
			hit.Y < 4 &&
			!sequences[selectedSequence].Functions[common.Function6_Static].State &&
			!sequences[selectedSequence].Functions[common.Function5_Color].State &&
			sequences[hit.Y].Type != "switch" {

			flashSequence := common.Sequence{
				Patten: common.Patten{
					Name:  "colors",
					Steps: pattens["colors"].Steps,
				},
			}

			red := flashSequence.Patten.Steps[hit.X].Fixtures[hit.X].Colors[0].R
			green := flashSequence.Patten.Steps[hit.X].Fixtures[hit.X].Colors[0].G
			blue := flashSequence.Patten.Steps[hit.X].Fixtures[hit.X].Colors[0].B
			pan := flashSequence.Patten.Steps[hit.X].Fixtures[hit.X].Pan
			tilt := flashSequence.Patten.Steps[hit.X].Fixtures[hit.X].Tilt
			shutter := flashSequence.Patten.Steps[hit.X].Fixtures[hit.X].Shutter
			gobo := flashSequence.Patten.Steps[hit.X].Fixtures[hit.X].Gobo

			common.LightOn(eventsForLauchpad, common.ALight{
				X:          hit.X,
				Y:          hit.Y,
				Brightness: masterBrightness,
				Red:        red,
				Green:      green,
				Blue:       blue,
			})
			fixture.MapFixtures(hit.Y, dmxController, hit.X, red, green, blue, pan, tilt, shutter, gobo, fixturesConfig, blackout, masterBrightness, masterBrightness)
			time.Sleep(200 * time.Millisecond)
			common.LightOff(eventsForLauchpad, hit.X, hit.Y)
			fixture.MapFixtures(hit.Y, dmxController, hit.X, 0, 0, 0, pan, tilt, shutter, gobo, fixturesConfig, blackout, masterBrightness, masterBrightness)
			continue
		}

		// C H O O S E   S T A T I C    C O L O R
		// Red
		if hit.X == 1 && hit.Y == -1 {

			staticButtons[selectedSequence].X = lastStaticColorButtonX
			staticButtons[selectedSequence].Y = lastStaticColorButtonY

			if staticButtons[selectedSequence].Color.R > 254 {
				staticButtons[selectedSequence].Color.R = 0
			} else {
				staticButtons[selectedSequence].Color.R = staticButtons[selectedSequence].Color.R + 10
			}
			common.LightOn(eventsForLauchpad, common.ALight{X: hit.X, Y: hit.Y, Brightness: full, Red: staticButtons[selectedSequence].Color.R, Green: 0, Blue: 0})
			updateStaticLamp(selectedSequence, staticButtons[selectedSequence], commandChannels)
			continue
		}

		// Green
		if hit.X == 2 && hit.Y == -1 {

			staticButtons[selectedSequence].X = lastStaticColorButtonX
			staticButtons[selectedSequence].Y = lastStaticColorButtonY

			if staticButtons[selectedSequence].Color.G > 254 {
				staticButtons[selectedSequence].Color.G = 0
			} else {
				staticButtons[selectedSequence].Color.G = staticButtons[selectedSequence].Color.G + 10
			}
			common.LightOn(eventsForLauchpad, common.ALight{X: hit.X, Y: hit.Y, Brightness: full, Red: 0, Green: staticButtons[selectedSequence].Color.G, Blue: 0})
			updateStaticLamp(selectedSequence, staticButtons[selectedSequence], commandChannels)
			continue
		}

		// Blue
		if hit.X == 3 && hit.Y == -1 {

			staticButtons[selectedSequence].X = lastStaticColorButtonX
			staticButtons[selectedSequence].Y = lastStaticColorButtonY

			if staticButtons[selectedSequence].Color.B > 254 {
				staticButtons[selectedSequence].Color.B = 0
			} else {
				staticButtons[selectedSequence].Color.B = staticButtons[selectedSequence].Color.B + 10
			}
			common.LightOn(eventsForLauchpad, common.ALight{X: hit.X, Y: hit.Y, Brightness: full, Red: 0, Green: 0, Blue: staticButtons[selectedSequence].Color.B})
			updateStaticLamp(selectedSequence, staticButtons[selectedSequence], commandChannels)
			continue
		}

		// S E T    S E Q U E N C E   C O L O R
		if hit.X >= 0 && hit.X < 8 && hit.Y != -1 &&
			sequences[selectedSequence].Functions[common.Function5_Color].State {

			fmt.Printf("S E T    S E Q E N C E   C O L O R for X:%d Y%d\n", hit.X, hit.Y)
			cmd := common.Command{
				UpdateSequenceColor: true,
				SelectedColor:       hit.X,
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

			// Turn off the edit colors bar.
			sequences[selectedSequence].Functions[common.Function5_Color].State = false
			cmd = common.Command{
				UpdateFunctions: true,
				Functions:       sequences[selectedSequence].Functions,
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

			// Restart the sequence.
			cmd = common.Command{
				Start: true,
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

			// And reveal the sequence on the launchpad keys
			common.RevealSequence(selectedSequence, commandChannels)
			// Turn off the function mode flag.
			functionSelectMode[selectedSequence] = false
			// Now forget we pressed twice and start again.
			selectButtonPressed[selectedSequence] = true

			HideColorSelectionButtons(selectedSequence, *sequences[selectedSequence], selectedSequence, eventsForLauchpad)

			continue
		}

		// S E T    S T A T I C   C O L O R
		if hit.X >= 0 && hit.X < 8 && !functionSelectMode[selectedSequence] && hit.Y != -1 &&
			sequences[selectedSequence].Functions[common.Function6_Static].State {

			// For this button increment the color.
			sequences[selectedSequence].StaticColors[hit.X].X = hit.X
			sequences[selectedSequence].StaticColors[hit.X].Y = hit.Y
			sequences[selectedSequence].StaticColors[hit.X].SelectedColor++
			if sequences[selectedSequence].StaticColors[hit.X].SelectedColor > 10 {
				sequences[selectedSequence].StaticColors[hit.X].SelectedColor = 0
			}
			sequences[selectedSequence].StaticColors[hit.X].Color = common.GetColorButtonsArray(sequences[selectedSequence].StaticColors[hit.X].SelectedColor)

			// Store the color data to allow for editing of static colors.
			staticButtons[selectedSequence].X = hit.X
			staticButtons[selectedSequence].Y = hit.Y
			staticButtons[selectedSequence].Color = sequences[selectedSequence].StaticColors[hit.X].Color
			staticButtons[selectedSequence].SelectedColor = sequences[selectedSequence].StaticColors[hit.X].SelectedColor

			// Tell the sequence about the new color and where we are in the
			// color cycle.
			cmd := common.Command{
				UpdateStaticColor: true,
				Static:            true,
				StaticLamp:        hit.X,
				StaticLampFlash:   false,
				SelectedColor:     sequences[selectedSequence].StaticColors[hit.X].SelectedColor,
				StaticColor:       sequences[selectedSequence].StaticColors[hit.X].Color,
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
			lastStaticColorButtonX = hit.X
			lastStaticColorButtonY = hit.Y
			fmt.Printf("LastStaticColorButton X:%d Y:%d\n", lastStaticColorButtonX, lastStaticColorButtonY)

			continue
		}

		// B L A C K O U T   B U T T O N.
		if hit.X == 8 && hit.Y == 7 {
			if !blackout {
				blackout = true
				cmd := common.Command{
					Blackout: true,
				}
				common.SendCommandToAllSequence(selectedSequence, cmd, commandChannels)
				common.LightOn(eventsForLauchpad, common.ALight{X: hit.X, Y: hit.Y, Brightness: full, Red: 0, Green: 0, Blue: 255})
			} else {
				blackout = false
				cmd := common.Command{
					Normal: true,
				}
				common.SendCommandToAllSequence(selectedSequence, cmd, commandChannels)
				common.LightOn(eventsForLauchpad, common.ALight{X: hit.X, Y: hit.Y, Brightness: full, Red: 255, Green: 255, Blue: 255})
			}
			continue
		}
	}
}

func updateStaticLamp(selectedSequence int, staticColorButtons common.StaticColorButton, commandChannels []chan common.Command) {

	// Static is set to true in the functions and this key is set to
	// the selected color.
	cmd := common.Command{
		UpdateStaticColor: true,
		Static:            true,
		StaticLamp:        staticColorButtons.X,
		SelectedColor:     staticColorButtons.SelectedColor,
		StaticColor: common.Color{
			R: staticColorButtons.Color.R,
			G: staticColorButtons.Color.G,
			B: staticColorButtons.Color.B,
		},
	}
	common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

}

func HandleSelect(sequences []*common.Sequence,
	selectedSequence int,
	eventsForLauchpad chan common.ALight,
	selectButtonPressed []bool,
	functionButtons [][]bool,
	functionSelectMode []bool,
	colorEditModeDone []bool,
	commandChannels []chan common.Command,
	channels common.Channels) {

	// fmt.Printf("HANDLE: selectButtons[%d] = %t \n", selectedSequence, selectButtonPressed[selectedSequence])
	// fmt.Printf("HANDLE: colorEditModeDone[%d] = %t \n", selectedSequence, colorEditModeDone[selectedSequence])
	// fmt.Printf("HANDLE: functionSelectMode[%d] = %t \n", selectedSequence, functionSelectMode[selectedSequence])
	// fmt.Printf("HANDLE: Func Static[%d] = %t\n", selectedSequence, sequences[selectedSequence].Functions[common.Function6_Static].State)

	// Light the sequence selector button.
	common.SequenceSelect(eventsForLauchpad, selectedSequence)

	// First time into function mode we head back to normal mode.
	if functionSelectMode[selectedSequence] && !selectButtonPressed[selectedSequence] {
		//fmt.Printf("Handle 1 Function Bar off\n")
		// Turn off function mode. Remove the function pads.
		common.HideFunctionButtons(selectedSequence, eventsForLauchpad)

		if sequences[selectedSequence].Functions[common.Function6_Static].State {
			common.SetMode(selectedSequence, commandChannels, "Static")
		}

		if sequences[selectedSequence].Functions[common.Function5_Color].State {
			ShowColorSelectionButtons(selectedSequence, *sequences[selectedSequence], selectedSequence, eventsForLauchpad)
		} else {
			// And reveal the sequence on the launchpad keys
			common.RevealSequence(selectedSequence, commandChannels)
			// Turn off the function mode flag.
			functionSelectMode[selectedSequence] = false
			// Now forget we pressed twice and start again.
			selectButtonPressed[selectedSequence] = true
		}
		return
	}

	// This the first time we have pressed the select button.
	if !selectButtonPressed[selectedSequence] {
		//fmt.Printf("Handle 2\n")
		// assume everything else is off.
		selectButtonPressed[0] = false
		selectButtonPressed[1] = false
		selectButtonPressed[2] = false
		selectButtonPressed[3] = false
		// But remember we have pressed this select button once.
		functionSelectMode[selectedSequence] = false
		selectButtonPressed[selectedSequence] = true
		return
	}

	// Are we in function mode ?
	if functionSelectMode[selectedSequence] {
		//fmt.Printf("Handle 3\n")
		// Turn off function mode. Remove the function pads.
		common.HideFunctionButtons(selectedSequence, eventsForLauchpad)
		// And reveal the sequence on the launchpad keys
		common.RevealSequence(selectedSequence, commandChannels)
		// Turn off the function mode flag.
		functionSelectMode[selectedSequence] = false
		// Now forget we pressed twice and start again.
		selectButtonPressed[selectedSequence] = false
		return
	}

	// We are in function mode for this sequence.
	if !functionSelectMode[selectedSequence] {
		//fmt.Printf("Handle 4 - Function Bar On!\n")

		// fmt.Printf("Color Edit Mode Done set for %t\n", colorEditModeDone[selectedSequence])

		// Set function mode.
		functionSelectMode[selectedSequence] = true

		// And hide the sequence so we can only see the function buttons.
		common.HideSequence(selectedSequence, commandChannels)

		// Create the function buttons.
		common.MakeFunctionButtons(*sequences[selectedSequence], selectedSequence, eventsForLauchpad, functionButtons, channels)
		// Now forget we pressed twice and start again.
		selectButtonPressed[selectedSequence] = false
		return
	}
}

func setColorEditMode(selectedSequence int,
	sequences []*common.Sequence,
	functionSelectMode []bool,
	colorEditModeDone []bool,
	selectButtonPressed []bool,
	commandChannels []chan common.Command,
	eventsForLauchpad chan common.ALight) {

	if functionSelectMode[selectedSequence] && sequences[selectedSequence].Functions[common.Function6_Static].State && colorEditModeDone[selectedSequence] {

		// fmt.Printf("Hide Edit Color Buttons.\n")
		showEditColorButtons(sequences[selectedSequence], selectedSequence, eventsForLauchpad, false, sequences[selectedSequence].Master)

		cmd := common.Command{
			SetEditColors: true,
			EditColors:    true,
		}
		common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

		colorEditModeDone[selectedSequence] = false
		selectButtonPressed[selectedSequence] = false
		functionSelectMode[selectedSequence] = false
		sequences[selectedSequence].Functions[common.Function6_Static].State = false

		common.RevealSequence(selectedSequence, commandChannels)
	}

	if !functionSelectMode[selectedSequence] && sequences[selectedSequence].Functions[common.Function6_Static].State && !colorEditModeDone[selectedSequence] {
		cmd := common.Command{
			SetEditColors: true,
			EditColors:    false,
		}
		common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
		if !colorEditModeDone[selectedSequence] {
			//fmt.Printf("Show Edit Color Buttons.\n")
			showEditColorButtons(sequences[selectedSequence], selectedSequence, eventsForLauchpad, true, sequences[selectedSequence].Master)
		}
	}
}

func allFixturesOff(eventsForLauchpad chan common.ALight, dmxController *ft232.DMXController, fixturesConfig *fixture.Fixtures) {
	for x := 0; x < 8; x++ {
		for y := 0; y < 4; y++ {
			common.LightOff(eventsForLauchpad, x, y)
			fixture.MapFixtures(y, dmxController, x, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, true, 0, 0)
		}
	}
}

// For the given sequence show the static colors on the relevant buttons.
func showEditColorButtons(sequence *common.Sequence, selectedSequence int, eventsForLauchpad chan common.ALight, show bool, master int) {
	for index, color := range sequence.StaticColors {
		launchpad.LightLamp(selectedSequence, index, color.Color.R, color.Color.G, color.Color.B, master, eventsForLauchpad)
	}
}

// For the given sequence show the available sequence colors on the relevant buttons.
func ShowColorSelectionButtons(mySequenceNumber int, sequence common.Sequence, selectedSequence int, eventsForLauchpad chan common.ALight) {
	fmt.Printf("ShowColorSelectionButtons\n")
	for myFixtureNumber, lamp := range sequence.AvailableSequenceColors {
		launchpad.LightLamp(mySequenceNumber, myFixtureNumber, lamp.Color.R, lamp.Color.G, lamp.Color.B, sequence.Master, eventsForLauchpad)
	}
}

// For the given sequence hide the available sequence colors..
func HideColorSelectionButtons(mySequenceNumber int, sequence common.Sequence, selectedSequence int, eventsForLauchpad chan common.ALight) {
	fmt.Printf("ShowColorSelectionButtons\n")
	for myFixtureNumber := range sequence.AvailableSequenceColors {
		launchpad.LightLamp(mySequenceNumber, myFixtureNumber, 0, 0, 0, sequence.Master, eventsForLauchpad)
	}
}
