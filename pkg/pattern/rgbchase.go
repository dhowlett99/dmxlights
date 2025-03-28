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

// Generate dynamic RGB chase pattern based on the number of fixtures.
func generateRgbChasePattern(numberOfFixtures int) common.Pattern {

	chasePattern := common.Pattern{
		Name:   "RGB Chase",
		Number: 2,
		Label:  "RGB.Chase",
		Steps:  []common.Step{},
	}

	// We need a step for each color in the chase.
	numberSteps := 3

	// Define red state for chase.
	redFixture := common.Fixture{
		MasterDimmer: full,
		Enabled:      true,
		Color:        colors.Red,
	}

	// Define green state for chase.
	greenFixture := common.Fixture{
		MasterDimmer: full,
		Enabled:      true,
		Color:        colors.Green,
	}

	// Define blue state for chase.
	blueFixture := common.Fixture{
		MasterDimmer: full,
		Enabled:      true,
		Color:        colors.Blue,
	}

	// We use a step counter which is reset every three step, to give the
	// red, green, blue,
	var stepCounter int

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

			if stepCounter == 0 {
				fixture = redFixture
			}
			if stepCounter == 1 {
				fixture = greenFixture
			}
			if stepCounter == 2 {
				fixture = blueFixture
			}

			// Add the fixture.
			fixtures[fixtureNumber] = fixture
		}

		stepCounter++

		// Now that all the fixtures have been added.
		// Add the completed set of fixtures to the step.
		newStep.Fixtures = fixtures

		// Add the step.
		chasePattern.Steps = append(chasePattern.Steps, newStep)

		if stepCounter > 2 {
			stepCounter = 0
		}
	}

	return chasePattern
}
