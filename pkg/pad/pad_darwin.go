//go:build darwin

/*
pad_darwin.go gives your Go program two extra macOS features that the standard MIDI library is missing: noticing when you plug the Launchpad
in late, and spotting right away when you pull the cable out.

1. Listening for the Launchpad to be plugged in

  - The problem: When dmxlights starts up on a Mac, macOS takes a one-time snapshot of connected MIDI gear.
    Because Go programs don't run a normal Mac app window loop, macOS never tells Go when you plug a
    USB device in later—the snapshot stays empty forever.

  - The fix: At startup, pad_darwin.go sets aside one background worker thread and tells macOS,
    "Keep this thread open and send it hardware updates." Even though the listener itself does nothing with those messages,
    simply having it open forces macOS to keep its master MIDI device list fresh whenever you plug in the Launchpad.

2. Checking if the USB cable is actually still connected

  - The problem: When you unplug the Launchpad, macOS's MIDI system often takes a while to
    realize it is gone and keeps showing it as connected.

  - The fix: Every half-second, your program asks the Mac's lower-level USB system (IOKit)
    a simple yes-or-no question: "Is any Novation USB device physically plugged into a USB port right now?"
    As soon as the cable is pulled, the answer flips to "no," and your app immediately knows the Launchpad is offline.
*/
package pad

/*
#cgo LDFLAGS: -framework CoreFoundation -framework CoreMIDI -framework IOKit
#include <CoreFoundation/CoreFoundation.h>
#include <CoreMIDI/CoreMIDI.h>
#include <IOKit/IOKitLib.h>

enum { NOVATION_USB_VENDOR_ID = 0x1235 };

static void midiNotifyCallback(const MIDINotification *message, void *refCon) {
	// Having a registered MIDIClientRef on an active CFRunLoop keeps CoreMIDI's
	// internal device list updated when USB MIDI devices are hot-plugged.
}

static void startCoreMIDIRunLoop() {
	MIDIClientRef client = 0;
	MIDIClientCreate(CFSTR("LaunchpadHotplugWatcher"), midiNotifyCallback, NULL, &client);
	CFRunLoopRun();
}

static int isLaunchpadOnlineDarwin() {
	CFMutableDictionaryRef matching = IOServiceMatching("IOUSBHostDevice");

	if (matching == NULL) {
		return 0;
	}

	io_iterator_t iter = 0;
	if (IOServiceGetMatchingServices(0, matching, &iter) != KERN_SUCCESS) {
		return 0;
	}

	int found = 0;
	io_service_t service;
	while ((service = IOIteratorNext(iter)) != 0) {
		CFTypeRef vidRef = IORegistryEntryCreateCFProperty(service, CFSTR("idVendor"), kCFAllocatorDefault, 0);
		if (vidRef != NULL) {
			SInt32 vid = 0;
			if (CFGetTypeID(vidRef) == CFNumberGetTypeID()) {
				CFNumberGetValue((CFNumberRef)vidRef, kCFNumberSInt32Type, &vid);
			}
			CFRelease(vidRef);

			if (vid == NOVATION_USB_VENDOR_ID) {
				found = 1;
			}
		}
		IOObjectRelease(service);
		if (found) {
			break;
		}
	}

	IOObjectRelease(iter);
	return found;
}
*/
import "C"
import (
	"runtime"
	"time"
)

func init() {
	// Anchor CoreMIDI's first MIDIClientCreate to a dedicated OS thread running CFRunLoopRun().
	started := make(chan struct{})
	go func() {
		runtime.LockOSThread()
		close(started)
		C.startCoreMIDIRunLoop()
	}()
	<-started
	time.Sleep(20 * time.Millisecond)
}

func checkHardwareOnline() bool {
	return C.isLaunchpadOnlineDarwin() == 1
}
