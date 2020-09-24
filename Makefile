## Folder content generated files
BUILD_FOLDER = ./build
## command
GO           = go
GO_VENDOR    = go mod
MKDIR_P      = mkdir -p

## Random Alphanumeric String
SECRET_KEY   = $(shell cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 32 | head -n 1)

## UNAME
UNAME := $(shell uname)

################################################

.PHONY: all
all: build

.PHONY: pre-build
pre-build:
	$(MAKE) download

.PHONY: build
build: pre-build
	$(MAKE) src.build

.PHONY: clean
clean:
	$(RM) -rf $(BUILD_FOLDER)

## vendor/ #####################################

.PHONY: download
download:
	$(GO_VENDOR) vendor

## src/ ########################################

.PHONY: src.build
src.build:
	$(MKDIR_P) $(BUILD_FOLDER)/pkg/cmd/sample-container-runtime/
	GO111MODULE=on $(GO) build -mod=vendor -v -o $(BUILD_FOLDER)/pkg/cmd/sample-container-runtime/sample-container-runtime \
	./cmd/container-runtime
