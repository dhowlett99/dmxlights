package sequence

import (
	"time"

	"github.com/dhowlett99/dmxlights/pkg/commands"
	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
	"github.com/rakyll/launchpad/mk2"
)

func CreateSequence(mySequenceNumber int, pad *mk2.Launchpad, eventsForLauchpad chan common.ALight, commandChannel chan common.Sequence, replyChannel chan common.Sequence, pattens map[string]common.Patten) {

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
			Steps:    pattens["standard"].Steps,
		},
		CurrentSpeed: 50 * time.Millisecond,
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
	for fixture := 0; fixture < command.Patten.Fixtures; fixture++ {
		channel := make(chan common.Event)
		channels = append(channels, channel)
	}

	// Now start the fixture threads listening.
	for thisFixture, channel := range channels {
		go fixture.FixtureReceiver(channel,
			thisFixture,
			command,
			commandChannel,
			mySequenceNumber,
			eventsForLauchpad)
	}

	for {

		// So this is the outer loop where sequence waits for commands and processes them if we're not playing a sequence.
		// i.e the sequence is in STOP mode and this is the way we change the RUN flag to START a sequence again.
		command = commands.ListenCommandChannelAndWait(command, commandChannel, replyChannel, command.CurrentSpeed, mySequenceNumber)

		// Start the color counter.
		// currentColor := 0

		// totalSteps := len(command.Patten.Steps)

		// lastColor := make(map[int]common.Color)

		if command.Run {
			for _, step := range pattens[command.Patten.Name].Steps {
				// This is the inner loop, when we are playing a sequence, we listen for commands here that affect the way the
				// Sequence is performed, and also the way we STOP a sequence.
				command = commands.ListenCommandChannelAndWait(command, commandChannel, replyChannel, command.CurrentSpeed, mySequenceNumber)
				playStep(step, command, channels, pattens)
			}
		}
	}
}

func playStep(step common.Step, command common.Sequence, channels []chan common.Event, pattens map[string]common.Patten) {
	if command.Run {
		// Start the color counter.
		currentColor := 0

		for fixture := range step.Fixtures {

			R := step.Fixtures[fixture].Colors[currentColor].R
			G := step.Fixtures[fixture].Colors[currentColor].G
			B := step.Fixtures[fixture].Colors[currentColor].B

			// Now trigger the fixture by sending an event.
			event := common.Event{
				Color: common.Color{
					R: R,
					G: G,
					B: B,
				},
			}
			channels[fixture] <- event
		}
	}
}
