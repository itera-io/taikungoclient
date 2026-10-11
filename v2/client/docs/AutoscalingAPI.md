# \AutoscalingAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutoscalingDisable**](AutoscalingAPI.md#AutoscalingDisable) | **Post** /api/v1/autoscaling/disable | Disable autoscaling
[**AutoscalingEdit**](AutoscalingAPI.md#AutoscalingEdit) | **Post** /api/v1/autoscaling/edit | Edit autoscaling
[**AutoscalingEnable**](AutoscalingAPI.md#AutoscalingEnable) | **Post** /api/v1/autoscaling/enable | Enable autoscaling



## AutoscalingDisable

> AutoscalingDisable(ctx).DisableAutoscalingCommand(disableAutoscalingCommand).Execute()

Disable autoscaling

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
	disableAutoscalingCommand := *openapiclient.NewDisableAutoscalingCommand(int32(123)) // DisableAutoscalingCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AutoscalingAPI.AutoscalingDisable(context.Background()).DisableAutoscalingCommand(disableAutoscalingCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoscalingAPI.AutoscalingDisable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutoscalingDisableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **disableAutoscalingCommand** | [**DisableAutoscalingCommand**](DisableAutoscalingCommand.md) |  | 

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


## AutoscalingEdit

> AutoscalingEdit(ctx).EditAutoscalingCommand(editAutoscalingCommand).Execute()

Edit autoscaling

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
	editAutoscalingCommand := *openapiclient.NewEditAutoscalingCommand(int32(123), int32(123), int32(123)) // EditAutoscalingCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AutoscalingAPI.AutoscalingEdit(context.Background()).EditAutoscalingCommand(editAutoscalingCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoscalingAPI.AutoscalingEdit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutoscalingEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **editAutoscalingCommand** | [**EditAutoscalingCommand**](EditAutoscalingCommand.md) |  | 

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


## AutoscalingEnable

> AutoscalingEnable(ctx).EnableAutoscalingCommand(enableAutoscalingCommand).Execute()

Enable autoscaling

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
	enableAutoscalingCommand := *openapiclient.NewEnableAutoscalingCommand(int32(123), int32(123), int32(123), float64(123), "Flavor_example", false) // EnableAutoscalingCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AutoscalingAPI.AutoscalingEnable(context.Background()).EnableAutoscalingCommand(enableAutoscalingCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoscalingAPI.AutoscalingEnable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutoscalingEnableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enableAutoscalingCommand** | [**EnableAutoscalingCommand**](EnableAutoscalingCommand.md) |  | 

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

