package sequence

import (
	"fmt"
	"time"

	"github.com/dhowlett99/dmxlights/pkg/commands"
	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
	"github.com/rakyll/launchpad/mk2"
)

func CreateSequence(mySequenceNumber int, pad *mk2.Launchpad, eventsForLauchpad chan common.ALight, commandChannel chan common.Sequence, replyChannel chan common.Sequence, pattens map[string]common.Patten) {

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
		//currentColor := 0

		//totalSteps := len(command.Patten.Steps)

		// lastColor := make(map[int]common.Color)

		// step := Pattens[command.Patten.Name].Steps

		//lastStep := common.Step{}

		if command.Run {
			for _, step := range pattens[command.Patten.Name].Steps {

				command = commands.ListenCommandChannelAndWait(command, commandChannel, replyChannel, command.CurrentSpeed, mySequenceNumber)
				if command.Stop {
					fmt.Printf("Seq: Stop is %t\n", command.Stop)
				}

				//for _, actualStep := range calcSteps(lastStep, step) {
				playStep(step, command, channels, pattens)
				//lastStep = step
				//}
			}
		}
	}
}

func playStep(step common.Step, command common.Sequence, channels []chan common.Event, pattens map[string]common.Patten) {
	if command.Run {

		//step := pattens[command.Patten.Name].Steps
		// Start the color counter.
		currentColor := 0

		for fixture := range step.Fixtures {
			// Listen on the command channel which controls the sequence.
			// command = commands.ListenCommandChannelAndWait(command, commandChannel, replyChannel, command.CurrentSpeed, mySequenceNumber)
			// if command.Stop {
			// 	fmt.Printf("Seq: Stop is %t\n", command.Stop)
			// }
			//tolalColors := len(step[stepCount].Fixtures[fixture].Colors)

			R := step.Fixtures[fixture].Colors[currentColor].R
			G := step.Fixtures[fixture].Colors[currentColor].G
			B := step.Fixtures[fixture].Colors[currentColor].B

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
			channels[fixture] <- event

			// time.Sleep(10 * time.Microsecond)

			// event.Shift = 1

			// event = common.Event{
			// 	Fadedown: true,
			// 	Color: common.Color{
			// 		R: R,
			// 		G: G,
			// 		B: B,
			// 	},
			// 	FadeTime: command.CurrentSpeed,
			// 	Shift:    event.Shift,
			// }
			// channels[fixture] <- event
		}
	}
}

func calcSteps(lastStep common.Step, nextStep common.Step) []common.Step {

	finalSteps := []common.Step{}

	for newStep := 0; newStep < 4; newStep++ {

		step := common.Step{}

		for fixture := 0; fixture < len(lastStep.Fixtures); fixture++ {
			//for lastStep.Fixtures[fixture].Colors[0].G < nextStep.Fixtures[fixture].Colors[0].G {
			out := calcNextStepFixtureValue(lastStep.Fixtures[fixture].Colors[0], nextStep.Fixtures[fixture].Colors[0])
			fmt.Printf("out is %+v\n", out)
			newFixture := common.Fixture{}
			newFixture.Colors = append(newFixture.Colors, out)
			step.Fixtures = append(step.Fixtures, newFixture)
			lastStep.Fixtures[fixture].Colors[0].G = out.G
			//}
		}

		finalSteps = append(finalSteps, step)
	}
	return finalSteps
}

func calcNextStepFixtureValue(in common.Color, wanted common.Color) (out common.Color) {

	// Fade up.
	if in.G < wanted.G {
		switch in.G {
		case 0:
			out.G = 1

		case 1:
			out.G = 2

		case 2:
			out.G = 3
		}
	}
	// Fade down.
	if in.G > wanted.G {
		switch in.G {
		case 3:
			out.G = 2

		case 2:
			out.G = 1

		case 1:
			out.G = 0
		}
	}
	fmt.Printf("in %d  wanted %d   out %d \n", in.G, wanted.G, out.G)
	return out
}
