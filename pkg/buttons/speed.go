// Copyright (C) 2022, 2023, 2024, 2025 dhowlett99.
// This is speed code used for sending messages to the sequence and
// than calling the update status bar function.
// It is called when we update :- increase or decrease speed.
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
	"github.com/dhowlett99/dmxlights/pkg/common"
)

// Deal with an RGB sequence.
func decreaseSpeed(sequences []*common.Sequence, this *CurrentState, commandChannels []chan common.Command) {

	// Strobe only every operates on the selected sequence, i.e chaser never applies strobe.
	// Decrease Strobe Speed.
	if this.Strobe[this.SelectedSequence] {

		if sequences[this.TargetSequence].Type == "rgb" || sequences[this.TargetSequence].Type == "scanner" {

			this.StrobeSpeed[this.SelectedSequence] -= 10
			if this.StrobeSpeed[this.SelectedSequence] < 0 {
				this.StrobeSpeed[this.SelectedSequence] = 0
			}

			cmd := common.Command{
				Action: common.UpdateStrobeSpeed,
				Args: []common.Arg{
					{Name: "STROBE", Value: this.Strobe[this.SelectedSequence]},
					{Name: "STROBE_SPEED", Value: this.StrobeSpeed[this.SelectedSequence]},
				},
			}
			common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)
			if this.SelectedType == "scanner" && this.ScannerChaser[this.SelectedSequence] {
				common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)
			}

		}

		// Override switch fixture.
		if sequences[this.TargetSequence].Type == "switch" {

			// Pull the overrides.
			overrides := *this.SwitchOverrides

			// Stop the strobe
			this.StrobeSpeed[this.SelectedSequence] -= 25
			if this.StrobeSpeed[this.SelectedSequence] < 0 {
				this.StrobeSpeed[this.SelectedSequence] = 0
			}

			overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Strobe = true

			// Copy in the current strobe speed.
			overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].StrobeSpeed = this.StrobeSpeed[this.SelectedSequence]

			cmd := common.Command{
				Action: common.OverrideStrobe,
				Args: []common.Arg{
					{Name: "SwitchNumber", Value: this.SelectedSwitch},
					{Name: "SwitchPosition", Value: this.SwitchPosition[this.SelectedSwitch]},
					{Name: "Strobe", Value: overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Strobe},
					{Name: "Strobe Speed", Value: overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].StrobeSpeed},
				},
			}
			// Send a message to the sequence.
			common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

			// Push the overrides.
			this.SwitchOverrides = &overrides

		}
	}
	// Don't give option to change speed when in music trigger mode.
	if !this.SoundTriggers[this.SelectedSequence].State {

		// Decrease RGB / Scanner Speed.
		this.Speed[this.TargetSequence]--
		if this.Speed[this.TargetSequence] < 1 {
			this.Speed[this.TargetSequence] = 1
		}

		cmd := common.Command{
			Action: common.UpdateSpeed,
			Args: []common.Arg{
				{Name: "Speed", Value: this.Speed[this.TargetSequence]},
			},
		}

		common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

		// Speed is used to control fade time in mini sequencer so send to switch sequence as well.
		common.SendCommandToSequence(this.SwitchSequenceNumber, cmd, commandChannels)

	}
}

func decreaseOverrideSpeedRGB(this *CurrentState, commandChannels []chan common.Command) {

	// Pull the overrides.
	overrides := *this.SwitchOverrides

	// Decrement the Switch Speed.
	overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Speed = overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Speed - 1
	if overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Speed < 0 {
		overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Speed = 0
	}
	this.SwitchOverrides = &overrides

	// Send a message to override / decrease the selected switch speed.
	cmd := common.Command{
		Action: common.OverrideSpeed,
		Args: []common.Arg{
			{Name: "SwitchNumber", Value: this.SelectedSwitch},
			{Name: "SwitchPosition", Value: this.SwitchPosition[this.SelectedSwitch]},
			{Name: "Speed", Value: overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Speed},
		},
	}
	common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

	// Push the overrides.
	this.SwitchOverrides = &overrides

}

func decreaseOverrideProgramSpeed(this *CurrentState, commandChannels []chan common.Command) {

	// Pull the overrides.
	overrides := *this.SwitchOverrides

	// Decrement the Switch Speed.
	overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].ProgramSpeed = overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].ProgramSpeed - 1
	if overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].ProgramSpeed < common.MIN_PROJECTOR_PROGRAM_SPEED {
		overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].ProgramSpeed = common.MIN_PROJECTOR_PROGRAM_SPEED
	}

	// Send a message to override / increase the selected switch shift.
	cmd := common.Command{
		Action: common.OverrideProgramSpeed,
		Args: []common.Arg{
			{Name: "SwitchNumber", Value: this.SelectedSwitch},
			{Name: "SwitchPosition", Value: this.SwitchPosition[this.SelectedSwitch]},
			{Name: "Program Speed", Value: overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].ProgramSpeed},
		},
	}
	common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

}

// func decreaseSpeedSwitchChaser(this *CurrentState, commandChannels []chan common.Command) {

// 	// Pull the overrides.
// 	overrides := *this.SwitchOverrides

// 	// Decrement the Switch Speed.
// 	overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Speed = overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Speed - 1
// 	if overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Speed < common.MIN_PROJECTOR_PROGRAM_SPEED {
// 		overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Speed = common.MIN_PROJECTOR_PROGRAM_SPEED
// 	}

