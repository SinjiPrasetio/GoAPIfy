# Load the environment variables from .env file
include .env
export

# Define the command for deleting files based on the operating system
ifeq ($(OS), Windows_NT)
    RM = del /S /Q 
else
    RM = rm -Rf
endif

.PHONY: key

# Get the path of the current module
MODULE_PATH = $(shell go list -m)

# Update the module dependencies and install any missing modules
install:
	@echo "Checking for .env file..."
	@if exist .env ( \
        echo ".env file found." \
    ) else ( \
        echo ".env file not found. Copying from .env.example..." \
        && copy .env.example .env \
    )
	make key
	go mod tidy

# Run the development server
dev:
	@echo "Running the development server..."
	@go run main.go
	
# Compile the server binary and place it in the build directory
server:
	make clean
	@echo "Compiling the server..."
	go build -o build/$(APP_NAME)
	@echo "Server compiled successfully."

# Clean the build directory, compile the server binary, and run it
run:
	@echo "Running the server..."
	make clean
	make server
	./build/$(APP_NAME)

# Start the Docker-compose containers
up:
	@echo "Starting docker-compose..."
	docker-compose up -d
	@echo "Docker-compose started!"

# Stop the Docker-compose containers
down:
	@echo "Stopping docker-compose..."
	docker-compose down
	@echo "Docker-compose stopped!"

# Delete the build directory
clean:
	@echo "Cleaning build..."
	$(RM) build
	@echo "Build is cleaned..."

# Update the module name to the value of APP_NAME variable in the .env file
rename:
	@echo "Renaming project..."
	go mod edit -module $(APP_NAME)
	go mod tidy

key:
	go run tools/goapi.go key