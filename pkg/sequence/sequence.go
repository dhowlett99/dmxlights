package sequence

import (
	"math"
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
		Run:          true,
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
		Master:   255,
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
						dmx.Fixtures(mySequenceNumber, dmxController, fixture, newColor.R, newColor.G, newColor.B, Pan, Tilt, Shutter, Gobo, fixtures, sequence.Blackout, sequence.Master)
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
						dmx.Fixtures(mySequenceNumber, dmxController, fixture, newColor.R, newColor.G, newColor.B, Pan, Tilt, Shutter, Gobo, fixtures, sequence.Blackout, sequence.Master)
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

func PlayShiftableSequence(sequence common.Sequence,
	mySequenceNumber int,
	pad *mk2.Launchpad,
	eventsForLauchpad chan common.ALight,
	pattens map[string]common.Patten,
	dmxController ft232.DMXController,
	fixtures *fixture.Fixtures,
	channels common.Channels) common.Sequence {

	for {

		// So this is the outer loop where sequence thread waits for commands and processes them if we're not playing a sequence.
		// i.e the sequence is in STOP mode and this is the way we change the RUN flag to START a sequence again.
		sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence, channels)
		if sequence.Run {

			// Create an array of steps.
			steps := []common.Step{}

			steps = fade(steps)

			for s := 0; s < len(steps); s++ {

				sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence, channels)
				if !sequence.Run {
					continue
				}

				f := calcFade(0, len(fixtures.Fixtures)-1, sequence.FadeSpeed)

				for _, color := range steps[s].Fixtures[f].Colors {
					R := color.R
					G := color.G
					B := color.B
					Pan := steps[0].Fixtures[0].Pan
					Tilt := steps[0].Fixtures[0].Tilt
					Shutter := steps[0].Fixtures[0].Shutter
					Gobo := steps[0].Fixtures[0].Gobo
					sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence, channels)
					if !sequence.Run {
						break
					}

					playFixture(mySequenceNumber, 0, 255, R, G, B, Pan, Tilt, Shutter, Gobo, fixtures, sequence, eventsForLauchpad, dmxController)
				}

				s = calcFade(s, len(steps)-1, sequence.FadeSpeed)
				f = calcFade(0, len(fixtures.Fixtures)-1, sequence.FadeSpeed)

				for _, color := range steps[s].Fixtures[f].Colors {
					R := color.R
					G := color.G
					B := color.B
					Pan := steps[0].Fixtures[0].Pan
					Tilt := steps[0].Fixtures[0].Tilt
					Shutter := steps[0].Fixtures[0].Shutter
					Gobo := steps[0].Fixtures[0].Gobo
					sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence, channels)
					if !sequence.Run {
						break
					}

					playFixture(mySequenceNumber, 1, 255, R, G, B, Pan, Tilt, Shutter, Gobo, fixtures, sequence, eventsForLauchpad, dmxController)
				}

				s = calcFade(s, len(steps)-1, sequence.FadeSpeed)
				f = calcFade(0, len(fixtures.Fixtures)-1, sequence.FadeSpeed)

				for _, color := range steps[s].Fixtures[f].Colors {
					R := color.R
					G := color.G
					B := color.B
					Pan := steps[0].Fixtures[0].Pan
					Tilt := steps[0].Fixtures[0].Tilt
					Shutter := steps[0].Fixtures[0].Shutter
					Gobo := steps[0].Fixtures[0].Gobo
					sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence, channels)
					if !sequence.Run {
						break
					}
					playFixture(mySequenceNumber, 2, 255, R, G, B, Pan, Tilt, Shutter, Gobo, fixtures, sequence, eventsForLauchpad, dmxController)
				}

				s = calcFade(s, len(steps)-1, sequence.FadeSpeed)
				f = calcFade(0, len(fixtures.Fixtures)-1, sequence.FadeSpeed)

				for _, color := range steps[s].Fixtures[f].Colors {
					R := color.R
					G := color.G
					B := color.B
					Pan := steps[0].Fixtures[0].Pan
					Tilt := steps[0].Fixtures[0].Tilt
					Shutter := steps[0].Fixtures[0].Shutter
					Gobo := steps[0].Fixtures[0].Gobo
					sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence, channels)
					if !sequence.Run {
						break
					}
					playFixture(mySequenceNumber, 3, 255, R, G, B, Pan, Tilt, Shutter, Gobo, fixtures, sequence, eventsForLauchpad, dmxController)
				}

				s = calcFade(s, len(steps)-1, sequence.FadeSpeed)
				f = calcFade(0, len(fixtures.Fixtures)-1, sequence.FadeSpeed)

				for _, color := range steps[s].Fixtures[f].Colors {
					R := color.R
					G := color.G
					B := color.B
					Pan := steps[0].Fixtures[0].Pan
					Tilt := steps[0].Fixtures[0].Tilt
					Shutter := steps[0].Fixtures[0].Shutter
					Gobo := steps[0].Fixtures[0].Gobo
					sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence, channels)
					if !sequence.Run {
						break
					}
					playFixture(mySequenceNumber, 4, 255, R, G, B, Pan, Tilt, Shutter, Gobo, fixtures, sequence, eventsForLauchpad, dmxController)
				}

				s = calcFade(s, len(steps)-1, sequence.FadeSpeed)
				f = calcFade(0, len(fixtures.Fixtures)-1, sequence.FadeSpeed)

				for _, color := range steps[s].Fixtures[f].Colors {
					R := color.R
					G := color.G
					B := color.B
					Pan := steps[0].Fixtures[0].Pan
					Tilt := steps[0].Fixtures[0].Tilt
					Shutter := steps[0].Fixtures[0].Shutter
					Gobo := steps[0].Fixtures[0].Gobo
					sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence, channels)
					if !sequence.Run {
						break
					}

					playFixture(mySequenceNumber, 5, 255, R, G, B, Pan, Tilt, Shutter, Gobo, fixtures, sequence, eventsForLauchpad, dmxController)
				}

				s = calcFade(s, len(steps)-1, sequence.FadeSpeed)
				f = calcFade(0, len(fixtures.Fixtures)-1, sequence.FadeSpeed)

				for _, color := range steps[s].Fixtures[f].Colors {
					R := color.R
					G := color.G
					B := color.B
					Pan := steps[0].Fixtures[0].Pan
					Tilt := steps[0].Fixtures[0].Tilt
					Shutter := steps[0].Fixtures[0].Shutter
					Gobo := steps[0].Fixtures[0].Gobo
					sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence, channels)
					if !sequence.Run {
						break
					}

					playFixture(mySequenceNumber, 6, 255, R, G, B, Pan, Tilt, Shutter, Gobo, fixtures, sequence, eventsForLauchpad, dmxController)
				}

				s = calcFade(s, len(steps)-1, sequence.FadeSpeed)
				f = calcFade(0, len(fixtures.Fixtures)-1, sequence.FadeSpeed)

				for _, color := range steps[s].Fixtures[f].Colors {
					R := color.R
					G := color.G
					B := color.B
					Pan := steps[0].Fixtures[0].Pan
					Tilt := steps[0].Fixtures[0].Tilt
					Shutter := steps[0].Fixtures[0].Shutter
					Gobo := steps[0].Fixtures[0].Gobo
					sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence, channels)
					if !sequence.Run {
						break
					}

					playFixture(mySequenceNumber, 7, 255, R, G, B, Pan, Tilt, Shutter, Gobo, fixtures, sequence, eventsForLauchpad, dmxController)
				}
			}
		}
	}
}

