package buttons

import (
	"fmt"

	"github.com/dhowlett99/dmxlights/pkg/common"
)

func showStatusBars(this *CurrentState, sequences []*common.Sequence, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight) {

	debug := false

	if debug {
		fmt.Printf("showStatusBar for sequence %d\n", this.SelectedSequence)
	}

	sensitivity := common.FindSensitivity(this.SoundGain)
	common.UpdateStatusBar(fmt.Sprintf("Sensitivity %02d", sensitivity), "sensitivity", false, guiButtons)
	common.UpdateStatusBar(fmt.Sprintf("Master %02d", this.MasterBrightness), "master", false, guiButtons)

	// Make sure modes are setup.
	if this.SelectedType == "scanner" && this.ScannerChaser[this.SelectedSequence] &&
		(this.SelectedMode[this.SelectedSequence] == CHASER_FUNCTION || this.SelectedMode[this.SelectedSequence] == CHASER_DISPLAY) {
		this.TargetSequence = this.ChaserSequenceNumber
		this.DisplaySequence = this.SelectedSequence
	} else {
		this.TargetSequence = this.SelectedSequence
		this.DisplaySequence = this.SelectedSequence
	}

	if debug {
		fmt.Printf("Target Sequence %d Mode %s Type %s\n", this.TargetSequence, printMode(this.SelectedMode[this.TargetSequence]), sequences[this.TargetSequence].Type)
		fmt.Printf("Display Sequence %d Mode %s Type %s\n", this.DisplaySequence, printMode(this.SelectedMode[this.DisplaySequence]), sequences[this.DisplaySequence].Type)
	}

	// Update status bar.
	UpdateSpeed(this, guiButtons)
	UpdateShift(this, guiButtons)
	UpdateSize(this, guiButtons)
	UpdateFade(this, guiButtons)

	showTopLabels(this, eventsForLaunchpad, guiButtons)
	showBottomLabels(this, eventsForLaunchpad, guiButtons)

	// Hide the color editing buttons.
	common.UpdateStatusBar(fmt.Sprintf("Tilt %02d", this.OffsetTilt), "tilt", false, guiButtons)
	common.UpdateStatusBar("        ", "red", false, guiButtons)
	common.UpdateStatusBar("        ", "green", false, guiButtons)
	common.UpdateStatusBar(fmt.Sprintf("Pan %02d", this.OffsetPan), "pan", false, guiButtons)

}

