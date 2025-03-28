// Copyright (C) 2022, 2023 , 2024 dhowlett99.
// This is button processor, used by the launchpad and gui interfaces.
// This file controls the display when pressing the select button.
// Decides which menu items / buttons are displayed.
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

// Select modes.
const (
	NORMAL                int = iota // Normal RGB or Scanner Rotation display.
	NORMAL_STATIC                    // Normal RGB in edit all static fixtures.
	FUNCTION                         // Show the RGB or Scanner functions.
	CHASER_DISPLAY                   //  Show the scanner shutter display.
	CHASER_DISPLAY_STATIC            //  Shutter chaser in edit all fixtures mode.
	CHASER_FUNCTION                  // Show the scammer shutter chaser functions.
	STATUS                           // Show the fixture status states.
)

func displayMode(sequenceNumber int, mode int, this *CurrentState, sequences []*common.Sequence, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command) {

	debug := false

	// Show this sequence running status in the start/stop button.
	common.ShowRunningStatus(this.Running[sequenceNumber], eventsForLaunchpad, guiButtons)
	common.ShowStrobeButtonStatus(this.Strobe[this.SelectedSequence], eventsForLaunchpad, guiButtons)

	// Update the status bar.
	showStatusBars(this, sequences, eventsForLaunchpad, guiButtons)

	// Light the sequence selector button.
	lightSelectedButton(eventsForLaunchpad, guiButtons, this)

	switch {

	case mode == NORMAL:

		if debug {
			fmt.Printf("%d: DisplayMode: NORMAL\n", sequenceNumber)
		}

		// Make sure we hide the shutter chaser.
		if this.SequenceType[sequenceNumber] == "scanner" {
			common.HideSequence(this.ChaserSequenceNumber, commandChannels)
		}

		common.RevealSequence(sequenceNumber, commandChannels)

		return

	case mode == NORMAL_STATIC:

		if debug {
			fmt.Printf("%d: DisplayMode: NORMAL STATIC\n", sequenceNumber)
		}

		// Make sure we hide any shutter chaser.
		if this.SelectedType == "scanner" {
			common.HideSequence(this.ChaserSequenceNumber, commandChannels)
		}

		// Reveal the selected sequence.
		common.RevealSequence(sequenceNumber, commandChannels)

		return

	case mode == CHASER_DISPLAY:

		if debug {
			fmt.Printf("%d: DisplayMode: CHASER_DISPLAY \n", sequenceNumber)
		}

		// Hide the selected sequence.
		common.HideSequence(sequenceNumber, commandChannels)

		// Reveal the chaser sequence.
		common.RevealSequence(this.ChaserSequenceNumber, commandChannels)

		return

	case mode == CHASER_DISPLAY_STATIC:

		if debug {
			fmt.Printf("%d: DisplayMode: CHASER_DISPLAY_STATIC\n", sequenceNumber)
		}

		// Hide the selected sequence.
		common.HideSequence(sequenceNumber, commandChannels)

		// Reveal the chaser sequence.
		common.RevealSequence(this.ChaserSequenceNumber, commandChannels)

		// Select all fixtures.
		this.SelectAllStaticFixtures = true

		return

	case mode == FUNCTION:

		if debug {
			fmt.Printf("%d: DisplayMode: FUNCTION  Shutter Chaser is %t\n", sequenceNumber, this.ScannerChaser[sequenceNumber])
		}

		hideAllFunctionKeys(this, sequences, eventsForLaunchpad, guiButtons, commandChannels)

		// If we have a shutter chaser running hide it.
		if this.SequenceType[sequenceNumber] == "scanner" {
			common.HideSequence(this.ChaserSequenceNumber, commandChannels)
		}

		// Hide the sequence.
		common.HideSequence(sequenceNumber, commandChannels)

		// Show the function buttons.
		ShowFunctionButtons(this, eventsForLaunchpad, guiButtons)

		return

	case mode == CHASER_FUNCTION:

		if debug {
			fmt.Printf("%d: DisplayMode: CHASER_FUNCTION\n", sequenceNumber)
		}
		// If we have a shutter chaser running hide it.
		if this.ScannerChaser[sequenceNumber] && this.SelectedType == "scanner" {
			common.HideSequence(this.ChaserSequenceNumber, commandChannels)
		}
		// Hide the normal sequence.
		common.HideSequence(sequenceNumber, commandChannels)

		// Show the chaser function buttons.
		this.TargetSequence = this.ChaserSequenceNumber
		ShowFunctionButtons(this, eventsForLaunchpad, guiButtons)

		return

	case mode == STATUS:

		if debug {
			fmt.Printf("%d: DisplayMode: STATUS\n", sequenceNumber)
		}

		// If we're a scanner sequence and trying to display the status bar we don't want a shutter chaser in view.
		if this.SelectedType == "scanner" {
			common.HideSequence(this.ChaserSequenceNumber, commandChannels)
		}

		// Hide the normal sequence.
		common.HideSequence(sequenceNumber, commandChannels)

		// Display the fixture status bar.
		showFixtureStatus(this.TargetSequence, sequences[sequenceNumber], eventsForLaunchpad, guiButtons, commandChannels)

		return
	}

	if debug {
		fmt.Printf("%d: No Mode Selected\n", sequenceNumber)
	}

}

func printMode(mode int) string {
	if mode == NORMAL {
		return "    NORMAL     "
	}
	if mode == NORMAL_STATIC {
		return "    STATIC     "
	}
	if mode == CHASER_DISPLAY {
		return "    CHASER     "
	}
	if mode == CHASER_DISPLAY_STATIC {
		return " CHASER_STATIC "
	}
	if mode == FUNCTION {
		return "   FUNCTION    "
	}
	if mode == CHASER_FUNCTION {
		return "CHASER_FUNCTION"
	}
	if mode == STATUS {
		return "     STATUS    "
	}
	return "UNKNOWN"
}

func clearAllModes(sequences []*common.Sequence, this *CurrentState) {
	for sequenceNumber := range sequences {
		this.SelectButtonPressed[sequenceNumber] = false
		this.SelectedMode[sequenceNumber] = NORMAL
		this.ShowRGBColorPicker = false
		this.Static[this.DisplaySequence] = false
		this.Static[this.TargetSequence] = false
		this.ShowStaticColorPicker = false
		this.EditGoboSelectionMode = false
		this.EditPatternMode = false
		for function := range this.Functions {
			this.Functions[sequenceNumber][function].State = false
		}
	}
}

// SetTarget - If we're a scanner and we're in shutter chase mode and if we're in either CHASER_DISPLAY or CHASER_FUNCTION mode then
// set the target sequence to the chaser sequence number.
// Else the target is just this sequence number.
// Returns the target sequence number.
func SetTarget(this *CurrentState) {
	if this.SelectedType == "scanner" &&
		this.ScannerChaser[this.SelectedSequence] &&
		(this.SelectedMode[this.SelectedSequence] == CHASER_FUNCTION || this.SelectedMode[this.SelectedSequence] == CHASER_DISPLAY) {
		this.TargetSequence = this.ChaserSequenceNumber
		this.DisplaySequence = this.SelectedSequence
	} else {
		this.TargetSequence = this.SelectedSequence
		this.DisplaySequence = this.SelectedSequence
	}
}
