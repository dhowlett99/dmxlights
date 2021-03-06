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
	Label       string  `yaml:"label"`
	Values      []Value `yaml:"values"`
	ButtonColor Color   `yaml:"buttoncolor"`
	Master      int     `yaml:"master"`
}

type Switch struct {
	Name        string  `yaml:"name"`
	Label       string  `yaml:"label"`
	Number      int     `yaml:"number"`
	Description string  `yaml:"description"`
	States      []State `yaml:"states"`
}

type Fixture struct {
	Name        string    `yaml:"name"`
	Label       string    `yaml:"label"`
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
	Label   string `yaml:"label"`
	Number  int    `yaml:"number"`
	Setting int    `yaml:"setting"`
}

type Channel struct {
	Number     int16     `yaml:"number"`
	Name       string    `yaml:"name"`
	Value      int16     `yaml:"value"`
	MaxDegrees *int      `yaml:"maxdegrees,omitempty"`
	Offset     *int      `yaml:"offset,omitempty"` // Offset allows you to position the fixture.
	Comment    string    `yaml:"comment"`
	Settings   []Setting `yaml:"settings"`
}

// LoadFixtures opens the fixtures config file and returns a pointer to the fixtures.
// or an error.
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
	guiButtons chan common.ALight,
	dmxController *ft232.DMXController,
	fixtures *Fixtures,
	fixtureStopChannels []chan bool) {

	cmd := common.FixtureCommand{}

	// Calculate fade curve values.
	fadeUp := getFade(127.5, false)
	fadeDown := getFade(127.5, true)

	for {
		cmd = <-channels[myFixtureNumber]

		if cmd.Tick {

			// If we're a RGB fixture implement the flood and static features.
			if cmd.Type == "rgb" {
				if cmd.StartFlood {
					MapFixtures(cmd.SequenceNumber, dmxController, myFixtureNumber, 255, 255, 255, 0, 0, 0, 0, nil, fixtures, sequence.Blackout, sequence.Master, sequence.Master)
					common.LightLamp(common.ALight{X: myFixtureNumber, Y: sequence.Number, Red: 255, Green: 255, Blue: 255, Brightness: 255}, eventsForLauchpad, guiButtons)
					continue
				}
				if cmd.StopFlood {
					MapFixtures(cmd.SequenceNumber, dmxController, myFixtureNumber, 0, 0, 0, 0, 0, 0, 0, nil, fixtures, sequence.Blackout, sequence.Master, sequence.Master)
					common.LightLamp(common.ALight{X: myFixtureNumber, Y: sequence.Number, Red: 0, Green: 0, Blue: 0, Brightness: 0}, eventsForLauchpad, guiButtons)
					continue
				}
				if cmd.Static {
					sequence := common.Sequence{}
					sequence.Type = cmd.Type
					sequence.Number = cmd.SequenceNumber
					sequence.Master = cmd.Master
					sequence.Blackout = cmd.Blackout
					sequence.Hide = cmd.Hide
					sequence.StaticColors = cmd.StaticColors
					sequence.Static = cmd.Static
					lightStaticFixture(sequence, myFixtureNumber, dmxController, eventsForLauchpad, guiButtons, fixtures, true)
					continue
				}
			}

			// Positions can have many fixtures play at the same time.
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

				cmd.FixtureDisabledMutex.RLock()
				disableThis := cmd.FixtureDisabled[myFixtureNumber]
				cmd.FixtureDisabledMutex.RUnlock()

				cmd.DisableOnceMutex.RLock()
				disableOnce := cmd.DisableOnce[myFixtureNumber]
				cmd.DisableOnceMutex.RUnlock()

				// If this fixture is disabled then shut the shutter off.
				if cmd.CurrentPosition == position.StartPosition &&
					cmd.Type == "scanner" &&
					disableOnce &&
					disableThis {

					MapFixtures(mySequenceNumber, dmxController, myFixtureNumber, 0, 0, 0, 0, 0, 0, 0, nil, fixtures, cmd.Blackout, 0, 0)

					cmd.DisableOnceMutex.RLock()
					cmd.DisableOnce[myFixtureNumber] = false
					cmd.DisableOnceMutex.RUnlock()

					continue
				}

				cmd.FixtureDisabledMutex.RLock()
				disableThis = cmd.FixtureDisabled[myFixtureNumber]
				cmd.FixtureDisabledMutex.RUnlock()

				if cmd.CurrentPosition == position.StartPosition && !disableThis {

					// S C A N N E R - Short ciruit the soft fade if we are a scanner.
					if cmd.Type == "scanner" {
						if position.Fixture == myFixtureNumber {

							MapFixtures(mySequenceNumber, dmxController, myFixtureNumber, position.Color.R, position.Color.G, position.Color.B, position.Pan, position.Tilt,
								position.Shutter, cmd.SelectedGobo, cmd.ScannerColor, fixtures, cmd.Blackout, cmd.Master, cmd.Master)

							if !cmd.Hide {
								if cmd.ScannerChase {
									// We are chase mode, we want the buttons to be the color selected for this scanner.

									// Remember that fixtures in the real world start with 1 not 0.
									realFixture := position.Fixture + 1
									// Find the color that has been selected for this fixture.
									// selected color is an index into the scanner colors selected.
									selectedColor := cmd.ScannerColor[position.Fixture]

									// Do we have a set of available colors for this fixture.
									_, ok := cmd.AvailableScannerColors[realFixture]
									if ok {
										availableColors := cmd.AvailableScannerColors[realFixture]
										red := availableColors[selectedColor].Color.R
										green := availableColors[selectedColor].Color.G
										blue := availableColors[selectedColor].Color.B
										common.LightLamp(common.ALight{X: myFixtureNumber, Y: sequence.Number, Red: red, Green: green, Blue: blue, Brightness: position.Shutter}, eventsForLauchpad, guiButtons)
										common.LabelButton(myFixtureNumber, sequence.Number, "", guiButtons)
									} else {
										// No color selected or available, use white.
										common.LightLamp(common.ALight{X: myFixtureNumber, Y: sequence.Number, Red: 255, Green: 255, Blue: 255, Brightness: position.Shutter}, eventsForLauchpad, guiButtons)
										common.LabelButton(myFixtureNumber, sequence.Number, "", guiButtons)
									}
								} else {
									// We're not in chase mode so use the color generated in the patten generator.common.
									common.LightLamp(common.ALight{X: myFixtureNumber, Y: mySequenceNumber, Red: position.Color.R, Green: position.Color.G, Blue: position.Color.B, Brightness: cmd.Master}, eventsForLauchpad, guiButtons)
									common.LabelButton(myFixtureNumber, sequence.Number, "", guiButtons)
								}
							}
						}
						continue
					}

					// Now create a thread to process the RGB fixture itself.
					if position.Fixture == myFixtureNumber {
						go func(fixtureStopChannel chan bool) {
							var R int
							var G int
							var B int

							if cmd.Inverted {
								for _, value := range fadeDown {
									R = int((float64(position.Color.R) / 100) * (float64(value) / 2.55))
									G = int((float64(position.Color.G) / 100) * (float64(value) / 2.55))
									B = int((float64(position.Color.B) / 100) * (float64(value) / 2.55))
									if !cmd.Hide {
										common.LightLamp(common.ALight{X: myFixtureNumber, Y: mySequenceNumber, Red: R, Green: G, Blue: B, Brightness: cmd.Master}, eventsForLauchpad, guiButtons)
									}
									MapFixtures(mySequenceNumber, dmxController, myFixtureNumber, R, G, B, 0, 0, 0, 0, cmd.ScannerColor, fixtures, cmd.Blackout, cmd.Master, cmd.Master)
									// Fade down time.
									if listenAndWaitForStop(cmd.FadeTime/4, fixtureStopChannel) {
										turnOffFixtures(cmd, myFixtureNumber, mySequenceNumber, fixtures, dmxController, eventsForLauchpad, guiButtons)
										return
									}
								}
								// Fade off time.
								if listenAndWaitForStop(cmd.FadeTime/4, fixtureStopChannel) {
									turnOffFixtures(cmd, myFixtureNumber, mySequenceNumber, fixtures, dmxController, eventsForLauchpad, guiButtons)
									return
								}
							}
							for _, value := range fadeUp {
								R = int((float64(position.Color.R) / 100) * (float64(value) / 2.55))
								G = int((float64(position.Color.G) / 100) * (float64(value) / 2.55))
								B = int((float64(position.Color.B) / 100) * (float64(value) / 2.55))
								if !cmd.Hide {
									common.LightLamp(common.ALight{X: myFixtureNumber, Y: mySequenceNumber, Red: R, Green: G, Blue: B, Brightness: cmd.Master}, eventsForLauchpad, guiButtons)
								}
								// Now ask DMX to actually light the real fixture.
								MapFixtures(mySequenceNumber, dmxController, myFixtureNumber, R, G, B, 0, 0, 0, 0, cmd.ScannerColor, fixtures, cmd.Blackout, cmd.Master, cmd.Master)
								// Fade up time.
								if listenAndWaitForStop(cmd.FadeTime/4, fixtureStopChannel) {
									turnOffFixtures(cmd, myFixtureNumber, mySequenceNumber, fixtures, dmxController, eventsForLauchpad, guiButtons)
									return
								}
							}
							for x := 0; x < cmd.Size; x++ {
								// Time between fades.
								if listenAndWaitForStop(cmd.CurrentSpeed*5, fixtureStopChannel) {
									turnOffFixtures(cmd, myFixtureNumber, mySequenceNumber, fixtures, dmxController, eventsForLauchpad, guiButtons)
									return
								}
							}
							// Fade on time.
							if listenAndWaitForStop(cmd.FadeTime/4, fixtureStopChannel) {
								turnOffFixtures(cmd, myFixtureNumber, mySequenceNumber, fixtures, dmxController, eventsForLauchpad, guiButtons)
								return
							}
							if !cmd.Inverted {
								for _, value := range fadeDown {
									R = int((float64(position.Color.R) / 100) * (float64(value) / 2.55))
									G = int((float64(position.Color.G) / 100) * (float64(value) / 2.55))
									B = int((float64(position.Color.B) / 100) * (float64(value) / 2.55))
									if !cmd.Hide {
										common.LightLamp(common.ALight{X: myFixtureNumber, Y: mySequenceNumber, Red: R, Green: G, Blue: B, Brightness: cmd.Master}, eventsForLauchpad, guiButtons)
									}
									MapFixtures(mySequenceNumber, dmxController, myFixtureNumber, R, G, B, 0, 0, 0, 0, cmd.ScannerColor, fixtures, cmd.Blackout, cmd.Master, cmd.Master)
									// Fade down time.
									if listenAndWaitForStop(cmd.FadeTime/4, fixtureStopChannel) {
										turnOffFixtures(cmd, myFixtureNumber, mySequenceNumber, fixtures, dmxController, eventsForLauchpad, guiButtons)
										return
									}
								}
								// Fade off time.
								if listenAndWaitForStop(cmd.FadeTime/4, fixtureStopChannel) {
									turnOffFixtures(cmd, myFixtureNumber, mySequenceNumber, fixtures, dmxController, eventsForLauchpad, guiButtons)
									return
								}
							}
						}(fixtureStopChannels[myFixtureNumber])
					}
				}
			}
		}
	}
}

