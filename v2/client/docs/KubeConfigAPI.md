# \KubeConfigAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KubeconfigCreate**](KubeConfigAPI.md#KubeconfigCreate) | **Post** /api/v1/kubeconfig | Create kube config
[**KubeconfigDelete**](KubeConfigAPI.md#KubeconfigDelete) | **Post** /api/v1/kubeconfig/delete | Delete kube config
[**KubeconfigDeleteByProjectId**](KubeConfigAPI.md#KubeconfigDeleteByProjectId) | **Post** /api/v1/kubeconfig/delete-by-project-id | Delete kube config by project id
[**KubeconfigDownload**](KubeConfigAPI.md#KubeconfigDownload) | **Post** /api/v1/kubeconfig/download | Download kube config file for user by project Id
[**KubeconfigExport**](KubeConfigAPI.md#KubeconfigExport) | **Post** /api/v1/kubeconfig/export | Export
[**KubeconfigInteractiveShell**](KubeConfigAPI.md#KubeconfigInteractiveShell) | **Post** /api/v1/kubeconfig/interactive-shell | Interactive shell for user kube config
[**KubeconfigList**](KubeConfigAPI.md#KubeconfigList) | **Get** /api/v1/kubeconfig | Retrieve a list of kube configs for project



## KubeconfigCreate

> ApiResponse KubeconfigCreate(ctx).CreateKubeConfigCommand(createKubeConfigCommand).Execute()

Create kube config

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
	createKubeConfigCommand := *openapiclient.NewCreateKubeConfigCommand() // CreateKubeConfigCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubeConfigAPI.KubeconfigCreate(context.Background()).CreateKubeConfigCommand(createKubeConfigCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubeConfigAPI.KubeconfigCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubeconfigCreate`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `KubeConfigAPI.KubeconfigCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubeconfigCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createKubeConfigCommand** | [**CreateKubeConfigCommand**](CreateKubeConfigCommand.md) |  | 

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


## KubeconfigDelete

> KubeconfigDelete(ctx).DeleteKubeConfigCommand(deleteKubeConfigCommand).Execute()

Delete kube config

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
	deleteKubeConfigCommand := *openapiclient.NewDeleteKubeConfigCommand(int32(123)) // DeleteKubeConfigCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KubeConfigAPI.KubeconfigDelete(context.Background()).DeleteKubeConfigCommand(deleteKubeConfigCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubeConfigAPI.KubeconfigDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubeconfigDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteKubeConfigCommand** | [**DeleteKubeConfigCommand**](DeleteKubeConfigCommand.md) |  | 

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


## KubeconfigDeleteByProjectId

> KubeconfigDeleteByProjectId(ctx).DeleteKubeConfigByProjectIdCommand(deleteKubeConfigByProjectIdCommand).Execute()

Delete kube config by project id

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
	deleteKubeConfigByProjectIdCommand := *openapiclient.NewDeleteKubeConfigByProjectIdCommand(int32(123)) // DeleteKubeConfigByProjectIdCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KubeConfigAPI.KubeconfigDeleteByProjectId(context.Background()).DeleteKubeConfigByProjectIdCommand(deleteKubeConfigByProjectIdCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubeConfigAPI.KubeconfigDeleteByProjectId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubeconfigDeleteByProjectIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteKubeConfigByProjectIdCommand** | [**DeleteKubeConfigByProjectIdCommand**](DeleteKubeConfigByProjectIdCommand.md) |  | 

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


## KubeconfigDownload

> string KubeconfigDownload(ctx).DownloadKubeConfigCommand(downloadKubeConfigCommand).Execute()

Download kube config file for user by project Id

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
	downloadKubeConfigCommand := *openapiclient.NewDownloadKubeConfigCommand(int32(123), int32(123)) // DownloadKubeConfigCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubeConfigAPI.KubeconfigDownload(context.Background()).DownloadKubeConfigCommand(downloadKubeConfigCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubeConfigAPI.KubeconfigDownload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubeconfigDownload`: string
	fmt.Fprintf(os.Stdout, "Response from `KubeConfigAPI.KubeconfigDownload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubeconfigDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **downloadKubeConfigCommand** | [**DownloadKubeConfigCommand**](DownloadKubeConfigCommand.md) |  | 

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


## KubeconfigExport

> CsvExporter KubeconfigExport(ctx).ExportKubeConfigCommand(exportKubeConfigCommand).Execute()

Export

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
	exportKubeConfigCommand := *openapiclient.NewExportKubeConfigCommand(int32(123), int32(123)) // ExportKubeConfigCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubeConfigAPI.KubeconfigExport(context.Background()).ExportKubeConfigCommand(exportKubeConfigCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubeConfigAPI.KubeconfigExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubeconfigExport`: CsvExporter
	fmt.Fprintf(os.Stdout, "Response from `KubeConfigAPI.KubeconfigExport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubeconfigExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **exportKubeConfigCommand** | [**ExportKubeConfigCommand**](ExportKubeConfigCommand.md) |  | 

### Return type

[**CsvExporter**](CsvExporter.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubeconfigInteractiveShell

> InteractiveShellDto KubeconfigInteractiveShell(ctx).KubeConfigInteractiveShellCommand(kubeConfigInteractiveShellCommand).Execute()

Interactive shell for user kube config

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
	kubeConfigInteractiveShellCommand := *openapiclient.NewKubeConfigInteractiveShellCommand(int32(123), "Token_example", int32(123)) // KubeConfigInteractiveShellCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubeConfigAPI.KubeconfigInteractiveShell(context.Background()).KubeConfigInteractiveShellCommand(kubeConfigInteractiveShellCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubeConfigAPI.KubeconfigInteractiveShell``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubeconfigInteractiveShell`: InteractiveShellDto
	fmt.Fprintf(os.Stdout, "Response from `KubeConfigAPI.KubeconfigInteractiveShell`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubeconfigInteractiveShellRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kubeConfigInteractiveShellCommand** | [**KubeConfigInteractiveShellCommand**](KubeConfigInteractiveShellCommand.md) |  | 

### Return type

[**InteractiveShellDto**](InteractiveShellDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubeconfigList

> KubeConfigForUserList KubeconfigList(ctx).ProjectId(projectId).OrganizationId(organizationId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Id(id).Execute()

Retrieve a list of kube configs for project

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
	projectId := int32(56) // int32 | 
	organizationId := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 |  (optional)
	offset := int32(56) // int32 |  (optional)
	sortBy := "sortBy_example" // string |  (optional)
	sortDirection := "sortDirection_example" // string |  (optional)
	search := "search_example" // string |  (optional)
	id := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubeConfigAPI.KubeconfigList(context.Background()).ProjectId(projectId).OrganizationId(organizationId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubeConfigAPI.KubeconfigList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubeconfigList`: KubeConfigForUserList
	fmt.Fprintf(os.Stdout, "Response from `KubeConfigAPI.KubeconfigList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubeconfigListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **int32** |  | 
 **organizationId** | **int32** |  | 
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **id** | **int32** |  | 

### Return type

[**KubeConfigForUserList**](KubeConfigForUserList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

