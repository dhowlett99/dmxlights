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

func ResetSwitchOveride(useFixture *fixture.Fixture, switchNumber int, stateNumber int, switchOverrides *[][]common.Override, fixturesConfig *fixture.Fixtures) common.Override {

	// Convert this switches action into a config we can query.
	action := fixture.GetSwitchAction(switchNumber, int16(stateNumber), fixturesConfig)
	cfg := fixture.GetConfig(action, useFixture, fixturesConfig)

	if debug {
		fmt.Printf("ResetSwitchOveride: Reset Override for Fixture Name %s Switch %d State %d\n", useFixture.Name, switchNumber, stateNumber)
	}

	overrides := *switchOverrides

	// Create a new override for this action.
	newOverride := overrides[switchNumber-1][stateNumber-1]

	populateOverride(useFixture, &newOverride, cfg)

	overrides[switchNumber-1][stateNumber-1] = newOverride

	return newOverride
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

	populateOverride(useFixture, &newOverride, cfg)

	return newOverride
}

func CreateOverrides(sequenceNumber int, fixturesConfig *fixture.Fixtures, switchOverrides *[][]common.Override) {

	if debug {
		fmt.Printf("CreateOverrides\n")
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
			thisFixture, err := fixture.GetFixtureDetailsByLabel(swiTch.UseFixture, fixturesConfig)
			if err != nil {
				fmt.Printf("error %s\n", err.Error())
			}

			// Load the config for this state of of this switch
			override := DiscoverSwitchOveride(thisFixture, swiTchNumber, int(stateNumber), fixturesConfig)

			// Assign this discovered override to the current switch state.
			overrides[swiTchNumber] = append(overrides[swiTchNumber], override)

			if debug {
				fmt.Printf("Setting Up Override for Switch No=%d Name=%s State No=%d Name=%s\n", swiTch.Number, swiTch.Name, state.Number, state.Name)
				fmt.Printf("\t Override Colors %+v\n", override.AvailableColors)
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
			thisFixture, err := fixture.GetFixtureDetailsByLabel(swiTch.UseFixture, fixturesConfig)
			if err != nil {
				fmt.Printf("error %s\n", err.Error())
			}

			// Load the config for this state of of this switch
			overrides[swiTchNumber][stateNumber] = ResetSwitchOveride(thisFixture, swiTch.Number, int(state.Number), switchOverrides, fixturesConfig)
			switchOverrides = &overrides

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

func populateOverride(useFixture *fixture.Fixture, newOverride *common.Override, cfg fixture.ActionConfig) {

	newOverride.Mode = cfg.Mode // Action mode. Setting, Off , Static, Control, Chase.

	newOverride.IsShutterOverrideAble = fixture.IsThisChannelOverrideAble(useFixture, "Strobe")
	newOverride.Strobe = cfg.Strobe
	newOverride.StrobeSpeed = cfg.StrobeSpeed
	newOverride.AvailableStrobeSpeedChannels = fixture.GetAvailableSettingsForChannelsByFixure(useFixture, "Speed")
	newOverride.MaxStrobeSpeeds = len(newOverride.AvailableStrobeSpeedChannels)

	newOverride.IsSpeedOverrideAble = fixture.IsThisChannelOverrideAble(useFixture, "Speed")
	newOverride.Speed = cfg.Speed
	newOverride.MaxSpeeds = common.MAX_SPEED

	newOverride.Shift = cfg.Shift

	newOverride.Size = cfg.Size

	newOverride.Fade = cfg.Fade

	newOverride.IsProgramSpeedOverrideAble = fixture.IsThisChannelOverrideAble(useFixture, "ProgramSpeed")
	newOverride.ProgramSpeed = cfg.ProgramSpeed
	newOverride.AvailableProgramSpeedChannels = fixture.GetAvailableSettingsForChannelsByFixure(useFixture, "ProgramSpeed")
	newOverride.MaxProgramSpeeds = len(newOverride.AvailableProgramSpeedChannels)

	newOverride.IsRotateOverrideAble = fixture.IsThisChannelOverrideAble(useFixture, "Rotate")
	newOverride.Rotate = cfg.Rotate
	newOverride.RotateName = fixture.GetRotateSpeedNameByNumber(useFixture, cfg.RotateSpeed)
	newOverride.RotateChannels = fixture.GetAvailableRotateChannelsByFixure(useFixture)
	newOverride.MaxRotateSpeed = len(newOverride.RotateChannels)

	newOverride.HasRGBChannels = useFixture.HasRGBChannels
	newOverride.HasColorChannel = useFixture.HasColorChannel
	newOverride.IsColorOverrideAble = fixture.IsThisChannelOverrideAble(useFixture, "Color")
	newOverride.Color = cfg.Color
	newOverride.ColorName = cfg.ColorName
	newOverride.AvailableColorNames = cfg.AvailableColorNames
	newOverride.MaxColors = len(cfg.AvailableColorNames) - 1

	newOverride.IsGoboOverrideAble = fixture.IsThisChannelOverrideAble(useFixture, "Gobo")
	newOverride.Gobo = cfg.Gobo
	newOverride.AvailableGobos = fixture.GetAvailableSettingsForChannelsByFixure(useFixture, "Gobo")
	newOverride.GoboName = fixture.GetGoboNameByNumber(useFixture, cfg.Gobo)
	newOverride.MaxGobos = len(newOverride.AvailableGobos)

	if debug {
		fmt.Printf("Populate Fixture %s\n", useFixture.Name)
		fmt.Printf("Speed OverrideAble %t\n", newOverride.IsSpeedOverrideAble)
		fmt.Printf("Speed %d\n", newOverride.Speed)
		fmt.Printf("MaxSpeeds %d\n", newOverride.MaxSpeeds)

		fmt.Printf("Shift %d\n", newOverride.Shift)
		fmt.Printf("Size %d\n", newOverride.Size)
		fmt.Printf("Fade %d\n", newOverride.Fade)

		fmt.Printf("ProgramSpeed OverrideAble %t\n", newOverride.IsProgramSpeedOverrideAble)
		fmt.Printf("ProgramSpeed %d\n", newOverride.ProgramSpeed)
		fmt.Printf("AvailableProgramSpeedChannels %s\n", newOverride.AvailableProgramSpeedChannels)
		fmt.Printf("MaxProgramSpeeds %d\n", newOverride.MaxProgramSpeeds)

		fmt.Printf("Rotate OverrideAble %t\n", newOverride.IsRotateOverrideAble)
		fmt.Printf("RotateSpeed %d\n", newOverride.Rotate)
		fmt.Printf("RotateName %s\n", newOverride.RotateName)
		fmt.Printf("RotateChannels %s\n", newOverride.RotateChannels)
		fmt.Printf("MaxRotateSpeed %d\n", newOverride.MaxRotateSpeed)

		fmt.Printf("Color OverrideAble %t\n", newOverride.IsColorOverrideAble)
		fmt.Printf("Color %d\n", newOverride.Color)
		fmt.Printf("ColorName %s\n", newOverride.ColorName)
		fmt.Printf("AvailableColors %+v\n", newOverride.AvailableColors)
		fmt.Printf("AvailableColorNames %+v\n", newOverride.AvailableColorNames)
		fmt.Printf("MaxColors %d\n", newOverride.MaxColors)

		fmt.Printf("Gobo OverrideAble %t\n", newOverride.IsGoboOverrideAble)
		fmt.Printf("Gobo %d\n", newOverride.Gobo)
		fmt.Printf("AvailableGobos %s\n", newOverride.AvailableGobos)
		fmt.Printf("GoboName %s\n", newOverride.GoboName)
		fmt.Printf("MaxGobos %d\n", newOverride.MaxGobos)
		fmt.Printf("\n")
	}
}
