// Copyright (C) 2022,2023, 2024, 2025 dhowlett99.
// This is part of the dmxlights sequencer, this file holds
// some helper function for setting up fixtures in a sequence.
//
// Implemented and depends on fyne.io
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

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
	"github.com/dhowlett99/dmxlights/pkg/pattern"
	"github.com/dhowlett99/dmxlights/pkg/sound"
	"github.com/oliread/usbdmx/ft232"
)

func LoadNewFixtures(sequence *common.Sequence,
	fixtureStepChannels []chan common.FixtureCommand,
	eventsForLaunchpad chan common.ALight,
	guiButtons chan common.ALight,
	switchChannels []common.SwitchChannel,
	soundTriggers []*common.Trigger,
	soundConfig *sound.SoundConfig,
	dmxController *ft232.DMXController,
	fixturesConfig *fixture.Fixtures,
	dmxInterfacePresent bool) []chan common.FixtureCommand {

	// Stop existing Fixture threads.
	StopFixtureReceivers(fixtureStepChannels)

	// Wait for fixture threads to stop.
	time.Sleep(25 * time.Millisecond)

	// Count the number of fixtures for this sequence.
	sequence.NumberFixtures = CountNumberOfFixtures(*sequence, fixturesConfig)

	// Create a new set of fixture command channels.
	fixtureStepChannels = CreateFixtureChannels(sequence.NumberFixtures, fixturesConfig)

	// Because the number of fixtures may have changes reload the patterns.
	sequence.RGBAvailablePatterns = LoadAvailablePatterns(*sequence, fixturesConfig)

	// Set the fixtures states to enables and not inverted.
	sequence.FixtureState = SetDefaultFixtureState(sequence.NumberFixtures)

	// Now create a thread for each one of the new fixtures.
	CreateFixtureReceiverThreads(sequence.NumberFixtures, fixtureStepChannels, eventsForLaunchpad, guiButtons, switchChannels, soundTriggers, soundConfig, dmxController, fixturesConfig, dmxInterfacePresent)

	return fixtureStepChannels
}

// StopFixtureReceivers sends a message to all the current fixtures
// and asks the fixture thread to stop.
// Takes
func StopFixtureReceivers(fixtureStepChannels []chan common.FixtureCommand) {

	if debug {
		fmt.Printf("StopFixtureReceivers\n")
	}

	// Prepare a message to be sent to the fixtures in the sequence.
	command := common.FixtureCommand{
		Stop: true,
	}

	// Now tell all the fixtures what they need to do.
	fixture.SendToAllFixtures(fixtureStepChannels, command)

}

// CountNumberOfFixtures counts the number of fixturs in this sequence.
// Takes the sequence number and the fixture config.
// Returns a number of fixtures available for this sequence.
func CountNumberOfFixtures(sequence common.Sequence, fixturesConfig *fixture.Fixtures) int {

	sequence.NumberFixtures = fixture.HowManyFixturesInGroup(sequence.Number, fixturesConfig)
	if debug {
		fmt.Printf("%d: Number of Patterms %d\n", sequence.Number, sequence.NumberFixtures)
	}

	return sequence.NumberFixtures
}

// LoadAvailablePatterns configures a set of patterns for an RGB sequence.
// Takes a sequence and fixtures config.
// Returns an array of available patterns for this RGB sequence.
func LoadAvailablePatterns(sequence common.Sequence, fixturesConfig *fixture.Fixtures) []common.Pattern {

	RGBAvailablePatterns := pattern.MakePatterns(sequence.NumberFixtures)
	if debug {
		fmt.Printf("%d: Number of Patterms %d\n", sequence.Number, len(RGBAvailablePatterns))
	}
	return RGBAvailablePatterns
}

// Create channels used for stepping the fixture threads for this sequnece.
// Takes the number of fixtures and the fixtures config.
// Returns an array of channels to commnicate with the fixtures for this sequence.
func CreateFixtureChannels(numberFixtures int, fixturesConfig *fixture.Fixtures) []chan common.FixtureCommand {

	fixtureStepChannels := []chan common.FixtureCommand{}
	for fixtureNumber := 0; fixtureNumber < numberFixtures; fixtureNumber++ {
		fixtureStepChannel := make(chan common.FixtureCommand)
		fixtureStepChannels = append(fixtureStepChannels, fixtureStepChannel)
	}
	if debug {
		fmt.Printf("CreateFixtureChannels: Number of Channels Created %d\n", len(fixtureStepChannels))
	}
	return fixtureStepChannels

}

// SetDefaultFixtureState sets the default fixture state i.e every fixture enabled and not inverted.
// Takes the number of fixtures.
// Returns an array of type fixture.State.
func SetDefaultFixtureState(numberFixtures int) map[int]common.FixtureState {

	if debug {
		fmt.Printf("SetDefaultFixtureState\n")
	}

	fixtureState := make(map[int]common.FixtureState, numberFixtures)

	// Set the default fixture state.
	for fixtureNumber := 0; fixtureNumber < numberFixtures; fixtureNumber++ {
		newFixture := common.FixtureState{}
		newFixture.Enabled = true
		newFixture.RGBInverted = false
		newFixture.ScannerPatternReversed = false
		fixtureState[fixtureNumber] = newFixture
	}

	return fixtureState
}

// Create a fixture thread for each fixture.
func CreateFixtureReceiverThreads(numberFixtures int,
	fixtureStepChannels []chan common.FixtureCommand,
	eventsForLaunchpad chan common.ALight,
	guiButtons chan common.ALight,
	switchChannels []common.SwitchChannel,
	soundTriggers []*common.Trigger,
	soundConfig *sound.SoundConfig,
	dmxController *ft232.DMXController,
	fixturesConfig *fixture.Fixtures,
	dmxInterfacePresent bool) int {

	if debug {
		fmt.Printf("CreateFixtureReceiverThreads\n")
	}

	for fixtureNumber := 0; fixtureNumber < numberFixtures; fixtureNumber++ {
		if debug {
			fmt.Printf("CreateFixtureReceiver %d\n", fixtureNumber)
		}
		go fixture.FixtureReceiver(fixtureNumber, fixtureStepChannels[fixtureNumber], eventsForLaunchpad, guiButtons, switchChannels, soundTriggers, soundConfig, dmxController, fixturesConfig, dmxInterfacePresent)
	}

	return numberFixtures
}
