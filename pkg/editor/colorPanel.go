// Copyright (C) 2022,2023 dhowlett99.
// This is the dmxlights fixture editor it is attached to a fixture and
// describes the fixtures properties which is then saved in the fixtures.yaml
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

package editor

import (
	"fmt"
	"image/color"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/dhowlett99/dmxlights/pkg/common"
)

type ColorPanel struct {
	Panel            *fyne.Container
	UpdateThisAction int
	UpdateColors     bool
	ColorSelection   string // Coma seperated string of color names, Upcase first letter.
	Buttons          map[string]Button
	ActionNumber     int
	Rectanges        []*canvas.Rectangle // Display rectangles.
	Modal            *widget.PopUp
}

type Button struct {
	rectangle *canvas.Rectangle
	check     *widget.Check
	container *fyne.Container
}

func NewColorPickerPanel(w fyne.Window) *ColorPanel {

	if debug {
		fmt.Printf("NewColorPickerPanel\n")
	}

	cp := ColorPanel{}
	cp.Buttons = make(map[string]Button, 10)

	red := Button{}
	red.rectangle = canvas.NewRectangle(color.RGBA{R: 255, G: 0, B: 0, A: 255})
	red.check = widget.NewCheck("", func(value bool) {})
	red.container = container.NewMax(red.rectangle, red.check)
	cp.Buttons["Red"] = red

	orange := Button{}
	orange.rectangle = canvas.NewRectangle(color.RGBA{R: 255, G: 111, B: 0, A: 255})
	orange.rectangle.SetMinSize(fyne.Size{Height: 20, Width: 20})
	orange.check = widget.NewCheck("", func(value bool) {})
	orange.container = container.NewMax(orange.rectangle, orange.check)
	cp.Buttons["Orange"] = orange

	yellow := Button{}
	yellow.rectangle = canvas.NewRectangle(color.RGBA{R: 255, G: 255, B: 0, A: 255})
	yellow.rectangle.SetMinSize(fyne.Size{Height: 20, Width: 20})
	yellow.check = widget.NewCheck("", func(value bool) {})
	yellow.container = container.NewMax(yellow.rectangle, yellow.check)
	cp.Buttons["Yellow"] = yellow

	green := Button{}
	green.rectangle = canvas.NewRectangle(color.RGBA{R: 0, G: 255, B: 0, A: 255})
	green.rectangle.SetMinSize(fyne.Size{Height: 20, Width: 20})
	green.check = widget.NewCheck("", func(value bool) {})
	green.container = container.NewMax(green.rectangle, green.check)
	cp.Buttons["Green"] = green

	cyan := Button{}
	cyan.rectangle = canvas.NewRectangle(color.RGBA{R: 0, G: 255, B: 255, A: 255})
	cyan.rectangle.SetMinSize(fyne.Size{Height: 20, Width: 20})
	cyan.check = widget.NewCheck("", func(value bool) {})
	cyan.container = container.NewMax(cyan.rectangle, cyan.check)
	cp.Buttons["Cyan"] = cyan

	blue := Button{}
	blue.rectangle = canvas.NewRectangle(color.RGBA{R: 0, G: 0, B: 255, A: 255})
	blue.rectangle.SetMinSize(fyne.Size{Height: 20, Width: 20})
	blue.check = widget.NewCheck("", func(value bool) {})
	blue.container = container.NewMax(blue.rectangle, blue.check)
	cp.Buttons["Blue"] = blue

	purple := Button{}
	purple.rectangle = canvas.NewRectangle(color.RGBA{R: 171, G: 0, B: 255, A: 255})
	purple.rectangle.SetMinSize(fyne.Size{Height: 20, Width: 20})
	purple.check = widget.NewCheck("", func(value bool) {})
	purple.container = container.NewMax(purple.rectangle, purple.check)
	cp.Buttons["Purple"] = purple

	pink := Button{}
	pink.rectangle = canvas.NewRectangle(color.RGBA{R: 255, G: 0, B: 255, A: 255})
	pink.rectangle.SetMinSize(fyne.Size{Height: 20, Width: 20})
	pink.check = widget.NewCheck("", func(value bool) {})
	pink.container = container.NewMax(pink.rectangle, pink.check)
	cp.Buttons["Pink"] = pink

	white := Button{}
	white.rectangle = canvas.NewRectangle(color.RGBA{R: 255, G: 255, B: 255, A: 255})
	white.rectangle.SetMinSize(fyne.Size{Height: 20, Width: 20})
	white.check = widget.NewCheck("", func(value bool) {})
	white.container = container.NewMax(white.rectangle, white.check)
	cp.Buttons["White"] = white

	black := Button{}
	black.rectangle = canvas.NewRectangle(color.RGBA{R: 0, G: 0, B: 0, A: 255})
	black.rectangle.SetMinSize(fyne.Size{Height: 20, Width: 20})
	black.check = widget.NewCheck("", func(value bool) {})
	black.container = container.NewMax(black.rectangle, black.check)
	cp.Buttons["Black"] = black

	colors := []string{}
	ok := widget.NewButton("OK", func() {
		// Now tell the Actions panel to update
		cp.ColorSelection = strings.Join(colors, ",")
		cp.UpdateColors = true
		cp.UpdateThisAction = cp.ActionNumber

		// Set the Rectangle display and update the color selection string.
		SetFromCheckState(&cp)
		cp.Modal.Hide()
	})

	cp.Panel = container.NewAdaptiveGrid(3,
		red.container,
		orange.container,
		yellow.container,
		green.container,
		cyan.container,
		blue.container,
		purple.container,
		pink.container,
		white.container,
		black.container,
		layout.NewSpacer(),
		ok,
	)

	return &cp
}

