package editor

import (
	"github.com/dhowlett99/dmxlights/pkg/common"
)

func SetRectangleColors(cp *ColorPanel, colors []string) {
	// Clear the color display boxes.
	for x := 0; x < 10; x++ {
		RGBColor, _ := common.GetRGBColorByName("White")
		cp.Rectanges[x].FillColor = common.ConvertRGBtoNRGBA(RGBColor)
	}
	// Now set the selected colors in the display.
	var count int
	for _, color := range colors {
		RGBColor, _ := common.GetRGBColorByName(color)
		cp.Rectanges[count].FillColor = common.ConvertRGBtoNRGBA(RGBColor)
		count++
	}
}
