# \GatewayAPIAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GatewayApiDisable**](GatewayAPIAPI.md#GatewayApiDisable) | **Post** /api/v1/gateway-api/disable | Disable Gateway API modern CRD bundle
[**GatewayApiEnable**](GatewayAPIAPI.md#GatewayApiEnable) | **Post** /api/v1/gateway-api/enable | Enable Gateway API modern CRD bundle



## GatewayApiDisable

> GatewayApiDisable(ctx).DisableGatewayApiModernCommand(disableGatewayApiModernCommand).Execute()

Disable Gateway API modern CRD bundle

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
	disableGatewayApiModernCommand := *openapiclient.NewDisableGatewayApiModernCommand(int32(123)) // DisableGatewayApiModernCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GatewayAPIAPI.GatewayApiDisable(context.Background()).DisableGatewayApiModernCommand(disableGatewayApiModernCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPIAPI.GatewayApiDisable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayApiDisableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **disableGatewayApiModernCommand** | [**DisableGatewayApiModernCommand**](DisableGatewayApiModernCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayApiEnable

> GatewayApiEnable(ctx).EnableGatewayApiModernCommand(enableGatewayApiModernCommand).Execute()

Enable Gateway API modern CRD bundle

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
	enableGatewayApiModernCommand := *openapiclient.NewEnableGatewayApiModernCommand(int32(123)) // EnableGatewayApiModernCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GatewayAPIAPI.GatewayApiEnable(context.Background()).EnableGatewayApiModernCommand(enableGatewayApiModernCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPIAPI.GatewayApiEnable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayApiEnableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enableGatewayApiModernCommand** | [**EnableGatewayApiModernCommand**](EnableGatewayApiModernCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

