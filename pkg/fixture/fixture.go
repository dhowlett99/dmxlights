package fixture

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/launchpad"
	"github.com/go-yaml/yaml"
	"github.com/oliread/usbdmx/ft232"
)

const debug = false

type Fixtures struct {
	Fixtures []Fixture `yaml:"fixtures"`
}

type Color struct {
	R int `yaml:"red"`
	G int `yaml:"green"`
	B int `yaml:"blue"`
}

type Value struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Channel     int16  `yaml:"channel"`
	Setting     int16  `yaml:"setting"`
}

type State struct {
	Name        string  `yaml:"name"`
	Values      []Value `yaml:"values"`
	ButtonColor Color   `yaml:"buttoncolor"`
	Master      int     `yaml:"master"`
}

type Switch struct {
	Name        string  `yaml:"name"`
	Number      int     `yaml:"number"`
	Description string  `yaml:"description"`
	States      []State `yaml:"states"`
}

type Fixture struct {
	Name        string    `yaml:"name"`
	Number      int       `yaml:"number"`
	Description string    `yaml:"description"`
	Type        string    `yaml:"type"`
	Group       int       `yaml:"group"`
	Address     int16     `yaml:"address"`
	Channels    []Channel `yaml:"channels"`
	Switches    []Switch  `yaml:"switches"`
}

type Setting struct {
	Name    string `yaml:"name"`
	Number  int    `yaml:"number"`
	Setting int    `yaml:"setting"`
}

type Channel struct {
	Number   int16     `yaml:"number"`
	Name     string    `yaml:"name"`
	Value    int16     `yaml:"value"`
	Comment  string    `yaml:"comment"`
	Settings []Setting `yaml:"settings"`
}

