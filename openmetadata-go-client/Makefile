DIRECTORY_NAME := $(notdir $(CURDIR))
GO_DIRECTORY := openmetadata-go-client

ifeq (openmetadata-go-client,$(DIRECTORY_NAME))
    ROOT_DIR := ..
else
    ROOT_DIR := .
endif


.PHONY: go_install_tools
go_install_tools:  ## Install go tools
	@echo "Installing go tools"
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

.PHONY: go_lint
go_lint:  ## Run go linter
	cd $(ROOT_DIR)/$(GO_DIRECTORY) && golangci-lint run ./...

.PHONY: go_lint_fix
go_lint_fix:  ## Run go linter with auto-fix
	cd $(ROOT_DIR)/$(GO_DIRECTORY) && golangci-lint run --fix ./...

.PHONY: go_unit_tests
go_unit_tests:  ## Run go unit tests
	cd $(ROOT_DIR)/$(GO_DIRECTORY) && go test -v -cover ./pkg/ometa/

.PHONY: go_integration_tests
go_integration_tests:  ## Run go integration tests
	cd $(ROOT_DIR)/$(GO_DIRECTORY) && go test -v ./tests/integration/

.PHONY: go_tests
go_tests:  ## Run linter + all tests
	$(MAKE) go_lint
	$(MAKE) go_unit_tests
	$(MAKE) go_integration_tests
