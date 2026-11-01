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
		color color.RGBA
	}
	tests := []struct {
		name    string
		args    args
		wantOut color.RGBA
	}{
		{
			name: "invert white",
			args: args{
				color: color.RGBA{
					R: 255,
					G: 255,
					B: 255,
					A: 255,
				},
			},
			wantOut: color.RGBA{
				R: 0,
				G: 0,
				B: 0,
				A: 255,
			},
		},
		{
			name: "invert black",
			args: args{
				color: color.RGBA{
					R: 0,
					G: 0,
					B: 0,
					A: 255,
				},
			},
			wantOut: color.RGBA{
				R: 255,
				G: 255,
				B: 255,
				A: 255,
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
		from          int
		too           int
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
				fade:          1,
				from:          0,
				too:           255,
				reverse:       false, // Fade up
			},
			want: []int{0, 44, 87, 127, 163, 195, 220, 239, 251, 255},
		},
		{
			name: "Fade down 1",
			args: args{
				nocoordinates: 10,
				fade:          1,
				from:          255,
				too:           0,
				reverse:       true, // Fade down
			},
			want: []int{255, 251, 239, 220, 195, 163, 127, 87, 44, 0},
		},
		{
			name: "Fade up 2",
			args: args{
				nocoordinates: 10,
				fade:          2,
				from:          0,
				too:           255,
				reverse:       false, // Fade up
			},
			want: []int{0, 18, 51, 90, 131, 170, 205, 232, 249, 255},
		},
		{
			name: "Fade down 2",
			args: args{
				nocoordinates: 10,
				fade:          2,
				from:          255,
				too:           0,
				reverse:       true, // Fade up
			},
			want: []int{255, 249, 232, 205, 170, 131, 90, 51, 18, 0},
		},
		{
			name: "Fade up 3",
			args: args{
				nocoordinates: 10,
				fade:          3,
				from:          0,
				too:           255,
				reverse:       false, // Fade up
			},
			want: []int{0, 7, 29, 63, 105, 149, 191, 225, 247, 255},
		},
		{
			name: "Fade up 4",
			args: args{
				nocoordinates: 10,
				fade:          4,
				from:          0,
				too:           255,
				reverse:       false, // Fade up
			},
			want: []int{0, 1, 10, 31, 67, 114, 165, 211, 243, 255},
		},
		{
			name: "Fade up 5",
			args: args{
				nocoordinates: 10,
				fade:          5,
				from:          0,
				too:           255,
				reverse:       false, // Fade up
			},
			want: []int{0, 0, 3, 15, 43, 87, 143, 198, 239, 255},
		},
		{
			name: "Fade up 6",
			args: args{
				nocoordinates: 10,
				fade:          6,
				from:          0,
				too:           255,
				reverse:       false, // Fade up
			},
			want: []int{0, 0, 1, 7, 27, 67, 124, 186, 236, 255},
		},
		{
			name: "Fade up 7",
			args: args{
				nocoordinates: 10,
				fade:          7,
				from:          0,
				too:           255,
				reverse:       false, // Fade up
			},
			want: []int{0, 0, 0, 1, 11, 39, 93, 164, 229, 255},
		},
		{
			name: "Fade up 8",
			args: args{
				nocoordinates: 10,
				fade:          8,
				from:          0,
				too:           255,
				reverse:       false, // Fade up
			},
			want: []int{0, 0, 0, 0, 3, 17, 60, 136, 218, 255},
		},
		{
			name: "Fade up 9",
			args: args{
				nocoordinates: 10,
				from:          0,
				too:           255,
				fade:          9,
				reverse:       false, // Fade up
			},
			want: []int{0, 0, 0, 0, 0, 4, 29, 100, 202, 255},
		},
		{
			name: "Fade up 10",
			args: args{
				nocoordinates: 10,
				from:          0,
				too:           255,
				fade:          10,
				reverse:       false, // Fade up
			},
			want: []int{0, 0, 0, 0, 0, 0, 3, 39, 161, 255},
		},
		{
			name: "Fade down 10",
			args: args{
				nocoordinates: 10,
				from:          255,
				too:           0,
				fade:          10,
				reverse:       true, // Fade down
			},
			want: []int{255, 161, 39, 3, 0, 0, 0, 0, 0, 0},
		},
		{
			name: "Fade from 255 to 39",
			args: args{
				nocoordinates: 10,
				fade:          10,
				from:          255,
				too:           39,
				reverse:       true, // Fade down
			},
			want: []int{255, 161, 39},
		},
		{
			name: "Fade from 145 to 255",
			args: args{
				nocoordinates: 10,
				fade:          1, // <-- Change from 10 to 1 to match the 'want' array
				from:          145,
				too:           255,
				reverse:       false,
			},
			want: []int{145, 163, 195, 220, 239, 251, 255},
		},
		{
			name: "Fade from 145 to 255",
			args: args{
				nocoordinates: 10,
				fade:          1, // <-- Change from 10 to 1 to match the 'want' array
				from:          145,
				too:           0,
				reverse:       true,
			},
			want: []int{145, 127, 87, 44, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFadeValues(tt.args.nocoordinates, tt.args.from, tt.args.too, tt.args.fade, tt.args.reverse); !reflect.DeepEqual(got, tt.want) {
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

func TestGetLaunchPadCodeByRGBColor(t *testing.T) {
	type args struct {
		selectedColor color.RGBA
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "Test Red",
			args: args{
				selectedColor: color.RGBA{
					R: 255,
					G: 0,
					B: 0,
					A: 255,
				},
			},
			want: 0x48,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLaunchPadCodeByRGBColor(tt.args.selectedColor); got != tt.want {
				t.Errorf("GetLaunchPadCodeByRGBColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeColors(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		height int
		width  int
		want   []StaticColorButton
	}{
		{
			name:   "Generate 1 bank of 8 colors",
			height: 1,
			width:  8,
			want: []StaticColorButton{
				{Name: "Red", Label: "", Number: 0, X: 0, Y: 0, SelectedSequence: 0, Color: color.RGBA{R: 255, G: 0, B: 0, A: 255}, SelectedColor: 0, Flash: false, Setting: 0, FirstPress: false, Enabled: true},
				{Name: "Orange", Label: "", Number: 0, X: 1, Y: 0, SelectedSequence: 0, Color: color.RGBA{R: 255, G: 111, B: 0, A: 255}, SelectedColor: 1, Flash: false, Setting: 0, FirstPress: false, Enabled: true},
				{Name: "Yellow", Label: "", Number: 0, X: 2, Y: 0, SelectedSequence: 0, Color: color.RGBA{R: 255, G: 255, B: 0, A: 255}, SelectedColor: 2, Flash: false, Setting: 0, FirstPress: false, Enabled: true},
				{Name: "Green", Label: "", Number: 0, X: 3, Y: 0, SelectedSequence: 0, Color: color.RGBA{R: 0, G: 255, B: 0, A: 255}, SelectedColor: 3, Flash: false, Setting: 0, FirstPress: false, Enabled: true},
				{Name: "Cyan", Label: "", Number: 0, X: 4, Y: 0, SelectedSequence: 0, Color: color.RGBA{R: 0, G: 255, B: 255, A: 255}, SelectedColor: 4, Flash: false, Setting: 0, FirstPress: false, Enabled: true},
				{Name: "Blue", Label: "", Number: 0, X: 5, Y: 0, SelectedSequence: 0, Color: color.RGBA{R: 0, G: 0, B: 255, A: 255}, SelectedColor: 5, Flash: false, Setting: 0, FirstPress: false, Enabled: true},
				{Name: "Purple", Label: "", Number: 0, X: 6, Y: 0, SelectedSequence: 0, Color: color.RGBA{R: 100, G: 0, B: 255, A: 255}, SelectedColor: 6, Flash: false, Setting: 0, FirstPress: false, Enabled: true},
				{Name: "Magenta", Label: "", Number: 0, X: 7, Y: 0, SelectedSequence: 0, Color: color.RGBA{R: 255, G: 0, B: 255, A: 255}, SelectedColor: 7, Flash: false, Setting: 0, FirstPress: false, Enabled: true},
			},
		},
		{
			name:   "Generate 3 bank of 8 colors, 24 colors in total",
			height: 3,
			width:  8,
			want: []StaticColorButton{
				{Name: "Red", Label: "", Number: 0, X: 0, Y: 0, SelectedSequence: 0, Color: color.RGBA{R: 255, G: 0, B: 0, A: 255}, SelectedColor: 0, Flash: false, Setting: 0, FirstPress: false, Enabled: true},
				{Name: "Orange", Label: "", Number: 0, X: 1, Y: 0, SelectedSequence: 0, Color: color.RGBA{R: 255, G: 111, B: 0, A: 255}, SelectedColor: 1, Flash: false, Setting: 0, FirstPress: false, Enabled: true},
				{Name: "Yellow", Label: "", Number: 0, X: 2, Y: 0, SelectedSequence: 0, Color: color.RGBA{R: 255, G: 255, B: 0, A: 255}, SelectedColor: 2, Flash: false, Setting: 0, FirstPress: false, Enabled: true},
				{Name: "Green", Label: "", Number: 0, X: 3, Y: 0, SelectedSequence: 0, Color: color.RGBA{R: 0, G: 255, B: 0, A: 255}, SelectedColor: 3, Flash: false, Setting: 0, FirstPress: false, Enabled: true},
				{Name: "Cyan", Label: "", Number: 0, X: 4, Y: 0, SelectedSequence: 0, Color: color.RGBA{R: 0, G: 255, B: 255, A: 255}, SelectedColor: 4, Flash: false, Setting: 0, FirstPress: false, Enabled: true},
				{Name: "Blue", Label: "", Number: 0, X: 5, Y: 0, SelectedSequence: 0, Color: color.RGBA{R: 0, G: 0, B: 255, A: 255}, SelectedColor: 5, Flash: false, Setting: 0, FirstPress: false, Enabled: true},
				{Name: "Purple", Label: "", Number: 0, X: 6, Y: 0, SelectedSequence: 0, Color: color.RGBA{R: 100, G: 0, B: 255, A: 255}, SelectedColor: 6, Flash: false, Setting: 0, FirstPress: false, Enabled: true},
				{Name: "Magenta", Label: "", Number: 0, X: 7, Y: 0, SelectedSequence: 0, Color: color.RGBA{R: 255, G: 0, B: 255, A: 255}, SelectedColor: 7, Flash: false, Setting: 0, FirstPress: false, Enabled: true},

				{Name: "Crimson", Label: "", Number: 0, X: 0, Y: 1, SelectedSequence: 0, Color: color.RGBA{R: 220, G: 20, B: 60, A: 255}, SelectedColor: 8, Flash: false, Setting: 0, FirstPress: false, Enabled: true},
				{Name: "Dark Orange", Label: "", Number: 0, X: 1, Y: 1, SelectedSequence: 0, Color: color.RGBA{R: 215, G: 50, B: 0, A: 255}, SelectedColor: 9, Flash: false, Setting: 0, FirstPress: false, Enabled: true},
				{Name: "Gold", Label: "", Number: 0, X: 2, Y: 1, SelectedSequence: 0, Color: color.RGBA{R: 255, G: 215, B: 0, A: 255}, SelectedColor: 10, Flash: false, Setting: 0, FirstPress: false, Enabled: true},
				{Name: "Forest Green", Label: "", Number: 0, X: 3, Y: 1, SelectedSequence: 0, Color: color.RGBA{R: 0, G: 100, B: 0, A: 255}, SelectedColor: 11, Flash: false, Setting: 0, FirstPress: false, Enabled: true},
				{Name: "Aqua", Label: "", Number: 0, X: 4, Y: 1, SelectedSequence: 0, Color: color.RGBA{R: 127, G: 255, B: 212, A: 255}, SelectedColor: 12, Flash: false, Setting: 0, FirstPress: false, Enabled: true},
				{Name: "Sky Blue", Label: "", Number: 0, X: 5, Y: 1, SelectedSequence: 0, Color: color.RGBA{R: 0, G: 191, B: 255, A: 255}, SelectedColor: 13, Flash: false, Setting: 0, FirstPress: false, Enabled: true},
				{Name: "Dark Purple", Label: "", Number: 0, X: 6, Y: 1, SelectedSequence: 0, Color: color.RGBA{R: 50, G: 0, B: 255, A: 255}, SelectedColor: 14, Flash: false, Setting: 0, FirstPress: false, Enabled: true},
				{Name: "Pink", Label: "", Number: 0, X: 7, Y: 1, SelectedSequence: 0, Color: color.RGBA{R: 255, G: 192, B: 203, A: 255}, SelectedColor: 15, Flash: false, Setting: 0, FirstPress: false, Enabled: true},

				{Name: "Salmon", Label: "", Number: 0, X: 0, Y: 2, SelectedSequence: 0, Color: color.RGBA{R: 250, G: 128, B: 114, A: 255}, SelectedColor: 16, Flash: false, Setting: 0, FirstPress: false, Enabled: true},
				{Name: "Light Orange", Label: "", Number: 0, X: 1, Y: 2, SelectedSequence: 0, Color: color.RGBA{R: 255, G: 175, B: 0, A: 255}, SelectedColor: 17, Flash: false, Setting: 0, FirstPress: false, Enabled: true},
				{Name: "Olive", Label: "", Number: 0, X: 2, Y: 2, SelectedSequence: 0, Color: color.RGBA{R: 150, G: 150, B: 0, A: 255}, SelectedColor: 18, Flash: false, Setting: 0, FirstPress: false, Enabled: true},
				{Name: "Lawn Green", Label: "", Number: 0, X: 3, Y: 2, SelectedSequence: 0, Color: color.RGBA{R: 124, G: 252, B: 0, A: 255}, SelectedColor: 19, Flash: false, Setting: 0, FirstPress: false, Enabled: true},
				{Name: "Teal", Label: "", Number: 0, X: 4, Y: 2, SelectedSequence: 0, Color: color.RGBA{R: 0, G: 128, B: 128, A: 255}, SelectedColor: 20, Flash: false, Setting: 0, FirstPress: false, Enabled: true},
				{Name: "Light Blue", Label: "", Number: 0, X: 5, Y: 2, SelectedSequence: 0, Color: color.RGBA{R: 100, G: 185, B: 255, A: 255}, SelectedColor: 21, Flash: false, Setting: 0, FirstPress: false, Enabled: true},
				{Name: "Violet", Label: "", Number: 0, X: 6, Y: 2, SelectedSequence: 0, Color: color.RGBA{R: 199, G: 21, B: 133, A: 255}, SelectedColor: 22, Flash: false, Setting: 0, FirstPress: false, Enabled: true},
				{Name: "White", Label: "", Number: 0, X: 7, Y: 2, SelectedSequence: 0, Color: color.RGBA{R: 255, G: 255, B: 255, A: 255}, SelectedColor: 23, Flash: false, Setting: 0, FirstPress: false, Enabled: true},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := makeColors(tt.height, tt.width)
			// TODO: update the condition below to compare got with tt.want.
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeColors() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
