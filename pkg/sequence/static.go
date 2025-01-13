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

	"github.com/dhowlett99/dmxlights/pkg/common"
)

func startStatic(mySequenceNumber int, sequence *common.Sequence, channels common.Channels, fixtureStepChannels []chan common.FixtureCommand) {

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

		// Prepare a message to be sent to the fixtures in the sequence.
		for fixtureNumber := range fixtureStepChannels {

			if debug {
				fmt.Printf("%d:%d Sequence Fade up Once static lastColor %s requested color %s\n", mySequenceNumber, fixtureNumber, common.GetColorNameByRGB(sequence.LastColors[fixtureNumber].RGBColor), common.GetColorNameByRGB(sequence.StaticColors[fixtureNumber].Color))
			}

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
				LastColor:       sequence.LastColors[fixtureNumber],
			}

			// Now tell the fixtures what to do.
			fixtureStepChannels[fixtureNumber] <- command

			// Copy the new color into the last color buffer.
			newColor := common.LastColor{
				RGBColor: sequence.StaticColors[fixtureNumber].Color,
			}
			sequence.LastColors = append(sequence.LastColors, newColor)

		}

		// Done fading for this static scene only reset when we set a static scene again.
		sequence.StaticFadeUpOnce = false
	} else {
		// else just play the static scene.

		if sequence.LastColors != nil {

			// Prepare a message to be sent to the fixtures in the sequence.
			for fixtureNumber := range fixtureStepChannels {

				if debug {
					fmt.Printf("%d:%d Sequence Just Play static lastColor %s requested color %s\n", mySequenceNumber, fixtureNumber, common.GetColorNameByRGB(sequence.LastColors[fixtureNumber].RGBColor), common.GetColorNameByRGB(sequence.StaticColors[fixtureNumber].Color))
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
					LastColor:       sequence.LastColors[fixtureNumber],
				}

				// Now tell the fixtures what to do.
				fixtureStepChannels[fixtureNumber] <- command

				// Copy the new color into the last color buffer.
				newColor := common.LastColor{
					RGBColor: sequence.StaticColors[fixtureNumber].Color,
				}
				sequence.LastColors = append(sequence.LastColors, newColor)

			}
		}
	}
}

func stopStatic(mySequenceNumber int, sequence *common.Sequence, channels common.Channels, fixtureStepChannels []chan common.FixtureCommand) {
	if debug {
		fmt.Printf("%d: Sequence RGB Static mode OFF Type %s Label %s \n", mySequenceNumber, sequence.Type, sequence.Label)
	}

	channels.SoundTriggers[mySequenceNumber].State = false

	if sequence.LastColors != nil {

		// Prepare a message to be sent to the fixtures in the sequence.
		for fixtureNumber := range fixtureStepChannels {
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
				LastColor:       sequence.LastColors[fixtureNumber],
			}

			// Now tell the fixtures what to do.
			fixtureStepChannels[fixtureNumber] <- command
		}
	}
}
