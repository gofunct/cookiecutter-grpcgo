# A Makefile is a collection of rules.  Each rule is a recipe to do a specific
# thing, sort of like a grunt task or an npm package.json script.
#
# A rule looks like this:
#
# <target>: <prerequisites...>
# 	<commands>
#
# The "target" is required.  The prerequisites are optional, and the commands
# are also optional, but you have to have one or the other.
#
# To suppress output, add @ signs before each line
#
#
# Within the ruleset, we can use some special syntax to refer to the input
# file and the output file.  Here are the special variables:
#
# $@  The file that is being made right now by this rule (aka the "target")
#     You can remember this because it's like the "$@" list in a
#     shell script.  @ is like a letter "a" for "arguments.
#     When you type "make foo", then "foo" is the argument.
#
# $<  The input file (that is, the first prerequisite in the list)
#     You can remember this becasue the < is like a file input
#     pipe in bash.  `head <foo.txt` is using the contents of
#     foo.txt as the input.  Also the < points INto the $
#
# $^  This is the list of ALL input files, not just the first one.
#     You can remember it because it's like $<, but turned up a notch.
#     If a file shows up more than once in the input list for some reason,
#     it's still only going to show one time in $^.
#
# $?  All the input files that are newer than the target
#     It's like a question. "Wait, why are you doing this?  What
#     files changed to make this necessary?"
#
# $$  A literal $ character inside of the rules section
#
# $*  The "stem" part that matched in the rule definition's % bit
#     You can remember this because in make rules, % is like * on
#     the shell, so $* is telling you what matched the pattern..

IMAGE_NAME := "{{cookiecutter.docker_hub_username}}/{{cookiecutter.app_name}}"
PROJECT_PATH := "github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}"
SERVICE_PATH := "github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/services"
BIN_NAME={{cookiecutter.app_name}}


.PHONY: build build-alpine clean test help default vendor init gen deps

SOURCES :=	$(shell find . -name "*.proto" -not -path ./vendor/\*)
TARGETS_TMPL :=	$(foreach source, $(SOURCES), $(source)_tmpl)
VERSION := $(shell grep "const Version " version/version.go | sed -E 's/.*"(.+)"$$/\1/')
GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_DIRTY=$(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
BUILD_DATE=$(shell date '+%Y-%m-%d-%H:%M:%S')
MODULES=$(shell export GOMODULES=on)

service_name =	$(word 2,$(subst /, ,$1))

default: test

init: gen $(TARGETS_TMPL)
	go mod vendor
	go install

gen:
	go mod vendor
	docker run -v `pwd`:/defs colemanword/protoc-all:1.17_0 -d services -l go

protofiles:
	docker run -v $(PWD):/defs colemanword/prototool:1.17_0 files ./services

$(TARGETS_TMPL): %_tmpl:
	git clone -q https://github.com/gofunct/templates.git
	protoc -I. --gotemplate_out=destination_dir=gen/$(call service_name,$*),debug=true,template_dir=templates:gen "$*"
	rm -rf templates

build: setup vendor $(TARGETS_GO) $(TARGETS_TMPL) certs ## Compile the project.
	@echo "building ${BIN_NAME} ${VERSION}"
	@echo "GOPATH=${GOPATH}"
	go build -ldflags "-X github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/version.GitCommit=${GIT_COMMIT}${GIT_DIRTY} -X github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/version.BuildDate=${BUILD_DATE}" -o bin/${BIN_NAME}

sources: $(SOURCES)

vendor: ## Runs go mod vendor, mostly used for ci.
	go mod vendor

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