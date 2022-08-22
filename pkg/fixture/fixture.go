package fixture

import (
	"errors"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/go-yaml/yaml"
	"github.com/oliread/usbdmx/ft232"
)

//const debug = false

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
	Name           string    `yaml:"name"`
	Label          string    `yaml:"label"`
	Number         int       `yaml:"number"`
	Description    string    `yaml:"description"`
	Type           string    `yaml:"type"`
	Group          int       `yaml:"group"`
	Address        int16     `yaml:"address"`
	Channels       []Channel `yaml:"channels"`
	Switches       []Switch  `yaml:"switches"`
	NumberChannels int       `yaml:"number_channels"`
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
func FixtureReceiver(
	sequence common.Sequence,
	mySequenceNumber int,
	myFixtureNumber int,
	fixtureControlChannel chan common.FixtureCommand,
	fixtureStepChannel chan common.NewFixtureCommand,
	fixtureStopChannel chan bool,
	eventsForLauchpad chan common.ALight,
	guiButtons chan common.ALight,
	dmxController *ft232.DMXController,
	fixtures *Fixtures) {

	// Calculate fade curve values.
	fadeUpValues := getFadeValues(127.5, false)
	fadeDownValues := getFadeValues(127.5, true)

	// Outer loop wait for configuration.
	for {

		// Wait for first step
		cmd := <-fixtureStepChannel

		// Positions can have many fixtures play at the same time.
		positions := cmd.RGBPositions[cmd.Step]

	start:
		for _, position := range positions {

			if position.StartPosition == cmd.Step {

				// Now create a thread to process the RGB fixture itself.
				if position.Fixture == myFixtureNumber {

					if cmd.Invert {
						// Fade down.
						for _, value := range fadeDownValues {
							R := int((float64(position.Color.R) / 100) * (float64(value) / 2.55))
							G := int((float64(position.Color.G) / 100) * (float64(value) / 2.55))
							B := int((float64(position.Color.B) / 100) * (float64(value) / 2.55))
							if !cmd.Hide {
								common.LightLamp(common.ALight{X: myFixtureNumber, Y: mySequenceNumber, Red: R, Green: G, Blue: B, Brightness: cmd.Master}, eventsForLauchpad, guiButtons)
							}
							MapFixtures(mySequenceNumber, dmxController, myFixtureNumber, R, G, B, 0, 0, 0, 0, cmd.ScannerColor, fixtures, cmd.Blackout, cmd.Master, cmd.Master, cmd.StrobeSpeed)
							// Fade down time.
							if listenAndWaitForStop(fixtureStepChannel, fixtureStopChannel) {
								turnOffFixtures(cmd, myFixtureNumber, mySequenceNumber, fixtures, dmxController, eventsForLauchpad, guiButtons)
								break start
							}
						}

						// Fade on or off time.
						for x := 0; x < cmd.RGBSize*10; x++ {
							if listenAndWaitForStop(fixtureStepChannel, fixtureStopChannel) {
								turnOffFixtures(cmd, myFixtureNumber, mySequenceNumber, fixtures, dmxController, eventsForLauchpad, guiButtons)
								break start
							}
						}
					}

					// Fade up.
					for _, value := range fadeUpValues {
						R := int((float64(position.Color.R) / 100) * (float64(value) / 2.55))
						G := int((float64(position.Color.G) / 100) * (float64(value) / 2.55))
						B := int((float64(position.Color.B) / 100) * (float64(value) / 2.55))
						if !cmd.Hide {
							common.LightLamp(common.ALight{X: myFixtureNumber, Y: mySequenceNumber, Red: R, Green: G, Blue: B, Brightness: cmd.Master}, eventsForLauchpad, guiButtons)
						}
						// Now ask DMX to actually light the real fixture.
						MapFixtures(mySequenceNumber, dmxController, myFixtureNumber, R, G, B, 0, 0, 0, 0, cmd.ScannerColor, fixtures, cmd.Blackout, cmd.Master, cmd.Master, cmd.StrobeSpeed)
						// Fade up time.
						if listenAndWaitForStop(fixtureStepChannel, fixtureStopChannel) {
							turnOffFixtures(cmd, myFixtureNumber, mySequenceNumber, fixtures, dmxController, eventsForLauchpad, guiButtons)
							break start
						}
					}

					if !cmd.Invert {
						// Fade on or off time.
						for x := 0; x < cmd.RGBSize*10; x++ {
							if listenAndWaitForStop(fixtureStepChannel, fixtureStopChannel) {
								turnOffFixtures(cmd, myFixtureNumber, mySequenceNumber, fixtures, dmxController, eventsForLauchpad, guiButtons)
								break start
							}
						}

						// Fade down.
						for _, value := range fadeDownValues {
							R := int((float64(position.Color.R) / 100) * (float64(value) / 2.55))
							G := int((float64(position.Color.G) / 100) * (float64(value) / 2.55))
							B := int((float64(position.Color.B) / 100) * (float64(value) / 2.55))
							if !cmd.Hide {
								common.LightLamp(common.ALight{X: myFixtureNumber, Y: mySequenceNumber, Red: R, Green: G, Blue: B, Brightness: cmd.Master}, eventsForLauchpad, guiButtons)
							}
							MapFixtures(mySequenceNumber, dmxController, myFixtureNumber, R, G, B, 0, 0, 0, 0, cmd.ScannerColor, fixtures, cmd.Blackout, cmd.Master, cmd.Master, cmd.StrobeSpeed)
							// Fade down time.
							if listenAndWaitForStop(fixtureStepChannel, fixtureStopChannel) {
								turnOffFixtures(cmd, myFixtureNumber, mySequenceNumber, fixtures, dmxController, eventsForLauchpad, guiButtons)
								break start
							}
						}
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
	fixtures *Fixtures, blackout bool, brightness int, master int, strobe int) {

	// We control the brightness of each color with the brightness value.
	// The overall fixture brightness is set from the master value.
	Red := (float64(R) / 100) * (float64(brightness) / 2.55)
	Green := (float64(G) / 100) * (float64(brightness) / 2.55)
	Blue := (float64(B) / 100) * (float64(brightness) / 2.55)

	for _, fixture := range fixtures.Fixtures {
		if fixture.Group-1 == mySequenceNumber {
			for channelNumber, channel := range fixture.Channels {
				if fixture.Number == displayFixture+1 {
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
					if strings.Contains(channel.Name, "Strobe") {
						if blackout {
							dmxController.SetChannel(fixture.Address+int16(channelNumber), byte(0))
						} else {
							dmxController.SetChannel(fixture.Address+int16(channelNumber), byte(strobe))
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

func getFadeValues(size float64, direction bool) []int {

	out := []int{}

	var theta float64
	var x float64
	if direction {
		for x = 0; x <= 180; x += 15 {

			theta = (x - 90) * math.Pi / 180

			x := int(-size*math.Sin(theta) + size)
			//if x > 10 {
			out = append(out, x)
			//}
		}
	} else {
		for x = 180; x >= 0; x -= 15 {
			theta = (x - 90) * math.Pi / 180

			x := int(-size*math.Sin(theta) + size)
			//if x > 10 {
			out = append(out, x)
			//}

		}
	}
	return out
}

// func lightStaticFixture(sequence common.Sequence, myFixtureNumber int, dmxController *ft232.DMXController, eventsForLauchpad chan common.ALight, guiButtons chan common.ALight, fixturesConfig *Fixtures, enabled bool) {

// 	lamp := sequence.StaticColors[myFixtureNumber]

// 	if sequence.Hide {
// 		if lamp.Flash {
// 			onColor := common.Color{R: lamp.Color.R, G: lamp.Color.G, B: lamp.Color.B}
// 			Black := common.Color{R: 0, G: 0, B: 0}
// 			common.FlashLight(sequence.Number, myFixtureNumber, onColor, Black, eventsForLauchpad, guiButtons)
// 		} else {
// 			common.LightLamp(common.ALight{X: myFixtureNumber, Y: sequence.Number, Red: lamp.Color.R, Green: lamp.Color.G, Blue: lamp.Color.B, Brightness: sequence.Master}, eventsForLauchpad, guiButtons)
// 		}
// 	}
// 	MapFixtures(sequence.Number, dmxController, myFixtureNumber, lamp.Color.R, lamp.Color.G, lamp.Color.B, 0, 0, 0, 0, nil, fixturesConfig, sequence.Blackout, sequence.Master, sequence.Master, sequence.StrobeSpeed)

// 	// Only play once, we don't want to flood the DMX universe with
// 	// continual commands.
// 	sequence.PlayStaticOnce = false
// }

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
func listenAndWaitForStop(stepChannel chan common.NewFixtureCommand, stopChannel chan bool) bool {

	select {
	case <-stopChannel:
		//fmt.Printf("got stop\n")
		return true
	case <-stepChannel:
		// fmt.Printf("got step\n")
		break
		// case <-time.After(2 * time.Second):
		// 	fmt.Printf("wait step timeout\n")
		// 	break
	}
	return false
}

// turnOffFixtures is used to turn off a fixture when we stop a sequence.
func turnOffFixtures(cmd common.NewFixtureCommand, myFixtureNumber int, mySequenceNumber int, fixtures *Fixtures, dmxController *ft232.DMXController, eventsForLauchpad chan common.ALight, guiButtons chan common.ALight) {
	if !cmd.Hide {
		common.LightLamp(common.ALight{X: myFixtureNumber, Y: mySequenceNumber, Red: 0, Green: 0, Blue: 0, Brightness: 0}, eventsForLauchpad, guiButtons)
	}
	MapFixtures(mySequenceNumber, dmxController, myFixtureNumber, 0, 0, 0, 0, 0, 0, 0, cmd.ScannerColor, fixtures, cmd.Blackout, cmd.Master, cmd.Master, cmd.StrobeSpeed)
}
