// Copyright (C) 2022,2023,2024,2025 dhowlett99.
// This is the dmxlights fixture library that retrieves info from the fixture.
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
	"image/color"
	"strconv"
	"strings"

	"github.com/dhowlett99/dmxlights/pkg/colors"
	"github.com/dhowlett99/dmxlights/pkg/common"
)

// GetFixtureDetailsById - find a fixture in the fixtures config.
// Returns details of the fixture.
// Returns an error.
func GetFixtureDetailsById(id int, fixtures *Fixtures) (Fixture, error) {

	if debug {
		fmt.Printf("GetFixtureDetailsById\n")
	}

	// scan the fixtures structure for the selected fixture.
	if debug {
		fmt.Printf("Looking for Fixture ID %d\n", id)
	}

	for _, fixture := range fixtures.Fixtures {
		if debug {
			fmt.Printf("Fixture ID %d and Name %s States %+v\n", fixture.ID, fixture.Name, fixture.States)
		}
		if fixture.ID == id {
			return fixture, nil
		}
	}
	return Fixture{}, fmt.Errorf("error: fixture id %d not found", id)
}

// GetFixtureDetailsByLabel - find a fixture in the fixtures config.
// Returns details of the fixture.
// Returns an error.
func GetFixtureDetailsByLabel(label string, fixtures *Fixtures) (*Fixture, error) {
	// scan the fixtures structure for the selected fixture.
	if debug {
		fmt.Printf("GetFixtureDetailsByLabel: Looking for Fixture by Label %s\n", label)
	}

	for _, fixture := range fixtures.Fixtures {
		if debug {
			fmt.Printf("Fixture Label %s and Name %s States %+v\n", fixture.Label, fixture.Name, fixture.States)
		}
		if fixture.Label == label {
			return &fixture, nil
		}
	}
	return &Fixture{}, fmt.Errorf("error: fixture label %s not found", label)
}

// GetChannelSettingByNameAndSetting Look through the fixtures channels and use the channel name and setting name.
// returns the setting value.
func GetChannelSettingByNameAndSetting(fixture *Fixture, channelName string, settingName string) (int, error) {

	if debug {
		fmt.Printf("GetChannelSettingByNameAndSetting for fixture %s on channel %s setting %s\n", fixture.Name, channelName, settingName)
	}

	if settingName == "" {
		return 0, fmt.Errorf("GetChannelSettingByNameAndSetting: settingName is empty for channel %s in fixture %s", channelName, fixture.Name)
	}

	for _, channel := range fixture.Channels {
		if debug {
			fmt.Printf("inspect channel %s for %s\n", channel.Name, settingName)
		}
		if channel.Name == channelName {
			if debug {
				fmt.Printf("channel.Settings %+v\n", channel.Settings)
			}
			for _, setting := range channel.Settings {
				if debug {
					fmt.Printf("inspect setting %+v \n", setting)
					fmt.Printf("setting.Name %s = name %s\n", setting.Name, settingName)
				}
				if setting.Name == settingName {
					if debug {
						fmt.Printf("FixtureName=%s ChannelName=%s SettingName=%s SettingValue=%s\n", fixture.Name, channel.Name, settingName, setting.Value)
					}

					var v int
					var err error
					// If the setting value contains a "-" remove it and then take the first valuel.
					if strings.Contains(setting.Value, "-") {
						// We've found a range of values.
						// Find the start value
						numbers := strings.Split(setting.Value, "-")
						v, err = strconv.Atoi(numbers[0])
						if err != nil {
							return 0, err
						}
					} else {
						v, err = strconv.Atoi(setting.Value)
						if err != nil {
							return 0, err
						}
					}
					if debug {
						fmt.Printf("Value Returned %d\n", v)
					}

					return v, nil
				}
			}
		}
	}

	return 0, fmt.Errorf("GetChannelSettingByNameAndSetting: setting %s not found in channel %s for fixture %s", settingName, channelName, fixture.Name)
}

func GetChannelSettingByNameAndSpeed(fixture *Fixture, channelName string, settingName string, settingSpeed string, fixtures *Fixtures) (int, error) {
	if debug {
		fmt.Printf("GetChannelSettingByNameAndSpeed for fixture %s channel name %s setting name %s and setting speed %s\n", fixture.Name, channelName, settingName, settingSpeed)
	}

	if channelName == "" {
		return 0, fmt.Errorf("GetChannelSettingByNameAndSpeed: Fixture %s error channel name is empty", fixture.Name)
	}
	if settingName == "" {
		return 0, fmt.Errorf("GetChannelSettingByNameAndSpeed: Fixture %s error setting name is empty", fixture.Name)
	}
	if settingSpeed == "" {
		return 0, fmt.Errorf("GetChannelSettingByNameAndSpeed: Fixture %s error setting speed is empty", fixture.Name)
	}

	for _, channel := range fixture.Channels {

		if channel.Name == channelName {
			if debug {
				fmt.Printf("fixture %s: inspect channel %s for %s\n", fixture.Name, channel.Name, settingName)
			}

			for _, setting := range channel.Settings {
				if debug {
					fmt.Printf("\tinspect setting %+v \n", setting)
					fmt.Printf("\tgot:setting.Name %s  want name %s speed %s\n", setting.Name, settingName, settingSpeed)
				}
				if strings.Contains(setting.Name, settingName) && strings.Contains(setting.Name, settingSpeed) {

					if debug {
						fmt.Printf("\t\tFixtureName=%s ChannelName=%s SettingName=%s SettingSpeed=%s, SettingValue=%s\n", fixture.Name, channel.Name, settingName, settingSpeed, setting.Value)
					}

					// If the setting value contains a "-" remove it and then take the first valuel.
					var err error
					var v int
					if strings.Contains(setting.Value, "-") {
						// We've found a range of values.
						// Find the start value
						numbers := strings.Split(setting.Value, "-")
						v, err = strconv.Atoi(numbers[0])
						if err != nil {
							return 0, err
						}
					} else {
						v, err = strconv.Atoi(setting.Value)
						if err != nil {
							return 0, err
						}
					}

					if debug {
						fmt.Printf("FOUND -> channel %s setting %s %s in fixture %s value %d", channelName, settingName, settingSpeed, fixture.Name, v)
					}
					return v, nil
				}
			}
		}
	}
	return 0, fmt.Errorf("warning: channel %s setting %s %s not found in fixture %s", channelName, settingName, settingSpeed, fixture.Name)
}

