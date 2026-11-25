// Simple Midi Interface for Novation Launchpad inspired by work done by https://github.com/scgolang/launchpad
// and https://github.com/rakyll/launchpad.

package pad

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/pkg/errors"
	"github.com/scgolang/midi"
)

const debug = false

// ErrNotConnected is returned by LED write methods while the Launchpad is unplugged.
var ErrNotConnected = errors.New("launchpad is not connected")

// midiTransport abstracts *midi.Device so Pad can be unit-tested without physical hardware.
type midiTransport interface {
	Open() error
	Close() error
	Write([]byte) (int, error)
	Packets() (<-chan []midi.Packet, error)
}

// Pad represents an auto-reconnecting connection to a Novation Launchpad Mini MK3.
type Pad struct {
	mu          sync.RWMutex
	transport   midiTransport
	online      bool
	dial        func() (midiTransport, error)
	checkOnline func() bool
	retryDelay  time.Duration
	closed      chan struct{}
	closeOnce   sync.Once
	OnReconnect func() // Optional hook to redraw LED state after a reconnect
}

// Open initializes a Pad and attempts an immediate connection to the Launchpad Mini MK3.
// If the device is not plugged in yet, Listen will continue polling and connect automatically.
func Open() (*Pad, error) {
	p := newPad(discoverLaunchpad, time.Second)
	if err := p.connect(); err != nil {
		if errors.Is(err, ErrNotConnected) {
			fmt.Println("No Launchpad found at startup; waiting for USB connection...")
			return p, nil
		}
		return nil, err
	}
	return p, nil
}

func newPad(dial func() (midiTransport, error), retryDelay time.Duration) *Pad {
	return &Pad{
		dial:        dial,
		checkOnline: checkHardwareOnline,
		retryDelay:  retryDelay,
		closed:      make(chan struct{}),
	}
}

func discoverLaunchpad() (midiTransport, error) {
	hwOnline := checkHardwareOnline()
	devices, err := midi.Devices()
	if err != nil {
		if debug {
			fmt.Printf("Launchpad: midi.Devices() error: %v\n", err)
		}
		return nil, errors.Wrap(err, "listing MIDI devices")
	}

	var names []string
	for _, d := range devices {
		names = append(names, d.Name)
	}
	if debug {
		fmt.Printf("Launchpad: midi.Devices() = %v (IOKit online=%v)\n", names, hwOnline)
	}

	for _, d := range devices {
		// On Launchpad Mini MK3, Port 1 is DAW and Port 2 is the MIDI / Programmer port.
		isMidiPort := strings.Contains(d.Name, "MIDI") || (hwOnline && d.Name == "Port 2")
		if isMidiPort {
			if err := d.Open(); err != nil {
				if debug {
					fmt.Printf("Launchpad: failed to open %q: %v\n", d.Name, err)
				}
				return nil, errors.Wrapf(err, "opening MIDI device %q", d.Name)
			}
			if debug {
				fmt.Printf("Launchpad: successfully opened %q\n", d.Name)
			}
			return d, nil
		}
	}

	return nil, ErrNotConnected
}

// IsConnected reports whether the Launchpad currently has an active, online MIDI connection.
func (pad *Pad) IsConnected() bool {
	pad.mu.RLock()
	defer pad.mu.RUnlock()
	return pad.transport != nil && pad.online
}

func (pad *Pad) connect() error {
	pad.mu.RLock()
	t := pad.transport
	pad.mu.RUnlock()

	if t == nil {
		var err error
		t, err = pad.dial()
		if err != nil {
			return err
		}
		pad.mu.Lock()
		pad.transport = t
		pad.mu.Unlock()
	}

	pad.mu.Lock()
	pad.online = true
	pad.mu.Unlock()

	if err := pad.Program(); err != nil {
		pad.mu.Lock()
		pad.online = false
		pad.mu.Unlock()
		if debug {
			fmt.Printf("Launchpad: Program() failed: %v\n", err)
		}
		return errors.Wrap(err, "entering programmer mode")
	}

	fmt.Println("Launchpad: connected and set to Programmer mode")

	if pad.OnReconnect != nil {
		if debug {
			fmt.Println("Launchpad: firing OnReconnect callback")
		}
		pad.OnReconnect()
	} else {
		if debug {
			fmt.Println("Launchpad: OnReconnect is nil")
		}
	}
	return nil
}

