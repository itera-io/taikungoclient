# \ProjectQuotasAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectquotasList**](ProjectQuotasAPI.md#ProjectquotasList) | **Get** /api/v1/projectquotas | Retrieve all project quotas
[**ProjectquotasUpdate**](ProjectQuotasAPI.md#ProjectquotasUpdate) | **Post** /api/v1/projectquotas/update | Edit project quota



## ProjectquotasList

> ProjectQuotaList ProjectquotasList(ctx).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).StartRam(startRam).EndRam(endRam).StartDiskSize(startDiskSize).EndDiskSize(endDiskSize).StartCpu(startCpu).EndCpu(endCpu).OrganizationId(organizationId).Id(id).DomainId(domainId).Execute()

Retrieve all project quotas

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
	sortBy := "sortBy_example" // string |  (optional)
	sortDirection := "sortDirection_example" // string |  (optional)
	search := "search_example" // string |  (optional)
	startRam := float64(1.2) // float64 |  (optional)
	endRam := float64(1.2) // float64 |  (optional)
	startDiskSize := int64(789) // int64 |  (optional)
	endDiskSize := int64(789) // int64 |  (optional)
	startCpu := int32(56) // int32 |  (optional)
	endCpu := int32(56) // int32 |  (optional)
	organizationId := int32(56) // int32 |  (optional)
	id := int32(56) // int32 |  (optional)
	domainId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectQuotasAPI.ProjectquotasList(context.Background()).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).StartRam(startRam).EndRam(endRam).StartDiskSize(startDiskSize).EndDiskSize(endDiskSize).StartCpu(startCpu).EndCpu(endCpu).OrganizationId(organizationId).Id(id).DomainId(domainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectQuotasAPI.ProjectquotasList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectquotasList`: ProjectQuotaList
	fmt.Fprintf(os.Stdout, "Response from `ProjectQuotasAPI.ProjectquotasList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectquotasListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **startRam** | **float64** |  | 
 **endRam** | **float64** |  | 
 **startDiskSize** | **int64** |  | 
 **endDiskSize** | **int64** |  | 
 **startCpu** | **int32** |  | 
 **endCpu** | **int32** |  | 
 **organizationId** | **int32** |  | 
 **id** | **int32** |  | 
 **domainId** | **int32** |  | 

### Return type

[**ProjectQuotaList**](ProjectQuotaList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectquotasUpdate

> ProjectquotasUpdate(ctx).UpdateQuotaCommand(updateQuotaCommand).Execute()

Edit project quota

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
	updateQuotaCommand := *openapiclient.NewUpdateQuotaCommand() // UpdateQuotaCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectQuotasAPI.ProjectquotasUpdate(context.Background()).UpdateQuotaCommand(updateQuotaCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectQuotasAPI.ProjectquotasUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectquotasUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateQuotaCommand** | [**UpdateQuotaCommand**](UpdateQuotaCommand.md) |  | 

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

