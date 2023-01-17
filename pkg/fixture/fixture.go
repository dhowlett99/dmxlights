// Copyright (C) 2022,2023 dhowlett99.
// This is the dmxlights fixture control code, it sends messages to fixtures
// using the usb dmx library.
// Implemented and depends on usbdmx.
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
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/sound"
	"github.com/go-yaml/yaml"
	"github.com/oliread/usbdmx/ft232"
)

const debug = false
const dmxDebug = false

type Fixtures struct {
	Fixtures []Fixture `yaml:"fixtures"`
}

type Color struct {
	R int `yaml:"red"`
	G int `yaml:"green"`
	B int `yaml:"blue"`
	W int `yaml:"white"`
}

type State struct {
	Name        string    `yaml:"name"`
	Number      int16     `yaml:"number"`
	Label       string    `yaml:"label"`
	ButtonColor string    `yaml:"buttoncolor"`
	Master      int       `yaml:"master"`
	Actions     []Action  `yaml:"actions,omitempty"`
	Settings    []Setting `yaml:"settings,omitempty"`
	Flash       bool      `yaml:"flash"`
}

type Action struct {
	Name        string `yaml:"name"`
	Number      int
	Colors      []string `yaml:"colors"`
	Mode        string   `yaml:"mode"`
	Fade        string   `yaml:"fade"`
	Size        string   `yaml:"size"`
	Speed       string   `yaml:"speed"`
	Rotate      string   `yaml:"rotate"`
	RotateSpeed string   `yaml:"rotatespeed"`
	Program     string   `yaml:"program"`
	Strobe      string   `yaml:"strobe"`
}

type ActionConfig struct {
	Name          string
	Colors        []common.Color
	Fade          int
	Size          int
	Speed         time.Duration
	TriggerState  bool
	RotateSpeed   int
	Rotatable     bool
	Clockwise     bool
	AntiClockwise bool
	Auto          bool
	Program       int
	Music         int
	MusicTrigger  bool
	Strobe        bool
	StrobeSpeed   int
}

type Fixture struct {
	ID             int       `yaml:"id"`
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
	Label   string `yaml:"labe,omitempty"`
	Number  int    `yaml:"number"`
	Channel string `yaml:"channel,omitempty"`
	Value   string `yaml:"value"`
}

type Channel struct {
	Number     int16     `yaml:"number"`
	Name       string    `yaml:"name"`
	Value      *int16    `yaml:"value,omitempty"`
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
	data, err := os.ReadFile(filename)
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
	_, err = os.Open(filename)
	if err != nil {
		return errors.New("error: opening fixtures.yaml file: " + err.Error())
	}

	// Write the fixtures.yaml file.
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return errors.New("error: writing fixtures.yaml file: " + err.Error())
	}

	// Fixtures file saved, no errors.
	return nil
}

// GetFixureDetailsById - find a fixture in the fixtures config.
// Returns details of the fixture.
// Returns an error.
func GetFixureDetailsById(id int, fixtures *Fixtures) (Fixture, error) {
	// scan the fixtures structure for the selected fixture.
	if debug {
		fmt.Printf("Looking for Fixture ID %d\n", id)
	}

	for _, fixture := range fixtures.Fixtures {
		if debug {
			fmt.Printf("Fixture ID %d and Name %s States %+v\n", fixture.ID, fixture.Name, fixture.States)
		}
		if fixture.ID == id {
			return fixture, nil
		}
	}
	return Fixture{}, fmt.Errorf("error: fixture id %d not found", id)
}

