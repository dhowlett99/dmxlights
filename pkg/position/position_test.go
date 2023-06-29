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

func TestCalculateRGBPositionsSimpleGreenChase(t *testing.T) {

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
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Shutter: 255, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
			},

			want: map[int]common.Position{
				0: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				1: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				2: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				3: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				4: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				5: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				6: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				7: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				8: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				9: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				10: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				11: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				12: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				13: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				14: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				15: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				16: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				17: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				18: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				19: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				20: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				21: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				22: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				23: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				24: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				25: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				26: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				27: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				28: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				29: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				30: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				31: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				32: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				33: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				34: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				35: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				36: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				37: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				38: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				39: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				40: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				41: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				42: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				43: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				44: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				45: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				46: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				47: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
			},
			want1: 48,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			fadeColors, numberFixtures, totalNumberOfSteps := CalculatePositions(tt.steps, tt.args.sequence, tt.args.scanner, 0)
			positions, numberPositions := AssemblePositions(fadeColors, numberFixtures, totalNumberOfSteps, false)
			if !reflect.DeepEqual(positions, tt.want) {
				t.Errorf("CalculatePositions() got = %v, want %v", positions, tt.want)

				for positionNumber := 0; positionNumber < len(positions); positionNumber++ {

					fmt.Printf("Position:%d ============================\n", positionNumber)
					position := positions[positionNumber]
					for fixtureNumber := 0; fixtureNumber < len(position.Fixtures); fixtureNumber++ {

						fixture := position.Fixtures[fixtureNumber]
						fmt.Printf("Fixture:%d color:%+v\n", fixtureNumber, fixture.Colors[0])
					}

				}
			}
			if numberPositions != tt.want1 {
				t.Errorf("CalculatePositions() got1 = %v, want %v", numberPositions, tt.want1)
			}
		})
	}
}

func TestCalculateMulticoloredPatten(t *testing.T) {

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
					Bounce:        false,
					ScannerInvert: false,
					FadeUp:        []int{255},
					//FadeDown:              []int{0},
					Optimisation:          false,
					FixtureState:          allFixturesEnabled,
					EnabledNumberFixtures: 8,
				},
			},
			steps: []common.Step{
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{

						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{

						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 111, B: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}},
					},
				},
			},

			want: map[int]common.Position{
				// Start of want.
				0: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 111, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 255, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 100, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},

				1: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 111, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 255, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 100, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},

				2: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 255, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 100, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 111, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},

				3: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 255, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 100, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 111, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				4: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 255, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 100, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 111, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				5: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 100, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 111, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 255, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				6: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 100, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 111, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 255, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				7: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 111, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 255, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 100, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				// End of Want
			},
			want1: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			fadeColors, numberFixtures, totalNumberOfSteps := CalculatePositions(tt.steps, tt.args.sequence, true, 0)
			positions, numberPositions := AssemblePositions(fadeColors, numberFixtures, totalNumberOfSteps, false)
			if !reflect.DeepEqual(positions, tt.want) {
				t.Errorf("CalculatePositions() got = %v, want %v", positions, tt.want)

				for positionNumber := 0; positionNumber < len(positions); positionNumber++ {

					fmt.Printf("Position:%d ============================\n", positionNumber)
					position := positions[positionNumber]
					for fixtureNumber := 0; fixtureNumber < len(position.Fixtures); fixtureNumber++ {

						fixture := position.Fixtures[fixtureNumber]
						fmt.Printf("Fixture:%d R:%d\n", fixtureNumber, fixture.Colors[0].R)
					}

				}
			}
			if numberPositions != tt.want1 {
				t.Errorf("CalculatePositions() got1 = %v, want %v", numberPositions, tt.want1)
			}
		})
	}
}

