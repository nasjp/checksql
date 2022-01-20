.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


.PHONY: migrate
migrate: ## migrate db
	docker run --rm -v $$PWD/db/migrations:/migrations --network host migrate/migrate:latest -path=/migrations/ -database 'mysql://root:@tcp(localhost:4529)/app' up

.PHONY: clear-db
clear-db: ## clear db
	docker compose exec db bash -c "mysql -D app -e 'SHOW TABLES' | xargs -I TBL mysql -D app -e 'DROP TABLE IF EXISTS TBL'"
