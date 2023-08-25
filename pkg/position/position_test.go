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
	"fmt"
	"reflect"
	"testing"

	"github.com/dhowlett99/dmxlights/pkg/common"
)

const full int = 255

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

var allFixturesRGBInverted = map[int]common.FixtureState{
	0: {
		Enabled:     true,
		RGBInverted: true,
	},
	1: {
		Enabled:     true,
		RGBInverted: true,
	},
	2: {
		Enabled:     true,
		RGBInverted: true,
	},
	3: {
		Enabled:     true,
		RGBInverted: true,
	},
	4: {
		Enabled:     true,
		RGBInverted: true,
	},
	5: {
		Enabled:     true,
		RGBInverted: true,
	},
	6: {
		Enabled:     true,
		RGBInverted: true,
	},
	7: {
		Enabled:     true,
		RGBInverted: true,
	},
}

func TestCalculateRGB3FixturesPositions(t *testing.T) {

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
				ScannerReverse:        false,
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
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fadeColors, got1, _ := CalculatePositions(tt.steps, tt.sequence, tt.scanner)

			if !reflect.DeepEqual(fadeColors, tt.want) {

				t.Errorf("CalculatePositions() ")

				fmt.Printf("++++++++++++++ WANT ++++++++++++++++++++\n")
				for fixtureNumber := 0; fixtureNumber < len(tt.want); fixtureNumber++ {

					fade := tt.want[fixtureNumber]

					fmt.Printf("fixtureNumber:%d ============================\n", fixtureNumber)

					for number, fixtureBuffer := range fade {
						fmt.Printf("Buffer:%d fixtureBuffer:%+v\n", number, fixtureBuffer.Color.R)
					}

				}

				fmt.Printf("++++++++++++++ GOT ++++++++++++++++++++\n")
				for fixtureNumber := 0; fixtureNumber < len(fadeColors); fixtureNumber++ {

					fade := fadeColors[fixtureNumber]

					fmt.Printf("fixtureNumber:%d ============================\n", fixtureNumber)

					step := 0
					for _, fixtureBuffer := range fade {
						fmt.Printf("\tstep %d rule %d %s: fixtureBuffer:%+v\n", fixtureBuffer.Step, fixtureBuffer.Rule, fixtureBuffer.DebugMsg, fixtureBuffer.Color.R)
						step++
					}

				}
			}

			if got1 != tt.want1 {
				t.Errorf("CalculatePositions() got1 = %+v, want %+v", got1, tt.want1)
			}
		})
	}
}

func TestCalculateRGBPositionsWithShift(t *testing.T) {

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
			name: "Standard 3 fixtures forward chase with shift.",

			scanner: false,
			sequence: common.Sequence{
				Bounce:                false,
				ScannerReverse:        false,
				FadeUp:                []int{0, 50, 255},
				FadeDown:              []int{255, 50, 0},
				Optimisation:          false,
				FixtureState:          allFixturesEnabled,
				EnabledNumberFixtures: 3,
				ScannerChaser:         false,
				RGBShift:              9,
			},
			steps: []common.Step{
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
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
					// Fixture 1 OFF
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
				},

				1: {
					// Fixture 2 OFF
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
				},

				2: {
					// Fixture 3 OFF
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 3 OFF
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fadeColors, got1, _ := CalculatePositions(tt.steps, tt.sequence, tt.scanner)

			if !reflect.DeepEqual(fadeColors, tt.want) {

				t.Errorf("CalculatePositions() ")

				// fmt.Printf("++++++++++++++ WANT ++++++++++++++++++++\n")
				for fixtureNumber := 0; fixtureNumber < len(tt.want); fixtureNumber++ {

					fade := tt.want[fixtureNumber]

					fmt.Printf("fixtureNumber:%d ============================\n", fixtureNumber)

					for number, fixtureBuffer := range fade {
						fmt.Printf("Buffer:%d fixtureBuffer:%+v\n", number, fixtureBuffer.Color.R)
					}

				}

				fmt.Printf("++++++++++++++ GOT ++++++++++++++++++++\n")
				for fixtureNumber := 0; fixtureNumber < len(fadeColors); fixtureNumber++ {

					fade := fadeColors[fixtureNumber]

					fmt.Printf("fixtureNumber:%d ============================\n", fixtureNumber)

					step := 0
					for _, fixtureBuffer := range fade {
						fmt.Printf("\tstep %d rule %d %s: fixtureBuffer:%+v\n", fixtureBuffer.Step, fixtureBuffer.Rule, fixtureBuffer.DebugMsg, fixtureBuffer.Color.R)
						step++
					}

				}
			}

			if got1 != tt.want1 {
				t.Errorf("CalculatePositions() got1 = %+v, want %+v", got1, tt.want1)
			}
		})
	}
}

func TestCalculatRGBBouncePositions(t *testing.T) {

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
				Bounce:                true,
				ScannerReverse:        false,
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
					StepNumber: 0,
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
					},
				},
				{
					StepNumber: 1,
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
					},
				},
				{
					StepNumber: 2,
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
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
					// Fixture 1 Fade up and down.
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
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
					// Fixture 2 Fade up and down.
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
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
				},
			},
			want1: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fadeColors, got1, _ := CalculatePositions(tt.steps, tt.sequence, tt.scanner)

			if !reflect.DeepEqual(fadeColors, tt.want) {

				t.Errorf("CalculatePositions() ")

				fmt.Printf("++++++++++++++ WANT ++++++++++++++++++++\n")
				for fixtureNumber := 0; fixtureNumber < len(tt.want); fixtureNumber++ {

					fade := tt.want[fixtureNumber]

					fmt.Printf("fixtureNumber:%d ============================\n", fixtureNumber)

					for _, fixtureBuffer := range fade {
						fmt.Printf("\tstep %d rule %d %s: fixtureBuffer:%+v\n", fixtureBuffer.Step, fixtureBuffer.Rule, fixtureBuffer.DebugMsg, fixtureBuffer.Color.R)
					}
				}

				fmt.Printf("++++++++++++++ GOT ++++++++++++++++++++\n")
				for fixtureNumber := 0; fixtureNumber < len(fadeColors); fixtureNumber++ {

					fade := fadeColors[fixtureNumber]

					fmt.Printf("fixtureNumber:%d ============================\n", fixtureNumber)

					for _, fixtureBuffer := range fade {
						fmt.Printf("\tstep %d rule %d %s: fixtureBuffer:%+v\n", fixtureBuffer.Step, fixtureBuffer.Rule, fixtureBuffer.DebugMsg, fixtureBuffer.Color.R)
					}

				}
			}

			if got1 != tt.want1 {
				t.Errorf("CalculatePositions() got1 = %+v, want %+v", got1, tt.want1)
			}
		})
	}
}

func TestApplyRGBChaseWithOnlyThreeEnabled(t *testing.T) {

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
			name: "Standard 8 fixtures but only theree enabled forward chase.",

			scanner: false,
			sequence: common.Sequence{
				Type:           "rgb",
				Bounce:         false,
				ScannerReverse: false,
				FadeUp:         []int{0, 50, 255},
				FadeDown:       []int{255, 50, 0},
				Optimisation:   true,
				FixtureState: map[int]common.FixtureState{
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
				EnabledNumberFixtures: 3,
				ScannerChaser:         false,
				RGBShift:              0,
			},
			steps: []common.Step{
				{
					StepNumber: 0,
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
					},
				},
				{
					StepNumber: 1,
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
					},
				},
				{
					StepNumber: 2,
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
					},
				},
				{
					StepNumber: 3,
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
					},
				},
				{
					StepNumber: 4,
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
					},
				},
				{
					StepNumber: 5,
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
					},
				},
				{
					StepNumber: 6,
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
					},
				},
				{
					StepNumber: 7,
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
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
				3: {
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},

					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},

					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
				},
				4: {
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},

					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},

					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
				},
				5: {
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},

					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},

					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
				},
				6: {
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},

					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},

					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
				},
				7: {
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},

					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},

					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
				},
			},
			want1: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			pattern := common.Pattern{}
			pattern.Steps = tt.steps

			for _, step := range pattern.Steps {
				fmt.Printf("Step %d\n", step.StepNumber)
				for fixtureNumber := 0; fixtureNumber < len(step.Fixtures); fixtureNumber++ {
					fixture := step.Fixtures[fixtureNumber]
					fmt.Printf("Fixture %d Color %+v\n", fixture.Number, fixture.Color)
				}
			}

			RGBPattern := ApplyFixtureState(pattern, tt.sequence.FixtureState)
			fadeColors, got1, _ := CalculatePositions(RGBPattern.Steps, tt.sequence, tt.scanner)

			if !reflect.DeepEqual(fadeColors, tt.want) {

				t.Errorf("CalculatePositions() ")

				fmt.Printf("++++++++++++++ WANT ++++++++++++++++++++\n")
				for fixtureNumber := 0; fixtureNumber < len(tt.want); fixtureNumber++ {

					fade := tt.want[fixtureNumber]

					fmt.Printf("fixtureNumber:%d ============================\n", fixtureNumber)

					for _, fixtureBuffer := range fade {
						fmt.Printf("\tstep %d rule %d %s: fixtureBuffer:%+v\n", fixtureBuffer.Step, fixtureBuffer.Rule, fixtureBuffer.DebugMsg, fixtureBuffer.Color.R)
					}
				}

				fmt.Printf("++++++++++++++ GOT ++++++++++++++++++++\n")
				for fixtureNumber := 0; fixtureNumber < len(fadeColors); fixtureNumber++ {

					fade := fadeColors[fixtureNumber]

					fmt.Printf("fixtureNumber:%d ============================\n", fixtureNumber)

					for _, fixtureBuffer := range fade {
						fmt.Printf("\tstep %d rule %d %s: fixtureBuffer:%+v\n", fixtureBuffer.Step, fixtureBuffer.Rule, fixtureBuffer.DebugMsg, fixtureBuffer.Color.R)
					}

				}
			}

			if got1 != tt.want1 {
				t.Errorf("CalculatePositions() got1 = %+v, want %+v", got1, tt.want1)
			}
		})
	}
}

func TestCalculateRGBNoBounceRGBInvertedPositions(t *testing.T) {

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
			name: "Standard 3 fixtures forward chase inverted no bounce.",

			scanner: false,
			sequence: common.Sequence{
				Type:                  "rgb",
				Bounce:                false,
				ScannerReverse:        false,
				FadeUp:                []int{0, 50, 255},
				FadeDown:              []int{255, 50, 0},
				Optimisation:          false,
				FixtureState:          allFixturesRGBInverted,
				EnabledNumberFixtures: 3,
				ScannerChaser:         false,
				RGBShift:              0,
			},
			steps: []common.Step{
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
					},
				},
			},
			want: map[int][]common.FixtureBuffer{
				0: { // Fixture 0 Fade down.
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 0 Fade Up.
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 0 On
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 0 On
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
				},

				1: {
					// Fixture 1 On
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 1 Fade down.
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 1 Fade Up.
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 1 On
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
				},

				2: {
					// Fixture 2 On
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 2 On
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},

					// Fixture 2 Fade down.
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 2 Fade Up.
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
				},
			},
			want1: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fadeColors, got1, _ := CalculatePositions(tt.steps, tt.sequence, tt.scanner)

			if !reflect.DeepEqual(fadeColors, tt.want) {

				t.Errorf("CalculatePositions() ")

				fmt.Printf("++++++++++++++ WANT ++++++++++++++++++++\n")
				for fixtureNumber := 0; fixtureNumber < len(tt.want); fixtureNumber++ {

					fade := tt.want[fixtureNumber]

					fmt.Printf("fixtureNumber:%d ============================\n", fixtureNumber)

					for number, fixtureBuffer := range fade {
						fmt.Printf("Buffer:%d fixtureBuffer:%+v\n", number, fixtureBuffer.Color.R)
					}

				}

				fmt.Printf("++++++++++++++ GOT ++++++++++++++++++++\n")
				for fixtureNumber := 0; fixtureNumber < len(fadeColors); fixtureNumber++ {

					fade := fadeColors[fixtureNumber]

					fmt.Printf("fixtureNumber:%d ============================\n", fixtureNumber)

					step := 0
					for number, fixtureBuffer := range fade {
						fmt.Printf("\tstep %d rule %d %s: Buffer:%d fixtureBuffer:%+v\n", fixtureBuffer.Step, fixtureBuffer.Rule, fixtureBuffer.DebugMsg, number, fixtureBuffer.Color.R)
						step++
					}

				}
			}

			if got1 != tt.want1 {
				t.Errorf("CalculatePositions() got1 = %+v, want %+v", got1, tt.want1)
			}
		})
	}
}

func TestAssemblePositions(t *testing.T) {

	tests := []struct {
		name               string
		fadeColors         map[int][]common.FixtureBuffer
		numberFixtures     int
		totalNumberOfSteps int
		want               map[int]common.Position
	}{
		{
			name:               "golden path assemble fade colors into positions",
			numberFixtures:     4,
			totalNumberOfSteps: 24,
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
					// Fixture 3 Fade up and down.
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 3 OFF
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
				},

				3: {
					// Fixture 4 OFF
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 4 OFF
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 4 OFF
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 4 Fade up and down.
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
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},

				1: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				2: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				3: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				4: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				5: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				6: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				7: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				8: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				9: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				10: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				11: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				12: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				13: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				14: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				15: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				16: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				17: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},

				18: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				19: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				20: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				21: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				22: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				23: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if positions, _ := AssemblePositions(tt.fadeColors, tt.numberFixtures, tt.totalNumberOfSteps, false); !reflect.DeepEqual(positions, tt.want) {
				t.Errorf("AssemblePositions() = %+v, want %+v", positions, tt.want)

				for positionNumber := 0; positionNumber < len(positions); positionNumber++ {

					position := positions[positionNumber]

					fmt.Printf("Position %d\n", positionNumber)

					for fixtureNumber := 0; fixtureNumber < len(position.Fixtures); fixtureNumber++ {

						fixture := position.Fixtures[fixtureNumber]

						fmt.Printf("fixture %+v\n", fixture.Color)

					}
				}

			}
		})
	}
}

func TestAssemblePositionsOnlyThreeEnabled(t *testing.T) {

	tests := []struct {
		name               string
		fadeColors         map[int][]common.FixtureBuffer
		numberFixtures     int
		totalNumberOfSteps int
		optimasation       bool
		want               map[int]common.Position
	}{
		{
			name:               "golden path assemble fade colors into positions",
			optimasation:       false,
			numberFixtures:     3,
			totalNumberOfSteps: 24,
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
				3: {
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},

					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},

					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
				},
				4: {
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},

					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},

					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
				},
				5: {
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},

					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},

					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
				},
				6: {
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},

					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},

					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
				},
				7: {
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},

					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},

					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
				},
			},
			want: map[int]common.Position{
				0: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},

				1: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				2: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				3: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				4: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				5: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				6: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				7: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				8: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				9: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				10: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				11: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				12: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				13: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				14: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				15: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				16: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				17: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if positions, _ := AssemblePositions(tt.fadeColors, tt.numberFixtures, tt.totalNumberOfSteps, tt.optimasation); !reflect.DeepEqual(positions, tt.want) {
				t.Errorf("AssemblePositions() = %+v, want %+v", positions, tt.want)

				for positionNumber := 0; positionNumber < len(positions); positionNumber++ {

					position := positions[positionNumber]

					fmt.Printf("Position %d\n", positionNumber)

					for fixtureNumber := 0; fixtureNumber < tt.numberFixtures; fixtureNumber++ {

						fixture := position.Fixtures[fixtureNumber]

						fmt.Printf("\t fixture %d %+v\n", fixture.Number, fixture.Color)

					}
				}

			}
		})
	}
}

func TestCalculateRGBPositionsSimpleGreenChase8Fitures(t *testing.T) {

	full := 255
	type args struct {
		sequence common.Sequence
		scanner  bool
	}
	tests := []struct {
		name  string
		args  args
		steps []common.Step
		want  map[int]common.Position
		want1 int
	}{
		{
			name: "golden path - 8 fixtures simple green chase. Should result in 48 steps",
			args: args{
				scanner: false,
				sequence: common.Sequence{
					FadeUp:                []int{1, 50, 255},
					FadeDown:              []int{255, 50, 1},
					Optimisation:          false,
					Bounce:                false,
					RGBInvert:             false,
					RGBShift:              0,
					FixtureState:          allFixturesEnabled,
					EnabledNumberFixtures: 8,
					ScannerChaser:         false,
				},
			},
			steps: []common.Step{
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
			},

			want: map[int]common.Position{
				0: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				1: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				2: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				3: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				4: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				5: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				6: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				7: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				8: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				9: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				10: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				11: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				12: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				13: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				14: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				15: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				16: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				17: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				18: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				19: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				20: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				21: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				22: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				23: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				24: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				25: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				26: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				27: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				28: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				29: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				30: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				31: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				32: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				33: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				34: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				35: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				36: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				37: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				38: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				39: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				40: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				41: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				42: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				43: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				44: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				45: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				46: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				47: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: common.Color{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
			},
			want1: 48,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			fadeColors, numberFixtures, totalNumberOfSteps := CalculatePositions(tt.steps, tt.args.sequence, tt.args.scanner)
			positions, numberPositions := AssemblePositions(fadeColors, numberFixtures, totalNumberOfSteps, false)
			if !reflect.DeepEqual(positions, tt.want) {
				t.Errorf("CalculatePositions() got = %v, want %v", positions, tt.want)

				for positionNumber := 0; positionNumber < len(positions); positionNumber++ {

					fmt.Printf("Position:%d ============================\n", positionNumber)
					position := positions[positionNumber]
					for fixtureNumber := 0; fixtureNumber < len(position.Fixtures); fixtureNumber++ {

						fixture := position.Fixtures[fixtureNumber]
						fmt.Printf("Fixture:%d color:%+v\n", fixtureNumber, fixture.Color)
					}

				}

				fmt.Printf("++++++++++++++ GOT ++++++++++++++++++++\n")
				for fixtureNumber := 0; fixtureNumber < len(fadeColors); fixtureNumber++ {
					positionNumber := 0
					fade := fadeColors[fixtureNumber]

					fmt.Printf("fixtureNumber:%d ============================\n", fixtureNumber)

					step := 0
					for _, fixtureBuffer := range fade {
						fmt.Printf("\tpositionNumber %d step %d rule %d %s: fixtureBuffer:%+v\n", positionNumber, fixtureBuffer.Step, fixtureBuffer.Rule, fixtureBuffer.DebugMsg, fixtureBuffer.Color)
						step++
						positionNumber++
					}

				}
			}
			if numberPositions != tt.want1 {
				t.Errorf("CalculatePositions() got1 = %v, want %v", numberPositions, tt.want1)
			}
		})
	}
}

func TestCalculateRGBMulticoloredPatten(t *testing.T) {

	full := 255
	type args struct {
		sequence common.Sequence
		scanner  bool
	}
	tests := []struct {
		name  string
		args  args
		steps []common.Step
		want  map[int]common.Position
		want1 int
	}{
		{
			name: "Multicolored Patten",
			args: args{
				scanner: false,
				sequence: common.Sequence{
					Bounce:         false,
					ScannerReverse: false,
					FadeUp:         []int{255},
					//FadeDown:              []int{0},
					Optimisation:          false,
					FixtureState:          allFixturesEnabled,
					EnabledNumberFixtures: 8,
				},
			},
			steps: []common.Step{
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 111, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 255, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 255, B: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 255, B: 255}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 255}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 100, G: 0, B: 255}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 255}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 111, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 255, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 255, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 255, B: 255}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 255}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 100, G: 0, B: 255}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 255}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 255, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 255, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 255, B: 255}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 255}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 100, G: 0, B: 255}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 255}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 111, B: 0}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 255, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 255, B: 255}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 255}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 100, G: 0, B: 255}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 255}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 111, B: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 255, B: 0}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 255, B: 255}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 255}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 100, G: 0, B: 255}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 255}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 111, B: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 255, B: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 255, B: 0}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 255}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 100, G: 0, B: 255}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 255}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 111, B: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 255, B: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 255, B: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 255, B: 255}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{

						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 100, G: 0, B: 255}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 255}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 111, B: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 255, B: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 255, B: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 255, B: 255}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 255}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{

						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 255}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 111, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 255, B: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 255, B: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 255, B: 255}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 255}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 100, G: 0, B: 255}},
					},
				},
			},

			want: map[int]common.Position{
				// Start of want.
				0: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 111, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 255, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 100, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},

				1: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 111, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 255, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 100, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},

				2: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 255, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 100, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 111, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},

				3: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 255, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 100, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 111, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				4: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 255, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 100, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 111, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				5: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 100, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 111, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 255, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				6: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 100, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 111, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 255, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				7: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 111, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 255, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 100, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				// End of Want
			},
			want1: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			fadeColors, numberFixtures, totalNumberOfSteps := CalculatePositions(tt.steps, tt.args.sequence, true)
			positions, numberPositions := AssemblePositions(fadeColors, numberFixtures, totalNumberOfSteps, false)
			if !reflect.DeepEqual(positions, tt.want) {
				t.Errorf("CalculatePositions() got = %v, want %v", positions, tt.want)

				for positionNumber := 0; positionNumber < len(positions); positionNumber++ {

					fmt.Printf("Position:%d ============================\n", positionNumber)
					position := positions[positionNumber]
					for fixtureNumber := 0; fixtureNumber < len(position.Fixtures); fixtureNumber++ {

						fixture := position.Fixtures[fixtureNumber]
						fmt.Printf("Fixture:%d Color:%+v\n", fixtureNumber, fixture.Color)
					}

				}
			}
			if numberPositions != tt.want1 {
				t.Errorf("CalculatePositions() got1 = %v, want %v", numberPositions, tt.want1)
			}
		})
	}
}

func TestCalculateRGBShift8(t *testing.T) {

	full := 255
	type args struct {
		sequence common.Sequence
		scanner  bool
	}
	tests := []struct {
		name  string
		args  args
		steps []common.Step
		want  map[int]common.Position
		want1 int
	}{
		{
			name: "Shift 8 - common par fixture RGB",
			args: args{
				scanner: false,
				sequence: common.Sequence{
					FadeUp:                []int{1, 50, 255},
					FadeDown:              []int{255, 50, 1},
					Optimisation:          false,
					Bounce:                false,
					RGBInvert:             false,
					RGBShift:              8,
					FixtureState:          allFixturesEnabled,
					EnabledNumberFixtures: 4,
				},
			},
			steps: []common.Step{
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
			},

			want: map[int]common.Position{
				0: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				1: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				2: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				3: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				4: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				5: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				6: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				7: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				8: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
			},

			want1: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fadeColors, numberFixtures, totalNumberOfSteps := CalculatePositions(tt.steps, tt.args.sequence, true)
			positions, numberPositions := AssemblePositions(fadeColors, numberFixtures, totalNumberOfSteps, false)
			if !reflect.DeepEqual(positions, tt.want) {
				t.Errorf("CalculatePositions() got = %v, want %v", positions, tt.want)

				fmt.Printf(" ================== Want =====================\n")
				for positionNumber := 0; positionNumber < len(tt.want); positionNumber++ {

					fmt.Printf("Position:%d ============================\n", positionNumber)
					position := tt.want[positionNumber]
					for fixtureNumber := 0; fixtureNumber < len(position.Fixtures); fixtureNumber++ {

						fixture := position.Fixtures[fixtureNumber]
						fmt.Printf("Fixture:%d R:%d G:%d B:%d\n", fixtureNumber, fixture.Color.R, fixture.Color.G, fixture.Color.B)
					}

				}

				fmt.Printf(" ================== Got =====================\n")
				for positionNumber := 0; positionNumber < len(positions); positionNumber++ {

					fmt.Printf("Position:%d ============================\n", positionNumber)
					position := positions[positionNumber]
					for fixtureNumber := 0; fixtureNumber < len(position.Fixtures); fixtureNumber++ {

						fixture := position.Fixtures[fixtureNumber]
						fmt.Printf("Fixture:%d R:%d G:%d B:%d\n", fixtureNumber, fixture.Color.R, fixture.Color.G, fixture.Color.B)
					}

				}
			}
			if numberPositions != tt.want1 {
				t.Errorf("CalculatePositions() got1 = %v, want %v", numberPositions, tt.want1)
			}
		})
	}
}

func TestCalculateRGBShift1(t *testing.T) {

	full := 255
	type args struct {
		sequence common.Sequence
		scanner  bool
	}
	tests := []struct {
		name  string
		args  args
		steps []common.Step
		want  map[int]common.Position
		want1 int
	}{
		{
			name: "Shift1 - Not RGBInverted common par fixture RGB",
			args: args{
				scanner: false,
				sequence: common.Sequence{
					FadeUp:                []int{1, 50, 255},
					FadeDown:              []int{255, 50, 1},
					Optimisation:          false,
					Bounce:                false,
					RGBInvert:             false,
					RGBShift:              8, // Eight is reversed so creates a shift of 2.
					ScannerReverse:        false,
					FixtureState:          allFixturesEnabled,
					EnabledNumberFixtures: 4,
				},
			},
			steps: []common.Step{
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
			},

			want: map[int]common.Position{
				0: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				1: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				2: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				3: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				4: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				5: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				6: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				7: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
				8: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}},
					},
				},
			},
			want1: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Optimisation is turned off for testing.
			fadeColors, numberFixtures, totalNumberOfSteps := CalculatePositions(tt.steps, tt.args.sequence, true)
			positions, numberPositions := AssemblePositions(fadeColors, numberFixtures, totalNumberOfSteps, false)
			if !reflect.DeepEqual(positions, tt.want) {
				t.Errorf("calculatePositions() got = %+v, want %+v", positions, tt.want)
			}
			if numberPositions != tt.want1 {
				t.Errorf("calculatePositions() got1 = %+v, want %+v", numberPositions, tt.want1)
			}
		})
	}
}

