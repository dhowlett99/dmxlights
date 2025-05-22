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
	AvailableRotates     []string
	NumberOfRotates      int
	Color                int
	ColorName            string
	MaxNumberColors      int
	AvailableColors      []string

	Gobo           int
	GoboName       string
	MaxNumberGobos int

	OverrideSpeed int
	OverrideSize  int
	OverrideFade  int
	OverrideGobo  int

	ProgramSpeedName              string
	ProgramSpeed                  int
	MaxPrograms                   int
	Program                       int
	IsProgramOverrideAble         bool
	ProgramName                   string
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

	// Pull overrides.
	overrides := *this.SwitchOverrides

	// Position
	number := this.SelectedSwitch
	position := this.SwitchPosition[this.SelectedSwitch]

	// Type.
	switchInfo.Type = this.SelectedType

	// Fixture Type.
	switchInfo.FixtureType = this.SelectedFixtureType

	// Mode.
	switchInfo.SelectedMode = this.SelectedMode[this.DisplaySequence]
	switchInfo.ActionMode = overrides[number][position].Mode

	// Speed.
	switchInfo.Speed = this.Speed[this.TargetSequence]
	switchInfo.OverrideSpeed = overrides[number][position].Speed

	// Shift.
	switchInfo.RGBShift = this.RGBShift[this.TargetSequence]
	switchInfo.ScannerShift = getScannerShiftLabel(this.ScannerShift[this.TargetSequence])

	// Size.
	switchInfo.Size = this.RGBSize[this.TargetSequence]
	switchInfo.OverrideSize = overrides[number][position].Size

	// Fade.
	switchInfo.ScannerFade = this.ScannerSize[this.TargetSequence]
	switchInfo.RGBFade = this.RGBFade[this.TargetSequence]
	switchInfo.OverrideFade = overrides[number][position].Fade

	// Scanner Coordinates.
	switchInfo.ScannerCoordinates = getScannerCoordinatesLabel(this.ScannerCoordinates[this.TargetSequence])

	// Gobo.
	switchInfo.OverrideGobo = overrides[number][position].Gobo
	switchInfo.Gobo = overrides[number][position].Gobo
	switchInfo.MaxNumberGobos = overrides[number][position].MaxGobos
	switchInfo.GoboName = "Unknown"
	if switchInfo.MaxNumberGobos > 0 && switchInfo.Gobo < switchInfo.MaxNumberGobos && switchInfo.Gobo != -1 {
		if switchInfo.Gobo != 0 {
			switchInfo.GoboName = overrides[number][position].AvailableGobos[switchInfo.Gobo-1]
		}
	}

	// Rotate.
	switchInfo.Rotate = overrides[number][position].Rotate
	switchInfo.AvailableRotates = overrides[number][position].RotateChannels
	switchInfo.RotateName = overrides[number][position].RotateName
	switchInfo.IsRotateOverrideAble = overrides[number][position].IsRotateOverrideAble
	switchInfo.NumberOfRotates = len(switchInfo.AvailableRotates)
	switchInfo.RotateName = "Unknown"
	if switchInfo.NumberOfRotates > 0 && switchInfo.Rotate <= switchInfo.NumberOfRotates && switchInfo.Rotate != -1 {
		if switchInfo.Rotate > 0 {
			switchInfo.RotateName = switchInfo.AvailableRotates[switchInfo.Rotate-1]
		}
	}

	// Color.
	switchInfo.Color = overrides[number][position].Color
	switchInfo.MaxNumberColors = overrides[number][position].MaxColors
	switchInfo.AvailableColors = overrides[number][position].AvailableColors
	switchInfo.HasColorChannel = overrides[number][position].HasColorChannel
	switchInfo.HasRGBChannels = overrides[number][position].HasRGBChannels
	switchInfo.ColorName = "Unknown"
	if switchInfo.MaxNumberColors > 0 && switchInfo.Color <= switchInfo.MaxNumberColors && switchInfo.Color != -1 {
		switchInfo.ColorName = overrides[number][position].AvailableColors[switchInfo.Color]
	}

	// Program.
	switchInfo.Program = overrides[number][position].Program
	switchInfo.IsProgramOverrideAble = overrides[number][position].IsProgramOverrideAble
	switchInfo.MaxPrograms = overrides[number][position].MaxPrograms
	switchInfo.ProgramName = "Unknown"
	if switchInfo.MaxPrograms > 0 && switchInfo.Program <= switchInfo.MaxPrograms && switchInfo.Program != -1 {
		availablePrograms := overrides[number][position].AvailableProgramChannels
		if debug {
			fmt.Printf("AvailableProgramChannels %+v switchInfo.MaxPrograms %d switchInfo.Program %d\n", overrides[number][position].AvailableProgramChannels, switchInfo.MaxPrograms, switchInfo.Program)
		}
		if switchInfo.Program > 0 {
			switchInfo.ProgramName = availablePrograms[switchInfo.Program-1]
		}
	}

	// Program Speed.
	switchInfo.ProgramSpeed = overrides[number][position].ProgramSpeed
	switchInfo.AvailableProgramSpeedChannels = len(overrides[number][position].AvailableProgramSpeedChannels)
	switchInfo.MaxNumberProgramSpeeds = overrides[number][position].MaxProgramSpeeds
	switchInfo.IsProgramSpeedOverrideAble = overrides[number][position].IsProgramSpeedOverrideAble
	switchInfo.NumberOfProgramSpeeds = switchInfo.AvailableProgramSpeedChannels
	switchInfo.ProgramSpeedName = "Unknown"
	if switchInfo.NumberOfProgramSpeeds > 0 && switchInfo.ProgramSpeed <= switchInfo.MaxNumberProgramSpeeds && switchInfo.ProgramSpeed != -1 {
		availableProgramSpeeds := overrides[number][position].AvailableProgramSpeedChannels
		if switchInfo.ProgramSpeed > 0 {
			switchInfo.ProgramSpeedName = availableProgramSpeeds[switchInfo.ProgramSpeed-1]
		}
	}

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

		if switchInfo.Type == "switch" &&
			this.SelectedFixtureType == "projector" {
			common.UpdateStatusBar(fmt.Sprintf("Speed %02d", switchInfo.OverrideSpeed), "speed", false, guiButtons)
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
		common.UpdateStatusBar(fmt.Sprintf("Speed %02d", 0), "speed", false, guiButtons)
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
			common.UpdateStatusBar(fmt.Sprintf("Color %02d:%s", switchInfo.Color, switchInfo.ColorName), "size", false, guiButtons)
		}
	}
	if switchInfo.Mode == CHASER_DISPLAY || switchInfo.Mode == CHASER_FUNCTION {
		common.UpdateStatusBar(fmt.Sprintf("Chase Size %02d", switchInfo.Size), "size", false, guiButtons)
	}
}

