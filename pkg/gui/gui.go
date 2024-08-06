// Copyright (C) 2022,2025 dhowlett99.
// This is the dmxlights graphical user interface.
// Implemented and depends on fyne.io
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

package gui

import (
	"fmt"
	"image/color"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/dhowlett99/dmxlights/pkg/buttons"
	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/editor"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
	"github.com/dhowlett99/dmxlights/pkg/presets"
	"github.com/dhowlett99/dmxlights/pkg/sound"
	"github.com/oliread/usbdmx"
	"github.com/oliread/usbdmx/ft232"
)

const ColumnWidth int = 9

type Button struct {
	button    *noHoverButton
	rectangle *canvas.Rectangle
	container *fyne.Container
}

type MyPanel struct {
	Buttons          [][]Button
	SpeedLabel       *widget.Label
	ShiftLabel       *widget.Label
	SizeLabel        *widget.Label
	FadeLabel        *widget.Label
	VersionLabel     *widget.Button
	ColorDisplay     *fyne.Container
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

	// Create an empty color display.
	emptyColorDisplay := common.ColorDisplayControl{}
	panel.ColorDisplay = ShowColorDisplay(emptyColorDisplay)

	return panel
}

func (panel *MyPanel) LabelRightHandButtons() {
	panel.UpdateButtonLabel(8, 1, ">")
	panel.UpdateButtonLabel(8, 2, ">")
	panel.UpdateButtonLabel(8, 3, ">")

	panel.UpdateButtonLabel(8, 4, "FLOOD")
	panel.UpdateButtonLabel(8, 5, "SAVE")
	panel.UpdateButtonLabel(8, 6, "START.STOP")
	panel.UpdateButtonLabel(8, 7, "STROBE")

	panel.UpdateButtonLabel(8, 8, "BLACK.OUT")
}

func (panel *MyPanel) ListenAndSendToGUI(guiButtons chan common.ALight, GuiFlashButtons [][]common.ALight) {
	go func() {
		for {
			alight := <-guiButtons
			if !alight.ColorDisplay {
				panel.UpdateButtonColor(alight, GuiFlashButtons)
			}

			if alight.ColorDisplay {
				panel.UpdateColorDisplay(alight.ColorDisplayControl)
			}
		}
	}()
}

const Red = 0
const Orange = 1
const Yellow = 2
const Green = 3
const Cyan = 4
const Blue = 5
const Purple = 6
const Magenta = 7

func (panel *MyPanel) UpdateColorDisplay(control common.ColorDisplayControl) {

	// Clear all colors.
	for _, button := range panel.ColorDisplay.Objects {
		button.(*fyne.Container).Objects[0].(*canvas.Rectangle).Hidden = true
	}

	if control.Red {
		panel.ColorDisplay.Objects[Red].(*fyne.Container).Objects[0].(*canvas.Rectangle).Hidden = false
	}
	if control.Orange {
		panel.ColorDisplay.Objects[Orange].(*fyne.Container).Objects[0].(*canvas.Rectangle).Hidden = false
	}
	if control.Yellow {
		panel.ColorDisplay.Objects[Yellow].(*fyne.Container).Objects[0].(*canvas.Rectangle).Hidden = false
	}
	if control.Green {
		panel.ColorDisplay.Objects[Green].(*fyne.Container).Objects[0].(*canvas.Rectangle).Hidden = false
	}
	if control.Cyan {
		panel.ColorDisplay.Objects[Cyan].(*fyne.Container).Objects[0].(*canvas.Rectangle).Hidden = false
	}
	if control.Blue {
		panel.ColorDisplay.Objects[Blue].(*fyne.Container).Objects[0].(*canvas.Rectangle).Hidden = false
	}
	if control.Purple {
		panel.ColorDisplay.Objects[Purple].(*fyne.Container).Objects[0].(*canvas.Rectangle).Hidden = false
	}
	if control.Magenta {
		panel.ColorDisplay.Objects[Magenta].(*fyne.Container).Objects[0].(*canvas.Rectangle).Hidden = false
	}

	panel.ColorDisplay.Hidden = false
	panel.ColorDisplay.Refresh()
}

