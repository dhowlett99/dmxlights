package sequence

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/dhowlett99/dmxlights/pkg/commands"
	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
	"github.com/dhowlett99/dmxlights/pkg/pattern"
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
	availableRGBPatterns map[int]common.Pattern,
	fixturesConfig *fixture.Fixtures,
	channels common.Channels) common.Sequence {

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

	// A map of the fixture colors.
	scannerColors := make(map[int]int)

	// Number of scanners in this sequence
	var scanners int

	if sequenceType == "scanner" {

		// Initilise Gobos
		scanners, availableScannerGobos = getAvailableScannerGobos(mySequenceNumber, fixturesConfig)

		// Initialise Colors.
		availableScannerColors, scannerColors = getAvailableScannerColors(fixturesConfig)

	}

	// A map of the state of fixtures in the sequence.
	// We can disable a fixture by setting fixture Enabled to false.
	scannerState := make(map[int]common.ScannerState, 8)
	for x := 0; x < 8; x++ {
		newScanner := common.ScannerState{}
		newScanner.Enabled = true
		newScanner.Inverted = false
		scannerState[x] = newScanner
	}

	disabledOnce := make(map[int]bool, 8)

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
		RGBFade:                common.DefaultRGBFade,
		MusicTrigger:           false,
		Run:                    true,
		Bounce:                 false,
		RGBAvailablePatterns:   availableRGBPatterns,
		ScannerSize:            common.DefaultScannerSize,
		RGBSize:                common.DefaultRGBSize,
		Speed:                  common.DefaultSpeed,
		ScannerShift:           common.DefaultScannerShift,
		RGBShift:               common.DefaultRGBShift,
		Blackout:               false,
		Master:                 255,
		ScannerGobo:            1,
		StartFlood:             false,
		RGBColor:               1,
		AutoColor:              false,
		AutoPattern:            false,
		SelectedPattern:        common.DefaultPattern,
		ScannerState:           scannerState,
		DisableOnce:            disabledOnce,
		ScannerCoordinates:     []int{12, 16, 24, 32},
		ScannerColor:           scannerColors,
		ScannerOffsetPan:       120,
		ScannerOffsetTilt:      120,
		GuiFixtureLabels:       fixtureLabels,
	}

	// Since we will be accessing these maps from the sequence thread and the fixture threads
	// We need to protect the maps from syncronous access.
	sequence.ScannerStateMutex = &sync.RWMutex{}
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

	// Find the number of fixtures for this sequence.
	sequence.NumberFixtures = getNumberOfFixtures(mySequenceNumber, fixturesConfig)
	sequence.ScannerPositions = make(map[int]map[int][]common.Position, sequence.NumberFixtures)
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

	// Create channels used for stepping the fixture threads for this sequnece.
	fixtureStepChannels := []chan common.FixtureCommand{}
	fixtureStepChannel0 := make(chan common.FixtureCommand)
	fixtureStepChannels = append(fixtureStepChannels, fixtureStepChannel0)
	fixtureStepChannel1 := make(chan common.FixtureCommand)
	fixtureStepChannels = append(fixtureStepChannels, fixtureStepChannel1)
	fixtureStepChannel2 := make(chan common.FixtureCommand)
	fixtureStepChannels = append(fixtureStepChannels, fixtureStepChannel2)
	fixtureStepChannel3 := make(chan common.FixtureCommand)
	fixtureStepChannels = append(fixtureStepChannels, fixtureStepChannel3)
	fixtureStepChannel4 := make(chan common.FixtureCommand)
	fixtureStepChannels = append(fixtureStepChannels, fixtureStepChannel4)
	fixtureStepChannel5 := make(chan common.FixtureCommand)
	fixtureStepChannels = append(fixtureStepChannels, fixtureStepChannel5)
	fixtureStepChannel6 := make(chan common.FixtureCommand)
	fixtureStepChannels = append(fixtureStepChannels, fixtureStepChannel6)
	fixtureStepChannel7 := make(chan common.FixtureCommand)
	fixtureStepChannels = append(fixtureStepChannels, fixtureStepChannel7)

	// Create eight fixture threads for this sequence.
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 0, fixtureStepChannels[0], eventsForLauchpad, guiButtons, dmxController, fixturesConfig)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 1, fixtureStepChannels[1], eventsForLauchpad, guiButtons, dmxController, fixturesConfig)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 2, fixtureStepChannels[2], eventsForLauchpad, guiButtons, dmxController, fixturesConfig)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 3, fixtureStepChannels[3], eventsForLauchpad, guiButtons, dmxController, fixturesConfig)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 4, fixtureStepChannels[4], eventsForLauchpad, guiButtons, dmxController, fixturesConfig)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 5, fixtureStepChannels[5], eventsForLauchpad, guiButtons, dmxController, fixturesConfig)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 6, fixtureStepChannels[6], eventsForLauchpad, guiButtons, dmxController, fixturesConfig)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 7, fixtureStepChannels[7], eventsForLauchpad, guiButtons, dmxController, fixturesConfig)

	// So this is the outer loop where sequence waits for commands and processes them if we're not playing a sequence.
	// i.e the sequence is in STOP mode and this is the way we change the RUN flag to START a sequence again.

	for {
		sequence.UpdateShift = false

		// Check for any waiting commands.
		sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, 10*time.Millisecond, sequence, channels)

		// Sequence in Switch Mode.
		if sequence.PlaySwitchOnce && sequence.Type == "switch" {
			// Show initial state of switches
			ShowSwitches(mySequenceNumber, &sequence, eventsForLauchpad, guiButtons, dmxController, fixturesConfig)
			sequence.PlaySwitchOnce = false
			sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, 1*time.Microsecond, sequence, channels)
			continue
		}

		// Start flood mode.
		if sequence.StartFlood && sequence.FloodPlayOnce {
			sequence.Run = false
			// Prepare a message to be sent to the fixtures in the sequence.
			command := common.FixtureCommand{
				Type:           sequence.Type,
				SequenceNumber: sequence.Number,
				RGBStartFlood:  sequence.StartFlood,
			}

			// Now tell all the fixtures what they need to do.
			sendToAllFixtures(sequence, fixtureStepChannels, channels, command)
			sequence.StartFlood = false
			sequence.FloodPlayOnce = false
			sequence.Run = false
			continue
		}

		// Stop flood mode.
		if sequence.StopFlood && sequence.FloodPlayOnce {
			// Prepare a message to be sent to the fixtures in the sequence.
			command := common.FixtureCommand{
				Type:           sequence.Type,
				SequenceNumber: sequence.Number,
				RGBStartFlood:  sequence.StartFlood,
				RGBStopFlood:   sequence.StopFlood,
			}
			// Now tell all the fixtures what they need to do.
			sendToAllFixtures(sequence, fixtureStepChannels, channels, command)
			sequence.StopFlood = false
			sequence.FloodPlayOnce = false
			sequence.Run = true
			continue
		}

		// Sequence in Static Mode.
		if sequence.PlayStaticOnce && sequence.Static && !sequence.StartFlood {
			// Prepare a message to be sent to the fixtures in the sequence.
			command := common.FixtureCommand{
				Type:            sequence.Type,
				SequenceNumber:  sequence.Number,
				RGBStatic:       sequence.Static,
				RGBStaticColors: sequence.StaticColors,
				Hide:            sequence.Hide,
				Master:          sequence.Master,
				StrobeSpeed:     sequence.StrobeSpeed,
				Blackout:        sequence.Blackout,
			}

			// Now tell all the fixtures what they need to do.
			sendToAllFixtures(sequence, fixtureStepChannels, channels, command)
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
						// TODO eventually Music speed will be set by the BPM analyser.
						// But this hasn't been written yet. We just have some framework code
						// in pkg/sound which counts peaks and this is where we display them.
						common.UpdateStatusBar(fmt.Sprintf("BPM %03d", trigger.BPM), "bpm", guiButtons)

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

				// Setup rgb patterns.
				if sequence.Type == "rgb" {
					sequence.Steps = sequence.RGBAvailablePatterns[sequence.SelectedPattern].Steps
					sequence.UpdatePattern = false
				}

				// Setup scanner patterns.
				if sequence.Type == "scanner" {

					// Get available scanner patterns.
					sequence.ScannerAvailablePatterns = getAvailableScannerPattens(sequence)
					sequence.UpdatePattern = false

					sequence.Pattern = sequence.ScannerAvailablePatterns[sequence.SelectedPattern]
					sequence.Steps = sequence.Pattern.Steps

					if sequence.AutoColor {
						sequence.ScannerGobo++
						if sequence.ScannerGobo > 7 {
							sequence.ScannerGobo = 0
						}

						scannerLastColor := 0

						// AvailableFixtures give the real number of configured scanners.
						for _, fixture := range sequence.ScannersAvailable {
							// First check that this fixture has some configured colors.
							colors, ok := sequence.ScannerAvailableColors[fixture.Number]
							if ok {
								// Found a scanner with some colors.
								totalColorForThisFixture := len(colors)

								sequence.ScannerColor[fixture.Number-1]++
								if sequence.ScannerColor[fixture.Number-1] > scannerLastColor {
									if sequence.ScannerColor[fixture.Number-1] >= totalColorForThisFixture {
										sequence.ScannerColor[fixture.Number-1] = 0
									}
									scannerLastColor++
									continue
								}
							}
						}
					}
				}

				// Check is any commands are waiting.
				sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, 10*time.Millisecond, sequence, channels)
				if !sequence.Run || sequence.StartFlood || sequence.Static || sequence.UpdatePattern || sequence.UpdateShift {
					break
				}

				// Calculate positions for each scanner based on the steps in the pattern.
				if sequence.Type == "scanner" {
					for fixture := 0; fixture < sequence.NumberFixtures; fixture++ {
						positions, num := calculateScannerPositions("scanner", sequence.Steps, sequence.Bounce, sequence.ScannerState[fixture].Inverted, 14)
						sequence.NumberSteps = num
						sequence.ScannerPositions[fixture] = make(map[int][]common.Position, 9)
						for key, value := range positions {
							sequence.ScannerPositions[fixture][key] = value
						}
					}
				}

				// At this point colors are solid colors from the patten and not faded yet.
				// an ideal point to replace colors in a sequence.
				// If we are updating the color in a sequence.
				if sequence.UpdateSequenceColor && sequence.Type == "rgb" {
					if sequence.RecoverSequenceColors {
						if sequence.SavedSequenceColors != nil {
							sequence.Steps = replaceRGBcolorsInSteps(sequence.Steps, sequence.SequenceColors)
							sequence.AutoColor = false
						}
					} else {
						sequence.Steps = replaceRGBcolorsInSteps(sequence.Steps, sequence.SequenceColors)
						// Save the current color selection.
						if sequence.SaveColors {
							sequence.SavedSequenceColors = common.HowManyColors(sequence.RGBPositions)
							sequence.SaveColors = false
						}
					}
				}

				if sequence.Invert && sequence.Type == "rgb" {
					patterns := pattern.MakePatterns()
					sequence.Steps = invertRGBColors(patterns[sequence.SelectedPattern].Steps, common.HowManyStepColors(patterns[sequence.SelectedPattern].Steps))
				}

				if sequence.Type == "rgb" {
					// Calculate fade curve values.
					slopeOn, slopeOff := calculateFadeValues(sequence.RGBFade, sequence.RGBSize)
					// Calulate positions for each RGB fixture.
					sequence.RGBPositions, sequence.NumberSteps = calculateRGBPositions(sequence, slopeOn, slopeOff)
				}

				// If we are setting the pattern automatically for rgb fixtures.
				if sequence.AutoPattern && sequence.Type == "rgb" {
					for patternNumber, pattern := range sequence.RGBAvailablePatterns {
						if pattern.Number == sequence.SelectedPattern {
							sequence.Pattern.Number = patternNumber
							if debug {
								fmt.Printf(">>>> I AM PATTEN %d\n", patternNumber)
							}
							break
						}
					}
					sequence.SelectedPattern++
					if sequence.SelectedPattern > len(sequence.RGBAvailablePatterns) {
						sequence.SelectedPattern = 0
					}
				}

				// If we are setting the pattern automatically for scanner fixtures.
				if sequence.AutoPattern && sequence.Type == "scanner" {
					sequence.SelectedPattern++
					if sequence.SelectedPattern > 3 {
						sequence.SelectedPattern = 0
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
					sequence.RGBPositions = replaceColors(sequence.RGBPositions, sequence.SequenceColors)
				}

				// Now that the pattern colors have been decided and the positions calculated, set the CurrentSequenceColors
				// with the colors from that pattern.
				for fixture := 0; fixture < sequence.NumberFixtures; fixture++ {
					sequence.CurrentColors = common.HowManyScannerColors(sequence.ScannerPositions[fixture])
				}

				sequence.ScannerStateMutex.RLock()
				scannerState := sequence.ScannerState
				sequence.ScannerStateMutex.RUnlock()

				sequence.DisableOnceMutex.RLock()
				disabledOnce := sequence.DisableOnce
				sequence.DisableOnceMutex.RUnlock()

				// Run through the steps in the sequence.
				// Remember every step contains infomation for all the fixtures in this group.
				for step := 0; step < sequence.NumberSteps; step++ {

					// This is were we set the speed of the sequence to current speed.
					sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence.CurrentSpeed/10, sequence, channels)
					if !sequence.Run || sequence.StartFlood || sequence.Static || sequence.UpdatePattern || sequence.UpdateShift {
						break
					}

					for fixtureNumber, fixture := range fixtureStepChannels {
						command := common.FixtureCommand{
							Step:                   step,
							StrobeSpeed:            sequence.StrobeSpeed,
							Master:                 sequence.Master,
							Blackout:               sequence.Blackout,
							Hide:                   sequence.Hide,
							Invert:                 sequence.Invert,
							Type:                   sequence.Type,
							RGBSize:                sequence.RGBSize,
							RGBFade:                sequence.RGBFade,
							RGBPosition:            sequence.RGBPositions[step],
							RGBStartFlood:          sequence.StartFlood,
							RGBStopFlood:           sequence.StopFlood,
							SequenceNumber:         sequence.Number,
							ScannerPosition:        sequence.ScannerPositions[fixtureNumber],
							ScannerSelectedGobo:    sequence.ScannerGobo,
							ScannerState:           scannerState,
							ScannerDisableOnce:     disabledOnce,
							ScannerChase:           sequence.ScannerChase,
							ScannerColor:           sequence.ScannerColor,
							ScannerAvailableColors: sequence.ScannerAvailableColors,
							ScannerOffsetPan:       sequence.ScannerOffsetPan,
							ScannerOffsetTilt:      sequence.ScannerOffsetTilt,
						}

						// Start the fixture group.
						fixture <- command
					}
				}
			}
		}
	}
}

