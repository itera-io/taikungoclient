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

java -jar openapi-generator-cli.jar generate -i ./"$FILE_SHOWBACK" \
-g go \
--additional-properties=packageName=taikunshowback  \
--additional-properties=enumClassPrefix=true \
--git-user-id="$GITHUB_USERNAME" \
--git-repo-id="$GITHUB_REPO/showbackclient" \
-o ./showbackclient
