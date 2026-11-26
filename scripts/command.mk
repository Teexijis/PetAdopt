.PHONY: generate-api
generate-api:
	@bash ./scripts/generate.sh

.PHONY: up
up:
	docker compose up -d

.PHONY: down
down:
	docker compose down

.PHONY: cov
cov:
	cd account && go test -cover ./...