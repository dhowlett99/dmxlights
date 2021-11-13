package common

import (
	"time"
)

type ALight struct {
	X          int
	Y          int
	Brightness int
	Red        int
	Green      int
	Blue       int
	Flash      bool
	OnColor    int
	OffColor   int
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
	Name            string
	Number          int16
	CurrentPosition int
	Description     string
	States          []State
}

type StaticColorButton struct {
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
	Chase    []int
	Steps    []Step
}

// Command tells sequences what to do.
type Command struct {
	UpdateMode            bool
	Mode                  string
	MasterBrightness      bool
	Master                int
	Static                bool
	UpdateStatic          bool
	UpdateStaticColor     bool
	UpdateSequenceColor   bool
	ClearSequenceColor    bool
	StaticLampFlash       bool
	SelectedColor         int
	PlayStaticOnce        bool
	PlaySwitchOnce        bool
	SetEditColors         bool
	EditColors            bool
	StaticColor           Color
	StaticLamp            int
	UnHide                bool
	Hide                  bool
	Name                  string
	Number                int
	Start                 bool
	Stop                  bool
	ReadConfig            bool
	LoadConfig            bool
	UpdateSpeed           bool
	Speed                 int
	UpdatePatten          bool
	Patten                Patten
	IncreaseFade          bool
	DecreaseFade          bool
	FadeTime              time.Duration
	FadeSpeed             int
	UpdateSize            bool
	Size                  int
	UpdateSequenceSize    bool
	SequenceSize          int
	X                     int
	Y                     int
	Blackout              bool
	Normal                bool
	SoftFadeOn            bool
	SoftFadeOff           bool
	UpdateColor           bool
	Color                 int
	UpdateFunctionMode    bool
	FunctionMode          bool
	UpdateFunctions       bool
	Functions             []Function
	UpdateSequence        bool
	UpdateSwitch          bool
	SwitchNumber          int
	SwitchPosition        int
	UpdateSwitchPositions bool
	Inverted              bool
	UpdateGobo            bool
	SelectedGobo          int
	UpdateFlood           bool
	Flood                 bool
	UpdateAutoColor       bool
	AutoColor             bool
}

type Gobo struct {
	Name    string
	Number  int
	Setting int
}

// Sequence describes sequences.
type Sequence struct {
	NumberFixtures          int
	NumberScanners          int
	Mode                    string // Sequence or Static
	Static                  bool
	EditSeqColors           bool
	PlayStaticOnce          bool
	PlaySwitchOnce          bool
	Flood                   bool
	PlayFloodOnce           bool
	StaticColors            []StaticColorButton
	AvailableSequenceColors []StaticColorButton
	EditColors              bool
	Hide                    bool
	Type                    string
	FadeTime                time.Duration
	FadeOnTime              time.Duration
	FadeOffTime             time.Duration
	Name                    string
	Number                  int
	Run                     bool
	Bounce                  bool
	Steps                   int    // Holds the number of steps this sequence has. Will change if you change size, fade times etc.
	Patten                  Patten // Contains fixtures and steps info.
	Colors                  []Color
	Shift                   bool // Used for shifting scanners patterns apart.
	CurrentSpeed            time.Duration
	Speed                   int
	FadeSpeed               int
	Size                    int
	SequenceSize            int
	X                       int
	Y                       int
	MusicTrigger            bool
	Blackout                bool
	Color                   int
	Gobo                    []Gobo
	SelectedGobo            int
	SelectedColor           int
	Master                  int // Master Brightness
	Functions               []Function
	FunctionMode            bool
	Switches                []Switch
	UpdateSequenceColor     bool
	SequenceColors          []Color
	Inverted                bool
	Positions               map[int][]Position
	CurrentSequenceColors   []Color
	SelectedFloodSequence   map[int]bool // A map that remembers who is in flood mode.
	AutoColor               bool
	AutoPatten              bool
	SelectedPatten          int
	RecoverSequenceColors   bool
	SaveColors              bool
}

