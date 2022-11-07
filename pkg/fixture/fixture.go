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
	Label       string   `yaml:"label"`
	Values      []Value  `yaml:"values"`
	ButtonColor Color    `yaml:"buttoncolor"`
	Master      int      `yaml:"master"`
	Actions     []Action `yaml:"actions"`
	Flash       bool     `yaml:"flash"`
}

type Action struct {
	Name    string   `yaml:"name"`
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

type Switch struct {
	Name        string  `yaml:"name"`
	Label       string  `yaml:"label"`
	Number      int     `yaml:"number"`
	Description string  `yaml:"description"`
	States      []State `yaml:"states"`
	Fixture     string  `yaml:"fixture"`
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
	NumberChannels int       `yaml:"use_channels"`
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
func LoadFixtures(filename string) (fixtures *Fixtures, err error) {

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
	fixtureStepChannel chan common.FixtureCommand,
	eventsForLauchpad chan common.ALight,
	guiButtons chan common.ALight,
	dmxController *ft232.DMXController,
	fixtures *Fixtures) {

	// Outer loop wait for configuration.
	for {

		// Wait for first step
		cmd := <-fixtureStepChannel

		// If we're a RGB fixture implement the flood and static features.
		if cmd.Type == "rgb" {
			if cmd.Clear {
				turnOffFixtures(cmd, myFixtureNumber, mySequenceNumber, fixtures, dmxController, eventsForLauchpad, guiButtons)
				continue
			}
			if cmd.StartFlood {
				MapFixtures(cmd.SequenceNumber, dmxController, myFixtureNumber, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, nil, fixtures, sequence.Blackout, sequence.Master, sequence.Master, sequence.StrobeSpeed)
				common.LightLamp(common.ALight{X: myFixtureNumber, Y: sequence.Number, Red: 255, Green: 255, Blue: 255, Brightness: 255}, eventsForLauchpad, guiButtons)
				continue
			}
			if cmd.StopFlood {
				MapFixtures(cmd.SequenceNumber, dmxController, myFixtureNumber, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, nil, fixtures, sequence.Blackout, sequence.Master, sequence.Master, sequence.StrobeSpeed)
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
				lightStaticFixture(sequence, myFixtureNumber, dmxController, eventsForLauchpad, guiButtons, fixtures, true)
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
				MapFixtures(mySequenceNumber, dmxController, myFixtureNumber, red, green, blue, white, 0, 0, 0, 0, 0, 0, 0, 0, 0, cmd.ScannerColor, fixtures, cmd.Blackout, cmd.Master, cmd.Master, cmd.StrobeSpeed)
			}
		}

		if cmd.Type == "scanner" {

			// Turn off the scanners in flood mode.
			if cmd.StartFlood {
				turnOnFixtures(cmd, myFixtureNumber, mySequenceNumber, fixtures, dmxController, eventsForLauchpad, guiButtons)
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
				MapFixtures(mySequenceNumber, dmxController, myFixtureNumber, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, nil, fixtures, cmd.Blackout, 0, 0, 0)
				// Locking for write.
				sequence.DisableOnceMutex.Lock()
				sequence.DisableOnce[myFixtureNumber] = false
				sequence.DisableOnceMutex.Unlock()
				continue
			}

			if enabled {

				// If enables activate the physical scanner.
				MapFixtures(mySequenceNumber, dmxController, myFixtureNumber, fixture.ScannerColor.R, fixture.ScannerColor.G, fixture.ScannerColor.B, fixture.ScannerColor.W, fixture.ScannerColor.A, fixture.ScannerColor.UV, fixture.Pan, fixture.Tilt,
					fixture.Shutter, cmd.Rotate, cmd.Music, cmd.Program, cmd.ScannerSelectedGobo, cmd.ScannerColor, fixtures, cmd.Blackout, cmd.Master, cmd.Master, cmd.StrobeSpeed)

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

func lookUpSettingLabelInFixtureDefinition(group int, switchNumber int, channelName string, name string, label string, fixtures *Fixtures) (int, error) {

	if debug {
		fmt.Printf("lookUpSettingLabelInFixtureDefinition for %s\n", channelName)
	}

	var fixtureName string
	for _, fixture := range fixtures.Fixtures {
		if fixture.Group == group {
			if debug {
				fmt.Printf("---> fixture %s\n", fixture.Name)
				fmt.Printf("---> fixture.group %d group %d\n", fixture.Group, group)
				fmt.Printf("---> channels %+v\n", fixture.Channels)
			}
			fixtureName = fixture.Name
			for _, channel := range fixture.Channels {
				if debug {
					fmt.Printf("---> inspect channel %s for %s\n", channel.Name, name)
				}
				if channel.Name == name {
					if debug {
						fmt.Printf("---> channel.Settings %+v\n", channel.Settings)
					}
					for _, setting := range channel.Settings {
						if debug {
							fmt.Printf("---> inspect setting %+v \n", setting)
							fmt.Printf("---> setting.Label %s = label %s\n", setting.Label, label)
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
				fmt.Printf("---> fixture %s\n", fixture.Name)
				fmt.Printf("---> fixture.group %d group %d\n", fixture.Group, group)
				fmt.Printf("---> channels %+v\n", fixture.Channels)
			}
			//if fixture.Number == switchNumber {
			fixtureName = fixture.Name
			for channelNumber, channel := range fixture.Channels {
				if debug {
					fmt.Printf("---> inspect channel %s for %s\n", channel.Name, channelName)
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
	displayFixture int, R int, G int, B int, W int, A int, uv int,
	Pan int, Tilt int, Shutter int, Rotate int, Music int, Program int,
	selectedGobo int, scannerColor map[int]int,
	fixtures *Fixtures, blackout bool, brightness int, master int, strobe int) {

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
					if strings.Contains(channel.Name, "Rotate") {
						dmxController.SetChannel(fixture.Address+int16(channelNumber), byte(Rotate))
					}
					if strings.Contains(channel.Name, "Music") {
						dmxController.SetChannel(fixture.Address+int16(channelNumber), byte(Music))
					}
					if strings.Contains(channel.Name, "Program") {
						dmxController.SetChannel(fixture.Address+int16(channelNumber), byte(Program))
					}
					if strings.Contains(channel.Name, "ProgramSpeed") {
						dmxController.SetChannel(fixture.Address+int16(channelNumber), byte(Program))
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
				if strings.Contains(channel.Name, "White"+strconv.Itoa(displayFixture+1)) {
					dmxController.SetChannel(fixture.Address+int16(channelNumber), byte(int(White)))
				}
				if strings.Contains(channel.Name, "Amber"+strconv.Itoa(displayFixture+1)) {
					dmxController.SetChannel(fixture.Address+int16(channelNumber), byte(int(Amber)))
				}
				if strings.Contains(channel.Name, "UV"+strconv.Itoa(displayFixture+1)) {
					dmxController.SetChannel(fixture.Address+int16(channelNumber), byte(int(UV)))
				}
			}
		}
	}
}

func MapSwitchFixture(mySequenceNumber int,
	dmxController *ft232.DMXController,
	switchNumber int, currentState int,
	fixtures *Fixtures, blackout bool, brightness int, master int,
	switchChannels map[int]common.SwitchChannel,
	SoundTriggers map[int]*common.Trigger,
	soundConfig *sound.SoundConfig) {

	var fixtureName string

	if debug {
		fmt.Printf("MapSwitchFixture switchNumber %d, currentState %d\n", switchNumber, currentState)
	}

	// Step through the fixture config file looking for the group that matches mysequence number.
	for _, fixture := range fixtures.Fixtures {
		if fixture.Group-1 == mySequenceNumber {
			for _, swiTch := range fixture.Switches {

				if swiTch.Fixture != "" {
					// start a fixture sequencer
					fixtureName = swiTch.Fixture
				}
				if switchNumber+1 == swiTch.Number {
					for stateNumber, state := range swiTch.States {
						if stateNumber == currentState {

							// Play Actions which send messages to a dedicated mini sequencer.
							for _, action := range state.Actions {
								newMiniSequencer(fixtureName, swiTch.Number, action, dmxController, fixtures, switchChannels, soundConfig, blackout, master)
							}

							// Play DMX values directly to the univers.
							for _, value := range state.Values {
								if blackout {
									v, _ := strconv.ParseFloat(value.Setting, 16)
									dmxController.SetChannel(fixture.Address+int16(v), byte(0))
								} else {
									// This should be controlled by the master brightness
									if strings.Contains(value.Name, "master") || strings.Contains(value.Name, "dimmer") {
										v, _ := strconv.ParseFloat(value.Setting, 32)
										howBright := int((float64(v) / 100) * (float64(brightness) / 2.55))
										if strings.Contains(value.Name, "reverse") || strings.Contains(value.Name, "invert") {
											c, _ := strconv.ParseFloat(value.Channel, 16)
											dmxController.SetChannel(fixture.Address+int16(c), byte(reverse_dmx(howBright)))
										} else {
											c, _ := strconv.Atoi(value.Channel)
											dmxController.SetChannel(fixture.Address+int16(c), byte(howBright))
										}
									} else {

										// If the setting has is a number set it directly.
										if IsNumericOnly(value.Setting) {

											v, _ := strconv.ParseFloat(value.Setting, 32)
											if IsNumericOnly(value.Channel) {
												// If the channel has is a number set it directly.
												c, _ := strconv.ParseFloat(value.Channel, 16)
												dmxController.SetChannel(fixture.Address+int16(c), byte(v))
											} else {
												// Handle the fact that the channel may be a label as well.
												fixture := findFixtureByName(fixtureName, fixtures)
												c, _ := lookUpChannelNumberByNameInFixtureDefinition(fixture.Group, switchNumber, value.Channel, fixtures)
												dmxController.SetChannel(fixture.Address+int16(c), byte(v))
											}

										} else {
											// If the setting contains a label, look up the label in the fixture definition.
											fixture := findFixtureByName(fixtureName, fixtures)
											v, err := lookUpSettingLabelInFixtureDefinition(fixture.Group, swiTch.Number, value.Channel, value.Name, value.Setting, fixtures)
											if err != nil {
												fmt.Printf("dmxlights: error failed to find Name=%s in switch Setting=%s in fixture config for switch Number=%d: %s\n", value.Name, value.Setting, swiTch.Number, err.Error())
												fmt.Printf("fixture.Group=%d, swiTch.Number=%d, value.Channel=%s, value.Name=%s, value.Setting=%s\n", fixture.Group, swiTch.Number, value.Channel, value.Name, value.Setting)
												os.Exit(1)
											}

											// Handle the fact that the channel may be a label as well.
											if IsNumericOnly(value.Channel) {
												c, _ := strconv.ParseFloat(value.Channel, 16)
												dmxController.SetChannel(fixture.Address+int16(c), byte(v))
											} else {
												fixture := findFixtureByName(fixtureName, fixtures)
												c, _ := lookUpChannelNumberByNameInFixtureDefinition(fixture.Group, switchNumber, value.Channel, fixtures)
												dmxController.SetChannel(fixture.Address+int16(c), byte(v))
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

func stopFixture(fixtureName string, fixtures *Fixtures, dmxController *ft232.DMXController) {
	fixture := findFixtureByName(fixtureName, fixtures)
	blackout := false
	master := 255
	strobeSpeed := 0
	ScannerColor := make(map[int]int)
	MapFixtures(fixture.Group-1, dmxController, fixture.Number-1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, ScannerColor, fixtures, blackout, master, master, strobeSpeed)
}

func newMiniSequencer(fixtureName string, switchNumber int, action Action, dmxController *ft232.DMXController, fixturesConfig *Fixtures,
	switchChannels map[int]common.SwitchChannel, soundConfig *sound.SoundConfig,
	blackout bool, master int) {

	switchName := fmt.Sprintf("switch%d", switchNumber)
	fixture := findFixtureByName(fixtureName, fixturesConfig)
	scannerColor := make(map[int]int)

	mySequenceNumber := fixture.Group - 1
	myFixtureNumber := fixture.Number - 1

	var speed time.Duration
	var music int
	var program int
	var rotateSpeed int
	var rotatable bool
	var fade int
	var size int
	var strobe int

	if debug {
		fmt.Printf("Action %+v\n", action)
	}

	switch action.Fade {
	case "soft":
		fade = 1
	case "sharp":
		fade = 10
	default:
		fade = 1
	}

	switch action.Size {
	case "short":
		size = 1
	case "medium":
		size = 3
	case "long":
		size = 10
	default:
		size = 3
	}

	switch action.Program {
	case "all":
		program = 255
	default:
		program = 0
	}

	switch action.Rotate {
	case "off":
		rotateSpeed = 0
		rotatable = false
	case "slow":
		rotateSpeed = 1
		rotatable = true
	case "medium":
		rotateSpeed = 50
		rotatable = true
	case "fast":
		rotateSpeed = 127
		rotatable = true
	default:
		rotateSpeed = 0
		rotatable = false
	}

	switch action.Music {

	case "internal":
		music = 255
	case "off":
		music = 0
	default:
		music = 0
	}

	switch action.Strobe {
	case "off":
		strobe = 0
	case "slow":
		strobe = 0
	case "fast":
		strobe = 0
	default:
		strobe = 0
	}

	if action.Mode == "off" {
		if debug {
			fmt.Printf("Stop mini sequence for switch number %d\n", switchNumber)
		}

		// DeRegister this mini sequencer with the sound service.
		// Use the switch number as the unique sequence name.
		soundConfig.DeRegisterSoundTrigger(switchName)

		select {
		case switchChannels[switchNumber].Stop <- true:
			stopFixture(fixtureName, fixturesConfig, dmxController)
		case <-time.After(100 * time.Millisecond):
		}

		return
	}

	if action.Mode == "static" {
		if debug {
			fmt.Printf("Static mini sequence for switch number %d\n", switchNumber)
		}

		// Remember that we have stopped this mini sequencer.

		select {
		case switchChannels[switchNumber].Stop <- true:
			stopFixture(fixtureName, fixturesConfig, dmxController)
		case <-time.After(100 * time.Millisecond):
		}
		MapFixtures(mySequenceNumber, dmxController, myFixtureNumber, 255, 255, 255, 255, 0, 0, 0, 0, 0, rotateSpeed, music, program, 0, scannerColor, fixturesConfig, blackout, master, master, strobe)
		return
	}

	if action.Mode == "chase" {

		// DeRegister this mini sequencer with the sound service.
		// Use the switch number as the unique sequence name.
		soundConfig.DeRegisterSoundTrigger(switchName)

		select {
		case switchChannels[switchNumber].Stop <- true:
			stopFixture(fixtureName, fixturesConfig, dmxController)
		case <-time.After(100 * time.Millisecond):
		}

		// Register this mini sequencer with the sound service.
		channel := make(chan common.Command)
		trigger := soundConfig.RegisterSoundTrigger(switchName, channel, switchNumber)

		if debug {
			fmt.Printf("Start mini sequence chase for switch number %d\n", switchNumber)
		}

		// Stop any left over sequence left over for this switch.
		select {
		case switchChannels[switchNumber].Stop <- true:
			stopFixture(fixtureName, fixturesConfig, dmxController)
		case <-time.After(100 * time.Millisecond):
		}

		switch action.Speed {
		case "slow":
			speed = 1 * time.Second
			trigger.State = false
		case "medium":
			speed = 500 * time.Millisecond
			trigger.State = false
		case "fast":
			speed = 50 * time.Millisecond
			trigger.State = false
		case "music":
			trigger.State = true
			speed = time.Duration(12 * time.Hour)
		default:
			trigger.State = false
			speed = time.Duration(12 * time.Hour)
		}

		sequence := common.Sequence{
			ScannerInvert: false,
			RGBInvert:     false,
			Bounce:        false,
			ScannerChase:  false,
			RGBShift:      1,
		}
		sequence.Pattern = pattern.MakeSingleFixtureChase()
		sequence.Steps = sequence.Pattern.Steps
		sequence.NumberFixtures = 1
		// Calculate fade curve values.
		slopeOn, slopeOff := common.CalculateFadeValues(fade, size)
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

			if rotatable {
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
							dmxController.SetChannel(fixture.Address+int16(rotateChannel), byte(0))
							return
						case <-switchChannels[switchNumber].KeepRotateAlive:
							time.Sleep(1 * time.Millisecond)
							continue
						case <-time.After(1500 * time.Millisecond):
							dmxController.SetChannel(fixture.Address+int16(rotateChannel), byte(0))
							time.Sleep(250 * time.Millisecond)
							dmxController.SetChannel(fixture.Address+int16(masterChannel), byte(0))
						}
					}
				}(switchNumber, switchChannels)
			}

			// Wait for rotator thread to start.
			time.Sleep(100 * time.Millisecond)

			// find the correct trigger for this switch number.
			var useTriggerNumber int
			for triggerNumber, trigger := range soundConfig.SoundTriggers {
				if trigger.Name == switchName {
					useTriggerNumber = triggerNumber
					break
				}
			}

			for {

				// Run through the steps in the sequence.
				// Remember every step contains infomation for all the fixtures in this group.
				for step := 0; step < sequence.NumberSteps; step++ {

					if rotatable {
						select {
						case switchChannels[switchNumber].KeepRotateAlive <- true:
						case <-time.After(10 * time.Millisecond):
						}
					}

					if rotateCounter > 500 {
						rotateCounter = 1
					}

					if rotateCounter < 128 {
						rotateSpeed = 127
					} else {
						rotateSpeed = 128
					}

					if debug {
						fmt.Printf("switch:%d waiting for beat on %d with speed %d\n", switchNumber, useTriggerNumber, speed)
					}

					// This is were we wait for a beat or a time out equivalent to the speed.
					select {
					case <-soundConfig.SoundTriggers[useTriggerNumber].Channel:
					case <-switchChannels[switchNumber].Stop:
						if rotatable {
							switchChannels[switchNumber].StopRotate <- true
						}
						MapFixtures(mySequenceNumber, dmxController, myFixtureNumber, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, scannerColor, fixturesConfig, blackout, master, master, strobe)
						return
					case <-time.After(speed):
					}

					// Play out fixture to DMX channels.
					position := sequence.RGBPositions[step]

					fixtures := position.Fixtures

					for fixtureNumber := 0; fixtureNumber < sequence.NumberFixtures; fixtureNumber++ {
						fixture := fixtures[fixtureNumber]
						for _, color := range fixture.Colors {
							MapFixtures(mySequenceNumber, dmxController, myFixtureNumber, color.R, color.G, color.B, color.W, 0, 0, 0, 0, 0, rotateSpeed, 0, 0, 0, scannerColor, fixturesConfig, blackout, master, master, strobe)
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

func lightStaticFixture(sequence common.Sequence, myFixtureNumber int, dmxController *ft232.DMXController, eventsForLauchpad chan common.ALight, guiButtons chan common.ALight, fixturesConfig *Fixtures, enabled bool) {

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
	MapFixtures(sequence.Number, dmxController, myFixtureNumber, lamp.Color.R, lamp.Color.G, lamp.Color.B, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, nil, fixturesConfig, sequence.Blackout, sequence.Master, sequence.Master, sequence.StrobeSpeed)

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
func turnOffFixtures(cmd common.FixtureCommand, myFixtureNumber int, mySequenceNumber int, fixtures *Fixtures, dmxController *ft232.DMXController, eventsForLauchpad chan common.ALight, guiButtons chan common.ALight) {
	if !cmd.Hide {
		common.LightLamp(common.ALight{X: myFixtureNumber, Y: mySequenceNumber, Red: 0, Green: 0, Blue: 0, Brightness: 0}, eventsForLauchpad, guiButtons)
	}
	MapFixtures(mySequenceNumber, dmxController, myFixtureNumber, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, cmd.ScannerColor, fixtures, cmd.Blackout, cmd.Master, cmd.Master, cmd.StrobeSpeed)
}

func turnOnFixtures(cmd common.FixtureCommand, myFixtureNumber int, mySequenceNumber int, fixtures *Fixtures, dmxController *ft232.DMXController, eventsForLauchpad chan common.ALight, guiButtons chan common.ALight) {
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

	MapFixtures(mySequenceNumber, dmxController, myFixtureNumber, red, green, blue, white, amber, uv, pan, tilt, shutter, rotate, music, program, gobo, cmd.ScannerColor, fixtures, false, brightness, master, cmd.StrobeSpeed)
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
