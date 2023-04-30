package buttons

import (
	"fmt"

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/sequence"
)

//	+-------------------+
//	|       NORMAL      |
//	+-------------------+
//	         |
//	         V
//	+-------------------+
//	|     FUNCTION      |
//	+-------------------+
//	    |            | If Scanner
//	    V            V
//	    |       +-------------------+
//	    |       |  CHASER FUNCTIONS |
//	    |       +-------------------+
//	    |				|
//	    V				V
//	 +-------------------+
//	 |       STATUS      |
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

	if this.SelectMode[this.SelectedSequence] == CHASER {
		this.TargetSequence = this.ChaserSequenceNumber
		this.DisplaySequence = this.SelectedSequence
	} else {
		this.TargetSequence = this.SelectedSequence
		this.DisplaySequence = this.SelectedSequence
	}

	if debug {
		fmt.Printf("HANDLE: this.Type = %s \n", this.SelectedType)
		for functionNumber := 0; functionNumber < 8; functionNumber++ {
			state := this.Functions[this.TargetSequence][functionNumber].State
			fmt.Printf("HANDLE: function %d state %t\n", functionNumber, state)
		}
		fmt.Printf("HANDLE: this.ChaserRunning %t \n", this.ScannerChaser)

		fmt.Printf("================== WHAT SELECT MODE =================\n")
		fmt.Printf("HANDLE: this.SelectButtonPressed[%d] = %t \n", this.TargetSequence, this.SelectButtonPressed[this.TargetSequence])
		if this.SelectMode[this.DisplaySequence] == NORMAL {
			fmt.Printf("HANDLE: this.SelectMode[%d] = NORMAL \n", this.TargetSequence)
		}
		if this.SelectMode[this.DisplaySequence] == FUNCTION {
			fmt.Printf("HANDLE: this.SelectMode[%d] = FUNCTION \n", this.SelectedSequence)
		}
		if this.SelectMode[this.DisplaySequence] == CHASER {
			fmt.Printf("HANDLE: this.SelectMode[%d] = CHASER \n", this.SelectedSequence)
		}
		if this.SelectMode[this.DisplaySequence] == STATUS {
			fmt.Printf("HANDLE: this.SelectMode[%d] = STATUS \n", this.SelectedSequence)
		}

		fmt.Printf("================== WHAT EDIT MODES =================\n")
		fmt.Printf("HANDLE: this.EditSequenceColorsMode[%d] = %t \n", this.TargetSequence, this.EditSequenceColorsMode)
		fmt.Printf("HANDLE: this.EditStaticColorsMode[%d] = %t \n", this.TargetSequence, this.EditStaticColorsMode)
		fmt.Printf("HANDLE: this.EditGoboSelectionMode[%d] = %t \n", this.TargetSequence, this.EditGoboSelectionMode)
		fmt.Printf("HANDLE: this.EditPatternMode[%d] = %t \n", this.TargetSequence, this.EditPatternMode)
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

	// Light the sequence selector button.
	sequence.SequenceSelect(eventsForLaunchpad, guiButtons, this.DisplaySequence)

	// Light the strobe button.
	common.ShowStrobeButtonStatus(this.Strobe[this.DisplaySequence], eventsForLaunchpad, guiButtons)

	// Light the start stop button.
	common.ShowRunningStatus(this.DisplaySequence, this.Running, eventsForLaunchpad, guiButtons)

	// 1st Press Select Sequence - This the first time we have pressed the select button.
	// Simply select the selected sequence.
	// But remember we have pressed this select button once.
	if this.SelectMode[this.DisplaySequence] == NORMAL &&
		!this.SelectButtonPressed[this.DisplaySequence] {
		if debug {
			fmt.Printf("%d: Show Sequence - Handle Step 1\n", this.SelectedSequence)
		}

		// Assume everything else is off.
		this.SelectButtonPressed[0] = false
		this.SelectButtonPressed[1] = false
		this.SelectButtonPressed[2] = false
		this.SelectButtonPressed[3] = false

		// But remember we have pressed this select button once.
		this.SelectMode[this.SelectedSequence] = NORMAL
		this.SelectButtonPressed[this.SelectedSequence] = true

		// Turn off any previous function or status bars.
		for sequenceNumber := range sequences {
			if this.SelectMode[sequenceNumber] == FUNCTION ||
				this.SelectMode[sequenceNumber] == STATUS {
				common.ClearSelectedRowOfButtons(sequenceNumber, eventsForLaunchpad, guiButtons)
				// And reveal all the other sequence that isn't us.
				if sequenceNumber != this.SelectedSequence {
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
			common.RevealSequence(this.SelectedSequence, commandChannels)

			// If the chase is running, reveal it.
			if this.ScannerChaser && this.SelectedType == "scanner" {
				common.RevealSequence(this.ChaserSequenceNumber, commandChannels)
			}

			// Editing pattern is over for this sequence.
			this.EditPatternMode = false

			// Clear buttons and remove any labels.
			common.ClearSelectedRowOfButtons(this.SelectedSequence, eventsForLaunchpad, guiButtons)
		}

		if this.SelectMode[this.SelectedSequence] == NORMAL &&
			this.Functions[this.SelectedSequence][common.Function5_Color].State && this.EditSequenceColorsMode {
			unSetEditSequenceColorsMode(sequences, this, commandChannels, eventsForLaunchpad, guiButtons)
		}

		// Tailor the top buttons to the sequence type.
		common.ShowTopButtons(sequences[this.SelectedSequence].Type, eventsForLaunchpad, guiButtons)

		// Tailor the bottom buttons to the sequence type.
		common.ShowBottomButtons(sequences[this.SelectedSequence].Type, eventsForLaunchpad, guiButtons)

		// Show this sequence running status in the start/stop button.
		common.ShowRunningStatus(this.SelectedSequence, this.Running, eventsForLaunchpad, guiButtons)

		return
	}

	// 2nd Press same select button go into Function Mode for this sequence.
	if this.SelectMode[this.SelectedSequence] == NORMAL && sequences[this.SelectedSequence].Type != "switch" || // Don't alow functions in switch mode.
		this.SelectMode[this.SelectedSequence] == NORMAL && // Function select mode is off
			this.EditStaticColorsMode[this.TargetSequence] && // AND static colol mode, the case when we leave static colors edit mode.
			this.SelectButtonPressed[this.SelectedSequence] {

		if debug {
			fmt.Printf("%d: 2nd Press Function Bar Mode - Handle Step 2\n", this.SelectedSequence)
		}

		// Set function mode.
		this.SelectMode[this.SelectedSequence] = FUNCTION

		// If static, show static colors.
		if this.EditStaticColorsMode[this.TargetSequence] {
			if debug {
				fmt.Printf("Show Static Color Selection Buttons\n")
			}
			common.SetMode(this.TargetSequence, commandChannels, "Static")
			//this.EditStaticColorsMode = false
		}

		// And hide the sequence so we can only see the function buttons.
		common.HideSequence(this.SelectedSequence, commandChannels)

		// If the chase is running, hide it.
		if this.ScannerChaser && this.SelectedType == "scanner" {
			common.HideSequence(this.ChaserSequenceNumber, commandChannels)
		}

		// Turn off any static sequence so we can see the functions.
		common.SetMode(this.SelectedSequence, commandChannels, "Sequence")

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

		// Show the function buttons.
		ShowFunctionButtons(this, eventsForLaunchpad, guiButtons)

		// Now forget we pressed twice and start again.
		this.SelectButtonPressed[this.SelectedSequence] = false

		return
	}

	// 3rd Press Status Mode and not a scanner - we display the fixture status enable/invert/disable buttons.
	if this.SelectMode[this.SelectedSequence] == FUNCTION &&
		!this.SelectButtonPressed[this.SelectedSequence] &&
		!this.EditSequenceColorsMode &&
		this.SelectedType != "scanner" {

		if debug {
			fmt.Printf("%d: Handle Step 3 Status Mode, Not a Scanner, Function Bar off, status buttons on\n", this.SelectedSequence)
		}

		// Turn on status mode
		this.SelectMode[this.SelectedSequence] = STATUS

		// If the chase is running, hide it.
		if this.ScannerChaser && this.SelectedType == "scanner" {
			common.HideSequence(this.ChaserSequenceNumber, commandChannels)
		}

		// Remove the function pads.
		common.ClearSelectedRowOfButtons(this.SelectedSequence, eventsForLaunchpad, guiButtons)

		// Show the Fixture Status Buttons.
		ShowFixtureStatus(this.TargetSequence, *sequences[this.SelectedSequence], this, eventsForLaunchpad, guiButtons, commandChannels)

		return
	}

	// 3rd Press Status Mode and we are scanner - we display the shutter chaser function buttons.
	if this.SelectMode[this.SelectedSequence] == FUNCTION &&
		!this.SelectButtonPressed[this.SelectedSequence] &&
		!this.EditSequenceColorsMode &&
		this.SelectedType == "scanner" {

		if debug {
			fmt.Printf("%d: Handle Step 4 Status Mode, We are a Scanner, Function Bar off, shutter chase function buttons on\n", this.SelectedSequence)
		}

		// Turn on shutter chaser mode.
		this.SelectMode[this.SelectedSequence] = CHASER

		// If the chase is running, hide it.
		if this.ScannerChaser && this.SelectedType == "scanner" {
			common.HideSequence(this.ChaserSequenceNumber, commandChannels)
		}

		// Remove the function pads.
		common.ClearSelectedRowOfButtons(this.DisplaySequence, eventsForLaunchpad, guiButtons)

		this.TargetSequence = this.ChaserSequenceNumber

		// Show the function buttons.
		ShowFunctionButtons(this, eventsForLaunchpad, guiButtons)

		return
	}

	// 4th Press Normal Mode - we head back to normal mode.
	if this.SelectMode[this.SelectedSequence] == STATUS &&
		!this.SelectButtonPressed[this.SelectedSequence] &&
		!this.EditSequenceColorsMode {

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
			return
		}

		// We're in RGB Color Selection Mode.
		if this.Functions[this.SelectedSequence][common.Function5_Color].State && sequences[this.SelectedSequence].Type == "rgb" {
			if debug {
				fmt.Printf("Show RGB Sequence Color Selection Buttons\n")
			}
			// Set the colors.
			sequences[this.SelectedSequence].CurrentColors = sequences[this.SelectedSequence].SequenceColors
			// Show the colors
			ShowRGBColorSelectionButtons(this.MasterBrightness, *sequences[this.SelectedSequence], this.DisplaySequence, eventsForLaunchpad, guiButtons)
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
				common.SetMode(sequenceNumber, commandChannels, "Static")
				common.RevealSequence(sequenceNumber, commandChannels)
				//this.EditStaticColorsMode = false
				static = true
			}
		}
		if static {
			return
		}

		// We're in Scanner Gobo Selection Mode.
		if this.Functions[this.SelectedSequence][common.Function6_Static_Gobo].State &&
			!this.EditStaticColorsMode[this.EditWhichSequenceStatic] &&
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

		// If the chase is running, reveal it.
		if this.ScannerChaser && this.SelectedType == "scanner" {
			if debug {
				fmt.Printf("%d: Reveal Sequence\n", this.ChaserSequenceNumber)
			}
			common.RevealSequence(this.ChaserSequenceNumber, commandChannels)
		}

		// Turn off the function mode flag.
		this.SelectMode[this.SelectedSequence] = NORMAL
		// Now forget we pressed twice and start again.
		this.SelectButtonPressed[this.SelectedSequence] = true

		return
	}

	// 4th Press Normal Mode and we are a scanner- we head fixture status mode.
	if this.SelectMode[this.SelectedSequence] == CHASER &&
		!this.SelectButtonPressed[this.SelectedSequence] &&
		!this.EditSequenceColorsMode &&
		this.SelectedType == "scanner" {

		if debug {
			fmt.Printf("%d: Handle Step 6 Normal Mode, From  Scanner, Function Bar off, status buttons on\n", this.SelectedSequence)
		}

		// Turn on status mode
		this.SelectMode[this.SelectedSequence] = STATUS

		// If the chase is running, hide it.
		if this.ScannerChaser && this.SelectedType == "scanner" {
			common.HideSequence(this.ChaserSequenceNumber, commandChannels)
		}

		// Remove the function pads.
		common.ClearSelectedRowOfButtons(this.SelectedSequence, eventsForLaunchpad, guiButtons)

		// Show the Fixture Status Buttons.
		ShowFixtureStatus(this.TargetSequence, *sequences[this.SelectedSequence], this, eventsForLaunchpad, guiButtons, commandChannels)
	}

	// Are we in function mode ?
	if this.SelectMode[this.SelectedSequence] == FUNCTION {
		if debug {
			fmt.Printf("%d: Handle 3\n", this.SelectedSequence)
		}
		// Turn off function mode. Remove the function pads.
		common.ClearSelectedRowOfButtons(this.SelectedSequence, eventsForLaunchpad, guiButtons)

		// And reveal the sequence on the launchpad keys
		common.RevealSequence(this.SelectedSequence, commandChannels)

		// If the chaser is running, reveal it.
		if this.ScannerChaser && this.SelectedType == "scanner" {
			common.RevealSequence(this.ChaserSequenceNumber, commandChannels)
		}

		// Turn off the function mode flag.
		this.SelectMode[this.SelectedSequence] = NORMAL

		// Turn off the edit sequence colors button.
		if this.EditSequenceColorsMode {
			this.EditSequenceColorsMode = false
			this.Functions[this.TargetSequence][common.Function5_Color].State = false
		}

		// Now forget we pressed twice and start again.
		this.SelectButtonPressed[this.SelectedSequence] = false

		return
	}

}
