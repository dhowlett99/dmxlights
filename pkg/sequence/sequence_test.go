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
	"reflect"
	"testing"

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
)

// func Test_calculateScannerPositions(t *testing.T) {

// 	full := 255
// 	type args struct {
// 		steps  []common.Step
// 		bounce bool
// 		invert bool
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want map[int][]common.Position
// 	}{
// 		{
// 			name: "golden path - common par fixture RGB",
// 			args: args{
// 				slopeOn:  []int{1, 50, 255},
// 				slopeOff: []int{1, 50, 255},
// 				sequence: common.Sequence{
// 					Bounce:   false,
// 					Invert:   false,
// 					RGBShift: 0,
// 					RGBSize:  255,
// 					RGBFade:  0,
// 					Steps: []common.Step{
// 						{
// 							Fixtures: []common.Fixture{
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 							},
// 						},
// 						{
// 							Fixtures: []common.Fixture{
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 							},
// 						},
// 						{
// 							Fixtures: []common.Fixture{
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 							},
// 						},
// 						{
// 							Fixtures: []common.Fixture{
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 							},
// 						},
// 						{
// 							Fixtures: []common.Fixture{
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 							},
// 						},
// 						{
// 							Fixtures: []common.Fixture{
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 							},
// 						},
// 						{
// 							Fixtures: []common.Fixture{
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 							},
// 						},
// 						{
// 							Fixtures: []common.Fixture{
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
// 							},
// 						},
// 						{
// 							Fixtures: []common.Fixture{
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 							},
// 						},
// 					},
// 				},
// 			},
// 			want: map[int]common.Position{

// 				0: {
// 					Fixtures: map[int]common.Fixture{
// 						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 1, B: 0}}},
// 						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 					},
// 				},

// 				0: {
// 					{
// 						ScannerNumber: 0,
// 						StartPosition: 0,
// 						Color:         common.Color{R: 0, G: 255, B: 0},
// 					},
// 				},
// 				14: {
// 					{
// 						ScannerNumber: 1,
// 						StartPosition: 14,
// 						Color:         common.Color{R: 0, G: 255, B: 0},
// 					},
// 				},
// 				28: {
// 					{
// 						ScannerNumber: 2,
// 						StartPosition: 28,
// 						Color:         common.Color{R: 0, G: 255, B: 0},
// 					},
// 				},
// 				42: {
// 					{
// 						ScannerNumber: 3,
// 						StartPosition: 42,
// 						Color:         common.Color{R: 0, G: 255, B: 0},
// 					},
// 				},
// 				56: {
// 					{
// 						ScannerNumber: 4,
// 						StartPosition: 56,
// 						Color:         common.Color{R: 0, G: 255, B: 0},
// 					},
// 				},
// 				70: {
// 					{
// 						ScannerNumber: 5,
// 						StartPosition: 70,
// 						Color:         common.Color{R: 0, G: 255, B: 0},
// 					},
// 				},
// 				84: {
// 					{
// 						ScannerNumber: 6,
// 						StartPosition: 84,
// 						Color:         common.Color{R: 0, G: 255, B: 0},
// 					},
// 				},

// 				98: {{
// 					ScannerNumber: 7,
// 					StartPosition: 98,
// 					Color:         common.Color{R: 0, G: 255, B: 0},
// 				},
// 				},
// 			},
// 		},
// 		{
// 			name: "multiple colors case",
// 			args: args{
// 				steps: []common.Step{
// 					{
// 						Fixtures: []common.Fixture{
// 							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}, {R: 0, G: 255, B: 0}, {R: 0, G: 0, B: 255}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 						},
// 					},
// 					{
// 						Fixtures: []common.Fixture{
// 							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}, {R: 0, G: 255, B: 0}, {R: 0, G: 0, B: 255}}},
// 						},
// 					},
// 				},
// 			},
// 			want: map[int][]common.Position{
// 				0: {
// 					{
// 						ScannerNumber: 0,
// 						StartPosition: 0,
// 						Color:         common.Color{R: 255, G: 0, B: 0},
// 					},
// 				},
// 				14: {
// 					{
// 						ScannerNumber: 0,
// 						StartPosition: 14,
// 						Color:         common.Color{R: 0, G: 255, B: 0},
// 					},
// 				},
// 				28: {
// 					{
// 						ScannerNumber: 0,
// 						StartPosition: 28,
// 						Color:         common.Color{R: 0, G: 0, B: 255},
// 					},
// 				},
// 				42: {
// 					{
// 						ScannerNumber: 1,
// 						StartPosition: 42,
// 						Color:         common.Color{R: 255, G: 0, B: 0},
// 					},
// 				},
// 				56: {
// 					{
// 						ScannerNumber: 1,
// 						StartPosition: 56,
// 						Color:         common.Color{R: 0, G: 255, B: 0},
// 					},
// 				},
// 				70: {
// 					{
// 						ScannerNumber: 1,
// 						StartPosition: 70,
// 						Color:         common.Color{R: 0, G: 0, B: 255},
// 					},
// 				},
// 			},
// 		},
// 		{
// 			name: "Scanner case, both scanners doing same things.",
// 			args: args{
// 				bounce: false,
// 				steps: []common.Step{
// 					{
// 						Type: "scanner",
// 						Fixtures: []common.Fixture{
// 							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 255},
// 							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 255},
// 						},
// 					},
// 					{
// 						Type: "scanner",
// 						Fixtures: []common.Fixture{
// 							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 50, Tilt: 200},
// 							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 50, Tilt: 200},
// 						},
// 					},
// 					{
// 						Type: "scanner",
// 						Fixtures: []common.Fixture{
// 							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 100, Tilt: 150},
// 							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 100, Tilt: 150},
// 						},
// 					},
// 					{
// 						Type: "scanner",
// 						Fixtures: []common.Fixture{
// 							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 36, Shutter: 255, Pan: 150, Tilt: 100},
// 							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 36, Shutter: 255, Pan: 150, Tilt: 100},
// 						},
// 					},
// 					{
// 						Type: "scanner",
// 						Fixtures: []common.Fixture{
// 							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 200, Tilt: 50},
// 							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 200, Tilt: 50},
// 						},
// 					},
// 					{
// 						Type: "scanner",
// 						Fixtures: []common.Fixture{
// 							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 190, B: 255}}, Gobo: 36, Shutter: 255, Pan: 255, Tilt: 0},
// 							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 190, B: 255}}, Gobo: 36, Shutter: 255, Pan: 255, Tilt: 0},
// 						},
// 					},
// 					{
// 						Type: "scanner",
// 						Fixtures: []common.Fixture{
// 							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 145, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 0},
// 							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 145, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 0},
// 						},
// 					},
// 				},
// 			},
// 			want: map[int][]common.Position{
// 				0: {
// 					{
// 						ScannerNumber: 0,
// 						StartPosition: 0,
// 						Color:         common.Color{R: 0, G: 0, B: 255},
// 						Gobo:          36, Shutter: 255, Pan: 0, Tilt: 255,
// 					},
// 					{
// 						ScannerNumber: 1,
// 						StartPosition: 0,
// 						Color:         common.Color{R: 0, G: 0, B: 255},
// 						Gobo:          36, Shutter: 255, Pan: 0, Tilt: 255,
// 					},
// 				},
// 				14: {
// 					{
// 						ScannerNumber: 0,
// 						StartPosition: 14,
// 						Color:         common.Color{R: 255, G: 0, B: 0},
// 						Gobo:          36, Shutter: 255, Pan: 50, Tilt: 200,
// 					},
// 					{
// 						ScannerNumber: 1,
// 						StartPosition: 14,
// 						Color:         common.Color{R: 255, G: 0, B: 0},
// 						Gobo:          36, Shutter: 255, Pan: 50, Tilt: 200,
// 					},
// 				},
// 				28: {
// 					{
// 						ScannerNumber: 0,
// 						StartPosition: 28,
// 						Color:         common.Color{R: 255, G: 0, B: 255},
// 						Gobo:          36,
// 						Shutter:       255, Pan: 100, Tilt: 150,
// 					},
// 					{
// 						ScannerNumber: 1,
// 						StartPosition: 28,
// 						Color:         common.Color{R: 255, G: 0, B: 255},
// 						Gobo:          36,
// 						Shutter:       255, Pan: 100, Tilt: 150,
// 					},
// 				},
// 				42: {
// 					{
// 						ScannerNumber: 0,
// 						StartPosition: 42,
// 						Color:         common.Color{R: 0, G: 255, B: 255},
// 						Gobo:          36,
// 						Shutter:       255, Pan: 150, Tilt: 100,
// 					},
// 					{
// 						ScannerNumber: 1,
// 						StartPosition: 42,
// 						Color:         common.Color{R: 0, G: 255, B: 255},
// 						Gobo:          36,
// 						Shutter:       255, Pan: 150, Tilt: 100,
// 					},
// 				},
// 				56: {
// 					{
// 						ScannerNumber: 0,
// 						StartPosition: 56,
// 						Color:         common.Color{R: 100, G: 0, B: 255},
// 						Gobo:          36, Shutter: 255, Pan: 200, Tilt: 50,
// 					},
// 					{
// 						ScannerNumber: 1,
// 						StartPosition: 56,
// 						Color:         common.Color{R: 100, G: 0, B: 255},
// 						Gobo:          36, Shutter: 255, Pan: 200, Tilt: 50,
// 					},
// 				},
// 				70: {
// 					{
// 						ScannerNumber: 0,
// 						StartPosition: 70,
// 						Color:         common.Color{R: 0, G: 190, B: 255},
// 						Gobo:          36, Shutter: 255, Pan: 255, Tilt: 0,
// 					},
// 					{
// 						ScannerNumber: 1,
// 						StartPosition: 70,
// 						Color:         common.Color{R: 0, G: 190, B: 255},
// 						Gobo:          36, Shutter: 255, Pan: 255, Tilt: 0,
// 					},
// 				},
// 				84: {
// 					{
// 						ScannerNumber: 0,
// 						StartPosition: 84,
// 						Color:         common.Color{R: 145, G: 0, B: 255},
// 						Gobo:          36, Shutter: 255, Pan: 0, Tilt: 0,
// 					},
// 					{
// 						ScannerNumber: 1,
// 						StartPosition: 84,
// 						Color:         common.Color{R: 145, G: 0, B: 255},
// 						Gobo:          36, Shutter: 255, Pan: 0, Tilt: 0,
// 					},
// 				},
// 			},
// 		},
// 		{
// 			name: "Scanner case, both scanners doing different things.",
// 			args: args{
// 				bounce: false,
// 				steps: []common.Step{
// 					{
// 						Type: "scanner",
// 						Fixtures: []common.Fixture{
// 							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 255},
// 							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 128, Tilt: 255},
// 						},
// 					},
// 					{
// 						Type: "scanner",
// 						Fixtures: []common.Fixture{
// 							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 50, Tilt: 200},
// 							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 128, Tilt: 200},
// 						},
// 					},
// 				},
// 			},
// 			want: map[int][]common.Position{
// 				0: {
// 					{
// 						ScannerNumber: 0,
// 						StartPosition: 0,
// 						Color:         common.Color{R: 0, G: 0, B: 255},
// 						Gobo:          36, Shutter: 255, Pan: 0, Tilt: 255,
// 					},
// 					{
// 						ScannerNumber: 1,
// 						StartPosition: 0,
// 						Color:         common.Color{R: 0, G: 0, B: 255},
// 						Gobo:          36, Shutter: 255, Pan: 128, Tilt: 255,
// 					},
// 				},
// 				14: {
// 					{
// 						ScannerNumber: 0,
// 						StartPosition: 14,
// 						Color:         common.Color{R: 255, G: 0, B: 0},
// 						Gobo:          36, Shutter: 255, Pan: 50, Tilt: 200,
// 					},
// 					{
// 						ScannerNumber: 1,
// 						StartPosition: 14,
// 						Color:         common.Color{R: 255, G: 0, B: 0},
// 						Gobo:          36, Shutter: 255, Pan: 128, Tilt: 200,
// 					},
// 				},
// 			},
// 		},
// 		{
// 			name: "Scanner case, one set of instruction in a pattern should create one set of positions.",
// 			args: args{
// 				bounce: false,
// 				steps: []common.Step{
// 					{
// 						Type: "scanner",
// 						Fixtures: []common.Fixture{
// 							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 255},
// 						},
// 					},
// 					{
// 						Type: "scanner",
// 						Fixtures: []common.Fixture{
// 							{Type: "scanner", MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 50, Tilt: 200},
// 						},
// 					},
// 				},
// 			},
// 			want: map[int][]common.Position{
// 				0: {
// 					{
// 						ScannerNumber: 0,
// 						StartPosition: 0,
// 						Color:         common.Color{R: 0, G: 0, B: 255},
// 						Gobo:          36, Shutter: 255, Pan: 0, Tilt: 255,
// 					},
// 				},
// 				14: {
// 					{
// 						ScannerNumber: 0,
// 						StartPosition: 14,
// 						Color:         common.Color{R: 255, G: 0, B: 0},
// 						Gobo:          36, Shutter: 255, Pan: 50, Tilt: 200,
// 					},
// 				},
// 			},
// 		},
// 		{
// 			name: "Pairs case",
// 			args: args{
// 				bounce: false,
// 				steps: []common.Step{
// 					{
// 						Fixtures: []common.Fixture{
// 							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 						},
// 					},
// 					{
// 						Fixtures: []common.Fixture{
// 							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
// 						},
// 					},
// 				},
// 			},
// 			want: map[int][]common.Position{

