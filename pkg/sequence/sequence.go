package sequence

import (
	"time"

	"github.com/dhowlett99/dmxlights/pkg/commands"
	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/dmx"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
	"github.com/oliread/usbdmx/ft232"
	"github.com/rakyll/launchpad/mk2"
)

func CreateSequence(
	sequenceType string,
	mySequenceNumber int,
	pattens map[string]common.Patten,
	channels common.Channels) common.Sequence {

	// set default values.
	sequence := common.Sequence{

		Name:         sequenceType,
		Number:       mySequenceNumber,
		FadeTime:     0 * time.Millisecond,
		MusicTrigger: false,
		Run:          false,
		Patten: common.Patten{
			Name:     sequenceType,
			Length:   2,
			Size:     2,
			Fixtures: 8,
			Chase:    []int{1, 2, 3, 4, 5, 6, 7, 8},
			Steps:    pattens[sequenceType].Steps,
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
	return sequence
}

func PlaySequence(sequence common.Sequence,
	mySequenceNumber int,
	pad *mk2.Launchpad,
	eventsForLauchpad chan common.ALight,
	pattens map[string]common.Patten,
	dmxController ft232.DMXController,
	fixtures *fixture.Fixtures,
	channels common.Channels) common.Sequence {

	for {

		// So this is the outer loop where sequence waits for commands and processes them if we're not playing a sequence.
		// i.e the sequence is in STOP mode and this is the way we change the RUN flag to START a sequence again.
		sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence, channels)

		if sequence.Run {

			for _, step := range pattens[sequence.Patten.Name].Steps {

				for fixture := range step.Fixtures {
					for color := range step.Fixtures[fixture].Colors {
						sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence, channels)
						if !sequence.Run {
							continue
						}
						R := step.Fixtures[fixture].Colors[color].R
						G := step.Fixtures[fixture].Colors[color].G
						B := step.Fixtures[fixture].Colors[color].B
						Pan := step.Fixtures[fixture].Pan
						Tilt := step.Fixtures[fixture].Tilt
						Shutter := step.Fixtures[fixture].Shutter
						Gobo := step.Fixtures[fixture].Gobo
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
						dmx.Fixtures(mySequenceNumber, dmxController, fixture, R, G, B, Pan, Tilt, Shutter, Gobo, fixtures, sequence.Blackout)
						sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence, channels)
						if !sequence.Run {
							continue
						}
					}
				}
			}

			for index := len(pattens[sequence.Patten.Name].Steps) - 1; index >= 0; index-- {
				step := pattens[sequence.Patten.Name].Steps[index]
				for fixture := range step.Fixtures {
					for color := range step.Fixtures[fixture].Colors {
						sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence, channels)

						if !sequence.Run {
							continue
						}
						R := step.Fixtures[fixture].Colors[color].R
						G := step.Fixtures[fixture].Colors[color].G
						B := step.Fixtures[fixture].Colors[color].B
						Pan := step.Fixtures[fixture].Pan
						Tilt := step.Fixtures[fixture].Tilt
						Shutter := step.Fixtures[fixture].Shutter
						Gobo := step.Fixtures[fixture].Tilt
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
						dmx.Fixtures(mySequenceNumber, dmxController, fixture, R, G, B, Pan, Tilt, Shutter, Gobo, fixtures, sequence.Blackout)
						sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence, channels)
						if !sequence.Run {
							continue
						}
					}
				}
			}
		}
	}
}
