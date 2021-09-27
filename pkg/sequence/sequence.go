package sequence

import (
	"fmt"
	"time"

	"github.com/dhowlett99/dmxlights/pkg/commands"
	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
	"github.com/dhowlett99/dmxlights/pkg/launchpad"
	"github.com/oliread/usbdmx/ft232"
	"github.com/rakyll/launchpad/mk2"
)

func CreateSequence(
	sequenceType string,
	mySequenceNumber int,
	pattens map[string]common.Patten,
	channels common.Channels) common.Sequence {

	// Make a map to hold static colors.
	staticColors := make(map[int]common.Color)

	// Set default values.
	sequence := common.Sequence{
		Hide:         false,
		StaticColors: staticColors,
		Name:         sequenceType,
		Number:       mySequenceNumber,
		FadeSpeed:    9,
		FadeTime:     150 * time.Millisecond,
		MusicTrigger: false,
		Run:          true,
		Bounce:       false,
		Steps:        8 * 14, // Eight lamps and 14 steps to fade up and down.
		Patten: common.Patten{
			Name:     sequenceType,
			Length:   2,
			Size:     2,
			Fixtures: 8,
			Chase:    []int{1, 2, 3, 4, 5, 6, 7, 8},
			Steps:    pattens[sequenceType].Steps,
		},
		Speed:        14,
		CurrentSpeed: 25 * time.Millisecond,
		Colors: []common.Color{
			{
				R: 0,
				G: 0,
				B: 0,
			},
		},
		Shift:    2,
		Blackout: false,
		Master:   255,
	}

	// Make functions for each of the sequences.
	for function := 0; function < 8; function++ {
		newFunction := common.Function{
			Name:           fmt.Sprintf("function %d", function),
			SequenceNumber: mySequenceNumber,
			Number:         function,
			State:          false,
		}
		sequence.Functions = append(sequence.Functions, newFunction)
	}
	return sequence
}

func PlayNewSequence(sequence common.Sequence,
	mySequenceNumber int,
	pad *mk2.Launchpad,
	eventsForLauchpad chan common.ALight,
	pattens map[string]common.Patten,
	dmxController ft232.DMXController,
	fixtureConfig *fixture.Fixtures,
	channels common.Channels) {

	positions := map[int][]common.Position{}

	// Create eight channels to control the fixtures.
	fixtureChannel1 := make(chan common.FixtureCommand)
	fixtureChannel2 := make(chan common.FixtureCommand)
	fixtureChannel3 := make(chan common.FixtureCommand)
	fixtureChannel4 := make(chan common.FixtureCommand)
	fixtureChannel5 := make(chan common.FixtureCommand)
	fixtureChannel6 := make(chan common.FixtureCommand)
	fixtureChannel7 := make(chan common.FixtureCommand)
	fixtureChannel8 := make(chan common.FixtureCommand)

	// Make an array to hold all the fixture channels.
	fixtureChannels := []chan common.FixtureCommand{}
	fixtureChannels = append(fixtureChannels, fixtureChannel1)
	fixtureChannels = append(fixtureChannels, fixtureChannel2)
	fixtureChannels = append(fixtureChannels, fixtureChannel3)
	fixtureChannels = append(fixtureChannels, fixtureChannel4)
	fixtureChannels = append(fixtureChannels, fixtureChannel5)
	fixtureChannels = append(fixtureChannels, fixtureChannel6)
	fixtureChannels = append(fixtureChannels, fixtureChannel7)
	fixtureChannels = append(fixtureChannels, fixtureChannel8)

	// Create eight fixture threads for this sequence.
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 0, fixtureChannels, eventsForLauchpad, dmxController, fixtureConfig)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 1, fixtureChannels, eventsForLauchpad, dmxController, fixtureConfig)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 2, fixtureChannels, eventsForLauchpad, dmxController, fixtureConfig)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 3, fixtureChannels, eventsForLauchpad, dmxController, fixtureConfig)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 4, fixtureChannels, eventsForLauchpad, dmxController, fixtureConfig)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 5, fixtureChannels, eventsForLauchpad, dmxController, fixtureConfig)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 6, fixtureChannels, eventsForLauchpad, dmxController, fixtureConfig)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 7, fixtureChannels, eventsForLauchpad, dmxController, fixtureConfig)

	// So this is the outer loop where sequence waits for commands and processes them if we're not playing a sequence.
	// i.e the sequence is in STOP mode and this is the way we change the RUN flag to START a sequence again.
	for {

		sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence.CurrentSpeed*10, sequence, channels)

		// Sequence in Static Mode.
		if sequence.Static {
			for myFixtureNumber, lamp := range sequence.StaticColors {
				if !sequence.Hide {
					launchpad.LightLamp(mySequenceNumber, myFixtureNumber, lamp.R, lamp.G, lamp.B, eventsForLauchpad)
				}
				fixture.MapFixtures(mySequenceNumber, dmxController, myFixtureNumber, lamp.R, lamp.G, lamp.B, 0, 0, 0, 0, fixtureConfig, sequence.Blackout, sequence.Master, sequence.Master)
			}
			// Only play once, we don't want to flood the DMX universe with
			// continual commands.
			sequence.Static = false
			continue
		}

		// Sequence in Normal Running Mode.
		if sequence.Run {

			// Map function keys 0-7 to sequencer functions.
			sequence.Bounce = sequence.Functions[common.Function7_Bounce].State
			sequence.MusicTrigger = sequence.Functions[common.Function8_Music_Trigger].State
			if sequence.MusicTrigger {
				sequence.CurrentSpeed = time.Duration(12 * time.Hour)
			}

			sequence.Functions[common.Function1_Forward_Chase].State = sequence.Run

			if sequence.Patten.Name == "scanner" {
				sequence.Type = "scanner"
			}

			sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence.CurrentSpeed, sequence, channels)
			if !sequence.Run {
				break
			}

			// Calulate positions for fixtures based on patten.
			positions, sequence.Steps = calculatePositions(pattens[sequence.Patten.Name].Steps, sequence.Bounce)

			// Run the sequence through.
			for step := 0; step < sequence.Steps; step++ {
				cmd := common.FixtureCommand{
					Hide:            sequence.Hide,
					Tick:            true,
					Positions:       positions,
					Type:            sequence.Type,
					FadeSpeed:       sequence.FadeSpeed,
					FadeTime:        sequence.FadeTime,
					Size:            sequence.Size,
					Steps:           sequence.Steps,
					CurrentSpeed:    sequence.CurrentSpeed,
					Speed:           sequence.Speed,
					Blackout:        sequence.Blackout,
					CurrentPosition: step,
				}
				fixtureChannel1 <- cmd
				sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, 1*time.Millisecond, sequence, channels)
				if !sequence.Run {
					break
				}
				fixtureChannel2 <- cmd
				sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, 1*time.Millisecond, sequence, channels)
				if !sequence.Run {
					break
				}
				fixtureChannel3 <- cmd
				sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, 1*time.Millisecond, sequence, channels)
				if !sequence.Run {
					break
				}
				fixtureChannel4 <- cmd
				sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, 1*time.Millisecond, sequence, channels)
				if !sequence.Run {
					break
				}
				fixtureChannel5 <- cmd
				sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, 1*time.Millisecond, sequence, channels)
				if !sequence.Run {
					break
				}
				fixtureChannel6 <- cmd
				sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, 1*time.Millisecond, sequence, channels)
				if !sequence.Run {
					break
				}
				fixtureChannel7 <- cmd
				sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, 1*time.Millisecond, sequence, channels)
				if !sequence.Run {
					break
				}
				fixtureChannel8 <- cmd

				sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence.CurrentSpeed, sequence, channels)
				if !sequence.Run {
					break
				}
			}
		}
	}
}

