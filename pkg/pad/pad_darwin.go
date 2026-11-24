//go:build darwin

package pad

/*
#cgo LDFLAGS: -framework CoreFoundation -framework IOKit
#include <CoreFoundation/CoreFoundation.h>
#include <IOKit/IOKitLib.h>

enum { NOVATION_USB_VENDOR_ID = 0x1235 };

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

func checkHardwareOnline() bool {
	return C.isLaunchpadOnlineDarwin() == 1
}
