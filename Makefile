NAME           := developers-account-mapper
VERSION        := v0.2.2
REVISION       := $(shell git rev-parse --short HEAD)

SRCS           := $(shell find . -type f -name '*.go')
SRCS_NO_VENDOR := $(shell find . -type f -name '*.go' -not -path "./vendor/*")
NOVENDOR       := $(shell go list ./... | grep -v vendor)

LDFLAGS        := -ldflags="-s -w -X \"main.Version=$(VERSION)\" -X \"main.GitCommit=$(REVISION)\""

DIST_DIRS      := find * -type d -exec

DOCKER_REPOSITORY := quay.io
DOCKER_IMAGE_NAME := $(DOCKER_REPOSITORY)/wantedly/developers-account-mapper
DOCKER_IMAGE_TAG ?= latest
DOCKER_IMAGE := $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)

.DEFAULT_GOAL  := bin/$(NAME)

bin/$(NAME): $(SRCS)
	go build $(LDFLAGS) -o bin/$(NAME)

.PHONY: install
install: deps
	go install $(LDFLAGS)

.PHONY: dep
dep:
ifeq ($(shell command -v dep 2> /dev/null),)
	go get -u github.com/golang/dep/cmd/dep
endif

.PHONY: deps
deps: dep
	dep ensure -v

.PHONY: test
test:
	go test -v $(NOVENDOR)

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

.PHONY: docker-build
docker-build:
ifeq ($(findstring ELF 64-bit LSB,$(shell file bin/$(NAME) 2> /dev/null)),)
	@echo "bin/$(NAME) is not a binary of Linux 64bit binary."
	@exit 1
endif
	docker build -t $(DOCKER_IMAGE) .

.PHONY: ci-docker-release
ci-docker-release: docker-build
	@docker login -u="$(DOCKER_QUAY_USERNAME)" -p="$(DOCKER_QUAY_PASSWORD)" $(DOCKER_REPOSITORY)
	docker push $(DOCKER_IMAGE)

.PHONY: clean
clean:
	rm -rf bin/*
	rm -rf dist/*
	rm -rf vendor/*
