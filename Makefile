.PHONY: all tidy local-build

all: tidy local-builds

tidy:
	go mod tidy

local-build:
	go install github.com/goreleaser/goreleaser@latest
	goreleaser build --snapshot --rm-dist
