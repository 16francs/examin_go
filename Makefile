.PHONY: setup
setup:
	docker-compose build

.PHONY: run
run:
	docker-compose up
