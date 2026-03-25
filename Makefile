BINARY=taikun

generate: generate-swagger generate-showback ## Generates Go bindings out of Swagger definition

generate-swagger: ## Generates Go binding out of core Swagger API definition
	openapi-generator generate -i ./swagger-taikun.json -g go \
	--additional-properties=packageName=taikuncore \
	--additional-properties=enumClassPrefix=true \
	--git-user-id=itera-io \
	--git-repo-id=taikungoclient/client \
	-o ./client
	rm -f ./client/go.*

generate-showback: ## Generates Go binding out of showback Swagger API definition
	openapi-generator generate -i ./swagger-showback.json -g go \
	--additional-properties=packageName=taikunshowback  \
	--additional-properties=enumClassPrefix=true \
	--git-user-id=itera-io \
	--git-repo-id=taikungoclient/showbackclient \
	-o ./showbackclient
	rm ./showbackclient/go.*

build: ## Builds Taikun Go Client wrapper
	go build -o ${BINARY} .

go-tidy: ## Runs go mod tidy
	go mod tidy

go-vendor: go-tidy ## Runs go mod tidy && go mod vendor
	go mod vendor

.PHONY: help
help: # Credits to https://gist.github.com/prwhite/8168133 for this handy oneliner
	@awk 'BEGIN {FS = ":.*##"; printf "Usage: make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

