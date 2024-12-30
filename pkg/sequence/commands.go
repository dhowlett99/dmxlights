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
	"github.com/dhowlett99/dmxlights/pkg/fixture"
	"github.com/dhowlett99/dmxlights/pkg/sound"
	"github.com/oliread/usbdmx/ft232"
)

func processCommands(sequence *common.Sequence, channels common.Channels, switchChannels []common.SwitchChannel, fixtureStepChannels []chan common.FixtureCommand, soundConfig *sound.SoundConfig, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, fixturesConfig *fixture.Fixtures, dmxController *ft232.DMXController, dmxInterfacePresent bool) []chan common.FixtureCommand {

	// Load new set of fixtures and setup fixture threads and channels to those fixtures.
	if sequence.LoadNewFixtures {

		if debug {
			fmt.Printf("%d: Load New Fixtures\n", sequence.Number)
		}
		fixtureStepChannels = LoadNewFixtures(sequence, fixtureStepChannels, eventsForLaunchpad, guiButtons, switchChannels, channels.SoundTriggers, soundConfig, dmxController, fixturesConfig, dmxInterfacePresent)
	}

	// Clear all fixtures.
	if sequence.Clear {
		if debug {
			fmt.Printf("%d: Clear\n", sequence.Number)
		}
		clearSequence(sequence.Number, sequence, fixtureStepChannels)
		sequence.Clear = false
	}

	// Show all switches.
	if sequence.PlaySwitchOnce && !sequence.PlaySingleSwitch && !sequence.Override && sequence.Type == "switch" {
		if debug {
			fmt.Printf("%d: Show All Switches\n", sequence.Number)
		}
		showAllSwitches(sequence.Number, sequence, fixtureStepChannels, eventsForLaunchpad, guiButtons)
		sequence.PlaySwitchOnce = false
	}

	// Show the selected switch.
	if sequence.PlaySwitchOnce && sequence.PlaySingleSwitch && !sequence.Override && sequence.Type == "switch" {
		if debug {
			fmt.Printf("%d: Show Single Switch\n", sequence.Number)
		}
		showSelectedSwitch(sequence.Number, sequence, fixtureStepChannels, eventsForLaunchpad, guiButtons)
		sequence.PlaySwitchOnce = false
		sequence.PlaySingleSwitch = false
	}

	// Override the selected switch.
	if sequence.PlaySwitchOnce && sequence.Override && sequence.Type == "switch" {
		if debug {
			fmt.Printf("%d: Override Single Switch=%d Override=%+v\n", sequence.Number, sequence.CurrentSwitch, sequence.Switches[sequence.CurrentSwitch].Override)
		}

		// Get switch data variables setup.
		swiTch := sequence.Switches[sequence.CurrentSwitch]
		state := swiTch.States[swiTch.CurrentPosition]

		// Pass through the override command.
		command := common.FixtureCommand{
			Type:           "override",
			Master:         sequence.Master,
			Blackout:       sequence.Blackout,
			CurrentSwitch:  sequence.CurrentSwitch,
			SwiTch:         sequence.Switches[sequence.CurrentSwitch],
			Override:       sequence.Switches[sequence.CurrentSwitch].Override,
			State:          state,
			RGBFade:        sequence.RGBFade,
			MasterChanging: sequence.MasterChanging,
		}
		fixtureStepChannels[sequence.CurrentSwitch] <- command

		sequence.PlaySwitchOnce = false
		sequence.Override = false
	}

	// Start flood.
	if sequence.StartFlood && sequence.FloodPlayOnce && sequence.Type != "switch" {
		if debug {
			fmt.Printf("%d: Start Flood\n", sequence.Number)
		}
		startFlood(sequence.Number, sequence, fixtureStepChannels)
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
			fmt.Printf("%d: Stop Flood\n", sequence.Number)
		}
		if sequence.SaveChase {
			sequence.SaveChase = false
			sequence.Chase = true
		}
		stopFlood(sequence.Number, sequence, fixtureStepChannels)
		sequence.StopFlood = false
		sequence.FloodPlayOnce = false
	}

	// Sequence in static mode.
	if sequence.PlayStaticOnce && sequence.Static && !sequence.StartFlood && sequence.Type != "switch" {
		if debug {
			fmt.Printf("%d: Start Static\n", sequence.Number)
		}
		startStatic(sequence.Number, sequence, channels, fixtureStepChannels)
		sequence.PlayStaticOnce = false
	}

	// Turn static mode off.
	if sequence.PlayStaticOnce && !sequence.Static && !sequence.StartFlood && sequence.Type != "switch" && !sequence.Run {
		if debug {
			fmt.Printf("%d: Stop Static\n", sequence.Number)
		}
		stopStatic(sequence.Number, sequence, channels, fixtureStepChannels)
		sequence.PlayStaticOnce = false
	}

	return fixtureStepChannels
}
