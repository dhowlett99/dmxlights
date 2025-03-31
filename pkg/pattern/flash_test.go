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

func Test_generateFlashPattern(t *testing.T) {
	type args struct {
		numberOfFixtures int
	}
	tests := []struct {
		name string
		args args
		want common.Pattern
	}{
		{
			name: "Generate standard flash pattern of 8 fixtures",
			args: args{
				numberOfFixtures: 8,
			},
			want: common.Pattern{
				Name:   "Flash",
				Number: 1,
				Label:  "Flash",
				Steps: []common.Step{
					{
						KeyStep: true,
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Color: colors.White},
							1: {MasterDimmer: full, Enabled: true, Color: colors.White},
							2: {MasterDimmer: full, Enabled: true, Color: colors.White},
							3: {MasterDimmer: full, Enabled: true, Color: colors.White},
							4: {MasterDimmer: full, Enabled: true, Color: colors.White},
							5: {MasterDimmer: full, Enabled: true, Color: colors.White},
							6: {MasterDimmer: full, Enabled: true, Color: colors.White},
							7: {MasterDimmer: full, Enabled: true, Color: colors.White},
						},
					},
					{
						KeyStep: true,
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							1: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							2: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							3: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							4: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							5: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							6: {MasterDimmer: full, Enabled: true, Color: colors.Black},
							7: {MasterDimmer: full, Enabled: true, Color: colors.Black},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateFlashPattern(tt.args.numberOfFixtures); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateFlashPattern() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
