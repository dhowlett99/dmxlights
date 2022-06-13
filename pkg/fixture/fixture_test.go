package fixture

import "testing"

func Test_calculateMaxDMX(t *testing.T) {

	fiveFourty := 540
	threeSixty := 360
	twoFourty := 240

	type args struct {
		MaxDegreeValueForFixture *int
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
				MaxDegreeValueForFixture: &fiveFourty,
				Value:                    255,
			},
			want: 170,
		},
		{
			name: "a scanner that can do 360 degrees",
			args: args{
				MaxDegreeValueForFixture: &threeSixty,
				Value:                    255,
			},
			want: 255,
		},
		{
			name: "a scanner that can do less than 360 degrees",
			args: args{
				MaxDegreeValueForFixture: &twoFourty,
				Value:                    255,
			},
			want: 255,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := limitDmxValue(tt.args.MaxDegreeValueForFixture, tt.args.Value); got != tt.want {
				t.Errorf("calculateMaxDMX() = %v, want %v", got, tt.want)
			}
		})
	}
}
