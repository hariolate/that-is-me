#!/usr/bin/env bash

project_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/../../../" >/dev/null 2>&1 && pwd)"

mkdir -p "${project_root}"/static

#"$(dirname "${BASH_SOURCE[0]}")"/gen_protocol.sh

npm --prefix ${project_root} install
npm --prefix ${project_root} run build