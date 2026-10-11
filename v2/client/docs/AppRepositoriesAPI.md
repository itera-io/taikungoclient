# \AppRepositoriesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RepositoryAvailableList**](AppRepositoriesAPI.md#RepositoryAvailableList) | **Get** /api/v1/repository/available | Retrieve available repositories
[**RepositoryBind**](AppRepositoriesAPI.md#RepositoryBind) | **Post** /api/v1/repository/bind | Bind repo to organization
[**RepositoryDelete**](AppRepositoriesAPI.md#RepositoryDelete) | **Post** /api/v1/repository/delete | Unbind repo from organization
[**RepositoryImport**](AppRepositoriesAPI.md#RepositoryImport) | **Post** /api/v1/repository/import | Import repo to artifact
[**RepositoryList**](AppRepositoriesAPI.md#RepositoryList) | **Get** /api/v1/repository/list | Retrieve repo names
[**RepositoryRecommendedList**](AppRepositoriesAPI.md#RepositoryRecommendedList) | **Get** /api/v1/repository/recommended | Retrieve taikun recommended repositories
[**RepositoryUnbind**](AppRepositoriesAPI.md#RepositoryUnbind) | **Post** /api/v1/repository/unbind | Unbind repo from organization
[**RepositoryUpdatePassword**](AppRepositoriesAPI.md#RepositoryUpdatePassword) | **Post** /api/v1/repository/update-password | Update repo password



## RepositoryAvailableList

> AppRepositoryList RepositoryAvailableList(ctx).Offset(offset).Limit(limit).SortBy(sortBy).SortDirection(sortDirection).Search(search).Id(id).IsPrivate(isPrivate).OrganizationId(organizationId).Execute()

Retrieve available repositories

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
	isPrivate := true // bool |  (optional)
	organizationId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppRepositoriesAPI.RepositoryAvailableList(context.Background()).Offset(offset).Limit(limit).SortBy(sortBy).SortDirection(sortDirection).Search(search).Id(id).IsPrivate(isPrivate).OrganizationId(organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppRepositoriesAPI.RepositoryAvailableList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoryAvailableList`: AppRepositoryList
	fmt.Fprintf(os.Stdout, "Response from `AppRepositoriesAPI.RepositoryAvailableList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRepositoryAvailableListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **id** | **string** |  | 
 **isPrivate** | **bool** |  | 
 **organizationId** | **int32** |  | 

### Return type

[**AppRepositoryList**](AppRepositoryList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoryBind

> RepositoryBind(ctx).BindAppRepositoryCommand(bindAppRepositoryCommand).Execute()

Bind repo to organization

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
	bindAppRepositoryCommand := *openapiclient.NewBindAppRepositoryCommand([]openapiclient.FilteringElementDto{*openapiclient.NewFilteringElementDto("OrganizationName_example", "Name_example")}, NullableInt32(123)) // BindAppRepositoryCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppRepositoriesAPI.RepositoryBind(context.Background()).BindAppRepositoryCommand(bindAppRepositoryCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppRepositoriesAPI.RepositoryBind``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRepositoryBindRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bindAppRepositoryCommand** | [**BindAppRepositoryCommand**](BindAppRepositoryCommand.md) |  | 

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


## RepositoryDelete

> RepositoryDelete(ctx).DeleteRepositoryCommand(deleteRepositoryCommand).Execute()

Unbind repo from organization

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
	deleteRepositoryCommand := *openapiclient.NewDeleteRepositoryCommand(int32(123)) // DeleteRepositoryCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppRepositoriesAPI.RepositoryDelete(context.Background()).DeleteRepositoryCommand(deleteRepositoryCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppRepositoriesAPI.RepositoryDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRepositoryDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteRepositoryCommand** | [**DeleteRepositoryCommand**](DeleteRepositoryCommand.md) |  | 

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


## RepositoryImport

> RepositoryImport(ctx).ImportRepoCommand(importRepoCommand).Execute()

Import repo to artifact

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
	importRepoCommand := *openapiclient.NewImportRepoCommand("Username_example", "Password_example", "Name_example", "Url_example", int32(123)) // ImportRepoCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppRepositoriesAPI.RepositoryImport(context.Background()).ImportRepoCommand(importRepoCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppRepositoriesAPI.RepositoryImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRepositoryImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importRepoCommand** | [**ImportRepoCommand**](ImportRepoCommand.md) |  | 

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


## RepositoryList

> []string RepositoryList(ctx).OrganizationId(organizationId).Execute()

Retrieve repo names

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppRepositoriesAPI.RepositoryList(context.Background()).OrganizationId(organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppRepositoriesAPI.RepositoryList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoryList`: []string
	fmt.Fprintf(os.Stdout, "Response from `AppRepositoriesAPI.RepositoryList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRepositoryListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 

### Return type

**[]string**

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoryRecommendedList

> []ArtifactRepositoryDto RepositoryRecommendedList(ctx).IsTaikun(isTaikun).OrganizationId(organizationId).Execute()

Retrieve taikun recommended repositories

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
	isTaikun := true // bool |  (optional)
	organizationId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppRepositoriesAPI.RepositoryRecommendedList(context.Background()).IsTaikun(isTaikun).OrganizationId(organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppRepositoriesAPI.RepositoryRecommendedList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoryRecommendedList`: []ArtifactRepositoryDto
	fmt.Fprintf(os.Stdout, "Response from `AppRepositoriesAPI.RepositoryRecommendedList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRepositoryRecommendedListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isTaikun** | **bool** |  | 
 **organizationId** | **int32** |  | 

### Return type

[**[]ArtifactRepositoryDto**](ArtifactRepositoryDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoryUnbind

> RepositoryUnbind(ctx).UnbindAppRepositoryCommand(unbindAppRepositoryCommand).Execute()

Unbind repo from organization

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
	unbindAppRepositoryCommand := *openapiclient.NewUnbindAppRepositoryCommand([]string{"Ids_example"}, int32(123)) // UnbindAppRepositoryCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppRepositoriesAPI.RepositoryUnbind(context.Background()).UnbindAppRepositoryCommand(unbindAppRepositoryCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppRepositoriesAPI.RepositoryUnbind``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRepositoryUnbindRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unbindAppRepositoryCommand** | [**UnbindAppRepositoryCommand**](UnbindAppRepositoryCommand.md) |  | 

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


## RepositoryUpdatePassword

> RepositoryUpdatePassword(ctx).UpdateRepoPasswordCommand(updateRepoPasswordCommand).Execute()

Update repo password

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
	updateRepoPasswordCommand := *openapiclient.NewUpdateRepoPasswordCommand("Username_example", "Password_example", "RepositoryId_example", int32(123)) // UpdateRepoPasswordCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppRepositoriesAPI.RepositoryUpdatePassword(context.Background()).UpdateRepoPasswordCommand(updateRepoPasswordCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppRepositoriesAPI.RepositoryUpdatePassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRepositoryUpdatePasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateRepoPasswordCommand** | [**UpdateRepoPasswordCommand**](UpdateRepoPasswordCommand.md) |  | 

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

