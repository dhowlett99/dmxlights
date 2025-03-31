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
	"math"
	"strconv"
	"strings"

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/oliread/usbdmx/ft232"
)

// Process settings.

// newMiniSetter -The mini setter takes a bunch of settings/values from a switch state and applies them to the
// DMX universe, using SetChannel primitive directly.
// It uses the the settings channel name to identify the channel to use in the fixture.
func newMiniSetter(thisFixture *Fixture, override *common.Override, setting common.Setting, masterChannel int,
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
				channel, err := GetChannelNumberByName(thisFixture, setting.Channel)
				if err != nil {
					fmt.Printf("error fixure %s has no channel %s\n", thisFixture.Name, setting.Channel)
					return
				}
				if debug_mini_setter {
					fmt.Printf("fixture %s: ChannelLabel Control: Channel=%s send Setting=%s Address=%d Value=%d\n", thisFixture.Name, setting.Channel, setting.Name, thisFixture.Address+int16(channel), value)
				}

				// Override settings code.
				var overrideHasHappened bool

				// Override Shutter.
				if setting.Channel == "Shutter" && (override.Shutter || override.Strobe) {

					if override.Strobe {
						strobeValues := GetADMXValueMaxMin(thisFixture, "Strobe", setting.Channel)
						shutter := makeStrobeSpeed(strobeValues, override.StrobeSpeed)
						if debug_mini_setter {
							fmt.Printf("Override is set Address=%d Strobe Shutter=%t DMX Value=%d\n", thisFixture.Address+int16(channel), override.Shutter, shutter)
						}
						SetChannel(thisFixture.Address+int16(channel), byte(shutter), dmxController, dmxInterfacePresent)
					} else {
						shutter := GetADMXValueByName(thisFixture, "Open", "Shutter")
						if debug_mini_setter {
							fmt.Printf("Override is set Address=%d Open Shutter=%t DMX Value=%d\n", thisFixture.Address+int16(channel), override.Shutter, shutter)
						}
						SetChannel(thisFixture.Address+int16(channel), byte(shutter), dmxController, dmxInterfacePresent)
					}
					overrideHasHappened = true
					override.Strobe = false
					override.Shutter = false
				}

				// Override Strobe.
				if setting.Channel == "Strobe" && override.Strobe {
					if debug_mini_setter {
						fmt.Printf("Override is set Address=%d Strobe=%t StrobeSpeed=%d DMX Value=%d\n", thisFixture.Address+int16(channel), override.Strobe, override.StrobeSpeed, override.StrobeSpeed)
					}
					SetChannel(thisFixture.Address+int16(channel), byte(override.StrobeSpeed), dmxController, dmxInterfacePresent)
					overrideHasHappened = true
					override.Strobe = false
				}

				// Override Speed.
				if setting.Channel == "Speed" && override.OverrideSpeed {
					speed := GetADMXValue(thisFixture, override.Speed, "Speed")
					if debug_mini_setter {
						fmt.Printf("Override is set Address=%d Speed=%d DMX Value=%d\n", thisFixture.Address+int16(channel), override.Speed, speed)
					}
					SetChannel(thisFixture.Address+int16(channel), byte(override.Speed), dmxController, dmxInterfacePresent)
					overrideHasHappened = true
					override.OverrideSpeed = false
				}

				// Override Shift.
				if setting.Channel == "Shift" && override.OverrideShift {
					shift := GetADMXValue(thisFixture, override.Rotate, "Shift")
					if debug_mini_setter {
						fmt.Printf("Override is set Address=%d Shift=%d DMX Value=%d\n", thisFixture.Address+int16(channel), override.Shift, shift)
					}
					SetChannel(thisFixture.Address+int16(channel), byte(shift), dmxController, dmxInterfacePresent)
					overrideHasHappened = true
					override.OverrideShift = false
				}

				// Override Size.
				if setting.Channel == "Size" && override.OverrideSize {
					size := GetADMXValue(thisFixture, override.Rotate, "Size")
					if debug_mini_setter {
						fmt.Printf("Override is set Address=%d Size=%d DMX Value=%d\n", thisFixture.Address+int16(channel), override.Size, size)
					}
					SetChannel(thisFixture.Address+int16(channel), byte(size), dmxController, dmxInterfacePresent)
					overrideHasHappened = true
					override.OverrideSize = false
				}

				// Override Fade
				if setting.Channel == "Fade" && override.OverrideFade {
					fade := GetADMXValue(thisFixture, override.Fade, "Fade")
					if debug_mini_setter {
						fmt.Printf("Override is set Address=%d Fade=%d\n DMX Value=%d\n", thisFixture.Address+int16(channel), override.Fade, fade)
					}
					SetChannel(thisFixture.Address+int16(channel), byte(fade), dmxController, dmxInterfacePresent)
					overrideHasHappened = true
					override.OverrideFade = false
				}

				// Override ProgramSpeed.
				if setting.Channel == "ProgramSpeed" && override.OverrideProgramSpeed {
					programSpeed := GetADMXValue(thisFixture, override.Speed, "ProgramSpeed")
					if debug_mini_setter {
						fmt.Printf("Override is set Address=%d ProgramSpeed=%d DMX Value=%d\n", thisFixture.Address+int16(channel), override.ProgramSpeed, programSpeed)
					}
					SetChannel(thisFixture.Address+int16(channel), byte(override.ProgramSpeed), dmxController, dmxInterfacePresent)
					overrideHasHappened = true
					override.OverrideProgramSpeed = false
				}

				// Override RotateSpeed.
				if setting.Channel == "Rotate" && override.OverrideRotateSpeed {
					rotate := GetADMXValue(thisFixture, override.Rotate, "Rotate")
					if debug_mini_setter {
						fmt.Printf("Override is set Address=%d Rotate Speed=%d  DMX Value=%d\n", thisFixture.Address+int16(channel), override.Rotate, rotate)
					}
					SetChannel(thisFixture.Address+int16(channel), byte(rotate), dmxController, dmxInterfacePresent)
					overrideHasHappened = true
					override.OverrideRotateSpeed = false
				}

				// Override Color.
				if setting.Channel == "Color" && override.OverrideColors {
					color := GetADMXValue(thisFixture, override.Color, "Color")
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
					gobo := GetADMXValue(thisFixture, override.Gobo, "Gobo")
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
						fmt.Printf("Set Channel Channel %d Value %d\n", thisFixture.Address+int16(channel), byte(value))
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
				channel, _ := GetChannelNumberByName(thisFixture, setting.Channel)
				if debug_mini_setter {
					fmt.Printf("fixture %s: SettingisID Control: Channel=%s send Setting=%s Address=%d Value=%d\n", thisFixture.Name, setting.Channel, setting.Name, thisFixture.Address+int16(channel), value)
				}
				SetChannel(thisFixture.Address+int16(channel), byte(value), dmxController, dmxInterfacePresent)
			}
		}
	}
}

