// Copyright (C) 2022, 2023 , 2024 dhowlett99.
// This is button processor, used by the launchpad and gui interfaces.
// This file holds the tests for the menu items generator.
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

import (
	"testing"
)

func Test_getNextMenuItem(t *testing.T) {
	type args struct {
		selectedMode    int
		chaser          bool
		editstaticcolor bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "get next item, send normal want function",
			args: args{
				selectedMode:    NORMAL,
				chaser:          false,
				editstaticcolor: false,
			},
			want: FUNCTION,
		},
		{
			name: "get next item, send function want status",
			args: args{
				selectedMode:    FUNCTION,
				chaser:          false,
				editstaticcolor: false,
			},
			want: STATUS,
		},
		{
			name: "get next item, send status want normal",
			args: args{
				selectedMode:    STATUS,
				chaser:          false,
				editstaticcolor: false,
			},
			want: NORMAL,
		},

		// Chaser mode.
		{
			name: "get next item, send normal want function",
			args: args{
				selectedMode:    NORMAL,
				chaser:          true,
				editstaticcolor: false,
			},
			want: FUNCTION,
		},
		{
			name: "get next item, send function chaser display",
			args: args{
				selectedMode:    FUNCTION,
				chaser:          true,
				editstaticcolor: false,
			},
			want: CHASER_DISPLAY,
		},
		{
			name: "get next item, send chaser display want chaser function",
			args: args{
				selectedMode:    CHASER_DISPLAY,
				chaser:          true,
				editstaticcolor: false,
			},
			want: CHASER_FUNCTION,
		},
		{
			name: "get next item, send chaser function want status",
			args: args{
				selectedMode:    CHASER_FUNCTION,
				chaser:          true,
				editstaticcolor: false,
			},
			want: STATUS,
		},
		{
			name: "get next item, send status want normal,",
			args: args{
				selectedMode:    STATUS,
				chaser:          true,
				editstaticcolor: false,
			},
			want: NORMAL,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNextMenuItem(tt.args.selectedMode, tt.args.chaser, tt.args.editstaticcolor); got != tt.want {
				t.Errorf("getNextMenuItem() = %v, want %v", got, tt.want)
			}
		})
	}
}
