BINARY_FILE_NAME=main

start:
	./bin/$(BINARY_FILE_NAME)

build:
	go build -o ./bin/$(BINARY_FILE_NAME) ./cmd/http/main.go

dev:
	air

clean:
	go clean
	rm ./bin/$(BINARY_FILE_NAME)

clear_log:
	rm log/* -f




