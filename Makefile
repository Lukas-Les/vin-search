.PHONY: build test

BINARY_NAME := vin-search

build:
	go build ./cmd/vin-search

test:
	go test -v ./cmd/vin-search
	go test -v ./internal
