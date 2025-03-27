# Output binary name
BINARY_NAME=main

# Path to main Go file
MAIN_PATH=./cmd/main.go

# Load .env variables
define LOAD_ENV
  export $$(cat .env | grep -v '^#' | xargs)
endef

# Build and run the application with .env variables
run:
	@echo "▶️  Loading .env and building..."
	@$(LOAD_ENV) && go build -o $(BINARY_NAME) $(MAIN_PATH)
	@echo "🚀 Running $(BINARY_NAME)..."
	@$(LOAD_ENV) && ./$(BINARY_NAME)

# Build only (binary output)
build:
	@echo "🔨 Building..."
	@$(LOAD_ENV) && go build -o $(BINARY_NAME) $(MAIN_PATH)

# Delete the compiled binary
clean:
	@echo "🧹 Cleaning up binary..."
	rm -f $(BINARY_NAME)

# Format Go code
fmt:
	@echo "🎨 Formatting code..."
	go fmt ./...

swagger:
	swag init -g cmd/main.go --parseDependency --parseInternal

# Run all tests (unit + e2e)
test:
	go test ./... -v

# Run only unit tests
test-unit:
	@echo "🧪 Running unit tests..."
	@go test ./tests/unit/... -v

# Run only e2e tests
test-e2e:
	@echo "🧪 Running E2E tests with .env.test..."
	@export $$(cat .env.test | grep -v '^#' | xargs) && go test -p 1 ./tests/e2e/... -v

# Run all tests with coverage report in terminal
test-cover:
	go test ./... -v -cover

# Generate HTML coverage report
test-cover-html:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

lint:
	@echo "🔍 Running linter..."
	golangci-lint run ./...

