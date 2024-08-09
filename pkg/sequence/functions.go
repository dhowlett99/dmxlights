package sequence

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"image/color"

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/pattern"
	"github.com/dhowlett99/dmxlights/pkg/position"
)

func setupRGBPatterns(sequence *common.Sequence, availablePatterns map[int]common.Pattern) []common.Step {

	RGBPattern := position.ApplyFixtureState(availablePatterns[sequence.SelectedPattern], sequence.FixtureState)

	sequence.EnabledNumberFixtures = pattern.GetNumberEnabledScanners(sequence.FixtureState, sequence.NumberFixtures)

	sequence.Pattern.Name = RGBPattern.Name
	sequence.Pattern.Label = RGBPattern.Label
	sequence.Pattern.Steps = RGBPattern.Steps

	// If we are updating the pattern, we also set the represention of the sequence colors.
	if sequence.UpdatePattern {
		sequence.SequenceColors = common.HowManyColorsInSteps(sequence.Pattern.Steps)
	}
	sequence.UpdatePattern = false

	// Initialise chaser.
	if sequence.Label == "chaser" {
		// Set the chase RGB steps used to chase the shutter.
		sequence.ScannerChaser = true
		// Chaser start with a standard chase pattern in white.
		sequence.Pattern.Steps = replaceRGBcolorsInSteps(sequence.Pattern.Steps, []color.RGBA{common.White})
	}

	return sequence.Pattern.Steps
}

func setupScannerPatterns(sequence *common.Sequence) []common.Step {

	// Get available scanner patterns.
	sequence.ScannerAvailablePatterns = getAvailableScannerPattens(sequence)
	sequence.UpdatePattern = false
	sequence.EnabledNumberFixtures = pattern.GetNumberEnabledScanners(sequence.FixtureState, sequence.NumberFixtures)

	// Set the scanner steps used to send out pan and tilt values.
	sequence.Pattern = sequence.ScannerAvailablePatterns[sequence.SelectedPattern]

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
	return sequence.Pattern.Steps
}

// Set the button color for the selected switch.
// Will also change the brightness to highlight the last selected switch.
func setSwitchLamp(sequence common.Sequence, switchNumber int, eventsForLauchpad chan common.ALight, guiButtons chan common.ALight) {

	swiTch := sequence.Switches[switchNumber]

	if debug {
		fmt.Printf("%d: switchNumber %d current %d selected %t\n", sequence.Number, swiTch.Number, swiTch.CurrentPosition, swiTch.Selected)
	}

	state := swiTch.States[swiTch.CurrentPosition]

	// Use the button color for this state to light the correct color on the launchpad.
	color, _ := common.GetRGBColorByName(state.ButtonColor)
	var brightness int
	if swiTch.Selected {
		brightness = common.MAX_DMX_BRIGHTNESS
	} else {
		brightness = common.MAX_DMX_BRIGHTNESS / 8
	}

	common.LightLamp(common.Button{X: switchNumber, Y: sequence.Number}, color, brightness, eventsForLauchpad, guiButtons)

	// Label the switch.
	common.LabelButton(switchNumber, sequence.Number, swiTch.Label+"\n"+state.Label, guiButtons)

}

// Set the DMX parameters for the selected switch.
func setSwitchDMX(sequence common.Sequence, switchNumber int, fixtureStepChannels []chan common.FixtureCommand) {

	swiTch := sequence.Switches[switchNumber]

	if debug {
		fmt.Printf("switchNumber %d current %d selected %t speed %d\n", swiTch.Number, swiTch.CurrentPosition, swiTch.Selected, sequence.Switches[swiTch.Number].Override.Speed)
	}

	state := swiTch.States[swiTch.CurrentPosition]

	// Now send a message to the fixture to play all the values for this state.
	command := common.FixtureCommand{
		Master:             sequence.Master,
		Blackout:           sequence.Blackout,
		Type:               sequence.Type,
		Label:              sequence.Label,
		SequenceNumber:     sequence.Number,
		SwiTch:             swiTch,
		State:              state,
		CurrentSwitchState: swiTch.CurrentPosition,
		MasterChanging:     sequence.MasterChanging,
		RGBFade:            sequence.RGBFade,
		Override:           sequence.Switches[swiTch.CurrentPosition].Override,
	}

	// Send a message to the fixture to operate the switch.
	fixtureStepChannels[switchNumber] <- command

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
