# \AllowedHostAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllowedhostCreate**](AllowedHostAPI.md#AllowedhostCreate) | **Post** /api/v1/allowedhost/create | Create access profile allowed host
[**AllowedhostDelete**](AllowedHostAPI.md#AllowedhostDelete) | **Delete** /api/v1/allowedhost/{id} | Delete access profile allowed host
[**AllowedhostEdit**](AllowedHostAPI.md#AllowedhostEdit) | **Put** /api/v1/allowedhost/edit/{id} | Edit access profile allowed host
[**AllowedhostList**](AllowedHostAPI.md#AllowedhostList) | **Get** /api/v1/allowedhost/list/{accessProfileId} | List allowed hosts by access profile id



## AllowedhostCreate

> ApiResponse AllowedhostCreate(ctx).CreateAllowedHostCommand(createAllowedHostCommand).Execute()

Create access profile allowed host

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
	createAllowedHostCommand := *openapiclient.NewCreateAllowedHostCommand(int32(123), "Description_example", "IpAddress_example", int32(123)) // CreateAllowedHostCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AllowedHostAPI.AllowedhostCreate(context.Background()).CreateAllowedHostCommand(createAllowedHostCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllowedHostAPI.AllowedhostCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AllowedhostCreate`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AllowedHostAPI.AllowedhostCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAllowedhostCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAllowedHostCommand** | [**CreateAllowedHostCommand**](CreateAllowedHostCommand.md) |  | 

### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AllowedhostDelete

> AllowedhostDelete(ctx, id).Execute()

Delete access profile allowed host

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
	r, err := apiClient.AllowedHostAPI.AllowedhostDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllowedHostAPI.AllowedhostDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAllowedhostDeleteRequest struct via the builder pattern


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


## AllowedhostEdit

> AllowedhostEdit(ctx, id).EditAllowedHostDto(editAllowedHostDto).Execute()

Edit access profile allowed host

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
	editAllowedHostDto := *openapiclient.NewEditAllowedHostDto("Description_example", "IpAddress_example", int32(123)) // EditAllowedHostDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AllowedHostAPI.AllowedhostEdit(context.Background(), id).EditAllowedHostDto(editAllowedHostDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllowedHostAPI.AllowedhostEdit``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAllowedhostEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **editAllowedHostDto** | [**EditAllowedHostDto**](EditAllowedHostDto.md) |  | 

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


## AllowedhostList

> AllowedHostList AllowedhostList(ctx, accessProfileId).Offset(offset).Limit(limit).Search(search).SortBy(sortBy).SortDirection(sortDirection).Execute()

List allowed hosts by access profile id

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
	accessProfileId := int32(56) // int32 | 
	offset := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 |  (optional)
	search := "search_example" // string |  (optional)
	sortBy := "sortBy_example" // string |  (optional)
	sortDirection := "sortDirection_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AllowedHostAPI.AllowedhostList(context.Background(), accessProfileId).Offset(offset).Limit(limit).Search(search).SortBy(sortBy).SortDirection(sortDirection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllowedHostAPI.AllowedhostList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AllowedhostList`: AllowedHostList
	fmt.Fprintf(os.Stdout, "Response from `AllowedHostAPI.AllowedhostList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessProfileId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAllowedhostListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** |  | 
 **limit** | **int32** |  | 
 **search** | **string** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 

### Return type

[**AllowedHostList**](AllowedHostList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

