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
	"image/color"
	"reflect"
	"testing"
)

func Test_invertColor(t *testing.T) {
	type args struct {
		color color.NRGBA
	}
	tests := []struct {
		name    string
		args    args
		wantOut color.NRGBA
	}{
		{
			name: "invert white",
			args: args{
				color: color.NRGBA{
					R: 255,
					G: 255,
					B: 255,
				},
			},
			wantOut: color.NRGBA{
				R: 0,
				G: 0,
				B: 0,
			},
		},
		{
			name: "invert black",
			args: args{
				color: color.NRGBA{
					R: 0,
					G: 0,
					B: 0,
				},
			},
			wantOut: color.NRGBA{
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
		nocoordinates int
		size          float64
		fade          int
		reverse       bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Fade up 1",
			args: args{
				nocoordinates: 10,
				size:          255,
				fade:          1,
				reverse:       false, // Fade up
			},
			want: []int{0, 44, 87, 127, 163, 195, 220, 239, 251, 255},
		},
		{
			name: "Fade down 1",
			args: args{
				nocoordinates: 10,
				size:          255,
				fade:          1,
				reverse:       true, // Fade down
			},
			want: []int{255, 251, 239, 220, 195, 163, 127, 87, 44, 0},
		},
		{
			name: "Fade up 2",
			args: args{
				nocoordinates: 10,
				size:          255,
				fade:          2,
				reverse:       false, // Fade up
			},
			want: []int{0, 18, 51, 90, 131, 170, 205, 232, 249, 255},
		},
		{
			name: "Fade down 2",
			args: args{
				nocoordinates: 10,
				size:          255,
				fade:          2,
				reverse:       true, // Fade up
			},
			want: []int{255, 249, 232, 205, 170, 131, 90, 51, 18, 0},
		},
		{
			name: "Fade up 3",
			args: args{
				nocoordinates: 10,
				size:          255,
				fade:          3,
				reverse:       false, // Fade up
			},
			want: []int{0, 7, 29, 63, 105, 149, 191, 225, 247, 255},
		},
		{
			name: "Fade up 4",
			args: args{
				nocoordinates: 10,
				size:          255,
				fade:          4,
				reverse:       false, // Fade up
			},
			want: []int{0, 1, 10, 31, 67, 114, 165, 211, 243, 255},
		},
		{
			name: "Fade up 5",
			args: args{
				nocoordinates: 10,
				size:          255,
				fade:          5,
				reverse:       false, // Fade up
			},
			want: []int{0, 0, 3, 15, 43, 87, 143, 198, 239, 255},
		},
		{
			name: "Fade up 6",
			args: args{
				nocoordinates: 10,
				size:          255,
				fade:          6,
				reverse:       false, // Fade up
			},
			want: []int{0, 0, 1, 7, 27, 67, 124, 186, 236, 255},
		},
		{
			name: "Fade up 7",
			args: args{
				nocoordinates: 10,
				size:          255,
				fade:          7,
				reverse:       false, // Fade up
			},
			want: []int{0, 0, 0, 1, 11, 39, 93, 164, 229, 255},
		},
		{
			name: "Fade up 8",
			args: args{
				nocoordinates: 10,
				size:          255,
				fade:          8,
				reverse:       false, // Fade up
			},
			want: []int{0, 0, 0, 0, 3, 17, 60, 136, 218, 255},
		},
		{
			name: "Fade up 9",
			args: args{
				nocoordinates: 10,
				size:          255,
				fade:          9,
				reverse:       false, // Fade up
			},
			want: []int{0, 0, 0, 0, 0, 4, 29, 100, 202, 255},
		},
		{
			name: "Fade up 10",
			args: args{
				nocoordinates: 10,
				size:          255,
				fade:          10,
				reverse:       false, // Fade up
			},
			want: []int{0, 0, 0, 0, 0, 0, 3, 39, 161, 255},
		},
		{
			name: "Fade down 10",
			args: args{
				nocoordinates: 10,
				size:          255,
				fade:          10,
				reverse:       true, // Fade down
			},
			want: []int{255, 161, 39, 3, 0, 0, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFadeValues(tt.args.nocoordinates, tt.args.size, tt.args.fade, tt.args.reverse); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %+v", got)
				t.Errorf("want %+v", tt.want)
			}
		})
	}
}

func Test_getFadeOnValues(t *testing.T) {
	type args struct {
		brightness int
		fade       int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "10 slots on at 255",
			args: args{
				brightness: 255,
				fade:       10,
			},
			want: []int{255, 255, 255, 255, 255, 255, 255, 255, 255, 255},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFadeOnValues(tt.args.brightness, tt.args.fade); !reflect.DeepEqual(got, tt.want) {
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

func TestReverseDmx(t *testing.T) {
	type args struct {
		n uint8
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		{
			name: "Reverse 255",
			args: args{
				n: 255,
			},
			want: 0,
		},
		{
			name: "Reverse 50",
			args: args{
				n: 50,
			},
			want: 205,
		},
		{
			name: "Reverse 0",
			args: args{
				n: 0,
			},
			want: 255,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseDmx(tt.args.n); got != tt.want {
				t.Errorf("ReverseDmx() = %v, want %v", got, tt.want)
			}
		})
	}
}
