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

type StaticColorButtons struct {
	X     int
	Y     int
	Color Color
}

type Patten struct {
	Name     string
	Length   int // 8, 4 or 2
	Size     int
	Fixtures int // 8 Fixtures
	Chase    []int
	Steps    []Step
}

// Command tells sequences what to do.
type Command struct {
	Static            bool
	UpdateStatic      bool
	UpdateStaticColor bool
	StaticColor       Color
	StaticLamp        int
	UnHide            bool
	Hide              bool
	Name              string
	Number            int
	Start             bool
	Stop              bool
	ReadConfig        bool
	LoadConfig        bool
	UpdateSpeed       bool
	Speed             int
	UpdatePatten      bool
	Patten            Patten
	IncreaseFade      bool
	DecreaseFade      bool
	FadeTime          time.Duration
	FadeSpeed         int
	UpdateSize        bool
	Size              int
	X                 int
	Y                 int
	Blackout          bool
	Normal            bool
	MusicTriggerOn    bool
	MusicTriggerOff   bool
	SoftFadeOn        bool
	SoftFadeOff       bool
	UpdateColor       bool
	Color             int
	UpdateFunctions   bool
	Functions         []Function
}

// Sequence describes sequences.
type Sequence struct {
	Static       bool
	StaticColors map[int]Color
	Hide         bool
	Type         string
	FadeTime     time.Duration
	FadeOnTime   time.Duration
	FadeOffTime  time.Duration
	Name         string
	Number       int
	Run          bool
	Bounce       bool
	Steps        int    // Holds the number of steps this sequence has. Will change if you change size, fade times etc.
	Patten       Patten // Contains fixtures and steps info.
	Colors       []Color
	Shift        int
	CurrentSpeed time.Duration
	Speed        int
	FadeSpeed    int
	Size         int
	X            int
	Y            int
	MusicTrigger bool
	Blackout     bool
	Color        int
	Master       int // Master Brightness
	Functions    []Function
}

type Function struct {
	Name           string
	SequenceNumber int
	Number         int
	State          bool
}

