// Copyright (C) 2022 dhowlett99.
// This is the dmxlights fixture editor it is attached to a fixture and
// describes the fixtures properties which is then saved in the fixtures.yaml
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
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/pattern"
	"github.com/dhowlett99/dmxlights/pkg/position"
	"github.com/dhowlett99/dmxlights/pkg/sound"
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
	W int `yaml:"white"`
}

type Value struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Channel     string `yaml:"channel"`
	Setting     string `yaml:"setting"`
}

type State struct {
	Name        string   `yaml:"name"`
	Number      int16    `yaml:"number"`
	Label       string   `yaml:"label"`
	Values      []Value  `yaml:"values,omitempty"`
	ButtonColor Color    `yaml:"buttoncolor"`
	Master      int      `yaml:"master"`
	Actions     []Action `yaml:"actions,omitempty"`
	Flash       bool     `yaml:"flash"`
}

type Action struct {
	Name    string `yaml:"name"`
	Number  int
	Colors  []string `yaml:"colors"`
	Mode    string   `yaml:"mode"`
	Fade    string   `yaml:"fade"`
	Size    string   `yaml:"size"`
	Speed   string   `yaml:"speed"`
	Rotate  string   `yaml:"rotate"`
	Music   string   `yaml:"music"`
	Program string   `yaml:"program"`
	Strobe  string   `yaml:"strobe"`
}

type ActionConfig struct {
	Name         string
	Colors       []common.Color
	Fade         int
	Size         int
	Speed        time.Duration
	TriggerState bool
	RotateSpeed  int
	Rotatable    bool
	Program      int
	Music        int
	MusicTrigger bool
	Strobe       int
}

type Fixture struct {
	Name           string    `yaml:"name"`
	Label          string    `yaml:"label,omitempty"`
	Number         int       `yaml:"number"`
	Description    string    `yaml:"description"`
	Type           string    `yaml:"type"`
	Group          int       `yaml:"group"`
	Address        int16     `yaml:"address"`
	Channels       []Channel `yaml:"channels"`
	States         []State   `yaml:"states,omitempty"`
	NumberChannels int       `yaml:"use_channels,omitempty"`
	UseFixture     string    `yaml:"use_fixture,omitempty"`
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
	Value      int16     `yaml:"value,omitempty"`
	MaxDegrees *int      `yaml:"maxdegrees,omitempty"`
	Offset     *int      `yaml:"offset,omitempty"` // Offset allows you to position the fixture.
	Comment    string    `yaml:"comment,omitempty"`
	Settings   []Setting `yaml:"settings,omitempty"`
}

// LoadFixtures opens the fixtures config file.
// Returns a pointer to the fixtures config.
// Returns an error.
func LoadFixtures(filename string) (fixtures *Fixtures, err error) {

	// Open the fixtures.yaml file.
	_, err = os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		return nil, errors.New("error: loading fixtures.yaml file: " + err.Error())
	}

	// Reads the fixtures.yaml file.
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.New("error: reading fixtures.yaml file: " + err.Error())
	}

	// Unmarshals the fixtures.yaml file into a data struct
	fixtures = &Fixtures{}
	err = yaml.Unmarshal(data, fixtures)
	if err != nil {
		return nil, errors.New("error: unmarshalling fixtures.yaml file: " + err.Error())
	}
	return fixtures, nil
}

// SaveFixtures - saves a complete list of fixtures to fixtures.yaml
// Returns an error.
func SaveFixtures(filename string, fixtures *Fixtures) error {

	// Marshal the fixtures data into a yaml data structure.
	data, err := yaml.Marshal(fixtures)
	if err != nil {
		return errors.New("error: marshalling fixtures.yaml file: " + err.Error())
	}

	// Open the fixtures.yaml file.
	_, err = os.Create(filename)
	if err != nil {
		return errors.New("error: opening fixtures.yaml file: " + err.Error())
	}

	// Write the fixtures.yaml file.
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return errors.New("error: writing fixtures.yaml file: " + err.Error())
	}

	// Fixtures file saved, no errors.
	return nil
}

// GetFixureDetails - find a fixture in the fixtures config.
// Returns details of the fixture.
// Returns an error.
func GetFixureDetails(group int, number int, fixtures *Fixtures) (Fixture, error) {
	// scan the fixtures structure for the selected fixture.
	if debug {
		fmt.Printf("Looking for Fixture Group %d, Number %d\n", group, number)
	}

	for _, fixture := range fixtures.Fixtures {
		if debug {
			fmt.Printf("Fixture Group %d, Number %d\n", fixture.Group, fixture.Number)
		}
		if fixture.Group == group && fixture.Number == number+1 {
			return fixture, nil
		}
	}
	return Fixture{}, fmt.Errorf("error: fixture not found")
}

// EditFixture - allows you to change the fixture details for the selected fixture.
// Returns a complete list of fixtures.
// Returns an error.
func EditFixture(groupNumber int, fixtureNumber int, newFixture Fixture, fixtures *Fixtures) (*Fixtures, error) {
	// scan the fixtures structure for the selected fixture.
	for _, fixture := range fixtures.Fixtures {
		if fixture.Group == groupNumber && fixture.Number == fixtureNumber {
			fixture = newFixture
			return fixtures, nil
		}
	}
	return fixtures, fmt.Errorf("error: fixture not found")
}