func TestCalculateRGBPairsPatten(t *testing.T) {

	full := 255
	type args struct {
		sequence common.Sequence
		scanner  bool
	}
	tests := []struct {
		name  string
		args  args
		steps []common.Step
		want  map[int]common.Position
	}{
		{
			name: "Pairs case",
			args: args{
				scanner: false,
				sequence: common.Sequence{
					Bounce:                false,
					FadeUp:                []int{0, 50, 255},
					FadeDown:              []int{255, 50, 0},
					Optimisation:          false,
					FixtureState:          allFixturesEnabled,
					EnabledNumberFixtures: 8,
					ScannerChaser:         false,
					RGBShift:              0,
				},
			},
			steps: []common.Step{
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
					},
				},
			},

			want: map[int]common.Position{

				0: {
					Fixtures: map[int]common.Fixture{

						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				1: {
					Fixtures: map[int]common.Fixture{

						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				2: {
					Fixtures: map[int]common.Fixture{

						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				3: {
					Fixtures: map[int]common.Fixture{

						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				4: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				5: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				6: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				7: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				8: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				9: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				10: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				11: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Optimisation is turned off for testing.
			fadeColors, numberFixtures, totalNumberOfSteps := CalculatePositions(tt.steps, tt.args.sequence, tt.args.scanner)
			positions, _ := AssemblePositions(fadeColors, numberFixtures, totalNumberOfSteps, false)

			if !reflect.DeepEqual(positions, tt.want) {
				t.Errorf("calculatePositions() got = %+v, want %+v", positions, tt.want)

				fmt.Printf("++++++++++++++ GOT ++++++++++++++++++++\n")
				for positionNumnber := 0; positionNumnber < len(fadeColors); positionNumnber++ {

					position := positions[positionNumnber]
					fmt.Printf("positionNumnber:%d ============================\n", positionNumnber)

					for fixtureNumber := 0; fixtureNumber < len(position.Fixtures); fixtureNumber++ {
						fixture := position.Fixtures[fixtureNumber]
						fmt.Printf("fixture %d Color:%+v\n", fixtureNumber, fixture.Color)
					}
				}
			}
		})
	}
}

// Red, Green, Blue, Yellow Chase.
func TestCalculateMultiColorChasePatten(t *testing.T) {

	full := 255
	type args struct {
		sequence common.Sequence
		scanner  bool
	}
	tests := []struct {
		name  string
		args  args
		steps []common.Step
		want  map[int]common.Position
	}{
		{
			name: "RGB Chase case",
			args: args{
				scanner: false,
				sequence: common.Sequence{
					Bounce:                false,
					FadeUp:                []int{0, 50, 255},
					FadeDown:              []int{255, 50, 0},
					Optimisation:          false,
					FixtureState:          allFixturesEnabled,
					EnabledNumberFixtures: 8,
					ScannerChaser:         false,
					RGBShift:              0,
				},
			},
			steps: []common.Step{
				{
					Fixtures: map[int]common.Fixture{ // Red
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{ // Green
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 255, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 255, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 255, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 255, B: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 255, B: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 255, B: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 255, B: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 255, B: 0}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{ // Blue
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 255}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 255}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 255}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 255}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 255}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 255}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 255}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 255}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{ // Yellow
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 255, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 255, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 255, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 255, B: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 255, B: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 255, B: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 255, B: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 255, B: 0}},
					},
				},
			},

			want: map[int]common.Position{

				0: {
					Fixtures: map[int]common.Fixture{

						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				1: {
					Fixtures: map[int]common.Fixture{

						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				2: {
					Fixtures: map[int]common.Fixture{

						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				3: {
					Fixtures: map[int]common.Fixture{

						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				4: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				5: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				6: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				7: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				8: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				9: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				10: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				11: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				12: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				13: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 50, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 50, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 50, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 50, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 50, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 50, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 50, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 50, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				14: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				15: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				16: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 50, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 50, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 50, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 50, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 50, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 50, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 50, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 50, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				17: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				18: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				19: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 50, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 50, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 50, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 50, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 50, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 50, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 00, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 50, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 50, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				20: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				21: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				22: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 50, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 50, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 50, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 50, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 50, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 50, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 50, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 50, G: 50, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				23: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Optimisation is turned off for testing.
			fadeColors, numberFixtures, totalNumberOfSteps := CalculatePositions(tt.steps, tt.args.sequence, tt.args.scanner)
			fmt.Printf("++++++++++++++ GOT FADES ++++++++++++++++++++\n")
			for fixtureNumber := 0; fixtureNumber < len(fadeColors); fixtureNumber++ {

				fade := fadeColors[fixtureNumber]

				fmt.Printf("fixtureNumber:%d ============================\n", fixtureNumber)

				step := 0
				for number, fixtureBuffer := range fade {
					fmt.Printf("\tstep %d rule %d %s: Buffer:%d fixtureBuffer:%+v\n", fixtureBuffer.Step, fixtureBuffer.Rule, fixtureBuffer.DebugMsg, number, fixtureBuffer.Color)
					step++
				}

			}

			positions, _ := AssemblePositions(fadeColors, numberFixtures, totalNumberOfSteps, false)
			if !reflect.DeepEqual(positions, tt.want) {
				t.Errorf("CalculatePositions() got = %v, want %v", positions, tt.want)

				for positionNumber := 0; positionNumber < len(positions); positionNumber++ {

					fmt.Printf("Position:%d ============================\n", positionNumber)
					position := positions[positionNumber]
					for fixtureNumber := 0; fixtureNumber < len(position.Fixtures); fixtureNumber++ {

						fixture := position.Fixtures[fixtureNumber]
						fmt.Printf("Fixture:%d Color:%+v\n", fixtureNumber, fixture.Color)
					}

				}
			}
		})
	}
}

func Test_calculateScannerBothDoingSameThing(t *testing.T) {

	full := 255
	type args struct {
		sequence common.Sequence
		scanner  bool
	}
	tests := []struct {
		name  string
		args  args
		steps []common.Step
		want  map[int]common.Position
	}{

		{
			name: "Scanner case, both scanners doing same things.",
			args: args{
				scanner: true,
				sequence: common.Sequence{
					Type:                  "scanner",
					Bounce:                false,
					FadeUp:                []int{255},
					FadeDown:              []int{0},
					Optimisation:          false,
					ScannerChaser:         false,
					FixtureState:          allFixturesEnabled,
					EnabledNumberFixtures: 2,
				},
			},
			steps: []common.Step{
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 255}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 255},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 255}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 255},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}, Gobo: 36, Shutter: 255, Pan: 50, Tilt: 200},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}, Gobo: 36, Shutter: 255, Pan: 50, Tilt: 200},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 255}, Gobo: 36, Shutter: 255, Pan: 100, Tilt: 150},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 255}, Gobo: 36, Shutter: 255, Pan: 100, Tilt: 150},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 255, B: 255}, Gobo: 36, Shutter: 255, Pan: 150, Tilt: 100},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 255, B: 255}, Gobo: 36, Shutter: 255, Pan: 150, Tilt: 100},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 100, G: 0, B: 255}, Gobo: 36, Shutter: 255, Pan: 200, Tilt: 50},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 100, G: 0, B: 255}, Gobo: 36, Shutter: 255, Pan: 200, Tilt: 50},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 190, B: 255}, Gobo: 36, Shutter: 255, Pan: 255, Tilt: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 190, B: 255}, Gobo: 36, Shutter: 255, Pan: 255, Tilt: 0},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 145, G: 0, B: 255}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 145, G: 0, B: 255}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 0},
					},
				},
			},

			want: map[int]common.Position{
				0: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 255, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 255, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
					},
				},
				1: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 50, Tilt: 200, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 50, Tilt: 200, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
					},
				},
				2: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 100, Tilt: 150, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 100, Tilt: 150, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
					},
				},
				3: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 255, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 150, Tilt: 100, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 255, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 150, Tilt: 100, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
					},
				},
				4: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 100, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 200, Tilt: 50, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 100, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 200, Tilt: 50, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
					},
				},
				5: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 190, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 255, Tilt: 0, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 190, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 255, Tilt: 0, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
					},
				},
				6: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 145, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 145, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fadeColors, numberFixtures, totalNumberOfSteps := CalculatePositions(tt.steps, tt.args.sequence, true)
			positions, _ := AssemblePositions(fadeColors, numberFixtures, totalNumberOfSteps, false)
			if !reflect.DeepEqual(positions, tt.want) {
				t.Errorf("got = %+v", positions)
				t.Errorf("want =%+v", tt.want)

				fmt.Printf("++++++++++++++ GOT FADES ++++++++++++++++++++\n")
				for fixtureNumber := 0; fixtureNumber < len(fadeColors); fixtureNumber++ {
					positionNumber := 0
					fade := fadeColors[fixtureNumber]

					fmt.Printf("fixtureNumber:%d ============================\n", fixtureNumber)

					step := 0
					for _, fixtureBuffer := range fade {
						fmt.Printf("\tpositionNumber %d step %d rule %d %s: fixtureBuffer:%+v Pan:%d Tilt:%d\n", positionNumber, fixtureBuffer.Step, fixtureBuffer.Rule, fixtureBuffer.DebugMsg, fixtureBuffer.Color, fixtureBuffer.Pan, fixtureBuffer.Tilt)
						step++
						positionNumber++
					}
				}

				fmt.Printf("++++++++++++++ WANT ++++++++++++++++++++\n")
				for positionNumber := 0; positionNumber < len(tt.want); positionNumber++ {

					fmt.Printf("Position:%d ============================\n", positionNumber)
					position := tt.want[positionNumber]
					for fixtureNumber := 0; fixtureNumber < len(position.Fixtures); fixtureNumber++ {

						fixture := position.Fixtures[fixtureNumber]
						fmt.Printf("Fixture:%d color:%+v Pan %d Tilt %d\n", fixtureNumber, fixture.Color, fixture.Pan, fixture.Tilt)
					}
				}

				fmt.Printf("++++++++++++++ GOT ++++++++++++++++++++\n")
				for positionNumber := 0; positionNumber < len(positions); positionNumber++ {

					fmt.Printf("Position:%d ============================\n", positionNumber)
					position := positions[positionNumber]
					for fixtureNumber := 0; fixtureNumber < len(position.Fixtures); fixtureNumber++ {

						fixture := position.Fixtures[fixtureNumber]
						fmt.Printf("Fixture:%d color:%+v Pan %d Tilt %d\n", fixtureNumber, fixture.Color, fixture.Pan, fixture.Tilt)
					}
				}
			}
		})
	}
}

