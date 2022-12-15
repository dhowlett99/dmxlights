package editor

import (
	"github.com/dhowlett99/dmxlights/pkg/common"
)

func SetRectangleColors(cp *ColorPanel, colors []string) Result {

	result := Result{}

	// Clear the color display boxes.
	for x := 0; x < 10; x++ {
		RGBColor, _ := common.GetRGBColorByName("White")
		cp.Rectanges[x].FillColor = common.ConvertRGBtoNRGBA(RGBColor)
	}
	// Now set the selected colors in the display.
	var count int
	for _, color := range colors {
		if color == "Red" {
			result.red = true
		}
		if color == "Orange" {
			result.orange = true
		}
		if color == "Yellow" {
			result.yellow = true
		}
		if color == "Purple" {
			result.purple = true
		}
		if color == "Green" {
			result.green = true
		}
		if color == "Blue" {
			result.blue = true
		}
		if color == "Pink" {
			result.pink = true
		}
		if color == "Cyan" {
			result.cyan = true
		}
		if color == "White" {
			result.white = true
		}
		if color == "Black" {
			result.black = true
		}
		RGBColor, _ := common.GetRGBColorByName(color)
		cp.Rectanges[count].FillColor = common.ConvertRGBtoNRGBA(RGBColor)
		count++
	}

	return result
}
