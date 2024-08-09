// Copyright (C) 2022,2023 dhowlett99.
// This is the dmxlight main sequencer test code.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package sequence

import (
	"image/color"
	"reflect"
	"testing"

	"github.com/dhowlett99/dmxlights/pkg/common"
)

func Test_replaceRGBcolorsInSteps(t *testing.T) {

	full := 255
	type args struct {
		steps  []common.Step
		colors []color.NRGBA
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
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Color: common.Green},
							1: {MasterDimmer: full, Color: common.Black},
							2: {MasterDimmer: full, Color: common.Black},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Color: common.Black},
							1: {MasterDimmer: full, Color: common.Green},
							2: {MasterDimmer: full, Color: common.Black},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Color: common.Black},
							1: {MasterDimmer: full, Color: common.Black},
							2: {MasterDimmer: full, Color: common.Green},
						},
					},
					{ // Only black.
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Color: common.Black},
							1: {MasterDimmer: full, Color: common.Black},
							2: {MasterDimmer: full, Color: common.Black},
						},
					},
				},
				colors: []color.NRGBA{
					common.Red,
					common.Green,
					common.Blue,
				},
			},
			want: []common.Step{
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Color: common.Red},
						1: {MasterDimmer: full, Color: common.Black},
						2: {MasterDimmer: full, Color: common.Black},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Color: common.Black},
						1: {MasterDimmer: full, Color: common.Green},
						2: {MasterDimmer: full, Color: common.Black},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Color: common.Black},
						1: {MasterDimmer: full, Color: common.Black},
						2: {MasterDimmer: full, Color: common.Blue},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Color: common.Black},
						1: {MasterDimmer: full, Color: common.Black},
						2: {MasterDimmer: full, Color: common.Black},
					},
				},
			},
		},
		{
			name: "replace a number of colors with just one.",
			args: args{
				steps: []common.Step{
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Color: common.Red},
							1: {MasterDimmer: full, Color: common.Black},
							2: {MasterDimmer: full, Color: common.Black},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Color: common.Black},
							1: {MasterDimmer: full, Color: common.Green},
							2: {MasterDimmer: full, Color: common.Black},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Color: common.Black},
							1: {MasterDimmer: full, Color: common.Black},
							2: {MasterDimmer: full, Color: common.Blue},
						},
					},
				},
				colors: []color.NRGBA{
					common.Green,
				},
			},
			want: []common.Step{
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Color: common.Green},
						1: {MasterDimmer: full, Color: common.Black},
						2: {MasterDimmer: full, Color: common.Black},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Color: common.Black},
						1: {MasterDimmer: full, Color: common.Green},
						2: {MasterDimmer: full, Color: common.Black},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Color: common.Black},
						1: {MasterDimmer: full, Color: common.Black},
						2: {MasterDimmer: full, Color: common.Green},
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
