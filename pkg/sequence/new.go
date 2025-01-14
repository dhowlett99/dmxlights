// Copyright (C) 2022,2023 dhowlett99.
// This is the dmxlights main sequencer responsible for controlling all
// of the fixtures in a group. This file holds the code for creating
// a new sequence.
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
	"image/color"

	"github.com/dhowlett99/dmxlights/pkg/colors"
	"github.com/dhowlett99/dmxlights/pkg/commands"
	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
)

// Before a sequence can run it needs to be created.
// Assigns default values for all types of sequence.
// Returns a sequence of type common.Sequence.
func NewSequence(
	sequenceType string,
	sequenceLabel string,
	mySequenceNumber int,
	fixturesConfig *fixture.Fixtures,
	channels common.Channels) common.Sequence {

	// Populate the static colors for this sequence with the defaults.
	staticColorsButtons := common.SetDefaultStaticColorButtons(mySequenceNumber)

	// Populate the edit sequence colors for this sequence with the defaults.
	sequenceColorButtons := common.SetDefaultStaticColorButtons(mySequenceNumber)

	// Every scanner has a number of colors in its wheel.
	availableScannerColors := make(map[int][]common.StaticColorButton)

	// Find the fixtures.
	availableFixtures := commands.SetAvalableFixtures(mySequenceNumber, fixturesConfig)

	// Setup fixtures labels.
	fixtureLabels := []string{}
	for _, fixture := range fixturesConfig.Fixtures {
		if fixture.Type == "scanner" {
			fixtureLabels = append(fixtureLabels, fixture.Label)
		}
	}

	// Every scanner has a number of gobos in its wheel.
	availableScannerGobos := make(map[int][]common.StaticColorButton)

	// Create a map of the fixture colors.
	scannerColors := make(map[int]int)
	// Create a map of the fixture gobos.
	scannerGobos := make(map[int]int)

	if sequenceType == "scanner" {
		// Initilise Gobos
		availableScannerGobos = fixture.GetAvailableScannerGobos(mySequenceNumber, fixturesConfig)

		// Initialise Colors.
		availableScannerColors, scannerColors = fixture.GetAvailableScannerColors(fixturesConfig)
	}

	var numberFixtures int
	// Find the number of fixtures for this sequence.
	if sequenceLabel == "chaser" {
		scannerSequenceNumber := common.GlobalScannerSequenceNumber // Scanner sequence number from config.
		numberFixtures = commands.GetNumberOfFixtures(scannerSequenceNumber, fixturesConfig)
	} else {
		numberFixtures = commands.GetNumberOfFixtures(mySequenceNumber, fixturesConfig)
	}

	// A map of the state of fixtures in the sequence.
	// We can disable a fixture by setting fixture Enabled to false.
	// Set the sequence default fixture states.
	// Enable all the defined fixtures.
	// RGB Invert set to false
	// Scanner Pattern not reversed.
	fixtureState := make(map[int]common.FixtureState, numberFixtures) // Make space for the fixture states.
	for fixtureNumber := 0; fixtureNumber < numberFixtures; fixtureNumber++ {
		newFixture := common.FixtureState{}
		newFixture.Enabled = true
		newFixture.RGBInverted = false
		newFixture.ScannerPatternReversed = false
		fixtureState[fixtureNumber] = newFixture
		scannerGobos[fixtureNumber] = common.DEFAULT_SCANNER_GOBO
	}

	// Set default sequence colors.
	var defaultSequenceColors []color.RGBA
	if sequenceLabel == "chaser" {
		defaultSequenceColors = []color.RGBA{colors.White}
	}

	if sequenceType == "rgb" && sequenceLabel != "chaser" {
		defaultSequenceColors = []color.RGBA{colors.Green}
	}

	// The actual sequence definition.
	sequence := common.Sequence{
		Label:                  sequenceLabel,
		StartPattern:           true, // Start by setting up the pattern
		UpdateColors:           true, // And the colors
		SequenceColors:         defaultSequenceColors,
		ScannerAvailableColors: availableScannerColors,
		ScannersAvailable:      availableFixtures,
		NumberFixtures:         numberFixtures,
		Type:                   sequenceType,
		Hidden:                 false,
		Chase:                  true,
		StaticColors:           staticColorsButtons,
		RGBAvailableColors:     sequenceColorButtons,
		ScannerAvailableGobos:  availableScannerGobos,
		Name:                   sequenceType,
		Number:                 mySequenceNumber,
		RGBFade:                common.DEFAULT_RGB_FADE,
		MusicTrigger:           false,
		Run:                    false,
		Bounce:                 false,
		ScannerSize:            common.DEFAULT_SCANNER_SIZE,
		RGBSize:                common.DEFAULT_RGB_SIZE,
		Speed:                  common.DEFAULT_SPEED,
		ScannerShift:           common.DEFAULT_SCANNER_SHIFT,
		RGBShift:               common.DEFAULT_RGB_SHIFT,
		RGBNumberStepsInFade:   common.DEFAULT_RGB_FADE_STEPS,
		Blackout:               false,
		Master:                 common.MAX_DMX_BRIGHTNESS,
		ScannerGobo:            scannerGobos,
		StartFlood:             false,
		RGBColor:               1,
		AutoColor:              false,
		AutoPattern:            false,
		SelectedPattern:        common.DEFAULT_PATTERN,
		FixtureState:           fixtureState,
		ScannerCoordinates:     []int{12, 16, 24, 32, 64},
		ScannerColor:           scannerColors,
		ScannerOffsetPan:       common.SCANNER_MID_POINT,
		ScannerOffsetTilt:      common.SCANNER_MID_POINT,
		GuiFixtureLabels:       fixtureLabels,
		LoadNewFixtures:        true,
	}

	// Load the switch information in from the fixtures config.
	if sequenceType == "switch" {
		sequence.Switches = commands.LoadSwitchConfiguration(mySequenceNumber, fixturesConfig)
		sequence.PlaySwitchOnce = true
	}

	if sequenceType == "scanner" {
		// Get available scanner patterns.
		sequence.ScannerAvailablePatterns = getAvailableScannerPattens(&sequence)
		sequence.StartPattern = false
	}

	return sequence
}
