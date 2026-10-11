# \ServersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServersConsole**](ServersAPI.md#ServersConsole) | **Post** /api/v1/servers/console | Console screenshot or terminal access for server
[**ServersCreate**](ServersAPI.md#ServersCreate) | **Post** /api/v1/servers/create | Create a new server in the given project.
[**ServersDetails**](ServersAPI.md#ServersDetails) | **Get** /api/v1/servers/{projectId} | Retrieve all servers by given project
[**ServersGetCskClusterConfig**](ServersAPI.md#ServersGetCskClusterConfig) | **Get** /api/v1/servers/get-csk-cluster-config | Get csk cluster config
[**ServersList**](ServersAPI.md#ServersList) | **Get** /api/v1/servers | Retrieve all servers
[**ServersReboot**](ServersAPI.md#ServersReboot) | **Post** /api/v1/servers/reboot | Reboot server
[**ServersRegisterNode**](ServersAPI.md#ServersRegisterNode) | **Post** /api/v1/servers/register-node | Register csk node
[**ServersReset**](ServersAPI.md#ServersReset) | **Post** /api/v1/servers/reset | Update server(s) status(es)
[**ServersStatus**](ServersAPI.md#ServersStatus) | **Get** /api/v1/servers/status/{serverId} | Show server status
[**ServersUpdateByProjectId**](ServersAPI.md#ServersUpdateByProjectId) | **Put** /api/v1/servers/update/{projectId} | Update server by projectId



## ServersConsole

> string ServersConsole(ctx).ConsoleScreenshotCommand(consoleScreenshotCommand).Execute()

Console screenshot or terminal access for server

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
	consoleScreenshotCommand := *openapiclient.NewConsoleScreenshotCommand(int32(123)) // ConsoleScreenshotCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServersAPI.ServersConsole(context.Background()).ConsoleScreenshotCommand(consoleScreenshotCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.ServersConsole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServersConsole`: string
	fmt.Fprintf(os.Stdout, "Response from `ServersAPI.ServersConsole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServersConsoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **consoleScreenshotCommand** | [**ConsoleScreenshotCommand**](ConsoleScreenshotCommand.md) |  | 

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


## ServersCreate

> ApiResponse ServersCreate(ctx).ServerForCreateDto(serverForCreateDto).Execute()

Create a new server in the given project.

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
	serverForCreateDto := *openapiclient.NewServerForCreateDto() // ServerForCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServersAPI.ServersCreate(context.Background()).ServerForCreateDto(serverForCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.ServersCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServersCreate`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ServersAPI.ServersCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverForCreateDto** | [**ServerForCreateDto**](ServerForCreateDto.md) |  | 

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


## ServersDetails

> ServersListForDetails ServersDetails(ctx, projectId).SortBy(sortBy).SortDirection(sortDirection).WithAutoscalingGroup(withAutoscalingGroup).Execute()

Retrieve all servers by given project

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
	sortBy := "sortBy_example" // string |  (optional)
	sortDirection := "sortDirection_example" // string |  (optional)
	withAutoscalingGroup := "withAutoscalingGroup_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServersAPI.ServersDetails(context.Background(), projectId).SortBy(sortBy).SortDirection(sortDirection).WithAutoscalingGroup(withAutoscalingGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.ServersDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServersDetails`: ServersListForDetails
	fmt.Fprintf(os.Stdout, "Response from `ServersAPI.ServersDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **withAutoscalingGroup** | **string** |  | 

### Return type

[**ServersListForDetails**](ServersListForDetails.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServersGetCskClusterConfig

> string ServersGetCskClusterConfig(ctx).ServerId(serverId).Execute()

Get csk cluster config

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
	serverId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServersAPI.ServersGetCskClusterConfig(context.Background()).ServerId(serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.ServersGetCskClusterConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServersGetCskClusterConfig`: string
	fmt.Fprintf(os.Stdout, "Response from `ServersAPI.ServersGetCskClusterConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServersGetCskClusterConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverId** | **int32** |  | 

### Return type

**string**

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServersList

> ServersList ServersList(ctx).Limit(limit).Offset(offset).ProjectId(projectId).SortBy(sortBy).SortDirection(sortDirection).Search(search).StartRam(startRam).EndRam(endRam).StartDiskSize(startDiskSize).EndDiskSize(endDiskSize).StartCpu(startCpu).EndCpu(endCpu).OrganizationId(organizationId).Id(id).FilterBy(filterBy).AutoscalingGroup(autoscalingGroup).DomainId(domainId).Execute()

Retrieve all servers

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
	limit := int32(56) // int32 |  (optional)
	offset := int32(56) // int32 |  (optional)
	projectId := int32(56) // int32 |  (optional)
	sortBy := "sortBy_example" // string |  (optional)
	sortDirection := "sortDirection_example" // string |  (optional)
	search := "search_example" // string |  (optional)
	startRam := float64(1.2) // float64 |  (optional)
	endRam := float64(1.2) // float64 |  (optional)
	startDiskSize := int64(789) // int64 |  (optional)
	endDiskSize := int64(789) // int64 |  (optional)
	startCpu := int32(56) // int32 |  (optional)
	endCpu := int32(56) // int32 |  (optional)
	organizationId := int32(56) // int32 |  (optional)
	id := int32(56) // int32 |  (optional)
	filterBy := "filterBy_example" // string |  (optional)
	autoscalingGroup := "autoscalingGroup_example" // string |  (optional)
	domainId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServersAPI.ServersList(context.Background()).Limit(limit).Offset(offset).ProjectId(projectId).SortBy(sortBy).SortDirection(sortDirection).Search(search).StartRam(startRam).EndRam(endRam).StartDiskSize(startDiskSize).EndDiskSize(endDiskSize).StartCpu(startCpu).EndCpu(endCpu).OrganizationId(organizationId).Id(id).FilterBy(filterBy).AutoscalingGroup(autoscalingGroup).DomainId(domainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.ServersList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServersList`: ServersList
	fmt.Fprintf(os.Stdout, "Response from `ServersAPI.ServersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **projectId** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **startRam** | **float64** |  | 
 **endRam** | **float64** |  | 
 **startDiskSize** | **int64** |  | 
 **endDiskSize** | **int64** |  | 
 **startCpu** | **int32** |  | 
 **endCpu** | **int32** |  | 
 **organizationId** | **int32** |  | 
 **id** | **int32** |  | 
 **filterBy** | **string** |  | 
 **autoscalingGroup** | **string** |  | 
 **domainId** | **int32** |  | 

### Return type

[**ServersList**](ServersList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServersReboot

> ServersReboot(ctx).RebootServerCommand(rebootServerCommand).Execute()

Reboot server

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
	rebootServerCommand := *openapiclient.NewRebootServerCommand(int32(123), "Type_example") // RebootServerCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServersAPI.ServersReboot(context.Background()).RebootServerCommand(rebootServerCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.ServersReboot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServersRebootRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rebootServerCommand** | [**RebootServerCommand**](RebootServerCommand.md) |  | 

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


## ServersRegisterNode

> ServersRegisterNode(ctx).CskRegisterNodeCommand(cskRegisterNodeCommand).Execute()

Register csk node

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
	cskRegisterNodeCommand := *openapiclient.NewCskRegisterNodeCommand(int32(123)) // CskRegisterNodeCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServersAPI.ServersRegisterNode(context.Background()).CskRegisterNodeCommand(cskRegisterNodeCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.ServersRegisterNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServersRegisterNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cskRegisterNodeCommand** | [**CskRegisterNodeCommand**](CskRegisterNodeCommand.md) |  | 

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


## ServersReset

> ServersReset(ctx).ResetServerStatusCommand(resetServerStatusCommand).Execute()

Update server(s) status(es)

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
	resetServerStatusCommand := *openapiclient.NewResetServerStatusCommand() // ResetServerStatusCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServersAPI.ServersReset(context.Background()).ResetServerStatusCommand(resetServerStatusCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.ServersReset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServersResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resetServerStatusCommand** | [**ResetServerStatusCommand**](ResetServerStatusCommand.md) |  | 

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


## ServersStatus

> string ServersStatus(ctx, serverId).Execute()

Show server status

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
	serverId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServersAPI.ServersStatus(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.ServersStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServersStatus`: string
	fmt.Fprintf(os.Stdout, "Response from `ServersAPI.ServersStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServersUpdateByProjectId

> ServersUpdateByProjectId(ctx, projectId).UpdateServerHealthDto(updateServerHealthDto).Execute()

Update server by projectId

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
	updateServerHealthDto := []openapiclient.UpdateServerHealthDto{*openapiclient.NewUpdateServerHealthDto()} // []UpdateServerHealthDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServersAPI.ServersUpdateByProjectId(context.Background(), projectId).UpdateServerHealthDto(updateServerHealthDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.ServersUpdateByProjectId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersUpdateByProjectIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateServerHealthDto** | [**[]UpdateServerHealthDto**](UpdateServerHealthDto.md) |  | 

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

