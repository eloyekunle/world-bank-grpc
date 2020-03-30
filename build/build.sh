#!/bin/sh

set -o errexit
set -o nounset

if [ -z "${VERSION:-}" ]; then
    echo "VERSION must be set"
    exit 1
fi

export CGO_ENABLED=0
export GO111MODULE=on

go build 														                           \
		-ldflags "-X $(go list -m)/pkg/version.VERSION=${VERSION}" \
	 	-o "${OUTBIN}" github.com/eloyekunle/world-bank-grpc/cmd
