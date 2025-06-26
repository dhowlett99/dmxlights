// Copyright (C) 2022, 2023 , 2024 dhowlett99.
// This is button processor, used by the launchpad and gui interfaces.
// This file processes fade buttons and controls their actions.
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

func decreaseFadeRGB(this *CurrentState, commandChannels []chan common.Command) {

	if debug {
		fmt.Printf("Decrease Fade Time Type=%s Sequence=%d \n", this.SelectedType, this.TargetSequence)
	}

	// If we're a scanner and we're in shutter chase mode.
	this.TargetSequence = CheckType(this.SequenceType[this.SelectedSequence], this)

	// Decrement the RGB Fade size.
	this.RGBFade[this.TargetSequence]--
	if this.RGBFade[this.TargetSequence] < 1 {
		this.RGBFade[this.TargetSequence] = 1
	}

	// Send fade update command to sequence.
	cmd := common.Command{
		Action: common.UpdateRGBFadeSpeed,
		Args: []common.Arg{
			{Name: "RGBFadeSpeed", Value: this.RGBFade[this.TargetSequence]},
		},
	}
	common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

}

func decreaseFadeScannerNumberCoordinates(this *CurrentState, commandChannels []chan common.Command) {

	// Update Coordinates.

	// Fade also send more or less coordinates for the scanner patterns.
	this.ScannerCoordinates[this.TargetSequence]--
	if this.ScannerCoordinates[this.TargetSequence] < 0 {
		this.ScannerCoordinates[this.TargetSequence] = 0
	}

	// Send a messages to the scanner sequence.
	cmd := common.Command{
		Action: common.UpdateNumberCoordinates,
		Args: []common.Arg{
			{Name: "NumberCoordinates", Value: this.ScannerCoordinates[this.TargetSequence]},
		},
	}
	common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

}

// func decreaseOverrideFadeRGB(this *CurrentState, commandChannels []chan common.Command) {

// 	// Decrease the fade size.
// 	overrides := *this.SwitchOverrides
// 	overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Fade--
// 	if overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Fade < common.MIN_RGB_FADE {
// 		overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Fade = common.MIN_RGB_FADE
// 	}
// 	this.SwitchOverrides = &overrides

// 	// Send a message to override / increase the selected switch shift.
// 	cmd := common.Command{
// 		Action: common.OverrideFade,
// 		Args: []common.Arg{
// 			{Name: "SwitchNumber", Value: this.SelectedSwitch},
// 			{Name: "SwitchPosition", Value: this.SwitchPosition[this.SelectedSwitch]},
// 			{Name: "Shift", Value: overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Fade},
// 		},
// 	}
// 	common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

// }

func decreaseGobo(this *CurrentState, commandChannels []chan common.Command) {

	// Decrease the gobo.
	overrides := *this.SwitchOverrides
	overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Gobo--
	if overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Gobo < common.MIN_PROJECTOR_GOBO {
		overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Gobo = common.MIN_PROJECTOR_GOBO
	}
	this.SwitchOverrides = &overrides

	// Send a message to select the gobo.
	cmd := common.Command{
		Action: common.OverrideGobo,
		Args: []common.Arg{
			{Name: "SwitchNumber", Value: this.SelectedSwitch},
			{Name: "SwitchPosition", Value: this.SwitchPosition[this.SelectedSwitch]},
			{Name: "Gobo", Value: overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Gobo},
		},
	}
	common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

}

func increaseFadeRGB(this *CurrentState, commandChannels []chan common.Command) {

	if debug {
		fmt.Printf("Increase Fade Time\n")
	}

	// Increase fade time.
	this.RGBFade[this.TargetSequence]++
	if this.RGBFade[this.TargetSequence] > common.MAX_RGB_FADE {
		this.RGBFade[this.TargetSequence] = common.MAX_RGB_FADE
	}

	// Send fade update command.
	cmd := common.Command{
		Action: common.UpdateRGBFadeSpeed,
		Args: []common.Arg{
			{Name: "FadeSpeed", Value: this.RGBFade[this.TargetSequence]},
		},
	}
	common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

}

func increaseFadeScannerNumberCoordinates(this *CurrentState, commandChannels []chan common.Command) {

	// Fade also send more or less coordinates for the scanner patterns.
	this.ScannerCoordinates[this.TargetSequence]++
	if this.ScannerCoordinates[this.TargetSequence] > 4 {
		this.ScannerCoordinates[this.TargetSequence] = 4
	}

	// Send a message to scanner seqiemce.
	cmd := common.Command{
		Action: common.UpdateNumberCoordinates,
		Args: []common.Arg{
			{Name: "NumberCoordinates", Value: this.ScannerCoordinates[this.TargetSequence]},
		},
	}
	common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

}

// func increaseFadeSwitchRGB(this *CurrentState, commandChannels []chan common.Command) {

// 	// Increase the switch size.
// 	overrides := *this.SwitchOverrides
// 	overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Fade = overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Fade + 1
// 	if overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Fade > common.MAX_RGB_SHIFT {
// 		overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Fade = common.MAX_RGB_SHIFT
// 	}
// 	this.SwitchOverrides = &overrides

// 	// Send a message to override / increase the selected switch shift.
// 	cmd := common.Command{
// 		Action: common.OverrideFade,
// 		Args: []common.Arg{
// 			{Name: "SwitchNumber", Value: this.SelectedSwitch},
// 			{Name: "SwitchPosition", Value: this.SwitchPosition[this.SelectedSwitch]},
// 			{Name: "Shift", Value: overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Fade},
// 		},
// 	}
// 	common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

// }

func increaseGobo(this *CurrentState, commandChannels []chan common.Command) {

	// Deal with an Switch sequence with a projector fixture. Increase Gobo
	if this.SelectedType == "switch" && this.SelectedFixtureType == "projector" {

		// Increase the switch size.
		overrides := *this.SwitchOverrides
		maxNumberGobos := overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].MaxGobos
		overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Gobo = overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Gobo + 1
		if overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Gobo > maxNumberGobos {
			overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Gobo = maxNumberGobos
		}

		// Send a message to override / increase the selected gobo.
		cmd := common.Command{
			Action: common.OverrideGobo,
			Args: []common.Arg{
				{Name: "SwitchNumber", Value: this.SelectedSwitch},
				{Name: "SwitchPosition", Value: this.SwitchPosition[this.SelectedSwitch]},
				{Name: "Gobo", Value: overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Gobo},
			},
		}
		common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

	}
}
