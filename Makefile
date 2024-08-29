.PHONY: status up reset

MIGRATE_DIR=migrate
DB_STRING="postgres://eatmore:pg@0.0.0.0:5432/data_base"
GOOSE_CMD=GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(DB_STRING) goose -dir=$(MIGRATE_DIR)

status:
	@$(GOOSE_CMD) status

up:
	@$(GOOSE_CMD) up

reset:
	@$(GOOSE_CMD) reset

