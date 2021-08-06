package sequence

import (
	"time"

	"github.com/dhowlett99/dmxlights/pkg/commands"
	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
	"github.com/oliread/usbdmx/ft232"
	"github.com/rakyll/launchpad/mk2"
)

func CreateSequence(
	mySequenceNumber int,
	pad *mk2.Launchpad,
	eventsForLauchpad chan common.ALight,
	commandChannel chan common.Command,
	replyChannel chan common.Command,
	pattens map[string]common.Patten,
	dmxController ft232.DMXController) {

	// set default values.
	sequence := common.Sequence{
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
	for fixture := 0; fixture < sequence.Patten.Fixtures; fixture++ {
		channel := make(chan common.Event)
		channels = append(channels, channel)
	}

	// Now start the fixture threads listening.
	for thisFixture, channel := range channels {
		go fixture.FixtureReceiver(
			channel,
			thisFixture,
			sequence,
			commandChannel,
			mySequenceNumber,
			eventsForLauchpad)
	}

	for {

		// So this is the outer loop where sequence waits for commands and processes them if we're not playing a sequence.
		// i.e the sequence is in STOP mode and this is the way we change the RUN flag to START a sequence again.
		sequence = commands.ListenCommandChannelAndWait(sequence, commandChannel, replyChannel, sequence.CurrentSpeed, mySequenceNumber)

		// Start the color counter.
		// currentColor := 0

		// totalSteps := len(command.Patten.Steps)

		// lastColor := make(map[int]common.Color)

		if sequence.Run {
			for _, step := range pattens[sequence.Patten.Name].Steps {
				// for fixture := range step.Fixtures {
				// 	fmt.Printf("Step is %v\n", fixture)
				// }
				// This is the inner loop, when we are playing a sequence, we listen for commands here that affect the way the
				// Sequence is performed, and also the way we STOP a sequence.
				sequence = commands.ListenCommandChannelAndWait(sequence, commandChannel, replyChannel, sequence.CurrentSpeed, mySequenceNumber)
				playStep(step, sequence, channels, pattens, dmxController)
			}
		}
	}
}

func playStep(step common.Step, command common.Sequence, channels []chan common.Event, pattens map[string]common.Patten, dmxController ft232.DMXController) {
	if command.Run {
		// Start the color counter.
		currentColor := 0

		for fixture := range step.Fixtures {

			R := step.Fixtures[fixture].Colors[currentColor].R
			G := step.Fixtures[fixture].Colors[currentColor].G
			B := step.Fixtures[fixture].Colors[currentColor].B

			// Now trigger the fixture lamp on the launch pad by sending an event.
			event := common.Event{
				Color: common.Color{
					R: R,
					G: G,
					B: B,
				},
			}
			channels[fixture] <- event

			R = convertToDMXValues(R)
			G = convertToDMXValues(G)
			B = convertToDMXValues(B)

			// Now ask DMX to actually light the real fixture.
			if fixture == 0 {
				dmxController.SetChannel(1, byte(R))
				dmxController.SetChannel(2, byte(G))
				dmxController.SetChannel(3, byte(B))
			}
			if fixture == 1 {
				dmxController.SetChannel(4, byte(R))
				dmxController.SetChannel(5, byte(G))
				dmxController.SetChannel(6, byte(B))
			}
			if fixture == 2 {
				dmxController.SetChannel(7, byte(R))
				dmxController.SetChannel(8, byte(G))
				dmxController.SetChannel(9, byte(B))
			}
			if fixture == 3 {
				dmxController.SetChannel(10, byte(R))
				dmxController.SetChannel(11, byte(G))
				dmxController.SetChannel(12, byte(B))
			}
			dmxController.SetChannel(13, 255)
			if fixture == 4 {
				dmxController.SetChannel(14, byte(R))
				dmxController.SetChannel(15, byte(G))
				dmxController.SetChannel(16, byte(B))
			}
			if fixture == 5 {
				dmxController.SetChannel(17, byte(R))
				dmxController.SetChannel(18, byte(G))
				dmxController.SetChannel(19, byte(B))
			}
			if fixture == 6 {
				dmxController.SetChannel(20, byte(R))
				dmxController.SetChannel(21, byte(G))
				dmxController.SetChannel(22, byte(B))
			}
			if fixture == 7 {
				dmxController.SetChannel(23, byte(R))
				dmxController.SetChannel(24, byte(G))
				dmxController.SetChannel(25, byte(B))
			}
			dmxController.SetChannel(26, 255)

		}
	}
}

func convertToDMXValues(input int) (output int) {

	if input == 0 {
		output = 0
	}
	if input == 1 {
		output = 64
	}
	if input == 2 {
		output = 128
	}
	if input == 3 {
		output = 168
	}
	if input == 4 {
		output = 255
	}

	return output
}