func (panel *MyPanel) UpdateButtonColor(alight common.ALight, GuiFlashButtons [][]common.ALight) {

	// Check for requests outside buttons avaialable.
	if alight.Button.X == -1 { // Addressing the top row.
		fmt.Printf("UpdateButtonColor: error X is -1\n")
		return
	}
	if alight.Button.Y == -1 { // Addressing the top row.
		fmt.Printf("UpdateButtonColor: error Y is -1\n")
		return
	}
	if alight.Button.X > 8 {
		fmt.Printf("UpdateButtonColor: error X is > 8 \n")
		return
	}
	if alight.Button.Y > 8 {
		fmt.Printf("UpdateButtonColor: error Y is > 8 \n")
		return
	}

	// Shortcut to label a button.
	if alight.UpdateLabel {
		panel.UpdateButtonLabel(alight.Button.X, alight.Button.Y, alight.Label)
		return
	}

	// Shortcut to label a status bar item.
	if alight.UpdateStatus {
		panel.UpdateStatusBar(alight.Status, alight.Hidden, alight.Which)
		return
	}

	// We're not flashing. Reset this button so it's not flashing.
	if !alight.Flash {

		// If the GuiFlashButtons array has a true value for this button,
		// Then there must be a thread flashing the lamp right now.
		// So we can assume its listening for a stop command.
		if GuiFlashButtons[alight.Button.X][alight.Button.Y].Flash {
			GuiFlashButtons[alight.Button.X][alight.Button.Y].FlashStopChannel <- true
			GuiFlashButtons[alight.Button.X][alight.Button.Y].Flash = false
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

		panel.Buttons[alight.Button.X][alight.Button.Y].rectangle.FillColor = color
		panel.Buttons[alight.Button.X][alight.Button.Y].rectangle.Refresh()
	} else {

		// Stop any existing flashing.
		if GuiFlashButtons[alight.Button.X][alight.Button.Y].Flash {
			GuiFlashButtons[alight.Button.X][alight.Button.Y].FlashStopChannel <- true
			GuiFlashButtons[alight.Button.X][alight.Button.Y].Flash = false
		}

		// Let everyone know that we're flashing.
		GuiFlashButtons[alight.Button.X][alight.Button.Y].Flash = true

		// We create a thread to flash the button.
		go func() {
			for {
				// Turn on.
				// Convert the  RGB color into NRGBA for the fyne.io GUI.
				panel.Buttons[alight.Button.X][alight.Button.Y].rectangle.FillColor = common.ConvertRGBtoNRGBA(alight.OnColor)
				panel.Buttons[alight.Button.X][alight.Button.Y].rectangle.Refresh()

				// We wait for a stop message or 250ms which ever comes first.
				select {
				case <-GuiFlashButtons[alight.Button.X][alight.Button.Y].FlashStopChannel:
					return
				case <-time.After(250 * time.Millisecond):
				}

				// Turn off.
				// Convert the  RGB color into NRGBA for the fyne.io GUI.
				panel.Buttons[alight.Button.X][alight.Button.Y].rectangle.FillColor = common.ConvertRGBtoNRGBA(alight.OffColor)
				panel.Buttons[alight.Button.X][alight.Button.Y].rectangle.Refresh()

				// We wait for a stop message or 250ms which ever comes first.
				select {
				case <-GuiFlashButtons[alight.Button.X][alight.Button.Y].FlashStopChannel:
					return
				case <-time.After(250 * time.Millisecond):
				}
			}
		}()
	}
}

func (panel *MyPanel) UpdateButtonLabel(X int, Y int, label string) {
	// If the label contains a space, replace it with a new line.
	label = strings.Replace(label, " ", "\n", 2)
	// If the label contains a period, replace it with a new line.
	label = strings.Replace(label, ".", "\n", 2)
	panel.Buttons[X][Y].button.Text = label
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
	if which == "version" {
		panel.VersionLabel.Hidden = hide
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

	iconImage, err := os.ReadFile(filename)
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
			container.NewHBox(dmxName, layout.NewSpacer(), dmxStatus),
			container.NewHBox(launchpadName, layout.NewSpacer(), launchpadStatus),
			widget.NewLabel(""),
			container.NewHBox(layout.NewSpacer(), layout.NewSpacer(), button),
		),
		myWindow.Canvas(),
	)

	modal.Show()

	return modal

}

