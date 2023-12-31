# Define variables
APP_NAME = main
DOCKER_IMAGE = myapp-image

build:
	@echo "Building Go binary..."
	GOOS=linux GOARCH=arm64 go build -o $(APP_NAME)

docker-compose:
	@echo "Shutting down Docker Compose..."
	docker compose down
	@echo "Running Docker Compose..."
	docker compose up -d --build
 
clean:
	@echo "Waiting for 2 seconds before cleaning up..."
	@sleep 2
	@echo "Cleaning up..."
	rm -f $(APP_NAME)

stop:
	@echo "Shutting down Docker Compose..."
	docker compose down

start: build docker-compose clean

.PHONY: build docker-compose clean start
