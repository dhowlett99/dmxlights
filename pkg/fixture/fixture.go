// Copyright (C) 2022,2023,2024,2025 dhowlett99.
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
	"image/color"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"github.com/dhowlett99/dmxlights/pkg/colors"
	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/go-yaml/yaml"
	"github.com/oliread/usbdmx/ft232"
)

const debug = false
const dmxDebug = false

type Fixtures struct {
	Fixtures []Fixture `yaml:"fixtures"`
}

type Groups struct {
	Groups []Group `yaml:"groups"`
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
	Name         string `yaml:"name"`
	Number       int
	Colors       []string `yaml:"colors"`
	Map          string   `yaml:"map"`
	Mode         string   `yaml:"mode"`
	Fade         string   `yaml:"fade"`
	Size         string   `yaml:"size"`
	Speed        string   `yaml:"speed"`
	Rotate       string   `yaml:"rotate"`
	RotateSpeed  string   `yaml:"rotatespeed"`
	Program      string   `yaml:"program"`
	ProgramSpeed string   `yaml:"programspeed"`
	Strobe       string   `yaml:"strobe"`
	StrobeSpeed  string   `yaml:"strobespeed"`
	Gobo         string   `yaml:"gobo"`
	GoboSpeed    string   `yaml:"gobospeed"`
}

type ActionConfig struct {
	Name                string
	Mode                string
	AvailableColors     []color.RGBA // Available colors for this fixture.
	AvailableColorNames []string
	Color               int // The selected color index for this fixture.
	ColorName           string
	Map                 bool
	Fade                int
	NumberSteps         int
	Shutter             int
	Size                int
	SpeedDuration       time.Duration
	Speed               int
	Shift               int
	TriggerState        bool
	Rotate              int
	RotateName          string
	RotateSpeed         int
	Rotatable           bool
	ReverseSpeed        int
	ForwardSpeed        int
	Forward             bool //Clockwise
	Reverse             bool //AntiClockwise
	AutoRotate          bool
	Program             int
	ProgramOptions      []string
	ProgramSpeed        int
	Music               int
	MusicTrigger        bool
	Strobe              bool
	StrobeSpeed         int
	Gobo                int
	GoboSpeed           int
	ScannerColor        int
	AutoGobo            bool
	GoboOptions         []string
	Pan                 int
	Tilt                int
	RotateSensitivity   int
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
	HasRGBChannels     bool      `yaml:"-"` // Calulated internally.
	HasColorChannel    bool      `yaml:"-"` // Calulated internally.
	UseFixture         string    `yaml:"use_fixture,omitempty"`
}

type Group struct {
	Name   string `yaml:"name"`
	Number string `yaml:"number"`
}

type FixtureInfo struct {
	HasRotate          bool
	RotateOptions      []string
	HasRotateSpeed     bool
	RotateSpeedOptions []string
	HasGobo            bool
	HasColorWheel      bool
	HasProgram         bool
	HasProgramSpeed    bool
}

type Setting struct {
	Name          string `yaml:"name"`
	Label         string `yaml:"label,omitempty"`
	Number        int    `yaml:"number"`
	Channel       string `yaml:"channel,omitempty"`
	Value         string `yaml:"value"`
	SelectedValue string `yaml:"selectedvalue"`
}

type Channel struct {
	Number     int16     `yaml:"number"`
	Name       string    `yaml:"name"`
	Value      *int16    `yaml:"value,omitempty"`
	MaxDegrees *int      `yaml:"maxdegrees,omitempty"`
	Offset     *int      `yaml:"offset,omitempty"` // Offset allows you to position the fixture.
	Comment    string    `yaml:"comment,omitempty"`
	Settings   []Setting `yaml:"settings,omitempty"`
	Override   bool      `yaml:"override,omitempty"`
}