type Function struct {
	Name           string
	SequenceNumber int
	Number         int
	State          bool
	Flash          bool
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
	Master              int
	Hide                bool
	Tick                bool
	Config              bool // Configure fixture.
	Start               bool
	Steps               int
	Positions           map[int][]Position
	Type                string
	StartPosition       int
	CurrentPosition     int
	CurrentSpeed        time.Duration
	Color               Color
	Speed               int
	Shift               int
	Size                int
	FadeSpeed           int
	FadeTime            time.Duration
	FadeUpTime          time.Duration
	FadeOnTime          time.Duration
	FadeDownTime        time.Duration
	FadeOffTime         time.Duration
	Blackout            bool
	Flood               bool
	PlayFloodOnce       bool
	UpdateSequenceColor bool
	SequenceColor       Color
	Inverted            bool
	SelectedGobo        int
}

type Position struct {
	Fixture       int
	StartPosition int
	Color         Color
	Pan           int
	Tilt          int
	Shutter       int
	Gobo          int
}

// A fixture can have any or some of the
// following, depending if its a light or
// a scanner.
type Fixture struct {
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
}

// Define the function keys.
const (
	Function1_Forward_Chase = 0 // Forward chase in RGB, cirlce in scanner mode.
	Function2_Pairs_Chase   = 1 // Pais chase in RGB, left and right scanner mode.
	Function3_Inward_Chase  = 2 // Inward chase in RGB, up and down scanner mode.
	Function4_Bounce        = 3 // Sequence auto reverses. or Sinewave pattern in scanner mode.
	Function5_Color         = 4 // Set chase color. or select the scanner GOBO or color.
	Function6_Static        = 5 // Set static colors.
	Function7_Invert        = 6 // Invert the RGB colors or shift scanners 360 deg of each other.
	Function8_Music_Trigger = 7 // Music trigger on and off. Both RGB and scanners.
)

// LightOn Turn on a common.Light.
func LightOn(eventsForLauchpad chan ALight, Light ALight) {
	event := ALight{
		X:          Light.X,
		Y:          Light.Y,
		Brightness: Light.Brightness,
		Red:        Light.Red,
		Green:      Light.Green,
		Blue:       Light.Blue,
		Flash:      Light.Flash,
		OnColor:    22,
		OffColor:   18,
	}
	eventsForLauchpad <- event
}

// LightOff Turn on a common.Light.
func LightOff(eventsForLauchpad chan ALight, X int, Y int) {
	event := ALight{X: X, Y: Y, Brightness: 0, Red: 0, Green: 0, Blue: 0}
	eventsForLauchpad <- event
}

func SequenceSelect(eventsForLauchpad chan ALight, selectedSequence int) {
	// Turn off all sequence lights.
	for seq := 0; seq < 4; seq++ {
		//LightOff(eventsForLauchpad, 8, seq)
		LightOn(eventsForLauchpad, ALight{X: 8, Y: seq, Brightness: 255, Red: 255, Green: 255, Blue: 255})
	}
	// Now turn blue the selected sequence select light.
	LightOn(eventsForLauchpad, ALight{X: 8, Y: selectedSequence, Brightness: 255, Red: 255, Green: 0, Blue: 255})
}

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

func MakeFunctionButtons(sequence Sequence, selectedSequence int, eventsForLauchpad chan ALight, functionButtons [][]bool, channels Channels) {
	HideFunctionButtons(selectedSequence, eventsForLauchpad)
	// Get an upto date copy of the sequence.
	cmd := Command{
		ReadConfig: true,
	}
	SendCommandToSequence(selectedSequence, cmd, channels.CommmandChannels)

	replyChannel := channels.ReplyChannels[selectedSequence]
	sequence = <-replyChannel

	ShowFunctionButtons(sequence, selectedSequence, eventsForLauchpad)
}

func HideFunctionButtons(selectedSequence int, eventsForLauchpad chan ALight) {
	for x := 0; x < 8; x++ {
		LightOn(eventsForLauchpad, ALight{X: x, Y: selectedSequence, Brightness: 0, Red: 0, Green: 0, Blue: 0})
	}
}

func ShowFunctionButtons(sequence Sequence, selectedSequence int, eventsForLauchpad chan ALight) {
	for index, function := range sequence.Functions {

		if function.State {
			LightOn(eventsForLauchpad, ALight{X: index, Y: selectedSequence, Brightness: 255, Red: 200, Green: 0, Blue: 255})
		} else {
			LightOn(eventsForLauchpad, ALight{X: index, Y: selectedSequence, Brightness: 255, Red: 3, Green: 255, Blue: 255})
		}
	}
}

