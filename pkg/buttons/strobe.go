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
