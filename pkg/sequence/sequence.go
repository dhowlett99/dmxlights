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
		FadeSpeed:    4,
		FadeTime:     500 * time.Millisecond,
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
		sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence.CurrentSpeed/2, sequence, channels)

		if sequence.Run {

			for _, step := range translatePatten(pattens[sequence.Patten.Name].Steps, sequence.FadeSpeed) {
				for fixture := range step.Fixtures {
					for color := range step.Fixtures[fixture].Colors {
						sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence.CurrentSpeed/2, sequence, channels)
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
						dmx.Fixtures(mySequenceNumber, dmxController, fixture, newColor.R, newColor.G, newColor.B, Pan, Tilt, Shutter, Gobo, fixtures, sequence.Blackout, sequence.Master, sequence.Master)
						sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence.CurrentSpeed/2, sequence, channels)
						if !sequence.Run {
							continue
						}
					}
				}
			}

			// for index := len(pattens[sequence.Patten.Name].Steps) - 1; index >= 0; index-- {
			// 	step := pattens[sequence.Patten.Name].Steps[index]
			// 	for fixture := range step.Fixtures {
			// 		for color := range step.Fixtures[fixture].Colors {
			// 			sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence.CurrentSpeed/2, sequence, channels)
			// 			if !sequence.Run {
			// 				continue
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
			// 			sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence.CurrentSpeed/2, sequence, channels)
			// 			if !sequence.Run {
			// 				continue
			// 			}
			// 		}
			// 	}
			// }
		}
	}
}

func PlayNewSequence(sequence common.Sequence,
	mySequenceNumber int,
	pad *mk2.Launchpad,
	eventsForLauchpad chan common.ALight,
	pattens map[string]common.Patten,
	dmxController ft232.DMXController,
	fixtures *fixture.Fixtures,
	channels common.Channels) common.Sequence {

	fixtureCommandChannels := []chan common.FixtureCommand{}
	fadeChannel0 := make(chan common.FixtureCommand)
	fadeChannel1 := make(chan common.FixtureCommand)
	fadeChannel2 := make(chan common.FixtureCommand)
	fadeChannel3 := make(chan common.FixtureCommand)
	fadeChannel4 := make(chan common.FixtureCommand)
	fadeChannel5 := make(chan common.FixtureCommand)
	fadeChannel6 := make(chan common.FixtureCommand)
	fadeChannel7 := make(chan common.FixtureCommand)

	fixtureCommandChannels = append(fixtureCommandChannels, fadeChannel0)
	fixtureCommandChannels = append(fixtureCommandChannels, fadeChannel1)
	fixtureCommandChannels = append(fixtureCommandChannels, fadeChannel2)
	fixtureCommandChannels = append(fixtureCommandChannels, fadeChannel3)
	fixtureCommandChannels = append(fixtureCommandChannels, fadeChannel4)
	fixtureCommandChannels = append(fixtureCommandChannels, fadeChannel5)
	fixtureCommandChannels = append(fixtureCommandChannels, fadeChannel6)
	fixtureCommandChannels = append(fixtureCommandChannels, fadeChannel7)

	for {

		// So this is the outer loop where sequence waits for commands and processes them if we're not playing a sequence.
		// i.e the sequence is in STOP mode and this is the way we change the RUN flag to START a sequence again.
		sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence.CurrentSpeed/2, sequence, channels)

		if sequence.Run {
			for _, step := range pattens[sequence.Patten.Name].Steps {
				//fmt.Printf("Step %d\n", stepIndex)
				for fixture := range step.Fixtures {
					//fmt.Printf("Fixture %d\n", fixture)
					for color := range step.Fixtures[fixture].Colors {
						//fmt.Printf("Color %d\n", color)
						sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence.CurrentSpeed/2, sequence, channels)
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

						go func() {
							newColor := mapColors(R, G, B, sequence.Color)
							if R > 0 || G > 0 || B > 0 {
								// fade up.
								nums := []int{0, 66, 127, 180, 220, 246, 255}

								for _, stepIndex := range nums {
									// Light the lauch pad lamp.
									fireLaunchPadLamp(mySequenceNumber, fixture, stepIndex, eventsForLauchpad, sequence, R, G, B, newColor)

									// Now ask DMX to actually light the real fixture.
									dmx.Fixtures(mySequenceNumber, dmxController, fixture, newColor.R, newColor.G, newColor.B, Pan, Tilt, Shutter, Gobo, fixtures, sequence.Blackout, stepIndex, sequence.Master)

									// Fade up time is set here.
									sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence.FadeTime, sequence, channels)
									if !sequence.Run {
										break
									}
								}

								// Send a message to say we are at the top of the fade.
								fixtureSignal := common.FixtureCommand{
									FadeUp: true,
								}
								fixtureCommandChannels[fixture] <- fixtureSignal

								// Fade on time is set here.
								sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence.FadeTime, sequence, channels)
								if !sequence.Run {
									return
								}
							} else {
								// Now fade down.
								nums := []int{255, 246, 220, 180, 127, 66, 0}

								for _, stepIndex := range nums {
									// Light the lauch pad lamp.
									fireLaunchPadLamp(mySequenceNumber, fixture, stepIndex, eventsForLauchpad, sequence, R, G, B, newColor)

									// Now ask DMX to actually light the real fixture.
									dmx.Fixtures(mySequenceNumber, dmxController, fixture, newColor.R, newColor.G, newColor.B, Pan, Tilt, Shutter, Gobo, fixtures, sequence.Blackout, stepIndex, sequence.Master)

									// Fade down time is set here.
									sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence.FadeTime, sequence, channels)
									if !sequence.Run {
										break
									}
								}

								// Fade off time is set here.
								sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence.FadeTime, sequence, channels)
								if !sequence.Run {
									return
								}

								// Send a message to say we are at the bottom of the fade.
								//fmt.Printf("send we're done\n")
								fixtureSignal := common.FixtureCommand{
									FadeDown: true,
								}
								fixtureCommandChannels[fixture] <- fixtureSignal
							}

						}()
						//time.Sleep(50 * time.Millisecond)
						//c := common.FixtureCommand{}

						//fmt.Printf("c is %+v\n", c)

						// var run bool
						// command := common.FixtureCommand{}
						// for run {

						// 	command = <-fixtureCommandChannels[fixture]
						// 	if command.FadeUp {
						// 		run = false
						// 	}
						// }

						// Now we wait before the bounce happens.
						// sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence.CurrentSpeed/2, sequence, channels)
						// if !sequence.Run {
						// 	continue
						// }
					}

					<-fixtureCommandChannels[fixture]
				}
			}
		}
	}
}

