package sequence

import (
	"fmt"

	"github.com/dhowlett99/dmxlights/pkg/common"
)

func showAllSwitches(mySequenceNumber int, sequence *common.Sequence, fixtureStepChannels []chan common.FixtureCommand, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight) {
	if debug {
		fmt.Printf("sequence %d Play all switches mode\n", mySequenceNumber)
	}
	// Show initial state of switches
	for switchNumber := 0; switchNumber < len(sequence.Switches); switchNumber++ {
		setSwitchLamp(*sequence, switchNumber, eventsForLaunchpad, guiButtons)
		setSwitchDMX(*sequence, switchNumber, fixtureStepChannels)
	}
	sequence.PlaySwitchOnce = false
}

func showSelectedSwitch(mySequenceNumber int, sequence *common.Sequence, fixtureStepChannels []chan common.FixtureCommand, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight) {
	if debug {
		fmt.Printf("%d: Play single switch number %d\n", mySequenceNumber, sequence.CurrentSwitch)
	}

	// Dim the last lamp.
	if sequence.CurrentSwitch != sequence.LastSwitchSelected || !sequence.FocusSwitch {
		// Clear the last selected switch.
		newSwitch := sequence.Switches[sequence.LastSwitchSelected]
		newSwitch.Selected = false
		sequence.Switches[sequence.LastSwitchSelected] = newSwitch
		setSwitchLamp(*sequence, sequence.LastSwitchSelected, eventsForLaunchpad, guiButtons)
	}

	// Now show the current switch state.
	if sequence.StepSwitch {
		// This is the second press so actually switch and send the DMX command.
		setSwitchLamp(*sequence, sequence.CurrentSwitch, eventsForLaunchpad, guiButtons)
		setSwitchDMX(*sequence, sequence.CurrentSwitch, fixtureStepChannels)
	} else {
		// first time we presses this switch button just move the focus here and use full brightness to indicate we
		// are the selected sequence and selected switch.
		setSwitchLamp(*sequence, sequence.CurrentSwitch, eventsForLaunchpad, guiButtons)
	}

	sequence.LastSwitchSelected = sequence.CurrentSwitch

	sequence.PlaySwitchOnce = false
	sequence.PlaySingleSwitch = false
	sequence.OverrideSpeed = false
}
