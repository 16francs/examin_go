.PHONY: build
build:
	docker-compose build

.PHONY: run
run:
	docker-compose up

.PHONY: db-create
db-create:
	docker-compose up -d
	- sh ./docker/bin/create-db.sh
	docker-compose down 

.PHONY: db-migrate
db-migrate:
	docker-compose run --rm api go run infrastructure/datastore/migrate/migrate.go

.PHONY: test
test:
	docker-compose run --rm api go test -v -cover ./...

.PHONY: lint
lint:
	docker-compose run --rm api golint ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: swagger-start
swagger-start:
	sh ./bin/swagger-start.sh

.PHONY: swagger-stop
swagger-stop:
	docker stop swagger
