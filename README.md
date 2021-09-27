# dmxlights
A DJ lighting package

A DMX lighting controller that uses a Novation Launchpad as a control surface and a cheap FTDI interface for 
connecting to the lighting fixtures.

Novation Launchpad Mini Mk3 - https://novationmusic.com/en/launch/launchpad-mini
FTDI interface card is USB to RS485 TTL Serial Converter Adapter FTDI interface FT232RL 75176 Module_AM - https://www.ebay.co.uk/itm/333947918351

# Dependances
- github.com/rakyll/launchpad/mk2
- github.com/oliread/usbdmx"
- github.com/oliread/usbdmx/ft232"


# FTDI Integration





## Notes on getting the FTDI interface working.

If you see the following error:
2021/08/08 17:29:33 Failed to connect DMX Controller: failed to select interface #0 alternate setting 0 of config 1 of device vid=0403,pid=6001,bus=38,addr=1: failed to claim interface 0 on vid=0403,pid=6001,bus=38,addr=1,config=1: libusb: bad access [code -3]
Dereks-iMac:dmxlights derek$ 

Stop the USB driver grabbing the device.

Looks like other drivers could be claiming the channel.
$ kextstat | grep -i ftdi
  161    0 0xffffff7f82d5a000 0x7000     0x7000     com.FTDI.driver.FTDIUSBSerialDriver (2.3) ECC3AF36-431D-370D-86F2-5237785E9CF8 <107 49 5 4 3 1>

Remove one or both
$ sudo kextunload -b com.FTDI.driver.FTDIUSBSerialDriver
$ sudo kextunload -b com.apple.driver.AppleUSBFTDI

##  Launch Pad Integration

Use mk2  open and change the following three functions in 
I need to contribute the changes back to create a mk3 version.

"github.com/rakyll/launchpad/mk2”

```
func (l *Launchpad) Program() error {
    // programmers mode
    return l.outputStream.WriteSysExBytes(portmidi.Time(), []byte{0xF0, 0x00, 0x20, 0x29, 0x02, 0x0D, 0x00, 0x7f, 0xF7})
}
```

```
// discovers the currently connected Launchpad device as a MIDI device.
func discover() (input portmidi.DeviceID, output portmidi.DeviceID, err error) {
    in := -1
    out := -1
    for i := 0; i < portmidi.CountDevices(); i++ {
        info := portmidi.Info(portmidi.DeviceID(i))
        if strings.Contains(info.Name, "Launchpad") {
            if info.IsInputAvailable {
                in = i
            }
            if info.IsOutputAvailable {
                out = i
            }
    }
}

if in == -1 || out == -1 {
    err = errors.New("launchpad: no launchpad is connected")
} else {
    input = portmidi.DeviceID(in)
    output = portmidi.DeviceID(out)
}
    return
}
```

```
// Read reads hits from the input stream. It returns max 64 hits for each read.
func (l *Launchpad) Read() (hits []Hit, err error) {
	var evts []portmidi.Event
	if evts, err = l.inputStream.Read(1024); err != nil {
		return
	}
	for _, evt := range evts {
		  if evt.Data2 > 0 {
			  var x, y int64
			  if evt.Status == 176 {
				  // top row button
				  // FIXME
				  // x = evt.Data1 - 104
			   	// y = -1
			  	x = evt.Data1%10 - 1
			  	y = (8 - (evt.Data1-x)/10)
		  	} else {
				    x = evt.Data1%10 - 1
				    y = (8 - (evt.Data1-x)/10)
			  }
			  hits = append(hits, Hit{X: int(x), Y: int(y)})
		  } 
	}
	return
}
```
```
func (l *Launchpad) Light(x, y, color int) error {
    // TODO(jbd): Support top row.
    led := int64((8-y)*10 + x + 1)
    //return l.outputStream.WriteShort(0x90, led, int64(color))
    return l.outputStream.WriteSysExBytes(portmidi.Time(), []byte{0xF0, 0x00, 0x20, 0x29, 0x02, 0x0D, 0x03, 0x00, byte(led), byte(color), 0xF7})
}

func (l *Launchpad) FlashLight(x int, y int, onColor int, offColor int) error {
	// TODO(jbd): Support top row.
	led := int64((8-y)*10 + x + 1)
	return l.outputStream.WriteSysExBytes(portmidi.Time(), []byte{0xF0, 0x00, 0x20, 0x29, 0x02, 0x0D, 0x03, 0x01, byte(led), byte(onColor), byte(offColor), 0xF7})
}

// Light lights the button at x,y with the given red, green, and blue values.
// x and y are [0, 7]. Color is [0, 128).
// All available colors are documented and visualized at Launchpad's Programmers Guide
// at https://global.novationmusic.com/sites/default/files/novation/downloads/10529/launchpad-mk2-programmers-reference-guide_0.pdf.
func (l *Launchpad) Light(x int, y int, red int, green int, blue int) error {
	  // TODO(jbd): Support top row.
	  led := int64((8-y)*10 + x + 1)
	  //return l.outputStream.WriteShort(0x90, led, int64(color))
	  return l.outputStream.WriteSysExBytes(portmidi.Time(), []byte{0xF0, 0x00, 0x20, 0x29, 0x02, 0x0D, 0x03, 0x03, byte(led), byte(red), byte(green), byte(blue), 0xF7})
}
