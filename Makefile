.PHONY: build platform-builds build-matrix compress-binary compress-all check-upx fmt lint test clean pre

GOLANGCI_LINT_TIMEOUT ?= 2m

BUILDTIME := $(shell date +%Y-%m-%dT%T%z)
GITTAG := $(shell git describe --tags --always --dirty 2>/dev/null || printf 'v0.0.0')
GITHASH := $(shell git rev-parse --short HEAD 2>/dev/null || printf 'unknown')
VERSION ?= $(GITTAG)-$(shell date +%y.%m.%d)
ROOT_DIR := $(CURDIR)
OUTPUT_DIR := $(ROOT_DIR)/build/$(VERSION)
APP_NAME := freshrss-cli
OSES := linux darwin windows
ARCHES := amd64 arm64
LINUX_ARCHES := amd64 arm64 arm
UPX_ENABLED ?=
UPX_LEVEL ?= 9
UPX_ARGS ?= -$(UPX_LEVEL)
UPX_SUPPORTED_PLATFORMS := linux/amd64 linux/arm64 linux/arm windows/amd64

VERSION_LDFLAGS := -X github.com/infituit/freshrss-cli/pkg/version.VERSION=$(VERSION) \
	-X github.com/infituit/freshrss-cli/pkg/version.BUILDTIME=$(BUILDTIME) \
	-X github.com/infituit/freshrss-cli/pkg/version.GITHASH=$(GITHASH)

build: | pre
	CGO_ENABLED=0 go build -trimpath -ldflags "$(VERSION_LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME) ./cmd/freshrss-cli
	@if [ -n "$(UPX_ENABLED)" ]; then \
		$(MAKE) compress-binary BINARY=$(OUTPUT_DIR)/$(APP_NAME); \
	fi

platform-builds: | pre
	@set -e; \
	for os in $(OSES); do \
		arches="$(ARCHES)"; \
		if [ "$$os" = "linux" ]; then arches="$(LINUX_ARCHES)"; fi; \
		for arch in $$arches; do \
			ext=""; \
			if [ "$$os" = "windows" ]; then ext=".exe"; fi; \
			binary="$(OUTPUT_DIR)/$(APP_NAME)_$${os}_$${arch}$$ext"; \
			echo "Building $(APP_NAME) $(VERSION) for $$os/$$arch..."; \
			CGO_ENABLED=0 GOOS=$$os GOARCH=$$arch go build -trimpath -ldflags "$(VERSION_LDFLAGS)" -o "$$binary" ./cmd/freshrss-cli; \
			echo "Built: $$binary"; \
			if [ -n "$(UPX_ENABLED)" ]; then \
				$(MAKE) compress-binary BINARY="$$binary" PLATFORM=$$os/$$arch; \
			fi; \
		done; \
	done

build-matrix: | pre
	@if [ -z "$(GOOS)" ] || [ -z "$(GOARCH)" ]; then \
		echo "Error: GOOS and GOARCH must be specified."; \
		echo "Example: make build-matrix GOOS=linux GOARCH=amd64"; \
		exit 1; \
	fi
	@ext=""; \
	if [ "$(GOOS)" = "windows" ]; then ext=".exe"; fi; \
	binary="$(OUTPUT_DIR)/$(APP_NAME)_$(GOOS)_$(GOARCH)$$ext"; \
	echo "Building $(APP_NAME) $(VERSION) for $(GOOS)/$(GOARCH)..."; \
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -trimpath -ldflags "$(VERSION_LDFLAGS)" -o "$$binary" ./cmd/freshrss-cli; \
	echo "Built: $$binary"; \
	if [ -n "$(UPX_ENABLED)" ]; then \
		$(MAKE) compress-binary BINARY="$$binary" PLATFORM=$(GOOS)/$(GOARCH); \
	fi

pre: check-upx
	@mkdir -p $(OUTPUT_DIR)

check-upx:
	@if [ -n "$(UPX_ENABLED)" ] && ! command -v upx >/dev/null 2>&1; then \
		echo "Error: UPX compression requested but upx was not found in PATH"; \
		echo "Run 'mise install' or set UPX_ENABLED="; \
		exit 1; \
	fi

compress-binary:
	@if [ -z "$(BINARY)" ]; then \
		echo "Error: BINARY path not specified"; \
		echo "Usage: make compress-binary BINARY=path/to/binary"; \
		exit 1; \
	fi
	@if [ ! -f "$(BINARY)" ]; then \
		echo "Error: Binary file not found: $(BINARY)"; \
		exit 1; \
	fi
	@if [ -n "$(PLATFORM)" ] && ! echo " $(UPX_SUPPORTED_PLATFORMS) " | grep -q " $(PLATFORM) "; then \
		echo "Skipping UPX compression for unsupported platform $(PLATFORM)"; \
	else \
		echo "Compressing $(BINARY) with UPX $(UPX_ARGS)..."; \
		upx $(UPX_ARGS) "$(BINARY)"; \
		echo "Compression complete"; \
	fi

compress-all: check-upx
	@if [ -n "$(UPX_ENABLED)" ]; then \
		for binary in $(OUTPUT_DIR)/$(APP_NAME)*; do \
			[ -f "$$binary" ] || continue; \
			upx $(UPX_ARGS) "$$binary"; \
		done; \
	else \
		echo "UPX compression disabled. Set UPX_ENABLED=1 to enable"; \
	fi

clean:
	rm -rf $(ROOT_DIR)/build

fmt:
	gofmt -w ./cmd ./internal ./pkg

lint:
	golangci-lint run --config .golangci.yml --timeout $(GOLANGCI_LINT_TIMEOUT)

test:
	go test ./...
