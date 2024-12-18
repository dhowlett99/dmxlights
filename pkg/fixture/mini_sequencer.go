// Copyright (C) 2022,2023 dhowlett99.
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
	"os"
	"strconv"
	"time"

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/pattern"
	"github.com/dhowlett99/dmxlights/pkg/position"
	"github.com/dhowlett99/dmxlights/pkg/sound"
	"github.com/oliread/usbdmx/ft232"
)

const debug_mini bool = false

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
func newMiniSequencer(fixture *Fixture, swiTch common.Switch, action Action,
	dmxController *ft232.DMXController, fixturesConfig *Fixtures,
	switchChannels []common.SwitchChannel, soundConfig *sound.SoundConfig,
	blackout bool, brightness int, master int, masterChanging bool, lastColor common.LastColor,
	dmxInterfacePresent bool,
	eventsForLaunchpad chan common.ALight,
	guiButtons chan common.ALight,
	fixtureStepChannel chan common.FixtureCommand) {

	if debug {
		fmt.Printf("newMiniSequencer: start fixture %s\n", fixture.Name)
	}

	switchName := fmt.Sprintf("switch%d", swiTch.Number)

	mySequenceNumber := fixture.Group - 1
	myFixtureNumber := fixture.Number - 1

	// Setup the configuration.
	cfg := getConfig(action, fixture, fixturesConfig)

	if debug_mini {
		fmt.Printf("Action %+v\n", action)
	}

	if action.Mode == "Off" {
		if debug_mini {
			fmt.Printf("Stop mini sequence for switch number %d\n", swiTch.Number)
		}

		// Remember that we have stopped this mini sequencer.
		setSwitchState(swiTch, false, blackout, master)

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
			MapFixtures(false, false, mySequenceNumber, myFixtureNumber, common.Black, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, false, 0, 0, 0, false, 0, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		// Stop any running fade downs.
		select {
		case switchChannels[swiTch.Number].StopFadeDown <- true:
			MapFixtures(false, false, mySequenceNumber, myFixtureNumber, common.Black, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, false, 0, 0, 0, false, 0, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		// Stop any running chases.
		select {
		case switchChannels[swiTch.Number].Stop <- true:
			MapFixtures(false, false, mySequenceNumber, myFixtureNumber, common.Black, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, false, 0, 0, 0, false, 0, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		// Stop any rotates.
		select {
		case switchChannels[swiTch.Number].StopRotate <- true:
			MapFixtures(false, false, mySequenceNumber, myFixtureNumber, common.Black, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, false, 0, 0, 0, false, 0, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		if lastColor.RGBColor != common.EmptyColor {
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
				MapFixtures(false, false, mySequenceNumber, myFixtureNumber, lastColor.RGBColor, 0, 0, 0, cfg.RotateSpeed, cfg.Program, 0, 0, fixturesConfig, blackout, brightness, fade, cfg.Music, cfg.Strobe, cfg.StrobeSpeed, dmxController, dmxInterfacePresent)
				// Control how long the fade take with the speed control.
				time.Sleep((5 * time.Millisecond) * (time.Duration(cfg.Fade)))
			}
			state := swiTch.States[0]
			buttonColor, _ := common.GetRGBColorByName(state.ButtonColor)
			common.LightLamp(common.Button{X: swiTch.Number - 1, Y: 3}, buttonColor, master, eventsForLaunchpad, guiButtons)
		} else {
			lastColor = MapFixtures(false, false, mySequenceNumber, myFixtureNumber, common.Black, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, false, 0, 0, 0, false, 0, dmxController, dmxInterfacePresent)
		}

		return
	}

	if action.Mode == "Control" {
		if debug_mini {
			fmt.Printf("Control selected for switch number %d\n", swiTch.Number)
		}

		// Remember that we have stopped this mini sequencer.
		setSwitchState(swiTch, false, blackout, master)

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
			MapFixtures(false, false, mySequenceNumber, myFixtureNumber, common.Black, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, false, 0, 0, 0, false, 0, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		// Stop any running fade downs.
		select {
		case switchChannels[swiTch.Number].StopFadeDown <- true:
			MapFixtures(false, false, mySequenceNumber, myFixtureNumber, common.Black, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, false, 0, 0, 0, false, 0, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		// Stop any running chases.
		select {
		case switchChannels[swiTch.Number].Stop <- true:
			MapFixtures(false, false, mySequenceNumber, myFixtureNumber, common.Black, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, false, 0, 0, 0, false, 0, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		// Stop any rotates.
		select {
		case switchChannels[swiTch.Number].StopRotate <- true:
			MapFixtures(false, false, mySequenceNumber, myFixtureNumber, common.Black, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, false, 0, 0, 0, false, 0, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		//MapFixtures(false, false, mySequenceNumber, myFixtureNumber, common.Black, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, false, 0, 0, 0, false, 0, dmxController, dmxInterfacePresent)

		// Find the program channel for this fixture.
		programChannel, err := FindChannelNumberByName(fixture, "Program")
		if err != nil {
			fmt.Printf("warning: Switch Number %d: %s\n", swiTch.Number, err)
		}

		if fixtureHasChannel(fixture, "Master") {
			// Find the program speed channel for this fixture.
			masterChannel, err := FindChannelNumberByName(fixture, "Master")
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
			shutterChannel, err := FindChannelNumberByName(fixture, "Shutter")
			if err != nil {
				fmt.Printf("warning: Switch Number %d: %s\n", swiTch.Number, err)
			}
			if debug {
				fmt.Printf("fixture %s: Control: send Shutter Address %d Value %d \n", fixture.Name, fixture.Address+int16(shutterChannel), master)
			}
			SetChannel(fixture.Address+int16(shutterChannel), byte(32), dmxController, dmxInterfacePresent)
		}

		if fixtureHasChannel(fixture, "Rotate") {
			// Find the rotate channel for this fixture.
			rotateChannel, err := FindChannelNumberByName(fixture, "Rotate")
			if err != nil {
				fmt.Printf("warning: Switch Number %d: %s\n", swiTch.Number, err)
			}
			if debug {
				fmt.Printf("fixture %s: Control: send Rotate Address %d Value %d \n", fixture.Name, fixture.Address+int16(rotateChannel), master)
			}
			SetChannel(fixture.Address+int16(rotateChannel), byte(0), dmxController, dmxInterfacePresent)
		}

		if fixtureHasChannel(fixture, "Gobo") {
			// Find the gobo channel for this fixture.
			goboChannel, err := FindChannelNumberByName(fixture, "Gobo")
			if err != nil {
				fmt.Printf("warning: Switch Number %d: %s\n", swiTch.Number, err)
			}
			if debug {
				fmt.Printf("fixture %s: Control: send Gobo Address %d Value %d \n", fixture.Name, fixture.Address+int16(goboChannel), master)
			}
			SetChannel(fixture.Address+int16(goboChannel), byte(0), dmxController, dmxInterfacePresent)
		}
		if fixtureHasChannel(fixture, "ProgramSpeed") {
			// Find the program speed channel for this fixture.
			programSpeedChannel, err := FindChannelNumberByName(fixture, "ProgramSpeed")
			if err != nil {
				fmt.Printf("warning: Switch Number %d: %s\n", swiTch.Number, err)
			}
			if debug {
				fmt.Printf("fixture %s: Control: send ProgramSpeed Address %d Value %d \n", fixture.Name, fixture.Address+int16(programSpeedChannel), master)
			}
			// Now play that DMX value on the program channel of this fixture.
			SetChannel(fixture.Address+int16(programSpeedChannel), byte(cfg.ProgramSpeed), dmxController, dmxInterfacePresent)
		}

		if fixtureHasChannel(fixture, "Program") {
			// Look up the program state required.
			programState, err := findChannelSettingByChannelNameAndSettingName(fixture, "Program", action.Program)
			if err != nil {
				fmt.Printf("warning: %s\n", err)
			}
			if debug {
				fmt.Printf("fixture %s: Control: send Program Address %d Value %d \n", fixture.Name, fixture.Address+int16(programState), master)
			}
			SetChannel(fixture.Address+int16(programChannel), byte(programState), dmxController, dmxInterfacePresent)
		}

		return
	}

	if action.Mode == "Static" {
		if debug_mini {
			fmt.Printf("Static mini sequence for switch number %d\n", swiTch.Number)
		}

		// Remember that we have stopped this mini sequencer.
		setSwitchState(swiTch, false, blackout, master)

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
			lastColor = MapFixtures(false, false, mySequenceNumber, myFixtureNumber, common.Black, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, false, 0, 0, 0, false, 0, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		// Stop any running fade downs.
		select {
		case switchChannels[swiTch.Number].StopFadeDown <- true:
			lastColor = MapFixtures(false, false, mySequenceNumber, myFixtureNumber, common.Black, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, false, 0, 0, 0, false, 0, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		// Stop any running chases.
		select {
		case switchChannels[swiTch.Number].Stop <- true:
			lastColor = MapFixtures(false, false, mySequenceNumber, myFixtureNumber, common.Black, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, false, 0, 0, 0, false, 0, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		// Stop any rotates.
		select {
		case switchChannels[swiTch.Number].StopRotate <- true:
			lastColor = MapFixtures(false, false, mySequenceNumber, myFixtureNumber, common.Black, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, false, 0, 0, 0, false, 0, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		color, err := common.GetRGBColorByName(action.Colors[0])
		if err != nil {
			fmt.Printf("error %d\n", err)
		}

		// Soft start
		// Calulate the steps
		fadeUpValues := common.GetFadeValues(64, float64(common.MAX_DMX_BRIGHTNESS), 1, false)
		fadeDownValues := common.GetFadeValues(64, float64(common.MAX_DMX_BRIGHTNESS), 1, true)

		// Now Fade up
		go func(lastColor common.LastColor) {

			if !masterChanging {
				// If last color is set then fade down first.
				if lastColor.RGBColor != common.EmptyColor {
					for _, fade := range fadeDownValues {
						// Listen for stop command.
						select {
						case <-switchChannels[swiTch.Number].StopFadeDown:
							return
						case <-time.After(10 * time.Millisecond):
						}
						common.LightLamp(common.Button{X: swiTch.Number - 1, Y: 3}, lastColor.RGBColor, fade, eventsForLaunchpad, guiButtons)
						MapFixtures(false, false, mySequenceNumber, myFixtureNumber, lastColor.RGBColor, 0, 0, 0, cfg.RotateSpeed, cfg.Program, 0, 0, fixturesConfig, blackout, fade, master, cfg.Music, cfg.Strobe, cfg.StrobeSpeed, dmxController, dmxInterfacePresent)
						// Control how long the fade take with the fade speed control.
						time.Sleep((5 * time.Millisecond) * (time.Duration(common.Reverse(cfg.Fade))))
					}
					// Fade down complete, set lastColor to empty in the fixture.
					command := common.FixtureCommand{
						Type:      "lastColor",
						LastColor: common.EmptyColor,
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
					MapFixtures(false, false, mySequenceNumber, myFixtureNumber, color, 0, 0, 0, cfg.RotateSpeed, cfg.Program, 0, 0, fixturesConfig, blackout, fade, master, cfg.Music, cfg.Strobe, cfg.StrobeSpeed, dmxController, dmxInterfacePresent)
					// Control how long the fade take with the fade speed control.
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
				MapFixtures(false, false, mySequenceNumber, myFixtureNumber, color, 0, 0, 0, cfg.RotateSpeed, cfg.Program, 0, 0, fixturesConfig, blackout, brightness, master, cfg.Music, cfg.Strobe, cfg.StrobeSpeed, dmxController, dmxInterfacePresent)
			}
		}(lastColor)
		return
	}

	if action.Mode == "Chase" {

		if debug_mini {
			fmt.Printf("Chase mini sequence for switch number %d\n", swiTch.Number)
		}

		// Stop any running fades.
		select {
		case switchChannels[swiTch.Number].StopFadeUp <- true:
			lastColor = MapFixtures(false, false, mySequenceNumber, myFixtureNumber, common.Black, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, false, 0, 0, 0, false, 0, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		// Stop any running fade downs.
		select {
		case switchChannels[swiTch.Number].StopFadeDown <- true:
			lastColor = MapFixtures(false, false, mySequenceNumber, myFixtureNumber, common.Black, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, false, 0, 0, 0, false, 0, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		// Don't stop this mini sequencer if there's one running already.
		// Unless we are changing switch positions.
		if getSwitchState(swiTch) { //&& SwitchPosition == swiTch.CurrentPosition {
			setSwitchState(swiTch, true, blackout, master)
			return
		}

		// If the last color isn't empty, fade down last color before starting chase.
		if lastColor.RGBColor != common.EmptyColor {
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
				MapFixtures(false, false, mySequenceNumber, myFixtureNumber, lastColor.RGBColor, 0, 0, 0, cfg.RotateSpeed, cfg.Program, 0, 0, fixturesConfig, blackout, brightness, fade, cfg.Music, cfg.Strobe, cfg.StrobeSpeed, dmxController, dmxInterfacePresent)
				// Control how long the fade take with the speed control.
				time.Sleep((5 * time.Millisecond) * (time.Duration(cfg.Fade)))
			}
			state := swiTch.States[0]
			buttonColor, _ := common.GetRGBColorByName(state.ButtonColor)
			common.LightLamp(common.Button{X: swiTch.Number - 1, Y: 3}, buttonColor, master, eventsForLaunchpad, guiButtons)
		}

		// Remember that we have started this mini sequencer.
		setSwitchState(swiTch, true, blackout, master)

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
			lastColor = MapFixtures(false, false, mySequenceNumber, myFixtureNumber, common.Black, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, false, 0, 0, 0, false, 0, dmxController, dmxInterfacePresent)
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
			lastColor = MapFixtures(false, false, mySequenceNumber, myFixtureNumber, common.Black, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, false, 0, 0, 0, false, 0, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		sequence := common.Sequence{
			ScannerReverse:       false,
			RGBInvert:            false,
			Bounce:               false,
			ScannerChaser:        true,
			RGBShift:             1,
			RGBNumberStepsInFade: cfg.NumberSteps,
			RGBFade:              cfg.Fade,
			RGBSize:              cfg.Size,
		}
		sequence.Pattern = pattern.MakeSingleFixtureChase(cfg.Colors)
		steps := sequence.Pattern.Steps
		sequence.NumberFixtures = 1
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
			4: {
				Enabled: true,
			},
		}
		fadeColors, totalNumberOfSteps := position.CalculatePositions(steps, sequence, common.IS_RGB)
		RGBPositions, numberSteps := position.AssemblePositions(fadeColors, sequence.NumberFixtures, totalNumberOfSteps, sequence.FixtureState, sequence.Optimisation)

		var rotateCounter int
		var goboChangeFrequency int
		var goboCounter int
		var clockwiseSpeed int
		var antiClockwiseSpeed int

		// Calculate the rotation speed based on direction and soeed.
		if cfg.Rotatable {
			antiClockwiseSpeed, err = findChannelSettingByNameAndSpeed(fixture.Name, "Rotate", "Rotate Anti Clockwise", action.RotateSpeed, fixturesConfig)
			if err != nil {
				fmt.Printf("Fixture %s anti clockwise rotate speed: %s\n", fixture.Name, err)
			}
			clockwiseSpeed, err = findChannelSettingByNameAndSpeed(fixture.Name, "Rotate", "Rotate Clockwise", action.RotateSpeed, fixturesConfig)
			if err != nil {
				fmt.Printf("Fixture %s clockwise rotate speed: %s\n", fixture.Name, err)
			}
			if debug {
				fmt.Printf("clockwiseSpeed %d\n", clockwiseSpeed)
				fmt.Printf("antiClockwiseSpeed %d\n", antiClockwiseSpeed)
			}
		}

		go func() {

			if cfg.Rotatable {

				rotateChannel, err := FindChannelNumberByName(fixture, "Rotate")
				if err != nil {
					fmt.Printf("rotator: %s,", err)
				}
				masterChannel, err := FindChannelNumberByName(fixture, "Master")
				if err != nil {
					fmt.Printf("master: %s,", err)
					return
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
						if !cfg.Clockwise && !cfg.AntiClockwise {
							cfg.RotateSpeed = 0
						}
						if cfg.Clockwise {
							cfg.RotateSpeed = clockwiseSpeed
						}
						if cfg.AntiClockwise {
							cfg.RotateSpeed = antiClockwiseSpeed
						}

						if cfg.AutoRotate {
							if rotateCounter < (cfg.RotateSensitivity / 2) {
								// Clockwise Speed.
								cfg.RotateSpeed = clockwiseSpeed
							} else {
								// Anti Clockwise Speed.
								cfg.RotateSpeed = antiClockwiseSpeed
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
						fmt.Printf("Rotate Value %d Counter %d  Clockwise %d Anti %d \n", cfg.RotateSpeed, rotateCounter, clockwiseSpeed, antiClockwiseSpeed)
						fmt.Printf("Gobo %d goboChangeFrequency %d\n", cfg.Gobo, goboChangeFrequency)
						fmt.Printf("switch:%d waiting for beat on %d with speed %d\n", swiTch.Number, swiTch.Number+10, cfg.Speed)
						fmt.Printf("switch:%d speed %d\n", swiTch.Number, cfg.Speed)
					}

					// This is were we wait for a beat or a time out equivalent to the speed.
					select {
					// First five triggers are occupied by sequence 0-FOH,1-Upluighters,2-Scanners,3-Switches,4-ShutterChaser
					// So switch channels use 5 -12
					case <-soundConfig.SoundTriggers[swiTch.Number+4].Channel:
					case <-switchChannels[swiTch.Number].Stop:
						// Stop.
						if cfg.Rotatable {
							switchChannels[swiTch.Number].StopRotate <- true
						}
						// And turn the fixture off.
						MapFixtures(false, false, mySequenceNumber, myFixtureNumber, common.Black, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, blackout, brightness, master, cfg.Music, cfg.Strobe, cfg.StrobeSpeed, dmxController, dmxInterfacePresent)
						return
					case <-time.After(cfg.Speed):
					}

					// Play out fixture to DMX channels.
					position := RGBPositions[step]

					fixtures := position.Fixtures

					var actualMaster int
					for fixtureNumber := 0; fixtureNumber < sequence.NumberFixtures; fixtureNumber++ {
						fixture := fixtures[fixtureNumber]
						common.LightLamp(common.Button{X: swiTch.Number - 1, Y: 3}, fixture.Color, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)
						if cfg.Map {
							// Use sound triggered brighness and apply master
							actualMaster = int((float64(fixture.Brightness) / 100) * (float64(master) / 2.55))
						} else {
							actualMaster = master
						}
						MapFixtures(false, false, mySequenceNumber, myFixtureNumber, fixture.Color, 0, 0, 0, cfg.RotateSpeed, 0, cfg.Gobo, 0, fixturesConfig, blackout, brightness, actualMaster, cfg.Music, cfg.Strobe, cfg.StrobeSpeed, dmxController, dmxInterfacePresent)
					}

					rotateCounter++
				}
			}
		}()
	}
}

func getConfig(action Action, fixture *Fixture, fixturesConfig *Fixtures) ActionConfig {

	config := ActionConfig{}

	var programSettings []common.Setting
	var goboSettings []common.Setting
	var err error

	fixtureInfo := FindFixtureInfo(fixture)
	if debug {
		fmt.Printf("This fixture has Rotate Feature %+v\n", fixtureInfo)
	}

	// Find all the specified settings for the program channel
	if fixtureInfo.HasProgram {

		programSettings, err = GetChannelSettinsByName(fixture, "Program", fixturesConfig)
		if err != nil && debug {
			fmt.Printf("newMiniSequencer: warning! no program settings found for fixture %s\n", fixture.Name)
		}

		// Program Speed - Speed of programs or shows.
		switch action.ProgramSpeed {
		case "Off":
			config.ProgramSpeed = 0
		case "Slow":
			config.ProgramSpeed = 10
		case "Medium":
			config.ProgramSpeed = 100
		case "Fast":
			config.ProgramSpeed = 200
		default:
			config.ProgramSpeed = 0
		}

		// Look through the available settins and see if you can find the specified program action.
		for _, setting := range programSettings {
			if action.Program == setting.Name || setting.Name == "Default" {
				config.Program = int(setting.Value)
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
		case "Default":
			config.Gobo = 0
			config.AutoGobo = false
		case "Auto":
			config.Gobo = -1
			config.AutoGobo = true
		default:
			// find current gobo number.
			config.Gobo = FindGobo(fixture.Number-1, fixture.Group-1, action.Gobo, fixturesConfig)
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

	if action.Colors != nil {
		// Find the color by name from the library of supported colors.
		colorLibrary, err := common.GetColorArrayByNames(action.Colors)
		if err != nil {
			fmt.Printf("error: %s\n", err.Error())
		}
		config.Colors = colorLibrary
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

	switch action.Rotate {
	case "Off":
		config.Rotatable = false
		config.AutoRotate = false
		config.Clockwise = false
		config.AntiClockwise = false
	case "Clockwise":
		config.Rotatable = true
		config.AutoRotate = false
		config.Clockwise = true
		config.AntiClockwise = false
	case "Anti Clockwise":
		config.Rotatable = true
		config.AutoRotate = false
		config.Clockwise = false
		config.AntiClockwise = true
	case "Auto":
		config.Rotatable = true
		config.AutoRotate = true
		config.Clockwise = false
		config.AntiClockwise = false
	default:
		config.Rotatable = false
		config.AutoRotate = false
		config.Clockwise = false
		config.AntiClockwise = false
	}

	switch action.Speed {
	case "Slow":
		config.TriggerState = false
		config.Speed = 1 * time.Second
		config.MusicTrigger = false
		config.NumberSteps = LARGE_NUMBER_STEPS
		config.RotateSensitivity = SENSITIVITY_SHORT
	case "Medium":
		config.TriggerState = false
		config.Speed = 500 * time.Millisecond
		config.MusicTrigger = false
		config.NumberSteps = LARGE_NUMBER_STEPS
		config.RotateSensitivity = SENSITIVITY_SHORT
	case "Fast":
		config.TriggerState = false
		config.Speed = 250 * time.Millisecond
		config.MusicTrigger = false
		config.NumberSteps = LARGE_NUMBER_STEPS
		config.RotateSensitivity = SENSITIVITY_SHORT
	case "VeryFast":
		config.TriggerState = false
		config.Speed = 50 * time.Millisecond
		config.MusicTrigger = false
		config.NumberSteps = LARGE_NUMBER_STEPS
		config.RotateSensitivity = SENSITIVITY_SHORT
	case "Music":
		config.TriggerState = true
		config.Speed = time.Duration(12 * time.Hour)
		config.MusicTrigger = true
		config.NumberSteps = MEDIUM_NUMBER_STEPS
		config.RotateSensitivity = SENSITIVITY_LONG
	default:
		config.TriggerState = false
		config.Speed = time.Duration(12 * time.Hour)
		config.MusicTrigger = false
		config.NumberSteps = MEDIUM_NUMBER_STEPS
		config.RotateSensitivity = SENSITIVITY_SHORT
	}

	switch action.Strobe {
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

	return config
}

func GetChannelSettinsByName(fixture *Fixture, channelName string, fixtures *Fixtures) ([]common.Setting, error) {
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

func setSwitchState(swiTch common.Switch, state bool, blackout bool, master int) {
	swiTch.MiniSequencerRunning = state
	swiTch.Blackout = blackout
	swiTch.Master = master
}
