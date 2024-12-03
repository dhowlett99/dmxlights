// Copyright (C) 2022,2023 dhowlett99.
// This is the dmxlights pattern generator test code.
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
	"fmt"
	"reflect"
	"testing"

	"github.com/dhowlett99/dmxlights/pkg/colors"
	"github.com/dhowlett99/dmxlights/pkg/common"
)

func Test_circleGenerator(t *testing.T) {
	type args struct {
		radius            int
		numberCoordinates int
		posX              float64
		posY              float64
	}
	tests := []struct {
		name    string
		args    args
		wantOut []Coordinate
	}{
		{
			name: "standard circle",
			args: args{
				radius:            126,
				numberCoordinates: 36,
				posX:              128,
				posY:              128,
			},
			wantOut: []Coordinate{
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
				//{127, 254},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := CircleGenerator(tt.args.radius, tt.args.numberCoordinates, tt.args.posX, tt.args.posY); !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("circleGenerator() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func Test_generateScannerPattern(t *testing.T) {

	allFixturesEnabled := map[int]common.FixtureState{
		0: {
			Enabled: true,
		},
		1: {
			Enabled: true,
		},
		2: {
			Enabled: true,
		},
		3: {
			Enabled: true,
		},
		4: {
			Enabled: true,
		},
		5: {
			Enabled: true,
		},
		6: {
			Enabled: true,
		},
		7: {
			Enabled: true,
		},
	}

	tests := []struct {
		name         string
		fixtures     int
		shift        int
		chase        bool
		scannerState map[int]common.FixtureState
		Coordinates  []Coordinate
		want         common.Pattern
	}{
		{
			name:         "circle pattern - 8 point , no shift",
			fixtures:     1,
			shift:        0,
			scannerState: allFixturesEnabled,
			Coordinates: []Coordinate{
				{
					Tilt: 0,
					Pan:  128,
				},
				{
					Tilt: 32,
					Pan:  192,
				},
				{
					Tilt: 128,
					Pan:  232,
				},
				{
					Tilt: 232,
					Pan:  192,
				},
				{
					Tilt: 255,
					Pan:  128,
				},
				{
					Tilt: 232,
					Pan:  64,
				},
				{
					Tilt: 128,
					Pan:  32,
				},
				{
					Tilt: 32,
					Pan:  64,
				},
			},
			want: common.Pattern{
				Steps: []common.Step{
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Red, Gobo: 0, Shutter: 255, Pan: 128, Tilt: 0, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Orange, Gobo: 0, Shutter: 255, Pan: 192, Tilt: 32, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Yellow, Gobo: 0, Shutter: 255, Pan: 232, Tilt: 128, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Green, Gobo: 0, Shutter: 255, Pan: 192, Tilt: 232, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Cyan, Gobo: 0, Shutter: 255, Pan: 128, Tilt: 255, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Blue, Gobo: 0, Shutter: 255, Pan: 64, Tilt: 232, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Purple, Gobo: 0, Shutter: 255, Pan: 32, Tilt: 128, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Magenta, Gobo: 0, Shutter: 255, Pan: 64, Tilt: 32, ScannerColor: colors.White},
						},
					},
				},
			},
		},
		{
			name:         "two fixtures, circle pattern - 8 point , with shift of 1/4",
			fixtures:     2,
			shift:        1,
			scannerState: allFixturesEnabled,
			Coordinates: []Coordinate{
				{
					Tilt: 0,
					Pan:  0,
				},
				{
					Tilt: 1,
					Pan:  1,
				},
				{
					Tilt: 2,
					Pan:  2,
				},
				{
					Tilt: 3,
					Pan:  3,
				},
				{
					Tilt: 4,
					Pan:  4,
				},
				{
					Tilt: 5,
					Pan:  5,
				},
				{
					Tilt: 6,
					Pan:  6,
				},
				{
					Tilt: 7,
					Pan:  7,
				},
			},
			want: common.Pattern{
				Steps: []common.Step{
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Red, Gobo: 0, Shutter: 255, Pan: 0, Tilt: 0, ScannerColor: colors.White},
							1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Yellow, Gobo: 0, Shutter: 255, Pan: 2, Tilt: 2, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Orange, Gobo: 0, Shutter: 255, Pan: 1, Tilt: 1, ScannerColor: colors.White},
							1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Green, Gobo: 0, Shutter: 255, Pan: 3, Tilt: 3, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Yellow, Gobo: 0, Shutter: 255, Pan: 2, Tilt: 2, ScannerColor: colors.White},
							1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Cyan, Gobo: 0, Shutter: 255, Pan: 4, Tilt: 4, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Green, Gobo: 0, Shutter: 255, Pan: 3, Tilt: 3, ScannerColor: colors.White},
							1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Blue, Gobo: 0, Shutter: 255, Pan: 5, Tilt: 5, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Cyan, Gobo: 0, Shutter: 255, Pan: 4, Tilt: 4, ScannerColor: colors.White},
							1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Purple, Gobo: 0, Shutter: 255, Pan: 6, Tilt: 6, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Blue, Gobo: 0, Shutter: 255, Pan: 5, Tilt: 5, ScannerColor: colors.White},
							1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Magenta, Gobo: 0, Shutter: 255, Pan: 7, Tilt: 7, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Purple, Gobo: 0, Shutter: 255, Pan: 6, Tilt: 6, ScannerColor: colors.White},
							1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Red, Gobo: 0, Shutter: 255, Pan: 0, Tilt: 0, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Magenta, Gobo: 0, Shutter: 255, Pan: 7, Tilt: 7, ScannerColor: colors.White},
							1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Orange, Gobo: 0, Shutter: 255, Pan: 1, Tilt: 1, ScannerColor: colors.White},
						},
					},
				},
			},
		},
		{
			name:         "two fixtures, circle pattern - 8 point , with shift of zero",
			fixtures:     2,
			shift:        0,
			scannerState: allFixturesEnabled,
			Coordinates: []Coordinate{
				{
					Tilt: 0,
					Pan:  0,
				},
				{
					Tilt: 1,
					Pan:  1,
				},
				{
					Tilt: 2,
					Pan:  2,
				},
				{
					Tilt: 3,
					Pan:  3,
				},
				{
					Tilt: 4,
					Pan:  4,
				},
				{
					Tilt: 5,
					Pan:  5,
				},
				{
					Tilt: 6,
					Pan:  6,
				},
				{
					Tilt: 7,
					Pan:  7,
				},
			},
			want: common.Pattern{
				Steps: []common.Step{
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Red, Gobo: 0, Shutter: 255, Pan: 0, Tilt: 0, ScannerColor: colors.White},
							1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Red, Gobo: 0, Shutter: 255, Pan: 0, Tilt: 0, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Orange, Gobo: 0, Shutter: 255, Pan: 1, Tilt: 1, ScannerColor: colors.White},
							1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Orange, Gobo: 0, Shutter: 255, Pan: 1, Tilt: 1, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Yellow, Gobo: 0, Shutter: 255, Pan: 2, Tilt: 2, ScannerColor: colors.White},
							1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Yellow, Gobo: 0, Shutter: 255, Pan: 2, Tilt: 2, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Green, Gobo: 0, Shutter: 255, Pan: 3, Tilt: 3, ScannerColor: colors.White},
							1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Green, Gobo: 0, Shutter: 255, Pan: 3, Tilt: 3, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Cyan, Gobo: 0, Shutter: 255, Pan: 4, Tilt: 4, ScannerColor: colors.White},
							1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Cyan, Gobo: 0, Shutter: 255, Pan: 4, Tilt: 4, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Blue, Gobo: 0, Shutter: 255, Pan: 5, Tilt: 5, ScannerColor: colors.White},
							1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Blue, Gobo: 0, Shutter: 255, Pan: 5, Tilt: 5, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Purple, Gobo: 0, Shutter: 255, Pan: 6, Tilt: 6, ScannerColor: colors.White},
							1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Purple, Gobo: 0, Shutter: 255, Pan: 6, Tilt: 6, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Magenta, Gobo: 0, Shutter: 255, Pan: 7, Tilt: 7, ScannerColor: colors.White},
							1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Magenta, Gobo: 0, Shutter: 255, Pan: 7, Tilt: 7, ScannerColor: colors.White},
						},
					},
				},
			},
		},
		{
			name:         "four fixtures, circle pattern - 8 point , with shift of 1/4 and chase turned on",
			fixtures:     4,
			shift:        1,
			scannerState: allFixturesEnabled,
			chase:        true,
			Coordinates: []Coordinate{
				{
					Tilt: 0,
					Pan:  0,
				},
				{
					Tilt: 1,
					Pan:  1,
				},
				{
					Tilt: 2,
					Pan:  2,
				},
				{
					Tilt: 3,
					Pan:  3,
				},
				{
					Tilt: 4,
					Pan:  4,
				},
				{
					Tilt: 5,
					Pan:  5,
				},
				{
					Tilt: 6,
					Pan:  6,
				},
				{
					Tilt: 7,
					Pan:  7,
				},
			},
			want: common.Pattern{
				Steps: []common.Step{
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Red, Gobo: 0, Shutter: 255, Pan: 0, Tilt: 0, ScannerColor: colors.White},
							1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Yellow, Gobo: 0, Shutter: 255, Pan: 2, Tilt: 2, ScannerColor: colors.White},
							2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Cyan, Gobo: 0, Shutter: 255, Pan: 4, Tilt: 4, ScannerColor: colors.White},
							3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Purple, Gobo: 0, Shutter: 255, Pan: 6, Tilt: 6, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Orange, Gobo: 0, Shutter: 255, Pan: 1, Tilt: 1, ScannerColor: colors.White},
							1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Green, Gobo: 0, Shutter: 255, Pan: 3, Tilt: 3, ScannerColor: colors.White},
							2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Blue, Gobo: 0, Shutter: 255, Pan: 5, Tilt: 5, ScannerColor: colors.White},
							3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Magenta, Gobo: 0, Shutter: 255, Pan: 7, Tilt: 7, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Yellow, Gobo: 0, Shutter: 255, Pan: 2, Tilt: 2, ScannerColor: colors.White},
							1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Cyan, Gobo: 0, Shutter: 255, Pan: 4, Tilt: 4, ScannerColor: colors.White},
							2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Purple, Gobo: 0, Shutter: 255, Pan: 6, Tilt: 6, ScannerColor: colors.White},
							3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Red, Gobo: 0, Shutter: 255, Pan: 0, Tilt: 0, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Green, Gobo: 0, Shutter: 255, Pan: 3, Tilt: 3, ScannerColor: colors.White},
							1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Blue, Gobo: 0, Shutter: 255, Pan: 5, Tilt: 5, ScannerColor: colors.White},
							2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Magenta, Gobo: 0, Shutter: 255, Pan: 7, Tilt: 7, ScannerColor: colors.White},
							3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Orange, Gobo: 0, Shutter: 255, Pan: 1, Tilt: 1, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Cyan, Gobo: 0, Shutter: 255, Pan: 4, Tilt: 4, ScannerColor: colors.White},
							1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Purple, Gobo: 0, Shutter: 255, Pan: 6, Tilt: 6, ScannerColor: colors.White},
							2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Red, Gobo: 0, Shutter: 255, Pan: 0, Tilt: 0, ScannerColor: colors.White},
							3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Yellow, Gobo: 0, Shutter: 255, Pan: 2, Tilt: 2, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Blue, Gobo: 0, Shutter: 255, Pan: 5, Tilt: 5, ScannerColor: colors.White},
							1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Magenta, Gobo: 0, Shutter: 255, Pan: 7, Tilt: 7, ScannerColor: colors.White},
							2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Orange, Gobo: 0, Shutter: 255, Pan: 1, Tilt: 1, ScannerColor: colors.White},
							3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Green, Gobo: 0, Shutter: 255, Pan: 3, Tilt: 3, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Purple, Gobo: 0, Shutter: 255, Pan: 6, Tilt: 6, ScannerColor: colors.White},
							1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Red, Gobo: 0, Shutter: 255, Pan: 0, Tilt: 0, ScannerColor: colors.White},
							2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Yellow, Gobo: 0, Shutter: 255, Pan: 2, Tilt: 2, ScannerColor: colors.White},
							3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Cyan, Gobo: 0, Shutter: 255, Pan: 4, Tilt: 4, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Magenta, Gobo: 0, Shutter: 255, Pan: 7, Tilt: 7, ScannerColor: colors.White},
							1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Orange, Gobo: 0, Shutter: 255, Pan: 1, Tilt: 1, ScannerColor: colors.White},
							2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Green, Gobo: 0, Shutter: 255, Pan: 3, Tilt: 3, ScannerColor: colors.White},
							3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Blue, Gobo: 0, Shutter: 255, Pan: 5, Tilt: 5, ScannerColor: colors.White},
						},
					},
				},
			},
		},
		{
			name:         "one fixture, circle pattern - 8 point shift of 1/4 ",
			fixtures:     1,
			shift:        1,
			scannerState: allFixturesEnabled,
			Coordinates: []Coordinate{
				{
					Tilt: 0,
					Pan:  0,
				},
				{
					Tilt: 1,
					Pan:  1,
				},
				{
					Tilt: 2,
					Pan:  2,
				},
				{
					Tilt: 3,
					Pan:  3,
				},
				{
					Tilt: 4,
					Pan:  4,
				},
				{
					Tilt: 5,
					Pan:  5,
				},
				{
					Tilt: 6,
					Pan:  6,
				},
				{
					Tilt: 7,
					Pan:  7,
				},
			},
			want: common.Pattern{
				Steps: []common.Step{
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Red, Gobo: 0, Shutter: 255, Pan: 0, Tilt: 0, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Orange, Gobo: 0, Shutter: 255, Pan: 1, Tilt: 1, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Yellow, Gobo: 0, Shutter: 255, Pan: 2, Tilt: 2, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Green, Gobo: 0, Shutter: 255, Pan: 3, Tilt: 3, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Cyan, Gobo: 0, Shutter: 255, Pan: 4, Tilt: 4, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Blue, Gobo: 0, Shutter: 255, Pan: 5, Tilt: 5, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Purple, Gobo: 0, Shutter: 255, Pan: 6, Tilt: 6, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Magenta, Gobo: 0, Shutter: 255, Pan: 7, Tilt: 7, ScannerColor: colors.White},
						},
					},
				},
			},
		},
		{
			name:         "two scanners doing the same circle",
			fixtures:     2,
			shift:        0,
			scannerState: allFixturesEnabled,
			Coordinates: []Coordinate{
				{
					Tilt: 0,
					Pan:  128,
				},
				{
					Tilt: 128,
					Pan:  255,
				},
				{
					Tilt: 128,
					Pan:  0,
				},
			},
			want: common.Pattern{
				Steps: []common.Step{
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Red, Gobo: 0, Shutter: 255, Pan: 128, Tilt: 0, ScannerColor: colors.White},
							1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Red, Gobo: 0, Shutter: 255, Pan: 128, Tilt: 0, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Orange, Gobo: 0, Shutter: 255, Pan: 255, Tilt: 128, ScannerColor: colors.White},
							1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Orange, Gobo: 0, Shutter: 255, Pan: 255, Tilt: 128, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Yellow, Gobo: 0, Shutter: 255, Pan: 0, Tilt: 128, ScannerColor: colors.White},
							1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Yellow, Gobo: 0, Shutter: 255, Pan: 0, Tilt: 128, ScannerColor: colors.White},
						},
					},
				},
			},
		},
		{
			name:         "four fixtures, circle pattern - 8 point , with shift of 2 ie 1/2",
			fixtures:     4,
			shift:        2,
			scannerState: allFixturesEnabled,
			Coordinates: []Coordinate{
				{
					Tilt: 0,
					Pan:  0,
				},
				{
					Tilt: 1,
					Pan:  1,
				},
				{
					Tilt: 2,
					Pan:  2,
				},
				{
					Tilt: 3,
					Pan:  3,
				},
				{
					Tilt: 4,
					Pan:  4,
				},
				{
					Tilt: 5,
					Pan:  5,
				},
				{
					Tilt: 6,
					Pan:  6,
				},
				{
					Tilt: 7,
					Pan:  7,
				},
			},
			want: common.Pattern{
				Steps: []common.Step{
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Red, Gobo: 0, Shutter: 255, Pan: 0, Tilt: 0, ScannerColor: colors.White},
							1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Cyan, Gobo: 0, Shutter: 255, Pan: 4, Tilt: 4, ScannerColor: colors.White},
							2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Red, Gobo: 0, Shutter: 255, Pan: 0, Tilt: 0, ScannerColor: colors.White},
							3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Cyan, Gobo: 0, Shutter: 255, Pan: 4, Tilt: 4, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Orange, Gobo: 0, Shutter: 255, Pan: 1, Tilt: 1, ScannerColor: colors.White},
							1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Blue, Gobo: 0, Shutter: 255, Pan: 5, Tilt: 5, ScannerColor: colors.White},
							2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Orange, Gobo: 0, Shutter: 255, Pan: 1, Tilt: 1, ScannerColor: colors.White},
							3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Blue, Gobo: 0, Shutter: 255, Pan: 5, Tilt: 5, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Yellow, Gobo: 0, Shutter: 255, Pan: 2, Tilt: 2, ScannerColor: colors.White},
							1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Purple, Gobo: 0, Shutter: 255, Pan: 6, Tilt: 6, ScannerColor: colors.White},
							2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Yellow, Gobo: 0, Shutter: 255, Pan: 2, Tilt: 2, ScannerColor: colors.White},
							3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Purple, Gobo: 0, Shutter: 255, Pan: 6, Tilt: 6, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Green, Gobo: 0, Shutter: 255, Pan: 3, Tilt: 3, ScannerColor: colors.White},
							1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Magenta, Gobo: 0, Shutter: 255, Pan: 7, Tilt: 7, ScannerColor: colors.White},
							2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Green, Gobo: 0, Shutter: 255, Pan: 3, Tilt: 3, ScannerColor: colors.White},
							3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Magenta, Gobo: 0, Shutter: 255, Pan: 7, Tilt: 7, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Cyan, Gobo: 0, Shutter: 255, Pan: 4, Tilt: 4, ScannerColor: colors.White},
							1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Red, Gobo: 0, Shutter: 255, Pan: 0, Tilt: 0, ScannerColor: colors.White},
							2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Cyan, Gobo: 0, Shutter: 255, Pan: 4, Tilt: 4, ScannerColor: colors.White},
							3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Red, Gobo: 0, Shutter: 255, Pan: 0, Tilt: 0, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Blue, Gobo: 0, Shutter: 255, Pan: 5, Tilt: 5, ScannerColor: colors.White},
							1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Orange, Gobo: 0, Shutter: 255, Pan: 1, Tilt: 1, ScannerColor: colors.White},
							2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Blue, Gobo: 0, Shutter: 255, Pan: 5, Tilt: 5, ScannerColor: colors.White},
							3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Orange, Gobo: 0, Shutter: 255, Pan: 1, Tilt: 1, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Purple, Gobo: 0, Shutter: 255, Pan: 6, Tilt: 6, ScannerColor: colors.White},
							1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Yellow, Gobo: 0, Shutter: 255, Pan: 2, Tilt: 2, ScannerColor: colors.White},
							2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Purple, Gobo: 0, Shutter: 255, Pan: 6, Tilt: 6, ScannerColor: colors.White},
							3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Yellow, Gobo: 0, Shutter: 255, Pan: 2, Tilt: 2, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Magenta, Gobo: 0, Shutter: 255, Pan: 7, Tilt: 7, ScannerColor: colors.White},
							1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Green, Gobo: 0, Shutter: 255, Pan: 3, Tilt: 3, ScannerColor: colors.White},
							2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Magenta, Gobo: 0, Shutter: 255, Pan: 7, Tilt: 7, ScannerColor: colors.White},
							3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Green, Gobo: 0, Shutter: 255, Pan: 3, Tilt: 3, ScannerColor: colors.White},
						},
					},
				},
			},
		},
		{
			name:         "four fixtures, circle pattern - 8 point , with shift of 3 ie 3/4 shift",
			fixtures:     4,
			shift:        3,
			scannerState: allFixturesEnabled,
			Coordinates: []Coordinate{
				{
					Tilt: 0,
					Pan:  0,
				},
				{
					Tilt: 1,
					Pan:  1,
				},
				{
					Tilt: 2,
					Pan:  2,
				},
				{
					Tilt: 3,
					Pan:  3,
				},
				{
					Tilt: 4,
					Pan:  4,
				},
				{
					Tilt: 5,
					Pan:  5,
				},
				{
					Tilt: 6,
					Pan:  6,
				},
				{
					Tilt: 7,
					Pan:  7,
				},
			},
			want: common.Pattern{
				Steps: []common.Step{
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Red, Gobo: 0, Shutter: 255, Pan: 0, Tilt: 0, ScannerColor: colors.White},
							1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Purple, Gobo: 0, Shutter: 255, Pan: 6, Tilt: 6, ScannerColor: colors.White},
							2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Cyan, Gobo: 0, Shutter: 255, Pan: 4, Tilt: 4, ScannerColor: colors.White},
							3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Yellow, Gobo: 0, Shutter: 255, Pan: 2, Tilt: 2, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Orange, Gobo: 0, Shutter: 255, Pan: 1, Tilt: 1, ScannerColor: colors.White},
							1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Magenta, Gobo: 0, Shutter: 255, Pan: 7, Tilt: 7, ScannerColor: colors.White},
							2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Blue, Gobo: 0, Shutter: 255, Pan: 5, Tilt: 5, ScannerColor: colors.White},
							3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Green, Gobo: 0, Shutter: 255, Pan: 3, Tilt: 3, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Yellow, Gobo: 0, Shutter: 255, Pan: 2, Tilt: 2, ScannerColor: colors.White},
							1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Red, Gobo: 0, Shutter: 255, Pan: 0, Tilt: 0, ScannerColor: colors.White},
							2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Purple, Gobo: 0, Shutter: 255, Pan: 6, Tilt: 6, ScannerColor: colors.White},
							3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Cyan, Gobo: 0, Shutter: 255, Pan: 4, Tilt: 4, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Green, Gobo: 0, Shutter: 255, Pan: 3, Tilt: 3, ScannerColor: colors.White},
							1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Orange, Gobo: 0, Shutter: 255, Pan: 1, Tilt: 1, ScannerColor: colors.White},
							2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Magenta, Gobo: 0, Shutter: 255, Pan: 7, Tilt: 7, ScannerColor: colors.White},
							3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Blue, Gobo: 0, Shutter: 255, Pan: 5, Tilt: 5, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Cyan, Gobo: 0, Shutter: 255, Pan: 4, Tilt: 4, ScannerColor: colors.White},
							1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Yellow, Gobo: 0, Shutter: 255, Pan: 2, Tilt: 2, ScannerColor: colors.White},
							2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Red, Gobo: 0, Shutter: 255, Pan: 0, Tilt: 0, ScannerColor: colors.White},
							3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Purple, Gobo: 0, Shutter: 255, Pan: 6, Tilt: 6, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Blue, Gobo: 0, Shutter: 255, Pan: 5, Tilt: 5, ScannerColor: colors.White},
							1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Green, Gobo: 0, Shutter: 255, Pan: 3, Tilt: 3, ScannerColor: colors.White},
							2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Orange, Gobo: 0, Shutter: 255, Pan: 1, Tilt: 1, ScannerColor: colors.White},
							3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Magenta, Gobo: 0, Shutter: 255, Pan: 7, Tilt: 7, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Purple, Gobo: 0, Shutter: 255, Pan: 6, Tilt: 6, ScannerColor: colors.White},
							1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Cyan, Gobo: 0, Shutter: 255, Pan: 4, Tilt: 4, ScannerColor: colors.White},
							2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Yellow, Gobo: 0, Shutter: 255, Pan: 2, Tilt: 2, ScannerColor: colors.White},
							3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Red, Gobo: 0, Shutter: 255, Pan: 0, Tilt: 0, ScannerColor: colors.White},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Magenta, Gobo: 0, Shutter: 255, Pan: 7, Tilt: 7, ScannerColor: colors.White},
							1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Blue, Gobo: 0, Shutter: 255, Pan: 5, Tilt: 5, ScannerColor: colors.White},
							2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Green, Gobo: 0, Shutter: 255, Pan: 3, Tilt: 3, ScannerColor: colors.White},
							3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: colors.Orange, Gobo: 0, Shutter: 255, Pan: 1, Tilt: 1, ScannerColor: colors.White},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateScannerPattern(tt.Coordinates, tt.fixtures, tt.shift, tt.chase, allFixturesEnabled); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Got = %+v", got)
				t.Errorf("Want = %+v", tt.want)
			}
		})
	}
}

