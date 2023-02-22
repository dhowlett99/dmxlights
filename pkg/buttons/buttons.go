// Copyright (C) 2022, 2023 dhowlett99.
// This is button processor, used by the launchpad and gui interfaces.
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
	"time"

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/config"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
	"github.com/dhowlett99/dmxlights/pkg/presets"
	"github.com/dhowlett99/dmxlights/pkg/sequence"
	"github.com/dhowlett99/dmxlights/pkg/sound"
	"github.com/oliread/usbdmx"
	"github.com/oliread/usbdmx/ft232"
	"github.com/rakyll/launchpad/mk3"
)

const debug = false

type CurrentState struct {
	SelectedSequence          int                          // The currently selected sequence.
	LastSelectedSequence      int                          // Store fof the last selected squence.
	Speed                     map[int]int                  // Local copy of sequence speed. Indexed by sequence.
	RGBShift                  map[int]int                  // Current rgb fixture shift. Indexed by sequence.
	ScannerShift              map[int]int                  // Current scanner shift for all fixtures.  Indexed by sequence
	RGBSize                   map[int]int                  // current RGB sequence this.Size[this.SelectedSequence]. Indexed by sequence
	ScannerSize               map[int]int                  // current scanner size for all fixtures. Indexed by sequence
	ScannerColor              int                          // current scanner color.
	RGBFade                   map[int]int                  // Indexed by sequence.
	ScannerFade               map[int]int                  // Indexed by sequence.
	ScannerCoordinates        map[int]int                  // Number of coordinates for scanner patterns is selected from 4 choices. ScannerCoordinates  0=12, 1=16,2=24,3=32,4=64, Indexed by sequence.
	Running                   map[int]bool                 // Which sequence is running. Indexed by sequence. True if running.
	Strobe                    map[int]bool                 // We are in strobe mode. True if strobing
	StrobeSpeed               map[int]int                  // Strobe speed. value is speed 0-255, indexed by sequence number.
	SavePreset                bool                         // Save a preset flag.
	Config                    bool                         // Flag to indicate we are in fixture config mode.
	Blackout                  bool                         // Blackout all fixtures.
	Flood                     bool                         // Flood all fixtures.
	FunctionSelectMode        []bool                       // Which sequence is in function selection mode.
	SelectButtonPressed       []bool                       // Which sequence has its Select button pressed.
	SwitchPositions           [9][9]int                    // Sorage for switch positions.
	EditSequenceColorsMode    []bool                       // This flag is true when the sequence is in sequence colors editing mode.
	EditScannerColorsMode     []bool                       // This flag is true when the sequence is in select scanner colors editing mode.
	EditGoboSelectionMode     []bool                       // This flag is true when the sequence is in sequence gobo selection mode.
	EditStaticColorsMode      []bool                       // This flag is true when the sequence is in static colors editing mode.
	EditPatternMode           []bool                       // This flag is true when the sequence is in pattern editing mode.
	EditFixtureSelectionMode  bool                         // This flag is true when the sequence is in select fixture mode.
	MasterBrightness          int                          // Affects all DMX fixtures and launchpad lamps.
	LastStaticColorButtonX    int                          // Which Static Color button did we change last.
	LastStaticColorButtonY    int                          // Which Static Color button did we change last.
	SoundGain                 float32                      // Fine gain -0.09 -> 0.09
	ScannerState              [][]common.ScannerState      // Which fixture is enabled: bool and inverted: bool on which sequence.
	SelectedFixture           int                          // Which fixture is selected when changing scanner color or gobo.
	FollowingAction           string                       // String to find next function, used in selecting a fixture.
	OffsetPan                 int                          // Offset for Pan.
	OffsetTilt                int                          // Offset for Tilt.
	Pad                       *mk3.Launchpad               // Pointer to the Novation Launchpad object.
	PresetsStore              map[string]presets.Preset    // Storage for the Presets.
	LastPreset                *string                      // Last preset used.
	SoundTriggers             []*common.Trigger            // Pointer to the Sound Triggers.
	SoundConfig               *sound.SoundConfig           // Pointer to the sound config struct.
	SequenceChannels          common.Channels              // Channles used to communicate with the sequence.
	Patterns                  map[int]common.Pattern       // A indexed map of the available patterns for this sequence.
	ScannerPattern            int                          // The selected scanner pattern Number. Used as the index for above.
	RGBPattern                int                          // The selected RGB pattern Number. Used as the index for above.
	StaticButtons             []common.StaticColorButton   // Storage for the color of the static buttons.
	SelectedGobo              int                          // The selected GOBO.
	ButtonTimer               *time.Time                   // Button Timer
	SelectColorBar            map[int]int                  // Storage for color bar in static color selection. Indexed by sequence.
	SwitchChannels            map[int]common.SwitchChannel // Used for communicating with mini-sequencers on switches.
	LaunchPadConnected        bool                         // Flag to indicate presence of Novation Launchpad.
	DmxInterfacePresent       bool                         // Flag to indicate precence of DMX interface card
	DmxInterfacePresentConfig *usbdmx.ControllerConfig     // DMX Interface card config.
	LaunchpadName             string                       // Storage for launchpad config.
}

