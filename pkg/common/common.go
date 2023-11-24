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

const MAX_NUMBER_OF_CHANNELS = 8
const MAX_DMX_ADDRESS = 512
const MAX_TEXT_ENTRY_LENGTH = 35
const DEFAULT_SCANNER_SIZE = 60
const MAX_SCANNER_SIZE = 127
const MIN_SPEED = 0
const MAX_SPEED = 12
const MIN_RGB_SIZE = 0
const MAX_RGB_SIZE = 10
const MIN_RGB_SHIFT = 1
const MAX_RGB_SHIFT = 10
const MAX_SCANNER_SHIFT = 3
const MIN_SCANNER_SHIFT = 0
const MAX_RGB_FADE = 10
const MAX_COLOR_BAR = 9 // Eight colors and a default color bar.
const MIN_DMX_BRIGHTNESS = 0
const CENTER_DMX_BRIGHTNESS = 127
const MAX_DMX_BRIGHTNESS = 255
const DEFAULT_PATTERN = 0
const DEFAULT_RGB_SIZE = 0
const DEFAULT_RGB_FADE = 1
const DEFAULT_SCANNER_FADE = 10
const DEFAULT_SPEED = 7
const DEFAULT_RGB_SHIFT = 0
const DEFAULT_SCANNER_COLOR = 1
const DEFAULT_SCANNER_GOBO = 1
const DEFAULT_SCANNER_SHIFT = 0
const DEFAULT_SCANNER_COORDNIATES = 0
const SCANNER_MID_POINT = 127
const DEFAULT_RGB_FADE_STEPS = 10
const DEFAULT_STROBE_SPEED = 255

const IS_SCANNER = true
const IS_RGB = false

var DefaultSequenceColors = []Color{{R: 0, G: 255, B: 0}}
var GlobalScannerSequenceNumber int

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

type ColorPicker struct {
	Name  string
	ID    int
	Code  byte // Launchpad hex code for this color
	Color Color
	X     int
	Y     int
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
	DebugMsg     string
	Step         int
	Rule         int
}

type Value struct {
	Channel string
	Setting string
}

type Setting struct {
	Name         string
	Label        string
	Number       int
	Channel      string
	Value        int16
	FixtureValue string
}

type State struct {
	Name        string
	Number      int16
	Label       string
	Values      []Value
	ButtonColor string
	Actions     []Action
	Settings    []Setting
	Flash       bool
}

type Action struct {
	Name        string
	Number      int
	Colors      []string
	Mode        string
	Fade        string
	Size        string
	Speed       string
	Rotate      string
	RotateSpeed string
	Music       string
	Program     string
	Strobe      string
}

type Switch struct {
	ID                   int
	Name                 string
	Address              int16
	Number               int
	Label                string
	CurrentPosition      int
	Description          string
	States               map[int]State
	Fixture              string
	UseFixture           string
	MiniSequencerRunning bool
	Blackout             bool
	Master               int
}

type StaticColorButton struct {
	Name             string
	Label            string
	Number           int
	X                int
	Y                int
	SelectedSequence int
	Color            Color
	SelectedColor    int
	Flash            bool
	Setting          int
	FirstPress       bool
	Enabled          bool
}

type FixtureState struct {
	Enabled                bool
	RGBInverted            bool
	ScannerPatternReversed bool
}

type Pattern struct {
	Name     string
	Label    string
	Number   int
	Length   int // 8, 4 or 2
	Size     int
	Fixtures int    // 8 Fixtures
	Steps    []Step `json:"-"` // Don't save the steps as they can be very large.
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
	UpdateStatic
	UpdateFlashAllStaticColorButtons
	UpdateBounce
	UpdateAllStaticColor
	UpdateStaticColor
	UpdateASingeSequenceColor
	UpdateSequenceColors
	PlayStaticOnce
	PlaySwitchOnce
	UnHide
	Hide
	Start
	StartChase
	Stop
	StopChase
	ReadConfig
	LoadConfig
	UpdateSpeed
	UpdatePattern
	UpdateRGBFadeSpeed
	UpdateRGBSize
	UpdateScannerSize
	Blackout
	Normal
	UpdateColor
	UpdateFunctions
	GetUpdatedSequence
	ResetAllSwitchPositions
	UpdateSwitch
	Inverted
	UpdateGobo
	Flood
	StopFlood
	Strobe
	StopStrobe
	UpdateAutoColor
	UpdateAutoPattern
	ToggleFixtureState
	UpdateRGBShift
	UpdateRGBInvert
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
	UpdateScannerChase
	UpdateMusicTrigger
	UpdateScannerHasShutterChase
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
var EmptyColor = Color{}

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
	MasterChanging              bool                        // flag to indicate we are changing brightness.
	Strobe                      bool                        // Strobe is enabled.
	StrobeSpeed                 int                         // Strobe speed.
	Rotate                      int                         // Rotate speed.
	RGBShift                    int                         // RGB shift.
	CurrentSpeed                time.Duration               // Sequence speed represented as a duration.
	Speed                       int                         // Sequence speed represented by a short number.
	MusicTrigger                bool                        // Is this sequence in music trigger mode.
	ChangeMusicTrigger          bool                        // true when we change the state of the music trigger.
	LastMusicTrigger            bool                        // Save copy of music trigger.
	Blackout                    bool                        // Flag to indicate we're in blackout mode.
	CurrentColors               []Color                     // Storage for the colors in a sequence.
	SequenceColors              []Color                     // Temporay storage for changing sequence colors.
	Color                       int                         // Index into current sequnece colors.
	ScannerSteps                []Step                      // Pan & Tilt steps in this  sequence.
	NumberSteps                 int                         // Holds the number of steps this sequence has. Will change if you change size, fade times etc.
	NumberFixtures              int                         // Total Number of fixtures for this sequence.
	EnabledNumberFixtures       int                         // Enabled Number of fixtures for this sequence.
	AutoColor                   bool                        // Sequence is going to automatically change the color.
	AutoPattern                 bool                        // Sequence is going to automatically change the pattern.
	GuiFixtureLabels            []string                    // Storage for the fixture labels. Used for scanner names.
	Pattern                     Pattern                     // Contains fixtures and RGB steps info.
	RGBAvailableColors          []StaticColorButton         // Available colors for the RGB fixtures.
	RGBColor                    int                         // The selected RGB fixture color.
	FadeUp                      []int                       // Fade up values.
	FadeOn                      []int                       // Fade on values.
	FadeDown                    []int                       // Fade down values.
	FadeOff                     []int                       // Fade off values.
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
	StaticFadeOnce              bool                        // Only Fade up once, used for don't fade during color config operations.
	StartFlood                  bool                        // We're in flood mode.
	StopFlood                   bool                        // We're not in flood mode.
	LastStatic                  bool                        // Last value of static before flood.
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
	ScannerChaser               bool                        // Chase the scanner shutters instead of allways being on.
	ScannerReverse              bool                        // Reverse the scanner, i.e scan in the opposite direction.
	ScannerColor                map[int]int                 // Eight scanners per sequence, each can have their own color.
	ScannerCoordinates          []int                       // Number of scanner coordinates.
	ScannerSelectedCoordinates  int                         // index into scanner coordinates.
	ScannerOffsetPan            int                         // Offset for pan values.
	ScannerOffsetTilt           int                         // Offset for tilt values.
	FixtureState                map[int]FixtureState        // Map of fixtures which are disabled.
	DisableOnceMutex            *sync.RWMutex               // Lock to protect DisableOnce.
	DisableOnce                 map[int]bool                // Map used to play disable only once.
	UpdateSize                  bool                        // Command to update size.
	UpdateShift                 bool                        // Command to update the shift.
	UpdatePattern               bool                        // Flag to indicate we're going to change the RGB pattern.
	UpdateSequenceColor         bool                        // Command to update the sequence colors.
	Switches                    map[int]Switch              // A switch sequence stores its data in here.
	CurrentSwitch               int                         // Play this current switch position.
	Optimisation                bool                        // Flag to decide on calculatePositions Optimisation.
	RGBNumberStepsInFade        int                         // Number of steps in a RGB fade.
	Hidden                      bool                        // Is this sequence hidden on the launchpad.
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
	Stop            chan bool
	StopRotate      chan bool
	StopFadeDown    chan bool
	StopFadeUp      chan bool
	KeepRotateAlive chan bool
}

type Hit struct {
	X int
	Y int
}

type Step struct {
	StepNumber int
	KeyStep    bool
	Fixtures   map[int]Fixture
}

type FixtureCommand struct {
	Step           int
	NumberSteps    int
	Type           string
	Label          string
	SequenceNumber int
	FixtureState   FixtureState
	LastColor      Color

	// Common commands.
	Strobe         bool
	StrobeSpeed    int
	Master         int
	MasterChanging bool
	Blackout       bool
	Hide           bool
	Clear          bool
	FadeSpeed      int

	StartFlood bool
	StopFlood  bool

	// RGB commands.
	RGBPosition       Position
	RGBStatic         bool
	RGBStaticColors   []StaticColorButton
	RGBPlayStaticOnce bool

	// Scanner Commands.
	ScannerColor             int
	ScannerPosition          Position
	ScannerDisableOnce       bool
	ScannerChaser            bool
	ScannerAvailableColors   []StaticColorButton
	ScannerGobo              int
	ScannerOffsetPan         int
	ScannerOffsetTilt        int
	ScannerNumberCoordinates int
	ScannerShutterPositions  map[int]Position
	ScannerHasShutterChase   bool

	// Derby Commands
	Rotate  int
	Music   int
	Program int

	// Switch Commands
	SetSwitch          bool
	SwitchData         Switch
	State              State
	CurrentSwitchState int
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
	Number       int
	Name         string
	Label        string
	MasterDimmer int
	Brightness   int
	ScannerColor Color
	Color        Color
	Pan          int
	Tilt         int
	Shutter      int
	Rotate       int
	Music        int
	Gobo         int
	Program      int
	Enabled      bool
	Inverted     bool
	State        int // Last thing we did :- MAKE SAME AGAIN ,FADEUP or FADEDOWN
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

func SendCommandToSequence(targetSequence int, command Command, commandChannels []chan Command) {
	commandChannels[targetSequence] <- command
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

func SendCommandToAllSequenceExcept(targetSequence int, command Command, commandChannels []chan Command) {
	for index := range commandChannels {
		if index != targetSequence {
			commandChannels[index] <- command
		}
	}
}

func RevealSequence(targetSequence int, commandChannels []chan Command) {
	cmd := Command{
		Action: UnHide,
	}
	SendCommandToSequence(targetSequence, cmd, commandChannels)
}

func HideSequence(targetSequence int, commandChannels []chan Command) {
	cmd := Command{
		Action: Hide,
	}
	SendCommandToSequence(targetSequence, cmd, commandChannels)
}

func HideAllSequences(commandChannels []chan Command) {

	cmd := Command{
		Action: Hide,
	}
	SendCommandToAllSequence(cmd, commandChannels)
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

func GetColorNameByRGB(color Color) string {
	switch color {
	case Color{R: 0, G: 196, B: 255}:
		return "Light Blue"
	case Color{R: 255, G: 0, B: 0}:
		return "Red"
	case Color{R: 255, G: 111, B: 0}:
		return "Orange"
	case Color{R: 255, G: 255, B: 0}:
		return "Yellow"
	case Color{R: 0, G: 255, B: 0}:
		return "Green"
	case Color{R: 0, G: 255, B: 255}:
		return "Cyan"
	case Color{R: 0, G: 0, B: 255}:
		return "Blue"
	case Color{R: 100, G: 0, B: 255}:
		return "Purple"
	case Color{R: 255, G: 0, B: 255}:
		return "Pink"
	case Color{R: 255, G: 255, B: 255}:
		return "White"
	case Color{R: 0, G: 0, B: 0}:
		return "Black"
	}

	return "White"
}

func HowManyColorsInSteps(steps []Step) (colors []Color) {

	if debug {
		fmt.Printf("HowManyColorsInSteps \n")
	}
	colorMap := make(map[Color]bool)
	for _, step := range steps {
		for _, fixture := range step.Fixtures {
			if fixture.Color.R > 0 || fixture.Color.G > 0 || fixture.Color.B > 0 {
				colorMap[fixture.Color] = true
			}
		}
	}

	for color := range colorMap {
		if debug {
			fmt.Printf("add color %+v\n", color)
		}
		colors = append(colors, color)
	}

	return colors
}

func HowManyColorsInPositions(positionsMap map[int]Position) (colors []Color) {

	colorMap := make(map[Color]bool)
	for _, position := range positionsMap {
		for _, fixture := range position.Fixtures {
			if fixture.Color.R > 0 || fixture.Color.G > 0 || fixture.Color.B > 0 {
				colorMap[fixture.Color] = true
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
			if fixture.Color.R > 0 || fixture.Color.G > 0 || fixture.Color.B > 0 {
				colorMap[fixture.Color] = true
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
			colorMap[fixture.Color] = true
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

func ClearSelectedRowOfButtons(selectedSequence int, eventsForLauchpad chan ALight, guiButtons chan ALight) {
	if selectedSequence == 4 {
		return
	}
	for x := 0; x < 8; x++ {
		LightLamp(ALight{X: x, Y: selectedSequence, Brightness: 0, Red: 0, Green: 0, Blue: 0}, eventsForLauchpad, guiButtons)
		LabelButton(x, selectedSequence, "", guiButtons)
	}
}

func ClearAllButtons(eventsForLauchpad chan ALight, guiButtons chan ALight) {
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			LightLamp(ALight{X: x, Y: y, Brightness: 0, Red: 0, Green: 0, Blue: 0}, eventsForLauchpad, guiButtons)
			LabelButton(x, y, "", guiButtons)
		}
	}
}

func ClearLabelsSelectedRowOfButtons(selectedSequence int, guiButtons chan ALight) {
	if selectedSequence == 4 {
		return
	}
	for x := 0; x < 8; x++ {
		LabelButton(x, selectedSequence, "", guiButtons)
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
	guiTopScannerButtons[0] = topButton{Label: "CLEAR.^", Color: White}
	guiTopScannerButtons[1] = topButton{Label: "V", Color: White}
	guiTopScannerButtons[2] = topButton{Label: "<", Color: White}
	guiTopScannerButtons[3] = topButton{Label: ">", Color: White}
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

	if debug {
		fmt.Printf("ShowBottomButtons\n")
	}

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
	guiBottomScannerButtons[6] = bottonButton{Label: "Coord\nDown", Color: Cyan}
	guiBottomScannerButtons[7] = bottonButton{Label: "Coord\nUp", Color: Cyan}

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

func ShowRunningStatus(runningState bool, eventsForLaunchpad chan ALight, guiButtons chan ALight) {
	if runningState {
		LightLamp(ALight{X: 8, Y: 5, Brightness: 255, Red: 0, Green: 255, Blue: 0}, eventsForLaunchpad, guiButtons)
	} else {
		LightLamp(ALight{X: 8, Y: 5, Brightness: 255, Red: 255, Green: 255, Blue: 255}, eventsForLaunchpad, guiButtons)
	}
}

func ShowStrobeButtonStatus(state bool, eventsForLaunchpad chan ALight, guiButtons chan ALight) {
	if state {
		FlashLight(8, 6, White, Pink, eventsForLaunchpad, guiButtons)
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
	eventsForLauchpad <- event // Event will be received by pkg/launchpad/launchpad.go ListenAndSendToLaunchPad()

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
	guiButtons <- event // Event will be received by dmxlights.go by pkg/gui/gui.go ListenAndSendToGUI()
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

func UpdateBottomButtons(selectedType string, guiButtons chan ALight) {

	LabelButton(0, 7, "Speed\nDown", guiButtons)
	LabelButton(1, 7, "Speed\nUp", guiButtons)

	LabelButton(2, 7, "Shift\nDown", guiButtons)
	LabelButton(3, 7, "Shift\nUp", guiButtons)

	LabelButton(4, 7, "Size\nDown", guiButtons)
	LabelButton(5, 7, "Size\nUp", guiButtons)

	if selectedType == "rgb" {
		LabelButton(6, 7, "Fade\nSoft", guiButtons)
		LabelButton(7, 7, "Fade\nSharp", guiButtons)
	}

	if selectedType == "scanner" {
		LabelButton(6, 7, "Coord\nDown", guiButtons)
		LabelButton(7, 7, "Coord\nUp", guiButtons)
	}
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
	eventsForLauchpad <- e // Event will be received by pkg/launchpad/launchpad.go ListenAndSendToLaunchPad()

	// Send message to GUI
	event := ALight{
		X:          e.X,
		Y:          e.Y + 1,
		Brightness: 255,
		Flash:      true,
		OnColor:    onColor,
		OffColor:   offColor,
	}
	guiButtons <- event // Event will be received by dmxlights.go by pkg/gui/gui.go ListenAndSendToGUI()

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

	var selectedColor int

	for Y := 0; Y < 3; Y++ {
		for X := 0; X < 8; X++ {
			if selectedColor >= 24 {
				break
			}

			staticColorButton := StaticColorButton{}

			colorPicker := GetColor(X, Y)
			staticColorButton.Name = colorPicker.Name
			staticColorButton.Color = colorPicker.Color
			staticColorButton.X = X
			staticColorButton.Y = Y
			staticColorButton.SelectedColor = selectedColor
			selectedColor++

			staticColorsButtons = append(staticColorsButtons, staticColorButton)
		}

	}

	return staticColorsButtons
}

func Reverse(in int) int {
	switch in {
	case 0:
		return 20
	case 1:
		return 15
	case 2:
		return 10
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
	default:
		return 20
	}
}

func Reverse12(in int) int {
	switch in {
	case 0:
		return 12
	case 1:
		return 11
	case 2:
		return 10
	case 3:
		return 9
	case 4:
		return 8
	case 5:
		return 7
	case 6:
		return 6
	case 7:
		return 5
	case 8:
		return 4
	case 9:
		return 3
	case 10:
		return 2
	case 11:
		return 1
	case 12:
		return 0
	default:
		return 12
	}
}

// CalculateFadeValues - calculate fade curve values.
func CalculateFadeValues(sequence *Sequence) {
	sequence.FadeUp = GetFadeValues(sequence.RGBNumberStepsInFade, MAX_DMX_BRIGHTNESS, sequence.RGBFade, false)
	sequence.FadeOn = GetFadeOnValues(MAX_DMX_BRIGHTNESS, sequence.RGBSize)
	sequence.FadeDown = GetFadeValues(sequence.RGBNumberStepsInFade, MAX_DMX_BRIGHTNESS, sequence.RGBFade, true)
	sequence.FadeOff = GetFadeOnValues(MIN_DMX_BRIGHTNESS, sequence.RGBSize)
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

func newColorPicker() []ColorPicker {

	colors := []ColorPicker{

		{ID: 0, X: 0, Y: 0, Name: "Red", Code: 0x48, Color: Color{R: 255, G: 0, B: 0}},
		{ID: 1, X: 1, Y: 0, Name: "Orange", Code: 0x09, Color: Color{R: 255, G: 100, B: 0}},
		{ID: 2, X: 2, Y: 0, Name: "Yellow", Code: 0x0d, Color: Color{R: 255, G: 255, B: 0}},
		{ID: 3, X: 3, Y: 0, Name: "Green", Code: 0x4C, Color: Color{R: 0, G: 255, B: 0}},
		{ID: 4, X: 4, Y: 0, Name: "Cyan", Code: 0x25, Color: Color{R: 0, G: 255, B: 255}},
		{ID: 5, X: 5, Y: 0, Name: "Blue", Code: 0x4f, Color: Color{R: 0, G: 0, B: 255}},
		{ID: 14, X: 6, Y: 1, Name: "Purple", Code: 0x32, Color: Color{R: 50, G: 0, B: 100}},
		{ID: 7, X: 7, Y: 0, Name: "Magenta", Code: 0x35, Color: Color{R: 255, G: 0, B: 255}},

		{ID: 8, X: 0, Y: 1, Name: "Crimson", Code: 0x38, Color: Color{R: 220, G: 20, B: 60}},
		{ID: 9, X: 1, Y: 1, Name: "Dark Orange", Code: 0x0a, Color: Color{R: 215, G: 50, B: 0}},
		{ID: 10, X: 2, Y: 1, Name: "Gold", Code: 0x61, Color: Color{R: 255, G: 215, B: 0}},
		{ID: 11, X: 3, Y: 1, Name: "Forest Green", Code: 0x1b, Color: Color{R: 0, G: 100, B: 0}},
		{ID: 12, X: 4, Y: 1, Name: "Aqua", Code: 0x20, Color: Color{R: 127, G: 255, B: 212}},
		{ID: 13, X: 5, Y: 1, Name: "Sky Blue", Code: 0x25, Color: Color{R: 0, G: 191, B: 255}},
		{ID: 6, X: 6, Y: 0, Name: "Purple", Code: 0x51, Color: Color{R: 100, G: 0, B: 255}},
		{ID: 15, X: 7, Y: 1, Name: "Pink", Code: 0x34, Color: Color{R: 255, G: 105, B: 180}},

		{ID: 16, X: 0, Y: 2, Name: "Salmon", Code: 0x6b, Color: Color{R: 250, G: 128, B: 114}},
		{ID: 17, X: 1, Y: 2, Name: "Light Orange", Code: 0x0c, Color: Color{R: 255, G: 175, B: 0}},
		{ID: 18, X: 2, Y: 2, Name: "Olive", Code: 0x10, Color: Color{R: 150, G: 150, B: 0}},
		{ID: 19, X: 3, Y: 2, Name: "Lawn green", Code: 0x13, Color: Color{R: 124, G: 252, B: 0}},
		{ID: 20, X: 4, Y: 2, Name: "Teal", Code: 0x44, Color: Color{R: 0, G: 128, B: 128}},
		{ID: 21, X: 5, Y: 2, Name: "Light Blue", Code: 0x20, Color: Color{R: 100, G: 185, B: 255}},
		{ID: 22, X: 6, Y: 2, Name: "Violet", Code: 0x5e, Color: Color{R: 199, G: 21, B: 133}},
		{ID: 23, X: 7, Y: 2, Name: "White", Code: 0x03, Color: Color{R: 255, G: 255, B: 255}},
	}

	return colors
}

func GetLaunchPadCodeByRGBColor(selectedColor Color) byte {

	colors := newColorPicker()
	if debug {
		fmt.Printf("Selected Color %+v\n", selectedColor)
	}
	for _, color := range colors {

		if selectedColor == color.Color {
			if debug {
				fmt.Printf("Color Name %s Code %x\n", color.Name, color.Code)
			}
			return color.Code
		}
	}
	return 0

}

func GetIDfromCoordinates(X int, Y int) int {

	colors := newColorPicker()

	for _, color := range colors {

		if color.X == X && color.Y == Y {
			return color.ID
		}
	}
	return 0
}

func GetColor(X int, Y int) ColorPicker {

	colors := newColorPicker()

	for _, color := range colors {

		if color.X == X && color.Y == Y {
			return color
		}
	}
	return ColorPicker{}
}
