BINARY_NAME=go-enigma-cli
MAIN=./cmd/enigma

.PHONY: all build run test clean

all: run

build:
	go build -o bin/$(BINARY_NAME) $(MAIN)

run: build
	./bin/$(BINARY_NAME)

test:
	go test ./test/...

test-coverage:
	go test -cover ./...

clean:
	rm -f bin/$(BINARY_NAME)
