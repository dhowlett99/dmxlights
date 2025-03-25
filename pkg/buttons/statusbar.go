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

type SwitchInfo struct {
	Number               int
	Mode                 int
	SelectedMode         int
	Type                 string
	FixtureType          string
	Speed                int
	RGBShift             int
	ScannerShift         string
	Size                 int
	RGBFade              int
	ScannerFade          int
	ScannerCoordinates   string
	Position             int
	IsRotateOverrideAble bool
	Rotate               int
	RotateName           string
	RotateSpeedName      string
	AvailableRotates     []string
	NumberOfRotates      int
	ColorIndex           int
	ColorName            string
	MaxNumberColors      int

	Gobo                   int
	OverrideAvailableGobos int
	GoboName               string
	NumberOfGobos          int
	MaxNumberGobos         int

	OverrideSpeed int
	OverrideSize  int
	OverrideFade  int
	OverrideGobo  int

	ProgramSpeedName              string
	ProgramSpeed                  int
	AvailableProgramSpeedChannels int
	MaxNumberProgramSpeeds        int
	NumberOfProgramSpeeds         int
	IsProgramSpeedOverrideAble    bool
	HasColorChannel               bool
	HasRGBChannels                bool
	ActionMode                    string
}

func getSwitchDetails(this *CurrentState) SwitchInfo {

	var switchInfo SwitchInfo

	// Position
	number := this.SelectedSwitch
	position := this.SwitchPosition[this.SelectedSwitch]

	// Pull overrides.
	overrides := *this.SwitchOverrides

	switchInfo.OverrideSpeed = overrides[number][position].Speed
	switchInfo.OverrideSize = overrides[number][position].Size
	switchInfo.OverrideGobo = overrides[number][position].Gobo
	switchInfo.AvailableRotates = overrides[number][position].RotateChannels
	switchInfo.Rotate = overrides[number][position].Rotate
	switchInfo.RotateSpeedName = overrides[number][position].RotateName
	switchInfo.IsRotateOverrideAble = overrides[number][position].IsRotateOverrideAble
	switchInfo.ColorIndex = overrides[number][position].Color
	switchInfo.MaxNumberColors = overrides[number][position].MaxColors
	switchInfo.HasColorChannel = overrides[number][position].HasColorChannel
	switchInfo.HasRGBChannels = overrides[number][position].HasRGBChannels
	switchInfo.ProgramSpeed = overrides[number][position].ProgramSpeed
	switchInfo.AvailableProgramSpeedChannels = len(overrides[number][position].AvailableProgramSpeedChannels)
	switchInfo.MaxNumberProgramSpeeds = overrides[number][position].MaxProgramSpeeds
	switchInfo.IsProgramSpeedOverrideAble = overrides[number][position].IsProgramSpeedOverrideAble
	switchInfo.ActionMode = overrides[number][position].Mode
	switchInfo.Gobo = overrides[number][position].Gobo
	switchInfo.OverrideAvailableGobos = len(overrides[number][position].AvailableGobos)
	switchInfo.MaxNumberGobos = overrides[number][position].MaxGobos
	switchInfo.OverrideFade = overrides[number][position].Fade

	switchInfo.SelectedMode = this.SelectedMode[this.DisplaySequence]
	switchInfo.Type = this.SelectedType
	switchInfo.FixtureType = this.SelectedFixtureType

	// Speed
	switchInfo.Speed = this.Speed[this.TargetSequence]

	// Size
	switchInfo.Size = this.RGBSize[this.TargetSequence]

	// Fade
	switchInfo.ScannerFade = this.ScannerSize[this.TargetSequence]
	switchInfo.RGBFade = this.RGBFade[this.TargetSequence]

	// Rotate
	switchInfo.NumberOfRotates = len(switchInfo.AvailableRotates)
	switchInfo.RotateName = "Unknown"
	if switchInfo.NumberOfRotates > 0 && switchInfo.Rotate <= switchInfo.NumberOfRotates && switchInfo.Rotate != -1 {
		if switchInfo.Rotate > 0 {
			switchInfo.RotateName = switchInfo.AvailableRotates[switchInfo.Rotate-1]
		}
	}

	// Shift
	switchInfo.RGBShift = this.RGBShift[this.TargetSequence]
	switchInfo.ScannerShift = getScannerShiftLabel(this.ScannerShift[this.TargetSequence])
	switchInfo.ProgramSpeedName = "Unknown"

	switchInfo.ColorName = "Unknown"

	// Colors
	if switchInfo.MaxNumberColors > 0 && switchInfo.ColorIndex <= switchInfo.MaxNumberColors && switchInfo.ColorIndex != -1 {
		if switchInfo.ColorIndex > 0 {
			switchInfo.ColorIndex--
		}
		switchInfo.ColorName = overrides[number][position].AvailableColors[switchInfo.ColorIndex]
	}

	// Program Speed
	switchInfo.NumberOfProgramSpeeds = switchInfo.AvailableProgramSpeedChannels
	if switchInfo.NumberOfProgramSpeeds > 0 && switchInfo.ProgramSpeed <= switchInfo.MaxNumberProgramSpeeds && switchInfo.ProgramSpeed != -1 {
		availableProgramSpeeds := overrides[number][position].AvailableProgramSpeedChannels
		if switchInfo.NumberOfProgramSpeeds > 0 {
			switchInfo.ProgramSpeedName = availableProgramSpeeds[switchInfo.ProgramSpeed-1]
		}
	}

	// Gobo
	switchInfo.GoboName = "Unknown"
	switchInfo.NumberOfGobos = switchInfo.OverrideAvailableGobos
	if switchInfo.NumberOfGobos > 0 && switchInfo.Gobo < switchInfo.MaxNumberGobos && switchInfo.Gobo != -1 {
		if switchInfo.Gobo != 0 {
			switchInfo.GoboName = overrides[number][position].AvailableGobos[switchInfo.Gobo-1]
		}
	}

	// Scanner
	switchInfo.ScannerCoordinates = getScannerCoordinatesLabel(this.ScannerCoordinates[this.TargetSequence])

	return switchInfo
}

