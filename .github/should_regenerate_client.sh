#!/usr/bin/env bash

set -e

# Sha256sum file paths
old_sha1sum_path="./.github/swagger_sha1sum.txt"
new_sha1sum_path="./.github/swagger_sha1sum_new.txt"

# Combine swagger files
cat swagger-web.json swagger-showback.json > combined-swagger.json
# Generate new sha1sum of swagger.json
sha1sum combined-swagger.json > "${new_sha1sum_path}"
# Remove combined-swagger.json file
rm -f combined-swagger.json
# If running for the first time, save the sha1sum and return true
if [[ ! -f "${old_sha1sum_path}" ]]; then
  mv "${new_sha1sum_path}" "${old_sha1sum_path}"
  echo "Running for the first time, saved ${old_sha1sum_path} and exiting with code 0"
  exit 0
fi

# If sha1sums differ, return true and save new sha1sum, otherwise return false
if diff "${new_sha1sum_path}" "${old_sha1sum_path}" &>/dev/null; then
  rm -f "${new_sha1sum_path}"
  echo "combined-swagger.json checksum hasn't changed, exiting with code 1"
  exit 1
else
  mv "${new_sha1sum_path}" "${old_sha1sum_path}"
  echo "combined-swagger.json checksum changed, saved ${old_sha1sum_path} and exiting with code 0"
  exit 0
fi
