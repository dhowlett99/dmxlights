// Copyright (C) 2022, 2023 , 2024 dhowlett99.
// This is button processor, used by the launchpad and gui interfaces.
// This file processes fixture buttons and controls their actions.
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

package buttons

import (
	"fmt"

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
)

func selectFixture(sequences []*common.Sequence, X int, Y int, this *CurrentState, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command, fixturesConfig *fixture.Fixtures) {

	this.SelectedFixture = X

	if debug {
		fmt.Printf("Sequence %d Selected Fixture is %d \n", Y, this.SelectedFixture)
	}

	// Clear the pattern function keys
	common.ClearSelectedRowOfButtons(this.SelectedSequence, eventsForLaunchpad, guiButtons)

	// Update the buttons.
	if this.FollowingAction == "ShowGoboSelectionButtons" {
		ShowGoboSelectionButtons(*sequences[this.SelectedSequence], this, eventsForLaunchpad, guiButtons)
	}
	if this.FollowingAction == "ShowScannerColorSelectionButtons" {
		err := ShowScannerColorSelectionButtons(*sequences[this.SelectedSequence], this, eventsForLaunchpad, fixturesConfig, guiButtons)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			common.RevealSequence(this.SelectedSequence, commandChannels)
		}
	}
	this.EditFixtureSelectionMode = false
}
