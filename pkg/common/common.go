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
}

// Fixture Command.
type FixtureCommand struct {
	Tick            bool
	Config          bool // Configure fixture.
	Start           bool
	Steps           int
	Positions       map[int]Position
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

func SequenceSelect(eventsForLauchpad chan ALight, sequenceNumber int) {
	// Turn off all sequence lights.
	for seq := 0; seq < 4; seq++ {
		LightOff(eventsForLauchpad, 8, seq)
	}
	// Now turn blue the selected seq light.
	LightOn(eventsForLauchpad, ALight{X: 8, Y: sequenceNumber - 1, Brightness: 255, Red: 0, Green: 0, Blue: 255})

}
