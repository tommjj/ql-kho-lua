BINARY_FILE_NAME=main

start:
	./bin/$(BINARY_FILE_NAME)

build:
	go build -o ./bin/$(BINARY_FILE_NAME) ./cmd/http/main.go

build_all:
	GOARCH=amd64 GOOS=darwin go build -o ./bin/$(BINARY_FILE_NAME)-darwin ./cmd/http/main.go
	GOARCH=amd64 GOOS=linux go build -o ./bin/$(BINARY_FILE_NAME)-linux ./cmd/http/main.go
	GOARCH=amd64 GOOS=windows go build -o ./bin/$(BINARY_FILE_NAME)-windows ./cmd/http/main.go

dev:
	air

clean:
	go clean
	rm bin/* -f

clear_log:
	rm log/* -f
