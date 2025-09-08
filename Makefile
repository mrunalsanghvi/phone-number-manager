# Makefile for Go Web Server + Docker

# App name (used for binary and Docker image)
APP_NAME := go-simple-server
PORT := 8080

# Go build flags
GO_BUILD_FLAGS := -o $(APP_NAME)

# Docker image name
DOCKER_IMAGE := $(APP_NAME):latest

# Default target
.PHONY: all
all: build

# Build Go binary
.PHONY: build
build:
	go build ./cmd/main.go $(GO_BUILD_FLAGS) .

# Run the Go app locally
.PHONY: run
run:
	go run ./cmd/main.go

# Clean up binary
.PHONY: clean
clean:
	go clean
	rm -f $(APP_NAME)

# Docker: build image
.PHONY: docker-build
docker-build:
	docker build -t $(DOCKER_IMAGE) .

# Docker: run container
.PHONY: docker-run
docker-run:
	docker run -p $(PORT):$(PORT) $(DOCKER_IMAGE)

# Docker: remove image
.PHONY: docker-clean
docker-clean:
	docker rmi $(DOCKER_IMAGE)
