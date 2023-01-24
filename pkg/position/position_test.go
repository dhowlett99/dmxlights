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
		slopeOn  []int
		slopeOff []int
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

				slopeOn:  []int{1, 50, 255},
				slopeOff: []int{1, 50, 255},
				sequence: common.Sequence{
					Bounce:       false,
					RGBInvert:    false,
					RGBShift:     0,
					RGBSize:      255,
					RGBFade:      0,
					ScannerState: allFixturesEnabled,
					Steps: []common.Step{
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
				slopeOn:  []int{1, 50, 255},
				slopeOff: []int{1, 50, 255},
				sequence: common.Sequence{
					Bounce:       false,
					RGBInvert:    false,
					RGBShift:     8,
					RGBSize:      255,
					RGBFade:      10,
					ScannerState: allFixturesEnabled,
					Steps: []common.Step{
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
				slopeOn:  []int{1, 50, 255},
				slopeOff: []int{255, 50, 1},
				sequence: common.Sequence{
					Bounce:       false,
					RGBInvert:    false,
					RGBShift:     8, // Eight is reversed so creates a shift of 2.
					RGBSize:      255,
					RGBFade:      10,
					ScannerState: allFixturesEnabled,
					Steps: []common.Step{
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
			got, got1 := CalculatePositions(tt.args.sequence, tt.args.slopeOn, tt.args.slopeOff, false, tt.args.sequence.ScannerState)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculatePositions() got = %+v, want %+v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("calculatePositions() got1 = %+v, want %+v", got1, tt.want1)
			}
		})
	}
}

// func TestCalculateScannerChasePositionsAllEnabled(t *testing.T) {

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
// 	}

// 	full := 255
// 	type args struct {
// 		sequence common.Sequence
// 		slopeOn  []int
// 		slopeOff []int
// 	}
// 	tests := []struct {
// 		name  string
// 		args  args
// 		want  map[int]common.Position
// 		want1 int
// 	}{
// 		{
// 			name: "Scanner Basic Check, All enabled shutter fires in sequence",
// 			args: args{
// 				slopeOn:  []int{255},
// 				slopeOff: []int{0},
// 				sequence: common.Sequence{
// 					Type:         "Scanner",
// 					ScannerChase: true,
// 					Bounce:       false,
// 					RGBInvert:    false,
// 					RGBShift:     0,
// 					RGBSize:      255,
// 					RGBFade:      10,
// 					ScannerState: allFixturesEnabled,
// 					Steps: []common.Step{
// 						{
// 							Fixtures: []common.Fixture{
// 								{Shutter: 255, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
// 								{Shutter: 255, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
// 								{Shutter: 255, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
// 								{Shutter: 255, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
// 							},
// 						},
// 						{
// 							Fixtures: []common.Fixture{
// 								{Shutter: 255, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
// 								{Shutter: 255, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
// 								{Shutter: 255, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
// 								{Shutter: 255, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
// 							},
// 						},
// 						{
// 							Fixtures: []common.Fixture{
// 								{Shutter: 255, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
// 								{Shutter: 255, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
// 								{Shutter: 255, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
// 								{Shutter: 255, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
// 							},
// 						},
// 						{
// 							Fixtures: []common.Fixture{
// 								{Shutter: 255, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
// 								{Shutter: 255, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
// 								{Shutter: 255, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
// 								{Shutter: 255, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
// 							},
// 						},
// 					},
// 				},
// 			},
// 			want: map[int]common.Position{
// 				0: {
// 					Fixtures: map[int]common.Fixture{
// 						0: {Shutter: 255, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
// 						1: {Shutter: 0, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
// 						2: {Shutter: 0, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
// 						3: {Shutter: 0, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
// 					},
// 				},
// 				1: {
// 					Fixtures: map[int]common.Fixture{
// 						0: {Shutter: 0, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
// 						1: {Shutter: 255, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
// 						2: {Shutter: 0, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
// 						3: {Shutter: 0, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
// 					},
// 				},
// 				2: {
// 					Fixtures: map[int]common.Fixture{
// 						0: {Shutter: 0, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
// 						1: {Shutter: 0, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
// 						2: {Shutter: 255, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
// 						3: {Shutter: 0, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
// 					},
// 				},
// 				3: {
// 					Fixtures: map[int]common.Fixture{
// 						0: {Shutter: 0, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
// 						1: {Shutter: 0, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
// 						2: {Shutter: 0, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
// 						3: {Shutter: 255, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
// 					},
// 				},
// 			},
// 			want1: 4,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			// Optimisation is turned off for testing.
// 			got, got1 := CalculatePositions(tt.args.sequence, tt.args.slopeOn, tt.args.slopeOff, false, tt.args.sequence.ScannerState)
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("CalculatePositions() got = %+v, want %+v", got, tt.want)
// 			}
// 			if got1 != tt.want1 {
// 				t.Errorf("CalculatePositions() got1 = %+v, want %+v", got1, tt.want1)
// 			}
// 		})
// 	}
// }

// func TestCalculateScannerChasePositionsHalfEnabled(t *testing.T) {

// 	allFixturesEnabled := map[int]common.ScannerState{
// 		0: {
// 			Enabled: true,
// 		},
// 		1: {
// 			Enabled: true,
// 		},
// 		2: {
// 			Enabled: false,
// 		},
// 		3: {
// 			Enabled: false,
// 		},
// 	}

// 	full := 255
// 	type args struct {
// 		sequence common.Sequence
// 		slopeOn  []int
// 		slopeOff []int
// 	}
// 	tests := []struct {
// 		name  string
// 		args  args
// 		want  map[int]common.Position
// 		want1 int
// 	}{
// 		{
// 			name: "Scanner Basic Check, All enabled shutter fires in sequence",
// 			args: args{
// 				slopeOn:  []int{255},
// 				slopeOff: []int{0},
// 				sequence: common.Sequence{
// 					Type:         "Scanner",
// 					ScannerChase: true,
// 					Bounce:       false,
// 					RGBInvert:    false,
// 					RGBShift:     0,
// 					RGBSize:      255,
// 					RGBFade:      10,
// 					ScannerState: allFixturesEnabled,
// 					Steps: []common.Step{
// 						{
// 							Fixtures: []common.Fixture{
// 								{Shutter: 255, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
// 								{Shutter: 255, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
// 								{Shutter: 255, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
// 								{Shutter: 255, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
// 							},
// 						},
// 						{
// 							Fixtures: []common.Fixture{
// 								{Shutter: 255, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
// 								{Shutter: 255, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
// 								{Shutter: 255, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
// 								{Shutter: 255, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
// 							},
// 						},
// 						{
// 							Fixtures: []common.Fixture{
// 								{Shutter: 255, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
// 								{Shutter: 255, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
// 								{Shutter: 255, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
// 								{Shutter: 255, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
// 							},
// 						},
// 						{
// 							Fixtures: []common.Fixture{
// 								{Shutter: 255, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
// 								{Shutter: 255, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
// 								{Shutter: 255, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
// 								{Shutter: 255, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
// 							},
// 						},
// 					},
// 				},
// 			},
// 			want: map[int]common.Position{
// 				0: {
// 					Fixtures: map[int]common.Fixture{
// 						0: {Shutter: 255, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
// 						1: {Shutter: 0, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
// 						2: {Shutter: 0, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
// 						3: {Shutter: 0, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
// 					},
// 				},
// 				1: {
// 					Fixtures: map[int]common.Fixture{
// 						0: {Shutter: 0, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
// 						1: {Shutter: 255, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
// 						2: {Shutter: 0, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
// 						3: {Shutter: 0, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
// 					},
// 				},
// 				2: {
// 					Fixtures: map[int]common.Fixture{
// 						0: {Shutter: 255, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
// 						1: {Shutter: 0, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
// 						2: {Shutter: 0, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
// 						3: {Shutter: 0, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
// 					},
// 				},
// 				3: {
// 					Fixtures: map[int]common.Fixture{
// 						0: {Shutter: 0, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
// 						1: {Shutter: 255, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
// 						2: {Shutter: 0, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
// 						3: {Shutter: 0, Pan: 128, Tilt: 128, MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
// 					},
// 				},
// 			},
// 			want1: 4,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			// Optimisation is turned off for testing.
// 			got, got1 := CalculatePositions(tt.args.sequence, tt.args.slopeOn, tt.args.slopeOff, false, tt.args.sequence.ScannerState)
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("calculatePositions() got = %+v, want %+v", got, tt.want)
// 			}
// 			if got1 != tt.want1 {
// 				t.Errorf("calculatePositions() got1 = %+v, want %+v", got1, tt.want1)
// 			}
// 		})
// 	}
// }
