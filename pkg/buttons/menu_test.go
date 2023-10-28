package buttons

import (
	"reflect"
	"testing"
)

func Test_newMenu(t *testing.T) {
	type args struct {
		chaser bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "not a chaser",
			args: args{
				chaser: false,
			},
			want: []int{NORMAL, FUNCTION, STATUS},
		},
		{
			name: "is a chaser",
			args: args{
				chaser: true,
			},
			want: []int{NORMAL, FUNCTION, CHASER_DISPLAY, CHASER_FUNCTION, STATUS},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newMenu(tt.args.chaser); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newMenu() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getNextMenuItem(t *testing.T) {
	type args struct {
		selectedMode int
		chaser       bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "get next item, normal in not a chaser mode,",
			args: args{
				selectedMode: NORMAL,
				chaser:       false,
			},
			want: FUNCTION,
		},
		{
			name: "get next item, normal in not a chaser mode,",
			args: args{
				selectedMode: FUNCTION,
				chaser:       false,
			},
			want: STATUS,
		},
		{
			name: "get next item, normal in not a chaser mode,",
			args: args{
				selectedMode: STATUS,
				chaser:       false,
			},
			want: NORMAL,
		},

		// Chaser mode.
		{
			name: "get next item, normal and in chaser mode,",
			args: args{
				selectedMode: NORMAL,
				chaser:       true,
			},
			want: FUNCTION,
		},
		{
			name: "get next item, function and in chaser mode,",
			args: args{
				selectedMode: FUNCTION,
				chaser:       true,
			},
			want: CHASER_DISPLAY,
		},
		{
			name: "get next item, chaser display and in chaser mode,",
			args: args{
				selectedMode: CHASER_DISPLAY,
				chaser:       true,
			},
			want: CHASER_FUNCTION,
		},
		{
			name: "get next item, chaser function and in chaser mode,",
			args: args{
				selectedMode: CHASER_FUNCTION,
				chaser:       true,
			},
			want: STATUS,
		},
		{
			name: "get next item, status and in chaser mode,",
			args: args{
				selectedMode: STATUS,
				chaser:       true,
			},
			want: NORMAL,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNextMenuItem(tt.args.selectedMode, tt.args.chaser); got != tt.want {
				t.Errorf("getNextMenuItem() = %v, want %v", got, tt.want)
			}
		})
	}
}
