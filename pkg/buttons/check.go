// Copyright (C) 2022, 2023, 2024, 2025 dhowlett99.
// This is check type helper function.
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
//

package buttons

// Check If we're a scanner and we're in shutter chase mode return the chaser sequence number.
// otherwise return this seleeted sequnce number.
func CheckType(sequenceType string, this *CurrentState) int {

	if sequenceType == "scanner" && this.ScannerChaser[this.SelectedSequence] &&
		(this.SelectedMode[this.SelectedSequence] == CHASER_FUNCTION || this.SelectedMode[this.SelectedSequence] == CHASER_DISPLAY) {
		return this.ChaserSequenceNumber
	} else {
		return this.SelectedSequence
	}

}
