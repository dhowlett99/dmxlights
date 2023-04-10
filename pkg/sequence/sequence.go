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

	// Create a map of the fixture colors.
	// This will be protected from synchronous access by sequence.ScannerColorMutex
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
	scannerState := make(map[int]common.ScannerState, 8)
	var numberFixtures int
	// Find the number of fixtures for this sequence.
	if sequenceLabel == "chaser" {
		// TODO find the scanner sequence number from the config.
		scannerSequenceNumber := 2
		numberFixtures = getNumberOfFixtures(scannerSequenceNumber, fixturesConfig, false)
	} else {
		numberFixtures = getNumberOfFixtures(mySequenceNumber, fixturesConfig, false)
	}

	// Initailise the scanner state for all defined fixtures.
	for x := 0; x < numberFixtures; x++ {
		newScanner := common.ScannerState{}
		newScanner.Enabled = true
		newScanner.Inverted = false
		scannerState[x] = newScanner
		// Set the first gobo for every fixture.
		scannerGobos[x] = 1
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
		Run:                    false,
		Bounce:                 false,
		ScannerSize:            common.DefaultScannerSize,
		SequenceColors:         common.DefaultSequenceColors,
		RGBSize:                common.DefaultRGBSize,
		Speed:                  common.DefaultSpeed,
		ScannerShift:           common.DefaultScannerShift,
		RGBShift:               common.DefaultRGBShift,
		RGBCoordinates:         common.DefaultRGBCoordinates,
		Blackout:               false,
		Master:                 common.MaxDMXBrightness,
		ScannerGobo:            scannerGobos,
		StartFlood:             false,
		RGBColor:               1,
		AutoColor:              false,
		AutoPattern:            false,
		SelectedPattern:        common.DefaultPattern,
		ScannerState:           scannerState,
		DisableOnce:            disabledOnce,
		ScannerCoordinates:     []int{12, 16, 24, 32, 64},
		ScannerColor:           scannerColors,
		ScannerOffsetPan:       common.ScannerMidPoint,
		ScannerOffsetTilt:      common.ScannerMidPoint,
		GuiFixtureLabels:       fixtureLabels,
	}

	if sequenceType == "switch" {
		// Load the switch information in from the fixtures.yaml file.
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
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 0, fixtureStepChannels[0], eventsForLauchpad, guiButtons, switchChannels, channels.SoundTriggers, soundConfig, dmxController, fixturesConfig, dmxInterfacePresent)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 1, fixtureStepChannels[1], eventsForLauchpad, guiButtons, switchChannels, channels.SoundTriggers, soundConfig, dmxController, fixturesConfig, dmxInterfacePresent)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 2, fixtureStepChannels[2], eventsForLauchpad, guiButtons, switchChannels, channels.SoundTriggers, soundConfig, dmxController, fixturesConfig, dmxInterfacePresent)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 3, fixtureStepChannels[3], eventsForLauchpad, guiButtons, switchChannels, channels.SoundTriggers, soundConfig, dmxController, fixturesConfig, dmxInterfacePresent)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 4, fixtureStepChannels[4], eventsForLauchpad, guiButtons, switchChannels, channels.SoundTriggers, soundConfig, dmxController, fixturesConfig, dmxInterfacePresent)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 5, fixtureStepChannels[5], eventsForLauchpad, guiButtons, switchChannels, channels.SoundTriggers, soundConfig, dmxController, fixturesConfig, dmxInterfacePresent)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 6, fixtureStepChannels[6], eventsForLauchpad, guiButtons, switchChannels, channels.SoundTriggers, soundConfig, dmxController, fixturesConfig, dmxInterfacePresent)
	go fixture.FixtureReceiver(sequence, mySequenceNumber, 7, fixtureStepChannels[7], eventsForLauchpad, guiButtons, switchChannels, channels.SoundTriggers, soundConfig, dmxController, fixturesConfig, dmxInterfacePresent)

	// So this is the outer loop where sequence waits for commands and processes them if we're not playing a sequence.
	// i.e the sequence is in STOP mode and this is the way we change the RUN flag to START a sequence again.
	for {
		sequence.UpdateShift = false

		// Check for any waiting commands.
		sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, 10*time.Millisecond, sequence, channels, fixturesConfig)

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
			ShowSwitches(mySequenceNumber, &sequence, eventsForLauchpad, guiButtons, dmxController, fixturesConfig, channels.SoundTriggers, soundConfig, fixtureStepChannels, dmxInterfacePresent)
			sequence.PlaySwitchOnce = false
			sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, 1*time.Microsecond, sequence, channels, fixturesConfig)
			continue
		}

		// Show the selected switch.
		if sequence.PlaySwitchOnce && sequence.PlaySingleSwitch && sequence.Type == "switch" {
			if debug {
				fmt.Printf("sequence %d Play single switch mode\n", mySequenceNumber)
			}
			ShowSingleSwitch(sequence.CurrentSwitch, mySequenceNumber, &sequence, eventsForLauchpad, guiButtons, dmxController, fixturesConfig, switchChannels, channels.SoundTriggers, soundConfig, dmxInterfacePresent)
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
				Type:            sequence.Type,
				SequenceNumber:  sequence.Number,
				RGBStatic:       sequence.Static,
				RGBStaticColors: sequence.StaticColors,
				StartFlood:      sequence.StartFlood,
				StrobeSpeed:     sequence.StrobeSpeed,
				Strobe:          sequence.Strobe,
			}

			// Now tell all the fixtures what they need to do.
			sendToAllFixtures(sequence, fixtureStepChannels, channels, command)
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
				Type:           sequence.Type,
				SequenceNumber: sequence.Number,
				StartFlood:     sequence.StartFlood,
				StopFlood:      sequence.StopFlood,
				StrobeSpeed:    sequence.StrobeSpeed,
				Strobe:         sequence.Strobe,
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

			sequence.Static = true
			// Turn off any music trigger for this sequence.
			sequence.MusicTrigger = false
			// this.Functions[common.Function8_Music_Trigger].State = false
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
				Strobe:          sequence.Strobe,
				Blackout:        sequence.Blackout,
			}

			// Now tell all the fixtures what they need to do.
			sendToAllFixtures(sequence, fixtureStepChannels, channels, command)
			sequence.PlayStaticOnce = false
			continue
		}

		// Turn Static Mode Off
		if sequence.PlayStaticOnce && !sequence.Static && !sequence.StartFlood {
			if debug {
				fmt.Printf("sequence %d Static Off mode\n", mySequenceNumber)
			}

			sequence.Static = false
			channels.SoundTriggers[mySequenceNumber].State = false

			// Prepare a message to be sent to the fixtures in the sequence.
			command := common.FixtureCommand{
				Type:              sequence.Type,
				SequenceNumber:    sequence.Number,
				RGBStatic:         sequence.Static,
				RGBPlayStaticOnce: sequence.PlayStaticOnce,
				RGBStaticColors:   sequence.StaticColors,
				Hide:              sequence.Hide,
				Master:            sequence.Master,
				StrobeSpeed:       sequence.StrobeSpeed,
				Strobe:            sequence.Strobe,
				Blackout:          sequence.Blackout,
			}

			// Now tell all the fixtures what they need to do.
			sendToAllFixtures(sequence, fixtureStepChannels, channels, command)
			sequence.PlayStaticOnce = false
			sequence.Static = false
			continue
		}

		// Sequence in Normal Running Mode.
		if sequence.Mode == "Sequence" {
			for sequence.Run && !sequence.Static {
				if debug {
					fmt.Printf("sequence %d type %s label %s Running mode\n", mySequenceNumber, sequence.Type, sequence.Label)
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

					var chasePattern common.Pattern
					sequence.EnabledNumberFixtures = pattern.GetNumberEnabledScanners(sequence.ScannerState, sequence.NumberFixtures)

					if sequence.Label == "chaser" {
						// Set the chase RGB steps used to chase the shutter.
						sequence.ScannerChaser = true
						pattenSteps := availablePatterns[sequence.SelectedPattern].Steps
						chasePattern = pattern.ApplyScannerState(pattenSteps, sequence.ScannerState)
					} else {
						chasePattern = availablePatterns[sequence.SelectedPattern]
					}

					steps = chasePattern.Steps
					sequence.Pattern.Name = chasePattern.Name
					sequence.Pattern.Label = chasePattern.Label
					sequence.UpdatePattern = false

				}

				// Setup scanner patterns.
				if sequence.Type == "scanner" {
					// Get available scanner patterns.
					sequence.ScannerAvailablePatterns = getAvailableScannerPattens(sequence)
					sequence.UpdatePattern = false
					sequence.EnabledNumberFixtures = pattern.GetNumberEnabledScanners(sequence.ScannerState, sequence.NumberFixtures)
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
						sequence.FadeUpAndDown = []int{255}
						sequence.FadeDownAndUp = []int{0}
						// Turn on optimasation.
						sequence.Optimisation = true

						// Pass through the inverted / reverse flag.
						sequence.ScannerInvert = sequence.ScannerState[fixture].Inverted
						// Calulate positions for each RGB fixture.
						positions, sequence.NumberSteps = position.CalculatePositions(steps, sequence)

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
							steps = replaceRGBcolorsInSteps(steps, sequence.SequenceColors)
							sequence.AutoColor = false
						}
					} else {
						steps = replaceRGBcolorsInSteps(steps, sequence.SequenceColors)
						// Save the current color selection.
						if sequence.SaveColors {
							sequence.SavedSequenceColors = common.HowManyColors(RGBPositions)
							sequence.SaveColors = false
						}
					}
				}

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
					newColor := []common.Color{}
					newColor = append(newColor, sequence.RGBAvailableColors[sequence.RGBColor].Color)
					sequence.SequenceColors = newColor

					// Step through the available colors.
					sequence.RGBColor++
					if sequence.RGBColor > 7 {
						sequence.RGBColor = 0
					}
					steps = replaceRGBcolorsInSteps(steps, sequence.SequenceColors)
				}

				if sequence.Type == "rgb" {
					// Calculate fade curve values.
					sequence.FadeUpAndDown, sequence.FadeDownAndUp = common.CalculateFadeValues(sequence.RGBCoordinates, sequence.RGBFade, sequence.RGBSize)
					// Calulate positions for each RGB fixture.
					sequence.Optimisation = true
					RGBPositions, sequence.NumberSteps = position.CalculatePositions(steps, sequence)
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

				if sequence.RGBInvert {
					sequence.SequenceColors = common.HowManyColors(RGBPositions)
					RGBPositions = invertRGBcolorsInPositions(RGBPositions, sequence.SequenceColors)
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
							Step:                     step,
							NumberSteps:              sequence.NumberSteps,
							Rotate:                   sequence.Rotate,
							StrobeSpeed:              sequence.StrobeSpeed,
							Strobe:                   sequence.Strobe,
							Master:                   sequence.Master,
							Blackout:                 sequence.Blackout,
							Hide:                     sequence.Hide,
							Type:                     sequence.Type,
							RGBPosition:              RGBPositions[step],
							StartFlood:               sequence.StartFlood,
							StopFlood:                sequence.StopFlood,
							SequenceNumber:           sequence.Number,
							ScannerPosition:          scannerPositions[fixtureNumber][step], // Scanner positions have an additional index for their fixture number.
							ScannerGobo:              sequence.ScannerGobo[fixtureNumber],
							ScannerState:             sequence.ScannerState[fixtureNumber],
							ScannerDisableOnce:       sequence.DisableOnce[fixtureNumber],
							ScannerChaser:            sequence.ScannerChaser,
							ScannerColor:             sequence.ScannerColor[fixtureNumber],
							ScannerAvailableColors:   sequence.ScannerAvailableColors[fixtureNumber],
							ScannerOffsetPan:         sequence.ScannerOffsetPan,
							ScannerOffsetTilt:        sequence.ScannerOffsetTilt,
							ScannerNumberCoordinates: sequence.ScannerCoordinates[sequence.ScannerSelectedCoordinates],
						}

						// Start the fixture group.
						fixtureStepChannels[fixtureNumber] <- command
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
	SoundTriggers []*common.Trigger,
	soundConfig *sound.SoundConfig, fixtureStepChannels []chan common.FixtureCommand, dmxInterfacePresent bool) {

	if debug {
		fmt.Printf("ShowSwitches for sequence %d\n", mySequenceNumber)
	}

	for switchNumber := 0; switchNumber < len(sequence.Switches); switchNumber++ {

		switchData := sequence.Switches[switchNumber]

		if debug {
			fmt.Printf("switchNumber %d state %d\n", switchData.Number, switchData.CurrentPosition)
		}

		state := switchData.States[switchData.CurrentPosition]

		color, _ := common.GetRGBColorByName(state.ButtonColor)
		common.LightLamp(common.ALight{X: switchNumber, Y: mySequenceNumber, Red: color.R, Green: color.G, Blue: color.B, Brightness: 255}, eventsForLauchpad, guiButtons)

		// Label the switch.
		common.LabelButton(switchNumber, mySequenceNumber, switchData.Label+"\n"+state.Label, guiButtons)

		// Now send a message to the fixture to play all the values for this state.
		command := common.FixtureCommand{

			SetSwitch:          true,
			SwitchData:         switchData,
			State:              state,
			CurrentSwitchState: switchData.CurrentPosition,
		}

		// Send a message to the fixture to operate the switch.
		fixtureStepChannels[switchNumber] <- command
	}
}

