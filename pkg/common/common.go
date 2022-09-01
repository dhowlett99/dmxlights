package common

import (
	"fmt"
	"sync"
	"time"

	"github.com/rakyll/launchpad/mk3"
)

const debug = false

const DefaultScannerSize = 120
const MaxScannerSize = 120
const MaxRGBSize = 120
const MaxRGBFade = 10
const DefaultPattern = 0
const DefaultRGBSize = 1
const DefaultRGBFade = 1
const DefaultScannerFade = 10
const DefaultSpeed = 7
const DefaultRGBShift = 0
const DefaultScannerShift = 0
const DefaultScannerCoordinates = 0

const MaxBrightness = 255

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
}

type Color struct {
	R            int
	G            int
	B            int
	MasterDimmer int
}

type Value struct {
	Channel int16
	Setting int16
}

type State struct {
	Name        string
	Label       string
	Values      []Value
	ButtonColor Color
}

type Switch struct {
	Name         string
	Number       int
	Label        string
	CurrentState int
	Description  string
	States       []State
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
	UpdateSwitch
	UpdateSwitchPositions
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
	ClearSequenceColor
	Static
	MasterBrightness
	UpdateNumberCoordinates
	UpdateOffsetPan
	UpdateOffsetTilt
	EnableAllScanners
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
	Name                       string                      // Sequence name.
	Label                      string                      // Sequence label.
	Description                string                      // Sequence description.
	Number                     int                         // Sequence number.
	Run                        bool                        // True if this sequence is running.
	Bounce                     bool                        // True if this sequence is bouncing.
	Invert                     bool                        // True if RGB sequence patten is inverted.
	Hide                       bool                        // Hide is used to hide sequence buttons when using function keys.
	Type                       string                      // Type of sequnece, current valid values are :- rgb, scanner,  or switch.
	Master                     int                         // Master Brightness
	StrobeSpeed                int                         // Strobe speed.
	RGBShift                   int                         // RGB shift.
	CurrentSpeed               time.Duration               // Sequence speed represented as a duration.
	Speed                      int                         // Sequence speed represented by a short number.
	MusicTrigger               bool                        // Is this sequence in music trigger mode.
	Ring                       bool                        // A ring is when a music triggers a ring of events for a scannner.
	Beat                       bool                        // Used by a ring to indicate the start of a ring.
	RingCounter                int                         // Used to keep track of the number of events in this ring, usually a multiple of the number of steps in the sequence.
	Blackout                   bool                        // Flag to indicate we're in blackout mode.
	CurrentColors              []Color                     // Storage for the colors in a sequence.
	SequenceColors             []Color                     // Temporay storage for changing sequence colors.
	Color                      int                         // Index into current sequnece colors.
	Steps                      []Step                      // Steps in this sequence.
	NumberSteps                int                         // Holds the number of steps this sequence has. Will change if you change size, fade times etc.
	NumberFixtures             int                         // Number of fixtures for this sequence.
	RGBPositions               map[int]Position            // One set of Fixture positions for RGB devices. index is position number.
	ScannerPositions           map[int]map[int][]Position  // Scanner Fixture positions decides where a fixture is in a give set of sequence steps. Index Positions.
	AutoColor                  bool                        // Sequence is going to automatically change the color.
	AutoPattern                bool                        // Sequence is going to automatically change the pattern.
	GuiFunctionLabels          [8]string                   // Storage for the function key labels for this sequence.
	GuiFixtureLabels           []string                    // Storage for the fixture labels. Used for scanner names.
	Pattern                    Pattern                     // Contains fixtures and steps info.
	RGBAvailablePatterns       map[int]Pattern             // Available patterns for the RGB fixtures.
	RGBAvailableColors         []StaticColorButton         // Available colors for the RGB fixtures.
	RGBColor                   int                         // The selected RGB fixture color.
	RGBFade                    int                         // RGB Fade time
	RGBSize                    int                         // RGB Fade size
	SavedSequenceColors        []Color                     // Used for updating the color in a sequence.
	RecoverSequenceColors      bool                        // Storage for recovering sequence colors, when you come out of automatic color change.
	SaveColors                 bool                        // Indicate we should save colors in this sequence. used for above.
	Mode                       string                      // Tells sequnece if we're in sequence (chase) or static (static colors) mode.
	StaticColors               []StaticColorButton         // Used in static color editing
	Static                     bool                        // We're a static sequence.
	PlayStaticOnce             bool                        // Play a static scene only once.
	PlaySwitchOnce             bool                        // Play a switch sequence scene only once.
	StartFlood                 bool                        // We're in flood mode.
	StopFlood                  bool                        // We're not in flood mode.
	StartStrobe                bool                        // We're in strobe mode.
	StopStrobe                 bool                        // We're not in strobe mode.
	FloodPlayOnce              bool                        // Play the flood sceme only once.
	FloodSelectedSequence      map[int]bool                // A map that remembers who is in flood mode.
	ScannersTotal              int                         // Total number of scanners in this sequence.
	ScannerAvailableColors     map[int][]StaticColorButton // Available colors for this scanner.
	ScannerAvailableGobos      map[int][]StaticColorButton // Available gobos for this scanner.
	ScannerAvailablePatterns   map[int]Pattern             // Available patterns for this scanner.
	ScannersAvailable          []StaticColorButton         // Holds a set of red buttons, one for every available fixture.
	SelectedPattern            int                         // The selected pattern.
	ScannerSize                int                         // The selected scanner size.
	ScannerShift               int                         // Used for shifting scanners patterns apart.
	ScannerGobo                int                         // The selected gobo.
	ScannerChase               bool                        // Chase the scanner shutters instead of allways being on.
	ScannerColor               map[int]int                 // Eight scanners per sequence, each can have their own color.
	ScannerCoordinates         []int                       // Number of scanner coordinates.
	ScannerSelectedCoordinates int                         // index into scanner coordinates.
	ScannerOffsetPan           int                         // Offset for pan values.
	ScannerOffsetTilt          int                         // Offset for tilt values.
	ScannerStateMutex          *sync.RWMutex               // Mutex to protect the  disable maps from syncronous access.
	ScannerState               map[int]ScannerState        // Map of fixtures which are disabled.
	DisableOnceMutex           *sync.RWMutex               // Mutex to protect the  disable maps from syncronous access.
	DisableOnce                map[int]bool                // Map used to play disable only once.
	UpdateShift                bool                        // Command to update the shift.
	UpdatePattern              bool                        // Flag to indicate we're going to change the RGB pattern.
	UpdateSequenceColor        bool                        // Command to update the sequence colors.
	Functions                  []Function                  // Storage for the sequence functions.
	FunctionMode               bool                        // This sequence is in function mode.
	Switches                   []Switch                    // A switch sequence stores its data in here.
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
	CommmandChannels     []chan Command
	ReplyChannels        []chan Sequence
	SoundTriggerChannels []chan Command
	UpdateChannels       []chan Sequence
}

