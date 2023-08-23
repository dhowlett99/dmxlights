// Copyright (C) 2022, 2023 dhowlett99.
// This status function implenents the fixture state, to enable, disable
// invert & revese fixtures. Used by the buttons package.
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

package buttons

import (
	"fmt"

	"github.com/dhowlett99/dmxlights/pkg/common"
)

func setFixtureStatus(this *CurrentState, Y int, X int, commandChannels []chan common.Command, sequence *common.Sequence) {

	// Disable fixture if we're already enabled and inverted.
	if this.FixtureState[Y][X].Enabled && this.FixtureState[Y][X].RGBInverted && X < sequence.NumberFixtures {
		if debug {
			fmt.Printf("Disable fixture Number %d State on Sequence %d to false\n", X, Y)
		}

		this.FixtureState[Y][X].Enabled = false
		this.FixtureState[Y][X].RGBInverted = false

		// Tell the sequence to turn on this scanner.
		cmd := common.Command{
			Action: common.ToggleFixtureState,
			Args: []common.Arg{
				{Name: "FixtureNumber", Value: X},
				{Name: "FixtureState", Value: false},
				{Name: "FixtureInverted", Value: false},
				{Name: "FixtureReversed", Value: false},
			},
		}
		common.SendCommandToSequence(Y, cmd, commandChannels)

		// If we're a scanner also tell the sequence shutter chaser to turn off this fixture.
		if this.SelectedType == "scanner" {
			common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)
		}

		return
	}

	// Enable scanner if not enabled and not inverted.
	if !this.FixtureState[Y][X].Enabled && !this.FixtureState[Y][X].RGBInverted && X < sequence.NumberFixtures {

		if debug {
			fmt.Printf("Enable fixture number %d State on Sequence %d to true [Scanners:%d]\n", X, Y, sequence.NumberFixtures)
		}

		this.FixtureState[Y][X].Enabled = true
		this.FixtureState[Y][X].RGBInverted = false

		// Tell the sequence to turn on this fixture.
		cmd := common.Command{
			Action: common.ToggleFixtureState,
			Args: []common.Arg{
				{Name: "FixtureNumber", Value: X},
				{Name: "FixtureState", Value: true},
				{Name: "FixtureInverted", Value: false},
				{Name: "FixtureReversed", Value: false},
			},
		}
		common.SendCommandToSequence(Y, cmd, commandChannels)

		// If we're a scanner also tell the sequence shutter chaser to turn on this fixture.
		if this.SelectedType == "scanner" {
			common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)
		}

		return

	}

	// Invert scanner if we're enabled but not inverted.
	if this.FixtureState[Y][X].Enabled && !this.FixtureState[Y][X].RGBInverted && X < sequence.NumberFixtures {

		if debug {
			fmt.Printf("Invert Scanner Number %d State on Sequence %d to false\n", X, Y)
		}

		this.FixtureState[Y][X].Enabled = true
		this.FixtureState[Y][X].RGBInverted = true

		// Tell the sequence to invert this scanner.
		cmd := common.Command{
			Action: common.ToggleFixtureState,
			Args: []common.Arg{
				{Name: "FixtureNumber", Value: X},
				{Name: "FixtureState", Value: true},
				{Name: "FixtureInverted", Value: true},
				{Name: "FixtureReversed", Value: false},
			},
		}
		common.SendCommandToSequence(Y, cmd, commandChannels)

		// If we're a scanner also tell the sequence shutter chaser to invert this fixture.
		if this.SelectedType == "scanner" {
			common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)
		}

		return
	}

}
