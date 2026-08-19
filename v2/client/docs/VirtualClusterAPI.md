# \VirtualClusterAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VirtualClusterCreate**](VirtualClusterAPI.md#VirtualClusterCreate) | **Post** /api/v1/virtual-cluster/create | Create virtual cluster
[**VirtualClusterDelete**](VirtualClusterAPI.md#VirtualClusterDelete) | **Post** /api/v1/virtual-cluster/delete | Delete virtual cluster
[**VirtualClusterDisableAi**](VirtualClusterAPI.md#VirtualClusterDisableAi) | **Post** /api/v1/virtual-cluster/disable-ai | Disable ai for virtual cluster
[**VirtualClusterDisableBackup**](VirtualClusterAPI.md#VirtualClusterDisableBackup) | **Post** /api/v1/virtual-cluster/disable-backup | Disable backup for virtual cluster
[**VirtualClusterDisableMonitoring**](VirtualClusterAPI.md#VirtualClusterDisableMonitoring) | **Post** /api/v1/virtual-cluster/disable-monitoring | Disable monitoring for virtual cluster
[**VirtualClusterDisableOpa**](VirtualClusterAPI.md#VirtualClusterDisableOpa) | **Post** /api/v1/virtual-cluster/disable-opa | Disable opa for virtual cluster
[**VirtualClusterEditQuotas**](VirtualClusterAPI.md#VirtualClusterEditQuotas) | **Post** /api/v1/virtual-cluster/edit/quotas | Edit virtual cluster resource quotas and limit ranges
[**VirtualClusterEnableAi**](VirtualClusterAPI.md#VirtualClusterEnableAi) | **Post** /api/v1/virtual-cluster/enable-ai | Enable ai for virtual cluster
[**VirtualClusterEnableBackup**](VirtualClusterAPI.md#VirtualClusterEnableBackup) | **Post** /api/v1/virtual-cluster/enable-backup | Enable backup for virtual cluster
[**VirtualClusterEnableMonitoring**](VirtualClusterAPI.md#VirtualClusterEnableMonitoring) | **Post** /api/v1/virtual-cluster/enable-monitoring | Enable monitoring for virtual cluster
[**VirtualClusterEnableOpa**](VirtualClusterAPI.md#VirtualClusterEnableOpa) | **Post** /api/v1/virtual-cluster/enable-opa | Enable opa for virtual cluster
[**VirtualClusterList**](VirtualClusterAPI.md#VirtualClusterList) | **Get** /api/v1/virtual-cluster/{parentProjectId} | Retrieve all vCluster by given project
[**VirtualClusterQuota**](VirtualClusterAPI.md#VirtualClusterQuota) | **Get** /api/v1/virtual-cluster/quota/{projectId} | Retrieve vcluster quota details
[**VirtualClusterQuotaPresets**](VirtualClusterAPI.md#VirtualClusterQuotaPresets) | **Get** /api/v1/virtual-cluster/quota-presets | Retrieve vCluster quota presets
[**VirtualClusterVisibility**](VirtualClusterAPI.md#VirtualClusterVisibility) | **Get** /api/v1/virtual-cluster/visibility/{projectId} | Create button condition visibility



## VirtualClusterCreate

> VirtualClusterCreate(ctx).CreateVirtualClusterCommand(createVirtualClusterCommand).Execute()

Create virtual cluster

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/itera-io/taikungoclient/v2/client"
)

func main() {
	createVirtualClusterCommand := *openapiclient.NewCreateVirtualClusterCommand(NullableInt32(123), int32(123), "Name_example", time.Now(), false, NullableInt32(123), "ExposeHostname_example", *openapiclient.NewCreateVirtualClusterWorkloadResources(int32(123), int32(123), int64(123), int32(123), int32(123)), *openapiclient.NewCreateVirtualClusterResourceLimits(openapiclient.EProjectResourceUnit("Cores"), float64(123), float64(123), float64(123), float64(123)), *openapiclient.NewCreateVirtualClusterResourceLimits(openapiclient.EProjectResourceUnit("Cores"), float64(123), float64(123), float64(123), float64(123)), *openapiclient.NewCreateVirtualClusterResourceLimits(openapiclient.EProjectResourceUnit("Cores"), float64(123), float64(123), float64(123), float64(123))) // CreateVirtualClusterCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualClusterAPI.VirtualClusterCreate(context.Background()).CreateVirtualClusterCommand(createVirtualClusterCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualClusterAPI.VirtualClusterCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVirtualClusterCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createVirtualClusterCommand** | [**CreateVirtualClusterCommand**](CreateVirtualClusterCommand.md) |  | 

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


## VirtualClusterDelete

> VirtualClusterDelete(ctx).DeleteVirtualClusterCommand(deleteVirtualClusterCommand).Execute()

Delete virtual cluster

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
	deleteVirtualClusterCommand := *openapiclient.NewDeleteVirtualClusterCommand(int32(123)) // DeleteVirtualClusterCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualClusterAPI.VirtualClusterDelete(context.Background()).DeleteVirtualClusterCommand(deleteVirtualClusterCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualClusterAPI.VirtualClusterDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVirtualClusterDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteVirtualClusterCommand** | [**DeleteVirtualClusterCommand**](DeleteVirtualClusterCommand.md) |  | 

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


## VirtualClusterDisableAi

> VirtualClusterDisableAi(ctx).VirtualClusterDisableAiCommand(virtualClusterDisableAiCommand).Execute()

Disable ai for virtual cluster

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
	virtualClusterDisableAiCommand := *openapiclient.NewVirtualClusterDisableAiCommand(int32(123)) // VirtualClusterDisableAiCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualClusterAPI.VirtualClusterDisableAi(context.Background()).VirtualClusterDisableAiCommand(virtualClusterDisableAiCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualClusterAPI.VirtualClusterDisableAi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVirtualClusterDisableAiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **virtualClusterDisableAiCommand** | [**VirtualClusterDisableAiCommand**](VirtualClusterDisableAiCommand.md) |  | 

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


## VirtualClusterDisableBackup

> VirtualClusterDisableBackup(ctx).VirtualClusterDisableBackupCommand(virtualClusterDisableBackupCommand).Execute()

Disable backup for virtual cluster

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
	virtualClusterDisableBackupCommand := *openapiclient.NewVirtualClusterDisableBackupCommand(int32(123)) // VirtualClusterDisableBackupCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualClusterAPI.VirtualClusterDisableBackup(context.Background()).VirtualClusterDisableBackupCommand(virtualClusterDisableBackupCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualClusterAPI.VirtualClusterDisableBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVirtualClusterDisableBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **virtualClusterDisableBackupCommand** | [**VirtualClusterDisableBackupCommand**](VirtualClusterDisableBackupCommand.md) |  | 

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


## VirtualClusterDisableMonitoring

> VirtualClusterDisableMonitoring(ctx).VirtualClusterDisableMonitoringCommand(virtualClusterDisableMonitoringCommand).Execute()

Disable monitoring for virtual cluster

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
	virtualClusterDisableMonitoringCommand := *openapiclient.NewVirtualClusterDisableMonitoringCommand(int32(123)) // VirtualClusterDisableMonitoringCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualClusterAPI.VirtualClusterDisableMonitoring(context.Background()).VirtualClusterDisableMonitoringCommand(virtualClusterDisableMonitoringCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualClusterAPI.VirtualClusterDisableMonitoring``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVirtualClusterDisableMonitoringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **virtualClusterDisableMonitoringCommand** | [**VirtualClusterDisableMonitoringCommand**](VirtualClusterDisableMonitoringCommand.md) |  | 

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


## VirtualClusterDisableOpa

> VirtualClusterDisableOpa(ctx).VirtualClusterDisableOpaCommand(virtualClusterDisableOpaCommand).Execute()

Disable opa for virtual cluster

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
	virtualClusterDisableOpaCommand := *openapiclient.NewVirtualClusterDisableOpaCommand(int32(123)) // VirtualClusterDisableOpaCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualClusterAPI.VirtualClusterDisableOpa(context.Background()).VirtualClusterDisableOpaCommand(virtualClusterDisableOpaCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualClusterAPI.VirtualClusterDisableOpa``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVirtualClusterDisableOpaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **virtualClusterDisableOpaCommand** | [**VirtualClusterDisableOpaCommand**](VirtualClusterDisableOpaCommand.md) |  | 

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


## VirtualClusterEditQuotas

> VirtualClusterEditQuotas(ctx).EditVirtualClusterQuotasCommand(editVirtualClusterQuotasCommand).Execute()

Edit virtual cluster resource quotas and limit ranges

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
	editVirtualClusterQuotasCommand := *openapiclient.NewEditVirtualClusterQuotasCommand(int32(123), *openapiclient.NewEditVirtualClusterWorkloadResources(int32(123), int32(123), int64(123), int32(123), int32(123)), *openapiclient.NewEditVirtualClusterResourceLimits(openapiclient.EProjectResourceUnit("Cores"), float64(123), float64(123), float64(123), float64(123)), *openapiclient.NewEditVirtualClusterResourceLimits(openapiclient.EProjectResourceUnit("Cores"), float64(123), float64(123), float64(123), float64(123)), *openapiclient.NewEditVirtualClusterResourceLimits(openapiclient.EProjectResourceUnit("Cores"), float64(123), float64(123), float64(123), float64(123))) // EditVirtualClusterQuotasCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualClusterAPI.VirtualClusterEditQuotas(context.Background()).EditVirtualClusterQuotasCommand(editVirtualClusterQuotasCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualClusterAPI.VirtualClusterEditQuotas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVirtualClusterEditQuotasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **editVirtualClusterQuotasCommand** | [**EditVirtualClusterQuotasCommand**](EditVirtualClusterQuotasCommand.md) |  | 

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


## VirtualClusterEnableAi

> VirtualClusterEnableAi(ctx).VirtualClusterEnableAiCommand(virtualClusterEnableAiCommand).Execute()

Enable ai for virtual cluster

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
	virtualClusterEnableAiCommand := *openapiclient.NewVirtualClusterEnableAiCommand(int32(123), int32(123)) // VirtualClusterEnableAiCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualClusterAPI.VirtualClusterEnableAi(context.Background()).VirtualClusterEnableAiCommand(virtualClusterEnableAiCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualClusterAPI.VirtualClusterEnableAi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVirtualClusterEnableAiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **virtualClusterEnableAiCommand** | [**VirtualClusterEnableAiCommand**](VirtualClusterEnableAiCommand.md) |  | 

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


## VirtualClusterEnableBackup

> VirtualClusterEnableBackup(ctx).VirtualClusterEnableBackupCommand(virtualClusterEnableBackupCommand).Execute()

Enable backup for virtual cluster

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
	virtualClusterEnableBackupCommand := *openapiclient.NewVirtualClusterEnableBackupCommand(int32(123), int32(123)) // VirtualClusterEnableBackupCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualClusterAPI.VirtualClusterEnableBackup(context.Background()).VirtualClusterEnableBackupCommand(virtualClusterEnableBackupCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualClusterAPI.VirtualClusterEnableBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVirtualClusterEnableBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **virtualClusterEnableBackupCommand** | [**VirtualClusterEnableBackupCommand**](VirtualClusterEnableBackupCommand.md) |  | 

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


## VirtualClusterEnableMonitoring

> VirtualClusterEnableMonitoring(ctx).VirtualClusterEnableMonitoringCommand(virtualClusterEnableMonitoringCommand).Execute()

Enable monitoring for virtual cluster

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
	virtualClusterEnableMonitoringCommand := *openapiclient.NewVirtualClusterEnableMonitoringCommand(int32(123)) // VirtualClusterEnableMonitoringCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualClusterAPI.VirtualClusterEnableMonitoring(context.Background()).VirtualClusterEnableMonitoringCommand(virtualClusterEnableMonitoringCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualClusterAPI.VirtualClusterEnableMonitoring``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVirtualClusterEnableMonitoringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **virtualClusterEnableMonitoringCommand** | [**VirtualClusterEnableMonitoringCommand**](VirtualClusterEnableMonitoringCommand.md) |  | 

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


## VirtualClusterEnableOpa

> VirtualClusterEnableOpa(ctx).VirtualClusterEnableOpaCommand(virtualClusterEnableOpaCommand).Execute()

Enable opa for virtual cluster

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
	virtualClusterEnableOpaCommand := *openapiclient.NewVirtualClusterEnableOpaCommand(int32(123), int32(123)) // VirtualClusterEnableOpaCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualClusterAPI.VirtualClusterEnableOpa(context.Background()).VirtualClusterEnableOpaCommand(virtualClusterEnableOpaCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualClusterAPI.VirtualClusterEnableOpa``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVirtualClusterEnableOpaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **virtualClusterEnableOpaCommand** | [**VirtualClusterEnableOpaCommand**](VirtualClusterEnableOpaCommand.md) |  | 

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


## VirtualClusterList

> VClusterList VirtualClusterList(ctx, parentProjectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Id(id).Execute()

Retrieve all vCluster by given project

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
	parentProjectId := int32(56) // int32 | 
	limit := int32(56) // int32 |  (optional)
	offset := int32(56) // int32 |  (optional)
	sortBy := "sortBy_example" // string |  (optional)
	sortDirection := "sortDirection_example" // string |  (optional)
	search := "search_example" // string |  (optional)
	searchId := "searchId_example" // string |  (optional)
	id := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualClusterAPI.VirtualClusterList(context.Background(), parentProjectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualClusterAPI.VirtualClusterList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VirtualClusterList`: VClusterList
	fmt.Fprintf(os.Stdout, "Response from `VirtualClusterAPI.VirtualClusterList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentProjectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVirtualClusterListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **id** | **int32** |  | 

### Return type

[**VClusterList**](VClusterList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VirtualClusterQuota

> KubernetesQuotaListDto VirtualClusterQuota(ctx, projectId).Execute()

Retrieve vcluster quota details

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualClusterAPI.VirtualClusterQuota(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualClusterAPI.VirtualClusterQuota``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VirtualClusterQuota`: KubernetesQuotaListDto
	fmt.Fprintf(os.Stdout, "Response from `VirtualClusterAPI.VirtualClusterQuota`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVirtualClusterQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KubernetesQuotaListDto**](KubernetesQuotaListDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VirtualClusterQuotaPresets

> VirtualClusterQuotasDto VirtualClusterQuotaPresets(ctx).Execute()

Retrieve vCluster quota presets

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
	resp, r, err := apiClient.VirtualClusterAPI.VirtualClusterQuotaPresets(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualClusterAPI.VirtualClusterQuotaPresets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VirtualClusterQuotaPresets`: VirtualClusterQuotasDto
	fmt.Fprintf(os.Stdout, "Response from `VirtualClusterAPI.VirtualClusterQuotaPresets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiVirtualClusterQuotaPresetsRequest struct via the builder pattern


### Return type

[**VirtualClusterQuotasDto**](VirtualClusterQuotasDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VirtualClusterVisibility

> VClusterActionVisibilityDto VirtualClusterVisibility(ctx, projectId).Execute()

Create button condition visibility

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualClusterAPI.VirtualClusterVisibility(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualClusterAPI.VirtualClusterVisibility``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VirtualClusterVisibility`: VClusterActionVisibilityDto
	fmt.Fprintf(os.Stdout, "Response from `VirtualClusterAPI.VirtualClusterVisibility`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVirtualClusterVisibilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VClusterActionVisibilityDto**](VClusterActionVisibilityDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

