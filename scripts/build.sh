#!/usr/bin/env bash
set -euo pipefail

CC=musl-gcc go build --ldflags '-linkmode external -extldflags "-static"' "$@" cmd/bibak/main.go