func UpdateShift(this *CurrentState, guiButtons chan common.ALight) {

	switchInfo := getSwitchDetails(this)

	if debug {
		fmt.Printf("UpdateShift RGBShift=%d scannerShift=%s switchShift=%d switchRotateSpeed %d switchRotateSpeedName=%s\n", switchInfo.RGBShift, switchInfo.ScannerShift, switchInfo.RGBShift, switchInfo.Rotate, switchInfo.RotateName)
		fmt.Printf("UpdateShift switchInfo.Type %s  switchInfo.FixtureType %s IsRotateOverrideAble %t ActionMode %s\n", switchInfo.Type, switchInfo.FixtureType, switchInfo.IsRotateOverrideAble, switchInfo.ActionMode)
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

		// Switch has a projector but not in control mode.
		if switchInfo.Type == "switch" && switchInfo.FixtureType == "projector" && switchInfo.IsRotateOverrideAble && switchInfo.ActionMode != "Control" {
			common.UpdateStatusBar(fmt.Sprintf("Rotate %02d:%s", switchInfo.Rotate, switchInfo.RotateName), "shift", false, guiButtons)
			return
		}

		// Switch has a projector in control mode. So shift becomes select Program or Show.
		if switchInfo.Type == "switch" && switchInfo.FixtureType == "projector" && switchInfo.IsProgramOverrideAble && switchInfo.ActionMode == "Control" {
			common.UpdateStatusBar(fmt.Sprintf("Program %02d:%s", switchInfo.Program, switchInfo.ProgramName), "shift", false, guiButtons)
			return
		}

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
			common.UpdateStatusBar(fmt.Sprintf("Gobo %02d:%s", switchInfo.Gobo, switchInfo.GoboName), "fade", false, guiButtons)
		}
	}
	if switchInfo.Mode == CHASER_DISPLAY || switchInfo.Mode == CHASER_FUNCTION {
		common.UpdateStatusBar(fmt.Sprintf("Chase Fade %02d", switchInfo.RGBFade), "fade", false, guiButtons)
	}
}