func UpdateSpeed(this *CurrentState, guiButtons chan common.ALight) {

	switchInfo := getSwitchDetails(this)

	// Are we changing the strobe speed.
	if this.Strobe[this.SelectedSequence] {
		// Update the status bar
		common.UpdateStatusBar(fmt.Sprintf("Strobe %02d", this.StrobeSpeed[this.SelectedSequence]), "speed", false, guiButtons)
		return
	}

	// Are we in music trigger so no speed to change. Sequence or Switch.
	if this.Functions[this.TargetSequence][common.Function8_Music_Trigger].State || this.SwitchHasMusicTrigger {
		common.UpdateStatusBar("  MUSIC  ", "speed", false, guiButtons)
		return
	}

	// Chaser mode.
	if switchInfo.Mode == CHASER_DISPLAY || switchInfo.Mode == CHASER_FUNCTION {

		if !this.Strobe[this.TargetSequence] {
			common.UpdateStatusBar(fmt.Sprintf("Chase Speed %02d", switchInfo.Speed), "speed", false, guiButtons)
		} else {
			common.UpdateStatusBar(fmt.Sprintf("Strobe %02d", this.StrobeSpeed[this.TargetSequence]), "speed", false, guiButtons)
		}
		return
	}

	// Not a Chaser.
	if switchInfo.Mode == NORMAL || switchInfo.Mode == FUNCTION || switchInfo.Mode == STATUS {

		// Sequence is strobing this fixture.
		if this.Strobe[this.TargetSequence] {
			common.UpdateStatusBar(fmt.Sprintf("Strobe %02d", this.StrobeSpeed[this.TargetSequence]), "speed", false, guiButtons)
			return
		}

		// Sequence has a RGB fixture.
		if switchInfo.Type == "rgb" && !this.Strobe[this.TargetSequence] {
			common.UpdateStatusBar(fmt.Sprintf("Speed %02d", switchInfo.Speed), "speed", false, guiButtons)
			return
		}

		// Sequence has a Scanner fixture.
		if switchInfo.Type == "scanner" {
			common.UpdateStatusBar(fmt.Sprintf("Rotate Speed %02d", switchInfo.Speed), "speed", false, guiButtons)
			return
		}

		// Switch has a RGB fixture.
		if switchInfo.Type == "switch" && this.SelectedFixtureType == "rgb" {
			common.UpdateStatusBar(fmt.Sprintf("Speed %02d", switchInfo.Speed), "speed", false, guiButtons)
			return
		}

		// Switch has a scanner fixture.
		if switchInfo.Type == "switch" && this.SelectedFixtureType == "scanner" {
			common.UpdateStatusBar(fmt.Sprintf("Rotate Speed %02d", switchInfo.Speed), "speed", false, guiButtons)
			return
		}

		// Switch has a projector in control mode.
		if switchInfo.Type == "switch" &&
			this.SelectedFixtureType == "projector" &&
			switchInfo.IsProgramSpeedOverrideAble &&
			switchInfo.ActionMode == "Control" {

			common.UpdateStatusBar(fmt.Sprintf("Program Speed %02d:%s", switchInfo.ProgramSpeed, switchInfo.ProgramSpeedName), "speed", false, guiButtons)
			return
		}

		// // Switch has a projector that has a dedicated color wheel and associated channel.
		// if switchInfo.Type == "switch" && this.SelectedFixtureType == "projector" && switchInfo.HasColorChannel {
		// 	common.UpdateStatusBar(fmt.Sprintf("Speed %02d", switchInfo.Speed), "speed", false, guiButtons)
		// 	return
		// }

		// // Switch has a projector that has a RGB channels.
		// if switchInfo.Type == "switch" && this.SelectedFixtureType == "projector" && switchInfo.HasRGBChannels {
		// 	common.UpdateStatusBar(fmt.Sprintf("Speed %02d", switchInfo.Speed), "speed", false, guiButtons)
		// 	return
		// }

		// Assume nothing is selected, display a empty place holder.
		common.UpdateStatusBar(fmt.Sprintf("NOT AVAILABLE %02d", 0), "speed", false, guiButtons)
		return

	}
}

