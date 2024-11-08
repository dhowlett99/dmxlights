// Copyright (C) 2022,2023,2024 dhowlett99.
// This is the dmxlights main sequencers step functions.
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

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
	"github.com/dhowlett99/dmxlights/pkg/sound"
)

func generateSteps(steps []common.Step, availablePatterns map[int]common.Pattern, sequence *common.Sequence, soundConfig *sound.SoundConfig, fixturesConfig *fixture.Fixtures) []common.Step {

	if debug {
		fmt.Printf("generateSteps\n")
	}

	// Setup music trigger.
	if sequence.MusicTrigger {
		enableMusicTrigger(sequence, soundConfig)
	} else {
		disableMusicTrigger(sequence, soundConfig)
	}

	// Set the pattern. steps are generated from patterns. the sequence.SelectedPattern will be used to create steps.
	if sequence.Type == "rgb" {
		steps = updateRGBSteps(steps, availablePatterns, sequence)
		return steps
	}

	if sequence.Type == "scanner" {
		steps = updateScannerPattern(steps, sequence, fixturesConfig)
		return steps
	}

	return steps
}

func updateRGBSteps(steps []common.Step, availablePatterns map[int]common.Pattern, sequence *common.Sequence) []common.Step {

	if debug {
		fmt.Printf("updateRGBSteps\n")
	}

	if sequence.StartPattern {
		if debug {
			fmt.Printf("Start Pattern\n")
		}
		steps = setupNewRGBPattern(sequence, availablePatterns)
		sequence.StartPattern = false
		return steps
	}

	// Auto RGB colors.
	if sequence.AutoColor && sequence.Type == "rgb" && sequence.Pattern.Label != "Multi.Color" && sequence.Pattern.Label != "Color.Chase" {
		if debug {
			fmt.Printf("RGB AutoColor\n")
		}
		steps = rgbAutoColors(sequence, steps)
	}

	// Auto Gobo Change for Chaser.
	if sequence.AutoColor && sequence.Label == "chaser" {
		if debug {
			fmt.Printf("RGB AutoColor Chaser\n")
		}
		steps = chaserAutoGobo(steps, sequence)
	}

	// Auto pattern change.
	if sequence.AutoPattern && sequence.Type == "rgb" {
		if debug {
			fmt.Printf("RGB AutoPattern\n")
		}
		steps = rgbAutoPattern(sequence, availablePatterns)
	}

	// At this point colors are solid colors from the patten and not faded yet.
	// an ideal point to replace colors in a sequence.
	// If we are updating the color in a sequence.
	if sequence.UpdateColors && sequence.Type == "rgb" {
		if debug {
			fmt.Printf("RGB UpdateColors\n")
		}
		if sequence.RecoverSequenceColors {
			if debug {
				fmt.Printf("RGB RecoverSequenceColors\n")
			}
			if sequence.SavedSequenceColors != nil {
				if debug {
					fmt.Printf("RGB SavedSequenceColors\n")
				}
				// Recover origial colors after auto color is switched off.
				steps = replaceRGBcolorsInSteps(sequence.Pattern.Name, steps, sequence.SequenceColors)
				sequence.AutoColor = false
			}
		} else {
			// We are updating color in sequence and sequence colors are set.
			if len(sequence.SequenceColors) > 0 {
				if debug {
					fmt.Printf("replaceRGBcolorsInSteps\n")
				}
				steps = replaceRGBcolorsInSteps(sequence.Pattern.Name, steps, sequence.SequenceColors)
				// Save the current color selection.
				if sequence.SaveColors {
					sequence.SavedSequenceColors = common.HowManyColorsInSteps(steps)
					sequence.SaveColors = false
				}
			}
		}
		sequence.UpdateColors = false
	}

	return steps
}

func clearFixture(fixtureNumber int, fixtureStepChannels []chan common.FixtureCommand) {
	command := common.FixtureCommand{
		Clear: true,
	}
	// Start the fixture group.
	fixtureStepChannels[fixtureNumber] <- command
}

func playStep(sequence *common.Sequence, step int, fixtureNumber int, rgbPositions map[int]common.Position, scannerPositions map[int]map[int]common.Position, fixtureStepChannels []chan common.FixtureCommand) {

	if debug {
		fmt.Printf("playStep number %d to fixture %d\n", step, fixtureNumber)
	}

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