// The latest version of the fyne.io toolkit implemnent a grey color
// which hovers over buttons when place the mouse over the button.
// Which stops us seeing the button colors.
// So we extend the button wiget to have a null MouseIn func.
type noHoverButton struct {
	widget.Button
}

// A null MouseIn func.
func (nhb *noHoverButton) MouseIn() {}

// A exteneded button functon with no hover affect.
func newNoHoverButton(label string, tapped func()) *noHoverButton {
	button := &noHoverButton{}
	button.ExtendBaseWidget(button)
	button.OnTapped = tapped
	button.SetText(label)
	return button
}

func (panel *MyPanel) GenerateRow(myWindow fyne.Window, rowNumber int,
	sequences []*common.Sequence,
	this *buttons.CurrentState,
	eventsForLauchpad chan common.ALight,
	guiButtons chan common.ALight,
	dmxController *ft232.DMXController,
	groupConfig *fixture.Groups,
	fixturesConfig *fixture.Fixtures,
	commandChannels []chan common.Command,
	replyChannels []chan common.Sequence,
	updateChannels []chan common.Sequence,
	dmxInterfacePresent bool) *fyne.Container {

	var popup *widget.PopUp

	containers := []*fyne.Container{}
	for columnNumber := 0; columnNumber < ColumnWidth; columnNumber++ {
		button := Button{}
		Y := rowNumber
		X := columnNumber

		var skipPopup bool
		button.button = newNoHoverButton("     ", func() {
			if X == 8 && Y == 5 || X > 7 || Y < 5 {
				skipPopup = true
			}
			if this.SavePreset {
				if !skipPopup {

					// Popup name
					title := widget.NewLabel("Save Presets")

					// Preset name.
					presetInput := widget.NewEntry()
					presetInput.Text = presets.GetPresetNumber(X, Y)
					presetLabel := widget.NewLabel("Preset Name")
					preset := container.NewAdaptiveGrid(3, presetLabel, presetInput, layout.NewSpacer())

					if this.PresetsStore[fmt.Sprint(X)+","+fmt.Sprint(Y-1)].Label != "" {
						presetInput.SetText(this.PresetsStore[fmt.Sprint(X)+","+fmt.Sprint(Y-1)].Label)
					}

					// Button color.
					buttonColorSelect := widget.NewSelect([]string{"White", "Red", "Orange", "Yellow", "Green", "Cyan", "Blue", "Purple", "Pink", "Black"}, func(value string) {})
					buttonColorSelect.Selected = this.PresetsStore[fmt.Sprint(X)+","+fmt.Sprint(Y-1)].ButtonColor

					buttonColorLabel := widget.NewLabel("Button Color")
					buttonColor := container.NewAdaptiveGrid(3, buttonColorLabel, buttonColorSelect, layout.NewSpacer())

					// Save button.
					buttonSave := widget.NewButton("OK", func() {})

					// Cancel button.
					buttonCancel := widget.NewButton("Cancel", func() {
						popup.Hide()
					})

					decideContents := container.NewHBox(layout.NewSpacer(), buttonCancel, buttonSave)
					decide := container.NewAdaptiveGrid(3, layout.NewSpacer(), layout.NewSpacer(), decideContents)

					form := container.NewVBox(title, preset, buttonColor, decide)

					// Layout of settings panel.
					popup = widget.NewModalPopUp(
						form,
						myWindow.Canvas(),
					)

					// Setup OK buttons action.
					buttonSave.OnTapped = func() {
						popup.Hide()
						if presetInput.Text == "" { // We clicked cancel so give up labelling.
							return
						}
						this.PresetsStore[fmt.Sprint(X)+","+fmt.Sprint(Y-1)] = presets.Preset{Label: presetInput.Text, State: true, Selected: true, ButtonColor: buttonColorSelect.Selected}
						presets.SavePresets(this.PresetsStore)
						presets.RefreshPresets(eventsForLauchpad, guiButtons, this.PresetsStore)
					}
					popup.Show()
				}
			}
			this.GUI = true
			buttons.ProcessButtons(X, Y-1, sequences, this, eventsForLauchpad, guiButtons, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels)

			skipPopup = false
		})
		if X == 8 && Y == 0 {
			button := widget.NewButton("DMXLIGI", func() {
				modal, err := editor.NewFixturePanel(sequences, myWindow, groupConfig, fixturesConfig, commandChannels)
				if err != nil {
					fmt.Printf("config not found for Group %d and Fixture %d  - %s\n", Y, X, err)
					return
				}
				modal.Resize(fyne.NewSize(800, 600))
				modal.Show()

			}) // button widget
			myLogo := canvas.NewImageFromFile("dmxlights.png")
			container1 := container.NewStack(
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
			button.container = container.NewStack(button.rectangle, button.button)
			containers = append(containers, button.container)
		}
		button.button.Importance = widget.LowImportance
		NewButton := Button{}
		NewButton.button = button.button
		NewButton.container = button.container
		NewButton.rectangle = button.rectangle
		panel.Buttons[columnNumber][rowNumber] = NewButton

	}

	row0 := container.New(layout.NewHBoxLayout(), containers[0], containers[1], containers[2], containers[3], containers[4], containers[5], containers[6], containers[7], containers[8])

	return row0
}

func NewFixtureEditor(sequences []*common.Sequence, myWindow fyne.Window, groupConfig *fixture.Groups, fixturesConfig *fixture.Fixtures, commandChannels []chan common.Command) error {

	modal, err := editor.NewFixturePanel(sequences, myWindow, groupConfig, fixturesConfig, commandChannels)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	modal.Resize(fyne.NewSize(800, 600))
	modal.Show()
	return nil
}

func PopupErrorMessage(myWindow fyne.Window, errorMessage string) {
	// Create a dialog for error messages.
	popupErrorPanel := &widget.PopUp{}
	// Ok button.
	button := widget.NewButton("OK", func() {
		popupErrorPanel.Hide()
	})

	popupErrorPanel = widget.NewModalPopUp(
		container.NewVBox(
			widget.NewLabel("Title"),
			widget.NewLabel("Error Message"),
			container.NewHBox(layout.NewSpacer(), button),
		),
		myWindow.Canvas(),
	)

	popupErrorPanel.Content.(*fyne.Container).Objects[0].(*widget.Label).Text = "Erorr Message"
	popupErrorPanel.Content.(*fyne.Container).Objects[1].(*widget.Label).Text = errorMessage
	popupErrorPanel.Show()
}

func AreYouSureDialog(myWindow fyne.Window, message string) *widget.PopUp {

	// Create a dialog for error messages.
	popupAreYouSurePanel := &widget.PopUp{}

	// Ok button.
	buttonOK := widget.NewButton("Quit Without Saving Project", func() {
		popupAreYouSurePanel.Hide()
		os.Exit(0)
	})

	buttonCancel := widget.NewButton("Cancel", func() {
		popupAreYouSurePanel.Hide()
	})

	popupAreYouSurePanel = widget.NewModalPopUp(
		container.NewVBox(
			widget.NewLabel("Project Has Changed"),
			widget.NewLabel("If you quit with out saving"),
			widget.NewLabel("your changes will be lost"),
			widget.NewLabel(message),
			container.NewHBox(buttonCancel, buttonOK),
		),
		myWindow.Canvas(),
	)

	return popupAreYouSurePanel
}

// MakeToolbar generates a tool bar at the top of the main window.
func MakeToolbar(myWindow fyne.Window, soundConfig *sound.SoundConfig,
	guiButtons chan common.ALight, eventsForLaunchPad chan common.ALight, commandChannels []chan common.Command,
	config *usbdmx.ControllerConfig, launchPadName string, fixturesConfig *fixture.Fixtures, startConfig *fixture.Fixtures) *widget.Toolbar {

	// Project open.
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.FolderOpenIcon(), func() {
			FileOpen(myWindow, startConfig, fixturesConfig, commandChannels)
		}),

		// Project save.
		widget.NewToolbarAction(theme.FileIcon(), func() {
			FileSave(myWindow, startConfig, fixturesConfig, commandChannels)
		}),

		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.SettingsIcon(), func() {
			modal := RunSettingsPopUp(myWindow, soundConfig, guiButtons, eventsForLaunchPad, config, launchPadName)
			modal.Resize(fyne.NewSize(250, 250))
			modal.Show()
		}),
	)
	return toolbar
}

