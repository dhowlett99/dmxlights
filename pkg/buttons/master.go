// Copyright (C) 2022, 2023 , 2024 dhowlett99.
// This is button processor, used by the launchpad and gui interfaces.
// This file processes master brightness buttons and controls their actions.
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

	"github.com/dhowlett99/dmxlights/pkg/colors"
	"github.com/dhowlett99/dmxlights/pkg/common"
)

func decraseBrightness(X int, Y int, this *CurrentState, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command) {

	if debug {
		fmt.Printf("Brightness Down Sequence %d Switch %d\n", this.SelectedSequence, this.SelectedSwitch)
	}

	buttonTouched(common.Button{X: X, Y: Y}, colors.White, colors.Cyan, eventsForLaunchpad, guiButtons)

	// pull overrides.
	overrides := *this.SwitchOverrides

	// find the psudo sequence number.
	sequence := getSelectedSequenceNumber(this.TargetSequence, this.SelectedType, this.SelectedSwitch)

	// Decrement the brightness.
	this.MasterBrightness[sequence] = this.MasterBrightness[sequence] - 10
	if this.MasterBrightness[sequence] < 0 {
		this.MasterBrightness[sequence] = 0
	}

	// Now look the mode to find switch action that is a mini sequencer in chase mode.
	if overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Mode == "Chase" &&
		overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].ChaseRunning {

		// We have found a mini-sequencer configured as a chaser.
		// Send an override command to change the master brightness for this fixture.
		cmd := common.Command{
			Action: common.OverrideMaster,
			Args: []common.Arg{
				{Name: "SwitchNumber", Value: this.SelectedSwitch},
				{Name: "StateNumber", Value: this.SwitchPosition[this.SelectedSwitch]},
				{Name: "Master", Value: this.MasterBrightness[sequence]},
			},
		}
		common.SendCommandToSequence(this.SwitchSequenceNumber, cmd, commandChannels)

	} else {

		cmd := common.Command{
			Action: common.Master,
			Args: []common.Arg{
				{Name: "Master", Value: this.MasterBrightness[sequence]},
				{Name: "SwitchNumber", Value: this.SelectedSwitch},
			},
		}
		common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

	}
	// Update the status bar
	common.UpdateStatusBar(fmt.Sprintf("Master %02d", this.MasterBrightness[sequence]), "master", false, guiButtons)

}

func increaseBrightness(X int, Y int, this *CurrentState, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command) {

	if debug {
		fmt.Printf("Brightness Up Sequence %d Switch %d\n", this.SelectedSequence, this.SelectedSwitch)
	}

	buttonTouched(common.Button{X: X, Y: Y}, colors.White, colors.Cyan, eventsForLaunchpad, guiButtons)

	// pull overrides.
	overrides := *this.SwitchOverrides

	// find the psudo sequence number.
	sequence := getSelectedSequenceNumber(this.TargetSequence, this.SelectedType, this.SelectedSwitch)

	// Increment the brightness.
	this.MasterBrightness[sequence] = this.MasterBrightness[sequence] + 10
	if this.MasterBrightness[sequence] > common.MAX_DMX_BRIGHTNESS {
		this.MasterBrightness[sequence] = common.MAX_DMX_BRIGHTNESS
	}

	// Now look the mode to find switch action that is a mini sequencer in chase mode.
	if overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Mode == "Chase" &&
		overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].ChaseRunning {

		// We have found a mini-sequencer configured as a chaser.
		// Send an override command to change the master brightness for this fixture.
		cmd := common.Command{
			Action: common.OverrideMaster,
			Args: []common.Arg{
				{Name: "SwitchNumber", Value: this.SelectedSwitch},
				{Name: "StateNumber", Value: this.SwitchPosition[this.SelectedSwitch]},
				{Name: "Master", Value: this.MasterBrightness[sequence]},
			},
		}
		common.SendCommandToSequence(this.SwitchSequenceNumber, cmd, commandChannels)

	} else {

		cmd := common.Command{
			Action: common.Master,
			Args: []common.Arg{
				{Name: "Master", Value: this.MasterBrightness[sequence]},
				{Name: "SwitchNumber", Value: this.SelectedSwitch},
			},
		}
		common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

	}

	// Update the status bar
	common.UpdateStatusBar(fmt.Sprintf("Master %02d", this.MasterBrightness[sequence]), "master", false, guiButtons)

}

func getSelectedSequenceNumber(selectedSequence int, selectedType string, selectedSwitch int) int {

	if selectedType == "switch" {
		return selectedSwitch + 5
	}

	return selectedSequence
}
