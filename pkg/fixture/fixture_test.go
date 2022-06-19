package fixture

import "testing"

func Test_calculateMaxDMX(t *testing.T) {

	type args struct {
		MaxDegreeValueForFixture int
		Value                    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "A scanner that can do 540 degrees",
			args: args{
				MaxDegreeValueForFixture: 540,
				Value:                    255,
			},
			want: 170,
		},
		{
			name: "A scanner that can do 540 degrees",
			args: args{
				MaxDegreeValueForFixture: 540,
				Value:                    128,
			},
			want: 85,
		},
		{
			name: "a scanner that can do 360 degrees",
			args: args{
				MaxDegreeValueForFixture: 360,
				Value:                    255,
			},
			want: 255,
		},
		{
			name: "a scanner that can only do less than 360 degrees",
			args: args{
				MaxDegreeValueForFixture: 240,
				Value:                    255, // this is indicating we want 360 degs.
			},
			want: 180,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := limitDmxValue(&tt.args.MaxDegreeValueForFixture, tt.args.Value); got != tt.want {
				t.Errorf("calculateMaxDMX() = %v, want %v", got, tt.want)
			}
		})
	}
}