func Test_calculateScannerBothDoingDifferentThing(t *testing.T) {

	full := 255
	type args struct {
		sequence common.Sequence
		scanner  bool
	}
	tests := []struct {
		name  string
		args  args
		steps []common.Step
		want  map[int]common.Position
	}{

		{
			name: "Scanner case, both scanners doing different things.",
			args: args{
				scanner: true,
				sequence: common.Sequence{
					Bounce:                false,
					FadeUp:                []int{255},
					FadeDown:              []int{0},
					Optimisation:          false,
					FixtureState:          allFixturesEnabled,
					EnabledNumberFixtures: 2,
					ScannerChaser:         false,
				},
			},
			steps: []common.Step{
				{

					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 255}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 255},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 255}, Gobo: 36, Shutter: 255, Pan: 128, Tilt: 255},
					},
				},
				{

					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}, Gobo: 36, Shutter: 255, Pan: 50, Tilt: 200},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}, Gobo: 36, Shutter: 255, Pan: 128, Tilt: 200},
					},
				},
			},

			want: map[int]common.Position{
				0: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 255, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 128, Tilt: 255, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
					},
				},

				1: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 50, Tilt: 200, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 128, Tilt: 200, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
					},
				},
			},
		},

		// End of test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fadeColors, numberFixtures, totalNumberOfSteps := CalculatePositions(tt.steps, tt.args.sequence, true)
			positions, _ := AssemblePositions(fadeColors, numberFixtures, totalNumberOfSteps, false)
			if !reflect.DeepEqual(positions, tt.want) {
				t.Errorf("got = %+v", positions)
				t.Errorf("want =%+v", tt.want)
			}
		})
	}
}

func Test_calculateScannerSimpleCase(t *testing.T) {

	full := 255
	type args struct {
		sequence common.Sequence
		scanner  bool
	}
	tests := []struct {
		name  string
		args  args
		steps []common.Step
		want  map[int]common.Position
	}{

		{
			name: "Scanner case, one set of instruction in a pattern should create one set of positions.",
			args: args{
				scanner: true,
				sequence: common.Sequence{
					Bounce:                false,
					FadeUp:                []int{255},
					FadeDown:              []int{0},
					Optimisation:          false,
					FixtureState:          allFixturesEnabled,
					EnabledNumberFixtures: 1,
				},
			},
			steps: []common.Step{
				{

					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 255}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 255},
					},
				},
				{

					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}, Gobo: 36, Shutter: 255, Pan: 50, Tilt: 200},
					},
				},
			},

			want: map[int]common.Position{

				0: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 255, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
					},
				},
				1: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 50, Tilt: 200, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fadeColors, numberFixtures, totalNumberOfSteps := CalculatePositions(tt.steps, tt.args.sequence, true)
			positions, _ := AssemblePositions(fadeColors, numberFixtures, totalNumberOfSteps, false)
			if !reflect.DeepEqual(positions, tt.want) {
				t.Errorf("got = %+v", positions)
				t.Errorf("want =%+v", tt.want)
			}
		})
	}
}

func Test_calculateScannerPairsCase(t *testing.T) {

	full := 255
	type args struct {
		sequence common.Sequence
		scanner  bool
	}
	tests := []struct {
		name  string
		args  args
		steps []common.Step
		want  map[int]common.Position
	}{
		{
			name: "Pairs case",
			args: args{
				scanner: true,
				sequence: common.Sequence{
					Bounce:                false,
					FadeUp:                []int{255},
					FadeDown:              []int{0},
					Optimisation:          false,
					FixtureState:          allFixturesEnabled,
					EnabledNumberFixtures: 8,
					ScannerChaser:         false,
				},
			},
			steps: []common.Step{
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}},
					},
				},
			},

			want: map[int]common.Position{

				0: {
					Fixtures: map[int]common.Fixture{

						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				1: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fadeColors, numberFixtures, totalNumberOfSteps := CalculatePositions(tt.steps, tt.args.sequence, true)
			positions, _ := AssemblePositions(fadeColors, numberFixtures, totalNumberOfSteps, false)
			if !reflect.DeepEqual(positions, tt.want) {
				t.Errorf("got = %+v", positions)
				t.Errorf("want =%+v", tt.want)
			}
		})
	}
}

