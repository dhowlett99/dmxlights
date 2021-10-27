package fixture

import (
	"errors"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/launchpad"
	"github.com/go-yaml/yaml"
	"github.com/oliread/usbdmx/ft232"
)

type Fixtures struct {
	Fixtures []Fixture `yaml:"fixtures"`
}

type Color struct {
	R int `yaml:"red"`
	G int `yaml:"green"`
	B int `yaml:"blue"`
}

type Value struct {
	Channel int16 `yaml:"channel"`
	Setting int16 `yaml:"setting"`
}

type State struct {
	Name        string  `yaml:"name"`
	Values      []Value `yaml:"values"`
	ButtonColor Color   `yaml:"buttoncolor"`
}

type Switch struct {
	Name        string  `yaml:"name"`
	Number      int16   `yaml:"number"`
	Description string  `yaml:"description"`
	States      []State `yaml:"states"`
}

type Fixture struct {
	Name        string    `yaml:"name"`
	Description string    `yaml:"description"`
	Type        string    `yaml:"type"`
	Group       int       `yaml:"group"`
	Address     int16     `yaml:"address"`
	Channels    []Channel `yaml:"channels"`
	Switches    []Switch  `yaml:"switches"`
}

type Channel struct {
	Number  int16  `yaml:"number"`
	Name    string `yaml:"name"`
	Value   int16  `yaml:"value"`
	Comment string `yaml:"comment"`
}

func LoadFixtures() (fixtures *Fixtures, err error) {
	filename := "fixtures.yaml"

	_, err = os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		return nil, errors.New("error: loading fixtures.yaml file: " + err.Error())
	}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.New("error: reading yaml file: " + err.Error())
	}

	fixtures = &Fixtures{}
	err = yaml.Unmarshal(data, fixtures)
	if err != nil {
		return nil, errors.New("error: marshalling fixtures: " + err.Error())
	}
	return fixtures, nil
}

// FixtureReceivers are created by the sequence and are used to receive step instructions.
// Each FixtureReceiver knows which step they belong too and when triggered they start a fade up
// and fade down events which get sent to the launchpad lamps and the DMX fixtures.
func FixtureReceiver(sequence common.Sequence,
	mySequenceNumber int,
	myFixtureNumber int,
	channels []chan common.FixtureCommand,
	eventsForLauchpad chan common.ALight,
	dmxController *ft232.DMXController,
	fixtures *Fixtures) {

	cmd := common.FixtureCommand{}
	fadeUp := []int{0, 66, 127, 180, 220, 246, 255}
	fadeDown := []int{255, 246, 220, 189, 127, 66, 0}

	for {
		select {
		case cmd = <-channels[myFixtureNumber]:
			if cmd.Tick {
				// positions can have many fixtures play at the same time.
				positions := cmd.Positions[cmd.CurrentPosition]
				for _, position := range positions {
					if cmd.CurrentPosition == position.StartPosition {
						// Short ciruit the soft fade if we are a scanner.
						if cmd.Type == "scanner" {
							if position.Fixture == myFixtureNumber {
								MapFixtures(mySequenceNumber, dmxController, myFixtureNumber, position.Color.R, position.Color.G, position.Color.B, position.Pan, position.Tilt, position.Shutter, position.Gobo, fixtures, cmd.Blackout, cmd.Master, cmd.Master)
								if !cmd.Hide {
									launchpad.LightLamp(mySequenceNumber, myFixtureNumber, position.Color.R, position.Color.G, position.Color.B, cmd.Master, eventsForLauchpad)
								}
							}
						} else {
							if position.Fixture == myFixtureNumber {
								// Now kick off the back end which drives the RGB fixture.
								go func() {
									var R int
									var G int
									var B int

									if cmd.Inverted {
										for _, value := range fadeDown {
											R = int((float64(position.Color.R) / 100) * (float64(value) / 2.55))
											G = int((float64(position.Color.G) / 100) * (float64(value) / 2.55))
											B = int((float64(position.Color.B) / 100) * (float64(value) / 2.55))
											if !cmd.Hide {
												launchpad.LightLamp(mySequenceNumber, myFixtureNumber, R, G, B, cmd.Master, eventsForLauchpad)
											}
											MapFixtures(mySequenceNumber, dmxController, myFixtureNumber, R, G, B, 0, 0, 0, 0, fixtures, cmd.Blackout, cmd.Master, cmd.Master)
											time.Sleep(cmd.FadeTime / 4) // Fade down time.
										}
										time.Sleep(cmd.FadeTime / 4) // Fade off time.
									}
									for _, value := range fadeUp {
										R = int((float64(position.Color.R) / 100) * (float64(value) / 2.55))
										G = int((float64(position.Color.G) / 100) * (float64(value) / 2.55))
										B = int((float64(position.Color.B) / 100) * (float64(value) / 2.55))
										if !cmd.Hide {
											launchpad.LightLamp(mySequenceNumber, myFixtureNumber, R, G, B, cmd.Master, eventsForLauchpad)
										}
										// Now ask DMX to actually light the real fixture.
										MapFixtures(mySequenceNumber, dmxController, myFixtureNumber, R, G, B, 0, 0, 0, 0, fixtures, cmd.Blackout, cmd.Master, cmd.Master)
										time.Sleep(cmd.FadeTime / 4) // Fade up Time.
									}
									for x := 0; x < cmd.Size; x++ {
										time.Sleep(cmd.CurrentSpeed * 5)
									}
									time.Sleep(cmd.FadeTime / 4) // Fade on time.
									if !cmd.Inverted {
										for _, value := range fadeDown {
											R = int((float64(position.Color.R) / 100) * (float64(value) / 2.55))
											G = int((float64(position.Color.G) / 100) * (float64(value) / 2.55))
											B = int((float64(position.Color.B) / 100) * (float64(value) / 2.55))
											if !cmd.Hide {
												launchpad.LightLamp(mySequenceNumber, myFixtureNumber, R, G, B, cmd.Master, eventsForLauchpad)
											}
											MapFixtures(mySequenceNumber, dmxController, myFixtureNumber, R, G, B, 0, 0, 0, 0, fixtures, cmd.Blackout, cmd.Master, cmd.Master)
											time.Sleep(cmd.FadeTime / 4) // Fade down time.
										}
										time.Sleep(cmd.FadeTime / 4) // Fade off time.
									}
								}()
							}
						}
					}
				}
			}
		}
	}
}

