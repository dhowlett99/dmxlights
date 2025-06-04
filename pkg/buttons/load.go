// Copyright (C) 2022, 2023, 2024, 2025 dhowlett99.
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
	"github.com/dhowlett99/dmxlights/pkg/labels"
	"github.com/dhowlett99/dmxlights/pkg/presets"
)

func loadPreset(sequences []*common.Sequence, this *CurrentState,
	X int, Y int,
	commandChannels []chan common.Command, eventsForLaunchpad chan common.ALight,
	guiButtons chan common.ALight, updateChannels []chan common.Sequence) {

	var lastSelectedSequence int
	var lastSelectedSwitch int

	// Stop all sequences, so we start in sync.
	cmd := common.Command{
		Action: common.Stop,
	}
	common.SendCommandToAllSequence(cmd, commandChannels)

	// Load the config.
	// Which forces all sequences to load their config in the stopped position. Run=false.
	config.AskToLoadConfig(commandChannels, X, Y, this.ProjectName)

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
		this.Paused[sequenceNumber] = false

		// Setup the correct mode for the displays.
		this.SequenceType[sequenceNumber] = sequences[sequenceNumber].Type

		// Assume we're starting in normal mode.
		this.SelectedMode[sequenceNumber] = NORMAL

		// Forget we've pressed twice.
		this.SelectButtonPressed[sequenceNumber] = false

		// Load the last selected sequence from the preset.
		lastSelectedSequence = sequences[sequenceNumber].LastSelectedSequence

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
			sequences[sequenceNumber] = common.RefreshSequence(sequenceNumber, commandChannels, updateChannels)

			// Load the last selected switch from the sequence.
			lastSelectedSwitch = sequences[sequenceNumber].LastSelectedSwitch
			// Get the overrides.
			RefreshLocalOverrides(this, sequences[sequenceNumber])

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

	// Auto select the last running or static sequence or switch which lights it's selected lamp.
	this.SelectedSequence, this.LastSelectedSwitch = autoSelect(this, commandChannels, lastSelectedSequence, lastSelectedSwitch)

	// And set its type.
	this.SelectedType = this.SequenceType[this.SelectedSequence]

	// Show this sequence running status in the start/stop button.
	common.ShowRunningStatus(this.Running[this.SelectedSequence], eventsForLaunchpad, guiButtons)
	common.ShowStrobeButtonStatus(this.Strobe[this.SelectedSequence], eventsForLaunchpad, guiButtons)

	// Update the status bar.
	showStatusBars(this, sequences, eventsForLaunchpad, guiButtons)

	// Light the sequence selector button.
	lightSelectedButton(eventsForLaunchpad, guiButtons, this)

	// Clear the pause button.
	this.AllPaused = false
	common.LightLamp(common.FREEZE_BUTTON, colors.White, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)
	common.LabelButton(common.FREEZE_BUTTON.X, common.FREEZE_BUTTON.Y, labels.GetLabel(this.Labels, "Freeze", "Off"), guiButtons)

}

func autoSelect(this *CurrentState, commandChannels []chan common.Command, lastSelectedSequence int, lastSelectedSwitch int) (int, int) {

	// Check for running sequences.
	for sequenceNumber, sequenceRunning := range this.Running {
		if sequenceRunning {
			lastSelectedSequence = sequenceNumber
			if lastSelectedSequence == this.ChaserSequenceNumber {
				lastSelectedSequence = this.ScannerSequenceNumber
				if debug {
					fmt.Printf("Found a running sequence %d\n", sequenceNumber)
				}
				return lastSelectedSequence, common.NOT_SELECTED
			}
		}
	}

	// Check for static sequences.
	for sequenceNumber, sequenceInStatic := range this.Static {
		if sequenceInStatic {
			lastSelectedSequence = sequenceNumber
			if lastSelectedSequence == this.ChaserSequenceNumber {
				lastSelectedSequence = this.ScannerSequenceNumber
				fmt.Printf("Found a static sequence %d\n", sequenceNumber)
				return lastSelectedSequence, common.NOT_SELECTED
			}
		}
	}

	// Check for a selected switch.
	// If this is a switch sequence also focus the last touched switch.
	if lastSelectedSequence == this.SwitchSequenceNumber {

		// Just send a message to focus the switch button.
		cmd := common.Command{
			Action: common.UpdateSwitch,
			Args: []common.Arg{
				{Name: "SwitchNumber", Value: lastSelectedSwitch},
				{Name: "SwitchPosition", Value: this.SwitchPosition[lastSelectedSwitch]},
				{Name: "Step", Value: false}, // Don't step the switch state.
				{Name: "Focus", Value: true}, // Focus the switch lamp.
			},
		}
		if commandChannels != nil {
			// Send a message to the switch sequence.
			common.SendCommandToSequence(this.SwitchSequenceNumber, cmd, commandChannels)
		}
	}

	// default to first sequnce if nothings running.
	return lastSelectedSequence, lastSelectedSwitch
}
