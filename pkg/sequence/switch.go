// Copyright (C) 2022,2023,2024 dhowlett99.
// This is the dmxlights main sequencers switch functions.
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

// Set the button color for the selected switch.
// Will also change the brightness to highlight the last selected switch.
func setSwitchLamp(sequence common.Sequence, switchNumber int, eventsForLauchpad chan common.ALight, guiButtons chan common.ALight) {

	swiTch := sequence.Switches[switchNumber]

	if debug {
		fmt.Printf("%d: switchNumber %d current %d selected %t\n", sequence.Number, swiTch.Number, swiTch.CurrentPosition, swiTch.Selected)
	}

	state := swiTch.States[swiTch.CurrentPosition]

	// Use the button color for this state to light the correct color on the launchpad.
	color, _ := common.GetRGBColorByName(state.ButtonColor)
	var brightness int
	if swiTch.Selected {
		brightness = common.MAX_DMX_BRIGHTNESS
	} else {
		brightness = common.MAX_DMX_BRIGHTNESS / 8
	}

	common.LightLamp(common.Button{X: switchNumber, Y: sequence.Number}, color, brightness, eventsForLauchpad, guiButtons)

	// Label the switch.
	common.LabelButton(switchNumber, sequence.Number, swiTch.Label+"\n"+state.Label, guiButtons)

}

// Set the DMX parameters for the selected switch.
func setSwitchDMX(sequence common.Sequence, switchNumber int, fixtureStepChannels []chan common.FixtureCommand) {

	swiTch := sequence.Switches[switchNumber]

	if debug {
		fmt.Printf("switchNumber %d current %d selected %t speed %d\n", swiTch.Number, swiTch.CurrentPosition, swiTch.Selected, sequence.Switches[swiTch.Number-1].Override.Speed)
	}

	state := swiTch.States[swiTch.CurrentPosition]

	// Now send a message to the fixture to play all the values for this state.
	command := common.FixtureCommand{
		Master:             sequence.Master,
		Blackout:           sequence.Blackout,
		Type:               sequence.Type,
		Label:              sequence.Label,
		SequenceNumber:     sequence.Number,
		SwiTch:             swiTch,
		State:              state,
		CurrentSwitchState: swiTch.CurrentPosition,
		MasterChanging:     sequence.MasterChanging,
		RGBFade:            sequence.RGBFade,
		Override:           sequence.Switches[swiTch.CurrentPosition].Override,
	}

	// Send a message to the fixture to operate the switch.
	fixtureStepChannels[switchNumber] <- command

}