func SetMode(selectedSequence int, commandChannels []chan Command, mode string) {
	cmd := Command{
		UpdateMode: true,
		Mode:       mode,
	}
	SendCommandToSequence(selectedSequence, cmd, commandChannels)
}

func RevealSequence(selectedSequence int, commandChannels []chan Command) {
	cmd := Command{
		UnHide: true,
	}
	SendCommandToSequence(selectedSequence, cmd, commandChannels)
}

func HideSequence(selectedSequence int, commandChannels []chan Command) {
	cmd := Command{
		Hide: true,
	}
	SendCommandToSequence(selectedSequence, cmd, commandChannels)
}

// Colors are selected from a pallete of 8 colors, this function takes 1-8 and
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
	}
	return Color{}
}

func GetLaunchPadColorCodeByRGB(color Color) (code byte) {

	switch color {
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

func SetFunctionKeys(functions []Function, sequence Sequence) Sequence {

	// Set normal chase.
	if sequence.Functions[Function1_Forward_Chase].State {
		functions[Function2_Pairs_Chase].State = false
		functions[Function3_Inward_Chase].State = false
		if sequence.Type == "scanner" {
			sequence.Functions[Function4_Bounce].State = false
		}
		return sequence
	}

	// Set pairs chase.
	if sequence.Functions[Function2_Pairs_Chase].State {
		sequence.Functions[Function1_Forward_Chase].State = false
		sequence.Functions[Function3_Inward_Chase].State = false
		if sequence.Type == "scanner" {
			sequence.Functions[Function4_Bounce].State = false
		}
		return sequence
	}

	// Set inward chase.
	if sequence.Functions[Function3_Inward_Chase].State {
		sequence.Functions[Function1_Forward_Chase].State = false
		sequence.Functions[Function2_Pairs_Chase].State = false
		if sequence.Type == "scanner" {
			sequence.Functions[Function4_Bounce].State = false
		}
		return sequence
	}

	if sequence.Type == "scanner" {
		if sequence.Functions[Function4_Bounce].State {
			sequence.Functions[Function1_Forward_Chase].State = false
			sequence.Functions[Function2_Pairs_Chase].State = false
			sequence.Functions[Function3_Inward_Chase].State = false
			return sequence
		}
	}

	return sequence
}

func SetFunctionKeyActions(functions []Function, sequence Sequence) Sequence {

	if sequence.Type != "scanner" {
		if sequence.Functions[Function1_Forward_Chase].State {
			sequence.Patten.Name = "standard"
		}

		if sequence.Functions[Function2_Pairs_Chase].State {
			sequence.Patten.Name = "pairs"
		}

		if sequence.Functions[Function3_Inward_Chase].State {
			sequence.Patten.Name = "inward"
		}
	}

	if sequence.Type == "scanner" {
		if sequence.Functions[Function1_Forward_Chase].State {
			sequence.Patten.Name = "circle"
		}

		if sequence.Functions[Function2_Pairs_Chase].State {
			sequence.Patten.Name = "leftandright"
		}

		if sequence.Functions[Function3_Inward_Chase].State {
			sequence.Patten.Name = "upanddown"
		}

		if sequence.Functions[Function4_Bounce].State {
			sequence.Patten.Name = "sinewave"
		}

		if sequence.Functions[Function7_Invert].State {
			sequence.Shift = true
		} else {
			sequence.Shift = false
		}
	}

	// Map bounce function to sequence bounce setting.
	if sequence.Type != "scanner" {
		sequence.Bounce = sequence.Functions[Function4_Bounce].State
	}

	// Map color selection function.
	if sequence.Functions[Function5_Color].State {
		sequence.PlayStaticOnce = true
		sequence.EditSeqColors = true
		sequence.Run = false
	}

	// Map static function.
	sequence.Static = sequence.Functions[Function6_Static].State

	// Map invert function.
	sequence.Inverted = sequence.Functions[Function7_Invert].State

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

func RefreshSequence(selectedSequence int, commandChannels []chan Command, updateChannels []chan Sequence) *Sequence {

	// Get an upto date copy of the sequence.
	cmd := Command{
		UpdateSequence: true,
	}
	SendCommandToSequence(selectedSequence, cmd, commandChannels)
	newSequence := <-updateChannels[selectedSequence]
	return &newSequence
}
