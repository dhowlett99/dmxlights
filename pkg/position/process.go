// Copyright (C) 2022, 2023 dhowlett99.
// This is the process a color part of dmxlight's position calculator,
// processing a color involves taking a color and fading it up, on and down
// at the correct time.
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

package position

import (
	"fmt"
	"math"

	"github.com/dhowlett99/dmxlights/pkg/common"
)

// processColor takes this color and next color and adds a fade color to the fadeColors map.
// This function uses simple rules to decide which fade value to add.
func processColor(stepNumber int, start bool, end bool, bounce bool, invert bool, fadeColors map[int][]common.FixtureBuffer, fixture common.Fixture, thisColor common.Color, lastColor common.Color, nextColor common.Color, sequence common.Sequence, shift int, scanner bool) map[int][]common.FixtureBuffer {

	// RULE #1 - If color is same as last time , play that color out again.
	if thisColor == lastColor {
		if debug {
			fmt.Printf("\t\tRULE#1 - fixture %d If color is same as last time , play that color out again. start %t end %t bounce %t invert %t\n", fixture.Number, start, end, bounce, invert)
		}
		fadeColors = makeAColor(stepNumber, 1, "SameColr1", shift, fadeColors, thisColor, sequence, fixture)
		return fadeColors
	}

	// RULE 2 - If color is different from last color and not black.
	if thisColor != lastColor && thisColor != common.Black {

		if debug {
			fmt.Printf("\t\tRULE#2 -fixture %d If color is different from last color and not black. start %t end %t bounce %t invert %t\n", fixture.Number, start, end, bounce, invert)
		}

		// Fade down last color but only if last color wasn't a black and we're not at the start.
		if lastColor != common.Black && !start && !end {
			fadeColors = fadeDownColor(stepNumber, 2, "FadeDwn1", shift, fadeColors, lastColor, sequence, fixture)
		}

		if !invert {
			fadeColors = fadeUpColor(stepNumber, 2, "FadeUp_1", shift, fadeColors, thisColor, sequence, fixture)
		}

		// Fade the color up.
		if !start && invert {
			fadeColors = fadeUpColor(stepNumber, 2, "FadeUp_1", shift, fadeColors, thisColor, sequence, fixture)
		}

		if invert { // Make a color
			fadeColors = makeAColor(stepNumber, 3, "MakeCLR1", shift, fadeColors, nextColor, sequence, fixture)
		}

		// If the next color is black. Fade dowm this color down ready.
		if nextColor == common.Black && !start && end {
			fadeColors = fadeDownColor(stepNumber, 2, "FadeDwn3", shift, fadeColors, thisColor, sequence, fixture)
		}

		// If the next color is another color but not black. Fade dowm this color down ready.
		if nextColor != common.Black && thisColor != nextColor && !start {
			fadeColors = fadeDownColor(stepNumber, 2, "FadeDwn4", shift, fadeColors, thisColor, sequence, fixture)
		}

		// If next color is not black and we're at the end or bouncing and we're the first fixture.
		if nextColor != common.Black && !start && end && bounce && fixture.Number == 0 {
			fadeColors = fadeDownColor(stepNumber, 2, "FadeDwn4", shift, fadeColors, nextColor, sequence, fixture)
		}

		return fadeColors
	}

	// RULE #3 - If this color is different from last color and is a black and at the start.
	if thisColor != lastColor && thisColor == common.Black && start {

		if debug {
			fmt.Printf("\t\tRULE#3 -fixture %d If this color is different from last colar and a is a black and at the start.. start %t end %t bounce %t invert %t\n", fixture.Number, start, end, bounce, invert)
		}

		if start && invert {
			//Fade down last color, so this black can be displayed.
			fadeColors = fadeDownColor(stepNumber, 3, "FadeDwn5", shift, fadeColors, lastColor, sequence, fixture)
		} else {
			fadeColors = makeAColor(stepNumber, 3, "MakeCLR1", shift, fadeColors, thisColor, sequence, fixture)
		}

		return fadeColors
	}

	// RULE #4 - If color is different from last color and color is a black.
	if thisColor != lastColor && thisColor == common.Black && !start {

		if debug {
			fmt.Printf("\t\tRULE#4 fixture %d If color is different from last color and color is a black. start %t end %t bounce %t invert %t\n", fixture.Number, start, end, bounce, invert)
		}

		// Fade down last color, so this black can be displayed.
		fadeColors = fadeDownColor(stepNumber, 4, "FadeDwn6", shift, fadeColors, lastColor, sequence, fixture)

		// Now that we have faded down. Populate one up,on & down cycle with the black we asked for.
		if !invert {
			fadeColors = makeAColor(stepNumber, 3, "makeABlack3", shift, fadeColors, thisColor, sequence, fixture)
		}

		// If the next color is color fade back up.
		if nextColor != common.Black && end && invert {
			fadeColors = fadeUpColor(stepNumber, 2, "FadeUp_3", shift, fadeColors, nextColor, sequence, fixture)
		}

		return fadeColors
	}

	if debug {
		fmt.Printf("\t\tNO RULE FIRED - fixture %d No rule fired. start %t end %t bounce %t invert %t\n", fixture.Number, start, end, bounce, invert)
	}

	return fadeColors
}

// fadeDownColor fades down the given color using the sequences fade down and fade off values.
func fadeDownColor(stepNumber int, rule int, debugMsg string, shift int, fadeColors map[int][]common.FixtureBuffer, color common.Color, sequence common.Sequence, fixture common.Fixture) map[int][]common.FixtureBuffer {

	if debug {
		fmt.Printf("\t\t\t\tfixture:%d fadeDownColor color %+v\n", fixture.Number, color)
	}

	var fade []int
	fade = append(fade, sequence.FadeDown...)
	fade = append(fade, sequence.FadeOff...)

	if color.R > 0 || color.G > 0 || color.B > 0 {
		for _, slope := range fade {
			newColor := addColor(stepNumber, rule, debugMsg, fixture, color, slope, sequence.ScannerChaser)
			fadeColors[fixture.Number] = append(fadeColors[fixture.Number], newColor)
		}
	} else {
		var shiftCounter int
		for _, slope := range fade {
			if shiftCounter == shift {
				break
			}
			newColor := addColor(stepNumber, rule, debugMsg, fixture, color, slope, sequence.ScannerChaser)
			fadeColors[fixture.Number] = append(fadeColors[fixture.Number], newColor)
			shiftCounter++
		}
	}

	return fadeColors
}

// fadeUpColor fades up the given color using the sequences fade up and fade on values.
func fadeUpColor(stepNumber int, rule int, debugMsg string, shift int, fadeColors map[int][]common.FixtureBuffer, color common.Color, sequence common.Sequence, fixture common.Fixture) map[int][]common.FixtureBuffer {

	if debug {
		fmt.Printf("\t\t\t\tfixture:%d fadeUpColor color %+v\n", fixture.Number, color)
	}

	var fade []int
	fade = append(fade, sequence.FadeUp...)
	fade = append(fade, sequence.FadeOn...)

	if color.R > 0 || color.G > 0 || color.B > 0 {
		for _, slope := range fade {
			newColor := addColor(stepNumber, rule, debugMsg, fixture, color, slope, sequence.ScannerChaser)
			fadeColors[fixture.Number] = append(fadeColors[fixture.Number], newColor)
		}
	} else {
		var shiftCounter int
		for _, slope := range fade {
			if shiftCounter == shift {
				break
			}
			newColor := addColor(stepNumber, rule, debugMsg, fixture, color, slope, sequence.ScannerChaser)
			fadeColors[fixture.Number] = append(fadeColors[fixture.Number], newColor)
			shiftCounter++
		}
	}

	return fadeColors
}

// makeAColor is used to add a color to the fixture buffer map of size fadeUp, fadeOn, FadeDown and fadeOff, which is the width of one cycle.
func makeAColor(stepNumber int, rule int, debugMsg string, shift int, fadeColors map[int][]common.FixtureBuffer, color common.Color, sequence common.Sequence, fixture common.Fixture) map[int][]common.FixtureBuffer {

	if debug {
		fmt.Printf("\t\t\t\tfixture:%d makeAColor color %+v\n", fixture.Number, color)
	}

	var fade []int
	fade = append(fade, sequence.FadeUp...)
	fade = append(fade, sequence.FadeOn...)
	fade = append(fade, sequence.FadeDown...)
	fade = append(fade, sequence.FadeOff...)

	var shiftCounter int
	for range fade {
		if shiftCounter == shift {
			break
		}
		newColor := addColor(stepNumber, rule, debugMsg, fixture, color, 255, sequence.ScannerChaser)
		fadeColors[fixture.Number] = append(fadeColors[fixture.Number], newColor)
		shiftCounter++
	}

	return fadeColors
}

// addColor adds a color to the fixtures buffer array, which is used later for assembling the positions.
func addColor(stepNumber int, rule int, debugMsg string, fixture common.Fixture, color common.Color, insertValue int, chase bool) common.FixtureBuffer {

	if debug {
		fmt.Printf("\t\t\t\t\tStep %d func=%s addColor fixture %d color %+v slope %d\n", stepNumber, debugMsg, fixture.Number, color, insertValue)
	}

	newColor := common.FixtureBuffer{}

	if debug {
		newColor.DebugMsg = debugMsg
		newColor.Step = stepNumber
		newColor.Rule = rule
	}

	newColor.Color = common.Color{}
	newColor.Gobo = fixture.Gobo
	newColor.Pan = fixture.Pan
	newColor.Tilt = fixture.Tilt
	newColor.Shutter = fixture.Shutter
	newColor.Color.R = int(math.Round((float64(color.R) / 100) * (float64(insertValue) / 2.55)))
	newColor.Color.G = int(math.Round((float64(color.G) / 100) * (float64(insertValue) / 2.55)))
	newColor.Color.B = int(math.Round((float64(color.B) / 100) * (float64(insertValue) / 2.55)))
	newColor.Color.W = int(math.Round((float64(color.W) / 100) * (float64(insertValue) / 2.55)))
	if !chase {
		newColor.Brightness = 255
	} else {
		newColor.Brightness = int(math.Round((float64(fixture.MasterDimmer) / 100) * (float64(insertValue) / 2.55)))
	}

	newColor.Enabled = fixture.Enabled
	newColor.MasterDimmer = fixture.MasterDimmer
	return newColor
}
