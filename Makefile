# Highlight
HL = @printf "\033[36m>> $1\033[0m\n"

default: help

help:
	@echo "Usage: make <TARGET>\n\nTargets:"
	@grep -E "^[\. a-zA-Z_-]+:.*?## .*$$" $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' |sort

go-mod-verify: ## Verify that the dependencies of the current module have not been modified
	$(call HL,go-mod-verify)
	@go mod verify

check-fmt: ## Check code formatting
	$(call HL,check-fmt)
	F=$(F) script/check-fmt.sh `find . -name "*.go"`

lint: ## Run go lint
	$(call HL,lint)
	F=$(F) script/lint.sh `find . -name "*.go"`

test: ## Run unit tests with ARGS=<go_test_args>
	$(call HL,test)
	@go test -count=1 $(ARGS) ./...
