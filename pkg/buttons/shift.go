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

	"github.com/dhowlett99/dmxlights/pkg/colors"
	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
)

func decreaseShift(sequences []*common.Sequence, X int, Y int, this *CurrentState, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command, fixturesConfig *fixture.Fixtures) {

	if debug {
		fmt.Printf("Decrease Shift\n")
	}

	buttonTouched(common.Button{X: X, Y: Y}, colors.White, colors.Cyan, eventsForLaunchpad, guiButtons)

	// If we're in shutter chase mode.
	if this.SelectedMode[this.SelectedSequence] == CHASER_FUNCTION || this.SelectedMode[this.SelectedSequence] == CHASER_DISPLAY {
		this.TargetSequence = this.ChaserSequenceNumber
	} else {
		this.TargetSequence = this.SelectedSequence
	}

	// Deal with an RGB sequence.
	if sequences[this.TargetSequence].Type == "rgb" {

		// Decrement the RGB Shift.
		this.RGBShift[this.TargetSequence] = this.RGBShift[this.TargetSequence] - 1
		if this.RGBShift[this.TargetSequence] < 0 {
			this.RGBShift[this.TargetSequence] = 0
		}

		// Send a message to the RGB sequence.
		cmd := common.Command{
			Action: common.UpdateRGBShift,
			Args: []common.Arg{
				{Name: "RGBShift", Value: this.RGBShift[this.TargetSequence]},
			},
		}
		common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

		// Update the status bar
		UpdateShift(this, guiButtons)

		return
	}

	// Deal with an Scanner sequence.
	if sequences[this.TargetSequence].Type == "scanner" {

		// Decrement the Scanner Shift.
		this.ScannerShift[this.TargetSequence] = this.ScannerShift[this.TargetSequence] - 1
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

		// Update the status bar
		UpdateShift(this, guiButtons)

		return
	}

	// Deal with an RGB Switch sequence.
	if this.SelectedType == "switch" && this.SelectedFixtureType == "rgb" {

		// Decrement the Switch Shift.
		this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Shift = this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Shift - 1
		if this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Shift < 0 {
			this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Shift = 0
		}

		// Send a message to override / increase the selected switch shift.
		cmd := common.Command{
			Action: common.OverrideShift,
			Args: []common.Arg{
				{Name: "SwitchNumber", Value: this.SelectedSwitch},
				{Name: "SwitchPosition", Value: this.SwitchPosition[this.SelectedSwitch]},
				{Name: "Shift", Value: this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Shift},
			},
		}
		common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

		// Update the status bar
		UpdateShift(this, guiButtons)

		return
	}

	// Deal with an Switch that holds a projector.
	if this.SelectedType == "switch" && this.SelectedFixtureType == "projector" {

		// Decrement the Switch Shift.
		this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].RotateSpeed = this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].RotateSpeed - 1
		if this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].RotateSpeed < common.MIN_PROJECTOR_ROTATE_SPEED {
			this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].RotateSpeed = 0
		}

		// Send a message to override / increase the selected switch shift.
		cmd := common.Command{
			Action: common.OverrideRotateSpeed,
			Args: []common.Arg{
				{Name: "SwitchNumber", Value: this.SelectedSwitch},
				{Name: "SwitchPosition", Value: this.SwitchPosition[this.SelectedSwitch]},
				{Name: "RotateSpeed", Value: this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].RotateSpeed},
			},
		}
		common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

		// Set the rotate speed name.
		theSwitch, _ := fixture.FindFixtureByGroupAndNumber(this.SelectedSequence, this.SelectedSwitch, fixturesConfig)
		useFixture, _ := fixture.FindFixtureByLabel(theSwitch.UseFixture, fixturesConfig)
		this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].RotateSpeedName = fixture.FindRotateSpeedNameByNumber(useFixture, this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].RotateSpeed)

		// Update the status bar
		UpdateShift(this, guiButtons)

		return
	}

}

func increaseShift(sequences []*common.Sequence, X int, Y int, this *CurrentState, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command, fixturesConfig *fixture.Fixtures) {

	if debug {
		fmt.Printf("Increase Shift \n")
	}

	buttonTouched(common.Button{X: X, Y: Y}, colors.White, colors.Cyan, eventsForLaunchpad, guiButtons)

	// If we're in shutter chase mode.
	if this.SelectedMode[this.SelectedSequence] == CHASER_FUNCTION || this.SelectedMode[this.SelectedSequence] == CHASER_DISPLAY {
		this.TargetSequence = this.ChaserSequenceNumber
	} else {
		this.TargetSequence = this.SelectedSequence
	}

	// Deal with an RGB sequence.
	if sequences[this.TargetSequence].Type == "rgb" {

		// Increment the RGB Shift.
		this.RGBShift[this.TargetSequence] = this.RGBShift[this.TargetSequence] + 1
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

		// Update the status bar
		UpdateShift(this, guiButtons)

		return
	}

	// Deal with an Scanner sequence.
	if sequences[this.TargetSequence].Type == "scanner" {

		// Increment the Scanner Shift.
		this.ScannerShift[this.TargetSequence] = this.ScannerShift[this.TargetSequence] + 1
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

		// Update the status bar
		UpdateShift(this, guiButtons)

		return
	}

	// Deal with an Switch sequence with a RGB fixture.
	if this.SelectedType == "switch" && this.SelectedFixtureType == "rgb" {

		this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Shift = this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Shift + 1
		if this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Shift > common.MAX_RGB_SHIFT {
			this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Shift = common.MAX_RGB_SHIFT
		}

		// Send a message to override / increase the selected switch shift.
		cmd := common.Command{
			Action: common.OverrideShift,
			Args: []common.Arg{
				{Name: "SwitchNumber", Value: this.SelectedSwitch},
				{Name: "SwitchPosition", Value: this.SwitchPosition[this.SelectedSwitch]},
				{Name: "Shift", Value: this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Shift},
			},
		}
		common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

		// Update the status bar
		UpdateShift(this, guiButtons)

		return
	}

	// Deal with an Switch sequence that has a projector fixture.
	if this.SelectedType == "switch" && this.SelectedFixtureType == "projector" {

		this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].RotateSpeed = this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].RotateSpeed + 1
		if this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].RotateSpeed > common.MAX_PROJECTOR_ROTATE_SPEED {
			this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].RotateSpeed = common.MAX_PROJECTOR_ROTATE_SPEED
		}

		// Send a message to override / increase the selected switch shift.
		cmd := common.Command{
			Action: common.OverrideRotateSpeed,
			Args: []common.Arg{
				{Name: "SwitchNumber", Value: this.SelectedSwitch},
				{Name: "SwitchPosition", Value: this.SwitchPosition[this.SelectedSwitch]},
				{Name: "RotateSpeed", Value: this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].RotateSpeed},
			},
		}
		common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

		// Set the rotate speed name.
		theSwitch, _ := fixture.FindFixtureByGroupAndNumber(this.SelectedSequence, this.SelectedSwitch, fixturesConfig)
		useFixture, _ := fixture.FindFixtureByLabel(theSwitch.UseFixture, fixturesConfig)
		this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].RotateSpeedName = fixture.FindRotateSpeedNameByNumber(useFixture, this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].RotateSpeed)

		// Update the status bar
		UpdateShift(this, guiButtons)

		return
	}
}
