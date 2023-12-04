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
	NumberSteps   int
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
	ID                 int       `yaml:"id"`
	Name               string    `yaml:"name"`
	Label              string    `yaml:"label,omitempty"`
	Number             int       `yaml:"number"`
	Description        string    `yaml:"description"`
	Type               string    `yaml:"type"`
	Group              int       `yaml:"group"`
	Address            int16     `yaml:"address"`
	Channels           []Channel `yaml:"channels"`
	States             []State   `yaml:"states,omitempty"`
	MultiFixtureDevice bool      `yaml:"-"` // Calulated internally.
	NumberSubFixtures  int       `yaml:"-"` // Calulated internally.
	UseFixture         string    `yaml:"use_fixture,omitempty"`
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

// FixtureReceivers are created by the sequence and are used to receive step instructions.
// Each FixtureReceiver knows which step they belong too and when triggered they start a fade up
// and fade down events which get sent to the launchpad lamps and the DMX fixtures.
func FixtureReceiver(
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

	// Used for static fades, remember the last color.
	var lastColor common.Color

	stopFadeUp := make(chan bool)
	stopFadeDown := make(chan bool)

	// Outer loop wait for configuration.
	for {

		// Wait for first step
		cmd := <-fixtureStepChannel

		// Command for setting fixture copy of last color.
		if cmd.Type == "lastColor" {
			if debug {
				fmt.Printf("Fixture:%d Command Set Last Color to %+v\n", myFixtureNumber, cmd.LastColor)
			}
			lastColor = cmd.LastColor
			continue
		}

		if cmd.SetSwitch && cmd.Type == "switch" {
			if debug {
				fmt.Printf("Fixture:%d Command Set Switch\n", myFixtureNumber)
			}
			lastColor = MapSwitchFixture(cmd.SwitchData, cmd.State, cmd.FadeSpeed, dmxController, fixtures, cmd.Blackout, cmd.Master, cmd.Master, cmd.MasterChanging, lastColor, switchChannels, soundTriggers, soundConfig, dmxInterfacePresent, eventsForLauchpad, guiButtons, fixtureStepChannel)
			continue
		}

		if cmd.Clear || cmd.Blackout {
			if debug {
				fmt.Printf("Fixture:%d Command Clear Fixture\n", myFixtureNumber)
			}
			// Stop any running fade ups.
			select {
			case stopFadeUp <- true:
				fmt.Printf("Send Stop COmmand\n")
			case <-time.After(100 * time.Millisecond):
			}
			turnOffFixtures(cmd, myFixtureNumber, cmd.SequenceNumber, fixtures, dmxController, eventsForLauchpad, guiButtons, dmxInterfacePresent)
			continue
		}

		// If we're a RGB fixture implement the flood and static features.
		if cmd.Type == "rgb" {

			// Start Flood.
			if cmd.StartFlood && cmd.Label != "chaser" {
				if debug {
					fmt.Printf("Fixture:%d Set RGB Flood\n", myFixtureNumber)
				}
				var lamp common.Color
				if cmd.RGBStatic {
					if debug {
						fmt.Printf("%d: Fixture:%d Set RGB Static\n", cmd.SequenceNumber, myFixtureNumber)
					}
					lamp = cmd.RGBStaticColors[myFixtureNumber].Color
				} else {
					lamp = common.White
				}
				lastColor = MapFixtures(false, false, cmd.SequenceNumber, myFixtureNumber, lamp, 0, 0, 0, 0, 0, 0, 0, fixtures, cmd.Blackout, cmd.Master, cmd.Master, 0, cmd.Strobe, cmd.StrobeSpeed, dmxController, dmxInterfacePresent)
				common.LightLamp(common.Button{X: myFixtureNumber, Y: cmd.SequenceNumber}, lamp, cmd.Master, eventsForLauchpad, guiButtons)
				continue
			}

			// Stop Flood.
			if cmd.StopFlood && cmd.Label != "chaser" {
				if debug {
					fmt.Printf("Fixture:%d Set Stop RGB Flood\n", myFixtureNumber)
				}
				lastColor = MapFixtures(false, false, cmd.SequenceNumber, myFixtureNumber, common.Black, 0, 0, 0, 0, 0, 0, 0, fixtures, cmd.Blackout, cmd.Master, cmd.Master, 0, cmd.Strobe, cmd.StrobeSpeed, dmxController, dmxInterfacePresent)
				common.LightLamp(common.Button{X: myFixtureNumber, Y: cmd.SequenceNumber}, common.Black, 0, eventsForLauchpad, guiButtons)
				continue
			}

			// Set Static Scene.
			if cmd.RGBStatic {
				if debug {
					fmt.Printf("%d: Fixture:%d Trying to Set RGB Static\n", cmd.SequenceNumber, myFixtureNumber)
				}
				if cmd.RGBStaticColors[myFixtureNumber].Enabled {

					if debug {
						fmt.Printf("%d: Fixture:%d Set RGB Static Color %+v\n", cmd.SequenceNumber, myFixtureNumber, cmd.RGBStaticColors[myFixtureNumber])
					}

					if cmd.SequenceNumber == 4 {
						cmd.SequenceNumber = 2
					}

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

					lastColor = lightStaticFixture(sequence, myFixtureNumber, eventsForLauchpad, guiButtons, fixtures, true, dmxController, dmxInterfacePresent)
					continue
				}
			}

			if cmd.RGBFadeUpStatic {
				if debug {
					fmt.Printf("1:%d: Fixture:%d State %t Trying to Set RGB Static\n", cmd.SequenceNumber, myFixtureNumber, cmd.RGBStaticColors[myFixtureNumber].Enabled)
				}
				if cmd.RGBStaticColors[myFixtureNumber].Enabled {

					if debug {
						fmt.Printf("%d: Fixture:%d Set RGB Static Color %+v\n", cmd.SequenceNumber, myFixtureNumber, cmd.RGBStaticColors[myFixtureNumber])
					}

					sequence := common.Sequence{}
					sequence.Type = cmd.Type

					if cmd.SequenceNumber == 4 {
						cmd.SequenceNumber = 2
					}

					// Stop any running fade ups.
					select {
					case stopFadeUp <- true:
					case <-time.After(100 * time.Millisecond):
					}

					// Stop any running fade downs.
					select {
					case stopFadeDown <- true:
					case <-time.After(100 * time.Millisecond):
					}

					sequence.Number = cmd.SequenceNumber
					sequence.Master = cmd.Master
					sequence.Blackout = cmd.Blackout
					sequence.Hide = cmd.Hide
					sequence.StaticColors = cmd.RGBStaticColors
					sequence.Static = cmd.RGBStatic
					sequence.StrobeSpeed = cmd.StrobeSpeed
					sequence.Strobe = cmd.Strobe
					sequence.RGBFade = cmd.FadeSpeed
					fadeUpStaticFixture(sequence, myFixtureNumber, stopFadeUp, stopFadeDown, lastColor, eventsForLauchpad, guiButtons, fixtures, true, dmxController, dmxInterfacePresent, fixtureStepChannel)
					continue
				} else {
					// This fixture is disabled, shut it off.
					lastColor = turnOffFixture(myFixtureNumber, cmd.SequenceNumber, common.Black, fixtures, dmxController, dmxInterfacePresent)
				}
			}

			if !cmd.RGBStatic && cmd.RGBPlayStaticOnce && cmd.Label != "chaser" {
				if debug {
					fmt.Printf("Fixture:%d Turn RGB Off\n", myFixtureNumber)
				}
				turnOffFixture(myFixtureNumber, cmd.SequenceNumber, common.EmptyColor, fixtures, dmxController, dmxInterfacePresent)
				common.LightLamp(common.Button{X: myFixtureNumber, Y: cmd.SequenceNumber}, common.Black, cmd.Master, eventsForLauchpad, guiButtons)
				continue
			}

			// Now play all the values for this state.

			// Play out fixture to DMX channels.
			fixture := cmd.RGBPosition.Fixtures[myFixtureNumber]
			if fixture.Enabled {

				// red := fixture.Color.R
				// green := fixture.Color.G
				// blue := fixture.Color.B

				// Integrate cmd.master with fixture.Brightness.
				fixture.Brightness = int((float64(fixture.Brightness) / 100) * (float64(cmd.Master) / 2.55))

				// If we're a shutter chaser flavoured RGB sequence, then disable everything except the brightness.
				if cmd.Label == "chaser" {
					scannerFixturesSequenceNumber := common.GlobalScannerSequenceNumber // Scanner sequence number from config.
					if !cmd.Hide {
						common.LightLamp(common.Button{X: myFixtureNumber, Y: scannerFixturesSequenceNumber}, fixture.Color, fixture.Brightness, eventsForLauchpad, guiButtons)
					}

					// Fixture brightness is sent as master in this case because a shutter chaser is controlling a scanner lamp.
					// and these generally don't have any RGB color channels that can be controlled with brightness.
					// So the only way to make the lamp in the scanner change intensity is to vary the master brightness channel.

					// Lookup chaser lamp color based on the request fixture color.
					// GetColorNameByRGB will return white if the color is not found.
					color := common.GetColorNameByRGB(fixture.Color)

					// Find a suitable gobo based on the requested chaser lamp color.
					scannerGobo := FindGobo(myFixtureNumber, scannerFixturesSequenceNumber, color, fixtures)
					// Find a suitable color wheel setting based on the requested static lamp color.
					scannerColor := FindColor(myFixtureNumber, scannerFixturesSequenceNumber, color, fixtures)

					lastColor = MapFixtures(true, cmd.ScannerChaser, scannerFixturesSequenceNumber, myFixtureNumber, fixture.Color, 0, 0, 0, 0, 0, scannerGobo, scannerColor, fixtures, cmd.Blackout, cmd.Master, fixture.Brightness, cmd.Music, cmd.Strobe, cmd.StrobeSpeed, dmxController, dmxInterfacePresent)
				} else {
					if !cmd.Hide {
						common.LightLamp(common.Button{X: myFixtureNumber, Y: cmd.SequenceNumber}, fixture.Color, cmd.Master, eventsForLauchpad, guiButtons)
					}
					lastColor = MapFixtures(false, cmd.ScannerChaser, cmd.SequenceNumber, myFixtureNumber, fixture.Color, 0, 0, 0, 0, 0, cmd.ScannerGobo, cmd.ScannerColor, fixtures, cmd.Blackout, cmd.Master, cmd.Master, cmd.Music, cmd.Strobe, cmd.StrobeSpeed, dmxController, dmxInterfacePresent)
				}
			}
		}

		// If we are a scanner, implement the scanner movements.
		if cmd.Type == "scanner" {

			// Flood on. Turn on the scanners in flood mode.
			if cmd.StartFlood {
				if debug {
					fmt.Printf("Fixture:%d Scanner Start Flood\n", myFixtureNumber)
				}
				turnOnFixtures(cmd, myFixtureNumber, cmd.SequenceNumber, fixtures, dmxController, eventsForLauchpad, guiButtons, dmxInterfacePresent)
				common.LightLamp(common.Button{X: myFixtureNumber, Y: cmd.SequenceNumber}, common.White, cmd.Master, eventsForLauchpad, guiButtons)
				common.LabelButton(myFixtureNumber, cmd.SequenceNumber, "", guiButtons)
				continue
			}

			// Stop flood.
			if cmd.StopFlood {
				if debug {
					fmt.Printf("Fixture:%d Scanner Start Flood\n", myFixtureNumber)
				}
				common.LightLamp(common.Button{X: myFixtureNumber, Y: cmd.SequenceNumber}, common.Black, common.MIN_DMX_BRIGHTNESS, eventsForLauchpad, guiButtons)
				continue
			}

			// Find the fixture
			fixture := cmd.ScannerPosition.Fixtures[myFixtureNumber]

			// If enabled activate the physical scanner.
			if fixture.Enabled {

				if debug {
					fmt.Printf("Fixture:%d Play Scanner \n", myFixtureNumber)
				}

				// In the case of a scanner, they usually have a shutter and a master dimmer to control the brightness
				// of the lamp. Problem is we can't use the shutter for the control of the overall brightness and the
				// master for the master dimmmer like we do with RGB fixture. The shutter noramlly is more of a switch
				// eg. Open , Closed , Strobe etc. If I want to slow fade through a set of scanners I need to use the
				// brightness for control. Which means I need to combine the master and the control brightness
				// at this stage.
				scannerBrightness := int(math.Round((float64(fixture.Brightness) / 100) * (float64(cmd.Master) / 2.55)))
				// Tell the scanner what to do.
				lastColor = MapFixtures(false, cmd.ScannerChaser, cmd.SequenceNumber, myFixtureNumber, fixture.ScannerColor, fixture.Pan, fixture.Tilt,
					fixture.Shutter, cmd.Rotate, cmd.Program, cmd.ScannerGobo, cmd.ScannerColor, fixtures, cmd.Blackout, cmd.Master, scannerBrightness, cmd.Music, cmd.Strobe, cmd.StrobeSpeed, dmxController, dmxInterfacePresent)

				// Scannner is rotating, work out what to do with the launchpad lamps.
				if !cmd.Hide {
					// Every quater turn, display a color to represent a position in the rotation.
					howOftern := cmd.NumberSteps / 4
					if howOftern != 0 {
						if cmd.Step%howOftern == 0 {
							// We're not in chase mode so use the color generated in the pattern generator.common.
							common.LightLamp(common.Button{X: myFixtureNumber, Y: cmd.SequenceNumber}, fixture.Color, cmd.Master, eventsForLauchpad, guiButtons)
							common.LabelButton(myFixtureNumber, cmd.SequenceNumber, "", guiButtons)
						}
					}
				}
			} else {
				// This scanner is disabled, shut it off.
				turnOffFixture(myFixtureNumber, cmd.SequenceNumber, lastColor, fixtures, dmxController, dmxInterfacePresent)
			}
		}
	}
}

