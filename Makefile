.PHONY: all clean install uninstall
BIN := main
HASH := $(shell git rev-parse --short HEAD)
COMMIT_DATE := $(shell git show -s --format=%ci ${HASH})
BUILD_DATE := $(shell date '+%Y-%m-%d %H:%M:%S')
VERSION := ${HASH} (${COMMIT_DATE})
build:
	go build -o bin/${BIN}

run:
	bin/${BIN}

test:
	go test

clean:
	go clean
	rm -f bin/${BIN}

all:
	go build -o bin/${BIN} && bin/${BIN}