package sequence

import (
	"fmt"

	"github.com/dhowlett99/dmxlights/pkg/common"
)

func clearSequence(mySequenceNumber int, sequence *common.Sequence, fixtureStepChannels []chan common.FixtureCommand) {

	if debug {
		fmt.Printf("sequence %d CLEAR\n", mySequenceNumber)
	}
	// Prepare a message to be sent to the fixtures in the sequence.
	command := common.FixtureCommand{
		Master:         sequence.Master,
		Blackout:       sequence.Blackout,
		Type:           sequence.Type,
		Label:          sequence.Label,
		SequenceNumber: sequence.Number,
		Clear:          sequence.Clear,
	}

	// Now tell all the fixtures what they need to do.
	sendToAllFixtures(fixtureStepChannels, command)
	sequence.Clear = false

}
