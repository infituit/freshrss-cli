.PHONY: build platform-builds build-matrix fmt lint test integration-test clean pre

GOLANGCI_LINT_TIMEOUT ?= 2m

BUILDTIME := $(shell date +%Y-%m-%dT%T%z)
GITTAG := $(shell git describe --tags --always --dirty 2>/dev/null || printf 'v0.0.0')
GITHASH := $(shell git rev-parse --short HEAD 2>/dev/null || printf 'unknown')
VERSION ?= $(GITTAG)-$(shell date +%y.%m.%d)
GOARM ?=
ROOT_DIR := $(CURDIR)
OUTPUT_DIR := $(ROOT_DIR)/build/$(VERSION)
APP_NAME := rss-cli
OSES := linux darwin windows
ARCHES := amd64 arm64
LINUX_ARCHES := amd64 arm64 arm

VERSION_LDFLAGS := -X github.com/infituit/rss-cli/pkg/version.VERSION=$(VERSION) \
	-X github.com/infituit/rss-cli/pkg/version.BUILDTIME=$(BUILDTIME) \
	-X github.com/infituit/rss-cli/pkg/version.GITHASH=$(GITHASH)

build: | pre
	CGO_ENABLED=0 go build -trimpath -ldflags "$(VERSION_LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME) ./cmd/rss-cli

platform-builds: | pre
	@set -e; \
	for os in $(OSES); do \
		arches="$(ARCHES)"; \
		if [ "$$os" = "linux" ]; then arches="$(LINUX_ARCHES)"; fi; \
		for arch in $$arches; do \
			goarm=""; \
			if [ "$$os" = "linux" ] && [ "$$arch" = "arm" ]; then goarm="7"; fi; \
			ext=""; \
			if [ "$$os" = "windows" ]; then ext=".exe"; fi; \
			suffix="$${os}_$${arch}"; \
			if [ -n "$$goarm" ]; then suffix="$${suffix}v$$goarm"; fi; \
			binary="$(OUTPUT_DIR)/$(APP_NAME)_$${suffix}$$ext"; \
			echo "Building $(APP_NAME) $(VERSION) for $$os/$$arch..."; \
			CGO_ENABLED=0 GOOS=$$os GOARCH=$$arch GOARM=$$goarm go build -trimpath -ldflags "$(VERSION_LDFLAGS)" -o "$$binary" ./cmd/rss-cli; \
			echo "Built: $$binary"; \
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
	suffix="$(GOOS)_$(GOARCH)"; \
	if [ -n "$(GOARM)" ]; then suffix="$${suffix}v$(GOARM)"; fi; \
	binary="$(OUTPUT_DIR)/$(APP_NAME)_$${suffix}$$ext"; \
	echo "Building $(APP_NAME) $(VERSION) for $(GOOS)/$(GOARCH)..."; \
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) GOARM=$(GOARM) go build -trimpath -ldflags "$(VERSION_LDFLAGS)" -o "$$binary" ./cmd/rss-cli; \
	echo "Built: $$binary"

pre:
	@mkdir -p $(OUTPUT_DIR)

clean:
	rm -rf $(ROOT_DIR)/build

fmt:
	gofmt -w ./cmd ./internal ./pkg ./testsuite

lint:
	golangci-lint run --config .golangci.yml --timeout $(GOLANGCI_LINT_TIMEOUT)

test:
	go test ./...

integration-test:
	go test -tags=integration ./testsuite/...
