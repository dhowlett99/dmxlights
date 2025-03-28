package buttons

import (
	"fmt"

	"github.com/dhowlett99/dmxlights/pkg/common"
)

//	+-------------------+
//  |       NORMAL      |
//  +-------------------+
//		      |
//	          V
//	+-------------------+
//	|     FUNCTION      |
//	+-------------------+
//	    |            | If Scanner
//	    |            | or if the DisplayChaserShortCut is set.
//	    V            V
//	    |       +-------------------+
//	    |       |  CHASER DISPLAY   |
//	    |       +-------------------+
//	    V            |
//	    |            | If Scanner
//	    V            V
//	    |       +-------------------+
//	    |       |  CHASER FUNCTIONS |
//	    |       +-------------------+
//	    |		     |
//	    V			 V
//	 +-------------------+
//	 |  FIXTURE STATUS   |
//	 +-------------------+
//	          |
//	          V
//	 +-------------------+
//	 |       NORMAL      |
//	 +-------------------+
//

// HandleSelect - Runs when you press a select button to select a sequence.
func HandleSelect(sequences []*common.Sequence, this *CurrentState, eventsForLaunchpad chan common.ALight,
	commandChannels []chan common.Command, guiButtons chan common.ALight) {

	debug := false

	// Setup sequence numbers.
	if this.ScannerChaser[this.SelectedSequence] {
		this.TargetSequence = this.ChaserSequenceNumber
		this.DisplaySequence = this.SelectedSequence
	} else {
		this.TargetSequence = this.SelectedSequence
		this.DisplaySequence = this.SelectedSequence
	}

	if debug {
		printHandleDebug(this)
	}

	// Clear scannr color selection mode.
	if this.EditScannerColorsMode {

		if debug {
			fmt.Printf("%d: If we're in color selection mode. turn off color func key\n", this.ChaserSequenceNumber)
		}

		// Reset the gobo function key.
		this.Functions[this.TargetSequence][common.Function5_Color].State = false

		// Editing gobo is over for this sequence.
		this.EditScannerColorsMode = false
	}

	// Clear scanner gobo selection mode.
	if this.EditGoboSelectionMode {
		if debug {
			fmt.Printf("%d: If we're in gobo selection mode. turn off gobo func key\n", this.ChaserSequenceNumber)
		}

		// Reset the gobo function key.
		this.Functions[this.TargetSequence][common.Function6_Static_Gobo].State = false

		// Editing gobo is over for this sequence.
		this.EditGoboSelectionMode = false
	}

	// Clear pattern selection mode.
	if this.EditPatternMode {

		if debug {
			fmt.Printf("%d: If we're in pattern selection mode. turn off pattern func key\n", this.ChaserSequenceNumber)
		}

		// Reset the pattern function key.
		this.Functions[this.SelectedSequence][common.Function1_Pattern].State = false

		// Editing pattern is over for this sequence.
		this.EditPatternMode = false

		this.SelectedMode[this.SelectedSequence] = NORMAL
	}

	// Clear color picker.
	if this.ShowRGBColorPicker || this.ShowStaticColorPicker {
		if debug {
			fmt.Printf("Turn off the edit sequence colors button. \n")
		}
		this.ShowRGBColorPicker = false
		this.ShowStaticColorPicker = false

		this.Functions[this.EditWhichStaticSequence][common.Function5_Color].State = false
		removeColorPicker(this, sequences, eventsForLaunchpad, guiButtons, commandChannels)

		// If the Selected Color has come back as empty this means we didn't select any colors.
		// So restore the colors that were already there.
		if debug {
			fmt.Printf("sequences[%d].SequenceColors %+v\n", this.SelectedSequence, sequences[this.SelectedSequence].SequenceColors)
		}
		if len(sequences[this.SelectedSequence].SequenceColors) == 0 {
			if debug {
				fmt.Printf("Restore Sequence Colors\n")
			}
			sequences[this.SelectedSequence].SequenceColors = this.SavedSequenceColors[this.SelectedSequence]
			if debug {
				fmt.Printf("Now set to sequences[%d].SequenceColors %+v\n", this.SelectedSequence, sequences[this.SelectedSequence].SequenceColors)
			}
			// Tell the sequence that we have restored the colors.
			cmd := common.Command{
				Action: common.UpdateSequenceColors,
				Args: []common.Arg{
					{Name: "Colors", Value: sequences[this.SelectedSequence].SequenceColors},
				},
			}
			common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)
		}
	}

	// Decide if we're on the first press of the select button.
	if this.SelectButtonPressed[this.SelectedSequence] {
		// Calculate the next mode.
		this.SelectedMode[this.SelectedSequence] = getNextMenuItem(this.SelectedMode[this.SelectedSequence], this.ScannerChaser[this.SelectedSequence], getStatic(this))
	}
	if !this.SelectButtonPressed[this.SelectedSequence] {
		this.SelectButtonPressed[0] = false
		this.SelectButtonPressed[1] = false
		this.SelectButtonPressed[2] = false
		this.SelectButtonPressed[3] = false
		this.SelectButtonPressed[this.SelectedSequence] = true
	}

	// Jump straight to chaser display.
	if this.DisplayChaserShortCut {
		this.SelectedMode[this.SelectedSequence] = CHASER_DISPLAY
		this.DisplayChaserShortCut = false
	}

	// Clear the buttons.
	common.ClearSelectedRowOfButtons(this.SelectedSequence, eventsForLaunchpad, guiButtons)

	// If selected show the static sequence.
	if this.Static[this.TargetSequence] {
		if this.SelectedMode[this.DisplaySequence] == NORMAL_STATIC || this.SelectedMode[this.DisplaySequence] == CHASER_DISPLAY_STATIC {
			this.StaticFlashing[this.TargetSequence] = true
			this.SelectAllStaticFixtures = true
		} else {
			this.StaticFlashing[this.TargetSequence] = false
			this.SelectAllStaticFixtures = false
		}
		common.ShowStaticButtons(sequences[this.TargetSequence], this.StaticFlashing[this.TargetSequence], eventsForLaunchpad, guiButtons)
	}

	// Now display the selected mode.
	displayMode(this.SelectedSequence, this.SelectedMode[this.SelectedSequence], this, sequences, eventsForLaunchpad, guiButtons, commandChannels)

}

