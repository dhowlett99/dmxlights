// Copyright (C) 2024 dhowlett99.
// This is the dmxlights rgb chase pattern generator.
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

func Test_generateRgbChasePattern(t *testing.T) {
	type args struct {
		numberOfFixtures int
	}
	tests := []struct {
		name string
		args args
		want common.Pattern
	}{
		{
			name: "Generate standard rgb chase for 8 fixtures",
			args: args{
				numberOfFixtures: 8,
			},
			want: common.Pattern{
				Name:   "RGB Chase",
				Number: 2,
				Label:  "RGB.Chase",
				Steps: []common.Step{
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Color: colors.Red},
							1: {MasterDimmer: full, Enabled: true, Color: colors.Red},
							2: {MasterDimmer: full, Enabled: true, Color: colors.Red},
							3: {MasterDimmer: full, Enabled: true, Color: colors.Red},
							4: {MasterDimmer: full, Enabled: true, Color: colors.Red},
							5: {MasterDimmer: full, Enabled: true, Color: colors.Red},
							6: {MasterDimmer: full, Enabled: true, Color: colors.Red},
							7: {MasterDimmer: full, Enabled: true, Color: colors.Red},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Color: colors.Green},
							1: {MasterDimmer: full, Enabled: true, Color: colors.Green},
							2: {MasterDimmer: full, Enabled: true, Color: colors.Green},
							3: {MasterDimmer: full, Enabled: true, Color: colors.Green},
							4: {MasterDimmer: full, Enabled: true, Color: colors.Green},
							5: {MasterDimmer: full, Enabled: true, Color: colors.Green},
							6: {MasterDimmer: full, Enabled: true, Color: colors.Green},
							7: {MasterDimmer: full, Enabled: true, Color: colors.Green},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Color: colors.Blue},
							1: {MasterDimmer: full, Enabled: true, Color: colors.Blue},
							2: {MasterDimmer: full, Enabled: true, Color: colors.Blue},
							3: {MasterDimmer: full, Enabled: true, Color: colors.Blue},
							4: {MasterDimmer: full, Enabled: true, Color: colors.Blue},
							5: {MasterDimmer: full, Enabled: true, Color: colors.Blue},
							6: {MasterDimmer: full, Enabled: true, Color: colors.Blue},
							7: {MasterDimmer: full, Enabled: true, Color: colors.Blue},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateRgbChasePattern(tt.args.numberOfFixtures); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateRgbChasePattern() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
