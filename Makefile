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


# Define the YAML file path
CONFIG_FILE = config.yaml

# Use grep and awk to extract values from the YAML file under the specified root
DB_HOST = $(shell grep '^database:' -A 6 $(CONFIG_FILE) | grep '^  host:' | awk '{print $$2}')
DB_PORT = $(shell grep '^database:' -A 6 $(CONFIG_FILE) | grep '^  port:' | awk '{print $$2}')
DB_USER = $(shell grep '^database:' -A 6 $(CONFIG_FILE) | grep '^  user:' | awk '{print $$2}')
DB_PASSWORD = $(shell grep '^database:' -A 6 $(CONFIG_FILE) | grep '^  password:' | awk '{print $$2}')
DB_NAME = $(shell grep '^database:' -A 6 $(CONFIG_FILE) | grep '^  dbname:' | awk '{print $$2}')

.PHONY: migrate.up
 migrate.up:
	@echo "==========> running goose up"
	@goose --dir ./migrations postgres "host=$(DB_HOST) port=$(DB_PORT) user=$(DB_USER) password=$(DB_PASSWORD) dbname=$(DB_NAME) sslmode=disable" up

.PHONY: migrate.down
 migrate.down:
	@echo "==========> running goose down"
	@goose --dir ./migrations postgres "host=$(DB_HOST) port=$(DB_PORT) user=$(DB_USER) password=$(DB_PASSWORD) dbname=$(DB_NAME) sslmode=disable" down

.PHONY: migrate.down.to
 migrate.down.to:
	@echo "==========> running goose down to $(VERSION)"
	@goose --dir ./migrations postgres "host=$(DB_HOST) port=$(DB_PORT) user=$(DB_USER) password=$(DB_PASSWORD) dbname=$(DB_NAME) sslmode=disable" down-to $(VERSION)

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

.PHONY: helm.install.postgres
helm.install.postgres:
	@echo "==========> running helm.install.postgres"
	@helm upgrade --install postgres \
		--set global.postgresql.auth.postgresPassword=postgres \
		--set global.postgresql.auth.username=postgres \
		--set global.postgresql.auth.password=postgres \
		--set global.postgresql.auth.database=postgres \
		--set architecture=standalone \
		--set persistence.enabled=false \
		bitnami/postgresql

.PHONY: skaffold.dev
skaffold.dev:
	@echo "==========> running skaffold.dev"
	@kubectl delete cm tinder-config 2>/dev/null
	@kubectl create cm tinder-config --from-file=config.yaml
	@skaffold dev