// calculateFadeValues - calculate fade curve values.
func calculateFadeValues(fade int, size int) (slopeOn []int, slopeOff []int) {

	fadeUpValues := getFadeValues(float64(common.MaxBrightness), fade, false)
	fadeOnValues := getFadeOnValues(common.MaxBrightness, size)
	fadeDownValues := getFadeValues(float64(common.MaxBrightness), fade, true)

	slopeOn = append(slopeOn, fadeUpValues...)
	slopeOn = append(slopeOn, fadeOnValues...)
	slopeOn = append(slopeOn, fadeDownValues...)

	slopeOff = append(slopeOff, fadeDownValues...)
	slopeOff = append(slopeOff, fadeUpValues...)

	return slopeOn, slopeOff
}

func invertRGBColors(steps []common.Step, colors []common.Color) []common.Step {

	var insertColor int
	numberColors := len(colors)

	for _, step := range steps {
		// if this step being a fading step ?

		for _, fixture := range step.Fixtures {
			for colorNumber, color := range fixture.Colors {
				//fmt.Printf("insert %d colno %d color %v\n", insertColor, colorNumber, color)
				if insertColor >= numberColors {
					insertColor = 0
				}

				if color.R > 0 || color.G > 0 || color.B > 0 {
					// insert a black.
					fixture.Colors[colorNumber] = common.Color{}
					insertColor++
				} else {
					// its a blank space so insert one of the colors.
					fixture.Colors[colorNumber] = colors[insertColor]
				}
			}
		}
	}

	return steps
}