func GetFixtureByGroupAndNumber(sequenceNumber int, fixtureNumber int, fixtures *Fixtures) (*Fixture, error) {

	if debug {
		fmt.Printf("FindFixtureByGroupAndNumber seq %d fixture %d\n", sequenceNumber, fixtureNumber)
	}
	for _, fixture := range fixtures.Fixtures {
		if fixture.Group == sequenceNumber+1 {
			if fixture.Number == fixtureNumber+1 {
				if debug {
					fmt.Printf("Found fixture %s Group %d Number %d Address %d\n", fixture.Name, fixture.Group, fixture.Number, fixture.Address)
				}
				return &fixture, nil
			}
		}
	}
	return nil, fmt.Errorf("FindFixtureByGroupAndNumber: failed to find fixture for sequence %d fixture %d", sequenceNumber, fixtureNumber)
}

func GetFadeValuesFixtureAddressByName(fixtureName string, fixtures *Fixtures) string {
	if debug {
		fmt.Printf("FindFixtureAddressByName: Looking for fixture by Name %s\n", fixtureName)
	}
	for _, fixture := range fixtures.Fixtures {
		if fixture.Label == fixtureName {
			if debug {
				fmt.Printf("Found fixture %s Group %d Number %d Address %d\n", fixture.Name, fixture.Group, fixture.Number, fixture.Address)
			}
			return fmt.Sprintf("%d", fixture.Address)
		}
	}
	if debug {
		fmt.Printf("fixture %s not found\n", fixtureName)
	}
	return "Not Set"
}

// GetShutter takes the name of a gobo channel setting like "Open" and returns the gobo number  for this type of scanner.
func GetShutter(myFixtureNumber int, mySequenceNumber int, shutterName string, fixtures *Fixtures) int {

	if debug {
		fmt.Printf("GetShutter\n")
	}

	for _, fixture := range fixtures.Fixtures {
		if fixture.Group == mySequenceNumber+1 {
			if fixture.Number == myFixtureNumber+1 {
				for _, channel := range fixture.Channels {
					if strings.Contains(channel.Name, "Shutter") {
						for _, setting := range channel.Settings {
							if strings.Contains(setting.Name, shutterName) {
								return setting.Number
							}
						}
					}
				}
			}
		}
	}
	return 255
}

// GetGoboNameByNumber takes the gobo number and returns the gobo name for this fixture.
func GetGoboNameByNumber(fixture *Fixture, number int) string {

	if debug {
		fmt.Printf("GetGoboNameByNumber Looking for gobo %d in fixture %s\n", number, fixture.Name)
	}

	if number == -1 {
		if debug {
			fmt.Printf("Gobo %d Name Auto\n", number)
		}
		return "Auto"
	}

	if fixture == nil {
		return "Not Found"
	}

	for _, channel := range fixture.Channels {
		if strings.Contains(channel.Name, "Gobo") {
			for _, setting := range channel.Settings {
				if setting.Number == number {
					if debug {
						fmt.Printf("Gobo %d Name %s\n", setting.Number, setting.Name)
					}
					return setting.Name
				}
			}
		}
	}

	return "Unknown"
}

// GetRotateSpeedNameByNumber takes the rotate speed number and returns the rotate speed name for this fixture.
func GetRotateSpeedNameByNumber(fixture *Fixture, number int) string {

	if debug {
		fmt.Printf("GetRotateSpeedNameByNumber Looking for rotate speed %d in fixture %s\n", number, fixture.Name)
	}

	if number == -1 {
		if debug {
			fmt.Printf("Rotate Speed %d Name Auto\n", number)
		}
		return "Auto"
	}

	if fixture == nil {
		return "Not Found"
	}

	for _, channel := range fixture.Channels {
		if strings.Contains(channel.Name, "Rotate") {
			for _, setting := range channel.Settings {
				if setting.Number == number {
					if debug {
						fmt.Printf("Rotate Speed %d Name %s\n", setting.Number, setting.Name)
					}
					return setting.Name
				}
			}
		}
	}

	return "Unknown"
}

