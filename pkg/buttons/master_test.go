// Copyright (C) 2022, 2023 , 2024 dhowlett99.
// This is button processor, used by the launchpad and gui interfaces.
// This file processes master brightness buttons and controls their actions.
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

func Test_getSelectedSequenceNumber(t *testing.T) {
	type args struct {
		selectedSequence int
		selectedType     string
		selectedSwitch   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//  Generate psudo sequence number.
		{
			name: "simple sequence selected",
			args: args{
				selectedSequence: 0,
				selectedType:     "rgb",
			},
			want: 0,
		},
		{
			name: "simple sequence selected",
			args: args{
				selectedSequence: 1,
				selectedType:     "rgb",
			},
			want: 1,
		},
		{
			name: "simple sequence selected",
			args: args{
				selectedSequence: 2,
				selectedType:     "scanner",
			},
			want: 2,
		},
		{
			name: "simple sequence selected",
			args: args{
				selectedSequence: 4,
				selectedType:     "chaser",
			},
			want: 4,
		},
		{
			name: "switch sequence selected",
			args: args{
				selectedSequence: 3,
				selectedType:     "switch",
				selectedSwitch:   0,
			},
			want: 5,
		},
		{
			name: "switch sequence selected",
			args: args{
				selectedSequence: 3,
				selectedType:     "switch",
				selectedSwitch:   1,
			},
			want: 6,
		},
		{
			name: "switch sequence selected",
			args: args{
				selectedSequence: 3,
				selectedType:     "switch",
				selectedSwitch:   2,
			},
			want: 7,
		},
		{
			name: "switch sequence selected",
			args: args{
				selectedSequence: 3,
				selectedType:     "switch",
				selectedSwitch:   3,
			},
			want: 8,
		},
		{
			name: "switch sequence selected",
			args: args{
				selectedSequence: 3,
				selectedType:     "switch",
				selectedSwitch:   4,
			},
			want: 9,
		},
		{
			name: "switch sequence selected",
			args: args{
				selectedSequence: 3,
				selectedType:     "switch",
				selectedSwitch:   5,
			},
			want: 10,
		},
		{
			name: "switch sequence selected",
			args: args{
				selectedSequence: 3,
				selectedType:     "switch",
				selectedSwitch:   6,
			},
			want: 11,
		},
		{
			name: "switch sequence selected",
			args: args{
				selectedSequence: 3,
				selectedType:     "switch",
				selectedSwitch:   7,
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSelectedSequenceNumber(tt.args.selectedSequence, tt.args.selectedType, tt.args.selectedSwitch); got != tt.want {
				t.Errorf("getSelectedSequenceNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
