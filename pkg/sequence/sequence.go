package sequence

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/dhowlett99/dmxlights/pkg/commands"
	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
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
	Label       string `yaml:"label"`
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
	availableRGBPattens map[int]common.Patten,
	fixturesConfig *fixture.Fixtures,
	channels common.Channels) common.Sequence {

	// Number of scanners in this sequence
	var scanners int

	// Populate the static colors for this sequence with the defaults.
	staticColorsButtons := setDefaultStaticColorButtons(mySequenceNumber)

	// Populate the edit sequence colors for this sequence with the defaults.
	sequenceColorButtons := setDefaultStaticColorButtons(mySequenceNumber)

	// Every scanner has a number of colors in its wheel.
	availableScannerColors := make(map[int][]common.StaticColorButton)

	// Find the fixtures.
	availableFixtures := setAvalableFixtures(fixturesConfig)

	fixtureLabels := []string{}
	for _, fixture := range fixturesConfig.Fixtures {
		if fixture.Type == "scanner" {
			fixtureLabels = append(fixtureLabels, fixture.Label)
		}
	}

	// Every scanner has a number of gobos in its wheel.
	availableScannerGobos := make(map[int][]common.StaticColorButton)

	if sequenceType == "scanner" {
		//initialPatten = "circle"

		// Initilaise Gobo's
		// Every scanner has a number of gobos in its wheel.
		scanners, availableScannerGobos = getAvailableScannerGobos(mySequenceNumber, fixturesConfig)

		// Get available scanner colors for all fixtures.
		availableScannerColors = getAvailableScannerColors(fixturesConfig)

	}

	// A map of the state of fixtures in the sequence.
	// We can disable a fixture by setting fixtureDisabled to true.
	fixtureDisabled := make(map[int]bool, 8)
	disabledOnce := make(map[int]bool, 8)

	// A map of the fixture colors.
	scannerColors := make(map[int]int)

	// The actual sequence definition.
	sequence := common.Sequence{
		ScannerAvailableColors: availableScannerColors,
		ScannersAvailable:      availableFixtures,
		ScannersTotal:          scanners,
		Type:                   sequenceType,
		Hide:                   false,
		Mode:                   "Sequence",
		StaticColors:           staticColorsButtons,
		RGBAvailableColors:     sequenceColorButtons,
		ScannerAvailableGobos:  availableScannerGobos,
		Name:                   sequenceType,
		Number:                 mySequenceNumber,
		FadeSpeed:              12,
		FadeTime:               75 * time.Millisecond,
		MusicTrigger:           false,
		Run:                    true,
		Bounce:                 false,
		NumberSteps:            8 * 14, // Eight lamps and 14 steps to fade up and down.
		RGBAvailablePattens:    availableRGBPattens,
		ScannerSize:            common.DefaultScannerSize,
		Speed:                  14,
		CurrentSpeed:           25 * time.Millisecond,
		ScannerShift:           0, // Start at zero ie no shift.
		Blackout:               false,
		Master:                 255,
		ScannerGobo:            1,
		StartFlood:             false,
		RGBColor:               1,
		AutoColor:              false,
		AutoPatten:             false,
		ScannerPatten:          0,
		FixtureDisabled:        fixtureDisabled,
		DisableOnce:            disabledOnce,
		ScannerCoordinates:     []int{12, 16, 24, 32},
		ScannerColor:           scannerColors,
		ScannerOffsetPan:       120,
		ScannerOffsetTilt:      120,
		GuiFixtureLabels:       fixtureLabels,
	}

	// Since we will be accessing these maps from the sequence thread and the fixture threads
	// We need to protect the maps from syncronous access.
	sequence.FixtureDisabledMutex = &sync.RWMutex{}
	sequence.DisableOnceMutex = &sync.RWMutex{}

	if sequence.Type == "rgb" {
		sequence.GuiFunctionLabels[0] = "Set\nPatten"
		sequence.GuiFunctionLabels[1] = "Auto\nColor"
		sequence.GuiFunctionLabels[2] = "Auto\nPatten"
		sequence.GuiFunctionLabels[3] = "Bounce"
		sequence.GuiFunctionLabels[4] = "Chase\nColor"
		sequence.GuiFunctionLabels[5] = "Static\nColor"
		sequence.GuiFunctionLabels[6] = "Invert"
		sequence.GuiFunctionLabels[7] = "Music"
	}

	if sequence.Type == "scanner" {
		sequence.GuiFunctionLabels[0] = "Set\nPatten"
		sequence.GuiFunctionLabels[1] = "Auto\nColor"
		sequence.GuiFunctionLabels[2] = "Auto\nPatten"
		sequence.GuiFunctionLabels[3] = "Bounce"
		sequence.GuiFunctionLabels[4] = "Color"
		sequence.GuiFunctionLabels[5] = "Gobo"
		sequence.GuiFunctionLabels[6] = "Chase"
		sequence.GuiFunctionLabels[7] = "Music"
	}

	// Make functions for each of the sequences.
	for function := 0; function < 8; function++ {
		newFunction := common.Function{
			Name:           strconv.Itoa(function),
			SequenceNumber: mySequenceNumber,
			Number:         function,
			State:          false,
			Label:          sequence.GuiFunctionLabels[function],
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
					newSwitch.Label = swiTch.Label
					newSwitch.Number = swiTch.Number
					newSwitch.Description = swiTch.Description

					newSwitch.States = []common.State{}
					for _, state := range swiTch.States {
						newState := common.State{}
						newState.Name = state.Name
						newState.Label = state.Label
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
	guiButtons chan common.ALight,
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

	// Create channels used for stopping the fixture threads for this sequnece.
	fixtureStopChannels := []chan bool{}
	fixtureStopChannel0 := make(chan bool)
	fixtureStopChannels = append(fixtureStopChannels, fixtureStopChannel0)
	fixtureStopChannel1 := make(chan bool)
	fixtureStopChannels = append(fixtureStopChannels, fixtureStopChannel1)
	fixtureStopChannel2 := make(chan bool)
	fixtureStopChannels = append(fixtureStopChannels, fixtureStopChannel2)
	fixtureStopChannel3 := make(chan bool)
	fixtureStopChannels = append(fixtureStopChannels, fixtureStopChannel3)
	fixtureStopChannel4 := make(chan bool)
	fixtureStopChannels = append(fixtureStopChannels, fixtureStopChannel4)
	fixtureStopChannel5 := make(chan bool)
	fixtureStopChannels = append(fixtureStopChannels, fixtureStopChannel5)
	fixtureStopChannel6 := make(chan bool)
	fixtureStopChannels = append(fixtureStopChannels, fixtureStopChannel6)
	fixtureStopChannel7 := make(chan bool)
	fixtureStopChannels = append(fixtureStopChannels, fixtureStopChannel7)

	// Create eight fixture threads for this sequence.
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 0, fixtureChannels, eventsForLauchpad, guiButtons, dmxController, fixturesConfig, fixtureStopChannels)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 1, fixtureChannels, eventsForLauchpad, guiButtons, dmxController, fixturesConfig, fixtureStopChannels)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 2, fixtureChannels, eventsForLauchpad, guiButtons, dmxController, fixturesConfig, fixtureStopChannels)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 3, fixtureChannels, eventsForLauchpad, guiButtons, dmxController, fixturesConfig, fixtureStopChannels)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 4, fixtureChannels, eventsForLauchpad, guiButtons, dmxController, fixturesConfig, fixtureStopChannels)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 5, fixtureChannels, eventsForLauchpad, guiButtons, dmxController, fixturesConfig, fixtureStopChannels)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 6, fixtureChannels, eventsForLauchpad, guiButtons, dmxController, fixturesConfig, fixtureStopChannels)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 7, fixtureChannels, eventsForLauchpad, guiButtons, dmxController, fixturesConfig, fixtureStopChannels)

	// So this is the outer loop where sequence waits for commands and processes them if we're not playing a sequence.
	// i.e the sequence is in STOP mode and this is the way we change the RUN flag to START a sequence again.
	for {

		sequence.UpdateShift = false

		if !sequence.Run {
			sendStopToAllFixtures(sequence, fixtureStopChannels)
		}

		// Check for any waiting commands.
		sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence.CurrentSpeed*10, sequence, channels)

		// Sequence in Switch Mode.
		if sequence.PlaySwitchOnce && sequence.Type == "switch" {
			// Show initial state of switches
			ShowSwitches(mySequenceNumber, &sequence, eventsForLauchpad, guiButtons, dmxController, fixturesConfig)
			sequence.PlaySwitchOnce = false
			sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, 1*time.Microsecond, sequence, channels)
			continue
		}

		// Sequence in flood mode.
		if sequence.StartFlood && sequence.FloodPlayOnce {
			sequence.Run = false
			// Prepare a message to be sent to the fixtures in the sequence.
			command := common.FixtureCommand{
				Tick:           true,
				Type:           sequence.Type,
				SequenceNumber: sequence.Number,
				StartFlood:     sequence.StartFlood,
			}
			// Stop the fixture threads doing what the are doing.
			sendStopToAllFixtures(sequence, fixtureStopChannels)

			// Now tell all the fixtures what they need to do.
			sendToAllFixtures(sequence, fixtureChannels, channels, command)
			sequence.StartFlood = false
			sequence.FloodPlayOnce = false
			sequence.Run = false
			continue
		}

		// Stop flood mode.
		if sequence.StopFlood && sequence.FloodPlayOnce {
			// Prepare a message to be sent to the fixtures in the sequence.
			command := common.FixtureCommand{
				Tick:           true,
				Type:           sequence.Type,
				SequenceNumber: sequence.Number,
				StartFlood:     sequence.StartFlood,
				StopFlood:      sequence.StopFlood,
			}
			// Now tell all the fixtures what they need to do.
			sendToAllFixtures(sequence, fixtureChannels, channels, command)
			sequence.StopFlood = false
			sequence.FloodPlayOnce = false
			sequence.Run = true
			continue
		}

		// Sequence in Static Mode.
		if sequence.PlayStaticOnce && sequence.Static && !sequence.StartFlood {
			// Prepare a message to be sent to the fixtures in the sequence.
			command := common.FixtureCommand{
				Tick:           true,
				Type:           sequence.Type,
				SequenceNumber: sequence.Number,
				Static:         sequence.Static,
				StaticColors:   sequence.StaticColors,
				Hide:           sequence.Hide,
				Master:         sequence.Master,
				Blackout:       sequence.Blackout,
			}

			// Stop the fixture threads doing what the are doing.
			sendStopToAllFixtures(sequence, fixtureStopChannels)

			// Now tell all the fixtures what they need to do.
			sendToAllFixtures(sequence, fixtureChannels, channels, command)
			sequence.PlayStaticOnce = false
			sequence.Run = false
			continue
		}

		// This is the inner loop where the sequence runs.
		// Sequence in Normal Running Mode.
		if sequence.Mode == "Sequence" {
			for sequence.Run && !sequence.Static {

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

				// Setup rgb pattens.
				if sequence.Type == "rgb" {
					sequence.Steps = sequence.RGBAvailablePattens[sequence.SelectedRGBPatten].Steps
					sequence.UpdatePatten = false
				}

				// Setup scanner pattens.
				if sequence.Type == "scanner" {

					// Get available scanner pattens.
					sequence.ScannerAvailablePattens = getAvailableScannerPattens(sequence)
					sequence.UpdatePatten = false

					sequence.Patten = sequence.ScannerAvailablePattens[sequence.ScannerPatten]
					sequence.Steps = sequence.Patten.Steps

					if sequence.AutoColor {
						sequence.ScannerGobo++
						if sequence.ScannerGobo > 7 {
							sequence.ScannerGobo = 0
						}

						scannerColorSelectedForThisFixtureLastTime := 0
						// AvailableFixtures give the real number of configured scanners.
						for _, fixture := range sequence.ScannersAvailable {
							// First check that this fixture has some configured colors.
							colors, ok := sequence.ScannerAvailableColors[fixture.Number]
							if ok {
								// Found a scanner with some colors.
								totalColorForThisFixture := len(colors)

								sequence.ScannerColor[fixture.Number] = sequence.ScannerColor[fixture.Number] + 1
								if sequence.ScannerColor[fixture.Number] > scannerColorSelectedForThisFixtureLastTime {
									if sequence.ScannerColor[fixture.Number] > totalColorForThisFixture {
										sequence.ScannerColor[fixture.Number] = 1
									}
									scannerColorSelectedForThisFixtureLastTime = scannerColorSelectedForThisFixtureLastTime + 1
									continue
								}
							}
						}
					}
				}

				// Check is any commands are waiting.
				sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence.CurrentSpeed, sequence, channels)
				if !sequence.Run || sequence.StartFlood || sequence.Static || sequence.UpdatePatten || sequence.UpdateShift {
					// Tell the fixtures to stop.
					sendStopToAllFixtures(sequence, fixtureStopChannels)
					break
				}

				// Calulate positions for fixtures based on the steps in the patten.
				sequence.Positions, sequence.NumberSteps = calculatePositions(sequence.Steps, sequence.Bounce)

				// If we are setting the patten automatically for rgb fixtures.
				if sequence.AutoPatten && sequence.Type == "rgb" {
					for pattenNumber, patten := range sequence.RGBAvailablePattens {
						if patten.Number == sequence.SelectedRGBPatten {
							sequence.Patten.Number = pattenNumber
							if debug {
								fmt.Printf(">>>> I AM PATTEN %d\n", pattenNumber)
							}
							break
						}
					}
					sequence.SelectedRGBPatten++
					if sequence.SelectedRGBPatten > len(sequence.RGBAvailablePattens) {
						sequence.SelectedRGBPatten = 0
					}
				}

				// If we are setting the patten automatically for scanner fixtures.
				if sequence.AutoPatten && sequence.Type == "scanner" {
					sequence.ScannerPatten++
					if sequence.ScannerPatten > 3 {
						sequence.ScannerPatten = 0
					}
				}

				// If we are setting the current colors in a rgb sequence.
				if sequence.AutoColor && sequence.Type == "rgb" {
					// Find a new color.
					newColor := []common.Color{}
					newColor = append(newColor, sequence.RGBAvailableColors[sequence.RGBColor].Color)
					sequence.SequenceColors = newColor

					// Step through the available colors.
					sequence.RGBColor++
					if sequence.RGBColor > 7 {
						sequence.RGBColor = 0
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
				sequence.CurrentColors = common.HowManyColors(sequence.Positions)

				// Run through the steps in the sequence.
				// Remember every step contains infomation for all the fixtures in this group.
				for step := 0; step < sequence.NumberSteps; step++ {
					// This is were we set the speed of the sequence to current speed.
					sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence.CurrentSpeed/2, sequence, channels)
					if !sequence.Run || sequence.StartFlood || sequence.Static || sequence.UpdatePatten || sequence.UpdateShift {
						// Tell the fixtures to stop.
						sendStopToAllFixtures(sequence, fixtureStopChannels)
						break
					}

					// Prepare a message to be sent to the fixtures in the sequence.
					command := common.FixtureCommand{
						SequenceNumber:         sequence.Number,
						Inverted:               sequence.PattenInverted,
						Master:                 sequence.Master,
						Hide:                   sequence.Hide,
						Tick:                   true,
						Positions:              sequence.Positions,
						Type:                   sequence.Type,
						FadeSpeed:              sequence.FadeSpeed,
						FadeTime:               sequence.FadeTime,
						Size:                   sequence.Size,
						Steps:                  sequence.NumberSteps,
						CurrentSpeed:           sequence.CurrentSpeed,
						Speed:                  sequence.Speed,
						Blackout:               sequence.Blackout,
						StartFlood:             sequence.StartFlood,
						StopFlood:              sequence.StopFlood,
						CurrentPosition:        step,
						SelectedGobo:           sequence.ScannerGobo,
						FixtureDisabled:        sequence.FixtureDisabled,
						FixtureDisabledMutex:   sequence.FixtureDisabledMutex,
						DisableOnce:            sequence.DisableOnce,
						DisableOnceMutex:       sequence.DisableOnceMutex,
						ScannerChase:           sequence.ScannerChase,
						ScannerColor:           sequence.ScannerColor,
						AvailableScannerColors: sequence.ScannerAvailableColors,
						OffsetPan:              sequence.ScannerOffsetPan,
						OffsetTilt:             sequence.ScannerOffsetTilt,
						FixtureLabels:          sequence.GuiFixtureLabels,
					}

					// Now tell all the fixtures in this group what they need to do.
					sendToAllFixtures(sequence, fixtureChannels, channels, command)
				}
			}
		}
	}
}