func Test_calculateScannerRGBInvertedCase(t *testing.T) {

	full := 255
	type args struct {
		sequence common.Sequence
		scanner  bool
	}
	tests := []struct {
		name  string
		args  args
		steps []common.Step
		want  map[int]common.Position
	}{

		{
			name: "Scanners inverted no bounce",
			args: args{
				scanner: true,
				sequence: common.Sequence{
					Bounce:                false,
					ScannerReverse:        true,
					FadeUp:                []int{255},
					FadeDown:              []int{0},
					Optimisation:          false,
					FixtureState:          allFixturesEnabled,
					EnabledNumberFixtures: 1,
					ScannerChaser:         false,
				},
			},
			steps: []common.Step{
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}, Pan: 0, Tilt: 0, Gobo: 1, Shutter: 255},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 255, B: 0}, Pan: 1, Tilt: 1, Gobo: 1, Shutter: 255},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 255}, Pan: 2, Tilt: 2, Gobo: 1, Shutter: 255},
					},
				},
			},

			want: map[int]common.Position{
				0: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}, Pan: 2, Tilt: 2, Shutter: 255, Rotate: 0, Music: 0, Gobo: 1, Program: 0},
					},
				},
				1: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 1, Tilt: 1, Shutter: 255, Rotate: 0, Music: 0, Gobo: 1, Program: 0},
					},
				},
				2: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Pan: 0, Tilt: 0, Shutter: 255, Rotate: 0, Music: 0, Gobo: 1, Program: 0},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fadeColors, numberFixtures, totalNumberOfSteps := CalculatePositions(tt.steps, tt.args.sequence, true)
			positions, _ := AssemblePositions(fadeColors, numberFixtures, totalNumberOfSteps, false)
			if !reflect.DeepEqual(positions, tt.want) {
				t.Errorf("got = %+v", positions)
				t.Errorf("want =%+v", tt.want)
			}
		})
	}
}

func Test_calculateScannerRGBInvertedBounceCase(t *testing.T) {

	full := 255
	type args struct {
		sequence common.Sequence
		scanner  bool
	}
	tests := []struct {
		name  string
		args  args
		steps []common.Step
		want  map[int]common.Position
	}{
		{
			name: "Scanners inverted with bounce",
			args: args{
				scanner: true,
				sequence: common.Sequence{
					Bounce:                true,
					ScannerReverse:        true,
					FadeUp:                []int{255},
					FadeDown:              []int{0},
					Optimisation:          false,
					FixtureState:          allFixturesEnabled,
					EnabledNumberFixtures: 1,
				},
			},
			steps: []common.Step{
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}, Pan: 0, Tilt: 0, Gobo: 1, Shutter: 255},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 255, B: 0}, Pan: 1, Tilt: 1, Gobo: 1, Shutter: 255},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 255}, Pan: 2, Tilt: 2, Gobo: 1, Shutter: 255},
					},
				},
			},

			want: map[int]common.Position{
				0: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 255}, Gobo: 1, Shutter: 255, Pan: 2, Tilt: 2},
					},
				},

				1: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 255, B: 0}, Gobo: 1, Shutter: 255, Pan: 1, Tilt: 1},
					},
				},
				2: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}, Gobo: 1, Shutter: 255, Pan: 0, Tilt: 0},
					},
				},
				3: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 255, G: 0, B: 0}, Gobo: 1, Shutter: 255, Pan: 0, Tilt: 0},
					},
				},
				4: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 255, B: 0}, Gobo: 1, Shutter: 255, Pan: 1, Tilt: 1},
					},
				},
				5: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Color{R: 0, G: 0, B: 255}, Gobo: 1, Shutter: 255, Pan: 2, Tilt: 2},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fadeColors, numberFixtures, totalNumberOfSteps := CalculatePositions(tt.steps, tt.args.sequence, true)
			positions, _ := AssemblePositions(fadeColors, numberFixtures, totalNumberOfSteps, false)
			if !reflect.DeepEqual(positions, tt.want) {
				t.Errorf("got = %+v", positions)
				t.Errorf("want =%+v", tt.want)
			}
		})
	}
}

// Which fixtures are inverted are controlled by the scanner state.
func Test_invertRGBColorsInSteps(t *testing.T) {

	full := 255
	type args struct {
		steps        []common.Step
		colors       []common.Color
		fixtureState map[int]common.FixtureState
	}
	tests := []struct {
		name string
		args args
		want []common.Step
	}{
		{
			name: "invert a single color.",
			args: args{
				fixtureState: allFixturesRGBInverted,
				steps: []common.Step{
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Color: common.Color{R: 0, G: 255, B: 0}},
							1: {MasterDimmer: full, Color: common.Color{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Color: common.Color{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Color: common.Color{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Color: common.Color{R: 0, G: 255, B: 0}},
							2: {MasterDimmer: full, Color: common.Color{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Color: common.Color{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Color: common.Color{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Color: common.Color{R: 0, G: 255, B: 0}},
						},
					},
				},
				colors: []common.Color{
					{R: 0, G: 255, B: 0},
				},
			},
			want: []common.Step{
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Color: common.Color{R: 0, G: 0, B: 0}, Inverted: true},
						1: {MasterDimmer: full, Color: common.Color{R: 0, G: 255, B: 0}, Inverted: true},
						2: {MasterDimmer: full, Color: common.Color{R: 0, G: 255, B: 0}, Inverted: true},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Color: common.Color{R: 0, G: 255, B: 0}, Inverted: true},
						1: {MasterDimmer: full, Color: common.Color{R: 0, G: 0, B: 0}, Inverted: true},
						2: {MasterDimmer: full, Color: common.Color{R: 0, G: 255, B: 0}, Inverted: true},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Color: common.Color{R: 0, G: 255, B: 0}, Inverted: true},
						1: {MasterDimmer: full, Color: common.Color{R: 0, G: 255, B: 0}, Inverted: true},
						2: {MasterDimmer: full, Color: common.Color{R: 0, G: 0, B: 0}, Inverted: true},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := invertRGBColorsInSteps(tt.args.steps, tt.args.colors, tt.args.fixtureState); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("invertRGBColors() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestApplyScannerState2And4Disabled(t *testing.T) {
	type args struct {
		steps        []common.Step
		scannerState map[int]common.FixtureState
	}
	tests := []struct {
		name string
		args args
		want common.Pattern
	}{
		{
			name: "standard 8 steps fixture 2 & 4 disabled",
			args: args{
				steps: []common.Step{
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: common.Color{R: 255, G: 255, B: 255}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: common.Color{R: 255, G: 255, B: 255}},
							2: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 255, Color: common.Color{R: 255, G: 255, B: 255}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: common.Color{R: 255, G: 255, B: 255}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 255, Color: common.Color{R: 255, G: 255, B: 255}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: common.Color{R: 255, G: 255, B: 255}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: common.Color{R: 255, G: 255, B: 255}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: common.Color{R: 255, G: 255, B: 255}},
						},
					},
				},
				scannerState: map[int]common.FixtureState{
					0: {
						Enabled: true,
					},
					1: {
						Enabled: true,
					},
					2: {
						Enabled: false, // Disabled.
					},
					3: {
						Enabled: true,
					},
					4: {
						Enabled: false, // Disabled.
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
				},
			},
			want: common.Pattern{

				Steps: []common.Step{
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: common.Color{R: 255, G: 255, B: 255}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: common.Color{R: 255, G: 255, B: 255}},
							2: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
						},
					},
					// {
					// 	Fixtures: map[int]common.Fixture{
					// 		0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
					// 		1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
					// 		2: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 255, Color: common.Color{R: 255, G: 255, B: 255}},
					// 		3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
					// 		4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
					// 		5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
					// 		6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
					// 		7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
					// 	},
					// },
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: common.Color{R: 255, G: 255, B: 255}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: common.Color{R: 255, G: 255, B: 255}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: common.Color{R: 255, G: 255, B: 255}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: common.Color{R: 255, G: 255, B: 255}},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		pattern := common.Pattern{}
		pattern.Steps = tt.args.steps
		t.Run(tt.name, func(t *testing.T) {
			if got := ApplyFixtureState(pattern, tt.args.scannerState); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateStandardChasePatterm() got = %+v", got)
				t.Errorf("GenerateStandardChasePatterm() want %+v", tt.want)
			}
		})
	}
}

