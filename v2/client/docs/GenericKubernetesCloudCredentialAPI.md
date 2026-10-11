# \GenericKubernetesCloudCredentialAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenericKubernetesList**](GenericKubernetesCloudCredentialAPI.md#GenericKubernetesList) | **Get** /api/v1/generic-kubernetes/list | Retrieve list of generic kubernetes cloud credentials
[**GenericKubernetesUpdate**](GenericKubernetesCloudCredentialAPI.md#GenericKubernetesUpdate) | **Put** /api/v1/generic-kubernetes/update | Update Generic kubernetes credentials



## GenericKubernetesList

> GenericKubernetesList GenericKubernetesList(ctx).Limit(limit).Offset(offset).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Id(id).DomainId(domainId).Execute()

Retrieve list of generic kubernetes cloud credentials

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
	limit := int32(56) // int32 |  (optional)
	offset := int32(56) // int32 |  (optional)
	organizationId := int32(56) // int32 |  (optional)
	sortBy := "sortBy_example" // string |  (optional)
	sortDirection := "sortDirection_example" // string |  (optional)
	search := "search_example" // string |  (optional)
	searchId := "searchId_example" // string |  (optional)
	id := int32(56) // int32 |  (optional)
	domainId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GenericKubernetesCloudCredentialAPI.GenericKubernetesList(context.Background()).Limit(limit).Offset(offset).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Id(id).DomainId(domainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GenericKubernetesCloudCredentialAPI.GenericKubernetesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenericKubernetesList`: GenericKubernetesList
	fmt.Fprintf(os.Stdout, "Response from `GenericKubernetesCloudCredentialAPI.GenericKubernetesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenericKubernetesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **organizationId** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **id** | **int32** |  | 
 **domainId** | **int32** |  | 

### Return type

[**GenericKubernetesList**](GenericKubernetesList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenericKubernetesUpdate

> TaikunResult GenericKubernetesUpdate(ctx).UpdateGenericKubernetesCommand(updateGenericKubernetesCommand).Execute()

Update Generic kubernetes credentials

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
	updateGenericKubernetesCommand := *openapiclient.NewUpdateGenericKubernetesCommand(int32(123), "Name_example") // UpdateGenericKubernetesCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GenericKubernetesCloudCredentialAPI.GenericKubernetesUpdate(context.Background()).UpdateGenericKubernetesCommand(updateGenericKubernetesCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GenericKubernetesCloudCredentialAPI.GenericKubernetesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenericKubernetesUpdate`: TaikunResult
	fmt.Fprintf(os.Stdout, "Response from `GenericKubernetesCloudCredentialAPI.GenericKubernetesUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenericKubernetesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateGenericKubernetesCommand** | [**UpdateGenericKubernetesCommand**](UpdateGenericKubernetesCommand.md) |  | 

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