func ProcessButtons(X int, Y int,
	sequences []*common.Sequence,
	this *CurrentState,
	eventsForLaunchpad chan common.ALight,
	guiButtons chan common.ALight,
	dmxController *ft232.DMXController,
	fixturesConfig *fixture.Fixtures,
	commandChannels []chan common.Command,
	replyChannels []chan common.Sequence,
	updateChannels []chan common.Sequence,
	gui bool) {

	if debug {
		fmt.Printf("ProcessButtons Called with X:%d Y:%d\n", X, Y)
	}

	// F L A S H   O N   B U T T O N S - Briefly light (flash) the fixtures based on color pattern.
	if X >= 0 &&
		X < 8 &&
		Y >= 0 &&
		Y < 4 &&
		!sequences[Y].Functions[common.Function1_Pattern].State &&
		!sequences[Y].Functions[common.Function6_Static_Gobo].State &&
		!sequences[Y].Functions[common.Function5_Color].State &&
		sequences[Y].Type != "switch" && // As long as we're not a switch sequence.
		sequences[Y].Type != "scanner" && // As long as we're not a scanner sequence.
		!this.FunctionSelectMode[Y] { // As long as we're not a scanner sequence for this sequence.

		if debug {
			fmt.Printf("Flash ON Fixture Pressed X:%d Y:%d\n", X, Y)
		}
		colorPattern := 5
		flashSequence := common.Sequence{
			Pattern: common.Pattern{
				Name:  "colors",
				Steps: this.Patterns[colorPattern].Steps, // Use the color pattern for flashing.
			},
		}

		red := flashSequence.Pattern.Steps[X].Fixtures[X].Colors[0].R
		green := flashSequence.Pattern.Steps[X].Fixtures[X].Colors[0].G
		blue := flashSequence.Pattern.Steps[X].Fixtures[X].Colors[0].B
		white := flashSequence.Pattern.Steps[X].Fixtures[X].Colors[0].W
		amber := flashSequence.Pattern.Steps[X].Fixtures[X].Colors[0].A
		uv := flashSequence.Pattern.Steps[X].Fixtures[X].Colors[0].UV
		pan := flashSequence.Pattern.Steps[X].Fixtures[X].Pan
		tilt := flashSequence.Pattern.Steps[X].Fixtures[X].Tilt
		shutter := flashSequence.Pattern.Steps[X].Fixtures[X].Shutter
		rotate := flashSequence.Pattern.Steps[X].Fixtures[X].Rotate
		music := flashSequence.Pattern.Steps[X].Fixtures[X].Music
		gobo := flashSequence.Pattern.Steps[X].Fixtures[X].Gobo
		program := flashSequence.Pattern.Steps[X].Fixtures[X].Program

		common.LightLamp(common.ALight{X: X, Y: Y, Brightness: this.MasterBrightness, Red: red, Green: green, Blue: blue}, eventsForLaunchpad, guiButtons)
		fixture.MapFixtures(Y, dmxController, X, red, green, blue, white, amber, uv, pan, tilt, shutter, rotate, music, program, gobo, 0, fixturesConfig, this.Blackout, this.MasterBrightness, this.MasterBrightness, this.Strobe[this.SelectedSequence], this.StrobeSpeed[this.SelectedSequence], this.DmxInterfacePresent)

		if gui {
			time.Sleep(200 * time.Millisecond)
			common.LightLamp(common.ALight{X: X, Y: Y, Brightness: 0, Red: 0, Green: 0, Blue: 0}, eventsForLaunchpad, guiButtons)
			fixture.MapFixtures(Y, dmxController, X, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, this.Blackout, this.MasterBrightness, this.MasterBrightness, this.Strobe[this.SelectedSequence], this.StrobeSpeed[this.SelectedSequence], this.DmxInterfacePresent)
		}

		return
	}

	// F L A S H  O F F   B U T T O N S - Briefly light (flash) the fixtures based on current pattern.
	if X >= 0 &&
		X != 108 && X != 117 &&
		X >= 100 && X < 117 &&
		Y >= 0 && Y < 4 &&
		!sequences[Y].Functions[common.Function1_Pattern].State &&
		!sequences[Y].Functions[common.Function6_Static_Gobo].State &&
		!sequences[Y].Functions[common.Function5_Color].State &&
		sequences[Y].Type != "switch" && // As long as we're not a switch sequence.
		sequences[Y].Type != "scanner" && // As long as we're not a scanner sequence.
		!this.FunctionSelectMode[Y] { // As long as we're not a scanner sequence for this sequence.

		if debug {
			fmt.Printf("Flash OFF Fixture Pressed X:%d Y:%d\n", X, Y)
		}

		X = X - 100

		common.LightLamp(common.ALight{X: X, Y: Y, Brightness: this.MasterBrightness, Red: 0, Green: 0, Blue: 0}, eventsForLaunchpad, guiButtons)
		fixture.MapFixtures(Y, dmxController, X, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, this.Blackout, this.MasterBrightness, this.MasterBrightness, this.Strobe[this.SelectedSequence], this.StrobeSpeed[this.SelectedSequence], this.DmxInterfacePresent)
		return
	}

	// P R E S E T S - recall (short press) or delete (long press) the preset.
	if X >= 100 && X < 108 &&
		(Y > 3 && Y < 7) {

		// Remove the button off offset.
		X = X - 100

		// We just pushed save this preset.
		if this.SavePreset {
			this.SavePreset = false
			return
		}

		// If this is a valid preset we are either recalling (short press) it or deleting it (long press)
		// If its not been set i.e. not valid we just ignore and return.
		if !this.PresetsStore[fmt.Sprint(X)+","+fmt.Sprint(Y)].State {
			return
		}

		// Stop the timer for this preset.
		elapsed := time.Since(*this.ButtonTimer)

		// Delete a preset - If the timer is longer than 1 seconds then we have a long press.
		if elapsed > 1*time.Second {

			// Delete the config file
			config.DeleteConfig(fmt.Sprintf("config%d.%d.json", X, Y))

			// Delete from preset store
			this.PresetsStore[fmt.Sprint(X)+","+fmt.Sprint(Y)] = presets.Preset{State: false, Selected: false, Label: ""}

			// Update the copy of presets on disk.
			presets.SavePresets(this.PresetsStore)

			// Show presets again.
			presets.RefreshPresets(eventsForLaunchpad, guiButtons, this.PresetsStore)

		} else {
			// Short press means load the config.
			loadConfig(sequences, this, X, Y, common.Red, common.PresetYellow, dmxController, fixturesConfig, commandChannels, eventsForLaunchpad, guiButtons, updateChannels, this.DmxInterfacePresent)
			this.SelectedSequence = 0
			// Indicate if this sequence is running.
			if this.Running[this.SelectedSequence] {
				common.LightLamp(common.ALight{X: 8, Y: 5, Brightness: this.MasterBrightness, Red: 0, Green: 255, Blue: 0}, eventsForLaunchpad, guiButtons)
			} else {
				common.LightLamp(common.ALight{X: 8, Y: 5, Brightness: this.MasterBrightness, Red: 255, Green: 255, Blue: 255}, eventsForLaunchpad, guiButtons)
			}
			this.SelectButtonPressed[this.SelectedSequence] = false
			HandleSelect(sequences, this, eventsForLaunchpad, commandChannels, guiButtons)
		}
		return
	}

	// C L E A R  - clear all from the GUI.
	if X == 0 && Y == -1 && gui {
		clear(X, Y, this, sequences, dmxController, fixturesConfig, commandChannels, eventsForLaunchpad, guiButtons, updateChannels)
		return
	}

	// C L E A R  - Start the timer, waiting for a long press to clear all.
	// Because a short press in scanner mode shifts the scanners up.
	if X == 0 && Y == -1 && !gui && sequences[this.SelectedSequence].Type == "scanner" {
		// Start a timer for this button.
		here := time.Now()
		this.ButtonTimer = &here
		return
	}

	//  C L E A R - clear all if we're not in the scanner mode.
	if X == 0 && Y == -1 && !gui && sequences[this.SelectedSequence].Type != "scanner" {
		clear(X, Y, this, sequences, dmxController, fixturesConfig, commandChannels, eventsForLaunchpad, guiButtons, updateChannels)
		return
	}

	// C L E A R  - We have a long press.
	if X == 100 && Y == -1 && !gui && sequences[this.SelectedSequence].Type == "scanner" {
		// Remove the off button offset.
		X = X - 100
		// Stop the timer for this preset.
		elapsed := time.Since(*this.ButtonTimer)
		// If the timer is longer than 1 seconds then we have a long press.
		if elapsed > 1*time.Second {
			clear(X, Y, this, sequences, dmxController, fixturesConfig, commandChannels, eventsForLaunchpad, guiButtons, updateChannels)
		} else {

			// S E L E C T   P O S I T I O N
			// UP ARROW
			if sequences[this.SelectedSequence].Type == "scanner" {

				if debug {
					fmt.Printf("UP ARROW\n")
				}

				buttonTouched(common.ALight{X: X, Y: Y, OnColor: common.Cyan, OffColor: common.White}, eventsForLaunchpad, guiButtons)

				this.OffsetTilt = this.OffsetTilt + 5

				if this.OffsetTilt > 255 {
					this.OffsetTilt = 255
				}
				// Clear the sequence colors for this sequence.
				cmd := common.Command{
					Action: common.UpdateOffsetTilt,
					Args: []common.Arg{
						{Name: "OffsetTilt", Value: this.OffsetTilt},
					},
				}
				common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

				// Update status bar.
				common.UpdateStatusBar(fmt.Sprintf("Tilt %02d", this.OffsetTilt), "tilt", false, guiButtons)
			}
		}
		return
	}

	// Swollow the button off events if not used for flash above.
	if X >= 100 {
		if debug {
			fmt.Printf("Swollow Event\n")
		}
		return
	}

	// F L O O D
	if X == 8 && Y == 3 && !this.SavePreset {

		// Shutdown any function bars.
		clearAllModes(sequences, this)
		HandleSelect(sequences, this, eventsForLaunchpad, commandChannels, guiButtons)

		if !this.Flood { // We're not already in flood so lets ask the sequence to flood.
			if debug {
				fmt.Printf("FLOOD ON\n")
			}
			// Find the currently selected preset and save it's location.
			for location, preset := range this.PresetsStore {
				if preset.State && preset.Selected {
					this.PresetsStore[location] = presets.Preset{State: preset.State, Selected: false, Label: preset.Label}
					presets.RefreshPresets(eventsForLaunchpad, guiButtons, this.PresetsStore)
					this.LastPreset = &location
					break
				}
			}
			floodOn(this, sequences, dmxController, fixturesConfig, commandChannels, eventsForLaunchpad, guiButtons, replyChannels)
			return
		}
		if this.Flood { // If we are flood already then tell the sequence to stop flood.
			if debug {
				fmt.Printf("FLOOD OFF\n")
			}
			// Restore the last preset
			if this.LastPreset != nil {
				lastPreset := this.PresetsStore[*this.LastPreset]
				this.PresetsStore[*this.LastPreset] = presets.Preset{State: lastPreset.State, Selected: true, Label: lastPreset.Label}
				presets.RefreshPresets(eventsForLaunchpad, guiButtons, this.PresetsStore)
			}
			floodOff(this, sequences, dmxController, fixturesConfig, commandChannels, eventsForLaunchpad, guiButtons, updateChannels)
			return
		}
	}

	// Sound sensitity up.
	if X == 4 && Y == -1 {
		if debug {
			fmt.Printf("Sound Up %f\n", this.SoundGain)
		}

		buttonTouched(common.ALight{X: X, Y: Y, OnColor: common.White, OffColor: common.Cyan}, eventsForLaunchpad, guiButtons)

		this.SoundGain = this.SoundGain - 0.01
		if this.SoundGain < -0.04 {
			this.SoundGain = -0.04
		}
		for _, trigger := range this.SoundTriggers {
			trigger.Gain = this.SoundGain
		}
		// Update the status bar
		sensitivity := common.FindSensitivity(this.SoundGain)
		common.UpdateStatusBar(fmt.Sprintf("Sensitivity %02d", sensitivity), "sensitivity", false, guiButtons)
		return
	}

	// Sound sensitity down.
	if X == 5 && Y == -1 {
		if debug {
			fmt.Printf("Sound Down%f\n", this.SoundGain)
		}

		buttonTouched(common.ALight{X: X, Y: Y, OnColor: common.White, OffColor: common.Cyan}, eventsForLaunchpad, guiButtons)

		this.SoundGain = this.SoundGain + 0.01
		if this.SoundGain > 0.09 {
			this.SoundGain = 0.09
		}
		for _, trigger := range this.SoundTriggers {
			trigger.Gain = this.SoundGain
		}
		// Update the status bar
		sensitivity := common.FindSensitivity(this.SoundGain)
		common.UpdateStatusBar(fmt.Sprintf("Sensitivity %02d", sensitivity), "sensitivity", false, guiButtons)
		return
	}

	// Master brightness down.
	if X == 6 && Y == -1 {

		if debug {
			fmt.Printf("Brightness Down \n")
		}

		buttonTouched(common.ALight{X: X, Y: Y, OnColor: common.White, OffColor: common.Cyan}, eventsForLaunchpad, guiButtons)

		this.MasterBrightness = this.MasterBrightness - 10
		if this.MasterBrightness < 0 {
			this.MasterBrightness = 0
		}
		cmd := common.Command{
			Action: common.Master,
			Args: []common.Arg{
				{Name: "Master", Value: this.MasterBrightness},
			},
		}
		common.SendCommandToAllSequence(cmd, commandChannels)

		// Update the status bar
		common.UpdateStatusBar(fmt.Sprintf("Master %02d", this.MasterBrightness), "master", false, guiButtons)
		return
	}

	// Master brightness up.
	if X == 7 && Y == -1 {

		if debug {
			fmt.Printf("Brightness Up \n")
		}

		buttonTouched(common.ALight{X: X, Y: Y, OnColor: common.White, OffColor: common.Cyan}, eventsForLaunchpad, guiButtons)

		this.MasterBrightness = this.MasterBrightness + 10
		if this.MasterBrightness > common.MaxDMXBrightness {
			this.MasterBrightness = common.MaxDMXBrightness
		}
		cmd := common.Command{
			Action: common.Master,
			Args: []common.Arg{
				{Name: "Master", Value: this.MasterBrightness},
			},
		}
		common.SendCommandToAllSequence(cmd, commandChannels)

		// Update the status bar
		common.UpdateStatusBar(fmt.Sprintf("Master %02d", this.MasterBrightness), "master", false, guiButtons)
		return
	}

	// Save mode.
	if X == 8 && Y == 4 {
		if debug {
			fmt.Printf("Save Mode\n")
		}

		if this.SavePreset { // Turn the save mode off.
			this.SavePreset = false
			presets.RefreshPresets(eventsForLaunchpad, guiButtons, this.PresetsStore)
			common.LightLamp(common.ALight{X: 8, Y: 4, Brightness: this.MasterBrightness, Red: 255, Green: 255, Blue: 255}, eventsForLaunchpad, guiButtons)
			return
		}
		this.SavePreset = true
		if this.Flood { // Turn off flood.
			floodOff(this, sequences, dmxController, fixturesConfig, commandChannels, eventsForLaunchpad, guiButtons, updateChannels)
		}
		presets.RefreshPresets(eventsForLaunchpad, guiButtons, this.PresetsStore)
		common.FlashLight(8, 4, common.Pink, common.White, eventsForLaunchpad, guiButtons)

		return
	}

	// P R E S E T S
	if X < 8 && (Y > 3 && Y < 7) {

		if debug {
			fmt.Printf("Ask For Config\n")
		}

		location := fmt.Sprint(X) + "," + fmt.Sprint(Y)

		if this.SavePreset {
			// S A V E - Ask all sequences for their current config and save in a file.

			current := this.PresetsStore[location]
			this.PresetsStore[location] = presets.Preset{State: true, Selected: true, Label: current.Label}
			this.LastPreset = &location
			common.LightLamp(common.ALight{X: X, Y: Y, Brightness: this.MasterBrightness, Red: 255, Green: 0, Blue: 0}, eventsForLaunchpad, guiButtons)
			config.AskToSaveConfig(commandChannels, replyChannels, X, Y)

			// turn off the save button from flashing.
			common.LightLamp(common.ALight{X: 8, Y: 4, Brightness: this.MasterBrightness, Red: 255, Green: 255, Blue: 255}, eventsForLaunchpad, guiButtons)

			presets.SavePresets(this.PresetsStore)

			// clear any selected preset.
			for location, preset := range this.PresetsStore {
				if preset.State && preset.Selected {
					this.PresetsStore[location] = presets.Preset{State: preset.State, Selected: false, Label: preset.Label}
				}
			}

			// Select this location and flash its button.
			this.PresetsStore[location] = presets.Preset{State: true, Selected: true, Label: current.Label}
			presets.RefreshPresets(eventsForLaunchpad, guiButtons, this.PresetsStore)

			if gui {
				this.SavePreset = false
			}

		} else {
			// L O A D - Load config, but only if it exists in the presets map.
			if this.PresetsStore[location].State {

				if gui { // GUI path.
					if this.SavePreset {
						this.SavePreset = false
					}
					loadConfig(sequences, this, X, Y, common.Red, common.PresetYellow, dmxController, fixturesConfig, commandChannels, eventsForLaunchpad, guiButtons, updateChannels, this.DmxInterfacePresent)
					this.SelectedSequence = 0
					// Indicate if this sequence is running.
					if this.Running[this.SelectedSequence] {
						common.LightLamp(common.ALight{X: 8, Y: 5, Brightness: this.MasterBrightness, Red: 0, Green: 255, Blue: 0}, eventsForLaunchpad, guiButtons)
					} else {
						common.LightLamp(common.ALight{X: 8, Y: 5, Brightness: this.MasterBrightness, Red: 255, Green: 255, Blue: 255}, eventsForLaunchpad, guiButtons)
					}
					this.SelectButtonPressed[this.SelectedSequence] = false
					HandleSelect(sequences, this, eventsForLaunchpad, commandChannels, guiButtons)
				} else { // Launchpad path.
					// This is a valid preset we might be trying to load it or delete it.
					// Start a timer for this button.
					here := time.Now()
					this.ButtonTimer = &here

					// And wait for the button release.
				}
			}
		}
		return
	}

	// Decrease Shift.
	if X == 2 && Y == 7 {

		if debug {
			fmt.Printf("Decrease Shift\n")
		}

		buttonTouched(common.ALight{X: X, Y: Y, OnColor: common.White, OffColor: common.Cyan}, eventsForLaunchpad, guiButtons)

		if sequences[this.SelectedSequence].Type == "rgb" {
			this.RGBShift[this.SelectedSequence] = this.RGBShift[this.SelectedSequence] - 1
			if this.RGBShift[this.SelectedSequence] < 0 {
				this.RGBShift[this.SelectedSequence] = 0
			}
			cmd := common.Command{
				Action: common.UpdateRGBShift,
				Args: []common.Arg{
					{Name: "RGBShift", Value: this.RGBShift[this.SelectedSequence]},
				},
			}
			common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

			// Update the status bar
			common.UpdateStatusBar(fmt.Sprintf("Shift %02d", this.RGBShift[this.SelectedSequence]), "shift", false, guiButtons)
			return
		}

		if sequences[this.SelectedSequence].Type == "scanner" {
			this.ScannerShift[this.SelectedSequence] = this.ScannerShift[this.SelectedSequence] - 1
			if this.ScannerShift[this.SelectedSequence] < 0 {
				this.ScannerShift[this.SelectedSequence] = 0
			}
			cmd := common.Command{
				Action: common.UpdateScannerShift,
				Args: []common.Arg{
					{Name: "ScannerShift", Value: this.ScannerShift[this.SelectedSequence]},
				},
			}
			common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

			// Update the status bar
			label := getScannerShiftLabel(this.ScannerShift[this.SelectedSequence])
			common.UpdateStatusBar(fmt.Sprintf("Shift %0s", label), "shift", false, guiButtons)
			return
		}
	}

	// Increase Shift.
	if X == 3 && Y == 7 {

		if debug {
			fmt.Printf("Increase Shift \n")
		}

		buttonTouched(common.ALight{X: X, Y: Y, OnColor: common.White, OffColor: common.Cyan}, eventsForLaunchpad, guiButtons)

		if sequences[this.SelectedSequence].Type == "rgb" {
			this.RGBShift[this.SelectedSequence] = this.RGBShift[this.SelectedSequence] + 1
			if this.RGBShift[this.SelectedSequence] > 50 {
				this.RGBShift[this.SelectedSequence] = 50
			}
			cmd := common.Command{
				Action: common.UpdateRGBShift,
				Args: []common.Arg{
					{Name: "Shift", Value: this.RGBShift[this.SelectedSequence]},
				},
			}
			common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

			// Update the status bar
			common.UpdateStatusBar(fmt.Sprintf("Shift %02d", this.RGBShift[this.SelectedSequence]), "shift", false, guiButtons)
			return
		}

		if sequences[this.SelectedSequence].Type == "scanner" {
			this.ScannerShift[this.SelectedSequence] = this.ScannerShift[this.SelectedSequence] + 1
			if this.ScannerShift[this.SelectedSequence] > 3 {
				this.ScannerShift[this.SelectedSequence] = 3
			}
			cmd := common.Command{
				Action: common.UpdateScannerShift,
				Args: []common.Arg{
					{Name: "ScannerShift", Value: this.ScannerShift[this.SelectedSequence]},
				},
			}
			common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

			// Update the status bar
			label := getScannerShiftLabel(this.ScannerShift[this.SelectedSequence])
			common.UpdateStatusBar(fmt.Sprintf("Shift %s", label), "shift", false, guiButtons)
			return
		}
	}

	// Decrease speed of selected sequence.
	if X == 0 && Y == 7 {

		if debug {
			fmt.Printf("Decrease Speed \n")
		}

		buttonTouched(common.ALight{X: X, Y: Y, OnColor: common.White, OffColor: common.Cyan}, eventsForLaunchpad, guiButtons)

		if this.Strobe[this.SelectedSequence] {
			this.StrobeSpeed[this.SelectedSequence] -= 10
			if this.StrobeSpeed[this.SelectedSequence] < 0 {
				this.StrobeSpeed[this.SelectedSequence] = 0
			}
			cmd := common.Command{
				Action: common.UpdateStrobeSpeed,
				Args: []common.Arg{
					{Name: "STROBE_SPEED", Value: this.StrobeSpeed[this.SelectedSequence]},
				},
			}
			common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)
			// Update the status bar
			common.UpdateStatusBar(fmt.Sprintf("Strobe %02d", this.StrobeSpeed[this.SelectedSequence]), "speed", false, guiButtons)
			return
		}

		// Get an upto date copy of the sequence.
		sequences[this.SelectedSequence] = common.RefreshSequence(this.SelectedSequence, commandChannels, updateChannels)

		if !sequences[this.SelectedSequence].MusicTrigger {
			this.Speed[this.SelectedSequence]--
			if this.Speed[this.SelectedSequence] < 0 {
				this.Speed[this.SelectedSequence] = 1
			}
			cmd := common.Command{
				Action: common.UpdateSpeed,
				Args: []common.Arg{
					{Name: "Speed", Value: this.Speed[this.SelectedSequence]},
				},
			}
			common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)
		}

		// Update the status bar
		common.UpdateStatusBar(fmt.Sprintf("Speed %02d", this.Speed[this.SelectedSequence]), "speed", false, guiButtons)

		return
	}

	// Increase speed of selected sequence.
	if X == 1 && Y == 7 {

		if debug {
			fmt.Printf("Increase Speed \n")
		}

		if this.Strobe[this.SelectedSequence] {
			this.StrobeSpeed[this.SelectedSequence] += 10
			if this.StrobeSpeed[this.SelectedSequence] > 255 {
				this.StrobeSpeed[this.SelectedSequence] = 255
			}
			cmd := common.Command{
				Action: common.UpdateStrobeSpeed,
				Args: []common.Arg{
					{Name: "STROBE_SPEED", Value: this.StrobeSpeed[this.SelectedSequence]},
				},
			}
			common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)
			// Update the status bar
			common.UpdateStatusBar(fmt.Sprintf("Strobe %02d", this.StrobeSpeed[this.SelectedSequence]), "speed", false, guiButtons)
			return
		}

		buttonTouched(common.ALight{X: X, Y: Y, OnColor: common.White, OffColor: common.Cyan}, eventsForLaunchpad, guiButtons)

		// Get an upto date copy of the sequence.
		sequences[this.SelectedSequence] = common.RefreshSequence(this.SelectedSequence, commandChannels, updateChannels)

		if !sequences[this.SelectedSequence].MusicTrigger {
			this.Speed[this.SelectedSequence]++
			if this.Speed[this.SelectedSequence] > 12 {
				this.Speed[this.SelectedSequence] = 12
			}
			cmd := common.Command{
				Action: common.UpdateSpeed,
				Args: []common.Arg{
					{Name: "Speed", Value: this.Speed[this.SelectedSequence]},
				},
			}
			common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)
		}

		// Update the status bar
		common.UpdateStatusBar(fmt.Sprintf("Speed %02d", this.Speed[this.SelectedSequence]), "speed", false, guiButtons)

		return
	}

	// S E L E C T    S E Q U E N C E.
	// Select sequence 1.
	if X == 8 && Y == 0 {

		this.SelectedSequence = 0

		if debug {
			fmt.Printf("Select Sequence %d \n", this.SelectedSequence)
		}

		HandleSelect(sequences, this, eventsForLaunchpad, commandChannels, guiButtons)
		cmd := common.Command{
			Action: common.PlayStaticOnce,
		}
		common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

		this.EditSequenceColorsMode[this.SelectedSequence] = false
		this.EditGoboSelectionMode[this.SelectedSequence] = false

		if this.LaunchPadConnected {
			this.Pad.Program()
		}

		return
	}

	// Select sequence 2.
	if X == 8 && Y == 1 {

		this.SelectedSequence = 1

		if debug {
			fmt.Printf("Select Sequence %d \n", this.SelectedSequence)
		}

		HandleSelect(sequences, this, eventsForLaunchpad, commandChannels, guiButtons)

		cmd := common.Command{
			Action: common.PlayStaticOnce,
		}
		common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

		this.EditSequenceColorsMode[this.SelectedSequence] = false
		this.EditGoboSelectionMode[this.SelectedSequence] = false

		if this.LaunchPadConnected {
			this.Pad.Program()
		}

		return
	}

	// Select sequence 3.
	if X == 8 && Y == 2 {

		this.SelectedSequence = 2

		if debug {
			fmt.Printf("Select Sequence %d \n", this.SelectedSequence)
		}

		HandleSelect(sequences, this, eventsForLaunchpad, commandChannels, guiButtons)

		cmd := common.Command{
			Action: common.PlayStaticOnce,
		}
		common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

		this.EditSequenceColorsMode[this.SelectedSequence] = false
		this.EditGoboSelectionMode[this.SelectedSequence] = false

		if this.LaunchPadConnected {
			this.Pad.Program()
		}

		return
	}

	// Select sequence 4.
	if X == 8 && Y == 3 {

		this.SelectedSequence = 3

		if debug {
			fmt.Printf("Select Sequence %d \n", this.SelectedSequence)
		}

		HandleSelect(sequences, this, eventsForLaunchpad, commandChannels, guiButtons)

		cmd := common.Command{
			Action: common.PlayStaticOnce,
		}
		common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

		this.EditSequenceColorsMode[this.SelectedSequence] = false
		this.EditGoboSelectionMode[this.SelectedSequence] = false

		return
	}

	// S T A R T - Start sequence.
	if X == 8 && Y == 5 {

		// S T O P - If sequence is running, stop it
		if this.Running[this.SelectedSequence] {
			if debug {
				fmt.Printf("Stop Sequence %d \n", this.SelectedSequence)
			}
			cmd := common.Command{
				Action: common.Stop,
				Args: []common.Arg{
					{Name: "Speed", Value: this.Speed[this.SelectedSequence]},
				},
			}
			common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)
			common.LightLamp(common.ALight{X: X, Y: Y, Brightness: this.MasterBrightness, Red: 255, Green: 255, Blue: 255}, eventsForLaunchpad, guiButtons)

			// The scanner sequence has stopped, so show the status of the scanners.
			if sequences[this.SelectedSequence].Type == "scanner" {
				// Show the status.
				ShowScannerStatus(this.SelectedSequence, *sequences[this.SelectedSequence], this, eventsForLaunchpad, guiButtons, commandChannels)
			}

			this.Running[this.SelectedSequence] = false
			return
		} else {
			// Start this sequence.
			if debug {
				fmt.Printf("Start Sequence %d \n", Y)
			}
			sequences[this.SelectedSequence].MusicTrigger = false
			cmd := common.Command{
				Action: common.Start,
				Args: []common.Arg{
					{Name: "Speed", Value: this.Speed[this.SelectedSequence]},
				},
			}
			common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)
			common.LightLamp(common.ALight{X: X, Y: Y, Brightness: this.MasterBrightness, Red: 0, Green: 255, Blue: 0}, eventsForLaunchpad, guiButtons)
			this.Running[this.SelectedSequence] = true
			return
		}

	}

	// S T R O B E - Strobe.
	if X == 8 && Y == 6 {

		// Shutdown any function bars.
		clearAllModes(sequences, this)
		HandleSelect(sequences, this, eventsForLaunchpad, commandChannels, guiButtons)

		// If strobing, stop it
		if this.Strobe[this.SelectedSequence] {
			// Stop strobing this sequence.
			cmd := common.Command{
				Action: common.StopStrobe,
			}
			common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)
			common.ShowStrobeButtonStatus(false, eventsForLaunchpad, guiButtons)
			this.Strobe[this.SelectedSequence] = false
			this.StrobeSpeed[this.SelectedSequence] = 0
			// Update the status bar
			common.UpdateStatusBar(fmt.Sprintf("Speed %02d", this.Speed[this.SelectedSequence]), "speed", false, guiButtons)
			return

		} else {
			// Start strobing for this sequence.
			this.Strobe[this.SelectedSequence] = true
			this.StrobeSpeed[this.SelectedSequence] = 255
			cmd := common.Command{
				Action: common.Strobe,
				Args: []common.Arg{
					{Name: "STROBE_SPEED", Value: this.StrobeSpeed[this.SelectedSequence]},
				},
			}
			// Store the strobe flag in all sequences.
			common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)
			common.ShowStrobeButtonStatus(true, eventsForLaunchpad, guiButtons)

			// Update the status bar
			common.UpdateStatusBar(fmt.Sprintf("Strobe %02d", this.StrobeSpeed[this.SelectedSequence]), "speed", false, guiButtons)
			return
		}
	}

	// Size decrease.
	if X == 4 && Y == 7 {

		if debug {
			fmt.Printf("Decrease Size\n")
		}

		buttonTouched(common.ALight{X: X, Y: Y, OnColor: common.White, OffColor: common.Cyan}, eventsForLaunchpad, guiButtons)

		if sequences[this.SelectedSequence].Type == "rgb" {
			// Send Update RGB Size.
			this.RGBSize[this.SelectedSequence]--
			if this.RGBSize[this.SelectedSequence] < 1 {
				this.RGBSize[this.SelectedSequence] = 1
			}
			cmd := common.Command{
				Action: common.UpdateRGBSize,
				Args: []common.Arg{
					{Name: "RGBSize", Value: this.RGBSize[this.SelectedSequence]},
				},
			}
			common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)
			// Update the status bar
			common.UpdateStatusBar(fmt.Sprintf("Size %02d", this.RGBSize[this.SelectedSequence]), "size", false, guiButtons)
			return
		}

		if sequences[this.SelectedSequence].Type == "scanner" {
			// Send Update Scanner Size.
			this.ScannerSize[this.SelectedSequence] = this.ScannerSize[this.SelectedSequence] - 10
			if this.ScannerSize[this.SelectedSequence] < 0 {
				this.ScannerSize[this.SelectedSequence] = 0
			}
			cmd := common.Command{
				Action: common.UpdateScannerSize,
				Args: []common.Arg{
					{Name: "ScannerSize", Value: this.ScannerSize[this.SelectedSequence]},
				},
			}
			common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

			// Update the status bar
			common.UpdateStatusBar(fmt.Sprintf("Size %02d", this.ScannerSize[this.SelectedSequence]), "size", false, guiButtons)
			return
		}
	}

	// Increase Size.
	if X == 5 && Y == 7 {

		if debug {
			fmt.Printf("Increase Size\n")
		}

		buttonTouched(common.ALight{X: X, Y: Y, OnColor: common.White, OffColor: common.Cyan}, eventsForLaunchpad, guiButtons)

		if sequences[this.SelectedSequence].Type == "rgb" {
			// Send Update RGB Size.
			this.RGBSize[this.SelectedSequence]++
			if this.RGBSize[this.SelectedSequence] > common.MaxRGBSize {
				this.RGBSize[this.SelectedSequence] = common.MaxRGBSize
			}
			cmd := common.Command{
				Action: common.UpdateRGBSize,
				Args: []common.Arg{
					{Name: "RGBSize", Value: this.RGBSize[this.SelectedSequence]},
				},
			}
			common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)
			// Update the status bar
			common.UpdateStatusBar(fmt.Sprintf("Size %02d", this.RGBSize[this.SelectedSequence]), "size", false, guiButtons)
			return
		}

		if sequences[this.SelectedSequence].Type == "scanner" {
			// Send Update Scanner Size.
			this.ScannerSize[this.SelectedSequence] = this.ScannerSize[this.SelectedSequence] + 10
			if this.ScannerSize[this.SelectedSequence] > common.MaxScannerSize {
				this.ScannerSize[this.SelectedSequence] = common.MaxScannerSize
			}
			cmd := common.Command{
				Action: common.UpdateScannerSize,
				Args: []common.Arg{
					{Name: "ScannerSize", Value: this.ScannerSize[this.SelectedSequence]},
				},
			}
			common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

			// Update the status bar
			common.UpdateStatusBar(fmt.Sprintf("Size %02d", this.ScannerSize[this.SelectedSequence]), "size", false, guiButtons)

			return
		}
	}

	// Fade time decrease.
	if X == 6 && Y == 7 {

		if debug {
			fmt.Printf("Decrease Fade Time\n")
		}

		buttonTouched(common.ALight{X: X, Y: Y, OnColor: common.White, OffColor: common.Cyan}, eventsForLaunchpad, guiButtons)

		if sequences[this.SelectedSequence].Type == "rgb" {
			this.RGBFade[this.SelectedSequence]--
			if this.RGBFade[this.SelectedSequence] < 1 {
				this.RGBFade[this.SelectedSequence] = 1
			}
			// Send fade update command.
			cmd := common.Command{
				Action: common.UpdateRGBFadeSpeed,
				Args: []common.Arg{
					{Name: "RGBFadeSpeed", Value: this.RGBFade[this.SelectedSequence]},
				},
			}
			common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)
			// Update the status bar
			common.UpdateStatusBar(fmt.Sprintf("Fade %02d", this.RGBFade[this.SelectedSequence]), "fade", false, guiButtons)

			return
		}

		// Update Coordinates.
		if sequences[this.SelectedSequence].Type == "scanner" {
			// Fade also send more or less coordinates for the scanner patterns.
			this.ScannerCoordinates[this.SelectedSequence]--
			if this.ScannerCoordinates[this.SelectedSequence] < 0 {
				this.ScannerCoordinates[this.SelectedSequence] = 0
			}
			cmd := common.Command{
				Action: common.UpdateNumberCoordinates,
				Args: []common.Arg{
					{Name: "NumberCoordinates", Value: this.ScannerCoordinates[this.SelectedSequence]},
				},
			}
			common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)
			// Update the status bar
			label := getScannerCoordinatesLabel(this.ScannerCoordinates[this.SelectedSequence])
			common.UpdateStatusBar(fmt.Sprintf("Coord %s", label), "fade", false, guiButtons)
			return
		}

	}

	// Fade time increase.
	if X == 7 && Y == 7 {

		if debug {
			fmt.Printf("Increase Fade Time\n")
		}

		buttonTouched(common.ALight{X: X, Y: Y, OnColor: common.White, OffColor: common.Cyan}, eventsForLaunchpad, guiButtons)

		if sequences[this.SelectedSequence].Type == "rgb" {
			this.RGBFade[this.SelectedSequence]++
			if this.RGBFade[this.SelectedSequence] > common.MaxRGBFade {
				this.RGBFade[this.SelectedSequence] = common.MaxRGBFade
			}
			// Send fade update command.
			cmd := common.Command{
				Action: common.UpdateRGBFadeSpeed,
				Args: []common.Arg{
					{Name: "FadeSpeed", Value: this.RGBFade[this.SelectedSequence]},
				},
			}
			common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)
			// Update the status bar
			common.UpdateStatusBar(fmt.Sprintf("Fade %02d", this.RGBFade[this.SelectedSequence]), "fade", false, guiButtons)
			return
		}

		// Update Coordinates.
		if sequences[this.SelectedSequence].Type == "scanner" {
			// Fade also send more or less coordinates for the scanner patterns.
			this.ScannerCoordinates[this.SelectedSequence]++
			if this.ScannerCoordinates[this.SelectedSequence] > 4 {
				this.ScannerCoordinates[this.SelectedSequence] = 4
			}
			cmd := common.Command{
				Action: common.UpdateNumberCoordinates,
				Args: []common.Arg{
					{Name: "NumberCoordinates", Value: this.ScannerCoordinates[this.SelectedSequence]},
				},
			}
			common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)
			// Update the status bar
			label := getScannerCoordinatesLabel(this.ScannerCoordinates[this.SelectedSequence])
			common.UpdateStatusBar(fmt.Sprintf("Coord %s", label), "fade", false, guiButtons)
			return
		}
	}

	// S W I T C H   B U T T O N's Toggle State of switches for this sequence.
	if X >= 0 && X < 8 && !this.FunctionSelectMode[this.SelectedSequence] &&
		Y >= 0 &&
		Y < 4 &&
		sequences[Y].Type == "switch" {

		if debug {
			fmt.Printf("Switch Key X:%d Y:%d\n", X, Y)
		}

		// Get an upto date copy of the switch information by updating our copy of the switch sequence.
		sequences[this.SelectedSequence] = common.RefreshSequence(this.SelectedSequence, commandChannels, updateChannels)

		// We have a valid switch.
		if X < len(sequences[Y].Switches) {
			this.SwitchPositions[Y][X] = this.SwitchPositions[Y][X] + 1
			valuesLength := len(sequences[Y].Switches[X].States)
			if this.SwitchPositions[Y][X] == valuesLength {
				this.SwitchPositions[Y][X] = 0
			}

			// Send a message to the sequence for it to toggle the selected switch.
			// Y is the sequence.
			// X is the switch.
			cmd := common.Command{
				Action: common.UpdateSwitch,
				Args: []common.Arg{
					{Name: "SwitchNumber", Value: X},
					{Name: "SwitchPosition", Value: this.SwitchPositions[Y][X]},
				},
			}
			// Send a message to the switch sequence.
			common.SendCommandToAllSequenceOfType(sequences, cmd, commandChannels, "switch")
		}
	}

	// D I S A B L E   F I X T U R E  - Used to toggle the scanner state from on, inverted or off.
	if X >= 0 && X < 8 && !this.FunctionSelectMode[this.SelectedSequence] &&
		Y >= 0 &&
		Y < 4 &&
		!sequences[this.SelectedSequence].Functions[common.Function1_Pattern].State &&
		!sequences[this.SelectedSequence].Functions[common.Function6_Static_Gobo].State &&
		!sequences[this.SelectedSequence].Functions[common.Function5_Color].State &&
		sequences[Y].Type == "scanner" {

		if debug {
			fmt.Printf("Disable Fixture X:%d Y:%d\n", X, Y)
			fmt.Printf("Fixture State Enabled %t  Inverted %t\n", this.ScannerState[X][Y].Enabled, this.ScannerState[X][Y].Inverted)
		}

		// Disable scanner if we're already enabled and inverted.
		if this.ScannerState[X][Y].Enabled && this.ScannerState[X][Y].Inverted && X < sequences[Y].NumberFixtures {
			if debug {
				fmt.Printf("Disable Scanner Number %d State on Sequence %d to false\n", X, Y)
			}

			this.ScannerState[X][Y].Enabled = false
			this.ScannerState[X][Y].Inverted = false

			// Tell the sequence to turn on this scanner.
			cmd := common.Command{
				Action: common.ToggleFixtureState,
				Args: []common.Arg{
					{Name: "SequenceNumber", Value: Y},
					{Name: "FixtureNumber", Value: X},
					{Name: "FixtureState", Value: false},
					{Name: "FixtureInverted", Value: false},
				},
			}
			common.SendCommandToSequence(Y, cmd, commandChannels)

			// Show the status.
			ShowScannerStatus(Y, *sequences[Y], this, eventsForLaunchpad, guiButtons, commandChannels)

			return
		}

		// Enable scanner if not enabled and not inverted.
		if !this.ScannerState[X][Y].Enabled && !this.ScannerState[X][Y].Inverted && X < sequences[Y].NumberFixtures {

			if debug {
				fmt.Printf("Enable Scanner Number %d State on Sequence %d to true [Scanners:%d]\n", X, Y, sequences[Y].NumberFixtures)
			}

			this.ScannerState[X][Y].Enabled = true
			this.ScannerState[X][Y].Inverted = false

			// Tell the sequence to turn off this scanner.
			cmd := common.Command{
				Action: common.ToggleFixtureState,
				Args: []common.Arg{
					{Name: "SequenceNumber", Value: Y},
					{Name: "FixtureNumber", Value: X},
					{Name: "FixtureState", Value: true},
					{Name: "FixtureInverted", Value: false},
				},
			}
			common.SendCommandToSequence(Y, cmd, commandChannels)

			// Show the status.
			ShowScannerStatus(Y, *sequences[Y], this, eventsForLaunchpad, guiButtons, commandChannels)

			return

		}

		// Invert scanner if we're enabled but not inverted.
		if this.ScannerState[X][Y].Enabled && !this.ScannerState[X][Y].Inverted && X < sequences[Y].NumberFixtures {
			if debug {
				fmt.Printf("Invert Scanner Number %d State on Sequence %d to false\n", X, Y)
			}

			this.ScannerState[X][Y].Enabled = true
			this.ScannerState[X][Y].Inverted = true

			// Tell the sequence to turn on this scanner.
			cmd := common.Command{
				Action: common.ToggleFixtureState,
				Args: []common.Arg{
					{Name: "SequenceNumber", Value: Y},
					{Name: "FixtureNumber", Value: X},
					{Name: "FixtureState", Value: true},
					{Name: "FixtureInverted", Value: true},
				},
			}
			common.SendCommandToSequence(Y, cmd, commandChannels)

			// Show the status.
			ShowScannerStatus(Y, *sequences[Y], this, eventsForLaunchpad, guiButtons, commandChannels)

			return
		}

	}

	// DOWN ARROW
	if X == 1 && Y == -1 && sequences[this.SelectedSequence].Type == "scanner" {
		if debug {
			fmt.Printf("DOWN ARROW\n")
		}

		buttonTouched(common.ALight{X: X, Y: Y, OnColor: common.Cyan, OffColor: common.White}, eventsForLaunchpad, guiButtons)

		this.OffsetTilt = this.OffsetTilt - 5

		if this.OffsetTilt < 0 {
			this.OffsetTilt = 0
		}
		// Clear the sequence colors for this sequence.
		cmd := common.Command{
			Action: common.UpdateOffsetTilt,
			Args: []common.Arg{
				{Name: "OffsetTilt", Value: this.OffsetTilt},
			},
		}
		common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

		// Update status bar.
		common.UpdateStatusBar(fmt.Sprintf("Tilt %02d", this.OffsetTilt), "tilt", false, guiButtons)

		return
	}

	// LEFT ARROW
	if X == 2 && Y == -1 && sequences[this.SelectedSequence].Type == "scanner" {
		if debug {
			fmt.Printf("LEFT ARROW\n")
		}

		buttonTouched(common.ALight{X: X, Y: Y, OnColor: common.Cyan, OffColor: common.White}, eventsForLaunchpad, guiButtons)

		this.OffsetPan = this.OffsetPan - 5

		if this.OffsetPan < 0 {
			this.OffsetPan = 0
		}
		// Clear the sequence colors for this sequence.
		cmd := common.Command{
			Action: common.UpdateOffsetPan,
			Args: []common.Arg{
				{Name: "OffsetPan", Value: this.OffsetPan},
			},
		}
		common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

		// Update status bar.
		common.UpdateStatusBar(fmt.Sprintf("Pan %02d", this.OffsetPan), "pan", false, guiButtons)

		return
	}

	// RIGHT ARROW
	if X == 3 && Y == -1 && sequences[this.SelectedSequence].Type == "scanner" {
		if debug {
			fmt.Printf("RIGHT ARROW\n")
		}

		buttonTouched(common.ALight{X: X, Y: Y, OnColor: common.Cyan, OffColor: common.White}, eventsForLaunchpad, guiButtons)

		this.OffsetPan = this.OffsetPan + 5

		if this.OffsetPan > 255 {
			this.OffsetPan = 255
		}
		// Clear the sequence colors for this sequence.
		cmd := common.Command{
			Action: common.UpdateOffsetPan,
			Args: []common.Arg{
				{Name: "OffsetPan", Value: this.OffsetPan},
			},
		}
		common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

		// Update status bar.
		common.UpdateStatusBar(fmt.Sprintf("Pan %02d", this.OffsetPan), "pan", false, guiButtons)

		return

	}

	// S E L E C T   S T A T I C   C O L O R
	// Red
	if X == 1 && Y == -1 && sequences[this.SelectedSequence].Type != "scanner" {

		if sequences[this.SelectedSequence].Functions[common.Function6_Static_Gobo].State {

			if debug {
				fmt.Printf("Choose Static Red X:%d Y:%d\n", X, Y)
			}

			buttonTouched(common.ALight{X: X, Y: Y, OnColor: common.White, OffColor: common.Red}, eventsForLaunchpad, guiButtons)

			this.StaticButtons[this.SelectedSequence].X = this.LastStaticColorButtonX
			this.StaticButtons[this.SelectedSequence].Y = this.LastStaticColorButtonY
			this.StaticButtons[this.SelectedSequence].Color.R = this.StaticButtons[this.SelectedSequence].Color.R - 10
			if this.StaticButtons[this.SelectedSequence].Color.R > 255 {
				this.StaticButtons[this.SelectedSequence].Color.R = 0
			}
			if this.StaticButtons[this.SelectedSequence].Color.R < 0 {
				this.StaticButtons[this.SelectedSequence].Color.R = 255
			}

			common.LightLamp(common.ALight{X: X, Y: Y, Brightness: this.MasterBrightness, Red: this.StaticButtons[this.SelectedSequence].Color.R, Green: 0, Blue: 0}, eventsForLaunchpad, guiButtons)
			updateStaticLamp(this.SelectedSequence, this.StaticButtons[this.SelectedSequence], commandChannels)

			// Update the status bar
			common.UpdateStatusBar(fmt.Sprintf("Red %02d", this.StaticButtons[this.SelectedSequence].Color.R), "red", false, guiButtons)
			return
		}
	}

	// Green
	if X == 2 && Y == -1 && sequences[this.SelectedSequence].Type != "scanner" {

		if sequences[this.SelectedSequence].Functions[common.Function6_Static_Gobo].State {
			if debug {
				fmt.Printf("Choose Static Green X:%d Y:%d\n", X, Y)
			}

			buttonTouched(common.ALight{X: X, Y: Y, OnColor: common.White, OffColor: common.Green}, eventsForLaunchpad, guiButtons)

			this.StaticButtons[this.SelectedSequence].X = this.LastStaticColorButtonX
			this.StaticButtons[this.SelectedSequence].Y = this.LastStaticColorButtonY
			this.StaticButtons[this.SelectedSequence].Color.G = this.StaticButtons[this.SelectedSequence].Color.G - 10
			if this.StaticButtons[this.SelectedSequence].Color.G > 255 {
				this.StaticButtons[this.SelectedSequence].Color.G = 0
			}
			if this.StaticButtons[this.SelectedSequence].Color.G < 0 {
				this.StaticButtons[this.SelectedSequence].Color.G = 255
			}
			common.LightLamp(common.ALight{X: X, Y: Y, Brightness: this.MasterBrightness, Red: 0, Green: this.StaticButtons[this.SelectedSequence].Color.G, Blue: 0}, eventsForLaunchpad, guiButtons)
			updateStaticLamp(this.SelectedSequence, this.StaticButtons[this.SelectedSequence], commandChannels)

			// Update the status bar
			common.UpdateStatusBar(fmt.Sprintf("Green %02d", this.StaticButtons[this.SelectedSequence].Color.G), "green", false, guiButtons)
			return
		}
	}

	// Blue
	if X == 3 && Y == -1 && sequences[this.SelectedSequence].Type != "scanner" {

		if sequences[this.SelectedSequence].Functions[common.Function6_Static_Gobo].State {
			if debug {
				fmt.Printf("Choose Static Blue X:%d Y:%d\n", X, Y)
			}

			buttonTouched(common.ALight{X: X, Y: Y, OnColor: common.White, OffColor: common.Blue}, eventsForLaunchpad, guiButtons)

			this.StaticButtons[this.SelectedSequence].X = this.LastStaticColorButtonX
			this.StaticButtons[this.SelectedSequence].Y = this.LastStaticColorButtonY
			this.StaticButtons[this.SelectedSequence].Color.B = this.StaticButtons[this.SelectedSequence].Color.B - 10
			if this.StaticButtons[this.SelectedSequence].Color.B > 255 {
				this.StaticButtons[this.SelectedSequence].Color.B = 0
			}
			if this.StaticButtons[this.SelectedSequence].Color.B < 0 {
				this.StaticButtons[this.SelectedSequence].Color.B = 255
			}
			common.LightLamp(common.ALight{X: X, Y: Y, Brightness: this.MasterBrightness, Red: 0, Green: 0, Blue: this.StaticButtons[this.SelectedSequence].Color.B}, eventsForLaunchpad, guiButtons)
			updateStaticLamp(this.SelectedSequence, this.StaticButtons[this.SelectedSequence], commandChannels)

			// Update the status bar
			common.UpdateStatusBar(fmt.Sprintf("Blue %02d", this.StaticButtons[this.SelectedSequence].Color.B), "blue", false, guiButtons)
			return
		}
	}

	// S E L E C T   R G B   S E Q U E N C E   C O L O R
	if X >= 0 && X < 8 && Y != -1 &&
		this.SelectedSequence == Y && // Make sure the buttons pressed are for this sequence.
		sequences[this.SelectedSequence].Type != "scanner" &&
		!this.EditFixtureSelectionMode &&
		sequences[this.SelectedSequence].Functions[common.Function5_Color].State {

		if debug {
			fmt.Printf("Set Sequence Color X:%d Y:%d\n", X, Y)
		}

		// If we're a scanner we can only select one color at a time.
		if sequences[this.SelectedSequence].Type == "scanner" {
			// Clear the sequence colors for this sequence.
			cmd := common.Command{
				Action: common.ClearSequenceColor,
			}
			common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)
		}

		// Add the selected color to the sequence.
		cmd := common.Command{
			Action: common.UpdateSequenceColor,
			Args: []common.Arg{
				{Name: "SelectedColor", Value: X},
			},
		}
		common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

		this.EditSequenceColorsMode[this.SelectedSequence] = true

		// Get an upto date copy of the sequence.
		sequences[this.SelectedSequence] = common.RefreshSequence(this.SelectedSequence, commandChannels, updateChannels)

		// Set the colors.
		sequences[this.SelectedSequence].CurrentColors = sequences[this.SelectedSequence].SequenceColors

		// We call ShowRGBColorSelectionButtons here so the selections will flash as you press them.
		ShowRGBColorSelectionButtons(this.MasterBrightness, this.SelectedSequence, *sequences[this.SelectedSequence], eventsForLaunchpad, guiButtons)

		return
	}

	// S E L E C T    S C A N N E R   C O L O R
	if X >= 0 && X < 8 && Y != -1 &&
		this.SelectedSequence == Y && // Make sure the buttons pressed are for this sequence.
		!this.EditFixtureSelectionMode &&
		sequences[this.SelectedSequence].Type == "scanner" &&
		sequences[this.SelectedSequence].Functions[common.Function5_Color].State {

		if debug {
			fmt.Printf("Set Scanner Color X:%d Y:%d\n", X, Y)
		}

		this.ScannerColor = X

		// Set the scanner color for this sequence.
		cmd := common.Command{
			Action: common.UpdateScannerColor,
			Args: []common.Arg{
				{Name: "SelectedColor", Value: this.ScannerColor},
				{Name: "SelectedFixture", Value: this.SelectedFixture},
			},
		}
		common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

		this.EditScannerColorsMode[this.SelectedSequence] = true

		// Get an upto date copy of the sequence.
		sequences[this.SelectedSequence] = common.RefreshSequence(this.SelectedSequence, commandChannels, updateChannels)

		// Set the colors.
		sequences[this.SelectedSequence].CurrentColors = sequences[this.SelectedSequence].SequenceColors

		// If the sequence isn't running this will force a single color DMX message.
		fixture.MapFixturesColorOnly(sequences[this.SelectedSequence], dmxController, fixturesConfig, this.ScannerColor, this.DmxInterfacePresent)

		// Clear the pattern function keys
		common.ClearSelectedRowOfButtons(this.SelectedSequence, eventsForLaunchpad, guiButtons)

		// We call ShowScannerColorSelectionButtons here so the selections will flash as you press them.
		ShowScannerColorSelectionButtons(*sequences[this.SelectedSequence], this, eventsForLaunchpad, fixturesConfig, guiButtons)

		return
	}

	// S E L E C T   F I X T U R E
	if X >= 0 && X < 8 && Y != -1 &&
		this.EditFixtureSelectionMode &&
		sequences[this.SelectedSequence].Functions[common.Function6_Static_Gobo].State || sequences[this.SelectedSequence].Functions[common.Function5_Color].State &&
		sequences[this.SelectedSequence].Type == "scanner" {

		this.SelectedFixture = X

		if debug {
			fmt.Printf("Selected Fixture is %d \n", this.SelectedFixture)
		}

		// Clear the pattern function keys
		common.ClearSelectedRowOfButtons(this.SelectedSequence, eventsForLaunchpad, guiButtons)

		// Update the buttons.
		if this.FollowingAction == "ShowGoboSelectionButtons" {
			ShowGoboSelectionButtons(*sequences[this.SelectedSequence], this, eventsForLaunchpad, guiButtons)
		}
		if this.FollowingAction == "ShowScannerColorSelectionButtons" {
			err := ShowScannerColorSelectionButtons(*sequences[this.SelectedSequence], this, eventsForLaunchpad, fixturesConfig, guiButtons)
			if err != nil {
				fmt.Printf("error: %v\n", err)
				common.RevealSequence(this.SelectedSequence, commandChannels)
			}
		}
		this.EditFixtureSelectionMode = false
		return
	}

	// S E L E C T   S C A N N E R   G O B O
	if X >= 0 && X < 8 && Y != -1 &&
		this.SelectedSequence == Y && // Make sure the buttons pressed are for this sequence.
		!this.EditFixtureSelectionMode &&
		sequences[this.SelectedSequence].Type == "scanner" &&
		sequences[this.SelectedSequence].Functions[common.Function6_Static_Gobo].State {

		if debug {
			fmt.Printf("Sequence %d Fixture %d Set Gobo %d\n", this.SelectedSequence, this.SelectedFixture, this.SelectedGobo)
		}

		this.SelectedGobo = X + 1

		// Set the selected gobo for this sequence.
		cmd := common.Command{
			Action: common.UpdateGobo,
			Args: []common.Arg{
				{Name: "SelectedGobo", Value: this.SelectedGobo},
				{Name: "FixtureNumber", Value: this.SelectedFixture},
			},
		}
		common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

		this.EditGoboSelectionMode[this.SelectedSequence] = true

		// Get an upto date copy of the sequence.
		sequences[this.SelectedSequence] = common.RefreshSequence(this.SelectedSequence, commandChannels, updateChannels)

		// If the sequence isn't running this will force a single gobo DMX message.
		fixture.MapFixturesGoboOnly(sequences[this.SelectedSequence], dmxController, fixturesConfig, this.SelectedGobo, this.DmxInterfacePresent)

		// Clear the pattern function keys
		common.ClearSelectedRowOfButtons(this.SelectedSequence, eventsForLaunchpad, guiButtons)

		// We call ShowGoboSelectionButtons here so the selections will flash as you press them.
		ShowGoboSelectionButtons(*sequences[this.SelectedSequence], this, eventsForLaunchpad, guiButtons)

		return
	}

	// S E L E C T   S T A T I C   C O L O R
	if X >= 0 && X < 8 &&
		Y != -1 &&
		!this.EditFixtureSelectionMode &&
		this.SelectedSequence == Y && // Make sure the buttons pressed are for this sequence.
		sequences[this.SelectedSequence].Type != "scanner" && // Not a scanner sequence.
		!this.FunctionSelectMode[this.SelectedSequence] && // Not in function Mode
		this.EditStaticColorsMode[this.SelectedSequence] { // Static Function On

		// For this button increment the color.
		sequences[this.SelectedSequence].StaticColors[X].X = X
		sequences[this.SelectedSequence].StaticColors[X].Y = Y
		if sequences[this.SelectedSequence].StaticColors[X].FirstPress {
			sequences[this.SelectedSequence].StaticColors[X].SelectedColor++
		}
		if sequences[this.SelectedSequence].StaticColors[X].SelectedColor > 10 {
			sequences[this.SelectedSequence].StaticColors[X].SelectedColor = 0
		}
		sequences[this.SelectedSequence].StaticColors[X].Color = common.GetColorButtonsArray(sequences[this.SelectedSequence].StaticColors[X].SelectedColor)
		if debug {
			fmt.Printf("Selected X:%d Y:%d Static Color is %d\n", X, Y, sequences[this.SelectedSequence].StaticColors[X].SelectedColor)
		}

		// Store the color data to allow for editing of static colors.
		this.StaticButtons[this.SelectedSequence].X = X
		this.StaticButtons[this.SelectedSequence].Y = Y
		this.StaticButtons[this.SelectedSequence].Color = sequences[this.SelectedSequence].StaticColors[X].Color
		this.StaticButtons[this.SelectedSequence].SelectedColor = sequences[this.SelectedSequence].StaticColors[X].SelectedColor

		// Update the status bar
		common.UpdateStatusBar(fmt.Sprintf("Red %02d", sequences[this.SelectedSequence].StaticColors[X].Color.R), "red", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Green %02d", sequences[this.SelectedSequence].StaticColors[X].Color.G), "green", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Blue %02d", sequences[this.SelectedSequence].StaticColors[X].Color.B), "blue", false, guiButtons)

		// Tell the sequence about the new color and where we are in the
		// color cycle.
		cmd := common.Command{
			Action: common.UpdateStaticColor,
			Args: []common.Arg{
				{Name: "Static", Value: true},
				{Name: "StaticLamp", Value: X},
				{Name: "StaticLampFlash", Value: true},
				{Name: "SelectedColor", Value: sequences[this.SelectedSequence].StaticColors[X].SelectedColor},
				{Name: "StaticColor", Value: sequences[this.SelectedSequence].StaticColors[X].Color},
			},
		}
		common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)
		this.LastStaticColorButtonX = X
		this.LastStaticColorButtonY = Y

		// Set the first pressed for only this fixture and cancel any others
		for x := 0; x < 8; x++ {
			sequences[this.SelectedSequence].StaticColors[x].FirstPress = false
		}
		sequences[this.SelectedSequence].StaticColors[X].FirstPress = true

		return
	}

	// S E L E C T   P A T T E N
	if X >= 0 && X < 8 && Y != -1 &&
		!this.EditFixtureSelectionMode &&
		this.EditPatternMode[this.SelectedSequence] {

		if debug {
			fmt.Printf("Set Pattern to %d\n", X)
		}

		// Tell the sequence to change the pattern.
		cmd := common.Command{
			Action: common.UpdatePattern,
			Args: []common.Arg{
				{Name: "SelectPattern", Value: X},
			},
		}
		common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

		this.FunctionSelectMode[this.SelectedSequence] = false

		// Get an upto date copy of the sequence.
		sequences[this.SelectedSequence] = common.RefreshSequence(this.SelectedSequence, commandChannels, updateChannels)

		// We call ShowPatternSelectionButtons here so the selections will flash as you press them.
		this.EditFixtureSelectionMode = false
		ShowPatternSelectionButtons(this.SelectedSequence, *sequences[this.SelectedSequence], eventsForLaunchpad, guiButtons)

		return
	}

	// F U N C T I O N  K E Y S
	if X >= 0 && X < 8 &&
		this.FunctionSelectMode[this.SelectedSequence] &&
		!this.EditPatternMode[this.SelectedSequence] &&
		!this.EditStaticColorsMode[this.SelectedSequence] &&
		!this.EditGoboSelectionMode[this.SelectedSequence] &&
		!sequences[this.SelectedSequence].Functions[common.Function5_Color].State {

		if debug {
			fmt.Printf("Function Key X:%d Y:%d\n", X, Y)
		}

		// Get an upto date copy of the sequence.
		sequences[this.SelectedSequence] = common.RefreshSequence(this.SelectedSequence, commandChannels, updateChannels)

		for _, functions := range sequences[this.SelectedSequence].Functions {
			if Y == functions.SequenceNumber {
				if !sequences[this.SelectedSequence].Functions[X].State {
					sequences[this.SelectedSequence].Functions[X].State = true
					break
				}
				if sequences[this.SelectedSequence].Functions[X].State {
					sequences[this.SelectedSequence].Functions[X].State = false
					break
				}
			}
		}

		// Send update functions command. This sets the temporary representation of
		// the function keys in the real sequence.
		cmd := common.Command{
			Action: common.UpdateFunctions,
			Args: []common.Arg{
				{Name: "Functions", Value: sequences[this.SelectedSequence].Functions},
			},
		}
		common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

		// Light the correct function key.
		common.ShowFunctionButtons(*sequences[this.SelectedSequence], this.SelectedSequence, eventsForLaunchpad, guiButtons)

		// Now some functions mean that we go into another menu ( set of buttons )
		// This is true for :-
		// Function 1 - setting the pattern.
		// Function 5 - setting the sequence colors or selecting scanner color.
		// Function 6 - setting the static colors or selecting scanner gobo.

		// Map Function 1 to pattern mode.
		this.EditPatternMode[this.SelectedSequence] = sequences[this.SelectedSequence].Functions[common.Function1_Pattern].State

		// Go straight into pattern select mode, don't wait for a another select press.
		if this.EditPatternMode[this.SelectedSequence] {
			common.ClearSelectedRowOfButtons(this.SelectedSequence, eventsForLaunchpad, guiButtons)
			this.EditFixtureSelectionMode = false
			ShowPatternSelectionButtons(this.SelectedSequence, *sequences[this.SelectedSequence], eventsForLaunchpad, guiButtons)
		}

		// Function 5.

		// Map Function 5 to color edit.
		this.EditSequenceColorsMode[this.SelectedSequence] = sequences[this.SelectedSequence].Functions[common.Function5_Color].State

		// Go straight into RGB color edit mode, don't wait for a another select press.
		if this.EditSequenceColorsMode[this.SelectedSequence] && sequences[this.SelectedSequence].Type == "rgb" {
			time.Sleep(500 * time.Millisecond) // But give the launchpad time to light the function key purple.
			common.ClearSelectedRowOfButtons(this.SelectedSequence, eventsForLaunchpad, guiButtons)
			// Set the colors.
			sequences[this.SelectedSequence].CurrentColors = sequences[this.SelectedSequence].SequenceColors
			// Show the colors
			ShowRGBColorSelectionButtons(this.MasterBrightness, this.SelectedSequence, *sequences[this.SelectedSequence], eventsForLaunchpad, guiButtons)
		}

		// Go straight into scanner color edit mode via select fixture, don't wait for a another select press.
		if this.EditSequenceColorsMode[this.SelectedSequence] && sequences[this.SelectedSequence].Type == "scanner" {
			time.Sleep(500 * time.Millisecond) // But give the launchpad time to light the function key purple.
			common.ClearSelectedRowOfButtons(this.SelectedSequence, eventsForLaunchpad, guiButtons)
			this.EditFixtureSelectionMode = true
			sequences[this.SelectedSequence].StaticColors[X].FirstPress = false
			this.FollowingAction = "ShowScannerColorSelectionButtons"
			this.SelectedFixture = ShowSelectFixtureButtons(*sequences[this.SelectedSequence], this, eventsForLaunchpad, this.FollowingAction, guiButtons)
		}

		// Function 6

		// Map Function 6 to select gobo mode if we are in scanner sequence.
		if sequences[this.SelectedSequence].Type == "scanner" && sequences[this.SelectedSequence].Functions[common.Function6_Static_Gobo].State {
			this.EditStaticColorsMode[this.SelectedSequence] = false // Turn off the other option for this function key.
			this.EditGoboSelectionMode[this.SelectedSequence] = sequences[this.SelectedSequence].Functions[common.Function6_Static_Gobo].State

			// Go straight to gobo selection mode via select fixture, don't wait for a another select press.
			time.Sleep(500 * time.Millisecond) // But give the launchpad time to light the function key purple.
			common.ClearSelectedRowOfButtons(this.SelectedSequence, eventsForLaunchpad, guiButtons)
			this.EditFixtureSelectionMode = true
			sequences[this.SelectedSequence].StaticColors[X].FirstPress = false
			this.FollowingAction = "ShowGoboSelectionButtons"
			this.SelectedFixture = ShowSelectFixtureButtons(*sequences[this.SelectedSequence], this, eventsForLaunchpad, this.FollowingAction, guiButtons)

		}

		// Map Function 6 to static color edit if we are a RGB sequence.
		if sequences[this.SelectedSequence].Type == "rgb" && sequences[this.SelectedSequence].Functions[common.Function6_Static_Gobo].State {
			this.EditGoboSelectionMode[this.SelectedSequence] = false // Turn off the other option for this function key.

			// Turn on edit static color mode.
			if sequences[this.SelectedSequence].Functions[common.Function6_Static_Gobo].State {
				this.EditStaticColorsMode[this.SelectedSequence] = true
			}
			// Go straight to static color selection mode, don't wait for a another select press.
			time.Sleep(250 * time.Millisecond) // But give the launchpad time to light the function key purple.
			common.ClearLabelsSelectedRowOfButtons(this.SelectedSequence, guiButtons)
			this.FunctionSelectMode[this.SelectedSequence] = false

			// Swicth on any static colors.
			cmd = common.Command{
				Action: common.UpdateStatic,
				Args: []common.Arg{
					{Name: "Static", Value: true},
				},
			}
			common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)
		}

		return
	}

	// B L A C K O U T   B U T T O N.
	if X == 8 && Y == 7 {

		if debug {
			fmt.Printf("BLACKOUT\n")
		}

		if !this.Blackout {
			this.Blackout = true
			cmd := common.Command{
				Action: common.Blackout,
			}
			common.SendCommandToAllSequence(cmd, commandChannels)
			common.LightLamp(common.ALight{X: X, Y: Y, Brightness: this.MasterBrightness, Red: 0, Green: 0, Blue: 0}, eventsForLaunchpad, guiButtons)
			common.FlashLight(8, 7, common.Pink, common.White, eventsForLaunchpad, guiButtons)
		} else {
			this.Blackout = false
			cmd := common.Command{
				Action: common.Normal,
			}
			common.SendCommandToAllSequence(cmd, commandChannels)
			common.LightLamp(common.ALight{X: X, Y: Y, Brightness: this.MasterBrightness, Red: 255, Green: 255, Blue: 255}, eventsForLaunchpad, guiButtons)
		}
		return
	}
}

func updateStaticLamp(selectedSequence int, staticColorButtons common.StaticColorButton, commandChannels []chan common.Command) {

	// Static is set to true in the functions and this key is set to
	// the selected color.
	cmd := common.Command{
		Action: common.UpdateStaticColor,
		Args: []common.Arg{
			{Name: "Static", Value: true},
			{Name: "StaticLamp", Value: staticColorButtons.X},
			{Name: "StaticLampFlash", Value: false},
			{Name: "SelectedColor", Value: staticColorButtons.SelectedColor},
			{Name: "StaticColor", Value: common.Color{R: staticColorButtons.Color.R, G: staticColorButtons.Color.G, B: staticColorButtons.Color.B}},
		},
	}
	common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

}

// HandleSelect - Runs when you press a select button to select a sequence.
func HandleSelect(sequences []*common.Sequence, this *CurrentState, eventsForLaunchpad chan common.ALight,
	commandChannels []chan common.Command, guiButtons chan common.ALight) {

	if debug {
		fmt.Printf("HANDLE: this.SelectButtonPressed[%d] = %t \n", this.SelectedSequence, this.SelectButtonPressed[this.SelectedSequence])
		fmt.Printf("HANDLE: this.FunctionSelectMode[%d] = %t \n", this.SelectedSequence, this.FunctionSelectMode[this.SelectedSequence])
		fmt.Printf("HANDLE: this.EditSequenceColorsMode[%d] = %t \n", this.SelectedSequence, this.EditSequenceColorsMode[this.SelectedSequence])
		fmt.Printf("HANDLE: this.EditStaticColorsMode[%d] = %t \n", this.SelectedSequence, this.EditStaticColorsMode[this.SelectedSequence])
		fmt.Printf("HANDLE: this.EditGoboSelectionMode[%d] = %t \n", this.SelectedSequence, this.EditGoboSelectionMode[this.SelectedSequence])
		fmt.Printf("HANDLE: this.EditPatternMode[%d] = %t \n", this.SelectedSequence, this.EditPatternMode[this.SelectedSequence])
		fmt.Printf("HANDLE: Function6_Static_Gobo[%d] = %t\n", this.SelectedSequence, sequences[this.SelectedSequence].Functions[common.Function6_Static_Gobo].State)
	}

	// Update the status bar
	if this.Strobe[this.SelectedSequence] {
		common.UpdateStatusBar(fmt.Sprintf("Strobe %02d", this.StrobeSpeed[this.SelectedSequence]), "speed", false, guiButtons)
	} else {
		// Update status bar.
		if sequences[this.SelectedSequence].Functions[common.Function8_Music_Trigger].State {
			common.UpdateStatusBar("  MUSIC  ", "speed", false, guiButtons)
		} else {
			common.UpdateStatusBar(fmt.Sprintf("Speed %02d", this.Speed[this.SelectedSequence]), "speed", false, guiButtons)
		}
	}

	sensitivity := common.FindSensitivity(this.SoundGain)
	common.UpdateStatusBar(fmt.Sprintf("Sensitivity %02d", sensitivity), "sensitivity", false, guiButtons)
	common.UpdateStatusBar(fmt.Sprintf("Master %02d", this.MasterBrightness), "master", false, guiButtons)

	if sequences[this.SelectedSequence].Type == "rgb" {
		common.UpdateStatusBar(fmt.Sprintf("Shift %02d", this.RGBShift[this.SelectedSequence]), "shift", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Size %02d", this.RGBSize[this.SelectedSequence]), "size", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Fade %02d", this.RGBFade[this.SelectedSequence]), "fade", false, guiButtons)
		common.UpdateStatusBar("       ", "tilt", false, guiButtons)

		common.UpdateStatusBar(fmt.Sprintf("Red %02d", this.StaticButtons[this.SelectedSequence].Color.R), "red", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Green %02d", this.StaticButtons[this.SelectedSequence].Color.G), "green", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Blue %02d", this.StaticButtons[this.SelectedSequence].Color.B), "blue", false, guiButtons)
	}
	if sequences[this.SelectedSequence].Type == "scanner" {
		label := getScannerShiftLabel(this.ScannerShift[this.SelectedSequence])
		common.UpdateStatusBar(fmt.Sprintf("Shift %s", label), "shift", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Size %02d", this.ScannerSize[this.SelectedSequence]), "size", false, guiButtons)
		label = getScannerCoordinatesLabel(this.ScannerCoordinates[this.SelectedSequence])
		common.UpdateStatusBar(fmt.Sprintf("Coord %s", label), "fade", false, guiButtons)

		// Hide the color editing buttons.
		common.UpdateStatusBar(fmt.Sprintf("Tilt %02d", this.OffsetTilt), "tilt", false, guiButtons)
		common.UpdateStatusBar("        ", "red", false, guiButtons)
		common.UpdateStatusBar("        ", "green", false, guiButtons)
		common.UpdateStatusBar(fmt.Sprintf("Pan %02d", this.OffsetPan), "pan", false, guiButtons)
	}

	// Light the top buttons.
	common.ShowTopButtons(sequences[this.SelectedSequence].Type, eventsForLaunchpad, guiButtons)

	// Light the sequence selector button.
	sequence.SequenceSelect(eventsForLaunchpad, guiButtons, this.SelectedSequence)

	// Light the strobe button.
	common.ShowStrobeButtonStatus(this.Strobe[this.SelectedSequence], eventsForLaunchpad, guiButtons)

	//Light the start stop button.
	common.ShowRunningStatus(this.SelectedSequence, this.Running, eventsForLaunchpad, guiButtons)

	// First time into function mode we head back to normal mode.
	if this.FunctionSelectMode[this.SelectedSequence] &&
		!this.SelectButtonPressed[this.SelectedSequence] &&
		!this.EditSequenceColorsMode[this.SelectedSequence] &&
		!this.EditStaticColorsMode[this.SelectedSequence] {
		if debug {
			fmt.Printf("Handle 1 Function Bar off\n")
		}
		// Turn off function mode. Remove the function pads.
		common.ClearSelectedRowOfButtons(this.SelectedSequence, eventsForLaunchpad, guiButtons)

		this.FunctionSelectMode[this.SelectedSequence] = false

		if sequences[this.SelectedSequence].Functions[common.Function1_Pattern].State {
			if debug {
				fmt.Printf("Show Pattern Selection Buttons\n")
			}
			this.EditPatternMode[this.SelectedSequence] = true
			common.HideSequence(this.SelectedSequence, commandChannels)
			ShowPatternSelectionButtons(this.SelectedSequence, *sequences[this.SelectedSequence], eventsForLaunchpad, guiButtons)
			return
		}

		if sequences[this.SelectedSequence].Functions[common.Function5_Color].State && sequences[this.SelectedSequence].Type == "rgb" {
			if debug {
				fmt.Printf("Show RGB Sequence Color Selection Buttons\n")
			}
			// Set the colors.
			sequences[this.SelectedSequence].CurrentColors = sequences[this.SelectedSequence].SequenceColors
			// Show the colors
			ShowRGBColorSelectionButtons(this.MasterBrightness, this.SelectedSequence, *sequences[this.SelectedSequence], eventsForLaunchpad, guiButtons)
			return
		}

		if sequences[this.SelectedSequence].Functions[common.Function6_Static_Gobo].State &&
			sequences[this.SelectedSequence].Type != "scanner" {
			if debug {
				fmt.Printf("Show Static Color Selection Buttons\n")
			}
			common.SetMode(this.SelectedSequence, commandChannels, "Static")
			sequences[this.SelectedSequence].Functions[common.Function6_Static_Gobo].State = false
			return
		}

		// Allow us to exit the pattern select mode without setting a pattern.
		if this.EditPatternMode[this.SelectedSequence] {
			this.EditPatternMode[this.SelectedSequence] = false
		}

		// Switch off the gobo selection mode.
		sequences[this.SelectedSequence].Functions[common.Function6_Static_Gobo].State = false

		// Else reveal the sequence on the launchpad keys
		if debug {
			fmt.Printf("Reveal Sequence\n")
		}
		common.RevealSequence(this.SelectedSequence, commandChannels)
		// Turn off the function mode flag.
		this.FunctionSelectMode[this.SelectedSequence] = false
		// Now forget we pressed twice and start again.
		this.SelectButtonPressed[this.SelectedSequence] = true

		// Reveal the scanner State
		if sequences[this.SelectedSequence].Type == "scanner" {
			// Show the status.
			ShowScannerStatus(this.SelectedSequence, *sequences[this.SelectedSequence], this, eventsForLaunchpad, guiButtons, commandChannels)
		}

		return
	}

	// This the first time we have pressed the select button.
	if !this.SelectButtonPressed[this.SelectedSequence] && !this.EditStaticColorsMode[this.SelectedSequence] {
		if debug {
			fmt.Printf("Handle 2\n")
		}
		// assume everything else is off.
		this.SelectButtonPressed[0] = false
		this.SelectButtonPressed[1] = false
		this.SelectButtonPressed[2] = false
		this.SelectButtonPressed[3] = false
		// But remember we have pressed this select button once.
		this.FunctionSelectMode[this.SelectedSequence] = false
		this.SelectButtonPressed[this.SelectedSequence] = true

		if sequences[this.SelectedSequence].Functions[common.Function1_Pattern].State {
			// Reset the pattern function key.
			sequences[this.SelectedSequence].Functions[common.Function1_Pattern].State = false

			// Clear the pattern function keys
			ClearPatternSelectionButtons(this.SelectedSequence, *sequences[this.SelectedSequence], eventsForLaunchpad, guiButtons)

			// And reveal the sequence.
			common.RevealSequence(this.SelectedSequence, commandChannels)

			// Editing pattern is over for this sequence.
			this.EditPatternMode[this.SelectedSequence] = false

			// Clear buttons and remove any labels.
			common.ClearSelectedRowOfButtons(this.SelectedSequence, eventsForLaunchpad, guiButtons)
		}

		if !this.FunctionSelectMode[this.SelectedSequence] && sequences[this.SelectedSequence].Functions[common.Function5_Color].State && this.EditSequenceColorsMode[this.SelectedSequence] {
			unSetEditSequenceColorsMode(sequences, this, commandChannels, eventsForLaunchpad, guiButtons)
		}

		// Tailor the top buttons to the sequence type.
		common.ShowTopButtons(sequences[this.SelectedSequence].Type, eventsForLaunchpad, guiButtons)

		// Tailor the bottom buttons to the sequence type.
		common.ShowBottomButtons(sequences[this.SelectedSequence].Type, eventsForLaunchpad, guiButtons)

		// Show this sequence running status in the start/stop button.
		common.ShowRunningStatus(this.SelectedSequence, this.Running, eventsForLaunchpad, guiButtons)

		return
	}

	// Are we in function mode ?
	if this.FunctionSelectMode[this.SelectedSequence] {
		if debug {
			fmt.Printf("Handle 3\n")
		}
		// Turn off function mode. Remove the function pads.
		common.ClearSelectedRowOfButtons(this.SelectedSequence, eventsForLaunchpad, guiButtons)
		// And reveal the sequence on the launchpad keys
		common.RevealSequence(this.SelectedSequence, commandChannels)
		// Turn off the function mode flag.
		this.FunctionSelectMode[this.SelectedSequence] = false
		// Now forget we pressed twice and start again.
		this.SelectButtonPressed[this.SelectedSequence] = false

		return
	}

	// We are in function mode for this sequence.
	if !this.FunctionSelectMode[this.SelectedSequence] &&
		sequences[this.SelectedSequence].Type != "switch" || // Don't alow functions in switch mode.
		!this.FunctionSelectMode[this.SelectedSequence] && // Function select mode is off
			this.EditStaticColorsMode[this.SelectedSequence] { // The case when we leave static colors edit mode.

		if debug {
			fmt.Printf("Handle 4 - Function Bar On!\n")
		}

		// Unset the function key.
		sequences[this.SelectedSequence].Functions[common.Function6_Static_Gobo].State = false

		// Unset the edit static color mode.
		this.EditStaticColorsMode[this.SelectedSequence] = false

		// Set function mode.
		this.FunctionSelectMode[this.SelectedSequence] = true

		// And hide the sequence so we can only see the function buttons.
		common.HideSequence(this.SelectedSequence, commandChannels)

		// Turn off any static sequence so we can see the functions.
		common.SetMode(this.SelectedSequence, commandChannels, "Sequence")

		// Turn off any previous function bars.
		for sequenceNumber := range sequences {
			if this.FunctionSelectMode[sequenceNumber] {
				common.ClearSelectedRowOfButtons(sequenceNumber, eventsForLaunchpad, guiButtons)
				// And reveal all the other sequence that isn't us.
				if sequenceNumber != this.SelectedSequence {
					common.RevealSequence(sequenceNumber, commandChannels)
					// And turn off the function selected.
					this.FunctionSelectMode[sequenceNumber] = false
				}
			}
		}

		// Create the function buttons.
		common.MakeFunctionButtons(this.SelectedSequence, eventsForLaunchpad, guiButtons, this.SequenceChannels)

		// Now forget we pressed twice and start again.
		this.SelectButtonPressed[this.SelectedSequence] = false

		return
	}
}

func unSetEditSequenceColorsMode(sequences []*common.Sequence, this *CurrentState, commandChannels []chan common.Command,
	eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight) {

	// Turn off the edit colors bar.
	sequences[this.SelectedSequence].Functions[common.Function5_Color].State = false
	cmd := common.Command{
		Action: common.UpdateFunctions,
		Args: []common.Arg{
			{Name: "Functions", Value: sequences[this.SelectedSequence].Functions},
		},
	}
	common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

	// Restart the sequence.
	cmd = common.Command{
		Action: common.Start,
	}
	common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

	// And reveal the sequence on the launchpad keys
	common.RevealSequence(this.SelectedSequence, commandChannels)
	// Turn off the function mode flag.
	this.FunctionSelectMode[this.SelectedSequence] = false
	// Now forget we pressed twice and start again.
	this.SelectButtonPressed[this.SelectedSequence] = true

	common.HideColorSelectionButtons(this.SelectedSequence, *sequences[this.SelectedSequence], eventsForLaunchpad, guiButtons)
}

func AllFixturesOff(sequences []*common.Sequence, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, dmxController *ft232.DMXController, fixturesConfig *fixture.Fixtures, dmxInterfacePresent bool) {
	for y := 0; y < len(sequences); y++ {
		if sequences[y].Type != "switch" {
			for x := 0; x < 8; x++ {
				common.LightLamp(common.ALight{X: x, Y: y, Brightness: 0, Red: 0, Green: 0, Blue: 0}, eventsForLaunchpad, guiButtons)
				fixture.MapFixtures(y, dmxController, x, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, true, 0, 0, false, 0, dmxInterfacePresent)
				common.LabelButton(x, y, "", guiButtons)
			}
		}
	}
}
func AllRGBFixturesOff(sequences []*common.Sequence, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, dmxController *ft232.DMXController, fixturesConfig *fixture.Fixtures, dmxInterfacePresent bool) {
	for x := 0; x < 8; x++ {
		for sequenceNumber := 0; sequenceNumber < len(sequences); sequenceNumber++ {
			if sequences[sequenceNumber].Type == "rgb" {
				common.LightLamp(common.ALight{X: x, Y: sequenceNumber, Brightness: 0, Red: 0, Green: 0, Blue: 0}, eventsForLaunchpad, guiButtons)
				fixture.MapFixtures(sequenceNumber, dmxController, x, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, true, 0, 0, false, 0, dmxInterfacePresent)
				common.LabelButton(x, sequenceNumber, "", guiButtons)
			}
		}
	}
}