// GetRotateSpeedNumberByName takes the rotate speed name and returns the rotate speed setting number for this fixture.
func GetRotateSpeedNumberByName(fixture *Fixture, rotateSettingName string) int {

	if debug {
		fmt.Printf("GetRotateSpeedNumberByName Looking for rotate speed %s in fixture %s\n", rotateSettingName, fixture.Name)
	}

	if fixture == nil {
		fmt.Printf("GetRotateSpeedNumberByName: fixture is empty\n")
		return 0
	}

	for _, channel := range fixture.Channels {
		if strings.Contains(channel.Name, "Rotate") {
			for _, setting := range channel.Settings {
				if setting.Name == rotateSettingName {
					if debug {
						fmt.Printf("Rotate Speed %d Name %s\n", setting.Number, setting.Name)
					}
					return setting.Number
				}
			}
		}
	}

	return 0
}

// GetColorDMXValueByNumber takes the gobo number and returns the DMX value which selects this Gobo.
func GetColorDMXValueByNumber(fixture *Fixture, number int) int {

	if debug {
		fmt.Printf("GetColorDMXValueByNumber Looking for color %d in fixture %s\n", number, fixture.Name)
	}

	for _, channel := range fixture.Channels {
		if strings.Contains(channel.Name, "Color") {
			for _, setting := range channel.Settings {
				if setting.Number == number {
					if debug {
						fmt.Printf("Color %d Name %s\n", setting.Number, setting.Name)
					}
					dmx, _ := strconv.Atoi(setting.Value)
					return dmx
				}
			}
		}
	}

	return 0
}

// GetADMXValue takes the setting number and channel name then returns the DMX value.
func GetADMXValue(fixture *Fixture, settinNumber int, channelName string) int {

	if debug {
		fmt.Printf("GetADMXValue Looking in channel %s for number %d in fixture %s\n", channelName, settinNumber, fixture.Name)
	}

	for _, channel := range fixture.Channels {
		if strings.Contains(channel.Name, channelName) {
			for _, setting := range channel.Settings {
				if setting.Number == settinNumber {
					dmx, _ := strconv.Atoi(setting.Value)
					if debug {
						fmt.Printf("Setting Name %d DMX value %d\n", setting.Number, dmx)
					}
					return dmx
				}
			}
		}
	}

	return 0
}

// GetADMXValueByName takes the setting number and channel name then returns the DMX value.
func GetADMXValueByName(fixture *Fixture, settingName string, channelName string) int {

	if debug {
		fmt.Printf("GetADMXValueByName Looking in channel %s for name %s in fixture %s\n", channelName, settingName, fixture.Name)
	}

	for _, channel := range fixture.Channels {
		if strings.Contains(channel.Name, channelName) {
			for _, setting := range channel.Settings {
				if setting.Name == settingName {
					dmx, _ := strconv.Atoi(setting.Value)
					if debug {
						fmt.Printf("Setting Name %d DMX value %d\n", setting.Number, dmx)
					}
					return dmx
				}
			}
		}
	}

	return 0
}

// GetADMXValueMaxMin takes the setting name and channel name then returns the DMX value.
func GetADMXValueMaxMin(fixture *Fixture, settingName string, channelName string) []int {

	if debug {
		fmt.Printf("GetADMXValueMaxMin Looking in channel %s for name %s in fixture %s\n", channelName, settingName, fixture.Name)
	}

	dmxValues := []int{}

	for _, channel := range fixture.Channels {
		if strings.Contains(channel.Name, channelName) {
			for _, setting := range channel.Settings {
				if setting.Name == settingName {

					if strings.Contains(setting.Value, "-") {
						values := strings.Split(setting.Value, "-")
						for _, value := range values {
							dmx, _ := strconv.Atoi(value)
							if debug {
								fmt.Printf("Found Values in Setting Name %s DMX value %d\n", setting.Name, dmx)
							}
							dmxValues = append(dmxValues, dmx)
						}
						return dmxValues
					} else {
						dmx, _ := strconv.Atoi(setting.Value)
						//if debug {
						fmt.Printf("Found A single Value in Setting Name %s DMX value %d\n", setting.Name, dmx)
						//}
						dmxValues = append(dmxValues, dmx)
						return dmxValues
					}

				}
			}
		}
	}

	return dmxValues
}

func IsThisChannelOverrideAble(fixture *Fixture, channelName string) bool {

	for _, channel := range fixture.Channels {
		if channel.Name == channelName {
			return channel.Override
		}
	}
	return false
}

func GetAvailableSettingsForChannelsByFixure(fixture *Fixture, channelName string) []string {

	var settings []string
	for _, channel := range fixture.Channels {
		if strings.Contains(channel.Name, channelName) {
			for _, setting := range channel.Settings {
				settings = append(settings, setting.Name)
			}
		}
	}
	return settings
}

func GetAvailableRotateChannelsByFixure(fixture *Fixture) []string {

	var rotateSpeeds []string
	for _, channel := range fixture.Channels {
		if strings.Contains(channel.Name, "Rotate") {
			for _, setting := range channel.Settings {
				rotateSpeeds = append(rotateSpeeds, setting.Name)
			}
		}
	}
	return rotateSpeeds
}

// GetColorNameByNumber takes the color number and returns the color name for this fixture.
func GetColorNameByNumber(fixture *Fixture, number int) string {

	if number == 0 {
		number = 1
	}
	if fixture == nil {
		return "Not Found"
	}
	if debug {
		fmt.Printf("GetColorNameByNumber looking for color number %d inside fixture %s\n", number, fixture.Name)
	}

	for _, channel := range fixture.Channels {
		if strings.Contains(channel.Name, "Color") {
			for _, setting := range channel.Settings {
				if setting.Number == number {
					if debug {
						fmt.Printf("Found color name %s\n", setting.Name)
					}
					return setting.Name
				}
			}
		}
	}

	if debug {
		fmt.Printf("NOT Found  color name Unkown\n")
	}
	return "Unknown"
}

// GetChannelNumberByName takes the name of a channel and setting  and returns the setting number.
func GetChannelSettingInfo(fixture *Fixture, channelName string, settingName string, fixtures *Fixtures) int {

	if debug {
		fmt.Printf("GetChannelNumberByName fixture name %s number %d channel name %s\n", fixture.Name, fixture.Number, channelName)
	}

	for _, channel := range fixture.Channels {
		if strings.Contains(channel.Name, channelName) {
			for _, setting := range channel.Settings {
				if setting.Name == settingName {
					return setting.Number
				}
			}
		}
	}

	return -1
}

// GetGoboByName takes the name of a gobo channel setting like "Open" and returns the gobo number  for this type of scanner.
func GetGoboByName(myFixtureNumber int, mySequenceNumber int, selectedGobo string, fixtures *Fixtures) int {

	if debug {
		fmt.Printf("GetGoboByName fixture %d seq %d gobo name %s\n", myFixtureNumber, mySequenceNumber, selectedGobo)
	}

	for _, fixture := range fixtures.Fixtures {
		if fixture.Group == mySequenceNumber+1 {
			if fixture.Number == myFixtureNumber+1 {
				for _, channel := range fixture.Channels {
					if strings.Contains(channel.Name, "Gobo") {
						for _, setting := range channel.Settings {
							if strings.Contains(setting.Name, selectedGobo) {
								if debug {
									fmt.Printf("GetGoboByName setting no %d\n", setting.Number)
								}
								return setting.Number
							}
						}
					}
				}
			}
		}
	}
	return 0
}

// GetNumberOfGobosForThisFixture takes the fixture number, sequence number and the fixturesConfig
// Returns returns the number of gobos this fixture has.
func GetNumberOfGobosForThisFixture(myFixtureNumber int, mySequenceNumber int, fixtures *Fixtures) int {

	if debug {
		fmt.Printf("GetNumberOfGobosForThisFixture\n")
	}

	for _, fixture := range fixtures.Fixtures {
		if fixture.Group == mySequenceNumber+1 {
			if fixture.Number == myFixtureNumber+1 {
				for _, channel := range fixture.Channels {
					if strings.Contains(channel.Name, "Gobo") {
						return len(channel.Settings)
					}
				}
			}
		}
	}
	return 0
}

func GetColorNumberFromFixture(fixture *Fixture, color string) int {

	if debug {
		fmt.Printf("FindColor looking for %s in fixture %s\n", color, fixture.Name)
	}

	for _, channel := range fixture.Channels {
		if strings.Contains(channel.Name, "Color") {
			for _, setting := range channel.Settings {
				if setting.Name == color {
					if debug {
						fmt.Printf("Found setting number %d\n", setting.Number)
					}
					return setting.Number
				}
			}
		}
	}
	if debug {
		fmt.Printf("Not FOund setting number returning 0\n")
	}
	return 0
}

// GetColor takes the name of a color channel setting like "White" and returns the color number for this type of scanner.
func GetColor(myFixtureNumber int, mySequenceNumber int, color string, fixtures *Fixtures) int {

	if debug {
		fmt.Printf("GetColor looking for %s seq %d fixture %d\n", color, mySequenceNumber, myFixtureNumber)
	}

	for _, fixture := range fixtures.Fixtures {
		if fixture.Group == mySequenceNumber+1 {
			if fixture.Number == myFixtureNumber+1 {
				for _, channel := range fixture.Channels {
					if strings.Contains(channel.Name, "Color") {
						for _, setting := range channel.Settings {
							if setting.Name == color {
								if debug {
									fmt.Printf("Found setting number %d\n", setting.Number)
								}
								return setting.Number
							}
						}
					}
				}
			}
		}
	}
	if debug {
		fmt.Printf("Not FOund setting number returning 0\n")
	}
	return 0
}

func GetChannelNumberByName(fixture *Fixture, channelName string) (int, error) {

	if debug {
		fmt.Printf("FindChannelNumberByName channelName %s\n", channelName)
	}

	for channelNumber, channel := range fixture.Channels {
		if strings.Contains(channel.Name, channelName) {
			return channelNumber, nil
		}
	}
	return 0, fmt.Errorf("channel %s not found in fixture %s", channelName, fixture.Name)
}

func GetFixtureInfo(thisFixture *Fixture) FixtureInfo {
	if debug {
		fmt.Printf("FindFixtureInfo\n")
	}

	fixtureInfo := FixtureInfo{}

	if thisFixture == nil {
		fmt.Printf("FindFixtureInfo: fixture is empty\n")
		return fixtureInfo
	}

	fixtureInfo.HasRotate = IsThisAChannel(thisFixture, "Rotate")
	fixtureInfo.HasRotateSpeed = IsThisAChannel(thisFixture, "RotateSpeed")

	// Find all the options for the channel called "Rotate".But only if we have a Rotate Channel exists.
	if fixtureInfo.HasRotate {
		availableRotateOptions := GetChannelOptions(*thisFixture, "Rotate")
		// Add the auto option for rotate
		var autoFound bool
		for _, option := range availableRotateOptions {
			if strings.Contains(option, "Auto") || strings.Contains(option, "auto") {
				autoFound = true
			}
			fixtureInfo.RotateOptions = append(fixtureInfo.RotateOptions, option)
		}
		// Now if we didn't find a dedicated channel for automatically rotating in different directions.
		// Add our internal keyword for Auto.
		if !autoFound {
			fixtureInfo.RotateOptions = append(fixtureInfo.RotateOptions, "Auto")
		}
	}

	fixtureInfo.RotateSpeedOptions = []string{"Slow", "Medium", "Fast"}

	fixtureInfo.HasColorWheel = IsThisAChannel(thisFixture, "Color")
	fixtureInfo.HasGobo = IsThisAChannel(thisFixture, "Gobo")
	fixtureInfo.HasProgram = IsThisAChannel(thisFixture, "Program")
	fixtureInfo.HasProgramSpeed = IsThisAChannel(thisFixture, "ProgramSpeed")
	return fixtureInfo
}

func GetChannelOptions(thisFixture Fixture, channelName string) []string {

	if debug {
		fmt.Printf("GetChannelOptions\n")
	}

	var options []string

	for _, channel := range thisFixture.Channels {
		if channel.Name == channelName {
			for _, setting := range channel.Settings {
				options = append(options, setting.Name)
			}
		}
	}
	return options
}

func GetSwitchFixtureType(switchNumber int, stateNumber int16, fixturesConfig *Fixtures) string {
	for _, fixture := range fixturesConfig.Fixtures {
		if fixture.Type == "switch" {
			if fixture.Number == switchNumber+1 {

				useFixture, _ := GetFixtureDetailsByLabel(fixture.UseFixture, fixturesConfig)
				if debug {
					fmt.Printf("fixture number %d name %s use fixture %s type %s\n", fixture.Number, fixture.Name, fixture.UseFixture, useFixture.Type)
				}
				return useFixture.Type
			}
		}
	}
	return ""
}

func GetSwitchStateIsMusicTriggerOn(switchNumber int, stateNumber int16, fixturesConfig *Fixtures) bool {

	for _, fixture := range fixturesConfig.Fixtures {
		if fixture.Type == "switch" {
			if fixture.Number == switchNumber+1 {
				if debug {
					fmt.Printf("fixture number %d %s\n", fixture.Number, fixture.Name)
				}
				for _, state := range fixture.States {
					if state.Number == stateNumber {
						if debug {
							fmt.Printf("state number %d %+v\n", stateNumber, state.Actions)
						}
						if state.Actions != nil {
							for _, action := range state.Actions {
								if action.Mode == "Chase" && action.Speed != "Music" {
									return true
								}
							}
						}
					}
				}
			}
		}
	}
	return false
}

func GetSwitchAction(switchNumber int, switchState int16, fixturesConfig *Fixtures) Action {

	for _, fixture := range fixturesConfig.Fixtures {
		if fixture.Type == "switch" {
			if fixture.Number == switchNumber {
				if debug {
					fmt.Printf("found fixture number %d name %s type %s\n", fixture.Number, fixture.Name, fixture.Type)
				}
				for _, state := range fixture.States {
					if state.Number == switchState {
						if debug {
							fmt.Printf("looking for state %d have state number %d Actions %+v\n", state.Number, state.Number, state.Actions)
						}
						var action Action
						var actionNumber int
						if state.Actions != nil {
							for actionNumber, action = range state.Actions {
								if action.Mode == "Control" {
									if action.Name == "Off" {
										action.Colors = []string{"Green"}
									}
									if action.Name == "On" {
										action.Colors = []string{"Red"}
									}
								}
							}
							if debug {
								fmt.Printf("Actions:- Mode %s action number %d name %s colors %+v\n", action.Mode, actionNumber, action.Name, action.Colors)
							}
							return action
						}

						if state.Settings != nil {
							if debug {
								for _, setting := range state.Settings {
									fmt.Printf("setting Number %d Label %s Channel %s Valuue %s\n", setting.Number, setting.Label, setting.Channel, setting.Value)
								}
							}
							action := convertSettingToAction(fixture, state.Settings)
							return action
						}
					}
				}
			}
		}
	}
	return Action{Name: "Not Found"}
}

func IsThisAChannel(thisFixture *Fixture, channelName string) bool {

	if thisFixture == nil {
		return false
	}

	for _, channel := range thisFixture.Channels {
		if channel.Name == channelName {
			if debug {
				fmt.Printf("\tisThisAChannel fixture %s channelName %s true\n", thisFixture.Name, channelName)
			}
			return true
		}
	}
	return false
}

// returns true is they are the same.
func CheckFixturesAreTheSame(fixtures *Fixtures, startConfig *Fixtures) (bool, string) {

	if len(fixtures.Fixtures) != len(startConfig.Fixtures) {
		return false, "Number of fixtures are different"
	}

	for fixtureNumber, fixture := range fixtures.Fixtures {

		if debug {
			fmt.Printf("Checking Fixture %s against %s\n", fixture.Name, startConfig.Fixtures[fixtureNumber].Name)
		}

		if fixture.Name != startConfig.Fixtures[fixtureNumber].Name {
			return false, fmt.Sprintf("Fixture:%d Name is different\n", fixtureNumber+1)
		}

		if fixture.ID != startConfig.Fixtures[fixtureNumber].ID {
			return false, fmt.Sprintf("Fixture:%d ID is different\n", fixtureNumber+1)
		}

		if fixture.Label != startConfig.Fixtures[fixtureNumber].Label {
			return false, fmt.Sprintf("Fixture:%d Label is different\n", fixtureNumber+1)
		}

		if fixture.Number != startConfig.Fixtures[fixtureNumber].Number {
			return false, fmt.Sprintf("Fixture:%d Number is different\n", fixtureNumber+1)
		}

		if fixture.Description != startConfig.Fixtures[fixtureNumber].Description {
			return false, fmt.Sprintf("Fixture:%d Description is different\n", fixtureNumber+1)
		}

		if fixture.Type != startConfig.Fixtures[fixtureNumber].Type {
			return false, fmt.Sprintf("Fixture:%d Type is different\n", fixtureNumber+1)
		}

		if fixture.Group != startConfig.Fixtures[fixtureNumber].Group {
			return false, fmt.Sprintf("Fixture:%d Group is different\n", fixtureNumber+1)
		}

		if fixture.Address != startConfig.Fixtures[fixtureNumber].Address {
			return false, fmt.Sprintf("Fixture:%d Address is different\n", fixtureNumber+1)
		}

		for channelNumber, channel := range fixture.Channels {

			if channel.Number != startConfig.Fixtures[fixtureNumber].Channels[channelNumber].Number {
				return false, fmt.Sprintf("Fixture:%d Channel Number is different\n", fixtureNumber+1)
			}

			if channel.Name != startConfig.Fixtures[fixtureNumber].Channels[channelNumber].Name {
				return false, fmt.Sprintf("Fixture:%d Channel Name is different\n", fixtureNumber+1)
			}

			if channel.Value != startConfig.Fixtures[fixtureNumber].Channels[channelNumber].Value {
				return false, fmt.Sprintf("Fixture:%d Channel Value is different\n", fixtureNumber+1)
			}

			if channel.MaxDegrees != startConfig.Fixtures[fixtureNumber].Channels[channelNumber].MaxDegrees {
				return false, fmt.Sprintf("Fixture:%d Channel MaxDegrees is different\n", fixtureNumber+1)
			}

			if channel.Offset != startConfig.Fixtures[fixtureNumber].Channels[channelNumber].Offset {
				return false, fmt.Sprintf("Fixture:%d Channel Offset is different\n", fixtureNumber+1)
			}

			if channel.Comment != startConfig.Fixtures[fixtureNumber].Channels[channelNumber].Comment {
				return false, fmt.Sprintf("Fixture:%d Channel Comment is different\n", fixtureNumber+1)
			}

			for settingNumber, setting := range channel.Settings {

				if setting.Name != startConfig.Fixtures[fixtureNumber].Channels[channelNumber].Settings[settingNumber].Name {
					return false, fmt.Sprintf("Fixture:%d Channel Settings Name is different\n", fixtureNumber+1)
				}

				if setting.Label != startConfig.Fixtures[fixtureNumber].Channels[channelNumber].Settings[settingNumber].Label {
					return false, fmt.Sprintf("Fixture:%d Channel Settings Label is different\n", fixtureNumber+1)
				}

				if setting.Number != startConfig.Fixtures[fixtureNumber].Channels[channelNumber].Settings[settingNumber].Number {
					return false, fmt.Sprintf("Fixture:%d Channel Settings Number is different\n", fixtureNumber+1)
				}

				if setting.Channel != startConfig.Fixtures[fixtureNumber].Channels[channelNumber].Settings[settingNumber].Channel {
					return false, fmt.Sprintf("Fixture:%d Channel Channel Number is different\n", fixtureNumber+1)
				}

				if setting.Value != startConfig.Fixtures[fixtureNumber].Channels[channelNumber].Settings[settingNumber].Value {
					return false, fmt.Sprintf("Fixture:%d Channel Value Number is different\n", fixtureNumber+1)
				}
			}

			for stateNumber, state := range fixture.States {

				if state.Number != startConfig.Fixtures[fixtureNumber].States[stateNumber].Number {
					return false, fmt.Sprintf("Fixture:%d State Number is different\n", fixtureNumber+1)
				}

				if state.Name != startConfig.Fixtures[fixtureNumber].States[stateNumber].Name {
					return false, fmt.Sprintf("Fixture:%d State Name is different\n", fixtureNumber+1)
				}

				if state.Label != startConfig.Fixtures[fixtureNumber].States[stateNumber].Label {
					return false, fmt.Sprintf("Fixture:%d State Label is different\n", fixtureNumber+1)
				}

				if state.ButtonColor != startConfig.Fixtures[fixtureNumber].States[stateNumber].ButtonColor {
					return false, fmt.Sprintf("Fixture:%d State ButtonColor is different\n", fixtureNumber+1)
				}

				if state.Master != startConfig.Fixtures[fixtureNumber].States[stateNumber].Master {
					return false, fmt.Sprintf("Fixture:%d State Master is different\n", fixtureNumber+1)
				}

				for actionNumber, action := range state.Actions {

					if action.Name != startConfig.Fixtures[fixtureNumber].States[stateNumber].Actions[actionNumber].Name {
						return false, fmt.Sprintf("Fixture:%d State Action Name is different\n", fixtureNumber+1)
					}

					if action.Number != startConfig.Fixtures[fixtureNumber].States[stateNumber].Actions[actionNumber].Number {
						return false, fmt.Sprintf("Fixture:%d State Action Number is different\n", fixtureNumber+1)
					}

					for colorNumber, color := range action.Colors {
						if color != startConfig.Fixtures[fixtureNumber].States[stateNumber].Actions[actionNumber].Colors[colorNumber] {

							return false, fmt.Sprintf("Fixture:%d State Action Color Number is different\n", fixtureNumber+1)
						}
					}

					if action.Mode != startConfig.Fixtures[fixtureNumber].States[stateNumber].Actions[actionNumber].Mode {
						return false, fmt.Sprintf("Fixture:%d State Action Mode is different\n", fixtureNumber+1)
					}

					if action.Fade != startConfig.Fixtures[fixtureNumber].States[stateNumber].Actions[actionNumber].Fade {
						return false, fmt.Sprintf("Fixture:%d State Action Fade is different\n", fixtureNumber+1)
					}

					if action.Size != startConfig.Fixtures[fixtureNumber].States[stateNumber].Actions[actionNumber].Size {
						return false, fmt.Sprintf("Fixture:%d State Action Size is different\n", fixtureNumber+1)
					}

					if action.Speed != startConfig.Fixtures[fixtureNumber].States[stateNumber].Actions[actionNumber].Speed {
						return false, fmt.Sprintf("Fixture:%d State Action Speed is different\n", fixtureNumber+1)
					}

					if action.Rotate != startConfig.Fixtures[fixtureNumber].States[stateNumber].Actions[actionNumber].Rotate {
						return false, fmt.Sprintf("Fixture:%d State Action Rotate is different\n", fixtureNumber+1)
					}

					if action.RotateSpeed != startConfig.Fixtures[fixtureNumber].States[stateNumber].Actions[actionNumber].RotateSpeed {
						return false, fmt.Sprintf("Fixture:%d State Action RotateSpeed is different\n", fixtureNumber+1)
					}

					if action.Program != startConfig.Fixtures[fixtureNumber].States[stateNumber].Actions[actionNumber].Program {
						return false, fmt.Sprintf("Fixture:%d State Action Program is different\n", fixtureNumber+1)
					}

					if action.Strobe != startConfig.Fixtures[fixtureNumber].States[stateNumber].Actions[actionNumber].Strobe {
						return false, fmt.Sprintf("Fixture:%d State Action Strobe is different\n", fixtureNumber+1)
					}

				}

				for settingNumber, setting := range state.Settings {

					if setting.Name != startConfig.Fixtures[fixtureNumber].States[stateNumber].Settings[settingNumber].Name {
						return false, fmt.Sprintf("Fixture:%d Channel Settings Strobe is different\n", fixtureNumber+1)
					}

					if setting.Label != startConfig.Fixtures[fixtureNumber].States[stateNumber].Settings[settingNumber].Label {
						return false, fmt.Sprintf("Fixture:%d Channel Settings Label is different\n", fixtureNumber+1)
					}

					if setting.Number != startConfig.Fixtures[fixtureNumber].States[stateNumber].Settings[settingNumber].Number {
						return false, fmt.Sprintf("Fixture:%d Channel Settings Number is different\n", fixtureNumber+1)
					}

					if setting.Channel != startConfig.Fixtures[fixtureNumber].States[stateNumber].Settings[settingNumber].Channel {
						return false, fmt.Sprintf("Fixture:%d Channel Channel Number is different\n", fixtureNumber+1)
					}

					if setting.Value != startConfig.Fixtures[fixtureNumber].States[stateNumber].Settings[settingNumber].Value {
						return false, fmt.Sprintf("Fixture:%d Channel Value Number is different\n", fixtureNumber+1)
					}
				}

				if state.Flash != startConfig.Fixtures[fixtureNumber].States[stateNumber].Flash {
					return false, fmt.Sprintf("Fixture:%d State Flash is different\n", fixtureNumber+1)
				}

			}

			if fixture.MultiFixtureDevice != startConfig.Fixtures[fixtureNumber].MultiFixtureDevice {
				return false, fmt.Sprintf("Fixture:%d MultiFixtureDevice is different\n", fixtureNumber+1)
			}

			if fixture.NumberSubFixtures != startConfig.Fixtures[fixtureNumber].NumberSubFixtures {
				return false, fmt.Sprintf("Fixture:%d NumberSubFixtures is different\n", fixtureNumber+1)
			}

			if fixture.UseFixture != startConfig.Fixtures[fixtureNumber].UseFixture {
				return false, fmt.Sprintf("Fixture: %d UseFixture is different\n", fixtureNumber+1)
			}

		}
	}

	return true, ""
}

// GetAvailableScannerGobos - populates a map indexed by fixture number for the sequenceNumber provided.
// Each fixture contains an array of StaticColorButtons, essentially info representing each gobo in this fixture.
// Gobo details provided are - Name, label, number. DMX value and color.
func GetAvailableScannerGobos(sequenceNumber int, fixtures *Fixtures) map[int][]common.StaticColorButton {
	if debug {
		fmt.Printf("getAvailableScannerGobos\n")
	}

	gobos := make(map[int][]common.StaticColorButton)

	for _, f := range fixtures.Fixtures {
		if debug {
			fmt.Printf("Fixture Name:%s\n", f.Name)
		}
		if f.Type == "scanner" {

			if debug {
				fmt.Printf("Sequence: %d - Scanner Name: %s Description: %s\n", sequenceNumber, f.Name, f.Description)
			}
			for _, channel := range f.Channels {
				if channel.Name == "Gobo" {
					newGobo := common.StaticColorButton{}
					for _, setting := range channel.Settings {
						newGobo.Name = setting.Name
						newGobo.Label = setting.Label
						newGobo.Number = setting.Number
						v, _ := strconv.Atoi(setting.Value)
						newGobo.Setting = v
						newGobo.Color = colors.Yellow
						gobos[f.Number] = append(gobos[f.Number], newGobo)
						if debug {
							fmt.Printf("\tGobo: %s Setting: %s\n", setting.Name, setting.Value)
						}
					}
				}
			}
		}
	}
	return gobos
}

func GetAvailableColors(fixture *Fixture) []string {

	colors := []string{}

	for _, channel := range fixture.Channels {

		if channel.Name == "Color" {
			for _, setting := range channel.Settings {
				colors = append(colors, setting.Name)
			}
		}
	}

	return colors
}

// getAvailableScannerColors looks through the fixtures list and finds scanners that
// have colors defined in their config. It then returns an array of these available colors.
// Also returns a map of the default values for each scanner that has colors.
func GetAvailableScannerColors(fixtures *Fixtures) (map[int][]common.StaticColorButton, map[int]int) {

	if debug {
		fmt.Printf("GetAvailableScannerColors for fixture\n")
	}
	scannerColors := make(map[int]int)

	availableScannerColors := make(map[int][]common.StaticColorButton)
	for _, fixture := range fixtures.Fixtures {
		if fixture.Type == "scanner" {
			for _, channel := range fixture.Channels {
				if strings.Contains(channel.Name, "Color") {
					for _, setting := range channel.Settings {
						newStaticColorButton := common.StaticColorButton{}
						newStaticColorButton.SelectedColor = setting.Number
						settingColor := common.GetRGBColorByName(setting.Name)
						newStaticColorButton.Color = settingColor
						availableScannerColors[fixture.Number] = append(availableScannerColors[fixture.Number], newStaticColorButton)
						scannerColors[fixture.Number-1] = 0
					}
				}
			}
		}
	}
	return availableScannerColors, scannerColors
}

// GetScannerColorName finds the color for given scanner and color number.
func GetScannerColorName(scannerNumber int, colorNumber int, fixtures *Fixtures) (color.RGBA, error) {

	if debug {
		fmt.Printf("GetScannerColorName() Looking for Color Number %d\n", colorNumber)
	}

	for _, fixture := range fixtures.Fixtures {
		if fixture.Type == "scanner" {
			if fixture.Number == scannerNumber+1 {
				for _, channel := range fixture.Channels {
					if strings.Contains(channel.Name, "Color") {
						for _, setting := range channel.Settings {
							if setting.Number == colorNumber+1 {
								settingColor := common.GetRGBColorByName(setting.Name)
								return settingColor, nil
							}
						}
					}
				}
			}
		}
	}
	return color.RGBA{}, fmt.Errorf("color not found")
}

func HowManyScannerColors(sequence *common.Sequence, fixturesConfig *Fixtures) []color.RGBA {

	if debug {
		fmt.Printf("HowManyScannerColors: \n")
	}

	// Clear out sequemce colors.
	sequence.SequenceColors = []color.RGBA{}

	for scannerNumber := 0; scannerNumber < sequence.NumberFixtures; scannerNumber++ {
		// Look at all the scannes and add their selected color to the color display.
		colorNumber := sequence.ScannerColor[scannerNumber]
		// Get the color name from the fixture config, ignore scanner that don't have a color set.
		color, err := GetScannerColorName(scannerNumber, colorNumber, fixturesConfig)
		if err == nil {
			if debug {
				fmt.Printf("HowManyScannerColors()Scanner %d is Color %s\n", scannerNumber, common.GetColorNameByRGB(color))
			}
			sequence.SequenceColors = append(sequence.SequenceColors, color)
		}
	}
	if debug {
		fmt.Printf("HowManyScannerColors() colors %+v\n", sequence.SequenceColors)
	}

	return sequence.SequenceColors
}

func GetNumberOfFixturesInGroup(sequenceNumber int, fixturesConfig *Fixtures) int {

	var recents []int

	if debug {
		fmt.Printf("\nHowManyFixturesInGroup for sequence %d\n", sequenceNumber)
	}

	var count int
	for _, fixture := range fixturesConfig.Fixtures {

		// Found the group.
		if fixture.Group == sequenceNumber+1 {

			if debug {
				fmt.Printf("\t%d: Found fixture in group %d\n", sequenceNumber, fixture.Number)
			}

			// Have we seen this fixture number already
			if !haveWeSeenThisBefore(recents, fixture.Number) {

				// If this is a multifixture device
				if fixture.MultiFixtureDevice {

					if debug {
						fmt.Printf("\t\t%d: Found MultiFixtureDevice %d\n", sequenceNumber, fixture.Number)
					}

					// Then count the number of RGB channels.
					for _, channel := range fixture.Channels {
						if strings.Contains(channel.Name, "Red") {
							if debug {
								fmt.Printf("\t\t\t%d: Found Red Channel %d\n", sequenceNumber, fixture.Number)
							}
							count++
						}
					}
					break
				}

				// Only count this one if we haven't counted it already.
				if debug {
					fmt.Printf("\t\tAdd fixture in recents %d\n", fixture.Number)
				}
				recents = append(recents, fixture.Number)
				count++
			}
		}
	}

	if debug {
		fmt.Printf("%d: Found %d fixtures\n", sequenceNumber, count)
	}
	return count
}

func haveWeSeenThisBefore(recents []int, fixtureNumber int) bool {

	for _, recent := range recents {
		if fixtureNumber == recent {
			return true
		}

	}
	return false
}
