GO ?= go
GITREF = $(shell git describe --always --long --tags --dirty)
GOBUILD_FLAGS = -ldflags "-X github.com/AndreaJegher/nlp/pkg/build.gitCommitID=$(GITREF)"

GOINSTALL = $(GO) install
GOBUILD = $(GO) build

.PHONY: install

all: build install

install: install-nlp

build: build-nlp

install-nlp:
	$(GOINSTALL) $(GOBUILD_FLAGS)

build-nlp:
	$(GOBUILD) $(GOBUILD_FLAGS) -o ./bin/nlp