func TestCalculateShift8(t *testing.T) {

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
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
			},

			want: map[int]common.Position{
				0: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				1: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				2: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				3: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				4: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				5: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				6: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				7: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				8: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
			},

			want1: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fadeColors, numberFixtures, totalNumberOfSteps := CalculatePositions(tt.steps, tt.args.sequence, true, 0)
			positions, numberPositions := AssemblePositions(fadeColors, numberFixtures, totalNumberOfSteps, false)
			if !reflect.DeepEqual(positions, tt.want) {
				t.Errorf("CalculatePositions() got = %v, want %v", positions, tt.want)

				fmt.Printf(" ================== Want =====================\n")
				for positionNumber := 0; positionNumber < len(tt.want); positionNumber++ {

					fmt.Printf("Position:%d ============================\n", positionNumber)
					position := tt.want[positionNumber]
					for fixtureNumber := 0; fixtureNumber < len(position.Fixtures); fixtureNumber++ {

						fixture := position.Fixtures[fixtureNumber]
						fmt.Printf("Fixture:%d R:%d G:%d B:%d\n", fixtureNumber, fixture.Colors[0].R, fixture.Colors[0].G, fixture.Colors[0].B)
					}

				}

				fmt.Printf(" ================== Got =====================\n")
				for positionNumber := 0; positionNumber < len(positions); positionNumber++ {

					fmt.Printf("Position:%d ============================\n", positionNumber)
					position := positions[positionNumber]
					for fixtureNumber := 0; fixtureNumber < len(position.Fixtures); fixtureNumber++ {

						fixture := position.Fixtures[fixtureNumber]
						fmt.Printf("Fixture:%d R:%d G:%d B:%d\n", fixtureNumber, fixture.Colors[0].R, fixture.Colors[0].G, fixture.Colors[0].B)
					}

				}
			}
			if numberPositions != tt.want1 {
				t.Errorf("CalculatePositions() got1 = %v, want %v", numberPositions, tt.want1)
			}
		})
	}
}

func TestCalculateShift1(t *testing.T) {

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
			name: "Shift1 - Not Inverted common par fixture RGB",
			args: args{
				scanner: false,
				sequence: common.Sequence{
					FadeUp:                []int{1, 50, 255},
					FadeDown:              []int{255, 50, 1},
					Optimisation:          false,
					Bounce:                false,
					RGBInvert:             false,
					RGBShift:              8, // Eight is reversed so creates a shift of 2.
					ScannerInvert:         false,
					FixtureState:          allFixturesEnabled,
					EnabledNumberFixtures: 4,
				},
			},
			steps: []common.Step{
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
			},

			want: map[int]common.Position{
				0: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				1: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				2: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				3: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				4: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				5: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				6: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				7: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
				8: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
			},
			want1: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Optimisation is turned off for testing.

			fadeColors, numberFixtures, totalNumberOfSteps := CalculatePositions(tt.steps, tt.args.sequence, true, 0)
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

func TestCalculatePairsPatten(t *testing.T) {

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
			name: "pairs patten test - common par fixture RGB",
			args: args{
				scanner: false,
				sequence: common.Sequence{
					FadeUp:                []int{1, 50, 255},
					FadeDown:              []int{255, 50, 1},
					Optimisation:          false,
					Bounce:                false,
					RGBInvert:             false,
					RGBShift:              10,
					FixtureState:          allFixturesEnabled,
					EnabledNumberFixtures: 2,
				},
			},
			steps: []common.Step{
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					},
				},
			},
			want: map[int]common.Position{
				0: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 1, B: 0, W: 0, A: 0, UV: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Pan: 0, Tilt: 0, Shutter: 255, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0}}},
					},
				},
			},

			want1: 24,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Optimisation is turned off for testing.
			fadeColors, numberFixtures, totalNumberOfSteps := CalculatePositions(tt.steps, tt.args.sequence, true, 0)
			_, numberPositions := AssemblePositions(fadeColors, numberFixtures, totalNumberOfSteps, false)

			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("calculatePositions() got = %+v, want %+v", got, tt.want)
			// }
			if numberPositions != tt.want1 {
				t.Errorf("calculatePositions() numberPositions = %+v, want %+v", numberPositions, tt.want1)
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
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 255},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 255},
					},
				},
				{

					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 50, Tilt: 200},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 50, Tilt: 200},
					},
				},
				{

					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 100, Tilt: 150},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 100, Tilt: 150},
					},
				},
				{

					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 36, Shutter: 255, Pan: 150, Tilt: 100},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 36, Shutter: 255, Pan: 150, Tilt: 100},
					},
				},
				{

					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 200, Tilt: 50},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 200, Tilt: 50},
					},
				},
				{

					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 190, B: 255}}, Gobo: 36, Shutter: 255, Pan: 255, Tilt: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 190, B: 255}}, Gobo: 36, Shutter: 255, Pan: 255, Tilt: 0},
					},
				},
				{

					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 145, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 145, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 0},
					},
				},
			},

			want: map[int]common.Position{
				0: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 255, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 255, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
					},
				},
				1: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 50, Tilt: 200, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 50, Tilt: 200, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
					},
				},
				2: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 100, Tilt: 150, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 100, Tilt: 150, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
					},
				},
				3: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 255, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 150, Tilt: 100, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 255, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 150, Tilt: 100, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
					},
				},
				4: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 100, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 200, Tilt: 50, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 100, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 200, Tilt: 50, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
					},
				},
				5: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 190, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 255, Tilt: 0, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 190, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 255, Tilt: 0, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
					},
				},
				6: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 145, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 145, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fadeColors, numberFixtures, totalNumberOfSteps := CalculatePositions(tt.steps, tt.args.sequence, true, 0)
			positions, _ := AssemblePositions(fadeColors, numberFixtures, totalNumberOfSteps, false)
			if !reflect.DeepEqual(positions, tt.want) {
				t.Errorf("got = %+v", positions)
				t.Errorf("want =%+v", tt.want)
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
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 255},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 128, Tilt: 255},
					},
				},
				{

					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 50, Tilt: 200},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 128, Tilt: 200},
					},
				},
			},

			want: map[int]common.Position{
				0: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 255, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 128, Tilt: 255, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
					},
				},

				1: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 50, Tilt: 200, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 128, Tilt: 200, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
					},
				},
			},
		},

		// End of test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fadeColors, numberFixtures, totalNumberOfSteps := CalculatePositions(tt.steps, tt.args.sequence, true, 0)
			positions, _ := AssemblePositions(fadeColors, numberFixtures, totalNumberOfSteps, false)
			if !reflect.DeepEqual(positions, tt.want) {
				t.Errorf("got = %+v", positions)
				t.Errorf("want =%+v", tt.want)
			}
		})
	}
}

