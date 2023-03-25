// Copyright (C) 2022, 2023 dhowlett99.
// This is the dmxlights common functions.
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

package common

import (
	"fmt"
	"image/color"
	"math"
	"strings"
	"sync"
	"time"
)

const debug = false
const MaxDMXAddress = 512
const MaxTextEntryLength = 35
const DefaultScannerSize = 120
const MaxScannerSize = 120
const MaxRGBSize = 120
const MaxRGBFade = 10
const MaxColorBar = 9 // Eight colors and a default color bar.
const MaxDMXBrightness = 255
const DefaultPattern = 0
const DefaultRGBSize = 1
const DefaultRGBFade = 1
const DefaultScannerFade = 10
const DefaultSpeed = 7
const DefaultRGBShift = 0
const DefaultScannerShift = 0
const DefaultScannerCoordinates = 0
const ScannerMidPoint = 127
const DefaultRGBCoordinates = 10

var DefaultSequenceColors = []Color{{R: 0, G: 255, B: 0}}

type ALight struct {
	X                int
	Y                int
	Brightness       int
	Red              int
	Green            int
	Blue             int
	Flash            bool
	OnColor          Color
	OffColor         Color
	UpdateLabel      bool
	Label            string
	UpdateStatus     bool
	Status           string
	Which            string
	FlashStopChannel chan bool
	Hidden           bool
}

type Color struct {
	R     int
	G     int
	B     int
	W     int
	A     int
	UV    int
	Flash bool
}

// Used in calculating Positions.
type FixtureBuffer struct {
	Color        Color
	MasterDimmer int
	Brightness   int
	Gobo         int
	Pan          int
	Tilt         int
	Shutter      int
	Enabled      bool
}

type Value struct {
	Channel string
	Setting string
}

type Setting struct {
	Name  string
	Value int16
}

type State struct {
	Name        string
	Number      int16
	Label       string
	Values      []Value
	ButtonColor string
	Actions     []Action
	Flash       bool
}

type Action struct {
	Name    string
	Number  int
	Colors  []string
	Mode    string
	Fade    string
	Speed   string
	Rotate  string
	Music   string
	Program string
}

type Switch struct {
	Name         string
	Number       int
	Label        string
	CurrentState int
	Description  string
	States       []State
	Fixture      string
	UseFixture   string
}

type StaticColorButton struct {
	Name          string
	Label         string
	Number        int
	X             int
	Y             int
	Color         Color
	SelectedColor int
	Flash         bool
	Setting       int
	FirstPress    bool
}

type ScannerState struct {
	Enabled  bool
	Inverted bool
}

type Pattern struct {
	Name     string
	Label    string
	Number   int
	Length   int // 8, 4 or 2
	Size     int
	Fixtures int // 8 Fixtures
	Steps    []Step
}

type Arg struct {
	Name  string
	Value interface{}
}

// Command tells sequences what to do.
type Command struct {
	Action int
	Args   []Arg
}

// Valid Command Actions.
const (
	Actions int = iota
	Clear
	Reset
	UpdateMode
	UpdateStatic
	UpdateStaticColor
	UpdateSequenceColor
	PlayStaticOnce
	PlaySwitchOnce
	UnHide
	Hide
	Start
	Stop
	ReadConfig
	LoadConfig
	UpdateSpeed
	UpdatePattern
	UpdateRGBFadeSpeed
	UpdateRGBSize
	UpdateScannerSize
	Blackout
	Normal
	SoftFadeOn
	SoftFadeOff
	UpdateColor
	UpdateFunctionMode
	FunctionMode
	UpdateFunctions
	GetUpdatedSequence
	ClearAllSwitchPositions
	ResetAllSwitchPositions
	UpdateSwitch
	Inverted
	UpdateGobo
	Flood
	StopFlood
	Strobe
	StopStrobe
	UpdateAutoColor
	AutoColor
	UpdateAutoPattern
	AutoPattern
	ToggleFixtureState
	FixtureState
	UpdateRGBShift
	UpdateScannerShift
	UpdateScannerColor
	UpdateStrobeSpeed
	ClearSequenceColor
	ClearStaticColor
	SetStaticColorBar
	Static
	Master
	UpdateNumberCoordinates
	UpdateOffsetPan
	UpdateOffsetTilt
	EnableAllScanners
	UpdateScannerChase
)

// A full step cycle is 39 ticks ie 39 values.
// 13 fade up values, 13 on values and 13 off values.
const StepSize = 39

var Pink = Color{R: 255, G: 0, B: 255}
var White = Color{R: 255, G: 255, B: 255}
var Black = Color{R: 0, G: 0, B: 0}
var Red = Color{R: 255, G: 0, B: 0}
var Green = Color{R: 0, G: 255, B: 0}
var Blue = Color{R: 0, G: 0, B: 255}
var PresetYellow = Color{R: 150, G: 150, B: 0}
var Cyan = Color{R: 0, G: 255, B: 255}

type Gobo struct {
	Name    string
	Label   string
	Number  int
	Setting int
	Flash   bool
	Color   Color
}

