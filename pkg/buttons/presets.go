// Copyright (C) 2022, 2023, 2024, 2025 dhowlett99.
// This is button processor, used by the launchpad and gui interfaces.
// This file processes preset buttons and controls their actions.
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
	"time"

	"github.com/dhowlett99/dmxlights/pkg/colors"
	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/config"
	"github.com/dhowlett99/dmxlights/pkg/presets"
)

func SavePresetOff(this *CurrentState, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight) {
	this.SavePreset = false
	// turn off the save button from flashing.
	common.LightLamp(common.SAVE_BUTTON, colors.White, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)
}

func togglePresetSaveMode(numberSequences int, this *CurrentState, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command) {

	if debug {
		fmt.Printf("Save Mode\n")
	}

	if this.SavePreset { // Turn the save mode off.
		this.SavePreset = false
		presets.RefreshPresets(eventsForLaunchpad, guiButtons, this.PresetsStore)
		common.LightLamp(common.SAVE_BUTTON, colors.White, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)
		return
	}
	this.SavePreset = true
	if this.Flood { // Turn off flood.
		floodOff(numberSequences, this, commandChannels, eventsForLaunchpad, guiButtons)
	}
	presets.RefreshPresets(eventsForLaunchpad, guiButtons, this.PresetsStore)
	common.FlashLight(common.SAVE_BUTTON, colors.Magenta, colors.White, eventsForLaunchpad, guiButtons)
}

func recallPreset(sequences []*common.Sequence, X int, Y int, this *CurrentState, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command, updateChannels []chan common.Sequence) {

	if debug {
		fmt.Printf("recallPreset() Preset Pressed X:%d Y:%d\n", X, Y)
	}

	// Remove the button off offset.
	X = X - 100

	// We just pushed save this preset.
	if this.SavePreset {
		this.SavePreset = false
		return
	}

	// If this is a valid preset we are either recalling (short press) it or deleting it (long press)
	// If its not been set i.e. not valid we just ignore and return.
	if !this.PresetsStore[fmt.Sprint(X)+","+fmt.Sprint(Y)].State {
		return
	}

	// Stop the timer for this preset.
	elapsed := time.Since(*this.ButtonTimer)

	// Delete a preset - If the timer is longer than 1 seconds then we have a long press.
	if elapsed > 1*time.Second {

		if debug {
			fmt.Printf("Clear Preset X:%d Y:%d\n", X, Y)
		}

		// Delete the config file
		config.DeleteConfig(fmt.Sprintf("config%d.%d.json", X, Y))

		// Delete from preset store
		this.PresetsStore[fmt.Sprint(X)+","+fmt.Sprint(Y)] = presets.Preset{State: false, Selected: false, Label: "", ButtonColor: ""}

		// Update the copy of presets on disk.
		presets.SavePresets(this.PresetsStore)

		// Show presets again.
		presets.RefreshPresets(eventsForLaunchpad, guiButtons, this.PresetsStore)

	} else {
		if debug {
			fmt.Printf("Load Preset X:%d Y:%d\n", X, Y)
		}

		// Short press means load the preset.
		loadPreset(sequences, this, X, Y, commandChannels, eventsForLaunchpad, guiButtons, updateChannels)
		common.StartStaticSequences(sequences, commandChannels)
	}

}

func savePresets(sequences []*common.Sequence, X int, Y int, this *CurrentState, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command, updateChannels []chan common.Sequence, replyChannels []chan common.Sequence) {

	if debug {
		fmt.Printf("Ask For Config Y=%d X=%d\n", Y, X)
	}

	if this.ShowRGBColorPicker {
		this.ShowRGBColorPicker = false
		removeColorPicker(this, sequences, eventsForLaunchpad, guiButtons, commandChannels)
	}

	location := fmt.Sprint(X) + "," + fmt.Sprint(Y)

	if this.SavePreset {
		// S A V E - Ask all sequences for their current config and save in a file.

		current := this.PresetsStore[location]
		this.PresetsStore[location] = presets.Preset{State: true, Selected: true, Label: current.Label, ButtonColor: current.ButtonColor}
		this.LastPreset = &location

		config.AskToSaveConfig(commandChannels, replyChannels, X, Y)

		// turn off the save button from flashing.
		common.LightLamp(common.SAVE_BUTTON, colors.White, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)

		presets.SavePresets(this.PresetsStore)

		// clear any selected preset.
		for location, preset := range this.PresetsStore {
			if preset.State && preset.Selected {
				this.PresetsStore[location] = presets.Preset{State: preset.State, Selected: false, Label: preset.Label, ButtonColor: preset.ButtonColor}
			}
		}

		// Select this location and flash its button.
		this.PresetsStore[location] = presets.Preset{State: true, Selected: true, Label: current.Label, ButtonColor: current.ButtonColor}
		presets.RefreshPresets(eventsForLaunchpad, guiButtons, this.PresetsStore)

		if this.GUI {
			this.SavePreset = false
		}

	} else {
		// L O A D - Load config, but only if it exists in the presets map.
		if this.PresetsStore[location].State {

			if this.GUI { // GUI path.
				if this.SavePreset {
					this.SavePreset = false
				}
				loadPreset(sequences, this, X, Y, commandChannels, eventsForLaunchpad, guiButtons, updateChannels)
				common.StartStaticSequences(sequences, commandChannels)
			} else { // Launchpad path.
				// This is a valid preset we might be trying to load it or delete it.
				// Start a timer for this button.
				here := time.Now()
				this.ButtonTimer = &here
				// And wait for the button release.
			}
		}
	}
}