// calculatePositions takes the steps defined in the patten and
// turns them into positions used by the sequencer.
func calculatePositions(steps []common.Step, bounce bool) (map[int][]common.Position, int) {

	position := common.Position{}

	// We have multiple positions for each fixture.
	var counter int
	positionsOut := make(map[int][]common.Position)

	for _, step := range steps {
		for fixtureIndex, fixture := range step.Fixtures {
			for _, color := range fixture.Colors {
				// Preserve the scanner commands.
				position.Gobo = fixture.Gobo
				position.Pan = fixture.Pan
				position.Tilt = fixture.Tilt
				position.Shutter = fixture.Shutter
				if color.R > 0 || color.G > 0 || color.B > 0 {
					position.StartPosition = counter
					position.Fixture = fixtureIndex
					position.Color.R = color.R
					position.Color.G = color.G
					position.Color.B = color.B
					positionsOut[counter] = append(positionsOut[counter], position)
					if fixture.Type != "scanner" {
						counter = counter + 14
					}
				}
			}
		}
		if step.Type == "scanner" {
			counter = counter + 14
		}
	}

	if bounce {
		for index := len(steps) - 1; index >= 0; index-- {
			step := steps[index]
			for fixtureIndex, fixture := range step.Fixtures {
				for _, color := range fixture.Colors {
					// Preserve the scanner commands.
					position.Gobo = fixture.Gobo
					position.Pan = fixture.Pan
					position.Tilt = fixture.Tilt
					position.Shutter = fixture.Shutter
					if color.R > 0 || color.G > 0 || color.B > 0 {
						position.StartPosition = counter
						position.Fixture = fixtureIndex
						position.Color.R = color.R
						position.Color.G = color.G
						position.Color.B = color.B
						positionsOut[counter] = append(positionsOut[counter], position)
						if fixture.Type != "scanner" {
							counter = counter + 14
						}
					}
				}
			}
			if step.Type == "scanner" {
				counter = counter + 14
			}
		}
	}
	return positionsOut, counter
}