func LoadFixtures() (fixtures *Fixtures, err error) {
	filename := "fixtures.yaml"

	_, err = os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		return nil, errors.New("error: loading fixtures.yaml file: " + err.Error())
	}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.New("error: reading fixtures.yaml file: " + err.Error())
	}

	fixtures = &Fixtures{}
	err = yaml.Unmarshal(data, fixtures)
	if err != nil {
		return nil, errors.New("error: unmarshalling fixtures.yaml file: " + err.Error())
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

	// Calculate fade curve values.
	fadeUp := getFade(127.5, false)
	fadeDown := getFade(127.5, true)

	for {
		select {
		case cmd = <-channels[myFixtureNumber]:
			if cmd.Tick {
				// positions can have many fixtures play at the same time.
				positions := cmd.Positions[cmd.CurrentPosition]
				for _, position := range positions {

					if debug {
						// Some debug to print the positions.
						if sequence.Type == "scanner" && position.Fixture == 0 {
							fmt.Printf("===========Seq Number %d ==============\n", cmd.SequenceNumber)
							fmt.Printf("StartPosition %+v\n", position.StartPosition)
							fmt.Printf("Color %+v\n", position.Color)
							fmt.Printf("Fixture %+v\n", position.Fixture)
							fmt.Printf("Gobo %+v\n", position.Gobo)
							fmt.Printf("Pan %+v\n", position.Pan)
							fmt.Printf("Tilt %+v\n", position.Tilt)
							fmt.Printf("Shutter %+v\n", position.Shutter)
						}
					}

					// If this fixture is disabled then shut the shutter off.
					if cmd.CurrentPosition == position.StartPosition &&
						cmd.Type == "scanner" &&
						cmd.FixtureDisabled[myFixtureNumber] {
						MapFixtures(mySequenceNumber, dmxController, myFixtureNumber, 0, 0, 0, 0, 0, 0, 0, fixtures, cmd.Blackout, 0, 0)
						continue
					}

					// Short ciruit the soft fade if we are a scanner.
					if cmd.CurrentPosition == position.StartPosition && !cmd.FixtureDisabled[myFixtureNumber] {
						if cmd.Type == "scanner" {

							if position.Fixture == myFixtureNumber {
								MapFixtures(mySequenceNumber, dmxController, myFixtureNumber, position.Color.R, position.Color.G, position.Color.B, position.Pan, position.Tilt, position.Shutter, cmd.SelectedGobo, fixtures, cmd.Blackout, cmd.Master, cmd.Master)
								if !cmd.Hide {
									if cmd.ScannerChase {
										launchpad.LightLamp(mySequenceNumber, myFixtureNumber, 255, 255, 255, position.Shutter, eventsForLauchpad)
									} else {
										launchpad.LightLamp(mySequenceNumber, myFixtureNumber, position.Color.R, position.Color.G, position.Color.B, cmd.Master, eventsForLauchpad)
									}
								}
							}
							continue
						}

						// Short ciruit the soft fade if we in flood mode.
						if cmd.Flood {
							continue
						}

						// Now process RGB fixtures.
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

// When want to light a DMX fixture we need for find it in our fuxture.yaml configuration file.
// This function maps the requested fixture into a DMX address.
func MapFixtures(mySequenceNumber int,
	dmxController *ft232.DMXController,
	displayFixture int, R int, G int, B int, Pan int, Tilt int, Shutter int, selectedGobo int,
	fixtures *Fixtures, blackout bool, brightness int, master int) {

	// We control the brightness of each color with the brightness value.
	// The overall fixture brightness is set from the master value.
	Red := (float64(R) / 100) * (float64(brightness) / 2.55)
	Green := (float64(G) / 100) * (float64(brightness) / 2.55)
	Blue := (float64(B) / 100) * (float64(brightness) / 2.55)

	for _, fixture := range fixtures.Fixtures {

		if fixture.Group-1 == mySequenceNumber {
			for channelNumber, channel := range fixture.Channels {

				if fixture.Number-1 == displayFixture {
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
						for _, setting := range channel.Settings {
							if setting.Number == selectedGobo {
								dmxController.SetChannel(fixture.Address+int16(channelNumber), byte(setting.Setting))
							}
						}
					}
					if strings.Contains(channel.Name, "Master") || strings.Contains(channel.Name, "Dimmer") {
						if blackout {
							dmxController.SetChannel(fixture.Address+int16(channelNumber), byte(0))
						} else {
							dmxController.SetChannel(fixture.Address+int16(channelNumber), byte(master))
						}
					}
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

			}
		}
	}
}

func MapSwitchFixture(mySequenceNumber int,
	dmxController *ft232.DMXController,
	switchNumber int, currentState int,
	fixtures *Fixtures, blackout bool, brightness int, master int) {

	// Step through the fixture config file looking for the group that matches mysequence number.
	for _, fixture := range fixtures.Fixtures {
		if fixture.Group-1 == mySequenceNumber {
			for _, swiTch := range fixture.Switches {
				if switchNumber+1 == swiTch.Number {
					for stateNumber, state := range swiTch.States {
						if stateNumber == currentState {

							for _, value := range state.Values {
								if blackout {
									dmxController.SetChannel(fixture.Address+int16(value.Channel), byte(0))
								} else {
									// This should be controlled by the master brightness
									if strings.Contains(value.Name, "master") || strings.Contains(value.Name, "dimmer") {
										howBright := int((float64(value.Setting) / 100) * (float64(brightness) / 2.55))
										if strings.Contains(value.Name, "reverse") || strings.Contains(value.Name, "invert") {
											dmxController.SetChannel(fixture.Address+int16(value.Channel), byte(reverse_dmx(howBright)))
										} else {
											dmxController.SetChannel(fixture.Address+int16(value.Channel), byte(howBright))
										}
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
	}
}

func reverse_dmx(n int) int {

	in := make(map[int]int, 255)
	var y = 255

	for x := 0; x <= 255; x++ {

		in[x] = y
		y--
	}
	return in[n]
}

func HowManyGobos(sequenceNumber int, fixturesConfig *Fixtures) (numberScanners int, gobos []common.Gobo) {
	if debug {
		fmt.Printf("HowManyGobos\n")
	}
	gobos = []common.Gobo{}
	for _, f := range fixturesConfig.Fixtures {
		if debug {
			fmt.Printf("Fixture Name:%s\n", f.Name)
		}
		if f.Type == "scanner" {
			numberScanners++
			if debug {
				fmt.Printf("Sequence: %d - Scanner Name: %s Description: %s\n", sequenceNumber, f.Name, f.Description)
			}
			for _, channel := range f.Channels {
				if channel.Name == "Gobo" {
					newGobo := common.Gobo{}
					for _, setting := range channel.Settings {
						newGobo.Name = setting.Name
						newGobo.Number = setting.Number
						newGobo.Setting = setting.Setting
						gobos = append(gobos, newGobo)
						if debug {
							fmt.Printf("\tGobo: %s Setting: %d\n", setting.Name, setting.Setting)
						}
					}
				}
			}
		}
	}
	return numberScanners, gobos
}

func getFade(size float64, direction bool) []int {

	out := []int{}

	var theta float64
	var x float64
	if direction {
		for x = 0; x <= 180; x += 15 {

			theta = (x - 90) * math.Pi / 180

			x := int(-size*math.Sin(theta) + size)
			out = append(out, x)
		}
	} else {
		for x = 180; x >= 0; x -= 15 {
			theta = (x - 90) * math.Pi / 180

			x := int(-size*math.Sin(theta) + size)
			out = append(out, x)
		}
	}
	return out
}
