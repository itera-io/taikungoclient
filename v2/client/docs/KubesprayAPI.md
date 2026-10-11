# \KubesprayAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KubesprayCreate**](KubesprayAPI.md#KubesprayCreate) | **Post** /api/v1/kubespray | Add kubespray
[**KubesprayDelete**](KubesprayAPI.md#KubesprayDelete) | **Delete** /api/v1/kubespray/{id} | Delete kubespray
[**KubesprayList**](KubesprayAPI.md#KubesprayList) | **Get** /api/v1/kubespray/list | Retrieve all kubespray versions



## KubesprayCreate

> ApiResponse KubesprayCreate(ctx).KubesprayCreateCommand(kubesprayCreateCommand).Execute()

Add kubespray

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
	kubesprayCreateCommand := *openapiclient.NewKubesprayCreateCommand("Version_example", "KubernetesVersion_example") // KubesprayCreateCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubesprayAPI.KubesprayCreate(context.Background()).KubesprayCreateCommand(kubesprayCreateCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubesprayAPI.KubesprayCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubesprayCreate`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `KubesprayAPI.KubesprayCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubesprayCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kubesprayCreateCommand** | [**KubesprayCreateCommand**](KubesprayCreateCommand.md) |  | 

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


## KubesprayDelete

> KubesprayDelete(ctx, id).Execute()

Delete kubespray

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
	r, err := apiClient.KubesprayAPI.KubesprayDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubesprayAPI.KubesprayDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiKubesprayDeleteRequest struct via the builder pattern


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


## KubesprayList

> Kubesprays KubesprayList(ctx).Execute()

Retrieve all kubespray versions

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubesprayAPI.KubesprayList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubesprayAPI.KubesprayList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubesprayList`: Kubesprays
	fmt.Fprintf(os.Stdout, "Response from `KubesprayAPI.KubesprayList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiKubesprayListRequest struct via the builder pattern


### Return type

[**Kubesprays**](Kubesprays.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