// GetFixureDetailsByLabel - find a fixture in the fixtures config.
// Returns details of the fixture.
// Returns an error.
func GetFixureDetailsByLabel(label string, fixtures *Fixtures) (Fixture, error) {
	// scan the fixtures structure for the selected fixture.
	if debug {
		fmt.Printf("Looking for Fixture by Label %s\n", label)
	}

	for _, fixture := range fixtures.Fixtures {
		if debug {
			fmt.Printf("Fixture Label %s and Name %s States %+v\n", fixture.Label, fixture.Name, fixture.States)
		}
		if fixture.Label == label {
			return fixture, nil
		}
	}
	return Fixture{}, fmt.Errorf("error: fixture label %s not found", label)
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
				turnOffFixtures(sequence, cmd, myFixtureNumber, mySequenceNumber, fixtures, dmxController, eventsForLauchpad, guiButtons, dmxInterfacePresent)
				continue
			}

			if cmd.StartFlood {
				var lamp common.Color
				if cmd.RGBStatic {
					lamp = cmd.RGBStaticColors[myFixtureNumber].Color
				} else {
					lamp = common.Color{R: 255, G: 255, B: 255}
				}
				MapFixtures(cmd.SequenceNumber, dmxController, myFixtureNumber, lamp.R, lamp.G, lamp.B, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, fixtures, sequence.Blackout, sequence.Master, sequence.Master, cmd.Strobe, cmd.StrobeSpeed, dmxInterfacePresent)
				common.LightLamp(common.ALight{X: myFixtureNumber, Y: sequence.Number, Red: lamp.R, Green: lamp.G, Blue: lamp.B, Brightness: 255}, eventsForLauchpad, guiButtons)
				continue
			}
			if cmd.StopFlood {
				MapFixtures(cmd.SequenceNumber, dmxController, myFixtureNumber, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, fixtures, sequence.Blackout, sequence.Master, sequence.Master, cmd.Strobe, cmd.StrobeSpeed, dmxInterfacePresent)
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
				scannerColor := common.MapCopy(cmd.ScannerColor, sequence.ScannerColorMutex)
				sequence.ScannerColorMutex.RLock()
				MapFixtures(mySequenceNumber, dmxController, myFixtureNumber, red, green, blue, white, 0, 0, 0, 0, 0, 0, 0, 0, 0, scannerColor[myFixtureNumber], fixtures, cmd.Blackout, cmd.Master, cmd.Master, cmd.Strobe, cmd.StrobeSpeed, dmxInterfacePresent)
				sequence.ScannerColorMutex.RUnlock()
			}
		}

		if cmd.Type == "scanner" {

			// Turn off the scanners in flood mode.
			if cmd.StartFlood {
				turnOnFixtures(sequence, cmd, myFixtureNumber, mySequenceNumber, fixtures, dmxController, eventsForLauchpad, guiButtons, dmxInterfacePresent)
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
				MapFixtures(mySequenceNumber, dmxController, myFixtureNumber, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, fixtures, cmd.Blackout, 0, 0, false, 0, dmxInterfacePresent)
				// Locking for write.
				sequence.DisableOnceMutex.Lock()
				sequence.DisableOnce[myFixtureNumber] = false
				sequence.DisableOnceMutex.Unlock()
				continue
			}

			if enabled {

				// If enables activate the physical scanner.
				scannerColor := common.MapCopy(cmd.ScannerColor, sequence.ScannerColorMutex)
				sequence.ScannerColorMutex.RLock()
				MapFixtures(mySequenceNumber, dmxController, myFixtureNumber, fixture.ScannerColor.R, fixture.ScannerColor.G, fixture.ScannerColor.B, fixture.ScannerColor.W, fixture.ScannerColor.A, fixture.ScannerColor.UV, fixture.Pan, fixture.Tilt,
					fixture.Shutter, cmd.Rotate, cmd.Music, cmd.Program, cmd.ScannerSelectedGobo, scannerColor[myFixtureNumber], fixtures, cmd.Blackout, cmd.Master, cmd.Master, cmd.Strobe, cmd.StrobeSpeed, dmxInterfacePresent)
				sequence.ScannerColorMutex.RUnlock()

				if !cmd.Hide {
					if cmd.ScannerChase {
						// We are chase mode, we want the buttons to be the color selected for this scanner.
						// Remember that fixtures in the real world start with 1 not 0.
						realFixture := myFixtureNumber + 1
						// Find the color that has been selected for this fixture.
						// selected color is an index into the scanner colors selected.
						sequence.ScannerColorMutex.RLock()
						selectedColor := cmd.ScannerColor[myFixtureNumber]
						sequence.ScannerColorMutex.RUnlock()

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
						// Only fire every quarter turn of the scanner coordinates to save on launchpad mini traffic.
						howOftern := cmd.NumberSteps / 4
						if howOftern != 0 {
							if cmd.Step%howOftern == 0 {
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
							v, _ := strconv.ParseFloat(setting.Value, 32)
							setChannel(fixture.Address+int16(channelNumber), byte(v), dmxController, dmxInterfacePresent)
						}
					}
				}
			}
		}
	}
}

func findChannelSettingByLabel(group int, switchNumber int, channelName string, label string, fixtures *Fixtures) (int, error) {

	if debug {
		fmt.Printf("findChannelSettingByLabel: looking for Label %s in Channel %s settings\n", label, channelName)
	}

	var fixtureName string
	for _, fixture := range fixtures.Fixtures {

		// Match the group.
		if fixture.Group == group {
			if debug {
				fmt.Printf("fixture %s\n", fixture.Name)
				fmt.Printf("fixture.group %d group %d\n", fixture.Group, group)
			}

			// Look through the channels.
			for _, channel := range fixture.Channels {
				if debug {
					fmt.Printf("inspect channel %s for %s\n", channel.Name, channelName)
				}
				// Match the channel.
				if channel.Name == channelName {
					if debug {
						fmt.Printf("channel.Settings %+v\n", channel.Settings)
					}

					// Look through the settings.
					for _, setting := range channel.Settings {
						if debug {
							fmt.Printf("inspect setting -> Label %s = label %s\n", setting.Label, label)
						}

						// Match Setting.
						if setting.Label == label {
							if debug {
								fmt.Printf("Found! Fixture.Name=%s Channel.Name=%s Label=%s Setting.Name %s Setting.Value %s\n", fixture.Name, channel.Name, label, setting.Name, setting.Value)
							}
							v, _ := strconv.Atoi(setting.Value)
							return v, nil
						}
					}
				}
			}
		}
	}

	return 0, fmt.Errorf("label setting \"%s\" not found i channel \"%s\" fixture :%s", label, channelName, fixtureName)
}

func findChannelSettingByName(group int, switchNumber int, channelName string, settingName string, fixtures *Fixtures) (int, error) {

	if debug {
		fmt.Printf("findChannelSettingByName for %s\n", channelName)
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
					fmt.Printf("inspect channel %s for %s\n", channel.Name, settingName)
				}
				if channel.Name == channelName {
					if debug {
						fmt.Printf("channel.Settings %+v\n", channel.Settings)
					}
					for _, setting := range channel.Settings {
						if debug {
							fmt.Printf("inspect setting %+v \n", setting)
							fmt.Printf("setting.Name %s = name %s\n", setting.Name, settingName)
						}
						if setting.Name == settingName {
							if debug {
								fmt.Printf("FixtureName=%s ChannelName=%s SettingName=%s SettingValue=%s\n", fixture.Name, channel.Name, settingName, setting.Value)
							}
							v, _ := strconv.Atoi(setting.Value)
							return v, nil
						}
					}
				}
			}
		}
	}

	return 0, fmt.Errorf("label not found in fixture :%s", fixtureName)
}

