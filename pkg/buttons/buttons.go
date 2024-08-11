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
	"image/color"
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
const NUMBER_SWITCHES int = 8
const NUMBER_SEQUENCES int = 5

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
	SelectedSwitch              int                        // The currently selected switch.
	TargetSequence              int                        // The current target sequence.
	DisplaySequence             int                        // The current display sequence.
	SequenceType                []string                   // The type, indexed by sequence.
	SelectedStaticFixtureNumber int                        // Temporary storage for the selected fixture number, used by color picker.
	SelectAllStaticFixtures     bool                       // Flag that indicate that all static fixtures have been selected.
	StaticFlashing              []bool                     // Static buttons are flashing, indexed by sequence.
	SavedSequenceColors         map[int][]color.RGBA       // Local storage for sequence colors.
	SelectedType                string                     // The currently selected sequenece type.
	SelectedFixtureType         string                     // The use fixture type for a switch.
	LastSelectedSwitch          int                        // The last selected switch.
	LastSelectedSequence        int                        // Store fof the last selected squence.
	MusicTrigger                bool                       // Does this seleted switch have a music trigger.
	Speed                       map[int]int                // Local copy of sequence speed. Indexed by sequence.
	SwitchOverrides             [][]common.Override        // Local copy of overriden switch fixture speeds. Indexed by switch number and state.
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
	SavePreset                  bool                       // Save a preset flag.
	Config                      bool                       // Flag to indicate we are in fixture config mode.
	Blackout                    bool                       // Blackout all fixtures.
	Flood                       bool                       // Flood all fixtures.
	SelectedMode                []int                      // What mode each sequence is in : normal mode, function mode, status selection mode.
	LastMode                    []int                      // Last mode sequence was in : normal mode, function mode, status selection mode.
	Functions                   map[int][]common.Function  // Map indexed by sequence of functions
	FunctionLabels              [8]string                  // Storage for the function key labels for this sequence.
	SelectButtonPressed         []bool                     // Which sequence has its Select button pressed.
	SwitchPosition              [NUMBER_SWITCHES]int       // Sorage for switch positions.
	EditScannerColorsMode       bool                       // This flag is true when the sequence is in select scanner colors editing mode.
	EditGoboSelectionMode       bool                       // This flag is true when the sequence is in sequence gobo selection mode.
	Static                      []bool                     // This flag is true when the sequence is in edit static colors mode.
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
		staticColors := []color.RGBA{}
		for _, button := range sequences[0].StaticColors {
			staticColors = append(staticColors, button.Color)
		}
		InitButtons(this, sequences[0].SequenceColors, staticColors, eventsForLaunchpad, guiButtons)

		// Show the static and switch settings.
		cmd := common.Command{
			Action: common.Reveal,
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
		!this.Static[this.EditWhichStaticSequence] &&
		!this.ShowRGBColorPicker &&
		!this.ShowStaticColorPicker &&
		sequences[Y].Type != "switch" && // As long as we're not a switch sequence.
		(this.SelectedMode[Y] == NORMAL || this.SelectedMode[Y] == CHASER_DISPLAY) { // As long as we're in normal or shutter chaser mode for this sequence.

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

		pan := common.SCANNER_MID_POINT
		tilt := common.SCANNER_MID_POINT
		color := flashSequence.Pattern.Steps[X].Fixtures[X].Color
		shutter := flashSequence.Pattern.Steps[X].Fixtures[X].Shutter
		rotate := flashSequence.Pattern.Steps[X].Fixtures[X].Rotate
		music := flashSequence.Pattern.Steps[X].Fixtures[X].Music
		gobo := flashSequence.Pattern.Steps[X].Fixtures[X].Gobo
		program := flashSequence.Pattern.Steps[X].Fixtures[X].Program

		if this.SelectedType == "rgb" {
			common.LightLamp(common.Button{X: X, Y: Y}, color, this.MasterBrightness, eventsForLaunchpad, guiButtons)
			fixture.MapFixtures(false, false, Y, X, color, pan, tilt, shutter, rotate, program, gobo, 0, fixturesConfig, this.Blackout, this.MasterBrightness, this.MasterBrightness, music, this.Strobe[this.SelectedSequence], this.StrobeSpeed[this.SelectedSequence], dmxController, this.DmxInterfacePresent)
		}
		if this.SelectedType == "scanner" {
			common.LightLamp(common.Button{X: X, Y: Y}, common.White, this.MasterBrightness, eventsForLaunchpad, guiButtons)
			fixture.MapFixtures(false, false, Y, X, color, pan, tilt, shutter, rotate, program, gobo, 0, fixturesConfig, this.Blackout, this.MasterBrightness, this.MasterBrightness, music, this.Strobe[this.SelectedSequence], this.StrobeSpeed[this.SelectedSequence], dmxController, this.DmxInterfacePresent)
		}

		if this.GUI {
			time.Sleep(200 * time.Millisecond)
			brightness := 0
			master := 0
			common.LightLamp(common.Button{X: X, Y: Y}, common.Black, common.MIN_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)
			fixture.MapFixtures(false, false, Y, X, color, pan, tilt, shutter, rotate, program, gobo, 0, fixturesConfig, this.Blackout, brightness, master, music, this.Strobe[this.SelectedSequence], this.StrobeSpeed[this.SelectedSequence], dmxController, this.DmxInterfacePresent)
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
		this.SelectedMode[Y] == NORMAL { // As long as we're in normal mode for this sequence.

		if debug {
			fmt.Printf("Flash OFF Fixture Pressed X:%d Y:%d\n", X, Y)
		}

		X = X - 100

		pan := common.SCANNER_MID_POINT
		tilt := common.SCANNER_MID_POINT
		shutter := 0
		rotate := 0
		music := 0
		gobo := 0
		program := 0
		brightness := 0
		master := 0

		common.LightLamp(common.Button{X: X, Y: Y}, common.Black, common.MIN_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)
		fixture.MapFixtures(false, false, Y, X, common.Black, pan, tilt, shutter, rotate, program, gobo, 0, fixturesConfig, this.Blackout, brightness, master, music, this.Strobe[this.SelectedSequence], this.StrobeSpeed[this.SelectedSequence], dmxController, this.DmxInterfacePresent)
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
			loadConfig(sequences, this, X, Y, commandChannels, eventsForLaunchpad, guiButtons, updateChannels)
			common.StartStaticSequences(sequences, commandChannels)
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

				buttonTouched(common.Button{X: X, Y: Y}, common.Cyan, common.White, eventsForLaunchpad, guiButtons)

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
		common.LightLamp(common.SAVE_BUTTON, common.White, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)

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
			floodOn(this, commandChannels, eventsForLaunchpad, guiButtons)
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
			floodOff(this, commandChannels, eventsForLaunchpad, guiButtons)
			return
		}
	}

	// Sound sensitity up.
	if X == 4 && Y == -1 {

		if debug {
			fmt.Printf("Sound Up %f\n", this.SoundGain)
		}

		buttonTouched(common.Button{X: X, Y: Y}, common.White, common.Cyan, eventsForLaunchpad, guiButtons)

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

		buttonTouched(common.Button{X: X, Y: Y}, common.White, common.Cyan, eventsForLaunchpad, guiButtons)

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

		buttonTouched(common.Button{X: X, Y: Y}, common.White, common.Cyan, eventsForLaunchpad, guiButtons)

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

		buttonTouched(common.Button{X: X, Y: Y}, common.White, common.Cyan, eventsForLaunchpad, guiButtons)

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
			common.LightLamp(common.SAVE_BUTTON, common.White, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)
			return
		}
		this.SavePreset = true
		if this.Flood { // Turn off flood.
			floodOff(this, commandChannels, eventsForLaunchpad, guiButtons)
		}
		presets.RefreshPresets(eventsForLaunchpad, guiButtons, this.PresetsStore)
		common.FlashLight(common.SAVE_BUTTON, common.Magenta, common.White, eventsForLaunchpad, guiButtons)

		return
	}

	// P R E S E T S
	if X < 8 && (Y > 3 && Y < 7) {

		if debug {
			fmt.Printf("Ask For Config\n")
		}

		if this.ShowRGBColorPicker {
			this.ShowRGBColorPicker = false
			removeColorPicker(this, sequences, eventsForLaunchpad, guiButtons, commandChannels)
		}

		location := fmt.Sprint(X) + "," + fmt.Sprint(Y)

		if this.SavePreset {
			// S A V E - Ask all sequences for their current config and save in a file.

			current := this.PresetsStore[location]
			this.PresetsStore[location] = presets.Preset{State: true, Selected: true, Label: current.Label, ButtonColor: current.ButtonColor}
			this.LastPreset = &location

			config.AskToSaveConfig(commandChannels, replyChannels, X, Y)

			// turn off the save button from flashing.
			common.LightLamp(common.SAVE_BUTTON, common.White, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)

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
					loadConfig(sequences, this, X, Y, commandChannels, eventsForLaunchpad, guiButtons, updateChannels)
					common.StartStaticSequences(sequences, commandChannels)
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

	// S E L E C T   S H I F T - Decrease Shift.
	if X == 2 && Y == 7 && !this.ShowRGBColorPicker {

		if debug {
			fmt.Printf("Decrease Shift\n")
		}

		buttonTouched(common.Button{X: X, Y: Y}, common.White, common.Cyan, eventsForLaunchpad, guiButtons)

		// If we're in shutter chase mode.
		if this.SelectedMode[this.SelectedSequence] == CHASER_FUNCTION || this.SelectedMode[this.SelectedSequence] == CHASER_DISPLAY {
			this.TargetSequence = this.ChaserSequenceNumber
		} else {
			this.TargetSequence = this.SelectedSequence
		}

		// Deal with an RGB sequence.
		if sequences[this.TargetSequence].Type == "rgb" {

			// Decrement the RGB Shift.
			this.RGBShift[this.TargetSequence] = this.RGBShift[this.TargetSequence] - 1
			if this.RGBShift[this.TargetSequence] < 0 {
				this.RGBShift[this.TargetSequence] = 0
			}

			// Send a message to the RGB sequence.
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

		// Deal with an Scanner sequence.
		if sequences[this.TargetSequence].Type == "scanner" {

			// Decrement the Scanner Shift.
			this.ScannerShift[this.TargetSequence] = this.ScannerShift[this.TargetSequence] - 1
			if this.ScannerShift[this.TargetSequence] < common.MIN_SCANNER_SHIFT {
				this.ScannerShift[this.TargetSequence] = common.MIN_SCANNER_SHIFT
			}

			// Send a message to the Scanner sequence.
			cmd := common.Command{
				Action: common.UpdateScannerShift,
				Args: []common.Arg{
					{Name: "ScannerShift", Value: this.ScannerShift[this.TargetSequence]},
				},
			}
			common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

			// Update the status bar
			UpdateShift(this, guiButtons)

			return
		}

		// Deal with an Switch sequence.
		if this.SelectedType == "switch" {

			// Decrement the Switch Shift.
			this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Shift = this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Shift - 1
			if this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Shift < 0 {
				this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Shift = 0
			}

			// Send a message to override / increase the selected switch shift.
			cmd := common.Command{
				Action: common.OverrideSwitchShift,
				Args: []common.Arg{
					{Name: "SwitchNumber", Value: this.SelectedSwitch},
					{Name: "SwitchPosition", Value: this.SwitchPosition[this.SelectedSwitch]},
					{Name: "Shift", Value: this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Shift},
				},
			}
			common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

			// Update the status bar
			UpdateShift(this, guiButtons)

			return
		}
	}

	// S E L E C T   S H I F T - Increase Shift.
	if X == 3 && Y == 7 && !this.ShowRGBColorPicker {

		if debug {
			fmt.Printf("Increase Shift \n")
		}

		buttonTouched(common.Button{X: X, Y: Y}, common.White, common.Cyan, eventsForLaunchpad, guiButtons)

		// If we're in shutter chase mode.
		if this.SelectedMode[this.SelectedSequence] == CHASER_FUNCTION || this.SelectedMode[this.SelectedSequence] == CHASER_DISPLAY {
			this.TargetSequence = this.ChaserSequenceNumber
		} else {
			this.TargetSequence = this.SelectedSequence
		}

		// Deal with an RGB sequence.
		if sequences[this.TargetSequence].Type == "rgb" {

			// Increment the RGB Shift.
			this.RGBShift[this.TargetSequence] = this.RGBShift[this.TargetSequence] + 1
			if this.RGBShift[this.TargetSequence] > common.MAX_RGB_SHIFT {
				this.RGBShift[this.TargetSequence] = common.MAX_RGB_SHIFT
			}

			// Send a message to the RGB sequence.
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

		// Deal with an Scanner sequence.
		if sequences[this.TargetSequence].Type == "scanner" {

			// Increment the Scanner Shift.
			this.ScannerShift[this.TargetSequence] = this.ScannerShift[this.TargetSequence] + 1
			if this.ScannerShift[this.TargetSequence] > common.MAX_SCANNER_SHIFT {
				this.ScannerShift[this.TargetSequence] = common.MAX_SCANNER_SHIFT
			}

			// Send a message to the Scanner sequence.
			cmd := common.Command{
				Action: common.UpdateScannerShift,
				Args: []common.Arg{
					{Name: "ScannerShift", Value: this.ScannerShift[this.TargetSequence]},
				},
			}
			common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

			// Update the status bar
			UpdateShift(this, guiButtons)

			return
		}

		// Deal with an Switch sequence.
		if this.SelectedType == "switch" {

			this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Shift = this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Shift + 1
			if this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Shift > common.MAX_RGB_SHIFT {
				this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Shift = common.MAX_RGB_SHIFT
			}

			// Send a message to override / increase the selected switch shift.
			cmd := common.Command{
				Action: common.OverrideSwitchShift,
				Args: []common.Arg{
					{Name: "SwitchNumber", Value: this.SelectedSwitch},
					{Name: "SwitchPosition", Value: this.SwitchPosition[this.SelectedSwitch]},
					{Name: "Shift", Value: this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Shift},
				},
			}
			common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

			// Update the status bar
			UpdateShift(this, guiButtons)

			return
		}
	}

	// S E L E C T   S P E E D - Decrease speed of selected sequence.
	if X == 0 && Y == 7 && !this.ShowRGBColorPicker {

		if debug {
			fmt.Printf("Decrease Speed \n")
		}

		buttonTouched(common.Button{X: X, Y: Y}, common.White, common.Cyan, eventsForLaunchpad, guiButtons)

		// If we're in shutter chase mode.
		if this.SelectedMode[this.SelectedSequence] == CHASER_FUNCTION || this.SelectedMode[this.SelectedSequence] == CHASER_DISPLAY {
			this.TargetSequence = this.ChaserSequenceNumber
		} else {
			this.TargetSequence = this.SelectedSequence
		}

		// Strobe only every operates on the selected sequence, i.e chaser never applies strobe.
		// Decrease Strobe Speed.
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
			if this.SelectedType == "scanner" && this.ScannerChaser[this.SelectedSequence] {
				common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)
			}

			// Update the status bar
			common.UpdateStatusBar(fmt.Sprintf("Strobe %02d", this.StrobeSpeed[this.SelectedSequence]), "speed", false, guiButtons)
			return
		}

		// Get an upto date copy of the target sequence.
		sequences[this.TargetSequence] = common.RefreshSequence(this.TargetSequence, commandChannels, updateChannels)

		if this.SelectedType == "switch" {
			// Copy the updated speed setting into the local switch speed storage
			this.Speed[this.TargetSequence] = this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Speed
		}

		// Decrease Speed.
		if !sequences[this.TargetSequence].MusicTrigger {
			this.Speed[this.TargetSequence]--
			if this.Speed[this.TargetSequence] < 1 {
				this.Speed[this.TargetSequence] = 1
			}

			// If you reached the min speed blink the increase button.
			if this.Speed[this.TargetSequence] == common.MIN_SPEED {
				common.FlashLight(common.Button{X: X, Y: Y}, common.White, common.Cyan, eventsForLaunchpad, guiButtons)
			} else {
				// If you reached the half speed blink both buttons.
				if this.Speed[this.TargetSequence] == common.MAX_SPEED/2 {
					common.FlashLight(common.Button{X: X, Y: Y}, common.White, common.Cyan, eventsForLaunchpad, guiButtons)
					common.FlashLight(common.Button{X: X + 1, Y: Y}, common.White, common.Cyan, eventsForLaunchpad, guiButtons)
				} else {
					common.LightLamp(common.Button{X: X + 1, Y: Y}, common.Cyan, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)
				}
			}

			if this.SelectedType == "switch" {
				// Copy the updated speed setting into the local switch speed storage
				this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Speed = this.Speed[this.TargetSequence]
				// Send a message to override / decrease the selected switch speed.
				cmd := common.Command{
					Action: common.OverrideSwitchSpeed,
					Args: []common.Arg{
						{Name: "SwitchNumber", Value: this.SelectedSwitch},
						{Name: "SwitchPosition", Value: this.SwitchPosition[this.SelectedSwitch]},
						{Name: "Speed", Value: this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Speed},
					},
				}
				common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)
			} else {
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

		buttonTouched(common.Button{X: X, Y: Y}, common.White, common.Cyan, eventsForLaunchpad, guiButtons)

		// If we're in shutter chase mode
		if this.SelectedMode[this.SelectedSequence] == CHASER_FUNCTION || this.SelectedMode[this.SelectedSequence] == CHASER_DISPLAY {
			this.TargetSequence = this.ChaserSequenceNumber
		} else {
			this.TargetSequence = this.SelectedSequence
		}

		// Strobe only every operates on the selected sequence, i.e chaser never applies strobe.
		// Increase Strobe Speed.
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
			if this.SelectedType == "scanner" && this.ScannerChaser[this.SelectedSequence] {
				common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)
			}

			// Update the status bar
			common.UpdateStatusBar(fmt.Sprintf("Strobe %02d", this.StrobeSpeed[this.SelectedSequence]), "speed", false, guiButtons)
			return
		}

		// Get an upto date copy of the sequence.
		sequences[this.TargetSequence] = common.RefreshSequence(this.TargetSequence, commandChannels, updateChannels)

		if this.SelectedType == "switch" {
			// Copy the updated speed setting into the local switch speed storage
			this.Speed[this.TargetSequence] = this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Speed
		}

		if !sequences[this.TargetSequence].MusicTrigger {
			this.Speed[this.TargetSequence]++
			if this.Speed[this.TargetSequence] > 12 {
				this.Speed[this.TargetSequence] = 12
			}

			// If you reached the max speed blink the increase button.
			if this.Speed[this.TargetSequence] == common.MAX_SPEED {
				common.FlashLight(common.Button{X: X, Y: Y}, common.White, common.Cyan, eventsForLaunchpad, guiButtons)
			} else {
				// If you reached the half speed blink both buttons.
				if this.Speed[this.TargetSequence] == common.MAX_SPEED/2 {
					common.FlashLight(common.Button{X: X, Y: Y}, common.White, common.Cyan, eventsForLaunchpad, guiButtons)
					common.FlashLight(common.Button{X: X - 1, Y: Y}, common.White, common.Cyan, eventsForLaunchpad, guiButtons)
				} else {
					common.LightLamp(common.Button{X: X - 1, Y: Y}, common.Cyan, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)
				}
			}
			if this.SelectedType == "switch" {
				// Copy the speed setting into the local switch speed storage
				this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Speed = this.Speed[this.TargetSequence]
				// Send a message to override / increase the selected switch speed.
				cmd := common.Command{
					Action: common.OverrideSwitchSpeed,
					Args: []common.Arg{
						{Name: "SwitchNumber", Value: this.SelectedSwitch},
						{Name: "SwitchPosition", Value: this.SwitchPosition[this.SelectedSwitch]},
						{Name: "Speed", Value: this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Speed},
					},
				}
				common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)
			} else {
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
		}

		// Update the status bar
		UpdateSpeed(this, guiButtons)

		return
	}

	// S E L E C T   S E Q U E N C E.
	// Select sequence 1.
	if X == 8 && Y == 0 {

		this.SelectedSequence = 0
		this.SelectedType = sequences[this.SelectedSequence].Type

		if debug {
			fmt.Printf("Select Sequence %d Type %s\n", this.SelectedSequence, this.SelectedType)
		}
		deFocusAllSwitches(this, sequences, commandChannels)

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

		deFocusAllSwitches(this, sequences, commandChannels)
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

		deFocusAllSwitches(this, sequences, commandChannels)
		HandleSelect(sequences, this, eventsForLaunchpad, commandChannels, guiButtons)

		this.ShowRGBColorPicker = false
		this.EditGoboSelectionMode = false
		if this.ScannerChaser[this.SelectedSequence] {
			this.EditWhichStaticSequence = 4
		}

		return
	}

	// S T A R T - Start sequence.
	if X == 8 && Y == 5 {

		if this.ShowRGBColorPicker {
			this.ShowRGBColorPicker = false
			removeColorPicker(this, sequences, eventsForLaunchpad, guiButtons, commandChannels)
		}

		// Start in normal mode, hide the shutter chaser.
		if this.SelectedType == "scanner" && this.ScannerChaser[this.SelectedSequence] {
			common.HideSequence(this.ChaserSequenceNumber, commandChannels)
			common.ClearSelectedRowOfButtons(this.SelectedSequence, eventsForLaunchpad, guiButtons)
			this.SelectedMode[this.SelectedSequence] = NORMAL
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

			if this.Strobe[this.SelectedSequence] {
				this.Strobe[this.SelectedSequence] = false
				StopStrobe(this, eventsForLaunchpad, guiButtons, commandChannels)
			}

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
				this.SelectedMode[this.SelectedSequence] = NORMAL
				this.Running[this.ChaserSequenceNumber] = false
			}

			// Clear the pattern function keys
			common.ClearSelectedRowOfButtons(this.SelectedSequence, eventsForLaunchpad, guiButtons)

			// Turn off the start lamp.
			common.LightLamp(common.Button{X: X, Y: Y}, common.White, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)

			// Set the correct color for the select button.
			lightSelectedButton(eventsForLaunchpad, guiButtons, this)

			return

		} else {
			// Start this sequence.
			if debug {
				fmt.Printf("Start Sequence %d \n", Y)
			}

			// Stop the music trigger.
			sequences[this.SelectedSequence].MusicTrigger = false

			// If strobing stop it.
			if this.Strobe[this.SelectedSequence] {
				this.Strobe[this.SelectedSequence] = false
				StopStrobe(this, eventsForLaunchpad, guiButtons, commandChannels)
			}

			// Start the sequence.
			cmd := common.Command{
				Action: common.Start,
				Args: []common.Arg{
					{Name: "Speed", Value: this.Speed[this.SelectedSequence]},
				},
			}
			common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)
			common.LightLamp(common.Button{X: X, Y: Y}, common.Green, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)

			this.Running[this.SelectedSequence] = true
			this.Functions[this.SelectedSequence][common.Function6_Static_Gobo].State = false
			this.Functions[this.SelectedSequence][common.Function8_Music_Trigger].State = false

			// Clear the pattern function keys
			common.ClearSelectedRowOfButtons(this.SelectedSequence, eventsForLaunchpad, guiButtons)

			// Reveal the now running sequence
			common.RevealSequence(this.SelectedSequence, commandChannels)

			// Set the correct color for the select button.
			lightSelectedButton(eventsForLaunchpad, guiButtons, this)

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
		common.LightLamp(common.SAVE_BUTTON, common.White, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)

		// Shutdown any function bars.
		clearAllModes(sequences, this)
		HandleSelect(sequences, this, eventsForLaunchpad, commandChannels, guiButtons)

		// If strobing, stop it
		if this.Strobe[this.SelectedSequence] {
			this.Strobe[this.SelectedSequence] = false
			// Stop strobing this sequence.
			StopStrobe(this, eventsForLaunchpad, guiButtons, commandChannels)
			return

		} else {
			// Start strobing for this sequence. Strobe on.
			this.Strobe[this.SelectedSequence] = true
			StartStrobe(this, eventsForLaunchpad, guiButtons, commandChannels)
			return
		}
	}

	// Size decrease.
	if X == 4 && Y == 7 && !this.ShowRGBColorPicker {

		if debug {
			fmt.Printf("Decrease Size\n")
		}

		buttonTouched(common.Button{X: X, Y: Y}, common.White, common.Cyan, eventsForLaunchpad, guiButtons)

		// If we're a scanner and we're in shutter chase mode.
		if sequences[this.SelectedSequence].Type == "scanner" && this.ScannerChaser[this.SelectedSequence] &&
			(this.SelectedMode[this.SelectedSequence] == CHASER_FUNCTION || this.SelectedMode[this.SelectedSequence] == CHASER_DISPLAY) {
			this.TargetSequence = this.ChaserSequenceNumber
		} else {
			this.TargetSequence = this.SelectedSequence
		}

		// Deal with the RGB sequence.
		if sequences[this.TargetSequence].Type == "rgb" {

			// Decrement RGB Size.
			this.RGBSize[this.TargetSequence]--
			if this.RGBSize[this.TargetSequence] < common.MIN_RGB_SIZE {
				this.RGBSize[this.TargetSequence] = common.MIN_RGB_SIZE
			}

			// Send Update RGB Size.
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

		// Deal with Scanner sequence.
		if sequences[this.TargetSequence].Type == "scanner" {

			// Send Update Scanner Size.
			this.ScannerSize[this.TargetSequence] = this.ScannerSize[this.TargetSequence] - 10
			if this.ScannerSize[this.TargetSequence] < 0 {
				this.ScannerSize[this.TargetSequence] = 0
			}

			// Send Update Scanner Size.
			cmd := common.Command{
				Action: common.UpdateScannerSize,
				Args: []common.Arg{
					{Name: "ScannerSize", Value: this.ScannerSize[this.TargetSequence]},
				},
			}
			common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

			// Update the status bar
			UpdateSize(this, guiButtons)

			return
		}

		// Deal with the switch sequence.
		if sequences[this.TargetSequence].Type == "switch" {

			// Decrement the switch size.
			this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Size--
			if this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Size < common.MIN_RGB_SIZE {
				this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Size = common.MIN_RGB_SIZE
			}

			// Send a message to override / increase the selected switch shift.
			cmd := common.Command{
				Action: common.OverrideSwitchSize,
				Args: []common.Arg{
					{Name: "SwitchNumber", Value: this.SelectedSwitch},
					{Name: "SwitchPosition", Value: this.SwitchPosition[this.SelectedSwitch]},
					{Name: "Shift", Value: this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Size},
				},
			}
			common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

			// Update the status bar.
			UpdateSize(this, guiButtons)

			return
		}
	}

	// Increase Size.
	if X == 5 && Y == 7 && !this.ShowRGBColorPicker {

		if debug {
			fmt.Printf("Increase Size\n")
		}

		buttonTouched(common.Button{X: X, Y: Y}, common.White, common.Cyan, eventsForLaunchpad, guiButtons)

		// If we're a scanner and we're in shutter chase mode.
		if sequences[this.SelectedSequence].Type == "scanner" && this.ScannerChaser[this.SelectedSequence] &&
			(this.SelectedMode[this.SelectedSequence] == CHASER_FUNCTION || this.SelectedMode[this.SelectedSequence] == CHASER_DISPLAY) {
			this.TargetSequence = this.ChaserSequenceNumber
		} else {
			this.TargetSequence = this.SelectedSequence
		}

		// Deal with the RGB sequence.
		if sequences[this.TargetSequence].Type == "rgb" {

			// Send Update RGB Size.
			this.RGBSize[this.TargetSequence]++
			if this.RGBSize[this.TargetSequence] > common.MAX_RGB_SIZE {
				this.RGBSize[this.TargetSequence] = common.MAX_RGB_SIZE
			}

			// Send a message to the RGB sequence.
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

		// Deal with the Scanner size.
		if sequences[this.TargetSequence].Type == "scanner" {

			// Increment the scanner size.
			this.ScannerSize[this.TargetSequence] = this.ScannerSize[this.TargetSequence] + 10
			if this.ScannerSize[this.TargetSequence] > common.MAX_SCANNER_SIZE {
				this.ScannerSize[this.TargetSequence] = common.MAX_SCANNER_SIZE
			}

			// Send Update Scanner Size.
			cmd := common.Command{
				Action: common.UpdateScannerSize,
				Args: []common.Arg{
					{Name: "ScannerSize", Value: this.ScannerSize[this.TargetSequence]},
				},
			}
			common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

			// Update the status bar
			UpdateSize(this, guiButtons)

			return
		}

		// Deal with an Switch sequence.
		if this.SelectedType == "switch" {

			// Increase the switch size.
			this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Size = this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Size + 1
			if this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Size > common.MAX_RGB_SHIFT {
				this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Size = common.MAX_RGB_SHIFT
			}

			// Send a message to override / increase the selected switch shift.
			cmd := common.Command{
				Action: common.OverrideSwitchSize,
				Args: []common.Arg{
					{Name: "SwitchNumber", Value: this.SelectedSwitch},
					{Name: "SwitchPosition", Value: this.SwitchPosition[this.SelectedSwitch]},
					{Name: "Shift", Value: this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Size},
				},
			}
			common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

			// Update the status bar
			UpdateSize(this, guiButtons)

			return
		}
	}

	// Fade time decrease.
	if X == 6 && Y == 7 && !this.ShowRGBColorPicker {

		if debug {
			fmt.Printf("Decrease Fade Time Type=%s Sequence=%d Type=%s\n", this.SelectedType, this.TargetSequence, sequences[this.TargetSequence].Type)
		}

		buttonTouched(common.Button{X: X, Y: Y}, common.White, common.Cyan, eventsForLaunchpad, guiButtons)

		// If we're a scanner and we're in shutter chase mode.
		if sequences[this.SelectedSequence].Type == "scanner" &&
			this.ScannerChaser[this.SelectedSequence] &&
			(this.SelectedMode[this.SelectedSequence] == CHASER_FUNCTION || this.SelectedMode[this.SelectedSequence] == CHASER_DISPLAY) {
			this.TargetSequence = this.ChaserSequenceNumber
		} else {
			this.TargetSequence = this.SelectedSequence
		}

		if sequences[this.TargetSequence].Type == "rgb" || sequences[this.TargetSequence].Label == "chaser" {

			// Decrement the RGB Fade size.
			this.RGBFade[this.TargetSequence]--
			if this.RGBFade[this.TargetSequence] < 1 {
				this.RGBFade[this.TargetSequence] = 1
			}

			// Send fade update command to sequence.
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
		if sequences[this.TargetSequence].Type == "scanner" && sequences[this.TargetSequence].Label != "chaser" {

			// Fade also send more or less coordinates for the scanner patterns.
			this.ScannerCoordinates[this.TargetSequence]--
			if this.ScannerCoordinates[this.TargetSequence] < 0 {
				this.ScannerCoordinates[this.TargetSequence] = 0
			}

			// Send a messages to the scanner sequence.
			cmd := common.Command{
				Action: common.UpdateNumberCoordinates,
				Args: []common.Arg{
					{Name: "NumberCoordinates", Value: this.ScannerCoordinates[this.TargetSequence]},
				},
			}
			common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

			// Update the status bar
			UpdateFade(this, guiButtons)

			return
		}

		// Deal with an Switch sequence.
		if this.SelectedType == "switch" {

			// Decrease the fade size.
			this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Fade--
			if this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Fade < common.MIN_RGB_FADE {
				this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Fade = common.MIN_RGB_FADE
			}

			// Send a message to override / increase the selected switch shift.
			cmd := common.Command{
				Action: common.OverrideSwitchFade,
				Args: []common.Arg{
					{Name: "SwitchNumber", Value: this.SelectedSwitch},
					{Name: "SwitchPosition", Value: this.SwitchPosition[this.SelectedSwitch]},
					{Name: "Shift", Value: this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Fade},
				},
			}
			common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

			// Update the status bar
			UpdateFade(this, guiButtons)

			return
		}

	}

	// Fade time increase.
	if X == 7 && Y == 7 && !this.ShowRGBColorPicker {

		if debug {
			fmt.Printf("Increase Fade Time\n")
		}

		buttonTouched(common.Button{X: X, Y: Y}, common.White, common.Cyan, eventsForLaunchpad, guiButtons)

		// If we're a scanner and we're in shutter chase mode.
		if sequences[this.SelectedSequence].Type == "scanner" && this.ScannerChaser[this.SelectedSequence] &&
			(this.SelectedMode[this.SelectedSequence] == CHASER_FUNCTION || this.SelectedMode[this.SelectedSequence] == CHASER_DISPLAY) {
			this.TargetSequence = this.ChaserSequenceNumber
		} else {
			this.TargetSequence = this.SelectedSequence
		}

		// Deal with the RGB sequence.
		if sequences[this.TargetSequence].Type == "rgb" || sequences[this.TargetSequence].Label == "chaser" {

			// Increase fade time.
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

		// Deal wth scanner coordinates.
		if sequences[this.TargetSequence].Type == "scanner" && sequences[this.TargetSequence].Label != "chaser" {

			// Fade also send more or less coordinates for the scanner patterns.
			this.ScannerCoordinates[this.TargetSequence]++
			if this.ScannerCoordinates[this.TargetSequence] > 4 {
				this.ScannerCoordinates[this.TargetSequence] = 4
			}

			// Send a message to scanner seqiemce.
			cmd := common.Command{
				Action: common.UpdateNumberCoordinates,
				Args: []common.Arg{
					{Name: "NumberCoordinates", Value: this.ScannerCoordinates[this.TargetSequence]},
				},
			}
			common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

			// Update the status bar
			UpdateFade(this, guiButtons)

			return
		}

		// Deal with an Switch sequence.
		if this.SelectedType == "switch" {

			// Increase the switch size.
			this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Fade = this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Fade + 1
			if this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Fade > common.MAX_RGB_SHIFT {
				this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Fade = common.MAX_RGB_SHIFT
			}

			// Send a message to override / increase the selected switch shift.
			cmd := common.Command{
				Action: common.OverrideSwitchFade,
				Args: []common.Arg{
					{Name: "SwitchNumber", Value: this.SelectedSwitch},
					{Name: "SwitchPosition", Value: this.SwitchPosition[this.SelectedSwitch]},
					{Name: "Shift", Value: this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Fade},
				},
			}
			common.SendCommandToSequence(this.TargetSequence, cmd, commandChannels)

			// Update the status bar
			UpdateFade(this, guiButtons)

			return
		}

	}

	// S E L E C T   S W I T C H   B U T T O N's Toggle State of switches for this sequence.
	if X >= 0 && X < 8 &&
		Y >= 0 &&
		Y < 4 &&
		sequences[Y].Type == "switch" {

		this.SelectedSequence = Y
		this.SelectedSwitch = X
		this.SelectedType = "switch"
		this.SelectedFixtureType = fixture.GetSwitchFixtureType(this.SelectedSwitch, int16(this.SwitchPosition[this.SelectedSwitch]), fixturesConfig)

		if debug {
			fmt.Printf("Switch Key X:%d Y:%d\n", this.SelectedSwitch, this.SelectedSequence)
		}

		if this.ShowRGBColorPicker {
			this.ShowRGBColorPicker = false
			removeColorPicker(this, sequences, eventsForLaunchpad, guiButtons, commandChannels)
		}

		// Get an upto date copy of the switch information by updating our copy of the switch sequence.
		sequences[this.SelectedSequence] = common.RefreshSequence(this.SelectedSequence, commandChannels, updateChannels)

		// We have a valid switch.
		if this.SelectedSwitch < len(sequences[this.SelectedSequence].Switches) {

			// Second time we've pressed this switch button, actually step the state.
			if this.SelectedSwitch == this.LastSelectedSwitch {
				this.SwitchPosition[this.SelectedSwitch] = this.SwitchPosition[this.SelectedSwitch] + 1
				valuesLength := len(sequences[this.SelectedSequence].Switches[this.SelectedSwitch].States)
				if this.SwitchPosition[this.SelectedSwitch] == valuesLength {
					this.SwitchPosition[this.SelectedSwitch] = 0
				}
				// Send a message to the sequence for it to step to the next state the selected switch.
				cmd := common.Command{
					Action: common.UpdateSwitch,
					Args: []common.Arg{
						{Name: "SwitchNumber", Value: this.SelectedSwitch},
						{Name: "SwitchPosition", Value: this.SwitchPosition[this.SelectedSwitch]},
						{Name: "Step", Value: true},  // Step the switch state.
						{Name: "Focus", Value: true}, // Focus the switch lamp.
					},
				}
				// Send a message to the switch sequence.
				common.SendCommandToAllSequenceOfType(sequences, cmd, commandChannels, "switch")
			} else {
				// Just send a message to focus the switch button.
				cmd := common.Command{
					Action: common.UpdateSwitch,
					Args: []common.Arg{
						{Name: "SwitchNumber", Value: this.SelectedSwitch},
						{Name: "SwitchPosition", Value: this.SwitchPosition[this.SelectedSwitch]},
						{Name: "Step", Value: false}, // Don't step the switch state.
						{Name: "Focus", Value: true}, // Focus the switch lamp.
					},
				}
				// Send a message to the switch sequence.
				common.SendCommandToAllSequenceOfType(sequences, cmd, commandChannels, "switch")

			}

		}

		// Light the correct selected switch.
		this.SelectedSequence = this.SwitchSequenceNumber
		this.LastSelectedSwitch = this.SelectedSwitch

		// Use the default behaviour of SelectSequence to turn of the other sequence select buttons.
		SelectSequence(this)

		// Find out if this switch state has a music trigger.
		this.MusicTrigger = fixture.GetSwitchStateIsMusicTriggerOn(this.SelectedSwitch, int16(this.SwitchPosition[this.SelectedSwitch]), fixturesConfig)

		// Switch overrides will get displayed here as well.
		UpdateSpeed(this, guiButtons)
		UpdateShift(this, guiButtons)
		UpdateSize(this, guiButtons)
		UpdateFade(this, guiButtons)

		// Update the labels.
		showStatusBars(this, sequences, eventsForLaunchpad, guiButtons)

		// Light the sequence selector button.
		lightSelectedButton(eventsForLaunchpad, guiButtons, this)
	}

	// D I S A B L E  / E N A B L E   F I X T U R E  S T A T U S - Used to toggle the scanner state from on, inverted or off.
	if X >= 0 && X < 8 &&
		Y >= 0 &&
		Y < 4 &&
		this.SelectedMode[this.SelectedSequence] == STATUS {

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

		buttonTouched(common.Button{X: X, Y: Y}, common.Cyan, common.White, eventsForLaunchpad, guiButtons)

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

		buttonTouched(common.Button{X: X, Y: Y}, common.Cyan, common.White, eventsForLaunchpad, guiButtons)

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

		buttonTouched(common.Button{X: X, Y: Y}, common.Cyan, common.White, eventsForLaunchpad, guiButtons)

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

	// S E L E C T   E D I T  S T A T I C   C O L O R
	// Red
	if X == 1 && Y == -1 && this.Static[this.TargetSequence] {

		if this.Functions[this.SelectedSequence][common.Function6_Static_Gobo].State {

			if debug {
				fmt.Printf("Choose Static Red X:%d Y:%d\n", X, Y)
			}

			buttonTouched(common.Button{X: X, Y: Y}, common.White, common.Red, eventsForLaunchpad, guiButtons)

			this.StaticButtons[this.SelectedSequence].X = this.LastStaticColorButtonX
			this.StaticButtons[this.SelectedSequence].Y = this.LastStaticColorButtonY
			this.StaticButtons[this.SelectedSequence].Color.R = this.StaticButtons[this.SelectedSequence].Color.R - 10
			if this.StaticButtons[this.SelectedSequence].Color.R == 254 {
				this.StaticButtons[this.SelectedSequence].Color.R = 0
			}
			if this.StaticButtons[this.SelectedSequence].Color.R == 0 {
				this.StaticButtons[this.SelectedSequence].Color.R = 254
			}

			redColor := color.RGBA{R: this.StaticButtons[this.SelectedSequence].Color.R, G: 0, B: 0}
			common.LightLamp(common.RED_BUTTON, redColor, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)
			updateStaticLamp(this.SelectedSequence, this.StaticButtons[this.SelectedSequence], commandChannels)

			// Update the status bar
			common.UpdateStatusBar(fmt.Sprintf("Red %02d", this.StaticButtons[this.SelectedSequence].Color.R), "red", false, guiButtons)
			return
		}
	}

	// Green
	if X == 2 && Y == -1 && this.Static[this.TargetSequence] {

		if this.Functions[this.SelectedSequence][common.Function6_Static_Gobo].State {

			if debug {
				fmt.Printf("Choose Static Green X:%d Y:%d\n", X, Y)
			}

			buttonTouched(common.Button{X: X, Y: Y}, common.White, common.Green, eventsForLaunchpad, guiButtons)

			this.StaticButtons[this.SelectedSequence].X = this.LastStaticColorButtonX
			this.StaticButtons[this.SelectedSequence].Y = this.LastStaticColorButtonY
			this.StaticButtons[this.SelectedSequence].Color.G = this.StaticButtons[this.SelectedSequence].Color.G - 10
			if this.StaticButtons[this.SelectedSequence].Color.G == 254 {
				this.StaticButtons[this.SelectedSequence].Color.G = 0
			}
			if this.StaticButtons[this.SelectedSequence].Color.G == 0 {
				this.StaticButtons[this.SelectedSequence].Color.G = 254
			}
			greenColor := color.RGBA{R: 0, G: this.StaticButtons[this.SelectedSequence].Color.G, B: 0}
			common.LightLamp(common.Button{X: X, Y: Y}, greenColor, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)
			updateStaticLamp(this.SelectedSequence, this.StaticButtons[this.SelectedSequence], commandChannels)

			// Update the status bar
			common.UpdateStatusBar(fmt.Sprintf("Green %02d", this.StaticButtons[this.SelectedSequence].Color.G), "green", false, guiButtons)
			return
		}
	}

	// Blue
	if X == 3 && Y == -1 && this.Static[this.TargetSequence] {

		if this.Functions[this.SelectedSequence][common.Function6_Static_Gobo].State {

			if debug {
				fmt.Printf("Choose Static Blue X:%d Y:%d\n", X, Y)
			}

			buttonTouched(common.Button{X: X, Y: Y}, common.White, common.Blue, eventsForLaunchpad, guiButtons)

			this.StaticButtons[this.SelectedSequence].X = this.LastStaticColorButtonX
			this.StaticButtons[this.SelectedSequence].Y = this.LastStaticColorButtonY
			this.StaticButtons[this.SelectedSequence].Color.B = this.StaticButtons[this.SelectedSequence].Color.B - 10
			if this.StaticButtons[this.SelectedSequence].Color.B > 254 {
				this.StaticButtons[this.SelectedSequence].Color.B = 0
			}
			if this.StaticButtons[this.SelectedSequence].Color.B == 0 {
				this.StaticButtons[this.SelectedSequence].Color.B = 254
			}
			blueColor := color.RGBA{R: 0, G: 0, B: this.StaticButtons[this.SelectedSequence].Color.B}
			common.LightLamp(common.Button{X: X, Y: Y}, blueColor, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)
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

		if this.SelectedMode[this.SelectedSequence] == CHASER_FUNCTION {
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
		ShowRGBColorPicker(*sequences[this.TargetSequence], this.DisplaySequence, eventsForLaunchpad, guiButtons, commandChannels)

		return
	}

	// S E L E C T   S C A N N E R   C O L O R
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
		fixture.MapFixturesColorOnly(this.SelectedSequence, this.SelectedFixture, this.ScannerColor, dmxController, fixturesConfig, this.DmxInterfacePresent)

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
		fixture.MapFixturesGoboOnly(this.SelectedSequence, this.SelectedFixture, this.SelectedGobo, fixturesConfig, dmxController, this.DmxInterfacePresent)

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
		(this.SelectedMode[this.SelectedSequence] == NORMAL ||
			this.SelectedMode[this.SelectedSequence] == NORMAL_STATIC ||
			this.SelectedMode[this.SelectedSequence] == CHASER_DISPLAY ||
			this.SelectedMode[this.SelectedSequence] == CHASER_DISPLAY_STATIC) && // Not in function Mode
		!this.ShowStaticColorPicker && // Not In Color Picker Mode.
		getStatic(this) { // Static Function On in any sequence

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
		if color == common.EmptyColor {
			color = FindCurrentColor(this.SelectedStaticFixtureNumber, this.SelectedSequence, *sequences[this.TargetSequence])
		}

		if debug {
			fmt.Printf("Sequence %d Fixture %d Setting Current Color as %+v\n", this.SelectedSequence, this.SelectedStaticFixtureNumber, color)
		}

		// Set the fixture color so that it flashs in the color picker.
		sequences[this.TargetSequence].CurrentColors = SetRGBColorPicker(color, *sequences[this.TargetSequence])

		// We call ShowRGBColorPicker so you can choose the static color for this fixture.
		ShowRGBColorPicker(*sequences[this.TargetSequence], this.DisplaySequence, eventsForLaunchpad, guiButtons, commandChannels)

		// Switch the mode so we know we are picking a static color from the color picker.
		this.ShowStaticColorPicker = true
		this.Static[this.TargetSequence] = true

		return
	}

	// S E L E C T   S T A T I C   C O L O R
	if X >= 0 && X < 8 && Y != -1 &&
		Y < 3 && // Make sure the buttons pressed inside the color picker.
		this.ShowStaticColorPicker && // Now We Are In Static Color Picker Mode.
		!this.EditFixtureSelectionMode && // Not In Fixture Selection Mode.
		getStatic(this) { // Static Function On in this or shutter chaser sequence

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

		// Save the color in local copy of the static color button
		this.StaticButtons[this.SelectedStaticFixtureNumber].Color = color

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

			// Save the color to all static fixtures in our local copy.
			for buttonNumber, button := range this.StaticButtons {
				this.StaticButtons[buttonNumber].Color = button.Color
			}

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

			// Save the color in local copy of the static color button
			this.StaticButtons[this.SelectedStaticFixtureNumber].Color = sequences[this.TargetSequence].StaticColors[this.SelectedStaticFixtureNumber].Color
		}

		this.LastStaticColorButtonX = this.SelectedStaticFixtureNumber
		this.LastStaticColorButtonY = Y

		// Hide the sequence.
		common.HideSequence(this.TargetSequence, commandChannels)

		// We call ShowRGBColorPicker so you can see which static color has been selected for this fixture.
		ShowRGBColorPicker(*sequences[this.TargetSequence], this.DisplaySequence, eventsForLaunchpad, guiButtons, commandChannels)

		// Set the first pressed for only this fixture and cancel any others
		for x := 0; x < 8; x++ {
			sequences[this.TargetSequence].StaticColors[x].FirstPress = false
		}
		sequences[this.TargetSequence].StaticColors[this.SelectedStaticFixtureNumber].FirstPress = true

		// Remove the color picker and reveal the sequence.
		removeColorPicker(this, sequences, eventsForLaunchpad, guiButtons, commandChannels)

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

	// S E L E C T   P A T T E R N
	if X >= 0 && X < 8 && Y != -1 &&
		!this.EditFixtureSelectionMode &&
		this.EditPatternMode {

		if this.SelectedMode[this.SelectedSequence] == CHASER_FUNCTION {
			this.TargetSequence = this.ChaserSequenceNumber
			this.DisplaySequence = this.SelectedSequence
			// This is the way we get the shutter chaser to be displayed as we exit
			// the pattern selection
			this.DisplayChaserShortCut = true
			this.SelectedMode[this.DisplaySequence] = CHASER_FUNCTION
		} else {
			this.TargetSequence = this.SelectedSequence
			this.DisplaySequence = this.SelectedSequence
			this.SelectedMode[this.DisplaySequence] = NORMAL
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
		// Update the labels.
		showStatusBars(this, sequences, eventsForLaunchpad, guiButtons)
		return
	}

	// F U N C T I O N  K E Y S
	if X >= 0 && X < 8 && Y >= 0 && Y < 3 &&
		this.SelectedSequence == Y && // Make sure the buttons pressed are for this sequence.
		this.SelectedMode[this.SelectedSequence] == FUNCTION || this.SelectedMode[this.SelectedSequence] == CHASER_FUNCTION {
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
		common.LightLamp(common.SAVE_BUTTON, common.White, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)

		if !this.Blackout {
			this.Blackout = true
			cmd := common.Command{
				Action: common.Blackout,
			}
			common.SendCommandToAllSequence(cmd, commandChannels)
			common.LightLamp(common.Button{X: X, Y: Y}, common.Black, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)
			common.FlashLight(common.BLACKOUT_BUTTON, common.Magenta, common.White, eventsForLaunchpad, guiButtons)
		} else {
			this.Blackout = false
			cmd := common.Command{
				Action: common.Normal,
			}
			common.SendCommandToAllSequence(cmd, commandChannels)
			common.LightLamp(common.Button{X: X, Y: Y}, common.White, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)
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
			{Name: "StaticColor", Value: color.RGBA{R: staticColorButtons.Color.R, G: staticColorButtons.Color.G, B: staticColorButtons.Color.B}},
		},
	}
	common.SendCommandToSequence(selectedSequence, cmd, commandChannels)

}

func AllFixturesOff(sequences []*common.Sequence, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, dmxController *ft232.DMXController, fixturesConfig *fixture.Fixtures, dmxInterfacePresent bool) {

	if debug {
		fmt.Printf("AllFixturesOff\n")
	}

	for y := 0; y < len(sequences); y++ {
		if sequences[y].Type != "switch" && sequences[y].Label != "chaser" {
			for x := 0; x < 8; x++ {
				common.LightLamp(common.Button{X: x, Y: y}, common.Black, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)
				fixture.MapFixtures(false, false, y, x, common.Black, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, true, 0, 0, 0, false, 0, dmxController, dmxInterfacePresent)
				common.LabelButton(x, y, "", guiButtons)
			}
		}
	}
}

func SetRGBColorPicker(selectedColor color.RGBA, targetSequence common.Sequence) []color.RGBA {

	if debug {
		fmt.Printf("SetRGBColorPicker\n")
	}

	// Clear out exiting colors.
	targetSequence.CurrentColors = []color.RGBA{}

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

func FindCurrentColor(X int, Y int, targetSequence common.Sequence) color.RGBA {

	if debug {
		fmt.Printf("FindCurrentColor\n")
	}

	for _, availableColor := range targetSequence.RGBAvailableColors {
		if availableColor.X == X && availableColor.Y == Y {
			return availableColor.Color
		}
	}

	return color.RGBA{}
}

// For the given sequence show the available sequence colors on the relevant buttons.
// With the new color picker there can be 24 colors displayed.
// ShowRGBColorPicker operates on the sequence.RGBAvailableColors which is an array of type []common.StaticColorButton
// the targetSequence .CurrentColors selects which colors are selected.
// Returns the RGBAvailableColors []common.StaticColorButton
func ShowRGBColorPicker(targetSequence common.Sequence, displaySequence int, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command) {

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
			Black := common.Black
			if debug {
				fmt.Printf("FLASH myFixtureNumber X:%d Y:%d Color %+v \n", lamp.X, lamp.Y, lamp.Color)
			}
			common.FlashLight(common.Button{X: lamp.X, Y: lamp.Y}, lamp.Color, Black, eventsForLaunchpad, guiButtons)
		} else {
			common.LightLamp(common.Button{X: lamp.X, Y: lamp.Y}, lamp.Color, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)
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
			White := common.White
			common.FlashLight(common.Button{X: fixtureNumber, Y: displaySequence}, fixture.Color, White, eventsForLaunchpad, guiButtons)
		} else {
			common.LightLamp(common.Button{X: fixtureNumber, Y: displaySequence}, fixture.Color, targetSequence.Master, eventsForLaunchpad, guiButtons)
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
			Black := common.Black
			common.FlashLight(common.Button{X: goboNumber, Y: this.SelectedSequence}, gobo.Color, Black, eventsForLaunchpad, guiButtons)
		} else {
			common.LightLamp(common.Button{X: goboNumber, Y: this.SelectedSequence}, gobo.Color, sequence.Master, eventsForLaunchpad, guiButtons)
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
				common.LightLamp(common.Button{X: fixture.Number - 1, Y: this.SelectedSequence}, common.White, sequence.Master, eventsForLaunchpad, guiButtons)
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
			Black := common.Black
			common.FlashLight(common.Button{X: fixtureNumber, Y: this.SelectedSequence}, lamp.Color, Black, eventsForLaunchpad, guiButtons)
		} else {
			common.LightLamp(common.Button{X: fixtureNumber, Y: this.SelectedSequence}, lamp.Color, sequence.Master, eventsForLaunchpad, guiButtons)
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
		common.LightLamp(common.Button{X: myFixtureNumber, Y: mySequenceNumber}, common.Black, sequence.Master, eventsForLaunchpad, guiButtons)
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

	if targetSequence.Type == "rgb" {
		for _, pattern := range this.RGBPatterns {
			if debug {
				fmt.Printf("pattern is %s\n", pattern.Name)
			}
			if pattern.Number == targetSequence.SelectedPattern {
				common.FlashLight(common.Button{X: pattern.Number, Y: displaySequence}, common.White, common.LightBlue, eventsForLaunchpad, guiButtons)
			} else {
				common.LightLamp(common.Button{X: pattern.Number, Y: displaySequence}, common.LightBlue, master, eventsForLaunchpad, guiButtons)
			}
			common.LabelButton(pattern.Number, displaySequence, pattern.Label, guiButtons)
		}
		return
	}

	if targetSequence.Type == "scanner" {
		for _, pattern := range targetSequence.ScannerAvailablePatterns {
			if pattern.Number == targetSequence.SelectedPattern {
				common.FlashLight(common.Button{X: pattern.Number, Y: displaySequence}, common.White, common.LightBlue, eventsForLaunchpad, guiButtons)
			} else {
				common.LightLamp(common.Button{X: pattern.Number, Y: displaySequence}, common.LightBlue, master, eventsForLaunchpad, guiButtons)
			}
			common.LabelButton(pattern.Number, displaySequence, pattern.Label, guiButtons)
		}
		return
	}
}

func InitButtons(this *CurrentState, sequenceColors []color.RGBA, staticColors []color.RGBA, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight) {

	// Light the logo blue.
	if this.LaunchPadConnected {
		this.Pad.Light(8, -1, 0, 0, 255)
	}

	// Initially set the Flood, Save, Running and Blackout buttons to white.
	common.LightLamp(common.FLOOD_BUTTON, common.White, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)
	common.LightLamp(common.SAVE_BUTTON, common.White, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)
	common.LightLamp(common.RUNNING_BUTTON, common.White, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)
	common.LightLamp(common.STROBE_BUTTON, common.White, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)
	common.LightLamp(common.BLACKOUT_BUTTON, common.White, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)

	// Light up any existing presets.
	presets.RefreshPresets(eventsForLaunchpad, guiButtons, this.PresetsStore)

	// Show the correct labels at the bottom.
	showBottomLabels(this, sequenceColors, staticColors, eventsForLaunchpad, guiButtons)

	// Light the top labels.
	showTopLabels(this, eventsForLaunchpad, guiButtons)

	// Light the first sequence as the default selected.
	this.SelectedSequence = 0
	lightSelectedButton(eventsForLaunchpad, guiButtons, this)

}

func floodOff(this *CurrentState, commandChannels []chan common.Command, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight) {

	// Turn the flood button back to white.
	common.LightLamp(common.FLOOD_BUTTON, common.White, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)

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

func floodOn(this *CurrentState, commandChannels []chan common.Command, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight) {

	// Remember which sequence is currently selected.
	this.LastSelectedSequence = this.SelectedSequence

	// Flash the flood button pink to indicate we're in flood.
	common.FlashLight(common.FLOOD_BUTTON, common.Magenta, common.White, eventsForLaunchpad, guiButtons)

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

func buttonTouched(button common.Button, onColor color.RGBA, offColor color.RGBA, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight) {
	common.LightLamp(button, onColor, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)
	time.Sleep(200 * time.Millisecond)
	common.LightLamp(button, offColor, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)
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
		this.SelectedMode[sequenceNumber] = NORMAL
		this.ShowRGBColorPicker = false
		this.Static[this.DisplaySequence] = false
		this.Static[this.TargetSequence] = false
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

func lightSelectedButton(eventsForLauchpad chan common.ALight, guiButtons chan common.ALight, this *CurrentState) {

	NumberOfSelectableSequences := 4
	// 2 x RGB (FOH & Uplighters) sequences,
	// 1 x scanner sequence (chaser sequence shares its button with the scanner sequence),
	// 1 x switch sequence.

	if debug {
		fmt.Printf("SequenceSelect\n")
	}

	if this.SelectedSequence > NumberOfSelectableSequences-1 {
		return
	}

	// Turn off all sequence lights.
	for seq := 0; seq < NumberOfSelectableSequences; seq++ {
		common.LightLamp(common.Button{X: 8, Y: seq}, common.Cyan, common.MAX_DMX_BRIGHTNESS, eventsForLauchpad, guiButtons)
	}

	// Provided we're not the switch sequence number turn on the selected lamp.
	if this.SelectedSequence != this.SwitchSequenceNumber {
		// Turn on the correct sequence select number.
		if this.SelectedType == "scanner" && this.ScannerChaser[this.SelectedSequence] &&
			(this.SelectedMode[this.SelectedSequence] == CHASER_FUNCTION || this.SelectedMode[this.SelectedSequence] == CHASER_DISPLAY) {
			// If we are in shutter chaser mode, light the lamp yellow.
			common.LightLamp(common.Button{X: 8, Y: this.SelectedSequence}, common.Magenta, common.MAX_DMX_BRIGHTNESS, eventsForLauchpad, guiButtons)
		} else {
			// Now turn pink the selected sequence select light.
			common.LightLamp(common.Button{X: 8, Y: this.SelectedSequence}, common.Magenta, common.MAX_DMX_BRIGHTNESS, eventsForLauchpad, guiButtons)
		}
	}
}

func UpdateSpeed(this *CurrentState, guiButtons chan common.ALight) {

	mode := this.SelectedMode[this.DisplaySequence]
	tYpe := this.SelectedType
	speed := this.Speed[this.TargetSequence]
	switchSpeed := this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Speed

	if debug {
		fmt.Printf("UpdateSpeed Type=%s Switch %d Speed=%d\n", this.SelectedType, this.SelectedSwitch, this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Speed)
	}

	if this.Functions[this.TargetSequence][common.Function8_Music_Trigger].State {
		common.UpdateStatusBar("  MUSIC  ", "speed", false, guiButtons)
	} else {

		if mode == NORMAL || mode == FUNCTION || mode == STATUS {
			if tYpe == "rgb" {
				if !this.Strobe[this.TargetSequence] {
					common.UpdateStatusBar(fmt.Sprintf("Speed %02d", speed), "speed", false, guiButtons)
				} else {
					common.UpdateStatusBar(fmt.Sprintf("Strobe %02d", this.StrobeSpeed[this.TargetSequence]), "speed", false, guiButtons)
				}
			}
			if tYpe == "scanner" {
				common.UpdateStatusBar(fmt.Sprintf("Rotate Speed %02d", speed), "speed", false, guiButtons)
			}
			if tYpe == "switch" {
				if this.MusicTrigger {
					common.UpdateStatusBar("MUSIC", "speed", false, guiButtons)
				} else {
					common.UpdateStatusBar(fmt.Sprintf("Speed %02d", switchSpeed), "speed", false, guiButtons)
				}

			}
		}
		if mode == CHASER_DISPLAY || mode == CHASER_FUNCTION {
			if !this.Strobe[this.TargetSequence] {
				common.UpdateStatusBar(fmt.Sprintf("Chase Speed %02d", speed), "speed", false, guiButtons)
			} else {
				common.UpdateStatusBar(fmt.Sprintf("Strobe %02d", this.StrobeSpeed[this.TargetSequence]), "speed", false, guiButtons)
			}
		}
	}
}

func UpdateSize(this *CurrentState, guiButttons chan common.ALight) {

	mode := this.SelectedMode[this.DisplaySequence]
	tYpe := this.SelectedType
	size := this.RGBSize[this.TargetSequence]
	scannerFade := this.ScannerSize[this.TargetSequence]
	switchSize := this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Size

	if mode == NORMAL || mode == FUNCTION || mode == STATUS {
		if tYpe == "rgb" || tYpe == "switch" {
			common.UpdateStatusBar(fmt.Sprintf("Size %02d", size), "size", false, guiButttons)
		}
		if tYpe == "scanner" {
			common.UpdateStatusBar(fmt.Sprintf("Rotate Size %02d", scannerFade), "size", false, guiButttons)
		}
		if tYpe == "switch" {
			common.UpdateStatusBar(fmt.Sprintf("Size %02d", switchSize), "size", false, guiButttons)
		}
	}
	if mode == CHASER_DISPLAY || mode == CHASER_FUNCTION {
		common.UpdateStatusBar(fmt.Sprintf("Chase Size %02d", size), "size", false, guiButttons)
	}
}

func UpdateShift(this *CurrentState, guiButttons chan common.ALight) {

	mode := this.SelectedMode[this.DisplaySequence]
	tYpe := this.SelectedType
	shift := this.RGBShift[this.TargetSequence]
	scannerShift := getScannerShiftLabel(this.ScannerShift[this.TargetSequence])
	switchShift := this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Shift

	if debug {
		fmt.Printf("UpdateShift RGBShift=%d scannerShift=%s switchShift=%d\n", shift, scannerShift, switchShift)
	}

	if mode == NORMAL || mode == FUNCTION || mode == STATUS {
		if tYpe == "rgb" {
			common.UpdateStatusBar(fmt.Sprintf("Shift %02d", shift), "shift", false, guiButttons)
		}
		if tYpe == "scanner" {
			common.UpdateStatusBar(fmt.Sprintf("Rotate Shift %s", scannerShift), "shift", false, guiButttons)
		}
		if tYpe == "switch" {
			common.UpdateStatusBar(fmt.Sprintf("Shift %02d", switchShift), "shift", false, guiButttons)
		}
	}
	if mode == CHASER_DISPLAY || mode == CHASER_FUNCTION {
		common.UpdateStatusBar(fmt.Sprintf("Chase Shift %02d", shift), "shift", false, guiButttons)
	}
}

func UpdateFade(this *CurrentState, guiButttons chan common.ALight) {

	mode := this.SelectedMode[this.DisplaySequence]
	tYpe := this.SelectedType
	fade := this.RGBFade[this.TargetSequence]
	scannerCoordinates := getScannerCoordinatesLabel(this.ScannerCoordinates[this.TargetSequence])
	switchFade := this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Fade

	if mode == NORMAL || mode == FUNCTION || mode == STATUS {
		if tYpe == "rgb" {
			common.UpdateStatusBar(fmt.Sprintf("Fade %02d", fade), "fade", false, guiButttons)
		}
		if tYpe == "scanner" {
			common.UpdateStatusBar(fmt.Sprintf("Rotate Coord %s", scannerCoordinates), "fade", false, guiButttons)
		}
		if tYpe == "switch" {
			common.UpdateStatusBar(fmt.Sprintf("Fade %02d", switchFade), "fade", false, guiButttons)
		}
	}
	if mode == CHASER_DISPLAY || mode == CHASER_FUNCTION {
		common.UpdateStatusBar(fmt.Sprintf("Chase Fade %02d", fade), "fade", false, guiButttons)
	}
}

func StopStrobe(this *CurrentState, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command) {

	// Stop strobing this sequence.
	cmd := common.Command{
		Action: common.Strobe,
		Args: []common.Arg{
			{Name: "STROBE_STATE", Value: this.Strobe[this.SelectedSequence]},
			{Name: "STROBE_SPEED", Value: this.StrobeSpeed[this.SelectedSequence]},
		},
	}
	common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)
	if this.SelectedType == "scanner" && this.ScannerChaser[this.SelectedSequence] {
		common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)
	}

	// Update the strobe button and status bar.
	common.ShowStrobeButtonStatus(this.Strobe[this.SelectedSequence], eventsForLaunchpad, guiButtons)
	common.UpdateStatusBar(fmt.Sprintf("Speed %02d", this.Speed[this.SelectedSequence]), "speed", false, guiButtons)
}

func StartStrobe(this *CurrentState, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command) {
	cmd := common.Command{
		Action: common.Strobe,
		Args: []common.Arg{
			{Name: "STROBE_STATE", Value: this.Strobe[this.SelectedSequence]},
			{Name: "STROBE_SPEED", Value: this.StrobeSpeed[this.SelectedSequence]},
		},
	}
	// Store the strobe flag in all sequences.
	common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)
	if this.SelectedType == "scanner" && this.ScannerChaser[this.SelectedSequence] {
		common.SendCommandToSequence(this.ChaserSequenceNumber, cmd, commandChannels)
	}

	// Update the strobe button and status bar.
	common.ShowStrobeButtonStatus(this.Strobe[this.SelectedSequence], eventsForLaunchpad, guiButtons)
	common.UpdateStatusBar(fmt.Sprintf("Strobe %02d", this.StrobeSpeed[this.SelectedSequence]), "speed", false, guiButtons)
}

// SetTarget - If we're a scanner and we're in shutter chase mode and if we're in either CHASER_DISPLAY or CHASER_FUNCTION mode then
// set the target sequence to the chaser sequence number.
// Else the target is just this sequence number.
// Returns the target sequence number.
func SetTarget(this *CurrentState) {
	if this.SelectedType == "scanner" &&
		this.ScannerChaser[this.SelectedSequence] &&
		(this.SelectedMode[this.SelectedSequence] == CHASER_FUNCTION || this.SelectedMode[this.SelectedSequence] == CHASER_DISPLAY) {
		this.TargetSequence = this.ChaserSequenceNumber
		this.DisplaySequence = this.SelectedSequence
	} else {
		this.TargetSequence = this.SelectedSequence
		this.DisplaySequence = this.SelectedSequence
	}
}
