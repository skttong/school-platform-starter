.PHONY: run dev tidy migrate-up seed

run:
	APP_PORT?=8080
	go run ./cmd/server

dev:
	air || go run ./cmd/server

tidy:
	go mod tidy

migrate-up:
	psql "$$DATABASE_URL" -f migrations/0001_init.sql

seed:
	psql "$$DATABASE_URL" -f migrations/0002_seed_rbac.sql


fmt:
	go fmt ./...

vet:
	go vet ./...

lint:
	golangci-lint run ./...

sec:
	gosec ./... || true

test:
	go test ./... -v

ci:
	go mod tidy
	go build ./...
	$(MAKE) test
	$(MAKE) lint
	$(MAKE) sec

hook-install:
	@cp scripts/pre-commit.sh .git/hooks/pre-commit || true
	@chmod +x .git/hooks/pre-commit || true
	@echo "Installed pre-commit hook."
