GO = go
GO_BUILD = $(GO) build
TARGET_DIR = build
LD_FLAGS := -s -w

.PHONY: all build-osx

all: build-osx

build-osx:
	GOOS=darwin GOARCH=amd64 $(GO_BUILD) -ldflags "$(LD_FLAGS)" -o $(TARGET_DIR)/ff .
