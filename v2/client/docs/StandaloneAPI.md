# \StandaloneAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StandaloneCreate**](StandaloneAPI.md#StandaloneCreate) | **Post** /api/v1/standalone/create | Create a new vm in the given project.
[**StandaloneDetails**](StandaloneAPI.md#StandaloneDetails) | **Get** /api/v1/standalone/{projectId} | Retrieve a list of standalone vm with detailed info
[**StandaloneIpManagement**](StandaloneAPI.md#StandaloneIpManagement) | **Post** /api/v1/standalone/ip/management | Enable/Disable stand alone public ip
[**StandaloneList**](StandaloneAPI.md#StandaloneList) | **Get** /api/v1/standalone | Retrieve all vms
[**StandaloneReset**](StandaloneAPI.md#StandaloneReset) | **Post** /api/v1/standalone/reset | Reset vm status
[**StandaloneUpdateFlavor**](StandaloneAPI.md#StandaloneUpdateFlavor) | **Post** /api/v1/standalone/update/flavor | Update standalone vm flavor



## StandaloneCreate

> ApiResponse StandaloneCreate(ctx).CreateStandAloneVmCommand(createStandAloneVmCommand).Execute()

Create a new vm in the given project.

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
	createStandAloneVmCommand := *openapiclient.NewCreateStandAloneVmCommand() // CreateStandAloneVmCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StandaloneAPI.StandaloneCreate(context.Background()).CreateStandAloneVmCommand(createStandAloneVmCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StandaloneAPI.StandaloneCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StandaloneCreate`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `StandaloneAPI.StandaloneCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStandaloneCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createStandAloneVmCommand** | [**CreateStandAloneVmCommand**](CreateStandAloneVmCommand.md) |  | 

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


## StandaloneDetails

> StandAloneVmListForDetails StandaloneDetails(ctx, projectId).SortBy(sortBy).SortDirection(sortDirection).Id(id).Execute()

Retrieve a list of standalone vm with detailed info

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
	id := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StandaloneAPI.StandaloneDetails(context.Background(), projectId).SortBy(sortBy).SortDirection(sortDirection).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StandaloneAPI.StandaloneDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StandaloneDetails`: StandAloneVmListForDetails
	fmt.Fprintf(os.Stdout, "Response from `StandaloneAPI.StandaloneDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStandaloneDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **id** | **int32** |  | 

### Return type

[**StandAloneVmListForDetails**](StandAloneVmListForDetails.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StandaloneIpManagement

> StandaloneIpManagement(ctx).StandAloneVmIpManagementCommand(standAloneVmIpManagementCommand).Execute()

Enable/Disable stand alone public ip

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
	standAloneVmIpManagementCommand := *openapiclient.NewStandAloneVmIpManagementCommand(int32(123), "Mode_example") // StandAloneVmIpManagementCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StandaloneAPI.StandaloneIpManagement(context.Background()).StandAloneVmIpManagementCommand(standAloneVmIpManagementCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StandaloneAPI.StandaloneIpManagement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStandaloneIpManagementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **standAloneVmIpManagementCommand** | [**StandAloneVmIpManagementCommand**](StandAloneVmIpManagementCommand.md) |  | 

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


## StandaloneList

> StandaloneVmsList StandaloneList(ctx).Limit(limit).Offset(offset).ProjectId(projectId).SortBy(sortBy).SortDirection(sortDirection).Search(search).StartRam(startRam).EndRam(endRam).StartVolumeSize(startVolumeSize).EndVolumeSize(endVolumeSize).StartCpu(startCpu).EndCpu(endCpu).OrganizationId(organizationId).Id(id).SearchId(searchId).FilterBy(filterBy).DomainId(domainId).Execute()

Retrieve all vms

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
	startVolumeSize := int64(789) // int64 |  (optional)
	endVolumeSize := int64(789) // int64 |  (optional)
	startCpu := int32(56) // int32 |  (optional)
	endCpu := int32(56) // int32 |  (optional)
	organizationId := int32(56) // int32 |  (optional)
	id := int32(56) // int32 |  (optional)
	searchId := "searchId_example" // string |  (optional)
	filterBy := "filterBy_example" // string |  (optional)
	domainId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StandaloneAPI.StandaloneList(context.Background()).Limit(limit).Offset(offset).ProjectId(projectId).SortBy(sortBy).SortDirection(sortDirection).Search(search).StartRam(startRam).EndRam(endRam).StartVolumeSize(startVolumeSize).EndVolumeSize(endVolumeSize).StartCpu(startCpu).EndCpu(endCpu).OrganizationId(organizationId).Id(id).SearchId(searchId).FilterBy(filterBy).DomainId(domainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StandaloneAPI.StandaloneList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StandaloneList`: StandaloneVmsList
	fmt.Fprintf(os.Stdout, "Response from `StandaloneAPI.StandaloneList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStandaloneListRequest struct via the builder pattern


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
 **startVolumeSize** | **int64** |  | 
 **endVolumeSize** | **int64** |  | 
 **startCpu** | **int32** |  | 
 **endCpu** | **int32** |  | 
 **organizationId** | **int32** |  | 
 **id** | **int32** |  | 
 **searchId** | **string** |  | 
 **filterBy** | **string** |  | 
 **domainId** | **int32** |  | 

### Return type

[**StandaloneVmsList**](StandaloneVmsList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StandaloneReset

> StandaloneReset(ctx).ResetStandAloneVmStatusCommand(resetStandAloneVmStatusCommand).Execute()

Reset vm status

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
	resetStandAloneVmStatusCommand := *openapiclient.NewResetStandAloneVmStatusCommand() // ResetStandAloneVmStatusCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StandaloneAPI.StandaloneReset(context.Background()).ResetStandAloneVmStatusCommand(resetStandAloneVmStatusCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StandaloneAPI.StandaloneReset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStandaloneResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resetStandAloneVmStatusCommand** | [**ResetStandAloneVmStatusCommand**](ResetStandAloneVmStatusCommand.md) |  | 

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


## StandaloneUpdateFlavor

> StandaloneUpdateFlavor(ctx).UpdateStandAloneVmFlavorCommand(updateStandAloneVmFlavorCommand).Execute()

Update standalone vm flavor

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
	updateStandAloneVmFlavorCommand := *openapiclient.NewUpdateStandAloneVmFlavorCommand(int32(123), "Flavor_example") // UpdateStandAloneVmFlavorCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StandaloneAPI.StandaloneUpdateFlavor(context.Background()).UpdateStandAloneVmFlavorCommand(updateStandAloneVmFlavorCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StandaloneAPI.StandaloneUpdateFlavor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStandaloneUpdateFlavorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateStandAloneVmFlavorCommand** | [**UpdateStandAloneVmFlavorCommand**](UpdateStandAloneVmFlavorCommand.md) |  | 

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

