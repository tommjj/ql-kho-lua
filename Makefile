BINARY_FILE_NAME=main

api_start:
	./bin/$(BINARY_FILE_NAME)

api_build:
	go build -o ./bin/$(BINARY_FILE_NAME) ./cmd/http/main.go

build_all: b_build f_build

api_dev:
	air

app_start:
	cd ./internal/web; npm start

app_build:
	cd ./internal/web; npm run build

app_dev:
	cd ./internal/web; npm run dev

up:
	docker compose up --build   

clean:
	go clean
	rm bin/* -f

clear_log:
	rm log/* -f
