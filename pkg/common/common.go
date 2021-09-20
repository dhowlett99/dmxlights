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
}

type Color struct {
	R int
	G int
	B int
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
	UnHide          bool
	Hide            bool
	Name            string
	Number          int
	Start           bool
	Stop            bool
	ReadConfig      bool
	LoadConfig      bool
	UpdateSpeed     bool
	Speed           int
	UpdatePatten    bool
	Patten          Patten
	IncreaseFade    bool
	DecreaseFade    bool
	FadeTime        time.Duration
	FadeSpeed       int
	UpdateSize      bool
	Size            int
	X               int
	Y               int
	Blackout        bool
	Normal          bool
	MusicTrigger    bool
	MusicTriggerOff bool
	SoftFadeOn      bool
	SoftFadeOff     bool
	UpdateColor     bool
	Color           int
}

// Sequence describes sequences.
type Sequence struct {
	Hide         bool
	Type         string
	FadeTime     time.Duration
	FadeOnTime   time.Duration
	FadeOffTime  time.Duration
	SoftFade     bool
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

// LightOn Turn on a common.Light.
func LightOn(eventsForLauchpad chan ALight, Light ALight) {
	event := ALight{X: Light.X, Y: Light.Y, Brightness: 255, Red: Light.Red, Green: Light.Green, Blue: Light.Blue}
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
	LightOn(eventsForLauchpad, ALight{X: 8, Y: selectedSequence - 1, Brightness: 255, Red: 0, Green: 0, Blue: 255})
}

func SendCommandToSequence(selectedSequence int, command Command, commandChannels []chan Command) {
	commandChannels[selectedSequence-1] <- command
}

func SendCommandToAllSequence(selectedSequence int, command Command, commandChannels []chan Command) {
	commandChannels[0] <- command
	commandChannels[1] <- command
	commandChannels[2] <- command
	commandChannels[3] <- command
}

func SendCommandToAllSequenceExcept(selectedSequence int, command Command, commandChannels []chan Command) {
	for index := range commandChannels {
		if index != selectedSequence-1 {
			commandChannels[index] <- command
		}
	}
}

func MakeFunctionButtons(selectedSequence int, eventsForLauchpad chan ALight, commandChannels []chan Command, functionButtons [][]bool, X int, Y int) {
	if !functionButtons[X][Y] {
		functionButtons[X][Y] = true
		HideFunctionButtons(eventsForLauchpad, functionButtons)
		ShowFunctionButtons(selectedSequence, eventsForLauchpad, functionButtons)
	} else {
		HideFunctionButtons(eventsForLauchpad, functionButtons)
		functionButtons[X][Y] = false
		// unhide the sequence
		cmd := Command{
			UnHide: true,
		}
		SendCommandToAllSequence(selectedSequence, cmd, commandChannels)
		// Wait for the sequence to end
		time.Sleep(700 * time.Millisecond)
	}
}
func ShowFunctionButtons(selectedSequence int, eventsForLauchpad chan ALight, functionButtons [][]bool) {
	for x := 0; x < 8; x++ {
		LightOn(eventsForLauchpad, ALight{X: x, Y: selectedSequence - 1, Brightness: 255, Red: 3, Green: 255, Blue: 255})
	}
}

func HideFunctionButtons(eventsForLauchpad chan ALight, functionButtons [][]bool) {
	for x := 0; x < 8; x++ {
		for y := 0; y < 4; y++ {
			LightOn(eventsForLauchpad, ALight{X: x, Y: y, Brightness: 0, Red: 0, Green: 0, Blue: 0})
		}
	}
}

func HideFunctionButtonsExcept(selectedSequence int, eventsForLauchpad chan ALight, functionButtons [][]bool) {
	for x := 0; x < 8; x++ {
		for y := 0; y < 4; y++ {
			if y != selectedSequence-1 {
				LightOn(eventsForLauchpad, ALight{X: x, Y: y, Brightness: 0, Red: 0, Green: 0, Blue: 0})
			}
		}
	}
}

func HideSequence(sequences []*Sequence, selectedSequence int, commandChannels []chan Command, sendMesg bool) {
	cmd := Command{}
	// Set the selected flag inside the sequence.
	for _, sequence := range sequences {
		//fmt.Printf("seq no %d   selected %d \n", sequence.Number, sequenceNumber)
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
