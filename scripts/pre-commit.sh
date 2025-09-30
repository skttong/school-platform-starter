#!/usr/bin/env bash
set -euo pipefail

echo "[pre-commit] go fmt"
go fmt ./...

echo "[pre-commit] go vet"
go vet ./...

echo "[pre-commit] golangci-lint"
if ! command -v golangci-lint >/dev/null 2>&1; then
  echo "golangci-lint not found. Install via: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"
  exit 1
fi
golangci-lint run ./...

echo "[pre-commit] go test"
go test ./... -count=1

echo "[pre-commit] gosec (advisory)"
if command -v gosec >/dev/null 2>&1; then
  gosec ./... || true
else
  echo "gosec not found (skipping). Install via: go install github.com/securego/gosec/v2/cmd/gosec@latest"
fi

echo "[pre-commit] OK"
