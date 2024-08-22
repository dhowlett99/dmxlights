// Copyright (C) 2022, 2023 , 2024 dhowlett99.
// This is button processor, used by the launchpad and gui interfaces.
// This file processes master brightness buttons and controls their actions.
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

func decraseBrightness(X int, Y int, this *CurrentState, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command) {

	if debug {
		fmt.Printf("Brightness Down \n")
	}

	buttonTouched(common.Button{X: X, Y: Y}, colors.White, colors.Cyan, eventsForLaunchpad, guiButtons)

	this.MasterBrightness = this.MasterBrightness - 10
	if this.MasterBrightness < 0 {
		this.MasterBrightness = 0
	}
	cmd := common.Command{
		Action: common.Master,
		Args: []common.Arg{
			{Name: "Master", Value: this.MasterBrightness},
		},
	}
	common.SendCommandToAllSequence(cmd, commandChannels)

	// Update the status bar
	common.UpdateStatusBar(fmt.Sprintf("Master %02d", this.MasterBrightness), "master", false, guiButtons)
}

func increaseBrightness(X int, Y int, this *CurrentState, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command) {

	if debug {
		fmt.Printf("Brightness Up \n")
	}

	buttonTouched(common.Button{X: X, Y: Y}, colors.White, colors.Cyan, eventsForLaunchpad, guiButtons)

	this.MasterBrightness = this.MasterBrightness + 10
	if this.MasterBrightness > common.MAX_DMX_BRIGHTNESS {
		this.MasterBrightness = common.MAX_DMX_BRIGHTNESS
	}
	cmd := common.Command{
		Action: common.Master,
		Args: []common.Arg{
			{Name: "Master", Value: this.MasterBrightness},
		},
	}
	common.SendCommandToAllSequence(cmd, commandChannels)

	// Update the status bar
	common.UpdateStatusBar(fmt.Sprintf("Master %02d", this.MasterBrightness), "master", false, guiButtons)

}
