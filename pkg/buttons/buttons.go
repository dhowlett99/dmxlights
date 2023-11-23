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

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/config"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
	"github.com/dhowlett99/dmxlights/pkg/pad"
	"github.com/dhowlett99/dmxlights/pkg/presets"
	"github.com/dhowlett99/dmxlights/pkg/sound"
	"github.com/oliread/usbdmx"
	"github.com/oliread/usbdmx/ft232"
)

const debug = false

// Select modes.
const (
	NORMAL                int = iota // Normal RGB or Scanner Rotation display.
	NORMAL_STATIC                    // Normal RGB in edit all static fixtures.
	FUNCTION                         // Show the RGB or Scanner functions.
	CHASER_DISPLAY                   //  Show the scanner shutter display.
	CHASER_DISPLAY_STATIC            //  Shutter chaser in edit all fixtures mode.
	CHASER_FUNCTION                  // Show the scammer shutter chaser functions.
	STATUS                           // Show the fixture status states.
)

type CurrentState struct {
	MyWindow                    fyne.Window                // Pointer to main window.
	GUI                         bool                       // Flag to indicate use of GUI.
	Crash1                      bool                       // Flags to detect launchpad crash.
	Crash2                      bool                       // Flags to detect launchpad crash.
	SelectedSequence            int                        // The currently selected sequence.
	TargetSequence              int                        // The current target sequence.
	DisplaySequence             int                        // The current display sequence.
	SequenceType                []string                   // The type, indexed by sequence.
	SelectedStaticFixtureNumber int                        // Temporary storage for the selected fixture number, used by color picker.
	SelectAllStaticFixtures     bool                       // Flag that indicate that all static fixtures have been selected.
	StaticFlashing              []bool                     // Static buttons are flashing, indexed by sequence.
	SavedSequenceColors         map[int][]common.Color     // Local storage for sequence colors.
	SelectedType                string                     // The currently selected sequenece type.
	LastSelectedSequence        int                        // Store fof the last selected squence.
	Speed                       map[int]int                // Local copy of sequence speed. Indexed by sequence.
	RGBShift                    map[int]int                // Current rgb fixture shift. Indexed by sequence.
	ScannerShift                map[int]int                // Current scanner shift for all fixtures.  Indexed by sequence
	RGBSize                     map[int]int                // current RGB sequence this.Size[this.SelectedSequence]. Indexed by sequence
	ScannerSize                 map[int]int                // current scanner size for all fixtures. Indexed by sequence
	ScannerColor                int                        // current scanner color.
	RGBFade                     map[int]int                // Indexed by sequence.
	ScannerFade                 map[int]int                // Indexed by sequence.
	ScannerCoordinates          map[int]int                // Number of coordinates for scanner patterns is selected from 4 choices. ScannerCoordinates  0=12, 1=16,2=24,3=32,4=64, Indexed by sequence.
	Running                     map[int]bool               // Which sequence is running. Indexed by sequence. True if running.
	Strobe                      map[int]bool               // We are in strobe mode. True if strobing
	StrobeSpeed                 map[int]int                // Strobe speed. value is speed 0-255, indexed by sequence number.
	LastStrobeSpeed             map[int]int                // Last Strobe speed selected. value is speed 0-255, indexed by sequence number.
	SavePreset                  bool                       // Save a preset flag.
	Config                      bool                       // Flag to indicate we are in fixture config mode.
	Blackout                    bool                       // Blackout all fixtures.
	Flood                       bool                       // Flood all fixtures.
	SelectMode                  []int                      // What mode each sequence is in : normal mode, function mode, status selection mode.
	LastMode                    []int                      // Last mode sequence was in : normal mode, function mode, status selection mode.
	Functions                   map[int][]common.Function  // Map indexed by sequence of functions
	FunctionLabels              [8]string                  // Storage for the function key labels for this sequence.
	SelectButtonPressed         []bool                     // Which sequence has its Select button pressed.
	SwitchPositions             [9][9]int                  // Sorage for switch positions.
	EditScannerColorsMode       bool                       // This flag is true when the sequence is in select scanner colors editing mode.
	EditGoboSelectionMode       bool                       // This flag is true when the sequence is in sequence gobo selection mode.
	EditStaticColorsMode        []bool                     // This flag is true when the sequence is in edit static colors mode.
	ShowRGBColorPicker          bool                       // This flag is true when the sequence is in when we are showing the color picker.
	ShowStaticColorPicker       bool                       // This flag is true when the sequence is showing the static color picker mode.
	EditWhichStaticSequence     int                        // Which static sequence is currently being edited.
	EditPatternMode             bool                       // This flag is true when the sequence is in pattern editing mode.
	EditFixtureSelectionMode    bool                       // This flag is true when the sequence is in select fixture mode.
	MasterBrightness            int                        // Affects all DMX fixtures and launchpad lamps.
	LastStaticColorButtonX      int                        // Which Static Color button did we change last.
	LastStaticColorButtonY      int                        // Which Static Color button did we change last.
	SoundGain                   float32                    // Fine gain -0.09 -> 0.09
	FixtureState                [][]common.FixtureState    // Which fixture is enabled: bool and inverted: bool on which sequence. [sequeneNumber],[fixtureNumber]
	SelectedFixture             int                        // Which fixture is selected when changing scanner color or gobo.
	FollowingAction             string                     // String to find next function, used in selecting a fixture.
	OffsetPan                   int                        // Offset for Pan.
	OffsetTilt                  int                        // Offset for Tilt.
	Pad                         *pad.Pad                   // Pointer to the Novation Launchpad object.
	PresetsStore                map[string]presets.Preset  // Storage for the Presets.
	LastPreset                  *string                    // Last preset used.
	SoundTriggers               []*common.Trigger          // Pointer to the Sound Triggers.
	SoundConfig                 *sound.SoundConfig         // Pointer to the sound config struct.
	SequenceChannels            common.Channels            // Channles used to communicate with the sequence.
	RGBPatterns                 map[int]common.Pattern     // Available RGB Patterns.
	ScannerPattern              int                        // The selected scanner pattern Number. Used as the index for above.
	Pattern                     int                        // The selected RGB pattern Number. Used as the index for above.
	StaticButtons               []common.StaticColorButton // Storage for the color of the static buttons.
	SelectedGobo                int                        // The selected GOBO.
	ButtonTimer                 *time.Time                 // Button Timer
	ClearPressed                map[int]bool               // Storage clear pressed in static color selection. Indexed by sequence.
	SwitchChannels              []common.SwitchChannel     // Used for communicating with mini-sequencers on switches.
	LaunchPadConnected          bool                       // Flag to indicate presence of Novation Launchpad.
	DmxInterfacePresent         bool                       // Flag to indicate precence of DMX interface card
	DmxInterfacePresentConfig   *usbdmx.ControllerConfig   // DMX Interface card config.
	LaunchpadName               string                     // Storage for launchpad config.
	ScannerChaser               map[int]bool               // Chaser is running.
	DisplayChaserShortCut       bool                       // Flag to indicate we've taken a shortcut to the chaser display
	SwitchSequenceNumber        int                        // Switch sequence number, setup at start.
	ChaserSequenceNumber        int                        // Chaser sequence number, setup at start.
	ScannerSequenceNumber       int                        // Scanner sequence number, setup at start.
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
	updateChannels []chan common.Sequence) {

	if debug {
		fmt.Printf("ProcessButtons Called with X:%d Y:%d\n", X, Y)
	}

	// Set the sequence type.
	this.SelectedType = sequences[this.SelectedSequence].Type

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
		!this.EditStaticColorsMode[this.EditWhichStaticSequence] &&
		!this.ShowRGBColorPicker &&
		!this.ShowStaticColorPicker &&
		sequences[Y].Type != "switch" && // As long as we're not a switch sequence.
		this.SelectMode[Y] == NORMAL { // As long as we're in normal mode for this sequence.

		this.SelectedType = sequences[Y].Type

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

		red := flashSequence.Pattern.Steps[X].Fixtures[X].Color.R
		green := flashSequence.Pattern.Steps[X].Fixtures[X].Color.G
		blue := flashSequence.Pattern.Steps[X].Fixtures[X].Color.B
		white := flashSequence.Pattern.Steps[X].Fixtures[X].Color.W
		amber := flashSequence.Pattern.Steps[X].Fixtures[X].Color.A
		uv := flashSequence.Pattern.Steps[X].Fixtures[X].Color.UV
		pan := common.SCANNER_MID_POINT
		tilt := common.SCANNER_MID_POINT
		shutter := flashSequence.Pattern.Steps[X].Fixtures[X].Shutter
		rotate := flashSequence.Pattern.Steps[X].Fixtures[X].Rotate
		music := flashSequence.Pattern.Steps[X].Fixtures[X].Music
		gobo := flashSequence.Pattern.Steps[X].Fixtures[X].Gobo
		program := flashSequence.Pattern.Steps[X].Fixtures[X].Program

		if this.SelectedType == "rgb" {
			common.LightLamp(common.ALight{X: X, Y: Y, Brightness: this.MasterBrightness, Red: red, Green: green, Blue: blue}, eventsForLaunchpad, guiButtons)
			fixture.MapFixtures(false, false, Y, dmxController, X, red, green, blue, white, amber, uv, pan, tilt, shutter, rotate, music, program, gobo, 0, fixturesConfig, this.Blackout, this.MasterBrightness, this.MasterBrightness, this.Strobe[this.SelectedSequence], this.StrobeSpeed[this.SelectedSequence], this.DmxInterfacePresent)
		}
		if this.SelectedType == "scanner" {
			common.LightLamp(common.ALight{X: X, Y: Y, Brightness: this.MasterBrightness, Red: 255, Green: 255, Blue: 255}, eventsForLaunchpad, guiButtons)
			fixture.MapFixtures(false, false, Y, dmxController, X, red, green, blue, white, amber, uv, pan, tilt, shutter, rotate, music, program, gobo, 0, fixturesConfig, this.Blackout, this.MasterBrightness, this.MasterBrightness, this.Strobe[this.SelectedSequence], this.StrobeSpeed[this.SelectedSequence], this.DmxInterfacePresent)
		}

		if this.GUI {
			time.Sleep(200 * time.Millisecond)
			brightness := 0
			master := 0
			common.LightLamp(common.ALight{X: X, Y: Y, Brightness: 0, Red: 0, Green: 0, Blue: 0}, eventsForLaunchpad, guiButtons)
			fixture.MapFixtures(false, false, Y, dmxController, X, red, green, blue, white, amber, uv, pan, tilt, shutter, rotate, music, program, gobo, 0, fixturesConfig, this.Blackout, brightness, master, this.Strobe[this.SelectedSequence], this.StrobeSpeed[this.SelectedSequence], this.DmxInterfacePresent)
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
		!this.ShowRGBColorPicker &&
		!this.ShowStaticColorPicker &&
		sequences[Y].Type != "switch" && // As long as we're not a switch sequence.
		this.SelectMode[Y] == NORMAL { // As long as we're in normal mode for this sequence.

		if debug {
			fmt.Printf("Flash OFF Fixture Pressed X:%d Y:%d\n", X, Y)
		}

		X = X - 100

		red := 0
		green := 0
		blue := 0
		white := 0
		amber := 0
		uv := 0
		pan := common.SCANNER_MID_POINT
		tilt := common.SCANNER_MID_POINT
		shutter := 0
		rotate := 0
		music := 0
		gobo := 0
		program := 0
		brightness := 0
		master := 0

		common.LightLamp(common.ALight{X: X, Y: Y, Brightness: this.MasterBrightness, Red: 0, Green: 0, Blue: 0}, eventsForLaunchpad, guiButtons)
		fixture.MapFixtures(false, false, Y, dmxController, X, red, green, blue, white, amber, uv, pan, tilt, shutter, rotate, music, program, gobo, 0, fixturesConfig, this.Blackout, brightness, master, this.Strobe[this.SelectedSequence], this.StrobeSpeed[this.SelectedSequence], this.DmxInterfacePresent)
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

			if debug {
				fmt.Printf("Clear Preset X:%d Y:%d\n", X, Y)
			}

			// Delete the config file
			config.DeleteConfig(fmt.Sprintf("config%d.%d.json", X, Y))

			// Delete from preset store
			this.PresetsStore[fmt.Sprint(X)+","+fmt.Sprint(Y)] = presets.Preset{State: false, Selected: false, Label: "", ButtonColor: ""}

			// Update the copy of presets on disk.
			presets.SavePresets(this.PresetsStore)

			// Show presets again.
			presets.RefreshPresets(eventsForLaunchpad, guiButtons, this.PresetsStore)

		} else {
			if debug {
				fmt.Printf("Load Preset X:%d Y:%d\n", X, Y)
			}

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
	if X == 0 && Y == -1 && this.GUI {

		if debug {
			fmt.Printf("GUI Clear Pressed X:%d Y:%d\n", X, Y)
		}

		Clear(X, Y, this, sequences, dmxController, fixturesConfig, commandChannels, eventsForLaunchpad, guiButtons, updateChannels)
		return
	}

	// C L E A R  - Start the timer, waiting for a long press to clear all.
	// Because a short press in scanner mode shifts the scanners up.
	if X == 0 && Y == -1 && !this.GUI && sequences[this.SelectedSequence].Type == "scanner" {

		if debug {
			fmt.Printf("Clear Pressed Start Timer X:%d Y:%d\n", X, Y)
		}
		// Start a timer for this button.
		here := time.Now()
		this.ButtonTimer = &here
		return
	}

	//  C L E A R - clear all if we're not in the scanner mode.
	if X == 0 && Y == -1 && !this.GUI && sequences[this.SelectedSequence].Type != "scanner" {
		if debug {
			fmt.Printf("Clear All If We're Not in Scanner Mode X:%d Y:%d\n", X, Y)
		}
		Clear(X, Y, this, sequences, dmxController, fixturesConfig, commandChannels, eventsForLaunchpad, guiButtons, updateChannels)
		return
	}

	// C L E A R  - We have a long press.
	if X == 100 && Y == -1 && !this.GUI && sequences[this.SelectedSequence].Type == "scanner" {

		if debug {
			fmt.Printf("Clear Pressed Long Press X:%d Y:%d\n", X, Y)
		}

		// Remove the off button offset.
		X = X - 100
		// Stop the timer for this preset.
		elapsed := time.Since(*this.ButtonTimer)
		// If the timer is longer than 1 seconds then we have a long press.
		if elapsed > 1*time.Second {
			Clear(X, Y, this, sequences, dmxController, fixturesConfig, commandChannels, eventsForLaunchpad, guiButtons, updateChannels)
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
					this.PresetsStore[location] = presets.Preset{State: preset.State, Selected: false, Label: preset.Label, ButtonColor: preset.ButtonColor}
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
				this.PresetsStore[*this.LastPreset] = presets.Preset{State: lastPreset.State, Selected: true, Label: lastPreset.Label, ButtonColor: lastPreset.ButtonColor}
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
		if this.MasterBrightness > common.MAX_DMX_BRIGHTNESS {
			this.MasterBrightness = common.MAX_DMX_BRIGHTNESS
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

		if this.ShowRGBColorPicker {
			this.ShowRGBColorPicker = false
			removeColorPicker(this, eventsForLaunchpad, guiButtons, commandChannels)
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
					this.PresetsStore[location] = presets.Preset{State: preset.State, Selected: false, Label: preset.Label, ButtonColor: preset.ButtonColor}
				}
			}

			// Select this location and flash its button.
			this.PresetsStore[location] = presets.Preset{State: true, Selected: true, Label: current.Label, ButtonColor: current.ButtonColor}
			presets.RefreshPresets(eventsForLaunchpad, guiButtons, this.PresetsStore)

			if this.GUI {
				this.SavePreset = false
			}

		} else {
			// L O A D - Load config, but only if it exists in the presets map.
			if this.PresetsStore[location].State {

				if this.GUI { // GUI path.
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
	if X == 2 && Y == 7 && !this.ShowRGBColorPicker {

		if debug {
			fmt.Printf("Decrease Shift\n")
		}

		buttonTouched(common.ALight{X: X, Y: Y, OnColor: common.White, OffColor: common.Cyan}, eventsForLaunchpad, guiButtons)

		// If we're a scanner and we're in shutter chase mode.
		if sequences[this.SelectedSequence].Type == "scanner" && this.ScannerChaser[this.SelectedSequence] &&
			(this.SelectMode[this.SelectedSequence] == CHASER_FUNCTION || this.SelectMode[this.SelectedSequence] == CHASER_DISPLAY) {
			this.TargetSequence = this.ChaserSequenceNumber
		} else {
			this.TargetSequence = this.SelectedSequence
		}

		if sequences[this.TargetSequence].Type == "rgb" {
			this.RGBShift[this.TargetSequence] = this.RGBShift[this.TargetSequence] - 1
			if this.RGBShift[this.TargetSequence] < 0 {
				this.RGBShift[this.TargetSequence] = 0
			}
			cmd := common.Command{
				Action: common.UpdateRGBShift,
				Args: []common.Arg{
					{Name: "RGBShift", Value: this.RGBShift[this.TargetSequence]},
				},
			}
			common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

			// Update the status bar
			UpdateShift(this, guiButtons)
			return
		}

		if sequences[this.TargetSequence].Type == "scanner" {
			this.ScannerShift[this.TargetSequence] = this.ScannerShift[this.TargetSequence] - 1
			if this.ScannerShift[this.TargetSequence] < common.MIN_SCANNER_SHIFT {
				this.ScannerShift[this.TargetSequence] = common.MIN_SCANNER_SHIFT
			}
			cmd := common.Command{
				Action: common.UpdateScannerShift,
				Args: []common.Arg{
					{Name: "ScannerShift", Value: this.ScannerShift[this.TargetSequence]},
				},
			}
			common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

			// Update the status bar
			label := getScannerShiftLabel(this.ScannerShift[this.TargetSequence])
			common.UpdateStatusBar(fmt.Sprintf("Rotate Shift %0s", label), "shift", false, guiButtons)
			return
		}
	}

	// Increase Shift.
	if X == 3 && Y == 7 && !this.ShowRGBColorPicker {

		if debug {
			fmt.Printf("Increase Shift \n")
		}

		buttonTouched(common.ALight{X: X, Y: Y, OnColor: common.White, OffColor: common.Cyan}, eventsForLaunchpad, guiButtons)

		// If we're a scanner and we're in shutter chase mode.
		if sequences[this.SelectedSequence].Type == "scanner" && this.ScannerChaser[this.SelectedSequence] &&
			(this.SelectMode[this.SelectedSequence] == CHASER_FUNCTION || this.SelectMode[this.SelectedSequence] == CHASER_DISPLAY) {
			this.TargetSequence = this.ChaserSequenceNumber
		} else {
			this.TargetSequence = this.SelectedSequence
		}
		if sequences[this.TargetSequence].Type == "rgb" {
			this.RGBShift[this.TargetSequence] = this.RGBShift[this.TargetSequence] + 1
			if this.RGBShift[this.TargetSequence] > common.MAX_RGB_SHIFT {
				this.RGBShift[this.TargetSequence] = common.MAX_RGB_SHIFT
			}
			cmd := common.Command{
				Action: common.UpdateRGBShift,
				Args: []common.Arg{
					{Name: "Shift", Value: this.RGBShift[this.TargetSequence]},
				},
			}
			common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

			// Update the status bar
			UpdateShift(this, guiButtons)
			return
		}

		if sequences[this.TargetSequence].Type == "scanner" {
			this.ScannerShift[this.TargetSequence] = this.ScannerShift[this.TargetSequence] + 1
			if this.ScannerShift[this.TargetSequence] > common.MAX_SCANNER_SHIFT {
				this.ScannerShift[this.TargetSequence] = common.MAX_SCANNER_SHIFT
			}
			cmd := common.Command{
				Action: common.UpdateScannerShift,
				Args: []common.Arg{
					{Name: "ScannerShift", Value: this.ScannerShift[this.TargetSequence]},
				},
			}
			common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

			// Update the status bar
			label := getScannerShiftLabel(this.ScannerShift[this.TargetSequence])
			common.UpdateStatusBar(fmt.Sprintf("Rotate Shift %s", label), "shift", false, guiButtons)
			return
		}
	}

	// S E L E C T   S P E E D - Decrease speed of selected sequence.
	if X == 0 && Y == 7 && !this.ShowRGBColorPicker {

		if debug {
			fmt.Printf("Decrease Speed \n")
		}

		buttonTouched(common.ALight{X: X, Y: Y, OnColor: common.White, OffColor: common.Cyan}, eventsForLaunchpad, guiButtons)

		// If we're a scanner and we're in shutter chase mode.
		if sequences[this.SelectedSequence].Type == "scanner" && this.ScannerChaser[this.SelectedSequence] &&
			(this.SelectMode[this.SelectedSequence] == CHASER_FUNCTION || this.SelectMode[this.SelectedSequence] == CHASER_DISPLAY) {
			this.TargetSequence = this.ChaserSequenceNumber
		} else {
			this.TargetSequence = this.SelectedSequence
		}
		// Decrease Strobe Speed.
		if this.Strobe[this.TargetSequence] {
			this.StrobeSpeed[this.TargetSequence] -= 10
			if this.StrobeSpeed[this.TargetSequence] < 0 {
				this.StrobeSpeed[this.TargetSequence] = 0
			}

			// Store the last strobe speed.
			this.LastStrobeSpeed[this.SelectedSequence] = this.StrobeSpeed[this.TargetSequence]

			cmd := common.Command{
				Action: common.UpdateStrobeSpeed,
				Args: []common.Arg{
					{Name: "STROBE_SPEED", Value: this.StrobeSpeed[this.TargetSequence]},
				},
			}
			common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)
			// Update the status bar
			common.UpdateStatusBar(fmt.Sprintf("Strobe %02d", this.StrobeSpeed[this.TargetSequence]), "speed", false, guiButtons)
			return
		}

		// Get an upto date copy of the sequence.
		sequences[this.TargetSequence] = common.RefreshSequence(this.TargetSequence, commandChannels, updateChannels)

		// Decrease Speed.
		if !sequences[this.TargetSequence].MusicTrigger {
			this.Speed[this.TargetSequence]--
			if this.Speed[this.TargetSequence] < 1 {
				this.Speed[this.TargetSequence] = 1
			}

			// If you reached the min speed blink the increase button.
			if this.Speed[this.TargetSequence] == common.MIN_SPEED {
				common.FlashLight(X, Y, common.White, common.Cyan, eventsForLaunchpad, guiButtons)
			} else {
				// If you reached the half speed blink both buttons.
				if this.Speed[this.TargetSequence] == common.MAX_SPEED/2 {
					common.FlashLight(X, Y, common.White, common.Cyan, eventsForLaunchpad, guiButtons)
					common.FlashLight(X+1, Y, common.White, common.Cyan, eventsForLaunchpad, guiButtons)
				} else {
					common.LightLamp(common.ALight{X: X + 1, Y: Y, Brightness: 255, Red: 0, Green: 255, Blue: 255}, eventsForLaunchpad, guiButtons)
				}
			}

			cmd := common.Command{
				Action: common.UpdateSpeed,
				Args: []common.Arg{
					{Name: "Speed", Value: this.Speed[this.TargetSequence]},
				},
			}
			common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)
			// Speed is used to control fade time in mini sequencer so send to switch sequence as well.
			common.SendCommandToSequence(this.SwitchSequenceNumber, cmd, commandChannels)
		}

		// Update the status bar
		UpdateSpeed(this, guiButtons)

		return
	}

	// S E L E C T   S P E E D - Increase speed of selected sequence.
	if X == 1 && Y == 7 && !this.ShowRGBColorPicker {

		if debug {
			fmt.Printf("Increase Speed \n")
		}

		buttonTouched(common.ALight{X: X, Y: Y, OnColor: common.White, OffColor: common.Cyan}, eventsForLaunchpad, guiButtons)

		// If we're a scanner and we're in shutter chase mode
		if sequences[this.SelectedSequence].Type == "scanner" && this.ScannerChaser[this.SelectedSequence] &&
			(this.SelectMode[this.SelectedSequence] == CHASER_FUNCTION || this.SelectMode[this.SelectedSequence] == CHASER_DISPLAY) {
			this.TargetSequence = this.ChaserSequenceNumber
		} else {
			this.TargetSequence = this.SelectedSequence
		}

		if this.Strobe[this.TargetSequence] {
			this.StrobeSpeed[this.TargetSequence] += 10
			if this.StrobeSpeed[this.TargetSequence] > 255 {
				this.StrobeSpeed[this.TargetSequence] = 255
			}

			// Store the last strobe speed.
			this.LastStrobeSpeed[this.SelectedSequence] = this.StrobeSpeed[this.TargetSequence]

			cmd := common.Command{
				Action: common.UpdateStrobeSpeed,
				Args: []common.Arg{
					{Name: "STROBE_SPEED", Value: this.StrobeSpeed[this.TargetSequence]},
				},
			}
			common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)
			// Update the status bar
			common.UpdateStatusBar(fmt.Sprintf("Strobe %02d", this.StrobeSpeed[this.TargetSequence]), "speed", false, guiButtons)
			return
		}

		// Get an upto date copy of the sequence.
		sequences[this.TargetSequence] = common.RefreshSequence(this.TargetSequence, commandChannels, updateChannels)

		if !sequences[this.TargetSequence].MusicTrigger {
			this.Speed[this.TargetSequence]++
			if this.Speed[this.TargetSequence] > 12 {
				this.Speed[this.TargetSequence] = 12
			}

			// If you reached the max speed blink the increase button.
			if this.Speed[this.TargetSequence] == common.MAX_SPEED {
				common.FlashLight(X, Y, common.White, common.Cyan, eventsForLaunchpad, guiButtons)
			} else {
				// If you reached the half speed blink both buttons.
				if this.Speed[this.TargetSequence] == common.MAX_SPEED/2 {
					common.FlashLight(X, Y, common.White, common.Cyan, eventsForLaunchpad, guiButtons)
					common.FlashLight(X-1, Y, common.White, common.Cyan, eventsForLaunchpad, guiButtons)
				} else {
					common.LightLamp(common.ALight{X: X - 1, Y: Y, Brightness: 255, Red: 0, Green: 255, Blue: 255}, eventsForLaunchpad, guiButtons)
				}
			}

			cmd := common.Command{
				Action: common.UpdateSpeed,
				Args: []common.Arg{
					{Name: "Speed", Value: this.Speed[this.TargetSequence]},
				},
			}
			common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)
			// Speed is used to control fade time in mini sequencer so send to switch sequence as well.
			common.SendCommandToSequence(this.SwitchSequenceNumber, cmd, commandChannels)
		}

		// Update the status bar
		UpdateSpeed(this, guiButtons)

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

		this.ShowRGBColorPicker = false
		this.EditGoboSelectionMode = false
		this.DisplayChaserShortCut = false
		this.EditWhichStaticSequence = 0

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

		this.ShowRGBColorPicker = false
		this.EditGoboSelectionMode = false
		this.DisplayChaserShortCut = false
		this.EditWhichStaticSequence = 1

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

		this.ShowRGBColorPicker = false
		this.EditGoboSelectionMode = false
		if this.ScannerChaser[this.SelectedSequence] {
			this.EditWhichStaticSequence = 4
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

		this.ShowRGBColorPicker = false
		this.EditGoboSelectionMode = false
		this.DisplayChaserShortCut = false
		this.EditWhichStaticSequence = 3

		return
	}

	// S T A R T - Start sequence.
	if X == 8 && Y == 5 {

		if this.ShowRGBColorPicker {
			this.ShowRGBColorPicker = false
			removeColorPicker(this, eventsForLaunchpad, guiButtons, commandChannels)
		}

		// Start in normal mode, hide the shutter chaser.
		if this.SelectedType == "scanner" && this.ScannerChaser[this.SelectedSequence] {
			common.HideSequence(this.ChaserSequenceNumber, commandChannels)
			common.ClearSelectedRowOfButtons(this.SelectedSequence, eventsForLaunchpad, guiButtons)
			this.SelectMode[this.SelectedSequence] = NORMAL
		}

		this.SelectButtonPressed[this.SelectedSequence] = false
		this.StaticFlashing[this.SelectedSequence] = false

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

			this.Running[this.SelectedSequence] = false
			this.Functions[this.SelectedSequence][common.Function6_Static_Gobo].State = false
			this.Functions[this.SelectedSequence][common.Function8_Music_Trigger].State = false
			this.Functions[this.ChaserSequenceNumber][common.Function6_Static_Gobo].State = false
			this.Functions[this.ChaserSequenceNumber][common.Function8_Music_Trigger].State = false

			// Stop should also stop the shutter chaser.
			if this.SelectedType == "scanner" && this.ScannerChaser[this.SelectedSequence] {

				cmd := common.Command{
					Action: common.Stop,
					Args: []common.Arg{
						{Name: "Speed", Value: this.Speed[this.ChaserSequenceNumber]},
					},
				}
				common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)

				this.Functions[this.SelectedSequence][common.Function6_Static_Gobo].State = false
				this.Functions[this.SelectedSequence][common.Function7_Invert_Chase].State = false
				this.Functions[this.ChaserSequenceNumber][common.Function6_Static_Gobo].State = false
				this.Functions[this.ChaserSequenceNumber][common.Function8_Music_Trigger].State = false
				this.ScannerChaser[this.SelectedSequence] = false
				this.SelectMode[this.SelectedSequence] = NORMAL
				this.Running[this.ChaserSequenceNumber] = false
			}

			// Clear the pattern function keys
			common.ClearSelectedRowOfButtons(this.SelectedSequence, eventsForLaunchpad, guiButtons)

			// Turn off the start lamp.
			common.LightLamp(common.ALight{X: X, Y: Y, Brightness: this.MasterBrightness, Red: 255, Green: 255, Blue: 255}, eventsForLaunchpad, guiButtons)

			// Set the correct color for the select button.
			SequenceSelect(eventsForLaunchpad, guiButtons, this)

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
			this.Functions[this.SelectedSequence][common.Function6_Static_Gobo].State = false
			this.Functions[this.SelectedSequence][common.Function8_Music_Trigger].State = false

			// Clear the pattern function keys
			common.ClearSelectedRowOfButtons(this.SelectedSequence, eventsForLaunchpad, guiButtons)

			// Reveal the now running sequence
			common.RevealSequence(this.SelectedSequence, commandChannels)

			// Set the correct color for the select button.
			SequenceSelect(eventsForLaunchpad, guiButtons, this)

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
			// Start strobing for this sequence. Strobe on.
			this.Strobe[this.SelectedSequence] = true

			if this.LastStrobeSpeed[this.SelectedSequence] != this.StrobeSpeed[this.SelectedSequence] {
				// Then use it.
				this.StrobeSpeed[this.SelectedSequence] = this.LastStrobeSpeed[this.SelectedSequence]
			} else {
				// Use the default.
				this.StrobeSpeed[this.SelectedSequence] = 255
			}

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
	if X == 4 && Y == 7 && !this.ShowRGBColorPicker {

		if debug {
			fmt.Printf("Decrease Size\n")
		}

		buttonTouched(common.ALight{X: X, Y: Y, OnColor: common.White, OffColor: common.Cyan}, eventsForLaunchpad, guiButtons)

		// If we're a scanner and we're in shutter chase mode.
		if sequences[this.SelectedSequence].Type == "scanner" && this.ScannerChaser[this.SelectedSequence] &&
			(this.SelectMode[this.SelectedSequence] == CHASER_FUNCTION || this.SelectMode[this.SelectedSequence] == CHASER_DISPLAY) {
			this.TargetSequence = this.ChaserSequenceNumber
		} else {
			this.TargetSequence = this.SelectedSequence
		}

		if sequences[this.TargetSequence].Type == "rgb" {
			// Send Update RGB Size.
			this.RGBSize[this.TargetSequence]--
			if this.RGBSize[this.TargetSequence] < common.MIN_RGB_SIZE {
				this.RGBSize[this.TargetSequence] = common.MIN_RGB_SIZE
			}
			cmd := common.Command{
				Action: common.UpdateRGBSize,
				Args: []common.Arg{
					{Name: "RGBSize", Value: this.RGBSize[this.TargetSequence]},
				},
			}
			common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)
			// Update the status bar
			UpdateSize(this, guiButtons)
			return
		}

		if sequences[this.TargetSequence].Type == "scanner" {
			// Send Update Scanner Size.
			this.ScannerSize[this.TargetSequence] = this.ScannerSize[this.TargetSequence] - 10
			if this.ScannerSize[this.TargetSequence] < 0 {
				this.ScannerSize[this.TargetSequence] = 0
			}
			cmd := common.Command{
				Action: common.UpdateScannerSize,
				Args: []common.Arg{
					{Name: "ScannerSize", Value: this.ScannerSize[this.TargetSequence]},
				},
			}
			common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

			// Update the status bar
			common.UpdateStatusBar(fmt.Sprintf("Rotate Size %02d", this.ScannerSize[this.TargetSequence]), "size", false, guiButtons)
			return
		}
	}

	// Increase Size.
	if X == 5 && Y == 7 && !this.ShowRGBColorPicker {

		if debug {
			fmt.Printf("Increase Size\n")
		}

		buttonTouched(common.ALight{X: X, Y: Y, OnColor: common.White, OffColor: common.Cyan}, eventsForLaunchpad, guiButtons)

		// If we're a scanner and we're in shutter chase mode.
		if sequences[this.SelectedSequence].Type == "scanner" && this.ScannerChaser[this.SelectedSequence] &&
			(this.SelectMode[this.SelectedSequence] == CHASER_FUNCTION || this.SelectMode[this.SelectedSequence] == CHASER_DISPLAY) {
			this.TargetSequence = this.ChaserSequenceNumber
		} else {
			this.TargetSequence = this.SelectedSequence
		}
		if sequences[this.TargetSequence].Type == "rgb" {

			// Send Update RGB Size.
			this.RGBSize[this.TargetSequence]++
			if this.RGBSize[this.TargetSequence] > common.MAX_RGB_SIZE {
				this.RGBSize[this.TargetSequence] = common.MAX_RGB_SIZE
			}
			cmd := common.Command{
				Action: common.UpdateRGBSize,
				Args: []common.Arg{
					{Name: "RGBSize", Value: this.RGBSize[this.TargetSequence]},
				},
			}
			common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)
			// Update the status bar
			UpdateSize(this, guiButtons)
			return
		}

		if sequences[this.TargetSequence].Type == "scanner" {
			// Send Update Scanner Size.
			this.ScannerSize[this.TargetSequence] = this.ScannerSize[this.TargetSequence] + 10
			if this.ScannerSize[this.TargetSequence] > common.MAX_SCANNER_SIZE {
				this.ScannerSize[this.TargetSequence] = common.MAX_SCANNER_SIZE
			}
			cmd := common.Command{
				Action: common.UpdateScannerSize,
				Args: []common.Arg{
					{Name: "ScannerSize", Value: this.ScannerSize[this.TargetSequence]},
				},
			}
			common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

			// Update the status bar
			common.UpdateStatusBar(fmt.Sprintf("Rotate Size %02d", this.ScannerSize[this.TargetSequence]), "size", false, guiButtons)
			return
		}
	}

	// Fade time decrease.
	if X == 6 && Y == 7 && !this.ShowRGBColorPicker {

		if debug {
			fmt.Printf("Decrease Fade Time\n")
		}

		buttonTouched(common.ALight{X: X, Y: Y, OnColor: common.White, OffColor: common.Cyan}, eventsForLaunchpad, guiButtons)

		// If we're a scanner and we're in shutter chase mode.
		if sequences[this.SelectedSequence].Type == "scanner" && this.ScannerChaser[this.SelectedSequence] &&
			(this.SelectMode[this.SelectedSequence] == CHASER_FUNCTION || this.SelectMode[this.SelectedSequence] == CHASER_DISPLAY) {
			this.TargetSequence = this.ChaserSequenceNumber
		} else {
			this.TargetSequence = this.SelectedSequence
		}

		if sequences[this.TargetSequence].Type == "rgb" {
			this.RGBFade[this.TargetSequence]--
			if this.RGBFade[this.TargetSequence] < 1 {
				this.RGBFade[this.TargetSequence] = 1
			}
			// Send fade update command.
			cmd := common.Command{
				Action: common.UpdateRGBFadeSpeed,
				Args: []common.Arg{
					{Name: "RGBFadeSpeed", Value: this.RGBFade[this.TargetSequence]},
				},
			}
			common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)
			// Update the status bar
			UpdateFade(this, guiButtons)
			return
		}

		// Update Coordinates.
		if sequences[this.TargetSequence].Type == "scanner" {
			// Fade also send more or less coordinates for the scanner patterns.
			this.ScannerCoordinates[this.TargetSequence]--
			if this.ScannerCoordinates[this.TargetSequence] < 0 {
				this.ScannerCoordinates[this.TargetSequence] = 0
			}
			cmd := common.Command{
				Action: common.UpdateNumberCoordinates,
				Args: []common.Arg{
					{Name: "NumberCoordinates", Value: this.ScannerCoordinates[this.TargetSequence]},
				},
			}
			common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)
			// Update the status bar
			label := getScannerCoordinatesLabel(this.ScannerCoordinates[this.TargetSequence])
			common.UpdateStatusBar(fmt.Sprintf("Rotate Coord %s", label), "fade", false, guiButtons)
			return
		}

	}

	// Fade time increase.
	if X == 7 && Y == 7 && !this.ShowRGBColorPicker {

		if debug {
			fmt.Printf("Increase Fade Time\n")
		}

		buttonTouched(common.ALight{X: X, Y: Y, OnColor: common.White, OffColor: common.Cyan}, eventsForLaunchpad, guiButtons)

		// If we're a scanner and we're in shutter chase mode.
		if sequences[this.SelectedSequence].Type == "scanner" && this.ScannerChaser[this.SelectedSequence] &&
			(this.SelectMode[this.SelectedSequence] == CHASER_FUNCTION || this.SelectMode[this.SelectedSequence] == CHASER_DISPLAY) {
			this.TargetSequence = this.ChaserSequenceNumber
		} else {
			this.TargetSequence = this.SelectedSequence
		}

		if sequences[this.TargetSequence].Type == "rgb" {
			this.RGBFade[this.TargetSequence]++
			if this.RGBFade[this.TargetSequence] > common.MAX_RGB_FADE {
				this.RGBFade[this.TargetSequence] = common.MAX_RGB_FADE
			}
			// Send fade update command.
			cmd := common.Command{
				Action: common.UpdateRGBFadeSpeed,
				Args: []common.Arg{
					{Name: "FadeSpeed", Value: this.RGBFade[this.TargetSequence]},
				},
			}
			common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)
			// Update the status bar
			UpdateFade(this, guiButtons)
			return
		}

		// Update Coordinates.
		if sequences[this.TargetSequence].Type == "scanner" {
			// Fade also send more or less coordinates for the scanner patterns.
			this.ScannerCoordinates[this.TargetSequence]++
			if this.ScannerCoordinates[this.TargetSequence] > 4 {
				this.ScannerCoordinates[this.TargetSequence] = 4
			}
			cmd := common.Command{
				Action: common.UpdateNumberCoordinates,
				Args: []common.Arg{
					{Name: "NumberCoordinates", Value: this.ScannerCoordinates[this.TargetSequence]},
				},
			}
			common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)
			// Update the status bar
			label := getScannerCoordinatesLabel(this.ScannerCoordinates[this.TargetSequence])
			common.UpdateStatusBar(fmt.Sprintf("Rotate Coord %s", label), "fade", false, guiButtons)
			return
		}
	}

	// S W I T C H   B U T T O N's Toggle State of switches for this sequence.
	if X >= 0 && X < 8 &&
		Y >= 0 &&
		Y < 4 &&
		sequences[Y].Type == "switch" {

		if debug {
			fmt.Printf("Switch Key X:%d Y:%d\n", X, Y)
		}

		if this.ShowRGBColorPicker {
			this.ShowRGBColorPicker = false
			removeColorPicker(this, eventsForLaunchpad, guiButtons, commandChannels)
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
			fmt.Printf("Fixture State Enabled %t  Inverted %t Reversed %t\n", this.FixtureState[Y][X].Enabled, this.FixtureState[Y][X].RGBInverted, this.FixtureState[Y][X].RGBInverted)
		}

		// Rotate the  fixture state based on last fixture state.
		setFixtureStatus(this, Y, X, commandChannels, sequences[Y])

		// Show the status.
		showFixtureStatus(Y, sequences[Y].Number, sequences[Y].NumberFixtures, this, eventsForLaunchpad, guiButtons, commandChannels)

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

	// RIGHT ARROW
	if X == 3 && Y == -1 && sequences[this.SelectedSequence].Type == "scanner" {

		if debug {
			fmt.Printf("RIGHT ARROW\n")
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

	// S E L E C T   S T A T I C   C O L O R
	// Red
	if X == 1 && Y == -1 && this.EditStaticColorsMode[this.TargetSequence] {

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
	if X == 2 && Y == -1 && this.EditStaticColorsMode[this.TargetSequence] {

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
	if X == 3 && Y == -1 && this.EditStaticColorsMode[this.TargetSequence] {

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

	// S E L E C T   R G B   S E Q U E N C E  C H A S E  C O L O R's
	if X >= 0 && X < 8 &&
		Y != -1 && Y < 3 &&
		!this.EditFixtureSelectionMode &&
		!this.EditScannerColorsMode &&
		this.ShowRGBColorPicker {

		if debug {
			fmt.Printf("Set Sequence Color X:%d Y:%d\n", X, Y)
		}

		if this.SelectMode[this.SelectedSequence] == CHASER_FUNCTION {
			this.TargetSequence = this.ChaserSequenceNumber
			this.DisplaySequence = this.SelectedSequence
		} else {
			this.TargetSequence = this.SelectedSequence
			this.DisplaySequence = this.SelectedSequence
		}

		// Reset the clear button so you can clear this selection if required.
		this.ClearPressed[this.TargetSequence] = false

		// Add the selected color to the sequence.
		cmd := common.Command{
			Action: common.UpdateASingeSequenceColor,
			Args: []common.Arg{
				{Name: "SelectedX", Value: X},
				{Name: "SelectedY", Value: Y},
			},
		}
		common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

		this.ShowRGBColorPicker = true

		// Get an upto date copy of the sequence.
		sequences[this.TargetSequence] = common.RefreshSequence(this.TargetSequence, commandChannels, updateChannels)

		// Set the colors.
		sequences[this.TargetSequence].CurrentColors = sequences[this.TargetSequence].SequenceColors

		// We call ShowRGBColorPicker here so the selections will flash as you press them.
		ShowRGBColorPicker(this.MasterBrightness, *sequences[this.TargetSequence], this.DisplaySequence, eventsForLaunchpad, guiButtons, commandChannels)

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
		if this.ScannerChaser[this.SelectedSequence] && this.SelectedType == "scanner" {
			cmd := common.Command{
				Action: common.UpdateScannerColor,
				Args: []common.Arg{
					{Name: "SelectedColor", Value: this.ScannerColor},
					{Name: "SelectedFixture", Value: this.SelectedFixture},
				},
			}
			common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)
		}

		this.EditScannerColorsMode = true

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
		if this.ScannerChaser[this.SelectedSequence] && this.SelectedType == "scanner" {
			cmd := common.Command{
				Action: common.UpdateGobo,
				Args: []common.Arg{
					{Name: "SelectedGobo", Value: this.SelectedGobo},
					{Name: "FixtureNumber", Value: this.SelectedFixture},
				},
			}
			common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)
		}

		this.EditGoboSelectionMode = true

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

	// S E L E C T   S T A T I C   F I X T U R E
	if X >= 0 && X < 8 &&
		Y != -1 &&
		!this.EditFixtureSelectionMode &&
		this.SelectedSequence == Y && // Make sure the buttons pressed are for this sequence.
		(this.SelectMode[this.SelectedSequence] == NORMAL ||
			this.SelectMode[this.SelectedSequence] == NORMAL_STATIC ||
			this.SelectMode[this.SelectedSequence] == CHASER_DISPLAY ||
			this.SelectMode[this.SelectedSequence] == CHASER_DISPLAY_STATIC) && // Not in function Mode
		!this.ShowStaticColorPicker && // Not In Color Picker Mode.
		this.EditStaticColorsMode[this.SelectedSequence] { // Static Function On in any sequence

		this.TargetSequence = this.EditWhichStaticSequence
		this.DisplaySequence = this.SelectedSequence

		if debug {
			fmt.Printf("EditWhichStaticSequence %d\n", this.EditWhichStaticSequence)
			fmt.Printf("TargetSequence %d\n", this.TargetSequence)
			fmt.Printf("DisplaySequence %d\n", this.DisplaySequence)
		}

		// Save the selected fixture number.
		this.SelectedStaticFixtureNumber = X

		// Reset Clear pressed flag so we can clear next selection
		this.ClearPressed[this.TargetSequence] = false

		// The current color is help in our local copy.
		color := sequences[this.TargetSequence].StaticColors[X].Color
		empty := common.Color{}
		if color == empty {
			color = FindCurrentColor(this.SelectedStaticFixtureNumber, this.SelectedSequence, *sequences[this.TargetSequence])
		}

		if debug {
			fmt.Printf("Sequence %d Fixture %d Setting Current Color as %+v\n", this.SelectedSequence, this.SelectedStaticFixtureNumber, color)
		}

		// Set the fixture color so that it flashs in the color picker.
		sequences[this.TargetSequence].CurrentColors = SetRGBColorPicker(color, *sequences[this.TargetSequence])

		// We call ShowRGBColorPicker so you can choose the static color for this fixture.
		ShowRGBColorPicker(this.MasterBrightness, *sequences[this.TargetSequence], this.DisplaySequence, eventsForLaunchpad, guiButtons, commandChannels)

		// Switch the mode so we know we are picking a static color from the color picker.
		this.ShowStaticColorPicker = true
		this.EditStaticColorsMode[this.TargetSequence] = true

		return
	}

	// S E T   S T A T I C   C O L O R
	if X >= 0 && X < 8 && Y != -1 &&
		Y < 3 && // Make sure the buttons pressed inside the color picker.
		this.ShowStaticColorPicker && // Now We Are In Static Color Picker Mode.
		!this.EditFixtureSelectionMode && // Not In Fixture Selection Mode.
		this.EditStaticColorsMode[this.SelectedSequence] { // Static Function On in this sequence

		this.TargetSequence = this.EditWhichStaticSequence
		this.DisplaySequence = this.SelectedSequence

		if debug {
			fmt.Printf("EditWhichStaticSequence %d\n", this.EditWhichStaticSequence)
			fmt.Printf("ShowStaticColorPicker %t\n", this.ShowStaticColorPicker)
			fmt.Printf("TargetSequence %d\n", this.TargetSequence)
			fmt.Printf("DisplaySequence %d\n", this.DisplaySequence)
		}

		// Find the color from the button pressed.
		color := FindCurrentColor(X, Y, *sequences[this.TargetSequence])

		if debug {
			fmt.Printf("Selected Static Color for X %d  Y %d to Color %+v\n", this.SelectedStaticFixtureNumber, Y, color)
		}

		// Set the fixture color so that it flashs in the color picker.
		sequences[this.TargetSequence].CurrentColors = SetRGBColorPicker(color, *sequences[this.TargetSequence])

		// Set our local copy of the color.
		sequences[this.TargetSequence].StaticColors[this.SelectedStaticFixtureNumber].Color = color

		// Tell the sequence about the new color and where we are in the
		// color cycle.

		if this.SelectAllStaticFixtures {
			// Set the same static color for all.
			cmd := common.Command{
				Action: common.UpdateAllStaticColor,
				Args: []common.Arg{
					{Name: "Static", Value: true},
					{Name: "StaticLampFlash", Value: false},
					{Name: "SelectedColor", Value: sequences[this.TargetSequence].StaticColors[this.SelectedStaticFixtureNumber].SelectedColor},
					{Name: "StaticColor", Value: sequences[this.TargetSequence].StaticColors[this.SelectedStaticFixtureNumber].Color},
				},
			}
			common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

			this.SelectAllStaticFixtures = false

		} else {
			// Set a static color for an individual fixture.
			cmd := common.Command{
				Action: common.UpdateStaticColor,
				Args: []common.Arg{
					{Name: "Static", Value: true},
					{Name: "FixtureNumber", Value: this.SelectedStaticFixtureNumber},
					{Name: "StaticLampFlash", Value: true},
					{Name: "SelectedColor", Value: sequences[this.TargetSequence].StaticColors[this.SelectedStaticFixtureNumber].SelectedColor},
					{Name: "StaticColor", Value: sequences[this.TargetSequence].StaticColors[this.SelectedStaticFixtureNumber].Color},
				},
			}
			common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)
		}

		this.LastStaticColorButtonX = this.SelectedStaticFixtureNumber
		this.LastStaticColorButtonY = Y

		// Hide the sequence.
		common.HideSequence(this.TargetSequence, commandChannels)

		// We call ShowRGBColorPicker so you can see which static color has been selected for this fixture.
		ShowRGBColorPicker(this.MasterBrightness, *sequences[this.TargetSequence], this.DisplaySequence, eventsForLaunchpad, guiButtons, commandChannels)

		// Set the first pressed for only this fixture and cancel any others
		for x := 0; x < 8; x++ {
			sequences[this.TargetSequence].StaticColors[x].FirstPress = false
		}
		sequences[this.TargetSequence].StaticColors[this.SelectedStaticFixtureNumber].FirstPress = true

		// Remove the color picker and reveal the sequence.
		removeColorPicker(this, eventsForLaunchpad, guiButtons, commandChannels)

		// Show static colors.
		cmd := common.Command{
			Action: common.UpdateStatic,
			Args: []common.Arg{
				{Name: "Static", Value: true},
			},
		}
		common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)
		common.RevealSequence(this.TargetSequence, commandChannels)

		// Switch off the color picker.
		this.ShowStaticColorPicker = false

		return
	}

	// S E L E C T   P A T T E N
	if X >= 0 && X < 8 && Y != -1 &&
		!this.EditFixtureSelectionMode &&
		this.EditPatternMode {

		if this.SelectMode[this.SelectedSequence] == CHASER_FUNCTION {
			this.TargetSequence = this.ChaserSequenceNumber
			this.DisplaySequence = this.SelectedSequence
			// This is the way we get the shutter chaser to be displayed as we exit
			// the pattern selection
			this.DisplayChaserShortCut = true
			this.SelectMode[this.DisplaySequence] = CHASER_FUNCTION
		} else {
			this.TargetSequence = this.SelectedSequence
			this.DisplaySequence = this.SelectedSequence
			this.SelectMode[this.DisplaySequence] = NORMAL
		}

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
		common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

		// Get an upto date copy of the sequence.
		sequences[this.TargetSequence] = common.RefreshSequence(this.TargetSequence, commandChannels, updateChannels)

		// We call ShowPatternSelectionButtons here so the selections will flash as you press them.
		this.EditFixtureSelectionMode = false
		ShowPatternSelectionButtons(this, sequences[this.SelectedSequence].Master, *sequences[this.TargetSequence], this.DisplaySequence, eventsForLaunchpad, guiButtons)

		return
	}

	// F U N C T I O N  K E Y S
	if X >= 0 && X < 8 && Y >= 0 && Y < 3 &&
		this.SelectedSequence == Y && // Make sure the buttons pressed are for this sequence.
		this.SelectMode[this.SelectedSequence] == FUNCTION || this.SelectMode[this.SelectedSequence] == CHASER_FUNCTION {
		processFunctions(X, Y, sequences, this, eventsForLaunchpad, guiButtons, commandChannels, updateChannels)
		return
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

func SetRGBColorPicker(selectedColor common.Color, targetSequence common.Sequence) []common.Color {

	// Clear out exiting colors.
	targetSequence.CurrentColors = []common.Color{}

	for _, availableColor := range targetSequence.RGBAvailableColors {
		if availableColor.Color == selectedColor {
			if debug {
				fmt.Printf("Adding color %+v\n", selectedColor)
			}
			targetSequence.CurrentColors = append(targetSequence.CurrentColors, selectedColor)
		}
	}
	return targetSequence.CurrentColors
}

func FindCurrentColor(X int, Y int, targetSequence common.Sequence) common.Color {

	for _, availableColor := range targetSequence.RGBAvailableColors {
		if availableColor.X == X && availableColor.Y == Y {
			return availableColor.Color
		}
	}

	return common.Color{}
}

// For the given sequence show the available sequence colors on the relevant buttons.
// With the new color picker there can be 24 colors displayed.
// ShowRGBColorPicker operates on the sequence.RGBAvailableColors which is an array of type []common.StaticColorButton
// the targetSequence .CurrentColors selects which colors are selected.
// Returns the RGBAvailableColors []common.StaticColorButton
func ShowRGBColorPicker(master int, targetSequence common.Sequence, displaySequence int, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command) {

	if debug {
		fmt.Printf("Color Picker - Show Color Selection Buttons\n")
	}

	common.HideSequence(0, commandChannels)
	common.HideSequence(1, commandChannels)
	common.HideSequence(1, commandChannels)

	for myFixtureNumber, lamp := range targetSequence.RGBAvailableColors {

		lamp.Flash = false

		// Check if we need to flash this button.
		for index, availableColor := range targetSequence.RGBAvailableColors {
			for _, sequenceColor := range targetSequence.CurrentColors {
				if availableColor.Color == sequenceColor {
					if myFixtureNumber == index {
						if debug {
							fmt.Printf("myFixtureNumber %d   current color %+v\n", myFixtureNumber, sequenceColor)
						}
						lamp.Flash = true
					}
				}
			}
		}
		if lamp.Flash {
			Black := common.Color{R: 0, G: 0, B: 0}
			if debug {
				fmt.Printf("FLASH myFixtureNumber X:%d Y:%d Color %+v \n", lamp.X, lamp.Y, lamp.Color)
			}
			common.FlashLight(lamp.X, lamp.Y, lamp.Color, Black, eventsForLaunchpad, guiButtons)
		} else {
			common.LightLamp(common.ALight{X: lamp.X, Y: lamp.Y, Brightness: master, Red: lamp.Color.R, Green: lamp.Color.G, Blue: lamp.Color.B}, eventsForLaunchpad, guiButtons)
		}
		common.LabelButton(lamp.X, lamp.Y, lamp.Name, guiButtons)

		time.Sleep(10 * time.Millisecond)
	}
}

// For the given sequence show the available scanner selection colors on the relevant buttons.
func ShowSelectFixtureButtons(targetSequence common.Sequence, displaySequence int, this *CurrentState, eventsForLaunchpad chan common.ALight, action string, guiButtons chan common.ALight) int {

	if debug {
		fmt.Printf("Sequence %d Show Fixture Selection Buttons on the way to %s\n", this.SelectedSequence, action)
	}

	for fixtureNumber, fixture := range targetSequence.ScannersAvailable {

		if debug {
			fmt.Printf("Fixture %+v\n", fixture)
		}
		if fixtureNumber == this.SelectedFixture {
			fixture.Flash = true
			this.SelectedFixture = fixtureNumber
		}
		if fixture.Flash {
			White := common.Color{R: 255, G: 255, B: 255}
			common.FlashLight(fixtureNumber, displaySequence, fixture.Color, White, eventsForLaunchpad, guiButtons)

		} else {
			common.LightLamp(common.ALight{X: fixtureNumber, Y: displaySequence, Red: fixture.Color.R, Green: fixture.Color.G, Blue: fixture.Color.B, Brightness: targetSequence.Master}, eventsForLaunchpad, guiButtons)
		}
		common.LabelButton(fixtureNumber, displaySequence, fixture.Label, guiButtons)
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

		// Turn off the color edit mode.
		this.ShowRGBColorPicker = false
		// And since we seem to be using two flags for the same thing, turn this off too.
		this.Functions[this.SelectedSequence][common.Function5_Color].State = false

		for _, fixture := range fixtures.Fixtures {
			if fixture.Group == this.SelectedSequence+1 {
				common.LightLamp(common.ALight{X: fixture.Number - 1, Y: this.SelectedSequence, Red: 0, Green: 0, Blue: 0, Brightness: sequence.Master}, eventsForLaunchpad, guiButtons)
			}
		}
		if this.GUI {
			displayErrorPopUp(this.MyWindow, "no colors available for this fixture")
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

func displayErrorPopUp(w fyne.Window, errorMessage string) (modal *widget.PopUp) {

	title := widget.NewLabel("Error")
	title.TextStyle = fyne.TextStyle{
		Bold: true,
	}

	errorText := widget.NewLabel(errorMessage)

	// Ok button.
	button := widget.NewButton("Dismiss", func() {
		modal.Hide()
	})

	// Layout of settings panel.
	modal = widget.NewModalPopUp(
		container.NewVBox(
			title,
			errorText,
			widget.NewLabel(""),
			container.NewHBox(layout.NewSpacer(), button),
		),
		w.Canvas(),
	)

	modal.Resize(fyne.NewSize(250, 250))
	modal.Show()

	return modal
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
// this.TargetSequence - is the squence you are updating the pattern, this could be different in the case
// of scanner shutter chaser sequence which doesn't have it's own buttons.
func ShowPatternSelectionButtons(this *CurrentState, master int, targetSequence common.Sequence, displaySequence int, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight) {

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
				common.FlashLight(pattern.Number, displaySequence, White, LightBlue, eventsForLaunchpad, guiButtons)
			} else {
				common.LightLamp(common.ALight{X: pattern.Number, Y: displaySequence, Red: 0, Green: 100, Blue: 255, Brightness: master}, eventsForLaunchpad, guiButtons)
			}
			common.LabelButton(pattern.Number, displaySequence, pattern.Label, guiButtons)
		}
		return
	}

	if targetSequence.Type == "scanner" {
		for _, pattern := range targetSequence.ScannerAvailablePatterns {
			if pattern.Number == targetSequence.SelectedPattern {
				common.FlashLight(pattern.Number, displaySequence, White, LightBlue, eventsForLaunchpad, guiButtons)
			} else {
				common.LightLamp(common.ALight{X: pattern.Number, Y: displaySequence, Red: 0, Green: 100, Blue: 255, Brightness: master}, eventsForLaunchpad, guiButtons)
			}
			common.LabelButton(pattern.Number, displaySequence, pattern.Label, guiButtons)
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
	SequenceSelect(eventsForLaunchpad, guiButtons, this)

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
		this.ShowRGBColorPicker = false
		this.EditStaticColorsMode[this.DisplaySequence] = false
		this.EditStaticColorsMode[this.TargetSequence] = false
		this.ShowStaticColorPicker = false
		this.EditGoboSelectionMode = false
		this.EditPatternMode = false
		for function := range this.Functions {
			this.Functions[sequenceNumber][function].State = false
		}
	}
}

func printMode(mode int) string {
	if mode == NORMAL {
		return "NORMAL"
	}
	if mode == NORMAL_STATIC {
		return "NORMAL_STATIC"
	}
	if mode == CHASER_DISPLAY {
		return "CHASER_DISPLAY"
	}
	if mode == CHASER_DISPLAY_STATIC {
		return "CHASER_DISPLAY_STATIC"
	}
	if mode == FUNCTION {
		return "FUNCTION"
	}
	if mode == CHASER_FUNCTION {
		return "CHASER_FUNCTION"
	}
	if mode == STATUS {
		return "STATUS"
	}
	return "UNKNOWN"
}

func SequenceSelect(eventsForLauchpad chan common.ALight, guiButtons chan common.ALight, this *CurrentState) {

	if debug {
		fmt.Printf("SequenceSelect\n")
	}

	// Turn off all sequence lights.
	for seq := 0; seq < 3; seq++ {
		common.LightLamp(common.ALight{X: 8, Y: seq, Brightness: 255, Red: 100, Green: 255, Blue: 255}, eventsForLauchpad, guiButtons)
	}

	if this.SelectedType == "scanner" && this.ScannerChaser[this.SelectedSequence] &&
		(this.SelectMode[this.SelectedSequence] == CHASER_FUNCTION || this.SelectMode[this.SelectedSequence] == CHASER_DISPLAY) {
		// If we are in shutter chaser mode, light the lamp yellow.
		common.LightLamp(common.ALight{X: 8, Y: this.SelectedSequence, Brightness: 255, Red: 255, Green: 255, Blue: 0}, eventsForLauchpad, guiButtons)
	} else {
		// Now turn pink the selected sequence select light.
		common.LightLamp(common.ALight{X: 8, Y: this.SelectedSequence, Brightness: 255, Red: 255, Green: 0, Blue: 255}, eventsForLauchpad, guiButtons)
	}

}

func UpdateSpeed(this *CurrentState, guiButttons chan common.ALight) {

	mode := this.SelectMode[this.DisplaySequence]
	tYpe := this.SelectedType
	speed := this.Speed[this.TargetSequence]

	if mode == NORMAL || mode == FUNCTION || mode == STATUS {
		if tYpe == "rgb" {
			common.UpdateStatusBar(fmt.Sprintf("Speed %02d", speed), "speed", false, guiButttons)
		}
		if tYpe == "scanner" {
			common.UpdateStatusBar(fmt.Sprintf("Rotate Speed %02d", speed), "speed", false, guiButttons)
		}
	}
	if mode == CHASER_DISPLAY || mode == CHASER_FUNCTION {
		common.UpdateStatusBar(fmt.Sprintf("Chase Speed %02d", speed), "speed", false, guiButttons)
	}
}

func UpdateSize(this *CurrentState, guiButttons chan common.ALight) {

	mode := this.SelectMode[this.DisplaySequence]
	tYpe := this.SelectedType
	size := this.RGBSize[this.TargetSequence]

	if mode == NORMAL || mode == FUNCTION || mode == STATUS {
		if tYpe == "rgb" {
			common.UpdateStatusBar(fmt.Sprintf("Size %02d", size), "size", false, guiButttons)
		}
		if tYpe == "scanner" {
			common.UpdateStatusBar(fmt.Sprintf("Rotate Size %02d", size), "size", false, guiButttons)
		}
	}
	if mode == CHASER_DISPLAY || mode == CHASER_FUNCTION {
		common.UpdateStatusBar(fmt.Sprintf("Chase Size %02d", size), "size", false, guiButttons)
	}
}

func UpdateShift(this *CurrentState, guiButttons chan common.ALight) {

	mode := this.SelectMode[this.DisplaySequence]
	tYpe := this.SelectedType
	shift := this.RGBShift[this.TargetSequence]

	if mode == NORMAL || mode == FUNCTION || mode == STATUS {
		if tYpe == "rgb" {
			common.UpdateStatusBar(fmt.Sprintf("Shift %02d", shift), "shift", false, guiButttons)
		}
		if tYpe == "scanner" {
			common.UpdateStatusBar(fmt.Sprintf("Rotate Shift %02d", shift), "shift", false, guiButttons)
		}
	}
	if mode == CHASER_DISPLAY || mode == CHASER_FUNCTION {
		common.UpdateStatusBar(fmt.Sprintf("Chase Shift %02d", shift), "shift", false, guiButttons)
	}
}

func UpdateFade(this *CurrentState, guiButttons chan common.ALight) {

	mode := this.SelectMode[this.DisplaySequence]
	tYpe := this.SelectedType
	fade := this.RGBFade[this.TargetSequence]

	if mode == NORMAL || mode == FUNCTION || mode == STATUS {
		if tYpe == "rgb" {
			common.UpdateStatusBar(fmt.Sprintf("Fade %02d", fade), "fade", false, guiButttons)
		}
		if tYpe == "scanner" {
			common.UpdateStatusBar(fmt.Sprintf("Rotate Fade %02d", fade), "fade", false, guiButttons)
		}
	}
	if mode == CHASER_DISPLAY || mode == CHASER_FUNCTION {
		common.UpdateStatusBar(fmt.Sprintf("Chase Fade %02d", fade), "fade", false, guiButttons)
	}
}
