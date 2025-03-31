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

	"github.com/dhowlett99/dmxlights/pkg/colors"
	"github.com/dhowlett99/dmxlights/pkg/common"
)

func toggleFixtureStatus(sequences []*common.Sequence, selectedFixture int, selectedSequence int, this *CurrentState, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command, updateChannels []chan common.Sequence) {

	if debug {
		fmt.Printf("Disable Fixture %d selectedSequence:%d\n", selectedFixture, selectedSequence)
		fmt.Printf("Fixture State Enabled %t  Inverted %t Reversed %t\n", sequences[selectedSequence].FixtureState[selectedFixture].Enabled, sequences[selectedSequence].FixtureState[selectedFixture].RGBInverted, sequences[selectedSequence].FixtureState[selectedFixture].ScannerPatternReversed)
	}

	sequences[selectedSequence].FixtureState = GetFixtureStatus(selectedSequence, commandChannels, updateChannels)

	// Rotate the  fixture state based on last fixture state.
	SetFixtureStatus(sequences[selectedSequence], this, selectedFixture, selectedSequence, commandChannels)

	// Show the status.
	showFixtureStatus(selectedSequence, sequences[selectedSequence], eventsForLaunchpad, guiButtons, commandChannels)
}

func EnableAllFixtures(sequenceNumber int, commandChannels []chan common.Command) {

	// Tell the sequence to invert this scanner.
	cmd := common.Command{
		Action: common.EnableAllFixtures,
	}
	common.SendCommandToSequence(sequenceNumber, cmd, commandChannels)

}

func InvertAllFixtures(sequenceNumber int, commandChannels []chan common.Command) {

	cmd := common.Command{
		Action: common.InvertAllFixtures,
	}
	common.SendCommandToSequence(sequenceNumber, cmd, commandChannels)

}

func GetFixtureStatus(selectedSequence int, commandChannels []chan common.Command, updateChannels []chan common.Sequence) map[int]common.FixtureState {

	// Get an upto date copy of the sequence.
	sequence := common.RefreshSequence(selectedSequence, commandChannels, updateChannels)

	return sequence.FixtureState
}

