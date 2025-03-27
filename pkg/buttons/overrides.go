// Copyright (C) 2022, 2023, 2024, 2025 dhowlett99.
// This implements the load preset feature, used by the buttons package.
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
//

package buttons

import (
	"fmt"

	"github.com/dhowlett99/dmxlights/pkg/common"
)

func RefreshLocalOverrides(this *CurrentState, sequence *common.Sequence) {

	// Pull the switch overrides pointer from the current state.
	overrides := *this.SwitchOverrides

	// Now set our local representation for all switches and states.
	for swiTchNumber, swiTch := range sequence.Switches {

		this.SwitchPosition[swiTchNumber] = swiTch.CurrentPosition

		for stateNumber := range swiTch.States {

			//  Restore any switch Overrides.
			if sequence.Switches[swiTchNumber].Override.Speed != 0 {
				overrides[swiTchNumber][stateNumber].Speed = sequence.Switches[swiTchNumber].Override.Speed
			}
			if sequence.Switches[swiTchNumber].Override.Shift != 0 {
				overrides[swiTchNumber][stateNumber].Shift = sequence.Switches[swiTchNumber].Override.Shift
			}
			if sequence.Switches[swiTchNumber].Override.Size != 0 {
				overrides[swiTchNumber][stateNumber].Size = sequence.Switches[swiTchNumber].Override.Size
			}
			if sequence.Switches[swiTchNumber].Override.Fade != 0 {
				overrides[swiTchNumber][stateNumber].Fade = sequence.Switches[swiTchNumber].Override.Fade
			}

			if sequence.Switches[swiTchNumber].Override.Rotate != 0 {
				overrides[swiTchNumber][stateNumber].Rotate = sequence.Switches[swiTchNumber].Override.Rotate
			}
			if sequence.Switches[swiTchNumber].Override.AvailableColors != nil {
				overrides[swiTchNumber][stateNumber].AvailableColors = sequence.Switches[swiTchNumber].Override.AvailableColors
			}
			if sequence.Switches[swiTchNumber].Override.Color != 0 {
				overrides[swiTchNumber][stateNumber].Color = sequence.Switches[swiTchNumber].Override.Color
			}
			if sequence.Switches[swiTchNumber].Override.Gobo != 0 {
				overrides[swiTchNumber][stateNumber].Gobo = sequence.Switches[swiTchNumber].Override.Gobo
			}
		}

		if debug {
			var stateNames []string
			for _, state := range swiTch.States {
				stateNames = append(stateNames, state.Name)
			}
			fmt.Printf("restoring switch number %d to postion %d states[%s]\n", swiTchNumber, this.SwitchPosition[swiTchNumber], stateNames)
		}
	}

	// Push the switch overrides pointer back into the current state.
	this.SwitchOverrides = &overrides

}
