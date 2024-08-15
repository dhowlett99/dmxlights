// Copyright (C) 2022, 2023 dhowlett99.
// This implements the load preset feature, used by the buttons package.
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

package buttons

import "fmt"

// getNextMenuItem get the next items in the menu sequence.
// Wraps if your at the end.
func getNextMenuItem(currentMode int, chaser bool, staticColorMode bool) int {

	if debug {
		fmt.Printf("getNextMenuItem current Mode %s chaser %t static %t\n", printMode(currentMode), chaser, staticColorMode)
	}

	menuOrder := []int{NORMAL, NORMAL_STATIC, FUNCTION, CHASER_DISPLAY, CHASER_DISPLAY_STATIC, CHASER_FUNCTION, STATUS}

	if !chaser && !staticColorMode {
		switch {
		case currentMode == NORMAL:
			return menuOrder[FUNCTION]

		case currentMode == FUNCTION:
			return menuOrder[STATUS]

		case currentMode == STATUS:
			return menuOrder[NORMAL]
		}
	}

	if !chaser && staticColorMode {
		switch {
		case currentMode == NORMAL:
			return menuOrder[NORMAL_STATIC]

		case currentMode == NORMAL_STATIC:
			return menuOrder[FUNCTION]

		case currentMode == FUNCTION:
			return menuOrder[STATUS]

		case currentMode == STATUS:
			return menuOrder[NORMAL]
		}
	}

	if chaser && !staticColorMode {
		switch {
		case currentMode == NORMAL:
			return menuOrder[FUNCTION]

		case currentMode == FUNCTION:
			return menuOrder[CHASER_DISPLAY]

		case currentMode == CHASER_DISPLAY:
			return menuOrder[CHASER_FUNCTION]

		case currentMode == CHASER_FUNCTION:
			return menuOrder[STATUS]

		case currentMode == STATUS:
			return menuOrder[NORMAL]
		}
	}

	if chaser && staticColorMode {
		switch {
		case currentMode == NORMAL:
			return menuOrder[FUNCTION]

		case currentMode == FUNCTION:
			return menuOrder[CHASER_DISPLAY]

		case currentMode == CHASER_DISPLAY:
			return menuOrder[CHASER_DISPLAY_STATIC]

		case currentMode == CHASER_DISPLAY_STATIC:
			return menuOrder[CHASER_FUNCTION]

		case currentMode == CHASER_FUNCTION:
			return menuOrder[STATUS]

		case currentMode == STATUS:
			return menuOrder[NORMAL]
		}
	}

	return 0
}
