# Makefile for simple golang Project

# Variables
APP_NAME := simple-bank
GO := go
# GOCMD := $(GO) build
GOBUILD_FLAGS := -ldflags="-s -w"
# GOFILES := $(wildcard *.go)
GOFILES := cmd/main.go
DIST_DIR := ./dist
step ?= 1
OAUTH_CLI := cmd/oauth/main.go

.PHONY: build clean run migration-create migration-up migration-down generate-rsa docker-up

# Build the project
build: clean
	@echo "Building $(APP_NAME)"
	$(GO) build $(GOBUILD_FLAGS) -o $(DIST_DIR)/$(APP_NAME) $(GOFILES)

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts"
	@rm -rf $(DIST_DIR)

# Run the built binary
run: build
	$(DIST_DIR)/$(APP_NAME)

migration-create: build
	$(DIST_DIR)/$(APP_NAME) migration-create -migration_name=$(name)
	
# migrate create -ext sql -dir pkg/common/db/migration -seq $(name)

migration-up: build
	$(DIST_DIR)/$(APP_NAME) migration-up -migrate_step=$(step)

migration-down: build
	$(DIST_DIR)/$(APP_NAME) migration-down -migrate_step=$(step)

# generate private/public key value pair
generate-rsa: build
	$(DIST_DIR)/$(APP_NAME) generate-rsa 

docker-up:
	docker-compose up -d	

# migrate -path pkg/common/db/migration -database "postgresql://pgsuperuser:Admin@1@localhost:5432/gotodo?sslmode=disable" -verbose up 	

# Help target to display available targets
help:
	@echo "Available targets:"
	@echo "  build        : Build the project"
	@echo "  clean        : Clean build artifacts"
	@echo "  run          : Build and run the project"
	@echo "  help         : Display this help message"
