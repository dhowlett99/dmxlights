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

	"github.com/dhowlett99/dmxlights/pkg/common"
)

func UpdateSpeed(this *CurrentState, guiButtons chan common.ALight) {

	selectedMode := this.SelectedMode[this.DisplaySequence]
	tYpe := this.SelectedType
	speed := this.Speed[this.TargetSequence]
	switchPosition := this.SwitchPosition[this.SelectedSwitch]
	overrides := *this.SwitchOverrides
	switchSpeed := overrides[this.SelectedSwitch][switchPosition].Speed
	switchNumber := this.SelectedSwitch
	switchProgramSpeedName := "Unknown"
	switchProgramSpeed := overrides[switchNumber][switchPosition].ProgramSpeed
	numberOfProgramSpeeds := len(overrides[switchNumber][switchPosition].AvailableProgramSpeedChannels)
	maxNumberProgramSpeeds := overrides[switchNumber][switchPosition].MaxProgramSpeeds
	isProgramSpeedOverrideAble := overrides[switchNumber][switchPosition].IsProgramSpeedOverrideAble
	actionMode := overrides[this.SelectedSwitch][switchPosition].Mode

	if this.Strobe[this.SelectedSequence] {
		// Update the status bar
		common.UpdateStatusBar(fmt.Sprintf("Strobe %02d", this.StrobeSpeed[this.SelectedSequence]), "speed", false, guiButtons)
		return
	}

	if numberOfProgramSpeeds > 0 && switchProgramSpeed <= maxNumberProgramSpeeds && switchProgramSpeed != -1 {
		availableProgramSpeeds := overrides[switchNumber][switchPosition].AvailableProgramSpeedChannels
		if switchProgramSpeed > 0 {
			switchProgramSpeedName = availableProgramSpeeds[switchProgramSpeed-1]
		}
	}

	if this.Functions[this.TargetSequence][common.Function8_Music_Trigger].State {
		common.UpdateStatusBar("  MUSIC  ", "speed", false, guiButtons)
	} else {

		if selectedMode == NORMAL || selectedMode == FUNCTION || selectedMode == STATUS {
			if tYpe == "rgb" {
				if !this.Strobe[this.TargetSequence] {
					common.UpdateStatusBar(fmt.Sprintf("Speed %02d", speed), "speed", false, guiButtons)
				} else {
					common.UpdateStatusBar(fmt.Sprintf("Strobe %02d", this.StrobeSpeed[this.TargetSequence]), "speed", false, guiButtons)
				}
				return
			}
			if tYpe == "scanner" {
				common.UpdateStatusBar(fmt.Sprintf("Rotate Speed %02d", speed), "speed", false, guiButtons)
			}
			if tYpe == "switch" {

				if this.SwitchStateName != "Off" {
					if isProgramSpeedOverrideAble && actionMode == "Control" {

						common.UpdateStatusBar(fmt.Sprintf("Program Speed %02d:%s", switchProgramSpeed, switchProgramSpeedName), "speed", false, guiButtons)
					} else {
						if this.MusicTrigger {
							common.UpdateStatusBar("MUSIC", "speed", false, guiButtons)
						} else {
							common.UpdateStatusBar(fmt.Sprintf("Speed %02d", switchSpeed), "speed", false, guiButtons)
						}
					}
				} else {
					common.ClearBottomStatusBar(guiButtons)
				}
				return
			}
		}
		if selectedMode == CHASER_DISPLAY || selectedMode == CHASER_FUNCTION {
			if !this.Strobe[this.TargetSequence] {
				common.UpdateStatusBar(fmt.Sprintf("Chase Speed %02d", speed), "speed", false, guiButtons)
			} else {
				common.UpdateStatusBar(fmt.Sprintf("Strobe %02d", this.StrobeSpeed[this.TargetSequence]), "speed", false, guiButtons)
			}
		}
	}
}

func UpdateSize(this *CurrentState, guiButtons chan common.ALight) {

	mode := this.SelectedMode[this.DisplaySequence]
	tYpe := this.SelectedType
	size := this.RGBSize[this.TargetSequence]
	scannerFade := this.ScannerSize[this.TargetSequence]
	overrides := *this.SwitchOverrides

	if mode == NORMAL || mode == FUNCTION || mode == STATUS {
		if tYpe == "rgb" || tYpe == "switch" {
			common.UpdateStatusBar(fmt.Sprintf("Size %02d", size), "size", false, guiButtons)
		}
		if tYpe == "scanner" {
			common.UpdateStatusBar(fmt.Sprintf("Rotate Size %02d", scannerFade), "size", false, guiButtons)
		}
		if tYpe == "switch" && this.SelectedFixtureType == "rgb" {
			switchSize := overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Size
			common.UpdateStatusBar(fmt.Sprintf("Size %02d", switchSize), "size", false, guiButtons)
		}
		if tYpe == "switch" && this.SelectedFixtureType == "projector" {
			switchColorIndex := overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Color
			switchColorName := "Unknown"
			switchMaxNumberColors := overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].MaxColors
			if switchMaxNumberColors > 0 && switchColorIndex <= switchMaxNumberColors && switchColorIndex != -1 {
				if switchColorIndex > 0 {
					switchColorIndex--
				}
				switchColorName = overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].AvailableColors[switchColorIndex]
			}
			if this.SwitchStateName != "Off" {
				common.UpdateStatusBar(fmt.Sprintf("Color %02d:%s", switchColorIndex, switchColorName), "size", false, guiButtons)
			} else {
				common.ClearBottomStatusBar(guiButtons)
			}
		}
	}
	if mode == CHASER_DISPLAY || mode == CHASER_FUNCTION {
		common.UpdateStatusBar(fmt.Sprintf("Chase Size %02d", size), "size", false, guiButtons)
	}
}

