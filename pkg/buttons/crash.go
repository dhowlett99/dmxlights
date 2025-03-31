// Copyright (C) 2022, 2023 , 2024 dhowlett99.
// This is button processor, used by the launchpad and gui interfaces.
// This file processes the reboot messages from the Novation Launchpad
// uses them to refresh the buttons in the event of a crash.
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
	"image/color"
	"time"

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/presets"
)

func handleLaunchPadCrash(sequences []*common.Sequence, X int, Y int, this *CurrentState, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command) {

	// The Novation Launchpad is not designed for the number of MIDI
	// Events we send when all the sequences are chasing at top
	// Speed, so we look out for the crys for help when the Launchpad
	// Crashes. When we see the following three events we reset the pad.
	// As this happens so quickly the user should be unware that a crash
	// has taken place.
	if X == -1 && Y == 8 {
		this.Crash1 = true
		return
	}
	if X == 1 && Y == 8 && this.Crash1 {
		this.Crash2 = true
		return
	}
	// Crash 2 message has appeared and this isn't a pad program ack.
	if X != 0 && Y == 8 && this.Crash2 {
		// Start a supervisor thread which will reset the launchpad every 1/2 second.
		time.Sleep(200 * time.Millisecond)
		if this.LaunchPadConnected {
			this.Pad.Program()
		}
		staticColors := []color.RGBA{}
		for _, button := range sequences[0].StaticColors {
			staticColors = append(staticColors, button.Color)
		}
		InitButtons(this, sequences[0].SequenceColors, staticColors, eventsForLaunchpad, guiButtons)

		// Show the static and switch settings.
		cmd := common.Command{
			Action: common.Reveal,
		}
		common.SendCommandToAllSequence(cmd, commandChannels)

		// Show the presets again.
		presets.RefreshPresets(eventsForLaunchpad, guiButtons, this.PresetsStore)
		this.Crash1 = false
		this.Crash2 = false
		return
	}
}
