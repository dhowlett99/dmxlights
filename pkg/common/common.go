package common

import (
	"fmt"
	"sync"
	"time"

	"github.com/dhowlett99/dmxlights/pkg/pad"
)

const debug = false

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
	FlashStopChannel chan bool
}

type Color struct {
	R int
	G int
	B int
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

type Patten struct {
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
const UpdateMode = 1
const UpdateStatic = 2
const UpdateStaticColor = 3
const UpdateSequenceColor = 4
const PlayStaticOnce = 5
const PlaySwitchOnce = 6
const UnHide = 7
const Hide = 8
const Start = 9
const Stop = 10
const ReadConfig = 11
const LoadConfig = 12
const UpdateSpeed = 13
const UpdatePatten = 14
const SelectPatten = 15
const IncreaseFade = 16
const DecreaseFade = 17
const UpdateSize = 18
const UpdateScannerSize = 19
const Blackout = 20
const Normal = 21
const SoftFadeOn = 22
const SoftFadeOff = 23
const UpdateColor = 24
const UpdateFunctionMode = 25
const FunctionMode = 26
const UpdateFunctions = 27
const GetUpdatedSequence = 28
const UpdateSwitch = 29
const UpdateSwitchPositions = 30
const Inverted = 31
const UpdateGobo = 32
const Flood = 33
const NoFlood = 34
const UpdateAutoColor = 35
const AutoColor = 36
const UpdateAutoPatten = 37
const AutoPatten = 38
const ToggleFixtureState = 39
const FixtureState = 40
const UpdateShift = 41
const UpdateScannerColor = 42
const ClearSequenceColor = 43
const Static = 44
const MasterBrightness = 45
const UpdateNumberCoordinates = 46
const UpdateOffsetPan = 47
const UpdateOffsetTilt = 48

const DefaultScannerSize = 120

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
	Hide                       bool                        // Hide is used to hide sequence buttons when using function keys.
	Type                       string                      // Type of sequnece, current valid values are :- rgb, scanner,  or switch.
	Master                     int                         // Master Brightness
	CurrentSpeed               time.Duration               // Sequence speed represented as a duration.
	Speed                      int                         // Sequence speed represented by a short number.
	MusicTrigger               bool                        // Is this sequence in music trigger mode.
	Blackout                   bool                        // Flag to indicate we're in blackout mode.
	CurrentColors              []Color                     // Storage for the colors in a sequence.
	SequenceColors             []Color                     // Temporay storage for changing sequence colors.
	Color                      int                         // Index into current sequnece colors.
	Steps                      []Step                      // Steps in this sequence.
	NumberSteps                int                         // Holds the number of steps this sequence has. Will change if you change size, fade times etc.
	Positions                  map[int][]Position          // Positions decides where a fixture is in a give set of sequence steps.
	AutoColor                  bool                        // Sequence is going to automatically change the color.
	AutoPatten                 bool                        // Sequence is going to automatically change the patten.
	GuiFunctionLabels          [8]string                   // Storage for the function key labels for this sequence.
	GuiFixtureLabels           []string                    // Storage for the fixture labels. Used for scanner names.
	Patten                     Patten                      // Contains fixtures and steps info.
	PattenInverted             bool                        // The patten is inverted.
	RGBAvailablePattens        map[int]Patten              // Available pattens for the RGB fixtures.
	RGBAvailableColors         []StaticColorButton         // Available colors for the RGB fixtures.
	RGBColor                   int                         // The selected RGB fixture color.
	FadeSpeed                  int                         // Fade Speed
	FadeTime                   time.Duration               // Fade time
	Size                       int                         // Fade size
	SavedSequenceColors        []Color                     // Used for updating the color in a sequence.
	SelectedRGBPatten          int                         // Selected RGB patten.
	RecoverSequenceColors      bool                        // Storage for recovering sequence colors, when you come out of automatic color change.
	SaveColors                 bool                        // Indicate we should save colors in this sequence. used for above.
	Mode                       string                      // Tells sequnece if we're in sequence (chase) or static (static colors) mode.
	StaticColors               []StaticColorButton         // Used in static color editing
	Static                     bool                        // We're a static sequence.
	PlayStaticOnce             bool                        // Play a static scene only once.
	PlaySwitchOnce             bool                        // Play a switch sequence scene only once.
	Flood                      bool                        // We're in flood mode.
	NoFlood                    bool                        // We're not in flood mode.
	FloodPlayOnce              bool                        // Play the flood sceme only once.
	FloodSelectedSequence      map[int]bool                // A map that remembers who is in flood mode.
	ScannersTotal              int                         // Total number of scanners in this sequence.
	ScannerAvailableColors     map[int][]StaticColorButton // Available colors for this scanner.
	ScannerAvailableGobos      map[int][]StaticColorButton // Available gobos for this scanner.
	ScannerAvailablePattens    map[int]Patten              // Available pattens for this scanner.
	ScannersAvailable          []StaticColorButton         // Holds a set of red buttons, one for every available fixture.
	ScannerPatten              int                         // The selected scanner patten.
	ScannerSize                int                         // The selected scanner size.
	ScannerShift               int                         // Used for shifting scanners patterns apart.
	ScannerGobo                int                         // The selected gobo.
	ScannerChase               bool                        // Chase the scanner shutters instead of allways being on.
	ScannerColor               map[int]int                 // Eight scanners per sequence, each can have their own color.
	ScannerCoordinates         []int                       // Number of scanner coordinates.
	ScannerSelectedCoordinates int                         // index into scanner coordinates.
	ScannerOffsetPan           int                         // Offset for pan values.
	ScannerOffsetTilt          int                         // Offset for tilt values.
	FixtureDisabledMutex       *sync.RWMutex               // Mutex to protect the  disable maps from syncronous access.
	FixtureDisabled            map[int]bool                // Map of fixtures which are disabled.
	DisableOnceMutex           *sync.RWMutex               // Mutex to protect the  disable maps from syncronous access.
	DisableOnce                map[int]bool                // Map used to play disable only once.
	UpdateShift                bool                        // Command to update the shift.
	UpdatePatten               bool                        // Flag to indicate we're going to change the RGB patten.
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

// Fixture Command.
type FixtureCommand struct {
	Master                 int
	Hide                   bool
	Tick                   bool
	Config                 bool // Configure fixture.
	Start                  bool
	Steps                  int
	Positions              map[int][]Position
	Type                   string
	StartPosition          int
	CurrentPosition        int
	CurrentSpeed           time.Duration
	Color                  Color
	Speed                  int
	Shift                  int
	Size                   int
	FadeSpeed              int
	FadeTime               time.Duration
	FadeUpTime             time.Duration
	FadeOnTime             time.Duration
	FadeDownTime           time.Duration
	FadeOffTime            time.Duration
	Blackout               bool
	Flood                  bool
	NoFlood                bool
	PlayFloodOnce          bool
	UpdateSequenceColor    bool
	SequenceColor          Color
	SequenceNumber         int
	Inverted               bool
	SelectedGobo           int
	FixtureDisabledMutex   *sync.RWMutex // Mutex to protect the  disable maps from syncronous access.
	FixtureDisabled        map[int]bool
	DisableOnceMutex       *sync.RWMutex // Mutex to protect the  disable once map from syncronous access.
	DisableOnce            map[int]bool
	ScannerChase           bool
	ScannerColor           map[int]int
	Static                 bool
	StaticColors           []StaticColorButton
	AvailableScannerColors map[int][]StaticColorButton
	OffsetPan              int
	OffsetTilt             int
	FixtureLabels          []string
}

type Position struct {
	Fixture        int
	StartPosition  int
	Color          Color
	Pan            int
	PanMaxDegrees  *int
	Tilt           int
	TiltMaxDegrees *int
	Shutter        int
	Gobo           int
}

// A fixture can have any or some of the
// following, depending if its a light or
// a scanner.
type Fixture struct {
	Name           string
	Label          string
	Type           string
	MasterDimmer   int
	Colors         []Color
	Pan            int
	PanMaxDegrees  int
	Tilt           int
	TiltMaxDegrees int
	Shutter        int
	Gobo           int
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
}

// Define the function keys.
const (
	Function1_Patten        = 0 // Set patten mode.
	Function2_Auto_Color    = 1 // Auto Color change.
	Function3_Auto_Patten   = 2 // Auto Patten change
	Function4_Bounce        = 3 // Sequence auto reverses.  doesn't apply in scanner mode.
	Function5_Color         = 4 // Set RGB chase color. or select the scanner color.
	Function6_Static_Gobo   = 5 // Set static color / set scanner gobo.
	Function7_Invert_Chase  = 6 // Invert the RGB colors  / Set scanner chase mode.
	Function8_Music_Trigger = 7 // Music trigger on and off. Both RGB and scanners.
)

func SendCommandToSequence(selectedSequence int, command Command, commandChannels []chan Command) {
	commandChannels[selectedSequence] <- command
}

func SendCommandToAllSequence(selectedSequence int, command Command, commandChannels []chan Command) {
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

func MakeFunctionButtons(sequence Sequence, selectedSequence int, eventsForLauchpad chan ALight, guiButtons chan ALight, channels Channels) {

	// The taget set of buttons.
	ClearSelectedRowOfButtons(selectedSequence, eventsForLauchpad, guiButtons)

	// Get an upto date copy of the sequence.
	cmd := Command{
		Action: ReadConfig,
	}
	SendCommandToSequence(selectedSequence, cmd, channels.CommmandChannels)

	replyChannel := channels.ReplyChannels[selectedSequence]
	sequence = <-replyChannel

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

func ConvertRGBtoPalette(red, green, blue int) (paletteColor int) {
	if red == 255 && green == 0 && blue == 0 {
		return 0x78
	} // Red
	if red == 255 && green == 111 && blue == 0 {
		return 0x60
	} // Orange
	if red == 255 && green == 255 && blue == 0 {
		return 0x7c
	} // Yellow
	if red == 0 && green == 255 && blue == 0 {
		return 0x15
	} // Green
	if red == 0 && green == 255 && blue == 255 {
		return 0x25
	} // Cyan
	if red == 0 && green == 0 && blue == 255 {
		return 0x42
	} // Blue
	if red == 100 && green == 0 && blue == 255 {
		return 0x2d
	} // Purple
	if red == 255 && green == 0 && blue == 255 {
		return 0x35
	} // Pink
	if red == 255 && green == 255 && blue == 255 {
		return 0x03
	} // White
	if red == 0 && green == 0 && blue == 0 {
		return 0x0
	} // Black
	return 0
}

func SetFunctionKeyActions(functions []Function, sequence Sequence) Sequence {

	// Map the auto color change setting.
	sequence.AutoColor = sequence.Functions[Function2_Auto_Color].State

	// Map the auto patten change setting.
	sequence.AutoPatten = sequence.Functions[Function3_Auto_Patten].State

	// Map bounce function to sequence bounce setting.
	if sequence.Type != "scanner" {
		sequence.Bounce = sequence.Functions[Function4_Bounce].State
	}

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
	sequence.PattenInverted = sequence.Functions[Function7_Invert_Chase].State
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

func HowManyColors(positionsMap map[int][]Position) (colors []Color) {

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
func HideColorSelectionButtons(mySequenceNumber int, sequence Sequence, selectedSequence int, eventsForLauchpad chan ALight, guiButtons chan ALight) {
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

func ClearLabelsSelectedRowOfButtons(selectedSequence int, eventsForLauchpad chan ALight, guiButtons chan ALight) {
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

func ShowBottomButtons(tYpe string, eventsForLauchpad chan ALight, guiButtons chan ALight) {

	// Storage for the rgb labels on the bottom row.
	var guiBottomRGBButtons [8]string
	guiBottomRGBButtons[0] = "Speed\nDown"
	guiBottomRGBButtons[1] = "Speed\nUp"
	guiBottomRGBButtons[2] = "Shift\nDown"
	guiBottomRGBButtons[3] = "Shift\nUp"
	guiBottomRGBButtons[4] = "Size\nDown"
	guiBottomRGBButtons[5] = "Size\nUp"
	guiBottomRGBButtons[6] = "Fade\nSoft"
	guiBottomRGBButtons[7] = "Fade\nSharp"

	// Storage for the scanner labels on the bottom row.
	var guiBottomScannerButtons [8]string
	guiBottomScannerButtons[0] = "Speed\nDown"
	guiBottomScannerButtons[1] = "Speed\nUp"
	guiBottomScannerButtons[2] = "Shift\nDown"
	guiBottomScannerButtons[3] = "Shift\nUp"
	guiBottomScannerButtons[4] = "Size\nDown"
	guiBottomScannerButtons[5] = "Size\nUp"
	guiBottomScannerButtons[6] = "Coord\nDown"
	guiBottomScannerButtons[7] = "Coord\nUp"

	//  The bottom row of the Novation Launchpad.
	bottomRow := 7

	if tYpe == "rgb" {
		// Loop through the available functions for this sequence
		for index, button := range guiBottomRGBButtons {
			if debug {
				fmt.Printf("button %+v\n", button)
			}
			LightLamp(ALight{X: index, Y: bottomRow, Brightness: 255, Red: 3, Green: 255, Blue: 255}, eventsForLauchpad, guiButtons)
			LabelButton(index, bottomRow, button, guiButtons)
		}
	}
	if tYpe == "scanner" {
		// Loop through the available functions for this sequence
		for index, button := range guiBottomScannerButtons {
			if debug {
				fmt.Printf("button %+v\n", button)
			}
			LightLamp(ALight{X: index, Y: bottomRow, Brightness: 255, Red: 3, Green: 255, Blue: 255}, eventsForLauchpad, guiButtons)
			LabelButton(index, bottomRow, button, guiButtons)
		}
	}
}

// ListenAndSendToLaunchPad is the thread that listens for events to send to
// the launch pad.  It is thread safe and is the only thread talking to the
// launch pad. A channel is used to queue the events to be sent.
func ListenAndSendToLaunchPad(eventsForLauchpad chan ALight, pad *pad.Pad, LaunchPadFlashButtons [][]ALight) {
	for {

		// Wait for the event.
		alight := <-eventsForLauchpad

		// We're in standard turn the light on.
		if !alight.Flash {

			// We're not flashing.
			// reset this button so it's not flashing.

			// If the LaunchPadFlashButtons array has a true value for this button,
			// Then there must be a thread flashing the lamp right now.
			// So we can assume its listening for a stop command.
			if LaunchPadFlashButtons != nil {
				if LaunchPadFlashButtons[alight.X][alight.Y+1].Flash {
					LaunchPadFlashButtons[alight.X][alight.Y+1].FlashStopChannel <- true
					LaunchPadFlashButtons[alight.X][alight.Y+1].Flash = false
				}
			}

			// Take into account the brightness. Divide by 2 because launch pad is 1-127.
			Red := ((float64(alight.Red) / 2) / 100) * (float64(alight.Brightness) / 2.55)
			Green := ((float64(alight.Green) / 2) / 100) * (float64(alight.Brightness) / 2.55)
			Blue := ((float64(alight.Blue) / 2) / 100) * (float64(alight.Brightness) / 2.55)

			// Now light the launchpad button.
			err := pad.Light(alight.X, alight.Y, int(Red), int(Green), int(Blue))
			if err != nil {
				fmt.Printf("turn on: error writing to launchpad %e\n" + err.Error())
			}

			// Now we're been asked go flash this button.
		} else {
			// Stop any existing flashing.
			if LaunchPadFlashButtons != nil {
				if LaunchPadFlashButtons[alight.X][alight.Y+1].Flash {
					// Set this button to not flashing.
					LaunchPadFlashButtons[alight.X][alight.Y+1].Flash = false
					// Send a message to stop the thread.
					LaunchPadFlashButtons[alight.X][alight.Y+1].FlashStopChannel <- true
					// Wait for the any currently flashing button to stop.
					time.Sleep(10 * time.Millisecond)
				}
			}

			// Now start a new flashing button. Let everyone know that we're flashing.
			if LaunchPadFlashButtons != nil {
				LaunchPadFlashButtons[alight.X][alight.Y+1].Flash = true
			}

			// We create a thread to flash the button.
			go func() {
				for {
					// Turn on.
					// For the math to work we need to convert our ints to floats and then back again.
					Red := ((float64(alight.OnColor.R) / 2) / 100) * (float64(alight.Brightness) / 2.55)
					Green := ((float64(alight.OnColor.G) / 2) / 100) * (float64(alight.Brightness) / 2.55)
					Blue := ((float64(alight.OnColor.B) / 2) / 100) * (float64(alight.Brightness) / 2.55)
					err := pad.Light(alight.X, alight.Y, int(Red), int(Green), int(Blue))
					if err != nil {
						fmt.Printf("flash on: error writing to launchpad %e\n" + err.Error())
					}

					// We wait for a stop message or 250ms which ever comes first.
					select {
					case <-LaunchPadFlashButtons[alight.X][alight.Y+1].FlashStopChannel:
						return
					case <-time.After(250 * time.Millisecond):
					}

					// Turn off.
					// For the math to work we need to convert our ints to floats and then back again.
					Red = ((float64(alight.OffColor.R) / 2) / 100) * (float64(alight.Brightness) / 2.55)
					Green = ((float64(alight.OffColor.G) / 2) / 100) * (float64(alight.Brightness) / 2.55)
					Blue = ((float64(alight.OffColor.B) / 2) / 100) * (float64(alight.Brightness) / 2.55)
					err = pad.Light(alight.X, alight.Y, int(Red), int(Green), int(Blue))
					if err != nil {
						fmt.Printf("flash off: error writing to launchpad %e\n" + err.Error())
					}

					// We wait for a stop message or 250ms which ever comes first.
					select {
					case <-LaunchPadFlashButtons[alight.X][alight.Y+1].FlashStopChannel:
						return
					case <-time.After(250 * time.Millisecond):
					}
				}
			}()
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
