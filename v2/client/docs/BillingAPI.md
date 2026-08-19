# \BillingAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BillingCreate**](BillingAPI.md#BillingCreate) | **Post** /api/v1/billing/create | Add billing summary
[**BillingExportCsv**](BillingAPI.md#BillingExportCsv) | **Get** /api/v1/billing/export | Export Csv
[**BillingGroupedList**](BillingAPI.md#BillingGroupedList) | **Post** /api/v1/billing/grouped | Retrieve a grouped list of billing summaries
[**BillingList**](BillingAPI.md#BillingList) | **Get** /api/v1/billing | Retrieve billing info



## BillingCreate

> BillingCreate(ctx).CreateBillingSummaryCommand(createBillingSummaryCommand).Execute()

Add billing summary

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
	createBillingSummaryCommand := *openapiclient.NewCreateBillingSummaryCommand() // CreateBillingSummaryCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BillingAPI.BillingCreate(context.Background()).CreateBillingSummaryCommand(createBillingSummaryCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBillingCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBillingSummaryCommand** | [**CreateBillingSummaryCommand**](CreateBillingSummaryCommand.md) |  | 

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


## BillingExportCsv

> CsvExporter BillingExportCsv(ctx).IsEmailEnabled(isEmailEnabled).OrganizationId(organizationId).StartDate(startDate).EndDate(endDate).IsDeleted(isDeleted).Execute()

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
	organizationId := int32(56) // int32 |  (optional)
	startDate := time.Now() // time.Time |  (optional)
	endDate := time.Now() // time.Time |  (optional)
	isDeleted := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingExportCsv(context.Background()).IsEmailEnabled(isEmailEnabled).OrganizationId(organizationId).StartDate(startDate).EndDate(endDate).IsDeleted(isDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingExportCsv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingExportCsv`: CsvExporter
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingExportCsv`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBillingExportCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isEmailEnabled** | **bool** |  | 
 **organizationId** | **int32** |  | 
 **startDate** | **time.Time** |  | 
 **endDate** | **time.Time** |  | 
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


## BillingGroupedList

> []GroupedBillingInfo BillingGroupedList(ctx).GroupedBillingListQuery(groupedBillingListQuery).Execute()

Retrieve a grouped list of billing summaries

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
	groupedBillingListQuery := *openapiclient.NewGroupedBillingListQuery(NullableInt32(123), map[string]interface{}{"key": interface{}(123)}, false) // GroupedBillingListQuery |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingGroupedList(context.Background()).GroupedBillingListQuery(groupedBillingListQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingGroupedList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingGroupedList`: []GroupedBillingInfo
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingGroupedList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBillingGroupedListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupedBillingListQuery** | [**GroupedBillingListQuery**](GroupedBillingListQuery.md) |  | 

### Return type

[**[]GroupedBillingInfo**](GroupedBillingInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingList

> BillingInfo BillingList(ctx).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).StartDate(startDate).EndDate(endDate).OrganizationId(organizationId).IsDeleted(isDeleted).ProjectId(projectId).Execute()

Retrieve billing info

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
	isDeleted := true // bool |  (optional)
	projectId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingList(context.Background()).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).StartDate(startDate).EndDate(endDate).OrganizationId(organizationId).IsDeleted(isDeleted).ProjectId(projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingList`: BillingInfo
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBillingListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **startDate** | **time.Time** |  | 
 **endDate** | **time.Time** |  | 
 **organizationId** | **int32** |  | 
 **isDeleted** | **bool** |  | 
 **projectId** | **int32** |  | 

### Return type

[**BillingInfo**](BillingInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

