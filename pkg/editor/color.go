package editor

import (
	"image/color"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Button struct {
	button    *widget.Button
	rectangle *canvas.Rectangle
	check     *widget.Check
	container *fyne.Container
}

type Result struct {
	red    bool
	orange bool
	yellow bool
	green  bool
	blue   bool
	cyan   bool
	purple bool
	pink   bool
	white  bool
	black  bool
}

func NewColorPicker(w fyne.Window, cp *ColorPanel, actionNumber int) (modal *widget.PopUp) {

	result := Result{}

	redButton := Button{}
	redButton.rectangle = canvas.NewRectangle(color.RGBA{R: 255, G: 0, B: 0, A: 255})
	redButton.button = widget.NewButton("Red", func() {})
	size := fyne.Size{}
	size.Height = 20
	size.Width = 20
	redButton.rectangle.SetMinSize(size)
	redButton.check = widget.NewCheck("", func(value bool) {
		result.red = true
	})
	redButton.container = container.NewMax(redButton.rectangle, redButton.button, redButton.check)

	orangeButton := Button{}
	orangeButton.rectangle = canvas.NewRectangle(color.RGBA{R: 255, G: 111, B: 0, A: 255})
	orangeButton.button = widget.NewButton("Orange", func() {})
	size = fyne.Size{}
	size.Height = 20
	size.Width = 20
	orangeButton.rectangle.SetMinSize(size)
	orangeButton.check = widget.NewCheck("", func(value bool) {
		result.orange = true
	})
	orangeButton.container = container.NewMax(orangeButton.rectangle, orangeButton.button, orangeButton.check)

	yellowButton := Button{}
	yellowButton.rectangle = canvas.NewRectangle(color.RGBA{R: 255, G: 255, B: 0, A: 255})
	yellowButton.button = widget.NewButton("Yellow", func() {})
	size = fyne.Size{}
	size.Height = 20
	size.Width = 20
	yellowButton.rectangle.SetMinSize(size)
	yellowButton.check = widget.NewCheck("", func(value bool) {
		result.yellow = true
	})
	yellowButton.container = container.NewMax(yellowButton.rectangle, yellowButton.button, yellowButton.check)

	greenButton := Button{}
	greenButton.rectangle = canvas.NewRectangle(color.RGBA{R: 0, G: 255, B: 0, A: 255})
	greenButton.button = widget.NewButton("Green", func() {})
	size = fyne.Size{}
	size.Height = 20
	size.Width = 20
	greenButton.rectangle.SetMinSize(size)
	greenButton.check = widget.NewCheck("", func(value bool) {
		result.green = true
	})
	greenButton.container = container.NewMax(greenButton.rectangle, greenButton.button, greenButton.check)

	cyanButton := Button{}
	cyanButton.rectangle = canvas.NewRectangle(color.RGBA{R: 0, G: 255, B: 255, A: 255})
	cyanButton.button = widget.NewButton("Cyan", func() {})
	size = fyne.Size{}
	size.Height = 20
	size.Width = 20
	cyanButton.rectangle.SetMinSize(size)
	cyanButton.check = widget.NewCheck("", func(value bool) {
		result.cyan = true
	})
	cyanButton.container = container.NewMax(cyanButton.rectangle, cyanButton.button, cyanButton.check)

	blueButton := Button{}
	blueButton.rectangle = canvas.NewRectangle(color.RGBA{R: 0, G: 0, B: 255, A: 255})
	blueButton.button = widget.NewButton("Blue", func() {})
	size = fyne.Size{}
	size.Height = 20
	size.Width = 20
	blueButton.rectangle.SetMinSize(size)
	blueButton.check = widget.NewCheck("", func(value bool) {
		result.blue = true
	})
	blueButton.container = container.NewMax(blueButton.rectangle, blueButton.button, blueButton.check)

	purpleButton := Button{}
	purpleButton.rectangle = canvas.NewRectangle(color.RGBA{R: 171, G: 0, B: 255, A: 255})
	purpleButton.button = widget.NewButton("Purple", func() {})
	size = fyne.Size{}
	size.Height = 20
	size.Width = 20
	purpleButton.rectangle.SetMinSize(size)
	purpleButton.check = widget.NewCheck("", func(value bool) {
		result.purple = true
	})
	purpleButton.container = container.NewMax(purpleButton.rectangle, purpleButton.button, purpleButton.check)

	pinkButton := Button{}
	pinkButton.rectangle = canvas.NewRectangle(color.RGBA{R: 255, G: 0, B: 255, A: 255})
	pinkButton.button = widget.NewButton("Pink", func() {})
	size = fyne.Size{}
	size.Height = 20
	size.Width = 20
	pinkButton.rectangle.SetMinSize(size)
	pinkButton.check = widget.NewCheck("", func(value bool) {
		result.pink = true
	})
	pinkButton.container = container.NewMax(pinkButton.rectangle, pinkButton.button, pinkButton.check)

	whiteButton := Button{}
	whiteButton.rectangle = canvas.NewRectangle(color.RGBA{R: 255, G: 255, B: 255, A: 255})
	whiteButton.button = widget.NewButton("White", func() {})
	size = fyne.Size{}
	size.Height = 20
	size.Width = 20
	whiteButton.rectangle.SetMinSize(size)
	whiteButton.check = widget.NewCheck("", func(value bool) {
		result.white = true
	})
	whiteButton.container = container.NewMax(whiteButton.rectangle, whiteButton.button, whiteButton.check)

	blackButton := Button{}
	blackButton.rectangle = canvas.NewRectangle(color.RGBA{R: 0, G: 0, B: 0, A: 255})
	blackButton.button = widget.NewButton("Black", func() {})
	size = fyne.Size{}
	size.Height = 20
	size.Width = 20
	blackButton.rectangle.SetMinSize(size)
	blackButton.check = widget.NewCheck("", func(value bool) {
		result.black = true
	})
	blackButton.container = container.NewMax(blackButton.rectangle, blackButton.button, blackButton.check)

	colors := []string{}
	ok := widget.NewButton("OK", func() {

		if result.red {
			colors = append(colors, "Red")
		}
		if result.orange {
			colors = append(colors, "Orange")
		}
		if result.yellow {
			colors = append(colors, "Yellow")
		}
		if result.green {
			colors = append(colors, "Green")
		}
		if result.blue {
			colors = append(colors, "Blue")
		}
		if result.cyan {
			colors = append(colors, "Cyan")
		}
		if result.purple {
			colors = append(colors, "Purple")
		}
		if result.pink {
			colors = append(colors, "Pink")
		}
		if result.white {
			colors = append(colors, "White")
		}
		if result.black {
			colors = append(colors, "Black")
		}

		// Now tell the Actions panel to update
		cp.ColorSelection = strings.Join(colors, ",")
		cp.UpdateColors = true
		cp.UpdateThisAction = actionNumber

		SetRectangleColors(cp, colors)

		modal.Hide()
	})

	panel := container.NewAdaptiveGrid(3,
		redButton.container,
		orangeButton.container,
		yellowButton.container,
		greenButton.container,
		cyanButton.container,
		blueButton.container,
		purpleButton.container,
		pinkButton.container,
		whiteButton.container,
		blackButton.container,
		layout.NewSpacer(),
		ok,
	)

	// Layout of settings panel.
	modal = widget.NewModalPopUp(
		panel,
		w.Canvas(),
	)
	return modal
}