func TestApplyScannerState4Enabled4Disabled(t *testing.T) {
	type args struct {
		steps        []common.Step
		scannerState map[int]common.FixtureState
	}
	tests := []struct {
		name string
		args args
		want common.Pattern
	}{
		{
			name: "4 enabled and 4 disabled.",
			args: args{
				steps: []common.Step{
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: common.Color{R: 255, G: 255, B: 255}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: common.Color{R: 255, G: 255, B: 255}},
							2: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: common.Color{R: 255, G: 255, B: 255}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: common.Color{R: 255, G: 255, B: 255}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 255, Color: common.Color{R: 255, G: 255, B: 255}},
							5: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 255, Color: common.Color{R: 255, G: 255, B: 255}},
							6: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 255, Color: common.Color{R: 255, G: 255, B: 255}},
							7: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 255, Color: common.Color{R: 255, G: 255, B: 255}},
						},
					},
				},
				scannerState: map[int]common.FixtureState{
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
				Steps: []common.Step{
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: common.Color{R: 255, G: 255, B: 255}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: common.Color{R: 255, G: 255, B: 255}},
							2: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: common.Color{R: 255, G: 255, B: 255}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: common.Color{R: 255, G: 255, B: 255}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		pattern := common.Pattern{}
		pattern.Steps = tt.args.steps
		t.Run(tt.name, func(t *testing.T) {
			if got := ApplyFixtureState(pattern, tt.args.scannerState); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateStandardChasePatterm() got = %+v", got)
				t.Errorf("GenerateStandardChasePatterm() want %+v", tt.want)
			}
		})
	}
}

func TestApplyScannerStateAllEnabled(t *testing.T) {
	type args struct {
		steps        []common.Step
		scannerState map[int]common.FixtureState
	}
	tests := []struct {
		name string
		args args
		want common.Pattern
	}{
		{
			name: "standard 2 steps All enabled.",
			args: args{
				steps: []common.Step{
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: common.Color{R: 255, G: 255, B: 255}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: common.Color{R: 255, G: 255, B: 255}},
						},
					},
				},
				scannerState: map[int]common.FixtureState{
					0: {
						Enabled: true,
					},
					1: {
						Enabled: true,
					},
				},
			},
			want: common.Pattern{
				Steps: []common.Step{
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: common.Color{R: 255, G: 255, B: 255}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: common.Color{R: 255, G: 255, B: 255}},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		pattern := common.Pattern{}
		pattern.Steps = tt.args.steps
		t.Run(tt.name, func(t *testing.T) {
			if got := ApplyFixtureState(pattern, tt.args.scannerState); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateStandardChasePatterm() got = %+v", got)
				t.Errorf("GenerateStandardChasePatterm() want %+v", tt.want)
			}
		})
	}
}

func TestApplyScannerStateFirst4Disabled(t *testing.T) {
	type args struct {
		steps        []common.Step
		scannerState map[int]common.FixtureState
	}
	tests := []struct {
		name string
		args args
		want common.Pattern
	}{

		{
			name: "standard 8 steps  first 4 disabled",
			args: args{
				steps: []common.Step{
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: common.Color{R: 255, G: 255, B: 255}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: common.Color{R: 255, G: 255, B: 255}},
							2: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: common.Color{R: 255, G: 255, B: 255}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: common.Color{R: 255, G: 255, B: 255}},
							4: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: common.Color{R: 255, G: 255, B: 255}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: common.Color{R: 255, G: 255, B: 255}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: common.Color{R: 255, G: 255, B: 255}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: common.Color{R: 255, G: 255, B: 255}},
						},
					},
				},
				scannerState: map[int]common.FixtureState{
					0: {
						Enabled: false, // Disabled.
					},
					1: {
						Enabled: false, // Disabled.
					},
					2: {
						Enabled: false, // Disabled.
					},
					3: {
						Enabled: false, // Disabled.
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
				},
			},
			want: common.Pattern{

				Steps: []common.Step{
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: common.Color{R: 255, G: 255, B: 255}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: common.Color{R: 255, G: 255, B: 255}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: common.Color{R: 255, G: 255, B: 255}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: common.Color{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: common.Color{R: 255, G: 255, B: 255}},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		pattern := common.Pattern{}
		pattern.Steps = tt.args.steps
		t.Run(tt.name, func(t *testing.T) {
			if got := ApplyFixtureState(pattern, tt.args.scannerState); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateStandardChasePatterm() got = %+v", got)
				t.Errorf("GenerateStandardChasePatterm() want %+v", tt.want)

				for stepNumber, step := range got.Steps {
					fmt.Printf("Step %d\n", stepNumber)
					for fixtureNumber := 0; fixtureNumber < len(step.Fixtures); fixtureNumber++ {
						fixture := step.Fixtures[fixtureNumber]
						fmt.Printf("\t\tFixture %d Enabled %t Master %d Brightness %d Shutter %d Color %+v\n", fixtureNumber, fixture.Enabled, fixture.MasterDimmer, fixture.Brightness, fixture.Shutter, fixture.Color)
					}
				}
			}
		})
	}
}
