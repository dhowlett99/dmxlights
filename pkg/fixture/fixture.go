package fixture

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/go-yaml/yaml"
)

type Groups struct {
	Groups []Group `yaml:"groups"`
}

type Group struct {
	Name        string    `yaml:"name"`
	Description string    `yaml:"desription"`
	Fixtures    []Fixture `yaml:"fixtures"`
}

type Fixture struct {
	Address  int       `yaml:"startaddress"`
	Channels []Channel `yaml:"channels"`
}

type Channel struct {
	Number int    `yaml:"number"`
	Name   string `yaml:"name"`
}

type Fixtures struct {
	Fixtures []Fixture `yaml:"fixtures"`
}

func MakeFixtures() {

	groups := Groups{
		Groups: []Group{
			{
				Name:        "FOH Pars",
				Description: "Stairville Xbricks",
				Fixtures: []Fixture{
					{
						Address: 1,
						Channels: []Channel{
							{
								Number: 1,
								Name:   "Red1",
							},
							{
								Number: 2,
								Name:   "Green1",
							},
							{
								Number: 3,
								Name:   "Blue1",
							},
							{
								Number: 4,
								Name:   "Red2",
							},
							{
								Number: 5,
								Name:   "Green2",
							},
							{
								Number: 6,
								Name:   "Blue2",
							},
							{
								Number: 7,
								Name:   "Red3",
							},
							{
								Number: 8,
								Name:   "Green3",
							},
							{
								Number: 9,
								Name:   "Blue3",
							},
							{
								Number: 10,
								Name:   "Red4",
							},
							{
								Number: 11,
								Name:   "Green4",
							},
							{
								Number: 12,
								Name:   "Blue4",
							},
							{
								Number: 13,
								Name:   "Master",
							},
						},
					},
					{
						Address: 14,
						Channels: []Channel{
							{
								Number: 1,
								Name:   "Red1",
							},
							{
								Number: 2,
								Name:   "Green1",
							},
							{
								Number: 3,
								Name:   "Blue1",
							},
							{
								Number: 4,
								Name:   "Red2",
							},
							{
								Number: 5,
								Name:   "Green2",
							},
							{
								Number: 6,
								Name:   "Blue23",
							},
							{
								Number: 7,
								Name:   "Red3",
							},
							{
								Number: 8,
								Name:   "Green3",
							},
							{
								Number: 9,
								Name:   "Blue3",
							},
							{
								Number: 10,
								Name:   "Red4",
							},
							{
								Number: 11,
								Name:   "Green4",
							},
							{
								Number: 12,
								Name:   "Blue4",
							},
							{
								Number: 13,
								Name:   "MasterDimmer",
							},
						},
					},
				},
			},
			{
				Name:        "Uplighers",
				Description: "Chauvet Color Rail IRC",
				Fixtures: []Fixture{
					{
						Address: 27,
						Channels: []Channel{
							{
								Number: 1,
								Name:   "MasterDimmer",
							},
							// Segment 1
							{
								Number: 2,
								Name:   "Red1",
							},
							{
								Number: 3,
								Name:   "Green1",
							},
							{
								Number: 4,
								Name:   "Blue1",
							},
							// Segment 2
							{
								Number: 5,
								Name:   "Red2",
							},
							{
								Number: 6,
								Name:   "Green2",
							},
							{
								Number: 7,
								Name:   "Blue2",
							},

							// Segment 3
							{
								Number: 8,
								Name:   "Red3",
							},
							{
								Number: 9,
								Name:   "Green3",
							},
							{
								Number: 10,
								Name:   "Blue3",
							},

							// Segment 4
							{
								Number: 11,
								Name:   "Red4",
							},
							{
								Number: 12,
								Name:   "Green4",
							},
							{
								Number: 13,
								Name:   "Blue4",
							},
							// Segment 5
							{
								Number: 14,
								Name:   "Red5",
							},
							{
								Number: 15,
								Name:   "Green5",
							},
							{
								Number: 16,
								Name:   "Blue5",
							},
							// Segment 6
							{
								Number: 17,
								Name:   "Red6",
							},
							{
								Number: 18,
								Name:   "Green6",
							},
							{
								Number: 19,
								Name:   "Blue6",
							},
							// Segment 7
							{
								Number: 20,
								Name:   "Red7",
							},
							{
								Number: 21,
								Name:   "Green7",
							},
							{
								Number: 22,
								Name:   "Blue7",
							},
							// Segment 8
							{
								Number: 23,
								Name:   "Red8",
							},
							{
								Number: 24,
								Name:   "Green8",
							},
							{
								Number: 25,
								Name:   "Blue8",
							},
							// Strobe Channel.
							{
								Number: 26,
								Name:   "Strobe",
							},
						},
					},
				},
			},
		},
	}

	b, err := yaml.Marshal(&groups)
	if err != nil {
		fmt.Printf("error marshalling fixtures\n")
	}

	filename := "fixtures.yaml"

	_, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("error creating yaml file\n")
	}
	err = ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		fmt.Printf("error writing yaml file\n")
	}
}

func LoadFixtures() *Groups {
	filename := "fixtures.yaml"

	_, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Printf("error loading fixtures.yaml file\n")
	}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("error writing yaml file\n")
	}

	groups := &Groups{}
	err = yaml.Unmarshal(data, groups)
	if err != nil {
		fmt.Printf("error marshalling fixtures: %s\n", err.Error())
	}

	return groups

}

func FixtureReceiver(
	channel chan common.Event,
	fixture int,
	command common.Sequence,
	commandChannel chan common.Command,
	mySequenceNumber int,
	eventsForLauchpad chan common.ALight) {

	for {

		event := <-channel

		e := common.ALight{
			X:          fixture,
			Y:          mySequenceNumber - 1,
			Brightness: 3,
			Red:        event.Color.R,
			Green:      event.Color.G,
			Blue:       event.Color.B,
		}
		eventsForLauchpad <- e
	}
}
