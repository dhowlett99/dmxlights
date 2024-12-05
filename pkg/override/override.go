// Copyright (C) 2022, 2023 , 2024 dhowlett99.
// This is the override tools package.
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

package override

import (
	"fmt"

	"github.com/dhowlett99/dmxlights/pkg/commands"
	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
)

const debug = false

func UpdateOverrides(sequenceNumber int, fixturesConfig *fixture.Fixtures, switchOverrides *[][]common.Override) {

	// Store the switch Config locally.
	switchConfig := commands.LoadSwitchConfiguration(sequenceNumber, fixturesConfig)

	// Copy in the overrides.
	overrides := *switchOverrides

	// Populate each switch with a number of states based on their config.
	for swiTchNumber := 0; swiTchNumber < len(switchConfig); swiTchNumber++ {

		// assign the switch.
		swiTch := switchConfig[swiTchNumber]

		// Now populate the states.
		for stateNumber := 0; stateNumber < len(swiTch.States); stateNumber++ {

			state := swiTch.States[stateNumber]

			// Find the details of the fixture for this switch.
			thisFixture, err := fixture.FindFixtureByLabel(swiTch.UseFixture, fixturesConfig)
			if err != nil {
				fmt.Printf("error %s\n", err.Error())
			}

			// Load the config for this state of of this switch
			override := fixture.DiscoverSwitchOveride(thisFixture, swiTch.Number, int(state.Number), fixturesConfig)

			// Assign this discovered override to the current switch state.
			overrides[swiTchNumber] = append(overrides[swiTchNumber], override)

			if debug {
				fmt.Printf("Setting Up Override for Switch No=%d Name=%s State No=%d Name=%s\n", swiTch.Number, swiTch.Name, state.Number, state.Name)
				fmt.Printf("\t Override Colors %+v\n", override.Colors)
			}
		}
	}

	if debug {
		for overrideNumber, override := range overrides {
			fmt.Printf("overrideNumber %d\n", overrideNumber)
			for stateNumber, state := range override {
				fmt.Printf("\tstateNumber %d state %+v\n", stateNumber, state)
			}
		}
	}

	// re-instate the pointer to the overrides.
	switchOverrides = &overrides
}
