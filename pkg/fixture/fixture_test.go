// Copyright (C) 2022, 2023 dhowlett99.
// This is the dmxlights fixture control test code.
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

import (
	"reflect"
	"testing"
)

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

func TestFindGobo(t *testing.T) {

	fixturesConfig := &Fixtures{
		Fixtures: []Fixture{
			{
				Name:   "fixture1",
				Group:  1,
				Number: 1,
				Channels: []Channel{
					{
						Name: "Red",
					},
					{
						Name: "Green",
					},
					{
						Name: "Gobo",
						Settings: []Setting{
							{
								Name:   "Yellow Circle",
								Number: 1,
							},
						},
					},
				},
			},
			{
				Name:   "fixture2",
				Group:  2,
				Number: 2,
				Channels: []Channel{
					{
						Name: "Gobo",
						Settings: []Setting{
							{
								Name:   "Yellow Circle",
								Number: 1,
							},
							{
								Name:   "White Circle",
								Number: 2,
							},
						},
					},
					{
						Name: "Shutter",
					},
				},
			},
			{
				Name:   "fixture3",
				Group:  3,
				Number: 3,
				Channels: []Channel{
					{
						Name: "ProgramSpeed",
					},
				},
			},
		},
	}

	type args struct {
		myFixtureNumber  int
		mySequenceNumber int
		selectedGobo     string
		fixtures         *Fixtures
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "find White gobo",
			args: args{
				myFixtureNumber:  1,
				mySequenceNumber: 1,
				selectedGobo:     "White",
				fixtures:         fixturesConfig,
			},
			want: 2,
		},
		{
			name: "find Yellow gobo",
			args: args{
				myFixtureNumber:  0,
				mySequenceNumber: 0,
				selectedGobo:     "Yellow",
				fixtures:         fixturesConfig,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindGobo(tt.args.myFixtureNumber, tt.args.mySequenceNumber, tt.args.selectedGobo, tt.args.fixtures); got != tt.want {
				t.Errorf("FindGobo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckFixturesAreTheSame(t *testing.T) {

	fixturesConfig := &Fixtures{
		Fixtures: []Fixture{
			{
				Name:   "fixture1",
				Group:  1,
				Number: 1,
				Channels: []Channel{
					{
						Name: "Red",
					},
					{
						Name: "Green",
					},
					{
						Name: "Gobo",
						Settings: []Setting{
							{
								Name:   "Yellow Circle",
								Number: 1,
							},
						},
					},
				},
			},
			{
				Name:   "fixture2",
				Group:  2,
				Number: 2,
				Channels: []Channel{
					{
						Name: "Gobo",
						Settings: []Setting{
							{
								Name:   "Yellow Circle",
								Number: 1,
							},
							{
								Name:   "White Circle",
								Number: 2,
							},
						},
					},
					{
						Name: "Shutter",
					},
				},
			},
			{
				Name:   "fixture3",
				Group:  3,
				Number: 3,
				Channels: []Channel{
					{
						Name: "ProgramSpeed",
					},
				},
			},
		},
	}

	differentConfig := &Fixtures{
		Fixtures: []Fixture{
			{
				Name:   "fixture1",
				Group:  1,
				Number: 1,
				Channels: []Channel{
					{
						Name: "DIFFERENT Red",
					},
					{
						Name: "Green",
					},
					{
						Name: "Gobo",
						Settings: []Setting{
							{
								Name:   "Yellow Circle",
								Number: 1,
							},
						},
					},
				},
			},
			{
				Name:   "fixture2",
				Group:  2,
				Number: 2,
				Channels: []Channel{
					{
						Name: "Gobo",
						Settings: []Setting{
							{
								Name:   "Yellow Circle",
								Number: 1,
							},
							{
								Name:   "White Circle",
								Number: 2,
							},
						},
					},
					{
						Name: "Shutter",
					},
				},
			},
			{
				Name:   "fixture3",
				Group:  3,
				Number: 3,
				Channels: []Channel{
					{
						Name: "ProgramSpeed",
					},
				},
			},
		},
	}

	type args struct {
		fixtures    *Fixtures
		startConfig *Fixtures
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "same configs",
			args: args{
				fixtures:    fixturesConfig,
				startConfig: fixturesConfig,
			},
			want: true,
		},
		{
			name: "different configs",
			args: args{
				fixtures:    fixturesConfig,
				startConfig: differentConfig,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := CheckFixturesAreTheSame(tt.args.fixtures, tt.args.startConfig); got != tt.want {
				t.Errorf("CheckFixturesAreTheSame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSwitchSpeeds(t *testing.T) {
	type args struct {
		fixturesConfig *Fixtures
	}

	fixturesConfig := &Fixtures{
		Fixtures: []Fixture{
			{
				Type:   "switch",
				Group:  3,
				Number: 1,
				States: []State{
					{
						Name: "On",
					},
					{
						Name: "Off",
					},
				},
			},
			{
				Type:   "switch",
				Group:  3,
				Number: 2,
				States: []State{
					{
						Name: "On",
					},
					{
						Name: "Off",
					},
				},
			},
			{
				Type:   "switch",
				Group:  3,
				Number: 3,
				States: []State{
					{
						Name: "Off",
					},
					{
						Name: "Fade",
						Actions: []Action{
							{
								Mode:  "Chase",
								Speed: "Fast",
							},
						},
					},
				},
			},
		},
	}

	tests := []struct {
		name string
		args args
		want map[int]int
	}{
		{
			name: "gold path",
			args: args{
				fixturesConfig: fixturesConfig,
			},

			want: map[int]int{
				2: 8,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSwitchSpeeds(tt.args.fixturesConfig); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSwitchSpeeds() = %v, want %v", got, tt.want)
			}
		})
	}
}
