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
	"time"

	"github.com/dhowlett99/dmxlights/pkg/commands"
	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
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

	var steps []common.Step
	rgbPositions := make(map[int]common.Position)
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

		// Sequence in static mode.
		if sequence.PlayStaticOnce && sequence.Static && !sequence.StartFlood {
			if debug {
				fmt.Printf("%d: Start Static\n", mySequenceNumber)
			}
			startStatic(mySequenceNumber, &sequence, channels, fixtureStepChannels)
			sequence.PlayStaticOnce = false
		}

		// Turn static mode off.
		if sequence.PlayStaticOnce && !sequence.Static && !sequence.StartFlood {
			if debug {
				fmt.Printf("%d: Stop Static\n", mySequenceNumber)
			}
			stopStatic(mySequenceNumber, &sequence, channels, fixtureStepChannels)
			sequence.PlayStaticOnce = false
		}

		// Sequence in normal running chase mode.
		if sequence.Chase {
			for sequence.Run && !sequence.Static {

				if debug {
					fmt.Printf("%d: Start CHASE Sequence type %s label %s Running %t\n", mySequenceNumber, sequence.Type, sequence.Label, sequence.Run)
				}

				// Setup music trigger.
				if sequence.MusicTrigger {
					enableMusicTrigger(&sequence, soundConfig)
				} else {
					disableMusicTrigger(&sequence, soundConfig)
				}

				// Set the pattern. steps are generated from patterns. the sequence.SelectedPattern will be used to create steps.
				if sequence.UpdatePattern {
					steps = updatePatterns(&sequence, availablePatterns)
					sequence.UpdatePattern = false
				}

				if sequence.UpdateShift && sequence.Type == "scanner" {
					steps = updateScannerPatterns(&sequence)
					sequence.UpdateShift = false
				}

				// Auto RGB colors.
				if sequence.AutoColor && sequence.Type == "rgb" && sequence.Pattern.Label != "Multi.Color" && sequence.Pattern.Label != "Color.Chase" {
					steps = rgbAutoColors(&sequence, steps)
				}

				// Calculate RGB positions.
				if sequence.Type == "rgb" {
					rgbPositions = calculateRGBPositions(&sequence, steps)
				}

				// At this point colors are solid colors from the patten and not faded yet.
				// an ideal point to replace colors in a sequence.
				// If we are updating the color in a sequence.
				if sequence.UpdateColors && sequence.Type == "rgb" {
					if sequence.RecoverSequenceColors {
						if sequence.SavedSequenceColors != nil {
							// Recover origial colors after auto color is switched off.
							steps = replaceRGBcolorsInSteps(steps, sequence.SequenceColors)
							sequence.AutoColor = false
						}
					} else {
						// We are updating color in sequence and sequence colors are set.
						if len(sequence.SequenceColors) > 0 {
							fmt.Printf("sequenc() 236: We are updating color in sequence to. %+v\n", sequence.SequenceColors)
							steps = replaceRGBcolorsInSteps(steps, sequence.SequenceColors)
							// Save the current color selection.
							if sequence.SaveColors {
								sequence.SavedSequenceColors = common.HowManyColorsInPositions(rgbPositions)
								sequence.SaveColors = false
							}
						}
					}
					sequence.UpdateColors = false
				}

				// Save the steps temporarily
				sequence.Pattern.Steps = steps

				// Auto Gobo Change for Chaser.
				if sequence.Label == "chaser" {
					chaserAutoGobo(&sequence)
				}

				// Calculate positions from steps.
				if sequence.Type == "scanner" {
					scannerPositions = calculateScannerPositions(&sequence, steps)
				}

				// Auto pattern change.
				if sequence.AutoPattern && sequence.Type == "rgb" {
					rgbAutoPattern(&sequence, availablePatterns)
				}
				if sequence.AutoPattern && sequence.Type == "scanner" {
					scannerAutoPattern(&sequence)
				}

				// Auto color change.
				if sequence.AutoColor && sequence.Type == "scanner" {
					scannerAutoColor(&sequence)
				}

				// Update scanner colors.
				if sequence.Type == "scanner" {
					sequence.SequenceColors = fixture.HowManyScannerColors(&sequence, fixturesConfig)
				}

				// Check for command.
				sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, 10*time.Millisecond, sequence, channels, fixturesConfig)
				if !sequence.Run || sequence.Clear || sequence.StartFlood || sequence.StopFlood || sequence.Static || sequence.UpdatePattern || sequence.UpdateColors || sequence.UpdateShift || sequence.UpdateSize {
					break
				}

				// This is the inner loop where the sequence runs.
				// Run through the steps in the sequence.
				// Remember every step contains infomation for all the fixtures in this sequence group.
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
							RGBPosition:              rgbPositions[step],
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
