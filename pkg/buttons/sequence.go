// Copyright (C) 2022, 2023 , 2024 dhowlett99.
// This is button processor, used by the launchpad and gui interfaces.
// This file processes the sequence start and stop buttons and controls their actions.
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

func toggleSequence(sequences []*common.Sequence, X int, Y int, this *CurrentState, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command) {

	if this.ShowRGBColorPicker {
		this.ShowRGBColorPicker = false
		removeColorPicker(this, sequences, eventsForLaunchpad, guiButtons, commandChannels)
	}

	// Start in normal mode, hide the shutter chaser.
	if this.SelectedType == "scanner" && this.ScannerChaser[this.SelectedSequence] {
		common.HideSequence(this.ChaserSequenceNumber, commandChannels)
		common.ClearSelectedRowOfButtons(this.SelectedSequence, eventsForLaunchpad, guiButtons)
		this.SelectedMode[this.SelectedSequence] = NORMAL
	}

	this.SelectButtonPressed[this.SelectedSequence] = false
	this.StaticFlashing[this.SelectedSequence] = false

	// S T O P - If sequence is running, stop it
	if this.Running[this.SelectedSequence] {
		if debug {
			fmt.Printf("Stop Sequence %d \n", this.SelectedSequence)
		}
		cmd := common.Command{
			Action: common.Stop,
			Args: []common.Arg{
				{Name: "Speed", Value: this.Speed[this.SelectedSequence]},
			},
		}
		common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

		if this.Strobe[this.SelectedSequence] {
			this.Strobe[this.SelectedSequence] = false
			StopStrobe(this, eventsForLaunchpad, guiButtons, commandChannels)
		}

		this.Running[this.SelectedSequence] = false
		this.Functions[this.SelectedSequence][common.Function6_Static_Gobo].State = false
		this.Functions[this.SelectedSequence][common.Function8_Music_Trigger].State = false
		this.Functions[this.ChaserSequenceNumber][common.Function6_Static_Gobo].State = false
		this.Functions[this.ChaserSequenceNumber][common.Function8_Music_Trigger].State = false

		// Stop should also stop the shutter chaser.
		if this.SelectedType == "scanner" && this.ScannerChaser[this.SelectedSequence] {

			cmd := common.Command{
				Action: common.Stop,
				Args: []common.Arg{
					{Name: "Speed", Value: this.Speed[this.ChaserSequenceNumber]},
				},
			}
			common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)

			this.Functions[this.SelectedSequence][common.Function6_Static_Gobo].State = false
			this.Functions[this.SelectedSequence][common.Function7_Invert_Chase].State = false
			this.Functions[this.ChaserSequenceNumber][common.Function6_Static_Gobo].State = false
			this.Functions[this.ChaserSequenceNumber][common.Function8_Music_Trigger].State = false
			this.ScannerChaser[this.SelectedSequence] = false
			this.SelectedMode[this.SelectedSequence] = NORMAL
			this.Running[this.ChaserSequenceNumber] = false
		}

		// Clear the pattern function keys
		common.ClearSelectedRowOfButtons(this.SelectedSequence, eventsForLaunchpad, guiButtons)

		// Turn off the start lamp.
		common.LightLamp(common.Button{X: X, Y: Y}, colors.White, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)

		// Set the correct color for the select button.
		lightSelectedButton(eventsForLaunchpad, guiButtons, this)

		return

	} else {
		// Start this sequence.
		if debug {
			fmt.Printf("Start Sequence %d \n", Y)
		}

		// Stop the music trigger.
		sequences[this.SelectedSequence].MusicTrigger = false

		// If strobing stop it.
		if this.Strobe[this.SelectedSequence] {
			this.Strobe[this.SelectedSequence] = false
			StopStrobe(this, eventsForLaunchpad, guiButtons, commandChannels)
		}

		// Start the sequence.
		cmd := common.Command{
			Action: common.Start,
		}
		common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)
		common.LightLamp(common.Button{X: X, Y: Y}, colors.Green, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)

		this.Running[this.SelectedSequence] = true
		this.Functions[this.SelectedSequence][common.Function6_Static_Gobo].State = false
		this.Functions[this.SelectedSequence][common.Function8_Music_Trigger].State = false

		this.Static[this.SelectedSequence] = false
		this.SelectedMode[this.SelectedSequence] = NORMAL

		HandleSelect(sequences, this, eventsForLaunchpad, commandChannels, guiButtons)

		return
	}
}