// makeStrobeSpeed takes a range of strobe values and the current strobe speed value,
// returns the strobe speed as a DMX value.
func makeStrobeSpeed(strobeValues []int, strobeSpeed int) int {

	var minValue float32
	var maxValue float32
	var step float32
	var stepNumber float32
	var slotSize float32
	var groupIndex int

	var MAX_STROBE_SPEED = float32(common.MAX_STROBE_SPEED)

	dmxValue := -1

	if len(strobeValues) > 0 {
		minValue = float32(strobeValues[0])
		maxValue = float32(strobeValues[1])
		speed := strobeSpeed
		slotSize = MAX_STROBE_SPEED / 10

		numberOfValues := maxValue - minValue
		inc := numberOfValues / 9

		groups := getGroups(10, 255)
		if debug {
			fmt.Printf("Groups %+v\n", groups)
		}

		// Populate array.
		if debug {
			fmt.Printf("Populate array.\n")
		}
		for step = minValue; step < maxValue; step += inc {

			s := float32(math.Round(float64(step)))
			if debug {
				fmt.Printf("stepNumber %d dmx %d speed %d\n", int(stepNumber), int(s), int(speed))
			}
			dmxValue = int(s)
			if int(stepNumber) >= groups[groupIndex].Min && int(stepNumber) <= groups[groupIndex].Max {
				groups[groupIndex].DMXValue = dmxValue
			}

			stepNumber += slotSize
			groupIndex++
		}

		// Print array.
		if debug {
			fmt.Printf("Print array.\n")

			for groupIndex := 0; groupIndex < len(groups); groupIndex++ {
				fmt.Printf("min %d max %d dmx %d\n", groups[groupIndex].Min, groups[groupIndex].Max, groups[groupIndex].DMXValue)
			}
		}

		// Find value based on index.
		if debug {
			fmt.Printf("Print array.\n")
		}
		for groupIndex := 0; groupIndex < len(groups); groupIndex++ {
			if speed > groups[groupIndex].Min-1 && speed < groups[groupIndex].Max-1 {
				if debug {
					fmt.Printf("Found DMX value %d\n", groups[groupIndex].DMXValue)
				}
				return groups[groupIndex].DMXValue
			}
		}
	}
	return dmxValue
}

type MaxMins struct {
	Min      int
	Max      int
	DMXValue int
}

func getGroups(patternSize int, numLights int) []MaxMins {
	width := int(math.Floor(float64(numLights) / float64(patternSize)))
	groups := make([]MaxMins, patternSize)

	for i := 0; i < int(patternSize); i++ {
		groups[i] = MaxMins{
			Min: i * width,
			Max: (i + 1) * width,
		}
	}
	if debug {
		fmt.Printf("Width %d Groups %d\n", width, groups)
	}

	return groups
}
