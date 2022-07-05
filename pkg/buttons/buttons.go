package buttons

import (
	"fmt"

	"time"

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/config"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
	"github.com/dhowlett99/dmxlights/pkg/presets"
	"github.com/dhowlett99/dmxlights/pkg/sequence"
	"github.com/oliread/usbdmx/ft232"
	"github.com/rakyll/launchpad/mk3"
)

const debug = false

const (
	full = 255
)

type CurrentState struct {
	SelectedSequence         int       // The currently selected sequence.
	SequenceSpeed            int       // Local copy of sequence speed.
	Size                     int       // current RGB sequence this.Size.
	ScannerSize              int       // Current scanner this.Size.
	SavePreset               bool      // Save a preset flag.
	SelectedShift            int       // Current fixture shift.
	Blackout                 bool      // Blackout all fixtures.
	Flood                    bool      // Flood all fixtures.
	FunctionSelectMode       []bool    // Which sequence is in function selection mode.
	StaticLamps              [][]bool  // Static color lamps.
	SelectButtonPressed      []bool    // Which sequence has its Select button pressed.
	SwitchPositions          [9][9]int // Sorage for switch positions.
	EditSequenceColorsMode   []bool    // This flag is true when the sequence is in sequence colors editing mode.
	EditScannerColorsMode    []bool    // This flag is true when the sequence is in select scanner colors editing mode.
	EditGoboSelectionMode    []bool    // This flag is true when the sequence is in sequence gobo selection mode.
	EditStaticColorsMode     []bool    // This flag is true when the sequence is in static colors editing mode.
	EditPattenMode           []bool    // This flag is true when the sequence is in patten editing mode.
	EditFixtureSelectionMode bool      // This flag is true when the sequence is in select fixture mode.
	FadeSpeed                int       // Default start at 50ms.
	MasterBrightness         int       // Affects all DMX fixtures and launchpad lamps.
	LastStaticColorButtonX   int       // Which Static Color button did we change last.
	LastStaticColorButtonY   int       // Which Static Color button did we change last.
	SoundGain                float32   // Fine gain -0.09 -> 0.09
	DisabledFixture          [][]bool  // Which fixture is disabled on which sequence.
	SelectedFixture          int       // Which fixture is selected when changing scanner color or gobo.
	FollowingAction          string    // String to find next function, used in selecting a fixture.
	SelectedCordinates       int       // Number of coordinates for scanner patterns is selected from 4 choices. 0=12, 1=26,2=24,3=32
	OffsetPan                int       // Start pan from the center
	OffsetTilt               int
	Pad                      *mk3.Launchpad
	PresetsStore             map[string]bool
	SoundTriggers            []*common.Trigger
	SequenceChannels         common.Channels
	Pattens                  map[string]common.Patten
	StaticButtons            []common.StaticColorButton
	SelectedFloodMap         map[int]bool
	SelectedGobo             int
}

// main thread is used to get commands from the lauchpad.
func ReadLaunchPadButtons(guiButtons chan common.ALight, this *CurrentState, sequences []*common.Sequence,
	eventsForLauchpad chan common.ALight, dmxController *ft232.DMXController,
	fixturesConfig *fixture.Fixtures, commandChannels []chan common.Command,
	replyChannels []chan common.Sequence, updateChannels []chan common.Sequence) {

	// Create a channel to listen for buttons being pressed.
	buttonChannel := this.Pad.Listen()

	// Main loop reading commands from the Novation Launchpad.
	for {
		hit := <-buttonChannel
		ProcessButtons(hit.X, hit.Y, sequences, this, eventsForLauchpad, guiButtons, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels)
	}
}

