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

package process

import (
	"fmt"
	"math"

	"github.com/dhowlett99/dmxlights/pkg/common"
)

const debug = false
const UNKNOWN = 0
const MAKE = 1
const FADEUP = 2
const FADEDOWN = 3

// processScannerColor takes this color and next color and adds a fade color to the fadeColors map.
// This function uses simple rules to decide which fade value to add.
func ProcessScannerColor(stepNumber int, start bool, end bool, bounce bool, invert bool, fadeColors map[int][]common.FixtureBuffer, thisFixture *common.Fixture, lastFixture *common.Fixture, nextFixture *common.Fixture, sequence common.Sequence, shift int) map[int][]common.FixtureBuffer {

	fadeColors = fadeUpColor(stepNumber, 2, "FadeUp_0", shift, fadeColors, thisFixture.Color, sequence, thisFixture)
	return fadeColors

}

func ProcessSimpleColor(stepNumber int, start bool, end bool, bounce bool, invert bool, fadeColors map[int][]common.FixtureBuffer, thisFixture *common.Fixture, lastFixture *common.Fixture, nextFixture *common.Fixture, sequence common.Sequence, shift int) map[int][]common.FixtureBuffer {
	fadeColors = fadeUpColor(stepNumber, 2, "FadeUp_0", shift, fadeColors, thisFixture.Color, sequence, thisFixture)
	fadeColors = fadeDownColor(stepNumber, 2, "FadeDwn1", shift, fadeColors, thisFixture.Color, sequence, thisFixture)
	return fadeColors
}