// Send a command to all the fixtures.
func sendToAllFixtures(sequence common.Sequence, fixtureChannels []chan common.FixtureCommand, channels common.Channels, command common.FixtureCommand) {
	for _, fixture := range fixtureChannels {
		fixture <- command
	}
}

// Send a Stop command to all the fixtures.
func sendStopToAllFixtures(sequence common.Sequence, fixtureStopChannels []chan bool) {
	for _, fixture := range fixtureStopChannels {
		select {
		case fixture <- true:
			continue
		case <-time.After(5 * time.Millisecond):
		}
	}
}

// showSwitches - This is for switch sequences, a type of sequence which is just a set of eight switches.
// Each switch can have a number of states as defined in the fixtures.yaml file.
// The color of the lamp indicates which state you are in.
// ShowSwitches relies on you giving the sequence number of the switch sequnence.
func ShowSwitches(mySequenceNumber int, sequence *common.Sequence, eventsForLauchpad chan common.ALight,
	guiButtons chan common.ALight, dmxController *ft232.DMXController, fixtures *fixture.Fixtures) {

	if debug {
		fmt.Printf("ShowSwitches for sequence %d\n", mySequenceNumber)
	}
	for switchNumber, switchData := range sequence.Switches {
		for stateNumber, state := range switchData.States {

			// For this state.
			if stateNumber == switchData.CurrentState {
				// Use the button color for this state to light the correct color on the launchpad.
				common.LightLamp(common.ALight{X: switchNumber, Y: mySequenceNumber, Red: state.ButtonColor.R, Green: state.ButtonColor.G, Blue: state.ButtonColor.B, Brightness: 255}, eventsForLauchpad, guiButtons)

				// Label the switch.
				common.LabelButton(switchNumber, mySequenceNumber, switchData.Label+"\n"+state.Label, guiButtons)

				// Now play all the values for this state.
				fixture.MapSwitchFixture(mySequenceNumber, dmxController, switchNumber, switchData.CurrentState, fixtures, sequence.Blackout, sequence.Master, sequence.Master)
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

func setAvalableFixtures(fixturesConfig *fixture.Fixtures) []common.StaticColorButton {

	// You need to select a fixture before you can choose a color or gobo.
	// availableFixtures holds a set of red buttons, one for every available fixture.
	availableFixtures := []common.StaticColorButton{}
	for _, fixture := range fixturesConfig.Fixtures {
		if fixture.Type == "scanner" {
			newFixture := common.StaticColorButton{}
			newFixture.Name = fixture.Name
			newFixture.Label = fixture.Label
			newFixture.Number = fixture.Number
			newFixture.SelectedColor = 1 // Red
			newFixture.Color = common.Color{R: 255, G: 0, B: 0}
			availableFixtures = append(availableFixtures, newFixture)
		}
	}

	return availableFixtures
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

// getAvailableScannerColors looks through the fixtures list and finds scanners that
// have colors defined in their config. It then returns an array of these available colors.
func getAvailableScannerColors(fixtures *fixture.Fixtures) map[int][]common.StaticColorButton {
	availableScannerColors := make(map[int][]common.StaticColorButton)
	for _, fixture := range fixtures.Fixtures {
		for _, channel := range fixture.Channels {
			if strings.Contains(channel.Name, "Color") {
				for _, setting := range channel.Settings {
					newStaticColorButton := common.StaticColorButton{}
					newStaticColorButton.SelectedColor = setting.Number
					newStaticColorButton.Color = common.GetRGBColorByName(setting.Name)
					availableScannerColors[fixture.Number] = append(availableScannerColors[fixture.Number], newStaticColorButton)
				}
			}
		}
	}
	return availableScannerColors
}

func getAvailableScannerGobos(sequenceNumber int, fixtures *fixture.Fixtures) (int, map[int][]common.StaticColorButton) {
	if debug {
		fmt.Printf("getAvailableScannerGobos\n")
	}

	var numberScanners int
	gobos := make(map[int][]common.StaticColorButton)

	for _, f := range fixtures.Fixtures {
		if debug {
			fmt.Printf("Fixture Name:%s\n", f.Name)
		}
		if f.Type == "scanner" {

			if debug {
				fmt.Printf("Sequence: %d - Scanner Name: %s Description: %s\n", sequenceNumber, f.Name, f.Description)
			}
			numberScanners++
			for _, channel := range f.Channels {
				if channel.Name == "Gobo" {
					newGobo := common.StaticColorButton{}
					for _, setting := range channel.Settings {
						newGobo.Name = setting.Name
						newGobo.Label = setting.Label
						newGobo.Number = setting.Number
						newGobo.Setting = setting.Setting
						newGobo.Color = common.Color{R: 255, G: 255, B: 0} // Yellow.
						gobos[f.Number] = append(gobos[f.Number], newGobo)
						if debug {
							fmt.Printf("\tGobo: %s Setting: %d\n", setting.Name, setting.Setting)
						}
					}
				}
			}
		}
	}
	return numberScanners, gobos
}

// getAvailableScannerPattens generates scanner pattens and stores them in the sequence.
// Each scanner can then select which patten to use.
// All scanner pattens have the same number of steps defined by NumberCoordinates.
func getAvailableScannerPattens(sequence common.Sequence) map[int]common.Patten {

	scannerPattens := make(map[int]common.Patten)

	// Scanner circle patten 0
	coordinates := patten.CircleGenerator(sequence.ScannerSize, sequence.ScannerCoordinates[sequence.ScannerSelectedCoordinates], float64(sequence.ScannerOffsetPan), float64(sequence.ScannerOffsetTilt))
	circlePatten := patten.GeneratePatten(coordinates, sequence.ScannersTotal, sequence.ScannerShift, sequence.ScannerChase)
	circlePatten.Name = "circle"
	circlePatten.Number = 0
	circlePatten.Label = "Circle"
	scannerPattens[0] = circlePatten

	// Scanner left right patten 1
	coordinates = patten.ScanGeneratorLeftRight(128, sequence.ScannerCoordinates[sequence.ScannerSelectedCoordinates])
	leftRightPatten := patten.GeneratePatten(coordinates, sequence.ScannersTotal, sequence.ScannerShift, sequence.ScannerChase)
	leftRightPatten.Name = "leftright"
	leftRightPatten.Number = 1
	leftRightPatten.Label = "Left.Right"
	scannerPattens[1] = leftRightPatten

	// Scanner up down patten 2
	coordinates = patten.ScanGeneratorUpDown(128, sequence.ScannerCoordinates[sequence.ScannerSelectedCoordinates])
	upDownPatten := patten.GeneratePatten(coordinates, sequence.ScannersTotal, sequence.ScannerShift, sequence.ScannerChase)
	upDownPatten.Name = "updown"
	upDownPatten.Number = 2
	upDownPatten.Label = "Up.Down"
	scannerPattens[2] = upDownPatten

	// Scanner zig zag patten 3
	coordinates = patten.ScanGenerateSineWave(255, 5000, sequence.ScannerCoordinates[sequence.ScannerSelectedCoordinates])
	zigZagPatten := patten.GeneratePatten(coordinates, sequence.ScannersTotal, sequence.ScannerShift, sequence.ScannerChase)
	zigZagPatten.Name = "zigzag"
	zigZagPatten.Number = 3
	zigZagPatten.Label = "Zig.Zag"
	scannerPattens[3] = zigZagPatten

	return scannerPattens

}

func SequenceSelect(eventsForLauchpad chan common.ALight, guiButtons chan common.ALight, selectedSequence int) {
	// Turn off all sequence lights.
	for seq := 0; seq < 3; seq++ {
		common.LightLamp(common.ALight{X: 8, Y: seq, Brightness: 255, Red: 100, Green: 255, Blue: 255}, eventsForLauchpad, guiButtons)
	}
	// Now turn pink the selected sequence select light.
	common.LightLamp(common.ALight{X: 8, Y: selectedSequence, Brightness: 255, Red: 255, Green: 0, Blue: 255}, eventsForLauchpad, guiButtons)
}
