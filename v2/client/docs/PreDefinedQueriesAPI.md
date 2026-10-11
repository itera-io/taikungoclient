# \PreDefinedQueriesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PredefinedqueriesCreate**](PreDefinedQueriesAPI.md#PredefinedqueriesCreate) | **Post** /api/v1/predefinedqueries/prometheus/dashboard/create | Create prometheus dashboard pre defined query
[**PredefinedqueriesDelete**](PreDefinedQueriesAPI.md#PredefinedqueriesDelete) | **Delete** /api/v1/predefinedqueries/prometheus/dashboard/delete/{id} | Delete prometheus dashboard pre defined query
[**PredefinedqueriesList**](PreDefinedQueriesAPI.md#PredefinedqueriesList) | **Get** /api/v1/predefinedqueries/prometheus/dashboard/list/{projectId} | Get list of pre defined organization prometheus dashboard elements
[**PredefinedqueriesPrometheusDashboardCommon**](PreDefinedQueriesAPI.md#PredefinedqueriesPrometheusDashboardCommon) | **Get** /api/v1/predefinedqueries/prometheus/dashboard/common/{projectId} | et list of pre defined common prometheus dashboard elements
[**PredefinedqueriesUpdate**](PreDefinedQueriesAPI.md#PredefinedqueriesUpdate) | **Post** /api/v1/predefinedqueries/prometheus/dashboard/update | Update prometheus dashboard pre defined query



## PredefinedqueriesCreate

> PredefinedqueriesCreate(ctx).PrometheusDashboardCreateCommand(prometheusDashboardCreateCommand).Execute()

Create prometheus dashboard pre defined query

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
	prometheusDashboardCreateCommand := *openapiclient.NewPrometheusDashboardCreateCommand(int32(123), "Name_example", "Expression_example", "Description_example", "CategoryName_example") // PrometheusDashboardCreateCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PreDefinedQueriesAPI.PredefinedqueriesCreate(context.Background()).PrometheusDashboardCreateCommand(prometheusDashboardCreateCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PreDefinedQueriesAPI.PredefinedqueriesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPredefinedqueriesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **prometheusDashboardCreateCommand** | [**PrometheusDashboardCreateCommand**](PrometheusDashboardCreateCommand.md) |  | 

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


## PredefinedqueriesDelete

> PredefinedqueriesDelete(ctx, id).Execute()

Delete prometheus dashboard pre defined query

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
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PreDefinedQueriesAPI.PredefinedqueriesDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PreDefinedQueriesAPI.PredefinedqueriesDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPredefinedqueriesDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PredefinedqueriesList

> []PrometheusDashboardListDto PredefinedqueriesList(ctx, projectId).Execute()

Get list of pre defined organization prometheus dashboard elements

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
	projectId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PreDefinedQueriesAPI.PredefinedqueriesList(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PreDefinedQueriesAPI.PredefinedqueriesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PredefinedqueriesList`: []PrometheusDashboardListDto
	fmt.Fprintf(os.Stdout, "Response from `PreDefinedQueriesAPI.PredefinedqueriesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPredefinedqueriesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PrometheusDashboardListDto**](PrometheusDashboardListDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PredefinedqueriesPrometheusDashboardCommon

> []PrometheusDashboardListDto PredefinedqueriesPrometheusDashboardCommon(ctx, projectId).Execute()

et list of pre defined common prometheus dashboard elements

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
	projectId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PreDefinedQueriesAPI.PredefinedqueriesPrometheusDashboardCommon(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PreDefinedQueriesAPI.PredefinedqueriesPrometheusDashboardCommon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PredefinedqueriesPrometheusDashboardCommon`: []PrometheusDashboardListDto
	fmt.Fprintf(os.Stdout, "Response from `PreDefinedQueriesAPI.PredefinedqueriesPrometheusDashboardCommon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPredefinedqueriesPrometheusDashboardCommonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PrometheusDashboardListDto**](PrometheusDashboardListDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PredefinedqueriesUpdate

> PredefinedqueriesUpdate(ctx).PrometheusDashboardUpdateCommand(prometheusDashboardUpdateCommand).Execute()

Update prometheus dashboard pre defined query

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
	prometheusDashboardUpdateCommand := *openapiclient.NewPrometheusDashboardUpdateCommand(int32(123), "Name_example", "Expression_example", "Description_example", "CategoryName_example") // PrometheusDashboardUpdateCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PreDefinedQueriesAPI.PredefinedqueriesUpdate(context.Background()).PrometheusDashboardUpdateCommand(prometheusDashboardUpdateCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PreDefinedQueriesAPI.PredefinedqueriesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPredefinedqueriesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **prometheusDashboardUpdateCommand** | [**PrometheusDashboardUpdateCommand**](PrometheusDashboardUpdateCommand.md) |  | 

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

