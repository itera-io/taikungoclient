# \StandaloneActionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StandaloneactionsConsole**](StandaloneActionsAPI.md#StandaloneactionsConsole) | **Post** /api/v1/standaloneactions/console | Console screenshot or terminal for vm
[**StandaloneactionsDownloadRdp**](StandaloneActionsAPI.md#StandaloneactionsDownloadRdp) | **Get** /api/v1/standaloneactions/download/rdp/{id} | Download RDP file
[**StandaloneactionsReboot**](StandaloneActionsAPI.md#StandaloneactionsReboot) | **Post** /api/v1/standaloneactions/reboot | Reboot standalone vm
[**StandaloneactionsShelve**](StandaloneActionsAPI.md#StandaloneactionsShelve) | **Post** /api/v1/standaloneactions/shelve | Shelve standalone vm
[**StandaloneactionsStart**](StandaloneActionsAPI.md#StandaloneactionsStart) | **Post** /api/v1/standaloneactions/start | Start standalone vm
[**StandaloneactionsStatus**](StandaloneActionsAPI.md#StandaloneactionsStatus) | **Get** /api/v1/standaloneactions/status/{id} | Show standalone vm status
[**StandaloneactionsStop**](StandaloneActionsAPI.md#StandaloneactionsStop) | **Post** /api/v1/standaloneactions/stop | Stop standalone vm
[**StandaloneactionsUnshelve**](StandaloneActionsAPI.md#StandaloneactionsUnshelve) | **Post** /api/v1/standaloneactions/unshelve | Unshelve standalone vm
[**StandaloneactionsWindowsInstancePassword**](StandaloneActionsAPI.md#StandaloneactionsWindowsInstancePassword) | **Post** /api/v1/standaloneactions/password | Retrieve aws windows admin instance password



## StandaloneactionsConsole

> string StandaloneactionsConsole(ctx).VmConsoleScreenshotCommand(vmConsoleScreenshotCommand).Execute()

Console screenshot or terminal for vm

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
	vmConsoleScreenshotCommand := *openapiclient.NewVmConsoleScreenshotCommand(int32(123)) // VmConsoleScreenshotCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StandaloneActionsAPI.StandaloneactionsConsole(context.Background()).VmConsoleScreenshotCommand(vmConsoleScreenshotCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StandaloneActionsAPI.StandaloneactionsConsole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StandaloneactionsConsole`: string
	fmt.Fprintf(os.Stdout, "Response from `StandaloneActionsAPI.StandaloneactionsConsole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStandaloneactionsConsoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vmConsoleScreenshotCommand** | [**VmConsoleScreenshotCommand**](VmConsoleScreenshotCommand.md) |  | 

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


## StandaloneactionsDownloadRdp

> CsvExporter StandaloneactionsDownloadRdp(ctx, id).Execute()

Download RDP file

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
	resp, r, err := apiClient.StandaloneActionsAPI.StandaloneactionsDownloadRdp(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StandaloneActionsAPI.StandaloneactionsDownloadRdp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StandaloneactionsDownloadRdp`: CsvExporter
	fmt.Fprintf(os.Stdout, "Response from `StandaloneActionsAPI.StandaloneactionsDownloadRdp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStandaloneactionsDownloadRdpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CsvExporter**](CsvExporter.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StandaloneactionsReboot

> StandaloneactionsReboot(ctx).RebootStandAloneVmCommand(rebootStandAloneVmCommand).Execute()

Reboot standalone vm

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
	rebootStandAloneVmCommand := *openapiclient.NewRebootStandAloneVmCommand(int32(123), "Type_example") // RebootStandAloneVmCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StandaloneActionsAPI.StandaloneactionsReboot(context.Background()).RebootStandAloneVmCommand(rebootStandAloneVmCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StandaloneActionsAPI.StandaloneactionsReboot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStandaloneactionsRebootRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rebootStandAloneVmCommand** | [**RebootStandAloneVmCommand**](RebootStandAloneVmCommand.md) |  | 

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


## StandaloneactionsShelve

> StandaloneactionsShelve(ctx).ShelveStandAloneVmCommand(shelveStandAloneVmCommand).Execute()

Shelve standalone vm

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
	shelveStandAloneVmCommand := *openapiclient.NewShelveStandAloneVmCommand(int32(123)) // ShelveStandAloneVmCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StandaloneActionsAPI.StandaloneactionsShelve(context.Background()).ShelveStandAloneVmCommand(shelveStandAloneVmCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StandaloneActionsAPI.StandaloneactionsShelve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStandaloneactionsShelveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shelveStandAloneVmCommand** | [**ShelveStandAloneVmCommand**](ShelveStandAloneVmCommand.md) |  | 

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


## StandaloneactionsStart

> StandaloneactionsStart(ctx).StartStandaloneVmCommand(startStandaloneVmCommand).Execute()

Start standalone vm

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
	startStandaloneVmCommand := *openapiclient.NewStartStandaloneVmCommand(int32(123)) // StartStandaloneVmCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StandaloneActionsAPI.StandaloneactionsStart(context.Background()).StartStandaloneVmCommand(startStandaloneVmCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StandaloneActionsAPI.StandaloneactionsStart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStandaloneactionsStartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startStandaloneVmCommand** | [**StartStandaloneVmCommand**](StartStandaloneVmCommand.md) |  | 

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


## StandaloneactionsStatus

> string StandaloneactionsStatus(ctx, id).Execute()

Show standalone vm status

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
	resp, r, err := apiClient.StandaloneActionsAPI.StandaloneactionsStatus(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StandaloneActionsAPI.StandaloneactionsStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StandaloneactionsStatus`: string
	fmt.Fprintf(os.Stdout, "Response from `StandaloneActionsAPI.StandaloneactionsStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStandaloneactionsStatusRequest struct via the builder pattern


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


## StandaloneactionsStop

> StandaloneactionsStop(ctx).StopStandaloneVmCommand(stopStandaloneVmCommand).Execute()

Stop standalone vm

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
	stopStandaloneVmCommand := *openapiclient.NewStopStandaloneVmCommand(int32(123)) // StopStandaloneVmCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StandaloneActionsAPI.StandaloneactionsStop(context.Background()).StopStandaloneVmCommand(stopStandaloneVmCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StandaloneActionsAPI.StandaloneactionsStop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStandaloneactionsStopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stopStandaloneVmCommand** | [**StopStandaloneVmCommand**](StopStandaloneVmCommand.md) |  | 

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


## StandaloneactionsUnshelve

> StandaloneactionsUnshelve(ctx).UnshelveStandaloneVmCommand(unshelveStandaloneVmCommand).Execute()

Unshelve standalone vm

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
	unshelveStandaloneVmCommand := *openapiclient.NewUnshelveStandaloneVmCommand(int32(123)) // UnshelveStandaloneVmCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StandaloneActionsAPI.StandaloneactionsUnshelve(context.Background()).UnshelveStandaloneVmCommand(unshelveStandaloneVmCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StandaloneActionsAPI.StandaloneactionsUnshelve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStandaloneactionsUnshelveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unshelveStandaloneVmCommand** | [**UnshelveStandaloneVmCommand**](UnshelveStandaloneVmCommand.md) |  | 

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


## StandaloneactionsWindowsInstancePassword

> string StandaloneactionsWindowsInstancePassword(ctx).Id(id).Key(key).Config(config).Execute()

Retrieve aws windows admin instance password

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
	id := int32(56) // int32 |  (optional)
	key := "key_example" // string |  (optional)
	config := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StandaloneActionsAPI.StandaloneactionsWindowsInstancePassword(context.Background()).Id(id).Key(key).Config(config).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StandaloneActionsAPI.StandaloneactionsWindowsInstancePassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StandaloneactionsWindowsInstancePassword`: string
	fmt.Fprintf(os.Stdout, "Response from `StandaloneActionsAPI.StandaloneactionsWindowsInstancePassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStandaloneactionsWindowsInstancePasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32** |  | 
 **key** | **string** |  | 
 **config** | ***os.File** |  | 

### Return type

**string**

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

