# \SlackAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SlackCreate**](SlackAPI.md#SlackCreate) | **Post** /api/v1/slack/create | Create slack configuration
[**SlackDeleteMultiple**](SlackAPI.md#SlackDeleteMultiple) | **Post** /api/v1/slack/delete-multiple | Delete slack configuration(s)
[**SlackDropdown**](SlackAPI.md#SlackDropdown) | **Get** /api/v1/slack/list | Retrieve all slack configs for organization
[**SlackList**](SlackAPI.md#SlackList) | **Get** /api/v1/slack | Retrieve all slack configs
[**SlackUpdate**](SlackAPI.md#SlackUpdate) | **Put** /api/v1/slack/update/{id} | Update slack configuration
[**SlackVerify**](SlackAPI.md#SlackVerify) | **Post** /api/v1/slack/verify | Verify slack configuration



## SlackCreate

> ApiResponse SlackCreate(ctx).CreateSlackConfigurationCommand(createSlackConfigurationCommand).Execute()

Create slack configuration

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
	createSlackConfigurationCommand := *openapiclient.NewCreateSlackConfigurationCommand(NullableInt32(123), "Name_example", "Url_example", "Channel_example", openapiclient.SlackType("Alert")) // CreateSlackConfigurationCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlackAPI.SlackCreate(context.Background()).CreateSlackConfigurationCommand(createSlackConfigurationCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlackAPI.SlackCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlackCreate`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `SlackAPI.SlackCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlackCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSlackConfigurationCommand** | [**CreateSlackConfigurationCommand**](CreateSlackConfigurationCommand.md) |  | 

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


## SlackDeleteMultiple

> SlackDeleteMultiple(ctx).DeleteSlackConfigCommand(deleteSlackConfigCommand).Execute()

Delete slack configuration(s)

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
	deleteSlackConfigCommand := *openapiclient.NewDeleteSlackConfigCommand([]int32{int32(123)}) // DeleteSlackConfigCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SlackAPI.SlackDeleteMultiple(context.Background()).DeleteSlackConfigCommand(deleteSlackConfigCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlackAPI.SlackDeleteMultiple``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlackDeleteMultipleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteSlackConfigCommand** | [**DeleteSlackConfigCommand**](DeleteSlackConfigCommand.md) |  | 

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


## SlackDropdown

> []CommonDropdownDto SlackDropdown(ctx).OrganizationId(organizationId).Search(search).Execute()

Retrieve all slack configs for organization

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
	resp, r, err := apiClient.SlackAPI.SlackDropdown(context.Background()).OrganizationId(organizationId).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlackAPI.SlackDropdown``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlackDropdown`: []CommonDropdownDto
	fmt.Fprintf(os.Stdout, "Response from `SlackAPI.SlackDropdown`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlackDropdownRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 
 **search** | **string** |  | 

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


## SlackList

> SlackConfigurationList SlackList(ctx).OrganizationId(organizationId).DomainId(domainId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Id(id).Execute()

Retrieve all slack configs

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
	id := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlackAPI.SlackList(context.Background()).OrganizationId(organizationId).DomainId(domainId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlackAPI.SlackList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlackList`: SlackConfigurationList
	fmt.Fprintf(os.Stdout, "Response from `SlackAPI.SlackList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlackListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 
 **domainId** | **int32** |  | 
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **id** | **int32** |  | 

### Return type

[**SlackConfigurationList**](SlackConfigurationList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlackUpdate

> SlackUpdate(ctx, id).UpdateSlackConfigurationDto(updateSlackConfigurationDto).Execute()

Update slack configuration

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
	updateSlackConfigurationDto := *openapiclient.NewUpdateSlackConfigurationDto() // UpdateSlackConfigurationDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SlackAPI.SlackUpdate(context.Background(), id).UpdateSlackConfigurationDto(updateSlackConfigurationDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlackAPI.SlackUpdate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSlackUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateSlackConfigurationDto** | [**UpdateSlackConfigurationDto**](UpdateSlackConfigurationDto.md) |  | 

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


## SlackVerify

> SlackVerify(ctx).VerifySlackCredentialsCommand(verifySlackCredentialsCommand).Execute()

Verify slack configuration

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
	verifySlackCredentialsCommand := *openapiclient.NewVerifySlackCredentialsCommand("Name_example", "Url_example", "Channel_example") // VerifySlackCredentialsCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SlackAPI.SlackVerify(context.Background()).VerifySlackCredentialsCommand(verifySlackCredentialsCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlackAPI.SlackVerify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlackVerifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **verifySlackCredentialsCommand** | [**VerifySlackCredentialsCommand**](VerifySlackCredentialsCommand.md) |  | 

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

