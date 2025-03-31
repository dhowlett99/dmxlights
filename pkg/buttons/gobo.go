// Copyright (C) 2022, 2023 , 2024 dhowlett99.
// This is button processor, used by the launchpad and gui interfaces.
// This file processes gobo buttons and controls their actions.
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
	"github.com/oliread/usbdmx/ft232"
)

func selectGobo(sequences []*common.Sequence, X int, Y int, this *CurrentState, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command, updateChannels []chan common.Sequence, fixturesConfig *fixture.Fixtures, dmxController *ft232.DMXController) {

	if debug {
		fmt.Printf("Sequence %d Fixture %d Set Gobo %d\n", Y, this.SelectedFixture, this.SelectedGobo)
	}

	this.SelectedGobo = X + 1

	// Set the selected gobo for this sequence.
	cmd := common.Command{
		Action: common.UpdateGobo,
		Args: []common.Arg{
			{Name: "SelectedGobo", Value: this.SelectedGobo},
			{Name: "FixtureNumber", Value: this.SelectedFixture},
		},
	}
	common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

	// If configured set scanner color in chaser.
	if this.ScannerChaser[this.SelectedSequence] && this.SelectedType == "scanner" {
		cmd := common.Command{
			Action: common.UpdateGobo,
			Args: []common.Arg{
				{Name: "SelectedGobo", Value: this.SelectedGobo},
				{Name: "FixtureNumber", Value: this.SelectedFixture},
			},
		}
		common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)
	}

	this.EditGoboSelectionMode = true

	// Get an upto date copy of the sequence.
	sequences[this.SelectedSequence] = common.RefreshSequence(this.SelectedSequence, commandChannels, updateChannels)

	// If the sequence isn't running this will force a single gobo DMX message.
	fixture.MapFixturesGoboOnly(this.SelectedSequence, this.SelectedFixture, this.SelectedGobo, fixturesConfig, dmxController, this.DmxInterfacePresent)

	// Clear the pattern function keys
	common.ClearSelectedRowOfButtons(this.SelectedSequence, eventsForLaunchpad, guiButtons)

	// We call ShowGoboSelectionButtons here so the selections will flash as you press them.
	ShowGoboSelectionButtons(*sequences[this.SelectedSequence], this, eventsForLaunchpad, guiButtons)
}
