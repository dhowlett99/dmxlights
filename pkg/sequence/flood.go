package sequence

import (
	"fmt"

	"github.com/dhowlett99/dmxlights/pkg/common"
)

func startFlood(mySequenceNumber int, sequence *common.Sequence, fixtureStepChannels []chan common.FixtureCommand) {
	if debug {
		fmt.Printf("sequence %d Start flood mode\n", mySequenceNumber)
	}
	// Prepare a message to be sent to the fixtures in the sequence.
	command := common.FixtureCommand{
		Master:         sequence.Master,
		Blackout:       sequence.Blackout,
		Type:           sequence.Type,
		Label:          sequence.Label,
		SequenceNumber: sequence.Number,
		StartFlood:     sequence.StartFlood,
		StrobeSpeed:    sequence.StrobeSpeed,
		Strobe:         sequence.Strobe,
	}

	// Now tell all the fixtures what they need to do.
	sendToAllFixtures(fixtureStepChannels, command)
	sequence.FloodPlayOnce = false
}

func stopFlood(mySequenceNumber int, sequence *common.Sequence, fixtureStepChannels []chan common.FixtureCommand) {
	if debug {
		fmt.Printf("sequence %d Stop flood mode\n", mySequenceNumber)
	}
	// Prepare a message to be sent to the fixtures in the sequence.
	command := common.FixtureCommand{
		Master:         sequence.Master,
		Blackout:       sequence.Blackout,
		Type:           sequence.Type,
		Label:          sequence.Label,
		SequenceNumber: sequence.Number,
		StartFlood:     sequence.StartFlood,
		StopFlood:      sequence.StopFlood,
		StrobeSpeed:    sequence.StrobeSpeed,
		Strobe:         sequence.Strobe,
	}
	// Now tell all the fixtures what they need to do.
	sendToAllFixtures(fixtureStepChannels, command)
	sequence.StartFlood = false
	sequence.StopFlood = false
	sequence.FloodPlayOnce = false

}
