// Copyright (C) 2024 dhowlett99.
// This is the dmxlights colors definition package.
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
package colors

import "image/color"

var Grey = color.RGBA{R: 229, G: 228, B: 226, A: 255}
var Black = color.RGBA{R: 0, G: 0, B: 0, A: 255}
var Red = color.RGBA{R: 255, G: 0, B: 0, A: 255}
var Red5 = color.RGBA{R: 5, G: 0, B: 0, A: 255}
var Red25 = color.RGBA{R: 25, G: 0, B: 0, A: 255}
var Red50 = color.RGBA{R: 50, G: 0, B: 0, A: 255}
var Red75 = color.RGBA{R: 75, G: 0, B: 0, A: 255}
var Red100 = color.RGBA{R: 100, G: 0, B: 0, A: 255}
var Red125 = color.RGBA{R: 125, G: 0, B: 0, A: 255}
var Red150 = color.RGBA{R: 150, G: 0, B: 0, A: 255}
var Red175 = color.RGBA{R: 175, G: 0, B: 0, A: 255}
var Green = color.RGBA{R: 0, G: 255, B: 0, A: 255}
var Green1 = color.RGBA{R: 0, G: 1, B: 0, A: 255}
var Green50 = color.RGBA{R: 0, G: 50, B: 0, A: 255}
var QuarterGreen = color.RGBA{R: 0, G: 50, B: 0, A: 255}
var Blue = color.RGBA{R: 0, G: 0, B: 255, A: 255}
var QuarterBlue = color.RGBA{R: 0, G: 0, B: 50, A: 255}
var PresetYellow = color.RGBA{R: 150, G: 150, B: 0, A: 255}
var Cyan = color.RGBA{R: 0, G: 255, B: 255, A: 255}
var Yellow = color.RGBA{R: 255, G: 255, B: 0, A: 255}
var QuarterYellow = color.RGBA{R: 50, G: 50, B: 0, A: 255}
var Orange = color.RGBA{R: 255, G: 111, B: 0, A: 255}
var Magenta = color.RGBA{R: 255, G: 0, B: 255, A: 255}

var Crimson = color.RGBA{R: 220, G: 20, B: 60, A: 255}
var DarkOrange = color.RGBA{R: 215, G: 50, B: 0, A: 255}
var Gold = color.RGBA{R: 255, G: 215, B: 0, A: 255}
var ForestGreen = color.RGBA{R: 0, G: 100, B: 0, A: 255}
var Aqua = color.RGBA{R: 127, G: 255, B: 212, A: 255}
var SkyBlue = color.RGBA{R: 0, G: 191, B: 255, A: 255}
var Purple = color.RGBA{R: 100, G: 0, B: 255, A: 255}
var DarkPurple = color.RGBA{R: 50, G: 0, B: 255, A: 255}

var Pink = color.RGBA{R: 255, G: 192, B: 203, A: 255}
var Salmon = color.RGBA{R: 250, G: 128, B: 114, A: 255}
var LightOrange = color.RGBA{R: 255, G: 175, B: 0, A: 255}
var Olive = color.RGBA{R: 150, G: 150, B: 0, A: 255}
var LawnGreen = color.RGBA{R: 124, G: 252, B: 0, A: 255}
var Teal = color.RGBA{R: 0, G: 128, B: 128, A: 255}
var LightBlue = color.RGBA{R: 100, G: 185, B: 255, A: 255}
var Violet = color.RGBA{R: 199, G: 21, B: 133, A: 255}
var White = color.RGBA{R: 255, G: 255, B: 255, A: 255}
var EmptyColor = color.RGBA{}

func GetAvailableColors() []color.RGBA {
	var colors []color.RGBA
	colors = append(colors, Red)
	colors = append(colors, Orange)
	colors = append(colors, Yellow)
	colors = append(colors, Green)
	colors = append(colors, Cyan)
	colors = append(colors, Blue)
	colors = append(colors, Purple)
	colors = append(colors, Pink)
	colors = append(colors, White)
	colors = append(colors, Black)

	return colors
}

func FindColorIndexByName(colorLibrary []string, colorIn string) int {
	for colorNumber, color := range colorLibrary {
		if color == colorIn {
			return colorNumber
		}
	}
	return -1
}

func GetAvailableColorsAsStrings() []string {
	var colors []string
	colors = append(colors, "Red")
	colors = append(colors, "Orange")
	colors = append(colors, "Yellow")
	colors = append(colors, "Green")
	colors = append(colors, "Cyan")
	colors = append(colors, "Blue")
	colors = append(colors, "Purple")
	colors = append(colors, "Pink")
	colors = append(colors, "White")
	colors = append(colors, "Black")

	return colors
}

func GetColorFromIndexNumberFromColorsLibrary(selectedColor int, colorLibrary []color.RGBA) color.RGBA {
	return colorLibrary[selectedColor]
}
