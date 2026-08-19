#!/bin/bash

# Module name
GITHUB_USERNAME="itera-io"
GITHUB_REPO="taikungoclient/v2"

# File names
FILE_WEB="v2/swagger-web.json"
FILE_SHOWBACK="v2/swagger-showback.json"

# Generator refuses to rewrite some files
rm -rf v2/client
rm -rf v2/showbackclient

# Generate client
java -jar /tmp/openapi-generator-cli.jar generate -i ./"$FILE_WEB" \
-g go \
--additional-properties=packageName=taikuncore \
--additional-properties=enumClassPrefix=true \
--additional-properties=disallowAdditionalPropertiesIfNotPresent=false \
--git-user-id="$GITHUB_USERNAME" \
--git-repo-id="$GITHUB_REPO/client" \
-o v2/client

# Delete automatically generated go.mod and go.sum. This interferes with our own package.
rm v2/client/go.mod
rm v2/client/go.sum

# Substitute one occuency of 'w.CreateFormFile' to send Content-Type: application/json (accepted in API)
# instead of Content-Type: application/octet-stream (rejected in Taikun API)
# This makes uploading files (like GCP import project json) through TaikunGoCLient work.
# The echo needs to escape backquotes used to escape other backquotes - hence the backquote hell.
# Import inputs missing dependencies
sed -i "s/part, err := w.CreateFormFile(formFile.formFileName, filepath.Base(formFile.fileName))/\/\/ Custom replace to make Content-Type application\/json - fixes GCP json upload\n				h := make(textproto.MIMEHeader)\n				h.Set(\"Content-Disposition\",fmt.Sprintf(\`form-data; name=\"%s\"; filename=\"%s\"\`,escapeQuotes(formFile.formFileName),escapeQuotes(filepath.Base(formFile.fileName))))\n				h.Set(\"Content-Type\", \"application\/json\")\n				part, err := w.CreatePart(h)\n\n				\/\/ Old version\n				\/\/ part, err := w.CreateFormFile(formFile.formFileName, filepath.Base(formFile.fileName))/" v2/client/client.go
echo -e "\nfunc escapeQuotes(s string) string {\n  var quoteEscaper = strings.NewReplacer(\"\\\\\\\", \"\\\\\\\\\\\\\\\", \`\"\`, \"\\\\\\\\\\\"\")\n  return quoteEscaper.Replace(s)\n}" >> v2/client/client.go
sed -i "s/import (/import (\n	\"net\/textproto\"/" v2/client/client.go

# Generate showback client
java -jar /tmp/openapi-generator-cli.jar generate -i ./"$FILE_SHOWBACK" \
-g go \
--additional-properties=packageName=taikunshowback  \
--additional-properties=enumClassPrefix=true \
--additional-properties=disallowAdditionalPropertiesIfNotPresent=false \
--git-user-id="$GITHUB_USERNAME" \
--git-repo-id="$GITHUB_REPO/showbackclient" \
-o v2/showbackclient

# Delete automatically generated go.mod and go.sum. This interferes with our own package.
rm v2/showbackclient/go.mod
rm v2/showbackclient/go.sum
