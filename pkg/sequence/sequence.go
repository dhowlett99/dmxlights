package sequence

import (
	"fmt"
	"time"

	"github.com/dhowlett99/dmxlights/pkg/commands"
	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
	"github.com/dhowlett99/dmxlights/pkg/launchpad"
	"github.com/oliread/usbdmx/ft232"
	"github.com/rakyll/launchpad/mk3"
)

func CreateSequence(
	sequenceType string,
	mySequenceNumber int,
	pattens map[string]common.Patten,
	fixtureConfig *fixture.Fixtures,
	channels common.Channels) common.Sequence {

	var initialPatten string

	// Populate the static colors for this sequence with the defaults.
	staticColorsButtons := setDefaultStaticColorButtons(mySequenceNumber)

	// Set default values
	if sequenceType == "rgb" {
		initialPatten = "standard"
	}
	if sequenceType == "scanner" {
		initialPatten = "scanner"
	}

	sequence := common.Sequence{
		Type:         sequenceType,
		Hide:         false,
		Mode:         "Sequence",
		StaticColors: staticColorsButtons,
		Name:         sequenceType,
		Number:       mySequenceNumber,
		FadeSpeed:    9,
		FadeTime:     150 * time.Millisecond,
		MusicTrigger: false,
		Run:          true,
		Bounce:       false,
		Steps:        8 * 14, // Eight lamps and 14 steps to fade up and down.
		Patten: common.Patten{
			Name:     initialPatten,
			Length:   2,
			Size:     2,
			Fixtures: 8,
			Chase:    []int{1, 2, 3, 4, 5, 6, 7, 8},
			Steps:    pattens[initialPatten].Steps,
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
			SequenceNumber: mySequenceNumber,
			Number:         function,
			State:          false,
		}
		sequence.Functions = append(sequence.Functions, newFunction)
	}

	if sequenceType == "switch" {

		fmt.Printf("Load switch data\n")
		// Load the switch information in from the fixtures.yaml file.
		// A new group of switches.
		newSwitchList := []common.Switch{}
		for _, fixture := range fixtureConfig.Fixtures {
			if fixture.Group == mySequenceNumber+1 {
				// find switch data.
				for _, swiTch := range fixture.Switches {
					newSwitch := common.Switch{}
					newSwitch.Name = swiTch.Name
					newSwitch.Number = swiTch.Number
					newSwitch.Description = swiTch.Description
					for _, value := range swiTch.Values {
						newValue := common.Value{}
						newValue.Name = value.Name
						newValue.Value = value.Value
						newValue.ButtonColor.R = value.ButtonColor.R
						newValue.ButtonColor.G = value.ButtonColor.G
						newValue.ButtonColor.B = value.ButtonColor.B
						newValue.Channel = value.Channel
						newSwitch.Values = append(newSwitch.Values, newValue)
					}
					// Add new switch to the list.
					newSwitchList = append(newSwitchList, newSwitch)
				}
			}
		}
		sequence.Type = sequenceType
		sequence.Switches = newSwitchList
	}

	return sequence
}

