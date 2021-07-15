PROJECT_ROOT = $(shell pwd)
BIN_DIR ?= $(PROJECT_ROOT)/bin

build: clean
	@echo "Building..."
	@go build -o "$(BIN_DIR)/ocp-calendar-api" "$(PROJECT_ROOT)/cmd/ocp-calendar-api/main.go"

run:
	@go run "$(PROJECT_ROOT)/cmd/ocp-calendar-api/main.go"

clean:
	@echo "Cleaning..."
	@rm -rf "$(BIN_DIR)/*"
