package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/dhowlett99/dmxlights/pkg/buttons"
	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/dmx"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
	"github.com/dhowlett99/dmxlights/pkg/gui"
	"github.com/dhowlett99/dmxlights/pkg/pad"
	"github.com/dhowlett99/dmxlights/pkg/patten"
	"github.com/dhowlett99/dmxlights/pkg/presets"
	"github.com/dhowlett99/dmxlights/pkg/sequence"
	"github.com/dhowlett99/dmxlights/pkg/sound"
	"github.com/oliread/usbdmx/ft232"
)

func main() {

	// Start the GUI.
	fmt.Println("Start GUI")
	panel := gui.NewPanel() // Panel represents the buttons in the GUI.
	myApp := app.New()
	myWindow := myApp.NewWindow("DMX Lights")
	myLogo := panel.ConvertButtonImageToIcon("dmxlights.png")
	myWindow.Resize(fyne.NewSize(400, 50))

	// Generate the toolbar at the top.
	toolbar := gui.MakeToolbar(myLogo)

	// Setup the current state
	this := buttons.CurrentState{}

	fmt.Println("DMX Lighting")

	// Setup State.
	this.SequenceSpeed = 12                         // Selected speed for the sequence.
	this.ScannerSize = common.DefaultScannerSize    // Default scanner size.
	this.SelectedShift = 0                          // Default shift size.
	this.Blackout = false                           // Blackout starts in off.
	this.Flood = false                              // Flood starts in off.
	this.FadeSpeed = 11                             // Default start at 50ms.
	this.MasterBrightness = 255                     // Affects all DMX fixtures and launchpad lamps.
	this.SoundGain = 0                              // Fine gain -0.09 -> 0.09
	this.SelectedCordinates = 0                     // Number of coordinates for scanner patterns is selected from 4 choices. 0=12, 1=26,2=24,3=32
	this.OffsetPan = 120                            // Start pan from the center
	this.OffsetTilt = 120                           // Start tilt from the center.
	this.SelectedFloodMap = make(map[int]bool, 4)   // Make a store for which sequences can be flood light.
	this.Pattens = patten.MakePatterns()            // Build the default set of Pattens.
	this.SelectButtonPressed = make([]bool, 4)      // Initialize four select buttons.
	this.FunctionSelectMode = make([]bool, 4)       // Initialize four function mode states.
	this.EditSequenceColorsMode = make([]bool, 4)   // Remember when we are in editing sequence colors mode.
	this.EditScannerColorsMode = make([]bool, 4)    // Remember when we are in setting scanner color mode.
	this.EditGoboSelectionMode = make([]bool, 4)    // Remember when we are in selecting gobo mode.
	this.EditStaticColorsMode = make([]bool, 4)     // Remember when we are in editing static colors mode.
	this.EditPattenMode = make([]bool, 4)           // Remember when we are in editing patten mode.
	this.StaticButtons = makeStaticButtonsStorage() // Make storgage for color editing button results.
	this.PresetsStore = presets.LoadPresets()       // Load the presets from their json files.

	// Initialize eight fixture states for the four sequences.
	this.DisabledFixture = make([][]bool, 9)
	for i := 0; i < 9; i++ {
		this.DisabledFixture[i] = make([]bool, 9)
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
	this.Pad, err = pad.Open()
	if err != nil {
		log.Fatalf("error initializing launchpad: %v", err)
	}
	defer this.Pad.Close()

	// Create a channel to send events to the launchpad.
	eventsForLauchpad := make(chan common.ALight)

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

	// Make space for info on which Launchpad button is flashing.
	LaunchPadFlashButtons := make([][]common.ALight, 10)
	for y := 0; y < 10; y++ {
		LaunchPadFlashButtons[y] = make([]common.ALight, 10)
		for x := 0; x < 10; x++ {
			// Make a stop flashing channel for every button.
			LaunchPadFlashButtons[y][x].FlashStopChannel = make(chan bool)
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
	for index, sequenceConf := range sequencesConfig.Sequences {
		fmt.Printf("Found sequence  name: %s, label:%s desc: %s, type: %s\n", sequenceConf.Name, sequenceConf.Label, sequenceConf.Description, sequenceConf.Type)
		if sequenceConf.Type == "rgb" {
			this.SelectedFloodMap[index] = true // This sequence is this.Flood able because it's a rgb.
		}
		newSequence := sequence.CreateSequence(sequenceConf.Type, index, this.Pattens, fixturesConfig, this.SequenceChannels)

		// Add the name, label and description to the new sequence.
		newSequence.Name = sequenceConf.Name
		newSequence.Description = sequenceConf.Description
		newSequence.Label = sequenceConf.Label

		sequences = append(sequences, &newSequence)
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

	// Create a sound trigger object and give it the sequences so it can access their configs.
	sound.NewSoundTrigger(this.SoundTriggers, this.SequenceChannels)

	// Create a thread to handle GUI button events.
	go func(panel gui.MyPanel, guiButtons chan common.ALight, GuiFlashButtons [][]common.ALight) {
		for {
			alight := <-guiButtons
			panel.UpdateButtonColor(alight, GuiFlashButtons)
		}
	}(panel, guiButtons, GuiFlashButtons)

	// Now create a thread to handle launchpad light button events.
	go func(eventsForLauchpad chan common.ALight, pad *pad.Pad, LaunchPadFlashButtons [][]common.ALight) {
		common.ListenAndSendToLaunchPad(eventsForLauchpad, this.Pad, LaunchPadFlashButtons)
	}(eventsForLauchpad, this.Pad, LaunchPadFlashButtons)

	// Add buttons to the main panel.
	row0 := panel.GenerateRow(myWindow, 0, sequences, &this, eventsForLauchpad, guiButtons, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels)
	row1 := panel.GenerateRow(myWindow, 1, sequences, &this, eventsForLauchpad, guiButtons, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels)
	row2 := panel.GenerateRow(myWindow, 2, sequences, &this, eventsForLauchpad, guiButtons, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels)
	row3 := panel.GenerateRow(myWindow, 3, sequences, &this, eventsForLauchpad, guiButtons, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels)
	row4 := panel.GenerateRow(myWindow, 4, sequences, &this, eventsForLauchpad, guiButtons, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels)
	row5 := panel.GenerateRow(myWindow, 5, sequences, &this, eventsForLauchpad, guiButtons, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels)
	row6 := panel.GenerateRow(myWindow, 6, sequences, &this, eventsForLauchpad, guiButtons, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels)
	row7 := panel.GenerateRow(myWindow, 7, sequences, &this, eventsForLauchpad, guiButtons, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels)
	row8 := panel.GenerateRow(myWindow, 8, sequences, &this, eventsForLauchpad, guiButtons, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels)

	// Gather all the rows into a container called squares.
	squares := container.New(layout.NewGridLayoutWithRows(gui.ColumnWidth), row0, row1, row2, row3, row4, row5, row6, row7, row8)

	// Now configure the content to contain the top toolbar and the squares.
	content := container.NewBorder(toolbar, nil, nil, nil, squares)

	// Start threads for each sequence.
	go sequence.PlaySequence(*sequences[0], 0, this.Pad, eventsForLauchpad, guiButtons, dmxController, fixturesConfig, this.SequenceChannels, this.SoundTriggers)
	go sequence.PlaySequence(*sequences[1], 1, this.Pad, eventsForLauchpad, guiButtons, dmxController, fixturesConfig, this.SequenceChannels, this.SoundTriggers)
	go sequence.PlaySequence(*sequences[2], 2, this.Pad, eventsForLauchpad, guiButtons, dmxController, fixturesConfig, this.SequenceChannels, this.SoundTriggers)
	go sequence.PlaySequence(*sequences[3], 3, this.Pad, eventsForLauchpad, guiButtons, dmxController, fixturesConfig, this.SequenceChannels, this.SoundTriggers)

	// Light the first sequence as the default selected.
	this.SelectedSequence = 0
	buttons.InitButtons(&this, sequences, eventsForLauchpad, guiButtons, dmxController, fixturesConfig, commandChannels)

	// Clear the pad.
	buttons.AllFixturesOff(eventsForLauchpad, guiButtons, dmxController, fixturesConfig)

	// Listen to launchpad buttons.
	go func(guiButtons chan common.ALight,
		this *buttons.CurrentState,
		sequences []*common.Sequence,
		eventsForLauchpad chan common.ALight,
		dmxController *ft232.DMXController,
		fixturesConfig *fixture.Fixtures,
		commandChannels []chan common.Command,
		replyChannels []chan common.Sequence,
		updateChannels []chan common.Sequence) {

		buttons.ReadLaunchPadButtons(guiButtons, this, sequences, eventsForLauchpad, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels)

	}(guiButtons, &this, sequences, eventsForLauchpad, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels)

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