func PlayNewSequence(sequence common.Sequence,
	mySequenceNumber int,
	pad *mk3.Launchpad,
	eventsForLauchpad chan common.ALight,
	pattens map[string]common.Patten,
	dmxController *ft232.DMXController,
	fixtureConfig *fixture.Fixtures,
	channels common.Channels,
	soundTriggers []*common.Trigger) {

	var positions map[int][]common.Position

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

		if sequence.Type == "switch" {
			// Show initial state of switches
			showSwitches(mySequenceNumber, sequence.Switches, eventsForLauchpad)
			continue
		}

		// Map bounce function to sequence bounce setting.
		sequence.Bounce = sequence.Functions[common.Function7_Bounce].State

		// Map music trigger function.
		sequence.MusicTrigger = sequence.Functions[common.Function8_Music_Trigger].State
		if sequence.MusicTrigger {
			sequence.Run = true
		}

		// Map static function.
		if sequence.Functions[common.Function6_Static].State {
			sequence.Static = true
		}

		// We are in static color editing mode flash this rows buttons.
		// if sequence.EditColors && sequence.Static {
		// 	showEditColorButtons(mySequenceNumber, eventsForLauchpad, true)
		// }
		// if !sequence.EditColors && !sequence.FunctionMode {
		// 	showEditColorButtons(mySequenceNumber, eventsForLauchpad, false)
		// }

		// Sequence in Static Mode.
		if sequence.PlayStaticOnce && sequence.Static && sequence.Mode == "Static" {
			for myFixtureNumber, lamp := range sequence.StaticColors {
				if !sequence.Hide {
					if lamp.Flash {
						fmt.Printf("FlashLight X:%d Y:%d\n", lamp.X, lamp.Y)
						onColor := common.ConvertRGBtoPalette(lamp.Color.R, lamp.Color.G, lamp.Color.B)
						launchpad.FlashLight(mySequenceNumber, myFixtureNumber, onColor, 0, eventsForLauchpad)
					} else {
						fmt.Printf("LightLamp X:%d Y:%d\n", lamp.X, lamp.Y)
						launchpad.LightLamp(mySequenceNumber, myFixtureNumber, lamp.Color.R, lamp.Color.G, lamp.Color.B, sequence.Master, eventsForLauchpad)
					}
				}
				fixture.MapFixtures(mySequenceNumber, dmxController, myFixtureNumber, lamp.Color.R, lamp.Color.G, lamp.Color.B, 0, 0, 0, 0, fixtureConfig, sequence.Blackout, sequence.Master, sequence.Master)
			}
			// Only play once, we don't want to flood the DMX universe with
			// continual commands.
			sequence.PlayStaticOnce = false
			continue
		}

		// This is the inner loop where the sequence runs.
		// Sequence in Normal Running Mode.
		if sequence.Mode == "Sequence" {
			for sequence.Run {

				sequence.Functions[common.Function1_Forward_Chase].State = sequence.Run

				// Map music trigger function.
				sequence.MusicTrigger = sequence.Functions[common.Function8_Music_Trigger].State

				// If the music trigger is being used then the timer is disabled.
				for _, trigger := range soundTriggers {
					if sequence.MusicTrigger {
						sequence.CurrentSpeed = time.Duration(12 * time.Hour)
						if trigger.SequenceNumber == mySequenceNumber {
							trigger.State = true
						}
					} else {
						if trigger.SequenceNumber == mySequenceNumber {
							trigger.State = false
						}
						sequence.CurrentSpeed = 25 * time.Millisecond
					}
				}

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
						Master:          sequence.Master,
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
}

func showSwitches(mySequenceNumber int, switches []common.Switch, eventsForLauchpad chan common.ALight) {

	for myFixtureNumber, swiTch := range switches {
		// fmt.Printf("swiTch: name %s\n", swiTch.Name)
		// fmt.Printf("swiTch: no %d\n", swiTch.Number)
		// fmt.Printf("swiTch: description %s\n", swiTch.Description)
		// fmt.Printf("swiTch: values %+v\n", swiTch.Values)
		for _, value := range swiTch.Values {
			if value.Name == "Off" {
				launchpad.LightLamp(mySequenceNumber, myFixtureNumber, value.ButtonColor.R, value.ButtonColor.G, value.ButtonColor.B, 255, eventsForLauchpad)
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

// Sets the static colors to default values.
func setDefaultStaticColorButtons(selectedSequence int) []common.StaticColorButton {

	// Make an array to hold static colors.
	staticColorsButtons := []common.StaticColorButton{}

	for X := 0; X < 8; X++ {
		staticColorButton := common.StaticColorButton{}
		staticColorButton.X = X
		staticColorButton.Y = selectedSequence
		staticColorButton.SelectedColor = X
		staticColorButton.Color = common.GetColorButtonsArray(X)
		staticColorsButtons = append(staticColorsButtons, staticColorButton)
	}

	return staticColorsButtons
}