func getStatic(this *CurrentState) bool {

	// If we're a scanner static can be from either the scanner or shutter chaser static value.
	if this.SelectedSequence == this.ScannerSequenceNumber {
		return this.Static[this.SelectedSequence] || this.Static[this.ChaserSequenceNumber]
	}
	return this.Static[this.SelectedSequence]
}

func removeColorPicker(this *CurrentState, sequences []*common.Sequence, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command) {

	if debug {
		fmt.Printf("removeColorPicker Turn off the color picker\n")
	}

	this.Functions[this.EditWhichStaticSequence][common.Function5_Color].State = false

	// Clear the first three launchpad rows used by the color picker.
	for sequenceNumber, sequence := range sequences {
		common.ClearSelectedRowOfButtons(sequenceNumber, eventsForLaunchpad, guiButtons)

		if sequenceNumber != this.SwitchSequenceNumber {

			// Show the static and switch settings.
			cmd := common.Command{
				Action: common.Reveal,
			}
			common.SendCommandToSequence(sequenceNumber, cmd, commandChannels)

			this.SelectedMode[sequenceNumber] = NORMAL

			if this.Static[sequenceNumber] {
				common.ShowStaticButtons(sequence, false, eventsForLaunchpad, guiButtons)
			}
		}
	}
}

