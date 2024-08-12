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

	"github.com/dhowlett99/dmxlights/pkg/common"
)

func showAllSwitches(mySequenceNumber int, sequence *common.Sequence, fixtureStepChannels []chan common.FixtureCommand, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight) {
	if debug {
		fmt.Printf("sequence %d Play all switches mode\n", mySequenceNumber)
	}
	// Show initial state of switches
	for switchNumber := 0; switchNumber < len(sequence.Switches); switchNumber++ {
		setSwitchLamp(*sequence, switchNumber, eventsForLaunchpad, guiButtons)
		setSwitchDMX(*sequence, switchNumber, fixtureStepChannels)
	}
	sequence.PlaySwitchOnce = false
}

func showSelectedSwitch(mySequenceNumber int, sequence *common.Sequence, fixtureStepChannels []chan common.FixtureCommand, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight) {
	if debug {
		fmt.Printf("%d: Play single switch number %d\n", mySequenceNumber, sequence.CurrentSwitch)
	}

	// Dim the last lamp.
	if sequence.CurrentSwitch != sequence.LastSwitchSelected || !sequence.FocusSwitch {
		// Clear the last selected switch.
		newSwitch := sequence.Switches[sequence.LastSwitchSelected]
		newSwitch.Selected = false
		sequence.Switches[sequence.LastSwitchSelected] = newSwitch
		setSwitchLamp(*sequence, sequence.LastSwitchSelected, eventsForLaunchpad, guiButtons)
	}

	// Now show the current switch state.
	if sequence.StepSwitch {
		// This is the second press so actually switch and send the DMX command.
		setSwitchLamp(*sequence, sequence.CurrentSwitch, eventsForLaunchpad, guiButtons)
		setSwitchDMX(*sequence, sequence.CurrentSwitch, fixtureStepChannels)
	} else {
		// first time we presses this switch button just move the focus here and use full brightness to indicate we
		// are the selected sequence and selected switch.
		setSwitchLamp(*sequence, sequence.CurrentSwitch, eventsForLaunchpad, guiButtons)
	}

	sequence.LastSwitchSelected = sequence.CurrentSwitch

	sequence.PlaySwitchOnce = false
	sequence.PlaySingleSwitch = false
	sequence.OverrideSpeed = false
}
