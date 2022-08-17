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
	"github.com/oliread/usbdmx/ft232"
)

const ColumnWidth int = 9

type Button struct {
	button    *widget.Button
	rectangle *canvas.Rectangle
	container *fyne.Container
}

type MyPanel struct {
	Buttons    [][]Button
	SpeedLabel *widget.Label
	ShiftLabel *widget.Label
	SizeLabel  *widget.Label
	FadeLabel  *widget.Label
	BPMLabel   *widget.Label
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
		panel.UpdateStatusBar(alight.Status, alight.Which)
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

func (panel *MyPanel) UpdateStatusBar(label string, which string) {
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
	if which == "bpm" {
		panel.BPMLabel.SetText(label)
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

func (panel *MyPanel) GenerateRow(myWindow fyne.Window, rowNumber int,
	sequences []*common.Sequence,
	this *buttons.CurrentState,
	eventsForLauchpad chan common.ALight,
	guiButtons chan common.ALight,
	dmxController *ft232.DMXController,
	fixturesConfig *fixture.Fixtures,
	commandChannels []chan common.Command,
	replyChannels []chan common.Sequence,
	updateChannels []chan common.Sequence) *fyne.Container {

	containers := []*fyne.Container{}
	for columnNumber := 0; columnNumber < ColumnWidth; columnNumber++ {
		button := Button{}
		Y := rowNumber
		X := columnNumber

		var skipPopup bool
		button.button = widget.NewButton("     ", func() {

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
			button := widget.NewButton("MYDMX", nil) // button widget
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
func MakeToolbar(myLogo []byte) *widget.Toolbar {

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
	return toolbar
}
