.PHONY: help build test 


build:
	go build ./cmd/bkmark

test:
	go test ./...