// 				0: {
// 					{
// 						ScannerNumber: 0,
// 						StartPosition: 0,
// 						Color:         common.Color{R: 255, G: 0, B: 0},
// 						Gobo:          0, Shutter: 0, Pan: 0, Tilt: 0,
// 					},
// 					{
// 						ScannerNumber: 2,
// 						StartPosition: 0,
// 						Color:         common.Color{R: 255, G: 0, B: 0},
// 						Gobo:          0, Shutter: 0, Pan: 0, Tilt: 0,
// 					},
// 					{
// 						ScannerNumber: 4,
// 						StartPosition: 0,
// 						Color:         common.Color{R: 255, G: 0, B: 0},
// 						Gobo:          0, Shutter: 0, Pan: 0, Tilt: 0,
// 					},
// 					{
// 						ScannerNumber: 6,
// 						StartPosition: 0,
// 						Color:         common.Color{R: 255, G: 0, B: 0},
// 						Gobo:          0, Shutter: 0, Pan: 0, Tilt: 0,
// 					},
// 				},
// 				14: {
// 					{
// 						ScannerNumber: 1,
// 						StartPosition: 14,
// 						Color:         common.Color{R: 255, G: 0, B: 0},
// 						Gobo:          0, Shutter: 0, Pan: 0, Tilt: 0,
// 					},
// 					{
// 						ScannerNumber: 3,
// 						StartPosition: 14,
// 						Color:         common.Color{R: 255, G: 0, B: 0},
// 						Gobo:          0, Shutter: 0, Pan: 0, Tilt: 0,
// 					},
// 					{
// 						ScannerNumber: 5,
// 						StartPosition: 14,
// 						Color:         common.Color{R: 255, G: 0, B: 0},
// 						Gobo:          0, Shutter: 0, Pan: 0, Tilt: 0,
// 					},
// 					{
// 						ScannerNumber: 7,
// 						StartPosition: 14,
// 						Color:         common.Color{R: 255, G: 0, B: 0},
// 						Gobo:          0, Shutter: 0, Pan: 0, Tilt: 0,
// 					},
// 				},
// 			},
// 		},
// 		{
// 			name: "Scanners inverted no bounce",
// 			args: args{
// 				bounce: false,
// 				invert: true,
// 				steps: []common.Step{
// 					{
// 						Fixtures: []common.Fixture{
// 							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Pan: 0, Tilt: 0, Gobo: 1, Shutter: 255},
// 						},
// 					},
// 					{
// 						Fixtures: []common.Fixture{
// 							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Pan: 1, Tilt: 1, Gobo: 1, Shutter: 255},
// 						},
// 					},
// 					{
// 						Fixtures: []common.Fixture{
// 							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Pan: 2, Tilt: 2, Gobo: 1, Shutter: 255},
// 						},
// 					},
// 				},
// 			},
// 			want: map[int][]common.Position{
// 				0:  {{ScannerNumber: 0, StartPosition: 0, Color: common.Color{R: 0, G: 0, B: 255}, Gobo: 1, Shutter: 255, Pan: 2, Tilt: 2}},
// 				14: {{ScannerNumber: 0, StartPosition: 14, Color: common.Color{R: 0, G: 255, B: 0}, Gobo: 1, Shutter: 255, Pan: 1, Tilt: 1}},
// 				28: {{ScannerNumber: 0, StartPosition: 28, Color: common.Color{R: 255, G: 0, B: 0}, Gobo: 1, Shutter: 255, Pan: 0, Tilt: 0}},
// 			},
// 		},
// 		{
// 			name: "Scanners inverted with bounce",
// 			args: args{
// 				bounce: true,
// 				invert: true,
// 				steps: []common.Step{
// 					{
// 						Fixtures: []common.Fixture{
// 							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Pan: 0, Tilt: 0, Gobo: 1, Shutter: 255},
// 						},
// 					},
// 					{
// 						Fixtures: []common.Fixture{
// 							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Pan: 1, Tilt: 1, Gobo: 1, Shutter: 255},
// 						},
// 					},
// 					{
// 						Fixtures: []common.Fixture{
// 							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Pan: 2, Tilt: 2, Gobo: 1, Shutter: 255},
// 						},
// 					},
// 				},
// 			},
// 			want: map[int][]common.Position{
// 				0:  {{ScannerNumber: 0, StartPosition: 0, Color: common.Color{R: 0, G: 0, B: 255}, Gobo: 1, Shutter: 255, Pan: 2, Tilt: 2}},
// 				14: {{ScannerNumber: 0, StartPosition: 14, Color: common.Color{R: 0, G: 255, B: 0}, Gobo: 1, Shutter: 255, Pan: 1, Tilt: 1}},
// 				28: {{ScannerNumber: 0, StartPosition: 28, Color: common.Color{R: 255, G: 0, B: 0}, Gobo: 1, Shutter: 255, Pan: 0, Tilt: 0}},
// 				42: {{ScannerNumber: 0, StartPosition: 42, Color: common.Color{R: 255, G: 0, B: 0}, Gobo: 1, Shutter: 255, Pan: 0, Tilt: 0}},
// 				56: {{ScannerNumber: 0, StartPosition: 56, Color: common.Color{R: 0, G: 255, B: 0}, Gobo: 1, Shutter: 255, Pan: 1, Tilt: 1}},
// 				70: {{ScannerNumber: 0, StartPosition: 70, Color: common.Color{R: 0, G: 0, B: 255}, Gobo: 1, Shutter: 255, Pan: 2, Tilt: 2}},
// 			},
// 		},
// 		{
// 			name: "Multicolored Patten",
// 			args: args{
// 				bounce: false,
// 				invert: false,
// 				steps: []common.Step{
// 					{
// 						Fixtures: []common.Fixture{
// 							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}},
// 						},
// 					},
// 					{
// 						Fixtures: []common.Fixture{
// 							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
// 						},
// 					},
// 					{
// 						Fixtures: []common.Fixture{
// 							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}},
// 						},
// 					},
// 					{
// 						Fixtures: []common.Fixture{
// 							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}},
// 						},
// 					},
// 					{
// 						Fixtures: []common.Fixture{
// 							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
// 						},
// 					},
// 					{
// 						Fixtures: []common.Fixture{
// 							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}},
// 						},
// 					},
// 					{
// 						Fixtures: []common.Fixture{

