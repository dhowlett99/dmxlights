package sound

import "testing"

func Test_findLargest(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				values: []int{119, 7, 0, 0},
			},
			want: 1,
		},
		{
			name: "test2",
			args: args{
				values: []int{12342, 7293, 4930, 3378, 2364, 1661, 1124, 732, 489, 309},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findGain(tt.args.values); got != tt.want {
				t.Errorf("findGain() = %v, want %v", got, tt.want)
			}
		})
	}
}