// FixtureReceivers are created by the sequence and are used to receive step instructions.
// Each FixtureReceiver knows which step they belong too and when triggered they start a fade up
// and fade down events which get sent to the launchpad lamps and the DMX fixtures.
func FixtureReceiver(
	sequence common.Sequence,
	mySequenceNumber int,
	myFixtureNumber int,
	fixtureStepChannel chan common.FixtureCommand,
	eventsForLauchpad chan common.ALight,
	guiButtons chan common.ALight,
	dmxController *ft232.DMXController,
	fixtures *Fixtures,
	dmxInterfacePresent bool) {

	// Outer loop wait for configuration.
	for {

		// Wait for first step
		cmd := <-fixtureStepChannel

		// Propogate the strobe speed.
		sequence.StrobeSpeed = cmd.StrobeSpeed

		// If we're a RGB fixture implement the flood and static features.
		if cmd.Type == "rgb" {
			if cmd.Clear {
				turnOffFixtures(cmd, myFixtureNumber, mySequenceNumber, fixtures, dmxController, eventsForLauchpad, guiButtons, dmxInterfacePresent)
				continue
			}

			if cmd.StartFlood {
				var lamp common.Color
				if cmd.RGBStatic {
					lamp = cmd.RGBStaticColors[myFixtureNumber].Color
				} else {
					lamp = common.Color{R: 255, G: 255, B: 255}
				}
				MapFixtures(cmd.SequenceNumber, dmxController, myFixtureNumber, lamp.R, lamp.G, lamp.B, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, nil, fixtures, sequence.Blackout, sequence.Master, sequence.Master, sequence.StrobeSpeed, dmxInterfacePresent)
				common.LightLamp(common.ALight{X: myFixtureNumber, Y: sequence.Number, Red: lamp.R, Green: lamp.G, Blue: lamp.B, Brightness: 255}, eventsForLauchpad, guiButtons)
				continue
			}
			if cmd.StopFlood {
				MapFixtures(cmd.SequenceNumber, dmxController, myFixtureNumber, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, nil, fixtures, sequence.Blackout, sequence.Master, sequence.Master, sequence.StrobeSpeed, dmxInterfacePresent)
				common.LightLamp(common.ALight{X: myFixtureNumber, Y: sequence.Number, Red: 0, Green: 0, Blue: 0, Brightness: 0}, eventsForLauchpad, guiButtons)
				continue
			}
			if cmd.RGBStatic {
				sequence := common.Sequence{}
				sequence.Type = cmd.Type
				sequence.Number = cmd.SequenceNumber
				sequence.Master = cmd.Master
				sequence.Blackout = cmd.Blackout
				sequence.Hide = cmd.Hide
				sequence.StaticColors = cmd.RGBStaticColors
				sequence.Static = cmd.RGBStatic
				sequence.StrobeSpeed = cmd.StrobeSpeed
				lightStaticFixture(sequence, myFixtureNumber, dmxController, eventsForLauchpad, guiButtons, fixtures, true, dmxInterfacePresent)
				continue
			}
			// Play out fixture to DMX channels.
			fixture := cmd.RGBPosition.Fixtures[myFixtureNumber]
			for _, color := range fixture.Colors {
				red := color.R
				green := color.G
				blue := color.B
				white := color.W

				if !cmd.Hide {
					common.LightLamp(common.ALight{X: myFixtureNumber, Y: mySequenceNumber, Red: red, Green: green, Blue: blue, Brightness: cmd.Master}, eventsForLauchpad, guiButtons)
				}
				MapFixtures(mySequenceNumber, dmxController, myFixtureNumber, red, green, blue, white, 0, 0, 0, 0, 0, 0, 0, 0, 0, cmd.ScannerColor, fixtures, cmd.Blackout, cmd.Master, cmd.Master, cmd.StrobeSpeed, dmxInterfacePresent)
			}
		}

		if cmd.Type == "scanner" {

			// Turn off the scanners in flood mode.
			if cmd.StartFlood {
				turnOnFixtures(cmd, myFixtureNumber, mySequenceNumber, fixtures, dmxController, eventsForLauchpad, guiButtons, dmxInterfacePresent)
				continue
			}

			// find the fixture
			fixture := cmd.ScannerPosition.Fixtures[myFixtureNumber]

			sequence.ScannerStateMutex.RLock()
			enabled := cmd.ScannerState[myFixtureNumber].Enabled
			sequence.ScannerStateMutex.RUnlock()

			sequence.DisableOnceMutex.RLock()
			disableOnce := cmd.ScannerDisableOnce[myFixtureNumber]
			sequence.DisableOnceMutex.RUnlock()

			// If this fixture is disabled then shut the shutter off.
			if disableOnce && !enabled {
				MapFixtures(mySequenceNumber, dmxController, myFixtureNumber, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, nil, fixtures, cmd.Blackout, 0, 0, 0, dmxInterfacePresent)
				// Locking for write.
				sequence.DisableOnceMutex.Lock()
				sequence.DisableOnce[myFixtureNumber] = false
				sequence.DisableOnceMutex.Unlock()
				continue
			}

			if enabled {

				// If enables activate the physical scanner.
				MapFixtures(mySequenceNumber, dmxController, myFixtureNumber, fixture.ScannerColor.R, fixture.ScannerColor.G, fixture.ScannerColor.B, fixture.ScannerColor.W, fixture.ScannerColor.A, fixture.ScannerColor.UV, fixture.Pan, fixture.Tilt,
					fixture.Shutter, cmd.Rotate, cmd.Music, cmd.Program, cmd.ScannerSelectedGobo, cmd.ScannerColor, fixtures, cmd.Blackout, cmd.Master, cmd.Master, cmd.StrobeSpeed, dmxInterfacePresent)

				if !cmd.Hide {
					if cmd.ScannerChase {
						// We are chase mode, we want the buttons to be the color selected for this scanner.
						// Remember that fixtures in the real world start with 1 not 0.
						realFixture := myFixtureNumber + 1
						// Find the color that has been selected for this fixture.
						// selected color is an index into the scanner colors selected.
						selectedColor := cmd.ScannerColor[myFixtureNumber]

						// Do we have a set of available colors for this fixture.
						_, ok := cmd.ScannerAvailableColors[realFixture]
						if ok {
							availableColors := cmd.ScannerAvailableColors[realFixture]
							red := availableColors[selectedColor].Color.R
							green := availableColors[selectedColor].Color.G
							blue := availableColors[selectedColor].Color.B
							common.LightLamp(common.ALight{X: myFixtureNumber, Y: sequence.Number, Red: red, Green: green, Blue: blue, Brightness: fixture.Shutter}, eventsForLauchpad, guiButtons)
							common.LabelButton(myFixtureNumber, sequence.Number, "", guiButtons)
						} else {
							// No color selected or available, use white.
							common.LightLamp(common.ALight{X: myFixtureNumber, Y: sequence.Number, Red: 255, Green: 255, Blue: 255, Brightness: fixture.Shutter}, eventsForLauchpad, guiButtons)
							common.LabelButton(myFixtureNumber, sequence.Number, "", guiButtons)
						}
					} else {
						// We're not in chase mode so use the color generated in the pattern generator.common.
						for _, color := range fixture.Colors {
							common.LightLamp(common.ALight{X: myFixtureNumber, Y: mySequenceNumber, Red: color.R, Green: color.G, Blue: color.B, Brightness: cmd.Master}, eventsForLauchpad, guiButtons)
							common.LabelButton(myFixtureNumber, sequence.Number, "", guiButtons)
						}
					}
				}
			}
		}
	}
}

