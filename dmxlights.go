// Copyright (C) 2022, 2023 dhowlett99.
// This is the dmxlights main program and calls all others.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"image/color"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/dhowlett99/dmxlights/pkg/buttons"
	"github.com/dhowlett99/dmxlights/pkg/commands"
	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/dmx"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
	"github.com/dhowlett99/dmxlights/pkg/gui"
	"github.com/dhowlett99/dmxlights/pkg/launchpad"
	"github.com/dhowlett99/dmxlights/pkg/pattern"
	"github.com/dhowlett99/dmxlights/pkg/presets"
	"github.com/dhowlett99/dmxlights/pkg/sequence"
	"github.com/dhowlett99/dmxlights/pkg/sound"
	"github.com/oliread/usbdmx/ft232"
)

const debug = false
const NumberOfSequences = 5
const NumberOfFixtures = 8
const NumberOfSwitches = 8

const DEFAULT_PROJECT = "Default.yaml"

func main() {

	fmt.Println("DMX Lighting")

	os.Setenv("FYNE_THEME", "light")

	// Start the GUI.
	fmt.Println("Starting GUI")
	panel := gui.NewPanel() // Panel represents the buttons in the GUI.
	myApp := app.New()

	myWindow := myApp.NewWindow("DMX Lights")
	myWindow.Resize(fyne.NewSize(400, 50))

	if desk, ok := myApp.(desktop.App); ok {
		menu := fyne.NewMenu("MyApp",
			fyne.NewMenuItem("Show", func() {
				myWindow.Show()
			}))
		desk.SetSystemTrayMenu(menu)
	}

	// Setup the current state.
	this := buttons.CurrentState{}
	this.MyWindow = myWindow                                              // Pointer to main window.
	this.Blackout = false                                                 // Blackout starts in off.
	this.Flood = false                                                    // Flood starts in off.
	this.Running = make(map[int]bool, NumberOfSequences)                  // Initialise storage for four sequences.
	this.Strobe = make(map[int]bool, NumberOfSequences)                   // Initialise storage for four sequences.
	this.MasterBrightness = 255                                           // Affects all DMX fixtures and launchpad lamps.
	this.SoundGain = 0                                                    // Fine gain -0.09 -> 0.09
	this.OffsetPan = common.SCANNER_MID_POINT                             // Start pan from the center
	this.OffsetTilt = common.SCANNER_MID_POINT                            // Start tilt from the center.
	this.RGBPatterns = pattern.MakePatterns()                             // Build the default set of Patterns.
	this.SelectButtonPressed = make([]bool, NumberOfSequences)            // Initialise four select buttons.
	this.SelectedMode = make([]int, NumberOfSequences)                    // Initialise four mode variables.
	this.LastMode = make([]int, NumberOfSequences)                        // Initialise four mode variables.
	this.ShowRGBColorPicker = false                                       // Remember when we are in editing sequence colors mode.
	this.EditScannerColorsMode = false                                    // Remember when we are in setting scanner color mode.
	this.EditGoboSelectionMode = false                                    // Remember when we are in selecting gobo mode.
	this.Static = make([]bool, NumberOfSequences)                         // Remember when this sequence is in static mode.
	this.StaticFlashing = make([]bool, NumberOfSequences)                 // Remember when we are in static buttons are flashing.
	this.SequenceType = make([]string, NumberOfSequences)                 // Remember sequence type.
	this.EditPatternMode = false                                          // Remember when we are in editing pattern mode.
	this.StaticButtons = makeStaticButtonsStorage()                       // Make storgage for color editing button results.
	this.PresetsStore = presets.LoadPresets()                             // Load the presets from their json files.
	this.Speed = make(map[int]int, NumberOfSequences+NumberOfSwitches)    // Initialise storage for four sequences and eight switches.
	this.SwitchOverrides = make([][]common.Override, NumberOfSwitches)    // Initialise local override storage for eight switches. Indexed by switch number.
	this.RGBSize = make(map[int]int, NumberOfSequences+NumberOfSwitches)  // Initialise storage for four sequences and eight switches.
	this.ScannerSize = make(map[int]int, NumberOfSequences)               // Initialise storage for four sequences.
	this.RGBShift = make(map[int]int, NumberOfSequences+NumberOfSwitches) // Initialise storage for four sequences and eight switches..
	this.ScannerShift = make(map[int]int, NumberOfSequences)              // Initialise storage for four sequences.
	this.RGBFade = make(map[int]int, NumberOfSequences)                   // Initialise storage for four sequences.
	this.ScannerFade = make(map[int]int, NumberOfSequences)               // Initialise storage for four sequences.
	this.StrobeSpeed = make(map[int]int, NumberOfSequences)               // Initialise storage for four sequences.
	this.ClearPressed = make(map[int]bool, NumberOfSequences)             // Initialise storage for four sequences.
	this.ScannerChaser = make(map[int]bool, NumberOfSequences)            // Initialise storage for four sequences.
	this.ScannerCoordinates = make(map[int]int, NumberOfSequences)        // Number of coordinates for scanner patterns is selected from 4 choices. 0=12, 1=16,2=24,3=32,4=64
	this.LaunchPadConnected = true                                        // Assume launchpad is present, until tested.
	this.DmxInterfacePresent = true                                       // Assume DMX interface card is present, until tested.
	this.LaunchpadName = "Novation Launchpad Mk3 Mini"                    // Name of launchpad.
	this.Functions = make(map[int][]common.Function)                      // Array holding functions for each sequence.
	this.SavedSequenceColors = make(map[int][]color.RGBA)                 // Array holding saved sequence colors for each sequence. Used by the color picker.
	this.LastSelectedSwitch = common.NOT_SELECTED                         // Set the last selected switch to not selected.

	// Now add channels to communicate with mini-sequencers on switch channels.
	this.SwitchChannels = []common.SwitchChannel{}
	for switchChannel := 0; switchChannel < 10; switchChannel++ {
		newSwitch := common.SwitchChannel{}
		newSwitch.Stop = make(chan bool)
		newSwitch.KeepRotateAlive = make(chan bool)
		newSwitch.StopRotate = make(chan bool)
		newSwitch.StopFadeUp = make(chan bool)
		newSwitch.StopFadeDown = make(chan bool)
		newSwitch.CommandChannel = make(chan common.Command)
		this.SwitchChannels = append(this.SwitchChannels, newSwitch)
	}
	// Initialize eight fixture states for the four sequences.
	this.FixtureState = make([][]common.FixtureState, NumberOfSequences)
	// Populate each sequence with fixtures.
	for sequenceNumber := 0; sequenceNumber < NumberOfSequences; sequenceNumber++ {
		this.FixtureState[sequenceNumber] = make([]common.FixtureState, NumberOfFixtures)
		for fixtureNumber := 0; fixtureNumber < NumberOfFixtures; fixtureNumber++ {
			newFixture := common.FixtureState{}
			newFixture.Enabled = true
			newFixture.RGBInverted = false
			newFixture.ScannerPatternReversed = false
			this.FixtureState[sequenceNumber][fixtureNumber] = newFixture
		}
	}

	// Setup DMX interface.
	fmt.Println("Setup DMX Interface")
	dmxController, dmxInterfaceConfig, err := dmx.NewDmXController()
	if err != nil {
		fmt.Printf("dmx interface: %v\n", err)
		this.DmxInterfacePresent = false
	}

	// Save the presets on exit.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Saving Presets")
		presets.SavePresets(this.PresetsStore)
		os.Exit(1)
	}()

	// Setup a connection to the Novation Launchpad.
	// Tested with a Novation Launchpad mini mk3.
	fmt.Println("Setup Novation Launchpad")
	this.Pad, err = launchpad.NewLaunchPad()
	if err != nil {
		fmt.Printf("launchpad: %v\n", err)
		this.LaunchPadConnected = false
		this.LaunchpadName = "Not Found"
	}

	// If launchpad found, defer the close.
	if this.LaunchPadConnected {
		defer this.Pad.Close()
	}

	// Report on connected devices.
	panel.PopupNotFoundMessage(myWindow,
		gui.Device{
			Name:   "DMX Interface",
			Status: this.DmxInterfacePresent},
		gui.Device{
			Name:   "LaunchPad",
			Status: this.LaunchPadConnected})

	// Create a channel to send events to the launchpad.
	eventsForLaunchpad := make(chan common.ALight)

	// We need to be in programmers mode to use the launchpad.
	if this.LaunchPadConnected {
		this.Pad.Program()
	}

	// Create a channel to send events to the GUI.
	guiButtons := make(chan common.ALight)

	// Make space for info on which GUI button is flashing.
	GuiFlashButtons := make([][]common.ALight, 10)
	for y := 0; y < 10; y++ {
		GuiFlashButtons[y] = make([]common.ALight, 10)
		for x := 0; x < 10; x++ {
			GuiFlashButtons[y][x].FlashStopChannel = make(chan bool)
		}
	}

	// Read sequences config file
	fmt.Println("Load Sequences Config File")
	sequencesConfig, err := sequence.LoadSequences()
	if err != nil {
		fmt.Printf("dmxlights: error failed to load sequences config: %s\n", err.Error())
		os.Exit(1)
	}

	// Get a list of all the fixtures in the groups.
	fixturesConfig, err := fixture.LoadFixtures(DEFAULT_PROJECT)
	if err != nil {
		fmt.Printf("dmxlights: error failed to load fixtures: %s\n", err.Error())
		os.Exit(1)
	}

	// Load groups.
	groupConfig, err := fixture.LoadFixtureGroups("groups.yaml")
	if err != nil {
		fmt.Printf("dmxlights: error failed to load groups: %s\n", err.Error())
		os.Exit(1)
	}

	// Update the fixture list with the sequence type.
	for _, sequence := range sequencesConfig.Sequences {
		for fixtureNumber, fixture := range fixturesConfig.Fixtures {
			if sequence.Group == fixture.Group {
				fixturesConfig.Fixtures[fixtureNumber].Type = sequence.Type
			}
		}
	}

	for fixtureNumber, fixture := range fixturesConfig.Fixtures {
		// Automatically set the number of sub fixtures inside a fixture.
		var numberSubFixtures int
		for _, channel := range fixture.Channels {
			if strings.Contains(channel.Name, "Red") {
				numberSubFixtures++
			}
		}
		if numberSubFixtures > 1 {
			if debug {
				fmt.Printf("\t fixture %s numberSubFixtures %d\n", fixture.Name, numberSubFixtures)
			}
			fixturesConfig.Fixtures[fixtureNumber].MultiFixtureDevice = true
			fixturesConfig.Fixtures[fixtureNumber].NumberSubFixtures = numberSubFixtures
		}
	}

	// Now that the fixtures config is setup, make a copy.
	startConfig := &fixture.Fixtures{}
	startConfig.Fixtures = []fixture.Fixture{}
	startConfig.Fixtures = append(startConfig.Fixtures, fixturesConfig.Fixtures...)

	myWindow.SetTitle("DMX Lights:" + DEFAULT_PROJECT)

	// If you try to quit without saving your changed project. Uses startConfig as a ref to determine changes.
	myWindow.SetCloseIntercept(func() {
		theSame, message := fixture.CheckFixturesAreTheSame(fixturesConfig, startConfig)
		if !theSame {
			model := gui.AreYouSureDialog(myWindow, message)
			model.Show()
		} else {
			os.Exit(0)
		}
	})

	// Create the sequences from config file.
	// Add Sequence to an array.
	sequences := []*common.Sequence{}
	for sequenceNumber, sequenceConf := range sequencesConfig.Sequences {
		fmt.Printf("Found sequence  name: %s, label:%s desc: %s, type: %s\n", sequenceConf.Name, sequenceConf.Label, sequenceConf.Description, sequenceConf.Type)
		newSequence := sequence.CreateSequence(sequenceConf.Type, sequenceConf.Label, sequenceNumber, fixturesConfig, this.SequenceChannels)

		// Add the name, label and description to the new sequence.
		newSequence.Name = sequenceConf.Name
		newSequence.DisableOnceMutex = &sync.RWMutex{}
		newSequence.Description = sequenceConf.Description
		newSequence.Label = sequenceConf.Label
		newSequence.Type = sequenceConf.Type

		sequences = append(sequences, &newSequence)

		// Setup Default State.
		this.Speed[sequenceNumber] = common.DEFAULT_SPEED                            // Selected speed for the sequence. Common to all types of sequence.
		this.Running[sequenceNumber] = false                                         // Set this sequence to be in the not running state. Common to all types of sequence.
		this.Strobe[sequenceNumber] = false                                          // Set strobe to be off for all sequences.
		this.StrobeSpeed[sequenceNumber] = common.DEFAULT_STROBE_SPEED               // Set the strobe speed to be the fastest for this sequence.
		this.RGBShift[sequenceNumber] = common.DEFAULT_RGB_SHIFT                     // Default RGB shift size.
		this.ScannerShift[sequenceNumber] = common.DEFAULT_SCANNER_SHIFT             // Default scanner shift size.
		this.SequenceType[sequenceNumber] = newSequence.Type                         // Set the sequence type.
		this.RGBSize[sequenceNumber] = common.DEFAULT_RGB_SIZE                       // Set the defaults size for the RGB fixtures.
		this.ScannerSize[sequenceNumber] = common.DEFAULT_SCANNER_SIZE               // Set the defaults size for the scanner fixtures.
		this.RGBFade[sequenceNumber] = common.DEFAULT_RGB_FADE                       // Set the default fade time for RGB fixtures.
		this.ScannerFade[sequenceNumber] = common.DEFAULT_SCANNER_FADE               // Set the default fade time for scanners.
		this.ScannerCoordinates[sequenceNumber] = common.DEFAULT_SCANNER_COORDNIATES // Set the default fade time for scanners.

		if newSequence.Label == "switch" {
			this.SwitchSequenceNumber = sequenceNumber

			// Store the switch Config locally.
			switchConfig := commands.LoadSwitchConfiguration(this.SwitchSequenceNumber, fixturesConfig)

			// Populate each switch with a number of states based on their config.
			for swiTchNumber := 0; swiTchNumber < len(switchConfig); swiTchNumber++ {

				// assign the switch.
				swiTch := switchConfig[swiTchNumber]

				// Now populate the states.
				for stateNumber := 0; stateNumber < len(swiTch.States); stateNumber++ {

					state := swiTch.States[stateNumber]

					// Find the details of the fixture for this switch.
					thisFixture, err := fixture.FindFixtureByLabel(swiTch.UseFixture, fixturesConfig)
					if err != nil {
						fmt.Printf("error %s\n", err.Error())
					}

					// Load the config for this state of of this switch
					override := fixture.DiscoverSwitchOveride(thisFixture, swiTch.Number, int(state.Number), fixturesConfig)

					// Assign this discovered override to the current switch state.
					this.SwitchOverrides[swiTchNumber] = append(this.SwitchOverrides[swiTchNumber], override)

					if debug {
						fmt.Printf("Setting Up Override for Switch No=%d Name=%s State No=%d Name=%s\n", swiTch.Number, swiTch.Name, state.Number, state.Name)
						fmt.Printf("\t Override Colors %+v\n", override.Colors)

					}
				}
			}
		}

		if newSequence.Label == "chaser" {
			this.ChaserSequenceNumber = sequenceNumber
		}

		if newSequence.Type == "scanner" {
			this.ScannerSequenceNumber = sequenceNumber
			common.GlobalScannerSequenceNumber = sequenceNumber
		}
		// Setup Functions Labels.
		if newSequence.Type == "rgb" {
			this.FunctionLabels[0] = "RGB\nPatten"
			this.FunctionLabels[1] = "RGB\nAuto\nColor"
			this.FunctionLabels[2] = "RGB\nAuto\nPatten"
			this.FunctionLabels[3] = "RGB\nBounce"
			this.FunctionLabels[4] = "RGB\nChase\nColor"
			this.FunctionLabels[5] = "RGB\nStatic\nColor"
			this.FunctionLabels[6] = "RGB\nInvert"
			this.FunctionLabels[7] = "RGB\nMusic"
		}

		if newSequence.Type == "scanner" && newSequence.Label != "chaser" {
			this.FunctionLabels[0] = "Scanner\nPatten"
			this.FunctionLabels[1] = "Scanner\nAuto\nColor"
			this.FunctionLabels[2] = "Scanner\nAuto\nPatten"
			this.FunctionLabels[3] = "Scanner\nBounce"
			this.FunctionLabels[4] = "Scanner\nColor"
			this.FunctionLabels[5] = "Scanner\nGobo"
			this.FunctionLabels[6] = "Scanner\nShutter\nChaser"
			this.FunctionLabels[7] = "Scanner\nMusic"
		}

		if newSequence.Type == "rgb" && newSequence.Label == "chaser" {
			this.FunctionLabels[0] = "Chase\nPatten"
			this.FunctionLabels[1] = "Chase\nAuto\nColor"
			this.FunctionLabels[2] = "Chase\nAuto\nPatten"
			this.FunctionLabels[3] = "Chase\nBounce"
			this.FunctionLabels[4] = "Chase\nColor"
			this.FunctionLabels[5] = "Chase\nStatic\nColor"
			this.FunctionLabels[6] = "Chase\nInvert"
			this.FunctionLabels[7] = "Chase\nMusic"
		}

		// Make functions for each of the sequences.
		for function := 0; function < 8; function++ {
			newFunction := common.Function{
				Name:           strconv.Itoa(function),
				SequenceNumber: sequenceNumber,
				Number:         function,
				State:          false,
				Label:          this.FunctionLabels[function],
			}
			this.Functions[sequenceNumber] = append(this.Functions[sequenceNumber], newFunction)
		}
	}

	// Create all the channels I need.
	commandChannels := []chan common.Command{}
	replyChannels := []chan common.Sequence{}
	updateChannels := []chan common.Sequence{}

	// Make four default channels & one for the scanner chaser for commands.
	for range sequences {
		commandChannel := make(chan common.Command)
		commandChannels = append(commandChannels, commandChannel)
		replyChannel := make(chan common.Sequence)
		replyChannels = append(replyChannels, replyChannel)
		updateChannel := make(chan common.Sequence)
		updateChannels = append(updateChannels, updateChannel)
	}

	// SoundTriggers is a an array of switches and channels which control which sequence gets a music trigger.
	this.SoundTriggers = []*common.Trigger{}

	NumberOfMusicTriggers := NumberOfSequences + NumberOfSwitches

	// Setting trigger names.
	for triggerNumber := 0; triggerNumber < NumberOfMusicTriggers; triggerNumber++ {
		newChannel := make(chan common.Command)
		var name string
		var newTrigger common.Trigger
		// First three triggers occupied by sequence 0=FOH, 1=Uplighter,2=Scanners, 3-10 switched 11=shutter chase
		if triggerNumber == 0 {
			name = fmt.Sprintf("sequence%d", triggerNumber) // FOH
		}
		if triggerNumber == 1 {
			name = fmt.Sprintf("sequence%d", triggerNumber) // Uplighters
		}
		if triggerNumber == 2 {
			name = fmt.Sprintf("sequence%d", triggerNumber) // Scanners
		}
		if triggerNumber == 3 {
			name = fmt.Sprintf("sequence%d", triggerNumber) // Switches
		}
		if triggerNumber == 4 {
			name = fmt.Sprintf("sequence%d", triggerNumber) // Shutter Chaser
		}
		// 5-12, eight switches
		if triggerNumber > 4 {
			name = fmt.Sprintf("switch%d", triggerNumber-4)
		}

		newTrigger = common.Trigger{
			Name:    name,
			State:   false,
			Gain:    this.SoundGain,
			Channel: newChannel,
		}

		this.SoundTriggers = append(this.SoundTriggers, &newTrigger)
	}

	if debug {
		for triggerNumber, trigger := range this.SoundTriggers {
			fmt.Printf("%d: trigger %s installed, enabled %t\n", triggerNumber, trigger.Name, trigger.State)
		}
	}

	// Now add them all to a handy channels struct.
	this.SequenceChannels = common.Channels{}
	this.SequenceChannels.CommmandChannels = commandChannels
	this.SequenceChannels.ReplyChannels = replyChannels
	this.SequenceChannels.UpdateChannels = updateChannels
	this.SequenceChannels.SoundTriggers = this.SoundTriggers

	// Create a timer for timing buttons, long and short presses.
	this.ButtonTimer = &time.Time{}

	// Create a sound trigger object and give it the sequences so it can access their configs.
	this.SoundConfig = sound.NewSoundTrigger(this.SequenceChannels, guiButtons, eventsForLaunchpad)

	// Generate the toolbar at the top.
	toolbar := gui.MakeToolbar(myWindow, this.SoundConfig, guiButtons, eventsForLaunchpad, commandChannels, dmxInterfaceConfig, this.LaunchpadName, fixturesConfig, startConfig)

	// Create objects for bottom status bar.
	panel.SpeedLabel = widget.NewLabel(fmt.Sprintf("Speed %02d", common.DEFAULT_SPEED))
	panel.ShiftLabel = widget.NewLabel(fmt.Sprintf("Shift %02d", common.DEFAULT_RGB_SHIFT))
	panel.SizeLabel = widget.NewLabel(fmt.Sprintf("Size %02d", common.DEFAULT_RGB_SIZE))
	panel.FadeLabel = widget.NewLabel(fmt.Sprintf("Fade %02d", common.DEFAULT_RGB_FADE))
	panel.VersionLabel = widget.NewButton("Version 2.1", func() {})
	panel.VersionLabel.Hidden = false
	panel.DisplayMode = widget.NewButton("NORMAL", func() {})
	panel.DisplayMode.Hidden = false
	panel.ColorDisplay.Hidden = false

	// Create objects for top status bar.
	upLabel := widget.NewLabel("       ")
	panel.TiltLabel = upLabel

	redLabel := widget.NewLabel(fmt.Sprintf("Red %02d", 0))
	panel.RedLabel = redLabel

	greenLabel := widget.NewLabel(fmt.Sprintf("Green %02d", 0))
	panel.GreenLabel = greenLabel

	blueLabel := widget.NewLabel(fmt.Sprintf("Blue %02d", 0))
	panel.BlueLabel = blueLabel

	sensitivity := common.FindSensitivity(this.SoundGain)
	sensitivityLabel := widget.NewLabel(fmt.Sprintf("Sensitivity %02d", sensitivity))
	panel.SensitivityLabel = sensitivityLabel

	masterLabel := widget.NewLabel(fmt.Sprintf("Master %02d", this.MasterBrightness))
	panel.MasterLabel = masterLabel

	// Create a thread to handle GUI button events.
	panel.ListenAndSendToGUI(guiButtons, GuiFlashButtons)

	// Now create a thread to handle launchpad light button events.
	launchpad.ListenAndSendToLaunchPad(eventsForLaunchpad, this.Pad, this.LaunchPadConnected)

	// Add buttons to the main panel.
	row0 := panel.GenerateRow(myWindow, 0, sequences, &this, eventsForLaunchpad, guiButtons, dmxController, groupConfig, fixturesConfig, commandChannels, replyChannels, updateChannels, this.DmxInterfacePresent)
	row1 := panel.GenerateRow(myWindow, 1, sequences, &this, eventsForLaunchpad, guiButtons, dmxController, groupConfig, fixturesConfig, commandChannels, replyChannels, updateChannels, this.DmxInterfacePresent)
	row2 := panel.GenerateRow(myWindow, 2, sequences, &this, eventsForLaunchpad, guiButtons, dmxController, groupConfig, fixturesConfig, commandChannels, replyChannels, updateChannels, this.DmxInterfacePresent)
	row3 := panel.GenerateRow(myWindow, 3, sequences, &this, eventsForLaunchpad, guiButtons, dmxController, groupConfig, fixturesConfig, commandChannels, replyChannels, updateChannels, this.DmxInterfacePresent)
	row4 := panel.GenerateRow(myWindow, 4, sequences, &this, eventsForLaunchpad, guiButtons, dmxController, groupConfig, fixturesConfig, commandChannels, replyChannels, updateChannels, this.DmxInterfacePresent)
	row5 := panel.GenerateRow(myWindow, 5, sequences, &this, eventsForLaunchpad, guiButtons, dmxController, groupConfig, fixturesConfig, commandChannels, replyChannels, updateChannels, this.DmxInterfacePresent)
	row6 := panel.GenerateRow(myWindow, 6, sequences, &this, eventsForLaunchpad, guiButtons, dmxController, groupConfig, fixturesConfig, commandChannels, replyChannels, updateChannels, this.DmxInterfacePresent)
	row7 := panel.GenerateRow(myWindow, 7, sequences, &this, eventsForLaunchpad, guiButtons, dmxController, groupConfig, fixturesConfig, commandChannels, replyChannels, updateChannels, this.DmxInterfacePresent)
	row8 := panel.GenerateRow(myWindow, 8, sequences, &this, eventsForLaunchpad, guiButtons, dmxController, groupConfig, fixturesConfig, commandChannels, replyChannels, updateChannels, this.DmxInterfacePresent)

	// Gather all the rows into a container called squares.
	squares := container.New(layout.NewGridLayoutWithRows(gui.ColumnWidth), row0, row1, row2, row3, row4, row5, row6, row7, row8)

	// Create top status bar.
	topStatusBar := container.New(layout.NewHBoxLayout(),
		panel.ColorDisplay,
		upLabel,
		redLabel,
		greenLabel,
		blueLabel,
		layout.NewSpacer(),
		layout.NewSpacer(),
		layout.NewSpacer(),
		layout.NewSpacer(),
		sensitivityLabel,
		layout.NewSpacer(),
		layout.NewSpacer(),
		masterLabel,
		layout.NewSpacer(),
		layout.NewSpacer(),
		layout.NewSpacer(),
		toolbar,
	)

	// Create bottom status bar.
	bottonStatusBar := container.New(
		layout.NewHBoxLayout(), panel.DisplayMode, layout.NewSpacer(), panel.SpeedLabel, layout.NewSpacer(), panel.ShiftLabel, layout.NewSpacer(), panel.SizeLabel, layout.NewSpacer(), panel.FadeLabel, layout.NewSpacer(), panel.VersionLabel)

	// Now configure the panel content to contain the top toolbar and the squares.
	main := container.NewBorder(topStatusBar, nil, nil, nil, squares)
	content := container.NewBorder(main, nil, nil, nil, bottonStatusBar)

	// Start threads for each sequence.
	go sequence.PlaySequence(*sequences[0], 0, this.RGBPatterns, eventsForLaunchpad, guiButtons, dmxController, fixturesConfig, this.SequenceChannels, this.SwitchChannels, this.SoundConfig, this.DmxInterfacePresent)
	go sequence.PlaySequence(*sequences[1], 1, this.RGBPatterns, eventsForLaunchpad, guiButtons, dmxController, fixturesConfig, this.SequenceChannels, this.SwitchChannels, this.SoundConfig, this.DmxInterfacePresent)
	go sequence.PlaySequence(*sequences[2], 2, this.RGBPatterns, eventsForLaunchpad, guiButtons, dmxController, fixturesConfig, this.SequenceChannels, this.SwitchChannels, this.SoundConfig, this.DmxInterfacePresent)
	go sequence.PlaySequence(*sequences[3], 3, this.RGBPatterns, eventsForLaunchpad, guiButtons, dmxController, fixturesConfig, this.SequenceChannels, this.SwitchChannels, this.SoundConfig, this.DmxInterfacePresent)
	go sequence.PlaySequence(*sequences[4], 4, this.RGBPatterns, eventsForLaunchpad, guiButtons, dmxController, fixturesConfig, this.SequenceChannels, this.SwitchChannels, this.SoundConfig, this.DmxInterfacePresent)

	// Light the first sequence as the default selected.
	this.SelectedSequence = 0
	this.SelectedType = "rgb"
	buttons.InitButtons(&this, sequences[0].SequenceColors, []color.RGBA{}, eventsForLaunchpad, guiButtons)

	// Label the right hand buttons.
	panel.LabelRightHandButtons()

	// Clear the pad. Strobe is set to 0.
	buttons.AllFixturesOff(sequences, eventsForLaunchpad, guiButtons, dmxController, fixturesConfig, this.DmxInterfacePresent)
	buttons.Clear(0, 0, &this, sequences, dmxController, fixturesConfig, commandChannels, eventsForLaunchpad, guiButtons, updateChannels)

	// If present create a thread to listen to launchpad button events.
	if this.LaunchPadConnected {
		go func(guiButtons chan common.ALight,
			this *buttons.CurrentState,
			sequences []*common.Sequence,
			eventsForLaunchpad chan common.ALight,
			dmxController *ft232.DMXController,
			fixturesConfig *fixture.Fixtures,
			commandChannels []chan common.Command,
			replyChannels []chan common.Sequence,
			updateChannels []chan common.Sequence,
			dmxInterfaceCardPresent bool) {

			launchpad.ReadLaunchPadButtons(guiButtons, this, sequences, eventsForLaunchpad, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels, dmxInterfaceCardPresent)

		}(guiButtons, &this, sequences, eventsForLaunchpad, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels, this.DmxInterfacePresent)
	}

	// Show this sequence running status in the start/stop button.
	common.ShowRunningStatus(this.Running[this.SelectedSequence], eventsForLaunchpad, guiButtons)
	common.ShowStrobeButtonStatus(this.Strobe[this.SelectedSequence], eventsForLaunchpad, guiButtons)

	myWindow.SetContent(content)

	// Main menu.
	openProject := fyne.NewMenuItem("Open", func() {
		gui.FileOpen(myWindow, startConfig, fixturesConfig, commandChannels)
	})
	saveProject := fyne.NewMenuItem("Save", func() {
		gui.FileSave(myWindow, startConfig, fixturesConfig, commandChannels)
	})
	editSettings := fyne.NewMenuItem("Edit", func() {
		modal := gui.RunSettingsPopUp(myWindow, this.SoundConfig, guiButtons, eventsForLaunchpad, dmxInterfaceConfig, this.LaunchpadName)
		modal.Resize(fyne.NewSize(250, 250))
		modal.Show()
	})
	editFixtures := fyne.NewMenuItem("Edit", func() {
		gui.NewFixtureEditor(sequences, myWindow, groupConfig, fixturesConfig, commandChannels)
	})
	projectMenu := fyne.NewMenu("Project", openProject, saveProject)
	settingsMenu := fyne.NewMenu("Settings", editSettings)
	fixturesMenu := fyne.NewMenu("Fixtures", editFixtures)
	helpMenu := fyne.NewMenu("Help")
	mainMenu := fyne.NewMainMenu(projectMenu, settingsMenu, fixturesMenu, helpMenu)
	myWindow.SetMainMenu(mainMenu)

	myWindow.ShowAndRun()

	fmt.Println("Saving Presets")
	presets.SavePresets(this.PresetsStore)

}

func makeStaticButtonsStorage() []common.StaticColorButton {

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

	return staticButtons
}
