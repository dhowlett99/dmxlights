// Copyright (C) 2022, 2023, 2024, 2025 dhowlett99.
// This is status bar update code, used to update the speed, shift, size and fade labels.
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
)

func getAction(this *CurrentState) int {

	var action int

	// Position
	number := this.SelectedSwitch
	position := this.SwitchPosition[this.SelectedSwitch]
	overrides := *this.SwitchOverrides

	//if this.SelectedType == "rgb" && sequences[this.SelectedSequence].Label != "chaser" {
	if this.SelectedType == "rgb" {
		action = common.ActionRGB
		if debug {
			fmt.Printf("Action RGB\n")
		}
	}
	if this.SelectedType == "scanner" {
		action = common.ActionScanner
		if debug {
			fmt.Printf("Action Scanner\n")
		}
	}

	//if this.SelectedType == "rgb" && sequences[this.SelectedSequence].Label == "chaser" {
	if this.SelectedType == "rgb" {
		action = common.ActionChaser
		if debug {
			fmt.Printf("Action Chaser\n")
		}
	}
	if this.SelectedType == "switch" {
		if debug {
			fmt.Printf("Action Switch ")
		}
		switch overrides[number][position].Mode {

		case "Setting":

			action = common.ActionSwitchSetting
			if debug {
				fmt.Printf("Setting \n")
			}

		case "Off":
			action = common.ActionSwitchOff
			if debug {
				fmt.Printf("Off \n")
			}

		case "Static":
			action = common.ActionSwitchStatic
			if debug {
				fmt.Printf("Static \n")
			}

		case "Control":
			action = common.ActionSwitchControl
			if debug {
				fmt.Printf("Control \n")
			}

		case "Chase":
			action = common.ActionSwitchChaser
			if debug {
				fmt.Printf("Chase \n")
			}
		}
	}

	return action
}

func getType(selectedFixtureType string) int {

	var fixtureType int
	switch selectedFixtureType {

	case "rgb":
		fixtureType = common.RGB

	case "scanner":
		fixtureType = common.Scanner

	case "derby":
		fixtureType = common.Derby

	case "projector":
		fixtureType = common.Projector
	}

	return fixtureType
}

// updateStatusBar processes speed, shift, size and fade actions and all there variations, depending on fixture.
func updateStatusBar(X int, Y int, sub int, direction int, sequences []*common.Sequence, this *CurrentState, commandChannels []chan common.Command, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight) {

	if debug {
		fmt.Printf("DecideOnAction action \n")
	}

	action := getAction(this)
	fixtureType := getType(this.SelectedFixtureType)

	// If we're in shutter chase mode.
	this.TargetSequence = CheckType(this.SequenceType[this.SelectedSequence], this)

	if !(X == 0 || Y == 0) {
		buttonTouched(common.Button{X: X, Y: Y}, colors.White, colors.Cyan, eventsForLaunchpad, guiButtons)
	}

	switch action {

	// Common RGB Fixure.
	case common.ActionRGB:
		rgbFixture(sequences, sub, direction, this, commandChannels, guiButtons)
		labelButtons(this.ButtonConfig.RGBButtons, eventsForLaunchpad, guiButtons)

	// Scanner RGB Shutter Chaser. With Speed, Shift, Size & Fade.
	case common.ActionChaser:
		chaserFixture(sequences, sub, direction, this, commandChannels, guiButtons)
		labelButtons(this.ButtonConfig.ChaserButtons, eventsForLaunchpad, guiButtons)

	// Scanner with Speed, Shift, Size & Coordinates.
	case common.ActionScanner:
		scannerFixture(sequences, sub, direction, this, commandChannels, guiButtons)
		labelButtons(this.ButtonConfig.ScannerButtons, eventsForLaunchpad, guiButtons)

	// A Switch that has Settings could be a RGB Lamp, Scanner (TODO), Derby RGB with Rotate or a Projector with Color & Gobo Wheels.
	case common.ActionSwitchSetting:

		switch fixtureType {

		case common.RGB:
			rgbFixture(sequences, sub, direction, this, commandChannels, guiButtons)
			labelButtons(this.ButtonConfig.RGBButtons, eventsForLaunchpad, guiButtons)

		case common.Scanner:
			scannerFixture(sequences, sub, direction, this, commandChannels, guiButtons)
			labelButtons(this.ButtonConfig.ScannerButtons, eventsForLaunchpad, guiButtons)

		case common.Derby:
			derbyFixture(sequences, sub, direction, this, commandChannels, guiButtons)
			labelButtons(this.ButtonConfig.DerbyButtons, eventsForLaunchpad, guiButtons)

		case common.Projector:
			projectorFixture(sub, direction, this, commandChannels, guiButtons)
			labelButtons(this.ButtonConfig.ProjectorButtons, eventsForLaunchpad, guiButtons)
		}

	case common.ActionSwitchOff:
		blankFixture(sub, guiButtons)
		labelButtons(this.ButtonConfig.BlankButtons, eventsForLaunchpad, guiButtons)

	case common.ActionSwitchStatic:
		staticFixture(sub, direction, this, commandChannels, guiButtons)

		switch fixtureType {
		case common.RGB:
			rgbFixture(sequences, sub, direction, this, commandChannels, guiButtons)
			labelButtons(this.ButtonConfig.RGBButtons, eventsForLaunchpad, guiButtons)

		case common.Scanner:
			scannerFixture(sequences, sub, direction, this, commandChannels, guiButtons)
			labelButtons(this.ButtonConfig.ScannerButtons, eventsForLaunchpad, guiButtons)

		case common.Derby:
			derbyFixture(sequences, sub, direction, this, commandChannels, guiButtons)
			labelButtons(this.ButtonConfig.DerbyButtons, eventsForLaunchpad, guiButtons)

		case common.Projector:
			projectorFixture(sub, direction, this, commandChannels, guiButtons)
			labelButtons(this.ButtonConfig.ProjectorButtons, eventsForLaunchpad, guiButtons)
		}

	case common.ActionSwitchControl:
		programFixture(sub, direction, this, commandChannels, guiButtons)

	// Control a fixture that has a mini sequencer in chase mode.
	// The chaser can control a RGB lamp, a scanner (TODO), a derby with RGB and rotate, a projector with gobo and color wheels
	case common.ActionSwitchChaser:

		switch fixtureType {

		case common.RGB:
			rgbFixture(sequences, sub, direction, this, commandChannels, guiButtons)
			labelButtons(this.ButtonConfig.RGBButtons, eventsForLaunchpad, guiButtons)
		case common.Scanner:
			scannerFixture(sequences, sub, direction, this, commandChannels, guiButtons)
			labelButtons(this.ButtonConfig.ScannerButtons, eventsForLaunchpad, guiButtons)

		case common.Derby:
			derbyFixture(sequences, sub, direction, this, commandChannels, guiButtons)
			labelButtons(this.ButtonConfig.DerbyButtons, eventsForLaunchpad, guiButtons)

		case common.Projector:
			projectorFixture(sub, direction, this, commandChannels, guiButtons)
			labelButtons(this.ButtonConfig.ProjectorButtons, eventsForLaunchpad, guiButtons)

		}

	}

	// Pull overrides.
	overrides := *this.SwitchOverrides
	// Show the color display.
	control := common.GetColorListByNames(overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].AvailableColors)
	if debug {
		fmt.Printf("Control %+v\n", control)
	}
	common.UpdateColorDisplay(control, guiButtons)

	// var control common.ColorDisplayControl
	// if !this.Static[this.TargetSequence] {
	// 	// Update the color display for the sequence.
	// 	control = common.GetColorList(sequenceColors)
	// } else {
	// 	// Use static colors for color display.
	// 	control = common.GetColorList(staticColors)
	// }
	// common.UpdateColorDisplay(control, guiButtons)

}

func rgbFixture(sequences []*common.Sequence, sub int, direction int, this *CurrentState, commandChannels []chan common.Command, guiButtons chan common.ALight) {

	switch sub {

	case common.DisplayAll:
		common.UpdateStatusBar(fmt.Sprintf("Speed %02d", this.Speed[this.SelectedSequence]), "speed", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Shift %02d", this.RGBShift[this.SelectedSequence]), "shift", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Size %02d", this.RGBSize[this.SelectedSequence]), "size", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Fade %02d", this.RGBFade[this.SelectedSequence]), "fade", false, guiButtons)

	case common.ChangeSpeed:
		switch direction {
		case common.Decrease:
			decreaseSpeed(sequences, this, commandChannels)
		case common.Increase:
			increaseSpeedRGB(sequences, this, commandChannels)
		}
		common.UpdateStatusBar(fmt.Sprintf("Speed %02d", this.Speed[this.SelectedSequence]), "speed", false, guiButtons)

	case common.ChangeShift:
		switch direction {
		case common.Decrease:
			decreaseShiftRGB(this, commandChannels)
		case common.Increase:
			increaseShiftRGB(this, commandChannels)
		}
		common.UpdateStatusBar(fmt.Sprintf("Shift %02d", this.RGBShift[this.SelectedSequence]), "shift", false, guiButtons)

	case common.ChangeSize:
		switch direction {
		case common.Decrease:
			decreaseSizeRGB(this, commandChannels)
		case common.Increase:
			increaseSizeRGB(this, commandChannels)
		}

		common.UpdateStatusBar(fmt.Sprintf("Size %02d", this.RGBSize[this.SelectedSequence]), "size", false, guiButtons)

	case common.ChangeFade:
		switch direction {
		case common.Decrease:
			decreaseFadeRGB(this, commandChannels)
		case common.Increase:
			increaseFadeRGB(this, commandChannels)
		}
		common.UpdateStatusBar(fmt.Sprintf("Fade %02d", this.RGBFade[this.SelectedSequence]), "fade", false, guiButtons)
	}
}

func chaserFixture(sequences []*common.Sequence, sub int, direction int, this *CurrentState, commandChannels []chan common.Command, guiButtons chan common.ALight) {
	switch sub {

	case common.DisplayAll:
		common.UpdateStatusBar(fmt.Sprintf("Speed %02d", this.Speed[this.TargetSequence]), "speed", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Shift %02d", this.RGBShift[this.TargetSequence]), "shift", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Size %02d", this.RGBSize[this.TargetSequence]), "size", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Fade %02d", this.RGBFade[this.TargetSequence]), "fade", false, guiButtons)

	case common.ChangeSpeed:
		switch direction {
		case common.Decrease:
			decreaseSpeed(sequences, this, commandChannels)
		case common.Increase:
			increaseSpeedRGB(sequences, this, commandChannels)
		}
		common.UpdateStatusBar(fmt.Sprintf("Speed %02d", this.Speed[this.TargetSequence]), "speed", false, guiButtons)

	case common.ChangeShift:
		switch direction {
		case common.Decrease:
			decreaseShiftRGB(this, commandChannels)
		case common.Increase:
			increaseShiftRGB(this, commandChannels)
		}
		common.UpdateStatusBar(fmt.Sprintf("Shift %02d", this.RGBShift[this.TargetSequence]), "shift", false, guiButtons)

	case common.ChangeSize:
		switch direction {
		case common.Decrease:
			decreaseSizeRGB(this, commandChannels)
		case common.Increase:
			increaseSizeRGB(this, commandChannels)
		}
		common.UpdateStatusBar(fmt.Sprintf("Size %02d", this.RGBSize[this.TargetSequence]), "size", false, guiButtons)

	case common.ChangeFade:
		switch direction {
		case common.Decrease:
			decreaseFadeRGB(this, commandChannels)
		case common.Increase:
			increaseFadeRGB(this, commandChannels)
		}
		common.UpdateStatusBar(fmt.Sprintf("Fade %02d", this.RGBFade[this.TargetSequence]), "fade", false, guiButtons)
	}
}

func scannerFixture(sequences []*common.Sequence, sub int, direction int, this *CurrentState, commandChannels []chan common.Command, guiButtons chan common.ALight) {

	switch sub {

	case common.DisplayAll:
		common.UpdateStatusBar(fmt.Sprintf("Speed %02d", this.Speed[this.TargetSequence]), "speed", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Shift %s", getScannerShiftLabel(this.ScannerShift[this.SelectedSequence])), "shift", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Size %02d", this.ScannerSize[this.TargetSequence]), "size", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Rotate Coord %s", getScannerCoordinatesLabel(this.ScannerCoordinates[this.TargetSequence])), "fade", false, guiButtons)

	case common.ChangeSpeed:
		switch direction {
		case common.Decrease:
			decreaseSpeed(sequences, this, commandChannels)
		case common.Increase:
			increaseSpeedScanner(this, commandChannels)
		}
		common.UpdateStatusBar(fmt.Sprintf("Speed %02d", this.Speed[this.TargetSequence]), "speed", false, guiButtons)

	case common.ChangeShift:
		switch direction {
		case common.Decrease:
			decreaseShiftScanner(this, commandChannels)
		case common.Increase:
			increaseShiftScanner(this, commandChannels)
		}
		common.UpdateStatusBar(fmt.Sprintf("Shift %s", getScannerShiftLabel(this.ScannerShift[this.SelectedSequence])), "shift", false, guiButtons)

	case common.ChangeSize:
		switch direction {
		case common.Decrease:
			decreaseSizeScanner(this, commandChannels)
		case common.Increase:
			increaseSizeScanner(this, commandChannels)
		}
		common.UpdateStatusBar(fmt.Sprintf("Scanner Size %02d", this.ScannerSize[this.TargetSequence]), "size", false, guiButtons)

	case common.ChangeFade:
		switch direction {
		case common.Decrease:
			decreaseFadeScannerNumberCoordinates(this, commandChannels)
		case common.Increase:
			increaseFadeScannerNumberCoordinates(this, commandChannels)
		}
		common.UpdateStatusBar(fmt.Sprintf("Rotate Coord %s", getScannerCoordinatesLabel(this.ScannerCoordinates[this.TargetSequence])), "fade", false, guiButtons)
	}
}

func derbyFixture(sequences []*common.Sequence, sub int, direction int, this *CurrentState, commandChannels []chan common.Command, guiButtons chan common.ALight) {

	color, colorName := getColor(this)
	rotate, rotateName := getRotate(this)
	gobo, goboName := getGobo(this)

	switch sub {

	// ShutterSpeed
	case common.ChangeSpeed:
		switch direction {
		case common.Decrease:
			decreaseSpeed(sequences, this, commandChannels)
		case common.Increase:
			increaseSpeedScanner(this, commandChannels)
		}
		common.UpdateStatusBar(fmt.Sprintf("Speed %02d", this.Speed[this.TargetSequence]), "speed", false, guiButtons)

	// RotateSpeed
	case common.ChangeShift:
		switch direction {
		case common.Decrease:
			decreaseRotateSpeed(this, commandChannels)
		case common.Increase:
			increaseRotateSpeed(this, commandChannels)
		}
		common.UpdateStatusBar(fmt.Sprintf("Rotate Speed %d:%s", rotate, rotateName), "shift", false, guiButtons)

		// ColorWheel
	case common.ChangeSize:
		switch direction {
		case common.Decrease:
			decreaseColor(this, commandChannels)
		case common.Increase:
			increaseColor(this, commandChannels)
		}
		common.UpdateStatusBar(fmt.Sprintf("Color %d:%s", color, colorName), "size", false, guiButtons)

	// GoboWheel
	case common.ChangeFade:
		switch direction {
		case common.Decrease:
			decreaseGobo(this, commandChannels)
		case common.Increase:
			increaseGobo(this, commandChannels)
		}
		common.UpdateStatusBar(fmt.Sprintf("Gobo %d:%s", gobo, goboName), "fade", false, guiButtons)
	}
}

