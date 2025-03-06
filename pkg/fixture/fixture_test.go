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

func Test_getNumberOfFixtures(t *testing.T) {

	oneColorChannels := []Channel{
		{Name: "Red1"},
	}

	eightColorChannels := []Channel{
		{Name: "Red1"},
		{Name: "Red2"},
		{Name: "Red3"},
		{Name: "Red4"},
		{Name: "Red5"},
		{Name: "Red6"},
		{Name: "Red7"},
		{Name: "Red8"},
	}

	type args struct {
		sequenceNumber     int
		fixtures           *Fixtures
		allPosibleFixtures bool
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "eight fixtures",
			args: args{
				allPosibleFixtures: false,
				sequenceNumber:     0,
				fixtures: &Fixtures{
					Fixtures: []Fixture{

						{Name: "fixture1", Number: 1, Group: 1, Channels: oneColorChannels},
						{Name: "fixture2", Number: 2, Group: 1, Channels: oneColorChannels},
						{Name: "fixture3", Number: 3, Group: 1, Channels: oneColorChannels},
						{Name: "fixture4", Number: 4, Group: 1, Channels: oneColorChannels},
						{Name: "fixture5", Number: 5, Group: 1, Channels: oneColorChannels},
						{Name: "fixture6", Number: 6, Group: 1, Channels: oneColorChannels},
						{Name: "fixture7", Number: 7, Group: 1, Channels: oneColorChannels},
						{Name: "fixture8", Number: 8, Group: 1, Channels: oneColorChannels},

						{Name: "fixture1", Number: 1, Group: 2, Channels: oneColorChannels},
						{Name: "fixture2", Number: 2, Group: 2, Channels: oneColorChannels},
						{Name: "fixture3", Number: 3, Group: 2, Channels: oneColorChannels},

						{Name: "fixture1", Number: 1, Group: 3, Channels: oneColorChannels},
						{Name: "fixture2", Number: 2, Group: 3, Channels: oneColorChannels},
						{Name: "fixture3", Number: 3, Group: 3, Channels: oneColorChannels},
						{Name: "fixture4", Number: 4, Group: 3, Channels: oneColorChannels},

						{Name: "fixture1", Number: 1, Group: 4, Channels: oneColorChannels},
						{Name: "fixture2", Number: 2, Group: 4, Channels: oneColorChannels},
					},
				},
			},
			want: 8,
		},

		{
			name: "three fixtures",
			args: args{
				allPosibleFixtures: false,
				sequenceNumber:     1,
				fixtures: &Fixtures{
					Fixtures: []Fixture{

						{Name: "fixture1", Number: 1, Group: 1},
						{Name: "fixture2", Number: 2, Group: 1},
						{Name: "fixture3", Number: 3, Group: 1},
						{Name: "fixture4", Number: 4, Group: 1},
						{Name: "fixture5", Number: 5, Group: 1},
						{Name: "fixture6", Number: 6, Group: 1},
						{Name: "fixture7", Number: 7, Group: 1},
						{Name: "fixture8", Number: 8, Group: 1},

						{Name: "fixture1", Number: 1, Group: 2},
						{Name: "fixture2", Number: 2, Group: 2},
						{Name: "fixture3", Number: 3, Group: 2},

						{Name: "fixture1", Number: 1, Group: 3},
						{Name: "fixture2", Number: 2, Group: 3},
						{Name: "fixture3", Number: 3, Group: 3},
						{Name: "fixture4", Number: 4, Group: 3},

						{Name: "fixture1", Number: 1, Group: 4},
						{Name: "fixture2", Number: 2, Group: 4},
					},
				},
			},
			want: 3,
		},

		{
			// Uplighters with their use_channels set to 8.
			name: "four uplighters with their use_channels (NumberChannels) set to 8 fixtures so 32 in all.",
			args: args{
				allPosibleFixtures: false,
				sequenceNumber:     1,
				fixtures: &Fixtures{
					Fixtures: []Fixture{
						{Name: "fixture1", Number: 1, Group: 2, MultiFixtureDevice: true, NumberSubFixtures: 8, Channels: eightColorChannels},
						{Name: "fixture2", Number: 2, Group: 2, MultiFixtureDevice: true, NumberSubFixtures: 8, Channels: eightColorChannels},
						{Name: "fixture3", Number: 3, Group: 2, MultiFixtureDevice: true, NumberSubFixtures: 8, Channels: eightColorChannels},
						{Name: "fixture3", Number: 3, Group: 2, MultiFixtureDevice: true, NumberSubFixtures: 8, Channels: eightColorChannels},
					},
				},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetNumberOfFixturesInGroup(tt.args.sequenceNumber, tt.args.fixtures); got != tt.want {
				t.Errorf("GetNumberOfFixtures() got=%+v, want=%+v", got, tt.want)
			}
		})
	}
}

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
			if got := GetGoboByName(tt.args.myFixtureNumber, tt.args.mySequenceNumber, tt.args.selectedGobo, tt.args.fixtures); got != tt.want {
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
		swiTchNumber   int
		stateNumber    int16
	}

	fixturesConfig := &Fixtures{
		Fixtures: []Fixture{
			{
				Type:   "switch",
				Group:  3,
				Number: 1,
				States: []State{
					{
						Name:   "On",
						Number: 1,
						Actions: []Action{
							{
								Name:  "Fade",
								Mode:  "Chase",
								Speed: "Fast",
							},
						},
					},
					{
						Name:   "Off",
						Number: 2,
					},
				},
			},
			{
				Type:   "switch",
				Group:  3,
				Number: 2,
				States: []State{
					{
						Name:   "On",
						Number: 1,
					},
					{
						Name:   "Off",
						Number: 2,
					},
				},
			},
			{
				Type:   "switch",
				Group:  3,
				Number: 3,
				States: []State{
					{
						Name:   "Off",
						Number: 1,
					},
					{
						Name:   "Fade",
						Number: 2,
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
		want Action
	}{
		{
			name: "find the right action based on switch and state number",
			args: args{
				fixturesConfig: fixturesConfig,
				swiTchNumber:   1,
				stateNumber:    1,
			},

			want: Action{
				Name:  "Fade",
				Mode:  "Chase",
				Speed: "Fast",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetSwitchAction(tt.args.swiTchNumber, tt.args.stateNumber, tt.args.fixturesConfig)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSwitchSpeeds() got = %+v\n", got)
				t.Errorf("GetSwitchSpeeds() want= %+v\n", tt.want)
			}
		})
	}
}