func showTopLabels(this *CurrentState, eventsForLauchpad chan common.ALight, guiButtons chan common.ALight) {

	debug := false

	type topButton struct {
		Label string
		Color common.Color
	}
	// Storage for the rgb labels on the top row.
	var guiTopRGBButtons [8]topButton
	guiTopRGBButtons[0] = topButton{Label: "CLEAR", Color: common.Magenta}
	guiTopRGBButtons[1] = topButton{Label: "RED", Color: common.Red}
	guiTopRGBButtons[2] = topButton{Label: "GREEN", Color: common.Green}
	guiTopRGBButtons[3] = topButton{Label: "BLUE", Color: common.Blue}
	guiTopRGBButtons[4] = topButton{Label: "SENS -", Color: common.Cyan}
	guiTopRGBButtons[5] = topButton{Label: "SENS +", Color: common.Cyan}
	guiTopRGBButtons[6] = topButton{Label: "MAST -", Color: common.Cyan}
	guiTopRGBButtons[7] = topButton{Label: "MAST +", Color: common.Cyan}

	// Storage for the scanner labels on the Top row.
	var guiTopScannerButtons [8]topButton
	guiTopScannerButtons[0] = topButton{Label: "CLEAR.^", Color: common.White}
	guiTopScannerButtons[1] = topButton{Label: "V", Color: common.White}
	guiTopScannerButtons[2] = topButton{Label: "<", Color: common.White}
	guiTopScannerButtons[3] = topButton{Label: ">", Color: common.White}
	guiTopScannerButtons[4] = topButton{Label: "SENS -", Color: common.Cyan}
	guiTopScannerButtons[5] = topButton{Label: "SENS +", Color: common.Cyan}
	guiTopScannerButtons[6] = topButton{Label: "MAST -", Color: common.Cyan}
	guiTopScannerButtons[7] = topButton{Label: "MAST +", Color: common.Cyan}

	// Storage for the switch labels on the top row.
	var guiTopSwitchButtons [8]topButton
	guiTopSwitchButtons[0] = topButton{Label: "CLEAR", Color: common.Magenta}
	guiTopSwitchButtons[1] = topButton{Label: "RED", Color: common.Red}
	guiTopSwitchButtons[2] = topButton{Label: "GREEN", Color: common.Green}
	guiTopSwitchButtons[3] = topButton{Label: "BLUE", Color: common.Blue}
	guiTopSwitchButtons[4] = topButton{Label: "SENS -", Color: common.Cyan}
	guiTopSwitchButtons[5] = topButton{Label: "SENS +", Color: common.Cyan}
	guiTopSwitchButtons[6] = topButton{Label: "MAST -", Color: common.Cyan}
	guiTopSwitchButtons[7] = topButton{Label: "MAST +", Color: common.Cyan}

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

func showBottomLabels(this *CurrentState, eventsForLauchpad chan common.ALight, guiButtons chan common.ALight) {

	if debug {
		fmt.Printf("showBottomLabels type=%s fixture type=%s\n", this.SelectedType, this.SelectedFixtureType)
	}

	type bottonButton struct {
		Label string
		Color common.Color
	}

	// Storage for the rgb labels on the bottom row.
	var guiBottomRGBButtons [8]bottonButton
	guiBottomRGBButtons[0] = bottonButton{Label: "Speed\nDown", Color: common.Cyan}
	guiBottomRGBButtons[1] = bottonButton{Label: "Speed\nUp", Color: common.Cyan}
	guiBottomRGBButtons[2] = bottonButton{Label: "Shift\nDown", Color: common.Cyan}
	guiBottomRGBButtons[3] = bottonButton{Label: "Shift\nUp", Color: common.Cyan}
	guiBottomRGBButtons[4] = bottonButton{Label: "Size\nDown", Color: common.Cyan}
	guiBottomRGBButtons[5] = bottonButton{Label: "Size\nUp", Color: common.Cyan}
	guiBottomRGBButtons[6] = bottonButton{Label: "Fade\nSoft", Color: common.Cyan}
	guiBottomRGBButtons[7] = bottonButton{Label: "Fade\nSharp", Color: common.Cyan}

	// Storage for the scanner labels on the bottom row.
	var guiBottomScannerButtons [8]bottonButton
	guiBottomScannerButtons[0] = bottonButton{Label: "Speed\nDown", Color: common.Cyan}
	guiBottomScannerButtons[1] = bottonButton{Label: "Speed\nUp", Color: common.Cyan}
	guiBottomScannerButtons[2] = bottonButton{Label: "Shift\nDown", Color: common.Cyan}
	guiBottomScannerButtons[3] = bottonButton{Label: "Shift\nUp", Color: common.Cyan}
	guiBottomScannerButtons[4] = bottonButton{Label: "Size\nDown", Color: common.Cyan}
	guiBottomScannerButtons[5] = bottonButton{Label: "Size\nUp", Color: common.Cyan}
	guiBottomScannerButtons[6] = bottonButton{Label: "Coord\nDown", Color: common.Cyan}
	guiBottomScannerButtons[7] = bottonButton{Label: "Coord\nUp", Color: common.Cyan}

	// Storage for chaser labels on the bottom row.
	var guiBottomChaserButtons [8]bottonButton
	guiBottomChaserButtons[0] = bottonButton{Label: "Chase\nSpeed\nDown", Color: common.Cyan}
	guiBottomChaserButtons[1] = bottonButton{Label: "Chase\nSpeed\nUp", Color: common.Cyan}
	guiBottomChaserButtons[2] = bottonButton{Label: "Chase\nShift\nDown", Color: common.Cyan}
	guiBottomChaserButtons[3] = bottonButton{Label: "Chase\nShift\nUp", Color: common.Cyan}
	guiBottomChaserButtons[4] = bottonButton{Label: "Chase\nSize\nDown", Color: common.Cyan}
	guiBottomChaserButtons[5] = bottonButton{Label: "Chase\nSize\nUp", Color: common.Cyan}
	guiBottomChaserButtons[6] = bottonButton{Label: "Chase\nFade\nSoft", Color: common.Cyan}
	guiBottomChaserButtons[7] = bottonButton{Label: "Chase\nFade\nSharp", Color: common.Cyan}

	// Storage for the rgb labels on the bottom row.
	var guiBottomSwitchButtons [8]bottonButton
	guiBottomSwitchButtons[0] = bottonButton{Label: "Speed\nDown", Color: common.Cyan}
	guiBottomSwitchButtons[1] = bottonButton{Label: "Speed\nUp", Color: common.Cyan}
	guiBottomSwitchButtons[2] = bottonButton{Label: "Shift\nDown", Color: common.Cyan}
	guiBottomSwitchButtons[3] = bottonButton{Label: "Shift\nUp", Color: common.Cyan}
	guiBottomSwitchButtons[4] = bottonButton{Label: "Size\nDown", Color: common.Cyan}
	guiBottomSwitchButtons[5] = bottonButton{Label: "Size\nUp", Color: common.Cyan}
	guiBottomSwitchButtons[6] = bottonButton{Label: "Fade\nSoft", Color: common.Cyan}
	guiBottomSwitchButtons[7] = bottonButton{Label: "Fade\nSharp", Color: common.Cyan}

	// Storage for the rgb labels on the bottom row.
	var guiBottomProjectorButtons [8]bottonButton
	guiBottomProjectorButtons[0] = bottonButton{Label: "Shutter\nSpeed\nDown", Color: common.Cyan}
	guiBottomProjectorButtons[1] = bottonButton{Label: "Shutter\nSpeed\nUp", Color: common.Cyan}
	guiBottomProjectorButtons[2] = bottonButton{Label: "Rotate\nSpeed\nDown", Color: common.Cyan}
	guiBottomProjectorButtons[3] = bottonButton{Label: "Rotate\nSpeed\nUp", Color: common.Cyan}
	guiBottomProjectorButtons[4] = bottonButton{Label: "Color\nDown", Color: common.Cyan}
	guiBottomProjectorButtons[5] = bottonButton{Label: "Color\nUp", Color: common.Cyan}
	guiBottomProjectorButtons[6] = bottonButton{Label: "Gobo\nDown", Color: common.Cyan}
	guiBottomProjectorButtons[7] = bottonButton{Label: "Gobo\nUp", Color: common.Cyan}

	//  The bottom row of the Novation Launchpad.
	bottomRow := 7

	// RGB Front of house or uplighters.
	if this.SelectedType == "rgb" {

		common.UpdateStatusBar(fmt.Sprintf("Shift %02d", this.RGBShift[this.TargetSequence]), "shift", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Size %02d", this.RGBSize[this.TargetSequence]), "size", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Fade %02d", this.RGBFade[this.TargetSequence]), "fade", false, guiButtons)
		common.UpdateStatusBar("       ", "tilt", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Red %02d", this.StaticButtons[this.TargetSequence].Color.R), "red", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Green %02d", this.StaticButtons[this.TargetSequence].Color.G), "green", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Blue %02d", this.StaticButtons[this.TargetSequence].Color.B), "blue", false, guiButtons)

		// Loop through the available button names this sequence
		for index, button := range guiBottomRGBButtons {
			if debug {
				fmt.Printf("rgb button %+v\n", button)
			}
			common.LightLamp(common.Button{X: index, Y: bottomRow}, button.Color, common.MAX_DMX_BRIGHTNESS, eventsForLauchpad, guiButtons)
			common.LabelButton(index, bottomRow, button.Label, guiButtons)
		}
	}

	// Scanner showing rotate functions.
	if this.SelectedType == "scanner" &&
		(this.SelectedMode[this.DisplaySequence] == NORMAL || this.SelectedMode[this.DisplaySequence] == FUNCTION || this.SelectedMode[this.DisplaySequence] == STATUS) {

		common.UpdateStatusBar(fmt.Sprintf("Rotate Shift %s", getScannerShiftLabel(this.ScannerShift[this.TargetSequence])), "shift", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Rotate Size %02d", this.ScannerSize[this.TargetSequence]), "size", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Rotate Coord %s", getScannerCoordinatesLabel(this.ScannerCoordinates[this.TargetSequence])), "fade", false, guiButtons)

		// Loop through the available functions for this sequence
		for index, button := range guiBottomScannerButtons {
			if debug {
				fmt.Printf("scanner button %+v\n", button)
			}
			common.LightLamp(common.Button{X: index, Y: bottomRow}, button.Color, common.MAX_DMX_BRIGHTNESS, eventsForLauchpad, guiButtons)
			common.LabelButton(index, bottomRow, button.Label, guiButtons)
		}
	}

	// Shutter chaser showing RGB chase functions.
	if this.SelectedType == "scanner" &&
		(this.SelectedMode[this.DisplaySequence] == CHASER_DISPLAY || this.SelectedMode[this.DisplaySequence] == CHASER_FUNCTION) {

		common.UpdateStatusBar(fmt.Sprintf("Chase Shift %02d", this.RGBShift[this.TargetSequence]), "shift", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Chase Size %02d", this.RGBSize[this.TargetSequence]), "size", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Chase Fade %02d", this.RGBFade[this.TargetSequence]), "fade", false, guiButtons)

		// Loop through the available functions for this sequence
		for index, button := range guiBottomChaserButtons {
			if debug {
				fmt.Printf("chaser button %+v\n", button)
			}
			common.LightLamp(common.Button{X: index, Y: bottomRow}, button.Color, common.MAX_DMX_BRIGHTNESS, eventsForLauchpad, guiButtons)
			common.LabelButton(index, bottomRow, button.Label, guiButtons)
		}
	}

	// Switch functions.
	if this.SelectedType == "switch" && this.SelectedFixtureType != "projector" {
		// Loop through the available functions for this sequence
		for index, button := range guiBottomSwitchButtons {
			if debug {
				fmt.Printf("switch button %+v\n", button)
			}
			common.LightLamp(common.Button{X: index, Y: bottomRow}, button.Color, common.MAX_DMX_BRIGHTNESS, eventsForLauchpad, guiButtons)
			common.LabelButton(index, bottomRow, button.Label, guiButtons)
		}
		showSwitchColorDisplay(this, guiButtons)
	}

	// Projector functions.
	if this.SelectedType == "switch" && this.SelectedFixtureType == "projector" {

		common.UpdateStatusBar(fmt.Sprintf("Shutter Speed %02d", this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Speed), "speed", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Rotate Speed %02d", this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Shift), "shift", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Gobo %s", this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].GoboName), "fade", false, guiButtons)
		common.UpdateStatusBar("Colors", "size", false, guiButtons)
		showSwitchColorDisplay(this, guiButtons)

		// Loop through the available functions for this sequence
		for index, button := range guiBottomProjectorButtons {
			if debug {
				fmt.Printf("projector button %+v\n", button)
			}
			common.LightLamp(common.Button{X: index, Y: bottomRow}, button.Color, common.MAX_DMX_BRIGHTNESS, eventsForLauchpad, guiButtons)
			common.LabelButton(index, bottomRow, button.Label, guiButtons)
		}
	}
}

func showSwitchColorDisplay(this *CurrentState, guiButtons chan common.ALight) {
	if debug {
		fmt.Printf("Get color list for switch %d state %d\n", this.SelectedSwitch, this.SwitchPosition[this.SelectedSwitch])
	}
	control := common.GetColorList(this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Colors)
	if debug {
		fmt.Printf("Control %+v\n", control)
	}
	common.UpdateColorDisplay(control, guiButtons)
}