func projectorFixture(sub int, direction int, this *CurrentState, commandChannels []chan common.Command, guiButtons chan common.ALight) {

	switch sub {

	case common.DisplayAll:
		color, colorName := getColor(this)
		rotate, rotateName := getRotate(this)
		gobo, goboName := getGobo(this)
		speed := getSpeed(this)
		common.UpdateStatusBar(fmt.Sprintf("Speed %02d", speed), "speed", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Rotate Speed %d:%s", rotate, rotateName), "shift", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Color %02d:%s", color, colorName), "size", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Gobo %02d:%s", gobo, goboName), "fade", false, guiButtons)

	// ShutterSpeed
	case common.ChangeSpeed:
		switch direction {
		case common.Decrease:
			decreaseOverrideSpeedRGB(this, commandChannels)
		case common.Increase:
			increaseOverrideSpeedRGB(this, commandChannels)
		}
		speed := getSpeed(this)
		common.UpdateStatusBar(fmt.Sprintf("Speed %02d", speed), "speed", false, guiButtons)

	// RotateSpeed.
	case common.ChangeShift:
		switch direction {
		case common.Decrease:
			decreaseRotateSpeed(this, commandChannels)
		case common.Increase:
			increaseRotateSpeed(this, commandChannels)
		}
		rotate, rotateName := getRotate(this)
		common.UpdateStatusBar(fmt.Sprintf("Rotate Speed %d:%s", rotate, rotateName), "shift", false, guiButtons)

	// ColorWheel
	case common.ChangeSize:
		switch direction {
		case common.Decrease:
			decreaseColor(this, commandChannels)
		case common.Increase:
			increaseColor(this, commandChannels)
		}
		color, colorName := getColor(this)
		common.UpdateStatusBar(fmt.Sprintf("Color %02d:%s", color, colorName), "size", false, guiButtons)

	// GoboWheel
	case common.ChangeFade:
		switch direction {
		case common.Decrease:
			decreaseGobo(this, commandChannels)
		case common.Increase:
			increaseGobo(this, commandChannels)
		}
		gobo, goboName := getGobo(this)
		common.UpdateStatusBar(fmt.Sprintf("Gobo %02d:%s", gobo, goboName), "fade", false, guiButtons)
	}
}

