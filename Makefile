
.PHONY: build-publisher 

build-publisher:
	@go build -o bin/publisher cmd/publisher/main.go

publisher:
	@./bin/publisher