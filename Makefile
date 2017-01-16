NAME           := developers-account-mapper
VERSION        := v0.2.0
REVISION       := $(shell git rev-parse --short HEAD)

SRCS           := $(shell find . -type f -name '*.go')
SRCS_NO_VENDOR := $(shell find . -type f -name '*.go' -not -path "./vendor/*")

LDFLAGS        := -ldflags="-s -w -X \"main.Version=$(VERSION)\" -X \"main.Revision=$(REVISION)\""

DIST_DIRS      := find * -type d -exec

.DEFAULT_GOAL  := bin/$(NAME)

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

.PHONY: test
test:
	go test -v `glide novendor`

.PHONY: gofmt
gofmt: $(SRCS_NO_VENDOR)
	gofmt -s -w $(SRCS_NO_VENDOR)

.PHONY: dist
dist:
	cd dist && \
	$(DIST_DIRS) cp ../LICENSE {} \; && \
	$(DIST_DIRS) cp ../README.md {} \; && \
	$(DIST_DIRS) tar -zcf $(NAME)-$(VERSION)-{}.tar.gz {} \; && \
	$(DIST_DIRS) zip -r $(NAME)-$(VERSION)-{}.zip {} \; && \
	cd ..

.PHONY: cross-build
cross-build:
	for os in darwin linux windows; do \
		for arch in amd64 386; do \
			GOOS=$$os GOARCH=$$arch go build -a -tags netgo -installsuffix netgo $(LDFLAGS) -o dist/$$os-$$arch/$(NAME); \
		done; \
	done

.PHONY: release
release:
	git checkout master
	git tag $(VERSION)
	git push origin $(VERSION)

.PHONY: clean
clean:
	rm -rf bin/*
	rm -rf dist/*
	rm -rf vendor/*
