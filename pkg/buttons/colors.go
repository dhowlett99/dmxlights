// Copyright (C) 2022, 2023 , 2024 dhowlett99.
// This is button processor, used by the launchpad and gui interfaces.
// This file processes color adjustments.
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

func redButton(this *CurrentState, X int, Y int, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command) {

	if this.Functions[this.SelectedSequence][common.Function6_Static_Gobo].State {

		if debug {
			fmt.Printf("Choose Static Red X:%d Y:%d\n", X, Y)
		}

		buttonTouched(common.Button{X: X, Y: Y}, colors.White, colors.Red, eventsForLaunchpad, guiButtons)

		this.StaticButtons[this.SelectedSequence].X = this.LastStaticColorButtonX
		this.StaticButtons[this.SelectedSequence].Y = this.LastStaticColorButtonY
		this.StaticButtons[this.SelectedSequence].Color.R = this.StaticButtons[this.SelectedSequence].Color.R - 10
		if this.StaticButtons[this.SelectedSequence].Color.R == 254 {
			this.StaticButtons[this.SelectedSequence].Color.R = 0
		}
		if this.StaticButtons[this.SelectedSequence].Color.R == 0 {
			this.StaticButtons[this.SelectedSequence].Color.R = 254
		}

		redColor := color.RGBA{R: this.StaticButtons[this.SelectedSequence].Color.R, G: 0, B: 0}
		common.LightLamp(common.RED_BUTTON, redColor, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)
		updateStaticLamp(this.SelectedSequence, this.StaticButtons[this.SelectedSequence], commandChannels)

		// Update the status bar
		common.UpdateStatusBar(fmt.Sprintf("Red %02d", this.StaticButtons[this.SelectedSequence].Color.R), "red", false, guiButtons)
		return
	}
}

func greenButton(this *CurrentState, X int, Y int, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command) {

	if this.Functions[this.SelectedSequence][common.Function6_Static_Gobo].State {

		if debug {
			fmt.Printf("Choose Static Green X:%d Y:%d\n", X, Y)
		}

		buttonTouched(common.Button{X: X, Y: Y}, colors.White, colors.Green, eventsForLaunchpad, guiButtons)

		this.StaticButtons[this.SelectedSequence].X = this.LastStaticColorButtonX
		this.StaticButtons[this.SelectedSequence].Y = this.LastStaticColorButtonY
		this.StaticButtons[this.SelectedSequence].Color.G = this.StaticButtons[this.SelectedSequence].Color.G - 10
		if this.StaticButtons[this.SelectedSequence].Color.G == 254 {
			this.StaticButtons[this.SelectedSequence].Color.G = 0
		}
		if this.StaticButtons[this.SelectedSequence].Color.G == 0 {
			this.StaticButtons[this.SelectedSequence].Color.G = 254
		}
		greenColor := color.RGBA{R: 0, G: this.StaticButtons[this.SelectedSequence].Color.G, B: 0}
		common.LightLamp(common.Button{X: X, Y: Y}, greenColor, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)
		updateStaticLamp(this.SelectedSequence, this.StaticButtons[this.SelectedSequence], commandChannels)

		// Update the status bar
		common.UpdateStatusBar(fmt.Sprintf("Green %02d", this.StaticButtons[this.SelectedSequence].Color.G), "green", false, guiButtons)
		return
	}
}
func blueButton(this *CurrentState, X int, Y int, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command) {

	if this.Functions[this.SelectedSequence][common.Function6_Static_Gobo].State {

		if debug {
			fmt.Printf("Choose Static Blue X:%d Y:%d\n", X, Y)
		}

		buttonTouched(common.Button{X: X, Y: Y}, colors.White, colors.Blue, eventsForLaunchpad, guiButtons)

		this.StaticButtons[this.SelectedSequence].X = this.LastStaticColorButtonX
		this.StaticButtons[this.SelectedSequence].Y = this.LastStaticColorButtonY
		this.StaticButtons[this.SelectedSequence].Color.B = this.StaticButtons[this.SelectedSequence].Color.B - 10
		if this.StaticButtons[this.SelectedSequence].Color.B > 254 {
			this.StaticButtons[this.SelectedSequence].Color.B = 0
		}
		if this.StaticButtons[this.SelectedSequence].Color.B == 0 {
			this.StaticButtons[this.SelectedSequence].Color.B = 254
		}
		blueColor := color.RGBA{R: 0, G: 0, B: this.StaticButtons[this.SelectedSequence].Color.B}
		common.LightLamp(common.Button{X: X, Y: Y}, blueColor, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)
		updateStaticLamp(this.SelectedSequence, this.StaticButtons[this.SelectedSequence], commandChannels)

		// Update the status bar
		common.UpdateStatusBar(fmt.Sprintf("Blue %02d", this.StaticButtons[this.SelectedSequence].Color.B), "blue", false, guiButtons)
		return
	}
}