// processRGBColor takes this color and next color and adds a fade color to the fadeColors map.
// This function uses simple rules to decide which fade value to add.
// This is complicated because we want to support complex patterns which have colors stay on
// from one step to another, the simple fade up the color then fade down the color doesn't work
// for these use cases.
//
// There are a number of different inputs which controll the way we process the fading of the color.
//  1. We notice when we are at the start and the end of sequences, which prevents us starting with the fading
//     down of the last fixture and also helps us manage the change in direction of bounces.
//  2. We notice when we are bouncing to control the pattern reversal.
//  3. We notice when we select inverted sequence so we can cater for the pattern being upside down.
//  4. We pass pointers to the thisFixture, lastFixture and nextFixture so we know what we did last and
//     also we use these pointers to modify the state of the fixture, used for tracking when we last faded down.
//     Only used in the multicolor patterns.
//  5. We care about the next color being Black as we don't need to fade to black if we are already black.
//
// Note We pass debug messages into the fixture buffer, for each operation, so that when in debug mode you
// can see how the fadeColors were created.
func ProcessRGBColor(stepNumber int, start bool, end bool, bounce bool, invert bool, fadeColors map[int][]common.FixtureBuffer, thisFixture *common.Fixture, lastFixture *common.Fixture, nextFixture *common.Fixture, sequence common.Sequence, shift int) map[int][]common.FixtureBuffer {

	// RULE #1 - If color is same as last time , play that color out again.
	thisColor := thisFixture.Color
	lastColor := lastFixture.Color
	if thisColor == lastColor {
		if debug {
			fmt.Printf("\t\tRULE#1 - fixture %d If color is same as last time , play that color out again. start %t end %t bounce %t invert %t\n", thisFixture.Number, start, end, bounce, invert)
		}
		fadeColors = makeAColor(stepNumber, 1, "SameColr1", shift, fadeColors, thisFixture.Color, sequence, thisFixture)

		return fadeColors
	}

	// RULE 2 - If color is different from last color and not black.
	if thisFixture.Color != lastFixture.Color && thisFixture.Color != common.Black {

		if debug {
			fmt.Printf("\t\tRULE#2 -fixture %d If color is different from last color and not black. start %t end %t bounce %t invert %t\n", thisFixture.Number, start, end, bounce, invert)
		}

		// Fade down last color but only if last color wasn't a black and we're not at the start.
		if lastFixture.Color != common.Black && !start && !end && lastFixture.State != FADEDOWN {
			fadeColors = fadeDownColor(stepNumber, 2, fmt.Sprintf("FadeDwn1 this state %d last state %d", thisFixture.State, lastFixture.State), shift, fadeColors, lastFixture.Color, sequence, thisFixture)
		}

		if !invert {
			fadeColors = fadeUpColor(stepNumber, 2, "FadeUp_1", shift, fadeColors, thisFixture.Color, sequence, thisFixture)
		}

		// Fade the color up.
		if !start && invert {
			fadeColors = fadeUpColor(stepNumber, 2, "FadeUp_2", shift, fadeColors, thisFixture.Color, sequence, thisFixture)
		}

		// Make a color
		if invert {
			fadeColors = makeAColor(stepNumber, 3, "MakeCLR1", shift, fadeColors, nextFixture.Color, sequence, thisFixture)
		}

		// If the next color is black. Fade dowm this color down ready.
		if nextFixture.Color == common.Black && !start && end && thisFixture.State != FADEDOWN {
			fadeColors = fadeDownColor(stepNumber, 2, "FadeDwn3", shift, fadeColors, thisFixture.Color, sequence, thisFixture)
		}

		// If the next color is another color but not black. Fade dowm this color down ready.
		if lastFixture.Color != common.Black && nextFixture.Color != common.Black && thisFixture.Color != nextFixture.Color && !start && thisFixture.State != FADEDOWN {
			fadeColors = fadeDownColor(stepNumber, 2, fmt.Sprintf("FadeDwn4 state %d", thisFixture.State), shift, fadeColors, thisFixture.Color, sequence, thisFixture)
		}

		// If next color is not black and we're at the end or bouncing and we're the first fixture.
		if nextFixture.Color != common.Black && !start && end && bounce && thisFixture.Number == 0 && thisFixture.State != FADEDOWN {
			fadeColors = fadeDownColor(stepNumber, 2, "FadeDwn5", shift, fadeColors, nextFixture.Color, sequence, thisFixture)
		}
		return fadeColors
	}

	// RULE #3 - If this color is different from last color and is a black and at the start.
	if thisFixture.Color != lastFixture.Color && thisFixture.Color == common.Black && start {

		if debug {
			fmt.Printf("\t\tRULE#3 -fixture %d If this color is different from last colar and a is a black and at the start.. start %t end %t bounce %t invert %t\n", thisFixture.Number, start, end, bounce, invert)
		}

		if start && invert {
			//Fade down last color, so this black can be displayed.
			fadeColors = fadeDownColor(stepNumber, 3, "FadeDwn6", shift, fadeColors, lastFixture.Color, sequence, thisFixture)
		} else {
			fadeColors = makeAColor(stepNumber, 3, "MakeCLR1", shift, fadeColors, thisFixture.Color, sequence, thisFixture)
		}
		return fadeColors
	}

	// RULE #4 - If color is different from last color and color is a black.
	if thisFixture.Color != lastFixture.Color && thisFixture.Color == common.Black && !start {

		if debug {
			fmt.Printf("\t\tRULE#4 fixture %d If color is different from last color and color is a black. start %t end %t bounce %t invert %t\n", thisFixture.Number, start, end, bounce, invert)
		}

		// Fade down last color, so this black can be displayed.
		fadeColors = fadeDownColor(stepNumber, 4, "FadeDwn7", shift, fadeColors, lastFixture.Color, sequence, thisFixture)

		// Now that we have faded down. Populate one up,on & down cycle with the black we asked for.
		if !invert {
			fadeColors = makeAColor(stepNumber, 3, "makeABlack3", shift, fadeColors, thisFixture.Color, sequence, thisFixture)
		}

		// If the next color is color fade back up.
		if nextFixture.Color != common.Black && end && invert {
			fadeColors = fadeUpColor(stepNumber, 2, "FadeUp_3", shift, fadeColors, nextFixture.Color, sequence, thisFixture)
		}

		return fadeColors
	}

	if debug {
		fmt.Printf("\t\tNO RULE FIRED - fixture %d No rule fired. start %t end %t bounce %t invert %t\n", thisFixture.Number, start, end, bounce, invert)
	}

	return fadeColors
}

// fadeDownColor fades down the given color using the sequences fade down and fade off values.
func fadeDownColor(stepNumber int, rule int, debugMsg string, shift int, fadeColors map[int][]common.FixtureBuffer, color common.Color, sequence common.Sequence, thisFixture *common.Fixture) map[int][]common.FixtureBuffer {

	if debug {
		fmt.Printf("\t\t\t\tfixture:%d fadeDownColor color %+v\n", thisFixture.Number, color)
	}

	thisFixture.State = FADEDOWN

	var fade []int
	fade = append(fade, sequence.FadeDown...)
	fade = append(fade, sequence.FadeOff...)

	if color.R > 0 || color.G > 0 || color.B > 0 {
		for _, slope := range fade {
			newColor := addColor(stepNumber, rule, debugMsg, thisFixture, color, slope, sequence.ScannerChaser)
			fadeColors[thisFixture.Number] = append(fadeColors[thisFixture.Number], newColor)
		}
	} else {
		var shiftCounter int
		for _, slope := range fade {
			if shiftCounter == shift {
				break
			}
			newColor := addColor(stepNumber, rule, debugMsg, thisFixture, color, slope, sequence.ScannerChaser)
			fadeColors[thisFixture.Number] = append(fadeColors[thisFixture.Number], newColor)
			shiftCounter++
		}
	}

	return fadeColors
}

