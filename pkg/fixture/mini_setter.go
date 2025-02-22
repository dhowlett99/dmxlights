// Copyright (C) 2022,2023,2024,2025 dhowlett99.
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
func newMiniSetter(thisFixture *Fixture, override common.Override, setting common.Setting, masterChannel int,
	dmxController *ft232.DMXController,
	master int,
	dmxInterfacePresent bool) {

	debug_mini_setter := false

	if debug_mini_setter {
		fmt.Printf("newMiniSetter: settings are available fixture %s settingName %s\n", thisFixture.Name, setting.Name)
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
		if debug_mini_setter {
			fmt.Printf("newMiniSetter: Fixture %s setting value %d master %d howBright %d\n", thisFixture.Name, int(value), master, howBright)
		}

		// If we are using the master reverse feature, we can label it in the setting name or the channel name.
		if strings.Contains(settingName, "reverse") ||
			strings.Contains(settingName, "invert") ||
			strings.Contains(channelName, "reverse") ||
			strings.Contains(channelName, "invert") {

			// Invert the brightness value,  some fixtures have the max brightness at 0 and off at 255.
			if debug_mini_setter {
				fmt.Printf("fixture %s: Control: send Setting %s Address %d Value %d \n", thisFixture.Name, setting.Name, thisFixture.Address+int16(masterChannel), int(howBright))
			}
			SetChannel(thisFixture.Address+int16(masterChannel), byte(reverse_dmx(howBright)), dmxController, dmxInterfacePresent)
		} else {
			// Set the master brightness value.
			if debug_mini_setter {
				fmt.Printf("fixture %s: Control: send Setting %s Address %d Value %d \n", thisFixture.Name, setting.Name, thisFixture.Address+int16(masterChannel), int(howBright))
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
				if debug_mini_setter {
					fmt.Printf("fixture %s: IsNumber Control: Channel=%s send Setting=%s Address=%d Value=%d\n", thisFixture.Name, setting.Channel, setting.Name, thisFixture.Address+int16(channel), value)
				}
				SetChannel(thisFixture.Address+int16(channel), byte(value), dmxController, dmxInterfacePresent)
			} else {
				// Handle the fact that the channel may be a label as well.
				// Look for this channels number in this fixture identified by ID.
				channel, _ := FindChannelNumberByName(thisFixture, setting.Channel)
				if debug_mini_setter {
					fmt.Printf("fixture %s: ChannelLabel Control: Channel=%s send Setting=%s Address=%d Value=%d\n", thisFixture.Name, setting.Channel, setting.Name, thisFixture.Address+int16(channel), value)
				}

				// Override settings code.
				var overrideHasHappened bool

				// Override Speed.
				if setting.Channel == "Speed" && override.OverrideSpeed {
					if debug_mini_setter {
						fmt.Printf("Override is set Address=%d Speed=%d\n", thisFixture.Address+int16(channel), override.Speed)
					}
					SetChannel(thisFixture.Address+int16(channel), byte(override.Speed), dmxController, dmxInterfacePresent)
					overrideHasHappened = true
					override.OverrideSpeed = false
				}

				// Override Shift.
				if setting.Channel == "RotateSpeed" && override.OverrideShift {
					if debug_mini_setter {
						fmt.Printf("Override is set Address=%d Shift=%d\n", thisFixture.Address+int16(channel), override.Shift)
					}
					SetChannel(thisFixture.Address+int16(channel), byte(override.Shift), dmxController, dmxInterfacePresent)
					overrideHasHappened = true
					override.OverrideShift = false
				}

				// Override Size.
				if setting.Channel == "Size" && override.OverrideSize {
					if debug_mini_setter {
						fmt.Printf("Override is set Address=%d Size=%d\n", thisFixture.Address+int16(channel), override.Size)
					}
					SetChannel(thisFixture.Address+int16(channel), byte(override.Size), dmxController, dmxInterfacePresent)
					overrideHasHappened = true
					override.OverrideSize = false
				}

				// Override Fade
				if setting.Channel == "Fade" && override.OverrideFade {
					if debug_mini_setter {
						fmt.Printf("Override is set Address=%d Fade=%d\n", thisFixture.Address+int16(channel), override.Fade)
					}
					SetChannel(thisFixture.Address+int16(channel), byte(override.Fade), dmxController, dmxInterfacePresent)
					overrideHasHappened = true
					override.OverrideFade = false
				}

				// Override RotateSpeed.
				if setting.Channel == "RotateSpeed" && override.OverrideRotateSpeed {
					if debug_mini_setter {
						fmt.Printf("Override is set Address=%d Rotate Speed=%d\n", thisFixture.Address+int16(channel), override.Rotate)
					}
					SetChannel(thisFixture.Address+int16(channel), byte(override.Rotate), dmxController, dmxInterfacePresent)
					overrideHasHappened = true
					override.OverrideRotateSpeed = false
				}

				// Override Color.
				if setting.Channel == "Color" && override.OverrideColors {
					color := GetColorDMXValueByNumber(thisFixture, override.Color)
					if debug_mini_setter {
						fmt.Printf("Override is set Address=%d ColorsNumber=%d DMX Value=%d\n", thisFixture.Address+int16(channel), override.Color, color)
					}
					SetChannel(thisFixture.Address+int16(channel), byte(color), dmxController, dmxInterfacePresent)
					overrideHasHappened = true
					override.OverrideColors = false
				}

				// Override Gobo.
				if setting.Channel == "Gobo" && override.OverrideGobo {
					// Lookup correct value for this Gobo number.
					gobo := GetGoboDMXValueByNumber(thisFixture, override.Gobo)
					if debug_mini_setter {
						fmt.Printf("Override is set Address=%d GoboNumber=%d  DMX Value=%d\n", thisFixture.Address+int16(channel), override.Gobo, gobo)
					}
					SetChannel(thisFixture.Address+int16(channel), byte(gobo), dmxController, dmxInterfacePresent)
					overrideHasHappened = true
					override.OverrideGobo = false
				}

				// Not overriding this channel, so use value from config.
				if !overrideHasHappened {
					if debug_mini_setter {
						fmt.Printf("Set Channel Channel %ds Value %d\n", thisFixture.Address+int16(channel), byte(value))
					}
					SetChannel(thisFixture.Address+int16(channel), byte(value), dmxController, dmxInterfacePresent)
				}

			}

		} else {
			// The setting value is a string.
			// Lets looks to see if the setting string contains a valid fixture label, look up the label in the fixture definition.
			value, err := findChannelSettingByLabel(thisFixture, setting.Channel, setting.Label)
			if err != nil {
				fmt.Printf("error: %s\n", err.Error())
			}

			// Handle the fact that the channel may be a label as well.
			if IsNumericOnly(setting.Channel) {
				// Find the channel
				channel, _ := strconv.ParseFloat(setting.Channel, 32)
				if debug_mini_setter {
					fmt.Printf("fixture %s: SettingisValue Control: Channel=%s send Setting=%s Address=%d Value=%d\n", thisFixture.Name, setting.Channel, setting.Name, thisFixture.Address+int16(channel), value)
				}
				SetChannel(thisFixture.Address+int16(channel), byte(value), dmxController, dmxInterfacePresent)
			} else {
				// Look for this channels number in this fixture identified by ID.
				channel, _ := FindChannelNumberByName(thisFixture, setting.Channel)
				if debug_mini_setter {
					fmt.Printf("fixture %s: SettingisID Control: Channel=%s send Setting=%s Address=%d Value=%d\n", thisFixture.Name, setting.Channel, setting.Name, thisFixture.Address+int16(channel), value)
				}
				SetChannel(thisFixture.Address+int16(channel), byte(value), dmxController, dmxInterfacePresent)
			}
		}
	}
}
