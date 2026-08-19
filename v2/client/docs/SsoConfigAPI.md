# \SsoConfigAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SsoConfigCheck**](SsoConfigAPI.md#SsoConfigCheck) | **Post** /api/v1/sso-config/check | Verify that an SSO configuration&#39;s issuer URL is reachable and returns a valid OIDC discovery document
[**SsoConfigCreate**](SsoConfigAPI.md#SsoConfigCreate) | **Post** /api/v1/sso-config/create | Create SSO configuration
[**SsoConfigDelete**](SsoConfigAPI.md#SsoConfigDelete) | **Delete** /api/v1/sso-config/{id} | Delete SSO configuration
[**SsoConfigDetails**](SsoConfigAPI.md#SsoConfigDetails) | **Get** /api/v1/sso-config/details | Retrieve a single SSO configuration by id
[**SsoConfigList**](SsoConfigAPI.md#SsoConfigList) | **Get** /api/v1/sso-config | Retrieve all SSO configurations for an domain
[**SsoConfigUpdate**](SsoConfigAPI.md#SsoConfigUpdate) | **Put** /api/v1/sso-config/update/{id} | Update SSO configuration



## SsoConfigCheck

> SsoConfigCheck(ctx).CheckDomainSsoConfigCommand(checkDomainSsoConfigCommand).Execute()

Verify that an SSO configuration's issuer URL is reachable and returns a valid OIDC discovery document

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
	checkDomainSsoConfigCommand := *openapiclient.NewCheckDomainSsoConfigCommand("IssuerUrl_example", "ClientId_example", "ClientSecret_example") // CheckDomainSsoConfigCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SsoConfigAPI.SsoConfigCheck(context.Background()).CheckDomainSsoConfigCommand(checkDomainSsoConfigCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SsoConfigAPI.SsoConfigCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSsoConfigCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkDomainSsoConfigCommand** | [**CheckDomainSsoConfigCommand**](CheckDomainSsoConfigCommand.md) |  | 

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


## SsoConfigCreate

> int32 SsoConfigCreate(ctx).CreateDomainSsoConfigCommand(createDomainSsoConfigCommand).Execute()

Create SSO configuration

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
	createDomainSsoConfigCommand := *openapiclient.NewCreateDomainSsoConfigCommand("Name_example", int32(123), "IssuerUrl_example", "ClientId_example", "ClientSecret_example", false, false) // CreateDomainSsoConfigCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SsoConfigAPI.SsoConfigCreate(context.Background()).CreateDomainSsoConfigCommand(createDomainSsoConfigCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SsoConfigAPI.SsoConfigCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SsoConfigCreate`: int32
	fmt.Fprintf(os.Stdout, "Response from `SsoConfigAPI.SsoConfigCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSsoConfigCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDomainSsoConfigCommand** | [**CreateDomainSsoConfigCommand**](CreateDomainSsoConfigCommand.md) |  | 

### Return type

**int32**

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SsoConfigDelete

> SsoConfigDelete(ctx, id).Execute()

Delete SSO configuration

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
	r, err := apiClient.SsoConfigAPI.SsoConfigDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SsoConfigAPI.SsoConfigDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSsoConfigDeleteRequest struct via the builder pattern


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


## SsoConfigDetails

> DomainSsoConfigDto SsoConfigDetails(ctx).Id(id).Execute()

Retrieve a single SSO configuration by id

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
	resp, r, err := apiClient.SsoConfigAPI.SsoConfigDetails(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SsoConfigAPI.SsoConfigDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SsoConfigDetails`: DomainSsoConfigDto
	fmt.Fprintf(os.Stdout, "Response from `SsoConfigAPI.SsoConfigDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSsoConfigDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32** |  | 

### Return type

[**DomainSsoConfigDto**](DomainSsoConfigDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SsoConfigList

> DomainSsoConfigList SsoConfigList(ctx).DomainId(domainId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Id(id).Execute()

Retrieve all SSO configurations for an domain

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
	domainId := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 |  (optional)
	offset := int32(56) // int32 |  (optional)
	sortBy := "sortBy_example" // string |  (optional)
	sortDirection := "sortDirection_example" // string |  (optional)
	search := "search_example" // string |  (optional)
	id := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SsoConfigAPI.SsoConfigList(context.Background()).DomainId(domainId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SsoConfigAPI.SsoConfigList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SsoConfigList`: DomainSsoConfigList
	fmt.Fprintf(os.Stdout, "Response from `SsoConfigAPI.SsoConfigList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSsoConfigListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainId** | **int32** |  | 
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **id** | **int32** |  | 

### Return type

[**DomainSsoConfigList**](DomainSsoConfigList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SsoConfigUpdate

> SsoConfigUpdate(ctx, id).UpdateDomainSsoConfigDto(updateDomainSsoConfigDto).Execute()

Update SSO configuration

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
	updateDomainSsoConfigDto := *openapiclient.NewUpdateDomainSsoConfigDto() // UpdateDomainSsoConfigDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SsoConfigAPI.SsoConfigUpdate(context.Background(), id).UpdateDomainSsoConfigDto(updateDomainSsoConfigDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SsoConfigAPI.SsoConfigUpdate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSsoConfigUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDomainSsoConfigDto** | [**UpdateDomainSsoConfigDto**](UpdateDomainSsoConfigDto.md) |  | 

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

