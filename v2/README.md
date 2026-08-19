# taikungoclient v2

Module: `github.com/itera-io/taikungoclient/v2`

This is the **new major version** of the Taikun Go client, generated from the post-3.1
Taikun API (the `Account` → `Domain` rename plus the 3.1 feature removals). The **old
client is unchanged** and remains at the root module `github.com/itera-io/taikungoclient`
— pin the last pre-v2 commit/tag until you migrate. Both can be imported side by side.

```go
import taikun "github.com/itera-io/taikungoclient/v2"
```

## Why a new major version
The API surface changed in breaking ways (see `MIGRATION.md`): 84 endpoints removed
(vSphere / Proxmox / OpenShift / Zadara providers, payment / subscription / invoices /
project-templates / taikun-lb, …), the `Account` concept renamed to `Domain` (25 endpoints
+ all `Account*` models/params/enums), and the generator input moved to a normalized 3.0
spec (`openapi-codegen.json`, produced by `make openapi-gen` in the taikun-api repo).

## Regenerating
The v2 client is generated the same way CI generates the root client, but into the `v2/`
module and from the codegen-normalized spec:

```bash
# swagger-web.json here must be taikun-api's docs/openapi-update/openapi-codegen.json
./.github/generate_client_v2.sh
```

`swagger-web.json` is the codegen-normalized 3.0 spec (openapi-generator cannot consume the
raw 3.1 document — `type:[null,X]` arrays, `integer|string` unions and single-element
`oneOf` produce invalid Go). `openapi3.1-after-removal.json` is kept as the raw 3.1
reference only.
