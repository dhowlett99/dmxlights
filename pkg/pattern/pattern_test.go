package pattern

import (
	"reflect"
	"testing"

	"github.com/dhowlett99/dmxlights/pkg/common"
)

func Test_circleGenerator(t *testing.T) {
	type args struct {
		radius            int
		numberCoordinates int
		posX              float64
		posY              float64
	}
	tests := []struct {
		name    string
		args    args
		wantOut []Coordinate
	}{
		{
			name: "standard circle",
			args: args{
				radius:            126,
				numberCoordinates: 36,
				posX:              128,
				posY:              128,
			},
			wantOut: []Coordinate{
				{128, 254},
				{149, 252},
				{171, 246},
				{191, 237},
				{208, 224},
				{224, 208},
				{237, 191},
				{246, 171},
				{252, 149},
				{254, 128},
				{252, 106},
				{246, 84},
				{237, 65},
				{224, 47},
				{208, 31},
				{191, 18},
				{171, 9},
				{149, 3},
				{128, 2},
				{106, 3},
				{84, 9},
				{65, 18},
				{47, 31},
				{31, 47},
				{18, 65},
				{9, 84},
				{3, 106},
				{2, 127},
				{3, 149},
				{9, 171},
				{18, 191},
				{31, 208},
				{47, 224},
				{64, 237},
				{84, 246},
				{106, 252},
				//{127, 254},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := CircleGenerator(tt.args.radius, tt.args.numberCoordinates, tt.args.posX, tt.args.posY); !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("circleGenerator() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func Test_generatePattern(t *testing.T) {

	allFixturesEnabled := map[int]common.ScannerState{
		0: {
			Enabled: true,
		},
		1: {
			Enabled: true,
		},
		2: {
			Enabled: true,
		},
		3: {
			Enabled: true,
		},
		4: {
			Enabled: true,
		},
		5: {
			Enabled: true,
		},
		6: {
			Enabled: true,
		},
		7: {
			Enabled: true,
		},
	}

	tests := []struct {
		name         string
		fixtures     int
		shift        int
		chase        bool
		scannerState map[int]common.ScannerState
		Coordinates  []Coordinate
		want         common.Pattern
	}{
		{
			name:         "circle pattern - 8 point , no shift",
			fixtures:     1,
			shift:        0,
			scannerState: allFixturesEnabled,
			Coordinates: []Coordinate{
				{
					Tilt: 0,
					Pan:  128,
				},
				{
					Tilt: 32,
					Pan:  192,
				},
				{
					Tilt: 128,
					Pan:  232,
				},
				{
					Tilt: 232,
					Pan:  192,
				},
				{
					Tilt: 255,
					Pan:  128,
				},
				{
					Tilt: 232,
					Pan:  64,
				},
				{
					Tilt: 128,
					Pan:  32,
				},
				{
					Tilt: 32,
					Pan:  64,
				},
			},
			want: common.Pattern{
				Steps: []common.Step{
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 128, Tilt: 0},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}, Gobo: 36, Shutter: 255, Pan: 192, Tilt: 32},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}, Gobo: 36, Shutter: 255, Pan: 232, Tilt: 128},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Gobo: 36, Shutter: 255, Pan: 192, Tilt: 232},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 36, Shutter: 255, Pan: 128, Tilt: 255},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 64, Tilt: 232},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 32, Tilt: 128},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 64, Tilt: 32},
						},
					},
				},
			},
		},
		{
			name:         "two fixtures, circle pattern - 8 point , with shift of 1/4",
			fixtures:     2,
			shift:        1,
			scannerState: allFixturesEnabled,
			Coordinates: []Coordinate{
				{
					Tilt: 0,
					Pan:  0,
				},
				{
					Tilt: 1,
					Pan:  1,
				},
				{
					Tilt: 2,
					Pan:  2,
				},
				{
					Tilt: 3,
					Pan:  3,
				},
				{
					Tilt: 4,
					Pan:  4,
				},
				{
					Tilt: 5,
					Pan:  5,
				},
				{
					Tilt: 6,
					Pan:  6,
				},
				{
					Tilt: 7,
					Pan:  7,
				},
			},
			want: common.Pattern{
				Steps: []common.Step{
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 0},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}, Gobo: 36, Shutter: 255, Pan: 2, Tilt: 2},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}, Gobo: 36, Shutter: 255, Pan: 1, Tilt: 1},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Gobo: 36, Shutter: 255, Pan: 3, Tilt: 3},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}, Gobo: 36, Shutter: 255, Pan: 2, Tilt: 2},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 36, Shutter: 255, Pan: 4, Tilt: 4},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Gobo: 36, Shutter: 255, Pan: 3, Tilt: 3},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 5, Tilt: 5},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 36, Shutter: 255, Pan: 4, Tilt: 4},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 6, Tilt: 6},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 5, Tilt: 5},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 7, Tilt: 7},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 6, Tilt: 6},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 0},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 7, Tilt: 7},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}, Gobo: 36, Shutter: 255, Pan: 1, Tilt: 1},
						},
					},
				},
			},
		},
		{
			name:         "two fixtures, circle pattern - 8 point , with shift of zero",
			fixtures:     2,
			shift:        0,
			scannerState: allFixturesEnabled,
			Coordinates: []Coordinate{
				{
					Tilt: 0,
					Pan:  0,
				},
				{
					Tilt: 1,
					Pan:  1,
				},
				{
					Tilt: 2,
					Pan:  2,
				},
				{
					Tilt: 3,
					Pan:  3,
				},
				{
					Tilt: 4,
					Pan:  4,
				},
				{
					Tilt: 5,
					Pan:  5,
				},
				{
					Tilt: 6,
					Pan:  6,
				},
				{
					Tilt: 7,
					Pan:  7,
				},
			},
			want: common.Pattern{
				Steps: []common.Step{
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 0},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 0},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}, Gobo: 36, Shutter: 255, Pan: 1, Tilt: 1},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}, Gobo: 36, Shutter: 255, Pan: 1, Tilt: 1},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}, Gobo: 36, Shutter: 255, Pan: 2, Tilt: 2},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}, Gobo: 36, Shutter: 255, Pan: 2, Tilt: 2},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Gobo: 36, Shutter: 255, Pan: 3, Tilt: 3},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Gobo: 36, Shutter: 255, Pan: 3, Tilt: 3},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 36, Shutter: 255, Pan: 4, Tilt: 4},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 36, Shutter: 255, Pan: 4, Tilt: 4},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 5, Tilt: 5},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 5, Tilt: 5},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 6, Tilt: 6},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 6, Tilt: 6},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 7, Tilt: 7},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 7, Tilt: 7},
						},
					},
				},
			},
		},
		{
			name:         "four fixtures, circle pattern - 8 point , with shift of 1/4 and chase turned on",
			fixtures:     4,
			shift:        1,
			scannerState: allFixturesEnabled,
			chase:        true,
			Coordinates: []Coordinate{
				{
					Tilt: 0,
					Pan:  0,
				},
				{
					Tilt: 1,
					Pan:  1,
				},
				{
					Tilt: 2,
					Pan:  2,
				},
				{
					Tilt: 3,
					Pan:  3,
				},
				{
					Tilt: 4,
					Pan:  4,
				},
				{
					Tilt: 5,
					Pan:  5,
				},
				{
					Tilt: 6,
					Pan:  6,
				},
				{
					Tilt: 7,
					Pan:  7,
				},
			},
			want: common.Pattern{
				Steps: []common.Step{
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 0},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}, Gobo: 36, Shutter: 0, Pan: 2, Tilt: 2},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 36, Shutter: 0, Pan: 4, Tilt: 4},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 36, Shutter: 0, Pan: 6, Tilt: 6},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}, Gobo: 36, Shutter: 255, Pan: 1, Tilt: 1},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Gobo: 36, Shutter: 0, Pan: 3, Tilt: 3},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 0, Pan: 5, Tilt: 5},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 36, Shutter: 0, Pan: 7, Tilt: 7},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}, Gobo: 36, Shutter: 0, Pan: 2, Tilt: 2},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 36, Shutter: 255, Pan: 4, Tilt: 4},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 36, Shutter: 0, Pan: 6, Tilt: 6},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 0, Pan: 0, Tilt: 0},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Gobo: 36, Shutter: 0, Pan: 3, Tilt: 3},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 5, Tilt: 5},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 36, Shutter: 0, Pan: 7, Tilt: 7},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}, Gobo: 36, Shutter: 0, Pan: 1, Tilt: 1},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 36, Shutter: 0, Pan: 4, Tilt: 4},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 36, Shutter: 0, Pan: 6, Tilt: 6},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 0},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}, Gobo: 36, Shutter: 0, Pan: 2, Tilt: 2},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 0, Pan: 5, Tilt: 5},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 36, Shutter: 0, Pan: 7, Tilt: 7},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}, Gobo: 36, Shutter: 255, Pan: 1, Tilt: 1},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Gobo: 36, Shutter: 0, Pan: 3, Tilt: 3},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 36, Shutter: 0, Pan: 6, Tilt: 6},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 0, Pan: 0, Tilt: 0},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}, Gobo: 36, Shutter: 0, Pan: 2, Tilt: 2},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 36, Shutter: 255, Pan: 4, Tilt: 4},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 36, Shutter: 0, Pan: 7, Tilt: 7},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}, Gobo: 36, Shutter: 0, Pan: 1, Tilt: 1},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Gobo: 36, Shutter: 0, Pan: 3, Tilt: 3},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 5, Tilt: 5},
						},
					},
				},
			},
		},
		{
			name:         "one fixture, circle pattern - 8 point shift of 1/4 ",
			fixtures:     1,
			shift:        1,
			scannerState: allFixturesEnabled,
			Coordinates: []Coordinate{
				{
					Tilt: 0,
					Pan:  0,
				},
				{
					Tilt: 1,
					Pan:  1,
				},
				{
					Tilt: 2,
					Pan:  2,
				},
				{
					Tilt: 3,
					Pan:  3,
				},
				{
					Tilt: 4,
					Pan:  4,
				},
				{
					Tilt: 5,
					Pan:  5,
				},
				{
					Tilt: 6,
					Pan:  6,
				},
				{
					Tilt: 7,
					Pan:  7,
				},
			},
			want: common.Pattern{
				Steps: []common.Step{
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 0},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}, Gobo: 36, Shutter: 255, Pan: 1, Tilt: 1},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}, Gobo: 36, Shutter: 255, Pan: 2, Tilt: 2},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Gobo: 36, Shutter: 255, Pan: 3, Tilt: 3},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 36, Shutter: 255, Pan: 4, Tilt: 4},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 5, Tilt: 5},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 6, Tilt: 6},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 7, Tilt: 7},
						},
					},
				},
			},
		},
		{
			name:         "two scanners doing the same circle",
			fixtures:     2,
			shift:        0,
			scannerState: allFixturesEnabled,
			Coordinates: []Coordinate{
				{
					Tilt: 0,
					Pan:  128,
				},
				{
					Tilt: 128,
					Pan:  255,
				},
				{
					Tilt: 128,
					Pan:  0,
				},
			},
			want: common.Pattern{
				Steps: []common.Step{
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 128, Tilt: 0},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 128, Tilt: 0},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}, Gobo: 36, Shutter: 255, Pan: 255, Tilt: 128},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}, Gobo: 36, Shutter: 255, Pan: 255, Tilt: 128},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 128},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 128},
						},
					},
				},
			},
		},
		{
			name:         "four fixtures, circle pattern - 8 point , with shift of 2 ie 1/2",
			fixtures:     4,
			shift:        2,
			scannerState: allFixturesEnabled,
			Coordinates: []Coordinate{
				{
					Tilt: 0,
					Pan:  0,
				},
				{
					Tilt: 1,
					Pan:  1,
				},
				{
					Tilt: 2,
					Pan:  2,
				},
				{
					Tilt: 3,
					Pan:  3,
				},
				{
					Tilt: 4,
					Pan:  4,
				},
				{
					Tilt: 5,
					Pan:  5,
				},
				{
					Tilt: 6,
					Pan:  6,
				},
				{
					Tilt: 7,
					Pan:  7,
				},
			},
			want: common.Pattern{
				Steps: []common.Step{
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 0},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 36, Shutter: 255, Pan: 4, Tilt: 4},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 0},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 36, Shutter: 255, Pan: 4, Tilt: 4},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}, Gobo: 36, Shutter: 255, Pan: 1, Tilt: 1},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 5, Tilt: 5},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}, Gobo: 36, Shutter: 255, Pan: 1, Tilt: 1},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 5, Tilt: 5},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}, Gobo: 36, Shutter: 255, Pan: 2, Tilt: 2},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 6, Tilt: 6},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}, Gobo: 36, Shutter: 255, Pan: 2, Tilt: 2},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 6, Tilt: 6},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Gobo: 36, Shutter: 255, Pan: 3, Tilt: 3},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 7, Tilt: 7},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Gobo: 36, Shutter: 255, Pan: 3, Tilt: 3},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 7, Tilt: 7},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 36, Shutter: 255, Pan: 4, Tilt: 4},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 0},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 36, Shutter: 255, Pan: 4, Tilt: 4},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 0},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 5, Tilt: 5},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}, Gobo: 36, Shutter: 255, Pan: 1, Tilt: 1},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 5, Tilt: 5},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}, Gobo: 36, Shutter: 255, Pan: 1, Tilt: 1},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 6, Tilt: 6},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}, Gobo: 36, Shutter: 255, Pan: 2, Tilt: 2},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 6, Tilt: 6},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}, Gobo: 36, Shutter: 255, Pan: 2, Tilt: 2},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 7, Tilt: 7},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Gobo: 36, Shutter: 255, Pan: 3, Tilt: 3},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 7, Tilt: 7},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Gobo: 36, Shutter: 255, Pan: 3, Tilt: 3},
						},
					},
				},
			},
		},
		{
			name:         "four fixtures, circle pattern - 8 point , with shift of 3 ie 3/4 shift",
			fixtures:     4,
			shift:        3,
			scannerState: allFixturesEnabled,
			Coordinates: []Coordinate{
				{
					Tilt: 0,
					Pan:  0,
				},
				{
					Tilt: 1,
					Pan:  1,
				},
				{
					Tilt: 2,
					Pan:  2,
				},
				{
					Tilt: 3,
					Pan:  3,
				},
				{
					Tilt: 4,
					Pan:  4,
				},
				{
					Tilt: 5,
					Pan:  5,
				},
				{
					Tilt: 6,
					Pan:  6,
				},
				{
					Tilt: 7,
					Pan:  7,
				},
			},
			want: common.Pattern{
				Steps: []common.Step{
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 0},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 6, Tilt: 6},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 36, Shutter: 255, Pan: 4, Tilt: 4},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}, Gobo: 36, Shutter: 255, Pan: 2, Tilt: 2},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}, Gobo: 36, Shutter: 255, Pan: 1, Tilt: 1},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 7, Tilt: 7},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 5, Tilt: 5},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Gobo: 36, Shutter: 255, Pan: 3, Tilt: 3},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}, Gobo: 36, Shutter: 255, Pan: 2, Tilt: 2},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 0},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 6, Tilt: 6},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 36, Shutter: 255, Pan: 4, Tilt: 4},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Gobo: 36, Shutter: 255, Pan: 3, Tilt: 3},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}, Gobo: 36, Shutter: 255, Pan: 1, Tilt: 1},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 7, Tilt: 7},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 5, Tilt: 5},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 36, Shutter: 255, Pan: 4, Tilt: 4},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}, Gobo: 36, Shutter: 255, Pan: 2, Tilt: 2},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 0},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 6, Tilt: 6},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 5, Tilt: 5},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Gobo: 36, Shutter: 255, Pan: 3, Tilt: 3},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}, Gobo: 36, Shutter: 255, Pan: 1, Tilt: 1},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 7, Tilt: 7},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 6, Tilt: 6},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 36, Shutter: 255, Pan: 4, Tilt: 4},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}, Gobo: 36, Shutter: 255, Pan: 2, Tilt: 2},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 0},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 7, Tilt: 7},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 5, Tilt: 5},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Gobo: 36, Shutter: 255, Pan: 3, Tilt: 3},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}, Gobo: 36, Shutter: 255, Pan: 1, Tilt: 1},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GeneratePattern(tt.Coordinates, tt.fixtures, tt.shift, tt.chase, tt.scannerState); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Got = %+v", got)
				t.Errorf("Want = %+v", tt.want)
			}
		})
	}
}

