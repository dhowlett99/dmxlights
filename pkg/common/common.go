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

	"github.com/dhowlett99/dmxlights/pkg/colors"
)

const debug = false
const NOT_SELECTED = -1

const MAX_NUMBER_OF_CHANNELS = 8
const MAX_DMX_ADDRESS = 512
const MAX_TEXT_ENTRY_LENGTH = 35
const DEFAULT_SCANNER_SIZE = 60
const MAX_SCANNER_SIZE = 127
const MIN_SPEED = 0
const MAX_SPEED = 12
const MIN_RGB_SHIFT = 1
const MAX_RGB_SHIFT = 10
const MIN_RGB_SIZE = 0
const MAX_RGB_SIZE = 10
const MIN_RGB_FADE = 1
const MAX_RGB_FADE = 10
const MAX_SCANNER_SHIFT = 3
const MIN_SCANNER_SHIFT = 0
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
const DEFAULT_SCANNER_COLOR = 0
const DEFAULT_SCANNER_GOBO = 0
const DEFAULT_SCANNER_SHIFT = 0
const DEFAULT_SCANNER_COORDNIATES = 0
const SCANNER_MID_POINT = 127
const DEFAULT_RGB_FADE_STEPS = 10
const DEFAULT_STROBE_SPEED = 255

const IS_SCANNER = true
const IS_RGB = false

var GlobalScannerSequenceNumber int

var FLOOD_BUTTON = Button{X: 8, Y: 3}
var SAVE_BUTTON = Button{X: 8, Y: 4}
var RUNNING_BUTTON = Button{X: 8, Y: 5}
var STROBE_BUTTON = Button{X: 8, Y: 6}
var BLACKOUT_BUTTON = Button{X: 8, Y: 7}

var RED_BUTTON = Button{X: 1, Y: -1}
var GREEN_BUTTON = Button{X: 2, Y: -1}
var BLUE_BUTTON = Button{X: 3, Y: -1}

type Button struct {
	X int
	Y int
}

type ColorDisplayControl struct {
	Red     bool
	Orange  bool
	Yellow  bool
	Green   bool
	Cyan    bool
	Blue    bool
	Purple  bool
	Magenta bool

	Crimson     bool
	DarkOrange  bool
	Gold        bool
	ForestGreen bool
	Aqua        bool
	SkyBlue     bool
	DarkPurple  bool
	Pink        bool

	Salmon      bool
	LightOrange bool
	Olive       bool
	LawnGreen   bool
	Teal        bool
	LightBlue   bool
	Violet      bool
	White       bool
}

type ALight struct {
	Button              Button
	Brightness          int
	Red                 uint8
	Green               uint8
	Blue                uint8
	Flash               bool
	OnColor             color.RGBA
	OffColor            color.RGBA
	UpdateLabel         bool
	Label               string
	UpdateStatus        bool
	Status              string
	Which               string
	FlashStopChannel    chan bool
	Hidden              bool
	ColorDisplay        bool
	ColorDisplayControl ColorDisplayControl
}

// Used for static fades, remember the last color.
type LastColor struct {
	RGBColor     color.RGBA
	ScannerColor int
}

type ColorPicker struct {
	Name  string
	ID    int
	Code  byte // Launchpad hex code for this color
	Color color.RGBA
	X     int
	Y     int
}

// Used in calculating Positions.
type FixtureBuffer struct {
	BaseColor    color.RGBA
	Color        color.RGBA
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
	Name         string
	Number       int
	Colors       []string
	Mode         string
	Fade         string
	Size         string
	Speed        string
	Rotate       string
	RotateSpeed  string
	Music        string
	Program      string
	ProgramSpeed string
	Strobe       string
	Map          string
	Gobo         string
	GoboSpeed    string
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
	Selected             bool
	Override             Override
}

type Override struct {
	Override    bool
	Speed       int
	Shift       int
	Size        int
	Fade        int
	RotateSpeed int
	Colors      []color.RGBA
	Gobo        int
	GoboName    string
}

type StaticColorButton struct {
	Name             string
	Label            string
	Number           int
	X                int
	Y                int
	SelectedSequence int
	Color            color.RGBA
	SelectedColor    int
	Flash            bool
	Setting          int
	FirstPress       bool
	Enabled          bool
	NumberOfGobos    int
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
	Reveal
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
	UpdateRotateSpeed
	UpdateColors
	UpdateGobos
	UpdateScannerSize
	Blackout
	Normal
	UpdateFunctions
	GetUpdatedSequence
	ResetAllSwitchPositions
	UpdateSwitch
	OverrideSwitchSpeed
	OverrideSwitchShift
	OverrideSwitchSize
	OverrideSwitchFade
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
	UpdateFixturesConfig
)

