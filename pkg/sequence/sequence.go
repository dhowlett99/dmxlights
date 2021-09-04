package sequence

import (
	"fmt"
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

	// Create eight fixtures.

	// Create eight channels to control the fixtures.
	fixtureChannel1 := make(chan common.FixtureCommand)
	fixtureChannel2 := make(chan common.FixtureCommand)
	fixtureChannel3 := make(chan common.FixtureCommand)
	fixtureChannel4 := make(chan common.FixtureCommand)
	fixtureChannel5 := make(chan common.FixtureCommand)
	fixtureChannel6 := make(chan common.FixtureCommand)
	fixtureChannel7 := make(chan common.FixtureCommand)
	fixtureChannel8 := make(chan common.FixtureCommand)

	fixtureChannels := []chan common.FixtureCommand{}
	fixtureChannels = append(fixtureChannels, fixtureChannel1)
	fixtureChannels = append(fixtureChannels, fixtureChannel2)
	fixtureChannels = append(fixtureChannels, fixtureChannel3)
	fixtureChannels = append(fixtureChannels, fixtureChannel4)
	fixtureChannels = append(fixtureChannels, fixtureChannel5)
	fixtureChannels = append(fixtureChannels, fixtureChannel6)
	fixtureChannels = append(fixtureChannels, fixtureChannel7)
	fixtureChannels = append(fixtureChannels, fixtureChannel8)

	go makeFixture(sequence, mySequenceNumber, 0, fixtureChannels, eventsForLauchpad)
	go makeFixture(sequence, mySequenceNumber, 1, fixtureChannels, eventsForLauchpad)
	go makeFixture(sequence, mySequenceNumber, 2, fixtureChannels, eventsForLauchpad)
	go makeFixture(sequence, mySequenceNumber, 3, fixtureChannels, eventsForLauchpad)
	go makeFixture(sequence, mySequenceNumber, 4, fixtureChannels, eventsForLauchpad)
	go makeFixture(sequence, mySequenceNumber, 5, fixtureChannels, eventsForLauchpad)
	go makeFixture(sequence, mySequenceNumber, 6, fixtureChannels, eventsForLauchpad)
	go makeFixture(sequence, mySequenceNumber, 7, fixtureChannels, eventsForLauchpad)

	for {

		//steps := pattens[sequence.Patten.Name].Steps
		//steps := translatePatten(pattens[sequence.Patten.Name].Steps, sequence.FadeSpeed)

		// So this is the outer loop where sequence waits for commands and processes them if we're not playing a sequence.
		// i.e the sequence is in STOP mode and this is the way we change the RUN flag to START a sequence again.
		sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence.CurrentSpeed*10, sequence, channels)

		if sequence.Run {

			sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence.CurrentSpeed, sequence, channels)
			if !sequence.Run {
				break
			}

			// Calulate positions for fixtures based on patten.
			positions := calculatePositions(sequence, pattens[sequence.Patten.Name])

			// Now we have calculates the positions for the fixtures we must config each fixture.
			for index, position := range positions {
				cmd := common.FixtureCommand{
					Config:        true,
					StartPosition: position.StartPosition,
					Color:         position.Color,
				}
				fixtureChannels[index] <- cmd
			}

			noSteps := 8 * 14

			// Run the sequence through.
			for step := 0; step < noSteps; step++ {

				cmd := common.FixtureCommand{
					Tick:         true,
					FadeSpeed:    sequence.FadeSpeed,
					FadeTime:     sequence.FadeTime,
					Size:         sequence.Size,
					CurrentSpeed: sequence.CurrentSpeed,
					Speed:        sequence.Speed,
					CurrentPosition: common.Position{
						StartPosition: step,
					},
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

func calculatePositions(sequence common.Sequence, patten common.Patten) (positions []common.Position) {

	position := common.Position{}
	var counter int
	counter = 0
	for x := 0; x < 8; x++ {
		// It takes 7 steps to fade up and 7 steps to fade down so 14 for each fixture.
		position.StartPosition = counter
		position.Color.R = 0
		position.Color.G = 255
		position.Color.B = 0
		positions = append(positions, position)
		counter = counter + 14
	}
	return positions
}

func makeFixture(sequence common.Sequence, mySequenceNumber int, myFixtureNumber int, channels []chan common.FixtureCommand, eventsForLauchpad chan common.ALight) {

	cmd := common.FixtureCommand{}
	fadeUp := []int{0, 66, 127, 180, 220, 246, 255}
	fadeDown := []int{255, 246, 220, 189, 127, 66, 0}

	var startPosition int
	var color common.Color

	for {
		select {
		case cmd = <-channels[myFixtureNumber]:
			if cmd.Config {
				// Configure the position this fixture will light
				startPosition = cmd.StartPosition
				color = cmd.Color
			}
			if cmd.Tick {
				if cmd.CurrentPosition.StartPosition == startPosition {
					// Now kick off the back end which drives the fixture.
					go func() {
						//fmt.Printf("Fixture %d FADE UP at Positions %d\n", myFixtureNumber, cmd.CurrentPosition)
						for _, value := range fadeUp {
							time.Sleep(cmd.FadeTime / 3)
							R := int((float64(color.R) / 100) * (float64(value) / 2.55))
							G := int((float64(color.G) / 100) * (float64(value) / 2.55))
							B := int((float64(color.B) / 100) * (float64(value) / 2.55))
							lightLamp(mySequenceNumber, myFixtureNumber, R, G, B, eventsForLauchpad)
						}
						//fmt.Printf("-----> Size %d\n", cmd.Size)
						for x := 0; x < cmd.Size; x++ {
							time.Sleep(cmd.CurrentSpeed * 5)
						}
						time.Sleep(cmd.FadeTime / 3)
						//fmt.Printf("Fixture %d FADE DOWN\n", myFixtureNumber)
						for _, value := range fadeDown {
							R := int((float64(cmd.CurrentPosition.Color.R) / 100) * (float64(value) / 2.55))
							G := int((float64(cmd.CurrentPosition.Color.G) / 100) * (float64(value) / 2.55))
							B := int((float64(cmd.CurrentPosition.Color.B) / 100) * (float64(value) / 2.55))
							lightLamp(mySequenceNumber, myFixtureNumber, R, G, B, eventsForLauchpad)
							time.Sleep(cmd.FadeTime / 3)
						}
						time.Sleep(cmd.FadeTime / 3)
					}()
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

		steps := pattens[sequence.Patten.Name].Steps
		//steps := translatePatten(pattens[sequence.Patten.Name].Steps, sequence.FadeSpeed)

		// So this is the outer loop where sequence waits for commands and processes them if we're not playing a sequence.
		// i.e the sequence is in STOP mode and this is the way we change the RUN flag to START a sequence again.
		sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence.CurrentSpeed/2, sequence, channels)

		if sequence.Run {

			// if sequence.SoftFade {

			// 	sequence.CurrentSpeed = sequence.CurrentSpeed * 50
			// } else {
			// 	steps = pattens[sequence.Patten.Name].Steps
			// }

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
		if !fadeDownFlag {
			// Loop around creating new steps for fixture values for fade up.
			for _, newFixtureValue := range fadeUp {

				// Create a new step.
				newStep = common.Step{}

				// OK we found a fixture in this step, so we now add a bunch of new steps for each of the fade up values.
				for newFixtureIndex, fixture := range stepOriginal.Fixtures {

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
											// If we reached full brightness on the last step and we are requesting zero brightness
											// Then we are in a fade down situation and not a fade up.
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

type NewFixture struct {
	colors []common.Color
}

func shiftPatten(steps []common.Step, shift int) []common.Step {

	fixture1 := []common.Color{}
	fixture2 := []common.Color{}
	fixture3 := []common.Color{}
	fixture4 := []common.Color{}
	fixture5 := []common.Color{}
	fixture6 := []common.Color{}
	fixture7 := []common.Color{}
	fixture8 := []common.Color{}

	// Find the values of all the fixtures.
	for _, step := range steps {
		//fmt.Printf("Step no:%d \n", stepIndex)
		for fixtureIndex, fixture := range step.Fixtures {

			if fixtureIndex == 0 {

				//fmt.Printf("\tFixture Number %d, value %+v\n", fixtureIndex, fixture)

				for _, color := range fixture.Colors {
					//fmt.Printf("\t\tcolorIndex:%d color:%+v\n", colorIndex, color.R)
					newColor := common.Color{}
					newColor.R = color.R
					newColor.G = color.G
					newColor.B = color.B
					fixture1 = append(fixture1, newColor)
				}
			}
			if fixtureIndex == 1 {

				//fmt.Printf("\tFixture Number %d, value %+v\n", fixtureIndex, fixture)

				for _, color := range fixture.Colors {
					//fmt.Printf("\t\tcolorIndex:%d color:%+v\n", colorIndex, color.R)
					newColor := common.Color{}
					newColor.R = color.R
					newColor.G = color.G
					newColor.B = color.B
					fixture2 = append(fixture2, newColor)
				}
			}
			if fixtureIndex == 2 {

				//fmt.Printf("\tFixture Number %d, value %+v\n", fixtureIndex, fixture)

				for _, color := range fixture.Colors {
					//fmt.Printf("\t\tcolorIndex:%d color:%+v\n", colorIndex, color.R)
					newColor := common.Color{}
					newColor.R = color.R
					newColor.G = color.G
					newColor.B = color.B
					fixture3 = append(fixture3, newColor)
				}
			}
			if fixtureIndex == 3 {

				//fmt.Printf("\tFixture Number %d, value %+v\n", fixtureIndex, fixture)

				for _, color := range fixture.Colors {
					//fmt.Printf("\t\tcolorIndex:%d color:%+v\n", colorIndex, color.R)
					newColor := common.Color{}
					newColor.R = color.R
					newColor.G = color.G
					newColor.B = color.B
					fixture4 = append(fixture4, newColor)
				}
			}
			if fixtureIndex == 4 {

				//fmt.Printf("\tFixture Number %d, value %+v\n", fixtureIndex, fixture)

				for _, color := range fixture.Colors {
					//fmt.Printf("\t\tcolorIndex:%d color:%+v\n", colorIndex, color.R)
					newColor := common.Color{}
					newColor.R = color.R
					newColor.G = color.G
					newColor.B = color.B
					fixture5 = append(fixture5, newColor)
				}
			}
			if fixtureIndex == 5 {

				//fmt.Printf("\tFixture Number %d, value %+v\n", fixtureIndex, fixture)

				for _, color := range fixture.Colors {
					//fmt.Printf("\t\tcolorIndex:%d color:%+v\n", colorIndex, color.R)
					newColor := common.Color{}
					newColor.R = color.R
					newColor.G = color.G
					newColor.B = color.B
					fixture6 = append(fixture6, newColor)
				}
			}
			if fixtureIndex == 6 {

				//fmt.Printf("\tFixture Number %d, value %+v\n", fixtureIndex, fixture)

				for _, color := range fixture.Colors {
					//fmt.Printf("\t\tcolorIndex:%d color:%+v\n", colorIndex, color.R)
					newColor := common.Color{}
					newColor.R = color.R
					newColor.G = color.G
					newColor.B = color.B
					fixture7 = append(fixture7, newColor)
				}
			}
			if fixtureIndex == 7 {

				//fmt.Printf("\tFixture Number %d, value %+v\n", fixtureIndex, fixture)

				for _, color := range fixture.Colors {
					//fmt.Printf("\t\tcolorIndex:%d color:%+v\n", colorIndex, color.R)
					newColor := common.Color{}
					newColor.R = color.R
					newColor.G = color.G
					newColor.B = color.B
					fixture8 = append(fixture8, newColor)
				}
			}
		}
	}
	//fmt.Printf("values %v\n", fixture1)
	//fmt.Printf("values %v\n", fixture2)

	// Create an array of all the new fixture values.
	NewFixturesValues := []NewFixture{}
	newF1 := NewFixture{colors: fixture1}
	newF2 := NewFixture{colors: fixture2}
	newF3 := NewFixture{colors: fixture3}
	newF4 := NewFixture{colors: fixture4}
	newF5 := NewFixture{colors: fixture5}
	newF6 := NewFixture{colors: fixture6}
	newF7 := NewFixture{colors: fixture7}
	newF8 := NewFixture{colors: fixture8}
	NewFixturesValues = append(NewFixturesValues, newF1)
	NewFixturesValues = append(NewFixturesValues, newF2)
	NewFixturesValues = append(NewFixturesValues, newF3)
	NewFixturesValues = append(NewFixturesValues, newF4)
	NewFixturesValues = append(NewFixturesValues, newF5)
	NewFixturesValues = append(NewFixturesValues, newF6)
	NewFixturesValues = append(NewFixturesValues, newF7)
	NewFixturesValues = append(NewFixturesValues, newF8)

	// Write out the values of the fixture applying the shift.
	stepsOut := []common.Step{}

	// Now using the original step list to recreate the modified shifted fixture list.
	for stepIndex, step := range steps {

		//fmt.Printf("step no %d\n", stepIndex)
		newStep := common.Step{}

		for fixtureIndex, fixture := range step.Fixtures {

			//fmt.Printf("\tfixture no %d\n", fixtureIndex)

			newFixture := common.Fixture{}
			newFixture.MasterDimmer = fixture.MasterDimmer

			// Add colors.
			for _ = range fixture.Colors {
				var actualShift int
				//fmt.Printf("\t\tcolor no %d\n", colorIndex)
				newColor := common.Color{}
				//newColor.R = NewFixturesValues[fixtureIndex].values[stepIndex]
				if fixtureIndex == 0 {
					actualShift = 0
				} else {
					actualShift = shift
				}
				// c := range NewFixturesValues[fixtureIndex].colors{

				// }
				shiftedValues := calculateShift(NewFixturesValues[fixtureIndex].colors, actualShift)
				newColor.R = shiftedValues[stepIndex].R
				newColor.G = shiftedValues[stepIndex].G
				newColor.B = shiftedValues[stepIndex].B
				newFixture.Colors = append(newFixture.Colors, newColor)

				// shiftedValues := calculateShift(NewFixturesValues[fixtureIndex].colors, actualShift)
				// for _, newShiftedColor := range shiftedValues {
				// 	newColor.R = newShiftedColor.R
				// 	newColor.G = newShiftedColor.G
				// 	newColor.B = newShiftedColor.B
				// 	newFixture.Colors = append(newFixture.Colors, newColor)
				// }

			}
			newStep.Fixtures = append(newStep.Fixtures, newFixture)
		}

		stepsOut = append(stepsOut, newStep)

	}

	//printSteps(stepsOut)
	return stepsOut
}

// calculateShift - Takes a array of colors and a number to shift by.
func calculateShift(colors []common.Color, shift int) []common.Color {

	out := make([]common.Color, len(colors))

	//fmt.Printf("\t\t\tshift is %d\n", shift)
	//fmt.Printf("\t\t\tin %+v\n", values)

	var counter int
	for x := 0; x < len(colors); x++ {
		y := x + shift
		if y > len(colors)-1 {
			y = 0 + counter
			counter++
		}
		//fmt.Printf("\t\t\tx is %d   y is %d\n", x, y)

		out[y].R = colors[x].R
		out[y].G = colors[x].G
		out[y].B = colors[x].B

	}

	//fmt.Printf("\t\t\tout %+v\n", out)
	//fmt.Println("----------")
	return out
}

func printSteps(steps []common.Step) {

	fmt.Println()
	for stepIndex, step := range steps {
		fmt.Printf("Step No:%d\n", stepIndex)
		for fixtureIndex, fixture := range step.Fixtures {
			fmt.Printf("\t\tFixture No:%d\n", fixtureIndex)
			for _, color := range fixture.Colors {
				fmt.Printf("\t\t\tColor   R:%d G:%d B:%d\n", color.R, color.G, color.B)
			}
		}
	}
}
