// Copyright (C) 2024 dhowlett99.
// This is the dmxlights flash pattern generator.
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
	"github.com/dhowlett99/dmxlights/pkg/colors"
	"github.com/dhowlett99/dmxlights/pkg/common"
)

// Generate dynamic flash pattern based on the number of fixtures.
func generateFlashPattern(numberOfFixtures int) common.Pattern {

	flashPattern := common.Pattern{
		Name:   "Flash",
		Number: 1,
		Label:  "Flash",
		Steps:  []common.Step{},
	}

	// We need a step for on and a step for off in the flash pattern so 2.
	numberSteps := 2

	// Define on state for chase.
	onFixture := common.Fixture{
		MasterDimmer: full,
		Enabled:      true,
		Color:        colors.White,
	}

	// Define off state for chase.
	offFixture := common.Fixture{
		MasterDimmer: full,
		Enabled:      true,
		Color:        colors.Black,
	}

	// Populate the steps in the pattern.
	for stepNumber := 0; stepNumber < numberSteps; stepNumber++ {

		// Define the step.
		var newStep common.Step
		newStep.KeyStep = true

		// Make space for fixtures in this step.
		fixtures := make(map[int]common.Fixture, numberOfFixtures)

		// Define the fixture.
		var fixture common.Fixture

		// Populate Fixtures for this step.
		for fixtureNumber := 0; fixtureNumber < numberOfFixtures; fixtureNumber++ {

			if stepNumber == 0 {
				fixture = onFixture
			} else {
				fixture = offFixture
			}

			// Add the fixture.
			fixtures[fixtureNumber] = fixture
		}

		// Now that all the fixtures have been added.
		// Add the completed set of fixtures to the step.
		newStep.Fixtures = fixtures

		// Add the step.
		flashPattern.Steps = append(flashPattern.Steps, newStep)

	}

	return flashPattern
}
