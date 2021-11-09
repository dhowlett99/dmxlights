package sequence

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/dhowlett99/dmxlights/pkg/commands"
	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
	"github.com/dhowlett99/dmxlights/pkg/launchpad"
	"github.com/dhowlett99/dmxlights/pkg/patten"
	"github.com/go-yaml/yaml"
	"github.com/oliread/usbdmx/ft232"
	"github.com/rakyll/launchpad/mk3"
)

type SequencesConfig struct {
	Sequences []SequenceConfig `yaml:"sequences"`
}

type SequenceConfig struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Type        string `yaml:"type"`
	Group       int    `yaml:"group"`
}

func LoadSequences() (sequences *SequencesConfig, err error) {
	filename := "sequences.yaml"

	_, err = os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		return nil, errors.New("error: loading sequences.yaml file: " + err.Error())
	}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.New("error: reading sequences.yaml file: " + err.Error())
	}

	sequences = &SequencesConfig{}
	err = yaml.Unmarshal(data, sequences)
	if err != nil {
		return nil, errors.New("error: unmarshalling sequences config: " + err.Error())
	}
	return sequences, nil
}

func CreateSequence(
	sequenceType string,
	mySequenceNumber int,
	pattens map[string]common.Patten,
	fixturesConfig *fixture.Fixtures,
	channels common.Channels) common.Sequence {

	var initialPatten string

	// Populate the static colors for this sequence with the defaults.
	staticColorsButtons := setDefaultStaticColorButtons(mySequenceNumber)

	// Populate the edit sequence colors for this sequence with the defaults.
	sequenceColorButtons := setDefaultStaticColorButtons(mySequenceNumber)

	// Set default values
	if sequenceType == "rgb" {
		initialPatten = "standard"
	}

	var gobos = []common.Gobo{}
	if sequenceType == "scanner" {
		initialPatten = "circle"

		// Initilaise Gobo's
		for _, f := range fixturesConfig.Fixtures {
			// fmt.Printf("Found fixture: %s, group: %d, desc: %s\n", f.Name, f.Group, f.Description)
			if f.Type == "scanner" {
				gobos = fixture.HowManyGobos(fixturesConfig, f)
			}
		}
	}

	sequence := common.Sequence{
		Type:                    sequenceType,
		Hide:                    false,
		Mode:                    "Sequence",
		StaticColors:            staticColorsButtons,
		AvailableSequenceColors: sequenceColorButtons,
		Name:                    sequenceType,
		Number:                  mySequenceNumber,
		FadeSpeed:               9,
		FadeTime:                150 * time.Millisecond,
		MusicTrigger:            false,
		Run:                     true,
		Bounce:                  false,
		Steps:                   8 * 14, // Eight lamps and 14 steps to fade up and down.
		Patten: common.Patten{
			Name:     initialPatten,
			Length:   2,
			Size:     2,
			Fixtures: 8,
			Chase:    []int{1, 2, 3, 4, 5, 6, 7, 8},
			Steps:    pattens[initialPatten].Steps,
		},
		SequenceSize: 60,
		Speed:        14,
		CurrentSpeed: 25 * time.Millisecond,
		Colors: []common.Color{
			{
				R: 0,
				G: 0,
				B: 0,
			},
		},
		Shift:        2,
		Blackout:     false,
		Master:       255,
		Gobo:         gobos,
		SelectedGobo: 1,
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

	// We start in forward chase.
	sequence.Functions[common.Function1_Forward_Chase].State = true

	if sequenceType == "switch" {

		fmt.Printf("Load switch data\n")
		// Load the switch information in from the fixtures.yaml file.
		// A new group of switches.
		newSwitchList := []common.Switch{}
		for _, fixture := range fixturesConfig.Fixtures {
			if fixture.Group == mySequenceNumber+1 {
				// find switch data.
				for _, swiTch := range fixture.Switches {
					newSwitch := common.Switch{}
					newSwitch.Name = swiTch.Name
					newSwitch.Number = swiTch.Number
					newSwitch.Description = swiTch.Description

					newSwitch.States = []common.State{}
					for _, state := range swiTch.States {
						newState := common.State{}
						newState.Name = state.Name
						newState.ButtonColor.R = state.ButtonColor.R
						newState.ButtonColor.G = state.ButtonColor.G
						newState.ButtonColor.B = state.ButtonColor.B

						newState.Values = []common.Value{}
						for _, value := range state.Values {
							newValue := common.Value{}
							newValue.Channel = value.Channel
							newValue.Setting = value.Setting
							newState.Values = append(newState.Values, newValue)
						}

						newSwitch.States = append(newSwitch.States, newState)
					}
					// Add new switch to the list.
					newSwitchList = append(newSwitchList, newSwitch)
				}
			}
		}
		sequence.Type = sequenceType
		sequence.Switches = newSwitchList
		sequence.PlaySwitchOnce = true
	}

	return sequence
}

func PlayNewSequence(sequence common.Sequence,
	mySequenceNumber int,
	pad *mk3.Launchpad,
	eventsForLauchpad chan common.ALight,
	pattens map[string]common.Patten,
	dmxController *ft232.DMXController,
	fixturesConfig *fixture.Fixtures,
	channels common.Channels,
	soundTriggers []*common.Trigger) {

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
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 0, fixtureChannels, eventsForLauchpad, dmxController, fixturesConfig)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 1, fixtureChannels, eventsForLauchpad, dmxController, fixturesConfig)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 2, fixtureChannels, eventsForLauchpad, dmxController, fixturesConfig)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 3, fixtureChannels, eventsForLauchpad, dmxController, fixturesConfig)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 4, fixtureChannels, eventsForLauchpad, dmxController, fixturesConfig)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 5, fixtureChannels, eventsForLauchpad, dmxController, fixturesConfig)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 6, fixtureChannels, eventsForLauchpad, dmxController, fixturesConfig)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 7, fixtureChannels, eventsForLauchpad, dmxController, fixturesConfig)

	// So this is the outer loop where sequence waits for commands and processes them if we're not playing a sequence.
	// i.e the sequence is in STOP mode and this is the way we change the RUN flag to START a sequence again.
	for {

		sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence.CurrentSpeed*10, sequence, channels)

		// Sequence in Switch Mode.
		if sequence.PlaySwitchOnce && sequence.Type == "switch" {
			// Show initial state of switches
			showSwitches(mySequenceNumber, sequence.Switches, eventsForLauchpad, dmxController, fixturesConfig, sequence.Blackout, sequence.Master)
			sequence.PlaySwitchOnce = false
			continue
		}

		// Sequence in Static Mode.
		if sequence.PlayStaticOnce && sequence.Static && sequence.Mode == "Static" {
			for myFixtureNumber, lamp := range sequence.StaticColors {
				if !sequence.Hide {
					if lamp.Flash {
						onColor := common.ConvertRGBtoPalette(lamp.Color.R, lamp.Color.G, lamp.Color.B)
						launchpad.FlashLight(mySequenceNumber, myFixtureNumber, onColor, 0, eventsForLauchpad)
					} else {
						launchpad.LightLamp(mySequenceNumber, myFixtureNumber, lamp.Color.R, lamp.Color.G, lamp.Color.B, sequence.Master, eventsForLauchpad)
					}
				}
				fixture.MapFixtures(mySequenceNumber, dmxController, myFixtureNumber, lamp.Color.R, lamp.Color.G, lamp.Color.B, 0, 0, 0, 0, fixturesConfig, sequence.Blackout, sequence.Master, sequence.Master)
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
						sequence.CurrentSpeed = commands.SetSpeed(sequence.Speed)
					}
				}

				steps := pattens[sequence.Patten.Name].Steps

				if sequence.Patten.Name == "circle" {
					coordinates := patten.CircleGenerator(sequence.SequenceSize)
					scannerPatten := patten.GeneratePatten(coordinates)
					steps = scannerPatten.Steps
				}
				if sequence.Patten.Name == "leftandright" {
					coordinates := patten.ScanGeneratorLeftRight(128)
					scannerPatten := patten.GeneratePatten(coordinates)
					steps = scannerPatten.Steps
				}
				if sequence.Patten.Name == "upanddown" {
					coordinates := patten.ScanGeneratorUpDown(128)
					scannerPatten := patten.GeneratePatten(coordinates)
					steps = scannerPatten.Steps
				}
				if sequence.Patten.Name == "sinewave" {
					coordinates := patten.ScanGenerateSineWave(255, 5000)
					scannerPatten := patten.GeneratePatten(coordinates)
					steps = scannerPatten.Steps
				}

				sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence.CurrentSpeed, sequence, channels)
				if !sequence.Run {
					break
				}

				// Calulate positions for fixtures based on patten.
				sequence.Positions, sequence.Steps = calculatePositions(steps, sequence.Bounce)
				if sequence.UpdateSequenceColor {
					sequence.Positions = replaceColors(sequence.Positions, sequence.SequenceColors)
				}

				// Set the current colors in the sequence.
				sequence.CurrentSequenceColors = common.HowManyColors(sequence.Positions)

				// Run the sequence through.
				for step := 0; step < sequence.Steps; step++ {

					// This is were we set the speed of the sequence to current speed.
					sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence.CurrentSpeed/2, sequence, channels)

					cmd := common.FixtureCommand{
						Inverted:        sequence.Inverted,
						Master:          sequence.Master,
						Hide:            sequence.Hide,
						Tick:            true,
						Positions:       sequence.Positions,
						Type:            sequence.Type,
						FadeSpeed:       sequence.FadeSpeed,
						FadeTime:        sequence.FadeTime,
						Size:            sequence.Size,
						Steps:           sequence.Steps,
						CurrentSpeed:    sequence.CurrentSpeed,
						Speed:           sequence.Speed,
						Blackout:        sequence.Blackout,
						CurrentPosition: step,
						SelectedGobo:    sequence.SelectedGobo,
					}
					fixtureChannel1 <- cmd
					if sequence.Type == "scanner" {
						sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, 1*time.Microsecond, sequence, channels)
					} else {
						sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, 100*time.Microsecond, sequence, channels)
					}
					if !sequence.Run {
						break
					}

					if sequence.Type == "scanner" {
						continue
					}
					fixtureChannel2 <- cmd
					sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, 100*time.Microsecond, sequence, channels)
					if !sequence.Run {
						break
					}
					fixtureChannel3 <- cmd
					sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, 100*time.Microsecond, sequence, channels)
					if !sequence.Run {
						break
					}
					fixtureChannel4 <- cmd
					sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, 100*time.Microsecond, sequence, channels)
					if !sequence.Run {
						break
					}
					fixtureChannel5 <- cmd
					sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, 100*time.Microsecond, sequence, channels)
					if !sequence.Run {
						break
					}
					fixtureChannel6 <- cmd
					sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, 100*time.Microsecond, sequence, channels)
					if !sequence.Run {
						break
					}
					fixtureChannel7 <- cmd
					sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, 100*time.Microsecond, sequence, channels)
					if !sequence.Run {
						break
					}
					fixtureChannel8 <- cmd
					sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, 100*time.Microsecond, sequence, channels)
					if !sequence.Run {
						break
					}
				}
			}
		}
	}
}

