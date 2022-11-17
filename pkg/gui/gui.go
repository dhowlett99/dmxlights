package gui

import (
	"bufio"
	"fmt"
	"image/color"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/dhowlett99/dmxlights/pkg/buttons"
	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
	"github.com/dhowlett99/dmxlights/pkg/presets"
	"github.com/dhowlett99/dmxlights/pkg/sound"
	"github.com/oliread/usbdmx"
	"github.com/oliread/usbdmx/ft232"
)

const ColumnWidth int = 9

type Button struct {
	button    *widget.Button
	rectangle *canvas.Rectangle
	container *fyne.Container
}

type MyPanel struct {
	Buttons          [][]Button
	SpeedLabel       *widget.Label
	ShiftLabel       *widget.Label
	SizeLabel        *widget.Label
	FadeLabel        *widget.Label
	BeatLabel        *widget.Button
	TiltLabel        *widget.Label
	RedLabel         *widget.Label
	GreenLabel       *widget.Label
	BlueLabel        *widget.Label
	SensitivityLabel *widget.Label
	MasterLabel      *widget.Label
}

func NewPanel() MyPanel {

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

	return panel
}

func (panel *MyPanel) LabelRightHandButtons() {
	panel.UpdateButtonLabel(8, 1, "  >  ")
	panel.UpdateButtonLabel(8, 2, "  >  ")
	panel.UpdateButtonLabel(8, 3, "  >  ")

	panel.UpdateButtonLabel(8, 4, "FLOOD")
	panel.UpdateButtonLabel(8, 5, "SAVE")
	panel.UpdateButtonLabel(8, 6, "START.STOP")
	panel.UpdateButtonLabel(8, 7, "STROBE")

	panel.UpdateButtonLabel(8, 8, "BLACK.OUT")
}

