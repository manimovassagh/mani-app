# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build

# Output directory
BUILDDIR=build

# Name of the executable file
BINARY_NAME=main

build:
	@echo "Building $(BINARY_NAME)..."
	go build -o $(BUILDDIR)/$(BINARY_NAME) main.go

run: build
	@echo "Running $(BINARY_NAME)..."
	./$(BUILDDIR)/$(BINARY_NAME)

.PHONY: build run