// Sequence describes sequences.
type Sequence struct {
	Name                        string                      // Sequence name.
	Label                       string                      // Sequence label.
	Description                 string                      // Sequence description.
	Number                      int                         // Sequence number.
	Run                         bool                        // True if this sequence is running.
	Bounce                      bool                        // True if this sequence is bouncing.
	RGBInvert                   bool                        // True if RGB sequence patten is inverted.
	Hide                        bool                        // Hide is used to hide sequence buttons when using function keys.
	Type                        string                      // Type of sequnece, current valid values are :- rgb, scanner,  or switch.
	Master                      int                         // Master Brightness
	Strobe                      bool                        // Strobe is enabled.
	StrobeSpeed                 int                         // Strobe speed.
	Rotate                      int                         // Rotate speed.
	RGBShift                    int                         // RGB shift.
	CurrentSpeed                time.Duration               // Sequence speed represented as a duration.
	Speed                       int                         // Sequence speed represented by a short number.
	MusicTrigger                bool                        // Is this sequence in music trigger mode.
	LastMusicTrigger            bool                        // Save copy of music trigger.
	Blackout                    bool                        // Flag to indicate we're in blackout mode.
	CurrentColors               []Color                     // Storage for the colors in a sequence.
	SequenceColors              []Color                     // Temporay storage for changing sequence colors.
	Color                       int                         // Index into current sequnece colors.
	Steps                       []Step                      // RGB or Shutter steps in this  sequence.
	ScannerSteps                []Step                      // Pan & Tilt steps in this  sequence.
	NumberSteps                 int                         // Holds the number of steps this sequence has. Will change if you change size, fade times etc.
	NumberFixtures              int                         // Total Number of fixtures for this sequence.
	EnabledNumberFixtures       int                         // Enabled Number of fixtures for this sequence.
	RGBPositions                map[int]Position            // One set of Fixture positions for RGB devices. index is position number.
	ScannerPositions            map[int]map[int]Position    // Scanner Fixture positions decides where a fixture is in a give set of sequence steps. First index is fixure, second index is positions.
	AutoColor                   bool                        // Sequence is going to automatically change the color.
	AutoPattern                 bool                        // Sequence is going to automatically change the pattern.
	GuiFunctionLabels           [8]string                   // Storage for the function key labels for this sequence.
	GuiFixtureLabels            []string                    // Storage for the fixture labels. Used for scanner names.
	Pattern                     Pattern                     // Contains fixtures and RGB steps info.
	RGBAvailablePatterns        map[int]Pattern             // Available patterns for the RGB fixtures.
	RGBAvailableColors          []StaticColorButton         // Available colors for the RGB fixtures.
	RGBColor                    int                         // The selected RGB fixture color.
	FadeUpAndDown               []int                       // curve fade on and stay on and time to fade off
	FadeDownAndUp               []int                       // curve fade off and on again
	RGBFade                     int                         // RGB Fade time
	RGBSize                     int                         // RGB Fade size
	SavedSequenceColors         []Color                     // Used for updating the color in a sequence.
	RecoverSequenceColors       bool                        // Storage for recovering sequence colors, when you come out of automatic color change.
	SaveColors                  bool                        // Indicate we should save colors in this sequence. used for above.
	Mode                        string                      // Tells sequnece if we're in sequence (chase) or static (static colors) mode.
	StaticColors                []StaticColorButton         // Used in static color editing
	Clear                       bool                        // Clear all fixtures in this sequence.
	Static                      bool                        // We're a static sequence.
	PlayStaticOnce              bool                        // Play a static scene only once.
	PlaySwitchOnce              bool                        // Play a switch sequence scene only once.
	PlaySingleSwitch            bool                        // Play a single switch.
	StartFlood                  bool                        // We're in flood mode.
	StopFlood                   bool                        // We're not in flood mode.
	StartStrobe                 bool                        // We're in strobe mode.
	StopStrobe                  bool                        // We're not in strobe mode.
	FloodPlayOnce               bool                        // Play the flood sceme only once.
	FloodSelectedSequence       map[int]bool                // A map that remembers who is in flood mode.
	ScannerAvailableColorsMutex *sync.RWMutex               // Mutex to protect the scanner available colors map from syncronous access.
	ScannerAvailableColors      map[int][]StaticColorButton // Available colors for this scanner.
	ScannerAvailableGobos       map[int][]StaticColorButton // Available gobos for this scanner.
	ScannerAvailablePatterns    map[int]Pattern             // Available patterns for this scanner.
	ScannersAvailable           []StaticColorButton         // Holds a set of red buttons, one for every available fixture.
	SelectedPattern             int                         // The selected pattern.
	ScannerSize                 int                         // The selected scanner size.
	ScannerShift                int                         // Used for shifting scanners patterns apart.
	ScannerGobo                 map[int]int                 // Eight scanners per sequence, each can have their own gobo.
	ScannerChase                bool                        // Chase the scanner shutters instead of allways being on.
	ScannerInvert               bool                        // Invert the scanner, i.e scanner in the opposite direction.
	ScannerColor                map[int]int                 // Eight scanners per sequence, each can have their own color.
	ScannerCoordinates          []int                       // Number of scanner coordinates.
	ScannerSelectedCoordinates  int                         // index into scanner coordinates.
	ScannerOffsetPan            int                         // Offset for pan values.
	ScannerOffsetTilt           int                         // Offset for tilt values.
	ScannerState                map[int]ScannerState        // Map of fixtures which are disabled.
	DisableOnceMutex            *sync.RWMutex               // Lock to protect DisableOnce.
	DisableOnce                 map[int]bool                // Map used to play disable only once.
	UpdateSize                  bool                        // Command to update size.
	UpdateShift                 bool                        // Command to update the shift.
	UpdatePattern               bool                        // Flag to indicate we're going to change the RGB pattern.
	UpdateSequenceColor         bool                        // Command to update the sequence colors.
	Functions                   []Function                  // Storage for the sequence functions.
	FunctionMode                bool                        // This sequence is in function mode.
	Switches                    []Switch                    // A switch sequence stores its data in here.
	CurrentSwitch               int                         // Play this current switch position.
	Optimisation                bool                        // Flag to decide on calculatePositions Optimisation.
	RGBCoordinates              int                         // Number of coordinates in RGB fade.
}

