package buttons

import (
	"fmt"

	"github.com/dhowlett99/dmxlights/pkg/common"
)

//		+-------------------+
//		|       NORMAL      |
//		+-------------------+
//		    |            |
//		    |            | If Scanner
//	     |            | or if the DisplayChaserShortCut is set.
//		    V            V
//		    |       +-------------------+
//		    |       |  CHASER DISPLAY   |
//		    |       +-------------------+
//		    V            |
//		+-------------------+
//		|     FUNCTION      |
//		+-------------------+
//		    |            | If Scanner
//		    V            V
//		    |       +-------------------+
//		    |       |  CHASER FUNCTIONS |
//		    |       +-------------------+
//		    |		     |
//		    V			 V
//		 +-------------------+
//		 |  FIXTURE STATUS   |
//		 +-------------------+
//		          |
//		          V
//		 +-------------------+
//		 |       NORMAL      |
//		 +-------------------+
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
		fmt.Printf("HANDLE: this.DisplayChaserShortCut = %t \n", this.DisplayChaserShortCut)
		fmt.Printf("HANDLE: this.EditSequenceColorPickerMode[%d] = %t \n", this.TargetSequence, this.EditSequenceColorPickerMode)
		fmt.Printf("HANDLE: this.EditStaticColorsMode[%d] = %t \n", this.TargetSequence, this.EditStaticColorsMode)
		fmt.Printf("HANDLE: this.EditGoboSelectionMode[%d] = %t \n", this.TargetSequence, this.EditGoboSelectionMode)
		fmt.Printf("HANDLE: this.EditPatternMode[%d] = %t \n", this.TargetSequence, this.EditPatternMode)
		fmt.Printf("HANDLE: this.EditColorPicker[%d] = %t \n", this.TargetSequence, this.EditColorPicker)
		fmt.Printf("===============================================\n")
	}

	// Update the status bar
	if this.Strobe[this.SelectedSequence] {
		common.UpdateStatusBar(fmt.Sprintf("Strobe %02d", this.StrobeSpeed[this.DisplaySequence]), "speed", false, guiButtons)
	} else {
		// Update status bar.
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

	// Light the top buttons.
	common.ShowTopButtons(sequences[this.DisplaySequence].Type, eventsForLaunchpad, guiButtons)

	// Light the strobe button.
	common.ShowStrobeButtonStatus(this.Strobe[this.DisplaySequence], eventsForLaunchpad, guiButtons)

	// Light the start stop button.
	common.ShowRunningStatus(this.DisplaySequence, this.Running, eventsForLaunchpad, guiButtons)

	// 1st Press Select Sequence - This the first time we have pressed the select button.
	// Simply select the selected sequence.
	// But remember we have pressed this select button once.
	if this.SelectMode[this.DisplaySequence] == NORMAL && !this.SelectButtonPressed[this.DisplaySequence] || // Normal mode and first time pressed.
		// OR we're in the CHASER_DISPLAY mode with the chaser on and the DisplayChaserShortCut has fired
		this.SelectMode[this.SelectedSequence] == CHASER_DISPLAY && this.ScannerChaser[this.SelectedSequence] && this.DisplayChaserShortCut {

		if debug {
			fmt.Printf("%d: Show Sequence - Handle Step 1\n", this.SelectedSequence)
		}

		// OK this is complicated,, but if we have switched on the shutter chaser from the scanner sequence function key 7 "Scanner Shutter Chase"
		// We would at arrive at Step 2 below, where we would have switched the mode to CHASER_DISPLAY and force the display to show the shutter chaser.
		// What we are doing here is using the DisplayChaserShortCut flag to detect that this has happened and force the next step in the select press sequence to
		// go back to the NORMAL mode and resume the sequence of select presses to as described in the header of this Handle() function.
		if this.DisplayChaserShortCut {
			this.SelectMode[this.DisplaySequence] = NORMAL
			// And now forget this ever happened. Well untill the next shortcut is called.
			this.DisplayChaserShortCut = false
		}

		// Flash all static fixtures. Represents all fixtures selected.
		if !this.SelectAllStaticFixtures && this.EditStaticColorsMode[this.DisplaySequence] && !this.EditColorPicker {
			if debug {
				fmt.Printf("%d: Select All\n", this.DisplaySequence)
			}
			this.SelectAllStaticFixtures = true
			// Update all the fixtures so they will flash.
			cmd := common.Command{
				Action: common.UpdateFlashAllStaticColorButtons,
				Args: []common.Arg{
					{Name: "StaticFlash", Value: this.SelectAllStaticFixtures},
				},
			}
			if this.SelectedType == "scanner" && this.ScannerChaser[this.DisplaySequence] {
				common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)
			} else {
				if this.TargetSequence != 2 {
					common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)
				}
			}
		}

		// Assume everything else is off.
		this.SelectButtonPressed[0] = false
		this.SelectButtonPressed[1] = false
		this.SelectButtonPressed[2] = false
		this.SelectButtonPressed[3] = false

		// Remember which select button has been pressed.
		this.SelectButtonPressed[this.SelectedSequence] = true

		// Turn off the color picker if set.
		if this.EditColorPicker {
			removeColorPicker(this, eventsForLaunchpad, guiButtons, commandChannels)
			this.EditColorPicker = false

			// Turn off function mode. Remove the function pads.
			common.ClearSelectedRowOfButtons(this.SelectedSequence, eventsForLaunchpad, guiButtons)

			// If the chaser is running, reveal it.
			if this.ScannerChaser[this.SelectedSequence] && this.SelectedType == "scanner" && this.SelectMode[this.SelectedSequence] == CHASER_DISPLAY {
				common.HideSequence(this.SelectedSequence, commandChannels)
				common.RevealSequence(this.ChaserSequenceNumber, commandChannels)
				this.SelectMode[this.SelectedSequence] = CHASER_DISPLAY
			} else {
				common.HideSequence(this.ChaserSequenceNumber, commandChannels)
				// And reveal the sequence on the launchpad keys
				common.RevealSequence(this.SelectedSequence, commandChannels)
				this.SelectMode[this.SelectedSequence] = NORMAL
			}
		}

		// Turn off any previous function or status bars.
		for sequenceNumber := range sequences {
			if this.SelectMode[sequenceNumber] == FUNCTION ||
				this.SelectMode[sequenceNumber] == STATUS {
				common.ClearSelectedRowOfButtons(sequenceNumber, eventsForLaunchpad, guiButtons)
				// And reveal all the other sequence that isn't us.
				if sequenceNumber != this.SelectedSequence {
					if this.ScannerChaser[this.SelectedSequence] {
						common.HideSequence(this.ChaserSequenceNumber, commandChannels)
					}
					common.RevealSequence(sequenceNumber, commandChannels)
					// And turn off the function selected.
					this.SelectMode[sequenceNumber] = NORMAL
				}
			}
		}

		if this.Functions[this.SelectedSequence][common.Function1_Pattern].State {
			// Reset the pattern function key.
			this.Functions[this.SelectedSequence][common.Function1_Pattern].State = false

			// Clear the pattern function keys
			ClearPatternSelectionButtons(this.SelectedSequence, *sequences[this.SelectedSequence], eventsForLaunchpad, guiButtons)

			// And reveal the sequence.
			if this.ScannerChaser[this.SelectedSequence] {
				common.HideSequence(this.ChaserSequenceNumber, commandChannels)
			}
			common.RevealSequence(this.SelectedSequence, commandChannels)

			// Editing pattern is over for this sequence.
			this.EditPatternMode = false

			// Clear buttons and remove any labels.
			common.ClearSelectedRowOfButtons(this.SelectedSequence, eventsForLaunchpad, guiButtons)
		}

		if this.SelectMode[this.SelectedSequence] == NORMAL &&
			this.Functions[this.EditWhichStaticSequence][common.Function5_Color].State &&
			this.EditSequenceColorPickerMode {

			fmt.Printf("Color Edit Mode Off for sequence %d\n", this.SelectedSequence)

			// Reset the color function key.
			this.Functions[this.EditWhichStaticSequence][common.Function5_Color].State = false

			// And reveal the sequence on the launchpad keys
			// and hide the shutter chaser.
			if this.ScannerChaser[this.SelectedSequence] && this.SelectedType == "scanner" {
				common.HideSequence(this.ChaserSequenceNumber, commandChannels)
			}
			common.RevealSequence(this.SelectedSequence, commandChannels)

			// Turn off the function mode flag.
			this.SelectMode[this.SelectedSequence] = NORMAL

			// Now forget we pressed twice and start again.
			this.SelectButtonPressed[this.SelectedSequence] = true

			common.HideColorSelectionButtons(this.SelectedSequence, *sequences[this.SelectedSequence], eventsForLaunchpad, guiButtons)

		}

		// Tailor the top buttons to the sequence type.
		common.ShowTopButtons(sequences[this.SelectedSequence].Type, eventsForLaunchpad, guiButtons)

		// Tailor the bottom buttons to the sequence type.
		common.ShowBottomButtons(sequences[this.SelectedSequence].Type, eventsForLaunchpad, guiButtons)

		// Show this sequence running status in the start/stop button.
		common.ShowRunningStatus(this.SelectedSequence, this.Running, eventsForLaunchpad, guiButtons)

		// Clear the select unless we're in static mode for this sequence.
		if !this.EditStaticColorsMode[this.SelectedSequence] {
			common.ClearSelectedRowOfButtons(this.SelectedSequence, eventsForLaunchpad, guiButtons)
		}

		// Now select the correct exit mode for scanner sequences.
		if this.ScannerChaser[this.SelectedSequence] && this.SelectedType == "scanner" {
			if this.SelectMode[this.SelectedSequence] == NORMAL {
				common.HideSequence(this.ChaserSequenceNumber, commandChannels)
				common.RevealSequence(this.SelectedSequence, commandChannels)
			}
		}
		if debug {
			printMode(this)
		}
		// Light the sequence selector button.
		SequenceSelect(eventsForLaunchpad, guiButtons, this)
		return
	}

	// First option of the 2nd Press in scanner mode go into Shutter Chaser Mode for this sequence.
	// We're in Normal mode, We've Pressed twice and the scanner chaser is on and we're a scanner sequence.
	// Or the DisplayChaserShortCut has fired.
	if (this.SelectMode[this.SelectedSequence] == NORMAL &&
		this.ScannerChaser[this.SelectedSequence] &&
		this.SelectedType == "scanner" &&
		this.SelectButtonPressed[this.SelectedSequence]) || this.DisplayChaserShortCut {

		if debug {
			fmt.Printf("%d: We are a SCANNER - 2nd Press Shutter Chase Mode - Handle Step 2\n", this.SelectedSequence)
		}

		// Set function mode. And take note if we arrived using the shortcut.
		if this.DisplayChaserShortCut {
			if debug {
				fmt.Printf("%d: Chaser Display entered via DisplayChaserShortCut\n", this.SelectedSequence)
			}
			this.SelectButtonPressed[this.SelectedSequence] = false
		}

		// Update the bottom buttons.
		common.UpdateBottomButtons("rgb", guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Speed %02d", this.Speed[this.ChaserSequenceNumber]), "speed", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Shift %02d", this.RGBShift[this.ChaserSequenceNumber]), "shift", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Size %02d", this.RGBSize[this.ChaserSequenceNumber]), "size", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Fade %02d", this.RGBFade[this.ChaserSequenceNumber]), "fade", false, guiButtons)

		// Set the Shutter Chaser mode.
		this.SelectMode[this.SelectedSequence] = CHASER_DISPLAY

		// Hide the rotating sequence.
		common.HideSequence(this.SelectedSequence, commandChannels)

		// Clear the sequence.
		common.ClearSelectedRowOfButtons(this.SelectedSequence, eventsForLaunchpad, guiButtons)

		// And show the chaser sequence.
		common.RevealSequence(this.ChaserSequenceNumber, commandChannels)

		if debug {
			printMode(this)
		}

		// Light the sequence selector button.
		SequenceSelect(eventsForLaunchpad, guiButtons, this)
		return
	}

	// Second option of the 2nd Press in rgb mode same select button go into Function Mode for this sequence.
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

		// Set function mode.
		this.SelectMode[this.SelectedSequence] = FUNCTION

		// Toggle the select all static fixuure off.
		if this.SelectAllStaticFixtures {
			this.SelectAllStaticFixtures = false
			// Update all the fixtures so they will stop flashing.
			cmd := common.Command{
				Action: common.UpdateFlashAllStaticColorButtons,
				Args: []common.Arg{
					{Name: "StaticFlash", Value: this.SelectAllStaticFixtures},
				},
			}
			common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)
		}

		// And hide the sequence so we can only see the function buttons.
		common.HideSequence(this.SelectedSequence, commandChannels)

		// If the shutter chaser is running, hide it.
		if this.ScannerChaser[this.SelectedSequence] && this.SelectedType == "scanner" {
			common.HideSequence(this.ChaserSequenceNumber, commandChannels)
		}

		// Turn off any previous function bars.
		for sequenceNumber := range sequences {
			if this.SelectMode[sequenceNumber] == FUNCTION {
				common.ClearSelectedRowOfButtons(sequenceNumber, eventsForLaunchpad, guiButtons)
				// And reveal all the other sequence that isn't us.
				if sequenceNumber != this.SelectedSequence {
					common.RevealSequence(sequenceNumber, commandChannels)
					// And turn off the function selected.
					this.SelectMode[sequenceNumber] = NORMAL
				}
			}
		}

		// Clear the buttons.
		common.ClearSelectedRowOfButtons(this.SelectedSequence, eventsForLaunchpad, guiButtons)

		// If we're a scanner sequence and trying to display the function bar we don't want a shutter chaser in view.
		if !this.ScannerChaser[this.SelectedSequence] && this.SelectedType == "scanner" {
			common.HideSequence(this.ChaserSequenceNumber, commandChannels)
		}

		// Show the function buttons.
		ShowFunctionButtons(this, eventsForLaunchpad, guiButtons)

		// Now forget we pressed twice and start again.
		this.SelectButtonPressed[this.SelectedSequence] = false

		if debug {
			printMode(this)
		}

		// Light the sequence selector button.
		SequenceSelect(eventsForLaunchpad, guiButtons, this)
		return
	}

	// 3rd Press Status Mode and not a scanner - we display the fixture status enable/invert/disable buttons.
	// We're in Function mode, pressed twice. and we're in edit sequence color mode. and we're NOT a scanner.
	if this.SelectMode[this.SelectedSequence] == FUNCTION &&
		!this.SelectButtonPressed[this.SelectedSequence] &&
		!this.EditSequenceColorPickerMode &&
		this.SelectedType != "scanner" {

		if debug {
			fmt.Printf("%d: Handle 3 - Status Buttons on\n", this.SelectedSequence)
		}

		// Turn on status mode
		this.SelectMode[this.SelectedSequence] = STATUS

		// If the chase is running, hide it.
		if this.ScannerChaser[this.SelectedSequence] && this.SelectedType == "scanner" {
			common.HideSequence(this.ChaserSequenceNumber, commandChannels)
		}

		// Remove the function pads.
		common.ClearSelectedRowOfButtons(this.SelectedSequence, eventsForLaunchpad, guiButtons)

		// Show the Fixture Status Buttons.
		showFixtureStatus(this.TargetSequence, sequences[this.SelectedSequence].Number, sequences[this.SelectedSequence].NumberFixtures, this, eventsForLaunchpad, guiButtons, commandChannels)

		if debug {
			printMode(this)
		}

		// Light the sequence selector button.
		SequenceSelect(eventsForLaunchpad, guiButtons, this)
		return
	}

	// 3rd Press Status Mode and we are scanner and the shutter chaser is running - we display the shutter chaser function buttons.
	if this.SelectMode[this.SelectedSequence] == FUNCTION &&
		this.ScannerChaser[this.SelectedSequence] &&
		!this.SelectButtonPressed[this.SelectedSequence] &&
		!this.EditSequenceColorPickerMode &&
		this.SelectedType == "scanner" {

		if debug {
			fmt.Printf("%d: Handle Step 4 Shutter Chase Function buttons on\n", this.SelectedSequence)
		}

		// Turn on shutter chaser mode.
		this.SelectMode[this.SelectedSequence] = CHASER_FUNCTION

		// Update the buttons: speed
		common.LabelButton(0, 7, "Chase\nSpeed\nDown", guiButtons)
		common.LabelButton(1, 7, "Chase\nSpeed\nUp", guiButtons)

		// Update the status bar
		common.UpdateStatusBar(fmt.Sprintf("Chase Speed %02d", this.Speed[this.ChaserSequenceNumber]), "speed", false, guiButtons)

		common.LabelButton(2, 7, "Chase\nShift\nDown", guiButtons)
		common.LabelButton(3, 7, "Chase\nShift\nUp", guiButtons)

		common.LabelButton(4, 7, "Chase\nSize\nDown", guiButtons)
		common.LabelButton(5, 7, "Chase\nSize\nUp", guiButtons)

		common.LabelButton(6, 7, "Chase\nFade\nSoft", guiButtons)
		common.LabelButton(7, 7, "Chase\nFade\nSharp", guiButtons)

		// If the chase is running, hide it.
		if this.ScannerChaser[this.SelectedSequence] && this.SelectedType == "scanner" {
			common.HideSequence(this.ChaserSequenceNumber, commandChannels)
		}

		// Remove the function pads.
		common.ClearSelectedRowOfButtons(this.DisplaySequence, eventsForLaunchpad, guiButtons)

		this.TargetSequence = this.ChaserSequenceNumber

		// Show the function buttons.
		ShowFunctionButtons(this, eventsForLaunchpad, guiButtons)

		if debug {
			printMode(this)
		}

		// Light the sequence selector button.
		SequenceSelect(eventsForLaunchpad, guiButtons, this)
		return
	}

	// 4th Press Normal Mode - we head back to normal mode.
	if this.SelectMode[this.SelectedSequence] == STATUS &&
		!this.SelectButtonPressed[this.SelectedSequence] &&
		!this.EditSequenceColorPickerMode {

		if debug {
			fmt.Printf("%d: Handle Step 5 - Normal Mode From Non Scanner, Function Bar off\n", this.SelectedSequence)
		}

		// Turn off function mode.
		this.SelectMode[this.SelectedSequence] = NORMAL

		// Remove the status buttons.
		common.ClearSelectedRowOfButtons(this.SelectedSequence, eventsForLaunchpad, guiButtons)

		// We're in Edit Pattern Mode.
		if this.Functions[this.SelectedSequence][common.Function1_Pattern].State {
			if debug {
				fmt.Printf("Show Pattern Selection Buttons\n")
			}
			this.EditPatternMode = true
			common.HideSequence(this.SelectedSequence, commandChannels)
			ShowPatternSelectionButtons(this, sequences[this.SelectedSequence].Master, *sequences[this.SelectedSequence], this.DisplaySequence, eventsForLaunchpad, guiButtons)
			// Light the sequence selector button.
			SequenceSelect(eventsForLaunchpad, guiButtons, this)
			return
		}

		// We're in RGB Color Selection Mode.
		if this.SelectedType == "rgb" && this.Functions[this.EditWhichStaticSequence][common.Function5_Color].State {
			if debug {
				fmt.Printf("Show RGB Sequence Color Selection Buttons\n")
			}
			// Turn off the color selection function key.
			this.Functions[this.EditWhichStaticSequence][common.Function5_Color].State = false
			// Set the colors.
			sequences[this.EditWhichStaticSequence].CurrentColors = sequences[this.EditWhichStaticSequence].SequenceColors
			// Show the colors
			ShowRGBColorPicker(this.MasterBrightness, *sequences[this.EditWhichStaticSequence], this.DisplaySequence, eventsForLaunchpad, guiButtons, commandChannels)
			// Light the sequence selector button.
			SequenceSelect(eventsForLaunchpad, guiButtons, this)
			return
		}

		// We're in RGB Static Color Mode.
		var static bool
		// Check all sequences to see if one is static.
		for sequenceNumber, isStatic := range this.EditStaticColorsMode {
			if isStatic {
				if debug {
					fmt.Printf("Show RGB Static Colors for sequence %d\n", sequenceNumber)
				}
				//common.SetMode(sequenceNumber, commandChannels, "Static")
				common.RevealSequence(sequenceNumber, commandChannels)
				//this.EditStaticColorsMode = false
				static = true
			}
		}
		if static {
			// Light the sequence selector button.
			SequenceSelect(eventsForLaunchpad, guiButtons, this)
			if debug {
				printMode(this)
			}
			// Turn off the function mode flag.
			this.SelectMode[this.SelectedSequence] = NORMAL
			// If the chase is running, hide it.
			if this.ScannerChaser[this.SelectedSequence] && this.SelectedType == "scanner" {
				if debug {
					fmt.Printf("%d: Hide Sequence\n", this.ChaserSequenceNumber)
				}
				common.HideSequence(this.ChaserSequenceNumber, commandChannels)
			}
			// Remember that we've preseed twice.
			this.SelectButtonPressed[this.SelectedSequence] = true
			return
		}

		// We're in Scanner Gobo Selection Mode.
		if this.Functions[this.SelectedSequence][common.Function6_Static_Gobo].State &&
			!this.EditStaticColorsMode[this.EditWhichStaticSequence] &&
			sequences[this.SelectedSequence].Type == "scanner" {
			this.Functions[this.SelectedSequence][common.Function6_Static_Gobo].State = false
			this.EditGoboSelectionMode = false
		}

		// Allow us to exit the pattern select mode without setting a pattern.
		if this.EditPatternMode {
			this.EditPatternMode = false
		}

		// Else reveal the sequence on the launchpad keys
		if debug {
			fmt.Printf("%d: Reveal Sequence\n", this.SelectedSequence)
		}
		common.RevealSequence(this.SelectedSequence, commandChannels)

		// If the chase is running, reveal it. But only if we're in CHASER_DISPLAY mode.
		if this.ScannerChaser[this.SelectedSequence] && this.SelectedType == "scanner" && this.SelectMode[this.SelectedSequence] == CHASER_DISPLAY {
			if debug {
				fmt.Printf("%d: Reveal Sequence\n", this.ChaserSequenceNumber)
			}
			common.RevealSequence(this.ChaserSequenceNumber, commandChannels)
		}

		// Turn off the function mode flag.
		this.SelectMode[this.SelectedSequence] = NORMAL
		// Remember that we've preseed twice.
		this.SelectButtonPressed[this.SelectedSequence] = true

		if debug {
			printMode(this)
		}
		// Light the sequence selector button.
		SequenceSelect(eventsForLaunchpad, guiButtons, this)
		return
	}

	// 4th Press Normal Mode and we are a scanner- we head fixture status mode.
	if (this.SelectMode[this.SelectedSequence] == CHASER_FUNCTION || !this.ScannerChaser[this.SelectedSequence]) &&
		!this.SelectButtonPressed[this.SelectedSequence] &&
		!this.EditSequenceColorPickerMode &&
		this.SelectedType == "scanner" {

		if debug {
			fmt.Printf("%d: Handle Step 6 Normal Mode, From  Scanner, Function Bar off, status buttons on\n", this.SelectedSequence)
		}

		// Turn on status mode
		this.SelectMode[this.SelectedSequence] = STATUS

		// Update the status bar
		common.UpdateStatusBar(fmt.Sprintf("Speed %02d", this.Speed[this.TargetSequence]), "speed", false, guiButtons)

		// Update the buttons: speed
		common.UpdateBottomButtons(this.SelectedType, guiButtons)

		// If the chase is running, hide it.
		if this.ScannerChaser[this.SelectedSequence] && this.SelectedType == "scanner" {
			common.HideSequence(this.ChaserSequenceNumber, commandChannels)
		}

		// Turn off the edit sequence colors button.
		if this.EditSequenceColorPickerMode {
			this.EditSequenceColorPickerMode = false
			this.Functions[this.EditWhichStaticSequence][common.Function5_Color].State = false
		}

		// Remove the function pads.
		common.ClearSelectedRowOfButtons(this.SelectedSequence, eventsForLaunchpad, guiButtons)

		// Show the Fixture Status Buttons.
		showFixtureStatus(this.TargetSequence, sequences[this.SelectedSequence].Number, sequences[this.SelectedSequence].NumberFixtures, this, eventsForLaunchpad, guiButtons, commandChannels)

		if debug {
			printMode(this)
		}
		// Light the sequence selector button.
		SequenceSelect(eventsForLaunchpad, guiButtons, this)
		return
	}

	// Are we in function mode ?
	if this.SelectMode[this.SelectedSequence] == FUNCTION || this.SelectMode[this.SelectedSequence] == CHASER_FUNCTION {
		if debug {
			fmt.Printf("HANDLE %d: Handle 7 - Back to NORMAL mode.\n", this.SelectedSequence)
		}
		// Turn off function mode. Remove the function pads.
		common.ClearSelectedRowOfButtons(this.SelectedSequence, eventsForLaunchpad, guiButtons)

		// And reveal the sequence on the launchpad keys
		common.RevealSequence(this.SelectedSequence, commandChannels)

		// If the chaser is running, reveal it.
		if this.ScannerChaser[this.SelectedSequence] && this.SelectedType == "scanner" {
			common.RevealSequence(this.ChaserSequenceNumber, commandChannels)
		}

		// Now select the correct exit mode based on which mode we came in on.
		if this.SelectMode[this.SelectedSequence] == FUNCTION {
			this.SelectMode[this.SelectedSequence] = NORMAL
		}
		if this.SelectMode[this.SelectedSequence] == CHASER_FUNCTION {
			this.SelectMode[this.SelectedSequence] = CHASER_DISPLAY
		}

		// Turn off the edit sequence colors button.
		if this.EditSequenceColorPickerMode {
			removeColorPicker(this, eventsForLaunchpad, guiButtons, commandChannels)
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

		// Now forget we pressed twice and start again.
		this.SelectButtonPressed[this.SelectedSequence] = false

		if debug {
			printMode(this)
		}
		// Light the sequence selector button.
		SequenceSelect(eventsForLaunchpad, guiButtons, this)
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

	// This doesn't take account of the shutter chaser which should be shown if the chaser is running.
	common.SendCommandToSequence(0, cmd, commandChannels)
	common.SendCommandToSequence(1, cmd, commandChannels)
	if !this.ScannerChaser[this.SelectedSequence] {
		common.SendCommandToSequence(2, cmd, commandChannels)
	} else {
		common.SendCommandToSequence(4, cmd, commandChannels)
	}

}
