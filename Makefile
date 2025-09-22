BINARY_NAME=passgen

# Build for macOS
build:
	go build -o bin/$(BINARY_NAME) ./cmd/passgen

# Build for macOS/Linux/Windows
build-all:
	GOOS=darwin GOARCH=arm64 go build -o bin/$(BINARY_NAME)-darwin-arm64 ./cmd/passgen
	GOOS=linux GOARCH=amd64 go build -o bin/$(BINARY_NAME)-linux-amd64 ./cmd/passgen
	GOOS=windows GOARCH=amd64 go build -o bin/$(BINARY_NAME)-windows-amd64.exe ./cmd/passgen

clean:
	rm -rf bin/

help:
	@echo "Commands:"
	@echo " build      - Build a binary for Mac"
	@echo " build-all  - Build binaries for Linux, Mac, and Windows"
	@echo " clean      - Delete all compiled binaries"