type Function struct {
	Name           string
	SequenceNumber int
	Number         int
	State          bool
	Flash          bool
	Label          string
}

type Channels struct {
	CommmandChannels []chan Command
	ReplyChannels    []chan Sequence
	UpdateChannels   []chan Sequence
	SoundTriggers    []*Trigger
}

type SwitchChannel struct {
	Stop             chan bool
	StopRotate       chan bool
	KeepRotateAlive  chan bool
	SequencerRunning bool
	Blackout         bool
	Master           int
	SwitchPosition   int
}

type Hit struct {
	X int
	Y int
}

type Step struct {
	Fixtures map[int]Fixture
}

type FixtureCommand struct {
	Step           int
	NumberSteps    int
	Type           string
	SequenceNumber int

	// Common commands.
	Strobe      bool
	StrobeSpeed int
	Master      int
	Blackout    bool
	Hide        bool
	Clear       bool

	StartFlood bool
	StopFlood  bool

	// RGB commands.
	RGBPosition     Position
	RGBStatic       bool
	RGBStaticColors []StaticColorButton

	// Scanner Commands.
	ScannerColor             int
	ScannerPosition          Position
	ScannerState             ScannerState
	ScannerDisableOnce       bool
	ScannerChase             bool
	ScannerAvailableColors   []StaticColorButton
	ScannerGobo              int
	ScannerOffsetPan         int
	ScannerOffsetTilt        int
	ScannerNumberCoordinates int
	ScannerShutterPositions  map[int]Position

	// Derby Commands
	Rotate  int
	Music   int
	Program int
}

type Position struct {
	// RGB
	Fixtures map[int]Fixture
}

// A common fixture can have any or some of the
// following, depending if its a light or
// a scanner.
type Fixture struct {
	ID           string
	Name         string
	Label        string
	MasterDimmer int
	Brightness   int
	ScannerColor Color
	Colors       []Color
	Pan          int
	Tilt         int
	Shutter      int
	Rotate       int
	Music        int
	Gobo         int
	Program      int
	Enabled      bool
}

type ButtonPresets struct {
	X int
	Y int
}

type Event struct {
	Fixture   int
	Run       bool
	Stop      bool
	Start     bool
	Fadeup    bool
	Fadedown  bool
	Shift     int
	FadeTime  time.Duration
	LastColor Color
	Color     Color
}

type Trigger struct {
	Name    string
	State   bool
	Gain    float32
	Channel chan Command
}

// Define the function keys.
const (
	Function1_Pattern       = 0 // Set pattern mode.
	Function2_Auto_Color    = 1 // Auto Color change.
	Function3_Auto_Pattern  = 2 // Auto Pattern change
	Function4_Bounce        = 3 // Sequence auto reverses.  doesn't apply in scanner mode.
	Function5_Color         = 4 // Set RGB chase color. or select the scanner color.
	Function6_Static_Gobo   = 5 // Set static color / set scanner gobo.
	Function7_Invert_Chase  = 6 // Invert the RGB colors  / Set scanner chase mode.
	Function8_Music_Trigger = 7 // Music trigger on and off. Both RGB and scanners.
)

func SendCommandToSequence(selectedSequence int, command Command, commandChannels []chan Command) {
	commandChannels[selectedSequence] <- command
}

func SendCommandToAllSequence(command Command, commandChannels []chan Command) {
	commandChannels[0] <- command
	commandChannels[1] <- command
	commandChannels[2] <- command
	commandChannels[3] <- command
	commandChannels[4] <- command
}

func SendCommandToAllSequenceOfType(sequences []*Sequence, command Command, commandChannels []chan Command, Type string) {
	for index, s := range sequences {
		if s.Type == Type {
			commandChannels[index] <- command
		}
	}
}

func SendCommandToAllSequenceExcept(selectedSequence int, command Command, commandChannels []chan Command) {
	for index := range commandChannels {
		if index != selectedSequence {
			commandChannels[index] <- command
		}
	}
}

func MakeFunctionButtons(selectedSequence int, eventsForLauchpad chan ALight, guiButtons chan ALight, channels Channels) {

	// The target set of buttons.
	ClearSelectedRowOfButtons(selectedSequence, eventsForLauchpad, guiButtons)

	// Get an upto date copy of the sequence.
	cmd := Command{
		Action: ReadConfig,
	}
	SendCommandToSequence(selectedSequence, cmd, channels.CommmandChannels)

	replyChannel := channels.ReplyChannels[selectedSequence]
	sequence := <-replyChannel

	ShowFunctionButtons(sequence, selectedSequence, eventsForLauchpad, guiButtons)
}

func SetMode(selectedSequence int, commandChannels []chan Command, mode string) {
	cmd := Command{
		Action: UpdateMode,
		Args: []Arg{
			{Name: "Mode", Value: mode},
		},
	}
	SendCommandToSequence(selectedSequence, cmd, commandChannels)
}

func RevealSequence(selectedSequence int, commandChannels []chan Command) {
	cmd := Command{
		Action: UnHide,
	}
	SendCommandToSequence(selectedSequence, cmd, commandChannels)
}

func HideSequence(selectedSequence int, commandChannels []chan Command) {
	cmd := Command{
		Action: Hide,
	}
	SendCommandToSequence(selectedSequence, cmd, commandChannels)
}

