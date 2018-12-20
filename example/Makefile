.PHONY: build build-alpine clean test help default vendor setup regen

service_name =	$(word 2,$(subst /, ,$1))
BIN_NAME=example

VERSION := $(shell grep "const Version " version/version.go | sed -E 's/.*"(.+)"$$/\1/')
GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_DIRTY=$(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
BUILD_DATE=$(shell date '+%Y-%m-%d-%H:%M:%S')
IMAGE_NAME := "gofunct/example"
SOURCES :=	$(shell find . -name "*.proto" -not -path ./vendor/\*)
TARGETS_GO :=	$(foreach source, $(SOURCES), $(source)_go)
TARGETS_TMPL :=	$(foreach source, $(SOURCES), $(source)_tmpl)
PROJECT_PATH := "github.com/gofunct/example"
SERVICE_PATH := "github.com/gofunct/example/services"

default: test

setup: ## download dependencies and tls certificates
	go mod vendor
	git clone -q https://github.com/gofunct/templates.git
	brew install protobuf
	brew install prototool
	go get -u \
		google.golang.org/grpc \
		github.com/golang/protobuf/protoc-gen-go \
		go get github.com/gogo/protobuf/... \
		github.com/ckaznocha/protoc-gen-lint \
		github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc \
		moul.io/protoc-gen-gotemplate

build: setup vendor $(TARGETS_GO) $(TARGETS_TMPL) certs ## Compile the project.
	@echo "building ${BIN_NAME} ${VERSION}"
	@echo "GOPATH=${GOPATH}"
	go build -ldflags "-X github.com/gofunct/example/version.GitCommit=${GIT_COMMIT}${GIT_DIRTY} -X github.com/gofunct/example/version.BuildDate=${BUILD_DATE}" -o bin/${BIN_NAME}

gen:
	prototool all services

regen: ## delete all files and regenerate project
	rm -rf *
	rm .*
	cookiecutter https://github.com/gofunct/cookiecutter-grpcgo.git


vendor: ## Runs go mod vendor, mostly used for ci.
	go mod vendor


build-alpine: ## Compile optimized for alpine linux.
	@echo "building ${BIN_NAME} ${VERSION}"
	@echo "GOPATH=${GOPATH}"
	go build -ldflags '-w -linkmode external -extldflags "-static" -X github.com/gofunct/example/version.GitCommit=${GIT_COMMIT}${GIT_DIRTY} -X github.com/gofunct/example/version.BuildDate=${BUILD_DATE}' -o bin/${BIN_NAME}

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

clean: ## Clean the directory tree.
	@test ! -e bin/${BIN_NAME} || rm bin/${BIN_NAME}


certs: ## Generate a server key and certificate
	openssl genrsa -out certs/server.key 2048
	openssl req -new -x509 -key server.key -out certs/server.pem -days 3650


test: ## Run tests on a compiled project.
	go test ./...

help: ## help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort