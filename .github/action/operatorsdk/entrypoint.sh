#!/bin/sh

# Put repo in GOPATH, so go tools work properly.
PROJECT_ROOT="/go/src/github.com/${GITHUB_REPOSITORY}"
mkdir -p "$PROJECT_ROOT"
rmdir "$PROJECT_ROOT"
ln -s "$GITHUB_WORKSPACE" "$PROJECT_ROOT"
cd "$PROJECT_ROOT" || exit 1

# Execute

make build "$@":"$VERSION"