// 							{MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}},
// 							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
// 						},
// 					},
// 					{
// 						Fixtures: []common.Fixture{

//								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}},
//								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
//								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}},
//								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}},
//								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
//								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}},
//								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
//								{MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}},
//							},
//						},
//					},
//				},
//				want: map[int][]common.Position{
//					0: {
//						{ScannerNumber: 0, StartPosition: 0, Color: common.Color{R: 255, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 1, StartPosition: 0, Color: common.Color{R: 255, G: 111, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 2, StartPosition: 0, Color: common.Color{R: 255, G: 255, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 3, StartPosition: 0, Color: common.Color{R: 0, G: 255, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 4, StartPosition: 0, Color: common.Color{R: 0, G: 255, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 5, StartPosition: 0, Color: common.Color{R: 0, G: 0, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 6, StartPosition: 0, Color: common.Color{R: 100, G: 0, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 7, StartPosition: 0, Color: common.Color{R: 255, G: 0, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//					},
//					14: {
//						{ScannerNumber: 0, StartPosition: 14, Color: common.Color{R: 255, G: 111, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 1, StartPosition: 14, Color: common.Color{R: 255, G: 255, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 2, StartPosition: 14, Color: common.Color{R: 0, G: 255, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 3, StartPosition: 14, Color: common.Color{R: 0, G: 255, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 4, StartPosition: 14, Color: common.Color{R: 0, G: 0, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 5, StartPosition: 14, Color: common.Color{R: 100, G: 0, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 6, StartPosition: 14, Color: common.Color{R: 255, G: 0, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 7, StartPosition: 14, Color: common.Color{R: 255, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//					},
//					28: {
//						{ScannerNumber: 0, StartPosition: 28, Color: common.Color{R: 255, G: 255, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 1, StartPosition: 28, Color: common.Color{R: 0, G: 255, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 2, StartPosition: 28, Color: common.Color{R: 0, G: 255, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 3, StartPosition: 28, Color: common.Color{R: 0, G: 0, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 4, StartPosition: 28, Color: common.Color{R: 100, G: 0, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 5, StartPosition: 28, Color: common.Color{R: 255, G: 0, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 6, StartPosition: 28, Color: common.Color{R: 255, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 7, StartPosition: 28, Color: common.Color{R: 255, G: 111, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//					},
//					42: {
//						{ScannerNumber: 0, StartPosition: 42, Color: common.Color{R: 0, G: 255, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 1, StartPosition: 42, Color: common.Color{R: 0, G: 255, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 2, StartPosition: 42, Color: common.Color{R: 0, G: 0, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 3, StartPosition: 42, Color: common.Color{R: 100, G: 0, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 4, StartPosition: 42, Color: common.Color{R: 255, G: 0, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 5, StartPosition: 42, Color: common.Color{R: 255, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 6, StartPosition: 42, Color: common.Color{R: 255, G: 111, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 7, StartPosition: 42, Color: common.Color{R: 255, G: 255, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//					},
//					56: {
//						{ScannerNumber: 0, StartPosition: 56, Color: common.Color{R: 0, G: 255, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 1, StartPosition: 56, Color: common.Color{R: 0, G: 0, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 2, StartPosition: 56, Color: common.Color{R: 100, G: 0, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 3, StartPosition: 56, Color: common.Color{R: 255, G: 0, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 4, StartPosition: 56, Color: common.Color{R: 255, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 5, StartPosition: 56, Color: common.Color{R: 255, G: 111, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 6, StartPosition: 56, Color: common.Color{R: 255, G: 255, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 7, StartPosition: 56, Color: common.Color{R: 0, G: 255, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//					},
//					70: {
//						{ScannerNumber: 0, StartPosition: 70, Color: common.Color{R: 0, G: 0, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 1, StartPosition: 70, Color: common.Color{R: 100, G: 0, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 2, StartPosition: 70, Color: common.Color{R: 255, G: 0, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 3, StartPosition: 70, Color: common.Color{R: 255, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 4, StartPosition: 70, Color: common.Color{R: 255, G: 111, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 5, StartPosition: 70, Color: common.Color{R: 255, G: 255, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 6, StartPosition: 70, Color: common.Color{R: 0, G: 255, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 7, StartPosition: 70, Color: common.Color{R: 0, G: 255, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//					},
//					84: {
//						{ScannerNumber: 0, StartPosition: 84, Color: common.Color{R: 100, G: 0, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 1, StartPosition: 84, Color: common.Color{R: 255, G: 0, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 2, StartPosition: 84, Color: common.Color{R: 255, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 3, StartPosition: 84, Color: common.Color{R: 255, G: 111, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 4, StartPosition: 84, Color: common.Color{R: 255, G: 255, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 5, StartPosition: 84, Color: common.Color{R: 0, G: 255, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 6, StartPosition: 84, Color: common.Color{R: 0, G: 255, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 7, StartPosition: 84, Color: common.Color{R: 0, G: 0, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//					},
//					98: {
//						{ScannerNumber: 0, StartPosition: 98, Color: common.Color{R: 255, G: 0, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 1, StartPosition: 98, Color: common.Color{R: 255, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 2, StartPosition: 98, Color: common.Color{R: 255, G: 111, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 3, StartPosition: 98, Color: common.Color{R: 255, G: 255, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 4, StartPosition: 98, Color: common.Color{R: 0, G: 255, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 5, StartPosition: 98, Color: common.Color{R: 0, G: 255, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 6, StartPosition: 98, Color: common.Color{R: 0, G: 0, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//						{ScannerNumber: 7, StartPosition: 98, Color: common.Color{R: 100, G: 0, B: 255}, Pan: 0, Tilt: 0, Shutter: 0, Gobo: 0},
//					},
//				},
//			},
//		}
//		for _, tt := range tests {
//			t.Run(tt.name, func(t *testing.T) {
//				if got, _ := calculatePositions("scanner", tt.args.slopeOn, tt.args.slopeOff, false); !reflect.DeepEqual(got, tt.want) {
//					t.Errorf("got = %+v", got)
//					t.Errorf("want =%+v", tt.want)
//				}
//			})
//		}
//	}
func Test_getNumberOfFixtures(t *testing.T) {

	type args struct {
		sequenceNumber int
		fixtures       *fixture.Fixtures
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "eight fixtures",
			args: args{
				sequenceNumber: 0,
				fixtures: &fixture.Fixtures{
					Fixtures: []fixture.Fixture{

						{Name: "fixture1", Number: 1, Group: 1},
						{Name: "fixture2", Number: 2, Group: 1},
						{Name: "fixture3", Number: 3, Group: 1},
						{Name: "fixture4", Number: 4, Group: 1},
						{Name: "fixture5", Number: 5, Group: 1},
						{Name: "fixture6", Number: 6, Group: 1},
						{Name: "fixture7", Number: 7, Group: 1},
						{Name: "fixture8", Number: 8, Group: 1},

						{Name: "fixture1", Number: 1, Group: 2},
						{Name: "fixture2", Number: 2, Group: 2},
						{Name: "fixture3", Number: 3, Group: 2},

						{Name: "fixture1", Number: 1, Group: 3},
						{Name: "fixture2", Number: 2, Group: 3},
						{Name: "fixture3", Number: 3, Group: 3},
						{Name: "fixture4", Number: 4, Group: 3},

						{Name: "fixture1", Number: 1, Group: 4},
						{Name: "fixture2", Number: 2, Group: 4},
					},
				},
			},
			want: 8,
		},

		{
			name: "three fixtures",
			args: args{
				sequenceNumber: 1,
				fixtures: &fixture.Fixtures{
					Fixtures: []fixture.Fixture{

						{Name: "fixture1", Number: 1, Group: 1},
						{Name: "fixture2", Number: 2, Group: 1},
						{Name: "fixture3", Number: 3, Group: 1},
						{Name: "fixture4", Number: 4, Group: 1},
						{Name: "fixture5", Number: 5, Group: 1},
						{Name: "fixture6", Number: 6, Group: 1},
						{Name: "fixture7", Number: 7, Group: 1},
						{Name: "fixture8", Number: 8, Group: 1},

						{Name: "fixture1", Number: 1, Group: 2},
						{Name: "fixture2", Number: 2, Group: 2},
						{Name: "fixture3", Number: 3, Group: 2},

						{Name: "fixture1", Number: 1, Group: 3},
						{Name: "fixture2", Number: 2, Group: 3},
						{Name: "fixture3", Number: 3, Group: 3},
						{Name: "fixture4", Number: 4, Group: 3},

						{Name: "fixture1", Number: 1, Group: 4},
						{Name: "fixture2", Number: 2, Group: 4},
					},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNumberOfFixtures(tt.args.sequenceNumber, tt.args.fixtures); got != tt.want {
				t.Errorf("getNumberOfFixtures() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func Test_replaceRGBcolorsInSteps(t *testing.T) {

	full := 255
	type args struct {
		steps  []common.Step
		colors []common.Color
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
						Fixtures: []common.Fixture{
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						},
					},
					{
						Fixtures: []common.Fixture{
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						},
					},
					{
						Fixtures: []common.Fixture{
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
						},
					},
				},
				colors: []common.Color{
					{R: 255, G: 0, B: 0},
					{R: 0, G: 255, B: 0},
					{R: 0, G: 0, B: 255},
				},
			},
			want: []common.Step{
				{
					Fixtures: []common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				{
					Fixtures: []common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				{
					Fixtures: []common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
					},
				},
			},
		},
		{
			name: "replace a number of colors with just one.",
			args: args{
				steps: []common.Step{
					{
						Fixtures: []common.Fixture{
							{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						},
					},
					{
						Fixtures: []common.Fixture{
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						},
					},
					{
						Fixtures: []common.Fixture{
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
						},
					},
				},
				colors: []common.Color{
					{R: 0, G: 255, B: 0},
				},
			},
			want: []common.Step{
				{
					Fixtures: []common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				{
					Fixtures: []common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				{
					Fixtures: []common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
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

func Test_invertRGBColors(t *testing.T) {

	full := 255
	type args struct {
		steps  []common.Step
		colors []common.Color
	}
	tests := []struct {
		name string
		args args
		want []common.Step
	}{
		{
			name: "invert a single color.",
			args: args{
				steps: []common.Step{
					{
						Fixtures: []common.Fixture{
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						},
					},
					{
						Fixtures: []common.Fixture{
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						},
					},
					{
						Fixtures: []common.Fixture{
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}},
						},
					},
				},
				colors: []common.Color{
					{R: 0, G: 255, B: 0},
				},
			},
			want: []common.Step{
				{
					Fixtures: []common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					},
				},
				{
					Fixtures: []common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					},
				},
				{
					Fixtures: []common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := invertRGBColors(tt.args.steps, tt.args.colors); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("invertRGBColors() = %v, want %v", got, tt.want)
			}
		})
	}
}
