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

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/config"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
	"github.com/dhowlett99/dmxlights/pkg/presets"
	"github.com/oliread/usbdmx/ft232"
)

func loadConfig(sequences []*common.Sequence, this *CurrentState,
	X int, Y int, Red common.Color, PresetYellow common.Color,
	dmxController *ft232.DMXController, fixturesConfig *fixture.Fixtures,
	commandChannels []chan common.Command, eventsForLaunchpad chan common.ALight,
	guiButtons chan common.ALight, updateChannels []chan common.Sequence,
	dmxInterfacePresent bool) {

	// Stop all sequences, so we start in sync.
	cmd := common.Command{
		Action: common.Stop,
	}
	common.SendCommandToAllSequence(cmd, commandChannels)

	//AllFixturesOff(sequences, eventsForLaunchpad, guiButtons, dmxController, fixturesConfig, dmxInterfacePresent)

	// Load the config.
	// Which forces all sequences to load their config.
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
	common.LightLamp(common.ALight{X: 8, Y: 3, Brightness: this.MasterBrightness, Red: 255, Green: 255, Blue: 255}, eventsForLaunchpad, guiButtons)

	// Remember we selected this preset
	last := fmt.Sprint(X) + "," + fmt.Sprint(Y)
	this.LastPreset = &last

	// Get an upto date copy of all of the sequences.
	for sequenceNumber, sequence := range sequences {
		sequences[sequenceNumber] = common.RefreshSequence(sequenceNumber, commandChannels, updateChannels)

		// restore the speed, shift, size, fade, coordinates label data.
		this.Speed[sequenceNumber] = sequences[sequenceNumber].Speed
		this.RGBShift[sequenceNumber] = sequences[sequenceNumber].RGBShift
		this.ScannerShift[this.SelectedSequence] = sequences[sequenceNumber].ScannerShift
		this.RGBSize[sequenceNumber] = sequences[sequenceNumber].RGBSize
		this.ScannerSize[this.SelectedSequence] = sequences[sequenceNumber].ScannerSize
		this.RGBFade[sequenceNumber] = sequences[sequenceNumber].RGBFade
		this.ScannerCoordinates[sequenceNumber] = sequences[sequenceNumber].ScannerSelectedCoordinates
		this.Running[sequenceNumber] = sequences[sequenceNumber].Run
		this.Strobe[sequenceNumber] = sequences[sequenceNumber].Strobe
		this.StrobeSpeed[sequenceNumber] = sequences[sequenceNumber].StrobeSpeed

		// Setup the correct mode for the displays.
		this.SequenceType[sequenceNumber] = sequences[sequenceNumber].Type
		this.SelectMode[sequenceNumber] = NORMAL
		this.StaticFlashing[sequenceNumber] = false
		this.ScannerChaser[sequenceNumber] = sequences[sequenceNumber].ScannerChaser
		this.EditStaticColorsMode[sequenceNumber] = false

		// If the scanner sequence isn't running but the shutter chaser is, then it makes sense to show the shutter chaser.
		var displaySet bool
		if this.SequenceType[sequenceNumber] == "scanner" && !this.Running[this.ScannerSequenceNumber] && this.ScannerChaser[this.ScannerSequenceNumber] {
			// So adjust the mode to be CHASER_DISPLAY
			this.SelectMode[this.ScannerSequenceNumber] = CHASER_DISPLAY
			displayMode(sequenceNumber, this.SelectMode[sequenceNumber], this, sequences, eventsForLaunchpad, guiButtons, commandChannels)
		}
		if sequenceNumber != this.ChaserSequenceNumber && !displaySet {
			displayMode(sequenceNumber, this.SelectMode[sequenceNumber], this, sequences, eventsForLaunchpad, guiButtons, commandChannels)
		}

		// Reload the fixture state.
		for fixtureNumber := 0; fixtureNumber < sequences[this.SelectedSequence].NumberFixtures; fixtureNumber++ {
			this.FixtureState[sequenceNumber][fixtureNumber] = sequences[sequenceNumber].FixtureState[fixtureNumber]
		}

		// Restore the functions states from the sequence.
		if sequence.Type == "rgb" {
			this.Functions[sequenceNumber][common.Function2_Auto_Color].State = sequences[sequenceNumber].AutoColor
			this.Functions[sequenceNumber][common.Function3_Auto_Pattern].State = sequences[sequenceNumber].AutoPattern
			this.Functions[sequenceNumber][common.Function4_Bounce].State = sequences[sequenceNumber].Bounce
			this.Functions[sequenceNumber][common.Function6_Static_Gobo].State = sequences[sequenceNumber].Static
			this.EditStaticColorsMode[sequenceNumber] = sequences[sequenceNumber].Static
			this.Functions[sequenceNumber][common.Function7_Invert_Chase].State = sequences[sequenceNumber].RGBInvert
			this.Functions[sequenceNumber][common.Function8_Music_Trigger].State = sequences[sequenceNumber].MusicTrigger
		}
		if sequence.Type == "scanner" {
			this.Functions[sequenceNumber][common.Function2_Auto_Color].State = sequences[sequenceNumber].AutoColor
			this.Functions[sequenceNumber][common.Function3_Auto_Pattern].State = sequences[sequenceNumber].AutoPattern
			this.Functions[sequenceNumber][common.Function4_Bounce].State = sequences[sequenceNumber].Bounce
			this.Functions[sequenceNumber][common.Function7_Invert_Chase].State = sequences[sequenceNumber].ScannerChaser
			this.Functions[sequenceNumber][common.Function8_Music_Trigger].State = sequences[sequenceNumber].MusicTrigger
		}

		// If we are loading a switch sequence, update our local copy of the switch settings.
		if sequence.Type == "switch" {

			// Get an upto date copy of the switch sequence.
			sequences[sequenceNumber] = common.RefreshSequence(sequenceNumber, commandChannels, updateChannels)

			// Now set our local representation of switches
			for swiTchNumber, swiTch := range sequences[sequenceNumber].Switches {
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

	// Restore the master brightness, remember that the master is for all sequences in this loaded config.
	// So the master we retrive from this selected sequence will be the same for all the others.
	this.MasterBrightness = sequences[this.SelectedSequence].Master

	// Show the correct running and strobe buttons.
	if this.Strobe[this.SelectedSequence] {
		this.StrobeSpeed[this.SelectedSequence] = sequences[this.SelectedSequence].StrobeSpeed
	}
	// Show this sequence running status in the start/stop button.
	common.ShowRunningStatus(this.Running[this.SelectedSequence], eventsForLaunchpad, guiButtons)
	common.ShowStrobeButtonStatus(this.Strobe[this.SelectedSequence], eventsForLaunchpad, guiButtons)

}
