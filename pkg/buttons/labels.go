// Copyright (C) 2022, 2023 , 2024 dhowlett99.
// This is button processor, used by the launchpad and gui interfaces.
// This file controls which labels are shown at the top and bottom of the window.
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
	"strconv"
	"strings"

	"github.com/dhowlett99/dmxlights/pkg/colors"
	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/labels"
)

func showStatusBars(this *CurrentState, sequences []*common.Sequence, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight) {

	debug := false

	if debug {
		fmt.Printf("showStatusBar for sequence %d Type %s\n", this.SelectedSequence, this.SelectedType)
	}

	sensitivity := common.FindSensitivity(this.SoundGain)
	common.UpdateStatusBar(fmt.Sprintf("Sensitivity %02d", sensitivity), "sensitivity", false, guiButtons)

	// Make sure modes are setup.
	if this.SelectedType == "scanner" && this.ScannerChaser[this.SelectedSequence] &&
		(this.SelectedMode[this.SelectedSequence] == CHASER_FUNCTION || this.SelectedMode[this.SelectedSequence] == CHASER_DISPLAY) {
		this.TargetSequence = this.ChaserSequenceNumber
		this.DisplaySequence = this.SelectedSequence
	} else {
		this.TargetSequence = this.SelectedSequence
		this.DisplaySequence = this.SelectedSequence
	}

	sequence := getSelectedSequenceNumber(this.TargetSequence, this.SelectedType, this.SelectedSwitch)
	common.UpdateStatusBar(fmt.Sprintf("Master %02d", this.MasterBrightness[sequence]), "master", false, guiButtons)

	if debug {
		fmt.Printf("Target Sequence %d Mode %s Type %s\n", this.TargetSequence, printMode(this.SelectedMode[this.TargetSequence]), sequences[this.TargetSequence].Type)
		fmt.Printf("Display Sequence %d Mode %s Type %s\n", this.DisplaySequence, printMode(this.SelectedMode[this.DisplaySequence]), sequences[this.DisplaySequence].Type)
	}

	// Update status bar.
	updateBottomStatusBar(0, 0, common.DisplayAll, common.Display, sequences, this, nil, eventsForLaunchpad, guiButtons)

	showTopLabels(this, eventsForLaunchpad, guiButtons)

	control := getControl(sequences, this)
	common.UpdateColorDisplay(control, guiButtons)

	// Hide the color editing buttons.
	common.UpdateStatusBar(fmt.Sprintf("Tilt %02d", this.OffsetTilt), "tilt", false, guiButtons)
	common.UpdateStatusBar("        ", "red", false, guiButtons)
	common.UpdateStatusBar("        ", "green", false, guiButtons)
	common.UpdateStatusBar(fmt.Sprintf("Pan %02d", this.OffsetPan), "pan", false, guiButtons)

}

// getControl looks at the sequences and switches and find the current colors.
// relies on sequences being an upto date copy of the sequences.
func getControl(sequences []*common.Sequence, this *CurrentState) common.ColorDisplayControl {

	// Sequence is a switch sequence, so check switch actions for appropriate colors.
	if sequences[this.TargetSequence].Type == "switch" {
		return getSwitchColors(sequences, this)
	}

	// Sequence with RGB or Scanner.
	if !this.Static[this.TargetSequence] {
		// Update the color display for the sequence.
		return common.GetColorList(sequences[this.TargetSequence].SequenceColors)
	}

	// Sequence in static mode.
	if this.Static[this.TargetSequence] {
		// Use static colors for color display.
		staticColors := []color.RGBA{}
		for buttonNumber, button := range sequences[this.TargetSequence].StaticColors {
			if buttonNumber > 7 { // Only copy the first eight fixtures.
				break
			}
			staticColors = append(staticColors, button.Color)
		}
		return common.GetColorList(staticColors)
	}

	// Nothig found, return an blank color display.
	return common.ColorDisplayControl{}
}

