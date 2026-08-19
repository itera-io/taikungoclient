# \PackageAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PackageDetails**](PackageAPI.md#PackageDetails) | **Get** /api/v1/package/{repoName}/{packageName} | Available helm package details
[**PackageList**](PackageAPI.md#PackageList) | **Get** /api/v1/package/list | Retrieve all available packages
[**PackageValue**](PackageAPI.md#PackageValue) | **Post** /api/v1/package/value | Get yaml based value
[**PackageValueAutocomplete**](PackageAPI.md#PackageValueAutocomplete) | **Post** /api/v1/package/value-autocomplete | Get autocomplete dictionary
[**PackageVersions**](PackageAPI.md#PackageVersions) | **Post** /api/v1/package/versions | Get available versions



## PackageDetails

> AvailablePackageDetailsDto PackageDetails(ctx, repoName, packageName).OrganizationId(organizationId).Execute()

Available helm package details

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
	repoName := "repoName_example" // string | 
	packageName := "packageName_example" // string | 
	organizationId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PackageAPI.PackageDetails(context.Background(), repoName, packageName).OrganizationId(organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PackageAPI.PackageDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PackageDetails`: AvailablePackageDetailsDto
	fmt.Fprintf(os.Stdout, "Response from `PackageAPI.PackageDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoName** | **string** |  | 
**packageName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackageDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **organizationId** | **int32** |  | 

### Return type

[**AvailablePackageDetailsDto**](AvailablePackageDetailsDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackageList

> AvailablePackagesList PackageList(ctx).Offset(offset).Limit(limit).SortBy(sortBy).SortDirection(sortDirection).Search(search).Id(id).CatalogId(catalogId).IsPrivate(isPrivate).FilterBy(filterBy).OrganizationId(organizationId).Execute()

Retrieve all available packages

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
	id := "id_example" // string |  (optional)
	catalogId := int32(56) // int32 |  (optional)
	isPrivate := true // bool |  (optional)
	filterBy := "filterBy_example" // string |  (optional)
	organizationId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PackageAPI.PackageList(context.Background()).Offset(offset).Limit(limit).SortBy(sortBy).SortDirection(sortDirection).Search(search).Id(id).CatalogId(catalogId).IsPrivate(isPrivate).FilterBy(filterBy).OrganizationId(organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PackageAPI.PackageList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PackageList`: AvailablePackagesList
	fmt.Fprintf(os.Stdout, "Response from `PackageAPI.PackageList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPackageListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **id** | **string** |  | 
 **catalogId** | **int32** |  | 
 **isPrivate** | **bool** |  | 
 **filterBy** | **string** |  | 
 **organizationId** | **int32** |  | 

### Return type

[**AvailablePackagesList**](AvailablePackagesList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackageValue

> string PackageValue(ctx).GetCatalogAppValueCommand(getCatalogAppValueCommand).Execute()

Get yaml based value

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
	getCatalogAppValueCommand := *openapiclient.NewGetCatalogAppValueCommand("PackageId_example", "Version_example") // GetCatalogAppValueCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PackageAPI.PackageValue(context.Background()).GetCatalogAppValueCommand(getCatalogAppValueCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PackageAPI.PackageValue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PackageValue`: string
	fmt.Fprintf(os.Stdout, "Response from `PackageAPI.PackageValue`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPackageValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getCatalogAppValueCommand** | [**GetCatalogAppValueCommand**](GetCatalogAppValueCommand.md) |  | 

### Return type

**string**

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackageValueAutocomplete

> []PackageAutocompleteDto PackageValueAutocomplete(ctx).GetCatalogAppValueAutocompleteCommand(getCatalogAppValueAutocompleteCommand).Execute()

Get autocomplete dictionary

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
	getCatalogAppValueAutocompleteCommand := *openapiclient.NewGetCatalogAppValueAutocompleteCommand("PackageId_example", "Version_example", NullableInt32(123), false) // GetCatalogAppValueAutocompleteCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PackageAPI.PackageValueAutocomplete(context.Background()).GetCatalogAppValueAutocompleteCommand(getCatalogAppValueAutocompleteCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PackageAPI.PackageValueAutocomplete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PackageValueAutocomplete`: []PackageAutocompleteDto
	fmt.Fprintf(os.Stdout, "Response from `PackageAPI.PackageValueAutocomplete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPackageValueAutocompleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getCatalogAppValueAutocompleteCommand** | [**GetCatalogAppValueAutocompleteCommand**](GetCatalogAppValueAutocompleteCommand.md) |  | 

### Return type

[**[]PackageAutocompleteDto**](PackageAutocompleteDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackageVersions

> []string PackageVersions(ctx).ListCatalogAppAvailableVersionsCommand(listCatalogAppAvailableVersionsCommand).Execute()

Get available versions

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
	listCatalogAppAvailableVersionsCommand := *openapiclient.NewListCatalogAppAvailableVersionsCommand("RepoName_example", "PackageName_example", "CurrentVersion_example") // ListCatalogAppAvailableVersionsCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PackageAPI.PackageVersions(context.Background()).ListCatalogAppAvailableVersionsCommand(listCatalogAppAvailableVersionsCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PackageAPI.PackageVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PackageVersions`: []string
	fmt.Fprintf(os.Stdout, "Response from `PackageAPI.PackageVersions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPackageVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listCatalogAppAvailableVersionsCommand** | [**ListCatalogAppAvailableVersionsCommand**](ListCatalogAppAvailableVersionsCommand.md) |  | 

### Return type

**[]string**

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

