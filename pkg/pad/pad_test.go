package pad

import (
	"bytes"
	"errors"
	"testing"
	"time"

	"github.com/dhowlett99/dmxlights/pkg/common"
	"github.com/scgolang/midi"
)

// mockTransport records every Write() chunk and lets tests push incoming MIDI packets.
type mockTransport struct {
	chunks   [][]byte
	packets  chan []midi.Packet
	writeErr error
	closed   bool
}

func newMockTransport() *mockTransport {
	return &mockTransport{
		packets: make(chan []midi.Packet, 8),
	}
}

func (m *mockTransport) Open() error { return nil }

func (m *mockTransport) Close() error {
	m.closed = true
	return nil
}

func (m *mockTransport) Write(p []byte) (int, error) {
	if m.writeErr != nil {
		return 0, m.writeErr
	}
	cp := append([]byte(nil), p...)
	m.chunks = append(m.chunks, cp)
	return len(p), nil
}

func (m *mockTransport) Packets() (<-chan []midi.Packet, error) {
	return m.packets, nil
}

func (m *mockTransport) fullStream() []byte {
	var out []byte
	for _, c := range m.chunks {
		out = append(out, c...)
	}
	return out
}

func newTestPad(t midiTransport) *Pad {

	return &Pad{
		transport:   t,
		online:      true,
		checkOnline: func() bool { return true },
		closed:      make(chan struct{}),
	}
}

func TestProgramAndReset(t *testing.T) {
	tests := []struct {
		name string
		call func(*Pad) error
		want []byte
	}{
		{
			name: "Program enters Programmer mode",
			call: (*Pad).Program,
			want: []byte{0xF0, 0x00, 0x20, 0x29, 0x02, 0x0D, 0x0E, 0x01, 0xF7},
		},
		{
			name: "Reset exits Programmer mode",
			call: (*Pad).Reset,
			want: []byte{0xF0, 0x00, 0x20, 0x29, 0x02, 0x0D, 0x0E, 0x00, 0xF7},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mock := newMockTransport()
			p := newTestPad(mock)

			if err := tc.call(p); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got := mock.fullStream(); !bytes.Equal(got, tc.want) {
				t.Errorf("got SysEx %X, want %X", got, tc.want)
			}
			for i, chunk := range mock.chunks {
				if len(chunk) > 3 {
					t.Errorf("chunk %d exceeded 3 bytes: len=%d (%X)", i, len(chunk), chunk)
				}
			}
		})
	}
}

func TestLight(t *testing.T) {
	tests := []struct {
		name      string
		x, y      int
		r, g, b   uint8
		want      []byte
		wantError bool
	}{
		{
			name: "bottom-left grid pad (0,7) maps to LED 11",
			x:    0, y: 7,
			r: 0x7F, g: 0x40, b: 0x10,
			want: []byte{0xF0, 0x00, 0x20, 0x29, 0x02, 0x0D, 0x03, 0x03, 11, 0x7F, 0x40, 0x10, 0xF7},
		},
		{
			name: "top-right grid pad (7,0) maps to LED 88 and masks 8-bit RGB to 7-bit",
			x:    7, y: 0,
			r: 0xFF, g: 0x80, b: 0x01,
			want: []byte{0xF0, 0x00, 0x20, 0x29, 0x02, 0x0D, 0x03, 0x03, 88, 0x7F, 0x00, 0x01, 0xF7},
		},
		{
			name: "out-of-bounds negative X returns error",
			x:    -1, y: 0,
			r: 10, g: 10, b: 10,
			wantError: true,
		},
		{
			name: "out-of-bounds Y > 8 returns error",
			x:    0, y: 9,
			r: 10, g: 10, b: 10,
			wantError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mock := newMockTransport()
			p := newTestPad(mock)

			err := p.Light(tc.x, tc.y, tc.r, tc.g, tc.b)
			if tc.wantError {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got := mock.fullStream(); !bytes.Equal(got, tc.want) {
				t.Errorf("got SysEx %X, want %X", got, tc.want)
			}
		})
	}
}

func TestFlashLight(t *testing.T) {
	mock := newMockTransport()
	p := newTestPad(mock)

	// Valid flash command at (2, 4) -> LED = (8-4)*10 + 2 + 1 = 43
	if err := p.FlashLight(2, 4, 0x05, 0x00); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := []byte{0xF0, 0x00, 0x20, 0x29, 0x02, 0x0D, 0x03, 0x01, 43, 0x05, 0x00, 0xF7}
	if got := mock.fullStream(); !bytes.Equal(got, want) {
		t.Errorf("got SysEx %X, want %X", got, want)
	}

	// Color > 0x7F must fail validation
	if err := p.FlashLight(2, 4, 0x80, 0x00); err == nil {
		t.Error("expected error for 8-bit onColor > 0x7F, got nil")
	}
}

