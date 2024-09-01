// Copyright (C) 2022 dhowlett99.
// This is the dmxlights fixture editor it is attached to a fixture and
// describes the fixtures properties which is then saved in the fixtures.yaml
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

package commands

import (
	"fmt"
	"image/color"
	"time"

	"github.com/dhowlett99/dmxlights/pkg/colors"
	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/config"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
)

const debug = false
const beatDebug = false

// listenCommandChannelAndWait listens on channel for instructions or timeout and go to next step of sequence.
func ListenCommandChannelAndWait(mySequenceNumber int, currentSpeed time.Duration, sequence common.Sequence, channels common.Channels, fixturesConfig *fixture.Fixtures) common.Sequence {

	// Setup channels.
	commandChannel := channels.CommmandChannels[mySequenceNumber]
	replyChannel := channels.ReplyChannels[mySequenceNumber]
	updateChannel := channels.UpdateChannels[mySequenceNumber]

	// Create an empty command.
	command := common.Command{}

	// Wait for a trigger: sound, command or timeout.
	select {
	case command = <-channels.SoundTriggers[mySequenceNumber].Channel:
		if sequence.MusicTrigger {
			if beatDebug {
				fmt.Printf("%d: BEAT\n", mySequenceNumber)
			}
			break
		}
	case command = <-commandChannel:
		break

	case <-time.After(currentSpeed):
		break
	}

	// Now process any command.
	switch command.Action {

	case common.Reset:
		if debug {
			fmt.Printf("%d: Command Reset\n", mySequenceNumber)
		}
		// Turn off any static scenes.
		sequence.Static = false
		sequence.PlayStaticOnce = true
		// Stop the sequence.
		sequence.MusicTrigger = false
		sequence.Run = false
		sequence.Clear = true
		// Remove the hidden flag.
		sequence.Hidden = false
		// Clear the sequence colors.
		sequence.UpdateColors = false
		// Reset the speed back to the default.
		sequence.Speed = common.DEFAULT_SPEED
		sequence.CurrentSpeed = common.SetSpeed(common.DEFAULT_SPEED)
		// Stop the strobe mode.
		sequence.Strobe = false
		sequence.StrobeSpeed = common.DEFAULT_STROBE_SPEED
		sequence.StartFlood = false
		sequence.StopFlood = true
		sequence.FloodPlayOnce = true
		// Set Master brightness back to max.
		sequence.Master = common.MAX_DMX_BRIGHTNESS

		// Enable all fixtures
		for fixture := 0; fixture < sequence.NumberFixtures; fixture++ {
			newScannerState := common.FixtureState{}
			newScannerState.Enabled = true
			newScannerState.RGBInverted = false
			newScannerState.ScannerPatternReversed = false
			sequence.FixtureState[fixture] = newScannerState
		}

		if sequence.Type == "rgb" {
			// Reset the RGB shift back to the default.
			sequence.RGBShift = common.DEFAULT_RGB_SHIFT
			// Reset the RGB Size back to the default.
			sequence.RGBSize = common.DEFAULT_RGB_SIZE
			// Reset the RGB fade speed back to the default
			sequence.RGBFade = common.DEFAULT_RGB_FADE
			// Stop the flood mode.
			sequence.StartFlood = false
			sequence.StopFlood = true
			sequence.FloodPlayOnce = true
		}
		if sequence.Type == "scanner" {
			// Reset pan and tilt to the center
			sequence.ScannerOffsetPan = common.SCANNER_MID_POINT
			sequence.ScannerOffsetTilt = common.SCANNER_MID_POINT
			// Reset colors and gobo's.
			for scanner := 0; scanner < sequence.NumberFixtures; scanner++ {
				sequence.ScannerColor[scanner] = common.DEFAULT_SCANNER_COLOR // Reset Selected Color
				sequence.ScannerGobo[scanner] = common.DEFAULT_SCANNER_GOBO   // Reset Selected Gobo
			}
			// Reset the number of coordinates.
			sequence.ScannerSelectedCoordinates = common.DEFAULT_SCANNER_COORDNIATES
			// Reset the scanner size and shift back to defaults.
			sequence.ScannerSize = common.DEFAULT_SCANNER_SIZE
			sequence.ScannerShift = common.DEFAULT_SCANNER_SHIFT
			// Reset the scanner pattern back to default.
			sequence.UpdateColors = false
			sequence.RecoverSequenceColors = false
			sequence.UpdatePattern = true
			sequence.SelectedPattern = common.DEFAULT_PATTERN
		}
		// Clear all the function buttons for this sequence.
		sequence.SelectedPattern = common.DEFAULT_PATTERN
		sequence.AutoColor = false
		sequence.AutoPattern = false
		sequence.Bounce = false
		sequence.ScannerChaser = false
		sequence.RGBInvert = false
		sequence.MusicTrigger = false

		if sequence.Type == "switch" {
			if debug {
				fmt.Printf("Clear switch positions\n")
			}

			// read the fixtures config currently in memory.
			sequence.Switches = LoadSwitchConfiguration(mySequenceNumber, fixturesConfig)

			// Clear switch positions to their first positions.
			for switchNumber := 0; switchNumber < len(sequence.Switches); switchNumber++ {
				newSwitch := common.Switch{}
				newSwitch.CurrentPosition = 0
				newSwitch.Selected = false
				newOverride := common.Override{}
				// Default overrides set here.
				newOverride.Speed = 0
				newOverride.Shift = 0
				newOverride.Size = 0
				newOverride.Fade = 0
				newSwitch.Override = newOverride
				newSwitch.Description = sequence.Switches[switchNumber].Description
				newSwitch.Fixture = sequence.Switches[switchNumber].Fixture
				newSwitch.Label = sequence.Switches[switchNumber].Label
				newSwitch.Name = sequence.Switches[switchNumber].Name
				newSwitch.Number = sequence.Switches[switchNumber].Number
				newSwitch.States = sequence.Switches[switchNumber].States
				newSwitch.UseFixture = sequence.Switches[switchNumber].UseFixture
				sequence.Switches[switchNumber] = newSwitch
			}
			sequence.PlaySingleSwitch = false
			sequence.PlaySwitchOnce = true
			sequence.StepSwitch = true
			sequence.FocusSwitch = false
			sequence.OverrideSpeed = false
			sequence.OverrideShift = false
			sequence.OverrideSize = false
			sequence.OverrideFade = false
		}
		return sequence

	case common.Clear:
		if debug {
			fmt.Printf("%d: Command Clear\n", mySequenceNumber)
		}
		sequence.Clear = true
		return sequence

	case common.Hide:
		if debug {
			fmt.Printf("%d: Command Hide\n", mySequenceNumber)
		}
		sequence.Hidden = true
		return sequence

	case common.Reveal:
		if debug {
			fmt.Printf("%d: Command Reveal Static=%t\n", mySequenceNumber, sequence.Static)
		}
		sequence.Hidden = false
		sequence.PlayStaticOnce = false
		return sequence

	case common.UpdateSpeed:
		const SPEED = 0
		if debug {
			fmt.Printf("%d: Command Update %s to %d\n", mySequenceNumber, command.Args[SPEED].Name, command.Args[SPEED].Value)
		}
		sequence.Speed = command.Args[SPEED].Value.(int)
		sequence.CurrentSpeed = common.SetSpeed(command.Args[SPEED].Value.(int))
		return sequence

	case common.UpdatePattern:
		const PATTEN_NUMBER = 0
		if debug {
			fmt.Printf("%d: Command Update Pattern to number %d\n", mySequenceNumber, command.Args[PATTEN_NUMBER].Value)
		}
		sequence.UpdateColors = false
		sequence.RecoverSequenceColors = false
		sequence.UpdatePattern = true
		sequence.NewPattern = true
		sequence.SelectedPattern = command.Args[PATTEN_NUMBER].Value.(int)
		return sequence

	case common.UpdateRGBShift:
		const SHIFT = 0
		if debug {
			fmt.Printf("%d: Command Update RGB Shift to %d\n", mySequenceNumber, command.Args[SHIFT].Value)
		}
		sequence.RGBShift = command.Args[SHIFT].Value.(int)
		return sequence

	case common.UpdateRGBInvert:
		const INVERT = 0
		if debug {
			fmt.Printf("%d: Command Update RGB Invert to %t\n", mySequenceNumber, command.Args[INVERT].Value)
		}
		sequence.RGBInvert = command.Args[INVERT].Value.(bool)

		// Set all the fixtures to Invert
		if sequence.RGBInvert {
			for fixtureNumber := 0; fixtureNumber < sequence.NumberFixtures; fixtureNumber++ {
				state := sequence.FixtureState[fixtureNumber]
				state.RGBInverted = true
				state.ScannerPatternReversed = false
				sequence.FixtureState[fixtureNumber] = state
			}
		} else {
			for fixtureNumber := 0; fixtureNumber < sequence.NumberFixtures; fixtureNumber++ {
				state := sequence.FixtureState[fixtureNumber]
				state.RGBInverted = false
				state.ScannerPatternReversed = false
				sequence.FixtureState[fixtureNumber] = state
			}
		}

		return sequence

	case common.UpdateScannerShift:
		const SHIFT = 0
		if debug {
			fmt.Printf("%d: Command Update Scanner Shift to %d\n", mySequenceNumber, command.Args[SHIFT].Value)
		}
		sequence.UpdateShift = true
		sequence.ScannerShift = command.Args[SHIFT].Value.(int)
		return sequence

	case common.UpdateRGBSize:
		const SIZE = 0
		if debug {
			fmt.Printf("%d: Command Update Size to %d\n", mySequenceNumber, command.Args[SIZE].Value)
		}
		sequence.RGBSize = common.GetSize(command.Args[SIZE].Value.(int))
		return sequence

	case common.UpdateScannerSize:
		const SCANNER_SIZE = 0
		if debug {
			fmt.Printf("%d: Command Update Scanner Size to %d\n", mySequenceNumber, command.Args[SCANNER_SIZE].Value)
		}
		sequence.ScannerSize = command.Args[SCANNER_SIZE].Value.(int)
		return sequence

	case common.UpdateRGBFadeSpeed:
		const FADE_SPEED = 0
		if debug {
			fmt.Printf("%d: Command Set Fade to %d\n", mySequenceNumber, command.Args[FADE_SPEED].Value)
		}
		sequence.RGBFade = command.Args[FADE_SPEED].Value.(int)
		return sequence

	case common.Start:
		if debug {
			fmt.Printf("%d: Command Start\n", mySequenceNumber)
		}
		sequence.Chase = true
		sequence.Static = false
		sequence.UpdatePattern = true
		sequence.Run = true
		return sequence

	case common.StartChase:
		if debug {
			fmt.Printf("%d: Command StartChase\n", mySequenceNumber)
		}
		sequence.ScannerChaser = true
		sequence.Chase = true
		sequence.Static = false
		sequence.UpdatePattern = true
		sequence.Run = true
		return sequence

	case common.StopChase:
		if debug {
			fmt.Printf("%d: Command Stop Chase\n", mySequenceNumber)
		}
		sequence.ScannerChaser = false
		sequence.Chase = false
		sequence.Static = false
		sequence.Run = false
		return sequence

	case common.Stop:
		if debug {
			fmt.Printf("%d: Command Stop\n", mySequenceNumber)
		}
		sequence.MusicTrigger = false
		sequence.Run = false
		sequence.Static = false
		return sequence

	case common.Blackout:
		if debug {
			fmt.Printf("%d: Command Blackout\n", mySequenceNumber)
		}
		sequence.PlayStaticOnce = true
		sequence.PlaySwitchOnce = true
		sequence.Blackout = true
		return sequence

	case common.Flood:
		if debug {
			fmt.Printf("%d: Command to Start Flood\n", mySequenceNumber)
		}
		sequence.LastStatic = sequence.Static
		sequence.Static = false
		sequence.StartFlood = true
		sequence.StopFlood = false
		sequence.FloodPlayOnce = true
		return sequence

	case common.StopFlood:
		if debug {
			fmt.Printf("%d: Command to Stop Flood\n", mySequenceNumber)
		}
		sequence.Static = sequence.LastStatic
		sequence.StartFlood = false
		sequence.StopFlood = true
		sequence.FloodPlayOnce = true
		return sequence

	case common.Strobe:
		const STROBE_STATE = 0
		const STROBE_SPEED = 1
		if debug {
			fmt.Printf("%d: Command to Strobe set to %t at speed %d\n", mySequenceNumber, command.Args[STROBE_STATE].Value.(bool), command.Args[STROBE_SPEED].Value.(int))
		}
		// Remember the state of the Music trigger flag.
		sequence.LastMusicTrigger = sequence.MusicTrigger
		sequence.StrobeSpeed = command.Args[STROBE_SPEED].Value.(int)
		sequence.Strobe = command.Args[STROBE_STATE].Value.(bool)
		if sequence.StartFlood {
			sequence.FloodPlayOnce = true
		}
		if sequence.Static {
			sequence.PlayStaticOnce = true
		}
		return sequence

	case common.StopStrobe:
		if debug {
			fmt.Printf("%d: Command to Stop Strobe\n", mySequenceNumber)
		}
		sequence.Strobe = false
		sequence.StartFlood = false
		sequence.StopFlood = true
		sequence.FloodPlayOnce = true
		if sequence.Static {
			sequence.PlayStaticOnce = true
		}
		// Restore the state of the music trigger flag.
		sequence.MusicTrigger = sequence.LastMusicTrigger
		return sequence

	case common.UpdateStrobeSpeed:
		const STROBE_SPEED = 0
		if debug {
			fmt.Printf("%d: Command to Update Strobe Speed to %d\n", mySequenceNumber, command.Args[STROBE_SPEED].Value)
		}
		sequence.StrobeSpeed = command.Args[STROBE_SPEED].Value.(int)
		if sequence.StartFlood {
			sequence.FloodPlayOnce = true
		}
		if sequence.Static {
			sequence.PlayStaticOnce = true
		}

		return sequence

	case common.Normal:
		if debug {
			fmt.Printf("%d: Command Normal\n", mySequenceNumber)
		}
		// Normal is used to recover from blackout.
		sequence.Hidden = false
		sequence.StaticFadeUpOnce = true
		sequence.PlayStaticOnce = true
		sequence.PlaySwitchOnce = true
		sequence.PlaySingleSwitch = false
		sequence.Blackout = false
		return sequence

	case common.UpdateBounce:
		const STATE = 0
		if debug {
			fmt.Printf("%d: Command Update Bounce to %t\n", mySequenceNumber, command.Args[STATE].Value)
		}
		sequence.Bounce = command.Args[STATE].Value.(bool)
		return sequence

	case common.UpdateStatic: // Update Static will force the sequence to play the static scene.
		const STATIC = 0
		if debug {
			fmt.Printf("%d: Command Update Static to %t\n", mySequenceNumber, command.Args[STATIC].Value)
		}
		sequence.PlayStaticOnce = true
		sequence.PlaySwitchOnce = true
		sequence.StaticFadeUpOnce = false
		sequence.Static = command.Args[STATIC].Value.(bool)
		sequence.Run = false
		sequence.Hidden = false
		return sequence

	case common.UpdateFlashAllStaticColorButtons:
		const STATIC_FLASH = 0
		const STATIC_HIDDEN = 1
		if debug {
			fmt.Printf("%d: Command Flash All Static Colors to %t\n", mySequenceNumber, command.Args[STATIC_FLASH].Value)
		}
		for staticColor := range sequence.StaticColors {
			sequence.StaticColors[staticColor].Flash = command.Args[STATIC_FLASH].Value.(bool)
		}
		sequence.StaticFadeUpOnce = false // We don't want to fade as we set colors.
		sequence.PlayStaticLampsOnce = command.Args[STATIC_FLASH].Value.(bool)
		sequence.StaticLampsOn = true
		sequence.Hidden = command.Args[STATIC_HIDDEN].Value.(bool)
		return sequence

	case common.UpdateAllStaticColor:
		const STATIC = 0                // Boolean
		const STATIC_FIXTURE_FLASH = 1  // Boolean
		const STATIC_SELECTED_COLOR = 2 // Integer
		const STATIC_COLOR = 3          // Color
		if debug {
			fmt.Printf("%d: Command Update All Static Colors\n", mySequenceNumber)
			fmt.Printf("Selected Color:%d Flash:%t\n", command.Args[STATIC_SELECTED_COLOR].Value, command.Args[STATIC_FIXTURE_FLASH].Value)
			fmt.Printf("Lamp Color   %+v\n", command.Args[STATIC_COLOR].Value.(color.RGBA))
			fmt.Printf("Lamp Flash   %+v\n", command.Args[STATIC_FIXTURE_FLASH].Value.(bool))
		}
		// Set fixtures.
		for fixture := 0; fixture < sequence.NumberFixtures; fixture++ {
			sequence.StaticColors[fixture].SelectedColor = command.Args[STATIC_SELECTED_COLOR].Value.(int)
			sequence.StaticColors[fixture].Color = command.Args[STATIC_COLOR].Value.(color.RGBA)
			sequence.StaticColors[fixture].Flash = command.Args[STATIC_FIXTURE_FLASH].Value.(bool)
		}
		sequence.StaticFadeUpOnce = false // We don't want to fade as we set colors.
		sequence.PlayStaticOnce = true
		sequence.Static = command.Args[STATIC].Value.(bool)
		sequence.Hidden = true

		return sequence

	case common.UpdateStaticColor:
		const STATIC = 0                // Boolean
		const STATIC_FIXTURE_NUMBER = 1 // Integer
		const STATIC_FIXTURE_FLASH = 2  // Boolean
		const STATIC_SELECTED_COLOR = 3 // Integer
		const STATIC_COLOR = 4          // Color
		if debug {
			fmt.Printf("%d: Command Update Static Color\n", mySequenceNumber)
			fmt.Printf("Lamp Color   %+v\n", command.Args[STATIC_COLOR].Value.(color.RGBA))
			fmt.Printf("Selected Color:%d Flash:%t\n", command.Args[STATIC_SELECTED_COLOR].Value, command.Args[STATIC_FIXTURE_FLASH].Value)
		}
		sequence.StaticFadeUpOnce = false // We don't want to fade as we set colors.
		sequence.PlayStaticOnce = true
		sequence.Static = command.Args[STATIC].Value.(bool)
		sequence.Hidden = true
		// turn all flashing off first.
		for fixture := 0; fixture < sequence.NumberFixtures; fixture++ {
			sequence.StaticColors[fixture].Flash = false
		}
		sequence.StaticColors[command.Args[STATIC_FIXTURE_NUMBER].Value.(int)].SelectedColor = command.Args[STATIC_SELECTED_COLOR].Value.(int)
		sequence.StaticColors[command.Args[STATIC_FIXTURE_NUMBER].Value.(int)].Color = command.Args[STATIC_COLOR].Value.(color.RGBA)
		sequence.StaticColors[command.Args[STATIC_FIXTURE_NUMBER].Value.(int)].Flash = command.Args[STATIC_FIXTURE_FLASH].Value.(bool)
		return sequence

	case common.UpdateSequenceColors:
		const COLORS = 0
		if debug {
			fmt.Printf("%d: Command Update Sequence Color Type %s to %+v\n", mySequenceNumber, sequence.Type, command.Args[COLORS].Value)
		}

		sequence.SequenceColors = command.Args[COLORS].Value.([]color.RGBA)
		sequence.UpdateColors = true
		sequence.SaveColors = true

		return sequence

	case common.UpdateScannerColor:
		const SELECTED_COLOR = 0
		const FIXTURE_NUMBER = 1
		sequence.SaveColors = true
		selectedColor := command.Args[SELECTED_COLOR].Value.(int)
		selectedScanner := command.Args[FIXTURE_NUMBER].Value.(int)

		// Update Color for this scanner.
		sequence.ScannerColor[selectedScanner] = selectedColor

		// Clear out existing sequemce colors.
		sequence.SequenceColors = []color.RGBA{}

		// Look through the scanners and find their current color.
		// We set sequence colors so the color display shows the correct colors for the scanners.
		sequence.SequenceColors = fixture.HowManyScannerColors(&sequence, fixturesConfig)

		if debug {
			fmt.Printf("%d: Command Update Scanner Colors for fixture %d to %d final colors %+v\n", mySequenceNumber, selectedScanner, selectedColor, sequence.SequenceColors)
		}
		return sequence

	case common.ClearStaticColor:
		if debug {
			fmt.Printf("%d: Command Clear Static Color \n", mySequenceNumber)
		}
		// Populate the static colors for this sequence with the defaults.
		sequence.StaticColors = common.SetDefaultStaticColorButtons(mySequenceNumber)
		sequence.PlayStaticOnce = true
		sequence.StaticFadeUpOnce = false
		sequence.Static = true
		sequence.Hidden = false
		return sequence

	case common.SetStaticColorBar:
		const SELECTED_COLOR = 0
		if debug {
			fmt.Printf("%d: Command Set Static Color Bar to %d\n", mySequenceNumber, command.Args[SELECTED_COLOR].Value)
		}
		// Find the color bar for this selection.
		color := common.GetColorButtonsArray(command.Args[SELECTED_COLOR].Value.(int) - 1)
		newStaticColor := common.StaticColorButton{
			Color:         color,
			SelectedColor: command.Args[SELECTED_COLOR].Value.(int) - 1,
		}
		for fixture := 0; fixture < sequence.NumberFixtures; fixture++ {
			sequence.StaticColors[fixture] = newStaticColor
		}
		sequence.PlayStaticOnce = true
		sequence.Static = true
		sequence.Hidden = true
		return sequence

	// If we are being asked for our config we must reply with our current sequence.
	case common.ReadConfig:
		if debug {
			fmt.Printf("%d: Command Read Config\n", mySequenceNumber)
		}
		replyChannel <- sequence
		return sequence

	// We are setting the Master brightness in this sequence.
	case common.Master:
		const MASTER = 0
		if debug {
			fmt.Printf("%d: Command Master Brightness set to %d\n", mySequenceNumber, command.Args[MASTER].Value)
		}
		sequence.StaticFadeUpOnce = false // Don't soft fade as we change the brightness.
		sequence.PlayStaticOnce = true
		sequence.PlaySwitchOnce = true
		sequence.FloodPlayOnce = true
		sequence.Master = command.Args[MASTER].Value.(int)
		sequence.MasterChanging = true
		return sequence

	// If we are being asked for a updated config we must reply with our current sequence.
	case common.GetUpdatedSequence:
		if debug {
			fmt.Printf("%d: Command Get Updated Sequence\n", mySequenceNumber)
		}
		updateChannel <- sequence
		return sequence

	case common.ResetAllSwitchPositions:
		const FIXTURES_CONFIG = 0
		if debug {
			fmt.Printf("%d: Command ResetAllSwitchPositions n", mySequenceNumber)
		}
		sequence.Switches = LoadSwitchConfiguration(mySequenceNumber, command.Args[FIXTURES_CONFIG].Value.(*fixture.Fixtures))
		sequence.PlaySwitchOnce = true
		sequence.PlaySingleSwitch = false
		return sequence

	case common.OverrideSpeed:
		const SWITCH_NUMBER = 0 // Integer
		const SWITCH_POSITION = 1
		const SWITCH_SPEED = 2 // Integer
		switchNumber := command.Args[SWITCH_NUMBER].Value.(int)
		if debug {
			fmt.Printf("%d: Command Override Switch Number %d Position %d Speed %d\n", mySequenceNumber, switchNumber, command.Args[SWITCH_POSITION].Value, command.Args[SWITCH_SPEED].Value)
		}
		sequence.PlaySwitchOnce = true
		sequence.Override = true
		sequence.OverrideSpeed = true

		switchPosition := command.Args[SWITCH_POSITION].Value.(int)
		sequence.CurrentSwitch = switchNumber
		sequence.LastSwitchSelected = switchNumber

		newSwitch := common.Switch{}
		newSwitch.CurrentPosition = switchPosition
		newSwitch.Description = sequence.Switches[switchNumber].Description
		newSwitch.Fixture = sequence.Switches[switchNumber].Fixture
		newSwitch.Label = sequence.Switches[switchNumber].Label
		newSwitch.Name = sequence.Switches[switchNumber].Name
		newSwitch.Number = sequence.Switches[switchNumber].Number
		newSwitch.States = sequence.Switches[switchNumber].States
		newSwitch.UseFixture = sequence.Switches[switchNumber].UseFixture
		newSwitch.Selected = true

		newOverride := common.Override{}
		newOverride.Override = true
		newOverride.Speed = command.Args[SWITCH_SPEED].Value.(int)
		newOverride.Shift = sequence.Switches[switchNumber].Override.Shift
		newOverride.Size = sequence.Switches[switchNumber].Override.Size
		newOverride.Fade = sequence.Switches[switchNumber].Override.Fade
		newSwitch.Override = newOverride

		sequence.Switches[switchNumber] = newSwitch

		return sequence

	case common.OverrideShift:
		const SWITCH_NUMBER = 0 // Integer
		const SWITCH_POSITION = 1
		const SWITCH_SHIFT = 2 // Integer
		switchNumber := command.Args[SWITCH_NUMBER].Value.(int)
		if debug {
			fmt.Printf("%d: Command Override Switch Number %d Position %d Shift %d\n", mySequenceNumber, switchNumber, command.Args[SWITCH_POSITION].Value, command.Args[SWITCH_SHIFT].Value)
		}
		sequence.PlaySwitchOnce = true
		sequence.Override = true
		sequence.OverrideShift = true

		switchPosition := command.Args[SWITCH_POSITION].Value.(int)
		sequence.CurrentSwitch = switchNumber
		sequence.LastSwitchSelected = switchNumber

		newSwitch := common.Switch{}
		newSwitch.CurrentPosition = switchPosition
		newSwitch.Description = sequence.Switches[switchNumber].Description
		newSwitch.Fixture = sequence.Switches[switchNumber].Fixture
		newSwitch.Label = sequence.Switches[switchNumber].Label
		newSwitch.Name = sequence.Switches[switchNumber].Name
		newSwitch.Number = sequence.Switches[switchNumber].Number
		newSwitch.States = sequence.Switches[switchNumber].States
		newSwitch.UseFixture = sequence.Switches[switchNumber].UseFixture
		newSwitch.Selected = true

		newOverride := common.Override{}
		newOverride.Override = true
		newOverride.Speed = sequence.Switches[switchNumber].Override.Speed
		newOverride.Shift = command.Args[SWITCH_SHIFT].Value.(int)
		newOverride.Size = sequence.Switches[switchNumber].Override.Size
		newOverride.Fade = sequence.Switches[switchNumber].Override.Fade
		newSwitch.Override = newOverride

		sequence.Switches[switchNumber] = newSwitch

		return sequence

	case common.OverrideSize:
		const SWITCH_NUMBER = 0 // Integer
		const SWITCH_POSITION = 1
		const SWITCH_SIZE = 2 // Integer
		switchNumber := command.Args[SWITCH_NUMBER].Value.(int)
		if debug {
			fmt.Printf("%d: Command Override Switch Number %d Position %d Size %d\n", mySequenceNumber, switchNumber, command.Args[SWITCH_POSITION].Value, command.Args[SWITCH_SIZE].Value)
		}
		sequence.PlaySwitchOnce = true
		sequence.Override = true
		sequence.OverrideSize = true

		switchPosition := command.Args[SWITCH_POSITION].Value.(int)
		sequence.CurrentSwitch = switchNumber
		sequence.LastSwitchSelected = switchNumber

		newSwitch := common.Switch{}
		newSwitch.CurrentPosition = switchPosition
		newSwitch.Description = sequence.Switches[switchNumber].Description
		newSwitch.Fixture = sequence.Switches[switchNumber].Fixture
		newSwitch.Label = sequence.Switches[switchNumber].Label
		newSwitch.Name = sequence.Switches[switchNumber].Name
		newSwitch.Number = sequence.Switches[switchNumber].Number
		newSwitch.States = sequence.Switches[switchNumber].States
		newSwitch.UseFixture = sequence.Switches[switchNumber].UseFixture
		newSwitch.Selected = true

		newOverride := common.Override{}
		newOverride.Override = true
		newOverride.Speed = sequence.Switches[switchNumber].Override.Speed
		newOverride.Shift = sequence.Switches[switchNumber].Override.Shift
		newOverride.Size = command.Args[SWITCH_SIZE].Value.(int)
		newOverride.Fade = sequence.Switches[switchNumber].Override.Fade
		newSwitch.Override = newOverride

		sequence.Switches[switchNumber] = newSwitch

		return sequence

	case common.OverrideFade:
		const SWITCH_NUMBER = 0 // Integer
		const SWITCH_POSITION = 1
		const SWITCH_FADE = 2 // Integer
		switchNumber := command.Args[SWITCH_NUMBER].Value.(int)
		if debug {
			fmt.Printf("%d: Command Override Switch Number %d Position %d Fade %d\n", mySequenceNumber, switchNumber, command.Args[SWITCH_POSITION].Value, command.Args[SWITCH_FADE].Value)
		}
		sequence.PlaySwitchOnce = true
		sequence.Override = true
		sequence.OverrideFade = true

		switchPosition := command.Args[SWITCH_POSITION].Value.(int)
		sequence.CurrentSwitch = switchNumber
		sequence.LastSwitchSelected = switchNumber

		newSwitch := common.Switch{}
		newSwitch.CurrentPosition = switchPosition
		newSwitch.Description = sequence.Switches[switchNumber].Description
		newSwitch.Fixture = sequence.Switches[switchNumber].Fixture
		newSwitch.Label = sequence.Switches[switchNumber].Label
		newSwitch.Name = sequence.Switches[switchNumber].Name
		newSwitch.Number = sequence.Switches[switchNumber].Number
		newSwitch.States = sequence.Switches[switchNumber].States
		newSwitch.UseFixture = sequence.Switches[switchNumber].UseFixture
		newSwitch.Selected = true

		newOverride := common.Override{}
		newOverride.Override = true
		newOverride.Speed = sequence.Switches[switchNumber].Override.Speed
		newOverride.Shift = sequence.Switches[switchNumber].Override.Shift
		newOverride.Size = sequence.Switches[switchNumber].Override.Size
		newOverride.Fade = command.Args[SWITCH_FADE].Value.(int)
		newSwitch.Override = newOverride

		sequence.Switches[switchNumber] = newSwitch

		return sequence

	case common.OverrideRotateSpeed:
		const SWITCH_NUMBER = 0 // Integer
		const SWITCH_POSITION = 1
		const SWITCH_ROTATE_SPEED = 2 // Integer
		switchNumber := command.Args[SWITCH_NUMBER].Value.(int)
		if debug {
			fmt.Printf("%d: Command Override Switch Number %d Position %d Rotate Speed %d\n", mySequenceNumber, switchNumber, command.Args[SWITCH_POSITION].Value, command.Args[SWITCH_ROTATE_SPEED].Value)
		}
		sequence.PlaySwitchOnce = true
		sequence.Override = true
		sequence.OverrideRotateSpeed = true

		switchPosition := command.Args[SWITCH_POSITION].Value.(int)
		sequence.CurrentSwitch = switchNumber
		sequence.LastSwitchSelected = switchNumber

		newSwitch := common.Switch{}
		newSwitch.CurrentPosition = switchPosition
		newSwitch.Description = sequence.Switches[switchNumber].Description
		newSwitch.Fixture = sequence.Switches[switchNumber].Fixture
		newSwitch.Label = sequence.Switches[switchNumber].Label
		newSwitch.Name = sequence.Switches[switchNumber].Name
		newSwitch.Number = sequence.Switches[switchNumber].Number
		newSwitch.States = sequence.Switches[switchNumber].States
		newSwitch.UseFixture = sequence.Switches[switchNumber].UseFixture
		newSwitch.Selected = true

		newOverride := common.Override{}
		newOverride.Override = true
		newOverride.Speed = sequence.Switches[switchNumber].Override.Speed
		newOverride.Shift = sequence.Switches[switchNumber].Override.Shift
		newOverride.RotateSpeed = command.Args[SWITCH_ROTATE_SPEED].Value.(int)
		newOverride.Size = sequence.Switches[switchNumber].Override.Size
		newOverride.Fade = sequence.Switches[switchNumber].Override.Fade
		newSwitch.Override = newOverride

		sequence.Switches[switchNumber] = newSwitch

		return sequence

	case common.OverrideColor:
		const SWITCH_NUMBER = 0 // Integer
		const SWITCH_POSITION = 1
		const SWITCH_COLOR = 2 // Integer
		switchNumber := command.Args[SWITCH_NUMBER].Value.(int)
		if debug {
			fmt.Printf("%d: Command Override Switch Number %d Position %d Color %d\n", mySequenceNumber, switchNumber, command.Args[SWITCH_POSITION].Value, command.Args[SWITCH_COLOR].Value)
		}
		sequence.PlaySwitchOnce = true
		sequence.Override = true
		sequence.OverrideRotateSpeed = true

		switchPosition := command.Args[SWITCH_POSITION].Value.(int)
		sequence.CurrentSwitch = switchNumber
		sequence.LastSwitchSelected = switchNumber

		newSwitch := common.Switch{}
		newSwitch.CurrentPosition = switchPosition
		newSwitch.Description = sequence.Switches[switchNumber].Description
		newSwitch.Fixture = sequence.Switches[switchNumber].Fixture
		newSwitch.Label = sequence.Switches[switchNumber].Label
		newSwitch.Name = sequence.Switches[switchNumber].Name
		newSwitch.Number = sequence.Switches[switchNumber].Number
		newSwitch.States = sequence.Switches[switchNumber].States
		newSwitch.UseFixture = sequence.Switches[switchNumber].UseFixture
		newSwitch.Selected = true

		newOverride := common.Override{}
		newOverride.Override = true
		newOverride.Speed = sequence.Switches[switchNumber].Override.Speed
		newOverride.Shift = sequence.Switches[switchNumber].Override.Shift
		newOverride.RotateSpeed = sequence.Switches[switchNumber].Override.Shift
		newOverride.Color = command.Args[SWITCH_COLOR].Value.(int)
		newOverride.Size = sequence.Switches[switchNumber].Override.Size
		newOverride.Fade = sequence.Switches[switchNumber].Override.Fade
		newSwitch.Override = newOverride

		sequence.Switches[switchNumber] = newSwitch

		return sequence

	// Update the named switch position for the current sequence.
	case common.UpdateSwitch:
		const SWITCH_NUMBER = 0   // Integer
		const SWITCH_POSITION = 1 // Integer
		const SWITCH_STEP = 2     // Boolean,
		const SWITCH_FOCUS = 3    // Boolean, true to focus switch, full brighness. false to defocue dim button.
		if debug {
			fmt.Printf("%d: Command Update Switch %d to Position %d Step %t Focus %t\n",
				mySequenceNumber,
				command.Args[SWITCH_NUMBER].Value,
				command.Args[SWITCH_POSITION].Value,
				command.Args[SWITCH_STEP].Value, command.Args[SWITCH_FOCUS].Value)
		}

		// Loop through all the switchies. and reset their current state back to 0.
		switchNumber := command.Args[SWITCH_NUMBER].Value.(int)
		switchPosition := command.Args[SWITCH_POSITION].Value.(int)

		newSwitch := common.Switch{}
		newSwitch.CurrentPosition = switchPosition
		newSwitch.Description = sequence.Switches[switchNumber].Description
		newSwitch.Fixture = sequence.Switches[switchNumber].Fixture
		newSwitch.Label = sequence.Switches[switchNumber].Label
		newSwitch.Name = sequence.Switches[switchNumber].Name
		newSwitch.Number = sequence.Switches[switchNumber].Number
		newSwitch.States = sequence.Switches[switchNumber].States
		newSwitch.UseFixture = sequence.Switches[switchNumber].UseFixture
		newSwitch.Override = sequence.Switches[switchNumber].Override
		newSwitch.Selected = command.Args[SWITCH_FOCUS].Value.(bool)
		sequence.Switches[switchNumber] = newSwitch
		sequence.CurrentSwitch = command.Args[SWITCH_NUMBER].Value.(int)
		sequence.PlaySwitchOnce = true
		sequence.PlaySingleSwitch = true
		sequence.StepSwitch = command.Args[SWITCH_STEP].Value.(bool)
		sequence.FocusSwitch = command.Args[SWITCH_FOCUS].Value.(bool)
		sequence.Run = false
		sequence.Type = "switch"
		sequence.MasterChanging = false
		return sequence

	// Here we want to disable/enable the selected scanner.
	case common.ToggleFixtureState:
		const FIXTURE_NUMBER = 0       // Integer
		const FIXTURE_STATE = 1        // Boolean
		const FIXTURE_RGB_INVERTED = 2 // Boolean
		const FIXTURE_SCANNER_REVERSED = 3
		if debug {
			fmt.Printf("%d: Command ToggleFixtureState for fixture number %d, state %t, inverted %t reversed %t\n", mySequenceNumber, command.Args[FIXTURE_NUMBER].Value, command.Args[FIXTURE_STATE].Value, command.Args[FIXTURE_RGB_INVERTED].Value, command.Args[FIXTURE_SCANNER_REVERSED].Value)
		}

		// Set the state in the static color buttons.
		sequence.StaticColors[command.Args[FIXTURE_NUMBER].Value.(int)].Enabled = command.Args[FIXTURE_STATE].Value.(bool)

		if command.Args[FIXTURE_NUMBER].Value.(int) < sequence.NumberFixtures {
			newScannerState := common.FixtureState{}
			newScannerState.Enabled = command.Args[FIXTURE_STATE].Value.(bool)
			newScannerState.RGBInverted = command.Args[FIXTURE_RGB_INVERTED].Value.(bool)
			newScannerState.ScannerPatternReversed = command.Args[FIXTURE_SCANNER_REVERSED].Value.(bool)
			sequence.FixtureState[command.Args[FIXTURE_NUMBER].Value.(int)] = newScannerState
		}

		return sequence

	case common.UpdateGobo:
		const SELECTED_GOBO = 0
		const FIXTURE_NUMBER = 1
		if debug {
			fmt.Printf("%d: Command Update Gobo to Number %d\n", mySequenceNumber, command.Args[SELECTED_GOBO].Value)
		}
		sequence.ScannerGobo[command.Args[FIXTURE_NUMBER].Value.(int)] = command.Args[SELECTED_GOBO].Value.(int)
		sequence.Static = false
		return sequence

	case common.UpdateAutoColor:
		const AUTO_COLOR = 0
		const SELECTED_TYPE = 1
		if debug {
			fmt.Printf("Sequence %d: of Type %s : Command Update Auto Color to  %t\n", mySequenceNumber, command.Args[SELECTED_TYPE].Value, command.Args[AUTO_COLOR].Value)
		}
		sequence.AutoColor = command.Args[AUTO_COLOR].Value.(bool)
		selectedType := command.Args[SELECTED_TYPE].Value.(string)

		// If we switch auto color off and we are a rgb rembember what colors are in our sequence.
		if !sequence.AutoColor && selectedType == "rgb" {
			// If RecoverSequenceColors is true then we recover the colors from the SavedSequenceColors.
			sequence.RecoverSequenceColors = true
		}
		// If we switch auto color on and we are a rgb rembember what colors are in our sequence.
		if sequence.AutoColor && selectedType == "rgb" {
			// setting RecoverSequenceColors to false forces the sequence to save the currented
			// seleced colors to the SavedSequenceColors
			sequence.RecoverSequenceColors = false
		}

		// If switch auto color off and we are a scanner then reset the gobo and color back to the defaults.
		if !sequence.AutoColor && selectedType == "scanner" {
			for scanner := 0; scanner < sequence.NumberFixtures; scanner++ {
				sequence.ScannerColor[scanner] = common.DEFAULT_SCANNER_COLOR // Reset Selected Color
				sequence.ScannerGobo[scanner] = common.DEFAULT_SCANNER_GOBO   // Reset Selected Gobo
			}
		}

		return sequence

	case common.UpdateAutoPattern:
		const AUTO_PATTEN = 0
		if debug {
			fmt.Printf("%d: Command Update Auto Patten to  %t\n", mySequenceNumber, command.Args[AUTO_PATTEN].Value)
		}
		sequence.AutoPattern = command.Args[AUTO_PATTEN].Value.(bool)
		if !command.Args[AUTO_PATTEN].Value.(bool) {
			sequence.SelectedPattern = 0
		}
		return sequence

	case common.UpdateNumberCoordinates:
		const NUMBER_COORDINATES = 0
		if debug {
			fmt.Printf("%d: Command Update Number Coordinates to  %d\n", mySequenceNumber, command.Args[NUMBER_COORDINATES].Value)
		}
		sequence.ScannerSelectedCoordinates = command.Args[NUMBER_COORDINATES].Value.(int)
		return sequence

	case common.UpdateOffsetPan:
		const OFFSET_PAN = 0
		if debug {
			fmt.Printf("%d: Command Update Offset Pan to  %d\n", mySequenceNumber, command.Args[OFFSET_PAN].Value)
		}
		sequence.ScannerOffsetPan = command.Args[OFFSET_PAN].Value.(int)
		return sequence

	case common.UpdateOffsetTilt:
		const OFFSET_TILT = 0
		if debug {
			fmt.Printf("%d: Command Update Offset Tilt to  %d\n", mySequenceNumber, command.Args[OFFSET_TILT].Value)
		}
		sequence.ScannerOffsetTilt = command.Args[OFFSET_TILT].Value.(int)
		return sequence

	case common.UpdateScannerChase:
		const SCANNER_CHASE = 0
		if debug {
			fmt.Printf("%d: Command Update ScannerChase to %t \n", mySequenceNumber, command.Args[SCANNER_CHASE].Value)
		}
		sequence.ScannerChaser = command.Args[SCANNER_CHASE].Value.(bool)
		return sequence

	case common.UpdateMusicTrigger:
		const STATE = 0
		if debug {
			fmt.Printf("%d: Command Update Music Trigger to %t \n", mySequenceNumber, command.Args[STATE].Value)
		}
		sequence.MusicTrigger = command.Args[STATE].Value.(bool)
		sequence.Run = true
		sequence.Chase = true
		sequence.ScannerChaser = false
		if sequence.Label == "chaser" && sequence.Run {
			sequence.ScannerChaser = true
		}
		if sequence.Type == "scanner" && sequence.Label != "chaser" && sequence.Run {
			sequence.ScannerChaser = false
		}
		sequence.UpdatePattern = true
		sequence.Static = false
		sequence.PlayStaticOnce = false
		sequence.ChangeMusicTrigger = true
		return sequence

	case common.UpdateScannerHasShutterChase:
		const STATE = 0
		if debug {
			fmt.Printf("%d: Command Update ScannerHasShutterChase to %t \n", mySequenceNumber, command.Args[STATE].Value)
		}
		sequence.ScannerChaser = command.Args[STATE].Value.(bool)
		return sequence

	// If we are being asked to load a config, use the new sequence.
	case common.LoadConfig:
		const X = 0
		const Y = 1
		if debug {
			fmt.Printf("%d: Command Load Config\n", mySequenceNumber)
		}
		x := command.Args[X].Value.(int)
		y := command.Args[Y].Value.(int)
		config := config.LoadConfig(fmt.Sprintf("config%d.%d.json", x, y))
		for _, seq := range config {
			if seq.Number == sequence.Number {
				sequence = seq
				// Don't assume we're blacked out.
				sequence.Blackout = false
				sequence.StaticFadeUpOnce = true
				// Because steps are not stored in the config file. UpdatePattern generates the steps.
				sequence.UpdatePattern = true
				// Restore any sequenceColors.
				sequence.UpdateColors = true
				return sequence
			}
		}

	case common.UpdateFixturesConfig:
		if debug {
			fmt.Printf("%d: Command Update Fixure Config\n", mySequenceNumber)
		}
		const FIXTURES_CONFIG = 0
		fixturesConfig = command.Args[FIXTURES_CONFIG].Value.(*fixture.Fixtures)

		// Find the fixtures.
		sequence.ScannersAvailable = SetAvalableFixtures(mySequenceNumber, fixturesConfig)

		// Find the number of fixtures for this sequence.
		if sequence.Label == "chaser" {
			scannerSequenceNumber := common.GlobalScannerSequenceNumber // Scanner sequence number from config.
			sequence.NumberFixtures = GetNumberOfFixtures(scannerSequenceNumber, fixturesConfig)
		} else {
			sequence.NumberFixtures = GetNumberOfFixtures(mySequenceNumber, fixturesConfig)
		}

		// Setup fixtures labels.
		sequence.GuiFixtureLabels = []string{}
		for _, fixture := range fixturesConfig.Fixtures {
			if fixture.Type == "scanner" {
				sequence.GuiFixtureLabels = append(sequence.GuiFixtureLabels, fixture.Label)
			}
		}

		// Enable all the defined fixtures.
		for x := 0; x < sequence.NumberFixtures; x++ {
			newScanner := common.FixtureState{}
			newScanner.Enabled = true
			newScanner.RGBInverted = false
			newScanner.ScannerPatternReversed = false
			sequence.FixtureState[x] = newScanner
			// Set the first gobo for every fixture.
			sequence.ScannerGobo[x] = 1
		}

		return sequence
	}

	return sequence
}

