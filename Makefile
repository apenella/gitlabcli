
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
	@grep -E '^[a-zA-Z1-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[1;32m%-20s\033[0m %s\n", $$1, $$2}' | sort
	@echo

## Execute static analysis
static-analysis: vet golint staticcheck

vet: ## Executes the go vet
	@echo
	@echo "$(COLOR_GREEN) Executing go vet $(COLOR_END)"
	@echo
	@$(DOCKER_COMPOSE_BINARY) --project-name gitlabcli-go-vet run --rm --build --entrypoint go ci vet $(LDFLAGS) ./...; $(DOCKER_COMPOSE_BINARY) --project-name gitlabcli-go-vet down --volumes --remove-orphans --timeout 3

golint: ## Executes golint
	@echo
	@echo "$(COLOR_GREEN) Executing golint$(COLOR_END)"
	@echo
	@$(DOCKER_COMPOSE_BINARY) --project-name gitlabcli-golint run --rm --build ci golint ./internal/...; $(DOCKER_COMPOSE_BINARY) --project-name gitlabcli-golint down --volumes --remove-orphans --timeout 3

staticcheck: ## Executes staticcheck
	@echo
	@echo "$(COLOR_GREEN) Executing staticcheck$(COLOR_END)"
	@echo
	@$(DOCKER_COMPOSE_BINARY) --project-name gitlabcli-staticcheck run --rm --build ci staticcheck ./internal/...; $(DOCKER_COMPOSE_BINARY) --project-name gitlabcli-staticcheck down --volumes --remove-orphans --timeout 3

gosec:
	@echo
	@echo "$(COLOR_GREEN) Executing gosec$(COLOR_END)"
	@echo
	@docker run --rm -it -w /app -v "${PWD}":/app securego/gosec /app/...