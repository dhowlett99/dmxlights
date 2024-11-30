// Copyright (C) 2024 dhowlett99.
// This is the dmxlights chase pattern generator.
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

package pattern

import (
	"reflect"
	"testing"

	"github.com/dhowlett99/dmxlights/pkg/colors"
	"github.com/dhowlett99/dmxlights/pkg/common"
)

func Test_generateChasePattern(t *testing.T) {
	type args struct {
		numberOfFixtures int
	}
	tests := []struct {
		name string
		args args
		want common.Pattern
	}{
		{
			name: "Generate standard chase of 8 fixtures",
			args: args{
				numberOfFixtures: 8,
			},
			want: common.Pattern{
				Name:   "Chase",
				Number: 0,
				Label:  "Std.Chase",
				Steps: []common.Step{
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Shutter: 255, Brightness: 255, Color: colors.Green},
							1: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							2: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							3: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							4: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							5: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							6: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							7: {MasterDimmer: full, Enabled: true, Color: colors.Black},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							1: {MasterDimmer: full, Enabled: true, Shutter: 255, Brightness: 255, Color: colors.Green},
							2: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							3: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							4: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							5: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							6: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							7: {MasterDimmer: full, Enabled: true, Color: colors.Black},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							1: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							2: {MasterDimmer: full, Enabled: true, Shutter: 255, Brightness: 255, Color: colors.Green},
							3: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							4: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							5: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							6: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							7: {MasterDimmer: full, Enabled: true, Color: colors.Black},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							1: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							2: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							3: {MasterDimmer: full, Enabled: true, Shutter: 255, Brightness: 255, Color: colors.Green},
							4: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							5: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							6: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							7: {MasterDimmer: full, Enabled: true, Color: colors.Black},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							1: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							2: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							3: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							4: {MasterDimmer: full, Enabled: true, Shutter: 255, Brightness: 255, Color: colors.Green},
							5: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							6: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							7: {MasterDimmer: full, Enabled: true, Color: colors.Black},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							1: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							2: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							3: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							4: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							5: {MasterDimmer: full, Enabled: true, Shutter: 255, Brightness: 255, Color: colors.Green},
							6: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							7: {MasterDimmer: full, Enabled: true, Color: colors.Black},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							1: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							2: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							3: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							4: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							5: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							6: {MasterDimmer: full, Enabled: true, Shutter: 255, Brightness: 255, Color: colors.Green},
							7: {MasterDimmer: full, Enabled: true, Color: colors.Black},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							1: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							2: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							3: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							4: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							5: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							6: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							7: {MasterDimmer: full, Enabled: true, Shutter: 255, Brightness: 255, Color: colors.Green},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateChasePattern(tt.args.numberOfFixtures); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateChasePattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