func UpdateShift(this *CurrentState, guiButtons chan common.ALight) {

	mode := this.SelectedMode[this.DisplaySequence]
	tYpe := this.SelectedType
	shift := this.RGBShift[this.TargetSequence]
	scannerShift := getScannerShiftLabel(this.ScannerShift[this.TargetSequence])
	overrides := *this.SwitchOverrides
	switchNumber := this.SelectedSwitch
	switchPosition := this.SwitchPosition[this.SelectedSwitch]
	switchRGBShift := overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Shift
	switchRotate := overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Rotate
	switchRotateSpeedName := overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].RotateName

	availableRotates := overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].RotateChannels
	numberOfRotates := len(availableRotates)
	maxNumberRotates := overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].MaxRotateSpeed
	rotateName := "Unknown"
	if numberOfRotates > 0 && switchRotate <= maxNumberRotates && switchRotate != -1 {
		if switchRotate > 0 {
			rotateName = availableRotates[switchRotate-1]
		}
	}
	isRotateOverrideAble := overrides[switchNumber][switchPosition].IsRotateOverrideAble

	if debug {
		fmt.Printf("UpdateShift RGBShift=%d scannerShift=%s switchShift=%d switchRotateSpeed %d switchRotateSpeedName=%s\n", shift, scannerShift, switchRGBShift, switchRotate, switchRotateSpeedName)
	}

	if mode == NORMAL || mode == FUNCTION || mode == STATUS {
		if tYpe == "rgb" {
			common.UpdateStatusBar(fmt.Sprintf("Shift %02d", shift), "shift", false, guiButtons)
		}
		if tYpe == "scanner" {
			common.UpdateStatusBar(fmt.Sprintf("Rotate Shift %s", scannerShift), "shift", false, guiButtons)
		}
		if tYpe == "switch" {

			if this.SwitchStateName != "Off" {

				if this.SelectedFixtureType == "rgb" {
					common.UpdateStatusBar(fmt.Sprintf("Shift %02d", switchRGBShift), "shift", false, guiButtons)
				}

				if isRotateOverrideAble {
					if this.SelectedFixtureType == "projector" {
						common.UpdateStatusBar(fmt.Sprintf("Rotate %02d:%s", switchRotate, rotateName), "shift", false, guiButtons)
					}
				}
			} else {
				// Display a empty place holder.
				common.UpdateStatusBar(fmt.Sprintf("Shift %02d", 0), "shift", false, guiButtons)
			}
		}
	}
	if mode == CHASER_DISPLAY || mode == CHASER_FUNCTION {
		common.UpdateStatusBar(fmt.Sprintf("Chase Shift %02d", shift), "shift", false, guiButtons)
	}
}

func UpdateFade(this *CurrentState, guiButtons chan common.ALight) {

	mode := this.SelectedMode[this.DisplaySequence]
	tYpe := this.SelectedType
	fixtureType := this.SelectedFixtureType
	fade := this.RGBFade[this.TargetSequence]
	scannerCoordinates := getScannerCoordinatesLabel(this.ScannerCoordinates[this.TargetSequence])
	overrides := *this.SwitchOverrides
	switchFade := overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Fade
	switchGobo := overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Gobo

	if mode == NORMAL || mode == FUNCTION || mode == STATUS {
		if tYpe == "rgb" {
			common.UpdateStatusBar(fmt.Sprintf("Fade %02d", fade), "fade", false, guiButtons)
		}
		if tYpe == "scanner" {
			common.UpdateStatusBar(fmt.Sprintf("Rotate Coord %s", scannerCoordinates), "fade", false, guiButtons)
		}
		if tYpe == "switch" && fixtureType == "rgb" {
			common.UpdateStatusBar(fmt.Sprintf("Fade %02d", switchFade), "fade", false, guiButtons)
		}
		if tYpe == "switch" && fixtureType == "projector" {
			switchGoboName := "Unknown"
			numberOfGobos := len(overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].AvailableGobos)
			maxNumberGobos := overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].MaxGobos
			if numberOfGobos > 0 && switchGobo < maxNumberGobos && switchGobo != -1 {
				if switchGobo != 0 {
					switchGoboName = overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].AvailableGobos[switchGobo-1]
				}
			}

			if this.SwitchStateName != "Off" {
				common.UpdateStatusBar(fmt.Sprintf("Gobo %02d:%s", switchGobo, switchGoboName), "fade", false, guiButtons)
			} else {
				common.ClearBottomStatusBar(guiButtons)
			}
		}
	}
	if mode == CHASER_DISPLAY || mode == CHASER_FUNCTION {
		common.UpdateStatusBar(fmt.Sprintf("Chase Fade %02d", fade), "fade", false, guiButtons)
	}
}
