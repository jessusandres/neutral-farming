include .env
export

DATABASE_URL:=postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=$(DB_SSL_MODE)

print-db:
	@echo "DATABASE_URL=$(DATABASE_URL)"

MIGRATE=docker compose run --rm migrate -path /migrations -database $(DATABASE_URL)

migrate-up:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down 1

migrate-create:
	$(MIGRATE) create -ext sql -dir /migrations -seq $(name)

migrate-down-all:
	$(MIGRATE) down