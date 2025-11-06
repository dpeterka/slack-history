.PHONY: build run test clean fmt lint test-run docker-build docker-run help

# Binary name
BINARY_NAME=history-slackbot
BINARY_PATH=bin/$(BINARY_NAME)

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOFMT=gofmt
GOLINT=golint

# Build the application
build:
	@echo "Building..."
	@mkdir -p bin
	$(GOBUILD) -o $(BINARY_PATH) cmd/bot/main.go
	@echo "Build complete: $(BINARY_PATH)"

# Run the application
run: build
	@echo "Running..."
	@if [ -f .env ]; then \
		set -a && . ./.env && set +a && ./$(BINARY_PATH); \
	else \
		echo "Warning: .env file not found. Copy .env.example to .env and configure it."; \
		./$(BINARY_PATH); \
	fi

# Run once (for testing)
test-run: build
	@echo "Running once for testing..."
	@if [ -f .env ]; then \
		set -a && . ./.env && set +a && RUN_ONCE=true ./$(BINARY_PATH); \
	else \
		echo "Warning: .env file not found. Copy .env.example to .env and configure it."; \
		RUN_ONCE=true ./$(BINARY_PATH); \
	fi

# Run tests
test:
	@echo "Running tests..."
	$(GOTEST) -v ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	$(GOTEST) -v -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Clean build artifacts
clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	rm -rf bin/
	rm -f coverage.out coverage.html
	@echo "Clean complete"

# Format code
fmt:
	@echo "Formatting code..."
	$(GOFMT) -w .
	@echo "Format complete"

# Lint code
lint:
	@echo "Linting code..."
	$(GOLINT) ./...

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	$(GOGET) -v ./...
	@echo "Dependencies downloaded"

# Tidy dependencies
tidy:
	@echo "Tidying dependencies..."
	$(GOCMD) mod tidy
	@echo "Dependencies tidied"

# Build Docker image
docker-build:
	@echo "Building Docker image..."
	docker build -t $(BINARY_NAME):latest .
	@echo "Docker image built: $(BINARY_NAME):latest"

# Run Docker container
docker-run:
	@echo "Running Docker container..."
	docker run --env-file .env $(BINARY_NAME):latest

# Stop Docker container
docker-stop:
	@echo "Stopping Docker container..."
	docker stop $(BINARY_NAME) || true
	docker rm $(BINARY_NAME) || true

# Build for multiple platforms
build-all:
	@echo "Building for multiple platforms..."
	@mkdir -p bin
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o bin/$(BINARY_NAME)-linux-amd64 cmd/bot/main.go
	GOOS=linux GOARCH=arm64 $(GOBUILD) -o bin/$(BINARY_NAME)-linux-arm64 cmd/bot/main.go
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o bin/$(BINARY_NAME)-darwin-amd64 cmd/bot/main.go
	GOOS=darwin GOARCH=arm64 $(GOBUILD) -o bin/$(BINARY_NAME)-darwin-arm64 cmd/bot/main.go
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o bin/$(BINARY_NAME)-windows-amd64.exe cmd/bot/main.go
	@echo "Multi-platform build complete"

# Show help
help:
	@echo "Available targets:"
	@echo "  build          - Build the application"
	@echo "  run            - Build and run the application"
	@echo "  test-run       - Build and run once (for testing)"
	@echo "  test           - Run tests"
	@echo "  test-coverage  - Run tests with coverage report"
	@echo "  clean          - Clean build artifacts"
	@echo "  fmt            - Format code"
	@echo "  lint           - Lint code"
	@echo "  deps           - Download dependencies"
	@echo "  tidy           - Tidy dependencies"
	@echo "  docker-build   - Build Docker image"
	@echo "  docker-run     - Run Docker container"
	@echo "  docker-stop    - Stop Docker container"
	@echo "  build-all      - Build for multiple platforms"
	@echo "  help           - Show this help message"