// A full step cycle is 39 ticks ie 39 values.
// 13 fade up values, 13 on values and 13 off values.
const StepSize = 39

type Gobo struct {
	Name    string
	Label   string
	Number  int
	Setting int
	Flash   bool
	Color   color.RGBA
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
	Hidden                      bool                        // Hidden is used to indicate sequence buttons are not visible.
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
	SequenceColors              []color.RGBA                // Temporay storage for changing sequence colors.
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
	SavedSequenceColors         []color.RGBA                // Used for updating the color in a sequence.
	RecoverSequenceColors       bool                        // Storage for recovering sequence colors, when you come out of automatic color change.
	SaveColors                  bool                        // Indicate we should save colors in this sequence. used for above.
	Chase                       bool                        // Tells sequnece if we're in sequence (chase) or static (static colors) mode.
	StaticColors                []StaticColorButton         // Used in static color editing
	Clear                       bool                        // Clear all fixtures in this sequence.
	Static                      bool                        // We're a static sequence.
	PlayStaticOnce              bool                        // Play a static scene only once.
	PlayStaticLampsOnce         bool                        // Play a static scene but only on indicator lamps.
	PlaySwitchOnce              bool                        // Play a switch sequence scene only once.
	OverrideSpeed               bool                        // Override a switch speed.
	OverrideShift               bool                        // Override a switch shift.
	OverrideSize                bool                        // Override a switch size.
	OverrideFade                bool                        // Override a switch fade.
	PlaySingleSwitch            bool                        // Play a single switch.
	StepSwitch                  bool                        // Step the switch if true.
	FocusSwitch                 bool                        // Focus the switch.
	StaticFadeUpOnce            bool                        // Only Fade up once, used for don't fade during color config operations.
	StaticLampsOn               bool                        // Show the static scene on the lamps, but don't send anything to the DMX universe.
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
	UpdateColors                bool                        // Command to update the sequence colors.
	Switches                    map[int]Switch              // A switch sequence stores its data in here.
	CurrentSwitch               int                         // Play this current switch position.
	Optimisation                bool                        // Flag to decide on calculatePositions Optimisation.
	RGBNumberStepsInFade        int                         // Number of steps in a RGB fade.
	LastSwitchSelected          int                         // Storage for the last selected switch.
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
	CommandChannel  chan Command
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
	LastColor      color.RGBA

	// Common commands.
	Hidden         bool
	Strobe         bool
	StrobeSpeed    int
	Master         int
	MasterChanging bool
	Blackout       bool
	Clear          bool

	StartFlood bool
	StopFlood  bool

	// RGB commands.
	RGBPosition       Position
	RGBStaticOff      bool
	RGBStaticOn       bool
	RGBStaticFadeUp   bool
	RGBStaticColors   []StaticColorButton
	RGBPlayStaticOnce bool
	RGBFade           int

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
	SwiTch             Switch
	State              State
	CurrentSwitchState int
	Override           Override
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
	ScannerColor color.RGBA
	BaseColor    color.RGBA
	Color        color.RGBA
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
	LastColor color.RGBA
	Color     color.RGBA
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
		Action: Reveal,
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

func StartStaticSequences(sequences []*Sequence, commandChannels []chan Command) {
	for sequenceNumber := range sequences {
		cmd := Command{
			Action: Normal,
		}
		SendCommandToSequence(sequenceNumber, cmd, commandChannels)
	}
}

// Colors are selected from a pallete of 8 colors, this function takes 0-9 (repeating 4 time) and
// returns the color array
func GetColorButtonsArray(colorIn int) color.RGBA {

	switch colorIn {
	case 0:
		return colors.Red
	case 1:
		return colors.Orange
	case 2:
		return colors.Yellow
	case 3:
		return colors.Green
	case 4:
		return colors.Cyan
	case 5:
		return colors.Blue
	case 6:
		return colors.Purple
	case 7:
		return colors.Magenta
	case 8:
		return colors.White
	case 9:
		return colors.Black
	case 10:
		return colors.Red
	case 11:
		return colors.Orange
	case 12:
		return colors.Yellow
	case 13:
		return colors.Green
	case 14:
		return colors.Cyan
	case 15:
		return colors.Blue
	case 16:
		return colors.Purple
	case 17:
		return colors.Magenta
	case 18:
		return colors.White
	case 19:
		return colors.Black
	case 20:
		return colors.Red
	case 21:
		return colors.Orange
	case 22:
		return colors.Yellow
	case 23:
		return colors.Green
	case 24:
		return colors.Cyan
	case 25:
		return colors.Blue
	case 26:
		return colors.Purple
	case 27:
		return colors.Magenta
	case 28:
		return colors.White
	case 29:
		return colors.Black
	case 30:
		return colors.Red
	case 31:
		return colors.Orange
	case 32:
		return colors.Yellow
	case 33:
		return colors.Green
	case 34:
		return colors.Cyan
	case 35:
		return colors.Blue
	case 36:
		return colors.Purple
	case 37:
		return colors.Magenta
	case 38:
		return colors.White
	case 39:
		return colors.Black
	case 40:
		return colors.Red
	case 41:
		return colors.Orange
	case 42:
		return colors.Yellow
	case 43:
		return colors.Green
	case 44:
		return colors.Cyan
	case 45:
		return colors.Blue
	case 46:
		return colors.Purple
	case 47:
		return colors.Magenta
	case 48:
		return colors.White
	case 49:
		return colors.Black
	case 50:
		return colors.Red
	case 51:
		return colors.Orange
	case 52:
		return colors.Yellow
	case 53:
		return colors.Green
	case 54:
		return colors.Cyan
	case 55:
		return colors.Blue
	case 56:
		return colors.Purple
	case 57:
		return colors.Magenta
	case 58:
		return colors.White
	case 59:
		return colors.Black
	case 60:
		return colors.Red
	case 61:
		return colors.Orange
	case 62:
		return colors.Yellow
	case 63:
		return colors.Green
	case 64:
		return colors.Cyan
	case 65:
		return colors.Blue
	case 66:
		return colors.Purple
	case 67:
		return colors.Magenta
	case 68:
		return colors.White
	case 69:
		return colors.Black
	}

	return color.RGBA{}
}

func GetColorArrayByNames(names []string) ([]color.RGBA, error) {

	colors := []color.RGBA{}
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

// Convert my common.Color RGB into color.RGBA used by the fyne.io GUI library.
func ConvertRGBtoRGBA(alight color.RGBA) color.RGBA {
	RGBAcolor := color.RGBA{}
	RGBAcolor.R = uint8(alight.R)
	RGBAcolor.G = uint8(alight.G)
	RGBAcolor.B = uint8(alight.B)
	RGBAcolor.A = 255
	return RGBAcolor
}

func GetRGBColorByName(colorIn string) (color.RGBA, error) {

	if debug {
		fmt.Printf("Looking for color %s\n", colorIn)
	}

	switch colorIn {
	case "Red":
		return colors.Red, nil

	case "Orange":
		return colors.Orange, nil

	case "Yellow":
		return colors.Yellow, nil

	case "Green":
		return colors.Green, nil

	case "Cyan":
		return colors.Cyan, nil

	case "Blue":
		return colors.Blue, nil

	case "Purple":
		return colors.Purple, nil

	case "Pink":
		return colors.Pink, nil

	case "White":
		return colors.White, nil

	case "Light Blue":
		return colors.LightBlue, nil

	case "Black":
		return colors.Black, nil

	}
	return color.RGBA{}, fmt.Errorf("GetRGBColorByName: color not found: %s", colorIn)
}

func GetColorNameByRGB(colorIn color.RGBA) string {
	switch colorIn {
	case colors.LightBlue:
		return "LightBlue"
	case colors.Red:
		return "Red"
	case colors.Orange:
		return "Orange"
	case colors.Yellow:
		return "Yellow"
	case colors.Green:
		return "Green"
	case colors.Cyan:
		return "Cyan"
	case colors.Blue:
		return "Blue"
	case colors.Purple:
		return "Purple"
	case colors.Pink:
		return "Pink"
	case colors.Magenta:
		return "Magenta"

	case colors.Crimson:
		return "Crimson"
	case colors.DarkOrange:
		return "DarkOrange"
	case colors.Gold:
		return "Gold"
	case colors.ForestGreen:
		return "ForestGreen"
	case colors.Aqua:
		return "Aqua"
	case colors.SkyBlue:
		return "SkyBlue"
	case colors.DarkPurple:
		return "DarkPurple"
	case colors.Pink:
		return "Pink"

	case colors.Salmon:
		return "Salmon"
	case colors.LightOrange:
		return "LightOrange"
	case colors.Olive:
		return "Olive"
	case colors.LawnGreen:
		return "LawnGreen"
	case colors.Teal:
		return "Teal"
	case colors.LightBlue:
		return "LightBlue"
	case colors.Violet:
		return "Violet"
	case colors.White:
		return "White"

	case colors.Black:
		return "Black"
	}

	return "White"
}

func HowManyColorsInSteps(steps []Step) (colors []color.RGBA) {

	if debug {
		fmt.Printf("HowManyColorsInSteps \n")
	}
	colorMap := make(map[color.RGBA]bool)
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

func HowManyColorsInPositions(positionsMap map[int]Position) (colors []color.RGBA) {

	colorMap := make(map[color.RGBA]bool)
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

func HowManyStepColors(steps []Step) (colors []color.RGBA) {

	colorMap := make(map[color.RGBA]bool)
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

// Get an upto date copy of the sequence.
func RefreshSequence(selectedSequence int, commandChannels []chan Command, updateChannels []chan Sequence) *Sequence {

	cmd := Command{
		Action: GetUpdatedSequence,
	}

	SendCommandToSequence(selectedSequence, cmd, commandChannels)
	newSequence := <-updateChannels[selectedSequence]
	return &newSequence
}

func ShowStaticButtons(sequence *Sequence, staticFlashing bool, eventsForLaunchpad chan ALight, guiButtons chan ALight) {

	var sequenceNumber int
	if sequence.Number == 4 {
		sequenceNumber = 2
	} else {
		sequenceNumber = sequence.Number
	}

	if debug {
		fmt.Printf("%d: ShowStaticButtons\n", sequenceNumber)
	}

	for fixtureNumber, staticColorButton := range sequence.StaticColors {

		// Only the first 8 colors are used for static color defaults.
		if fixtureNumber > 7 {
			break
		}

		if staticColorButton.Enabled {
			if staticColorButton.Flash || staticFlashing {
				onColor := color.RGBA{R: staticColorButton.Color.R, G: staticColorButton.Color.G, B: staticColorButton.Color.B}
				FlashLight(Button{X: fixtureNumber, Y: sequenceNumber}, onColor, colors.Black, eventsForLaunchpad, guiButtons)
			} else {
				LightLamp(Button{X: fixtureNumber, Y: sequenceNumber}, staticColorButton.Color, sequence.Master, eventsForLaunchpad, guiButtons)
			}
		}
	}
}

func ClearSelectedRowOfButtons(selectedSequence int, eventsForLauchpad chan ALight, guiButtons chan ALight) {
	if debug {
		fmt.Printf("%d: ClearSelectedRowOfButtons\n", selectedSequence)
	}
	// TODO replace with constants for switch and chase sequence numbers.
	if selectedSequence == 4 || selectedSequence == 3 {
		return
	}
	for x := 0; x < 8; x++ {
		LightLamp(Button{X: x, Y: selectedSequence}, colors.Black, MIN_DMX_BRIGHTNESS, eventsForLauchpad, guiButtons)
		LabelButton(x, selectedSequence, "", guiButtons)
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

func ShowRunningStatus(runningState bool, eventsForLaunchpad chan ALight, guiButtons chan ALight) {
	if runningState {
		LightLamp(RUNNING_BUTTON, colors.Green, MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)
	} else {
		LightLamp(RUNNING_BUTTON, colors.White, MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)
	}
}

func ShowStrobeButtonStatus(state bool, eventsForLaunchpad chan ALight, guiButtons chan ALight) {
	if state {
		FlashLight(STROBE_BUTTON, colors.White, colors.Magenta, eventsForLaunchpad, guiButtons)
		return
	}
	LightLamp(STROBE_BUTTON, colors.White, MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)
}

func LabelButton(X int, Y int, label string, guiButtons chan ALight) {
	if debug {
		fmt.Printf("Label Button  X:%d  Y:%d  with %s\n", X, Y, label)
	}
	// Send message to GUI
	event := ALight{
		UpdateLabel: true,
		Button: Button{
			X: X,
			Y: Y + 1,
		},
		Label: label,
	}
	guiButtons <- event
}

// LightOn Turn on a Light.
func LightLamp(button Button, color color.RGBA, master int, eventsForLauchpad chan ALight, guiButtons chan ALight) {
	if debug {
		fmt.Printf("LightLamp  X:%d  Y:%d Red %d Green %d Blue %d Brightnes %d\n", button.X, button.Y, color.R, color.G, color.B, master)
	}
	// Send message to Novation Launchpad.
	event := ALight{
		Button:     button,
		Brightness: master,
		Red:        color.R,
		Green:      color.G,
		Blue:       color.B,
		Flash:      false,
	}
	eventsForLauchpad <- event // Event will be received by pkg/launchpad/launchpad.go ListenAndSendToLaunchPad()

	// Send message to fyne.io GUI.
	event = ALight{
		Button: Button{
			X: button.X,
			Y: button.Y + 1,
		},
		Brightness: master,
		Red:        color.R,
		Green:      color.G,
		Blue:       color.B,
		Flash:      false,
	}
	guiButtons <- event // Event will be received by dmxlights.go by pkg/gui/gui.go ListenAndSendToGUI()
}

func UpdateColorDisplay(control ColorDisplayControl, guiButtons chan ALight) {
	if debug {
		fmt.Printf("UpdateColorDisplay: control %+v\n", control)
	}
	event := ALight{
		ColorDisplay:        true,
		ColorDisplayControl: control,
	}
	guiButtons <- event // Event will be received by dmxlights.go by pkg/gui/gui.go ListenAndSendToGUI()
}

func GetColorList(colors []color.RGBA) ColorDisplayControl {

	if debug {
		fmt.Printf("GetColorList Colors=%+v\n", colors)
	}

	control := ColorDisplayControl{}

	for _, color := range colors {
		found := GetColorNameByRGB(color)
		switch {
		case found == "Red":
			control.Red = true
		case found == "Orange":
			control.Orange = true
		case found == "Yellow":
			control.Yellow = true
		case found == "Green":
			control.Green = true
		case found == "Cyan":
			control.Cyan = true
		case found == "Blue":
			control.Blue = true
		case found == "Purple":
			control.Purple = true
		case found == "Magenta":
			control.Magenta = true

		case found == "Crimson":
			control.Crimson = true
		case found == "DarkOrange":
			control.DarkOrange = true
		case found == "Gold":
			control.Gold = true
		case found == "ForestGreen":
			control.ForestGreen = true
		case found == "Aqua":
			control.Aqua = true
		case found == "SkyBlue":
			control.SkyBlue = true
		case found == "DarkPurple":
			control.DarkPurple = true
		case found == "Pink":
			control.Pink = true

		case found == "Salmon":
			control.Salmon = true
		case found == "LightOrange":
			control.LightOrange = true
		case found == "Olive":
			control.Olive = true
		case found == "LawnGreen":
			control.LawnGreen = true
		case found == "Teal":
			control.Teal = true
		case found == "LightBlue":
			control.LightBlue = true
		case found == "Violet":
			control.Violet = true
		case found == "White":
			control.White = true
		}

	}

	if debug {
		fmt.Printf("GetColorList Control=%+v\n", control)
	}

	return control

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

func FlashLight(button Button, onColor color.RGBA, offColor color.RGBA, eventsForLauchpad chan ALight, guiButtons chan ALight) {

	// Now ask the fixture lamp to flash on the launch pad by sending an event.
	e := ALight{
		Button:     button,
		Brightness: 255,
		Flash:      true,
		OnColor:    onColor,
		OffColor:   offColor,
	}
	eventsForLauchpad <- e // Event will be received by pkg/launchpad/launchpad.go ListenAndSendToLaunchPad()

	// Send message to GUI
	event := ALight{
		Button:     Button{X: button.X, Y: button.Y + 1},
		Brightness: 255,
		Flash:      true,
		OnColor:    onColor,
		OffColor:   offColor,
	}
	guiButtons <- event // Event will be received by dmxlights.go by pkg/gui/gui.go ListenAndSendToGUI()

}

// InvertColor just reverses the DMX values.
func InvertColor(colorIn color.RGBA) (out color.RGBA) {

	out.R = ReverseDmx(colorIn.R)
	out.G = ReverseDmx(colorIn.G)
	out.B = ReverseDmx(colorIn.B)
	out.A = 255

	return out
}

// Takes a DMX value 1-255 and reverses the value.
func ReverseDmx(n uint8) uint8 {
	in := make(map[uint8]uint8, 255)
	var y uint8 = 255
	var x uint8

	for x = 0; x < 255; x++ {

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
			staticColorButton.Enabled = true
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

func GetFadeOnValues(brightness int, fadeOnTime int) []int {

	out := []int{}

	var x int

	for x = 0; x < fadeOnTime; x++ {
		out = append(out, brightness)
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

// Used to convert a speed to a millisecond time.
func SetSpeed(commandSpeed int) (Speed time.Duration) {
	if commandSpeed == 0 {
		Speed = 3500
	}
	if commandSpeed == 1 {
		Speed = 3000
	}
	if commandSpeed == 2 {
		Speed = 2500
	}
	if commandSpeed == 3 {
		Speed = 1800
	}
	if commandSpeed == 4 {
		Speed = 1500
	}
	if commandSpeed == 5 {
		Speed = 1000
	}
	if commandSpeed == 6 {
		Speed = 750
	}
	if commandSpeed == 7 {
		Speed = 500
	}
	if commandSpeed == 8 {
		Speed = 250
	}
	if commandSpeed == 9 {
		Speed = 150
	}
	if commandSpeed == 10 {
		Speed = 125
	}
	if commandSpeed == 11 {
		Speed = 100
	}
	if commandSpeed == 12 {
		Speed = 50
	}
	return Speed * time.Millisecond
}

func GetSize(size int) int {

	switch size {
	case 1:
		return 1
	case 2:
		return 5
	case 3:
		return 15
	case 4:
		return 25
	case 5:
		return 35
	case 6:
		return 45
	case 7:
		return 55
	case 8:
		return 65
	case 9:
		return 75
	case 10:
		return 85
	}
	return 0
}

func FormatLabel(label string) string {
	// replace any spaces with new lines.
	// new lines are represented by a dot in code beneath us.
	return strings.Replace(label, " ", ".", -1)
}

func newColorPicker() []ColorPicker {

	colors := []ColorPicker{

		{ID: 0, X: 0, Y: 0, Name: "Red", Code: 0x48, Color: colors.Red},
		{ID: 1, X: 1, Y: 0, Name: "Orange", Code: 0x09, Color: colors.Orange},
		{ID: 2, X: 2, Y: 0, Name: "Yellow", Code: 0x0d, Color: colors.Yellow},
		{ID: 3, X: 3, Y: 0, Name: "Green", Code: 0x4C, Color: colors.Green},
		{ID: 4, X: 4, Y: 0, Name: "Cyan", Code: 0x25, Color: colors.Cyan},
		{ID: 5, X: 5, Y: 0, Name: "Blue", Code: 0x4f, Color: colors.Blue},
		{ID: 6, X: 6, Y: 0, Name: "Purple", Code: 0x32, Color: colors.Purple},
		{ID: 7, X: 7, Y: 0, Name: "Magenta", Code: 0x35, Color: colors.Magenta},

		{ID: 8, X: 0, Y: 1, Name: "Crimson", Code: 0x38, Color: colors.Crimson},
		{ID: 9, X: 1, Y: 1, Name: "Dark Orange", Code: 0x0a, Color: colors.DarkOrange},
		{ID: 10, X: 2, Y: 1, Name: "Gold", Code: 0x61, Color: colors.Gold},
		{ID: 11, X: 3, Y: 1, Name: "Forest Green", Code: 0x1b, Color: colors.ForestGreen},
		{ID: 12, X: 4, Y: 1, Name: "Aqua", Code: 0x20, Color: colors.Aqua},
		{ID: 13, X: 5, Y: 1, Name: "Sky Blue", Code: 0x25, Color: colors.SkyBlue},
		{ID: 14, X: 6, Y: 1, Name: "Dark Purple", Code: 0x32, Color: colors.DarkPurple},
		{ID: 15, X: 7, Y: 1, Name: "Pink", Code: 0x34, Color: colors.Pink},

		{ID: 16, X: 0, Y: 2, Name: "Salmon", Code: 0x6b, Color: colors.Salmon},
		{ID: 17, X: 1, Y: 2, Name: "Light Orange", Code: 0x0c, Color: colors.LightOrange},
		{ID: 18, X: 2, Y: 2, Name: "Olive", Code: 0x10, Color: colors.Olive},
		{ID: 19, X: 3, Y: 2, Name: "Lawn green", Code: 0x13, Color: colors.LawnGreen},
		{ID: 20, X: 4, Y: 2, Name: "Teal", Code: 0x44, Color: colors.Teal},
		{ID: 21, X: 5, Y: 2, Name: "Light Blue", Code: 0x20, Color: colors.LightBlue},
		{ID: 22, X: 6, Y: 2, Name: "Violet", Code: 0x5e, Color: colors.Violet},
		{ID: 23, X: 7, Y: 2, Name: "White", Code: 0x03, Color: colors.White},
	}

	return colors
}

func GetLaunchPadCodeByRGBColor(selectedColor color.RGBA) byte {

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
