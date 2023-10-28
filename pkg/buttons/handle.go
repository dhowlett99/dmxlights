package buttons

import (
	"fmt"

	"github.com/dhowlett99/dmxlights/pkg/common"
)

//		+-------------------+
//		|       NORMAL      |
//		+-------------------+
//		    |            |

//	+-------------------+
//	|     FUNCTION      |
//	+-------------------+
//	    |            | If Scanner
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

	if this.SelectMode[this.SelectedSequence] == CHASER_FUNCTION {
		this.TargetSequence = this.ChaserSequenceNumber
		this.DisplaySequence = this.SelectedSequence
	} else {
		this.TargetSequence = this.SelectedSequence
		this.DisplaySequence = this.SelectedSequence
	}

	if debug {
		printHandleDebug(this)
	}

	// Select Chase Pattern.
	if this.EditPatternMode {

		if debug {
			fmt.Printf("%d: If we're in pattern selection mode. turn off pattern func key\n", this.ChaserSequenceNumber)
		}

		// Reset the pattern function key.
		this.Functions[this.SelectedSequence][common.Function1_Pattern].State = false

		// Editing pattern is over for this sequence.
		this.EditPatternMode = false

		this.SelectMode[this.SelectedSequence] = NORMAL
	}

	// Select Chase Sequence Colors. Turn off the edit sequence colors button.
	if this.ShowRGBColorPicker {
		if debug {
			fmt.Printf("Turn off the edit sequence colors button. \n")
		}
		this.ShowRGBColorPicker = false
		this.Functions[this.EditWhichStaticSequence][common.Function5_Color].State = false
		removeColorPicker(this, eventsForLaunchpad, guiButtons, commandChannels)
		this.SelectMode[this.SelectedSequence] = NORMAL

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
				fmt.Printf("Now set to ----> sequences[%d].SequenceColors %+v\n", this.SelectedSequence, sequences[this.SelectedSequence].SequenceColors)
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

	// 1st Press Select Sequence - This the first time we have pressed the select button.
	// Simply select the selected sequence.
	// But remember we have pressed this select button once.
	if this.SelectMode[this.DisplaySequence] == NORMAL && !this.SelectButtonPressed[this.DisplaySequence] {

		if debug {
			fmt.Printf("%d: Show Sequence - Handle Step 1\n", this.SelectedSequence)
		}

		// Assume everything else is off.
		this.SelectButtonPressed[0] = false
		this.SelectButtonPressed[1] = false
		this.SelectButtonPressed[2] = false
		this.SelectButtonPressed[3] = false

		// Remember which select button has been pressed.
		this.SelectButtonPressed[this.SelectedSequence] = true

		// Now display the selected mode.
		displayMode(this.SelectMode[this.SelectedSequence], this, sequences, eventsForLaunchpad, guiButtons, commandChannels)

		return
	}

	// Fisrt option of the 2nd Press we are in rgb mode same select button go into Function Mode for this sequence.
	// We're in Normal mode, We've Pressed twice. we're a rgb sequence.
	// Or
	// We're in Normal mode, We've Pressed twice. but the scanner chaser is off
	// Or
	// We're in Chaser Display mode. and scanner sequence.
	if (this.SelectMode[this.SelectedSequence] == NORMAL && this.SelectedType == "rgb") ||
		(this.SelectMode[this.SelectedSequence] == NORMAL && !this.ScannerChaser[this.SelectedSequence]) ||
		(this.SelectMode[this.SelectedSequence] == CHASER_DISPLAY && this.SelectedType == "scanner") {

		if debug {
			fmt.Printf("%d: 2nd Press Function Bar Mode - Handle Step 2\n", this.SelectedSequence)
		}

		// Calculate the next mode.
		this.SelectMode[this.SelectedSequence] = getNextMenuItem(this.SelectMode[this.SelectedSequence], this.ScannerChaser[this.SelectedSequence])

		// Now forget we pressed twice and start again.
		this.SelectButtonPressed[this.SelectedSequence] = false

		// Now display the selected mode.
		displayMode(this.SelectMode[this.SelectedSequence], this, sequences, eventsForLaunchpad, guiButtons, commandChannels)

		return
	}

	// Second option of the 2nd Press in scanner mode go into Shutter Chaser Mode for this sequence.
	// We're in Normal mode, We've Pressed twice and the scanner chaser is on and we're a scanner sequence.
	if this.SelectMode[this.SelectedSequence] == NORMAL &&
		this.ScannerChaser[this.SelectedSequence] &&
		this.SelectedType == "scanner" &&
		this.SelectButtonPressed[this.SelectedSequence] {

		if debug {
			fmt.Printf("%d: We are a SCANNER, Shutter Chaser On- 2nd Press Shutter Chase Mode - Handle Step 2\n", this.SelectedSequence)
		}

		// Calculate the next mode.
		this.SelectMode[this.SelectedSequence] = getNextMenuItem(this.SelectMode[this.SelectedSequence], this.ScannerChaser[this.SelectedSequence])

		// Now display the selected mode.
		displayMode(this.SelectMode[this.SelectedSequence], this, sequences, eventsForLaunchpad, guiButtons, commandChannels)

		return
	}

	// 3rd Press Status Mode and not a scanner - we display the fixture status enable/invert/disable buttons.
	// We're in Function mode, pressed twice. and we're in edit sequence color mode. and we're NOT a scanner.
	if this.SelectMode[this.SelectedSequence] == FUNCTION &&
		!this.SelectButtonPressed[this.SelectedSequence] &&
		!this.ShowRGBColorPicker &&
		this.SelectedType != "scanner" {

		if debug {
			fmt.Printf("%d: Handle 3 - Status Buttons on\n", this.SelectedSequence)
		}

		// Calculate the next mode.
		this.SelectMode[this.SelectedSequence] = getNextMenuItem(this.SelectMode[this.SelectedSequence], this.ScannerChaser[this.SelectedSequence])

		// Now display the selected mode.
		displayMode(this.SelectMode[this.SelectedSequence], this, sequences, eventsForLaunchpad, guiButtons, commandChannels)

		return
	}

	// 3rd Press Status Mode and we are scanner and the shutter chaser is running - we display the shutter chaser function buttons.
	if this.SelectMode[this.SelectedSequence] == FUNCTION &&
		this.ScannerChaser[this.SelectedSequence] &&
		!this.ShowRGBColorPicker &&
		this.SelectedType == "scanner" {

		if debug {
			fmt.Printf("%d: Handle Step 4 Shutter Chase Function buttons on\n", this.SelectedSequence)
		}

		// Calculate the next mode.
		this.SelectMode[this.SelectedSequence] = getNextMenuItem(this.SelectMode[this.SelectedSequence], this.ScannerChaser[this.SelectedSequence])

		// Now display the selected mode.
		displayMode(this.SelectMode[this.SelectedSequence], this, sequences, eventsForLaunchpad, guiButtons, commandChannels)

		return
	}

	// 4th Press Normal Mode - we head back to normal mode.
	if this.SelectMode[this.SelectedSequence] == STATUS &&
		!this.SelectButtonPressed[this.SelectedSequence] &&
		!this.ShowRGBColorPicker {

		if debug {
			fmt.Printf("%d: Handle Step 5 - Normal Mode From Non Scanner, Function Bar off\n", this.SelectedSequence)
		}

		// Remember that we've preseed twice.
		this.SelectButtonPressed[this.SelectedSequence] = true

		// Calculate the next mode.
		this.SelectMode[this.SelectedSequence] = getNextMenuItem(this.SelectMode[this.SelectedSequence], this.ScannerChaser[this.SelectedSequence])

		// Now display the selected mode.
		displayMode(this.SelectMode[this.SelectedSequence], this, sequences, eventsForLaunchpad, guiButtons, commandChannels)

		return
	}

	// 4th Press Normal Mode and we are a scanner- we head fixture status mode.
	if (this.SelectMode[this.SelectedSequence] == CHASER_FUNCTION || !this.ScannerChaser[this.SelectedSequence]) &&
		!this.SelectButtonPressed[this.SelectedSequence] &&
		!this.ShowRGBColorPicker &&
		this.SelectedType == "scanner" {

		if debug {
			fmt.Printf("%d: Handle Step 6 Normal Mode, From  Scanner, Function Bar off, status buttons on\n", this.SelectedSequence)
		}

		// Calculate the next mode.
		this.SelectMode[this.SelectedSequence] = getNextMenuItem(this.SelectMode[this.SelectedSequence], this.ScannerChaser[this.SelectedSequence])

		// Now display the selected mode.
		displayMode(this.SelectMode[this.SelectedSequence], this, sequences, eventsForLaunchpad, guiButtons, commandChannels)

		return
	}

	if debug {
		fmt.Printf("HANDLE: Sequence %d Nothing Handled  \n", this.TargetSequence)
	}

}

func removeColorPicker(this *CurrentState, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command) {

	this.SelectButtonPressed[this.SelectedSequence] = false
	this.SelectMode[this.SelectedSequence] = NORMAL
	this.Functions[this.EditWhichStaticSequence][common.Function5_Color].State = false

	// Clear the first three launchpad rows used by the color picker.
	for y := 0; y < 3; y++ {
		common.ClearSelectedRowOfButtons(y, eventsForLaunchpad, guiButtons)
	}

	// Show the static and switch settings.
	cmd := common.Command{
		Action: common.UnHide,
	}

	// Take account of the shutter chaser which should be shown if the chaser is running.
	common.SendCommandToSequence(0, cmd, commandChannels)
	common.SendCommandToSequence(1, cmd, commandChannels)
	if !this.ScannerChaser[this.SelectedSequence] {
		common.SendCommandToSequence(2, cmd, commandChannels)
	} else {
		common.SendCommandToSequence(4, cmd, commandChannels)
	}

}

func displayMode(mode int, this *CurrentState, sequences []*common.Sequence, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command) {

	// Clear the buttons.
	common.ClearSelectedRowOfButtons(this.SelectedSequence, eventsForLaunchpad, guiButtons)

	if debug {
		printMode(this)
	}

	// Hide any function keys.
	hideAllFunctionKeys(this, sequences, eventsForLaunchpad, guiButtons, commandChannels)

	// Tailor the top buttons to the sequence type.
	common.ShowTopButtons(sequences[this.SelectedSequence].Type, eventsForLaunchpad, guiButtons)

	// Tailor the bottom buttons to the sequence type.
	common.ShowBottomButtons(sequences[this.SelectedSequence].Type, eventsForLaunchpad, guiButtons)

	// Show this sequence running status in the start/stop button.
	common.ShowRunningStatus(this.SelectedSequence, this.Running, eventsForLaunchpad, guiButtons)

	// Update the status bar.
	showStatusBar(this, sequences, guiButtons)

	// Light the sequence selector button.
	SequenceSelect(eventsForLaunchpad, guiButtons, this)

	switch {

	case mode == NORMAL:

		if debug {
			fmt.Printf("displayMode: NORMAL\n")
		}

		// If we have a shutter chaser running hide it.
		if this.ScannerChaser[this.SelectedSequence] && this.SelectedType == "scanner" {
			common.HideSequence(this.ChaserSequenceNumber, commandChannels)
		}
		// Reveal the selected sequence.
		common.RevealSequence(this.SelectedSequence, commandChannels)

		return

	case mode == CHASER_DISPLAY:

		if debug {
			fmt.Printf("displayMode: CHASER_DISPLAY\n")
		}
		// Hide the selected sequence.
		common.HideSequence(this.SelectedSequence, commandChannels)

		// Reveal the shutter chaser.
		//if this.ScannerChaser[this.SelectedSequence] && this.SelectedType == "scanner" {
		common.RevealSequence(this.ChaserSequenceNumber, commandChannels)
		//}

		return

	case mode == FUNCTION:

		if debug {
			fmt.Printf("displayMode: FUNCTION  Seq:%d Shutter Chaser is %t\n", this.SelectedSequence, this.ScannerChaser[this.SelectedSequence])
		}
		// If we have a shutter chaser running hide it.
		//if this.ScannerChaser[this.SelectedSequence] && this.SelectedType == "scanner" {
		common.HideSequence(this.ChaserSequenceNumber, commandChannels)
		//}
		// Hide the sequence.
		common.HideSequence(this.SelectedSequence, commandChannels)

		// Update the bottom buttons.
		common.UpdateBottomButtons("rgb", guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Speed %02d", this.Speed[this.ChaserSequenceNumber]), "speed", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Shift %02d", this.RGBShift[this.ChaserSequenceNumber]), "shift", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Size %02d", this.RGBSize[this.ChaserSequenceNumber]), "size", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Fade %02d", this.RGBFade[this.ChaserSequenceNumber]), "fade", false, guiButtons)

		// Show the function buttons.
		ShowFunctionButtons(this, eventsForLaunchpad, guiButtons)

		return

	case mode == CHASER_FUNCTION:

		if debug {
			fmt.Printf("displayMode: CHASER_FUNCTION\n")
		}
		// If we have a shutter chaser running hide it.
		if this.ScannerChaser[this.SelectedSequence] && this.SelectedType == "scanner" {
			common.HideSequence(this.ChaserSequenceNumber, commandChannels)
		}
		// Hide the normal sequence.
		common.HideSequence(this.SelectedSequence, commandChannels)

		// Update the labels.
		common.LabelButton(0, 7, "Chase\nSpeed\nDown", guiButtons)
		common.LabelButton(1, 7, "Chase\nSpeed\nUp", guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Chase Speed %02d", this.Speed[this.ChaserSequenceNumber]), "speed", false, guiButtons)
		common.LabelButton(2, 7, "Chase\nShift\nDown", guiButtons)
		common.LabelButton(3, 7, "Chase\nShift\nUp", guiButtons)
		common.LabelButton(4, 7, "Chase\nSize\nDown", guiButtons)
		common.LabelButton(5, 7, "Chase\nSize\nUp", guiButtons)
		common.LabelButton(6, 7, "Chase\nFade\nSoft", guiButtons)
		common.LabelButton(7, 7, "Chase\nFade\nSharp", guiButtons)

		// Show the chaser function buttons.
		this.TargetSequence = this.ChaserSequenceNumber
		ShowFunctionButtons(this, eventsForLaunchpad, guiButtons)

		return

	case mode == STATUS:

		if debug {
			fmt.Printf("displayMode: STATUS\n")
		}
		// If we're a scanner sequence and trying to display the status bar we don't want a shutter chaser in view.
		if this.ScannerChaser[this.SelectedSequence] && this.SelectedType == "scanner" {
			common.HideSequence(this.ChaserSequenceNumber, commandChannels)
		}
		// Hide the normal sequence.
		common.HideSequence(this.SelectedSequence, commandChannels)

		// Display the fixture status bar.
		showFixtureStatus(this.TargetSequence, sequences[this.SelectedSequence].Number, sequences[this.SelectedSequence].NumberFixtures, this, eventsForLaunchpad, guiButtons, commandChannels)

		return
	}

}

