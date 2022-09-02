package sequence

import (
	"reflect"
	"testing"

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
)

func Test_calculateScannerPositions(t *testing.T) {

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
						ScannerNumber: 0,
						StartPosition: 0,
						Color:         common.Color{R: 0, G: 255, B: 0},
					},
				},
				14: {
					{
						ScannerNumber: 1,
						StartPosition: 14,
						Color:         common.Color{R: 0, G: 255, B: 0},
					},
				},
				28: {
					{
						ScannerNumber: 2,
						StartPosition: 28,
						Color:         common.Color{R: 0, G: 255, B: 0},
					},
				},
				42: {
					{
						ScannerNumber: 3,
						StartPosition: 42,
						Color:         common.Color{R: 0, G: 255, B: 0},
					},
				},
				56: {
					{
						ScannerNumber: 4,
						StartPosition: 56,
						Color:         common.Color{R: 0, G: 255, B: 0},
					},
				},
				70: {
					{
						ScannerNumber: 5,
						StartPosition: 70,
						Color:         common.Color{R: 0, G: 255, B: 0},
					},
				},
				84: {
					{
						ScannerNumber: 6,
						StartPosition: 84,
						Color:         common.Color{R: 0, G: 255, B: 0},
					},
				},

				98: {{
					ScannerNumber: 7,
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
						ScannerNumber: 0,
						StartPosition: 0,
						Color:         common.Color{R: 255, G: 0, B: 0},
					},
				},
				14: {
					{
						ScannerNumber: 0,
						StartPosition: 14,
						Color:         common.Color{R: 0, G: 255, B: 0},
					},
				},
				28: {
					{
						ScannerNumber: 0,
						StartPosition: 28,
						Color:         common.Color{R: 0, G: 0, B: 255},
					},
				},
				42: {
					{
						ScannerNumber: 1,
						StartPosition: 42,
						Color:         common.Color{R: 255, G: 0, B: 0},
					},
				},
				56: {
					{
						ScannerNumber: 1,
						StartPosition: 56,
						Color:         common.Color{R: 0, G: 255, B: 0},
					},
				},
				70: {
					{
						ScannerNumber: 1,
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
						ScannerNumber: 0,
						StartPosition: 0,
						Color:         common.Color{R: 0, G: 0, B: 255},
						Gobo:          36, Shutter: 255, Pan: 0, Tilt: 255,
					},
					{
						ScannerNumber: 1,
						StartPosition: 0,
						Color:         common.Color{R: 0, G: 0, B: 255},
						Gobo:          36, Shutter: 255, Pan: 0, Tilt: 255,
					},
				},
				14: {
					{
						ScannerNumber: 0,
						StartPosition: 14,
						Color:         common.Color{R: 255, G: 0, B: 0},
						Gobo:          36, Shutter: 255, Pan: 50, Tilt: 200,
					},
					{
						ScannerNumber: 1,
						StartPosition: 14,
						Color:         common.Color{R: 255, G: 0, B: 0},
						Gobo:          36, Shutter: 255, Pan: 50, Tilt: 200,
					},
				},
				28: {
					{
						ScannerNumber: 0,
						StartPosition: 28,
						Color:         common.Color{R: 255, G: 0, B: 255},
						Gobo:          36,
						Shutter:       255, Pan: 100, Tilt: 150,
					},
					{
						ScannerNumber: 1,
						StartPosition: 28,
						Color:         common.Color{R: 255, G: 0, B: 255},
						Gobo:          36,
						Shutter:       255, Pan: 100, Tilt: 150,
					},
				},
				42: {
					{
						ScannerNumber: 0,
						StartPosition: 42,
						Color:         common.Color{R: 0, G: 255, B: 255},
						Gobo:          36,
						Shutter:       255, Pan: 150, Tilt: 100,
					},
					{
						ScannerNumber: 1,
						StartPosition: 42,
						Color:         common.Color{R: 0, G: 255, B: 255},
						Gobo:          36,
						Shutter:       255, Pan: 150, Tilt: 100,
					},
				},
				56: {
					{
						ScannerNumber: 0,
						StartPosition: 56,
						Color:         common.Color{R: 100, G: 0, B: 255},
						Gobo:          36, Shutter: 255, Pan: 200, Tilt: 50,
					},
					{
						ScannerNumber: 1,
						StartPosition: 56,
						Color:         common.Color{R: 100, G: 0, B: 255},
						Gobo:          36, Shutter: 255, Pan: 200, Tilt: 50,
					},
				},
				70: {
					{
						ScannerNumber: 0,
						StartPosition: 70,
						Color:         common.Color{R: 0, G: 190, B: 255},
						Gobo:          36, Shutter: 255, Pan: 255, Tilt: 0,
					},
					{
						ScannerNumber: 1,
						StartPosition: 70,
						Color:         common.Color{R: 0, G: 190, B: 255},
						Gobo:          36, Shutter: 255, Pan: 255, Tilt: 0,
					},
				},
				84: {
					{
						ScannerNumber: 0,
						StartPosition: 84,
						Color:         common.Color{R: 145, G: 0, B: 255},
						Gobo:          36, Shutter: 255, Pan: 0, Tilt: 0,
					},
					{
						ScannerNumber: 1,
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
						ScannerNumber: 0,
						StartPosition: 0,
						Color:         common.Color{R: 0, G: 0, B: 255},
						Gobo:          36, Shutter: 255, Pan: 0, Tilt: 255,
					},
					{
						ScannerNumber: 1,
						StartPosition: 0,
						Color:         common.Color{R: 0, G: 0, B: 255},
						Gobo:          36, Shutter: 255, Pan: 128, Tilt: 255,
					},
				},
				14: {
					{
						ScannerNumber: 0,
						StartPosition: 14,
						Color:         common.Color{R: 255, G: 0, B: 0},
						Gobo:          36, Shutter: 255, Pan: 50, Tilt: 200,
					},
					{
						ScannerNumber: 1,
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
						ScannerNumber: 0,
						StartPosition: 0,
						Color:         common.Color{R: 0, G: 0, B: 255},
						Gobo:          36, Shutter: 255, Pan: 0, Tilt: 255,
					},
				},
				14: {
					{
						ScannerNumber: 0,
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
						ScannerNumber: 0,
						StartPosition: 0,
						Color:         common.Color{R: 255, G: 0, B: 0},
						Gobo:          0, Shutter: 0, Pan: 0, Tilt: 0,
					},
					{
						ScannerNumber: 2,
						StartPosition: 0,
						Color:         common.Color{R: 255, G: 0, B: 0},
						Gobo:          0, Shutter: 0, Pan: 0, Tilt: 0,
					},
					{
						ScannerNumber: 4,
						StartPosition: 0,
						Color:         common.Color{R: 255, G: 0, B: 0},
						Gobo:          0, Shutter: 0, Pan: 0, Tilt: 0,
					},
					{
						ScannerNumber: 6,
						StartPosition: 0,
						Color:         common.Color{R: 255, G: 0, B: 0},
						Gobo:          0, Shutter: 0, Pan: 0, Tilt: 0,
					},
				},
				14: {
					{
						ScannerNumber: 1,
						StartPosition: 14,
						Color:         common.Color{R: 255, G: 0, B: 0},
						Gobo:          0, Shutter: 0, Pan: 0, Tilt: 0,
					},
					{
						ScannerNumber: 3,
						StartPosition: 14,
						Color:         common.Color{R: 255, G: 0, B: 0},
						Gobo:          0, Shutter: 0, Pan: 0, Tilt: 0,
					},
					{
						ScannerNumber: 5,
						StartPosition: 14,
						Color:         common.Color{R: 255, G: 0, B: 0},
						Gobo:          0, Shutter: 0, Pan: 0, Tilt: 0,
					},
					{
						ScannerNumber: 7,
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
				0:  {{ScannerNumber: 0, StartPosition: 0, Color: common.Color{R: 0, G: 0, B: 255}, Gobo: 1, Shutter: 255, Pan: 2, Tilt: 2}},
				14: {{ScannerNumber: 0, StartPosition: 14, Color: common.Color{R: 0, G: 255, B: 0}, Gobo: 1, Shutter: 255, Pan: 1, Tilt: 1}},
				28: {{ScannerNumber: 0, StartPosition: 28, Color: common.Color{R: 255, G: 0, B: 0}, Gobo: 1, Shutter: 255, Pan: 0, Tilt: 0}},
			},
		},
		{
			name: "Scanners inverted with bounce",
			args: args{
				bounce: true,
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
				0:  {{ScannerNumber: 0, StartPosition: 0, Color: common.Color{R: 0, G: 0, B: 255}, Gobo: 1, Shutter: 255, Pan: 2, Tilt: 2}},
				14: {{ScannerNumber: 0, StartPosition: 14, Color: common.Color{R: 0, G: 255, B: 0}, Gobo: 1, Shutter: 255, Pan: 1, Tilt: 1}},
				28: {{ScannerNumber: 0, StartPosition: 28, Color: common.Color{R: 255, G: 0, B: 0}, Gobo: 1, Shutter: 255, Pan: 0, Tilt: 0}},
				42: {{ScannerNumber: 0, StartPosition: 42, Color: common.Color{R: 255, G: 0, B: 0}, Gobo: 1, Shutter: 255, Pan: 0, Tilt: 0}},
				56: {{ScannerNumber: 0, StartPosition: 56, Color: common.Color{R: 0, G: 255, B: 0}, Gobo: 1, Shutter: 255, Pan: 1, Tilt: 1}},
				70: {{ScannerNumber: 0, StartPosition: 70, Color: common.Color{R: 0, G: 0, B: 255}, Gobo: 1, Shutter: 255, Pan: 2, Tilt: 2}},
			},
		},
		{
			name: "Multicolored Patten",
			args: args{
				bounce: false,
				invert: false,
				steps: []common.Step{
					{
						Fixtures: []common.Fixture{
							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}},
						},
					},
					{
						Fixtures: []common.Fixture{
							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						},
					},
					{
						Fixtures: []common.Fixture{
							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}},
						},
					},
					{
						Fixtures: []common.Fixture{
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}},
						},
					},
					{
						Fixtures: []common.Fixture{
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
						},
					},
					{
						Fixtures: []common.Fixture{
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}},
						},
					},
					{
						Fixtures: []common.Fixture{

							{MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
						},
					},
					{
						Fixtures: []common.Fixture{

							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}},
						},
					},
				},
			},
			want: map[int][]common.Position{
				0: {
					{ScannerNumber: 0, StartPosition: 0, Color: common.Color{R: 255, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 1, StartPosition: 0, Color: common.Color{R: 255, G: 111, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 2, StartPosition: 0, Color: common.Color{R: 255, G: 255, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 3, StartPosition: 0, Color: common.Color{R: 0, G: 255, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 4, StartPosition: 0, Color: common.Color{R: 0, G: 255, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 5, StartPosition: 0, Color: common.Color{R: 0, G: 0, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 6, StartPosition: 0, Color: common.Color{R: 100, G: 0, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 7, StartPosition: 0, Color: common.Color{R: 255, G: 0, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
				},
				14: {
					{ScannerNumber: 0, StartPosition: 14, Color: common.Color{R: 255, G: 111, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 1, StartPosition: 14, Color: common.Color{R: 255, G: 255, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 2, StartPosition: 14, Color: common.Color{R: 0, G: 255, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 3, StartPosition: 14, Color: common.Color{R: 0, G: 255, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 4, StartPosition: 14, Color: common.Color{R: 0, G: 0, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 5, StartPosition: 14, Color: common.Color{R: 100, G: 0, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 6, StartPosition: 14, Color: common.Color{R: 255, G: 0, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 7, StartPosition: 14, Color: common.Color{R: 255, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
				},
				28: {
					{ScannerNumber: 0, StartPosition: 28, Color: common.Color{R: 255, G: 255, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 1, StartPosition: 28, Color: common.Color{R: 0, G: 255, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 2, StartPosition: 28, Color: common.Color{R: 0, G: 255, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 3, StartPosition: 28, Color: common.Color{R: 0, G: 0, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 4, StartPosition: 28, Color: common.Color{R: 100, G: 0, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 5, StartPosition: 28, Color: common.Color{R: 255, G: 0, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 6, StartPosition: 28, Color: common.Color{R: 255, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 7, StartPosition: 28, Color: common.Color{R: 255, G: 111, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
				},
				42: {
					{ScannerNumber: 0, StartPosition: 42, Color: common.Color{R: 0, G: 255, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 1, StartPosition: 42, Color: common.Color{R: 0, G: 255, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 2, StartPosition: 42, Color: common.Color{R: 0, G: 0, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 3, StartPosition: 42, Color: common.Color{R: 100, G: 0, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 4, StartPosition: 42, Color: common.Color{R: 255, G: 0, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 5, StartPosition: 42, Color: common.Color{R: 255, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 6, StartPosition: 42, Color: common.Color{R: 255, G: 111, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 7, StartPosition: 42, Color: common.Color{R: 255, G: 255, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
				},
				56: {
					{ScannerNumber: 0, StartPosition: 56, Color: common.Color{R: 0, G: 255, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 1, StartPosition: 56, Color: common.Color{R: 0, G: 0, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 2, StartPosition: 56, Color: common.Color{R: 100, G: 0, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 3, StartPosition: 56, Color: common.Color{R: 255, G: 0, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 4, StartPosition: 56, Color: common.Color{R: 255, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 5, StartPosition: 56, Color: common.Color{R: 255, G: 111, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 6, StartPosition: 56, Color: common.Color{R: 255, G: 255, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 7, StartPosition: 56, Color: common.Color{R: 0, G: 255, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
				},
				70: {
					{ScannerNumber: 0, StartPosition: 70, Color: common.Color{R: 0, G: 0, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 1, StartPosition: 70, Color: common.Color{R: 100, G: 0, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 2, StartPosition: 70, Color: common.Color{R: 255, G: 0, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 3, StartPosition: 70, Color: common.Color{R: 255, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 4, StartPosition: 70, Color: common.Color{R: 255, G: 111, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 5, StartPosition: 70, Color: common.Color{R: 255, G: 255, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 6, StartPosition: 70, Color: common.Color{R: 0, G: 255, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 7, StartPosition: 70, Color: common.Color{R: 0, G: 255, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
				},
				84: {
					{ScannerNumber: 0, StartPosition: 84, Color: common.Color{R: 100, G: 0, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 1, StartPosition: 84, Color: common.Color{R: 255, G: 0, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 2, StartPosition: 84, Color: common.Color{R: 255, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 3, StartPosition: 84, Color: common.Color{R: 255, G: 111, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 4, StartPosition: 84, Color: common.Color{R: 255, G: 255, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 5, StartPosition: 84, Color: common.Color{R: 0, G: 255, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 6, StartPosition: 84, Color: common.Color{R: 0, G: 255, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 7, StartPosition: 84, Color: common.Color{R: 0, G: 0, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
				},
				98: {
					{ScannerNumber: 0, StartPosition: 98, Color: common.Color{R: 255, G: 0, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 1, StartPosition: 98, Color: common.Color{R: 255, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 2, StartPosition: 98, Color: common.Color{R: 255, G: 111, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 3, StartPosition: 98, Color: common.Color{R: 255, G: 255, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 4, StartPosition: 98, Color: common.Color{R: 0, G: 255, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 5, StartPosition: 98, Color: common.Color{R: 0, G: 255, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 6, StartPosition: 98, Color: common.Color{R: 0, G: 0, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
					{ScannerNumber: 7, StartPosition: 98, Color: common.Color{R: 100, G: 0, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := calculateScannerPositions("scanner", tt.args.steps, tt.args.bounce, tt.args.invert, 14); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %+v", got)
				t.Errorf("want =%+v", tt.want)
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

func TestCalculateRGBPositions(t *testing.T) {

	full := 255
	type args struct {
		sequence common.Sequence
		slopeOn  []int
		slopeOff []int
	}
	tests := []struct {
		name  string
		args  args
		want  map[int]common.Position
		want1 int
	}{
		{
			name: "golden path - common par fixture RGB",
			args: args{

				slopeOn:  []int{1, 50, 255},
				slopeOff: []int{1, 50, 255},
				sequence: common.Sequence{
					Bounce:   false,
					Invert:   false,
					RGBShift: 0,
					RGBSize:  255,
					RGBFade:  0,
					Steps: []common.Step{
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
					},
				},
			},
			want: map[int]common.Position{
				0: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 1, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						4: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						5: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						6: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						7: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				1: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 50, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						4: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						5: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						6: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						7: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				2: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 254, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						4: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						5: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						6: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						7: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				3: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 1, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						4: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						5: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						6: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						7: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				4: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 50, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						4: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						5: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						6: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						7: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				5: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 254, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						4: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						5: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						6: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						7: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				6: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 1, B: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						4: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						5: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						6: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						7: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				7: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 50, B: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						4: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						5: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						6: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						7: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				8: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 254, B: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						4: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						5: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						6: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						7: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				9: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 1, B: 0}}},
						4: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						5: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						6: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						7: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				10: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 50, B: 0}}},
						4: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						5: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						6: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						7: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				11: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 254, B: 0}}},
						4: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						5: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						6: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						7: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				12: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						4: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 1, B: 0}}},
						5: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						6: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						7: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				13: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						4: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 50, B: 0}}},
						5: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						6: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						7: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				14: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						4: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 254, B: 0}}},
						5: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						6: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						7: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				15: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						4: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						5: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 1, B: 0}}},
						6: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						7: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				16: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						4: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						5: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 50, B: 0}}},
						6: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						7: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				17: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						4: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						5: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 254, B: 0}}},
						6: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						7: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				18: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						4: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						5: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						6: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 1, B: 0}}},
						7: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				19: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						4: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						5: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						6: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 50, B: 0}}},
						7: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				20: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						4: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						5: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						6: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 254, B: 0}}},
						7: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				21: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						4: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						5: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						6: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						7: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 1, B: 0}}},
					},
				},
				22: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						4: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						5: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						6: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						7: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 50, B: 0}}},
					},
				},
				23: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						4: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						5: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						6: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						7: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 254, B: 0}}},
					},
				},
			},

			want1: 24,
		},
		{
			name: "Shift1 - common par fixture RGB",
			args: args{
				slopeOn:  []int{1, 50, 255},
				slopeOff: []int{1, 50, 255},
				sequence: common.Sequence{
					Bounce:   false,
					Invert:   false,
					RGBShift: 8,
					RGBSize:  255,
					RGBFade:  10,
					Steps: []common.Step{
						{
							Fixtures: []common.Fixture{
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
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
							},
						},
						{
							Fixtures: []common.Fixture{
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
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
							},
						},
					},
				},
			},
			want: map[int]common.Position{
				0: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 1, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				1: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 50, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				2: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 254, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 1, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				3: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 50, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				4: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 254, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 1, B: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				5: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 50, B: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				6: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 254, B: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 1, B: 0}}},
					},
				},
				7: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 50, B: 0}}},
					},
				},
				8: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 254, B: 0}}},
					},
				},
			},

			want1: 9,
		},
		{
			name: "Shift1 - Inverted common par fixture RGB",
			args: args{
				slopeOn:  []int{1, 50, 255},
				slopeOff: []int{255, 50, 1},
				sequence: common.Sequence{
					Bounce:   false,
					Invert:   false,
					RGBShift: 1,
					RGBSize:  255,
					RGBFade:  10,
					Steps: []common.Step{
						{
							Fixtures: []common.Fixture{
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
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
							},
						},
						{
							Fixtures: []common.Fixture{
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
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
							},
						},
					},
				},
			},
			want: map[int]common.Position{
				0: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 1, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				1: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 50, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				2: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 254, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 1, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				3: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 50, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				4: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 254, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 1, B: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				5: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 50, B: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				6: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 254, B: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 1, B: 0}}},
					},
				},
				7: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 50, B: 0}}},
					},
				},
				8: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 254, B: 0}}},
					},
				},
			},

			want1: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := calculateRGBPositions(tt.args.sequence, tt.args.slopeOn, tt.args.slopeOff)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculateRGBPositions() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("calculateRGBPositions() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_getFadeValues(t *testing.T) {
	type args struct {
		size    float64
		fade    int
		reverse bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Fade up 1",
			args: args{
				size:    255,
				fade:    1,
				reverse: false, // Fade up
			},
			want: []int{0, 7, 29, 63, 105, 149, 191, 225, 247, 255},
		},
		{
			name: "Fade down 1",
			args: args{
				size:    255,
				fade:    1,
				reverse: true, // Fade up
			},
			want: []int{255, 247, 225, 191, 149, 105, 63, 29, 7, 0},
		},
		{
			name: "Fade up 2",
			args: args{
				size:    255,
				fade:    2,
				reverse: false, // Fade up
			},
			want: []int{0, 11, 45, 94, 149, 200, 237, 254, 255, 255},
		},
		{
			name: "Fade down 2",
			args: args{
				size:    255,
				fade:    2,
				reverse: true, // Fade up
			},
			want: []int{255, 255, 255, 243, 209, 160, 105, 54, 17, 0},
		},
		{
			name: "Fade up 3",
			args: args{
				size:    255,
				fade:    3,
				reverse: false, // Fade up
			},
			want: []int{0, 17, 63, 127, 191, 237, 255, 255, 255, 255},
		},
		{
			name: "Fade up 4",
			args: args{
				size:    255,
				fade:    4,
				reverse: false, // Fade up
			},
			want: []int{0, 23, 83, 160, 225, 254, 255, 255, 255, 255},
		},
		{
			name: "Fade up 5",
			args: args{
				size:    255,
				fade:    5,
				reverse: false, // Fade up
			},
			want: []int{0, 29, 105, 191, 247, 255, 255, 255, 255, 255},
		},
		{
			name: "Fade up 6",
			args: args{
				size:    255,
				fade:    6,
				reverse: false, // Fade up
			},
			want: []int{0, 37, 127, 217, 255, 255, 255, 255, 255, 255},
		},
		{
			name: "Fade up 7",
			args: args{
				size:    255,
				fade:    7,
				reverse: false, // Fade up
			},
			want: []int{0, 45, 149, 237, 255, 255, 255, 255, 255, 255},
		},
		{
			name: "Fade up 8",
			args: args{
				size:    255,
				fade:    8,
				reverse: false, // Fade up
			},
			want: []int{0, 54, 171, 250, 255, 255, 255, 255, 255, 255},
		},
		{
			name: "Fade up 9",
			args: args{
				size:    255,
				fade:    9,
				reverse: false, // Fade up
			},
			want: []int{0, 63, 191, 255, 255, 255, 255, 255, 255, 255},
		},
		{
			name: "Fade up 10",
			args: args{
				size:    255,
				fade:    10,
				reverse: false, // Fade up
			},
			want: []int{0, 73, 209, 255, 255, 255, 255, 255, 255, 255},
		},
		{
			name: "Fade down 10",
			args: args{
				size:    255,
				fade:    10,
				reverse: true, // Fade down
			},
			want: []int{255, 255, 255, 255, 255, 255, 255, 255, 181, 45},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFadeValues(tt.args.size, tt.args.fade, tt.args.reverse); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFadeValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getFadeOnValues(t *testing.T) {
	type args struct {
		size int
		fade int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "10 slots on at 255",
			args: args{
				size: 255,
				fade: 10,
			},
			want: []int{255, 255, 255, 255, 255, 255, 255, 255, 255, 255},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFadeOnValues(tt.args.size, tt.args.fade); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFadeOnValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_replaceColors(t *testing.T) {

	full := 255
	type args struct {
		positionsMap map[int]common.Position
		colors       []common.Color
	}
	tests := []struct {
		name string
		args args
		want map[int]common.Position
	}{
		{
			name: "first pass",
			args: args{
				positionsMap: map[int]common.Position{
					0: {
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
							1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						},
					},
					1: {
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
							2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						},
					},
					2: {
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
						},
					},
				},
				colors: []common.Color{
					{R: 255, G: 0, B: 0},
					{R: 0, G: 255, B: 0},
					{R: 0, G: 0, B: 255},
				},
			},

			want: map[int]common.Position{
				0: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				1: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				2: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceColors(tt.args.positionsMap, tt.args.colors); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("replaceColors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_replaceRGBcolorsInSteps(t *testing.T) {

	full := 255
	type args struct {
		steps  []common.Step
		colors []common.Color
	}
	tests := []struct {
		name string
		args args
		want []common.Step
	}{
		{
			name: "simple case",
			args: args{
				steps: []common.Step{
					{
						Fixtures: []common.Fixture{
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						},
					},
					{
						Fixtures: []common.Fixture{
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						},
					},
					{
						Fixtures: []common.Fixture{
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
						},
					},
				},
				colors: []common.Color{
					{R: 255, G: 0, B: 0},
					{R: 0, G: 255, B: 0},
					{R: 0, G: 0, B: 255},
				},
			},
			want: []common.Step{
				{
					Fixtures: []common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				{
					Fixtures: []common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				{
					Fixtures: []common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
					},
				},
			},
		},
		{
			name: "replace a number of colors with just one.",
			args: args{
				steps: []common.Step{
					{
						Fixtures: []common.Fixture{
							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						},
					},
					{
						Fixtures: []common.Fixture{
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						},
					},
					{
						Fixtures: []common.Fixture{
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
						},
					},
				},
				colors: []common.Color{
					{R: 0, G: 255, B: 0},
				},
			},
			want: []common.Step{
				{
					Fixtures: []common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				{
					Fixtures: []common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				{
					Fixtures: []common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceRGBcolorsInSteps(tt.args.steps, tt.args.colors); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("replaceRGBcolorsInSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_invertRGBColors(t *testing.T) {

	full := 255
	type args struct {
		steps  []common.Step
		colors []common.Color
	}
	tests := []struct {
		name string
		args args
		want []common.Step
	}{
		{
			name: "invert a single color.",
			args: args{
				steps: []common.Step{
					{
						Fixtures: []common.Fixture{
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						},
					},
					{
						Fixtures: []common.Fixture{
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						},
					},
					{
						Fixtures: []common.Fixture{
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}},
						},
					},
				},
				colors: []common.Color{
					{R: 0, G: 255, B: 0},
				},
			},
			want: []common.Step{
				{
					Fixtures: []common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					},
				},
				{
					Fixtures: []common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					},
				},
				{
					Fixtures: []common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := invertRGBColors(tt.args.steps, tt.args.colors); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("invertRGBColors() = %v, want %v", got, tt.want)
			}
		})
	}
}
