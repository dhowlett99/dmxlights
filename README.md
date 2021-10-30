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
```
2021/08/08 17:29:33 Failed to connect DMX Controller: failed to select interface #0 alternate setting 0 of config 1 of device vid=0403,pid=6001,bus=38,addr=1: failed to claim interface 0 on vid=0403,pid=6001,bus=38,addr=1,config=1: libusb: bad access [code -3]
```

Stop the USB driver grabbing the device.

Looks like other drivers could be claiming the channel.
```
$ kextstat | grep -i ftdi
  161    0 0xffffff7f82d5a000 0x7000     0x7000     com.FTDI.driver.FTDIUSBSerialDriver (2.3) ECC3AF36-431D-370D-86F2-5237785E9CF8 <107 49 5 4 3 1>
```
Remove one or both
```
$ sudo kextunload -b com.FTDI.driver.FTDIUSBSerialDriver
$ sudo kextunload -b com.apple.driver.AppleUSBFTDI
```
##  Launch Pad Integration

I use the mk3 version.

"github.com/rakyll/launchpad/mk3”