// Send a command to all the fixtures.
func sendToAllFixtures(sequence common.Sequence, fixtureChannels []chan common.FixtureCommand, channels common.Channels, command common.FixtureCommand) {
	for _, fixture := range fixtureChannels {
		fixture <- command
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

func reverse(in int) int {
	switch in {
	case 0:
		return 10
	case 1:
		return 9
	case 2:
		return 8
	case 3:
		return 7
	case 4:
		return 6
	case 5:
		return 5
	case 6:
		return 4
	case 7:
		return 3
	case 8:
		return 2
	case 9:
		return 1
	case 10:
		return 0
	}

	return 10
}

func calculateRGBPositions(sequence common.Sequence, slopeOn []int, slopeOff []int) (map[int]common.Position, int) {

	positionsOut := make(map[int]common.Position)
	//fixtures := make(map[int][]common.Color)
	fadeColors := make(map[int][]common.Color)

	shift := reverse(sequence.RGBShift)

	var numberFixtures int
	var numberFixturesInThisStep int
	var shiftCounter int

	// First loop make a space in the slope values for each fixture.
	for _, step := range sequence.Steps {
		numberFixturesInThisStep = 0
		for fixtureNumber, fixture := range step.Fixtures {
			numberFixturesInThisStep++

			for _, color := range fixture.Colors {
				// make space for a color
				if color.R > 0 || color.G > 0 || color.B > 0 {
					if !sequence.Invert {
						for _, slope := range slopeOn {
							newColor := common.Color{}
							newColor.R = int((float64(color.R) / 100) * (float64(slope) / 2.55))
							newColor.G = int((float64(color.G) / 100) * (float64(slope) / 2.55))
							newColor.B = int((float64(color.B) / 100) * (float64(slope) / 2.55))
							newColor.MasterDimmer = fixture.MasterDimmer
							fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
						}
					} else {
						for range slopeOn {
							newColor := common.Color{}
							newColor.R = int((float64(color.R) / 100) * (float64(255) / 2.55))
							newColor.G = int((float64(color.G) / 100) * (float64(255) / 2.55))
							newColor.B = int((float64(color.B) / 100) * (float64(255) / 2.55))
							newColor.MasterDimmer = fixture.MasterDimmer
							fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
						}
					}
				} else {
					if !sequence.Invert {
						shiftCounter = 0
						for range slopeOff {
							if shiftCounter == shift {
								break
							}
							newColor := common.Color{}
							newColor.R = int((float64(color.R) / 100) * (float64(0) / 2.55))
							newColor.G = int((float64(color.G) / 100) * (float64(0) / 2.55))
							newColor.B = int((float64(color.B) / 100) * (float64(0) / 2.55))
							newColor.MasterDimmer = fixture.MasterDimmer
							fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
							shiftCounter++
						}
					} else {
						for _, slope := range slopeOff {
							newColor := common.Color{}
							newColor.R = int((float64(color.R) / 100) * (float64(slope) / 2.55))
							newColor.G = int((float64(color.G) / 100) * (float64(slope) / 2.55))
							newColor.B = int((float64(color.B) / 100) * (float64(slope) / 2.55))
							newColor.MasterDimmer = fixture.MasterDimmer
							fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
						}
					}
				}
			}
		}
		if numberFixturesInThisStep > numberFixtures {
			numberFixtures = numberFixturesInThisStep
		}
	}

	if sequence.Bounce {
		for stepNumber := len(sequence.Steps); stepNumber > 0; stepNumber-- {
			step := sequence.Steps[stepNumber-1]
			numberFixturesInThisStep = 0
			for fixtureNumber, fixture := range step.Fixtures {
				numberFixturesInThisStep++

				for _, color := range fixture.Colors {
					// make space for a color
					if color.R > 0 || color.G > 0 || color.B > 0 {
						if !sequence.Invert {
							for _, slope := range slopeOn {
								newColor := common.Color{}
								newColor.R = int((float64(color.R) / 100) * (float64(slope) / 2.55))
								newColor.G = int((float64(color.G) / 100) * (float64(slope) / 2.55))
								newColor.B = int((float64(color.B) / 100) * (float64(slope) / 2.55))
								newColor.MasterDimmer = fixture.MasterDimmer
								fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
							}
						} else {
							for range slopeOn {
								newColor := common.Color{}
								newColor.R = int((float64(color.R) / 100) * (float64(255) / 2.55))
								newColor.G = int((float64(color.G) / 100) * (float64(255) / 2.55))
								newColor.B = int((float64(color.B) / 100) * (float64(255) / 2.55))
								newColor.MasterDimmer = fixture.MasterDimmer
								fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
							}
						}
					} else {
						if !sequence.Invert {
							shiftCounter = 0
							for range slopeOff {
								if shiftCounter == shift {
									break
								}
								newColor := common.Color{}
								newColor.R = int((float64(color.R) / 100) * (float64(0) / 2.55))
								newColor.G = int((float64(color.G) / 100) * (float64(0) / 2.55))
								newColor.B = int((float64(color.B) / 100) * (float64(0) / 2.55))
								newColor.MasterDimmer = fixture.MasterDimmer
								fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
								shiftCounter++
							}
						} else {
							for _, slope := range slopeOff {
								newColor := common.Color{}
								newColor.R = int((float64(color.R) / 100) * (float64(slope) / 2.55))
								newColor.G = int((float64(color.G) / 100) * (float64(slope) / 2.55))
								newColor.B = int((float64(color.B) / 100) * (float64(slope) / 2.55))
								newColor.MasterDimmer = fixture.MasterDimmer
								fadeColors[fixtureNumber] = append(fadeColors[fixtureNumber], newColor)
							}
						}
					}
				}
			}
			if numberFixturesInThisStep > numberFixtures {
				numberFixtures = numberFixturesInThisStep
			}
		}
	}

	counter := len(fadeColors[0])

	if debug {
		// Print out the fixtures so far.
		for fixture := 0; fixture < numberFixtures; fixture++ {
			fmt.Printf("Fixture ")
			for out := 0; out < counter; out++ {
				fmt.Printf("%v", fadeColors[fixture][out])
			}
			fmt.Printf("\n")
		}
	}

	lampOn := make(map[int]bool)
	// Assemble the positions.
	for step := 0; step < counter; step++ {
		// Create a new position.
		newPosition := common.Position{}
		// Add some space for the fixtures.
		newPosition.Fixtures = make(map[int]common.Fixture)

		for fixture := 0; fixture < numberFixtures; fixture++ {
			newFixture := common.Fixture{}

			newColor := common.Color{}
			newColor.R = fadeColors[fixture][step].R
			newColor.G = fadeColors[fixture][step].G
			newColor.B = fadeColors[fixture][step].B

			// Optimisation is applied in this step. We only play out off's to the universe if the lamp is already on.
			/// And in the case of inverted playout only colors if the lamp is already on.
			if !sequence.Invert {
				// We've found a color.
				if fadeColors[fixture][step].R > 0 || fadeColors[fixture][step].G > 0 || fadeColors[fixture][step].B > 0 {
					newFixture.Colors = append(newFixture.Colors, newColor)
					lampOn[fixture] = true
				} else {
					// turn the lamp off, but only if its already on.
					if lampOn[fixture] {
						newFixture.Colors = append(newFixture.Colors, common.Color{})
						lampOn[fixture] = false
					}
				}
			} else {
				// We've found a color. turn it on but only if its already off.
				if fadeColors[fixture][step].R > 0 || fadeColors[fixture][step].G > 0 || fadeColors[fixture][step].B > 0 {
					if !lampOn[fixture] {
						newFixture.Colors = append(newFixture.Colors, newColor)
						lampOn[fixture] = true
					}
				} else {
					// turn the lamp off
					newFixture.Colors = append(newFixture.Colors, common.Color{})
					lampOn[fixture] = false
				}
			}
			newFixture.MasterDimmer = fadeColors[fixture][step].MasterDimmer
			newPosition.Fixtures[fixture] = newFixture
		}

		// Add this position
		positionsOut[step] = newPosition
	}

	if debug {
		// Print out the positions in fixtures order.
		for fixture := 0; fixture < numberFixtures; fixture++ {
			fmt.Printf("Position ")
			for step := 0; step < counter; step++ {
				fmt.Printf("%v", positionsOut[step].Fixtures[fixture].Colors)
			}
			fmt.Printf("\n")
		}
	}

	return positionsOut, counter
}

func makeShift(index int, length int, shift int) int {
	var use int

	if debug {
		fmt.Printf("index %d shift %d length %d\n", index, shift, length)
	}
	if index+shift > length {
		use = index + shift - length
	} else {
		use = index + shift
	}
	if use == length {
		use = 0
	}
	if debug {
		fmt.Printf("I will use %d\n", use)
	}

	return use
}

// calculateScannerPositions takes the steps defined in the pattern and
// turns them into positions used by the sequencer.
func calculateScannerPositions(tYpe string, steps []common.Step, bounce bool, invert bool, shift int) (map[int][]common.Position, int) {

	if debug {
		fmt.Printf("calculateScannerPositions\n")
	}

	position := common.Position{}

	// We have multiple positions for each fixture.
	var counter int
	var waitForColors bool

	positionsOut := make(map[int][]common.Position, 9)

	if !invert {
		waitForColors = false
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
						position.ScannerNumber = fixtureIndex
						position.Color.R = color.R
						position.Color.G = color.G
						position.Color.B = color.B
						positionsOut[counter] = append(positionsOut[counter], position)
						if noColors > 1 {
							if fixture.Type != "scanner" {
								counter = counter + shift
								waitForColors = true
							}
						}
					}
				}
			}
			if !waitForColors {
				counter = counter + shift
			}
		}
	}

	if bounce && tYpe == "rgb" {
		waitForColors = false
		for stepNumber := len(steps); stepNumber > 0; stepNumber-- {
			step := steps[stepNumber-1]
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
						position.ScannerNumber = fixtureIndex
						position.Color.R = color.R
						position.Color.G = color.G
						position.Color.B = color.B
						positionsOut[counter] = append(positionsOut[counter], position)
						if noColors > 1 {
							if fixture.Type != "scanner" {
								counter = counter + shift
								waitForColors = true
							}
						}
					}
				}
			}
			if !waitForColors {
				counter = counter + shift
			}
		}
	}

	if bounce || invert {
		if tYpe == "scanner" {
			// Generate the positions in reverse.
			// Reverse the steps.
			for stepNumber := len(steps); stepNumber > 0; stepNumber-- {
				step := steps[stepNumber-1]

				// Reverse the fixtures.
				for fixtureIndex := len(step.Fixtures); fixtureIndex > 0; fixtureIndex-- {
					fixture := step.Fixtures[fixtureIndex-1]

					position := common.Position{}
					// Reverse the colors.
					noColors := len(fixture.Colors)
					for colorNumber := noColors; colorNumber > 0; colorNumber-- {
						color := fixture.Colors[colorNumber-1]

						position.Gobo = fixture.Gobo
						position.Pan = fixture.Pan
						position.Tilt = fixture.Tilt
						position.Shutter = fixture.Shutter
						position.StartPosition = counter
						position.ScannerNumber = fixtureIndex - 1
						position.Color = color
						positionsOut[counter] = append(positionsOut[counter], position)
						if noColors >= 1 {
							if fixture.Type != "scanner" {
								counter = counter + shift
								waitForColors = true
							}
						}
					}
				}
				if !waitForColors {
					counter = counter + shift
				}
			}
		}
	}

	if bounce && invert {
		if tYpe == "scanner" {
			waitForColors = false
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
							position.ScannerNumber = fixtureIndex
							position.Color.R = color.R
							position.Color.G = color.G
							position.Color.B = color.B
							positionsOut[counter] = append(positionsOut[counter], position)
							if noColors > 1 {
								if fixture.Type != "scanner" {
									counter = counter + shift
									waitForColors = true
								}
							}
						}
					}
				}
				if !waitForColors {
					counter = counter + shift
				}
			}
		}
	}

	return positionsOut, counter
}

