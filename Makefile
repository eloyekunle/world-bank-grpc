# The binary to build (just the basename).
BIN := world-bank-grpc

OUTBIN = bin/$(BIN)

# Where to push the docker image.
REGISTRY ?= playmice

VERSION := 1.0.0

SRC_DIRS := cmd pkg # directories which hold app source (not vendored)

IMAGE := $(REGISTRY)/$(BIN)

BUILD_IMAGE ?= golang:1.12-alpine

all: build

# Directories that we need created to build/test.
BUILD_DIRS := bin/     \
              .go/bin  \
              .go/cache

# This will build the binary under ./.go and update the real binary iff needed.
.PHONY: build
build: $(BUILD_DIRS)
	@echo "making $(OUTBIN)"
	@docker run                                                 \
	    -i                                                      \
	    --rm                                                    \
	    -u $$(id -u):$$(id -g)                                  \
	    -v $$(pwd):/src                                         \
	    -w /src                                                 \
	    -v $$(pwd)/.go/bin:/go/bin								\
	    -v $$(pwd)/.go/cache:/.cache                            \
	    --env HTTP_PROXY=$(HTTP_PROXY)                          \
	    --env HTTPS_PROXY=$(HTTPS_PROXY)                        \
	    $(BUILD_IMAGE)                                          \
	    /bin/sh -c "                                            \
	        VERSION=$(VERSION)                                  \
	        OUTBIN=$(OUTBIN)			                        \
	        ./build/build.sh                                    \
	    "
	@if ! cmp -s .go/$(OUTBIN) $(OUTBIN); then \
	    mv .go/$(OUTBIN) $(OUTBIN);            \
	    date >$@;                              \
	fi

container: .container say_container_name
.container: build
	@docker build --build-arg BIN=$(BIN) -t $(IMAGE):$(VERSION) .

say_container_name:
	@echo "container: $(IMAGE):$(VERSION)"

push: .push say_push_name
.push: .container
	@docker push $(IMAGE):$(VERSION)

say_push_name:
	@echo "pushed: $(IMAGE):$(VERSION)"

version:
	@echo $(VERSION)

test: $(BUILD_DIRS)
	@docker run                                                 \
	    -i                                                      \
	    --rm                                                    \
	    -u $$(id -u):$$(id -g)                                  \
	    -v $$(pwd):/src                                         \
	    -w /src                                                 \
	    -v $$(pwd)/.go/bin/$(OS)_$(ARCH):/go/bin                \
	    -v $$(pwd)/.go/bin/$(OS)_$(ARCH):/go/bin/$(OS)_$(ARCH)  \
	    -v $$(pwd)/.go/cache:/.cache                            \
	    --env HTTP_PROXY=$(HTTP_PROXY)                          \
	    --env HTTPS_PROXY=$(HTTPS_PROXY)                        \
	    $(BUILD_IMAGE)                                          \
	    /bin/sh -c "                                            \
	        VERSION=$(VERSION)                                  \
	        ./build/test.sh $(SRC_DIRS)                         \
	    "

$(BUILD_DIRS):
	@mkdir -p $@

clean:
	rm -rf .go bin
