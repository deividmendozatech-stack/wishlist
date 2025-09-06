.PHONY: run build swag tidy \
        test-service test-handler test-repo test-all

# Levanta el servidor local
run:
	go run ./cmd/server

# Compila el binario en ./bin/server
build:
	go build -o bin/server ./cmd/server

# Genera documentación Swagger en ./docs
swag:
	swag init -g cmd/server/main.go -o docs

# Limpia y actualiza dependencias
tidy:
	go mod tidy && go mod verify

# Tests unitarios de la capa Service
test-service:
	go test ./internal/service -v

# Tests unitarios de la capa Handler
test-handler:
	go test ./internal/handler -v

# Tests de integración de la implementación GORM
test-repo:
	go test ./internal/repository/gorm -v

# Ejecuta todos los tests del proyecto
test-all:
	go test ./... -v

--------USO-----------

Ejecutar servidor:

make run

Generar Swagger:

make swag

Correr todos los tests:

make test-all

Solo service:

make test-service
