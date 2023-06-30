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

func processColor(start bool, end bool, invert bool, fadeColors map[int][]common.FixtureBuffer, fixture common.Fixture, fixtureNumber int, thisColor common.Color, lastColor common.Color, nextColor common.Color, sequence common.Sequence, shift int, patternShift int, scanner bool) map[int][]common.FixtureBuffer {

	// If color is same as last time , play that color out again.
	if thisColor == lastColor {
		if debug {
			fmt.Printf("\t\t\tfixture %d If color is same as last time , play that color out again. start %t end %t\n", fixtureNumber, start, end)
		}
		fadeColors = keepSameAsLastTime(shift, fadeColors, thisColor, sequence, fixture, fixtureNumber)
		return fadeColors
	}

	// If color is different from last color and not black.
	if thisColor != lastColor && thisColor != common.Black {

		if debug {
			fmt.Printf("\t\t\tfixture %d If color is different from last color and not black. start %t end %t\n", fixtureNumber, start, end)
		}

		// Fade down last color but only if last color wasn't a black and we're not at the start.
		if lastColor != common.Black && !start && !end {
			fadeColors = fadeDownColor(shift, fadeColors, lastColor, sequence, fixture, fixtureNumber)
		}

		// Fade the color up.
		fadeColors = fadeUpColor(fadeColors, thisColor, sequence, fixture, fixtureNumber)

		// Leave the color for on time.
		fadeColors = fadeOnColor(fadeColors, thisColor, sequence, fixture, fixtureNumber)

		// If the next color is black. Fade dowm this color down ready.
		if nextColor == common.Black && !start && end {
			fadeColors = fadeDownColor(shift, fadeColors, thisColor, sequence, fixture, fixtureNumber)
		}

		return fadeColors
	}

	if thisColor != lastColor && thisColor == common.Black && start {
		fadeColors = makeABlack(shift, fadeColors, thisColor, sequence, fixture, fixtureNumber)
	}

	// If color is different from last color and color is a black.
	if thisColor != lastColor && thisColor == common.Black && !start {

		if debug {
			fmt.Printf("\t\t\tfixture %d If color is different from last color and color is a black. start %t end %t\n", fixtureNumber, start, end)
		}

		// Fade down last color, so this black can be displayed.
		fadeColors = fadeDownColor(shift, fadeColors, lastColor, sequence, fixture, fixtureNumber)

		// Now that we have faded down. Populate one up,on & down cycle with the black we asked for.
		fadeColors = makeABlack(shift, fadeColors, thisColor, sequence, fixture, fixtureNumber)

		return fadeColors
	}
	return fadeColors
}

func makeABlack(shift int, fadeColors map[int][]common.FixtureBuffer, color common.Color, sequence common.Sequence, fixture common.Fixture, fixtureNumber int) map[int][]common.FixtureBuffer {

	if debug {
		fmt.Printf("\t\t\tfixture:%d makeABlack color %+v\n", fixtureNumber, color)
	}

	var fade []int
	fade = append(fade, sequence.FadeUp...)
	fade = append(fade, sequence.FadeOn...)
	fade = append(fade, sequence.FadeDown...)
	var shiftCounter int
	for range fade {
		if shiftCounter == shift {
			break
		}
		newColor := makeNewColor(fixture, fixtureNumber, color, 0, sequence.ScannerChaser)
		fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
		shiftCounter++
	}

	return fadeColors
}

func fadeDownColor(shift int, fadeColors map[int][]common.FixtureBuffer, color common.Color, sequence common.Sequence, fixture common.Fixture, fixtureNumber int) map[int][]common.FixtureBuffer {

	if debug {
		fmt.Printf("\t\t\tfixture:%d fadeDownColor color %+v\n", fixtureNumber, color)
	}

	var shiftCounter int
	for _, slope := range sequence.FadeDown {
		if shiftCounter == shift {
			break
		}
		newColor := makeNewColor(fixture, fixtureNumber, color, slope, sequence.ScannerChaser)
		fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
		shiftCounter++
	}
	return fadeColors
}

func fadeUpColor(fadeColors map[int][]common.FixtureBuffer, color common.Color, sequence common.Sequence, fixture common.Fixture, fixtureNumber int) map[int][]common.FixtureBuffer {

	if debug {
		fmt.Printf("\t\t\tfixture:%d fadeUpColor color %+v\n", fixtureNumber, color)
	}

	for _, slope := range sequence.FadeUp {
		newColor := makeNewColor(fixture, fixtureNumber, color, slope, sequence.ScannerChaser)
		fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
	}
	return fadeColors
}

func fadeOnColor(fadeColors map[int][]common.FixtureBuffer, color common.Color, sequence common.Sequence, fixture common.Fixture, fixtureNumber int) map[int][]common.FixtureBuffer {

	if debug {
		fmt.Printf("\t\t\tfixture:%d fadeOnColor color %+v\n", fixtureNumber, color)
	}

	for _, slope := range sequence.FadeOn {
		newColor := makeNewColor(fixture, fixtureNumber, color, slope, sequence.ScannerChaser)
		fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
	}
	return fadeColors
}

func keepSameAsLastTime(shift int, fadeColors map[int][]common.FixtureBuffer, color common.Color, sequence common.Sequence, fixture common.Fixture, fixtureNumber int) map[int][]common.FixtureBuffer {

	if debug {
		fmt.Printf("\t\t\tfixture:%d keepSameAsLastTime color %+v\n", fixtureNumber, color)
	}

	var fade []int
	fade = append(fade, sequence.FadeUp...)
	fade = append(fade, sequence.FadeOn...)
	fade = append(fade, sequence.FadeDown...)
	var shiftCounter int
	for range fade {
		if shiftCounter == shift {
			break
		}
		newColor := makeNewColor(fixture, fixtureNumber, color, 255, sequence.ScannerChaser)
		fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
		shiftCounter++
	}
	return fadeColors
}
