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

var allFixturesEnabled = map[int]common.ScannerState{
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

	allFixturesEnabled := map[int]common.ScannerState{
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

	full := 255
	type args struct {
		sequence common.Sequence
	}
	tests := []struct {
		name  string
		args  args
		want  map[int]common.Position
		want1 int
	}{
		{
			name: "golden path - common par fixture RGB",
			args: args{
				sequence: common.Sequence{
					FadeUpAndDown: []int{1, 50, 255},
					FadeDownAndUp: []int{1, 50, 255},
					Optimisation:  false,
					Bounce:        false,
					RGBInvert:     false,
					RGBShift:      0,
					RGBSize:       255,
					RGBFade:       0,
					ScannerState:  allFixturesEnabled,
					RGBSteps: []common.Step{
						{
							Fixtures: []common.Fixture{
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
							},
						},
						{
							Fixtures: []common.Fixture{
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
							},
						},
						{
							Fixtures: []common.Fixture{
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
							},
						},
						{
							Fixtures: []common.Fixture{
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
							},
						},
						{
							Fixtures: []common.Fixture{
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
							},
						},
						{
							Fixtures: []common.Fixture{
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
							},
						},
						{
							Fixtures: []common.Fixture{
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
							},
						},
						{
							Fixtures: []common.Fixture{
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
							},
						},
					},
				},
			},
			want: map[int]common.Position{
				0: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				1: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				2: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				3: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				4: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				5: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				6: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				7: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				8: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				9: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				10: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				11: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				12: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				13: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				14: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				15: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				16: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				17: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				18: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				19: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				20: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				21: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				22: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				23: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
			},

			want1: 24,
		},
		{
			name: "Shift1 - common par fixture RGB",
			args: args{
				sequence: common.Sequence{
					FadeUpAndDown: []int{1, 50, 255},
					FadeDownAndUp: []int{1, 50, 255},
					Optimisation:  false,
					Bounce:        false,
					RGBInvert:     false,
					RGBShift:      8,
					RGBSize:       255,
					RGBFade:       10,
					ScannerState:  allFixturesEnabled,
					RGBSteps: []common.Step{
						{
							Fixtures: []common.Fixture{
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
							},
						},
						{
							Fixtures: []common.Fixture{
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
							},
						},
						{
							Fixtures: []common.Fixture{
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
							},
						},
						{
							Fixtures: []common.Fixture{
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
							},
						},
					},
				},
			},
			want: map[int]common.Position{
				0: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				1: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				2: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				3: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				4: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				5: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				6: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				7: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				8: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
			},

			want1: 9,
		},
		{
			name: "Shift1 - Not Inverted common par fixture RGB",
			args: args{
				sequence: common.Sequence{
					FadeUpAndDown: []int{1, 50, 255},
					FadeDownAndUp: []int{255, 50, 1},
					Optimisation:  false,
					Bounce:        false,
					RGBInvert:     false,
					RGBShift:      8, // Eight is reversed so creates a shift of 2.
					RGBSize:       255,
					RGBFade:       10,
					ScannerState:  allFixturesEnabled,
					RGBSteps: []common.Step{
						{
							Fixtures: []common.Fixture{
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
							},
						},
						{
							Fixtures: []common.Fixture{
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
							},
						},
						{
							Fixtures: []common.Fixture{
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
							},
						},
						{
							Fixtures: []common.Fixture{
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
							},
						},
					},
				},
			},
			want: map[int]common.Position{
				0: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				1: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				2: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				3: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				4: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				5: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				6: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				7: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				8: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
			},
			want1: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Optimisation is turned off for testing.
			got, got1 := CalculatePositions(tt.args.sequence)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculatePositions() got = %+v, want %+v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("calculatePositions() got1 = %+v, want %+v", got1, tt.want1)
			}
		})
	}
}

func TestCalculatePositions2(t *testing.T) {

	full := 255
	type args struct {
		sequence common.Sequence
	}
	tests := []struct {
		name  string
		args  args
		want  map[int]common.Position
		want1 int
	}{
		{
			name: "pairs patten test - common par fixture RGB",
			args: args{
				sequence: common.Sequence{
					FadeUpAndDown: []int{1, 50, 255, 255, 50, 1},
					FadeDownAndUp: []int{255, 50, 1, 1, 50, 255},
					Optimisation:  false,
					Bounce:        false,
					RGBInvert:     false,
					RGBShift:      10,
					RGBSize:       255,
					RGBFade:       1,
					ScannerState:  allFixturesEnabled,
					RGBSteps: []common.Step{
						{
							Fixtures: []common.Fixture{
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							},
						},
						{
							Fixtures: []common.Fixture{
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
							},
						},
						{
							Fixtures: []common.Fixture{
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							},
						},
						{
							Fixtures: []common.Fixture{
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
							},
						},
						{
							Fixtures: []common.Fixture{
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							},
						},
						{
							Fixtures: []common.Fixture{
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
							},
						},
						{
							Fixtures: []common.Fixture{
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							},
						},
						{
							Fixtures: []common.Fixture{
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
							},
						},
					},
				},
			},
			want: map[int]common.Position{
				0: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
			},

			want1: 24,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Optimisation is turned off for testing.
			_, got1 := CalculatePositions(tt.args.sequence)
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("calculatePositions() got = %+v, want %+v", got, tt.want)
			// }
			if got1 != tt.want1 {
				t.Errorf("calculatePositions() got1 = %+v, want %+v", got1, tt.want1)
			}
		})
	}
}

