#!/bin/bash

DOCKER_REPO=${DOCKER_REPO}
NAMESPACE=${NAMESPACE:-colemanword}
GRPC_VERSION=${GRPC_VERSION:-1.17}
GRPC_JAVA_VERSION=${GRPC_JAVA_VERSION:-1.17}
BUILD_VERSION=${BUILD_VERSION:-0}
CONTAINER=${DOCKER_REPO}${NAMESPACE}
LATEST=${1:false}
BUILDS=("protoc-all" "protoc" "prototool" "grpc-cli" "gen-grpc-gateway")