func (panel *MyPanel) UpdateButtonColor(alight common.ALight, GuiFlashButtons [][]common.ALight) {

	// Shortcut to label a button.
	if alight.UpdateLabel {
		panel.UpdateButtonLabel(alight.X, alight.Y, alight.Label)
		return
	}

	// Shortcut to label a status bar item.
	if alight.UpdateStatus {
		panel.UpdateStatusBar(alight.Status, alight.Hidden, alight.Which)
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

	// We're not flashing. Reset this button so it's not flashing.
	if !alight.Flash {

		// If the GuiFlashButtons array has a true value for this button,
		// Then there must be a thread flashing the lamp right now.
		// So we can assume its listening for a stop command.
		if GuiFlashButtons[alight.X][alight.Y].Flash {
			GuiFlashButtons[alight.X][alight.Y].FlashStopChannel <- true
			GuiFlashButtons[alight.X][alight.Y].Flash = false
		}

		// Take into account the brightness.
		Red := (float64(alight.Red) / 100) * (float64(alight.Brightness) / 2.55)
		Green := (float64(alight.Green) / 100) * (float64(alight.Brightness) / 2.55)
		Blue := (float64(alight.Blue) / 100) * (float64(alight.Brightness) / 2.55)

		color := color.RGBA{}
		color.R = uint8(Red)
		color.G = uint8(Green)
		color.B = uint8(Blue)
		color.A = 255

		panel.Buttons[alight.X][alight.Y].rectangle.FillColor = color
		panel.Buttons[alight.X][alight.Y].rectangle.Refresh()
	} else {

		// Stop any existing flashing.
		if GuiFlashButtons[alight.X][alight.Y].Flash {
			GuiFlashButtons[alight.X][alight.Y].FlashStopChannel <- true
			GuiFlashButtons[alight.X][alight.Y].Flash = false
		}

		// Let everyone know that we're flashing.
		GuiFlashButtons[alight.X][alight.Y].Flash = true

		// We create a thread to flash the button.
		go func() {
			for {
				// Turn on.
				// Convert the  RGB color into NRGBA for the fyne.io GUI.
				panel.Buttons[alight.X][alight.Y].rectangle.FillColor = convertRGBtoNRGBA(alight.OnColor)
				panel.Buttons[alight.X][alight.Y].rectangle.Refresh()

				// We wait for a stop message or 250ms which ever comes first.
				select {
				case <-GuiFlashButtons[alight.X][alight.Y].FlashStopChannel:
					return
				case <-time.After(250 * time.Millisecond):
				}

				// Turn off.
				// Convert the  RGB color into NRGBA for the fyne.io GUI.
				panel.Buttons[alight.X][alight.Y].rectangle.FillColor = convertRGBtoNRGBA(alight.OffColor)
				panel.Buttons[alight.X][alight.Y].rectangle.Refresh()

				// We wait for a stop message or 250ms which ever comes first.
				select {
				case <-GuiFlashButtons[alight.X][alight.Y].FlashStopChannel:
					return
				case <-time.After(250 * time.Millisecond):
				}
			}
		}()
	}
}

// Convert my common.Color RGB into color.NRGBA used by the fyne.io GUI library.
func convertRGBtoNRGBA(alight common.Color) color.NRGBA {
	NRGBAcolor := color.NRGBA{}
	NRGBAcolor.R = uint8(alight.R)
	NRGBAcolor.G = uint8(alight.G)
	NRGBAcolor.B = uint8(alight.B)
	NRGBAcolor.A = 255
	return NRGBAcolor
}

func (panel *MyPanel) UpdateButtonLabel(X int, Y int, label string) {
	// If the label contains a period, replace it with a new line.
	panel.Buttons[X][Y].button.Text = strings.Replace(label, ".", "\n", 2)
	panel.Buttons[X][Y].button.Refresh()
}

func (panel *MyPanel) UpdateStatusBar(label string, hide bool, which string) {
	if which == "speed" {
		panel.SpeedLabel.SetText(label)
	}
	if which == "shift" {
		panel.ShiftLabel.SetText(label)
	}
	if which == "size" {
		panel.SizeLabel.SetText(label)
	}
	if which == "fade" {
		panel.FadeLabel.SetText(label)
	}
	if which == "beat" {
		panel.BeatLabel.Hidden = hide
	}
	if which == "tilt" {
		panel.TiltLabel.SetText(label)
	}
	if which == "red" {
		panel.RedLabel.SetText(label)
	}
	if which == "green" {
		panel.GreenLabel.SetText(label)
	}
	if which == "blue" || which == "pan" {
		panel.BlueLabel.SetText(label)
	}
	if which == "sensitivity" {
		panel.SensitivityLabel.SetText(label)
	}
	if which == "master" {
		panel.MasterLabel.SetText(label)
	}
}

func (panel *MyPanel) ConvertButtonImageToIcon(filename string) []byte {
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

func (panel *MyPanel) GetButtonColor(X int, Y int) color.Color {
	return panel.Buttons[X][Y].rectangle.FillColor
}

type Device struct {
	Name   string
	Status bool
}

func (panel *MyPanel) PopupNotFoundMessage(myWindow fyne.Window, dmxInterface Device, launchPad Device) (modal *widget.PopUp) {

	title := widget.NewLabel("Information")

	// Ok button.
	button := widget.NewButton("OK", func() {
		modal.Hide()
	})

	var dmxStatus *widget.Label
	dmxName := widget.NewLabel(dmxInterface.Name)
	if dmxInterface.Status {
		dmxStatus = widget.NewLabel("Connected")
	} else {
		dmxStatus = widget.NewLabel("Not Connected")
	}

	var launchpadStatus *widget.Label
	launchpadName := widget.NewLabel(launchPad.Name)
	if launchPad.Status {
		launchpadStatus = widget.NewLabel("Connected")
	} else {
		launchpadStatus = widget.NewLabel("Not Connected")
	}

	modal = widget.NewModalPopUp(
		container.NewVBox(
			title,
			container.NewHBox(dmxName, dmxStatus),
			container.NewHBox(launchpadName, launchpadStatus),
			widget.NewLabel(""),
			container.NewHBox(layout.NewSpacer(), button),
		),
		myWindow.Canvas(),
	)

	modal.Show()

	return modal

}

func (panel *MyPanel) GenerateRow(myWindow fyne.Window, rowNumber int,
	sequences []*common.Sequence,
	this *buttons.CurrentState,
	eventsForLauchpad chan common.ALight,
	guiButtons chan common.ALight,
	dmxController *ft232.DMXController,
	fixturesConfig *fixture.Fixtures,
	commandChannels []chan common.Command,
	replyChannels []chan common.Sequence,
	updateChannels []chan common.Sequence,
	dmxInterfacePresent bool) *fyne.Container {

	containers := []*fyne.Container{}
	for columnNumber := 0; columnNumber < ColumnWidth; columnNumber++ {
		button := Button{}
		Y := rowNumber
		X := columnNumber

		var skipPopup bool
		button.button = widget.NewButton("     ", func() {

			if this.Config {
				modal, err := runConfigPopUp(myWindow, Y, X, fixturesConfig)
				if err != nil {
					fmt.Printf("config not found for Group %d and Fixture %d  - %s\n", Y, X, err)
					return
				}
				modal.Resize(fyne.NewSize(500, 500))
				modal.Show()
			}

			if X == 8 && Y == 5 || X > 7 || Y < 5 {
				skipPopup = true
			}
			if this.SavePreset {
				if !skipPopup {
					items := []*widget.FormItem{}
					// Keep existing text.
					name := widget.NewEntry()
					if this.PresetsStore[fmt.Sprint(X)+","+fmt.Sprint(Y-1)].Label != "" {
						name.SetText(this.PresetsStore[fmt.Sprint(X)+","+fmt.Sprint(Y-1)].Label)
					}
					item := widget.NewFormItem("Name", name)
					items = append(items, item)
					popup := dialog.NewForm("Enter Preset", "Ok", "Cancel", items, func(bool) {
						if name.Text == "" { // We clicked cancel so give up labelling.
							return
						}
						this.PresetsStore[fmt.Sprint(X)+","+fmt.Sprint(Y-1)] = presets.Preset{Label: name.Text, State: true, Selected: true}
						presets.InitPresets(eventsForLauchpad, guiButtons, this.PresetsStore)
						presets.SavePresets(this.PresetsStore)
						Red := common.Color{R: 255, G: 0, B: 0}
						PresetYellow := common.Color{R: 150, G: 150, B: 0}
						common.FlashLight(X, Y-1, Red, PresetYellow, eventsForLauchpad, guiButtons)
					}, myWindow)
					popup.Show()
				}
			}
			buttons.ProcessButtons(X, Y-1, sequences, this, eventsForLauchpad, guiButtons, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels, true)
			skipPopup = false
		})
		if X == 8 && Y == 0 {
			button := widget.NewButton("MYDMX", func() {
				if this.Config {
					this.Config = false
				} else {
					this.Config = true
				}
			}) // button widget
			myLogo := canvas.NewImageFromFile("dmxlights.png")
			container1 := container.NewMax(
				button,
				myLogo,
			)
			containers = append(containers, container1)
		} else {
			button.rectangle = canvas.NewRectangle(color.White)
			size := fyne.Size{}
			size.Height = 80
			size.Width = 80
			button.rectangle.SetMinSize(size)
			button.container = container.NewMax(button.rectangle, button.button)
			containers = append(containers, button.container)
		}
		NewButton := Button{}
		NewButton.button = button.button
		NewButton.container = button.container
		NewButton.rectangle = button.rectangle
		panel.Buttons[columnNumber][rowNumber] = NewButton

	}

	row0 := container.New(layout.NewHBoxLayout(), containers[0], containers[1], containers[2], containers[3], containers[4], containers[5], containers[6], containers[7], containers[8])

	return row0
}

// MakeToolbar generates a tool bar at the top of the main window.
func MakeToolbar(myWindow fyne.Window, soundConfig *sound.SoundConfig,
	guiButtons chan common.ALight, config *usbdmx.ControllerConfig, launchPadName string) *widget.Toolbar {
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.SettingsIcon(), func() {
			modal := runSettingsPopUp(myWindow, soundConfig, guiButtons, config, launchPadName)
			modal.Resize(fyne.NewSize(250, 250))
			modal.Show()
		}),
	)
	return toolbar
}

