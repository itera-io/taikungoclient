#!/usr/bin/env bash

set -e

# Module name
module_name="github.com/itera-io/taikungoclient"
client_definition_file="client/github_com_itera_io_taikungoclient_client.go"

# Delete client/ and models/ directories if they already exist
rm -rfv client/
rm -rfv models/

# Delete go.mod and go.sum files if they already exist
rm -rfv go.mod
rm -rfv go.sum

# Ensure go-swagger names the package for the 'Documentation' endpoint 'doc' instead of 'documentation'
sed 's/"Documentation"/"Doc"/g' swagger.json > swagger-patch.json

# Initialize go module
go mod init "${module_name}"

# Generate the client
swagger generate client -c "client" -f swagger-patch.json -A "${module_name}" $@
rm -f swagger-patch.json

# Insert code to support extra media types
reference_line="transport := httptransport.New"
sed -i "/${reference_line}/r .github/codegen/media_type_support_code.go" "${client_definition_file}"
go fmt "${client_definition_file}"

# Tidy go module
go mod tidy
