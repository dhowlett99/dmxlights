package editor

import (
	"image/color"
	"strings"

	"github.com/dhowlett99/dmxlights/pkg/common"
)

func SetFromCheckState(cp *ColorPanel) {

	var colorSelection []string

	// Clear Rectangles
	for x := 0; x < 10; x++ {
		cp.Rectanges[x].FillColor = color.White
	}

	// Now set the selected colors in the display.
	var count int
	for key, button := range cp.Buttons {
		if button.check.Checked {
			color, _ := common.GetRGBColorByName(key)
			cp.Rectanges[count].FillColor = common.ConvertRGBtoNRGBA(color)
			colorSelection = append(colorSelection, key)
			count++
		}
	}

	cp.ColorSelection = strings.Join(colorSelection, ",")
}

func SetRectangleColorsFromString(cp *ColorPanel, colors []string) {
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
	}

	var count int
	for _, c := range colors {
		if c != "" {
			RGBcolor, _ := common.GetRGBColorByName(c)
			cp.Rectanges[count].FillColor = common.ConvertRGBtoNRGBA(RGBcolor)
			cp.Buttons[c].check.Checked = true
		} else {
			cp.Rectanges[count].FillColor = color.White
		}
		count++
	}
}
