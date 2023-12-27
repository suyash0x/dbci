.PHONY: build-publisher build-buildServer buildServer publisher all

all: build-publisher build-buildServer

build-publisher:
	@go build -o bin/publisher cmd/publisher/main.go

build-buildServer:
	@go build -o bin/buildServer cmd/buildServer/main.go

publisher:
	@./bin/publisher

buildServer:
	@./bin/buildServer