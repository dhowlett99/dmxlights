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
	"math"

	"github.com/dhowlett99/dmxlights/pkg/colors"
	"github.com/dhowlett99/dmxlights/pkg/common"
)

type MaxMins struct {
	Max int
	Min int
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

	return groups
}

// Generate dynamic chase pattern based on the number of fixtures.
func generateVuChasePattern(numberOfFixtures int) common.Pattern {

	// Divide up the fixtures into 8 groups used by this pattern.
	// groups := numberOfFixtures / 8
	// maxGroup := groups * numberOfFixtures
	// fmt.Printf("Number of Groups %d MaxGroup %d\n", groups, maxGroup)
	patterSize := 8
	groups := getGroups(patterSize, numberOfFixtures)

	chasePattern := common.Pattern{
		Name:   "VU.Meter",
		Label:  "VU.Meter",
		Number: 7,
		Steps:  []common.Step{},
	}

	// We need a step for every fixture in the chase.
	numberSteps := numberOfFixtures

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

	// Populate the steps in the pattern.
	for stepNumber := 0; stepNumber < numberSteps; stepNumber++ {

		// Define the step.
		var newStep common.Step

		// Make space for fixtures in this step.
		fixtures := make(map[int]common.Fixture, numberOfFixtures)

		// Populate Fixtures for this step.
		if stepNumber >= groups[0].Min && stepNumber <= groups[0].Max {
			fixtures[0] = greenFixture
			fixtures[1] = blackFixture
			fixtures[2] = blackFixture
			fixtures[3] = blackFixture
			fixtures[4] = blackFixture
			fixtures[5] = blackFixture
			fixtures[6] = blackFixture
			fixtures[7] = blackFixture
		}

		if stepNumber >= groups[1].Min && stepNumber <= groups[1].Max {
			fixtures[0] = greenFixture
			fixtures[1] = greenFixture
			fixtures[2] = blackFixture
			fixtures[3] = blackFixture
			fixtures[4] = blackFixture
			fixtures[5] = blackFixture
			fixtures[6] = blackFixture
			fixtures[7] = blackFixture
		}

		if stepNumber >= groups[2].Min && stepNumber <= groups[2].Max {
			fixtures[0] = greenFixture
			fixtures[1] = greenFixture
			fixtures[2] = greenFixture
			fixtures[3] = blackFixture
			fixtures[4] = blackFixture
			fixtures[5] = blackFixture
			fixtures[6] = blackFixture
			fixtures[7] = blackFixture
		}

		if stepNumber >= groups[3].Min && stepNumber <= groups[3].Max {
			fixtures[0] = greenFixture
			fixtures[1] = greenFixture
			fixtures[2] = greenFixture
			fixtures[3] = greenFixture
			fixtures[4] = blackFixture
			fixtures[5] = blackFixture
			fixtures[6] = blackFixture
			fixtures[7] = blackFixture
		}

		if stepNumber >= groups[4].Min && stepNumber <= groups[4].Max {
			fixtures[0] = greenFixture
			fixtures[1] = greenFixture
			fixtures[2] = greenFixture
			fixtures[3] = greenFixture
			fixtures[4] = greenFixture
			fixtures[5] = blackFixture
			fixtures[6] = blackFixture
			fixtures[7] = blackFixture
		}

		if stepNumber >= groups[5].Min && stepNumber <= groups[5].Max {
			fixtures[0] = greenFixture
			fixtures[1] = greenFixture
			fixtures[2] = greenFixture
			fixtures[3] = greenFixture
			fixtures[4] = greenFixture
			fixtures[5] = yellowFixture
			fixtures[6] = blackFixture
			fixtures[7] = blackFixture
		}

		if stepNumber >= groups[6].Min && stepNumber <= groups[6].Max {
			fixtures[0] = greenFixture
			fixtures[1] = greenFixture
			fixtures[2] = greenFixture
			fixtures[3] = greenFixture
			fixtures[4] = greenFixture
			fixtures[5] = yellowFixture
			fixtures[6] = orangeFixture
			fixtures[7] = blackFixture
		}

		if stepNumber >= groups[7].Min && stepNumber <= groups[7].Max {
			fixtures[0] = greenFixture
			fixtures[1] = greenFixture
			fixtures[2] = greenFixture
			fixtures[3] = greenFixture
			fixtures[4] = greenFixture
			fixtures[5] = yellowFixture
			fixtures[6] = orangeFixture
			fixtures[7] = redFixture
		}

		// Now that all the fixtures have been added.
		// Add the completed set of fixtures to the step.
		newStep.Fixtures = fixtures

		// Add the step.
		chasePattern.Steps = append(chasePattern.Steps, newStep)
	}

	return chasePattern
}
