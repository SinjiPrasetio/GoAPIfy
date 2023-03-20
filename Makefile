# Load the environment variables from .env file
include .env
export

# Define the command for deleting files based on the operating system
ifeq ($(OS), Windows_NT)
    RM = del /S /Q 
	EXE_NAME = "goapi.exe"
else
    RM = rm -Rf
	EXE_NAME = "goapi"
endif

.PHONY: key

# Get the path of the current module
MODULE_PATH = $(shell go list -m)

# Update the module dependencies and install any missing modules
install:
	@echo "Checking for .env file..."
	go mod tidy
	go build -o ${EXE_NAME} tools/goapi.go
	./goapi key