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
	"image/color"
	"math"

	"github.com/dhowlett99/dmxlights/pkg/colors"
	"github.com/dhowlett99/dmxlights/pkg/common"
)

const debug = false

const (
	full = 255
)

func MakeSingleFixtureChase(colors []color.RGBA) common.Pattern {

	steps := []common.Step{}
	for _, color := range colors {
		fixture := common.Fixture{
			MasterDimmer: full,
			Enabled:      true,
			Color:        color,
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
					0: {MasterDimmer: full, Enabled: true, Shutter: 255, Brightness: 255, Color: colors.Green},
					1: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					2: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					3: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					4: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					5: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					6: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					7: {MasterDimmer: full, Enabled: true, Color: colors.Black},
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					1: {MasterDimmer: full, Enabled: true, Shutter: 255, Brightness: 255, Color: colors.Green},
					2: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					3: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					4: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					5: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					6: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					7: {MasterDimmer: full, Enabled: true, Color: colors.Black},
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					1: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					2: {MasterDimmer: full, Enabled: true, Shutter: 255, Brightness: 255, Color: colors.Green},
					3: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					4: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					5: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					6: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					7: {MasterDimmer: full, Enabled: true, Color: colors.Black},
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					1: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					2: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					3: {MasterDimmer: full, Enabled: true, Shutter: 255, Brightness: 255, Color: colors.Green},
					4: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					5: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					6: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					7: {MasterDimmer: full, Enabled: true, Color: colors.Black},
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					1: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					2: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					3: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					4: {MasterDimmer: full, Enabled: true, Shutter: 255, Brightness: 255, Color: colors.Green},
					5: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					6: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					7: {MasterDimmer: full, Enabled: true, Color: colors.Black},
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					1: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					2: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					3: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					4: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					5: {MasterDimmer: full, Enabled: true, Shutter: 255, Brightness: 255, Color: colors.Green},
					6: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					7: {MasterDimmer: full, Enabled: true, Color: colors.Black},
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					1: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					2: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					3: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					4: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					5: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					6: {MasterDimmer: full, Enabled: true, Shutter: 255, Brightness: 255, Color: colors.Green},
					7: {MasterDimmer: full, Enabled: true, Color: colors.Black},
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					1: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					2: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					3: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					4: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					5: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					6: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					7: {MasterDimmer: full, Enabled: true, Shutter: 255, Brightness: 255, Color: colors.Green},
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
				KeyStep: true,
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Color: colors.White},
					1: {MasterDimmer: full, Enabled: true, Color: colors.White},
					2: {MasterDimmer: full, Enabled: true, Color: colors.White},
					3: {MasterDimmer: full, Enabled: true, Color: colors.White},
					4: {MasterDimmer: full, Enabled: true, Color: colors.White},
					5: {MasterDimmer: full, Enabled: true, Color: colors.White},
					6: {MasterDimmer: full, Enabled: true, Color: colors.White},
					7: {MasterDimmer: full, Enabled: true, Color: colors.White},
				},
			},
			{
				KeyStep: true,
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					1: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					2: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					3: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					4: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					5: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					6: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					7: {MasterDimmer: full, Enabled: true, Color: colors.Black},
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
					0: {MasterDimmer: full, Enabled: true, Color: colors.Red},
					1: {MasterDimmer: full, Enabled: true, Color: colors.Red},
					2: {MasterDimmer: full, Enabled: true, Color: colors.Red},
					3: {MasterDimmer: full, Enabled: true, Color: colors.Red},
					4: {MasterDimmer: full, Enabled: true, Color: colors.Red},
					5: {MasterDimmer: full, Enabled: true, Color: colors.Red},
					6: {MasterDimmer: full, Enabled: true, Color: colors.Red},
					7: {MasterDimmer: full, Enabled: true, Color: colors.Red},
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Color: colors.Green},
					1: {MasterDimmer: full, Enabled: true, Color: colors.Green},
					2: {MasterDimmer: full, Enabled: true, Color: colors.Green},
					3: {MasterDimmer: full, Enabled: true, Color: colors.Green},
					4: {MasterDimmer: full, Enabled: true, Color: colors.Green},
					5: {MasterDimmer: full, Enabled: true, Color: colors.Green},
					6: {MasterDimmer: full, Enabled: true, Color: colors.Green},
					7: {MasterDimmer: full, Enabled: true, Color: colors.Green},
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Color: colors.Blue},
					1: {MasterDimmer: full, Enabled: true, Color: colors.Blue},
					2: {MasterDimmer: full, Enabled: true, Color: colors.Blue},
					3: {MasterDimmer: full, Enabled: true, Color: colors.Blue},
					4: {MasterDimmer: full, Enabled: true, Color: colors.Blue},
					5: {MasterDimmer: full, Enabled: true, Color: colors.Blue},
					6: {MasterDimmer: full, Enabled: true, Color: colors.Blue},
					7: {MasterDimmer: full, Enabled: true, Color: colors.Blue},
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
					0: {MasterDimmer: full, Enabled: true, Color: colors.Blue},
					1: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					2: {MasterDimmer: full, Enabled: true, Color: colors.Blue},
					3: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					4: {MasterDimmer: full, Enabled: true, Color: colors.Blue},
					5: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					6: {MasterDimmer: full, Enabled: true, Color: colors.Blue},
					7: {MasterDimmer: full, Enabled: true, Color: colors.Black},
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					1: {MasterDimmer: full, Enabled: true, Color: colors.Blue},
					2: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					3: {MasterDimmer: full, Enabled: true, Color: colors.Blue},
					4: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					5: {MasterDimmer: full, Enabled: true, Color: colors.Blue},
					6: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					7: {MasterDimmer: full, Enabled: true, Color: colors.Blue},
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
					0: {MasterDimmer: full, Enabled: true, Color: colors.Red},
					1: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					2: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					3: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					4: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					5: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					6: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					7: {MasterDimmer: full, Enabled: true, Color: colors.Red},
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					1: {MasterDimmer: full, Enabled: true, Color: colors.Red},
					2: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					3: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					4: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					5: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					6: {MasterDimmer: full, Enabled: true, Color: colors.Red},
					7: {MasterDimmer: full, Enabled: true, Color: colors.Black},
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					1: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					2: {MasterDimmer: full, Enabled: true, Color: colors.Red},
					3: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					4: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					5: {MasterDimmer: full, Enabled: true, Color: colors.Red},
					6: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					7: {MasterDimmer: full, Enabled: true, Color: colors.Black},
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					1: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					2: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					3: {MasterDimmer: full, Enabled: true, Color: colors.Red},
					4: {MasterDimmer: full, Enabled: true, Color: colors.Red},
					5: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					6: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					7: {MasterDimmer: full, Enabled: true, Color: colors.Black},
				},
			},
		},
	}

	color := common.Pattern{
		Name:   "Color Chase",
		Label:  "Color.Chase",
		Number: 5,
		Steps: []common.Step{
			{ // Step 1, - Red
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Color: colors.Red},
					1: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					2: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					3: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					4: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					5: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					6: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					7: {MasterDimmer: full, Enabled: true, Color: colors.Black},
				},
			},
			{ // Step 2 - Orange
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					1: {MasterDimmer: full, Enabled: true, Color: colors.Orange},
					2: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					3: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					4: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					5: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					6: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					7: {MasterDimmer: full, Enabled: true, Color: colors.Black},
				},
			},
			{ // Step 3 - Yellow
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					1: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					2: {MasterDimmer: full, Enabled: true, Color: colors.Yellow},
					3: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					4: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					5: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					6: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					7: {MasterDimmer: full, Enabled: true, Color: colors.Black},
				},
			},
			{ // Step 4 - Green
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					1: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					2: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					3: {MasterDimmer: full, Enabled: true, Color: colors.Green},
					4: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					5: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					6: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					7: {MasterDimmer: full, Enabled: true, Color: colors.Black},
				},
			},
			{ // Step 5 - Cyan
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					1: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					2: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					3: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					4: {MasterDimmer: full, Enabled: true, Color: colors.Cyan},
					5: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					6: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					7: {MasterDimmer: full, Enabled: true, Color: colors.Black},
				},
			},
			{ // Step 6 - Blue
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					1: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					2: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					3: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					4: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					5: {MasterDimmer: full, Enabled: true, Color: colors.Blue},
					6: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					7: {MasterDimmer: full, Enabled: true, Color: colors.Black},
				},
			},
			{ // Step 7 - Purple
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					1: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					2: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					3: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					4: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					5: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					6: {MasterDimmer: full, Enabled: true, Color: colors.Purple},
					7: {MasterDimmer: full, Enabled: true, Color: colors.Black},
				},
			},
			{ // Step 8 - Magenta
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					1: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					2: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					3: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					4: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					5: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					6: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					7: {MasterDimmer: full, Enabled: true, Color: colors.Magenta},
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
					0: {MasterDimmer: full, Enabled: true, Color: colors.Red},     // Red
					1: {MasterDimmer: full, Enabled: true, Color: colors.Orange},  // Orange
					2: {MasterDimmer: full, Enabled: true, Color: colors.Yellow},  // Yellow
					3: {MasterDimmer: full, Enabled: true, Color: colors.Green},   // Green
					4: {MasterDimmer: full, Enabled: true, Color: colors.Cyan},    // Cyan
					5: {MasterDimmer: full, Enabled: true, Color: colors.Blue},    // Blue
					6: {MasterDimmer: full, Enabled: true, Color: colors.Purple},  // Purple
					7: {MasterDimmer: full, Enabled: true, Color: colors.Magenta}, // Magenta
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Color: colors.Orange},  // Orange
					1: {MasterDimmer: full, Enabled: true, Color: colors.Yellow},  // Yellow
					2: {MasterDimmer: full, Enabled: true, Color: colors.Green},   // Green
					3: {MasterDimmer: full, Enabled: true, Color: colors.Cyan},    // Cyan
					4: {MasterDimmer: full, Enabled: true, Color: colors.Blue},    // Blue
					5: {MasterDimmer: full, Enabled: true, Color: colors.Purple},  // Purple
					6: {MasterDimmer: full, Enabled: true, Color: colors.Magenta}, // Magenta
					7: {MasterDimmer: full, Enabled: true, Color: colors.Red},     // Red
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Color: colors.Yellow},  // Yellow
					1: {MasterDimmer: full, Enabled: true, Color: colors.Green},   // Green
					2: {MasterDimmer: full, Enabled: true, Color: colors.Cyan},    // Cyan
					3: {MasterDimmer: full, Enabled: true, Color: colors.Blue},    // Blue
					4: {MasterDimmer: full, Enabled: true, Color: colors.Purple},  // Purple
					5: {MasterDimmer: full, Enabled: true, Color: colors.Magenta}, // Magenta
					6: {MasterDimmer: full, Enabled: true, Color: colors.Red},     // Red
					7: {MasterDimmer: full, Enabled: true, Color: colors.Orange},  // Orange
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Color: colors.Green},   // Green
					1: {MasterDimmer: full, Enabled: true, Color: colors.Cyan},    // Cyan
					2: {MasterDimmer: full, Enabled: true, Color: colors.Blue},    // Blue
					3: {MasterDimmer: full, Enabled: true, Color: colors.Purple},  // Purple
					4: {MasterDimmer: full, Enabled: true, Color: colors.Magenta}, // Magenta
					5: {MasterDimmer: full, Enabled: true, Color: colors.Red},     // Red
					6: {MasterDimmer: full, Enabled: true, Color: colors.Orange},  // Orange
					7: {MasterDimmer: full, Enabled: true, Color: colors.Yellow},  // Yellow
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Color: colors.Cyan},    // Cyan
					1: {MasterDimmer: full, Enabled: true, Color: colors.Blue},    // Blue
					2: {MasterDimmer: full, Enabled: true, Color: colors.Purple},  // Purple
					3: {MasterDimmer: full, Enabled: true, Color: colors.Magenta}, // Magenta
					4: {MasterDimmer: full, Enabled: true, Color: colors.Red},     // Red
					5: {MasterDimmer: full, Enabled: true, Color: colors.Orange},  // Orange
					6: {MasterDimmer: full, Enabled: true, Color: colors.Yellow},  // Yellow
					7: {MasterDimmer: full, Enabled: true, Color: colors.Green},   // Green
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Color: colors.Blue},    // Blue
					1: {MasterDimmer: full, Enabled: true, Color: colors.Purple},  // Purple
					2: {MasterDimmer: full, Enabled: true, Color: colors.Magenta}, // Magenta
					3: {MasterDimmer: full, Enabled: true, Color: colors.Red},     // Red
					4: {MasterDimmer: full, Enabled: true, Color: colors.Orange},  // Orange
					5: {MasterDimmer: full, Enabled: true, Color: colors.Yellow},  // Yellow
					6: {MasterDimmer: full, Enabled: true, Color: colors.Green},   // Green
					7: {MasterDimmer: full, Enabled: true, Color: colors.Cyan},    // Cyan
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Color: colors.Purple},  // Purple
					1: {MasterDimmer: full, Enabled: true, Color: colors.Magenta}, // Magenta
					2: {MasterDimmer: full, Enabled: true, Color: colors.Red},     // Red
					3: {MasterDimmer: full, Enabled: true, Color: colors.Orange},  // Orange
					4: {MasterDimmer: full, Enabled: true, Color: colors.Yellow},  // Yellow
					5: {MasterDimmer: full, Enabled: true, Color: colors.Green},   // Green
					6: {MasterDimmer: full, Enabled: true, Color: colors.Cyan},    // Cyan
					7: {MasterDimmer: full, Enabled: true, Color: colors.Blue},    // Blue
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Color: colors.Magenta}, // Magenta
					1: {MasterDimmer: full, Enabled: true, Color: colors.Red},     // Red
					2: {MasterDimmer: full, Enabled: true, Color: colors.Orange},  // Orange
					3: {MasterDimmer: full, Enabled: true, Color: colors.Yellow},  // Yellow
					4: {MasterDimmer: full, Enabled: true, Color: colors.Green},   // Green
					5: {MasterDimmer: full, Enabled: true, Color: colors.Cyan},    // Cyan
					6: {MasterDimmer: full, Enabled: true, Color: colors.Blue},    // Blue
					7: {MasterDimmer: full, Enabled: true, Color: colors.Purple},  // Purple
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
					0: {MasterDimmer: full, Enabled: true, Color: colors.Green},
					1: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					2: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					3: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					4: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					5: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					6: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					7: {MasterDimmer: full, Enabled: true, Color: colors.Black},
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Color: colors.Green},
					1: {MasterDimmer: full, Enabled: true, Color: colors.Green},
					2: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					3: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					4: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					5: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					6: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					7: {MasterDimmer: full, Enabled: true, Color: colors.Black},
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Color: colors.Green},
					1: {MasterDimmer: full, Enabled: true, Color: colors.Green},
					2: {MasterDimmer: full, Enabled: true, Color: colors.Green},
					3: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					4: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					5: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					6: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					7: {MasterDimmer: full, Enabled: true, Color: colors.Black},
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Color: colors.Green},
					1: {MasterDimmer: full, Enabled: true, Color: colors.Green},
					2: {MasterDimmer: full, Enabled: true, Color: colors.Green},
					3: {MasterDimmer: full, Enabled: true, Color: colors.Green},
					4: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					5: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					6: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					7: {MasterDimmer: full, Enabled: true, Color: colors.Black},
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Color: colors.Green},
					1: {MasterDimmer: full, Enabled: true, Color: colors.Green},
					2: {MasterDimmer: full, Enabled: true, Color: colors.Green},
					3: {MasterDimmer: full, Enabled: true, Color: colors.Green},
					4: {MasterDimmer: full, Enabled: true, Color: colors.Green},
					5: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					6: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					7: {MasterDimmer: full, Enabled: true, Color: colors.Black},
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Color: colors.Green},
					1: {MasterDimmer: full, Enabled: true, Color: colors.Green},
					2: {MasterDimmer: full, Enabled: true, Color: colors.Green},
					3: {MasterDimmer: full, Enabled: true, Color: colors.Green},
					4: {MasterDimmer: full, Enabled: true, Color: colors.Green},
					5: {MasterDimmer: full, Enabled: true, Color: colors.Yellow},
					6: {MasterDimmer: full, Enabled: true, Color: colors.Black},
					7: {MasterDimmer: full, Enabled: true, Color: colors.Black},
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Color: colors.Green},
					1: {MasterDimmer: full, Enabled: true, Color: colors.Green},
					2: {MasterDimmer: full, Enabled: true, Color: colors.Green},
					3: {MasterDimmer: full, Enabled: true, Color: colors.Green},
					4: {MasterDimmer: full, Enabled: true, Color: colors.Green},
					5: {MasterDimmer: full, Enabled: true, Color: colors.Yellow},
					6: {MasterDimmer: full, Enabled: true, Color: colors.Orange},
					7: {MasterDimmer: full, Enabled: true, Color: colors.Black},
				},
			},
			{
				Fixtures: map[int]common.Fixture{
					0: {MasterDimmer: full, Enabled: true, Color: colors.Green},
					1: {MasterDimmer: full, Enabled: true, Color: colors.Green},
					2: {MasterDimmer: full, Enabled: true, Color: colors.Green},
					3: {MasterDimmer: full, Enabled: true, Color: colors.Green},
					4: {MasterDimmer: full, Enabled: true, Color: colors.Green},
					5: {MasterDimmer: full, Enabled: true, Color: colors.Yellow},
					6: {MasterDimmer: full, Enabled: true, Color: colors.Orange},
					7: {MasterDimmer: full, Enabled: true, Color: colors.Red},
				},
			},
		},
	}

	Patterns[0] = standard
	Patterns[1] = flash
	Patterns[2] = rgbchase
	Patterns[3] = pairs
	Patterns[4] = inward
	Patterns[5] = color
	Patterns[6] = multi
	Patterns[7] = vu
	return Patterns

}

