.PHONY: all golint vet fmt test coverage scan build linux osx windows source clean
BUILD_HASH=$(shell git rev-parse HEAD)
BUILD_VERSION=0.3

all: clean get test fmt coverage build source

clean:
	@-rm -rf build/

get:
	# Manually get Windows dependencies (somehow not pulled)
	@go get github.com/inconshreveable/mousetrap
	@go get github.com/konsorten/go-windows-terminal-sequences
	@go get -t -v ./...


fmt:
	@go fmt ./...

test:
	@export UNIT_TEST=1; go test -json ./... > test-report.json

coverage:
	@cd export UNIT_TEST=1; go test -coverprofile=coverage.out ./...

build: linux osx windows windows32 source

LDFLAGS=-ldflags "-w -s -X main.BuildHash=${BUILD_HASH} -X main.BUILDVersion=${BUILD_VERSION}"
linux:
	@mkdir -p build/${BUILD_VERSION}/linux && env GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o build/${BUILD_VERSION}/linux/crawl
	@cd build/${BUILD_VERSION} && mkdir -p release && zip release/linux.zip linux/crawl ../../README.md

osx:
	@mkdir -p build/${BUILD_VERSION}/osx && env GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o build/${BUILD_VERSION}/osx/crawl
	@cd build/${BUILD_VERSION} && mkdir -p release && zip release/osx.zip osx/crawl ../../README.md


windows:
	@mkdir -p build/${BUILD_VERSION}/windows/64 && env GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o build/${BUILD_VERSION}/windows/64/crawl.exe
	@cd build/${BUILD_VERSION} && mkdir -p release && zip release/win64.zip windows/64/crawl.exe ../../README.md

windows32:
	@mkdir -p build/${BUILD_VERSION}/windows/32 && env GOOS=windows GOARCH=386 go build ${LDFLAGS} -o build/${BUILD_VERSION}/windows/32/crawl.exe
	@cd build/${BUILD_VERSION} && mkdir -p release && zip release/win32.zip windows/32/crawl.exe ../../README.md

source:
	@mkdir -p build/${BUILD_VERSION}/release && zip -r build/${BUILD_VERSION}/release/source.zip * -x /build\* crawl crawl-cli /.git\*