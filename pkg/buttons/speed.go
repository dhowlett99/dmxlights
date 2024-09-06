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
	if this.SelectedMode[this.SelectedSequence] == CHASER_FUNCTION || this.SelectedMode[this.SelectedSequence] == CHASER_DISPLAY {
		this.TargetSequence = this.ChaserSequenceNumber
	} else {
		this.TargetSequence = this.SelectedSequence
	}

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

	// Get an upto date copy of the target sequence.
	sequences[this.TargetSequence] = common.RefreshSequence(this.TargetSequence, commandChannels, updateChannels)

	if this.SelectedType == "switch" {
		// Copy the updated speed setting into the local switch speed storage
		this.Speed[this.TargetSequence] = this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Speed
	}

	// Decrease Speed.
	if !sequences[this.TargetSequence].MusicTrigger {
		this.Speed[this.TargetSequence]--
		if this.Speed[this.TargetSequence] < 1 {
			this.Speed[this.TargetSequence] = 1
		}

		// If you reached the min speed blink the increase button.
		if this.Speed[this.TargetSequence] == common.MIN_SPEED {
			common.FlashLight(common.Button{X: X, Y: Y}, colors.White, colors.Cyan, eventsForLaunchpad, guiButtons)
		} else {
			// If you reached the half speed blink both buttons.
			if this.Speed[this.TargetSequence] == common.MAX_SPEED/2 {
				common.FlashLight(common.Button{X: X, Y: Y}, colors.White, colors.Cyan, eventsForLaunchpad, guiButtons)
				common.FlashLight(common.Button{X: X + 1, Y: Y}, colors.White, colors.Cyan, eventsForLaunchpad, guiButtons)
			} else {
				common.LightLamp(common.Button{X: X + 1, Y: Y}, colors.Cyan, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)
			}
		}

		if this.SelectedType == "switch" {
			// Copy the updated speed setting into the local switch speed storage
			this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Speed = this.Speed[this.TargetSequence]
			// Send a message to override / decrease the selected switch speed.
			cmd := common.Command{
				Action: common.OverrideSpeed,
				Args: []common.Arg{
					{Name: "SwitchNumber", Value: this.SelectedSwitch},
					{Name: "SwitchPosition", Value: this.SwitchPosition[this.SelectedSwitch]},
					{Name: "Speed", Value: this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Speed},
				},
			}
			common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)
		} else {
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

	// Update the status bar
	UpdateSpeed(this, guiButtons)

}

func increaseSpeed(sequences []*common.Sequence, X int, Y int, this *CurrentState, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command, updateChannels []chan common.Sequence) {

	if debug {
		fmt.Printf("Increase Speed \n")
	}

	buttonTouched(common.Button{X: X, Y: Y}, colors.White, colors.Cyan, eventsForLaunchpad, guiButtons)

	// If we're in shutter chase mode
	if this.SelectedMode[this.SelectedSequence] == CHASER_FUNCTION || this.SelectedMode[this.SelectedSequence] == CHASER_DISPLAY {
		this.TargetSequence = this.ChaserSequenceNumber
	} else {
		this.TargetSequence = this.SelectedSequence
	}

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

	// Get an upto date copy of the sequence.
	sequences[this.TargetSequence] = common.RefreshSequence(this.TargetSequence, commandChannels, updateChannels)

	if this.SelectedType == "switch" {
		// Copy the updated speed setting into the local switch speed storage
		this.Speed[this.TargetSequence] = this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Speed
	}

	if !sequences[this.TargetSequence].MusicTrigger {
		this.Speed[this.TargetSequence]++
		if this.Speed[this.TargetSequence] > 12 {
			this.Speed[this.TargetSequence] = 12
		}

		// If you reached the max speed blink the increase button.
		if this.Speed[this.TargetSequence] == common.MAX_SPEED {
			common.FlashLight(common.Button{X: X, Y: Y}, colors.White, colors.Cyan, eventsForLaunchpad, guiButtons)
		} else {
			// If you reached the half speed blink both buttons.
			if this.Speed[this.TargetSequence] == common.MAX_SPEED/2 {
				common.FlashLight(common.Button{X: X, Y: Y}, colors.White, colors.Cyan, eventsForLaunchpad, guiButtons)
				common.FlashLight(common.Button{X: X - 1, Y: Y}, colors.White, colors.Cyan, eventsForLaunchpad, guiButtons)
			} else {
				common.LightLamp(common.Button{X: X - 1, Y: Y}, colors.Cyan, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)
			}
		}
		if this.SelectedType == "switch" {
			// Copy the speed setting into the local switch speed storage
			this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Speed = this.Speed[this.TargetSequence]
			// Send a message to override / increase the selected switch speed.
			cmd := common.Command{
				Action: common.OverrideSpeed,
				Args: []common.Arg{
					{Name: "SwitchNumber", Value: this.SelectedSwitch},
					{Name: "SwitchPosition", Value: this.SwitchPosition[this.SelectedSwitch]},
					{Name: "Speed", Value: this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Speed},
				},
			}
			common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)
		} else {
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

	// Update the status bar
	UpdateSpeed(this, guiButtons)

}
