VERSION_FILE := internal/version.go
VERSION := $(shell grep 'const Version' $(VERSION_FILE) | sed 's/.*Version = "\(.*\)"/\1/')

BUILD_LINUX_AMD64 := GOOS=linux GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-linux-musl-gcc
BUILD_LINUX_ARM64 := GOOS=linux GOARCH=arm64 CGO_ENABLED=1 CC=aarch64-linux-musl-gcc

BINARY_NAME := octopus

PATH_BIN := bin

BINARY_PATH_AMD64 := $(PATH_BIN)/linux/amd64/$(BINARY_NAME)
BINARY_PATH_ARM64 := $(PATH_BIN)/linux/arm64/$(BINARY_NAME)

MAIN_FILE := cmd/main.go

UNAME_S := $(shell uname -s)

ifeq ($(UNAME_S),Darwin)
    CHOWN_GROUP := wheel
else
    CHOWN_GROUP := root
endif

.PHONY: build-linux-amd64
build-linux-amd64:
	@echo "Building version $(VERSION) linux/amd64 binary..."
	$(BUILD_LINUX_AMD64) go build -o $(BINARY_PATH_AMD64) $(MAIN_FILE)
	@chown root:$(CHOWN_GROUP) $(BINARY_PATH_AMD64)
	@chmod u+s $(BINARY_PATH_AMD64)
	@echo "Done!"

.PHONY: build-linux-arm64
build-linux-arm64:
	@echo "Building version $(VERSION) linux/arm64 binary..."
	$(BUILD_LINUX_ARM64) go build -o $(BINARY_PATH_ARM64) $(MAIN_FILE)
	@chown root:$(CHOWN_GROUP) $(BINARY_PATH_ARM64)
	@chmod u+s $(BINARY_PATH_ARM64)
	@echo "Done!"

clean:
	@echo "Cleaning..."
	rm -rf $(PATH_BIN)
	@echo "Done!"

version:
	@echo "Versión: $(VERSION)"

build: build-linux-amd64 build-linux-arm64