// Show Scanner status - Dim White is disabled, White is enabled.
func ShowScannerStatus(selectedSequence int, sequence common.Sequence, this *CurrentState, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command) {

	if debug {
		fmt.Printf("Show Scanner Status for sequence %d number of scanners %d\n", sequence.Number, sequence.NumberFixtures)
	}

	common.HideSequence(selectedSequence, commandChannels)

	for scannerNumber := 0; scannerNumber < sequence.NumberFixtures; scannerNumber++ {

		if debug {
			fmt.Printf("Enabled %t Inverted %t\n", this.ScannerState[scannerNumber][sequence.Number].Enabled, this.ScannerState[scannerNumber][sequence.Number].Inverted)
		}

		// Enabled but not inverted then On and green.
		if this.ScannerState[scannerNumber][sequence.Number].Enabled && !this.ScannerState[scannerNumber][sequence.Number].Inverted {
			common.LightLamp(common.ALight{X: scannerNumber, Y: sequence.Number, Brightness: this.MasterBrightness, Red: 0, Green: 255, Blue: 0}, eventsForLaunchpad, guiButtons)
			common.LabelButton(scannerNumber, sequence.Number, "On", guiButtons)
		}

		// Enabled and inverted then Invert and red.
		if this.ScannerState[scannerNumber][sequence.Number].Enabled && this.ScannerState[scannerNumber][sequence.Number].Inverted {
			common.LightLamp(common.ALight{X: scannerNumber, Y: sequence.Number, Brightness: this.MasterBrightness, Red: 255, Green: 0, Blue: 0}, eventsForLaunchpad, guiButtons)
			common.LabelButton(scannerNumber, sequence.Number, "Invert", guiButtons)
		}

		// Not enabled and not inverted then off and blue.
		if !this.ScannerState[scannerNumber][sequence.Number].Enabled && !this.ScannerState[scannerNumber][sequence.Number].Inverted {
			common.LightLamp(common.ALight{X: scannerNumber, Y: sequence.Number, Brightness: this.MasterBrightness, Red: 0, Green: 100, Blue: 150}, eventsForLaunchpad, guiButtons)
			common.LabelButton(scannerNumber, sequence.Number, "Off", guiButtons)
		}

	}
	time.Sleep(200 * time.Millisecond) // Wait so we can see the changes.
	common.RevealSequence(selectedSequence, commandChannels)
}

