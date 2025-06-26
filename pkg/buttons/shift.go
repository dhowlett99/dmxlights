// Copyright (C) 2022, 2023 , 2024 dhowlett99.
// This is button processor, used by the launchpad and gui interfaces.
// This file processes the shift buttons and controls their actions.
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

func decreaseShiftRGB(this *CurrentState, commandChannels []chan common.Command) {

	if debug {
		fmt.Printf("Decrease Shift\n")
	}

	// Decrement the RGB Shift.
	this.RGBShift[this.TargetSequence]--
	if this.RGBShift[this.TargetSequence] < common.MIN_RGB_SHIFT {
		this.RGBShift[this.TargetSequence] = common.MIN_RGB_SHIFT
	}

	// Send a message to the RGB sequence.
	cmd := common.Command{
		Action: common.UpdateRGBShift,
		Args: []common.Arg{
			{Name: "RGBShift", Value: this.RGBShift[this.TargetSequence]},
		},
	}
	common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

}

// Deal with an Scanner sequence.
func decreaseShiftScanner(this *CurrentState, commandChannels []chan common.Command) {

	// Decrement the Scanner Shift.
	this.ScannerShift[this.TargetSequence]--
	if this.ScannerShift[this.TargetSequence] < common.MIN_SCANNER_SHIFT {
		this.ScannerShift[this.TargetSequence] = common.MIN_SCANNER_SHIFT
	}

	// Send a message to the Scanner sequence.
	cmd := common.Command{
		Action: common.UpdateScannerShift,
		Args: []common.Arg{
			{Name: "ScannerShift", Value: this.ScannerShift[this.TargetSequence]},
		},
	}
	common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

}

// func decreaseShiftSwitchRGB(this *CurrentState, commandChannels []chan common.Command) {

// 	// Deal with an RGB Switch sequence.
// 	if this.SelectedType == "switch" && this.SelectedFixtureType == "rgb" {

// 		// Decrement the Switch Shift.
// 		overrides := *this.SwitchOverrides
// 		overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Shift = overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Shift - 1
// 		if overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Shift < 0 {
// 			overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Shift = 0
// 		}
// 		this.SwitchOverrides = &overrides

// 		// Send a message to override / increase the selected switch shift.
// 		cmd := common.Command{
// 			Action: common.OverrideShift,
// 			Args: []common.Arg{
// 				{Name: "SwitchNumber", Value: this.SelectedSwitch},
// 				{Name: "SwitchPosition", Value: this.SwitchPosition[this.SelectedSwitch]},
// 				{Name: "Shift", Value: overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Shift},
// 			},
// 		}
// 		common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

// 	}
// }

func decreaseRotateSpeed(this *CurrentState, commandChannels []chan common.Command) {

	// Pull the overrides
	overrides := *this.SwitchOverrides

	switchPosition := this.SwitchPosition[this.SelectedSwitch]
	IsRotateOverrideAble := overrides[this.SelectedSwitch][switchPosition].IsRotateOverrideAble
	actionMode := overrides[this.SelectedSwitch][switchPosition].Mode

	if IsRotateOverrideAble && actionMode != "Control" {

		// Decrement the Rotate Speed.
		overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Rotate--
		if overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Rotate < common.MIN_PROJECTOR_ROTATE_SPEED {
			overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Rotate = common.MIN_PROJECTOR_ROTATE_SPEED
		}

		// Send a message to override / increase the selected switch shift.
		cmd := common.Command{
			Action: common.OverrideRotateSpeed,
			Args: []common.Arg{
				{Name: "SwitchNumber", Value: this.SelectedSwitch},
				{Name: "SwitchPosition", Value: this.SwitchPosition[this.SelectedSwitch]},
				{Name: "RotateSpeed", Value: overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Rotate},
			},
		}
		common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)
	}
}

func decreaseProgram(this *CurrentState, commandChannels []chan common.Command) {

	// Pull the overrides
	overrides := *this.SwitchOverrides
	switchPosition := this.SwitchPosition[this.SelectedSwitch]
	IsProgramOverrideAble := overrides[this.SelectedSwitch][switchPosition].IsProgramOverrideAble

	if IsProgramOverrideAble {

		// Decrement the program/show number .
		overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Program--
		if overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Program < common.MIN_PROGRAM_NUMBER {
			overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Program = common.MIN_PROGRAM_NUMBER
		}

		// Send a message to override / increase the selected rotate speed.
		cmd := common.Command{
			Action: common.OverrideProgram,
			Args: []common.Arg{
				{Name: "SwitchNumber", Value: this.SelectedSwitch},
				{Name: "SwitchPosition", Value: this.SwitchPosition[this.SelectedSwitch]},
				{Name: "Rotate", Value: overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Rotate},
			},
		}
		common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

		// Push the overrides.
		this.SwitchOverrides = &overrides
	}
}

