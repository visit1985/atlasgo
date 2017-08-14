# Makefile

all: test build

build:
	cd ./example/client; go build

test:
	go test ./common/credentials
	go test ./common/client

.PHONY: all build test

