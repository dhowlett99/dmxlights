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
			panel.UpdateButtonColor(alight, GuiFlashButtons)
		}
	}()
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
				modal, err := editor.NewFixturesPanel(sequences, myWindow, Y, X, fixturesConfig, commandChannels)
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

// MakeToolbar generates a tool bar at the top of the main window.
func MakeToolbar(myWindow fyne.Window, soundConfig *sound.SoundConfig,
	guiButtons chan common.ALight, eventsForLaunchPad chan common.ALight,
	config *usbdmx.ControllerConfig, launchPadName string, fixturesConfig *fixture.Fixtures) *widget.Toolbar {

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.FolderOpenIcon(), func() {
			fileOpener := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
				if err == nil && reader != nil {
					fmt.Printf("Filename %s\n", reader.URI())
					filename := filepath.Base(reader.URI().String())
					fixturesConfig, err = fixture.LoadFixtures(filename)
					if err != nil {
						fmt.Printf("dmxlights: error failed to load fixtures: %s\n", err.Error())
						PopupErrorMessage(myWindow, "error failed to load fixture file "+filename)
						return
					} else {
						myWindow.SetTitle("DMX Lights - Project Name :" + filename)
					}
				}
			}, myWindow)
			pwd, _ := os.Getwd()
			currentmfolder, _ := filepath.Abs(pwd)
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
		}),

		widget.NewToolbarAction(theme.FileIcon(), func() {
			dialog.ShowFileSave(func(writer fyne.URIWriteCloser, err error) {
				if err == nil && writer != nil {
					fixturesConfig, err = fixture.SaveFixturesWriter(writer)
					if err != nil {
						fmt.Printf("dmxlights: error failed to save fixtures: %s\n", err.Error())
					}
				}
			}, myWindow)
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.SettingsIcon(), func() {
			modal := runSettingsPopUp(myWindow, soundConfig, guiButtons, eventsForLaunchPad, config, launchPadName)
			modal.Resize(fyne.NewSize(250, 250))
			modal.Show()
		}),
	)
	return toolbar
}

func runSettingsPopUp(w fyne.Window, soundConfig *sound.SoundConfig,
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
