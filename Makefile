NAME := github-username-converter
VERSION := 0.1.0
REVISION := $(shell git rev-parse --short HEAD)

LDFLAGS := -ldflags="-s -w -X \"main.Version=$(VERSION)\" -X \"main.Revision=$(REVISION)\""

.DEFAULT_GOAL := bin/$(NAME)

bin/$(NAME):
	go build $(LDFLAGS) -o bin/$(NAME)

.PHONY: clean
clean:
	rm -rf bin/*
	rm -rf dist/*
	rm -rf vendor/*

.PHONY: install
install:
	go install $(LDFLAGS)
