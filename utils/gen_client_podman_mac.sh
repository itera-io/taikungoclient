#!/bin/bash

# Generates the Taikun Go API client libraries from OpenAPI specifications
# using the openapi-generator-cli container image run via Podman.
#
# Two clients are generated:
#   - client/          - Main API client, generated from $API_SPEC_FILE
#   - showbackclient/  - Showback API client, generated from $API_SHOWBACK_SPEC_FILE
#
# After generation, a patch is applied to client/client.go to set
# Content-Type: application/json on multipart form file uploads, which is
# required for GCP JSON key uploads.
#
# Prerequisites:
#   - Podman must be installed (see https://podman.io or use homebrew)
#   - The OpenAPI spec files must be present in the working directory:
#       swagger-account-api.json
#       swagger-account-showback-api.json
#
# Usage:
#   ./utils/gen_client_podman_mac.sh --api-spec <file> --showback-spec <file>
#
# Arguments:
#   --api-spec       Path to the main API OpenAPI spec file (required)
#   --showback-spec  Path to the showback API OpenAPI spec file (required)
#
# Example:
#   ./utils/gen_client_podman_mac.sh \
#     --api-spec swagger-account-api.json \
#     --showback-spec swagger-account-showback-api.json
#
# Runs from the repository root. The existing client/ and showbackclient/
# directories will be replaced with freshly generated versions.

set -o errexit
set -o pipefail
set -o nounset

IMAGE="docker.io/openapitools/openapi-generator-cli:v7.20.0"

API_SPEC_FILE=""
API_SHOWBACK_SPEC_FILE=""

while [[ $# -gt 0 ]]; do
    case "$1" in
        --api-spec)
            API_SPEC_FILE="$2"
            shift 2
            ;;
        --showback-spec)
            API_SHOWBACK_SPEC_FILE="$2"
            shift 2
            ;;
        *)
            echo "ERROR: Unknown argument: $1" >&2
            echo "Usage: $0 --api-spec <file> --showback-spec <file>" >&2
            exit 1
            ;;
    esac
done

if [[ -z "$API_SPEC_FILE" ]]; then
    echo "ERROR: --api-spec is required." >&2
    echo "Usage: $0 --api-spec <file> --showback-spec <file>" >&2
    exit 1
fi

if [[ -z "$API_SHOWBACK_SPEC_FILE" ]]; then
    echo "ERROR: --showback-spec is required." >&2
    echo "Usage: $0 --api-spec <file> --showback-spec <file>" >&2
    exit 1
fi

echo "Starting API client generation..."

echo "Checking for openapi-generator-cli..."
if ! command -v podman run $IMAGE &> /dev/null; then
    echo "ERROR: openapi-generator-cli must be installed to proceed!" >&2
    exit 1
fi
echo "openapi-generator-cli found."

echo "Validating General API spec file: $API_SPEC_FILE..."
if ! podman run -v $(pwd)/$API_SPEC_FILE:/tmp/$API_SPEC_FILE $IMAGE validate -i /tmp/$API_SPEC_FILE; then
    echo "ERROR: API specification validation failed." >&2
    exit 1
fi
echo "General API specification is valid."

echo "Validating Showback API spec file: $API_SHOWBACK_SPEC_FILE..."
if ! podman run -v $(pwd)/$API_SHOWBACK_SPEC_FILE:/tmp/$API_SHOWBACK_SPEC_FILE $IMAGE validate -i /tmp/$API_SHOWBACK_SPEC_FILE; then
    echo "ERROR: API specification validation failed." >&2
    exit 1
fi
echo "Showback API specification is valid."

echo "Generating General API client from $API_SPEC_FILE..."
rm -rf "client_new" && mkdir -p "client_new"
podman run -v $(pwd)/$API_SPEC_FILE:/tmp/$API_SPEC_FILE \
    -v $(pwd)/client_new:/tmp/client $IMAGE generate \
    -i /tmp/$API_SPEC_FILE \
    -g go \
    --additional-properties=packageName=taikunshowback \
    --additional-properties=enumClassPrefix=true \
    --additional-properties=disallowAdditionalPropertiesIfNotPresent=false \
    --skip-validate-spec \
    -o /tmp/client

echo "Patching w.CreateFormFile to accept Content-Type: application/json"
sed -i.bak "s/part, err := w.CreateFormFile(formFile.formFileName, filepath.Base(formFile.fileName))/\/\/ Custom replace to make Content-Type application\/json - fixes GCP json upload\n				h := make(textproto.MIMEHeader)\n				h.Set(\"Content-Disposition\",fmt.Sprintf(\`form-data; name=\"%s\"; filename=\"%s\"\`,escapeQuotes(formFile.formFileName),escapeQuotes(filepath.Base(formFile.fileName))))\n				h.Set(\"Content-Type\", \"application\/json\")\n				part, err := w.CreatePart(h)\n\n				\/\/ Old version\n				\/\/ part, err := w.CreateFormFile(formFile.formFileName, filepath.Base(formFile.fileName))/" ./client_new/client.go && rm ./client_new/client.go.bak
echo -e "\nfunc escapeQuotes(s string) string {\n  var quoteEscaper = strings.NewReplacer(\"\\\\\\\", \"\\\\\\\\\\\\\\\", \`\"\`, \"\\\\\\\\\\\"\")\n  return quoteEscaper.Replace(s)\n}" >> ./client_new/client.go
sed -i.bak "s/import (/import (\n	\"net\/textproto\"/" ./client_new/client.go && rm ./client_new/client.go.bak

echo "Generating Showback API client from $API_SHOWBACK_SPEC_FILE..."
rm -rf "client_showback_new" && mkdir -p "client_showback_new"
podman run -v $(pwd)/$API_SHOWBACK_SPEC_FILE:/tmp/$API_SHOWBACK_SPEC_FILE \
    -v $(pwd)/client_showback_new:/tmp/client_showback $IMAGE generate \
    -i /tmp/$API_SHOWBACK_SPEC_FILE \
    -g go \
    --additional-properties=packageName=taikunshowback \
    --additional-properties=enumClassPrefix=true \
    --additional-properties=disallowAdditionalPropertiesIfNotPresent=false \
    --skip-validate-spec \
    -o /tmp/client_showback

echo "Deleting automatically generated go.mod and go.sum"
rm ./client_new/go.mod
rm ./client_new/go.sum
rm ./client_showback_new/go.mod
rm ./client_showback_new/go.sum

echo "Removing old client directories"
rm -rf ./client
rm -rf ./showbackclient

echo "Renaming new clients"
mv client_new client
mv client_showback_new showbackclient

echo "API clients generated successfully!"
