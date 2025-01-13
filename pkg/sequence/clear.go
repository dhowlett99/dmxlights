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

func clearSequence(mySequenceNumber int, sequence *common.Sequence, fixtureStepChannels []chan common.FixtureCommand) {

	if debug {
		fmt.Printf("sequence %d CLEAR\n", mySequenceNumber)
	}
	// Set color.
	newColor := common.LastColor{
		RGBColor: colors.White,
	}

	// Init last colors
	sequence.LastColors = make([]common.LastColor, sequence.NumberFixtures)

	if sequence.LastColors != nil {
		// Prepare a message to be sent to all the fixtures in the sequence.
		for fixtureNumber := 0; fixtureNumber < sequence.NumberFixtures; fixtureNumber++ {

			command := common.FixtureCommand{
				Master:         sequence.Master,
				Blackout:       sequence.Blackout,
				Type:           sequence.Type,
				Label:          sequence.Label,
				SequenceNumber: sequence.Number,
				Clear:          sequence.Clear,
				LastColor:      sequence.LastColors[fixtureNumber],
			}

			// Now tell the fixtures what to do.
			fixtureStepChannels[fixtureNumber] <- command

			// Now set the LastColors.
			sequence.LastColors = append(sequence.LastColors, newColor)
		}
	}
}
