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

func loadConfig(sequences []*common.Sequence, this *CurrentState,
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

		// Now start any thing that needs to run.
		if sequences[sequenceNumber].SavedRun {
			cmd := common.Command{
				Action: common.Start,
			}
			common.SendCommandToSequence(sequenceNumber, cmd, commandChannels)
		}

		restoreThis(this, sequenceNumber, *sequences[sequenceNumber])

		if debug {
			fmt.Printf("Loading Sequence %d Name %s Label %s Static %t\n", sequenceNumber, sequences[sequenceNumber].Name, sequences[sequenceNumber].Label, this.Static[sequenceNumber])
		}

		deFocusAllSwitches(this, sequences, commandChannels)

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

// restoreThis takes a sequence and copy's it contents into the current state represented by
// a pointer to a var called 'this' also passed in.
func restoreThis(this *CurrentState, sequenceNumber int, sequence common.Sequence) {

	// Restore the speed, shift, size, fade, coordinates label data.
	this.Speed[sequenceNumber] = sequence.Speed
	this.RGBShift[sequenceNumber] = sequence.RGBShift
	this.ScannerShift[sequenceNumber] = sequence.ScannerShift
	this.RGBSize[sequenceNumber] = sequence.RGBSize
	this.ScannerSize[sequenceNumber] = sequence.ScannerSize
	this.RGBFade[sequenceNumber] = sequence.RGBFade
	this.ScannerCoordinates[sequenceNumber] = sequence.ScannerSelectedCoordinates
	this.Running[sequenceNumber] = sequence.Run
	this.Strobe[sequenceNumber] = sequence.Strobe
	this.StrobeSpeed[sequenceNumber] = sequence.StrobeSpeed

	// Setup the correct mode for the displays.
	this.SequenceType[sequenceNumber] = sequence.Type

	// Assume we're starting in normal mode.
	this.SelectedMode[sequenceNumber] = NORMAL

	// Forget we've pressed twice.
	this.SelectButtonPressed[sequenceNumber] = false

	this.ScannerChaser[sequenceNumber] = sequence.ScannerChaser
	this.Static[sequenceNumber] = sequence.Static
	this.StaticFlashing[sequenceNumber] = false

	this.ShowStaticColorPicker = false
	this.ShowRGBColorPicker = false

	// If the scanner sequence isn't running but the shutter chaser is, then it makes sense to show the shutter chaser.
	if this.SequenceType[sequenceNumber] == "scanner" && !this.Running[this.ScannerSequenceNumber] && this.ScannerChaser[this.ScannerSequenceNumber] {
		// So adjust the mode to be CHASER_DISPLAY
		this.SelectedMode[sequenceNumber] = CHASER_DISPLAY
	}

	// Reload the fixture state.
	for fixtureNumber := 0; fixtureNumber < sequence.NumberFixtures; fixtureNumber++ {
		this.FixtureState[sequenceNumber][fixtureNumber] = sequence.FixtureState[fixtureNumber]
	}

	// Restore the functions states from the sequence.
	if sequence.Type == "rgb" {
		this.Functions[sequenceNumber][common.Function2_Auto_Color].State = sequence.AutoColor
		this.Functions[sequenceNumber][common.Function3_Auto_Pattern].State = sequence.AutoPattern
		this.Functions[sequenceNumber][common.Function4_Bounce].State = sequence.Bounce
		this.Functions[sequenceNumber][common.Function6_Static_Gobo].State = sequence.Static
		this.Functions[sequenceNumber][common.Function7_Invert_Chase].State = sequence.RGBInvert
		this.Functions[sequenceNumber][common.Function8_Music_Trigger].State = sequence.MusicTrigger
	}
	if sequence.Type == "scanner" {
		this.Functions[sequenceNumber][common.Function2_Auto_Color].State = sequence.AutoColor
		this.Functions[sequenceNumber][common.Function3_Auto_Pattern].State = sequence.AutoPattern
		this.Functions[sequenceNumber][common.Function4_Bounce].State = sequence.Bounce
		this.Functions[sequenceNumber][common.Function7_Invert_Chase].State = sequence.ScannerChaser
		this.Functions[sequenceNumber][common.Function8_Music_Trigger].State = sequence.MusicTrigger
	}

	// If we are loading a switch sequence, update our local copy of the switch settings.
	// and defocus each switch in turn.
	if sequence.Type == "switch" {

		// Get an upto date copy of the switch sequence.
		//sequences[sequenceNumber] = common.RefreshSequence(sequenceNumber, commandChannels, updateChannels)

		// Now set our local representation of switches
		for swiTchNumber, swiTch := range sequence.Switches {
			this.SwitchPosition[swiTchNumber] = swiTch.CurrentPosition

			//  Restore any switch Overrides.
			if sequence.Switches[swiTchNumber].Override.Speed != 0 {
				this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Speed = sequence.Switches[swiTchNumber].Override.Speed
			}
			if sequence.Switches[swiTchNumber].Override.Shift != 0 {
				this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Shift = sequence.Switches[swiTchNumber].Override.Shift
			}
			if sequence.Switches[swiTchNumber].Override.Size != 0 {
				this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Size = sequence.Switches[swiTchNumber].Override.Size
			}
			if sequence.Switches[swiTchNumber].Override.Fade != 0 {
				this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Fade = sequence.Switches[swiTchNumber].Override.Fade
			}

			if sequence.Switches[swiTchNumber].Override.RotateSpeed != 0 {
				this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].RotateSpeed = sequence.Switches[swiTchNumber].Override.RotateSpeed
			}
			if sequence.Switches[swiTchNumber].Override.Colors != nil {
				this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Colors = sequence.Switches[swiTchNumber].Override.Colors
			}
			if sequence.Switches[swiTchNumber].Override.Gobo != 0 {
				this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Gobo = sequence.Switches[swiTchNumber].Override.Gobo
			}

			// Defocus this switch.
			//this.LastSelectedSwitch = swiTchNumber
			//deFocusSingleSwitch(this, sequences, commandChannels)

			if debug {
				var stateNames []string
				for _, state := range swiTch.States {
					stateNames = append(stateNames, state.Name)
				}
				fmt.Printf("restoring switch number %d to postion %d states[%s]\n", swiTchNumber, this.SwitchPosition[swiTchNumber], stateNames)
			}
		}
	}
}
