#!/bin/bash

VERSION="$(awk '/BinaryVersion/{gsub("\"",""); print"v"$3}' defaults/const.go)"
if [ -z "${VERSION}" ]
then
	echo "Unknown version."
	exit 1
fi

echo -n "Checking git tag \"${VERSION}\"  - "
TAG="$(git describe --tags "${VERSION}")"
if [ ! -z "${TAG}" ]
then
	echo "Version ${VERSION} already exists."
	echo "Delete with:"
	echo "git tag -d ${VERSION}"
  echo ""
	echo "OR modify defaults/const.go file."
  echo ""
	echo "OR hit enter to just commit and push."
  echo ""
  echo -n "OK? (Ctrl-C to terminate) "
  read OK

  git add .
  git commit -m "Committed ${VERSION}"
  git push

  exit 1
fi

echo ""
echo "Releasing version ${VERSION}"
echo ""
echo -n "OK? (Ctrl-C to terminate) "
read OK

git add .
git commit -m "Committed ${VERSION}"
git tag "${VERSION}"
git push
git push origin "${VERSION}"

if ! command -v goreleaser &>/dev/null; then
	echo "Error: goreleaser not found. Install with: brew install goreleaser/tap/goreleaser"
	exit 1
fi

# Ensure go is in PATH (needed for before hooks and build) – Cursor/integrated terminals often have minimal PATH
if ! command -v go &>/dev/null; then
	for dir in /opt/homebrew/bin /usr/local/go/bin /usr/local/bin; do
		if [ -x "${dir}/go" ]; then
			export PATH="${dir}:${PATH}"
			break
		fi
	done
fi
if ! command -v go &>/dev/null; then
	echo "Error: go not found. Run this script from a terminal where 'go version' works, or install Go."
	exit 1
fi

goreleaser release --clean

git add .
git commit -m "Released ${VERSION}"
git push