func Test_calculateScannerPositions(t *testing.T) {

	full := 255
	type args struct {
		sequence common.Sequence
	}
	tests := []struct {
		name string
		args args
		want map[int]common.Position
	}{

		{
			name: "Scanner case, both scanners doing same things.",
			args: args{
				sequence: common.Sequence{
					Bounce:        false,
					FadeUpAndDown: []int{255},
					FadeDownAndUp: []int{0},
					Optimisation:  false,
					ScannerState:  allFixturesEnabled,
					RGBSteps: []common.Step{
						{

							Fixtures: []common.Fixture{
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 255},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 255},
							},
						},
						{

							Fixtures: []common.Fixture{
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 50, Tilt: 200},
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 50, Tilt: 200},
							},
						},
						{

							Fixtures: []common.Fixture{
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 100, Tilt: 150},
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 100, Tilt: 150},
							},
						},
						{

							Fixtures: []common.Fixture{
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 36, Shutter: 255, Pan: 150, Tilt: 100},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 36, Shutter: 255, Pan: 150, Tilt: 100},
							},
						},
						{

							Fixtures: []common.Fixture{
								{MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 200, Tilt: 50},
								{MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 200, Tilt: 50},
							},
						},
						{

							Fixtures: []common.Fixture{
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 190, B: 255}}, Gobo: 36, Shutter: 255, Pan: 255, Tilt: 0},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 190, B: 255}}, Gobo: 36, Shutter: 255, Pan: 255, Tilt: 0},
							},
						},
						{

							Fixtures: []common.Fixture{
								{MasterDimmer: full, Colors: []common.Color{{R: 145, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 0},
								{MasterDimmer: full, Colors: []common.Color{{R: 145, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 0},
							},
						},
					},
				},
			},
			want: map[int]common.Position{
				0: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 255, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
						1: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 255, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
					},
				},
				1: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 50, Tilt: 200, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
						1: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 50, Tilt: 200, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
					},
				},
				2: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 100, Tilt: 150, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
						1: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 100, Tilt: 150, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
					},
				},
				3: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 255, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 150, Tilt: 100, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
						1: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 255, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 150, Tilt: 100, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
					},
				},
				4: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 100, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 200, Tilt: 50, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
						1: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 100, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 200, Tilt: 50, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
					},
				},
				5: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 190, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 255, Tilt: 0, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
						1: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 190, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 255, Tilt: 0, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
					},
				},
				6: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 145, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
						1: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 145, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
					},
				},
			},
		},
		{
			name: "Scanner case, both scanners doing different things.",
			args: args{
				sequence: common.Sequence{
					Bounce:        false,
					FadeUpAndDown: []int{255},
					FadeDownAndUp: []int{0},
					Optimisation:  false,
					ScannerState:  allFixturesEnabled,
					RGBSteps: []common.Step{
						{

							Fixtures: []common.Fixture{
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 255},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 128, Tilt: 255},
							},
						},
						{

							Fixtures: []common.Fixture{
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 50, Tilt: 200},
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 128, Tilt: 200},
							},
						},
					},
				},
			},

			want: map[int]common.Position{
				0: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 255, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
						1: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 128, Tilt: 255, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
					},
				},

				1: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 50, Tilt: 200, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
						1: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 128, Tilt: 200, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
					},
				},
			},
		},

		{
			name: "Scanner case, one set of instruction in a pattern should create one set of positions.",
			args: args{
				sequence: common.Sequence{
					Bounce:        false,
					FadeUpAndDown: []int{255},
					FadeDownAndUp: []int{0},
					Optimisation:  false,
					ScannerState:  allFixturesEnabled,
					RGBSteps: []common.Step{
						{

							Fixtures: []common.Fixture{
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 255},
							},
						},
						{

							Fixtures: []common.Fixture{
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 50, Tilt: 200},
							},
						},
					},
				},
			},
			want: map[int]common.Position{

				0: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 255, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
					},
				},
				1: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 50, Tilt: 200, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
					},
				},
			},
		},

		{
			name: "Pairs case",
			args: args{
				sequence: common.Sequence{
					Bounce:        false,
					FadeUpAndDown: []int{255},
					FadeDownAndUp: []int{0},
					Optimisation:  false,
					ScannerState:  allFixturesEnabled,
					RGBSteps: []common.Step{
						{
							Fixtures: []common.Fixture{
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							},
						},
						{
							Fixtures: []common.Fixture{
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
							},
						},
					},
				},
			},
			want: map[int]common.Position{

				0: {
					Fixtures: map[int]common.Fixture{

						0: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				1: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
			},
		},

		{
			name: "Scanners inverted no bounce",
			args: args{
				sequence: common.Sequence{
					Bounce:        false,
					ScannerInvert: true,
					FadeUpAndDown: []int{255},
					FadeDownAndUp: []int{0},
					Optimisation:  false,
					ScannerState:  allFixturesEnabled,
					RGBSteps: []common.Step{
						{
							Fixtures: []common.Fixture{
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Pan: 0, Tilt: 0, Gobo: 1, Shutter: 255},
							},
						},
						{
							Fixtures: []common.Fixture{
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Pan: 1, Tilt: 1, Gobo: 1, Shutter: 255},
							},
						},
						{
							Fixtures: []common.Fixture{
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Pan: 2, Tilt: 2, Gobo: 1, Shutter: 255},
							},
						},
					},
				},
			},

			want: map[int]common.Position{
				0: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 2, Tilt: 2, Shutter: 255, Rotate: 0, Music: 0, Gobo: 1, Program: 0},
					},
				},
				1: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 1, Tilt: 1, Shutter: 255, Rotate: 0, Music: 0, Gobo: 1, Program: 0},
					},
				},
				2: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 255, Rotate: 0, Music: 0, Gobo: 1, Program: 0},
					},
				},
			},
		},

		{
			name: "Scanners inverted with bounce",
			args: args{
				sequence: common.Sequence{
					Bounce:        true,
					ScannerInvert: true,
					FadeUpAndDown: []int{255},
					FadeDownAndUp: []int{0},
					Optimisation:  false,
					ScannerState:  allFixturesEnabled,
					RGBSteps: []common.Step{
						{
							Fixtures: []common.Fixture{
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Pan: 0, Tilt: 0, Gobo: 1, Shutter: 255},
							},
						},
						{
							Fixtures: []common.Fixture{
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Pan: 1, Tilt: 1, Gobo: 1, Shutter: 255},
							},
						},
						{
							Fixtures: []common.Fixture{
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Pan: 2, Tilt: 2, Gobo: 1, Shutter: 255},
							},
						},
					},
				},
			},
			want: map[int]common.Position{
				0: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 1, Shutter: 255, Pan: 2, Tilt: 2},
					},
				},

				1: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Gobo: 1, Shutter: 255, Pan: 1, Tilt: 1},
					},
				},
				2: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 1, Shutter: 255, Pan: 0, Tilt: 0},
					},
				},
				3: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 1, Shutter: 255, Pan: 0, Tilt: 0},
					},
				},
				4: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Gobo: 1, Shutter: 255, Pan: 1, Tilt: 1},
					},
				},
				5: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 1, Shutter: 255, Pan: 2, Tilt: 2},
					},
				},
			},
		},
		{
			name: "Multicolored Patten",
			args: args{
				sequence: common.Sequence{
					Bounce:        false,
					ScannerInvert: false,
					FadeUpAndDown: []int{255},
					FadeDownAndUp: []int{0},
					Optimisation:  false,
					ScannerState:  allFixturesEnabled,
					RGBSteps: []common.Step{
						{
							Fixtures: []common.Fixture{
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}},
							},
						},
						{
							Fixtures: []common.Fixture{
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
							},
						},
						{
							Fixtures: []common.Fixture{
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}},
							},
						},
						{
							Fixtures: []common.Fixture{
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}},
							},
						},
						{
							Fixtures: []common.Fixture{
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
							},
						},
						{
							Fixtures: []common.Fixture{
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}},
							},
						},
						{
							Fixtures: []common.Fixture{

								{MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
							},
						},
						{
							Fixtures: []common.Fixture{

								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
								{MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}},
							},
						},
					},
				},
			},
			want: map[int]common.Position{
				// Start of want.
				0: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 111, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 255, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 100, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},

				1: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 111, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 255, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 100, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},

				2: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 255, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 100, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 111, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},

				3: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 255, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 100, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 111, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				4: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 255, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 100, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 111, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				5: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 100, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 111, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 255, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				6: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 100, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 111, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 255, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				7: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 111, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 255, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: 255, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 100, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				// End of Want
			},
		},

		// End of test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := CalculatePositions(tt.args.sequence); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %+v", got)
				t.Errorf("want =%+v", tt.want)
			}
		})
	}
}

