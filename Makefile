.PHONY: user help

user: ## rebuild the example from scratch
	rm -rf user
	cookiecutter https://github.com/gofunct/cookiecutter-grpcgo.git
	cd example; make setup

help: ## help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort