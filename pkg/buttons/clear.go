// Copyright (C) 2022, 2023 dhowlett99.
// This is clear function, used by the buttons package.
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

package buttons

import (
	"fmt"
	"image/color"

	"github.com/dhowlett99/dmxlights/pkg/colors"
	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
	"github.com/dhowlett99/dmxlights/pkg/presets"

	"github.com/oliread/usbdmx/ft232"
)

func Clear(X int, Y int, this *CurrentState, sequences []*common.Sequence, dmxController *ft232.DMXController, fixturesConfig *fixture.Fixtures,
	commandChannels []chan common.Command, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, updateChannels []chan common.Sequence) {

	debug := false

	if debug {
		fmt.Printf("CLEAR LAUNCHPAD\n")
	}

	if debug {
		fmt.Printf("Target Sequence %d\n", this.TargetSequence)
		fmt.Printf("Display Sequence %d\n", this.DisplaySequence)
		fmt.Printf("Which Static Sequnece are we Editing %d\n", this.EditWhichStaticSequence)
	}

	// Shortcut to clear rgb chase colors. We want to clear a color selection for a selected sequence.
	if this.ShowRGBColorPicker && !this.ClearPressed[this.TargetSequence] {

		if debug {
			fmt.Printf("Shortcut to clear rgb chase colors\n")
		}

		// Clear the local copy of colors.
		sequences[this.EditWhichStaticSequence].SequenceColors = []color.RGBA{}

		// Flash the correct color buttons
		ShowRGBColorPicker(*sequences[this.EditWhichStaticSequence], eventsForLaunchpad, guiButtons, commandChannels)

		// Clear has been pressed, next time we press clear we will get the full clear.
		this.ClearPressed[this.TargetSequence] = true

		return
	}

	// Shortcut to clear the selected set of static colors. We want to clear a static color selection for a target sequence.
	if this.Static[this.EditWhichStaticSequence] && !this.ClearPressed[this.TargetSequence] && !this.ShowStaticColorPicker {

		if debug {
			fmt.Printf("Shortcut to clear static colors\n")
		}

		if this.Static[this.EditWhichStaticSequence] && this.ShowRGBColorPicker {
			if debug {
				fmt.Printf("removeColorPicker\n")
			}
			removeColorPicker(this, sequences, eventsForLaunchpad, guiButtons, commandChannels)
		}

		// First press resets the colors to the default color bar.
		if !this.ClearPressed[this.TargetSequence] {

			if debug {
				fmt.Printf("Clear the sequence colors for this sequence %d\n", this.TargetSequence)
			}
			// Clear the sequence colors for this sequence.
			cmd := common.Command{
				Action: common.ClearStaticColor,
			}
			common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)
		}

		// Get an upto date copy of the sequence.
		sequences[this.TargetSequence] = common.RefreshSequence(this.TargetSequence, commandChannels, updateChannels)

		// Clear the pressed flag for all the fixtures.
		for x := 0; x < 8; x++ {
			sequences[this.TargetSequence].StaticColors[x].FirstPress = false
		}

		// Flash the correct color buttons
		common.ClearLabelsSelectedRowOfButtons(this.SelectedSequence, guiButtons)
		this.SelectedMode[this.DisplaySequence] = this.LastMode[this.DisplaySequence]

		// Clear the select all fixtures flag.
		this.SelectAllStaticFixtures = false

		// Clear has been pressed, next time we press clear we will get the full clear.
		this.ClearPressed[this.TargetSequence] = true

		// The sequence will automatically display the static colors now!
		return
	}

	if debug {
		fmt.Printf("Start full clear process\n")
	}

	// Start full clear process.
	if sequences[this.SelectedSequence].Type == "scanner" {
		buttonTouched(common.Button{X: X, Y: Y}, colors.Cyan, colors.White, eventsForLaunchpad, guiButtons)
	} else {
		buttonTouched(common.Button{X: X, Y: Y}, colors.White, colors.Magenta, eventsForLaunchpad, guiButtons)
	}

	// Reset the launchpad.
	if this.LaunchPadConnected {
		this.Pad.Program()
	}

	// Send reset to all sequences.
	// Look at the reset process in commands.go as a lot of stuff in the sequence is reset there.
	cmd := common.Command{
		Action: common.Reset,
	}
	common.SendCommandToAllSequence(cmd, commandChannels)

	// Turn off the flashing save button.
	this.SavePreset = false
	common.LightLamp(common.SAVE_BUTTON, colors.White, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)

	// Turn off the Running light.
	common.LightLamp(common.RUNNING_BUTTON, colors.White, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)

	// Turn off the this.Flood
	if this.Flood {
		this.Flood = false
		// Turn the flood button back to white.
		common.LightLamp(common.FLOOD_BUTTON, colors.White, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)
	}

	// Turn off the strobe light.
	common.LightLamp(common.STROBE_BUTTON, colors.White, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)

	// Clear out soundtriggers
	for _, trigger := range this.SoundTriggers {
		trigger.State = false
	}

	// Update status bar.
	common.UpdateStatusBar("Version"+" "+common.VERSION, "version", false, guiButtons)

	// Now go through all sequences and turn off stuff.
	for sequenceNumber, sequence := range sequences {
		this.SelectedSequence = 0                                                    // Update the status bar for the first sequnce. Because that will be the one selected after a clear.
		this.Strobe[sequenceNumber] = false                                          // Turn off the strobe.
		this.StrobeSpeed[sequenceNumber] = 255                                       // Reset to fastest strobe.
		this.Running[sequenceNumber] = false                                         // Stop the sequence.
		this.Speed[sequenceNumber] = common.DEFAULT_SPEED                            // Reset the speed back to the default.
		this.RGBShift[sequenceNumber] = common.DEFAULT_RGB_SHIFT                     // Reset the RGB shift back to the default.
		this.RGBSize[sequenceNumber] = common.DEFAULT_RGB_SIZE                       // Reset the RGB Size back to the default.
		this.RGBFade[sequenceNumber] = common.DEFAULT_RGB_FADE                       // Reset the RGB fade speed back to the default
		this.OffsetPan = common.SCANNER_MID_POINT                                    // Reset pan to the center
		this.OffsetTilt = common.SCANNER_MID_POINT                                   // Reset tilt to the center
		this.ScannerCoordinates[sequenceNumber] = common.DEFAULT_SCANNER_COORDNIATES // Reset the number of coordinates.
		this.ScannerSize[this.SelectedSequence] = common.DEFAULT_SCANNER_SIZE        // Reset the scanner size back to default.
		this.ScannerChaser[sequenceNumber] = false                                   // Clear the scanner chase mode.
		this.ScannerPattern = common.DEFAULT_PATTERN                                 // Reset the scanner pattern back to default.
		this.SwitchPosition = [NUMBER_SWITCHES]int{}                                 // Clear switch positions to their first positions.
		this.EditFixtureSelectionMode = false                                        // Clear fixture selecetd mode.
		this.SelectedMode[sequenceNumber] = NORMAL                                   // Clear function selecetd mode.
		this.SelectButtonPressed[sequenceNumber] = false                             // Clear buttoned selecetd mode.
		this.EditGoboSelectionMode = false                                           // Clear edit gobo mode.
		this.EditPatternMode = false                                                 // Clear edit pattern mode.
		this.EditScannerColorsMode = false                                           // Clear scanner color mode.
		this.ShowRGBColorPicker = false                                              // Clear rgb color mode.
		this.Static[sequenceNumber] = false                                          // Clear static color mode.
		this.ShowStaticColorPicker = false                                           // Clear the static color picker.
		this.MasterBrightness = common.MAX_DMX_BRIGHTNESS                            // Reset brightness to max.
		this.StaticFlashing[sequenceNumber] = false                                  // Turn off any flashing static buttons.
		this.ClearPressed[this.TargetSequence] = false                               // Reset the counter that counts how many times we've pressed the clear button.

		// Enable all fixtures.
		for fixtureNumber := 0; fixtureNumber < sequence.NumberFixtures; fixtureNumber++ {
			this.FixtureState[sequence.Number][fixtureNumber].Enabled = true
			this.FixtureState[sequence.Number][fixtureNumber].RGBInverted = false
			this.FixtureState[sequence.Number][fixtureNumber].ScannerPatternReversed = false
		}

		// Clear all the function buttons for this sequence.
		if sequence.Type != "switch" { // Switch sequences don't have funcion keys.
			this.Functions[sequenceNumber][common.Function1_Pattern].State = false
			this.Functions[sequenceNumber][common.Function2_Auto_Color].State = false
			this.Functions[sequenceNumber][common.Function3_Auto_Pattern].State = false
			this.Functions[sequenceNumber][common.Function4_Bounce].State = false
			this.Functions[sequenceNumber][common.Function5_Color].State = false
			this.Functions[sequenceNumber][common.Function6_Static_Gobo].State = false
			this.Functions[sequenceNumber][common.Function7_Invert_Chase].State = false
			this.Functions[sequenceNumber][common.Function8_Music_Trigger].State = false
		}

		// Reset the sequence switch states back to config from the fixture config in memory.
		// And ditch any out of date copy from a loaded preset.
		if sequence.Type == "switch" {
			// Get an upto date copy of the sequence.
			sequences[this.SelectedSequence] = common.RefreshSequence(this.SelectedSequence, commandChannels, updateChannels)

			// Now set our local representation of switches
			for swiTchNumber, swiTch := range sequence.Switches {
				this.SwitchPosition[swiTchNumber] = swiTch.CurrentPosition

				var stateNames []string

				// Find the details of the fixture for this switch.
				thisFixture, err := fixture.FindFixtureByLabel(swiTch.UseFixture, fixturesConfig)
				if err != nil {
					fmt.Printf("error %s\n", err.Error())
				}

				// Reset the speeds of switch sequences.
				for stateNumber, state := range swiTch.States {

					stateNames = append(stateNames, state.Name)

					action := fixture.GetSwitchConfig(swiTchNumber, int16(stateNumber), fixturesConfig)
					cfg := fixture.GetConfig(action, thisFixture, fixturesConfig)
					this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Speed = cfg.Speed
					this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Shift = cfg.Shift
					this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Size = cfg.Size
					this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Fade = cfg.Fade

					this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].RotateSpeed = cfg.RotateSpeed
					this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Colors = cfg.Colors
					this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Gobo = cfg.Gobo
				}
				if debug {
					fmt.Printf("restoring switch number %d to postion %d states[%s]\n", swiTchNumber, this.SwitchPosition[swiTchNumber], stateNames)
				}

			}
		}

		// Set the colors.
		//sequences[sequenceNumber].CurrentColors = sequences[sequenceNumber].SequenceColors

		// Handle the display for this sequence.
		this.SelectedSequence = sequenceNumber
		HandleSelect(sequences, this, eventsForLaunchpad, commandChannels, guiButtons)
	}

	deFocusSingleSwitch(this, sequences, commandChannels)

	// Clear the presets and display them.
	presets.ClearPresets(eventsForLaunchpad, guiButtons, this.PresetsStore)
	presets.RefreshPresets(eventsForLaunchpad, guiButtons, this.PresetsStore)

	// Set the first sequnence.
	this.SelectedSequence = 0
	this.LastSelectedSwitch = common.NOT_SELECTED
	HandleSelect(sequences, this, eventsForLaunchpad, commandChannels, guiButtons)

}
