// Copyright (C) 2022, 2023 , 2024 dhowlett99.
// This is button processor, used by the launchpad and gui interfaces.
// This file processes the strobe buttons and controls their actions.
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

func toggleStrobe(sequences []*common.Sequence, X int, Y int, this *CurrentState, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command) {

	if debug {
		fmt.Printf("Strobe X:%d Y:%d\n", X, Y)
	}

	// Turn off the flashing save button
	this.SavePreset = false
	this.SavePreset = false
	common.LightLamp(common.SAVE_BUTTON, colors.White, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)

	// Shutdown any function bars.
	clearAllModes(sequences, this)
	HandleSelect(sequences, this, eventsForLaunchpad, commandChannels, guiButtons)

	// If strobing, stop it
	if this.Strobe[this.SelectedSequence] {
		this.Strobe[this.SelectedSequence] = false
		// Stop strobing this sequence.
		StopStrobe(this, eventsForLaunchpad, guiButtons, commandChannels)
		return

	} else {
		// Start strobing for this sequence. Strobe on.
		this.Strobe[this.SelectedSequence] = true
		StartStrobe(this, eventsForLaunchpad, guiButtons, commandChannels)
		return

	}
}

func StopStrobe(this *CurrentState, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command) {

	if this.SelectedType == "rgb" {

		// Stop strobing this sequence.
		cmd := common.Command{
			Action: common.Strobe,
			Args: []common.Arg{
				{Name: "STROBE_STATE", Value: this.Strobe[this.SelectedSequence]},
				{Name: "STROBE_SPEED", Value: this.StrobeSpeed[this.SelectedSequence]},
			},
		}
		// Send a message to the sequence.
		common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

	}

	if this.SelectedType == "scanner" {

		// Stop strobing this sequence.
		cmd := common.Command{
			Action: common.Strobe,
			Args: []common.Arg{
				{Name: "STROBE_STATE", Value: this.Strobe[this.SelectedSequence]},
				{Name: "STROBE_SPEED", Value: this.StrobeSpeed[this.SelectedSequence]},
			},
		}
		// Send a message to the sequence.
		common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

		// Also send to chaser sequencer.
		if this.ScannerChaser[this.SelectedSequence] {
			common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)
		}
	}

	if this.SelectedType == "switch" {

		// Pull the overrides.
		overrides := *this.SwitchOverrides

		// Stop the strobe
		overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Strobe = false

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

	// Update the strobe button and status bar.
	common.ShowStrobeButtonStatus(this.Strobe[this.SelectedSequence], eventsForLaunchpad, guiButtons)
	common.UpdateStatusBar(fmt.Sprintf("Speed %02d", this.Speed[this.SelectedSequence]), "speed", false, guiButtons)
}

func StartStrobe(this *CurrentState, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command) {

	if this.SelectedType == "rgb" {

		cmd := common.Command{
			Action: common.Strobe,
			Args: []common.Arg{
				{Name: "STROBE_STATE", Value: this.Strobe[this.SelectedSequence]},
				{Name: "STROBE_SPEED", Value: this.StrobeSpeed[this.SelectedSequence]},
			},
		}
		// Send a message to the sequence.
		common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)
	}

	if this.SelectedType == "scanner" {

		cmd := common.Command{
			Action: common.Strobe,
			Args: []common.Arg{
				{Name: "STROBE_STATE", Value: this.Strobe[this.SelectedSequence]},
				{Name: "STROBE_SPEED", Value: this.StrobeSpeed[this.SelectedSequence]},
			},
		}
		// Send a message to the sequence.
		common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

		// Also send to chaser sequencer.
		if this.ScannerChaser[this.SelectedSequence] {
			common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)
		}
	}

	if this.SelectedType == "switch" {

		// Pull the overrides.
		overrides := *this.SwitchOverrides

		// Enable the strobe.
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

	// Update the strobe button and status bar.
	common.ShowStrobeButtonStatus(this.Strobe[this.SelectedSequence], eventsForLaunchpad, guiButtons)
	common.UpdateStatusBar(fmt.Sprintf("Strobe %02d", this.StrobeSpeed[this.SelectedSequence]), "speed", false, guiButtons)
}
