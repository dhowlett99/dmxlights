// Copyright (C) 2022, 2023 , 2024 dhowlett99.
// This is button processor, used by the launchpad and gui interfaces.
// This file processes the sound sensitivity buttons and controls their actions.
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

func increaseSensitivity(this *CurrentState, X int, Y int, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight) {

	if debug {
		fmt.Printf("Sound Up %f\n", this.SoundGain)
	}

	buttonTouched(common.Button{X: X, Y: Y}, colors.White, colors.Cyan, eventsForLaunchpad, guiButtons)

	this.SoundGain = this.SoundGain - 0.01
	if this.SoundGain < -0.04 {
		this.SoundGain = -0.04
	}
	for _, trigger := range this.SoundTriggers {
		trigger.Gain = this.SoundGain
	}
	// Update the status bar
	sensitivity := common.FindSensitivity(this.SoundGain)
	common.UpdateStatusBar(fmt.Sprintf("Sensitivity %02d", sensitivity), "sensitivity", false, guiButtons)
}

func decreaseSensitivity(this *CurrentState, X int, Y int, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight) {

	if debug {
		fmt.Printf("Sound Down%f\n", this.SoundGain)
	}

	buttonTouched(common.Button{X: X, Y: Y}, colors.White, colors.Cyan, eventsForLaunchpad, guiButtons)

	this.SoundGain = this.SoundGain + 0.01
	if this.SoundGain > 0.09 {
		this.SoundGain = 0.09
	}
	for _, trigger := range this.SoundTriggers {
		trigger.Gain = this.SoundGain
	}
	// Update the status bar
	sensitivity := common.FindSensitivity(this.SoundGain)
	common.UpdateStatusBar(fmt.Sprintf("Sensitivity %02d", sensitivity), "sensitivity", false, guiButtons)
}