func findChannelSettingByNameAndSpeed(fixtureName string, channelName string, settingName string, settingSpeed string, fixtures *Fixtures) (int, error) {

	if debug {
		fmt.Printf("findChannelSettingByNameAndSpeed for name %s and speed %s\n", settingName, settingSpeed)
	}

	for _, fixture := range fixtures.Fixtures {

		if debug {
			fmt.Printf("fixture %s\n", fixture.Name)
			fmt.Printf("channels %+v\n", fixture.Channels)
		}
		if fixtureName == fixture.Name {
			for _, channel := range fixture.Channels {
				if debug {
					fmt.Printf("inspect channel %s for %s\n", channel.Name, settingName)
				}
				if channel.Name == channelName {
					if debug {
						fmt.Printf("channel.Settings %+v\n", channel.Settings)
					}
					for _, setting := range channel.Settings {
						if debug {
							fmt.Printf("inspect setting %+v \n", setting)
							fmt.Printf("got:setting.Name %s  want name %s speed %s\n", setting.Name, settingName, settingSpeed)
						}
						if strings.Contains(setting.Name, settingName) && strings.Contains(setting.Name, settingSpeed) {

							if debug {
								fmt.Printf("FixtureName=%s ChannelName=%s SettingName=%s SettingSpeed=%s, SettingValue=%s\n", fixture.Name, channel.Name, settingName, settingSpeed, setting.Value)
							}
							v, _ := strconv.Atoi(setting.Value)
							return v, nil
						}
					}
				}
			}
		}
	}

	return 0, fmt.Errorf("channel %s setting %s not found in fixture :%s", channelName, settingSpeed, fixtureName)
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
							v, _ := strconv.Atoi(setting.Value)
							setChannel(fixture.Address+int16(channelNumber), byte(v), dmxController, dmxInterfacePresent)
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
	selectedGobo int, scannerColor int,
	fixtures *Fixtures, blackout bool, brightness int, master int, strobe bool, strobeSpeed int,
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
						// If we have defined settings for the shutter channel, then use them.
						if channel.Settings != nil {
							// Look through any settings configured for Shutter.
							for _, s := range channel.Settings {
								if !strobe && s.Name == "On" || s.Name == "Open" {
									v, _ := strconv.Atoi(s.Value)
									setChannel(fixture.Address+int16(channelNumber), byte(Shutter+v), dmxController, dmxInterfacePresent)
								}
								if strobe && strings.Contains(s.Name, "Strobe") {
									// Found some stobe settings.
									if strings.Contains(s.Value, "-") {
										// We've found a range of values.
										// Find the start value// Find the start
										numbers := strings.Split(s.Value, "-")

										// Now apply the range depending on the speed
										// First turn the stings into numbers.
										start, _ := strconv.Atoi(numbers[0])
										stop, _ := strconv.Atoi(numbers[1])

										// Calculate the value depending on the strobe speed.
										r := float32(stop) - float32(start)
										var full float32 = 255
										value := (r / full) * float32(strobeSpeed)

										// Now apply the speed
										final := int(value) + start
										setChannel(fixture.Address+int16(channelNumber), byte(final), dmxController, dmxInterfacePresent)
									}
								}
							}
						} else {
							// Ok no settings. so send out the strobe speed as a 0-255 on the Shutter channel.
							setChannel(fixture.Address+int16(channelNumber), byte(Shutter), dmxController, dmxInterfacePresent)
						}
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
								v, _ := strconv.Atoi(setting.Value)
								setChannel(fixture.Address+int16(channelNumber), byte(v), dmxController, dmxInterfacePresent)
							}
						}
					}
					if strings.Contains(channel.Name, "Color") {

						for _, setting := range channel.Settings {

							if setting.Number-1 == scannerColor {
								v, _ := strconv.Atoi(setting.Value)
								setChannel(fixture.Address+int16(channelNumber), byte(v), dmxController, dmxInterfacePresent)
							}

						}

					}
					if strings.Contains(channel.Name, "Strobe") {
						if blackout {
							setChannel(fixture.Address+int16(channelNumber), byte(0), dmxController, dmxInterfacePresent)
						} else {
							setChannel(fixture.Address+int16(channelNumber), byte(strobeSpeed), dmxController, dmxInterfacePresent)
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
					if channel.Value != nil {
						setChannel(fixture.Address+int16(channelNumber), byte(*channel.Value), dmxController, dmxInterfacePresent)
					}
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
		if dmxDebug {
			fmt.Printf("DMX Debug    Channel %d Value %d\n", index, data)
		}
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
		if fixture.Group-1 == mySequenceNumber && fixture.Number-1 == switchNumber {
			if fixture.UseFixture != "" {
				// use this fixture for the sequencer actions
				useFixture = fixture.UseFixture

				if debug {
					fmt.Printf("useFixture %s\n", fixture.UseFixture)
				}

				// Look through the switch states for this switch.
				for stateNumber, state := range fixture.States {
					if stateNumber == currentState {

						// Play Actions which send messages to a dedicated mini sequencer.
						for _, action := range state.Actions {
							newMiniSequencer(useFixture, fixture.Number, currentState, action, dmxController, fixtures, switchChannels, soundConfig, blackout, master, dmxInterfacePresent)
						}

						// Play DMX values directly to the univers.
						for _, setting := range state.Settings {
							if blackout {
								v, _ := strconv.ParseFloat(setting.Value, 32)
								setChannel(fixture.Address+int16(v), byte(0), dmxController, dmxInterfacePresent)
							} else {
								// This should be controlled by the master brightness
								if strings.Contains(setting.Name, "master") || strings.Contains(setting.Name, "dimmer") {
									v, _ := strconv.ParseFloat(setting.Value, 32)
									howBright := int((float64(v) / 100) * (float64(brightness) / 2.55))
									if strings.Contains(setting.Name, "reverse") || strings.Contains(setting.Name, "invert") {
										c, _ := strconv.ParseFloat(setting.Channel, 32)
										setChannel(fixture.Address+int16(c), byte(reverse_dmx(howBright)), dmxController, dmxInterfacePresent)
									} else {
										c, _ := strconv.Atoi(setting.Channel)
										setChannel(fixture.Address+int16(c), byte(howBright), dmxController, dmxInterfacePresent)
									}
								} else {

									// If the setting has is a number set it directly.
									if IsNumericOnly(setting.Value) {

										v, _ := strconv.ParseFloat(setting.Value, 32)
										if IsNumericOnly(setting.Channel) {
											// If the channel has is a number set it directly.
											c, _ := strconv.ParseFloat(setting.Channel, 32)
											setChannel(fixture.Address+int16(c), byte(v), dmxController, dmxInterfacePresent)
										} else {
											// Handle the fact that the channel may be a label as well.
											fixture, err := findFixtureByName(useFixture, fixtures)
											if err != nil {
												fmt.Printf("error %s\n", err.Error())
												return
											}
											c, _ := lookUpChannelNumberByNameInFixtureDefinition(fixture.Group, switchNumber, setting.Channel, fixtures)
											setChannel(fixture.Address+int16(c), byte(v), dmxController, dmxInterfacePresent)
										}

									} else {
										// If the setting contains a label, look up the label in the fixture definition.
										fixture, err := findFixtureByName(useFixture, fixtures)
										if err != nil {
											fmt.Printf("error %s\n", err.Error())
											return
										}
										v, err := findChannelSettingByLabel(fixture.Group, fixture.Number, setting.Channel, setting.Label, fixtures)
										if err != nil {
											fmt.Printf("findChannelSettingByLabel error: %s\n", err.Error())
											fmt.Printf("fixture.Group=%d, fixture.Number=%d, setting.Channel=%s, setting.Label=%s, setting.Value=%s\n", fixture.Group, fixture.Number, setting.Channel, setting.Label, setting.Value)
										}

										// Handle the fact that the channel may be a label as well.
										if IsNumericOnly(setting.Channel) {
											c, _ := strconv.ParseFloat(setting.Channel, 32)
											setChannel(fixture.Address+int16(c), byte(v), dmxController, dmxInterfacePresent)
										} else {
											fixture, err := findFixtureByName(useFixture, fixtures)
											if err != nil {
												fmt.Printf("error %s\n", err.Error())
												return
											}
											c, _ := lookUpChannelNumberByNameInFixtureDefinition(fixture.Group, switchNumber, setting.Channel, fixtures)
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
	fixture, err := findFixtureByName(fixtureName, fixtures)
	if err != nil {
		fmt.Printf("error %s\n", err.Error())
		return
	}
	blackout := false
	master := 255
	strobeSpeed := 0
	strobe := false

	MapFixtures(fixture.Group-1, dmxController, fixture.Number-1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, fixtures, blackout, master, master, strobe, strobeSpeed, dmxInterfacePresent)
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

func findFixtureByName(fixtureName string, fixtures *Fixtures) (*Fixture, error) {
	if debug {
		fmt.Printf("Look for fixture by Name %s\n", fixtureName)
	}
	for _, fixture := range fixtures.Fixtures {
		if strings.Contains(fixture.Label, fixtureName) {
			if debug {
				fmt.Printf("Found fixture %s Group %d Number %d Address %d\n", fixture.Name, fixture.Group, fixture.Number, fixture.Address)
			}
			return &fixture, nil
		}
	}
	return nil, fmt.Errorf("findFixtureByName: failed to find fixture %s", fixtureName)
}

func FindFixtureAddressByName(fixtureName string, fixtures *Fixtures) string {
	if debug {
		fmt.Printf("Looking for fixture by Name %s\n", fixtureName)
	}
	for _, fixture := range fixtures.Fixtures {
		if fixture.Label == fixtureName {
			if debug {
				fmt.Printf("Found fixture %s Group %d Number %d Address %d\n", fixture.Name, fixture.Group, fixture.Number, fixture.Address)
			}
			return fmt.Sprintf("%d", fixture.Address)
		}
	}
	if debug {
		fmt.Printf("fixture %s not found\n", fixtureName)
	}
	return "Not Set"
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
	MapFixtures(sequence.Number, dmxController, myFixtureNumber, lamp.Color.R, lamp.Color.G, lamp.Color.B, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, sequence.Blackout, sequence.Master, sequence.Master, sequence.Strobe, sequence.StrobeSpeed, dmxInterfacePresent)

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
func turnOffFixtures(sequence common.Sequence, cmd common.FixtureCommand, myFixtureNumber int, mySequenceNumber int, fixtures *Fixtures, dmxController *ft232.DMXController, eventsForLauchpad chan common.ALight, guiButtons chan common.ALight, dmxInterfacePresent bool) {
	if !cmd.Hide {
		common.LightLamp(common.ALight{X: myFixtureNumber, Y: mySequenceNumber, Red: 0, Green: 0, Blue: 0, Brightness: 0}, eventsForLauchpad, guiButtons)
	}
	MapFixtures(mySequenceNumber, dmxController, myFixtureNumber, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, fixtures, cmd.Blackout, cmd.Master, cmd.Master, cmd.Strobe, cmd.StrobeSpeed, dmxInterfacePresent)
}

func turnOnFixtures(sequence common.Sequence, cmd common.FixtureCommand, myFixtureNumber int, mySequenceNumber int, fixtures *Fixtures, dmxController *ft232.DMXController, eventsForLauchpad chan common.ALight, guiButtons chan common.ALight, dmxInterfacePresent bool) {
	if !cmd.Hide {
		common.LightLamp(common.ALight{X: myFixtureNumber, Y: mySequenceNumber, Red: 0, Green: 0, Blue: 0, Brightness: 0}, eventsForLauchpad, guiButtons)
	}
	// Create a temporary map of fixture colors.
	cmd.ScannerColor = make(map[int]int)
	sequence.ScannerColorMutex.Lock()
	cmd.ScannerColor[myFixtureNumber] = FindColor(myFixtureNumber, mySequenceNumber, "White", fixtures, cmd.ScannerColor)
	sequence.ScannerColorMutex.Unlock()

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

	scannerColor := common.MapCopy(cmd.ScannerColor, sequence.ScannerColorMutex)
	MapFixtures(mySequenceNumber, dmxController, myFixtureNumber, red, green, blue, white, amber, uv, pan, tilt, shutter, rotate, music, program, gobo, scannerColor[myFixtureNumber], fixtures, false, brightness, master, cmd.Strobe, cmd.StrobeSpeed, dmxInterfacePresent)
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

// FindChannel - takes a channel name and returns the channel number of the fixture.
// Returns an error if not found.
func FindChannel(channelName string, myFixtureNumber int, mySequenceNumber int, fixtures *Fixtures) (int, error) {
	for _, fixture := range fixtures.Fixtures {
		if fixture.Group-1 == mySequenceNumber {
			for channelNumber, channel := range fixture.Channels {
				if fixture.Number == myFixtureNumber+1 {
					if strings.Contains(channel.Name, channelName) {
						return channelNumber, nil
					}
				}
			}
		}
	}
	return 0, fmt.Errorf("error looking for channel %s", channelName)
}