func UpdateSize(this *CurrentState, guiButtons chan common.ALight) {

	switchInfo := getSwitchDetails(this)

	if switchInfo.Mode == NORMAL || switchInfo.Mode == FUNCTION || switchInfo.Mode == STATUS {
		if switchInfo.Type == "rgb" || switchInfo.Type == "switch" {
			common.UpdateStatusBar(fmt.Sprintf("Size %02d", switchInfo.Size), "size", false, guiButtons)
		}
		if switchInfo.Type == "scanner" {
			common.UpdateStatusBar(fmt.Sprintf("Rotate Size %02d", switchInfo.ScannerFade), "size", false, guiButtons)
		}
		if switchInfo.Type == "switch" && this.SelectedFixtureType == "rgb" {
			common.UpdateStatusBar(fmt.Sprintf("Size %02d", switchInfo.OverrideSize), "size", false, guiButtons)
		}
		if switchInfo.Type == "switch" && this.SelectedFixtureType == "projector" {
			if this.SwitchStateName != "Off" {
				common.UpdateStatusBar(fmt.Sprintf("Color %02d:%s", switchInfo.ColorIndex, switchInfo.ColorName), "size", false, guiButtons)
			} else {
				common.ClearBottomStatusBar(guiButtons)
			}
		}
	}
	if switchInfo.Mode == CHASER_DISPLAY || switchInfo.Mode == CHASER_FUNCTION {
		common.UpdateStatusBar(fmt.Sprintf("Chase Size %02d", switchInfo.Size), "size", false, guiButtons)
	}
}

