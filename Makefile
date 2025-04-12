# Makefile to build dmxlights.

GO111MODULE=on
COVERAGE = -coverprofile=../c.out -covermode=atomic
SHELL := /usr/bin/env bash
VERSION := "3.0"
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

clean: 
	go clean -cache

init:
	rm -rf go.mod
	rm -rf go.sum
	go mod init
	go mod tidy

dep:
	rm -rf go.mod go.sum 
	go mod init
	go mod tidy

build:
	go mod tidy
	go build dmxlights.go

legacy-deploy:
	rm -rf dmxlights.app/
	fyne package --appVersion 2.0 --id com.github.dhowlett99.dmxlights -os darwin -icon dmxlights.png --tags legacy
	cp Default_Project.yaml dmxlights.app/Contents/Resources/
	cp sequences.yaml dmxlights.app/Contents/Resources/
	cp dmxlights.png dmxlights.app/Contents/Resources/
	cp *.json dmxlights.app/Contents/Resources/

installer:
	go install fyne.io/fyne/v2/cmd/fyne@latest

certificate:
	rm -rf dmxlights.csr dmxlights.rsa dmxlights.crt
	- sudo security delete-certificate -c dmxlights
	openssl genrsa -out dmxlights.rsa 2048
	openssl req -new -key dmxlights.rsa -out dmxlights.csr -subj "/CN=dmxlights" -config openssl.cnf
	openssl req -x509 -new -nodes -key dmxlights.rsa -config /usr/local/etc/openssl/openssl.cnf -config code_sign_csr.conf -subj "/CN=dmxlights" -days 3650 -out dmxlights.crt -extensions v3_ca -extensions codesign_reqext 
	sudo security import dmxlights.rsa -k "/Users/derek/Library/Keychains/login.keychain"
	sudo security add-trusted-cert -r trustRoot -k "/Users/derek/Library/Keychains/login.keychain" dmxlights.crt

deploy: installer
	rm -rf dmxlights.app/
	codesign --remove-signature /usr/local/opt/portaudio/lib/libportaudio.2.dylib
	codesign --force --deep --entitlements entitlements.plist --sign ${CERT} -i ${APP_ID} /usr/local/opt/portaudio/lib/libportaudio.2.dylib
	$(GOPATH)/bin/fyne package --appVersion ${VERSION} --id com.github.dhowlett99.dmxlights -os darwin -icon dmxlights.png -use-raw-icon
	# fix the Info.plist
	./fix.sh dmxlights.app/Contents/Info.plist > /tmp/file
	mv /tmp/file dmxlights.app/Contents/Info.plist
	# sign the app:
	codesign --force --deep --entitlements entitlements.plist --verbose=2 --options runtime --sign ${CERT} -i ${APP_ID} ${APP_NAME}.app
	mkdir -p dmxlights.app/Contents/Resources/projects
	cp projects/Default.yaml dmxlights.app/Contents/Resources/projects
	- ./deploy_presets.sh
	cp groups.yaml dmxlights.app/Contents/Resources/
	cp sequences.yaml dmxlights.app/Contents/Resources/
	cp dmxlights.png dmxlights.app/Contents/Resources/
	cp labels.yaml dmxlights.app/Contents/Resources/