func ShowSingleSwitch(currentSwitch int, mySequenceNumber int, sequence *common.Sequence, eventsForLauchpad chan common.ALight,
	guiButtons chan common.ALight, dmxController *ft232.DMXController, fixtures *fixture.Fixtures,
	switchChannels []common.SwitchChannel, SoundTriggers []*common.Trigger, soundConfig *sound.SoundConfig, dmxInterfacePresent bool) {

	if debug {
		fmt.Printf("ShowSingleSwitch for sequence %d switch %d\n", mySequenceNumber, currentSwitch)
	}

	swiTch := sequence.Switches[currentSwitch]

	state := sequence.Switches[currentSwitch].States[swiTch.CurrentPosition]

	// Use the button color for this state to light the correct color on the launchpad.
	color, _ := common.GetRGBColorByName(state.ButtonColor)
	common.LightLamp(common.ALight{X: currentSwitch, Y: mySequenceNumber, Red: color.R, Green: color.G, Blue: color.B, Brightness: 255}, eventsForLauchpad, guiButtons)

	// Label the switch.
	common.LabelButton(currentSwitch, mySequenceNumber, swiTch.Label+"\n"+state.Label, guiButtons)

	// Now play all the values for this state.
	fixture.MapSwitchFixture(swiTch, state, dmxController, fixtures, sequence.Blackout, sequence.Master, sequence.Master, switchChannels, SoundTriggers, soundConfig, dmxInterfacePresent, eventsForLauchpad, guiButtons)

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

func getNumberOfFixtures(sequenceNumber int, fixtures *fixture.Fixtures, allPosibleFixtures bool) int {

	if debug {
		fmt.Printf("getNumberOfFixturesn for sequence %d\n", sequenceNumber)
	}

	var numberFixtures int

	for _, fixture := range fixtures.Fixtures {
		if fixture.Group-1 == sequenceNumber {
			// config has use_channels set.
			if fixture.NumberChannels > 0 {
				fmt.Printf("Sequence %d Found Number of Channels def. : %d\n", sequenceNumber, fixture.NumberChannels)
				if allPosibleFixtures {
					numberFixtures = numberFixtures + fixture.NumberChannels
				} else {
					return fixture.NumberChannels
				}

			} else {
				// Examine the channels and count number of color channels.
				// We use Red for the count.
				var subFixture int
				if allPosibleFixtures {
					for _, channel := range fixture.Channels {
						if strings.Contains(channel.Name, "Red") {
							// Found a fixture def.
							subFixture++
						}
					}
				}
				if subFixture > 1 {
					numberFixtures = numberFixtures + subFixture
				} else {
					if fixture.Number > numberFixtures {
						numberFixtures++
					}
				}
			}
		}
	}

	if debug {
		fmt.Printf("numberFixtures found %d\n", numberFixtures)
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
	coordinates := pattern.CircleGenerator(sequence.ScannerSize, sequence.ScannerCoordinates[sequence.ScannerSelectedCoordinates], float64(sequence.ScannerOffsetPan), float64(sequence.ScannerOffsetTilt))
	circlePatten := pattern.GeneratePattern(coordinates, sequence.NumberFixtures, sequence.ScannerShift, sequence.ScannerChaser, sequence.ScannerState)
	circlePatten.Name = "circle"
	circlePatten.Number = 0
	circlePatten.Label = "Circle"
	scannerPattens[0] = circlePatten

	// Scanner left right pattern 1
	coordinates = pattern.ScanGeneratorLeftRight(float64(sequence.ScannerSize), float64(sequence.ScannerCoordinates[sequence.ScannerSelectedCoordinates]))
	leftRightPatten := pattern.GeneratePattern(coordinates, sequence.NumberFixtures, sequence.ScannerShift, sequence.ScannerChaser, sequence.ScannerState)
	leftRightPatten.Name = "leftright"
	leftRightPatten.Number = 1
	leftRightPatten.Label = "Left.Right"
	scannerPattens[1] = leftRightPatten

	// // Scanner up down pattern 2
	coordinates = pattern.ScanGeneratorUpDown(float64(sequence.ScannerSize), float64(sequence.ScannerCoordinates[sequence.ScannerSelectedCoordinates]))
	upDownPatten := pattern.GeneratePattern(coordinates, sequence.NumberFixtures, sequence.ScannerShift, sequence.ScannerChaser, sequence.ScannerState)
	upDownPatten.Name = "updown"
	upDownPatten.Number = 2
	upDownPatten.Label = "Up.Down"
	scannerPattens[2] = upDownPatten

	// // Scanner zig zag pattern 3
	coordinates = pattern.ScanGenerateSineWave(float64(sequence.ScannerSize), 5000, float64(sequence.ScannerCoordinates[sequence.ScannerSelectedCoordinates]))
	zigZagPatten := pattern.GeneratePattern(coordinates, sequence.NumberFixtures, sequence.ScannerShift, sequence.ScannerChaser, sequence.ScannerState)
	zigZagPatten.Name = "zigzag"
	zigZagPatten.Number = 3
	zigZagPatten.Label = "Zig.Zag"
	scannerPattens[3] = zigZagPatten

	coordinates = []pattern.Coordinate{{Pan: 127, Tilt: 127}}
	stopPatten := pattern.GeneratePattern(coordinates, sequence.NumberFixtures, sequence.ScannerShift, sequence.ScannerChaser, sequence.ScannerState)
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

func SequenceSelect(eventsForLauchpad chan common.ALight, guiButtons chan common.ALight, selectedSequence int) {
	// Turn off all sequence lights.
	for seq := 0; seq < 3; seq++ {
		common.LightLamp(common.ALight{X: 8, Y: seq, Brightness: 255, Red: 100, Green: 255, Blue: 255}, eventsForLauchpad, guiButtons)
	}
	// Now turn pink the selected sequence select light.
	common.LightLamp(common.ALight{X: 8, Y: selectedSequence, Brightness: 255, Red: 255, Green: 0, Blue: 255}, eventsForLauchpad, guiButtons)
}
