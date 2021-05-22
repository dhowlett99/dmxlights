package common

import "time"

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
	Steps    []Steps
}

type Sequence struct {
	// commands
	Start        bool
	Stop         bool
	ReadConfig   bool
	LoadConfig   bool
	UpdateSpeed  bool
	UpdatePatten bool
	UpdateFade   bool
	// parameters
	FadeTime     time.Duration
	Name         string
	Number       int
	Run          bool
	Patten       Patten
	Colors       []Color
	Speed        int
	CurrentSpeed time.Duration
	X            int
	Y            int
}

type Hit struct {
	X int
	Y int
}

type Steps struct {
	Fixtures []Fixture
}

type Fixture struct {
	Brightness int
	Colors     []Color
}

type ButtonPresets struct {
	X int
	Y int
}

type Event struct {
	Fixture int
	Run     bool
}

// LightOn Turn on a common.Light.
func LightOn(eventsForLauchpad chan ALight, Light ALight) {
	event := ALight{X: Light.X, Y: Light.Y, Brightness: 3, Red: Light.Red, Green: Light.Green, Blue: Light.Blue}
	eventsForLauchpad <- event
}

// LightOff Turn on a common.Light.
func LightOff(eventsForLauchpad chan ALight, X int, Y int) {
	event := ALight{X: X, Y: Y, Brightness: 0, Red: 0, Green: 0, Blue: 0}
	eventsForLauchpad <- event
}

// LightOn Turn on a common.Light.
func SequenceSelect(eventsForLauchpad chan ALight, sequenceNumber int) {

	// Turn off all sequence lights.
	for seq := 0; seq < 4; seq++ {
		LightOff(eventsForLauchpad, 8, seq)
	}
	// Now turn blue the selected seq light.
	LightOn(eventsForLauchpad, ALight{X: 8, Y: sequenceNumber - 1, Brightness: 3, Red: 0, Green: 0, Blue: 3})

}
