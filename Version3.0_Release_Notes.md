# dmxlights Version 3.0 Release Notes.

# Introduction 

DMX lights is a Open Source Project to control DMX light fixtures. It is designed for the small to medium mobile disco rig and be very easy to setup and use.

Firstly we use the Novation Launchpad Mini (Mk3) as the control surface and a cheap FTDI RS422 controller as the interface for the DMX lights.

Second design goal was to do away with having to have a large fixture library, with DMX lights you simply add the description of what channels do what to a simple text file.

As of versions 3.0 of dmxlights we now support ...

This version of DMX Lights was developed on macOS 10.15.7 (19H2026)

## New in Version 3.0

* Updated Manual.
* Selectable switch sequences and ability to override parameters.
* Central color definition in its own package 
* Button functions broken into their own files.
* Sequencer performance fixes.
* Sequencer functions broken into their own files.
* Color Display show sequences current color(s).
* Updates Comscans project.
* Buttons are labels in the GUI from labels.yaml
* We now correctly count the number of fixtures in a sequence. Max still 8.
* Presets are now saved inside a project, so presets relate to fixtures correctly.
* Support for Equinox Helix
* Bug fix to entering ranges in settings panel.

## Known Problems

* The Novation Launchpad can only take a certain amount of notes per second, its possible to crash the launchpad if you have a lot of activity. The workaround is that DMX lights will use the reset message from the launchpad, then reset the launchpad and refresh the current state of the launchpad lamps. You may not even notice that the crash as happened.

* When you have the scanner sequence selected, the top left clear button is used for tilting up the scanners. The clear button is activated by a long press instead of a short press. In the GUI the tilt up is not available and the clear is always a single short press.

* Presets can be deleted by a long press on the launchpad, but you cannot delete a preset from the GUI.

* When opening the open or save project you may see an error message, this is because Fyne.io hard codes the favorites locations. 

    In my case I didn't have a 'Videos' directory in my home directory so I got:-
   
    **Fyne error:  Getting List favorite locations
    At: /Users/derek/project/pkg/mod/fyne.io/fyne/v2@v2.5.1/dialog/file_darwin.go:40
    Fyne error:  Getting favorite locations
    Cause: uri is not listable
    At: /Users/derek/project/pkg/mod/fyne.io/fyne/v2@v2.5.1/dialog/file.go:359**

    Workaround To avoid this error you need to have the following directories in your "Home" dirctory:- "Documents","Desktop","Downloads","Music","Pictures","Movies","Videos"

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