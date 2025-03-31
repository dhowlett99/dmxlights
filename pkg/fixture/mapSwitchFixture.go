// Copyright (C) 2022,2023,2024,2025 dhowlett99.
// This is the dmxlights switch mapping code, it will ultimately end up
// calling the mini sequencer or the mini setter functions.
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

	"github.com/dhowlett99/dmxlights/pkg/colors"
	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/sound"
	"github.com/oliread/usbdmx/ft232"
)

// MapSwitchFixture is repsonsible for playing out the state of a swicth.
// The switch is idendifed by the sequence and switch number.
func MapSwitchFixture(swiTch common.Switch,
	state common.State,
	override common.Override,
	RGBFade int,
	dmxController *ft232.DMXController,
	fixturesConfig *Fixtures, blackout bool,
	brightness int, master int, masterChanging bool, lastColor common.LastColor,
	switchChannels []common.SwitchChannel,
	SoundTriggers []*common.Trigger,
	soundConfig *sound.SoundConfig,
	dmxInterfacePresent bool,
	eventsForLaunchpad chan common.ALight,
	guiButtons chan common.ALight,
	fixtureStepChannel chan common.FixtureCommand) common.LastColor {

	var useFixtureLabel string

	if debug {
		fmt.Printf("MapSwitchFixture switchNumber %d, current position %d fade speed %d\n", swiTch.Number, swiTch.CurrentPosition, RGBFade)
	}

	// Master dinmmer is the best way to blackout a switch.
	if blackout {
		master = 0
	}

	// We start by having the switch and its current state passed in.

	// Now we find the fixture used by the switch
	if swiTch.UseFixture != "" {
		// use this fixture for the sequencer actions
		// BTW UseFixtureLabel is the label for the fixture NOT the name.
		useFixtureLabel = swiTch.UseFixture

		if debug {
			fmt.Printf("useFixtureLabel %s  blackout is %t\n", useFixtureLabel, blackout)
		}

		// Find the details of the fixture for this switch.
		thisFixture, err := GetFixtureDetailsByLabel(useFixtureLabel, fixturesConfig)
		if err != nil {
			fmt.Printf("error %s\n", err.Error())
			return lastColor
		}

		if debug {
			fmt.Printf("Found fixture Name %s \n", thisFixture.Name)
		}

		// Look for Master channel in this fixture identified by ID.
		masterChannel, err := GetChannelNumberByName(thisFixture, "Master")
		if err != nil && debug {
			fmt.Printf("warning! fixture %s: %s\n", thisFixture.Name, err)
		}

		// Play Actions which send messages to a dedicated mini sequencer.
		for _, action := range state.Actions {
			newAction := Action{}
			newAction.Name = action.Name
			newAction.Number = action.Number
			newAction.Colors = action.Colors
			newAction.Mode = action.Mode
			newAction.Fade = action.Fade
			newAction.Size = action.Size
			newAction.Speed = action.Speed
			newAction.Rotate = action.Rotate
			newAction.RotateSpeed = action.RotateSpeed
			newAction.Program = action.Program
			newAction.ProgramSpeed = action.ProgramSpeed
			newAction.Strobe = action.Strobe
			newAction.Map = action.Map
			newAction.Gobo = action.Gobo
			newAction.GoboSpeed = action.GoboSpeed
			newMiniSequencer(thisFixture, swiTch, override, newAction, dmxController, fixturesConfig, switchChannels, soundConfig, blackout, brightness, master, masterChanging, lastColor, dmxInterfacePresent, eventsForLaunchpad, guiButtons, fixtureStepChannel)
			if action.Mode != "Static" {
				lastColor.RGBColor = colors.EmptyColor
			}
		}

		// If there are no actions, turn off any previos mini sequencers for this switch.
		if len(state.Actions) == 0 {
			newAction := Action{}
			newAction.Name = "Off"
			newAction.Number = 1
			newAction.Mode = "Off"
			lastColor := common.LastColor{}
			newMiniSequencer(thisFixture, swiTch, override, newAction, dmxController, fixturesConfig, switchChannels, soundConfig, blackout, brightness, master, masterChanging, lastColor, dmxInterfacePresent, eventsForLaunchpad, guiButtons, fixtureStepChannel)
		}

		// Now play any preset DMX values directly to the universe.
		// Step through all the settings.
		for _, newSetting := range state.Settings {
			newMiniSetter(thisFixture, &override, newSetting, masterChannel, dmxController, master, dmxInterfacePresent)
		}
	}
	return lastColor
}
