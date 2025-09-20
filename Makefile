include .env

# Paths
MIGRATIONS_DIR=./cmd/migrate/migrations

# Database URL (from .env)
DB_URL=$(DB_ADDR)

# Goose binary
GOOSE=goose -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)"

# ─────────────────────────────
# Migration commands
# ─────────────────────────────

.PHONY: migrate-up
migrate-up:
	$(GOOSE) up

.PHONY: migrate-down
migrate-down:
	$(GOOSE) down

.PHONY: migrate-status
migrate-status:
	$(GOOSE) status

.PHONY: migrate-create
# Usage: make migrate-create name=create_users_and_posts
migrate-create:
	@read -p "Enter migration name: " name; \
	$(GOOSE) create $$name sql

# ─────────────────────────────
# App commands
# ─────────────────────────────

.PHONY: run
run:
	go run ./cmd/api

.PHONY: build
build:
	go build -o social ./cmd/api

.PHONY: docker-up
docker-up:
	docker compose up --build

.PHONY: docker-down
docker-down:
	docker compose down
