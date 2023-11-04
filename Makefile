#!/usr/bin/make -f

NAME = gotabit
APPNAME = gotabitd
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
COMMIT := $(shell git log -1 --format='%H')
DOCKER := $(shell which docker)
LEDGER_ENABLED ?= true
BINDIR ?= $(GOPATH)/bin
BUILDDIR ?= $(CURDIR)/build
DOCKER := $(shell which docker)
GOPATH ?= $(shell $(GO) env GOPATH)

# don't override user values
ifeq (,$(VERSION))
	VERSION := $(shell git describe --tags)
	# if VERSION is empty, then populate it with branch's name and raw commit hash
	ifeq (,$(VERSION))
		VERSION := $(BRANCH)-$(COMMIT)
	endif
endif

TM_VERSION := $(shell go list -m github.com/cometbft/cometbft | sed 's:.* ::')

export GO111MODULE = on

# process build tags

build_tags = netgo
ifeq ($(LEDGER_ENABLED),true)
  ifeq ($(OS),Windows_NT)
    GCCEXE = $(shell where gcc.exe 2> NUL)
    ifeq ($(GCCEXE),)
      $(error gcc.exe not installed for ledger support, please install or set LEDGER_ENABLED=false)
    else
      build_tags += ledger
    endif
  else
    UNAME_S = $(shell uname -s)
    ifeq ($(UNAME_S),OpenBSD)
      $(warning OpenBSD detected, disabling ledger support (https://github.com/cosmos/cosmos-sdk/issues/1988))
    else
      GCC = $(shell command -v gcc 2> /dev/null)
      ifeq ($(GCC),)
        $(error gcc not installed for ledger support, please install or set LEDGER_ENABLED=false)
      else
        build_tags += ledger
      endif
    endif
  endif
endif

ifeq ($(WITH_CLEVELDB),yes)
  build_tags += gcc
endif
build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

whitespace :=
empty = $(whitespace) $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(empty),$(comma),$(build_tags))

# process linker flags

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=$(NAME) \
	-X github.com/cosmos/cosmos-sdk/version.AppName=$(APPNAME) \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
	-X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)" \
	-X github.com/cometbft/cometbft/version.TMCoreSemVer=$(TM_VERSION)

ifeq ($(WITH_CLEVELDB),yes)
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=cleveldb
endif
ifeq ($(LINK_STATICALLY),true)
	ldflags += -linkmode=external -extldflags "-Wl,-z,muldefs -static"
endif
ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

BUILD_FLAGS := -tags "$(build_tags_comma_sep)" -ldflags '$(ldflags)' -trimpath


###############################################################################
###                                Building                                 ###
###############################################################################

all: install

build: go.sum
ifeq ($(OS),Windows_NT)
	exit 1
else
	go build -mod=readonly $(BUILD_FLAGS) -o build/$(APPNAME) ./cmd/$(APPNAME)
endif

go.sum:
	go mod tidy

build-linux: build-linux-amd64 build-linux-arm64

build-linux-amd64: go.sum
	mkdir -p $(BUILDDIR)
	$(DOCKER) buildx create --name $(NAME)-builder || true
	$(DOCKER) buildx use $(NAME)-builder
	$(DOCKER) buildx build \
		--build-arg GO_VERSION=$(shell cat go.mod | grep -E 'go [0-9].[0-9]+' | cut -d ' ' -f 2) \
		--platform linux/amd64 \
		-t  $(NAME)-dev-amd64 --rm \
		--load \
		-f Dockerfile .
	$(DOCKER) rm -f $(NAME)-temp-amd64 || true
	$(DOCKER) create -ti --name $(NAME)-temp-amd64 $(NAME)-dev-amd64
	$(DOCKER) cp $(NAME)-temp-amd64:/usr/bin/$(APPNAME) $(BUILDDIR)/$(APPNAME)-linux-amd64
	$(DOCKER) rm -f $(NAME)-temp-amd64

build-linux-arm64: go.sum $(BUILDDIR)/
	mkdir -p $(BUILDDIR)
	$(DOCKER) buildx create --name $(NAME)-builder || true
	$(DOCKER) buildx use $(NAME)-builder
	$(DOCKER) buildx build \
		--build-arg GO_VERSION=$(shell cat go.mod | grep -E 'go [0-9].[0-9]+' | cut -d ' ' -f 2) \
		--platform linux/arm64 \
		-t  $(NAME)-dev-arm64 --rm \
		--load \
		--build-arg arch=aarch64 \
		-f Dockerfile .
	$(DOCKER) rm -f $(NAME)-temp-arm64 || true
	$(DOCKER) create -ti --name $(NAME)-temp-arm64 $(NAME)-dev-arm64
	$(DOCKER) cp $(NAME)-temp-arm64:/usr/bin/$(APPNAME) $(BUILDDIR)/$(APPNAME)-linux-arm64
	$(DOCKER) rm -f $(NAME)-temp-arm64

install: go.sum
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/$(APPNAME)

start: build
	./scripts/local-test.sh

clean:
	rm -rf $(BUILDDIR)/

clean-local:
	rm -rf $(HOME)/.$(NAME)

.PHONY: build build-linux install clean start

###############################################################################
###                                  Proto                                  ###
###############################################################################

proto-all: proto docs

proto: proto-tools
	@echo "Generate Protobuf"
	./scripts/protoc-gen.sh

docs: proto-tools
	@echo "Generate Protobuf swagger files"
	./scripts/protoc-swagger-gen.sh

proto-tools:
	@echo "Install Protobuf tools"
	./scripts/protoc-tools.sh

.PHONY: proto docs proto-tools

###############################################################################
###                                Linting                                  ###
###############################################################################

lint:
ifeq (,$(shell which golangci-lint))
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.2
endif
	golangci-lint run --out-format=tab --timeout=10m

format:
ifeq (,$(shell which goimports))
	go install golang.org/x/tools/cmd/goimports@latest
endif
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./client/docs/*" -not -path "./tests/mocks/*" -not -name '*.pb.go' | xargs gofmt -w -s
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./client/docs/statik/statik.go" -not -path "./tests/mocks/*" -not -name '*.pb.go' | xargs goimports -w -local github.com/cosmos/cosmos-sdk

lint-fix:
ifeq (,$(shell which golangci-lint))
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.2
endif
	golangci-lint run --fix --out-format=tab --issues-exit-code=0

.PHONY: lint format lint-fix
