# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get

BIN_DIR=bin
BINARY_NAME=moviecall
MAIN_PACKAGE=./cmd/moviecall.go

# Default target (build the project)
default: build

# Build the project
build:
	$(GOBUILD) -o $(BIN_DIR)/$(BINARY_NAME) $(MAIN_PACKAGE)

# Clean the project
clean:
	$(GOCLEAN)
	rm -f $(BIN_DIR)/$(BINARY_NAME)
