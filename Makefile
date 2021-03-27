.PHONY: lint ## Run linter
lint:
	@golangci-lint --exclude-use-default=false run ./...

.PHONY: test
test: runTests ## Run tests
	@if [ -f coverage.out ]; then go tool cover -func=coverage.out && rm coverage.out; fi

.PHONY: coverage
coverage: runTests ## Show coverage
	@if [ -f coverage.out ]; then go tool cover -html=coverage.out && rm coverage.out; fi

runTests:
	@go test -coverprofile=coverage.out ./...

all:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "    \033[36m%-20s\033[0m %s\n", $$1, $$2}'
	@echo