func TestScanGenerateSineWave(t *testing.T) {
	type args struct {
		size      float64
		frequency float64
	}
	tests := []struct {
		name    string
		args    args
		wantOut []Coordinate
	}{
		{
			name: "5000hz sawtooth",
			args: args{
				size:      128,
				frequency: 5000,
			},
			wantOut: []Coordinate{
				{28, 1},
				{113, 11},
				{226, 21},
				{137, 31},
				{27, 41},
				{120, 51},
				{227, 61},
				{130, 71},
				{27, 81},
				{128, 91},
				{227, 101},
				{123, 111},
				{27, 121},
				{135, 131},
				{227, 141},
				{116, 151},
				{28, 161},
				{142, 171},
				{226, 181},
				{109, 191},
				{29, 201},
				{149, 211},
				{224, 221},
				{102, 231},
				{31, 241},
				{156, 251},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := ScanGenerateSineWave(tt.args.size, tt.args.frequency, 10); !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("ScanGenerateSineWave() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func Test_calulateShutterValue(t *testing.T) {
	type args struct {
		currentCoordinate int
		currentStep       int
		bounce            bool
		NumberFixtures    int
		NumberCoordinates int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "golden path",
			args: args{
				bounce:            false,
				currentCoordinate: 0,
				currentStep:       0,
				NumberFixtures:    8,
				NumberCoordinates: 8,
			},
			want: 255,
		},
		{
			name: "golden path",
			args: args{
				currentCoordinate: 8,
				currentStep:       0,
				NumberFixtures:    8,
				NumberCoordinates: 8,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalulateShutterValue(tt.args.currentCoordinate, tt.args.currentStep, tt.args.NumberFixtures, tt.args.NumberCoordinates, tt.args.bounce); got != tt.want {
				t.Errorf("calulateShutterValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
