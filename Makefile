.DEFAULT_GOAL := build

.PHONY: fmt vet build clean rebuild
fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build -o bin/minitodo main.go

clean: 
	rm -rf bin/
	go clean

rebuild: clean build