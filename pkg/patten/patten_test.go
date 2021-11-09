package patten

import (
	"reflect"
	"testing"

	"github.com/dhowlett99/dmxlights/pkg/common"
)

func Test_circleGenerator(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name    string
		args    args
		wantOut []coordinate
	}{
		{
			name: "standard circle",
			args: args{
				size: 126,
			},
			wantOut: []coordinate{
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
				{127, 254},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := CircleGenerator(tt.args.size); !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("circleGenerator() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func Test_generatePatten(t *testing.T) {
	tests := []struct {
		name        string
		coordinates []coordinate
		want        common.Patten
	}{
		{
			name: "circle patten",
			coordinates: []coordinate{
				{
					x: 0,
					y: 128,
				},
				{
					x: 128,
					y: 255,
				},
				{
					x: 128,
					y: 0,
				},
			},
			want: common.Patten{
				Name: "circle",
				Steps: []common.Step{
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 128, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 128, Tilt: 0},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 128, B: 0}}, Gobo: 36, Shutter: 255, Pan: 255, Tilt: 128},
						},
					},
					{
						Type: "scanner",
						Fixtures: []common.Fixture{
							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 128, B: 0}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 128},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GeneratePatten(tt.coordinates); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GeneratePatten() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScanGenerateSineWave(t *testing.T) {
	type args struct {
		size      int
		frequency int
	}
	tests := []struct {
		name    string
		args    args
		wantOut []coordinate
	}{
		{
			name: "50hz sawtooth",
			args: args{
				size:      255,
				frequency: 5000,
			},
			wantOut: []coordinate{
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
			if gotOut := ScanGenerateSineWave(tt.args.size, tt.args.frequency); !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("ScanGenerateSineWave() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