type Hit struct {
	X int
	Y int
}

type Step struct {
	Fixtures []Fixture
	Type     string
}

type FixtureCommand struct {
	Step           int
	Type           string
	SequenceNumber int

	// Command commands.
	StrobeSpeed int
	Master      int
	Blackout    bool
	Hide        bool
	Invert      bool

	// RGB commands.
	RGBPosition     Position
	RGBSize         int
	RGBFade         int
	RGBStartFlood   bool
	RGBStopFlood    bool
	RGBStatic       bool
	RGBStaticColors []StaticColorButton

	// Scanner Commands.
	ScannerColor           map[int]int
	ScannerPosition        map[int][]Position
	ScannerState           map[int]ScannerState
	ScannerDisableOnce     map[int]bool
	ScannerChase           bool
	ScannerAvailableColors map[int][]StaticColorButton
	ScannerSelectedGobo    int
	ScannerOffsetPan       int
	ScannerOffsetTilt      int
}

type Position struct {
	// RGB
	Fixtures       map[int]Fixture
	Color          Color
	PositionNumber int
	// Scanner
	ScannerNumber  int
	StartPosition  int
	Pan            int
	PanMaxDegrees  *int
	Tilt           int
	TiltMaxDegrees *int
	Shutter        int
	Gobo           int
}

type PreFadeDetails struct {
	FadeValue    int
	MasterDimmer int
	Color        Color
}

// A fixture can have any or some of the
// following, depending if its a light or
// a scanner.
type Fixture struct {
	Name         string
	Label        string
	Type         string
	MasterDimmer int
	Colors       []Color
	Pan          int
	Tilt         int
	Shutter      int
	Gobo         int
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
	SequenceNumber int
	State          bool
	Gain           float32
	BPM            int
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

	}
	return Color{}
}

func GetRGBColorByName(color string) Color {
	switch color {
	case "Red":
		return Color{R: 255, G: 0, B: 0}

	case "Orange":
		return Color{R: 255, G: 111, B: 0}

	case "Yellow":
		return Color{R: 255, G: 255, B: 0}

	case "Green":
		return Color{R: 0, G: 255, B: 0}

	case "Cyan":
		return Color{R: 0, G: 255, B: 255}

	case "Blue":
		return Color{R: 0, G: 0, B: 255}

	case "Purple":
		return Color{R: 100, G: 0, B: 255}

	case "Pink":
		return Color{R: 255, G: 0, B: 255}

	case "White":
		return Color{R: 255, G: 255, B: 255}

	case "Light Blue":
		return Color{R: 0, G: 196, B: 255}

	case "Black":
		return Color{R: 0, G: 0, B: 0}

	}
	return Color{R: 0, G: 0, B: 0}
}