type Channels struct {
	CommmandChannels     []chan Command
	ReplyChannels        []chan Sequence
	SoundTriggerChannels []chan Command
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
	Hide            bool
	Tick            bool
	Config          bool // Configure fixture.
	Start           bool
	Steps           int
	Positions       map[int][]Position
	Type            string
	StartPosition   int
	CurrentPosition int
	CurrentSpeed    time.Duration
	Color           Color
	Speed           int
	Shift           int
	Size            int
	FadeSpeed       int
	FadeTime        time.Duration
	FadeUpTime      time.Duration
	FadeOnTime      time.Duration
	FadeDownTime    time.Duration
	FadeOffTime     time.Duration
	Blackout        bool
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

// Define the function keys.
const (
	Function1_Forward_Chase = 0
	Function2_Undef         = 1
	Function3_Undef         = 2
	Function4_Undef         = 3
	Function5_Undef         = 4
	Function6_Static        = 5 // Static Colors.
	Function7_Bounce        = 6
	Function8_Music_Trigger = 7
)

// LightOn Turn on a common.Light.
func LightOn(eventsForLauchpad chan ALight, Light ALight) {
	event := ALight{X: Light.X, Y: Light.Y, Brightness: Light.Brightness, Red: Light.Red, Green: Light.Green, Blue: Light.Blue}
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
		LightOff(eventsForLauchpad, 8, seq)
	}
	// Now turn blue the selected sequence select light.
	LightOn(eventsForLauchpad, ALight{X: 8, Y: selectedSequence, Brightness: 255, Red: 0, Green: 0, Blue: 255})
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

func SendCommandToAllSequenceExcept(selectedSequence int, command Command, commandChannels []chan Command) {
	for index := range commandChannels {
		if index != selectedSequence {
			commandChannels[index] <- command
		}
	}
}

func MakeFunctionButtons(sequence Sequence, selectedSequence int, eventsForLauchpad chan ALight, functionButtons [][]bool, channels Channels) {
	HideFunctionButtons(eventsForLauchpad, functionButtons)
	// Get an upto date copy of the sequence.
	cmd := Command{
		ReadConfig: true,
	}
	SendCommandToSequence(selectedSequence, cmd, channels.CommmandChannels)

	replyChannel := channels.ReplyChannels[selectedSequence]
	sequence = <-replyChannel

	ShowFunctionButtons(sequence, selectedSequence, eventsForLauchpad, functionButtons)
}

func HideFunctionButtons(eventsForLauchpad chan ALight, functionButtons [][]bool) {
	for x := 0; x < 8; x++ {
		for y := 0; y < 4; y++ {
			LightOn(eventsForLauchpad, ALight{X: x, Y: y, Brightness: 0, Red: 0, Green: 0, Blue: 0})
		}
	}
}

func ShowFunctionButtons(sequence Sequence, selectedSequence int, eventsForLauchpad chan ALight, functionButtons [][]bool) {

	for index, function := range sequence.Functions {
		if function.State {
			LightOn(eventsForLauchpad, ALight{X: index, Y: selectedSequence, Brightness: 255, Red: 200, Green: 0, Blue: 255})
		} else {
			LightOn(eventsForLauchpad, ALight{X: index, Y: selectedSequence, Brightness: 255, Red: 3, Green: 255, Blue: 255})
		}
	}
}

func RevealSequence(sequences []*Sequence, selectedSequence int, commandChannels []chan Command, sendMesg bool) {
	cmd := Command{
		UnHide: true,
	}
	SendCommandToSequence(selectedSequence, cmd, commandChannels)
}

func HideSequence(sequences []*Sequence, selectedSequence int, commandChannels []chan Command, sendMesg bool) {
	cmd := Command{}
	// Set the selected flag inside the sequence.
	for _, sequence := range sequences {
		if sequence.Number == selectedSequence {
			sequence.Hide = true
			cmd = Command{
				Hide: true,
			}
			if sendMesg {
				SendCommandToSequence(selectedSequence, cmd, commandChannels)
			}
		} else {
			cmd = Command{
				UnHide: true,
			}
			sequence.Hide = false
			if sendMesg {
				SendCommandToAllSequenceExcept(selectedSequence, cmd, commandChannels)
			}
		}
	}
}

func HandleSelect(sequences []*Sequence,
	selectedSequence int,
	X int,
	Y int,
	eventsForLauchpad chan ALight,
	selectButtons [][]bool,
	functionButtons [][]bool,
	functionMode [][]bool,
	commandChannels []chan Command,
	channels Channels) {

	// Light the sequence selector button.
	SequenceSelect(eventsForLauchpad, selectedSequence)

	if functionMode[X][Y] && !selectButtons[X][Y] {
		// Turn off function mode. Remove the function pads.
		HideFunctionButtons(eventsForLauchpad, functionButtons)
		// And reveal the sequence on the launchpad keys
		RevealSequence(sequences, selectedSequence, commandChannels, true)
		// Turn off the function mode flag.
		functionMode[X][Y] = false
		// Now forget we pressed twice and start again.
		selectButtons[X][Y] = true
		return
	}
	// This the first time we have pressed the select button.
	if !selectButtons[X][Y] {
		// assume everything else is off.
		selectButtons[X][0] = false
		selectButtons[X][1] = false
		selectButtons[X][2] = false
		selectButtons[X][3] = false
		// But remember we have pressed this select button once.
		functionMode[X][Y] = false
		selectButtons[X][Y] = true
		return
	}

	// Are we in function mode ?
	if functionMode[X][Y] {
		// Turn off function mode. Remove the function pads.
		HideFunctionButtons(eventsForLauchpad, functionButtons)
		// And reveal the sequence on the launchpad keys
		RevealSequence(sequences, selectedSequence, commandChannels, true)
		// Turn off the function mode flag.
		functionMode[X][Y] = false
		// Now forget we pressed twice and start again.
		selectButtons[X][Y] = false
		return
	}

	// We're not in function mode for this sequence.
	if !functionMode[X][Y] {
		// Set function mode.
		functionMode[X][Y] = true

		// And hide the sequence so we can only see the function buttons.
		HideSequence(sequences, selectedSequence, commandChannels, true)

		// Create the function buttons.
		MakeFunctionButtons(*sequences[selectedSequence], selectedSequence, eventsForLauchpad, functionButtons, channels)
		// Now forget we pressed twice and start again.
		selectButtons[X][Y] = false
		return
	}
}
