# dmxlights Version 2.0 Release Notes.

# Introduction 

DMX lights is a Open Source Project to control DMX light fixtures. It is designed for the small to medium mobile disco rig and be very easy to setup and use.

Firstly we use the Novation Launchpad Mini (Mk3) as the control surface and a cheap FTDI RS422 controller as the interface for the DMX lights.

Second design goal was to do away with having to have a large fixture library, with DMX lights you simply add the description of what channels do what to a simple text file.

As of versions 2.0 of dmxlights we now support a fixture editior. The editor is started by clicking on the DMX Lights logo in the top right.

This version of DMX Lights was developed on macOS 10.15.7 (19H2026)

## New in Version 2.0

* Completely new version of the sequencer, with fixtures kept in sync.
* Now supports 8 RGB patterns.
* Now support 5 Scanner patterns.
* Scanner patterns now correctly take size into account
* Support for manual Pan and Tilt.
* Selectable scanner colors and gobos
* Graphical user interface using https://fyne.io/
* Ability to configure fixtures using UI. Click on DMX Lights logo to bring up editor.
* Sound Trigger now has a 800Hz low pass filter and a automatic gain control.
* 24 color picker for static and chase colors.
* Enable & disable controls for each fixture.
* Added shutter chaser to scanner sequences.
* You can now name presets by labelling buttions in the GUI.
* Simplified launchpad interface for better performance.
* Flood, Strobe and Blackout fully supported
* Implemented booth buddy, a mini sequencer for single fixtures, allowing color changes and rotations to be controlled using a switch.

## Known Problems

* The Novation Launchpad can only take a certain amount of notes per second, its possible to crash the launchpad if you have a lot of activity. The workaround is that DMX lights will use the reset message from the launchpad, then reset the launchpad and refresh the current state of the launchpad lamps. You may not even notice that the crash as happened.

* When you have the scanner sequence selected, the top left clear button is used for tilting up the scanners. The clear button is activated by a long press instead of a short press. In the GUI the tilt up is not available and the clear is always a single short press.

* Presets can be deleted by a long press on the launchpad, but you cannot delete a preset from the GUI.

# Dependances.

* The fyne toolkit for creating graphical apps for desktop, mobile and web. https://fyne.io/
* USBDMX, a versatile USB DMX library written in Go for programatic show control and special effects. https://github.com/oliread/usbdmx
* A simple MIDI package for Go. Currently only supports Linux and Mac. https://github.com/scgolang/midi
* Interface to the PortAudio audio I/O library. https://github.com/gordonklaus/portaudio
* PortAudio is a free, cross-platform, open-source, audio I/O library. https://www.portaudio.com/

The Launchpad interface has been inspired by work done by https://github.com/scgolang/launchpad
and https://github.com/rakyll/launchpad.

# Shared Libraries Used

```
 otool -L ./dmxlights
```

* /System/Library/Frameworks/Foundation.framework/Versions/C/Foundation (compatibility version 300.0.0, current version 1677.104.0)

* /System/Library/Frameworks/UserNotifications.framework/Versions/A/UserNotifications (compatibility version 1.0.0, current version 1.0.0)

* /usr/lib/libobjc.A.dylib (compatibility version 1.0.0, current version 228.0.0)
       
* /usr/lib/libresolv.9.dylib (compatibility version 1.0.0, current version 1.0.0)
       
* /System/Library/Frameworks/AppKit.framework/Versions/C/AppKit (compatibility version 45.0.0, current version 1894.60.100)
        
* /usr/lib/libSystem.B.dylib (compatibility version 1.0.0, current version 1281.100.1)
       
* /usr/local/opt/portaudio/lib/libportaudio.2.dylib (compatibility version 3.0.0, current version 3.0.0)
       
* /System/Library/Frameworks/CoreAudio.framework/Versions/A/CoreAudio (compatibility version 1.0.0, current version 1.0.0)
        
* /System/Library/Frameworks/AudioToolbox.framework/Versions/A/AudioToolbox (compatibility version 1.0.0, current version 1000.0.0)
        
* /System/Library/Frameworks/AudioUnit.framework/Versions/A/AudioUnit (compatibility version 1.0.0, current version 1.0.0)
        
* /System/Library/Frameworks/CoreFoundation.framework/Versions/A/CoreFoundation (compatibility version 150.0.0, current version 1677.104.0)
       
* /System/Library/Frameworks/CoreServices.framework/Versions/A/CoreServices (compatibility version 1.0.0, current version 1069.24.0)
       
* /opt/local/lib/libusb-1.0.0.dylib (compatibility version 4.0.0, current version 4.0.0)
        
* /System/Library/Frameworks/Cocoa.framework/Versions/A/Cocoa (compatibility version 1.0.0, current version 23.0.0)
       
* /System/Library/Frameworks/IOKit.framework/Versions/A/IOKit (compatibility version 1.0.0, current version 275.0.0)
        
* /System/Library/Frameworks/CoreVideo.framework/Versions/A/CoreVideo (compatibility version 1.2.0, current version 1.5.0)
        
* /System/Library/Frameworks/OpenGL.framework/Versions/A/OpenGL (compatibility version 1.0.0, current version 1.0.0)
        
* /System/Library/Frameworks/CoreMIDI.framework/Versions/A/CoreMIDI (compatibility version 1.0.0, current version 69.0.0)
        
* /System/Library/Frameworks/Security.framework/Versions/A/Security (compatibility version 1.0.0, current version 59306.140.5)
        
* /System/Library/Frameworks/CoreGraphics.framework/Versions/A/CoreGraphics (compatibility version 64.0.0, current version 1355.22.0)