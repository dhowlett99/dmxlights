// Copyright (C) 2022 dhowlett99.
// This is the dmxlights fixture editor it is attached to a fixture and
// describes the fixtures properties which is then saved in the fixtures.yaml
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
	"math"

	"github.com/dhowlett99/dmxlights/pkg/common"
)

const debug = false

const (
	full = 255
)

func MakeSingleFixtureChase(colors []common.Color) common.Pattern {

	steps := []common.Step{}
	for _, color := range colors {
		fixture := common.Fixture{
			MasterDimmer: full,
			Enabled:      true,
			Colors:       []common.Color{color},
		}
		fixtures := make(map[int]common.Fixture)
		// Create identical four fixtures
		fixtures[0] = fixture
		fixtures[1] = fixture
		fixtures[2] = fixture
		fixtures[3] = fixture
		step := common.Step{
			Fixtures: fixtures,
		}
		steps = append(steps, step)
	}

	single := common.Pattern{
		Name:   "Single",
		Number: 0,
		Label:  "Single.Chase",
		Steps:  steps,
	}
	return single
}

func MakePatterns() map[int]common.Pattern {

	Patterns := make(map[int]common.Pattern)

	standard := common.Pattern{
		Name:   "Chase",
		Number: 0,
		Label:  "Std.Chase",
		Steps: []common.Step{
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Shutter: 255, Brightness: 255, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					1: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					2: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					3: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					4: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					5: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					6: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					7: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					1: {MasterDimmer: full, Enabled: true, Shutter: 255, Brightness: 255, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					2: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					3: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					4: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					5: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					6: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					7: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					1: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					2: {MasterDimmer: full, Enabled: true, Shutter: 255, Brightness: 255, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					3: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					4: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					5: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					6: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					7: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					1: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					2: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					3: {MasterDimmer: full, Enabled: true, Shutter: 255, Brightness: 255, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					4: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					5: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					6: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					7: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					1: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					2: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					3: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					4: {MasterDimmer: full, Enabled: true, Shutter: 255, Brightness: 255, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					5: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					6: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					7: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					1: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					2: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					3: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					4: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					5: {MasterDimmer: full, Enabled: true, Shutter: 255, Brightness: 255, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					6: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					7: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
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
					6: {MasterDimmer: full, Enabled: true, Shutter: 255, Brightness: 255, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					7: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
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
					7: {MasterDimmer: full, Enabled: true, Shutter: 255, Brightness: 255, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
				},
			},
		},
	}

	flash := common.Pattern{
		Name:   "Flash",
		Number: 1,
		Label:  "Flash",
		Steps: []common.Step{
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
	}

	rgbchase := common.Pattern{
		Name:   "RGB Chase",
		Number: 2,
		Label:  "RGB.Chase",
		Steps: []common.Step{
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 0, B: 0}, {R: 0, G: 255, B: 0}, {R: 0, G: 0, B: 255}}},
					1: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
					2: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
					3: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
					4: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
					5: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
					6: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
					7: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
					1: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 0, B: 0}, {R: 0, G: 255, B: 0}, {R: 0, G: 0, B: 255}}},
					2: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
					3: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
					4: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
					5: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
					6: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
					7: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
					1: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
					2: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 0, B: 0}, {R: 0, G: 255, B: 0}, {R: 0, G: 0, B: 255}}},
					3: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
					4: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
					5: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
					6: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
					7: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
					1: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
					2: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
					3: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 0, B: 0}, {R: 0, G: 255, B: 0}, {R: 0, G: 0, B: 255}}},
					4: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
					5: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
					6: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
					7: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
					1: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
					2: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
					3: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
					4: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 0, B: 0}, {R: 0, G: 255, B: 0}, {R: 0, G: 0, B: 255}}},
					5: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
					6: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
					7: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
					1: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
					2: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
					3: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
					4: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
					5: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 0, B: 0}, {R: 0, G: 255, B: 0}, {R: 0, G: 0, B: 255}}},
					6: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
					7: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
					1: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
					2: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
					3: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
					4: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
					5: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
					6: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 0, B: 0}, {R: 0, G: 255, B: 0}, {R: 0, G: 0, B: 255}}},
					7: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
					1: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
					2: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
					3: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
					4: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
					5: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
					6: {MasterDimmer: full, Enabled: true, Colors: []common.Color{}},
					7: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 0, B: 0}, {R: 0, G: 255, B: 0}, {R: 0, G: 0, B: 255}}},
				},
			},
		},
	}

	pairs := common.Pattern{
		Name:   "Pairs",
		Label:  "Pairs",
		Number: 3,
		Steps: []common.Step{
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					1: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					2: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					3: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					4: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					5: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					6: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					7: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					1: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					2: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					3: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					4: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					5: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					6: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					7: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
				},
			},
		},
	}

	inward := common.Pattern{
		Name:   "Inward",
		Label:  "Inward",
		Number: 4,
		Steps: []common.Step{
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					1: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					2: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					3: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					4: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					5: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					6: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					7: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					1: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					2: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					3: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					4: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					5: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					6: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					7: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					1: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					2: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					3: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					4: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					5: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					6: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					7: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					1: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					2: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					3: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					4: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					5: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					6: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					7: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
		},
	}

	colors := common.Pattern{
		Name:   "Color Chase",
		Label:  "Color.Chase",
		Number: 5,
		Steps: []common.Step{
			{ // Step 1, - Red
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					1: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					2: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					3: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					4: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					5: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					6: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					7: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{ // Step 2 - Orange
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					1: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 111, B: 0}}},
					2: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					3: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					4: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					5: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					6: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					7: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{ // Step 3 - Yellow
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					1: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					2: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 255, B: 0}}},
					3: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					4: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					5: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					6: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					7: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{ // Step 4 - Green
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					1: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					2: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					3: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					4: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					5: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					6: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					7: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{ // Step 5 - Cyan
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					1: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					2: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					3: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					4: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 255, B: 255}}},
					5: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					6: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					7: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{ // Step 6 - Blue
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					1: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					2: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					3: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					4: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					5: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
					6: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					7: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{ // Step 7 - Purple
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					1: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					2: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					3: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					4: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					5: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					6: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 100, G: 0, B: 255}}},
					7: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{ // Step 8 - Pink
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					1: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					2: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					3: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					4: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					5: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					6: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					7: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 0, B: 255}}},
				},
			},
		},
	}

	multi := common.Pattern{
		Name:   "Multi Color",
		Label:  "Multi.Color",
		Number: 6,
		Steps: []common.Step{
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
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					1: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 111, B: 0}}},
					2: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 255, B: 0}}},
					3: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					4: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 255, B: 255}}},
					5: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
					6: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 100, G: 0, B: 255}}},
					7: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 0, B: 255}}},
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 111, B: 0}}},
					1: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 255, B: 0}}},
					2: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					3: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 255, B: 255}}},
					4: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
					5: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 100, G: 0, B: 255}}},
					6: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 0, B: 255}}},
					7: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 255, B: 0}}},
					1: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					2: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 255, B: 255}}},
					3: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
					4: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 100, G: 0, B: 255}}},
					5: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 0, B: 255}}},
					6: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					7: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 111, B: 0}}},
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					1: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 255, B: 255}}},
					2: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
					3: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 100, G: 0, B: 255}}},
					4: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 0, B: 255}}},
					5: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					6: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 111, B: 0}}},
					7: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 255, B: 0}}},
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 255, B: 255}}},
					1: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
					2: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 100, G: 0, B: 255}}},
					3: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 0, B: 255}}},
					4: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					5: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 111, B: 0}}},
					6: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 255, B: 0}}},
					7: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
					1: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 100, G: 0, B: 255}}},
					2: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 0, B: 255}}},
					3: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					4: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 111, B: 0}}},
					5: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 255, B: 0}}},
					6: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					7: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 255, B: 255}}},
				},
			},
			{
				Fixtures: map[int]common.Fixture{

					0: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 100, G: 0, B: 255}}},
					1: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 0, B: 255}}},
					2: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					3: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 111, B: 0}}},
					4: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 255, B: 0}}},
					5: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					6: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 255, B: 255}}},
					7: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
				},
			},
			{
				Fixtures: map[int]common.Fixture{

					0: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 0, B: 255}}},
					1: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					2: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 111, B: 0}}},
					3: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 255, B: 0}}},
					4: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					5: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 255, B: 255}}},
					6: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
					7: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 100, G: 0, B: 255}}},
				},
			},
		},
	}

	vu := common.Pattern{
		Name:   "VU.Meter",
		Label:  "VU.Meter",
		Number: 7,
		Steps: []common.Step{
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					1: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					2: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					3: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					4: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					5: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					6: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					7: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					1: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					2: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					3: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					4: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
					5: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
					6: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
					7: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					1: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					2: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					3: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					4: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					5: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					6: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					7: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					1: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					2: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					3: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					4: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					5: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					6: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					7: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					1: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					2: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					3: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					4: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					5: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					6: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					7: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					1: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					2: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					3: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					4: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					5: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 255, B: 0}}},
					6: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					7: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					1: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					2: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					3: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					4: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					5: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 255, B: 0}}},
					6: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 111, B: 0}}},
					7: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					1: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					2: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					3: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					4: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					5: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 255, B: 0}}},
					6: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 111, B: 0}}},
					7: {MasterDimmer: full, Enabled: true, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
				},
			},
		},
	}

	Patterns[0] = standard
	Patterns[1] = flash
	Patterns[2] = rgbchase
	Patterns[3] = pairs
	Patterns[4] = inward
	Patterns[5] = colors
	Patterns[6] = multi
	Patterns[7] = vu
	return Patterns

}

