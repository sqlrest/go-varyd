#!/usr/bin/env bash
# Regenerate protobuf Go from api/proto into src/proto via buf, run ad hoc —
# never a go.mod tool dependency, which would re-add the buf→docker/quic-go
# CVE graph to go.sum. Requires buf, protoc-gen-go, and protoc-gen-go-grpc on
# PATH (or in GOBIN, which is appended below).
set -o errexit -o nounset -o pipefail
CDPATH='' cd "$(dirname "${BASH_SOURCE[0]}")/.." >/dev/null

PATH="${PATH}:$(go env GOBIN):$(go env GOPATH)/bin" buf generate
