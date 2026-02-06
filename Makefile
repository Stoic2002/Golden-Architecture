.PHONY: build run test clean tidy migrate-up migrate-down

# Variables
APP_NAME=todo-api
MAIN_PATH=./cmd/api
BUILD_DIR=./bin

# Build the application
build:
	@echo "Building $(APP_NAME)..."
	@go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_PATH)

# Run the application
run:
	@echo "Running $(APP_NAME)..."
	@go run $(MAIN_PATH)/main.go

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)

# Tidy dependencies
tidy:
	@echo "Tidying dependencies..."
	@go mod tidy

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	@go mod download

# Run database migrations up
migrate-up:
	@echo "Running migrations up..."
	@./scripts/migrate.sh up

# Run database migrations down
migrate-down:
	@echo "Running migrations down..."
	@./scripts/migrate.sh down

# Format code
fmt:
	@echo "Formatting code..."
	@go fmt ./...

# Run linter
lint:
	@echo "Running linter..."
	@golangci-lint run

# All-in-one dev setup
dev: tidy run
