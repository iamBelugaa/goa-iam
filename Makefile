BINARY_NAME := iam-service
MAIN_PACKAGE := ./cmd/iam/main.go

BUILD_DIR := ./dist
GOA_GEN_DIR := ./gen
BUILD_FLAGS := -v -ldflags="-s -w"

COVERAGE_DIR := coverage
COVERAGE_PROFILE := coverprofile.out
COVERAGE_HTML := cover.html

# ANSI Color Codes
CYAN := \033[36m
RESET := \033[0m
GREEN := \033[32m
YELLOW := \033[33m
MAGENTA := \033[35m

.PHONY: tidy deps fmt clean gen-goa build run all clean-coverage test coverage

## Tidy Go modules
tidy:
	@echo "$(CYAN)[⚙️ ] Tidying Go modules...$(RESET)"
	@go mod tidy
	@echo "$(GREEN)[✅] Go modules tidied.$(RESET)"

## Download Go dependencies
deps:
	@echo "$(CYAN)[⚙️ ] Downloading Go modules...$(RESET)"
	@go mod download
	@go mod verify
	@echo "$(GREEN)[✅] Go modules downloaded and verified.$(RESET)"

## Format codebase
fmt:
	@echo "$(CYAN)[🎨] Formatting Go code...$(RESET)"
	@go fmt ./...
	@echo "$(GREEN)[✅] Formatting complete.$(RESET)"

## Clean build and generated files
clean:
	@echo "$(YELLOW)[🧹] Removing build artifacts...$(RESET)"
	@go clean
	@rm -rf $(BUILD_DIR)
	@rm -rf $(GOA_GEN_DIR)
	@echo "$(GREEN)[✅] Clean complete.$(RESET)"

## Generate Goa code from design definitions
gen-goa:
	@echo "$(MAGENTA)[⚙️ ] Generating Goa code...$(RESET)"
	@goa gen github.com/iamBelugaa/goa-iam/internal/design
	@goa gen github.com/iamBelugaa/goa-iam/internal/services/authsvc/design -o $(GOA_GEN_DIR)/auth
	@goa gen github.com/iamBelugaa/goa-iam/internal/services/usersvc/design -o $(GOA_GEN_DIR)/user
	@echo "$(GREEN)[✅] Goa code generation complete.$(RESET)"

## Build the binary
build: tidy clean gen-goa
	@echo "$(CYAN)[🔨] Building $(BINARY_NAME) for $(shell go env GOOS)/$(shell go env GOARCH)...$(RESET)"
	@GOOS=$(shell go env GOOS) GOARCH=$(shell go env GOARCH) go build $(BUILD_FLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PACKAGE)
	@echo "$(GREEN)[✅] Build complete: $(BUILD_DIR)/$(BINARY_NAME)$(RESET)"

## Run the built service
run: build
	@echo "$(CYAN)[🚀] Running $(BINARY_NAME)...$(RESET)"
	@$(BUILD_DIR)/$(BINARY_NAME)

## Run full cycle: clean, build, run
all: run

## Clean test coverage files
clean-coverage:
	@echo "$(YELLOW)[🧹] Removing test coverage files...$(RESET)"
	@rm -rf $(COVERAGE_DIR)
	@echo "$(GREEN)[✅] Test coverage cleanup complete.$(RESET)"

## Run tests with Ginkgo
test:
	@echo "$(CYAN)[🧪] Running unit tests...$(RESET)"
	@ginkgo -v -r
	@echo "$(GREEN)[✅] Tests completed.$(RESET)"

## Run tests with coverage report
coverage: clean-coverage
	@echo "$(CYAN)[📊] Running tests with coverage...$(RESET)"
	@mkdir -p $(COVERAGE_DIR)
	@ginkgo -r -v --cover --coverprofile=$(COVERAGE_PROFILE) --output-dir=$(COVERAGE_DIR)
	@go tool cover -html=$(COVERAGE_DIR)/$(COVERAGE_PROFILE) -o $(COVERAGE_DIR)/$(COVERAGE_HTML)
	@echo "$(GREEN)[✅] Coverage report generated at $(COVERAGE_DIR)/$(COVERAGE_HTML)$(RESET)"
	@echo "$(CYAN)[🌐] Opening coverage report in browser...$(RESET)"
	@{ \
		if command -v xdg-open >/dev/null 2>&1; then xdg-open $(COVERAGE_DIR)/$(COVERAGE_HTML); \
		elif command -v open >/dev/null 2>&1; then open $(COVERAGE_DIR)/$(COVERAGE_HTML); \
		else echo "$(YELLOW)[⚠️] Could not detect a command to open browser.$(RESET)"; fi; \
	}