func ProcessButtons(X int, Y int,
	sequences []*common.Sequence,
	this *CurrentState,
	eventsForLauchpad chan common.ALight,
	guiButtons chan common.ALight,
	dmxController *ft232.DMXController,
	fixturesConfig *fixture.Fixtures,
	commandChannels []chan common.Command,
	replyChannels []chan common.Sequence,
	updateChannels []chan common.Sequence) {

	if debug {
		fmt.Printf("ProcessButtons Called with X:%d Y:%d\n", X, Y)
	}

	// C L E A R  - Clear all the lights on the common.
	if X == 0 && Y == -1 && sequences[this.SelectedSequence].Type != "scanner" {

		if debug {
			fmt.Printf("CLEAR LAUNCHPAD\n")
		}

		// Turn off the flashing save button.
		this.SavePreset = false
		common.LightLamp(common.ALight{X: 8, Y: 4, Brightness: 0, Red: 0, Green: 0, Blue: 0}, eventsForLauchpad, guiButtons)

		// Turn off the this.Flood
		if this.Flood {
			cmd := common.Command{
				Action: common.NoFlood,
				Args: []common.Arg{
					{Name: "Flood", Value: false},
				},
			}
			common.SendCommandToAllSequence(this.SelectedSequence, cmd, commandChannels)
			this.Flood = false
		}

		// We want to clear a color selection for a selected sequence.
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
			ShowRGBColorSelectionButtons(this.SelectedSequence, *sequences[this.SelectedSequence], eventsForLauchpad, guiButtons)

			return
		}

		common.ClearAll(this.Pad, this.PresetsStore, eventsForLauchpad, guiButtons, commandChannels)
		AllFixturesOff(eventsForLauchpad, guiButtons, dmxController, fixturesConfig)
		presets.ClearPresets(eventsForLauchpad, guiButtons, this.PresetsStore)
		presets.InitPresets(eventsForLauchpad, guiButtons, this.PresetsStore)

		// Make sure we stop all sequences.
		cmd := common.Command{
			Action: common.Stop,
		}
		common.SendCommandToAllSequence(this.SelectedSequence, cmd, commandChannels)

		// Swicth off any static colors.
		cmd = common.Command{
			Action: common.UpdateStatic,
			Args: []common.Arg{
				{Name: "Static", Value: false},
			},
		}
		common.SendCommandToAllSequence(this.SelectedSequence, cmd, commandChannels)

		// refesh the switch positions.
		cmd = common.Command{
			Action: common.UpdateSwitchPositions,
		}
		common.SendCommandToAllSequenceOfType(sequences, this.SelectedSequence, cmd, commandChannels, "switch")

		// Clear the sequence colors.
		cmd = common.Command{
			Action: common.ClearSequenceColor,
		}
		common.SendCommandToAllSequence(this.SelectedSequence, cmd, commandChannels)

		// Clear out soundtriggers
		for _, trigger := range this.SoundTriggers {
			trigger.State = false
		}

		// Clear all the function buttons.
		for sequenceNumber := range sequences {
			sequences[sequenceNumber].Functions[common.Function1_Patten].State = false
			sequences[sequenceNumber].Functions[common.Function2_Auto_Color].State = false
			sequences[sequenceNumber].Functions[common.Function3_Auto_Patten].State = false
			sequences[sequenceNumber].Functions[common.Function4_Bounce].State = false
			sequences[sequenceNumber].Functions[common.Function5_Color].State = false
			sequences[sequenceNumber].Functions[common.Function6_Gobo].State = false
			sequences[sequenceNumber].Functions[common.Function6_Static].State = false
			sequences[sequenceNumber].Functions[common.Function7_RGB_Invert].State = false
			sequences[sequenceNumber].Functions[common.Function8_Music_Trigger].State = false

			// Send update functions command. This sets the temporary representation of
			// the function keys in the real sequence.
			cmd := common.Command{
				Action: common.UpdateFunctions,
				Args: []common.Arg{
					{Name: "Functions", Value: sequences[sequenceNumber].Functions},
				},
			}
			common.SendCommandToSequence(sequenceNumber, cmd, commandChannels)

		}

		// Disable fixtures.
		for x := 0; x < 4; x++ {
			for y := 0; y < 9; y++ {
				this.DisabledFixture[x][y] = false
				// Tell the sequence to turn on this scanner.
				cmd := common.Command{
					Action: common.ToggleFixtureState,
					Args: []common.Arg{
						{Name: "SequenceNumber", Value: x},
						{Name: "FixtureNumber", Value: y},
						{Name: "FixtureState", Value: false},
					},
				}
				common.SendCommandToSequence(x, cmd, commandChannels)
			}
		}

		// Reset the Scanner Size back to default.
		for _, sequence := range sequences {
			// Set local copy.
			this.ScannerSize = common.DefaultScannerSize
			// Set copy in sequences.
			if sequence.Type == "scanner" {
				cmd = common.Command{
					Action: common.UpdateScannerSize,
					Args: []common.Arg{
						{Name: "ScannerSize", Value: common.DefaultScannerSize},
					},
				}
				common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)
			}
		}

		return
	}

	// F L O O D
	if X == 8 && Y == 3 {

		if debug {
			fmt.Printf("FLOOD\n")
		}

		if !this.Flood { // We're not already in this.Flood so lets ask the sequence to this.Flood.

			// Flash the this.Flood button to indicate we're not in this.Flood.
			common.FlashLight(8, 3, 0x03, 0x5f, eventsForLauchpad, guiButtons)

			// First save our config
			config.AskToSaveConfig(commandChannels, replyChannels, 0, 0)

			// Stop all sequences, so we start in sync.
			cmd := common.Command{
				Action: common.Stop,
			}
			common.SendCommandToAllSequence(this.SelectedSequence, cmd, commandChannels)

			// Start this.Flood.
			cmd = common.Command{
				Action: common.Flood,
				Args: []common.Arg{
					{Name: "StartFlood", Value: true},
				},
			}
			common.SendCommandToAllSequence(this.SelectedSequence, cmd, commandChannels)

			this.Flood = true

			return
		}
		if this.Flood { // If we are this.Flood already then tell the sequence to stop this.Flood.

			// Turn the this.Flood button back to white.
			common.LightLamp(common.ALight{X: X, Y: Y, Brightness: full, Red: 255, Green: 255, Blue: 255}, eventsForLauchpad, guiButtons)

			cmd := common.Command{
				Action: common.NoFlood,
				Args: []common.Arg{
					{Name: "StopFlood", Value: false},
				},
			}
			common.SendCommandToAllSequence(this.SelectedSequence, cmd, commandChannels)

			this.Flood = false

			// Recall our previous config
			config.AskToLoadConfig(commandChannels, 0, 0)

			// Restart all the sequences.
			cmd = common.Command{
				Action: common.Start,
			}
			common.SendCommandToAllSequence(this.SelectedSequence, cmd, commandChannels)

			// Clear any function modes out and reveal sequence.
			for this.SelectedSequence = range sequences {
				// Turn off function mode. Remove the function pads.
				common.HideFunctionButtons(this.SelectedSequence, eventsForLauchpad, guiButtons)
				// And reveal the sequence on the launchpad keys
				common.RevealSequence(this.SelectedSequence, commandChannels)
				// Turn off the function mode flag.
				this.FunctionSelectMode[this.SelectedSequence] = false
				// Now forget we pressed twice and start again.
				this.SelectButtonPressed[this.SelectedSequence] = false
			}

			return
		}
	}

	// Sound sensitity up.
	if X == 4 && Y == -1 {
		if debug {
			fmt.Printf("Sound Up %f\n", this.SoundGain)
		}

		this.SoundGain = this.SoundGain - 0.01
		if this.SoundGain < -0.04 {
			this.SoundGain = -0.04
		}
		for _, trigger := range this.SoundTriggers {
			trigger.Gain = this.SoundGain
		}
	}

	// Sound sensitity down.
	if X == 5 && Y == -1 {
		if debug {
			fmt.Printf("Sound Down%f\n", this.SoundGain)
		}
		this.SoundGain = this.SoundGain + 0.01
		if this.SoundGain > 0.9 {
			this.SoundGain = 0.9
		}
		for _, trigger := range this.SoundTriggers {
			trigger.Gain = this.SoundGain
		}
	}

	// Master brightness down.
	if X == 6 && Y == -1 {

		if debug {
			fmt.Printf("Brightness Down \n")
		}

		this.MasterBrightness = this.MasterBrightness - 10
		if this.MasterBrightness < 0 {
			this.MasterBrightness = 0
		}
		cmd := common.Command{
			Action: common.MasterBrightness,
			Args: []common.Arg{
				{Name: "Master", Value: this.MasterBrightness},
			},
		}
		common.SendCommandToAllSequence(this.SelectedSequence, cmd, commandChannels)
		return
	}

	// Master brightness up.
	if X == 7 && Y == -1 {

		if debug {
			fmt.Printf("Brightness Up \n")
		}

		this.MasterBrightness = this.MasterBrightness + 10
		if this.MasterBrightness > 255 {
			this.MasterBrightness = 255
		}
		cmd := common.Command{
			Action: common.MasterBrightness,
			Args: []common.Arg{
				{Name: "Master", Value: this.MasterBrightness},
			},
		}
		common.SendCommandToAllSequence(this.SelectedSequence, cmd, commandChannels)
		return
	}

	// Save mode.
	if X == 8 && Y == 4 {
		if debug {
			fmt.Printf("Save Mode\n")
		}

		if this.SavePreset {
			this.SavePreset = false
			presets.InitPresets(eventsForLauchpad, guiButtons, this.PresetsStore)
			common.LightLamp(common.ALight{X: 8, Y: 4, Brightness: full, Red: 255, Green: 255, Blue: 255}, eventsForLauchpad, guiButtons)
			return
		}
		presets.InitPresets(eventsForLauchpad, guiButtons, this.PresetsStore)
		common.FlashLight(8, 4, 0x03, 0x5f, eventsForLauchpad, guiButtons)
		this.SavePreset = true
		return
	}

	// Ask all sequences for their current config and save in a file.
	if X < 8 && (Y > 3 && Y < 7) {

		if debug {
			fmt.Printf("Ask For Config\n")
		}

		if this.SavePreset {
			this.PresetsStore[fmt.Sprint(X)+","+fmt.Sprint(Y)] = true
			common.LightLamp(common.ALight{X: X, Y: Y, Brightness: full, Red: 255, Green: 0, Blue: 0}, eventsForLauchpad, guiButtons)
			config.AskToSaveConfig(commandChannels, replyChannels, X, Y)
			this.SavePreset = false

			// turn off the save button from flashing.
			common.LightLamp(common.ALight{X: 8, Y: 4, Brightness: 0, Red: 0, Green: 0, Blue: 0}, eventsForLauchpad, guiButtons)

			presets.SavePresets(this.PresetsStore)
			presets.ClearPresets(eventsForLauchpad, guiButtons, this.PresetsStore)
			presets.InitPresets(eventsForLauchpad, guiButtons, this.PresetsStore)
			common.FlashLight(X, Y, 0x0c, 0x78, eventsForLauchpad, guiButtons)
		} else {
			// L O A D - Load config, but only if it exists in the presets map.
			if this.PresetsStore[fmt.Sprint(X)+","+fmt.Sprint(Y)] {

				// Stop all sequences, so we start in sync.
				cmd := common.Command{
					Action: common.Stop,
				}
				common.SendCommandToAllSequence(this.SelectedSequence, cmd, commandChannels)

				AllFixturesOff(eventsForLauchpad, guiButtons, dmxController, fixturesConfig)
				time.Sleep(300 * time.Millisecond)

				// Load the config.
				config.AskToLoadConfig(commandChannels, X, Y)

				// Turn the selected preset light red.
				common.LightLamp(common.ALight{X: X, Y: Y, Brightness: full, Red: 3, Green: 0, Blue: 0}, eventsForLauchpad, guiButtons)
				presets.InitPresets(eventsForLauchpad, guiButtons, this.PresetsStore)
				common.FlashLight(X, Y, 0x0c, 0x78, eventsForLauchpad, guiButtons)

				// Preserve this.Blackout.
				if !this.Blackout {
					cmd := common.Command{
						Action: common.Normal,
					}
					common.SendCommandToAllSequence(this.SelectedSequence, cmd, commandChannels)
				}

				// Turn off the local copy of the this.Flood flag.
				this.Flood = false

			}
		}
		return
	}

	// Decrease Shift.
	if X == 2 && Y == 7 {

		if debug {
			fmt.Printf("Decrease Shift\n")
		}

		this.SelectedShift = this.SelectedShift - 1
		if this.SelectedShift < 0 {
			this.SelectedShift = 0
		}
		cmd := common.Command{
			Action: common.UpdateShift,
			Args: []common.Arg{
				{Name: "Shift", Value: this.SelectedShift},
			},
		}
		common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

		return
	}

	// Increase Shift.
	if X == 3 && Y == 7 {

		if debug {
			fmt.Printf("Increase Shift \n")
		}

		this.SelectedShift = this.SelectedShift + 1
		if this.SelectedShift > 3 {
			this.SelectedShift = 3
		}
		cmd := common.Command{
			Action: common.UpdateShift,
			Args: []common.Arg{
				{Name: "Shift", Value: this.SelectedShift},
			},
		}
		common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

		return
	}

	// Decrease speed of selected sequence.
	if X == 0 && Y == 7 {

		if debug {
			fmt.Printf("Decrease Speed \n")
		}

		// Get an upto date copy of the sequence.
		sequences[this.SelectedSequence] = common.RefreshSequence(this.SelectedSequence, commandChannels, updateChannels)

		if !sequences[this.SelectedSequence].MusicTrigger {
			this.SequenceSpeed--
			if this.SequenceSpeed < 0 {
				this.SequenceSpeed = 1
			}
			cmd := common.Command{
				Action: common.UpdateSpeed,
				Args: []common.Arg{
					{Name: "Speed", Value: this.SequenceSpeed},
				},
			}
			common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)
		}
		return
	}

	// Increase speed of selected sequence.
	if X == 1 && Y == 7 {

		if debug {
			fmt.Printf("Increase Speed \n")
		}

		// Get an upto date copy of the sequence.
		sequences[this.SelectedSequence] = common.RefreshSequence(this.SelectedSequence, commandChannels, updateChannels)

		if !sequences[this.SelectedSequence].MusicTrigger {
			this.SequenceSpeed++
			if this.SequenceSpeed > 21 {
				this.SequenceSpeed = 21
			}
			cmd := common.Command{
				Action: common.UpdateSpeed,
				Args: []common.Arg{
					{Name: "Speed", Value: this.SequenceSpeed},
				},
			}
			common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)
		}
		return
	}

	// S E L E C T    S E Q U E N C E.
	// Select sequence 1.
	if X == 8 && Y == 0 {

		this.SelectedSequence = 0

		if debug {
			fmt.Printf("Select Sequence %d \n", this.SelectedSequence)
		}

		HandleSelect(sequences, this, eventsForLauchpad, commandChannels, fixturesConfig, guiButtons)
		cmd := common.Command{
			Action: common.PlayStaticOnce,
		}
		common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

		this.EditSequenceColorsMode[this.SelectedSequence] = false
		this.EditGoboSelectionMode[this.SelectedSequence] = false

		return
	}

	// Select sequence 2.
	if X == 8 && Y == 1 {

		this.SelectedSequence = 1

		if debug {
			fmt.Printf("Select Sequence %d \n", this.SelectedSequence)
		}

		HandleSelect(sequences, this, eventsForLauchpad, commandChannels, fixturesConfig, guiButtons)

		cmd := common.Command{
			Action: common.PlayStaticOnce,
		}
		common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

		this.EditSequenceColorsMode[this.SelectedSequence] = false
		this.EditGoboSelectionMode[this.SelectedSequence] = false

		return
	}

	// Select sequence 3.
	if X == 8 && Y == 2 {

		this.SelectedSequence = 2

		if debug {
			fmt.Printf("Select Sequence %d \n", this.SelectedSequence)
		}

		HandleSelect(sequences, this, eventsForLauchpad, commandChannels, fixturesConfig, guiButtons)

		cmd := common.Command{
			Action: common.PlayStaticOnce,
		}
		common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

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

		HandleSelect(sequences, this, eventsForLauchpad, commandChannels, fixturesConfig, guiButtons)

		cmd := common.Command{
			Action: common.PlayStaticOnce,
		}
		common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

		this.EditSequenceColorsMode[this.SelectedSequence] = false
		this.EditGoboSelectionMode[this.SelectedSequence] = false

		return
	}

	// Start sequence.
	if X == 8 && Y == 5 {

		if debug {
			fmt.Printf("Start Sequence %d \n", Y)
		}

		sequences[this.SelectedSequence].MusicTrigger = false
		cmd := common.Command{
			Action: common.Start,
			Args: []common.Arg{
				{Name: "Speed", Value: this.SequenceSpeed},
			},
		}
		common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)
		common.LightLamp(common.ALight{X: X, Y: Y, Brightness: full, Red: 255, Green: 0, Blue: 0}, eventsForLauchpad, guiButtons)
		time.Sleep(100 * time.Millisecond)
		common.LightLamp(common.ALight{X: X, Y: Y, Brightness: full, Red: 255, Green: 255, Blue: 255}, eventsForLauchpad, guiButtons)
		return
	}

	// Stop sequence.
	if X == 8 && Y == 6 {

		if debug {
			fmt.Printf("Stop Sequence %d \n", this.SelectedSequence)
		}

		cmd := common.Command{
			Action: common.Stop,
			Args: []common.Arg{
				{Name: "Speed", Value: this.SequenceSpeed},
			},
		}
		common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)
		common.LightLamp(common.ALight{X: X, Y: Y, Brightness: full, Red: 255, Green: 0, Blue: 0}, eventsForLauchpad, guiButtons)
		time.Sleep(100 * time.Millisecond)
		common.LightLamp(common.ALight{X: X, Y: Y, Brightness: full, Red: 255, Green: 255, Blue: 255}, eventsForLauchpad, guiButtons)
		return
	}

	// Size decrease.
	if X == 4 && Y == 7 {

		if debug {
			fmt.Printf("Decrease Size\n")
		}

		// Send Update RGB Size.
		this.Size--
		if this.Size < 1 {
			this.Size = 1
		}
		cmd := common.Command{
			Action: common.UpdateSize,
			Args: []common.Arg{
				{Name: "Size", Value: this.Size},
			},
		}
		common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

		// Send Update Scanner Size.
		this.ScannerSize = this.ScannerSize - 10
		if this.ScannerSize < 10 {
			this.ScannerSize = 10
		}
		cmd = common.Command{
			Action: common.UpdateScannerSize,
			Args: []common.Arg{
				{Name: "ScannerSize", Value: this.ScannerSize},
			},
		}
		common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

		return
	}

	// Increase Size.
	if X == 5 && Y == 7 {

		if debug {
			fmt.Printf("Increase Size\n")
		}

		// Send Update RGB Size.
		this.Size++
		if this.Size > 25 {
			this.Size = 25
		}
		cmd := common.Command{
			Action: common.UpdateSize,
			Args: []common.Arg{
				{Name: "Size", Value: this.Size},
			},
		}
		common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

		// Send Update Scanner Size.
		this.ScannerSize = this.ScannerSize + 10
		if this.ScannerSize > 120 {
			this.ScannerSize = 120
		}
		cmd = common.Command{
			Action: common.UpdateScannerSize,
			Args: []common.Arg{
				{Name: "ScannerSize", Value: this.ScannerSize},
			},
		}
		common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

		return
	}

	// Fade time decrease.
	if X == 6 && Y == 7 {

		if debug {
			fmt.Printf("Decrease Fade Time\n")
		}

		this.FadeSpeed--
		if this.FadeSpeed < 0 {
			this.FadeSpeed = 0
		}
		// Send fade update command.
		cmd := common.Command{
			Action: common.DecreaseFade,
			Args: []common.Arg{
				{Name: "FadeSpeed", Value: this.FadeSpeed},
			},
		}
		common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

		// Fade also send more or less coordinates for the scanner patterns.
		this.SelectedCordinates--
		if this.SelectedCordinates < 0 {
			this.SelectedCordinates = 0
		}
		cmd = common.Command{
			Action: common.UpdateNumberCoordinates,
			Args: []common.Arg{
				{Name: "NumberCoordinates", Value: this.SelectedCordinates},
			},
		}
		common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

		return
	}

	// Fade time increase.
	if X == 7 && Y == 7 {

		if debug {
			fmt.Printf("Increase Fade Time\n")
		}

		this.FadeSpeed++
		if this.FadeSpeed > 20 {
			this.FadeSpeed = 20
		}
		// Send fade update command.
		cmd := common.Command{
			Action: common.IncreaseFade,
			Args: []common.Arg{
				{Name: "FadeSpeed", Value: this.FadeSpeed},
			},
		}
		common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

		// Fade also send more or less coordinates for the scanner patterns.
		this.SelectedCordinates++
		if this.SelectedCordinates > 3 {
			this.SelectedCordinates = 3
		}
		cmd = common.Command{
			Action: common.UpdateNumberCoordinates,
			Args: []common.Arg{
				{Name: "NumberCoordinates", Value: this.SelectedCordinates},
			},
		}
		common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

		return
	}

	// F U N C T I O N  K E Y S
	if X >= 0 && X < 8 &&
		this.FunctionSelectMode[this.SelectedSequence] &&
		!this.EditPattenMode[this.SelectedSequence] &&
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
		common.ShowFunctionButtons(*sequences[this.SelectedSequence], this.SelectedSequence, eventsForLauchpad, guiButtons)

		// Now some functions mean that we go into another menu ( set of buttons )
		// This is true for :-
		// Function 1 - setting the patten.
		// Function 5 - setting the sequence colors or selecting scanner color.
		// Function 6 - setting the static colors or selecting scanner gobo.

		// Map Function 1 to patten mode.
		this.EditPattenMode[this.SelectedSequence] = sequences[this.SelectedSequence].Functions[common.Function1_Patten].State

		// Go straight into patten select mode, don't wait for a another select press.
		if this.EditPattenMode[this.SelectedSequence] {
			time.Sleep(500 * time.Millisecond) // But give the launchpad time to light the function key purple.
			common.HideFunctionButtons(this.SelectedSequence, eventsForLauchpad, guiButtons)
			this.EditFixtureSelectionMode = false
			ShowPattenSelectionButtons(this.SelectedSequence, *sequences[this.SelectedSequence], eventsForLauchpad, guiButtons)
		}

		// Map Function 5 to color edit.
		this.EditSequenceColorsMode[this.SelectedSequence] = sequences[this.SelectedSequence].Functions[common.Function5_Color].State

		// Go straight into color edit mode, don't wait for a another select press.
		if this.EditSequenceColorsMode[this.SelectedSequence] && sequences[this.SelectedSequence].Type == "rgb" {
			time.Sleep(500 * time.Millisecond) // But give the launchpad time to light the function key purple.
			common.HideFunctionButtons(this.SelectedSequence, eventsForLauchpad, guiButtons)
			ShowRGBColorSelectionButtons(this.SelectedSequence, *sequences[this.SelectedSequence], eventsForLauchpad, guiButtons)
		}

		// Go straight into color edit mode via select fixture, don't wait for a another select press.
		if this.EditSequenceColorsMode[this.SelectedSequence] && sequences[this.SelectedSequence].Type == "scanner" {
			time.Sleep(500 * time.Millisecond) // But give the launchpad time to light the function key purple.
			common.HideFunctionButtons(this.SelectedSequence, eventsForLauchpad, guiButtons)
			this.EditFixtureSelectionMode = true
			this.FollowingAction = "ShowScannerColorSelectionButtons"
			this.SelectedFixture = ShowSelectFixtureButtons(*sequences[this.SelectedSequence], this, eventsForLauchpad, fixturesConfig, this.FollowingAction, guiButtons)
		}

		// Map Function 6 to select gobo mode.
		if sequences[this.SelectedSequence].Type == "scanner" {
			this.EditGoboSelectionMode[this.SelectedSequence] = sequences[this.SelectedSequence].Functions[common.Function6_Gobo].State
		}

		// Go straight to gobo selection mode via select fixture, don't wait for a another select press.
		if this.EditGoboSelectionMode[this.SelectedSequence] {
			time.Sleep(500 * time.Millisecond) // But give the launchpad time to light the function key purple.
			common.HideFunctionButtons(this.SelectedSequence, eventsForLauchpad, guiButtons)
			this.EditFixtureSelectionMode = true
			this.FollowingAction = "ShowGoboSelectionButtons"
			this.SelectedFixture = ShowSelectFixtureButtons(*sequences[this.SelectedSequence], this, eventsForLauchpad, fixturesConfig, this.FollowingAction, guiButtons)
		}

		// Map Function 6 to static color edit.
		if sequences[this.SelectedSequence].Type != "scanner" {
			this.EditStaticColorsMode[this.SelectedSequence] = sequences[this.SelectedSequence].Functions[common.Function6_Static].State
		}

		// Go straight into static color edit mode, don't wait for a another select press.
		if this.EditStaticColorsMode[this.SelectedSequence] {
			//time.Sleep(2000 * time.Millisecond) // But give the launchpad time to light the function key purple.
			//common.HideFunctionButtons(this.SelectedSequence, eventsForLauchpad)
			this.FunctionSelectMode[this.SelectedSequence] = false
			// The sequence will automatically display the colors now!
		}

		return
	}

	// S W I T C H   B U T T O N's Toggle State of switches for this sequence.
	if X >= 0 && X < 8 && !this.FunctionSelectMode[this.SelectedSequence] &&
		Y >= 0 &&
		Y < 4 &&
		sequences[Y].Type == "switch" {

		if debug {
			fmt.Printf("Switch Key X:%d Y:%d\n", X, Y)
		}

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
			common.SendCommandToAllSequenceOfType(sequences, Y, cmd, commandChannels, "switch")
		}
	}

	// D I S A B L E   F I X T U R E  - Used to toggle the scanner on or off.
	if X >= 0 && X < 8 && !this.FunctionSelectMode[this.SelectedSequence] &&
		Y >= 0 &&
		Y < 4 &&
		!sequences[this.SelectedSequence].Functions[common.Function1_Patten].State &&
		!sequences[this.SelectedSequence].Functions[common.Function6_Static].State &&
		!sequences[this.SelectedSequence].Functions[common.Function5_Color].State &&
		sequences[Y].Type == "scanner" {

		if debug {
			fmt.Printf("Disable Fixture X:%d Y:%d\n", X, Y)
		}

		if !this.DisabledFixture[X][Y] && X < sequences[Y].NumberScanners {

			if debug {
				fmt.Printf("Toggle Scanner Number %d State on Sequence %d to true [Scanners:%d]\n", X, Y, sequences[Y].NumberScanners)
			}

			this.DisabledFixture[X][Y] = true

			// Tell the sequence to turn off this scanner.
			cmd := common.Command{
				Action: common.ToggleFixtureState,
				Args: []common.Arg{
					{Name: "SequenceNumber", Value: Y},
					{Name: "FixtureNumber", Value: X},
					{Name: "FixtureState", Value: true},
				},
			}
			common.SendCommandToSequence(Y, cmd, commandChannels)

			// Turn off the lamp.
			common.LightLamp(common.ALight{X: X, Y: Y, Brightness: 0, Red: 0, Green: 0, Blue: 0}, eventsForLauchpad, guiButtons)

			return

		}

		if this.DisabledFixture[X][Y] && X < sequences[Y].NumberScanners {
			if debug {
				fmt.Printf("Toggle Scanner Number %d State on Sequence %d to false\n", X, Y)
			}

			this.DisabledFixture[X][Y] = false

			// Tell the sequence to turn on this scanner.
			cmd := common.Command{
				Action: common.ToggleFixtureState,
				Args: []common.Arg{
					{Name: "SequenceNumber", Value: Y},
					{Name: "FixtureNumber", Value: X},
					{Name: "FixtureState", Value: false},
				},
			}
			common.SendCommandToSequence(Y, cmd, commandChannels)

			// Turn the lamp on.
			common.LightLamp(common.ALight{X: X, Y: Y, Brightness: full, Red: 255, Green: 255, Blue: 255}, eventsForLauchpad, guiButtons)

			return
		}

	}

	// F L A S H   B U T T O N S - Briefly light (flash) the fixtures based on current patten.
	if X >= 0 &&
		X < 8 &&
		Y >= 0 &&
		Y < 4 &&
		!sequences[Y].Functions[common.Function1_Patten].State &&
		!sequences[Y].Functions[common.Function6_Static].State &&
		!sequences[Y].Functions[common.Function5_Color].State &&
		sequences[Y].Type != "switch" && // As long as we're not a switch sequence.
		sequences[Y].Type != "scanner" && // As long as we're not a scanner sequence.
		!this.FunctionSelectMode[Y] { // As long as we're not a scanner sequence.

		if debug {
			fmt.Printf("Flash Button X:%d Y:%d\n", X, Y)
		}

		flashSequence := common.Sequence{
			Patten: common.Patten{
				Name:  "colors",
				Steps: this.Pattens["colors"].Steps,
			},
		}

		red := flashSequence.Patten.Steps[X].Fixtures[X].Colors[0].R
		green := flashSequence.Patten.Steps[X].Fixtures[X].Colors[0].G
		blue := flashSequence.Patten.Steps[X].Fixtures[X].Colors[0].B
		pan := flashSequence.Patten.Steps[X].Fixtures[X].Pan
		tilt := flashSequence.Patten.Steps[X].Fixtures[X].Tilt
		shutter := flashSequence.Patten.Steps[X].Fixtures[X].Shutter
		gobo := flashSequence.Patten.Steps[X].Fixtures[X].Gobo

		common.LightLamp(common.ALight{X: X, Y: Y, Brightness: this.MasterBrightness, Red: red, Green: green, Blue: blue}, eventsForLauchpad, guiButtons)
		fixture.MapFixtures(Y, dmxController, X, red, green, blue, pan, tilt, shutter, gobo, nil, fixturesConfig, this.Blackout, this.MasterBrightness, this.MasterBrightness)
		time.Sleep(200 * time.Millisecond)
		common.LightLamp(common.ALight{X: X, Y: Y, Brightness: 0, Red: 0, Green: 0, Blue: 0}, eventsForLauchpad, guiButtons)
		fixture.MapFixtures(Y, dmxController, X, 0, 0, 0, pan, tilt, shutter, gobo, nil, fixturesConfig, this.Blackout, this.MasterBrightness, this.MasterBrightness)
		return
	}

	// S E L E C T   P O S I T I O N
	// UP ARROW
	if X == 0 && Y == -1 && sequences[this.SelectedSequence].Type == "scanner" {

		if debug {
			fmt.Printf("UP ARROW\n")
		}

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
	}

	// DOWN ARROW
	if X == 1 && Y == -1 && sequences[this.SelectedSequence].Type == "scanner" {
		if debug {
			fmt.Printf("DOWN ARROW\n")
		}
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
	}

	// LEFT ARROW
	if X == 2 && Y == -1 && sequences[this.SelectedSequence].Type == "scanner" {
		if debug {
			fmt.Printf("LEFT ARROW\n")
		}
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
	}

	// RIGHT ARROW
	if X == 3 && Y == -1 && sequences[this.SelectedSequence].Type == "scanner" {
		if debug {
			fmt.Printf("RIGHT ARROW\n")
		}
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
	}

	// S E L E C T   S T A T I C   C O L O R
	// Red
	if X == 1 && Y == -1 && sequences[this.SelectedSequence].Type != "scanner" {

		if sequences[this.SelectedSequence].Functions[common.Function6_Static].State {

			if debug {
				fmt.Printf("Choose Static Red X:%d Y:%d\n", X, Y)
			}

			this.StaticButtons[this.SelectedSequence].X = this.LastStaticColorButtonX
			this.StaticButtons[this.SelectedSequence].Y = this.LastStaticColorButtonY

			code := common.GetLaunchPadColorCodeByRGB(this.StaticButtons[this.SelectedSequence].Color)
			common.FlashLight(this.LastStaticColorButtonX, this.LastStaticColorButtonY, int(code), 0x0, eventsForLauchpad, guiButtons)

			if this.StaticButtons[this.SelectedSequence].Color.R > 254 {
				this.StaticButtons[this.SelectedSequence].Color.R = 0
			} else {
				this.StaticButtons[this.SelectedSequence].Color.R = this.StaticButtons[this.SelectedSequence].Color.R + 10
			}
			common.LightLamp(common.ALight{X: X, Y: Y, Brightness: full, Red: this.StaticButtons[this.SelectedSequence].Color.R, Green: 0, Blue: 0}, eventsForLauchpad, guiButtons)
			updateStaticLamp(this.SelectedSequence, this.StaticButtons[this.SelectedSequence], commandChannels)
			return
		}
	}

	// Green
	if X == 2 && Y == -1 && sequences[this.SelectedSequence].Type != "scanner" {

		if sequences[this.SelectedSequence].Functions[common.Function6_Static].State {
			if debug {
				fmt.Printf("Choose Static Green X:%d Y:%d\n", X, Y)
			}

			this.StaticButtons[this.SelectedSequence].X = this.LastStaticColorButtonX
			this.StaticButtons[this.SelectedSequence].Y = this.LastStaticColorButtonY
			code := common.GetLaunchPadColorCodeByRGB(this.StaticButtons[this.SelectedSequence].Color)
			common.FlashLight(this.LastStaticColorButtonX, this.LastStaticColorButtonY, int(code), 0x0, eventsForLauchpad, guiButtons)

			if this.StaticButtons[this.SelectedSequence].Color.G > 254 {
				this.StaticButtons[this.SelectedSequence].Color.G = 0
			} else {
				this.StaticButtons[this.SelectedSequence].Color.G = this.StaticButtons[this.SelectedSequence].Color.G + 10
			}
			common.LightLamp(common.ALight{X: X, Y: Y, Brightness: full, Red: 0, Green: this.StaticButtons[this.SelectedSequence].Color.G, Blue: 0}, eventsForLauchpad, guiButtons)
			updateStaticLamp(this.SelectedSequence, this.StaticButtons[this.SelectedSequence], commandChannels)
			return
		}
	}

	// Blue
	if X == 3 && Y == -1 && sequences[this.SelectedSequence].Type != "scanner" {

		if sequences[this.SelectedSequence].Functions[common.Function6_Static].State {
			if debug {
				fmt.Printf("Choose Static Blue X:%d Y:%d\n", X, Y)
			}

			this.StaticButtons[this.SelectedSequence].X = this.LastStaticColorButtonX
			this.StaticButtons[this.SelectedSequence].Y = this.LastStaticColorButtonY
			code := common.GetLaunchPadColorCodeByRGB(this.StaticButtons[this.SelectedSequence].Color)
			common.FlashLight(this.LastStaticColorButtonX, this.LastStaticColorButtonY, int(code), 0x0, eventsForLauchpad, guiButtons)

			if this.StaticButtons[this.SelectedSequence].Color.B > 254 {
				this.StaticButtons[this.SelectedSequence].Color.B = 0
			} else {
				this.StaticButtons[this.SelectedSequence].Color.B = this.StaticButtons[this.SelectedSequence].Color.B + 10
			}
			common.LightLamp(common.ALight{X: X, Y: Y, Brightness: full, Red: 0, Green: 0, Blue: this.StaticButtons[this.SelectedSequence].Color.B}, eventsForLauchpad, guiButtons)
			updateStaticLamp(this.SelectedSequence, this.StaticButtons[this.SelectedSequence], commandChannels)
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
		sequences[this.SelectedSequence].CurrentSequenceColors = sequences[this.SelectedSequence].SequenceColors

		// We call ShowRGBColorSelectionButtons here so the selections will flash as you press them.
		ShowRGBColorSelectionButtons(this.SelectedSequence, *sequences[this.SelectedSequence], eventsForLauchpad, guiButtons)

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

		scannerColor := X

		// Set the scanner color for this sequence.
		cmd := common.Command{
			Action: common.UpdateScannerColor,
			Args: []common.Arg{
				{Name: "SelectedColor", Value: scannerColor},
				{Name: "SelectedFixture", Value: this.SelectedFixture},
			},
		}
		common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

		this.EditScannerColorsMode[this.SelectedSequence] = true

		// Get an upto date copy of the sequence.
		sequences[this.SelectedSequence] = common.RefreshSequence(this.SelectedSequence, commandChannels, updateChannels)

		// Set the colors.
		sequences[this.SelectedSequence].CurrentSequenceColors = sequences[this.SelectedSequence].SequenceColors

		// If the sequence isn't running this will force a single color DMX message.
		fixture.MapFixturesColorOnly(sequences[this.SelectedSequence], dmxController, fixturesConfig, scannerColor)

		// We call ShowScannerColorSelectionButtons here so the selections will flash as you press them.
		ShowScannerColorSelectionButtons(*sequences[this.SelectedSequence], this, eventsForLauchpad, fixturesConfig, guiButtons)

		return
	}

	// S E L E C T   F I X T U R E
	if X >= 0 && X < 8 && Y != -1 &&
		this.EditFixtureSelectionMode &&
		sequences[this.SelectedSequence].Functions[common.Function6_Gobo].State || sequences[this.SelectedSequence].Functions[common.Function5_Color].State &&
		sequences[this.SelectedSequence].Type == "scanner" {

		this.SelectedFixture = X

		if debug {
			fmt.Printf("Selected Fixture is %d \n", this.SelectedFixture)
		}

		// Update the buttons.
		if this.FollowingAction == "ShowGoboSelectionButtons" {
			ShowGoboSelectionButtons(*sequences[this.SelectedSequence], this, eventsForLauchpad, guiButtons)
		}
		if this.FollowingAction == "ShowScannerColorSelectionButtons" {
			err := ShowScannerColorSelectionButtons(*sequences[this.SelectedSequence], this, eventsForLauchpad, fixturesConfig, guiButtons)
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
		!this.EditFixtureSelectionMode &&
		sequences[this.SelectedSequence].Functions[common.Function6_Gobo].State &&
		sequences[this.SelectedSequence].Type == "scanner" {

		this.SelectedGobo = X + 1

		if debug {
			fmt.Printf("Sequence %d Set Gobo %d\n", this.SelectedSequence, this.SelectedGobo)
		}

		// Add the selected gobo to the sequence.
		cmd := common.Command{
			Action: common.UpdateGobo,
			Args: []common.Arg{
				{Name: "SelectedGobo", Value: this.SelectedGobo},
			},
		}
		common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

		this.EditGoboSelectionMode[this.SelectedSequence] = true

		// Get an upto date copy of the sequence.
		sequences[this.SelectedSequence] = common.RefreshSequence(this.SelectedSequence, commandChannels, updateChannels)

		// Set the colors.
		sequences[this.SelectedSequence].CurrentSequenceColors = sequences[this.SelectedSequence].SequenceColors

		// If the sequence isn't running this will force a single gobo DMX message.
		fixture.MapFixturesColorOnly(sequences[this.SelectedSequence], dmxController, fixturesConfig, this.SelectedGobo)

		// We call ShowGoboSelectionButtons here so the selections will flash as you press them.
		ShowGoboSelectionButtons(*sequences[this.SelectedSequence], this, eventsForLauchpad, guiButtons)

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

		if debug {
			fmt.Printf("Set Static Color X:%d Y:%d\n", X, Y)
		}

		// For this button increment the color.
		sequences[this.SelectedSequence].StaticColors[X].X = X
		sequences[this.SelectedSequence].StaticColors[X].Y = Y
		sequences[this.SelectedSequence].StaticColors[X].SelectedColor++
		if sequences[this.SelectedSequence].StaticColors[X].SelectedColor > 10 {
			sequences[this.SelectedSequence].StaticColors[X].SelectedColor = 0
		}
		sequences[this.SelectedSequence].StaticColors[X].Color = common.GetColorButtonsArray(sequences[this.SelectedSequence].StaticColors[X].SelectedColor)

		// Store the color data to allow for editing of static colors.
		this.StaticButtons[this.SelectedSequence].X = X
		this.StaticButtons[this.SelectedSequence].Y = Y
		this.StaticButtons[this.SelectedSequence].Color = sequences[this.SelectedSequence].StaticColors[X].Color
		this.StaticButtons[this.SelectedSequence].SelectedColor = sequences[this.SelectedSequence].StaticColors[X].SelectedColor

		// Tell the sequence about the new color and where we are in the
		// color cycle.
		cmd := common.Command{
			Action: common.UpdateStaticColor,
			Args: []common.Arg{
				{Name: "Static", Value: true},
				{Name: "StaticLamp", Value: X},
				{Name: "StaticLampFlash", Value: false},
				{Name: "SelectedColor", Value: sequences[this.SelectedSequence].StaticColors[X].SelectedColor},
				{Name: "StaticColor", Value: sequences[this.SelectedSequence].StaticColors[X].Color},
			},
		}
		common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)
		this.LastStaticColorButtonX = X
		this.LastStaticColorButtonY = Y
		code := common.GetLaunchPadColorCodeByRGB(this.StaticButtons[this.SelectedSequence].Color)
		common.FlashLight(this.LastStaticColorButtonX, this.LastStaticColorButtonY, int(code), 0x0, eventsForLauchpad, guiButtons)

		return
	}

	// S E L E C T   P A T T E N
	if X >= 0 && X < 8 && Y != -1 &&
		!this.EditFixtureSelectionMode &&
		this.EditPattenMode[this.SelectedSequence] {

		if debug {
			fmt.Printf("Set Patten X:%d Y:%d\n", X, Y)
		}

		// Tell the sequence to change the patten.
		cmd := common.Command{
			Action: common.SelectPatten,
			Args: []common.Arg{
				{Name: "SelectedPatten", Value: X},
			},
		}
		common.SendCommandToSequence(this.SelectedSequence, cmd, commandChannels)

		this.FunctionSelectMode[this.SelectedSequence] = false

		// Get an upto date copy of the sequence.
		sequences[this.SelectedSequence] = common.RefreshSequence(this.SelectedSequence, commandChannels, updateChannels)

		// We call ShowPattenSelectionButtons here so the selections will flash as you press them.
		this.EditFixtureSelectionMode = false
		ShowPattenSelectionButtons(this.SelectedSequence, *sequences[this.SelectedSequence], eventsForLauchpad, guiButtons)

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
			common.SendCommandToAllSequence(this.SelectedSequence, cmd, commandChannels)
			common.LightLamp(common.ALight{X: X, Y: Y, Brightness: full, Red: 0, Green: 0, Blue: 0}, eventsForLauchpad, guiButtons)
			common.FlashLight(8, 7, 0x03, 0x5f, eventsForLauchpad, guiButtons)
		} else {
			this.Blackout = false
			cmd := common.Command{
				Action: common.Normal,
			}
			common.SendCommandToAllSequence(this.SelectedSequence, cmd, commandChannels)
			common.LightLamp(common.ALight{X: X, Y: Y, Brightness: full, Red: 255, Green: 255, Blue: 255}, eventsForLauchpad, guiButtons)
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

func HandleSelect(sequences []*common.Sequence, this *CurrentState, eventsForLauchpad chan common.ALight,
	commandChannels []chan common.Command, fixtures *fixture.Fixtures, guiButtons chan common.ALight) {

	if debug {
		fmt.Printf("HANDLE: selectButtons[%d] = %t \n", this.SelectedSequence, this.SelectButtonPressed[this.SelectedSequence])
		fmt.Printf("HANDLE: this.FunctionSelectMode[%d] = %t \n", this.SelectedSequence, this.FunctionSelectMode[this.SelectedSequence])
		fmt.Printf("HANDLE: this.EditSequenceColorsMode[%d] = %t \n", this.SelectedSequence, this.EditSequenceColorsMode[this.SelectedSequence])
		fmt.Printf("HANDLE: this.EditStaticColorsMode[%d] = %t \n", this.SelectedSequence, this.EditStaticColorsMode[this.SelectedSequence])
		fmt.Printf("HANDLE: this.EditGoboSelectionMode[%d] = %t \n", this.SelectedSequence, this.EditGoboSelectionMode[this.SelectedSequence])
		fmt.Printf("HANDLE: this.FunctionSelectMode[%d] = %t \n", this.SelectedSequence, this.FunctionSelectMode[this.SelectedSequence])
		fmt.Printf("HANDLE: this.EditPattenMode[%d] = %t \n", this.SelectedSequence, this.EditPattenMode[this.SelectedSequence])
		fmt.Printf("HANDLE: Func Static[%d] = %t\n", this.SelectedSequence, sequences[this.SelectedSequence].Functions[common.Function6_Static].State)
	}

	// Light the sequence selector button.
	sequence.SequenceSelect(eventsForLauchpad, guiButtons, this.SelectedSequence)

	// First time into function mode we head back to normal mode.
	if this.FunctionSelectMode[this.SelectedSequence] && !this.SelectButtonPressed[this.SelectedSequence] &&
		!this.EditSequenceColorsMode[this.SelectedSequence] && !this.EditStaticColorsMode[this.SelectedSequence] {
		if debug {
			fmt.Printf("Handle 1 Function Bar off\n")
		}
		// Turn off function mode. Remove the function pads.
		common.HideFunctionButtons(this.SelectedSequence, eventsForLauchpad, guiButtons)

		this.FunctionSelectMode[this.SelectedSequence] = false

		if sequences[this.SelectedSequence].Functions[common.Function1_Patten].State {
			if debug {
				fmt.Printf("Show Patten Selection Buttons\n")
			}
			this.EditPattenMode[this.SelectedSequence] = true
			common.HideSequence(this.SelectedSequence, commandChannels)
			ShowPattenSelectionButtons(this.SelectedSequence, *sequences[this.SelectedSequence], eventsForLauchpad, guiButtons)
			return
		}

		if sequences[this.SelectedSequence].Functions[common.Function5_Color].State && sequences[this.SelectedSequence].Type == "rgb" {
			if debug {
				fmt.Printf("Show RGB Sequence Color Selection Buttons\n")
			}
			ShowRGBColorSelectionButtons(this.SelectedSequence, *sequences[this.SelectedSequence], eventsForLauchpad, guiButtons)
			return
		}

		if sequences[this.SelectedSequence].Functions[common.Function6_Static].State &&
			sequences[this.SelectedSequence].Type != "scanner" {
			if debug {
				fmt.Printf("Show Static Color Selection Buttons\n")
			}
			common.SetMode(this.SelectedSequence, commandChannels, "Static")
			return
		}

		// Allow us to exit the patten select mode without setting a patten.
		if this.EditPattenMode[this.SelectedSequence] {
			this.EditPattenMode[this.SelectedSequence] = false
		}

		// Switch off the gobo selection mode.
		sequences[this.SelectedSequence].Functions[common.Function6_Static].State = false

		// Else reveal the sequence on the launchpad keys
		if debug {
			fmt.Printf("Reveal Sequence\n")
		}
		common.RevealSequence(this.SelectedSequence, commandChannels)
		// Turn off the function mode flag.
		this.FunctionSelectMode[this.SelectedSequence] = false
		// Now forget we pressed twice and start again.
		this.SelectButtonPressed[this.SelectedSequence] = true

		return
	}

	// This the first time we have pressed the select button.
	if !this.SelectButtonPressed[this.SelectedSequence] &&
		!this.EditStaticColorsMode[this.SelectedSequence] {
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

		if sequences[this.SelectedSequence].Functions[common.Function1_Patten].State {
			// Reset the patten function key.
			sequences[this.SelectedSequence].Functions[common.Function1_Patten].State = false

			// Clear the patten function keys
			ClearPattenSelectionButtons(this.SelectedSequence, *sequences[this.SelectedSequence], eventsForLauchpad, guiButtons)

			// And reveal the sequence.
			common.RevealSequence(this.SelectedSequence, commandChannels)

			// Editing patten is over for this sequence.
			this.EditPattenMode[this.SelectedSequence] = false
		}

		if !this.FunctionSelectMode[this.SelectedSequence] && sequences[this.SelectedSequence].Functions[common.Function5_Color].State && this.EditSequenceColorsMode[this.SelectedSequence] {
			unSetEditSequenceColorsMode(sequences, this, commandChannels, eventsForLauchpad, guiButtons)
		}
		return
	}

	// Are we in function mode ?
	if this.FunctionSelectMode[this.SelectedSequence] {
		if debug {
			fmt.Printf("Handle 3\n")
		}
		// Turn off function mode. Remove the function pads.
		common.HideFunctionButtons(this.SelectedSequence, eventsForLauchpad, guiButtons)
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
		!this.FunctionSelectMode[this.SelectedSequence] && this.EditStaticColorsMode[this.SelectedSequence] { // The case when we leave static colors edit mode.

		if debug {
			fmt.Printf("Handle 4 - Function Bar On!\n")
		}

		// Unset the edit static color mode.
		this.EditStaticColorsMode[this.SelectedSequence] = false

		// Set function mode.
		this.FunctionSelectMode[this.SelectedSequence] = true

		// And hide the sequence so we can only see the function buttons.
		common.HideSequence(this.SelectedSequence, commandChannels)

		// Turn off any static sequence so we can see the functions.
		common.SetMode(this.SelectedSequence, commandChannels, "Sequence")

		// Create the function buttons.
		common.MakeFunctionButtons(*sequences[this.SelectedSequence], this.SelectedSequence, eventsForLauchpad, guiButtons, this.SequenceChannels)

		// Now forget we pressed twice and start again.
		this.SelectButtonPressed[this.SelectedSequence] = false

		return
	}
}

func unSetEditSequenceColorsMode(sequences []*common.Sequence, this *CurrentState, commandChannels []chan common.Command,
	eventsForLauchpad chan common.ALight, guiButtons chan common.ALight) {

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

	common.HideColorSelectionButtons(this.SelectedSequence, *sequences[this.SelectedSequence], this.SelectedSequence, eventsForLauchpad, guiButtons)
}

func AllFixturesOff(eventsForLauchpad chan common.ALight, guiButtons chan common.ALight, dmxController *ft232.DMXController, fixturesConfig *fixture.Fixtures) {
	for x := 0; x < 8; x++ {
		for y := 0; y < 4; y++ {
			common.LightLamp(common.ALight{X: x, Y: y, Brightness: 0, Red: 0, Green: 0, Blue: 0}, eventsForLauchpad, guiButtons)
			fixture.MapFixtures(y, dmxController, x, 0, 0, 0, 0, 0, 0, 0, nil, fixturesConfig, true, 0, 0)
		}
	}
}

// For the given sequence show the available sequence colors on the relevant buttons.
func ShowRGBColorSelectionButtons(mySequenceNumber int, sequence common.Sequence, eventsForLauchpad chan common.ALight, guiButtons chan common.ALight) {

	if debug {
		fmt.Printf("Show Color Selection Buttons\n")
	}
	// Check if we need to flash this button.
	for myFixtureNumber, lamp := range sequence.AvailableSequenceColors {

		for index, availableColor := range sequence.AvailableSequenceColors {
			for _, sequenceColor := range sequence.CurrentSequenceColors {
				if debug {
					fmt.Printf("myFixtureNumber %d   current color %d\n", myFixtureNumber, sequenceColor)
				}
				if availableColor.Color == sequenceColor {
					if myFixtureNumber == index {
						lamp.Flash = true
					}
				}
			}
		}
		if lamp.Flash {
			code := common.GetLaunchPadColorCodeByRGB(lamp.Color)
			common.FlashLight(myFixtureNumber, mySequenceNumber, int(code), 0x0, eventsForLauchpad, guiButtons)
		} else {
			common.LightLamp(common.ALight{X: myFixtureNumber, Y: mySequenceNumber, Brightness: full, Red: lamp.Color.R, Green: lamp.Color.G, Blue: lamp.Color.B}, eventsForLauchpad, guiButtons)
		}
	}
}

// For the given sequence show the available scanner selection colors on the relevant buttons.
func ShowSelectFixtureButtons(sequence common.Sequence, this *CurrentState, eventsForLauchpad chan common.ALight, fixtures *fixture.Fixtures, action string, guiButtons chan common.ALight) int {

	if debug {
		fmt.Printf("Sequence %d Show Fixture Selection Buttons on the way to %s\n", this.SelectedSequence, action)
	}

	for fixtureNumber, lamp := range sequence.AvailableFixtures {

		if debug {
			fmt.Printf("Lamp %+v\n", lamp)
		}
		if fixtureNumber == this.SelectedFixture {
			lamp.Flash = true
			this.SelectedFixture = fixtureNumber
		}
		if lamp.Flash {
			code := common.GetLaunchPadColorCodeByRGB(lamp.Color)
			common.FlashLight(fixtureNumber, this.SelectedSequence, int(code), 0x0, eventsForLauchpad, guiButtons)

		} else {
			common.LightLamp(common.ALight{X: fixtureNumber, Y: this.SelectedSequence, Red: lamp.Color.R, Green: lamp.Color.G, Blue: lamp.Color.B, Brightness: sequence.Master}, eventsForLauchpad, guiButtons)
		}
	}
	if debug {
		fmt.Printf("Selected Fixture is %d\n", this.SelectedFixture)
	}
	return this.SelectedFixture
}

// ShowGoboSelectionButtons puts up a set of red buttons used to select a fixture.
func ShowGoboSelectionButtons(sequence common.Sequence, this *CurrentState, eventsForLauchpad chan common.ALight, guiButtons chan common.ALight) {

	if debug {
		fmt.Printf("Sequence %d Show Gobo Selection Buttons\n", this.SelectedSequence)
	}
	// Check if we need to flash this button.
	for myFixtureNumber, lamp := range sequence.AvailableGoboSelectionColors {
		if debug {
			fmt.Printf("myFixtureNumber %d   current gobo %d\n", myFixtureNumber, sequence.SelectedGobo)
		}
		if myFixtureNumber == sequence.SelectedGobo-1 {
			lamp.Flash = true
		}
		if lamp.Flash {
			code := common.GetLaunchPadColorCodeByRGB(lamp.Color)
			common.FlashLight(myFixtureNumber, this.SelectedSequence, int(code), 0x0, eventsForLauchpad, guiButtons)
		} else {
			common.LightLamp(common.ALight{X: myFixtureNumber, Y: this.SelectedSequence, Brightness: sequence.Master, Red: lamp.Color.R, Green: lamp.Color.G, Blue: lamp.Color.B}, eventsForLauchpad, guiButtons)
		}
	}
}

// For the given sequence show the available scanner selection colors on the relevant buttons.
func ShowScannerColorSelectionButtons(sequence common.Sequence, this *CurrentState, eventsForLauchpad chan common.ALight, fixtures *fixture.Fixtures, guiButtons chan common.ALight) error {

	if debug {
		fmt.Printf("Show Scanner Color Selection Buttons,  Sequence is %d  fixture is %d   color is %d \n", this.SelectedSequence, this.SelectedFixture, sequence.ScannerColor[this.SelectedFixture])
	}

	// if there are no colors available for this fixture turn everything off and print an error.
	if sequence.AvailableScannerColors[this.SelectedFixture+1] == nil {
		for _, fixture := range fixtures.Fixtures {
			if fixture.Group == this.SelectedSequence+1 {
				common.LightLamp(common.ALight{X: fixture.Number - 1, Y: this.SelectedSequence, Red: 0, Green: 0, Blue: 0, Brightness: sequence.Master}, eventsForLauchpad, guiButtons)
			}
		}
		return fmt.Errorf("error: no colors available for fixture number %d", this.SelectedFixture+1)
	}

	// selected fixture is +1 here because the fixtures in the yaml config file start with 1 not 0.
	for fixtureNumber, lamp := range sequence.AvailableScannerColors[this.SelectedFixture+1] {

		if debug {
			fmt.Printf("Lamp %+v\n", lamp)
		}
		if fixtureNumber == sequence.ScannerColor[this.SelectedFixture] {
			lamp.Flash = true
		}
		if lamp.Flash {
			code := common.GetLaunchPadColorCodeByRGB(lamp.Color)
			common.FlashLight(fixtureNumber, this.SelectedSequence, int(code), 0x0, eventsForLauchpad, guiButtons)
		} else {
			common.LightLamp(common.ALight{X: fixtureNumber, Y: this.SelectedSequence, Red: lamp.Color.R, Green: lamp.Color.G, Blue: lamp.Color.B, Brightness: sequence.Master}, eventsForLauchpad, guiButtons)
		}
	}
	return nil
}

// For the given sequence clear the available this.Pattens on the relevant buttons.
func ClearPattenSelectionButtons(mySequenceNumber int, sequence common.Sequence, eventsForLauchpad chan common.ALight, guiButtons chan common.ALight) {
	// Check if we need to flash this button.
	for myFixtureNumber := 0; myFixtureNumber < 4; myFixtureNumber++ {
		common.LightLamp(common.ALight{X: myFixtureNumber, Y: mySequenceNumber, Red: 0, Green: 0, Blue: 0, Brightness: sequence.Master}, eventsForLauchpad, guiButtons)
	}
}

// For the given sequence show the available this.Pattens on the relevant buttons.
func ShowPattenSelectionButtons(mySequenceNumber int, sequence common.Sequence, eventsForLauchpad chan common.ALight, guiButtons chan common.ALight) {
	// Check if we need to flash this button.
	for myFixtureNumber := 0; myFixtureNumber < 4; myFixtureNumber++ {
		if myFixtureNumber == sequence.SelectedPatten {
			code := common.GetLaunchPadColorCodeByRGB(common.Color{R: 255, G: 255, B: 255})
			common.FlashLight(myFixtureNumber, mySequenceNumber, int(code), 0x0, eventsForLauchpad, guiButtons)
		} else {
			common.LightLamp(common.ALight{X: myFixtureNumber, Y: mySequenceNumber, Red: 255, Green: 255, Blue: 255, Brightness: sequence.Master}, eventsForLauchpad, guiButtons)
		}
	}
}
