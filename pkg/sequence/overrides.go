// Copyright (C) 2022,2023 dhowlett99.
// This is the dmxlights main sequencer responsible for controlling all
// of the fixtures in a group.
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

package sequence

import (
	"fmt"
	"time"

	"github.com/dhowlett99/dmxlights/pkg/common"
)

func overrideSwitch(mySequenceNumber int, sequence *common.Sequence, switchChannels []common.SwitchChannel) {

	// Override the selected switch speed.
	if sequence.OverrideSpeed {
		if debug {
			fmt.Printf("sequence %d Override switch number %d Speed %d \n", mySequenceNumber, sequence.CurrentSwitch, sequence.Switches[sequence.CurrentSwitch].Override.Speed)
		}
		// Send a message to the selected switch device.
		cmd := common.Command{
			Action: common.UpdateSpeed,
			Args: []common.Arg{
				// Add one since we count from 0
				{Name: "Speed", Value: sequence.Switches[sequence.CurrentSwitch].Override.Speed},
			},
		}
		select {
		case switchChannels[sequence.CurrentSwitch+1].CommandChannel <- cmd:
		case <-time.After(10 * time.Millisecond):
		}
		return
	}

	if sequence.OverrideShift {

		if debug {
			fmt.Printf("sequence %d Override switch number %d Shift %d \n", mySequenceNumber, sequence.CurrentSwitch, sequence.Switches[sequence.CurrentSwitch].Override.Shift)
		}
		// Send a message to the selected switch device.
		cmd := common.Command{
			Action: common.UpdateRGBShift,
			Args: []common.Arg{
				// Add one since we count from 0
				{Name: "Shift", Value: sequence.Switches[sequence.CurrentSwitch].Override.Shift},
			},
		}
		select {
		case switchChannels[sequence.CurrentSwitch+1].CommandChannel <- cmd:
		case <-time.After(10 * time.Millisecond):
		}
		sequence.PlaySwitchOnce = false
		sequence.OverrideShift = false
		return
	}

	// Override the selected switch size.
	if sequence.OverrideSize {

		if debug {
			fmt.Printf("sequence %d Override switch number %d Size %d \n", mySequenceNumber, sequence.CurrentSwitch, sequence.Switches[sequence.CurrentSwitch].Override.Size)
		}
		// Send a message to the selected switch device.
		cmd := common.Command{
			Action: common.UpdateRGBSize,
			Args: []common.Arg{
				// Add one since we count from 0
				{Name: "Size", Value: sequence.Switches[sequence.CurrentSwitch].Override.Size},
			},
		}
		select {
		case switchChannels[sequence.CurrentSwitch+1].CommandChannel <- cmd:
		case <-time.After(10 * time.Millisecond):
		}
		sequence.PlaySwitchOnce = false
		sequence.OverrideSize = false
		return
	}

	// Override the selected switch fade size.
	if sequence.OverrideFade {

		if debug {
			fmt.Printf("sequence %d Override switch number %d Fade %d \n", mySequenceNumber, sequence.CurrentSwitch, sequence.Switches[sequence.CurrentSwitch].Override.Fade)
		}
		// Send a message to the selected switch device.
		cmd := common.Command{
			Action: common.UpdateRGBFadeSpeed,
			Args: []common.Arg{
				// Add one since we count from 0
				{Name: "Fade", Value: sequence.Switches[sequence.CurrentSwitch].Override.Fade},
			},
		}
		select {
		case switchChannels[sequence.CurrentSwitch+1].CommandChannel <- cmd:
		case <-time.After(10 * time.Millisecond):
		}
		sequence.PlaySwitchOnce = false
		sequence.OverrideFade = false
		return
	}
}
