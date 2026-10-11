# \AdminAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersOverloadProtectionStatus**](AdminAPI.md#UsersOverloadProtectionStatus) | **Get** /api/v1/users/overload-protection | Get overload protection status
[**UsersToggleOverloadProtection**](AdminAPI.md#UsersToggleOverloadProtection) | **Post** /api/v1/users/toggleoverloadprotection | Set overload protection mode



## UsersOverloadProtectionStatus

> UsersOverloadProtectionStatus(ctx).Execute()

Get overload protection status

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
	r, err := apiClient.AdminAPI.UsersOverloadProtectionStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.UsersOverloadProtectionStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUsersOverloadProtectionStatusRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersToggleOverloadProtection

> UsersToggleOverloadProtection(ctx).ToggleOverloadProtectionCommand(toggleOverloadProtectionCommand).Execute()

Set overload protection mode

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
	toggleOverloadProtectionCommand := *openapiclient.NewToggleOverloadProtectionCommand(openapiclient.OverloadProtectionMode("Normal")) // ToggleOverloadProtectionCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdminAPI.UsersToggleOverloadProtection(context.Background()).ToggleOverloadProtectionCommand(toggleOverloadProtectionCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.UsersToggleOverloadProtection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersToggleOverloadProtectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **toggleOverloadProtectionCommand** | [**ToggleOverloadProtectionCommand**](ToggleOverloadProtectionCommand.md) |  | 

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

