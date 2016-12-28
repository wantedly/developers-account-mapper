NAME := developers-account-mapper
VERSION := 0.1.0
REVISION := $(shell git rev-parse --short HEAD)

SRCS           := $(shell find . -type f -name '*.go')
SRCS_NO_VENDOR := $(shell find . -type f -name '*.go' -not -path "./vendor/*")

LDFLAGS := -ldflags="-s -w -X \"main.Version=$(VERSION)\" -X \"main.Revision=$(REVISION)\""

.DEFAULT_GOAL := bin/$(NAME)

bin/$(NAME): $(SRCS)
	go build $(LDFLAGS) -o bin/$(NAME)

.PHONY: install
install: deps
	go install $(LDFLAGS)

.PHONY: deps
deps: glide
	glide install

.PHONY: glide
glide:
ifeq ($(shell command -v glide 2> /dev/null),)
	curl https://glide.sh/get | sh
endif

.PHONY: gofmt
gofmt: $(SRCS_NO_VENDOR)
	gofmt -s -w $(SRCS_NO_VENDOR)

.PHONY: clean
clean:
	rm -rf bin/*
	rm -rf dist/*
	rm -rf vendor/*
