// Copyright (C) 2022,2023,2024 dhowlett99.
// This is the dmxlights main sequencers calculate position front end functions.
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
	"github.com/dhowlett99/dmxlights/pkg/position"
)

func calculateRGBPositions(sequence *common.Sequence, steps []common.Step) map[int]common.Position {

	if debug {
		fmt.Printf("calculateRGBPositions: \n")
	}

	// Calculate fade curve values.
	common.CalculateFadeValues(sequence)
	// Calulate positions for each RGB fixture.
	sequence.Optimisation = true
	var numberSteps int
	fadeColors, totalNumberOfSteps := position.CalculatePositions(steps, *sequence, common.IS_RGB)
	rgbPositions, numberSteps := position.AssemblePositions(fadeColors, sequence.NumberFixtures, totalNumberOfSteps, sequence.FixtureState, sequence.Optimisation)
	sequence.NumberSteps = numberSteps

	return rgbPositions
}

func calculateScannerPositions(sequence *common.Sequence, steps []common.Step) map[int]map[int]common.Position {

	if debug {
		fmt.Printf("calculateScannerPositions: \n")
	}

	if debug {
		fmt.Printf("Scanner Steps\n")
		for stepNumber, step := range sequence.ScannerSteps {
			fmt.Printf("Scanner Steps %+v\n", stepNumber)
			for _, fixture := range step.Fixtures {
				fmt.Printf("Fixture %+v\n", fixture)
			}
		}
	}

	scannerPositions := make(map[int]map[int]common.Position, sequence.NumberFixtures)

	for fixture := 0; fixture < sequence.NumberFixtures; fixture++ {
		var positions map[int]common.Position
		// We're playing out the scanner positions, so we won't need curve values.
		sequence.FadeUp = []int{255}
		sequence.FadeDown = []int{0}
		// Turn on optimasation.
		sequence.Optimisation = true

		// Pass through the inverted / reverse flag.
		sequence.ScannerReverse = sequence.FixtureState[fixture].ScannerPatternReversed
		// Calulate positions for each scanner fixture.
		fadeColors, totalNumberOfSteps := position.CalculatePositions(steps, *sequence, common.IS_SCANNER)
		positions, numberSteps := position.AssemblePositions(fadeColors, sequence.NumberFixtures, totalNumberOfSteps, sequence.FixtureState, sequence.Optimisation)
		sequence.NumberSteps = numberSteps

		// Setup positions for each scanner. This is so we can shift the patterns on each scannner.
		scannerPositions[fixture] = make(map[int]common.Position, 9)
		for positionNumber, position := range positions {
			scannerPositions[fixture][positionNumber] = position
		}
	}

	return scannerPositions
}