// fadeUpColor fades up the given color using the sequences fade up and fade on values.
func fadeUpColor(stepNumber int, rule int, debugMsg string, shift int, fadeColors map[int][]common.FixtureBuffer, color common.Color, sequence common.Sequence, thisFixture *common.Fixture) map[int][]common.FixtureBuffer {

	if debug {
		fmt.Printf("\t\t\t\tfixture:%d fadeUpColor color %+v\n", thisFixture.Number, color)
	}

	thisFixture.State = FADEUP

	var fade []int
	fade = append(fade, sequence.FadeUp...)
	fade = append(fade, sequence.FadeOn...)

	if color.R > 0 || color.G > 0 || color.B > 0 {
		for _, slope := range fade {
			newColor := addColor(stepNumber, rule, debugMsg, thisFixture, color, slope, sequence.ScannerChaser)
			fadeColors[thisFixture.Number] = append(fadeColors[thisFixture.Number], newColor)
		}
	} else {
		var shiftCounter int
		for _, slope := range fade {
			if shiftCounter == shift {
				break
			}
			newColor := addColor(stepNumber, rule, debugMsg, thisFixture, color, slope, sequence.ScannerChaser)
			fadeColors[thisFixture.Number] = append(fadeColors[thisFixture.Number], newColor)
			shiftCounter++
		}
	}

	return fadeColors
}

// makeAColor is used to add a color to the fixture buffer map of size fadeUp, fadeOn, FadeDown and fadeOff, which is the width of one cycle.
func makeAColor(stepNumber int, rule int, debugMsg string, shift int, fadeColors map[int][]common.FixtureBuffer, color common.Color, sequence common.Sequence, thisFixture *common.Fixture) map[int][]common.FixtureBuffer {

	if debug {
		fmt.Printf("\t\t\t\tfixture:%d makeAColor color %+v\n", thisFixture.Number, color)
	}

	thisFixture.State = MAKE

	var fade []int
	fade = append(fade, sequence.FadeUp...)
	fade = append(fade, sequence.FadeOn...)
	fade = append(fade, sequence.FadeDown...)
	fade = append(fade, sequence.FadeOff...)

	var shiftCounter int
	var insertValue int
	for range fade {
		if shiftCounter == shift {
			break
		}
		// If we insert a color that is black, also reduce the insert value
		// so the brightness is also set to zero for scanners that only
		// have shutter control.
		if color.R == 0 && color.G == 0 && color.B == 0 {
			insertValue = 0
		} else {
			insertValue = 255
		}
		newColor := addColor(stepNumber, rule, debugMsg, thisFixture, color, insertValue, sequence.ScannerChaser)
		fadeColors[thisFixture.Number] = append(fadeColors[thisFixture.Number], newColor)
		shiftCounter++
	}

	return fadeColors
}

// addColor adds a color to the fixtures buffer array, which is used later for assembling the positions.
func addColor(stepNumber int, rule int, debugMsg string, thisFixture *common.Fixture, color common.Color, insertValue int, chase bool) common.FixtureBuffer {

	if debug {
		fmt.Printf("\t\t\t\t\tStep %d func=%s addColor fixture %d color %+v slope %d\n", stepNumber, debugMsg, thisFixture.Number, color, insertValue)
	}

	newColor := common.FixtureBuffer{}

	if debug {
		newColor.DebugMsg = debugMsg
		newColor.Step = stepNumber
		newColor.Rule = rule
	}

	newColor.Color = common.Color{}
	newColor.Gobo = thisFixture.Gobo
	newColor.Pan = thisFixture.Pan
	newColor.Tilt = thisFixture.Tilt
	newColor.Shutter = thisFixture.Shutter
	newColor.Color.R = int(math.Round((float64(color.R) / 100) * (float64(insertValue) / 2.55)))
	newColor.Color.G = int(math.Round((float64(color.G) / 100) * (float64(insertValue) / 2.55)))
	newColor.Color.B = int(math.Round((float64(color.B) / 100) * (float64(insertValue) / 2.55)))
	newColor.Color.W = int(math.Round((float64(color.W) / 100) * (float64(insertValue) / 2.55)))

	if !chase {
		newColor.Brightness = 255
	} else {
		newColor.Brightness = int(math.Round((float64(thisFixture.MasterDimmer) / 100) * (float64(insertValue) / 2.55)))
	}

	newColor.Enabled = thisFixture.Enabled
	newColor.MasterDimmer = thisFixture.MasterDimmer

	return newColor
}
