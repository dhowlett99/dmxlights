package main

import (
	"bufio"
	"fmt"
	"image/color"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/dhowlett99/dmxlights/pkg/buttons"
	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/dmx"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
	"github.com/dhowlett99/dmxlights/pkg/patten"
	"github.com/dhowlett99/dmxlights/pkg/presets"
	"github.com/dhowlett99/dmxlights/pkg/sequence"
	"github.com/dhowlett99/dmxlights/pkg/sound"
	"github.com/oliread/usbdmx/ft232"
	"github.com/rakyll/launchpad/mk3"
)

type Button struct {
	button    *widget.Button
	rectangle *canvas.Rectangle
	container *fyne.Container
}

type MyPanel struct {
	Buttons [][]Button
}

func (panel *MyPanel) updateButtonColor(alight common.ALight, GuiFlashButtons [][]bool) {

	// Shortcut to label a button.
	if alight.UpdateLabel {
		panel.updateButtonLabel(alight.X, alight.Y, alight.Label)
		return
	}

	if alight.X == -1 { // Addressing the top row.
		fmt.Printf("error X is -1\n")
		return
	}
	if alight.Y == -1 { // Addressing the top row.
		fmt.Printf("error Y is -1\n")
		return
	}
	if alight.X > 8 {
		fmt.Printf("error X is > 8 \n")
		return
	}
	if alight.Y > 8 {
		fmt.Printf("error Y is > 8 \n")
		return
	}

	if !alight.Flash {
		// We're not flashing.
		// reset this button so it's not flashing.
		GuiFlashButtons[alight.X][alight.Y+1] = false

		// Take into account the brightness.
		Red := (float64(alight.Red) / 100) * (float64(alight.Brightness) / 2.55)
		Green := (float64(alight.Green) / 100) * (float64(alight.Brightness) / 2.55)
		Blue := (float64(alight.Blue) / 100) * (float64(alight.Brightness) / 2.55)

		color := color.NRGBA{}
		color.R = uint8(Red)
		color.G = uint8(Green)
		color.B = uint8(Blue)
		color.A = 255
		panel.Buttons[alight.X][alight.Y].rectangle.FillColor = color
		panel.Buttons[alight.X][alight.Y].rectangle.Refresh()
	} else {

		// We create a thread to flash the button.
		go func() {

			GuiFlashButtons[alight.X][alight.Y+1] = true

			for {
				if !GuiFlashButtons[alight.X][alight.Y+1] {
					return
				}
				// Turn on.
				// Convert the launchpad code into RGB and the into NRGBA for the GUI.
				panel.Buttons[alight.X][alight.Y].rectangle.FillColor = ConvertRGBtoNRGBA(common.GetLaunchPadColorCodeByInt(alight.OnColor))
				panel.Buttons[alight.X][alight.Y].rectangle.Refresh()

				// Wait 1/2 second.
				time.Sleep(500 * time.Millisecond)

				// Turn off.
				// Convert the launchpad code into RGB and the into NRGBA for the GUI.
				panel.Buttons[alight.X][alight.Y].rectangle.FillColor = ConvertRGBtoNRGBA(common.GetLaunchPadColorCodeByInt(alight.OffColor))
				panel.Buttons[alight.X][alight.Y].rectangle.Refresh()

				// Wait 1/2 second.
				time.Sleep(500 * time.Millisecond)
			}
		}()
	}
}

// Convert my common.Color RGB into color.NRGBA used by the fyne.io GUI library.
func ConvertRGBtoNRGBA(alight common.Color) color.NRGBA {
	NRGBAcolor := color.NRGBA{}
	NRGBAcolor.R = uint8(alight.R)
	NRGBAcolor.G = uint8(alight.G)
	NRGBAcolor.B = uint8(alight.B)
	NRGBAcolor.A = 255
	return NRGBAcolor
}

func (panel *MyPanel) updateButtonLabel(X int, Y int, label string) {
	panel.Buttons[X][Y].button.Text = label
	panel.Buttons[X][Y].button.Refresh()
}

func (panel *MyPanel) convertButtonImageToIcon(filename string) []byte {
	iconFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	r := bufio.NewReader(iconFile)

	iconImage, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	return iconImage
}

