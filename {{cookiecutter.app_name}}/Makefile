.PHONY: build build-alpine clean test help default init gen deps

VERSION := $(shell grep "const Version " version/version.go | sed -E 's/.*"(.+)"$$/\1/')
GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_DIRTY=$(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
BUILD_DATE=$(shell date '+%Y-%m-%d-%H:%M:%S')
IMAGE_NAME := "{{cookiecutter.docker_hub_username}}/{{cookiecutter.app_name}}"
PROJECT_PATH := "github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}"
API_PATH := "github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/api"
BIN_NAME={{cookiecutter.app_name}}
service_name =	$(word 2,$(subst /, ,$1))

default: help

init: deps gen ## initialize
	go install
	{{cookiecutter.app_name}} serve

deps: ## download dependencies
	dep init
	brew install protobuf
	brew install prototool
	go get -u \
		google.golang.org/grpc \
		github.com/golang/protobuf/protoc-gen-go \
		github.com/ckaznocha/protoc-gen-lint \
		github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc \
		github.com/gogo/protobuf/...

	dep ensure

gen: ## compile all protobuf files in api/
	prototool all api

{% if cookiecutter.use_docker == "y" %}
build-alpine: ## Compile optimized for alpine linux.
	@echo "building ${BIN_NAME} ${VERSION}"
	@echo "GOPATH=${GOPATH}"
	go build -ldflags '-w -linkmode external -extldflags "-static" -X github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/version.GitCommit=${GIT_COMMIT}${GIT_DIRTY} -X github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/version.BuildDate=${BUILD_DATE}' -o bin/${BIN_NAME}

package: ## Build final docker image with just the go binary inside
	@echo "building image ${BIN_NAME} ${VERSION} $(GIT_COMMIT)"
	docker build --build-arg VERSION=${VERSION} --build-arg GIT_COMMIT=$(GIT_COMMIT) -t $(IMAGE_NAME):local .

tag: ## Tag image created by package with latest, git commit and version'
	@echo "Tagging: latest ${VERSION} $(GIT_COMMIT)"
	docker tag $(IMAGE_NAME):local $(IMAGE_NAME):$(GIT_COMMIT)
	docker tag $(IMAGE_NAME):local $(IMAGE_NAME):${VERSION}
	docker tag $(IMAGE_NAME):local $(IMAGE_NAME):latest

push: tag ## Push tagged images to registry'
	@echo "Pushing docker image to registry: latest ${VERSION} $(GIT_COMMIT)"
	docker push $(IMAGE_NAME):$(GIT_COMMIT)
	docker push $(IMAGE_NAME):${VERSION}
	docker push $(IMAGE_NAME):latest
{% endif %}
clean: ## Clean the directory tree.
	@test ! -e bin/${BIN_NAME} || rm bin/${BIN_NAME}

{% if cookiecutter.use_tls == "y" %}
certs: ## Generate a server key and certificate
	openssl genrsa -out certs/server.key 2048
	openssl req -new -x509 -key server.key -out certs/server.pem -days 3650
{% endif %}

test: ## Run tests on a compiled project.
	go test ./...

help: ## help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort