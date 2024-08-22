// Copyright (C) 2022, 2023 , 2024 dhowlett99.
// This is button processor, used by the launchpad and gui interfaces.
// This file processes the select buttons and handles their actions.
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

// HandleSelect - Runs when you press a sequence select button to select a sequence.
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

		this.Functions[this.DisplaySequence][common.Function5_Color].State = false
		this.Functions[this.ChaserSequenceNumber][common.Function5_Color].State = false

		removeColorPicker(this, sequences, eventsForLaunchpad, guiButtons, commandChannels)

		// If the Selected Color has come back as empty this means we didn't select any colors.
		// So restore the colors that were already there.
		if debug {
			fmt.Printf("sequences[%d].SequenceColors %+v\n", this.TargetSequence, sequences[this.TargetSequence].SequenceColors)
		}
		if len(sequences[this.TargetSequence].SequenceColors) == 0 {
			if debug {
				fmt.Printf("Restore Sequence Colors\n")
			}
			sequences[this.TargetSequence].SequenceColors = this.SavedSequenceColors[this.TargetSequence]
			if debug {
				fmt.Printf("Now set to sequences[%d].SequenceColors %+v\n", this.TargetSequence, sequences[this.TargetSequence].SequenceColors)
			}
			// Tell the sequence that we have restored the colors.
			cmd := common.Command{
				Action: common.UpdateSequenceColors,
				Args: []common.Arg{
					{Name: "Colors", Value: sequences[this.TargetSequence].SequenceColors},
				},
			}
			common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)
		} else {
			if debug {
				fmt.Printf("%d: handle(): Set Sequence Colors %+v\n", sequences[this.TargetSequence].Number, sequences[this.TargetSequence].SequenceColors)
			}
			// Tell the sequence the colors we have selected.
			cmd := common.Command{
				Action: common.UpdateSequenceColors,
				Args: []common.Arg{
					{Name: "Colors", Value: sequences[this.TargetSequence].SequenceColors},
				},
			}
			common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)
		}

	}

	// Decide if we're on the first press of the select button.
	SelectSequence(this)

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

	// Show the mode in the status bar.
	common.UpdateStatusBar(printMode(this.SelectedMode[this.SelectedSequence]), "displaymode", false, guiButtons)

}

func deFocusAllSwitches(this *CurrentState, sequences []*common.Sequence, commandChannels []chan common.Command) {

	for switchNumber := range sequences[this.SwitchSequenceNumber].Switches {
		this.LastSelectedSwitch = switchNumber
		deFocusSingleSwitch(this, sequences, commandChannels)
	}
}

// Just send a message to defocus the last selected switch button.
func deFocusSingleSwitch(this *CurrentState, sequences []*common.Sequence, commandChannels []chan common.Command) {

	if this.LastSelectedSwitch != common.NOT_SELECTED {

		if debug {
			fmt.Printf("%d: deFocusSwitch single switch number %d\n", this.SwitchSequenceNumber, this.LastSelectedSwitch)
		}

		cmd := common.Command{
			Action: common.UpdateSwitch,
			Args: []common.Arg{
				{Name: "SwitchNumber", Value: this.LastSelectedSwitch},
				{Name: "SwitchPosition", Value: this.SwitchPosition[this.LastSelectedSwitch]},
				{Name: "Step", Value: false},  // Don't step the switch state.
				{Name: "Focus", Value: false}, // Focus the switch lamp.
			},
		}
		// Send a message to the switch sequence.
		common.SendCommandToAllSequenceOfType(sequences, cmd, commandChannels, "switch")
	}
}

func SelectSequence(this *CurrentState) {

	// Decide if we're on the first press of the select button.
	if this.SelectedType != "switch" && this.SelectButtonPressed[this.SelectedSequence] {
		// Calculate the next mode.
		this.SelectedMode[this.SelectedSequence] = getNextMenuItem(this.SelectedMode[this.SelectedSequence], this.ScannerChaser[this.SelectedSequence], getStatic(this))
	}
	if !this.SelectButtonPressed[this.SelectedSequence] {
		this.SelectButtonPressed[0] = false
		this.SelectButtonPressed[1] = false
		this.SelectButtonPressed[2] = false
		this.SelectButtonPressed[3] = false
		this.SelectButtonPressed[4] = false // Switch Sequence.
		this.SelectButtonPressed[this.SelectedSequence] = true
	}
}

func getStatic(this *CurrentState) bool {

	if debug {
		fmt.Printf("Static SelectedSequence %t\n", this.Static[this.SelectedSequence])
		fmt.Printf("Static ChaserSequenceNumber %t\n\n", this.Static[this.ChaserSequenceNumber])
	}

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
