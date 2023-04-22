package buttons

import (
	"fmt"

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
	"github.com/dhowlett99/dmxlights/pkg/presets"
	"github.com/dhowlett99/dmxlights/pkg/sequence"
	"github.com/oliread/usbdmx/ft232"
)

func clear(X int, Y int, this *CurrentState, sequences []*common.Sequence, dmxController *ft232.DMXController, fixturesConfig *fixture.Fixtures,
	commandChannels []chan common.Command, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, updateChannels []chan common.Sequence) {

	if debug {
		fmt.Printf("CLEAR LAUNCHPAD\n")
	}

	// Shortcut to clear rgb chase colors. We want to clear a color selection for a selected sequence.
	if this.Functions[this.SelectedSequence][common.Function5_Color].State &&
		sequences[this.SelectedSequence].Type != "scanner" {

		// Clear the sequence colors for this sequence.
		cmd := common.Command{
			Action: common.ClearSequenceColor,
		}
		common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

		// Get an upto date copy of the sequence.
		sequences[this.SelectedSequence] = common.RefreshSequence(this.SelectedSequence, commandChannels, updateChannels)

		// Flash the correct color buttons
		ShowRGBColorSelectionButtons(this.MasterBrightness, this.SelectedSequence, *sequences[this.SelectedSequence], this.SelectedSequence, eventsForLaunchpad, guiButtons)

		return
	}

	// Shortcut to clear static colors. We want to clear a static color selection for a selected sequence.
	if this.Functions[this.SelectedSequence][common.Function6_Static_Gobo].State &&
		sequences[this.SelectedSequence].Type != "scanner" {

		// Back to the begining of the rotation.
		if this.SelectColorBar[this.SelectedSequence] > common.MaxColorBar {
			this.SelectColorBar[this.SelectedSequence] = 0
		}

		// First press resets the colors to the default color bar.
		if this.SelectColorBar[this.SelectedSequence] == 0 {
			// Clear the sequence colors for this sequence.
			cmd := common.Command{
				Action: common.ClearStaticColor,
			}
			common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)
		}

		// Rotate around solid colors.
		if this.SelectColorBar[this.SelectedSequence] > 0 {

			// Clear the sequence colors for this sequence.
			cmd := common.Command{
				Action: common.SetStaticColorBar,
				Args: []common.Arg{
					{Name: "Selection", Value: this.SelectColorBar[this.SelectedSequence]},
				},
			}
			common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)
		}

		// Now increment the color bar.
		this.SelectColorBar[this.SelectedSequence]++

		// Get an upto date copy of the sequence.
		sequences[this.SelectedSequence] = common.RefreshSequence(this.SelectedSequence, commandChannels, updateChannels)

		// Clear the pressed flag for all the fixtures.
		for x := 0; x < 8; x++ {
			sequences[this.SelectedSequence].StaticColors[x].FirstPress = false
		}

		// Flash the correct color buttons
		common.ClearLabelsSelectedRowOfButtons(this.SelectedSequence, guiButtons)
		this.SelectMode[this.SelectedSequence] = NORMAL
		// The sequence will automatically display the static colors now!

		return
	}

	// Start clear process.
	if sequences[this.SelectedSequence].Type == "scanner" {
		buttonTouched(common.ALight{X: X, Y: Y, OnColor: common.Cyan, OffColor: common.White}, eventsForLaunchpad, guiButtons)
	} else {
		buttonTouched(common.ALight{X: X, Y: Y, OnColor: common.White, OffColor: common.Pink}, eventsForLaunchpad, guiButtons)
	}

	// Turn off the flashing save button.
	this.SavePreset = false
	common.LightLamp(common.ALight{X: 8, Y: 4, Brightness: 255, Red: 255, Green: 255, Blue: 255}, eventsForLaunchpad, guiButtons)

	// Turn off the Running light.
	common.LightLamp(common.ALight{X: 8, Y: 5, Brightness: this.MasterBrightness, Red: 255, Green: 255, Blue: 255}, eventsForLaunchpad, guiButtons)

	// Turn off the this.Flood
	if this.Flood {
		this.Flood = false
		// Turn the flood button back to white.
		common.LightLamp(common.ALight{X: 8, Y: 3, Brightness: this.MasterBrightness, Red: 255, Green: 255, Blue: 255}, eventsForLaunchpad, guiButtons)
	}

	// Turn off the strobe light.
	common.LightLamp(common.ALight{X: 8, Y: 6, Brightness: 255, Red: 255, Green: 255, Blue: 255, Flash: false}, eventsForLaunchpad, guiButtons)

	// Clear out soundtriggers
	for _, trigger := range this.SoundTriggers {
		trigger.State = false
	}
	// Update status bar.
	common.UpdateStatusBar("Version 2.0", "version", false, guiButtons)

	// Now go through all sequences and turn off stuff.
	for sequenceNumber, sequence := range sequences {
		this.SelectedSequence = 0                                                  // Update the status bar for the first sequnce. Because that will be the one selected after a clear.
		this.Strobe[sequenceNumber] = false                                        // Turn off the strobe.
		this.StrobeSpeed[sequenceNumber] = 255                                     // Reset to fastest strobe.
		this.Running[sequenceNumber] = false                                       // Stop the sequence.
		this.Speed[sequenceNumber] = common.DefaultSpeed                           // Reset the speed back to the default.
		this.RGBShift[sequenceNumber] = common.DefaultRGBShift                     // Reset the RGB shift back to the default.
		this.RGBSize[sequenceNumber] = common.DefaultRGBSize                       // Reset the RGB Size back to the default.
		this.RGBFade[sequenceNumber] = common.DefaultRGBFade                       // Reset the RGB fade speed back to the default
		this.OffsetPan = common.ScannerMidPoint                                    // Reset pan to the center
		this.OffsetTilt = common.ScannerMidPoint                                   // Reset tilt to the center
		this.ScannerCoordinates[sequenceNumber] = common.DefaultScannerCoordinates // Reset the number of coordinates.
		this.ScannerSize[this.SelectedSequence] = common.DefaultScannerSize        // Reset the scanner size back to default.
		this.ScannerPattern = common.DefaultPattern                                // Reset the scanner pattern back to default.
		this.SwitchPositions = [9][9]int{}                                         // Clear switch positions to their first positions.
		this.EditFixtureSelectionMode = false                                      // Clear fixture selecetd mode.
		this.SelectMode[sequenceNumber] = NORMAL                                   // Clear function selecetd mode.
		this.SelectButtonPressed[sequenceNumber] = false                           // Clear buttoned selecetd mode.
		this.EditGoboSelectionMode[sequenceNumber] = false                         // Clear edit gobo mode.
		this.EditPatternMode[sequenceNumber] = false                               // Clear edit pattern mode.
		this.EditScannerColorsMode[sequenceNumber] = false                         // Clear scanner color mode.
		this.EditSequenceColorsMode[sequenceNumber] = false                        // Clear rgb color mode.
		this.EditStaticColorsMode[sequenceNumber] = false                          // Clear static color mode.
		this.MasterBrightness = common.MaxDMXBrightness                            // Reset brightness to max.
		this.ScannerChaser = false                                                 // Clear the scanner chase mode.

		// Enable all fixtures.
		for fixtureNumber := 0; fixtureNumber < sequence.NumberFixtures; fixtureNumber++ {
			this.FixtureState[fixtureNumber][sequence.Number].Enabled = true
			this.FixtureState[fixtureNumber][sequence.Number].Inverted = false
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
				this.SwitchPositions[sequenceNumber][swiTchNumber] = swiTch.CurrentPosition
				if debug {
					var stateNames []string
					for _, state := range swiTch.States {
						stateNames = append(stateNames, state.Name)
					}
					fmt.Printf("restoring switch number %d to postion %d states[%s]\n", swiTchNumber, this.SwitchPositions[sequenceNumber][swiTchNumber], stateNames)
				}
			}
		}
	}

	// Send reset to all sequences.
	cmd := common.Command{
		Action: common.Reset,
	}
	common.SendCommandToAllSequence(cmd, commandChannels)

	// Clear the presets and display them.
	presets.ClearPresets(eventsForLaunchpad, guiButtons, this.PresetsStore)
	presets.RefreshPresets(eventsForLaunchpad, guiButtons, this.PresetsStore)

	// Turn off all fixtures.
	cmd = common.Command{
		Action: common.Clear,
	}
	common.SendCommandToAllSequence(cmd, commandChannels)

	// Light the correct sequence selector button.
	sequence.SequenceSelect(eventsForLaunchpad, guiButtons, this.SelectedSequence)

	// Clear the graphics labels.
	HandleSelect(sequences, this, eventsForLaunchpad, commandChannels, guiButtons)

	// Reset the launchpad.
	if this.LaunchPadConnected {
		this.Pad.Program()
	}
}