func showStatusBar(this *CurrentState, sequences []*common.Sequence, guiButtons chan common.ALight) {

	debug := false

	if debug {
		fmt.Printf("showStatusBar for sequence %d\n", this.SelectedSequence)
	}

	sensitivity := common.FindSensitivity(this.SoundGain)
	common.UpdateStatusBar(fmt.Sprintf("Sensitivity %02d", sensitivity), "sensitivity", false, guiButtons)
	common.UpdateStatusBar(fmt.Sprintf("Master %02d", this.MasterBrightness), "master", false, guiButtons)

	common.UpdateBottomButtons(this.SelectedType, guiButtons)

	// Make sure modes are setup.
	if sequences[this.SelectedSequence].Type == "scanner" && this.ScannerChaser[this.SelectedSequence] &&
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

	// Speed is common to all selectable sequences.
	UpdateSpeed(this, guiButtons)

	// RGB
	if sequences[this.DisplaySequence].Type == "rgb" &&
		(this.SelectedMode[this.DisplaySequence] == NORMAL || this.SelectedMode[this.DisplaySequence] == FUNCTION || this.SelectedMode[this.DisplaySequence] == STATUS) {

		if debug {
			fmt.Printf("showStatusBar show RGB labels\n")
		}

		common.UpdateStatusBar(fmt.Sprintf("Shift %02d", this.RGBShift[this.TargetSequence]), "shift", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Size %02d", this.RGBSize[this.TargetSequence]), "size", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Fade %02d", this.RGBFade[this.TargetSequence]), "fade", false, guiButtons)
		common.UpdateStatusBar("       ", "tilt", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Red %02d", this.StaticButtons[this.TargetSequence].Color.R), "red", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Green %02d", this.StaticButtons[this.TargetSequence].Color.G), "green", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Blue %02d", this.StaticButtons[this.TargetSequence].Color.B), "blue", false, guiButtons)

		common.LabelButton(0, 7, "Speed\nDown", guiButtons)
		common.LabelButton(1, 7, "Speed\nUp", guiButtons)
		common.LabelButton(2, 7, "Shift\nDown", guiButtons)
		common.LabelButton(3, 7, "Shift\nUp", guiButtons)
		common.LabelButton(4, 7, "Size\nDown", guiButtons)
		common.LabelButton(5, 7, "Size\nUp", guiButtons)
		common.LabelButton(6, 7, "Fade\nSoft", guiButtons)
		common.LabelButton(7, 7, "Fade\nSharp", guiButtons)
	}

	// SCANNER ROTATE
	if sequences[this.DisplaySequence].Type == "scanner" &&
		(this.SelectedMode[this.DisplaySequence] == NORMAL || this.SelectedMode[this.DisplaySequence] == FUNCTION || this.SelectedMode[this.DisplaySequence] == STATUS) {

		if debug {
			fmt.Printf("showStatusBar show Rotate labels\n")
		}

		if this.SelectedMode[this.TargetSequence] == NORMAL || this.SelectedMode[this.TargetSequence] == FUNCTION || this.SelectedMode[this.TargetSequence] == STATUS {
			label := getScannerShiftLabel(this.ScannerShift[this.TargetSequence])
			common.UpdateStatusBar(fmt.Sprintf("Rotate Shift %s", label), "shift", false, guiButtons)
			common.UpdateStatusBar(fmt.Sprintf("Rotate Size %02d", this.ScannerSize[this.TargetSequence]), "size", false, guiButtons)
			label = getScannerCoordinatesLabel(this.ScannerCoordinates[this.TargetSequence])
			common.UpdateStatusBar(fmt.Sprintf("Rotate Coord %s", label), "fade", false, guiButtons)

			common.LabelButton(0, 7, "Rotate\nSpeed\nDown", guiButtons)
			common.LabelButton(1, 7, "Rotate\nSpeed\nUp", guiButtons)
			common.LabelButton(2, 7, "Rotate\nShift\nDown", guiButtons)
			common.LabelButton(3, 7, "Rotate\nShift\nUp", guiButtons)
			common.LabelButton(4, 7, "Rotate\nSize\nDown", guiButtons)
			common.LabelButton(5, 7, "Rotate\nSize\nUp", guiButtons)
			common.LabelButton(6, 7, "Rotate\nCooord\nDown", guiButtons)
			common.LabelButton(7, 7, "Rotate\nCooord\nUp", guiButtons)

		}
	}
	// SHUTTER CHASER
	if sequences[this.DisplaySequence].Type == "scanner" &&
		(this.SelectedMode[this.DisplaySequence] == CHASER_DISPLAY || this.SelectedMode[this.DisplaySequence] == CHASER_FUNCTION) {

		if debug {
			fmt.Printf("showStatusBar show Chaser labels\n")
		}

		common.UpdateStatusBar(fmt.Sprintf("Chase Shift %02d", this.RGBShift[this.TargetSequence]), "shift", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Chase Size %02d", this.RGBSize[this.TargetSequence]), "size", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Chase Fade %02d", this.RGBFade[this.TargetSequence]), "fade", false, guiButtons)
		common.LabelButton(0, 7, "Chase\nSpeed\nDown", guiButtons)
		common.LabelButton(1, 7, "Chase\nSpeed\nUp", guiButtons)
		common.LabelButton(2, 7, "Chase\nShift\nDown", guiButtons)
		common.LabelButton(3, 7, "Chase\nShift\nUp", guiButtons)
		common.LabelButton(4, 7, "Chase\nSize\nDown", guiButtons)
		common.LabelButton(5, 7, "Chase\nSize\nUp", guiButtons)
		common.LabelButton(6, 7, "Chase\nFade\nSoft", guiButtons)
		common.LabelButton(7, 7, "Chase\nFade\nSharp", guiButtons)
	}

	// Hide the color editing buttons.
	common.UpdateStatusBar(fmt.Sprintf("Tilt %02d", this.OffsetTilt), "tilt", false, guiButtons)
	common.UpdateStatusBar("        ", "red", false, guiButtons)
	common.UpdateStatusBar("        ", "green", false, guiButtons)
	common.UpdateStatusBar(fmt.Sprintf("Pan %02d", this.OffsetPan), "pan", false, guiButtons)

}

func hideAllFunctionKeys(this *CurrentState, sequences []*common.Sequence, eventsForLaunchPad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command) {

	if debug {
		fmt.Printf("hideAllFunctionKeys\n")
	}

	for sequenceNumber := range sequences {
		if this.SelectedMode[sequenceNumber] == FUNCTION {
			common.ClearSelectedRowOfButtons(sequenceNumber, eventsForLaunchPad, guiButtons)
			// And reveal all the other sequence that isn't us.
			if sequenceNumber != this.SelectedSequence {
				// And turn off the function selected.
				displayMode(this.SelectedSequence, NORMAL, this, sequences, eventsForLaunchPad, guiButtons, commandChannels)
			}
		}

		if this.SelectedMode[sequenceNumber] == CHASER_FUNCTION {
			common.ClearSelectedRowOfButtons(sequenceNumber, eventsForLaunchPad, guiButtons)
			// And reveal all the other sequence that isn't us.
			if sequenceNumber != this.SelectedSequence {
				// And turn off the function selected.
				this.SelectedMode[sequenceNumber] = CHASER_DISPLAY
				displayMode(sequenceNumber, CHASER_DISPLAY, this, sequences, eventsForLaunchPad, guiButtons, commandChannels)
			}
		}
	}
}

// ALl states are based on the SelectedSequence. DisplaySequence.
func printHandleDebug(this *CurrentState) {
	fmt.Printf("HANDLE: this.Type = %s \n", this.SelectedType)
	for functionNumber := 0; functionNumber < 8; functionNumber++ {
		state := this.Functions[this.TargetSequence][functionNumber]
		fmt.Printf("%d function %d state %+v\n", this.SelectedSequence, functionNumber, state.State)
	}
	fmt.Printf("HANDLE: this.ScannerChaser[%d] running %t\n", this.DisplaySequence, this.ScannerChaser[this.DisplaySequence])
	fmt.Printf("HANDLE: ================== WHAT SELECT MODE =================\n")
	fmt.Printf("HANDLE: this.EditFixtureSelectionMode %t\n", this.EditFixtureSelectionMode)
	fmt.Printf("HANDLE: this.StaticFlashing %t\n", this.StaticFlashing[this.SelectedSequence])
	fmt.Printf("HANDLE: this.EditWhichStaticSequence = %d\n", this.EditWhichStaticSequence)
	fmt.Printf("HANDLE: this.SelectButtonPressed[%d] = %t\n", this.SelectedSequence, this.SelectButtonPressed[this.SelectedSequence])
	fmt.Printf("HANDLE: this.SelectedMode[%d] = %s\n", this.SelectedSequence, printMode(this.SelectedMode[this.SelectedSequence]))
	fmt.Printf("HANDLE: ================== WHAT EDIT MODES =================\n")
	fmt.Printf("HANDLE: this.ShowRGBColorPicker[%d] = %t\n", this.SelectedSequence, this.ShowRGBColorPicker)
	fmt.Printf("HANDLE: this.ShowStaticColorPicker[%d] = %t\n", this.SelectedSequence, this.ShowStaticColorPicker)
	fmt.Printf("HANDLE: this.Static[%d] = %t\n", this.SelectedSequence, this.Static[this.SelectedSequence])
	fmt.Printf("HANDLE: this.Static[%d] = %t\n", this.ChaserSequenceNumber, this.Static[this.ChaserSequenceNumber])
	fmt.Printf("HANDLE: this.EditGoboSelectionMode[%d] = %t\n", this.SelectedSequence, this.EditGoboSelectionMode)
	fmt.Printf("HANDLE: this.EditPatternMode[%d] = %t\n", this.SelectedSequence, this.EditPatternMode)
	fmt.Printf("HANDLE:===============================================\n")
}
