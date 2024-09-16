# Variables
DOCKER_IMAGE_NAME = worker
DOCKER_CONTAINER_NAME = worker-container
CONFIG_FILE = internal/config/config.json

# Extract the server port from config.json
SERVER_PORT = $(shell jq -r '.server_port' $(CONFIG_FILE))

.PHONY: all build test clean

all: build test

build:
	@echo "Building the project..."
	go build -o bin/worker cmd/main.go

test-unit:
	@echo "Running unit tests..."
	go test -v ./test/unit

test-integration:
	@echo "Running integration tests..."
	go test -v ./test/integration

clean:
	@echo "Cleaning up..."
	rm -rf bin

docker-build: build
	docker build -t $(DOCKER_IMAGE_NAME) .

docker-run:
	docker compose up -d

run: build
	./bin/worker