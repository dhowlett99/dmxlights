// Copyright (C) 2022, 2023 , 2024 dhowlett99.
// This is button processor, used by the launchpad and gui interfaces.
// This file processes the select buttons and controls their actions.
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

func selectSequence(sequences []*common.Sequence, Y int, this *CurrentState, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command) {

	this.SelectedSequence = Y
	this.SelectedType = sequences[this.SelectedSequence].Type

	if this.ScannerChaser[this.SelectedSequence] {
		this.EditWhichStaticSequence = this.ChaserSequenceNumber
	} else {
		this.EditWhichStaticSequence = this.SelectedSequence
	}

	if debug {
		fmt.Printf("Select Sequence %d Type %s\n", this.SelectedSequence, this.SelectedType)
	}

	// // If we're in shutter chase mode
	if this.SelectedMode[this.SelectedSequence] == CHASER_FUNCTION || this.SelectedMode[this.SelectedSequence] == CHASER_DISPLAY {
		this.TargetSequence = this.ChaserSequenceNumber
	} else {
		this.TargetSequence = this.SelectedSequence
	}

	deFocusAllSwitches(this, sequences, commandChannels)
	HandleSelect(sequences, this, eventsForLaunchpad, commandChannels, guiButtons)

	this.ShowRGBColorPicker = false
	this.EditGoboSelectionMode = false
	this.DisplayChaserShortCut = false

}
