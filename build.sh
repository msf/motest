#!/bin/bash
set -ex
go generate ./...
go vet ./...
go build ./...
go test ./...
find cmd -type f -exec go build {} \;
