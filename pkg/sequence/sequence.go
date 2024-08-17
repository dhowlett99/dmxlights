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
	"fmt"
	"image/color"
	"time"

	"github.com/dhowlett99/dmxlights/pkg/commands"
	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
	"github.com/dhowlett99/dmxlights/pkg/position"
	"github.com/dhowlett99/dmxlights/pkg/sound"

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

		// Check for any waiting commands. Setting a large timeout means that we only return when we hava a command.
		sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, 500*time.Hour, sequence, channels, fixturesConfig)

		// Soft fade downs should be disabled for blackout.
		// Send blackout messages to all fixtures.
		// And then continue on to process further commands.
		if sequence.Blackout {
			if debug {
				fmt.Printf("%d: Blackout\n", mySequenceNumber)
			}
			blackout(fixtureStepChannels)
			sequence.Blackout = false
		}

		// Clear all fixtures.
		if sequence.Clear {
			if debug {
				fmt.Printf("%d: Clear\n", mySequenceNumber)
			}
			clearSequence(mySequenceNumber, &sequence, fixtureStepChannels)
			sequence.Clear = false
		}

		// Show all switches.
		if sequence.PlaySwitchOnce && !sequence.PlaySingleSwitch && !sequence.OverrideSpeed && !sequence.OverrideShift && !sequence.OverrideSize && !sequence.OverrideFade && sequence.Type == "switch" {
			if debug {
				fmt.Printf("%d: Show All Switches\n", mySequenceNumber)
			}
			showAllSwitches(mySequenceNumber, &sequence, fixtureStepChannels, eventsForLauchpad, guiButtons)
			sequence.PlaySwitchOnce = false
		}

		// Show the selected switch.
		if sequence.PlaySwitchOnce && sequence.PlaySingleSwitch && !sequence.OverrideSpeed && sequence.Type == "switch" {
			if debug {
				fmt.Printf("%d: Show Single Switch\n", mySequenceNumber)
			}
			showSelectedSwitch(mySequenceNumber, &sequence, fixtureStepChannels, eventsForLauchpad, guiButtons)
			sequence.PlaySwitchOnce = false
			sequence.PlaySingleSwitch = false
		}

		// Override the selected switch.
		if sequence.PlaySwitchOnce && sequence.OverrideSpeed && sequence.Type == "switch" {
			if debug {
				fmt.Printf("%d: Override Single Switch\n", mySequenceNumber)
			}
			overrideSwitch(mySequenceNumber, &sequence, switchChannels)
			sequence.PlaySwitchOnce = false
			sequence.OverrideSpeed = false
		}

		// Start flood.
		if sequence.StartFlood && sequence.FloodPlayOnce && sequence.Type != "switch" {
			if debug {
				fmt.Printf("%d: Start Flood\n", mySequenceNumber)
			}
			startFlood(mySequenceNumber, &sequence, fixtureStepChannels)
			sequence.StartFlood = false
			sequence.FloodPlayOnce = false
		}

		// Stop flood.
		if sequence.StopFlood && sequence.FloodPlayOnce && sequence.Type != "switch" {
			if debug {
				fmt.Printf("%d: Stop Flood\n", mySequenceNumber)
			}
			stopFlood(mySequenceNumber, &sequence, fixtureStepChannels)
			sequence.StopFlood = false
			sequence.FloodPlayOnce = false
		}

		// Sequence in Static Mode.
		if sequence.PlayStaticOnce && sequence.Static && !sequence.StartFlood {
			if debug {
				fmt.Printf("%d: Start Static\n", mySequenceNumber)
			}
			startStatic(mySequenceNumber, &sequence, channels, fixtureStepChannels)
			sequence.PlayStaticOnce = false
		}

		// Turn Static Off Mode.
		if sequence.PlayStaticOnce && !sequence.Static && !sequence.StartFlood {
			if debug {
				fmt.Printf("%d: Stop Static\n", mySequenceNumber)
			}
			stopStatic(mySequenceNumber, &sequence, channels, fixtureStepChannels)
			sequence.PlayStaticOnce = false
		}

		// Sequence in Normal Running Chase Mode.
		if sequence.Chase {

			setupChase(mySequenceNumber, &sequence, availablePatterns, RGBPositions)

			for sequence.Run && !sequence.Static {
				if debug {
					fmt.Printf("%d: Sequence type %s label %s Chase %t Running %t NumberSteps %d\n", mySequenceNumber, sequence.Type, sequence.Label, sequence.Chase, sequence.Run, sequence.NumberSteps)
				}

				// If the music trigger is being used then the timer is disabled.
				if sequence.MusicTrigger {
					enableMusicTrigger(&sequence, soundConfig)
				} else {
					disableMusicTrigger(&sequence, soundConfig)
				}

				setupChase(mySequenceNumber, &sequence, availablePatterns, RGBPositions)

				// Check is any commands are waiting.
				sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, 10*time.Millisecond, sequence, channels, fixturesConfig)
				if !sequence.Run || sequence.Clear || sequence.StartFlood || sequence.StopFlood ||
					sequence.Static || sequence.UpdatePattern || sequence.UpdateSize {
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
						fadeColors, totalNumberOfSteps := position.CalculatePositions(sequence.Pattern.Steps, sequence, common.IS_SCANNER)
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
				if sequence.UpdateColors && sequence.Type == "rgb" {
					sequence.Pattern.Steps = setupColors(&sequence, RGBPositions)
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
					newColors := []color.RGBA{}
					newColors = append(newColors, sequence.RGBAvailableColors[sequence.RGBColor].Color)
					sequence.SequenceColors = newColors

					// Step through the available colors.
					sequence.RGBColor++
					if sequence.RGBColor > 7 {
						sequence.RGBColor = 0
					}
					sequence.Pattern.Steps = replaceRGBcolorsInSteps(sequence.Pattern.Steps, sequence.SequenceColors)
				}

				if sequence.Type == "rgb" {
					// Calculate fade curve values.
					common.CalculateFadeValues(&sequence)
					// Calulate positions for each RGB fixture.
					sequence.Optimisation = true
					var numberSteps int
					fadeColors, totalNumberOfSteps := position.CalculatePositions(sequence.Pattern.Steps, sequence, common.IS_RGB)
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
					sequence.Pattern.Steps = setupRGBPatterns(&sequence, availablePatterns)
				}

				// If we are setting the pattern automatically for scanner fixtures.
				if sequence.AutoPattern && sequence.Type == "scanner" {
					sequence.SelectedPattern++
					if sequence.SelectedPattern > 3 {
						sequence.SelectedPattern = 0
					}
				}

				// Now that the scanner pattern colors have been decided and the positions calculated, set the cCurrent SequenceColors
				// with the colors from that pattern.
				if sequence.Type == "scanner" {
					sequence.SequenceColors = fixture.HowManyScannerColors(&sequence, fixturesConfig)
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

					// Listen for any commands inside sequence steps, time out for the next step at the speed of the chase.
					// or additionally timeout when get a beat that also triggers the next step.
					sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, speed, sequence, channels, fixturesConfig)
					if !sequence.Run || sequence.Clear || sequence.StartFlood || sequence.StopFlood ||
						sequence.Static || sequence.UpdateShift || sequence.UpdatePattern || sequence.UpdateColors || sequence.UpdateSize {
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
