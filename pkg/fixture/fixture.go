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
	Name              string
	Colors            []color.RGBA // Available colors for this fixture.
	Color             int          // The selected color index for this fixture.
	Map               bool
	Fade              int
	NumberSteps       int
	Size              int
	SpeedDuration     time.Duration
	Speed             int
	Shift             int
	TriggerState      bool
	RotateName        string
	RotateNumber      int
	RotateSpeed       int
	Rotatable         bool
	ReverseSpeed      int
	ForwardSpeed      int
	Forward           bool //Clockwise
	Reverse           bool //AntiClockwise
	AutoRotate        bool
	Program           int
	ProgramOptions    []string
	ProgramSpeed      int
	Music             int
	MusicTrigger      bool
	Strobe            bool
	StrobeSpeed       int
	Gobo              int
	GoboSpeed         int
	AutoGobo          bool
	GoboOptions       []string
	RotateSensitivity int
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

// HowManyGobosForThisFixture takes the fixture number, sequence number and the fixturesConfig
// Returns returns the number of gobos this fixture has.
func HowManyGobosForThisFixture(myFixtureNumber int, mySequenceNumber int, fixtures *Fixtures) int {

	if debug {
		fmt.Printf("HowManyGobosForThisFixture\n")
	}

	for _, fixture := range fixtures.Fixtures {
		if fixture.Group == mySequenceNumber+1 {
			if fixture.Number == myFixtureNumber+1 {
				for _, channel := range fixture.Channels {
					if strings.Contains(channel.Name, "Gobo") {
						return len(channel.Settings)
					}
				}
			}
		}
	}
	return 0
}

func FindColorinFixture(fixture *Fixture, color string) int {

	if debug {
		fmt.Printf("FindColor looking for %s in fixture %s\n", color, fixture.Name)
	}

	for _, channel := range fixture.Channels {
		if strings.Contains(channel.Name, "Color") {
			for _, setting := range channel.Settings {
				if setting.Name == color {
					if debug {
						fmt.Printf("Found setting number %d\n", setting.Number)
					}
					return setting.Number
				}
			}
		}
	}
	if debug {
		fmt.Printf("Not FOund setting number returning 0\n")
	}
	return 0
}

// FindColor takes the name of a color channel setting like "White" and returns the color number for this type of scanner.
func FindColor(myFixtureNumber int, mySequenceNumber int, color string, fixtures *Fixtures) int {

	if debug {
		fmt.Printf("FindColor looking for %s seq %d fixture %d\n", color, mySequenceNumber, myFixtureNumber)
	}

	for _, fixture := range fixtures.Fixtures {
		if fixture.Group == mySequenceNumber+1 {
			if fixture.Number == myFixtureNumber+1 {
				for _, channel := range fixture.Channels {
					if strings.Contains(channel.Name, "Color") {
						for _, setting := range channel.Settings {
							if setting.Name == color {
								if debug {
									fmt.Printf("Found setting number %d\n", setting.Number)
								}
								return setting.Number
							}
						}
					}
				}
			}
		}
	}
	if debug {
		fmt.Printf("Not FOund setting number returning 0\n")
	}
	return 0
}

func FindChannelNumberByName(fixture *Fixture, channelName string) (int, error) {

	if debug {
		fmt.Printf("FindChannelNumberByName channelName %s\n", channelName)
	}

	for channelNumber, channel := range fixture.Channels {
		if strings.Contains(channel.Name, channelName) {
			return channelNumber, nil
		}
	}
	return 0, fmt.Errorf("channel %s not found in fixture %s", channelName, fixture.Name)
}