// Colors are selected from a pallete of 8 colors, this function takes 0-9 (repeating 4 time) and
// returns the color array
func GetColorButtonsArray(color int) Color {

	switch color {
	case 0:
		return Color{R: 255, G: 0, B: 0} // Red
	case 1:
		return Color{R: 255, G: 111, B: 0} // Orange
	case 2:
		return Color{R: 255, G: 255, B: 0} // Yellow
	case 3:
		return Color{R: 0, G: 255, B: 0} // Green
	case 4:
		return Color{R: 0, G: 255, B: 255} // Cyan
	case 5:
		return Color{R: 0, G: 0, B: 255} // Blue
	case 6:
		return Color{R: 100, G: 0, B: 255} // Purple
	case 7:
		return Color{R: 255, G: 0, B: 255} // Pink
	case 8:
		return Color{R: 255, G: 255, B: 255} // White
	case 9:
		return Color{R: 0, G: 0, B: 0} // Black
	case 10:
		return Color{R: 255, G: 0, B: 0} // Red
	case 11:
		return Color{R: 255, G: 111, B: 0} // Orange
	case 12:
		return Color{R: 255, G: 255, B: 0} // Yellow
	case 13:
		return Color{R: 0, G: 255, B: 0} // Green
	case 14:
		return Color{R: 0, G: 255, B: 255} // Cyan
	case 15:
		return Color{R: 0, G: 0, B: 255} // Blue
	case 16:
		return Color{R: 100, G: 0, B: 255} // Purple
	case 17:
		return Color{R: 255, G: 0, B: 255} // Pink
	case 18:
		return Color{R: 255, G: 255, B: 255} // White
	case 19:
		return Color{R: 0, G: 0, B: 0} // Black
	case 20:
		return Color{R: 255, G: 0, B: 0} // Red
	case 21:
		return Color{R: 255, G: 111, B: 0} // Orange
	case 22:
		return Color{R: 255, G: 255, B: 0} // Yellow
	case 23:
		return Color{R: 0, G: 255, B: 0} // Green
	case 24:
		return Color{R: 0, G: 255, B: 255} // Cyan
	case 25:
		return Color{R: 0, G: 0, B: 255} // Blue
	case 26:
		return Color{R: 100, G: 0, B: 255} // Purple
	case 27:
		return Color{R: 255, G: 0, B: 255} // Pink
	case 28:
		return Color{R: 255, G: 255, B: 255} // White
	case 29:
		return Color{R: 0, G: 0, B: 0} // Black
	case 30:
		return Color{R: 255, G: 0, B: 0} // Red
	case 31:
		return Color{R: 255, G: 111, B: 0} // Orange
	case 32:
		return Color{R: 255, G: 255, B: 0} // Yellow
	case 33:
		return Color{R: 0, G: 255, B: 0} // Green
	case 34:
		return Color{R: 0, G: 255, B: 255} // Cyan
	case 35:
		return Color{R: 0, G: 0, B: 255} // Blue
	case 36:
		return Color{R: 100, G: 0, B: 255} // Purple
	case 37:
		return Color{R: 255, G: 0, B: 255} // Pink
	case 38:
		return Color{R: 255, G: 255, B: 255} // White
	case 39:
		return Color{R: 0, G: 0, B: 0} // Black
	case 40:
		return Color{R: 255, G: 0, B: 0} // Red
	case 41:
		return Color{R: 255, G: 111, B: 0} // Orange
	case 42:
		return Color{R: 255, G: 255, B: 0} // Yellow
	case 43:
		return Color{R: 0, G: 255, B: 0} // Green
	case 44:
		return Color{R: 0, G: 255, B: 255} // Cyan
	case 45:
		return Color{R: 0, G: 0, B: 255} // Blue
	case 46:
		return Color{R: 100, G: 0, B: 255} // Purple
	case 47:
		return Color{R: 255, G: 0, B: 255} // Pink
	case 48:
		return Color{R: 255, G: 255, B: 255} // White
	case 49:
		return Color{R: 0, G: 0, B: 0} // Black
	case 50:
		return Color{R: 255, G: 0, B: 0} // Red
	case 51:
		return Color{R: 255, G: 111, B: 0} // Orange
	case 52:
		return Color{R: 255, G: 255, B: 0} // Yellow
	case 53:
		return Color{R: 0, G: 255, B: 0} // Green
	case 54:
		return Color{R: 0, G: 255, B: 255} // Cyan
	case 55:
		return Color{R: 0, G: 0, B: 255} // Blue
	case 56:
		return Color{R: 100, G: 0, B: 255} // Purple
	case 57:
		return Color{R: 255, G: 0, B: 255} // Pink
	case 58:
		return Color{R: 255, G: 255, B: 255} // White
	case 59:
		return Color{R: 0, G: 0, B: 0} // Black
	case 60:
		return Color{R: 255, G: 0, B: 0} // Red
	case 61:
		return Color{R: 255, G: 111, B: 0} // Orange
	case 62:
		return Color{R: 255, G: 255, B: 0} // Yellow
	case 63:
		return Color{R: 0, G: 255, B: 0} // Green
	case 64:
		return Color{R: 0, G: 255, B: 255} // Cyan
	case 65:
		return Color{R: 0, G: 0, B: 255} // Blue
	case 66:
		return Color{R: 100, G: 0, B: 255} // Purple
	case 67:
		return Color{R: 255, G: 0, B: 255} // Pink
	case 68:
		return Color{R: 255, G: 255, B: 255} // White
	case 69:
		return Color{R: 0, G: 0, B: 0} // Black
	}
	return Color{}
}

