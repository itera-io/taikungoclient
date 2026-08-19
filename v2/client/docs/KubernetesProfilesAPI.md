# \KubernetesProfilesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KubernetesprofilesCreate**](KubernetesProfilesAPI.md#KubernetesprofilesCreate) | **Post** /api/v1/kubernetesprofiles | Add kubernetes profile
[**KubernetesprofilesDelete**](KubernetesProfilesAPI.md#KubernetesprofilesDelete) | **Delete** /api/v1/kubernetesprofiles/{id} | Delete kubernetes profile
[**KubernetesprofilesDropdown**](KubernetesProfilesAPI.md#KubernetesprofilesDropdown) | **Get** /api/v1/kubernetesprofiles | Retrieve all kubernetes profiles for organization
[**KubernetesprofilesList**](KubernetesProfilesAPI.md#KubernetesprofilesList) | **Get** /api/v1/kubernetesprofiles/list | Retrieve all kubernetes profiles
[**KubernetesprofilesLockManager**](KubernetesProfilesAPI.md#KubernetesprofilesLockManager) | **Post** /api/v1/kubernetesprofiles/lockmanager | Kubernetes profile lock/unlock



## KubernetesprofilesCreate

> ApiResponse KubernetesprofilesCreate(ctx).CreateKubernetesProfileCommand(createKubernetesProfileCommand).Execute()

Add kubernetes profile

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
	createKubernetesProfileCommand := *openapiclient.NewCreateKubernetesProfileCommand("Name_example", false, false, NullableInt32(123), false, false, false, false, false) // CreateKubernetesProfileCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesProfilesAPI.KubernetesprofilesCreate(context.Background()).CreateKubernetesProfileCommand(createKubernetesProfileCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesProfilesAPI.KubernetesprofilesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesprofilesCreate`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesProfilesAPI.KubernetesprofilesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesprofilesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createKubernetesProfileCommand** | [**CreateKubernetesProfileCommand**](CreateKubernetesProfileCommand.md) |  | 

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


## KubernetesprofilesDelete

> KubernetesprofilesDelete(ctx, id).Execute()

Delete kubernetes profile

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
	r, err := apiClient.KubernetesProfilesAPI.KubernetesprofilesDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesProfilesAPI.KubernetesprofilesDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiKubernetesprofilesDeleteRequest struct via the builder pattern


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


## KubernetesprofilesDropdown

> []KubernetesProfilesEntity KubernetesprofilesDropdown(ctx).OrganizationId(organizationId).Search(search).CloudId(cloudId).Offset(offset).Limit(limit).Execute()

Retrieve all kubernetes profiles for organization

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
	organizationId := int32(56) // int32 |  (optional)
	search := "search_example" // string |  (optional)
	cloudId := int32(56) // int32 |  (optional)
	offset := int32(56) // int32 |  (optional) (default to 0)
	limit := int32(56) // int32 |  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesProfilesAPI.KubernetesprofilesDropdown(context.Background()).OrganizationId(organizationId).Search(search).CloudId(cloudId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesProfilesAPI.KubernetesprofilesDropdown``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesprofilesDropdown`: []KubernetesProfilesEntity
	fmt.Fprintf(os.Stdout, "Response from `KubernetesProfilesAPI.KubernetesprofilesDropdown`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesprofilesDropdownRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 
 **search** | **string** |  | 
 **cloudId** | **int32** |  | 
 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 50]

### Return type

[**[]KubernetesProfilesEntity**](KubernetesProfilesEntity.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesprofilesList

> KubernetesProfilesList KubernetesprofilesList(ctx).OrganizationId(organizationId).DomainId(domainId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Id(id).Execute()

Retrieve all kubernetes profiles

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
	organizationId := int32(56) // int32 |  (optional)
	domainId := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 |  (optional)
	offset := int32(56) // int32 |  (optional)
	sortBy := "sortBy_example" // string |  (optional)
	sortDirection := "sortDirection_example" // string |  (optional)
	search := "search_example" // string |  (optional)
	searchId := "searchId_example" // string |  (optional)
	id := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesProfilesAPI.KubernetesprofilesList(context.Background()).OrganizationId(organizationId).DomainId(domainId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesProfilesAPI.KubernetesprofilesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesprofilesList`: KubernetesProfilesList
	fmt.Fprintf(os.Stdout, "Response from `KubernetesProfilesAPI.KubernetesprofilesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesprofilesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 
 **domainId** | **int32** |  | 
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **id** | **int32** |  | 

### Return type

[**KubernetesProfilesList**](KubernetesProfilesList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesprofilesLockManager

> KubernetesprofilesLockManager(ctx).KubernetesProfilesLockManagerCommand(kubernetesProfilesLockManagerCommand).Execute()

Kubernetes profile lock/unlock

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
	kubernetesProfilesLockManagerCommand := *openapiclient.NewKubernetesProfilesLockManagerCommand(int32(123), "Mode_example") // KubernetesProfilesLockManagerCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KubernetesProfilesAPI.KubernetesprofilesLockManager(context.Background()).KubernetesProfilesLockManagerCommand(kubernetesProfilesLockManagerCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesProfilesAPI.KubernetesprofilesLockManager``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesprofilesLockManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kubernetesProfilesLockManagerCommand** | [**KubernetesProfilesLockManagerCommand**](KubernetesProfilesLockManagerCommand.md) |  | 

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

