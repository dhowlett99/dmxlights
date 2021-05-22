package sequence

import (
	"fmt"
	"time"

	"github.com/dhowlett99/dmxlights/pkg/commands"
	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
	"github.com/rakyll/launchpad/mk2"
)

func CreateSequence(mySequenceNumber int, pad *mk2.Launchpad, eventsForLauchpad chan common.ALight, commandChannel chan common.Sequence, replyChannel chan common.Sequence, Pattens map[string]common.Patten) {

	fmt.Printf("Setup default command\n")
	// set default values.
	command := common.Sequence{
		Name:     "cans",
		Number:   mySequenceNumber,
		FadeTime: 0 * time.Millisecond,
		Run:      false,
		Patten: common.Patten{
			Name:     "standard",
			Length:   2,
			Size:     2,
			Fixtures: 8,
			Chase:    []int{1, 2, 3, 4, 5, 6, 7, 8},
			Steps:    Pattens["standard"].Steps,
		},
		CurrentSpeed: 100 * time.Millisecond,
		Colors: []common.Color{
			{
				R: 0,
				G: 0,
				B: 0,
			},
		},
		Speed: 3,
	}

	// Create a channel for every fixture.
	fmt.Printf("Create a channel for every fixture.\n")
	fixtureChannels := []chan common.Event{}
	for fixture := 0; fixture < command.Patten.Fixtures; fixture++ {
		channel := make(chan common.Event)
		fixtureChannels = append(fixtureChannels, channel)
	}

	// Now start the fixture threads listening.
	fmt.Printf("Now start the fixture threads listening.")
	for thisFixture, channel := range fixtureChannels {

		//fmt.Printf("Start a thread %d fixture.\n", fixture)
		//time.Sleep(1 * time.Second)

		go fixture.FixtureReceiver(channel, thisFixture, command, commandChannel, replyChannel, mySequenceNumber, Pattens, eventsForLauchpad)

	}

	event := common.Event{}
	// Now start the fixture threads by sending an event.
	for {
		//fmt.Printf("Step to fixture loop %d fixtureChannels = %+v\n", event.Fixture, fixtureChannels)
		command = commands.ListenCommandChannelAndWait(command, commandChannel, replyChannel, command.CurrentSpeed, mySequenceNumber)
		for index, channel := range fixtureChannels {
			event.Fixture = index
			event.Run = command.Run
			channel <- event
		}
	}
}
