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
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/dhowlett99/dmxlights/pkg/buttons"
	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
	"github.com/oliread/usbdmx/ft232"
)

const ColumnWidth int = 9

//LightBlue := color.NRGBA{R: 0, G: 196, B: 255, A: 255}
//Red := color.NRGBA{R: 255, G: 0, B: 0, A: 255}
//Orange := color.NRGBA{R: 255, G: 111, B: 0, A: 255}
//Yellow := color.NRGBA{R: 255, G: 255, B: 0, A: 255}
//Green := color.NRGBA{R: 0, G: 255, B: 0, A: 255}
//Blue := color.NRGBA{R: 0, G: 0, B: 255, A: 255}
//Purple := color.NRGBA{R: 100, G: 0, B: 255, A: 255}
//Pink := color.NRGBA{R: 255, G: 0, B: 255, A: 255}
//Cyan := color.NRGBA{R: 0, G: 255, B: 255, A: 255}

type Button struct {
	button    *widget.Button
	rectangle *canvas.Rectangle
	container *fyne.Container
}

type MyPanel struct {
	Buttons [][]Button
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

func (panel *MyPanel) LabelButtons() {
	// panel.updateButtonColor(0, 0, Pink)
	// panel.updateButtonColor(1, 0, Red)
	// panel.updateButtonColor(2, 0, Green)
	// panel.updateButtonColor(3, 0, Blue)
	// panel.updateButtonColor(4, 0, LightBlue)
	// panel.updateButtonColor(5, 0, LightBlue)
	// panel.updateButtonColor(6, 0, LightBlue)
	// panel.updateButtonColor(7, 0, LightBlue)
	// panel.updateButtonColor(8, 0, Cyan)

	panel.UpdateButtonLabel(8, 1, "  >  ")
	panel.UpdateButtonLabel(8, 2, "  >  ")
	panel.UpdateButtonLabel(8, 3, "  >  ")

	panel.UpdateButtonLabel(0, 0, "CLEAR")
	panel.UpdateButtonLabel(1, 0, "RED")
	panel.UpdateButtonLabel(2, 0, "GREEN")
	panel.UpdateButtonLabel(3, 0, "BLUE")
	panel.UpdateButtonLabel(4, 0, "SENS -")
	panel.UpdateButtonLabel(5, 0, "SENS +")
	panel.UpdateButtonLabel(6, 0, "MAST -")
	panel.UpdateButtonLabel(7, 0, "MAST +")

	panel.UpdateButtonLabel(8, 4, "FLOOD")
	panel.UpdateButtonLabel(8, 5, "SAVE")
	panel.UpdateButtonLabel(8, 6, "START")
	panel.UpdateButtonLabel(8, 7, "STOP")

	panel.UpdateButtonLabel(0, 8, "SPEED-")
	panel.UpdateButtonLabel(1, 8, "SPEED+")
	panel.UpdateButtonLabel(2, 8, "SHIFT-")
	panel.UpdateButtonLabel(3, 8, "SHIFT+")
	panel.UpdateButtonLabel(4, 8, "SIZE-")
	panel.UpdateButtonLabel(5, 8, "SIZE+")
	panel.UpdateButtonLabel(6, 8, "FADE-")
	panel.UpdateButtonLabel(7, 8, "FADE+")
	panel.UpdateButtonLabel(8, 8, "BLACK")
}

func (panel *MyPanel) UpdateButtonColor(alight common.ALight, GuiFlashButtons [][]bool) {

	// Shortcut to label a button.
	if alight.UpdateLabel {
		panel.UpdateButtonLabel(alight.X, alight.Y, alight.Label)
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
		GuiFlashButtons[alight.X][alight.Y] = false

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

		// Let everyone know that we're flashing.
		GuiFlashButtons[alight.X][alight.Y] = true

		// We create a thread to flash the button.
		go func() {

			for {
				// Turn on.
				// Convert the launchpad code into RGB and the into NRGBA for the GUI.
				panel.Buttons[alight.X][alight.Y].rectangle.FillColor = convertRGBtoNRGBA(common.GetLaunchPadColorCodeByInt(alight.OnColor))
				panel.Buttons[alight.X][alight.Y].rectangle.Refresh()

				// Wake up every 10 ms to see if we need to stop.
				for t := 0; t < 300; t++ {
					if !GuiFlashButtons[alight.X][alight.Y] {
						panel.Buttons[alight.X][alight.Y].rectangle.FillColor = convertRGBtoNRGBA(common.GetLaunchPadColorCodeByInt(alight.OffColor))
						return
					}
					time.Sleep(1 * time.Millisecond)
				}

				// Turn off.
				// Convert the launchpad code into RGB and the into NRGBA for the GUI.
				panel.Buttons[alight.X][alight.Y].rectangle.FillColor = convertRGBtoNRGBA(common.GetLaunchPadColorCodeByInt(alight.OffColor))
				panel.Buttons[alight.X][alight.Y].rectangle.Refresh()

				// Wake up every 10 ms to see if we need to stop.
				for t := 0; t < 300; t++ {
					if !GuiFlashButtons[alight.X][alight.Y] {
						panel.Buttons[alight.X][alight.Y].rectangle.FillColor = convertRGBtoNRGBA(common.GetLaunchPadColorCodeByInt(alight.OffColor))
						return
					}
					time.Sleep(1 * time.Millisecond)
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

func (panel *MyPanel) SetButtonIcon(icon []byte, X int, Y int) {
	panel.Buttons[X][Y].button.Icon = fyne.NewStaticResource("", icon)
	size := panel.Buttons[X][Y].button.MinSize()
	fmt.Printf("size -> %+v\n", size)
	panel.Buttons[X][Y].button.Refresh()
}

func (panel *MyPanel) GetButtonColor(X int, Y int) color.Color {
	return panel.Buttons[X][Y].rectangle.FillColor
}

func (panel *MyPanel) GenerateRow(rowNumber int,
	sequences []*common.Sequence,
	this *buttons.CurrentState,
	eventsForLauchpad chan common.ALight,
	guiButtons chan common.ALight,
	dmxController *ft232.DMXController,
	fixturesConfig *fixture.Fixtures,
	commandChannels []chan common.Command,
	replyChannels []chan common.Sequence,
	updateChannels []chan common.Sequence) *fyne.Container {

	//White := color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	//Red := color.NRGBA{R: 255, G: 0, B: 0, A: 255}

	containers := []*fyne.Container{}
	for columnNumber := 0; columnNumber < ColumnWidth; columnNumber++ {
		button := Button{}
		Y := rowNumber
		X := columnNumber
		button.button = widget.NewButton("     ", func() {
			buttons.ProcessButtons(X, Y-1, sequences, this, eventsForLauchpad, guiButtons, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels)
		})
		button.rectangle = canvas.NewRectangle(color.Black)
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
