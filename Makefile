.PHONY: build
build:
	docker-compose build

.PHONY: run
run:
	docker-compose up

.PHONY: db-create
db-create:
	sh ./docker/bin/create-db.sh

.PHONY: db-migrate
db-migrate:
	docker-compose run --rm api go run infrastructure/datastore/migrate/migrate.go

.PHONY: test
test:
	docker-compose run --rm api go test -v -cover ./...

.PHONY: lint
lint:
	docker-compose run --rm api golint ./...
