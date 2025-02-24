// Copyright (C) 2022, 2023 , 2024 dhowlett99.
// This is button processor, used by the launchpad and gui interfaces.
// This file processes the speed buttons and controls their actions.
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

func decreaseSpeed(sequences []*common.Sequence, X int, Y int, this *CurrentState, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command, updateChannels []chan common.Sequence) {

	if debug {
		fmt.Printf("Decrease Speed \n")
	}

	buttonTouched(common.Button{X: X, Y: Y}, colors.White, colors.Cyan, eventsForLaunchpad, guiButtons)

	// If we're in shutter chase mode.
	this.TargetSequence = CheckType(this.SequenceType[this.SelectedSequence], this)

	// Strobe only every operates on the selected sequence, i.e chaser never applies strobe.
	// Decrease Strobe Speed.
	if this.Strobe[this.SelectedSequence] {
		this.StrobeSpeed[this.SelectedSequence] -= 10
		if this.StrobeSpeed[this.SelectedSequence] < 0 {
			this.StrobeSpeed[this.SelectedSequence] = 0
		}

		cmd := common.Command{
			Action: common.UpdateStrobeSpeed,
			Args: []common.Arg{
				{Name: "STROBE_SPEED", Value: this.StrobeSpeed[this.SelectedSequence]},
			},
		}
		common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)
		if this.SelectedType == "scanner" && this.ScannerChaser[this.SelectedSequence] {
			common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)
		}

		// Update the status bar
		common.UpdateStatusBar(fmt.Sprintf("Strobe %02d", this.StrobeSpeed[this.SelectedSequence]), "speed", false, guiButtons)

		return
	}

	// Get an upto date copy of the sequence so we know if the music trigger is on in the sequence.
	sequences[this.TargetSequence] = common.RefreshSequence(this.TargetSequence, commandChannels, updateChannels)

	// Don't give option to change speed when in music trigger mode.
	if !sequences[this.TargetSequence].MusicTrigger {

		// Deal with an RGB sequence.
		if sequences[this.TargetSequence].Type == "rgb" || sequences[this.TargetSequence].Type == "scanner" {

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

			UpdateSpeed(this, guiButtons)

			return
		}

		// Deal with an RGB Switch sequence.
		if this.SelectedType == "switch" && this.SelectedFixtureType == "rgb" {

			// Decrement the Switch Speed.
			overrides := *this.SwitchOverrides
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

			UpdateSpeed(this, guiButtons)

			return
		}

		// Deal with an Switch that holds a projector.
		if this.SelectedType == "switch" && this.SelectedFixtureType == "projector" {

			// Decrement the Switch Speed.
			overrides := *this.SwitchOverrides
			overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Speed = overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Speed - 1
			if overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Speed < common.MIN_PROJECTOR_ROTATE_SPEED {
				overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Speed = common.MIN_PROJECTOR_ROTATE_SPEED
			}
			this.SwitchOverrides = &overrides

			// Send a message to override / increase the selected switch shift.
			cmd := common.Command{
				Action: common.OverrideSpeed,
				Args: []common.Arg{
					{Name: "SwitchNumber", Value: this.SelectedSwitch},
					{Name: "SwitchPosition", Value: this.SwitchPosition[this.SelectedSwitch]},
					{Name: "Speed", Value: overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Speed},
				},
			}
			common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

			// Get the current switch state in order to decide what we display on status bar.
			this.SwitchStateName = sequences[this.SelectedSequence].Switches[this.SelectedSwitch].States[this.SwitchPosition[this.SelectedSwitch]].Name

			// Update the status bar
			UpdateSpeed(this, guiButtons)

			return
		}
	}
}

func increaseSpeed(sequences []*common.Sequence, X int, Y int, this *CurrentState, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command, updateChannels []chan common.Sequence) {

	if debug {
		fmt.Printf("Increase Speed SelectedType %s SelectedFixtureType %s\n", this.SelectedType, this.SelectedFixtureType)
	}

	buttonTouched(common.Button{X: X, Y: Y}, colors.White, colors.Cyan, eventsForLaunchpad, guiButtons)

	// If we're in shutter chase mode.
	this.TargetSequence = CheckType(this.SequenceType[this.SelectedSequence], this)

	// Strobe only every operates on the selected sequence, i.e chaser never applies strobe.
	// Increase Strobe Speed.
	if this.Strobe[this.SelectedSequence] {
		this.StrobeSpeed[this.SelectedSequence] += 10
		if this.StrobeSpeed[this.SelectedSequence] > 255 {
			this.StrobeSpeed[this.SelectedSequence] = 255
		}

		cmd := common.Command{
			Action: common.UpdateStrobeSpeed,
			Args: []common.Arg{
				{Name: "STROBE_SPEED", Value: this.StrobeSpeed[this.SelectedSequence]},
			},
		}
		common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)
		if this.SelectedType == "scanner" && this.ScannerChaser[this.SelectedSequence] {
			common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)
		}

		// Update the status bar
		common.UpdateStatusBar(fmt.Sprintf("Strobe %02d", this.StrobeSpeed[this.SelectedSequence]), "speed", false, guiButtons)

		return
	}

	// Get an upto date copy of the sequence so we know if the music trigger is on in the sequence.
	sequences[this.TargetSequence] = common.RefreshSequence(this.TargetSequence, commandChannels, updateChannels)

	if !sequences[this.TargetSequence].MusicTrigger {

		// Deal with an RGB / Scanner sequence.
		if sequences[this.TargetSequence].Type == "rgb" || sequences[this.TargetSequence].Type == "scanner" {

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

			// Update the status bar
			UpdateSpeed(this, guiButtons)

			return
		}

		// Deal with an Switch sequence with a RGB fixture.
		if this.SelectedType == "switch" && this.SelectedFixtureType == "rgb" {

			overrides := *this.SwitchOverrides
			overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Speed = overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Speed + 1
			if overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Speed > common.MAX_SPEED {
				overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Speed = common.MAX_SPEED
			}
			this.SwitchOverrides = &overrides

			// Send a message to override / increase the selected switch speed.
			cmd := common.Command{
				Action: common.OverrideSpeed,
				Args: []common.Arg{
					{Name: "SwitchNumber", Value: this.SelectedSwitch},
					{Name: "SwitchPosition", Value: this.SwitchPosition[this.SelectedSwitch]},
					{Name: "Speed", Value: overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Speed},
				},
			}
			common.SendCommandToSequence(this.SwitchSequenceNumber, cmd, commandChannels)

			// Update the status bar
			UpdateSpeed(this, guiButtons)

			return
		}

		// Deal with an Switch sequence that has a projector fixture.
		if this.SelectedType == "switch" && this.SelectedFixtureType == "projector" {

			overrides := *this.SwitchOverrides
			overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Speed = overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Speed + 1
			if overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Speed > overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].MaxSpeeds {
				overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Speed = overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].MaxSpeeds
			}
			this.SwitchOverrides = &overrides

			// Send a message to override / increase the selected switch shift.
			cmd := common.Command{
				Action: common.OverrideSpeed,
				Args: []common.Arg{
					{Name: "SwitchNumber", Value: this.SelectedSwitch},
					{Name: "SwitchPosition", Value: this.SwitchPosition[this.SelectedSwitch]},
					{Name: "Speed", Value: overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Speed},
				},
			}
			common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

			// Get the current switch state in order to decide what we display on status bar.
			this.SwitchStateName = sequences[this.SelectedSequence].Switches[this.SelectedSwitch].States[this.SwitchPosition[this.SelectedSwitch]].Name

			// Update the status bar
			UpdateSpeed(this, guiButtons)

			return
		}
	}
}
