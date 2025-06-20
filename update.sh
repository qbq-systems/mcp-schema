#!/bin/bash

set -euo pipefail

REPO_DIR="./upstream"

cd "$REPO_DIR" || { echo "Repo dir not found"; exit 1; }
LOCAL_HASH=$(git rev-parse HEAD)
REMOTE_HASH=$(git ls-remote origin HEAD | awk '{print $1}')

echo "Local hash:  $LOCAL_HASH"
echo "Remote hash: $REMOTE_HASH"

if [ "$LOCAL_HASH" == "$REMOTE_HASH" ]; then
  echo "No changes"
  exit 0
fi

echo "Changes detected, pulling..."
git pull --ff-only

function gen() {
  dir="$1"
  echo "Dir: $dir"
  file="$dir/schema.json"
  if [ ! -f "$file" ]; then
    return 1
  fi
  date_part=$(basename "$dir" | tr -d '-')
  version="v${date_part}"
  echo "Version: $version"
  out="../$version"
  rm -rf "$out"
  mkdir "$out"
  quicktype \
    -l go \
    -o "$out/$version.go" \
    --package "$version" \
    --telemetry disable \
    --multi-file-output \
    --omit-empty \
    "$file"
}

for dir in "schema"/*/; do
  [ -d "$dir" ] || continue
  gen "$dir"
done

tag="v$(date +%y).$(date +%-m).$(date +%-d)+$(date +%H%M%S)"
echo "Tag: $tag"

cd ../
if [ -z "$(git status --porcelain)" ]; then
  echo "No changes for commit"
  exit 0
fi

git add .
git commit -m "$tag"
git push github main
git tag "$tag"
git push github "$tag"

echo "Done"