// When want to light a DMX fixture we need for find it in our fuxture.yaml configuration file.
// This function maps the requested fixture into a DMX address.
func MapFixtures(mySequenceNumber int,
	dmxController *ft232.DMXController,
	displayFixture int, R int, G int, B int, Pan int, Tilt int, Shutter int, Gobo int,
	fixtures *Fixtures, blackout bool, brightness int, master int) {

	// We control the brightness of each color with the brightness value.
	// The overall fixture brightness is set from the master value.
	Red := (float64(R) / 100) * (float64(brightness) / 2.55)
	Green := (float64(G) / 100) * (float64(brightness) / 2.55)
	Blue := (float64(B) / 100) * (float64(brightness) / 2.55)

	for _, fixture := range fixtures.Fixtures {

		if fixture.Group-1 == mySequenceNumber {
			for channelNumber, channel := range fixture.Channels {
				// Scanner channels
				if strings.Contains(channel.Name, "Pan") {
					dmxController.SetChannel(fixture.Address+int16(channelNumber), byte(Pan))
				}
				if strings.Contains(channel.Name, "Tilt") {
					dmxController.SetChannel(fixture.Address+int16(channelNumber), byte(Tilt))
				}
				if strings.Contains(channel.Name, "Shutter") {
					dmxController.SetChannel(fixture.Address+int16(channelNumber), byte(Shutter))
				}
				if strings.Contains(channel.Name, "Gobo") {
					dmxController.SetChannel(fixture.Address+int16(channelNumber), byte(Gobo))
				}
				// Static value.
				if strings.Contains(channel.Name, "Static") {
					dmxController.SetChannel(fixture.Address+int16(channelNumber), byte(channel.Value))
				}
				// Fixture channels.
				if strings.Contains(channel.Name, "Red"+strconv.Itoa(displayFixture+1)) {
					dmxController.SetChannel(fixture.Address+int16(channelNumber), byte(int(Red)))
				}
				if strings.Contains(channel.Name, "Green"+strconv.Itoa(displayFixture+1)) {
					dmxController.SetChannel(fixture.Address+int16(channelNumber), byte(int(Green)))
				}
				if strings.Contains(channel.Name, "Blue"+strconv.Itoa(displayFixture+1)) {
					dmxController.SetChannel(fixture.Address+int16(channelNumber), byte(int(Blue)))
				}
				if strings.Contains(channel.Name, "Master") || strings.Contains(channel.Name, "Dimmer") {
					if blackout {
						dmxController.SetChannel(fixture.Address+int16(channelNumber), byte(0))
					} else {
						dmxController.SetChannel(fixture.Address+int16(channelNumber), byte(master))
					}
				}
			}

		}
	}
}

func MapSwitchFixture(mySequenceNumber int,
	dmxController *ft232.DMXController,
	displayFixture int, selectedSwitch int,
	fixtures *Fixtures, blackout bool, brightness int, master int) {

	for _, fixture := range fixtures.Fixtures {
		if fixture.Group-1 == mySequenceNumber {
			for _, swiTch := range fixture.Switches {
				for stateNumber, state := range swiTch.States {
					if stateNumber == selectedSwitch {
						for _, value := range state.Values {
							if blackout {
								dmxController.SetChannel(fixture.Address+int16(value.Channel), byte(0))
							} else {
								dmxController.SetChannel(fixture.Address+int16(value.Channel), byte(value.Setting))
							}
						}
					}
				}
			}
		}
	}
}
