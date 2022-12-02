
# LibUSB suppory macOS Catalina(10.15) and later

https://github.com/libusb/libusb/wiki/FAQ#How_can_I_run_libusb_applications_under_Mac_OS_X_if_there_is_already_a_kernel_extension_installed_for_the_device_and_claim_exclusive_access

Take note the location may be different in different macOS version.

For macOS Catalina(10.15) and later, Apple FTDI driver is now a Driverkit extension named com.apple.DriverKit-AppleUSBFTDI. (Note: it seems to work with libusb out of the box under Mac Mini M1 Big Sur even though the driver and the serial port are present).

Updates in Feb 2022: with the merging of pull request #911, support for detaching kernel drivers with authorization has been added. This has been included in the upcoming 1.0.25 release. Take note the underlying macOS capture APIs only work on the whole device and not on individual interfaces. So this will force all the kernel extensions (drivers) bound to all the interfaces of a USB Composite devices to be released (Issue #920). You will need to get the entitlement from Apple Developer support, create a provisioning profile with that entitlement, and build your app with that profile.

Please take note that command line apps cannot use provisioning profiles and therefore cannot hold this entitlement so you have to run as root. As for the GUI application, as of now APple support is saying that `com.apple.vm.device-access` entitlement is not the right entitlement. So basically there is no solution as of now. This means libusb (and libraries like libuvc which uses libusb) may not be the right library to use if you hit this issue (Issue #972 and Issue #1014). 