func GetColorArrayByNames(names []string) ([]Color, error) {

	colors := []Color{}
	for _, color := range names {
		// Find the color by name from the library of supported colors.
		colorLibrary, err := GetRGBColorByName(color)
		if err != nil {
			return colors, err
		}
		newColor := colorLibrary

		// Add the color to the chase colors.
		colors = append(colors, newColor)
	}
	return colors, nil
}

// Convert my common.Color RGB into color.NRGBA used by the fyne.io GUI library.
func ConvertRGBtoNRGBA(alight Color) color.NRGBA {
	NRGBAcolor := color.NRGBA{}
	NRGBAcolor.R = uint8(alight.R)
	NRGBAcolor.G = uint8(alight.G)
	NRGBAcolor.B = uint8(alight.B)
	NRGBAcolor.A = 255
	return NRGBAcolor
}

func GetRGBColorByName(color string) (Color, error) {
	switch color {
	case "Red":
		return Color{R: 255, G: 0, B: 0}, nil

	case "Orange":
		return Color{R: 255, G: 111, B: 0}, nil

	case "Yellow":
		return Color{R: 255, G: 255, B: 0}, nil

	case "Green":
		return Color{R: 0, G: 255, B: 0}, nil

	case "Cyan":
		return Color{R: 0, G: 255, B: 255}, nil

	case "Blue":
		return Color{R: 0, G: 0, B: 255}, nil

	case "Purple":
		return Color{R: 100, G: 0, B: 255}, nil

	case "Pink":
		return Color{R: 255, G: 0, B: 255}, nil

	case "White":
		return Color{R: 255, G: 255, B: 255}, nil

	case "Light Blue":
		return Color{R: 0, G: 196, B: 255}, nil

	case "Black":
		return Color{R: 0, G: 0, B: 0}, nil

	}
	return Color{}, fmt.Errorf("GetRGBColorByName: color not found: %s", color)
}

func GetLaunchPadColorCodeByRGB(color Color) (code byte) {
	switch color {
	case Color{R: 0, G: 196, B: 255}:
		return 0x25 // Light Blue
	case Color{R: 255, G: 0, B: 0}:
		return 0x48 // Red
	case Color{R: 255, G: 111, B: 0}:
		return 0x60 // Orange
	case Color{R: 255, G: 255, B: 0}:
		return 0x0d // Yellow
	case Color{R: 0, G: 255, B: 0}:
		return 0x15 // Green
	case Color{R: 0, G: 255, B: 255}:
		return 0x25 // Cyan
	case Color{R: 0, G: 0, B: 255}:
		return 0x4f // Blue
	case Color{R: 100, G: 0, B: 255}:
		return 0x51 // Purple
	case Color{R: 255, G: 0, B: 255}:
		return 0x34 // Pink
	case Color{R: 255, G: 255, B: 255}:
		return 0x03 // White
	case Color{R: 0, G: 0, B: 0}:
		return 0x00 // Black
	}

	return code
}

func SetFunctionKeyActions(functions []Function, sequence Sequence) Sequence {

	// Map the auto color change setting.
	sequence.AutoColor = functions[Function2_Auto_Color].State

	// Map the auto pattern change setting.
	sequence.AutoPattern = functions[Function3_Auto_Pattern].State

	// Map bounce function to sequence bounce setting.
	sequence.Bounce = functions[Function4_Bounce].State

	// Map color selection function.
	if functions[Function5_Color].State {
		sequence.PlayStaticOnce = true
	}

	// Map static function.
	if sequence.Type != "scanner" {
		sequence.Static = functions[Function6_Static_Gobo].State
		if functions[Function6_Static_Gobo].State {
			sequence.PlayStaticOnce = true
			sequence.Hide = true
		}
	}

	// Map RGB invert function.
	if sequence.Type == "rgb" {
		sequence.RGBInvert = functions[Function7_Invert_Chase].State
	}

	// Map scanner chase mode. Uses same function key as above.
	if sequence.Type == "scanner" {
		sequence.ScannerChase = functions[Function7_Invert_Chase].State
	}

	// Map music trigger function.
	sequence.MusicTrigger = functions[Function8_Music_Trigger].State
	if functions[Function8_Music_Trigger].State {
		sequence.Run = true
	}

	sequence.Functions = functions

	return sequence
}

func HowManyColors(positionsMap map[int]Position) (colors []Color) {

	colorMap := make(map[Color]bool)
	for _, position := range positionsMap {
		for _, fixture := range position.Fixtures {
			for _, color := range fixture.Colors {
				if color.R > 0 || color.G > 0 || color.B > 0 {
					colorMap[color] = true
				}
			}
		}
	}

	for color := range colorMap {
		colors = append(colors, color)
	}

	return colors
}

func HowManyStepColors(steps []Step) (colors []Color) {

	colorMap := make(map[Color]bool)
	for _, step := range steps {
		for _, fixture := range step.Fixtures {
			for _, color := range fixture.Colors {
				if color.R > 0 || color.G > 0 || color.B > 0 {
					colorMap[color] = true
				}
			}
		}
	}

	for color := range colorMap {
		colors = append(colors, color)
	}

	if debug {
		fmt.Printf("HowManyStepColors %d\n", len(colors))
	}

	return colors
}

