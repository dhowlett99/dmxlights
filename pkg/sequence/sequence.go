// Copyright (C) 2022,2023 dhowlett99.
// This is the dmxlights main sequencer responsible for controlling all
// of the fixtures in a group.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package sequence

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dhowlett99/dmxlights/pkg/commands"
	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
	"github.com/dhowlett99/dmxlights/pkg/pattern"
	"github.com/dhowlett99/dmxlights/pkg/position"
	"github.com/dhowlett99/dmxlights/pkg/sound"

	"github.com/go-yaml/yaml"
	"github.com/oliread/usbdmx/ft232"
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
//
//	name: sequence name,  a singe word.
//	description: free text describing the sequence.
//	group: assignes to one of the top 4 rows of the launchpad. 1-4
//	type:  rgb, scanner or switch
func LoadSequences() (sequences *SequencesConfig, err error) {
	filename := "sequences.yaml"

	_, err = os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		return nil, errors.New("error: loading sequences.yaml file: " + err.Error())
	}
	data, err := os.ReadFile(filename)
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
	sequenceLabel string,
	mySequenceNumber int,
	fixturesConfig *fixture.Fixtures,
	channels common.Channels) common.Sequence {

	// Populate the static colors for this sequence with the defaults.
	staticColorsButtons := common.SetDefaultStaticColorButtons(mySequenceNumber)

	// Populate the edit sequence colors for this sequence with the defaults.
	sequenceColorButtons := common.SetDefaultStaticColorButtons(mySequenceNumber)

	// Every scanner has a number of colors in its wheel.
	availableScannerColors := make(map[int][]common.StaticColorButton)

	// Find the fixtures.
	availableFixtures := commands.SetAvalableFixtures(fixturesConfig)

	// Setup fixtures labels.
	fixtureLabels := []string{}
	for _, fixture := range fixturesConfig.Fixtures {
		if fixture.Type == "scanner" {
			fixtureLabels = append(fixtureLabels, fixture.Label)
		}
	}

	// Every scanner has a number of gobos in its wheel.
	availableScannerGobos := make(map[int][]common.StaticColorButton)

	// Create a map of the fixture colors.
	scannerColors := make(map[int]int)
	// Create a map of the fixture gobos.
	scannerGobos := make(map[int]int)

	if sequenceType == "scanner" {
		// Initilise Gobos
		availableScannerGobos = getAvailableScannerGobos(mySequenceNumber, fixturesConfig)

		// Initialise Colors.
		availableScannerColors, scannerColors = getAvailableScannerColors(fixturesConfig)
	}

	// A map of the state of fixtures in the sequence.
	// We can disable a fixture by setting fixture Enabled to false.
	FixtureState := make(map[int]common.FixtureState, 8)
	var numberFixtures int
	// Find the number of fixtures for this sequence.
	if sequenceLabel == "chaser" {
		scannerSequenceNumber := common.GlobalScannerSequenceNumber // Scanner sequence number from config.
		numberFixtures = commands.GetNumberOfFixtures(scannerSequenceNumber, fixturesConfig)
	} else {
		numberFixtures = commands.GetNumberOfFixtures(mySequenceNumber, fixturesConfig)
	}

	// Enable all the defined fixtures.
	for x := 0; x < numberFixtures; x++ {
		newScanner := common.FixtureState{}
		newScanner.Enabled = true
		newScanner.RGBInverted = false
		newScanner.ScannerPatternReversed = false
		FixtureState[x] = newScanner
		// Set the first gobo for every fixture.
		scannerGobos[x] = common.DEFAULT_SCANNER_GOBO
	}

	disabledOnce := make(map[int]bool, 8)

	// The actual sequence definition.
	sequence := common.Sequence{
		Label:                  sequenceLabel,
		ScannerAvailableColors: availableScannerColors,
		ScannersAvailable:      availableFixtures,
		NumberFixtures:         numberFixtures,
		Type:                   sequenceType,
		Hidden:                 false,
		Mode:                   "Sequence",
		StaticColors:           staticColorsButtons,
		RGBAvailableColors:     sequenceColorButtons,
		ScannerAvailableGobos:  availableScannerGobos,
		Name:                   sequenceType,
		Number:                 mySequenceNumber,
		RGBFade:                common.DEFAULT_RGB_FADE,
		MusicTrigger:           false,
		Run:                    false,
		Bounce:                 false,
		ScannerSize:            common.DEFAULT_SCANNER_SIZE,
		SequenceColors:         common.DefaultSequenceColors,
		RGBSize:                common.DEFAULT_RGB_SIZE,
		Speed:                  common.DEFAULT_SPEED,
		ScannerShift:           common.DEFAULT_SCANNER_SHIFT,
		RGBShift:               common.DEFAULT_RGB_SHIFT,
		RGBNumberStepsInFade:   common.DEFAULT_RGB_FADE_STEPS,
		Blackout:               false,
		Master:                 common.MAX_DMX_BRIGHTNESS,
		ScannerGobo:            scannerGobos,
		StartFlood:             false,
		RGBColor:               1,
		AutoColor:              false,
		AutoPattern:            false,
		SelectedPattern:        common.DEFAULT_PATTERN,
		FixtureState:           FixtureState,
		DisableOnce:            disabledOnce,
		ScannerCoordinates:     []int{12, 16, 24, 32, 64},
		ScannerColor:           scannerColors,
		ScannerOffsetPan:       common.SCANNER_MID_POINT,
		ScannerOffsetTilt:      common.SCANNER_MID_POINT,
		GuiFixtureLabels:       fixtureLabels,
	}

	// Load the switch information in from the fixtures config.
	if sequenceType == "switch" {
		sequence.Switches = commands.LoadSwitchConfiguration(mySequenceNumber, fixturesConfig)
		sequence.PlaySwitchOnce = true
	}

	if sequenceType == "scanner" {
		// Get available scanner patterns.
		sequence.ScannerAvailablePatterns = getAvailableScannerPattens(sequence)
		sequence.UpdatePattern = false
	}

	return sequence
}

