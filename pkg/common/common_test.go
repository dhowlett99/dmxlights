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
