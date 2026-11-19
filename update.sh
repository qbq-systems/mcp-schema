#!/bin/bash

set -euo pipefail

REPO_DIR="./upstream"
IMAGE="qbq-mcp-schema-gen"

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

cd ../
docker build -q -t "$IMAGE" . >/dev/null
docker run --rm \
  --network none \
  --user "$(id -u):$(id -g)" \
  -e HOME=/tmp \
  -v "$PWD:/work" \
  "$IMAGE"

tag="v$(date +%y).$(date +%-m).$(date +%-d)+$(date +%H%M%S)"
echo "Tag: $tag"

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
