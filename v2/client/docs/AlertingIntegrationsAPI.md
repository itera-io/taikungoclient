# \AlertingIntegrationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AlertingintegrationsCreate**](AlertingIntegrationsAPI.md#AlertingintegrationsCreate) | **Post** /api/v1/alertingintegrations/create | Create alerting profile alerting integration
[**AlertingintegrationsDelete**](AlertingIntegrationsAPI.md#AlertingintegrationsDelete) | **Delete** /api/v1/alertingintegrations/{id} | Delete alerting profile alerting integration
[**AlertingintegrationsEdit**](AlertingIntegrationsAPI.md#AlertingintegrationsEdit) | **Put** /api/v1/alertingintegrations/edit | Edit alerting profile alerting integration
[**AlertingintegrationsList**](AlertingIntegrationsAPI.md#AlertingintegrationsList) | **Get** /api/v1/alertingintegrations/{alertingProfileId} | List alerting integrations by profile id



## AlertingintegrationsCreate

> ApiResponse AlertingintegrationsCreate(ctx).CreateAlertingIntegrationCommand(createAlertingIntegrationCommand).Execute()

Create alerting profile alerting integration

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
	createAlertingIntegrationCommand := *openapiclient.NewCreateAlertingIntegrationCommand("Url_example", "Token_example", openapiclient.AlertingIntegrationType("Opsgenie"), int32(123)) // CreateAlertingIntegrationCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertingIntegrationsAPI.AlertingintegrationsCreate(context.Background()).CreateAlertingIntegrationCommand(createAlertingIntegrationCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertingIntegrationsAPI.AlertingintegrationsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlertingintegrationsCreate`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AlertingIntegrationsAPI.AlertingintegrationsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertingintegrationsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAlertingIntegrationCommand** | [**CreateAlertingIntegrationCommand**](CreateAlertingIntegrationCommand.md) |  | 

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


## AlertingintegrationsDelete

> AlertingintegrationsDelete(ctx, id).Execute()

Delete alerting profile alerting integration

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
	r, err := apiClient.AlertingIntegrationsAPI.AlertingintegrationsDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertingIntegrationsAPI.AlertingintegrationsDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAlertingintegrationsDeleteRequest struct via the builder pattern


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


## AlertingintegrationsEdit

> AlertingintegrationsEdit(ctx).EditAlertingIntegrationCommand(editAlertingIntegrationCommand).Execute()

Edit alerting profile alerting integration

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
	editAlertingIntegrationCommand := *openapiclient.NewEditAlertingIntegrationCommand() // EditAlertingIntegrationCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AlertingIntegrationsAPI.AlertingintegrationsEdit(context.Background()).EditAlertingIntegrationCommand(editAlertingIntegrationCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertingIntegrationsAPI.AlertingintegrationsEdit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertingintegrationsEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **editAlertingIntegrationCommand** | [**EditAlertingIntegrationCommand**](EditAlertingIntegrationCommand.md) |  | 

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


## AlertingintegrationsList

> []AlertingIntegrationsListDto AlertingintegrationsList(ctx, alertingProfileId).Search(search).Execute()

List alerting integrations by profile id

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
	alertingProfileId := int32(56) // int32 | 
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertingIntegrationsAPI.AlertingintegrationsList(context.Background(), alertingProfileId).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertingIntegrationsAPI.AlertingintegrationsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlertingintegrationsList`: []AlertingIntegrationsListDto
	fmt.Fprintf(os.Stdout, "Response from `AlertingIntegrationsAPI.AlertingintegrationsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertingProfileId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertingintegrationsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** |  | 

### Return type

[**[]AlertingIntegrationsListDto**](AlertingIntegrationsListDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

