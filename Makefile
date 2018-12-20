example: ## rebuild the example from scratch
	rm -rf example
	cookiecutter https://github.com/gofunct/cookiecutter-grpcgo.git
	cd example; make build

help: ## help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort