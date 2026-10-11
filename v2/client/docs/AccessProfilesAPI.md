# \AccessProfilesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccessprofilesCreate**](AccessProfilesAPI.md#AccessprofilesCreate) | **Post** /api/v1/accessprofiles/create | Create access profile
[**AccessprofilesDelete**](AccessProfilesAPI.md#AccessprofilesDelete) | **Delete** /api/v1/accessprofiles/{id} | Delete access profile by Id
[**AccessprofilesDropdown**](AccessProfilesAPI.md#AccessprofilesDropdown) | **Get** /api/v1/accessprofiles/list | Retrieve access profiles by organization Id
[**AccessprofilesList**](AccessProfilesAPI.md#AccessprofilesList) | **Get** /api/v1/accessprofiles | Retrieve all access profiles
[**AccessprofilesLockManager**](AccessProfilesAPI.md#AccessprofilesLockManager) | **Post** /api/v1/accessprofiles/lockmanager | Lock/unlock access profiles
[**AccessprofilesUpdate**](AccessProfilesAPI.md#AccessprofilesUpdate) | **Put** /api/v1/accessprofiles/update/{id} | Update access profile



## AccessprofilesCreate

> ApiResponse AccessprofilesCreate(ctx).CreateAccessProfileCommand(createAccessProfileCommand).Execute()

Create access profile

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
	createAccessProfileCommand := *openapiclient.NewCreateAccessProfileCommand() // CreateAccessProfileCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessProfilesAPI.AccessprofilesCreate(context.Background()).CreateAccessProfileCommand(createAccessProfileCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessProfilesAPI.AccessprofilesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccessprofilesCreate`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AccessProfilesAPI.AccessprofilesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccessprofilesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAccessProfileCommand** | [**CreateAccessProfileCommand**](CreateAccessProfileCommand.md) |  | 

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


## AccessprofilesDelete

> AccessprofilesDelete(ctx, id).Execute()

Delete access profile by Id

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
	r, err := apiClient.AccessProfilesAPI.AccessprofilesDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessProfilesAPI.AccessprofilesDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAccessprofilesDeleteRequest struct via the builder pattern


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


## AccessprofilesDropdown

> []CommonDropdownDto AccessprofilesDropdown(ctx).OrganizationId(organizationId).Search(search).Offset(offset).Limit(limit).Execute()

Retrieve access profiles by organization Id

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
	offset := int32(56) // int32 |  (optional) (default to 0)
	limit := int32(56) // int32 |  (optional) (default to 1000)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessProfilesAPI.AccessprofilesDropdown(context.Background()).OrganizationId(organizationId).Search(search).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessProfilesAPI.AccessprofilesDropdown``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccessprofilesDropdown`: []CommonDropdownDto
	fmt.Fprintf(os.Stdout, "Response from `AccessProfilesAPI.AccessprofilesDropdown`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccessprofilesDropdownRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 
 **search** | **string** |  | 
 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 1000]

### Return type

[**[]CommonDropdownDto**](CommonDropdownDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccessprofilesList

> AccessProfilesList AccessprofilesList(ctx).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Id(id).DomainId(domainId).Offset(offset).Limit(limit).Execute()

Retrieve all access profiles

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
	organizationId := int32(56) // int32 | Id should be positive integer (optional)
	sortBy := "sortBy_example" // string | Sort By: ( name ) (optional)
	sortDirection := "sortDirection_example" // string | asc|desc (optional)
	search := "search_example" // string | Search related data (optional)
	searchId := "searchId_example" // string | Entity id (optional)
	id := int32(56) // int32 | Id should be positive integer (optional)
	domainId := int32(56) // int32 | Id should be positive integer (optional)
	offset := int32(56) // int32 | By default value is 0 (optional) (default to 0)
	limit := int32(56) // int32 | By default value is 50 (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessProfilesAPI.AccessprofilesList(context.Background()).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Id(id).DomainId(domainId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessProfilesAPI.AccessprofilesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccessprofilesList`: AccessProfilesList
	fmt.Fprintf(os.Stdout, "Response from `AccessProfilesAPI.AccessprofilesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccessprofilesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** | Id should be positive integer | 
 **sortBy** | **string** | Sort By: ( name ) | 
 **sortDirection** | **string** | asc|desc | 
 **search** | **string** | Search related data | 
 **searchId** | **string** | Entity id | 
 **id** | **int32** | Id should be positive integer | 
 **domainId** | **int32** | Id should be positive integer | 
 **offset** | **int32** | By default value is 0 | [default to 0]
 **limit** | **int32** | By default value is 50 | [default to 50]

### Return type

[**AccessProfilesList**](AccessProfilesList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccessprofilesLockManager

> AccessprofilesLockManager(ctx).AccessProfilesLockManagementCommand(accessProfilesLockManagementCommand).Execute()

Lock/unlock access profiles

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
	accessProfilesLockManagementCommand := *openapiclient.NewAccessProfilesLockManagementCommand(int32(123), "Mode_example") // AccessProfilesLockManagementCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccessProfilesAPI.AccessprofilesLockManager(context.Background()).AccessProfilesLockManagementCommand(accessProfilesLockManagementCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessProfilesAPI.AccessprofilesLockManager``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccessprofilesLockManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessProfilesLockManagementCommand** | [**AccessProfilesLockManagementCommand**](AccessProfilesLockManagementCommand.md) |  | 

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


## AccessprofilesUpdate

> AccessprofilesUpdate(ctx, id).UpdateAccessProfileDto(updateAccessProfileDto).Execute()

Update access profile

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
	updateAccessProfileDto := *openapiclient.NewUpdateAccessProfileDto() // UpdateAccessProfileDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccessProfilesAPI.AccessprofilesUpdate(context.Background(), id).UpdateAccessProfileDto(updateAccessProfileDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessProfilesAPI.AccessprofilesUpdate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAccessprofilesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAccessProfileDto** | [**UpdateAccessProfileDto**](UpdateAccessProfileDto.md) |  | 

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