func fireLaunchPadLamp(
	mySequenceNumber int,
	fixture int,
	brightness int,
	eventsForLauchpad chan common.ALight,
	sequence common.Sequence,
	R int,
	G int,
	B int,
	newColor common.Color) {

	e := common.ALight{
		X:          fixture,
		Y:          mySequenceNumber - 1,
		Brightness: brightness,
		Red:        newColor.R,
		Green:      newColor.G,
		Blue:       newColor.B,
	}
	eventsForLauchpad <- e
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

func translatePatten(steps []common.Step, shift int) []common.Step {

	var newStep common.Step
	var newFixture common.Fixture
	//var newColor common.Color

	fadeUp := []int{0, 66, 127, 180, 220, 246, 255}
	fadeDown := []int{255, 246, 220, 189, 127, 66, 0}

	outputSteps := []common.Step{}
	lastStep := common.Step{}
	fadeDownFlag := false

	for _, stepOriginal := range steps {

		//fmt.Printf("working on stepIndex %d  stepOriginal %+v \n", stepIndex, stepOriginal)

		if !fadeDownFlag {

			// Loop around creating new steps for fixture values for fade up.
			for _, newFixtureValue := range fadeUp {

				//fmt.Printf("FADE UP working on newStepIndex %d\n", newFixtureValue)

				// Create a new step.
				newStep = common.Step{}

				// OK we found a fixture in this step, so we now add a bunch of new steps for each of the fade up values.
				for newFixtureIndex, fixture := range stepOriginal.Fixtures {

					//fmt.Printf("working on new newFixtureIndex %d  fixture %+v\n", newFixtureIndex, fixture)

					// Create new fixture.
					newFixture = common.Fixture{}

					// Set the Master Dimmer.
					newFixture.MasterDimmer = fixture.MasterDimmer
					newFixture.Pan = fixture.Pan
					newFixture.Tilt = fixture.Tilt
					newFixture.Shutter = fixture.Shutter
					newFixture.Gobo = fixture.Gobo

					// Add the new fixture.
					newStep.Fixtures = append(newStep.Fixtures, newFixture)

					// // Now we have to match the values from the original fixture to the new fixture and
					// // make the necessary increments.
					for _, color := range fixture.Colors {

						//fmt.Printf("last step %+v\n", lastStep)

						// OK in our last step, look through the fixtures,
						for lastStepFixtureIndex := range lastStep.Fixtures {
							if lastStep.Fixtures != nil {
								if lastStep.Fixtures[lastStepFixtureIndex].Colors != nil {
									for colorIndex := range lastStep.Fixtures[lastStepFixtureIndex].Colors {
										//fmt.Printf("last step colorIndex Red is %+v\n", lastStep.Fixtures[lastStepFixtureIndex].Colors[colorIndex])
										// If we reached full brightness on the last step and we are requesting zero brightness
										// Then we are in a fade down situation and not a fade up.
										// fmt.Printf("UP last %d,  want %d \n", lastStep.Fixtures[lastStepFixtureIndex].Colors[colorIndex].R, color.R)
										if lastStep.Fixtures[lastStepFixtureIndex].Colors[colorIndex].R == 255 && color.R == 0 {
											//fmt.Printf("Whoops we should be fading down instead !!!\n")
											// Lets set the fade down flag.
											fadeDownFlag = true
										} else {
											fadeDownFlag = false
										}
									}
								}
							}
						}

						// if color.R == 0 { //|| color.G == 0 || color.B == 0 {
						newColors := []common.Color{}
						newColor := common.Color{}
						// 	newColors = append(newColors, newColor)
						// 	newStep.Fixtures[newFixtureIndex].Colors = newColors
						// } else {

						// newFixtureValue is essentially a percentage express as 0-255
						if color.R != 0 {
							newColor.R = newFixtureValue
						}
						if color.G != 0 {
							newColor.G = newFixtureValue
						}
						if color.B != 0 {
							newColor.B = newFixtureValue
						}
						newColors = append(newColors, newColor)
						newStep.Fixtures[newFixtureIndex].Colors = newColors

						if fadeDownFlag {
							break
						}

						// Save the state of the step so we can use it to calc if we need to fade up or fade down.
						lastStep = newStep
					}
					if fadeDownFlag {
						break
					}
				}
				if fadeDownFlag {
					continue
				}

				// Add new step to outputSteps.
				outputSteps = append(outputSteps, newStep)
			}

			if fadeDownFlag {
				// Loop around creating new steps for fixture values for fade down.
				for _, newFixtureValue := range fadeDown {
					//fmt.Printf("FADE DOWN working on newStepIndex %d\n", newFixtureValue)

					// Create a new step.
					newStep = common.Step{}

					// OK we found a fixture in this step, so we now add a bunch of new steps for each of the fade up values.
					for newFixtureIndex, fixture := range stepOriginal.Fixtures {

						//fmt.Printf("working on new newFixtureIndex %d  fixture %+v\n", newFixtureIndex, fixture)

						// Create new fixture.
						newFixture = common.Fixture{}

						// Set the Master Dimmer.
						newFixture.MasterDimmer = fixture.MasterDimmer
						newFixture.Pan = fixture.Pan
						newFixture.Tilt = fixture.Tilt
						newFixture.Shutter = fixture.Shutter
						newFixture.Gobo = fixture.Gobo

						// Add the new fixture.
						newStep.Fixtures = append(newStep.Fixtures, newFixture)

						// // Now we have to match the values from the original fixture to the new fixture and
						// // make the necessary increments.
						for _, color := range fixture.Colors {
							//fmt.Printf("color is %+v\n", color)
							// OK in our last step, look through the fixtures,
							for lastStepFixtureIndex := range lastStep.Fixtures {
								if lastStep.Fixtures != nil {
									if lastStep.Fixtures[lastStepFixtureIndex].Colors != nil {
										for colorIndex := range lastStep.Fixtures[lastStepFixtureIndex].Colors {
											//fmt.Printf("last step colorIndex Red is %+v\n", lastStep.Fixtures[lastStepFixtureIndex].Colors[colorIndex])
											// If we reached full brightness on the last step and we are requesting zero brightness
											// Then we are in a fade down situation and not a fade up.
											//fmt.Printf("DOWN last %d,  want %d \n", lastStep.Fixtures[lastStepFixtureIndex].Colors[colorIndex].R, color.R)
											if lastStep.Fixtures[lastStepFixtureIndex].Colors[colorIndex].R == 0 && color.R == 255 {
												//fmt.Printf("Whoops we should be fading up instead !!!\n")
												// Lets set the fade up flag.
												fadeDownFlag = false
											} else {
												fadeDownFlag = true
											}
										}
									}
								}
							}

							// if lastStep.Fixtures[newFixtureIndex].Colors[colorIndex].R == 255 && color.R == 0 {
							// 	fmt.Printf("Fade up\n")
							// }

							//fmt.Printf("working on new colors %d  color %+v\n", colorIndex, color)
							// Map the color.

							if color.R == 0 && color.G == 0 && color.B == 0 {
								newColors := []common.Color{}
								newColor := common.Color{}
								newColors = append(newColors, newColor)
								newStep.Fixtures[newFixtureIndex].Colors = newColors
							} else {
								newColors := []common.Color{}
								newColor := common.Color{}
								// newFixtureValue is essentially a percentage express as 0-255
								//newColor.R = newFixtureValue
								if color.R > 0 {
									newColor.R = newFixtureValue
								}

								if color.G > 0 {
									newColor.G = newFixtureValue
								}
								if color.B > 0 {
									newColor.B = newFixtureValue
								}
								newColors = append(newColors, newColor)
								newStep.Fixtures[newFixtureIndex].Colors = newColors
							}

							if !fadeDownFlag {
								continue
							}

							// Save the state of the step so we can use it to calc if we need to fade up or fade down.
							lastStep = newStep
						}
					}

					// Add new step to outputSteps.
					outputSteps = append(outputSteps, newStep)
				}
			}
			fadeDownFlag = false
		}
	}
	return outputSteps
}
