// Copyright (C) 2024 dhowlett99.
// This is the dmxlights VU chaser pattern generator.
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

package pattern

import (
	"fmt"
	"math"

	"github.com/dhowlett99/dmxlights/pkg/colors"
	"github.com/dhowlett99/dmxlights/pkg/common"
)

type MaxMins struct {
	Min int
	Max int
}

func getGroups(patternSize int, numLights int) []MaxMins {
	width := int(math.Floor(float64(numLights) / float64(patternSize)))
	groups := make([]MaxMins, patternSize)

	for i := 0; i < int(patternSize); i++ {
		groups[i] = MaxMins{
			Min: i * width,
			Max: (i + 1) * width,
		}
	}
	if debug {
		fmt.Printf("Width %d Groups %d\n", width, groups)
	}

	return groups
}

// Generate dynamic chase pattern based on the number of fixtures.
func generateVuChasePattern(numberOfFixtures int) common.Pattern {

	if debug {
		fmt.Printf("generateVuChasePattern for fixtures %d\n", numberOfFixtures)
	}

	// Divide up the fixtures into 8 groups used by this pattern.
	// groups := numberOfFixtures / 8
	// maxGroup := groups * numberOfFixtures
	// fmt.Printf("Number of Groups %d MaxGroup %d\n", groups, maxGroup)
	patterSize := 8
	groups := getGroups(patterSize, numberOfFixtures)

	// Define fixtures.
	redFixture := common.Fixture{
		MasterDimmer: full,
		Enabled:      true,
		Color:        colors.Red,
	}
	orangeFixture := common.Fixture{
		MasterDimmer: full,
		Enabled:      true,
		Color:        colors.Orange,
	}
	yellowFixture := common.Fixture{
		MasterDimmer: full,
		Enabled:      true,
		Color:        colors.Yellow,
	}
	greenFixture := common.Fixture{
		MasterDimmer: full,
		Enabled:      true,
		Color:        colors.Green,
	}
	blackFixture := common.Fixture{
		MasterDimmer: full,
		Enabled:      true,
		Color:        colors.Black,
	}

	chaseSteps := []common.Step{
		{
			StepNumber: 0,
			KeyStep:    false,
			Fixtures: map[int]common.Fixture{
				0: greenFixture,
				1: blackFixture,
				2: blackFixture,
				3: blackFixture,
				4: blackFixture,
				5: blackFixture,
				6: blackFixture,
				7: blackFixture,
			},
		},
		{
			StepNumber: 1,
			KeyStep:    false,
			Fixtures: map[int]common.Fixture{
				0: greenFixture,
				1: greenFixture,
				2: blackFixture,
				3: blackFixture,
				4: blackFixture,
				5: blackFixture,
				6: blackFixture,
				7: blackFixture,
			},
		},
		{
			StepNumber: 2,
			KeyStep:    false,
			Fixtures: map[int]common.Fixture{
				0: greenFixture,
				1: greenFixture,
				2: greenFixture,
				3: blackFixture,
				4: blackFixture,
				5: blackFixture,
				6: blackFixture,
				7: blackFixture,
			},
		},
		{
			StepNumber: 3,
			KeyStep:    false,
			Fixtures: map[int]common.Fixture{
				0: greenFixture,
				1: greenFixture,
				2: greenFixture,
				3: greenFixture,
				4: blackFixture,
				5: blackFixture,
				6: blackFixture,
				7: blackFixture,
			},
		},
		{
			StepNumber: 4,
			KeyStep:    false,
			Fixtures: map[int]common.Fixture{
				0: greenFixture,
				1: greenFixture,
				2: greenFixture,
				3: greenFixture,
				4: greenFixture,
				5: blackFixture,
				6: blackFixture,
				7: blackFixture,
			},
		},
		{
			StepNumber: 5,
			KeyStep:    false,
			Fixtures: map[int]common.Fixture{
				0: greenFixture,
				1: greenFixture,
				2: greenFixture,
				3: greenFixture,
				4: greenFixture,
				5: yellowFixture,
				6: blackFixture,
				7: blackFixture,
			},
		},
		{
			StepNumber: 6,
			KeyStep:    false,
			Fixtures: map[int]common.Fixture{
				0: greenFixture,
				1: greenFixture,
				2: greenFixture,
				3: greenFixture,
				4: greenFixture,
				5: yellowFixture,
				6: orangeFixture,
				7: blackFixture,
			},
		},
		{
			StepNumber: 7,
			KeyStep:    false,
			Fixtures: map[int]common.Fixture{
				0: greenFixture,
				1: greenFixture,
				2: greenFixture,
				3: greenFixture,
				4: greenFixture,
				5: yellowFixture,
				6: orangeFixture,
				7: redFixture,
			},
		},
	}

	patternOut := common.Pattern{
		Name:   "VU.Meter",
		Label:  "VU.Meter",
		Number: 7,
	}

	/// We need a step for every step in the chase pattern.
	numberSteps := len(chaseSteps)

	if debug {
		fmt.Printf("Number of Fixtures %d\n", numberOfFixtures)
		fmt.Printf("Number of Groups %d\n", len(groups))
		fmt.Printf("Number of Steps %d\n", numberSteps)
	}

	width := int(math.Floor(float64(numberOfFixtures) / float64(numberSteps)))

	// Populate the steps in the pattern.
	for stepNumber := 0; stepNumber < numberSteps; stepNumber++ {

		// Define the step.
		var newStep common.Step

		// Make space for fixtures in this step.
		fixtures := make(map[int]common.Fixture, numberOfFixtures)

		if debug {
			fmt.Printf("Step number %d\n", stepNumber)
			fmt.Printf("Group number %d Min=%d Max=%d\n", stepNumber, groups[stepNumber].Min, groups[stepNumber].Max)
		}

		if (stepNumber*width) >= groups[stepNumber].Min && (stepNumber*width) <= groups[stepNumber].Max {

			if debug {
				fmt.Printf("\tGroup number %d width %d\n", stepNumber, width)
			}

			// Populate Fixtures for this step.
			for fixtureNumber := 0; fixtureNumber < numberOfFixtures; fixtureNumber++ {
				patternIndex := int(math.Floor(float64(fixtureNumber) / float64(width)))
				if debug {
					fmt.Printf("\tFixture number %d patternIndex %d color %s\n", fixtureNumber, patternIndex, common.GetColorNameByRGB(chaseSteps[stepNumber].Fixtures[patternIndex].Color))
				}
				fixtures[fixtureNumber] = chaseSteps[stepNumber].Fixtures[patternIndex]
			}
		}

		// Now that all the fixtures have been added.
		// Add the completed set of fixtures to the step.
		newStep.Fixtures = fixtures

		// Add the step.
		patternOut.Steps = append(patternOut.Steps, newStep)

	}

	return patternOut
}
