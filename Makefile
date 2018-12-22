.PHONY: user help

gen: ## rebuild the example grpclab project from scratch
	rm -rf grpclab
	cookiecutter --no-input https://github.com/gofunct/cookiecutter-grpcgo.git
	cd grpclab; make init

help: ## help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort