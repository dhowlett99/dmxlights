// Copyright (C) 2022, 2023 , 2024 dhowlett99.
// This is button processor, used by the launchpad and gui interfaces.
// This file processes the select buttons and controls their actions.
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

	"github.com/dhowlett99/dmxlights/pkg/colors"
	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
)

func selectSequence(sequences []*common.Sequence, Y int, this *CurrentState, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command) {

	this.SelectedSequence = Y
	this.SelectedType = sequences[this.SelectedSequence].Type

	if this.ScannerChaser[this.SelectedSequence] {
		this.EditWhichStaticSequence = this.ChaserSequenceNumber
	} else {
		this.EditWhichStaticSequence = this.SelectedSequence
	}

	if debug {
		fmt.Printf("Select Sequence %d Type %s\n", this.SelectedSequence, this.SelectedType)
	}

	// // If we're in shutter chase mode
	if this.SelectedMode[this.SelectedSequence] == CHASER_FUNCTION || this.SelectedMode[this.SelectedSequence] == CHASER_DISPLAY {
		this.TargetSequence = this.ChaserSequenceNumber
	} else {
		this.TargetSequence = this.SelectedSequence
	}

	deFocusAllSwitches(this, sequences, commandChannels)
	HandleSelect(sequences, this, eventsForLaunchpad, commandChannels, guiButtons)

	this.ShowRGBColorPicker = false
	this.EditGoboSelectionMode = false
	this.DisplayChaserShortCut = false

}

// For the given sequence show the available scanner selection colors on the relevant buttons.
func ShowSelectFixtureButtons(targetSequence common.Sequence, displaySequence int, this *CurrentState, eventsForLaunchpad chan common.ALight, action string, guiButtons chan common.ALight) int {

	if debug {
		fmt.Printf("Sequence %d Show Fixture Selection Buttons on the way to %s\n", this.SelectedSequence, action)
	}

	for fixtureNumber, fixture := range targetSequence.ScannersAvailable {

		if debug {
			fmt.Printf("Fixture %+v\n", fixture)
		}
		if fixtureNumber == this.SelectedFixture {
			fixture.Flash = true
			this.SelectedFixture = fixtureNumber
		}
		if fixture.Flash {
			common.FlashLight(common.Button{X: fixtureNumber, Y: displaySequence}, fixture.Color, colors.White, eventsForLaunchpad, guiButtons)
		} else {
			common.LightLamp(common.Button{X: fixtureNumber, Y: displaySequence}, fixture.Color, targetSequence.Master, eventsForLaunchpad, guiButtons)
		}
		common.LabelButton(fixtureNumber, displaySequence, fixture.Label, guiButtons)
	}
	if debug {
		fmt.Printf("Selected Fixture is %d\n", this.SelectedFixture)
	}
	return this.SelectedFixture
}

// ShowGoboSelectionButtons puts up a set of red buttons used to select a fixture.
func ShowGoboSelectionButtons(sequence common.Sequence, this *CurrentState, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight) {

	if debug {
		fmt.Printf("Sequence %d Show Gobo Selection Buttons\n", this.SelectedSequence)
	}
	// Check if we need to flash this button.
	for goboNumber, gobo := range sequence.ScannerAvailableGobos[this.SelectedFixture+1] {

		if gobo.Number > 8 {
			return // We only have 8 buttons so we can't select from any more.
		}
		if gobo.Number == sequence.ScannerGobo[this.SelectedFixture] {
			gobo.Flash = true
		}
		if debug {
			fmt.Printf("goboNumber %d   current gobo %d  flash gobo %t\n", goboNumber, sequence.ScannerGobo, gobo.Flash)
		}
		if gobo.Flash {
			Black := colors.Black
			common.FlashLight(common.Button{X: goboNumber, Y: this.SelectedSequence}, gobo.Color, Black, eventsForLaunchpad, guiButtons)
		} else {
			common.LightLamp(common.Button{X: goboNumber, Y: this.SelectedSequence}, gobo.Color, sequence.Master, eventsForLaunchpad, guiButtons)
		}
		goboName := common.FormatLabel(gobo.Name)
		common.LabelButton(goboNumber, this.SelectedSequence, goboName, guiButtons)
	}
}

// For the given sequence show the available scanner selection colors on the relevant buttons.
func ShowScannerColorSelectionButtons(sequence common.Sequence, this *CurrentState, eventsForLaunchpad chan common.ALight, fixtures *fixture.Fixtures, guiButtons chan common.ALight) error {

	if debug {
		fmt.Printf("Show Scanner Color Selection Buttons,  Sequence is %d  fixture is %d   color is %d \n", this.SelectedSequence, this.SelectedFixture, sequence.ScannerColor[this.SelectedFixture])
	}

	// if there are no colors available for this fixture turn everything off and print an error.
	if sequence.ScannerAvailableColors[this.SelectedFixture+1] == nil {

		// Turn off the color edit mode.
		this.ShowRGBColorPicker = false
		// And since we seem to be using two flags for the same thing, turn this off too.
		this.Functions[this.SelectedSequence][common.Function5_Color].State = false

		for _, fixture := range fixtures.Fixtures {
			if fixture.Group == this.SelectedSequence+1 {
				common.LightLamp(common.Button{X: fixture.Number - 1, Y: this.SelectedSequence}, colors.White, sequence.Master, eventsForLaunchpad, guiButtons)
			}
		}
		if this.GUI {
			displayErrorPopUp(this.MyWindow, "no colors available for this fixture")
		}

		return fmt.Errorf("error: no colors available for fixture number %d", this.SelectedFixture+1)
	}

	// selected fixture is +1 here because the fixtures in the yaml config file start with 1 not 0.
	for fixtureNumber, lamp := range sequence.ScannerAvailableColors[this.SelectedFixture+1] {

		if debug {
			fmt.Printf("Lamp %+v\n", lamp)
		}
		if fixtureNumber == sequence.ScannerColor[this.SelectedFixture] {
			lamp.Flash = true
		}

		if lamp.Flash {
			Black := colors.Black
			common.FlashLight(common.Button{X: fixtureNumber, Y: this.SelectedSequence}, lamp.Color, Black, eventsForLaunchpad, guiButtons)
		} else {
			common.LightLamp(common.Button{X: fixtureNumber, Y: this.SelectedSequence}, lamp.Color, sequence.Master, eventsForLaunchpad, guiButtons)
		}
		// Remove any labels.
		common.LabelButton(fixtureNumber, this.SelectedSequence, "", guiButtons)
	}
	return nil
}

