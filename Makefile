# The binary to build (just the basename).
BIN := world-bank-grpc

OUTBIN = bin/$(BIN)

# Where to push the docker image.
REGISTRY ?= playmice

VERSION := 1.0.0

SRC_DIRS := cmd pkg # directories which hold app source (not vendored)

IMAGE := $(REGISTRY)/$(BIN)

all: build

# Directories that we need created to build/test.
BUILD_DIRS := bin/

# This will build the binary under ./.go and update the real binary iff needed.
.PHONY: build
build: $(BUILD_DIRS)
	@echo "making $(OUTBIN)"
	@VERSION=$(VERSION)				\
		OUTBIN=$(OUTBIN) 			\
		./build/build.sh

.PHONY: install
install:
	@echo "installing $(BIN)"
	@VERSION=$(VERSION) ./build/install.sh

container: .container say_container_name
.container: build
	docker build --build-arg BIN=$(BIN) -t $(IMAGE):$(VERSION) .

say_container_name:
	@echo "container: $(IMAGE):$(VERSION)"

push: .push say_push_name
.push: .container
	echo "$(DOCKER_PASSWORD)" | docker login -u "$(DOCKER_USERNAME)" --password-stdin
	docker push $(IMAGE):$(VERSION)

say_push_name:
	@echo "pushed: $(IMAGE):$(VERSION)"

version:
	@echo $(VERSION)

test: $(BUILD_DIRS)
	@VERSION=$(VERSION)				\
		OUTBIN=$(OUTBIN) 			\
		./build/test.sh $(SRC_DIRS)

lint:
	@hash golangci-lint > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin v1.24.0; \
	fi
	GO111MODULE=on golangci-lint run

$(BUILD_DIRS):
	@mkdir -p $@

generate:
	@ if ! which protoc > /dev/null; then \
		echo "error: protoc not installed" >&2; \
		exit 1; \
	fi
	go generate github.com/eloyekunle/world-bank-grpc/...

clean:
	rm -rf .go bin
