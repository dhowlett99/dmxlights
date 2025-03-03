// Copyright (C) 2022, 2023 , 2024 dhowlett99.
// This is the override tools package.
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

package override

import (
	"fmt"

	"github.com/dhowlett99/dmxlights/pkg/commands"
	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
)

const debug = false

func ResetSwitchOveride(useFixture *fixture.Fixture, switchNumber int, stateNumber int, switchOverrides *[][]common.Override, fixturesConfig *fixture.Fixtures) {

	// Convert this switches action into a config we can query.
	action := fixture.GetSwitchAction(switchNumber, int16(stateNumber), fixturesConfig)
	cfg := fixture.GetConfig(action, useFixture, fixturesConfig)

	if debug {
		fmt.Printf("ResetSwitchOveride: Reset Override for Fixture Name %s Switch %d State %d\n", useFixture.Name, switchNumber, stateNumber)
	}

	overrides := *switchOverrides

	// Create a new override for this action.
	newOverride := overrides[switchNumber-1][stateNumber-1]

	newOverride.Speed = cfg.Speed
	newOverride.AvailableSpeedChannels = fixture.GetAvailableSpeedChannelsByFixure(useFixture)
	newOverride.MaxSpeeds = len(newOverride.AvailableSpeedChannels)

	newOverride.Shift = cfg.Shift
	newOverride.Size = cfg.Size
	newOverride.Fade = cfg.Fade

	newOverride.Rotate = cfg.RotateSpeed
	newOverride.RotateName = fixture.GetRotateSpeedNameByNumber(useFixture, cfg.RotateSpeed)
	newOverride.RotateChannels = fixture.GetAvailableRotateChannelsByFixure(useFixture)
	newOverride.MaxRotateSpeed = len(newOverride.RotateChannels)

	newOverride.Color = cfg.Color
	newOverride.Colors = cfg.Colors

	newOverride.ColorName = fixture.GetColorNameByNumber(useFixture, newOverride.Color)
	newOverride.AvailableColors = fixture.GetAvailableColorsByFixure(useFixture)
	newOverride.MaxColors = len(newOverride.AvailableColors)

	newOverride.Gobo = cfg.Gobo
	newOverride.AvailableGobos = fixture.GetAvailableGobosByFixure(useFixture)
	newOverride.GoboName = fixture.GetGoboNameByNumber(useFixture, cfg.Gobo)
	newOverride.MaxGobos = len(newOverride.AvailableGobos)

	overrides[switchNumber-1][stateNumber-1] = newOverride

	if debug {

		fmt.Printf("Speed %d\n", newOverride.Speed)
		fmt.Printf("AvailableSpeedChannels %s\n", newOverride.AvailableSpeedChannels)
		fmt.Printf("MaxSpeeds %d\n", newOverride.MaxSpeeds)

		fmt.Printf("Shift %d\n", newOverride.Shift)
		fmt.Printf("Size %d\n", newOverride.Size)
		fmt.Printf("Fade %d\n", newOverride.Fade)

		fmt.Printf("RotateSpeed %d\n", newOverride.Rotate)
		fmt.Printf("RotateName %s\n", newOverride.RotateName)
		fmt.Printf("RotateChannels %s\n", newOverride.RotateChannels)
		fmt.Printf("MaxRotateSpeed %d\n", newOverride.MaxRotateSpeed)

		fmt.Printf("Color %d\n", newOverride.Color)
		fmt.Printf("Colors %+v\n", newOverride.Colors)
		fmt.Printf("ColorName %s\n", newOverride.ColorName)
		fmt.Printf("AvailableColors %s\n", newOverride.AvailableColors)
		fmt.Printf("MaxColors %d\n", newOverride.MaxColors)

		fmt.Printf("Gobo %d\n", newOverride.Gobo)
		fmt.Printf("AvailableGobos %s\n", newOverride.AvailableGobos)
		fmt.Printf("GoboName %s\n", newOverride.GoboName)
		fmt.Printf("MaxGobos %d\n", newOverride.MaxGobos)

	}

	overrides[switchNumber-1][stateNumber-1] = newOverride

	switchOverrides = &overrides
}