func MapFixturesColorOnly(sequence *common.Sequence, dmxController *ft232.DMXController,
	fixtures *Fixtures, scannerColor int, dmxInterfacePresent bool) {
	for _, fixture := range fixtures.Fixtures {
		if fixture.Group-1 == sequence.Number {
			for channelNumber, channel := range fixture.Channels {
				if strings.Contains(channel.Name, "Color") {
					for _, setting := range channel.Settings {
						if setting.Number-1 == scannerColor {
							setChannel(fixture.Address+int16(channelNumber), byte(setting.Setting), dmxController, dmxInterfacePresent)
						}
					}
				}
			}
		}
	}
}

func lookUpSettingLabelInFixtureDefinition(group int, switchNumber int, channelName string, name string, label string, fixtures *Fixtures) (int, error) {

	if debug {
		fmt.Printf("lookUpSettingLabelInFixtureDefinition for %s\n", channelName)
	}

	var fixtureName string
	for _, fixture := range fixtures.Fixtures {
		if fixture.Group == group {
			if debug {
				fmt.Printf("fixture %s\n", fixture.Name)
				fmt.Printf("fixture.group %d group %d\n", fixture.Group, group)
				fmt.Printf("channels %+v\n", fixture.Channels)
			}
			fixtureName = fixture.Name
			for _, channel := range fixture.Channels {
				if debug {
					fmt.Printf("inspect channel %s for %s\n", channel.Name, name)
				}
				if channel.Name == name {
					if debug {
						fmt.Printf("channel.Settings %+v\n", channel.Settings)
					}
					for _, setting := range channel.Settings {
						if debug {
							fmt.Printf("inspect setting %+v \n", setting)
							fmt.Printf("setting.Label %s = label %s\n", setting.Label, label)
						}
						if setting.Label == label {
							if debug {
								fmt.Printf("Fixture=%s Channel.Name=%s Label=%s Name=%s Setting %d\n", fixture.Name, channel.Name, label, name, setting.Setting)
							}
							return setting.Setting, nil
						}
					}
				}
			}
		}
	}

	return 0, fmt.Errorf("label not found in fixture :%s", fixtureName)
}

func lookUpChannelNumberByNameInFixtureDefinition(group int, switchNumber int, channelName string, fixtures *Fixtures) (int, error) {

	if debug {
		fmt.Printf("lookUpChannelNumberByName for %s\n", channelName)
	}
	var fixtureName string
	for _, fixture := range fixtures.Fixtures {

		if fixture.Group == group {
			if debug {
				fmt.Printf("fixture %s\n", fixture.Name)
				fmt.Printf("fixture.group %d group %d\n", fixture.Group, group)
				fmt.Printf("channels %+v\n", fixture.Channels)
			}
			//if fixture.Number == switchNumber {
			fixtureName = fixture.Name
			for channelNumber, channel := range fixture.Channels {
				if debug {
					fmt.Printf("inspect channel %s for %s\n", channel.Name, channelName)
				}
				if channel.Name == channelName {
					if debug {
						fmt.Printf("Fixture=%s Channel=%s Number %d\n", fixture.Name, channel.Name, channel.Number)
					}
					return int(channelNumber), nil
				}
			}
		}
	}

	if debug {
		fmt.Printf("channel not found in fixture :%s", fixtureName)
	}
	return 0, fmt.Errorf("channel not found in fixture :%s", fixtureName)
}

