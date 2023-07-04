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

	"github.com/dhowlett99/dmxlights/pkg/common"
)

func processColor(stepNumber int, start bool, end bool, bounce bool, invert bool, fadeColors map[int][]common.FixtureBuffer, fixture common.Fixture, fixtureNumber int, thisColor common.Color, lastColor common.Color, nextColor common.Color, sequence common.Sequence, shift int, patternShift int, scanner bool) map[int][]common.FixtureBuffer {

	// RULE #1 - If color is same as last time , play that color out again.
	if thisColor == lastColor {
		if debug {
			fmt.Printf("\t\tRULE#1 - fixture %d If color is same as last time , play that color out again. start %t end %t bounce %t invert %t\n", fixtureNumber, start, end, bounce, invert)
		}
		fadeColors = keepSameAsLastTime(stepNumber, 1, "SameColr", shift, fadeColors, thisColor, sequence, fixture, fixtureNumber)
		return fadeColors
	}

	// RULE 2 - If color is different from last color and not black.
	if thisColor != lastColor && thisColor != common.Black {

		if debug {
			fmt.Printf("\t\tRULE#2 -fixture %d If color is different from last color and not black. start %t end %t bounce %t invert %t\n", fixtureNumber, start, end, bounce, invert)
		}

		// Fade down last color but only if last color wasn't a black and we're not at the start.
		if lastColor != common.Black && !start && !end {
			fadeColors = fadeDownColor(stepNumber, 2, "FadeDwn1", shift, fadeColors, lastColor, sequence, fixture, fixtureNumber)
		}

		// Fade the color up.
		if !invert {
			fadeColors = fadeUpColor(stepNumber, 2, "FadeUp_1", fadeColors, thisColor, sequence, fixture, fixtureNumber)
		}

		// Leave the color for on time.
		fadeColors = fadeOnColor(stepNumber, 2, "FadeUp_2", fadeColors, thisColor, sequence, fixture, fixtureNumber)

		if invert { // Make a color
			fadeColors = makeAColor(stepNumber, 3, "MakeCLR1", shift, fadeColors, thisColor, sequence, fixture, fixtureNumber)
		}

		// If the next color is black. Fade dowm this color down ready.
		if nextColor == common.Black && !start && end {
			fadeColors = fadeDownColor(stepNumber, 2, "FadeDwn3", shift, fadeColors, thisColor, sequence, fixture, fixtureNumber)
		}

		//fmt.Printf("---->>> fixture %d nextColor %+v start %t end %t bounce %t\n", fixtureNumber, nextColor, start, end, bounce)
		if nextColor != common.Black && !start && end && bounce && fixtureNumber == 0 {
			fadeColors = fadeDownColor(stepNumber, 2, "FadeDwn4", shift, fadeColors, nextColor, sequence, fixture, fixtureNumber)
		}

		return fadeColors
	}

	// RULE #3 - If this color is different from last colar and is a black and at the start.
	if thisColor != lastColor && thisColor == common.Black && start {

		if debug {
			fmt.Printf("\t\tRULE#3 -fixture %d If this color is different from last colar and a is a black and at the start.. start %t end %t bounce %t invert %t\n", fixtureNumber, start, end, bounce, invert)
		}

		if start && invert {
			// Fade down last color, so this black can be displayed.
			fadeColors = fadeDownColor(stepNumber, 3, "FadeDwn5", shift, fadeColors, lastColor, sequence, fixture, fixtureNumber)
		}
		//if !invert {
		fadeColors = fadeOffColor(stepNumber, 3, "FadeOff", fadeColors, thisColor, sequence, fixture, fixtureNumber)
		//}

		return fadeColors
	}

	// RULE #4 - If color is different from last color and color is a black.
	if thisColor != lastColor && thisColor == common.Black && !start {

		if debug {
			fmt.Printf("\t\tRULE#4 fixture %d If color is different from last color and color is a black. start %t end %t bounce %t invert %t\n", fixtureNumber, start, end, bounce, invert)
		}

		// Fade down last color, so this black can be displayed.
		fadeColors = fadeDownColor(stepNumber, 4, "FadeDwn6", shift, fadeColors, lastColor, sequence, fixture, fixtureNumber)

		// Now that we have faded down. Play the black for the off.
		// Now that we have faded down. Populate one up,on & down cycle with the black we asked for.
		fadeColors = makeABlack(stepNumber, 3, "makeABlack3", shift, fadeColors, thisColor, sequence, fixture, fixtureNumber)

		fadeColors = fadeOffColor(stepNumber, 3, "FadeOff2", fadeColors, thisColor, sequence, fixture, fixtureNumber)

		// If the next color is color fade back up.
		if nextColor != common.Black && end && invert {
			fadeColors = fadeUpColor(stepNumber, 2, "FadeUp_3", fadeColors, nextColor, sequence, fixture, fixtureNumber)
		}

		return fadeColors
	}

	if debug {
		fmt.Printf("\t\tRULE#5 fixture %d No rule fired. start %t end %t bounce %t invert %t\n", fixtureNumber, start, end, bounce, invert)
	}

	return fadeColors
}

