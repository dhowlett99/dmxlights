package fixture

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/launchpad"
	"github.com/go-yaml/yaml"
	"github.com/oliread/usbdmx/ft232"
)

type Fixtures struct {
	Fixtures []Fixture `yaml:"fixtures"`
}

type Fixture struct {
	Name        string    `yaml:"name"`
	Description string    `yaml:"desrciption"`
	Group       int       `yaml:"group"`
	Address     int16     `yaml:"address"`
	Channels    []Channel `yaml:"channels"`
}

type Channel struct {
	Number int16  `yaml:"number"`
	Name   string `yaml:"name"`
}

func LoadFixtures() *Fixtures {
	filename := "fixtures.yaml"

	_, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Printf("error loading fixtures.yaml file\n")
	}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("error writing yaml file\n")
	}

	fixtures := &Fixtures{}
	err = yaml.Unmarshal(data, fixtures)
	if err != nil {
		fmt.Printf("error marshalling fixtures: %s\n", err.Error())
	}
	return fixtures
}

// FixtureReceivers are created by the sequence and are used to receive step instructions.
// Each FixtureReceiver knows which step they belong too and when triggered they start a fade up
// and fade down events which get sent to the launchpad lamps and the DMX fixtures.
func FixtureReceiver(sequence common.Sequence,
	mySequenceNumber int,
	myFixtureNumber int,
	channels []chan common.FixtureCommand,
	eventsForLauchpad chan common.ALight,
	dmxController ft232.DMXController,
	fixtures *Fixtures) {

	cmd := common.FixtureCommand{}
	fadeUp := []int{0, 66, 127, 180, 220, 246, 255}
	fadeDown := []int{255, 246, 220, 189, 127, 66, 0}

	for {
		select {
		case cmd = <-channels[myFixtureNumber]:
			if cmd.Tick {
				for _, position := range cmd.Positions {
					if cmd.CurrentPosition == position.StartPosition {
						// Short ciruit the soft fade if we are a scanner.
						if cmd.Type == "scanner" {
							if position.Fixture == myFixtureNumber {

								//fmt.Printf("Type %s\n", cmd.Type)

								//fmt.Printf("Pos %d   fixture %d  Pan is %d  Tilt is %d  Shutter is %d,   Gobo is  %d  Color is %v\n", position.StartPosition, myFixtureNumber, position.Pan, position.Tilt, position.Shutter, position.Gobo, position.Color)
								MapFixtures(mySequenceNumber, dmxController, myFixtureNumber, position.Color.R, position.Color.G, position.Color.B, position.Pan, position.Tilt, position.Shutter, position.Gobo, fixtures, cmd.Blackout, sequence.Master, sequence.Master)
								launchpad.LightLamp(mySequenceNumber, myFixtureNumber, position.Color.R, position.Color.G, position.Color.B, eventsForLauchpad)
							}
						} else {
							if position.Fixture == myFixtureNumber {
								// Now kick off the back end which drives the RGB fixture.
								go func() {
									for _, value := range fadeUp {
										R := int((float64(position.Color.R) / 100) * (float64(value) / 2.55))
										G := int((float64(position.Color.G) / 100) * (float64(value) / 2.55))
										B := int((float64(position.Color.B) / 100) * (float64(value) / 2.55))
										// Pan := position.Pan
										// Tilt := position.Tilt
										// Shutter := position.Shutter
										// Gobo := position.Gobo
										launchpad.LightLamp(mySequenceNumber, myFixtureNumber, R, G, B, eventsForLauchpad)
										// Now ask DMX to actually light the real fixture.
										MapFixtures(mySequenceNumber, dmxController, myFixtureNumber, R, G, B, 0, 0, 0, 0, fixtures, cmd.Blackout, sequence.Master, sequence.Master)
										time.Sleep(cmd.FadeTime / 4) // Fade up Time.
									}
									for x := 0; x < cmd.Size; x++ {
										time.Sleep(cmd.CurrentSpeed * 5)
									}
									time.Sleep(cmd.FadeTime / 4) // Fade on time.
									for _, value := range fadeDown {
										R := int((float64(position.Color.R) / 100) * (float64(value) / 2.55))
										G := int((float64(position.Color.G) / 100) * (float64(value) / 2.55))
										B := int((float64(position.Color.B) / 100) * (float64(value) / 2.55))
										// Pan := position.Pan
										// Tilt := position.Tilt
										// Shutter := position.Shutter
										// Gobo := position.Gobo
										launchpad.LightLamp(mySequenceNumber, myFixtureNumber, R, G, B, eventsForLauchpad)
										MapFixtures(mySequenceNumber, dmxController, myFixtureNumber, R, G, B, 0, 0, 0, 0, fixtures, cmd.Blackout, sequence.Master, sequence.Master)
										time.Sleep(cmd.FadeTime / 4) // Fade down time.
									}
									time.Sleep(cmd.FadeTime / 4) // Fade off time.
								}()
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
func MapFixtures(mySequenceNumber int,
	dmxController ft232.DMXController,
	displayFixture int, R int, G int, B int, Pan int, Tilt int, Shutter int, Gobo int,
	fixtures *Fixtures, blackout bool, brightness int, master int) {

	// We control the brightness of each color with the brightness value.
	// The overall fixture brightness is set from the master value.
	Red := (float64(R) / 100) * (float64(brightness) / 2.55)
	Green := (float64(G) / 100) * (float64(brightness) / 2.55)
	Blue := (float64(B) / 100) * (float64(brightness) / 2.55)

	for _, fixture := range fixtures.Fixtures {

		if fixture.Group == mySequenceNumber {

			//fmt.Printf("fixture  %+v\n\n", fixture)
			//fmt.Printf("fixture.Group %d   mySequenceNumber %d\n\n", fixture.Group, mySequenceNumber)

			//fmt.Printf("found fixture %d\n", fixtureNumber)
			for channelNumber, channel := range fixture.Channels {
				//fmt.Printf("found channel %s\n", channel.Name)
				// Scanner channels
				if strings.Contains(channel.Name, "Pan") {
					//fmt.Printf("DMX debug Pan Channel %d Value %d\n", fixture.Address+int16(channelNumber), Pan)
					dmxController.SetChannel(fixture.Address+int16(channelNumber), byte(Pan))
				}
				if strings.Contains(channel.Name, "Tilt") {
					//fmt.Printf("Tilt is %d\n", Tilt)
					//fmt.Printf("DMX debug Tilt Channel %d Value %d\n", fixture.Address+int16(channelNumber), Tilt)
					dmxController.SetChannel(fixture.Address+int16(channelNumber), byte(Tilt))
				}
				if strings.Contains(channel.Name, "Shutter") {
					//fmt.Printf("DMX debug Shutter Channel %d Value %d\n", fixture.Address+int16(channelNumber), Shutter)
					dmxController.SetChannel(fixture.Address+int16(channelNumber), byte(Shutter))
				}
				if strings.Contains(channel.Name, "Gobo") {
					//fmt.Printf("DMX debug Gobo Channel %d Value %d\n", fixture.Address+int16(channelNumber), Gobo)
					dmxController.SetChannel(fixture.Address+int16(channelNumber), byte(Gobo))
				}

				// Fixture channels.
				if strings.Contains(channel.Name, "Red"+strconv.Itoa(displayFixture+1)) {
					//fmt.Printf("DMX debug Channel %d Value %d\n", fixture.Address+int16(channelNumber), R)
					dmxController.SetChannel(fixture.Address+int16(channelNumber), byte(int(Red)))
				}
				if strings.Contains(channel.Name, "Green"+strconv.Itoa(displayFixture+1)) {
					//fmt.Printf("DMX debug Channel %d Value %d\n", fixture.Address+int16(channelNumber), G)
					dmxController.SetChannel(fixture.Address+int16(channelNumber), byte(int(Green)))
				}
				if strings.Contains(channel.Name, "Blue"+strconv.Itoa(displayFixture+1)) {
					//fmt.Printf("DMX debug Channel %d Value %d\n", fixture.Address+int16(channelNumber), B)
					dmxController.SetChannel(fixture.Address+int16(channelNumber), byte(int(Blue)))
				}
				if strings.Contains(channel.Name, "Master") {
					//fmt.Printf("DMX debug Channel %d Value %d\n", fixture.Address+int16(channelNumber), 255)
					if blackout {
						dmxController.SetChannel(fixture.Address+int16(channelNumber), byte(0))
					} else {
						dmxController.SetChannel(fixture.Address+int16(channelNumber), byte(master))
					}
				}
			}
		}
	}
}
