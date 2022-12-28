package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/dhowlett99/dmxlights/pkg/buttons"
	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/dmx"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
	"github.com/dhowlett99/dmxlights/pkg/gui"
	"github.com/dhowlett99/dmxlights/pkg/pattern"
	"github.com/dhowlett99/dmxlights/pkg/presets"
	"github.com/dhowlett99/dmxlights/pkg/sequence"
	"github.com/dhowlett99/dmxlights/pkg/sound"
	"github.com/oliread/usbdmx/ft232"
	"github.com/rakyll/launchpad/mk3"
)

func main() {

	fmt.Println("DMX Lighting")

	os.Setenv("FYNE_THEME", "light")

	// Start the GUI.
	fmt.Println("Starting GUI")
	panel := gui.NewPanel() // Panel represents the buttons in the GUI.
	myApp := app.New()

	myWindow := myApp.NewWindow("DMX Lights")
	myWindow.Resize(fyne.NewSize(400, 50))

	// Setup the current state.
	this := buttons.CurrentState{}

	this.Blackout = false                              // Blackout starts in off.
	this.Flood = false                                 // Flood starts in off.
	this.Running = make(map[int]bool, 4)               // Initialise storage for four sequences.
	this.Strobe = make(map[int]bool, 4)                // Initialise storage for four sequences.
	this.MasterBrightness = 255                        // Affects all DMX fixtures and launchpad lamps.
	this.SoundGain = 0                                 // Fine gain -0.09 -> 0.09
	this.OffsetPan = 120                               // Start pan from the center
	this.OffsetTilt = 120                              // Start tilt from the center.
	this.Patterns = pattern.MakePatterns()             // Build the default set of Patterns.
	this.SelectButtonPressed = make([]bool, 4)         // Initialise four select buttons.
	this.FunctionSelectMode = make([]bool, 4)          // Initialise four function mode states.
	this.EditSequenceColorsMode = make([]bool, 4)      // Remember when we are in editing sequence colors mode.
	this.EditScannerColorsMode = make([]bool, 4)       // Remember when we are in setting scanner color mode.
	this.EditGoboSelectionMode = make([]bool, 4)       // Remember when we are in selecting gobo mode.
	this.EditStaticColorsMode = make([]bool, 4)        // Remember when we are in editing static colors mode.
	this.EditPatternMode = make([]bool, 4)             // Remember when we are in editing pattern mode.
	this.StaticButtons = makeStaticButtonsStorage()    // Make storgage for color editing button results.
	this.PresetsStore = presets.LoadPresets()          // Load the presets from their json files.
	this.Speed = make(map[int]int, 4)                  // Initialise storage for four sequences.
	this.RGBSize = make(map[int]int, 4)                // Initialise storage for four sequences.
	this.ScannerSize = make(map[int]int, 4)            // Initialise storage for four sequences.
	this.RGBShift = make(map[int]int, 4)               // Initialise storage for four sequences.
	this.ScannerShift = make(map[int]int, 4)           // Initialise storage for four sequences.
	this.RGBFade = make(map[int]int, 4)                // Initialise storage for four sequences.
	this.ScannerFade = make(map[int]int, 4)            // Initialise storage for four sequences.
	this.StrobeSpeed = make(map[int]int, 4)            // Initialise storage for four sequences.
	this.ScannerCoordinates = make(map[int]int, 4)     // Number of coordinates for scanner patterns is selected from 4 choices. 0=12, 1=16,2=24,3=32,4=64
	this.LaunchPadConnected = true                     // Assume launchpad is present, until tested.
	this.DmxInterfacePresent = true                    // Assume DMX interface card is present, until tested.
	this.LaunchpadName = "Novation Launchpad Mk3 Mini" // Name of launchpad.
	// Now add channels to communicate with mini-sequencers on switch channels.
	this.SwitchChannels = make(map[int]common.SwitchChannel, 10)
	for switchChannel := 0; switchChannel < 10; switchChannel++ {
		newSwitch := common.SwitchChannel{}
		newSwitch.Stop = make(chan bool)
		newSwitch.KeepRotateAlive = make(chan bool)
		newSwitch.StopRotate = make(chan bool)
		newSwitch.SequencerRunning = false
		this.SwitchChannels[switchChannel] = newSwitch
	}

	// Initialize eight fixture states for the four sequences.
	this.ScannerState = make([][]common.ScannerState, 9)
	for x := 0; x < 9; x++ {
		this.ScannerState[x] = make([]common.ScannerState, 9)
		for y := 0; y < 9; y++ {
			newScanner := common.ScannerState{}
			newScanner.Enabled = true
			newScanner.Inverted = false
			this.ScannerState[x][y] = newScanner
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
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		presets.SavePresets(this.PresetsStore)
		os.Exit(1)
	}()

	// Setup a connection to the Novation Launchpad.
	// Tested with a Novation Launchpad mini mk3.
	fmt.Println("Setup Novation Launchpad")
	this.Pad, err = mk3.Open()
	if err != nil {
		fmt.Printf("launchpad: %v", err)
		this.LaunchPadConnected = false
		this.LaunchpadName = "Not Found"
	}

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
	fixturesConfig, err := fixture.LoadFixtures("fixtures.yaml")
	if err != nil {
		fmt.Printf("dmxlights: error failed to load fixtures: %s\n", err.Error())
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

	// Create the sequences from config file.
	// Add Sequence to an array.
	sequences := []*common.Sequence{}
	for sequenceNumber, sequenceConf := range sequencesConfig.Sequences {
		fmt.Printf("Found sequence  name: %s, label:%s desc: %s, type: %s\n", sequenceConf.Name, sequenceConf.Label, sequenceConf.Description, sequenceConf.Type)
		newSequence := sequence.CreateSequence(sequenceConf.Type, sequenceNumber, this.Patterns, fixturesConfig, this.SequenceChannels)

		// Add the name, label and description to the new sequence.
		newSequence.Name = sequenceConf.Name
		newSequence.Description = sequenceConf.Description
		newSequence.Label = sequenceConf.Label
		newSequence.Type = sequenceConf.Type

		sequences = append(sequences, &newSequence)

		// Setup Default State.
		this.Speed[sequenceNumber] = common.DefaultSpeed                           // Selected speed for the sequence. Common to all types of sequence.
		this.Running[sequenceNumber] = false                                       // Set this sequence to be in the not running state. Common to all types of sequence.
		this.Strobe[sequenceNumber] = false                                        // Set strobe to be off for all sequences.
		this.StrobeSpeed[sequenceNumber] = 255                                     // Set the strob speed to be the fastest for this sequence.
		this.RGBShift[sequenceNumber] = common.DefaultRGBShift                     // Default RGB shift size.
		this.ScannerShift[sequenceNumber] = common.DefaultScannerShift             // Default scanner shift size.
		this.RGBSize[sequenceNumber] = common.DefaultRGBSize                       // Set the defaults size for the RGB fixtures.
		this.ScannerSize[sequenceNumber] = common.DefaultScannerSize               // Set the defaults size for the scanner fixtures.
		this.RGBFade[sequenceNumber] = common.DefaultRGBFade                       // Set the default fade time for RGB fixtures.
		this.ScannerFade[sequenceNumber] = common.DefaultScannerFade               // Set the default fade time for scanners.
		this.ScannerCoordinates[sequenceNumber] = common.DefaultScannerCoordinates // Set the default fade time for scanners.
	}

	// Create all the channels I need.
	commandChannels := []chan common.Command{}
	replyChannels := []chan common.Sequence{}
	updateChannels := []chan common.Sequence{}

	// Make four default channels for commands.
	for sequenceNumber := 0; sequenceNumber < 4; sequenceNumber++ {
		commandChannel := make(chan common.Command)
		commandChannels = append(commandChannels, commandChannel)
		replyChannel := make(chan common.Sequence)
		replyChannels = append(replyChannels, replyChannel)
		updateChannel := make(chan common.Sequence)
		updateChannels = append(updateChannels, updateChannel)
	}

	// SoundTriggers is a an array of switches and channels which control which sequence gets a music trigger.
	this.SoundTriggers = []*common.Trigger{}

	for triggerNumber := 0; triggerNumber < 10; triggerNumber++ {
		newChannel := make(chan common.Command)
		var name string
		var newTrigger common.Trigger
		if triggerNumber < 3 {
			name = fmt.Sprintf("sequence%d", triggerNumber)
		} else {
			name = fmt.Sprintf("switch%d", triggerNumber-3)
		}
		newTrigger = common.Trigger{
			Name:    name,
			State:   false,
			Gain:    this.SoundGain,
			Channel: newChannel,
		}

		this.SoundTriggers = append(this.SoundTriggers, &newTrigger)
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
	this.SoundConfig = sound.NewSoundTrigger(this.SequenceChannels, guiButtons)

	// Generate the toolbar at the top.
	toolbar := gui.MakeToolbar(myWindow, this.SoundConfig, guiButtons, dmxInterfaceConfig, this.LaunchpadName)

	// Create objects for bottom status bar.
	panel.SpeedLabel = widget.NewLabel(fmt.Sprintf("Speed %02d", common.DefaultSpeed))
	panel.ShiftLabel = widget.NewLabel(fmt.Sprintf("Shift %02d", common.DefaultRGBShift))
	panel.SizeLabel = widget.NewLabel(fmt.Sprintf("Size %02d", common.DefaultRGBSize))
	panel.FadeLabel = widget.NewLabel(fmt.Sprintf("Fade %02d", common.DefaultRGBFade))
	panel.BeatLabel = widget.NewButton("BEAT", func() {})
	panel.BeatLabel.Hidden = true

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
	go func(panel gui.MyPanel, guiButtons chan common.ALight, GuiFlashButtons [][]common.ALight) {
		for {
			alight := <-guiButtons
			panel.UpdateButtonColor(alight, GuiFlashButtons)
		}
	}(panel, guiButtons, GuiFlashButtons)

	// Now create a thread to handle launchpad light button events.
	go func() {
		common.ListenAndSendToLaunchPad(eventsForLaunchpad, this.Pad, this.LaunchPadConnected)
	}()

	// Add buttons to the main panel.
	row0 := panel.GenerateRow(myWindow, 0, sequences, &this, eventsForLaunchpad, guiButtons, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels, this.DmxInterfacePresent)
	row1 := panel.GenerateRow(myWindow, 1, sequences, &this, eventsForLaunchpad, guiButtons, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels, this.DmxInterfacePresent)
	row2 := panel.GenerateRow(myWindow, 2, sequences, &this, eventsForLaunchpad, guiButtons, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels, this.DmxInterfacePresent)
	row3 := panel.GenerateRow(myWindow, 3, sequences, &this, eventsForLaunchpad, guiButtons, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels, this.DmxInterfacePresent)
	row4 := panel.GenerateRow(myWindow, 4, sequences, &this, eventsForLaunchpad, guiButtons, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels, this.DmxInterfacePresent)
	row5 := panel.GenerateRow(myWindow, 5, sequences, &this, eventsForLaunchpad, guiButtons, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels, this.DmxInterfacePresent)
	row6 := panel.GenerateRow(myWindow, 6, sequences, &this, eventsForLaunchpad, guiButtons, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels, this.DmxInterfacePresent)
	row7 := panel.GenerateRow(myWindow, 7, sequences, &this, eventsForLaunchpad, guiButtons, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels, this.DmxInterfacePresent)
	row8 := panel.GenerateRow(myWindow, 8, sequences, &this, eventsForLaunchpad, guiButtons, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels, this.DmxInterfacePresent)

	// Gather all the rows into a container called squares.
	squares := container.New(layout.NewGridLayoutWithRows(gui.ColumnWidth), row0, row1, row2, row3, row4, row5, row6, row7, row8)

	topStatusBar := fyne.NewContainerWithLayout(
		layout.NewHBoxLayout(),
		layout.NewSpacer(),
		upLabel,
		redLabel,
		greenLabel,
		blueLabel,
		layout.NewSpacer(),
		layout.NewSpacer(),
		layout.NewSpacer(),
		sensitivityLabel,
		layout.NewSpacer(),
		masterLabel,
		layout.NewSpacer(),
		layout.NewSpacer(),
		layout.NewSpacer(),
		toolbar)

	bottonStatusBar := fyne.NewContainerWithLayout(
		layout.NewHBoxLayout(), panel.SpeedLabel, layout.NewSpacer(), panel.ShiftLabel, layout.NewSpacer(), panel.SizeLabel, layout.NewSpacer(), panel.FadeLabel, layout.NewSpacer(), panel.BeatLabel)

	// Now configure the panel content to contain the top toolbar and the squares.
	main := container.NewBorder(topStatusBar, nil, nil, nil, squares)
	content := container.NewBorder(main, nil, nil, nil, bottonStatusBar)

	// Start threads for each sequence.
	go sequence.PlaySequence(*sequences[0], 0, this.Pad, eventsForLaunchpad, guiButtons, dmxController, fixturesConfig, this.SequenceChannels, this.SwitchChannels, this.SoundConfig, this.DmxInterfacePresent)
	go sequence.PlaySequence(*sequences[1], 1, this.Pad, eventsForLaunchpad, guiButtons, dmxController, fixturesConfig, this.SequenceChannels, this.SwitchChannels, this.SoundConfig, this.DmxInterfacePresent)
	go sequence.PlaySequence(*sequences[2], 2, this.Pad, eventsForLaunchpad, guiButtons, dmxController, fixturesConfig, this.SequenceChannels, this.SwitchChannels, this.SoundConfig, this.DmxInterfacePresent)
	go sequence.PlaySequence(*sequences[3], 3, this.Pad, eventsForLaunchpad, guiButtons, dmxController, fixturesConfig, this.SequenceChannels, this.SwitchChannels, this.SoundConfig, this.DmxInterfacePresent)

	// Light the first sequence as the default selected.
	this.SelectedSequence = 0
	buttons.InitButtons(&this, eventsForLaunchpad, guiButtons)

	// Label the right hand buttons.
	panel.LabelRightHandButtons()

	// Clear the pad. Strobe is set to 0.
	buttons.AllFixturesOff(sequences, eventsForLaunchpad, guiButtons, dmxController, fixturesConfig, this.DmxInterfacePresent)

	if this.LaunchPadConnected {
		// Listen to launchpad buttons.
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

			buttons.ReadLaunchPadButtons(guiButtons, this, sequences, eventsForLaunchpad, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels, dmxInterfaceCardPresent)

		}(guiButtons, &this, sequences, eventsForLaunchpad, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels, this.DmxInterfacePresent)
	}

	// Initially set the Flood, Save, Start, Stop and Blackout buttons to white.
	common.LightLamp(common.ALight{X: 8, Y: 3, Red: 255, Green: 255, Blue: 255, Brightness: 255}, eventsForLaunchpad, guiButtons)
	common.LightLamp(common.ALight{X: 8, Y: 4, Red: 255, Green: 255, Blue: 255, Brightness: 255}, eventsForLaunchpad, guiButtons)
	common.LightLamp(common.ALight{X: 8, Y: 5, Red: 255, Green: 255, Blue: 255, Brightness: 255}, eventsForLaunchpad, guiButtons)
	common.LightLamp(common.ALight{X: 8, Y: 6, Red: 255, Green: 255, Blue: 255, Brightness: 255}, eventsForLaunchpad, guiButtons)
	common.LightLamp(common.ALight{X: 8, Y: 7, Red: 255, Green: 255, Blue: 255, Brightness: 255}, eventsForLaunchpad, guiButtons)

	// Show this sequence running status in the start/stop button.
	common.ShowRunningStatus(this.SelectedSequence, this.Running, eventsForLaunchpad, guiButtons)

	myWindow.SetContent(content)

	myWindow.ShowAndRun()

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