// Storage for scanner values.
type scanner struct {
	values []int
}

// ApplyFixtureState - Apply the state of the fixtures to the pattern, fixture disabling works by disabling the
// steps that have no enabled fixtures AND also disabling in the fixure package. If we only disable here we don't
// catch steps that have more than one fixture alight in any one step.
// So make sure you also turn off the fixture in the fixture receiver.
func ApplyFixtureState(generatedSteps []common.Step, scannerState map[int]common.ScannerState) common.Pattern {

	var pattern common.Pattern

	pattern.Name = "Chase"
	pattern.Label = "std.Scanner.Chase"
	pattern.Steps = []common.Step{}

	for _, step := range generatedSteps {

		newStep := step
		newStep.Fixtures = make(map[int]common.Fixture)
		hasColors := make(map[int]bool)
		for fixtureNumber, fixture := range step.Fixtures {
			newFixture := common.Fixture{}
			newFixture.Enabled = scannerState[fixtureNumber].Enabled
			newFixture.MasterDimmer = 255
			newFixture.Shutter = fixture.Shutter
			newFixture.Colors = fixture.Colors
			for _, color := range newFixture.Colors {
				if color.R > 0 || color.G > 0 || color.B > 0 {
					//newFixture.Colors = []common.Color{{R: 255, G: 255, B: 255}}
					hasColors[fixtureNumber] = true
				} else {
					hasColors[fixtureNumber] = false
				}
			}
			newStep.Fixtures[fixtureNumber] = newFixture
		}

		if debug {
			fmt.Printf("Fixtures \n")
			for fixture := 0; fixture < len(newStep.Fixtures); fixture++ {
				fmt.Printf("Fixture %d Enabled %t Colors %+v\n", fixture, newStep.Fixtures[fixture].Enabled, newStep.Fixtures[fixture].Colors)
			}
		}

		for fixtureNumber, fixture := range newStep.Fixtures {
			// Don't add steps with no enabled fixtures.
			if hasColors[fixtureNumber] && fixture.Enabled {
				pattern.Steps = append(pattern.Steps, newStep)
				break
			}
		}
	}

	if debug {
		for _, step := range pattern.Steps {
			fmt.Printf("Fixtures \n")
			for fixture := 0; fixture < len(step.Fixtures); fixture++ {
				fmt.Printf("Fixture %d Enabled %t Values %+v\n", fixture, step.Fixtures[fixture].Enabled, step.Fixtures[fixture])
			}
		}
	}

	return pattern

}

// GeneratePattern takes an array of Coordinates and turns them into a pattern
// which is the starting point for all sequence steps.
func GeneratePattern(Coordinates []Coordinate, NumberFixtures int, requestedShift int, chase bool, scannerState map[int]common.ScannerState) common.Pattern {

	NumberCoordinates := len(Coordinates)

	if debug {
		fmt.Printf("Number Fixtures %d\n", NumberFixtures)
		fmt.Printf("Number Coordinates %d\n", NumberCoordinates)
	}

	// Storage space for the fixtures
	scanners := []scanner{}

	// First generate the values for all posible fixtures ie assune 8 scanner, split into two pieces.
	// This because we can only shift by four quaters of the scan.
	// First four,
	for fixture := 0; fixture < 4; fixture++ {

		// new scanner
		s := scanner{}

		actualShift := (NumberCoordinates / 4) * requestedShift

		shift := fixture * actualShift

		if shift == NumberCoordinates {
			shift = 0
		}

		if shift == NumberCoordinates+NumberCoordinates/2 {
			shift = NumberCoordinates / 2
		}

		if shift == (NumberCoordinates*2)+(NumberCoordinates/4) {
			shift = NumberCoordinates / 4
		}
		for Coordinate := shift; Coordinate < NumberCoordinates; Coordinate++ {
			s.values = append(s.values, Coordinate)
		}
		for Coordinate := 0; Coordinate < shift; Coordinate++ {
			s.values = append(s.values, Coordinate)
		}

		// append the scanner to the list of scanners.
		scanners = append(scanners, s)
	}

	// Second four.
	// But the same four quaters as the first pass.
	for fixture := 0; fixture < 4; fixture++ {

		// new scanner
		s := scanner{}

		actualShift := (NumberCoordinates / 4) * requestedShift

		shift := fixture * actualShift

		if shift == NumberCoordinates {
			shift = 0
		}

		if shift == NumberCoordinates+NumberCoordinates/2 {
			shift = NumberCoordinates / 2
		}

		if shift == (NumberCoordinates*2)+(NumberCoordinates/4) {
			shift = NumberCoordinates / 4
		}
		for Coordinate := shift; Coordinate < NumberCoordinates; Coordinate++ {
			s.values = append(s.values, Coordinate)
		}
		for Coordinate := 0; Coordinate < shift; Coordinate++ {
			s.values = append(s.values, Coordinate)
		}

		// append the scanner to the list of scanners.
		scanners = append(scanners, s)
	}

	if debug {
		for scannerNumber, scanner := range scanners {
			fmt.Printf("%dscanner %+v\n", scannerNumber, scanner)
		}
	}

	// Last phase is to create the actaual steps for the pattern.

	// First create the actual pattern.
	pattern := common.Pattern{}
	// And then the array for the steps.
	steps := []common.Step{}

	// Now create the steps in the pattern.
	// Now using the actual number of scanners.
	for stepNumber := 0; stepNumber < NumberCoordinates; stepNumber++ {

		// Make space for the fixtures.
		fixtures := make(map[int]common.Fixture)

		// Now add the fixtures.
		for fixture := 0; fixture < NumberFixtures; fixture++ {

			newFixture := common.Fixture{
				MasterDimmer: full,
				Enabled:      true,
				Brightness:   full,
				Shutter:      full,
				// Apply a color to represent each position in the pattern.
				Colors: []common.Color{
					common.GetColorButtonsArray(scanners[fixture].values[stepNumber]),
				},
				Pan:          Coordinates[scanners[fixture].values[stepNumber]].Pan,
				Tilt:         Coordinates[scanners[fixture].values[stepNumber]].Tilt,
				ScannerColor: common.Color{R: 255, G: 255, B: 255}, // White
				Gobo:         0,                                    // First gobo is usually open,
				// TODO find correct gobo and shutter values from the config.
			}
			fixtures[fixture] = newFixture
		}

		newStep := common.Step{
			Fixtures: fixtures,
		}
		steps = append(steps, newStep)
		pattern.Steps = steps
	}

	return pattern
}

