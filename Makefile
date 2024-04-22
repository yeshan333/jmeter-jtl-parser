all: build_macos build_linux

build_macos:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/jtl-parser

build_linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/jtl-parser-linux
