.PHONY: user help
MODULES=$(shell export GOMODULES=on)

gen: ## rebuild the example grpclab project from scratch
	cd ../grpclab; rm -rf *
	cd ../grpclab; cookiecutter --no-input -v https://github.com/gofunct/cookiecutter-grpcgo.git
	cd ../grpclab; mv grpclab/* . && rm -rf grpclab
	cd ../grpclab && export GOMODULES=on && make init
	cd ../grpclab; git add -A && git commit -m "build" && git push origin master
	grpclab serve

debug: ## rebuild the example grpclab project from scratch
	cd ../grpclab; rm -rf *
	cd ../grpclab; cookiecutter --no-input -v https://github.com/gofunct/cookiecutter-grpcgo.git
	cd ../grpclab; mv grpclab/* . && rm -rf grpclab
	cd ../grpclab && export GOMODULES=on && make proto-files

build-docker:
	cd hack; make build

help: ## help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort