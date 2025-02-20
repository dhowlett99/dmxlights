// Copyright (C) 2022,2023,2024,2025 dhowlett99.
// This is the dmxlights fixture receiver code, it receives messages from the sequence and
// decides what the fixture should do.
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
	"github.com/dhowlett99/dmxlights/pkg/sound"
	"github.com/oliread/usbdmx/ft232"
)

// FixtureReceivers are created by the sequence and are used to receive step instructions.
// Each FixtureReceiver knows which step they belong too and when triggered they start a fade up
// and fade down events which get sent to the launchpad lamps and the DMX fixtures.
func FixtureReceiver(
	myFixtureNumber int,
	fixtureStepChannel chan common.FixtureCommand,
	eventsForLaunchpad chan common.ALight,
	guiButtons chan common.ALight,
	switchChannels []common.SwitchChannel,
	soundTriggers []*common.Trigger,
	soundConfig *sound.SoundConfig,
	dmxController *ft232.DMXController,
	fixturesConfig *Fixtures,
	dmxInterfacePresent bool) {

	if debug {
		fmt.Printf("FixtureReceiver Started %d\n", myFixtureNumber)
	}

	// Used for static fades, remember the last color.
	var lastColor common.LastColor

	stopFadeUp := make(chan bool)
	stopFadeDown := make(chan bool)

	// Loop waiting for configuration.
	for {

		// Wait for first step
		cmd := <-fixtureStepChannel

		// Stop fixture channel.
		if cmd.Stop {
			if debug {
				fmt.Printf("Fixture %d Stopping\n", myFixtureNumber)
			}
			break
		}

		if cmd.Blackout {
			// Soft fade downs should be disabled for blackout.
			lastColor.RGBColor = cmd.LastColor
			lastColor.ScannerColor = 0
			// Make sure we are blacked out.
			cmd.Master = 0
		}

		switch {

		case cmd.Type == "override":
			if debug {
				fmt.Printf("FixtureReceiver: override %+v\n", cmd)
			}

			// We have actions we can override.
			if len(cmd.State.Actions) > 0 {
				// Overide is done differently for chase actions.
				if cmd.State.Actions[0].Mode == "Chase" {
					// Send a message to the mini sequencer in chase mode.
					overrideMiniSequencer(cmd, switchChannels)
				} else {
					// Call the minisequencer directly in in Static or Control Mode with the new override values.
					lastColor = MapSwitchFixture(cmd.SwiTch, cmd.State, cmd.Override, cmd.RGBFade, dmxController, fixturesConfig, cmd.Blackout, cmd.Master, cmd.Master, cmd.MasterChanging, lastColor, switchChannels, soundTriggers, soundConfig, dmxInterfacePresent, eventsForLaunchpad, guiButtons, fixtureStepChannel)
				}
			}
			// We have settings we can override.
			if len(cmd.State.Settings) > 0 {
				// Call the mini setter. Step through all the settings.
				overrideMiniSetter(cmd, fixturesConfig, dmxController, dmxInterfacePresent)
			}

			continue

		case cmd.Type == "lastColor":
			if debug {
				fmt.Printf("%d:%d LastColor set to %s\n", cmd.SequenceNumber, myFixtureNumber, common.GetColorNameByRGB(cmd.LastColor))
			}
			lastColor.RGBColor = cmd.LastColor
			lastColor.ScannerColor = 0
			continue

		case cmd.Type == "switch":
			if debug {
				fmt.Printf("%d:%d Activate switch number %d name %s Postition %d Speed %d Shift %d\n", cmd.SequenceNumber, myFixtureNumber, cmd.SwiTch.Number, cmd.SwiTch.Name, cmd.SwiTch.CurrentPosition, cmd.Override.Speed, cmd.Override.Shift)
			}
			// Since this is a swich being changed, we clear any override for this sequence.
			cmd.Override = common.Override{}
			lastColor = MapSwitchFixture(cmd.SwiTch, cmd.State, cmd.Override, cmd.RGBFade, dmxController, fixturesConfig, cmd.Blackout, cmd.Master, cmd.Master, cmd.MasterChanging, lastColor, switchChannels, soundTriggers, soundConfig, dmxInterfacePresent, eventsForLaunchpad, guiButtons, fixtureStepChannel)
			continue

		case cmd.Clear:
			if debug {
				fmt.Printf("%d:%d Clear %t Blackout %t\n", cmd.SequenceNumber, myFixtureNumber, cmd.Clear, cmd.Blackout)
			}
			lastColor = clearFixture(myFixtureNumber, cmd, stopFadeDown, stopFadeUp, fixturesConfig, dmxController, dmxInterfacePresent)
			continue

		case cmd.StartFlood:
			if debug {
				fmt.Printf("%d:%d StartFlood\n", cmd.SequenceNumber, myFixtureNumber)
			}
			// Stop any running fade ups.
			select {
			case stopFadeUp <- true:
			case <-time.After(100 * time.Millisecond):
			}
			// Stop any running fade downs.
			select {
			case stopFadeDown <- true:
			case <-time.After(100 * time.Millisecond):
			}
			lastColor = startFlood(myFixtureNumber, cmd, fixturesConfig, eventsForLaunchpad, guiButtons, dmxController, dmxInterfacePresent)
			continue

		case cmd.StopFlood:
			if debug {
				fmt.Printf("%d:%d StopFlood\n", cmd.SequenceNumber, myFixtureNumber)
			}
			lastColor = stopFlood(myFixtureNumber, cmd, fixturesConfig, eventsForLaunchpad, guiButtons, dmxController, dmxInterfacePresent)
			continue

		case cmd.RGBStaticOn:
			if debug {
				fmt.Printf("%d:%d Static On Master=%d Hidden=%t\n", cmd.SequenceNumber, myFixtureNumber, cmd.Master, cmd.Hidden)
			}
			lastColor = setStaticOn(myFixtureNumber, cmd, fixturesConfig, eventsForLaunchpad, guiButtons, dmxController, dmxInterfacePresent)
			continue

		case cmd.RGBStaticFadeUp:
			if debug {
				fmt.Printf("%d:%d Static Fade Up Hidden=%t\n", cmd.SequenceNumber, myFixtureNumber, cmd.Hidden)
			}
			// FadeUpStatic doesn't return a lastColor, instead it sends a message directly to the fixture to set lastColor once it's finished fading up.
			fadeUpStatic(myFixtureNumber, cmd, lastColor, stopFadeDown, stopFadeUp, fixturesConfig, fixtureStepChannel, eventsForLaunchpad, guiButtons, dmxController, dmxInterfacePresent)
			continue

		case cmd.RGBStaticOff:
			if debug {
				fmt.Printf("%d:%d Static Off Hidden=%t\n", cmd.SequenceNumber, myFixtureNumber, cmd.Hidden)
			}
			// staticOff doesn't return a lastColor, instead it sends a message directly to the fixture to set lastColor once it's finished fading down.
			staticOff(myFixtureNumber, cmd, lastColor, stopFadeDown, stopFadeUp, fixturesConfig, fixtureStepChannel, eventsForLaunchpad, guiButtons, dmxController, dmxInterfacePresent)
			continue

		case cmd.Type == "scanner":
			if debug {
				fmt.Printf("%d:%d Play Scanner Hidden=%t\n", cmd.SequenceNumber, myFixtureNumber, cmd.Hidden)
			}
			lastColor = playScanner(myFixtureNumber, cmd, fixturesConfig, eventsForLaunchpad, guiButtons, dmxController, dmxInterfacePresent)
			continue

		case cmd.Type == "rgb":
			if debug {
				fmt.Printf("%d:%d Play RGB Hidden=%t\n", cmd.SequenceNumber, myFixtureNumber, cmd.Hidden)
			}
			lastColor = playRGB(myFixtureNumber, cmd, fixturesConfig, eventsForLaunchpad, guiButtons, dmxController, dmxInterfacePresent)
			continue
		}
	}
}
