#!/usr/bin/env bash

set -e

# Module name
module_name="github.com/itera-io/taikungoclient"

# Application name
app_name="taikungoclient"
showback_app_name="showbackgoclient"
showback_package_name="showbackclient"

# Temporary swagger patch files
readonly swagger_patch_file="swagger-patch.json"
readonly showback_swagger_patch_file="showback-swagger-patch.json"

# TODO: DELETE ME, THIS WAS JUST FOR LOCAL TESTS!!!
# rm -fv "${swagger_patch_file}" "${showback_swagger_patch_file}"

# Delete client/ and models/ directories if they already exist
rm -rfv client/
rm -rfv models/

# Delete go.mod and go.sum files if they already exist
rm -rfv go.mod
rm -rfv go.sum

# Ensure go-swagger names the package for the 'Documentation' endpoint 'doc' instead of 'documentation'
sed 's/"Documentation"/"Doc"/g' swagger.json >"${swagger_patch_file}"
sed 's/"Documentation"/"Doc"/g' showback-swagger.json >"${showback_swagger_patch_file}"

# Make all datetimes nullable
sed -i 's/"format": "date-time"/"x-nullable": true, "format": "date-time"/g' "${swagger_patch_file}"
sed -i 's/"format": "date-time"/"x-nullable": true, "format": "date-time"/g' "${showback_swagger_patch_file}"

# Make showback_credential_id nullable
refline="\"showbackCredentialId\":"
sed -i "/${refline}/a \"x-nullable\": true," "${swagger_patch_file}"
sed -i "/${refline}/a \"x-nullable\": true," "${showback_swagger_patch_file}"

# Remove omitempty
refline="\"(ruleDiscountRate|globalDiscountRate|discountRate|displayName)\":"
sed -Ei "/${refline}/a \"x-omitempty\": false," "${swagger_patch_file}"
sed -Ei "/${refline}/a \"x-omitempty\": false," "${showback_swagger_patch_file}"

# Remove omitempty
refline="\"type\": \"boolean\""
sed -i "s/${refline}/\"x-omitempty\": false,${refline}/g" "${swagger_patch_file}"
sed -i "s/${refline}/\"x-omitempty\": false,${refline}/g" "${showback_swagger_patch_file}"

# Add new CustomProblemDetailsMg definition in both swagger files.
# MG in CustomProblemDetailsMg stands for Manually Generated, the name of this
# resource has to be unique enough not to conflict with any future definitions
refline="\"definitions\": {"
sed -i "/${refline}/r .github/codegen/custom_error_type.json" "${swagger_patch_file}"
sed -i "/${refline}/r .github/codegen/custom_error_type.json" "${showback_swagger_patch_file}"

# Lines containing definitions/ProblemDetails" or
# definitions/ValidationProblemDetails" are replaced with the content of
# .github/codegen/custom_error_response_schema.json. The latter contains a JSON
# array of CustomProblemDetailsMg. This is done for both swagger files.
for refline in 'definitions\/ProblemDetails"'; do
# for refline in 'definitions\/ProblemDetails"' 'definitions\/ValidationProblemDetails"'; do
  for swagger_file in "${swagger_patch_file}" "${showback_swagger_patch_file}"; do
    sed -i -e "/${refline}/ {" \
      -e 'r .github/codegen/custom_error_response_schema.json' \
      -e 'd' -e '}' "${swagger_file}"
  done
done

# TODO: DELETE ME, THIS WAS JUST FOR LOCAL TESTS!!!
# python -mjson.tool "${swagger_patch_file}" > /dev/null
# python -mjson.tool "${showback_swagger_patch_file}" > /dev/null
# exit 0

# Initialize go module
go mod init "${module_name}"

# Generate the client
./swagger generate client -f "${swagger_patch_file}" -A "${app_name}" $@
./swagger generate client -f "${showback_swagger_patch_file}" -c "${showback_package_name}" -A "${showback_app_name}" $@
rm -f "${swagger_patch_file}"
rm -f "${showback_swagger_patch_file}"

# Insert code to support extra media types
refline="transport := httptransport.New"
client_definition_file="client/${app_name}_client.go"
showback_client_definition_file="${showback_package_name}/${showback_app_name}_client.go"
sed -i "/${refline}/r .github/codegen/media_type_support_code.go" "${client_definition_file}"
go fmt "${client_definition_file}"

sed -i "/${refline}/r .github/codegen/media_type_support_code.go" "${showback_client_definition_file}"
go fmt "${showback_client_definition_file}"

# Set default scheme of showback client to HTTPS from HTTP
showbackclient="./showbackclient/showbackgoclient_client.go"
if [[ -f "${showbackclient}" ]]; then
  sed -i 's/var DefaultSchemes = \[\]string{"http"}/var DefaultSchemes = []string{"https"}/' "${showbackclient}"
fi

# Tidy go module
go mod tidy -compat=1.17
