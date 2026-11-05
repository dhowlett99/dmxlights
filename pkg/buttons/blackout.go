// Copyright (C) 2022, 2023 , 2024 dhowlett99.
// This is button processor, used by the launchpad and gui interfaces.
// This file processes the blackout key.
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

func blackout(X int, Y int, this *CurrentState, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command) {

	if debug {
		fmt.Printf("BLACKOUT\n")
	}

	// Turn off the flashing save button
	this.SavePreset = false
	common.LightLamp(common.SAVE_BUTTON, colors.White, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)

	if !this.Blackout {
		if debug {
			fmt.Printf("BLACKOUT\n")
		}
		this.Blackout = true
		cmd := common.Command{
			Action: common.Blackout,
		}
		common.SendCommandToAllSequence(cmd, commandChannels)
		common.LightLamp(common.Button{X: X, Y: Y}, colors.Black, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)
		common.FlashLight(common.BLACKOUT_BUTTON, colors.Magenta, colors.White, eventsForLaunchpad, guiButtons)
	} else {
		if debug {
			fmt.Printf("NORMAL\n")
		}
		this.Blackout = false
		cmd := common.Command{
			Action: common.Normal,
		}
		common.SendCommandToAllSequence(cmd, commandChannels)
		common.LightLamp(common.Button{X: X, Y: Y}, colors.White, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)
	}
}
