// Copyright (C) 2022, 2023, 2024, 2025 dhowlett99.
// This is status bar update code, used to update the speed, shift, size and fade labels.
// Along with the labels on the bottom row buttons.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software \t\tFoundation, either version 3 of the License, or
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
	"reflect"
	"testing"
)

func Test_findSlots(t *testing.T) {
	type args struct {
		channels []string
	}
	tests := []struct {
		name string
		args args
		want []Slot
	}{
		{
			name: "allocate first slot to Program",
			args: args{
				channels: []string{
					"Master", "Program", "Rotate", "Gobo", "Color",
				},
			},
			want: []Slot{
				{
					Button:          0,
					Label:           "speed",
					ChannelAssigned: "Program",
				},
				{
					Button:          2,
					Label:           "shift",
					ChannelAssigned: "Rotate",
				},
				{
					Button:          4,
					Label:           "size",
					ChannelAssigned: "Gobo",
				},
				{
					Button:          6,
					Label:           "fade",
					ChannelAssigned: "Color",
				},
			},
		},
		{
			name: "allocate first slot to ProgramSpeed and second to Program",
			args: args{
				channels: []string{
					"ProgramSpeed", "Program",
				},
			},
			want: []Slot{
				{
					Button:          0,
					Label:           "speed",
					ChannelAssigned: "ProgramSpeed",
				},
				{
					Button:          2,
					Label:           "shift",
					ChannelAssigned: "Program",
				},
			},
		},
		{
			name: "allocate first slot to Color",
			args: args{
				channels: []string{
					"Red", "Color",
				},
			},
			want: []Slot{
				{
					Button:          0,
					Label:           "speed",
					ChannelAssigned: "Color",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSlots(tt.args.channels); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSlots() got = %v", got)
				t.Errorf("findSlots() want %v", tt.want)
			}
		})
	}
}
