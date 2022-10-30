# Makefile to build dmxlights.

COVERAGE = -coverprofile=../c.out -covermode=atomic
SHELL := /usr/bin/env bash

all: test build

test:
	go test --tags legacy `go list ./...` ${COVERAGE}

build:
	go build --tags legacy dmxlights.go

