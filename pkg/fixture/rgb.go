// Copyright (C) 2022,2023,2024,2025 dhowlett99.
// This is the dmxlights rgb fixture code.
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

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/oliread/usbdmx/ft232"
)

// Now play all the values for this state.
func playRGB(fixtureNumber int, cmd common.FixtureCommand, fixtures *Fixtures, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, dmxController *ft232.DMXController, dmxInterfacePresent bool) (lastColor common.LastColor) {

	if debug {
		fmt.Printf("playRGB: fixtureNumber %d", fixtureNumber)
	}

	// Play out fixture to DMX channels.
	fixture := cmd.RGBPosition.Fixtures[fixtureNumber]

	if cmd.Type == "rgb" && fixture.Enabled {

		if debug {
			fmt.Printf("%d: Fixture:%d RGB Mode Strobe %t\n", cmd.SequenceNumber, fixtureNumber, cmd.Strobe)
		}

		// Integrate cmd.master with fixture.Brightness.
		fixture.Brightness = int((float64(fixture.Brightness) / 100) * (float64(cmd.Master) / 2.55))

		// If we're a shutter chaser flavoured RGB sequence, then disable everything except the brightness.
		if cmd.Label == "chaser" {
			scannerFixturesSequenceNumber := common.GlobalScannerSequenceNumber // Scanner sequence number from config.
			if !cmd.Hidden {
				common.LightLamp(common.Button{X: fixtureNumber, Y: scannerFixturesSequenceNumber}, fixture.Color, fixture.Brightness, eventsForLaunchpad, guiButtons)
			}

			// Fixture brightness is sent as master in this case because a shutter chaser is controlling a scanner lamp.
			// and these generally don't have any RGB color channels that can be controlled with brightness.
			// So the only way to make the lamp in the scanner change intensity is to vary the master brightness channel.

			// Lookup chaser lamp color based on the requested fixture base color.
			// We can't use the faded color as its impossibe to lookup the base color from a faded color.
			// GetColorNameByRGB will return white if the color is not found.
			color := common.GetColorNameByRGB(fixture.BaseColor)

			// Find a suitable gobo based on the requested chaser lamp color.
			scannerGobo := GetGoboByName(fixtureNumber, scannerFixturesSequenceNumber, color, fixtures)
			// Find a suitable color wheel setting based on the requested static lamp color.
			scannerColor := GetColor(fixtureNumber, scannerFixturesSequenceNumber, color, fixtures)

			lastColor = MapFixtures(true, cmd.ScannerChaser, scannerFixturesSequenceNumber, fixtureNumber, fixture.Color, fixture.Color, 0, 0, 0, 0, 0, scannerGobo, scannerColor, fixtures, cmd.Blackout, cmd.Master, fixture.Brightness, cmd.Music, cmd.Strobe, cmd.StrobeSpeed, dmxController, dmxInterfacePresent)
		} else {
			if !cmd.Hidden {
				if fixtureNumber > 7 {
					fixtureNumber = fixtureNumber - 8
					cmd.SequenceNumber = cmd.SequenceNumber + 1
				}
				common.LightLamp(common.Button{X: fixtureNumber, Y: cmd.SequenceNumber}, fixture.Color, cmd.Master, eventsForLaunchpad, guiButtons)
			}
			lastColor = MapFixtures(false, cmd.ScannerChaser, cmd.SequenceNumber, fixtureNumber, fixture.Color, fixture.Color, 0, 0, 0, 0, 0, cmd.ScannerGobo, cmd.ScannerColor, fixtures, cmd.Blackout, cmd.Master, cmd.Master, cmd.Music, cmd.Strobe, cmd.StrobeSpeed, dmxController, dmxInterfacePresent)
		}
	}

	return lastColor
}
