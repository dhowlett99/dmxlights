// Copyright (C) 2022,2023,2024 dhowlett99.
// This is the dmxlights mini settings player, used by the fixture to control
// settings for fixtures.
// Implemented and depends usbdmx.
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
	"fmt"
	"strconv"
	"strings"

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/oliread/usbdmx/ft232"
)

// Process settings.
func newMiniSetter(thisFixture *Fixture, setting common.Setting, masterChannel int,
	dmxController *ft232.DMXController,
	master int,
	dmxInterfacePresent bool) {

	debug := true

	if debug {
		fmt.Printf("settings are available\n")
	}

	// Not Blackout.
	// This should be controlled by the master brightness
	settingName := strings.ToLower(setting.Name)
	channelName := strings.ToLower(setting.Channel)
	if strings.Contains(settingName, "master") ||
		strings.Contains(settingName, "dimmer") ||
		strings.Contains(channelName, "master") ||
		strings.Contains(channelName, "dimmer") {

		// Master brightness.
		value, _ := strconv.ParseFloat(setting.FixtureValue, 32)

		howBright := int((float64(value) / 100) * (float64(master) / 2.55))
		if debug {
			fmt.Printf("Fixture %s setting value %d master %d howBright %d\n", thisFixture.Name, int(value), master, howBright)
		}

		if strings.Contains(settingName, "reverse") || strings.Contains(settingName, "invert") {
			// Invert the brightness value,  some fixtures have the max brightness at 0 and off at 255.
			if debug {
				fmt.Printf("fixture %s --->:Control: send Setting %s Address %d Value %d \n", thisFixture.Name, setting.Name, thisFixture.Address+int16(masterChannel), int(howBright))
			}
			SetChannel(thisFixture.Address+int16(masterChannel), byte(reverse_dmx(howBright)), dmxController, dmxInterfacePresent)
		} else {
			// Set the master brightness value.
			if debug {
				fmt.Printf("fixture %s --->:Control: send Setting %s Address %d Value %d \n", thisFixture.Name, setting.Name, thisFixture.Address+int16(masterChannel), int(howBright))
			}
			SetChannel(thisFixture.Address+int16(masterChannel), byte(howBright), dmxController, dmxInterfacePresent)
		}

	} else {

		// If the setting value has is a number set it directly.
		if IsNumericOnly(setting.FixtureValue) {

			value, _ := strconv.Atoi(setting.FixtureValue)
			if IsNumericOnly(setting.Channel) {
				channel, _ := strconv.Atoi(setting.Channel)
				channel = channel - 1 // Channels are relative to the base address so deduct one to make it work.
				if debug {
					fmt.Printf("fixture %s --->:Control: send Setting %s Address %d Value %d \n", thisFixture.Name, setting.Name, thisFixture.Address+int16(channel), value)
				}
				SetChannel(thisFixture.Address+int16(channel), byte(value), dmxController, dmxInterfacePresent)
			} else {
				// Handle the fact that the channel may be a label as well.
				// Look for this channels number in this fixture identified by ID.
				channel, _ := FindChannelNumberByName(thisFixture, setting.Channel)
				if debug {
					fmt.Printf("fixture %s --->:Control: send Setting %s Address %d Value %d \n", thisFixture.Name, setting.Name, thisFixture.Address+int16(channel), value)
				}
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
			fmt.Printf("Handle the fact that the channel may be a label as well.\n")
			if IsNumericOnly(setting.Channel) {
				// Find the channel
				channel, _ := strconv.ParseFloat(setting.Channel, 32)
				//if debug {
				fmt.Printf("fixture %s --->:Control: send Setting %s Address %d Value %d \n", thisFixture.Name, setting.Name, thisFixture.Address+int16(channel), value)
				//}
				SetChannel(thisFixture.Address+int16(channel), byte(value), dmxController, dmxInterfacePresent)
			} else {
				// Look for this channels number in this fixture identified by ID.
				channel, _ := FindChannelNumberByName(thisFixture, setting.Channel)
				//if debug {
				fmt.Printf("fixture %s --->:Control: send Setting %s Address %d Value %d \n", thisFixture.Name, setting.Name, thisFixture.Address+int16(channel), value)
				//}
				SetChannel(thisFixture.Address+int16(channel), byte(value), dmxController, dmxInterfacePresent)
			}
		}
	}
}
