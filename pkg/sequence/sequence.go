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
		CurrentSpeed: 100 * time.Millisecond,
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
			mySequenceNumber,
			eventsForLauchpad)
	}

	for {

		// Start the color counter.
		currentColor := 0

		//totalSteps := len(command.Patten.Steps)

		lastColor := make(map[int]common.Color)

		step := Pattens[command.Patten.Name].Steps

		if command.Run {
			for stepCount := 0; stepCount < len(Pattens[command.Patten.Name].Steps); stepCount++ {

				for fixture := range step[stepCount].Fixtures {
					// Listen on the command channel which controls the sequence.
					command = commands.ListenCommandChannelAndWait(command, commandChannel, replyChannel, command.CurrentSpeed, mySequenceNumber)
					if command.Stop {
						fmt.Printf("Seq: Stop is %t\n", command.Stop)
					}
					//tolalColors := len(step[stepCount].Fixtures[fixture].Colors)

					R := step[stepCount].Fixtures[fixture].Colors[currentColor].R
					G := step[stepCount].Fixtures[fixture].Colors[currentColor].G
					B := step[stepCount].Fixtures[fixture].Colors[currentColor].B

					//fmt.Printf("Step is %d Fixture is %d Green is %d\n", stepCount, fixture, G)

					// Now trigger the fixture by sending an event.
					event := common.Event{
						Fadeup: true,
						Color: common.Color{
							R: R,
							G: G,
							B: B,
						},
						FadeTime: command.CurrentSpeed / 4,
					}
					lastColor[fixture] = event.Color
					channels[fixture] <- event

					time.Sleep(100 * time.Microsecond)

					event.Shift = 1

					event = common.Event{
						Fadedown:  true,
						LastColor: lastColor[fixture],
						Color: common.Color{
							R: R,
							G: G,
							B: B,
						},
						FadeTime: command.CurrentSpeed,
						Shift:    event.Shift,
					}
					lastColor[fixture] = event.Color
					channels[fixture] <- event

					//time.Sleep(100 * time.Millisecond)
				}

				// if currentColor <= tolalColors {
				// 	currentColor++
				// }

				// if currentColor == tolalColors {
				// 	stepCount++
				// 	currentColor = 0
				// }

				// if stepCount >= totalSteps {
				// 	stepCount = 0
				// 	currentColor = 0
				// }
			}
		}
	}
}
