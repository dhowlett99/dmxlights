// Copyright (C) 2024 dhowlett99.
// This is the dmxlights VU chase pattern generator.
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

func Test_generateVuChasePattern(t *testing.T) {
	type args struct {
		numberOfFixtures int
	}
	tests := []struct {
		name string
		args args
		want common.Pattern
	}{
		{
			name: "Generate standard VU chase of 8 fixtures",
			args: args{
				numberOfFixtures: 8,
			},
			want: common.Pattern{
				Name:   "VU.Meter",
				Label:  "VU.Meter",
				Number: 7,
				Steps: []common.Step{
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Color: colors.Green},
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
							0: {MasterDimmer: full, Enabled: true, Color: colors.Green},
							1: {MasterDimmer: full, Enabled: true, Color: colors.Green},
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
							0: {MasterDimmer: full, Enabled: true, Color: colors.Green},
							1: {MasterDimmer: full, Enabled: true, Color: colors.Green},
							2: {MasterDimmer: full, Enabled: true, Color: colors.Green},
							3: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							4: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							5: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							6: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							7: {MasterDimmer: full, Enabled: true, Color: colors.Black},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Color: colors.Green},
							1: {MasterDimmer: full, Enabled: true, Color: colors.Green},
							2: {MasterDimmer: full, Enabled: true, Color: colors.Green},
							3: {MasterDimmer: full, Enabled: true, Color: colors.Green},
							4: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							5: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							6: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							7: {MasterDimmer: full, Enabled: true, Color: colors.Black},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Color: colors.Green},
							1: {MasterDimmer: full, Enabled: true, Color: colors.Green},
							2: {MasterDimmer: full, Enabled: true, Color: colors.Green},
							3: {MasterDimmer: full, Enabled: true, Color: colors.Green},
							4: {MasterDimmer: full, Enabled: true, Color: colors.Green},
							5: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							6: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							7: {MasterDimmer: full, Enabled: true, Color: colors.Black},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Color: colors.Green},
							1: {MasterDimmer: full, Enabled: true, Color: colors.Green},
							2: {MasterDimmer: full, Enabled: true, Color: colors.Green},
							3: {MasterDimmer: full, Enabled: true, Color: colors.Green},
							4: {MasterDimmer: full, Enabled: true, Color: colors.Green},
							5: {MasterDimmer: full, Enabled: true, Color: colors.Yellow},
							6: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							7: {MasterDimmer: full, Enabled: true, Color: colors.Black},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Color: colors.Green},
							1: {MasterDimmer: full, Enabled: true, Color: colors.Green},
							2: {MasterDimmer: full, Enabled: true, Color: colors.Green},
							3: {MasterDimmer: full, Enabled: true, Color: colors.Green},
							4: {MasterDimmer: full, Enabled: true, Color: colors.Green},
							5: {MasterDimmer: full, Enabled: true, Color: colors.Yellow},
							6: {MasterDimmer: full, Enabled: true, Color: colors.Orange},
							7: {MasterDimmer: full, Enabled: true, Color: colors.Black},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Color: colors.Green},
							1: {MasterDimmer: full, Enabled: true, Color: colors.Green},
							2: {MasterDimmer: full, Enabled: true, Color: colors.Green},
							3: {MasterDimmer: full, Enabled: true, Color: colors.Green},
							4: {MasterDimmer: full, Enabled: true, Color: colors.Green},
							5: {MasterDimmer: full, Enabled: true, Color: colors.Yellow},
							6: {MasterDimmer: full, Enabled: true, Color: colors.Orange},
							7: {MasterDimmer: full, Enabled: true, Color: colors.Red},
						},
					},
				},
			},
		},
		{
			name: "Generate standard VU chase of 16 fixtures",
			args: args{
				numberOfFixtures: 16,
			},
			want: common.Pattern{
				Name:   "VU.Meter",
				Label:  "VU.Meter",
				Number: 7,
				Steps: []common.Step{
					{
						Fixtures: map[int]common.Fixture{
							0:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							1:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							2:  {MasterDimmer: full, Enabled: true, Color: colors.Black},
							3:  {MasterDimmer: full, Enabled: true, Color: colors.Black},
							4:  {MasterDimmer: full, Enabled: true, Color: colors.Black},
							5:  {MasterDimmer: full, Enabled: true, Color: colors.Black},
							6:  {MasterDimmer: full, Enabled: true, Color: colors.Black},
							7:  {MasterDimmer: full, Enabled: true, Color: colors.Black},
							8:  {MasterDimmer: full, Enabled: true, Color: colors.Black},
							9:  {MasterDimmer: full, Enabled: true, Color: colors.Black},
							10: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							11: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							12: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							13: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							14: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							15: {MasterDimmer: full, Enabled: true, Color: colors.Black},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							1:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							2:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							3:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							4:  {MasterDimmer: full, Enabled: true, Color: colors.Black},
							5:  {MasterDimmer: full, Enabled: true, Color: colors.Black},
							6:  {MasterDimmer: full, Enabled: true, Color: colors.Black},
							7:  {MasterDimmer: full, Enabled: true, Color: colors.Black},
							8:  {MasterDimmer: full, Enabled: true, Color: colors.Black},
							9:  {MasterDimmer: full, Enabled: true, Color: colors.Black},
							10: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							11: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							12: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							13: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							14: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							15: {MasterDimmer: full, Enabled: true, Color: colors.Black},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							1:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							2:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							3:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							4:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							5:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							6:  {MasterDimmer: full, Enabled: true, Color: colors.Black},
							7:  {MasterDimmer: full, Enabled: true, Color: colors.Black},
							8:  {MasterDimmer: full, Enabled: true, Color: colors.Black},
							9:  {MasterDimmer: full, Enabled: true, Color: colors.Black},
							10: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							11: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							12: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							13: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							14: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							15: {MasterDimmer: full, Enabled: true, Color: colors.Black},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							1:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							2:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							3:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							4:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							5:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							6:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							7:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							8:  {MasterDimmer: full, Enabled: true, Color: colors.Black},
							9:  {MasterDimmer: full, Enabled: true, Color: colors.Black},
							10: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							11: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							12: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							13: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							14: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							15: {MasterDimmer: full, Enabled: true, Color: colors.Black},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							1:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							2:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							3:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							4:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							5:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							6:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							7:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							8:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							9:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							10: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							11: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							12: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							13: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							14: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							15: {MasterDimmer: full, Enabled: true, Color: colors.Black},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							1:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							2:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							3:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							4:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							5:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							6:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							7:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							8:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							9:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							10: {MasterDimmer: full, Enabled: true, Color: colors.Yellow},
							11: {MasterDimmer: full, Enabled: true, Color: colors.Yellow},
							12: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							13: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							14: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							15: {MasterDimmer: full, Enabled: true, Color: colors.Black},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							1:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							2:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							3:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							4:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							5:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							6:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							7:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							8:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							9:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							10: {MasterDimmer: full, Enabled: true, Color: colors.Yellow},
							11: {MasterDimmer: full, Enabled: true, Color: colors.Yellow},
							12: {MasterDimmer: full, Enabled: true, Color: colors.Orange},
							13: {MasterDimmer: full, Enabled: true, Color: colors.Orange},
							14: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							15: {MasterDimmer: full, Enabled: true, Color: colors.Black},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							1:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							2:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							3:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							4:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							5:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							6:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							7:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							8:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							9:  {MasterDimmer: full, Enabled: true, Color: colors.Green},
							10: {MasterDimmer: full, Enabled: true, Color: colors.Yellow},
							11: {MasterDimmer: full, Enabled: true, Color: colors.Yellow},
							12: {MasterDimmer: full, Enabled: true, Color: colors.Orange},
							13: {MasterDimmer: full, Enabled: true, Color: colors.Orange},
							14: {MasterDimmer: full, Enabled: true, Color: colors.Red},
							15: {MasterDimmer: full, Enabled: true, Color: colors.Red},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateVuChasePattern(tt.args.numberOfFixtures); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateVuChasePattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
