# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get

# Binary name
BINARY_NAME=moviecall

# Main package directory
MAIN_PACKAGE=./cmd/moviecall.go
DOTENV_PATH=./.env

# Build the project
build:
moviecall:
	export DOTENV_PATH=$(DOTENV_PATH) && $(GOBUILD) -o moviecall $(MAIN_PACKAGE)

# Clean the project
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

