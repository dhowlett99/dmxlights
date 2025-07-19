// Copyright (C) 2022,2023,2024,2025 dhowlett99.
// This is the dmxlights fixture configuration code, it is used by
// the mini sequencer and the mini setter.
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
	"time"

	"github.com/dhowlett99/dmxlights/pkg/colors"
	"github.com/dhowlett99/dmxlights/pkg/common"
)

func GetConfig(action Action, fixture *Fixture, fixturesConfig *Fixtures) ActionConfig {

	config := ActionConfig{}

	var programSettings []common.Setting
	var goboSettings []common.Setting
	var err error

	fixtureInfo := GetFixtureInfo(fixture, fixturesConfig)
	if debug {
		fmt.Printf("GetConfig() This fixture %s has FixtureInfo %+v\n", fixture.Name, fixtureInfo)
	}

	// Pass through which mode we are in.
	config.Mode = action.Mode

	// Find all the specified settings for the program channel
	if fixtureInfo.HasProgram {

		programSettings, err = GetChannelSettinsByName(fixture, "Program", fixturesConfig)
		if err != nil && debug {
			fmt.Printf("newMiniSequencer: warning! no program settings found for fixture %s\n", fixture.Name)
		}

		if debug {
			fmt.Printf("Fixture %s Available Program Settings %+v\n", fixture.Name, programSettings)
		}

		// Program Speed - Speed of programs or shows.
		config.ProgramSpeed = GetChannelSettingInfo(fixture, "ProgramSpeed", action.ProgramSpeed, fixturesConfig)

		// Look through the available program channel settins and see if you can find the specified program action.
		for settingNumber, setting := range programSettings {
			if debug {
				fmt.Printf("Fixture %s Setting Number %d Setting Name %s Action Name %s Action Program %s\n", fixture.Name, settingNumber, setting.Name, action.Name, action.Program)
			}

			// If a action program is a number look up program name in program channel.
			if IsNumericOnly(action.Program) {
				name := GetProgramNameByNumber(fixture, action.Program)
				action.Program = name
			}

			// Now assign the program based on name.
			if setting.Name == action.Program || setting.Name == "Default" || setting.Name == "Blackout" {
				config.Program = settingNumber + 1
				config.ProgramName = setting.Name
				if debug {
					fmt.Printf("---> Setting Program to %d\n", config.Program)
				}
			}
		}
	}

	// Setup the gobos based on their name.
	if fixtureInfo.HasGobo {

		switch action.GoboSpeed {
		case "Slow":
			config.GoboSpeed = SENSITIVITY_LONG
		case "Medium":
			config.GoboSpeed = SENSITIVITY_MEDIUM
		case "Fast":
			config.GoboSpeed = SENSITIVITY_SHORT
		default:
			config.GoboSpeed = 0
		}

		switch action.Gobo {
		case "":
			config.Gobo = -1
			config.AutoGobo = false
		case "Default":
			config.Gobo = 0
			config.AutoGobo = false
		case "Auto":
			config.Gobo = -1
			config.AutoGobo = true
		default:
			// find current gobo number.
			config.Gobo = GetGoboByName(fixture.Number-1, fixture.Group-1, action.Gobo, fixturesConfig)
			config.AutoGobo = false
		}

		// Find all the specified options for the gobo channel
		goboSettings, err = GetChannelSettinsByName(fixture, "Gobo", fixturesConfig)
		if err != nil && debug {
			fmt.Printf("newMiniSequencer: warning! no gobos found for fixture %s\n", fixture.Name)
		}

		// Gobo allways has a defult option of the first gobo.
		config.GoboOptions = append(config.GoboOptions, "Default")

		// Look through the available gobos and populate the available gobos values.
		for _, setting := range goboSettings {
			config.GoboOptions = append(config.GoboOptions, setting.Name)
		}

	}

	// If we have a color channel in the fixture get the color setting number based on color in the action.
	if fixture.FixtureInfo.HasColorChannel {
		// Find the avialable colors as listed in the color channel settings.
		config.AvailableColors = GetAvailableColors(fixture)

		// Look up the first color from the action in the available colors for this fixture.
		if len(action.Colors) > 0 {
			config.ColorName = action.Colors[0]
			// Color start at 0 , color settings start at 1 so -1 below.
			config.Color = GetColorNumberFromFixture(fixture, action.Colors[0]) - 1
		} else {
			config.Color = 0 // Just use the first color available.
		}
	}

	// If we have Red, Green, Blue color channels look for a matching color in
	// the standard color pallete based on the first color in the actions.
	if fixture.FixtureInfo.HasRGBChannels {

		config.AvailableColors = colors.GetAvailableColorsAsStrings()
		if len(action.Colors) > 0 {
			// Look up the first color from the action in the standard color library.
			config.ColorName = action.Colors[0]
			config.Color = colors.FindColorIndexByName(config.AvailableColors, action.Colors[0])
		}
	}

	// If we don't have any RGB channels or a dedicated color wheel just configure black and white as
	// available colors so the chaser can run.
	if !fixture.FixtureInfo.HasColorChannel && !fixture.FixtureInfo.HasRGBChannels {
		config.AvailableColors = append(config.AvailableColors, common.GetColorNameByRGB(colors.Black))
		config.AvailableColors = append(config.AvailableColors, common.GetColorNameByRGB(colors.White))
	}

	// Map - A switch to map the brightness to the master dimmer, useful for fixtures that don't have RGB.
	switch action.Map {
	case "Off":
		config.Map = false // Don't map
	case "On":
		config.Map = true // Map brightness to master dimmer.
	default:
		config.Map = false // Don't map
	}

	// Fade - Time taken to fade up and down.
	switch action.Fade {
	case "":
		config.Fade = FADE_SHARP
	case "Off":
		config.Fade = FADE_SHARP
	case "Soft":
		config.Fade = FADE_SOFT
	case "Normal":
		config.Fade = FADE_NORMAL
	case "Sharp":
		config.Fade = FADE_SHARP
	default:
		config.Fade = FADE_SHARP
	}

	// Size - How long does the lamp stay on.
	switch action.Size {
	case "Off":
		config.Size = SIZE_OFF
	case "Short":
		config.Size = SIZE_SHORT
	case "Medium":
		config.Size = SIZE_MEDIUM
	case "Long":
		config.Size = SIZE_LONG
	default:
		config.Size = SIZE_MEDIUM
	}

	// Shift - with only four fixture in the mini sequencer
	// no more of a shift of 1 makes sense.
	// TODO work out how to make shift work correctly for 4 fixtures.
	config.Shift = 1

	// Deal with the rotate parameters but only if there is a Rotate channel.
	if IsThisAChannel(fixture, "Rotate") {
		config = setRotate(fixture, config, action, fixturesConfig)
	}

	switch action.Speed {
	case "Slow":
		config.TriggerState = false
		config.Speed = 2
		config.SpeedDuration = common.SetSpeed(config.Speed)
		config.MusicTrigger = false
		config.NumberSteps = LARGE_NUMBER_STEPS
		config.RotateSensitivity = SENSITIVITY_SHORT
	case "Medium":
		config.TriggerState = false
		config.Speed = 4
		config.SpeedDuration = common.SetSpeed(config.Speed)
		config.MusicTrigger = false
		config.NumberSteps = LARGE_NUMBER_STEPS
		config.RotateSensitivity = SENSITIVITY_SHORT
	case "Fast":
		config.TriggerState = false
		config.Speed = 8
		config.SpeedDuration = common.SetSpeed(config.Speed)
		config.MusicTrigger = false
		config.NumberSteps = LARGE_NUMBER_STEPS
		config.RotateSensitivity = SENSITIVITY_SHORT
	case "VeryFast":
		config.TriggerState = false
		config.Speed = 12
		config.SpeedDuration = common.SetSpeed(config.Speed)
		config.MusicTrigger = false
		config.NumberSteps = LARGE_NUMBER_STEPS
		config.RotateSensitivity = SENSITIVITY_SHORT
	case "Music":
		config.TriggerState = true
		config.SpeedDuration = time.Duration(12 * time.Hour)
		config.MusicTrigger = true
		config.NumberSteps = MEDIUM_NUMBER_STEPS
		config.RotateSensitivity = SENSITIVITY_LONG
	default:
		config.TriggerState = false
		config.Speed = 1
		config.SpeedDuration = time.Duration(12 * time.Hour)
		config.MusicTrigger = false
		config.NumberSteps = MEDIUM_NUMBER_STEPS
		config.RotateSensitivity = SENSITIVITY_SHORT
	}

	switch action.Strobe {

	// TODO Lookup Strobe Speed in Fixture HERE
	case "Off":
		config.Strobe = false
		config.StrobeSpeed = STROBE_SPEED_SLOW
	case "Slow":
		config.Strobe = true
		config.StrobeSpeed = STROBE_SPEED_SLOW
	case "Medium":
		config.Strobe = true
		config.StrobeSpeed = STROBE_SPEED_MEDIUM
	case "Fast":
		config.Strobe = true
		config.StrobeSpeed = STROBE_SPEED_FAST
	default:
		config.Strobe = false
		config.StrobeSpeed = 0
	}

	// TODO add scanner support to mini sequencer.
	config.Pan = 0
	config.Tilt = 0

	if debug {
		fmt.Printf("Config %+v\n", config)
	}

	return config
}

// Given the fixture and the list of settings for this state
// buill a new action that represents the set of settings.
func convertSettingToAction(fixture Fixture, settings []Setting) Action {

	newAction := Action{}

	newAction.Mode = "Setting"
	newAction.Name = "Setting"
	newAction.Number = 1

	var hasRGBChannels bool

	var red int
	var green int
	var blue int

	// Look through settings and buuld up the new action.
	for _, setting := range settings {

		if debug {
			fmt.Printf("convertSettingToAction: Fixture name %s setting name %s label %s Channel %s name %s value %s\n", fixture.Name, setting.Name, setting.Label, setting.Channel, setting.Name, setting.Value)
		}

		if setting.Channel == "Speed" {
			newAction.Speed = setting.Value
		}
		if setting.Channel == "Fade" {
			newAction.Fade = setting.Value
		}

		if setting.Channel == "Size" {
			newAction.Fade = setting.Value
		}

		if setting.Channel == "Rotate" {
			// The rotate channel has to have settings which include the
			// direction and speed e.g. Forward Slow
			// TODO Remove this hard coding and make it configurable.
			newAction.Rotate = setting.Name
			if strings.Contains(setting.Name, "Slow") ||
				strings.Contains(setting.Name, "Medium") ||
				strings.Contains(setting.Name, "Fast") {

				// Find the rotate speed, Slow Medium Fast etc.
				rotateSettings := strings.Split(setting.Name, " ")
				newAction.RotateSpeed = rotateSettings[1]
			}
		}

		if setting.Channel == "Program" {
			newAction.Program = setting.Value
		}

		if setting.Channel == "ProgramSpeed" {
			newAction.ProgramSpeed = setting.Name
		}

		if strings.Contains(setting.Name, "Red") || strings.Contains(setting.Name, "red") {
			red, _ = strconv.Atoi(setting.Value)
			hasRGBChannels = true
		}
		if strings.Contains(setting.Name, "Green") || strings.Contains(setting.Name, "green") {
			green, _ = strconv.Atoi(setting.Value)
			hasRGBChannels = true
		}
		if strings.Contains(setting.Name, "Blue") || strings.Contains(setting.Name, "blue") {
			blue, _ = strconv.Atoi(setting.Value)
			hasRGBChannels = true
		}

		if hasRGBChannels {
			actualColor := color.RGBA{
				R: uint8(red),
				G: uint8(green),
				B: uint8(blue),
				A: 255,
			}
			colorName := common.GetColorNameByRGB(actualColor)
			newAction.Colors = []string{colorName}
		}

		// A channel setting can only contain one value
		// so only one color.
		if setting.Channel == "Color" {
			hasRGBChannels = false
			// If a setting has a channel name which is a number we lookup that color name.
			if IsNumericOnly(setting.Name) {
				if colorNumber, err := strconv.Atoi(setting.Value); err == nil {
					// Lookup color number in list of available colors.
					colorName := GetColorNameByNumber(&fixture, colorNumber)
					newAction.Colors = []string{colorName}
				}
			} else {
				// we use that string as the color.
				newAction.Colors = []string{setting.Name}
			}
			if setting.Name == "Off" {
				newAction.Name = "Off"
				newAction.Program = "Off"
				newAction.Colors = []string{"Green"}
			}
			if setting.Name == "On" {
				newAction.Name = "On"
				newAction.Program = "On"
				newAction.Colors = []string{"Red"}
			}
		}

		if setting.Channel == "Strobe" {
			newAction.Strobe = setting.Value
		}

		if setting.Channel == "StrobeSpeed" {
			newAction.StrobeSpeed = setting.Value
		}

		if setting.Channel == "Gobo" {
			newAction.Gobo = setting.Name
		}

		if setting.Channel == "GoboSpeed" {
			newAction.GoboSpeed = setting.Value
		}

	}
	return newAction
}

// Rotate is a channel with a number of settings,
// Each setting represents a speed, direction or both.
// At this point the rotate action contains the name of the setting.
// The rotate config needs to hold the rotate setting number so we can recall it and
// step through the settings with the override.
// We use the key words Forward and Reverse to represent Clockwise and Anticlockwise respectively.
func setRotate(fixture *Fixture, config ActionConfig, action Action, fixturesConfig *Fixtures) ActionConfig {

	// Lookup the setting number for this rotate name.
	config.RotateName = action.Rotate
	config.Rotate = GetRotateSpeedNumberByName(fixture, action.Rotate)
	config.Rotatable = false
	config.AutoRotate = false
	config.RotateSpeed = GetADMXValue(fixture, config.Rotate, "Rotate")

	if strings.Contains(config.RotateName, "Off") {
		config.Rotatable = false
		config.AutoRotate = false
		config.Forward = false
		config.Reverse = false
		return config
	}

	if strings.Contains(config.RotateName, "Forward") {
		config.Rotatable = true
		config.AutoRotate = false
		config.Forward = true
		config.Reverse = false
		return config
	}

	if strings.Contains(config.RotateName, "Reverse") {
		config.Rotatable = true
		config.AutoRotate = false
		config.Forward = false
		config.Reverse = true
		return config
	}
	if strings.Contains(config.RotateName, "Auto") {
		config.Rotatable = true
		config.AutoRotate = true
		config.Forward = false
		config.Reverse = false
		config.Rotatable = true

		var err error
		// Calculate the rotation speed based on direction and speed.
		if config.Rotatable {
			config.ReverseSpeed, err = GetChannelSettingByNameAndSpeed(fixture, "Rotate", "Reverse", action.RotateSpeed, fixturesConfig)
			if err != nil {
				fmt.Printf("Looking in channel:Rotate for Setting:Reverse %s error: %s\n", action.RotateSpeed, err)
			}
			config.ForwardSpeed, err = GetChannelSettingByNameAndSpeed(fixture, "Rotate", "Forward", action.RotateSpeed, fixturesConfig)
			if err != nil {
				fmt.Printf("Looking in channel:Rotate for Setting:Forward %s error: %s\n", action.RotateSpeed, err)
			}

			if !config.Forward && !config.Reverse {
				config.RotateSpeed = config.Rotate
			}
			if config.Forward {
				config.RotateSpeed = config.ForwardSpeed
			}
			if config.Reverse {
				config.RotateSpeed = config.ReverseSpeed
			}
			if debug_rotate {
				fmt.Printf("RotateSpeed %d\n", config.RotateSpeed)
			}
		}

		return config

	}

	return config
}