func TestAddScannerPositions(t *testing.T) {

	positions := make(map[int]common.Position)

	// First position.
	newPosition1 := common.Position{}
	newPosition1.Fixtures = make(map[int]common.Fixture)

	// Fixture 1.
	newFixture1 := common.Fixture{
		Shutter: 255,
		Colors:  []common.Color{{R: 0, G: 1, B: 0}},
	}
	newPosition1.Fixtures[0] = newFixture1

	// Fixture 2.
	newFixture2 := common.Fixture{

		Shutter: 255,
		Colors:  []common.Color{{R: 0, G: 50, B: 0}},
	}
	newPosition1.Fixtures[1] = newFixture2

	// Fixture 3.
	newFixture3 := common.Fixture{
		Shutter: 255,
		Colors:  []common.Color{{R: 0, G: 255, B: 0}},
	}
	newPosition1.Fixtures[2] = newFixture3

	// Add first position.
	positions[0] = newPosition1

	newPosition2 := common.Position{}
	newPosition2.Fixtures = make(map[int]common.Fixture)

	newFixture1 = common.Fixture{
		Shutter: 255,
		Colors:  []common.Color{{R: 0, G: 1, B: 0}},
	}
	newPosition2.Fixtures[0] = newFixture1

	newFixture2 = common.Fixture{
		Shutter: 255,
		Colors:  []common.Color{{R: 0, G: 50, B: 0}},
	}
	newPosition2.Fixtures[1] = newFixture2

	newFixture3 = common.Fixture{
		Shutter: 255,
		Colors:  []common.Color{{R: 0, G: 255, B: 0}},
	}
	newPosition2.Fixtures[2] = newFixture3

	// Add Second position.
	positions[1] = newPosition2

	newPosition3 := common.Position{}
	newPosition3.Fixtures = make(map[int]common.Fixture)

	newFixture1 = common.Fixture{
		Shutter: 255,
		Colors:  []common.Color{{R: 0, G: 1, B: 0}},
	}
	newPosition3.Fixtures[0] = newFixture1

	newFixture2 = common.Fixture{
		Shutter: 255,
		Colors:  []common.Color{{R: 0, G: 50, B: 0}},
	}
	newPosition3.Fixtures[1] = newFixture2

	newFixture3 = common.Fixture{
		Shutter: 255,
		Colors:  []common.Color{{R: 0, G: 255, B: 0}},
	}
	newPosition3.Fixtures[2] = newFixture3

	// Add third position.
	positions[2] = newPosition3

	type args struct {
		scannerPattern common.Pattern
		positionsIn    map[int]common.Position
	}
	tests := []struct {
		name string
		args args
		want map[int]common.Position
	}{
		{
			name: "add 12 scanner positions to a 10 point curver",
			args: args{
				positionsIn: positions,
				scannerPattern: common.Pattern{
					Name:  "test",
					Label: "test",
					Steps: []common.Step{
						{
							Fixtures: []common.Fixture{
								{
									Name: "scanner1",
									Pan:  1,
									Tilt: 1,
								},
								{
									Name: "scanner2",
									Pan:  1,
									Tilt: 1,
								},
								{
									Name: "scanner3",
									Pan:  1,
									Tilt: 1,
								},
							},
						},
						{
							Fixtures: []common.Fixture{
								{
									Name: "scanner1",
									Pan:  2,
									Tilt: 2,
								},
								{
									Name: "scanner2",
									Pan:  2,
									Tilt: 2,
								},
								{
									Name: "scanner3",
									Pan:  2,
									Tilt: 2,
								},
							},
						},
						{
							Fixtures: []common.Fixture{
								{
									Name: "scanner1",
									Pan:  3,
									Tilt: 3,
								},
								{
									Name: "scanner2",
									Pan:  3,
									Tilt: 3,
								},
								{
									Name: "scanner3",
									Pan:  3,
									Tilt: 3,
								},
							},
						},
					},
				},
			},
			want: map[int]common.Position{
				0: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 0, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 1, Tilt: 1, Shutter: 255, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: 0, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 1, Tilt: 1, Shutter: 255, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: 0, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 1, Tilt: 1, Shutter: 255, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				1: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 0, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 2, Tilt: 2, Shutter: 255, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: 0, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 2, Tilt: 2, Shutter: 255, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: 0, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 2, Tilt: 2, Shutter: 255, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				2: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: 0, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 3, Tilt: 3, Shutter: 255, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: 0, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 3, Tilt: 3, Shutter: 255, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: 0, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 3, Tilt: 3, Shutter: 255, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddScannerPositions(tt.args.scannerPattern, tt.args.positionsIn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddScannerPositions() = got \n%+v", got)
				t.Errorf("AddScannerPositions() = want \n%+v", tt.want)
			}
		})
	}
}