func DiscoverSwitchOveride(useFixture *fixture.Fixture, switchNumber int, stateNumber int, fixturesConfig *fixture.Fixtures) common.Override {

	// Convert this switches action into a config we can query.
	action := fixture.GetSwitchAction(switchNumber, int16(stateNumber), fixturesConfig)
	cfg := fixture.GetConfig(action, useFixture, fixturesConfig)

	if debug {
		fmt.Printf("DiscoverSwitchOveride: Discover Fixture Name %s Switch %d State %d\n", useFixture.Name, switchNumber, stateNumber)
	}

	// Create a new override for this action.
	newOverride := common.Override{}

	newOverride.Speed = cfg.Speed
	newOverride.AvailableSpeedChannels = fixture.GetAvailableSpeedChannelsByFixure(useFixture)
	newOverride.MaxSpeeds = len(newOverride.AvailableSpeedChannels)

	newOverride.Shift = cfg.Shift
	newOverride.Size = cfg.Size
	newOverride.Fade = cfg.Fade

	newOverride.Rotate = cfg.RotateSpeed
	newOverride.RotateName = fixture.GetRotateSpeedNameByNumber(useFixture, cfg.RotateSpeed)
	newOverride.RotateChannels = fixture.GetAvailableRotateChannelsByFixure(useFixture)
	newOverride.MaxRotateSpeed = len(newOverride.RotateChannels)

	newOverride.Color = cfg.Color
	newOverride.Colors = cfg.Colors

	newOverride.ColorName = fixture.GetColorNameByNumber(useFixture, newOverride.Color)
	newOverride.AvailableColors = fixture.GetAvailableColorsByFixure(useFixture)
	newOverride.MaxColors = len(newOverride.AvailableColors)

	newOverride.Gobo = cfg.Gobo
	newOverride.AvailableGobos = fixture.GetAvailableGobosByFixure(useFixture)
	newOverride.GoboName = fixture.GetGoboNameByNumber(useFixture, cfg.Gobo)
	newOverride.MaxGobos = len(newOverride.AvailableGobos)

	if debug {
		fmt.Printf("Action Mode %s\n", action.Mode)
		fmt.Printf("Action Name %s\n", action.Name)
		fmt.Printf("Switch Number %d State Number %d\n", switchNumber, stateNumber)
		fmt.Printf("Rotate Speed %d\n", newOverride.Rotate)
		fmt.Printf("Max Rotate Speeds %d\n", newOverride.MaxRotateSpeed)
		fmt.Printf("Action Color %s Color %d\n", action.Colors, newOverride.Color)
		fmt.Printf("Colors %+v\n", newOverride.Colors)
		fmt.Printf("AvailableColors %s\n", newOverride.AvailableColors)
		fmt.Printf("MaxColors %+v\n", newOverride.MaxColors)
		fmt.Printf("Color Names %s\n", newOverride.ColorName)
		fmt.Printf("Gobo action %s newOverride Gobo %d Gobo Name %s\n", action.Gobo, newOverride.Gobo, newOverride.GoboName)
		fmt.Printf("==========================================\n")
	}
	return newOverride
}

func CreateOverrides(sequenceNumber int, fixturesConfig *fixture.Fixtures, switchOverrides *[][]common.Override) {

	if debug {
		fmt.Printf("UpdateOverrides\n")
	}

	// Store the switch Config locally.
	switchConfig := commands.LoadSwitchConfiguration(sequenceNumber, fixturesConfig)

	// Copy in the overrides.
	overrides := *switchOverrides

	// Populate each switch with a number of states based on their config.
	for swiTchNumber := 0; swiTchNumber < len(switchConfig); swiTchNumber++ {

		// assign the switch.
		swiTch := switchConfig[swiTchNumber]

		// Now populate the states.
		for stateNumber := 0; stateNumber < len(swiTch.States); stateNumber++ {

			state := swiTch.States[stateNumber]

			// Find the details of the fixture for this switch.
			thisFixture, err := fixture.GetFixtureByLabel(swiTch.UseFixture, fixturesConfig)
			if err != nil {
				fmt.Printf("error %s\n", err.Error())
			}

			// Load the config for this state of of this switch
			override := DiscoverSwitchOveride(thisFixture, swiTch.Number, int(state.Number), fixturesConfig)

			// Assign this discovered override to the current switch state.
			overrides[swiTchNumber] = append(overrides[swiTchNumber], override)

			if debug {
				fmt.Printf("Setting Up Override for Switch No=%d Name=%s State No=%d Name=%s\n", swiTch.Number, swiTch.Name, state.Number, state.Name)
				fmt.Printf("\t Override Colors %+v\n", override.Colors)
			}
		}
	}

	if debug {
		for overrideNumber, override := range overrides {
			fmt.Printf("overrideNumber %d\n", overrideNumber)
			for stateNumber, state := range override {
				fmt.Printf("\tstateNumber %d state %+v\n", stateNumber, state)
			}
		}
	}

	// re-instate the pointer to the overrides.
	switchOverrides = &overrides
}

func ResetOverrides(sequenceNumber int, fixturesConfig *fixture.Fixtures, switchOverrides *[][]common.Override) {

	if debug {
		fmt.Printf("ResetOverrides\n")
	}

	// Store the switch Config locally.
	switchConfig := commands.LoadSwitchConfiguration(sequenceNumber, fixturesConfig)

	// Copy in the overrides.
	overrides := *switchOverrides

	// Populate each switch with a number of states based on their config.
	for swiTchNumber := 0; swiTchNumber < len(switchConfig); swiTchNumber++ {

		// assign the switch.
		swiTch := switchConfig[swiTchNumber]

		// Now populate the states.
		for stateNumber := 0; stateNumber < len(swiTch.States); stateNumber++ {

			state := swiTch.States[stateNumber]

			// Find the details of the fixture for this switch.
			thisFixture, err := fixture.GetFixtureByLabel(swiTch.UseFixture, fixturesConfig)
			if err != nil {
				fmt.Printf("error %s\n", err.Error())
			}

			// Load the config for this state of of this switch
			ResetSwitchOveride(thisFixture, swiTch.Number, int(state.Number), switchOverrides, fixturesConfig)

			if debug {
				fmt.Printf("Setting Up Override for Switch No=%d Name=%s State No=%d Name=%s\n", swiTch.Number, swiTch.Name, state.Number, state.Name)
			}
		}
	}

	if debug {
		for overrideNumber, override := range overrides {
			fmt.Printf("overrideNumber %d\n", overrideNumber)
			for stateNumber, state := range override {
				fmt.Printf("\tstateNumber %d state %+v\n", stateNumber, state)
			}
		}
	}

	// re-instate the pointer to the overrides.
	switchOverrides = &overrides
}
