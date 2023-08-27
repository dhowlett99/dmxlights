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

	fadeColors := make(map[int][]common.FixtureBuffer)
	shift := common.Reverse(sequence.RGBShift)

	if debug {
		fmt.Printf("CalculatePositions Number Steps %d\n", len(stepsIn))
	}

	// Apply inverted selection from fixtureState to the RGB sequence.
	if sequence.Type == "rgb" {
		sequence.SequenceColors = common.HowManyColorsInSteps(stepsIn)
		steps = invertRGBColorsInSteps(stepsIn, sequence.SequenceColors, sequence.FixtureState)
	} else {
		steps = stepsIn
	}

	if !sequence.ScannerReverse {
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
					fadeColors = process.ProcessScannerColor(stepNumber, start, end, sequence.Bounce, thisFixture.Inverted, fadeColors, &thisFixture, &lastFixture, &nextFixture, sequence, shift)
				} else {
					if (sequence.Pattern.Name == "Pairs" ||
						sequence.Pattern.Name == "Flash" ||
						sequence.Pattern.Name == "Inward") && !sequence.Bounce {
						fadeColors = process.ProcessSimpleColor(stepNumber, start, end, sequence.Bounce, thisFixture.Inverted, fadeColors, &thisFixture, &lastFixture, &nextFixture, sequence, shift)
					} else {
						if scanner {
							fadeColors = process.ProcessScannerColor(stepNumber, start, end, sequence.Bounce, thisFixture.Inverted, fadeColors, &thisFixture, &lastFixture, &nextFixture, sequence, shift)
						} else {
							fadeColors = process.ProcessRGBColor(stepNumber, start, end, sequence.Bounce, thisFixture.Inverted, fadeColors, &thisFixture, &lastFixture, &nextFixture, sequence, shift)
							// Remember State of fixture.
							step.Fixtures[fixtureNumber] = thisFixture
						}
					}
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

	if sequence.Bounce || sequence.ScannerReverse {
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
					fadeColors = process.ProcessScannerColor(stepNumber, start, end, sequence.Bounce, thisFixture.Inverted, fadeColors, &thisFixture, &lastFixture, &nextFixture, sequence, shift)
				} else {
					if (sequence.Pattern.Name == "Pairs" ||
						sequence.Pattern.Name == "Flash" ||
						sequence.Pattern.Name == "Inward") && !sequence.Bounce {
						fadeColors = process.ProcessSimpleColor(stepNumber, start, end, sequence.Bounce, thisFixture.Inverted, fadeColors, &thisFixture, &lastFixture, &nextFixture, sequence, shift)
					} else {
						if scanner {
							fadeColors = process.ProcessScannerColor(stepNumber, start, end, sequence.Bounce, thisFixture.Inverted, fadeColors, &thisFixture, &lastFixture, &nextFixture, sequence, shift)
						} else {
							fadeColors = process.ProcessRGBColor(stepNumber, start, end, sequence.Bounce, thisFixture.Inverted, fadeColors, &thisFixture, &lastFixture, &nextFixture, sequence, shift)
							// Remember State of fixture.
							step.Fixtures[fixtureNumber] = thisFixture
						}
					}
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

	if sequence.Bounce && sequence.ScannerReverse {
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
					fadeColors = process.ProcessScannerColor(stepNumber, start, end, sequence.Bounce, thisFixture.Inverted, fadeColors, &thisFixture, &lastFixture, &nextFixture, sequence, shift)
				} else {
					if (sequence.Pattern.Name == "Pairs" ||
						sequence.Pattern.Name == "Flash" ||
						sequence.Pattern.Name == "Inward") && !sequence.Bounce {
						fadeColors = process.ProcessSimpleColor(stepNumber, start, end, sequence.Bounce, thisFixture.Inverted, fadeColors, &thisFixture, &lastFixture, &nextFixture, sequence, shift)
					} else {
						if scanner {
							fadeColors = process.ProcessScannerColor(stepNumber, start, end, sequence.Bounce, thisFixture.Inverted, fadeColors, &thisFixture, &lastFixture, &nextFixture, sequence, shift)
						} else {
							fadeColors = process.ProcessRGBColor(stepNumber, start, end, sequence.Bounce, thisFixture.Inverted, fadeColors, &thisFixture, &lastFixture, &nextFixture, sequence, shift)
							// Remember State of fixture.
							step.Fixtures[fixtureNumber] = thisFixture
						}
					}
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
		fmt.Printf("FadeColors Fixture 0=%d ", len(fadeColors[0]))
		fmt.Printf("FadeColors Fixture 1=%d ", len(fadeColors[1]))
		fmt.Printf("FadeColors Fixture 2=%d ", len(fadeColors[2]))
		fmt.Printf("FadeColors Fixture 3=%d ", len(fadeColors[3]))
		fmt.Printf("FadeColors Fixture 4=%d ", len(fadeColors[4]))
		fmt.Printf("FadeColors Fixture 5=%d ", len(fadeColors[5]))
		fmt.Printf("FadeColors Fixture 6=%d ", len(fadeColors[6]))
		fmt.Printf("FadeColors Fixture 7=%d\n", len(fadeColors[7]))
	}

	// Setup the counters for the lengths for each fixture.
	// The number of steps is different for each fixture, depending on how
	// many fades (tramsistions) take place in a pattern.
	// Use the shortest for safety.
	totalNumberOfSteps := 1000
	for fixtureNumber := 0; fixtureNumber < numberFixtures; fixtureNumber++ {
		if sequence.FixtureState[fixtureNumber].Enabled {
			if len(fadeColors[fixtureNumber]) != 0 && len(fadeColors[fixtureNumber]) < totalNumberOfSteps {
				totalNumberOfSteps = len(fadeColors[fixtureNumber])
			}
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

func AssemblePositions(fadeColors map[int][]common.FixtureBuffer, numberFixtures int, totalNumberOfSteps int, fixtureState map[int]common.FixtureState, optimisation bool) (map[int]common.Position, int) {

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

	var newStep int

	// Assemble the positions.
	for step := 0; step < totalNumberOfSteps; step++ {
		// Create a new position.
		newPosition := common.Position{}
		// Add some space for the fixtures.
		newPosition.Fixtures = make(map[int]common.Fixture)

		for fixtureNumber := 0; fixtureNumber < numberFixtures; fixtureNumber++ {

			if !fixtureState[fixtureNumber].Enabled {
				continue
			}

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
					if lampOn[fixtureNumber] || !lampOff[fixtureNumber] || !optimisation {
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
			positionsOut[newStep] = newPosition
			newStep++
		}
	}

	if debug {
		for positionNumber := 0; positionNumber < len(positionsOut); positionNumber++ {
			position := positionsOut[positionNumber]
			fmt.Printf("Position %d\n", positionNumber)
			for fixtureNumber := 0; fixtureNumber < len(position.Fixtures); fixtureNumber++ {
				fmt.Printf("\tFixture %d Enabled %t Values %+v Brightness %d \n", fixtureNumber, position.Fixtures[fixtureNumber].Enabled, position.Fixtures[fixtureNumber].Color, position.Fixtures[fixtureNumber].Brightness)
			}
		}
	}

	return positionsOut, len(positionsOut)
}

func invertRGBColorsInSteps(steps []common.Step, colors []common.Color, fixtureState map[int]common.FixtureState) []common.Step {

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

			if fixtureState[fixtureNumber].RGBInverted {
				if fixture.Color.R > 0 || fixture.Color.G > 0 || fixture.Color.B > 0 {
					// insert a black.
					newFixture.Color = common.Color{}
				} else {
					// its a blank space so insert one of the colors.
					newFixture.Color = colors[insertColor]
					insertColor++
				}
				newFixture.Inverted = true
				newStep.Fixtures[fixtureNumber] = newFixture
			} else {
				// Just copy what was there.
				fixture.Inverted = false
				newStep.Fixtures[fixtureNumber] = fixture
			}
		}

		stepsOut = append(stepsOut, newStep)
	}

	return stepsOut
}

// ApplyFixtureState - Apply the state of the fixtures to the pattern, fixture disabling works by disabling the
// steps that have no enabled fixtures AND also disabling in the fixure package. If we only disable here we don't
// catch steps that have more than one fixture alight in any one step.
// So make sure you also turn off the fixture in the fixture receiver.
func ApplyFixtureState(patternIn common.Pattern, scannerState map[int]common.FixtureState) common.Pattern {

	debug := false

	generatedSteps := patternIn.Steps

	var patternOut common.Pattern

	patternOut.Name = patternIn.Name
	patternOut.Label = patternIn.Label
	patternOut.Steps = []common.Step{}

	if debug {
		for fixture := 0; fixture < len(scannerState); fixture++ {
			fmt.Printf("ApplyFixtureState: Fixture:%d State %t\n", fixture, scannerState[fixture].Enabled)
		}
	}

	for _, step := range generatedSteps {

		newStep := step
		newStep.Fixtures = make(map[int]common.Fixture)
		hasColors := make(map[int]bool)
		for fixtureNumber, fixture := range step.Fixtures {
			newFixture := common.Fixture{}
			newFixture.Enabled = scannerState[fixtureNumber].Enabled
			newFixture.MasterDimmer = 255
			newFixture.Shutter = fixture.Shutter
			newFixture.Color = fixture.Color

			if fixture.Color.R > 0 || fixture.Color.G > 0 || fixture.Color.B > 0 {
				hasColors[fixtureNumber] = true
			} else {
				hasColors[fixtureNumber] = false
			}

			newStep.Fixtures[fixtureNumber] = newFixture
		}

		if debug {
			fmt.Printf("Fixtures \n")
			for fixture := 0; fixture < len(newStep.Fixtures); fixture++ {
				fmt.Printf("Fixture %d Enabled %t Colors %+v\n", fixture, newStep.Fixtures[fixture].Enabled, newStep.Fixtures[fixture].Color)
			}
		}

		for fixtureNumber, fixture := range newStep.Fixtures {
			// Don't add steps with no enabled fixtures.
			if hasColors[fixtureNumber] && fixture.Enabled {
				patternOut.Steps = append(patternOut.Steps, newStep)
				break
			}
		}
		// // Always add key steps.
		if step.KeyStep {
			patternOut.Steps = append(patternOut.Steps, newStep)
		}
	}

	if debug {
		for _, step := range patternOut.Steps {
			fmt.Printf("Fixtures \n")
			for fixture := 0; fixture < len(step.Fixtures); fixture++ {
				fmt.Printf("Fixture %d Enabled %t Values %+v\n", fixture, step.Fixtures[fixture].Enabled, step.Fixtures[fixture])
			}
		}
	}

	return patternOut

}
