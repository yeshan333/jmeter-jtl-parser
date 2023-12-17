build_macos:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o jtl-parser

build_linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o jtl-parser-linux
