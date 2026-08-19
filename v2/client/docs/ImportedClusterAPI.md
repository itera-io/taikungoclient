# \ImportedClusterAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ImportedClusterAsCloudCredential**](ImportedClusterAPI.md#ImportedClusterAsCloudCredential) | **Get** /api/v1/imported-cluster/as-cloud-credential/{projectId} | Imported cluster as cloud credential
[**ImportedClusterAsFullyManaged**](ImportedClusterAPI.md#ImportedClusterAsFullyManaged) | **Get** /api/v1/imported-cluster/as-fully-managed/{projectId} | Imported cluster as fully managed
[**ImportedClusterAsReadOnly**](ImportedClusterAPI.md#ImportedClusterAsReadOnly) | **Get** /api/v1/imported-cluster/as-read-only/{projectId} | Imported cluster as read only
[**ImportedClusterDelete**](ImportedClusterAPI.md#ImportedClusterDelete) | **Post** /api/v1/imported-cluster/delete | Delete imported cluster
[**ImportedClusterDetails**](ImportedClusterAPI.md#ImportedClusterDetails) | **Get** /api/v1/imported-cluster/{id} | Retrieve imported-cluster by given id
[**ImportedClusterDisableAi**](ImportedClusterAPI.md#ImportedClusterDisableAi) | **Post** /api/v1/imported-cluster/disable-ai | Disable ai for imported cluster
[**ImportedClusterDisableBackup**](ImportedClusterAPI.md#ImportedClusterDisableBackup) | **Post** /api/v1/imported-cluster/disable-backup | Disable backup for imported cluster
[**ImportedClusterDisableMonitoring**](ImportedClusterAPI.md#ImportedClusterDisableMonitoring) | **Post** /api/v1/imported-cluster/disable-monitoring | Disable monitoring for imported cluster
[**ImportedClusterDisableOpa**](ImportedClusterAPI.md#ImportedClusterDisableOpa) | **Post** /api/v1/imported-cluster/disable-opa | Disable opa for imported cluster
[**ImportedClusterEnableAi**](ImportedClusterAPI.md#ImportedClusterEnableAi) | **Post** /api/v1/imported-cluster/enable-ai | Enable ai for imported cluster
[**ImportedClusterEnableBackup**](ImportedClusterAPI.md#ImportedClusterEnableBackup) | **Post** /api/v1/imported-cluster/enable-backup | Enable backup for imported cluster
[**ImportedClusterEnableMonitoring**](ImportedClusterAPI.md#ImportedClusterEnableMonitoring) | **Post** /api/v1/imported-cluster/enable-monitoring | Enable monitoring for imported cluster
[**ImportedClusterEnableOpa**](ImportedClusterAPI.md#ImportedClusterEnableOpa) | **Post** /api/v1/imported-cluster/enable-opa | Enable opa for imported cluster



## ImportedClusterAsCloudCredential

> ImportedAsCloudCredentialList ImportedClusterAsCloudCredential(ctx, projectId).Execute()

Imported cluster as cloud credential

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
	resp, r, err := apiClient.ImportedClusterAPI.ImportedClusterAsCloudCredential(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportedClusterAPI.ImportedClusterAsCloudCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportedClusterAsCloudCredential`: ImportedAsCloudCredentialList
	fmt.Fprintf(os.Stdout, "Response from `ImportedClusterAPI.ImportedClusterAsCloudCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportedClusterAsCloudCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ImportedAsCloudCredentialList**](ImportedAsCloudCredentialList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportedClusterAsFullyManaged

> ImportedAsFullyManagedList ImportedClusterAsFullyManaged(ctx, projectId).Execute()

Imported cluster as fully managed

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
	resp, r, err := apiClient.ImportedClusterAPI.ImportedClusterAsFullyManaged(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportedClusterAPI.ImportedClusterAsFullyManaged``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportedClusterAsFullyManaged`: ImportedAsFullyManagedList
	fmt.Fprintf(os.Stdout, "Response from `ImportedClusterAPI.ImportedClusterAsFullyManaged`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportedClusterAsFullyManagedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ImportedAsFullyManagedList**](ImportedAsFullyManagedList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportedClusterAsReadOnly

> ImportedAsReadOnlyList ImportedClusterAsReadOnly(ctx, projectId).Execute()

Imported cluster as read only

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
	resp, r, err := apiClient.ImportedClusterAPI.ImportedClusterAsReadOnly(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportedClusterAPI.ImportedClusterAsReadOnly``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportedClusterAsReadOnly`: ImportedAsReadOnlyList
	fmt.Fprintf(os.Stdout, "Response from `ImportedClusterAPI.ImportedClusterAsReadOnly`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportedClusterAsReadOnlyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ImportedAsReadOnlyList**](ImportedAsReadOnlyList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportedClusterDelete

> ImportedClusterDelete(ctx).DeleteImportedProjectCommand(deleteImportedProjectCommand).Execute()

Delete imported cluster

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
	deleteImportedProjectCommand := *openapiclient.NewDeleteImportedProjectCommand(int32(123), false) // DeleteImportedProjectCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ImportedClusterAPI.ImportedClusterDelete(context.Background()).DeleteImportedProjectCommand(deleteImportedProjectCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportedClusterAPI.ImportedClusterDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportedClusterDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteImportedProjectCommand** | [**DeleteImportedProjectCommand**](DeleteImportedProjectCommand.md) |  | 

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


## ImportedClusterDetails

> ImportedClusterList ImportedClusterDetails(ctx, id).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Execute()

Retrieve imported-cluster by given id

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
	limit := int32(56) // int32 |  (optional)
	offset := int32(56) // int32 |  (optional)
	sortBy := "sortBy_example" // string |  (optional)
	sortDirection := "sortDirection_example" // string |  (optional)
	search := "search_example" // string |  (optional)
	searchId := "searchId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImportedClusterAPI.ImportedClusterDetails(context.Background(), id).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportedClusterAPI.ImportedClusterDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportedClusterDetails`: ImportedClusterList
	fmt.Fprintf(os.Stdout, "Response from `ImportedClusterAPI.ImportedClusterDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportedClusterDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 

### Return type

[**ImportedClusterList**](ImportedClusterList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportedClusterDisableAi

> ImportedClusterDisableAi(ctx).ImportedClusterDisableAiCommand(importedClusterDisableAiCommand).Execute()

Disable ai for imported cluster

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
	importedClusterDisableAiCommand := *openapiclient.NewImportedClusterDisableAiCommand(int32(123)) // ImportedClusterDisableAiCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ImportedClusterAPI.ImportedClusterDisableAi(context.Background()).ImportedClusterDisableAiCommand(importedClusterDisableAiCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportedClusterAPI.ImportedClusterDisableAi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportedClusterDisableAiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importedClusterDisableAiCommand** | [**ImportedClusterDisableAiCommand**](ImportedClusterDisableAiCommand.md) |  | 

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


## ImportedClusterDisableBackup

> ImportedClusterDisableBackup(ctx).ImportedClusterDisableBackupCommand(importedClusterDisableBackupCommand).Execute()

Disable backup for imported cluster

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
	importedClusterDisableBackupCommand := *openapiclient.NewImportedClusterDisableBackupCommand(int32(123)) // ImportedClusterDisableBackupCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ImportedClusterAPI.ImportedClusterDisableBackup(context.Background()).ImportedClusterDisableBackupCommand(importedClusterDisableBackupCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportedClusterAPI.ImportedClusterDisableBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportedClusterDisableBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importedClusterDisableBackupCommand** | [**ImportedClusterDisableBackupCommand**](ImportedClusterDisableBackupCommand.md) |  | 

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


## ImportedClusterDisableMonitoring

> ImportedClusterDisableMonitoring(ctx).ImportedClusterDisableMonitoringCommand(importedClusterDisableMonitoringCommand).Execute()

Disable monitoring for imported cluster

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
	importedClusterDisableMonitoringCommand := *openapiclient.NewImportedClusterDisableMonitoringCommand(int32(123)) // ImportedClusterDisableMonitoringCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ImportedClusterAPI.ImportedClusterDisableMonitoring(context.Background()).ImportedClusterDisableMonitoringCommand(importedClusterDisableMonitoringCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportedClusterAPI.ImportedClusterDisableMonitoring``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportedClusterDisableMonitoringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importedClusterDisableMonitoringCommand** | [**ImportedClusterDisableMonitoringCommand**](ImportedClusterDisableMonitoringCommand.md) |  | 

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


## ImportedClusterDisableOpa

> ImportedClusterDisableOpa(ctx).ImportedClusterDisableOpaCommand(importedClusterDisableOpaCommand).Execute()

Disable opa for imported cluster

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
	importedClusterDisableOpaCommand := *openapiclient.NewImportedClusterDisableOpaCommand(int32(123)) // ImportedClusterDisableOpaCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ImportedClusterAPI.ImportedClusterDisableOpa(context.Background()).ImportedClusterDisableOpaCommand(importedClusterDisableOpaCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportedClusterAPI.ImportedClusterDisableOpa``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportedClusterDisableOpaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importedClusterDisableOpaCommand** | [**ImportedClusterDisableOpaCommand**](ImportedClusterDisableOpaCommand.md) |  | 

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


## ImportedClusterEnableAi

> ImportedClusterEnableAi(ctx).ImportedClusterEnableAiCommand(importedClusterEnableAiCommand).Execute()

Enable ai for imported cluster

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
	importedClusterEnableAiCommand := *openapiclient.NewImportedClusterEnableAiCommand(int32(123), int32(123)) // ImportedClusterEnableAiCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ImportedClusterAPI.ImportedClusterEnableAi(context.Background()).ImportedClusterEnableAiCommand(importedClusterEnableAiCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportedClusterAPI.ImportedClusterEnableAi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportedClusterEnableAiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importedClusterEnableAiCommand** | [**ImportedClusterEnableAiCommand**](ImportedClusterEnableAiCommand.md) |  | 

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


## ImportedClusterEnableBackup

> ImportedClusterEnableBackup(ctx).ImportedClusterEnableBackupCommand(importedClusterEnableBackupCommand).Execute()

Enable backup for imported cluster

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
	importedClusterEnableBackupCommand := *openapiclient.NewImportedClusterEnableBackupCommand(int32(123), int32(123)) // ImportedClusterEnableBackupCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ImportedClusterAPI.ImportedClusterEnableBackup(context.Background()).ImportedClusterEnableBackupCommand(importedClusterEnableBackupCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportedClusterAPI.ImportedClusterEnableBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportedClusterEnableBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importedClusterEnableBackupCommand** | [**ImportedClusterEnableBackupCommand**](ImportedClusterEnableBackupCommand.md) |  | 

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


## ImportedClusterEnableMonitoring

> ImportedClusterEnableMonitoring(ctx).ImportedClusterEnableMonitoringCommand(importedClusterEnableMonitoringCommand).Execute()

Enable monitoring for imported cluster

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
	importedClusterEnableMonitoringCommand := *openapiclient.NewImportedClusterEnableMonitoringCommand(int32(123)) // ImportedClusterEnableMonitoringCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ImportedClusterAPI.ImportedClusterEnableMonitoring(context.Background()).ImportedClusterEnableMonitoringCommand(importedClusterEnableMonitoringCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportedClusterAPI.ImportedClusterEnableMonitoring``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportedClusterEnableMonitoringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importedClusterEnableMonitoringCommand** | [**ImportedClusterEnableMonitoringCommand**](ImportedClusterEnableMonitoringCommand.md) |  | 

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


## ImportedClusterEnableOpa

> ImportedClusterEnableOpa(ctx).ImportedClusterEnableOpaCommand(importedClusterEnableOpaCommand).Execute()

Enable opa for imported cluster

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
	importedClusterEnableOpaCommand := *openapiclient.NewImportedClusterEnableOpaCommand(int32(123), int32(123)) // ImportedClusterEnableOpaCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ImportedClusterAPI.ImportedClusterEnableOpa(context.Background()).ImportedClusterEnableOpaCommand(importedClusterEnableOpaCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportedClusterAPI.ImportedClusterEnableOpa``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportedClusterEnableOpaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importedClusterEnableOpaCommand** | [**ImportedClusterEnableOpaCommand**](ImportedClusterEnableOpaCommand.md) |  | 

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