// For the given sequence show the available sequence colors on the relevant buttons.
func ShowRGBColorSelectionButtons(master int, mySequenceNumber int, sequence common.Sequence, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight) {

	if debug {
		fmt.Printf("Show Color Selection Buttons\n")
	}
	// Check if we need to flash this button.
	for myFixtureNumber, lamp := range sequence.RGBAvailableColors {

		for index, availableColor := range sequence.RGBAvailableColors {
			for _, sequenceColor := range sequence.CurrentColors {
				if debug {
					fmt.Printf("myFixtureNumber %d   current color %+v\n", myFixtureNumber, sequenceColor)
				}
				if availableColor.Color == sequenceColor {
					if myFixtureNumber == index {
						lamp.Flash = true
					}
				}
			}
		}
		if lamp.Flash {
			Black := common.Color{R: 0, G: 0, B: 0}
			common.FlashLight(myFixtureNumber, mySequenceNumber, lamp.Color, Black, eventsForLaunchpad, guiButtons)
		} else {
			common.LightLamp(common.ALight{X: myFixtureNumber, Y: mySequenceNumber, Brightness: master, Red: lamp.Color.R, Green: lamp.Color.G, Blue: lamp.Color.B}, eventsForLaunchpad, guiButtons)
		}
	}
}

