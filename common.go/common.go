package common

import "time"

type Light struct {
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
	Start   bool
	Fixture int
	Step    int
}
