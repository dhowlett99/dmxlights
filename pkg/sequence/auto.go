// Copyright (C) 2022,2023,2024 dhowlett99.
// This is the dmxlights main sequencers automatic change functions.
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
	"image/color"

	"github.com/dhowlett99/dmxlights/pkg/common"
)

func rgbAutoColors(sequence *common.Sequence, steps []common.Step) []common.Step {

	if debug {
		fmt.Printf("rgbAutoColors: \n")
	}

	// Set the color.
	sequence.SequenceColors = []color.RGBA{sequence.RGBAvailableColors[sequence.RGBColor].Color}

	// Increment the color.
	sequence.RGBColor++
	if sequence.RGBColor > 7 {
		sequence.RGBColor = 0
	}
	fmt.Printf("sequence.RGBColor: %d Color %+v\n", sequence.RGBColor, sequence.SequenceColors)

	// Now replace the color in the steps.
	steps = replaceRGBcolorsInSteps(steps, sequence.SequenceColors)

	return steps
}

func rgbAutoPattern(sequence *common.Sequence, availablePatterns map[int]common.Pattern) {

	if debug {
		fmt.Printf("rgbAutoPattern: \n")
	}

	for patternNumber, pattern := range availablePatterns {
		if pattern.Number == sequence.SelectedPattern {
			sequence.Pattern.Number = patternNumber
			if debug {
				fmt.Printf(">>>> I AM PATTEN %d\n", patternNumber)
			}
			break
		}
	}
	sequence.SelectedPattern++
	if sequence.SelectedPattern > len(availablePatterns) {
		sequence.SelectedPattern = 0
	}
}

func chaserAutoGobo(sequence *common.Sequence) {

	if debug {
		fmt.Printf("chaserAutoGobo: \n")
	}

	if sequence.AutoColor {
		// Change all the fixtures to the next gobo.
		for fixtureNumber := range sequence.ScannersAvailable {
			sequence.ScannerGobo[fixtureNumber]++
			if sequence.ScannerGobo[fixtureNumber] > 8 {
				sequence.ScannerGobo[fixtureNumber] = 1
			}
		}
	}
}

func scannerAutoPattern(sequence *common.Sequence) {

	if debug {
		fmt.Printf("scannerAutoPattern: \n")
	}

	sequence.SelectedPattern++
	if sequence.SelectedPattern > 3 {
		sequence.SelectedPattern = 0
	}
}

func scannerAutoColor(sequence *common.Sequence) {

	if debug {
		fmt.Printf("scannerAutoColor: \n")
	}

	if sequence.AutoColor {
		// Change all the fixtures to the next gobo.
		for fixtureNumber := range sequence.ScannersAvailable {
			sequence.ScannerGobo[fixtureNumber]++
			if sequence.ScannerGobo[fixtureNumber] > 7 {
				sequence.ScannerGobo[fixtureNumber] = 0
			}
		}
		scannerLastColor := 0

		// AvailableFixtures gives the real number of configured scanners.
		for _, fixture := range sequence.ScannersAvailable {

			// First check that this fixture has some configured colors.
			colors, ok := sequence.ScannerAvailableColors[fixture.Number]
			if ok {
				// Found a scanner with some colors.
				totalColorForThisFixture := len(colors)

				// Now can mess with the scanner color map.
				sequence.ScannerColor[fixture.Number-1]++
				if sequence.ScannerColor[fixture.Number-1] > scannerLastColor {
					if sequence.ScannerColor[fixture.Number-1] >= totalColorForThisFixture {
						sequence.ScannerColor[fixture.Number-1] = 0
					}
					scannerLastColor++
					continue
				}
			}
		}
	}
}
