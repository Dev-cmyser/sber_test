dev:
	docker compose -f docker/dev/dev.yml up --no-log-prefix --attach app
.PHONY: run

prod:
	docker compose -f docker/prod/prod.yml up --build --no-log-prefix
.PHONY: prod

lint:
	golangci-lint run -c .golangci.yml --fix
.PHONY: lint

remove:
	docker rm calc_ipoteca
.PHONY: remove

stop:
	docker stop calc_ipoteca
.PHONY: stop

clean: stop remove

PROJECT_ROOT := $(shell git rev-parse --show-toplevel)

COVERAGE_FILE := $(PROJECT_ROOT)/coverage.out

GINKGO = ginkgo

test:
	$(GINKGO)  -r -succinct --cover  --coverprofile=$(COVERAGE_FILE)
	go tool cover -html=$(COVERAGE_FILE) -o coverage.html

show_cover:
	open coverage.html