func (panel *MyPanel) setButtonIcon(icon []byte, X int, Y int) {
	panel.Buttons[X][Y].button.Icon = fyne.NewStaticResource("", icon)
	size := panel.Buttons[X][Y].button.MinSize()
	fmt.Printf("size -> %+v\n", size)
	panel.Buttons[X][Y].button.Refresh()
}

func (panel *MyPanel) getButtonColor(X int, Y int) color.Color {
	return panel.Buttons[X][Y].rectangle.FillColor
}

const columnWidth int = 9

func main() {

	// Start the GUI.
	fmt.Println("Start GUI")
	panel := MyPanel{}

	empty := Button{}

	panel.Buttons = [][]Button{
		{empty, empty, empty, empty, empty, empty, empty, empty, empty},
		{empty, empty, empty, empty, empty, empty, empty, empty, empty},
		{empty, empty, empty, empty, empty, empty, empty, empty, empty},
		{empty, empty, empty, empty, empty, empty, empty, empty, empty},
		{empty, empty, empty, empty, empty, empty, empty, empty, empty},
		{empty, empty, empty, empty, empty, empty, empty, empty, empty},
		{empty, empty, empty, empty, empty, empty, empty, empty, empty},
		{empty, empty, empty, empty, empty, empty, empty, empty, empty},
		{empty, empty, empty, empty, empty, empty, empty, empty, empty},
	}

	// Setup the current state
	this := buttons.CurrentState{}
	this.SequenceSpeed = 12
	this.ScannerSize = common.DefaultScannerSize
	this.SelectedShift = 0
	this.Blackout = false
	this.Flood = false
	this.FadeSpeed = 11         // Default start at 50ms.
	this.MasterBrightness = 255 // Affects all DMX fixtures and launchpad lamps.
	this.SoundGain = 0          // Fine gain -0.09 -> 0.09
	this.SelectedCordinates = 0 // Number of coordinates for scanner patterns is selected from 4 choices. 0=12, 1=26,2=24,3=32
	this.OffsetPan = 120        // Start pan from the center
	this.OffsetTilt = 120       // Start tilt from the center

	// Make an empty presets store.
	this.PresetsStore = make(map[string]bool)

	// Make a store for which sequences can be flood light.
	this.SelectedFloodMap = make(map[int]bool, 4)

	// Save the presets on exit.
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		presets.SavePresets(this.PresetsStore)
		os.Exit(1)
	}()

	fmt.Println("DMX Lighting")
	this.PresetsStore = presets.LoadPresets()

	// Setup DMX interface.
	fmt.Println("Setup DMX Interface")
	dmxController, err := dmx.NewDmXController()
	if err != nil {
		fmt.Printf("error initializing dmx interface: %v\n", err)
		os.Exit(1)
	}

	// Setup a connection to the common.
	// Tested with a Novation Launchpad mini mk3.
	fmt.Println("Setup Novation Launchpad")
	this.Pad, err = mk3.Open()
	if err != nil {
		log.Fatalf("error initializing launchpad: %v", err)
	}
	defer this.Pad.Close()

	// We need to be in programmers mode to use the common.
	this.Pad.Program()

	// Create a channel to send events to the GUI.
	guiButtons := make(chan common.ALight)

	// Make space for info on which GUI button is flashing.
	GuiFlashButtons := make([][]bool, 10)
	for i := 0; i < 10; i++ {
		GuiFlashButtons[i] = make([]bool, 10)
	}

	// Create a channel to send events to the launchthis.Pad.
	eventsForLauchpad := make(chan common.ALight)

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
		fmt.Printf("Found sequence: %s, desc: %s, type: %s\n", sequenceConf.Name, sequenceConf.Description, sequenceConf.Type)
		if sequenceConf.Type == "rgb" {
			this.SelectedFloodMap[index] = true // This sequence is this.Flood able because it's a rgb.
		}
		tempSequence := sequence.CreateSequence(sequenceConf.Type, index, this.Pattens, fixturesConfig, this.SequenceChannels, this.SelectedFloodMap)
		sequences = append(sequences, &tempSequence)
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

	// Initialize four select buttons.
	this.SelectButtonPressed = make([]bool, 4)

	// Initialize four function mode states.
	this.FunctionSelectMode = make([]bool, 4)

	// Initialize eight fixture states for the four sequences.
	this.DisabledFixture = make([][]bool, 9)
	for i := 0; i < 9; i++ {
		this.DisabledFixture[i] = make([]bool, 9)
	}

	// Initialize a ten length slice of empty slices for static lamps.
	this.StaticLamps = make([][]bool, 9)
	// Initialize those 10 empty function button slices
	for i := 0; i < 9; i++ {
		this.StaticLamps[i] = make([]bool, 9)
	}

	// Remember when we are in editing sequence colors mode.
	this.EditSequenceColorsMode = make([]bool, 4)

	// Remember when we are in setting scanner color mode.
	this.EditScannerColorsMode = make([]bool, 4)

	// Remember when we are in selecting gobo mode.
	this.EditGoboSelectionMode = make([]bool, 4)

	// Remember when we are in editing static colors mode.
	this.EditStaticColorsMode = make([]bool, 4)

	// Remember when we are in editing patten mode.
	this.EditPattenMode = make([]bool, 4)

	// Build the default set of Pattens.
	this.Pattens = patten.MakePatterns()

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
	this.StaticButtons = []common.StaticColorButton{}
	this.StaticButtons = append(this.StaticButtons, staticButton1)
	this.StaticButtons = append(this.StaticButtons, staticButton2)
	this.StaticButtons = append(this.StaticButtons, staticButton3)
	this.StaticButtons = append(this.StaticButtons, staticButton4)
	this.StaticButtons = append(this.StaticButtons, staticButton5)
	this.StaticButtons = append(this.StaticButtons, staticButton6)
	this.StaticButtons = append(this.StaticButtons, staticButton7)
	this.StaticButtons = append(this.StaticButtons, staticButton8)

	// this.SoundTriggers  is a an array of switches which control which sequence gets a music trigger.
	this.SoundTriggers = []*common.Trigger{}
	this.SoundTriggers = append(this.SoundTriggers, &common.Trigger{SequenceNumber: 0, State: false, Gain: this.SoundGain})
	this.SoundTriggers = append(this.SoundTriggers, &common.Trigger{SequenceNumber: 1, State: false, Gain: this.SoundGain})
	this.SoundTriggers = append(this.SoundTriggers, &common.Trigger{SequenceNumber: 2, State: false, Gain: this.SoundGain})
	this.SoundTriggers = append(this.SoundTriggers, &common.Trigger{SequenceNumber: 3, State: false, Gain: this.SoundGain})

	// Create a sound trigger object and give it the sequences so it can access their configs.
	sound.NewSoundTrigger(this.SoundTriggers, this.SequenceChannels)

	// Create a thread to handle GUI button events.
	go func(panel MyPanel, guiButtons chan common.ALight, GuiFlashButtons [][]bool) {
		for {
			alight := <-guiButtons
			panel.updateButtonColor(alight, GuiFlashButtons)
		}
	}(panel, guiButtons, GuiFlashButtons)

	// Now create a thread to handle launchpad events.
	go func() {
		common.ListenAndSendToLaunchPad(eventsForLauchpad, this.Pad)
	}()

	//LightBlue := color.NRGBA{R: 0, G: 196, B: 255, A: 255}
	//Red := color.NRGBA{R: 255, G: 0, B: 0, A: 255}
	//Orange := color.NRGBA{R: 255, G: 111, B: 0, A: 255}
	//Yellow := color.NRGBA{R: 255, G: 255, B: 0, A: 255}
	//Green := color.NRGBA{R: 0, G: 255, B: 0, A: 255}
	//Blue := color.NRGBA{R: 0, G: 0, B: 255, A: 255}
	//Purple := color.NRGBA{R: 100, G: 0, B: 255, A: 255}
	//Pink := color.NRGBA{R: 255, G: 0, B: 255, A: 255}
	//Cyan := color.NRGBA{R: 0, G: 255, B: 255, A: 255}

	myApp := app.New()
	myWindow := myApp.NewWindow("DMX Lights")
	myLogo := panel.convertButtonImageToIcon("dmxlights.png")
	myWindow.Resize(fyne.NewSize(400, 50))

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			log.Println("New document")
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentCutIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.ContentRedoIcon(), func() {}),
		widget.NewToolbarAction(fyne.NewStaticResource("icon", myLogo), func() {
			log.Println("Display help")
		}),
	)

	row0 := panel.generateRow(0, sequences, &this, eventsForLauchpad, guiButtons, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels)
	row1 := panel.generateRow(1, sequences, &this, eventsForLauchpad, guiButtons, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels)
	row2 := panel.generateRow(2, sequences, &this, eventsForLauchpad, guiButtons, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels)
	row3 := panel.generateRow(3, sequences, &this, eventsForLauchpad, guiButtons, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels)
	row4 := panel.generateRow(4, sequences, &this, eventsForLauchpad, guiButtons, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels)
	row5 := panel.generateRow(5, sequences, &this, eventsForLauchpad, guiButtons, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels)
	row6 := panel.generateRow(6, sequences, &this, eventsForLauchpad, guiButtons, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels)
	row7 := panel.generateRow(7, sequences, &this, eventsForLauchpad, guiButtons, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels)
	row8 := panel.generateRow(8, sequences, &this, eventsForLauchpad, guiButtons, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels)

	// panel.updateButtonColor(0, 0, Pink)
	// panel.updateButtonColor(1, 0, Red)
	// panel.updateButtonColor(2, 0, Green)
	// panel.updateButtonColor(3, 0, Blue)
	// panel.updateButtonColor(4, 0, LightBlue)
	// panel.updateButtonColor(5, 0, LightBlue)
	// panel.updateButtonColor(6, 0, LightBlue)
	// panel.updateButtonColor(7, 0, LightBlue)
	// panel.updateButtonColor(8, 0, Cyan)

	panel.updateButtonLabel(8, 1, "  >  ")
	panel.updateButtonLabel(8, 2, "  >  ")
	panel.updateButtonLabel(8, 3, "  >  ")

	panel.updateButtonLabel(0, 0, "CLEAR")
	panel.updateButtonLabel(1, 0, "RED")
	panel.updateButtonLabel(2, 0, "GREEN")
	panel.updateButtonLabel(3, 0, "BLUE")
	panel.updateButtonLabel(4, 0, "SENS -")
	panel.updateButtonLabel(5, 0, "SENS +")
	panel.updateButtonLabel(6, 0, "MAST -")
	panel.updateButtonLabel(7, 0, "MAST +")

	panel.updateButtonLabel(8, 4, "FLOOD")
	panel.updateButtonLabel(8, 5, "SAVE")
	panel.updateButtonLabel(8, 6, "START")
	panel.updateButtonLabel(8, 7, "STOP")

	panel.updateButtonLabel(0, 8, "SPEED-")
	panel.updateButtonLabel(1, 8, "SPEED+")
	panel.updateButtonLabel(2, 8, "SHIFT-")
	panel.updateButtonLabel(3, 8, "SHIFT+")
	panel.updateButtonLabel(4, 8, "SIZE-")
	panel.updateButtonLabel(5, 8, "SIZE+")
	panel.updateButtonLabel(6, 8, "FADE-")
	panel.updateButtonLabel(7, 8, "FADE+")
	panel.updateButtonLabel(8, 8, "BLACK")

	squares := container.New(layout.NewGridLayoutWithRows(columnWidth), row0, row1, row2, row3, row4, row5, row6, row7, row8)

	content := container.NewBorder(toolbar, nil, nil, nil, squares)

	// Start off by turning off all of the Lights
	this.Pad.Reset()

	// Start threads for each sequence.
	go sequence.PlaySequence(*sequences[0], 0, this.Pad, eventsForLauchpad, guiButtons, this.Pattens, dmxController, fixturesConfig, this.SequenceChannels, this.SoundTriggers)
	go sequence.PlaySequence(*sequences[1], 1, this.Pad, eventsForLauchpad, guiButtons, this.Pattens, dmxController, fixturesConfig, this.SequenceChannels, this.SoundTriggers)
	go sequence.PlaySequence(*sequences[2], 2, this.Pad, eventsForLauchpad, guiButtons, this.Pattens, dmxController, fixturesConfig, this.SequenceChannels, this.SoundTriggers)
	go sequence.PlaySequence(*sequences[3], 3, this.Pad, eventsForLauchpad, guiButtons, this.Pattens, dmxController, fixturesConfig, this.SequenceChannels, this.SoundTriggers)

	// Light up any existing presets.
	presets.InitPresets(eventsForLauchpad, guiButtons, this.PresetsStore)

	// Light the buttons at the bottom.
	common.ShowBottomButtons(*sequences[1], 7, eventsForLauchpad, guiButtons)

	// Light the logo blue.
	this.Pad.Light(8, -1, 0, 0, 255)

	// Light the clear button purple.
	common.LightLamp(common.ALight{X: 0, Y: -1, Brightness: 255, Red: 200, Green: 0, Blue: 255}, eventsForLauchpad, guiButtons)

	// Light the static color buttons.
	common.LightLamp(common.ALight{X: 1, Y: -1, Brightness: 255, Red: 255, Green: 0, Blue: 0}, eventsForLauchpad, guiButtons)
	common.LightLamp(common.ALight{X: 2, Y: -1, Brightness: 255, Red: 0, Green: 255, Blue: 0}, eventsForLauchpad, guiButtons)
	common.LightLamp(common.ALight{X: 3, Y: -1, Brightness: 255, Red: 0, Green: 0, Blue: 255}, eventsForLauchpad, guiButtons)

	// Light top functions.
	common.LightLamp(common.ALight{X: 4, Y: -1, Brightness: 255, Red: 3, Green: 255, Blue: 255}, eventsForLauchpad, guiButtons)
	common.LightLamp(common.ALight{X: 5, Y: -1, Brightness: 255, Red: 3, Green: 255, Blue: 255}, eventsForLauchpad, guiButtons)
	common.LightLamp(common.ALight{X: 6, Y: -1, Brightness: 255, Red: 3, Green: 255, Blue: 255}, eventsForLauchpad, guiButtons)
	common.LightLamp(common.ALight{X: 7, Y: -1, Brightness: 255, Red: 3, Green: 255, Blue: 255}, eventsForLauchpad, guiButtons)

	// Light the save, start, stop and this.Blackout buttons.
	common.LightLamp(common.ALight{X: 8, Y: 5, Brightness: 255, Red: 255, Green: 255, Blue: 255}, eventsForLauchpad, guiButtons)
	common.LightLamp(common.ALight{X: 8, Y: 6, Brightness: 255, Red: 255, Green: 255, Blue: 255}, eventsForLauchpad, guiButtons)
	common.LightLamp(common.ALight{X: 8, Y: 7, Brightness: 255, Red: 255, Green: 255, Blue: 255}, eventsForLauchpad, guiButtons)
	common.LightLamp(common.ALight{X: 8, Y: 8, Brightness: 255, Red: 255, Green: 255, Blue: 255}, eventsForLauchpad, guiButtons)

	// Initialise the this.Flood button to be green.
	common.LightLamp(common.ALight{X: 7, Y: 3, Brightness: 255, Red: 255, Green: 0, Blue: 0}, eventsForLauchpad, guiButtons)

	// Light the first sequence as the default selected.
	this.SelectedSequence = 0
	sequence.SequenceSelect(eventsForLauchpad, guiButtons, this.SelectedSequence)

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

func (panel *MyPanel) generateRow(rowNumber int,
	sequences []*common.Sequence,
	this *buttons.CurrentState,
	eventsForLauchpad chan common.ALight,
	guiButtons chan common.ALight,
	dmxController *ft232.DMXController,
	fixturesConfig *fixture.Fixtures,
	commandChannels []chan common.Command,
	replyChannels []chan common.Sequence,
	updateChannels []chan common.Sequence) *fyne.Container {

	White := color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	//Red := color.NRGBA{R: 255, G: 0, B: 0, A: 255}

	containers := []*fyne.Container{}
	for columnNumber := 0; columnNumber < columnWidth; columnNumber++ {
		button := Button{}
		Y := rowNumber
		X := columnNumber
		button.button = widget.NewButton("     ", func() {
			buttons.ProcessButtons(X, Y-1, sequences, this, eventsForLauchpad, guiButtons, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels)
		})
		button.rectangle = canvas.NewRectangle(White)
		size := fyne.Size{}
		size.Height = 80
		size.Width = 80
		button.rectangle.SetMinSize(size)
		button.container = container.NewMax(button.rectangle, button.button)
		containers = append(containers, button.container)
		NewButton := Button{}
		NewButton.button = button.button
		NewButton.container = button.container
		NewButton.rectangle = button.rectangle
		panel.Buttons[columnNumber][rowNumber] = NewButton
	}

	row0 := container.New(layout.NewHBoxLayout(), containers[0], containers[1], containers[2], containers[3], containers[4], containers[5], containers[6], containers[7], containers[8])

	return row0
}
