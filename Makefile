# ---------------------------------------------------------
# Set default values.
# ---------------------------------------------------------

.DEFAULT_GOAL := build
VERSION := $(shell date +%Y-%m-%d-%H%M%S)

# ---------------------------------------------------------
# Build the artifact.
# ---------------------------------------------------------

.PHONY: build
.SILENT: build

build: 
	go install -ldflags="-s -w -X 'github.com/deathlabs/emu/v4/cmd.version=$(VERSION)'" . 

# ---------------------------------------------------------
# Update dependencies.
# ---------------------------------------------------------

.PHONY: update
.SILENT: update

update:
	go get -u ./...
	go mod tidy
