# \TrustedRegistriesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TrustedregistriesCreate**](TrustedRegistriesAPI.md#TrustedregistriesCreate) | **Post** /api/v1/trustedregistries/create | Create trusted registries for access profile
[**TrustedregistriesDelete**](TrustedRegistriesAPI.md#TrustedregistriesDelete) | **Delete** /api/v1/trustedregistries/{id} | Delete trusted registry
[**TrustedregistriesEdit**](TrustedRegistriesAPI.md#TrustedregistriesEdit) | **Put** /api/v1/trustedregistries/edit/{id} | Edit trusted registry
[**TrustedregistriesList**](TrustedRegistriesAPI.md#TrustedregistriesList) | **Get** /api/v1/trustedregistries/{accessProfileId} | List trusted registries by profile id



## TrustedregistriesCreate

> ApiResponse TrustedregistriesCreate(ctx).CreateTrustedRegistriesCommand(createTrustedRegistriesCommand).Execute()

Create trusted registries for access profile

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
	createTrustedRegistriesCommand := *openapiclient.NewCreateTrustedRegistriesCommand("Registry_example", int32(123)) // CreateTrustedRegistriesCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrustedRegistriesAPI.TrustedregistriesCreate(context.Background()).CreateTrustedRegistriesCommand(createTrustedRegistriesCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrustedRegistriesAPI.TrustedregistriesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TrustedregistriesCreate`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `TrustedRegistriesAPI.TrustedregistriesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTrustedregistriesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTrustedRegistriesCommand** | [**CreateTrustedRegistriesCommand**](CreateTrustedRegistriesCommand.md) |  | 

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


## TrustedregistriesDelete

> TrustedregistriesDelete(ctx, id).Execute()

Delete trusted registry

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
	r, err := apiClient.TrustedRegistriesAPI.TrustedregistriesDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrustedRegistriesAPI.TrustedregistriesDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiTrustedregistriesDeleteRequest struct via the builder pattern


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


## TrustedregistriesEdit

> TrustedregistriesEdit(ctx, id).TrustedRegistryEditDto(trustedRegistryEditDto).Execute()

Edit trusted registry

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
	trustedRegistryEditDto := *openapiclient.NewTrustedRegistryEditDto("Registry_example") // TrustedRegistryEditDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TrustedRegistriesAPI.TrustedregistriesEdit(context.Background(), id).TrustedRegistryEditDto(trustedRegistryEditDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrustedRegistriesAPI.TrustedregistriesEdit``: %v\n", err)
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

Other parameters are passed through a pointer to a apiTrustedregistriesEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **trustedRegistryEditDto** | [**TrustedRegistryEditDto**](TrustedRegistryEditDto.md) |  | 

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


## TrustedregistriesList

> []TrustedRegistriesListDto TrustedregistriesList(ctx, accessProfileId).Search(search).Execute()

List trusted registries by profile id

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
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrustedRegistriesAPI.TrustedregistriesList(context.Background(), accessProfileId).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrustedRegistriesAPI.TrustedregistriesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TrustedregistriesList`: []TrustedRegistriesListDto
	fmt.Fprintf(os.Stdout, "Response from `TrustedRegistriesAPI.TrustedregistriesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessProfileId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTrustedregistriesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** |  | 

### Return type

[**[]TrustedRegistriesListDto**](TrustedRegistriesListDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

