#!/usr/bin/env bash

set -eo pipefail

need_cmd() {
  if ! check_cmd "$1"; then
    if [ "$2" != "" ]; then
      `$2`
    else
      err "need '$1' (command not found)"
    fi
  fi
}

check_cmd() {
  command -v "$1" >/dev/null 2>&1
}

need_cmd buf "go install github.com/bufbuild/buf/cmd/buf@v1.6.0"
need_cmd protoc-gen-go \
  "go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28"
need_cmd protoc-gen-go-grpc \
  "go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2"
need_cmd protoc-gen-grpc-gateway \
  "go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@v1.16.0"
need_cmd protoc-gen-gocosmos \
  "go install github.com/regen-network/cosmos-proto/protoc-gen-gocosmos"

need_cmd statik "go install github.com/rakyll/statik"
need_cmd protoc-gen-swagger  \
  "go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@latest"
  