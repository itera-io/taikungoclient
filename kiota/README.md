# taikungoclient — Kiota client (evaluation)

Module: `github.com/itera-io/taikungoclient/kiota`

A **[Kiota](https://github.com/microsoft/kiota)-generated** Go client for the post-3.1 Taikun
API, generated **natively from OpenAPI 3.1** (no 3.0 down-convert). Provided alongside the
openapi-generator `v2` client so we can compare the two before committing to one for the
long-lived SDK. The original root client is unchanged.

## Why Kiota
Kiota consumes the true 3.1 spec directly. It builds cleanly from taikun-api's
`openapi-codegen-3.1.json` (the codegen-normalized 3.1 spec — the raw .NET-native 3.1 output
first needs the spurious `integer|string` unions, enum int-defaults, single-element `oneOf`
and empty schemas collapsed; that normalization lives in taikun-api's `make openapi-gen`).

## Shape (differs from the openapi-generator client)
Fluent request builders + the Kiota runtime, e.g.:
```go
// adapter is a *kiotahttp.NetHttpRequestAdapter with your auth provider
client := client.NewTaikunClient(adapter)
domain, err := client.Api().V1().Domains().ByDomainId(id).Get(ctx, nil)
```
Runtime deps: `github.com/microsoft/kiota-abstractions-go` + kiota serialization packages.

## Regenerating
```bash
# swagger-web.json = taikun-api docs/openapi-update/openapi-codegen-3.1.json
./generate.sh
```

## Status / TODO
- Generated client **builds + vets clean**.
- Auth is applied via a Kiota `AuthenticationProvider` on the request adapter — the
  JWT-login/refresh + robot basic-auth wrapper (today's `taikungoclient.go`) is **not yet
  ported**; that's the follow-up if we adopt Kiota.
