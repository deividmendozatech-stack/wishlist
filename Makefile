.PHONY: run build swag tidy test-service test-handler test-repo test-all

run:
	go run ./cmd/server

build:
	go build -o bin/server ./cmd/server

swag:
	swag init -g cmd/server/main.go -o docs

tidy:
	go mod tidy && go mod verify

test-service:
	go test ./internal/service -v

test-handler:
	go test ./internal/handler -v

test-repo:
	go test ./internal/repository/gorm -v

test-all:
	go test ./... -v
