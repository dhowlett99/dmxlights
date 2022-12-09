# Makefile to build dmxlights.

GO111MODULE=on
COVERAGE = -coverprofile=../c.out -covermode=atomic
SHELL := /usr/bin/env bash
export PKG_CONFIG_PATH=/usr/local/Cellar/portaudio/19.7.0/lib/pkgconfig

# The name of the application
APP_NAME="dmxlights"
# The bundle / app ID for the app...
APP_ID="com.github.dhowlett99.dmxlights"
# This is the CN from the code signing cert
CERT="dmxlights"

all: test build deploy

legacy: legacy-test legacy-build legacy-deploy

legacy-test:
	go test --tags legacy `go list ./...` ${COVERAGE}
test:
	go test `go list ./...` ${COVERAGE}

legacy-build:
	go build --tags legacy dmxlights.go

build:
	go mod tidy
	go build dmxlights.go

legacy-deploy:
	rm -rf dmxlights.app/
	fyne package --appVersion 2.0 --id com.github.dhowlett99.dmxlights -os darwin -icon dmxlights.png --tags legacy
	cp fixtures.yaml dmxlights.app/Contents/Resources/
	cp sequences.yaml dmxlights.app/Contents/Resources/
	cp dmxlights.png dmxlights.app/Contents/Resources/
	cp *.json dmxlights.app/Contents/Resources/

deploy:
	rm -rf dmxlights.app/

	codesign --remove-signature /usr/local/opt/portaudio/lib/libportaudio.2.dylib
	codesign --force --deep --entitlements entitlements.plist --sign ${CERT} -i ${APP_ID} /usr/local/opt/portaudio/lib/libportaudio.2.dylib
	fyne package --appVersion 2.0 --id com.github.dhowlett99.dmxlights -os darwin -icon dmxlights.png
	# fix the Info.plist
	./fix.sh dmxlights.app/Contents/Info.plist > /tmp/file
	mv /tmp/file dmxlights.app/Contents/Info.plist
	# sign the app:
	codesign --force --deep --entitlements entitlements.plist --verbose=2 --options runtime --sign ${CERT} -i ${APP_ID} ${APP_NAME}.app
	cp fixtures.yaml dmxlights.app/Contents/Resources/
	cp sequences.yaml dmxlights.app/Contents/Resources/
	cp dmxlights.png dmxlights.app/Contents/Resources/
	cp *.json dmxlights.app/Contents/Resources/

