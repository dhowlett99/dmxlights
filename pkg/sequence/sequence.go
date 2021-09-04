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
		FadeSpeed:    9,
		FadeTime:     150 * time.Millisecond,
		SoftFade:     true,
		MusicTrigger: false,
		Run:          true,
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

	// Create eight fixtures.
	go makeFixture(sequence, mySequenceNumber, 0, fixtureChannels, eventsForLauchpad, dmxController, fixtures)
	go makeFixture(sequence, mySequenceNumber, 1, fixtureChannels, eventsForLauchpad, dmxController, fixtures)
	go makeFixture(sequence, mySequenceNumber, 2, fixtureChannels, eventsForLauchpad, dmxController, fixtures)
	go makeFixture(sequence, mySequenceNumber, 3, fixtureChannels, eventsForLauchpad, dmxController, fixtures)
	go makeFixture(sequence, mySequenceNumber, 4, fixtureChannels, eventsForLauchpad, dmxController, fixtures)
	go makeFixture(sequence, mySequenceNumber, 5, fixtureChannels, eventsForLauchpad, dmxController, fixtures)
	go makeFixture(sequence, mySequenceNumber, 6, fixtureChannels, eventsForLauchpad, dmxController, fixtures)
	go makeFixture(sequence, mySequenceNumber, 7, fixtureChannels, eventsForLauchpad, dmxController, fixtures)

	// So this is the outer loop where sequence waits for commands and processes them if we're not playing a sequence.
	// i.e the sequence is in STOP mode and this is the way we change the RUN flag to START a sequence again.
	for {

		sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence.CurrentSpeed*10, sequence, channels)

		if sequence.Run {

			sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence.CurrentSpeed, sequence, channels)
			if !sequence.Run {
				break
			}

			// Calulate positions for fixtures based on patten.
			positions := calculatePositions(pattens[sequence.Patten.Name].Steps)

			// TODO actually caluclate the number of steps required based on the calc above.
			noSteps := 8 * 14

			// Run the sequence through.
			for step := 0; step < noSteps; step++ {

				cmd := common.FixtureCommand{
					Tick:            true,
					Positions:       positions,
					FadeSpeed:       sequence.FadeSpeed,
					FadeTime:        sequence.FadeTime,
					Size:            sequence.Size,
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

func calculatePositions(steps []common.Step) []common.Position {

	position := common.Position{}

	// We have multiple positions for each fixture.
	var counter int
	positions := []common.Position{}

	for _, step := range steps {
		for fixtureIndex, fixture := range step.Fixtures {
			for _, color := range fixture.Colors {
				if color.R > 0 || color.G > 0 || color.B > 0 {
					position.StartPosition = counter
					position.Fixture = fixtureIndex
					position.Color.R = color.R
					position.Color.G = color.G
					position.Color.B = color.B
					positions = append(positions, position)
					// TODO calc actual size based on fade steps.
					counter = counter + 14
				}
			}
		}
	}
	return positions
}

func makeFixture(sequence common.Sequence,
	mySequenceNumber int,
	myFixtureNumber int,
	channels []chan common.FixtureCommand,
	eventsForLauchpad chan common.ALight,
	dmxController ft232.DMXController,
	fixtures *fixture.Fixtures) {

	cmd := common.FixtureCommand{}
	fadeUp := []int{0, 66, 127, 180, 220, 246, 255}
	fadeDown := []int{255, 246, 220, 189, 127, 66, 0}

	for {
		select {
		case cmd = <-channels[myFixtureNumber]:
			if cmd.Tick {
				for _, position := range cmd.Positions {
					if cmd.CurrentPosition == position.StartPosition {
						if position.Fixture == myFixtureNumber {
							// Now kick off the back end which drives the fixture.
							go func() {
								for _, value := range fadeUp {
									time.Sleep(cmd.FadeTime / 4)
									R := int((float64(position.Color.R) / 100) * (float64(value) / 2.55))
									G := int((float64(position.Color.G) / 100) * (float64(value) / 2.55))
									B := int((float64(position.Color.B) / 100) * (float64(value) / 2.55))
									Pan := position.Pan
									Tilt := position.Tilt
									Shutter := position.Shutter
									Gobo := position.Gobo
									lightLamp(mySequenceNumber, myFixtureNumber, R, G, B, eventsForLauchpad)
									// Now ask DMX to actually light the real fixture.
									dmx.Fixtures(mySequenceNumber, dmxController, myFixtureNumber, R, G, B, Pan, Tilt, Shutter, Gobo, fixtures, cmd.Blackout, sequence.Master, sequence.Master)
								}
								for x := 0; x < cmd.Size; x++ {
									time.Sleep(cmd.CurrentSpeed * 5)
								}
								time.Sleep(cmd.FadeTime / 4)
								for _, value := range fadeDown {
									R := int((float64(position.Color.R) / 100) * (float64(value) / 2.55))
									G := int((float64(position.Color.G) / 100) * (float64(value) / 2.55))
									B := int((float64(position.Color.B) / 100) * (float64(value) / 2.55))
									Pan := position.Pan
									Tilt := position.Tilt
									Shutter := position.Shutter
									Gobo := position.Gobo
									lightLamp(mySequenceNumber, myFixtureNumber, R, G, B, eventsForLauchpad)
									dmx.Fixtures(mySequenceNumber, dmxController, myFixtureNumber, R, G, B, Pan, Tilt, Shutter, Gobo, fixtures, cmd.Blackout, sequence.Master, sequence.Master)
									time.Sleep(cmd.FadeTime / 4)
								}
								time.Sleep(cmd.FadeTime / 4)
							}()
						}
					}
				}
			}
		}
	}
}

func lightLamp(X, Y, R, G, B int, eventsForLauchpad chan common.ALight) {

	// Now trigger the fixture lamp on the launch pad by sending an event.
	e := common.ALight{
		X:          Y,
		Y:          X - 1,
		Brightness: 255,
		Red:        R,
		Green:      G,
		Blue:       B,
	}
	eventsForLauchpad <- e
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

		// Select the patten name.
		steps := pattens[sequence.Patten.Name].Steps

		// So this is the outer loop where sequence waits for commands and processes them if we're not playing a sequence.
		// i.e the sequence is in STOP mode and this is the way we change the RUN flag to START a sequence again.
		sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence.CurrentSpeed/2, sequence, channels)

		if sequence.Run {

			if sequence.Patten.Name == "scanner" {
				steps = pattens[sequence.Patten.Name].Steps
			}

			for _, step := range steps {
				for fixture := range step.Fixtures {
					for color := range step.Fixtures[fixture].Colors {
						sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence.CurrentSpeed, sequence, channels)
						if !sequence.Run {
							break
						}
						R := step.Fixtures[fixture].Colors[color].R
						G := step.Fixtures[fixture].Colors[color].G
						B := step.Fixtures[fixture].Colors[color].B
						Pan := step.Fixtures[fixture].Pan
						Tilt := step.Fixtures[fixture].Tilt
						Shutter := step.Fixtures[fixture].Shutter
						Gobo := step.Fixtures[fixture].Gobo

						newColor := mapColors(R, G, B, sequence.Color)
						// Now trigger the fixture lamp on the launch pad by sending an event.
						e := common.ALight{
							X:          fixture,
							Y:          mySequenceNumber - 1,
							Brightness: 255,
							Red:        newColor.R,
							Green:      newColor.G,
							Blue:       newColor.B,
						}
						eventsForLauchpad <- e

						// Now ask DMX to actually light the real fixture.
						dmx.Fixtures(mySequenceNumber, dmxController, fixture, newColor.R, newColor.G, newColor.B, Pan, Tilt, Shutter, Gobo, fixtures, sequence.Blackout, sequence.Master, sequence.Master)
						sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence.CurrentSpeed, sequence, channels)
						if !sequence.Run {
							break
						}
					}
					if !sequence.Run {
						break
					}
				}
				if !sequence.Run {
					break
				}
			}

			// for index := len(steps) - 1; index >= 0; index-- {
			// 	step := steps[index]
			// 	for fixture := range step.Fixtures {
			// 		for color := range step.Fixtures[fixture].Colors {
			// 			sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence.CurrentSpeed, sequence, channels)
			// 			if !sequence.Run {
			// 				break
			// 			}
			// 			R := step.Fixtures[fixture].Colors[color].R
			// 			G := step.Fixtures[fixture].Colors[color].G
			// 			B := step.Fixtures[fixture].Colors[color].B
			// 			Pan := step.Fixtures[fixture].Pan
			// 			Tilt := step.Fixtures[fixture].Tilt
			// 			Shutter := step.Fixtures[fixture].Shutter
			// 			Gobo := step.Fixtures[fixture].Tilt

			// 			newColor := mapColors(R, G, B, sequence.Color)
			// 			// Now trigger the fixture lamp on the launch pad by sending an event.
			// 			e := common.ALight{
			// 				X:          fixture,
			// 				Y:          mySequenceNumber - 1,
			// 				Brightness: 255,
			// 				Red:        newColor.R,
			// 				Green:      newColor.G,
			// 				Blue:       newColor.B,
			// 			}
			// 			eventsForLauchpad <- e

			// 			// Now ask DMX to actually light the real fixture.
			// 			dmx.Fixtures(mySequenceNumber, dmxController, fixture, newColor.R, newColor.G, newColor.B, Pan, Tilt, Shutter, Gobo, fixtures, sequence.Blackout, sequence.Master, sequence.Master)
			// 			sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence.CurrentSpeed, sequence, channels)
			// 			if !sequence.Run {
			// 				break
			// 			}
			// 		}
			// 		if !sequence.Run {
			// 			break
			// 		}
			// 	}
			// 	if !sequence.Run {
			// 		break
			// 	}
			// }
		}
	}
}

func mapColors(R int, G int, B int, colorSelector int) common.Color {

	colorOut := common.Color{}
	intensity := findLargest(R, G, B)

	if colorSelector == 0 {
		colorOut = common.Color{R: R, G: G, B: B}
	}
	if colorSelector == 1 {
		colorOut = common.Color{R: intensity, G: 0, B: 0}
	}
	if colorSelector == 2 {
		colorOut = common.Color{R: 0, G: intensity, B: 0}
	}
	if colorSelector == 3 {
		colorOut = common.Color{R: 0, G: intensity, B: intensity}
	}
	if colorSelector == 4 {
		colorOut = common.Color{R: 0, G: 0, B: intensity}
	}
	if colorSelector == 5 {
		colorOut = common.Color{R: intensity, G: 0, B: intensity}
	}
	return colorOut
}

func findLargest(R int, G int, B int) (answer int) {
	/* check the boolean condition using if statement */
	if R >= G && R >= B {
		return R
	}
	if G >= R && G >= B {
		return G
	}
	if B >= R && B >= G {
		return B
	}
	return 0
}
