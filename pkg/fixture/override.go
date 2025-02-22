// Copyright (C) 2022,2023,2024 dhowlett99.
// This is the fixture override processor.
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

package fixture

import (
	"fmt"
	"time"

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/oliread/usbdmx/ft232"
)

func overrideMiniSequencer(cmd common.FixtureCommand, switchChannels []common.SwitchChannel) {

	// Override the selected switch speed.
	if cmd.Override.Speed > 0 {

		if debug {
			fmt.Printf("Override switch number %d Speed %d \n", cmd.CurrentSwitch, cmd.SwiTch.Override.Speed)
		}
		// Send a message to the selected switch device.
		switchCommand := common.Command{
			Action: common.UpdateSpeed,
			Args: []common.Arg{
				// Add one since we count from 0
				{Name: "Speed", Value: cmd.SwiTch.Override.Speed},
			},
		}
		select {
		case switchChannels[cmd.CurrentSwitch+1].CommandChannel <- switchCommand:
		case <-time.After(10 * time.Millisecond):
		}
		return
	}

	if cmd.Override.Shift > 0 {

		if debug {
			fmt.Printf("Override switch number %d Shift %d \n", cmd.CurrentSwitch, cmd.SwiTch.Override.Shift)
		}
		// Send a message to the selected switch device.
		switchCommand := common.Command{
			Action: common.UpdateRGBShift,
			Args: []common.Arg{
				// Add one since we count from 0
				{Name: "Shift", Value: cmd.SwiTch.Override.Shift},
			},
		}
		select {
		case switchChannels[cmd.CurrentSwitch+1].CommandChannel <- switchCommand:
		case <-time.After(10 * time.Millisecond):
		}
		return
	}

	// Override the selected switch size.
	if cmd.Override.Size > 0 {

		if debug {
			fmt.Printf("Override switch number %d Size %d \n", cmd.CurrentSwitch, cmd.SwiTch.Override.Size)
		}
		// Send a message to the selected switch device.
		switchCommand := common.Command{
			Action: common.UpdateRGBSize,
			Args: []common.Arg{
				// Add one since we count from 0
				{Name: "Size", Value: cmd.SwiTch.Override.Size},
			},
		}
		select {
		case switchChannels[cmd.CurrentSwitch+1].CommandChannel <- switchCommand:
		case <-time.After(10 * time.Millisecond):
		}
		return
	}

	// Override the selected switch fade size.
	if cmd.Override.Fade > 0 {

		if debug {
			fmt.Printf("Override switch number %d Fade %d \n", cmd.CurrentSwitch, cmd.SwiTch.Override.Fade)
		}
		// Send a message to the selected switch device.
		switchCommand := common.Command{
			Action: common.UpdateRGBFadeSpeed,
			Args: []common.Arg{
				// Add one since we count from 0
				{Name: "Fade", Value: cmd.SwiTch.Override.Fade},
			},
		}
		select {
		case switchChannels[cmd.CurrentSwitch+1].CommandChannel <- switchCommand:
		case <-time.After(10 * time.Millisecond):
		}

		return
	}

	if cmd.Override.Rotate > 0 {

		if debug {
			fmt.Printf("Override switch number %d RotateSpeed %d \n", cmd.CurrentSwitch, cmd.SwiTch.Override.Rotate)
		}
		// Send a message to the selected switch device.
		switchCommand := common.Command{
			Action: common.UpdateRotateSpeed,
			Args: []common.Arg{
				// Add one since we count from 0
				{Name: "RotateSpeed", Value: cmd.SwiTch.Override.Rotate},
			},
		}
		select {
		case switchChannels[cmd.CurrentSwitch+1].CommandChannel <- switchCommand:
		case <-time.After(10 * time.Millisecond):
		}
		return
	}

	if cmd.Override.Color > 0 {

		if debug {
			fmt.Printf("Override switch number %d Color %d \n", cmd.CurrentSwitch, cmd.SwiTch.Override.Color)
		}
		// Send a message to the selected switch device.
		switchCommand := common.Command{
			Action: common.UpdateColors,
			Args: []common.Arg{
				// Add one since we count from 0
				{Name: "Color", Value: cmd.SwiTch.Override.Color},
			},
		}
		select {
		case switchChannels[cmd.CurrentSwitch+1].CommandChannel <- switchCommand:
		case <-time.After(10 * time.Millisecond):
		}
		return
	}

	if cmd.Override.Gobo > 0 {

		if debug {
			fmt.Printf("Override switch number %d Gobo %d \n", cmd.CurrentSwitch, cmd.SwiTch.Override.Gobo)
		}
		// Send a message to the selected switch device.
		switchCommand := common.Command{
			Action: common.UpdateGobo,
			Args: []common.Arg{
				// Add one since we count from 0
				{Name: "Gobo", Value: cmd.SwiTch.Override.Gobo},
			},
		}
		select {
		case switchChannels[cmd.CurrentSwitch+1].CommandChannel <- switchCommand:
		case <-time.After(10 * time.Millisecond):
		}
		return
	}
}

func overrideMiniSetter(cmd common.FixtureCommand, fixturesConfig *Fixtures, dmxController *ft232.DMXController, dmxInterfacePresent bool) {

	if debug {
		fmt.Printf("overrideMiniSetter Switch=%d Override=%+v\n", cmd.CurrentSwitch, cmd.SwiTch.Override)
	}

	state := cmd.SwiTch.States[cmd.SwiTch.CurrentPosition]
	useFixtureLabel := cmd.SwiTch.UseFixture
	if debug {
		fmt.Printf("overrideMiniSetter Switch=%d UseFixture=%s State=%+v\n", cmd.CurrentSwitch, useFixtureLabel, state)
	}

	// Find the details of the fixture for this switch.
	thisFixture, err := GetFixtureByLabel(useFixtureLabel, fixturesConfig)
	if err != nil {
		fmt.Printf("error %s\n", err.Error())
	}

	// Look for Master channel in this fixture identified by ID.
	masterChannel, err := FindChannelNumberByName(thisFixture, "Master")
	if err != nil && debug {
		fmt.Printf("warning! fixture %s: %s\n", thisFixture.Name, err)
	}

	for _, newSetting := range state.Settings {
		newMiniSetter(thisFixture, cmd.Override, newSetting, masterChannel, dmxController, cmd.Master, dmxInterfacePresent)
	}

}
