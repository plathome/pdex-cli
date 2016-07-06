NAME=pdex
HOMEPAGE=https://github.com/plathome/pdex-cli
VERSION=0.1.4-rc
TAG=v$(VERSION)
ARCH=$(shell uname -m)
PREFIX=/usr/local
VETARGS?=-all

all: lint vet test

clean:
	go clean

cleanall:
	rm -rf bin/$(NAME)
	rm -rf build/*
	rm -rf release/*

test: unit acceptance

install: build
	mkdir -p $(PREFIX)/bin
	cp -v bin/$(NAME) $(PREFIX)/bin/$(NAME)

uninstall:
	rm -vf $(PREFIX)/bin/$(NAME)

unit: dependencies
	go test

acceptance: build
	bats test

build: dependencies
	go build -o bin/$(NAME)

build_releases: dependencies
	mkdir -p build/Linux  && GOOS=linux  GOARCH=386 	go build -v -ldflags "-X main.Version=$(VERSION)" -o build/Linux/$(NAME)_$(VERSION)_Linux_x86
	mkdir -p build/Linux  && GOOS=linux  GOARCH=amd64 	go build -v -ldflags "-X main.Version=$(VERSION)" -o build/Linux/$(NAME)_$(VERSION)_Linux_x86_64
	mkdir -p build/Linux  && GOOS=linux  GOARCH=arm 	go build -v -ldflags "-X main.Version=$(VERSION)" -o build/Linux/$(NAME)_$(VERSION)_Linux_arm
	mkdir -p build/Linux  && GOOS=linux  GOARCH=arm64	go build -v -ldflags "-X main.Version=$(VERSION)" -o build/Linux/$(NAME)_$(VERSION)_Linux_arm64

	mkdir -p build/Darwin && GOOS=darwin GOARCH=386 	go build -v -ldflags "-X main.Version=$(VERSION)" -o build/Darwin/$(NAME)_$(VERSION)_Darwin_x86
	mkdir -p build/Darwin && GOOS=darwin GOARCH=amd64 	go build -v -ldflags "-X main.Version=$(VERSION)" -o build/Darwin/$(NAME)_$(VERSION)_Darwin_x86_64

	mkdir -p build/Windows && GOOS=windows GOARCH=386 	go build -v -ldflags "-X main.Version=$(VERSION)" -o build/Windows/$(NAME)_$(VERSION)_Windows-x86.exe
	mkdir -p build/Windows && GOOS=windows GOARCH=amd64 go build -v -ldflags "-X main.Version=$(VERSION)" -o build/Windows/$(NAME)_$(VERSION)_Windows-x86_64.exe

	rm -rf release && mkdir release
	mv build/Linux/$(NAME)_$(VERSION)_Linux_* release/
	mv build/Darwin/$(NAME)_$(VERSION)_Darwin_* release/
	mv build/Windows/* release/

dependencies:
	go get -t
	@go tool cover 2>/dev/null; if [ $$? -eq 3 ]; then \
		go get -u golang.org/x/tools/cmd/cover; \
	fi
	go get github.com/golang/lint/golint

release: build_releases
	go get github.com/progrium/gh-release
	gh-release create plathome/$(NAME) $(VERSION) $(shell git rev-parse --abbrev-ref HEAD)

lint: dependencies
	golint -set_exit_status

# vet runs the Go source code static analysis tool `vet` to find
# any common errors.
vet:
	@go tool vet 2>/dev/null ; if [ $$? -eq 3 ]; then \
		go get golang.org/x/tools/cmd/vet; \
	fi
	@echo "go tool vet $(VETARGS)"
	@go tool vet $(VETARGS) . ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

.PHONY: acceptance build build_releases dependencies install test uninstall unit
