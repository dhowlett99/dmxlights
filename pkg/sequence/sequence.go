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
		Name:         "cans",
		Number:       mySequenceNumber,
		FadeTime:     0 * time.Millisecond,
		MusicTrigger: true,
		Run:          true,
		Patten: common.Patten{
			Name:     "colors",
			Length:   2,
			Size:     2,
			Fixtures: 8,
			Chase:    []int{1, 2, 3, 4, 5, 6, 7, 8},
			Steps:    pattens["color"].Steps,
		},
		CurrentSpeed: 50 * time.Millisecond,
		Colors: []common.Color{
			{
				R: 0,
				G: 0,
				B: 0,
			},
		},
		Shift:    2,
		Blackout: false,
	}

	for {

		// So this is the outer loop where sequence waits for commands and processes them if we're not playing a sequence.
		// i.e the sequence is in STOP mode and this is the way we change the RUN flag to START a sequence again.
		sequence = commands.ListenCommandChannelAndWait(sequence, commandChannel, replyChannel, soundTriggerChannel, soundTriggerControls)

		if sequence.Run {
			for _, step := range pattens[sequence.Patten.Name].Steps {

				for fixture := range step.Fixtures {
					for color := range step.Fixtures[fixture].Colors {
						sequence = commands.ListenCommandChannelAndWait(sequence, commandChannel, replyChannel, soundTriggerChannel, soundTriggerControls)
						R := step.Fixtures[fixture].Colors[color].R
						G := step.Fixtures[fixture].Colors[color].G
						B := step.Fixtures[fixture].Colors[color].B
						// Now trigger the fixture lamp on the launch pad by sending an event.
						e := common.ALight{
							X:          fixture,
							Y:          mySequenceNumber - 1,
							Brightness: 255,
							Red:        R,
							Green:      G,
							Blue:       B,
						}
						eventsForLauchpad <- e

						// Now ask DMX to actually light the real fixture.
						dmx.Fixtures(mySequenceNumber, dmxController, fixture, R, G, B, groups, sequence.Blackout)
						sequence = commands.ListenCommandChannelAndWait(sequence, commandChannel, replyChannel, soundTriggerChannel, soundTriggerControls)
					}
				}
			}

			for index := len(pattens[sequence.Patten.Name].Steps) - 1; index >= 0; index-- {
				step := pattens[sequence.Patten.Name].Steps[index]
				for fixture := range step.Fixtures {
					for color := range step.Fixtures[fixture].Colors {
						sequence = commands.ListenCommandChannelAndWait(sequence, commandChannel, replyChannel, soundTriggerChannel, soundTriggerControls)
						R := step.Fixtures[fixture].Colors[color].R
						G := step.Fixtures[fixture].Colors[color].G
						B := step.Fixtures[fixture].Colors[color].B
						// Now trigger the fixture lamp on the launch pad by sending an event.
						e := common.ALight{
							X:          fixture,
							Y:          mySequenceNumber - 1,
							Brightness: 255,
							Red:        R,
							Green:      G,
							Blue:       B,
						}
						eventsForLauchpad <- e

						// Now ask DMX to actually light the real fixture.
						dmx.Fixtures(mySequenceNumber, dmxController, fixture, R, G, B, groups, sequence.Blackout)
						sequence = commands.ListenCommandChannelAndWait(sequence, commandChannel, replyChannel, soundTriggerChannel, soundTriggerControls)
					}
				}
			}
		}
	}
}
