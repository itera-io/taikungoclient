#!/bin/bash

# Module name
GITHUB_USERNAME="itera-io"
GITHUB_REPO="taikungoclient"

# File names
FILE_WEB="swagger-web.json"
FILE_SHOWBACK="swagger-showback.json"

# Generator refuses to rewrite some files
rm -rf ./client
rm -rf ./showbackclient

# Generate client
java -jar openapi-generator-cli.jar generate -i ./"$FILE_WEB" \
-g go \
--additional-properties=packageName=taikuncore \
--additional-properties=enumClassPrefix=true \
--git-user-id="$GITHUB_USERNAME" \
--git-repo-id="$GITHUB_REPO/client" \
-o ./client

# Delete automatically generated go.mod and go.sum. This interferes with our own package.
rm ./client/go.mod
rm ./client/go.sum

# Substitute one occuency of 'w.CreateFormFile' to send Content-Type: application/json (accepted in API)
# instead of Content-Type: application/octet-stream (rejected in Taikun API)
# This makes uploading files (like GCP import project json) through TaikunGoCLient work.
# The echo needs to escape backquotes used to escape other backquotes - hence the backquote hell.
# Import inputs missing dependencies
sed -i "s/part, err := w.CreateFormFile(formFile.formFileName, filepath.Base(formFile.fileName))/\/\/ Custom replace to make Content-Type application\/json - fixes GCP json upload\n				h := make(textproto.MIMEHeader)\n				h.Set(\"Content-Disposition\",fmt.Sprintf(\`form-data; name=\"%s\"; filename=\"%s\"\`,escapeQuotes(formFile.formFileName),escapeQuotes(filepath.Base(formFile.fileName))))\n				h.Set(\"Content-Type\", \"application\/json\")\n				part, err := w.CreatePart(h)\n\n				\/\/ Old version\n				\/\/ part, err := w.CreateFormFile(formFile.formFileName, filepath.Base(formFile.fileName))/" ./client/client.go
echo -e "\nfunc escapeQuotes(s string) string {\n  var quoteEscaper = strings.NewReplacer(\"\\\\\\\", \"\\\\\\\\\\\\\\\", \`\"\`, \"\\\\\\\\\\\"\")\n  return quoteEscaper.Replace(s)\n}" >> ./client/client.go
sed -i "s/import (/import (\n	\"net\/textproto\"/" ./client/client.go

# Generate showback client
java -jar openapi-generator-cli.jar generate -i ./"$FILE_SHOWBACK" \
-g go \
--additional-properties=packageName=taikunshowback  \
--additional-properties=enumClassPrefix=true \
--git-user-id="$GITHUB_USERNAME" \
--git-repo-id="$GITHUB_REPO/showbackclient" \
-o ./showbackclient

# Delete automatically generated go.mod and go.sum. This interferes with our own package.
rm ./showbackclient/go.mod
rm ./showbackclient/go.sum