func SetFromCheckState(cp *ColorPanel) {

	if debug {
		fmt.Printf("SetFromCheckState\n")
	}

	var colorSelection []string

	// Clear Rectangles
	for x := 0; x < 10; x++ {
		cp.Rectanges[x].FillColor = color.White
		cp.Rectanges[x].StrokeColor = color.White
		cp.Rectanges[x].StrokeWidth = 1
	}

	// Now set the selected colors in the display.
	var count int

	// Enforce which order colors come back from the map in.
	labels := []string{"Red", "Orange", "Yellow", "Green", "Cyan", "Blue", "Purple", "Pink", "White", "Black"}
	for key := 0; key < len(labels); key++ {
		button := cp.Buttons[labels[key]]
		if button.check.Checked {
			currentColor, _ := common.GetRGBColorByName(labels[key])
			cp.Rectanges[count].FillColor = common.ConvertRGBtoNRGBA(currentColor)
			cp.Rectanges[count].StrokeColor = color.Black
			cp.Rectanges[count].StrokeWidth = 1
			colorSelection = append(colorSelection, labels[key])
			count++
		}
	}

	cp.ColorSelection = strings.Join(colorSelection, ",")
}

func SetRectangleColorsFromString(cp *ColorPanel, colors []string) {

	if debug {
		fmt.Printf("SetRectangleColorsFromString\n")
	}

	cp.Buttons["Red"].check.Checked = false
	cp.Buttons["Orange"].check.Checked = false
	cp.Buttons["Yellow"].check.Checked = false
	cp.Buttons["Green"].check.Checked = false
	cp.Buttons["Cyan"].check.Checked = false
	cp.Buttons["Blue"].check.Checked = false
	cp.Buttons["Purple"].check.Checked = false
	cp.Buttons["Pink"].check.Checked = false
	cp.Buttons["White"].check.Checked = false
	cp.Buttons["Black"].check.Checked = false

	// Clear Rectangles
	for x := 0; x < 10; x++ {
		cp.Rectanges[x].FillColor = color.White
		cp.Rectanges[x].StrokeColor = color.White
		cp.Rectanges[x].StrokeWidth = 1
	}

	var count int
	for _, c := range colors {
		if c != "" && c != "Off" {
			RGBcolor, _ := common.GetRGBColorByName(c)
			cp.Rectanges[count].FillColor = common.ConvertRGBtoNRGBA(RGBcolor)
			cp.Rectanges[count].StrokeColor = color.Black
			cp.Rectanges[count].StrokeWidth = 1
			cp.Buttons[c].check.Checked = true
		} else {
			cp.Rectanges[count].FillColor = color.White
			cp.Rectanges[count].StrokeColor = color.White
			cp.Rectanges[count].StrokeWidth = 1
		}
		count++
	}
}
