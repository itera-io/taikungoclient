#!/usr/bin/env bash

set -e

# Module name
module_name="github.com/itera-io/taikungoclient"

# Delete client/ and models/ directories if they already exist
[ -d client ] && rm -rf client && echo "Removed client/ directory"
[ -d models ] && rm -rf models && echo "Removed models/ directory"

# Delete go.mod and go.sum files if they already exist
[ -f go.mod ] && rm -rf go.mod && echo "Removed go.mod file"
[ -f go.sum ] && rm -rf go.sum && echo "Removed go.sum file"

# Ensure go-swagger names the package for the 'Documentation' endpoint 'doc' instead of 'documentation'
sed 's/"Documentation"/"Doc"/g' swagger.json >swagger-patch.json

# Initialize go module
go mod init "${module_name}" && echo "Initialized ${module_name} module"

# Generate the client
echo -n "Generating client..."
swagger generate client -c "client" -f swagger-patch.json -A "${module_name}" $@
rm -f swagger-patch.json
echo -e " Done."

# Insert code to support extra media types
reference_line="transport := httptransport.New"
client_definition_file="client/${module_name}"_client.go
sed -i "/${reference_line}/r media_type_support_code.go" "${client_definition_file}"
go fmt "${client_definition_file}"

# Tidy go module
go mod tidy && echo "Tidied ${module_name} module"
