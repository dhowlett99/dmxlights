package sequence

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
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

const debug = false

type SequencesConfig struct {
	Sequences []SequenceConfig `yaml:"sequences"`
}

type SequenceConfig struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Type        string `yaml:"type"`
	Group       int    `yaml:"group"`
}

// LoadSequences loads sequence configuration information.
// Each sequence has a :-
//       name: sequence name,  a singe word.
//       description: free text describing the sequence.
//       group: assignes to one of the top 4 rows of the launchpad. 1-4
//       type:  rgb, scanner or switch
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

// Before a sequence can run it needs to be created.
// Assigns default values for all types of sequence.
func CreateSequence(
	sequenceType string,
	mySequenceNumber int,
	pattens map[string]common.Patten,
	fixturesConfig *fixture.Fixtures,
	channels common.Channels,
	selectedFloodMap map[int]bool) common.Sequence {

	var initialPatten string
	var scanners int // Number of scanners in this sequence

	// Populate the static colors for this sequence with the defaults.
	staticColorsButtons := setDefaultStaticColorButtons(mySequenceNumber)

	// Populate the edit sequence colors for this sequence with the defaults.
	sequenceColorButtons := setDefaultStaticColorButtons(mySequenceNumber)

	// Populate the set scanner gobo colors buttons for this sequence with the defaults.
	sequenceGoboColorButtons := setDefaultGoboColorButtons(mySequenceNumber)

	// Set default values
	if sequenceType == "rgb" {
		initialPatten = "standard"
	}

	// Initilaise Scanners's
	var gobos = []common.Gobo{}
	if sequenceType == "scanner" {
		initialPatten = "circle"

		// Initilaise Gobo's
		scanners, gobos = fixture.HowManyGobos(mySequenceNumber, fixturesConfig)
	}

	// A map of the state of fixtures in the sequence.
	// We can disable a fixture by setting fixtureDisabled to true.
	fixtureDisabled := make(map[int]bool, 8)

	// The actual sequence definition.
	sequence := common.Sequence{
		NumberFixtures:               8,
		NumberScanners:               scanners,
		Type:                         sequenceType,
		Hide:                         false,
		Mode:                         "Sequence",
		StaticColors:                 staticColorsButtons,
		AvailableSequenceColors:      sequenceColorButtons,
		AvailableGoboSelectionColors: sequenceGoboColorButtons,
		Name:                         sequenceType,
		Number:                       mySequenceNumber,
		FadeSpeed:                    12,
		FadeTime:                     75 * time.Millisecond,
		MusicTrigger:                 false,
		Run:                          true,
		Bounce:                       false,
		NumberSteps:                  8 * 14, // Eight lamps and 14 steps to fade up and down.
		Patten: common.Patten{
			Name:     initialPatten,
			Length:   2,
			Size:     2,
			Fixtures: 8,
			Steps:    pattens[initialPatten].Steps,
		},
		ScannerSize:  120,
		Speed:        14,
		CurrentSpeed: 25 * time.Millisecond,
		Colors: []common.Color{
			{
				R: 0,
				G: 0,
				B: 0,
			},
		},
		Shift:                 0, // Start at zero ie no shift.
		Blackout:              false,
		Master:                255,
		Gobo:                  gobos,
		SelectedGobo:          1,
		SelectedFloodSequence: selectedFloodMap,
		Flood:                 false,
		AutoColor:             false,
		AutoPatten:            false,
		SelectedScannerPatten: 0,
		FixtureDisabled:       fixtureDisabled,
		NumberCoordinates:     10,
	}

	// Make functions for each of the sequences.
	for function := 0; function <= 8; function++ {
		newFunction := common.Function{
			Name:           strconv.Itoa(function),
			SequenceNumber: mySequenceNumber,
			Number:         function,
			State:          false,
		}
		sequence.Functions = append(sequence.Functions, newFunction)
	}

	if sequenceType == "switch" {

		if debug {
			fmt.Printf("Load switch data\n")
		}

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

// Now the sequence has been created, this functions starts the sequence.
func PlaySequence(sequence common.Sequence,
	mySequenceNumber int,
	pad *mk3.Launchpad,
	eventsForLauchpad chan common.ALight,
	pattens map[string]common.Patten,
	dmxController *ft232.DMXController,
	fixturesConfig *fixture.Fixtures,
	channels common.Channels,
	soundTriggers []*common.Trigger) {

	stepDelay := 10 * time.Microsecond

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

		sequence.UpdateShift = false

		// Check for any waiting commands.
		sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence.CurrentSpeed*10, sequence, channels)

		// Sequence in Switch Mode.
		if sequence.PlaySwitchOnce && sequence.Type == "switch" {
			// Show initial state of switches
			showSwitches(mySequenceNumber, &sequence, eventsForLauchpad, dmxController, fixturesConfig)
			sequence.PlaySwitchOnce = false
			sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, 1*time.Microsecond, sequence, channels)
			continue
		}

		// Sequence in Static Mode.
		if sequence.PlayStaticOnce && sequence.Static {
			for myFixtureNumber, lamp := range sequence.StaticColors {
				if sequence.Hide {
					if lamp.Flash {
						onColor := common.ConvertRGBtoPalette(lamp.Color.R, lamp.Color.G, lamp.Color.B)
						launchpad.FlashLight(mySequenceNumber, myFixtureNumber, onColor, 0, eventsForLauchpad)
					} else {
						launchpad.LightLamp(mySequenceNumber, myFixtureNumber, lamp.Color.R, lamp.Color.G, lamp.Color.B, sequence.Master, eventsForLauchpad)
					}
				}
				fixture.MapFixtures(mySequenceNumber, dmxController, myFixtureNumber, lamp.Color.R, lamp.Color.G, lamp.Color.B, 0, 0, 0, 0, 0, fixturesConfig, sequence.Blackout, sequence.Master, sequence.Master)
			}
			// Only play once, we don't want to flood the DMX universe with
			// continual commands.
			sequence.PlayStaticOnce = false
			continue
		}

		// This is the inner loop where the sequence runs.
		// Sequence in Normal Running Mode.
		if sequence.Mode == "Sequence" {
			for sequence.Run && !sequence.Static {

				// Map music trigger function.
				sequence.MusicTrigger = sequence.Functions[common.Function10_Music_Trigger].State

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

				if sequence.Type != "scanner" {
					sequence.Steps = pattens[sequence.Patten.Name].Steps
				}

				// Setup rgb pattens.
				if sequence.Type == "rgb" {
					sequence.ChangePatten = false
					if sequence.SelectedPatten == 0 {
						sequence.Patten.Name = "standard"
					}
					if sequence.SelectedPatten == 1 {
						sequence.Patten.Name = "pairs"
					}
					if sequence.SelectedPatten == 2 {
						sequence.Patten.Name = "inward"
					}
					if sequence.SelectedPatten == 3 {
						sequence.Patten.Name = "colors"
					}
				}

				// Setup scanner pattens.
				if sequence.Type == "scanner" {
					sequence.ChangePatten = false

					sequence.Steps = setPattern(sequence)

					if sequence.AutoColor {
						sequence.SelectedGobo++
						if sequence.SelectedGobo > 7 {
							sequence.SelectedGobo = 0
						}
					}
				}

				// Check is any commands are waiting.
				sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence.CurrentSpeed, sequence, channels)
				if !sequence.Run || sequence.Flood || sequence.ChangePatten || sequence.UpdateShift {
					break
				}

				// Calulate positions for fixtures based on the steps in the patten.
				sequence.Positions, sequence.NumberSteps = calculatePositions(sequence.Steps, sequence.Bounce)

				// If we are setting the patten automatically for rgb fixtures.
				if sequence.AutoPatten && sequence.Type == "rgb" {
					for name, patten := range pattens {
						if patten.Number == sequence.SelectedPatten {
							sequence.Patten.Name = name
							if debug {
								fmt.Printf(">>>> I AM PATTEN %s\n", name)
							}
							break
						}
					}
					sequence.SelectedPatten++
					if sequence.SelectedPatten > len(pattens) {
						sequence.SelectedPatten = 0
					}
				}

				// If we are setting the patten automatically for scanner fixtures.
				if sequence.AutoPatten && sequence.Type == "scanner" {
					sequence.SelectedScannerPatten++
					if sequence.SelectedScannerPatten > 3 {
						sequence.SelectedScannerPatten = 0
					}
				}

				// If we are setting the current colors in a rgb sequence.
				if sequence.AutoColor && sequence.Type == "rgb" {
					// Find a new color.
					newColor := []common.Color{}
					newColor = append(newColor, sequence.AvailableSequenceColors[sequence.SelectedColor].Color)
					sequence.SequenceColors = newColor

					// Step through the available colors.
					sequence.SelectedColor++
					if sequence.SelectedColor > 7 {
						sequence.SelectedColor = 0
					}
					sequence.Positions = replaceColors(sequence.Positions, sequence.SequenceColors)
				}

				// If we are updating the color in a sequence.
				if sequence.UpdateSequenceColor {
					if sequence.RecoverSequenceColors {
						if sequence.SavedSequenceColors != nil {
							sequence.Positions = replaceColors(sequence.Positions, sequence.SavedSequenceColors)
							sequence.AutoColor = false
						}
					} else {
						sequence.Positions = replaceColors(sequence.Positions, sequence.SequenceColors)
						// Save the current color selection.
						if sequence.SaveColors {
							sequence.SavedSequenceColors = common.HowManyColors(sequence.Positions)
							sequence.SaveColors = false
						}
					}
				}

				// Now that the patten colors have been decided and the positions calculated, set the CurrentSequenceColors
				// with the colors from that patten.
				sequence.CurrentSequenceColors = common.HowManyColors(sequence.Positions)

				// Run the sequence through.
				for step := 0; step < sequence.NumberSteps; step++ {

					//fmt.Printf("----STEP>>> %+v\n", step)

					// This is were we set the speed of the sequence to current speed.
					sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence.CurrentSpeed/2, sequence, channels)
					if !sequence.Run || sequence.Flood || sequence.ChangePatten || sequence.UpdateShift {
						break
					}

					// Prepare a message to be sent to the fixtures in the sequence.
					cmd := common.FixtureCommand{
						SequenceNumber:  sequence.Number,
						Inverted:        sequence.Inverted,
						Master:          sequence.Master,
						Hide:            sequence.Hide,
						Tick:            true,
						Positions:       sequence.Positions,
						Type:            sequence.Type,
						FadeSpeed:       sequence.FadeSpeed,
						FadeTime:        sequence.FadeTime,
						Size:            sequence.Size,
						Steps:           sequence.NumberSteps,
						CurrentSpeed:    sequence.CurrentSpeed,
						Speed:           sequence.Speed,
						Blackout:        sequence.Blackout,
						Flood:           sequence.Flood,
						CurrentPosition: step,
						SelectedGobo:    sequence.SelectedGobo,
						FixtureDisabled: sequence.FixtureDisabled,
						ScannerChase:    sequence.ScannerChase,
						ScannerColor:    sequence.ScannerColor,
					}

					// Now tell all the fixtures what they need to do.
					fixtureChannel1 <- cmd
					sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, stepDelay, sequence, channels)
					if !sequence.Run || sequence.Flood || sequence.ChangePatten || sequence.UpdateShift {
						break
					}
					fixtureChannel2 <- cmd
					sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, stepDelay, sequence, channels)
					if !sequence.Run || sequence.Flood || sequence.ChangePatten || sequence.UpdateShift {
						break
					}
					fixtureChannel3 <- cmd
					sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, stepDelay, sequence, channels)
					if !sequence.Run || sequence.Flood || sequence.ChangePatten || sequence.UpdateShift {
						break
					}
					fixtureChannel4 <- cmd
					sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, stepDelay, sequence, channels)
					if !sequence.Run || sequence.Flood || sequence.ChangePatten || sequence.UpdateShift {
						break
					}
					fixtureChannel5 <- cmd
					sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, stepDelay, sequence, channels)
					if !sequence.Run || sequence.Flood || sequence.ChangePatten || sequence.UpdateShift {
						break
					}
					fixtureChannel6 <- cmd
					sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, stepDelay, sequence, channels)
					if !sequence.Run || sequence.Flood || sequence.ChangePatten || sequence.UpdateShift {
						break
					}
					fixtureChannel7 <- cmd
					sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, stepDelay, sequence, channels)
					if !sequence.Run || sequence.Flood || sequence.ChangePatten || sequence.UpdateShift {
						break
					}
					fixtureChannel8 <- cmd
					sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, stepDelay, sequence, channels)
					if !sequence.Run || sequence.Flood || sequence.ChangePatten || sequence.UpdateShift {
						break
					}
				}
			}
		}
	}
}

