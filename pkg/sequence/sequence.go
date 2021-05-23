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

	//fmt.Printf("Setup default command\n")
	// set default values.
	command := common.Sequence{
		Name:     "cans",
		Number:   mySequenceNumber,
		FadeTime: 0 * time.Millisecond,
		Run:      true,
		Patten: common.Patten{
			Name:     "standard",
			Length:   2,
			Size:     2,
			Fixtures: 8,
			Chase:    []int{1, 2, 3, 4, 5, 6, 7, 8},
			Steps:    Pattens["standard"].Steps,
		},
		CurrentSpeed: 300 * time.Millisecond,
		Colors: []common.Color{
			{
				R: 0,
				G: 0,
				B: 0,
			},
		},
		Speed: 3,
		Shift: 2,
	}

	channels := []chan common.Event{}
	// Create a channel for every fixture.
	//fmt.Printf("Create a channel for every fixture.\n")
	for fixture := 0; fixture < command.Patten.Fixtures; fixture++ {
		channel := make(chan common.Event)
		channels = append(channels, channel)
	}

	// Now start the fixture threads listening.
	//fmt.Printf("Now start the fixture threads listening.")
	for thisFixture, channel := range channels {
		go fixture.FixtureReceiver(channel,
			thisFixture,
			command,
			commandChannel,
			replyChannel,
			mySequenceNumber,
			Pattens,
			eventsForLauchpad)
	}

	cmd := common.Event{}
	// Now trigger the fixture by sending an event.
	for {
		// Listen on the command channel which controls the sequence.
		command = commands.ListenCommandChannelAndWait(command, commandChannel, replyChannel, command.CurrentSpeed, mySequenceNumber)
		if command.Stop {
			fmt.Printf("Seq: Stop is %t\n", command.Stop)
		}
		//fmt.Printf("Seq: Command is %t\n", command.Run)
		if command.Run {

			//fmt.Printf("Seq says start fixture 0\n")
			for index, channel := range channels {
				cmd.Fixture = index
				channel <- cmd
			}
			time.Sleep(1 * time.Second)

		}

		// for command.Run {

		// 	for _, channel := range channels {
		// 		channel <- cmd
		// 	}
		// }
		// // for {
		// // 	for command.Start {
		// // 		fmt.Printf("Seq says start \n")
		// // 		channels[0] <- cmd
		// // 	}
		// // 	command.Start = false
		// // }
	}
}
