BINARY_FILE_NAME=main

b_start:
	./bin/$(BINARY_FILE_NAME)

b_build:
	go build -o ./bin/$(BINARY_FILE_NAME) ./cmd/http/main.go

build_all:
	GOARCH=amd64 GOOS=darwin go build -o ./bin/$(BINARY_FILE_NAME)-darwin ./cmd/http/main.go
	GOARCH=amd64 GOOS=linux go build -o ./bin/$(BINARY_FILE_NAME)-linux ./cmd/http/main.go
	GOARCH=amd64 GOOS=windows go build -o ./bin/$(BINARY_FILE_NAME)-windows ./cmd/http/main.go
	cd ./internal/web; npm run build

b_dev:
	air

f_start:
	cd ./internal/web; npm start

f_build:
	cd ./internal/web; npm run build

f_dev:
	cd ./internal/web; npm run dev

clean:
	go clean
	rm bin/* -f

clear_log:
	rm log/* -f
