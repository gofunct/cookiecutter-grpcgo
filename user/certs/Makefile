gen: # generate a server.key and server.pem file
	openssl genrsa -out server.key 2048
	openssl req -new -x509 -key server.key -out server.pem -days 3650

help: ## help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort