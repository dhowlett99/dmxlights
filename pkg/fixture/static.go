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
func setStaticOn(fixtureNumber int, cmd common.FixtureCommand, fixtures *Fixtures, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, dmxController *ft232.DMXController, dmxInterfacePresent bool) common.LastColor {

	if debug {
		fmt.Printf("Fixture:%d setStaticOn\n", fixtureNumber)
	}

	if cmd.RGBStaticColors[fixtureNumber].Enabled {
		if debug {
			fmt.Printf("%d: Fixture:%d RGB Switch Static On - Trying to Set RGB Static Master=%d\n", cmd.SequenceNumber, fixtureNumber, cmd.Master)
		}

		// TODO find sequence numbers from config.
		if cmd.SequenceNumber == 4 {
			cmd.SequenceNumber = 2
		}

		lamp := cmd.RGBStaticColors[fixtureNumber]

		// If we're not hiding the sequence on the launchpad, show the static colors on the buttons.
		if !cmd.Hidden {
			if lamp.Flash {
				onColor := color.RGBA{R: lamp.Color.R, G: lamp.Color.G, B: lamp.Color.B}
				common.FlashLight(common.Button{X: fixtureNumber, Y: cmd.SequenceNumber}, onColor, colors.Black, eventsForLaunchpad, guiButtons)
			} else {
				common.LightLamp(common.Button{X: fixtureNumber, Y: cmd.SequenceNumber}, lamp.Color, cmd.Master, eventsForLaunchpad, guiButtons)
			}
		}

		// Look for a matching color
		color := common.GetColorNameByRGB(lamp.Color)
		// Find a suitable gobo based on the requested static lamp color.
		scannerGobo := GetGoboByName(fixtureNumber, cmd.SequenceNumber, color, fixtures)
		// Find a suitable color wheel settin based on the requested static lamp color.
		scannerColor := FindColor(fixtureNumber, cmd.SequenceNumber, color, fixtures)

		return MapFixtures(false, false, cmd.SequenceNumber, fixtureNumber, lamp.Color, common.SCANNER_MID_POINT, common.SCANNER_MID_POINT, 0, 0, 0, scannerGobo, scannerColor, fixtures, cmd.Blackout, cmd.Master, cmd.Master, 0, cmd.Strobe, cmd.StrobeSpeed, dmxController, dmxInterfacePresent)
	}

	return common.LastColor{}
}

// Fade Up RGB Static Scene
func fadeUpStatic(fixtureNumber int, cmd common.FixtureCommand, lastColor common.LastColor, stopFadeDown chan bool, stopFadeUp chan bool, fixtures *Fixtures, fixtureStepChannel chan common.FixtureCommand, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, dmxController *ft232.DMXController, dmxInterfacePresent bool) {

	if debug {
		fmt.Printf("%d: fadeUpStaticFixture: Fixture No %d LastColor %+v\n", cmd.SequenceNumber, fixtureNumber, lastColor)
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

	lamp := cmd.RGBStaticColors[fixtureNumber]

	if debug {
		fmt.Printf("fadeUpStatic %d:%d Last Color %s This Color %s\n", cmd.SequenceNumber, fixtureNumber, common.GetColorNameByRGB(lastColor.RGBColor), common.GetColorNameByRGB(lamp.Color))
	}
	if lastColor.RGBColor != lamp.Color {
		// Now Fade Down
		go func() {
			// Soft start
			if debug {
				fmt.Printf("Fade Down %d\n", fixtureNumber)
			}
			// Calulate the steps
			fadeUpValues := common.GetFadeValues(64, float64(common.MAX_DMX_BRIGHTNESS), 1, false)
			fadeDownValues := common.GetFadeValues(64, float64(common.MAX_DMX_BRIGHTNESS), 1, true)

			master := cmd.Master

			if lastColor.RGBColor != colors.EmptyColor {
				for _, fade := range fadeDownValues {

					// Look for a matching color
					color := common.GetColorNameByRGB(lastColor.RGBColor)
					// Find a suitable gobo based on the requested static lamp color.
					scannerGobo := GetGoboByName(fixtureNumber, cmd.SequenceNumber, color, fixtures)
					// Find a suitable color wheel settin based on the requested static lamp color.
					scannerColor := FindColor(fixtureNumber, cmd.SequenceNumber, color, fixtures)

					// Listen for stop command.
					select {
					case <-stopFadeDown:
						return
					case <-time.After(10 * time.Millisecond):
					}
					if !cmd.Hidden {
						common.LightLamp(common.Button{X: fixtureNumber, Y: cmd.SequenceNumber}, lastColor.RGBColor, fade, eventsForLaunchpad, guiButtons)
					}
					if cmd.Label == "chaser" {
						// If we are a RGB chaser used as a shutter chasser apply fade values to the scanner's master dimmer channel because
						// scanners doesn't have a rgb color mixing capability so the wheel has to be faded using the master.
						master = int(float64(cmd.Master) / 100 * (float64(fade) / 2.55))
					}
					MapFixtures(false, false, cmd.SequenceNumber, fixtureNumber, lastColor.RGBColor, common.SCANNER_MID_POINT, common.SCANNER_MID_POINT, 0, 0, 0, scannerGobo, scannerColor, fixtures, cmd.Blackout, fade, master, 0, cmd.Strobe, cmd.StrobeSpeed, dmxController, dmxInterfacePresent)

					// Control how long the fade take with the speed control.
					time.Sleep((5 * time.Millisecond) * (time.Duration(cmd.RGBFade)))
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

			// If enabled fade up.
			if cmd.RGBStaticColors[fixtureNumber].Enabled {
				// Look for a matching color
				color := common.GetColorNameByRGB(lamp.Color)
				// Find a suitable gobo based on the requested static lamp color.
				scannerGobo := GetGoboByName(fixtureNumber, cmd.SequenceNumber, color, fixtures)
				// Find a suitable color wheel settin based on the requested static lamp color.
				scannerColor := FindColor(fixtureNumber, cmd.SequenceNumber, color, fixtures)

				// Fade up fixture.
				for _, fade := range fadeUpValues {
					// Listen for stop command.
					select {
					case <-stopFadeUp:
						lastColor = MapFixtures(false, false, cmd.SequenceNumber, fixtureNumber, colors.Black, 0, 0, 0, 0, 0, 0, 0, fixtures, false, 0, 0, 0, false, 0, dmxController, dmxInterfacePresent)
						return
					case <-time.After(10 * time.Millisecond):
					}
					if cmd.Label == "chaser" {
						// If we are a RGB chaser used as a shutter chasser apply fade values to the scanner's master dimmer channel because
						// scanners doesn't have a rgb color mixing capability so the wheel has to be faded using the master.
						master = int(float64(cmd.Master) / 100 * (float64(fade) / 2.55))
					}
					if !cmd.Hidden {
						common.LightLamp(common.Button{X: fixtureNumber, Y: cmd.SequenceNumber}, lamp.Color, fade, eventsForLaunchpad, guiButtons)
					}
					lastColor = MapFixtures(false, false, cmd.SequenceNumber, fixtureNumber, lamp.Color, common.SCANNER_MID_POINT, common.SCANNER_MID_POINT, 0, 0, 0, scannerGobo, scannerColor, fixtures, cmd.Blackout, fade, master, 0, cmd.Strobe, cmd.StrobeSpeed, dmxController, dmxInterfacePresent)

					// Control how long the fade take with the speed control.
					time.Sleep((5 * time.Millisecond) * (time.Duration(cmd.RGBFade)))
				}
				// Fade up complete, set lastColor up in the fixture.
				command := common.FixtureCommand{
					Type:         "lastColor",
					LastColor:    lastColor.RGBColor,
					ScannerColor: lastColor.ScannerColor,
				}
				select {
				case fixtureStepChannel <- command:
				case <-time.After(100 * time.Millisecond):
				}

			}
		}()
	}
}

func staticOff(fixtureNumber int, cmd common.FixtureCommand, lastColor common.LastColor, stopFadeDown chan bool, stopFadeUp chan bool, fixtures *Fixtures, fixtureStepChannel chan common.FixtureCommand, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, dmxController *ft232.DMXController, dmxInterfacePresent bool) {

	if debug {
		fmt.Printf("staticOff Fixture No %d", fixtureNumber)
	}

	go func() {
		var master int
		fadeDownValues := common.GetFadeValues(64, float64(common.MAX_DMX_BRIGHTNESS), 1, true)
		if lastColor.RGBColor != colors.Black {

			if debug {
				fmt.Printf("Fixture:%d =====>   RGB Static OFF -> Fade Down from LastColor %+v\n", fixtureNumber, lastColor)
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
				scannerColor := FindColor(fixtureNumber, cmd.SequenceNumber, color, fixtures)

				// Listen for stop commands.
				select {
				case <-stopFadeDown:
					return
				case <-stopFadeUp:
					return
				case <-time.After(10 * time.Millisecond):
				}
				common.LightLamp(common.Button{X: fixtureNumber, Y: sequenceNumber}, lastColor.RGBColor, fade, eventsForLaunchpad, guiButtons)
				if cmd.Label == "chaser" {
					// If we are a RGB chaser used as a shutter chasser apply fade values to the scanner's master dimmer channel because
					// scanners doesn't have a rgb color mixing capability so the wheel has to be faded using the master.
					master = int(float64(cmd.Master) / 100 * (float64(fade) / 2.55))
				}
				MapFixtures(false, false, cmd.SequenceNumber, fixtureNumber, lastColor.RGBColor, common.SCANNER_MID_POINT, common.SCANNER_MID_POINT, 0, 0, 0, scannerGobo, scannerColor, fixtures, cmd.Blackout, fade, master, 0, cmd.Strobe, cmd.StrobeSpeed, dmxController, dmxInterfacePresent)

				// Control how long the fade take with the speed control.
				time.Sleep((5 * time.Millisecond) * (time.Duration(cmd.RGBFade)))
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
	}()

}
