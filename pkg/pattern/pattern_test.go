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
	"reflect"
	"testing"

	"github.com/dhowlett99/dmxlights/pkg/common"
)

// func Test_circleGenerator(t *testing.T) {
// 	type args struct {
// 		radius            int
// 		numberCoordinates int
// 		posX              float64
// 		posY              float64
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		wantOut []Coordinate
// 	}{
// 		{
// 			name: "standard circle",
// 			args: args{
// 				radius:            126,
// 				numberCoordinates: 36,
// 				posX:              128,
// 				posY:              128,
// 			},
// 			wantOut: []Coordinate{
// 				{128, 254},
// 				{149, 252},
// 				{171, 246},
// 				{191, 237},
// 				{208, 224},
// 				{224, 208},
// 				{237, 191},
// 				{246, 171},
// 				{252, 149},
// 				{254, 128},
// 				{252, 106},
// 				{246, 84},
// 				{237, 65},
// 				{224, 47},
// 				{208, 31},
// 				{191, 18},
// 				{171, 9},
// 				{149, 3},
// 				{128, 2},
// 				{106, 3},
// 				{84, 9},
// 				{65, 18},
// 				{47, 31},
// 				{31, 47},
// 				{18, 65},
// 				{9, 84},
// 				{3, 106},
// 				{2, 127},
// 				{3, 149},
// 				{9, 171},
// 				{18, 191},
// 				{31, 208},
// 				{47, 224},
// 				{64, 237},
// 				{84, 246},
// 				{106, 252},
// 				//{127, 254},
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if gotOut := CircleGenerator(tt.args.radius, tt.args.numberCoordinates, tt.args.posX, tt.args.posY); !reflect.DeepEqual(gotOut, tt.wantOut) {
// 				t.Errorf("circleGenerator() = %v, want %v", gotOut, tt.wantOut)
// 			}
// 		})
// 	}
// }

// func Test_generatePattern(t *testing.T) {

// 	allFixturesEnabled := map[int]common.ScannerState{
// 		0: {
// 			Enabled: true,
// 		},
// 		1: {
// 			Enabled: true,
// 		},
// 		2: {
// 			Enabled: true,
// 		},
// 		3: {
// 			Enabled: true,
// 		},
// 		4: {
// 			Enabled: true,
// 		},
// 		5: {
// 			Enabled: true,
// 		},
// 		6: {
// 			Enabled: true,
// 		},
// 		7: {
// 			Enabled: true,
// 		},
// 	}

// 	tests := []struct {
// 		name         string
// 		fixtures     int
// 		shift        int
// 		chase        bool
// 		scannerState map[int]common.ScannerState
// 		Coordinates  []Coordinate
// 		want         common.Pattern
// 	}{
// 		{
// 			name:         "circle pattern - 8 point , no shift",
// 			fixtures:     1,
// 			shift:        0,
// 			scannerState: allFixturesEnabled,
// 			Coordinates: []Coordinate{
// 				{
// 					Tilt: 0,
// 					Pan:  128,
// 				},
// 				{
// 					Tilt: 32,
// 					Pan:  192,
// 				},
// 				{
// 					Tilt: 128,
// 					Pan:  232,
// 				},
// 				{
// 					Tilt: 232,
// 					Pan:  192,
// 				},
// 				{
// 					Tilt: 255,
// 					Pan:  128,
// 				},
// 				{
// 					Tilt: 232,
// 					Pan:  64,
// 				},
// 				{
// 					Tilt: 128,
// 					Pan:  32,
// 				},
// 				{
// 					Tilt: 32,
// 					Pan:  64,
// 				},
// 			},
// 			want: common.Pattern{
// 				Steps: []common.Step{
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 0, Shutter: 255, Pan: 128, Tilt: 0, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}, Gobo: 0, Shutter: 255, Pan: 192, Tilt: 32, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}, Gobo: 0, Shutter: 255, Pan: 232, Tilt: 128, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Gobo: 0, Shutter: 255, Pan: 192, Tilt: 232, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 0, Shutter: 255, Pan: 128, Tilt: 255, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 64, Tilt: 232, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 32, Tilt: 128, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 64, Tilt: 32, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 				},
// 			},
// 		},
// 		{
// 			name:         "two fixtures, circle pattern - 8 point , with shift of 1/4",
// 			fixtures:     2,
// 			shift:        1,
// 			scannerState: allFixturesEnabled,
// 			Coordinates: []Coordinate{
// 				{
// 					Tilt: 0,
// 					Pan:  0,
// 				},
// 				{
// 					Tilt: 1,
// 					Pan:  1,
// 				},
// 				{
// 					Tilt: 2,
// 					Pan:  2,
// 				},
// 				{
// 					Tilt: 3,
// 					Pan:  3,
// 				},
// 				{
// 					Tilt: 4,
// 					Pan:  4,
// 				},
// 				{
// 					Tilt: 5,
// 					Pan:  5,
// 				},
// 				{
// 					Tilt: 6,
// 					Pan:  6,
// 				},
// 				{
// 					Tilt: 7,
// 					Pan:  7,
// 				},
// 			},
// 			want: common.Pattern{
// 				Steps: []common.Step{
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 0, Shutter: 255, Pan: 0, Tilt: 0, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}, Gobo: 0, Shutter: 255, Pan: 2, Tilt: 2, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}, Gobo: 0, Shutter: 255, Pan: 1, Tilt: 1, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Gobo: 0, Shutter: 255, Pan: 3, Tilt: 3, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}, Gobo: 0, Shutter: 255, Pan: 2, Tilt: 2, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 0, Shutter: 255, Pan: 4, Tilt: 4, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Gobo: 0, Shutter: 255, Pan: 3, Tilt: 3, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 5, Tilt: 5, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 0, Shutter: 255, Pan: 4, Tilt: 4, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 6, Tilt: 6, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 5, Tilt: 5, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 7, Tilt: 7, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 6, Tilt: 6, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 0, Shutter: 255, Pan: 0, Tilt: 0, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 7, Tilt: 7, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}, Gobo: 0, Shutter: 255, Pan: 1, Tilt: 1, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 				},
// 			},
// 		},
// 		{
// 			name:         "two fixtures, circle pattern - 8 point , with shift of zero",
// 			fixtures:     2,
// 			shift:        0,
// 			scannerState: allFixturesEnabled,
// 			Coordinates: []Coordinate{
// 				{
// 					Tilt: 0,
// 					Pan:  0,
// 				},
// 				{
// 					Tilt: 1,
// 					Pan:  1,
// 				},
// 				{
// 					Tilt: 2,
// 					Pan:  2,
// 				},
// 				{
// 					Tilt: 3,
// 					Pan:  3,
// 				},
// 				{
// 					Tilt: 4,
// 					Pan:  4,
// 				},
// 				{
// 					Tilt: 5,
// 					Pan:  5,
// 				},
// 				{
// 					Tilt: 6,
// 					Pan:  6,
// 				},
// 				{
// 					Tilt: 7,
// 					Pan:  7,
// 				},
// 			},
// 			want: common.Pattern{
// 				Steps: []common.Step{
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 0, Shutter: 255, Pan: 0, Tilt: 0, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 0, Shutter: 255, Pan: 0, Tilt: 0, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}, Gobo: 0, Shutter: 255, Pan: 1, Tilt: 1, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}, Gobo: 0, Shutter: 255, Pan: 1, Tilt: 1, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}, Gobo: 0, Shutter: 255, Pan: 2, Tilt: 2, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}, Gobo: 0, Shutter: 255, Pan: 2, Tilt: 2, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Gobo: 0, Shutter: 255, Pan: 3, Tilt: 3, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Gobo: 0, Shutter: 255, Pan: 3, Tilt: 3, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 0, Shutter: 255, Pan: 4, Tilt: 4, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 0, Shutter: 255, Pan: 4, Tilt: 4, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 5, Tilt: 5, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 5, Tilt: 5, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 6, Tilt: 6, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 6, Tilt: 6, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 7, Tilt: 7, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 7, Tilt: 7, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 				},
// 			},
// 		},
// 		{
// 			name:         "four fixtures, circle pattern - 8 point , with shift of 1/4 and chase turned on",
// 			fixtures:     4,
// 			shift:        1,
// 			scannerState: allFixturesEnabled,
// 			chase:        true,
// 			Coordinates: []Coordinate{
// 				{
// 					Tilt: 0,
// 					Pan:  0,
// 				},
// 				{
// 					Tilt: 1,
// 					Pan:  1,
// 				},
// 				{
// 					Tilt: 2,
// 					Pan:  2,
// 				},
// 				{
// 					Tilt: 3,
// 					Pan:  3,
// 				},
// 				{
// 					Tilt: 4,
// 					Pan:  4,
// 				},
// 				{
// 					Tilt: 5,
// 					Pan:  5,
// 				},
// 				{
// 					Tilt: 6,
// 					Pan:  6,
// 				},
// 				{
// 					Tilt: 7,
// 					Pan:  7,
// 				},
// 			},
// 			want: common.Pattern{
// 				Steps: []common.Step{
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 0, Shutter: 255, Pan: 0, Tilt: 0, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}, Gobo: 0, Shutter: 255, Pan: 2, Tilt: 2, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 0, Shutter: 255, Pan: 4, Tilt: 4, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 6, Tilt: 6, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}, Gobo: 0, Shutter: 255, Pan: 1, Tilt: 1, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Gobo: 0, Shutter: 255, Pan: 3, Tilt: 3, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 5, Tilt: 5, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 7, Tilt: 7, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}, Gobo: 0, Shutter: 255, Pan: 2, Tilt: 2, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 0, Shutter: 255, Pan: 4, Tilt: 4, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 6, Tilt: 6, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 0, Shutter: 255, Pan: 0, Tilt: 0, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Gobo: 0, Shutter: 255, Pan: 3, Tilt: 3, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 5, Tilt: 5, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 7, Tilt: 7, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}, Gobo: 0, Shutter: 255, Pan: 1, Tilt: 1, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 0, Shutter: 255, Pan: 4, Tilt: 4, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 6, Tilt: 6, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 0, Shutter: 255, Pan: 0, Tilt: 0, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}, Gobo: 0, Shutter: 255, Pan: 2, Tilt: 2, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 5, Tilt: 5, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 7, Tilt: 7, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}, Gobo: 0, Shutter: 255, Pan: 1, Tilt: 1, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Gobo: 0, Shutter: 255, Pan: 3, Tilt: 3, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 6, Tilt: 6, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 0, Shutter: 255, Pan: 0, Tilt: 0, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}, Gobo: 0, Shutter: 255, Pan: 2, Tilt: 2, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 0, Shutter: 255, Pan: 4, Tilt: 4, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 7, Tilt: 7, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}, Gobo: 0, Shutter: 255, Pan: 1, Tilt: 1, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Gobo: 0, Shutter: 255, Pan: 3, Tilt: 3, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 5, Tilt: 5, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 				},
// 			},
// 		},
// 		{
// 			name:         "one fixture, circle pattern - 8 point shift of 1/4 ",
// 			fixtures:     1,
// 			shift:        1,
// 			scannerState: allFixturesEnabled,
// 			Coordinates: []Coordinate{
// 				{
// 					Tilt: 0,
// 					Pan:  0,
// 				},
// 				{
// 					Tilt: 1,
// 					Pan:  1,
// 				},
// 				{
// 					Tilt: 2,
// 					Pan:  2,
// 				},
// 				{
// 					Tilt: 3,
// 					Pan:  3,
// 				},
// 				{
// 					Tilt: 4,
// 					Pan:  4,
// 				},
// 				{
// 					Tilt: 5,
// 					Pan:  5,
// 				},
// 				{
// 					Tilt: 6,
// 					Pan:  6,
// 				},
// 				{
// 					Tilt: 7,
// 					Pan:  7,
// 				},
// 			},
// 			want: common.Pattern{
// 				Steps: []common.Step{
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 0, Shutter: 255, Pan: 0, Tilt: 0, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}, Gobo: 0, Shutter: 255, Pan: 1, Tilt: 1, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}, Gobo: 0, Shutter: 255, Pan: 2, Tilt: 2, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Gobo: 0, Shutter: 255, Pan: 3, Tilt: 3, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 0, Shutter: 255, Pan: 4, Tilt: 4, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 5, Tilt: 5, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 6, Tilt: 6, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 7, Tilt: 7, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 				},
// 			},
// 		},
// 		{
// 			name:         "two scanners doing the same circle",
// 			fixtures:     2,
// 			shift:        0,
// 			scannerState: allFixturesEnabled,
// 			Coordinates: []Coordinate{
// 				{
// 					Tilt: 0,
// 					Pan:  128,
// 				},
// 				{
// 					Tilt: 128,
// 					Pan:  255,
// 				},
// 				{
// 					Tilt: 128,
// 					Pan:  0,
// 				},
// 			},
// 			want: common.Pattern{
// 				Steps: []common.Step{
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 0, Shutter: 255, Pan: 128, Tilt: 0, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 0, Shutter: 255, Pan: 128, Tilt: 0, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}, Gobo: 0, Shutter: 255, Pan: 255, Tilt: 128, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}, Gobo: 0, Shutter: 255, Pan: 255, Tilt: 128, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}, Gobo: 0, Shutter: 255, Pan: 0, Tilt: 128, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}, Gobo: 0, Shutter: 255, Pan: 0, Tilt: 128, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 				},
// 			},
// 		},
// 		{
// 			name:         "four fixtures, circle pattern - 8 point , with shift of 2 ie 1/2",
// 			fixtures:     4,
// 			shift:        2,
// 			scannerState: allFixturesEnabled,
// 			Coordinates: []Coordinate{
// 				{
// 					Tilt: 0,
// 					Pan:  0,
// 				},
// 				{
// 					Tilt: 1,
// 					Pan:  1,
// 				},
// 				{
// 					Tilt: 2,
// 					Pan:  2,
// 				},
// 				{
// 					Tilt: 3,
// 					Pan:  3,
// 				},
// 				{
// 					Tilt: 4,
// 					Pan:  4,
// 				},
// 				{
// 					Tilt: 5,
// 					Pan:  5,
// 				},
// 				{
// 					Tilt: 6,
// 					Pan:  6,
// 				},
// 				{
// 					Tilt: 7,
// 					Pan:  7,
// 				},
// 			},
// 			want: common.Pattern{
// 				Steps: []common.Step{
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 0, Shutter: 255, Pan: 0, Tilt: 0, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 0, Shutter: 255, Pan: 4, Tilt: 4, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 0, Shutter: 255, Pan: 0, Tilt: 0, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 0, Shutter: 255, Pan: 4, Tilt: 4, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}, Gobo: 0, Shutter: 255, Pan: 1, Tilt: 1, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 5, Tilt: 5, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}, Gobo: 0, Shutter: 255, Pan: 1, Tilt: 1, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 5, Tilt: 5, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}, Gobo: 0, Shutter: 255, Pan: 2, Tilt: 2, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 6, Tilt: 6, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}, Gobo: 0, Shutter: 255, Pan: 2, Tilt: 2, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 6, Tilt: 6, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Gobo: 0, Shutter: 255, Pan: 3, Tilt: 3, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 7, Tilt: 7, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Gobo: 0, Shutter: 255, Pan: 3, Tilt: 3, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 7, Tilt: 7, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 0, Shutter: 255, Pan: 4, Tilt: 4, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 0, Shutter: 255, Pan: 0, Tilt: 0, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 0, Shutter: 255, Pan: 4, Tilt: 4, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 0, Shutter: 255, Pan: 0, Tilt: 0, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 5, Tilt: 5, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}, Gobo: 0, Shutter: 255, Pan: 1, Tilt: 1, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 5, Tilt: 5, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}, Gobo: 0, Shutter: 255, Pan: 1, Tilt: 1, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 6, Tilt: 6, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}, Gobo: 0, Shutter: 255, Pan: 2, Tilt: 2, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 6, Tilt: 6, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}, Gobo: 0, Shutter: 255, Pan: 2, Tilt: 2, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 7, Tilt: 7, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Gobo: 0, Shutter: 255, Pan: 3, Tilt: 3, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 7, Tilt: 7, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Gobo: 0, Shutter: 255, Pan: 3, Tilt: 3, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 				},
// 			},
// 		},
// 		{
// 			name:         "four fixtures, circle pattern - 8 point , with shift of 3 ie 3/4 shift",
// 			fixtures:     4,
// 			shift:        3,
// 			scannerState: allFixturesEnabled,
// 			Coordinates: []Coordinate{
// 				{
// 					Tilt: 0,
// 					Pan:  0,
// 				},
// 				{
// 					Tilt: 1,
// 					Pan:  1,
// 				},
// 				{
// 					Tilt: 2,
// 					Pan:  2,
// 				},
// 				{
// 					Tilt: 3,
// 					Pan:  3,
// 				},
// 				{
// 					Tilt: 4,
// 					Pan:  4,
// 				},
// 				{
// 					Tilt: 5,
// 					Pan:  5,
// 				},
// 				{
// 					Tilt: 6,
// 					Pan:  6,
// 				},
// 				{
// 					Tilt: 7,
// 					Pan:  7,
// 				},
// 			},
// 			want: common.Pattern{
// 				Steps: []common.Step{
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 0, Shutter: 255, Pan: 0, Tilt: 0, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 6, Tilt: 6, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 0, Shutter: 255, Pan: 4, Tilt: 4, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}, Gobo: 0, Shutter: 255, Pan: 2, Tilt: 2, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}, Gobo: 0, Shutter: 255, Pan: 1, Tilt: 1, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 7, Tilt: 7, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 5, Tilt: 5, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Gobo: 0, Shutter: 255, Pan: 3, Tilt: 3, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}, Gobo: 0, Shutter: 255, Pan: 2, Tilt: 2, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 0, Shutter: 255, Pan: 0, Tilt: 0, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 6, Tilt: 6, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 0, Shutter: 255, Pan: 4, Tilt: 4, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Gobo: 0, Shutter: 255, Pan: 3, Tilt: 3, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}, Gobo: 0, Shutter: 255, Pan: 1, Tilt: 1, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 7, Tilt: 7, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 5, Tilt: 5, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 0, Shutter: 255, Pan: 4, Tilt: 4, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}, Gobo: 0, Shutter: 255, Pan: 2, Tilt: 2, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 0, Shutter: 255, Pan: 0, Tilt: 0, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 6, Tilt: 6, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 5, Tilt: 5, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Gobo: 0, Shutter: 255, Pan: 3, Tilt: 3, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}, Gobo: 0, Shutter: 255, Pan: 1, Tilt: 1, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 7, Tilt: 7, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 6, Tilt: 6, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 0, Shutter: 255, Pan: 4, Tilt: 4, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}, Gobo: 0, Shutter: 255, Pan: 2, Tilt: 2, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 0, Shutter: 255, Pan: 0, Tilt: 0, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 					{
// 						Fixtures: map[int]common.Fixture{
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 7, Tilt: 7, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 0, Shutter: 255, Pan: 5, Tilt: 5, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Gobo: 0, Shutter: 255, Pan: 3, Tilt: 3, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 							{MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}, Gobo: 0, Shutter: 255, Pan: 1, Tilt: 1, ScannerColor: common.Color{R: 255, G: 255, B: 255, W: 0, A: 0, UV: 0}},
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := GeneratePattern(tt.Coordinates, tt.fixtures, tt.shift, tt.chase, allFixturesEnabled); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Got = %+v", got)
// 				t.Errorf("Want = %+v", tt.want)
// 			}
// 		})
// 	}
// }

