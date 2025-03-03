.PHONY: build run clean lint lint-fix test

# Default target
all: build

# Build the program
build:
	go build -o error-demo .

# Run the program
run: build
	./error-demo

# Clean up build artifacts
clean:
	rm -f error-demo
	go clean

# Run the error linter to check issues
lint:
	go tool go-errorlint -errorf ./...

# Run the error linter with auto-fix
lint-fix:
	go tool go-errorlint -fix ./...

# Run tests (if you add them later)
test:
	go test -v ./...