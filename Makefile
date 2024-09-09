BINARY_NAME=qrcode_generator_reader

DEFAULT_QR_CODE_DIR=qr_codes

build:
	@echo "Building the Go binary..."
	@go build -o $(BINARY_NAME) main.go
	@echo "Binary built: $(BINARY_NAME)"

run: build
	@./$(BINARY_NAME)

clean:
	@echo "Cleaning up..."
	@rm -f $(BINARY_NAME)
	@rm -rf $(DEFAULT_QR_CODE_DIR)
	@echo "Clean up completed."

deps:
	@echo "Getting dependencies..."
	@go mod tidy
	@echo "Dependencies installed."

fmt:
	@echo "Formatting Go files..."
	@go fmt ./...
	@echo "Go files formatted."

.PHONY: build run clean deps fmt
