// Copyright (C) 2022,2025 dhowlett99.
// This is the dmxlights launchpad interface.
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
package launchpad

import (
	"fmt"
	"time"

	"github.com/dhowlett99/dmxlights/pkg/buttons"
	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
	"github.com/dhowlett99/pad"
	"github.com/oliread/usbdmx/ft232"
)

const debug = false

func NewLaunchPad() (*pad.Pad, error) {

	// Setup a connection to the Novation Launchpad.
	// Tested with a Novation Launchpad mini pad.
	pad, err := pad.Open()
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	return pad, nil
}

// main thread is used to get commands from the lauchpad.
func ReadLaunchPadButtons(guiButtons chan common.ALight, this *buttons.CurrentState, sequences []*common.Sequence,
	eventsForLaunchpad chan common.ALight, dmxController *ft232.DMXController,
	fixturesConfig *fixture.Fixtures, commandChannels []chan common.Command,
	replyChannels []chan common.Sequence, updateChannels []chan common.Sequence,
	dmxInterfaceCardPresent bool) {

	// Create a channel to listen for buttons being pressed.
	// Send the button pressed hit to the button channel.
	buttonChannel := make(chan pad.Hit)
	go func() {
		this.Pad.Listen(buttonChannel)
	}()

	// Main loop reading commands from the Novation Launchpad.
	for {
		hit := <-buttonChannel
		buttons.ProcessButtons(hit.X, hit.Y, sequences, this, eventsForLaunchpad, guiButtons, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels, false)
	}
}

type coordinate struct {
	X int
	Y int
}

// ListenAndSendToLaunchPad is the thread that listens for events to send to
// the launch pad.  It is thread safe and is the only thread talking to the
// launch pad. A channel is used to queue the events to be sent.
func ListenAndSendToLaunchPad(eventsForLauchpad chan common.ALight, pad *pad.Pad, LaunchPadConnected bool) {

	launchPadMap := make(map[coordinate]common.Color, 81)

	for key := range launchPadMap {
		launchPadMap[key] = common.Color{R: 0, G: 0, B: 0}
	}

	for {
		// Wait for the event.
		alight := <-eventsForLauchpad

		if LaunchPadConnected {
			// Wait for a few millisecond so the launchpad and the gui step at the same time
			time.Sleep(14 * time.Microsecond)

			// What was the lamp previously set too ?
			whichLamp := coordinate{X: alight.X, Y: alight.Y}
			storedColor := common.Color{R: launchPadMap[whichLamp].R, G: launchPadMap[whichLamp].G, B: launchPadMap[whichLamp].B, Flash: launchPadMap[whichLamp].Flash}

			// Take into account the brightness. Divide by 2 because launch pad is 1-127.
			Red := int(((float64(alight.Red) / 2) / 100) * (float64(alight.Brightness) / 2.55))
			Green := int(((float64(alight.Green) / 2) / 100) * (float64(alight.Brightness) / 2.55))
			Blue := int(((float64(alight.Blue) / 2) / 100) * (float64(alight.Brightness) / 2.55))

			// We're in standard turn the light on.
			if !alight.Flash {

				// If we have this color already don't write again.
				newColor := common.Color{R: Red, G: Green, B: Blue}

				if storedColor != newColor || storedColor.Flash {
					// Now light the launchpad button.
					if debug {
						fmt.Printf("X:%d Y:%d Stored Color is %+v  New Color is %+v\n", whichLamp.X, whichLamp.Y, storedColor, newColor)
					}
					err := pad.Light(alight.X, alight.Y, Red, Green, Blue)
					if err != nil {
						fmt.Printf("error writing to launchpad %s\n" + err.Error())
					}
				}

			} else {
				// Now we're been asked go flash this button.
				if debug {
					fmt.Printf("Want Color %+v LaunchPad On Code is %x\n", alight.OnColor, common.GetLaunchPadColorCodeByRGB(alight.OnColor))
					fmt.Printf("Want Color %+v LaunchPad Off Code is %x\n", alight.OffColor, common.GetLaunchPadColorCodeByRGB(alight.OffColor))
				}
				err := pad.FlashLight(alight.X, alight.Y, int(common.GetLaunchPadColorCodeByRGB(alight.OnColor)), int(common.GetLaunchPadColorCodeByRGB(alight.OffColor)))
				if err != nil {
					fmt.Printf("flash: error writing to launchpad %s\n" + err.Error())
				}

			}
			// Remember what lamps are light.
			launchPadMap[coordinate{X: alight.X, Y: alight.Y}] = common.Color{
				R:     Red,
				G:     Green,
				B:     Blue,
				Flash: true,
			}
		}
	}
}
