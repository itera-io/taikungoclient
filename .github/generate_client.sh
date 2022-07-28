#!/usr/bin/env bash

set -e

# Module name
module_name="github.com/itera-io/taikungoclient"

# Application name
app_name="taikungoclient"
showback_app_name="showbackgoclient"
showback_package_name="showbackclient"

# Delete client/ and models/ directories if they already exist
rm -rfv client/
rm -rfv models/

# Delete go.mod and go.sum files if they already exist
rm -rfv go.mod
rm -rfv go.sum

# Ensure go-swagger names the package for the 'Documentation' endpoint 'doc' instead of 'documentation'
sed 's/"Documentation"/"Doc"/g' swagger.json >swagger-patch.json
sed 's/"Documentation"/"Doc"/g' showback-swagger.json >showback-swagger-patch.json

# Make all datetimes nullable
sed -i 's/"format": "date-time"/"x-nullable": true, "format": "date-time"/g' swagger-patch.json
sed -i 's/"format": "date-time"/"x-nullable": true, "format": "date-time"/g' showback-swagger-patch.json

# Make showback_credential_id nullable
refline="\"showbackCredentialId\":"
sed -i "/${refline}/a \"x-nullable\": true," swagger-patch.json
sed -i "/${refline}/a \"x-nullable\": true," showback-swagger-patch.json

# Remove omitempty
refline="\"(ruleDiscountRate|globalDiscountRate|discountRate|displayName)\":"
sed -Ei "/${refline}/a \"x-omitempty\": false," swagger-patch.json
sed -Ei "/${refline}/a \"x-omitempty\": false," showback-swagger-patch.json

# Remove omitempty
refline="\"type\": \"boolean\""
sed -i "s/${refline}/\"x-omitempty\": false,${refline}/g" swagger-patch.json
sed -i "s/${refline}/\"x-omitempty\": false,${refline}/g" showback-swagger-patch.json

# Initialize go module
go mod init "${module_name}"

# Generate the client
./swagger generate client -f swagger-patch.json -A "${app_name}" $@
./swagger generate client -f showback-swagger-patch.json -c "${showback_package_name}" -A "${showback_app_name}" $@
rm -f swagger-patch.json
rm -f showback-swagger-patch.json

# Insert code to support extra media types
refline="transport := httptransport.New"
client_definition_file="client/${app_name}_client.go"
showback_client_definition_file="${showback_package_name}/${showback_app_name}_client.go"
sed -i "/${refline}/r .github/codegen/media_type_support_code.go" "${client_definition_file}"
go fmt "${client_definition_file}"

sed -i "/${refline}/r .github/codegen/media_type_support_code.go" "${showback_client_definition_file}"
go fmt "${showback_client_definition_file}"

# Tidy go module
go mod tidy
