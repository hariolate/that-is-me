#!/usr/bin/env bash

set -e

project_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/../../../" >/dev/null 2>&1 && pwd)"

mkdir -p "${project_root}"/bin

"$(dirname "${BASH_SOURCE[0]}")"/gen_protocol.sh

go get -tags netgo -v ./...
go build -tags netgo -v -v -o "${project_root}"/bin/app-$(uname) "${project_root}"/src/cmd/app
