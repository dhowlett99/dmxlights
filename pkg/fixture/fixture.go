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
	"github.com/dhowlett99/dmxlights/pkg/sound"
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
	Colors            []color.RGBA // Colos available for this fixture.
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
func LoadFixtures(filename string) (fixtures *Fixtures, err error) {

	if debug {
		fmt.Printf("LoadFixtures from file %s\n", "projects/"+filename)
	}

	// Open the fixtures yaml file.
	_, err = os.OpenFile("projects/"+filename, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	// Reads the fixtures yaml file.
	data, err := os.ReadFile("projects/" + filename)
	if err != nil {
		return nil, err
	}

	// Unmarshals the fixtures.yaml file into a data struct
	fixtures = &Fixtures{}
	err = yaml.Unmarshal(data, fixtures)
	if err != nil {
		return nil, errors.New("error: unmarshalling file: " + "projects/" + filename + err.Error())
	}

	if len(fixtures.Fixtures) == 0 {
		return nil, errors.New("error: unmarshalling file: " + "projects/" + filename + " error: fixtures are empty")
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
				MapFixtures(false, false, y, x, colors.Black, 0, 0, 0, 0, 0, 0, 0, fixturesConfig, true, 0, 0, 0, false, 0, dmxController, dmxInterfacePresent)
				common.LabelButton(x, y, "", guiButtons)
			}
		}
	}
}

