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
	sequenceType string,
	mySequenceNumber int,
	pattens map[string]common.Patten,
	channels common.Channels) common.Sequence {

	// set default values.
	sequence := common.Sequence{

		Name:         sequenceType,
		Number:       mySequenceNumber,
		FadeSpeed:    9,
		FadeTime:     150 * time.Millisecond,
		SoftFade:     true,
		MusicTrigger: false,
		Run:          true,
		Bounce:       true,
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
	return sequence
}

func PlayNewSequence(sequence common.Sequence,
	mySequenceNumber int,
	pad *mk2.Launchpad,
	eventsForLauchpad chan common.ALight,
	pattens map[string]common.Patten,
	dmxController ft232.DMXController,
	fixtures *fixture.Fixtures,
	channels common.Channels) {

	positions := map[int]common.Position{}

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
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 0, fixtureChannels, eventsForLauchpad, dmxController, fixtures)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 1, fixtureChannels, eventsForLauchpad, dmxController, fixtures)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 2, fixtureChannels, eventsForLauchpad, dmxController, fixtures)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 3, fixtureChannels, eventsForLauchpad, dmxController, fixtures)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 4, fixtureChannels, eventsForLauchpad, dmxController, fixtures)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 5, fixtureChannels, eventsForLauchpad, dmxController, fixtures)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 6, fixtureChannels, eventsForLauchpad, dmxController, fixtures)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 7, fixtureChannels, eventsForLauchpad, dmxController, fixtures)

	// So this is the outer loop where sequence waits for commands and processes them if we're not playing a sequence.
	// i.e the sequence is in STOP mode and this is the way we change the RUN flag to START a sequence again.
	for {

		sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence.CurrentSpeed*10, sequence, channels)

		if sequence.Run {

			if sequence.Patten.Name == "scanner" {
				sequence.Type = "scanner"
			}

			sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence.CurrentSpeed, sequence, channels)
			if !sequence.Run {
				break
			}

			// Calulate positions for fixtures based on patten.
			positions, sequence.Steps = calculatePositions(pattens[sequence.Patten.Name].Steps, true)

			// Run the sequence through.
			for step := 0; step < sequence.Steps; step++ {
				cmd := common.FixtureCommand{
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
				fixtureChannel2 <- cmd
				fixtureChannel3 <- cmd
				fixtureChannel4 <- cmd
				fixtureChannel5 <- cmd
				fixtureChannel6 <- cmd
				fixtureChannel7 <- cmd
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
func calculatePositions(steps []common.Step, bounce bool) (map[int]common.Position, int) {

	position := common.Position{}

	// We have multiple positions for each fixture.
	var counter int
	positions := make(map[int]common.Position)
	//positions := []common.Position{}

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

					//positions = append(positions, position)
					positions[counter] = position
					// TODO calc actual size based on fade steps.
					counter = counter + 14
				}
			}
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

						//positions = append(positions, position)
						positions[counter] = position
						// TODO calc actual size based on fade steps.
						counter = counter + 14
					}
				}
			}
		}
	}
	return positions, counter
}

// func mapColors(R int, G int, B int, colorSelector int) common.Color {

// 	colorOut := common.Color{}
// 	intensity := findLargest(R, G, B)

// 	if colorSelector == 0 {
// 		colorOut = common.Color{R: R, G: G, B: B}
// 	}
// 	if colorSelector == 1 {
// 		colorOut = common.Color{R: intensity, G: 0, B: 0}
// 	}
// 	if colorSelector == 2 {
// 		colorOut = common.Color{R: 0, G: intensity, B: 0}
// 	}
// 	if colorSelector == 3 {
// 		colorOut = common.Color{R: 0, G: intensity, B: intensity}
// 	}
// 	if colorSelector == 4 {
// 		colorOut = common.Color{R: 0, G: 0, B: intensity}
// 	}
// 	if colorSelector == 5 {
// 		colorOut = common.Color{R: intensity, G: 0, B: intensity}
// 	}
// 	return colorOut
// }

// func findLargest(R int, G int, B int) (answer int) {
// 	/* check the boolean condition using if statement */
// 	if R >= G && R >= B {
// 		return R
// 	}
// 	if G >= R && G >= B {
// 		return G
// 	}
// 	if B >= R && B >= G {
// 		return B
// 	}
// 	return 0
// }
