.PHONY: help
help:
	@echo "Usage: make [target]"
	@echo "Targets:"
	@echo "  install-migrate-tool   Install the migrate tool"
	@echo "  db-up                  Starts docker-compose and run the migrations"
	@echo "  db-down                Rollback the migrations and stop docker-compose"
	@echo "  psql-local             Connect to the local postgres database"

.PHONY: install-migrate-tool
install-migrate-tool:
	@which migrate 1>/dev/null || go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

.PHONY: db-up
db-up: install-migrate-tool
	@docker-compose up -d
	@sleep 2
	@migrate -database "postgresql://user:password@localhost:5432/postgres?sslmode=disable" -path migrations up

.PHONY: db-down
db-down: install-migrate-tool
	@migrate -database "postgresql://user:password@localhost:5432/postgres?sslmode=disable" -path migrations down -all
	@docker-compose down

.PHONY: psql-local
psql-local:
	@docker exec -it reportpipe.postgres psql -d postgres -U user
