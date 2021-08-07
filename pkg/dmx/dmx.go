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

func Fixtures(mySequenceNumber int, dmxController ft232.DMXController, displayFixture int, R int, G int, B int, groups *fixture.Groups) {
	R = convertToDMXValues(R)
	G = convertToDMXValues(G)
	B = convertToDMXValues(B)

	// The sequence number passed in is directly mapped to the groups.
	for groupNumber, group := range groups.Groups {
		if mySequenceNumber-1 == groupNumber {
			//fmt.Printf("found group %d\n", groupNumber)
			for _, fixture := range group.Fixtures {
				//fmt.Printf("Base Address %d\n", fixture.Address)

				//fmt.Printf("found fixture %d\n", fixtureNumber)
				for channelNumber, channel := range fixture.Channels {
					//fmt.Printf("No %d\n", channel.Number)
					//fmt.Printf("Name %s Display Fixture %d\n", channel.Name, displayFixture+1)
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
						dmxController.SetChannel(fixture.Address+int16(channelNumber), 255)
					}
				}
			}
		}
	}

	// Now ask DMX to actually light the real fixture.
	// if fixture == 0 {
	// 	dmxController.SetChannel(1, byte(R))
	// 	dmxController.SetChannel(2, byte(G))
	// 	dmxController.SetChannel(3, byte(B))
	// }
	// if fixture == 1 {
	// 	dmxController.SetChannel(4, byte(R))
	// 	dmxController.SetChannel(5, byte(G))
	// 	dmxController.SetChannel(6, byte(B))
	// }
	// if fixture == 2 {
	// 	dmxController.SetChannel(7, byte(R))
	// 	dmxController.SetChannel(8, byte(G))
	// 	dmxController.SetChannel(9, byte(B))
	// }
	// if fixture == 3 {
	// 	dmxController.SetChannel(10, byte(R))
	// 	dmxController.SetChannel(11, byte(G))
	// 	dmxController.SetChannel(12, byte(B))
	// }
	// dmxController.SetChannel(13, 255)
	// if fixture == 4 {
	// 	dmxController.SetChannel(14, byte(R))
	// 	dmxController.SetChannel(15, byte(G))
	// 	dmxController.SetChannel(16, byte(B))
	// }
	// if fixture == 5 {
	// 	dmxController.SetChannel(17, byte(R))
	// 	dmxController.SetChannel(18, byte(G))
	// 	dmxController.SetChannel(19, byte(B))
	// }
	// if fixture == 6 {
	// 	dmxController.SetChannel(20, byte(R))
	// 	dmxController.SetChannel(21, byte(G))
	// 	dmxController.SetChannel(22, byte(B))
	// }
	// if fixture == 7 {
	// 	dmxController.SetChannel(23, byte(R))
	// 	dmxController.SetChannel(24, byte(G))
	// 	dmxController.SetChannel(25, byte(B))
	// }
	// dmxController.SetChannel(26, 255)
	// Controller how long the fixture remains on, smaller numbers
	// Give a more dramatic show.
	time.Sleep(20 * time.Millisecond)
}

func convertToDMXValues(input int) (output int) {

	if input == 0 {
		output = 0
	}
	if input == 1 {
		output = 85
	}
	if input == 2 {
		output = 170
	}
	if input == 3 {
		output = 255
	}

	return output
}
