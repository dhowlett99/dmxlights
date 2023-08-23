// Copyright (C) 2022, 2023 dhowlett99.
// These status functions implenents the fixture state, to enable, disable
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

	// There are three possiblities OFF, ON and INVERTED.
	if sequence.Type == "rgb" {

		// Disable fixture if we're already enabled and inverted.
		if this.FixtureState[Y][X].Enabled && this.FixtureState[Y][X].RGBInverted && X < sequence.NumberFixtures {
			if debug {
				fmt.Printf("Disable RGB fixture Number %d State on Sequence %d to false\n", X, Y)
			}

			this.FixtureState[Y][X].Enabled = false
			this.FixtureState[Y][X].RGBInverted = false

			// Tell the sequence to invert this scanner.
			cmd := common.Command{
				Action: common.ToggleFixtureState,
				Args: []common.Arg{
					{Name: "FixtureNumber", Value: X},
					{Name: "FixtureState", Value: this.FixtureState[Y][X].Enabled},
					{Name: "FixtureInverted", Value: this.FixtureState[Y][X].RGBInverted},
					{Name: "FixtureReversed", Value: this.FixtureState[Y][X].ScannerPatternReversed},
				},
			}
			common.SendCommandToSequence(Y, cmd, commandChannels)

			// If we're a scanner also tell the sequence shutter chaser to invert this fixture.
			if this.SelectedType == "scanner" {
				common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)
			}

			return
		}

		// Enable scanner if not enabled and not inverted.
		if !this.FixtureState[Y][X].Enabled && !this.FixtureState[Y][X].RGBInverted && X < sequence.NumberFixtures {
			if debug {
				fmt.Printf("Disable RGB fixture Number %d State on Sequence %d to false\n", X, Y)
			}

			this.FixtureState[Y][X].Enabled = true
			this.FixtureState[Y][X].RGBInverted = false

			// Tell the sequence to invert this scanner.
			cmd := common.Command{
				Action: common.ToggleFixtureState,
				Args: []common.Arg{
					{Name: "FixtureNumber", Value: X},
					{Name: "FixtureState", Value: this.FixtureState[Y][X].Enabled},
					{Name: "FixtureInverted", Value: this.FixtureState[Y][X].RGBInverted},
					{Name: "FixtureReversed", Value: this.FixtureState[Y][X].ScannerPatternReversed},
				},
			}
			common.SendCommandToSequence(Y, cmd, commandChannels)

			// If we're a scanner also tell the sequence shutter chaser to invert this fixture.
			if this.SelectedType == "scanner" {
				common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)
			}

			return
		}

		// Invert scanner if we're enabled but not inverted.
		if this.FixtureState[Y][X].Enabled && !this.FixtureState[Y][X].RGBInverted && X < sequence.NumberFixtures {

			if debug {
				fmt.Printf("Invert RGB Number %d State on Sequence %d to false\n", X, Y)
			}

			this.FixtureState[Y][X].Enabled = true
			this.FixtureState[Y][X].RGBInverted = true

			// Tell the sequence to invert this scanner.
			cmd := common.Command{
				Action: common.ToggleFixtureState,
				Args: []common.Arg{
					{Name: "FixtureNumber", Value: X},
					{Name: "FixtureState", Value: this.FixtureState[Y][X].Enabled},
					{Name: "FixtureInverted", Value: this.FixtureState[Y][X].RGBInverted},
					{Name: "FixtureReversed", Value: this.FixtureState[Y][X].ScannerPatternReversed},
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

	// There are five possiblities OFF, ON , RGB_INVERTED, SCANNER_REVERSED and RGB_INVERTED AND SCANNER_REVERSED.
	if sequence.Type == "scanner" {

		// OOF - Disable fixture if we're already enabled and inverted and reversed.
		if this.FixtureState[Y][X].Enabled && this.FixtureState[Y][X].RGBInverted && this.FixtureState[Y][X].ScannerPatternReversed && X < sequence.NumberFixtures {
			if debug {
				fmt.Printf("Disable scanner fixture Number %d State on Sequence %d to false\n", X, Y)
			}

			this.FixtureState[Y][X].Enabled = false
			this.FixtureState[Y][X].RGBInverted = false
			this.FixtureState[Y][X].ScannerPatternReversed = false

			// Tell the sequence to invert this scanner.
			cmd := common.Command{
				Action: common.ToggleFixtureState,
				Args: []common.Arg{
					{Name: "FixtureNumber", Value: X},
					{Name: "FixtureState", Value: this.FixtureState[Y][X].Enabled},
					{Name: "FixtureInverted", Value: this.FixtureState[Y][X].RGBInverted},
					{Name: "FixtureReversed", Value: this.FixtureState[Y][X].ScannerPatternReversed},
				},
			}
			common.SendCommandToSequence(Y, cmd, commandChannels)

			// If we're a scanner also tell the sequence shutter chaser to invert this fixture.
			if this.SelectedType == "scanner" {
				common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)
			}

			return
		}

		// ON - Enable scanner if not enabled and not inverted and not reversed.
		if !this.FixtureState[Y][X].Enabled && !this.FixtureState[Y][X].RGBInverted && !this.FixtureState[Y][X].ScannerPatternReversed && X < sequence.NumberFixtures {
			if debug {
				fmt.Printf("Enable scanner fixture Number %d State on Sequence %d to false\n", X, Y)
			}

			this.FixtureState[Y][X].Enabled = true
			this.FixtureState[Y][X].RGBInverted = false
			this.FixtureState[Y][X].ScannerPatternReversed = false

			// Tell the sequence to invert this scanner.
			cmd := common.Command{
				Action: common.ToggleFixtureState,
				Args: []common.Arg{
					{Name: "FixtureNumber", Value: X},
					{Name: "FixtureState", Value: this.FixtureState[Y][X].Enabled},
					{Name: "FixtureInverted", Value: this.FixtureState[Y][X].RGBInverted},
					{Name: "FixtureReversed", Value: this.FixtureState[Y][X].ScannerPatternReversed},
				},
			}
			common.SendCommandToSequence(Y, cmd, commandChannels)

			// If we're a scanner also tell the sequence shutter chaser to invert this fixture.
			if this.SelectedType == "scanner" {
				common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)
			}

			return
		}

		// Invert scanner if we're enabled but not inverted and not reversed.
		if this.FixtureState[Y][X].Enabled && !this.FixtureState[Y][X].RGBInverted && !this.FixtureState[Y][X].ScannerPatternReversed && X < sequence.NumberFixtures {

			if debug {
				fmt.Printf("Invert Scanner Number %d State on Sequence %d to false\n", X, Y)
			}

			this.FixtureState[Y][X].Enabled = true
			this.FixtureState[Y][X].RGBInverted = true
			this.FixtureState[Y][X].ScannerPatternReversed = false

			// Tell the sequence to invert this scanner.
			cmd := common.Command{
				Action: common.ToggleFixtureState,
				Args: []common.Arg{
					{Name: "FixtureNumber", Value: X},
					{Name: "FixtureState", Value: this.FixtureState[Y][X].Enabled},
					{Name: "FixtureInverted", Value: this.FixtureState[Y][X].RGBInverted},
					{Name: "FixtureReversed", Value: this.FixtureState[Y][X].ScannerPatternReversed},
				},
			}
			common.SendCommandToSequence(Y, cmd, commandChannels)

			// If we're a scanner also tell the sequence shutter chaser to invert this fixture.
			if this.SelectedType == "scanner" {
				common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)
			}
			return
		}

		// reverse scanner if we're enabled , inverted and not reversed.
		if this.FixtureState[Y][X].Enabled && this.FixtureState[Y][X].RGBInverted && !this.FixtureState[Y][X].ScannerPatternReversed && X < sequence.NumberFixtures {

			if debug {
				fmt.Printf("Invert Scanner Number %d State on Sequence %d to false\n", X, Y)
			}

			this.FixtureState[Y][X].Enabled = true
			this.FixtureState[Y][X].RGBInverted = false
			this.FixtureState[Y][X].ScannerPatternReversed = true

			// Tell the sequence to invert this scanner.
			cmd := common.Command{
				Action: common.ToggleFixtureState,
				Args: []common.Arg{
					{Name: "FixtureNumber", Value: X},
					{Name: "FixtureState", Value: this.FixtureState[Y][X].Enabled},
					{Name: "FixtureInverted", Value: this.FixtureState[Y][X].RGBInverted},
					{Name: "FixtureReversed", Value: this.FixtureState[Y][X].ScannerPatternReversed},
				},
			}
			common.SendCommandToSequence(Y, cmd, commandChannels)

			// If we're a scanner also tell the sequence shutter chaser to invert this fixture.
			if this.SelectedType == "scanner" {
				common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)
			}
			return
		}

		// reverse and invert scanner if we're enabled , inverted and reversed.
		if this.FixtureState[Y][X].Enabled && !this.FixtureState[Y][X].RGBInverted && this.FixtureState[Y][X].ScannerPatternReversed && X < sequence.NumberFixtures {

			if debug {
				fmt.Printf("Reverse and Invert Scanner Number %d State on Sequence %d to false\n", X, Y)
			}

			this.FixtureState[Y][X].Enabled = true
			this.FixtureState[Y][X].RGBInverted = true
			this.FixtureState[Y][X].ScannerPatternReversed = true

			// Tell the sequence to invert this scanner.
			cmd := common.Command{
				Action: common.ToggleFixtureState,
				Args: []common.Arg{
					{Name: "FixtureNumber", Value: X},
					{Name: "FixtureState", Value: this.FixtureState[Y][X].Enabled},
					{Name: "FixtureInverted", Value: this.FixtureState[Y][X].RGBInverted},
					{Name: "FixtureReversed", Value: this.FixtureState[Y][X].ScannerPatternReversed},
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
}

// Show Scanner status - Dim White is disabled, White is enabled.
// Uses the >-this<- representation of the fixture status. Not actual sequences which are stored in the threads below us.
func showFixtureStatus(selectedSequence int, sequenceNumber int, NumberFixtures int, this *CurrentState, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command) {

	if debug {
		fmt.Printf("Show Scanner Status for sequence %d number of scanners %d\n", sequenceNumber, NumberFixtures)
	}

	common.HideSequence(selectedSequence, commandChannels)

	for fixtureNumber := 0; fixtureNumber < NumberFixtures; fixtureNumber++ {

		if debug {
			fmt.Printf("%d: Scanner %d Enabled %t Inverted %t\n", sequenceNumber, fixtureNumber, this.FixtureState[sequenceNumber][fixtureNumber].Enabled, this.FixtureState[sequenceNumber][fixtureNumber].RGBInverted)
		}

		// Enabled but not inverted then On and green.
		if this.FixtureState[sequenceNumber][fixtureNumber].Enabled && !this.FixtureState[sequenceNumber][fixtureNumber].RGBInverted {
			common.LightLamp(common.ALight{X: fixtureNumber, Y: sequenceNumber, Brightness: this.MasterBrightness, Red: 0, Green: 255, Blue: 0}, eventsForLaunchpad, guiButtons)
			common.LabelButton(fixtureNumber, sequenceNumber, "On", guiButtons)
		}

		// Enabled and inverted then Invert and puple. Not reversed
		if this.FixtureState[sequenceNumber][fixtureNumber].Enabled && this.FixtureState[sequenceNumber][fixtureNumber].RGBInverted && !this.FixtureState[sequenceNumber][fixtureNumber].ScannerPatternReversed {
			common.LightLamp(common.ALight{X: fixtureNumber, Y: sequenceNumber, Brightness: this.MasterBrightness, Red: 255, Green: 0, Blue: 255}, eventsForLaunchpad, guiButtons)
			common.LabelButton(fixtureNumber, sequenceNumber, "Invert", guiButtons)
		}

		// Enabled not inverted but revesed then reverse and yellow.
		if this.FixtureState[sequenceNumber][fixtureNumber].Enabled && !this.FixtureState[sequenceNumber][fixtureNumber].RGBInverted && this.FixtureState[sequenceNumber][fixtureNumber].ScannerPatternReversed {
			common.LightLamp(common.ALight{X: fixtureNumber, Y: sequenceNumber, Brightness: this.MasterBrightness, Red: 255, Green: 255, Blue: 0}, eventsForLaunchpad, guiButtons)
			common.LabelButton(fixtureNumber, sequenceNumber, "Reversed", guiButtons)
		}

		// Enabled  inverted and revesed then reverse and white.
		if this.FixtureState[sequenceNumber][fixtureNumber].Enabled && this.FixtureState[sequenceNumber][fixtureNumber].RGBInverted && this.FixtureState[sequenceNumber][fixtureNumber].ScannerPatternReversed {
			common.LightLamp(common.ALight{X: fixtureNumber, Y: sequenceNumber, Brightness: this.MasterBrightness, Red: 255, Green: 255, Blue: 255}, eventsForLaunchpad, guiButtons)
			common.LabelButton(fixtureNumber, sequenceNumber, "Invert & Reversed", guiButtons)
		}

		// Not enabled and not inverted then off and blue.
		if !this.FixtureState[sequenceNumber][fixtureNumber].Enabled {
			common.LightLamp(common.ALight{X: fixtureNumber, Y: sequenceNumber, Brightness: this.MasterBrightness, Red: 255, Green: 0, Blue: 0}, eventsForLaunchpad, guiButtons)
			common.LabelButton(fixtureNumber, sequenceNumber, "Off", guiButtons)
		}

	}
}
