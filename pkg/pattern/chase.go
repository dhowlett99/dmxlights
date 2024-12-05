// Copyright (C) 2024 dhowlett99.
// This is the dmxlights chase pattern generator.
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

// Generate dynamic chase pattern based on the number of fixtures.
func generateChasePattern(numberOfFixtures int) common.Pattern {

	chasePattern := common.Pattern{
		Name:   "Chase",
		Number: 0,
		Label:  "Std.Chase",
		Steps:  []common.Step{},
	}

	// We need a step for every fixture in the chase.
	numberSteps := numberOfFixtures

	// Define on state for chase.
	onFixture := common.Fixture{
		MasterDimmer: full,
		Enabled:      true,
		Shutter:      255,
		Brightness:   255,
		Color:        colors.Green,
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

		// Make space for fixtures in this step.
		fixtures := make(map[int]common.Fixture, numberOfFixtures)

		// Define the fixture.
		var fixture common.Fixture

		// Populate Fixtures for this step.
		for fixtureNumber := 0; fixtureNumber < numberOfFixtures; fixtureNumber++ {

			if stepNumber == fixtureNumber {
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
		chasePattern.Steps = append(chasePattern.Steps, newStep)

	}

	return chasePattern
}
