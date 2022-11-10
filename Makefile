# Makefile to build dmxlights.

COVERAGE = -coverprofile=../c.out -covermode=atomic
SHELL := /usr/bin/env bash

all: test build deploy

legacy: legacy-test legacy-build legacy-deploy

legacy-test:
	go test --tags legacy `go list ./...` ${COVERAGE}
test:
	go test `go list ./...` ${COVERAGE}

legacy-build:
	go build --tags legacy dmxlights.go

build:
	go build dmxlights.go

legacy-deploy:
	fyne package --appVersion 2.0 --id com.github.dhowlett99.dmxlights -os darwin -icon dmxlights.png --tags legacy
	cp fixtures.yaml dmxlights.app/Contents/Resources/
	cp sequences.yaml dmxlights.app/Contents/Resources/
	cp dmxlights.png dmxlights.app/Contents/Resources/
	cp *.json dmxlights.app/Contents/Resources/

deploy:
	fyne package --appVersion 2.0 --id com.github.dhowlett99.dmxlights -os darwin -icon dmxlights.png
	cp fixtures.yaml dmxlights.app/Contents/Resources/
	cp sequences.yaml dmxlights.app/Contents/Resources/
	cp dmxlights.png dmxlights.app/Contents/Resources/
	cp *.json dmxlights.app/Contents/Resources/

