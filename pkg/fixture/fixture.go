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
	switchChannels []common.SwitchChannel,
	soundTriggers []*common.Trigger,
	soundConfig *sound.SoundConfig,
	dmxController *ft232.DMXController,
	fixtures *Fixtures,
	dmxInterfacePresent bool) {

	// Outer loop wait for configuration.
	for {

		// Wait for first step
		cmd := <-fixtureStepChannel

		// Propogate the strobe speed.
		sequence.StrobeSpeed = cmd.StrobeSpeed

		if cmd.SetSwitch && sequence.Type == "switch" {
			MapSwitchFixture(cmd.SwitchData, cmd.State, dmxController, fixtures, sequence.Blackout, sequence.Master, sequence.Master, switchChannels, soundTriggers, soundConfig, dmxInterfacePresent, eventsForLauchpad, guiButtons)
			continue
		}

		if cmd.Clear {
			turnOffFixtures(sequence, cmd, myFixtureNumber, mySequenceNumber, fixtures, dmxController, eventsForLauchpad, guiButtons, dmxInterfacePresent)
			continue
		}

		// If we're a RGB fixture implement the flood and static features.
		if cmd.Type == "rgb" {

			if cmd.StartFlood && sequence.Label != "chaser" {
				var lamp common.Color
				if cmd.RGBStatic {
					lamp = cmd.RGBStaticColors[myFixtureNumber].Color
				} else {
					lamp = common.Color{R: 255, G: 255, B: 255}
				}
				MapFixtures(false, false, cmd.SequenceNumber, dmxController, myFixtureNumber, lamp.R, lamp.G, lamp.B, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, fixtures, sequence.Blackout, sequence.Master, sequence.Master, cmd.Strobe, cmd.StrobeSpeed, dmxInterfacePresent)
				common.LightLamp(common.ALight{X: myFixtureNumber, Y: sequence.Number, Red: lamp.R, Green: lamp.G, Blue: lamp.B, Brightness: 255}, eventsForLauchpad, guiButtons)
				continue
			}
			if cmd.StopFlood && sequence.Label != "chaser" {
				MapFixtures(false, false, cmd.SequenceNumber, dmxController, myFixtureNumber, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, fixtures, sequence.Blackout, sequence.Master, sequence.Master, cmd.Strobe, cmd.StrobeSpeed, dmxInterfacePresent)
				common.LightLamp(common.ALight{X: myFixtureNumber, Y: sequence.Number, Red: 0, Green: 0, Blue: 0, Brightness: 0}, eventsForLauchpad, guiButtons)
				continue
			}

			if cmd.RGBStatic && sequence.Label != "chaser" {
				sequence := common.Sequence{}
				sequence.Type = cmd.Type
				sequence.Number = cmd.SequenceNumber
				sequence.Master = cmd.Master
				sequence.Blackout = cmd.Blackout
				sequence.Hide = cmd.Hide
				sequence.StaticColors = cmd.RGBStaticColors
				sequence.Static = cmd.RGBStatic
				sequence.StrobeSpeed = cmd.StrobeSpeed
				sequence.Strobe = cmd.Strobe
				lightStaticFixture(sequence, myFixtureNumber, dmxController, eventsForLauchpad, guiButtons, fixtures, true, dmxInterfacePresent)
				continue
			}
			if !cmd.RGBStatic && cmd.RGBPlayStaticOnce && sequence.Label != "chaser" {
				turnOffFixture(myFixtureNumber, mySequenceNumber, fixtures, dmxController, dmxInterfacePresent)
				common.LightLamp(common.ALight{X: myFixtureNumber, Y: mySequenceNumber, Red: 0, Green: 0, Blue: 0, Brightness: cmd.Master}, eventsForLauchpad, guiButtons)
				continue
			}

			// Now play all the values for this state.

			// Play out fixture to DMX channels.
			fixture := cmd.RGBPosition.Fixtures[myFixtureNumber]
			if fixture.Enabled {
				for _, color := range fixture.Colors {
					red := color.R
					green := color.G
					blue := color.B
					white := color.W

					// If we're a shutter chaser flavoured RGB sequence, then disable everything except the brightness.
					if sequence.Label == "chaser" {
						// TODO find the scanner sequence number from the config.
						scannerFixturesSequenceNumber := 2 // Scanner sequence.
						if !cmd.Hide {
							if cmd.FixtureState.Inverted {
								common.LightLamp(common.ALight{X: myFixtureNumber, Y: scannerFixturesSequenceNumber, Red: red, Green: green, Blue: blue, Brightness: common.ReverseDmx(fixture.Brightness)}, eventsForLauchpad, guiButtons)
							} else {
								common.LightLamp(common.ALight{X: myFixtureNumber, Y: scannerFixturesSequenceNumber, Red: red, Green: green, Blue: blue, Brightness: fixture.Brightness}, eventsForLauchpad, guiButtons)
							}
						}
						// Fixture brightness is sent as master in this case.
						// TODO Integrate cmd.master with fixture.Brightness.
						if cmd.FixtureState.Inverted {
							MapFixtures(true, cmd.ScannerChaser, scannerFixturesSequenceNumber, dmxController, myFixtureNumber, red, green, blue, white, 0, 0, 0, 0, 0, 0, 0, 0, cmd.ScannerGobo, sequence.ScannerColor[myFixtureNumber], fixtures, cmd.Blackout, cmd.Master, common.ReverseDmx(fixture.Brightness), cmd.Strobe, cmd.StrobeSpeed, dmxInterfacePresent)
						} else {
							MapFixtures(true, cmd.ScannerChaser, scannerFixturesSequenceNumber, dmxController, myFixtureNumber, red, green, blue, white, 0, 0, 0, 0, 0, 0, 0, 0, cmd.ScannerGobo, sequence.ScannerColor[myFixtureNumber], fixtures, cmd.Blackout, cmd.Master, fixture.Brightness, cmd.Strobe, cmd.StrobeSpeed, dmxInterfacePresent)
						}

					} else {
						if !cmd.Hide {
							common.LightLamp(common.ALight{X: myFixtureNumber, Y: mySequenceNumber, Red: red, Green: green, Blue: blue, Brightness: cmd.Master}, eventsForLauchpad, guiButtons)
						}
						MapFixtures(false, cmd.ScannerChaser, mySequenceNumber, dmxController, myFixtureNumber, red, green, blue, white, 0, 0, 0, 0, 0, 0, 0, 0, cmd.ScannerGobo, sequence.ScannerColor[myFixtureNumber], fixtures, cmd.Blackout, cmd.Master, cmd.Master, cmd.Strobe, cmd.StrobeSpeed, dmxInterfacePresent)
					}
				}
			}
		}

		if cmd.Type == "scanner" {

			// Turn off the scanners in flood mode.
			if cmd.StartFlood {
				turnOnFixtures(sequence, cmd, myFixtureNumber, mySequenceNumber, fixtures, dmxController, eventsForLauchpad, guiButtons, dmxInterfacePresent)
				common.LightLamp(common.ALight{X: myFixtureNumber, Y: sequence.Number, Red: 255, Green: 255, Blue: 255, Brightness: 255}, eventsForLauchpad, guiButtons)
				common.LabelButton(myFixtureNumber, sequence.Number, "", guiButtons)
				continue
			}

			if cmd.StopFlood {
				common.LightLamp(common.ALight{X: myFixtureNumber, Y: sequence.Number, Red: 0, Green: 0, Blue: 0, Brightness: 0}, eventsForLauchpad, guiButtons)
				continue
			}

			// find the fixture
			fixture := cmd.ScannerPosition.Fixtures[myFixtureNumber]

			// If this fixture is disabled then shut the shutter off.
			if cmd.ScannerDisableOnce && !cmd.FixtureState.Enabled {
				turnOffFixture(myFixtureNumber, mySequenceNumber, fixtures, dmxController, dmxInterfacePresent)
				// Locking for write.
				sequence.DisableOnceMutex.Lock()
				sequence.DisableOnce[myFixtureNumber] = false
				sequence.DisableOnceMutex.Unlock()
				continue
			}

			if fixture.Enabled {

				// If enables activate the physical scanner.
				scannerColor := cmd.ScannerColor

				// In the case of a scanner, they usually have a shutter and a master dimmer to control the brightness
				// of the lamp. Problem is we can't use the shutter for the control of the overall brightness and the
				// master for the master dimmmer like we do with RGB fixture. The shutter noramlly is more of a switch
				// eg. Open , Closed , Strobe etc. If I want to slow fade through a set of scanners I need to use the
				// brightness for control. Which means I need to combine the master and the control brightness
				// at this stage.
				scannerBrightness := int(math.Round((float64(fixture.Brightness) / 100) * (float64(cmd.Master) / 2.55)))
				MapFixtures(false, cmd.ScannerChaser, mySequenceNumber, dmxController, myFixtureNumber, fixture.ScannerColor.R, fixture.ScannerColor.G, fixture.ScannerColor.B, fixture.ScannerColor.W, fixture.ScannerColor.A, fixture.ScannerColor.UV, fixture.Pan, fixture.Tilt,
					fixture.Shutter, cmd.Rotate, cmd.Music, cmd.Program, cmd.ScannerGobo, scannerColor, fixtures, cmd.Blackout, cmd.Master, scannerBrightness, cmd.Strobe, cmd.StrobeSpeed, dmxInterfacePresent)

				if !cmd.Hide {
					if cmd.ScannerChaser && sequence.Label == "chaser" {
						// We are chase mode, we want the buttons to be the color selected for this scanner.
						// Remember that fixtures in the real world start with 1 not 0.
						//realFixture := myFixtureNumber + 1
						// Find the color that has been selected for this fixture.
						// selected color is an index into the scanner colors selected.
						selectedColor := cmd.ScannerColor

						// Do we have a set of available colors for this fixture.
						if len(cmd.ScannerAvailableColors) != 0 {
							availableColors := cmd.ScannerAvailableColors
							red := availableColors[selectedColor].Color.R
							green := availableColors[selectedColor].Color.G
							blue := availableColors[selectedColor].Color.B
							common.LightLamp(common.ALight{X: myFixtureNumber, Y: sequence.Number, Red: red, Green: green, Blue: blue, Brightness: fixture.Shutter}, eventsForLauchpad, guiButtons)
							common.LabelButton(myFixtureNumber, sequence.Number, "", guiButtons)
						} else {
							// If the pattern has colors use them.
							if len(fixture.Colors) != 0 {
								for _, color := range fixture.Colors {
									common.LightLamp(common.ALight{X: myFixtureNumber, Y: mySequenceNumber, Red: color.R, Green: color.G, Blue: color.B, Brightness: cmd.Master}, eventsForLauchpad, guiButtons)
									common.LabelButton(myFixtureNumber, sequence.Number, "", guiButtons)
								}
							} else {
								// No color selected or available, use white.
								common.LightLamp(common.ALight{X: myFixtureNumber, Y: sequence.Number, Red: 222, Green: 255, Blue: 255, Brightness: fixture.Shutter}, eventsForLauchpad, guiButtons)
								common.LabelButton(myFixtureNumber, sequence.Number, "", guiButtons)
							}
						}
					} else {
						// Only fire every quarter turn of the scanner coordinates to save on launchpad mini traffic.
						if !cmd.ScannerChaser {
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
			} else {
				// This scanner is disabled, shut it off.
				turnOffFixture(myFixtureNumber, mySequenceNumber, fixtures, dmxController, dmxInterfacePresent)
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
							SetChannel(fixture.Address+int16(channelNumber), byte(v), dmxController, dmxInterfacePresent)
						}
					}
				}
			}
		}
	}
}

