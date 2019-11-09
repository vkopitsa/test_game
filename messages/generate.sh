#!/bin/bash

set -x

# Path to this plugin
PROTOC_GEN_TS_PATH="../node_modules/.bin/protoc-gen-ts"

# Directory to write generated code to (.js and .d.ts files)
OUT_DIR="../front/src/proto/"

#     --js_out=library=${OUT_DIR}"pb",binary:. \


protoc \
    --plugin="protoc-gen-ts=${PROTOC_GEN_TS_PATH}" \
    --go_out=. \
    --js_out=import_style=commonjs,binary,${OUT_DIR}:. \
    --ts_out=${OUT_DIR} \
    *.proto
