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
	"fmt"
	"os"
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

func Test_lookUpChannelNumberByNameInFixtureDefinition(t *testing.T) {

	// Get a list of all the fixtures in the groups.
	fixturesConfig, err := LoadFixtures("../../fixtures.yaml")
	if err != nil {
		fmt.Printf("dmxlights: error failed to load fixtures: %s\n", err.Error())
		os.Exit(1)
	}

	type args struct {
		group        int
		switchNumber int
		channelName  string
		fixtures     *Fixtures
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "simple test",
			args: args{
				group:        100,
				switchNumber: 1,
				channelName:  "Master",
				fixtures:     fixturesConfig,
			},
			want:    5,
			wantErr: false,
		},
		{
			name: "simple test",
			args: args{
				group:        100,
				switchNumber: 1,
				channelName:  "White1",
				fixtures:     fixturesConfig,
			},
			want:    4,
			wantErr: false,
		},
		{
			name: "simple test",
			args: args{
				group:        100,
				switchNumber: 1,
				channelName:  "ProgramSpeed",
				fixtures:     fixturesConfig,
			},
			want:    7,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := lookUpChannelNumberByNameInFixtureDefinition(tt.args.group, tt.args.switchNumber, tt.args.channelName, tt.args.fixtures)
			if (err != nil) != tt.wantErr {
				t.Errorf("lookUpChannelNumberByNameInFixtureDefinition() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("lookUpChannelNumberByNameInFixtureDefinition() = %v, want %v", got, tt.want)
			}
		})
	}
}