func FileOpen(myWindow fyne.Window, startConfig *fixture.Fixtures, fixturesConfig *fixture.Fixtures, commandChannels []chan common.Command) {
	fileOpener := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
		if err == nil && reader != nil {
			newFixturesConfig, err := fixture.LoadFixturesReader(reader)
			if err != nil {
				fmt.Printf("dmxlights: error failed to load fixtures: %s\n", err.Error())
				PopupErrorMessage(myWindow, err.Error())
				return
			} else {
				// Reset the startConfig.
				startConfig.Fixtures = []fixture.Fixture{}
				startConfig.Fixtures = append(startConfig.Fixtures, newFixturesConfig.Fixtures...)
				filename := filepath.Base(reader.URI().String())
				myWindow.SetTitle("DMX Lights:" + filename)

				// Copy the newFixtures into the old pointer to the fixtures config.
				fixturesConfig.Fixtures = newFixturesConfig.Fixtures

				// Stop all the sequences.
				cmd := common.Command{
					Action: common.Reset,
				}
				common.SendCommandToAllSequence(cmd, commandChannels)
				// Update the fixtures config in all the sequences.
				cmd = common.Command{
					Action: common.UpdateFixturesConfig,
					Args: []common.Arg{
						{Name: "FixturesConfig", Value: fixturesConfig},
					},
				}
				common.SendCommandToAllSequence(cmd, commandChannels)
			}
		}
	}, myWindow)
	pwd, _ := os.Getwd()
	currentmfolder, _ := filepath.Abs(pwd + "/projects")
	if currentmfolder != "" {
		mfileURI := storage.NewFileURI(currentmfolder)
		mfileLister, _ := storage.ListerForURI(mfileURI)
		fileOpener.SetLocation(mfileLister)
		fileOpener.SetFilter(&storage.ExtensionFileFilter{
			Extensions: []string{
				".yaml",
			},
		})
	}
	fileOpener.Show()
}

