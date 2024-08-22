// Copyright (C) 2022, 2023 , 2024 dhowlett99.
// This is button processor, used by the launchpad and gui interfaces.
// This file processes the scanner pan and tilt adjustments.
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

func upArrow(X int, Y int, this *CurrentState, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command) {

	if debug {
		fmt.Printf("UP ARROW\n")
	}

	buttonTouched(common.Button{X: X, Y: Y}, colors.Cyan, colors.White, eventsForLaunchpad, guiButtons)

	this.OffsetTilt = this.OffsetTilt + 5

	if this.OffsetTilt > 255 {
		this.OffsetTilt = 255
	}
	// Clear the sequence colors for this sequence.
	cmd := common.Command{
		Action: common.UpdateOffsetTilt,
		Args: []common.Arg{
			{Name: "OffsetTilt", Value: this.OffsetTilt},
		},
	}
	common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

	// Update status bar.
	common.UpdateStatusBar(fmt.Sprintf("Tilt %02d", this.OffsetTilt), "tilt", false, guiButtons)
}

func downArrow(X int, Y int, this *CurrentState, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command) {

	if debug {
		fmt.Printf("DOWN ARROW\n")
	}

	buttonTouched(common.Button{X: X, Y: Y}, colors.Cyan, colors.White, eventsForLaunchpad, guiButtons)

	this.OffsetTilt = this.OffsetTilt - 5

	if this.OffsetTilt < 0 {
		this.OffsetTilt = 0
	}
	// Clear the sequence colors for this sequence.
	cmd := common.Command{
		Action: common.UpdateOffsetTilt,
		Args: []common.Arg{
			{Name: "OffsetTilt", Value: this.OffsetTilt},
		},
	}
	common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

	// Update status bar.
	common.UpdateStatusBar(fmt.Sprintf("Tilt %02d", this.OffsetTilt), "tilt", false, guiButtons)
}

func leftArrow(X int, Y int, this *CurrentState, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command) {

	if debug {
		fmt.Printf("LEFT ARROW\n")
	}

	buttonTouched(common.Button{X: X, Y: Y}, colors.Cyan, colors.White, eventsForLaunchpad, guiButtons)

	this.OffsetPan = this.OffsetPan + 5

	if this.OffsetPan > 255 {
		this.OffsetPan = 255
	}

	// Clear the sequence colors for this sequence.
	cmd := common.Command{
		Action: common.UpdateOffsetPan,
		Args: []common.Arg{
			{Name: "OffsetPan", Value: this.OffsetPan},
		},
	}
	common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

	// Update status bar.
	common.UpdateStatusBar(fmt.Sprintf("Pan %02d", this.OffsetPan), "pan", false, guiButtons)

}

func rightArrow(X int, Y int, this *CurrentState, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command) {

	if debug {
		fmt.Printf("RIGHT ARROW\n")
	}

	buttonTouched(common.Button{X: X, Y: Y}, colors.Cyan, colors.White, eventsForLaunchpad, guiButtons)

	this.OffsetPan = this.OffsetPan - 5

	if this.OffsetPan < 0 {
		this.OffsetPan = 0
	}

	// Clear the sequence colors for this sequence.
	cmd := common.Command{
		Action: common.UpdateOffsetPan,
		Args: []common.Arg{
			{Name: "OffsetPan", Value: this.OffsetPan},
		},
	}
	common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

	// Update status bar.
	common.UpdateStatusBar(fmt.Sprintf("Pan %02d", this.OffsetPan), "pan", false, guiButtons)
}
