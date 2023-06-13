// Copyright (C) 2022,2023 dhowlett99.
// This is the dmxlight main sequencer test code.
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

package sequence

import (
	"reflect"
	"testing"

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
)

func Test_getNumberOfFixtures(t *testing.T) {

	oneColorChannels := []fixture.Channel{
		{Name: "Red1"},
	}

	eightColorChannels := []fixture.Channel{
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
		fixtures           *fixture.Fixtures
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
				fixtures: &fixture.Fixtures{
					Fixtures: []fixture.Fixture{

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
				fixtures: &fixture.Fixtures{
					Fixtures: []fixture.Fixture{

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
				fixtures: &fixture.Fixtures{
					Fixtures: []fixture.Fixture{
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
			if got := getNumberOfFixtures(tt.args.sequenceNumber, tt.args.fixtures); got != tt.want {
				t.Errorf("getNumberOfFixtures() got=%+v, want=%+v", got, tt.want)
			}
		})
	}
}

func Test_replaceRGBcolorsInSteps(t *testing.T) {

	full := 255
	type args struct {
		steps  []common.Step
		colors []common.Color
	}
	tests := []struct {
		name string
		args args
		want []common.Step
	}{
		{
			name: "simple case",
			args: args{
				steps: []common.Step{
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
							1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
							2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
						},
					},
				},
				colors: []common.Color{
					{R: 255, G: 0, B: 0},
					{R: 0, G: 255, B: 0},
					{R: 0, G: 0, B: 255},
				},
			},
			want: []common.Step{
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
					},
				},
			},
		},
		{
			name: "replace a number of colors with just one.",
			args: args{
				steps: []common.Step{
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
							1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
							2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
						},
					},
				},
				colors: []common.Color{
					{R: 0, G: 255, B: 0},
				},
			},
			want: []common.Step{
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceRGBcolorsInSteps(tt.args.steps, tt.args.colors); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("replaceRGBcolorsInSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_invertRGBColorsInSteps(t *testing.T) {

	full := 255
	type args struct {
		steps  []common.Step
		colors []common.Color
	}
	tests := []struct {
		name string
		args args
		want []common.Step
	}{
		{
			name: "invert a single color.",
			args: args{
				steps: []common.Step{
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
							1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
							2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						},
					},
					{
						Fixtures: map[int]common.Fixture{
							0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}},
						},
					},
				},
				colors: []common.Color{
					{R: 0, G: 255, B: 0},
				},
			},
			want: []common.Step{
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					},
				},
				{
					Fixtures: map[int]common.Fixture{
						0: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
						1: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
						2: {MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := invertRGBColorsInSteps(tt.args.steps, tt.args.colors); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("invertRGBColors() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
