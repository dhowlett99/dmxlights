package main

import (
	"bufio"
	"fmt"
	"image/color"
	"io/ioutil"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/dhowlett99/dmxlights/pkg/buttons"
	"github.com/dhowlett99/dmxlights/pkg/common"
)

type Button struct {
	button    *widget.Button
	rectangle *canvas.Rectangle
	container *fyne.Container
}

type MyPanel struct {
	Buttons [][]Button
}

func (panel *MyPanel) updateButtonColor(X int, Y int, color color.Color) {
	if X == -1 { // Addressing the top row.
		fmt.Printf("error X is -1\n")
		return
	}
	if Y == -1 { // Addressing the top row.
		fmt.Printf("error Y is -1\n")
		return
	}
	if X > 8 {
		fmt.Printf("error X is > 8 \n")
		return
	}
	if Y > 8 {
		fmt.Printf("error Y is > 8 \n")
		return
	}
	panel.Buttons[X][Y].rectangle.FillColor = color
	panel.Buttons[X][Y].rectangle.Refresh()
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

	LightBlue := color.NRGBA{R: 0, G: 196, B: 255, A: 255}
	Red := color.NRGBA{R: 255, G: 0, B: 0, A: 255}
	//Orange := color.NRGBA{R: 255, G: 111, B: 0, A: 255}
	//Yellow := color.NRGBA{R: 255, G: 255, B: 0, A: 255}
	Green := color.NRGBA{R: 0, G: 255, B: 0, A: 255}
	Blue := color.NRGBA{R: 0, G: 0, B: 255, A: 255}
	//Purple := color.NRGBA{R: 100, G: 0, B: 255, A: 255}
	Pink := color.NRGBA{R: 255, G: 0, B: 255, A: 255}
	Cyan := color.NRGBA{R: 0, G: 255, B: 255, A: 255}

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

	row0 := panel.generateRow(0)
	row1 := panel.generateRow(1)
	row2 := panel.generateRow(2)
	row3 := panel.generateRow(3)
	row4 := panel.generateRow(4)
	row5 := panel.generateRow(5)
	row6 := panel.generateRow(6)
	row7 := panel.generateRow(7)
	row8 := panel.generateRow(8)

	panel.updateButtonColor(0, 0, Pink)
	panel.updateButtonColor(1, 0, Red)
	panel.updateButtonColor(2, 0, Green)
	panel.updateButtonColor(3, 0, Blue)
	panel.updateButtonColor(4, 0, LightBlue)
	panel.updateButtonColor(5, 0, LightBlue)
	panel.updateButtonColor(6, 0, LightBlue)
	panel.updateButtonColor(7, 0, LightBlue)
	panel.updateButtonColor(8, 0, Cyan)

	panel.updateButtonLabel(8, 1, "  >  ")
	panel.updateButtonLabel(8, 2, "  >  ")
	panel.updateButtonLabel(8, 3, "  >  ")

	//panel.setButtonIcon(lamp, 1, 1)

	squares := container.New(layout.NewGridLayoutWithRows(columnWidth), row0, row1, row2, row3, row4, row5, row6, row7, row8)

	content := container.NewBorder(toolbar, nil, nil, nil, squares)

	guiButtons := make(chan common.ALight)

	// Start the sequences.
	go func(guiButtons chan common.ALight) {
		buttons.StartSequences(guiButtons)
	}(guiButtons)

	// Listen for GUI light a lamp messages.
	go func(panel MyPanel, guiButtons chan common.ALight) {
		for {
			button := <-guiButtons
			color := color.NRGBA{}
			color.R = uint8(button.Red)
			color.G = uint8(button.Green)
			color.B = uint8(button.Blue)
			color.A = 255
			panel.updateButtonColor(button.X, button.Y, color)
		}
	}(panel, guiButtons)

	myWindow.SetContent(content)

	myWindow.ShowAndRun()

}

func (panel *MyPanel) generateRow(rowNumber int) *fyne.Container {

	White := color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	Red := color.NRGBA{R: 255, G: 0, B: 0, A: 255}

	containers := []*fyne.Container{}
	for columnNumber := 0; columnNumber < columnWidth; columnNumber++ {
		button := Button{}
		Y := rowNumber
		X := columnNumber
		button.button = widget.NewButton("     ", func() {
			panel.updateButtonColor(X, Y, Red)
			fmt.Printf("tapped Y:%d X:%d\n", Y, X)
		})
		size := fyne.Size{}
		size.Height = 10
		size.Width = 10
		button.button.Resize(size)
		button.rectangle = canvas.NewRectangle(White)
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