type Coordinate struct {
	Tilt int
	Pan  int
}

func GetNumberEnabledScanners(scannerState map[int]common.ScannerState, numberOfFixtures int) int {

	var getNumberEnabledScanners int
	for fixture := 0; fixture < numberOfFixtures; fixture++ {
		if scannerState[fixture].Enabled {
			getNumberEnabledScanners++
		}
	}
	if debug {
		fmt.Printf("getNumberEnabledScanners %d\n", getNumberEnabledScanners)
	}
	return getNumberEnabledScanners
}

func makeEnabledScannerList(scannerState map[int]common.ScannerState, NumberCoordinates int, numberEnabledScanners, numberScanners int) []int {

	enabledScannerList := []int{}

	size := findStepSize(NumberCoordinates, numberEnabledScanners)

	for fixture := 0; fixture < numberScanners; fixture++ {
		if scannerState[fixture].Enabled {
			for s := 0; s < size; s++ {
				enabledScannerList = append(enabledScannerList, fixture)
			}
		}
	}

	if debug {
		fmt.Printf("makeEnabledScannerList %d\n", enabledScannerList)
	}
	return enabledScannerList
}

func findStepSize(NumberCoordinates int, numberEnabledScanners int) int {

	actualNumberCoodinates := float64(NumberCoordinates)
	acutalnumberEnabledScanners := float64(numberEnabledScanners)

	return int(math.Round(actualNumberCoodinates / acutalnumberEnabledScanners))
}