func Test_calculateScannerCase(t *testing.T) {

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
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 255},
					},
				},
				{

					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 50, Tilt: 200},
					},
				},
			},

			want: map[int]common.Position{

				0: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 255, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
					},
				},
				1: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 50, Tilt: 200, Shutter: 255, Rotate: 0, Music: 0, Gobo: 36, Program: 0},
					},
				},
			},
		},

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
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					},
				},
			},

			want: map[int]common.Position{

				0: {
					Fixtures: map[int]common.Fixture{

						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
				1: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						3: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						4: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						5: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						6: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
						7: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 0, Rotate: 0, Music: 0, Gobo: 0, Program: 0},
					},
				},
			},
		},

		{
			name: "Scanners inverted no bounce",
			args: args{
				scanner: true,
				sequence: common.Sequence{
					Bounce:                false,
					ScannerInvert:         true,
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
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Pan: 0, Tilt: 0, Gobo: 1, Shutter: 255},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Pan: 1, Tilt: 1, Gobo: 1, Shutter: 255},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Pan: 2, Tilt: 2, Gobo: 1, Shutter: 255},
					},
				},
			},

			want: map[int]common.Position{
				0: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 0, B: 255, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 2, Tilt: 2, Shutter: 255, Rotate: 0, Music: 0, Gobo: 1, Program: 0},
					},
				},
				1: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 1, Tilt: 1, Shutter: 255, Rotate: 0, Music: 0, Gobo: 1, Program: 0},
					},
				},
				2: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, ScannerColor: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, Colors: []common.Color{{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}}, Pan: 0, Tilt: 0, Shutter: 255, Rotate: 0, Music: 0, Gobo: 1, Program: 0},
					},
				},
			},
		},

		{
			name: "Scanners inverted with bounce",
			args: args{
				scanner: true,
				sequence: common.Sequence{
					Bounce:                true,
					ScannerInvert:         true,
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
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Pan: 0, Tilt: 0, Gobo: 1, Shutter: 255},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Pan: 1, Tilt: 1, Gobo: 1, Shutter: 255},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Pan: 2, Tilt: 2, Gobo: 1, Shutter: 255},
					},
				},
			},

			want: map[int]common.Position{
				0: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 1, Shutter: 255, Pan: 2, Tilt: 2},
					},
				},

				1: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Gobo: 1, Shutter: 255, Pan: 1, Tilt: 1},
					},
				},
				2: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 1, Shutter: 255, Pan: 0, Tilt: 0},
					},
				},
				3: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 1, Shutter: 255, Pan: 0, Tilt: 0},
					},
				},
				4: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}, Gobo: 1, Shutter: 255, Pan: 1, Tilt: 1},
					},
				},
				5: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 1, Shutter: 255, Pan: 2, Tilt: 2},
					},
				},
			},
		},

		// End of test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fadeColors, numberFixtures, totalNumberOfSteps := CalculatePositions(tt.steps, tt.args.sequence, true, 0)
			positions, _ := AssemblePositions(fadeColors, numberFixtures, totalNumberOfSteps, false)
			if !reflect.DeepEqual(positions, tt.want) {
				t.Errorf("got = %+v", positions)
				t.Errorf("want =%+v", tt.want)
			}
		})
	}
}

