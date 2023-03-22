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

func CalculatePositions(sequence common.Sequence) (map[int]common.Position, int) {

	if debug {
		fmt.Printf("CalculatePositions Number Steps %d\n", len(sequence.Steps))
	}

	fadeColors := make(map[int][]common.FixtureBuffer)
	shift := common.Reverse(sequence.RGBShift)

	var numberFixtures int
	var numberFixturesInThisStep int
	var shiftCounter int

	if !sequence.ScannerInvert {
		// First loop make a space in the slope values for each fixture.
		for _, step := range sequence.Steps {
			numberFixturesInThisStep = 0
			for fixtureNumber := 0; fixtureNumber < len(step.Fixtures); fixtureNumber++ {
				fixture := step.Fixtures[fixtureNumber]
				fixture.Enabled = sequence.ScannerState[fixtureNumber].Enabled
				for _, color := range fixture.Colors {
					if color.R > 0 || color.G > 0 || color.B > 0 || color.W > 0 || fixture.Shutter == 255 {
						// make space for a colored lamp.
						if !sequence.RGBInvert {
							// A faded up and down color.
							for _, slope := range sequence.FadeUpAndDown {
								//fmt.Printf("--CALC--> Fixture %d Enabled %t\n", fixtureNumber, sequence.ScannerState[fixtureNumber].Enabled)
								newColor := makeNewColor(fixture, fixtureNumber, color, slope, sequence.ScannerChase)
								fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
							}
						} else {
							// A solid on color.
							for range sequence.FadeUpAndDown {
								newColor := makeNewColor(fixture, fixtureNumber, color, 255, sequence.ScannerChase)
								fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
							}
						}
					} else {
						if !sequence.RGBInvert {
							shiftCounter = 0
							// make space for a off lamp.
							for range sequence.FadeDownAndUp {
								if shiftCounter == shift {
									break
								}
								// A black lamp.
								newColor := makeNewColor(fixture, fixtureNumber, color, 0, sequence.ScannerChase)
								fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
								shiftCounter++
							}
						} else {
							// A fading to black lamp.
							for _, slope := range sequence.FadeDownAndUp {
								newColor := makeNewColor(fixture, fixtureNumber, color, slope, sequence.ScannerChase)
								fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
							}
						}
					}
				}
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
		for stepNumber := len(sequence.Steps); stepNumber > 0; stepNumber-- {
			step := sequence.Steps[stepNumber-1]
			numberFixturesInThisStep = 0
			for fixtureNumber := 0; fixtureNumber < len(step.Fixtures); fixtureNumber++ {
				fixture := step.Fixtures[fixtureNumber]
				fixture.Enabled = sequence.ScannerState[fixtureNumber].Enabled
				// Reverse the colors.
				noColors := len(fixture.Colors)
				for colorNumber := noColors; colorNumber > 0; colorNumber-- {
					color := fixture.Colors[colorNumber-1]
					if color.R > 0 || color.G > 0 || color.B > 0 || color.W > 0 || fixture.Shutter == 255 {
						// make space for a colored lamp.
						if !sequence.RGBInvert {
							// A faded up and down color.
							for _, slope := range sequence.FadeUpAndDown {
								newColor := makeNewColor(fixture, fixtureNumber, color, slope, sequence.ScannerChase)
								fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
							}
						} else {
							// A solid on color.
							for range sequence.FadeUpAndDown {
								newColor := makeNewColor(fixture, fixtureNumber, color, 255, sequence.ScannerChase)
								fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
							}
						}
					} else {
						if !sequence.RGBInvert {
							shiftCounter = 0
							// make space for a off lamp.
							for range sequence.FadeDownAndUp {
								if shiftCounter == shift {
									break
								}
								// A black lamp.
								newColor := makeNewColor(fixture, fixtureNumber, color, 0, sequence.ScannerChase)
								fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
								shiftCounter++
							}
						} else {
							// A fading to black lamp.
							for _, slope := range sequence.FadeDownAndUp {
								newColor := makeNewColor(fixture, fixtureNumber, color, slope, sequence.ScannerChase)
								fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
							}
						}
					}
				}
				numberFixturesInThisStep++
			}
			if numberFixturesInThisStep > numberFixtures {
				numberFixtures = numberFixturesInThisStep
			}
		}
	}

	if sequence.Bounce && sequence.ScannerInvert {
		// First loop make a space in the slope values for each fixture.
		for _, step := range sequence.Steps {
			numberFixturesInThisStep = 0
			for fixtureNumber := 0; fixtureNumber < len(step.Fixtures); fixtureNumber++ {
				fixture := step.Fixtures[fixtureNumber]
				fixture.Enabled = sequence.ScannerState[fixtureNumber].Enabled
				for _, color := range fixture.Colors {
					if color.R > 0 || color.G > 0 || color.B > 0 || color.W > 0 || fixture.Shutter == 255 {
						// make space for a colored lamp.
						if !sequence.RGBInvert {
							// A faded up and down color.
							for _, slope := range sequence.FadeUpAndDown {
								newColor := makeNewColor(fixture, fixtureNumber, color, slope, sequence.ScannerChase)
								fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
							}
						} else {
							// A solid on color.
							for range sequence.FadeUpAndDown {
								newColor := makeNewColor(fixture, fixtureNumber, color, 255, sequence.ScannerChase)
								fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
							}
						}
					} else {
						if !sequence.RGBInvert {
							shiftCounter = 0
							// make space for a off lamp.
							for range sequence.FadeDownAndUp {
								if shiftCounter == shift {
									break
								}
								// A black lamp.
								newColor := makeNewColor(fixture, fixtureNumber, color, 0, sequence.ScannerChase)
								fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
								shiftCounter++
							}
						} else {
							// A fading to black lamp.
							for _, slope := range sequence.FadeDownAndUp {
								newColor := makeNewColor(fixture, fixtureNumber, color, slope, sequence.ScannerChase)
								fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
							}
						}
					}
				}
				numberFixturesInThisStep++
			}
			if numberFixturesInThisStep > numberFixtures {
				numberFixtures = numberFixturesInThisStep
			}
		}
	}

	// Setup the counters for the lengths for each fixture.
	// The number of steps is different for each fixture, depending on how
	// many fades (tramsistions) take place in a pattern.
	// Use the shortest for safety.
	if debug {
		fmt.Printf("FadeColors 0=%d", len(fadeColors[0]))
		fmt.Printf("FadeColors 0=%d", len(fadeColors[1]))
		fmt.Printf("FadeColors 0=%d", len(fadeColors[2]))
		fmt.Printf("FadeColors 0=%d", len(fadeColors[3]))
		fmt.Printf("FadeColors 0=%d", len(fadeColors[4]))
		fmt.Printf("FadeColors 0=%d", len(fadeColors[5]))
		fmt.Printf("FadeColors 0=%d", len(fadeColors[6]))
		fmt.Printf("FadeColors 0=%d\n", len(fadeColors[7]))
	}

	counter := 200
	for fixture := 0; fixture < numberFixtures; fixture++ {
		if len(fadeColors[fixture]) != 0 && len(fadeColors[fixture]) < counter {
			counter = len(fadeColors[fixture])
		}
	}

	if debug {
		// Print out the fixtures so far.
		for fixture := 0; fixture < numberFixtures; fixture++ {
			fmt.Printf("Fixture %d\n", fixture)
			for out := 0; out < counter; out++ {
				fmt.Printf("%+v\n", fadeColors[fixture][out])
			}
			fmt.Printf("\n")
		}
	}

	positionsOut := assemblePositions(fadeColors, counter, sequence.EnabledNumberFixtures, sequence.ScannerState, sequence.RGBInvert, sequence.ScannerChase, sequence.Optimisation)

	// Add scanner positions. Chase mode means we overlay the scanner pan and tilt on top of the RGB fade values.
	if sequence.ScannerChase {
		positionsOut = overlayScannerPositions(sequence.ScannerSteps, positionsOut)
	}

	return positionsOut, len(positionsOut)

}

func overlayScannerPositions(steps []common.Step, positionsIn map[int]common.Position) map[int]common.Position {

	if debug {
		fmt.Printf("overlayScannerPositions has %d steps\n", len(steps))
	}

	positionsOut := make(map[int]common.Position)

	numberPositions := len(positionsIn)
	numberScannerSteps := len(steps)

	var scannerPosition int = 0

	// Step through the positions
	for step := 0; step <= numberPositions; step++ {
		if scannerPosition >= numberScannerSteps {
			scannerPosition = 0
		}

		// Create a new position.
		newPosition := common.Position{}
		// Add some space for the fixtures.
		newPosition.Fixtures = make(map[int]common.Fixture)

		// All fixtures have the same rotation for now.

		for fixtureNumber := 0; fixtureNumber < len(positionsIn[step].Fixtures); fixtureNumber++ {
			newFixture := common.Fixture{}
			newFixture.Enabled = positionsIn[step].Fixtures[fixtureNumber].Enabled
			newFixture.ID = positionsIn[step].Fixtures[fixtureNumber].ID
			newFixture.Name = positionsIn[step].Fixtures[fixtureNumber].Name
			newFixture.Label = positionsIn[step].Fixtures[fixtureNumber].Label
			newFixture.MasterDimmer = positionsIn[step].Fixtures[fixtureNumber].MasterDimmer
			newFixture.Brightness = positionsIn[step].Fixtures[fixtureNumber].Brightness
			newFixture.ScannerColor = positionsIn[step].Fixtures[fixtureNumber].ScannerColor
			newFixture.Colors = positionsIn[step].Fixtures[fixtureNumber].Colors
			newFixture.Shutter = positionsIn[step].Fixtures[fixtureNumber].Shutter
			newFixture.Rotate = positionsIn[step].Fixtures[fixtureNumber].Rotate
			newFixture.Music = positionsIn[step].Fixtures[fixtureNumber].Music
			newFixture.Gobo = positionsIn[step].Fixtures[fixtureNumber].Gobo
			newFixture.Program = positionsIn[step].Fixtures[fixtureNumber].Program
			newFixture.Pan = steps[scannerPosition].Fixtures[fixtureNumber].Pan
			newFixture.Tilt = steps[scannerPosition].Fixtures[fixtureNumber].Tilt

			newPosition.Fixtures[fixtureNumber] = newFixture

		}

		// Move to the next scanner position.
		scannerPosition++

		// Only add a position if there are some enabled scanners in the fixture list.
		if len(newPosition.Fixtures) != 0 {
			positionsOut[step] = newPosition
		}
	}

	return positionsOut
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

func assemblePositions(fadeColors map[int][]common.FixtureBuffer, totalNumberOfSteps int, enabledNumberFixtures int, scannerState map[int]common.ScannerState, RGBInvert bool, chase bool, Optimisation bool) map[int]common.Position {

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

		for fixture := 0; fixture <= enabledNumberFixtures; fixture++ {

			newFixture := common.Fixture{}
			newColor := common.Color{}

			lenghtOfSteps := len(fadeColors[fixture])
			if step < lenghtOfSteps {

				newColor.R = fadeColors[fixture][step].Color.R
				newColor.G = fadeColors[fixture][step].Color.G
				newColor.B = fadeColors[fixture][step].Color.B
				newColor.W = fadeColors[fixture][step].Color.W

				// Optimisation is applied in this step. We only play out off's to the universe if the lamp is already on.
				// And in the case of inverted playout only colors if the lamp is already on.
				if !RGBInvert {
					// We've found a color.
					if fadeColors[fixture][step].Color.R > 0 || fadeColors[fixture][step].Color.G > 0 || fadeColors[fixture][step].Color.B > 0 || fadeColors[fixture][step].Color.W > 0 {
						newFixture.Colors = append(newFixture.Colors, newColor)
						newFixture.Enabled = fadeColors[fixture][step].Enabled
						newFixture.Gobo = fadeColors[fixture][step].Gobo
						newFixture.Pan = fadeColors[fixture][step].Pan
						newFixture.Tilt = fadeColors[fixture][step].Tilt
						newFixture.Shutter = fadeColors[fixture][step].Shutter
						lampOn[fixture] = true
						lampOff[fixture] = false
						newFixture.MasterDimmer = fadeColors[fixture][step].MasterDimmer
						newFixture.Brightness = fadeColors[fixture][step].Brightness
						newPosition.Fixtures[fixture] = newFixture
					} else {
						// turn the lamp off, but only if its already on.
						if lampOn[fixture] || !lampOff[fixture] || !Optimisation {
							newFixture.Colors = append(newFixture.Colors, common.Color{})
							newFixture.Enabled = fadeColors[fixture][step].Enabled
							newFixture.Gobo = fadeColors[fixture][step].Gobo
							newFixture.Pan = fadeColors[fixture][step].Pan
							newFixture.Tilt = fadeColors[fixture][step].Tilt
							newFixture.Shutter = fadeColors[fixture][step].Shutter
							lampOn[fixture] = false
							newFixture.MasterDimmer = fadeColors[fixture][step].MasterDimmer
							newFixture.Brightness = fadeColors[fixture][step].Brightness
							newPosition.Fixtures[fixture] = newFixture
						}
						lampOff[fixture] = true
					}
				} else {
					// We've found a color. turn it on but only if its already off.
					if fadeColors[fixture][step].Color.R > 0 || fadeColors[fixture][step].Color.G > 0 || fadeColors[fixture][step].Color.B > 0 || fadeColors[fixture][step].Color.W > 0 {
						if !lampOn[fixture] || !Optimisation {
							newFixture.Colors = append(newFixture.Colors, newColor)
							newFixture.Enabled = fadeColors[fixture][step].Enabled
							newFixture.Gobo = fadeColors[fixture][step].Gobo
							newFixture.Pan = fadeColors[fixture][step].Pan
							newFixture.Tilt = fadeColors[fixture][step].Tilt
							newFixture.Shutter = fadeColors[fixture][step].Shutter
							lampOn[fixture] = true
							newFixture.MasterDimmer = fadeColors[fixture][step].MasterDimmer
							newFixture.Brightness = fadeColors[fixture][step].Brightness
							newPosition.Fixtures[fixture] = newFixture
						}
					} else {
						// turn the lamp off
						newFixture.Colors = append(newFixture.Colors, common.Color{})
						newFixture.Enabled = fadeColors[fixture][step].Enabled
						newFixture.Gobo = fadeColors[fixture][step].Gobo
						newFixture.Pan = fadeColors[fixture][step].Pan
						newFixture.Tilt = fadeColors[fixture][step].Tilt
						newFixture.Shutter = fadeColors[fixture][step].Shutter
						lampOn[fixture] = false
						newFixture.MasterDimmer = fadeColors[fixture][step].MasterDimmer
						newFixture.Brightness = fadeColors[fixture][step].Brightness
						newPosition.Fixtures[fixture] = newFixture
					}
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
			for fixture := 0; fixture < len(position.Fixtures); fixture++ {
				fmt.Printf("\tFixture %d Enabled %t Values %+v\n", fixture, position.Fixtures[fixture].Enabled, position.Fixtures[fixture].Colors)
			}
		}
	}

	return positionsOut
}
