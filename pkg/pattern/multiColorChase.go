// Copyright (C) 2024 dhowlett99.
// This is the dmxlights multi color chaser pattern generator.
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

	"github.com/dhowlett99/dmxlights/pkg/colors"
	"github.com/dhowlett99/dmxlights/pkg/common"
)

// Generate dynamic chase pattern based on the number of fixtures.
func generateMultiColorChasePattern(numberOfFixtures int) common.Pattern {

	chasePattern := common.Pattern{
		Name:   "Multi Color",
		Label:  "Multi.Color",
		Number: 6,
		Steps:  []common.Step{},
	}

	// We need a step for every fixture in the chase.
	numberSteps := numberOfFixtures

	// Define colors.
	var chaseFixtures []common.Fixture

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

	// Add colors to an array.
	chaseFixtures = append(chaseFixtures, redFixture)
	chaseFixtures = append(chaseFixtures, orangeFixture)
	chaseFixtures = append(chaseFixtures, yellowFixture)
	chaseFixtures = append(chaseFixtures, greenFixture)
	chaseFixtures = append(chaseFixtures, cyanFixture)
	chaseFixtures = append(chaseFixtures, blueFixture)
	chaseFixtures = append(chaseFixtures, purpleFixture)
	chaseFixtures = append(chaseFixtures, magentaFixture)

	var colorCounter int

	// Populate the steps in the pattern.
	for stepNumber := 0; stepNumber < numberSteps; stepNumber++ {

		// Define the step.
		var newStep common.Step
		var index int

		// Make space for fixtures in this step.
		fixtures := make(map[int]common.Fixture, numberOfFixtures)

		fmt.Printf("Step %d \n", stepNumber)

		// Populate Fixtures for this step.
		for fixtureNumber := 0; fixtureNumber < numberOfFixtures; fixtureNumber++ {
			index = colorCounter + fixtureNumber
			if index > 7 {
				index = index - numberOfFixtures
			}
			fixtures[fixtureNumber] = chaseFixtures[index]
		}

		// Now that all the fixtures have been added.
		// Add the completed set of fixtures to the step.
		newStep.Fixtures = fixtures

		colorCounter++
		if colorCounter > 7 {
			colorCounter = 0
		}

		// Add the step.
		chasePattern.Steps = append(chasePattern.Steps, newStep)

	}

	return chasePattern
}