// For the given sequence show the available patterns on the relevant buttons.
// mySequenceDisplayNumber is the sequence whos buttons you want the pattern selection to show on.
// master is the master brightness for the same buttons.
// this.TargetSequence - is the squence you are updating the pattern, this could be different in the case
// of scanner shutter chaser sequence which doesn't have it's own buttons.
func ShowPatternSelectionButtons(sequence *common.Sequence, master int, targetSequence common.Sequence, displaySequence int, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight) {

	if debug {
		fmt.Printf("Sequence Name %s Type %s  Label %s\n", targetSequence.Name, targetSequence.Type, targetSequence.Label)
		for _, pattern := range sequence.RGBAvailablePatterns {
			fmt.Printf("Found a pattern called %s\n", pattern.Name)
		}
	}

	if targetSequence.Type == "rgb" {
		for _, pattern := range sequence.RGBAvailablePatterns {
			if debug {
				fmt.Printf("pattern is %s\n", pattern.Name)
			}
			if pattern.Number == targetSequence.SelectedPattern {
				common.FlashLight(common.Button{X: pattern.Number, Y: displaySequence}, colors.White, colors.LightBlue, eventsForLaunchpad, guiButtons)
			} else {
				common.LightLamp(common.Button{X: pattern.Number, Y: displaySequence}, colors.LightBlue, master, eventsForLaunchpad, guiButtons)
			}
			common.LabelButton(pattern.Number, displaySequence, pattern.Label, guiButtons)
		}
		return
	}

	if targetSequence.Type == "scanner" {
		for _, pattern := range targetSequence.ScannerAvailablePatterns {
			if pattern.Number == targetSequence.SelectedPattern {
				common.FlashLight(common.Button{X: pattern.Number, Y: displaySequence}, colors.White, colors.LightBlue, eventsForLaunchpad, guiButtons)
			} else {
				common.LightLamp(common.Button{X: pattern.Number, Y: displaySequence}, colors.LightBlue, master, eventsForLaunchpad, guiButtons)
			}
			common.LabelButton(pattern.Number, displaySequence, pattern.Label, guiButtons)
		}
		return
	}
}

func lightSelectedButton(eventsForLauchpad chan common.ALight, guiButtons chan common.ALight, this *CurrentState) {

	NumberOfSelectableSequences := 4
	// 2 x RGB (FOH & Uplighters) sequences,
	// 1 x scanner sequence (chaser sequence shares its button with the scanner sequence),
	// 1 x switch sequence.

	if debug {
		fmt.Printf("SequenceSelect\n")
	}

	if this.SelectedSequence > NumberOfSelectableSequences-1 {
		return
	}

	// Turn off all sequence lights.
	for seq := 0; seq < NumberOfSelectableSequences; seq++ {
		common.LightLamp(common.Button{X: 8, Y: seq}, colors.Cyan, common.MAX_DMX_BRIGHTNESS, eventsForLauchpad, guiButtons)
	}

	// Provided we're not the switch sequence number turn on the selected lamp.
	if this.SelectedSequence != this.SwitchSequenceNumber {
		// Turn on the correct sequence select number.
		if this.SelectedType == "scanner" && this.ScannerChaser[this.SelectedSequence] &&
			(this.SelectedMode[this.SelectedSequence] == CHASER_FUNCTION || this.SelectedMode[this.SelectedSequence] == CHASER_DISPLAY) {
			// If we are in shutter chaser mode, light the lamp yellow.
			common.LightLamp(common.Button{X: 8, Y: this.SelectedSequence}, colors.Magenta, common.MAX_DMX_BRIGHTNESS, eventsForLauchpad, guiButtons)
		} else {
			// Now turn pink the selected sequence select light.
			common.LightLamp(common.Button{X: 8, Y: this.SelectedSequence}, colors.Magenta, common.MAX_DMX_BRIGHTNESS, eventsForLauchpad, guiButtons)
		}
	}
}

// For the given sequence clear the available this.Patterns on the relevant buttons.
func ClearPatternSelectionButtons(mySequenceNumber int, sequence common.Sequence, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight) {
	// Check if we need to flash this button.
	for myFixtureNumber := 0; myFixtureNumber < 4; myFixtureNumber++ {
		common.LightLamp(common.Button{X: myFixtureNumber, Y: mySequenceNumber}, colors.Black, sequence.Master, eventsForLaunchpad, guiButtons)
	}
}
