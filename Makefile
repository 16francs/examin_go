.PHONY: build
build:
	docker-compose build

.PHONY: run
run:
	docker-compose up

.PHONY: test
test:
	docker-compose run --rm api go test -v -cover ./...

.PHONY: lint
lint:
	docker-compose run --rm api golint ./...
