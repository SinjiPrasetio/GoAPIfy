# Define export command based on the operating system
ifeq ($(OS),Windows_NT)
	EXPORT_CMD = set
else
	EXPORT_CMD = export
endif

# Development server
dev:
	@echo "Running the development server..."
	@go run main.go

# Build server binary
build:
	@echo "Building the server..."
	@if [ -z "$(APP_NAME)" ]; then \
		$(EXPORT_CMD) $(grep -v '^#' .env | xargs) && \
		go build -ldflags="-X 'main.AppName=$${APP_NAME}'" -o build/$${APP_NAME}; \
	else \
		go build -ldflags="-X 'main.AppName=$(APP_NAME)'" -o build/$(APP_NAME); \
	fi
	@echo "Build finished!"

# Run server binary
run:
	@echo "Running the server..."
	@$(EXPORT_CMD) $(grep -v '^#' .env | xargs) && \
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

# Rename application and dependencies
rename:
	@echo "Renaming application and dependencies..."
	@OLD_APP_NAME=$$(grep -oP '^APP_NAME=\K.*' .env) && \
	NEW_APP_NAME=$(APP_NAME) && \
	find . -type f -exec sed -i 's/$$OLD_APP_NAME/$$NEW_APP_NAME/g' {} + && \
	if [ -e "./build/$$OLD_APP_NAME" ]; then \
		echo "Renaming server binary from $$OLD_APP_NAME to $$NEW_APP_NAME..." && \
		mv -f "./build/$$OLD_APP_NAME" "./build/$$NEW_APP_NAME"; \
	fi
	@echo "Application and dependencies renamed to $(APP_NAME)!"