// Close permanently shuts down the Pad and stops any active reconnection loop.
func (pad *Pad) Close() error {
	var closeErr error
	pad.closeOnce.Do(func() {
		close(pad.closed)
		pad.mu.Lock()
		t := pad.transport
		pad.transport = nil
		pad.online = false
		pad.mu.Unlock()

		if t != nil {
			closeErr = errors.Wrap(t.Close(), "closing midi device")
		}
	})
	return closeErr
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
	msg = append(msg, 0xF0, 0x00, 0x20, 0x29, 0x02, 0x0D)
	msg = append(msg, payload...)
	msg = append(msg, 0xF7)

	return pad.WriteSysExChunked(msg)
}

// WriteSysExChunked writes a complete SysEx byte slice 3 bytes at a time for USB-MIDI packet drivers.
func (pad *Pad) WriteSysExChunked(sysex []byte) error {
	pad.mu.RLock()
	t := pad.transport
	online := pad.online
	pad.mu.RUnlock()

	if t == nil || !online {
		return ErrNotConnected
	}

	for i := 0; i < len(sysex); i += 3 {
		end := i + 3
		if end > len(sysex) {
			end = len(sysex)
		}
		if _, err := t.Write(sysex[i:end]); err != nil {
			pad.mu.Lock()
			pad.online = false
			pad.mu.Unlock()
			return errors.Wrap(err, "writing SysEx chunk")
		}
	}
	return nil
}

// Listen continuously reads button events and restores Programmer mode + LEDs whenever the USB cable is replugged.
func (pad *Pad) Listen(buttonChannel chan<- common.Button) error {
	var eventChannel <-chan []midi.Packet

	for {
		select {
		case <-pad.closed:
			return nil
		default:
		}

		// Ensure initial transport is opened if the app started while unplugged.
		pad.mu.RLock()
		t := pad.transport
		pad.mu.RUnlock()

		if t == nil {
			if err := pad.connect(); err != nil {
				select {
				case <-pad.closed:
					return nil
				case <-time.After(pad.retryDelay):
					continue
				}
			}
			pad.mu.RLock()
			t = pad.transport
			pad.mu.RUnlock()
		}

		if eventChannel == nil {
			var err error
			eventChannel, err = t.Packets()
			if err != nil {
				select {
				case <-pad.closed:
					return nil
				case <-time.After(pad.retryDelay):
					continue
				}
			}
			if debug {
				fmt.Println("Launchpad DEBUG: Packets() stream started")
			}
		}

		pad.readLoop(eventChannel, buttonChannel)
		// If readLoop returned because eventChannel closed, clear it so Packets() is re-acquired.
		eventChannel = nil
	}
}

// func (pad *Pad) readLoop(eventChannel <-chan []midi.Packet, buttonChannel chan<- common.Button) {
// 	ticker := time.NewTicker(500 * time.Millisecond)
// 	defer ticker.Stop()

// 	// Only trigger USB disconnect detection if IOKit has observed the device online at least once,
// 	// so IOKit can never falsely block a working MIDI connection.
// 	armed := pad.checkOnline != nil && pad.checkOnline()

// 	for {
// 		select {
// 		case <-pad.closed:
// 			return

// 		case <-ticker.C:
// 			if pad.checkOnline == nil {
// 				continue
// 			}
// 			hwOnline := pad.checkOnline()
// 			if hwOnline {
// 				armed = true
// 			}
// 			if !armed {
// 				continue
// 			}

