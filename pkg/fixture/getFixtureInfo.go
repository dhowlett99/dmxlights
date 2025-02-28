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
	"strconv"
	"strings"
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
func GetFixtureDetailsByLabel(label string, fixtures *Fixtures) (Fixture, error) {
	// scan the fixtures structure for the selected fixture.
	if debug {
		fmt.Printf("GetFixtureDetailsByLabel: Looking for Fixture by Label %s\n", label)
	}

	for _, fixture := range fixtures.Fixtures {
		if debug {
			fmt.Printf("Fixture Label %s and Name %s States %+v\n", fixture.Label, fixture.Name, fixture.States)
		}
		if fixture.Label == label {
			return fixture, nil
		}
	}
	return Fixture{}, fmt.Errorf("error: fixture label %s not found", label)
}

func GetChannelSettingByChannelNameAndSettingName(fixture *Fixture, channelName string, settingName string) (int, error) {

	if debug {
		fmt.Printf("GetChannelSettingByChannelNameAndSettingName for fixture %s on channel %s setting %s\n", fixture.Name, channelName, settingName)
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

	return 0, fmt.Errorf("GetChannelSettingByChannelNameAndSettingName: setting %s not found in channel %s for fixture %s", settingName, channelName, fixture.Name)
}

func GetChannelSettingByNameAndSpeed(fixtureName string, channelName string, settingName string, settingSpeed string, fixtures *Fixtures) (int, error) {

	if debug {
		fmt.Printf("findChannelSettingByNameAndSpeed for fixture %s channel name %s setting name %s and setting speed %s\n", fixtureName, channelName, settingName, settingSpeed)
	}

	if channelName == "" {
		return 0, fmt.Errorf("findChannelSettingByNameAndSpeed: error channel name is empty")
	}
	if settingName == "" {
		return 0, fmt.Errorf("findChannelSettingByNameAndSpeed: error setting name is empty")
	}
	if settingSpeed == "" {
		return 0, fmt.Errorf("findChannelSettingByNameAndSpeed: error setting speed is empty")
	}

	for _, fixture := range fixtures.Fixtures {

		if fixtureName == fixture.Name {
			for _, channel := range fixture.Channels {

				if channel.Name == channelName {
					if debug {
						fmt.Printf("fixture %s: inspect channel %s for %s\n", fixture.Name, channel.Name, settingName)
					}

					for _, setting := range channel.Settings {
						if debug {
							fmt.Printf("inspect setting %+v \n", setting)
							fmt.Printf("got:setting.Name %s  want name %s speed %s\n", setting.Name, settingName, settingSpeed)
						}
						if strings.Contains(setting.Name, settingName) && strings.Contains(setting.Name, settingSpeed) {

							if debug {
								fmt.Printf("FixtureName=%s ChannelName=%s SettingName=%s SettingSpeed=%s, SettingValue=%s\n", fixture.Name, channel.Name, settingName, settingSpeed, setting.Value)
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
								fmt.Printf("speed found is %d\n", v)
							}
							return v, nil
						}
					}
				}
			}
		}
	}

	return 0, fmt.Errorf("warning: channel %s setting %s not found in fixture :%s", channelName, settingSpeed, fixtureName)
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

func GetFixtureByLabel(label string, fixtures *Fixtures) (*Fixture, error) {

	if debug {
		fmt.Printf("FindFixtureByLabel label is %s\n", label)
	}
	for _, fixture := range fixtures.Fixtures {
		if fixture.Label == label {
			if debug {
				fmt.Printf("Found fixture %s Group %d Number %d Address %d\n", fixture.Name, fixture.Group, fixture.Number, fixture.Address)
			}
			return &fixture, nil
		}
	}
	return nil, fmt.Errorf("FindFixtureByGroupAndNumber: failed to find fixture by label %s", label)
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

// GetRotateDMXValueByIndex takes the rotate setting number and returns the DMX value which selects this speed.
func GetRotateDMXValueByIndex(fixture *Fixture, index int) int {

	if debug {
		fmt.Printf("GetRotateDMXValueByIndex Looking for rotate speed index %d in fixture %s\n", index, fixture.Name)
	}

	for _, channel := range fixture.Channels {
		if strings.Contains(channel.Name, "Rotate") {
			for _, setting := range channel.Settings {
				if setting.Number == index {
					if debug {
						fmt.Printf("Rotate Speed %d Name %s\n", setting.Number, setting.Name)
					}
					dmx, _ := strconv.Atoi(setting.Value)
					return dmx
				}
			}
		}
	}

	return 0
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
					if debug {
						fmt.Printf("Gobo %d Name %s\n", setting.Number, setting.Name)
					}
					dmx, _ := strconv.Atoi(setting.Value)
					return dmx
				}
			}
		}
	}

	return 0
}

func GetAvailableSpeedChannelsByFixure(fixture *Fixture) []string {

	var speeds []string
	for _, channel := range fixture.Channels {
		if strings.Contains(channel.Name, "Speed") {
			for _, setting := range channel.Settings {
				speeds = append(speeds, setting.Name)
			}
		}
	}
	return speeds
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

func GetAvailableColorsByFixure(fixture *Fixture) []string {

	var colors []string
	for _, channel := range fixture.Channels {
		if strings.Contains(channel.Name, "Color") {
			for _, setting := range channel.Settings {
				colors = append(colors, setting.Name)
			}
		}
	}
	return colors
}

func GetAvailableGobosByFixure(fixture *Fixture) []string {

	var colors []string
	for _, channel := range fixture.Channels {
		if strings.Contains(channel.Name, "Gobo") {
			for _, setting := range channel.Settings {
				colors = append(colors, setting.Name)
			}
		}
	}
	return colors
}

// GetColorNameByNumber takes the color number and returns the color name for this fixture.
func GetColorNameByNumber(fixture *Fixture, number int) string {

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

	return "Unknown"
}

// GetGoboByName takes the name of a gobo channel setting like "Open" and returns the gobo number  for this type of scanner.
func GetGoboByName(myFixtureNumber int, mySequenceNumber int, selectedGobo string, fixtures *Fixtures) int {

	if debug {
		fmt.Printf("GetGoboByName\n")
	}

	for _, fixture := range fixtures.Fixtures {
		if fixture.Group == mySequenceNumber+1 {
			if fixture.Number == myFixtureNumber+1 {
				for _, channel := range fixture.Channels {
					if strings.Contains(channel.Name, "Gobo") {
						for _, setting := range channel.Settings {
							if strings.Contains(setting.Name, selectedGobo) {
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