func increaseShiftRGB(this *CurrentState, commandChannels []chan common.Command) {

	if debug {
		fmt.Printf("Increase Shift %d\n", this.RGBShift[this.TargetSequence])
	}

	// Increment the RGB Shift.
	this.RGBShift[this.TargetSequence]++
	if this.RGBShift[this.TargetSequence] > common.MAX_RGB_SHIFT {
		this.RGBShift[this.TargetSequence] = common.MAX_RGB_SHIFT
	}

	// Send a message to the RGB sequence.
	cmd := common.Command{
		Action: common.UpdateRGBShift,
		Args: []common.Arg{
			{Name: "Shift", Value: this.RGBShift[this.TargetSequence]},
		},
	}
	common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)
}

func increaseShiftScanner(this *CurrentState, commandChannels []chan common.Command) {

	// Increment the Scanner Shift.
	this.ScannerShift[this.TargetSequence]++
	if this.ScannerShift[this.TargetSequence] > common.MAX_SCANNER_SHIFT {
		this.ScannerShift[this.TargetSequence] = common.MAX_SCANNER_SHIFT
	}

	// Send a message to the Scanner sequence.
	cmd := common.Command{
		Action: common.UpdateScannerShift,
		Args: []common.Arg{
			{Name: "ScannerShift", Value: this.ScannerShift[this.TargetSequence]},
		},
	}
	common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

}

// func increaseShiftSwitchRGB(this *CurrentState, commandChannels []chan common.Command) {

// 	// Pull the overrides.
// 	overrides := *this.SwitchOverrides

// 	overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Shift = overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Shift + 1
// 	if overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Shift > common.MAX_RGB_SHIFT {
// 		overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Shift = common.MAX_RGB_SHIFT
// 	}

// 	// Send a message to override / increase the selected switch shift.
// 	cmd := common.Command{
// 		Action: common.OverrideShift,
// 		Args: []common.Arg{
// 			{Name: "SwitchNumber", Value: this.SelectedSwitch},
// 			{Name: "SwitchPosition", Value: this.SwitchPosition[this.SelectedSwitch]},
// 			{Name: "Shift", Value: overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Shift},
// 		},
// 	}
// 	common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

// 	// Push the overrides.
// 	this.SwitchOverrides = &overrides

// }

func increaseRotateSpeed(this *CurrentState, commandChannels []chan common.Command) {

	// Pull the overrides.
	overrides := *this.SwitchOverrides
	switchPosition := this.SwitchPosition[this.SelectedSwitch]

	isRotateOverrideAble := overrides[this.SelectedSwitch][switchPosition].IsRotateOverrideAble

	if isRotateOverrideAble {

		// Increment the Rotate Speed.
		overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Rotate++
		if overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Rotate > overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].MaxRotateSpeed {
			overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Rotate = overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].MaxRotateSpeed
		}

		// Send a message to override / increase the selected switch shift.
		cmd := common.Command{
			Action: common.OverrideRotateSpeed,
			Args: []common.Arg{
				{Name: "SwitchNumber", Value: this.SelectedSwitch},
				{Name: "SwitchPosition", Value: this.SwitchPosition[this.SelectedSwitch]},
				{Name: "RotateSpeed", Value: overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Rotate},
			},
		}
		common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)
	}
}

func increaseProgram(this *CurrentState, commandChannels []chan common.Command) {

	// Pull the overrides.
	overrides := *this.SwitchOverrides
	switchPosition := this.SwitchPosition[this.SelectedSwitch]

	IsProgrameNumberOverrideAble := overrides[this.SelectedSwitch][switchPosition].IsProgramOverrideAble
	maxProgramNumber := overrides[this.SelectedSwitch][switchPosition].MaxPrograms

	if IsProgrameNumberOverrideAble {

		// Increment the program/show number .
		overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Program++
		if overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Program > maxProgramNumber {
			overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Program = maxProgramNumber
		}

		// Send a message to override / increase the program / show number.
		cmd := common.Command{
			Action: common.OverrideProgram,
			Args: []common.Arg{
				{Name: "SwitchNumber", Value: this.SelectedSwitch},
				{Name: "SwitchPosition", Value: this.SwitchPosition[this.SelectedSwitch]},
				{Name: "Shift", Value: overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Shift},
			},
		}
		common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

	}

	// Push the overrides.
	this.SwitchOverrides = &overrides

}
