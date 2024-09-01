// Copyright (C) 2022,2023 dhowlett99.
// This is the dmxlights main sequencers supporting functions.
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
	"bytes"
	"encoding/gob"
	"fmt"
	"image/color"

	"github.com/dhowlett99/dmxlights/pkg/colors"
	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/pattern"
	"github.com/dhowlett99/dmxlights/pkg/position"
)

func updatePatterns(sequence *common.Sequence, availablePatterns map[int]common.Pattern) []common.Step {

	if debug {
		fmt.Printf("updatePatterns: \n")
	}

	var steps []common.Step

	if sequence.Type == "rgb" {
		steps = updateRGBPatterns(sequence, availablePatterns)
	}
	if sequence.Type == "scanner" {
		steps = updateScannerPatterns(sequence)
	}

	return steps
}

func updateRGBPatterns(sequence *common.Sequence, availablePatterns map[int]common.Pattern) []common.Step {

	if debug {
		fmt.Printf("updateRGBPatterns: \n")
	}

	RGBPattern := position.ApplyFixtureState(availablePatterns[sequence.SelectedPattern], sequence.FixtureState)
	sequence.EnabledNumberFixtures = pattern.GetNumberEnabledScanners(sequence.FixtureState, sequence.NumberFixtures)
	steps := RGBPattern.Steps
	sequence.Pattern.Name = RGBPattern.Name
	sequence.Pattern.Label = RGBPattern.Label
	if sequence.NewPattern {
		sequence.SequenceColors = common.HowManyColorsInSteps(steps)
		sequence.NewPattern = false
	}

	// Initialise chaser.
	if sequence.Label == "chaser" {
		// Set the chase RGB steps used to chase the shutter.
		sequence.ScannerChaser = true
		// Chaser start with a standard chase pattern in white.
		steps = replaceRGBcolorsInSteps(steps, []color.RGBA{colors.White})
	}

	return steps
}

func updateScannerPatterns(sequence *common.Sequence) []common.Step {

	if debug {
		fmt.Printf("updateScannerPatterns: \n")
	}

	// Get available scanner patterns.
	sequence.ScannerAvailablePatterns = getAvailableScannerPattens(sequence)
	sequence.UpdatePattern = false
	sequence.EnabledNumberFixtures = pattern.GetNumberEnabledScanners(sequence.FixtureState, sequence.NumberFixtures)
	// Set the scanner steps used to send out pan and tilt values.
	sequence.Pattern = sequence.ScannerAvailablePatterns[sequence.SelectedPattern]
	steps := sequence.Pattern.Steps

	return steps
}

// Send a command to all the fixtures.
func sendToAllFixtures(fixtureChannels []chan common.FixtureCommand, command common.FixtureCommand) {
	for _, fixture := range fixtureChannels {
		fixture <- command
	}
}

func makeACopy(src, dist interface{}) (err error) {
	buf := bytes.Buffer{}
	if err = gob.NewEncoder(&buf).Encode(src); err != nil {
		return
	}
	return gob.NewDecoder(&buf).Decode(dist)
}

func replaceRGBcolorsInSteps(steps []common.Step, colors []color.RGBA) []common.Step {

	if debug {
		fmt.Printf("replaceRGBcolorsInSteps: with colors %+v\n", colors)
	}

	stepsOut := []common.Step{}
	err := makeACopy(steps, &stepsOut)
	if err != nil {
		fmt.Printf("replaceRGBcolorsInSteps: error failed to copy steps.\n")
	}

	var insertColor int
	numberColors := len(colors)

	for stepNumber, step := range steps {
		for fixtureNumber, fixture := range step.Fixtures {

			// found a color.
			if fixture.Color.R > 0 || fixture.Color.G > 0 || fixture.Color.B > 0 {
				if insertColor >= numberColors {
					insertColor = 0
				}
				newFixture := stepsOut[stepNumber].Fixtures[fixtureNumber]
				newFixture.Color = colors[insertColor]
				stepsOut[stepNumber].Fixtures[fixtureNumber] = newFixture
				insertColor++
			}

		}
	}

	if debug {
		for stepNumber, step := range stepsOut {
			fmt.Printf("Step %d\n", stepNumber)
			for fixtureNumber, fixture := range step.Fixtures {
				fmt.Printf("\tFixture %d\n", fixtureNumber)
				fmt.Printf("\t\tColor %+v\n", fixture.Color)
			}
		}
	}

	return stepsOut
}

// getAvailableScannerPattens generates scanner patterns and stores them in the sequence.
// Each scanner can then select which pattern to use.
// All scanner patterns have the same number of steps defined by NumberCoordinates.
func getAvailableScannerPattens(sequence *common.Sequence) map[int]common.Pattern {

	if debug {
		fmt.Printf("getAvailableScannerPattens\n")
	}

	scannerPattens := make(map[int]common.Pattern)

	// Scanner circle pattern 0
	coordinates := pattern.CircleGenerator(sequence.ScannerSize, sequence.ScannerCoordinates[sequence.ScannerSelectedCoordinates], float64(sequence.ScannerOffsetTilt), float64(sequence.ScannerOffsetPan))
	circlePatten := pattern.GeneratePattern(coordinates, sequence.NumberFixtures, sequence.ScannerShift, sequence.ScannerChaser, sequence.FixtureState)
	circlePatten.Name = "circle"
	circlePatten.Number = 0
	circlePatten.Label = "Circle"
	scannerPattens[0] = circlePatten

	// Scanner left right pattern 1
	coordinates = pattern.ScanGeneratorLeftRight(float64(sequence.ScannerSize), float64(sequence.ScannerCoordinates[sequence.ScannerSelectedCoordinates]), float64(sequence.ScannerOffsetTilt), float64(sequence.ScannerOffsetPan))
	leftRightPatten := pattern.GeneratePattern(coordinates, sequence.NumberFixtures, sequence.ScannerShift, sequence.ScannerChaser, sequence.FixtureState)
	leftRightPatten.Name = "leftright"
	leftRightPatten.Number = 1
	leftRightPatten.Label = "Left.Right"
	scannerPattens[1] = leftRightPatten

	// // Scanner up down pattern 2
	coordinates = pattern.ScanGeneratorUpDown(float64(sequence.ScannerSize), float64(sequence.ScannerCoordinates[sequence.ScannerSelectedCoordinates]), float64(sequence.ScannerOffsetTilt), float64(sequence.ScannerOffsetPan))
	upDownPatten := pattern.GeneratePattern(coordinates, sequence.NumberFixtures, sequence.ScannerShift, sequence.ScannerChaser, sequence.FixtureState)
	upDownPatten.Name = "updown"
	upDownPatten.Number = 2
	upDownPatten.Label = "Up.Down"
	scannerPattens[2] = upDownPatten

	// // Scanner zig zag pattern 3
	coordinates = pattern.ScanGenerateSawTooth(float64(sequence.ScannerSize), float64(sequence.ScannerCoordinates[sequence.ScannerSelectedCoordinates]), float64(sequence.ScannerCoordinates[sequence.ScannerSelectedCoordinates]), float64(sequence.ScannerOffsetTilt), float64(sequence.ScannerOffsetPan))
	zigZagPatten := pattern.GeneratePattern(coordinates, sequence.NumberFixtures, sequence.ScannerShift, sequence.ScannerChaser, sequence.FixtureState)
	zigZagPatten.Name = "zigzag"
	zigZagPatten.Number = 3
	zigZagPatten.Label = "Zig.Zag"
	scannerPattens[3] = zigZagPatten

	coordinates = []pattern.Coordinate{{Pan: 127, Tilt: 127}}
	stopPatten := pattern.GeneratePattern(coordinates, sequence.NumberFixtures, sequence.ScannerShift, sequence.ScannerChaser, sequence.FixtureState)
	stopPatten.Name = "stop"
	stopPatten.Number = 4
	stopPatten.Label = "Stop"
	scannerPattens[4] = stopPatten

	if debug {
		for _, pattern := range scannerPattens {
			fmt.Printf("Made a pattern called %s\n", pattern.Name)
		}
	}

	return scannerPattens

}
