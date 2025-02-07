// Copyright (C) 2022, 2023 dhowlett99.
// This implements the load preset feature, used by the buttons package.
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

	"github.com/dhowlett99/dmxlights/pkg/colors"
	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/config"
	"github.com/dhowlett99/dmxlights/pkg/presets"
)

func loadPreset(sequences []*common.Sequence, this *CurrentState,
	X int, Y int,
	commandChannels []chan common.Command, eventsForLaunchpad chan common.ALight,
	guiButtons chan common.ALight, updateChannels []chan common.Sequence) {

	// Stop all sequences, so we start in sync.
	cmd := common.Command{
		Action: common.Stop,
	}
	common.SendCommandToAllSequence(cmd, commandChannels)

	// Load the config.
	// Which forces all sequences to load their config in the stopped position. Run=false.
	config.AskToLoadConfig(commandChannels, X, Y)

	// Turn the selected preset light flashing it's current color and yellow.
	if this.LastPreset != nil {
		last := this.PresetsStore[*this.LastPreset]
		this.PresetsStore[*this.LastPreset] = presets.Preset{State: last.State, Selected: false, Label: last.Label, ButtonColor: last.ButtonColor}
	}
	current := this.PresetsStore[fmt.Sprint(X)+","+fmt.Sprint(Y)]
	this.PresetsStore[fmt.Sprint(X)+","+fmt.Sprint(Y)] = presets.Preset{State: current.State, Selected: true, Label: current.Label, ButtonColor: current.ButtonColor}
	presets.RefreshPresets(eventsForLaunchpad, guiButtons, this.PresetsStore)

	// Preserve this.Blackout.
	if this.Blackout {
		cmd := common.Command{
			Action: common.Blackout,
		}
		common.SendCommandToAllSequence(cmd, commandChannels)
	}

	// Turn off the local copy of the this.Flood flag.
	this.Flood = false
	// And stop the flood button flashing.
	common.LightLamp(common.FLOOD_BUTTON, colors.White, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)

	// Remember we selected this preset
	last := fmt.Sprint(X) + "," + fmt.Sprint(Y)
	this.LastPreset = &last

	// Get an upto date copy of all of the sequences.
	for sequenceNumber := range sequences {

		sequences[sequenceNumber] = common.RefreshSequence(sequenceNumber, commandChannels, updateChannels)

		// If we're a scanner sequence and not in static mode clear the buttom
		if this.SequenceType[sequenceNumber] == "scanner" && !this.Static[sequenceNumber] {
			common.ClearSelectedRowOfButtons(sequenceNumber, eventsForLaunchpad, guiButtons)
		}

		// Clear any left over labels.
		common.ClearLabelsSelectedRowOfButtons(this.SelectedSequence, guiButtons)

		// Play out this sequence.
		displayMode(sequenceNumber, this.SelectedMode[this.SelectedSequence], this, sequences, eventsForLaunchpad, guiButtons, commandChannels)

		// Restore the speed, shift, size, fade, coordinates label data.
		this.Speed[sequenceNumber] = sequences[sequenceNumber].Speed
		this.RGBShift[sequenceNumber] = sequences[sequenceNumber].RGBShift
		this.ScannerShift[sequenceNumber] = sequences[sequenceNumber].ScannerShift
		this.RGBSize[sequenceNumber] = sequences[sequenceNumber].RGBSize
		this.ScannerSize[sequenceNumber] = sequences[sequenceNumber].ScannerSize
		this.RGBFade[sequenceNumber] = sequences[sequenceNumber].RGBFade
		this.ScannerCoordinates[sequenceNumber] = sequences[sequenceNumber].ScannerSelectedCoordinates
		this.Running[sequenceNumber] = sequences[sequenceNumber].SavedRun
		this.Strobe[sequenceNumber] = sequences[sequenceNumber].Strobe
		this.StrobeSpeed[sequenceNumber] = sequences[sequenceNumber].StrobeSpeed

		// Setup the correct mode for the displays.
		this.SequenceType[sequenceNumber] = sequences[sequenceNumber].Type

		// Assume we're starting in normal mode.
		this.SelectedMode[sequenceNumber] = NORMAL

		// Forget we've pressed twice.
		this.SelectButtonPressed[sequenceNumber] = false

		this.ScannerChaser[sequenceNumber] = sequences[sequenceNumber].ScannerChaser
		this.Static[sequenceNumber] = sequences[sequenceNumber].Static
		this.StaticFlashing[sequenceNumber] = false

		this.ShowStaticColorPicker = false
		this.ShowRGBColorPicker = false

		// If the scanner sequence isn't running but the shutter chaser is, then it makes sense to show the shutter chaser.
		if this.SequenceType[sequenceNumber] == "scanner" && !this.Running[this.ScannerSequenceNumber] && this.ScannerChaser[this.ScannerSequenceNumber] {
			// So adjust the mode to be CHASER_DISPLAY
			this.SelectedMode[sequenceNumber] = CHASER_DISPLAY
		}

		// Restore the functions states from the sequence.
		if sequences[sequenceNumber].Type == "rgb" {
			this.Functions[sequenceNumber][common.Function2_Auto_Color].State = sequences[sequenceNumber].AutoColor
			this.Functions[sequenceNumber][common.Function3_Auto_Pattern].State = sequences[sequenceNumber].AutoPattern
			this.Functions[sequenceNumber][common.Function4_Bounce].State = sequences[sequenceNumber].Bounce
			this.Functions[sequenceNumber][common.Function6_Static_Gobo].State = sequences[sequenceNumber].Static
			this.Functions[sequenceNumber][common.Function7_Invert_Chase].State = sequences[sequenceNumber].RGBInvert
			this.Functions[sequenceNumber][common.Function8_Music_Trigger].State = sequences[sequenceNumber].MusicTrigger
		}
		if sequences[sequenceNumber].Type == "scanner" {
			this.Functions[sequenceNumber][common.Function2_Auto_Color].State = sequences[sequenceNumber].AutoColor
			this.Functions[sequenceNumber][common.Function3_Auto_Pattern].State = sequences[sequenceNumber].AutoPattern
			this.Functions[sequenceNumber][common.Function4_Bounce].State = sequences[sequenceNumber].Bounce
			this.Functions[sequenceNumber][common.Function7_Invert_Chase].State = sequences[sequenceNumber].ScannerChaser
			this.Functions[sequenceNumber][common.Function8_Music_Trigger].State = sequences[sequenceNumber].MusicTrigger
		}

		// If we are loading a switch sequence, update our local copy of the switch settings.
		// and defocus each switch in turn.
		if sequences[sequenceNumber].Type == "switch" {

			// Get an upto date copy of the switch sequence.
			//sequences[sequenceNumber] = common.RefreshSequence(sequenceNumber, commandChannels, updateChannels)

			// Now set our local representation of switches
			for swiTchNumber, swiTch := range sequences[sequenceNumber].Switches {
				this.SwitchPosition[swiTchNumber] = swiTch.CurrentPosition

				overrides := *this.SwitchOverrides
				//  Restore any switch Overrides.
				if sequences[sequenceNumber].Switches[swiTchNumber].Override.Speed != 0 {
					overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Speed = sequences[sequenceNumber].Switches[swiTchNumber].Override.Speed
				}
				if sequences[sequenceNumber].Switches[swiTchNumber].Override.Shift != 0 {
					overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Shift = sequences[sequenceNumber].Switches[swiTchNumber].Override.Shift
				}
				if sequences[sequenceNumber].Switches[swiTchNumber].Override.Size != 0 {
					overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Size = sequences[sequenceNumber].Switches[swiTchNumber].Override.Size
				}
				if sequences[sequenceNumber].Switches[swiTchNumber].Override.Fade != 0 {
					overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Fade = sequences[sequenceNumber].Switches[swiTchNumber].Override.Fade
				}

				if sequences[sequenceNumber].Switches[swiTchNumber].Override.RotateSpeed != 0 {
					overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].RotateSpeed = sequences[sequenceNumber].Switches[swiTchNumber].Override.RotateSpeed
				}
				if sequences[sequenceNumber].Switches[swiTchNumber].Override.Colors != nil {
					overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Colors = sequences[sequenceNumber].Switches[swiTchNumber].Override.Colors
				}
				if sequences[sequenceNumber].Switches[swiTchNumber].Override.Gobo != 0 {
					overrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Gobo = sequences[sequenceNumber].Switches[swiTchNumber].Override.Gobo
				}
				this.SwitchOverrides = &overrides

				if debug {
					var stateNames []string
					for _, state := range swiTch.States {
						stateNames = append(stateNames, state.Name)
					}
					fmt.Printf("restoring switch number %d to postion %d states[%s]\n", swiTchNumber, this.SwitchPosition[swiTchNumber], stateNames)
				}
			}
			deFocusAllSwitches(this, sequences, commandChannels)
		}

		// Now start any thing that needs to run with load fixtures on.
		if sequences[sequenceNumber].SavedRun {
			cmd := common.Command{
				Action: common.Start,
			}
			common.SendCommandToSequence(sequenceNumber, cmd, commandChannels)
		}

		if debug {
			fmt.Printf("Loading Sequence %d Name %s Label %s Static %t\n", sequenceNumber, sequences[sequenceNumber].Name, sequences[sequenceNumber].Label, this.Static[sequenceNumber])
		}
	}

	// Restore the master brightness, remember that the master is for all sequences in this loaded config.
	// So the master we retrive from this selected sequence will be the same for all the others.
	this.MasterBrightness = sequences[this.SelectedSequence].Master

	// Show the correct running and strobe buttons.
	if this.Strobe[this.SelectedSequence] {
		this.StrobeSpeed[this.SelectedSequence] = sequences[this.SelectedSequence].StrobeSpeed
	}

	// Auto select the last running or static sequence which lights it's select lamp.
	this.SelectedSequence = autoSelect(this)
	// And set its type.
	this.SelectedType = this.SequenceType[this.SelectedSequence]

	// Show this sequence running status in the start/stop button.
	common.ShowRunningStatus(this.Running[this.SelectedSequence], eventsForLaunchpad, guiButtons)
	common.ShowStrobeButtonStatus(this.Strobe[this.SelectedSequence], eventsForLaunchpad, guiButtons)

	// Update the status bar.
	showStatusBars(this, sequences, eventsForLaunchpad, guiButtons)

	// Light the sequence selector button.
	lightSelectedButton(eventsForLaunchpad, guiButtons, this)

}

func autoSelect(this *CurrentState) (selectedSequence int) {

	// Check for running sequences.
	for sequenceNumber, sequenceRunning := range this.Running {
		if sequenceRunning {
			selectedSequence = sequenceNumber
			if selectedSequence == this.ChaserSequenceNumber {
				selectedSequence = this.ScannerSequenceNumber
			}
		}
	}

	// Check for static sequences.
	for sequenceNumber, sequenceInStatic := range this.Static {
		if sequenceInStatic {
			selectedSequence = sequenceNumber
			if selectedSequence == this.ChaserSequenceNumber {
				selectedSequence = this.ScannerSequenceNumber
			}
		}
	}

	// default to first sequnce if nothings running.
	return selectedSequence
}
