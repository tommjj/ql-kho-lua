BINARY_FILE_NAME=main

b_start:
	./bin/$(BINARY_FILE_NAME)

b_build:
	go build -o ./bin/$(BINARY_FILE_NAME) ./cmd/http/main.go

build_all: b_build f_build

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
