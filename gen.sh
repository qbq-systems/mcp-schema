#!/bin/sh
# Generates Go types from the MCP JSON Schema with quicktype.
# Runs inside the Docker container (see Makefile: make gen).
#
# SCHEMA_DIR - schemas laid out as <SCHEMA_DIR>/<YYYY-MM-DD|draft>/schema.json
# OUT_DIR    - where to write the v<YYYYMMDD>/ and vdraft/ packages

set -eu

SCHEMA_DIR="${SCHEMA_DIR:-upstream/schema}"
OUT_DIR="${OUT_DIR:-.}"

gen() {
  dir="${1%/}"
  file="$dir/schema.json"
  echo "Dir: $dir"
  [ -f "$file" ] || return 0

  version="v$(basename "$dir" | tr -d '-')"
  out="$OUT_DIR/$version"
  echo "Version: $version"

  rm -rf "$out"
  mkdir -p "$out"
  quicktype \
    -l go \
    -o "$out/$version.go" \
    --package "$version" \
    --telemetry disable \
    --multi-file-output \
    --omit-empty \
    "$file"
}

for dir in "$SCHEMA_DIR"/*/; do
  [ -d "$dir" ] || continue
  gen "$dir"
done
