GO=$(or $(shell which go), $(error "Missing dependency - no go in PATH"))
GOROOT := $(shell $(GO) env GOROOT)
GOBINPATH := $(shell $(GO) env GOPATH)/bin

BASE_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

TOOLS_DIR := tools

export PATH := $(BASE_DIR)/$(TOOLS_DIR):$(GOBINPATH):$(PATH)

# Kernel (OS) Name
OS := $(shell uname -s | tr '[:upper:]' '[:lower:]')

# Allow architecture to be overwritten
ifeq ($(HOST_ARCH),)
HOST_ARCH := $(shell uname -m)
endif

ifeq (x86_64, $(HOST_ARCH))
EXEC_ARCH := amd64
else ifeq (arm64, $(HOST_ARCH))
EXEC_ARCH := arm64
endif

# golangci-lint
GOLANGCI_LINT_VERSION=1.63.4
GOLANGCI_LINT_PATH=$(TOOLS_DIR)/golangci-lint-v$(GOLANGCI_LINT_VERSION)
GOLANGCI_LINT_BIN=$(GOLANGCI_LINT_PATH)/golangci-lint
GOLANGCI_LINT_ARCHIVE=golangci-lint-$(GOLANGCI_LINT_VERSION)-$(OS)-$(EXEC_ARCH).tar.gz
GOLANGCI_LINT_ARCHIVEBASE=golangci-lint-$(GOLANGCI_LINT_VERSION)-$(OS)-$(EXEC_ARCH)
export PATH := $(BASE_DIR)/$(GOLANGCI_LINT_PATH):$(PATH)

# goose
GOOSE_VERSION=3.24.1
GOOSE_PATH=$(TOOLS_DIR)/goose-v$(GOOSE_VERSION)
GOOSE_BIN=$(GOOSE_PATH)/goose
GOOSE_ARCHIVE=goose_$(OS)_$(EXEC_ARCH)
GOOSE_ARCHIVEBASE=goose_$(OS)_$(EXEC_ARCH)
export PATH := $(BASE_DIR)/$(GOOSE_PATH):$(PATH)

# wire
WIRE_VERSION=0.6.0
WIRE_BIN := $(or $(shell which wire 2>/dev/null), "$(GOBINPATH)/wire")

# Force Go modules even when checked out inside GOPATH
GO111MODULE := on
export GO111MODULE

# Install tools
.PHONY: tools
tools: $(GOLANGCI_LINT_BIN) $(WIRE_BIN)

# Install golangci-lint
$(GOLANGCI_LINT_BIN):
	@echo "==========> installing golangci-lint v$(GOLANGCI_LINT_VERSION)"
	@mkdir -p "$(GOLANGCI_LINT_PATH)"
	@curl -sSfL "https://github.com/golangci/golangci-lint/releases/download/v$(GOLANGCI_LINT_VERSION)/$(GOLANGCI_LINT_ARCHIVE)" \
		| tar -x -z --strip-components=1 -C "$(GOLANGCI_LINT_PATH)" "$(GOLANGCI_LINT_ARCHIVEBASE)/golangci-lint"

# Install wire
$(WIRE_BIN):
	@echo "==========> installing wire v$(WIRE_VERSION)"
	@if [ ! -f "$(WIRE_BIN)" ]; then \
		$(GO) install github.com/google/wire/cmd/wire@v$(WIRE_VERSION); \
		echo "Wire installed successfully"; \
	else \
		echo "Wire is already installed at $(WIRE_BIN)"; \
	fi

.PHONY: lint
lint: $(GOLANGCI_LINT_BIN)
	@echo "==========> running golangci-lint"
	@"${GOLANGCI_LINT_BIN}" run

.PHONY: wire
wire: $(WIRE_BIN)
	@echo "==========> running wire"
	@${WIRE_BIN} ./cmd/tinder/...

.PHONY: test
test:
	@echo "==========> running tests"
	@$(GO) test -v -race -short ./...

.PHONY: migrate.up
 migrate.up:
	@echo "==========> running goose up"
	@goose --dir ./migrations postgres "host=0.0.0.0 port=5432 user=postgres password=postgres dbname=postgres sslmode=disable" up

.PHONY: migrate.down
 migrate.down:
	@echo "==========> running goose down"
	@goose --dir ./migrations postgres "host=0.0.0.0 port=5432 user=postgres password=postgres dbname=postgres sslmode=disable" down

.PHONY: migrate.down.to
 migrate.down.to:
	@echo "==========> running goose down to $(VERSION)"
	@goose --dir ./migrations postgres "host=0.0.0.0 port=5432 user=postgres password=postgres dbname=postgres sslmode=disable" down-to $(VERSION)

.PHONY: migrate.reset
 migrate.reset:
	@echo "==========> running goose reset"
	@$(MAKE) migrate.down.to VERSION=0 || { echo "Failed to migrate down to version 0"; exit 1; }
	@$(MAKE) -f Makefile migrate.up || { echo "Failed to migrate up"; exit 1; }

.PHONY: docker-compose.infra.up
docker-compose.infra.up:
	@echo "==========> running docker-compose.infra"
	@docker compose -f docker-compose.infra.yaml up -d

.PHONY: docker-compose.infra.down
docker-compose.infra.down:
	@echo "==========> running docker-compose.infra"
	@docker compose -f docker-compose.infra.yaml down --remove-orphans --volumes

.PHONY: docker.build
docker.build:
	@echo "==========> running docker.build"
	@docker build -t tinder-api:latest -t tinder-api:$(shell git rev-parse --short HEAD) .