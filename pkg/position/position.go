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

func CalculatePositions(sequence common.Sequence, slopeOn []int, slopeOff []int, Optimisation bool, scannerState map[int]common.ScannerState) (map[int]common.Position, int) {

	fadeColors := make(map[int][]common.FixtureBuffer)
	shift := common.Reverse(sequence.RGBShift)

	var numberFixtures int
	var numberFixturesInThisStep int
	var shiftCounter int

	if !sequence.ScannerInvert {
		// First loop make a space in the slope values for each fixture.
		for _, step := range sequence.Steps {
			numberFixturesInThisStep = 0
			for fixtureNumber, fixture := range step.Fixtures {
				for _, color := range fixture.Colors {
					if color.R > 0 || color.G > 0 || color.B > 0 || color.W > 0 {
						// make space for a colored lamp.
						if !sequence.RGBInvert {
							// A faded up and down color.
							for _, slope := range slopeOn {
								newColor := makeNewColor(fixture, fixtureNumber, color, slope)
								fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
							}
						} else {
							// A solid on color.
							for range slopeOn {
								newColor := makeNewColor(fixture, fixtureNumber, color, 255)
								fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
							}
						}
					} else {
						if !sequence.RGBInvert {
							shiftCounter = 0
							// make space for a off lamp.
							for range slopeOff {
								if shiftCounter == shift {
									break
								}
								// A black lamp.
								newColor := makeNewColor(fixture, fixtureNumber, color, 0)
								fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
								shiftCounter++
							}
						} else {
							// A fading to black lamp.
							for _, slope := range slopeOff {
								newColor := makeNewColor(fixture, fixtureNumber, color, slope)
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
			for fixtureNumber, fixture := range step.Fixtures {
				// Reverse the colors.
				noColors := len(fixture.Colors)
				for colorNumber := noColors; colorNumber > 0; colorNumber-- {
					color := fixture.Colors[colorNumber-1]
					if color.R > 0 || color.G > 0 || color.B > 0 || color.W > 0 {
						// make space for a colored lamp.
						if !sequence.RGBInvert {
							// A faded up and down color.
							for _, slope := range slopeOn {
								newColor := makeNewColor(fixture, fixtureNumber, color, slope)
								fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
							}
						} else {
							// A solid on color.
							for range slopeOn {
								newColor := makeNewColor(fixture, fixtureNumber, color, 255)
								fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
							}
						}
					} else {
						if !sequence.RGBInvert {
							shiftCounter = 0
							// make space for a off lamp.
							for range slopeOff {
								if shiftCounter == shift {
									break
								}
								// A black lamp.
								newColor := makeNewColor(fixture, fixtureNumber, color, 0)
								fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
								shiftCounter++
							}
						} else {
							// A fading to black lamp.
							for _, slope := range slopeOff {
								newColor := makeNewColor(fixture, fixtureNumber, color, slope)
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
			for fixtureNumber, fixture := range step.Fixtures {
				for _, color := range fixture.Colors {
					if color.R > 0 || color.G > 0 || color.B > 0 || color.W > 0 {
						// make space for a colored lamp.
						if !sequence.RGBInvert {
							// A faded up and down color.
							for _, slope := range slopeOn {
								newColor := makeNewColor(fixture, fixtureNumber, color, slope)
								fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
							}
						} else {
							// A solid on color.
							for range slopeOn {
								newColor := makeNewColor(fixture, fixtureNumber, color, 255)
								fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
							}
						}
					} else {
						if !sequence.RGBInvert {
							shiftCounter = 0
							// make space for a off lamp.
							for range slopeOff {
								if shiftCounter == shift {
									break
								}
								// A black lamp.
								newColor := makeNewColor(fixture, fixtureNumber, color, 0)
								fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
								shiftCounter++
							}
						} else {
							// A fading to black lamp.
							for _, slope := range slopeOff {
								newColor := makeNewColor(fixture, fixtureNumber, color, slope)
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

	// It appears counters arn't always the same.
	counter1 := len(fadeColors[0])
	counter2 := len(fadeColors[1])
	counter3 := len(fadeColors[2])
	counter4 := len(fadeColors[3])

	// Use the shortest for safety.
	counter := counter1
	if counter2 < counter1 {
		counter = counter2
	}
	if counter3 < counter {
		counter = counter3
	}
	if counter4 < counter {
		counter = counter4
	}

	if debug {
		// Print out the fixtures so far.
		for fixture := 0; fixture < numberFixtures; fixture++ {
			fmt.Printf("Fixture ")
			for out := 0; out < counter; out++ {
				fmt.Printf("%v", fadeColors[fixture][out])
			}
			fmt.Printf("\n")
		}
	}

	positionsOut := AssemblePositions(fadeColors, counter, numberFixtures, scannerState, sequence.RGBInvert, sequence.ScannerChase, Optimisation)
	return positionsOut, len(positionsOut)

}

func makeNewColor(fixture common.Fixture, fixtureNumber int, color common.Color, insertValue int) common.FixtureBuffer {
	newColor := common.FixtureBuffer{}
	newColor.Color = common.Color{}
	newColor.Gobo = fixture.Gobo
	newColor.Pan = fixture.Pan
	newColor.Tilt = fixture.Tilt
	newColor.Shutter = fixture.Shutter
	newColor.ScannerNumber = fixtureNumber
	newColor.Color.R = int(math.Round((float64(color.R) / 100) * (float64(insertValue) / 2.55)))
	newColor.Color.G = int(math.Round((float64(color.G) / 100) * (float64(insertValue) / 2.55)))
	newColor.Color.B = int(math.Round((float64(color.B) / 100) * (float64(insertValue) / 2.55)))
	newColor.Color.W = int(math.Round((float64(color.W) / 100) * (float64(insertValue) / 2.55)))
	newColor.MasterDimmer = fixture.MasterDimmer
	return newColor
}

func AssemblePositions(fadeColors map[int][]common.FixtureBuffer, totalNumberOfSteps int, numberFixtures int, scannerState map[int]common.ScannerState, RGBInvert bool, chase bool, Optimisation bool) map[int]common.Position {

	positionsOut := make(map[int]common.Position)
	lampOn := make(map[int]bool)
	lampOff := make(map[int]bool)

	// Assemble the positions.
	for step := 0; step < totalNumberOfSteps; step++ {
		// Create a new position.
		newPosition := common.Position{}
		// Add some space for the fixtures.
		newPosition.Fixtures = make(map[int]common.Fixture)

		if debug {
			fmt.Printf("totalNumberOfSteps %d\n", totalNumberOfSteps)
		}

		for fixture := 0; fixture < numberFixtures; fixture++ {
			if scannerState[fixture].Enabled {
				newFixture := common.Fixture{}

				newColor := common.Color{}
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
						newFixture.Gobo = fadeColors[fixture][step].Gobo
						newFixture.Pan = fadeColors[fixture][step].Pan
						newFixture.Tilt = fadeColors[fixture][step].Tilt
						newFixture.Shutter = fadeColors[fixture][step].Shutter
						lampOn[fixture] = true
						lampOff[fixture] = false
						newFixture.MasterDimmer = fadeColors[fixture][step].MasterDimmer
						newPosition.Fixtures[fixture] = newFixture
					} else {
						// turn the lamp off, but only if its already on.
						if lampOn[fixture] || !lampOff[fixture] || !Optimisation {
							newFixture.Colors = append(newFixture.Colors, common.Color{})
							newFixture.Gobo = fadeColors[fixture][step].Gobo
							newFixture.Pan = fadeColors[fixture][step].Pan
							newFixture.Tilt = fadeColors[fixture][step].Tilt
							newFixture.Shutter = fadeColors[fixture][step].Shutter
							lampOn[fixture] = false
							newFixture.MasterDimmer = fadeColors[fixture][step].MasterDimmer
							newPosition.Fixtures[fixture] = newFixture
						}
						lampOff[fixture] = true
					}
				} else {
					// We've found a color. turn it on but only if its already off.
					if fadeColors[fixture][step].Color.R > 0 || fadeColors[fixture][step].Color.G > 0 || fadeColors[fixture][step].Color.B > 0 || fadeColors[fixture][step].Color.W > 0 {
						if !lampOn[fixture] || !Optimisation {
							newFixture.Colors = append(newFixture.Colors, newColor)
							newFixture.Gobo = fadeColors[fixture][step].Gobo
							newFixture.Pan = fadeColors[fixture][step].Pan
							newFixture.Tilt = fadeColors[fixture][step].Tilt
							newFixture.Shutter = fadeColors[fixture][step].Shutter
							lampOn[fixture] = true
							newFixture.MasterDimmer = fadeColors[fixture][step].MasterDimmer
							newPosition.Fixtures[fixture] = newFixture
						}
					} else {
						// turn the lamp off
						newFixture.Colors = append(newFixture.Colors, common.Color{})
						newFixture.Gobo = fadeColors[fixture][step].Gobo
						newFixture.Pan = fadeColors[fixture][step].Pan
						newFixture.Tilt = fadeColors[fixture][step].Tilt
						newFixture.Shutter = fadeColors[fixture][step].Shutter
						lampOn[fixture] = false
						newFixture.MasterDimmer = fadeColors[fixture][step].MasterDimmer
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
		// Print out the positions in fixtures order.
		for step := 0; step < len(positionsOut); step++ {
			fmt.Printf("Position %d: Fixtures %+v\n", step, positionsOut[step].Fixtures)
		}
	}

	return positionsOut
}
