.DEFAULT_GOAL := build

.PHONY: fmt vet build clean rebuild test test-verbose test-cover test-cover-html

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet test
	go build -o bin/minitodo main.go

clean: 
	rm -rf bin/
	rm -f coverage.out
	rm -f coverage.html
	go clean

rebuild: clean build

test:
	go test ./... -v

test-cover:
	go test ./... -cover

test-cover-html:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report saved to coverage.html"