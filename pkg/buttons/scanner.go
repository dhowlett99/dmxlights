// Copyright (C) 2022, 2023, 2024, 2025 dhowlett99.
// This is button processor, used by the launchpad and gui interfaces.
// This file holds scanner functions.
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

func getScannerShiftLabel(shift int) string {

	switch {
	case shift == 0:
		return "Sync"

	case shift == 1:
		return "1/4"

	case shift == 2:
		return "1/2"

	case shift == 3:
		return "3/4"

	}
	return ""
}

func getScannerCoordinatesLabel(shift int) string {

	switch {
	case shift == 0:
		return "12"

	case shift == 1:
		return "16"

	case shift == 2:
		return "24"

	case shift == 3:
		return "32"

	case shift == 4:
		return "64"

	}
	return ""
}
