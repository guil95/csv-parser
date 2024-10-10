mock-generate: ## Generate mocks
	go mod tidy
	docker run --rm -v "$(PWD):/app" -w /app -t vektra/mockery --all --dir ./internal/parser --case underscore

run:
	go run cmd/main.go -file=$(FILE)