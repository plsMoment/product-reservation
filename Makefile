up:
	@docker compose up

test:
	@go test ./utils

.PHONY: run test