// showSwitches - This is for switch sequences, a type of sequence which is just a set of eight switches.
// Each switch can have a number of states as defined in the fixtures.yaml file.
// The color of the lamp indicates which state you are in.
func showSwitches(mySequenceNumber int, sequence *common.Sequence, eventsForLauchpad chan common.ALight, dmxController *ft232.DMXController, fixtures *fixture.Fixtures) (flood bool) {

	for switchNumber, switchData := range sequence.Switches {
		for stateNumber, state := range switchData.States {

			// For this state.
			if stateNumber == switchData.CurrentState {
				// Use the button color for this state to light the correct color on the launchpad.
				launchpad.LightLamp(mySequenceNumber, switchNumber, state.ButtonColor.R, state.ButtonColor.G, state.ButtonColor.B, 255, eventsForLauchpad)

				// Now play all the values for this state.
				fixture.MapSwitchFixture(mySequenceNumber, dmxController, switchNumber, switchData.CurrentState, fixtures, sequence.Blackout, sequence.Master, sequence.Master)
			}
		}
	}

	return flood
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

	// Bounce repeates the steps in the sequence but backwards.
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

// replaceColors can take a sequence and replace its current patten colors with the colors specified.
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

// Sets the gobo select colors to default values. Namely Yellow !
func setDefaultGoboColorButtons(selectedSequence int) []common.StaticColorButton {

	// Make an array to hold gobo button colors.
	staticColorsButtons := []common.StaticColorButton{}

	for X := 0; X < 8; X++ {
		staticColorButton := common.StaticColorButton{}
		staticColorButton.X = X
		staticColorButton.Y = selectedSequence
		staticColorButton.SelectedColor = X
		staticColorButton.Color = common.Color{R: 255, G: 255, B: 0}
		staticColorsButtons = append(staticColorsButtons, staticColorButton)
	}

	return staticColorsButtons
}

// invertColor just reverses the DMX values.
func invertColor(color common.Color) (out common.Color) {
	out.R = reverseDmx(color.R)
	out.G = reverseDmx(color.G)
	out.B = reverseDmx(color.B)

	return out
}

// Takes a DMX value 1-255 and reverses the value.
func reverseDmx(n int) int {
	in := make(map[int]int, 255)
	var y = 255

	for x := 0; x <= 255; x++ {

		in[x] = y
		y--
	}
	return in[n]
}

// Flood - We are being asked to be in flood mode.
func Flood(sequence *common.Sequence, dmxController *ft232.DMXController, eventsForLauchpad chan common.ALight, fixturesConfig *fixture.Fixtures, enabled bool) {
	if sequence.Flood && sequence.PlayFloodOnce {
		for myFixtureNumber := 0; myFixtureNumber < sequence.NumberFixtures; myFixtureNumber++ {
			for s := range sequence.SelectedFloodSequence {
				if !sequence.Hide {
					launchpad.LightLamp(s, myFixtureNumber, 255, 255, 255, sequence.Master, eventsForLauchpad)
				}
				fixture.MapFixtures(s, dmxController, myFixtureNumber, 255, 255, 255, 0, 0, 0, 0, 0, fixturesConfig, sequence.Blackout, sequence.Master, sequence.Master)
			}
		}
		sequence.PlayFloodOnce = false
	}

	if !sequence.Flood && sequence.PlayFloodOnce {
		for myFixtureNumber := 0; myFixtureNumber < sequence.NumberFixtures; myFixtureNumber++ {
			for s := range sequence.SelectedFloodSequence {
				if !sequence.Hide {
					launchpad.LightLamp(s, myFixtureNumber, 0, 0, 0, sequence.Master, eventsForLauchpad)
				}
				fixture.MapFixtures(s, dmxController, myFixtureNumber, 0, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, sequence.Blackout, sequence.Master, sequence.Master)
			}
		}
		sequence.PlayFloodOnce = false
	}
}

func setPattern(sequence common.Sequence) (steps []common.Step) {
	if sequence.SelectedScannerPatten == 0 {
		coordinates := patten.CircleGenerator(sequence.ScannerSize, sequence.NumberCoordinates)
		scannerPatten := patten.GeneratePatten(coordinates, sequence.NumberScanners, sequence.Shift, sequence.ScannerChase)
		steps = scannerPatten.Steps
		return steps
	}
	if sequence.SelectedScannerPatten == 1 {
		coordinates := patten.ScanGeneratorLeftRight(128, sequence.NumberCoordinates)
		scannerPatten := patten.GeneratePatten(coordinates, sequence.NumberScanners, sequence.Shift, sequence.ScannerChase)
		steps = scannerPatten.Steps
		return steps
	}
	if sequence.SelectedScannerPatten == 2 {
		coordinates := patten.ScanGeneratorUpDown(128, sequence.NumberCoordinates)
		scannerPatten := patten.GeneratePatten(coordinates, sequence.NumberScanners, sequence.Shift, sequence.ScannerChase)
		steps = scannerPatten.Steps
		return steps
	}
	if sequence.SelectedScannerPatten == 3 {
		coordinates := patten.ScanGenerateSineWave(255, 5000, sequence.NumberCoordinates)
		scannerPatten := patten.GeneratePatten(coordinates, sequence.NumberScanners, sequence.Shift, sequence.ScannerChase)
		steps = scannerPatten.Steps
		return steps
	}

	return nil
}
