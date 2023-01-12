// Copyright (C) 2022 dhowlett99.
// This is the dmxlights fixture editor it is attached to a fixture and
// describes the fixtures properties which is then saved in the fixtures.yaml
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

package dmx

import (
	"errors"
	"log"
	"time"

	"github.com/oliread/usbdmx"
	"github.com/oliread/usbdmx/ft232"
)

func NewDmXController() (*ft232.DMXController, *usbdmx.ControllerConfig, error) {
	// Constants, these should really be defined in the module and will be
	// as of the next release
	vid := uint16(0x0403) // Future Technology Devices International Limited
	pid := uint16(0x6001)
	outputInterfaceID := 2
	inputInterfaceID := 1
	debugLevel := 0

	controller := ft232.DMXController{}

	// Create a configuration from our flags
	config := usbdmx.NewConfig(vid, pid, outputInterfaceID, inputInterfaceID, debugLevel)

	// Get a usb context for our configuration
	config.GetUSBContext()

	// Create a controller and connect to it
	controller = ft232.NewDMXController(config)
	err := controller.Connect()
	if err != nil {
		return nil, nil, errors.New("failed to connect DMX Controller: " + err.Error())
	}

	// Create a go routine that will ensure our controller keeps sending data
	// to our fixture with a short delay. No delay, or too much delay, may cause
	// flickering in fixtures. Check the specification of your fixtures and controller
	go func(c *ft232.DMXController) {
		for {
			if err := controller.Render(); err != nil {
				log.Fatalf("Failed to render output: %s", err)
			}
			// DMX refresh rate.
			time.Sleep(30 * time.Millisecond)
		}
	}(&controller)

	return &controller, &config, nil
}
