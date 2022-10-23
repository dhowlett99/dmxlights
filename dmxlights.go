package main

import (
	"fmt"
	"log"
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

	this.Blackout = false                           // Blackout starts in off.
	this.Flood = false                              // Flood starts in off.
	this.Running = make(map[int]bool, 4)            // Initialise storage for four sequences.
	this.MasterBrightness = 255                     // Affects all DMX fixtures and launchpad lamps.
	this.SoundGain = 0                              // Fine gain -0.09 -> 0.09
	this.OffsetPan = 120                            // Start pan from the center
	this.OffsetTilt = 120                           // Start tilt from the center.
	this.Patterns = pattern.MakePatterns()          // Build the default set of Patterns.
	this.SelectButtonPressed = make([]bool, 4)      // Initialise four select buttons.
	this.FunctionSelectMode = make([]bool, 4)       // Initialise four function mode states.
	this.EditSequenceColorsMode = make([]bool, 4)   // Remember when we are in editing sequence colors mode.
	this.EditScannerColorsMode = make([]bool, 4)    // Remember when we are in setting scanner color mode.
	this.EditGoboSelectionMode = make([]bool, 4)    // Remember when we are in selecting gobo mode.
	this.EditStaticColorsMode = make([]bool, 4)     // Remember when we are in editing static colors mode.
	this.EditPatternMode = make([]bool, 4)          // Remember when we are in editing pattern mode.
	this.StaticButtons = makeStaticButtonsStorage() // Make storgage for color editing button results.
	this.PresetsStore = presets.LoadPresets()       // Load the presets from their json files.
	this.Speed = make(map[int]int, 4)               // Initialise storage for four sequences.
	this.RGBSize = make(map[int]int, 4)             // Initialise storage for four sequences.
	this.ScannerSize = make(map[int]int, 4)         // Initialise storage for four sequences.
	this.RGBShift = make(map[int]int, 4)            // Initialise storage for four sequences.
	this.ScannerShift = make(map[int]int, 4)        // Initialise storage for four sequences.
	this.RGBFade = make(map[int]int, 4)             // Initialise storage for four sequences.
	this.ScannerFade = make(map[int]int, 4)         // Initialise storage for four sequences.
	this.ScannerCoordinates = make(map[int]int, 4)  // Number of coordinates for scanner patterns is selected from 4 choices. 0=12, 1=16,2=24,3=32
	this.SwitchStopChannel = make(chan bool)        // Channel to stop switch fixtures.

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
	dmxController, err := dmx.NewDmXController()
	if err != nil {
		fmt.Printf("error initializing dmx interface: %v\n", err)
		os.Exit(1)
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
		log.Fatalf("error initializing launchpad: %v", err)
	}
	defer this.Pad.Close()

	// Create a channel to send events to the launchpad.
	eventsForLaunchpad := make(chan common.ALight)

	// We need to be in programmers mode to use the launchpad.
	this.Pad.Program()

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
	fixturesConfig, err := fixture.LoadFixtures()
	if err != nil {
		fmt.Printf("dmxlights: error failed to load fixtures: %s\n", err.Error())
		os.Exit(1)
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

		sequences = append(sequences, &newSequence)

		// Setup Default State.
		this.Speed[sequenceNumber] = common.DefaultSpeed                           // Selected speed for the sequence. Common to all types of sequence.
		this.Running[sequenceNumber] = true                                        // Set this sequence to be in the running state. Common to all types of sequence.
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
	this.SequenceChannels = common.Channels{}
	this.SequenceChannels.CommmandChannels = commandChannels
	this.SequenceChannels.ReplyChannels = replyChannels
	this.SequenceChannels.SoundTriggerChannels = soundTriggerChannels
	this.SequenceChannels.UpdateChannels = updateChannels

	// this.SoundTriggers  is a an array of switches which control which sequence gets a music trigger.
	this.SoundTriggers = []*common.Trigger{}
	this.SoundTriggers = append(this.SoundTriggers, &common.Trigger{SequenceNumber: 0, State: false, Gain: this.SoundGain})
	this.SoundTriggers = append(this.SoundTriggers, &common.Trigger{SequenceNumber: 1, State: false, Gain: this.SoundGain})
	this.SoundTriggers = append(this.SoundTriggers, &common.Trigger{SequenceNumber: 2, State: false, Gain: this.SoundGain})
	this.SoundTriggers = append(this.SoundTriggers, &common.Trigger{SequenceNumber: 3, State: false, Gain: this.SoundGain})

	// Create a timer for timing buttons, long and short presses.
	this.ButtonTimer = &time.Time{}

	// Create a sound trigger object and give it the sequences so it can access their configs.
	soundConfig := sound.NewSoundTrigger(this.SoundTriggers, this.SequenceChannels)

	// Generate the toolbar at the top.
	toolbar := gui.MakeToolbar(myWindow, soundConfig)

	// Create objects for bottom status bar.
	speedLabel := widget.NewLabel(fmt.Sprintf("Speed %02d", common.DefaultSpeed))
	panel.SpeedLabel = speedLabel

	shiftLabel := widget.NewLabel(fmt.Sprintf("Shift %02d", common.DefaultRGBShift))
	panel.ShiftLabel = shiftLabel

	sizeLabel := widget.NewLabel(fmt.Sprintf("Size %02d", common.DefaultRGBSize))
	panel.SizeLabel = sizeLabel

	fadeLabel := widget.NewLabel(fmt.Sprintf("Fade %02d", common.DefaultRGBFade))
	panel.FadeLabel = fadeLabel

	bpmLabel := widget.NewLabel(fmt.Sprintf("BPM %03d", 0))
	panel.BPMLabel = bpmLabel

	// Create a thread to handle GUI button events.
	go func(panel gui.MyPanel, guiButtons chan common.ALight, GuiFlashButtons [][]common.ALight) {
		for {
			alight := <-guiButtons
			panel.UpdateButtonColor(alight, GuiFlashButtons)
		}
	}(panel, guiButtons, GuiFlashButtons)

	// Now create a thread to handle launchpad light button events.
	go func() {
		common.ListenAndSendToLaunchPad(eventsForLaunchpad, this.Pad)
	}()

	// Add buttons to the main panel.
	row0 := panel.GenerateRow(myWindow, 0, sequences, &this, eventsForLaunchpad, guiButtons, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels)
	row1 := panel.GenerateRow(myWindow, 1, sequences, &this, eventsForLaunchpad, guiButtons, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels)
	row2 := panel.GenerateRow(myWindow, 2, sequences, &this, eventsForLaunchpad, guiButtons, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels)
	row3 := panel.GenerateRow(myWindow, 3, sequences, &this, eventsForLaunchpad, guiButtons, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels)
	row4 := panel.GenerateRow(myWindow, 4, sequences, &this, eventsForLaunchpad, guiButtons, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels)
	row5 := panel.GenerateRow(myWindow, 5, sequences, &this, eventsForLaunchpad, guiButtons, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels)
	row6 := panel.GenerateRow(myWindow, 6, sequences, &this, eventsForLaunchpad, guiButtons, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels)
	row7 := panel.GenerateRow(myWindow, 7, sequences, &this, eventsForLaunchpad, guiButtons, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels)
	row8 := panel.GenerateRow(myWindow, 8, sequences, &this, eventsForLaunchpad, guiButtons, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels)

	// Gather all the rows into a container called squares.
	squares := container.New(layout.NewGridLayoutWithRows(gui.ColumnWidth), row0, row1, row2, row3, row4, row5, row6, row7, row8)

	statusBar := fyne.NewContainerWithLayout(
		layout.NewHBoxLayout(), speedLabel, layout.NewSpacer(), shiftLabel, layout.NewSpacer(), sizeLabel, layout.NewSpacer(), fadeLabel, layout.NewSpacer(), bpmLabel)

	// Now configure the panel content to contain the top toolbar and the squares.
	main := container.NewBorder(toolbar, nil, nil, nil, squares)
	content := container.NewBorder(main, nil, nil, nil, statusBar)

	// Start threads for each sequence.
	go sequence.PlaySequence(*sequences[0], 0, this.Pad, eventsForLaunchpad, guiButtons, dmxController, fixturesConfig, this.SequenceChannels, this.SoundTriggers, this.SwitchStopChannel)
	go sequence.PlaySequence(*sequences[1], 1, this.Pad, eventsForLaunchpad, guiButtons, dmxController, fixturesConfig, this.SequenceChannels, this.SoundTriggers, this.SwitchStopChannel)
	go sequence.PlaySequence(*sequences[2], 2, this.Pad, eventsForLaunchpad, guiButtons, dmxController, fixturesConfig, this.SequenceChannels, this.SoundTriggers, this.SwitchStopChannel)
	go sequence.PlaySequence(*sequences[3], 3, this.Pad, eventsForLaunchpad, guiButtons, dmxController, fixturesConfig, this.SequenceChannels, this.SoundTriggers, this.SwitchStopChannel)

	// Light the first sequence as the default selected.
	this.SelectedSequence = 0
	buttons.InitButtons(&this, eventsForLaunchpad, guiButtons)

	// Label the right hand buttons.
	panel.LabelRightHandButtons()

	// Clear the pad. Strobe is set to 0.
	buttons.AllFixturesOff(sequences, eventsForLaunchpad, guiButtons, dmxController, fixturesConfig, 0)

	// Listen to launchpad buttons.
	go func(guiButtons chan common.ALight,
		this *buttons.CurrentState,
		sequences []*common.Sequence,
		eventsForLaunchpad chan common.ALight,
		dmxController *ft232.DMXController,
		fixturesConfig *fixture.Fixtures,
		commandChannels []chan common.Command,
		replyChannels []chan common.Sequence,
		updateChannels []chan common.Sequence) {

		buttons.ReadLaunchPadButtons(guiButtons, this, sequences, eventsForLaunchpad, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels)

	}(guiButtons, &this, sequences, eventsForLaunchpad, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels)

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
