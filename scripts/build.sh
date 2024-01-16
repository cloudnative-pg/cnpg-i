#!/usr/bin/env bash

cd "$(dirname "$0")/.." || exit

# Recompile protobuf specification
protoc --go_out=. --go_opt=module=github.com/cloudnative-pg/cnpg-i \
    --go-grpc_out=. --go-grpc_opt=module=github.com/cloudnative-pg/cnpg-i \
    proto/wal.proto
protoc --go_out=. --go_opt=module=github.com/cloudnative-pg/cnpg-i \
    --go-grpc_out=. --go-grpc_opt=module=github.com/cloudnative-pg/cnpg-i \
    proto/identity.proto

# Compile client
go build `go list ./... | grep -v 'vendor'`