func runSettingsPopUp(w fyne.Window, soundConfig *sound.SoundConfig,
	guiButtons chan common.ALight, config *usbdmx.ControllerConfig, launchPadName string) (modal *widget.PopUp) {

	selectedInput := soundConfig.GetDeviceName()

	title := widget.NewLabel("Settings")
	title.TextStyle = fyne.TextStyle{
		Bold: true,
	}

	launchpadLabel := widget.NewLabel("Midi Interface Installed")

	// Launchpad configuration.
	var launchPads []string
	launchPads = append(launchPads, launchPadName)
	launchpadSelect := widget.NewSelect(launchPads, func(value string) {
		selectedInput = launchPadName
	})
	launchpadSelect.PlaceHolder = launchPadName

	// DMX interface configuration.
	dmxInterfaceLabel := widget.NewLabel("DMX Interface Installed ")
	var dmxLabels []string
	dmxLabels = append(dmxLabels, "Not Found")
	if config != nil {
		dmxLabels[0] = fmt.Sprintf("FT323:%d", config.InputInterfaceID)
	}
	dmxInterfaceSelect := widget.NewSelect(dmxLabels, func(value string) {
		selectedInput = dmxLabels[0]
	})
	dmxInterfaceSelect.PlaceHolder = dmxLabels[0]

	// Audio interface configuration.
	audioInterfaceLabel := widget.NewLabel("Select Audio Input")
	audioInterfaceSelect := widget.NewSelect(soundConfig.GetSoundConfig(), func(value string) {
		selectedInput = value
	})
	audioInterfaceSelect.PlaceHolder = selectedInput

	// Ok button.
	button := widget.NewButton("OK", func() {
		modal.Hide()
		soundConfig.StopSoundConfig()
		soundConfig.StartSoundConfig(selectedInput, guiButtons)
	})

	// Layout of settings panel.
	modal = widget.NewModalPopUp(
		container.NewVBox(
			title,
			container.NewHBox(dmxInterfaceLabel, dmxInterfaceSelect),
			container.NewHBox(launchpadLabel, launchpadSelect),
			container.NewHBox(audioInterfaceLabel, audioInterfaceSelect),
			widget.NewLabel(""),
			container.NewHBox(layout.NewSpacer(), button),
		),
		w.Canvas(),
	)
	return modal
}

