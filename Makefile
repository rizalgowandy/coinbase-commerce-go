.PHONY: help
help: # print all available make commands and their usages.
	@printf "\e[32musage: make [target]\n\n\e[0m"
	@grep -E '^[a-zA-Z_-]+:.*?# .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?# "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: setup
setup: # install configuration and dependencies for development.
	@./scripts/setup.sh

.PHONY: linter
linter: # run linter to keep code clean.
	@./scripts/linter.sh

.PHONY: test
test: # run all tests.
	@./scripts/test.sh

.PHONY: build
build: # ensure all binary can be build.
	@go build -o bin/coinbase-commerce-go && rm bin/coinbase-commerce-go

.PHONY: generate
generate: # generate all go generate command inside internal package.
	@go generate -v ./...