func HowManyScannerColors(positionsMap map[int]Position) (colors []Color) {

	colorMap := make(map[Color]bool)
	for _, positionMap := range positionsMap {
		fixtureLen := len(positionMap.Fixtures)
		for fixtureNumber := 0; fixtureNumber < fixtureLen; fixtureNumber++ {
			fixture := positionMap.Fixtures[fixtureNumber]
			for _, color := range fixture.Colors {
				colorMap[color] = true
			}
		}
	}

	for color := range colorMap {
		colors = append(colors, color)
	}

	return colors
}

// Get an upto date copy of the sequence.
func RefreshSequence(selectedSequence int, commandChannels []chan Command, updateChannels []chan Sequence) *Sequence {

	cmd := Command{
		Action: GetUpdatedSequence,
	}

	SendCommandToSequence(selectedSequence, cmd, commandChannels)
	newSequence := <-updateChannels[selectedSequence]
	return &newSequence
}

// For the given sequence hide the available sequence colors..
func HideColorSelectionButtons(mySequenceNumber int, sequence Sequence, eventsForLauchpad chan ALight, guiButtons chan ALight) {
	for myFixtureNumber := range sequence.RGBAvailableColors {
		LightLamp(ALight{X: myFixtureNumber, Y: mySequenceNumber, Red: 0, Green: 0, Blue: 0, Brightness: sequence.Master}, eventsForLauchpad, guiButtons)
	}
}

func ClearSelectedRowOfButtons(selectedSequence int, eventsForLauchpad chan ALight, guiButtons chan ALight) {
	for x := 0; x < 8; x++ {
		LightLamp(ALight{X: x, Y: selectedSequence, Brightness: 0, Red: 0, Green: 0, Blue: 0}, eventsForLauchpad, guiButtons)
		LabelButton(x, selectedSequence, "", guiButtons)
	}
}

func ClearLabelsSelectedRowOfButtons(selectedSequence int, guiButtons chan ALight) {
	for x := 0; x < 8; x++ {
		LabelButton(x, selectedSequence, "", guiButtons)
	}
}

func ShowFunctionButtons(sequence Sequence, selectedSequence int, eventsForLauchpad chan ALight, guiButtons chan ALight) {

	// Loop through the available functions for this sequence
	for index, function := range sequence.Functions {
		if debug {
			fmt.Printf("ShowFunctionButtons: function %+v\n", function)
		}

		if function.State {
			LightLamp(ALight{X: index, Y: selectedSequence, Brightness: 255, Red: 200, Green: 0, Blue: 255}, eventsForLauchpad, guiButtons)
		} else {
			LightLamp(ALight{X: index, Y: selectedSequence, Brightness: 255, Red: 3, Green: 255, Blue: 255}, eventsForLauchpad, guiButtons)
		}
		LabelButton(index, selectedSequence, function.Label, guiButtons)
	}
}

func ShowTopButtons(tYpe string, eventsForLauchpad chan ALight, guiButtons chan ALight) {

	type topButton struct {
		Label string
		Color Color
	}
	// Storage for the rgb labels on the top row.
	var guiTopRGBButtons [8]topButton
	guiTopRGBButtons[0] = topButton{Label: "CLEAR", Color: Pink}
	guiTopRGBButtons[1] = topButton{Label: "RED", Color: Red}
	guiTopRGBButtons[2] = topButton{Label: "GREEN", Color: Green}
	guiTopRGBButtons[3] = topButton{Label: "BLUE", Color: Blue}
	guiTopRGBButtons[4] = topButton{Label: "SENS -", Color: Cyan}
	guiTopRGBButtons[5] = topButton{Label: "SENS +", Color: Cyan}
	guiTopRGBButtons[6] = topButton{Label: "MAST -", Color: Cyan}
	guiTopRGBButtons[7] = topButton{Label: "MAST +", Color: Cyan}

	// Storage for the scanner labels on the Top row.
	var guiTopScannerButtons [8]topButton
	guiTopScannerButtons[0] = topButton{Label: "CLEAR. ^ ", Color: White}
	guiTopScannerButtons[1] = topButton{Label: " V", Color: White}
	guiTopScannerButtons[2] = topButton{Label: " < ", Color: White}
	guiTopScannerButtons[3] = topButton{Label: " > ", Color: White}
	guiTopScannerButtons[4] = topButton{Label: "SENS -", Color: Cyan}
	guiTopScannerButtons[5] = topButton{Label: "SENS +", Color: Cyan}
	guiTopScannerButtons[6] = topButton{Label: "MAST -", Color: Cyan}
	guiTopScannerButtons[7] = topButton{Label: "MAST +", Color: Cyan}

	//  The Top row of the Novation Launchpad.
	TopRow := -1

	if tYpe == "rgb" {
		// Loop through the available functions for this sequence
		for index, button := range guiTopRGBButtons {
			if debug {
				fmt.Printf("button %+v\n", button)
			}
			LightLamp(ALight{X: index, Y: TopRow, Brightness: 255, Red: button.Color.R, Green: button.Color.G, Blue: button.Color.B}, eventsForLauchpad, guiButtons)
			LabelButton(index, TopRow, button.Label, guiButtons)
		}
	}
	if tYpe == "scanner" {
		// Loop through the available functions for this sequence
		for index, button := range guiTopScannerButtons {
			if debug {
				fmt.Printf("button %+v\n", button)
			}
			LightLamp(ALight{X: index, Y: TopRow, Brightness: 255, Red: button.Color.R, Green: button.Color.G, Blue: button.Color.B}, eventsForLauchpad, guiButtons)
			LabelButton(index, TopRow, button.Label, guiButtons)
		}
	}
}

