# \NtpServersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NtpserversCreate**](NtpServersAPI.md#NtpserversCreate) | **Post** /api/v1/ntpservers/create | Create access profile ntp server
[**NtpserversDelete**](NtpServersAPI.md#NtpserversDelete) | **Delete** /api/v1/ntpservers/{id} | Delete access profile ntp server
[**NtpserversEdit**](NtpServersAPI.md#NtpserversEdit) | **Put** /api/v1/ntpservers/edit/{id} | Edit access profile ntp server
[**NtpserversList**](NtpServersAPI.md#NtpserversList) | **Get** /api/v1/ntpservers/list/{accessProfileId} | List ntp server by access profile id



## NtpserversCreate

> ApiResponse NtpserversCreate(ctx).CreateNtpServerCommand(createNtpServerCommand).Execute()

Create access profile ntp server

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
	createNtpServerCommand := *openapiclient.NewCreateNtpServerCommand("Address_example", int32(123)) // CreateNtpServerCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NtpServersAPI.NtpserversCreate(context.Background()).CreateNtpServerCommand(createNtpServerCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NtpServersAPI.NtpserversCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NtpserversCreate`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `NtpServersAPI.NtpserversCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNtpserversCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNtpServerCommand** | [**CreateNtpServerCommand**](CreateNtpServerCommand.md) |  | 

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


## NtpserversDelete

> NtpserversDelete(ctx, id).Execute()

Delete access profile ntp server

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
	r, err := apiClient.NtpServersAPI.NtpserversDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NtpServersAPI.NtpserversDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiNtpserversDeleteRequest struct via the builder pattern


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


## NtpserversEdit

> NtpserversEdit(ctx, id).DnsNtpAddressEditDto(dnsNtpAddressEditDto).Execute()

Edit access profile ntp server

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
	dnsNtpAddressEditDto := *openapiclient.NewDnsNtpAddressEditDto("Address_example") // DnsNtpAddressEditDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NtpServersAPI.NtpserversEdit(context.Background(), id).DnsNtpAddressEditDto(dnsNtpAddressEditDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NtpServersAPI.NtpserversEdit``: %v\n", err)
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

Other parameters are passed through a pointer to a apiNtpserversEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dnsNtpAddressEditDto** | [**DnsNtpAddressEditDto**](DnsNtpAddressEditDto.md) |  | 

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


## NtpserversList

> []NtpServersListDto NtpserversList(ctx, accessProfileId).Search(search).Execute()

List ntp server by access profile id

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
	resp, r, err := apiClient.NtpServersAPI.NtpserversList(context.Background(), accessProfileId).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NtpServersAPI.NtpserversList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NtpserversList`: []NtpServersListDto
	fmt.Fprintf(os.Stdout, "Response from `NtpServersAPI.NtpserversList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessProfileId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNtpserversListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** |  | 

### Return type

[**[]NtpServersListDto**](NtpServersListDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

