run:
	@go build ./cmd/imagery && ./imagery

dev:
	@./bin/air -c .air.toml