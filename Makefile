.PHONY: install check

install:
	go mod download

check:
	gofmt -e -l .
