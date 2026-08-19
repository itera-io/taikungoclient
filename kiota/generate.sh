#!/usr/bin/env bash
# Regenerate the Kiota Go client from the codegen-normalized 3.1 spec.
# swagger-web.json must be taikun-api's docs/openapi-update/openapi-codegen-3.1.json
# (Kiota consumes true 3.1 natively; the raw .NET 3.1 output needs the codegen normalization
# — integer|string unions, enum int-defaults, single-element oneOf, empty schemas — applied
# by taikun-api's `make openapi-gen`).
set -euo pipefail
kiota generate -l Go \
  -d swagger-web.json \
  -o client \
  -c TaikunClient \
  -n github.com/itera-io/taikungoclient/kiota/client \
  --clean-output true
