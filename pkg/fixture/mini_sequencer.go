// Copyright (C) 2022,2023,2024,2025 dhowlett99.
// This is the dmxlights mini sequencer, used by the actions to control
// single fixtures.
// Implemented and depends usbdmx.
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
	"os"
	"strconv"
	"time"

	"github.com/dhowlett99/dmxlights/pkg/colors"
	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/pattern"
	"github.com/dhowlett99/dmxlights/pkg/position"
	"github.com/dhowlett99/dmxlights/pkg/sound"
	"github.com/oliread/usbdmx/ft232"
)

const debug_mini bool = false
const debug_override bool = false
const debug_rotate bool = false

const FADE_SHARP int = 10
const FADE_NORMAL int = 5
const FADE_SOFT int = 1

const SIZE_OFF int = 0
const SIZE_SHORT int = 1
const SIZE_MEDIUM int = 3
const SIZE_LONG int = 10

const SENSITIVITY_LONG int = 500
const SENSITIVITY_MEDIUM int = 100
const SENSITIVITY_SHORT int = 10

const LARGE_NUMBER_STEPS int = 64
const MEDIUM_NUMBER_STEPS int = 32

const STROBE_SPEED_FAST int = 255
const STROBE_SPEED_MEDIUM int = 127
const STROBE_SPEED_SLOW int = 0

