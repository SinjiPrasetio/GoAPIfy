# Load the environment variables from .env file
include .env
export

ifeq ($(OS), Windows_NT)
    RM = del /S /Q 
else
    RM = rm -Rf
endif


# Development server
dev:
	@echo "Running the development server..."
	@go run main.go

# Build server binary
server:
	@echo "Compiling the server..."
	go build -o build/$(APP_NAME)
	@echo "Server compiled successfully."

# Run server binary
run:
	@echo "Running the server..."
	make clean
	make server
	./build/$(APP_NAME)

# Start docker-compose
up:
	@echo "Starting docker-compose..."
	docker-compose up -d
	@echo "Docker-compose started!"

# Stop docker-compose
down:
	@echo "Stopping docker-compose..."
	docker-compose down
	@echo "Docker-compose stopped!"

clean:
	@echo "Cleaning build..."
	$(RM) build
	@echo "Build is cleaned..."

rename:
	@echo "Renaming project..."
	go mod edit -module $(APP_NAME)