func radians(input float64) float64 {
	result := input * math.Pi / 180
	return result
}

func fade(steps []common.Step) []common.Step {

	nums := []int{0, 0, 0, 66, 127, 180, 220, 246, 255, 255, 255, 255, 246, 220, 180, 127, 66}
	//nums := []int{0, 0, 0, 66, 86, 97, 127, 155, 180, 220, 246, 255, 255, 255, 255, 255, 255, 255, 255, 255, 246, 220, 180, 155, 127, 97, 86, 66}
	//nums := []int{0, 0, 0, 10, 25, 45, 66, 70, 85, 90, 127, 145, 165, 180, 195, 220, 225, 232, 246, 255, 255, 255, 255, 255, 255, 255, 255, 255, 232, 225, 220, 195, 180, 165, 145, 127, 90, 85, 70, 66, 45, 25, 10}

	//nums := []int{0, 255}
	// degrees := []float64{0, 15, 30, 45, 60, 75, 90}
	// for _, x := range degrees {
	// 	fmt.Printf("radians %.00f\n", math.Sin(radians(x))*100)
	// }

	for _, stepIndex := range nums {

		//fmt.Printf("stepIndex is %d\n", stepIndex)
		// Create a step
		step := common.Step{}

		// Populate the fixtures.
		for fixtureIndex := 0; fixtureIndex < 8; fixtureIndex++ {

			// Create an array of Fixtures for this step.
			fixture := common.Fixture{
				MasterDimmer: 255,
				Colors: []common.Color{
					{
						R: stepIndex,
						G: 0,
						B: 0,
					},
				},
			}

			// Add fixtures to step.
			step.Fixtures = append(step.Fixtures, fixture)
		}

		// Add step to steps array.
		steps = append(steps, step)
	}
	return steps
}

//playFixture(mySequenceNumber, fixture, R, G, B, Pan, Tilt, Shutter, Gobo, fixtures, sequence, eventsForLauchpad, dmxController)
func playFixture(
	mySequenceNumber int,
	fixture int,
	brightness int,
	R int,
	G int,
	B int,
	Pan int,
	Tilt int,
	Shutter int,
	Gobo int,
	fixtures *fixture.Fixtures,
	sequence common.Sequence,
	eventsForLauchpad chan common.ALight,
	dmxController ft232.DMXController) {

	newColor := mapColors(R, G, B, sequence.Color)

	e := common.ALight{
		X:          fixture,
		Y:          mySequenceNumber - 1,
		Brightness: brightness,
		Red:        newColor.R,
		Green:      newColor.G,
		Blue:       newColor.B,
	}
	eventsForLauchpad <- e

	// Now ask DMX to actually light the real fixture.
	dmx.Fixtures(mySequenceNumber, dmxController, fixture, newColor.R, newColor.G, newColor.B, Pan, Tilt, Shutter, Gobo, fixtures, sequence.Blackout, brightness)

}

func calcFade(index, max, fade int) int {

	out := index + fade
	if out > max {
		out = 0
	}
	///fmt.Printf("index %d  max %d fade %d out %d\n", index, max, fade, out)
	return out
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
