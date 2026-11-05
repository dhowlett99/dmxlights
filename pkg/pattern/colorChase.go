// Copyright (C) 2024 dhowlett99.
// This is the dmxlights color chaser pattern generator.
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
func generateColorChasePattern(numberOfFixtures int) common.Pattern {

	chasePattern := common.Pattern{
		Name:   "Color Chase",
		Label:  "Color.Chase",
		Number: 5,
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
	cyanFixture := common.Fixture{
		MasterDimmer: full,
		Enabled:      true,
		Color:        colors.Cyan,
	}
	blueFixture := common.Fixture{
		MasterDimmer: full,
		Enabled:      true,
		Color:        colors.Blue,
	}
	purpleFixture := common.Fixture{
		MasterDimmer: full,
		Enabled:      true,
		Color:        colors.Purple,
	}
	magentaFixture := common.Fixture{
		MasterDimmer: full,
		Enabled:      true,
		Color:        colors.Magenta,
	}
	blackFixture := common.Fixture{
		MasterDimmer: full,
		Enabled:      true,
		Color:        colors.Black,
	}

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

			if stepNumber == fixtureNumber {
				switch stepCounter {
				case 0:
					fixture = redFixture
				case 1:
					fixture = orangeFixture
				case 2:
					fixture = yellowFixture
				case 3:
					fixture = greenFixture
				case 4:
					fixture = cyanFixture
				case 5:
					fixture = blueFixture
				case 6:
					fixture = purpleFixture
				case 7:
					fixture = magentaFixture
				}
			} else {
				fixture = blackFixture
			}

			// Add the fixture.
			fixtures[fixtureNumber] = fixture
		}

		stepCounter++
		if stepCounter > 7 {
			stepCounter = 0
		}

		// Now that all the fixtures have been added.
		// Add the completed set of fixtures to the step.
		newStep.Fixtures = fixtures

		// Add the step.
		chasePattern.Steps = append(chasePattern.Steps, newStep)

	}

	return chasePattern
}
