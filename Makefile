.PHONY: build test

BINARY_NAME := vin-search

build:
	cd src && go build -o ../$(BINARY_NAME) .

test:
	cd src && go test -v
