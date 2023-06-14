// Copyright (C) 2022,2023 dhowlett99.
// This is the dmxlights position calculator test code.
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

package position

import (
	"reflect"
	"testing"

	"github.com/dhowlett99/dmxlights/pkg/common"
)

var allFixturesEnabled = map[int]common.FixtureState{
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

func TestCalculatePositions(t *testing.T) {

	var full = 255

	tests := []struct {
		name     string
		steps    []common.Step
		sequence common.Sequence
		scanner  bool
		want     map[int][]common.FixtureBuffer
		want1    int
	}{
		{
			name: "Standard 3 fixtures forward chase.",

			scanner: false,
			sequence: common.Sequence{
				Bounce:                false,
				ScannerInvert:         false,
				FadeUp:                []int{0, 50, 255},
				FadeDown:              []int{255, 50, 0},
				Optimisation:          false,
				FixtureState:          allFixturesEnabled,
				EnabledNumberFixtures: 3,
				ScannerChaser:         false,
				RGBShift:              0,
			},
			steps: []common.Step{
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
			},
			want: map[int][]common.FixtureBuffer{
				0: { // Fixture 1 Fade up and down.
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 1 OFF
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 1 OFF
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
				},

				1: {
					// Fixture 2 OFF
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 2 Fade up and down.
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 2 OFF
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
				},

				2: {
					// Fixture 3 OFF
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 3 OFF
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 2 Fade up and down.
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
				},
			},
			want1: 3,
		},
		{
			name:    "Test two colors - flash",
			scanner: false,
			sequence: common.Sequence{
				Bounce:                false,
				ScannerInvert:         false,
				FadeUp:                []int{0, 50, 255},
				FadeDown:              []int{255, 50, 0},
				Optimisation:          false,
				FixtureState:          allFixturesEnabled,
				EnabledNumberFixtures: 3,
				ScannerChaser:         false,
				RGBShift:              0,
			},
			steps: []common.Step{
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 255, B: 255}}},
						1: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 255, B: 255}}},
						2: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 255, B: 255}}},
						3: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 255, B: 255}}},
						4: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 255, B: 255}}},
						5: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 255, B: 255}}},
						6: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 255, B: 255}}},
						7: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 255, B: 255}}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						3: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						4: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						5: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						6: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						7: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fadeColors, got1 := CalculatePositions(tt.steps, tt.sequence, tt.scanner)
			if !reflect.DeepEqual(fadeColors, tt.want) {
				t.Errorf("CalculatePositions() got = %+v, want %+v", fadeColors, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CalculatePositions() got1 = %+v, want %+v", got1, tt.want1)
			}
		})
	}
}

func TestAssemblePositions(t *testing.T) {

	tests := []struct {
		name           string
		fadeColors     map[int][]common.FixtureBuffer
		numberFixtures int
		want           map[int]common.Position
	}{
		{
			name:           "golden path assemble fade colors into positions",
			numberFixtures: 3,
			fadeColors: map[int][]common.FixtureBuffer{
				0: { // Fixture 1 Fade up and down.
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 1 OFF
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 1 OFF
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
				},

				1: {
					// Fixture 2 OFF
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 2 Fade up and down.
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 2 OFF
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
				},

				2: {
					// Fixture 3 OFF
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 3 OFF
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 2 Fade up and down.
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
				},
			},

			want: map[int]common.Position{
				0: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},

				1: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				2: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				3: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				4: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				5: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				6: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				7: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				8: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				9: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				10: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				11: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				12: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				13: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				14: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				15: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				16: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				17: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := AssemblePositions(tt.fadeColors, tt.numberFixtures); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AssemblePositions() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
