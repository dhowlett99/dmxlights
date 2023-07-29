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

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/process"
)

const debug = false

// CalculatePositions takes a series of steps, examaines them to see if the step should fade up
func CalculatePositions(stepsIn []common.Step, sequence common.Sequence, scanner bool) (map[int][]common.FixtureBuffer, int, int) {

	var steps []common.Step
	var numberFixtures int
	var numberFixturesInThisStep int
	var lastStep common.Step
	var nextStep common.Step

	start := true
	end := false
	invert := false

	fadeColors := make(map[int][]common.FixtureBuffer)
	shift := common.Reverse(sequence.RGBShift)

	if debug {
		fmt.Printf("CalculatePositions Number Steps %d\n", len(stepsIn))
	}

	// Invert the RGB sequence.
	if sequence.RGBInvert && sequence.RGBInvertOnce {
		sequence.SequenceColors = common.HowManyColorsInSteps(stepsIn)
		steps = invertRGBColorsInSteps(stepsIn, sequence.SequenceColors)
		sequence.RGBInvertOnce = false
		invert = true
	} else {
		steps = stepsIn
		invert = false
	}

	if !sequence.ScannerInvert {
		// Steps forward.
		for stepNumber, step := range steps {
			if debug {
				fmt.Printf("==================================================================================================================================================\n")
				fmt.Printf("Step Number %d No Fixtures %d\n", stepNumber, len(step.Fixtures))
			}

			// Calculate last step.
			if stepNumber == 0 {
				end = false
				lastStep = steps[len(steps)-1]
			}
			// If we're at the end. next step is the first step.
			if stepNumber == len(steps)-1 {
				end = true
				nextStep = steps[0]
			} else {
				nextStep = steps[stepNumber+1]
			}

			// Start the fixtures counter.
			numberFixturesInThisStep = 0

			// Fixtures forward.
			for fixtureNumber := 0; fixtureNumber < len(step.Fixtures); fixtureNumber++ {
				thisFixture := step.Fixtures[fixtureNumber]
				thisFixture.Number = fixtureNumber
				thisFixture.Enabled = sequence.FixtureState[fixtureNumber].Enabled
				lastFixture := lastStep.Fixtures[fixtureNumber]
				nextFixture := nextStep.Fixtures[fixtureNumber]
				if scanner {
					fadeColors = process.ProcessScannerColor(stepNumber, start, end, sequence.Bounce, invert, fadeColors, &thisFixture, &lastFixture, &nextFixture, sequence, shift)
				} else {
					fadeColors = process.ProcessRGBColor(stepNumber, start, end, sequence.Bounce, invert, fadeColors, &thisFixture, &lastFixture, &nextFixture, sequence, shift)
					// Remember State of fixture.
					step.Fixtures[fixtureNumber] = thisFixture
				}
				// Incremet the the fixture counter.
				numberFixturesInThisStep++
			}

			// Done all fixtures, so not at the start any more.
			start = false
			if numberFixturesInThisStep > numberFixtures {
				numberFixtures = numberFixturesInThisStep
			}

			lastStep = step
		}
	}

	if sequence.Bounce || sequence.ScannerInvert {
		if debug {
			fmt.Printf("<<<<<<<<<<<<<<<<<<<<<<<<< START BOUNCE >>>>>>>>>>>>>>>>>>>>>>>>>\n")
		}
		// Generate the positions in reverse.
		// Reverse the steps.
		start = true
		for stepNumber := len(steps); stepNumber > 0; stepNumber-- {
			step := steps[stepNumber-1]
			if debug {
				fmt.Printf("==================================================================================================================================================\n")
				fmt.Printf("Step Number %d No Fixtures %d\n", stepNumber, len(step.Fixtures))
			}

			// If your at the start of this reversed list of steps (i.e at the end) and your invert the previous steps
			// have not played so the last step isn't set so set it here.
			if stepNumber == len(steps) {
				end = true // Because we play the step backwards the end is true for the first step.
				lastStep = steps[0]
			}

			// If we're at the begining. next step is the last step.
			if stepNumber == 0 {
				nextStep = steps[len(steps)-1]
				end = false
			} else {
				nextStep = steps[stepNumber-1]
			}

			numberFixturesInThisStep = 0

			for fixtureNumber := 0; fixtureNumber <= len(step.Fixtures)-1; fixtureNumber++ {
				thisFixture := step.Fixtures[fixtureNumber]
				thisFixture.Enabled = sequence.FixtureState[fixtureNumber].Enabled
				thisFixture.Number = fixtureNumber
				lastFixture := lastStep.Fixtures[fixtureNumber]
				nextFixture := nextStep.Fixtures[fixtureNumber]
				if scanner {
					fadeColors = process.ProcessScannerColor(stepNumber, start, end, sequence.Bounce, invert, fadeColors, &thisFixture, &lastFixture, &nextFixture, sequence, shift)
				} else {
					fadeColors = process.ProcessRGBColor(stepNumber, start, end, sequence.Bounce, invert, fadeColors, &thisFixture, &lastFixture, &nextFixture, sequence, shift)
					// Remember State of fixture.
					step.Fixtures[fixtureNumber] = thisFixture
				}
				numberFixturesInThisStep++
			}
			start = false
			if numberFixturesInThisStep > numberFixtures {
				numberFixtures = numberFixturesInThisStep
			}

			lastStep = step
		}
	}

	if sequence.Bounce && sequence.ScannerInvert {
		// Steps forward.
		for stepNumber, step := range steps {
			if debug {
				fmt.Printf("================== Step Number %d ============== No Fixtures %d\n", stepNumber, len(step.Fixtures))
			}

			// Start the fixtures counter.
			numberFixturesInThisStep = 0

			// Fixtures forward.
			for fixtureNumber := 0; fixtureNumber < len(step.Fixtures); fixtureNumber++ {
				thisFixture := step.Fixtures[fixtureNumber]
				thisFixture.Number = fixtureNumber
				thisFixture.Enabled = sequence.FixtureState[fixtureNumber].Enabled
				lastFixture := lastStep.Fixtures[fixtureNumber]
				nextFixture := nextStep.Fixtures[fixtureNumber]
				if scanner {
					fadeColors = process.ProcessScannerColor(stepNumber, start, end, sequence.Bounce, invert, fadeColors, &thisFixture, &lastFixture, &nextFixture, sequence, shift)
				} else {
					fadeColors = process.ProcessRGBColor(stepNumber, start, end, sequence.Bounce, invert, fadeColors, &thisFixture, &lastFixture, &nextFixture, sequence, shift)
					// Remember State of fixture.
					step.Fixtures[fixtureNumber] = thisFixture
				}
				// Incremet the the fixture counter.
				numberFixturesInThisStep++
			}

			if numberFixturesInThisStep > numberFixtures {
				numberFixtures = numberFixturesInThisStep
			}
			lastStep = step
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

	// if debug {
	// 	// Print out the fixtures so far.
	// 	for fixture := 0; fixture < numberFixtures; fixture++ {
	// 		fmt.Printf("Fixture %d\n", fixture)
	// 		for out := 0; out < totalNumberOfSteps; out++ {
	// 			fmt.Printf("%+v\n", fadeColors[fixture][out])
	// 		}
	// 		fmt.Printf("\n")
	// 	}
	// }

	return fadeColors, numberFixtures, totalNumberOfSteps
}

func AssemblePositions(fadeColors map[int][]common.FixtureBuffer, numberFixtures int, totalNumberOfSteps int, Optimisation bool) (map[int]common.Position, int) {

	if debug {
		fmt.Printf("assemblePositions\n")
	}

	positionsOut := make(map[int]common.Position)
	lampOn := make(map[int]bool)
	// Athough this looks odd, we need a seperate flag for lamp off
	// to make sure we apply the off message exactly once and not optimise it away
	// as this leaves a light on at the end of a sequence.
	lampOff := make(map[int]bool)

	if debug {
		fmt.Printf("totalNumberOfSteps %d\n", totalNumberOfSteps)
	}

	// Assemble the positions.
	for step := 0; step < totalNumberOfSteps; step++ {
		// Create a new position.
		newPosition := common.Position{}
		// Add some space for the fixtures.
		newPosition.Fixtures = make(map[int]common.Fixture)

		for fixtureNumber := 0; fixtureNumber <= numberFixtures; fixtureNumber++ {

			newFixture := common.Fixture{}
			newColor := common.Color{}
			lenghtOfSteps := len(fadeColors[fixtureNumber])
			if step < lenghtOfSteps {

				// We've found a color.
				if fadeColors[fixtureNumber][step].Color.R > 0 || fadeColors[fixtureNumber][step].Color.G > 0 || fadeColors[fixtureNumber][step].Color.B > 0 || fadeColors[fixtureNumber][step].Color.W > 0 {

					newColor.R = fadeColors[fixtureNumber][step].Color.R
					newColor.G = fadeColors[fixtureNumber][step].Color.G
					newColor.B = fadeColors[fixtureNumber][step].Color.B
					newColor.W = fadeColors[fixtureNumber][step].Color.W
					newFixture.Color = newColor
					newFixture.Enabled = fadeColors[fixtureNumber][step].Enabled
					newFixture.Gobo = fadeColors[fixtureNumber][step].Gobo
					newFixture.Pan = fadeColors[fixtureNumber][step].Pan
					newFixture.Tilt = fadeColors[fixtureNumber][step].Tilt
					newFixture.Shutter = fadeColors[fixtureNumber][step].Shutter
					lampOn[fixtureNumber] = true
					lampOff[fixtureNumber] = false
					newFixture.MasterDimmer = fadeColors[fixtureNumber][step].MasterDimmer
					newFixture.Brightness = fadeColors[fixtureNumber][step].Brightness
					newPosition.Fixtures[fixtureNumber] = newFixture
				} else {
					// turn the lamp off, but only if its already on.
					if lampOn[fixtureNumber] || !lampOff[fixtureNumber] || !Optimisation {
						newFixture.Color = common.Color{}
						newFixture.Enabled = fadeColors[fixtureNumber][step].Enabled
						newFixture.Gobo = fadeColors[fixtureNumber][step].Gobo
						newFixture.Pan = fadeColors[fixtureNumber][step].Pan
						newFixture.Tilt = fadeColors[fixtureNumber][step].Tilt
						newFixture.Shutter = fadeColors[fixtureNumber][step].Shutter
						lampOn[fixtureNumber] = false
						newFixture.MasterDimmer = fadeColors[fixtureNumber][step].MasterDimmer
						newFixture.Brightness = fadeColors[fixtureNumber][step].Brightness
						newPosition.Fixtures[fixtureNumber] = newFixture
					}
					lampOff[fixtureNumber] = true
				}
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
			for fixtureNumber := 0; fixtureNumber < len(position.Fixtures); fixtureNumber++ {
				fmt.Printf("\tFixture %d Enabled %t Values %+v\n", fixtureNumber, position.Fixtures[fixtureNumber].Enabled, position.Fixtures[fixtureNumber].Color)
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

			if insertColor >= numberColors {
				insertColor = 0
			}
			if fixture.Color.R > 0 || fixture.Color.G > 0 || fixture.Color.B > 0 {
				// insert a black.
				newFixture.Color = common.Color{}
			} else {
				// its a blank space so insert one of the colors.
				newFixture.Color = colors[insertColor]
				insertColor++
			}

			newStep.Fixtures[fixtureNumber] = newFixture
		}

		stepsOut = append(stepsOut, newStep)
	}

	return stepsOut
}
