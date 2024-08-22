// Copyright (C) 2022, 2023 , 2024 dhowlett99.
// This is button processor, used by the launchpad and gui interfaces.
// This file processes pattern buttons and controls their actions.
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
)

func selectPattern(sequences []*common.Sequence, X int, Y int, this *CurrentState, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command, updateChannels []chan common.Sequence) {

	this.SelectedSequence = Y

	if this.SelectedMode[this.SelectedSequence] == CHASER_FUNCTION {
		this.TargetSequence = this.ChaserSequenceNumber
		this.DisplaySequence = this.SelectedSequence
		// This is the way we get the shutter chaser to be displayed as we exit
		// the pattern selection
		this.DisplayChaserShortCut = true
		this.SelectedMode[this.DisplaySequence] = CHASER_FUNCTION
	} else {
		this.TargetSequence = this.SelectedSequence
		this.DisplaySequence = this.SelectedSequence
		this.SelectedMode[this.DisplaySequence] = NORMAL
	}

	if debug {
		fmt.Printf("Set Pattern to %d\n", X)
	}

	// Tell the sequence to change the pattern.
	cmd := common.Command{
		Action: common.UpdatePattern,
		Args: []common.Arg{
			{Name: "SelectPattern", Value: X},
		},
	}
	common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

	// Get an upto date copy of the sequence.
	sequences[this.TargetSequence] = common.RefreshSequence(this.TargetSequence, commandChannels, updateChannels)

	// We call ShowPatternSelectionButtons here so the selections will flash as you press them.
	this.EditFixtureSelectionMode = false
	ShowPatternSelectionButtons(this, sequences[this.SelectedSequence].Master, *sequences[this.TargetSequence], this.DisplaySequence, eventsForLaunchpad, guiButtons)

	// Get an upto date copy of the sequence.
	sequences[this.TargetSequence] = common.RefreshSequence(this.TargetSequence, commandChannels, updateChannels)

	// Update the labels.
	showStatusBars(this, sequences, eventsForLaunchpad, guiButtons)
}
