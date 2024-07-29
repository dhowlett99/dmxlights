package buttons

import (
	"fmt"

	"github.com/dhowlett99/dmxlights/pkg/common"
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
		showFixtureStatus(this.TargetSequence, sequences[sequenceNumber].Number, sequences[sequenceNumber].NumberFixtures, this, eventsForLaunchpad, guiButtons, commandChannels)

		return
	}

	if debug {
		fmt.Printf("%d: No Mode Selected\n", sequenceNumber)
	}

}