func LoadSwitchConfiguration(mySequenceNumber int, fixturesConfig *fixture.Fixtures) map[int]common.Switch {

	if debug {
		fmt.Printf("Load switch data\n")
	}

	// A new group of switches.
	newSwitchList := make(map[int]common.Switch)

	switchNumber := 0

	for _, fixture := range fixturesConfig.Fixtures {
		if fixture.Group == mySequenceNumber+1 {
			// find switch data.
			newSwitch := common.Switch{}
			newSwitch.Name = fixture.Name
			newSwitch.ID = fixture.ID
			newSwitch.Address = fixture.Address
			newSwitch.Label = fixture.Label
			newSwitch.Number = fixture.Number
			newSwitch.Description = fixture.Description
			newSwitch.UseFixture = fixture.UseFixture

			// A switch has a number of states.
			newSwitch.States = make(map[int]common.State)
			for stateNumber, state := range fixture.States {
				newState := common.State{}
				newState.Name = state.Name
				newState.Number = state.Number
				newState.Label = state.Label

				// Copy settings.
				for _, setting := range state.Settings {
					newSetting := common.Setting{}
					newSetting.Name = setting.Name
					newSetting.Label = setting.Label
					newSetting.Number = setting.Number
					newSetting.Channel = setting.Channel
					newSetting.FixtureValue = setting.Value
					newState.Settings = append(newState.Settings, newSetting)
				}
				newState.ButtonColor = state.ButtonColor
				newState.Flash = state.Flash

				// Copy values.
				newState.Values = []common.Value{}
				for _, setting := range state.Settings {
					newValue := common.Value{}
					newValue.Channel = setting.Channel
					newValue.Setting = setting.Value
					newState.Values = append(newState.Values, newValue)
				}

				// Copy actions.
				newState.Actions = []common.Action{}
				for _, action := range state.Actions {
					newAction := common.Action{}
					newAction.Name = action.Name
					newAction.Colors = action.Colors
					newAction.Mode = action.Mode
					newAction.Fade = action.Fade
					newAction.Speed = action.Speed
					newAction.Rotate = action.Rotate
					newAction.RotateSpeed = action.RotateSpeed
					newAction.Program = action.Program
					newAction.ProgramSpeed = action.ProgramSpeed
					newAction.Gobo = action.Gobo
					newAction.GoboSpeed = action.GoboSpeed
					newAction.Map = action.Map
					newState.Actions = append(newState.Actions, newAction)
				}

				newSwitch.States[stateNumber] = newState
			}
			// Add new switch to the list.
			newSwitchList[switchNumber] = newSwitch
			switchNumber++
		}
	}

	return newSwitchList
}

