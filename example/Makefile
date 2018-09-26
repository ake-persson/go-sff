all: linux

clean:
	rm -f example

linux: clean
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

.PHONY: linux
