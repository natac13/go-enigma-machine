BINARY_NAME=go-enigma-cli
MAIN=.

.PHONY: all build run test clean

all: run

build:
	go build -o bin/$(BINARY_NAME) $(MAIN)

run: build
	./bin/$(BINARY_NAME)

test:
	go test ./...

test-coverage:
	go test -coverprofile=coverage.out -coverpkg ./pkg/enigma ./...

test-coverage-html: test-coverage
	go tool cover -html=coverage.out

clean:
	rm -f bin/$(BINARY_NAME)
