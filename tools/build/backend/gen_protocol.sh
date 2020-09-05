#!/usr/bin/env bash

set -e

project_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/../../../" >/dev/null 2>&1 && pwd)"

protoc -I="${project_root}"/src/protocol --go_out="${project_root}"/src/protocol --descriptor_set_out="${project_root}"/src/protocol/descriptor.pb "${project_root}/src/protocol/"*.proto
