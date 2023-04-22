package buttons

import (
	"fmt"

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/sequence"
)

// HandleSelect - Runs when you press a select button to select a sequence.
func HandleSelect(sequences []*common.Sequence, this *CurrentState, eventsForLaunchpad chan common.ALight,
	commandChannels []chan common.Command, guiButtons chan common.ALight) {

	debug := true

	if debug {
		fmt.Printf("HANDLE: this.Type = %s \n", this.SelectedType)
		for functionNumber := 0; functionNumber < 8; functionNumber++ {
			state := this.Functions[this.SelectedSequence][functionNumber].State
			fmt.Printf("HANDLE: function %d state %t\n", functionNumber, state)
		}
		fmt.Printf("HANDLE: this.ChaserRunning %t \n", this.ScannerChaser)
		fmt.Printf("================== WHAT MODE =================\n")
		fmt.Printf("HANDLE: this.SelectButtonPressed[%d] = %t \n", this.SelectedSequence, this.SelectButtonPressed[this.SelectedSequence])
		if this.SelectMode[this.SelectedSequence] == NORMAL {
			fmt.Printf("HANDLE: this.SelectMode[%d] = NORMAL \n", this.SelectedSequence)
		}
		if this.SelectMode[this.SelectedSequence] == FUNCTION {
			fmt.Printf("HANDLE: this.SelectMode[%d] = FUNCTION \n", this.SelectedSequence)
		}
		if this.SelectMode[this.SelectedSequence] == STATUS {
			fmt.Printf("HANDLE: this.SelectMode[%d] = STATUS \n", this.SelectedSequence)
		}

		fmt.Printf("HANDLE: this.EditSequenceColorsMode[%d] = %t \n", this.SelectedSequence, this.EditSequenceColorsMode[this.SelectedSequence])
		fmt.Printf("HANDLE: this.EditStaticColorsMode[%d] = %t \n", this.SelectedSequence, this.EditStaticColorsMode[this.SelectedSequence])
		fmt.Printf("HANDLE: this.EditGoboSelectionMode[%d] = %t \n", this.SelectedSequence, this.EditGoboSelectionMode[this.SelectedSequence])
		fmt.Printf("HANDLE: this.EditPatternMode[%d] = %t \n", this.SelectedSequence, this.EditPatternMode[this.SelectedSequence])
		fmt.Printf("===============================================\n")
	}

	//        +-------------------+
	//        |       NORMAL      |
	//        +-------------------+
	//                 |
	//                 V
	//        +-------------------+
	//        |     FUNCTION      |
	//        +-------------------+
	//            |            | If Chaser Enabled
	//            V            V
	//            |       +-------------------+
	//            |       |  CHASER FUNCTIONS |
	//            |       +-------------------+
	//            |				|
	//            V				V
	//         +-------------------+
	//         |       STATUS      |
	//         +-------------------+
	//                  |
	//                  V
	//         +-------------------+
	//         |       NORMAL      |
	//         +-------------------+

	// Update the status bar
	if this.Strobe[this.SelectedSequence] {
		common.UpdateStatusBar(fmt.Sprintf("Strobe %02d", this.StrobeSpeed[this.SelectedSequence]), "speed", false, guiButtons)
	} else {
		// Update status bar.
		if this.Functions[this.SelectedSequence][common.Function8_Music_Trigger].State {
			common.UpdateStatusBar("  MUSIC  ", "speed", false, guiButtons)
		} else {
			common.UpdateStatusBar(fmt.Sprintf("Speed %02d", this.Speed[this.SelectedSequence]), "speed", false, guiButtons)
		}
	}

	sensitivity := common.FindSensitivity(this.SoundGain)
	common.UpdateStatusBar(fmt.Sprintf("Sensitivity %02d", sensitivity), "sensitivity", false, guiButtons)
	common.UpdateStatusBar(fmt.Sprintf("Master %02d", this.MasterBrightness), "master", false, guiButtons)

	if sequences[this.SelectedSequence].Type == "rgb" {
		common.UpdateStatusBar(fmt.Sprintf("Shift %02d", this.RGBShift[this.SelectedSequence]), "shift", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Size %02d", this.RGBSize[this.SelectedSequence]), "size", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Fade %02d", this.RGBFade[this.SelectedSequence]), "fade", false, guiButtons)
		common.UpdateStatusBar("       ", "tilt", false, guiButtons)

		common.UpdateStatusBar(fmt.Sprintf("Red %02d", this.StaticButtons[this.SelectedSequence].Color.R), "red", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Green %02d", this.StaticButtons[this.SelectedSequence].Color.G), "green", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Blue %02d", this.StaticButtons[this.SelectedSequence].Color.B), "blue", false, guiButtons)
	}
	if sequences[this.SelectedSequence].Type == "scanner" {
		label := getScannerShiftLabel(this.ScannerShift[this.SelectedSequence])
		common.UpdateStatusBar(fmt.Sprintf("Shift %s", label), "shift", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Size %02d", this.ScannerSize[this.SelectedSequence]), "size", false, guiButtons)
		label = getScannerCoordinatesLabel(this.ScannerCoordinates[this.SelectedSequence])
		common.UpdateStatusBar(fmt.Sprintf("Coord %s", label), "fade", false, guiButtons)

		// Hide the color editing buttons.
		common.UpdateStatusBar(fmt.Sprintf("Tilt %02d", this.OffsetTilt), "tilt", false, guiButtons)
		common.UpdateStatusBar("        ", "red", false, guiButtons)
		common.UpdateStatusBar("        ", "green", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Pan %02d", this.OffsetPan), "pan", false, guiButtons)
	}

	// Light the top buttons.
	common.ShowTopButtons(sequences[this.SelectedSequence].Type, eventsForLaunchpad, guiButtons)

	// Light the sequence selector button.
	sequence.SequenceSelect(eventsForLaunchpad, guiButtons, this.SelectedSequence)

	// Light the strobe button.
	common.ShowStrobeButtonStatus(this.Strobe[this.SelectedSequence], eventsForLaunchpad, guiButtons)

	// Light the start stop button.
	common.ShowRunningStatus(this.SelectedSequence, this.Running, eventsForLaunchpad, guiButtons)

	// 1st Press Select Sequence - This the first time we have pressed the select button.
	// Simply select the selected sequence.
	// But remember we have pressed this select button once.
	if this.SelectMode[this.SelectedSequence] == NORMAL &&
		!this.SelectButtonPressed[this.SelectedSequence] &&
		!this.EditStaticColorsMode[this.SelectedSequence] {

		if debug {
			fmt.Printf("%d: Show Sequence - Handle 2\n", this.SelectedSequence)
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
			this.EditPatternMode[this.SelectedSequence] = false

			// Clear buttons and remove any labels.
			common.ClearSelectedRowOfButtons(this.SelectedSequence, eventsForLaunchpad, guiButtons)
		}

		if this.SelectMode[this.SelectedSequence] == NORMAL &&
			this.Functions[this.SelectedSequence][common.Function5_Color].State && this.EditSequenceColorsMode[this.SelectedSequence] {
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
			this.EditStaticColorsMode[this.SelectedSequence] && // AND static colol mode, the case when we leave static colors edit mode.
			this.SelectButtonPressed[this.SelectedSequence] {

		if debug {
			fmt.Printf("%d: 2nd Press Function Bar Mode - Handle 4\n", this.SelectedSequence)
		}

		// Set function mode.
		this.SelectMode[this.SelectedSequence] = FUNCTION

		// If static, show static colors.
		if this.EditStaticColorsMode[this.SelectedSequence] && sequences[this.SelectedSequence].Type != "scanner" {
			if debug {
				fmt.Printf("Show Static Color Selection Buttons\n")
			}
			common.SetMode(this.SelectedSequence, commandChannels, "Static")
			this.EditStaticColorsMode[this.SelectedSequence] = false
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
		ShowFunctionButtons(this, this.SelectedSequence, eventsForLaunchpad, guiButtons)

		// Now forget we pressed twice and start again.
		this.SelectButtonPressed[this.SelectedSequence] = false

		return
	}

	// 3rd Press Status Mode - we display the fixture status enable/invert/disable buttons.
	if this.SelectMode[this.SelectedSequence] == FUNCTION &&
		!this.SelectButtonPressed[this.SelectedSequence] &&
		!this.EditSequenceColorsMode[this.SelectedSequence] &&
		!this.EditStaticColorsMode[this.SelectedSequence] {

		if debug {
			fmt.Printf("%d: Handle 5 Status Mode, Function Bar off, status buttons on\n", this.SelectedSequence)
		}

		// Turn on status mode
		this.SelectMode[this.SelectedSequence] = STATUS

		// If the chase is running, hide it.
		if this.ScannerChaser && this.SelectedType == "scanner" {
			common.HideSequence(this.ChaserSequenceNumber, commandChannels)
		}

		// Remove the function pads.
		common.ClearSelectedRowOfButtons(this.SelectedSequence, eventsForLaunchpad, guiButtons)

		ShowFixtureStatus(this.SelectedSequence, *sequences[this.SelectedSequence], this, eventsForLaunchpad, guiButtons, commandChannels)

		return
	}

	// 4th Press Normal Mode - we head back to normal mode.
	if this.SelectMode[this.SelectedSequence] == STATUS &&
		!this.SelectButtonPressed[this.SelectedSequence] &&
		!this.EditSequenceColorsMode[this.SelectedSequence] &&
		!this.EditStaticColorsMode[this.SelectedSequence] {

		if debug {
			fmt.Printf("%d: Handle 1 Normal Mode, Function Bar off\n", this.SelectedSequence)
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
			this.EditPatternMode[this.SelectedSequence] = true
			common.HideSequence(this.SelectedSequence, commandChannels)
			ShowPatternSelectionButtons(this, sequences[this.SelectedSequence].Master, *sequences[this.SelectedSequence], eventsForLaunchpad, guiButtons)
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
			ShowRGBColorSelectionButtons(this.MasterBrightness, this.SelectedSequence, *sequences[this.SelectedSequence], eventsForLaunchpad, guiButtons)
			return
		}

		// We're in RGB Static Color Mode.
		if this.Functions[this.SelectedSequence][common.Function6_Static_Gobo].State && sequences[this.SelectedSequence].Type == "rgb" {
			if debug {
				fmt.Printf("Show RGB Static Colors\n")
			}
			this.EditStaticColorsMode[this.SelectedSequence] = true

			// Tell the sequence about the new color and where we are in the
			// color cycle.
			cmd := common.Command{
				Action: common.UpdateStaticColor,
				Args: []common.Arg{
					{Name: "Static", Value: true},
					{Name: "StaticLamp", Value: this.LastStaticColorButtonX},
					{Name: "StaticLampFlash", Value: true},
					{Name: "SelectedColor", Value: sequences[this.SelectedSequence].StaticColors[this.LastStaticColorButtonX].SelectedColor},
					{Name: "StaticColor", Value: sequences[this.SelectedSequence].StaticColors[this.LastStaticColorButtonX].Color},
				},
			}
			common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

			return
		}

		// We're in Scanner Gobo Selection Mode.
		if this.Functions[this.SelectedSequence][common.Function6_Static_Gobo].State && sequences[this.SelectedSequence].Type == "scanner" {
			this.Functions[this.SelectedSequence][common.Function6_Static_Gobo].State = false
			this.EditGoboSelectionMode[this.SelectedSequence] = false
		}

		// Allow us to exit the pattern select mode without setting a pattern.
		if this.EditPatternMode[this.SelectedSequence] {
			this.EditPatternMode[this.SelectedSequence] = false
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
		if this.EditSequenceColorsMode[this.SelectedSequence] {
			this.EditSequenceColorsMode[this.SelectedSequence] = false
			this.Functions[this.SelectedSequence][common.Function5_Color].State = false
		}

		// Now forget we pressed twice and start again.
		this.SelectButtonPressed[this.SelectedSequence] = false

		return
	}

}
