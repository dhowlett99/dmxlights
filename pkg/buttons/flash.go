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
	"image/color"
	"time"

	"github.com/dhowlett99/dmxlights/pkg/colors"
	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
	"github.com/oliread/usbdmx/ft232"
)

func flashOn(sequence common.Sequence, X int, Y int, this *CurrentState, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, fixturesConfig *fixture.Fixtures, dmxController *ft232.DMXController) {

	// If we press a fixture button that doesn't have a real fixture behind it, then give up.
	if X == sequence.NumberFixtures {
		return
	}

	this.SelectedType = sequence.Type

	if debug {
		fmt.Printf("Flash ON Fixture Pressed X:%d Y:%d Target Seq %d Fixtures=%d\n", X, Y, this.TargetSequence, sequence.NumberFixtures)
	}

	colorPattern := 5
	flashSequence := common.Sequence{
		Pattern: common.Pattern{
			Name:  "colors",
			Steps: sequence.RGBAvailablePatterns[colorPattern].Steps, // Use the color pattern for flashing.
		},
	}

	pan := common.SCANNER_MID_POINT
	tilt := common.SCANNER_MID_POINT
	color := colors.White
	shutter := flashSequence.Pattern.Steps[X].Fixtures[X].Shutter
	rotate := flashSequence.Pattern.Steps[X].Fixtures[X].Rotate
	music := flashSequence.Pattern.Steps[X].Fixtures[X].Music
	gobo := flashSequence.Pattern.Steps[X].Fixtures[X].Gobo
	program := flashSequence.Pattern.Steps[X].Fixtures[X].Program
	programSpeed := flashSequence.Pattern.Steps[X].Fixtures[X].ProgramSpeed

	if this.SelectedType == "rgb" {
		common.LightLamp(common.Button{X: X, Y: Y}, color, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)
		fixture.MapFixtures(false, false, false, Y, X, color, color, pan, tilt, shutter, rotate, program, programSpeed, gobo, 0, fixturesConfig, this.Blackout, common.MAX_DMX_BRIGHTNESS, common.MAX_DMX_BRIGHTNESS, music, this.Strobe[this.SelectedSequence], this.StrobeSpeed[this.SelectedSequence], dmxController, this.DmxInterfacePresent)
	}
	if this.SelectedType == "scanner" {
		common.LightLamp(common.Button{X: X, Y: Y}, colors.White, this.MasterBrightness[this.TargetSequence], eventsForLaunchpad, guiButtons)
		fixture.MapFixtures(false, false, false, Y, X, color, color, pan, tilt, shutter, rotate, program, programSpeed, gobo, 0, fixturesConfig, this.Blackout, common.MAX_DMX_BRIGHTNESS, common.MAX_DMX_BRIGHTNESS, music, this.Strobe[this.SelectedSequence], this.StrobeSpeed[this.SelectedSequence], dmxController, this.DmxInterfacePresent)
	}

	if this.GUI {
		time.Sleep(200 * time.Millisecond)
		fixtureColor, master, brightness := handleStatic(sequence, Y, X)
		common.LightLamp(common.Button{X: X, Y: Y}, fixtureColor, brightness, eventsForLaunchpad, guiButtons)
		fixture.MapFixtures(false, false, false, Y, X, color, color, pan, tilt, shutter, rotate, program, programSpeed, gobo, 0, fixturesConfig, this.Blackout, brightness, master, music, this.Strobe[this.SelectedSequence], this.StrobeSpeed[this.SelectedSequence], dmxController, this.DmxInterfacePresent)
	}
}

func handleStatic(sequence common.Sequence, sequenceNumber int, fixtureNumber int) (color.RGBA, int, int) {

	var master int
	var brightness int
	var fixtureColor color.RGBA

	// Check for any static color for this fixture.
	if sequence.Static {
		// Restore the static color.
		master = 255
		brightness = sequence.Master
		fixtureColor = sequence.StaticColors[fixtureNumber].Color
	} else {
		// Switch off.
		master = 0
		brightness = 0
		fixtureColor = colors.Black
	}
	return fixtureColor, master, brightness
}

func flashOff(sequence common.Sequence, X int, Y int, this *CurrentState, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, fixturesConfig *fixture.Fixtures, dmxController *ft232.DMXController) {

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
	programSpeed := 0
	brightness := 0
	master := 0

	// Check for static.
	fixtureColor, master, brightness := handleStatic(sequence, Y, X)

	common.LightLamp(common.Button{X: X, Y: Y}, fixtureColor, brightness, eventsForLaunchpad, guiButtons)
	fixture.MapFixtures(false, false, false, Y, X, fixtureColor, fixtureColor, pan, tilt, shutter, rotate, program, programSpeed, gobo, 0, fixturesConfig, this.Blackout, brightness, master, music, this.Strobe[this.SelectedSequence], this.StrobeSpeed[this.SelectedSequence], dmxController, this.DmxInterfacePresent)
}