func FindFixtureInfo(thisFixture *Fixture) FixtureInfo {
	if debug {
		fmt.Printf("FindFixtureInfo\n")
	}

	fixtureInfo := FixtureInfo{}

	if thisFixture == nil {
		fmt.Printf("FindFixtureInfo: fixture is empty\n")
		return fixtureInfo
	}

	fixtureInfo.HasRotate = isThisAChannel(thisFixture, "Rotate")
	fixtureInfo.HasRotateSpeed = isThisAChannel(thisFixture, "RotateSpeed")

	// Find all the options for the channel called "Rotate".But only if we have a Rotate Channel exists.
	if fixtureInfo.HasRotate {
		availableRotateOptions := getOptionsForAChannel(*thisFixture, "Rotate")
		// Add the auto option for rotate
		var autoFound bool
		for _, option := range availableRotateOptions {
			if strings.Contains(option, "Auto") || strings.Contains(option, "auto") {
				autoFound = true
			}
			fixtureInfo.RotateOptions = append(fixtureInfo.RotateOptions, option)
		}
		// Now if we didn't find a dedicated channel for automatically rotating in different directions.
		// Add our internal keyword for Auto.
		if !autoFound {
			fixtureInfo.RotateOptions = append(fixtureInfo.RotateOptions, "Auto")
		}
	}

	fixtureInfo.RotateSpeedOptions = []string{"Slow", "Medium", "Fast"}

	fixtureInfo.HasColorWheel = isThisAChannel(thisFixture, "Color")
	fixtureInfo.HasGobo = isThisAChannel(thisFixture, "Gobo")
	fixtureInfo.HasProgram = isThisAChannel(thisFixture, "Program")
	fixtureInfo.HasProgramSpeed = isThisAChannel(thisFixture, "ProgramSpeed")
	return fixtureInfo
}

func getOptionsForAChannel(thisFixture Fixture, channelName string) []string {

	if debug {
		fmt.Printf("getOptionsForAChannel\n")
	}

	var options []string

	for _, channel := range thisFixture.Channels {
		if channel.Name == channelName {
			for _, setting := range channel.Settings {
				options = append(options, setting.Name)
			}
		}
	}
	return options
}

func isThisAChannel(thisFixture *Fixture, channelName string) bool {

	if thisFixture == nil {
		return false
	}

	for _, channel := range thisFixture.Channels {
		if channel.Name == channelName {
			if debug {
				fmt.Printf("\tisThisAChannel fixture %s channelName %s true\n", thisFixture.Name, channelName)
			}
			return true
		}
	}
	return false
}