func TestCalculateStandardPositions(t *testing.T) {

	var full = 255
	type args struct {
		steps    []common.Step
		sequence common.Sequence
		scanner  bool
	}
	tests := []struct {
		name  string
		args  args
		want  map[int]common.Position
		want1 int
	}{
		{
			name: "Standard 3 fixtures forward chase.",
			args: args{
				scanner: false,
				sequence: common.Sequence{
					Bounce:                false,
					ScannerInvert:         false,
					FadeUp:                []int{0, 50, 255},
					FadeDown:              []int{255, 50, 0},
					Optimisation:          false,
					FixtureState:          allFixturesEnabled,
					EnabledNumberFixtures: 3,
					NumberFixtures:        3,
					ScannerChaser:         false,
					RGBShift:              0,
					RGBInvert:             false,
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
			},
			want: map[int]common.Position{
				0: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},

				1: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 50, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				2: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				3: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				4: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 50, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				5: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				6: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				7: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 50, G: 0, B: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				8: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				9: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				10: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 50, G: 0, B: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				11: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				12: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				13: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 50, G: 0, B: 0}}},
					},
				},
				14: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					},
				},
				15: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					},
				},
				16: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 50, G: 0, B: 0}}},
					},
				},
				17: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
			},
			want1: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			fadeColors, numberFixtures, totalNumberOfSteps := CalculatePositions(tt.args.steps, tt.args.sequence, true, 0)
			positions, numberPositions := AssemblePositions(fadeColors, numberFixtures, totalNumberOfSteps, false)
			if !reflect.DeepEqual(positions, tt.want) {
				t.Errorf("CalculatePositions() got = %v, want %v", positions, tt.want)

				for positionNumber := 0; positionNumber < len(positions); positionNumber++ {

					fmt.Printf("Position:%d ============================\n", positionNumber)
					position := positions[positionNumber]
					for fixtureNumber := 0; fixtureNumber < len(position.Fixtures); fixtureNumber++ {

						fixture := position.Fixtures[fixtureNumber]
						fmt.Printf("Fixture:%d R:%d\n", fixtureNumber, fixture.Colors[0].R)
					}

				}
			}
			if numberPositions != tt.want1 {
				t.Errorf("CalculatePositions() got1 = %v, want %v", numberPositions, tt.want1)
			}
		})
	}
}

func TestCalculateInvertedPositions(t *testing.T) {

	var full = 255
	type args struct {
		steps    []common.Step
		sequence common.Sequence
		scanner  bool
	}
	tests := []struct {
		name  string
		args  args
		want  map[int]common.Position
		want1 int
	}{
		{
			name: "Standard 3 fixtures forward chase. Inverted",
			args: args{
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
					RGBInvert:             false,
				},

				steps: []common.Step{
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
							2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
							1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
							1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
							2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
							1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
							2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						},
					},
				},
			},
			want: map[int]common.Position{
				0: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					},
				},

				1: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 50, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					},
				},
				2: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					},
				},
				3: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					},
				},
				4: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 50, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					},
				},
				5: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					},
				},
				6: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					},
				},
				7: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 50, G: 0, B: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					},
				},
				8: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					},
				},
				9: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					},
				},
				10: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 50, G: 0, B: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					},
				},
				11: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					},
				},
				12: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					},
				},
				13: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 50, G: 0, B: 0}}},
					},
				},
				14: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				15: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				16: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 50, G: 0, B: 0}}},
					},
				},
				17: {
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						1: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						2: {MasterDimmer: full, Enabled: true, Brightness: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					},
				},
			},
			want1: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			fadeColors, numberFixtures, totalNumberOfSteps := CalculatePositions(tt.args.steps, tt.args.sequence, true, 0)
			positions, numberPositions := AssemblePositions(fadeColors, numberFixtures, totalNumberOfSteps, false)
			if !reflect.DeepEqual(positions, tt.want) {
				t.Errorf("CalculatePositions() got = %v, want %v", positions, tt.want)

				for positionNumber := 0; positionNumber < len(positions); positionNumber++ {

					fmt.Printf("Position:%d ============================\n", positionNumber)
					position := positions[positionNumber]
					for fixtureNumber := 0; fixtureNumber < len(position.Fixtures); fixtureNumber++ {

						fixture := position.Fixtures[fixtureNumber]
						fmt.Printf("Fixture:%d R:%d\n", fixtureNumber, fixture.Colors[0].R)
					}

				}
			}
			if numberPositions != tt.want1 {
				t.Errorf("CalculatePositions() got1 = %v, want %v", numberPositions, tt.want1)
			}
		})
	}
}