// 			wasOnline := pad.IsConnected()
// 			if wasOnline && !hwOnline {
// 				fmt.Println("Launchpad: disconnected (USB unplugged)")
// 				pad.mu.Lock()
// 				pad.online = false
// 				pad.mu.Unlock()
// 			} else if !wasOnline && hwOnline {
// 				// Give macOS AppleUSBMIDIDriver a brief moment to re-attach the endpoint
// 				time.Sleep(300 * time.Millisecond)
// 				if err := pad.connect(); err != nil {
// 					fmt.Printf("Launchpad: reconnect error: %v\n", err)
// 				}
// 			}

// 		case events, ok := <-eventChannel:
// 			if !ok {
// 				if debug {
// 					fmt.Println("Launchpad: MIDI packet channel closed")
// 				}
// 				pad.mu.Lock()
// 				pad.online = false
// 				pad.mu.Unlock()
// 				return
// 			}
// 			for _, packet := range events {
// 				if packet.Err != nil {
// 					fmt.Printf("Launchpad: packet error (%v)\n", packet.Err)
// 					continue
// 				}
// 				if len(packet.Data) < 3 {
// 					continue
// 				}

// 				const (
// 					noteIdx = 1
// 					velIdx  = 2
// 				)

// 				note := int(packet.Data[noteIdx])
// 				velocity := packet.Data[velIdx]

// 				x := note%10 - 1
// 				y := 8 - (note-(x+1))/10

// 				if velocity == 0 {
// 					x += 100
// 				}

// 				select {
// 				case buttonChannel <- common.Button{X: x, Y: y}:
// 				case <-pad.closed:
// 					return
// 				}
// 			}
// 		}
// 	}
// }

func (pad *Pad) readLoop(eventChannel <-chan []midi.Packet, buttonChannel chan<- common.Button) {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	// Only trigger USB disconnect detection if IOKit has observed the device online at least once,
	// so IOKit can never falsely block a working MIDI connection.
	armed := pad.checkOnline != nil && pad.checkOnline()

	for {
		select {
		case <-pad.closed:
			return

		case <-ticker.C:
			if pad.checkOnline == nil {
				continue
			}
			hwOnline := pad.checkOnline()
			if hwOnline {
				armed = true
			}
			if !armed {
				continue
			}

			wasOnline := pad.IsConnected()
			if wasOnline && !hwOnline {
				fmt.Println("Launchpad: disconnected (USB unplugged)")
				pad.mu.Lock()
				pad.online = false
				pad.mu.Unlock()
			} else if !wasOnline && hwOnline {
				// Give macOS AppleUSBMIDIDriver a brief moment to re-attach the endpoint
				time.Sleep(300 * time.Millisecond)
				if err := pad.connect(); err != nil {
					fmt.Printf("Launchpad: reconnect error: %v\n", err)
				}
			}

		case events, ok := <-eventChannel:
			if !ok {
				if debug {
					fmt.Println("Launchpad: MIDI packet channel closed")
				}
				pad.mu.Lock()
				old := pad.transport
				pad.transport = nil
				pad.online = false
				pad.mu.Unlock()
				if old != nil {
					_ = old.Close()
				}
				return
			}
			for _, packet := range events {
				if packet.Err != nil {
					fmt.Printf("Launchpad: packet error (%v)\n", packet.Err)
					continue
				}
				if len(packet.Data) < 3 {
					continue
				}

				const (
					noteIdx = 1
					velIdx  = 2
				)

				note := int(packet.Data[noteIdx])
				velocity := packet.Data[velIdx]

				x := note%10 - 1
				y := 8 - (note-(x+1))/10

				if velocity == 0 {
					x += 100
				}

				select {
				case buttonChannel <- common.Button{X: x, Y: y}:
				case <-pad.closed:
					return
				}
			}
		}
	}
}

func padIndex(x, y int) (byte, error) {
	if x < 0 || x > 8 || y < 0 || y > 8 {
		return 0, fmt.Errorf("coordinates (%d, %d) out of range (0-8)", x, y)
	}
	return byte((8-y)*10 + x + 1), nil
}
