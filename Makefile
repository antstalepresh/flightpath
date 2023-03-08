SHELL=/bin/bash

.PHONY: build
build:
	CGO_ENABLED=0 go build -a -o ./build/flightpath cmd/main.go

.PHONY: test
test:
	go test -v -count=1 ./...

.PHONY: run
run:
	go run cmd/main.go

.PHONY: lint
lint:
	golangci-lint run ./... --fix