func MapFixturesColorOnly(sequence *common.Sequence, dmxController *ft232.DMXController, fixtures *Fixtures, scannerColor int) {
	for _, fixture := range fixtures.Fixtures {
		if fixture.Group-1 == sequence.Number {
			for channelNumber, channel := range fixture.Channels {
				if strings.Contains(channel.Name, "Color") {
					for _, setting := range channel.Settings {
						//fmt.Printf("Scanner Color  %d\n", scannerColor)
						//fmt.Printf("Setting  %+v\n", setting)
						if setting.Number-1 == scannerColor {
							dmxController.SetChannel(fixture.Address+int16(channelNumber), byte(setting.Setting))
						}
					}
				}
			}
		}
	}
}

func MapFixturesGoboOnly(sequence *common.Sequence, dmxController *ft232.DMXController, fixtures *Fixtures, selectedGobo int) {

	for _, fixture := range fixtures.Fixtures {
		if fixture.Group-1 == sequence.Number {
			for channelNumber, channel := range fixture.Channels {
				if strings.Contains(channel.Name, "Gobo") {
					for _, setting := range channel.Settings {
						if setting.Number == selectedGobo {
							dmxController.SetChannel(fixture.Address+int16(channelNumber), byte(setting.Setting))
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
	displayFixture int, R int, G int, B int, Pan int, Tilt int, Shutter int, selectedGobo int, scannerColor map[int]int,
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
						if channel.Offset != nil {
							dmxController.SetChannel(fixture.Address+int16(channelNumber), byte(limitDmxValue(channel.MaxDegrees, Pan+*channel.Offset)))
						} else {
							dmxController.SetChannel(fixture.Address+int16(channelNumber), byte(limitDmxValue(channel.MaxDegrees, Pan)))
						}
					}
					if strings.Contains(channel.Name, "Tilt") {
						if channel.Offset != nil {
							dmxController.SetChannel(fixture.Address+int16(channelNumber), byte(limitDmxValue(channel.MaxDegrees, Tilt+*channel.Offset)))
						}
						dmxController.SetChannel(fixture.Address+int16(channelNumber), byte(limitDmxValue(channel.MaxDegrees, Tilt)))
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
					if strings.Contains(channel.Name, "Color") {
						for colorNumber := range scannerColor {
							if colorNumber == displayFixture {
								for _, setting := range channel.Settings {
									if setting.Number-1 == scannerColor[displayFixture] {
										dmxController.SetChannel(fixture.Address+int16(channelNumber), byte(setting.Setting))
									}
								}
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

func lightStaticFixture(sequence common.Sequence, myFixtureNumber int, dmxController *ft232.DMXController, eventsForLauchpad chan common.ALight, guiButtons chan common.ALight, fixturesConfig *Fixtures, enabled bool) {

	lamp := sequence.StaticColors[myFixtureNumber]

	if sequence.Hide {
		if lamp.Flash {
			onColor := common.Color{R: lamp.Color.R, G: lamp.Color.G, B: lamp.Color.B}
			Black := common.Color{R: 0, G: 0, B: 0}
			common.FlashLight(sequence.Number, myFixtureNumber, onColor, Black, eventsForLauchpad, guiButtons)
		} else {
			common.LightLamp(common.ALight{X: myFixtureNumber, Y: sequence.Number, Red: lamp.Color.R, Green: lamp.Color.G, Blue: lamp.Color.B, Brightness: sequence.Master}, eventsForLauchpad, guiButtons)
		}
	}
	MapFixtures(sequence.Number, dmxController, myFixtureNumber, lamp.Color.R, lamp.Color.G, lamp.Color.B, 0, 0, 0, 0, nil, fixturesConfig, sequence.Blackout, sequence.Master, sequence.Master)

	// Only play once, we don't want to flood the DMX universe with
	// continual commands.
	sequence.PlayStaticOnce = false
}

// limitDmxValue - calculates the maximum DMX value based on the number of degrees the fixtire can achieve.
func limitDmxValue(MaxDegrees *int, Value int) int {

	in := float64(Value)

	// Limit the DMX value for Pan so the max degree we send is always less than or equal to 360 degrees.
	OriginalDMXValueRatio := float64(360) / float64(255)

	// If maxiumum number of degrees have been specified
	// then do the math so that 360 degrees are never exceeded.
	if MaxDegrees == nil {
		return int(in)
	}

	if *MaxDegrees < 360 { // If its less then 360 we can't limit.
		return int(in / OriginalDMXValueRatio)
	}

	DegreesRequired := math.Round(OriginalDMXValueRatio * float64(Value))

	var MaxDMX float64 = 255

	DegreesPerDMXClick := float64(*MaxDegrees) / MaxDMX

	NewDMXValue := int(math.Round(DegreesRequired / DegreesPerDMXClick))

	return NewDMXValue

}

// listenAndWaitForStop is used in the fixture fade loops.
// We make the sleeps interuptable so that fixtures can be stopped immediately
func listenAndWaitForStop(wait time.Duration, fixtureStopChannels chan bool) bool {
	select {
	case <-fixtureStopChannels:
		return true
	case <-time.After(wait):
	}
	return false
}

// turnOffFixtures is used to turn off a fixture when we stop a sequence.
func turnOffFixtures(cmd common.FixtureCommand, myFixtureNumber int, mySequenceNumber int, fixtures *Fixtures, dmxController *ft232.DMXController, eventsForLauchpad chan common.ALight, guiButtons chan common.ALight) {
	if !cmd.Hide {
		common.LightLamp(common.ALight{X: myFixtureNumber, Y: mySequenceNumber, Red: 0, Green: 0, Blue: 0, Brightness: 0}, eventsForLauchpad, guiButtons)
	}
	MapFixtures(mySequenceNumber, dmxController, myFixtureNumber, 0, 0, 0, 0, 0, 0, 0, cmd.ScannerColor, fixtures, cmd.Blackout, cmd.Master, cmd.Master)
}
