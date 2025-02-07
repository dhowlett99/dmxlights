// Copyright (C) 2022, 2023 , 2024 dhowlett99.
// This is button processor, used by the launchpad and gui interfaces.
// This file processes flood buttons and controls their actions.
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
	"github.com/dhowlett99/dmxlights/pkg/presets"
)

func FloodOff(numberSequences int, this *CurrentState, commandChannels []chan common.Command, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight) {

	// Turn the flood button back to white.
	common.LightLamp(common.FLOOD_BUTTON, colors.White, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)

	// Send a message to stop flood.
	cmd := common.Command{
		Action: common.StopFlood,
		Args: []common.Arg{
			{Name: "Stop Flood", Value: false},
		},
	}
	common.SendCommandToAllSequence(cmd, commandChannels)

	this.Flood = false

	// Preserve this.Blackout.
	if !this.Blackout {
		cmd := common.Command{
			Action: common.Normal,
		}
		common.SendCommandToAllSequence(cmd, commandChannels)
	}

	// Retore any shutter chaser mode.
	if this.ScannerChaser[this.SelectedSequence] {
		// Tell the scanner sequence to hide their indicator lamps.
		cmd = common.Command{
			Action: common.Hide,
			Args: []common.Arg{
				{Name: "Hide", Value: this.ScannerChaser[this.SelectedSequence]},
			},
		}
		common.SendCommandToSequence(this.ScannerSequenceNumber, cmd, commandChannels)
	}

	// ReStart any sequences that were running before the flood.
	for sequenceNumber := 0; sequenceNumber < numberSequences; sequenceNumber++ {
		if this.Running[sequenceNumber] {
			cmd = common.Command{
				Action: common.Start,
			}
			common.SendCommandToSequence(sequenceNumber, cmd, commandChannels)
		}
	}
}

func floodOn(numberSequences int, this *CurrentState, commandChannels []chan common.Command, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight) {

	// Remember which sequence is currently selected.
	this.LastSelectedSequence = this.SelectedSequence

	// Flash the flood button pink to indicate we're in flood.
	common.FlashLight(common.FLOOD_BUTTON, colors.Magenta, colors.White, eventsForLaunchpad, guiButtons)

	var cmd common.Command

	// Stop running sequences.
	for sequenceNumber := 0; sequenceNumber < numberSequences; sequenceNumber++ {
		if this.Running[sequenceNumber] {
			cmd = common.Command{
				Action: common.Stop,
				Args: []common.Arg{
					{Name: "Stop", Value: true},
				},
			}
			common.SendCommandToSequence(sequenceNumber, cmd, commandChannels)
		}
	}

	// Start flood.
	cmd = common.Command{
		Action: common.Flood,
		Args: []common.Arg{
			{Name: "StartFlood", Value: true},
		},
	}
	common.SendCommandToAllSequence(cmd, commandChannels)

	this.Flood = true
}

func toggleFlood(sequences []*common.Sequence, X int, Y int, this *CurrentState, commandChannels []chan common.Command, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight) {

	if debug {
		fmt.Printf("Start FLood X:%d Y:%d\n", X, Y)
	}

	// Turn off the flashing save button
	this.SavePreset = false
	common.LightLamp(common.SAVE_BUTTON, colors.White, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)

	// Shutdown any function bars.
	clearAllModes(sequences, this)
	HandleSelect(sequences, this, eventsForLaunchpad, commandChannels, guiButtons)

	if !this.Flood { // We're not already in flood so lets ask the sequence to flood.
		if debug {
			fmt.Printf("FLOOD ON\n")
		}
		// Find the currently selected preset and save it's location.
		for location, preset := range this.PresetsStore {
			if preset.State && preset.Selected {
				this.PresetsStore[location] = presets.Preset{State: preset.State, Selected: false, Label: preset.Label, ButtonColor: preset.ButtonColor}
				presets.RefreshPresets(eventsForLaunchpad, guiButtons, this.PresetsStore)
				this.LastPreset = &location
				break
			}
		}
		floodOn(len(sequences), this, commandChannels, eventsForLaunchpad, guiButtons)
		return
	}
	if this.Flood { // If we are flood already then tell the sequence to stop flood.
		if debug {
			fmt.Printf("FLOOD OFF\n")
		}
		// Restore the last preset
		if this.LastPreset != nil {
			lastPreset := this.PresetsStore[*this.LastPreset]
			this.PresetsStore[*this.LastPreset] = presets.Preset{State: lastPreset.State, Selected: true, Label: lastPreset.Label, ButtonColor: lastPreset.ButtonColor}
			presets.RefreshPresets(eventsForLaunchpad, guiButtons, this.PresetsStore)
		}
		FloodOff(len(sequences), this, commandChannels, eventsForLaunchpad, guiButtons)
		return
	}
}
