package editor

import (
	"fmt"

	"github.com/dhowlett99/dmxlights/pkg/common"
)

func SetRectangleColorsFromCheckState(cp *ColorPanel) {

	// Clear the color display boxes.
	for x := 0; x < 10; x++ {
		RGBColor, _ := common.GetRGBColorByName("White")
		cp.Rectanges[x].FillColor = common.ConvertRGBtoNRGBA(RGBColor)
	}
	// Now set the selected colors in the display.
	var count int
	for key, button := range cp.Buttons {
		fmt.Printf("Setting Color %s\n", key)
		if button.check.Checked {
			// Light the first available display rectangle
			color, _ := common.GetRGBColorByName(key)
			cp.Rectanges[count].FillColor = common.ConvertRGBtoNRGBA(color)
			count++
		}
	}

}

func SetRectangleColorsFromString(cp *ColorPanel, colors []string) {
	// Now set the selected colors in the rectanges display.
	if len(colors) == 0 {
		// Clear the color display boxes.
		for x := 0; x < 10; x++ {
			RGBColor, _ := common.GetRGBColorByName("White")
			cp.Rectanges[x].FillColor = common.ConvertRGBtoNRGBA(RGBColor)
		}
		return
	}
	var count int
	for _, color := range colors {
		if color == "" {
			return
		}
		fmt.Printf("Setting Color FromString %s\n", color)
		RGBcolor, _ := common.GetRGBColorByName(color)
		cp.Rectanges[count].FillColor = common.ConvertRGBtoNRGBA(RGBcolor)
		cp.Buttons[color].check.Checked = true
		count++
	}
}