func SetFixtureStatus(sequence *common.Sequence, this *CurrentState, selectedFixture int, selectedSequence int, commandChannels []chan common.Command) {

	// There are three possiblities OFF, ON and INVERTED.
	if sequence.Type == "rgb" {

		// Disable fixture if we're already enabled and inverted.
		if sequence.FixtureState[selectedFixture].Enabled && sequence.FixtureState[selectedFixture].RGBInverted && selectedFixture < sequence.NumberFixtures {
			if debug {
				fmt.Printf("Disable RGB fixture Number %d State on Sequence %d to false\n", selectedFixture, selectedSequence)
			}

			newState := common.FixtureState{}
			newState.Enabled = false
			newState.RGBInverted = false
			newState.ScannerPatternReversed = false
			sequence.FixtureState[selectedFixture] = newState
			// If any fixture is not inverted turn off the global invert function.
			this.Functions[selectedSequence][common.Function7_Invert_Chase].State = false

			// Tell the sequence to invert this scanner.
			cmd := common.Command{
				Action: common.ToggleFixtureState,
				Args: []common.Arg{
					{Name: "FixtureNumber", Value: selectedFixture},
					{Name: "FixtureState", Value: sequence.FixtureState[selectedFixture].Enabled},
					{Name: "FixtureInverted", Value: sequence.FixtureState[selectedFixture].RGBInverted},
					{Name: "FixtureReversed", Value: sequence.FixtureState[selectedFixture].ScannerPatternReversed},
				},
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

			return
		}

		// Enable RGB fixture if not enabled but inverted by the global invert all enabled fixtures function key.
		if !sequence.FixtureState[selectedFixture].Enabled && sequence.FixtureState[selectedFixture].RGBInverted && selectedFixture < sequence.NumberFixtures {

			if debug {
				fmt.Printf("Enable  RGB fixture Number %d State on Sequence %d to false\n", selectedFixture, selectedSequence)
			}

			newState := common.FixtureState{}
			newState.Enabled = true
			newState.RGBInverted = false
			newState.ScannerPatternReversed = false
			sequence.FixtureState[selectedFixture] = newState

			// Tell the sequence to invert this scanner.
			cmd := common.Command{
				Action: common.ToggleFixtureState,
				Args: []common.Arg{
					{Name: "FixtureNumber", Value: selectedFixture},
					{Name: "FixtureState", Value: sequence.FixtureState[selectedFixture].Enabled},
					{Name: "FixtureInverted", Value: sequence.FixtureState[selectedFixture].RGBInverted},
					{Name: "FixtureReversed", Value: sequence.FixtureState[selectedFixture].ScannerPatternReversed},
				},
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

			return
		}

		// Enable RGB fixture if not enabled and not inverted.
		if !sequence.FixtureState[selectedFixture].Enabled && !sequence.FixtureState[selectedFixture].RGBInverted && selectedFixture < sequence.NumberFixtures {
			if debug {
				fmt.Printf("Disable RGB fixture Number %d State on Sequence %d to false\n", selectedFixture, selectedSequence)
			}

			newState := common.FixtureState{}
			newState.Enabled = true
			newState.RGBInverted = false
			newState.ScannerPatternReversed = false
			sequence.FixtureState[selectedFixture] = newState

			// Tell the sequence to invert this scanner.
			cmd := common.Command{
				Action: common.ToggleFixtureState,
				Args: []common.Arg{
					{Name: "FixtureNumber", Value: selectedFixture},
					{Name: "FixtureState", Value: sequence.FixtureState[selectedFixture].Enabled},
					{Name: "FixtureInverted", Value: sequence.FixtureState[selectedFixture].RGBInverted},
					{Name: "FixtureReversed", Value: sequence.FixtureState[selectedFixture].ScannerPatternReversed},
				},
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

			return
		}

		// Invert RGB fixture if we're enabled but not inverted.
		if sequence.FixtureState[selectedFixture].Enabled && !sequence.FixtureState[selectedFixture].RGBInverted && selectedFixture < sequence.NumberFixtures {

			if debug {
				fmt.Printf("Invert RGB Number %d State on Sequence %d to false\n", selectedFixture, selectedSequence)
			}

			newState := common.FixtureState{}
			newState.Enabled = true
			newState.RGBInverted = true
			newState.ScannerPatternReversed = false
			sequence.FixtureState[selectedFixture] = newState

			// Tell the sequence to invert this scanner.
			cmd := common.Command{
				Action: common.ToggleFixtureState,
				Args: []common.Arg{
					{Name: "FixtureNumber", Value: selectedFixture},
					{Name: "FixtureState", Value: sequence.FixtureState[selectedFixture].Enabled},
					{Name: "FixtureInverted", Value: sequence.FixtureState[selectedFixture].RGBInverted},
					{Name: "FixtureReversed", Value: sequence.FixtureState[selectedFixture].ScannerPatternReversed},
				},
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

			return
		}
	}

	// There are five possiblities OFF, ON , RGB_INVERTED, SCANNER_REVERSED and RGB_INVERTED AND SCANNER_REVERSED.
	if sequence.Type == "scanner" {

		// OOF - Disable scanner if we're already enabled and inverted and reversed.
		if sequence.FixtureState[selectedFixture].Enabled && sequence.FixtureState[selectedFixture].RGBInverted && sequence.FixtureState[selectedFixture].ScannerPatternReversed && selectedFixture < sequence.NumberFixtures {
			if debug {
				fmt.Printf("Disable scanner fixture Number %d State on Sequence %d to false\n", selectedFixture, selectedSequence)
			}

			newState := common.FixtureState{}
			newState.Enabled = false
			newState.RGBInverted = false
			newState.ScannerPatternReversed = false
			sequence.FixtureState[selectedFixture] = newState

			// Tell the sequence to invert this scanner.
			cmd := common.Command{
				Action: common.ToggleFixtureState,
				Args: []common.Arg{
					{Name: "FixtureNumber", Value: selectedFixture},
					{Name: "FixtureState", Value: sequence.FixtureState[selectedFixture].Enabled},
					{Name: "FixtureInverted", Value: sequence.FixtureState[selectedFixture].RGBInverted},
					{Name: "FixtureReversed", Value: sequence.FixtureState[selectedFixture].ScannerPatternReversed},
				},
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

			// If we're a scanner also tell the sequence shutter chaser to invert this fixture.
			if this.SelectedType == "scanner" {
				common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)
			}

			return
		}

		// ON - Enable scanner if not enabled and not inverted and not reversed.
		if !sequence.FixtureState[selectedFixture].Enabled && !sequence.FixtureState[selectedFixture].RGBInverted && !sequence.FixtureState[selectedFixture].ScannerPatternReversed && selectedFixture < sequence.NumberFixtures {
			if debug {
				fmt.Printf("Enable scanner fixture Number %d State on Sequence %d to false\n", selectedFixture, selectedSequence)
			}

			newState := common.FixtureState{}
			newState.Enabled = true
			newState.RGBInverted = false
			newState.ScannerPatternReversed = false
			sequence.FixtureState[selectedFixture] = newState

			// Tell the sequence to invert this scanner.
			cmd := common.Command{
				Action: common.ToggleFixtureState,
				Args: []common.Arg{
					{Name: "FixtureNumber", Value: selectedFixture},
					{Name: "FixtureState", Value: sequence.FixtureState[selectedFixture].Enabled},
					{Name: "FixtureInverted", Value: sequence.FixtureState[selectedFixture].RGBInverted},
					{Name: "FixtureReversed", Value: sequence.FixtureState[selectedFixture].ScannerPatternReversed},
				},
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

			// If we're a scanner also tell the sequence shutter chaser to invert this fixture.
			if this.SelectedType == "scanner" {
				common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)
			}

			return
		}

		// Invert scanner if we're enabled but not inverted and not reversed.
		if sequence.FixtureState[selectedFixture].Enabled && !sequence.FixtureState[selectedFixture].RGBInverted && !sequence.FixtureState[selectedFixture].ScannerPatternReversed && selectedFixture < sequence.NumberFixtures {

			if debug {
				fmt.Printf("Invert Scanner Number %d State on Sequence %d to false\n", selectedFixture, selectedSequence)
			}

			newState := common.FixtureState{}
			newState.Enabled = true
			newState.RGBInverted = true
			newState.ScannerPatternReversed = false
			sequence.FixtureState[selectedFixture] = newState

			// Tell the sequence to invert this scanner.
			cmd := common.Command{
				Action: common.ToggleFixtureState,
				Args: []common.Arg{
					{Name: "FixtureNumber", Value: selectedFixture},
					{Name: "FixtureState", Value: sequence.FixtureState[selectedFixture].Enabled},
					{Name: "FixtureInverted", Value: sequence.FixtureState[selectedFixture].RGBInverted},
					{Name: "FixtureReversed", Value: sequence.FixtureState[selectedFixture].ScannerPatternReversed},
				},
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

			// If we're a scanner also tell the sequence shutter chaser to invert this fixture.
			if this.SelectedType == "scanner" {
				common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)
			}
			return
		}

		// reverse scanner if we're enabled , inverted and not reversed.
		if sequence.FixtureState[selectedFixture].Enabled && sequence.FixtureState[selectedFixture].RGBInverted && !sequence.FixtureState[selectedFixture].ScannerPatternReversed && selectedFixture < sequence.NumberFixtures {

			if debug {
				fmt.Printf("Invert Scanner Number %d State on Sequence %d to false\n", selectedFixture, selectedSequence)
			}

			newState := common.FixtureState{}
			newState.Enabled = true
			newState.RGBInverted = false
			newState.ScannerPatternReversed = true
			sequence.FixtureState[selectedFixture] = newState

			// Tell the sequence to invert this scanner.
			cmd := common.Command{
				Action: common.ToggleFixtureState,
				Args: []common.Arg{
					{Name: "FixtureNumber", Value: selectedFixture},
					{Name: "FixtureState", Value: sequence.FixtureState[selectedFixture].Enabled},
					{Name: "FixtureInverted", Value: sequence.FixtureState[selectedFixture].RGBInverted},
					{Name: "FixtureReversed", Value: sequence.FixtureState[selectedFixture].ScannerPatternReversed},
				},
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

			// If we're a scanner also tell the sequence shutter chaser to invert this fixture.
			if this.SelectedType == "scanner" {
				common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)
			}
			return
		}

		// reverse and invert scanner if we're enabled , inverted and reversed.
		if sequence.FixtureState[selectedFixture].Enabled && !sequence.FixtureState[selectedFixture].RGBInverted && sequence.FixtureState[selectedFixture].ScannerPatternReversed && selectedFixture < sequence.NumberFixtures {

			if debug {
				fmt.Printf("Reverse and Invert Scanner Number %d State on Sequence %d to false\n", selectedFixture, selectedSequence)
			}
			newState := common.FixtureState{}
			newState.Enabled = true
			newState.RGBInverted = true
			newState.ScannerPatternReversed = true
			sequence.FixtureState[selectedFixture] = newState

			// Tell the sequence to invert this scanner.
			cmd := common.Command{
				Action: common.ToggleFixtureState,
				Args: []common.Arg{
					{Name: "FixtureNumber", Value: selectedFixture},
					{Name: "FixtureState", Value: sequence.FixtureState[selectedFixture].Enabled},
					{Name: "FixtureInverted", Value: sequence.FixtureState[selectedFixture].RGBInverted},
					{Name: "FixtureReversed", Value: sequence.FixtureState[selectedFixture].ScannerPatternReversed},
				},
			}
			common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

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
func showFixtureStatus(selectedSequence int, sequence *common.Sequence, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command) {

	if debug {
		fmt.Printf("Show Fixture Status for sequence %d number of fixtures %d\n", sequence.Number, sequence.NumberFixtures)
	}

	common.HideSequence(selectedSequence, commandChannels)

	for fixtureNumber := 0; fixtureNumber < sequence.NumberFixtures; fixtureNumber++ {

		// in version 3.0 We are not supporting fixture status for any more than 8 fixtures.
		if fixtureNumber > 7 {
			return
		}

		if debug {
			fmt.Printf("Sequence %d: Fixture %d Enabled %t Inverted %t\n", sequence.Number, fixtureNumber, sequence.FixtureState[fixtureNumber].Enabled, sequence.FixtureState[fixtureNumber].RGBInverted)
		}

		// Enabled but not inverted then On and green.
		if sequence.FixtureState[fixtureNumber].Enabled && !sequence.FixtureState[fixtureNumber].RGBInverted {
			common.LightLamp(common.Button{X: fixtureNumber, Y: sequence.Number}, colors.Green, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)
			common.LabelButton(fixtureNumber, sequence.Number, "On", guiButtons)
		}

		// Enabled and inverted then Invert and puple. Not reversed
		if sequence.FixtureState[fixtureNumber].Enabled && sequence.FixtureState[fixtureNumber].RGBInverted && !sequence.FixtureState[fixtureNumber].ScannerPatternReversed {
			common.LightLamp(common.Button{X: fixtureNumber, Y: sequence.Number}, colors.Purple, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)
			common.LabelButton(fixtureNumber, sequence.Number, "Invert", guiButtons)
		}

		// Enabled not inverted but revesed then reverse and yellow.
		if sequence.FixtureState[fixtureNumber].Enabled && !sequence.FixtureState[fixtureNumber].RGBInverted && sequence.FixtureState[fixtureNumber].ScannerPatternReversed {
			common.LightLamp(common.Button{X: fixtureNumber, Y: sequence.Number}, colors.Yellow, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)
			common.LabelButton(fixtureNumber, sequence.Number, "Reversed", guiButtons)
		}

		// Enabled  inverted and revesed then reverse and white.
		if sequence.FixtureState[fixtureNumber].Enabled && sequence.FixtureState[fixtureNumber].RGBInverted && sequence.FixtureState[fixtureNumber].ScannerPatternReversed {
			common.LightLamp(common.Button{X: fixtureNumber, Y: sequence.Number}, colors.White, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)
			common.LabelButton(fixtureNumber, sequence.Number, "Invert & Reversed", guiButtons)
		}

		// Not enabled and not inverted then off and blue.
		if !sequence.FixtureState[fixtureNumber].Enabled {
			common.LightLamp(common.Button{X: fixtureNumber, Y: sequence.Number}, colors.Red, 255, eventsForLaunchpad, guiButtons)
			common.LabelButton(fixtureNumber, sequence.Number, "Off", guiButtons)
		}

	}
}
