#!/bin/bash
# Wrapper so goreleaser finds Go (e.g. when PATH is minimal in Cursor/integrated terminals)
set -e
for dir in /opt/homebrew/bin /usr/local/go/bin /usr/local/bin "$HOME/go/bin"; do
	[ -x "${dir}/go" ] && export PATH="${dir}:${PATH}" && break
done
if ! command -v go &>/dev/null; then
	echo "Error: go not found. Install Go or run from a terminal where 'go version' works."
	exit 1
fi
exec goreleaser "$@"
