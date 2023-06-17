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

func CalculatePositions(stepsIn []common.Step, sequence common.Sequence, scanner bool) (map[int][]common.FixtureBuffer, int) {

	var steps []common.Step

	if debug {
		fmt.Printf("CalculatePositions Number Steps %d\n", len(stepsIn))
	}

	if sequence.RGBInvert {
		sequence.SequenceColors = common.HowManyColorsInSteps(stepsIn)
		steps = invertRGBColorsInSteps(stepsIn, sequence.SequenceColors)
		sequence.RGBInvert = false
		sequence.RGBNoFadeDown = true
	} else {
		steps = stepsIn
		sequence.RGBNoFadeDown = false
	}

	fadeColors := make(map[int][]common.FixtureBuffer)
	shift := common.Reverse(sequence.RGBShift)

	var numberFixtures int
	var numberFixturesInThisStep int
	var shiftCounter int
	var lastStep common.Step
	var lastStepNumber int

	if !sequence.ScannerInvert {

		// First loop make a space in the slope values for each fixture.
		for stepNumber, step := range steps {

			if debug {
				fmt.Printf("================== Step Number %d ============== No Fixtures %d\n", stepNumber, len(step.Fixtures))
			}

			lastStepNumber = stepNumber - 1
			if lastStepNumber < 0 {
				lastStepNumber = len(steps) - 1
			}

			lastStep = steps[lastStepNumber]
			numberFixturesInThisStep = 0

			for fixtureNumber := 0; fixtureNumber < len(step.Fixtures); fixtureNumber++ {

				if debug {
					fmt.Printf("\tFixture Number %d\n", fixtureNumber)
				}

				fixture := step.Fixtures[fixtureNumber]
				fixture.Enabled = sequence.FixtureState[fixtureNumber].Enabled
				for colorNumber, color := range fixture.Colors {

					if debug {
						fmt.Printf("\t\tColor Number %d\n", colorNumber)
					}

					// If color is same as last time do nothing.
					myfixture := lastStep.Fixtures[fixtureNumber]
					if color == myfixture.Colors[colorNumber] {

						if debug {
							fmt.Printf("\t\t\tIf color is same as last time do nothing. %+v\n", color)
						}

						var fade []int
						fade = append(fade, sequence.FadeUp...)
						fade = append(fade, sequence.FadeOn...)
						fade = append(fade, sequence.FadeDown...)

						shiftCounter = 0
						for range fade {
							if shiftCounter == shift {
								break
							}
							newColor := makeNewColor(fixture, fixtureNumber, color, 255, sequence.ScannerChaser)
							if debug {
								fmt.Printf("\t\t\t\tAdd1 %+v\n", newColor)
							}
							fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
							shiftCounter++
						}
						continue
					}

					// If color is different from last color and not black.
					if color != lastStep.Fixtures[fixtureNumber].Colors[colorNumber] && color != common.Black {
						if debug {
							fmt.Printf("\t\t\tIf color is different from last color and not black.. %+v\n", color)
						}
						if !sequence.RGBNoFadeDown {
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
						for _, slope := range sequence.FadeOn {
							newColor := makeNewColor(fixture, fixtureNumber, color, slope, sequence.ScannerChaser)
							if debug {
								fmt.Printf("\t\t\t\tAdd4 %+v\n", newColor)
							}
							fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
						}
						continue
					}

					// If color is different from last color and color is a black.
					if color != lastStep.Fixtures[fixtureNumber].Colors[colorNumber] && color == common.Black {
						if debug {
							fmt.Printf("\t\t\tIf color is different from last color and color is a black.. %+v\n", color)
						}
						for _, slope := range sequence.FadeDown {
							newColor := makeNewColor(fixture, fixtureNumber, lastStep.Fixtures[fixtureNumber].Colors[colorNumber], slope, sequence.ScannerChaser)
							if debug {
								fmt.Printf("\t\t\t\tAdd5 %+v\n", newColor)
							}
							fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
						}
						continue
					}

					if debug {
						fmt.Printf("\t\t\tDo Nothing %+v\n", color)
					}

				}
				numberFixturesInThisStep++
			}
			if numberFixturesInThisStep > numberFixtures {
				numberFixtures = numberFixturesInThisStep
			}
		}
	}

	return fadeColors, numberFixtures
}

// if sequence.Bounce || sequence.ScannerInvert {
// 	// Generate the positions in reverse.
// 	// Reverse the steps.
// 	for stepNumber := len(steps); stepNumber > 0; stepNumber-- {
// 		step := steps[stepNumber-1]
// 		numberFixturesInThisStep = 0
// 		for fixtureNumber := 0; fixtureNumber < len(step.Fixtures); fixtureNumber++ {
// 			fixture := step.Fixtures[fixtureNumber]
// 			fixture.Enabled = sequence.FixtureState[fixtureNumber].Enabled
// 			// Reverse the colors.
// 			noColors := len(fixture.Colors)
// 			for colorNumber := noColors; colorNumber > 0; colorNumber-- {
// 				color := fixture.Colors[colorNumber-1]
// 				if color.R > 0 || color.G > 0 || color.B > 0 || color.W > 0 || fixture.Shutter == 255 {
// 					// make space for a colored lamp.
// 					if !sequence.RGBInvert {
// 						// A faded up and down color.
// 						for _, slope := range sequence.FadeUp {
// 							newColor := makeNewColor(fixture, fixtureNumber, color, slope, sequence.ScannerChaser)
// 							fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
// 						}
// 					} else {
// 						// A solid on color.
// 						for range sequence.FadeUp {
// 							newColor := makeNewColor(fixture, fixtureNumber, color, 255, sequence.ScannerChaser)
// 							fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
// 						}
// 					}
// 				} else {
// 					if !sequence.RGBInvert {
// 						shiftCounter = 0
// 						// make space for a off lamp.
// 						for range sequence.FadeDown {
// 							if shiftCounter == shift {
// 								break
// 							}
// 							// A black lamp.
// 							newColor := makeNewColor(fixture, fixtureNumber, color, 0, sequence.ScannerChaser)
// 							fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
// 							shiftCounter++
// 						}
// 					} else {
// 						// A fading to black lamp.
// 						shiftCounter = 0
// 						for _, slope := range sequence.FadeDown {
// 							if shiftCounter == (shift + shift) {
// 								break
// 							}
// 							newColor := makeNewColor(fixture, fixtureNumber, color, slope, sequence.ScannerChaser)
// 							fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
// 							shiftCounter++
// 						}
// 					}
// 				}
// 			}
// 			numberFixturesInThisStep++
// 		}
// 		if numberFixturesInThisStep > numberFixtures {
// 			numberFixtures = numberFixturesInThisStep
// 		}
// 	}
// }

// if sequence.Bounce && sequence.ScannerInvert {
// 	// First loop make a space in the slope values for each fixture.
// 	for _, step := range steps {
// 		numberFixturesInThisStep = 0
// 		for fixtureNumber := 0; fixtureNumber < len(step.Fixtures); fixtureNumber++ {
// 			fixture := step.Fixtures[fixtureNumber]
// 			fixture.Enabled = sequence.FixtureState[fixtureNumber].Enabled
// 			for _, color := range fixture.Colors {
// 				if color.R > 0 || color.G > 0 || color.B > 0 || color.W > 0 || fixture.Shutter == 255 {
// 					// make space for a colored lamp.
// 					if !sequence.RGBInvert {
// 						// A faded up and down color.
// 						for _, slope := range sequence.FadeUp {
// 							newColor := makeNewColor(fixture, fixtureNumber, color, slope, sequence.ScannerChaser)
// 							fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
// 						}
// 					} else {
// 						// A solid on color.
// 						for range sequence.FadeUp {
// 							newColor := makeNewColor(fixture, fixtureNumber, color, 255, sequence.ScannerChaser)
// 							fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
// 						}
// 					}
// 				} else {
// 					if !sequence.RGBInvert {
// 						shiftCounter = 0
// 						// make space for a off lamp.
// 						for range sequence.FadeDown {
// 							if shiftCounter == shift {
// 								break
// 							}
// 							// A black lamp.
// 							newColor := makeNewColor(fixture, fixtureNumber, color, 0, sequence.ScannerChaser)
// 							fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
// 							shiftCounter++
// 						}
// 					} else {
// 						// A fading to black lamp.
// 						shiftCounter = 0
// 						for _, slope := range sequence.FadeDown {
// 							if shiftCounter == (shift + shift) {
// 								break
// 							}
// 							newColor := makeNewColor(fixture, fixtureNumber, color, slope, sequence.ScannerChaser)
// 							fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
// 							shiftCounter++
// 						}
// 					}
// 				}
// 			}
// 			numberFixturesInThisStep++
// 		}
// 		if numberFixturesInThisStep > numberFixtures {
// 			numberFixtures = numberFixturesInThisStep
// 		}
// 	}
// }

// Setup the counters for the lengths for each fixture.
// The number of steps is different for each fixture, depending on how
// many fades (tramsistions) take place in a pattern.
// Use the shortest for safety.
//if debug {
// 	fmt.Printf("FadeColors 0=%d ", len(fadeColors[0]))
// 	fmt.Printf("FadeColors 0=%d ", len(fadeColors[1]))
// 	fmt.Printf("FadeColors 0=%d ", len(fadeColors[2]))
// 	fmt.Printf("FadeColors 0=%d ", len(fadeColors[3]))
// 	fmt.Printf("FadeColors 0=%d ", len(fadeColors[4]))
// 	fmt.Printf("FadeColors 0=%d ", len(fadeColors[5]))
// 	fmt.Printf("FadeColors 0=%d ", len(fadeColors[6]))
// 	fmt.Printf("FadeColors 0=%d\n", len(fadeColors[7]))
// 	//}

// 	counter := 200
// 	for fixture := 0; fixture < numberFixtures; fixture++ {
// 		if len(fadeColors[fixture]) != 0 && len(fadeColors[fixture]) < counter {
// 			counter = len(fadeColors[fixture])
// 		}
// 	}

// 	//if debug {
// 	// Print out the fixtures so far.
// 	for fixture := 0; fixture < numberFixtures; fixture++ {
// 		fmt.Printf("Fixture %d\n", fixture)
// 		for out := 0; out < counter; out++ {
// 			fmt.Printf("%+v\n", fadeColors[fixture][out])
// 		}
// 		fmt.Printf("\n")
// 	}
// 	//}

// 	positionsOut := assemblePositions(fadeColors, counter, numberFixtures, sequence.FixtureState, sequence.RGBInvert, sequence.ScannerChaser, sequence.Optimisation)

// 	return positionsOut, len(positionsOut)

// }

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

func AssemblePositions(fadeColors map[int][]common.FixtureBuffer, numberFixtures int) (map[int]common.Position, int) {

	totalNumberOfSteps := 200
	for fixture := 0; fixture < numberFixtures; fixture++ {
		if len(fadeColors[fixture]) != 0 && len(fadeColors[fixture]) < totalNumberOfSteps {
			totalNumberOfSteps = len(fadeColors[fixture])
		}
	}

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
