# Makefile

SAMPLES := $(wildcard ./example/*)

all: test build

build: $(SAMPLES)

test:
	go test ./common/...

$(SAMPLES):
	cd $@ && go vet && go build

.PHONY: all build test $(SAMPLES)