// For the given sequence show the available scanner selection colors on the relevant buttons.
func ShowSelectFixtureButtons(sequence common.Sequence, this *CurrentState, eventsForLaunchpad chan common.ALight, action string, guiButtons chan common.ALight) int {

	if debug {
		fmt.Printf("Sequence %d Show Fixture Selection Buttons on the way to %s\n", this.SelectedSequence, action)
	}

	for fixtureNumber, fixture := range sequence.ScannersAvailable {

		if debug {
			fmt.Printf("Fixture %+v\n", fixture)
		}
		if fixtureNumber == this.SelectedFixture {
			fixture.Flash = true
			this.SelectedFixture = fixtureNumber
		}
		if fixture.Flash {
			White := common.Color{R: 255, G: 255, B: 255}
			common.FlashLight(fixtureNumber, this.SelectedSequence, fixture.Color, White, eventsForLaunchpad, guiButtons)

		} else {
			common.LightLamp(common.ALight{X: fixtureNumber, Y: this.SelectedSequence, Red: fixture.Color.R, Green: fixture.Color.G, Blue: fixture.Color.B, Brightness: sequence.Master}, eventsForLaunchpad, guiButtons)
		}
		common.LabelButton(fixtureNumber, this.SelectedSequence, fixture.Label, guiButtons)
	}
	if debug {
		fmt.Printf("Selected Fixture is %d\n", this.SelectedFixture)
	}
	return this.SelectedFixture
}

// ShowGoboSelectionButtons puts up a set of red buttons used to select a fixture.
func ShowGoboSelectionButtons(sequence common.Sequence, this *CurrentState, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight) {

	if debug {
		fmt.Printf("Sequence %d Show Gobo Selection Buttons\n", this.SelectedSequence)
	}
	// Check if we need to flash this button.
	for goboNumber, gobo := range sequence.ScannerAvailableGobos[this.SelectedFixture+1] {

		if gobo.Number > 8 {
			return // We only have 8 buttons so we can't select from any more.
		}
		if gobo.Number == sequence.ScannerGobo[this.SelectedFixture] {
			gobo.Flash = true
		}
		if debug {
			fmt.Printf("goboNumber %d   current gobo %d  flash gobo %t\n", goboNumber, sequence.ScannerGobo, gobo.Flash)
		}
		if gobo.Flash {
			Black := common.Color{R: 0, G: 0, B: 0}
			common.FlashLight(goboNumber, this.SelectedSequence, gobo.Color, Black, eventsForLaunchpad, guiButtons)
		} else {
			common.LightLamp(common.ALight{X: goboNumber, Y: this.SelectedSequence, Brightness: sequence.Master, Red: gobo.Color.R, Green: gobo.Color.G, Blue: gobo.Color.B}, eventsForLaunchpad, guiButtons)
		}
		goboName := common.FormatLabel(gobo.Name)
		common.LabelButton(goboNumber, this.SelectedSequence, goboName, guiButtons)
	}
}

