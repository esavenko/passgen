BINARY_NAME=passgen

# Build for macOS
build:
	go build -o bin/$(BINARY_NAME) .

# Build for macOS/Linux/Windows
build-all:
	GOOS=darwin GOARCH=arm64 go build -o bin/$(BINARY_NAME)-darwin-arm64 .
	GOOS=linux GOARCH=amd64 go build -o bin/$(BINARY_NAME)-linux-amd64 .
	GOOS=windows GOARCH=amd64 go build -o bin/$(BINARY_NAME)-windows-amd64.exe .

install-hooks:
	ln -sf ../../scripts/pre-commit .git/hooks/pre-commit
	ln -sf ../../scripts/commit-msg .git/hooks/commit-msg
	@echo "Git hooks installed"

clean:
	rm -rf bin/

help:
	@echo "Commands:"
	@echo " build      - Build a binary for Mac"
	@echo " build-all  - Build binaries for Linux, Mac, and Windows"
	@echo " clean         - Delete all compiled binaries"
	@echo " install-hooks - Install git pre-commit and commit-msg hooks"