package common

import (
	"fmt"
	"time"

	"github.com/rakyll/launchpad/mk3"
)

const debug = false

type ALight struct {
	X           int
	Y           int
	Brightness  int
	Red         int
	Green       int
	Blue        int
	Flash       bool
	OnColor     int
	OffColor    int
	UpdateLabel bool
	Label       string
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
	Values      []Value
	ButtonColor Color
}

type Switch struct {
	Name         string
	Number       int
	CurrentState int
	Description  string
	States       []State
}

type StaticColorButton struct {
	Number        int
	X             int
	Y             int
	Color         Color
	SelectedColor int
	Flash         bool
}

type Patten struct {
	Name     string
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
const SetEditColors = 7
const EditColors = 8
const UnHide = 9
const Hide = 10
const Start = 11
const Stop = 12
const ReadConfig = 13
const LoadConfig = 14
const UpdateSpeed = 15
const UpdatePatten = 16
const SelectPatten = 17
const IncreaseFade = 18
const DecreaseFade = 19
const UpdateSize = 20
const UpdateScannerSize = 21
const Blackout = 22
const Normal = 23
const SoftFadeOn = 24
const SoftFadeOff = 25
const UpdateColor = 26
const UpdateFunctionMode = 27
const FunctionMode = 28
const UpdateFunctions = 29
const GetUpdatedSequence = 30
const UpdateSwitch = 31
const UpdateSwitchPositions = 32
const Inverted = 33
const UpdateGobo = 34
const Flood = 35
const NoFlood = 36
const UpdateAutoColor = 37
const AutoColor = 38
const UpdateAutoPatten = 39
const AutoPatten = 40
const ToggleFixtureState = 41
const FixtureState = 42
const UpdateShift = 43
const UpdateScannerColor = 44
const ClearSequenceColor = 45
const Static = 46
const MasterBrightness = 47
const UpdateNumberCoordinates = 48
const UpdateOffsetPan = 49
const UpdateOffsetTilt = 50

const DefaultScannerSize = 120

type Gobo struct {
	Name    string
	Number  int
	Setting int
}

// Sequence describes sequences.
type Sequence struct {
	NumberFixtures               int
	NumberScanners               int
	Mode                         string // Sequence or Static
	Static                       bool
	EditSeqColors                bool
	PlayStaticOnce               bool
	PlaySwitchOnce               bool
	Flood                        bool
	NoFlood                      bool
	PlayFloodOnce                bool
	StaticColors                 []StaticColorButton
	AvailableSequenceColors      []StaticColorButton
	AvailableScannerColors       map[int][]StaticColorButton
	AvailableGoboSelectionColors []StaticColorButton
	AvailableFixtures            []StaticColorButton // Holds a set of red buttons, one for every available fixture.
	EditColors                   bool
	Hide                         bool
	Type                         string
	FadeTime                     time.Duration
	FadeOnTime                   time.Duration
	FadeOffTime                  time.Duration
	Name                         string
	Number                       int
	Run                          bool
	Bounce                       bool
	NumberSteps                  int    // Holds the number of steps this sequence has. Will change if you change size, fade times etc.
	Patten                       Patten // Contains fixtures and steps info.
	Colors                       []Color
	UpdateShift                  bool
	Shift                        int // Used for shifting scanners patterns apart.
	CurrentSpeed                 time.Duration
	Speed                        int
	FadeSpeed                    int
	Size                         int
	ScannerSize                  int
	X                            int
	Y                            int
	MusicTrigger                 bool
	Blackout                     bool
	Color                        int
	Gobo                         []Gobo
	SelectedGobo                 int
	SelectedColor                int
	Master                       int // Master Brightness
	Functions                    []Function
	FunctionMode                 bool
	Switches                     []Switch
	UpdateSequenceColor          bool
	SequenceColors               []Color
	Inverted                     bool
	Positions                    map[int][]Position
	CurrentSequenceColors        []Color
	SavedSequenceColors          []Color
	SelectedFloodSequence        map[int]bool // A map that remembers who is in flood mode.
	AutoColor                    bool
	AutoPatten                   bool
	RecoverSequenceColors        bool
	SaveColors                   bool
	SelectedScannerPatten        int
	FixtureDisabled              map[int]bool
	DisableOnce                  map[int]bool
	ScannerChase                 bool
	UpdateScannerColor           bool
	ScannerColor                 map[int]int // eight scanners per sequence, each can have their own color.
	NumberCoordinates            []int
	SelectedCoordinates          int
	Steps                        []Step
	UpdatePatten                 bool
	SelectPatten                 bool
	SelectedPatten               int
	OffsetPan                    int
	OffsetTilt                   int
	FunctionLabels               [8]string
	BottomButtons                [8]string
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
	FixtureDisabled        map[int]bool
	DisableOnce            map[int]bool
	ScannerChase           bool
	ScannerColor           map[int]int
	Static                 bool
	StaticColors           []StaticColorButton
	AvailableScannerColors map[int][]StaticColorButton
	OffsetPan              int
	OffsetTilt             int
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

func SendCommandToAllSequenceOfType(sequences []*Sequence, selectedSequence int, command Command, commandChannels []chan Command, Type string) {

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
	HideFunctionButtons(selectedSequence, eventsForLauchpad, guiButtons)
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

func GetLaunchPadColorCodeByInt(code int) (color Color) {
	switch code {
	case 0x4e: // Light Blue
		return Color{R: 0, G: 196, B: 255}
	case 0x78: // Red
		return Color{R: 255, G: 0, B: 0}
	case 0x48: // Red
		return Color{R: 255, G: 0, B: 0}
	case 0x60: // Orange
		return Color{R: 255, G: 111, B: 0}
	case 0x0c: // Light Yellow
		return Color{R: 100, G: 100, B: 0}
	case 0x0d: // Yellow
		return Color{R: 255, G: 255, B: 0}
	case 0x15: // Green
		return Color{R: 0, G: 255, B: 0}
	case 0x25: // Cyan
		return Color{R: 0, G: 255, B: 255}
	case 0x4f: // Blue
		return Color{R: 0, G: 0, B: 255}
	case 0x51: // Purple
		return Color{R: 100, G: 0, B: 255}
	case 0x34: // Pink
		return Color{R: 255, G: 0, B: 255}
	case 0x03: // White
		return Color{R: 255, G: 255, B: 255}
	case 0x00: // Black
		return Color{R: 0, G: 0, B: 0}
	}
	return Color{R: 0, G: 0, B: 0}
}

func GetLaunchPadColorCodeByRGB(color Color) (code byte) {
	switch color {
	case Color{R: 0, G: 196, B: 255}:
		return 0x4e // Light Blue
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
		sequence.EditSeqColors = true
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
	sequence.Inverted = sequence.Functions[Function7_Invert_Chase].State
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
	for myFixtureNumber := range sequence.AvailableSequenceColors {
		LightLamp(ALight{X: myFixtureNumber, Y: mySequenceNumber, Red: 0, Green: 0, Blue: 0, Brightness: sequence.Master}, eventsForLauchpad, guiButtons)
	}
}

func HideFunctionButtons(selectedSequence int, eventsForLauchpad chan ALight, guiButtons chan ALight) {
	for x := 0; x < 8; x++ {
		LightLamp(ALight{X: x, Y: selectedSequence, Brightness: 0, Red: 0, Green: 0, Blue: 0}, eventsForLauchpad, guiButtons)
		LabelButton(x, selectedSequence, "", guiButtons)
	}
}

func ShowFunctionButtons(sequence Sequence, selectedSequence int, eventsForLauchpad chan ALight, guiButtons chan ALight) {

	// Loop through the available functions for this sequence
	for index, function := range sequence.Functions {
		if debug {
			fmt.Printf("function %+v\n", function)
		}

		if function.State {
			LightLamp(ALight{X: index, Y: selectedSequence, Brightness: 255, Red: 200, Green: 0, Blue: 255}, eventsForLauchpad, guiButtons)
		} else {
			LightLamp(ALight{X: index, Y: selectedSequence, Brightness: 255, Red: 3, Green: 255, Blue: 255}, eventsForLauchpad, guiButtons)
		}
		LabelButton(index, selectedSequence, function.Label, guiButtons)
	}
}

func ShowBottomButtons(sequence Sequence, selectedSequence int, eventsForLauchpad chan ALight, guiButtons chan ALight) {

	// Loop through the available functions for this sequence
	for index, button := range sequence.BottomButtons {
		if debug {
			fmt.Printf("button %+v\n", button)
		}
		LightLamp(ALight{X: index, Y: selectedSequence, Brightness: 255, Red: 3, Green: 255, Blue: 255}, eventsForLauchpad, guiButtons)
		LabelButton(index, selectedSequence, button, guiButtons)
	}
}

func ClearAll(pad *mk3.Launchpad, presetsStore map[string]bool, eventsForLauchpad chan ALight, guiButtons chan ALight, sequences []chan Command) {
	pad.Reset()
	command := Command{
		Action: Stop,
	}

	for _, sequence := range sequences {
		sequence <- command
	}

	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			if presetsStore[fmt.Sprint(x)+","+fmt.Sprint(y)] {
				LightLamp(ALight{X: x, Y: y, Red: 255, Green: 0, Blue: 0, Brightness: 255}, eventsForLauchpad, guiButtons)
			}
		}
	}
}

// ListenAndSendToLaunchPad is the thread that listens for events to send to
// the launch pad.  It is thread safe and is the only thread talking to the
// launch pad. A channel is used to queue the events to be sent.
func ListenAndSendToLaunchPad(eventsForLauchpad chan ALight, pad *mk3.Launchpad) {

	for {
		event := <-eventsForLauchpad

		if event.Flash {
			pad.FlashLight(event.X, event.Y, event.OnColor, event.OffColor)
		} else {
			// For the math to work we need to convert our ints to floats and then back again.
			Red := ((float64(event.Red) / 2) / 100) * (float64(event.Brightness) / 2.55)
			Green := ((float64(event.Green) / 2) / 100) * (float64(event.Brightness) / 2.55)
			Blue := ((float64(event.Blue) / 2) / 100) * (float64(event.Brightness) / 2.55)
			pad.Light(event.X, event.Y, int(Red), int(Green), int(Blue))
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
		OnColor:    22,
		OffColor:   18,
	}
	eventsForLauchpad <- event

	// Send message to GUI
	event = ALight{
		X:          Light.X,
		Y:          Light.Y + 1,
		Brightness: Light.Brightness,
		Red:        Light.Red,
		Green:      Light.Green,
		Blue:       Light.Blue,
		Flash:      false,
		OnColor:    22,
		OffColor:   18,
		Label:      Light.Label,
	}
	guiButtons <- event

}

func FlashLight(X int, Y int, onColor int, offColor int, eventsForLauchpad chan ALight, guiButtons chan ALight) {

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
