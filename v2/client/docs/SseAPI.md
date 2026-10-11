# \SseAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SseHealth**](SseAPI.md#SseHealth) | **Get** /api/v1/sse/health | Subscribe to real-time infrastructure health state changes
[**SseNotifications**](SseAPI.md#SseNotifications) | **Get** /api/v1/sse/notifications | Subscribe to real-time notification messages for the current user
[**SseRecentEvents**](SseAPI.md#SseRecentEvents) | **Get** /api/v1/sse/recent-events/{organizationId} | Subscribe to real-time audit log events
[**SseUpdateUserPermission**](SseAPI.md#SseUpdateUserPermission) | **Get** /api/v1/sse/update-user-permission | Subscribe to real-time cache invalidation events for the current user



## SseHealth

> SseHealth(ctx).Execute()

Subscribe to real-time infrastructure health state changes

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
	r, err := apiClient.SseAPI.SseHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SseAPI.SseHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSseHealthRequest struct via the builder pattern


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


## SseNotifications

> SseNotifications(ctx).Execute()

Subscribe to real-time notification messages for the current user

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
	r, err := apiClient.SseAPI.SseNotifications(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SseAPI.SseNotifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSseNotificationsRequest struct via the builder pattern


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


## SseRecentEvents

> SseRecentEvents(ctx, organizationId).Status(status).Category(category).ProjectId(projectId).Execute()

Subscribe to real-time audit log events

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
	organizationId := int32(56) // int32 | 
	status := openapiclient.ActionStatus("succeeded") // ActionStatus |  (optional)
	category := openapiclient.ActionType("StartAddServer") // ActionType |  (optional)
	projectId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SseAPI.SseRecentEvents(context.Background(), organizationId).Status(status).Category(category).ProjectId(projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SseAPI.SseRecentEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSseRecentEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | [**ActionStatus**](ActionStatus.md) |  | 
 **category** | [**ActionType**](ActionType.md) |  | 
 **projectId** | **int32** |  | 

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


## SseUpdateUserPermission

> SseUpdateUserPermission(ctx).Execute()

Subscribe to real-time cache invalidation events for the current user

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
	r, err := apiClient.SseAPI.SseUpdateUserPermission(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SseAPI.SseUpdateUserPermission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSseUpdateUserPermissionRequest struct via the builder pattern


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