func FileSave(myWindow fyne.Window, startConfig *fixture.Fixtures, fixturesConfig *fixture.Fixtures, commandChannels []chan common.Command) {
	fileSaver := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
		if err == nil && writer != nil {
			err = fixture.SaveFixturesWriter(writer, fixturesConfig)
			if err != nil {
				fmt.Printf("dmxlights: error failed to save fixtures: %s\n", err.Error())
			}
			// Reset the startConfig.
			startConfig.Fixtures = []fixture.Fixture{}
			startConfig.Fixtures = append(startConfig.Fixtures, fixturesConfig.Fixtures...)
		}
	}, myWindow)
	pwd, _ := os.Getwd()
	currentmfolder, _ := filepath.Abs(pwd + "/projects")
	if currentmfolder != "" {
		mfileURI := storage.NewFileURI(currentmfolder)
		mfileLister, _ := storage.ListerForURI(mfileURI)

		fileSaver.SetLocation(mfileLister)
		fileSaver.SetFilter(&storage.ExtensionFileFilter{
			Extensions: []string{
				".yaml",
			},
		})
		filename := strings.Split(myWindow.Title(), ":")
		fileSaver.SetFileName(filename[1])
	}
	fileSaver.Show()

}

func RunSettingsPopUp(w fyne.Window, soundConfig *sound.SoundConfig,
	guiButtons chan common.ALight, eventsForLaunchPad chan common.ALight, config *usbdmx.ControllerConfig, launchPadName string) (modal *widget.PopUp) {

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
		soundConfig.StartSoundConfig(selectedInput, guiButtons, eventsForLaunchPad)
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

func ShowColorDisplay(control common.ColorDisplayControl) *fyne.Container {

	red := Button{}
	red.rectangle = canvas.NewRectangle(color.RGBA{R: 255, G: 0, B: 0, A: 255})
	red.rectangle.SetMinSize(fyne.Size{Height: 5, Width: 5})
	red.rectangle.StrokeColor = color.Black
	red.rectangle.StrokeWidth = 1
	red.rectangle.Hidden = true
	red.container = container.NewStack(red.rectangle)

	orange := Button{}
	orange.rectangle = canvas.NewRectangle(color.RGBA{R: 255, G: 111, B: 0, A: 255})
	orange.rectangle.SetMinSize(fyne.Size{Height: 5, Width: 5})
	orange.rectangle.StrokeColor = color.Black
	orange.rectangle.StrokeWidth = 1
	orange.rectangle.Hidden = true
	orange.container = container.NewStack(orange.rectangle)

	yellow := Button{}
	yellow.rectangle = canvas.NewRectangle(color.RGBA{R: 255, G: 255, B: 0, A: 255})
	yellow.rectangle.SetMinSize(fyne.Size{Height: 5, Width: 5})
	yellow.rectangle.StrokeColor = color.Black
	yellow.rectangle.StrokeWidth = 1
	yellow.rectangle.Hidden = true
	yellow.container = container.NewStack(yellow.rectangle)

	green := Button{}
	green.rectangle = canvas.NewRectangle(color.RGBA{R: 0, G: 255, B: 0, A: 255})
	green.rectangle.Hidden = true
	green.rectangle.SetMinSize(fyne.Size{Height: 5, Width: 5})
	green.rectangle.StrokeColor = color.Black
	green.rectangle.StrokeWidth = 1
	green.rectangle.Hidden = true
	green.container = container.NewStack(green.rectangle)

	cyan := Button{}
	cyan.rectangle = canvas.NewRectangle(color.RGBA{R: 0, G: 255, B: 255, A: 255})
	cyan.rectangle.SetMinSize(fyne.Size{Height: 5, Width: 5})
	cyan.rectangle.StrokeColor = color.Black
	cyan.rectangle.StrokeWidth = 1
	cyan.rectangle.Hidden = true
	cyan.container = container.NewStack(cyan.rectangle)

	blue := Button{}
	blue.rectangle = canvas.NewRectangle(color.RGBA{R: 0, G: 0, B: 255, A: 255})
	blue.rectangle.SetMinSize(fyne.Size{Height: 5, Width: 5})
	blue.rectangle.StrokeColor = color.Black
	blue.rectangle.StrokeWidth = 1
	blue.rectangle.Hidden = true
	blue.container = container.NewStack(blue.rectangle)

	purple := Button{}
	purple.rectangle = canvas.NewRectangle(color.RGBA{R: 171, G: 0, B: 255, A: 255})
	purple.rectangle.SetMinSize(fyne.Size{Height: 5, Width: 5})
	purple.rectangle.StrokeColor = color.Black
	purple.rectangle.StrokeWidth = 1
	purple.rectangle.Hidden = true
	purple.container = container.NewStack(purple.rectangle)

	pink := Button{}
	pink.rectangle = canvas.NewRectangle(color.RGBA{R: 255, G: 0, B: 255, A: 255})
	pink.rectangle.SetMinSize(fyne.Size{Height: 5, Width: 5})
	pink.rectangle.StrokeColor = color.Black
	pink.rectangle.StrokeWidth = 1
	pink.rectangle.Hidden = true
	pink.container = container.NewStack(pink.rectangle)

	crimsom := Button{}
	crimsom.rectangle = canvas.NewRectangle(color.RGBA{R: 220, G: 20, B: 60, A: 255})
	crimsom.rectangle.SetMinSize(fyne.Size{Height: 5, Width: 5})
	crimsom.rectangle.StrokeColor = color.Black
	crimsom.rectangle.StrokeWidth = 1
	crimsom.rectangle.Hidden = true
	crimsom.container = container.NewStack(crimsom.rectangle)

	darkOrange := Button{}
	darkOrange.rectangle = canvas.NewRectangle(color.RGBA{R: 215, G: 50, B: 0, A: 255})
	darkOrange.rectangle.SetMinSize(fyne.Size{Height: 5, Width: 5})
	darkOrange.rectangle.StrokeColor = color.Black
	darkOrange.rectangle.StrokeWidth = 1
	darkOrange.rectangle.Hidden = true
	darkOrange.container = container.NewStack(darkOrange.rectangle)

	gold := Button{}
	gold.rectangle = canvas.NewRectangle(color.RGBA{R: 255, G: 215, B: 0, A: 255})
	gold.rectangle.SetMinSize(fyne.Size{Height: 5, Width: 5})
	gold.rectangle.StrokeColor = color.Black
	gold.rectangle.StrokeWidth = 1
	gold.rectangle.Hidden = true
	gold.container = container.NewStack(gold.rectangle)

	forestGreen := Button{}
	forestGreen.rectangle = canvas.NewRectangle(color.RGBA{R: 0, G: 100, B: 0, A: 255})
	forestGreen.rectangle.SetMinSize(fyne.Size{Height: 5, Width: 5})
	forestGreen.rectangle.StrokeColor = color.Black
	forestGreen.rectangle.StrokeWidth = 1
	forestGreen.rectangle.Hidden = true
	forestGreen.container = container.NewStack(forestGreen.rectangle)

	aqua := Button{}
	aqua.rectangle = canvas.NewRectangle(color.RGBA{R: 127, G: 255, B: 212, A: 255})
	aqua.rectangle.SetMinSize(fyne.Size{Height: 5, Width: 5})
	aqua.rectangle.StrokeColor = color.Black
	aqua.rectangle.StrokeWidth = 1
	aqua.rectangle.Hidden = true
	aqua.container = container.NewStack(aqua.rectangle)

	skyBlue := Button{}
	skyBlue.rectangle = canvas.NewRectangle(color.RGBA{R: 0, G: 191, B: 255, A: 255})
	skyBlue.rectangle.SetMinSize(fyne.Size{Height: 5, Width: 5})
	skyBlue.rectangle.StrokeColor = color.Black
	skyBlue.rectangle.StrokeWidth = 1
	skyBlue.rectangle.Hidden = true
	skyBlue.container = container.NewStack(skyBlue.rectangle)

	darkPurple := Button{}
	darkPurple.rectangle = canvas.NewRectangle(color.RGBA{R: 50, G: 0, B: 255, A: 255})
	darkPurple.rectangle.SetMinSize(fyne.Size{Height: 5, Width: 5})
	darkPurple.rectangle.StrokeColor = color.Black
	darkPurple.rectangle.StrokeWidth = 1
	darkPurple.rectangle.Hidden = true
	darkPurple.container = container.NewStack(darkPurple.rectangle)

	salmon := Button{}
	salmon.rectangle = canvas.NewRectangle(color.RGBA{R: 250, G: 128, B: 114, A: 255})
	salmon.rectangle.SetMinSize(fyne.Size{Height: 5, Width: 5})
	salmon.rectangle.StrokeColor = color.Black
	salmon.rectangle.StrokeWidth = 1
	salmon.rectangle.Hidden = true
	salmon.container = container.NewStack(salmon.rectangle)

	lightOrange := Button{}
	lightOrange.rectangle = canvas.NewRectangle(color.RGBA{R: 255, G: 175, B: 0, A: 255})
	lightOrange.rectangle.SetMinSize(fyne.Size{Height: 5, Width: 5})
	lightOrange.rectangle.StrokeColor = color.Black
	lightOrange.rectangle.StrokeWidth = 1
	lightOrange.rectangle.Hidden = true
	lightOrange.container = container.NewStack(lightOrange.rectangle)

	olive := Button{}
	olive.rectangle = canvas.NewRectangle(color.RGBA{R: 150, G: 150, B: 0, A: 255})
	olive.rectangle.SetMinSize(fyne.Size{Height: 5, Width: 5})
	olive.rectangle.StrokeColor = color.Black
	olive.rectangle.StrokeWidth = 1
	olive.rectangle.Hidden = true
	olive.container = container.NewStack(olive.rectangle)

	lawnGreen := Button{}
	lawnGreen.rectangle = canvas.NewRectangle(color.RGBA{R: 124, G: 252, B: 0, A: 255})
	lawnGreen.rectangle.SetMinSize(fyne.Size{Height: 5, Width: 5})
	lawnGreen.rectangle.StrokeColor = color.Black
	lawnGreen.rectangle.StrokeWidth = 1
	lawnGreen.rectangle.Hidden = true
	lawnGreen.container = container.NewStack(lawnGreen.rectangle)

	teal := Button{}
	teal.rectangle = canvas.NewRectangle(color.RGBA{R: 0, G: 128, B: 128, A: 255})
	teal.rectangle.SetMinSize(fyne.Size{Height: 5, Width: 5})
	teal.rectangle.StrokeColor = color.Black
	teal.rectangle.StrokeWidth = 1
	teal.rectangle.Hidden = true
	teal.container = container.NewStack(teal.rectangle)

	lightBlue := Button{}
	lightBlue.rectangle = canvas.NewRectangle(color.RGBA{R: 100, G: 185, B: 255, A: 255})
	lightBlue.rectangle.SetMinSize(fyne.Size{Height: 5, Width: 5})
	lightBlue.rectangle.StrokeColor = color.Black
	lightBlue.rectangle.StrokeWidth = 1
	lightBlue.rectangle.Hidden = true
	lightBlue.container = container.NewStack(lightBlue.rectangle)

	violet := Button{}
	violet.rectangle = canvas.NewRectangle(color.RGBA{R: 199, G: 21, B: 133, A: 255})
	violet.rectangle.SetMinSize(fyne.Size{Height: 5, Width: 5})
	violet.rectangle.StrokeColor = color.Black
	violet.rectangle.StrokeWidth = 1
	violet.rectangle.Hidden = true
	violet.container = container.NewStack(violet.rectangle)

	white := Button{}
	white.rectangle = canvas.NewRectangle(color.RGBA{R: 255, G: 255, B: 255, A: 255})
	white.rectangle.SetMinSize(fyne.Size{Height: 5, Width: 5})
	white.rectangle.StrokeColor = color.Black
	white.rectangle.StrokeWidth = 1
	white.rectangle.Hidden = true
	white.container = container.NewStack(white.rectangle)

	magenta := Button{}
	magenta.rectangle = canvas.NewRectangle(color.RGBA{R: 255, G: 0, B: 255, A: 255})
	magenta.rectangle.SetMinSize(fyne.Size{Height: 5, Width: 5})
	magenta.rectangle.StrokeColor = color.Black
	magenta.rectangle.StrokeWidth = 1
	magenta.rectangle.Hidden = true
	magenta.container = container.NewStack(magenta.rectangle)

	return container.New(
		layout.NewHBoxLayout(),

		// Top row.
		red.container,
		orange.container,
		yellow.container,
		green.container,
		cyan.container,
		blue.container,
		purple.container,
		magenta.container,

		// Middle row.
		crimsom.container,
		darkOrange.container,
		gold.container,
		forestGreen.container,
		aqua.container,
		skyBlue.container,
		darkPurple.container,
		pink.container,

		// Bottom row.
		salmon.container,
		lightOrange.container,
		olive.container,
		lawnGreen.container,
		teal.container,
		lightBlue.container,
		violet.container,
		white.container,
	)
}
