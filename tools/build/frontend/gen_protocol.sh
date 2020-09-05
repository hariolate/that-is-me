#!/usr/bin/env bash

set -e

project_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/../../../" >/dev/null 2>&1 && pwd)"

#protoc -I="${project_root}"/src/protocol --js_out=import_style=commonjs:"${project_root}"/src/protocol "${project_root}/src/protocol/"*.proto
#
## workaround "error  'proto' is not defined         no-undef"
#for f in "${project_root}"/src/protocol/*.js
#do
#    echo '/* eslint-disable */' | cat - "${f}" > temp && mv temp "${f}"
#done

cd ${project_root}
npx pbjs -t json-module -w commonjs -o "${project_root}"/src/protocol/protocol.js "${project_root}"/src/protocol/*.proto