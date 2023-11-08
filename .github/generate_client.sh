#!/bin/bash

# Module name
GITHUB_USERNAME="itera-io"
GITHUB_REPO="taikungoclient"

FILE_WEB="swagger-web.json"
FILE_SHOWBACK="swagger-showback.json"

#MAIN_API_URL="https://api.taikun.dev/swagger/v1/swagger.json"
#SHOWBACK_API_URL="https://api.taikun.dev/showback/swagger/v1/swagger.json"
#rm "$FILE_WEB" "$FILE_SHOWBACK"
#
#wget "$MAIN_API_URL" -O "./$FILE_WEB"
#wget "$SHOWBACK_API_URL" -O "./$FILE_SHOWBACK"

rm -rf ./client         # Generator refuses to rewrite some files
#rm -rf ./showbackclient # Generator refuses to rewrite some files

# Generate client
java -jar openapi-generator-cli.jar --version
java -jar openapi-generator-cli.jar
java --version
ls -la
java -jar openapi-generator-cli.jar generate -i ./"$FILE_WEB" \
-g go \
--additional-properties=packageName=taikuncore \
--additional-properties=enumClassPrefix=true \
--git-user-id="$GITHUB_USERNAME" \
--git-repo-id="$GITHUB_REPO/client" \
-o ./client

#openapi-generator-cli generate -i ./"$FILE_SHOWBACK" \
#-g go \
#--additional-properties=packageName=taikunshowback  \
#--additional-properties=enumClassPrefix=true \
#--git-user-id="$GITHUB_USERNAME" \
#--git-repo-id="$GITHUB_REPO/showbackclient" \
#-o ./showbackclient