func TestLightBatch(t *testing.T) {
	mock := newMockTransport()
	p := newTestPad(mock)

	// Solid LED 11 color 5 + Flashing LED 12 color 9->0
	err := p.LightBatch(
		0x00, 11, 0x05,
		0x01, 12, 0x09, 0x00,
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := []byte{
		0xF0, 0x00, 0x20, 0x29, 0x02, 0x0D,
		0x03,
		0x00, 11, 0x05,
		0x01, 12, 0x09, 0x00,
		0xF7,
	}
	if got := mock.fullStream(); !bytes.Equal(got, want) {
		t.Errorf("got SysEx %X, want %X", got, want)
	}
}

func TestWriteSysExChunkedPropagatesError(t *testing.T) {
	mock := newMockTransport()
	mock.writeErr = errors.New("usb disconnect")
	p := newTestPad(mock)

	if err := p.Program(); err == nil {
		t.Fatal("expected write error to propagate, got nil")
	}
}

func TestListen(t *testing.T) {
	mock := newMockTransport()
	p := newTestPad(mock)

	buttons := make(chan common.Button, 4)
	done := make(chan error, 1)

	go func() {
		done <- p.Listen(buttons)
	}()

	// LED 43 corresponds to x=2, y=4 ((8-4)*10 + 2 + 1 = 43)
	mock.packets <- []midi.Packet{
		{Data: [3]byte{0x90, 43, 127}}, // Press
		{Data: [3]byte{0x90, 43, 0}},   // Release (X offset by +100)
	}

	readButton := func() common.Button {
		t.Helper()
		select {
		case btn := <-buttons:
			return btn
		case <-time.After(time.Second):
			t.Fatal("timed out waiting for button event")
			return common.Button{}
		}
	}

	press := readButton()
	if press.X != 2 || press.Y != 4 {
		t.Errorf("press event = (%d, %d), want (2, 4)", press.X, press.Y)
	}

	release := readButton()
	if release.X != 102 || release.Y != 4 {
		t.Errorf("release event = (%d, %d), want (102, 4)", release.X, release.Y)
	}

	if err := p.Close(); err != nil {
		t.Fatalf("Close returned unexpected error: %v", err)
	}

	select {
	case err := <-done:
		if err != nil {
			t.Fatalf("Listen returned unexpected error: %v", err)
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for Listen to exit")
	}
}

func TestAutoReconnect(t *testing.T) {
	first := newMockTransport()
	second := newMockTransport()

	var attempts int
	dial := func() (midiTransport, error) {
		attempts++
		switch attempts {
		case 1:
			return first, nil
		case 2:
			return nil, ErrNotConnected // Simulate device still unplugged for one retry tick
		default:
			return second, nil
		}
	}

	p := newPad(dial, 10*time.Millisecond)
	reconnected := make(chan struct{}, 2)
	p.OnReconnect = func() {
		reconnected <- struct{}{}
	}

	if err := p.connect(); err != nil {
		p.online = true
		t.Fatalf("initial connect failed: %v", err)
	}
	<-reconnected // Drain initial connect notification

	buttons := make(chan common.Button, 4)
	go p.Listen(buttons)
	defer p.Close()

	// Simulate USB unplug by closing the first packet channel
	close(first.packets)

	select {
	case <-reconnected:
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for auto-reconnect")
	}

	// Verify second transport received Programmer mode SysEx on reconnect
	wantProgram := []byte{0xF0, 0x00, 0x20, 0x29, 0x02, 0x0D, 0x0E, 0x01, 0xF7}
	if got := second.fullStream(); !bytes.Equal(got, wantProgram) {
		t.Errorf("second transport SysEx = %X, want %X", got, wantProgram)
	}

	// Verify button presses flow again on the reconnected transport
	second.packets <- []midi.Packet{{Data: [3]byte{0x90, 43, 127}}}
	select {
	case btn := <-buttons:
		if btn.X != 2 || btn.Y != 4 {
			t.Errorf("got button (%d, %d), want (2, 4)", btn.X, btn.Y)
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for button event after reconnect")
	}
}
