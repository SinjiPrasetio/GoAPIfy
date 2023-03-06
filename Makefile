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
