# Simple Makefile for a Go project

# Build the application
all: build test

build:
	@echo "Building..."
	
	@go build -o main api/main.go

# Run the application
run:
	@go run api/main.go

# Test the application
test:
	@echo "Testing..."
	@go test ./... -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Live Reload
watch:
	@if command -v air > /dev/null; then \
            air; \
            echo "Watching...";\
        else \
            read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
            if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
                go install github.com/air-verse/air@latest; \
                air; \
                echo "Watching...";\
            else \
                echo "You chose not to install air. Exiting..."; \
                exit 1; \
            fi; \
        fi

# Watch and run tests continuously
watch-test:
	@if command -v entr > /dev/null; then \
            find . -name '*.go' | entr -c go test ./... -v; \
            echo "Watching tests...";\
        else \
            read -p "'entr' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
            if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
                brew install entr; \
                find . -name '*.go' | entr -c go test ./... -v; \
                echo "Watching tests...";\
            else \
                echo "You chose not to install entr. Exiting..."; \
                exit 1; \
            fi; \
        fi

.PHONY: all build run test clean watch watch-test
