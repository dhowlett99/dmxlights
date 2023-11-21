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

// newMiniSequencer is a simple sequencer which can be attached to a switch and a single fixture to allow simple effects.
// The miniSequenceer implements the actions attaced to a switch state.
// Currently we support 1. Off 2. Control, ability to set programs 3. Static colors 4. Chase. soft, hard and timed or music triggered.
// Long term objective of actions is to replace the direct value settings.
func newMiniSequencer(fixture *Fixture, swiTch common.Switch, action Action,
	dmxController *ft232.DMXController, fixturesConfig *Fixtures,
	switchChannels []common.SwitchChannel, soundConfig *sound.SoundConfig,
	blackout bool, brightness int, master int, masterChanging bool, lastColor common.Color,
	dmxInterfacePresent bool,
	eventsForLaunchpad chan common.ALight,
	guiButtons chan common.ALight,
	fixtureStepChannel chan common.FixtureCommand) {

	switchName := fmt.Sprintf("switch%d", swiTch.Number)

	mySequenceNumber := fixture.Group - 1
	myFixtureNumber := fixture.Number - 1

	// Find all the specified settings for the program channel
	programSettings, err := GetChannelSettinsByName(fixture, "Program", fixturesConfig)
	if err != nil && debug {
		fmt.Printf("newMiniSequencer: warning! no program settings found for fixture %s\n", fixture.Name)
	}

	cfg := getConfig(action, programSettings)

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
			turnOffFixture(myFixtureNumber, mySequenceNumber, lastColor, fixturesConfig, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		// Stop any running fade downs.
		select {
		case switchChannels[swiTch.Number].StopFadeDown <- true:
			turnOffFixture(myFixtureNumber, mySequenceNumber, lastColor, fixturesConfig, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		// Stop any running chases.
		select {
		case switchChannels[swiTch.Number].Stop <- true:
			turnOffFixture(myFixtureNumber, mySequenceNumber, lastColor, fixturesConfig, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		// Stop any rotates.
		select {
		case switchChannels[swiTch.Number].StopRotate <- true:
			turnOffFixture(myFixtureNumber, mySequenceNumber, lastColor, fixturesConfig, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		if lastColor != common.EmptyColor {
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
				common.LightLamp(common.ALight{X: swiTch.Number - 1, Y: 3, Brightness: fade, Red: lastColor.R, Green: lastColor.G, Blue: lastColor.B}, eventsForLaunchpad, guiButtons)
				MapFixtures(false, false, mySequenceNumber, dmxController, myFixtureNumber, lastColor.R, lastColor.G, lastColor.B, 0, 0, 0, 0, 0, 0, cfg.RotateSpeed, cfg.Music, cfg.Program, 0, 0, fixturesConfig, blackout, brightness, fade, cfg.Strobe, cfg.StrobeSpeed, dmxInterfacePresent)
				// Control how long the fade take with the speed control.
				time.Sleep((5 * time.Millisecond) * (time.Duration(cfg.FadeSpeed)))
			}
			state := swiTch.States[0]
			buttonColor, _ := common.GetRGBColorByName(state.ButtonColor)
			common.LightLamp(common.ALight{X: swiTch.Number - 1, Y: 3, Brightness: master, Red: buttonColor.R, Green: buttonColor.G, Blue: buttonColor.B}, eventsForLaunchpad, guiButtons)
		} else {
			turnOffFixture(myFixtureNumber, mySequenceNumber, lastColor, fixturesConfig, dmxController, dmxInterfacePresent)
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
			turnOffFixture(myFixtureNumber, mySequenceNumber, lastColor, fixturesConfig, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		// Stop any running fade downs.
		select {
		case switchChannels[swiTch.Number].StopFadeDown <- true:
			turnOffFixture(myFixtureNumber, mySequenceNumber, lastColor, fixturesConfig, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		// Stop any running chases.
		select {
		case switchChannels[swiTch.Number].Stop <- true:
			turnOffFixture(myFixtureNumber, mySequenceNumber, lastColor, fixturesConfig, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		// Stop any rotates.
		select {
		case switchChannels[swiTch.Number].StopRotate <- true:
			turnOffFixture(myFixtureNumber, mySequenceNumber, lastColor, fixturesConfig, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		turnOffFixture(myFixtureNumber, mySequenceNumber, lastColor, fixturesConfig, dmxController, dmxInterfacePresent)

		// Find the program channel for this fixture.
		programChannel, err := FindChannelNumberByName(fixture, "Program")
		if err != nil {
			fmt.Printf("fixture %s program channel not found: %s,", fixture.Name, err)
			return
		}

		// Look up the program state required.
		v, err := findChannelSettingByChannelNameAndSettingName(fixture, "Program", action.Program)
		if err != nil {
			fmt.Printf("fixture %s program state not found: %s,", fixture.Name, err)
			return
		}

		// Now play that DMX value on the program channel of this fixture.
		SetChannel(fixture.Address+int16(programChannel), byte(v), dmxController, dmxInterfacePresent)

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
			turnOffFixture(myFixtureNumber, mySequenceNumber, lastColor, fixturesConfig, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		// Stop any running fade downs.
		select {
		case switchChannels[swiTch.Number].StopFadeDown <- true:
			turnOffFixture(myFixtureNumber, mySequenceNumber, lastColor, fixturesConfig, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		// Stop any running chases.
		select {
		case switchChannels[swiTch.Number].Stop <- true:
			turnOffFixture(myFixtureNumber, mySequenceNumber, lastColor, fixturesConfig, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		// Stop any rotates.
		select {
		case switchChannels[swiTch.Number].StopRotate <- true:
			turnOffFixture(myFixtureNumber, mySequenceNumber, lastColor, fixturesConfig, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		color, err := common.GetRGBColorByName(action.Colors[0])
		if err != nil {
			fmt.Printf("error %d\n", err)
		}

		// Soft start
		// Calulate the steps
		fadeUpValues := common.GetFadeValues(64, float64(master), 1, false)
		fadeDownValues := common.GetFadeValues(64, float64(master), 1, true)

		// Now Fade up
		go func(lastColor common.Color) {

			if !masterChanging {
				// If last color is set then fade down first.
				if lastColor != common.EmptyColor {
					for _, fade := range fadeDownValues {
						// Listen for stop command.
						select {
						case <-switchChannels[swiTch.Number].StopFadeDown:
							return
						case <-time.After(10 * time.Millisecond):
						}
						common.LightLamp(common.ALight{X: swiTch.Number - 1, Y: 3, Brightness: fade, Red: lastColor.R, Green: lastColor.G, Blue: lastColor.B}, eventsForLaunchpad, guiButtons)
						MapFixtures(false, false, mySequenceNumber, dmxController, myFixtureNumber, lastColor.R, lastColor.G, lastColor.B, 0, 0, 0, 0, 0, 0, cfg.RotateSpeed, cfg.Music, cfg.Program, 0, 0, fixturesConfig, blackout, brightness, fade, cfg.Strobe, cfg.StrobeSpeed, dmxInterfacePresent)
						// Control how long the fade take with the fade speed control.
						time.Sleep((5 * time.Millisecond) * (time.Duration(common.Reverse(cfg.FadeSpeed))))
					}
				}
				for _, fade := range fadeUpValues {
					// Listen for stop command.
					select {
					case <-switchChannels[swiTch.Number].StopFadeUp:
						return
					case <-time.After(10 * time.Millisecond):
					}
					common.LightLamp(common.ALight{X: swiTch.Number - 1, Y: 3, Brightness: fade, Red: color.R, Green: color.G, Blue: color.B}, eventsForLaunchpad, guiButtons)
					MapFixtures(false, false, mySequenceNumber, dmxController, myFixtureNumber, color.R, color.G, color.B, 0, 0, 0, 0, 0, 0, cfg.RotateSpeed, cfg.Music, cfg.Program, 0, 0, fixturesConfig, blackout, brightness, fade, cfg.Strobe, cfg.StrobeSpeed, dmxInterfacePresent)
					// Control how long the fade take with the fade speed control.
					time.Sleep((5 * time.Millisecond) * (time.Duration(common.Reverse(cfg.FadeSpeed))))
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
				common.LightLamp(common.ALight{X: swiTch.Number - 1, Y: 3, Brightness: master, Red: color.R, Green: color.G, Blue: color.B}, eventsForLaunchpad, guiButtons)
				MapFixtures(false, false, mySequenceNumber, dmxController, myFixtureNumber, color.R, color.G, color.B, 0, 0, 0, 0, 0, 0, cfg.RotateSpeed, cfg.Music, cfg.Program, 0, 0, fixturesConfig, blackout, brightness, master, cfg.Strobe, cfg.StrobeSpeed, dmxInterfacePresent)
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
			turnOffFixture(myFixtureNumber, mySequenceNumber, lastColor, fixturesConfig, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		// Stop any running fade downs.
		select {
		case switchChannels[swiTch.Number].StopFadeDown <- true:
			turnOffFixture(myFixtureNumber, mySequenceNumber, lastColor, fixturesConfig, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		// Don't stop this mini sequencer if there's one running already.
		// Unless we are changing switch positions.
		if getSwitchState(swiTch, swiTch.Number) { //&& SwitchPosition == swiTch.CurrentPosition {
			setSwitchState(swiTch, true, blackout, master)
			return
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
			turnOffFixture(myFixtureNumber, mySequenceNumber, lastColor, fixturesConfig, dmxController, dmxInterfacePresent)
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
			turnOffFixture(myFixtureNumber, mySequenceNumber, lastColor, fixturesConfig, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		sequence := common.Sequence{
			ScannerReverse:       false,
			RGBInvert:            false,
			Bounce:               false,
			ScannerChaser:        false,
			RGBShift:             1,
			RGBNumberStepsInFade: 32,
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
		fadeColors, numberFixtures, totalNumberOfSteps := position.CalculatePositions(steps, sequence, common.IS_RGB)
		RGBPositions, numberSteps := position.AssemblePositions(fadeColors, numberFixtures, totalNumberOfSteps, sequence.FixtureState, sequence.Optimisation)

		var rotateCounter int
		var clockwiseSpeed int
		var antiClockwiseSpeed int

		// Calculate the rotation speed based on direction and soeed.
		if cfg.Rotatable {
			antiClockwiseSpeed, err = findChannelSettingByNameAndSpeed(fixture.Name, "Rotate", "Anti Clockwise", action.RotateSpeed, fixturesConfig)
			if err != nil {
				fmt.Printf("anti clockwise rotate speed: %s\n", err)
			}
			clockwiseSpeed, err = findChannelSettingByNameAndSpeed(fixture.Name, "Rotate", "Clockwise", action.RotateSpeed, fixturesConfig)
			if err != nil {
				fmt.Printf("clockwise rotate speed: %s\n", err)
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

						if rotateCounter > 500 {
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

						if cfg.Auto {
							if rotateCounter < 250 {
								// Clockwise Speed.
								cfg.RotateSpeed = clockwiseSpeed
							} else {
								// Anti Clockwise Speed.
								cfg.RotateSpeed = antiClockwiseSpeed
							}
						}
					}

					if debug_mini {
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
						MapFixtures(false, false, mySequenceNumber, dmxController, myFixtureNumber, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, blackout, brightness, master, cfg.Strobe, cfg.StrobeSpeed, dmxInterfacePresent)
						return
					case <-time.After(cfg.Speed):
					}

					// Play out fixture to DMX channels.
					position := RGBPositions[step]

					fixtures := position.Fixtures

					for fixtureNumber := 0; fixtureNumber < sequence.NumberFixtures; fixtureNumber++ {
						fixture := fixtures[fixtureNumber]
						common.LightLamp(common.ALight{X: swiTch.Number - 1, Y: 3, Brightness: 255, Red: fixture.Color.R, Green: fixture.Color.G, Blue: fixture.Color.B}, eventsForLaunchpad, guiButtons)
						MapFixtures(false, false, mySequenceNumber, dmxController, myFixtureNumber, fixture.Color.R, fixture.Color.G, fixture.Color.B, fixture.Color.W, 0, 0, 0, 0, 0, cfg.RotateSpeed, 0, 0, 0, 0, fixturesConfig, blackout, brightness, master, cfg.Strobe, cfg.StrobeSpeed, dmxInterfacePresent)
					}

					rotateCounter++
				}
			}
		}()
	}
}

func getConfig(action Action, programSettings []common.Setting) ActionConfig {

	config := ActionConfig{}

	if action.Colors != nil {
		// Find the color by name from the library of supported colors.
		colorLibrary, err := common.GetColorArrayByNames(action.Colors)
		if err != nil {
			fmt.Printf("error: %s\n", err.Error())
		}
		config.Colors = colorLibrary
	}

	// Fade - Time taken to fade up and down.
	switch action.Fade {
	case "Off":
		config.Fade = 1
	case "Soft":
		config.Fade = 1
	case "Sharp":
		config.Fade = 10
	default:
		config.Fade = 1
	}

	// Size - How long does the lamp stay on.
	switch action.Size {
	case "Off":
		config.Size = 0
	case "Short":
		config.Size = 1
	case "Medium":
		config.Size = 3
	case "Long":
		config.Size = 10
	default:
		config.Size = 3
	}

	// Look through the available settins and see if you can find the specified program action.
	for _, setting := range programSettings {
		if action.Program == setting.Name || setting.Name == "Default" {
			config.Program = int(setting.Value)
		}
	}

	switch action.Rotate {
	case "Off":
		config.Rotatable = false
		config.Auto = false
		config.Clockwise = false
		config.AntiClockwise = false
	case "Clockwise":
		config.Rotatable = true
		config.Auto = false
		config.Clockwise = true
		config.AntiClockwise = false
	case "Anti Clockwise":
		config.Rotatable = true
		config.Auto = false
		config.Clockwise = false
		config.AntiClockwise = true
	case "Auto":
		config.Rotatable = true
		config.Auto = true
		config.Clockwise = false
		config.AntiClockwise = false
	default:
		config.Rotatable = false
		config.Auto = false
		config.Clockwise = false
		config.AntiClockwise = false
	}

	switch action.Speed {
	case "Slow":
		config.TriggerState = false
		config.Speed = 1 * time.Second
		config.MusicTrigger = false
	case "Medium":
		config.TriggerState = false
		config.Speed = 500 * time.Millisecond
		config.MusicTrigger = false
	case "Fast":
		config.TriggerState = false
		config.Speed = 250 * time.Millisecond
		config.MusicTrigger = false
	case "VeryFast":
		config.TriggerState = false
		config.Speed = 50 * time.Millisecond
		config.MusicTrigger = false
	case "Music":
		config.TriggerState = true
		config.Speed = time.Duration(12 * time.Hour)
		config.MusicTrigger = true
	default:
		config.TriggerState = false
		config.Speed = time.Duration(12 * time.Hour)
		config.MusicTrigger = false
	}

	switch action.Strobe {
	case "Off":
		config.Strobe = false
		config.StrobeSpeed = 0
	case "Slow":
		config.Strobe = true
		config.StrobeSpeed = 0
	case "Medium":
		config.Strobe = true
		config.StrobeSpeed = 127
	case "Fast":
		config.Strobe = true
		config.StrobeSpeed = 255
	default:
		config.Strobe = false
		config.StrobeSpeed = 0
	}

	config.FadeSpeed = action.FadeSpeed

	return config
}

func GetChannelSettinsByName(fixture *Fixture, name string, fixtures *Fixtures) ([]common.Setting, error) {
	if debug_mini {
		fmt.Printf("GetChannelSettinsByName: Looking for program settings for fixture %s\n", fixture.Name)
	}

	settingNames := []common.Setting{}

	// Find the channel by name.
	for _, channel := range fixture.Channels {
		if debug_mini {
			fmt.Printf("GetChannelSettinsByName: looking at channel %s\n", channel.Name)
		}
		if channel.Name == name {
			if debug_mini {
				fmt.Printf("Found a Program Channel\n")
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
			// Otherwise find the settings available for this channel.
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
		}
	}
	// found some settings.
	if len(settingNames) > 0 {
		if debug_mini {
			fmt.Printf("Found Settings %+v\n", settingNames)
		}
		return settingNames, nil
	}

	return nil, fmt.Errorf("failed to find program settings for fixture%s", fixture.Name)

}

// getSwitchState reports on if this switch is running a mini sequence.
func getSwitchState(swiTch common.Switch, switchNumber int) bool {
	return swiTch.MiniSequencerRunning
}

func setSwitchState(swiTch common.Switch, state bool, blackout bool, master int) {
	swiTch.MiniSequencerRunning = state
	swiTch.Blackout = blackout
	swiTch.Master = master
}
