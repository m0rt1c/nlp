GO ?= go
GITREF = $(shell git describe --always --long --tags --dirty)
GOBUILD_FLAGS = -ldflags "-X github.com/AndreaJegher/chrome-netlog-parser/pkg/build.GitCommitID=$(GITREF)"

GOINSTALL = $(GO) install
GOBUILD = $(GO) build

.PHONY: install

all: build install

install: install-nlp

build: build-nlp

install-nlp:
	cd cmd/nlp; $(GOINSTALL)

build-nlp:
	$(GOBUILD) -o ./bin/nlp ./cmd/nlp
