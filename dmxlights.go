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

const debug = false

const (
	full = 255
)

// main thread is used to get commands from the lauchpad.
func main() {

	var sequenceSpeed int = 12        // Local copy of sequence speed.
	var size int                      // current RGB sequence size.
	var scannerSize int = 60          // Current scanner size.
	var savePreset bool               // Save a preset flag.
	var selectedShift = 0             // Current fixture shift.
	var blackout bool = false         // Blackout all fixtures.
	var flood bool = false            // Flood all fixtures.
	var functionButtons [][]bool      // Function buttons.
	var functionSelectMode []bool     // Which sequence is in function selection mode.
	var selectButtonPressed []bool    // Which sequence has its Select button pressed.
	var staticLamps [][]bool          // Static color lamps.
	var switchPositions [9][9]int     // Sorage for switch positions.
	var editSequenceColorsMode []bool // This flag is true when the sequence is in sequence colors editing mode.
	var editScannerColorsMode []bool  // This flag is true when the sequence is in select scanner colors editing mode.
	var editGoboSelectionMode []bool  // This flag is true when the sequence is in sequence gobo selection mode.
	var editStaticColorsMode []bool   // This flag is true when the sequence is in static colors editing mode.
	var editPattenMode []bool         // This flag is true when the sequence is in patten editing mode.
	var fadeSpeed = 11                // Default start at 50ms.
	var masterBrightness = 255        // Affects all DMX fixtures and launchpad lamps.
	var lastStaticColorButtonX int    // Which Static Color button did we change last.
	var lastStaticColorButtonY int    // Which Static Color button did we change last.
	var soundGain float32 = 0         // Fine gain -0.09 -> 0.09
	var disabledFixture [][]bool      // Which fixture is disabled on which sequence.

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
	go sequence.PlaySequence(*sequences[0], 0, pad, eventsForLauchpad, pattens, dmxController, fixturesConfig, sequenceChannels, soundTriggers)
	go sequence.PlaySequence(*sequences[1], 1, pad, eventsForLauchpad, pattens, dmxController, fixturesConfig, sequenceChannels, soundTriggers)
	go sequence.PlaySequence(*sequences[2], 2, pad, eventsForLauchpad, pattens, dmxController, fixturesConfig, sequenceChannels, soundTriggers)
	go sequence.PlaySequence(*sequences[3], 3, pad, eventsForLauchpad, pattens, dmxController, fixturesConfig, sequenceChannels, soundTriggers)

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

	// Initialize four select buttons.
	selectButtonPressed = make([]bool, 4)

	// Initialize four function mode states.
	functionSelectMode = make([]bool, 4)

	// Initialize eight fixture states for the four sequences.
	disabledFixture = make([][]bool, 9)
	for i := 0; i < 9; i++ {
		disabledFixture[i] = make([]bool, 9)
	}

	// Initialize a ten length slice of empty slices for static lamps.
	staticLamps = make([][]bool, 9)
	// Initialize those 10 empty function button slices
	for i := 0; i < 9; i++ {
		staticLamps[i] = make([]bool, 9)
	}

	// Remember when we've in editing sequence colors mode.
	editSequenceColorsMode = make([]bool, 4)

	// Remember when we've in setting scanner color mode.
	editScannerColorsMode = make([]bool, 4)

	// Remember when we've in selecting gobo mode.
	editGoboSelectionMode = make([]bool, 4)

	// Remember when we've in editing static colors mode.
	editStaticColorsMode = make([]bool, 4)

	// Remember when we've in editing patten mode.
	editPattenMode = make([]bool, 4)

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

	// Clear the pad.
	allFixturesOff(eventsForLauchpad, dmxController, fixturesConfig)

	// Main loop reading commands from the Novation Launchpad.
	for {

		hit := <-buttonChannel

		// Clear all the lights on the launchpad.
		if hit.X == 0 && hit.Y == -1 {

			if debug {
				fmt.Printf("CLEAR LAUNCHPAD\n")
			}

			// Turn off the flood
			if flood {
				cmd := common.Command{
					Action: common.UpdateFlood,
					Args: []common.Arg{
						{Name: "Flood", Value: false},
					},
				}
				common.SendCommandToAllSequence(selectedSequence, cmd, commandChannels)
				flood = false
			}

			// We want to clear a color selection.
			if sequences[selectedSequence].Functions[common.Function5_Color].State &&
				sequences[selectedSequence].Type != "scanner" {

				// Clear the sequence colors for this sequence.
				cmd := common.Command{
					Action: common.ClearSequenceColor,
				}

				common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

				// Get an upto date copy of the sequence.
				sequences[selectedSequence] = common.RefreshSequence(selectedSequence, commandChannels, updateChannels)

				// Flash the correct color buttons
				ShowColorSelectionButtons(selectedSequence, *sequences[selectedSequence], eventsForLauchpad)

				continue
			}

			launchpad.ClearAll(pad, presetsStore, eventsForLauchpad, commandChannels)
			allFixturesOff(eventsForLauchpad, dmxController, fixturesConfig)
			presets.ClearPresets(eventsForLauchpad, presetsStore)
			presets.InitPresets(eventsForLauchpad, presetsStore)

			// Make sure we stop all sequences.
			cmd := common.Command{
				Action: common.Stop,
			}
			common.SendCommandToAllSequence(selectedSequence, cmd, commandChannels)

			// Swicth off any static colors.
			cmd = common.Command{
				Action: common.UpdateStatic,
				Args: []common.Arg{
					{Name: "Static", Value: false},
				},
			}
			common.SendCommandToAllSequence(selectedSequence, cmd, commandChannels)

			// refesh the switch positions.
			cmd = common.Command{
				Action: common.UpdateSwitchPositions,
			}
			common.SendCommandToAllSequenceOfType(sequences, selectedSequence, cmd, commandChannels, "switch")

			// Clear the sequence colors.
			cmd = common.Command{
				Action: common.ClearSequenceColor,
			}
			common.SendCommandToAllSequence(selectedSequence, cmd, commandChannels)

			// Clear out soundtriggers
			for _, trigger := range soundTriggers {
				trigger.State = false
			}

			// Switch of auto color change and auto patten change.
			for _, sequence := range sequences {
				sequence.AutoColor = false
				sequence.Functions[common.Function2_Auto_Color].State = false
				sequence.AutoPatten = false
				sequence.Functions[common.Function3_Auto_Patten].State = false
			}

			// Disable fixtures.
			for x := 0; x < 4; x++ {
				for y := 0; y < 9; y++ {
					disabledFixture[x][y] = true
				}
			}

			continue
		}

		// F L O O D
		if hit.X == 7 && hit.Y == 3 {

			if debug {
				fmt.Printf("FLOOD\n")
			}

			if !flood {
				cmd := common.Command{
					Action: common.UpdateFlood,
					Args: []common.Arg{
						{Name: "Flood", Value: true},
					},
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
					Action: common.UpdateFlood,
					Args: []common.Arg{
						{Name: "Flood", Value: true},
					},
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
			if debug {
				fmt.Printf("Sound Up %f\n", soundGain)
			}

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
			if debug {
				fmt.Printf("Sound Down%f\n", soundGain)
			}
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

			if debug {
				fmt.Printf("Brightness Down \n")
			}

			masterBrightness = masterBrightness - 10
			if masterBrightness < 0 {
				masterBrightness = 0
			}
			cmd := common.Command{
				Action: common.MasterBrightness,
				Args: []common.Arg{
					{Name: "Master", Value: masterBrightness},
				},
			}
			common.SendCommandToAllSequence(selectedSequence, cmd, commandChannels)
			continue
		}

		// Master brightness up.
		if hit.X == 7 && hit.Y == -1 {

			if debug {
				fmt.Printf("Brightness Up \n")
			}

			masterBrightness = masterBrightness + 10
			if masterBrightness > 255 {
				masterBrightness = 255
			}
			cmd := common.Command{
				Action: common.MasterBrightness,
				Args: []common.Arg{
					{Name: "Master", Value: masterBrightness},
				},
			}
			common.SendCommandToAllSequence(selectedSequence, cmd, commandChannels)
			continue
		}

		// Save mode.
		if hit.X == 8 && hit.Y == 4 {
			if debug {
				fmt.Printf("Save Mode\n")
			}

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
		if hit.X < 8 && (hit.Y > 3 && hit.Y < 7) {

			if debug {
				fmt.Printf("Ask For Config\n")
			}

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
						Action: common.Stop,
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
							Action: common.Normal,
						}
						common.SendCommandToAllSequence(selectedSequence, cmd, commandChannels)
					}
				}
			}
			continue
		}

		// Decrease Shift.
		if hit.X == 2 && hit.Y == 7 {

			if debug {
				fmt.Printf("Decrease Shift\n")
			}

			selectedShift = selectedShift - 1
			if selectedShift < 0 {
				selectedShift = 0
			}
			cmd := common.Command{
				Action: common.UpdateShift,
				Args: []common.Arg{
					{Name: "Shift", Value: selectedShift},
				},
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

			continue
		}

		// Increase Shift.
		if hit.X == 3 && hit.Y == 7 {

			if debug {
				fmt.Printf("Increase Shift \n")
			}

			selectedShift = selectedShift + 1
			if selectedShift > 3 {
				selectedShift = 3
			}
			cmd := common.Command{
				Action: common.UpdateShift,
				Args: []common.Arg{
					{Name: "Shift", Value: selectedShift},
				},
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

			continue
		}

		// Decrease speed of selected sequence.
		if hit.X == 0 && hit.Y == 7 {

			if debug {
				fmt.Printf("Decrease Speed \n")
			}

			// Get an upto date copy of the sequence.
			sequences[selectedSequence] = common.RefreshSequence(selectedSequence, commandChannels, updateChannels)

			if !sequences[selectedSequence].MusicTrigger {
				sequenceSpeed--
				if sequenceSpeed < 0 {
					sequenceSpeed = 1
				}
				cmd := common.Command{
					Action: common.UpdateSpeed,
					Args: []common.Arg{
						{Name: "Speed", Value: sequenceSpeed},
					},
				}
				common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
			}
			continue
		}

		// Increase speed of selected sequence.
		if hit.X == 1 && hit.Y == 7 {

			if debug {
				fmt.Printf("Increase Speed \n")
			}

			// Get an upto date copy of the sequence.
			sequences[selectedSequence] = common.RefreshSequence(selectedSequence, commandChannels, updateChannels)

			if !sequences[selectedSequence].MusicTrigger {
				sequenceSpeed++
				if sequenceSpeed > 21 {
					sequenceSpeed = 21
				}
				cmd := common.Command{
					Action: common.UpdateSpeed,
					Args: []common.Arg{
						{Name: "Speed", Value: sequenceSpeed},
					},
				}
				common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
			}
			continue
		}

		// S E L E C T    S E Q U E N C E.
		// Select sequence 1.
		if hit.X == 8 && hit.Y == 0 {

			if debug {
				fmt.Printf("Select Sequence %d \n", hit.Y)
			}

			selectedSequence = 0
			HandleSelect(sequences, selectedSequence, eventsForLauchpad, selectButtonPressed, functionButtons,
				functionSelectMode, editSequenceColorsMode, editScannerColorsMode, editGoboSelectionMode, editStaticColorsMode, editPattenMode, commandChannels, sequenceChannels)

			cmd := common.Command{
				Action: common.PlayStaticOnce,
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

			editSequenceColorsMode[selectedSequence] = false
			editGoboSelectionMode[selectedSequence] = false

			continue
		}

		// Select sequence 2.
		if hit.X == 8 && hit.Y == 1 {

			if debug {
				fmt.Printf("Select Sequence %d \n", hit.Y)
			}

			selectedSequence = 1
			HandleSelect(sequences, selectedSequence, eventsForLauchpad, selectButtonPressed, functionButtons,
				functionSelectMode, editSequenceColorsMode, editScannerColorsMode, editStaticColorsMode, editGoboSelectionMode, editPattenMode, commandChannels, sequenceChannels)

			cmd := common.Command{
				Action: common.PlayStaticOnce,
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

			editSequenceColorsMode[selectedSequence] = false
			editGoboSelectionMode[selectedSequence] = false

			continue
		}

		// Select sequence 3.
		if hit.X == 8 && hit.Y == 2 {

			if debug {
				fmt.Printf("Select Sequence %d \n", hit.Y)
			}

			selectedSequence = 2
			HandleSelect(sequences, selectedSequence, eventsForLauchpad, selectButtonPressed, functionButtons,
				functionSelectMode, editSequenceColorsMode, editScannerColorsMode, editGoboSelectionMode, editStaticColorsMode, editPattenMode, commandChannels, sequenceChannels)

			cmd := common.Command{
				Action: common.PlayStaticOnce,
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

			editSequenceColorsMode[selectedSequence] = false
			editGoboSelectionMode[selectedSequence] = false

			continue
		}

		// Select sequence 4.
		if hit.X == 8 && hit.Y == 3 {

			if debug {
				fmt.Printf("Select Sequence %d \n", hit.Y)
			}

			selectedSequence = 3
			HandleSelect(sequences, selectedSequence, eventsForLauchpad, selectButtonPressed, functionButtons,
				functionSelectMode, editSequenceColorsMode, editScannerColorsMode, editStaticColorsMode, editGoboSelectionMode, editPattenMode, commandChannels, sequenceChannels)

			cmd := common.Command{
				Action: common.PlayStaticOnce,
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

			editSequenceColorsMode[selectedSequence] = false
			editGoboSelectionMode[selectedSequence] = false

			continue
		}

		// Start sequence.
		if hit.X == 8 && hit.Y == 5 {

			if debug {
				fmt.Printf("Start Sequence %d \n", hit.Y)
			}

			sequences[selectedSequence].MusicTrigger = false
			cmd := common.Command{
				Action: common.Start,
				Args: []common.Arg{
					{Name: "Speed", Value: sequenceSpeed},
				},
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
			common.LightOn(eventsForLauchpad, common.ALight{X: hit.X, Y: hit.Y, Brightness: full, Red: 255, Green: 0, Blue: 0})
			time.Sleep(100 * time.Millisecond)
			common.LightOn(eventsForLauchpad, common.ALight{X: hit.X, Y: hit.Y, Brightness: full, Red: 255, Green: 255, Blue: 255})
			continue
		}

		// Stop sequence.
		if hit.X == 8 && hit.Y == 6 {

			if debug {
				fmt.Printf("Stop Sequence %d \n", hit.Y)
			}

			cmd := common.Command{
				Action: common.Stop,
				Args: []common.Arg{
					{Name: "Speed", Value: sequenceSpeed},
				},
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
			common.LightOn(eventsForLauchpad, common.ALight{X: hit.X, Y: hit.Y, Brightness: full, Red: 255, Green: 0, Blue: 0})
			time.Sleep(100 * time.Millisecond)
			common.LightOn(eventsForLauchpad, common.ALight{X: hit.X, Y: hit.Y, Brightness: full, Red: 255, Green: 255, Blue: 255})
			continue
		}

		// Size decrease.
		if hit.X == 4 && hit.Y == 7 {

			if debug {
				fmt.Printf("Decrease Size\n")
			}

			// Send Update RGB Size.
			size--
			if size < 1 {
				size = 1
			}
			cmd := common.Command{
				Action: common.UpdateSize,
				Args: []common.Arg{
					{Name: "Size", Value: size},
				},
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

			// Send Update Scanner Size.
			scannerSize = scannerSize - 10
			if scannerSize < 10 {
				scannerSize = 10
			}
			cmd = common.Command{
				Action: common.UpdateScannerSize,
				Args: []common.Arg{
					{Name: "ScannerSize", Value: scannerSize},
				},
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

			continue
		}

		// Increase Size.
		if hit.X == 5 && hit.Y == 7 {

			if debug {
				fmt.Printf("Increase Size\n")
			}

			// Send Update RGB Size.
			size++
			if size > 25 {
				size = 25
			}
			cmd := common.Command{
				Action: common.UpdateSize,
				Args: []common.Arg{
					{Name: "Size", Value: size},
				},
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

			// Send Update Scanner Size.
			scannerSize = scannerSize + 10
			if scannerSize > 120 {
				scannerSize = 120
			}
			cmd = common.Command{
				Action: common.UpdateScannerSize,
				Args: []common.Arg{
					{Name: "ScannerSize", Value: scannerSize},
				},
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

			continue
		}

		// Fade time decrease.
		if hit.X == 6 && hit.Y == 7 {

			if debug {
				fmt.Printf("Decrease Fade Time\n")
			}

			fadeSpeed--
			if fadeSpeed < 0 {
				fadeSpeed = 0
			}
			// Send fade update command.
			cmd := common.Command{
				Action: common.DecreaseFade,
				Args: []common.Arg{
					{Name: "FadeSpeed", Value: fadeSpeed},
				},
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
			continue
		}

		// Fade time increase.
		if hit.X == 7 && hit.Y == 7 {

			if debug {
				fmt.Printf("Increase Fade Time\n")
			}

			fadeSpeed++
			if fadeSpeed > 20 {
				fadeSpeed = 20
			}
			// Send fade update command.
			cmd := common.Command{
				Action: common.IncreaseFade,
				Args: []common.Arg{
					{Name: "FadeSpeed", Value: fadeSpeed},
				},
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
			continue
		}

		// F U N C T I O N  K E Y S
		if hit.X >= 0 && hit.X < 8 &&
			functionSelectMode[selectedSequence] &&
			!editPattenMode[selectedSequence] &&
			!editStaticColorsMode[selectedSequence] &&
			!editGoboSelectionMode[selectedSequence] &&
			!sequences[selectedSequence].Functions[common.Function5_Color].State {

			if debug {
				fmt.Printf("Function Key X:%d Y:%d\n", hit.X, hit.Y)
			}

			// Get an upto date copy of the sequence.
			sequences[selectedSequence] = common.RefreshSequence(selectedSequence, commandChannels, updateChannels)

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
				Action: common.UpdateFunctions,
				Args: []common.Arg{
					{Name: "Functions", Value: sequences[selectedSequence].Functions},
				},
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

			// Light the correct function key.
			common.ShowFunctionButtons(*sequences[selectedSequence], selectedSequence, eventsForLauchpad)

			// Now some functions mean that we go into another menu ( set of buttons )
			// This is true for :-
			// Function 1 - setting the patten.
			// Function 5 - setting the sequence colors or selecting scanner color.
			// Function 6 - setting the static colors or selecting scanner gobo.

			// Map Function 1 to patten mode.
			editPattenMode[selectedSequence] = sequences[selectedSequence].Functions[common.Function1_Patten].State

			// Go straight into patten select mode, don't wait for a another select press.
			if editPattenMode[selectedSequence] {
				time.Sleep(500 * time.Millisecond) // But give the launchpad time to light the function key purple.
				common.HideFunctionButtons(selectedSequence, eventsForLauchpad)
				ShowPattenSelectionButtons(selectedSequence, *sequences[selectedSequence], eventsForLauchpad)
			}

			// Map Function 5 to color edit.
			editSequenceColorsMode[selectedSequence] = sequences[selectedSequence].Functions[common.Function5_Color].State

			// Go straight into color edit mode, don't wait for a another select press.
			if editSequenceColorsMode[selectedSequence] {
				time.Sleep(500 * time.Millisecond) // But give the launchpad time to light the function key purple.
				common.HideFunctionButtons(selectedSequence, eventsForLauchpad)
				ShowColorSelectionButtons(selectedSequence, *sequences[selectedSequence], eventsForLauchpad)
			}

			// Map Function 6 to select gobo mode.
			if sequences[selectedSequence].Type == "scanner" {
				editGoboSelectionMode[selectedSequence] = sequences[selectedSequence].Functions[common.Function7_Gobo].State
			}

			// Go straight to gobo selection mode.
			if editGoboSelectionMode[selectedSequence] {
				time.Sleep(500 * time.Millisecond) // But give the launchpad time to light the function key purple.
				common.HideFunctionButtons(selectedSequence, eventsForLauchpad)
				ShowGoboSelectionButtons(selectedSequence, *sequences[selectedSequence], eventsForLauchpad)
			}

			// Map Function 6 to static color edit.
			if sequences[selectedSequence].Type != "scanner" {
				editStaticColorsMode[selectedSequence] = sequences[selectedSequence].Functions[common.Function6_Static].State
			}

			// Go straight into static color edit mode, don't wait for a another select press.
			if editStaticColorsMode[selectedSequence] {
				//time.Sleep(2000 * time.Millisecond) // But give the launchpad time to light the function key purple.
				//common.HideFunctionButtons(selectedSequence, eventsForLauchpad)
				functionSelectMode[selectedSequence] = false
				// The sequence will automatically display the colors now!
			}

			continue
		}

		// S W I T C H   B U T T O N's Toggle State of switches for this sequence.
		if hit.X >= 0 && hit.X < 8 && !functionSelectMode[selectedSequence] &&
			hit.Y >= 0 &&
			hit.Y < 4 &&
			sequences[hit.Y].Type == "switch" {

			if debug {
				fmt.Printf("Switch Key X:%d Y:%d\n", hit.X, hit.Y)
			}

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
					Action: common.UpdateSwitch,
					Args: []common.Arg{
						{Name: "SwitchNumber", Value: hit.X},
						{Name: "SwitchPosition", Value: switchPositions[hit.Y][hit.X]},
					},
				}
				// Send a message to the switch sequence.
				common.SendCommandToAllSequenceOfType(sequences, hit.Y, cmd, commandChannels, "switch")
			}
		}

		// D I S A B L E   F I X T U R E  - Used to toggle the scanner on or off.
		if hit.X >= 0 && hit.X < 8 && !functionSelectMode[selectedSequence] &&
			hit.Y >= 0 &&
			hit.Y < 4 &&
			!sequences[selectedSequence].Functions[common.Function1_Patten].State &&
			!sequences[selectedSequence].Functions[common.Function6_Static].State &&
			!sequences[selectedSequence].Functions[common.Function5_Color].State &&
			sequences[hit.Y].Type == "scanner" {

			if debug {
				fmt.Printf("Disable Fixture X:%d Y:%d\n", hit.X, hit.Y)
			}

			if !disabledFixture[hit.X][hit.Y] && hit.X < sequences[hit.Y].NumberScanners {

				if debug {
					fmt.Printf("Toggle Scanner Number %d State on Sequence %d to true [Scanners:%d]\n", hit.X, hit.Y, sequences[hit.Y].NumberScanners)
				}

				disabledFixture[hit.X][hit.Y] = true

				// Tell the sequence to turn off this scanner.
				cmd := common.Command{
					Action: common.ToggleFixtureState,
					Args: []common.Arg{
						{Name: "SequenceNumber", Value: hit.Y},
						{Name: "FixtureNumber", Value: hit.X},
						{Name: "FixtureState", Value: true},
					},
				}
				common.SendCommandToSequence(hit.Y, cmd, commandChannels)

				// Turn off the lamp.
				common.LightOff(eventsForLauchpad, hit.X, hit.Y)

				continue

			}

			if disabledFixture[hit.X][hit.Y] && hit.X < sequences[hit.Y].NumberScanners {
				if debug {
					fmt.Printf("Toggle Scanner Number %d State on Sequence %d to false\n", hit.X, hit.Y)
				}

				disabledFixture[hit.X][hit.Y] = false

				// Tell the sequence to turn on this scanner.
				cmd := common.Command{
					Action: common.ToggleFixtureState,
					Args: []common.Arg{
						{Name: "SequenceNumber", Value: hit.Y},
						{Name: "FixtureNumber", Value: hit.X},
						{Name: "FixtureState", Value: false},
					},
				}
				common.SendCommandToSequence(hit.Y, cmd, commandChannels)

				// Turn the lamp on.
				common.LightOn(eventsForLauchpad, common.ALight{X: hit.X, Y: hit.Y, Brightness: full, Red: 255, Green: 255, Blue: 255})

				continue
			}

		}

		// F L A S H   B U T T O N S - Briefly light (flash) the fixtures based on current patten.
		if hit.X >= 0 &&
			hit.X < 8 &&
			hit.Y >= 0 &&
			hit.Y < 4 &&
			!sequences[hit.Y].Functions[common.Function1_Patten].State &&
			!sequences[hit.Y].Functions[common.Function6_Static].State &&
			!sequences[hit.Y].Functions[common.Function5_Color].State &&
			sequences[hit.Y].Type != "switch" && // As long as we're not a switch sequence.
			sequences[hit.Y].Type != "scanner" && // As long as we're not a scanner sequence.
			!functionSelectMode[hit.Y] { // As long as we're not a scanner sequence.

			if debug {
				fmt.Printf("Flash Button X:%d Y:%d\n", hit.X, hit.Y)
			}

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
			fixture.MapFixtures(hit.Y, dmxController, hit.X, red, green, blue, pan, tilt, shutter, gobo, 0, fixturesConfig, blackout, masterBrightness, masterBrightness)
			time.Sleep(200 * time.Millisecond)
			common.LightOff(eventsForLauchpad, hit.X, hit.Y)
			fixture.MapFixtures(hit.Y, dmxController, hit.X, 0, 0, 0, pan, tilt, shutter, gobo, 0, fixturesConfig, blackout, masterBrightness, masterBrightness)
			continue
		}

		// C H O O S E   S T A T I C    C O L O R
		// Red
		if hit.X == 1 && hit.Y == -1 {

			if debug {
				fmt.Printf("Choose Static Red X:%d Y:%d\n", hit.X, hit.Y)
			}

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

			if debug {
				fmt.Printf("Choose Static Green X:%d Y:%d\n", hit.X, hit.Y)
			}

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

			if debug {
				fmt.Printf("Choose Static Blue X:%d Y:%d\n", hit.X, hit.Y)
			}

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

		// S E T  R G B  S E Q U E N C E   C O L O R
		if hit.X >= 0 && hit.X < 8 && hit.Y != -1 &&
			selectedSequence == hit.Y && // Make sure the buttons pressed are for this sequence.
			sequences[selectedSequence].Type != "scanner" &&
			sequences[selectedSequence].Functions[common.Function5_Color].State {

			if debug {
				fmt.Printf("Set Sequence Color X:%d Y:%d\n", hit.X, hit.Y)
			}

			// If we're a scanner we can only select one color at a time.
			if sequences[selectedSequence].Type == "scanner" {
				// Clear the sequence colors for this sequence.
				cmd := common.Command{
					Action: common.ClearSequenceColor,
				}
				common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
			}

			// Add the selected color to the sequence.
			cmd := common.Command{
				Action: common.UpdateSequenceColor,
				Args: []common.Arg{
					{Name: "SelectedColor", Value: hit.X},
				},
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

			editSequenceColorsMode[selectedSequence] = true

			// Get an upto date copy of the sequence.
			sequences[selectedSequence] = common.RefreshSequence(selectedSequence, commandChannels, updateChannels)

			// Set the colors.
			sequences[selectedSequence].CurrentSequenceColors = sequences[selectedSequence].SequenceColors

			// We call ShowColorSelectionButtons here so the selections will flash as you press them.
			ShowColorSelectionButtons(selectedSequence, *sequences[selectedSequence], eventsForLauchpad)

			continue
		}

		// S E T  S C A N N E R   C O L O R
		if hit.X >= 0 && hit.X < 8 && hit.Y != -1 &&
			selectedSequence == hit.Y && // Make sure the buttons pressed are for this sequence.
			sequences[selectedSequence].Type == "scanner" &&
			sequences[selectedSequence].Functions[common.Function5_Color].State {

			if debug {
				fmt.Printf("Set Scanner Color X:%d Y:%d\n", hit.X, hit.Y)
			}

			// Set the scanner color for this sequence.
			cmd := common.Command{
				Action: common.UpdateScannerColor,
				Args: []common.Arg{
					{Name: "SelectedColor", Value: hit.X},
				},
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

			editScannerColorsMode[selectedSequence] = true

			// Get an upto date copy of the sequence.
			sequences[selectedSequence] = common.RefreshSequence(selectedSequence, commandChannels, updateChannels)

			// Set the colors.
			sequences[selectedSequence].CurrentSequenceColors = sequences[selectedSequence].SequenceColors

			// We call ShowColorSelectionButtons here so the selections will flash as you press them.
			ShowColorSelectionButtons(selectedSequence, *sequences[selectedSequence], eventsForLauchpad)

			continue
		}

		// S E T   S C A N N E R   G O B O
		if hit.X >= 0 && hit.X < 8 && hit.Y != -1 &&
			sequences[selectedSequence].Functions[common.Function7_Gobo].State &&
			sequences[selectedSequence].Type == "scanner" {

			if debug {
				fmt.Printf("Set Gobo X:%d Y:%d\n", hit.X, hit.Y)
			}

			// Add the selected color to the sequence.
			cmd := common.Command{
				Action: common.UpdateGobo,
				Args: []common.Arg{
					{Name: "SelectedGobo", Value: hit.X},
				},
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

			editGoboSelectionMode[selectedSequence] = true

			// Get an upto date copy of the sequence.
			sequences[selectedSequence] = common.RefreshSequence(selectedSequence, commandChannels, updateChannels)

			// Set the colors.
			sequences[selectedSequence].CurrentSequenceColors = sequences[selectedSequence].SequenceColors

			// We call ShowGoboSelectionButtons here so the selections will flash as you press them.
			ShowGoboSelectionButtons(selectedSequence, *sequences[selectedSequence], eventsForLauchpad)

			continue
		}

		// S E T  S T A T I C   C O L O R
		if hit.X >= 0 && hit.X < 8 &&
			hit.Y != -1 &&
			selectedSequence == hit.Y && // Make sure the buttons pressed are for this sequence.
			sequences[selectedSequence].Type != "scanner" && // Not a scanner sequence.
			!functionSelectMode[selectedSequence] && // Not in function Mode
			editStaticColorsMode[selectedSequence] { // Static Function On

			if debug {
				fmt.Printf("Set Static Color X:%d Y:%d\n", hit.X, hit.Y)
			}

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
				Action: common.UpdateStaticColor,
				Args: []common.Arg{
					{Name: "Static", Value: true},
					{Name: "StaticLamp", Value: hit.X},
					{Name: "StaticLampFlash", Value: false},
					{Name: "SelectedColor", Value: sequences[selectedSequence].StaticColors[hit.X].SelectedColor},
					{Name: "StaticColor", Value: sequences[selectedSequence].StaticColors[hit.X].Color},
				},
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)
			lastStaticColorButtonX = hit.X
			lastStaticColorButtonY = hit.Y

			continue
		}

		// S E T   P A T T E N
		if hit.X >= 0 && hit.X < 8 && hit.Y != -1 &&
			editPattenMode[selectedSequence] {

			if debug {
				fmt.Printf("Set Patten X:%d Y:%d\n", hit.X, hit.Y)
			}

			// Tell the sequence to change the patten
			cmd := common.Command{
				Action: common.SelectPatten,
				Args: []common.Arg{
					{Name: "SelectedPatten", Value: hit.X},
				},
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

			//editPattenMode[selectedSequence] = false
			functionSelectMode[selectedSequence] = false

			// Get an upto date copy of the sequence.
			sequences[selectedSequence] = common.RefreshSequence(selectedSequence, commandChannels, updateChannels)

			// We call ShowPattenSelectionButtons here so the selections will flash as you press them.
			ShowPattenSelectionButtons(selectedSequence, *sequences[selectedSequence], eventsForLauchpad)

			continue
		}

		// B L A C K O U T   B U T T O N.
		if hit.X == 8 && hit.Y == 7 {

			if debug {
				fmt.Printf("BLACKOUT\n")
			}

			if !blackout {
				blackout = true
				cmd := common.Command{
					Action: common.Blackout,
				}
				common.SendCommandToAllSequence(selectedSequence, cmd, commandChannels)
				common.LightOn(eventsForLauchpad, common.ALight{X: hit.X, Y: hit.Y, Brightness: full, Red: 0, Green: 0, Blue: 255})
			} else {
				blackout = false
				cmd := common.Command{
					Action: common.Normal,
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
		Action: common.UpdateStaticColor,
		Args: []common.Arg{
			{Name: "Static", Value: true},
			{Name: "StaticLamp", Value: staticColorButtons.X},
			{Name: "SelectedColor", Value: staticColorButtons.SelectedColor},
			{Name: "StaticColor", Value: common.Color{R: staticColorButtons.Color.R, G: staticColorButtons.Color.G, B: staticColorButtons.Color.B}},
		},
	}
	common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

}

// HandleSelect - Runs when you press a select button to select a sequence.
func HandleSelect(sequences []*common.Sequence,
	selectedSequence int,
	eventsForLauchpad chan common.ALight,
	selectButtonPressed []bool,
	functionButtons [][]bool,
	functionSelectMode []bool,
	editSequenceColorsMode []bool,
	editScannerColorsMode []bool,
	editGoboSelectionMode []bool,
	editStaticColorsMode []bool,
	editPattenMode []bool,
	commandChannels []chan common.Command,
	channels common.Channels) {

	if debug {
		fmt.Printf("HANDLE: selectButtons[%d] = %t \n", selectedSequence, selectButtonPressed[selectedSequence])
		fmt.Printf("HANDLE: editSequenceColorsMode[%d] = %t \n", selectedSequence, editSequenceColorsMode[selectedSequence])
		fmt.Printf("HANDLE: editStaticColorsMode[%d] = %t \n", selectedSequence, editStaticColorsMode[selectedSequence])
		fmt.Printf("HANDLE: editGoboSelectionMode[%d] = %t \n", selectedSequence, editGoboSelectionMode[selectedSequence])
		fmt.Printf("HANDLE: functionSelectMode[%d] = %t \n", selectedSequence, functionSelectMode[selectedSequence])
		fmt.Printf("HANDLE: editPattenMode[%d] = %t \n", selectedSequence, editPattenMode[selectedSequence])
		fmt.Printf("HANDLE: Func Static[%d] = %t\n", selectedSequence, sequences[selectedSequence].Functions[common.Function6_Static].State)
	}

	// Light the sequence selector button.
	common.SequenceSelect(eventsForLauchpad, selectedSequence)

	// First time into function mode we head back to normal mode.
	if functionSelectMode[selectedSequence] && !selectButtonPressed[selectedSequence] &&
		!editSequenceColorsMode[selectedSequence] && !editStaticColorsMode[selectedSequence] {
		if debug {
			fmt.Printf("Handle 1 Function Bar off\n")
		}
		// Turn off function mode. Remove the function pads.
		common.HideFunctionButtons(selectedSequence, eventsForLauchpad)

		functionSelectMode[selectedSequence] = false

		if sequences[selectedSequence].Functions[common.Function1_Patten].State {
			if debug {
				fmt.Printf("Show Patten Selection Buttons\n")
			}
			editPattenMode[selectedSequence] = true
			common.HideSequence(selectedSequence, commandChannels)
			ShowPattenSelectionButtons(selectedSequence, *sequences[selectedSequence], eventsForLauchpad)
			return
		}

		if sequences[selectedSequence].Functions[common.Function5_Color].State {
			if debug {
				fmt.Printf("Show Sequence Color Selection Buttons\n")
			}
			ShowColorSelectionButtons(selectedSequence, *sequences[selectedSequence], eventsForLauchpad)
			return
		}

		if sequences[selectedSequence].Functions[common.Function6_Static].State &&
			sequences[selectedSequence].Type != "scanner" {
			if debug {
				fmt.Printf("Show Static Color Selection Buttons\n")
			}
			common.SetMode(selectedSequence, commandChannels, "Static")
			return
		}

		// Allow us to exit the patten select mode without setting a patten.
		if editPattenMode[selectedSequence] {
			editPattenMode[selectedSequence] = false
		}

		// Else reveal the sequence on the launchpad keys
		if debug {
			fmt.Printf("Reveal Sequence\n")
		}
		common.RevealSequence(selectedSequence, commandChannels)
		// Turn off the function mode flag.
		functionSelectMode[selectedSequence] = false
		// Now forget we pressed twice and start again.
		selectButtonPressed[selectedSequence] = true

		return
	}

	// This the first time we have pressed the select button.
	if !selectButtonPressed[selectedSequence] &&
		!editStaticColorsMode[selectedSequence] {
		if debug {
			fmt.Printf("Handle 2\n")
		}
		// assume everything else is off.
		selectButtonPressed[0] = false
		selectButtonPressed[1] = false
		selectButtonPressed[2] = false
		selectButtonPressed[3] = false
		// But remember we have pressed this select button once.
		functionSelectMode[selectedSequence] = false
		selectButtonPressed[selectedSequence] = true

		if sequences[selectedSequence].Functions[common.Function1_Patten].State {
			// Reset the patten function key.
			sequences[selectedSequence].Functions[common.Function1_Patten].State = false

			// Clear the patten function keys
			ClearPattenSelectionButtons(selectedSequence, *sequences[selectedSequence], eventsForLauchpad)

			// And reveal the sequence.
			common.RevealSequence(selectedSequence, commandChannels)

			// Editing patten is over for this sequence.
			editPattenMode[selectedSequence] = false
		}

		if !functionSelectMode[selectedSequence] && sequences[selectedSequence].Functions[common.Function5_Color].State && editSequenceColorsMode[selectedSequence] {
			unSetEditSequenceColorsMode(selectedSequence, sequences, functionSelectMode, editSequenceColorsMode, selectButtonPressed, commandChannels, eventsForLauchpad)
		}
		return
	}

	// Are we in function mode ?
	if functionSelectMode[selectedSequence] {
		if debug {
			fmt.Printf("Handle 3\n")
		}
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
		sequences[selectedSequence].Type != "switch" || // Don't alow functions in switch mode.
		!functionSelectMode[selectedSequence] && editStaticColorsMode[selectedSequence] { // The case when we leave static colors edit mode.

		if debug {
			fmt.Printf("Handle 4 - Function Bar On!\n")
		}

		// Unset the edit static color mode.
		editStaticColorsMode[selectedSequence] = false

		// Set function mode.
		functionSelectMode[selectedSequence] = true

		// And hide the sequence so we can only see the function buttons.
		common.HideSequence(selectedSequence, commandChannels)

		// Turn off any static sequence so we can see the functions.
		common.SetMode(selectedSequence, commandChannels, "Sequence")

		// Create the function buttons.
		common.MakeFunctionButtons(*sequences[selectedSequence], selectedSequence, eventsForLauchpad, functionButtons, channels)

		// Now forget we pressed twice and start again.
		selectButtonPressed[selectedSequence] = false

		return
	}
}

func unSetEditSequenceColorsMode(selectedSequence int,
	sequences []*common.Sequence,
	functionSelectMode []bool,
	editSequenceColorsMode []bool,
	selectButtonPressed []bool,
	commandChannels []chan common.Command,
	eventsForLauchpad chan common.ALight) {

	// Turn off the edit colors bar.
	sequences[selectedSequence].Functions[common.Function5_Color].State = false
	cmd := common.Command{
		Action: common.UpdateFunctions,
		Args: []common.Arg{
			{Name: "Functions", Value: sequences[selectedSequence].Functions},
		},
	}
	common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

	// Restart the sequence.
	cmd = common.Command{
		Action: common.Start,
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

func allFixturesOff(eventsForLauchpad chan common.ALight, dmxController *ft232.DMXController, fixturesConfig *fixture.Fixtures) {
	for x := 0; x < 8; x++ {
		for y := 0; y < 4; y++ {
			common.LightOff(eventsForLauchpad, x, y)
			fixture.MapFixtures(y, dmxController, x, 0, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, true, 0, 0)
		}
	}
}

// For the given sequence show the available sequence colors on the relevant buttons.
func ShowColorSelectionButtons(mySequenceNumber int, sequence common.Sequence, eventsForLauchpad chan common.ALight) {

	if debug {
		fmt.Printf("Show Color Selection Buttons\n")
	}
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

// For the given sequence show the available gobo selection colors on the relevant buttons.
func ShowGoboSelectionButtons(mySequenceNumber int, sequence common.Sequence, eventsForLauchpad chan common.ALight) {

	if debug {
		fmt.Printf("Show Gobo Selection Buttons\n")
	}
	// Check if we need to flash this button.
	for myFixtureNumber, lamp := range sequence.AvailableGoboSelectionColors {
		if debug {
			fmt.Printf("myFixtureNumber %d   currenr gobo %d\n", myFixtureNumber, sequence.SelectedGobo)
		}
		if myFixtureNumber == sequence.SelectedGobo {
			lamp.Flash = true
		}
		if lamp.Flash {
			code := common.GetLaunchPadColorCodeByRGB(lamp.Color)
			launchpad.FlashLight(mySequenceNumber, myFixtureNumber, int(code), 0x0, eventsForLauchpad)
		} else {
			launchpad.LightLamp(mySequenceNumber, myFixtureNumber, lamp.Color.R, lamp.Color.G, lamp.Color.B, sequence.Master, eventsForLauchpad)
		}
	}
}

// For the given sequence clear the available pattens on the relevant buttons.
func ClearPattenSelectionButtons(mySequenceNumber int, sequence common.Sequence, eventsForLauchpad chan common.ALight) {
	// Check if we need to flash this button.
	for myFixtureNumber := 0; myFixtureNumber < 4; myFixtureNumber++ {
		launchpad.LightLamp(mySequenceNumber, myFixtureNumber, 0, 0, 0, sequence.Master, eventsForLauchpad)
	}
}

// For the given sequence show the available pattens on the relevant buttons.
func ShowPattenSelectionButtons(mySequenceNumber int, sequence common.Sequence, eventsForLauchpad chan common.ALight) {
	// Check if we need to flash this button.
	for myFixtureNumber := 0; myFixtureNumber < 4; myFixtureNumber++ {
		if myFixtureNumber == sequence.SelectedPatten {
			code := common.GetLaunchPadColorCodeByRGB(common.Color{R: 255, G: 255, B: 255})
			launchpad.FlashLight(mySequenceNumber, myFixtureNumber, int(code), 0x0, eventsForLauchpad)
		} else {
			launchpad.LightLamp(mySequenceNumber, myFixtureNumber, 255, 255, 255, sequence.Master, eventsForLauchpad)
		}
	}
}

// For the given sequence hide the available sequence colors..
func HideColorSelectionButtons(mySequenceNumber int, sequence common.Sequence, selectedSequence int, eventsForLauchpad chan common.ALight) {
	for myFixtureNumber := range sequence.AvailableSequenceColors {
		launchpad.LightLamp(mySequenceNumber, myFixtureNumber, 0, 0, 0, sequence.Master, eventsForLauchpad)
	}
}