// GetFixtureDetailsById - find a fixture in the fixtures config.
// Returns details of the fixture.
// Returns an error.
func GetFixtureDetailsById(id int, fixtures *Fixtures) (Fixture, error) {

	if debug {
		fmt.Printf("GetFixtureDetailsById\n")
	}

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

// GetFixtureDetailsByLabel - find a fixture in the fixtures config.
// Returns details of the fixture.
// Returns an error.
func GetFixtureDetailsByLabel(label string, fixtures *Fixtures) (Fixture, error) {
	// scan the fixtures structure for the selected fixture.
	if debug {
		fmt.Printf("GetFixtureDetailsByLabel: Looking for Fixture by Label %s\n", label)
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
	eventsForLaunchpad chan common.ALight,
	guiButtons chan common.ALight,
	switchChannels []common.SwitchChannel,
	soundTriggers []*common.Trigger,
	soundConfig *sound.SoundConfig,
	dmxController *ft232.DMXController,
	fixturesConfig *Fixtures,
	dmxInterfacePresent bool) {

	if debug {
		fmt.Printf("FixtureReceiver Started %d\n", myFixtureNumber)
	}

	// Used for static fades, remember the last color.
	var lastColor common.LastColor

	stopFadeUp := make(chan bool)
	stopFadeDown := make(chan bool)

	// Loop waiting for configuration.
	for {

		// Wait for first step
		cmd := <-fixtureStepChannel

		if cmd.Blackout {
			// Soft fade downs should be disabled for blackout.
			lastColor.RGBColor = cmd.LastColor
			lastColor.ScannerColor = 0
			// Make sure we are blacked out.
			cmd.Master = 0
		}

		switch {

		case cmd.Type == "override":
			if debug {
				fmt.Printf("FixtureReceiver: override %+v\n", cmd.Override)
			}
			// Talk to the mini sequencer.
			overrideMiniSequencer(cmd, switchChannels)
			// Call the mini setter. Step through all the settings.
			overrideMiniSetter(cmd, fixturesConfig, dmxController, dmxInterfacePresent)
			continue

		case cmd.Type == "lastColor":
			if debug {
				fmt.Printf("%d:%d LastColor set to %s\n", cmd.SequenceNumber, myFixtureNumber, common.GetColorNameByRGB(cmd.LastColor))
			}
			lastColor.RGBColor = cmd.LastColor
			lastColor.ScannerColor = 0
			continue

		case cmd.Type == "switch":
			if debug {
				fmt.Printf("%d:%d Activate switch number %d name %s Postition %d Speed %d Shift %d\n", cmd.SequenceNumber, myFixtureNumber, cmd.SwiTch.Number, cmd.SwiTch.Name, cmd.SwiTch.CurrentPosition, cmd.Override.Speed, cmd.Override.Shift)
			}
			lastColor = MapSwitchFixture(cmd.SwiTch, cmd.State, cmd.Override, cmd.RGBFade, dmxController, fixturesConfig, cmd.Blackout, cmd.Master, cmd.Master, cmd.MasterChanging, lastColor, switchChannels, soundTriggers, soundConfig, dmxInterfacePresent, eventsForLaunchpad, guiButtons, fixtureStepChannel)
			continue

		case cmd.Clear:
			if debug {
				fmt.Printf("%d:%d Clear %t Blackout %t\n", cmd.SequenceNumber, myFixtureNumber, cmd.Clear, cmd.Blackout)
			}
			lastColor = clear(myFixtureNumber, cmd, stopFadeDown, stopFadeUp, fixturesConfig, dmxController, dmxInterfacePresent)
			continue

		case cmd.StartFlood:
			if debug {
				fmt.Printf("%d:%d StartFlood\n", cmd.SequenceNumber, myFixtureNumber)
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
			lastColor = startFlood(myFixtureNumber, cmd, fixturesConfig, eventsForLaunchpad, guiButtons, dmxController, dmxInterfacePresent)
			continue

		case cmd.StopFlood:
			if debug {
				fmt.Printf("%d:%d StopFlood\n", cmd.SequenceNumber, myFixtureNumber)
			}
			lastColor = stopFlood(myFixtureNumber, cmd, fixturesConfig, eventsForLaunchpad, guiButtons, dmxController, dmxInterfacePresent)
			continue

		case cmd.RGBStaticOn:
			if debug {
				fmt.Printf("%d:%d Static On Master=%d Hidden=%t\n", cmd.SequenceNumber, myFixtureNumber, cmd.Master, cmd.Hidden)
			}
			lastColor = setStaticOn(myFixtureNumber, cmd, fixturesConfig, eventsForLaunchpad, guiButtons, dmxController, dmxInterfacePresent)
			continue

		case cmd.RGBStaticFadeUp:
			if debug {
				fmt.Printf("%d:%d Static Fade Up Hidden=%t\n", cmd.SequenceNumber, myFixtureNumber, cmd.Hidden)
			}
			// FadeUpStatic doesn't return a lastColor, instead it sends a message directly to the fixture to set lastColor once it's finished fading up.
			fadeUpStatic(myFixtureNumber, cmd, lastColor, stopFadeDown, stopFadeUp, fixturesConfig, fixtureStepChannel, eventsForLaunchpad, guiButtons, dmxController, dmxInterfacePresent)
			continue

		case cmd.RGBStaticOff:
			if debug {
				fmt.Printf("%d:%d Static Off Hidden=%t\n", cmd.SequenceNumber, myFixtureNumber, cmd.Hidden)
			}
			staticOff(myFixtureNumber, cmd, lastColor, stopFadeDown, stopFadeUp, fixturesConfig, fixtureStepChannel, eventsForLaunchpad, guiButtons, dmxController, dmxInterfacePresent)
			continue

		case cmd.Type == "scanner":
			if debug {
				fmt.Printf("%d:%d Play Scanner Hidden=%t\n", cmd.SequenceNumber, myFixtureNumber, cmd.Hidden)
			}
			lastColor = playScanner(myFixtureNumber, cmd, fixturesConfig, eventsForLaunchpad, guiButtons, dmxController, dmxInterfacePresent)
			continue

		case cmd.Type == "rgb":
			if debug {
				fmt.Printf("%d:%d Play RGB Hidden=%t\n", cmd.SequenceNumber, myFixtureNumber, cmd.Hidden)
			}
			lastColor = playRGB(myFixtureNumber, cmd, fixturesConfig, eventsForLaunchpad, guiButtons, dmxController, dmxInterfacePresent)
			continue
		}
	}
}

// Clear fixture.
func clear(fixtureNumber int, cmd common.FixtureCommand, stopFadeDown chan bool, stopFadeUp chan bool, fixtures *Fixtures, dmxController *ft232.DMXController, dmxInterfacePresent bool) common.LastColor {

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

	return MapFixtures(false, false, cmd.SequenceNumber, fixtureNumber, colors.Black, 0, 0, 0, 0, 0, 0, cmd.ScannerColor, fixtures, cmd.Blackout, cmd.Master, cmd.Master, cmd.Music, cmd.Strobe, cmd.StrobeSpeed, dmxController, dmxInterfacePresent)
}

// Start Flood.
func startFlood(fixtureNumber int, cmd common.FixtureCommand, fixtures *Fixtures, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, dmxController *ft232.DMXController, dmxInterfacePresent bool) common.LastColor {
	if debug {
		fmt.Printf("Fixture:%d Set RGB Flood\n", fixtureNumber)
	}

	// TODO find sequence numbers from config.
	if cmd.SequenceNumber == 4 {
		cmd.SequenceNumber = 2
	}

	pan := 128
	tilt := 128
	shutter := FindShutter(fixtureNumber, cmd.SequenceNumber, "Open", fixtures)
	gobo := FindGoboByName(fixtureNumber, cmd.SequenceNumber, "White", fixtures)
	scannerColor := FindColor(fixtureNumber, cmd.SequenceNumber, "White", fixtures)
	rotate := 0
	program := 0

	if !cmd.Hidden {
		common.LightLamp(common.Button{X: fixtureNumber, Y: cmd.SequenceNumber}, colors.White, cmd.Master, eventsForLaunchpad, guiButtons)
		common.LabelButton(fixtureNumber, cmd.SequenceNumber, "", guiButtons)
	}

	return MapFixtures(false, false, cmd.SequenceNumber, fixtureNumber, colors.White, pan, tilt, shutter, rotate, program, gobo, scannerColor, fixtures, false, cmd.Master, cmd.Master, 0, cmd.Strobe, cmd.StrobeSpeed, dmxController, dmxInterfacePresent)

}

// Stop Flood.
func stopFlood(fixtureNumber int, cmd common.FixtureCommand, fixtures *Fixtures, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, dmxController *ft232.DMXController, dmxInterfacePresent bool) common.LastColor {

	if debug {
		fmt.Printf("Fixture:%d Set Stop RGB Flood\n", fixtureNumber)
	}

	// TODO find sequence numbers from config.
	if cmd.SequenceNumber == 4 {
		cmd.SequenceNumber = 2
	}

	if !cmd.Hidden {
		common.LightLamp(common.Button{X: fixtureNumber, Y: cmd.SequenceNumber}, colors.Black, 0, eventsForLaunchpad, guiButtons)
		common.LabelButton(fixtureNumber, cmd.SequenceNumber, "", guiButtons)
	}
	return MapFixtures(false, false, cmd.SequenceNumber, fixtureNumber, colors.Black, 0, 0, 0, 0, 0, 0, 0, fixtures, cmd.Blackout, 0, 0, 0, cmd.Strobe, cmd.StrobeSpeed, dmxController, dmxInterfacePresent)
}

// Switch On Static Scene.
func setStaticOn(fixtureNumber int, cmd common.FixtureCommand, fixtures *Fixtures, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, dmxController *ft232.DMXController, dmxInterfacePresent bool) common.LastColor {

	if debug {
		fmt.Printf("Fixture:%d setStaticOn\n", fixtureNumber)
	}

	if cmd.RGBStaticColors[fixtureNumber].Enabled {
		if debug {
			fmt.Printf("%d: Fixture:%d RGB Switch Static On - Trying to Set RGB Static Master=%d\n", cmd.SequenceNumber, fixtureNumber, cmd.Master)
		}

		// TODO find sequence numbers from config.
		if cmd.SequenceNumber == 4 {
			cmd.SequenceNumber = 2
		}

		lamp := cmd.RGBStaticColors[fixtureNumber]

		// If we're not hiding the sequence on the launchpad, show the static colors on the buttons.
		if !cmd.Hidden {
			if lamp.Flash {
				onColor := color.RGBA{R: lamp.Color.R, G: lamp.Color.G, B: lamp.Color.B}
				common.FlashLight(common.Button{X: fixtureNumber, Y: cmd.SequenceNumber}, onColor, colors.Black, eventsForLaunchpad, guiButtons)
			} else {
				common.LightLamp(common.Button{X: fixtureNumber, Y: cmd.SequenceNumber}, lamp.Color, cmd.Master, eventsForLaunchpad, guiButtons)
			}
		}

		// Look for a matching color
		color := common.GetColorNameByRGB(lamp.Color)
		// Find a suitable gobo based on the requested static lamp color.
		scannerGobo := FindGoboByName(fixtureNumber, cmd.SequenceNumber, color, fixtures)
		// Find a suitable color wheel settin based on the requested static lamp color.
		scannerColor := FindColor(fixtureNumber, cmd.SequenceNumber, color, fixtures)

		return MapFixtures(false, false, cmd.SequenceNumber, fixtureNumber, lamp.Color, common.SCANNER_MID_POINT, common.SCANNER_MID_POINT, 0, 0, 0, scannerGobo, scannerColor, fixtures, cmd.Blackout, cmd.Master, cmd.Master, 0, cmd.Strobe, cmd.StrobeSpeed, dmxController, dmxInterfacePresent)
	}

	return common.LastColor{}
}

// Fade Up RGB Static Scene
func fadeUpStatic(fixtureNumber int, cmd common.FixtureCommand, lastColor common.LastColor, stopFadeDown chan bool, stopFadeUp chan bool, fixtures *Fixtures, fixtureStepChannel chan common.FixtureCommand, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, dmxController *ft232.DMXController, dmxInterfacePresent bool) {

	if debug {
		fmt.Printf("%d: fadeUpStaticFixture: Fixture No %d LastColor %+v\n", cmd.SequenceNumber, fixtureNumber, lastColor)
	}

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

	lamp := cmd.RGBStaticColors[fixtureNumber]

	if lastColor.RGBColor != lamp.Color {
		// Now Fade Down
		go func() {
			// Soft start
			// Calulate the steps
			fadeUpValues := common.GetFadeValues(64, float64(common.MAX_DMX_BRIGHTNESS), 1, false)
			fadeDownValues := common.GetFadeValues(64, float64(common.MAX_DMX_BRIGHTNESS), 1, true)

			master := cmd.Master

			if lastColor.RGBColor != colors.EmptyColor {
				for _, fade := range fadeDownValues {

					// Look for a matching color
					color := common.GetColorNameByRGB(lastColor.RGBColor)
					// Find a suitable gobo based on the requested static lamp color.
					scannerGobo := FindGoboByName(fixtureNumber, cmd.SequenceNumber, color, fixtures)
					// Find a suitable color wheel settin based on the requested static lamp color.
					scannerColor := FindColor(fixtureNumber, cmd.SequenceNumber, color, fixtures)

					// Listen for stop command.
					select {
					case <-stopFadeDown:
						return
					case <-time.After(10 * time.Millisecond):
					}
					if !cmd.Hidden {
						common.LightLamp(common.Button{X: fixtureNumber, Y: cmd.SequenceNumber}, lastColor.RGBColor, fade, eventsForLaunchpad, guiButtons)
					}
					if cmd.Label == "chaser" {
						// If we are a RGB chaser used as a shutter chasser apply fade values to the scanner's master dimmer channel because
						// scanners doesn't have a rgb color mixing capability so the wheel has to be faded using the master.
						master = int(float64(cmd.Master) / 100 * (float64(fade) / 2.55))
					}
					MapFixtures(false, false, cmd.SequenceNumber, fixtureNumber, lastColor.RGBColor, common.SCANNER_MID_POINT, common.SCANNER_MID_POINT, 0, 0, 0, scannerGobo, scannerColor, fixtures, cmd.Blackout, fade, master, 0, cmd.Strobe, cmd.StrobeSpeed, dmxController, dmxInterfacePresent)

					// Control how long the fade take with the speed control.
					time.Sleep((5 * time.Millisecond) * (time.Duration(cmd.RGBFade)))
				}
				// Fade down complete, set lastColor to empty in the fixture.
				command := common.FixtureCommand{
					Type:      "lastColor",
					LastColor: colors.EmptyColor,
				}
				select {
				case fixtureStepChannel <- command:
				case <-time.After(100 * time.Millisecond):
				}
			}

			// If enabled fade up.
			if cmd.RGBStaticColors[fixtureNumber].Enabled {
				// Look for a matching color
				color := common.GetColorNameByRGB(lamp.Color)
				// Find a suitable gobo based on the requested static lamp color.
				scannerGobo := FindGoboByName(fixtureNumber, cmd.SequenceNumber, color, fixtures)
				// Find a suitable color wheel settin based on the requested static lamp color.
				scannerColor := FindColor(fixtureNumber, cmd.SequenceNumber, color, fixtures)

				// Fade up fixture.
				for _, fade := range fadeUpValues {
					// Listen for stop command.
					select {
					case <-stopFadeUp:
						lastColor = MapFixtures(false, false, cmd.SequenceNumber, fixtureNumber, colors.Black, 0, 0, 0, 0, 0, 0, 0, fixtures, false, 0, 0, 0, false, 0, dmxController, dmxInterfacePresent)
						return
					case <-time.After(10 * time.Millisecond):
					}
					if cmd.Label == "chaser" {
						// If we are a RGB chaser used as a shutter chasser apply fade values to the scanner's master dimmer channel because
						// scanners doesn't have a rgb color mixing capability so the wheel has to be faded using the master.
						master = int(float64(cmd.Master) / 100 * (float64(fade) / 2.55))
					}
					if !cmd.Hidden {
						common.LightLamp(common.Button{X: fixtureNumber, Y: cmd.SequenceNumber}, lamp.Color, fade, eventsForLaunchpad, guiButtons)
					}
					lastColor = MapFixtures(false, false, cmd.SequenceNumber, fixtureNumber, lamp.Color, common.SCANNER_MID_POINT, common.SCANNER_MID_POINT, 0, 0, 0, scannerGobo, scannerColor, fixtures, cmd.Blackout, fade, master, 0, cmd.Strobe, cmd.StrobeSpeed, dmxController, dmxInterfacePresent)

					// Control how long the fade take with the speed control.
					time.Sleep((5 * time.Millisecond) * (time.Duration(cmd.RGBFade)))
				}
				// Fade up complete, set lastColor up in the fixture.
				command := common.FixtureCommand{
					Type:         "lastColor",
					LastColor:    lastColor.RGBColor,
					ScannerColor: lastColor.ScannerColor,
				}
				select {
				case fixtureStepChannel <- command:
				case <-time.After(100 * time.Millisecond):
				}

			}
		}()
	}
}

func staticOff(fixtureNumber int, cmd common.FixtureCommand, lastColor common.LastColor, stopFadeDown chan bool, stopFadeUp chan bool, fixtures *Fixtures, fixtureStepChannel chan common.FixtureCommand, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, dmxController *ft232.DMXController, dmxInterfacePresent bool) {

	if debug {
		fmt.Printf("staticOff Fixture No %d", fixtureNumber)
	}

	go func() {
		var master int
		fadeDownValues := common.GetFadeValues(64, float64(common.MAX_DMX_BRIGHTNESS), 1, true)
		if lastColor.RGBColor != colors.Black {

			if debug {
				fmt.Printf("Fixture:%d =====>   RGB Static OFF -> Fade Down from LastColor %+v\n", fixtureNumber, lastColor)
			}

			var sequenceNumber int
			if cmd.Label == "chaser" {
				sequenceNumber = common.GlobalScannerSequenceNumber // Scanner sequence number from config.
			} else {
				sequenceNumber = cmd.SequenceNumber
			}

			for _, fade := range fadeDownValues {

				// Look for a matching color
				color := common.GetColorNameByRGB(lastColor.RGBColor)
				// Find a suitable gobo based on the requested static lamp color.
				scannerGobo := FindGoboByName(fixtureNumber, cmd.SequenceNumber, color, fixtures)
				// Find a suitable color wheel settin based on the requested static lamp color.
				scannerColor := FindColor(fixtureNumber, cmd.SequenceNumber, color, fixtures)

				// Listen for stop commands.
				select {
				case <-stopFadeDown:
					return
				case <-stopFadeUp:
					return
				case <-time.After(10 * time.Millisecond):
				}
				common.LightLamp(common.Button{X: fixtureNumber, Y: sequenceNumber}, lastColor.RGBColor, fade, eventsForLaunchpad, guiButtons)
				if cmd.Label == "chaser" {
					// If we are a RGB chaser used as a shutter chasser apply fade values to the scanner's master dimmer channel because
					// scanners doesn't have a rgb color mixing capability so the wheel has to be faded using the master.
					master = int(float64(cmd.Master) / 100 * (float64(fade) / 2.55))
				}
				MapFixtures(false, false, cmd.SequenceNumber, fixtureNumber, lastColor.RGBColor, common.SCANNER_MID_POINT, common.SCANNER_MID_POINT, 0, 0, 0, scannerGobo, scannerColor, fixtures, cmd.Blackout, fade, master, 0, cmd.Strobe, cmd.StrobeSpeed, dmxController, dmxInterfacePresent)

				// Control how long the fade take with the speed control.
				time.Sleep((5 * time.Millisecond) * (time.Duration(cmd.RGBFade)))
			}
			// Fade down complete, set lastColor to empty in the fixture.
			command := common.FixtureCommand{
				Type:      "lastColor",
				LastColor: colors.EmptyColor,
			}
			select {
			case fixtureStepChannel <- command:
			case <-time.After(100 * time.Millisecond):
			}
		}
	}()

}

func playRGB(fixtureNumber int, cmd common.FixtureCommand, fixtures *Fixtures, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, dmxController *ft232.DMXController, dmxInterfacePresent bool) (lastColor common.LastColor) {

	if debug {
		fmt.Printf("playRGB: fixtureNumber %d", fixtureNumber)
	}
	// Now play all the values for this state.

	// Play out fixture to DMX channels.
	fixture := cmd.RGBPosition.Fixtures[fixtureNumber]

	if cmd.Type == "rgb" && fixture.Enabled {

		if debug {
			fmt.Printf("%d: Fixture:%d RGB Mode Strobe %t\n", cmd.SequenceNumber, fixtureNumber, cmd.Strobe)
		}

		// Integrate cmd.master with fixture.Brightness.
		fixture.Brightness = int((float64(fixture.Brightness) / 100) * (float64(cmd.Master) / 2.55))

		// If we're a shutter chaser flavoured RGB sequence, then disable everything except the brightness.
		if cmd.Label == "chaser" {
			scannerFixturesSequenceNumber := common.GlobalScannerSequenceNumber // Scanner sequence number from config.
			if !cmd.Hidden {
				common.LightLamp(common.Button{X: fixtureNumber, Y: scannerFixturesSequenceNumber}, fixture.Color, fixture.Brightness, eventsForLaunchpad, guiButtons)
			}

			// Fixture brightness is sent as master in this case because a shutter chaser is controlling a scanner lamp.
			// and these generally don't have any RGB color channels that can be controlled with brightness.
			// So the only way to make the lamp in the scanner change intensity is to vary the master brightness channel.

			// Lookup chaser lamp color based on the requested fixture base color.
			// We can't use the faded color as its impossibe to lookup the base color from a faded color.
			// GetColorNameByRGB will return white if the color is not found.
			color := common.GetColorNameByRGB(fixture.BaseColor)

			// Find a suitable gobo based on the requested chaser lamp color.
			scannerGobo := FindGoboByName(fixtureNumber, scannerFixturesSequenceNumber, color, fixtures)
			// Find a suitable color wheel setting based on the requested static lamp color.
			scannerColor := FindColor(fixtureNumber, scannerFixturesSequenceNumber, color, fixtures)

			lastColor = MapFixtures(true, cmd.ScannerChaser, scannerFixturesSequenceNumber, fixtureNumber, fixture.Color, 0, 0, 0, 0, 0, scannerGobo, scannerColor, fixtures, cmd.Blackout, cmd.Master, fixture.Brightness, cmd.Music, cmd.Strobe, cmd.StrobeSpeed, dmxController, dmxInterfacePresent)
		} else {
			if !cmd.Hidden {
				common.LightLamp(common.Button{X: fixtureNumber, Y: cmd.SequenceNumber}, fixture.Color, cmd.Master, eventsForLaunchpad, guiButtons)
			}
			lastColor = MapFixtures(false, cmd.ScannerChaser, cmd.SequenceNumber, fixtureNumber, fixture.Color, 0, 0, 0, 0, 0, cmd.ScannerGobo, cmd.ScannerColor, fixtures, cmd.Blackout, cmd.Master, cmd.Master, cmd.Music, cmd.Strobe, cmd.StrobeSpeed, dmxController, dmxInterfacePresent)
		}
	}

	return lastColor
}

func playScanner(fixtureNumber int, cmd common.FixtureCommand, fixtures *Fixtures, eventsForLaunchpad chan common.ALight, guiButtons chan common.ALight, dmxController *ft232.DMXController, dmxInterfacePresent bool) (lastColor common.LastColor) {

	if debug {
		fmt.Printf("Fixture:%d playScanner\n", fixtureNumber)
	}

	// Find the fixture
	fixture := cmd.ScannerPosition.Fixtures[fixtureNumber]

	if fixture.Enabled {

		if debug {
			fmt.Printf("Fixture:%d Play Scanner \n", fixtureNumber)
		}

		// In the case of a scanner, they usually have a shutter and a master dimmer to control the brightness
		// of the lamp. Problem is we can't use the shutter for the control of the overall brightness and the
		// master for the master dimmmer like we do with RGB fixture. The shutter noramlly is more of a switch
		// eg. Open , Closed , Strobe etc. If I want to slow fade through a set of scanners I need to use the
		// brightness for control. Which means I need to combine the master and the control brightness
		// at this stage.
		scannerBrightness := int(math.Round((float64(fixture.Brightness) / 100) * (float64(cmd.Master) / 2.55)))
		// Tell the scanner what to do.
		lastColor = MapFixtures(false, cmd.ScannerChaser, cmd.SequenceNumber, fixtureNumber, fixture.ScannerColor, fixture.Pan, fixture.Tilt,
			fixture.Shutter, cmd.Rotate, cmd.Program, cmd.ScannerGobo, cmd.ScannerColor, fixtures, cmd.Blackout, cmd.Master, scannerBrightness, cmd.Music, cmd.Strobe, cmd.StrobeSpeed, dmxController, dmxInterfacePresent)

		// Scannner is rotating, work out what to do with the launchpad lamps.
		if !cmd.Hidden {
			// Every quater turn, display a color to represent a position in the rotation.
			howOftern := cmd.NumberSteps / 4
			if howOftern != 0 {
				if cmd.Step%howOftern == 0 {
					// We're not in chase mode so use the color generated in the pattern generator.common.
					common.LightLamp(common.Button{X: fixtureNumber, Y: cmd.SequenceNumber}, fixture.Color, cmd.Master, eventsForLaunchpad, guiButtons)
					common.LabelButton(fixtureNumber, cmd.SequenceNumber, "", guiButtons)
				}
			}
		}
	} else {
		// This scanner is disabled, shut it off.
		lastColor = MapFixtures(false, false, cmd.SequenceNumber, fixtureNumber, colors.Black, 0, 0, 0, 0, 0, 0, 0, fixtures, false, 0, 0, 0, false, 0, dmxController, dmxInterfacePresent)
	}

	return lastColor
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

func findChannelSettingByChannelNameAndSettingName(fixture *Fixture, channelName string, settingName string) (int, error) {

	if debug {
		fmt.Printf("findChannelSettingByChannelNameAndSettingName for fixture %s on channel %s setting %s\n", fixture.Name, channelName, settingName)
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

					var v int
					var err error
					// If the setting value contains a "-" remove it and then take the first valuel.
					if strings.Contains(setting.Value, "-") {
						// We've found a range of values.
						// Find the start value
						numbers := strings.Split(setting.Value, "-")
						v, err = strconv.Atoi(numbers[0])
						if err != nil {
							return 0, err
						}
					} else {
						v, err = strconv.Atoi(setting.Value)
						if err != nil {
							return 0, err
						}
					}
					if debug {
						fmt.Printf("Value Returned %d\n", v)
					}

					return v, nil
				}
			}
		}
	}

	return 0, fmt.Errorf("setting %s not found in channel %s for fixture %s", settingName, channelName, fixture.Name)
}

func findChannelSettingByNameAndSpeed(fixtureName string, channelName string, settingName string, settingSpeed string, fixtures *Fixtures) (int, error) {

	if debug {
		fmt.Printf("findChannelSettingByNameAndSpeed for fixture %s channel name %s setting name %s and setting speed %s\n", fixtureName, channelName, settingName, settingSpeed)
	}

	if channelName == "" {
		return 0, fmt.Errorf("findChannelSettingByNameAndSpeed: error channel name is empty")
	}
	if settingName == "" {
		return 0, fmt.Errorf("findChannelSettingByNameAndSpeed: error setting name is empty")
	}
	if settingSpeed == "" {
		return 0, fmt.Errorf("findChannelSettingByNameAndSpeed: error setting speed is empty")
	}

	for _, fixture := range fixtures.Fixtures {

		if fixtureName == fixture.Name {
			for _, channel := range fixture.Channels {

				if channel.Name == channelName {
					if debug {
						fmt.Printf("fixture %s: inspect channel %s for %s\n", fixture.Name, channel.Name, settingName)
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

							// If the setting value contains a "-" remove it and then take the first valuel.
							var err error
							var v int
							if strings.Contains(setting.Value, "-") {
								// We've found a range of values.
								// Find the start value
								numbers := strings.Split(setting.Value, "-")
								v, err = strconv.Atoi(numbers[0])
								if err != nil {
									return 0, err
								}
							} else {
								v, err = strconv.Atoi(setting.Value)
								if err != nil {
									return 0, err
								}
							}

							if debug {
								fmt.Printf("speed found is %d\n", v)
							}
							return v, nil
						}
					}
				}
			}
		}
	}

	return 0, fmt.Errorf("warning: channel %s setting %s not found in fixture :%s", channelName, settingSpeed, fixtureName)
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

// When want to light a DMX fixture we need for find it in our fuxture.yaml configuration file.
// This function maps the requested fixture into a DMX address.
func MapFixtures(chaser bool, hadShutterChase bool,
	mySequenceNumber int,
	displayFixture int,
	color color.RGBA,
	pan int, tilt int, shutter int, rotate int, program int, selectedGobo int, scannerColor int,
	fixtures *Fixtures, blackout bool, brightness int, master int, music int, strobe bool, strobeSpeed int,
	dmxController *ft232.DMXController, dmxInterfacePresent bool) (lastColor common.LastColor) {

	if debug {
		fmt.Printf("MapFixtures Fixture No %d Sequence No %d Color %+v\n", displayFixture, mySequenceNumber, color)
	}

	// We control the brightness of each color with the brightness value.
	// The overall fixture brightness is set from the master value.
	var Red float64
	var Green float64
	var Blue float64
	var Master int
	var Brightness int

	// The best blackout solution is to manually pull all the colors down.
	if blackout {
		Master = 0
		Brightness = 0
		Red = 0
		Green = 0
		Blue = 0
	} else {
		Master = master
		Brightness = brightness
		Red = (float64(color.R) / 100) * (float64(Brightness) / 2.55)
		Green = (float64(color.G) / 100) * (float64(Brightness) / 2.55)
		Blue = (float64(color.B) / 100) * (float64(Brightness) / 2.55)
	}
	if debug {
		fmt.Printf("MapFixtures Fixture No %d Sequence No %d Red %f Green %f Blue %f Brightness %d Master %d Blackout %t\n", displayFixture, mySequenceNumber, Red, Green, Blue, Brightness, Master, blackout)
	}

	for _, fixture := range fixtures.Fixtures {
		if fixture.Group == mySequenceNumber+1 {
			for channelNumber, channel := range fixture.Channels {

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
									if debug {
										fmt.Printf("MapFixtures: fixture %s: send ChannelName %s Address %d Value %d \n", fixture.Name, channel.Name, fixture.Address+int16(channelNumber), int(reverse_dmx(Master)))
									}
									SetChannel(fixture.Address+int16(channelNumber), byte(reverse_dmx(Master)), dmxController, dmxInterfacePresent)
								} else {
									if debug {
										fmt.Printf("MapFixtures: fixture %s: send ChannelName %s Address %d Value %d \n", fixture.Name, channel.Name, fixture.Address+int16(channelNumber), Master)
									}
									SetChannel(fixture.Address+int16(channelNumber), byte(Master), dmxController, dmxInterfacePresent)
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
								SetChannel(fixture.Address+int16(channelNumber), byte(reverse_dmx(Master)), dmxController, dmxInterfacePresent)
							} else {
								SetChannel(fixture.Address+int16(channelNumber), byte(Master), dmxController, dmxInterfacePresent)
							}
						}
						// Shutter
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
				}
			}
		}
	}

	lastColor.RGBColor = color
	lastColor.ScannerColor = scannerColor

	return lastColor
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

// MapSwitchFixture is repsonsible for playing out the state of a swicth.
// The switch is idendifed by the sequence and switch number.
func MapSwitchFixture(swiTch common.Switch,
	state common.State,
	override common.Override,
	RGBFade int,
	dmxController *ft232.DMXController,
	fixturesConfig *Fixtures, blackout bool,
	brightness int, master int, masterChanging bool, lastColor common.LastColor,
	switchChannels []common.SwitchChannel,
	SoundTriggers []*common.Trigger,
	soundConfig *sound.SoundConfig,
	dmxInterfacePresent bool,
	eventsForLaunchpad chan common.ALight,
	guiButtons chan common.ALight,
	fixtureStepChannel chan common.FixtureCommand) common.LastColor {

	var useFixtureLabel string

	if debug {
		fmt.Printf("MapSwitchFixture switchNumber %d, current position %d fade speed %d\n", swiTch.Number, swiTch.CurrentPosition, RGBFade)
	}

	// Master dinmmer is the best way to blackout a switch.
	if blackout {
		master = 0
	}

	// We start by having the switch and its current state passed in.

	// Now we find the fixture used by the switch
	if swiTch.UseFixture != "" {
		// use this fixture for the sequencer actions
		// BTW UseFixtureLabel is the label for the fixture NOT the name.
		useFixtureLabel = swiTch.UseFixture

		if debug {
			fmt.Printf("useFixtureLabel %s  blackout is %t\n", useFixtureLabel, blackout)
		}

		// Find the details of the fixture for this switch.
		thisFixture, err := FindFixtureByLabel(useFixtureLabel, fixturesConfig)
		if err != nil {
			fmt.Printf("error %s\n", err.Error())
			return lastColor
		}

		if debug {
			fmt.Printf("Found fixture Name %s \n", thisFixture.Name)
		}

		// Look for Master channel in this fixture identified by ID.
		masterChannel, err := FindChannelNumberByName(thisFixture, "Master")
		if err != nil && debug {
			fmt.Printf("warning! fixture %s: %s\n", thisFixture.Name, err)
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
			newAction.ProgramSpeed = action.ProgramSpeed
			newAction.Strobe = action.Strobe
			newAction.Map = action.Map
			newAction.Gobo = action.Gobo
			newAction.GoboSpeed = action.GoboSpeed
			newMiniSequencer(thisFixture, swiTch, override, newAction, dmxController, fixturesConfig, switchChannels, soundConfig, blackout, brightness, master, masterChanging, lastColor, dmxInterfacePresent, eventsForLaunchpad, guiButtons, fixtureStepChannel)
			if action.Mode != "Static" {
				lastColor.RGBColor = colors.EmptyColor
			}
		}

		// If there are no actions, turn off any previos mini sequencers for this switch.
		if len(state.Actions) == 0 {
			newAction := Action{}
			newAction.Name = "Off"
			newAction.Number = 1
			newAction.Mode = "Off"
			lastColor := common.LastColor{}
			newMiniSequencer(thisFixture, swiTch, override, newAction, dmxController, fixturesConfig, switchChannels, soundConfig, blackout, brightness, master, masterChanging, lastColor, dmxInterfacePresent, eventsForLaunchpad, guiButtons, fixtureStepChannel)
		}

		// Now play any preset DMX values directly to the universe.
		// Step through all the settings.
		for _, newSetting := range state.Settings {
			newMiniSetter(thisFixture, override, newSetting, masterChannel, dmxController, master, dmxInterfacePresent)
		}
	}
	return lastColor
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

func FindFixtureByGroupAndNumber(sequenceNumber int, fixtureNumber int, fixtures *Fixtures) (*Fixture, error) {

	if debug {
		fmt.Printf("FindFixtureByGroupAndNumber seq %d fixture %d\n", sequenceNumber, fixtureNumber)
	}
	for _, fixture := range fixtures.Fixtures {
		if fixture.Group == sequenceNumber+1 {
			if fixture.Number == fixtureNumber+1 {
				if debug {
					fmt.Printf("Found fixture %s Group %d Number %d Address %d\n", fixture.Name, fixture.Group, fixture.Number, fixture.Address)
				}
				return &fixture, nil
			}
		}
	}
	return nil, fmt.Errorf("FindFixtureByGroupAndNumber: failed to find fixture for sequence %d fixture %d", sequenceNumber, fixtureNumber)
}

func FindFixtureByLabel(label string, fixtures *Fixtures) (*Fixture, error) {

	if debug {
		fmt.Printf("FindFixtureByLabel label is %s\n", label)
	}
	for _, fixture := range fixtures.Fixtures {
		if fixture.Label == label {
			if debug {
				fmt.Printf("Found fixture %s Group %d Number %d Address %d\n", fixture.Name, fixture.Group, fixture.Number, fixture.Address)
			}
			return &fixture, nil
		}
	}
	return nil, fmt.Errorf("FindFixtureByGroupAndNumber: failed to find fixture by label %s", label)
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

func FindFixtureAddressByName(fixtureName string, fixtures *Fixtures) string {
	if debug {
		fmt.Printf("FindFixtureAddressByName: Looking for fixture by Name %s\n", fixtureName)
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

// FindShutter takes the name of a gobo channel setting like "Open" and returns the gobo number  for this type of scanner.
func FindShutter(myFixtureNumber int, mySequenceNumber int, shutterName string, fixtures *Fixtures) int {

	if debug {
		fmt.Printf("FindShutter\n")
	}

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

// findGobo takes the gobo dmx value and returns the gobo name for this fixture.
func FindGoboByDMXValue(fixture *Fixture, dmxValue string) string {

	if debug {
		fmt.Printf("FindGoboByDMXValue\n")
	}

	for _, channel := range fixture.Channels {
		if strings.Contains(channel.Name, "Gobo") {
			for _, setting := range channel.Settings {
				if setting.Value == dmxValue {
					return setting.Name
				}
			}
		}
	}

	return "Unknown"
}

// FindGoboNameByNumber takes the gobo number and returns the gobo name for this fixture.
func FindGoboNameByNumber(fixture *Fixture, number int) string {

	if debug {
		fmt.Printf("FindGoboByNumber Looking for gobo %d in fixture %s\n", number, fixture.Name)
	}

	if number == -1 {
		if debug {
			fmt.Printf("Gobo %d Name Auto\n", number)
		}
		return "Auto"
	}

	for _, channel := range fixture.Channels {
		if strings.Contains(channel.Name, "Gobo") {
			for _, setting := range channel.Settings {
				if setting.Number == number {
					if debug {
						fmt.Printf("Gobo %d Name %s\n", setting.Number, setting.Name)
					}
					return setting.Name
				}
			}
		}
	}

	return "Unknown"
}

// FindRotateDMXValueByIndex takes the rotate setting number and returns the DMX value which selects this speed.
func FindRotateDMXValueByIndex(fixture *Fixture, index int) int {

	if debug {
		fmt.Printf("FindRotateDMXValueByIndex Looking for rotate speed index %d in fixture %s\n", index, fixture.Name)
	}

	for _, channel := range fixture.Channels {
		if strings.Contains(channel.Name, "Rotate") {
			for _, setting := range channel.Settings {
				if setting.Number == index {
					if debug {
						fmt.Printf("Rotate Speed %d Name %s\n", setting.Number, setting.Name)
					}
					dmx, _ := strconv.Atoi(setting.Value)
					return dmx
				}
			}
		}
	}

	return 0
}

// FindRotateSpeedNameByNumber takes the rotate speed number and returns the rotate speed name for this fixture.
func FindRotateSpeedNameByNumber(fixture *Fixture, number int) string {

	if debug {
		fmt.Printf("FindRotateSpeedNameByNumber Looking for rotate speed %d in fixture %s\n", number, fixture.Name)
	}

	if number == -1 {
		if debug {
			fmt.Printf("Rotate Speed %d Name Auto\n", number)
		}
		return "Auto"
	}

	for _, channel := range fixture.Channels {
		if strings.Contains(channel.Name, "Rotate") {
			for _, setting := range channel.Settings {
				if setting.Number == number {
					if debug {
						fmt.Printf("Rotate Speed %d Name %s\n", setting.Number, setting.Name)
					}
					return setting.Name
				}
			}
		}
	}

	return "Unknown"
}

// FindRotateSpeedNumberByName takes the rotate speed name and returns the rotate speed setting number for this fixture.
func FindRotateSpeedNumberByName(fixture *Fixture, rotateSettingName string) int {

	if debug {
		fmt.Printf("FindRotateSpeedNameByNumber Looking for rotate speed %s in fixture %s\n", rotateSettingName, fixture.Name)
	}

	for _, channel := range fixture.Channels {
		if strings.Contains(channel.Name, "Rotate") {
			for _, setting := range channel.Settings {
				if setting.Name == rotateSettingName {
					if debug {
						fmt.Printf("Rotate Speed %d Name %s\n", setting.Number, setting.Name)
					}
					return setting.Number
				}
			}
		}
	}

	return 0
}

// FindGoboDMXValueByNumber takes the gobo number and returns the DMX value which selects this Gobo.
func FindGoboDMXValueByNumber(fixture *Fixture, number int) int {

	if debug {
		fmt.Printf("FindGoboDMXValueByNumber Looking for gobo %d in fixture %s\n", number, fixture.Name)
	}

	for _, channel := range fixture.Channels {
		if strings.Contains(channel.Name, "Gobo") {
			for _, setting := range channel.Settings {
				if setting.Number == number {
					if debug {
						fmt.Printf("Gobo %d Name %s\n", setting.Number, setting.Name)
					}
					dmx, _ := strconv.Atoi(setting.Value)
					return dmx
				}
			}
		}
	}

	return 0
}

// FindColorNameByNumber takes the color number and returns the color name for this fixture.
func FindColorNameByNumber(fixture *Fixture, number int) string {

	if debug {
		fmt.Printf("FindColorNameByNumber looking for color number %d inside fixture %s\n", number, fixture.Name)
	}

	for _, channel := range fixture.Channels {
		if strings.Contains(channel.Name, "Color") {
			for _, setting := range channel.Settings {
				if setting.Number == number {
					if debug {
						fmt.Printf("Found name %s\n", setting.Name)
					}
					return setting.Name
				}
			}
		}
	}

	return "Unknown"
}

// findGobo takes the name of a gobo channel setting like "Open" and returns the gobo number  for this type of scanner.
func FindGoboByName(myFixtureNumber int, mySequenceNumber int, selectedGobo string, fixtures *Fixtures) int {

	if debug {
		fmt.Printf("FindGobo\n")
	}

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

// FindColor takes the name of a color channel setting like "White" and returns the color number for this type of scanner.
func FindColor(myFixtureNumber int, mySequenceNumber int, color string, fixtures *Fixtures) int {

	if debug {
		fmt.Printf("FindColor\n")
	}

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
	fixtureInfo.HasRotate = isThisAChannel(*thisFixture, "Rotate")
	fixtureInfo.HasRotateSpeed = isThisAChannel(*thisFixture, "RotateSpeed")

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

	fixtureInfo.HasColorWheel = isThisAChannel(*thisFixture, "Color")
	fixtureInfo.HasGobo = isThisAChannel(*thisFixture, "Gobo")
	fixtureInfo.HasProgram = isThisAChannel(*thisFixture, "Program")
	fixtureInfo.HasProgramSpeed = isThisAChannel(*thisFixture, "ProgramSpeed")
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

func isThisAChannel(thisFixture Fixture, channelName string) bool {

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

func DiscoverSwitchOveride(fixture *Fixture, switchNumber int, stateNumber int, fixturesConfig *Fixtures) common.Override {

	// Convert this switches action into a config we can query.
	action := GetSwitchConfig(switchNumber, int16(stateNumber), fixturesConfig)
	cfg := GetConfig(action, fixture, fixturesConfig)

	if debug {
		fmt.Printf("---> Discover Action Name %s Color %+v Cfg.Colors %+v\n", action.Name, action.Colors, cfg.Colors)
	}

	// Create a new override for this action.
	newOverride := common.Override{}

	// if action.Mode == "Off" {
	// 	newOverride.Colors = []color.RGBA{}
	// }

	if action.Mode == "Static" {
		newOverride.Speed = cfg.Speed
		newOverride.Shift = cfg.Shift
		newOverride.Size = cfg.Size
		newOverride.Fade = cfg.Fade
		newOverride.RotateSpeed = cfg.RotateSpeed
		newOverride.RotateSpeedName = FindRotateSpeedNameByNumber(fixture, cfg.RotateSpeed)

		newOverride.Color = cfg.Color
		newOverride.Colors = cfg.Colors
		newOverride.ColorName = FindColorNameByNumber(fixture, cfg.Color)
		newOverride.Gobo = cfg.Gobo
		newOverride.GoboName = FindGoboNameByNumber(fixture, cfg.Gobo)
	}

	if action.Mode == "Control" {
		newOverride.Speed = cfg.Speed
		newOverride.Shift = cfg.Shift
		newOverride.Size = cfg.Size
		newOverride.Fade = cfg.Fade
		newOverride.RotateSpeed = cfg.RotateSpeed
		newOverride.RotateSpeedName = FindRotateSpeedNameByNumber(fixture, cfg.RotateSpeed)

		newOverride.Color = cfg.Color
		newOverride.Colors = cfg.Colors
		newOverride.ColorName = FindColorNameByNumber(fixture, cfg.Color)
		newOverride.Gobo = cfg.Gobo
		newOverride.GoboName = FindGoboNameByNumber(fixture, cfg.Gobo)
	}

	if action.Mode == "Chase" {
		newOverride.Speed = cfg.Speed
		newOverride.Shift = cfg.Shift
		newOverride.Size = cfg.Size
		newOverride.Fade = cfg.Fade
		newOverride.RotateSpeed = cfg.RotateSpeed
		newOverride.RotateSpeedName = FindRotateSpeedNameByNumber(fixture, cfg.RotateSpeed)

		newOverride.Color = cfg.Color
		newOverride.Colors = cfg.Colors
		newOverride.ColorName = FindColorNameByNumber(fixture, cfg.Color)
		newOverride.Gobo = cfg.Gobo
		newOverride.GoboName = FindGoboNameByNumber(fixture, cfg.Gobo)
	}

	if action.Mode == "Setting" {
		newOverride.Speed = cfg.Speed
		newOverride.Shift = cfg.Shift
		newOverride.Size = cfg.Size
		newOverride.Fade = cfg.Fade
		newOverride.RotateSpeed = cfg.RotateSpeed
		newOverride.RotateSpeedName = FindRotateSpeedNameByNumber(fixture, cfg.RotateSpeed)

		newOverride.Color = cfg.Color
		newOverride.Colors = cfg.Colors
		newOverride.ColorName = FindColorNameByNumber(fixture, cfg.Color)
		newOverride.Gobo = cfg.Gobo
		newOverride.GoboName = FindGoboNameByNumber(fixture, cfg.Gobo)
	}

	if debug {
		fmt.Printf("Action Name %s\n", action.Name)
		fmt.Printf("\tSwitch Number %d State Number %d\n", switchNumber, stateNumber)
		fmt.Printf("\t\t Rotate Speed %d\n", newOverride.RotateSpeed)
		fmt.Printf("\t\t Colors %+v\n", newOverride.Colors)
		fmt.Printf("\t\t Color %+v\n", newOverride.Color)
		fmt.Printf("\t\t Color Name %s\n", newOverride.ColorName)
		fmt.Printf("\t\t Gobo action %s newOverride Gobo %d Gobo Name %s\n", action.Gobo, newOverride.Gobo, newOverride.GoboName)
	}
	return newOverride
}

func GetSwitchConfig(switchNumber int, switchState int16, fixturesConfig *Fixtures) Action {

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
							action := convertSettingToAction(fixture, state.Settings)
							if debug {
								fmt.Printf("Settings:- action number %d colors %+v\n", action.Number, action.Colors)
							}
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
		if debug {
			fmt.Printf("setting name %s label %s Channel %s value %s\n", setting.Name, setting.Label, setting.Channel, setting.Value)
		}

		if setting.Channel == "Color" {
			// If a setting has a channel name which is a number we lookup that color name.
			if colorNumber, err := strconv.Atoi(setting.Value); err == nil {
				// Lookup color number in list of available colors.
				colorName := FindColorNameByNumber(&fixture, colorNumber)
				newAction.Colors = []string{colorName}
			} else {
				// we use that name as the color.
				newAction.Colors = []string{setting.Value}
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