func showStatusBar(this *CurrentState, sequences []*common.Sequence, guiButtons chan common.ALight) {

	// Update the status bar
	common.UpdateStatusBar(fmt.Sprintf("Speed %02d", this.Speed[this.TargetSequence]), "speed", false, guiButtons)

	// Update the buttons: speed
	common.UpdateBottomButtons(this.SelectedType, guiButtons)

	if this.Strobe[this.SelectedSequence] {
		common.UpdateStatusBar(fmt.Sprintf("Strobe %02d", this.StrobeSpeed[this.DisplaySequence]), "speed", false, guiButtons)
	} else {
		if this.Functions[this.SelectedSequence][common.Function8_Music_Trigger].State {
			common.UpdateStatusBar("  MUSIC  ", "speed", false, guiButtons)
		} else {
			common.UpdateStatusBar(fmt.Sprintf("Speed %02d", this.Speed[this.DisplaySequence]), "speed", false, guiButtons)
		}
	}

	sensitivity := common.FindSensitivity(this.SoundGain)
	common.UpdateStatusBar(fmt.Sprintf("Sensitivity %02d", sensitivity), "sensitivity", false, guiButtons)
	common.UpdateStatusBar(fmt.Sprintf("Master %02d", this.MasterBrightness), "master", false, guiButtons)

	if sequences[this.TargetSequence].Type == "rgb" {
		common.UpdateStatusBar(fmt.Sprintf("Shift %02d", this.RGBShift[this.TargetSequence]), "shift", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Size %02d", this.RGBSize[this.TargetSequence]), "size", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Fade %02d", this.RGBFade[this.TargetSequence]), "fade", false, guiButtons)
		common.UpdateStatusBar("       ", "tilt", false, guiButtons)

		common.UpdateStatusBar(fmt.Sprintf("Red %02d", this.StaticButtons[this.TargetSequence].Color.R), "red", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Green %02d", this.StaticButtons[this.TargetSequence].Color.G), "green", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Blue %02d", this.StaticButtons[this.TargetSequence].Color.B), "blue", false, guiButtons)
	}
	if sequences[this.SelectedSequence].Type == "scanner" {
		label := getScannerShiftLabel(this.ScannerShift[this.TargetSequence])
		common.UpdateStatusBar(fmt.Sprintf("Shift %s", label), "shift", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Size %02d", this.ScannerSize[this.TargetSequence]), "size", false, guiButtons)
		label = getScannerCoordinatesLabel(this.ScannerCoordinates[this.TargetSequence])
		common.UpdateStatusBar(fmt.Sprintf("Coord %s", label), "fade", false, guiButtons)

		// Hide the color editing buttons.
		common.UpdateStatusBar(fmt.Sprintf("Tilt %02d", this.OffsetTilt), "tilt", false, guiButtons)
		common.UpdateStatusBar("        ", "red", false, guiButtons)
		common.UpdateStatusBar("        ", "green", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Pan %02d", this.OffsetPan), "pan", false, guiButtons)
	}
}

func hideAllFunctionKeys(this *CurrentState, sequences []*common.Sequence, eventsForLaunchPad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command) {

	for sequenceNumber := range sequences {
		if this.SelectMode[sequenceNumber] == FUNCTION {
			common.ClearSelectedRowOfButtons(sequenceNumber, eventsForLaunchPad, guiButtons)
			// And reveal all the other sequence that isn't us.
			if sequenceNumber != this.SelectedSequence {
				// And turn off the function selected.
				this.SelectMode[sequenceNumber] = NORMAL
				displayMode(NORMAL, this, sequences, eventsForLaunchPad, guiButtons, commandChannels)
			}
		}
	}
}

func printHandleDebug(this *CurrentState) {
	fmt.Printf("HANDLE: this.Type = %s \n", this.SelectedType)
	for functionNumber := 0; functionNumber < 8; functionNumber++ {
		state := this.Functions[this.TargetSequence][functionNumber]
		fmt.Printf("%d function %d state %+v\n", this.TargetSequence, functionNumber, state.State)
	}
	fmt.Printf("HANDLE:  SEQ: %d this.ScannerChaser running %t \n", this.DisplaySequence, this.ScannerChaser[this.DisplaySequence])

	fmt.Printf("================== WHAT SELECT MODE =================\n")
	fmt.Printf("HANDLE: this.SelectButtonPressed[%d] = %t \n", this.TargetSequence, this.SelectButtonPressed[this.TargetSequence])
	printMode(this)

	fmt.Printf("================== WHAT EDIT MODES =================\n")
	fmt.Printf("HANDLE: this.ShowRGBColorPicker[%d] = %t \n", this.TargetSequence, this.ShowRGBColorPicker)
	fmt.Printf("HANDLE: this.ShowStaticColorPicker[%d] = %t \n", this.TargetSequence, this.ShowStaticColorPicker)
	fmt.Printf("HANDLE: this.EditStaticColorsMode[%d] = %t \n", this.TargetSequence, this.EditStaticColorsMode)
	fmt.Printf("HANDLE: this.EditGoboSelectionMode[%d] = %t \n", this.TargetSequence, this.EditGoboSelectionMode)
	fmt.Printf("HANDLE: this.EditPatternMode[%d] = %t \n", this.TargetSequence, this.EditPatternMode)
	fmt.Printf("===============================================\n")
}