func CircleGenerator(radius int, NumberCoordinates int, posX float64, posY float64) (out []Coordinate) {
	var theta float64
	for theta = 0; theta < 360; theta += (360 / float64(NumberCoordinates)) {
		n := Coordinate{}
		n.Tilt, n.Pan = circleXY(float64(radius), theta, posX, posY)
		out = append(out, n)
	}
	if debug {
		for _, cood := range out {
			fmt.Printf("%d,%d\n", cood.Pan, cood.Tilt)
		}
	}

	return out
}

func ScanGenerateSineWave(size float64, frequency float64, NumberCoordinates float64) (out []Coordinate) {
	var t float64
	size = size * 2
	T := float64(size)
	for t = 1; t < T-1; t += float64((255 / NumberCoordinates)) {
		n := Coordinate{}
		x := (float64(size)/2 + math.Sin(t*float64(frequency))*100)
		n.Tilt = int(x)
		n.Pan = int(t)
		out = append(out, n)
	}
	return out
}

func ScanGeneratorUpDown(size float64, NumberCoordinates float64) (out []Coordinate) {
	var tilt float64
	var divideBy float64
	pan := 128
	size = size * 2
	if size > 255 {
		size = 255
	}
	divideBy = 255 / NumberCoordinates

	for tilt = 0; tilt < size; tilt += divideBy {
		n := Coordinate{}
		n.Tilt = int(tilt)
		n.Pan = int(pan)
		out = append(out, n)
	}
	return out
}

func ScanGeneratorLeftRight(size float64, NumberCoordinates float64) (out []Coordinate) {
	var tilt float64
	var pan float64
	tilt = 128
	size = size * 2
	if size > 255 {
		size = 255
	}
	for pan = 0; pan < size; pan += (255 / NumberCoordinates) {
		n := Coordinate{}
		n.Tilt = int(tilt)
		n.Pan = int(pan)
		out = append(out, n)
	}
	return out
}

func circleXY(radius float64, theta float64, posX float64, posY float64) (int, int) {
	// Convert angle to radians
	theta = (theta - 90) * math.Pi / 180
	// Adding the raduis always positions the circle so no we don't get any negitive numbers.
	x := int(radius*math.Cos(theta) + posX)
	y := int(-radius*math.Sin(theta) + posY)
	return x, y
}
