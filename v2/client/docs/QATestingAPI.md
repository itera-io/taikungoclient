# \QATestingAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QaEndSession**](QATestingAPI.md#QaEndSession) | **Post** /api/v1/qa/sessions/end | End an active QA session without seeded tenant
[**QaSeed**](QATestingAPI.md#QaSeed) | **Post** /api/v1/qa/seed | Seed domain, organization, and domain-admin or robot user for QA
[**QaStartSession**](QATestingAPI.md#QaStartSession) | **Post** /api/v1/qa/sessions/start | Start an ephemeral QA session and receive a short-lived session token
[**QaTeardown**](QATestingAPI.md#QaTeardown) | **Post** /api/v1/qa/teardown | Teardown seeded QA tenant for the current session



## QaEndSession

> TaikunResult QaEndSession(ctx).Execute()

End an active QA session without seeded tenant

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/itera-io/taikungoclient/v2/client"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QATestingAPI.QaEndSession(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QATestingAPI.QaEndSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QaEndSession`: TaikunResult
	fmt.Fprintf(os.Stdout, "Response from `QATestingAPI.QaEndSession`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiQaEndSessionRequest struct via the builder pattern


### Return type

[**TaikunResult**](TaikunResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QaSeed

> SeedQaTenantResult QaSeed(ctx).SeedQaTenantCommand(seedQaTenantCommand).Execute()

Seed domain, organization, and domain-admin or robot user for QA

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/itera-io/taikungoclient/v2/client"
)

func main() {
	seedQaTenantCommand := *openapiclient.NewSeedQaTenantCommand() // SeedQaTenantCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QATestingAPI.QaSeed(context.Background()).SeedQaTenantCommand(seedQaTenantCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QATestingAPI.QaSeed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QaSeed`: SeedQaTenantResult
	fmt.Fprintf(os.Stdout, "Response from `QATestingAPI.QaSeed`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQaSeedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **seedQaTenantCommand** | [**SeedQaTenantCommand**](SeedQaTenantCommand.md) |  | 

### Return type

[**SeedQaTenantResult**](SeedQaTenantResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QaStartSession

> StartQaSessionResult QaStartSession(ctx).Execute()

Start an ephemeral QA session and receive a short-lived session token

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/itera-io/taikungoclient/v2/client"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QATestingAPI.QaStartSession(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QATestingAPI.QaStartSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QaStartSession`: StartQaSessionResult
	fmt.Fprintf(os.Stdout, "Response from `QATestingAPI.QaStartSession`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiQaStartSessionRequest struct via the builder pattern


### Return type

[**StartQaSessionResult**](StartQaSessionResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QaTeardown

> TaikunResult QaTeardown(ctx).TeardownQaTenantCommand(teardownQaTenantCommand).Execute()

Teardown seeded QA tenant for the current session

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/itera-io/taikungoclient/v2/client"
)

func main() {
	teardownQaTenantCommand := *openapiclient.NewTeardownQaTenantCommand("CleanupToken_example") // TeardownQaTenantCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QATestingAPI.QaTeardown(context.Background()).TeardownQaTenantCommand(teardownQaTenantCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QATestingAPI.QaTeardown``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QaTeardown`: TaikunResult
	fmt.Fprintf(os.Stdout, "Response from `QATestingAPI.QaTeardown`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQaTeardownRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teardownQaTenantCommand** | [**TeardownQaTenantCommand**](TeardownQaTenantCommand.md) |  | 

### Return type

[**TaikunResult**](TaikunResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

