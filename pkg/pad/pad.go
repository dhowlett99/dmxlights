// Simple Midi Interface for Novation Launchpad.
package pad

import (
	"fmt"
	"log"
	"strings"

	"github.com/pkg/errors"
	"github.com/scgolang/midi"
)

type Pad struct {
	*midi.Device
	hits chan Hit
}
type Hit struct {
	X int
	Y int
}

func Open() (*Pad, error) {
	devices, err := midi.Devices()
	if err != nil {
		return nil, errors.Wrap(err, "listing MIDI devices")
	}
	var device *midi.Device
	for _, d := range devices {
		if strings.Contains(d.Name, "MIDI") {
			device = d
			break
		}
	}
	if device == nil {
		return nil, errors.New("Pad not found")
	}
	pad := &Pad{Device: device}
	if err := pad.Open(); err != nil {
		return nil, err
	}
	return pad, nil
}

// Close closes the connection to the Pad.
func (pad *Pad) Close() error {
	if pad.hits != nil {
		close(pad.hits)
	}
	return errors.Wrap(pad.Device.Close(), "closing midi device")
}

func (pad *Pad) Reset() error {
	_, err := pad.Write([]byte{0xf0, 0x00, 0x20})
	if err != nil {
		return err
	}
	_, err = pad.Write([]byte{0x29, 0x02, 0x18})
	if err != nil {
		return err
	}
	_, err = pad.Write([]byte{0x0e, 0x00, 0xf7})
	if err != nil {
		return err
	}
	return err
}

func (pad *Pad) Program() error {
	_, err := pad.Write([]byte{0xf0, 0x00, 0x20})
	if err != nil {
		return err
	}
	_, err = pad.Write([]byte{0x29, 0x02, 0x0d})
	if err != nil {
		return err
	}
	_, err = pad.Write([]byte{0x0e, 0x01, 0xf7})
	if err != nil {
		return err
	}
	return nil
}

func (pad *Pad) Listen(buttonchannel chan Hit) error {
	eventChannel, err := pad.Packets()
	if err != nil {
		log.Fatal("error can't open button channel")
		return err
	}

	for {

		events := <-eventChannel

		for _, packet := range events {
			if packet.Err != nil {
				fmt.Printf("packet error")
				continue
			}
			var x, y int

			// Button pressed codes.
			if packet.Data[2] > 0 {
				x = int(packet.Data[1])%10 - 1
				y = 8 - (int(packet.Data[1])-x)/10
				buttonchannel <- Hit{X: x, Y: y}
			}

			// Button released codes.
			if packet.Data[2] == 0 {
				x = int(packet.Data[1])%10 - 1
				y = 8 - (int(packet.Data[1])-x)/10
				x = x + 100
				buttonchannel <- Hit{X: x, Y: y}
			}
		}
	}
}

// Light lights the button at x,y with the given red, green, and blue values.
func (pad *Pad) Light(x, y, red int, green int, blue int) error {
	led := int64((8-y)*10 + x + 1)

	_, err := pad.Write([]byte{0xF0, 0x00, 0x20})
	if err != nil {
		return err
	}

	_, err = pad.Write([]byte{0x29, 0x02, 0x0D})
	if err != nil {
		return err
	}

	_, err = pad.Write([]byte{0x03, 0x03, byte(led)})
	if err != nil {
		return err
	}

	_, err = pad.Write([]byte{byte(red), byte(green), byte(blue)})
	if err != nil {
		return err
	}

	_, err = pad.Write([]byte{0xF7, 0, 0})
	if err != nil {
		return err
	}
	return nil
}
