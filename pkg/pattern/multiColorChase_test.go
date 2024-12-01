// Copyright (C) 2024 dhowlett99.
// This is the dmxlights multi color chaser pattern generator.
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

func Test_generateMultiColorChasePattern(t *testing.T) {
	type args struct {
		numberOfFixtures int
	}
	tests := []struct {
		name string
		args args
		want common.Pattern
	}{
		{
			name: "Generate standard color chase of 8 fixtures",
			args: args{
				numberOfFixtures: 8,
			},
			want: common.Pattern{
				Name:   "Multi Color",
				Label:  "Multi.Color",
				Number: 6,
				Steps: []common.Step{
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Color: colors.Red},     // Red
							1: {MasterDimmer: full, Enabled: true, Color: colors.Orange},  // Orange
							2: {MasterDimmer: full, Enabled: true, Color: colors.Yellow},  // Yellow
							3: {MasterDimmer: full, Enabled: true, Color: colors.Green},   // Green
							4: {MasterDimmer: full, Enabled: true, Color: colors.Cyan},    // Cyan
							5: {MasterDimmer: full, Enabled: true, Color: colors.Blue},    // Blue
							6: {MasterDimmer: full, Enabled: true, Color: colors.Purple},  // Purple
							7: {MasterDimmer: full, Enabled: true, Color: colors.Magenta}, // Magenta
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Color: colors.Orange},  // Orange
							1: {MasterDimmer: full, Enabled: true, Color: colors.Yellow},  // Yellow
							2: {MasterDimmer: full, Enabled: true, Color: colors.Green},   // Green
							3: {MasterDimmer: full, Enabled: true, Color: colors.Cyan},    // Cyan
							4: {MasterDimmer: full, Enabled: true, Color: colors.Blue},    // Blue
							5: {MasterDimmer: full, Enabled: true, Color: colors.Purple},  // Purple
							6: {MasterDimmer: full, Enabled: true, Color: colors.Magenta}, // Magenta
							7: {MasterDimmer: full, Enabled: true, Color: colors.Red},     // Red
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Color: colors.Yellow},  // Yellow
							1: {MasterDimmer: full, Enabled: true, Color: colors.Green},   // Green
							2: {MasterDimmer: full, Enabled: true, Color: colors.Cyan},    // Cyan
							3: {MasterDimmer: full, Enabled: true, Color: colors.Blue},    // Blue
							4: {MasterDimmer: full, Enabled: true, Color: colors.Purple},  // Purple
							5: {MasterDimmer: full, Enabled: true, Color: colors.Magenta}, // Magenta
							6: {MasterDimmer: full, Enabled: true, Color: colors.Red},     // Red
							7: {MasterDimmer: full, Enabled: true, Color: colors.Orange},  // Orange
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Color: colors.Green},   // Green
							1: {MasterDimmer: full, Enabled: true, Color: colors.Cyan},    // Cyan
							2: {MasterDimmer: full, Enabled: true, Color: colors.Blue},    // Blue
							3: {MasterDimmer: full, Enabled: true, Color: colors.Purple},  // Purple
							4: {MasterDimmer: full, Enabled: true, Color: colors.Magenta}, // Magenta
							5: {MasterDimmer: full, Enabled: true, Color: colors.Red},     // Red
							6: {MasterDimmer: full, Enabled: true, Color: colors.Orange},  // Orange
							7: {MasterDimmer: full, Enabled: true, Color: colors.Yellow},  // Yellow
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Color: colors.Cyan},    // Cyan
							1: {MasterDimmer: full, Enabled: true, Color: colors.Blue},    // Blue
							2: {MasterDimmer: full, Enabled: true, Color: colors.Purple},  // Purple
							3: {MasterDimmer: full, Enabled: true, Color: colors.Magenta}, // Magenta
							4: {MasterDimmer: full, Enabled: true, Color: colors.Red},     // Red
							5: {MasterDimmer: full, Enabled: true, Color: colors.Orange},  // Orange
							6: {MasterDimmer: full, Enabled: true, Color: colors.Yellow},  // Yellow
							7: {MasterDimmer: full, Enabled: true, Color: colors.Green},   // Green
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Color: colors.Blue},    // Blue
							1: {MasterDimmer: full, Enabled: true, Color: colors.Purple},  // Purple
							2: {MasterDimmer: full, Enabled: true, Color: colors.Magenta}, // Magenta
							3: {MasterDimmer: full, Enabled: true, Color: colors.Red},     // Red
							4: {MasterDimmer: full, Enabled: true, Color: colors.Orange},  // Orange
							5: {MasterDimmer: full, Enabled: true, Color: colors.Yellow},  // Yellow
							6: {MasterDimmer: full, Enabled: true, Color: colors.Green},   // Green
							7: {MasterDimmer: full, Enabled: true, Color: colors.Cyan},    // Cyan
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Color: colors.Purple},  // Purple
							1: {MasterDimmer: full, Enabled: true, Color: colors.Magenta}, // Magenta
							2: {MasterDimmer: full, Enabled: true, Color: colors.Red},     // Red
							3: {MasterDimmer: full, Enabled: true, Color: colors.Orange},  // Orange
							4: {MasterDimmer: full, Enabled: true, Color: colors.Yellow},  // Yellow
							5: {MasterDimmer: full, Enabled: true, Color: colors.Green},   // Green
							6: {MasterDimmer: full, Enabled: true, Color: colors.Cyan},    // Cyan
							7: {MasterDimmer: full, Enabled: true, Color: colors.Blue},    // Blue
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Color: colors.Magenta}, // Magenta
							1: {MasterDimmer: full, Enabled: true, Color: colors.Red},     // Red
							2: {MasterDimmer: full, Enabled: true, Color: colors.Orange},  // Orange
							3: {MasterDimmer: full, Enabled: true, Color: colors.Yellow},  // Yellow
							4: {MasterDimmer: full, Enabled: true, Color: colors.Green},   // Green
							5: {MasterDimmer: full, Enabled: true, Color: colors.Cyan},    // Cyan
							6: {MasterDimmer: full, Enabled: true, Color: colors.Blue},    // Blue
							7: {MasterDimmer: full, Enabled: true, Color: colors.Purple},  // Purple
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateMultiColorChasePattern(tt.args.numberOfFixtures); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateMultiColorChasePattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