func showSwitches(mySequenceNumber int, switches []common.Switch, eventsForLauchpad chan common.ALight, dmxController *ft232.DMXController, fixtures *fixture.Fixtures, blackout bool, master int) {
	for switchNumber, switchData := range switches {
		for valuePosition, value := range switchData.States {
			if valuePosition == switchData.CurrentPosition {
				launchpad.LightLamp(mySequenceNumber, switchNumber, value.ButtonColor.R, value.ButtonColor.G, value.ButtonColor.B, 255, eventsForLauchpad)
				fixture.MapSwitchFixture(mySequenceNumber, dmxController, switchNumber, switchData.CurrentPosition, fixtures, blackout, master, master)
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
	var waitForColors bool
	for _, step := range steps {
		for fixtureIndex, fixture := range step.Fixtures {
			noColors := len(fixture.Colors)
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
					if noColors > 1 {
						if fixture.Type != "scanner" {
							counter = counter + 14
							waitForColors = true
						}
					}
				}
			}
		}
		if !waitForColors {
			counter = counter + 14
		}
	}

	if bounce {
		for index := len(steps) - 1; index >= 0; index-- {
			step := steps[index]
			for fixtureIndex, fixture := range step.Fixtures {
				noColors := len(fixture.Colors)
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
						if noColors > 1 {
							if fixture.Type != "scanner" {
								counter = counter + 14
								waitForColors = true
							}
						}
					}
				}
				if step.Type == "scanner" {
					counter = counter + 14
				}
			}
			if !waitForColors {
				counter = counter + 14
			}
		}
	}

	return positionsOut, counter
}

func replaceColors(positionsMap map[int][]common.Position, colors []common.Color) map[int][]common.Position {

	var insetColor int
	numberColors := len(colors)

	replace := make(map[int][]common.Position)
	for currentPosition, positions := range positionsMap {
		for _, position := range positions {
			if insetColor >= numberColors {
				insetColor = 0
			}
			position.Color = colors[insetColor]
			insetColor++

			replace[currentPosition] = append(replace[currentPosition], position)
		}
	}

	return replace
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

func invertColor(color common.Color) (out common.Color) {

	out.R = reverse_dmx(color.R)
	out.G = reverse_dmx(color.G)
	out.B = reverse_dmx(color.B)

	return out

}

func reverse_dmx(n int) int {

	in := make(map[int]int, 255)
	var y = 255

	for x := 0; x <= 255; x++ {

		in[x] = y
		y--
	}
	return in[n]
}