func SetAvalableFixtures(sequenceNumber int, fixturesConfig *fixture.Fixtures) []common.StaticColorButton {

	// You need to select a fixture before you can choose a color or gobo.
	// availableFixtures holds a set of red buttons, one for every available fixture.
	availableFixtures := []common.StaticColorButton{}
	for _, f := range fixturesConfig.Fixtures {
		if f.Type == "scanner" {
			newFixture := common.StaticColorButton{}
			newFixture.Name = f.Name
			newFixture.Label = f.Label
			newFixture.Number = f.Number
			newFixture.SelectedColor = 1 // Red
			newFixture.Color = colors.Red
			newFixture.NumberOfGobos = fixture.HowManyGobosForThisFixture(f.Number, sequenceNumber, fixturesConfig)
			availableFixtures = append(availableFixtures, newFixture)
		}
	}

	return availableFixtures
}

func GetNumberOfFixtures(sequenceNumber int, fixtures *fixture.Fixtures) int {

	if debug {
		fmt.Printf("getNumberOfFixturesn for sequence %d\n", sequenceNumber)
	}

	var numberFixtures int

	for _, fixture := range fixtures.Fixtures {
		if fixture.Group-1 == sequenceNumber {
			// config has use_channels set.
			if fixture.MultiFixtureDevice {
				if debug {
					fmt.Printf("Sequence %d Found Number of Channels def. : %d\n", sequenceNumber, fixture.NumberSubFixtures)
				}
				// Since we don't yet have code that understands how to place a multi fixture device into a sequence
				// we always return the max channels in a sequence, currently 8
				return common.MAX_NUMBER_OF_CHANNELS
			} else {
				// Examine the channels and count number of color channels.
				// We use Red for the count.
				var subFixture int
				if subFixture > 1 {
					numberFixtures = numberFixtures + subFixture
				} else {
					if fixture.Number > numberFixtures {
						numberFixtures++
					}
				}
			}
		}
	}

	if debug {
		fmt.Printf("numberFixtures found %d\n", numberFixtures)
	}
	return numberFixtures
}