func findChannelSettingByLabel(fixture *Fixture, channelName string, label string) (int, error) {

	if debug {
		fmt.Printf("findChannelSettingByLabel: looking for Label -->%s<--- in Channel -->%s<-- settings\n", label, channelName)
	}

	var fixtureName string

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

	return 0, fmt.Errorf("label setting \"%s\" not found i channel \"%s\" fixture :%s", label, channelName, fixtureName)
}

func findChannelSettingByChannelNameAndSettingName(fixture *Fixture, channelName string, settingName string) (int, error) {

	if debug {
		fmt.Printf("findChannelSettingByChannelNameAndSettingName for %s\n", channelName)
	}

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

	return 0, fmt.Errorf("label not found in fixture :%s", fixture.Name)
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

func MapFixturesGoboOnly(sequence *common.Sequence, dmxController *ft232.DMXController,
	fixtures *Fixtures, selectedGobo int, dmxInterfacePresent bool) {

	for _, fixture := range fixtures.Fixtures {
		if fixture.Group-1 == sequence.Number {
			for channelNumber, channel := range fixture.Channels {
				if strings.Contains(channel.Name, "Gobo") {
					for _, setting := range channel.Settings {
						if setting.Number == selectedGobo {
							v, _ := strconv.Atoi(setting.Value)
							SetChannel(fixture.Address+int16(channelNumber), byte(v), dmxController, dmxInterfacePresent)
						}
					}
				}
			}
		}
	}
}

// When want to light a DMX fixture we need for find it in our fuxture.yaml configuration file.
// This function maps the requested fixture into a DMX address.
func MapFixtures(chaser bool, hadShutterChase bool, mySequenceNumber int,
	dmxController *ft232.DMXController,
	displayFixture int, R int, G int, B int, W int, A int, uv int,
	pan int, tilt int, shutter int, rotate int, music int, program int,
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
		if fixture.Group == mySequenceNumber+1 {
			for channelNumber, channel := range fixture.Channels {
				// Match the fixture number unless there are mulitple sub fixtures.
				if fixture.Number == displayFixture+1 || fixture.NumberChannels > 0 {
					if !chaser {
						// Scanner channels
						if strings.Contains(channel.Name, "Pan") {
							if channel.Offset != nil {
								SetChannel(fixture.Address+int16(channelNumber), byte(limitDmxValue(channel.MaxDegrees, pan+*channel.Offset)), dmxController, dmxInterfacePresent)
							} else {
								SetChannel(fixture.Address+int16(channelNumber), byte(limitDmxValue(channel.MaxDegrees, pan)), dmxController, dmxInterfacePresent)
							}
						}
						if strings.Contains(channel.Name, "Tilt") {
							if channel.Offset != nil {
								SetChannel(fixture.Address+int16(channelNumber), byte(limitDmxValue(channel.MaxDegrees, tilt+*channel.Offset)), dmxController, dmxInterfacePresent)
							}
							SetChannel(fixture.Address+int16(channelNumber), byte(limitDmxValue(channel.MaxDegrees, tilt)), dmxController, dmxInterfacePresent)
						}
						if strings.Contains(channel.Name, "Shutter") {
							// If we have defined settings for the shutter channel, then use them.
							if channel.Settings != nil {
								// Look through any settings configured for Shutter.
								for _, s := range channel.Settings {
									if !strobe && (s.Name == "On" || s.Name == "Open") {
										v := calcFinalValueBasedOnConfigAndSettingValue(s.Value, shutter)
										SetChannel(fixture.Address+int16(channelNumber), byte(v), dmxController, dmxInterfacePresent)
									}
									if strobe && strings.Contains(s.Name, "Strobe") {
										v := calcFinalValueBasedOnConfigAndSettingValue(s.Value, strobeSpeed)
										SetChannel(fixture.Address+int16(channelNumber), byte(v), dmxController, dmxInterfacePresent)
									}
								}
							} else {
								// Ok no settings. so send out the strobe speed as a 0-255 on the Shutter channel.
								SetChannel(fixture.Address+int16(channelNumber), byte(shutter), dmxController, dmxInterfacePresent)
							}
						}
						if strings.Contains(channel.Name, "Rotate") {
							SetChannel(fixture.Address+int16(channelNumber), byte(rotate), dmxController, dmxInterfacePresent)
						}
						if strings.Contains(channel.Name, "Music") {
							SetChannel(fixture.Address+int16(channelNumber), byte(music), dmxController, dmxInterfacePresent)
						}
						if strings.Contains(channel.Name, "Program") {
							SetChannel(fixture.Address+int16(channelNumber), byte(program), dmxController, dmxInterfacePresent)
						}
						if strings.Contains(channel.Name, "ProgramSpeed") {
							SetChannel(fixture.Address+int16(channelNumber), byte(program), dmxController, dmxInterfacePresent)
						}
						if !hadShutterChase {
							if strings.Contains(channel.Name, "Gobo") {
								for _, setting := range channel.Settings {
									if setting.Number == selectedGobo {
										v, _ := strconv.Atoi(setting.Value)
										SetChannel(fixture.Address+int16(channelNumber), byte(v), dmxController, dmxInterfacePresent)
									}
								}
							}
						}
						if !hadShutterChase {
							if strings.Contains(channel.Name, "Color") {
								for _, setting := range channel.Settings {
									if setting.Number-1 == scannerColor {
										v, _ := strconv.Atoi(setting.Value)
										SetChannel(fixture.Address+int16(channelNumber), byte(v), dmxController, dmxInterfacePresent)
									}
								}
							}
						}
						if strings.Contains(channel.Name, "Strobe") {
							if blackout {
								SetChannel(fixture.Address+int16(channelNumber), byte(0), dmxController, dmxInterfacePresent)
							} else {
								if strobe {
									SetChannel(fixture.Address+int16(channelNumber), byte(strobeSpeed), dmxController, dmxInterfacePresent)
								} else {
									SetChannel(fixture.Address+int16(channelNumber), byte(0), dmxController, dmxInterfacePresent)
								}
							}
						}
						// Master Dimmer.
						if !hadShutterChase {
							if strings.Contains(channel.Name, "Master") || strings.Contains(channel.Name, "Dimmer") {
								if blackout {
									SetChannel(fixture.Address+int16(channelNumber), byte(0), dmxController, dmxInterfacePresent)
								} else {
									if strings.Contains(channel.Name, "reverse") ||
										strings.Contains(channel.Name, "Reverse") ||
										strings.Contains(channel.Name, "invert") ||
										strings.Contains(channel.Name, "Invert") {
										SetChannel(fixture.Address+int16(channelNumber), byte(reverse_dmx(master)), dmxController, dmxInterfacePresent)
									} else {
										SetChannel(fixture.Address+int16(channelNumber), byte(master), dmxController, dmxInterfacePresent)
									}
								}
							}
						}
					} else { // We are a scanner chaser, so operate on brightness to master dimmer and scanner color and gobo.
						// Master Dimmer.
						if strings.Contains(channel.Name, "Master") || strings.Contains(channel.Name, "Dimmer") {
							if blackout {
								SetChannel(fixture.Address+int16(channelNumber), byte(0), dmxController, dmxInterfacePresent)
							} else {
								if strings.Contains(channel.Name, "reverse") ||
									strings.Contains(channel.Name, "Reverse") ||
									strings.Contains(channel.Name, "invert") ||
									strings.Contains(channel.Name, "Invert") {
									SetChannel(fixture.Address+int16(channelNumber), byte(reverse_dmx(master)), dmxController, dmxInterfacePresent)
								} else {
									SetChannel(fixture.Address+int16(channelNumber), byte(master), dmxController, dmxInterfacePresent)
								}
							}
						}
						// Scanner Color
						if strings.Contains(channel.Name, "Color") {
							for _, setting := range channel.Settings {
								if setting.Number-1 == scannerColor {
									v, _ := strconv.Atoi(setting.Value)
									SetChannel(fixture.Address+int16(channelNumber), byte(v), dmxController, dmxInterfacePresent)
								}
							}
						}
						// Scanner Gobo
						if strings.Contains(channel.Name, "Gobo") {
							for _, setting := range channel.Settings {
								if setting.Number == selectedGobo {
									v, _ := strconv.Atoi(setting.Value)
									SetChannel(fixture.Address+int16(channelNumber), byte(v), dmxController, dmxInterfacePresent)
								}
							}
						}
					}
				}
				if !chaser {
					// Static value.
					if strings.Contains(channel.Name, "Static") {
						if channel.Value != nil {
							SetChannel(fixture.Address+int16(channelNumber), byte(*channel.Value), dmxController, dmxInterfacePresent)
						}
					}
					// Fixture channels.
					if strings.Contains(channel.Name, "Red"+strconv.Itoa(displayFixture+1)) {
						SetChannel(fixture.Address+int16(channelNumber), byte(int(Red)), dmxController, dmxInterfacePresent)
					}
					if strings.Contains(channel.Name, "Green"+strconv.Itoa(displayFixture+1)) {
						SetChannel(fixture.Address+int16(channelNumber), byte(int(Green)), dmxController, dmxInterfacePresent)
					}
					if strings.Contains(channel.Name, "Blue"+strconv.Itoa(displayFixture+1)) {
						SetChannel(fixture.Address+int16(channelNumber), byte(int(Blue)), dmxController, dmxInterfacePresent)
					}
					if strings.Contains(channel.Name, "White"+strconv.Itoa(displayFixture+1)) {
						SetChannel(fixture.Address+int16(channelNumber), byte(int(White)), dmxController, dmxInterfacePresent)
					}
					if strings.Contains(channel.Name, "Amber"+strconv.Itoa(displayFixture+1)) {
						SetChannel(fixture.Address+int16(channelNumber), byte(int(Amber)), dmxController, dmxInterfacePresent)
					}
					if strings.Contains(channel.Name, "UV"+strconv.Itoa(displayFixture+1)) {
						SetChannel(fixture.Address+int16(channelNumber), byte(int(UV)), dmxController, dmxInterfacePresent)
					}
				}
			}
		}
	}
}

