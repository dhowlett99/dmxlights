// Copyright (C) 2022,2023 dhowlett99.
// This is the dmxlight chaser used to chase the scanner lamps.
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
	"strconv"
	"time"

	"github.com/dhowlett99/dmxlights/pkg/commands"
	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
	"github.com/dhowlett99/dmxlights/pkg/pattern"
	"github.com/dhowlett99/dmxlights/pkg/position"
	"github.com/dhowlett99/dmxlights/pkg/sound"
	"github.com/oliread/usbdmx/ft232"
)

func NewChaser(mySequenceNumber int, dmxController *ft232.DMXController, fixturesConfig *fixture.Fixtures, channels common.Channels, soundConfig *sound.SoundConfig, dmxInterfacePresent bool) common.Sequence {

	fmt.Printf("NewMultiFixureSequencer: created\n")

	// Start of chase.
	sequence := common.Sequence{}
	sequence.Type = "scanner"
	sequence.ChaserSpeed = 100 * time.Millisecond
	sequence.ScannerChase = false
	sequence.NumberFixtures = 8
	sequence.RGBCoordinates = 20
	sequence.RGBFade = 1
	sequence.RGBSize = 1
	sequence.RGBShift = 1
	sequence.EnabledNumberFixtures = 3
	sequence.Optimisation = false
	sequence.ScannerInvert = false
	sequence.Bounce = false
	sequence.RGBInvert = false

	// Make functions for this chaser sequences.
	for function := 0; function < 8; function++ {
		newFunction := common.Function{
			Name:           strconv.Itoa(function),
			SequenceNumber: mySequenceNumber,
			Number:         function,
			State:          false,
			Label:          sequence.GuiFunctionLabels[function],
		}
		sequence.Functions = append(sequence.Functions, newFunction)
	}

	//sequence.Functions[common.Function7_Invert_Chase].State = false

	sequence.ScannerState = map[int]common.ScannerState{
		0: {
			Enabled: true,
		},
		1: {
			Enabled: true,
		},
		2: {
			Enabled: true,
		},
		3: {
			Enabled: true,
		},
		4: {
			Enabled: true,
		},
		5: {
			Enabled: true,
		},
		6: {
			Enabled: true,
		},
		7: {
			Enabled: true,
		},
	}

	// Set the chase RGB steps used to chase the shutter.
	scannerChasePattern := pattern.GenerateStandardChasePatterm(sequence.NumberFixtures, sequence.ScannerState)
	sequence.Steps = scannerChasePattern.Steps

	// Calculate fade curve values. The number of Shutter (RGB) steps has to match the number of scanner steps.
	sequence.FadeUpAndDown, sequence.FadeDownAndUp = common.CalculateFadeValues(sequence.RGBCoordinates, sequence.RGBFade, sequence.RGBSize)

	return sequence
}

func StartChaser(sequence common.Sequence,
	mySequenceNumber int,
	eventsForLauchpad chan common.ALight,
	guiButtons chan common.ALight,
	dmxController *ft232.DMXController,
	fixturesConfig *fixture.Fixtures,
	channels common.Channels,
	switchChannels map[int]common.SwitchChannel,
	soundConfig *sound.SoundConfig,
	dmxInterfacePresent bool) {

	fmt.Printf("NewMultiFixureSequencer: started\n")

	// Create a new thread to service the commands.
	go func(sequence common.Sequence) {
		// Calulate positions for each Scanner Shutter.
		shutterPositions, _ := position.CalculatePositions(sequence)

		for {

			// Check for any waiting commands.
			sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, 10*time.Hour, sequence, channels, fixturesConfig)
			if sequence.ScannerChase {
				// Run through the steps in the sequence.
				// Remember every step contains infomation for all the fixtures in this group.
				for step := 0; step < len(shutterPositions); step++ {

					// Play out fixture to DMX channels.
					position := shutterPositions[step]

					fixtures := position.Fixtures

					for fixtureNumber := 0; fixtureNumber < len(position.Fixtures); fixtureNumber++ {

						myfixture := fixtures[fixtureNumber]

						scannerFixturesSequenceNumber := 2 // Scanner sequence.

						masterChannel, err := fixture.FindChannelNumberByName("Master", fixtureNumber, scannerFixturesSequenceNumber, fixturesConfig)
						if err != nil {
							fmt.Printf("StartChaser master: %s,", err)
							return
						}

						// Find the fixtures details.

						fixtureAddress, err := fixture.FindFixtureAddressByGroupAndNumber(scannerFixturesSequenceNumber, fixtureNumber, fixturesConfig)
						if err != nil {
							fmt.Printf("StartChaser: error %s\n", err.Error())
							return
						}
						//fmt.Printf("Fixture %d Set Master Address %d to Value %d\n", fixtureNumber, fixtureAddress+int16(masterChannel), myfixture.Brightness)
						fixture.SetChannel(fixtureAddress+int16(masterChannel), byte(myfixture.Brightness), dmxController, dmxInterfacePresent)
					}
					// Check for any waiting commands.
					sequence = commands.ListenCommandChannelAndWait(mySequenceNumber, sequence.ChaserSpeed, sequence, channels, fixturesConfig)
					if !sequence.ScannerChase {
						break
					}
				}
			}
		}
	}(sequence)
}
