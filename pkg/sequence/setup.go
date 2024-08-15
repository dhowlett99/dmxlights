// Copyright (C) 2022,2023 dhowlett99.
// This is the dmxlights main sequencer responsible for controlling all
// of the fixtures in a group.
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
	"fmt"

	"github.com/dhowlett99/dmxlights/pkg/common"
)

func setupChase(mySequenceNumber int, sequence *common.Sequence, availablePatterns map[int]common.Pattern, RGBPositions map[int]common.Position) {

	if debug {
		fmt.Printf("%d: setupChase UpdatePattern %t Type %s\n", mySequenceNumber, sequence.UpdatePattern, sequence.Type)
	}

	// Setup rgb patterns.
	if sequence.UpdatePattern && sequence.Type == "rgb" {
		if debug {
			fmt.Printf("%d: Setup RGB Patterns\n", mySequenceNumber)
		}
		sequence.Pattern.Steps = setupRGBPatterns(sequence, availablePatterns)
		fmt.Printf("%d:\t\t Sequence Colors are %+v\n", mySequenceNumber, sequence.SequenceColors)
		sequence.UpdatePattern = false
	}

	// Setup scanner patterns.
	if (sequence.UpdatePattern || sequence.UpdateShift) && sequence.Type == "scanner" {
		if debug {
			fmt.Printf("%d: Setup Scanner Patterns 2\n", mySequenceNumber)
		}
		sequence.Pattern.Steps = setupScannerPatterns(sequence)
		sequence.UpdatePattern = false
		sequence.UpdateShift = false
	}

	// Setup colors in a sequence.
	if sequence.UpdateColors && sequence.Type == "rgb" {
		if debug {
			fmt.Printf("%d: Update RGB Sequence Colors to %+v\n", mySequenceNumber, sequence.SequenceColors)
		}
		sequence.Pattern.Steps = setupColors(sequence, RGBPositions)
		sequence.UpdateColors = false
	}

}
