// Copyright (C) 2022,2023,2024,2025 dhowlett99.
// This is the dmxlights scanner fixture code.
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
	"math"

	"github.com/dhowlett99/dmxlights/pkg/colors"
	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/oliread/usbdmx/ft232"
)

func playScanner(fixtureNumber int, cmd common.FixtureCommand, fixtures *Fixtures, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, dmxController *ft232.DMXController, dmxInterfacePresent bool) (lastColor common.LastColor) {

	if debug {
		fmt.Printf("Fixture:%d playScanner\n", fixtureNumber)
	}

	// Find the fixture
	fixture := cmd.ScannerPosition.Fixtures[fixtureNumber]

	if fixture.Enabled {

		if debug {
			fmt.Printf("Fixture:%d Play Scanner \n", fixtureNumber)
		}

		// In the case of a scanner, they usually have a shutter and a master dimmer to control the brightness
		// of the lamp. Problem is we can't use the shutter for the control of the overall brightness and the
		// master for the master dimmmer like we do with RGB fixture. The shutter noramlly is more of a switch
		// eg. Open , Closed , Strobe etc. If I want to slow fade through a set of scanners I need to use the
		// brightness for control. Which means I need to combine the master and the control brightness
		// at this stage.
		scannerBrightness := int(math.Round((float64(fixture.Brightness) / 100) * (float64(cmd.Master) / 2.55)))
		// Tell the scanner what to do.
		lastColor = MapFixtures(false, cmd.ScannerChaser, cmd.SequenceNumber, fixtureNumber, fixture.ScannerColor, fixture.Pan, fixture.Tilt,
			fixture.Shutter, cmd.Rotate, cmd.Program, cmd.ScannerGobo, cmd.ScannerColor, fixtures, cmd.Blackout, cmd.Master, scannerBrightness, cmd.Music, cmd.Strobe, cmd.StrobeSpeed, dmxController, dmxInterfacePresent)

		// Scannner is rotating, work out what to do with the launchpad lamps.
		if !cmd.Hidden {
			// Every quater turn, display a color to represent a position in the rotation.
			howOftern := cmd.NumberSteps / 4
			if howOftern != 0 {
				if cmd.Step%howOftern == 0 {
					// We're not in chase mode so use the color generated in the pattern generator.common.
					common.LightLamp(common.Button{X: fixtureNumber, Y: cmd.SequenceNumber}, fixture.Color, cmd.Master, eventsForLaunchpad, guiButtons)
					common.LabelButton(fixtureNumber, cmd.SequenceNumber, "", guiButtons)
				}
			}
		}
	} else {
		// This scanner is disabled, shut it off.
		lastColor = MapFixtures(false, false, cmd.SequenceNumber, fixtureNumber, colors.Black, 0, 0, 0, 0, 0, 0, 0, fixtures, false, 0, 0, 0, false, 0, dmxController, dmxInterfacePresent)
	}

	return lastColor
}
