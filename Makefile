APP := wishlist
BIN := bin/$(APP)
PORT := 8080

.PHONY: build run test clean swagger docker-up docker-down

build:
	go build -o $(BIN) ./cmd/server

run: build
	./$(BIN)

test:
	go test -v ./...

swagger:
	swag init -g cmd/server/main.go -o docs

clean:
	rm -rf bin

docker-up:
	docker build -t $(APP):latest .
	docker run --rm -it -p $(PORT):8080 --env-file .env $(APP):latest

docker-down:
	docker stop $$(docker ps -q --filter ancestor=$(APP):latest) || true
