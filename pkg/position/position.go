// Copyright (C) 2022, 2023 dhowlett99.
// This is the dmxlights position calculator, positions are generated from
// patterns. Positions control size, fade times and shifts.
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

const debug = false

// calculateLastStep Calculate last step.
func calculateLastStep(stepNumber int, steps []common.Step, reverse bool) common.Step {

	var lastStepNumber int
	var lastStep common.Step

	if !reverse {
		lastStepNumber = stepNumber - 1

		if lastStepNumber < 0 {
			lastStepNumber = len(steps) - 1
		}
	} else {
		lastStepNumber = stepNumber + 1

		if lastStepNumber > len(steps)-1 {
			lastStepNumber = 0
		}
	}

	lastStep = steps[lastStepNumber]

	return lastStep
}

// CalculatePositions takes a series of steps, examaines them to see if the step should fade up
func CalculatePositions(stepsIn []common.Step, sequence common.Sequence, scanner bool, patternShift int) (map[int][]common.FixtureBuffer, int, int) {

	var steps []common.Step
	var numberFixtures int
	var numberFixturesInThisStep int
	var lastStep common.Step

	fadeColors := make(map[int][]common.FixtureBuffer)
	shift := common.Reverse(sequence.RGBShift)

	if debug {
		fmt.Printf("CalculatePositions Number Steps %d\n", len(stepsIn))
	}

	if sequence.RGBInvert && sequence.RGBInvertOnce {
		sequence.SequenceColors = common.HowManyColorsInSteps(stepsIn)
		steps = invertRGBColorsInSteps(stepsIn, sequence.SequenceColors)
		sequence.RGBInvertOnce = false
	} else {
		steps = stepsIn
	}

	if !sequence.ScannerInvert {
		// Steps forward.
		for stepNumber, step := range steps {
			if debug {
				fmt.Printf("================== Step Number %d ============== No Fixtures %d\n", stepNumber, len(step.Fixtures))
			}

			// Calculate last step.
			lastStep = calculateLastStep(stepNumber, steps, false)

			// Start the fixtures counter.
			numberFixturesInThisStep = 0

			// Fixtures forward.
			for fixtureNumber := 0; fixtureNumber < len(step.Fixtures); fixtureNumber++ {
				fixture := step.Fixtures[fixtureNumber]

				if debug {
					fmt.Printf("\tFixture Number %d\n", fixtureNumber)
				}

				fixture.Enabled = sequence.FixtureState[fixtureNumber].Enabled

				// Color in forward.
				for colorNumber, color := range fixture.Colors {
					if debug {
						fmt.Printf("\t\tColor Number %d\n", colorNumber)
					}
					processColor(fadeColors, fixture, fixtureNumber, color, colorNumber, lastStep, sequence, shift, patternShift, scanner)
				}

				// Incremet the the fixture counter.
				numberFixturesInThisStep++
			}
			if numberFixturesInThisStep > numberFixtures {
				numberFixtures = numberFixturesInThisStep
			}
		}
	}

	if sequence.Bounce || sequence.ScannerInvert {
		// Generate the positions in reverse.
		// Reverse the steps.
		for stepNumber := len(steps); stepNumber > 0; stepNumber-- {
			step := steps[stepNumber-1]
			if debug {
				fmt.Printf("================== Step Number %d ============== No Fixtures %d\n", stepNumber, len(step.Fixtures))
			}

			lastStep = calculateLastStep(stepNumber, steps, true)

			numberFixturesInThisStep = 0

			for fixtureNumber := 0; fixtureNumber < len(step.Fixtures); fixtureNumber++ {
				if debug {
					fmt.Printf("\tFixture Number %d\n", fixtureNumber)
				}
				fixture := step.Fixtures[fixtureNumber]
				fixture.Enabled = sequence.FixtureState[fixtureNumber].Enabled

				// Reverse the colors.
				noColors := len(fixture.Colors)
				for colorNumber := noColors; colorNumber > 0; colorNumber-- {
					color := fixture.Colors[colorNumber-1]
					processColor(fadeColors, fixture, fixtureNumber, color, colorNumber-1, lastStep, sequence, shift, patternShift, scanner)
				}
				numberFixturesInThisStep++
			}
			if numberFixturesInThisStep > numberFixtures {
				numberFixtures = numberFixturesInThisStep
			}
		}
	}

	if sequence.Bounce && sequence.ScannerInvert {
		// Steps forward.
		for stepNumber, step := range steps {
			if debug {
				fmt.Printf("================== Step Number %d ============== No Fixtures %d\n", stepNumber, len(step.Fixtures))
			}

			// Calculate last step.
			lastStep = calculateLastStep(stepNumber, steps, false)

			// Start the fixtures counter.
			numberFixturesInThisStep = 0

			// Fixtures forward.
			for fixtureNumber := 0; fixtureNumber < len(step.Fixtures); fixtureNumber++ {
				fixture := step.Fixtures[fixtureNumber]

				if debug {
					fmt.Printf("\tFixture Number %d\n", fixtureNumber)
				}

				fixture.Enabled = sequence.FixtureState[fixtureNumber].Enabled

				// Colors forward.
				for colorNumber, color := range fixture.Colors {
					if debug {
						fmt.Printf("\t\tColor Number %d\n", colorNumber)
					}
					processColor(fadeColors, fixture, fixtureNumber, color, colorNumber, lastStep, sequence, shift, patternShift, scanner)
				}

				// Incremet the the fixture counter.
				numberFixturesInThisStep++
			}

			if numberFixturesInThisStep > numberFixtures {
				numberFixtures = numberFixturesInThisStep
			}
		}
	}

	if debug {
		fmt.Printf("FadeColors 0=%d ", len(fadeColors[0]))
		fmt.Printf("FadeColors 0=%d ", len(fadeColors[1]))
		fmt.Printf("FadeColors 0=%d ", len(fadeColors[2]))
		fmt.Printf("FadeColors 0=%d ", len(fadeColors[3]))
		fmt.Printf("FadeColors 0=%d ", len(fadeColors[4]))
		fmt.Printf("FadeColors 0=%d ", len(fadeColors[5]))
		fmt.Printf("FadeColors 0=%d ", len(fadeColors[6]))
		fmt.Printf("FadeColors 0=%d\n", len(fadeColors[7]))
	}

	// Setup the counters for the lengths for each fixture.
	// The number of steps is different for each fixture, depending on how
	// many fades (tramsistions) take place in a pattern.
	// Use the shortest for safety.
	totalNumberOfSteps := 1000
	for fixture := 0; fixture < numberFixtures; fixture++ {
		if len(fadeColors[fixture]) != 0 && len(fadeColors[fixture]) < totalNumberOfSteps {
			totalNumberOfSteps = len(fadeColors[fixture])
		}
	}

	if debug {
		// Print out the fixtures so far.
		for fixture := 0; fixture < numberFixtures; fixture++ {
			fmt.Printf("Fixture %d\n", fixture)
			for out := 0; out < totalNumberOfSteps; out++ {
				fmt.Printf("%+v\n", fadeColors[fixture][out])
			}
			fmt.Printf("\n")
		}
	}

	return fadeColors, numberFixtures, totalNumberOfSteps
}

func processColor(fadeColors map[int][]common.FixtureBuffer, fixture common.Fixture, fixtureNumber int, color common.Color, colorNumber int, lastStep common.Step, sequence common.Sequence, shift int, patternShift int, scanner bool) {

	// If color is same as last time , play that color out again.
	myfixture := lastStep.Fixtures[fixtureNumber]
	if color == myfixture.Colors[colorNumber] {

		if debug {
			fmt.Printf("\t\t\tIf color is same as last time do nothing. %+v\n", color)
		}

		var fade []int
		fade = append(fade, sequence.FadeUp...)

		if !scanner {
			if !sequence.RGBInvert {
				fade = append(fade, sequence.FadeOn...)
			}
			fade = append(fade, sequence.FadeDown...)
		}

		var shiftCounter = 0
		var actualShift int
		for range fade {
			if shift == 10 {
				actualShift = shift + patternShift
			} else {
				actualShift = shift
			}
			if shiftCounter == actualShift {
				break
			}
			newColor := makeNewColor(fixture, fixtureNumber, color, 255, sequence.ScannerChaser)
			if debug {
				fmt.Printf("\t\t\t\tAdd1 %+v\n", newColor)
			}
			fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
			shiftCounter++
		}
		return
	}

	// If color is different from last color and not black.
	if color != lastStep.Fixtures[fixtureNumber].Colors[colorNumber] && color != common.Black {
		if debug {
			fmt.Printf("\t\t\tIf color is different from last color and not black.. %+v\n", color)
		}
		if !scanner {
			// Fade down last color but only if last color wasn't a black.
			if lastStep.Fixtures[fixtureNumber].Colors[colorNumber] != common.Black {
				for _, slope := range sequence.FadeDown {
					newColor := makeNewColor(fixture, fixtureNumber, lastStep.Fixtures[fixtureNumber].Colors[colorNumber], slope, sequence.ScannerChaser)
					if debug {
						fmt.Printf("\t\t\t\tAdd2 %+v\n", newColor)
					}
					fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
				}
			}
		}

		// Fade up new color.
		for _, slope := range sequence.FadeUp {
			newColor := makeNewColor(fixture, fixtureNumber, color, slope, sequence.ScannerChaser)
			if debug {
				fmt.Printf("\t\t\t\tAdd3 %+v\n", newColor)
			}
			fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
		}
		if !scanner {
			if !sequence.RGBInvert {
				for _, slope := range sequence.FadeOn {
					newColor := makeNewColor(fixture, fixtureNumber, color, slope, sequence.ScannerChaser)
					if debug {
						fmt.Printf("\t\t\t\tAdd4 %+v\n", newColor)
					}
					fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
				}
			}
		}

		return
	}

	// If color is different from last color and color is a black.
	if color != lastStep.Fixtures[fixtureNumber].Colors[colorNumber] && color == common.Black {
		if debug {
			fmt.Printf("\t\t\tIf color is different from last color and color is a black.. %+v\n", color)
		}
		// Fade down last color, so this black can be displayed.
		for _, slope := range sequence.FadeDown {
			newColor := makeNewColor(fixture, fixtureNumber, lastStep.Fixtures[fixtureNumber].Colors[colorNumber], slope, sequence.ScannerChaser)
			if debug {
				fmt.Printf("\t\t\t\tAdd5 %+v\n", newColor)
			}
			fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
		}
		if !scanner {
			if sequence.RGBInvert {
				// Stay off for the on time.
				for range sequence.FadeOn {
					newColor := makeNewColor(fixture, fixtureNumber, lastStep.Fixtures[fixtureNumber].Colors[colorNumber], 0, sequence.ScannerChaser)
					if debug {
						fmt.Printf("\t\t\t\tAdd6 %+v\n", newColor)
					}
					fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
				}
			}
		}

		return
	}

	if debug {
		fmt.Printf("\t\t\tDo Nothing %+v\n", color)
	}

}
func makeNewColor(fixture common.Fixture, fixtureNumber int, color common.Color, insertValue int, chase bool) common.FixtureBuffer {

	if debug {
		fmt.Printf("makeNewColor fixture %d color %+v\n", fixtureNumber, color)
	}

	newColor := common.FixtureBuffer{}
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

func AssemblePositions(fadeColors map[int][]common.FixtureBuffer, numberFixtures int, totalNumberOfSteps int) (map[int]common.Position, int) {

	if debug {
		fmt.Printf("assemblePositions\n")
	}

	positionsOut := make(map[int]common.Position)

	if debug {
		fmt.Printf("totalNumberOfSteps %d\n", totalNumberOfSteps)
	}

	// Assemble the positions.
	for step := 0; step < totalNumberOfSteps; step++ {
		// Create a new position.
		newPosition := common.Position{}
		// Add some space for the fixtures.
		newPosition.Fixtures = make(map[int]common.Fixture)

		for fixture := 0; fixture <= numberFixtures; fixture++ {

			newFixture := common.Fixture{}
			newColor := common.Color{}

			lenghtOfSteps := len(fadeColors[fixture])
			if step < lenghtOfSteps {
				newColor.R = fadeColors[fixture][step].Color.R
				newColor.G = fadeColors[fixture][step].Color.G
				newColor.B = fadeColors[fixture][step].Color.B
				newColor.W = fadeColors[fixture][step].Color.W
				newFixture.Colors = append(newFixture.Colors, newColor)
				newFixture.Enabled = fadeColors[fixture][step].Enabled
				newFixture.Gobo = fadeColors[fixture][step].Gobo
				newFixture.Pan = fadeColors[fixture][step].Pan
				newFixture.Tilt = fadeColors[fixture][step].Tilt
				newFixture.Shutter = fadeColors[fixture][step].Shutter
				newFixture.MasterDimmer = fadeColors[fixture][step].MasterDimmer
				newFixture.Brightness = fadeColors[fixture][step].Brightness
				newPosition.Fixtures[fixture] = newFixture

			}
		}

		// Only add a position if there are some enabled scanners in the fixture list.
		if len(newPosition.Fixtures) != 0 {
			positionsOut[step] = newPosition
		}
	}

	if debug {
		for positionNumber := 0; positionNumber < len(positionsOut); positionNumber++ {
			position := positionsOut[positionNumber]
			fmt.Printf("Position %d\n", positionNumber)
			for fixture := 0; fixture < len(position.Fixtures); fixture++ {
				fmt.Printf("\tFixture %d Enabled %t Values %+v\n", fixture, position.Fixtures[fixture].Enabled, position.Fixtures[fixture].Colors)
			}
		}
	}

	return positionsOut, len(positionsOut)
}

func invertRGBColorsInSteps(steps []common.Step, colors []common.Color) []common.Step {

	var insertColor int
	numberColors := len(colors)

	var stepsOut []common.Step

	for _, step := range steps {

		newStep := common.Step{}

		newFixtures := make(map[int]common.Fixture)

		newStep.Fixtures = newFixtures

		for fixtureNumber, fixture := range step.Fixtures {

			newFixture := common.Fixture{}
			newFixture.MasterDimmer = fixture.MasterDimmer

			for _, color := range fixture.Colors {

				if insertColor >= numberColors {
					insertColor = 0
				}
				if color.R > 0 || color.G > 0 || color.B > 0 {
					// insert a black.
					newFixture.Colors = append(newFixture.Colors, common.Color{})
					insertColor++
				} else {
					// its a blank space so insert one of the colors.
					newFixture.Colors = append(newFixture.Colors, colors[insertColor])
				}

			}
			newStep.Fixtures[fixtureNumber] = newFixture
		}

		stepsOut = append(stepsOut, newStep)
	}

	return stepsOut
}