func selectRGBChaseColor(sequences []*common.Sequence, X int, Y int, this *CurrentState, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command) {

	if debug {
		fmt.Printf("Set Sequence Color X:%d Y:%d\n", X, Y)
	}

	if this.SelectedMode[this.SelectedSequence] == CHASER_FUNCTION {
		this.TargetSequence = this.ChaserSequenceNumber
		this.DisplaySequence = this.SelectedSequence
	} else {
		this.TargetSequence = this.SelectedSequence
		this.DisplaySequence = this.SelectedSequence
	}

	// Reset the clear button so you can clear this selection if required.
	this.ClearPressed[this.TargetSequence] = false

	// Add the selected color to the sequence.
	newColor := common.GetColor(X, Y)
	sequences[this.TargetSequence].SequenceColors = append(sequences[this.TargetSequence].SequenceColors, newColor.Color)

	if debug {
		fmt.Printf("%d: RGB Adding colors are now %+v\n", this.TargetSequence, sequences[this.TargetSequence].SequenceColors)
	}

	this.ShowRGBColorPicker = true

	// We call ShowRGBColorPicker here so the selections will flash as you press them.
	ShowRGBColorPicker(*sequences[this.TargetSequence], eventsForLaunchpad, guiButtons, commandChannels)

}

func selectScannerColor(sequences []*common.Sequence, X int, Y int, this *CurrentState, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command, updateChannels []chan common.Sequence, dmxController *ft232.DMXController, fixturesConfig *fixture.Fixtures) {

	if debug {
		fmt.Printf("Set Scanner Color X:%d Y:%d\n", X, Y)
	}

	this.ScannerColor = X

	// Set the scanner color for this sequence.
	cmd := common.Command{
		Action: common.UpdateScannerColor,
		Args: []common.Arg{
			{Name: "SelectedColor", Value: this.ScannerColor},
			{Name: "SelectedFixture", Value: this.SelectedFixture},
		},
	}
	common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

	// Add the selected color to the sequence.
	if this.ScannerChaser[this.SelectedSequence] && this.SelectedType == "scanner" {
		cmd := common.Command{
			Action: common.UpdateScannerColor,
			Args: []common.Arg{
				{Name: "SelectedColor", Value: this.ScannerColor},
				{Name: "SelectedFixture", Value: this.SelectedFixture},
			},
		}
		common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)
	}

	this.EditScannerColorsMode = true

	// Sequence colors are calculated in the sequence thread so get an upto date copy of the sequence.
	sequences[this.SelectedSequence] = common.RefreshSequence(this.SelectedSequence, commandChannels, updateChannels)

	// If the sequence isn't running this will force a single color DMX message.
	fixture.MapFixturesColorOnly(this.SelectedSequence, this.SelectedFixture, this.ScannerColor, dmxController, fixturesConfig, this.DmxInterfacePresent)

	// Clear the pattern function keys
	common.ClearSelectedRowOfButtons(this.SelectedSequence, eventsForLaunchpad, guiButtons)

	// Update the new scanner colors in the labels.
	showStatusBars(this, sequences, eventsForLaunchpad, guiButtons)

	// We call ShowScannerColorSelectionButtons here so the selections will flash as you press them.
	ShowScannerColorSelectionButtons(*sequences[this.SelectedSequence], this, eventsForLaunchpad, fixturesConfig, guiButtons)

}

func FindCurrentColor(X int, Y int, targetSequence common.Sequence) color.RGBA {

	if debug {
		fmt.Printf("FindCurrentColor\n")
	}

	for _, availableColor := range targetSequence.RGBAvailableColors {
		if availableColor.X == X && availableColor.Y == Y {
			return availableColor.Color
		}
	}

	return color.RGBA{}
}

// For the given sequence show the available sequence colors on the relevant buttons.
// With the new color picker there can be 24 colors displayed.
// ShowRGBColorPicker operates on the sequence.RGBAvailableColors which is an array of type []common.StaticColorButton
// the targetSequence .CurrentColors selects which colors are selected.
// Returns nothing, simply displays the available colors on the buttons.
func ShowRGBColorPicker(targetSequence common.Sequence, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command) {

	if debug {
		fmt.Printf("Color Picker - Show Color Selection Buttons\n")
	}

	common.HideSequence(0, commandChannels)
	common.HideSequence(1, commandChannels)
	common.HideSequence(1, commandChannels)

	for myFixtureNumber, lamp := range targetSequence.RGBAvailableColors {

		lamp.Flash = false

		// Check if we need to flash this button.
		for index, availableColor := range targetSequence.RGBAvailableColors {
			for _, sequenceColor := range targetSequence.SequenceColors {
				if availableColor.Color == sequenceColor {
					if myFixtureNumber == index {
						if debug {
							fmt.Printf("myFixtureNumber %d   current color %+v\n", myFixtureNumber, sequenceColor)
						}
						lamp.Flash = true
					}
				}
			}
		}
		if lamp.Flash {
			Black := colors.Black
			if debug {
				fmt.Printf("FLASH myFixtureNumber X:%d Y:%d Color %+v \n", lamp.X, lamp.Y, lamp.Color)
			}
			common.FlashLight(common.Button{X: lamp.X, Y: lamp.Y}, lamp.Color, Black, eventsForLaunchpad, guiButtons)
		} else {
			common.LightLamp(common.Button{X: lamp.X, Y: lamp.Y}, lamp.Color, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)
		}
		common.LabelButton(lamp.X, lamp.Y, lamp.Name, guiButtons)

		time.Sleep(10 * time.Millisecond)
	}
}