// 	// Send a message to override / increase the selected switch speed.
// 	cmd := common.Command{
// 		Action: common.OverrideSpeed,
// 		Args: []common.Arg{
// 			{Name: "SwitchNumber", Value: this.SelectedSwitch},
// 			{Name: "SwitchPosition", Value: this.SwitchPosition[this.SelectedSwitch]},
// 			{Name: "Speed", Value: overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Speed},
// 		},
// 	}
// 	common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

// 	// Push the overrides.
// 	this.SwitchOverrides = &overrides

// }

func increaseSpeedRGB(sequences []*common.Sequence, this *CurrentState, commandChannels []chan common.Command) {

	// Strobe only every operates on the selected sequence, i.e chaser never applies strobe.
	// Increase Strobe Speed.
	if this.Strobe[this.SelectedSequence] {

		if sequences[this.TargetSequence].Type == "rgb" || sequences[this.TargetSequence].Type == "scanner" {

			this.StrobeSpeed[this.SelectedSequence] += 10
			if this.StrobeSpeed[this.SelectedSequence] > common.MAX_STROBE_SPEED {
				this.StrobeSpeed[this.SelectedSequence] = common.MAX_STROBE_SPEED
			}

			cmd := common.Command{
				Action: common.UpdateStrobeSpeed,
				Args: []common.Arg{
					{Name: "STROBE", Value: this.Strobe[this.SelectedSequence]},
					{Name: "STROBE_SPEED", Value: this.StrobeSpeed[this.SelectedSequence]},
				},
			}
			common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)
			if this.SelectedType == "scanner" && this.ScannerChaser[this.SelectedSequence] {
				common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)
			}
		}

		// Override switch fixture.
		if sequences[this.TargetSequence].Type == "switch" {

			// Pull the overrides.
			overrides := *this.SwitchOverrides

			this.StrobeSpeed[this.SelectedSequence] += 25
			if this.StrobeSpeed[this.SelectedSequence] > 255 {
				this.StrobeSpeed[this.SelectedSequence] = 255
			}

			overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Strobe = true

			// Copy in the current strobe speed.
			overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].StrobeSpeed = this.StrobeSpeed[this.SelectedSequence]

			cmd := common.Command{
				Action: common.OverrideStrobe,
				Args: []common.Arg{
					{Name: "SwitchNumber", Value: this.SelectedSwitch},
					{Name: "SwitchPosition", Value: this.SwitchPosition[this.SelectedSwitch]},
					{Name: "Strobe", Value: overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Strobe},
					{Name: "Strobe Speed", Value: overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].StrobeSpeed},
				},
			}
			// Send a message to the sequence.
			common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

			// Push the overrides.
			this.SwitchOverrides = &overrides

		}
	}

	// Increment the RGB Speed.
	this.Speed[this.TargetSequence]++
	if this.Speed[this.TargetSequence] > common.MAX_SPEED {
		this.Speed[this.TargetSequence] = common.MAX_SPEED
	}

	cmd := common.Command{
		Action: common.UpdateSpeed,
		Args: []common.Arg{
			{Name: "Speed", Value: this.Speed[this.TargetSequence]},
		},
	}
	common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

}

func increaseSpeedScanner(this *CurrentState, commandChannels []chan common.Command) {

	// Increment the RGB Speed.
	this.Speed[this.TargetSequence]++
	if this.Speed[this.TargetSequence] > common.MAX_SPEED {
		this.Speed[this.TargetSequence] = common.MAX_SPEED
	}

	cmd := common.Command{
		Action: common.UpdateSpeed,
		Args: []common.Arg{
			{Name: "Speed", Value: this.Speed[this.TargetSequence]},
		},
	}
	common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

}

func increaseOverrideProgramSpeed(this *CurrentState, commandChannels []chan common.Command) {

	// Pull the overrides.
	overrides := *this.SwitchOverrides

	overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].ProgramSpeed++
	if overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].ProgramSpeed > overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].MaxProgramSpeeds {
		overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].ProgramSpeed = overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].MaxProgramSpeeds
	}

	// Send a message to override / increase the selected switch shift.
	cmd := common.Command{
		Action: common.OverrideProgramSpeed,
		Args: []common.Arg{
			{Name: "SwitchNumber", Value: this.SelectedSwitch},
			{Name: "SwitchPosition", Value: this.SwitchPosition[this.SelectedSwitch]},
			{Name: "Speed", Value: overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].ProgramSpeed},
		},
	}
	common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)
}

func increaseOverrideSpeedRGB(this *CurrentState, commandChannels []chan common.Command) {

	// Pull the overrides.
	overrides := *this.SwitchOverrides

	overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Speed = overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Speed + 1
	if overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Speed > common.MAX_SPEED {
		overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Speed = common.MAX_SPEED
	}

	// Send a message to override / decrease the selected switch speed.
	cmd := common.Command{
		Action: common.OverrideSpeed,
		Args: []common.Arg{
			{Name: "SwitchNumber", Value: this.SelectedSwitch},
			{Name: "SwitchPosition", Value: this.SwitchPosition[this.SelectedSwitch]},
			{Name: "ProgramSpeed", Value: overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Speed},
		},
	}
	common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

	// Push the overrides.
	this.SwitchOverrides = &overrides

}
