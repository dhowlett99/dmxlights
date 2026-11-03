// Copyright (C) 2022,2023,2024,2025 dhowlett99.
// This is the dmxlights fixture static scene code, it sends messages to fixtures
// to fade up and down static scences.
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
	"time"

	"github.com/dhowlett99/dmxlights/pkg/colors"
	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/oliread/usbdmx/ft232"
)

// Switch On Static Scene.
func setStaticOn(fixtureNumber int, cmd common.FixtureCommand, fixtures *Fixtures, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, dmxController *ft232.DMXController, dmxInterfacePresent bool) common.FixtureColor {

	if debug && fixtureNumber == 00 {
		fmt.Printf("setStaticOn: Fixture %d\n", fixtureNumber)
	}

	if cmd.RGBStaticColors[fixtureNumber].Enabled {
		if debug && fixtureNumber == 00 {
			fmt.Printf("setStaticOn: Fixture %d Master=%d\n", fixtureNumber, cmd.Master)
		}

		// TODO find sequence numbers from config.
		if cmd.SequenceNumber == 4 {
			cmd.SequenceNumber = 2
		}

		nextColor := common.FixtureColor{
			RGBColor:   cmd.RGBStaticColors[fixtureNumber].Color,
			Brightness: cmd.Master,
			Flash:      cmd.RGBStaticColors[fixtureNumber].Flash,
		}

		// If we're not hiding the sequence on the launchpad, show the static colors on the buttons.
		if !cmd.Hidden {
			if nextColor.Flash {
				onColor := color.RGBA{R: nextColor.RGBColor.R, G: nextColor.RGBColor.G, B: nextColor.RGBColor.B, A: nextColor.RGBColor.A}
				common.FlashLight(common.Button{X: fixtureNumber, Y: cmd.SequenceNumber}, onColor, colors.Black, eventsForLaunchpad, guiButtons)
			} else {
				common.LightLamp(common.Button{X: fixtureNumber, Y: cmd.SequenceNumber}, nextColor.RGBColor, cmd.Master, eventsForLaunchpad, guiButtons)
			}
		}

		// Look for a matching color
		color := common.GetColorNameByRGB(nextColor.RGBColor)
		// Find a suitable gobo based on the requested static lamp color.
		scannerGobo := GetGoboByName(fixtureNumber, cmd.SequenceNumber, color, fixtures)
		// Find a suitable color wheel settin based on the requested static lamp color.
		scannerColor := GetColor(fixtureNumber, cmd.SequenceNumber, color, fixtures)

		var iAmAChaserSequence bool
		if cmd.Label == "chaser" {
			iAmAChaserSequence = true
		} else {
			iAmAChaserSequence = false
		}
		return MapFixtures(iAmAChaserSequence, cmd.ScannerChaser, cmd.RotateRunning, cmd.SequenceNumber, fixtureNumber, nextColor.RGBColor, nextColor.RGBColor, cmd.ScannerOffsetPan, cmd.ScannerOffsetTilt, 0, 0, 0, 0, scannerGobo, scannerColor, fixtures, cmd.Blackout, nextColor.Brightness, nextColor.Brightness, 0, cmd.Strobe, cmd.StrobeSpeed, dmxController, dmxInterfacePresent)
	}

	return common.FixtureColor{}
}

