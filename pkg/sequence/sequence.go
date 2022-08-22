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
	sequence.FixtureScannerPositions = make(map[int]map[int][]common.Position, sequence.NumberFixtures)

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
	fixtureControlChannels := []chan common.FixtureCommand{}
	fixtureControlChannels = append(fixtureControlChannels, fixtureChannel1)
	fixtureControlChannels = append(fixtureControlChannels, fixtureChannel2)
	fixtureControlChannels = append(fixtureControlChannels, fixtureChannel3)
	fixtureControlChannels = append(fixtureControlChannels, fixtureChannel4)
	fixtureControlChannels = append(fixtureControlChannels, fixtureChannel5)
	fixtureControlChannels = append(fixtureControlChannels, fixtureChannel6)
	fixtureControlChannels = append(fixtureControlChannels, fixtureChannel7)
	fixtureControlChannels = append(fixtureControlChannels, fixtureChannel8)

	// Create channels used for stepping the fixture threads for this sequnece.
	fixtureStepChannels := []chan common.NewFixtureCommand{}
	fixtureStepChannel0 := make(chan common.NewFixtureCommand)
	fixtureStepChannels = append(fixtureStepChannels, fixtureStepChannel0)
	fixtureStepChannel1 := make(chan common.NewFixtureCommand)
	fixtureStepChannels = append(fixtureStepChannels, fixtureStepChannel1)
	fixtureStepChannel2 := make(chan common.NewFixtureCommand)
	fixtureStepChannels = append(fixtureStepChannels, fixtureStepChannel2)
	fixtureStepChannel3 := make(chan common.NewFixtureCommand)
	fixtureStepChannels = append(fixtureStepChannels, fixtureStepChannel3)
	fixtureStepChannel4 := make(chan common.NewFixtureCommand)
	fixtureStepChannels = append(fixtureStepChannels, fixtureStepChannel4)
	fixtureStepChannel5 := make(chan common.NewFixtureCommand)
	fixtureStepChannels = append(fixtureStepChannels, fixtureStepChannel5)
	fixtureStepChannel6 := make(chan common.NewFixtureCommand)
	fixtureStepChannels = append(fixtureStepChannels, fixtureStepChannel6)
	fixtureStepChannel7 := make(chan common.NewFixtureCommand)
	fixtureStepChannels = append(fixtureStepChannels, fixtureStepChannel7)

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
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 0, fixtureControlChannels[0], fixtureStepChannels[0], fixtureStopChannels[0], eventsForLauchpad, guiButtons, dmxController, fixturesConfig)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 1, fixtureControlChannels[1], fixtureStepChannels[1], fixtureStopChannels[1], eventsForLauchpad, guiButtons, dmxController, fixturesConfig)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 2, fixtureControlChannels[2], fixtureStepChannels[2], fixtureStopChannels[2], eventsForLauchpad, guiButtons, dmxController, fixturesConfig)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 3, fixtureControlChannels[3], fixtureStepChannels[3], fixtureStopChannels[3], eventsForLauchpad, guiButtons, dmxController, fixturesConfig)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 4, fixtureControlChannels[4], fixtureStepChannels[4], fixtureStopChannels[4], eventsForLauchpad, guiButtons, dmxController, fixturesConfig)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 5, fixtureControlChannels[5], fixtureStepChannels[5], fixtureStopChannels[5], eventsForLauchpad, guiButtons, dmxController, fixturesConfig)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 6, fixtureControlChannels[6], fixtureStepChannels[6], fixtureStopChannels[6], eventsForLauchpad, guiButtons, dmxController, fixturesConfig)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 7, fixtureControlChannels[7], fixtureStepChannels[7], fixtureStopChannels[7], eventsForLauchpad, guiButtons, dmxController, fixturesConfig)

	// So this is the outer loop where sequence waits for commands and processes them if we're not playing a sequence.
	// i.e the sequence is in STOP mode and this is the way we change the RUN flag to START a sequence again.

	for {
		sequence.UpdateShift = false

		if !sequence.Run {
			sendStopToAllFixtures(sequence, fixtureStopChannels)
		}

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
				Tick:           true,
				Type:           sequence.Type,
				SequenceNumber: sequence.Number,
				StartFlood:     sequence.StartFlood,
			}
			// Stop the fixture threads doing what the are doing.
			sendStopToAllFixtures(sequence, fixtureStopChannels)

			// Now tell all the fixtures what they need to do.
			sendToAllFixtures(sequence, fixtureControlChannels, channels, command)
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
			sendToAllFixtures(sequence, fixtureControlChannels, channels, command)
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
				StrobeSpeed:    sequence.StrobeSpeed,
				Blackout:       sequence.Blackout,
			}

			// Stop the fixture threads doing what the are doing.
			sendStopToAllFixtures(sequence, fixtureStopChannels)

			// Now tell all the fixtures what they need to do.
			sendToAllFixtures(sequence, fixtureControlChannels, channels, command)
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
					// Tell the fixtures to stop.
					sendStopToAllFixtures(sequence, fixtureStopChannels)
					break
				}

				// Calulate positions for each scanner based on the steps in the pattern.
				if sequence.Type == "scanner" {
					for fixture := 0; fixture < sequence.NumberFixtures; fixture++ {
						positions, num := calculatePositions("scanner", sequence.Steps, sequence.Bounce, sequence.ScannerState[fixture].Inverted, 14)
						sequence.NumberSteps = num
						sequence.FixtureScannerPositions[fixture] = make(map[int][]common.Position, 9)
						for key, value := range positions {
							sequence.FixtureScannerPositions[fixture][key] = value
						}
					}
				}

				// Calulate positions for each RGB fixture.
				if sequence.Type == "rgb" {
					// Invert is done in a differnent way for RGB fixtures so invert flag is always fales here.
					sequence.FixtureRGBPositions, sequence.NumberSteps = calculatePositions("rgb", sequence.Steps, sequence.Bounce, false, sequence.RGBShift)
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
					sequence.FixtureRGBPositions = replaceColors(sequence.FixtureRGBPositions, sequence.SequenceColors)
				}

				// If we are updating the color in a sequence.
				if sequence.UpdateSequenceColor && sequence.Type == "rgb" {
					if sequence.RecoverSequenceColors {
						if sequence.SavedSequenceColors != nil {
							sequence.FixtureRGBPositions = replaceColors(sequence.FixtureRGBPositions, sequence.SequenceColors)
							sequence.AutoColor = false
						}
					} else {
						sequence.FixtureRGBPositions = replaceColors(sequence.FixtureRGBPositions, sequence.SequenceColors)
						// Save the current color selection.
						if sequence.SaveColors {
							sequence.SavedSequenceColors = common.HowManyColors(sequence.FixtureRGBPositions)
							sequence.SaveColors = false
						}
					}
				}

				// Now that the pattern colors have been decided and the positions calculated, set the CurrentSequenceColors
				// with the colors from that pattern.
				for fixture := 0; fixture < sequence.NumberFixtures; fixture++ {
					sequence.CurrentColors = common.HowManyColors(sequence.FixtureScannerPositions[fixture])
				}

				sequence.ScannerStateMutex.RLock()
				scannerState := sequence.ScannerState
				sequence.ScannerStateMutex.RUnlock()

				sequence.DisableOnceMutex.RLock()
				disabledOnce := sequence.DisableOnce
				sequence.DisableOnceMutex.RUnlock()

				// //fmt.Printf("Step %d \n", step)
				// for fixtureNumber, fixture := range fixtureControlChannels {
				// 	positions := sequence.FixtureRGBPositions[fixtureNumber]

				// 	// Prepare a message to be sent to the fixtures in the sequence.
				// 	command := common.FixtureCommand{
				// 		SequenceNumber:         sequence.Number,
				// 		Invert:                 sequence.Invert,
				// 		Master:                 sequence.Master,
				// 		StrobeSpeed:            sequence.StrobeSpeed,
				// 		Hide:                   sequence.Hide,
				// 		Tick:                   true,
				// 		ScannerPositions:       sequence.FixtureScannerPositions,
				// 		RGBPositions:           positions,
				// 		Type:                   sequence.Type,
				// 		FadeTime:               sequence.FadeTime,
				// 		Size:                   sequence.Size,
				// 		Steps:                  sequence.NumberSteps,
				// 		CurrentSpeed:           sequence.CurrentSpeed,
				// 		Speed:                  sequence.Speed,
				// 		Blackout:               sequence.Blackout,
				// 		StartFlood:             sequence.StartFlood,
				// 		StopFlood:              sequence.StopFlood,
				// 		SelectedGobo:           sequence.ScannerGobo,
				// 		ScannerState:           scannerState,
				// 		DisableOnce:            disabledOnce,
				// 		ScannerChase:           sequence.ScannerChase,
				// 		ScannerColor:           sequence.ScannerColor,
				// 		AvailableScannerColors: sequence.ScannerAvailableColors,
				// 		OffsetPan:              sequence.ScannerOffsetPan,
				// 		OffsetTilt:             sequence.ScannerOffsetTilt,
				// 		FixtureLabels:          sequence.GuiFixtureLabels,
				// 	}

				// 	fixture <- command
				// }

				// Run through the steps in the sequence.
				// Remember every step contains infomation for all the fixtures in this group.
				for step := 0; step < sequence.NumberSteps; step++ {

					//fmt.Printf("Step %d \n", step)

					// This is were we set the speed of the sequence to current speed.
					sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence.CurrentSpeed/20, sequence, channels)
					if !sequence.Run || sequence.StartFlood || sequence.Static || sequence.UpdatePattern || sequence.UpdateShift {
						// Tell the fixtures to stop.
						sendStopToAllFixtures(sequence, fixtureStopChannels)
						break
					}

					for fixtureNumber, fixture := range fixtureStepChannels {
						command := common.NewFixtureCommand{
							Step:                   step,
							StrobeSpeed:            sequence.StrobeSpeed,
							Master:                 sequence.Master,
							Blackout:               sequence.Blackout,
							Hide:                   sequence.Hide,
							Invert:                 sequence.Invert,
							Type:                   sequence.Type,
							RGBSize:                sequence.RGBSize,
							RGBFade:                sequence.RGBFade,
							RGBPositions:           sequence.FixtureRGBPositions,
							RGBStartFlood:          sequence.StartFlood,
							RGBStopFlood:           sequence.StopFlood,
							ScannerPositions:       sequence.FixtureScannerPositions[fixtureNumber],
							ScannerSelectedGobo:    sequence.ScannerGobo,
							ScannerState:           scannerState,
							ScannerDisableOnce:     disabledOnce,
							ScannerChase:           sequence.ScannerChase,
							ScannerColor:           sequence.ScannerColor,
							ScannerAvailableColors: sequence.ScannerAvailableColors,
							ScannerOffsetPan:       sequence.ScannerOffsetPan,
							ScannerOffsetTilt:      sequence.ScannerOffsetTilt,
						}
						fixture <- command
					}
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

// calculatePositions takes the steps defined in the pattern and
// turns them into positions used by the sequencer.
func calculatePositions(tYpe string, steps []common.Step, bounce bool, invert bool, shift int) (map[int][]common.Position, int) {

	position := common.Position{}

	// We have multiple positions for each fixture.
	var counter int
	var waitForColors bool
	positionsOut := make(map[int][]common.Position)

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
						position.Fixture = fixtureIndex
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
						position.Fixture = fixtureIndex
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
				for fixtureNumber := len(step.Fixtures); fixtureNumber > 0; fixtureNumber-- {
					fixture := step.Fixtures[fixtureNumber-1]

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
						position.Fixture = fixtureNumber - 1
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
							position.Fixture = fixtureIndex
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

// replaceColors can take a sequence and replace its current pattern colors with the colors specified.
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