func ShowBottomButtons(tYpe string, eventsForLauchpad chan ALight, guiButtons chan ALight) {

	type bottonButton struct {
		Label string
		Color Color
	}

	// Storage for the rgb labels on the bottom row.
	var guiBottomRGBButtons [8]bottonButton
	guiBottomRGBButtons[0] = bottonButton{Label: "Speed\nDown", Color: Cyan}
	guiBottomRGBButtons[1] = bottonButton{Label: "Speed\nUp", Color: Cyan}
	guiBottomRGBButtons[2] = bottonButton{Label: "Shift\nDown", Color: Cyan}
	guiBottomRGBButtons[3] = bottonButton{Label: "Shift\nUp", Color: Cyan}
	guiBottomRGBButtons[4] = bottonButton{Label: "Size\nDown", Color: Cyan}
	guiBottomRGBButtons[5] = bottonButton{Label: "Size\nUp", Color: Cyan}
	guiBottomRGBButtons[6] = bottonButton{Label: "Fade\nSoft", Color: Cyan}
	guiBottomRGBButtons[7] = bottonButton{Label: "Fade\nSharp", Color: Cyan}

	// Storage for the scanner labels on the bottom row.
	var guiBottomScannerButtons [8]bottonButton
	guiBottomScannerButtons[0] = bottonButton{Label: "Speed\nDown", Color: Cyan}
	guiBottomScannerButtons[1] = bottonButton{Label: "Speed\nUp", Color: Cyan}
	guiBottomScannerButtons[2] = bottonButton{Label: "Shift\nDown", Color: Cyan}
	guiBottomScannerButtons[3] = bottonButton{Label: "Shift\nUp", Color: Cyan}
	guiBottomScannerButtons[4] = bottonButton{Label: "Size\nDown", Color: Cyan}
	guiBottomScannerButtons[5] = bottonButton{Label: "Size\nUp", Color: Cyan}
	guiBottomScannerButtons[6] = bottonButton{Label: "Coord\nDown", Color: White}
	guiBottomScannerButtons[7] = bottonButton{Label: "Coord\nUp", Color: White}

	//  The bottom row of the Novation Launchpad.
	bottomRow := 7

	if tYpe == "rgb" {
		// Loop through the available functions for this sequence
		for index, button := range guiBottomRGBButtons {
			if debug {
				fmt.Printf("button %+v\n", button)
			}
			LightLamp(ALight{X: index, Y: bottomRow, Brightness: 255, Red: button.Color.R, Green: button.Color.G, Blue: button.Color.B}, eventsForLauchpad, guiButtons)
			LabelButton(index, bottomRow, button.Label, guiButtons)
		}
	}
	if tYpe == "scanner" {
		// Loop through the available functions for this sequence
		for index, button := range guiBottomScannerButtons {
			if debug {
				fmt.Printf("button %+v\n", button)
			}
			LightLamp(ALight{X: index, Y: bottomRow, Brightness: 255, Red: button.Color.R, Green: button.Color.G, Blue: button.Color.B}, eventsForLauchpad, guiButtons)
			LabelButton(index, bottomRow, button.Label, guiButtons)
		}
	}
}

func ShowRunningStatus(sequenceNumber int, runningState map[int]bool, eventsForLaunchpad chan ALight, guiButtons chan ALight) {
	for key, value := range runningState {
		if sequenceNumber == key {
			if value {
				LightLamp(ALight{X: 8, Y: 5, Brightness: 255, Red: 0, Green: 255, Blue: 0}, eventsForLaunchpad, guiButtons)
			} else {
				LightLamp(ALight{X: 8, Y: 5, Brightness: 255, Red: 255, Green: 255, Blue: 255}, eventsForLaunchpad, guiButtons)
			}
		}
	}
}

func ShowStrobeButtonStatus(state bool, eventsForLaunchpad chan ALight, guiButtons chan ALight) {
	if state {
		FlashLight(8, 6, White, Black, eventsForLaunchpad, guiButtons)
		return
	}
	LightLamp(ALight{X: 8, Y: 6, Brightness: 255, Red: 255, Green: 255, Blue: 255}, eventsForLaunchpad, guiButtons)
}

func LabelButton(X int, Y int, label string, guiButtons chan ALight) {
	if debug {
		fmt.Printf("Label Button  X:%d  Y:%d  with %s\n", X, Y, label)
	}
	// Send message to GUI
	event := ALight{
		UpdateLabel: true,
		X:           X,
		Y:           Y + 1,
		Label:       label,
	}
	guiButtons <- event
}

// LightOn Turn on a Light.
func LightLamp(Light ALight, eventsForLauchpad chan ALight, guiButtons chan ALight) {
	if debug {
		fmt.Printf("LightLamp  X:%d  Y:%d Red %d Green %d Blue %d Brightnes %d\n", Light.X, Light.Y, Light.Red, Light.Green, Light.Blue, Light.Brightness)
	}
	// Send message to Novation Launchpad.
	event := ALight{
		X:          Light.X,
		Y:          Light.Y,
		Brightness: Light.Brightness,
		Red:        Light.Red,
		Green:      Light.Green,
		Blue:       Light.Blue,
		Flash:      Light.Flash,
		OnColor:    Light.OnColor,
		OffColor:   Light.OffColor,
	}
	eventsForLauchpad <- event

	// Send message to fyne.io GUI.
	event = ALight{
		X:          Light.X,
		Y:          Light.Y + 1,
		Brightness: Light.Brightness,
		Red:        Light.Red,
		Green:      Light.Green,
		Blue:       Light.Blue,
		Flash:      Light.Flash,
		OnColor:    Light.OnColor,
		OffColor:   Light.OffColor,
		Label:      Light.Label,
	}
	guiButtons <- event
}