// newMiniSequencer is a simple sequencer which can be attached to a switch and a single fixture to allow simple effects.
// The miniSequenceer implements the actions attaced to a switch state.
// Currently we support 1. Off 2. Control, ability to set programs 3. Static colors 4. Chase. soft, hard and timed or music triggered.
// Long term objective of actions is to replace the direct value settings.
func newMiniSequencer(fixture *Fixture,
	swiTch common.Switch,
	override common.Override,
	action Action,
	dmxController *ft232.DMXController, fixturesConfig *Fixtures,
	switchChannels []common.SwitchChannel, soundConfig *sound.SoundConfig,
	blackout bool, brightness int, master int, masterChanging bool, lastColor common.LastColor,
	dmxInterfacePresent bool,
	eventsForLaunchpad chan common.ALight,
	guiButtons chan common.ALight,
	fixtureStepChannel chan common.FixtureCommand) {

	if debug_mini {
		fmt.Printf("newMiniSequencer: actions are available fixture %s actionName %s\n", fixture.Name, action.Name)
	}

	switchName := fmt.Sprintf("switch%d", swiTch.Number)

	mySequenceNumber := fixture.Group - 1
	myFixtureNumber := fixture.Number - 1

	// Setup the configuration.
	cfg := GetConfig(action, fixture, fixturesConfig)

	if debug_mini {
		fmt.Printf("Action %+v\n", action)
	}

	if action.Mode == "Off" {
		if debug_mini {
			fmt.Printf("Stop mini sequence for switch number %d\n", swiTch.Number)
		}

		// Remember that we have stopped this mini sequencer.
		setSwitchState(&swiTch, false, blackout, master)

		// Disable this mini sequencer with the sound service.
		if soundConfig.GetSoundTriggerState(switchName) {
			// Use the switch name as the unique sequence name.
			err := soundConfig.DisableSoundTrigger(switchName)
			if err != nil {
				fmt.Printf("Error while trying to disable sound trigger %s\n", err.Error())
				os.Exit(1)
			}
			if debug_mini {
				fmt.Printf("Sound trigger %s disabled\n", switchName)
			}
		}

		// Stop any running fade ups.
		select {
		case switchChannels[swiTch.Number].StopFadeUp <- true:
			MapFixtures(false, false, mySequenceNumber, myFixtureNumber, colors.Black, colors.Black, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, false, 0, 0, 0, false, 0, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		// Stop any running fade downs.
		select {
		case switchChannels[swiTch.Number].StopFadeDown <- true:
			MapFixtures(false, false, mySequenceNumber, myFixtureNumber, colors.Black, colors.Black, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, false, 0, 0, 0, false, 0, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		// Stop any running chases.
		select {
		case switchChannels[swiTch.Number].Stop <- true:
			MapFixtures(false, false, mySequenceNumber, myFixtureNumber, colors.Black, colors.Black, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, false, 0, 0, 0, false, 0, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		// Stop any rotates.
		select {
		case switchChannels[swiTch.Number].StopRotate <- true:
			MapFixtures(false, false, mySequenceNumber, myFixtureNumber, colors.Black, colors.Black, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, false, 0, 0, 0, false, 0, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		if lastColor.RGBColor != colors.EmptyColor {
			if debug {
				fmt.Printf("Action OFF fade to black\n")
			}
			fadeDownValues := common.GetFadeValues(64, float64(master), 1, true)
			for _, fade := range fadeDownValues {
				// Listen for stop command.
				select {
				case <-switchChannels[swiTch.Number].StopFadeDown:
					return
				case <-time.After(10 * time.Millisecond):
				}
				common.LightLamp(common.Button{X: swiTch.Number - 1, Y: 3}, lastColor.RGBColor, fade, eventsForLaunchpad, guiButtons)
				MapFixtures(false, false, mySequenceNumber, myFixtureNumber, lastColor.RGBColor, lastColor.RGBColor, 0, 0, 0, cfg.RotateSpeed, cfg.Program, 0, 0, fixturesConfig, blackout, brightness, fade, cfg.Music, cfg.Strobe, cfg.StrobeSpeed, dmxController, dmxInterfacePresent)
				// Control how long the fade take with the fade control.
				time.Sleep((5 * time.Millisecond) * (time.Duration(cfg.Fade)))
			}
			state := swiTch.States[0]
			buttonColor, _ := common.GetRGBColorByName(state.ButtonColor)
			common.LightLamp(common.Button{X: swiTch.Number - 1, Y: 3}, buttonColor, master, eventsForLaunchpad, guiButtons)
		} else {
			lastColor = MapFixtures(false, false, mySequenceNumber, myFixtureNumber, colors.Black, colors.Black, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, false, 0, 0, 0, false, 0, dmxController, dmxInterfacePresent)
		}

		return
	}

	if action.Mode == "Control" {
		if debug_mini {
			fmt.Printf("Control selected for switch number %d\n", swiTch.Number)
		}

		// Remember that we have stopped this mini sequencer.
		setSwitchState(&swiTch, false, blackout, master)

		// Disable this mini sequencer with the sound service.
		// Use the switch name as the unique sequence name.
		err := soundConfig.DisableSoundTrigger(switchName)
		if err != nil {
			fmt.Printf("Error while trying to disable sound trigger %s\n", err.Error())
			os.Exit(1)
		}
		if debug {
			fmt.Printf("Sound trigger %s disabled\n", switchName)
		}

		// Stop any running fades.
		select {
		case switchChannels[swiTch.Number].StopFadeUp <- true:
			MapFixtures(false, false, mySequenceNumber, myFixtureNumber, colors.Black, colors.Black, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, false, 0, 0, 0, false, 0, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		// Stop any running fade downs.
		select {
		case switchChannels[swiTch.Number].StopFadeDown <- true:
			MapFixtures(false, false, mySequenceNumber, myFixtureNumber, colors.Black, colors.Black, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, false, 0, 0, 0, false, 0, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		// Stop any running chases.
		select {
		case switchChannels[swiTch.Number].Stop <- true:
			MapFixtures(false, false, mySequenceNumber, myFixtureNumber, colors.Black, colors.Black, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, false, 0, 0, 0, false, 0, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		// Stop any rotates.
		select {
		case switchChannels[swiTch.Number].StopRotate <- true:
			MapFixtures(false, false, mySequenceNumber, myFixtureNumber, colors.Black, colors.Black, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, false, 0, 0, 0, false, 0, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		if fixtureHasChannel(fixture, "Master") {
			// Find the program speed channel for this fixture.
			masterChannel, err := GetChannelNumberByName(fixture, "Master")
			if err != nil {
				fmt.Printf("warning: Switch Number %d: %s\n", swiTch.Number, err)
			}
			if debug {
				fmt.Printf("fixture %s: Control: send master Address %d Value %d \n", fixture.Name, fixture.Address+int16(masterChannel), master)
			}
			SetChannel(fixture.Address+int16(masterChannel), byte(master), dmxController, dmxInterfacePresent)
		}

		if fixtureHasChannel(fixture, "Shutter") {
			// Find the program speed channel for this fixture.
			shutterChannel, err := GetChannelNumberByName(fixture, "Shutter")
			if err != nil {
				fmt.Printf("warning: Switch Number %d: %s\n", swiTch.Number, err)
			}
			if debug_mini {
				fmt.Printf("fixture %s: Control: send Shutter Address %d Value %d Strobe %t\n", fixture.Name, fixture.Address+int16(shutterChannel), master, override.Strobe)
			}

			if override.Strobe {
				strobeValues := GetADMXValueMaxMin(fixture, "Strobe", "Shutter")
				shutter := makeStrobeSpeed(strobeValues, override.StrobeSpeed)
				if debug_mini {
					fmt.Printf("Override is set Address=%d Strobe Shutter=%t DMX Value=%d\n", fixture.Address+int16(shutterChannel), override.Shutter, shutter)
				}
				SetChannel(fixture.Address+int16(shutterChannel), byte(shutter), dmxController, dmxInterfacePresent)
			} else {
				shutter := GetADMXValueByName(fixture, "Open", "Shutter")
				if debug_mini {
					fmt.Printf("Override is set Address=%d Open Shutter=%t DMX Value=%d\n", fixture.Address+int16(shutterChannel), override.Shutter, shutter)
				}
				SetChannel(fixture.Address+int16(shutterChannel), byte(shutter), dmxController, dmxInterfacePresent)
			}
		}

		if fixtureHasChannel(fixture, "Rotate") {
			// Find the rotate channel for this fixture.
			rotateChannel, err := GetChannelNumberByName(fixture, "Rotate")
			if err != nil {
				fmt.Printf("warning: Switch Number %d: %s\n", swiTch.Number, err)
			}

			var rotate int = 0
			if debug {
				fmt.Printf("fixture %s: Control: send Rotate Address %d Value %d \n", fixture.Name, fixture.Address+int16(rotateChannel), rotate)
			}
			if override.Gobo != 0 {
				rotate = GetADMXValue(fixture, override.Rotate, "Rotate")
				if debug_mini {
					fmt.Printf("Override is set so Rotate is %d DMX Vaue %d\n", override.Rotate, rotate)
				}
			}
			SetChannel(fixture.Address+int16(rotateChannel), byte(rotate), dmxController, dmxInterfacePresent)
		}

		if fixtureHasChannel(fixture, "Color") {
			// Find the color channel for this fixture.
			colorChannel, err := GetChannelNumberByName(fixture, "Color")
			if err != nil {
				fmt.Printf("warning: Switch Number %d: %s\n", swiTch.Number, err)
			}

			var color int = 0
			if debug {
				fmt.Printf("fixture %s: Control: send Color Address %d Value %d \n", fixture.Name, fixture.Address+int16(colorChannel), color)
			}
			if override.Color != 0 {
				color = GetADMXValue(fixture, override.Color, "Color")
				if debug_mini {
					fmt.Printf("Override is set so Color is %d DMX Vaue %d\n", override.Gobo, color)
				}
			}
			SetChannel(fixture.Address+int16(colorChannel), byte(color), dmxController, dmxInterfacePresent)
		}

		if fixtureHasChannel(fixture, "Gobo") {
			// Find the gobo channel for this fixture.
			goboChannel, err := GetChannelNumberByName(fixture, "Gobo")
			if err != nil {
				fmt.Printf("warning: Switch Number %d: %s\n", swiTch.Number, err)
			}
			if debug {
				fmt.Printf("fixture %s: Control: send Gobo Address %d Value %d \n", fixture.Name, fixture.Address+int16(goboChannel), master)
			}
			var gobo int = 0
			if override.Gobo != 0 {
				gobo = GetADMXValue(fixture, override.Gobo, "Gobo")
				if debug_mini {
					fmt.Printf("Override is set so Gobo is %d DMX Vaue %d\n", override.Gobo, gobo)
				}
			}
			SetChannel(fixture.Address+int16(goboChannel), byte(gobo), dmxController, dmxInterfacePresent)
		}

		if fixtureHasChannel(fixture, "ProgramSpeed") {
			// Find the program speed channel for this fixture.
			programSpeedChannel, err := GetChannelNumberByName(fixture, "ProgramSpeed")
			if err != nil {
				fmt.Printf("warning: Switch Number %d: %s\n", swiTch.Number, err)
			}
			programSpeed, err := GetChannelSettingByNameAndSetting(fixture, "ProgramSpeed", action.ProgramSpeed)
			if err != nil {
				fmt.Printf("warning: %s\n", err)
			}
			if debug {
				fmt.Printf("fixture %s: Control: send ProgramSpeed Address %d Value %d \n", fixture.Name, fixture.Address+int16(programSpeedChannel), master)
			}
			if override.ProgramSpeed != 0 {
				programSpeed = GetADMXValue(fixture, override.ProgramSpeed, "ProgramSpeed")
				if debug_mini {
					fmt.Printf("Override is set so ProgramSpeed is %d DMX Vaue %d\n", override.ProgramSpeed, programSpeed)
				}
			}
			// Now play that DMX value on the program channel of this fixture.
			SetChannel(fixture.Address+int16(programSpeedChannel), byte(programSpeed), dmxController, dmxInterfacePresent)
		}

		if fixtureHasChannel(fixture, "Program") {
			// Find the program channel for this fixture.
			programChannel, err := GetChannelNumberByName(fixture, "Program")
			if err != nil {
				fmt.Printf("warning: Switch Number %d: %s\n", swiTch.Number, err)
			}
			// Look up the program state required.
			programState, err := GetChannelSettingByNameAndSetting(fixture, "Program", action.Program)
			if err != nil {
				fmt.Printf("warning: %s\n", err)
			}
			if debug {
				fmt.Printf("fixture %s: Control: send Program Address %d Value %d \n", fixture.Name, fixture.Address+int16(programState), master)
			}
			if blackout {
				SetChannel(fixture.Address+int16(programChannel), 0, dmxController, dmxInterfacePresent)
			} else {
				SetChannel(fixture.Address+int16(programChannel), byte(programState), dmxController, dmxInterfacePresent)
			}
		}

		return
	}

	if action.Mode == "Static" {
		if debug_mini {
			fmt.Printf("Static mini sequence for switch number %d\n", swiTch.Number)
		}

		// Remember that we have stopped this mini sequencer.
		setSwitchState(&swiTch, false, blackout, master)

		// Disable this mini sequencer with the sound service.
		// Use the switch name as the unique sequence name.
		err := soundConfig.DisableSoundTrigger(switchName)
		if err != nil {
			fmt.Printf("Error while trying to disable sound trigger %s\n", err.Error())
			os.Exit(1)
		}
		if debug {
			fmt.Printf("Sound trigger %s disable\n", switchName)
		}

		// Stop any running fades.
		select {
		case switchChannels[swiTch.Number].StopFadeUp <- true:
			lastColor = MapFixtures(false, false, mySequenceNumber, myFixtureNumber, colors.Black, colors.Black, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, false, 0, 0, 0, false, 0, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		// Stop any running fade downs.
		select {
		case switchChannels[swiTch.Number].StopFadeDown <- true:
			lastColor = MapFixtures(false, false, mySequenceNumber, myFixtureNumber, colors.Black, colors.Black, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, false, 0, 0, 0, false, 0, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		// Stop any running chases.
		select {
		case switchChannels[swiTch.Number].Stop <- true:
			lastColor = MapFixtures(false, false, mySequenceNumber, myFixtureNumber, colors.Black, colors.Black, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, false, 0, 0, 0, false, 0, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		// Stop any rotates.
		select {
		case switchChannels[swiTch.Number].StopRotate <- true:
			lastColor = MapFixtures(false, false, mySequenceNumber, myFixtureNumber, colors.Black, colors.Black, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, false, 0, 0, 0, false, 0, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		// Decide on the shutter value.
		if fixtureHasChannel(fixture, "Shutter") {
			if override.Strobe {
				strobeValues := GetADMXValueMaxMin(fixture, "Strobe", "Shutter")
				cfg.Shutter = makeStrobeSpeed(strobeValues, override.StrobeSpeed)
				if debug_mini {
					fmt.Printf("Override is set so Shutter is %t DMX Vaue %d\n", override.Shutter, cfg.Shutter)
				}
			} else {
				shutter := GetADMXValueByName(fixture, "Open", "Shutter")
				if debug_mini {
					fmt.Printf("Override is set Open Shutter=%t DMX Value=%d\n", override.Shutter, shutter)
				}
			}
		}

		// Decide on the static color.
		var color color.RGBA
		if override.Color > 0 {
			colorName := override.AvailableColorNames[override.Color]
			color, err = common.GetRGBColorByName(colorName)
			if err != nil {
				fmt.Printf("error %d\n", err)
			}
			if debug_mini {
				fmt.Printf("Override is set so Color name is %s index is %d Color %+v Available Colors %+v\n", colorName, override.Color, color, override.AvailableColorNames)
			}
		} else {
			// Use the fitst static color from the action.
			color, err = common.GetRGBColorByName(action.Colors[0])
			if err != nil {
				fmt.Printf("error %d\n", err)
			}
		}

		// Decide on the static gobo.
		if fixtureHasChannel(fixture, "Gobo") {
			if override.Gobo > 0 {
				cfg.Gobo = GetADMXValue(fixture, override.Gobo, "Gobo")
				//if debug_mini {
				fmt.Printf("Override is set so Gobo is %d DMX Vaue %d\n", override.Rotate, cfg.Gobo)
				//}
			}
		}

		// Decide on the static rotate speed.
		if fixtureHasChannel(fixture, "Rotate") {
			if override.Rotate != 0 {
				cfg.RotateSpeed = GetADMXValue(fixture, override.Rotate, "Rotate")
				//if debug_mini {
				fmt.Printf("Override is set so Rotate is %d DMX Vaue %d\n", override.Rotate, cfg.RotateSpeed)
				//}
			}
		}

		// Soft start
		// Calulate the steps
		fadeUpValues := common.GetFadeValues(64, float64(common.MAX_DMX_BRIGHTNESS), 1, false)
		fadeDownValues := common.GetFadeValues(64, float64(common.MAX_DMX_BRIGHTNESS), 1, true)

		// Now Fade up
		go func(lastColor common.LastColor) {

			if !masterChanging {
				// If last color is set then fade down first.
				if lastColor.RGBColor != colors.EmptyColor {
					for _, fade := range fadeDownValues {
						// Listen for stop command.
						select {
						case <-switchChannels[swiTch.Number].StopFadeDown:
							return
						case <-time.After(10 * time.Millisecond):
						}
						common.LightLamp(common.Button{X: swiTch.Number - 1, Y: 3}, lastColor.RGBColor, fade, eventsForLaunchpad, guiButtons)
						MapFixtures(false, false, mySequenceNumber, myFixtureNumber, lastColor.RGBColor, lastColor.RGBColor, cfg.Pan, cfg.Tilt, cfg.Shutter, cfg.RotateSpeed, cfg.Program, 0, 0, fixturesConfig, blackout, fade, master, cfg.Music, cfg.Strobe, cfg.StrobeSpeed, dmxController, dmxInterfacePresent)
						// Control how long the fade take with the fade control.
						time.Sleep((5 * time.Millisecond) * (time.Duration(common.Reverse(cfg.Fade))))
					}
					// Fade down complete, set lastColor to empty in the fixture.
					command := common.FixtureCommand{
						Type:      "lastColor",
						LastColor: colors.EmptyColor,
					}
					select {
					case fixtureStepChannel <- command:
					case <-time.After(100 * time.Millisecond):
					}
				}
				for _, fade := range fadeUpValues {
					// Listen for stop command.
					select {
					case <-switchChannels[swiTch.Number].StopFadeUp:
						return
					case <-time.After(10 * time.Millisecond):
					}
					common.LightLamp(common.Button{X: swiTch.Number - 1, Y: 3}, color, fade, eventsForLaunchpad, guiButtons)
					MapFixtures(false, false, mySequenceNumber, myFixtureNumber, color, color, cfg.Pan, cfg.Tilt, cfg.Shutter, cfg.RotateSpeed, cfg.Program, cfg.Gobo, cfg.ScannerColor, fixturesConfig, blackout, fade, master, cfg.Music, cfg.Strobe, cfg.StrobeSpeed, dmxController, dmxInterfacePresent)
					// Control how long the fade take with the fade control.
					time.Sleep((5 * time.Millisecond) * (time.Duration(common.Reverse(cfg.Fade))))
				}
				// Fade up complete, set lastColor up in the fixture.
				command := common.FixtureCommand{
					Type:      "lastColor",
					LastColor: color,
				}
				select {
				case fixtureStepChannel <- command:
				case <-time.After(100 * time.Millisecond):
				}
			} else {
				common.LightLamp(common.Button{X: swiTch.Number - 1, Y: 3}, color, master, eventsForLaunchpad, guiButtons)
				MapFixtures(false, false, mySequenceNumber, myFixtureNumber, color, color, cfg.Pan, cfg.Tilt, cfg.Shutter, cfg.RotateSpeed, cfg.Program, cfg.Gobo, cfg.ScannerColor, fixturesConfig, blackout, brightness, master, cfg.Music, cfg.Strobe, cfg.StrobeSpeed, dmxController, dmxInterfacePresent)
			}
		}(lastColor)
		return
	}

	if action.Mode == "Chase" {

		if debug_mini {
			fmt.Printf("Chase mini sequence for switch number %d\n", swiTch.Number)
		}

		// Stop any running fade ups.
		select {
		case switchChannels[swiTch.Number].StopFadeUp <- true:
			lastColor = MapFixtures(false, false, mySequenceNumber, myFixtureNumber, colors.Black, colors.Black, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, false, 0, 0, 0, false, 0, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		// Stop any running fade downs.
		select {
		case switchChannels[swiTch.Number].StopFadeDown <- true:
			lastColor = MapFixtures(false, false, mySequenceNumber, myFixtureNumber, colors.Black, colors.Black, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, false, 0, 0, 0, false, 0, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		// Don't stop this mini sequencer if there's one running already.
		// Unless we are changing switch positions.
		if getSwitchState(swiTch) { //&& SwitchPosition == swiTch.CurrentPosition {
			setSwitchState(&swiTch, true, blackout, master)
			return
		}

		// If the last color isn't empty, fade down last color before starting chase.
		if lastColor.RGBColor != colors.EmptyColor {
			if debug {
				fmt.Printf("Action Chase STARTUP: fade down to black from %+v\n", lastColor)
			}
			fadeDownValues := common.GetFadeValues(64, float64(master), 1, true)
			for _, fade := range fadeDownValues {
				// Listen for stop command.
				select {
				case <-switchChannels[swiTch.Number].StopFadeDown:
					return
				case <-time.After(10 * time.Millisecond):
				}
				common.LightLamp(common.Button{X: swiTch.Number - 1, Y: 3}, lastColor.RGBColor, fade, eventsForLaunchpad, guiButtons)
				MapFixtures(false, false, mySequenceNumber, myFixtureNumber, lastColor.RGBColor, lastColor.RGBColor, 0, 0, 0, cfg.RotateSpeed, cfg.Program, 0, 0, fixturesConfig, blackout, brightness, fade, cfg.Music, cfg.Strobe, cfg.StrobeSpeed, dmxController, dmxInterfacePresent)
				// Control how long the fade take with the fade control.
				time.Sleep((5 * time.Millisecond) * (time.Duration(cfg.Fade)))
			}
			state := swiTch.States[0]
			buttonColor, _ := common.GetRGBColorByName(state.ButtonColor)
			common.LightLamp(common.Button{X: swiTch.Number - 1, Y: 3}, buttonColor, master, eventsForLaunchpad, guiButtons)
		}

		// Remember that we have started this mini sequencer.
		setSwitchState(&swiTch, true, blackout, master)

		// DeRegister this mini sequencer with the sound service.
		// Use the switch name as the unique sequence name.
		err := soundConfig.DisableSoundTrigger(switchName)
		if err != nil {
			fmt.Printf("Error while trying to disable sound trigger %s\n", err.Error())
			os.Exit(1)
		}
		if debug {
			fmt.Printf("Sound trigger %s disable\n", switchName)
		}

		// Turn off the fixture.
		select {
		case switchChannels[swiTch.Number].Stop <- true:
			lastColor = MapFixtures(false, false, mySequenceNumber, myFixtureNumber, colors.Black, colors.Black, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, false, 0, 0, 0, false, 0, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		// Register this mini sequencer with the sound service.
		// Use the switch name as the unique sequence name.
		if cfg.MusicTrigger {
			err := soundConfig.EnableSoundTrigger(switchName)
			if err != nil {
				fmt.Printf("Error while trying to enable sound trigger %s\n", err.Error())
				os.Exit(1)
			}
			if debug_mini {
				fmt.Printf("Sound trigger %s enabled\n", switchName)
			}
		}

		// Stop any left over sequence left over for this switch.
		select {
		case switchChannels[swiTch.Number].Stop <- true:
			lastColor = MapFixtures(false, false, mySequenceNumber, myFixtureNumber, colors.Black, colors.Black, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, false, 0, 0, 0, false, 0, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		// Create a sequence and calculate steps.
		sequence, RGBPositions, numberSteps := createSequence(cfg)

		var rotateCounter int
		var goboChangeFrequency int
		var goboCounter int

		// Main chaser thread.
		go func() {

			if cfg.Rotatable {

				if debug_rotate {
					fmt.Printf("Fixture %s is rotatable\n", fixture.Name)
				}

				rotateChannel, err := GetChannelNumberByName(fixture, "Rotate")
				if err != nil {
					fmt.Printf("rotator: %s,", err)
				}
				masterChannel, err := GetChannelNumberByName(fixture, "Master")
				if err != nil {
					fmt.Printf("master: %s,", err)
					return
				}

				if debug_rotate {
					fmt.Printf("rotateChannel %d masterChannel %d cfg.RotateSpeed %d\n", rotateChannel, masterChannel, cfg.RotateSpeed)
					fmt.Printf("cfg.ForwardSpeed %d, cfg.ReverseSpeed %d \n", cfg.ForwardSpeed, cfg.ReverseSpeed)
				}

				// Consume any left over stop commands before starting.
				select {
				case <-switchChannels[swiTch.Number].StopRotate:
				case <-time.After(10 * time.Millisecond):
				}

				// Thread to run the rotator
				go func(switchNumber int) {

					for {

						select {
						case <-switchChannels[swiTch.Number].StopRotate:
							time.Sleep(1 * time.Millisecond)
							SetChannel(fixture.Address+int16(rotateChannel), byte(0), dmxController, dmxInterfacePresent)
							return
						case <-switchChannels[swiTch.Number].KeepRotateAlive:
							time.Sleep(1 * time.Millisecond)
							continue
						case <-time.After(1500 * time.Millisecond):
							SetChannel(fixture.Address+int16(rotateChannel), byte(0), dmxController, dmxInterfacePresent)
							time.Sleep(250 * time.Millisecond)
							SetChannel(fixture.Address+int16(masterChannel), byte(0), dmxController, dmxInterfacePresent)
						}
					}
				}(swiTch.Number)
			}

			// Wait for rotator thread to start.
			time.Sleep(100 * time.Millisecond)

			for {
				// Apply the overrides.
				if !cfg.MusicTrigger && override.Speed != 0 {
					if debug_mini {
						fmt.Printf("Override is set so Speed is %d\n", override.Speed)
					}
					cfg.SpeedDuration = common.SetSpeed(override.Speed)
				}

				if override.Shutter || override.Strobe {
					if override.Strobe {
						strobeValues := GetADMXValueMaxMin(fixture, "Strobe", "Shutter")
						shutter := makeStrobeSpeed(strobeValues, override.StrobeSpeed)
						if debug_mini {
							fmt.Printf("Override is set Strobe Shutter=%t DMX Value=%d\n", override.Shutter, shutter)
						}
						cfg.Shutter = shutter
					} else {
						shutter := GetADMXValueByName(fixture, "Open", "Shutter")
						if debug_mini {
							fmt.Printf("Override is set Open Shutter=%t DMX Value=%d\n", override.Shutter, shutter)
						}
						cfg.Shutter = shutter
					}
				}

				if override.ProgramSpeed != 0 {
					if debug_mini {
						fmt.Printf("Override is set so Speed is %d\n", override.Speed)
					}
					cfg.ProgramSpeed = GetADMXValue(fixture, override.ProgramSpeed, "ProgramSpeed")
				}

				if override.Shift != 0 {
					if debug_mini {
						fmt.Printf("Override is set so Shift is %d\n", override.Shift)
					}
					cfg.Shift = override.Shift
				}
				if override.Size != 0 {
					if debug_mini {
						fmt.Printf("Override is set so Size is %d\n", override.Size)
					}
					cfg.Size = override.Size
				}
				if override.Fade != 0 {
					if debug_mini {
						fmt.Printf("Override is set so Fade is %d\n", override.Fade)
					}
					cfg.Fade = override.Fade
				}

				if override.Rotate != 0 {
					if debug_rotate {
						fmt.Printf("Override is set so Rotate Speed is %d\n", override.Rotate)
					}
					// At this point we need to convert a 1-10 rotate value that means something to this specific fixture.
					// cfg.RotateSpeed is the DMX value from the Rotate channel settings.
					// override.RotateSpeed is the index so for example 1 is setting 1.
					cfg.RotateSpeed = GetADMXValue(fixture, override.Rotate, "Rotate")
					if debug_override {
						fmt.Printf("Apply Rotate Speed DMX Value=%d\n", cfg.RotateSpeed)
					}
				}

				if override.Color != 0 {
					if debug_mini {
						fmt.Printf("Override is set so Colors is %+v\n", override.AvailableColors)
					}
					// You can only override a single color.
					cfg.AvailableColors[cfg.Color] = override.AvailableColors[cfg.Color]
					newColor, err := common.GetRGBColorByName(override.AvailableColorNames[cfg.Color-1])
					if err != nil {
						fmt.Printf("error %d\n", err)
					}
					cfg.AvailableColors = []color.RGBA{
						newColor,
					}
				}

				if override.Gobo != 0 {
					if debug_mini {
						fmt.Printf("Override is set so Gobo is %+v\n", override.Gobo)
					}
					cfg.Gobo = override.Gobo
				}

				if debug_mini {
					fmt.Printf("Speed %d Duration %d Shift %d numberSteps %d\n", cfg.Speed, cfg.SpeedDuration, cfg.Shift, numberSteps)
				}

				// Run through the steps in the sequence.
				// Remember every step contains infomation for all the fixtures in this group.
				for step := 0; step < numberSteps; step++ {

					if cfg.Rotatable {
						select {
						case switchChannels[swiTch.Number].KeepRotateAlive <- true:
						case <-time.After(10 * time.Millisecond):
						}

						if rotateCounter > cfg.RotateSensitivity {
							rotateCounter = 1
						}

						if cfg.AutoRotate {
							if debug_rotate {
								fmt.Printf("AutoRotate is on\n")
							}
							if rotateCounter < (cfg.RotateSensitivity / 2) {
								// Clockwise Speed.
								cfg.RotateSpeed = cfg.ForwardSpeed
							} else {
								// Anti Clockwise Speed.
								cfg.RotateSpeed = cfg.ReverseSpeed
							}
						}
					}

					if cfg.AutoGobo {
						if goboCounter >= len(cfg.GoboOptions) {
							goboCounter = 0
						}
						if goboChangeFrequency > cfg.GoboSpeed {
							goboChangeFrequency = 0
							goboCounter++
						}
						goboChangeFrequency++
						cfg.Gobo = goboCounter
					}

					if debug_mini {
						fmt.Printf("Rotate Value %d Counter %d  Clockwise %d Anti %d \n", cfg.RotateSpeed, rotateCounter, cfg.ForwardSpeed, cfg.ReverseSpeed)
						fmt.Printf("Gobo %d goboChangeFrequency %d\n", cfg.Gobo, goboChangeFrequency)
						fmt.Printf("switch:%d waiting for beat on %d with speed %d\n", swiTch.Number, swiTch.Number+10, cfg.SpeedDuration)
						fmt.Printf("switch:%d speed %d\n", swiTch.Number, cfg.SpeedDuration)
					}

					// This is were we wait for a beat or a time out equivalent to the speed.
					select {
					// First five triggers are occupied by sequence 0-FOH,1-Upluighters,2-Scanners,3-Switches,4-ShutterChaser
					// So switch channels use 5 -12
					case cmd := <-switchChannels[swiTch.Number].CommandChannel:
						if debug_override {
							fmt.Printf("CMD is %+v\n", cmd)
						}
						// Update RGB Speed or Scanner Shutter Speed but not in music trigger mode.
						if !cfg.MusicTrigger && cmd.Action == common.UpdateSpeed {
							const SPEED = 0
							override.Speed = cmd.Args[SPEED].Value.(int)
							cfg.SpeedDuration = common.SetSpeed(override.Speed)
							if debug_override {
								fmt.Printf("Speed %d Duration %d\n", cmd.Args[SPEED].Value.(int), cfg.SpeedDuration)
							}
						}

						// Update Shutter when strobe speed changes.
						if cmd.Action == common.UpdateStrobeSpeed {
							const STROBE = 0
							const STROBE_SPEED = 1
							override.Strobe = cmd.Args[STROBE].Value.(bool)
							override.Shutter = cmd.Args[STROBE].Value.(bool)

							cfg.Strobe = cmd.Args[STROBE].Value.(bool)
							cfg.StrobeSpeed = cmd.Args[STROBE_SPEED].Value.(int)
							if debug_override {
								fmt.Printf("Strobe Speed %d\n", cmd.Args[STROBE_SPEED].Value.(int))
							}
						}

						// Update Shift.
						if cmd.Action == common.UpdateRGBShift {
							const SHIFT = 0
							override.Shift = cmd.Args[SHIFT].Value.(int)
							cfg.Shift = cmd.Args[SHIFT].Value.(int)
							if debug_override {
								fmt.Printf("Shift %d\n", cmd.Args[SHIFT].Value.(int))
							}
						}
						// Update Size.
						if cmd.Action == common.UpdateRGBSize {
							const SIZE = 0
							override.Size = cmd.Args[SIZE].Value.(int)
							cfg.Size = common.GetSize(cmd.Args[SIZE].Value.(int))
							if debug_override {
								fmt.Printf("Size %d\n", cmd.Args[SIZE].Value.(int))
							}
						}
						// Update Fade
						if cmd.Action == common.UpdateRGBFadeSpeed {
							const FADE = 0
							override.Fade = cmd.Args[FADE].Value.(int)
							cfg.Fade = cmd.Args[FADE].Value.(int)
							if debug_override {
								fmt.Printf("Fade %d\n", cmd.Args[FADE].Value.(int))
							}
						}

						// Update Rotate Speed
						if cmd.Action == common.UpdateRotateSpeed {
							const ROTATE_SPEED = 0
							override.Rotate = cmd.Args[ROTATE_SPEED].Value.(int)
							if debug_override || debug_rotate {
								fmt.Printf("Override Speed Index=%d\n", override.Rotate)
							}
							// At this point we need to convert a 1-10 rotate value that means something to this specific fixture.
							// cfg.RotateSpeed is the DMX value from the Rotate channel settings.
							// override.RotateSpeed is the index so for example 1 is setting 1.
							cfg.RotateSpeed = GetADMXValue(fixture, override.Rotate, "Rotate")
							if debug_override {
								fmt.Printf("Rotate Speed DMX Value=%d\n", cfg.RotateSpeed)
							}
						}

						// Update Colors
						if cmd.Action == common.UpdateColors {
							const COLORS = 0
							override.Color = cmd.Args[COLORS].Value.(int)
							newColor, err := common.GetRGBColorByName(override.AvailableColorNames[cfg.Color-1])
							if err != nil {
								fmt.Printf("error %d\n", err)
							}
							cfg.AvailableColors = []color.RGBA{
								newColor,
							}
							if debug_override {
								fmt.Printf("Color Selected is %d %+v\n", override.Color-1, newColor)
							}
						}

						// Update Gobos
						if cmd.Action == common.UpdateGobo {
							const GOBO = 0
							override.Gobo = cmd.Args[GOBO].Value.(int)
							cfg.Gobo = cmd.Args[GOBO].Value.(int)
							if debug_override {
								fmt.Printf("Gobo %d\n", cmd.Args[GOBO].Value.(int))
							}
						}

						// Recreate the sequence and recalculate steps.
						sequence, RGBPositions, numberSteps = createSequence(cfg)

					case <-soundConfig.SoundTriggers[swiTch.Number+4].Channel:
					case <-switchChannels[swiTch.Number].Stop:
						// Stop.
						if cfg.Rotatable {
							switchChannels[swiTch.Number].StopRotate <- true
						}
						// And turn the fixture off.
						MapFixtures(false, false, mySequenceNumber, myFixtureNumber, colors.Black, colors.Black, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, blackout, brightness, master, cfg.Music, cfg.Strobe, cfg.StrobeSpeed, dmxController, dmxInterfacePresent)
						return
					case <-time.After(cfg.SpeedDuration / 50):
					}

					// Play out fixture to DMX channels.
					position := RGBPositions[step]

					fixtures := position.Fixtures

					var actualMaster int
					for fixtureNumber := 0; fixtureNumber < sequence.NumberFixtures; fixtureNumber++ {
						thisFixture := fixtures[fixtureNumber]
						common.LightLamp(common.Button{X: swiTch.Number - 1, Y: 3}, thisFixture.Color, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)
						if cfg.Map {
							// Use sound triggered brighness and apply master
							actualMaster = int((float64(thisFixture.Brightness) / 100) * (float64(master) / 2.55))
						} else {
							actualMaster = master
						}
						MapFixtures(false, false, mySequenceNumber, myFixtureNumber, thisFixture.Color, thisFixture.BaseColor, cfg.Pan, cfg.Tilt, cfg.Shutter, cfg.RotateSpeed, 0, cfg.Gobo, 0, fixturesConfig, blackout, brightness, actualMaster, cfg.Music, cfg.Strobe, cfg.StrobeSpeed, dmxController, dmxInterfacePresent)
					}

					rotateCounter++
				}
			}
		}()
	}
}

func createSequence(cfg ActionConfig) (common.Sequence, map[int]common.Position, int) {

	if debug_mini {
		fmt.Printf("createSequence cfg.Shift %d cfg.NumberSteps %d cfg.Fade %d  cfg.Size %d cfg.AvailableColors %+v\n", cfg.Shift, cfg.NumberSteps, cfg.Fade, cfg.Size, cfg.AvailableColors)
	}

	sequence := common.Sequence{
		ScannerReverse:       false,
		RGBInvert:            false,
		Bounce:               false,
		ScannerChaser:        true,
		RGBShift:             cfg.Shift,
		RGBNumberStepsInFade: cfg.NumberSteps,
		RGBFade:              cfg.Fade,
		RGBSize:              cfg.Size,
	}
	sequence.Pattern = pattern.MakeSingleFixtureChase(cfg.AvailableColors)
	steps := sequence.Pattern.Steps
	sequence.NumberFixtures = 4
	// Calculate fade curve values.
	common.CalculateFadeValues(&sequence)
	// Calulate positions for each RGB fixture.
	sequence.Optimisation = false
	sequence.FixtureState = map[int]common.FixtureState{
		0: {
			Enabled: true,
		},
		1: {
			Enabled: true,
		},
		2: {
			Enabled: true,
		},
		3: {
			Enabled: true,
		},
	}
	fadeColors, totalNumberOfSteps := position.CalculatePositions(steps, sequence, common.IS_RGB)
	RGBPositions, numberSteps := position.AssemblePositions(fadeColors, sequence.NumberFixtures, totalNumberOfSteps, sequence.FixtureState, sequence.Optimisation)

	return sequence, RGBPositions, numberSteps
}

func GetChannelSettinsByName(fixture *Fixture, channelName string, fixturesConfig *Fixtures) ([]common.Setting, error) {
	if debug_mini {
		fmt.Printf("GetChannelSettinsByName: Looking for program settings for fixture %s\n", fixture.Name)
	}

	settingNames := []common.Setting{}

	// Find the channel by name.
	for _, channel := range fixture.Channels {
		if debug_mini {
			fmt.Printf("GetChannelSettinsByName: looking at channel %s\n", channel.Name)
		}
		if channel.Name == channelName {
			if debug_mini {
				fmt.Printf("Found a Program Channel\n")
			}
			// First look for any settings available for this channel.
			if channel.Settings != nil {
				for _, setting := range channel.Settings {
					if debug_mini {
						fmt.Printf("Looking through Settings %s\n", setting.Name)
					}
					v, _ := strconv.Atoi(setting.Value)
					value := common.Setting{
						Name:  setting.Name,
						Value: int16(v),
					}
					settingNames = append(settingNames, value)
				}
				return settingNames, nil
			}
			// If the program has a hard coded value return that as a default.
			if channel.Value != nil {
				if debug_mini {
					fmt.Printf("Found a Default Program Value of %d\n", *channel.Value)
				}
				value := common.Setting{
					Name:  "Default",
					Value: *channel.Value,
				}
				settingNames = append(settingNames, value)
				return settingNames, nil
			}
		}
	}

	return nil, fmt.Errorf("failed to find settings for channel %s in fixture%s", channelName, fixture.Name)
}

// getSwitchState reports on if this switch is running a mini sequence.
func getSwitchState(swiTch common.Switch) bool {
	return swiTch.MiniSequencerRunning
}

func setSwitchState(swiTch *common.Switch, state bool, blackout bool, master int) {
	swiTch.MiniSequencerRunning = state
	swiTch.Blackout = blackout
	swiTch.Master = master
}
