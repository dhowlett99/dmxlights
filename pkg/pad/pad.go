// Simple Midi Interface for Novation Launchpad inspired by work done by https://github.com/scgolang/launchpad
// and https://github.com/rakyll/launchpad.

package pad

import (
	"fmt"
	"strings"

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/pkg/errors"
	"github.com/scgolang/midi"
)

// midiTransport abstracts *midi.Device so Pad can be unit-tested without physical hardware.
type midiTransport interface {
	Open() error
	Close() error
	Write([]byte) (int, error)
	Packets() (<-chan []midi.Packet, error)
}

// Pad represents a connection to a Novation Launchpad Mini MK3.
type Pad struct {
	midiTransport
}

// Open opens a connection to the Novation Launchpad Mini MK3.
func Open() (*Pad, error) {
	devices, err := midi.Devices()
	if err != nil {
		return nil, errors.Wrap(err, "listing MIDI devices")
	}

	var device *midi.Device
	for _, d := range devices {
		if (strings.Contains(d.Name, "Launchpad") || strings.Contains(d.Name, "LPMiniMK3")) &&
			strings.Contains(d.Name, "MIDI") {
			device = d
			break
		}
	}
	if device == nil {
		return nil, errors.New("Launchpad MIDI device not found")
	}

	p := &Pad{midiTransport: device}
	if err := p.Open(); err != nil {
		return nil, errors.Wrap(err, "opening Launchpad MIDI device")
	}
	return p, nil
}

// Close closes the connection to the Pad.
func (pad *Pad) Close() error {
	return errors.Wrap(pad.midiTransport.Close(), "closing midi device")
}

// Program puts the Launchpad Mini MK3 into Programmer mode.
func (pad *Pad) Program() error {
	return pad.sendSysEx(0x0E, 0x01)
}

// Reset exits Programmer mode and returns the Launchpad Mini MK3 to Live mode.
func (pad *Pad) Reset() error {
	return pad.sendSysEx(0x0E, 0x00)
}

// Light lights the button at x,y with the given 7-bit red, green, and blue values (0-127).
func (pad *Pad) Light(x, y int, red, green, blue uint8) error {
	led, err := padIndex(x, y)
	if err != nil {
		return err
	}
	return pad.sendSysEx(0x03, 0x03, led, red&0x7F, green&0x7F, blue&0x7F)
}

// FlashLight flashes a specific pad LED between onColor and offColor (palette indices 0-127).
func (pad *Pad) FlashLight(x, y int, onColor, offColor byte) error {
	led, err := padIndex(x, y)
	if err != nil {
		return err
	}
	if onColor > 0x7F || offColor > 0x7F {
		return fmt.Errorf("color values must be 7-bit MIDI data bytes (0x00-0x7F)")
	}
	return pad.sendSysEx(0x03, 0x01, led, onColor, offColor)
}

// LightBatch sends multiple solid, flashing, or RGB LED updates in a single SysEx 03h message.
// Each entry in specs must follow Novation's 03h format (e.g., [0x00, led, color] or [0x03, led, r, g, b]).
func (pad *Pad) LightBatch(specs ...byte) error {
	payload := append([]byte{0x03}, specs...)
	return pad.sendSysEx(payload...)
}

// sendSysEx wraps payload bytes in the Mini MK3 SysEx header/footer and writes in 3-byte USB-MIDI chunks.
func (pad *Pad) sendSysEx(payload ...byte) error {
	msg := make([]byte, 0, len(payload)+7)
	msg = append(msg, 0xF0, 0x00, 0x20, 0x29, 0x02, 0x0D) // Mini MK3 header
	msg = append(msg, payload...)
	msg = append(msg, 0xF7) // End of SysEx

	return pad.WriteSysExChunked(msg)
}

// WriteSysExChunked writes a complete SysEx byte slice 3 bytes at a time for USB-MIDI packet drivers.
func (pad *Pad) WriteSysExChunked(sysex []byte) error {
	for i := 0; i < len(sysex); i += 3 {
		end := i + 3
		if end > len(sysex) {
			end = len(sysex)
		}
		if _, err := pad.Write(sysex[i:end]); err != nil {
			return err
		}
	}
	return nil
}

// Listen reads button press and release events from the Launchpad and forwards them to buttonChannel.
func (pad *Pad) Listen(buttonChannel chan<- common.Button) error {
	eventChannel, err := pad.Packets()
	if err != nil {
		return errors.Wrap(err, "opening MIDI packet channel")
	}

	for events := range eventChannel {
		for _, packet := range events {
			if packet.Err != nil {
				fmt.Printf("packet error: %v\n", packet.Err)
				continue
			}
			if len(packet.Data) < 3 {
				continue
			}

			note := int(packet.Data[1])
			velocity := packet.Data[2]

			x := note%10 - 1
			y := 8 - (note-(x+1))/10

			if velocity == 0 {
				// Button released
				x += 100
			}
			buttonChannel <- common.Button{X: x, Y: y}
		}
	}
	return nil
}

func padIndex(x, y int) (byte, error) {
	if x < 0 || x > 8 || y < 0 || y > 8 {
		return 0, fmt.Errorf("coordinates (%d, %d) out of range (0-8)", x, y)
	}
	return byte((8-y)*10 + x + 1), nil
}
