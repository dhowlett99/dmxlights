// Copyright (C) 2022,2023,2024,2025 dhowlett99.
// This is the dmxlights fixture control code.
// this function is used to set up the RGB patterns using the fixture config.
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

package fixture

import (
	"fmt"

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/pattern"
)

// LoadAvailablePatterns configures a set of patterns for an RGB sequence.
// Takes a sequence and fixtures config.
// Returns an array of available patterns for this RGB sequence.
func LoadAvailablePatterns(sequence common.Sequence, fixturesConfig *Fixtures) []common.Pattern {

	RGBAvailablePatterns := pattern.MakePatterns(sequence.NumberFixtures)
	if debug {
		fmt.Printf("%d: Number of Patterms %d\n", sequence.Number, len(RGBAvailablePatterns))
	}
	return RGBAvailablePatterns
}