func programFixture(sub int, direction int, this *CurrentState, commandChannels []chan common.Command, guiButtons chan common.ALight) {

	switch sub {

	case common.DisplayAll:
		programSpeed, programSpeedName := getProgramSpeed(this)
		program, programName := getProgram(this)
		common.UpdateStatusBar(fmt.Sprintf("Speed %d:%s", programSpeed, programSpeedName), "speed", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Program %d:%s", program, programName), "shift", false, guiButtons)
		common.UpdateStatusBar("     ", "size", false, guiButtons)
		common.UpdateStatusBar("     ", "fade", false, guiButtons)

	// ProgramSpeed
	case common.ChangeSpeed:
		switch direction {
		case common.Decrease:
			decreaseProgramSpeed(this, commandChannels)
		case common.Increase:
			increaseProgramSpeed(this, commandChannels)
		}
		programSpeed, programSpeedName := getProgramSpeed(this)
		common.UpdateStatusBar(fmt.Sprintf("Speed %d:%s", programSpeed, programSpeedName), "speed", false, guiButtons)

	// Program
	case common.ChangeShift:
		switch direction {
		case common.Decrease:
			decreaseProgram(this, commandChannels)
		case common.Increase:
			increaseProgram(this, commandChannels)
		}
		program, programName := getProgram(this)
		common.UpdateStatusBar(fmt.Sprintf("Program %d:%s", program, programName), "shift", false, guiButtons)
	}
}

func blankFixture(sub int, guiButtons chan common.ALight) {

	switch sub {

	case common.DisplayAll:
		common.UpdateStatusBar("    ", "speed", false, guiButtons)
		common.UpdateStatusBar("    ", "shift", false, guiButtons)
		common.UpdateStatusBar("    ", "size", false, guiButtons)
		common.UpdateStatusBar("    ", "fade", false, guiButtons)

	case common.ChangeSpeed:
		common.UpdateStatusBar("    ", "speed", false, guiButtons)

	case common.ChangeShift:
		common.UpdateStatusBar("    ", "shift", false, guiButtons)

	case common.ChangeSize:
		common.UpdateStatusBar("    ", "size", false, guiButtons)

	case common.ChangeFade:
		common.UpdateStatusBar("    ", "fade", false, guiButtons)
	}
}

func staticFixture(sub int, direction int, this *CurrentState, commandChannels []chan common.Command, guiButtons chan common.ALight) {

	switch sub {

	case common.DisplayAll:
		rotate, rotateName := getRotate(this)
		color, colorName := getColor(this)
		common.UpdateStatusBar(fmt.Sprintf("Speed %02d", this.Speed[this.TargetSequence]), "speed", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Rotate Speed %d:%s", rotate, rotateName), "shift", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Color %d:%s", color, colorName), "size", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Fade %02d", this.RGBFade[this.TargetSequence]), "fade", false, guiButtons)

	case common.ChangeSpeed:
		switch direction {
		case common.Decrease:
			decreaseOverrideSpeedRGB(this, commandChannels)
		case common.Increase:
			increaseOverrideSpeedRGB(this, commandChannels)
		}
		common.UpdateStatusBar(fmt.Sprintf("Speed %02d", this.Speed[this.TargetSequence]), "speed", false, guiButtons)

	// RotateSpeed
	case common.ChangeShift:
		switch direction {
		case common.Decrease:
			decreaseRotateSpeed(this, commandChannels)
		case common.Increase:
			increaseRotateSpeed(this, commandChannels)
		}
		rotate, rotateName := getRotate(this)
		common.UpdateStatusBar(fmt.Sprintf("Rotate Speed %d:%s", rotate, rotateName), "shift", false, guiButtons)

	// ColorWheel
	case common.ChangeSize:
		switch direction {
		case common.Decrease:
			decreaseColor(this, commandChannels)
		case common.Increase:
			increaseSizeScanner(this, commandChannels)
		}
		color, colorName := getColor(this)
		common.UpdateStatusBar(fmt.Sprintf("Color %d:%s", color, colorName), "size", false, guiButtons)

	// GoboWheel
	case common.ChangeFade:
		switch direction {
		case common.Decrease:
			decreaseGobo(this, commandChannels)
		case common.Increase:
			increaseGobo(this, commandChannels)
		}
		common.UpdateStatusBar(fmt.Sprintf("Fade %02d", this.RGBFade[this.TargetSequence]), "fade", false, guiButtons)
	}
}

func getRotate(this *CurrentState) (int, string) {

	// Pull Overrides
	overrides := *this.SwitchOverrides

	// Position
	number := this.SelectedSwitch
	position := this.SwitchPosition[this.SelectedSwitch]

	// Rotate.
	rotate := overrides[number][position].Rotate
	availableRotates := overrides[number][position].RotateChannels
	// rotateName := overrides[number][position].RotateName
	//isRotateOverrideAble := overrides[number][position].IsRotateOverrideAble
	numberOfRotates := len(availableRotates)
	rotateName := "Unknown"
	if numberOfRotates > 0 && rotate <= numberOfRotates && rotate != -1 {
		if rotate > 0 {
			rotateName = availableRotates[rotate-1]
		}
	}

	return rotate, rotateName
}

func getColor(this *CurrentState) (int, string) {

	// Pull Overrides
	overrides := *this.SwitchOverrides

	// Position
	number := this.SelectedSwitch
	position := this.SwitchPosition[this.SelectedSwitch]

	// Color.
	color := overrides[number][position].Color
	maxNumberColors := overrides[number][position].MaxColors
	availableColors := overrides[number][position].AvailableColors
	//hasColorChannel := overrides[number][position].HasColorChannel
	//hasRGBChannels := overrides[number][position].HasRGBChannels
	colorName := "Unknown"
	if maxNumberColors > 0 && color <= maxNumberColors && color != -1 {
		colorName = availableColors[color]
	}
	return color, colorName
}

func getGobo(this *CurrentState) (int, string) {

	// Pull Overrides
	overrides := *this.SwitchOverrides

	// Position
	number := this.SelectedSwitch
	position := this.SwitchPosition[this.SelectedSwitch]

	// Gobo.
	gobo := overrides[number][position].Gobo
	maxNumberGobos := overrides[number][position].MaxGobos
	goboName := "Unknown"
	if maxNumberGobos > 0 && gobo < maxNumberGobos && gobo != -1 {
		if gobo != 0 {
			goboName = overrides[number][position].AvailableGobos[gobo-1]
		}
	}
	return gobo, goboName
}

func getProgram(this *CurrentState) (int, string) {

	// Pull Overrides
	overrides := *this.SwitchOverrides

	// Position
	number := this.SelectedSwitch
	position := this.SwitchPosition[this.SelectedSwitch]

	// Program.
	program := overrides[number][position].Program
	//isProgramOverrideAble := overrides[number][position].IsProgramOverrideAble
	maxPrograms := overrides[number][position].MaxPrograms
	programName := "Unknown"
	if maxPrograms > 0 && program <= maxPrograms && program != -1 {
		availablePrograms := overrides[number][position].AvailableProgramChannels
		if debug {
			fmt.Printf("AvailableProgramChannels %+v switchInfo.MaxPrograms %d switchInfo.Program %d\n", availablePrograms, maxPrograms, program)
		}
		if program > 0 {
			programName = availablePrograms[program-1]
		}
	}

	return program, programName
}

func getProgramSpeed(this *CurrentState) (int, string) {

	// Pull Overrides
	overrides := *this.SwitchOverrides

	// Position
	number := this.SelectedSwitch
	position := this.SwitchPosition[this.SelectedSwitch]

	// Program Speed.
	programSpeed := overrides[number][position].ProgramSpeed
	availableProgramSpeedChannels := len(overrides[number][position].AvailableProgramSpeedChannels)
	maxNumberProgramSpeeds := overrides[number][position].MaxProgramSpeeds
	//isProgramSpeedOverrideAble := overrides[number][position].IsProgramSpeedOverrideAble
	numberOfProgramSpeeds := availableProgramSpeedChannels
	programSpeedName := "Unknown"
	if numberOfProgramSpeeds > 0 && programSpeed <= maxNumberProgramSpeeds && programSpeed != -1 {
		availableProgramSpeeds := overrides[number][position].AvailableProgramSpeedChannels
		if programSpeed > 0 {
			programSpeedName = availableProgramSpeeds[programSpeed-1]
		}
	}

	return programSpeed, programSpeedName
}

func getSpeed(this *CurrentState) int {

	// Pull Overrides
	overrides := *this.SwitchOverrides

	return overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Speed

}

func labelButtons(b []bottonButton, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight) {

	//  The bottom row of the Novation Launchpad.
	bottomRow := 7

	// Loop through the available button names this sequence
	for index, button := range b {
		if debug {
			fmt.Printf("labelButtons %+v\n", button)
		}
		common.LightLamp(common.Button{X: index, Y: bottomRow}, button.Color, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)
		common.LabelButton(index, bottomRow, button.Label, guiButtons)
	}
}
