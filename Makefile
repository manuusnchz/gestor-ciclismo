.PHONY: install check test

install:
	go mod download

check:
	gofmt -e -l .

test:
	go test ./...