// func TestScanGenerateSineWave(t *testing.T) {
// 	type args struct {
// 		size      float64
// 		frequency float64
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		wantOut []Coordinate
// 	}{
// 		{
// 			name: "5000hz sawtooth",
// 			args: args{
// 				size:      255,
// 				frequency: 5000,
// 			},
// 			wantOut: []Coordinate{
// 				{156, 1},
// 				{273, 26},
// 				{352, 52},
// 				{229, 77},
// 				{159, 103},
// 				{286, 128},
// 				{348, 154},
// 				{217, 179},
// 				{163, 205},
// 				{298, 230},
// 				{343, 256},
// 				{205, 281},
// 				{169, 307},
// 				{310, 332},
// 				{336, 358},
// 				{194, 383},
// 				{177, 409},
// 				{320, 434},
// 				{328, 460},
// 				{184, 485},
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if gotOut := ScanGenerateSineWave(tt.args.size, tt.args.frequency, 10); !reflect.DeepEqual(gotOut, tt.wantOut) {
// 				t.Errorf("ScanGenerateSineWave() = %v, want %v", gotOut, tt.wantOut)
// 			}
// 		})
// 	}
// }

// func Test_getEnabledScanner(t *testing.T) {
// 	type args struct {
// 		scannerState          map[int]common.ScannerState
// 		numberCoordinates     int
// 		numberEnabledScanners int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want []int
// 	}{
// 		{
// 			name: "test 8 scanners 3 disabled",
// 			args: args{
// 				scannerState: map[int]common.ScannerState{
// 					0: {Enabled: false, Inverted: false},
// 					1: {Enabled: false, Inverted: false},
// 					2: {Enabled: false, Inverted: false},
// 					3: {Enabled: true, Inverted: false},
// 					4: {Enabled: true, Inverted: false},
// 					5: {Enabled: true, Inverted: false},
// 					6: {Enabled: true, Inverted: false},
// 					7: {Enabled: true, Inverted: false},
// 				},
// 				numberCoordinates:     64,
// 				numberEnabledScanners: 5,
// 			},
// 			want: []int{3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
// 				4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
// 				5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5,
// 				6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
// 				7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := makeEnabledScannerList(tt.args.scannerState, tt.args.numberCoordinates, tt.args.numberEnabledScanners, 8); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("getEnabledScanner() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestGenerateStandardChasePatterm(t *testing.T) {
	type args struct {
		numberSteps  int
		scannerState map[int]common.ScannerState
	}
	tests := []struct {
		name string
		args args
		want common.Pattern
	}{
		{
			name: "standard 8 steps fixture 2 & 4 disabled",
			args: args{
				numberSteps: 8,
				scannerState: map[int]common.ScannerState{
					0: {
						Enabled: true,
					},
					1: {
						Enabled: false, // Disabled.
					},
					2: {
						Enabled: true,
					},
					3: {
						Enabled: false,
					},
					4: {
						Enabled: false,
					},
					5: {
						Enabled: false,
					},
					6: {
						Enabled: false,
					},
					7: {
						Enabled: false,
					},
				},
			},
			want: common.Pattern{
				Name:  "Chase",
				Label: "std.Scanner.Chase",
				Steps: []common.Step{
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Colors: []common.Color{{R: 255, G: 255, B: 255}}},
							1: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							2: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							3: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							5: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							6: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							7: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						},
					},
					// {
					// 	Fixtures: map[int]common.Fixture{
					// 		0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// 		1: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 255, Colors: []common.Color{{R: 255, G: 255, B: 255}}},
					// 		2: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// 		3: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// 		4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// 		5: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// 		6: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// 		7: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// 	},
					// },
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							1: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							2: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Colors: []common.Color{{R: 255, G: 255, B: 255}}},
							3: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							5: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							6: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							7: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						},
					},
					// {
					// 	Fixtures: map[int]common.Fixture{
					// 		0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// 		1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// 		2: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// 		3: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 255, Colors: []common.Color{{R: 255, G: 255, B: 255}}},
					// 		4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// 		5: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// 		6: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// 		7: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// 	},
					// },
					// {
					// 	Fixtures: map[int]common.Fixture{
					// 		0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// 		1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// 		2: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// 		3: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// 		4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 255, Colors: []common.Color{{R: 255, G: 255, B: 255}}},
					// 		5: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// 		6: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// 		7: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// 	},
					// },
					// {
					// 	Fixtures: map[int]common.Fixture{
					// 		0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// 		1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// 		2: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// 		3: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// 		4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// 		5: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 255, Colors: []common.Color{{R: 255, G: 255, B: 255}}},
					// 		6: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// 		7: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// 	},
					// },
					// {
					// 	Fixtures: map[int]common.Fixture{
					// 		0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// 		1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// 		2: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// 		3: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// 		4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// 		5: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// 		6: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 255, Colors: []common.Color{{R: 255, G: 255, B: 255}}},
					// 		7: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// 	},
					// },
					// {
					// 	Fixtures: map[int]common.Fixture{
					// 		0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// 		1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// 		2: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// 		3: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// 		4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// 		5: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// 		6: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// 		7: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 255, Colors: []common.Color{{R: 255, G: 255, B: 255}}},
					// 	},
					// },
				},
			},
		},
		// {
		// 	name: "standard 2 steps All enabled.",
		// 	args: args{
		// 		numberSteps: 2,
		// 		scannerState: map[int]common.ScannerState{
		// 			0: {
		// 				Enabled: true,
		// 			},
		// 			1: {
		// 				Enabled: true,
		// 			},
		// 		},
		// 	},
		// 	want: common.Pattern{
		// 		Name:  "Chase",
		// 		Label: "std.Scanner.Chase",
		// 		Steps: []common.Step{
		// 			{
		// 				Fixtures: map[int]common.Fixture{
		// 					0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Colors: []common.Color{{R: 255, G: 255, B: 255}}},
		// 					1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 				},
		// 			},
		// 			{
		// 				Fixtures: map[int]common.Fixture{
		// 					0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 					1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Colors: []common.Color{{R: 255, G: 255, B: 255}}},
		// 				},
		// 			},
		// 		},
		// 	},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateStandardChasePatterm(tt.args.numberSteps, tt.args.scannerState); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateStandardChasePatterm() got = %+v", got)
				t.Errorf("GenerateStandardChasePatterm() want %+v", tt.want)
			}
		})
	}
}
