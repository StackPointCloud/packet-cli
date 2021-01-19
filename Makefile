EXECUTABLES = git go find pwd
K := $(foreach exec,$(EXECUTABLES),\
        $(if $(shell which $(exec)),some string,$(error "No $(exec) in PATH)))

ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

PACKAGE_NAME?=github.com/packethost/packet-cli
BINARY?=equinix
GIT_VERSION?=$(shell git log -1 --format="%h")
VERSION?=$(GIT_VERSION)
RELEASE_TAG?=$(shell git tag --points-at HEAD)
ifneq (,$(RELEASE_TAG))
VERSION:=$(RELEASE_TAG)-$(VERSION)
endif

BUILD=`git rev-parse --short HEAD`
PLATFORMS?=darwin linux windows freebsd
ARCHITECTURES?=amd64 arm64
GOBIN?=$(shell go env GOPATH)/bin
LINTER?=$(GOBIN)/golangci-lint
FORMATTER?=$(GOBIN)/goimports


# Setup linker flags option for build that interoperate with variable names in src code
LDFLAGS?=-ldflags "-X $(PACKAGE_NAME)/cmd.Version=$(VERSION) -X $(PACKAGE_NAME)/cmd.Build=$(BUILD)"

.PHONY: default fmt fmt-check lint test vet golint tag version
default: lint generate-docs

## fmt files
fmt: $(FORMATTER)
	$(FORMATTER) -lw .
$(FORMATTER):
	go get golang.org/x/tools/cmd/goimports

lint: $(LINTER)
	$(LINTER) run ./...
$(LINTER):
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.30.0

build:
	go build ${LDFLAGS} -o bin/${BINARY}

build-all:
	$(foreach GOOS, $(PLATFORMS),\
	$(foreach GOARCH, $(ARCHITECTURES), $(shell export GOOS=$(GOOS); export GOARCH=$(GOARCH); go build -v -o bin/$(BINARY)-$(GOOS)-$(GOARCH))))

clean:
	rm -rf bin/
	find ${ROOT_DIR} -name '${BINARY}[-?][a-zA-Z0-9]*[-?][a-zA-Z0-9]*' -delete

clean-docs:
	rm -rf docs/

install:
	go install ${LDFLAGS}
	mv ${GOPATH}/bin/packet-cli ${GOPATH}/bin/packet

generate-docs: clean-docs
	mkdir -p docs
	go run main.go docs ./docs

test:
	go test ./tests
