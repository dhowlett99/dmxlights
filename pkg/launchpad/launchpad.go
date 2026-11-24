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
	"github.com/dhowlett99/dmxlights/pkg/dmx"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
	"github.com/dhowlett99/dmxlights/pkg/pad"
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
	eventsForLaunchpad chan common.ALight, dmxController *dmx.ControllerConfig,
	fixturesConfig *fixture.Fixtures, commandChannels []chan common.Command,
	replyChannels []chan common.Sequence, updateChannels []chan common.Sequence,
	dmxInterfaceCardPresent bool) {

	// Create a channel to listen for buttons being pressed.
	// Send the button pressed hit to the button channel.
	buttonChannel := make(chan common.Button)
	go func() {
		this.Pad.Listen(buttonChannel)
	}()

	// Main loop reading commands from the Novation Launchpad.
	for {
		hit := <-buttonChannel
		this.GUI = false
		buttons.ProcessButtons(hit.X, hit.Y, sequences, this, eventsForLaunchpad, guiButtons, dmxController, fixturesConfig, commandChannels, replyChannels, updateChannels)
	}
}

type coordinate struct {
	X int
	Y int
}

func ListenAndSendToLaunchPad(eventsForLaunchpad chan common.ALight, pad *pad.Pad, launchPadConnected bool) {
	type PrivateColorMap struct {
		R     uint8
		G     uint8
		B     uint8
		Flash bool
	}
	go func() {
		launchPadMap := make(map[coordinate]PrivateColorMap, 81)
		pending := make(map[coordinate]common.ALight, 81)

		// Flush coalesced updates every 20ms (~50 FPS)
		ticker := time.NewTicker(20 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case alight, ok := <-eventsForLaunchpad:
				if !ok {
					return
				}
				if pad == nil || !pad.IsConnected() {
					continue
				}
				// Coalesce rapid updates to the same coordinate;
				// all 81 coordinates during startup/reset remain distinct.
				coord := coordinate{X: alight.Button.X, Y: alight.Button.Y}
				pending[coord] = alight

			case <-ticker.C:
				if pad == nil || !pad.IsConnected() {
					// Reset cache while unplugged so all LEDs repaint when reconnected
					if len(launchPadMap) > 0 {
						launchPadMap = make(map[coordinate]PrivateColorMap, 81)
					}
					if len(pending) > 0 {
						pending = make(map[coordinate]common.ALight, 81)
					}
					continue
				}

				if len(pending) == 0 {
					continue
				}

				// Build one combined SysEx 03h message for all dirty buttons
				sysex := make([]byte, 0, 8+len(pending)*5)
				sysex = append(sysex, 0xF0, 0x00, 0x20, 0x29, 0x02, 0x0D, 0x03)

				for whichLamp, alight := range pending {
					storedColor := launchPadMap[whichLamp]

					red := uint8(((float64(alight.Red) / 2) / 100) * (float64(alight.Brightness) / 2.55))
					green := uint8(((float64(alight.Green) / 2) / 100) * (float64(alight.Brightness) / 2.55))
					blue := uint8(((float64(alight.Blue) / 2) / 100) * (float64(alight.Brightness) / 2.55))

					newColor := PrivateColorMap{
						R:     red,
						G:     green,
						B:     blue,
						Flash: alight.Flash,
					}

					// Exact coordinate-to-LED formula from pad.Light
					led := byte((8-alight.Button.Y)*10 + alight.Button.X + 1)

					if !alight.Flash {
						if storedColor != newColor {
							if debug {
								fmt.Printf("X:%d Y:%d Stored Color is %+v  New Color is %+v\n", whichLamp.X, whichLamp.Y, storedColor, newColor)
							}
							// 0x03 = Static RGB LED spec (5 bytes)
							sysex = append(sysex, 0x03, led, red&0x7F, green&0x7F, blue&0x7F)
						}
					} else {
						onCode := common.GetLaunchPadCodeByRGBColor(alight.OnColor)
						offCode := common.GetLaunchPadCodeByRGBColor(alight.OffColor)
						if debug {
							fmt.Printf("Want Color %+v LaunchPad On Code is %x\n", alight.OnColor, onCode)
							fmt.Printf("Want Color %+v LaunchPad Off Code is %x\n", alight.OffColor, offCode)
						}
						// 0x01 = Flashing LED spec (4 bytes: offCode then onCode)
						sysex = append(sysex, 0x01, led, offCode&0x7F, onCode&0x7F)
					}

					launchPadMap[whichLamp] = newColor
					delete(pending, whichLamp)
				}

				// Only send if at least one LED actually changed
				if len(sysex) > 7 {
					sysex = append(sysex, 0xF7)
					if err := pad.WriteSysExChunked(sysex); err != nil {
						// Write failed (device just unplugged); reset cache so next reconnect repaints everything
						launchPadMap = make(map[coordinate]PrivateColorMap, 81)
					}
				}
			}
		}
	}()
}