func getSwitchColors(sequences []*common.Sequence, this *CurrentState) common.ColorDisplayControl {

	var control common.ColorDisplayControl

	// Look inside the switch configuration to see what type of action has been set.
	for swiTchNumber, swiTch := range sequences[this.TargetSequence].Switches {
		if swiTchNumber == this.SelectedSwitch {
			// Look inside switch state.
			for stateNumber, state := range swiTch.States {
				if stateNumber == this.SwitchPosition[this.SelectedSwitch] {

					// Look at the settings.
					// Settings override any actions so check them first.
					if len(state.Settings) > 0 {
						control := checkStateSettings(state)
						empty := common.ColorDisplayControl{}
						if control != empty {
							return control
						}
					}
					// Look at the actions.
					if len(state.Actions) > 0 {
						return checkStateActions(state)
					}
				}
			}
		}
	}
	return control
}

func checkStateSettings(state common.State) common.ColorDisplayControl {

	if debug {
		fmt.Printf("checkStateSettings\n")
	}

	var red int
	var green int
	var blue int

	var foundRGB bool
	var colorFound color.RGBA
	var colorsList []color.RGBA

	for _, setting := range state.Settings {

		// If a color wheel is present.
		if setting.Channel == "Color" {
			foundRGB = false
			// TODO work out what color has been selected.
		}
		// If settins for RGB channels exist
		// construct a color from them.
		if strings.Contains(setting.Name, "Red") ||
			strings.Contains(setting.Name, "red") ||
			strings.Contains(setting.Channel, "Red") ||
			strings.Contains(setting.Channel, "red") {

			foundRGB = true

			red, _ = strconv.Atoi(setting.FixtureValue)
		}
		if strings.Contains(setting.Name, "Green") ||
			strings.Contains(setting.Name, "green") ||
			strings.Contains(setting.Channel, "Green") ||
			strings.Contains(setting.Channel, "green") {

			foundRGB = true
			green, _ = strconv.Atoi(setting.FixtureValue)
		}
		if strings.Contains(setting.Name, "Blue") ||
			strings.Contains(setting.Name, "blue") ||
			strings.Contains(setting.Channel, "Blue") ||
			strings.Contains(setting.Channel, "blue") {

			foundRGB = true
			blue, _ = strconv.Atoi(setting.FixtureValue)
		}

		if foundRGB {
			foundRGB = false
			colorFound.R = uint8(red)
			colorFound.G = uint8(green)
			colorFound.B = uint8(blue)
			colorFound.A = 255
			colorsList = append(colorsList, colorFound)
		}
	}
	return common.GetColorList(colorsList)
}

func checkStateActions(state common.State) common.ColorDisplayControl {

	if debug {
		fmt.Printf("checkStateActions\n")
	}

	var control common.ColorDisplayControl

	// Look inside state action.
	switch state.Actions[0].Mode {

	case "Off":
		// Return the blank control struct.
		return control
	case "Static":
		control := common.GetColorListByNames(state.Actions[0].Colors)
		return control
	case "Control":
		// Because shows contain an unknow number of colors return the blank control struct.
		return control
	case "Chase":
		control := common.GetColorListByNames(state.Actions[0].Colors)
		return control
	default:
		// Return the blank control struct.
		return control
	}
}

