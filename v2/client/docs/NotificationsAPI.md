# \NotificationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NotificationsExportCsv**](NotificationsAPI.md#NotificationsExportCsv) | **Get** /api/v1/notifications/download | Export Csv
[**NotificationsList**](NotificationsAPI.md#NotificationsList) | **Get** /api/v1/notifications/list | Retrieve all notifications
[**NotificationsNotifyOwner**](NotificationsAPI.md#NotificationsNotifyOwner) | **Post** /api/v1/notifications/notifyowner | Notify owner



## NotificationsExportCsv

> CsvExporter NotificationsExportCsv(ctx).IsEmailEnabled(isEmailEnabled).SortBy(sortBy).SortDirection(sortDirection).StartDate(startDate).EndDate(endDate).OrganizationId(organizationId).FilterBy(filterBy).ProjectId(projectId).UserId(userId).IsDeleted(isDeleted).Execute()

Export Csv

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
	isEmailEnabled := true // bool | 
	sortBy := "sortBy_example" // string |  (optional)
	sortDirection := "sortDirection_example" // string |  (optional)
	startDate := time.Now() // time.Time |  (optional)
	endDate := time.Now() // time.Time |  (optional)
	organizationId := int32(56) // int32 |  (optional)
	filterBy := "filterBy_example" // string |  (optional)
	projectId := int32(56) // int32 |  (optional)
	userId := "userId_example" // string |  (optional)
	isDeleted := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.NotificationsExportCsv(context.Background()).IsEmailEnabled(isEmailEnabled).SortBy(sortBy).SortDirection(sortDirection).StartDate(startDate).EndDate(endDate).OrganizationId(organizationId).FilterBy(filterBy).ProjectId(projectId).UserId(userId).IsDeleted(isDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.NotificationsExportCsv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotificationsExportCsv`: CsvExporter
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.NotificationsExportCsv`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNotificationsExportCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isEmailEnabled** | **bool** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **startDate** | **time.Time** |  | 
 **endDate** | **time.Time** |  | 
 **organizationId** | **int32** |  | 
 **filterBy** | **string** |  | 
 **projectId** | **int32** |  | 
 **userId** | **string** |  | 
 **isDeleted** | **bool** |  | 

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


## NotificationsList

> NotificationHistory NotificationsList(ctx).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).StartDate(startDate).EndDate(endDate).OrganizationId(organizationId).FilterBy(filterBy).ProjectId(projectId).UserId(userId).IsDeleted(isDeleted).Search(search).Execute()

Retrieve all notifications

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
	limit := int32(56) // int32 |  (optional)
	offset := int32(56) // int32 |  (optional)
	sortBy := "sortBy_example" // string |  (optional)
	sortDirection := "sortDirection_example" // string |  (optional)
	startDate := time.Now() // time.Time |  (optional)
	endDate := time.Now() // time.Time |  (optional)
	organizationId := int32(56) // int32 |  (optional)
	filterBy := "filterBy_example" // string |  (optional)
	projectId := int32(56) // int32 |  (optional)
	userId := "userId_example" // string |  (optional)
	isDeleted := true // bool |  (optional)
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.NotificationsList(context.Background()).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).StartDate(startDate).EndDate(endDate).OrganizationId(organizationId).FilterBy(filterBy).ProjectId(projectId).UserId(userId).IsDeleted(isDeleted).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.NotificationsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotificationsList`: NotificationHistory
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.NotificationsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNotificationsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **startDate** | **time.Time** |  | 
 **endDate** | **time.Time** |  | 
 **organizationId** | **int32** |  | 
 **filterBy** | **string** |  | 
 **projectId** | **int32** |  | 
 **userId** | **string** |  | 
 **isDeleted** | **bool** |  | 
 **search** | **string** |  | 

### Return type

[**NotificationHistory**](NotificationHistory.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotificationsNotifyOwner

> NotificationsNotifyOwner(ctx).Command(command).Execute()

Notify owner

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
	command := "command_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NotificationsAPI.NotificationsNotifyOwner(context.Background()).Command(command).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.NotificationsNotifyOwner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNotificationsNotifyOwnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **command** | **string** |  | 

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

