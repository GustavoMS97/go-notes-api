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

# Run tests (if available)
test:
	@echo "🧪 Running tests..."
	go test ./...

swagger:
	swag init -g cmd/main.go --parseDependency --parseInternal
