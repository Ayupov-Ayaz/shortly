# load environments
include .env.development
export

run-development:
	go run ./cmd/api/main.go

run-production:
	go run ./cmd/api/main.go -env=production

generate-api:
	@echo "Generating API code from OpenAPI spec ..."
	go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen \
		-config ./api/config.yaml \
		./api/openapi.yaml
	@echo "✅ API code generated"


TOOLS_DIR := $(PWD)/tools
# migration

DB_URL := postgres://$(SHORTLY_POSTGRES_USER):$(SHORTLY_POSTGRES_PASSWORD)@$(SHORTLY_POSTGRES_HOST):$(SHORTLY_POSTGRES_PORT)/$(SHORTLY_POSTGRES_DB)?sslmode=$(SHORTLY_POSTGRES_SSLMODE)

# colors for stdout
GREEN := \033[0;32m
YELLOW := \033[1;33m
NC := \033[0m # No Color

.PHONY: migrate-create migrate-up migrate-down migrate-down-all migrate-force migrate-version

# create new migration
migrate-create:
	@echo "$(YELLOW)📝 Creating new migration...$(NC)"
	@read -p "Enter migration name: " name; \
	cd $(TOOLS_DIR) && go run github.com/golang-migrate/migrate/v4/cmd/migrate@v4.18.2 \
		create -ext sql -dir ../migrations -seq $${name}
	@echo "$(GREEN)✅ Migration created successfully$(NC)"

# use migration
migrate-up:
	@echo "⬆️  Applying migrations..."
	cd tools && go run \
		-tags 'postgres' \
		github.com/golang-migrate/migrate/v4/cmd/migrate \
		-database "$(DB_URL)" \
		-path ../migrations \
		-verbose up

migrate-down:
	@echo "⬆️  Applying migrations..."
	cd tools && go run \
		-tags 'postgres' \
		github.com/golang-migrate/migrate/v4/cmd/migrate \
		-database "$(DB_URL)" \
		-path ../migrations \
		-verbose down

db-reset:
	@echo "$(YELLOW)🔄 Resetting database...$(NC)"
	docker exec -it shortly-postgres-1 psql -U shortly -d shortly -c "DROP SCHEMA public CASCADE; CREATE SCHEMA public;"
	@echo "$(GREEN)✅ Database reset complete$(NC)"
