# \BackupPolicyAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BackupByName**](BackupPolicyAPI.md#BackupByName) | **Get** /api/v1/backup/{projectId}/{name} | Get backup info by name
[**BackupCreate**](BackupPolicyAPI.md#BackupCreate) | **Post** /api/v1/backup/create | Add backup policy
[**BackupDeleteBackup**](BackupPolicyAPI.md#BackupDeleteBackup) | **Post** /api/v1/backup/delete/backup | Remove policy backup
[**BackupDeleteBackupLocation**](BackupPolicyAPI.md#BackupDeleteBackupLocation) | **Post** /api/v1/backup/delete/location | Remove backup location from project
[**BackupDeleteRestore**](BackupPolicyAPI.md#BackupDeleteRestore) | **Post** /api/v1/backup/delete/restore | Remove policy restore
[**BackupDeleteSchedule**](BackupPolicyAPI.md#BackupDeleteSchedule) | **Post** /api/v1/backup/delete/schedule | Remove policy schedule
[**BackupDescribeBackup**](BackupPolicyAPI.md#BackupDescribeBackup) | **Get** /api/v1/backup/describe/backup/{projectId}/{name} | Get backup info by name
[**BackupDescribeRestore**](BackupPolicyAPI.md#BackupDescribeRestore) | **Get** /api/v1/backup/describe/restore/{projectId}/{name} | Get restore info by name
[**BackupDescribeSchedule**](BackupPolicyAPI.md#BackupDescribeSchedule) | **Get** /api/v1/backup/describe/schedule/{projectId}/{name} | Get schedule info by name
[**BackupImportBackupStorage**](BackupPolicyAPI.md#BackupImportBackupStorage) | **Post** /api/v1/backup/location | Import backup storage from source project to target project
[**BackupListAllBackupStorages**](BackupPolicyAPI.md#BackupListAllBackupStorages) | **Get** /api/v1/backup/location/{projectId} | List all backup locations
[**BackupListAllBackups**](BackupPolicyAPI.md#BackupListAllBackups) | **Get** /api/v1/backup/backups/{projectId} | List all backups
[**BackupListAllDeleteBackupRequests**](BackupPolicyAPI.md#BackupListAllDeleteBackupRequests) | **Get** /api/v1/backup/delete-requests/{projectId} | List all delete backup requests
[**BackupListAllRestores**](BackupPolicyAPI.md#BackupListAllRestores) | **Get** /api/v1/backup/restores/{projectId} | List all restores
[**BackupListAllSchedules**](BackupPolicyAPI.md#BackupListAllSchedules) | **Get** /api/v1/backup/schedules/{projectId} | List all schedules
[**BackupRestoreBackup**](BackupPolicyAPI.md#BackupRestoreBackup) | **Post** /api/v1/backup/restore | Restore backup



## BackupByName

> BackupDto BackupByName(ctx, projectId, name).Execute()

Get backup info by name

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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupPolicyAPI.BackupByName(context.Background(), projectId, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyAPI.BackupByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackupByName`: BackupDto
	fmt.Fprintf(os.Stdout, "Response from `BackupPolicyAPI.BackupByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BackupDto**](BackupDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupCreate

> BackupCreate(ctx).CreateBackupPolicyCommand(createBackupPolicyCommand).Execute()

Add backup policy

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
	createBackupPolicyCommand := *openapiclient.NewCreateBackupPolicyCommand("Name_example", []string{"IncludeNamespaces_example"}, "CronPeriod_example", int32(123)) // CreateBackupPolicyCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackupPolicyAPI.BackupCreate(context.Background()).CreateBackupPolicyCommand(createBackupPolicyCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyAPI.BackupCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackupCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBackupPolicyCommand** | [**CreateBackupPolicyCommand**](CreateBackupPolicyCommand.md) |  | 

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


## BackupDeleteBackup

> BackupDeleteBackup(ctx).DeleteBackupCommand(deleteBackupCommand).Execute()

Remove policy backup

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
	deleteBackupCommand := *openapiclient.NewDeleteBackupCommand(int32(123), "Name_example") // DeleteBackupCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackupPolicyAPI.BackupDeleteBackup(context.Background()).DeleteBackupCommand(deleteBackupCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyAPI.BackupDeleteBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackupDeleteBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteBackupCommand** | [**DeleteBackupCommand**](DeleteBackupCommand.md) |  | 

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


## BackupDeleteBackupLocation

> BackupDeleteBackupLocation(ctx).DeleteBackupStorageLocationCommand(deleteBackupStorageLocationCommand).Execute()

Remove backup location from project

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
	deleteBackupStorageLocationCommand := *openapiclient.NewDeleteBackupStorageLocationCommand(int32(123), "StorageLocation_example") // DeleteBackupStorageLocationCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackupPolicyAPI.BackupDeleteBackupLocation(context.Background()).DeleteBackupStorageLocationCommand(deleteBackupStorageLocationCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyAPI.BackupDeleteBackupLocation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackupDeleteBackupLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteBackupStorageLocationCommand** | [**DeleteBackupStorageLocationCommand**](DeleteBackupStorageLocationCommand.md) |  | 

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


## BackupDeleteRestore

> BackupDeleteRestore(ctx).DeleteRestoreCommand(deleteRestoreCommand).Execute()

Remove policy restore

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
	deleteRestoreCommand := *openapiclient.NewDeleteRestoreCommand(int32(123), "Name_example") // DeleteRestoreCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackupPolicyAPI.BackupDeleteRestore(context.Background()).DeleteRestoreCommand(deleteRestoreCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyAPI.BackupDeleteRestore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackupDeleteRestoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteRestoreCommand** | [**DeleteRestoreCommand**](DeleteRestoreCommand.md) |  | 

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


## BackupDeleteSchedule

> BackupDeleteSchedule(ctx).DeleteScheduleCommand(deleteScheduleCommand).Execute()

Remove policy schedule

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
	deleteScheduleCommand := *openapiclient.NewDeleteScheduleCommand(int32(123), "Name_example") // DeleteScheduleCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackupPolicyAPI.BackupDeleteSchedule(context.Background()).DeleteScheduleCommand(deleteScheduleCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyAPI.BackupDeleteSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackupDeleteScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteScheduleCommand** | [**DeleteScheduleCommand**](DeleteScheduleCommand.md) |  | 

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


## BackupDescribeBackup

> string BackupDescribeBackup(ctx, projectId, name).Execute()

Get backup info by name

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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupPolicyAPI.BackupDescribeBackup(context.Background(), projectId, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyAPI.BackupDescribeBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackupDescribeBackup`: string
	fmt.Fprintf(os.Stdout, "Response from `BackupPolicyAPI.BackupDescribeBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupDescribeBackupRequest struct via the builder pattern


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


## BackupDescribeRestore

> string BackupDescribeRestore(ctx, projectId, name).Execute()

Get restore info by name

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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupPolicyAPI.BackupDescribeRestore(context.Background(), projectId, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyAPI.BackupDescribeRestore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackupDescribeRestore`: string
	fmt.Fprintf(os.Stdout, "Response from `BackupPolicyAPI.BackupDescribeRestore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupDescribeRestoreRequest struct via the builder pattern


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


## BackupDescribeSchedule

> string BackupDescribeSchedule(ctx, projectId, name).Execute()

Get schedule info by name

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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupPolicyAPI.BackupDescribeSchedule(context.Background(), projectId, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyAPI.BackupDescribeSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackupDescribeSchedule`: string
	fmt.Fprintf(os.Stdout, "Response from `BackupPolicyAPI.BackupDescribeSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupDescribeScheduleRequest struct via the builder pattern


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


## BackupImportBackupStorage

> BackupImportBackupStorage(ctx).ImportBackupStorageLocationCommand(importBackupStorageLocationCommand).Execute()

Import backup storage from source project to target project

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
	importBackupStorageLocationCommand := *openapiclient.NewImportBackupStorageLocationCommand(int32(123), int32(123)) // ImportBackupStorageLocationCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackupPolicyAPI.BackupImportBackupStorage(context.Background()).ImportBackupStorageLocationCommand(importBackupStorageLocationCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyAPI.BackupImportBackupStorage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackupImportBackupStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importBackupStorageLocationCommand** | [**ImportBackupStorageLocationCommand**](ImportBackupStorageLocationCommand.md) |  | 

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


## BackupListAllBackupStorages

> ListAllBackupStorageLocations BackupListAllBackupStorages(ctx, projectId).SortBy(sortBy).SortDirection(sortDirection).Search(search).Limit(limit).Offset(offset).Execute()

List all backup locations

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
	search := "search_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)
	offset := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupPolicyAPI.BackupListAllBackupStorages(context.Background(), projectId).SortBy(sortBy).SortDirection(sortDirection).Search(search).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyAPI.BackupListAllBackupStorages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackupListAllBackupStorages`: ListAllBackupStorageLocations
	fmt.Fprintf(os.Stdout, "Response from `BackupPolicyAPI.BackupListAllBackupStorages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupListAllBackupStoragesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **limit** | **int32** |  | [default to 50]
 **offset** | **int32** |  | [default to 0]

### Return type

[**ListAllBackupStorageLocations**](ListAllBackupStorageLocations.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupListAllBackups

> ListAllBackups BackupListAllBackups(ctx, projectId).SortBy(sortBy).SortDirection(sortDirection).Search(search).Limit(limit).Offset(offset).Execute()

List all backups

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
	search := "search_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)
	offset := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupPolicyAPI.BackupListAllBackups(context.Background(), projectId).SortBy(sortBy).SortDirection(sortDirection).Search(search).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyAPI.BackupListAllBackups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackupListAllBackups`: ListAllBackups
	fmt.Fprintf(os.Stdout, "Response from `BackupPolicyAPI.BackupListAllBackups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupListAllBackupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **limit** | **int32** |  | [default to 50]
 **offset** | **int32** |  | [default to 0]

### Return type

[**ListAllBackups**](ListAllBackups.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupListAllDeleteBackupRequests

> ListAllDeleteBackupRequests BackupListAllDeleteBackupRequests(ctx, projectId).SortBy(sortBy).SortDirection(sortDirection).Search(search).Limit(limit).Offset(offset).Execute()

List all delete backup requests

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
	search := "search_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)
	offset := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupPolicyAPI.BackupListAllDeleteBackupRequests(context.Background(), projectId).SortBy(sortBy).SortDirection(sortDirection).Search(search).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyAPI.BackupListAllDeleteBackupRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackupListAllDeleteBackupRequests`: ListAllDeleteBackupRequests
	fmt.Fprintf(os.Stdout, "Response from `BackupPolicyAPI.BackupListAllDeleteBackupRequests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupListAllDeleteBackupRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **limit** | **int32** |  | [default to 50]
 **offset** | **int32** |  | [default to 0]

### Return type

[**ListAllDeleteBackupRequests**](ListAllDeleteBackupRequests.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupListAllRestores

> ListAllRestores BackupListAllRestores(ctx, projectId).SortBy(sortBy).SortDirection(sortDirection).Search(search).Limit(limit).Offset(offset).Execute()

List all restores

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
	search := "search_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)
	offset := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupPolicyAPI.BackupListAllRestores(context.Background(), projectId).SortBy(sortBy).SortDirection(sortDirection).Search(search).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyAPI.BackupListAllRestores``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackupListAllRestores`: ListAllRestores
	fmt.Fprintf(os.Stdout, "Response from `BackupPolicyAPI.BackupListAllRestores`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupListAllRestoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **limit** | **int32** |  | [default to 50]
 **offset** | **int32** |  | [default to 0]

### Return type

[**ListAllRestores**](ListAllRestores.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupListAllSchedules

> ListAllSchedules BackupListAllSchedules(ctx, projectId).SortBy(sortBy).SortDirection(sortDirection).Search(search).Limit(limit).Offset(offset).Execute()

List all schedules

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
	search := "search_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)
	offset := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupPolicyAPI.BackupListAllSchedules(context.Background(), projectId).SortBy(sortBy).SortDirection(sortDirection).Search(search).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyAPI.BackupListAllSchedules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackupListAllSchedules`: ListAllSchedules
	fmt.Fprintf(os.Stdout, "Response from `BackupPolicyAPI.BackupListAllSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupListAllSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **limit** | **int32** |  | [default to 50]
 **offset** | **int32** |  | [default to 0]

### Return type

[**ListAllSchedules**](ListAllSchedules.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupRestoreBackup

> BackupRestoreBackup(ctx).RestoreBackupCommand(restoreBackupCommand).Execute()

Restore backup

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
	restoreBackupCommand := *openapiclient.NewRestoreBackupCommand(int32(123), "BackupName_example", "RestoreName_example", []string{"IncludeNamespaces_example"}, []string{"ExcludeNamespaces_example"}) // RestoreBackupCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackupPolicyAPI.BackupRestoreBackup(context.Background()).RestoreBackupCommand(restoreBackupCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupPolicyAPI.BackupRestoreBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restoreBackupCommand** | [**RestoreBackupCommand**](RestoreBackupCommand.md) |  | 

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