func calcFinalValueBasedOnConfigAndSettingValue(configValue string, settingValue int) (final int) {

	if strings.Contains(configValue, "-") {
		// We've found a range of values.
		// Find the start value
		numbers := strings.Split(configValue, "-")

		// Now apply the range depending on the speed
		// First turn the stings into numbers.
		start, _ := strconv.Atoi(numbers[0])
		stop, _ := strconv.Atoi(numbers[1])

		// Calculate the value depending on the setting value
		r := float32(stop) - float32(start)
		var full float32 = 255
		value := (r / full) * float32(settingValue)

		// Now apply the speed
		final = int(value) + start
		return final
	} else {
		strValue, _ := strconv.Atoi(configValue)
		return strValue
	}
}

func SetChannel(index int16, data byte, dmxController *ft232.DMXController, dmxInterfacePresent bool) {
	if dmxDebug {
		fmt.Printf("DMX Debug    Channel %d Value %d\n", index, data)
	}
	if dmxInterfacePresent {
		dmxController.SetChannel(index, data)
	}
}

// MapSwitchFixture is repsonsible for playing out the state of a swicth.
// The switch is idendifed by the sequence and switch number.
func MapSwitchFixture(swiTch common.Switch,
	state common.State,
	dmxController *ft232.DMXController,
	fixturesConfig *Fixtures, blackout bool,
	brightness int, master int,
	switchChannels []common.SwitchChannel,
	SoundTriggers []*common.Trigger,
	soundConfig *sound.SoundConfig,
	dmxInterfacePresent bool,
	eventsForLauchpad chan common.ALight,
	guiButtons chan common.ALight) {

	var useFixtureLabel string

	if debug {
		fmt.Printf("MapSwitchFixture switchNumber %d, current position %d\n", swiTch.Number, swiTch.CurrentPosition)
	}

	// We start by having the switch and its current state passed in.

	// Now we find the fixture used by the switch
	if swiTch.UseFixture != "" {
		// use this fixture for the sequencer actions
		// BTW UseFixtureLabel is the label for the fixture NOT the name.
		useFixtureLabel = swiTch.UseFixture

		if debug {
			fmt.Printf("useFixtureLabel -->%s<---  blackout is %t\n", useFixtureLabel, blackout)
		}

		// Find the details of the fixture for this switch.
		thisFixture, err := findFixtureByLabel(useFixtureLabel, fixturesConfig)
		if err != nil {
			fmt.Printf("error %s\n", err.Error())
			return
		}

		if debug {
			fmt.Printf("Found fixture Name -->%s<--- \n", thisFixture.Name)
		}

		// Look for Master channel in this fixture identified by ID.
		masterChannel, err := FindChannelNumberByName(thisFixture, "Master")
		if err != nil && debug {
			fmt.Printf("warning! fixture:%s master channel not defined: %s\n", thisFixture.Name, err)
		}

		// If blackout, set master to off.
		if blackout {
			// Blackout the fixture by setting master brightness to zero.
			if debug {
				fmt.Printf("---> SetChannel %d To Value %d\n", thisFixture.Address+int16(masterChannel), 0)
			}
			SetChannel(thisFixture.Address+int16(masterChannel), byte(0), dmxController, dmxInterfacePresent)
		}

		// Play Actions which send messages to a dedicated mini sequencer.
		for _, action := range state.Actions {
			if debug {
				fmt.Printf("actions are available\n")
			}

			newAction := Action{}
			newAction.Name = action.Name
			newAction.Number = action.Number
			newAction.Colors = action.Colors
			newAction.Mode = action.Mode
			newAction.Fade = action.Fade
			newAction.Size = action.Size
			newAction.Speed = action.Speed
			newAction.Rotate = action.Rotate
			newAction.RotateSpeed = action.RotateSpeed
			newAction.Program = action.Program
			newAction.Strobe = action.Strobe
			newMiniSequencer(thisFixture, swiTch, newAction, dmxController, fixturesConfig, switchChannels, soundConfig, blackout, brightness, master, dmxInterfacePresent, eventsForLauchpad, guiButtons)
		}

		// If there are no actions, turn off any previos mini sequencers for this switch.
		if len(state.Actions) == 0 {
			newAction := Action{}
			newAction.Name = "Off"
			newAction.Number = 1
			newAction.Mode = "Off"
			newMiniSequencer(thisFixture, swiTch, newAction, dmxController, fixturesConfig, switchChannels, soundConfig, blackout, brightness, master, dmxInterfacePresent, eventsForLauchpad, guiButtons)
		}

		// Now play any preset DMX values directly to the universe.
		// Step through all the settings.
		for _, setting := range state.Settings {

			// Process settings.
			if debug {
				fmt.Printf("settings are available\n")
			}

			// Not Blackout.
			// This should be controlled by the master brightness
			if strings.Contains(setting.Name, "master") || strings.Contains(setting.Name, "dimmer") {

				// Master brightness.
				value, _ := strconv.ParseFloat(setting.FixtureValue, 32)
				howBright := int((float64(value) / 100) * (float64(brightness) / 2.55))

				if strings.Contains(setting.Name, "reverse") || strings.Contains(setting.Name, "invert") {
					// Invert the brightness value,  some fixtures have the max brightness at 0 and off at 255.
					SetChannel(thisFixture.Address+int16(masterChannel), byte(reverse_dmx(howBright)), dmxController, dmxInterfacePresent)
				} else {
					// Set the master brightness value.
					SetChannel(thisFixture.Address+int16(masterChannel), byte(howBright), dmxController, dmxInterfacePresent)
				}

			} else {

				// If the setting value has is a number set it directly.
				if IsNumericOnly(setting.FixtureValue) {

					value, _ := strconv.Atoi(setting.FixtureValue)
					if IsNumericOnly(setting.Channel) {
						channel, _ := strconv.Atoi(setting.Channel)
						channel = channel - 1 // Channels are relative to the base address so deduct one to make it work.
						SetChannel(thisFixture.Address+int16(channel), byte(value), dmxController, dmxInterfacePresent)
					} else {
						// Handle the fact that the channel may be a label as well.
						// Look for this channels number in this fixture identified by ID.
						channel, _ := FindChannelNumberByName(thisFixture, setting.Channel)
						SetChannel(thisFixture.Address+int16(channel), byte(value), dmxController, dmxInterfacePresent)
					}

				} else {
					// The setting value is a string.
					// Lets looks to see if the setting string contains a valid fixture label, look up the label in the fixture definition.
					value, err := findChannelSettingByLabel(thisFixture, setting.Channel, setting.Label)
					if err != nil {
						fmt.Printf("error: %s\n", err.Error())
					}

					// Handle the fact that the channel may be a label as well.
					if IsNumericOnly(setting.Channel) {
						// Find the channel
						channel, _ := strconv.ParseFloat(setting.Channel, 32)
						SetChannel(thisFixture.Address+int16(channel), byte(value), dmxController, dmxInterfacePresent)
					} else {
						// Look for this channels number in this fixture identified by ID.
						channel, _ := FindChannelNumberByName(thisFixture, setting.Channel)
						SetChannel(thisFixture.Address+int16(channel), byte(value), dmxController, dmxInterfacePresent)
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

func FindFixtureAddressByGroupAndNumber(sequenceNumber int, fixtureNumber int, fixtures *Fixtures) (int16, error) {
	if debug {
		fmt.Printf("findFixtureAddressByGroupAndNumber seq %d fixture %d\n", sequenceNumber, fixtureNumber)
	}
	for _, fixture := range fixtures.Fixtures {
		if fixture.Group == sequenceNumber+1 {
			if fixture.Number == fixtureNumber+1 {
				if debug {
					fmt.Printf("Found fixture %s Group %d Number %d Address %d\n", fixture.Name, fixture.Group, fixture.Number, fixture.Address)
				}
				return fixture.Address, nil
			}
		}
	}
	return 0, fmt.Errorf("findFixtureByName: failed to find address for sequence %d fixture %d", sequenceNumber, fixtureNumber)
}

func findFixtureByLabel(label string, fixtures *Fixtures) (*Fixture, error) {

	if debug {
		fmt.Printf("Look for fixture by Label %s\n", label)
	}

	if label == "" {
		return nil, fmt.Errorf("findFixtureByLabel: fixture label is empty")
	}

	for _, fixture := range fixtures.Fixtures {
		if strings.Contains(fixture.Label, label) {
			if debug {
				fmt.Printf("Found fixture %s Group %d Number %d Address %d\n", fixture.Name, fixture.Group, fixture.Number, fixture.Address)
			}
			return &fixture, nil
		}
	}
	return nil, fmt.Errorf("findFixtureByLabel: failed to find fixture by label– %s", label)
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

	if debug {
		fmt.Printf("Reverse in is %d\n", n)
	}

	in := make(map[int]int, 255)
	var y = 255

	for x := 0; x <= 255; x++ {

		in[x] = y
		y--
	}
	if debug {
		fmt.Printf("Reverse out is %d\n", in[n])
	}
	return in[n]
}

func lightStaticFixture(sequence common.Sequence, myFixtureNumber int, dmxController *ft232.DMXController, eventsForLauchpad chan common.ALight, guiButtons chan common.ALight, fixturesConfig *Fixtures, enabled bool, dmxInterfacePresent bool) {

	if debug {
		fmt.Printf("lightStaticFixture seq %d fixture %d \n", sequence.Number, myFixtureNumber)
	}

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
	if debug {
		fmt.Printf("strobe %t speed %d\n", sequence.Strobe, sequence.StrobeSpeed)
	}
	MapFixtures(false, false, sequence.Number, dmxController, myFixtureNumber, lamp.Color.R, lamp.Color.G, lamp.Color.B, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, sequence.Blackout, sequence.Master, sequence.Master, sequence.Strobe, sequence.StrobeSpeed, dmxInterfacePresent)

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

	//if debug {
	fmt.Printf("Sequence %d: Fixture %d turnOffFixtures\n", sequence.Number, myFixtureNumber)
	//}
	common.LabelButton(myFixtureNumber, sequence.Number, "", guiButtons)
	MapFixtures(false, false, mySequenceNumber, dmxController, myFixtureNumber, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, fixtures, cmd.Blackout, cmd.Master, cmd.Master, cmd.Strobe, cmd.StrobeSpeed, dmxInterfacePresent)
}

// FindShutter takes the name of a gobo channel setting like "Open" and returns the gobo number  for this type of scanner.
func FindShutter(myFixtureNumber int, mySequenceNumber int, shutterName string, fixtures *Fixtures) int {
	for _, fixture := range fixtures.Fixtures {
		if fixture.Group == mySequenceNumber+1 {
			if fixture.Number == myFixtureNumber+1 {
				for _, channel := range fixture.Channels {
					if strings.Contains(channel.Name, "Shutter") {
						for _, setting := range channel.Settings {
							if strings.Contains(setting.Name, shutterName) {
								return setting.Number
							}
						}
					}
				}
			}
		}
	}
	return 255
}

// findGobo takes the name of a gobo channel setting like "Open" and returns the gobo number  for this type of scanner.
func FindGobo(myFixtureNumber int, mySequenceNumber int, selectedGobo string, fixtures *Fixtures) int {
	for _, fixture := range fixtures.Fixtures {
		if fixture.Group == mySequenceNumber+1 {
			if fixture.Number == myFixtureNumber+1 {
				for _, channel := range fixture.Channels {
					if strings.Contains(channel.Name, "Gobo") {
						for _, setting := range channel.Settings {
							if strings.Contains(setting.Name, selectedGobo) {
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
func FindColor(myFixtureNumber int, mySequenceNumber int, color string, fixtures *Fixtures) int {
	for _, fixture := range fixtures.Fixtures {
		if fixture.Group == mySequenceNumber+1 {
			if fixture.Number == myFixtureNumber+1 {
				for _, channel := range fixture.Channels {
					if strings.Contains(channel.Name, "Color") {
						for _, setting := range channel.Settings {
							if setting.Name == color {
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

func FindChannelNumberByName(fixture *Fixture, channelName string) (int, error) {
	{
		for channelNumber, channel := range fixture.Channels {
			if strings.Contains(channel.Name, channelName) {
				return channelNumber, nil
			}
		}
	}
	return 0, fmt.Errorf("error looking for channel %s", channelName)
}

func turnOnFixtures(sequence common.Sequence, cmd common.FixtureCommand, myFixtureNumber int, mySequenceNumber int, fixtures *Fixtures, dmxController *ft232.DMXController, eventsForLauchpad chan common.ALight, guiButtons chan common.ALight, dmxInterfacePresent bool) {
	if !cmd.Hide {
		common.LightLamp(common.ALight{X: myFixtureNumber, Y: mySequenceNumber, Red: 0, Green: 0, Blue: 0, Brightness: 0}, eventsForLauchpad, guiButtons)
	}
	// Find the color number for White.
	red := 255
	green := 255
	blue := 255
	white := 0
	amber := 0
	uv := 0
	pan := 128
	tilt := 128
	shutter := FindShutter(myFixtureNumber, mySequenceNumber, "Open", fixtures)
	//fmt.Printf("shutter  %d\n", shutter)
	gobo := FindGobo(myFixtureNumber, mySequenceNumber, "White", fixtures)
	//fmt.Printf("gobo  %d\n", gobo)
	scannerColor := FindColor(myFixtureNumber, mySequenceNumber, "White", fixtures)
	//fmt.Printf("color %d\n", scannerColor)
	brightness := 255
	master := 255
	rotate := 0
	music := 0
	program := 0

	MapFixtures(false, false, mySequenceNumber, dmxController, myFixtureNumber, red, green, blue, white, amber, uv, pan, tilt, shutter, rotate, music, program, gobo, scannerColor, fixtures, false, brightness, master, cmd.Strobe, cmd.StrobeSpeed, dmxInterfacePresent)
}

func turnOffFixture(myFixtureNumber int, mySequenceNumber int, fixtures *Fixtures, dmxController *ft232.DMXController, dmxInterfacePresent bool) {

	blackout := false
	master := 0
	brightness := 0
	strobeSpeed := 0
	strobe := false

	// Find the color number for White.
	scannerColor := FindColor(myFixtureNumber, mySequenceNumber, "White", fixtures)
	red := 0
	green := 0
	blue := 0
	white := 0
	amber := 0
	uv := 0
	shutter := FindShutter(myFixtureNumber, mySequenceNumber, "Open", fixtures)
	pan := 127
	tilt := 127
	gobo := FindGobo(myFixtureNumber, mySequenceNumber, "Open", fixtures)
	rotate := 0
	music := 0
	program := 0

	if debug {
		fmt.Printf("mySequenceNumber %d\n", mySequenceNumber)
		fmt.Printf("myFixtureNumber %d\n", myFixtureNumber)
		fmt.Printf("red %d\n", red)
		fmt.Printf("green %d\n", green)
		fmt.Printf("blue %d\n", blue)
		fmt.Printf("pan %d\n", pan)
		fmt.Printf("tilt %d\n", tilt)
		fmt.Printf("shutter %d\n", shutter)
		fmt.Printf("rotate %d\n", rotate)
		fmt.Printf("music %d\n", music)
		fmt.Printf("program %d\n", program)
		fmt.Printf("gobo %d\n", gobo)
		fmt.Printf("scannerColor %d\n", scannerColor)
		fmt.Printf("brightness %d\n", brightness)
		fmt.Printf("master %d\n", master)
		fmt.Printf("Strobe %t\n", strobe)
		fmt.Printf("StrobeSpeed %d\n", strobeSpeed)
	}
	MapFixtures(false, false, mySequenceNumber, dmxController, myFixtureNumber, red, green, blue, white, amber, uv, pan, tilt, shutter, rotate, music, program, gobo, scannerColor, fixtures, blackout, brightness, master, strobe, strobeSpeed, dmxInterfacePresent)
}
