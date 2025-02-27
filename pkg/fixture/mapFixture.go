// Copyright (C) 2022,2023,2024,2025 dhowlett99.
// This is the dmxlights fixture mapping code, it is called to find fixtures and
// then sends messages to fixtures using the usb dmx library.
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
	"fmt"
	"image/color"
	"strconv"
	"strings"

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/oliread/usbdmx/ft232"
)

// When want to light a DMX fixture we need to find it in our fuxture.yaml configuration file.
// This function maps the requested fixture into a DMX address.
func MapFixtures(chaser bool, hasShutterChaser bool,
	mySequenceNumber int,
	displayFixture int,
	color color.RGBA,
	baseColorName color.RGBA,
	pan int, tilt int, shutter int, rotate int, program int, selectedGobo int, scannerColor int,
	fixtures *Fixtures, blackout bool, brightness int, master int, music int, strobe bool, strobeSpeed int,
	dmxController *ft232.DMXController, dmxInterfacePresent bool) (lastColor common.LastColor) {

	if debug {
		fmt.Printf("MapFixtures Fixture No %d Sequence No %d color %+v\n", displayFixture, mySequenceNumber, color)
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
						if !hasShutterChaser {
							if strings.Contains(channel.Name, "Gobo") {
								for _, setting := range channel.Settings {
									if setting.Number == selectedGobo {
										v, _ := strconv.Atoi(setting.Value)
										SetChannel(fixture.Address+int16(channelNumber), byte(v), dmxController, dmxInterfacePresent)
									}
								}
							}
						}
						if !hasShutterChaser {
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
						if !hasShutterChaser {
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
					// If the fixure supports red, green and blue channels we can set the color directly.
					if fixture.HasRGBChannels {
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
					} else {
						// Set the color using the original actions base color.
						if strings.Contains(channel.Name, "Color") {
							baseColor := common.GetColorNameByRGB(baseColorName)
							// Look for a setting that matches the color.
							for _, setting := range channel.Settings {
								if setting.Name == baseColor {
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

	lastColor.RGBColor = color
	lastColor.ScannerColor = scannerColor

	return lastColor
}
