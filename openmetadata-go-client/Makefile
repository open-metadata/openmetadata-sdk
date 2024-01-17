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
	go install github.com/mgechev/revive@latest
	go install golang.org/x/tools/cmd/cover@latest

.PHONY: go_unit_tests
go_unit_tests:  ## Run go unit tests
	cd $(ROOT_DIR)/$(GO_DIRECTORY) && go test -v -cover ./...

.PHONY: go_all_tests
go_all_tests:  ## Run go integration tests
	cd $(ROOT_DIR)/$(GO_DIRECTORY) && go test -v -cover -tags=integration ./...

.PHONY: go_check_format
go_check_format:  ## Run go format check
	revive -formatter friendly $(ROOT_DIR)/$(GO_DIRECTORY)/...
	gofmt -l -d $(ROOT_DIR)/$(GO_DIRECTORY)/

.PHONY: go_code_format
go_code_format:  ## Run go code format
	gofmt -w $(ROOT_DIR)/$(GO_DIRECTORY)/

.PHONY: go_check_format_ci
go_check_format_ci:  ## Run go format check for CI/CD. Returns non-zero exit code if there are formatting issues
	files=$$(revive -formatter friendly $(ROOT_DIR)/$(GO_DIRECTORY)/...) && [ -z "$$files" ]
	files=$$(gofmt -l $(ROOT_DIR)/$(GO_DIRECTORY)/) && [ -z "$$files" ]

.PHONY: go_tests
go_tests:  ## Run go tests
	$(MAKE) go_check_format
	$(MAKE) go_all_tests

.PHONY: go_tests_ci
go_tests_ci:  ## Run go tests for CI/CD
	$(MAKE) go_check_format_ci
	$(MAKE) go_all_tests
