GO ?= go
GITREF = $(shell git describe --always --long --tags --dirty)
GOBUILD_FLAGS = -ldflags "-X github.com/AndreaJegher/chromium-netlog-parser/pkg/build.gitCommitID=$(GITREF)"

GOINSTALL = $(GO) install
GOBUILD = $(GO) build

.PHONY: install

all: build install

install: install-nlp

build: build-nlp

install-nlp:
	cd cmd/nlp; $(GOINSTALL) $(GOBUILD_FLAGS)

build-nlp:
	$(GOBUILD) $(GOBUILD_FLAGS) -o ./bin/nlp ./cmd/nlp