// Fade Up RGB Static Scene
func fadeUpStatic(fixtureNumber int, cmd common.FixtureCommand, lastColor common.FixtureColor, stopFadeDown chan bool, stopFadeUp chan bool, fixtures *Fixtures, fixtureStepChannel chan common.FixtureCommand, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, dmxController *ft232.DMXController, dmxInterfacePresent bool) {

	nextColor := common.FixtureColor{
		RGBColor:   cmd.RGBStaticColors[fixtureNumber].Color,
		Brightness: cmd.Master,
		Flash:      cmd.RGBStaticColors[fixtureNumber].Flash,
	}

	if debug && fixtureNumber == 0 {
		fmt.Printf("%d: fadeUpStatic fixture number %d \n", cmd.SequenceNumber, fixtureNumber)
		fmt.Printf("%d: Last Color %+v Last Brightnes %d\n", cmd.SequenceNumber, lastColor.RGBColor, lastColor.Brightness)
		fmt.Printf("%d: Next Color %+v Next Brightnes %d\n", cmd.SequenceNumber, nextColor.RGBColor, nextColor.Brightness)
	}

	if cmd.SequenceNumber == 4 {
		cmd.SequenceNumber = 2
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

	// They are the same color but the brightness is different.
	if (lastColor.RGBColor == nextColor.RGBColor) && (lastColor.Brightness != nextColor.Brightness) {

		// If the requested next color brightness is greater than the last color brightness just fade up to that brightness.
		if nextColor.Brightness > lastColor.Brightness {

			if debug && fixtureNumber == 0 {
				fmt.Printf("%d: requested colors are the same but the brightness is greater than the last brightness\n", cmd.SequenceNumber)
			}

			go func() {

				// If enabled fade up.
				if cmd.RGBStaticColors[fixtureNumber].Enabled {
					fadeUp(fixtureNumber, cmd, lastColor, nextColor, stopFadeUp, fixtures, fixtureStepChannel, eventsForLaunchpad, guiButtons, dmxController, dmxInterfacePresent)
				}
			}()
			return
		}

		// if the requested brightness is less than the last brightness just fade down to that brightness.
		if nextColor.Brightness < lastColor.Brightness {
			if debug && fixtureNumber == 0 {
				fmt.Printf("%d: requested colors are the same but requested color brightness is less than the last colors brightness\n", cmd.SequenceNumber)
			}
			go func() {
				// If requested color is not empty fade down.nextColor is not black
				if lastColor.RGBColor != colors.EmptyColor {
					lastColor = fadeDown(fixtureNumber, cmd, lastColor, nextColor, stopFadeDown, fixtures, fixtureStepChannel, eventsForLaunchpad, guiButtons, dmxController, dmxInterfacePresent)
				}
			}()
			return
		}
	}

	// If colors are completely different.
	if lastColor.RGBColor != nextColor.RGBColor {
		if debug && fixtureNumber == 0 {
			fmt.Printf("%d: colors are completely different.\n", cmd.SequenceNumber)
		}
		// Now Fade Down
		go func() {

			// If last color is not black fade down first.
			if lastColor.RGBColor != colors.Black {
				if debug && fixtureNumber == 0 {
					fmt.Printf("%d: nextColor is not black.\n", cmd.SequenceNumber)
				}
				// Make the off color black, brightness zero so fade to black works.
				offColor := common.FixtureColor{}
				offColor.RGBColor = color.RGBA{0, 0, 0, 0} // Black
				offColor.Brightness = common.MIN_DMX_BRIGHTNESS
				lastColor = fadeDown(fixtureNumber, cmd, lastColor, offColor, stopFadeDown, fixtures, fixtureStepChannel, eventsForLaunchpad, guiButtons, dmxController, dmxInterfacePresent)
			}

			// If enabled fade up.
			if cmd.RGBStaticColors[fixtureNumber].Enabled {
				fadeUp(fixtureNumber, cmd, lastColor, nextColor, stopFadeUp, fixtures, fixtureStepChannel, eventsForLaunchpad, guiButtons, dmxController, dmxInterfacePresent)
			}
		}()
		return
	}

	if debug && fixtureNumber == 0 {
		fmt.Printf("%d: everything is the same, nothing to see here.\n", cmd.SequenceNumber)
	}
}

func fadeDown(fixtureNumber int, cmd common.FixtureCommand, lastColor common.FixtureColor, nextColor common.FixtureColor, stopFadeDown chan bool, fixtures *Fixtures, fixtureStepChannel chan common.FixtureCommand, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, dmxController *ft232.DMXController, dmxInterfacePresent bool) common.FixtureColor {

	if debug && fixtureNumber == 0 {
		fmt.Printf("%d fadeDown fixture number %d sequence Number %d\n", cmd.SequenceNumber, fixtureNumber, cmd.SequenceNumber)
		fmt.Printf("%d Last Color %+v Last Brightnes %d\n", cmd.SequenceNumber, lastColor.RGBColor, lastColor.Brightness)
		fmt.Printf("%d Next Color %+v Next Brightnes %d\n", cmd.SequenceNumber, nextColor.RGBColor, nextColor.Brightness)
	}
	fadeDownValues := common.GetFadeValues(64, lastColor.Brightness, nextColor.Brightness, 1, true)
	if debug && fixtureNumber == 0 {
		fmt.Printf("%d fadeDown: Fadedown values %+v\n", cmd.SequenceNumber, fadeDownValues)
	}

	// Look for a matching color
	color := common.GetColorNameByRGB(lastColor.RGBColor)
	// Find a suitable gobo based on the requested static lamp color.
	scannerGobo := GetGoboByName(fixtureNumber, cmd.SequenceNumber, color, fixtures)
	// Find a suitable color wheel settin based on the requested static lamp color.
	scannerColor := GetColor(fixtureNumber, cmd.SequenceNumber, color, fixtures)

	for _, fade := range fadeDownValues {

		// Listen for stop command.
		select {
		case <-stopFadeDown:
			return lastColor
		case <-time.After(10 * time.Millisecond):
		}
		if !cmd.Hidden {
			fade := applyMasterToFade(fade, cmd.Master)
			common.LightLamp(common.Button{X: fixtureNumber, Y: cmd.SequenceNumber}, lastColor.RGBColor, fade, eventsForLaunchpad, guiButtons)
		}
		var iAmAChaserSequence bool
		if cmd.Label == "chaser" {
			iAmAChaserSequence = true
			// If we are a RGB chaser used as a shutter chasser apply fade values to the scanner's master dimmer channel because
			// scanners doesn't have a rgb color mixing capability so the wheel has to be faded using the master.
			cmd.Master = fade
		} else {
			iAmAChaserSequence = false
		}
		MapFixtures(iAmAChaserSequence, cmd.ScannerChaser, cmd.RotateRunning, cmd.SequenceNumber, fixtureNumber, nextColor.RGBColor, nextColor.RGBColor, cmd.ScannerOffsetPan, cmd.ScannerOffsetTilt, 0, 0, 0, 0, scannerGobo, scannerColor, fixtures, cmd.Blackout, fade, cmd.Master, 0, cmd.Strobe, cmd.StrobeSpeed, dmxController, dmxInterfacePresent)

		// Control how long the fade take with the speed control.
		time.Sleep((5 * time.Millisecond) * (time.Duration(cmd.RGBFade)))
	}
	// Fade down complete, set lastColor to the final color and brightness.
	lastColor = common.FixtureColor{
		RGBColor:     lastColor.RGBColor,
		ScannerColor: 0,
		Brightness:   nextColor.Brightness,
	}
	command := common.FixtureCommand{
		Type:      "lastColor",
		LastColor: lastColor,
	}
	select {
	case fixtureStepChannel <- command:
	case <-time.After(100 * time.Millisecond):
	}

	if debug && fixtureNumber == 0 {
		fmt.Printf("%d fadeDown: End of Fade Down LastColor was %+v Master %d\n", cmd.SequenceNumber, lastColor, cmd.Master)
	}

	return lastColor

}

func fadeUp(fixtureNumber int, cmd common.FixtureCommand, lastColor common.FixtureColor, nextColor common.FixtureColor, stopFadeUp chan bool, fixtures *Fixtures, fixtureStepChannel chan common.FixtureCommand, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, dmxController *ft232.DMXController, dmxInterfacePresent bool) {

	if debug && fixtureNumber == 0 {
		fmt.Printf("%d: fadeUP fixture number %d \n", cmd.SequenceNumber, fixtureNumber)
		fmt.Printf("%d: Last Color %+v Last Brightnes %d\n", cmd.SequenceNumber, lastColor.RGBColor, lastColor.Brightness)
		fmt.Printf("%d: Next Color %+v Next Brightnes %d\n", cmd.SequenceNumber, nextColor.RGBColor, nextColor.Brightness)
	}

	// Calulate the steps
	fadeUpValues := common.GetFadeValues(64, lastColor.Brightness, nextColor.Brightness, 1, false)
	if debug && fixtureNumber == 0 {
		fmt.Printf("%d: fadeup: Fadeup values %+v\n", cmd.SequenceNumber, fadeUpValues)
	}

	// Look for a matching color

	color := common.GetColorNameByRGB(nextColor.RGBColor)
	// Find a suitable gobo based on the requested static lamp color.
	scannerGobo := GetGoboByName(fixtureNumber, cmd.SequenceNumber, color, fixtures)
	// Find a suitable color wheel settin based on the requested static lamp color.
	scannerColor := GetColor(fixtureNumber, cmd.SequenceNumber, color, fixtures)

	// Fade up fixture.
	for _, fade := range fadeUpValues {
		// Listen for stop command.
		select {
		case <-stopFadeUp:
			lastColor = MapFixtures(false, false, false, cmd.SequenceNumber, fixtureNumber, colors.Black, colors.Black, 0, 0, 0, 0, 0, 0, 0, 0, fixtures, false, 0, 0, 0, false, 0, dmxController, dmxInterfacePresent)
			// Fade up complete, set lastColor up in the fixture.
			command := common.FixtureCommand{
				Type:      "lastColor",
				LastColor: lastColor,
			}
			select {
			case fixtureStepChannel <- command:
			case <-time.After(100 * time.Millisecond):
			}
			return
		case <-time.After(10 * time.Millisecond):
		}

		// If we are a RGB chaser used as a shutter chasser apply fade values to the scanner's master dimmer channel because
		// scanners doesn't have a rgb color mixing capability so the wheel has to be faded using the master.
		var iAmAChaserSequence bool
		var master int
		if cmd.Label == "chaser" {
			iAmAChaserSequence = true
			master = fade
		} else {
			iAmAChaserSequence = false
			master = cmd.Master
		}

		// Now play out the fade up commands.
		if !cmd.Hidden {
			fade := applyMasterToFade(fade, master)
			common.LightLamp(common.Button{X: fixtureNumber, Y: cmd.SequenceNumber}, nextColor.RGBColor, fade, eventsForLaunchpad, guiButtons)
		}
		lastColor = MapFixtures(iAmAChaserSequence, cmd.ScannerChaser, cmd.RotateRunning, cmd.SequenceNumber, fixtureNumber, nextColor.RGBColor, nextColor.RGBColor, cmd.ScannerOffsetPan, cmd.ScannerOffsetTilt, 0, 0, 0, 0, scannerGobo, scannerColor, fixtures, cmd.Blackout, fade, master, 0, cmd.Strobe, cmd.StrobeSpeed, dmxController, dmxInterfacePresent)

		// Control how long the fade take with the speed control.
		time.Sleep((5 * time.Millisecond) * (time.Duration(cmd.RGBFade)))
	}

	lastColor.Brightness = nextColor.Brightness

	// Fade up complete, set lastColor up in the fixture.
	command := common.FixtureCommand{
		Type:      "lastColor",
		LastColor: lastColor,
	}
	select {
	case fixtureStepChannel <- command:
	case <-time.After(100 * time.Millisecond):
	}

	if debug && fixtureNumber == 0 {
		fmt.Printf("%d: fadeUp: Fade up complete, set lastColor up in the fixture lastColor is %v\n", cmd.SequenceNumber, lastColor)
	}
}

func staticOff(fixtureNumber int, cmd common.FixtureCommand, lastColor common.FixtureColor, stopFadeDown chan bool, stopFadeUp chan bool, fixtures *Fixtures, fixtureStepChannel chan common.FixtureCommand, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, dmxController *ft232.DMXController, dmxInterfacePresent bool) {

	if debug && fixtureNumber == 0 {
		fmt.Printf("staticOff Fixture No %d lastColor %v\n", fixtureNumber, lastColor.RGBColor)
	}

	go func() {

		// Changing just the brightnes.
		if lastColor.Brightness != cmd.Master {
			return
		}

		if lastColor.RGBColor != colors.Black {
			var master int
			fadeDownValues := common.GetFadeValues(64, lastColor.Brightness, cmd.Master, 1, true)
			if debug && fixtureNumber == 0 {
				fmt.Printf("staticOff: Fadedown values %+v\n", fadeDownValues)
			}

			if debug && fixtureNumber == 0 {
				fmt.Printf("Fixture:%d =====>   RGB Static OFF -> Fade Down from LastColor %+v fadeDownValues %v\n", fixtureNumber, lastColor, fadeDownValues)
			}

			var sequenceNumber int
			if cmd.Label == "chaser" {
				sequenceNumber = common.GlobalScannerSequenceNumber // Scanner sequence number from config.
			} else {
				sequenceNumber = cmd.SequenceNumber
			}

			for _, fade := range fadeDownValues {

				// Look for a matching color
				color := common.GetColorNameByRGB(lastColor.RGBColor)
				// Find a suitable gobo based on the requested static lamp color.
				scannerGobo := GetGoboByName(fixtureNumber, cmd.SequenceNumber, color, fixtures)
				// Find a suitable color wheel settin based on the requested static lamp color.
				scannerColor := GetColor(fixtureNumber, cmd.SequenceNumber, color, fixtures)

				// Listen for stop commands.
				select {
				case <-stopFadeDown:
					return
				case <-stopFadeUp:
					return
				case <-time.After(10 * time.Millisecond):
				}
				fade := applyMasterToFade(fade, master)
				common.LightLamp(common.Button{X: fixtureNumber, Y: sequenceNumber}, lastColor.RGBColor, fade, eventsForLaunchpad, guiButtons)
				if cmd.Label == "chaser" {
					// If we are a RGB chaser used as a shutter chasser apply fade values to the scanner's master dimmer channel because
					// scanners doesn't have a rgb color mixing capability so the wheel has to be faded using the master.
					master = int(float64(cmd.Master) / 100 * (float64(fade) / 2.55))
				}
				var iAmAChaserSequence bool
				if cmd.Label == "chaser" {
					iAmAChaserSequence = true
				} else {
					iAmAChaserSequence = false
				}
				MapFixtures(iAmAChaserSequence, cmd.ScannerChaser, cmd.RotateRunning, cmd.SequenceNumber, fixtureNumber, lastColor.RGBColor, lastColor.RGBColor, cmd.ScannerOffsetPan, cmd.ScannerOffsetTilt, 0, 0, 0, 0, scannerGobo, scannerColor, fixtures, cmd.Blackout, fade, master, 0, cmd.Strobe, cmd.StrobeSpeed, dmxController, dmxInterfacePresent)

				// Control how long the fade take with the speed control.
				time.Sleep((5 * time.Millisecond) * (time.Duration(cmd.RGBFade)))
			}
			// Fade down complete, set lastColor to empty in the fixture.
			empty := common.FixtureColor{
				RGBColor:     colors.EmptyColor,
				ScannerColor: 0,
				Brightness:   common.MIN_DMX_BRIGHTNESS,
			}
			command := common.FixtureCommand{
				Type:      "lastColor",
				LastColor: empty,
			}
			select {
			case fixtureStepChannel <- command:
			case <-time.After(100 * time.Millisecond):
			}
		}
	}()

}
