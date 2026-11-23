// Copyright (C) 2022,2023 dhowlett99.
// This is the dmxlights dmx interface (macOS native serial OpenDMX driver).

package dmx

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"syscall"
	"time"
	"unsafe"
)

// macOS Darwin serial ioctl constants
const (
	darwinTIOCSBRK    = 0x2000747b // Assert serial BREAK
	darwinTIOCCBRK    = 0x2000747a // Clear serial BREAK
	darwinIOSSIOSPEED = 0x80085402 // Set arbitrary baud rate (speed_t = uint64 on 64-bit macOS)
	dmxBaudRate       = 250000     // Standard DMX512 baud rate (250 kbps)
)

// ControllerConfig holds the active port and forwards channel updates to the controller.
type ControllerConfig struct {
	PortName   string
	controller *DMXController
}

// SetChannel updates a DMX channel (1-512) via dmxControllerConfig.SetChannel(address, value).
func (cfg *ControllerConfig) SetChannel(index int16, data byte) error {
	if cfg == nil || cfg.controller == nil {
		return errors.New("DMX controller config not initialized")
	}
	return cfg.controller.SetChannel(index, data)
}

// DMXController manages an FT232RL OpenDMX interface over macOS's native serial port.
type DMXController struct {
	mu       sync.Mutex
	file     *os.File
	portName string
	universe [513]byte // Byte 0 = DMX Start Code (0x00), Bytes 1..512 = Channels 1..512
}

func findFTDIPort() (string, error) {
	patterns := []string{
		"/dev/cu.usbserial*",
		"/dev/cu.usbmodem*",
		"/dev/cu.FT*",
	}
	for _, pattern := range patterns {
		matches, err := filepath.Glob(pattern)
		if err == nil && len(matches) > 0 {
			return matches[0], nil
		}
	}

	allPorts, _ := filepath.Glob("/dev/cu.*")
	return "", fmt.Errorf("no FTDI serial device found (available ports: %v)", allPorts)
}

// Connect opens the FTDI serial port and configures it for DMX512 (250000 baud, 8N2).
func (c *DMXController) Connect() error {
	port, err := findFTDIPort()
	if err != nil {
		return err
	}
	c.portName = port

	f, err := os.OpenFile(port, os.O_RDWR|syscall.O_NOCTTY|syscall.O_NONBLOCK, 0666)
	if err != nil {
		return fmt.Errorf("failed to open %s: %w", port, err)
	}

	fd := f.Fd()

	if err := syscall.SetNonblock(int(fd), false); err != nil {
		f.Close()
		return err
	}

	var termios syscall.Termios
	if _, _, errno := syscall.Syscall6(syscall.SYS_IOCTL, fd, syscall.TIOCGETA, uintptr(unsafe.Pointer(&termios)), 0, 0, 0); errno != 0 {
		f.Close()
		return fmt.Errorf("TIOCGETA failed: %v", errno)
	}

	// Raw mode: 8 data bits, no parity, 2 stop bits (8N2), no flow control
	termios.Iflag = 0
	termios.Oflag = 0
	termios.Lflag = 0
	termios.Cflag = syscall.CS8 | syscall.CSTOPB | syscall.CLOCAL | syscall.CREAD

	if _, _, errno := syscall.Syscall6(syscall.SYS_IOCTL, fd, syscall.TIOCSETA, uintptr(unsafe.Pointer(&termios)), 0, 0, 0); errno != 0 {
		f.Close()
		return fmt.Errorf("TIOCSETA failed: %v", errno)
	}

	// Set exact 250,000 baud using macOS IOSSIOSPEED ioctl
	speed := uint64(dmxBaudRate)
	if _, _, errno := syscall.Syscall(syscall.SYS_IOCTL, fd, darwinIOSSIOSPEED, uintptr(unsafe.Pointer(&speed))); errno != 0 {
		f.Close()
		return fmt.Errorf("failed to set 250000 baud: %v", errno)
	}

	c.file = f
	return nil
}

// SetChannel sets a single DMX channel (1-512) to a value (0-255).
func (c *DMXController) SetChannel(index int16, data byte) error {
	if index < 1 || index > 512 {
		return fmt.Errorf("channel index %d out of range (1-512)", index)
	}
	c.mu.Lock()
	c.universe[index] = data
	c.mu.Unlock()
	return nil
}

// Render transmits one complete DMX512 frame (BREAK -> MAB -> Start Code + 512 channels).
func (c *DMXController) Render() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.file == nil {
		return errors.New("DMX controller not connected")
	}

	fd := c.file.Fd()

	// 1. DMX BREAK (>= 88 microseconds low)
	if _, _, errno := syscall.Syscall(syscall.SYS_IOCTL, fd, darwinTIOCSBRK, 0); errno != 0 {
		return fmt.Errorf("TIOCSBRK failed: %v", errno)
	}
	time.Sleep(110 * time.Microsecond)

	// 2. Clear BREAK -> Mark After Break (MAB >= 8 microseconds high)
	if _, _, errno := syscall.Syscall(syscall.SYS_IOCTL, fd, darwinTIOCCBRK, 0); errno != 0 {
		return fmt.Errorf("TIOCCBRK failed: %v", errno)
	}
	time.Sleep(16 * time.Microsecond)

	// 3. Write Start Code (0x00) + all 512 channels
	_, err := c.file.Write(c.universe[:])
	return err
}

// Close closes the serial connection.
func (c *DMXController) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.file != nil {
		err := c.file.Close()
		c.file = nil
		return err
	}
	return nil
}

func NewDmXController() (*DMXController, *ControllerConfig, error) {
	controller := &DMXController{}

	if err := controller.Connect(); err != nil {
		return nil, nil, errors.New("failed to connect DMX Controller: " + err.Error())
	}

	config := &ControllerConfig{
		PortName:   controller.portName,
		controller: controller,
	}
	log.Printf("Connected to DMX interface on %s", controller.portName)

	// Continuous ~44Hz DMX512 refresh loop
	go func(c *DMXController) {
		for {
			if err := c.Render(); err != nil {
				log.Fatalf("Failed to render output: %s", err)
			}
			time.Sleep(23 * time.Millisecond)
		}
	}(controller)

	return controller, config, nil
}