func Test_getEnabledScanner(t *testing.T) {
	type args struct {
		scannerState          map[int]common.FixtureState
		numberCoordinates     int
		numberEnabledScanners int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test 8 scanners 3 disabled",
			args: args{
				scannerState: map[int]common.FixtureState{
					0: {Enabled: false, RGBInverted: false},
					1: {Enabled: false, RGBInverted: false},
					2: {Enabled: false, RGBInverted: false},
					3: {Enabled: true, RGBInverted: false},
					4: {Enabled: true, RGBInverted: false},
					5: {Enabled: true, RGBInverted: false},
					6: {Enabled: true, RGBInverted: false},
					7: {Enabled: true, RGBInverted: false},
				},
				numberCoordinates:     64,
				numberEnabledScanners: 5,
			},
			want: []int{3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
				4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
				5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5,
				6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
				7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeEnabledScannerList(tt.args.scannerState, tt.args.numberCoordinates, tt.args.numberEnabledScanners, 8); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getEnabledScanner() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScanGenerateSawTooth(t *testing.T) {
	tests := []struct {
		name              string
		size              float64
		frequency         float64
		numberCoordinates float64
		posX              float64
		posY              float64
		wantOut           []Coordinate
	}{
		{
			name:              "simple case - frequency 64",
			size:              127,
			frequency:         64,
			numberCoordinates: 64,
			posX:              127,
			posY:              127,
			wantOut: []Coordinate{
				{Pan: 0, Tilt: 0},
				{Pan: 3, Tilt: 31},
				{Pan: 7, Tilt: 63},
				{Pan: 11, Tilt: 94},
				{Pan: 15, Tilt: 126},
				{Pan: 19, Tilt: 158},
				{Pan: 23, Tilt: 189},
				{Pan: 27, Tilt: 221},
				{Pan: 31, Tilt: 253},
				{Pan: 35, Tilt: 223},
				{Pan: 39, Tilt: 191},
				{Pan: 43, Tilt: 160},
				{Pan: 47, Tilt: 128},
				{Pan: 51, Tilt: 96},
				{Pan: 55, Tilt: 65},
				{Pan: 59, Tilt: 33},
				{Pan: 63, Tilt: 1},
				{Pan: 67, Tilt: 29},
				{Pan: 71, Tilt: 61},
				{Pan: 75, Tilt: 92},
				{Pan: 79, Tilt: 124},
				{Pan: 83, Tilt: 156},
				{Pan: 87, Tilt: 187},
				{Pan: 91, Tilt: 219},
				{Pan: 95, Tilt: 251},
				{Pan: 99, Tilt: 225},
				{Pan: 103, Tilt: 193},
				{Pan: 107, Tilt: 162},
				{Pan: 111, Tilt: 130},
				{Pan: 115, Tilt: 98},
				{Pan: 119, Tilt: 67},
				{Pan: 123, Tilt: 35},
				{Pan: 127, Tilt: 3},
				{Pan: 131, Tilt: 27},
				{Pan: 135, Tilt: 59},
				{Pan: 139, Tilt: 90},
				{Pan: 143, Tilt: 122},
				{Pan: 147, Tilt: 154},
				{Pan: 151, Tilt: 185},
				{Pan: 155, Tilt: 217},
				{Pan: 159, Tilt: 249},
				{Pan: 163, Tilt: 227},
				{Pan: 167, Tilt: 195},
				{Pan: 171, Tilt: 164},
				{Pan: 175, Tilt: 132},
				{Pan: 179, Tilt: 100},
				{Pan: 183, Tilt: 69},
				{Pan: 187, Tilt: 37},
				{Pan: 191, Tilt: 5},
				{Pan: 195, Tilt: 25},
				{Pan: 199, Tilt: 57},
				{Pan: 203, Tilt: 88},
				{Pan: 207, Tilt: 120},
				{Pan: 211, Tilt: 152},
				{Pan: 215, Tilt: 183},
				{Pan: 219, Tilt: 215},
				{Pan: 223, Tilt: 247},
				{Pan: 227, Tilt: 229},
				{Pan: 231, Tilt: 197},
				{Pan: 235, Tilt: 166},
				{Pan: 239, Tilt: 134},
				{Pan: 243, Tilt: 102},
				{Pan: 247, Tilt: 71},
				{Pan: 251, Tilt: 39},
			},
		},
		{
			name:              "simple case - frequency 32",
			size:              127,
			frequency:         32,
			numberCoordinates: 32,
			posX:              127,
			posY:              127,
			wantOut: []Coordinate{
				{Pan: 0, Tilt: 0},
				{Pan: 7, Tilt: 126},
				{Pan: 15, Tilt: 253},
				{Pan: 23, Tilt: 128},
				{Pan: 31, Tilt: 1},
				{Pan: 39, Tilt: 124},
				{Pan: 47, Tilt: 251},
				{Pan: 55, Tilt: 130},
				{Pan: 63, Tilt: 3},
				{Pan: 71, Tilt: 122},
				{Pan: 79, Tilt: 249},
				{Pan: 87, Tilt: 132},
				{Pan: 95, Tilt: 5},
				{Pan: 103, Tilt: 120},
				{Pan: 111, Tilt: 247},
				{Pan: 119, Tilt: 134},
				{Pan: 127, Tilt: 7},
				{Pan: 135, Tilt: 118},
				{Pan: 143, Tilt: 245},
				{Pan: 151, Tilt: 136},
				{Pan: 159, Tilt: 9},
				{Pan: 167, Tilt: 116},
				{Pan: 175, Tilt: 243},
				{Pan: 183, Tilt: 138},
				{Pan: 191, Tilt: 11},
				{Pan: 199, Tilt: 114},
				{Pan: 207, Tilt: 241},
				{Pan: 215, Tilt: 140},
				{Pan: 223, Tilt: 13},
				{Pan: 231, Tilt: 112},
				{Pan: 239, Tilt: 239},
				{Pan: 247, Tilt: 142},
			},
		},
		{
			name:              "simple case - frequency 24",
			size:              127,
			frequency:         24,
			numberCoordinates: 24,
			posX:              127,
			posY:              127,
			wantOut: []Coordinate{
				{Pan: 0, Tilt: 0},
				{Pan: 10, Tilt: 224},
				{Pan: 21, Tilt: 58},
				{Pan: 31, Tilt: 166},
				{Pan: 42, Tilt: 116},
				{Pan: 53, Tilt: 108},
				{Pan: 63, Tilt: 174},
				{Pan: 74, Tilt: 50},
				{Pan: 85, Tilt: 232},
				{Pan: 95, Tilt: 7},
				{Pan: 106, Tilt: 216},
				{Pan: 116, Tilt: 66},
				{Pan: 127, Tilt: 158},
				{Pan: 138, Tilt: 124},
				{Pan: 148, Tilt: 100},
				{Pan: 159, Tilt: 182},
				{Pan: 170, Tilt: 42},
				{Pan: 180, Tilt: 240},
				{Pan: 191, Tilt: 15},
				{Pan: 201, Tilt: 209},
				{Pan: 212, Tilt: 74},
				{Pan: 223, Tilt: 150},
				{Pan: 233, Tilt: 132},
				{Pan: 244, Tilt: 92},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := ScanGenerateSawTooth(tt.size, tt.frequency, tt.numberCoordinates, tt.posX, tt.posY); !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("ScanGenerateSawTooth() = %v, want %v", gotOut, tt.wantOut)

				for _, coordinate := range gotOut {

					fmt.Printf("{Pan: %d,Tilt:%d},\n", coordinate.Pan, coordinate.Tilt)

				}
				for _, coordinate := range gotOut {

					fmt.Printf("%d,%d\n", coordinate.Pan, coordinate.Tilt)

				}
			}
		})
	}
}

func Test_findStart(t *testing.T) {
	type args struct {
		posY   int
		maxDMX int
	}
	tests := []struct {
		name  string
		args  args
		start float64
		stop  float64
	}{

		{
			name: "pan right",
			args: args{
				posY:   255,
				maxDMX: 127,
			},
			start: 252,
			stop:  255,
		},

		{
			name: "half pan right",
			args: args{
				posY:   190,
				maxDMX: 127,
			},
			start: 188,
			stop:  255,
		},

		{
			name: "centre point",
			args: args{
				posY:   127,
				maxDMX: 127,
			},
			start: 0,
			stop:  255,
		},

		{
			name: "half pan left",
			args: args{
				posY:   64,
				maxDMX: 127,
			},
			start: 0,
			stop:  128,
		},

		{
			name: "pan left",
			args: args{
				posY:   1,
				maxDMX: 127,
			},
			start: 0,
			stop:  2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if start, stop := findStart(tt.args.posY, tt.args.maxDMX); start != tt.start || stop != tt.stop {

				if !reflect.DeepEqual(start, tt.start) {
					t.Errorf("findStart(start) = %v, want %v", start, tt.start)
				}

				if !reflect.DeepEqual(stop, tt.stop) {
					t.Errorf("findStart(stop) = %v, want %v", stop, tt.stop)
				}

			}
		})
	}
}

func Test_scaleBetween(t *testing.T) {
	type args struct {
		unscaledNum []int
		minAllowed  int
		maxAllowed  int
		min         int
		max         int
	}
	tests := []struct {
		name    string
		args    args
		wantOut []int
	}{
		{
			name: "fisrt pass",
			args: args{
				unscaledNum: []int{0, 64, 127},
				minAllowed:  0,
				maxAllowed:  10,
				min:         0,
				max:         127,
			},
			wantOut: []int{0, 5, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := scaleBetween(tt.args.unscaledNum, tt.args.minAllowed, tt.args.maxAllowed, tt.args.min, tt.args.max); !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("scaleBetween() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func Test_traingle(t *testing.T) {
	type args struct {
		y    float64
		size float64
		freq float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "standard triangle",
			args: args{
				y:    0,
				size: 255,
				freq: 50,
			},
			want: 0,
		},
		{
			name: "half triangle",
			args: args{
				y:    25,
				size: 255,
				freq: 50,
			},
			want: 255,
		},
		{
			name: "2nd cylce triangle",
			args: args{
				y:    50,
				size: 255,
				freq: 50,
			},
			want: 0,
		},

		{
			name: "2nd half cylce triangle",
			args: args{
				y:    75,
				size: 255,
				freq: 50,
			},
			want: 255,
		},
		{
			name: "3rd cylce triangle",
			args: args{
				y:    100,
				size: 255,
				freq: 50,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := traingle(tt.args.y, tt.args.size, tt.args.freq); got != tt.want {
				t.Errorf("traingle() = %v, want %v", got, tt.want)
			}
		})
	}
}
