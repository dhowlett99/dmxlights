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

// rgbAutoColors - when called sets the sequence colors to the selected color indicated by the variable sequence.RGBColor
// and then increments the sequence.RGBColor so the next time around the sequence steps loop the color automatically changes.
// Currently supports only eight colors.
// Returns - A set of steps with the pattern color set to the selected color.
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

	if debug {
		fmt.Printf("sequence.RGBColor: %d Color %+v\n", sequence.RGBColor, sequence.SequenceColors)
	}

	// Now replace the color in the steps.
	steps = replaceRGBcolorsInSteps(steps, sequence.SequenceColors)

	return steps
}

// rgbAutoPattern - when called sets the sequence pattern to the selected color indicated by the variable sequence.SelectedPattern
// and then increments the sequence.SelectedPattern so the next time around the sequence steps loop the pattern automatically changes.
// Currently supports as many patterns as defined in availablePatterns as passed in.
// Returns - A set of steps with the pattern selected.
func rgbAutoPattern(sequence *common.Sequence, rgbAvailablePatterns map[int]common.Pattern) []common.Step {

	if debug {
		fmt.Printf("rgbAutoPattern: \n")
	}

	for patternNumber, pattern := range rgbAvailablePatterns {
		if pattern.Number == sequence.SelectedPattern {
			sequence.Pattern.Number = patternNumber
			if debug {
				fmt.Printf(">>>> I AM PATTEN %d\n", patternNumber)
			}
			break
		}
	}
	sequence.SelectedPattern++
	if sequence.SelectedPattern > len(rgbAvailablePatterns) {
		sequence.SelectedPattern = 0
	}

	return updateRGBPatterns(sequence, rgbAvailablePatterns)
}

// chaserAutoGobo - when called sets increments the sequences scanner gobo indicated by the variable sequence.ScannerGobo
// Currently supports as many Gobos only 8 gobos.
// Returns - A set of steps with the selected gobo in the pattern.
func chaserAutoGobo(sequence *common.Sequence) []common.Step {

	if debug {
		fmt.Printf("chaserAutoGobo: \n")
	}

	if sequence.AutoColor {
		// Change all the fixtures to the next gobo.
		for fixtureNumber, scanner := range sequence.ScannersAvailable {
			sequence.ScannerGobo[fixtureNumber]++
			if sequence.ScannerGobo[fixtureNumber] > scanner.NumberOfGobos {
				sequence.ScannerGobo[fixtureNumber] = 1
			}
		}
	}

	return updateScannerPatterns(sequence)
}

// scannerAutoPattern - when called sets increments the sequences scanner pattern indicated by the variable sequence.SelectedPattern
// Currently supports 4 scanner patterns.
// Returns - Nothing, pattern is determined by sequence.SelectedPattern.
func scannerAutoPattern(sequence *common.Sequence) []common.Step {

	if debug {
		fmt.Printf("scannerAutoPattern: \n")
	}

	sequence.SelectedPattern++
	if sequence.SelectedPattern > len(sequence.ScannerAvailablePatterns) {
		sequence.SelectedPattern = 0
	}

	if debug {
		fmt.Printf("SelectedPattern: %d\n", sequence.SelectedPattern)
	}

	return updateScannerPatterns(sequence)
}

// scannerAutoColor - when called changes all the fixtures to the next gobo and changes all the fixtures to the next color.
// Fixtures / scanners gobos and colors are indicated by the variables sequence.ScannerGobo and sequence.ScannerColor.
// Currently supports up to 8 scanner gobos and as many colors defined by sequence.ScannerAvailableColors
// Returns - Nothing, gobo is determined by sequence.ScannerGobo. Color is determined by sequence.ScannerColor
func scannerAutoColor(sequence *common.Sequence) []common.Step {
	if debug {
		fmt.Printf("scannerAutoColor: \n")
	}

	if sequence.AutoColor {
		// Change all the fixtures to the next gobo.
		for fixtureNumber, scanner := range sequence.ScannersAvailable {
			sequence.ScannerGobo[fixtureNumber]++
			if sequence.ScannerGobo[fixtureNumber] > scanner.NumberOfGobos {
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

	return updateScannerPatterns(sequence)
}
