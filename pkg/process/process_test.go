package process

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/dhowlett99/dmxlights/pkg/common"
)

func Test_processDifferentColor(t *testing.T) {

	fadeColors := make(map[int][]common.FixtureBuffer, 10)
	type args struct {
		start         bool
		end           bool
		bounce        bool
		invert        bool
		stepNumber    int
		fadeColors    map[int][]common.FixtureBuffer
		fixture       common.Fixture
		fixtureNumber int
		color         common.Color
		colorNumber   int
		lastStep      common.Step
		nextStep      common.Step
		sequence      common.Sequence
		shift         int
	}
	tests := []struct {
		name string
		args args
		want map[int][]common.FixtureBuffer
	}{
		{
			// If color is different from last color and not black.
			// And last color was green, so fade down green first.
			name: "process a single color red, last color was black",
			args: args{
				start:      false,
				end:        false,
				fadeColors: fadeColors,
				sequence: common.Sequence{
					FadeUp:   []int{0, 50, 255},
					FadeOn:   []int{255},
					FadeDown: []int{255, 50, 0},
				},
				fixture: common.Fixture{
					Color: common.Color{
						R: 255,
						G: 0,
						B: 0,
					},
				},
				colorNumber: 0,
				color: common.Color{
					R: 255,
					G: 0,
					B: 0,
				},
				lastStep: common.Step{
					Fixtures: map[int]common.Fixture{
						0: {
							Color: common.Color{
								R: 0,
								G: 255,
								B: 0,
							},
						},
					},
				},
				nextStep: common.Step{
					Fixtures: map[int]common.Fixture{
						0: {
							Color: common.Color{
								R: 0,
								G: 0,
								B: 0,
							},
						},
					},
				},
			},

			want: map[int][]common.FixtureBuffer{

				0: {
					// Fade Down Green.
					{Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 0, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 0, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 0, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					// Fade Up Red.
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 0, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 0, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 0, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					// Keep Red on for on time.
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 0, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lastFixture := tt.args.lastStep.Fixtures[tt.args.fixtureNumber]
			nextFixture := tt.args.nextStep.Fixtures[tt.args.fixtureNumber]
			if fadeColors := ProcessRGBColor(tt.args.stepNumber, tt.args.start, tt.args.end, tt.args.bounce, tt.args.invert, tt.args.fadeColors, &tt.args.fixture, &lastFixture, &nextFixture, tt.args.sequence, tt.args.shift); !reflect.DeepEqual(fadeColors, tt.want) {
				t.Errorf("processColor() got = %v, want %v", fadeColors, tt.want)

				fmt.Printf("++++++++++++++ GOT ++++++++++++++++++++\n")
				for fixtureNumber := 0; fixtureNumber < len(fadeColors); fixtureNumber++ {

					fade := fadeColors[fixtureNumber]

					fmt.Printf("fixtureNumber:%d ============================\n", fixtureNumber)

					step := 0
					for _, fixtureBuffer := range fade {
						fmt.Printf("\tstep %d rule %d %s: fixtureBuffer:%+v\n", fixtureBuffer.Step, fixtureBuffer.Rule, fixtureBuffer.DebugMsg, fixtureBuffer.Color)
						step++
					}

				}
			}
		})
	}
}

func Test_processSameColorNotBlack(t *testing.T) {

	fadeColors := make(map[int][]common.FixtureBuffer, 10)
	type args struct {
		stepNumber    int
		start         bool
		end           bool
		bounce        bool
		invert        bool
		fadeColors    map[int][]common.FixtureBuffer
		fixture       common.Fixture
		fixtureNumber int
		color         common.Color
		colorNumber   int
		lastStep      common.Step
		nextStep      common.Step
		sequence      common.Sequence
		shift         int
	}
	tests := []struct {
		name string
		args args
		want map[int][]common.FixtureBuffer
	}{
		{
			// If color is same as last time , play that color out again.
			name: "process a single color red, last color was red",
			args: args{
				start:      false,
				end:        false,
				shift:      10, // inverted to represents shift 0
				fadeColors: fadeColors,
				sequence: common.Sequence{
					FadeUp:   []int{0, 50, 255},
					FadeOn:   []int{255},
					FadeDown: []int{255, 50, 0},
				},
				fixtureNumber: 0,
				// Fixture contains color Red.
				fixture: common.Fixture{
					Color: common.Color{
						R: 255,
						G: 0,
						B: 0,
					},
				},
				colorNumber: 0,
				// Color is therefor Red.
				color: common.Color{
					R: 255,
					G: 0,
					B: 0,
				},
				// Last step was also red.
				lastStep: common.Step{
					Fixtures: map[int]common.Fixture{
						0: {
							Color: common.Color{
								R: 255,
								G: 0,
								B: 0,
							},
						},
					},
				},
				nextStep: common.Step{
					Fixtures: map[int]common.Fixture{
						0: {
							Color: common.Color{
								R: 0,
								G: 0,
								B: 0,
							},
						},
					},
				},
			},

			want: map[int][]common.FixtureBuffer{

				0: {
					// Play out the existing Red
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 0, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 0, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 0, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					// Keep Red on for the on time.
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 0, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					// Keep Red on for the down time.
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 0, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 0, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 0, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lastFixture := tt.args.lastStep.Fixtures[tt.args.fixtureNumber]
			nextFixture := tt.args.nextStep.Fixtures[tt.args.fixtureNumber]
			if fadeColors := ProcessRGBColor(tt.args.stepNumber, tt.args.start, tt.args.end, tt.args.bounce, tt.args.invert, tt.args.fadeColors, &tt.args.fixture, &lastFixture, &nextFixture, tt.args.sequence, tt.args.shift); !reflect.DeepEqual(fadeColors, tt.want) {
				t.Errorf("processColor() got = %v, want %v", fadeColors, tt.want)

				fmt.Printf("++++++++++++++ GOT ++++++++++++++++++++\n")
				for fixtureNumber := 0; fixtureNumber < len(fadeColors); fixtureNumber++ {

					fade := fadeColors[fixtureNumber]

					fmt.Printf("fixtureNumber:%d ============================\n", fixtureNumber)

					step := 0
					for _, fixtureBuffer := range fade {
						fmt.Printf("\tstep %d rule %d %s: fixtureBuffer:%+v\n", fixtureBuffer.Step, fixtureBuffer.Rule, fixtureBuffer.DebugMsg, fixtureBuffer.Color)
						step++
					}

				}
			}
		})
	}
}

func Test_processDiffColorBlack(t *testing.T) {

	fadeColors := make(map[int][]common.FixtureBuffer, 10)
	type args struct {
		stepNumber    int
		start         bool
		end           bool
		bounce        bool
		invert        bool
		fadeColors    map[int][]common.FixtureBuffer
		fixture       common.Fixture
		fixtureNumber int
		color         common.Color
		colorNumber   int
		lastStep      common.Step
		nextStep      common.Step
		sequence      common.Sequence
		shift         int
	}
	tests := []struct {
		name string
		args args
		want map[int][]common.FixtureBuffer
	}{
		{
			// If color is different from last color and color is a black.
			name: "process a single color black, last color was red",
			args: args{
				start:      false,
				end:        false,
				shift:      10, // inverted to represents shift 0
				fadeColors: fadeColors,
				sequence: common.Sequence{
					//FadeOff:  []int{0},
					FadeUp:   []int{0, 50, 255},
					FadeOn:   []int{255},
					FadeDown: []int{255, 50, 0},
				},
				fixtureNumber: 0,
				// Fixture contains color Black.
				fixture: common.Fixture{
					Color: common.Color{
						R: 0,
						G: 0,
						B: 0,
					},
				},
				colorNumber: 0,
				// Color is therefor Black.
				color: common.Color{
					R: 0,
					G: 0,
					B: 0,
				},
				// Last step was also red.
				lastStep: common.Step{
					Fixtures: map[int]common.Fixture{
						0: {
							Color: common.Color{
								R: 255,
								G: 0,
								B: 0,
							},
						},
					},
				},
				nextStep: common.Step{
					Fixtures: map[int]common.Fixture{
						0: {
							Color: common.Color{
								R: 0,
								G: 0,
								B: 0,
							},
						},
					},
				},
			},

			want: map[int][]common.FixtureBuffer{

				0: {
					// Fade Down Red to Black.
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 0, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 0, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 0, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					// Keep off for the off time, same as on time.
					//{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 0, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					// Keep off for the fade up time.
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 0, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 0, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 0, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					// Keep off for the fade on time.
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 0, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					// Keep off for the fade down time.
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 0, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 0, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 0, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lastFixture := tt.args.lastStep.Fixtures[tt.args.fixtureNumber]
			nextFixture := tt.args.nextStep.Fixtures[tt.args.fixtureNumber]
			if fadeColors := ProcessRGBColor(tt.args.stepNumber, tt.args.start, tt.args.end, tt.args.bounce, tt.args.invert, tt.args.fadeColors, &tt.args.fixture, &lastFixture, &nextFixture, tt.args.sequence, tt.args.shift); !reflect.DeepEqual(fadeColors, tt.want) {
				t.Errorf("processColor() got = %v, want %v", fadeColors, tt.want)

				fmt.Printf("++++++++++++++ GOT ++++++++++++++++++++\n")
				for fixtureNumber := 0; fixtureNumber < len(fadeColors); fixtureNumber++ {

					fade := fadeColors[fixtureNumber]

					fmt.Printf("fixtureNumber:%d ============================\n", fixtureNumber)

					step := 0
					for _, fixtureBuffer := range fade {
						fmt.Printf("\tstep %d rule %d %s: fixtureBuffer:%+v\n", fixtureBuffer.Step, fixtureBuffer.Rule, fixtureBuffer.DebugMsg, fixtureBuffer.Color)
						step++
					}

				}
			}
		})
	}
}
