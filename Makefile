.PHONY: all tidy local-build

all: tidy local-builds local-move

tidy:
	go mod tidy

local: local-build local-move

local-build:
	go install github.com/goreleaser/goreleaser@latest
	goreleaser build --snapshot --clean

local-move:
	cp dist/ff_darwin_amd64_v1/ff /usr/local/bin/ff