func UpdateShift(this *CurrentState, guiButtons chan common.ALight) {

	switchInfo := getSwitchDetails(this)

	if debug {
		fmt.Printf("UpdateShift RGBShift=%d scannerShift=%s switchShift=%d switchRotateSpeed %d switchRotateSpeedName=%s\n", switchInfo.RGBShift, switchInfo.ScannerShift, switchInfo.RGBShift, switchInfo.Rotate, switchInfo.RotateSpeedName)
	}

	// Chaser mode.
	if switchInfo.Mode == CHASER_DISPLAY || switchInfo.Mode == CHASER_FUNCTION {
		common.UpdateStatusBar(fmt.Sprintf("Chase Shift %02d", switchInfo.RGBShift), "shift", false, guiButtons)
	}

	// Not a Chaser.
	if switchInfo.Mode == NORMAL || switchInfo.Mode == FUNCTION || switchInfo.Mode == STATUS {

		// Sequence has a RGB fixture.
		if switchInfo.Type == "rgb" {
			common.UpdateStatusBar(fmt.Sprintf("Shift %02d", switchInfo.RGBShift), "shift", false, guiButtons)
			return
		}

		// Sequence has a Scanner fixture.
		if switchInfo.Type == "scanner" {
			common.UpdateStatusBar(fmt.Sprintf("Rotate Shift %s", switchInfo.ScannerShift), "shift", false, guiButtons)
			return
		}

		// Switch has a RGB fixture.
		if switchInfo.Type == "switch" && this.SelectedFixtureType == "rgb" {
			common.UpdateStatusBar(fmt.Sprintf("Shift %02d", switchInfo.RGBShift), "shift", false, guiButtons)
			return
		}

		// Switch has a scanner fixture.
		if switchInfo.Type == "switch" && switchInfo.FixtureType == "scanner" {
			common.UpdateStatusBar(fmt.Sprintf("Rotate Shift %s", switchInfo.ScannerShift), "shift", false, guiButtons)
			return
		}

		// Switch has a projector in control mode.
		if switchInfo.Type == "switch" && switchInfo.FixtureType == "projector" && switchInfo.IsRotateOverrideAble && switchInfo.ActionMode != "Control" {
			common.UpdateStatusBar(fmt.Sprintf("Rotate %02d:%s", switchInfo.Rotate, switchInfo.RotateName), "shift", false, guiButtons)
			return
		}

		// Assume nothing is selected, display a empty place holder.
		common.UpdateStatusBar(fmt.Sprintf("Shift N/A%02d", 0), "shift", false, guiButtons)
		return
	}

}

func UpdateFade(this *CurrentState, guiButtons chan common.ALight) {

	switchInfo := getSwitchDetails(this)

	if switchInfo.Mode == NORMAL || switchInfo.Mode == FUNCTION || switchInfo.Mode == STATUS {
		if switchInfo.Type == "rgb" {
			common.UpdateStatusBar(fmt.Sprintf("Fade %02d", switchInfo.RGBFade), "fade", false, guiButtons)
		}
		if switchInfo.Type == "scanner" {
			common.UpdateStatusBar(fmt.Sprintf("Rotate Coord %s", switchInfo.ScannerCoordinates), "fade", false, guiButtons)
		}
		if switchInfo.Type == "switch" && switchInfo.FixtureType == "rgb" {
			common.UpdateStatusBar(fmt.Sprintf("Fade %02d", switchInfo.RGBFade), "fade", false, guiButtons)
		}
		if switchInfo.Type == "switch" && switchInfo.FixtureType == "projector" {
			if this.SwitchStateName != "Off" {
				common.UpdateStatusBar(fmt.Sprintf("Gobo %02d:%s", switchInfo.Gobo, switchInfo.GoboName), "fade", false, guiButtons)
			} else {
				common.ClearBottomStatusBar(guiButtons)
			}
		}
	}
	if switchInfo.Mode == CHASER_DISPLAY || switchInfo.Mode == CHASER_FUNCTION {
		common.UpdateStatusBar(fmt.Sprintf("Chase Fade %02d", switchInfo.RGBFade), "fade", false, guiButtons)
	}
}
