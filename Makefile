BINARY_NAME := iam-service
MAIN_PACKAGE := ./cmd/iam/main.go

BUILD_DIR := ./dist
GOA_GEN_DIR := ./gen
BUILD_FLAGS := -v -ldflags="-s -w"

COVERAGE_DIR := coverage
COVERAGE_PROFILE := coverprofile.out
COVERAGE_HTML := cover.html

# Docker variables
DOCKER_TAG ?= latest
DOCKER_IMAGE_NAME := iam-service
DOCKERFILE_PATH := ./deploy/docker/Dockerfile

.PHONY: tidy deps install-goa fmt clean gen-goa build docker-build dockerize docker-run run all clean-coverage test coverage

## Tidy Go modules
tidy:
	@echo "Tidying Go modules..."
	@go mod tidy
	@echo "Go modules tidied."

## Download Go dependencies
deps:
	@echo "Downloading Go modules..."
	@go mod download
	@go mod verify
	@echo "Go modules downloaded and verified."

## Install Goa framework
install-goa:
	@echo "Installing Goa framework..."
	@go install goa.design/goa/v3/cmd/goa@latest
	@echo "Goa installation complete."

## Install Ginkgo framework
install-ginkgo:
	@echo "Installing Ginkgo framework..."
	@go install github.com/onsi/ginkgo/v2/ginkgo
	@echo "Goa installation complete."

## Format codebase
fmt:
	@echo "Formatting Go code..."
	@go fmt ./...
	@echo "Formatting complete."

## Clean build and generated files
clean:
	@echo "Removing build artifacts..."
	@go clean
	@rm -rf $(BUILD_DIR)
	@rm -rf $(GOA_GEN_DIR)
	@echo "Clean complete."

## Generate Goa code from design definitions
gen-goa:
	@echo "Generating Goa code..."
	@goa gen github.com/iamBelugaa/goa-iam/internal/design
	@echo "Goa code generation complete."

## Build the binary
build: clean gen-goa
	@echo "Building $(BINARY_NAME) for $(shell go env GOOS)/$(shell go env GOARCH)..."
	GOOS=$(shell go env GOOS) GOARCH=$(shell go env GOARCH) go build $(BUILD_FLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PACKAGE)
	@echo "Build complete: $(BUILD_DIR)/$(BINARY_NAME)"

## Build the binary for Docker (Linux)
docker-build: clean gen-goa
	@echo "Building $(BINARY_NAME) for linux/amd64..."
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
	-ldflags='-w -s -extldflags "-static"' \
	-a -installsuffix cgo \
	-o $(BINARY_NAME) $(MAIN_PACKAGE)
	@echo "Docker build complete: $(BINARY_NAME)"

## Run the built service
run: build
	@echo "Running $(BINARY_NAME)..."
	@$(BUILD_DIR)/$(BINARY_NAME)

## Run full cycle: clean, build, run
all: run

## Clean test coverage files
clean-coverage:
	@echo "Removing test coverage files..."
	@rm -rf $(COVERAGE_DIR)
	@echo "Test coverage cleanup complete."

## Run tests with Ginkgo
test:
	@echo "Running unit tests..."
	@ginkgo -v -r
	@echo "Tests completed."

## Run tests with coverage report
coverage: clean-coverage
	@echo "Running tests with coverage..."
	@mkdir -p $(COVERAGE_DIR)
	@ginkgo -r -v --cover --coverprofile=$(COVERAGE_PROFILE) --output-dir=$(COVERAGE_DIR)
	@go tool cover -html=$(COVERAGE_DIR)/$(COVERAGE_PROFILE) -o $(COVERAGE_DIR)/$(COVERAGE_HTML)
	@echo "Coverage report generated at $(COVERAGE_DIR)/$(COVERAGE_HTML)"

## Build Docker image
dockerize:
	@echo "Building Docker image..."
	@FULL_IMAGE_NAME="$(DOCKER_IMAGE_NAME):$(DOCKER_TAG)"; \
	echo "Building image: $$FULL_IMAGE_NAME"; \
	docker build -f $(DOCKERFILE_PATH) -t $$FULL_IMAGE_NAME .
	@echo "Docker image build complete."

## Run Docker container
docker-run: dockerize
	@echo "Running Docker container..."
	@FULL_IMAGE_NAME="$(DOCKER_IMAGE_NAME):$(DOCKER_TAG)"; \
	echo "Starting container from image: $$FULL_IMAGE_NAME"; \
	docker run --rm -p 8080:8080 --name $(DOCKER_IMAGE_NAME)-container $$FULL_IMAGE_NAME