# Go parameters
CMD=go
BUILD=$(CMD) build

# Output directory
BUILD_DIR=build

# Name of the executable file
BINARY_NAME=main

build:
	@echo "Building $(BINARY_NAME)..."
	go build -o $(BUILD_DIR)/$(BINARY_NAME) main.go

database:
	docker run --name my-postgres -e POSTGRES_PASSWORD=mysecretpassword -d -p 5432:5432 postgres
run: build
	@echo "Running $(BINARY_NAME)..."
	./$(BUILD_DIR)/$(BINARY_NAME)

.PHONY: build run
