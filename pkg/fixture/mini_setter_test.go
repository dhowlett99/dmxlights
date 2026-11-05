// Copyright (C) 2022,2023,2024,2025 dhowlett99.
// This is the dmxlights mini settings player, used by the fixture to control
// settings for fixtures.
// Implemented and depends usbdmx.
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

import "testing"

func Test_makeStrobeSpeed(t *testing.T) {
	type args struct {
		strobeValues []int
		strobeSpeed  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "strobe speed 0",
			args: args{
				strobeValues: []int{64, 95},
				strobeSpeed:  0,
			},
			want: 64,
		},
		{
			name: "strobe speed 35",
			args: args{
				strobeValues: []int{64, 95},
				strobeSpeed:  35,
			},
			want: 67,
		},
		{
			name: "strobe speed 70",
			args: args{
				strobeValues: []int{64, 95},
				strobeSpeed:  70,
			},
			want: 71,
		},
		{
			name: "strobe speed 105",
			args: args{
				strobeValues: []int{64, 95},
				strobeSpeed:  105,
			},
			want: 78,
		},
		{
			name: "strobe speed 140",
			args: args{
				strobeValues: []int{64, 95},
				strobeSpeed:  140,
			},
			want: 81,
		},
		{
			name: "strobe speed 175",
			args: args{
				strobeValues: []int{64, 95},
				strobeSpeed:  175,
			},
			want: 88,
		},
		{
			name: "strobe speed 210",
			args: args{
				strobeValues: []int{64, 95},
				strobeSpeed:  210,
			},
			want: 92,
		},
		{
			name: "strobe speed 245",
			args: args{
				strobeValues: []int{64, 95},
				strobeSpeed:  245,
			},
			want: 95,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeStrobeSpeed(tt.args.strobeValues, tt.args.strobeSpeed); got != tt.want {
				t.Errorf("makeStrobeSpeed() = %v, want %v", got, tt.want)
			}
		})
	}
}