// For the given sequence show the available scanner selection colors on the relevant buttons.
func ShowScannerColorSelectionButtons(sequence common.Sequence, this *CurrentState, eventsForLaunchpad chan common.ALight, fixtures *fixture.Fixtures, guiButtons chan common.ALight) error {

	if debug {
		fmt.Printf("Show Scanner Color Selection Buttons,  Sequence is %d  fixture is %d   color is %d \n", this.SelectedSequence, this.SelectedFixture, sequence.ScannerColor[this.SelectedFixture])
	}

	// if there are no colors available for this fixture turn everything off and print an error.
	if sequence.ScannerAvailableColors[this.SelectedFixture+1] == nil {
		for _, fixture := range fixtures.Fixtures {
			if fixture.Group == this.SelectedSequence+1 {
				common.LightLamp(common.ALight{X: fixture.Number - 1, Y: this.SelectedSequence, Red: 0, Green: 0, Blue: 0, Brightness: sequence.Master}, eventsForLaunchpad, guiButtons)
			}
		}
		return fmt.Errorf("error: no colors available for fixture number %d", this.SelectedFixture+1)
	}

	// selected fixture is +1 here because the fixtures in the yaml config file start with 1 not 0.
	for fixtureNumber, lamp := range sequence.ScannerAvailableColors[this.SelectedFixture+1] {

		if debug {
			fmt.Printf("Lamp %+v\n", lamp)
		}
		if fixtureNumber == sequence.ScannerColor[this.SelectedFixture] {
			lamp.Flash = true
		}

		if lamp.Flash {
			Black := common.Color{R: 0, G: 0, B: 0}
			common.FlashLight(fixtureNumber, this.SelectedSequence, lamp.Color, Black, eventsForLaunchpad, guiButtons)
		} else {
			common.LightLamp(common.ALight{X: fixtureNumber, Y: this.SelectedSequence, Red: lamp.Color.R, Green: lamp.Color.G, Blue: lamp.Color.B, Brightness: sequence.Master}, eventsForLaunchpad, guiButtons)
		}
		// Remove any labels.
		common.LabelButton(fixtureNumber, this.SelectedSequence, "", guiButtons)
	}
	return nil
}

// For the given sequence clear the available this.Patterns on the relevant buttons.
func ClearPatternSelectionButtons(mySequenceNumber int, sequence common.Sequence, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight) {
	// Check if we need to flash this button.
	for myFixtureNumber := 0; myFixtureNumber < 4; myFixtureNumber++ {
		common.LightLamp(common.ALight{X: myFixtureNumber, Y: mySequenceNumber, Red: 0, Green: 0, Blue: 0, Brightness: sequence.Master}, eventsForLaunchpad, guiButtons)
	}
}

// For the given sequence show the available patterns on the relevant buttons.
func ShowPatternSelectionButtons(mySequenceNumber int, sequence common.Sequence, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight) {

	if debug {
		fmt.Printf("Sequence Name %s Type %s  Label %s\n", sequence.Name, sequence.Type, sequence.Label)
	}

	LightBlue := common.Color{R: 0, G: 100, B: 255}
	White := common.Color{R: 255, G: 255, B: 255}

	if sequence.Type == "rgb" {
		for _, pattern := range sequence.RGBAvailablePatterns {
			if pattern.Number == sequence.SelectedPattern {
				common.FlashLight(pattern.Number, mySequenceNumber, White, LightBlue, eventsForLaunchpad, guiButtons)
			} else {
				common.LightLamp(common.ALight{X: pattern.Number, Y: mySequenceNumber, Red: 0, Green: 100, Blue: 255, Brightness: sequence.Master}, eventsForLaunchpad, guiButtons)
			}
			common.LabelButton(pattern.Number, mySequenceNumber, pattern.Label, guiButtons)
		}
		return
	}

	if sequence.Type == "scanner" {
		for _, pattern := range sequence.ScannerAvailablePatterns {
			if pattern.Number == sequence.SelectedPattern {
				common.FlashLight(pattern.Number, mySequenceNumber, White, LightBlue, eventsForLaunchpad, guiButtons)
			} else {
				common.LightLamp(common.ALight{X: pattern.Number, Y: mySequenceNumber, Red: 0, Green: 100, Blue: 255, Brightness: sequence.Master}, eventsForLaunchpad, guiButtons)
			}
			common.LabelButton(pattern.Number, mySequenceNumber, pattern.Label, guiButtons)
		}
		return
	}
}

func InitButtons(this *CurrentState, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight) {

	// Light the logo blue.
	if this.LaunchPadConnected {
		this.Pad.Light(8, -1, 0, 0, 255)
	}

	// Light up any existing presets.
	presets.RefreshPresets(eventsForLaunchpad, guiButtons, this.PresetsStore)

	// Light the buttons at the bottom.
	common.ShowBottomButtons("rgb", eventsForLaunchpad, guiButtons)

	// Light the top buttons.
	common.ShowTopButtons("rgb", eventsForLaunchpad, guiButtons)

	// Light the first sequence as the default selected.
	this.SelectedSequence = 0
	sequence.SequenceSelect(eventsForLaunchpad, guiButtons, this.SelectedSequence)

}

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

	AllFixturesOff(sequences, eventsForLaunchpad, guiButtons, dmxController, fixturesConfig, dmxInterfacePresent)

	// Load the config.
	// Which forces all sequences to load their config.
	config.AskToLoadConfig(commandChannels, X, Y)

	// Turn the selected preset light flashing red / yellow.
	if this.LastPreset != nil {
		current := this.PresetsStore[*this.LastPreset]
		this.PresetsStore[*this.LastPreset] = presets.Preset{State: current.State, Selected: false, Label: current.Label}
	}
	current := this.PresetsStore[fmt.Sprint(X)+","+fmt.Sprint(Y)]
	this.PresetsStore[fmt.Sprint(X)+","+fmt.Sprint(Y)] = presets.Preset{State: current.State, Selected: true, Label: current.Label}
	presets.RefreshPresets(eventsForLaunchpad, guiButtons, this.PresetsStore)

	// Preserve this.Blackout.
	if !this.Blackout {
		cmd := common.Command{
			Action: common.Normal,
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
	for sequenceNumber := range sequences {
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

		// switch off any color editing.
		sequences[this.SelectedSequence].Functions[common.Function6_Static_Gobo].State = false

		// If we are loading a switch sequence, update our local copy of the switch settings.
		if sequences[sequenceNumber].Type == "switch" {
			sequences[sequenceNumber] = common.RefreshSequence(sequenceNumber, commandChannels, updateChannels)

			// Now set our local representation of switches
			for swiTchNumber, swiTch := range sequences[sequenceNumber].Switches {
				this.SwitchPositions[sequenceNumber][swiTchNumber] = swiTch.CurrentState
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
		// Show this sequence running status in the start/stop button.
		common.ShowRunningStatus(this.SelectedSequence, this.Running, eventsForLaunchpad, guiButtons)
		common.ShowStrobeButtonStatus(true, eventsForLaunchpad, guiButtons)
	} else {
		common.ShowStrobeButtonStatus(false, eventsForLaunchpad, guiButtons)
	}
}

func floodOff(this *CurrentState, sequences []*common.Sequence, dmxController *ft232.DMXController, fixturesConfig *fixture.Fixtures,
	commandChannels []chan common.Command, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, updateChannels []chan common.Sequence) {

	// Turn the flood button back to white.
	common.LightLamp(common.ALight{X: 8, Y: 3, Brightness: this.MasterBrightness, Red: 255, Green: 255, Blue: 255, Flash: false}, eventsForLaunchpad, guiButtons)

	// Send a message to stop
	cmd := common.Command{
		Action: common.StopFlood,
		Args: []common.Arg{
			{Name: "Stop Flood", Value: false},
		},
	}
	common.SendCommandToAllSequence(cmd, commandChannels)

	this.Flood = false

	// Preserve this.Blackout.
	if !this.Blackout {
		cmd := common.Command{
			Action: common.Normal,
		}
		common.SendCommandToAllSequence(cmd, commandChannels)
	}
}

func floodOn(this *CurrentState, sequences []*common.Sequence, dmxController *ft232.DMXController, fixturesConfig *fixture.Fixtures,
	commandChannels []chan common.Command, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, replyChannels []chan common.Sequence) {

	// Remember which sequence is currently selected.
	this.LastSelectedSequence = this.SelectedSequence

	// Flash the flood button pink to indicate we're in flood.
	common.FlashLight(8, 3, common.Pink, common.White, eventsForLaunchpad, guiButtons)

	// Start flood.
	cmd := common.Command{
		Action: common.Flood,
		Args: []common.Arg{
			{Name: "StartFlood", Value: true},
		},
	}
	common.SendCommandToAllSequence(cmd, commandChannels)

	this.Flood = true
}

func buttonTouched(alight common.ALight, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight) {
	common.LightLamp(common.ALight{X: alight.X, Y: alight.Y, Brightness: 255, Red: alight.OnColor.R, Green: alight.OnColor.G, Blue: alight.OnColor.B}, eventsForLaunchpad, guiButtons)
	time.Sleep(200 * time.Millisecond)
	common.LightLamp(common.ALight{X: alight.X, Y: alight.Y, Brightness: 255, Red: alight.OffColor.R, Green: alight.OffColor.G, Blue: alight.OffColor.B}, eventsForLaunchpad, guiButtons)
}

func clear(X int, Y int, this *CurrentState, sequences []*common.Sequence, dmxController *ft232.DMXController, fixturesConfig *fixture.Fixtures,
	commandChannels []chan common.Command, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, updateChannels []chan common.Sequence) {

	if debug {
		fmt.Printf("CLEAR LAUNCHPAD\n")
	}

	// Shortcut to clear rgb chase colors. We want to clear a color selection for a selected sequence.
	if sequences[this.SelectedSequence].Functions[common.Function5_Color].State &&
		sequences[this.SelectedSequence].Type != "scanner" {

		// Clear the sequence colors for this sequence.
		cmd := common.Command{
			Action: common.ClearSequenceColor,
		}
		common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

		// Get an upto date copy of the sequence.
		sequences[this.SelectedSequence] = common.RefreshSequence(this.SelectedSequence, commandChannels, updateChannels)

		// Flash the correct color buttons
		ShowRGBColorSelectionButtons(this.MasterBrightness, this.SelectedSequence, *sequences[this.SelectedSequence], eventsForLaunchpad, guiButtons)

		return
	}

	// Shortcut to clear static colors. We want to clear a static color selection for a selected sequence.
	if sequences[this.SelectedSequence].Functions[common.Function6_Static_Gobo].State &&
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
		this.FunctionSelectMode[this.SelectedSequence] = false
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
	common.UpdateStatusBar("BEAT", "beat", false, guiButtons)

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
		this.FunctionSelectMode[sequenceNumber] = false                            // Clear function selecetd mode.
		this.SelectButtonPressed[sequenceNumber] = false                           // Clear buttoned selecetd mode.
		this.EditGoboSelectionMode[sequenceNumber] = false                         // Clear edit gobo mode.
		this.EditPatternMode[sequenceNumber] = false                               // Clear edit pattern mode.
		this.EditScannerColorsMode[sequenceNumber] = false                         // Clear scanner color mode.
		this.EditSequenceColorsMode[sequenceNumber] = false                        // Clear rgb color mode.
		this.EditStaticColorsMode[sequenceNumber] = false                          // Clear static color mode.
		this.MasterBrightness = common.MaxDMXBrightness                            // Reset brightness to max.

		if sequence.Type == "scanner" {
			// Enable all scanners.
			for scannerNumber := 0; scannerNumber < sequence.NumberFixtures; scannerNumber++ {
				this.ScannerState[scannerNumber][sequence.Number].Enabled = true
				this.ScannerState[scannerNumber][sequence.Number].Inverted = false
			}
			// Tell the scanner buttons what to show.
			ShowScannerStatus(sequenceNumber, *sequences[sequenceNumber], this, eventsForLaunchpad, guiButtons, commandChannels)
		}

		// Clear all the function buttons for this sequence.
		if sequence.Type != "switch" { // Switch sequences don't have funcion keys.
			sequences[sequenceNumber].Functions[common.Function1_Pattern].State = false
			sequences[sequenceNumber].Functions[common.Function2_Auto_Color].State = false
			sequences[sequenceNumber].Functions[common.Function3_Auto_Pattern].State = false
			sequences[sequenceNumber].Functions[common.Function4_Bounce].State = false
			sequences[sequenceNumber].Functions[common.Function5_Color].State = false
			sequences[sequenceNumber].Functions[common.Function6_Static_Gobo].State = false
			sequences[sequenceNumber].Functions[common.Function7_Invert_Chase].State = false
			sequences[sequenceNumber].Functions[common.Function8_Music_Trigger].State = false
		}

		// Reset the sequence switch states back to config from the fixture config in memory.
		// And ditch any out of date copy from a loaded preset.
		if sequence.Type == "switch" {
			// Get an upto date copy of the sequence.
			sequences[this.SelectedSequence] = common.RefreshSequence(this.SelectedSequence, commandChannels, updateChannels)

			// Now set our local representation of switches
			for swiTchNumber, swiTch := range sequence.Switches {
				this.SwitchPositions[sequenceNumber][swiTchNumber] = swiTch.CurrentState
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

func getScannerShiftLabel(shift int) string {

	switch {
	case shift == 0:
		return "Sync"

	case shift == 1:
		return "1/4"

	case shift == 2:
		return "1/2"

	case shift == 3:
		return "3/4"

	}
	return ""
}

func getScannerCoordinatesLabel(shift int) string {

	switch {
	case shift == 0:
		return "12"

	case shift == 1:
		return "16"

	case shift == 2:
		return "24"

	case shift == 3:
		return "32"

	case shift == 4:
		return "64"

	}
	return ""
}

func clearAllModes(sequences []*common.Sequence, this *CurrentState) {
	for sequenceNumber, sequence := range sequences {
		this.SelectButtonPressed[sequenceNumber] = false
		this.FunctionSelectMode[sequenceNumber] = false
		this.EditSequenceColorsMode[sequenceNumber] = false
		this.EditStaticColorsMode[sequenceNumber] = false
		this.EditGoboSelectionMode[sequenceNumber] = false
		this.EditPatternMode[sequenceNumber] = false
		for function := range sequence.Functions {
			sequences[sequenceNumber].Functions[function].State = false
		}
	}
}
