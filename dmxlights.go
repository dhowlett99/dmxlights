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
	"github.com/dhowlett99/dmxlights/pkg/launchpad"
	"github.com/dhowlett99/dmxlights/pkg/patten"
	"github.com/dhowlett99/dmxlights/pkg/presets"
	"github.com/dhowlett99/dmxlights/pkg/sequence"
	"github.com/dhowlett99/dmxlights/pkg/sound"
	"github.com/oliread/usbdmx/ft232"
	"github.com/rakyll/launchpad/mk3"
)

const (
	full = 255
)

var sequenceSpeed int = 12
var fadeSpeed int
var size int
var sequenceSize int = 60
var masterBrightness int

//var color int
var savePreset bool
var selectedPatten = 0
var blackout bool = false
var flood bool = false

// main thread is used to get commands from the lauchpad.
func main() {

	var functionButtons [][]bool   // Function buttons.
	var functionSelectMode []bool  // Which sequence is in function selection mode.
	var selectButtonPressed []bool // Which sequence has its Select button pressed.
	var staticLamps [][]bool       // Static color lamps.
	var switchPositions [9][9]int  // Sorage for switch positions.
	var colorEditMode []bool       // This flag is true when the sequence is in color editing mode.
	fadeSpeed = 11                 // Default start at 50ms.
	masterBrightness = 255         // Affects all DMX fixtures and launchpad lamps.
	var lastStaticColorButtonX int // Which Static Color button did we change last.
	var lastStaticColorButtonY int // Which Static Color button did we change last.
	var soundGain float32 = 0      // Fine gain -0.09 -> 0.09
	var autocolor []bool           // Is auto color change selected for this sequence.
	var autopatten []bool          // Is auto patten change selected for this sequence.

	// Make an empty presets store.
	presetsStore := make(map[string]bool)

	// Make a store for which sequences can be flood light.
	selectedFloodMap := make(map[int]bool, 4)

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

	// Read sequences config file
	sequencesConfig, err := sequence.LoadSequences()
	if err != nil {
		fmt.Printf("dmxlights: error failed to load sequences config: %s\n", err.Error())
		os.Exit(1)
	}

	// Get a list of all the fixtures in the groups.
	fixturesConfig, err := fixture.LoadFixtures()
	if err != nil {
		fmt.Printf("dmxlights: error failed to load fixtures: %s\n", err.Error())
		os.Exit(1)
	}

	// Find scanner GOBO's
	for _, f := range fixturesConfig.Fixtures {
		fmt.Printf("Found fixture: %s, group: %d, desc: %s\n", f.Name, f.Group, f.Description)
		if f.Type == "scanner" {
			gobos := fixture.HowManyGobos(fixturesConfig, f)
			for _, gobo := range gobos {
				fmt.Printf("Fixture %s Number of Gobos %s\n", f.Name, gobo.Name)
			}
		}
	}

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

	// Create the sequences from config file.
	// Add Sequence to an array.
	sequences := []*common.Sequence{}
	for index, sequenceConf := range sequencesConfig.Sequences {
		fmt.Printf("Found sequence: %s, desc: %s, type: %s\n", sequenceConf.Name, sequenceConf.Description, sequenceConf.Type)
		if sequenceConf.Type == "rgb" {
			selectedFloodMap[index] = true // This sequence is flood able because it's a rgb.
		}
		tempSequence := sequence.CreateSequence(sequenceConf.Type, index, pattens, fixturesConfig, sequenceChannels, selectedFloodMap)
		sequences = append(sequences, &tempSequence)
	}

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
	soundTriggers = append(soundTriggers, &common.Trigger{SequenceNumber: 0, State: false, Gain: soundGain})
	soundTriggers = append(soundTriggers, &common.Trigger{SequenceNumber: 1, State: false, Gain: soundGain})
	soundTriggers = append(soundTriggers, &common.Trigger{SequenceNumber: 2, State: false, Gain: soundGain})
	soundTriggers = append(soundTriggers, &common.Trigger{SequenceNumber: 3, State: false, Gain: soundGain})

	// Create a sound trigger object and give it the sequences so it can access their configs.
	sound.NewSoundTrigger(soundTriggers, sequenceChannels)

	// Start threads for each sequence.
	go sequence.PlayNewSequence(*sequences[0], 0, pad, eventsForLauchpad, pattens, dmxController, fixturesConfig, sequenceChannels, soundTriggers)
	go sequence.PlayNewSequence(*sequences[1], 1, pad, eventsForLauchpad, pattens, dmxController, fixturesConfig, sequenceChannels, soundTriggers)
	go sequence.PlayNewSequence(*sequences[2], 2, pad, eventsForLauchpad, pattens, dmxController, fixturesConfig, sequenceChannels, soundTriggers)
	go sequence.PlayNewSequence(*sequences[3], 3, pad, eventsForLauchpad, pattens, dmxController, fixturesConfig, sequenceChannels, soundTriggers)

	// Light up any existing presets.
	presets.InitPresets(eventsForLauchpad, presetsStore)

	// Light the function buttons at the bottom.
	common.ShowFunctionButtons(*sequences[1], 7, eventsForLauchpad)

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

	// Remember when we've finished editing static colors.
	colorEditMode = make([]bool, 4)

	// Remember when we have set autocolor change for this sequence.
	autocolor = make([]bool, 4)

	// Remember when we have set autocolor change for this sequence.
	autopatten = make([]bool, 4)

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

	// Initialise the flood button to be green.
	common.LightOn(eventsForLauchpad, common.ALight{X: 7, Y: 3, Brightness: full, Red: 255, Green: 0, Blue: 0})

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

			// Turn off the flood
			if flood {
				cmd := common.Command{
					UpdateFlood: true,
					Flood:       false,
				}
				common.SendCommandToAllSequence(selectedSequence, cmd, commandChannels)
				flood = false
			}

			// We want to clear a color selection.
			if sequences[selectedSequence].Functions[common.Function5_Color].State &&
				sequences[selectedSequence].Type != "scanner" {

				// Clear the sequence colors for this sequence.
				cmd := common.Command{
					ClearSequenceColor: true,
				}
				common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

				// Get an upto date copy of the sequence.
				sequences[selectedSequence] = common.RefreshSequence(selectedSequence, commandChannels, updateChannels)

				// Flash the correct function buttons
				ShowColorSelectionButtons(selectedSequence, *sequences[selectedSequence], eventsForLauchpad)

				continue
			}

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

			// refesh the switch positions.
			cmd = common.Command{
				UpdateSwitchPositions: true,
			}
			common.SendCommandToAllSequenceOfType(sequences, selectedSequence, cmd, commandChannels, "switch")

			// Clear the sequence colors.
			cmd = common.Command{
				ClearSequenceColor: true,
			}
			common.SendCommandToAllSequence(selectedSequence, cmd, commandChannels)

			continue
		}

		// A U T O   C O L O R   M O D E
		if hit.X == 6 && hit.Y == 3 {
			if !autocolor[selectedSequence] {
				autocolor[selectedSequence] = true
				cmd := common.Command{
					UpdateAutoColor: true,
					AutoColor:       true,
				}
				common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
				common.LightOn(eventsForLauchpad, common.ALight{X: hit.X, Y: hit.Y, Brightness: full, Red: 255, Green: 0, Blue: 0})
				time.Sleep(300 * time.Millisecond)
				common.LightOn(eventsForLauchpad, common.ALight{X: hit.X, Y: hit.Y, Brightness: full, Red: 0, Green: 255, Blue: 0})
				continue
			}
			if autocolor[selectedSequence] {
				autocolor[selectedSequence] = false
				cmd := common.Command{
					UpdateAutoColor: true,
					AutoColor:       false,
				}
				common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
				common.LightOn(eventsForLauchpad, common.ALight{X: hit.X, Y: hit.Y, Brightness: full, Red: 255, Green: 0, Blue: 0})
				time.Sleep(300 * time.Millisecond)
				common.LightOn(eventsForLauchpad, common.ALight{X: hit.X, Y: hit.Y, Brightness: full, Red: 0, Green: 255, Blue: 0})
				continue
			}
		}

		// A U T O   P A T T E N   M O D E
		if hit.X == 5 && hit.Y == 3 {
			if !autopatten[selectedSequence] {
				autopatten[selectedSequence] = true
				cmd := common.Command{
					UpdateAutoPatten: true,
					AutoPatten:       true,
				}
				common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
				common.LightOn(eventsForLauchpad, common.ALight{X: hit.X, Y: hit.Y, Brightness: full, Red: 255, Green: 0, Blue: 0})
				time.Sleep(300 * time.Millisecond)
				common.LightOn(eventsForLauchpad, common.ALight{X: hit.X, Y: hit.Y, Brightness: full, Red: 0, Green: 255, Blue: 0})
				continue
			}
			if autopatten[selectedSequence] {
				autopatten[selectedSequence] = false
				cmd := common.Command{
					UpdateAutoPatten: true,
					AutoPatten:       false,
				}
				common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
				common.LightOn(eventsForLauchpad, common.ALight{X: hit.X, Y: hit.Y, Brightness: full, Red: 255, Green: 0, Blue: 0})
				time.Sleep(300 * time.Millisecond)
				common.LightOn(eventsForLauchpad, common.ALight{X: hit.X, Y: hit.Y, Brightness: full, Red: 0, Green: 255, Blue: 0})
				continue
			}
		}

		// F L O O D
		if hit.X == 7 && hit.Y == 3 {
			if !flood {
				cmd := common.Command{
					UpdateFlood: true,
					Flood:       true,
				}
				common.SendCommandToAllSequence(selectedSequence, cmd, commandChannels)

				// Wait for sequence to pause.
				time.Sleep(500 * time.Millisecond)

				flood = true
				sequences[selectedSequence].Flood = true
				sequences[selectedSequence].PlayFloodOnce = true
				common.LightOn(eventsForLauchpad, common.ALight{X: hit.X, Y: hit.Y, Brightness: full, Red: 255, Green: 255, Blue: 255})
				sequence.Flood(sequences[selectedSequence], dmxController, eventsForLauchpad, fixturesConfig, true)

				continue
			}
			if flood {
				cmd := common.Command{
					UpdateFlood: true,
					Flood:       false,
				}
				common.SendCommandToAllSequence(selectedSequence, cmd, commandChannels)
				common.LightOn(eventsForLauchpad, common.ALight{X: hit.X, Y: hit.Y, Brightness: full, Red: 0, Green: 255, Blue: 0})
				flood = false
				sequences[selectedSequence].Flood = false
				sequences[selectedSequence].PlayFloodOnce = true
				sequence.Flood(sequences[selectedSequence], dmxController, eventsForLauchpad, fixturesConfig, true)
				continue
			}
		}

		// Sound sensitity up.
		if hit.X == 4 && hit.Y == -1 {
			fmt.Printf("Sound Up %f\n", soundGain)
			soundGain = soundGain - 0.01
			if soundGain < -0.9 {
				soundGain = -0.9
			}
			for _, trigger := range soundTriggers {
				trigger.Gain = soundGain
			}
		}

		// Sound sensitity down.
		if hit.X == 5 && hit.Y == -1 {
			fmt.Printf("Sound Down%f\n", soundGain)
			soundGain = soundGain + 0.01
			if soundGain > 0.9 {
				soundGain = 0.9
			}
			for _, trigger := range soundTriggers {
				trigger.Gain = soundGain
			}
		}

		// Master brightness down.
		if hit.X == 6 && hit.Y == -1 {
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
			// Get an upto date copy of the sequence.
			sequences[selectedSequence] = common.RefreshSequence(selectedSequence, commandChannels, updateChannels)

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
			// Get an upto date copy of the sequence.
			sequences[selectedSequence] = common.RefreshSequence(selectedSequence, commandChannels, updateChannels)

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
				functionSelectMode, colorEditMode, commandChannels, sequenceChannels)
			setColorEditMode(selectedSequence, sequences, functionSelectMode, colorEditMode, selectButtonPressed, commandChannels, eventsForLauchpad)

			cmd := common.Command{
				PlayStaticOnce: true,
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

			colorEditMode[selectedSequence] = false
			continue
		}

		// Select sequence 2.
		if hit.X == 8 && hit.Y == 1 {
			selectedSequence = 1
			HandleSelect(sequences, selectedSequence, eventsForLauchpad, selectButtonPressed, functionButtons,
				functionSelectMode, colorEditMode, commandChannels, sequenceChannels)
			setColorEditMode(selectedSequence, sequences, functionSelectMode, colorEditMode, selectButtonPressed, commandChannels, eventsForLauchpad)

			cmd := common.Command{
				PlayStaticOnce: true,
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

			colorEditMode[selectedSequence] = false
			continue
		}

		// Select sequence 3.
		if hit.X == 8 && hit.Y == 2 {
			selectedSequence = 2
			HandleSelect(sequences, selectedSequence, eventsForLauchpad, selectButtonPressed, functionButtons,
				functionSelectMode, colorEditMode, commandChannels, sequenceChannels)
			setColorEditMode(selectedSequence, sequences, functionSelectMode, colorEditMode, selectButtonPressed, commandChannels, eventsForLauchpad)

			cmd := common.Command{
				PlayStaticOnce: true,
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

			colorEditMode[selectedSequence] = false
			continue
		}

		// Select sequence 4.
		if hit.X == 8 && hit.Y == 3 {
			selectedSequence = 3
			HandleSelect(sequences, selectedSequence, eventsForLauchpad, selectButtonPressed, functionButtons,
				functionSelectMode, colorEditMode, commandChannels, sequenceChannels)
			setColorEditMode(selectedSequence, sequences, functionSelectMode, colorEditMode, selectButtonPressed, commandChannels, eventsForLauchpad)

			cmd := common.Command{
				PlayStaticOnce: true,
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

			colorEditMode[selectedSequence] = false
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

			sequenceSize = sequenceSize - 10
			if sequenceSize < 0 {
				sequenceSize = 0
			}
			// Send Update Fade Speed.
			cmd = common.Command{
				UpdateSequenceSize: true,
				SequenceSize:       sequenceSize,
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

			continue
		}

		// Size increase.
		if hit.X == 5 && hit.Y == 7 {
			// Update the PAR can size.
			size++
			if size > 25 {
				size = 25
			}
			cmd := common.Command{
				UpdateSize: true,
				Size:       size,
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

			// Update the Scanner size.
			sequenceSize = sequenceSize + 10
			if sequenceSize > 120 {
				sequenceSize = 120
			}
			cmd = common.Command{
				UpdateSequenceSize: true,
				SequenceSize:       sequenceSize,
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
			sequences[selectedSequence] = common.RefreshSequence(selectedSequence, commandChannels, updateChannels)

			// We clear first three function keys first. So that the toggle of the
			// chase modes will work.
			if hit.X < 3 {
				sequences[selectedSequence].Functions[common.Function1_Forward_Chase].State = false
				sequences[selectedSequence].Functions[common.Function2_Pairs_Chase].State = false
				sequences[selectedSequence].Functions[common.Function3_Inward_Chase].State = false
			}

			if hit.X < 4 && sequences[selectedSequence].Type == "scanner" {
				sequences[selectedSequence].Functions[common.Function1_Forward_Chase].State = false
				sequences[selectedSequence].Functions[common.Function2_Pairs_Chase].State = false
				sequences[selectedSequence].Functions[common.Function3_Inward_Chase].State = false
				sequences[selectedSequence].Functions[common.Function4_Bounce].State = false
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
			cmd := common.Command{
				UpdateFunctions: true,
				Functions:       sequences[selectedSequence].Functions,
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

			// For the chase functions we only allow one at a time.
			common.SetFunctionKeys(sequences[selectedSequence].Functions, *sequences[selectedSequence])

			// Light the correct function key.
			common.ShowFunctionButtons(*sequences[selectedSequence], selectedSequence, eventsForLauchpad)

			// TODO find a way to get instant patten changes
			// without stopping and starting sequences.
			// cmd = common.Command{
			// 	Stop: true,
			// }
			// common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

			// cmd = common.Command{
			// 	Start: true,
			// }
			// common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

			continue
		}

		// S W I T C H   B U T T O N's Toggle State of switches for this sequence.
		if hit.X >= 0 && hit.X < 8 && !functionSelectMode[selectedSequence] &&
			hit.Y >= 0 &&
			hit.Y < 4 &&
			!sequences[selectedSequence].Functions[common.Function5_Color].State &&
			!sequences[selectedSequence].Functions[common.Function6_Static].State &&
			sequences[hit.Y].Type == "switch" {

			// We have a valid switch.
			if hit.X < len(sequences[hit.Y].Switches) {
				switchPositions[hit.Y][hit.X] = switchPositions[hit.Y][hit.X] + 1
				valuesLength := len(sequences[hit.Y].Switches[hit.X].States)
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
				// Send a message to the switch sequence.
				common.SendCommandToSequence(hit.Y, cmd, commandChannels)
			}
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
			sequences[selectedSequence].Functions[common.Function5_Color].State &&
			sequences[selectedSequence].Type != "scanner" {

			// Add the selected color to the sequence.
			cmd := common.Command{
				UpdateSequenceColor: true,
				SelectedColor:       hit.X,
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

			colorEditMode[selectedSequence] = true

			// Get an upto date copy of the sequence.
			sequences[selectedSequence] = common.RefreshSequence(selectedSequence, commandChannels, updateChannels)

			// Set the colors.
			sequences[selectedSequence].CurrentSequenceColors = sequences[selectedSequence].SequenceColors

			// We call ShowColorSelectionButtons here so the selections will flash as you press them.
			ShowColorSelectionButtons(selectedSequence, *sequences[selectedSequence], eventsForLauchpad)

			continue
		}

		// S E T    S C A N N E R   G O B O
		if hit.X >= 0 && hit.X < 8 && hit.Y != -1 &&
			sequences[selectedSequence].Functions[common.Function5_Color].State &&
			sequences[selectedSequence].Type == "scanner" {

			// Add the selected color to the sequence.
			cmd := common.Command{
				UpdateGobo:   true,
				SelectedGobo: hit.X,
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

			colorEditMode[selectedSequence] = true

			// Get an upto date copy of the sequence.
			sequences[selectedSequence] = common.RefreshSequence(selectedSequence, commandChannels, updateChannels)

			// Set the colors.
			sequences[selectedSequence].CurrentSequenceColors = sequences[selectedSequence].SequenceColors

			// We call ShowColorSelectionButtons here so the selections will flash as you press them.
			ShowColorSelectionButtons(selectedSequence, *sequences[selectedSequence], eventsForLauchpad)

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
	colorEditMode []bool,
	commandChannels []chan common.Command,
	channels common.Channels) {

	// fmt.Printf("HANDLE: selectButtons[%d] = %t \n", selectedSequence, selectButtonPressed[selectedSequence])
	// fmt.Printf("HANDLE: colorEditMode[%d] = %t \n", selectedSequence, colorEditMode[selectedSequence])
	// fmt.Printf("HANDLE: functionSelectMode[%d] = %t \n", selectedSequence, functionSelectMode[selectedSequence])
	// fmt.Printf("HANDLE: Func Static[%d] = %t\n", selectedSequence, sequences[selectedSequence].Functions[common.Function6_Static].State)

	// Light the sequence selector button.
	common.SequenceSelect(eventsForLauchpad, selectedSequence)

	// First time into function mode we head back to normal mode.
	if functionSelectMode[selectedSequence] && !selectButtonPressed[selectedSequence] && !colorEditMode[selectedSequence] {
		//fmt.Printf("Handle 1 Function Bar off\n")
		// Turn off function mode. Remove the function pads.
		common.HideFunctionButtons(selectedSequence, eventsForLauchpad)

		if sequences[selectedSequence].Functions[common.Function6_Static].State {
			common.SetMode(selectedSequence, commandChannels, "Static")
		}

		if sequences[selectedSequence].Functions[common.Function5_Color].State {
			ShowColorSelectionButtons(selectedSequence, *sequences[selectedSequence], eventsForLauchpad)
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
	if !functionSelectMode[selectedSequence] &&
		sequences[selectedSequence].Type != "switch" { // Don't alow functions in switch mode.

		//fmt.Printf("Handle 4 - Function Bar On!\n")

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

func unSetColorEditMode(selectedSequence int,
	sequences []*common.Sequence,
	functionSelectMode []bool,
	colorEditMode []bool,
	selectButtonPressed []bool,
	commandChannels []chan common.Command,
	eventsForLauchpad chan common.ALight) {

	// Turn off the edit colors bar.
	sequences[selectedSequence].Functions[common.Function5_Color].State = false
	cmd := common.Command{
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
}

func setColorEditMode(selectedSequence int,
	sequences []*common.Sequence,
	functionSelectMode []bool,
	colorEditMode []bool,
	selectButtonPressed []bool,
	commandChannels []chan common.Command,
	eventsForLauchpad chan common.ALight) {

	if !functionSelectMode[selectedSequence] && sequences[selectedSequence].Functions[common.Function5_Color].State && colorEditMode[selectedSequence] {
		unSetColorEditMode(selectedSequence, sequences, functionSelectMode, colorEditMode, selectButtonPressed, commandChannels, eventsForLauchpad)
		return
	}

	if functionSelectMode[selectedSequence] && sequences[selectedSequence].Functions[common.Function6_Static].State && colorEditMode[selectedSequence] {
		showEditColorButtons(sequences[selectedSequence], selectedSequence, eventsForLauchpad, sequences[selectedSequence].Master)
		cmd := common.Command{
			SetEditColors: true,
			EditColors:    true,
		}
		common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
		colorEditMode[selectedSequence] = false
		selectButtonPressed[selectedSequence] = false
		functionSelectMode[selectedSequence] = false
		sequences[selectedSequence].Functions[common.Function6_Static].State = false
		common.RevealSequence(selectedSequence, commandChannels)
	}

	if !functionSelectMode[selectedSequence] && sequences[selectedSequence].Functions[common.Function6_Static].State && !colorEditMode[selectedSequence] {
		cmd := common.Command{
			SetEditColors: true,
			EditColors:    false,
		}
		common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
		if !colorEditMode[selectedSequence] {
			showEditColorButtons(sequences[selectedSequence], selectedSequence, eventsForLauchpad, sequences[selectedSequence].Master)
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
func showEditColorButtons(sequence *common.Sequence, selectedSequence int, eventsForLauchpad chan common.ALight, master int) {
	for index, color := range sequence.StaticColors {
		launchpad.LightLamp(selectedSequence, index, color.Color.R, color.Color.G, color.Color.B, master, eventsForLauchpad)
	}
}

// For the given sequence show the available sequence colors on the relevant buttons.
func ShowColorSelectionButtons(mySequenceNumber int, sequence common.Sequence, eventsForLauchpad chan common.ALight) {
	// Check if we need to flash this button.
	for myFixtureNumber, lamp := range sequence.AvailableSequenceColors {
		for index, availableColor := range sequence.AvailableSequenceColors {
			for _, sequenceColor := range sequence.CurrentSequenceColors {
				if availableColor.Color == sequenceColor {
					if myFixtureNumber == index {
						lamp.Flash = true
					}
				}
			}
		}
		if lamp.Flash {
			code := common.GetLaunchPadColorCodeByRGB(lamp.Color)
			launchpad.FlashLight(mySequenceNumber, myFixtureNumber, int(code), 0x0, eventsForLauchpad)
		} else {
			launchpad.LightLamp(mySequenceNumber, myFixtureNumber, lamp.Color.R, lamp.Color.G, lamp.Color.B, sequence.Master, eventsForLauchpad)
		}
	}
}

// For the given sequence hide the available sequence colors..
func HideColorSelectionButtons(mySequenceNumber int, sequence common.Sequence, selectedSequence int, eventsForLauchpad chan common.ALight) {
	for myFixtureNumber := range sequence.AvailableSequenceColors {
		launchpad.LightLamp(mySequenceNumber, myFixtureNumber, 0, 0, 0, sequence.Master, eventsForLauchpad)
	}
}
