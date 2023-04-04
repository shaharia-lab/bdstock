# Set variables
APP_NAME := bdstock
GOOS := linux
GOARCH := amd64
BUILD_DIR := build
DIST_DIR := dist

# Set flags
LDFLAGS := -w -s
GCFLAGS := -trimpath=$(GOPATH) -installsuffix netgo

# Set commands
GO := go
GOBUILD := CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build
GOCLEAN := $(GO) clean
GOTEST := $(GO) test
GOMOD := $(GO) mod
MKDIR := mkdir -p
CP := cp
RM := rm -rf

.PHONY: all build clean test mod vendor dist

all: clean build test

build: $(BUILD_DIR)/$(APP_NAME)

$(BUILD_DIR)/$(APP_NAME):
	$(MKDIR) $(BUILD_DIR)
	$(GOBUILD) -gcflags="$(GCFLAGS)" -ldflags="$(LDFLAGS)" -o $@

clean:
	$(GOCLEAN)
	$(RM) $(BUILD_DIR)
	$(RM) $(DIST_DIR)

test:
	$(GOTEST) ./...

testc:
	$(GOTEST) ./... -coverprofile=coverage.out

mod:
	$(GOMOD) download

vendor:
	$(GOMOD) vendor

dist: clean build
	$(MKDIR) $(DIST_DIR)
	$(CP) $(BUILD_DIR)/$(APP_NAME) $(DIST_DIR)/$(APP_NAME)
