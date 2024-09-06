// Copyright (C) 2022, 2023 , 2024 dhowlett99.
// This is button processor, used by the launchpad and gui interfaces.
// This file processes the switch buttons and controls their actions.
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

func selectSwitch(sequences []*common.Sequence, X int, Y int, this *CurrentState, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command, updateChannels []chan common.Sequence, fixturesConfig *fixture.Fixtures) {

	this.SelectedSequence = Y
	this.SelectedSwitch = X
	this.SelectedType = "switch"
	this.SelectedFixtureType = fixture.GetSwitchFixtureType(this.SelectedSwitch, int16(this.SwitchPosition[this.SelectedSwitch]), fixturesConfig)

	if debug {
		fmt.Printf("Switch Key X:%d Y:%d\n", this.SelectedSwitch, this.SelectedSequence)
	}

	if this.ShowRGBColorPicker {
		this.ShowRGBColorPicker = false
		removeColorPicker(this, sequences, eventsForLaunchpad, guiButtons, commandChannels)
	}

	// Get an upto date copy of the switch information by updating our copy of the switch sequence.
	sequences[this.SelectedSequence] = common.RefreshSequence(this.SelectedSequence, commandChannels, updateChannels)

	// We have a valid switch.
	if this.SelectedSwitch < len(sequences[this.SelectedSequence].Switches) {

		// Second time we've pressed this switch button, actually step the state.
		if this.SelectedSwitch == this.LastSelectedSwitch {
			this.SwitchPosition[this.SelectedSwitch] = this.SwitchPosition[this.SelectedSwitch] + 1
			valuesLength := len(sequences[this.SelectedSequence].Switches[this.SelectedSwitch].States)
			if this.SwitchPosition[this.SelectedSwitch] == valuesLength {
				this.SwitchPosition[this.SelectedSwitch] = 0
			}
			// Send a message to the sequence for it to step to the next state the selected switch.
			cmd := common.Command{
				Action: common.UpdateSwitch,
				Args: []common.Arg{
					{Name: "SwitchNumber", Value: this.SelectedSwitch},
					{Name: "SwitchPosition", Value: this.SwitchPosition[this.SelectedSwitch]},
					{Name: "Step", Value: true},  // Step the switch state.
					{Name: "Focus", Value: true}, // Focus the switch lamp.
				},
			}
			// Send a message to the switch sequence.
			common.SendCommandToAllSequenceOfType(sequences, cmd, commandChannels, "switch")
		} else {
			// Just send a message to focus the switch button.
			cmd := common.Command{
				Action: common.UpdateSwitch,
				Args: []common.Arg{
					{Name: "SwitchNumber", Value: this.SelectedSwitch},
					{Name: "SwitchPosition", Value: this.SwitchPosition[this.SelectedSwitch]},
					{Name: "Step", Value: false}, // Don't step the switch state.
					{Name: "Focus", Value: true}, // Focus the switch lamp.
				},
			}
			// Send a message to the switch sequence.
			common.SendCommandToAllSequenceOfType(sequences, cmd, commandChannels, "switch")

		}

	}

	// Light the correct selected switch.
	this.SelectedSequence = this.SwitchSequenceNumber
	this.LastSelectedSwitch = this.SelectedSwitch

	// Use the default behaviour of SelectSequence to turn of the other sequence select buttons.
	SelectSequence(this)

	// Find out if this switch state has a music trigger.
	this.MusicTrigger = fixture.GetSwitchStateIsMusicTriggerOn(this.SelectedSwitch, int16(this.SwitchPosition[this.SelectedSwitch]), fixturesConfig)

	// Switch overrides will get displayed here as well.
	UpdateSpeed(this, guiButtons)
	UpdateShift(this, guiButtons)
	UpdateSize(this, guiButtons)
	UpdateFade(this, guiButtons)

	// Update the labels.
	showStatusBars(this, sequences, eventsForLaunchpad, guiButtons)

	// Light the sequence selector button.
	lightSelectedButton(eventsForLaunchpad, guiButtons, this)

}
