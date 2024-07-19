.PHONY: build test

build:
    ./scripts/build.sh

test:
	go test -v ./cmd/vin-search
	go test -v ./internal
