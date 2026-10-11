# \ExecutorsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExecutorsCreate**](ExecutorsAPI.md#ExecutorsCreate) | **Post** /api/v1/executors/create | Create an executor
[**ExecutorsDelete**](ExecutorsAPI.md#ExecutorsDelete) | **Post** /api/v1/executors/delete | Delete an executor
[**ExecutorsEditHealth**](ExecutorsAPI.md#ExecutorsEditHealth) | **Put** /api/v1/executors/edit/health | Update health status of the executor by Id
[**ExecutorsList**](ExecutorsAPI.md#ExecutorsList) | **Get** /api/v1/executors | Retrieve a list of kube configs of executors
[**ExecutorsToggle**](ExecutorsAPI.md#ExecutorsToggle) | **Post** /api/v1/executors/toggle | Toggle an executor



## ExecutorsCreate

> ExecutorsCreate(ctx).CreateExecutorCommand(createExecutorCommand).Execute()

Create an executor

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
	createExecutorCommand := *openapiclient.NewCreateExecutorCommand("Name_example", false, false, "Continent_example", "KubeConfig_example") // CreateExecutorCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExecutorsAPI.ExecutorsCreate(context.Background()).CreateExecutorCommand(createExecutorCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecutorsAPI.ExecutorsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExecutorsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createExecutorCommand** | [**CreateExecutorCommand**](CreateExecutorCommand.md) |  | 

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


## ExecutorsDelete

> ExecutorsDelete(ctx).DeleteExecutorCommand(deleteExecutorCommand).Execute()

Delete an executor

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
	deleteExecutorCommand := *openapiclient.NewDeleteExecutorCommand(int32(123)) // DeleteExecutorCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExecutorsAPI.ExecutorsDelete(context.Background()).DeleteExecutorCommand(deleteExecutorCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecutorsAPI.ExecutorsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExecutorsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteExecutorCommand** | [**DeleteExecutorCommand**](DeleteExecutorCommand.md) |  | 

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


## ExecutorsEditHealth

> ExecutorsEditHealth(ctx).UpdateExecutorHealthStatusCommand(updateExecutorHealthStatusCommand).Execute()

Update health status of the executor by Id

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
	updateExecutorHealthStatusCommand := *openapiclient.NewUpdateExecutorHealthStatusCommand(int32(123), openapiclient.ExecutorHealth("None")) // UpdateExecutorHealthStatusCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExecutorsAPI.ExecutorsEditHealth(context.Background()).UpdateExecutorHealthStatusCommand(updateExecutorHealthStatusCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecutorsAPI.ExecutorsEditHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExecutorsEditHealthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateExecutorHealthStatusCommand** | [**UpdateExecutorHealthStatusCommand**](UpdateExecutorHealthStatusCommand.md) |  | 

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


## ExecutorsList

> ExecutorListResponse ExecutorsList(ctx).Offset(offset).Limit(limit).SortBy(sortBy).SortDirection(sortDirection).Search(search).Id(id).Execute()

Retrieve a list of kube configs of executors

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
	offset := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 |  (optional)
	sortBy := "sortBy_example" // string |  (optional)
	sortDirection := "sortDirection_example" // string |  (optional)
	search := "search_example" // string |  (optional)
	id := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExecutorsAPI.ExecutorsList(context.Background()).Offset(offset).Limit(limit).SortBy(sortBy).SortDirection(sortDirection).Search(search).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecutorsAPI.ExecutorsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExecutorsList`: ExecutorListResponse
	fmt.Fprintf(os.Stdout, "Response from `ExecutorsAPI.ExecutorsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExecutorsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **id** | **int32** |  | 

### Return type

[**ExecutorListResponse**](ExecutorListResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExecutorsToggle

> ExecutorsToggle(ctx).ToggleExecutorCommand(toggleExecutorCommand).Execute()

Toggle an executor

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
	toggleExecutorCommand := *openapiclient.NewToggleExecutorCommand(int32(123), "Mode_example") // ToggleExecutorCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExecutorsAPI.ExecutorsToggle(context.Background()).ToggleExecutorCommand(toggleExecutorCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecutorsAPI.ExecutorsToggle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExecutorsToggleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **toggleExecutorCommand** | [**ToggleExecutorCommand**](ToggleExecutorCommand.md) |  | 

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