func MapFixturesGoboOnly(sequence *common.Sequence, dmxController *ft232.DMXController,
	fixtures *Fixtures, selectedGobo int, dmxInterfacePresent bool) {

	for _, fixture := range fixtures.Fixtures {
		if fixture.Group-1 == sequence.Number {
			for channelNumber, channel := range fixture.Channels {
				if strings.Contains(channel.Name, "Gobo") {
					for _, setting := range channel.Settings {
						if setting.Number == selectedGobo {
							setChannel(fixture.Address+int16(channelNumber), byte(setting.Setting), dmxController, dmxInterfacePresent)
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
	displayFixture int, R int, G int, B int, W int, A int, uv int,
	Pan int, Tilt int, Shutter int, Rotate int, Music int, Program int,
	selectedGobo int, scannerColor map[int]int,
	fixtures *Fixtures, blackout bool, brightness int, master int, strobe int,
	dmxInterfacePresent bool) {

	// We control the brightness of each color with the brightness value.
	// The overall fixture brightness is set from the master value.
	Red := (float64(R) / 100) * (float64(brightness) / 2.55)
	Green := (float64(G) / 100) * (float64(brightness) / 2.55)
	Blue := (float64(B) / 100) * (float64(brightness) / 2.55)
	White := (float64(W) / 100) * (float64(brightness) / 2.55)
	Amber := (float64(A) / 100) * (float64(brightness) / 2.55)
	UV := (float64(uv) / 100) * (float64(brightness) / 2.55)

	for _, fixture := range fixtures.Fixtures {
		if fixture.Group-1 == mySequenceNumber {
			for channelNumber, channel := range fixture.Channels {
				if fixture.Number == displayFixture+1 {
					// Scanner channels
					if strings.Contains(channel.Name, "Pan") {
						if channel.Offset != nil {
							setChannel(fixture.Address+int16(channelNumber), byte(limitDmxValue(channel.MaxDegrees, Pan+*channel.Offset)), dmxController, dmxInterfacePresent)
						} else {
							setChannel(fixture.Address+int16(channelNumber), byte(limitDmxValue(channel.MaxDegrees, Pan)), dmxController, dmxInterfacePresent)
						}
					}
					if strings.Contains(channel.Name, "Tilt") {
						if channel.Offset != nil {
							setChannel(fixture.Address+int16(channelNumber), byte(limitDmxValue(channel.MaxDegrees, Tilt+*channel.Offset)), dmxController, dmxInterfacePresent)
						}
						setChannel(fixture.Address+int16(channelNumber), byte(limitDmxValue(channel.MaxDegrees, Tilt)), dmxController, dmxInterfacePresent)
					}
					if strings.Contains(channel.Name, "Shutter") {
						setChannel(fixture.Address+int16(channelNumber), byte(Shutter), dmxController, dmxInterfacePresent)
					}
					if strings.Contains(channel.Name, "Rotate") {
						setChannel(fixture.Address+int16(channelNumber), byte(Rotate), dmxController, dmxInterfacePresent)
					}
					if strings.Contains(channel.Name, "Music") {
						setChannel(fixture.Address+int16(channelNumber), byte(Music), dmxController, dmxInterfacePresent)
					}
					if strings.Contains(channel.Name, "Program") {
						setChannel(fixture.Address+int16(channelNumber), byte(Program), dmxController, dmxInterfacePresent)
					}
					if strings.Contains(channel.Name, "ProgramSpeed") {
						setChannel(fixture.Address+int16(channelNumber), byte(Program), dmxController, dmxInterfacePresent)
					}
					if strings.Contains(channel.Name, "Gobo") {
						for _, setting := range channel.Settings {
							if setting.Number == selectedGobo {
								setChannel(fixture.Address+int16(channelNumber), byte(setting.Setting), dmxController, dmxInterfacePresent)
							}
						}
					}
					if strings.Contains(channel.Name, "Color") {
						for colorNumber := range scannerColor {
							if colorNumber == displayFixture {
								for _, setting := range channel.Settings {
									if setting.Number-1 == scannerColor[displayFixture] {
										setChannel(fixture.Address+int16(channelNumber), byte(setting.Setting), dmxController, dmxInterfacePresent)
									}
								}
							}
						}
					}
					if strings.Contains(channel.Name, "Strobe") {
						if blackout {
							setChannel(fixture.Address+int16(channelNumber), byte(0), dmxController, dmxInterfacePresent)
						} else {
							setChannel(fixture.Address+int16(channelNumber), byte(strobe), dmxController, dmxInterfacePresent)
						}
					}
					if strings.Contains(channel.Name, "Master") || strings.Contains(channel.Name, "Dimmer") {
						if blackout {
							setChannel(fixture.Address+int16(channelNumber), byte(0), dmxController, dmxInterfacePresent)
						} else {
							setChannel(fixture.Address+int16(channelNumber), byte(master), dmxController, dmxInterfacePresent)
						}
					}
				}
				// Static value.
				if strings.Contains(channel.Name, "Static") {
					setChannel(fixture.Address+int16(channelNumber), byte(channel.Value), dmxController, dmxInterfacePresent)
				}
				// Fixture channels.
				if strings.Contains(channel.Name, "Red"+strconv.Itoa(displayFixture+1)) {
					setChannel(fixture.Address+int16(channelNumber), byte(int(Red)), dmxController, dmxInterfacePresent)
				}
				if strings.Contains(channel.Name, "Green"+strconv.Itoa(displayFixture+1)) {
					setChannel(fixture.Address+int16(channelNumber), byte(int(Green)), dmxController, dmxInterfacePresent)
				}
				if strings.Contains(channel.Name, "Blue"+strconv.Itoa(displayFixture+1)) {
					setChannel(fixture.Address+int16(channelNumber), byte(int(Blue)), dmxController, dmxInterfacePresent)
				}
				if strings.Contains(channel.Name, "White"+strconv.Itoa(displayFixture+1)) {
					setChannel(fixture.Address+int16(channelNumber), byte(int(White)), dmxController, dmxInterfacePresent)
				}
				if strings.Contains(channel.Name, "Amber"+strconv.Itoa(displayFixture+1)) {
					setChannel(fixture.Address+int16(channelNumber), byte(int(Amber)), dmxController, dmxInterfacePresent)
				}
				if strings.Contains(channel.Name, "UV"+strconv.Itoa(displayFixture+1)) {
					setChannel(fixture.Address+int16(channelNumber), byte(int(UV)), dmxController, dmxInterfacePresent)
				}
			}
		}
	}
}

func setChannel(index int16, data byte, dmxController *ft232.DMXController, dmxInterfacePresent bool) {
	if dmxInterfacePresent {
		dmxController.SetChannel(index, data)
	}
}

func MapSwitchFixture(mySequenceNumber int,
	dmxController *ft232.DMXController,
	switchNumber int, currentState int,
	fixtures *Fixtures, blackout bool, brightness int, master int,
	switchChannels map[int]common.SwitchChannel,
	SoundTriggers []*common.Trigger,
	soundConfig *sound.SoundConfig,
	dmxInterfacePresent bool) {

	var useFixture string

	if debug {
		fmt.Printf("MapSwitchFixture switchNumber %d, currentState %d\n", switchNumber, currentState)
	}

	// Step through the fixture config file looking for the group that matches mysequence number.
	for _, fixture := range fixtures.Fixtures {
		if fixture.Group-1 == mySequenceNumber {
			if fixture.UseFixture != "" {
				// use this fixture for the sequencer actions
				useFixture = fixture.UseFixture

				// Look through the switch states for this switch.
				for stateNumber, state := range fixture.States {
					if stateNumber == currentState {

						// Play Actions which send messages to a dedicated mini sequencer.
						for _, action := range state.Actions {
							newMiniSequencer(useFixture, fixture.Number, currentState, action, dmxController, fixtures, switchChannels, soundConfig, blackout, master, dmxInterfacePresent)
						}

						// Play DMX values directly to the univers.
						for _, value := range state.Values {
							if blackout {
								v, _ := strconv.ParseFloat(value.Setting, 32)
								setChannel(fixture.Address+int16(v), byte(0), dmxController, dmxInterfacePresent)
							} else {
								// This should be controlled by the master brightness
								if strings.Contains(value.Name, "master") || strings.Contains(value.Name, "dimmer") {
									v, _ := strconv.ParseFloat(value.Setting, 32)
									howBright := int((float64(v) / 100) * (float64(brightness) / 2.55))
									if strings.Contains(value.Name, "reverse") || strings.Contains(value.Name, "invert") {
										c, _ := strconv.ParseFloat(value.Channel, 32)
										setChannel(fixture.Address+int16(c), byte(reverse_dmx(howBright)), dmxController, dmxInterfacePresent)
									} else {
										c, _ := strconv.Atoi(value.Channel)
										setChannel(fixture.Address+int16(c), byte(howBright), dmxController, dmxInterfacePresent)
									}
								} else {

									// If the setting has is a number set it directly.
									if IsNumericOnly(value.Setting) {

										v, _ := strconv.ParseFloat(value.Setting, 32)
										if IsNumericOnly(value.Channel) {
											// If the channel has is a number set it directly.
											c, _ := strconv.ParseFloat(value.Channel, 32)
											setChannel(fixture.Address+int16(c), byte(v), dmxController, dmxInterfacePresent)
										} else {
											// Handle the fact that the channel may be a label as well.
											fixture := findFixtureByName(useFixture, fixtures)
											c, _ := lookUpChannelNumberByNameInFixtureDefinition(fixture.Group, switchNumber, value.Channel, fixtures)
											setChannel(fixture.Address+int16(c), byte(v), dmxController, dmxInterfacePresent)
										}

									} else {
										// If the setting contains a label, look up the label in the fixture definition.
										fixture := findFixtureByName(useFixture, fixtures)
										v, err := lookUpSettingLabelInFixtureDefinition(fixture.Group, fixture.Number, value.Channel, value.Name, value.Setting, fixtures)
										if err != nil {
											fmt.Printf("lookUpSettingLabelInFixtureDefinition error: %s\n", err.Error())
											fmt.Printf("dmxlights: error failed to find Name=%s in switch Setting=%s \n", value.Name, value.Setting)
											fmt.Printf("fixture.Name %s, fixture.Number %d\n", fixture.Name, fixture.Number)
											fmt.Printf("fixture.Group=%d, swiTch.Number=%d, value.Channel=%s, value.Name=%s, value.Setting=%s\n", fixture.Group, fixture.Number, value.Channel, value.Name, value.Setting)
											os.Exit(1)
										}

										// Handle the fact that the channel may be a label as well.
										if IsNumericOnly(value.Channel) {
											c, _ := strconv.ParseFloat(value.Channel, 32)
											setChannel(fixture.Address+int16(c), byte(v), dmxController, dmxInterfacePresent)
										} else {
											fixture := findFixtureByName(useFixture, fixtures)
											c, _ := lookUpChannelNumberByNameInFixtureDefinition(fixture.Group, switchNumber, value.Channel, fixtures)
											setChannel(fixture.Address+int16(c), byte(v), dmxController, dmxInterfacePresent)
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
}

func IsNumericOnly(str string) bool {

	if len(str) == 0 {
		return false
	}

	for _, s := range str {
		if s < '0' || s > '9' {
			return false
		}
	}
	return true
}

func turnOffFixture(fixtureName string, fixtures *Fixtures, dmxController *ft232.DMXController, dmxInterfacePresent bool) {
	fixture := findFixtureByName(fixtureName, fixtures)
	blackout := false
	master := 255
	strobeSpeed := 0
	ScannerColor := make(map[int]int)
	MapFixtures(fixture.Group-1, dmxController, fixture.Number-1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, ScannerColor, fixtures, blackout, master, master, strobeSpeed, dmxInterfacePresent)
}

func setSwitchState(switchChannels map[int]common.SwitchChannel, switchNumber int, switchPosition int, state bool, blackout bool, master int) {

	currentSwitchState := switchChannels[switchNumber]
	newSwitchState := common.SwitchChannel{
		Stop:             currentSwitchState.Stop,
		StopRotate:       currentSwitchState.StopRotate,
		KeepRotateAlive:  currentSwitchState.KeepRotateAlive,
		SequencerRunning: state,
		Blackout:         blackout,
		Master:           master,
		SwitchPosition:   switchPosition,
	}
	switchChannels[switchNumber] = newSwitchState

}

func getSwitchState(switchChannels map[int]common.SwitchChannel, switchNumber int) bool {
	return switchChannels[switchNumber].SequencerRunning
}

func getConfig(action Action) ActionConfig {

	config := ActionConfig{}

	if action.Colors != nil {
		// Find the color by name from the library of supported colors.
		colorLibrary, err := common.GetColorArrayByNames(action.Colors)
		if err != nil {
			fmt.Printf("error: %s\n", err.Error())
		}
		config.Colors = colorLibrary
	}

	switch action.Fade {
	case "Soft":
		config.Fade = 1
	case "Sharp":
		config.Fade = 10
	default:
		config.Fade = 1
	}

	switch action.Size {
	case "Short":
		config.Size = 1
	case "Medium":
		config.Size = 3
	case "Long":
		config.Size = 10
	default:
		config.Size = 3
	}

	switch action.Program {
	case "All":
		config.Program = 255
	default:
		config.Program = 0
	}

	switch action.Rotate {
	case "Off":
		config.RotateSpeed = 0
		config.Rotatable = false
	case "Slow":
		config.RotateSpeed = 1
		config.Rotatable = true
	case "Medium":
		config.RotateSpeed = 50
		config.Rotatable = true
	case "Fast":
		config.RotateSpeed = 127
		config.Rotatable = true
	default:
		config.RotateSpeed = 0
		config.Rotatable = false
	}

	switch action.Music {

	case "Internal":
		config.Music = 255
	case "Off":
		config.Music = 0
	default:
		config.Music = 0
	}

	switch action.Strobe {
	case "Off":
		config.Strobe = 0
	case "Slow":
		config.Strobe = 0
	case "Fast":
		config.Strobe = 0
	default:
		config.Strobe = 0
	}

	switch action.Speed {
	case "Slow":
		config.TriggerState = false
		config.Speed = 1 * time.Second
		config.MusicTrigger = false
	case "Medium":
		config.TriggerState = false
		config.Speed = 500 * time.Millisecond
		config.MusicTrigger = false
	case "Fast":
		config.TriggerState = false
		config.Speed = 250 * time.Millisecond
		config.MusicTrigger = false
	case "VeryFast":
		config.TriggerState = false
		config.Speed = 50 * time.Millisecond
		config.MusicTrigger = false
	case "Music":
		config.TriggerState = true
		config.Speed = time.Duration(12 * time.Hour)
		config.MusicTrigger = true
	default:
		config.TriggerState = false
		config.Speed = time.Duration(12 * time.Hour)
		config.MusicTrigger = false
	}

	return config
}

// newMiniSequencer is a simple sequencer which can be attached to a switch and a fixture to allow simple effects.
func newMiniSequencer(fixtureName string, switchNumber int, switchPosition int, action Action, dmxController *ft232.DMXController, fixturesConfig *Fixtures,
	switchChannels map[int]common.SwitchChannel, soundConfig *sound.SoundConfig,
	blackout bool, master int, dmxInterfacePresent bool) {

	switchName := fmt.Sprintf("switch%d", switchNumber)
	fixture := findFixtureByName(fixtureName, fixturesConfig)
	scannerColor := make(map[int]int)

	mySequenceNumber := fixture.Group - 1
	myFixtureNumber := fixture.Number - 1

	cfg := getConfig(action)

	if debug {
		fmt.Printf("Action %+v\n", action)
	}

	if action.Mode == "Off" {
		if debug {
			fmt.Printf("Stop mini sequence for switch number %d\n", switchNumber)
		}

		// Remember that we have stopped this mini sequencer.
		setSwitchState(switchChannels, switchNumber, switchPosition, false, blackout, master)

		// Disable this mini sequencer with the sound service.
		// Use the switch number as the unique sequence name.
		soundConfig.DisableSoundTrigger(switchName)

		select {
		case switchChannels[switchNumber].Stop <- true:
			turnOffFixture(fixtureName, fixturesConfig, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		return
	}

	if action.Mode == "Static" {
		if debug {
			fmt.Printf("Static mini sequence for switch number %d\n", switchNumber)
		}

		// Remember that we have stopped this mini sequencer.
		setSwitchState(switchChannels, switchNumber, switchPosition, false, blackout, master)

		select {
		case switchChannels[switchNumber].Stop <- true:
			turnOffFixture(fixtureName, fixturesConfig, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		color, err := common.GetRGBColorByName(action.Colors[0])
		if err != nil {
			fmt.Printf("error %d\n", err)
		}
		MapFixtures(mySequenceNumber, dmxController, myFixtureNumber, color.R, color.G, color.B, 0, 0, 0, 0, 0, 0, cfg.RotateSpeed, cfg.Music, cfg.Program, 0, scannerColor, fixturesConfig, blackout, master, master, cfg.Strobe, dmxInterfacePresent)
		return
	}

	if action.Mode == "Chase" {

		if debug {
			fmt.Printf("Chase mini sequence for switch number %d\n", switchNumber)
		}

		// Don't stop this mini sequencer if there's one running already.
		// Unless we are changing switch positions.
		if getSwitchState(switchChannels, switchNumber) &&
			switchChannels[switchNumber].SwitchPosition == switchPosition {
			setSwitchState(switchChannels, switchNumber, switchPosition, true, blackout, master)
			return
		}

		// Remember that we have started this mini sequencer.
		setSwitchState(switchChannels, switchNumber, switchPosition, true, blackout, master)

		// DeRegister this mini sequencer with the sound service.
		// Use the switch number as the unique sequence name.
		soundConfig.DisableSoundTrigger(switchName)

		// Turn off the fixture.
		select {
		case switchChannels[switchNumber].Stop <- true:
			turnOffFixture(fixtureName, fixturesConfig, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		// Register this mini sequencer with the sound service.
		if cfg.MusicTrigger {
			soundConfig.EnableSoundTrigger(switchName)
		}

		// Stop any left over sequence left over for this switch.
		select {
		case switchChannels[switchNumber].Stop <- true:
			turnOffFixture(fixtureName, fixturesConfig, dmxController, dmxInterfacePresent)
		case <-time.After(100 * time.Millisecond):
		}

		sequence := common.Sequence{
			ScannerInvert: false,
			RGBInvert:     false,
			Bounce:        false,
			ScannerChase:  false,
			RGBShift:      1,
		}
		sequence.Pattern = pattern.MakeSingleFixtureChase(cfg.Colors)
		sequence.Steps = sequence.Pattern.Steps
		sequence.NumberFixtures = 1
		// Calculate fade curve values.
		slopeOn, slopeOff := common.CalculateFadeValues(cfg.Fade, cfg.Size)
		// Calulate positions for each RGB fixture.
		optimisation := false
		scannerState := map[int]common.ScannerState{
			0: {
				Enabled: true,
			},
			1: {
				Enabled: true,
			},
			2: {
				Enabled: true,
			},
			3: {
				Enabled: true,
			},
			4: {
				Enabled: true,
			},
		}

		sequence.RGBPositions, sequence.NumberSteps = position.CalculatePositions(sequence, slopeOn, slopeOff, optimisation, scannerState)

		var rotateCounter int

		go func() {

			if cfg.Rotatable {
				// Thread to run the rotator
				go func(switchNumber int, switchChannels map[int]common.SwitchChannel) {

					for {
						rotateChannel, err := FindChannel("Rotate", myFixtureNumber, mySequenceNumber, fixturesConfig)
						if err != nil {
							fmt.Printf("rotator: %s,", err)
							return
						}
						masterChannel, err := FindChannel("Master", myFixtureNumber, mySequenceNumber, fixturesConfig)
						if err != nil {
							fmt.Printf("rotator: %s,", err)
							return
						}

						select {
						case <-switchChannels[switchNumber].StopRotate:
							time.Sleep(1 * time.Millisecond)
							setChannel(fixture.Address+int16(rotateChannel), byte(0), dmxController, dmxInterfacePresent)
							return
						case <-switchChannels[switchNumber].KeepRotateAlive:
							time.Sleep(1 * time.Millisecond)
							continue
						case <-time.After(1500 * time.Millisecond):
							setChannel(fixture.Address+int16(rotateChannel), byte(0), dmxController, dmxInterfacePresent)
							time.Sleep(250 * time.Millisecond)
							setChannel(fixture.Address+int16(masterChannel), byte(0), dmxController, dmxInterfacePresent)
						}
					}
				}(switchNumber, switchChannels)
			}

			// Wait for rotator thread to start.
			time.Sleep(100 * time.Millisecond)

			for {

				// Run through the steps in the sequence.
				// Remember every step contains infomation for all the fixtures in this group.
				for step := 0; step < sequence.NumberSteps; step++ {

					blackout = switchChannels[switchNumber].Blackout
					master = switchChannels[switchNumber].Master

					if cfg.Rotatable {
						select {
						case switchChannels[switchNumber].KeepRotateAlive <- true:
						case <-time.After(10 * time.Millisecond):
						}
					}

					if rotateCounter > 500 {
						rotateCounter = 1
					}

					if rotateCounter < 128 {
						cfg.RotateSpeed = 127
					} else {
						cfg.RotateSpeed = 128
					}

					if debug {
						fmt.Printf("switch:%d waiting for beat on %d with speed %d\n", switchNumber, switchNumber+10, cfg.Speed)
						fmt.Printf("switch:%d speed %d\n", switchNumber, cfg.Speed)
					}

					// This is were we wait for a beat or a time out equivalent to the speed.
					select {
					case <-soundConfig.SoundTriggers[switchNumber+3].Channel:
					case <-switchChannels[switchNumber].Stop:
						if cfg.Rotatable {
							switchChannels[switchNumber].StopRotate <- true
						}
						MapFixtures(mySequenceNumber, dmxController, myFixtureNumber, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, scannerColor, fixturesConfig, blackout, master, master, cfg.Strobe, dmxInterfacePresent)
						return
					case <-time.After(cfg.Speed):
					}

					// Play out fixture to DMX channels.
					position := sequence.RGBPositions[step]

					fixtures := position.Fixtures

					for fixtureNumber := 0; fixtureNumber < sequence.NumberFixtures; fixtureNumber++ {
						fixture := fixtures[fixtureNumber]
						for _, color := range fixture.Colors {
							MapFixtures(mySequenceNumber, dmxController, myFixtureNumber, color.R, color.G, color.B, color.W, 0, 0, 0, 0, 0, cfg.RotateSpeed, 0, 0, 0, scannerColor, fixturesConfig, blackout, master, master, cfg.Strobe, dmxInterfacePresent)
						}
					}

					rotateCounter++
				}
			}
		}()
	}
}

func findFixtureByName(fixtureName string, fixtures *Fixtures) *Fixture {

	for _, fixture := range fixtures.Fixtures {
		if fixture.Label == fixtureName {
			return &fixture
		}
	}
	return nil
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

func lightStaticFixture(sequence common.Sequence, myFixtureNumber int, dmxController *ft232.DMXController, eventsForLauchpad chan common.ALight, guiButtons chan common.ALight, fixturesConfig *Fixtures, enabled bool, dmxInterfacePresent bool) {

	lamp := sequence.StaticColors[myFixtureNumber]

	if sequence.Hide {
		if lamp.Flash {
			onColor := common.Color{R: lamp.Color.R, G: lamp.Color.G, B: lamp.Color.B}
			Black := common.Color{R: 0, G: 0, B: 0}
			common.FlashLight(myFixtureNumber, sequence.Number, onColor, Black, eventsForLauchpad, guiButtons)
		} else {
			common.LightLamp(common.ALight{X: myFixtureNumber, Y: sequence.Number, Red: lamp.Color.R, Green: lamp.Color.G, Blue: lamp.Color.B, Brightness: sequence.Master}, eventsForLauchpad, guiButtons)
		}
	}
	MapFixtures(sequence.Number, dmxController, myFixtureNumber, lamp.Color.R, lamp.Color.G, lamp.Color.B, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, nil, fixturesConfig, sequence.Blackout, sequence.Master, sequence.Master, sequence.StrobeSpeed, dmxInterfacePresent)

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

// turnOffFixtures is used to turn off a fixture when we stop a sequence.
func turnOffFixtures(cmd common.FixtureCommand, myFixtureNumber int, mySequenceNumber int, fixtures *Fixtures, dmxController *ft232.DMXController, eventsForLauchpad chan common.ALight, guiButtons chan common.ALight, dmxInterfacePresent bool) {
	if !cmd.Hide {
		common.LightLamp(common.ALight{X: myFixtureNumber, Y: mySequenceNumber, Red: 0, Green: 0, Blue: 0, Brightness: 0}, eventsForLauchpad, guiButtons)
	}
	MapFixtures(mySequenceNumber, dmxController, myFixtureNumber, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, cmd.ScannerColor, fixtures, cmd.Blackout, cmd.Master, cmd.Master, cmd.StrobeSpeed, dmxInterfacePresent)
}

func turnOnFixtures(cmd common.FixtureCommand, myFixtureNumber int, mySequenceNumber int, fixtures *Fixtures, dmxController *ft232.DMXController, eventsForLauchpad chan common.ALight, guiButtons chan common.ALight, dmxInterfacePresent bool) {
	if !cmd.Hide {
		common.LightLamp(common.ALight{X: myFixtureNumber, Y: mySequenceNumber, Red: 0, Green: 0, Blue: 0, Brightness: 0}, eventsForLauchpad, guiButtons)
	}
	// A map of the fixture colors.
	cmd.ScannerColor = make(map[int]int)
	cmd.ScannerColor[myFixtureNumber] = FindColor(myFixtureNumber, mySequenceNumber, "White", fixtures, cmd.ScannerColor)

	red := 255
	green := 255
	blue := 255
	white := 0
	amber := 0
	uv := 0
	pan := 128
	tilt := 128
	shutter := 255
	gobo := FindGobo(myFixtureNumber, mySequenceNumber, "Open", fixtures)
	brightness := 255
	master := 255
	rotate := 0
	music := 0
	program := 0

	MapFixtures(mySequenceNumber, dmxController, myFixtureNumber, red, green, blue, white, amber, uv, pan, tilt, shutter, rotate, music, program, gobo, cmd.ScannerColor, fixtures, false, brightness, master, cmd.StrobeSpeed, dmxInterfacePresent)
}

// findGobo takes the name of a gobo channel setting like "Open" and returns the gobo number  for this type of scanner.
func FindGobo(myFixtureNumber int, mySequenceNumber int, selectedGobo string, fixtures *Fixtures) int {
	for _, fixture := range fixtures.Fixtures {
		if fixture.Group-1 == mySequenceNumber {
			for _, channel := range fixture.Channels {
				if fixture.Number == myFixtureNumber+1 {

					if strings.Contains(channel.Name, "Gobo") {
						for _, setting := range channel.Settings {
							if setting.Name == selectedGobo {
								return setting.Number
							}
						}
					}
				}
			}
		}
	}
	return 0
}

// FindColor takes the name of a color channel setting like "White" and returns the color number for this type of scanner.
func FindColor(myFixtureNumber int, mySequenceNumber int, color string, fixtures *Fixtures, scannerColors map[int]int) int {
	for _, fixture := range fixtures.Fixtures {
		if fixture.Group-1 == mySequenceNumber {
			for _, channel := range fixture.Channels {
				if fixture.Number == myFixtureNumber+1 {
					if strings.Contains(channel.Name, "Color") {
						for colorNumber := range scannerColors {
							if colorNumber == myFixtureNumber {
								for _, setting := range channel.Settings {
									if setting.Number-1 == scannerColors[myFixtureNumber] {
										return setting.Number
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return 0
}

func FindChannel(name string, myFixtureNumber int, mySequenceNumber int, fixtures *Fixtures) (int, error) {
	for _, fixture := range fixtures.Fixtures {
		if fixture.Group-1 == mySequenceNumber {
			for channelNumber, channel := range fixture.Channels {
				if fixture.Number == myFixtureNumber+1 {
					if strings.Contains(channel.Name, name) {
						return channelNumber, nil
					}
				}
			}
		}
	}
	return 0, fmt.Errorf("error looking for rotate channel")
}
