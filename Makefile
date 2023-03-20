# Define the command for deleting files based on the operating system
ifeq ($(OS), Windows_NT)
	EXE_NAME = "goapi.exe"
else
	EXE_NAME = "goapi"
endif

# Update the module dependencies and install any missing modules
install:
	@echo "Checking for .env file..."
	go mod tidy
	go build -o ${EXE_NAME} tools/goapi.go
	./goapi key