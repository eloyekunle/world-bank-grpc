#!/bin/sh

set -o errexit
set -o nounset
set -o pipefail

export CGO_ENABLED=0
export GO111MODULE=on

TARGETS=$(for d in "$@"; do echo ./$d/...; done)

echo "Running tests:"
go test -coverprofile=coverage.txt -covermode=atomic -installsuffix "static" ${TARGETS}
echo