// LoadFixturesReader opens the fixtures config file using the io reader passed.
// Returns a pointer to the fixtures config.
// Returns an error.
func LoadFixturesReader(reader fyne.URIReadCloser) (fixtures *Fixtures, err error) {

	if debug {
		fmt.Printf("LoadFixturesReader\n")
	}

	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("dmxlights: error failed to load fixtures: %s", err.Error())
	}

	// Unmarshals the fixtures yaml file into a data struct
	fixtures = &Fixtures{}
	err = yaml.Unmarshal(data, fixtures)
	if err != nil {
		return nil, errors.New("error: unmarshalling file: " + reader.URI().Name() + err.Error())
	}

	if len(fixtures.Fixtures) == 0 {
		return nil, errors.New("error: unmarshalling file: " + reader.URI().Name() + " error: fixtures are empty")
	}

	return fixtures, nil
}

func SaveFixturesWriter(writer fyne.URIWriteCloser, fixtures *Fixtures) error {

	if debug {
		fmt.Printf("SaveFixturesWriter\n")
	}

	// Marshal the fixtures data into a yaml data structure.
	data, err := yaml.Marshal(fixtures)
	if err != nil {
		return errors.New("error: marshalling file: " + writer.URI().Name() + err.Error())
	}

	// Write the fixtures.yaml file.
	_, err = io.WriteString(writer, string(data))
	if err != nil {
		return errors.New("error: writing file: " + writer.URI().Name() + err.Error())
	}

	// Fixtures file saved, no errors.
	return nil
}

// LoadFixtures opens the fixtures config file using the filename passed.
// Returns a pointer to the fixtures config.
// Returns an error.
func LoadFixtures(projectName string) (fixtures *Fixtures, err error) {

	filename := "projects" + "/" + projectName + ".yaml"

	if debug {
		fmt.Printf("LoadFixtures from file %s\n", filename)
	}

	// Open the fixtures yaml file.
	_, err = os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	// Reads the fixtures yaml file.
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Unmarshals the fixtures.yaml file into a data struct
	fixtures = &Fixtures{}
	err = yaml.Unmarshal(data, fixtures)
	if err != nil {
		return nil, errors.New("error: unmarshalling file: " + filename + err.Error())
	}

	if len(fixtures.Fixtures) == 0 {
		return nil, errors.New("error: unmarshalling file: " + filename + " error: fixtures are empty")
	}

	return fixtures, nil
}

// LoadFixtureGroups opens the fixtures group config file using the filename passed.
// Returns a pointer to the fixtures group config.
// Returns an error.
func LoadFixtureGroups(filename string) (groups *Groups, err error) {

	if debug {
		fmt.Printf("LoadFixtures from file %s\n", filename)
	}

	// Open the fixtures yaml file.
	_, err = os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	// Reads the fixtures yaml file.
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Unmarshals the fixtures group file into a data struct.
	groups = &Groups{}
	err = yaml.Unmarshal(data, groups)
	if err != nil {
		return nil, errors.New("error: unmarshalling file: " + filename + err.Error())
	}

	if len(groups.Groups) == 0 {
		return nil, errors.New("error: unmarshalling file: " + filename + " error: groups are empty")
	}

	return groups, nil
}

// SaveFixtures - saves a complete list of fixtures to filename.
// Returns an error.
func SaveFixtures(filename string, fixtures *Fixtures) error {

	if debug {
		fmt.Printf("SaveFixtures\n")
	}

	// Marshal the fixtures data into a yaml data structure.
	data, err := yaml.Marshal(fixtures)
	if err != nil {
		return errors.New("error: marshalling file: " + "projects/" + filename + err.Error())
	}

	// Write the fixtures.yaml file.
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return errors.New("error: writing file: " + "projects/" + filename + err.Error())
	}

	// Fixtures file saved, no errors.
	return nil
}

func AllFixturesOff(sequences []*common.Sequence, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, dmxController *ft232.DMXController, fixturesConfig *Fixtures, dmxInterfacePresent bool) {

	if debug {
		fmt.Printf("AllFixturesOff\n")
	}

	for y := 0; y < len(sequences); y++ {
		if sequences[y].Type != "switch" && sequences[y].Label != "chaser" {
			for x := 0; x < 8; x++ {
				common.LightLamp(common.Button{X: x, Y: y}, colors.Black, common.MAX_DMX_BRIGHTNESS, eventsForLaunchpad, guiButtons)
				MapFixtures(false, false, y, x, colors.Black, colors.Black, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, true, 0, 0, 0, false, 0, dmxController, dmxInterfacePresent)
				common.LabelButton(x, y, "", guiButtons)
			}
		}
	}
}

// Clear fixture.
func clearFixture(fixtureNumber int, cmd common.FixtureCommand, stopFadeDown chan bool, stopFadeUp chan bool, fixtures *Fixtures, dmxController *ft232.DMXController, dmxInterfacePresent bool) common.LastColor {

	if debug {
		fmt.Printf("Fixture:%d clear\n", fixtureNumber)
	}

	// Send stop any running fade ups.
	select {
	case stopFadeUp <- true:
	case <-time.After(100 * time.Millisecond):
	}

	// Send stop any running fade downs.
	select {
	case stopFadeDown <- true:
	case stopFadeDown <- true:
	case <-time.After(100 * time.Millisecond):
	}

	return MapFixtures(false, false, cmd.SequenceNumber, fixtureNumber, colors.Black, colors.Black, 0, 0, 0, 0, 0, 0, cmd.ScannerColor, fixtures, cmd.Blackout, cmd.Master, cmd.Master, cmd.Music, cmd.Strobe, cmd.StrobeSpeed, dmxController, dmxInterfacePresent)
}

func MapFixturesColorOnly(sequenceNumber, selectedFixture, selectedColor int, dmxController *ft232.DMXController, fixtures *Fixtures, dmxInterfacePresent bool) {
	if debug {
		fmt.Printf("MapFixturesColorOnly Sequence %d Fixture %d Gobo %d \n", sequenceNumber, selectedFixture, selectedColor)
	}

	for _, fixture := range fixtures.Fixtures {
		// Match only this fixture.
		if fixture.Group-1 == sequenceNumber {
			// Match only this sequence.
			if fixture.Group-1 == sequenceNumber {

				// Match only this fixture.
				if fixture.Number == selectedFixture+1 {
					for channelNumber, channel := range fixture.Channels {
						if strings.Contains(channel.Name, "Color") {
							for _, setting := range channel.Settings {
								if setting.Number-1 == selectedColor {
									v, _ := strconv.ParseFloat(setting.Value, 32)
									SetChannel(fixture.Address+int16(channelNumber), byte(v), dmxController, dmxInterfacePresent)
								}
							}
						}
					}
				}
			}
		}
	}
}

func findChannelSettingByLabel(fixture *Fixture, channelName string, label string) (int, error) {

	if debug {
		fmt.Printf("findChannelSettingByLabel: looking for Label %s in Channel %s settings\n", label, channelName)
	}

	// Look through the channels.
	for _, channel := range fixture.Channels {
		if debug {
			fmt.Printf("inspect channel %s for %s\n", channel.Name, channelName)
		}
		// Match the channel. So covert to lowercase first and then look if it is the search Name.
		channelName := strings.ToLower(channel.Name)
		searchName := strings.ToLower(channelName)
		if debug {
			fmt.Printf("channelName=%s searchName=%s\n", channelName, searchName)
		}
		if channelName == searchName {
			if debug {
				fmt.Printf("Found a matching name: channel.Settings %+v\n", channel.Settings)
			}

			// Look through the settings.
			for _, setting := range channel.Settings {
				if debug {
					fmt.Printf("inspect setting -> Label %s = label %s\n", setting.Label, label)
				}

				// Match Setting. So covert to lowercase first and then look if it is the search label.
				channelLabel := strings.ToLower(setting.Label)
				searchLabel := strings.ToLower(label)
				if debug {
					fmt.Printf("channelLabel=%s searchLabel=%s\n", channelLabel, searchLabel)
				}
				if channelLabel == searchLabel {
					if debug {
						fmt.Printf("Found a matcing label: Fixture.Name=%s Channel.Name=%s Label=%s Setting.Name %s Setting.Value %s\n", fixture.Name, channel.Name, label, setting.Name, setting.Value)
					}
					v, _ := strconv.Atoi(setting.Value)
					return v, nil
				}
			}
		}
	}

	return 0, fmt.Errorf("findChannelSettingByLabel: label setting \"%s\" not found in channel \"%s\" fixture :%s", label, channelName, fixture.Name)
}

func fixtureHasChannel(fixture *Fixture, channelName string) bool {

	if debug {
		fmt.Printf("fixtureHasChannel for fixture %s channel %s\n", fixture.Name, channelName)
	}

	for _, channel := range fixture.Channels {
		if channel.Name == channelName {
			return true
		}
	}

	return false
}

func MapFixturesGoboOnly(sequenceNumber, selectedFixture, selectedGobo int, fixtures *Fixtures, dmxController *ft232.DMXController,
	dmxInterfacePresent bool) {

	if debug {
		fmt.Printf("MapFixturesGoboOnly Sequence %d Fixture %d Gobo %d \n", sequenceNumber, selectedFixture, selectedGobo)
	}

	for _, fixture := range fixtures.Fixtures {
		// Match only this sequence.
		if fixture.Group-1 == sequenceNumber {
			// Match only this fixture.
			if fixture.Number == selectedFixture+1 {
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
}

func calcFinalValueBasedOnConfigAndSettingValue(configValue string, settingValue int) (final int) {

	if debug {
		fmt.Printf("calcFinalValueBasedOnConfigAndSettingValue\n")
	}

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

func IsNumericOnly(str string) bool {

	if debug {
		fmt.Printf("IsNumericOnly\n")
	}

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

func reverse_dmx(n int) int {

	if debug {
		fmt.Printf("reverse_dmx: Reverse in is %d\n", n)
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

// limitDmxValue - calculates the maximum DMX value based on the number of degrees the fixtire can achieve.
func limitDmxValue(MaxDegrees *int, Value int) int {

	if debug {
		fmt.Printf("limitDmxValue\n")
	}

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

// Send a command to all the fixtures.
func SendToAllFixtures(fixtureChannels []chan common.FixtureCommand, command common.FixtureCommand) {
	for _, fixture := range fixtureChannels {
		fixture <- command
	}
}

// Automatically set the number of sub fixtures inside a fixture.
func SetMultiFixtureFlag(fixturesConfig *Fixtures) {

	for fixtureNumber, fixture := range fixturesConfig.Fixtures {
		// Automatically set the number of sub fixtures inside a fixture.
		var numberSubFixtures int
		for _, channel := range fixture.Channels {
			if strings.Contains(channel.Name, "Red") {
				numberSubFixtures++
			}
		}
		if numberSubFixtures > 1 {
			if debug {
				fmt.Printf("\t fixture %s numberSubFixtures %d\n", fixture.Name, numberSubFixtures)
			}
			fixturesConfig.Fixtures[fixtureNumber].MultiFixtureDevice = true
			fixturesConfig.Fixtures[fixtureNumber].NumberSubFixtures = numberSubFixtures
		}
	}
}

// Does the fixture have red, green and blue channels.
func SetHasRGBFlag(fixturesConfig *Fixtures) {

	for fixtureNumber, fixture := range fixturesConfig.Fixtures {

		var hasRed bool
		var hasGreen bool
		var hasBlue bool
		var hasColor bool

		for _, channel := range fixture.Channels {
			if strings.Contains(channel.Name, "Red") || strings.Contains(channel.Name, "red") {
				hasRed = true
			}
			if strings.Contains(channel.Name, "Green") || strings.Contains(channel.Name, "green") {
				hasGreen = true
			}
			if strings.Contains(channel.Name, "Blue") || strings.Contains(channel.Name, "blue") {
				hasBlue = true
			}
			if strings.Contains(channel.Name, "Color") || strings.Contains(channel.Name, "color") {
				hasColor = true
			}
		}

		// Now we have looked at every channel.
		if hasRed && hasGreen && hasBlue {
			fixturesConfig.Fixtures[fixtureNumber].HasRGBChannels = true
		}
		if hasColor {
			fixturesConfig.Fixtures[fixtureNumber].HasColorChannel = true
		}

	}
}
