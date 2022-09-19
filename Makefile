.PHONY: help
.DEFAULT_GOAL := help

help: ## Show options.
	@grep -E '^[a-zA-Z]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
