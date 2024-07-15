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

.PHONY: buf
buf:
	@echo "Checking buf..."
	@buf generate
	@mkdir -p $(CURDIR)/ui/proto
	@cp $(CURDIR)/proto/* $(CURDIR)/ui/proto
	@cd $(CURDIR)/ui && protoc -I=$(CURDIR)/ui/proto $(CURDIR)/ui/proto/*.proto --js_out=import_style=commonjs,binary:$(CURDIR)/ui/src/proto --grpc-web_out=import_style=typescript,mode=grpcweb:$(CURDIR)/ui/src/proto
	@rm -rf $(CURDIR)/ui/proto
	@echo "Done!"

.PHONY: app
app: buf
	@echo "Building ui react"
	@sed -i.bak 's/"version": "[^"]*"/"version": "$(VERSION)"/' ui/package.json
	@rm -rf $(CURDIR)/internal/http/public/*
	if [ ! -d "ui/node_modules" ] || [ -z "$$(ls -A ui/node_modules)" ]; then \
  		cd ui && npm install && npm run clean && npm run build; \
	else \
		cd ui && npm run clean && npm run build; \
	fi
	@cp -r $(CURDIR)/ui/public/* $(CURDIR)/internal/http/public/
	@echo "Done!"

.PHONY: build-linux-amd64 app
build-linux-amd64: app
	@echo "Building version $(VERSION) linux/amd64 binary..."
	$(BUILD_LINUX_AMD64) go build -o $(BINARY_PATH_AMD64) $(MAIN_FILE)
	@chown root:$(CHOWN_GROUP) $(BINARY_PATH_AMD64)
	@chmod u+s $(BINARY_PATH_AMD64)
	@echo "Done!"

.PHONY: build-linux-arm64
build-linux-arm64: app
	@echo "Building version $(VERSION) linux/arm64 binary..."
	$(BUILD_LINUX_ARM64) go build -o $(BINARY_PATH_ARM64) $(MAIN_FILE)
	@chown root:$(CHOWN_GROUP) $(BINARY_PATH_ARM64)
	@chmod u+s $(BINARY_PATH_ARM64)
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
build: build-linux-amd64 build-linux-arm64


