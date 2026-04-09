BINARY=taikun

generate: generate-swagger generate-showback ## Generates Go bindings out of Swagger definition

generate-swagger: ## Generates Go binding out of core Swagger API definition
	rm -rf ./client
	openapi-generator generate -i ./swagger-taikun.json -g go \
	--additional-properties=packageName=taikuncore \
	--additional-properties=enumClassPrefix=true \
	--additional-properties=disallowAdditionalPropertiesIfNotPresent=false \
	--git-user-id=itera-io \
	--git-repo-id=taikungoclient/client \
	-o ./client
	rm -f ./client/go.*
	sed -i "s/part, err := w.CreateFormFile(formFile.formFileName, filepath.Base(formFile.fileName))/\/\/ Custom replace to make Content-Type application\/json - fixes GCP json upload\n				h := make(textproto.MIMEHeader)\n				h.Set(\"Content-Disposition\",fmt.Sprintf(\`form-data; name=\"%s\"; filename=\"%s\"\`,escapeQuotes(formFile.formFileName),escapeQuotes(filepath.Base(formFile.fileName))))\n				h.Set(\"Content-Type\", \"application\/json\")\n				part, err := w.CreatePart(h)\n\n				\/\/ Old version\n				\/\/ part, err := w.CreateFormFile(formFile.formFileName, filepath.Base(formFile.fileName))/" ./client/client.go
	echo -e "\nfunc escapeQuotes(s string) string {\n  var quoteEscaper = strings.NewReplacer(\"\\\\\\\", \"\\\\\\\\\\\\\\\", \`\"\`, \"\\\\\\\\\\\"\")\n  return quoteEscaper.Replace(s)\n}" >> ./client/client.go
	sed -i "s/import (/import (\n	\"net\/textproto\"/" ./client/client.go

generate-showback: ## Generates Go binding out of showback Swagger API definition
	rm -rf ./showbackclient
	openapi-generator generate -i ./swagger-showback.json -g go \
	--additional-properties=packageName=taikunshowback  \
	--additional-properties=enumClassPrefix=true \
    --additional-properties=disallowAdditionalPropertiesIfNotPresent=false \
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

