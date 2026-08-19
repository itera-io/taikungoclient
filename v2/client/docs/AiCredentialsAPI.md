# \AiCredentialsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AiCredentialCreate**](AiCredentialsAPI.md#AiCredentialCreate) | **Post** /api/v1/ai-credential/create | Create ai credential
[**AiCredentialDelete**](AiCredentialsAPI.md#AiCredentialDelete) | **Delete** /api/v1/ai-credential/{id} | Remove ai credential
[**AiCredentialDropdown**](AiCredentialsAPI.md#AiCredentialDropdown) | **Get** /api/v1/ai-credential | Retrieve all AI credentials for organization
[**AiCredentialList**](AiCredentialsAPI.md#AiCredentialList) | **Get** /api/v1/ai-credential/list | Retrieve all AI credentials



## AiCredentialCreate

> ApiResponse AiCredentialCreate(ctx).CreateAiCredentialCommand(createAiCredentialCommand).Execute()

Create ai credential

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
	createAiCredentialCommand := *openapiclient.NewCreateAiCredentialCommand("Name_example", "ApiKey_example", NullableInt32(123)) // CreateAiCredentialCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiCredentialsAPI.AiCredentialCreate(context.Background()).CreateAiCredentialCommand(createAiCredentialCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiCredentialsAPI.AiCredentialCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiCredentialCreate`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AiCredentialsAPI.AiCredentialCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiCredentialCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAiCredentialCommand** | [**CreateAiCredentialCommand**](CreateAiCredentialCommand.md) |  | 

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


## AiCredentialDelete

> AiCredentialDelete(ctx, id).Execute()

Remove ai credential

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
	r, err := apiClient.AiCredentialsAPI.AiCredentialDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiCredentialsAPI.AiCredentialDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAiCredentialDeleteRequest struct via the builder pattern


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


## AiCredentialDropdown

> []AiCredentialsForOrganizationEntity AiCredentialDropdown(ctx).OrganizationId(organizationId).Search(search).Execute()

Retrieve all AI credentials for organization

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiCredentialsAPI.AiCredentialDropdown(context.Background()).OrganizationId(organizationId).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiCredentialsAPI.AiCredentialDropdown``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiCredentialDropdown`: []AiCredentialsForOrganizationEntity
	fmt.Fprintf(os.Stdout, "Response from `AiCredentialsAPI.AiCredentialDropdown`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiCredentialDropdownRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 
 **search** | **string** |  | 

### Return type

[**[]AiCredentialsForOrganizationEntity**](AiCredentialsForOrganizationEntity.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiCredentialList

> AiCredentials AiCredentialList(ctx).OrganizationId(organizationId).Search(search).SearchId(searchId).Id(id).SortBy(sortBy).DomainId(domainId).SortDirection(sortDirection).Limit(limit).Offset(offset).Execute()

Retrieve all AI credentials

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
	searchId := "searchId_example" // string |  (optional)
	id := int32(56) // int32 |  (optional)
	sortBy := "sortBy_example" // string |  (optional)
	domainId := int32(56) // int32 |  (optional)
	sortDirection := "sortDirection_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)
	offset := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiCredentialsAPI.AiCredentialList(context.Background()).OrganizationId(organizationId).Search(search).SearchId(searchId).Id(id).SortBy(sortBy).DomainId(domainId).SortDirection(sortDirection).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiCredentialsAPI.AiCredentialList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiCredentialList`: AiCredentials
	fmt.Fprintf(os.Stdout, "Response from `AiCredentialsAPI.AiCredentialList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiCredentialListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **id** | **int32** |  | 
 **sortBy** | **string** |  | 
 **domainId** | **int32** |  | 
 **sortDirection** | **string** |  | 
 **limit** | **int32** |  | [default to 50]
 **offset** | **int32** |  | [default to 0]

### Return type

[**AiCredentials**](AiCredentials.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