func makeAColor(stepNumber int, rule int, debugMsg string, shift int, fadeColors map[int][]common.FixtureBuffer, color common.Color, sequence common.Sequence, fixture common.Fixture, fixtureNumber int) map[int][]common.FixtureBuffer {

	if debug {
		fmt.Printf("\t\t\tfixture:%d makeAColor color %+v\n", fixtureNumber, color)
	}

	var fade []int
	fade = append(fade, sequence.FadeUp...)
	fade = append(fade, sequence.FadeOn...)
	fade = append(fade, sequence.FadeDown...)
	var shiftCounter int
	for range fade {
		// if shiftCounter == shift {
		// 	break
		// }
		newColor := makeNewColor(stepNumber, rule, debugMsg, fixture, fixtureNumber, color, 255, sequence.ScannerChaser)
		fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
		shiftCounter++
	}

	return fadeColors
}

func makeABlack(stepNumber int, rule int, debugMsg string, shift int, fadeColors map[int][]common.FixtureBuffer, color common.Color, sequence common.Sequence, fixture common.Fixture, fixtureNumber int) map[int][]common.FixtureBuffer {

	if debug {
		fmt.Printf("\t\t\tfixture:%d makeABlack color %+v\n", fixtureNumber, color)
	}

	var fade []int
	fade = append(fade, sequence.FadeUp...)
	fade = append(fade, sequence.FadeOn...)
	fade = append(fade, sequence.FadeDown...)
	var shiftCounter int
	for range fade {
		// if shiftCounter == shift {
		// 	break
		// }
		newColor := makeNewColor(stepNumber, rule, debugMsg, fixture, fixtureNumber, color, 0, sequence.ScannerChaser)
		fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
		shiftCounter++
	}

	return fadeColors
}

func fadeDownColor(stepNumber int, rule int, debugMsg string, shift int, fadeColors map[int][]common.FixtureBuffer, color common.Color, sequence common.Sequence, fixture common.Fixture, fixtureNumber int) map[int][]common.FixtureBuffer {

	if debug {
		fmt.Printf("\t\t\tfixture:%d fadeDownColor color %+v\n", fixtureNumber, color)
	}

	var shiftCounter int
	for _, slope := range sequence.FadeDown {
		// if shiftCounter == shift {
		// 	break
		// }
		newColor := makeNewColor(stepNumber, rule, debugMsg, fixture, fixtureNumber, color, slope, sequence.ScannerChaser)
		fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
		shiftCounter++
	}
	return fadeColors
}

func fadeUpColor(stepNumber int, rule int, debugMsg string, fadeColors map[int][]common.FixtureBuffer, color common.Color, sequence common.Sequence, fixture common.Fixture, fixtureNumber int) map[int][]common.FixtureBuffer {

	if debug {
		fmt.Printf("\t\t\tfixture:%d fadeUpColor color %+v\n", fixtureNumber, color)
	}

	for _, slope := range sequence.FadeUp {
		newColor := makeNewColor(stepNumber, rule, debugMsg, fixture, fixtureNumber, color, slope, sequence.ScannerChaser)
		fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
	}
	return fadeColors
}

func fadeOnColor(stepNumber int, rule int, debugMsg string, fadeColors map[int][]common.FixtureBuffer, color common.Color, sequence common.Sequence, fixture common.Fixture, fixtureNumber int) map[int][]common.FixtureBuffer {

	if debug {
		fmt.Printf("\t\t\tfixture:%d fadeOnColor color %+v\n", fixtureNumber, color)
	}

	for _, slope := range sequence.FadeOn {
		newColor := makeNewColor(stepNumber, rule, debugMsg, fixture, fixtureNumber, color, slope, sequence.ScannerChaser)
		fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
	}
	return fadeColors
}

func fadeOffColor(stepNumber int, rule int, debugMsg string, fadeColors map[int][]common.FixtureBuffer, color common.Color, sequence common.Sequence, fixture common.Fixture, fixtureNumber int) map[int][]common.FixtureBuffer {

	if debug {
		fmt.Printf("\t\t\tfixture:%d fadeOnColor color %+v\n", fixtureNumber, color)
	}

	for _, slope := range sequence.FadeOn {
		newColor := makeNewColor(stepNumber, rule, debugMsg, fixture, fixtureNumber, color, slope, sequence.ScannerChaser)
		fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
	}
	return fadeColors
}

func keepSameAsLastTime(stepNumber int, rule int, debugMsg string, shift int, fadeColors map[int][]common.FixtureBuffer, color common.Color, sequence common.Sequence, fixture common.Fixture, fixtureNumber int) map[int][]common.FixtureBuffer {

	if debug {
		fmt.Printf("\t\t\tfixture:%d keepSameAsLastTime color %+v\n", fixtureNumber, color)
	}

	var fade []int
	fade = append(fade, sequence.FadeUp...)
	fade = append(fade, sequence.FadeOn...)
	fade = append(fade, sequence.FadeDown...)
	var shiftCounter int
	for range fade {
		// if shiftCounter == shift {
		// 	break
		// }
		newColor := makeNewColor(stepNumber, rule, debugMsg, fixture, fixtureNumber, color, 255, sequence.ScannerChaser)
		fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
		shiftCounter++
	}
	return fadeColors
}
