TAG = $(shell git describe --tags)
BRANCH = $(shell git rev-parse --abbrev-ref HEAD)
HASH = $(shell git rev-parse --short HEAD)
VERSION = $(TAG)-$(BRANCH):$(HASH)
BASENAME = $(shell basename $(shell pwd))
TARGET_DIR = ./bin

all: prepare clean build

build:
	GOOS=windows GOARCH=amd64 go build -ldflags="-X main.version=$(VERSION)" -o $(TARGET_DIR)/$(BASENAME).exe
	GOOS=linux GOARCH=amd64 go build -ldflags="-X main.version=$(VERSION)" -o $(TARGET_DIR)/$(BASENAME)

check: cyclo lint test

clean:
	go clean
	rm -rf $(TARGET_DIR)/*

cyclo:
	-gocyclo -top 10 -avg .

lint:
	go vet ./
	golint ./

prepare:
	@if [ ! -d "$(TARGET_DIR)" ]; then \
		mkdir $(TARGET_DIR); \
	fi

test:
	go test ./


.PHONY: all build check clean cyclo lint test
