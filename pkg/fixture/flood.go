// Copyright (C) 2022,2023,2024,2025 dhowlett99.
// This is the dmxlights fixture flood control code, it sends messages to fixtures
// to turn on and off flood.
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

	"github.com/dhowlett99/dmxlights/pkg/colors"
	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/oliread/usbdmx/ft232"
)

// Start Flood.
func startFlood(fixtureNumber int, cmd common.FixtureCommand, fixtures *Fixtures, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, dmxController *ft232.DMXController, dmxInterfacePresent bool) common.LastColor {
	if debug {
		fmt.Printf("Fixture:%d Set RGB Flood\n", fixtureNumber)
	}

	if fixtureNumber > 7 {
		fixtureNumber = fixtureNumber - 8
		cmd.SequenceNumber = cmd.SequenceNumber + 1
	}

	// TODO find sequence numbers from config.
	if cmd.SequenceNumber == 4 {
		cmd.SequenceNumber = 2
	}

	pan := 128
	tilt := 128
	shutter := GetShutter(fixtureNumber, cmd.SequenceNumber, "Open", fixtures)
	gobo := GetGoboByName(fixtureNumber, cmd.SequenceNumber, "White", fixtures)
	scannerColor := GetColor(fixtureNumber, cmd.SequenceNumber, "White", fixtures)
	rotate := 0
	program := 0

	if !cmd.Hidden {
		common.LightLamp(common.Button{X: fixtureNumber, Y: cmd.SequenceNumber}, colors.White, cmd.Master, eventsForLaunchpad, guiButtons)
		common.LabelButton(fixtureNumber, cmd.SequenceNumber, "", guiButtons)
	}

	return MapFixtures(false, false, cmd.SequenceNumber, fixtureNumber, colors.White, colors.White, pan, tilt, shutter, rotate, program, gobo, scannerColor, fixtures, false, cmd.Master, cmd.Master, 0, cmd.Strobe, cmd.StrobeSpeed, dmxController, dmxInterfacePresent)

}

// Stop Flood.
func stopFlood(fixtureNumber int, cmd common.FixtureCommand, fixtures *Fixtures, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, dmxController *ft232.DMXController, dmxInterfacePresent bool) common.LastColor {

	if debug {
		fmt.Printf("Fixture:%d Set Stop RGB Flood\n", fixtureNumber)
	}

	if fixtureNumber > 7 {
		fixtureNumber = fixtureNumber - 8
		cmd.SequenceNumber = cmd.SequenceNumber + 1
	}

	// TODO find sequence numbers from config.
	if cmd.SequenceNumber == 4 {
		cmd.SequenceNumber = 2
	}

	if !cmd.Hidden {
		common.LightLamp(common.Button{X: fixtureNumber, Y: cmd.SequenceNumber}, colors.Black, 0, eventsForLaunchpad, guiButtons)
		common.LabelButton(fixtureNumber, cmd.SequenceNumber, "", guiButtons)
	}
	return MapFixtures(false, false, cmd.SequenceNumber, fixtureNumber, colors.Black, colors.Black, 0, 0, 0, 0, 0, 0, 0, fixtures, cmd.Blackout, 0, 0, 0, cmd.Strobe, cmd.StrobeSpeed, dmxController, dmxInterfacePresent)
}
