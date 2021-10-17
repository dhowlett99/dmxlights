package sequence

import (
	"reflect"
	"testing"

	"github.com/dhowlett99/dmxlights/pkg/common"
)

func Test_calculatePositions(t *testing.T) {

	full := 255
	type args struct {
		steps  []common.Step
		bounce bool
	}
	tests := []struct {
		name string
		args args
		want map[int][]common.Position
	}{
		{
			name: "golden path - common par fixture RGB",
			args: args{
				steps: []common.Step{
					{
						Fixtures: []common.Fixture{
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						},
					},
					{
						Fixtures: []common.Fixture{
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						},
					},
					{
						Fixtures: []common.Fixture{
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						},
					},
					{
						Fixtures: []common.Fixture{
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						},
					},
					{
						Fixtures: []common.Fixture{
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						},
					},
					{
						Fixtures: []common.Fixture{
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						},
					},
					{
						Fixtures: []common.Fixture{
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						},
					},
					{
						Fixtures: []common.Fixture{
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
						},
					},
					{
						Fixtures: []common.Fixture{
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						},
					},
				},
			},
			want: map[int][]common.Position{
				0: {
					{
						Fixture:       0,
						StartPosition: 0,
						Color:         common.Color{R: 0, G: 255, B: 0},
					},
				},
				14: {
					{
						Fixture:       1,
						StartPosition: 14,
						Color:         common.Color{R: 0, G: 255, B: 0},
					},
				},
				28: {
					{
						Fixture:       2,
						StartPosition: 28,
						Color:         common.Color{R: 0, G: 255, B: 0},
					},
				},
				42: {
					{
						Fixture:       3,
						StartPosition: 42,
						Color:         common.Color{R: 0, G: 255, B: 0},
					},
				},
				56: {
					{
						Fixture:       4,
						StartPosition: 56,
						Color:         common.Color{R: 0, G: 255, B: 0},
					},
				},
				70: {
					{
						Fixture:       5,
						StartPosition: 70,
						Color:         common.Color{R: 0, G: 255, B: 0},
					},
				},
				84: {
					{
						Fixture:       6,
						StartPosition: 84,
						Color:         common.Color{R: 0, G: 255, B: 0},
					},
				},

				98: {{
					Fixture:       7,
					StartPosition: 98,
					Color:         common.Color{R: 0, G: 255, B: 0},
				},
				},
			},
		},
		{
			name: "Scanner case",
			args: args{
				bounce: false,
				steps: []common.Step{
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 255},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 255},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 50, Tilt: 200},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 50, Tilt: 200},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 100, Tilt: 150},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 100, Tilt: 150},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 36, Shutter: 255, Pan: 150, Tilt: 100},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 36, Shutter: 255, Pan: 150, Tilt: 100},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 200, Tilt: 50},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 200, Tilt: 50},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 190, B: 255}}, Gobo: 36, Shutter: 255, Pan: 255, Tilt: 0},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 190, B: 255}}, Gobo: 36, Shutter: 255, Pan: 255, Tilt: 0},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 145, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 0},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 145, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 0},
						},
					},
				},
			},
			want: map[int][]common.Position{
				0: {
					{
						Fixture:       0,
						StartPosition: 0,
						Color:         common.Color{R: 0, G: 0, B: 255},
						Gobo:          36, Shutter: 255, Pan: 0, Tilt: 255,
					},
					{
						Fixture:       1,
						StartPosition: 0,
						Color:         common.Color{R: 0, G: 0, B: 255},
						Gobo:          36, Shutter: 255, Pan: 0, Tilt: 255,
					},
				},
				14: {
					{
						Fixture:       0,
						StartPosition: 14,
						Color:         common.Color{R: 255, G: 0, B: 0},
						Gobo:          36, Shutter: 255, Pan: 50, Tilt: 200,
					},
					{
						Fixture:       1,
						StartPosition: 14,
						Color:         common.Color{R: 255, G: 0, B: 0},
						Gobo:          36, Shutter: 255, Pan: 50, Tilt: 200,
					},
				},
				28: {
					{
						Fixture:       0,
						StartPosition: 28,
						Color:         common.Color{R: 255, G: 0, B: 255},
						Gobo:          36,
						Shutter:       255, Pan: 100, Tilt: 150,
					},
					{
						Fixture:       1,
						StartPosition: 28,
						Color:         common.Color{R: 255, G: 0, B: 255},
						Gobo:          36,
						Shutter:       255, Pan: 100, Tilt: 150,
					},
				},
				42: {
					{
						Fixture:       0,
						StartPosition: 42,
						Color:         common.Color{R: 0, G: 255, B: 255},
						Gobo:          36,
						Shutter:       255, Pan: 150, Tilt: 100,
					},
					{
						Fixture:       1,
						StartPosition: 42,
						Color:         common.Color{R: 0, G: 255, B: 255},
						Gobo:          36,
						Shutter:       255, Pan: 150, Tilt: 100,
					},
				},
				56: {
					{
						Fixture:       0,
						StartPosition: 56,
						Color:         common.Color{R: 100, G: 0, B: 255},
						Gobo:          36, Shutter: 255, Pan: 200, Tilt: 50,
					},
					{
						Fixture:       1,
						StartPosition: 56,
						Color:         common.Color{R: 100, G: 0, B: 255},
						Gobo:          36, Shutter: 255, Pan: 200, Tilt: 50,
					},
				},
				70: {
					{
						Fixture:       0,
						StartPosition: 70,
						Color:         common.Color{R: 0, G: 190, B: 255},
						Gobo:          36, Shutter: 255, Pan: 255, Tilt: 0,
					},
					{
						Fixture:       1,
						StartPosition: 70,
						Color:         common.Color{R: 0, G: 190, B: 255},
						Gobo:          36, Shutter: 255, Pan: 255, Tilt: 0,
					},
				},
				84: {
					{
						Fixture:       0,
						StartPosition: 84,
						Color:         common.Color{R: 145, G: 0, B: 255},
						Gobo:          36, Shutter: 255, Pan: 0, Tilt: 0,
					},
					{
						Fixture:       1,
						StartPosition: 84,
						Color:         common.Color{R: 145, G: 0, B: 255},
						Gobo:          36, Shutter: 255, Pan: 0, Tilt: 0,
					},
				},
			},
		},
		{
			name: "Pairs case",
			args: args{
				bounce: false,
				steps: []common.Step{
					{
						Fixtures: []common.Fixture{
							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						},
					},
					{
						Fixtures: []common.Fixture{
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						},
					},
				},
			},
			want: map[int][]common.Position{

				0: {
					{
						Fixture:       0,
						StartPosition: 0,
						Color:         common.Color{R: 255, G: 0, B: 0},
						Gobo:          0, Shutter: 0, Pan: 0, Tilt: 0,
					},
					{
						Fixture:       2,
						StartPosition: 0,
						Color:         common.Color{R: 255, G: 0, B: 0},
						Gobo:          0, Shutter: 0, Pan: 0, Tilt: 0,
					},
					{
						Fixture:       4,
						StartPosition: 0,
						Color:         common.Color{R: 255, G: 0, B: 0},
						Gobo:          0, Shutter: 0, Pan: 0, Tilt: 0,
					},
					{
						Fixture:       6,
						StartPosition: 0,
						Color:         common.Color{R: 255, G: 0, B: 0},
						Gobo:          0, Shutter: 0, Pan: 0, Tilt: 0,
					},
				},
				14: {
					{
						Fixture:       1,
						StartPosition: 14,
						Color:         common.Color{R: 255, G: 0, B: 0},
						Gobo:          0, Shutter: 0, Pan: 0, Tilt: 0,
					},
					{
						Fixture:       3,
						StartPosition: 14,
						Color:         common.Color{R: 255, G: 0, B: 0},
						Gobo:          0, Shutter: 0, Pan: 0, Tilt: 0,
					},
					{
						Fixture:       5,
						StartPosition: 14,
						Color:         common.Color{R: 255, G: 0, B: 0},
						Gobo:          0, Shutter: 0, Pan: 0, Tilt: 0,
					},
					{
						Fixture:       7,
						StartPosition: 14,
						Color:         common.Color{R: 255, G: 0, B: 0},
						Gobo:          0, Shutter: 0, Pan: 0, Tilt: 0,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := calculatePositions(tt.args.steps, tt.args.bounce); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %v", got)
				t.Errorf("want =%v", tt.want)
			}
		})
	}
}
