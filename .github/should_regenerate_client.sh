#!/usr/bin/env bash

set -e

# Sha256sum file paths
old_sha1sum_path="./.github/.swagger_sha1sum.txt"
new_sha1sum_path="./.github/.swagger_sha1sum_new.txt"

# Generate new sha1sum of swagger.json
sha1sum swagger.json > "${new_sha1sum_path}"

# If running for the first time, save the sha1sum and return true
if [[ ! -f "${old_sha1sum_path}" ]]; then
  mv "${new_sha1sum_path}" "${old_sha1sum_path}"
  exit 0
fi

# If sha1sums differ, return true and save new sha1sum, otherwise return false
if diff "${new_sha1sum_path}" "${old_sha1sum_path}" &>/dev/null; then
  rm -f "${new_sha1sum_path}"
  exit 1
else
  mv "${new_sha1sum_path}" "${old_sha1sum_path}"
  exit 0
fi
