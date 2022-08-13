package sequence

import (
	"reflect"
	"testing"

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
)

func Test_calculatePositions(t *testing.T) {

	full := 255
	type args struct {
		steps  []common.Step
		bounce bool
		invert bool
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
			name: "multiple colors case",
			args: args{
				steps: []common.Step{
					{
						Fixtures: []common.Fixture{
							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}, {R: 0, G: 255, B: 0}, {R: 0, G: 0, B: 255}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						},
					},
					{
						Fixtures: []common.Fixture{
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}, {R: 0, G: 255, B: 0}, {R: 0, G: 0, B: 255}}},
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
					},
				},
				14: {
					{
						Fixture:       0,
						StartPosition: 14,
						Color:         common.Color{R: 0, G: 255, B: 0},
					},
				},
				28: {
					{
						Fixture:       0,
						StartPosition: 28,
						Color:         common.Color{R: 0, G: 0, B: 255},
					},
				},
				42: {
					{
						Fixture:       1,
						StartPosition: 42,
						Color:         common.Color{R: 255, G: 0, B: 0},
					},
				},
				56: {
					{
						Fixture:       1,
						StartPosition: 56,
						Color:         common.Color{R: 0, G: 255, B: 0},
					},
				},
				70: {
					{
						Fixture:       1,
						StartPosition: 70,
						Color:         common.Color{R: 0, G: 0, B: 255},
					},
				},
			},
		},
		{
			name: "Scanner case, both scanners doing same things.",
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
			name: "Scanner case, both scanners doing different things.",
			args: args{
				bounce: false,
				steps: []common.Step{
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 255},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 128, Tilt: 255},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 50, Tilt: 200},
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 128, Tilt: 200},
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
						Gobo:          36, Shutter: 255, Pan: 128, Tilt: 255,
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
						Gobo:          36, Shutter: 255, Pan: 128, Tilt: 200,
					},
				},
			},
		},
		{
			name: "Scanner case, one set of instruction in a pattern should create one set of positions.",
			args: args{
				bounce: false,
				steps: []common.Step{
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 255},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 50, Tilt: 200},
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
				},
				14: {
					{
						Fixture:       0,
						StartPosition: 14,
						Color:         common.Color{R: 255, G: 0, B: 0},
						Gobo:          36, Shutter: 255, Pan: 50, Tilt: 200,
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
		{
			name: "Scanners inverted no bounce",
			args: args{
				bounce: false,
				invert: true,
				steps: []common.Step{
					{
						Fixtures: []common.Fixture{
							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Pan: 0, Tilt: 0, Gobo: 1, Shutter: 255},
						},
					},
					{
						Fixtures: []common.Fixture{
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Pan: 1, Tilt: 1, Gobo: 1, Shutter: 255},
						},
					},
					{
						Fixtures: []common.Fixture{
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Pan: 2, Tilt: 2, Gobo: 1, Shutter: 255},
						},
					},
				},
			},
			want: map[int][]common.Position{
				0:  {{Fixture: 0, StartPosition: 0, Color: common.Color{R: 0, G: 0, B: 255}, Gobo: 1, Shutter: 255, Pan: 2, Tilt: 2}},
				14: {{Fixture: 0, StartPosition: 14, Color: common.Color{R: 0, G: 255, B: 0}, Gobo: 1, Shutter: 255, Pan: 1, Tilt: 1}},
				28: {{Fixture: 0, StartPosition: 28, Color: common.Color{R: 255, G: 0, B: 0}, Gobo: 1, Shutter: 255, Pan: 0, Tilt: 0}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := calculatePositions(tt.args.steps, tt.args.bounce, tt.args.invert); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %+v", got)
				t.Errorf("want =%+v", tt.want)
			}
		})
	}
}

func Test_invertColor(t *testing.T) {
	type args struct {
		color common.Color
	}
	tests := []struct {
		name    string
		args    args
		wantOut common.Color
	}{
		{
			name: "golden path",
			args: args{
				color: common.Color{
					R: 255,
					G: 255,
					B: 255,
				},
			},
			wantOut: common.Color{
				R: 0,
				G: 0,
				B: 0,
			},
		},
		{
			name: "golden path",
			args: args{
				color: common.Color{
					R: 0,
					G: 0,
					B: 0,
				},
			},
			wantOut: common.Color{
				R: 255,
				G: 255,
				B: 255,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := invertColor(tt.args.color); !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("invertColor() = %+v, want %+v", gotOut, tt.wantOut)
			}
		})
	}
}

func Test_getNumberOfFixtures(t *testing.T) {

	type args struct {
		sequenceNumber int
		fixtures       *fixture.Fixtures
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ten fixtures",
			args: args{
				sequenceNumber: 0,
				fixtures: &fixture.Fixtures{
					Fixtures: []fixture.Fixture{

						{Name: "fixture1", Group: 1},
						{Name: "fixture2", Group: 1},
						{Name: "fixture3", Group: 1},
						{Name: "fixture4", Group: 1},
						{Name: "fixture5", Group: 1},
						{Name: "fixture6", Group: 1},
						{Name: "fixture7", Group: 1},
						{Name: "fixture8", Group: 1},

						{Name: "fixture1", Group: 2},
						{Name: "fixture2", Group: 2},
						{Name: "fixture3", Group: 2},

						{Name: "fixture1", Group: 3},
						{Name: "fixture2", Group: 3},
						{Name: "fixture3", Group: 3},
						{Name: "fixture4", Group: 3},

						{Name: "fixture1", Group: 4},
						{Name: "fixture2", Group: 4},
					},
				},
			},
			want: 8,
		},

		{
			name: "ten fixtures",
			args: args{
				sequenceNumber: 1,
				fixtures: &fixture.Fixtures{
					Fixtures: []fixture.Fixture{

						{Name: "fixture1", Group: 1},
						{Name: "fixture2", Group: 1},
						{Name: "fixture3", Group: 1},
						{Name: "fixture4", Group: 1},
						{Name: "fixture5", Group: 1},
						{Name: "fixture6", Group: 1},
						{Name: "fixture7", Group: 1},
						{Name: "fixture8", Group: 1},

						{Name: "fixture1", Group: 2},
						{Name: "fixture2", Group: 2},
						{Name: "fixture3", Group: 2},

						{Name: "fixture1", Group: 3},
						{Name: "fixture2", Group: 3},
						{Name: "fixture3", Group: 3},
						{Name: "fixture4", Group: 3},

						{Name: "fixture1", Group: 4},
						{Name: "fixture2", Group: 4},
					},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNumberOfFixtures(tt.args.sequenceNumber, tt.args.fixtures); got != tt.want {
				t.Errorf("getNumberOfFixtures() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
