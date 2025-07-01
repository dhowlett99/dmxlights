// Copyright (C) 2022, 2023 , 2024 dhowlett99.
// This is button processor, used by the launchpad and gui interfaces.
// This file processes the size buttons and controls their actions.
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

func decreaseSizeRGB(this *CurrentState, commandChannels []chan common.Command) {

	if debug {
		fmt.Printf("Decrease Size\n")
	}

	// Deal with the RGB sequence. Decrease Size.

	// Decrement RGB Size.
	this.RGBSize[this.TargetSequence]--
	if this.RGBSize[this.TargetSequence] < common.MIN_RGB_SIZE {
		this.RGBSize[this.TargetSequence] = common.MIN_RGB_SIZE
	}

	// Send Update RGB Size.
	cmd := common.Command{
		Action: common.UpdateRGBSize,
		Args: []common.Arg{
			{Name: "RGBSize", Value: this.RGBSize[this.TargetSequence]},
		},
	}
	common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)
}

func decreaseSizeScanner(this *CurrentState, commandChannels []chan common.Command) {

	// Deal with Scanner sequence. Decrease Size.

	// Send Update Scanner Size.
	this.ScannerSize[this.TargetSequence] = this.ScannerSize[this.TargetSequence] - 10
	if this.ScannerSize[this.TargetSequence] < 0 {
		this.ScannerSize[this.TargetSequence] = 0
	}

	// Send Update Scanner Size.
	cmd := common.Command{
		Action: common.UpdateScannerSize,
		Args: []common.Arg{
			{Name: "ScannerSize", Value: this.ScannerSize[this.TargetSequence]},
		},
	}
	common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

}

func decreaseOverrideSize(this *CurrentState, commandChannels []chan common.Command) {

	// Decrement the switch size.
	overrides := *this.SwitchOverrides
	overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Size--
	if overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Size < common.MIN_RGB_SIZE {
		overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Size = common.MIN_RGB_SIZE
	}
	this.SwitchOverrides = &overrides

	// Send a message to override / increase the selected switch shift.
	cmd := common.Command{
		Action: common.OverrideSize,
		Args: []common.Arg{
			{Name: "SwitchNumber", Value: this.SelectedSwitch},
			{Name: "SwitchPosition", Value: this.SwitchPosition[this.SelectedSwitch]},
			{Name: "Shift", Value: overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Size},
		},
	}
	common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

}

// Deal with the switch sequence that has a projector fixture.
func decreaseOverrideColor(this *CurrentState, commandChannels []chan common.Command) {

	// Pull overrides.
	overrides := *this.SwitchOverrides

	// Decrement the switch color.
	overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Color--
	if overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Color < common.MIN_PROJECTOR_COLOR {
		overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Color = common.MIN_PROJECTOR_COLOR
	}
	this.SwitchOverrides = &overrides

	// Send a message to override / increase the selected switch shift.
	cmd := common.Command{
		Action: common.OverrideColor,
		Args: []common.Arg{
			{Name: "SwitchNumber", Value: this.SelectedSwitch},
			{Name: "SwitchPosition", Value: this.SwitchPosition[this.SelectedSwitch]},
			{Name: "Color", Value: overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Color},
			{Name: "Available Color Names", Value: overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].AvailableColors},
		},
	}
	common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

	// Push the switch overrides pointer back into the current state.
	this.SwitchOverrides = &overrides
}

func increaseSizeRGB(this *CurrentState, commandChannels []chan common.Command) {
	if debug {
		fmt.Printf("Increase Size\n")
	}

	// Send Update RGB Size.
	this.RGBSize[this.TargetSequence]++
	if this.RGBSize[this.TargetSequence] > common.MAX_RGB_SIZE {
		this.RGBSize[this.TargetSequence] = common.MAX_RGB_SIZE
	}

	// Send a message to the RGB sequence.
	cmd := common.Command{
		Action: common.UpdateRGBSize,
		Args: []common.Arg{
			{Name: "RGBSize", Value: this.RGBSize[this.TargetSequence]},
		},
	}
	common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

}

// Deal with the Scanner size.
func increaseSizeScanner(this *CurrentState, commandChannels []chan common.Command) {

	// Increment the scanner size.
	this.ScannerSize[this.TargetSequence] = this.ScannerSize[this.TargetSequence] + 10
	if this.ScannerSize[this.TargetSequence] > common.MAX_SCANNER_SIZE {
		this.ScannerSize[this.TargetSequence] = common.MAX_SCANNER_SIZE
	}

	// Send Update Scanner Size.
	cmd := common.Command{
		Action: common.UpdateScannerSize,
		Args: []common.Arg{
			{Name: "ScannerSize", Value: this.ScannerSize[this.TargetSequence]},
		},
	}
	common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

}

func increaseOverrideSize(this *CurrentState, commandChannels []chan common.Command) {

	if this.SelectedType == "switch" && this.SelectedFixtureType == "rgb" {

		// Increase the switch size.
		overrides := *this.SwitchOverrides
		overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Size++
		if overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Size > common.MAX_RGB_SHIFT {
			overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Size = common.MAX_RGB_SHIFT
		}
		this.SwitchOverrides = &overrides

		// Send a message to override / increase the selected switch shift.
		cmd := common.Command{
			Action: common.OverrideSize,
			Args: []common.Arg{
				{Name: "SwitchNumber", Value: this.SelectedSwitch},
				{Name: "SwitchPosition", Value: this.SwitchPosition[this.SelectedSwitch]},
				{Name: "Shift", Value: overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Size},
			},
		}
		common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

	}
}

func increaseOverrideColor(this *CurrentState, commandChannels []chan common.Command) {

	//  Pull overrides.
	overrides := *this.SwitchOverrides

	//  Increase the switch color.
	overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Color++
	if overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Color > overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].MaxColors {
		overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Color = overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].MaxColors
	}
	this.SwitchOverrides = &overrides

	// Send a message to override / increase the selected switch color.
	cmd := common.Command{
		Action: common.OverrideColor,
		Args: []common.Arg{
			{Name: "SwitchNumber", Value: this.SelectedSwitch},
			{Name: "SwitchPosition", Value: this.SwitchPosition[this.SelectedSwitch]},
			{Name: "Color", Value: overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Color},
			{Name: "Available Color Names", Value: overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].AvailableColors},
		},
	}
	common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

	// Push the switch overrides pointer back into the current state.
	this.SwitchOverrides = &overrides

}