func MapFixturesColorOnly(sequence *common.Sequence, dmxController *ft232.DMXController, fixtures *Fixtures, scannerColor int, dmxInterfacePresent bool) {
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

	if settingName == "" {
		return 0, fmt.Errorf("findChannelSettingByNameAndSpeed: error setting name is empty")
	}
	if settingSpeed == "" {
		return 0, fmt.Errorf("findChannelSettingByNameAndSpeed: error setting speed is empty")
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
							if debug {
								fmt.Printf("findChannelSettingByNameAndSpeed: speed found is %d\n", v)
							}
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
func MapFixtures(chaser bool, hadShutterChase bool,
	mySequenceNumber int,
	displayFixture int,
	color common.Color,
	pan int, tilt int, shutter int, rotate int, program int, selectedGobo int, scannerColor int,
	fixtures *Fixtures, blackout bool, brightness int, master int, music int, strobe bool, strobeSpeed int,
	dmxController *ft232.DMXController, dmxInterfacePresent bool) common.Color {

	// We control the brightness of each color with the brightness value.
	// The overall fixture brightness is set from the master value.
	Red := (float64(color.R) / 100) * (float64(brightness) / 2.55)
	Green := (float64(color.G) / 100) * (float64(brightness) / 2.55)
	Blue := (float64(color.B) / 100) * (float64(brightness) / 2.55)
	White := (float64(color.W) / 100) * (float64(brightness) / 2.55)
	Amber := (float64(color.A) / 100) * (float64(brightness) / 2.55)
	UV := (float64(color.UV) / 100) * (float64(brightness) / 2.55)

	for _, fixture := range fixtures.Fixtures {
		if fixture.Group == mySequenceNumber+1 {
			for channelNumber, channel := range fixture.Channels {

				// Right of the bat if we're blacked out, set the channel to 0 and our work here is done.
				if blackout {
					SetChannel(fixture.Address+int16(channelNumber), byte(0), dmxController, dmxInterfacePresent)
					continue
				}

				// Match the fixture number unless there are mulitple sub fixtures.
				if fixture.Number == displayFixture+1 || fixture.MultiFixtureDevice {
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
							if strobe {
								SetChannel(fixture.Address+int16(channelNumber), byte(strobeSpeed), dmxController, dmxInterfacePresent)
							} else {
								SetChannel(fixture.Address+int16(channelNumber), byte(0), dmxController, dmxInterfacePresent)
							}
						}
						// Master Dimmer.
						if !hadShutterChase {
							if strings.Contains(channel.Name, "Master") || strings.Contains(channel.Name, "Dimmer") {
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
					} else { // We are a scanner chaser, so operate on brightness to master dimmer and scanner color and gobo.
						// Master Dimmer.
						if strings.Contains(channel.Name, "Master") || strings.Contains(channel.Name, "Dimmer") {
							if strings.Contains(channel.Name, "reverse") ||
								strings.Contains(channel.Name, "Reverse") ||
								strings.Contains(channel.Name, "invert") ||
								strings.Contains(channel.Name, "Invert") {
								SetChannel(fixture.Address+int16(channelNumber), byte(reverse_dmx(master)), dmxController, dmxInterfacePresent)
							} else {
								SetChannel(fixture.Address+int16(channelNumber), byte(master), dmxController, dmxInterfacePresent)
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

	return color
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
	fadeSpeed int,
	dmxController *ft232.DMXController,
	fixturesConfig *Fixtures, blackout bool,
	brightness int, master int, masterChanging bool, lastColor common.Color,
	switchChannels []common.SwitchChannel,
	SoundTriggers []*common.Trigger,
	soundConfig *sound.SoundConfig,
	dmxInterfacePresent bool,
	eventsForLauchpad chan common.ALight,
	guiButtons chan common.ALight,
	fixtureStepChannel chan common.FixtureCommand) common.Color {

	var useFixtureLabel string

	if debug {
		fmt.Printf("MapSwitchFixture switchNumber %d, current position %d fade speed %d\n", swiTch.Number, swiTch.CurrentPosition, fadeSpeed)
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
			return lastColor
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
			return lastColor
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
			newMiniSequencer(thisFixture, swiTch, newAction, dmxController, fixturesConfig, switchChannels, soundConfig, blackout, brightness, master, masterChanging, lastColor, dmxInterfacePresent, eventsForLauchpad, guiButtons, fixtureStepChannel)
			if action.Mode != "Static" {
				lastColor = common.EmptyColor
			}
		}

		// If there are no actions, turn off any previos mini sequencers for this switch.
		if len(state.Actions) == 0 {
			newAction := Action{}
			newAction.Name = "Off"
			newAction.Number = 1
			newAction.Mode = "Off"
			lastColor := common.Color{}
			newMiniSequencer(thisFixture, swiTch, newAction, dmxController, fixturesConfig, switchChannels, soundConfig, blackout, brightness, master, masterChanging, lastColor, dmxInterfacePresent, eventsForLauchpad, guiButtons, fixtureStepChannel)
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
	return lastColor
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

func fadeUpStaticFixture(sequence common.Sequence, myFixtureNumber int, StopFadeUp chan bool, StopFadeDown chan bool, lastColor common.Color, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, fixturesConfig *Fixtures, enabled bool, dmxController *ft232.DMXController, dmxInterfacePresent bool, fixtureStepChannel chan common.FixtureCommand) {

	if debug {
		fmt.Printf("fadeUpStaticFixture seq %d fixture %d \n", sequence.Number, myFixtureNumber)
	}

	lamp := sequence.StaticColors[myFixtureNumber]

	// Look for a matching color
	color := common.GetColorNameByRGB(lamp.Color)

	// Find a suitable gobo based on the requested static lamp color.
	scannerGobo := FindGobo(myFixtureNumber, sequence.Number, color, fixturesConfig)
	// Find a suitable color wheel settin based on the requested static lamp color.
	scannerColor := FindColor(myFixtureNumber, sequence.Number, color, fixturesConfig)

	// Now Fade up
	go func() {
		// Soft start
		// Calulate the steps
		fadeUpValues := common.GetFadeValues(64, float64(common.MAX_DMX_BRIGHTNESS), 1, false)
		fadeDownValues := common.GetFadeValues(64, float64(common.MAX_DMX_BRIGHTNESS), 1, true)

		if lastColor != common.EmptyColor {
			for _, fade := range fadeDownValues {
				// Listen for stop command.
				select {
				case <-StopFadeDown:
					return
				case <-time.After(10 * time.Millisecond):
				}
				common.LightLamp(common.Button{X: myFixtureNumber, Y: sequence.Number}, lastColor, fade, eventsForLaunchpad, guiButtons)
				MapFixtures(false, false, sequence.Number, myFixtureNumber, lastColor, common.SCANNER_MID_POINT, common.SCANNER_MID_POINT, 0, 0, 0, scannerGobo, scannerColor, fixturesConfig, sequence.Blackout, fade, sequence.Master, 0, sequence.Strobe, sequence.StrobeSpeed, dmxController, dmxInterfacePresent)

				// Control how long the fade take with the speed control.
				time.Sleep((5 * time.Millisecond) * (time.Duration(sequence.RGBFade)))
			}
			// Fade down complete, set lastColor to empty in the fixture.
			command := common.FixtureCommand{
				Type:      "lastColor",
				LastColor: common.EmptyColor,
			}
			select {
			case fixtureStepChannel <- command:
			case <-time.After(100 * time.Millisecond):
			}
		}

		// Fade up fixture.
		for _, fade := range fadeUpValues {
			// Listen for stop command.
			select {
			case <-StopFadeUp:
				turnOffFixture(myFixtureNumber, sequence.Number, lastColor, fixturesConfig, dmxController, dmxInterfacePresent)
				return
			case <-time.After(10 * time.Millisecond):
			}

			common.LightLamp(common.Button{X: myFixtureNumber, Y: sequence.Number}, lamp.Color, fade, eventsForLaunchpad, guiButtons)
			MapFixtures(false, false, sequence.Number, myFixtureNumber, lamp.Color, common.SCANNER_MID_POINT, common.SCANNER_MID_POINT, 0, 0, 0, scannerGobo, scannerColor, fixturesConfig, sequence.Blackout, fade, sequence.Master, 0, sequence.Strobe, sequence.StrobeSpeed, dmxController, dmxInterfacePresent)

			// Control how long the fade take with the speed control.
			time.Sleep((5 * time.Millisecond) * (time.Duration(sequence.RGBFade)))
		}
		// Fade up complete, set lastColor up in the fixture.
		command := common.FixtureCommand{
			Type:      "lastColor",
			LastColor: lamp.Color,
		}
		select {
		case fixtureStepChannel <- command:
		case <-time.After(100 * time.Millisecond):
		}

	}()

	// Only play once, we don't want to flood the DMX universe with
	// continual commands.
	sequence.PlayStaticOnce = false

}

func lightStaticFixture(sequence common.Sequence, myFixtureNumber int, eventsForLauchpad chan common.ALight, guiButtons chan common.ALight, fixturesConfig *Fixtures, enabled bool, dmxController *ft232.DMXController, dmxInterfacePresent bool) (lastColor common.Color) {

	if debug {
		fmt.Printf("lightStaticFixture seq %d fixture %d \n", sequence.Number, myFixtureNumber)
	}

	lamp := sequence.StaticColors[myFixtureNumber]

	// If we're not hiding the sequence on the launchpad, show the static colors on the buttons.
	if !sequence.Hide {
		if lamp.Flash {
			onColor := common.Color{R: lamp.Color.R, G: lamp.Color.G, B: lamp.Color.B}
			if debug {
				fmt.Printf("FlashLamp Y:%d X:%d\n", sequence.Number, myFixtureNumber)
			}
			common.FlashLight(common.Button{X: myFixtureNumber, Y: sequence.Number}, onColor, common.Black, eventsForLauchpad, guiButtons)
		} else {
			if debug {
				fmt.Printf("LightLamp Y:%d X:%d\n", sequence.Number, myFixtureNumber)
			}
			common.LightLamp(common.Button{X: myFixtureNumber, Y: sequence.Number}, lamp.Color, sequence.Master, eventsForLauchpad, guiButtons)
		}
	}
	if debug {
		fmt.Printf("seq %d fixture %d strobe %t speed %d master %d blackout %t\n", sequence.Number, myFixtureNumber, sequence.Strobe, sequence.StrobeSpeed, sequence.Master, sequence.Blackout)
	}

	if debug {
		fmt.Printf("lightStaticFixtur: Looking for Color seq %d fixture %d color %+v\n", sequence.Number, myFixtureNumber, lamp.Color)
	}

	// Look for a matching color
	color := common.GetColorNameByRGB(lamp.Color)
	if debug {
		fmt.Printf("lightStaticFixture seq %d fixture %d Matching color -> lamp.Color %+v Found Name color %s \n", sequence.Number, myFixtureNumber, lamp.Color, color)
	}

	// Find a suitable gobo based on the requested static lamp color.
	scannerGobo := FindGobo(myFixtureNumber, sequence.Number, color, fixturesConfig)
	// Find a suitable color wheel settin based on the requested static lamp color.
	scannerColor := FindColor(myFixtureNumber, sequence.Number, color, fixturesConfig)

	if debug {
		fmt.Printf("lightStaticFixture seq %d fixture %d Found -> scannerGobo %d scannerColor %d\n", sequence.Number, myFixtureNumber, scannerGobo, scannerColor)
	}

	lastColor = MapFixtures(false, false, sequence.Number, myFixtureNumber, lamp.Color, common.SCANNER_MID_POINT, common.SCANNER_MID_POINT, 0, 0, 0, scannerGobo, scannerColor, fixturesConfig, sequence.Blackout, sequence.Master, sequence.Master, 0, sequence.Strobe, sequence.StrobeSpeed, dmxController, dmxInterfacePresent)

	// Only play once, we don't want to flood the DMX universe with
	// continual commands.
	sequence.PlayStaticOnce = false

	return lastColor
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
func turnOffFixtures(cmd common.FixtureCommand, myFixtureNumber int, mySequenceNumber int, fixtures *Fixtures, dmxController *ft232.DMXController, eventsForLauchpad chan common.ALight, guiButtons chan common.ALight, dmxInterfacePresent bool) (lastColor common.Color) {
	// Only turn off real sequences.
	if mySequenceNumber < 3 {
		if debug {
			fmt.Printf("Sequence %d: Fixture %d turnOffFixtures\n", mySequenceNumber, myFixtureNumber)
		}
		common.LabelButton(myFixtureNumber, mySequenceNumber, "", guiButtons)
		lastColor = MapFixtures(false, false, mySequenceNumber, myFixtureNumber, common.Black, 0, 0, 0, 0, 0, 0, 0, fixtures, cmd.Blackout, cmd.Master, cmd.Master, cmd.Music, cmd.Strobe, cmd.StrobeSpeed, dmxController, dmxInterfacePresent)
	}

	return lastColor
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

func turnOnFixtures(cmd common.FixtureCommand, myFixtureNumber int, mySequenceNumber int, fixtures *Fixtures, dmxController *ft232.DMXController, eventsForLauchpad chan common.ALight, guiButtons chan common.ALight, dmxInterfacePresent bool) common.Color {
	if !cmd.Hide {
		common.LightLamp(common.Button{X: myFixtureNumber, Y: mySequenceNumber}, common.White, common.MAX_DMX_BRIGHTNESS, eventsForLauchpad, guiButtons)
	}
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

	return MapFixtures(false, false, mySequenceNumber, myFixtureNumber, common.White, pan, tilt, shutter, rotate, program, gobo, scannerColor, fixtures, false, brightness, master, music, cmd.Strobe, cmd.StrobeSpeed, dmxController, dmxInterfacePresent)

}

func turnOffFixture(myFixtureNumber int, mySequenceNumber int, lastColor common.Color, fixtures *Fixtures, dmxController *ft232.DMXController, dmxInterfacePresent bool) common.Color {

	if lastColor != common.EmptyColor {
		return lastColor
	}

	blackout := false
	master := 0
	brightness := 0
	strobeSpeed := 0
	strobe := false

	// Find the color number for White.
	scannerColor := FindColor(myFixtureNumber, mySequenceNumber, "White", fixtures)
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
	return MapFixtures(false, false, mySequenceNumber, myFixtureNumber, common.Black, pan, tilt, shutter, rotate, program, gobo, scannerColor, fixtures, blackout, brightness, master, music, strobe, strobeSpeed, dmxController, dmxInterfacePresent)
}
