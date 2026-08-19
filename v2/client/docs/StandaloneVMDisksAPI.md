# \StandaloneVMDisksAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StandalonevmdisksCreate**](StandaloneVMDisksAPI.md#StandalonevmdisksCreate) | **Post** /api/v1/standalonevmdisks/create | Add disk for standalone vm
[**StandalonevmdisksUpdateSize**](StandaloneVMDisksAPI.md#StandalonevmdisksUpdateSize) | **Post** /api/v1/standalonevmdisks/update-size | Update disk size



## StandalonevmdisksCreate

> ApiResponse StandalonevmdisksCreate(ctx).CreateStandAloneDiskCommand(createStandAloneDiskCommand).Execute()

Add disk for standalone vm

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
	createStandAloneDiskCommand := *openapiclient.NewCreateStandAloneDiskCommand() // CreateStandAloneDiskCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StandaloneVMDisksAPI.StandalonevmdisksCreate(context.Background()).CreateStandAloneDiskCommand(createStandAloneDiskCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StandaloneVMDisksAPI.StandalonevmdisksCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StandalonevmdisksCreate`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `StandaloneVMDisksAPI.StandalonevmdisksCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStandalonevmdisksCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createStandAloneDiskCommand** | [**CreateStandAloneDiskCommand**](CreateStandAloneDiskCommand.md) |  | 

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


## StandalonevmdisksUpdateSize

> StandalonevmdisksUpdateSize(ctx).UpdateStandaloneVmDiskSizeCommand(updateStandaloneVmDiskSizeCommand).Execute()

Update disk size

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
	updateStandaloneVmDiskSizeCommand := *openapiclient.NewUpdateStandaloneVmDiskSizeCommand(int32(123), int32(123)) // UpdateStandaloneVmDiskSizeCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StandaloneVMDisksAPI.StandalonevmdisksUpdateSize(context.Background()).UpdateStandaloneVmDiskSizeCommand(updateStandaloneVmDiskSizeCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StandaloneVMDisksAPI.StandalonevmdisksUpdateSize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStandalonevmdisksUpdateSizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateStandaloneVmDiskSizeCommand** | [**UpdateStandaloneVmDiskSizeCommand**](UpdateStandaloneVmDiskSizeCommand.md) |  | 

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