func runConfigPopUp(w fyne.Window, group int, number int, fixtures *fixture.Fixtures) (modal *widget.PopUp, err error) {

	fixture, err := fixture.GetFixureDetails(group, number, fixtures)
	if err != nil {
		return nil, fmt.Errorf("GetFixureDetails %s", err.Error())
	}

	title := widget.NewLabel(fmt.Sprintf("Edit Config for Sequence %d Fixture %d", fixture.Group, fixture.Number))
	title.TextStyle = fyne.TextStyle{
		Bold: true,
	}

	nameInput := widget.NewEntry()
	nameInput.SetPlaceHolder(fixture.Name)
	descInput := widget.NewEntry()
	descInput.SetPlaceHolder(fixture.Description)
	addrInput := widget.NewEntry()
	addrInput.SetPlaceHolder(fmt.Sprintf("%d", fixture.Address))

	var formTopItems []*widget.FormItem
	// Top Form.
	name1 := widget.NewEntry()
	name1.SetText(fixture.Name)
	formTopItem := widget.NewFormItem("Name", name1)
	formTopItems = append(formTopItems, formTopItem)

	name2 := widget.NewEntry()
	name2.SetText(fixture.Description)
	formTopItem2 := widget.NewFormItem("Description", name2)
	formTopItems = append(formTopItems, formTopItem2)

	name3 := widget.NewEntry()
	name3.SetText(fmt.Sprintf("%d", fixture.Address))
	formTopItem3 := widget.NewFormItem("DMX Address", name3)
	formTopItems = append(formTopItems, formTopItem3)

	formTop := &widget.Form{
		Items: formTopItems,
	}

	var formItems []*widget.FormItem

	if fixture.Type == "rgb" || fixture.Type == "scanner" {
		for _, channel := range fixture.Channels {
			name := widget.NewEntry()
			name.SetText(channel.Name)
			formItem := widget.NewFormItem(fmt.Sprintf("%d", channel.Number), name)
			formItems = append(formItems, formItem)
		}
	}

	if fixture.Type == "switch" {
		for _, state := range fixture.States {
			name := widget.NewEntry()
			name.SetText(state.Name)
			formItem := widget.NewFormItem(state.Label, name)
			formItems = append(formItems, formItem)
		}
	}

	// Channels Form.
	form := &widget.Form{
		Items: formItems,
	}
	scrollableForm := container.NewScroll(form)
	scrollableForm.SetMinSize(fyne.Size{Height: 400, Width: 200})

	// Save button.
	buttonSave := widget.NewButton("Save", func() {
		modal.Hide()
	})
	// Cancel button.
	buttonCancel := widget.NewButton("Cancel", func() {
		modal.Hide()
	})

	saveCancel := container.NewHBox(layout.NewSpacer(), buttonCancel, buttonSave)

	top := container.NewBorder(formTop, nil, nil, nil, scrollableForm)
	content := container.NewBorder(top, nil, nil, nil, saveCancel)

	// Layout of settings panel.
	modal = widget.NewModalPopUp(
		content,
		w.Canvas(),
	)

	return modal, nil
}
