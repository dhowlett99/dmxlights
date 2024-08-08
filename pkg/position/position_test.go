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
	"image/color"
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

var threeFixturesRGBInverted = map[int]common.FixtureState{
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
				NumberFixtures:        3,
				ScannerChaser:         false,
				RGBShift:              0,
			},
			steps: []common.Step{
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
					},
				},
			},
			want: map[int][]common.FixtureBuffer{
				0: { // Fixture 1 Fade up and down.
					{BaseColor: common.Red, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 50, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 255, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 255, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 50, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 1 OFF
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 1 OFF
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
				},

				1: {
					// Fixture 2 OFF
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 2 Fade up and down.
					{BaseColor: common.Red, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 50, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 255, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 255, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 50, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 2 OFF
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
				},

				2: {
					// Fixture 3 OFF
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 3 OFF
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 2 Fade up and down.
					{BaseColor: common.Red, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 50, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 255, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 255, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 50, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
				},
			},
			want1: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fadeColors, _ := CalculatePositions(tt.steps, tt.sequence, tt.scanner)

			if !reflect.DeepEqual(fadeColors, tt.want) {

				t.Errorf("CalculatePositions() ")

				fmt.Printf("++++++++++++++ WANT ++++++++++++++++++++\n")
				for fixtureNumber := 0; fixtureNumber < len(tt.want); fixtureNumber++ {

					fade := tt.want[fixtureNumber]

					fmt.Printf("fixtureNumber:%d ============================\n", fixtureNumber)

					for number, fixtureBuffer := range fade {
						fmt.Printf("Buffer:%d fixtureBuffer:%+v\n", number, fixtureBuffer.Color)
					}

				}

				fmt.Printf("++++++++++++++ GOT ++++++++++++++++++++\n")
				for fixtureNumber := 0; fixtureNumber < len(fadeColors); fixtureNumber++ {

					fade := fadeColors[fixtureNumber]

					fmt.Printf("fixtureNumber:%d ============================\n", fixtureNumber)

					step := 0
					for _, fixtureBuffer := range fade {
						fmt.Printf("\tstep %d rule %d %s: fixtureBuffer:%+v\n", step, fixtureBuffer.Rule, fixtureBuffer.DebugMsg, fixtureBuffer.Color)
						step++
					}

				}
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
				NumberFixtures:        3,
				ScannerChaser:         false,
				RGBShift:              9,
			},
			steps: []common.Step{
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
					},
				},
			},
			want: map[int][]common.FixtureBuffer{
				0: { // Fixture 1 Fade up and down.
					{BaseColor: common.Red, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 1 OFF
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 1 OFF
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
				},

				1: {
					// Fixture 2 OFF
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 2 Fade up and down.
					{BaseColor: common.Red, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 2 OFF
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
				},

				2: {
					// Fixture 3 OFF
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 3 OFF
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 2 Fade up and down.
					{BaseColor: common.Red, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
				},
			},
			want1: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fadeColors, _ := CalculatePositions(tt.steps, tt.sequence, tt.scanner)

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
				NumberFixtures:        3,
				ScannerChaser:         false,
				RGBShift:              0,
			},
			steps: []common.Step{
				{
					StepNumber: 0,
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				{
					StepNumber: 1,
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				{
					StepNumber: 2,
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
					},
				},
			},
			want: map[int][]common.FixtureBuffer{
				0: { // Fixture 1 Fade up and down.
					{BaseColor: common.Red, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 1 OFF
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 1 OFF
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 1 OFF
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 1 OFF
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 1 Fade up and down.
					{BaseColor: common.Red, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
				},

				1: {
					// Fixture 2 OFF
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 2 Fade up and down.
					{BaseColor: common.Red, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 2 OFF
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 2 OFF
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 2 Fade up and down.
					{BaseColor: common.Red, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 2 OFF
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
				},

				2: {
					// Fixture 3 OFF
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 3 OFF
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 2 Fade up and down.
					{BaseColor: common.Red, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 2 Fade up and down.
					{BaseColor: common.Red, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 3 OFF
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 3 OFF
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
				},
			},
			want1: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fadeColors, _ := CalculatePositions(tt.steps, tt.sequence, tt.scanner)

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
				NumberFixtures:        8,
				ScannerChaser:         false,
				RGBShift:              0,
			},
			steps: []common.Step{
				{
					StepNumber: 0,
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				{
					StepNumber: 1,
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				{
					StepNumber: 2,
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				{
					StepNumber: 3,
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				{
					StepNumber: 4,
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				{
					StepNumber: 5,
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				{
					StepNumber: 6,
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				{
					StepNumber: 7,
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
					},
				},
			},
			want: map[int][]common.FixtureBuffer{
				0: { // Fixture 1 Fade up and down.
					{BaseColor: common.Red, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 1 OFF
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 1 OFF
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
				},

				1: {
					// Fixture 2 OFF
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 2 Fade up and down.
					{BaseColor: common.Red, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 2 OFF
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
				},

				2: {
					// Fixture 3 OFF
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 3 OFF
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 2 Fade up and down.
					{BaseColor: common.Red, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
				},
				3: {
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},

					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},

					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
				},
				4: {
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},

					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},

					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
				},
				5: {
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},

					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},

					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
				},
				6: {
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},

					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},

					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
				},
				7: {
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},

					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},

					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
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
			fadeColors, _ := CalculatePositions(RGBPattern.Steps, tt.sequence, tt.scanner)

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
		})
	}
}

func TestApplyRGBChaseWithOnlyFourEnabledBounce(t *testing.T) {

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
			name: "Standard 8 fixtures but only four enabled, bounce.",

			scanner: false,
			sequence: common.Sequence{
				Type:           "rgb",
				Bounce:         true,
				ScannerReverse: false,
				FadeUp:         []int{0, 5, 25, 50, 75, 100, 125, 150, 175, 255},
				FadeDown:       []int{255, 175, 125, 100, 75, 50, 25, 50, 25, 0},
				RGBShift:       5,
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
				EnabledNumberFixtures: 3,
				NumberFixtures:        8,
				ScannerChaser:         false,
			},
			steps: []common.Step{
				{
					StepNumber: 0,
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				{
					StepNumber: 1,
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				{
					StepNumber: 2,
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				{
					StepNumber: 3,
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				{
					StepNumber: 4,
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				{
					StepNumber: 5,
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				{
					StepNumber: 6,
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				{
					StepNumber: 7,
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
					},
				},
			},
			want: map[int][]common.FixtureBuffer{
				0: {
					{BaseColor: common.Red, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 5, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 25, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 50, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 75, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 100, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 125, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 150, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 175, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 255, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 255, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 175, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 125, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 100, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 75, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 50, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 25, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 50, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 25, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},

					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},

					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},

					{BaseColor: common.Red, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 5, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 25, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 50, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 75, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 100, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 125, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 150, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 175, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 255, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},

					{BaseColor: common.Red, Color: color.NRGBA{R: 255, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 175, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 125, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 100, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 75, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 50, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 25, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 50, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 25, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
				},
				1: {
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},

					{BaseColor: common.Red, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 5, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 25, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 50, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 75, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},

					{BaseColor: common.Red, Color: color.NRGBA{R: 100, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 125, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 150, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 175, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 255, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 255, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 175, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 125, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 100, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 75, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},

					{BaseColor: common.Red, Color: color.NRGBA{R: 50, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 25, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 50, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 25, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},

					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},

					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},

					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},

					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},

					{BaseColor: common.Red, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 5, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 25, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 50, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 75, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},

					{BaseColor: common.Red, Color: color.NRGBA{R: 100, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 125, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 150, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 175, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 255, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},

					{BaseColor: common.Red, Color: color.NRGBA{R: 255, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 175, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 125, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 100, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 75, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},

					{BaseColor: common.Red, Color: color.NRGBA{R: 50, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 25, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 50, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 25, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},

					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
				},
				2: {
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},

					{BaseColor: common.Red, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 5, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 25, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 50, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 75, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 100, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 125, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 150, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 175, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 255, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},

					{BaseColor: common.Red, Color: color.NRGBA{R: 255, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 175, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 125, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 100, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 75, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 50, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 25, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 50, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 25, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},

					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},

					{BaseColor: common.Red, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 5, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 25, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 50, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 75, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 100, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 125, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 150, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 175, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 255, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},

					{BaseColor: common.Red, Color: color.NRGBA{R: 255, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 175, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 125, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 100, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 75, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 50, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 25, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 50, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 25, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},

					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},

					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
				},
				3: {
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},

					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},

					{BaseColor: common.Red, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 5, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 25, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 50, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 75, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 100, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 125, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 150, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 175, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 255, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 255, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 175, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 125, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 100, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 75, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},

					{BaseColor: common.Red, Color: color.NRGBA{R: 50, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 25, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 50, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 25, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 5, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 25, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 50, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 75, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 100, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 125, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 150, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 175, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 255, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 255, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 175, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 125, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 100, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 75, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},

					{BaseColor: common.Red, Color: color.NRGBA{R: 50, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 25, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 50, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 25, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Red, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},

					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
				},
				4: {
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},

					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
				},
				5: {
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},

					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
				},
				6: {
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},

					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
				},
				7: {
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},

					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: color.NRGBA{R: 0, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: false},
				},
			},
			want1: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			pattern := common.Pattern{}
			pattern.Steps = tt.steps

			// for _, step := range pattern.Steps {
			// 	fmt.Printf("Step %d\n", step.StepNumber)
			// 	for fixtureNumber := 0; fixtureNumber < len(step.Fixtures); fixtureNumber++ {
			// 		fixture := step.Fixtures[fixtureNumber]
			// 		fmt.Printf("Fixture %d Color %+v\n", fixture.Number, fixture.Color)
			// 	}
			// }

			RGBPattern := ApplyFixtureState(pattern, tt.sequence.FixtureState)
			got, _ := CalculatePositions(RGBPattern.Steps, tt.sequence, tt.scanner)

			if !reflect.DeepEqual(got, tt.want) {

				t.Errorf("CalculatePositions() ")

				fmt.Printf("++++++++++++++ WANT ++++++++++++++++++++\n")
				for fixtureNumber := 0; fixtureNumber < len(tt.want); fixtureNumber++ {

					fade := tt.want[fixtureNumber]

					fmt.Printf("fixtureNumber:%d ============================\n", fixtureNumber)

					for step, fixtureBuffer := range fade {
						//fmt.Printf("\tstep %d rule %d %s: Color:%+v BaseColor:%+v\n", step, fixtureBuffer.Rule, fixtureBuffer.DebugMsg, common.GetColorNameByRGB(fixtureBuffer.Color), common.GetColorNameByRGB(fixtureBuffer.BaseColor))
						fmt.Printf("\tstep %d FixtureBufffer:%+v\n", step, fixtureBuffer)
					}
				}

				fmt.Printf("++++++++++++++ GOT ++++++++++++++++++++\n")
				for fixtureNumber := 0; fixtureNumber < len(got); fixtureNumber++ {

					fade := got[fixtureNumber]

					fmt.Printf("%d: {\n", fixtureNumber)

					for step, fixtureBuffer := range fade {
						//fmt.Printf("\tstep %d rule %d %s: Color:%+v BaseColor:%+v\n", step, fixtureBuffer.Rule, fixtureBuffer.DebugMsg, common.GetColorNameByRGB(fixtureBuffer.Color), common.GetColorNameByRGB(fixtureBuffer.BaseColor))
						fmt.Printf("\tstep %d FixtureBufffer:%+v\n", step, fixtureBuffer)
					}
				}
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
				NumberFixtures:        3,
				ScannerChaser:         false,
				RGBShift:              0,
			},
			steps: []common.Step{
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
					},
				},
			},
			want: map[int][]common.FixtureBuffer{
				0: { // Fixture 0 Fade down.
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 0 Fade Up.
					{BaseColor: common.Red, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 0 On
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 0 On
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
				},

				1: {
					// Fixture 1 On
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 1 Fade down.
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 1 Fade Up.
					{BaseColor: common.Red, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 1 On
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
				},

				2: {
					// Fixture 2 On
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 2 On
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},

					// Fixture 2 Fade down.
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 2 Fade Up.
					{BaseColor: common.Red, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Red, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
				},
			},
			want1: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := CalculatePositions(tt.steps, tt.sequence, tt.scanner)

			if !reflect.DeepEqual(got, tt.want) {

				t.Errorf("CalculatePositions() ")

				fmt.Printf("++++++++++++++ WANT ++++++++++++++++++++\n")
				for fixtureNumber := 0; fixtureNumber < len(tt.want); fixtureNumber++ {

					fade := tt.want[fixtureNumber]

					fmt.Printf("fixtureNumber:%d ============================\n", fixtureNumber)

					for step, fixtureBuffer := range fade {
						fmt.Printf("step %d FixtureBuffer:%+v \n", step, fixtureBuffer)
					}

				}

				fmt.Printf("++++++++++++++ GOT ++++++++++++++++++++\n")
				for fixtureNumber := 0; fixtureNumber < len(got); fixtureNumber++ {

					fade := got[fixtureNumber]

					fmt.Printf("fixtureNumber:%d ============================\n", fixtureNumber)

					for step, fixtureBuffer := range fade {
						fmt.Printf("step %d FixtureBuffer:%+v \n", step, fixtureBuffer)
					}

				}
			}
		})
	}
}

func TestAssemblePositions(t *testing.T) {

	tests := []struct {
		name               string
		fadeColors         map[int][]common.FixtureBuffer
		fixtureState       map[int]common.FixtureState
		numberFixtures     int
		totalNumberOfSteps int
		want               map[int]common.Position
	}{
		{
			name:               "golden path assemble fade colors into positions",
			numberFixtures:     4,
			totalNumberOfSteps: 24,
			fixtureState:       allFixturesEnabled,
			fadeColors: map[int][]common.FixtureBuffer{
				0: { // Fixture 1 Fade up and down.
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 1 OFF
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 1 OFF
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 1 OFF
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
				},

				1: {
					// Fixture 2 OFF
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 2 Fade up and down.
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 2 OFF
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 2 OFF
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
				},

				2: {
					// Fixture 3 OFF
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 3 OFF
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 3 Fade up and down.
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 3 OFF
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
				},

				3: {
					// Fixture 4 OFF
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 4 OFF
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 4 OFF
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 4 Fade up and down.
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
				},
			},

			want: map[int]common.Position{
				0: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},

				1: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				2: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				3: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				4: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				5: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				6: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				7: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				8: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				9: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				10: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				11: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				12: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				13: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				14: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				15: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				16: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				17: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},

				18: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				19: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				20: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				21: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				22: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				23: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if positions, _ := AssemblePositions(tt.fadeColors, tt.numberFixtures, tt.totalNumberOfSteps, tt.fixtureState, false); !reflect.DeepEqual(positions, tt.want) {
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
		fixtureState       map[int]common.FixtureState
		numberFixtures     int
		totalNumberOfSteps int
		optimisation       bool
		want               map[int]common.Position
	}{
		{
			name:               "golden path assemble fade colors into positions",
			optimisation:       false,
			numberFixtures:     3,
			totalNumberOfSteps: 18,
			fixtureState: map[int]common.FixtureState{
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
			fadeColors: map[int][]common.FixtureBuffer{
				0: { // Fixture 1 Fade up and down.
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 1 OFF
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 1 OFF
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
				},

				1: {
					// Fixture 2 OFF
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 2 Fade up and down.
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 2 OFF
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
				},

				2: {
					// Fixture 3 OFF
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 3 OFF
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 2 Fade up and down.
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
				},
				3: {
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},

					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},

					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
				},
				4: {
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},

					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},

					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
				},
				5: {
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},

					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},

					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
				},
				6: {
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},

					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},

					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
				},
				7: {
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},

					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},

					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
				},
			},
			want: map[int]common.Position{
				0: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},

				1: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				2: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				3: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				4: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				5: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				6: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				7: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				8: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				9: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				10: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				11: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				12: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				13: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				14: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				15: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				16: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				17: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if positions, _ := AssemblePositions(tt.fadeColors, tt.numberFixtures, tt.totalNumberOfSteps, tt.fixtureState, tt.optimisation); !reflect.DeepEqual(positions, tt.want) {
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

func TestAssemblePositionsOnlyThreeEnabledWithOptimisation(t *testing.T) {

	tests := []struct {
		name               string
		fadeColors         map[int][]common.FixtureBuffer
		numberFixtures     int
		totalNumberOfSteps int
		optimisation       bool
		want               map[int]common.Position
		want1              int
	}{
		{
			name:               "golden path assemble fade colors into positions",
			optimisation:       true,
			numberFixtures:     3,
			totalNumberOfSteps: 18,
			want1:              15,
			fadeColors: map[int][]common.FixtureBuffer{
				0: { // Fixture 1 Fade up and down.
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 1 OFF
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 1 OFF
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
				},

				1: {
					// Fixture 2 OFF
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 2 Fade up and down.
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 2 OFF
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
				},

				2: {
					// Fixture 3 OFF
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 3 OFF
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					// Fixture 2 Fade up and down.
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Red, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: true},
				},
				3: {
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},

					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},

					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
				},
				4: {
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},

					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},

					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
				},
				5: {
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},

					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},

					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
				},
				6: {
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},

					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},

					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
				},
				7: {
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},

					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},

					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
				},
			},
			want: map[int]common.Position{
				0: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},

				1: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						// 	1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						// 	2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				2: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						// 1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						// 2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},

				3: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						// 1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						// 2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},

				4: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						//1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						//2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				5: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						//1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						//2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},

				6: {
					Fixtures: map[int]common.Fixture{
						//0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						//2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				7: {
					Fixtures: map[int]common.Fixture{
						//0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						//2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				8: {
					Fixtures: map[int]common.Fixture{
						//0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						//2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				9: {
					Fixtures: map[int]common.Fixture{
						//0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						//2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				10: {
					Fixtures: map[int]common.Fixture{
						//0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						//2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},

				11: {
					Fixtures: map[int]common.Fixture{
						//0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						//1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				12: {
					Fixtures: map[int]common.Fixture{
						//0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						//1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				13: {
					Fixtures: map[int]common.Fixture{
						//0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						//1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				14: {
					Fixtures: map[int]common.Fixture{
						//0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						//1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				15: {
					Fixtures: map[int]common.Fixture{
						//0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						//1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if positions, numberPositions := AssemblePositions(tt.fadeColors, tt.numberFixtures, tt.totalNumberOfSteps, allFixturesEnabled, tt.optimisation); !reflect.DeepEqual(positions, tt.want) {
				t.Errorf("AssemblePositions() = %+v, want %+v", positions, tt.want)

				for positionNumber := 0; positionNumber < len(positions); positionNumber++ {

					position := positions[positionNumber]

					fmt.Printf("Position %d\n", positionNumber)

					for fixtureNumber := 0; fixtureNumber < 3; fixtureNumber++ {

						fixture := position.Fixtures[fixtureNumber]

						fmt.Printf("\t fixture %d %+v\n", fixtureNumber, fixture.Color)

					}

				}

				if numberPositions != tt.want1 {
					t.Errorf("CalculatePositions() got1 = %v, want %v", numberPositions, tt.want1)
				}
			}

		})
	}
}

func TestCalculatePositionsOnlyFourEnabledBounceAndShiftOfFive(t *testing.T) {

	tests := []struct {
		name               string
		fadeColors         map[int][]common.FixtureBuffer
		numberFixtures     int
		totalNumberOfSteps int
		optimisation       bool
		want               map[int]common.Position
		want1              int
		fixtureState       map[int]common.FixtureState
	}{
		{
			name:               "Eight Fitures but only four enabled, bounce and shift on",
			optimisation:       false,
			numberFixtures:     8,
			totalNumberOfSteps: 70,
			want1:              70,
			fixtureState: map[int]common.FixtureState{
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
			fadeColors: map[int][]common.FixtureBuffer{
				0: {
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 5, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 25, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 75, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 100, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 125, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 150, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 175, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Red, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Red, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 175, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 125, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 100, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 75, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 25, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 25, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},

					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},

					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 5, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 25, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 75, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 100, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 125, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 150, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 175, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Red, MasterDimmer: 255, Brightness: 255, Enabled: true},

					{BaseColor: common.Black, Color: common.Red, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 175, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 125, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 100, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 75, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 25, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 25, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
				},
				1: {
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 5, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 25, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 75, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},

					{BaseColor: common.Black, Color: color.NRGBA{R: 100, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 125, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 150, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 175, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Red, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Red, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 175, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 125, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 100, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 75, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 25, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 25, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},

					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 5, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 25, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 75, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},

					{BaseColor: common.Black, Color: color.NRGBA{R: 100, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 125, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 150, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 175, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Red, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Red, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 175, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 125, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 100, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 75, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 25, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 25, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
				},
				2: {
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 5, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 25, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 75, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 100, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 125, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 150, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 175, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Red, MasterDimmer: 255, Brightness: 255, Enabled: true},

					{BaseColor: common.Black, Color: common.Red, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 175, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 125, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 100, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 75, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 25, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 25, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},

					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 5, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 25, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 75, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 100, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 125, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 150, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 175, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Red, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Red, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 175, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 125, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 100, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 75, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 25, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 25, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},

					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
				},
				3: {
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},

					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 5, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 25, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 75, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 100, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 125, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 150, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 175, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Red, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Red, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 175, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 125, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 100, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 75, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},

					{BaseColor: common.Black, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 25, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 25, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 5, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 25, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 75, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 100, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 125, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 150, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 175, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Red, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Red, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 175, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 125, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 100, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 75, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},

					{BaseColor: common.Black, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 25, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.QuarterRed, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: color.NRGBA{R: 25, G: 0, B: 0}, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: true},
				},
				4: {
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},

					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
				},
				5: {
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},

					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
				},
				6: {
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},

					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
				},
				7: {
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},

					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
					{BaseColor: common.Black, Color: common.Black, MasterDimmer: 255, Brightness: 255, Enabled: false},
				},
			},
			want: map[int]common.Position{
				0: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				1: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 5, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				2: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 25, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				3: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				4: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 75, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				5: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 100, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				6: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 125, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 5, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				7: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 150, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 25, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				8: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 175, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				9: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 75, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				10: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 100, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				11: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 175, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 125, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 5, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				12: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 125, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 150, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 25, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				13: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 100, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 175, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				14: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 75, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 75, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				15: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 100, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				16: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 25, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 175, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 125, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 5, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				17: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 125, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 150, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 25, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				18: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 25, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 100, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 175, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				19: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 75, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 75, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				20: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 100, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				21: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 25, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 175, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 125, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				22: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 125, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 150, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				23: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 25, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 100, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 175, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				24: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 75, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				25: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				26: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 25, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 175, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				27: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 125, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				28: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 25, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 100, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				29: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 75, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				30: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				31: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 25, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				32: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				33: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 25, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				34: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				35: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				36: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 5, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				37: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 25, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				38: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				39: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 75, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				40: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 100, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				41: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 5, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 125, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				42: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 25, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 150, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				43: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 175, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				44: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 75, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				45: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 100, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				46: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 5, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 125, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 175, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				47: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 25, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 150, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 125, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				48: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 175, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 100, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				49: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 75, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 75, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				50: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 100, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				51: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 5, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 125, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 175, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 25, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				52: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 25, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 150, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 125, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				53: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 175, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 100, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 25, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				54: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 75, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 75, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				55: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 100, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				56: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 125, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 175, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 25, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				57: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 150, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 125, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				58: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 175, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 100, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 25, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				59: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 75, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				60: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				61: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 175, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 25, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				62: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 125, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				63: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 100, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 25, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				64: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 75, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				65: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				66: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 25, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				67: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				68: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: color.NRGBA{R: 25, G: 0, B: 0}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
				69: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						1: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						2: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
						3: {MasterDimmer: 255, Brightness: 255, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0, Enabled: true},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if positions, numberPositions := AssemblePositions(tt.fadeColors, tt.numberFixtures, tt.totalNumberOfSteps, tt.fixtureState, tt.optimisation); !reflect.DeepEqual(positions, tt.want) {
				t.Errorf("AssemblePositions() = %+v, want %+v", positions, tt.want)

				for positionNumber := 0; positionNumber < len(positions); positionNumber++ {

					position := positions[positionNumber]

					fmt.Printf("%d: {\n", positionNumber)

					fmt.Printf("\tFixtures: map[int]common.Fixture{\n")
					for fixtureNumber := 0; fixtureNumber < 4; fixtureNumber++ {

						fixture := position.Fixtures[fixtureNumber]

						fmt.Printf("\t\t%d: {MasterDimmer: %d, Brightness: %d, Color: color.NRGBA{R: %d, G: %d, B: %d},Pan: %d, Tilt: %d, Shutter: %d, Rotate: %d, Music: %d, Gobo: %d, Program: %d ,Enabled: %t},\n",
							fixtureNumber,
							fixture.MasterDimmer,
							fixture.Brightness,
							fixture.Color.R,
							fixture.Color.G,
							fixture.Color.B,
							fixture.Pan,
							fixture.Tilt,
							fixture.Shutter,
							fixture.Rotate,
							fixture.Music,
							fixture.Gobo,
							fixture.Program,
							fixture.Enabled)

					}
					fmt.Printf("\t\t},\n")
					fmt.Printf("\t},\n")
				}

				if numberPositions != tt.want1 {
					t.Errorf("CalculatePositions() got1 = %v, want %v", numberPositions, tt.want1)
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
					NumberFixtures:        8,
					ScannerChaser:         false,
				},
			},
			steps: []common.Step{
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 255, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 255, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 255, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 255, B: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 255, B: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 255, B: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 255, B: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Color: color.NRGBA{R: 0, G: 255, B: 0}},
					},
				},
			},

			want: map[int]common.Position{
				0: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 1, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				1: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 50, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				2: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 255, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				3: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 255, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				4: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 50, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				5: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 1, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				6: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 1, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				7: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 50, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				8: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 255, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				9: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 255, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				10: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 50, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				11: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 1, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				12: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 1, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				13: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 50, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				14: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 255, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				15: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 255, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				16: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 50, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				17: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 1, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				18: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 1, B: 0}},
						4: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				19: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 50, B: 0}},
						4: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				20: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 255, B: 0}},
						4: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				21: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 255, B: 0}},
						4: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				22: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 50, B: 0}},
						4: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				23: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 1, B: 0}},
						4: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				24: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 1, B: 0}},
						5: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				25: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 50, B: 0}},
						5: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				26: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 255, B: 0}},
						5: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				27: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 255, B: 0}},
						5: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				28: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 50, B: 0}},
						5: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				29: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 1, B: 0}},
						5: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				30: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 1, B: 0}},
						6: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				31: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 50, B: 0}},
						6: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				32: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 255, B: 0}},
						6: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				33: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 255, B: 0}},
						6: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				34: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 50, B: 0}},
						6: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				35: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 1, B: 0}},
						6: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				36: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 1, B: 0}},
						7: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				37: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 50, B: 0}},
						7: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				38: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 255, B: 0}},
						7: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				39: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 255, B: 0}},
						7: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				40: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 50, B: 0}},
						7: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				41: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 1, B: 0}},
						7: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				42: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 1, B: 0}},
					},
				},
				43: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 50, B: 0}},
					},
				},
				44: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 255, B: 0}},
					},
				},
				45: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 255, B: 0}},
					},
				},
				46: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 50, B: 0}},
					},
				},
				47: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Color: color.NRGBA{R: 0, G: 1, B: 0}},
					},
				},
			},
			want1: 48,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			fadeColors, totalNumberOfSteps := CalculatePositions(tt.steps, tt.args.sequence, tt.args.scanner)
			positions, numberPositions := AssemblePositions(fadeColors, tt.args.sequence.NumberFixtures, totalNumberOfSteps, tt.args.sequence.FixtureState, false)
			if !reflect.DeepEqual(positions, tt.want) {
				t.Errorf("CalculatePositions() got = %v, want %v", positions, tt.want)

				fmt.Printf("++++++++++++++ WANT ++++++++++++++++++++\n")
				for positionNumber := 0; positionNumber < len(positions); positionNumber++ {

					fmt.Printf("Position:%d ============================\n", positionNumber)
					position := positions[positionNumber]
					for fixtureNumber := 0; fixtureNumber < len(position.Fixtures); fixtureNumber++ {

						fixture := position.Fixtures[fixtureNumber]
						fmt.Printf("Fixture:%d color:%+v BaseColor %+v\n", fixtureNumber, fixture.Color, fixture.BaseColor)
					}

				}

				fmt.Printf("++++++++++++++ GOT ++++++++++++++++++++\n")
				for fixtureNumber := 0; fixtureNumber < len(fadeColors); fixtureNumber++ {
					positionNumber := 0
					fade := fadeColors[fixtureNumber]

					fmt.Printf("fixtureNumber:%d ============================\n", fixtureNumber)

					step := 0
					for _, fixtureBuffer := range fade {
						fmt.Printf("\tpositionNumber %d step %d rule %d %s: fixtureBuffer:%+v BaseColor %+v\n", positionNumber, fixtureBuffer.Step, fixtureBuffer.Rule, fixtureBuffer.DebugMsg, fixtureBuffer.Color, fixtureBuffer.BaseColor)
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
					Bounce:                false,
					ScannerReverse:        false,
					FadeUp:                []int{255},
					Optimisation:          false,
					FixtureState:          allFixturesEnabled,
					EnabledNumberFixtures: 8,
					NumberFixtures:        8,
				},
			},
			steps: []common.Step{
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Orange},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Yellow},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Green},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Cyan},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Blue},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Purple},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Magenta},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Orange},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Yellow},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Green},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Cyan},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Blue},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Purple},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Magenta},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Yellow},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Green},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Cyan},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Blue},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Purple},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Magenta},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Orange},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Green},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Cyan},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Blue},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Purple},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Magenta},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Orange},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Yellow},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Cyan},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Blue},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Purple},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Magenta},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Orange},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Yellow},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Green},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Blue},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Purple},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Magenta},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Orange},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Yellow},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Green},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Cyan},
					},
				},
				{
					Fixtures: map[int]common.Fixture{

						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Purple},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Magenta},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Orange},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Yellow},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Green},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Cyan},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Blue},
					},
				},
				{
					Fixtures: map[int]common.Fixture{

						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Magenta},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Orange},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Yellow},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Green},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Cyan},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Blue},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Purple},
					},
				},
			},

			want: map[int]common.Position{
				// Start of want.
				0: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {BaseColor: common.Orange, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Orange, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Yellow, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Green, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {BaseColor: common.Cyan, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Cyan, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Blue, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {BaseColor: common.Purple, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Purple, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {BaseColor: common.Magenta, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Magenta, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},

				1: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Orange, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Orange, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Yellow, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Green, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {BaseColor: common.Cyan, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Cyan, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Blue, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {BaseColor: common.Purple, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Purple, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {BaseColor: common.Magenta, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Magenta, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},

				2: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Yellow, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Green, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {BaseColor: common.Cyan, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Cyan, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Blue, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {BaseColor: common.Purple, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Purple, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {BaseColor: common.Magenta, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Magenta, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {BaseColor: common.Orange, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Orange, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},

				3: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Green, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {BaseColor: common.Cyan, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Cyan, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Blue, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {BaseColor: common.Purple, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Purple, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {BaseColor: common.Magenta, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Magenta, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {BaseColor: common.Orange, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Orange, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Yellow, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				4: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Cyan, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Cyan, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Blue, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {BaseColor: common.Purple, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Purple, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {BaseColor: common.Magenta, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Magenta, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {BaseColor: common.Orange, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Orange, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Yellow, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Green, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				5: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Blue, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {BaseColor: common.Purple, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Purple, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {BaseColor: common.Magenta, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Magenta, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {BaseColor: common.Orange, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Orange, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Yellow, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Green, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {BaseColor: common.Cyan, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Cyan, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				6: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Purple, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Purple, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {BaseColor: common.Magenta, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Magenta, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {BaseColor: common.Orange, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Orange, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Yellow, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Green, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {BaseColor: common.Cyan, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Cyan, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Blue, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				7: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Magenta, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Magenta, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {BaseColor: common.Orange, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Orange, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Yellow, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Green, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {BaseColor: common.Cyan, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Cyan, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Blue, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {BaseColor: common.Purple, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Purple, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				// End of Want
			},
			want1: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			fadeColors, totalNumberOfSteps := CalculatePositions(tt.steps, tt.args.sequence, true)
			positions, numberPositions := AssemblePositions(fadeColors, tt.args.sequence.NumberFixtures, totalNumberOfSteps, tt.args.sequence.FixtureState, false)
			if !reflect.DeepEqual(positions, tt.want) {
				t.Errorf("CalculatePositions() got = %v, want %v", positions, tt.want)

				fmt.Printf("==========WANT===============\n")
				for positionNumber := 0; positionNumber < len(tt.want); positionNumber++ {
					position := tt.want[positionNumber]
					fmt.Printf("Positon %d\n", positionNumber)
					for fixtureNumber := 0; fixtureNumber < len(position.Fixtures); fixtureNumber++ {
						fixture := position.Fixtures[fixtureNumber]
						fmt.Printf("step:%d FixtureBuffer%+v\n", fixtureNumber, fixture)
					}
				}

				fmt.Printf("==========GOT===============\n")
				for positionNumber := 0; positionNumber < len(positions); positionNumber++ {
					position := positions[positionNumber]
					fmt.Printf("Positon %d\n", positionNumber)
					for fixtureNumber := 0; fixtureNumber < len(position.Fixtures); fixtureNumber++ {
						fixture := position.Fixtures[fixtureNumber]
						fmt.Printf("step:%d FixtureBuffer%+v\n", fixtureNumber, fixture)
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
					NumberFixtures:        4,
				},
			},
			steps: []common.Step{
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 255, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 255, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 255, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 255, B: 0}},
					},
				},
			},

			want: map[int]common.Position{
				0: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 1, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				1: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 50, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				2: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 255, B: 0}},
						1: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 1, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				3: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 50, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				4: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 255, B: 0}},
						2: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 1, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				5: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 50, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				6: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 255, B: 0}},
						3: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 1, B: 0}},
					},
				},
				7: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 50, B: 0}},
					},
				},
				8: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 255, B: 0}},
					},
				},
			},

			want1: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fadeColors, totalNumberOfSteps := CalculatePositions(tt.steps, tt.args.sequence, true)
			positions, numberPositions := AssemblePositions(fadeColors, tt.args.sequence.NumberFixtures, totalNumberOfSteps, tt.args.sequence.FixtureState, false)
			if !reflect.DeepEqual(positions, tt.want) {
				t.Errorf("CalculatePositions() got = %v, want %v", positions, tt.want)

				fmt.Printf(" ================== Want =====================\n")
				for positionNumber := 0; positionNumber < len(tt.want); positionNumber++ {

					fmt.Printf("Position:%d ============================\n", positionNumber)
					position := tt.want[positionNumber]
					for fixtureNumber := 0; fixtureNumber < len(position.Fixtures); fixtureNumber++ {
						fixture := position.Fixtures[fixtureNumber]
						fmt.Printf("step:%d fixture %+v\n", fixtureNumber, fixture)
					}

				}

				fmt.Printf(" ================== Got =====================\n")
				for positionNumber := 0; positionNumber < len(positions); positionNumber++ {
					fmt.Printf("Position:%d ============================\n", positionNumber)
					position := positions[positionNumber]
					for fixtureNumber := 0; fixtureNumber < len(position.Fixtures); fixtureNumber++ {
						fixture := position.Fixtures[fixtureNumber]
						fmt.Printf("step:%d fixture %+v\n", fixtureNumber, fixture)
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
					NumberFixtures:        4,
				},
			},
			steps: []common.Step{
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 255, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 255, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 255, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 255, B: 0}},
					},
				},
			},

			want: map[int]common.Position{
				0: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 1, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				1: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 50, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				2: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 255, B: 0}},
						1: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 1, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				3: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 50, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				4: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 255, B: 0}},
						2: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 1, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				5: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 50, B: 0}},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				6: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 255, B: 0}},
						3: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 1, B: 0}},
					},
				},
				7: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 50, B: 0}},
					},
				},
				8: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 255, B: 0}},
					},
				},
			},
			want1: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Optimisation is turned off for testing.
			fadeColors, totalNumberOfSteps := CalculatePositions(tt.steps, tt.args.sequence, true)
			positions, numberPositions := AssemblePositions(fadeColors, tt.args.sequence.NumberFixtures, totalNumberOfSteps, tt.args.sequence.FixtureState, false)
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
					NumberFixtures:        8,
					ScannerChaser:         false,
					RGBShift:              0,
				},
			},
			steps: []common.Step{
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
					},
				},
			},

			want: map[int]common.Position{

				0: {
					Fixtures: map[int]common.Fixture{

						0: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				1: {
					Fixtures: map[int]common.Fixture{

						0: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				2: {
					Fixtures: map[int]common.Fixture{

						0: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				3: {
					Fixtures: map[int]common.Fixture{

						0: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				4: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				5: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				6: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				7: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				8: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				9: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				10: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				11: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Optimisation is turned off for testing.
			fadeColors, totalNumberOfSteps := CalculatePositions(tt.steps, tt.args.sequence, tt.args.scanner)
			got, _ := AssemblePositions(fadeColors, tt.args.sequence.NumberFixtures, totalNumberOfSteps, tt.args.sequence.FixtureState, false)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculatePositions() got = %+v, want %+v", got, tt.want)

				fmt.Printf("++++++++++++++ WANT ++++++++++++++++++++\n")
				for positionNumnber := 0; positionNumnber < len(tt.want); positionNumnber++ {
					position := tt.want[positionNumnber]
					fmt.Printf("positionNumnber:%d ============================\n", positionNumnber)
					for fixtureNumber := 0; fixtureNumber < len(position.Fixtures); fixtureNumber++ {
						fixture := position.Fixtures[fixtureNumber]
						fmt.Printf("step %d Fixture:%+v\n", fixtureNumber, fixture)
					}
				}

				fmt.Printf("++++++++++++++ GOT ++++++++++++++++++++\n")
				for positionNumnber := 0; positionNumnber < len(got); positionNumnber++ {
					position := got[positionNumnber]
					fmt.Printf("positionNumnber:%d ============================\n", positionNumnber)
					for fixtureNumber := 0; fixtureNumber < len(position.Fixtures); fixtureNumber++ {
						fixture := position.Fixtures[fixtureNumber]
						fmt.Printf("step %d Fixture:%+v\n", fixtureNumber, fixture)
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
					NumberFixtures:        8,
					ScannerChaser:         false,
					RGBShift:              0,
				},
			},
			steps: []common.Step{
				{
					Fixtures: map[int]common.Fixture{ // Red
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
					},
				},
				{
					Fixtures: map[int]common.Fixture{ // Green
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Green},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Green},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Green},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Green},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Green},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Green},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Green},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Green},
					},
				},
				{
					Fixtures: map[int]common.Fixture{ // Blue
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Blue},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Blue},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Blue},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Blue},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Blue},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Blue},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Blue},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Blue},
					},
				},
				{
					Fixtures: map[int]common.Fixture{ // Yellow
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Yellow},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Yellow},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Yellow},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Yellow},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Yellow},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Yellow},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Yellow},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Yellow},
					},
				},
			},

			want: map[int]common.Position{

				0: {
					Fixtures: map[int]common.Fixture{

						0: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				1: {
					Fixtures: map[int]common.Fixture{

						0: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				2: {
					Fixtures: map[int]common.Fixture{

						0: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				3: {
					Fixtures: map[int]common.Fixture{

						0: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				4: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterRed, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				5: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				6: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				7: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterGreen, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterGreen, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterGreen, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterGreen, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterGreen, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterGreen, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterGreen, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterGreen, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				8: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Green, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Green, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Green, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Green, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Green, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Green, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Green, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Green, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				9: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Green, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Green, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Green, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Green, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Green, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Green, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Green, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Green, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				10: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterGreen, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterGreen, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterGreen, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterGreen, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterGreen, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterGreen, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterGreen, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterGreen, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				11: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				12: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				13: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterBlue, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterBlue, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterBlue, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterBlue, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterBlue, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterBlue, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterBlue, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterBlue, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				14: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Blue, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Blue, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Blue, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Blue, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Blue, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Blue, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Blue, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Blue, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				15: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Blue, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Blue, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Blue, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Blue, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Blue, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Blue, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Blue, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Blue, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				16: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterBlue, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterBlue, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterBlue, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterBlue, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterBlue, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterBlue, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterBlue, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterBlue, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				17: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				18: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				19: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterYellow, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterYellow, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterYellow, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterYellow, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterYellow, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterYellow, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterYellow, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterYellow, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				20: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Yellow, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Yellow, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Yellow, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Yellow, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Yellow, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Yellow, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Yellow, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Yellow, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				21: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Yellow, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Yellow, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Yellow, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Yellow, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Yellow, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Yellow, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Yellow, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Yellow, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				22: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterYellow, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterYellow, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterYellow, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterYellow, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterYellow, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterYellow, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterYellow, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.QuarterYellow, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				23: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {BaseColor: common.Yellow, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Optimisation is turned off for testing.
			fadeColors, totalNumberOfSteps := CalculatePositions(tt.steps, tt.args.sequence, tt.args.scanner)
			got, _ := AssemblePositions(fadeColors, tt.args.sequence.NumberFixtures, totalNumberOfSteps, tt.args.sequence.FixtureState, false)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculatePositions() got = %v, want %v", got, tt.want)

				fmt.Printf("++++++++++++++ WANT ++++++++++++++++++++\n")
				for positionNumnber := 0; positionNumnber < len(tt.want); positionNumnber++ {
					position := tt.want[positionNumnber]
					fmt.Printf("positionNumnber:%d ============================\n", positionNumnber)
					for fixtureNumber := 0; fixtureNumber < len(position.Fixtures); fixtureNumber++ {
						fixture := position.Fixtures[fixtureNumber]
						fmt.Printf("step %d Fixture:%+v\n", fixtureNumber, fixture)
					}
				}

				fmt.Printf("++++++++++++++ GOT ++++++++++++++++++++\n")
				for positionNumnber := 0; positionNumnber < len(got); positionNumnber++ {
					position := got[positionNumnber]
					fmt.Printf("positionNumnber:%d ============================\n", positionNumnber)
					for fixtureNumber := 0; fixtureNumber < len(position.Fixtures); fixtureNumber++ {
						fixture := position.Fixtures[fixtureNumber]
						fmt.Printf("step %d Fixture:%+v\n", fixtureNumber, fixture)
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
					NumberFixtures:        2,
				},
			},
			steps: []common.Step{
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Blue, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 255},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Blue, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 255},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red, Gobo: 36, Shutter: 255, Pan: 50, Tilt: 200},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red, Gobo: 36, Shutter: 255, Pan: 50, Tilt: 200},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Magenta, Gobo: 36, Shutter: 255, Pan: 100, Tilt: 150},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Magenta, Gobo: 36, Shutter: 255, Pan: 100, Tilt: 150},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Cyan, Gobo: 36, Shutter: 255, Pan: 150, Tilt: 100},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Cyan, Gobo: 36, Shutter: 255, Pan: 150, Tilt: 100},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Purple, Gobo: 36, Shutter: 255, Pan: 200, Tilt: 50},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Purple, Gobo: 36, Shutter: 255, Pan: 200, Tilt: 50},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 190, B: 255}, Gobo: 36, Shutter: 255, Pan: 255, Tilt: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 190, B: 255}, Gobo: 36, Shutter: 255, Pan: 255, Tilt: 0},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 145, G: 0, B: 255}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 145, G: 0, B: 255}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 0},
					},
				},
			},

			want: map[int]common.Position{
				0: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Blue, Pan: 0, Tilt: 255, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
						1: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Blue, Pan: 0, Tilt: 255, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
					},
				},
				1: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 50, Tilt: 200, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
						1: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 50, Tilt: 200, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
					},
				},
				2: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Magenta, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Magenta, Pan: 100, Tilt: 150, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
						1: {BaseColor: common.Magenta, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Magenta, Pan: 100, Tilt: 150, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
					},
				},
				3: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Cyan, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Cyan, Pan: 150, Tilt: 100, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
						1: {BaseColor: common.Cyan, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Cyan, Pan: 150, Tilt: 100, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
					},
				},
				4: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Purple, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Purple, Pan: 200, Tilt: 50, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
						1: {BaseColor: common.Purple, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Purple, Pan: 200, Tilt: 50, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
					},
				},
				5: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: color.NRGBA{R: 0, G: 190, B: 255}, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: color.NRGBA{R: 0, G: 190, B: 255}, Pan: 255, Tilt: 0, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
						1: {BaseColor: color.NRGBA{R: 0, G: 190, B: 255}, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: color.NRGBA{R: 0, G: 190, B: 255}, Pan: 255, Tilt: 0, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
					},
				},
				6: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: color.NRGBA{R: 145, G: 0, B: 255}, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: color.NRGBA{R: 145, G: 0, B: 255}, Pan: 0, Tilt: 0, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
						1: {BaseColor: color.NRGBA{R: 145, G: 0, B: 255}, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: color.NRGBA{R: 145, G: 0, B: 255}, Pan: 0, Tilt: 0, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fadeColors, totalNumberOfSteps := CalculatePositions(tt.steps, tt.args.sequence, true)
			positions, _ := AssemblePositions(fadeColors, tt.args.sequence.NumberFixtures, totalNumberOfSteps, tt.args.sequence.FixtureState, false)
			if !reflect.DeepEqual(positions, tt.want) {
				t.Errorf("got = %+v", positions)
				t.Errorf("want =%+v", tt.want)

				// fmt.Printf("++++++++++++++ GOT FADES ++++++++++++++++++++\n")
				// for fixtureNumber := 0; fixtureNumber < len(fadeColors); fixtureNumber++ {
				// 	positionNumber := 0
				// 	fade := fadeColors[fixtureNumber]

				// 	fmt.Printf("fixtureNumber:%d ============================\n", fixtureNumber)

				// 	step := 0
				// 	for _, fixtureBuffer := range fade {
				// 		fmt.Printf("\tpositionNumber %d step %d rule %d %s: fixtureBuffer:%+v Pan:%d Tilt:%d\n", positionNumber, fixtureBuffer.Step, fixtureBuffer.Rule, fixtureBuffer.DebugMsg, fixtureBuffer.Color, fixtureBuffer.Pan, fixtureBuffer.Tilt)
				// 		step++
				// 		positionNumber++
				// 	}
				// }

				fmt.Printf("++++++++++++++ WANT ++++++++++++++++++++\n")
				for positionNumber := 0; positionNumber < len(tt.want); positionNumber++ {
					fmt.Printf("Position:%d ============================\n", positionNumber)
					position := tt.want[positionNumber]
					for fixtureNumber := 0; fixtureNumber < len(position.Fixtures); fixtureNumber++ {
						fixture := position.Fixtures[fixtureNumber]
						fmt.Printf("Fixture:%d color:%+v Pan %d Tilt %d BaseColor %+v\n", fixtureNumber, fixture.Color, fixture.Pan, fixture.Tilt, fixture.BaseColor)
					}
				}

				fmt.Printf("++++++++++++++ GOT ++++++++++++++++++++\n")
				for positionNumber := 0; positionNumber < len(positions); positionNumber++ {
					fmt.Printf("Position:%d ============================\n", positionNumber)
					position := positions[positionNumber]
					for fixtureNumber := 0; fixtureNumber < len(position.Fixtures); fixtureNumber++ {
						fixture := position.Fixtures[fixtureNumber]
						fmt.Printf("Fixture:%d color:%+v Pan %d Tilt %d BaseColor %+v\n", fixtureNumber, fixture.Color, fixture.Pan, fixture.Tilt, fixture.BaseColor)
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
					NumberFixtures:        2,
					ScannerChaser:         false,
				},
			},
			steps: []common.Step{
				{

					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Blue, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 255},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Blue, Gobo: 36, Shutter: 255, Pan: 128, Tilt: 255},
					},
				},
				{

					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red, Gobo: 36, Shutter: 255, Pan: 50, Tilt: 200},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red, Gobo: 36, Shutter: 255, Pan: 128, Tilt: 200},
					},
				},
			},

			want: map[int]common.Position{
				0: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Blue, Pan: 0, Tilt: 255, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
						1: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Blue, Pan: 128, Tilt: 255, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
					},
				},

				1: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 50, Tilt: 200, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
						1: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 128, Tilt: 200, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
					},
				},
			},
		},

		// End of test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fadeColors, totalNumberOfSteps := CalculatePositions(tt.steps, tt.args.sequence, true)
			positions, _ := AssemblePositions(fadeColors, tt.args.sequence.NumberFixtures, totalNumberOfSteps, tt.args.sequence.FixtureState, false)
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
					NumberFixtures:        1,
				},
			},
			steps: []common.Step{
				{

					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Blue, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 255},
					},
				},
				{

					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red, Gobo: 36, Shutter: 255, Pan: 50, Tilt: 200},
					},
				},
			},

			want: map[int]common.Position{

				0: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Blue, Pan: 0, Tilt: 255, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
					},
				},
				1: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 50, Tilt: 200, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fadeColors, totalNumberOfSteps := CalculatePositions(tt.steps, tt.args.sequence, true)
			positions, _ := AssemblePositions(fadeColors, tt.args.sequence.NumberFixtures, totalNumberOfSteps, tt.args.sequence.FixtureState, false)
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
					NumberFixtures:        8,
					ScannerChaser:         false,
				},
			},
			steps: []common.Step{
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red},
					},
				},
			},

			want: map[int]common.Position{

				0: {
					Fixtures: map[int]common.Fixture{

						0: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				1: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {BaseColor: common.Black, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Black, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fadeColors, totalNumberOfSteps := CalculatePositions(tt.steps, tt.args.sequence, true)
			positions, _ := AssemblePositions(fadeColors, tt.args.sequence.NumberFixtures, totalNumberOfSteps, tt.args.sequence.FixtureState, false)
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
					NumberFixtures:        1,
					ScannerChaser:         false,
				},
			},
			steps: []common.Step{
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red, Pan: 0, Tilt: 0, Gobo: 1, Shutter: 255},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Green, Pan: 1, Tilt: 1, Gobo: 1, Shutter: 255},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Blue, Pan: 2, Tilt: 2, Gobo: 1, Shutter: 255},
					},
				},
			},

			want: map[int]common.Position{
				0: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Blue, Pan: 2, Tilt: 2, Shutter: 255, Rotate: 0, Music: 0, Gobo: 1, Program: 0},
					},
				},
				1: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Green, Pan: 1, Tilt: 1, Shutter: 255, Rotate: 0, Music: 0, Gobo: 1, Program: 0},
					},
				},
				2: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Black, Color: common.Red, Pan: 0, Tilt: 0, Shutter: 255, Rotate: 0, Music: 0, Gobo: 1, Program: 0},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fadeColors, totalNumberOfSteps := CalculatePositions(tt.steps, tt.args.sequence, true)
			positions, _ := AssemblePositions(fadeColors, tt.args.sequence.NumberFixtures, totalNumberOfSteps, tt.args.sequence.FixtureState, false)
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
					NumberFixtures:        1,
				},
			},
			steps: []common.Step{
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red, Pan: 0, Tilt: 0, Gobo: 1, Shutter: 255},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Green, Pan: 1, Tilt: 1, Gobo: 1, Shutter: 255},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Blue, Pan: 2, Tilt: 2, Gobo: 1, Shutter: 255},
					},
				},
			},

			want: map[int]common.Position{
				0: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Blue, Gobo: 1, Shutter: 255, Pan: 2, Tilt: 2},
					},
				},

				1: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Green, Gobo: 1, Shutter: 255, Pan: 1, Tilt: 1},
					},
				},
				2: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red, Gobo: 1, Shutter: 255, Pan: 0, Tilt: 0},
					},
				},
				3: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Red, MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Red, Gobo: 1, Shutter: 255, Pan: 0, Tilt: 0},
					},
				},
				4: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Green, MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Green, Gobo: 1, Shutter: 255, Pan: 1, Tilt: 1},
					},
				},
				5: {
					Fixtures: map[int]common.Fixture{
						0: {BaseColor: common.Blue, MasterDimmer: full, Enabled: true, Brightness: full, Color: common.Blue, Gobo: 1, Shutter: 255, Pan: 2, Tilt: 2},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fadeColors, totalNumberOfSteps := CalculatePositions(tt.steps, tt.args.sequence, true)
			positions, _ := AssemblePositions(fadeColors, tt.args.sequence.NumberFixtures, totalNumberOfSteps, tt.args.sequence.FixtureState, false)
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
		steps          []common.Step
		colors         []color.NRGBA
		numberFixtures int
		fixtureState   map[int]common.FixtureState
	}
	tests := []struct {
		name string
		args args
		want []common.Step
	}{
		{
			name: "invert a single color.",
			args: args{
				fixtureState:   threeFixturesRGBInverted,
				numberFixtures: 3,
				steps: []common.Step{
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Color: common.Green},
							1: {MasterDimmer: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Color: common.Green},
							2: {MasterDimmer: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Color: common.Green},
						},
					},
				},
				colors: []color.NRGBA{
					{R: 0, G: 255, B: 0},
				},
			},
			want: []common.Step{
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Color: color.NRGBA{R: 0, G: 0, B: 0}, Inverted: true, Enabled: true},
						1: {MasterDimmer: full, Color: common.Green, Inverted: true, Enabled: true},
						2: {MasterDimmer: full, Color: common.Green, Inverted: true, Enabled: true},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Color: common.Green, Inverted: true, Enabled: true},
						1: {MasterDimmer: full, Color: color.NRGBA{R: 0, G: 0, B: 0}, Inverted: true, Enabled: true},
						2: {MasterDimmer: full, Color: common.Green, Inverted: true, Enabled: true},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Color: common.Green, Inverted: true, Enabled: true},
						1: {MasterDimmer: full, Color: common.Green, Inverted: true, Enabled: true},
						2: {MasterDimmer: full, Color: color.NRGBA{R: 0, G: 0, B: 0}, Inverted: true, Enabled: true},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := invertRGBColorsInSteps(tt.args.steps, tt.args.numberFixtures, tt.args.colors, tt.args.fixtureState); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("invertRGBColors() = %+v, want %+v", got, tt.want)

				fmt.Printf("++++++++++++++ 3 Inverted GOT ++++++++++++++++++++\n")

				for stepNumber, step := range got {

					fmt.Printf("Step:%d ============================\n", stepNumber)

					for fixtureNumber := 0; fixtureNumber < len(step.Fixtures); fixtureNumber++ {
						fixture := step.Fixtures[fixtureNumber]
						fmt.Printf("\fixture %d master %d inverted %t color:%+v\n", fixtureNumber, fixture.MasterDimmer, fixture.Inverted, fixture.Color)
					}
				}
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
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: color.NRGBA{R: 255, G: 255, B: 255}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: color.NRGBA{R: 255, G: 255, B: 255}},
							2: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 255, Color: color.NRGBA{R: 255, G: 255, B: 255}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: color.NRGBA{R: 255, G: 255, B: 255}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 255, Color: color.NRGBA{R: 255, G: 255, B: 255}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: color.NRGBA{R: 255, G: 255, B: 255}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: color.NRGBA{R: 255, G: 255, B: 255}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: color.NRGBA{R: 255, G: 255, B: 255}},
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
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: color.NRGBA{R: 255, G: 255, B: 255}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: color.NRGBA{R: 255, G: 255, B: 255}},
							2: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						},
					},
					// {
					// 	Fixtures: map[int]common.Fixture{
					// 		0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					// 		1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					// 		2: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 255, Color: color.NRGBA{R: 255, G: 255, B: 255}},
					// 		3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					// 		4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					// 		5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					// 		6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					// 		7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
					// 	},
					// },
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: color.NRGBA{R: 255, G: 255, B: 255}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: color.NRGBA{R: 255, G: 255, B: 255}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: color.NRGBA{R: 255, G: 255, B: 255}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: color.NRGBA{R: 255, G: 255, B: 255}},
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
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: color.NRGBA{R: 255, G: 255, B: 255}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: color.NRGBA{R: 255, G: 255, B: 255}},
							2: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: color.NRGBA{R: 255, G: 255, B: 255}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: color.NRGBA{R: 255, G: 255, B: 255}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 255, Color: color.NRGBA{R: 255, G: 255, B: 255}},
							5: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 255, Color: color.NRGBA{R: 255, G: 255, B: 255}},
							6: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 255, Color: color.NRGBA{R: 255, G: 255, B: 255}},
							7: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 255, Color: color.NRGBA{R: 255, G: 255, B: 255}},
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
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: color.NRGBA{R: 255, G: 255, B: 255}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: color.NRGBA{R: 255, G: 255, B: 255}},
							2: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: color.NRGBA{R: 255, G: 255, B: 255}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: color.NRGBA{R: 255, G: 255, B: 255}},
							4: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
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
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: color.NRGBA{R: 255, G: 255, B: 255}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: color.NRGBA{R: 255, G: 255, B: 255}},
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
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: color.NRGBA{R: 255, G: 255, B: 255}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: color.NRGBA{R: 255, G: 255, B: 255}},
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
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: color.NRGBA{R: 255, G: 255, B: 255}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: color.NRGBA{R: 255, G: 255, B: 255}},
							2: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: color.NRGBA{R: 255, G: 255, B: 255}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: color.NRGBA{R: 255, G: 255, B: 255}},
							4: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: color.NRGBA{R: 255, G: 255, B: 255}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: color.NRGBA{R: 255, G: 255, B: 255}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: color.NRGBA{R: 255, G: 255, B: 255}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: color.NRGBA{R: 255, G: 255, B: 255}},
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
							0: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: color.NRGBA{R: 255, G: 255, B: 255}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: color.NRGBA{R: 255, G: 255, B: 255}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: color.NRGBA{R: 255, G: 255, B: 255}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							1: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							2: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							3: {MasterDimmer: full, Enabled: false, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							4: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							5: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							6: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 0, Color: color.NRGBA{R: 0, G: 0, B: 0}},
							7: {MasterDimmer: full, Enabled: true, Brightness: 0, Shutter: 255, Color: color.NRGBA{R: 255, G: 255, B: 255}},
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