func replaceRGBcolorsInSteps(steps []common.Step, colors []common.Color) []common.Step {

	var insertColor int
	numberColors := len(colors)

	for stepNumber, step := range steps {
		for fixtureNumber, fixture := range step.Fixtures {
			for colorNumber, color := range fixture.Colors {
				// found a color.
				if color.R > 0 || color.G > 0 || color.B > 0 {
					if insertColor >= numberColors {
						insertColor = 0
					}
					steps[stepNumber].Fixtures[fixtureNumber].Colors[colorNumber] = colors[insertColor]
					insertColor++
				}
			}
		}
	}

	if debug {
		for stepNumber, step := range steps {
			fmt.Printf("Step %d\n", stepNumber)
			for fixtureNumber, fixture := range step.Fixtures {
				fmt.Printf("\tFixture %d\n", fixtureNumber)
				for _, color := range fixture.Colors {
					fmt.Printf("\t\tColor %+v\n", color)
				}
			}
		}
	}

	return steps
}

// replaceColors can take a sequence and replace its current pattern colors with the colors specified.
func replaceColors(positionsMap map[int]common.Position, colors []common.Color) map[int]common.Position {

	var insertColor int
	numberColors := len(colors)

	replace := make(map[int]common.Position)
	for currentPosition, position := range positionsMap {

		if insertColor >= numberColors {
			insertColor = 0
		}
		//fmt.Printf("Insert color %+v\n", colors[insertColor])
		for fixtureNumber, fixture := range position.Fixtures {
			for colorNumber, color := range fixture.Colors {
				if color.R > 0 || color.G > 0 || color.B > 0 {
					position.Fixtures[fixtureNumber].Colors[colorNumber] = colors[insertColor]
					continue
				}
			}
		}

		insertColor++

		replace[currentPosition] = position

	}

	if debug {
		length := len(positionsMap)
		for step := 0; step < length; step++ {
			fmt.Printf("%v\n", replace[step])
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

// getAvailableScannerColors looks through the fixtures list and finds scanners that
// have colors defined in their config. It then returns an array of these available colors.
// Also returns a map of the default values for each scanner that has colors.
func getAvailableScannerColors(fixtures *fixture.Fixtures) (map[int][]common.StaticColorButton, map[int]int) {

	scannerColors := make(map[int]int)

	availableScannerColors := make(map[int][]common.StaticColorButton)
	for _, fixture := range fixtures.Fixtures {
		for _, channel := range fixture.Channels {
			if strings.Contains(channel.Name, "Color") {
				for _, setting := range channel.Settings {
					newStaticColorButton := common.StaticColorButton{}
					newStaticColorButton.SelectedColor = setting.Number
					newStaticColorButton.Color = common.GetRGBColorByName(setting.Name)
					availableScannerColors[fixture.Number] = append(availableScannerColors[fixture.Number], newStaticColorButton)
					scannerColors[fixture.Number-1] = 0
				}
			}
		}
	}
	return availableScannerColors, scannerColors
}

func getNumberOfFixtures(sequenceNumber int, fixtures *fixture.Fixtures) int {
	if debug {
		fmt.Printf("getNumberOfFixtures\n")
	}

	var numberFixtures int

	for _, fixture := range fixtures.Fixtures {
		if debug {
			fmt.Printf("Fixture Name:%s\n", fixture.Name)
		}

		if debug {
			fmt.Printf("Sequence: %d - Scanner Name: %s Description: %s\n", sequenceNumber, fixture.Name, fixture.Description)
		}

		if fixture.NumberChannels > numberFixtures {
			return fixture.NumberChannels
		}

		if fixture.Group == sequenceNumber+1 {
			if fixture.Number > numberFixtures {
				numberFixtures++
			}

		}
	}

	fmt.Printf("Sequence: %d - Number of Fixtures %d\n", sequenceNumber, numberFixtures)
	return numberFixtures
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

// getAvailableScannerPattens generates scanner patterns and stores them in the sequence.
// Each scanner can then select which pattern to use.
// All scanner patterns have the same number of steps defined by NumberCoordinates.
func getAvailableScannerPattens(sequence common.Sequence) map[int]common.Pattern {

	scannerPattens := make(map[int]common.Pattern)

	// Scanner circle pattern 0
	coordinates := pattern.CircleGenerator(sequence.ScannerSize, sequence.ScannerCoordinates[sequence.ScannerSelectedCoordinates], float64(sequence.ScannerOffsetPan), float64(sequence.ScannerOffsetTilt))
	circlePatten := pattern.GeneratePattern(coordinates, sequence.ScannersTotal, sequence.ScannerShift, sequence.ScannerChase)
	circlePatten.Name = "circle"
	circlePatten.Number = 0
	circlePatten.Label = "Circle"
	scannerPattens[0] = circlePatten

	// Scanner left right pattern 1
	coordinates = pattern.ScanGeneratorLeftRight(128, sequence.ScannerCoordinates[sequence.ScannerSelectedCoordinates])
	leftRightPatten := pattern.GeneratePattern(coordinates, sequence.ScannersTotal, sequence.ScannerShift, sequence.ScannerChase)
	leftRightPatten.Name = "leftright"
	leftRightPatten.Number = 1
	leftRightPatten.Label = "Left.Right"
	scannerPattens[1] = leftRightPatten

	// Scanner up down pattern 2
	coordinates = pattern.ScanGeneratorUpDown(128, sequence.ScannerCoordinates[sequence.ScannerSelectedCoordinates])
	upDownPatten := pattern.GeneratePattern(coordinates, sequence.ScannersTotal, sequence.ScannerShift, sequence.ScannerChase)
	upDownPatten.Name = "updown"
	upDownPatten.Number = 2
	upDownPatten.Label = "Up.Down"
	scannerPattens[2] = upDownPatten

	// Scanner zig zag pattern 3
	coordinates = pattern.ScanGenerateSineWave(255, 5000, sequence.ScannerCoordinates[sequence.ScannerSelectedCoordinates])
	zigZagPatten := pattern.GeneratePattern(coordinates, sequence.ScannersTotal, sequence.ScannerShift, sequence.ScannerChase)
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

func getFadeValues(size float64, fade int, reverse bool) []int {

	out := []int{}
	outPadded := []int{}
	size = size / 2

	var numberCoordinates float64

	if fade == 1 {
		numberCoordinates = 20
	}
	if fade == 2 {
		numberCoordinates = 25
	}
	if fade == 3 {
		numberCoordinates = 30
	}
	if fade == 4 {
		numberCoordinates = 35
	}
	if fade == 5 {
		numberCoordinates = 40
	}
	if fade == 6 {
		numberCoordinates = 45
	}
	if fade == 7 {
		numberCoordinates = 50
	}
	if fade == 8 {
		numberCoordinates = 55
	}
	if fade == 9 {
		numberCoordinates = 60
	}
	if fade == 10 {
		numberCoordinates = 65
	}

	var theta float64
	var x float64
	if reverse {
		for x = 0; x <= 180; x += numberCoordinates {
			theta = (x - 90) * math.Pi / 180
			x := int(-size*math.Sin(theta) + size)
			out = append(out, x)
		}
	} else {
		for x = 180; x >= 0; x -= numberCoordinates {
			theta = (x - 90) * math.Pi / 180
			x := int(-size*math.Sin(theta) + size)
			out = append(out, x)
		}
	}

	if reverse {
		for value := 10; value > 0; value-- {
			if value >= len(out) {
				outPadded = append(outPadded, 255)
			} else {
				outPadded = append(outPadded, out[len(out)-value])
			}
		}

	} else {
		for value := 0; value < 10; value++ {
			if value >= len(out) {
				outPadded = append(outPadded, 255)
			} else {
				outPadded = append(outPadded, out[value])
			}
		}
	}

	return outPadded
}

func getFadeOnValues(size int, fade int) []int {

	out := []int{}

	var x int

	for x = 0; x < fade; x++ {
		x := size
		out = append(out, x)
	}

	return out
}
