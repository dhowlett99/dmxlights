// Copyright (C) 2022, 2023 , 2024 dhowlett99.
// This is button processor, used by the launchpad and gui interfaces.
// This file processes static scene buttons and controls their actions.
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

	"github.com/dhowlett99/dmxlights/pkg/colors"
	"github.com/dhowlett99/dmxlights/pkg/common"
)

func selectStaticFixture(sequences []*common.Sequence, X int, Y int, this *CurrentState, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command) {

	this.TargetSequence = this.EditWhichStaticSequence
	this.DisplaySequence = Y

	if debug {
		fmt.Printf("EditWhichStaticSequence %d\n", this.EditWhichStaticSequence)
		fmt.Printf("TargetSequence %d\n", this.TargetSequence)
		fmt.Printf("DisplaySequence %d\n", this.DisplaySequence)
	}

	// Save the selected fixture number.
	this.SelectedStaticFixtureNumber = X

	// Reset Clear pressed flag so we can clear next selection
	this.ClearPressed[this.TargetSequence] = false

	// The current color is help in our local copy.
	color := sequences[this.TargetSequence].StaticColors[X].Color
	if color == colors.EmptyColor {
		color = FindCurrentColor(this.SelectedStaticFixtureNumber, this.SelectedSequence, *sequences[this.TargetSequence])
	}

	if debug {
		fmt.Printf("Sequence %d Fixture %d Setting Current Color as %+v\n", this.SelectedSequence, this.SelectedStaticFixtureNumber, color)
	}

	// We call ShowRGBColorPicker so you can choose the static color for this fixture.
	ShowRGBColorPicker(*sequences[this.TargetSequence], eventsForLaunchpad, guiButtons, commandChannels)

	// Switch the mode so we know we are picking a static color from the color picker.
	this.ShowStaticColorPicker = true
	this.Static[this.TargetSequence] = true

}

func selectStaticColor(sequences []*common.Sequence, X int, Y int, this *CurrentState, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command) {

	this.TargetSequence = this.EditWhichStaticSequence
	this.DisplaySequence = this.SelectedSequence

	if debug {
		fmt.Printf("Selected Mode is %s\n", printMode(this.SelectedMode[this.TargetSequence]))
		fmt.Printf("EditWhichStaticSequence %d\n", this.EditWhichStaticSequence)
		fmt.Printf("ShowStaticColorPicker %t\n", this.ShowStaticColorPicker)
		fmt.Printf("TargetSequence %d\n", this.TargetSequence)
		fmt.Printf("DisplaySequence %d\n", this.DisplaySequence)
	}

	// Find the color from the button pressed.
	color := FindCurrentColor(X, Y, *sequences[this.TargetSequence])

	if debug {
		fmt.Printf("Selected Static Color for X %d  Y %d to Color %+v\n", this.SelectedStaticFixtureNumber, Y, color)
	}

	// Set our local copy of the color.
	sequences[this.TargetSequence].StaticColors[this.SelectedStaticFixtureNumber].Color = color

	// Save the color in local copy of the static color button
	this.StaticButtons[this.SelectedStaticFixtureNumber].Color = color

	// Tell the sequence about the new color and where we are in the
	// color cycle.

	if this.SelectAllStaticFixtures {
		// Set the same static color for all.
		cmd := common.Command{
			Action: common.UpdateAllStaticColor,
			Args: []common.Arg{
				{Name: "Static", Value: true},
				{Name: "StaticLampFlash", Value: false},
				{Name: "SelectedColor", Value: sequences[this.TargetSequence].StaticColors[this.SelectedStaticFixtureNumber].SelectedColor},
				{Name: "StaticColor", Value: sequences[this.TargetSequence].StaticColors[this.SelectedStaticFixtureNumber].Color},
			},
		}
		common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

		this.SelectAllStaticFixtures = false

	} else {
		// Set a static color for an individual fixture.
		cmd := common.Command{
			Action: common.UpdateStaticColor,
			Args: []common.Arg{
				{Name: "Static", Value: true},
				{Name: "FixtureNumber", Value: this.SelectedStaticFixtureNumber},
				{Name: "StaticLampFlash", Value: true},
				{Name: "SelectedColor", Value: sequences[this.TargetSequence].StaticColors[this.SelectedStaticFixtureNumber].SelectedColor},
				{Name: "StaticColor", Value: sequences[this.TargetSequence].StaticColors[this.SelectedStaticFixtureNumber].Color},
			},
		}
		common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

	}

	this.LastStaticColorButtonX = this.SelectedStaticFixtureNumber
	this.LastStaticColorButtonY = Y

	// Hide the sequence.
	common.HideSequence(this.TargetSequence, commandChannels)

	// We call ShowRGBColorPicker so you can see which static color has been selected for this fixture.
	ShowRGBColorPicker(*sequences[this.TargetSequence], eventsForLaunchpad, guiButtons, commandChannels)

	// Set the first pressed for only this fixture and cancel any others
	for x := 0; x < 8; x++ {
		sequences[this.TargetSequence].StaticColors[x].FirstPress = false
	}
	sequences[this.TargetSequence].StaticColors[this.SelectedStaticFixtureNumber].FirstPress = true

	// Remove the color picker and reveal the sequence.
	removeColorPicker(this, sequences, eventsForLaunchpad, guiButtons, commandChannels)

	// Show static colors.
	cmd := common.Command{
		Action: common.UpdateStatic,
		Args: []common.Arg{
			{Name: "Static", Value: true},
		},
	}
	common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)
	common.RevealSequence(this.TargetSequence, commandChannels)

	// Update the labels.
	showStatusBars(this, sequences, eventsForLaunchpad, guiButtons)

	// Switch off the color picker.
	this.ShowStaticColorPicker = false

}

func updateStaticLamp(selectedSequence int, staticColorButtons common.StaticColorButton, commandChannels []chan common.Command) {

	// Static is set to true in the functions and this key is set to
	// the selected color.
	cmd := common.Command{
		Action: common.UpdateStaticColor,
		Args: []common.Arg{
			{Name: "Static", Value: true},
			{Name: "StaticLamp", Value: staticColorButtons.X},
			{Name: "StaticLampFlash", Value: false},
			{Name: "SelectedColor", Value: staticColorButtons.SelectedColor},
			{Name: "StaticColor", Value: color.RGBA{R: staticColorButtons.Color.R, G: staticColorButtons.Color.G, B: staticColorButtons.Color.B}},
		},
	}
	common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

}
