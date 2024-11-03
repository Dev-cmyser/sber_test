run:
	docker compose -f docker/dev/dev.yml up --no-log-prefix --attach app
.PHONY: run

prod:
	docker compose -f docker/prod/prod.yml up --build --no-log-prefix
.PHONY: prod

lint:
	golangci-lint run -c .golangci.yml
.PHONY: lint