func showTopLabels(this *CurrentState, eventsForLauchpad chan common.ALight, guiButtons chan common.ALight) {

	debug := false

	type topButton struct {
		Label string
		Color color.RGBA
	}

	// Storage for the rgb labels on the top row.
	var guiTopRGBButtons [8]topButton
	guiTopRGBButtons[0] = topButton{Label: labels.GetLabel(this.Labels, "Clear", "Clear"), Color: colors.Magenta}
	guiTopRGBButtons[1] = topButton{Label: labels.GetLabel(this.Labels, "Colors", "Red"), Color: colors.Red}
	guiTopRGBButtons[2] = topButton{Label: labels.GetLabel(this.Labels, "Colors", "Green"), Color: colors.Green}
	guiTopRGBButtons[3] = topButton{Label: labels.GetLabel(this.Labels, "Colors", "Blue"), Color: colors.Blue}
	guiTopRGBButtons[4] = topButton{Label: labels.GetLabel(this.Labels, "Sensitivity", "Decrease"), Color: colors.Cyan}
	guiTopRGBButtons[5] = topButton{Label: labels.GetLabel(this.Labels, "Sensitivity", "Increase"), Color: colors.Cyan}
	guiTopRGBButtons[6] = topButton{Label: labels.GetLabel(this.Labels, "Master", "Decrease"), Color: colors.Cyan}
	guiTopRGBButtons[7] = topButton{Label: labels.GetLabel(this.Labels, "Master", "Increase"), Color: colors.Cyan}

	// Storage for the scanner labels on the Top row.
	var guiTopScannerButtons [8]topButton
	guiTopScannerButtons[0] = topButton{Label: labels.GetLabel(this.Labels, "Clear", "Clear ^"), Color: colors.White}
	guiTopScannerButtons[1] = topButton{Label: labels.GetLabel(this.Labels, "Scanner Buttons", "Down"), Color: colors.White}
	guiTopScannerButtons[2] = topButton{Label: labels.GetLabel(this.Labels, "Scanner Buttons", "Left"), Color: colors.White}
	guiTopScannerButtons[3] = topButton{Label: labels.GetLabel(this.Labels, "Scanner Buttons", "Right"), Color: colors.White}
	guiTopScannerButtons[4] = topButton{Label: labels.GetLabel(this.Labels, "Sensitivity", "Decrease"), Color: colors.Cyan}
	guiTopScannerButtons[5] = topButton{Label: labels.GetLabel(this.Labels, "Sensitivity", "Increase"), Color: colors.Cyan}
	guiTopScannerButtons[6] = topButton{Label: labels.GetLabel(this.Labels, "Master", "Decrease"), Color: colors.Cyan}
	guiTopScannerButtons[7] = topButton{Label: labels.GetLabel(this.Labels, "Master", "Increase"), Color: colors.Cyan}

	// Storage for the switch labels on the top row.
	var guiTopSwitchButtons [8]topButton
	guiTopSwitchButtons[0] = topButton{Label: labels.GetLabel(this.Labels, "Clear", "Clear"), Color: colors.Magenta}
	guiTopSwitchButtons[1] = topButton{Label: labels.GetLabel(this.Labels, "Colors", "Red"), Color: colors.Red}
	guiTopSwitchButtons[2] = topButton{Label: labels.GetLabel(this.Labels, "Colors", "Green"), Color: colors.Green}
	guiTopSwitchButtons[3] = topButton{Label: labels.GetLabel(this.Labels, "Colors", "Blue"), Color: colors.Blue}
	guiTopSwitchButtons[4] = topButton{Label: labels.GetLabel(this.Labels, "Sensitivity", "Decrease"), Color: colors.Cyan}
	guiTopSwitchButtons[5] = topButton{Label: labels.GetLabel(this.Labels, "Sensitivity", "Increase"), Color: colors.Cyan}
	guiTopSwitchButtons[6] = topButton{Label: labels.GetLabel(this.Labels, "Master", "Decrease"), Color: colors.Cyan}
	guiTopSwitchButtons[7] = topButton{Label: labels.GetLabel(this.Labels, "Master", "Increase"), Color: colors.Cyan}

	//  The Top row of the Novation Launchpad.
	TopRow := -1

	if this.SelectedType == "rgb" {
		// Loop through the available functions for this sequence
		for index, button := range guiTopRGBButtons {
			if debug {
				fmt.Printf("rgb button %+v\n", button)
			}
			common.LightLamp(common.Button{X: index, Y: TopRow}, button.Color, common.MAX_DMX_BRIGHTNESS, eventsForLauchpad, guiButtons)
			common.LabelButton(index, TopRow, button.Label, guiButtons)
		}
	}
	if this.SelectedType == "scanner" {
		// Loop through the available functions for this sequence
		for index, button := range guiTopScannerButtons {
			if debug {
				fmt.Printf("scanner button %+v\n", button)
			}
			common.LightLamp(common.Button{X: index, Y: TopRow}, button.Color, common.MAX_DMX_BRIGHTNESS, eventsForLauchpad, guiButtons)
			common.LabelButton(index, TopRow, button.Label, guiButtons)
		}
	}
	if this.SelectedType == "switch" {
		// Loop through the available functions for this sequence
		for index, button := range guiTopSwitchButtons {
			if debug {
				fmt.Printf("switch button %+v\n", button)
			}
			common.LightLamp(common.Button{X: index, Y: TopRow}, button.Color, common.MAX_DMX_BRIGHTNESS, eventsForLauchpad, guiButtons)
			common.LabelButton(index, TopRow, button.Label, guiButtons)
		}
	}
}
