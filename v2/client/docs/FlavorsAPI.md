# \FlavorsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FlavorsAwsInstanceTypes**](FlavorsAPI.md#FlavorsAwsInstanceTypes) | **Get** /api/v1/flavors/aws/{cloudId} | Retrieve awz instance types
[**FlavorsAzureVmSizes**](FlavorsAPI.md#FlavorsAzureVmSizes) | **Get** /api/v1/flavors/azure/{cloudId} | Retrieve azure vm sizes
[**FlavorsBindToProject**](FlavorsAPI.md#FlavorsBindToProject) | **Post** /api/v1/flavors/bind | Bind flavors to project
[**FlavorsDropdownFlavors**](FlavorsAPI.md#FlavorsDropdownFlavors) | **Get** /api/v1/flavors/credentials/dropdown/list | Retrieve cloud credentials dropdown list
[**FlavorsGoogleMachineTypes**](FlavorsAPI.md#FlavorsGoogleMachineTypes) | **Get** /api/v1/flavors/google/{cloudId} | Retrieve google machine types
[**FlavorsOpenstackFlavors**](FlavorsAPI.md#FlavorsOpenstackFlavors) | **Get** /api/v1/flavors/openstack/{cloudId} | Retrieve openstack flavors
[**FlavorsSelectedFlavorsForProject**](FlavorsAPI.md#FlavorsSelectedFlavorsForProject) | **Get** /api/v1/flavors/projects/list | Retrieve selected flavors for project
[**FlavorsUnbindFromProject**](FlavorsAPI.md#FlavorsUnbindFromProject) | **Post** /api/v1/flavors/unbind | Unbind flavors from project



## FlavorsAwsInstanceTypes

> AwsFlavorList FlavorsAwsInstanceTypes(ctx, cloudId).Limit(limit).Offset(offset).StartRam(startRam).EndRam(endRam).StartCpu(startCpu).EndCpu(endCpu).Search(search).SortBy(sortBy).SortDirection(sortDirection).ProjectId(projectId).HasGpuSupport(hasGpuSupport).StartGpuCount(startGpuCount).EndGpuCount(endGpuCount).GpuModel(gpuModel).Execute()

Retrieve awz instance types

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
	cloudId := int32(56) // int32 | 
	limit := int32(56) // int32 |  (optional)
	offset := int32(56) // int32 |  (optional)
	startRam := float64(1.2) // float64 |  (optional)
	endRam := float64(1.2) // float64 |  (optional)
	startCpu := int32(56) // int32 |  (optional)
	endCpu := int32(56) // int32 |  (optional)
	search := "search_example" // string |  (optional)
	sortBy := "sortBy_example" // string |  (optional)
	sortDirection := "sortDirection_example" // string |  (optional)
	projectId := int32(56) // int32 |  (optional)
	hasGpuSupport := true // bool |  (optional)
	startGpuCount := int32(56) // int32 |  (optional)
	endGpuCount := int32(56) // int32 |  (optional)
	gpuModel := "gpuModel_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlavorsAPI.FlavorsAwsInstanceTypes(context.Background(), cloudId).Limit(limit).Offset(offset).StartRam(startRam).EndRam(endRam).StartCpu(startCpu).EndCpu(endCpu).Search(search).SortBy(sortBy).SortDirection(sortDirection).ProjectId(projectId).HasGpuSupport(hasGpuSupport).StartGpuCount(startGpuCount).EndGpuCount(endGpuCount).GpuModel(gpuModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlavorsAPI.FlavorsAwsInstanceTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlavorsAwsInstanceTypes`: AwsFlavorList
	fmt.Fprintf(os.Stdout, "Response from `FlavorsAPI.FlavorsAwsInstanceTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlavorsAwsInstanceTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **startRam** | **float64** |  | 
 **endRam** | **float64** |  | 
 **startCpu** | **int32** |  | 
 **endCpu** | **int32** |  | 
 **search** | **string** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **projectId** | **int32** |  | 
 **hasGpuSupport** | **bool** |  | 
 **startGpuCount** | **int32** |  | 
 **endGpuCount** | **int32** |  | 
 **gpuModel** | **string** |  | 

### Return type

[**AwsFlavorList**](AwsFlavorList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlavorsAzureVmSizes

> AzureFlavorList FlavorsAzureVmSizes(ctx, cloudId).Limit(limit).Offset(offset).StartRam(startRam).EndRam(endRam).StartCpu(startCpu).EndCpu(endCpu).Search(search).SortBy(sortBy).SortDirection(sortDirection).ProjectId(projectId).HasGpuSupport(hasGpuSupport).StartGpuCount(startGpuCount).EndGpuCount(endGpuCount).GpuModel(gpuModel).Execute()

Retrieve azure vm sizes

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
	cloudId := int32(56) // int32 | 
	limit := int32(56) // int32 |  (optional)
	offset := int32(56) // int32 |  (optional)
	startRam := float64(1.2) // float64 |  (optional)
	endRam := float64(1.2) // float64 |  (optional)
	startCpu := int32(56) // int32 |  (optional)
	endCpu := int32(56) // int32 |  (optional)
	search := "search_example" // string |  (optional)
	sortBy := "sortBy_example" // string |  (optional)
	sortDirection := "sortDirection_example" // string |  (optional)
	projectId := int32(56) // int32 |  (optional)
	hasGpuSupport := true // bool |  (optional)
	startGpuCount := int32(56) // int32 |  (optional)
	endGpuCount := int32(56) // int32 |  (optional)
	gpuModel := "gpuModel_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlavorsAPI.FlavorsAzureVmSizes(context.Background(), cloudId).Limit(limit).Offset(offset).StartRam(startRam).EndRam(endRam).StartCpu(startCpu).EndCpu(endCpu).Search(search).SortBy(sortBy).SortDirection(sortDirection).ProjectId(projectId).HasGpuSupport(hasGpuSupport).StartGpuCount(startGpuCount).EndGpuCount(endGpuCount).GpuModel(gpuModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlavorsAPI.FlavorsAzureVmSizes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlavorsAzureVmSizes`: AzureFlavorList
	fmt.Fprintf(os.Stdout, "Response from `FlavorsAPI.FlavorsAzureVmSizes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlavorsAzureVmSizesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **startRam** | **float64** |  | 
 **endRam** | **float64** |  | 
 **startCpu** | **int32** |  | 
 **endCpu** | **int32** |  | 
 **search** | **string** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **projectId** | **int32** |  | 
 **hasGpuSupport** | **bool** |  | 
 **startGpuCount** | **int32** |  | 
 **endGpuCount** | **int32** |  | 
 **gpuModel** | **string** |  | 

### Return type

[**AzureFlavorList**](AzureFlavorList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlavorsBindToProject

> FlavorsBindToProject(ctx).BindFlavorToProjectCommand(bindFlavorToProjectCommand).Execute()

Bind flavors to project

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
	bindFlavorToProjectCommand := *openapiclient.NewBindFlavorToProjectCommand(int32(123), []string{"Flavors_example"}) // BindFlavorToProjectCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FlavorsAPI.FlavorsBindToProject(context.Background()).BindFlavorToProjectCommand(bindFlavorToProjectCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlavorsAPI.FlavorsBindToProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlavorsBindToProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bindFlavorToProjectCommand** | [**BindFlavorToProjectCommand**](BindFlavorToProjectCommand.md) |  | 

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


## FlavorsDropdownFlavors

> []CloudCredentialsDropdownRecordDto FlavorsDropdownFlavors(ctx).OrganizationId(organizationId).FilterBy(filterBy).Search(search).IsInfra(isInfra).Execute()

Retrieve cloud credentials dropdown list

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
	filterBy := "filterBy_example" // string |  (optional)
	search := "search_example" // string |  (optional)
	isInfra := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlavorsAPI.FlavorsDropdownFlavors(context.Background()).OrganizationId(organizationId).FilterBy(filterBy).Search(search).IsInfra(isInfra).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlavorsAPI.FlavorsDropdownFlavors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlavorsDropdownFlavors`: []CloudCredentialsDropdownRecordDto
	fmt.Fprintf(os.Stdout, "Response from `FlavorsAPI.FlavorsDropdownFlavors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlavorsDropdownFlavorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 
 **filterBy** | **string** |  | 
 **search** | **string** |  | 
 **isInfra** | **bool** |  | 

### Return type

[**[]CloudCredentialsDropdownRecordDto**](CloudCredentialsDropdownRecordDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlavorsGoogleMachineTypes

> GoogleFlavorList FlavorsGoogleMachineTypes(ctx, cloudId).Limit(limit).Offset(offset).StartRam(startRam).EndRam(endRam).StartCpu(startCpu).EndCpu(endCpu).Search(search).SortBy(sortBy).SortDirection(sortDirection).ProjectId(projectId).HasGpuSupport(hasGpuSupport).StartGpuCount(startGpuCount).EndGpuCount(endGpuCount).GpuModel(gpuModel).Execute()

Retrieve google machine types

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
	cloudId := int32(56) // int32 | 
	limit := int32(56) // int32 |  (optional)
	offset := int32(56) // int32 |  (optional)
	startRam := float64(1.2) // float64 |  (optional)
	endRam := float64(1.2) // float64 |  (optional)
	startCpu := int32(56) // int32 |  (optional)
	endCpu := int32(56) // int32 |  (optional)
	search := "search_example" // string |  (optional)
	sortBy := "sortBy_example" // string |  (optional)
	sortDirection := "sortDirection_example" // string |  (optional)
	projectId := int32(56) // int32 |  (optional)
	hasGpuSupport := true // bool |  (optional)
	startGpuCount := int32(56) // int32 |  (optional)
	endGpuCount := int32(56) // int32 |  (optional)
	gpuModel := "gpuModel_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlavorsAPI.FlavorsGoogleMachineTypes(context.Background(), cloudId).Limit(limit).Offset(offset).StartRam(startRam).EndRam(endRam).StartCpu(startCpu).EndCpu(endCpu).Search(search).SortBy(sortBy).SortDirection(sortDirection).ProjectId(projectId).HasGpuSupport(hasGpuSupport).StartGpuCount(startGpuCount).EndGpuCount(endGpuCount).GpuModel(gpuModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlavorsAPI.FlavorsGoogleMachineTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlavorsGoogleMachineTypes`: GoogleFlavorList
	fmt.Fprintf(os.Stdout, "Response from `FlavorsAPI.FlavorsGoogleMachineTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlavorsGoogleMachineTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **startRam** | **float64** |  | 
 **endRam** | **float64** |  | 
 **startCpu** | **int32** |  | 
 **endCpu** | **int32** |  | 
 **search** | **string** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **projectId** | **int32** |  | 
 **hasGpuSupport** | **bool** |  | 
 **startGpuCount** | **int32** |  | 
 **endGpuCount** | **int32** |  | 
 **gpuModel** | **string** |  | 

### Return type

[**GoogleFlavorList**](GoogleFlavorList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlavorsOpenstackFlavors

> OpenstackFlavorList FlavorsOpenstackFlavors(ctx, cloudId).Limit(limit).Offset(offset).StartRam(startRam).EndRam(endRam).StartCpu(startCpu).EndCpu(endCpu).Search(search).SortBy(sortBy).SortDirection(sortDirection).ProjectId(projectId).HasGpuSupport(hasGpuSupport).Execute()

Retrieve openstack flavors

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
	cloudId := int32(56) // int32 | 
	limit := int32(56) // int32 |  (optional)
	offset := int32(56) // int32 |  (optional)
	startRam := float64(1.2) // float64 |  (optional)
	endRam := float64(1.2) // float64 |  (optional)
	startCpu := int32(56) // int32 |  (optional)
	endCpu := int32(56) // int32 |  (optional)
	search := "search_example" // string |  (optional)
	sortBy := "sortBy_example" // string |  (optional)
	sortDirection := "sortDirection_example" // string |  (optional)
	projectId := int32(56) // int32 |  (optional)
	hasGpuSupport := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlavorsAPI.FlavorsOpenstackFlavors(context.Background(), cloudId).Limit(limit).Offset(offset).StartRam(startRam).EndRam(endRam).StartCpu(startCpu).EndCpu(endCpu).Search(search).SortBy(sortBy).SortDirection(sortDirection).ProjectId(projectId).HasGpuSupport(hasGpuSupport).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlavorsAPI.FlavorsOpenstackFlavors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlavorsOpenstackFlavors`: OpenstackFlavorList
	fmt.Fprintf(os.Stdout, "Response from `FlavorsAPI.FlavorsOpenstackFlavors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlavorsOpenstackFlavorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **startRam** | **float64** |  | 
 **endRam** | **float64** |  | 
 **startCpu** | **int32** |  | 
 **endCpu** | **int32** |  | 
 **search** | **string** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **projectId** | **int32** |  | 
 **hasGpuSupport** | **bool** |  | 

### Return type

[**OpenstackFlavorList**](OpenstackFlavorList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlavorsSelectedFlavorsForProject

> BoundFlavorsForProjectsList FlavorsSelectedFlavorsForProject(ctx).Limit(limit).Offset(offset).ProjectId(projectId).SortBy(sortBy).SortDirection(sortDirection).Search(search).FilterBy(filterBy).OrganizationId(organizationId).FlavorName(flavorName).WithPrice(withPrice).HasGpuSupport(hasGpuSupport).Execute()

Retrieve selected flavors for project

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
	filterBy := "filterBy_example" // string |  (optional)
	organizationId := int32(56) // int32 |  (optional)
	flavorName := "flavorName_example" // string |  (optional)
	withPrice := true // bool |  (optional)
	hasGpuSupport := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlavorsAPI.FlavorsSelectedFlavorsForProject(context.Background()).Limit(limit).Offset(offset).ProjectId(projectId).SortBy(sortBy).SortDirection(sortDirection).Search(search).FilterBy(filterBy).OrganizationId(organizationId).FlavorName(flavorName).WithPrice(withPrice).HasGpuSupport(hasGpuSupport).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlavorsAPI.FlavorsSelectedFlavorsForProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlavorsSelectedFlavorsForProject`: BoundFlavorsForProjectsList
	fmt.Fprintf(os.Stdout, "Response from `FlavorsAPI.FlavorsSelectedFlavorsForProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlavorsSelectedFlavorsForProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **projectId** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **filterBy** | **string** |  | 
 **organizationId** | **int32** |  | 
 **flavorName** | **string** |  | 
 **withPrice** | **bool** |  | 
 **hasGpuSupport** | **bool** |  | 

### Return type

[**BoundFlavorsForProjectsList**](BoundFlavorsForProjectsList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlavorsUnbindFromProject

> FlavorsUnbindFromProject(ctx).UnbindFlavorFromProjectCommand(unbindFlavorFromProjectCommand).Execute()

Unbind flavors from project

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
	unbindFlavorFromProjectCommand := *openapiclient.NewUnbindFlavorFromProjectCommand([]int32{int32(123)}) // UnbindFlavorFromProjectCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FlavorsAPI.FlavorsUnbindFromProject(context.Background()).UnbindFlavorFromProjectCommand(unbindFlavorFromProjectCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlavorsAPI.FlavorsUnbindFromProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlavorsUnbindFromProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unbindFlavorFromProjectCommand** | [**UnbindFlavorFromProjectCommand**](UnbindFlavorFromProjectCommand.md) |  | 

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

