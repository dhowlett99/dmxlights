package sequence

import (
	"time"

	"github.com/dhowlett99/dmxlights/pkg/commands"
	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/dmx"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
	"github.com/dhowlett99/dmxlights/pkg/sound"
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
	dmxController ft232.DMXController,
	soundTriggerChannel chan common.Command,
	soundTriggerControls *sound.Sound,
	groups *fixture.Groups) {

	// set default values.
	sequence := common.Sequence{
		Name:     "cans",
		Number:   mySequenceNumber,
		FadeTime: 0 * time.Millisecond,
		Run:      true,
		Patten: common.Patten{
			Name:     "colors",
			Length:   2,
			Size:     2,
			Fixtures: 8,
			Chase:    []int{1, 2, 3, 4, 5, 6, 7, 8},
			Steps:    pattens["colors"].Steps,
		},
		CurrentSpeed: 250 * time.Millisecond,
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

	fixtureChannels := []chan common.Event{}
	// Create a channel for every fixture.
	for fixture := 0; fixture < sequence.Patten.Fixtures; fixture++ {
		thisFixtureChannel := make(chan common.Event)
		fixtureChannels = append(fixtureChannels, thisFixtureChannel)
	}

	// Now start the fixture threads listening.
	for thisFixture, channel := range fixtureChannels {
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
		sequence = commands.ListenCommandChannelAndWait(sequence, commandChannel, replyChannel, sequence.CurrentSpeed, mySequenceNumber, soundTriggerChannel, soundTriggerControls)

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
				sequence = commands.ListenCommandChannelAndWait(sequence, commandChannel, replyChannel, sequence.CurrentSpeed, mySequenceNumber, soundTriggerChannel, soundTriggerControls)
				playStep(mySequenceNumber, step, sequence, fixtureChannels, pattens, dmxController, groups)
			}
		}
	}
}

func playStep(mySequenceNumber int,
	step common.Step,
	command common.Sequence,
	fixtureChannels []chan common.Event,
	pattens map[string]common.Patten,
	dmxController ft232.DMXController,
	groups *fixture.Groups) {

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
			fixtureChannels[fixture] <- event

			dmx.Fixtures(mySequenceNumber, dmxController, fixture, R, G, B, groups)
		}
	}
}
