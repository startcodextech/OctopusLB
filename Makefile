VERSION_FILE := internal/version.go
VERSION := $(shell grep 'const Version' $(VERSION_FILE) | sed 's/.*Version = "\(.*\)"/\1/')

BUILD_LINUX_AMD64 := GOOS=linux GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-linux-musl-gcc
BUILD_LINUX_ARM64 := GOOS=linux GOARCH=arm64 CGO_ENABLED=1 CC=aarch64-linux-musl-gcc
BUILD_DARWIN_ARM := GOOS=darwin GOARCH=arm64 CGO_ENABLED=1 CC=clang
BUILD_WINDOWS_AMD64 := GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc
BUILD_WINDOWS_ARM64 := GOOS=windows GOARCH=arm64 CGO_ENABLED=1 CC=aarch64-w64-mingw32-gcc

BINARY_NAME := octopus

PATH_BIN := bin

BINARY_LINUX_PATH_AMD64 := $(PATH_BIN)/linux/amd64/$(BINARY_NAME)
BINARY_LINUX_PATH_ARM64 := $(PATH_BIN)/linux/arm64/$(BINARY_NAME)
BINARY_DARWIN_PATH_ARM64 := $(PATH_BIN)/darwin/arm64/$(BINARY_NAME)
BINARY_WINDOWS_PATH_AMD64 := $(PATH_BIN)/windows/amd64/$(BINARY_NAME)
BINARY_WINDOWS_PATH_ARM64 := $(PATH_BIN)/windows/arm64/$(BINARY_NAME)

MAIN_FILE := cmd/main.go

UNAME_S := $(shell uname -s)

ifeq ($(UNAME_S),Darwin)
    CHOWN_GROUP := wheel
else
    CHOWN_GROUP := root
endif

.PHONY: build-linux-amd64 app
build-linux-amd64:
	@echo "Building version $(VERSION) linux/amd64 binary..."
	$(BUILD_LINUX_AMD64) xcaddy build --with github.com/mholt/caddy-l4 --output $(BINARY_LINUX_PATH_AMD64)-caddy
	$(BUILD_LINUX_AMD64) go build -o $(BINARY_LINUX_PATH_AMD64) $(MAIN_FILE)
	@chown root:$(CHOWN_GROUP) $(BINARY_LINUX_PATH_AMD64)
	@chmod u+s $(BINARY_LINUX_PATH_AMD64)
	@echo "Done!"

.PHONY: build-linux-arm64
build-linux-arm64: app
	@echo "Building version $(VERSION) linux/arm64 binary..."
	$(BUILD_LINUX_ARM64) xcaddy build --with github.com/mholt/caddy-l4 --output $(BINARY_LINUX_PATH_ARM64)-caddy
	$(BUILD_LINUX_ARM64) go build -o $(BINARY_LINUX_PATH_ARM64) $(MAIN_FILE)
	@chown root:$(CHOWN_GROUP) $(BINARY_LINUX_PATH_ARM64)
	@chmod u+s $(BINARY_LINUX_PATH_ARM64)
	@echo "Done!"

.PHONY: build-darwin-arm
build-darwin-arm: app
	@echo "Building version $(VERSION) darwin/arm64 binary..."
	$(BUILD_DARWIN_ARM) xcaddy build --with github.com/mholt/caddy-l4 --output $(BINARY_DARWIN_PATH_ARM64)-caddy
	$(BUILD_DARWIN_ARM) go build -o $(BINARY_DARWIN_PATH_ARM64) $(MAIN_FILE)
	@chown root:$(CHOWN_GROUP) $(BINARY_DARWIN_PATH_ARM64)
	@chmod u+s $(BINARY_DARWIN_PATH_ARM64)
	@echo "Done!"

.PHONY: build-windows-amd64
build-windows-amd64: app
	@echo "Building version $(VERSION) windows/amd64 binary..."
	$(BUILD_WINDOWS_AMD64) xcaddy build --with github.com/mholt/caddy-l4 --output $(BINARY_WINDOWS_PATH_AMD64)-caddy.exe
	$(BUILD_WINDOWS_AMD64) go build -o $(BINARY_WINDOWS_PATH_AMD64).exe $(MAIN_FILE)
	@echo "Done!"

.PHONY: build-windows-arm64
build-windows-arm64: app
	@echo "Building version $(VERSION) windows/arm64 binary..."
	$(BUILD_WINDOWS_ARM64) xcaddy build --with github.com/mholt/caddy-l4 --output $(BINARY_WINDOWS_PATH_ARM64)-caddy.exe
	$(BUILD_WINDOWS_ARM64) go build -o $(BINARY_WINDOWS_PATH_ARM64).exe $(MAIN_FILE)
	@echo "Done!"

.PHONY: clean
clean:
	@echo "Cleaning..."
	rm -rf $(PATH_BIN)
	@echo "Done!"

.PHONY: version
version:
	@echo "Versión: $(VERSION)"

.PHONY: build
build: build-linux-amd64 build-linux-arm64 build-darwin-arm build-windows-amd64 build-windows-arm64


