# Makefile to build dmxlights.

COVERAGE = -coverprofile=../c.out -covermode=atomic
SHELL := /usr/bin/env bash

all: test build

test:
	go test --tags legacy `go list ./...` ${COVERAGE}

build:
	go build --tags legacy dmxlights.go

deploy:
	fyne package --appVersion 2.0 --id com.github.dhowlett99.dmxlights -os darwin -icon dmxlights.png --tags legacy
	cp fixtures.yaml dmxlights.app/Contents/Resources/
	cp sequences.yaml dmxlights.app/Contents/Resources/
	cp dmxlights.png dmxlights.app/Contents/Resources/
	cp *.json dmxlights.app/Contents/Resources/