// Storage for scanner values.
type scanner struct {
	values []int
}

// GeneratePattern takes an array of Coordinates and turns them into a pattern
// which is the starting point for all sequence steps.
func GeneratePattern(Coordinates []Coordinate, NumberFixtures int, requestedShift int, chase bool, scannerState map[int]common.FixtureState) common.Pattern {

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
				Color:        common.GetColorButtonsArray(scanners[fixture].values[stepNumber]),
				Pan:          Coordinates[scanners[fixture].values[stepNumber]].Pan,
				Tilt:         Coordinates[scanners[fixture].values[stepNumber]].Tilt,
				ScannerColor: colors.White, // White
				Gobo:         0,            // First gobo is usually open,
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

func GetNumberEnabledScanners(scannerState map[int]common.FixtureState, numberOfFixtures int) int {

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

func makeEnabledScannerList(scannerState map[int]common.FixtureState, NumberCoordinates int, numberEnabledScanners, numberScanners int) []int {

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

// posY runs from 0 to 255 and starts in the centre at 127.
// The goal here is to return the start and stop values for the scanner pattern generators
// so that we can pan the pattern from left to right.
func findStart(posY int, maxDMX int) (start float64, stop float64) {

	if posY == common.CENTER_DMX_BRIGHTNESS {
		start = 0
		stop = common.MAX_DMX_BRIGHTNESS
		return start, stop
	}
	if posY < common.CENTER_DMX_BRIGHTNESS {
		in := []int{posY}
		out := scaleBetween(in, 1, maxDMX, 0, maxDMX)
		start = common.MIN_DMX_BRIGHTNESS
		stop = float64(out[0]) * 2
	}
	if posY > common.CENTER_DMX_BRIGHTNESS {
		in := []int{posY / 2}
		out := scaleBetween(in, 0, maxDMX-1, 0, maxDMX)
		start = float64(out[0]) * 2
		stop = common.MAX_DMX_BRIGHTNESS
	}
	return start, stop
}

func scaleBetween(unscaledNum []int, minAllowed int, maxAllowed int, min int, max int) (out []int) {
	for _, number := range unscaledNum {
		arg := (maxAllowed-minAllowed)*(number-min)/max - min + minAllowed
		out = append(out, arg)
	}
	return out
}

func ScanGenerateSawTooth(size float64, frequency float64, numberCoordinates float64, posX float64, posY float64) (out []Coordinate) {

	var y float64
	var x float64

	size = size * 2

	lift := (common.MAX_DMX_BRIGHTNESS - size) / 2

	start, stop := findStart(int(posY), 127)

	for y = start; y < stop; y += float64(255 / numberCoordinates) {
		n := Coordinate{}
		x = traingle(y, size, frequency)
		n.Tilt = int(x) + int(lift) - 127 + int(posX)
		n.Pan = int(y)
		out = append(out, n)
	}
	return out
}

// traingle creates a symmetrical triangle
func traingle(y float64, size float64, freq float64) float64 {
	arg := math.Round(y/freq) - (y / freq)
	x := size * 2 * math.Abs(arg)
	return x
}

func ScanGeneratorUpDown(size float64, NumberCoordinates float64, posX float64, posY float64) (out []Coordinate) {
	var tilt float64
	var divideBy float64
	pan := posY
	size = size * 2

	lift := (common.MAX_DMX_BRIGHTNESS - size) / 2

	if size > 255 {
		size = 255
	}
	divideBy = 255 / NumberCoordinates

	for tilt = 0; tilt < size; tilt += divideBy {
		n := Coordinate{}
		n.Tilt = int(tilt) + int(lift) - 127 + int(posX)
		n.Pan = int(pan)
		out = append(out, n)
	}
	return out
}

func ScanGeneratorLeftRight(size float64, NumberCoordinates float64, posX float64, posY float64) (out []Coordinate) {
	var tilt float64
	var pan float64
	tilt = posX
	size = (size * 2)

	lift := (common.MAX_DMX_BRIGHTNESS - size) / 2

	if size > 255 {
		size = 255
	}
	for pan = 0; pan < size; pan += (255 / NumberCoordinates) {
		n := Coordinate{}
		n.Tilt = int(tilt)
		n.Pan = int(pan) + int(lift) - 127 + int(posY)
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
