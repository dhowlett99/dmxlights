package patten

import (
	"reflect"
	"testing"

	"github.com/dhowlett99/dmxlights/pkg/common"
)

func Test_circleGenerator(t *testing.T) {
	type args struct {
		size int
		posX float64
		posY float64
	}
	tests := []struct {
		name    string
		args    args
		wantOut []common.Coordinate
	}{
		{
			name: "standard circle",
			args: args{
				size: 126,
				posX: 128,
				posY: 128,
			},
			wantOut: []common.Coordinate{
				{Tilt: 128, Pan: 254},
				{Tilt: 149, Pan: 252},
				{Tilt: 171, Pan: 246},
				{Tilt: 191, Pan: 237},
				{Tilt: 208, Pan: 224},
				{Tilt: 224, Pan: 208},
				{Tilt: 237, Pan: 191},
				{Tilt: 246, Pan: 171},
				{Tilt: 252, Pan: 149},
				{Tilt: 254, Pan: 128},
				{Tilt: 252, Pan: 106},
				{Tilt: 246, Pan: 84},
				{Tilt: 237, Pan: 65},
				{Tilt: 224, Pan: 47},
				{Tilt: 208, Pan: 31},
				{Tilt: 191, Pan: 18},
				{Tilt: 171, Pan: 9},
				{Tilt: 149, Pan: 3},
				{Tilt: 128, Pan: 2},
				{Tilt: 106, Pan: 3},
				{Tilt: 84, Pan: 9},
				{Tilt: 65, Pan: 18},
				{Tilt: 47, Pan: 31},
				{Tilt: 31, Pan: 47},
				{Tilt: 18, Pan: 65},
				{Tilt: 9, Pan: 84},
				{Tilt: 3, Pan: 106},
				{Tilt: 2, Pan: 127},
				{Tilt: 3, Pan: 149},
				{Tilt: 9, Pan: 171},
				{Tilt: 18, Pan: 191},
				{Tilt: 31, Pan: 208},
				{Tilt: 47, Pan: 224},
				{Tilt: 64, Pan: 237},
				{Tilt: 84, Pan: 246},
				{Tilt: 106, Pan: 252},
				//{127, 254},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := CircleGenerator(tt.args.size, len(tt.wantOut), tt.args.posX, tt.args.posY); !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("circleGenerator() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func Test_generateSteps(t *testing.T) {

	fixtureDisabled := make(map[int]bool, 8)
	for i := 0; i < 9; i++ {
		fixtureDisabled[i] = false
	}
	fixtureAvailable := make(map[int]bool, 8)
	for i := 0; i < 9; i++ {
		fixtureAvailable[i] = true
	}

	tests := []struct {
		name        string
		fixtures    int
		shift       int
		chase       bool
		coordinates []common.Coordinate
		want        []common.Step
	}{
		{
			name:     "circle patten - 8 point , no shift",
			fixtures: 1,
			shift:    0,
			coordinates: []common.Coordinate{
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
			want: []common.Step{
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
		{
			name:     "two fixtures, circle patten - 8 point , with shift of 1/4",
			fixtures: 2,
			shift:    1,
			coordinates: []common.Coordinate{
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
			want: []common.Step{
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
		{
			name:     "two fixtures, circle patten - 8 point , with shift of zero",
			fixtures: 2,
			shift:    0,
			coordinates: []common.Coordinate{
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
			want: []common.Step{
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
		{
			name:     "four fixtures, circle patten - 8 point , with shift of 1/4",
			fixtures: 4,
			shift:    1,
			coordinates: []common.Coordinate{
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
			want: []common.Step{
				{
					Type: "scanner",
					Fixtures: []common.Fixture{
						{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 0},
						{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}, Gobo: 36, Shutter: 255, Pan: 2, Tilt: 2},
						{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 36, Shutter: 255, Pan: 4, Tilt: 4},
						{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 6, Tilt: 6},
					},
				},
				{
					Type: "scanner",
					Fixtures: []common.Fixture{
						{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}, Gobo: 36, Shutter: 255, Pan: 1, Tilt: 1},
						{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Gobo: 36, Shutter: 255, Pan: 3, Tilt: 3},
						{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 5, Tilt: 5},
						{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 7, Tilt: 7},
					},
				},
				{
					Type: "scanner",
					Fixtures: []common.Fixture{
						{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}, Gobo: 36, Shutter: 255, Pan: 2, Tilt: 2},
						{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 36, Shutter: 255, Pan: 4, Tilt: 4},
						{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 6, Tilt: 6},
						{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 0},
					},
				},
				{
					Type: "scanner",
					Fixtures: []common.Fixture{
						{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Gobo: 36, Shutter: 255, Pan: 3, Tilt: 3},
						{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 5, Tilt: 5},
						{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 7, Tilt: 7},
						{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}, Gobo: 36, Shutter: 255, Pan: 1, Tilt: 1},
					},
				},
				{
					Type: "scanner",
					Fixtures: []common.Fixture{
						{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 36, Shutter: 255, Pan: 4, Tilt: 4},
						{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 6, Tilt: 6},
						{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 0},
						{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}, Gobo: 36, Shutter: 255, Pan: 2, Tilt: 2},
					},
				},
				{
					Type: "scanner",
					Fixtures: []common.Fixture{
						{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 5, Tilt: 5},
						{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 7, Tilt: 7},
						{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}, Gobo: 36, Shutter: 255, Pan: 1, Tilt: 1},
						{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Gobo: 36, Shutter: 255, Pan: 3, Tilt: 3},
					},
				},
				{
					Type: "scanner",
					Fixtures: []common.Fixture{
						{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 6, Tilt: 6},
						{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 0},
						{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}, Gobo: 36, Shutter: 255, Pan: 2, Tilt: 2},
						{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 36, Shutter: 255, Pan: 4, Tilt: 4},
					},
				},
				{
					Type: "scanner",
					Fixtures: []common.Fixture{
						{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 7, Tilt: 7},
						{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}, Gobo: 36, Shutter: 255, Pan: 1, Tilt: 1},
						{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Gobo: 36, Shutter: 255, Pan: 3, Tilt: 3},
						{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 5, Tilt: 5},
					},
				},
			},
		},
		{
			name:     "one fixture, circle patten - 8 point shift of 1/4 ",
			fixtures: 1,
			shift:    1,
			coordinates: []common.Coordinate{
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
			want: []common.Step{

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
		{
			name:     "two scanners doing the same circle",
			fixtures: 2,
			shift:    0,
			coordinates: []common.Coordinate{
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
			want: []common.Step{
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
		{
			name:     "four fixtures, circle patten - 8 point , with shift of 2 ie 1/2",
			fixtures: 4,
			shift:    2,
			coordinates: []common.Coordinate{
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
			want: []common.Step{
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
		{
			name:     "four fixtures, circle patten - 8 point , with shift of 3 ie 3/4 shift",
			fixtures: 4,
			shift:    3,
			coordinates: []common.Coordinate{
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
			want: []common.Step{
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sequence := common.Sequence{}
			sequence.ScannerCoordinates = tt.coordinates
			sequence.NumberScanners = tt.fixtures
			sequence.Shift = tt.shift
			sequence.ScannerChase = tt.chase
			sequence.FixtureDisabled = fixtureDisabled
			sequence.FixtureAvailable = fixtureAvailable
			if got := GenerateSteps(sequence); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Got = %v", got)
				t.Errorf("Want = %v", tt.want)
			}
		})
	}
}

func TestScanGenerateSineWave(t *testing.T) {
	type args struct {
		size        int
		frequency   int
		scannerSize int
	}
	tests := []struct {
		name    string
		args    args
		wantOut []common.Coordinate
	}{
		{
			name: "5000hz sawtooth",
			args: args{
				size:        255,
				frequency:   5000,
				scannerSize: 10,
			},
			wantOut: []common.Coordinate{
				{Tilt: 28, Pan: 1},
				{Tilt: 113, Pan: 11},
				{Tilt: 226, Pan: 21},
				{Tilt: 137, Pan: 31},
				{Tilt: 27, Pan: 41},
				{Tilt: 120, Pan: 51},
				{Tilt: 227, Pan: 61},
				{Tilt: 130, Pan: 71},
				{Tilt: 27, Pan: 81},
				{Tilt: 128, Pan: 91},
				{Tilt: 227, Pan: 101},
				{Tilt: 123, Pan: 111},
				{Tilt: 27, Pan: 121},
				{Tilt: 135, Pan: 131},
				{Tilt: 227, Pan: 141},
				{Tilt: 116, Pan: 151},
				{Tilt: 28, Pan: 161},
				{Tilt: 142, Pan: 171},
				{Tilt: 226, Pan: 181},
				{Tilt: 109, Pan: 191},
				{Tilt: 29, Pan: 201},
				{Tilt: 149, Pan: 211},
				{Tilt: 224, Pan: 221},
				{Tilt: 102, Pan: 231},
				{Tilt: 31, Pan: 241},
				{Tilt: 156, Pan: 251},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := ScanGenerateSineWave(tt.args.size, tt.args.frequency, tt.args.scannerSize); !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("ScanGenerateSineWave() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func Test_calulateShutterValue(t *testing.T) {
	type args struct {
		CurrentCoordinate int
		CurrentStep       int
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
				CurrentCoordinate: 0,
				CurrentStep:       0,
				NumberFixtures:    8,
				NumberCoordinates: 8,
			},
			want: 255,
		},
		{
			name: "golden path",
			args: args{
				CurrentCoordinate: 8,
				CurrentStep:       0,
				NumberFixtures:    8,
				NumberCoordinates: 8,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calulateShutterValue(tt.args.CurrentCoordinate, tt.args.CurrentStep, tt.args.NumberFixtures, tt.args.NumberCoordinates); got != tt.want {
				t.Errorf("calulateShutterValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
