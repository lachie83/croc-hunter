default: docker_build

DOCKER_IMAGE ?= lachlanevenson/croc-hunter
DOCKER_TAG ?= `git rev-parse --abbrev-ref HEAD`
VCS_REF ?= `git rev-parse --short HEAD`

.PHONY: docker_build
docker_build:
	@docker build \
	  --build-arg VCS_REF=$(VCS_REF) \
	  --build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
	  -t $(DOCKER_IMAGE):$(DOCKER_TAG) .

.PHONY: docker_push
docker_push:
	# Push to DockerHub
	docker tag $(DOCKER_IMAGE):$(DOCKER_TAG) $(DOCKER_IMAGE):latest
	docker push $(DOCKER_IMAGE):$(DOCKER_TAG)
	docker push $(DOCKER_IMAGE):latest

# go option
GO        ?= go
PKG       := $(shell glide novendor)
TAGS      :=
TESTS     := .
TESTFLAGS :=
LDFLAGS   :=
GOFLAGS   :=
BINDIR    := $(CURDIR)/bin

.PHONY: all
all: build

.PHONY: build
build:
	GOBIN=$(BINDIR) $(GO) install $(GOFLAGS) -tags '$(TAGS)' -ldflags '$(LDFLAGS)' github.com/lachie83/croc-hunter/...

HAS_GLIDE := $(shell command -v glide;)
HAS_GIT := $(shell command -v git;)

.PHONY: bootstrap
bootstrap:
ifndef HAS_GLIDE
	go get -u github.com/Masterminds/glide
endif
ifndef HAS_GIT
	$(error You must install Git)
endif
	glide install
