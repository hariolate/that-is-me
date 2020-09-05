#!/usr/bin/env bash

project_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/../" >/dev/null 2>&1 && pwd)"

${project_root}/tools/build/backend/build.sh
${project_root}/tools/build/frontend/build.sh