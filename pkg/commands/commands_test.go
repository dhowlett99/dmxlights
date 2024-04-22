package commands

import (
	"testing"

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
			if got := GetNumberOfFixtures(tt.args.sequenceNumber, tt.args.fixtures); got != tt.want {
				t.Errorf("GetNumberOfFixtures() got=%+v, want=%+v", got, tt.want)
			}
		})
	}
}
