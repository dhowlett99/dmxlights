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
	"github.com/dhowlett99/dmxlights/pkg/colors"
	"github.com/dhowlett99/dmxlights/pkg/common"
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
	RGBPatterns                 map[int]common.Pattern     // Available RGB Patterns.
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

	handleLaunchPadCrash(sequences, X, Y, this, eventsForLaunchpad, guiButtons, commandChannels)

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

		SavePresetOff(this, eventsForLaunchpad, guiButtons)
		flashOn(sequences, X, Y, this, eventsForLaunchpad, guiButtons, fixturesConfig, dmxController)
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

		SavePresetOff(this, eventsForLaunchpad, guiButtons)
		flashOff(X, Y, this, eventsForLaunchpad, guiButtons, fixturesConfig, dmxController)
		return
	}

	// S E L E C T   P R E S E T S
	// recall (short press) or delete (long press) the preset.
	if X >= 100 && X < 108 &&
		(Y > 3 && Y < 7) {

		recallPreset(sequences, X, Y, this, eventsForLaunchpad, guiButtons, commandChannels, updateChannels)
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

	// S E L E C T   C L E A R  - Start the timer, waiting for a long press to clear all.
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

	// S E L E C T   C L E A R -  Clear all if we're not in the scanner mode.
	if X == 0 && Y == -1 && !this.GUI && sequences[this.SelectedSequence].Type != "scanner" {
		if debug {
			fmt.Printf("Clear All If We're Not in Scanner Mode X:%d Y:%d\n", X, Y)
		}

		Clear(X, Y, this, sequences, dmxController, fixturesConfig, commandChannels, eventsForLaunchpad, guiButtons, updateChannels)
		return
	}

	// S E L E C T   C L E A R  - We have a long press.
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
			// S E L E C T   S C A N N E R  P O S I T I O N  U P  A R R O W
			if sequences[this.SelectedSequence].Type == "scanner" {
				upArrow(X, Y, this, eventsForLaunchpad, guiButtons, commandChannels)
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

	// S E L E C T   F L O O D
	if X == 8 && Y == 3 {
		SavePresetOff(this, eventsForLaunchpad, guiButtons)
		toggleFlood(sequences, X, Y, this, commandChannels, eventsForLaunchpad, guiButtons)
		return
	}

	// S E L E C T   S O U N D  U P
	if X == 4 && Y == -1 {
		SavePresetOff(this, eventsForLaunchpad, guiButtons)
		increaseSensitivity(this, X, Y, eventsForLaunchpad, guiButtons)
		return
	}

	// S E L E C T   S O U N D  D O W N
	if X == 5 && Y == -1 {
		SavePresetOff(this, eventsForLaunchpad, guiButtons)
		decreaseSensitivity(this, X, Y, eventsForLaunchpad, guiButtons)
		return
	}

	// S E L E C T   M A S T E R  D O W N
	if X == 6 && Y == -1 {
		SavePresetOff(this, eventsForLaunchpad, guiButtons)
		decraseBrightness(X, Y, this, eventsForLaunchpad, guiButtons, commandChannels)
		return
	}

	// S E L E C T   M A S T E R  U P
	if X == 7 && Y == -1 {
		SavePresetOff(this, eventsForLaunchpad, guiButtons)
		increaseBrightness(X, Y, this, eventsForLaunchpad, guiButtons, commandChannels)
		return
	}

	// S E L E C T   S A V E  P R E S E T  M O D E
	if X == 8 && Y == 4 {
		togglePresetSaveMode(len(sequences), this, eventsForLaunchpad, guiButtons, commandChannels)
		return
	}

	// S E L E C T  S A V E   R E C A L L  P R E S E T S
	if X < 8 && (Y > 3 && Y < 7) {
		savePresets(sequences, X, Y, this, eventsForLaunchpad, guiButtons, commandChannels, updateChannels, replyChannels)
		return
	}

	// S E L E C T   D E C R E A S E  S H I F T
	if X == 2 && Y == 7 && !this.ShowRGBColorPicker {
		SavePresetOff(this, eventsForLaunchpad, guiButtons)
		decreaseShift(sequences, X, Y, this, eventsForLaunchpad, guiButtons, commandChannels)
		return
	}

	// S E L E C T   I N C R E A S E   S H I F T
	if X == 3 && Y == 7 && !this.ShowRGBColorPicker {
		SavePresetOff(this, eventsForLaunchpad, guiButtons)
		increaseShift(sequences, X, Y, this, eventsForLaunchpad, guiButtons, commandChannels)
		return
	}

	// S E L E C T   D E C R E A S E  S P E E D
	if X == 0 && Y == 7 && !this.ShowRGBColorPicker {
		SavePresetOff(this, eventsForLaunchpad, guiButtons)
		decreaseSpeed(sequences, X, Y, this, eventsForLaunchpad, guiButtons, commandChannels, updateChannels)
		return
	}

	// S E L E C T   I N C R E A S E   S P E E D
	if X == 1 && Y == 7 && !this.ShowRGBColorPicker {
		SavePresetOff(this, eventsForLaunchpad, guiButtons)
		increaseSpeed(sequences, X, Y, this, eventsForLaunchpad, guiButtons, commandChannels, updateChannels)
		return
	}

	// S E L E C T   S E Q U E N C E
	if X == 8 && (Y == 0 || Y == 1 || Y == 2) {
		SavePresetOff(this, eventsForLaunchpad, guiButtons)
		selectSequence(sequences, Y, this, eventsForLaunchpad, guiButtons, commandChannels)
		return
	}

	// S T A R T - Start sequence.
	if X == 8 && Y == 5 {
		SavePresetOff(this, eventsForLaunchpad, guiButtons)
		toggleSequence(sequences, X, Y, this, eventsForLaunchpad, guiButtons, commandChannels)
		return
	}

	// S T R O B E - Strobe.
	if X == 8 && Y == 6 {
		SavePresetOff(this, eventsForLaunchpad, guiButtons)
		toggleStrobe(sequences, X, Y, this, eventsForLaunchpad, guiButtons, commandChannels)
		return
	}

	// S E L E C T   D E C R E A S E  S I Z E
	if X == 4 && Y == 7 && !this.ShowRGBColorPicker {
		SavePresetOff(this, eventsForLaunchpad, guiButtons)
		decreaseSize(sequences, X, Y, this, eventsForLaunchpad, guiButtons, commandChannels)
		return
	}

	// S E L E C T   I N C R E A S E  S I Z E
	if X == 5 && Y == 7 && !this.ShowRGBColorPicker {
		SavePresetOff(this, eventsForLaunchpad, guiButtons)
		increaseSize(sequences, X, Y, this, eventsForLaunchpad, guiButtons, commandChannels)
		return
	}

	// S E L E C T   D E C R E A S E  F A D E
	if X == 6 && Y == 7 && !this.ShowRGBColorPicker {
		SavePresetOff(this, eventsForLaunchpad, guiButtons)
		decreaseFade(sequences, X, Y, this, eventsForLaunchpad, guiButtons, commandChannels, fixturesConfig)
		return
	}

	// S E L E C T   I N C R E A S E  F A D E
	if X == 7 && Y == 7 && !this.ShowRGBColorPicker {
		SavePresetOff(this, eventsForLaunchpad, guiButtons)
		increaseFade(sequences, X, Y, this, eventsForLaunchpad, guiButtons, commandChannels, fixturesConfig)
		return
	}

	// S E L E C T   S W I T C H   B U T T O N
	if X >= 0 && X < 8 && Y >= 0 && Y < 4 && sequences[Y].Type == "switch" {
		SavePresetOff(this, eventsForLaunchpad, guiButtons)
		selectSwitch(sequences, X, Y, this, eventsForLaunchpad, guiButtons, commandChannels, updateChannels, fixturesConfig)
		return
	}

	// S E L E C T   D I S A B L E  / E N A B L E   F I X T U R E  S T A T U S
	if X >= 0 && X < 8 && Y >= 0 && Y < 4 && this.SelectedMode[this.SelectedSequence] == STATUS {
		SavePresetOff(this, eventsForLaunchpad, guiButtons)
		toggleFixtureStatus(sequences, X, Y, this, eventsForLaunchpad, guiButtons, commandChannels)
		return
	}

	// S E L E C T   S C A N N E R  P O S I T I O N  D O W N  A R R O W
	if X == 1 && Y == -1 && sequences[this.SelectedSequence].Type == "scanner" {
		SavePresetOff(this, eventsForLaunchpad, guiButtons)
		downArrow(X, Y, this, eventsForLaunchpad, guiButtons, commandChannels)
		return
	}

	// S E L E C T   S C A N N E R  P O S I T I O N  L E F T  A R R O W
	if X == 2 && Y == -1 && sequences[this.SelectedSequence].Type == "scanner" {
		SavePresetOff(this, eventsForLaunchpad, guiButtons)
		leftArrow(X, Y, this, eventsForLaunchpad, guiButtons, commandChannels)
		return
	}

	// S E L E C T   S C A N N E R  P O S I T I O N  R I G H T  A R R O W
	if X == 3 && Y == -1 && sequences[this.SelectedSequence].Type == "scanner" {
		SavePresetOff(this, eventsForLaunchpad, guiButtons)
		rightArrow(X, Y, this, eventsForLaunchpad, guiButtons, commandChannels)
		return
	}

	// S E L E C T   E D I T  R E D  S T A T I C   C O L O R
	if X == 1 && Y == -1 && this.Static[this.TargetSequence] {
		SavePresetOff(this, eventsForLaunchpad, guiButtons)
		redButton(this, X, Y, eventsForLaunchpad, guiButtons, commandChannels)
		return
	}

	// S E L E C T   E D I T  G R E E N  S T A T I C   C O L O R
	if X == 2 && Y == -1 && this.Static[this.TargetSequence] {
		SavePresetOff(this, eventsForLaunchpad, guiButtons)
		greenButton(this, X, Y, eventsForLaunchpad, guiButtons, commandChannels)
		return
	}

	// S E L E C T   E D I T  B L U E  S T A T I C   C O L O R
	if X == 3 && Y == -1 && this.Static[this.TargetSequence] {
		SavePresetOff(this, eventsForLaunchpad, guiButtons)
		blueButton(this, X, Y, eventsForLaunchpad, guiButtons, commandChannels)
		return
	}

	// S E L E C T   R G B  C H A S E  C O L O R S
	if X >= 0 && X < 8 &&
		Y != -1 && Y < 3 &&
		!this.EditFixtureSelectionMode &&
		!this.EditScannerColorsMode &&
		this.ShowRGBColorPicker {

		SavePresetOff(this, eventsForLaunchpad, guiButtons)
		selectRGBChaseColor(sequences, X, Y, this, eventsForLaunchpad, guiButtons, commandChannels)
		return
	}

	// S E L E C T   S C A N N E R   C O L O R
	if X >= 0 && X < 8 && Y != -1 &&
		this.SelectedSequence == Y && // Make sure the buttons pressed are for this sequence.
		!this.EditFixtureSelectionMode &&
		sequences[this.SelectedSequence].Type == "scanner" &&
		this.Functions[this.SelectedSequence][common.Function5_Color].State {

		SavePresetOff(this, eventsForLaunchpad, guiButtons)
		selectScannerColor(sequences, X, Y, this, eventsForLaunchpad, guiButtons, commandChannels, updateChannels, dmxController, fixturesConfig)
		return
	}

	// S E L E C T   F I X T U R E
	if X >= 0 && X < 8 && Y != -1 &&
		this.EditFixtureSelectionMode &&
		this.Functions[this.SelectedSequence][common.Function6_Static_Gobo].State || this.Functions[this.SelectedSequence][common.Function5_Color].State &&
		sequences[this.SelectedSequence].Type == "scanner" {

		SavePresetOff(this, eventsForLaunchpad, guiButtons)
		selectFixture(sequences, X, Y, this, eventsForLaunchpad, guiButtons, commandChannels, fixturesConfig)
		return
	}

	// S E L E C T   G O B O
	if X >= 0 && X < 8 && Y != -1 &&
		this.SelectedSequence == Y && // Make sure the buttons pressed are for this sequence.
		!this.EditFixtureSelectionMode &&
		sequences[this.SelectedSequence].Type == "scanner" &&
		this.Functions[this.SelectedSequence][common.Function6_Static_Gobo].State {

		SavePresetOff(this, eventsForLaunchpad, guiButtons)
		selectGobo(sequences, X, Y, this, eventsForLaunchpad, guiButtons, commandChannels, updateChannels, fixturesConfig, dmxController)
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

		SavePresetOff(this, eventsForLaunchpad, guiButtons)
		selectStaticFixture(sequences, X, Y, this, eventsForLaunchpad, guiButtons, commandChannels)
		return
	}

	// S E L E C T   S T A T I C   C O L O R
	if X >= 0 && X < 8 && Y != -1 &&
		Y < 3 && // Make sure the buttons pressed inside the color picker.
		this.ShowStaticColorPicker && // Now We Are In Static Color Picker Mode.
		!this.EditFixtureSelectionMode && // Not In Fixture Selection Mode.
		getStatic(this) { // Static Function On in this or shutter chaser sequence

		SavePresetOff(this, eventsForLaunchpad, guiButtons)
		selectStaticColor(sequences, X, Y, this, eventsForLaunchpad, guiButtons, commandChannels)
		return
	}

	// S E L E C T   P A T T E R N
	if X >= 0 && X < 8 && Y != -1 &&
		!this.EditFixtureSelectionMode &&
		this.EditPatternMode {

		SavePresetOff(this, eventsForLaunchpad, guiButtons)
		selectPattern(sequences, X, Y, this, eventsForLaunchpad, guiButtons, commandChannels, updateChannels)
		return
	}

	// S E L E C T   F U N C T I O N  K E Y S
	if X >= 0 && X < 8 && Y >= 0 && Y < 3 &&
		this.SelectedSequence == Y && // Make sure the buttons pressed are for this sequence.
		this.SelectedMode[this.SelectedSequence] == FUNCTION || this.SelectedMode[this.SelectedSequence] == CHASER_FUNCTION {

		SavePresetOff(this, eventsForLaunchpad, guiButtons)
		processFunctions(sequences, X, Y, this, eventsForLaunchpad, guiButtons, commandChannels, updateChannels)
		return
	}

	// S E L E C T   B L A C K O U T   B U T T O N.
	if X == 8 && Y == 7 {
		SavePresetOff(this, eventsForLaunchpad, guiButtons)
		blackout(X, Y, this, eventsForLaunchpad, guiButtons, commandChannels)
		return
	}
}

// For the given sequence show the available sequence colors on the relevant buttons.
// With the new color picker there can be 24 colors displayed.
// ShowRGBColorPicker operates on the sequence.RGBAvailableColors which is an array of type []common.StaticColorButton
// the targetSequence .CurrentColors selects which colors are selected.
// Returns the RGBAvailableColors []common.StaticColorButton
func ShowRGBColorPicker(targetSequence common.Sequence, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, commandChannels []chan common.Command) {

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
			for _, sequenceColor := range targetSequence.SequenceColors {
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
			Black := colors.Black
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
			common.FlashLight(common.Button{X: fixtureNumber, Y: displaySequence}, fixture.Color, colors.White, eventsForLaunchpad, guiButtons)
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
			Black := colors.Black
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
				common.LightLamp(common.Button{X: fixture.Number - 1, Y: this.SelectedSequence}, colors.White, sequence.Master, eventsForLaunchpad, guiButtons)
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
			Black := colors.Black
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
		common.LightLamp(common.Button{X: myFixtureNumber, Y: mySequenceNumber}, colors.Black, sequence.Master, eventsForLaunchpad, guiButtons)
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
				common.FlashLight(common.Button{X: pattern.Number, Y: displaySequence}, colors.White, colors.LightBlue, eventsForLaunchpad, guiButtons)
			} else {
				common.LightLamp(common.Button{X: pattern.Number, Y: displaySequence}, colors.LightBlue, master, eventsForLaunchpad, guiButtons)
			}
			common.LabelButton(pattern.Number, displaySequence, pattern.Label, guiButtons)
		}
		return
	}

	if targetSequence.Type == "scanner" {
		for _, pattern := range targetSequence.ScannerAvailablePatterns {
			if pattern.Number == targetSequence.SelectedPattern {
				common.FlashLight(common.Button{X: pattern.Number, Y: displaySequence}, colors.White, colors.LightBlue, eventsForLaunchpad, guiButtons)
			} else {
				common.LightLamp(common.Button{X: pattern.Number, Y: displaySequence}, colors.LightBlue, master, eventsForLaunchpad, guiButtons)
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
	common.LightLamp(common.FLOOD_BUTTON, colors.White, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)
	common.LightLamp(common.SAVE_BUTTON, colors.White, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)
	common.LightLamp(common.RUNNING_BUTTON, colors.White, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)
	common.LightLamp(common.STROBE_BUTTON, colors.White, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)
	common.LightLamp(common.BLACKOUT_BUTTON, colors.White, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)

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
		return "    NORMAL     "
	}
	if mode == NORMAL_STATIC {
		return "    STATIC     "
	}
	if mode == CHASER_DISPLAY {
		return "    CHASER     "
	}
	if mode == CHASER_DISPLAY_STATIC {
		return " CHASER_STATIC "
	}
	if mode == FUNCTION {
		return "   FUNCTION    "
	}
	if mode == CHASER_FUNCTION {
		return "CHASER_FUNCTION"
	}
	if mode == STATUS {
		return "     STATUS    "
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
		common.LightLamp(common.Button{X: 8, Y: seq}, colors.Cyan, common.MAX_DMX_BRIGHTNESS, eventsForLauchpad, guiButtons)
	}

	// Provided we're not the switch sequence number turn on the selected lamp.
	if this.SelectedSequence != this.SwitchSequenceNumber {
		// Turn on the correct sequence select number.
		if this.SelectedType == "scanner" && this.ScannerChaser[this.SelectedSequence] &&
			(this.SelectedMode[this.SelectedSequence] == CHASER_FUNCTION || this.SelectedMode[this.SelectedSequence] == CHASER_DISPLAY) {
			// If we are in shutter chaser mode, light the lamp yellow.
			common.LightLamp(common.Button{X: 8, Y: this.SelectedSequence}, colors.Magenta, common.MAX_DMX_BRIGHTNESS, eventsForLauchpad, guiButtons)
		} else {
			// Now turn pink the selected sequence select light.
			common.LightLamp(common.Button{X: 8, Y: this.SelectedSequence}, colors.Magenta, common.MAX_DMX_BRIGHTNESS, eventsForLauchpad, guiButtons)
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
	switchColor := this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Color
	switchColorName := this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].ColorName

	if mode == NORMAL || mode == FUNCTION || mode == STATUS {
		if tYpe == "rgb" || tYpe == "switch" {
			common.UpdateStatusBar(fmt.Sprintf("Size %02d", size), "size", false, guiButttons)
		}
		if tYpe == "scanner" {
			common.UpdateStatusBar(fmt.Sprintf("Rotate Size %02d", scannerFade), "size", false, guiButttons)
		}
		if tYpe == "switch" && this.SelectedFixtureType == "rgb" {
			common.UpdateStatusBar(fmt.Sprintf("Size %02d", switchSize), "size", false, guiButttons)
		}
		if tYpe == "switch" && this.SelectedFixtureType == "projector" {
			common.UpdateStatusBar(fmt.Sprintf("Color %02d:%s", switchColor, switchColorName), "size", false, guiButttons)
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
	switchRGBShift := this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Shift
	switchRoatetSpeed := this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].RotateSpeed

	if debug {
		fmt.Printf("UpdateShift RGBShift=%d scannerShift=%s switchShift=%d switchRotateSpeed %d\n", shift, scannerShift, switchRGBShift, switchRoatetSpeed)
	}

	if mode == NORMAL || mode == FUNCTION || mode == STATUS {
		if tYpe == "rgb" {
			common.UpdateStatusBar(fmt.Sprintf("Shift %02d", shift), "shift", false, guiButttons)
		}
		if tYpe == "scanner" {
			common.UpdateStatusBar(fmt.Sprintf("Rotate Shift %s", scannerShift), "shift", false, guiButttons)
		}
		if tYpe == "switch" {
			if this.SelectedFixtureType == "rgb" {
				common.UpdateStatusBar(fmt.Sprintf("Shift %02d", switchRGBShift), "shift", false, guiButttons)
			}
			if this.SelectedFixtureType == "projector" {
				common.UpdateStatusBar(fmt.Sprintf("Rotate Speed %02d", switchRoatetSpeed), "shift", false, guiButttons)
			}
		}
	}
	if mode == CHASER_DISPLAY || mode == CHASER_FUNCTION {
		common.UpdateStatusBar(fmt.Sprintf("Chase Shift %02d", shift), "shift", false, guiButttons)
	}
}

func UpdateFade(this *CurrentState, guiButttons chan common.ALight) {

	mode := this.SelectedMode[this.DisplaySequence]
	tYpe := this.SelectedType
	fixtureType := this.SelectedFixtureType
	fade := this.RGBFade[this.TargetSequence]
	scannerCoordinates := getScannerCoordinatesLabel(this.ScannerCoordinates[this.TargetSequence])
	switchFade := this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Fade
	switchGobo := this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].Gobo
	switchGoboName := this.SwitchOverrides[this.SelectedSwitch][this.SwitchPosition[this.SelectedSwitch]].GoboName

	if mode == NORMAL || mode == FUNCTION || mode == STATUS {
		if tYpe == "rgb" {
			common.UpdateStatusBar(fmt.Sprintf("Fade %02d", fade), "fade", false, guiButttons)
		}
		if tYpe == "scanner" {
			common.UpdateStatusBar(fmt.Sprintf("Rotate Coord %s", scannerCoordinates), "fade", false, guiButttons)
		}
		if tYpe == "switch" && fixtureType == "rgb" {
			common.UpdateStatusBar(fmt.Sprintf("Fade %02d", switchFade), "fade", false, guiButttons)
		}
		if tYpe == "switch" && fixtureType == "projector" {
			common.UpdateStatusBar(fmt.Sprintf("Gobo %02d:%s", switchGobo, switchGoboName), "fade", false, guiButttons)
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
