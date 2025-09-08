# Makefile for common Wishlist API tasks

.PHONY: run build swag tidy test-service test-handler test-repo test-all

# Run the server locally
run:
	go run ./cmd/server

# Compile the binary into bin/server
build:
	go build -o bin/server ./cmd/server

# Generate Swagger docs from main.go into ./docs
swag:
	swag init -g cmd/server/main.go -o docs

# Clean and verify Go modules
tidy:
	go mod tidy && go mod verify

# Unit tests for each layer
test-service:
	go test ./internal/service -v

test-handler:
	go test ./internal/handler -v

test-repo:
	go test ./internal/repository/gorm -v

# Run all tests recursively
test-all:
	go test ./... -v
