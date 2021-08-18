.PHONY: build image release proto

VERSION_TAG = 1.0.0
MILESTONE_TAG = 08.2021
REGISTRY = docker.io/rubinus

# auto-generated
COMMIT_ID := $(shell git rev-parse HEAD)
BUILD_TS := $(shell date +'%Y%m%d%H%M%S')
BRANCH_TAG := $(shell git rev-parse --abbrev-ref HEAD)
GIT_TAG := $(shell git describe --tags --exact-match 2> /dev/null \
                     || git symbolic-ref -q --short HEAD \
                     || git rev-parse --short HEAD)
VERSION := $(VERSION_TAG)-build-$(BRANCH_TAG)-$(BUILD_TS)
BUILD := $(VERSION_TAG)-$(MILESTONE_TAG)-build-$(BUILD_TS)-$(BRANCH_TAG)-$(COMMIT_ID)
PROJECT_ROOT := $(shell pwd -L)
OUTPUT := $(PROJECT_ROOT)/_output

default: image

build: compile 
# vet lint compile

compile:
ifeq ($(GOOS),$(GOHOSTOS))
	@echo "... build with race condition detector on ..."
	@go build -race -p 8  -o $(OUTPUT)/origin
else
	@echo "... race condition detector is disabled on cross compiling"
	@go build -p 8 -o $(OUTPUT)/origin
endif

image:
	@echo "Build image of version $(VERSION)"
	@echo "Build origin"
	@$(MAKE) GOOS=linux GOARCH=amd64 CGO_ENABLED=0 build
	@echo "Build image $(REGISTRY)/origin:$(VERSION)"
	@docker build -f $(PROJECT_ROOT)/Dockerfile --build-arg BUILD=$(BUILD) -t $(REGISTRY)/origin:$(VERSION) . >/dev/null
	@echo "Done"

release:
	@echo "Build image of version $(GIT_TAG)"
	@echo "Build origin"
	@$(MAKE) GOOS=linux GOARCH=amd64 CGO_ENABLED=0 build
	@echo "Build image $(REGISTRY)/origin:$(GIT_TAG)"
	@docker build -f $(PROJECT_ROOT)/Dockerfile --build-arg BUILD=$(BUILD) -t $(REGISTRY)/origin:$(GIT_TAG) . >/dev/null
	@echo "Done"

proto:
	./compile_proto.sh
