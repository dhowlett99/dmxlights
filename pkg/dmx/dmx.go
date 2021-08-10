package dmx

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/dhowlett99/dmxlights/pkg/fixture"
	"github.com/oliread/usbdmx"
	"github.com/oliread/usbdmx/ft232"
)

func NewDmXController() (controller ft232.DMXController) {
	// Constants, these should really be defined in the module and will be
	// as of the next release
	vid := uint16(0x0403)
	pid := uint16(0x6001)
	outputInterfaceID := 2
	inputInterfaceID := 1
	debugLevel := 0

	// Create a configuration from our flags
	config := usbdmx.NewConfig(vid, pid, outputInterfaceID, inputInterfaceID, debugLevel)

	// Get a usb context for our configuration
	config.GetUSBContext()

	// Create a controller and connect to it
	controller = ft232.NewDMXController(config)
	if err := controller.Connect(); err != nil {
		log.Fatalf("Failed to connect DMX Controller: %s", err)
	}

	// Create a go routine that will ensure our controller keeps sending data
	// to our fixture with a short delay. No delay, or too much delay, may cause
	// flickering in fixtures. Check the specification of your fixtures and controller
	go func(c *ft232.DMXController) {
		for {
			if err := controller.Render(); err != nil {
				log.Fatalf("Failed to render output: %s", err)
			}

			time.Sleep(30 * time.Millisecond)
		}
	}(&controller)

	return controller
}

func SetDMXChannel(controller ft232.DMXController, channel int16, value byte) {
	controller.SetChannel(channel, value)
}

func Fixtures(mySequenceNumber int, dmxController ft232.DMXController, displayFixture int, R int, G int, B int, Pan int, Tilt int, Shutter int, Gobo int, fixtures *fixture.Fixtures, blackout bool) {

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
					dmxController.SetChannel(fixture.Address+int16(channelNumber), byte(R))
				}
				if strings.Contains(channel.Name, "Green"+strconv.Itoa(displayFixture+1)) {
					//fmt.Printf("DMX debug Channel %d Value %d\n", fixture.Address+int16(channelNumber), G)
					dmxController.SetChannel(fixture.Address+int16(channelNumber), byte(G))
				}
				if strings.Contains(channel.Name, "Blue"+strconv.Itoa(displayFixture+1)) {
					//fmt.Printf("DMX debug Channel %d Value %d\n", fixture.Address+int16(channelNumber), B)
					dmxController.SetChannel(fixture.Address+int16(channelNumber), byte(B))
				}
				if strings.Contains(channel.Name, "Master") {
					//fmt.Printf("DMX debug Channel %d Value %d\n", fixture.Address+int16(channelNumber), 255)
					if blackout {
						dmxController.SetChannel(fixture.Address+int16(channelNumber), byte(0))
					} else {
						dmxController.SetChannel(fixture.Address+int16(channelNumber), byte(255))
					}
				}
			}
		}
	}

	// Controller how long the fixture remains on, smaller numbers
	// Give a more dramatic show.
	// time.Sleep(20 * time.Millisecond)
}