// returns true is they are the same.
func CheckFixturesAreTheSame(fixtures *Fixtures, startConfig *Fixtures) (bool, string) {

	if len(fixtures.Fixtures) != len(startConfig.Fixtures) {
		return false, "Number of fixtures are different"
	}

	for fixtureNumber, fixture := range fixtures.Fixtures {

		if debug {
			fmt.Printf("Checking Fixture %s against %s\n", fixture.Name, startConfig.Fixtures[fixtureNumber].Name)
		}

		if fixture.Name != startConfig.Fixtures[fixtureNumber].Name {
			return false, fmt.Sprintf("Fixture:%d Name is different\n", fixtureNumber+1)
		}

		if fixture.ID != startConfig.Fixtures[fixtureNumber].ID {
			return false, fmt.Sprintf("Fixture:%d ID is different\n", fixtureNumber+1)
		}

		if fixture.Label != startConfig.Fixtures[fixtureNumber].Label {
			return false, fmt.Sprintf("Fixture:%d Label is different\n", fixtureNumber+1)
		}

		if fixture.Number != startConfig.Fixtures[fixtureNumber].Number {
			return false, fmt.Sprintf("Fixture:%d Number is different\n", fixtureNumber+1)
		}

		if fixture.Description != startConfig.Fixtures[fixtureNumber].Description {
			return false, fmt.Sprintf("Fixture:%d Description is different\n", fixtureNumber+1)
		}

		if fixture.Type != startConfig.Fixtures[fixtureNumber].Type {
			return false, fmt.Sprintf("Fixture:%d Type is different\n", fixtureNumber+1)
		}

		if fixture.Group != startConfig.Fixtures[fixtureNumber].Group {
			return false, fmt.Sprintf("Fixture:%d Group is different\n", fixtureNumber+1)
		}

		if fixture.Address != startConfig.Fixtures[fixtureNumber].Address {
			return false, fmt.Sprintf("Fixture:%d Address is different\n", fixtureNumber+1)
		}

		for channelNumber, channel := range fixture.Channels {

			if channel.Number != startConfig.Fixtures[fixtureNumber].Channels[channelNumber].Number {
				return false, fmt.Sprintf("Fixture:%d Channel Number is different\n", fixtureNumber+1)
			}

			if channel.Name != startConfig.Fixtures[fixtureNumber].Channels[channelNumber].Name {
				return false, fmt.Sprintf("Fixture:%d Channel Name is different\n", fixtureNumber+1)
			}

			if channel.Value != startConfig.Fixtures[fixtureNumber].Channels[channelNumber].Value {
				return false, fmt.Sprintf("Fixture:%d Channel Value is different\n", fixtureNumber+1)
			}

			if channel.MaxDegrees != startConfig.Fixtures[fixtureNumber].Channels[channelNumber].MaxDegrees {
				return false, fmt.Sprintf("Fixture:%d Channel MaxDegrees is different\n", fixtureNumber+1)
			}

			if channel.Offset != startConfig.Fixtures[fixtureNumber].Channels[channelNumber].Offset {
				return false, fmt.Sprintf("Fixture:%d Channel Offset is different\n", fixtureNumber+1)
			}

			if channel.Comment != startConfig.Fixtures[fixtureNumber].Channels[channelNumber].Comment {
				return false, fmt.Sprintf("Fixture:%d Channel Comment is different\n", fixtureNumber+1)
			}

			for settingNumber, setting := range channel.Settings {

				if setting.Name != startConfig.Fixtures[fixtureNumber].Channels[channelNumber].Settings[settingNumber].Name {
					return false, fmt.Sprintf("Fixture:%d Channel Settings Name is different\n", fixtureNumber+1)
				}

				if setting.Label != startConfig.Fixtures[fixtureNumber].Channels[channelNumber].Settings[settingNumber].Label {
					return false, fmt.Sprintf("Fixture:%d Channel Settings Label is different\n", fixtureNumber+1)
				}

				if setting.Number != startConfig.Fixtures[fixtureNumber].Channels[channelNumber].Settings[settingNumber].Number {
					return false, fmt.Sprintf("Fixture:%d Channel Settings Number is different\n", fixtureNumber+1)
				}

				if setting.Channel != startConfig.Fixtures[fixtureNumber].Channels[channelNumber].Settings[settingNumber].Channel {
					return false, fmt.Sprintf("Fixture:%d Channel Channel Number is different\n", fixtureNumber+1)
				}

				if setting.Value != startConfig.Fixtures[fixtureNumber].Channels[channelNumber].Settings[settingNumber].Value {
					return false, fmt.Sprintf("Fixture:%d Channel Value Number is different\n", fixtureNumber+1)
				}
			}

			for stateNumber, state := range fixture.States {

				if state.Number != startConfig.Fixtures[fixtureNumber].States[stateNumber].Number {
					return false, fmt.Sprintf("Fixture:%d State Number is different\n", fixtureNumber+1)
				}

				if state.Name != startConfig.Fixtures[fixtureNumber].States[stateNumber].Name {
					return false, fmt.Sprintf("Fixture:%d State Name is different\n", fixtureNumber+1)
				}

				if state.Label != startConfig.Fixtures[fixtureNumber].States[stateNumber].Label {
					return false, fmt.Sprintf("Fixture:%d State Label is different\n", fixtureNumber+1)
				}

				if state.ButtonColor != startConfig.Fixtures[fixtureNumber].States[stateNumber].ButtonColor {
					return false, fmt.Sprintf("Fixture:%d State ButtonColor is different\n", fixtureNumber+1)
				}

				if state.Master != startConfig.Fixtures[fixtureNumber].States[stateNumber].Master {
					return false, fmt.Sprintf("Fixture:%d State Master is different\n", fixtureNumber+1)
				}

				for actionNumber, action := range state.Actions {

					if action.Name != startConfig.Fixtures[fixtureNumber].States[stateNumber].Actions[actionNumber].Name {
						return false, fmt.Sprintf("Fixture:%d State Action Name is different\n", fixtureNumber+1)
					}

					if action.Number != startConfig.Fixtures[fixtureNumber].States[stateNumber].Actions[actionNumber].Number {
						return false, fmt.Sprintf("Fixture:%d State Action Number is different\n", fixtureNumber+1)
					}

					for colorNumber, color := range action.Colors {
						if color != startConfig.Fixtures[fixtureNumber].States[stateNumber].Actions[actionNumber].Colors[colorNumber] {

							return false, fmt.Sprintf("Fixture:%d State Action Color Number is different\n", fixtureNumber+1)
						}
					}

					if action.Mode != startConfig.Fixtures[fixtureNumber].States[stateNumber].Actions[actionNumber].Mode {
						return false, fmt.Sprintf("Fixture:%d State Action Mode is different\n", fixtureNumber+1)
					}

					if action.Fade != startConfig.Fixtures[fixtureNumber].States[stateNumber].Actions[actionNumber].Fade {
						return false, fmt.Sprintf("Fixture:%d State Action Fade is different\n", fixtureNumber+1)
					}

					if action.Size != startConfig.Fixtures[fixtureNumber].States[stateNumber].Actions[actionNumber].Size {
						return false, fmt.Sprintf("Fixture:%d State Action Size is different\n", fixtureNumber+1)
					}

					if action.Speed != startConfig.Fixtures[fixtureNumber].States[stateNumber].Actions[actionNumber].Speed {
						return false, fmt.Sprintf("Fixture:%d State Action Speed is different\n", fixtureNumber+1)
					}

					if action.Rotate != startConfig.Fixtures[fixtureNumber].States[stateNumber].Actions[actionNumber].Rotate {
						return false, fmt.Sprintf("Fixture:%d State Action Rotate is different\n", fixtureNumber+1)
					}

					if action.RotateSpeed != startConfig.Fixtures[fixtureNumber].States[stateNumber].Actions[actionNumber].RotateSpeed {
						return false, fmt.Sprintf("Fixture:%d State Action RotateSpeed is different\n", fixtureNumber+1)
					}

					if action.Program != startConfig.Fixtures[fixtureNumber].States[stateNumber].Actions[actionNumber].Program {
						return false, fmt.Sprintf("Fixture:%d State Action Program is different\n", fixtureNumber+1)
					}

					if action.Strobe != startConfig.Fixtures[fixtureNumber].States[stateNumber].Actions[actionNumber].Strobe {
						return false, fmt.Sprintf("Fixture:%d State Action Strobe is different\n", fixtureNumber+1)
					}

				}

				for settingNumber, setting := range state.Settings {

					if setting.Name != startConfig.Fixtures[fixtureNumber].States[stateNumber].Settings[settingNumber].Name {
						return false, fmt.Sprintf("Fixture:%d Channel Settings Strobe is different\n", fixtureNumber+1)
					}

					if setting.Label != startConfig.Fixtures[fixtureNumber].States[stateNumber].Settings[settingNumber].Label {
						return false, fmt.Sprintf("Fixture:%d Channel Settings Label is different\n", fixtureNumber+1)
					}

					if setting.Number != startConfig.Fixtures[fixtureNumber].States[stateNumber].Settings[settingNumber].Number {
						return false, fmt.Sprintf("Fixture:%d Channel Settings Number is different\n", fixtureNumber+1)
					}

					if setting.Channel != startConfig.Fixtures[fixtureNumber].States[stateNumber].Settings[settingNumber].Channel {
						return false, fmt.Sprintf("Fixture:%d Channel Channel Number is different\n", fixtureNumber+1)
					}

					if setting.Value != startConfig.Fixtures[fixtureNumber].States[stateNumber].Settings[settingNumber].Value {
						return false, fmt.Sprintf("Fixture:%d Channel Value Number is different\n", fixtureNumber+1)
					}
				}

				if state.Flash != startConfig.Fixtures[fixtureNumber].States[stateNumber].Flash {
					return false, fmt.Sprintf("Fixture:%d State Flash is different\n", fixtureNumber+1)
				}

			}

			if fixture.MultiFixtureDevice != startConfig.Fixtures[fixtureNumber].MultiFixtureDevice {
				return false, fmt.Sprintf("Fixture:%d MultiFixtureDevice is different\n", fixtureNumber+1)
			}

			if fixture.NumberSubFixtures != startConfig.Fixtures[fixtureNumber].NumberSubFixtures {
				return false, fmt.Sprintf("Fixture:%d NumberSubFixtures is different\n", fixtureNumber+1)
			}

			if fixture.UseFixture != startConfig.Fixtures[fixtureNumber].UseFixture {
				return false, fmt.Sprintf("Fixture: %d UseFixture is different\n", fixtureNumber+1)
			}

		}
	}

	return true, ""
}

func GetSwitchFixtureType(switchNumber int, stateNumber int16, fixturesConfig *Fixtures) string {
	for _, fixture := range fixturesConfig.Fixtures {
		if fixture.Type == "switch" {
			if fixture.Number == switchNumber+1 {

				useFixture, _ := GetFixtureDetailsByLabel(fixture.UseFixture, fixturesConfig)
				if debug {
					fmt.Printf("fixture number %d name %s use fixture %s type %s\n", fixture.Number, fixture.Name, fixture.UseFixture, useFixture.Type)
				}
				return useFixture.Type
			}
		}
	}
	return ""
}

func GetSwitchStateIsMusicTriggerOn(switchNumber int, stateNumber int16, fixturesConfig *Fixtures) bool {

	for _, fixture := range fixturesConfig.Fixtures {
		if fixture.Type == "switch" {
			if fixture.Number == switchNumber+1 {
				if debug {
					fmt.Printf("fixture number %d %s\n", fixture.Number, fixture.Name)
				}
				for _, state := range fixture.States {
					if state.Number == stateNumber {
						if debug {
							fmt.Printf("state number %d %+v\n", stateNumber, state.Actions)
						}
						if state.Actions != nil {
							for _, action := range state.Actions {
								if action.Mode == "Chase" && action.Speed != "Music" {
									return true
								}
							}
						}
					}
				}
			}
		}
	}
	return false
}

func GetSwitchAction(switchNumber int, switchState int16, fixturesConfig *Fixtures) Action {

	for _, fixture := range fixturesConfig.Fixtures {
		if fixture.Type == "switch" {
			if fixture.Number == switchNumber {
				if debug {
					fmt.Printf("found fixture number %d name %s type %s\n", fixture.Number, fixture.Name, fixture.Type)
				}
				for _, state := range fixture.States {
					if state.Number == switchState {
						if debug {
							fmt.Printf("looking for state %d have state number %d Actions %+v\n", state.Number, state.Number, state.Actions)
						}
						var action Action
						var actionNumber int
						if state.Actions != nil {
							for actionNumber, action = range state.Actions {
								if action.Mode == "Control" {
									if action.Name == "Off" {
										action.Colors = []string{"Green"}
									}
									if action.Name == "On" {
										action.Colors = []string{"Red"}
									}
								}
							}
							if debug {
								fmt.Printf("Actions:- Mode %s action number %d name %s colors %+v\n", action.Mode, actionNumber, action.Name, action.Colors)
							}
							return action
						}

						if state.Settings != nil {
							if debug {
								for _, setting := range state.Settings {
									fmt.Printf("setting Number %d Label %s Channel %s Valuue %s\n", setting.Number, setting.Label, setting.Channel, setting.Value)
								}
							}
							action := convertSettingToAction(fixture, state.Settings)
							return action
						}
					}
				}
			}
		}
	}
	return Action{Name: "Not Found"}
}

// Given the fixture and the list of settings for this state
// buill a new action that represents the set of settings.
func convertSettingToAction(fixture Fixture, settings []Setting) Action {

	newAction := Action{}

	newAction.Mode = "Setting"
	newAction.Name = "Setting"
	newAction.Number = 1

	// Look through settings and buuld up the new action.
	for _, setting := range settings {

		if debug {
			fmt.Printf("convertSettingToAction: Fixture name %s setting name %s label %s Channel %s name %s value %s\n", fixture.Name, setting.Name, setting.Label, setting.Channel, setting.Name, setting.Value)
		}

		if setting.Channel == "Speed" {
			newAction.RotateSpeed = setting.Value
		}
		if setting.Channel == "Fade" {
			newAction.Fade = setting.Value
		}

		if setting.Channel == "Size" {
			newAction.Fade = setting.Value
		}

		if setting.Channel == "Rotate" {
			newAction.Rotate = setting.Value
		}

		if setting.Channel == "RotateSpeed" {
			newAction.RotateSpeed = setting.Value
		}

		if setting.Channel == "Program" {
			newAction.Program = setting.Value
		}

		if setting.Channel == "ProgramSpeed" {
			newAction.ProgramSpeed = setting.Value
		}

		// A channel setting can only contain one value
		// so only one color.
		if setting.Channel == "Color" {
			// If a setting has a channel name which is a number we lookup that color name.
			if IsNumericOnly(setting.Name) {
				if colorNumber, err := strconv.Atoi(setting.Value); err == nil {
					// Lookup color number in list of available colors.
					colorName := GetColorNameByNumber(&fixture, colorNumber)
					newAction.Colors = []string{colorName}
				}
			} else {
				// we use that string as the color.
				newAction.Colors = []string{setting.Name}
			}
			if setting.Name == "Off" {
				newAction.Colors = []string{"Green"}
			}
			if setting.Name == "On" {
				newAction.Colors = []string{"Red"}
			}
		}

		if setting.Channel == "Strobe" {
			newAction.Strobe = setting.Value
		}

		if setting.Channel == "StrobeSpeed" {
			newAction.StrobeSpeed = setting.Value
		}

		if setting.Channel == "Gobo" {
			newAction.Gobo = setting.Name
		}

		if setting.Channel == "GoboSpeed" {
			newAction.GoboSpeed = setting.Value
		}

	}
	return newAction
}

// GetAvailableScannerGobos - populates a map indexed by fixture number for the sequenceNumber provided.
// Each fixture contains an array of StaticColorButtons, essentially info representing each gobo in this fixture.
// Gobo details provided are - Name, label, number. DMX value and color.
func GetAvailableScannerGobos(sequenceNumber int, fixtures *Fixtures) map[int][]common.StaticColorButton {
	if debug {
		fmt.Printf("getAvailableScannerGobos\n")
	}

	gobos := make(map[int][]common.StaticColorButton)

	for _, f := range fixtures.Fixtures {
		if debug {
			fmt.Printf("Fixture Name:%s\n", f.Name)
		}
		if f.Type == "scanner" {

			if debug {
				fmt.Printf("Sequence: %d - Scanner Name: %s Description: %s\n", sequenceNumber, f.Name, f.Description)
			}
			for _, channel := range f.Channels {
				if channel.Name == "Gobo" {
					newGobo := common.StaticColorButton{}
					for _, setting := range channel.Settings {
						newGobo.Name = setting.Name
						newGobo.Label = setting.Label
						newGobo.Number = setting.Number
						v, _ := strconv.Atoi(setting.Value)
						newGobo.Setting = v
						newGobo.Color = colors.Yellow
						gobos[f.Number] = append(gobos[f.Number], newGobo)
						if debug {
							fmt.Printf("\tGobo: %s Setting: %s\n", setting.Name, setting.Value)
						}
					}
				}
			}
		}
	}
	return gobos
}

// getAvailableScannerColors looks through the fixtures list and finds scanners that
// have colors defined in their config. It then returns an array of these available colors.
// Also returns a map of the default values for each scanner that has colors.
func GetAvailableScannerColors(fixtures *Fixtures) (map[int][]common.StaticColorButton, map[int]int) {

	if debug {
		fmt.Printf("GetAvailableScannerColors for fixture\n")
	}
	scannerColors := make(map[int]int)

	availableScannerColors := make(map[int][]common.StaticColorButton)
	for _, fixture := range fixtures.Fixtures {
		if fixture.Type == "scanner" {
			for _, channel := range fixture.Channels {
				if strings.Contains(channel.Name, "Color") {
					for _, setting := range channel.Settings {
						newStaticColorButton := common.StaticColorButton{}
						newStaticColorButton.SelectedColor = setting.Number
						settingColor, err := common.GetRGBColorByName(setting.Name)
						if err != nil {
							fmt.Printf("error: %s\n", err)
							continue
						}
						newStaticColorButton.Color = settingColor
						availableScannerColors[fixture.Number] = append(availableScannerColors[fixture.Number], newStaticColorButton)
						scannerColors[fixture.Number-1] = 0
					}
				}
			}
		}
	}
	return availableScannerColors, scannerColors
}

// GetScannerColorName finds the color for given scanner and color number.
func GetScannerColorName(scannerNumber int, colorNumber int, fixtures *Fixtures) (color.RGBA, error) {

	if debug {
		fmt.Printf("GetScannerColorName() Looking for Color Number %d\n", colorNumber)
	}

	for _, fixture := range fixtures.Fixtures {
		if fixture.Type == "scanner" {
			if fixture.Number == scannerNumber+1 {
				for _, channel := range fixture.Channels {
					if strings.Contains(channel.Name, "Color") {
						for _, setting := range channel.Settings {
							if setting.Number == colorNumber+1 {
								settingColor, err := common.GetRGBColorByName(setting.Name)
								if err != nil {
									fmt.Printf("error: %s\n", err)
									continue
								}
								return settingColor, nil
							}
						}
					}
				}
			}
		}
	}
	return color.RGBA{}, fmt.Errorf("color not found")
}

func HowManyScannerColors(sequence *common.Sequence, fixturesConfig *Fixtures) []color.RGBA {

	if debug {
		fmt.Printf("HowManyScannerColors: \n")
	}

	// Clear out sequemce colors.
	sequence.SequenceColors = []color.RGBA{}

	for scannerNumber := 0; scannerNumber < sequence.NumberFixtures; scannerNumber++ {
		// Look at all the scannes and add their selected color to the color display.
		colorNumber := sequence.ScannerColor[scannerNumber]
		// Get the color name from the fixture config, ignore scanner that don't have a color set.
		color, err := GetScannerColorName(scannerNumber, colorNumber, fixturesConfig)
		if err == nil {
			if debug {
				fmt.Printf("HowManyScannerColors()Scanner %d is Color %s\n", scannerNumber, common.GetColorNameByRGB(color))
			}
			sequence.SequenceColors = append(sequence.SequenceColors, color)
		}
	}
	if debug {
		fmt.Printf("HowManyScannerColors() colors %+v\n", sequence.SequenceColors)
	}

	return sequence.SequenceColors
}

// Send a command to all the fixtures.
func SendToAllFixtures(fixtureChannels []chan common.FixtureCommand, command common.FixtureCommand) {
	for _, fixture := range fixtureChannels {
		fixture <- command
	}
}

func HowManyFixturesInGroup(sequenceNumber int, fixturesConfig *Fixtures) int {

	var recents []int

	if debug {
		fmt.Printf("\nHowManyFixturesInGroup for sequence %d\n", sequenceNumber)
	}

	var count int
	for _, fixture := range fixturesConfig.Fixtures {

		// Found the group.
		if fixture.Group == sequenceNumber+1 {

			if debug {
				fmt.Printf("\t%d: Found fixture in group %d\n", sequenceNumber, fixture.Number)
			}

			// Have we seen this fixture number already
			if !haveWeSeenThisBefore(recents, fixture.Number) {

				// If this is a multifixture device
				if fixture.MultiFixtureDevice {

					if debug {
						fmt.Printf("\t\t%d: Found MultiFixtureDevice %d\n", sequenceNumber, fixture.Number)
					}

					// Then count the number of RGB channels.
					for _, channel := range fixture.Channels {
						if strings.Contains(channel.Name, "Red") {
							if debug {
								fmt.Printf("\t\t\t%d: Found Red Channel %d\n", sequenceNumber, fixture.Number)
							}
							count++
						}
					}
					break
				}

				// Only count this one if we haven't counted it already.
				if debug {
					fmt.Printf("\t\tAdd fixture in recents %d\n", fixture.Number)
				}
				recents = append(recents, fixture.Number)
				count++
			}
		}
	}

	if debug {
		fmt.Printf("%d: Found %d fixtures\n", sequenceNumber, count)
	}
	return count
}

func haveWeSeenThisBefore(recents []int, fixtureNumber int) bool {

	for _, recent := range recents {
		if fixtureNumber == recent {
			return true
		}

	}
	return false
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

		for _, channel := range fixture.Channels {
			if strings.Contains(channel.Name, "Red") {
				hasRed = true
			}
			if strings.Contains(channel.Name, "Green") {
				hasGreen = true
			}
			if strings.Contains(channel.Name, "Blue") {
				hasBlue = true
			}
		}

		// Now we have looked at every channel.
		if hasRed && hasGreen && hasBlue {
			fixturesConfig.Fixtures[fixtureNumber].HasRGBChannels = true
		}

	}
}
