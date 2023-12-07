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

import "testing"

func Test_autoSelect(t *testing.T) {

	type args struct {
		this *CurrentState
	}
	tests := []struct {
		name                 string
		args                 args
		wantSelectedSequence int
	}{
		{
			name: "first sequence is running",
			args: args{
				this: &CurrentState{
					Running: map[int]bool{
						0: true,
						1: false,
						2: false,
					},
				},
			},
			wantSelectedSequence: 0,
		},
		{
			name: "second sequence is running",
			args: args{
				this: &CurrentState{
					Running: map[int]bool{
						0: false,
						1: true,
						2: false,
					},
				},
			},
			wantSelectedSequence: 1,
		},
		{
			name: "last sequence is running",
			args: args{
				this: &CurrentState{
					Running: map[int]bool{
						0: false,
						1: false,
						2: true,
					},
				},
			},
			wantSelectedSequence: 2,
		},
		{
			name: "last sequence is in static mode",
			args: args{
				this: &CurrentState{
					Running: map[int]bool{
						0: false,
						1: false,
						2: false,
					},
					Static: []bool{
						0: false,
						1: false,
						2: true,
					},
				},
			},
			wantSelectedSequence: 2,
		},
		{
			name: "shutter chaser is running",
			args: args{

				this: &CurrentState{
					ChaserSequenceNumber:  4,
					ScannerSequenceNumber: 2,
					Running: map[int]bool{
						0: false,
						1: false,
						2: false,
						3: false,
						4: true,
					},
				},
			},
			wantSelectedSequence: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSelectedSequence := autoSelect(tt.args.this); gotSelectedSequence != tt.wantSelectedSequence {
				t.Errorf("autoSelect() = %v, want %v", gotSelectedSequence, tt.wantSelectedSequence)
			}
		})
	}
}
