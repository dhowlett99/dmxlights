// Copyright (C) 2022, 2023 , 2024 dhowlett99.
// This is button processor, used by the launchpad and gui interfaces.
// This file processes flash buttons and controls their actions.
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

package buttons

import (
	"fmt"
	"time"

	"github.com/dhowlett99/dmxlights/pkg/colors"
	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
	"github.com/oliread/usbdmx/ft232"
)

func flashOn(sequences []*common.Sequence, X int, Y int, this *CurrentState, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, fixturesConfig *fixture.Fixtures, dmxController *ft232.DMXController) {

	this.SelectedType = sequences[Y].Type

	if debug {
		fmt.Printf("Flash ON Fixture Pressed X:%d Y:%d\n", X, Y)
	}
	colorPattern := 5
	flashSequence := common.Sequence{
		Pattern: common.Pattern{
			Name:  "colors",
			Steps: this.RGBPatterns[colorPattern].Steps, // Use the color pattern for flashing.
		},
	}

	pan := common.SCANNER_MID_POINT
	tilt := common.SCANNER_MID_POINT
	color := flashSequence.Pattern.Steps[X].Fixtures[X].Color
	shutter := flashSequence.Pattern.Steps[X].Fixtures[X].Shutter
	rotate := flashSequence.Pattern.Steps[X].Fixtures[X].Rotate
	music := flashSequence.Pattern.Steps[X].Fixtures[X].Music
	gobo := flashSequence.Pattern.Steps[X].Fixtures[X].Gobo
	program := flashSequence.Pattern.Steps[X].Fixtures[X].Program

	if this.SelectedType == "rgb" {
		common.LightLamp(common.Button{X: X, Y: Y}, color, this.MasterBrightness, eventsForLaunchpad, guiButtons)
		fixture.MapFixtures(false, false, Y, X, color, pan, tilt, shutter, rotate, program, gobo, 0, fixturesConfig, this.Blackout, this.MasterBrightness, this.MasterBrightness, music, this.Strobe[this.SelectedSequence], this.StrobeSpeed[this.SelectedSequence], dmxController, this.DmxInterfacePresent)
	}
	if this.SelectedType == "scanner" {
		common.LightLamp(common.Button{X: X, Y: Y}, colors.White, this.MasterBrightness, eventsForLaunchpad, guiButtons)
		fixture.MapFixtures(false, false, Y, X, color, pan, tilt, shutter, rotate, program, gobo, 0, fixturesConfig, this.Blackout, this.MasterBrightness, this.MasterBrightness, music, this.Strobe[this.SelectedSequence], this.StrobeSpeed[this.SelectedSequence], dmxController, this.DmxInterfacePresent)
	}

	if this.GUI {
		time.Sleep(200 * time.Millisecond)
		brightness := 0
		master := 0
		common.LightLamp(common.Button{X: X, Y: Y}, colors.Black, common.MIN_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)
		fixture.MapFixtures(false, false, Y, X, color, pan, tilt, shutter, rotate, program, gobo, 0, fixturesConfig, this.Blackout, brightness, master, music, this.Strobe[this.SelectedSequence], this.StrobeSpeed[this.SelectedSequence], dmxController, this.DmxInterfacePresent)
	}
}

func flashOff(X int, Y int, this *CurrentState, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, fixturesConfig *fixture.Fixtures, dmxController *ft232.DMXController) {

	if debug {
		fmt.Printf("Flash OFF Fixture Pressed X:%d Y:%d\n", X, Y)
	}

	X = X - 100

	pan := common.SCANNER_MID_POINT
	tilt := common.SCANNER_MID_POINT
	shutter := 0
	rotate := 0
	music := 0
	gobo := 0
	program := 0
	brightness := 0
	master := 0

	common.LightLamp(common.Button{X: X, Y: Y}, colors.Black, common.MIN_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)
	fixture.MapFixtures(false, false, Y, X, colors.Black, pan, tilt, shutter, rotate, program, gobo, 0, fixturesConfig, this.Blackout, brightness, master, music, this.Strobe[this.SelectedSequence], this.StrobeSpeed[this.SelectedSequence], dmxController, this.DmxInterfacePresent)
}
