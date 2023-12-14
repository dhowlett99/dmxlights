package buttons

import (
	"fmt"

	"github.com/dhowlett99/dmxlights/pkg/common"
)

func displayMode(sequenceNumber int, mode int, this *CurrentState, sequences []*common.Sequence, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command) {

	debug := true

	// Tailor the top buttons to the sequence type.
	common.ShowTopButtons(sequences[sequenceNumber].Type, eventsForLaunchpad, guiButtons)

	// Tailor the bottom buttons to the sequence type.
	common.ShowBottomButtons(sequences[sequenceNumber].Type, eventsForLaunchpad, guiButtons)

	// Show this sequence running status in the start/stop button.
	common.ShowRunningStatus(this.Running[sequenceNumber], eventsForLaunchpad, guiButtons)
	common.ShowStrobeButtonStatus(this.Strobe[this.SelectedSequence], eventsForLaunchpad, guiButtons)

	// Update the status bar.
	showStatusBar(this, sequences, guiButtons)

	// Light the sequence selector button.
	SequenceSelect(eventsForLaunchpad, guiButtons, this)

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

		// Select all fixtures.
		this.SelectAllStaticFixtures = true

		// Flash the static buttons,
		this.StaticFlashing[sequenceNumber] = true

		return

	case mode == CHASER_DISPLAY:

		if debug {
			fmt.Printf("%d: DisplayMode: CHASER_DISPLAY\n", sequenceNumber)
		}
		// Hide the selected sequence.
		common.HideSequence(sequenceNumber, commandChannels)

		// Reveal the chaser sequence.
		common.RevealSequence(this.ChaserSequenceNumber, commandChannels)

		if this.StaticFlashing[sequenceNumber] {
			// Unselect all fixtures.
			this.SelectAllStaticFixtures = false
			// Stop the flash of the static buttons,
			this.StaticFlashing[sequenceNumber] = false
		}

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

		// Flash the static buttons,
		flashwStaticButtons(this.ChaserSequenceNumber, true, false, commandChannels)
		this.StaticFlashing[sequenceNumber] = true

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

		// Turn off any flashing static buttons.
		if this.StaticFlashing[sequenceNumber] {
			// Unselect all fixtures.
			this.SelectAllStaticFixtures = false
			// Stop the flash of the static buttons,
			flashwStaticButtons(sequenceNumber, false, true, commandChannels)
			this.StaticFlashing[sequenceNumber] = false
		}

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

		if this.StaticFlashing[sequenceNumber] {
			// Unselect all fixtures.
			this.SelectAllStaticFixtures = false

			// Stop the flash of the static buttons, taking care to select the correct sequence.
			if this.ScannerChaser[sequenceNumber] && this.SelectedType == "scanner" {
				flashwStaticButtons(this.ChaserSequenceNumber, false, true, commandChannels)
			} else {
				flashwStaticButtons(sequenceNumber, false, true, commandChannels)
			}
			this.StaticFlashing[sequenceNumber] = false
		}

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

		if this.StaticFlashing[sequenceNumber] {
			// Unselect all fixtures.
			this.SelectAllStaticFixtures = false

			// Stop the flash of the static buttons, taking care to select the correct sequence.
			if this.ScannerChaser[sequenceNumber] && this.SelectedType == "scanner" {
				flashwStaticButtons(this.ChaserSequenceNumber, false, true, commandChannels)
			} else {
				flashwStaticButtons(sequenceNumber, false, true, commandChannels)
			}
			this.StaticFlashing[sequenceNumber] = false
		}

		// Display the fixture status bar.
		showFixtureStatus(this.TargetSequence, sequences[sequenceNumber].Number, sequences[sequenceNumber].NumberFixtures, this, eventsForLaunchpad, guiButtons, commandChannels)

		return
	}

	if debug {
		fmt.Printf("%d: No Mode Selected\n", sequenceNumber)
	}

}