func UpdateStatusBar(status string, which string, hide bool, guiButtons chan ALight) {
	// Send message to fyne.io GUI.
	event := ALight{
		UpdateStatus: true,
		Status:       status,
		Which:        which,
		Hidden:       hide,
	}
	guiButtons <- event
}

func FlashLight(X int, Y int, onColor Color, offColor Color, eventsForLauchpad chan ALight, guiButtons chan ALight) {

	// Now ask the fixture lamp to flash on the launch pad by sending an event.
	e := ALight{
		X:          X,
		Y:          Y,
		Brightness: 255,
		Flash:      true,
		OnColor:    onColor,
		OffColor:   offColor,
	}
	eventsForLauchpad <- e

	// Send message to GUI
	event := ALight{
		X:          e.X,
		Y:          e.Y + 1,
		Brightness: 255,
		Flash:      true,
		OnColor:    onColor,
		OffColor:   offColor,
	}
	guiButtons <- event
}

// InvertColor just reverses the DMX values.
func InvertColor(color Color) (out Color) {

	out.R = ReverseDmx(color.R)
	out.G = ReverseDmx(color.G)
	out.B = ReverseDmx(color.B)

	return out
}

// Takes a DMX value 1-255 and reverses the value.
func ReverseDmx(n int) int {
	in := make(map[int]int, 255)
	var y = 255

	for x := 0; x <= 255; x++ {

		in[x] = y
		y--
	}
	return in[n]
}

// Sets the static colors to default values.
func SetDefaultStaticColorButtons(selectedSequence int) []StaticColorButton {

	// Make an array to hold static colors.
	staticColorsButtons := []StaticColorButton{}

	for X := 0; X < 8; X++ {
		staticColorButton := StaticColorButton{}
		staticColorButton.X = X
		staticColorButton.Y = selectedSequence
		staticColorButton.SelectedColor = X
		staticColorButton.Color = GetColorButtonsArray(X)
		staticColorsButtons = append(staticColorsButtons, staticColorButton)
	}

	return staticColorsButtons
}

func Reverse(in int) int {
	switch in {
	case 0:
		return 10
	case 1:
		return 9
	case 2:
		return 8
	case 3:
		return 7
	case 4:
		return 6
	case 5:
		return 5
	case 6:
		return 4
	case 7:
		return 3
	case 8:
		return 2
	case 9:
		return 1
	case 10:
		return 0
	}

	return 10
}

// CalculateFadeValues - calculate fade curve values.
func CalculateFadeValues(noCoordianates int, fade int, size int) (slopeOn []int, slopeOff []int) {

	fadeUpValues := GetFadeValues(noCoordianates, MaxDMXBrightness, fade, false)
	fadeOnValues := GetFadeOnValues(MaxDMXBrightness, size)
	fadeDownValues := GetFadeValues(noCoordianates, MaxDMXBrightness, fade, true)

	slopeOn = append(slopeOn, fadeUpValues...)
	slopeOn = append(slopeOn, fadeOnValues...)
	slopeOn = append(slopeOn, fadeDownValues...)

	slopeOff = append(slopeOff, fadeDownValues...)
	slopeOff = append(slopeOff, fadeUpValues...)

	return slopeOn, slopeOff
}

func GetFadeValues(noCoordinates int, size float64, fade int, reverse bool) (out []int) {

	var x float64
	var counter float64
	var slope float64

	coordinates := float64(noCoordinates)

	switch fade {
	case 10:
		slope = 30
	case 9:
		slope = 15
	case 8:
		slope = 10
	case 7:
		slope = 7
	case 6:
		slope = 5
	case 5:
		slope = 4
	case 4:
		slope = 3
	case 3:
		slope = 2
	case 2:
		slope = 1.5
	case 1:
		slope = 1
	default:
		slope = 0
	}

	if !reverse {
		for counter = 0; counter <= coordinates-1; counter++ {
			x = (counter / 2) / (coordinates - 1)
			y := math.Pow(math.Sin(x*math.Pi), slope)
			dmx := int(size * y)
			out = append(out, dmx)
		}
	} else {
		for counter = coordinates - 1; counter >= 0; counter-- {
			x = (counter / 2) / (coordinates - 1)
			y := math.Pow(math.Sin(x*math.Pi), slope)
			dmx := int(size * y)
			out = append(out, dmx)
		}
	}

	return out
}

func GetFadeOnValues(size int, fade int) []int {

	out := []int{}

	var x int

	for x = 0; x < fade; x++ {
		x := size
		out = append(out, x)
	}

	return out
}

func FindSensitivity(soundGain float32) int {

	in := fmt.Sprintf("%f", soundGain)

	switch in {
	case "-0.040000":
		return 0
	case "-0.030000":
		return 1
	case "-0.020000":
		return 2
	case "-0.010000":
		return 3
	case "0.000000":
		return 4
	case "0.010000":
		return 5
	case "0.020000":
		return 6
	case "0.030000":
		return 7
	case "0.040000":
		return 8
	case "0.050000":
		return 9
	case "0.060000":
		return 10
	case "0.070000":
		return 11
	case "0.080000":
		return 12
	case "0.090000":
		return 13
	}

	return 99
}

func FormatLabel(label string) string {
	// replace any spaces with new lines.
	// new lines are represented by a dot in code beneath us.
	return strings.Replace(label, " ", ".", -1)
}
