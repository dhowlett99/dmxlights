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
