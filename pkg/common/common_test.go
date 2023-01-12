// Copyright (C) 2022, 2023 dhowlett99.
// This is the test program for common functions.
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

package common

import (
	"reflect"
	"testing"
)

func Test_invertColor(t *testing.T) {
	type args struct {
		color Color
	}
	tests := []struct {
		name    string
		args    args
		wantOut Color
	}{
		{
			name: "invert white",
			args: args{
				color: Color{
					R: 255,
					G: 255,
					B: 255,
				},
			},
			wantOut: Color{
				R: 0,
				G: 0,
				B: 0,
			},
		},
		{
			name: "invert black",
			args: args{
				color: Color{
					R: 0,
					G: 0,
					B: 0,
				},
			},
			wantOut: Color{
				R: 255,
				G: 255,
				B: 255,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := InvertColor(tt.args.color); !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("InvertColor() = %+v, want %+v", gotOut, tt.wantOut)
			}
		})
	}
}

func Test_getFadeValues(t *testing.T) {
	type args struct {
		size    float64
		fade    int
		reverse bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Fade up 1",
			args: args{
				size:    255,
				fade:    1,
				reverse: false, // Fade up
			},
			want: []int{0, 7, 29, 63, 105, 149, 191, 225, 247, 255},
		},
		{
			name: "Fade down 1",
			args: args{
				size:    255,
				fade:    1,
				reverse: true, // Fade up
			},
			want: []int{255, 247, 225, 191, 149, 105, 63, 29, 7, 0},
		},
		{
			name: "Fade up 2",
			args: args{
				size:    255,
				fade:    2,
				reverse: false, // Fade up
			},
			want: []int{0, 11, 45, 94, 149, 200, 237, 254, 255, 255},
		},
		{
			name: "Fade down 2",
			args: args{
				size:    255,
				fade:    2,
				reverse: true, // Fade up
			},
			want: []int{255, 255, 255, 243, 209, 160, 105, 54, 17, 0},
		},
		{
			name: "Fade up 3",
			args: args{
				size:    255,
				fade:    3,
				reverse: false, // Fade up
			},
			want: []int{0, 17, 63, 127, 191, 237, 255, 255, 255, 255},
		},
		{
			name: "Fade up 4",
			args: args{
				size:    255,
				fade:    4,
				reverse: false, // Fade up
			},
			want: []int{0, 23, 83, 160, 225, 254, 255, 255, 255, 255},
		},
		{
			name: "Fade up 5",
			args: args{
				size:    255,
				fade:    5,
				reverse: false, // Fade up
			},
			want: []int{0, 29, 105, 191, 247, 255, 255, 255, 255, 255},
		},
		{
			name: "Fade up 6",
			args: args{
				size:    255,
				fade:    6,
				reverse: false, // Fade up
			},
			want: []int{0, 37, 127, 217, 255, 255, 255, 255, 255, 255},
		},
		{
			name: "Fade up 7",
			args: args{
				size:    255,
				fade:    7,
				reverse: false, // Fade up
			},
			want: []int{0, 45, 149, 237, 255, 255, 255, 255, 255, 255},
		},
		{
			name: "Fade up 8",
			args: args{
				size:    255,
				fade:    8,
				reverse: false, // Fade up
			},
			want: []int{0, 54, 171, 250, 255, 255, 255, 255, 255, 255},
		},
		{
			name: "Fade up 9",
			args: args{
				size:    255,
				fade:    9,
				reverse: false, // Fade up
			},
			want: []int{0, 63, 191, 255, 255, 255, 255, 255, 255, 255},
		},
		{
			name: "Fade up 10",
			args: args{
				size:    255,
				fade:    10,
				reverse: false, // Fade up
			},
			want: []int{0, 73, 209, 255, 255, 255, 255, 255, 255, 255},
		},
		{
			name: "Fade down 10",
			args: args{
				size:    255,
				fade:    10,
				reverse: true, // Fade down
			},
			want: []int{255, 255, 255, 255, 255, 255, 255, 255, 181, 45},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFadeValues(tt.args.size, tt.args.fade, tt.args.reverse); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %+v", got)
				t.Errorf("want %+v", tt.want)
			}
		})
	}
}

func Test_getFadeOnValues(t *testing.T) {
	type args struct {
		size int
		fade int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "10 slots on at 255",
			args: args{
				size: 255,
				fade: 10,
			},
			want: []int{255, 255, 255, 255, 255, 255, 255, 255, 255, 255},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFadeOnValues(tt.args.size, tt.args.fade); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %+v", got)
				t.Errorf("want %+v", tt.want)
			}
		})
	}
}

func TestFindSensitivity(t *testing.T) {
	type args struct {
		soundGain float32
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test for 0.05",
			args: args{
				soundGain: 0.05,
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindSensitivity(tt.args.soundGain); got != tt.want {
				t.Errorf("FindSensitivity() = %v, want %v", got, tt.want)
			}
		})
	}
}