// Now the sequence has been created, this functions starts the sequence.
func PlaySequence(sequence common.Sequence,
	mySequenceNumber int,
	availablePatterns map[int]common.Pattern,
	eventsForLauchpad chan common.ALight,
	guiButtons chan common.ALight,
	dmxController *ft232.DMXController,
	fixturesConfig *fixture.Fixtures,
	channels common.Channels,
	switchChannels []common.SwitchChannel,
	soundConfig *sound.SoundConfig,
	dmxInterfacePresent bool) {

	var steps []common.Step
	RGBPositions := make(map[int]common.Position)
	scannerPositions := make(map[int]map[int]common.Position, sequence.NumberFixtures)

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
	go fixture.FixtureReceiver(0, fixtureStepChannels[0], eventsForLauchpad, guiButtons, switchChannels, channels.SoundTriggers, soundConfig, dmxController, fixturesConfig, dmxInterfacePresent)
	go fixture.FixtureReceiver(1, fixtureStepChannels[1], eventsForLauchpad, guiButtons, switchChannels, channels.SoundTriggers, soundConfig, dmxController, fixturesConfig, dmxInterfacePresent)
	go fixture.FixtureReceiver(2, fixtureStepChannels[2], eventsForLauchpad, guiButtons, switchChannels, channels.SoundTriggers, soundConfig, dmxController, fixturesConfig, dmxInterfacePresent)
	go fixture.FixtureReceiver(3, fixtureStepChannels[3], eventsForLauchpad, guiButtons, switchChannels, channels.SoundTriggers, soundConfig, dmxController, fixturesConfig, dmxInterfacePresent)
	go fixture.FixtureReceiver(4, fixtureStepChannels[4], eventsForLauchpad, guiButtons, switchChannels, channels.SoundTriggers, soundConfig, dmxController, fixturesConfig, dmxInterfacePresent)
	go fixture.FixtureReceiver(5, fixtureStepChannels[5], eventsForLauchpad, guiButtons, switchChannels, channels.SoundTriggers, soundConfig, dmxController, fixturesConfig, dmxInterfacePresent)
	go fixture.FixtureReceiver(6, fixtureStepChannels[6], eventsForLauchpad, guiButtons, switchChannels, channels.SoundTriggers, soundConfig, dmxController, fixturesConfig, dmxInterfacePresent)
	go fixture.FixtureReceiver(7, fixtureStepChannels[7], eventsForLauchpad, guiButtons, switchChannels, channels.SoundTriggers, soundConfig, dmxController, fixturesConfig, dmxInterfacePresent)

	// So this is the outer loop where sequence waits for commands and processes them if we're not playing a sequence.
	// i.e the sequence is in STOP mode and this is the way we change the RUN flag to START a sequence again.
	for {
		sequence.UpdateShift = false

		// Copy in the fixture status into the static color buffer.
		for fixtureNumber := range sequence.StaticColors {
			sequence.StaticColors[fixtureNumber].Enabled = sequence.FixtureState[fixtureNumber].Enabled
		}

		// Check for any waiting commands.
		sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, 10*time.Millisecond, sequence, channels, fixturesConfig)

		// Soft fade downs should be disabled for blackout.
		if sequence.Blackout {
			command := common.FixtureCommand{
				Type:      "lastColor",
				LastColor: common.Black,
			}
			sendToAllFixtures(fixtureStepChannels, command)
		}

		// Clear all fixtures.
		if sequence.Clear {
			if debug {
				fmt.Printf("sequence %d CLEAR\n", mySequenceNumber)
			}
			// Prepare a message to be sent to the fixtures in the sequence.
			command := common.FixtureCommand{
				Master:         sequence.Master,
				Blackout:       sequence.Blackout,
				Type:           sequence.Type,
				Label:          sequence.Label,
				SequenceNumber: sequence.Number,
				Clear:          sequence.Clear,
			}

			// Now tell all the fixtures what they need to do.
			sendToAllFixtures(fixtureStepChannels, command)
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
			for switchNumber := 0; switchNumber < len(sequence.Switches); switchNumber++ {

				switchData := sequence.Switches[switchNumber]

				if debug {
					fmt.Printf("switchNumber %d state %d\n", switchData.Number, switchData.CurrentPosition)
				}

				state := switchData.States[switchData.CurrentPosition]

				color, _ := common.GetRGBColorByName(state.ButtonColor)
				common.LightLamp(common.Button{X: switchNumber, Y: mySequenceNumber}, color, common.MAX_DMX_BRIGHTNESS, eventsForLauchpad, guiButtons)

				// Label the switch.
				common.LabelButton(switchNumber, mySequenceNumber, switchData.Label+"\n"+state.Label, guiButtons)

				// Now send a message to the fixture to play all the values for this state.
				command := common.FixtureCommand{
					Master:             sequence.Master,
					Blackout:           sequence.Blackout,
					Type:               sequence.Type,
					Label:              sequence.Label,
					SequenceNumber:     sequence.Number,
					SwitchData:         switchData,
					State:              state,
					CurrentSwitchState: switchData.CurrentPosition,
					MasterChanging:     sequence.MasterChanging,
					RGBFade:            sequence.RGBFade,
				}

				// Send a message to the fixture to operate the switch.
				fixtureStepChannels[switchNumber] <- command
			}
			sequence.PlaySwitchOnce = false
			sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, 1*time.Microsecond, sequence, channels, fixturesConfig)
			continue
		}

		// Show the selected switch.
		if sequence.PlaySwitchOnce && sequence.PlaySingleSwitch && sequence.Type == "switch" {
			if debug {
				fmt.Printf("sequence %d Play single switch mode\n", mySequenceNumber)
			}

			swiTch := sequence.Switches[sequence.CurrentSwitch]

			state := sequence.Switches[sequence.CurrentSwitch].States[swiTch.CurrentPosition]

			// Use the button color for this state to light the correct color on the launchpad.
			color, _ := common.GetRGBColorByName(state.ButtonColor)
			common.LightLamp(common.Button{X: sequence.CurrentSwitch, Y: mySequenceNumber}, color, common.MAX_DMX_BRIGHTNESS, eventsForLauchpad, guiButtons)

			// Label the switch.
			common.LabelButton(sequence.CurrentSwitch, mySequenceNumber, swiTch.Label+"\n"+state.Label, guiButtons)

			// Now send a message to the fixture to play all the values for this state.
			command := common.FixtureCommand{
				Master:             sequence.Master,
				Blackout:           sequence.Blackout,
				Type:               sequence.Type,
				Label:              sequence.Label,
				SequenceNumber:     sequence.Number,
				SwitchData:         sequence.Switches[sequence.CurrentSwitch],
				State:              sequence.Switches[sequence.CurrentSwitch].States[swiTch.CurrentPosition],
				CurrentSwitchState: swiTch.CurrentPosition,
				MasterChanging:     sequence.MasterChanging,
				RGBFade:            sequence.RGBFade,
			}

			// Send a message to the fixture to operate the switch.
			fixtureStepChannels[sequence.CurrentSwitch] <- command
			sequence.PlaySwitchOnce = false
			sequence.PlaySingleSwitch = false
			sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, 1*time.Microsecond, sequence, channels, fixturesConfig)
			continue
		}

		// Start flood mode.
		if sequence.StartFlood && sequence.FloodPlayOnce && sequence.Type != "switch" {
			if debug {
				fmt.Printf("sequence %d Start flood mode\n", mySequenceNumber)
			}
			// Prepare a message to be sent to the fixtures in the sequence.
			command := common.FixtureCommand{
				Master:         sequence.Master,
				Blackout:       sequence.Blackout,
				Type:           sequence.Type,
				Label:          sequence.Label,
				SequenceNumber: sequence.Number,
				StartFlood:     sequence.StartFlood,
				StrobeSpeed:    sequence.StrobeSpeed,
				Strobe:         sequence.Strobe,
			}

			// Now tell all the fixtures what they need to do.
			sendToAllFixtures(fixtureStepChannels, command)
			sequence.FloodPlayOnce = false
			continue
		}

		// Stop flood mode.
		if sequence.StopFlood && sequence.FloodPlayOnce && sequence.Type != "switch" {
			if debug {
				fmt.Printf("sequence %d Stop flood mode\n", mySequenceNumber)
			}
			// Prepare a message to be sent to the fixtures in the sequence.
			command := common.FixtureCommand{
				Master:         sequence.Master,
				Blackout:       sequence.Blackout,
				Type:           sequence.Type,
				Label:          sequence.Label,
				SequenceNumber: sequence.Number,
				StartFlood:     sequence.StartFlood,
				StopFlood:      sequence.StopFlood,
				StrobeSpeed:    sequence.StrobeSpeed,
				Strobe:         sequence.Strobe,
			}
			// Now tell all the fixtures what they need to do.
			sendToAllFixtures(fixtureStepChannels, command)
			sequence.StartFlood = false
			sequence.StopFlood = false
			sequence.FloodPlayOnce = false
			continue
		}

		// Sequence in Static Mode.
		if sequence.PlayStaticOnce && sequence.Static && !sequence.StartFlood {
			if debug {
				fmt.Printf("%d: Sequence Static mode StaticFadeUpOnce %t\n", mySequenceNumber, sequence.StaticFadeUpOnce)
			}

			sequence.Static = true
			sequence.PlayStaticOnce = false

			// Turn off any music trigger for this sequence.
			sequence.MusicTrigger = false
			// this.Functions[common.Function8_Music_Trigger].State = false
			channels.SoundTriggers[mySequenceNumber].State = false

			// Now send the Fade up command to the fixture.
			if sequence.StaticFadeUpOnce {
				if debug {
					fmt.Printf("%d: Sequence Fade up static \n", mySequenceNumber)
				}
				// Prepare a message to be sent to the fixtures in the sequence.
				command := common.FixtureCommand{
					Master:          sequence.Master,
					Blackout:        sequence.Blackout,
					Type:            sequence.Type,
					Label:           sequence.Label,
					SequenceNumber:  sequence.Number,
					RGBStaticFadeUp: true,
					RGBFade:         sequence.RGBFade,
					RGBStaticColors: sequence.StaticColors,
					Hidden:          false,
					StrobeSpeed:     sequence.StrobeSpeed,
					Strobe:          sequence.Strobe,
					ScannerChaser:   sequence.ScannerChaser,
				}

				// Now tell all the fixtures what they need to do.
				sendToAllFixtures(fixtureStepChannels, command)

				// Done fading for this static scene only reset when we set a static scene again.
				sequence.StaticFadeUpOnce = false
			} else {
				// else just play the static scene.
				if debug {
					fmt.Printf("%d: Sequence Turn on static \n", mySequenceNumber)
				}
				command := common.FixtureCommand{
					Master:          sequence.Master,
					Blackout:        sequence.Blackout,
					Type:            sequence.Type,
					Label:           sequence.Label,
					SequenceNumber:  sequence.Number,
					Hidden:          false,
					StrobeSpeed:     sequence.StrobeSpeed,
					Strobe:          sequence.Strobe,
					ScannerChaser:   sequence.ScannerChaser,
					RGBStaticOn:     true,
					RGBStaticColors: sequence.StaticColors,
				}

				// Now tell all the fixtures what they need to do.
				sendToAllFixtures(fixtureStepChannels, command)
			}
			sequence.PlayStaticOnce = false
			continue
		}

		// Turn Static Off Mode
		if sequence.PlayStaticOnce && !sequence.Static && !sequence.StartFlood {
			if debug {
				fmt.Printf("%d: Sequence RGB Static mode OFF Type %s Label %s \n", mySequenceNumber, sequence.Type, sequence.Label)
			}

			channels.SoundTriggers[mySequenceNumber].State = false

			// Prepare a message to be sent to the fixtures in the sequence.
			command := common.FixtureCommand{
				Master:          sequence.Master,
				Blackout:        sequence.Blackout,
				Type:            sequence.Type,
				Label:           sequence.Label,
				SequenceNumber:  sequence.Number,
				Hidden:          sequence.Hidden,
				StrobeSpeed:     sequence.StrobeSpeed,
				Strobe:          sequence.Strobe,
				ScannerChaser:   sequence.ScannerChaser,
				RGBStaticOff:    true,
				RGBStaticColors: sequence.StaticColors,
				RGBFade:         sequence.RGBFade,
			}

			// Now tell all the fixtures what they need to do.
			sendToAllFixtures(fixtureStepChannels, command)
			sequence.PlayStaticOnce = false

			continue
		}

		// Sequence in Normal Running Mode.
		if sequence.Mode == "Sequence" {
			for sequence.Run && !sequence.Static {
				if debug {
					fmt.Printf("%d: Sequence type %s label %s Running %t\n", mySequenceNumber, sequence.Type, sequence.Label, sequence.Run)
				}

				// If the music trigger is being used then the timer is disabled.
				if sequence.MusicTrigger {
					sequence.CurrentSpeed = time.Duration(12 * time.Hour)
					err := soundConfig.EnableSoundTrigger(sequence.Name)
					if err != nil {
						fmt.Printf("Error while trying to enable sound trigger %s\n", err.Error())
						os.Exit(1)
					}
					if debug {
						fmt.Printf("Sound trigger %s enabled \n", sequence.Name)
					}
					sequence.ChangeMusicTrigger = false
				} else {
					err := soundConfig.DisableSoundTrigger(sequence.Name)
					if err != nil {
						fmt.Printf("Error while trying to disable sound trigger %s\n", err.Error())
						os.Exit(1)
					}
					if debug {
						fmt.Printf("Sound trigger %s disabled\n", sequence.Name)
					}
					sequence.CurrentSpeed = commands.SetSpeed(sequence.Speed)
					sequence.ChangeMusicTrigger = false
				}

				// Setup rgb patterns.
				if sequence.Type == "rgb" {
					RGBPattern := position.ApplyFixtureState(availablePatterns[sequence.SelectedPattern], sequence.FixtureState)
					sequence.EnabledNumberFixtures = pattern.GetNumberEnabledScanners(sequence.FixtureState, sequence.NumberFixtures)
					steps = RGBPattern.Steps
					sequence.Pattern.Name = RGBPattern.Name
					sequence.Pattern.Label = RGBPattern.Label

					// If we are updating the pattern, we also set the represention of the sequence colors.
					if sequence.UpdatePattern {
						sequence.SequenceColors = common.HowManyColorsInSteps(steps)
					}
					sequence.UpdatePattern = false

					// Initialise chaser.
					if sequence.Label == "chaser" {
						// Set the chase RGB steps used to chase the shutter.
						sequence.ScannerChaser = true
						// Chaser start with a standard chase pattern in white.
						steps = replaceRGBcolorsInSteps(steps, []common.Color{{R: 255, G: 255, B: 255}})
					}
				}

				// Setup scanner patterns.
				if sequence.Type == "scanner" {
					// Get available scanner patterns.
					sequence.ScannerAvailablePatterns = getAvailableScannerPattens(sequence)
					sequence.UpdatePattern = false
					sequence.EnabledNumberFixtures = pattern.GetNumberEnabledScanners(sequence.FixtureState, sequence.NumberFixtures)
					// Set the scanner steps used to send out pan and tilt values.
					sequence.Pattern = sequence.ScannerAvailablePatterns[sequence.SelectedPattern]
					steps = sequence.Pattern.Steps

					if sequence.AutoColor {
						// Change all the fixtures to the next gobo.
						for fixtureNumber := range sequence.ScannersAvailable {
							sequence.ScannerGobo[fixtureNumber]++
							if sequence.ScannerGobo[fixtureNumber] > 7 {
								sequence.ScannerGobo[fixtureNumber] = 0
							}
						}
						scannerLastColor := 0

						// AvailableFixtures gives the real number of configured scanners.
						for _, fixture := range sequence.ScannersAvailable {

							// First check that this fixture has some configured colors.
							colors, ok := sequence.ScannerAvailableColors[fixture.Number]
							if ok {
								// Found a scanner with some colors.
								totalColorForThisFixture := len(colors)

								// Now can mess with the scanner color map.
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
				sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, 10*time.Millisecond, sequence, channels, fixturesConfig)
				if !sequence.Run || sequence.Clear || sequence.StartFlood || sequence.StopFlood || sequence.Static || sequence.UpdatePattern || sequence.UpdateShift || sequence.UpdateSize {
					break
				}

				// Calculate positions for each scanner based on the steps in the pattern.
				if sequence.Type == "scanner" {
					if debug {
						fmt.Printf("Scanner Steps\n")
						for stepNumber, step := range sequence.ScannerSteps {
							fmt.Printf("Scanner Steps %+v\n", stepNumber)
							for _, fixture := range step.Fixtures {
								fmt.Printf("Fixture %+v\n", fixture)
							}
						}
					}
					for fixture := 0; fixture < sequence.NumberFixtures; fixture++ {
						var positions map[int]common.Position
						// We're playing out the scanner positions, so we won't need curve values.
						sequence.FadeUp = []int{255}
						sequence.FadeDown = []int{0}
						// Turn on optimasation.
						sequence.Optimisation = true

						// Pass through the inverted / reverse flag.
						sequence.ScannerReverse = sequence.FixtureState[fixture].ScannerPatternReversed
						// Calulate positions for each scanner fixture.
						fadeColors, totalNumberOfSteps := position.CalculatePositions(steps, sequence, common.IS_SCANNER)
						positions, numberSteps := position.AssemblePositions(fadeColors, sequence.NumberFixtures, totalNumberOfSteps, sequence.FixtureState, sequence.Optimisation)
						sequence.NumberSteps = numberSteps

						// Setup positions for each scanner. This is so we can shift the patterns on each scannner.
						scannerPositions[fixture] = make(map[int]common.Position, 9)
						for positionNumber, position := range positions {
							scannerPositions[fixture][positionNumber] = position
						}
					}
				}

				// At this point colors are solid colors from the patten and not faded yet.
				// an ideal point to replace colors in a sequence.
				// If we are updating the color in a sequence.
				if sequence.UpdateSequenceColor && sequence.Type == "rgb" {
					if sequence.RecoverSequenceColors {
						if sequence.SavedSequenceColors != nil {
							// Recover origial colors after auto color is switched off.
							steps = replaceRGBcolorsInSteps(steps, sequence.SequenceColors)
							sequence.AutoColor = false
						}
					} else {
						// We are updating color in sequence and sequence colors are set.
						if len(sequence.SequenceColors) > 0 {
							steps = replaceRGBcolorsInSteps(steps, sequence.SequenceColors)
							// Save the current color selection.
							if sequence.SaveColors {
								sequence.SavedSequenceColors = common.HowManyColorsInPositions(RGBPositions)
								sequence.SaveColors = false
							}
						}
					}
				}
				// Save the steps temporarily
				sequence.Pattern.Steps = steps

				if sequence.Label == "chaser" {
					if sequence.AutoColor {
						// Change all the fixtures to the next gobo.
						for fixtureNumber := range sequence.ScannersAvailable {
							sequence.ScannerGobo[fixtureNumber]++
							if sequence.ScannerGobo[fixtureNumber] > 8 {
								sequence.ScannerGobo[fixtureNumber] = 1
							}
						}
					}
				}

				// If we are setting the current colors in a rgb sequence.
				if sequence.AutoColor &&
					sequence.Type == "rgb" &&
					sequence.Pattern.Label != "Multi.Color" &&
					sequence.Pattern.Label != "Color.Chase" {

					// Find a new color.
					newColors := []common.Color{}
					newColors = append(newColors, sequence.RGBAvailableColors[sequence.RGBColor].Color)
					sequence.SequenceColors = newColors

					// Step through the available colors.
					sequence.RGBColor++
					if sequence.RGBColor > 7 {
						sequence.RGBColor = 0
					}
					steps = replaceRGBcolorsInSteps(steps, sequence.SequenceColors)
				}

				if sequence.Type == "rgb" {
					// Calculate fade curve values.
					common.CalculateFadeValues(&sequence)
					// Calulate positions for each RGB fixture.
					sequence.Optimisation = true
					var numberSteps int
					fadeColors, totalNumberOfSteps := position.CalculatePositions(steps, sequence, common.IS_RGB)
					RGBPositions, numberSteps = position.AssemblePositions(fadeColors, sequence.NumberFixtures, totalNumberOfSteps, sequence.FixtureState, sequence.Optimisation)
					sequence.NumberSteps = numberSteps
				}

				// If we are setting the pattern automatically for rgb fixtures.
				if sequence.AutoPattern && sequence.Type == "rgb" {
					for patternNumber, pattern := range availablePatterns {
						if pattern.Number == sequence.SelectedPattern {
							sequence.Pattern.Number = patternNumber
							if debug {
								fmt.Printf(">>>> I AM PATTEN %d\n", patternNumber)
							}
							break
						}
					}
					sequence.SelectedPattern++
					if sequence.SelectedPattern > len(availablePatterns) {
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

				// Now that the pattern colors have been decided and the positions calculated, set the CurrentSequenceColors
				// with the colors from that pattern.
				for fixture := 0; fixture < sequence.NumberFixtures; fixture++ {
					sequence.CurrentColors = common.HowManyScannerColors(scannerPositions[fixture])
				}

				// This is the inner loop where the sequence runs.
				// Run through the steps in the sequence.
				// Remember every step contains infomation for all the fixtures in this group.
				for step := 0; step < sequence.NumberSteps; step++ {

					// This is were we set the speed of the sequence to current speed.
					speed := sequence.CurrentSpeed / 10
					if sequence.Type == "scanner" {
						speed = sequence.CurrentSpeed / 5 // Slow the scanners down.
					}
					sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, speed, sequence, channels, fixturesConfig)
					if !sequence.Run || sequence.Clear || sequence.StartFlood || sequence.StopFlood || sequence.Static || sequence.UpdatePattern || sequence.UpdateShift || sequence.UpdateSize {
						break
					}

					if debug {
						fmt.Printf("Step %d This many Fixtures %d\n", step, len(RGBPositions[step].Fixtures))
						for _, fixture := range RGBPositions[step].Fixtures {
							fmt.Printf("\t Fixture: %+v\n", fixture)
						}
					}

					for fixtureNumber := 0; fixtureNumber < sequence.NumberFixtures; fixtureNumber++ {

						// Even if the fixture is disabled we still need to send this message to the fixture.
						// beacuse the fixture is the one who is responsible for turning it off.
						command := common.FixtureCommand{
							Master:                   sequence.Master,
							Blackout:                 sequence.Blackout,
							Type:                     sequence.Type,
							Label:                    sequence.Label,
							SequenceNumber:           sequence.Number,
							Step:                     step,
							NumberSteps:              sequence.NumberSteps,
							Rotate:                   sequence.Rotate,
							StrobeSpeed:              sequence.StrobeSpeed,
							Strobe:                   sequence.Strobe,
							RGBFade:                  sequence.RGBFade,
							Hidden:                   sequence.Hidden,
							RGBPosition:              RGBPositions[step],
							StartFlood:               sequence.StartFlood,
							StopFlood:                sequence.StopFlood,
							ScannerPosition:          scannerPositions[fixtureNumber][step], // Scanner positions have an additional index for their fixture number.
							ScannerGobo:              sequence.ScannerGobo[fixtureNumber],
							FixtureState:             sequence.FixtureState[fixtureNumber],
							ScannerDisableOnce:       sequence.DisableOnce[fixtureNumber],
							ScannerChaser:            sequence.ScannerChaser,
							ScannerColor:             sequence.ScannerColor[fixtureNumber],
							ScannerAvailableColors:   sequence.ScannerAvailableColors[fixtureNumber],
							ScannerOffsetPan:         sequence.ScannerOffsetPan,
							ScannerOffsetTilt:        sequence.ScannerOffsetTilt,
							ScannerNumberCoordinates: sequence.ScannerCoordinates[sequence.ScannerSelectedCoordinates],
							MasterChanging:           sequence.MasterChanging,
						}

						// Start the fixture group.
						fixtureStepChannels[fixtureNumber] <- command
					}
				}
			}
		}
	}
}

// Send a command to all the fixtures.
func sendToAllFixtures(fixtureChannels []chan common.FixtureCommand, command common.FixtureCommand) {
	for _, fixture := range fixtureChannels {
		fixture <- command
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

			// found a color.
			if fixture.Color.R > 0 || fixture.Color.G > 0 || fixture.Color.B > 0 {
				if insertColor >= numberColors {
					insertColor = 0
				}
				newFixture := stepsOut[stepNumber].Fixtures[fixtureNumber]
				newFixture.Color = colors[insertColor]
				stepsOut[stepNumber].Fixtures[fixtureNumber] = newFixture
				insertColor++
			}

		}
	}

	if debug {
		for stepNumber, step := range stepsOut {
			fmt.Printf("Step %d\n", stepNumber)
			for fixtureNumber, fixture := range step.Fixtures {
				fmt.Printf("\tFixture %d\n", fixtureNumber)
				fmt.Printf("\t\tColor %+v\n", fixture.Color)
			}
		}
	}

	return stepsOut
}

// getAvailableScannerColors looks through the fixtures list and finds scanners that
// have colors defined in their config. It then returns an array of these available colors.
// Also returns a map of the default values for each scanner that has colors.
func getAvailableScannerColors(fixtures *fixture.Fixtures) (map[int][]common.StaticColorButton, map[int]int) {

	scannerColors := make(map[int]int)

	availableScannerColors := make(map[int][]common.StaticColorButton)
	for _, fixture := range fixtures.Fixtures {
		if fixture.Type == "scanner" {
			for _, channel := range fixture.Channels {
				if strings.Contains(channel.Name, "Color") {
					for _, setting := range channel.Settings {
						newStaticColorButton := common.StaticColorButton{}
						newStaticColorButton.SelectedColor = setting.Number
						settingColor, err := common.GetRGBColorByName(setting.Name)
						if err != nil {
							fmt.Printf("error: %s\n", err)
							continue
						}
						newStaticColorButton.Color = settingColor
						availableScannerColors[fixture.Number] = append(availableScannerColors[fixture.Number], newStaticColorButton)
						scannerColors[fixture.Number-1] = 0
					}
				}
			}
		}
	}
	return availableScannerColors, scannerColors
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
						v, _ := strconv.Atoi(setting.Value)
						newGobo.Setting = v
						newGobo.Color = common.Color{R: 255, G: 255, B: 0} // Yellow.
						gobos[f.Number] = append(gobos[f.Number], newGobo)
						if debug {
							fmt.Printf("\tGobo: %s Setting: %s\n", setting.Name, setting.Value)
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

	if debug {
		fmt.Printf("getAvailableScannerPattens\n")
	}

	scannerPattens := make(map[int]common.Pattern)

	// Scanner circle pattern 0
	coordinates := pattern.CircleGenerator(sequence.ScannerSize, sequence.ScannerCoordinates[sequence.ScannerSelectedCoordinates], float64(sequence.ScannerOffsetTilt), float64(sequence.ScannerOffsetPan))
	circlePatten := pattern.GeneratePattern(coordinates, sequence.NumberFixtures, sequence.ScannerShift, sequence.ScannerChaser, sequence.FixtureState)
	circlePatten.Name = "circle"
	circlePatten.Number = 0
	circlePatten.Label = "Circle"
	scannerPattens[0] = circlePatten

	// Scanner left right pattern 1
	coordinates = pattern.ScanGeneratorLeftRight(float64(sequence.ScannerSize), float64(sequence.ScannerCoordinates[sequence.ScannerSelectedCoordinates]), float64(sequence.ScannerOffsetTilt), float64(sequence.ScannerOffsetPan))
	leftRightPatten := pattern.GeneratePattern(coordinates, sequence.NumberFixtures, sequence.ScannerShift, sequence.ScannerChaser, sequence.FixtureState)
	leftRightPatten.Name = "leftright"
	leftRightPatten.Number = 1
	leftRightPatten.Label = "Left.Right"
	scannerPattens[1] = leftRightPatten

	// // Scanner up down pattern 2
	coordinates = pattern.ScanGeneratorUpDown(float64(sequence.ScannerSize), float64(sequence.ScannerCoordinates[sequence.ScannerSelectedCoordinates]), float64(sequence.ScannerOffsetTilt), float64(sequence.ScannerOffsetPan))
	upDownPatten := pattern.GeneratePattern(coordinates, sequence.NumberFixtures, sequence.ScannerShift, sequence.ScannerChaser, sequence.FixtureState)
	upDownPatten.Name = "updown"
	upDownPatten.Number = 2
	upDownPatten.Label = "Up.Down"
	scannerPattens[2] = upDownPatten

	// // Scanner zig zag pattern 3
	coordinates = pattern.ScanGenerateSawTooth(float64(sequence.ScannerSize), float64(sequence.ScannerCoordinates[sequence.ScannerSelectedCoordinates]), float64(sequence.ScannerCoordinates[sequence.ScannerSelectedCoordinates]), float64(sequence.ScannerOffsetTilt), float64(sequence.ScannerOffsetPan))
	zigZagPatten := pattern.GeneratePattern(coordinates, sequence.NumberFixtures, sequence.ScannerShift, sequence.ScannerChaser, sequence.FixtureState)
	zigZagPatten.Name = "zigzag"
	zigZagPatten.Number = 3
	zigZagPatten.Label = "Zig.Zag"
	scannerPattens[3] = zigZagPatten

	coordinates = []pattern.Coordinate{{Pan: 127, Tilt: 127}}
	stopPatten := pattern.GeneratePattern(coordinates, sequence.NumberFixtures, sequence.ScannerShift, sequence.ScannerChaser, sequence.FixtureState)
	stopPatten.Name = "stop"
	stopPatten.Number = 4
	stopPatten.Label = "Stop"
	scannerPattens[4] = stopPatten

	if debug {
		for _, pattern := range scannerPattens {
			fmt.Printf("Made a pattern called %s\n", pattern.Name)
		}
	}

	return scannerPattens

}
