package dmx

import (
	"log"
	"time"

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

func Fixtures(dmxController ft232.DMXController, fixture int, R int, G int, B int) {
	R = convertToDMXValues(R)
	G = convertToDMXValues(G)
	B = convertToDMXValues(B)

	// Now ask DMX to actually light the real fixture.
	if fixture == 0 {
		dmxController.SetChannel(1, byte(R))
		dmxController.SetChannel(2, byte(G))
		dmxController.SetChannel(3, byte(B))
	}
	if fixture == 1 {
		dmxController.SetChannel(4, byte(R))
		dmxController.SetChannel(5, byte(G))
		dmxController.SetChannel(6, byte(B))
	}
	if fixture == 2 {
		dmxController.SetChannel(7, byte(R))
		dmxController.SetChannel(8, byte(G))
		dmxController.SetChannel(9, byte(B))
	}
	if fixture == 3 {
		dmxController.SetChannel(10, byte(R))
		dmxController.SetChannel(11, byte(G))
		dmxController.SetChannel(12, byte(B))
	}
	dmxController.SetChannel(13, 20)
	if fixture == 4 {
		dmxController.SetChannel(14, byte(R))
		dmxController.SetChannel(15, byte(G))
		dmxController.SetChannel(16, byte(B))
	}
	if fixture == 5 {
		dmxController.SetChannel(17, byte(R))
		dmxController.SetChannel(18, byte(G))
		dmxController.SetChannel(19, byte(B))
	}
	if fixture == 6 {
		dmxController.SetChannel(20, byte(R))
		dmxController.SetChannel(21, byte(G))
		dmxController.SetChannel(22, byte(B))
	}
	if fixture == 7 {
		dmxController.SetChannel(23, byte(R))
		dmxController.SetChannel(24, byte(G))
		dmxController.SetChannel(25, byte(B))
	}
	dmxController.SetChannel(26, 20)
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