func Test_processDifferentColor(t *testing.T) {

	fadeColors := make(map[int][]common.FixtureBuffer, 10)
	type args struct {
		start         bool
		end           bool
		invert        bool
		fadeColors    map[int][]common.FixtureBuffer
		fixture       common.Fixture
		fixtureNumber int
		color         common.Color
		colorNumber   int
		lastStep      common.Step
		nextStep      common.Step
		sequence      common.Sequence
		shift         int
		patternShift  int
		scanner       bool
	}
	tests := []struct {
		name string
		args args
		want map[int][]common.FixtureBuffer
	}{
		{
			// If color is different from last color and not black.
			// And last color was green, so fade down green first.
			name: "process a single color red, last color was black",
			args: args{
				start:      true,
				end:        false,
				fadeColors: fadeColors,
				sequence: common.Sequence{
					FadeUp:   []int{0, 50, 255},
					FadeOn:   []int{255},
					FadeDown: []int{255, 50, 0},
				},
				fixtureNumber: 0,
				fixture: common.Fixture{
					Colors: []common.Color{
						{
							R: 255,
							G: 0,
							B: 0,
						},
					},
				},
				colorNumber: 0,
				color: common.Color{
					R: 255,
					G: 0,
					B: 0,
				},
				lastStep: common.Step{
					Fixtures: map[int]common.Fixture{
						0: {
							Colors: []common.Color{
								{
									R: 0,
									G: 255,
									B: 0,
								},
							},
						},
					},
				},
			},

			want: map[int][]common.FixtureBuffer{

				0: {
					// Fade Down Green.
					{Color: common.Color{R: 0, G: 255, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 0, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 50, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 0, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 0, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					// Fade Up Red.
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 0, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 0, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 0, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					// Keep Red on for on time.
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 0, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lastColor := tt.args.lastStep.Fixtures[tt.args.fixtureNumber].Colors[tt.args.colorNumber]
			nextColor := tt.args.nextStep.Fixtures[tt.args.fixtureNumber].Colors[tt.args.colorNumber]
			if got := processColor(tt.args.start, tt.args.end, tt.args.invert, tt.args.fadeColors, tt.args.fixture, tt.args.fixtureNumber, tt.args.color, lastColor, nextColor, tt.args.sequence, tt.args.shift, tt.args.patternShift, tt.args.scanner); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("processColor() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_processSameColorNotBlack(t *testing.T) {

	fadeColors := make(map[int][]common.FixtureBuffer, 10)
	type args struct {
		start         bool
		end           bool
		invert        bool
		fadeColors    map[int][]common.FixtureBuffer
		fixture       common.Fixture
		fixtureNumber int
		color         common.Color
		colorNumber   int
		lastStep      common.Step
		nextStep      common.Step
		sequence      common.Sequence
		shift         int
		patternShift  int
		scanner       bool
	}
	tests := []struct {
		name string
		args args
		want map[int][]common.FixtureBuffer
	}{
		{
			// If color is same as last time , play that color out again.
			name: "process a single color red, last color was red",
			args: args{
				start:      false,
				end:        true,
				shift:      10, // inverted to represents shift 0
				fadeColors: fadeColors,
				sequence: common.Sequence{
					FadeUp:   []int{0, 50, 255},
					FadeOn:   []int{255},
					FadeDown: []int{255, 50, 0},
				},
				fixtureNumber: 0,
				// Fixture contains color Red.
				fixture: common.Fixture{
					Colors: []common.Color{
						{
							R: 255,
							G: 0,
							B: 0,
						},
					},
				},
				colorNumber: 0,
				// Color is therefor Red.
				color: common.Color{
					R: 255,
					G: 0,
					B: 0,
				},
				// Last step was also red.
				lastStep: common.Step{
					Fixtures: map[int]common.Fixture{
						0: {
							Colors: []common.Color{
								{
									R: 255,
									G: 0,
									B: 0,
								},
							},
						},
					},
				},
			},

			want: map[int][]common.FixtureBuffer{

				0: {
					// Play out the existing Red
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 0, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 0, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 0, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					// Keep Red on for the on time.
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 0, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					// Keep Red on for the down time.
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 0, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 0, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 0, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lastColor := tt.args.lastStep.Fixtures[tt.args.fixtureNumber].Colors[tt.args.colorNumber]
			nextColor := tt.args.nextStep.Fixtures[tt.args.fixtureNumber].Colors[tt.args.colorNumber]
			if got := processColor(tt.args.start, tt.args.end, tt.args.invert, tt.args.fadeColors, tt.args.fixture, tt.args.fixtureNumber, tt.args.color, lastColor, nextColor, tt.args.sequence, tt.args.shift, tt.args.patternShift, tt.args.scanner); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("processColor() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_processDiffColorBlack(t *testing.T) {

	fadeColors := make(map[int][]common.FixtureBuffer, 10)
	type args struct {
		start         bool
		end           bool
		invert        bool
		fadeColors    map[int][]common.FixtureBuffer
		fixture       common.Fixture
		fixtureNumber int
		color         common.Color
		colorNumber   int
		lastStep      common.Step
		nextStep      common.Step
		sequence      common.Sequence
		shift         int
		patternShift  int
		scanner       bool
	}
	tests := []struct {
		name string
		args args
		want map[int][]common.FixtureBuffer
	}{
		{
			// If color is different from last color and color is a black.
			name: "process a single color black, last color was red",
			args: args{
				start:      false,
				end:        true,
				shift:      10, // inverted to represents shift 0
				fadeColors: fadeColors,
				sequence: common.Sequence{
					FadeUp:   []int{0, 50, 255},
					FadeOn:   []int{255},
					FadeDown: []int{255, 50, 0},
				},
				fixtureNumber: 0,
				// Fixture contains color Black.
				fixture: common.Fixture{
					Colors: []common.Color{
						{
							R: 0,
							G: 0,
							B: 0,
						},
					},
				},
				colorNumber: 0,
				// Color is therefor Black.
				color: common.Color{
					R: 0,
					G: 0,
					B: 0,
				},
				// Last step was also red.
				lastStep: common.Step{
					Fixtures: map[int]common.Fixture{
						0: {
							Colors: []common.Color{
								{
									R: 255,
									G: 0,
									B: 0,
								},
							},
						},
					},
				},
			},

			want: map[int][]common.FixtureBuffer{

				0: {
					// Fade Down Red to Black.
					{Color: common.Color{R: 255, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 0, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 50, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 0, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 0, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					// Keep off for the off time, same as on time.
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 0, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					// Keep off for the fade up time.
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 0, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 0, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
					{Color: common.Color{R: 0, G: 0, B: 0, W: 0, A: 0, UV: 0, Flash: false}, MasterDimmer: 0, Brightness: 255, Gobo: 0, Pan: 0, Tilt: 0, Shutter: 0, Enabled: false},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lastColor := tt.args.lastStep.Fixtures[tt.args.fixtureNumber].Colors[tt.args.colorNumber]
			nextColor := tt.args.nextStep.Fixtures[tt.args.fixtureNumber].Colors[tt.args.colorNumber]
			if got := processColor(tt.args.start, tt.args.end, tt.args.invert, tt.args.fadeColors, tt.args.fixture, tt.args.fixtureNumber, tt.args.color, lastColor, nextColor, tt.args.sequence, tt.args.shift, tt.args.patternShift, tt.args.scanner); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("processColor() got = %v, want %v", got, tt.want)
			}
		})
	}
}
