package sequence

import (
	"bytes"
	"encoding/gob"
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
	"github.com/dhowlett99/dmxlights/pkg/position"
	"github.com/dhowlett99/dmxlights/pkg/sound"

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
	staticColorsButtons := common.SetDefaultStaticColorButtons(mySequenceNumber)

	// Populate the edit sequence colors for this sequence with the defaults.
	sequenceColorButtons := common.SetDefaultStaticColorButtons(mySequenceNumber)

	// Every scanner has a number of colors in its wheel.
	availableScannerColors := make(map[int][]common.StaticColorButton)

	// Find the fixtures.
	availableFixtures := setAvalableFixtures(fixturesConfig)

	fixtureLabels := []string{}
	shutterAddress := make(map[int]int16)

	for _, fixture := range fixturesConfig.Fixtures {
		if fixture.Type == "scanner" {
			fixtureLabels = append(fixtureLabels, fixture.Label)
			for channelNumber, channel := range fixture.Channels {
				if strings.Contains(channel.Name, "Shutter") {
					shutterAddress[fixture.Number] = fixture.Address + int16(channelNumber)
				}
			}
		}
	}

	// Every scanner has a number of gobos in its wheel.
	availableScannerGobos := make(map[int][]common.StaticColorButton)

	// A map of the fixture colors.
	scannerColors := make(map[int]int)

	if sequenceType == "scanner" {
		// Initilise Gobos
		availableScannerGobos = getAvailableScannerGobos(mySequenceNumber, fixturesConfig)

		// Initialise Colors.
		availableScannerColors, scannerColors = getAvailableScannerColors(fixturesConfig)
	}

	// A map of the state of fixtures in the sequence.
	// We can disable a fixture by setting fixture Enabled to false.
	scannerState := make(map[int]common.ScannerState, 8)
	// Find the number of fixtures for this sequence.
	numberFixtures := getNumberOfFixtures(mySequenceNumber, fixturesConfig)
	// Initailise the scanner state for all defined fixtures.
	for x := 0; x < numberFixtures; x++ {
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
		NumberFixtures:         numberFixtures,
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
		SequenceColors:         []common.Color{{R: 0, G: 255, B: 0}},
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
		ScannerCoordinates:     []int{12, 16, 24, 32, 64},
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

	sequence.ScannerPositions = make(map[int]map[int]common.Position, sequence.NumberFixtures)
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
				newSwitch := common.Switch{}
				newSwitch.Name = fixture.Name
				newSwitch.Label = fixture.Label
				newSwitch.Number = fixture.Number
				newSwitch.Description = fixture.Description
				newSwitch.UseFixture = fixture.UseFixture

				newSwitch.States = []common.State{}
				for _, state := range fixture.States {
					newState := common.State{}
					newState.Name = state.Name
					newState.Label = state.Label
					newState.ButtonColor.R = state.ButtonColor.R
					newState.ButtonColor.G = state.ButtonColor.G
					newState.ButtonColor.B = state.ButtonColor.B
					newState.Flash = state.Flash

					// Copy values.
					newState.Values = []common.Value{}
					for _, value := range state.Values {
						newValue := common.Value{}
						newValue.Channel = value.Channel
						newValue.Setting = value.Setting
						newState.Values = append(newState.Values, newValue)
					}

					// Copy actions.
					newState.Actions = []common.Action{}
					for _, action := range state.Actions {
						newAction := common.Action{}
						newAction.Name = action.Name
						newAction.Colors = action.Colors
						newAction.Mode = action.Mode
						newAction.Fade = action.Fade
						newAction.Speed = action.Speed
						newAction.Rotate = action.Rotate
						newAction.Music = action.Music
						newAction.Program = action.Program
						newState.Actions = append(newState.Actions, newAction)
					}

					newSwitch.States = append(newSwitch.States, newState)
				}
				// Add new switch to the list.
				newSwitchList = append(newSwitchList, newSwitch)
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
	SwitchChannels map[int]common.SwitchChannel,
	soundConfig *sound.SoundConfig,
	dmxInterfacePresent bool) {

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
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 0, fixtureStepChannels[0], eventsForLauchpad, guiButtons, dmxController, fixturesConfig, dmxInterfacePresent)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 1, fixtureStepChannels[1], eventsForLauchpad, guiButtons, dmxController, fixturesConfig, dmxInterfacePresent)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 2, fixtureStepChannels[2], eventsForLauchpad, guiButtons, dmxController, fixturesConfig, dmxInterfacePresent)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 3, fixtureStepChannels[3], eventsForLauchpad, guiButtons, dmxController, fixturesConfig, dmxInterfacePresent)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 4, fixtureStepChannels[4], eventsForLauchpad, guiButtons, dmxController, fixturesConfig, dmxInterfacePresent)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 5, fixtureStepChannels[5], eventsForLauchpad, guiButtons, dmxController, fixturesConfig, dmxInterfacePresent)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 6, fixtureStepChannels[6], eventsForLauchpad, guiButtons, dmxController, fixturesConfig, dmxInterfacePresent)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 7, fixtureStepChannels[7], eventsForLauchpad, guiButtons, dmxController, fixturesConfig, dmxInterfacePresent)

	// So this is the outer loop where sequence waits for commands and processes them if we're not playing a sequence.
	// i.e the sequence is in STOP mode and this is the way we change the RUN flag to START a sequence again.
	for {
		sequence.UpdateShift = false

		// Check for any waiting commands.
		sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, 10*time.Millisecond, sequence, channels)

		// Clear all fixtures.
		if sequence.Clear {
			if debug {
				fmt.Printf("sequence %d CLEAR\n", mySequenceNumber)
			}
			// Prepare a message to be sent to the fixtures in the sequence.
			command := common.FixtureCommand{
				Type:           sequence.Type,
				SequenceNumber: sequence.Number,
				Clear:          sequence.Clear,
			}

			// Now tell all the fixtures what they need to do.
			sendToAllFixtures(sequence, fixtureStepChannels, channels, command)
			sequence.Clear = false
			continue
		}

		// Sequence in Switch Mode.
		// Show all switches.
		if sequence.PlaySwitchOnce && !sequence.PlaySingleSwitch && sequence.Type == "switch" {
			if debug {
				fmt.Printf("sequence %d Play all switches mode\n", mySequenceNumber)
			}
			// Show initial state of switches
			ShowSwitches(mySequenceNumber, &sequence, eventsForLauchpad, guiButtons, dmxController, fixturesConfig, SwitchChannels, channels.SoundTriggers, soundConfig, dmxInterfacePresent)
			sequence.PlaySwitchOnce = false
			sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, 1*time.Microsecond, sequence, channels)
			continue
		}

		// Show the selected switch.
		if sequence.PlaySwitchOnce && sequence.PlaySingleSwitch && sequence.Type == "switch" {
			if debug {
				fmt.Printf("sequence %d Play single switch mode\n", mySequenceNumber)
			}
			ShowSingleSwitch(sequence.CurrentSwitch, mySequenceNumber, &sequence, eventsForLauchpad, guiButtons, dmxController, fixturesConfig, SwitchChannels, channels.SoundTriggers, soundConfig, dmxInterfacePresent)
			sequence.PlaySwitchOnce = false
			sequence.PlaySingleSwitch = false
			sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, 1*time.Microsecond, sequence, channels)
			continue
		}

		// Start flood mode.
		if sequence.StartFlood && sequence.FloodPlayOnce {
			if debug {
				fmt.Printf("sequence %d Start flood mode\n", mySequenceNumber)
			}
			// Prepare a message to be sent to the fixtures in the sequence.
			command := common.FixtureCommand{
				Type:            sequence.Type,
				SequenceNumber:  sequence.Number,
				RGBStatic:       sequence.Static,
				RGBStaticColors: sequence.StaticColors,
				StartFlood:      sequence.StartFlood,
				StrobeSpeed:     sequence.StrobeSpeed,
			}

			// Now tell all the fixtures what they need to do.
			sendToAllFixtures(sequence, fixtureStepChannels, channels, command)
			sequence.FloodPlayOnce = false
			continue
		}

		// Stop flood mode.
		if sequence.StopFlood && sequence.FloodPlayOnce {
			if debug {
				fmt.Printf("sequence %d Stop flood mode\n", mySequenceNumber)
			}
			// Prepare a message to be sent to the fixtures in the sequence.
			command := common.FixtureCommand{
				Type:           sequence.Type,
				SequenceNumber: sequence.Number,
				StartFlood:     sequence.StartFlood,
				StopFlood:      sequence.StopFlood,
				StrobeSpeed:    sequence.StrobeSpeed,
			}
			// Now tell all the fixtures what they need to do.
			sendToAllFixtures(sequence, fixtureStepChannels, channels, command)
			sequence.StartFlood = false
			sequence.StopFlood = false
			sequence.FloodPlayOnce = false
			continue
		}

		// Sequence in Static Mode.
		if sequence.PlayStaticOnce && sequence.Static && !sequence.StartFlood {
			if debug {
				fmt.Printf("sequence %d Static mode\n", mySequenceNumber)
			}
			// Turn off any music trigger for this sequence.
			sequence.MusicTrigger = false
			sequence.Functions[common.Function8_Music_Trigger].State = false
			channels.SoundTriggers[mySequenceNumber].State = false

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
			continue
		}

		// This is the inner loop where the sequence runs.
		// Sequence in Normal Running Mode.
		if sequence.Mode == "Sequence" {
			for sequence.Run && !sequence.Static {
				if debug {
					fmt.Printf("sequence %d Running mode\n", mySequenceNumber)
				}
				// Map music trigger function.
				sequence.MusicTrigger = sequence.Functions[common.Function8_Music_Trigger].State

				// If the music trigger is being used then the timer is disabled.
				for triggerNumber, trigger := range channels.SoundTriggers {
					if sequence.MusicTrigger {
						sequence.CurrentSpeed = time.Duration(12 * time.Hour)
						if triggerNumber == mySequenceNumber {
							trigger.State = true
						}
					} else {
						if triggerNumber == mySequenceNumber {
							trigger.State = false
						}
						sequence.CurrentSpeed = commands.SetSpeed(sequence.Speed)
					}
				}

				// Setup rgb patterns.
				if sequence.Type == "rgb" {
					sequence.Steps = sequence.RGBAvailablePatterns[sequence.SelectedPattern].Steps
					sequence.Pattern.Name = sequence.RGBAvailablePatterns[sequence.SelectedPattern].Name
					sequence.Pattern.Label = sequence.RGBAvailablePatterns[sequence.SelectedPattern].Label
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
				if !sequence.Run || sequence.StartFlood || sequence.StopFlood || sequence.Static || sequence.UpdatePattern || sequence.UpdateShift || sequence.UpdateSize {
					break
				}

				// Calculate positions for each scanner based on the steps in the pattern.
				if sequence.Type == "scanner" {
					for fixture := 0; fixture < sequence.NumberFixtures; fixture++ {
						// Calculate fade curve values.
						slopeOn := []int{255}
						slopeOff := []int{0}
						// Calulate positions for each RGB fixture.
						optimisation := true
						// Retrieve the scanner state.
						sequence.ScannerStateMutex.RLock()
						scannerState := sequence.ScannerState
						sequence.ScannerStateMutex.RUnlock()
						// Pass through the inverted / reverse flag.
						sequence.ScannerInvert = sequence.ScannerState[fixture].Inverted
						positions, num := position.CalculatePositions(sequence, slopeOn, slopeOff, optimisation, scannerState)
						sequence.NumberSteps = num

						sequence.ScannerPositions[fixture] = make(map[int]common.Position, 9)
						for positionNumber, position := range positions {
							sequence.ScannerPositions[fixture][positionNumber] = position
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

				if sequence.Type == "rgb" {
					// Calculate fade curve values.
					slopeOn, slopeOff := common.CalculateFadeValues(sequence.RGBFade, sequence.RGBSize)
					// Calulate positions for each RGB fixture.
					optimisation := true
					// Retrieve the scanner state.
					sequence.ScannerStateMutex.RLock()
					scannerState := sequence.ScannerState
					sequence.ScannerStateMutex.RUnlock()
					sequence.RGBPositions, sequence.NumberSteps = position.CalculatePositions(sequence, slopeOn, slopeOff, optimisation, scannerState)
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
				if sequence.AutoColor &&
					sequence.Type == "rgb" &&
					sequence.Pattern.Label != "Multi.Color" &&
					sequence.Pattern.Label != "Color.Chase" {

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

				if sequence.RGBInvert {
					sequence.SequenceColors = common.HowManyColors(sequence.RGBPositions)
					sequence.RGBPositions = invertRGBcolorsInPositions(sequence.RGBPositions, sequence.SequenceColors)
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
					if !sequence.Run || sequence.StartFlood || sequence.StopFlood || sequence.Static || sequence.UpdatePattern || sequence.UpdateShift || sequence.UpdateSize {
						break
					}

					for fixtureNumber, fixture := range fixtureStepChannels {
						if scannerState[fixtureNumber].Enabled {
							command := common.FixtureCommand{
								Step:                   step,
								Rotate:                 sequence.Rotate,
								StrobeSpeed:            sequence.StrobeSpeed,
								Master:                 sequence.Master,
								Blackout:               sequence.Blackout,
								Hide:                   sequence.Hide,
								Type:                   sequence.Type,
								RGBPosition:            sequence.RGBPositions[step],
								StartFlood:             sequence.StartFlood,
								StopFlood:              sequence.StopFlood,
								SequenceNumber:         sequence.Number,
								ScannerPosition:        sequence.ScannerPositions[fixtureNumber][step], // Scanner positions have an additional index for their fixture number.
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
}

func invertRGBColors(steps []common.Step, colors []common.Color) []common.Step {

	var insertColor int
	numberColors := len(colors)

	for _, step := range steps {
		for _, fixture := range step.Fixtures {
			for colorNumber, color := range fixture.Colors {
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
	guiButtons chan common.ALight, dmxController *ft232.DMXController, fixtures *fixture.Fixtures,
	switchChannels map[int]common.SwitchChannel, SoundTriggers []*common.Trigger,
	soundConfig *sound.SoundConfig, dmxInterfacePresent bool) {

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
				fixture.MapSwitchFixture(mySequenceNumber, dmxController, switchNumber, switchData.CurrentState, fixtures, sequence.Blackout, sequence.Master, sequence.Master, switchChannels, SoundTriggers, soundConfig, dmxInterfacePresent)
			}
		}
	}
}

func ShowSingleSwitch(currentSwitch int, mySequenceNumber int, sequence *common.Sequence, eventsForLauchpad chan common.ALight,
	guiButtons chan common.ALight, dmxController *ft232.DMXController, fixtures *fixture.Fixtures,
	switchChannels map[int]common.SwitchChannel, SoundTriggers []*common.Trigger,
	soundConfig *sound.SoundConfig, dmxInterfacePresent bool) {

	if debug {
		fmt.Printf("ShowSingleSwitch for sequence %d\n", mySequenceNumber)
	}

	currentState := sequence.Switches[currentSwitch].CurrentState
	switchNumber := sequence.Switches[currentSwitch].Number - 1
	switchLabel := sequence.Switches[currentSwitch].Label

	for stateNumber, state := range sequence.Switches[currentSwitch].States {

		// For this state.
		if stateNumber == currentState {
			// Use the button color for this state to light the correct color on the launchpad.
			common.LightLamp(common.ALight{X: switchNumber, Y: mySequenceNumber, Red: state.ButtonColor.R, Green: state.ButtonColor.G, Blue: state.ButtonColor.B, Brightness: 255}, eventsForLauchpad, guiButtons)

			// Label the switch.
			common.LabelButton(switchNumber, mySequenceNumber, switchLabel+"\n"+state.Label, guiButtons)

			// Now play all the values for this state.
			fixture.MapSwitchFixture(mySequenceNumber, dmxController, switchNumber, currentState, fixtures, sequence.Blackout, sequence.Master, sequence.Master, switchChannels, SoundTriggers, soundConfig, dmxInterfacePresent)
		}
	}
}

func MakeACopy(src, dist interface{}) (err error) {
	buf := bytes.Buffer{}
	if err = gob.NewEncoder(&buf).Encode(src); err != nil {
		return
	}
	return gob.NewDecoder(&buf).Decode(dist)
}

func replaceRGBcolorsInSteps(steps []common.Step, colors []common.Color) []common.Step {
	stepsOut := []common.Step{}
	err := MakeACopy(steps, &stepsOut)
	if err != nil {
		fmt.Printf("replaceRGBcolorsInSteps: error failed to copy steps.\n")
	}

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
					stepsOut[stepNumber].Fixtures[fixtureNumber].Colors[colorNumber] = colors[insertColor]
					insertColor++
				}
			}
		}
	}

	if debug {
		for stepNumber, step := range stepsOut {
			fmt.Printf("Step %d\n", stepNumber)
			for fixtureNumber, fixture := range step.Fixtures {
				fmt.Printf("\tFixture %d\n", fixtureNumber)
				for _, color := range fixture.Colors {
					fmt.Printf("\t\tColor %+v\n", color)
				}
			}
		}
	}

	return stepsOut
}

func invertRGBcolorsInPositions(positions map[int]common.Position, colors []common.Color) map[int]common.Position {

	var insertColor int
	numberColors := len(colors)
	numberPositions := len(positions)

	for positionNumber := 0; positionNumber < numberPositions; positionNumber++ {
		position := positions[positionNumber]
		for fixtureNumber, fixture := range position.Fixtures {
			for colorNumber, color := range fixture.Colors {
				// found a color.
				if color.R > 0 || color.G > 0 || color.B > 0 {
					// insert a black.
					position.Fixtures[fixtureNumber].Colors[colorNumber] = common.Color{
						R: 0,
						G: 0,
						B: 0,
					}
					insertColor++
					continue
				}
				// found a black.
				if color.R == 0 && color.G == 0 && color.B == 0 {
					// insert one of the colors from the sequence.
					if insertColor >= numberColors {
						insertColor = 0
					}
					position.Fixtures[fixtureNumber].Colors[colorNumber] = colors[insertColor]
					insertColor++
					continue
				}
			}
		}
	}

	return positions
}

// replaceColors can take a sequence and replace its current pattern colors with the colors specified.
func replaceColors(positionsMap map[int]common.Position, colors []common.Color) map[int]common.Position {

	var insertColor int
	numberColors := len(colors)

	replace := make(map[int]common.Position)
	lengthPositions := len(positionsMap)
	for currentPosition := 0; currentPosition < lengthPositions; currentPosition++ {
		position := positionsMap[currentPosition]
		if insertColor >= numberColors {
			insertColor = 0
		}
		lengthFixtures := len(position.Fixtures)
		for fixtureNumber := 0; fixtureNumber < lengthFixtures; fixtureNumber++ {
			fixture := position.Fixtures[fixtureNumber]
			lengthColors := len(fixture.Colors)
			for colorNumber := 0; colorNumber < lengthColors; colorNumber++ {
				color := fixture.Colors[colorNumber]
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
					settingColor, err := common.GetRGBColorByName(setting.Name)
					if err != nil {
						fmt.Printf("error: %d\n", err)
						continue
					}
					newStaticColorButton.Color = settingColor
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
		if fixture.Group-1 == sequenceNumber {
			if fixture.NumberChannels > 0 {
				fmt.Printf("Found Number of Channels def. : %d\n", fixture.NumberChannels)
				return fixture.NumberChannels
			}
			if fixture.Number > numberFixtures {
				numberFixtures++
			}
		}
	}
	return numberFixtures
}

func getAvailableScannerGobos(sequenceNumber int, fixtures *fixture.Fixtures) map[int][]common.StaticColorButton {
	if debug {
		fmt.Printf("getAvailableScannerGobos\n")
	}

	gobos := make(map[int][]common.StaticColorButton)

	for _, f := range fixtures.Fixtures {
		if debug {
			fmt.Printf("Fixture Name:%s\n", f.Name)
		}
		if f.Type == "scanner" {

			if debug {
				fmt.Printf("Sequence: %d - Scanner Name: %s Description: %s\n", sequenceNumber, f.Name, f.Description)
			}
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
	return gobos
}

// getAvailableScannerPattens generates scanner patterns and stores them in the sequence.
// Each scanner can then select which pattern to use.
// All scanner patterns have the same number of steps defined by NumberCoordinates.
func getAvailableScannerPattens(sequence common.Sequence) map[int]common.Pattern {

	scannerPattens := make(map[int]common.Pattern)

	// Scanner circle pattern 0
	coordinates := pattern.CircleGenerator(sequence.ScannerSize, sequence.ScannerCoordinates[sequence.ScannerSelectedCoordinates], float64(sequence.ScannerOffsetPan), float64(sequence.ScannerOffsetTilt))
	circlePatten := pattern.GeneratePattern(coordinates, sequence.NumberFixtures, sequence.ScannerShift, sequence.ScannerChase, sequence.ScannerState)
	circlePatten.Name = "circle"
	circlePatten.Number = 0
	circlePatten.Label = "Circle"
	scannerPattens[0] = circlePatten

	// Scanner left right pattern 1
	coordinates = pattern.ScanGeneratorLeftRight(float64(sequence.ScannerSize), float64(sequence.ScannerCoordinates[sequence.ScannerSelectedCoordinates]))
	leftRightPatten := pattern.GeneratePattern(coordinates, sequence.NumberFixtures, sequence.ScannerShift, sequence.ScannerChase, sequence.ScannerState)
	leftRightPatten.Name = "leftright"
	leftRightPatten.Number = 1
	leftRightPatten.Label = "Left.Right"
	scannerPattens[1] = leftRightPatten

	// // Scanner up down pattern 2
	coordinates = pattern.ScanGeneratorUpDown(float64(sequence.ScannerSize), float64(sequence.ScannerCoordinates[sequence.ScannerSelectedCoordinates]))
	upDownPatten := pattern.GeneratePattern(coordinates, sequence.NumberFixtures, sequence.ScannerShift, sequence.ScannerChase, sequence.ScannerState)
	upDownPatten.Name = "updown"
	upDownPatten.Number = 2
	upDownPatten.Label = "Up.Down"
	scannerPattens[2] = upDownPatten

	// // Scanner zig zag pattern 3
	coordinates = pattern.ScanGenerateSineWave(float64(sequence.ScannerSize), 5000, float64(sequence.ScannerCoordinates[sequence.ScannerSelectedCoordinates]))
	zigZagPatten := pattern.GeneratePattern(coordinates, sequence.NumberFixtures, sequence.ScannerShift, sequence.ScannerChase, sequence.ScannerState)
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
