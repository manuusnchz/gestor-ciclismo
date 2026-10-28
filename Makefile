.PHONY: install check

install:
	go mod download

check:
	go build ./...
