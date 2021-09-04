package sequence

import (
	"reflect"
	"testing"

	"github.com/dhowlett99/dmxlights/pkg/common"
)

func Test_calculatePositions(t *testing.T) {

	full := 255
	type args struct {
		steps []common.Step
	}
	tests := []struct {
		name string
		args args
		want []common.Position
	}{
		{
			name: "golden path",
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
			want: []common.Position{
				{
					Fixture:       0,
					StartPosition: 0,
					Color:         common.Color{R: 0, G: 255, B: 0},
				},
				{
					Fixture:       1,
					StartPosition: 14,
					Color:         common.Color{R: 0, G: 255, B: 0},
				},
				{
					Fixture:       2,
					StartPosition: 28,
					Color:         common.Color{R: 0, G: 255, B: 0},
				},
				{
					Fixture:       3,
					StartPosition: 42,
					Color:         common.Color{R: 0, G: 255, B: 0},
				},
				{
					Fixture:       4,
					StartPosition: 56,
					Color:         common.Color{R: 0, G: 255, B: 0},
				},
				{
					Fixture:       5,
					StartPosition: 70,
					Color:         common.Color{R: 0, G: 255, B: 0},
				},
				{
					Fixture:       6,
					StartPosition: 84,
					Color:         common.Color{R: 0, G: 255, B: 0},
				},
				{
					Fixture:       7,
					StartPosition: 98,
					Color:         common.Color{R: 0, G: 255, B: 0},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculatePositions(tt.args.steps); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %v", got)
				t.Errorf("want =%v", tt.want)
			}
		})
	}
}
