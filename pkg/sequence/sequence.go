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

		// Process any commands.
		processCommands(mySequenceNumber, &sequence, channels, switchChannels, fixtureStepChannels, eventsForLauchpad, guiButtons)

		// Sequence in normal running chase mode.
		if sequence.Chase && sequence.Run && !sequence.Static && !sequence.StartFlood {

			// Update the steps.
			steps = generateSteps(steps, availablePatterns, &sequence, soundConfig, fixturesConfig)

			if debug {
				fmt.Printf("%d: Begin CHASE Sequence type %s label %s Running %t Colors %+v NumberSteps=%d \n", mySequenceNumber, sequence.Type, sequence.Label, sequence.Run, sequence.SequenceColors, sequence.NumberSteps)
			}

			// Calculate positions from steps.  Soeed, shift, size and fade can be addjusted in this loop.
			rgbPositions, scannerPositions := calculatePositions(&sequence, steps)

			// This is the inner loop where the sequence runs.
			// Run through the steps in the sequence.
			// Remember every step contains infomation for all the fixtures in this sequence group.
			for step := 0; step < sequence.NumberSteps; step++ {

				// This is were we set the speed of the sequence to current speed.
				speed := sequence.CurrentSpeed / 10
				if sequence.Type == "scanner" {
					speed = sequence.CurrentSpeed / 5 // Slow the scanners down.
				}

				// Listen for any commands during chase so inside sequence steps loop, time out for the next step at the speed of the chase.
				// or additionally timeout when get a beat that also triggers the next step.
				sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, speed, sequence, channels, fixturesConfig)
				if !sequence.Run || sequence.Clear || sequence.StartFlood || sequence.StopFlood ||
					sequence.Static || sequence.UpdateShift || sequence.StartPattern || sequence.UpdateColors || sequence.UpdateSize {
					for fixtureNumber := 0; fixtureNumber < sequence.NumberFixtures; fixtureNumber++ {
						clearFixture(fixtureNumber, fixtureStepChannels)
					}
					if debug {
						fmt.Printf("%d: Break\n", mySequenceNumber)
						fmt.Printf("%d: Run %t \n", mySequenceNumber, sequence.Run)
						fmt.Printf("%d: Clear %t \n", mySequenceNumber, sequence.Clear)
						fmt.Printf("%d: StartFlood %t\n", mySequenceNumber, sequence.StartFlood)
						fmt.Printf("%d: StopFlood %t\n", mySequenceNumber, sequence.StopFlood)
						fmt.Printf("%d: Statics %t\n", mySequenceNumber, sequence.Static)
						fmt.Printf("%d: UpdateShift %t\n", mySequenceNumber, sequence.UpdateShift)
						fmt.Printf("%d: StartPattern %t\n", mySequenceNumber, sequence.StartPattern)
						fmt.Printf("%d: UpdateColors %t\n", mySequenceNumber, sequence.UpdateColors)
						fmt.Printf("%d: UpdateSize %t\n", mySequenceNumber, sequence.UpdateSize)
					}
					// Break out of the step loop to process commands.
					break
				}

				for fixtureNumber := 0; fixtureNumber < sequence.NumberFixtures; fixtureNumber++ {
					playStep(&sequence, step, fixtureNumber, rgbPositions, scannerPositions, fixtureStepChannels)
				}
			}
		} else {
			if debug {
				fmt.Printf("%d: Start Listen for commands\n", mySequenceNumber)
			}

			// This is where we wait for command when the sequence isn't running.
			// Check for any waiting commands. Setting a large timeout means that we only return when we hava a command.
			sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, 50*time.Hour, sequence, channels, fixturesConfig)
		}
	}
}
