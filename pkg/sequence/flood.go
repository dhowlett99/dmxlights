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

	"github.com/dhowlett99/dmxlights/pkg/colors"
	"github.com/dhowlett99/dmxlights/pkg/common"
)

func startFlood(mySequenceNumber int, sequence *common.Sequence, fixtureStepChannels []chan common.FixtureCommand) {
	//if debug {
	fmt.Printf("sequence %d Start flood mode for %d fixtures\n", mySequenceNumber, sequence.NumberFixtures)
	//}

	// Init last colors.
	sequence.LastColors = make([]common.LastColor, sequence.NumberFixtures)

	// Prepare a message to be sent to the fixtures in the sequence.
	for fixtureNumber := 0; fixtureNumber < sequence.NumberFixtures; fixtureNumber++ {

		// Set color.
		newColor := common.LastColor{
			RGBColor: colors.White,
		}
		sequence.LastColors = append(sequence.LastColors, newColor)

		command := common.FixtureCommand{
			Master:         sequence.Master,
			Blackout:       sequence.Blackout,
			Type:           sequence.Type,
			Label:          sequence.Label,
			SequenceNumber: sequence.Number,
			StartFlood:     sequence.StartFlood,
			StrobeSpeed:    sequence.StrobeSpeed,
			Strobe:         sequence.Strobe,
			LastColor:      sequence.LastColors[fixtureNumber],
		}

		// Now tell the fixtures what to do.
		fixtureStepChannels[fixtureNumber] <- command
	}
}

func stopFlood(mySequenceNumber int, sequence *common.Sequence, fixtureStepChannels []chan common.FixtureCommand) {
	if debug {
		fmt.Printf("sequence %d Stop flood mode\n", mySequenceNumber)
	}

	// Use the last color when restoring scene.
	if sequence.LastColors != nil {
		// Prepare a message to be sent to the fixtures in the sequence.
		for fixtureNumber := 0; fixtureNumber < sequence.NumberFixtures; fixtureNumber++ {
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
				LastColor:      sequence.LastColors[fixtureNumber],
			}
			// Now tell the fixtures what to do.
			fixtureStepChannels[fixtureNumber] <- command
		}
	}
}
