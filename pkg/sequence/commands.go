// Copyright (C) 2022,2023,2024 dhowlett99.
// This is the dmxlights main sequencers which processes commands.
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

func processCommands(mySequenceNumber int, sequence *common.Sequence, channels common.Channels, switchChannels []common.SwitchChannel, fixtureStepChannels []chan common.FixtureCommand, eventsForLauchpad chan common.ALight, guiButtons chan common.ALight) {

	// Clear all fixtures.
	if sequence.Clear {
		if debug {
			fmt.Printf("%d: Clear\n", mySequenceNumber)
		}
		clearSequence(mySequenceNumber, sequence, fixtureStepChannels)
		sequence.Clear = false
	}

	// Show all switches.
	if sequence.PlaySwitchOnce && !sequence.PlaySingleSwitch && !sequence.Override && sequence.Type == "switch" {
		if debug {
			fmt.Printf("%d: Show All Switches\n", mySequenceNumber)
		}
		showAllSwitches(mySequenceNumber, sequence, fixtureStepChannels, eventsForLauchpad, guiButtons)
		sequence.PlaySwitchOnce = false
	}

	// Show the selected switch.
	if sequence.PlaySwitchOnce && sequence.PlaySingleSwitch && !sequence.Override && sequence.Type == "switch" {
		if debug {
			fmt.Printf("%d: Show Single Switch\n", mySequenceNumber)
		}
		showSelectedSwitch(mySequenceNumber, sequence, fixtureStepChannels, eventsForLauchpad, guiButtons)
		sequence.PlaySwitchOnce = false
		sequence.PlaySingleSwitch = false
	}

	// Override the selected switch.
	if sequence.PlaySwitchOnce && sequence.Override && sequence.Type == "switch" {
		if debug {
			fmt.Printf("%d: Override Single Switch\n", mySequenceNumber)
		}
		overrideSwitch(mySequenceNumber, sequence, switchChannels)
		sequence.PlaySwitchOnce = false
		sequence.Override = false
	}

	// Start flood.
	if sequence.StartFlood && sequence.FloodPlayOnce && sequence.Type != "switch" {
		if debug {
			fmt.Printf("%d: Start Flood\n", mySequenceNumber)
		}
		startFlood(mySequenceNumber, sequence, fixtureStepChannels)
		if sequence.Chase {
			sequence.SaveChase = true
		}
		sequence.Chase = false
		sequence.StartFlood = false
		sequence.FloodPlayOnce = false
	}

	// Stop flood.
	if sequence.StopFlood && sequence.FloodPlayOnce && sequence.Type != "switch" {
		if debug {
			fmt.Printf("%d: Stop Flood\n", mySequenceNumber)
		}
		if sequence.SaveChase {
			sequence.SaveChase = false
			sequence.Chase = true
		}
		stopFlood(mySequenceNumber, sequence, fixtureStepChannels)
		sequence.StopFlood = false
		sequence.FloodPlayOnce = false
	}

	// Sequence in static mode.
	if sequence.PlayStaticOnce && sequence.Static && !sequence.StartFlood && sequence.Type != "switch" {
		if debug {
			fmt.Printf("%d: Start Static\n", mySequenceNumber)
		}
		startStatic(mySequenceNumber, sequence, channels, fixtureStepChannels)
		sequence.PlayStaticOnce = false
	}

	// Turn static mode off.
	if sequence.PlayStaticOnce && !sequence.Static && !sequence.StartFlood && sequence.Type != "switch" {
		if debug {
			fmt.Printf("%d: Stop Static\n", mySequenceNumber)
		}
		stopStatic(mySequenceNumber, sequence, channels, fixtureStepChannels)
		sequence.PlayStaticOnce = false
	}
}
