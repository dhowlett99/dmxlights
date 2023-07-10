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

func processColor(stepNumber int, start bool, end bool, bounce bool, invert bool, fadeColors map[int][]common.FixtureBuffer, fixture common.Fixture, thisColor common.Color, lastColor common.Color, nextColor common.Color, sequence common.Sequence, shift int, patternShift int, scanner bool) map[int][]common.FixtureBuffer {

	// RULE #1 - If color is same as last time , play that color out again.
	if thisColor == lastColor {
		if debug {
			fmt.Printf("\t\tRULE#1 - fixture %d If color is same as last time , play that color out again. start %t end %t bounce %t invert %t\n", fixture.Number, start, end, bounce, invert)
		}
		fadeColors = keepSameAsLastTime(stepNumber, 1, "SameColr", shift, fadeColors, thisColor, sequence, fixture)
		return fadeColors
	}

	// RULE 2 - If color is different fade down old fade up new.
	if thisColor != lastColor {

		if debug {
			fmt.Printf("\t\tRULE#2 -fixture %d If color is different from last color and not black. start %t end %t bounce %t invert %t\n", fixture.Number, start, end, bounce, invert)
		}

		if lastColor != common.Black {
			fadeColors = fadeDownColor(stepNumber, 2, "FadeDwn2", shift, fadeColors, thisColor, sequence, fixture)
			fadeColors = fadeUpColor(stepNumber, 2, "FadeUp_1", shift, fadeColors, thisColor, sequence, fixture, fixture.Number)
		} else {
			fadeColors = fadeUpColor(stepNumber, 2, "FadeUp_1", shift, fadeColors, thisColor, sequence, fixture, fixture.Number)
			fadeColors = fadeDownColor(stepNumber, 2, "FadeDwn2", shift, fadeColors, thisColor, sequence, fixture)
		}
		return fadeColors
	}

	if debug {
		fmt.Printf("\t\tRULE#5 fixture %d No rule fired. start %t end %t bounce %t invert %t\n", fixture.Number, start, end, bounce, invert)
	}

	return fadeColors
}

func fadeDownColor(stepNumber int, rule int, debugMsg string, shift int, fadeColors map[int][]common.FixtureBuffer, color common.Color, sequence common.Sequence, fixture common.Fixture) map[int][]common.FixtureBuffer {

	if debug {
		fmt.Printf("\t\t\t\tfixture:%d fadeDownColor color %+v\n", fixture.Number, color)
	}

	if color.R > 0 || color.G > 0 || color.B > 0 {
		for _, slope := range sequence.FadeDown {
			newColor := makeNewColor(stepNumber, rule, debugMsg, fixture, color, slope, sequence.ScannerChaser)
			fadeColors[fixture.Number] = append(fadeColors[fixture.Number], newColor)
		}
	} else {
		var shiftCounter int
		for _, slope := range sequence.FadeDown {
			if shiftCounter == shift {
				break
			}
			newColor := makeNewColor(stepNumber, rule, debugMsg, fixture, color, slope, sequence.ScannerChaser)
			fadeColors[fixture.Number] = append(fadeColors[fixture.Number], newColor)
			shiftCounter++
		}
	}

	return fadeColors
}

func fadeUpColor(stepNumber int, rule int, debugMsg string, shift int, fadeColors map[int][]common.FixtureBuffer, color common.Color, sequence common.Sequence, fixture common.Fixture, fixtureNumber int) map[int][]common.FixtureBuffer {

	if debug {
		fmt.Printf("\t\t\t\tfixture:%d fadeUpColor color %+v\n", fixtureNumber, color)
	}

	if color.R > 0 || color.G > 0 || color.B > 0 {
		for _, slope := range sequence.FadeUp {
			newColor := makeNewColor(stepNumber, rule, debugMsg, fixture, color, slope, sequence.ScannerChaser)
			fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
		}
	} else {
		var shiftCounter int
		for _, slope := range sequence.FadeUp {
			if shiftCounter == shift {
				break
			}
			newColor := makeNewColor(stepNumber, rule, debugMsg, fixture, color, slope, sequence.ScannerChaser)
			fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
			shiftCounter++
		}
	}

	return fadeColors
}

func keepSameAsLastTime(stepNumber int, rule int, debugMsg string, shift int, fadeColors map[int][]common.FixtureBuffer, color common.Color, sequence common.Sequence, fixture common.Fixture) map[int][]common.FixtureBuffer {

	if debug {
		fmt.Printf("\t\t\t\tfixture:%d keepSameAsLastTime color %+v\n", fixture.Number, color)
	}

	var fade []int
	fade = append(fade, sequence.FadeUp...)
	fade = append(fade, sequence.FadeOn...)
	fade = append(fade, sequence.FadeDown...)

	if color.R > 0 || color.G > 0 || color.B > 0 {
		for range fade {
			newColor := makeNewColor(stepNumber, rule, debugMsg, fixture, color, 255, sequence.ScannerChaser)
			fadeColors[fixture.Number] = append(fadeColors[fixture.Number], newColor)
		}
	} else {
		var shiftCounter int
		for range fade {
			if shiftCounter == shift {
				break
			}
			newColor := makeNewColor(stepNumber, rule, debugMsg, fixture, color, 255, sequence.ScannerChaser)
			fadeColors[fixture.Number] = append(fadeColors[fixture.Number], newColor)
			shiftCounter++
		}
	}

	return fadeColors
}
