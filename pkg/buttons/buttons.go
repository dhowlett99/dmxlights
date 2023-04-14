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
	"github.com/dhowlett99/dmxlights/pkg/pad"
	"github.com/dhowlett99/dmxlights/pkg/presets"
	"github.com/dhowlett99/dmxlights/pkg/sequence"
	"github.com/dhowlett99/dmxlights/pkg/sound"
	"github.com/oliread/usbdmx"
	"github.com/oliread/usbdmx/ft232"
)

const debug = false

// Select modes.
const NORMAL = 0
const FUNCTION = 1
const STATUS = 2

type CurrentState struct {
	Crash1                    bool                       // Flags to detect launchpad crash.
	Crash2                    bool                       // Flags to detect launchpad crash.
	SelectedSequence          int                        // The currently selected sequence.
	SelectedType              string                     // The currently selected sequenece type.
	LastSelectedSequence      int                        // Store fof the last selected squence.
	Speed                     map[int]int                // Local copy of sequence speed. Indexed by sequence.
	RGBShift                  map[int]int                // Current rgb fixture shift. Indexed by sequence.
	ScannerShift              map[int]int                // Current scanner shift for all fixtures.  Indexed by sequence
	RGBSize                   map[int]int                // current RGB sequence this.Size[this.SelectedSequence]. Indexed by sequence
	ScannerSize               map[int]int                // current scanner size for all fixtures. Indexed by sequence
	ScannerColor              int                        // current scanner color.
	RGBFade                   map[int]int                // Indexed by sequence.
	ScannerFade               map[int]int                // Indexed by sequence.
	ScannerCoordinates        map[int]int                // Number of coordinates for scanner patterns is selected from 4 choices. ScannerCoordinates  0=12, 1=16,2=24,3=32,4=64, Indexed by sequence.
	Running                   map[int]bool               // Which sequence is running. Indexed by sequence. True if running.
	Strobe                    map[int]bool               // We are in strobe mode. True if strobing
	StrobeSpeed               map[int]int                // Strobe speed. value is speed 0-255, indexed by sequence number.
	SavePreset                bool                       // Save a preset flag.
	Config                    bool                       // Flag to indicate we are in fixture config mode.
	Blackout                  bool                       // Blackout all fixtures.
	Flood                     bool                       // Flood all fixtures.
	SelectMode                []int                      // What mode each sequence is in : normal mode, function mode, status selection mode.
	Functions                 map[int][]common.Function  // Map indexed sequence of functions
	FunctionLabels            [8]string                  // Storage for the function key labels for this sequence.
	SelectButtonPressed       []bool                     // Which sequence has its Select button pressed.
	SwitchPositions           [9][9]int                  // Sorage for switch positions.
	EditSequenceColorsMode    []bool                     // This flag is true when the sequence is in sequence colors editing mode.
	EditScannerColorsMode     []bool                     // This flag is true when the sequence is in select scanner colors editing mode.
	EditGoboSelectionMode     []bool                     // This flag is true when the sequence is in sequence gobo selection mode.
	EditStaticColorsMode      []bool                     // This flag is true when the sequence is in static colors editing mode.
	EditPatternMode           []bool                     // This flag is true when the sequence is in pattern editing mode.
	EditFixtureSelectionMode  bool                       // This flag is true when the sequence is in select fixture mode.
	MasterBrightness          int                        // Affects all DMX fixtures and launchpad lamps.
	LastStaticColorButtonX    int                        // Which Static Color button did we change last.
	LastStaticColorButtonY    int                        // Which Static Color button did we change last.
	SoundGain                 float32                    // Fine gain -0.09 -> 0.09
	FixtureState              [][]common.FixtureState    // Which fixture is enabled: bool and inverted: bool on which sequence.
	SelectedFixture           int                        // Which fixture is selected when changing scanner color or gobo.
	FollowingAction           string                     // String to find next function, used in selecting a fixture.
	OffsetPan                 int                        // Offset for Pan.
	OffsetTilt                int                        // Offset for Tilt.
	Pad                       *pad.Pad                   // Pointer to the Novation Launchpad object.
	PresetsStore              map[string]presets.Preset  // Storage for the Presets.
	LastPreset                *string                    // Last preset used.
	SoundTriggers             []*common.Trigger          // Pointer to the Sound Triggers.
	SoundConfig               *sound.SoundConfig         // Pointer to the sound config struct.
	SequenceChannels          common.Channels            // Channles used to communicate with the sequence.
	RGBPatterns               map[int]common.Pattern     // Available RGB Patterns.
	ScannerPattern            int                        // The selected scanner pattern Number. Used as the index for above.
	Pattern                   int                        // The selected RGB pattern Number. Used as the index for above.
	StaticButtons             []common.StaticColorButton // Storage for the color of the static buttons.
	SelectedGobo              int                        // The selected GOBO.
	ButtonTimer               *time.Time                 // Button Timer
	SelectColorBar            map[int]int                // Storage for color bar in static color selection. Indexed by sequence.
	SwitchChannels            []common.SwitchChannel     // Used for communicating with mini-sequencers on switches.
	LaunchPadConnected        bool                       // Flag to indicate presence of Novation Launchpad.
	DmxInterfacePresent       bool                       // Flag to indicate precence of DMX interface card
	DmxInterfacePresentConfig *usbdmx.ControllerConfig   // DMX Interface card config.
	LaunchpadName             string                     // Storage for launchpad config.
	Chaser                    common.Sequence            // Sequence for chaser.
	ScannerChaser             bool                       // Chaser is running.
	SwitchSequenceNumber      int                        // Switch sequence number, setup at start.
	ChaserSequenceNumber      int                        // Chaser sequence number, setup at start.
	ScannerSequenceNumber     int                        // Scanner sequence number, setup at start.
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

	// The Novation Launchpad is not designed for the number of MIDI
	// Events we send when all the sequences are chasing at top
	// Speed, so we look out for the crys for help when the Launchpad
	// Crashes. When we see the following three events we reset the pad.
	// As this happens so quickly the user should be unware that a crash
	// has taken place.
	if X == -1 && Y == 8 {
		this.Crash1 = true
		return
	}
	if X == 1 && Y == 8 && this.Crash1 {
		this.Crash2 = true
		return
	}
	// Crash 2 message has appeared and this isn't a pad program ack.
	if X != 0 && Y == 8 && this.Crash2 {
		// Start a supervisor thread which will reset the launchpad every 1/2 second.
		time.Sleep(200 * time.Millisecond)
		if this.LaunchPadConnected {
			this.Pad.Program()
		}
		InitButtons(this, eventsForLaunchpad, guiButtons)

		// Show the static and switch settings.
		cmd := common.Command{
			Action: common.UnHide,
		}
		common.SendCommandToAllSequence(cmd, commandChannels)

		// Show the presets again.
		presets.RefreshPresets(eventsForLaunchpad, guiButtons, this.PresetsStore)
		this.Crash1 = false
		this.Crash2 = false
		return
	}

	// F L A S H   O N   B U T T O N S - Briefly light (flash) the fixtures based on color pattern.
	if X >= 0 &&
		X < 8 &&
		Y >= 0 &&
		Y < 4 &&
		!this.Functions[Y][common.Function1_Pattern].State &&
		!this.Functions[Y][common.Function6_Static_Gobo].State &&
		!this.Functions[Y][common.Function5_Color].State &&
		sequences[Y].Type != "switch" && // As long as we're not a switch sequence.
		sequences[Y].Type != "scanner" && // As long as we're not a scanner sequence.
		this.SelectMode[Y] == NORMAL { // As long as we're in normal mode for this sequence.

		if debug {
			fmt.Printf("Flash ON Fixture Pressed X:%d Y:%d\n", X, Y)
		}
		colorPattern := 5
		flashSequence := common.Sequence{
			Pattern: common.Pattern{
				Name:  "colors",
				Steps: this.RGBPatterns[colorPattern].Steps, // Use the color pattern for flashing.
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
		fixture.MapFixtures(false, false, Y, dmxController, X, red, green, blue, white, amber, uv, pan, tilt, shutter, rotate, music, program, gobo, 0, fixturesConfig, this.Blackout, this.MasterBrightness, this.MasterBrightness, this.Strobe[this.SelectedSequence], this.StrobeSpeed[this.SelectedSequence], this.DmxInterfacePresent)

		if gui {
			time.Sleep(200 * time.Millisecond)
			common.LightLamp(common.ALight{X: X, Y: Y, Brightness: 0, Red: 0, Green: 0, Blue: 0}, eventsForLaunchpad, guiButtons)
			fixture.MapFixtures(false, false, Y, dmxController, X, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, this.Blackout, this.MasterBrightness, this.MasterBrightness, this.Strobe[this.SelectedSequence], this.StrobeSpeed[this.SelectedSequence], this.DmxInterfacePresent)
		}

		return
	}

	// F L A S H  O F F   B U T T O N S - Briefly light (flash) the fixtures based on current pattern.
	if X >= 0 &&
		X != 108 && X != 117 &&
		X >= 100 && X < 117 &&
		Y >= 0 && Y < 4 &&
		!this.Functions[Y][common.Function1_Pattern].State &&
		!this.Functions[Y][common.Function6_Static_Gobo].State &&
		!this.Functions[Y][common.Function5_Color].State &&
		sequences[Y].Type != "switch" && // As long as we're not a switch sequence.
		sequences[Y].Type != "scanner" && // As long as we're not a scanner sequence.
		this.SelectMode[Y] == NORMAL { // As long as we're in normal mode for this sequence.

		if debug {
			fmt.Printf("Flash OFF Fixture Pressed X:%d Y:%d\n", X, Y)
		}

		X = X - 100

		common.LightLamp(common.ALight{X: X, Y: Y, Brightness: this.MasterBrightness, Red: 0, Green: 0, Blue: 0}, eventsForLaunchpad, guiButtons)
		fixture.MapFixtures(false, false, Y, dmxController, X, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, this.Blackout, this.MasterBrightness, this.MasterBrightness, this.Strobe[this.SelectedSequence], this.StrobeSpeed[this.SelectedSequence], this.DmxInterfacePresent)
		return
	}

	// P R E S E T S - recall (short press) or delete (long press) the preset.
	if X >= 100 && X < 108 &&
		(Y > 3 && Y < 7) {

		if debug {
			fmt.Printf("Preset Pressed X:%d Y:%d\n", X, Y)
		}

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
			this.PresetsStore[fmt.Sprint(X)+","+fmt.Sprint(Y)] = presets.Preset{State: false, Selected: false, Label: "", ButtonColor: ""}

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

		if debug {
			fmt.Printf("GUI Clear Pressed X:%d Y:%d\n", X, Y)
		}

		clear(X, Y, this, sequences, dmxController, fixturesConfig, commandChannels, eventsForLaunchpad, guiButtons, updateChannels)
		return
	}

	// C L E A R  - Start the timer, waiting for a long press to clear all.
	// Because a short press in scanner mode shifts the scanners up.
	if X == 0 && Y == -1 && !gui && sequences[this.SelectedSequence].Type == "scanner" {

		if debug {
			fmt.Printf("Clear Pressed Start Timer X:%d Y:%d\n", X, Y)
		}
		// Start a timer for this button.
		here := time.Now()
		this.ButtonTimer = &here
		return
	}

	//  C L E A R - clear all if we're not in the scanner mode.
	if X == 0 && Y == -1 && !gui && sequences[this.SelectedSequence].Type != "scanner" {
		if debug {
			fmt.Printf("Clear All If We're Not in Scanner Mode X:%d Y:%d\n", X, Y)
		}
		clear(X, Y, this, sequences, dmxController, fixturesConfig, commandChannels, eventsForLaunchpad, guiButtons, updateChannels)
		return
	}

	// C L E A R  - We have a long press.
	if X == 100 && Y == -1 && !gui && sequences[this.SelectedSequence].Type == "scanner" {

		if debug {
			fmt.Printf("Clear Pressed Long Press X:%d Y:%d\n", X, Y)
		}

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
	if X == 8 && Y == 3 {

		if debug {
			fmt.Printf("Start FLood X:%d Y:%d\n", X, Y)
		}

		// Turn off the flashing save button
		this.SavePreset = false
		common.LightLamp(common.ALight{X: 8, Y: 4, Brightness: this.MasterBrightness, Red: 255, Green: 255, Blue: 255}, eventsForLaunchpad, guiButtons)

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
			this.PresetsStore[location] = presets.Preset{State: true, Selected: true, Label: current.Label, ButtonColor: current.ButtonColor}
			this.LastPreset = &location

			config.AskToSaveConfig(commandChannels, replyChannels, X, Y)

			// turn off the save button from flashing.
			common.LightLamp(common.ALight{X: 8, Y: 4, Brightness: this.MasterBrightness, Red: 255, Green: 255, Blue: 255}, eventsForLaunchpad, guiButtons)

			presets.SavePresets(this.PresetsStore)

			// clear any selected preset.
			for location, preset := range this.PresetsStore {
				if preset.State && preset.Selected {
					this.PresetsStore[location] = presets.Preset{State: preset.State, Selected: false, Label: preset.Label, ButtonColor: "Red"}
				}
			}

			// Select this location and flash its button.
			this.PresetsStore[location] = presets.Preset{State: true, Selected: true, Label: current.Label, ButtonColor: current.ButtonColor}
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

		// If we're a scanner and we're in shutter chase mode.
		var targetSequence int
		if sequences[this.SelectedSequence].Type == "scanner" &&
			this.ScannerChaser {
			targetSequence = this.ChaserSequenceNumber
		} else {
			targetSequence = this.SelectedSequence
		}

		if sequences[targetSequence].Type == "rgb" || this.Functions[targetSequence][common.Function7_Invert_Chase].State {
			this.RGBShift[targetSequence] = this.RGBShift[targetSequence] - 1
			if this.RGBShift[targetSequence] < 0 {
				this.RGBShift[targetSequence] = 0
			}
			cmd := common.Command{
				Action: common.UpdateRGBShift,
				Args: []common.Arg{
					{Name: "RGBShift", Value: this.RGBShift[targetSequence]},
				},
			}
			common.SendCommandToSequence(targetSequence, cmd, commandChannels)

			// Update the status bar
			common.UpdateStatusBar(fmt.Sprintf("Shift %02d", this.RGBShift[targetSequence]), "shift", false, guiButtons)
			return
		}

		if sequences[targetSequence].Type == "scanner" {
			this.ScannerShift[targetSequence] = this.ScannerShift[targetSequence] - 1
			if this.ScannerShift[targetSequence] < 0 {
				this.ScannerShift[targetSequence] = 0
			}
			cmd := common.Command{
				Action: common.UpdateScannerShift,
				Args: []common.Arg{
					{Name: "ScannerShift", Value: this.ScannerShift[targetSequence]},
				},
			}
			common.SendCommandToSequence(targetSequence, cmd, commandChannels)

			// Update the status bar
			label := getScannerShiftLabel(this.ScannerShift[targetSequence])
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

		// If we're a scanner and we're in shutter chase mode.
		var targetSequence int
		if sequences[this.SelectedSequence].Type == "scanner" &&
			this.ScannerChaser {
			targetSequence = this.ChaserSequenceNumber
		} else {
			targetSequence = this.SelectedSequence
		}

		if sequences[targetSequence].Type == "rgb" || this.Functions[targetSequence][common.Function7_Invert_Chase].State {
			this.RGBShift[targetSequence] = this.RGBShift[targetSequence] + 1
			if this.RGBShift[targetSequence] > 50 {
				this.RGBShift[targetSequence] = 50
			}
			cmd := common.Command{
				Action: common.UpdateRGBShift,
				Args: []common.Arg{
					{Name: "Shift", Value: this.RGBShift[targetSequence]},
				},
			}
			common.SendCommandToSequence(targetSequence, cmd, commandChannels)

			// Update the status bar
			common.UpdateStatusBar(fmt.Sprintf("Shift %02d", this.RGBShift[targetSequence]), "shift", false, guiButtons)
			return
		}

		if sequences[targetSequence].Type == "scanner" {
			this.ScannerShift[targetSequence] = this.ScannerShift[targetSequence] + 1
			if this.ScannerShift[targetSequence] > 3 {
				this.ScannerShift[targetSequence] = 3
			}
			cmd := common.Command{
				Action: common.UpdateScannerShift,
				Args: []common.Arg{
					{Name: "ScannerShift", Value: this.ScannerShift[targetSequence]},
				},
			}
			common.SendCommandToSequence(targetSequence, cmd, commandChannels)

			// Update the status bar
			label := getScannerShiftLabel(this.ScannerShift[targetSequence])
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

		// If we're a scanner and we're in shutter chase mode.
		var targetSequence int
		if sequences[this.SelectedSequence].Type == "scanner" &&
			this.ScannerChaser {
			targetSequence = this.ChaserSequenceNumber
		} else {
			targetSequence = this.SelectedSequence
		}

		// Decrease Strobe Speed.
		if this.Strobe[targetSequence] {
			this.StrobeSpeed[targetSequence] -= 10
			if this.StrobeSpeed[targetSequence] < 0 {
				this.StrobeSpeed[targetSequence] = 0
			}
			cmd := common.Command{
				Action: common.UpdateStrobeSpeed,
				Args: []common.Arg{
					{Name: "STROBE_SPEED", Value: this.StrobeSpeed[targetSequence]},
				},
			}
			common.SendCommandToSequence(targetSequence, cmd, commandChannels)
			// Update the status bar
			common.UpdateStatusBar(fmt.Sprintf("Strobe %02d", this.StrobeSpeed[targetSequence]), "speed", false, guiButtons)
			return
		}

		// Get an upto date copy of the sequence.
		sequences[targetSequence] = common.RefreshSequence(targetSequence, commandChannels, updateChannels)

		// Decrease Speed.
		if !sequences[targetSequence].MusicTrigger {
			this.Speed[targetSequence]--
			if this.Speed[targetSequence] < 0 {
				this.Speed[targetSequence] = 1
			}
			cmd := common.Command{
				Action: common.UpdateSpeed,
				Args: []common.Arg{
					{Name: "Speed", Value: this.Speed[targetSequence]},
				},
			}
			common.SendCommandToSequence(targetSequence, cmd, commandChannels)
		}

		// Update the status bar
		common.UpdateStatusBar(fmt.Sprintf("Speed %02d", this.Speed[targetSequence]), "speed", false, guiButtons)

		return
	}

	// Increase speed of selected sequence.
	if X == 1 && Y == 7 {

		if debug {
			fmt.Printf("Increase Speed \n")
		}

		buttonTouched(common.ALight{X: X, Y: Y, OnColor: common.White, OffColor: common.Cyan}, eventsForLaunchpad, guiButtons)

		// If we're a scanner and we're in shutter chase mode.
		var targetSequence int
		if sequences[this.SelectedSequence].Type == "scanner" &&
			this.ScannerChaser {
			targetSequence = this.ChaserSequenceNumber
		} else {
			targetSequence = this.SelectedSequence
		}

		if this.Strobe[targetSequence] {
			this.StrobeSpeed[targetSequence] += 10
			if this.StrobeSpeed[targetSequence] > 255 {
				this.StrobeSpeed[targetSequence] = 255
			}
			cmd := common.Command{
				Action: common.UpdateStrobeSpeed,
				Args: []common.Arg{
					{Name: "STROBE_SPEED", Value: this.StrobeSpeed[targetSequence]},
				},
			}
			common.SendCommandToSequence(targetSequence, cmd, commandChannels)
			// Update the status bar
			common.UpdateStatusBar(fmt.Sprintf("Strobe %02d", this.StrobeSpeed[targetSequence]), "speed", false, guiButtons)
			return
		}

		// Get an upto date copy of the sequence.
		sequences[targetSequence] = common.RefreshSequence(targetSequence, commandChannels, updateChannels)

		if !sequences[targetSequence].MusicTrigger {
			this.Speed[targetSequence]++
			if this.Speed[targetSequence] > 12 {
				this.Speed[targetSequence] = 12
			}
			cmd := common.Command{
				Action: common.UpdateSpeed,
				Args: []common.Arg{
					{Name: "Speed", Value: this.Speed[targetSequence]},
				},
			}
			common.SendCommandToSequence(targetSequence, cmd, commandChannels)
		}

		// Update the status bar
		common.UpdateStatusBar(fmt.Sprintf("Speed %02d", this.Speed[targetSequence]), "speed", false, guiButtons)

		return
	}

	// S E L E C T    S E Q U E N C E.
	// Select sequence 1.
	if X == 8 && Y == 0 {

		this.SelectedSequence = 0
		this.SelectedType = sequences[this.SelectedSequence].Type

		if debug {
			fmt.Printf("Select Sequence %d Type %s\n", this.SelectedSequence, this.SelectedType)
		}

		HandleSelect(sequences, this, eventsForLaunchpad, commandChannels, guiButtons)

		this.EditSequenceColorsMode[this.SelectedSequence] = false
		this.EditGoboSelectionMode[this.SelectedSequence] = false

		return
	}

	// Select sequence 2.
	if X == 8 && Y == 1 {

		this.SelectedSequence = 1
		this.SelectedType = sequences[this.SelectedSequence].Type

		if debug {
			fmt.Printf("Select Sequence %d Type %s\n", this.SelectedSequence, this.SelectedType)
		}

		HandleSelect(sequences, this, eventsForLaunchpad, commandChannels, guiButtons)

		this.EditSequenceColorsMode[this.SelectedSequence] = false
		this.EditGoboSelectionMode[this.SelectedSequence] = false

		return
	}

	// Select sequence 3.
	if X == 8 && Y == 2 {

		this.SelectedSequence = 2
		this.SelectedType = sequences[this.SelectedSequence].Type

		if debug {
			fmt.Printf("Select Sequence %d Type %s\n", this.SelectedSequence, this.SelectedType)
		}

		HandleSelect(sequences, this, eventsForLaunchpad, commandChannels, guiButtons)
		this.EditSequenceColorsMode[this.SelectedSequence] = false
		this.EditGoboSelectionMode[this.SelectedSequence] = false

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

			// Clear the pattern function keys
			common.ClearSelectedRowOfButtons(this.SelectedSequence, eventsForLaunchpad, guiButtons)

			// Reveal the now running sequence
			common.RevealSequence(this.SelectedSequence, commandChannels)

			return
		}

	}

	// S T R O B E - Strobe.
	if X == 8 && Y == 6 {

		if debug {
			fmt.Printf("Strobe X:%d Y:%d\n", X, Y)
		}

		// Turn off the flashing save button
		this.SavePreset = false
		this.SavePreset = false
		common.LightLamp(common.ALight{X: 8, Y: 4, Brightness: this.MasterBrightness, Red: 255, Green: 255, Blue: 255}, eventsForLaunchpad, guiButtons)

		// Shutdown any function bars.
		clearAllModes(sequences, this)
		HandleSelect(sequences, this, eventsForLaunchpad, commandChannels, guiButtons)

		// If strobing, stop it
		if this.Strobe[this.SelectedSequence] {
			// Stop strobing this sequence.
			cmd := common.Command{
				Action: common.Strobe,
				Args: []common.Arg{
					{Name: "STROBE_STATE", Value: false},
					{Name: "STROBE_SPEED", Value: this.StrobeSpeed[this.SelectedSequence]},
				},
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
					{Name: "STROBE_STATE", Value: true},
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

		// If we're a scanner and we're in shutter chase mode.
		var targetSequence int
		if sequences[this.SelectedSequence].Type == "scanner" &&
			this.ScannerChaser {
			targetSequence = this.ChaserSequenceNumber
		} else {
			targetSequence = this.SelectedSequence
		}

		if sequences[targetSequence].Type == "rgb" || this.Functions[targetSequence][common.Function7_Invert_Chase].State {
			// Send Update RGB Size.
			this.RGBSize[targetSequence]--
			if this.RGBSize[targetSequence] < 1 {
				this.RGBSize[targetSequence] = 1
			}
			cmd := common.Command{
				Action: common.UpdateRGBSize,
				Args: []common.Arg{
					{Name: "RGBSize", Value: this.RGBSize[targetSequence]},
				},
			}
			common.SendCommandToSequence(targetSequence, cmd, commandChannels)
			// Update the status bar
			common.UpdateStatusBar(fmt.Sprintf("Size %02d", this.RGBSize[targetSequence]), "size", false, guiButtons)
			return
		}

		if sequences[targetSequence].Type == "scanner" {
			// Send Update Scanner Size.
			this.ScannerSize[targetSequence] = this.ScannerSize[targetSequence] - 10
			if this.ScannerSize[targetSequence] < 0 {
				this.ScannerSize[targetSequence] = 0
			}
			cmd := common.Command{
				Action: common.UpdateScannerSize,
				Args: []common.Arg{
					{Name: "ScannerSize", Value: this.ScannerSize[targetSequence]},
				},
			}
			common.SendCommandToSequence(targetSequence, cmd, commandChannels)

			// Update the status bar
			common.UpdateStatusBar(fmt.Sprintf("Size %02d", this.ScannerSize[targetSequence]), "size", false, guiButtons)
			return
		}
	}

	// Increase Size.
	if X == 5 && Y == 7 {

		if debug {
			fmt.Printf("Increase Size\n")
		}

		buttonTouched(common.ALight{X: X, Y: Y, OnColor: common.White, OffColor: common.Cyan}, eventsForLaunchpad, guiButtons)

		// If we're a scanner and we're in shutter chase mode.
		var targetSequence int
		if sequences[this.SelectedSequence].Type == "scanner" &&
			this.ScannerChaser {
			targetSequence = this.ChaserSequenceNumber
		} else {
			targetSequence = this.SelectedSequence
		}

		if sequences[targetSequence].Type == "rgb" ||
			this.Functions[targetSequence][common.Function7_Invert_Chase].State {

			// Send Update RGB Size.
			this.RGBSize[targetSequence]++
			if this.RGBSize[targetSequence] > common.MaxRGBSize {
				this.RGBSize[targetSequence] = common.MaxRGBSize
			}
			cmd := common.Command{
				Action: common.UpdateRGBSize,
				Args: []common.Arg{
					{Name: "RGBSize", Value: this.RGBSize[targetSequence]},
				},
			}
			common.SendCommandToSequence(targetSequence, cmd, commandChannels)
			// Update the status bar
			common.UpdateStatusBar(fmt.Sprintf("Size %02d", this.RGBSize[targetSequence]), "size", false, guiButtons)
			return
		}

		if sequences[targetSequence].Type == "scanner" {
			// Send Update Scanner Size.
			this.ScannerSize[targetSequence] = this.ScannerSize[targetSequence] + 10
			if this.ScannerSize[targetSequence] > common.MaxScannerSize {
				this.ScannerSize[targetSequence] = common.MaxScannerSize
			}
			cmd := common.Command{
				Action: common.UpdateScannerSize,
				Args: []common.Arg{
					{Name: "ScannerSize", Value: this.ScannerSize[targetSequence]},
				},
			}
			common.SendCommandToSequence(targetSequence, cmd, commandChannels)

			// Update the status bar
			common.UpdateStatusBar(fmt.Sprintf("Size %02d", this.ScannerSize[targetSequence]), "size", false, guiButtons)

			return
		}
	}

	// Fade time decrease.
	if X == 6 && Y == 7 {

		if debug {
			fmt.Printf("Decrease Fade Time\n")
		}

		buttonTouched(common.ALight{X: X, Y: Y, OnColor: common.White, OffColor: common.Cyan}, eventsForLaunchpad, guiButtons)

		// If we're a scanner and we're in shutter chase mode.
		var targetSequence int
		if sequences[this.SelectedSequence].Type == "scanner" &&
			this.ScannerChaser {
			targetSequence = this.ChaserSequenceNumber
		} else {
			targetSequence = this.SelectedSequence
		}

		if sequences[targetSequence].Type == "rgb" || this.Functions[targetSequence][common.Function7_Invert_Chase].State {
			this.RGBFade[targetSequence]--
			if this.RGBFade[targetSequence] < 1 {
				this.RGBFade[targetSequence] = 1
			}
			// Send fade update command.
			cmd := common.Command{
				Action: common.UpdateRGBFadeSpeed,
				Args: []common.Arg{
					{Name: "RGBFadeSpeed", Value: this.RGBFade[targetSequence]},
				},
			}
			common.SendCommandToSequence(targetSequence, cmd, commandChannels)
			// Update the status bar
			common.UpdateStatusBar(fmt.Sprintf("Fade %02d", this.RGBFade[targetSequence]), "fade", false, guiButtons)

			return
		}

		// Update Coordinates.
		if sequences[targetSequence].Type == "scanner" {
			// Fade also send more or less coordinates for the scanner patterns.
			this.ScannerCoordinates[targetSequence]--
			if this.ScannerCoordinates[targetSequence] < 0 {
				this.ScannerCoordinates[targetSequence] = 0
			}
			cmd := common.Command{
				Action: common.UpdateNumberCoordinates,
				Args: []common.Arg{
					{Name: "NumberCoordinates", Value: this.ScannerCoordinates[targetSequence]},
				},
			}
			common.SendCommandToSequence(targetSequence, cmd, commandChannels)
			// Update the status bar
			label := getScannerCoordinatesLabel(this.ScannerCoordinates[targetSequence])
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

		// If we're a scanner and we're in shutter chase mode.
		var targetSequence int
		if sequences[this.SelectedSequence].Type == "scanner" &&
			this.ScannerChaser {
			targetSequence = this.ChaserSequenceNumber
		} else {
			targetSequence = this.SelectedSequence
		}

		if sequences[targetSequence].Type == "rgb" || this.Functions[targetSequence][common.Function7_Invert_Chase].State {
			this.RGBFade[targetSequence]++
			if this.RGBFade[targetSequence] > common.MaxRGBFade {
				this.RGBFade[targetSequence] = common.MaxRGBFade
			}
			// Send fade update command.
			cmd := common.Command{
				Action: common.UpdateRGBFadeSpeed,
				Args: []common.Arg{
					{Name: "FadeSpeed", Value: this.RGBFade[targetSequence]},
				},
			}
			common.SendCommandToSequence(targetSequence, cmd, commandChannels)
			// Update the status bar
			common.UpdateStatusBar(fmt.Sprintf("Fade %02d", this.RGBFade[targetSequence]), "fade", false, guiButtons)
			return
		}

		// Update Coordinates.
		if sequences[targetSequence].Type == "scanner" {
			// Fade also send more or less coordinates for the scanner patterns.
			this.ScannerCoordinates[targetSequence]++
			if this.ScannerCoordinates[targetSequence] > 4 {
				this.ScannerCoordinates[targetSequence] = 4
			}
			cmd := common.Command{
				Action: common.UpdateNumberCoordinates,
				Args: []common.Arg{
					{Name: "NumberCoordinates", Value: this.ScannerCoordinates[targetSequence]},
				},
			}
			common.SendCommandToSequence(targetSequence, cmd, commandChannels)
			// Update the status bar
			label := getScannerCoordinatesLabel(this.ScannerCoordinates[targetSequence])
			common.UpdateStatusBar(fmt.Sprintf("Coord %s", label), "fade", false, guiButtons)
			return
		}
	}

	// S W I T C H   B U T T O N's Toggle State of switches for this sequence.
	if X >= 0 && X < 8 && this.SelectMode[this.SelectedSequence] == NORMAL &&
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

	// D I S A B L E  / E N A B L E   F I X T U R E  - Used to toggle the scanner state from on, inverted or off.
	if X >= 0 && X < 8 &&
		Y >= 0 &&
		Y < 4 &&
		this.SelectMode[this.SelectedSequence] == STATUS &&
		!this.Functions[this.SelectedSequence][common.Function1_Pattern].State &&
		!this.Functions[this.SelectedSequence][common.Function6_Static_Gobo].State &&
		!this.Functions[this.SelectedSequence][common.Function5_Color].State {

		if debug {
			fmt.Printf("Disable Fixture X:%d Y:%d\n", X, Y)
			fmt.Printf("Fixture State Enabled %t  Inverted %t\n", this.FixtureState[X][Y].Enabled, this.FixtureState[X][Y].Inverted)
		}

		// Disable fixture if we're already enabled and inverted.
		if this.FixtureState[X][Y].Enabled && this.FixtureState[X][Y].Inverted && X < sequences[Y].NumberFixtures {
			if debug {
				fmt.Printf("Disable fixture Number %d State on Sequence %d to false\n", X, Y)
			}

			this.FixtureState[X][Y].Enabled = false
			this.FixtureState[X][Y].Inverted = false

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

			if this.ScannerChaser && this.SelectedType == "scanner" {
				// Tell the sequence chaser to turn on this scanner.
				cmd = common.Command{
					Action: common.ToggleFixtureState,
					Args: []common.Arg{
						{Name: "SequenceNumber", Value: 4},
						{Name: "FixtureNumber", Value: X},
						{Name: "FixtureState", Value: false},
						{Name: "FixtureInverted", Value: false},
					},
				}
				common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)
			}

			// Show the status.
			ShowFixtureStatus(Y, *sequences[Y], this, eventsForLaunchpad, guiButtons, commandChannels)

			return
		}

		// Enable scanner if not enabled and not inverted.
		if !this.FixtureState[X][Y].Enabled && !this.FixtureState[X][Y].Inverted && X < sequences[Y].NumberFixtures {

			if debug {
				fmt.Printf("Enable fixture number %d State on Sequence %d to true [Scanners:%d]\n", X, Y, sequences[Y].NumberFixtures)
			}

			this.FixtureState[X][Y].Enabled = true
			this.FixtureState[X][Y].Inverted = false

			// Tell the sequence to turn on this scanner.
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

			if this.ScannerChaser && this.SelectedType == "scanner" {
				// Tell the sequence chaseer to turn on this scanner.
				cmd = common.Command{
					Action: common.ToggleFixtureState,
					Args: []common.Arg{
						{Name: "SequenceNumber", Value: 4},
						{Name: "FixtureNumber", Value: X},
						{Name: "FixtureState", Value: true},
						{Name: "FixtureInverted", Value: false},
					},
				}
				common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)
			}

			// Show the status.
			ShowFixtureStatus(Y, *sequences[Y], this, eventsForLaunchpad, guiButtons, commandChannels)

			return

		}

		// Invert scanner if we're enabled but not inverted.
		if this.FixtureState[X][Y].Enabled && !this.FixtureState[X][Y].Inverted && X < sequences[Y].NumberFixtures {

			if debug {
				fmt.Printf("Invert Scanner Number %d State on Sequence %d to false\n", X, Y)
			}

			this.FixtureState[X][Y].Enabled = true
			this.FixtureState[X][Y].Inverted = true

			// Tell the sequence to invert this scanner.
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

			if this.ScannerChaser && this.SelectedType == "scanner" {
				// Tell the sequence chaser to invert this scanner.
				cmd = common.Command{
					Action: common.ToggleFixtureState,
					Args: []common.Arg{
						{Name: "SequenceNumber", Value: 4},
						{Name: "FixtureNumber", Value: X},
						{Name: "FixtureState", Value: true},
						{Name: "FixtureInverted", Value: true},
					},
				}
				common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)
			}

			// Show the status.
			ShowFixtureStatus(Y, *sequences[Y], this, eventsForLaunchpad, guiButtons, commandChannels)

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

		if this.Functions[this.SelectedSequence][common.Function6_Static_Gobo].State {

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

		if this.Functions[this.SelectedSequence][common.Function6_Static_Gobo].State {

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

		if this.Functions[this.SelectedSequence][common.Function6_Static_Gobo].State {

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
		this.Functions[this.SelectedSequence][common.Function5_Color].State {

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
		this.Functions[this.SelectedSequence][common.Function5_Color].State {

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

		// If configured set scanner color in chaser.
		if this.ScannerChaser && this.SelectedType == "scanner" {
			cmd := common.Command{
				Action: common.UpdateScannerColor,
				Args: []common.Arg{
					{Name: "SelectedColor", Value: this.ScannerColor},
					{Name: "SelectedFixture", Value: this.SelectedFixture},
				},
			}
			common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)
		}

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
		this.Functions[this.SelectedSequence][common.Function6_Static_Gobo].State || this.Functions[this.SelectedSequence][common.Function5_Color].State &&
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
		this.Functions[this.SelectedSequence][common.Function6_Static_Gobo].State {

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

		// If configured set scanner color in chaser.
		if this.ScannerChaser && this.SelectedType == "scanner" {
			cmd := common.Command{
				Action: common.UpdateGobo,
				Args: []common.Arg{
					{Name: "SelectedGobo", Value: this.SelectedGobo},
					{Name: "FixtureNumber", Value: this.SelectedFixture},
				},
			}
			common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)
		}

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
		this.SelectMode[this.SelectedSequence] == NORMAL && // Not in function Mode
		this.EditStaticColorsMode[this.SelectedSequence] { // Static Function On

		if debug {
			fmt.Printf("Update Static for X %d\n", X)
		}

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

		if this.ScannerChaser && this.SelectedType == "scanner" {
			// Tell the sequence to change the pattern.
			cmd := common.Command{
				Action: common.UpdatePattern,
				Args: []common.Arg{
					{Name: "SelectPattern", Value: X},
				},
			}
			common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)

			this.SelectMode[this.SelectedSequence] = NORMAL

			// Get an upto date copy of the sequence.
			sequences[this.ChaserSequenceNumber] = common.RefreshSequence(this.ChaserSequenceNumber, commandChannels, updateChannels)

			// We call ShowPatternSelectionButtons here so the selections will flash as you press them.
			this.EditFixtureSelectionMode = false
			ShowPatternSelectionButtons(this, sequences[this.SelectedSequence].Master, *sequences[this.ChaserSequenceNumber], eventsForLaunchpad, guiButtons)

		} else {
			// Tell the sequence to change the pattern.
			cmd := common.Command{
				Action: common.UpdatePattern,
				Args: []common.Arg{
					{Name: "SelectPattern", Value: X},
				},
			}
			common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

			this.SelectMode[this.SelectedSequence] = NORMAL

			// Get an upto date copy of the sequence.
			sequences[this.SelectedSequence] = common.RefreshSequence(this.SelectedSequence, commandChannels, updateChannels)

			// We call ShowPatternSelectionButtons here so the selections will flash as you press them.
			this.EditFixtureSelectionMode = false
			ShowPatternSelectionButtons(this, sequences[this.SelectedSequence].Master, *sequences[this.SelectedSequence], eventsForLaunchpad, guiButtons)

		}

		return
	}

	// F U N C T I O N  K E Y S
	if X >= 0 && X < 8 &&
		this.SelectMode[this.SelectedSequence] == FUNCTION {

		if debug {
			fmt.Printf("Function Key X:%d Y:%d\n", X, Y)
		}

		// Map Function 1 - Go straight into pattern select mode, don't wait for a another select press.
		if X == common.Function1_Pattern {
			this.EditPatternMode[this.SelectedSequence] = true
			this.Functions[this.SelectedSequence][common.Function1_Pattern].State = true
			common.ClearSelectedRowOfButtons(this.SelectedSequence, eventsForLaunchpad, guiButtons)
			this.EditFixtureSelectionMode = false
			if !this.ScannerChaser {
				ShowPatternSelectionButtons(this, sequences[this.SelectedSequence].Master, *sequences[this.SelectedSequence], eventsForLaunchpad, guiButtons)
			} else {
				ShowPatternSelectionButtons(this, sequences[this.SelectedSequence].Master, *sequences[this.ChaserSequenceNumber], eventsForLaunchpad, guiButtons)
			}
			return
		}

		// Function 2 Set Auto Color - Toggle the auto color feature in scanner and rgb sequences.
		if X == common.Function2_Auto_Color &&
			!this.Functions[this.SelectedSequence][common.Function2_Auto_Color].State &&
			!this.ScannerChaser {

			this.Functions[this.SelectedSequence][common.Function2_Auto_Color].State = true
			cmd := common.Command{
				Action: common.UpdateAutoColor,
				Args: []common.Arg{
					{Name: "AutoColor", Value: true},
					{Name: "Type", Value: this.SelectedType},
				},
			}
			common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)
			ShowFunctionButtons(this, this.SelectedSequence, eventsForLaunchpad, guiButtons)
			return
		}
		if X == common.Function2_Auto_Color &&
			this.Functions[this.SelectedSequence][common.Function2_Auto_Color].State &&
			!this.ScannerChaser {

			this.Functions[this.SelectedSequence][common.Function2_Auto_Color].State = false
			cmd := common.Command{
				Action: common.UpdateAutoColor,
				Args: []common.Arg{
					{Name: "AutoColor", Value: false},
					{Name: "Type", Value: this.SelectedType},
				},
			}
			common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)
			ShowFunctionButtons(this, this.SelectedSequence, eventsForLaunchpad, guiButtons)
			return
		}

		// Function 2 Set Auto Color - Toggle the auto color feature in chaser sequences.
		if X == common.Function2_Auto_Color &&
			!this.Functions[this.SelectedSequence][common.Function2_Auto_Color].State2 &&
			this.ScannerChaser {

			this.Functions[this.SelectedSequence][common.Function2_Auto_Color].State2 = true
			cmd := common.Command{
				Action: common.UpdateAutoColor,
				Args: []common.Arg{
					{Name: "AutoColor", Value: true},
					{Name: "Type", Value: this.SelectedType},
				},
			}
			common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)
			ShowFunctionButtons(this, this.SelectedSequence, eventsForLaunchpad, guiButtons)
			return
		}
		if X == common.Function2_Auto_Color &&
			this.Functions[this.SelectedSequence][common.Function2_Auto_Color].State2 &&
			this.ScannerChaser {

			this.Functions[this.SelectedSequence][common.Function2_Auto_Color].State2 = false
			cmd := common.Command{
				Action: common.UpdateAutoColor,
				Args: []common.Arg{
					{Name: "AutoColor", Value: false},
					{Name: "Type", Value: this.SelectedType},
				},
			}
			common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)
			ShowFunctionButtons(this, this.SelectedSequence, eventsForLaunchpad, guiButtons)
			return
		}

		// Function 3 Set Auto Pattern - Toggle Auto Pattern for scanner and rgb sequences.
		if X == common.Function3_Auto_Pattern &&
			!this.Functions[this.SelectedSequence][common.Function3_Auto_Pattern].State &&
			!this.ScannerChaser {

			this.Functions[this.SelectedSequence][common.Function3_Auto_Pattern].State = true

			cmd := common.Command{
				Action: common.UpdateAutoPattern,
				Args: []common.Arg{
					{Name: "AutoPattern", Value: true},
				},
			}
			common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)
			ShowFunctionButtons(this, this.SelectedSequence, eventsForLaunchpad, guiButtons)
			return
		}
		if X == common.Function3_Auto_Pattern &&
			this.Functions[this.SelectedSequence][common.Function3_Auto_Pattern].State &&
			!this.ScannerChaser {

			this.Functions[this.SelectedSequence][common.Function3_Auto_Pattern].State = false
			cmd := common.Command{
				Action: common.UpdateAutoPattern,
				Args: []common.Arg{
					{Name: "AutoPattern", Value: false},
				},
			}
			common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)
			ShowFunctionButtons(this, this.SelectedSequence, eventsForLaunchpad, guiButtons)
			return
		}

		// Function 3 Set Auto Pattern - Toggle Auto Pattern for chaser sequences
		if X == common.Function3_Auto_Pattern &&
			!this.Functions[this.SelectedSequence][common.Function3_Auto_Pattern].State2 &&
			this.ScannerChaser {

			this.Functions[this.SelectedSequence][common.Function3_Auto_Pattern].State2 = true

			cmd := common.Command{
				Action: common.UpdateAutoPattern,
				Args: []common.Arg{
					{Name: "AutoPattern", Value: true},
				},
			}
			common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)
			ShowFunctionButtons(this, this.SelectedSequence, eventsForLaunchpad, guiButtons)
			return
		}
		if X == common.Function3_Auto_Pattern &&
			this.Functions[this.SelectedSequence][common.Function3_Auto_Pattern].State2 &&
			this.ScannerChaser {

			this.Functions[this.SelectedSequence][common.Function3_Auto_Pattern].State2 = false

			cmd := common.Command{
				Action: common.UpdateAutoPattern,
				Args: []common.Arg{
					{Name: "AutoPattern", Value: false},
				},
			}
			common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)
			ShowFunctionButtons(this, this.SelectedSequence, eventsForLaunchpad, guiButtons)
			return
		}

		// Function 4 Bounce - Toggle bounce feature.
		if X == common.Function4_Bounce {

			startCmd := common.Command{
				Action: common.UpdateBounce,
				Args: []common.Arg{
					{Name: "Bounce", Value: true},
				},
			}

			stopCmd := common.Command{
				Action: common.UpdateBounce,
				Args: []common.Arg{
					{Name: "Bounce", Value: false},
				},
			}

			if this.ScannerChaser && this.SelectedType == "scanner" {
				// Chaser Sequence.
				if this.Functions[this.SelectedSequence][common.Function4_Bounce].State2 {
					this.Functions[this.SelectedSequence][common.Function4_Bounce].State2 = false
					common.SendCommandToSequence(this.ChaserSequenceNumber, stopCmd, commandChannels)
				} else {
					this.Functions[this.SelectedSequence][common.Function4_Bounce].State2 = true
					common.SendCommandToSequence(this.ChaserSequenceNumber, startCmd, commandChannels)
				}

			} else {
				// Scanner Sequence.
				if this.Functions[this.SelectedSequence][common.Function4_Bounce].State {
					this.Functions[this.SelectedSequence][common.Function4_Bounce].State = false
					common.SendCommandToSequence(this.SelectedSequence, stopCmd, commandChannels)
				} else {
					this.Functions[this.SelectedSequence][common.Function4_Bounce].State = true
					common.SendCommandToSequence(this.SelectedSequence, startCmd, commandChannels)
				}
			}
			ShowFunctionButtons(this, this.SelectedSequence, eventsForLaunchpad, guiButtons)
			return
		}

		// Map Function 5 RGB - Go straight into RGB color edit mode, don't wait for a another select press.
		if X == common.Function5_Color && !this.Functions[this.SelectedSequence][common.Function5_Color].State &&
			sequences[this.SelectedSequence].Type == "rgb" {
			this.EditSequenceColorsMode[this.SelectedSequence] = true
			this.Functions[this.SelectedSequence][common.Function5_Color].State = true
			time.Sleep(500 * time.Millisecond) // But give the launchpad time to light the function key purple.
			common.ClearSelectedRowOfButtons(this.SelectedSequence, eventsForLaunchpad, guiButtons)
			// Set the colors.
			sequences[this.SelectedSequence].CurrentColors = sequences[this.SelectedSequence].SequenceColors
			// Show the colors
			ShowRGBColorSelectionButtons(this.MasterBrightness, this.SelectedSequence, *sequences[this.SelectedSequence], eventsForLaunchpad, guiButtons)
			return
		}

		// Map Function 5 Scanner Color Selection - Go straight into scanner color edit mode via select fixture, don't wait for a another select press.
		if X == common.Function5_Color && !this.Functions[this.SelectedSequence][common.Function5_Color].State &&
			sequences[this.SelectedSequence].Type == "scanner" {

			this.Functions[this.SelectedSequence][common.Function5_Color].State = true
			this.EditSequenceColorsMode[this.SelectedSequence] = true

			ShowFunctionButtons(this, this.SelectedSequence, eventsForLaunchpad, guiButtons)
			time.Sleep(500 * time.Millisecond) // But give the launchpad time to light the function key purple.

			common.ClearSelectedRowOfButtons(this.SelectedSequence, eventsForLaunchpad, guiButtons)

			// Select a fixture.
			this.EditFixtureSelectionMode = true
			this.SelectMode[this.SelectedSequence] = FUNCTION
			sequences[this.SelectedSequence].StaticColors[X].FirstPress = false

			this.FollowingAction = "ShowScannerColorSelectionButtons"
			this.SelectedFixture = ShowSelectFixtureButtons(*sequences[this.SelectedSequence], this, eventsForLaunchpad, this.FollowingAction, guiButtons)
			return
		}

		// Map Function 6 Scanner GOBO Selection - Go to select gobo mode if we are in scanner sequence.
		if X == common.Function6_Static_Gobo && !this.Functions[this.SelectedSequence][common.Function6_Static_Gobo].State &&
			sequences[this.SelectedSequence].Type == "scanner" {

			this.Functions[this.SelectedSequence][common.Function6_Static_Gobo].State = true
			this.EditStaticColorsMode[this.SelectedSequence] = false // Turn off the other option for this function key.
			this.EditGoboSelectionMode[this.SelectedSequence] = true

			ShowFunctionButtons(this, this.SelectedSequence, eventsForLaunchpad, guiButtons)
			time.Sleep(500 * time.Millisecond) // But give the launchpad time to light the function key purple.

			common.ClearSelectedRowOfButtons(this.SelectedSequence, eventsForLaunchpad, guiButtons)

			// Select a fixture.
			this.EditFixtureSelectionMode = true
			this.SelectMode[this.SelectedSequence] = FUNCTION
			sequences[this.SelectedSequence].StaticColors[X].FirstPress = false

			this.FollowingAction = "ShowGoboSelectionButtons"
			this.SelectedFixture = ShowSelectFixtureButtons(*sequences[this.SelectedSequence], this, eventsForLaunchpad, this.FollowingAction, guiButtons)
			return
		}

		// Function 6 RGB - Turn on edit static color mode.
		if X == common.Function6_Static_Gobo &&
			!this.Functions[this.SelectedSequence][common.Function6_Static_Gobo].State &&
			sequences[this.SelectedSequence].Type == "rgb" {
			this.Functions[this.SelectedSequence][common.Function6_Static_Gobo].State = true
			// Starting a static sequence will turn off a running chaser, so turn off the start lamp
			common.LightLamp(common.ALight{X: X, Y: Y, Brightness: this.MasterBrightness, Red: 255, Green: 255, Blue: 255}, eventsForLaunchpad, guiButtons)
			//  and remember that this sequence is off.
			this.Running[this.SelectedSequence] = false
			this.EditGoboSelectionMode[this.SelectedSequence] = false // Turn off the other option for this function key.
			this.EditStaticColorsMode[this.SelectedSequence] = true   // Turn on edit static color mode.
			this.SelectMode[this.SelectedSequence] = NORMAL           // Turn off functions.
			// Go straight to static color selection mode, don't wait for a another select press.
			ShowFunctionButtons(this, this.SelectedSequence, eventsForLaunchpad, guiButtons)
			time.Sleep(250 * time.Millisecond) // But give the launchpad time to light the function key purple.
			common.ClearLabelsSelectedRowOfButtons(this.SelectedSequence, guiButtons)
			// Switch on any static colors.
			cmd := common.Command{
				Action: common.UpdateStatic,
				Args: []common.Arg{
					{Name: "Static", Value: true},
				},
			}
			common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)
			return
		}

		// Function 6 RGB - Turn off edit static color mode.
		if X == common.Function6_Static_Gobo &&
			this.Functions[this.SelectedSequence][common.Function6_Static_Gobo].State &&
			sequences[this.SelectedSequence].Type == "rgb" {

			if debug {
				fmt.Printf("Function 6 RGB - Turn off edit static color mode.\n")

			}

			this.Functions[this.SelectedSequence][common.Function6_Static_Gobo].State = false

			this.EditGoboSelectionMode[this.SelectedSequence] = false // Turn off the other option for this function key.
			this.EditStaticColorsMode[this.SelectedSequence] = false  // Turn off edit static color mode.
			this.SelectMode[this.SelectedSequence] = NORMAL           // Turn off function selection mode.
			// Go straight to static color selection mode, don't wait for a another select press.
			ShowFunctionButtons(this, this.SelectedSequence, eventsForLaunchpad, guiButtons)
			time.Sleep(250 * time.Millisecond) // But give the launchpad time to light the function key purple.
			common.ClearLabelsSelectedRowOfButtons(this.SelectedSequence, guiButtons)
			// Switch off static colors.
			cmd := common.Command{
				Action: common.UpdateStatic,
				Args: []common.Arg{
					{Name: "Static", Value: false},
				},
			}
			common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)
			common.RevealSequence(this.SelectedSequence, commandChannels)
			return
		}

		// Function 7 - Turn on the RGB Invert mode.
		if X == common.Function7_Invert_Chase &&
			!this.Functions[this.SelectedSequence][common.Function7_Invert_Chase].State &&
			sequences[this.SelectedSequence].Type == "rgb" {

			this.Functions[this.SelectedSequence][common.Function7_Invert_Chase].State = true

			cmd := common.Command{
				Action: common.UpdateRGBInvert,
				Args: []common.Arg{
					{Name: "RGBInvert", Value: true},
				},
			}
			common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

			ShowFunctionButtons(this, this.SelectedSequence, eventsForLaunchpad, guiButtons)
			return
		}

		// Function 7 - Turn off the RGB Invert mode.
		if X == common.Function7_Invert_Chase &&
			this.Functions[this.SelectedSequence][common.Function7_Invert_Chase].State &&
			sequences[this.SelectedSequence].Type == "rgb" {

			this.Functions[this.SelectedSequence][common.Function7_Invert_Chase].State = false

			cmd := common.Command{
				Action: common.UpdateRGBInvert,
				Args: []common.Arg{
					{Name: "RGBInvert", Value: false},
				},
			}
			common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

			ShowFunctionButtons(this, this.SelectedSequence, eventsForLaunchpad, guiButtons)
			return
		}

		// Function 7 - Toggle the shutter chaser mode. Start the chaser.
		if X == common.Function7_Invert_Chase &&
			!this.Functions[this.SelectedSequence][common.Function7_Invert_Chase].State2 &&
			sequences[this.SelectedSequence].Type == "scanner" {

			this.ScannerChaser = true
			this.Functions[this.SelectedSequence][common.Function7_Invert_Chase].State2 = true // Chaser

			ShowFunctionButtons(this, this.SelectedSequence, eventsForLaunchpad, guiButtons)
			time.Sleep(500 * time.Millisecond) // But give the launchpad time to light the function key purple.

			// Tell the scannern & chaser sequences that the scanner shutter chase flag is on.
			cmd := common.Command{
				Action: common.UpdateScannerHasShutterChase,
				Args: []common.Arg{
					{Name: "ScannerHasShutterChase", Value: this.ScannerChaser},
				},
			}
			common.SendCommandToSequence(this.ScannerSequenceNumber, cmd, commandChannels)
			common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)

			// Tell the chaser to start.
			cmd = common.Command{
				Action: common.StartChase,
			}
			common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)

			// Update the buttons: speed
			common.LabelButton(0, 7, "Chase\nSpeed\nDown", guiButtons)
			common.LabelButton(1, 7, "Chase\nSpeed\nUp", guiButtons)

			common.LabelButton(2, 7, "Chase\nShift\nDown", guiButtons)
			common.LabelButton(3, 7, "Chase\nShift\nUp", guiButtons)

			common.LabelButton(4, 7, "Chase\nSize\nDown", guiButtons)
			common.LabelButton(5, 7, "Chase\nSize\nUp", guiButtons)

			common.LabelButton(6, 7, "Chase\nFase\nSoft", guiButtons)
			common.LabelButton(7, 7, "Chase\nFade\nSharp", guiButtons)

			HandleSelect(sequences, this, eventsForLaunchpad, commandChannels, guiButtons)
			return
		}

		// Function 7 - Toggle the shutter chaser mode. Stop the chaser.
		if X == common.Function7_Invert_Chase &&
			this.Functions[this.SelectedSequence][common.Function7_Invert_Chase].State2 &&
			sequences[this.SelectedSequence].Type == "scanner" {

			this.ScannerChaser = false
			this.Functions[this.SelectedSequence][common.Function7_Invert_Chase].State2 = false

			ShowFunctionButtons(this, this.SelectedSequence, eventsForLaunchpad, guiButtons)
			time.Sleep(500 * time.Millisecond) // But give the launchpad time to light the function key purple.

			// Tell scanner & chaser sequence that the scanner shutter chase flag is off.
			this.Running[this.ChaserSequenceNumber] = false
			cmd := common.Command{
				Action: common.UpdateScannerHasShutterChase,
				Args: []common.Arg{
					{Name: "ScannerHasShutterChase", Value: this.ScannerChaser},
				},
			}
			common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)
			common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)

			// Stop the chaser.
			cmd = common.Command{
				Action: common.StopChase,
			}
			common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)

			// Update the buttons: speed
			common.LabelButton(0, 7, "Speed\nDown", guiButtons)
			common.LabelButton(1, 7, "Speed\nUp", guiButtons)

			common.LabelButton(2, 7, "Shift\nDown", guiButtons)
			common.LabelButton(3, 7, "Shift\nUp", guiButtons)

			common.LabelButton(4, 7, "Size\nDown", guiButtons)
			common.LabelButton(5, 7, "Size\nUp", guiButtons)

			common.LabelButton(6, 7, "Fase\nSoft", guiButtons)
			common.LabelButton(7, 7, "Fade\nSharp", guiButtons)

			HandleSelect(sequences, this, eventsForLaunchpad, commandChannels, guiButtons)
			return
		}

		// Function 8 MUSIC TRIGGER  - Send start music trigger for scanner & rgb sequences.
		if X == common.Function8_Music_Trigger &&
			!this.ScannerChaser &&
			!this.Functions[this.SelectedSequence][common.Function8_Music_Trigger].State {

			this.Functions[this.SelectedSequence][common.Function8_Music_Trigger].State = true

			ShowFunctionButtons(this, this.SelectedSequence, eventsForLaunchpad, guiButtons)
			time.Sleep(250 * time.Millisecond) // But give the launchpad time to light the function key purple.

			// Starting a music trigger will start the sequence, so turn on the start lamp
			// and remember that this sequence is on.
			this.Running[this.SelectedSequence] = true
			common.ShowRunningStatus(this.SelectedSequence, this.Running, eventsForLaunchpad, guiButtons)

			// Start the music trigger for the target sequence.
			cmd := common.Command{
				Action: common.UpdateMusicTrigger,
				Args: []common.Arg{
					{Name: "MusicTriger", Value: true},
				},
			}
			common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

			// We want to exit from functioms immediately so we call handle.
			HandleSelect(sequences, this, eventsForLaunchpad, commandChannels, guiButtons)
			return
		}

		// Function 8 MUSIC TRIGGER  - Send start music trigger for chaser sequences.
		if X == common.Function8_Music_Trigger &&
			this.ScannerChaser &&
			!this.Functions[this.SelectedSequence][common.Function8_Music_Trigger].State2 {

			this.Functions[this.SelectedSequence][common.Function8_Music_Trigger].State2 = true
			this.ScannerChaser = true

			ShowFunctionButtons(this, this.SelectedSequence, eventsForLaunchpad, guiButtons)
			time.Sleep(500 * time.Millisecond) // But give the launchpad time to light the function key purple.

			// Tell the chaser sequences that the music trigger is on.
			cmd := common.Command{
				Action: common.UpdateMusicTrigger,
				Args: []common.Arg{
					{Name: "MusicTriger", Value: this.ScannerChaser},
				},
			}
			common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)

			// We want to exit from functioms immediately so we call handle.
			HandleSelect(sequences, this, eventsForLaunchpad, commandChannels, guiButtons)
			return
		}

		// Function 8 MUSIC TRIGGER  - Send stop music trigger for scanner and rgb sequences.
		if X == common.Function8_Music_Trigger &&

			this.Functions[this.SelectedSequence][common.Function8_Music_Trigger].State {

			this.Functions[this.SelectedSequence][common.Function8_Music_Trigger].State = false

			this.Running[this.SelectedSequence] = false
			common.ShowRunningStatus(this.SelectedSequence, this.Running, eventsForLaunchpad, guiButtons)

			// Stop the music trigger for the target sequence.
			cmd := common.Command{
				Action: common.UpdateMusicTrigger,
				Args: []common.Arg{
					{Name: "MusicTriger", Value: false},
				},
			}
			common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

			// We want to exit from functioms immediately so we call handle.
			HandleSelect(sequences, this, eventsForLaunchpad, commandChannels, guiButtons)
			return
		}

		// Function 8 MUSIC TRIGGER  - Send stop music trigger chaser sequences.
		if X == common.Function8_Music_Trigger &&
			this.ScannerChaser &&
			this.Functions[this.SelectedSequence][common.Function8_Music_Trigger].State2 {

			this.Functions[this.SelectedSequence][common.Function8_Music_Trigger].State2 = false
			this.ScannerChaser = false

			ShowFunctionButtons(this, this.SelectedSequence, eventsForLaunchpad, guiButtons)
			time.Sleep(500 * time.Millisecond) // But give the launchpad time to light the function key purple.

			// Tell scanner & chaser sequence that the scanner shutter chase flag is off.
			this.Running[this.ChaserSequenceNumber] = this.ScannerChaser
			cmd := common.Command{
				Action: common.UpdateScannerHasShutterChase,
				Args: []common.Arg{
					{Name: "ScannerHasShutterChase", Value: this.ScannerChaser},
				},
			}
			common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)
			common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)

			// Stop the music trigger for the chaser sequence.
			cmd = common.Command{
				Action: common.UpdateMusicTrigger,
				Args: []common.Arg{
					{Name: "MusicTriger", Value: this.ScannerChaser},
				},
			}
			common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)

			// We want to exit from functioms immediately so we call handle.
			HandleSelect(sequences, this, eventsForLaunchpad, commandChannels, guiButtons)
			return
		}
	}

	// B L A C K O U T   B U T T O N.
	if X == 8 && Y == 7 {

		if debug {
			fmt.Printf("BLACKOUT\n")
		}

		// Turn off the flashing save button
		this.SavePreset = false
		common.LightLamp(common.ALight{X: 8, Y: 4, Brightness: this.MasterBrightness, Red: 255, Green: 255, Blue: 255}, eventsForLaunchpad, guiButtons)

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
		fmt.Printf("HANDLE: this.Type = %s \n", this.SelectedType)
		for functionNumber := 0; functionNumber < 8; functionNumber++ {
			state := this.Functions[this.SelectedSequence][functionNumber].State
			fmt.Printf("HANDLE: function %d state %t\n", functionNumber, state)
		}
		fmt.Printf("HANDLE: this.ChaserRunning %t \n", this.ScannerChaser)
		fmt.Printf("================== WHAT MODE =================\n")
		fmt.Printf("HANDLE: this.SelectButtonPressed[%d] = %t \n", this.SelectedSequence, this.SelectButtonPressed[this.SelectedSequence])
		if this.SelectMode[this.SelectedSequence] == NORMAL {
			fmt.Printf("HANDLE: this.SelectMode[%d] = NORMAL \n", this.SelectedSequence)
		}
		if this.SelectMode[this.SelectedSequence] == FUNCTION {
			fmt.Printf("HANDLE: this.SelectMode[%d] = FUNCTION \n", this.SelectedSequence)
		}
		if this.SelectMode[this.SelectedSequence] == STATUS {
			fmt.Printf("HANDLE: this.SelectMode[%d] = STATUS \n", this.SelectedSequence)
		}

		fmt.Printf("HANDLE: this.EditSequenceColorsMode[%d] = %t \n", this.SelectedSequence, this.EditSequenceColorsMode[this.SelectedSequence])
		fmt.Printf("HANDLE: this.EditStaticColorsMode[%d] = %t \n", this.SelectedSequence, this.EditStaticColorsMode[this.SelectedSequence])
		fmt.Printf("HANDLE: this.EditGoboSelectionMode[%d] = %t \n", this.SelectedSequence, this.EditGoboSelectionMode[this.SelectedSequence])
		fmt.Printf("HANDLE: this.EditPatternMode[%d] = %t \n", this.SelectedSequence, this.EditPatternMode[this.SelectedSequence])
		fmt.Printf("===============================================\n")
	}

	//        +-------------------+
	//        |       NORMAL      |
	//        +-------------------+
	//                 |
	//                 V
	//        +-------------------+
	//        |     FUNCTION      |
	//        +-------------------+
	//            |            | If Chaser Enabled
	//            V            V
	//            |       +-------------------+
	//            |       |  CHASER FUNCTIONS |
	//            |       +-------------------+
	//            |				|
	//            V				V
	//         +-------------------+
	//         |       STATUS      |
	//         +-------------------+
	//                  |
	//                  V
	//         +-------------------+
	//         |       NORMAL      |
	//         +-------------------+

	// Update the status bar
	if this.Strobe[this.SelectedSequence] {
		common.UpdateStatusBar(fmt.Sprintf("Strobe %02d", this.StrobeSpeed[this.SelectedSequence]), "speed", false, guiButtons)
	} else {
		// Update status bar.
		if this.Functions[this.SelectedSequence][common.Function8_Music_Trigger].State ||
			this.Functions[this.SelectedSequence][common.Function8_Music_Trigger].State2 {
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

	// Light the start stop button.
	common.ShowRunningStatus(this.SelectedSequence, this.Running, eventsForLaunchpad, guiButtons)

	// 1st Press Select Sequence - This the first time we have pressed the select button.
	// Simply select the selected sequence.
	// But remember we have pressed this select button once.
	if this.SelectMode[this.SelectedSequence] == NORMAL &&
		!this.SelectButtonPressed[this.SelectedSequence] &&
		!this.EditStaticColorsMode[this.SelectedSequence] {

		if debug {
			fmt.Printf("%d: Show Sequence - Handle 2\n", this.SelectedSequence)
		}
		// Assume everything else is off.
		this.SelectButtonPressed[0] = false
		this.SelectButtonPressed[1] = false
		this.SelectButtonPressed[2] = false
		this.SelectButtonPressed[3] = false

		// But remember we have pressed this select button once.
		this.SelectMode[this.SelectedSequence] = NORMAL
		this.SelectButtonPressed[this.SelectedSequence] = true

		// Turn off any previous function or status bars.
		for sequenceNumber := range sequences {
			if this.SelectMode[sequenceNumber] == FUNCTION ||
				this.SelectMode[sequenceNumber] == STATUS {
				common.ClearSelectedRowOfButtons(sequenceNumber, eventsForLaunchpad, guiButtons)
				// And reveal all the other sequence that isn't us.
				if sequenceNumber != this.SelectedSequence {
					common.RevealSequence(sequenceNumber, commandChannels)
					// And turn off the function selected.
					this.SelectMode[sequenceNumber] = NORMAL
				}
			}
		}

		if this.Functions[this.SelectedSequence][common.Function1_Pattern].State {
			// Reset the pattern function key.
			this.Functions[this.SelectedSequence][common.Function1_Pattern].State = false

			// Clear the pattern function keys
			ClearPatternSelectionButtons(this.SelectedSequence, *sequences[this.SelectedSequence], eventsForLaunchpad, guiButtons)

			// And reveal the sequence.
			common.RevealSequence(this.SelectedSequence, commandChannels)

			// If the chase is running, reveal it.
			if this.ScannerChaser && this.SelectedType == "scanner" {
				common.RevealSequence(this.ChaserSequenceNumber, commandChannels)
			}

			// Editing pattern is over for this sequence.
			this.EditPatternMode[this.SelectedSequence] = false

			// Clear buttons and remove any labels.
			common.ClearSelectedRowOfButtons(this.SelectedSequence, eventsForLaunchpad, guiButtons)
		}

		if this.SelectMode[this.SelectedSequence] == NORMAL &&
			this.Functions[this.SelectedSequence][common.Function5_Color].State && this.EditSequenceColorsMode[this.SelectedSequence] {
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

	// 2nd Press same select button go into Function Mode for this sequence.
	if this.SelectMode[this.SelectedSequence] == NORMAL && sequences[this.SelectedSequence].Type != "switch" || // Don't alow functions in switch mode.
		this.SelectMode[this.SelectedSequence] == NORMAL && // Function select mode is off
			this.EditStaticColorsMode[this.SelectedSequence] && // AND static colol mode, the case when we leave static colors edit mode.
			this.SelectButtonPressed[this.SelectedSequence] {

		if debug {
			fmt.Printf("%d: 2nd Press Function Bar Mode - Handle 4\n", this.SelectedSequence)
		}

		// Set function mode.
		this.SelectMode[this.SelectedSequence] = FUNCTION

		// If static, show static colors.
		if this.EditStaticColorsMode[this.SelectedSequence] && sequences[this.SelectedSequence].Type != "scanner" {
			if debug {
				fmt.Printf("Show Static Color Selection Buttons\n")
			}
			common.SetMode(this.SelectedSequence, commandChannels, "Static")
			//this.EditStaticColorsMode[this.SelectedSequence] = false
		}

		// And hide the sequence so we can only see the function buttons.
		common.HideSequence(this.SelectedSequence, commandChannels)

		// If the chase is running, hide it.
		if this.ScannerChaser && this.SelectedType == "scanner" {
			common.HideSequence(this.ChaserSequenceNumber, commandChannels)
		}

		// Turn off any static sequence so we can see the functions.
		common.SetMode(this.SelectedSequence, commandChannels, "Sequence")

		// Turn off any previous function bars.
		for sequenceNumber := range sequences {
			if this.SelectMode[sequenceNumber] == FUNCTION {
				common.ClearSelectedRowOfButtons(sequenceNumber, eventsForLaunchpad, guiButtons)
				// And reveal all the other sequence that isn't us.
				if sequenceNumber != this.SelectedSequence {
					common.RevealSequence(sequenceNumber, commandChannels)
					// And turn off the function selected.
					this.SelectMode[sequenceNumber] = NORMAL
				}
			}
		}

		// Clear the buttons.
		common.ClearSelectedRowOfButtons(this.SelectedSequence, eventsForLaunchpad, guiButtons)

		// Show the function buttons.
		ShowFunctionButtons(this, this.SelectedSequence, eventsForLaunchpad, guiButtons)

		// Now forget we pressed twice and start again.
		this.SelectButtonPressed[this.SelectedSequence] = false

		return
	}

	// 3rd Press Status Mode - we display the fixture status enable/invert/disable buttons.
	if this.SelectMode[this.SelectedSequence] == FUNCTION &&
		!this.SelectButtonPressed[this.SelectedSequence] &&
		!this.EditSequenceColorsMode[this.SelectedSequence] &&
		!this.EditStaticColorsMode[this.SelectedSequence] {

		if debug {
			fmt.Printf("%d: Handle 5 Status Mode, Function Bar off, status buttons on\n", this.SelectedSequence)
		}

		// Turn on status mode
		this.SelectMode[this.SelectedSequence] = STATUS

		// If the chase is running, hide it.
		if this.ScannerChaser && this.SelectedType == "scanner" {
			common.HideSequence(this.ChaserSequenceNumber, commandChannels)
		}

		// Remove the function pads.
		common.ClearSelectedRowOfButtons(this.SelectedSequence, eventsForLaunchpad, guiButtons)

		ShowFixtureStatus(this.SelectedSequence, *sequences[this.SelectedSequence], this, eventsForLaunchpad, guiButtons, commandChannels)

		return
	}

	// 4th Press Normal Mode - we head back to normal mode.
	if this.SelectMode[this.SelectedSequence] == STATUS &&
		!this.SelectButtonPressed[this.SelectedSequence] &&
		!this.EditSequenceColorsMode[this.SelectedSequence] &&
		!this.EditStaticColorsMode[this.SelectedSequence] {

		if debug {
			fmt.Printf("%d: Handle 1 Normal Mode, Function Bar off\n", this.SelectedSequence)
		}

		// Turn off function mode.
		this.SelectMode[this.SelectedSequence] = NORMAL

		// Remove the status buttons.
		common.ClearSelectedRowOfButtons(this.SelectedSequence, eventsForLaunchpad, guiButtons)

		// We're in Edit Pattern Mode.
		if this.Functions[this.SelectedSequence][common.Function1_Pattern].State {
			if debug {
				fmt.Printf("Show Pattern Selection Buttons\n")
			}
			this.EditPatternMode[this.SelectedSequence] = true
			common.HideSequence(this.SelectedSequence, commandChannels)
			ShowPatternSelectionButtons(this, sequences[this.SelectedSequence].Master, *sequences[this.SelectedSequence], eventsForLaunchpad, guiButtons)
			return
		}

		// We're in RGB Color Selection Mode.
		if this.Functions[this.SelectedSequence][common.Function5_Color].State && sequences[this.SelectedSequence].Type == "rgb" {
			if debug {
				fmt.Printf("Show RGB Sequence Color Selection Buttons\n")
			}
			// Set the colors.
			sequences[this.SelectedSequence].CurrentColors = sequences[this.SelectedSequence].SequenceColors
			// Show the colors
			ShowRGBColorSelectionButtons(this.MasterBrightness, this.SelectedSequence, *sequences[this.SelectedSequence], eventsForLaunchpad, guiButtons)
			return
		}

		// We're in RGB Static Color Mode.
		if this.Functions[this.SelectedSequence][common.Function6_Static_Gobo].State && sequences[this.SelectedSequence].Type == "rgb" {
			if debug {
				fmt.Printf("Show RGB Static Colors\n")
			}
			this.EditStaticColorsMode[this.SelectedSequence] = true

			// Tell the sequence about the new color and where we are in the
			// color cycle.
			cmd := common.Command{
				Action: common.UpdateStaticColor,
				Args: []common.Arg{
					{Name: "Static", Value: true},
					{Name: "StaticLamp", Value: this.LastStaticColorButtonX},
					{Name: "StaticLampFlash", Value: true},
					{Name: "SelectedColor", Value: sequences[this.SelectedSequence].StaticColors[this.LastStaticColorButtonX].SelectedColor},
					{Name: "StaticColor", Value: sequences[this.SelectedSequence].StaticColors[this.LastStaticColorButtonX].Color},
				},
			}
			common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

			return
		}

		// We're in Scanner Gobo Selection Mode.
		if this.Functions[this.SelectedSequence][common.Function6_Static_Gobo].State && sequences[this.SelectedSequence].Type == "scanner" {
			this.Functions[this.SelectedSequence][common.Function6_Static_Gobo].State = false
			this.EditGoboSelectionMode[this.SelectedSequence] = false
		}

		// Allow us to exit the pattern select mode without setting a pattern.
		if this.EditPatternMode[this.SelectedSequence] {
			this.EditPatternMode[this.SelectedSequence] = false
		}

		// Else reveal the sequence on the launchpad keys
		if debug {
			fmt.Printf("%d: Reveal Sequence\n", this.SelectedSequence)
		}
		common.RevealSequence(this.SelectedSequence, commandChannels)

		// If the chase is running, reveal it.
		if this.ScannerChaser && this.SelectedType == "scanner" {
			if debug {
				fmt.Printf("%d: Reveal Sequence\n", this.ChaserSequenceNumber)
			}
			common.RevealSequence(this.ChaserSequenceNumber, commandChannels)
		}

		// Turn off the function mode flag.
		this.SelectMode[this.SelectedSequence] = NORMAL
		// Now forget we pressed twice and start again.
		this.SelectButtonPressed[this.SelectedSequence] = true

		return
	}

	// Are we in function mode ?
	if this.SelectMode[this.SelectedSequence] == FUNCTION {
		if debug {
			fmt.Printf("%d: Handle 3\n", this.SelectedSequence)
		}
		// Turn off function mode. Remove the function pads.
		common.ClearSelectedRowOfButtons(this.SelectedSequence, eventsForLaunchpad, guiButtons)

		// And reveal the sequence on the launchpad keys
		common.RevealSequence(this.SelectedSequence, commandChannels)

		// If the chaser is running, reveal it.
		if this.ScannerChaser && this.SelectedType == "scanner" {
			common.RevealSequence(this.ChaserSequenceNumber, commandChannels)
		}

		// Turn off the function mode flag.
		this.SelectMode[this.SelectedSequence] = NORMAL

		// Turn off the edit sequence colors button.
		if this.EditSequenceColorsMode[this.SelectedSequence] {
			this.EditSequenceColorsMode[this.SelectedSequence] = false
			this.Functions[this.SelectedSequence][common.Function5_Color].State = false
		}

		// Now forget we pressed twice and start again.
		this.SelectButtonPressed[this.SelectedSequence] = false

		return
	}

}

func unSetEditSequenceColorsMode(sequences []*common.Sequence, this *CurrentState, commandChannels []chan common.Command,
	eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight) {

	// Turn off the edit colors bar.
	this.Functions[this.SelectedSequence][common.Function5_Color].State = false

	// Restart the sequence.
	// cmd := common.Command{
	// 	Action: common.Start,
	// }
	// common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

	// And reveal the sequence on the launchpad keys
	common.RevealSequence(this.SelectedSequence, commandChannels)
	// Turn off the function mode flag.
	this.SelectMode[this.SelectedSequence] = NORMAL
	// Now forget we pressed twice and start again.
	this.SelectButtonPressed[this.SelectedSequence] = true

	common.HideColorSelectionButtons(this.SelectedSequence, *sequences[this.SelectedSequence], eventsForLaunchpad, guiButtons)
}

func AllFixturesOff(sequences []*common.Sequence, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, dmxController *ft232.DMXController, fixturesConfig *fixture.Fixtures, dmxInterfacePresent bool) {
	for y := 0; y < len(sequences); y++ {
		if sequences[y].Type != "switch" && sequences[y].Label != "chaser" {
			for x := 0; x < 8; x++ {
				common.LightLamp(common.ALight{X: x, Y: y, Brightness: 0, Red: 0, Green: 0, Blue: 0}, eventsForLaunchpad, guiButtons)
				fixture.MapFixtures(false, false, y, dmxController, x, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, true, 0, 0, false, 0, dmxInterfacePresent)
				common.LabelButton(x, y, "", guiButtons)
			}
		}
	}
}
func AllRGBFixturesOff(sequences []*common.Sequence, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, dmxController *ft232.DMXController, fixturesConfig *fixture.Fixtures, dmxInterfacePresent bool) {
	for x := 0; x < 8; x++ {
		for sequenceNumber := 0; sequenceNumber < len(sequences); sequenceNumber++ {
			if sequences[sequenceNumber].Type == "rgb" && sequences[sequenceNumber].Label != "chaser" {
				common.LightLamp(common.ALight{X: x, Y: sequenceNumber, Brightness: 0, Red: 0, Green: 0, Blue: 0}, eventsForLaunchpad, guiButtons)
				fixture.MapFixtures(false, false, sequenceNumber, dmxController, x, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, true, 0, 0, false, 0, dmxInterfacePresent)
				common.LabelButton(x, sequenceNumber, "", guiButtons)
			}
		}
	}
}

// Show Scanner status - Dim White is disabled, White is enabled.
func ShowFixtureStatus(selectedSequence int, sequence common.Sequence, this *CurrentState, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command) {

	if debug {
		fmt.Printf("Show Scanner Status for sequence %d number of scanners %d\n", sequence.Number, sequence.NumberFixtures)
	}

	common.HideSequence(selectedSequence, commandChannels)

	for scannerNumber := 0; scannerNumber < sequence.NumberFixtures; scannerNumber++ {

		if debug {
			fmt.Printf("Scanner %d Enabled %t Inverted %t\n", scannerNumber, this.FixtureState[scannerNumber][sequence.Number].Enabled, this.FixtureState[scannerNumber][sequence.Number].Inverted)
		}

		// Enabled but not inverted then On and green.
		if this.FixtureState[scannerNumber][sequence.Number].Enabled && !this.FixtureState[scannerNumber][sequence.Number].Inverted {
			common.LightLamp(common.ALight{X: scannerNumber, Y: sequence.Number, Brightness: this.MasterBrightness, Red: 0, Green: 255, Blue: 0}, eventsForLaunchpad, guiButtons)
			common.LabelButton(scannerNumber, sequence.Number, "On", guiButtons)
		}

		// Enabled and inverted then Invert and red.
		if this.FixtureState[scannerNumber][sequence.Number].Enabled && this.FixtureState[scannerNumber][sequence.Number].Inverted {
			common.LightLamp(common.ALight{X: scannerNumber, Y: sequence.Number, Brightness: this.MasterBrightness, Red: 255, Green: 0, Blue: 0}, eventsForLaunchpad, guiButtons)
			common.LabelButton(scannerNumber, sequence.Number, "Invert", guiButtons)
		}

		// Not enabled and not inverted then off and blue.
		if !this.FixtureState[scannerNumber][sequence.Number].Enabled && !this.FixtureState[scannerNumber][sequence.Number].Inverted {
			common.LightLamp(common.ALight{X: scannerNumber, Y: sequence.Number, Brightness: this.MasterBrightness, Red: 0, Green: 100, Blue: 150}, eventsForLaunchpad, guiButtons)
			common.LabelButton(scannerNumber, sequence.Number, "Off", guiButtons)
		}

	}
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
// mySequenceDisplayNumber is the sequence whos buttons you want the pattern selection to show on.
// master is the master brightness for the same buttons.
// targetSequence - is the squence you are updating the pattern, this could be different in the case
// of scanner shutter chaser sequence which doesn't have it's own buttons.
func ShowPatternSelectionButtons(this *CurrentState, master int, targetSequence common.Sequence, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight) {

	if debug {
		fmt.Printf("Sequence Name %s Type %s  Label %s\n", targetSequence.Name, targetSequence.Type, targetSequence.Label)
		for _, pattern := range this.RGBPatterns {
			fmt.Printf("Found a pattern called %s\n", pattern.Name)
		}
	}

	LightBlue := common.Color{R: 0, G: 100, B: 255}
	White := common.Color{R: 255, G: 255, B: 255}

	if targetSequence.Type == "rgb" {
		for _, pattern := range this.RGBPatterns {
			if debug {
				fmt.Printf("pattern is %s\n", pattern.Name)
			}
			if pattern.Number == targetSequence.SelectedPattern {
				common.FlashLight(pattern.Number, this.SelectedSequence, White, LightBlue, eventsForLaunchpad, guiButtons)
			} else {
				common.LightLamp(common.ALight{X: pattern.Number, Y: this.SelectedSequence, Red: 0, Green: 100, Blue: 255, Brightness: master}, eventsForLaunchpad, guiButtons)
			}
			common.LabelButton(pattern.Number, this.SelectedSequence, pattern.Label, guiButtons)
		}
		return
	}

	if targetSequence.Type == "scanner" {
		for _, pattern := range targetSequence.ScannerAvailablePatterns {
			if pattern.Number == targetSequence.SelectedPattern {
				common.FlashLight(pattern.Number, this.SelectedSequence, White, LightBlue, eventsForLaunchpad, guiButtons)
			} else {
				common.LightLamp(common.ALight{X: pattern.Number, Y: this.SelectedSequence, Red: 0, Green: 100, Blue: 255, Brightness: master}, eventsForLaunchpad, guiButtons)
			}
			common.LabelButton(pattern.Number, this.SelectedSequence, pattern.Label, guiButtons)
		}
		return
	}
}

func InitButtons(this *CurrentState, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight) {

	// Light the logo blue.
	if this.LaunchPadConnected {
		this.Pad.Light(8, -1, 0, 0, 255)
	}

	// Initially set the Flood, Save, Start, Stop and Blackout buttons to white.
	common.LightLamp(common.ALight{X: 8, Y: 3, Red: 255, Green: 255, Blue: 255, Brightness: 255}, eventsForLaunchpad, guiButtons)
	common.LightLamp(common.ALight{X: 8, Y: 4, Red: 255, Green: 255, Blue: 255, Brightness: 255}, eventsForLaunchpad, guiButtons)
	common.LightLamp(common.ALight{X: 8, Y: 5, Red: 255, Green: 255, Blue: 255, Brightness: 255}, eventsForLaunchpad, guiButtons)
	common.LightLamp(common.ALight{X: 8, Y: 6, Red: 255, Green: 255, Blue: 255, Brightness: 255}, eventsForLaunchpad, guiButtons)
	common.LightLamp(common.ALight{X: 8, Y: 7, Red: 255, Green: 255, Blue: 255, Brightness: 255}, eventsForLaunchpad, guiButtons)

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
		this.PresetsStore[*this.LastPreset] = presets.Preset{State: current.State, Selected: false, Label: current.Label, ButtonColor: current.ButtonColor}
	}
	current := this.PresetsStore[fmt.Sprint(X)+","+fmt.Sprint(Y)]
	this.PresetsStore[fmt.Sprint(X)+","+fmt.Sprint(Y)] = presets.Preset{State: current.State, Selected: true, Label: current.Label, ButtonColor: current.ButtonColor}
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

		// Restore the functions states from the sequence.
		if sequence.Type == "rgb" {
			this.Functions[sequenceNumber][common.Function2_Auto_Color].State = sequences[sequenceNumber].AutoColor
			this.Functions[sequenceNumber][common.Function3_Auto_Pattern].State = sequences[sequenceNumber].AutoPattern
			this.Functions[sequenceNumber][common.Function4_Bounce].State = sequences[sequenceNumber].Bounce
			this.Functions[sequenceNumber][common.Function6_Static_Gobo].State = sequences[sequenceNumber].Static
			this.Functions[sequenceNumber][common.Function7_Invert_Chase].State = sequences[sequenceNumber].RGBInvert
			this.Functions[sequenceNumber][common.Function8_Music_Trigger].State = sequences[sequenceNumber].MusicTrigger
		}
		if sequence.Type == "scanner" {
			this.Functions[sequenceNumber][common.Function2_Auto_Color].State = sequences[sequenceNumber].AutoColor
			this.Functions[sequenceNumber][common.Function3_Auto_Pattern].State = sequences[sequenceNumber].AutoPattern
			this.Functions[sequenceNumber][common.Function4_Bounce].State = sequences[sequenceNumber].Bounce
			this.Functions[sequenceNumber][common.Function7_Invert_Chase].State2 = sequences[sequenceNumber].ScannerChaser
			this.Functions[sequenceNumber][common.Function8_Music_Trigger].State = sequences[sequenceNumber].MusicTrigger
			this.Functions[sequenceNumber][common.Function8_Music_Trigger].State2 = sequences[this.ChaserSequenceNumber].MusicTrigger
			this.ScannerChaser = sequences[sequenceNumber].ScannerChaser
		}

		// If we are loading a switch sequence, update our local copy of the switch settings.
		if sequences[sequenceNumber].Type == "switch" {
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
		ShowRGBColorSelectionButtons(this.MasterBrightness, this.SelectedSequence, *sequences[this.SelectedSequence], eventsForLaunchpad, guiButtons)

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
			this.Functions[sequenceNumber][common.Function1_Pattern].State2 = false

			this.Functions[sequenceNumber][common.Function2_Auto_Color].State = false
			this.Functions[sequenceNumber][common.Function2_Auto_Color].State2 = false

			this.Functions[sequenceNumber][common.Function3_Auto_Pattern].State = false
			this.Functions[sequenceNumber][common.Function3_Auto_Pattern].State2 = false

			this.Functions[sequenceNumber][common.Function4_Bounce].State = false
			this.Functions[sequenceNumber][common.Function4_Bounce].State2 = false

			this.Functions[sequenceNumber][common.Function5_Color].State = false
			this.Functions[sequenceNumber][common.Function5_Color].State2 = false

			this.Functions[sequenceNumber][common.Function6_Static_Gobo].State = false
			this.Functions[sequenceNumber][common.Function6_Static_Gobo].State2 = false

			this.Functions[sequenceNumber][common.Function7_Invert_Chase].State = false
			this.Functions[sequenceNumber][common.Function7_Invert_Chase].State2 = false

			this.Functions[sequenceNumber][common.Function8_Music_Trigger].State = false
			this.Functions[sequenceNumber][common.Function8_Music_Trigger].State2 = false
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
	for sequenceNumber := range sequences {
		this.SelectButtonPressed[sequenceNumber] = false
		this.SelectMode[sequenceNumber] = NORMAL
		this.EditSequenceColorsMode[sequenceNumber] = false
		this.EditStaticColorsMode[sequenceNumber] = false
		this.EditGoboSelectionMode[sequenceNumber] = false
		this.EditPatternMode[sequenceNumber] = false
		for function := range this.Functions {
			this.Functions[sequenceNumber][function].State = false
		}
	}
}

func ShowFunctionButtons(this *CurrentState, selectedSequence int, eventsForLauchpad chan common.ALight, guiButtons chan common.ALight) {

	if debug {
		fmt.Printf("ShowFunctionButtons sequence %d\n", selectedSequence)
	}

	// Loop through the available functions for this sequence
	for index, function := range this.Functions[selectedSequence] {

		if debug {
			fmt.Printf("ShowFunctionButtons: function %s state %t state2 %t\n", function.Name, function.State, function.State2)
		}

		if !function.State && !function.State2 { // Both off so Cyan.
			common.LightLamp(common.ALight{X: index, Y: selectedSequence, Brightness: 255, Red: 3, Green: 255, Blue: 255}, eventsForLauchpad, guiButtons)
		}

		if function.State && !function.State2 { // Scanner or rgb on and chaser off.
			if !this.ScannerChaser { // Purple
				common.LightLamp(common.ALight{X: index, Y: selectedSequence, Brightness: 255, Red: 200, Green: 0, Blue: 255}, eventsForLauchpad, guiButtons)
			} else { // Cyan
				common.LightLamp(common.ALight{X: index, Y: selectedSequence, Brightness: 255, Red: 3, Green: 255, Blue: 255}, eventsForLauchpad, guiButtons)
			}
		}

		if !function.State && function.State2 { // Scanner or rgb off and chaser on.
			if !this.ScannerChaser { // Cyan
				common.LightLamp(common.ALight{X: index, Y: selectedSequence, Brightness: 255, Red: 3, Green: 255, Blue: 255}, eventsForLauchpad, guiButtons)
			} else { // Purple
				common.LightLamp(common.ALight{X: index, Y: selectedSequence, Brightness: 255, Red: 200, Green: 0, Blue: 255}, eventsForLauchpad, guiButtons)
			}
		}

		if function.State && function.State2 { // Both scanner or rgb and chaser both on.
			common.LightLamp(common.ALight{X: index, Y: selectedSequence, Brightness: 255, Red: 200, Green: 0, Blue: 255}, eventsForLauchpad, guiButtons)
		}

		common.LabelButton(index, selectedSequence, function.Label, guiButtons)

		if this.ScannerChaser && this.SelectedType == "scanner" {
			common.LabelButton(common.Function1_Pattern, this.SelectedSequence, "Chase\nPattern", guiButtons)
			common.LabelButton(common.Function2_Auto_Color, this.SelectedSequence, "Chaser\nAutoColor", guiButtons)
			common.LabelButton(common.Function3_Auto_Pattern, this.SelectedSequence, "Chaser\nAuto\nPattern", guiButtons)
			common.LabelButton(common.Function4_Bounce, this.SelectedSequence, "Chaser\nBounce", guiButtons)
			common.LabelButton(common.Function5_Color, selectedSequence, "Chaser\nColor", guiButtons)
			common.LabelButton(common.Function6_Static_Gobo, selectedSequence, "Chaser\nGobo", guiButtons)
			common.LabelButton(common.Function7_Invert_Chase, this.SelectedSequence, "Chaser", guiButtons)
			common.LabelButton(common.Function8_Music_Trigger, this.SelectedSequence, "Chaser\nMusic\nTrigger", guiButtons)
		}
	}

}