func GetLaunchPadColorCodeByRGB(color Color) (code byte) {
	switch color {
	case Color{R: 0, G: 100, B: 255}:
		return 0x2a // Light Blue
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
	sequence.AutoColor = sequence.Functions[Function2_Auto_Color].State

	// Map the auto pattern change setting.
	sequence.AutoPattern = sequence.Functions[Function3_Auto_Pattern].State

	// Map bounce function to sequence bounce setting.
	sequence.Bounce = sequence.Functions[Function4_Bounce].State

	// Map color selection function.
	if sequence.Functions[Function5_Color].State {
		sequence.PlayStaticOnce = true
	}

	// Map static function.
	if sequence.Type != "scanner" {
		sequence.Static = sequence.Functions[Function6_Static_Gobo].State
		if sequence.Functions[Function6_Static_Gobo].State {
			sequence.PlayStaticOnce = true
			sequence.Hide = true
		}
	}

	// Map invert function.
	sequence.Invert = sequence.Functions[Function7_Invert_Chase].State
	// Map scanner chase mode. Uses same function key as above.
	sequence.ScannerChase = sequence.Functions[Function7_Invert_Chase].State

	// Map music trigger function.
	sequence.MusicTrigger = sequence.Functions[Function8_Music_Trigger].State
	if sequence.Functions[Function8_Music_Trigger].State {
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

func HowManyScannerColors(positionsMap map[int][]Position) (colors []Color) {

	colorMap := make(map[Color]bool)
	for _, positions := range positionsMap {
		for _, position := range positions {
			colorMap[position.Color] = true
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

func ShowStrobeStatus(state bool, eventsForLaunchpad chan ALight, guiButtons chan ALight) {
	if state {
		FlashLight(8, 6, White, Black, eventsForLaunchpad, guiButtons)
		return
	}
	LightLamp(ALight{X: 8, Y: 6, Brightness: 255, Red: 255, Green: 255, Blue: 255}, eventsForLaunchpad, guiButtons)
}

// ListenAndSendToLaunchPad is the thread that listens for events to send to
// the launch pad.  It is thread safe and is the only thread talking to the
// launch pad. A channel is used to queue the events to be sent.
func ListenAndSendToLaunchPad(eventsForLauchpad chan ALight, pad *mk3.Launchpad) {
	for {

		// Wait for the event.
		alight := <-eventsForLauchpad

		// Wait for a few millisecond so the launchpad and the gui step at the same time
		time.Sleep(14 * time.Microsecond)

		// We're in standard turn the light on.
		if !alight.Flash {

			// Take into account the brightness. Divide by 2 because launch pad is 1-127.
			Red := ((float64(alight.Red) / 2) / 100) * (float64(alight.Brightness) / 2.55)
			Green := ((float64(alight.Green) / 2) / 100) * (float64(alight.Brightness) / 2.55)
			Blue := ((float64(alight.Blue) / 2) / 100) * (float64(alight.Brightness) / 2.55)

			// Now light the launchpad button.
			err := pad.Light(alight.X, alight.Y, int(Red), int(Green), int(Blue))
			if err != nil {
				fmt.Printf("error writing to launchpad %e\n" + err.Error())
			}

			// Now we're been asked go flash this button.
		} else {
			// Now light the launchpad button.
			if debug {
				fmt.Printf("Want Color %+v LaunchPad On Code is %x\n", alight.OnColor, GetLaunchPadColorCodeByRGB(alight.OnColor))
				fmt.Printf("Want Color %+v LaunchPad Off Code is %x\n", alight.OffColor, GetLaunchPadColorCodeByRGB(alight.OffColor))
			}
			err := pad.FlashLight(alight.X, alight.Y, int(GetLaunchPadColorCodeByRGB(alight.OnColor)), int(GetLaunchPadColorCodeByRGB(alight.OffColor)))
			if err != nil {
				fmt.Printf("flash: error writing to launchpad %e\n" + err.Error())
			}

		}
	}
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
		fmt.Printf("LightLamp  X:%d  Y:%d\n", Light.X, Light.Y)
	}
	// Send message to Novation Launchpad.
	event := ALight{
		X:          Light.X,
		Y:          Light.Y,
		Brightness: Light.Brightness,
		Red:        Light.Red,
		Green:      Light.Green,
		Blue:       Light.Blue,
		Flash:      false,
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
		Flash:      false,
		OnColor:    Light.OnColor,
		OffColor:   Light.OffColor,
		Label:      Light.Label,
	}
	guiButtons <- event
}

func UpdateStatusBar(status string, which string, guiButtons chan ALight) {
	// Send message to fyne.io GUI.
	event := ALight{
		UpdateStatus: true,
		Status:       status,
		Which:        which,
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
