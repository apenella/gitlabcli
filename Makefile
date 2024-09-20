
DOCKER_COMPOSE_BINARY := $(shell docker compose version > /dev/null 2>&1 && echo "docker compose" || (which docker-compose > /dev/null 2>&1 && echo "docker-compose" || (echo "docker compose not found. Aborting." >&2; exit 1)))

## Colors
COLOR_GREEN=\033[0;32m
COLOR_RED=\033[0;31m
COLOR_BLUE=\033[0;34m
COLOR_PURPLE=\033[0;35m
COLOR_END=\033[0m

.DEFAULT_GOAL := help

help: ## Lists available targets
	@echo
	@echo "Makefile usage:"
	@grep -E '^[a-zA-Z1-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[1;32m%-25s\033[0m %s\n", $$1, $$2}' | sort
	@echo

## Execute static analysis
static-analysis: vet linter staticcheck gosec

ci-go-tools-docker-image: ## Build the docker image
	@echo
	@echo "$(COLOR_BLUE) Building the docker image $(COLOR_END)"
	@echo
	@docker build -t ci-go-tools-docker-image -f build/Dockerfile .

vet: ci-go-tools-docker-image ## Executes the go vet
	@echo
	@echo "$(COLOR_BLUE) Executing go vet $(COLOR_END)"
	@echo
	@docker run --rm -v "${PWD}":/app -w /app ci-go-tools-docker-image go vet ./internal/... && echo "$(COLOR_GREEN) go vet: all files linted$(COLOR_END)" || echo "$(COLOR_RED)go vet: some files not linted$(COLOR_END)"

linter: ci-go-tools-docker-image ## Executes Go linter (golint)
	@echo
	@echo "$(COLOR_BLUE) Executing golint$(COLOR_END)"
	@echo
	@docker run --rm -v "${PWD}":/app -w /app ci-go-tools-docker-image golint ./internal/... && echo "$(COLOR_GREEN) golint: all files linted$(COLOR_END)" || echo "$(COLOR_RED)golint: some files not linted$(COLOR_END)"

staticcheck: ci-go-tools-docker-image ## Executes staticcheck
	@echo
	@echo "$(COLOR_BLUE) Executing staticcheck$(COLOR_END)"
	@echo
	@docker run --rm -v "${PWD}":/app -w /app ci-go-tools-docker-image staticcheck ./internal/... && echo "$(COLOR_GREEN) staticcheck: all files linted$(COLOR_END)" || echo "$(COLOR_RED)staticcheck: some files not linted$(COLOR_END)"

gosec:
	@echo
	@echo "$(COLOR_BLUE) Executing gosec$(COLOR_END)"
	@echo
	@docker run --rm -w /app -v "${PWD}":/app securego/gosec /app/... && echo "$(COLOR_GREEN) gosec: no issues found$(COLOR_END)" || echo "$(COLOR_RED)gosec: some issues found$(COLOR_END)"
