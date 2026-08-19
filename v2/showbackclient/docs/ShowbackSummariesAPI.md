# \ShowbackSummariesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ShowbacksummariesByLabel**](ShowbackSummariesAPI.md#ShowbacksummariesByLabel) | **Get** /showback/v1/showbacksummaries/grouped/byLabel | Grouped showback summary by label
[**ShowbacksummariesByProject**](ShowbackSummariesAPI.md#ShowbacksummariesByProject) | **Get** /showback/v1/showbacksummaries/grouped/byProject | Grouped showback summary by project
[**ShowbacksummariesCreate**](ShowbackSummariesAPI.md#ShowbacksummariesCreate) | **Post** /showback/v1/showbacksummaries/create | Create showback summary
[**ShowbacksummariesExportCsv**](ShowbackSummariesAPI.md#ShowbacksummariesExportCsv) | **Get** /showback/v1/showbacksummaries/export | Export Csv
[**ShowbacksummariesGrouped**](ShowbackSummariesAPI.md#ShowbacksummariesGrouped) | **Get** /showback/v1/showbacksummaries/grouped/list | Grouped showback summary
[**ShowbacksummariesGroupedList**](ShowbackSummariesAPI.md#ShowbacksummariesGroupedList) | **Get** /showback/v1/showbacksummaries/grouped | Grouped list of showback summary



## ShowbacksummariesByLabel

> GroupedShowbackByLabelList ShowbacksummariesByLabel(ctx).OrganizationId(organizationId).FromDate(fromDate).ToDate(toDate).Execute()

Grouped showback summary by label

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/itera-io/taikungoclient/v2/showbackclient"
)

func main() {
	organizationId := int32(56) // int32 |  (optional)
	fromDate := time.Now() // time.Time |  (optional)
	toDate := time.Now() // time.Time |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShowbackSummariesAPI.ShowbacksummariesByLabel(context.Background()).OrganizationId(organizationId).FromDate(fromDate).ToDate(toDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShowbackSummariesAPI.ShowbacksummariesByLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowbacksummariesByLabel`: GroupedShowbackByLabelList
	fmt.Fprintf(os.Stdout, "Response from `ShowbackSummariesAPI.ShowbacksummariesByLabel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShowbacksummariesByLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 
 **fromDate** | **time.Time** |  | 
 **toDate** | **time.Time** |  | 

### Return type

[**GroupedShowbackByLabelList**](GroupedShowbackByLabelList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowbacksummariesByProject

> GroupedShowbackByProjectList ShowbacksummariesByProject(ctx).OrganizationId(organizationId).IsDeleted(isDeleted).FromDate(fromDate).ToDate(toDate).Execute()

Grouped showback summary by project

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/itera-io/taikungoclient/v2/showbackclient"
)

func main() {
	organizationId := int32(56) // int32 |  (optional)
	isDeleted := true // bool |  (optional)
	fromDate := time.Now() // time.Time |  (optional)
	toDate := time.Now() // time.Time |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShowbackSummariesAPI.ShowbacksummariesByProject(context.Background()).OrganizationId(organizationId).IsDeleted(isDeleted).FromDate(fromDate).ToDate(toDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShowbackSummariesAPI.ShowbacksummariesByProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowbacksummariesByProject`: GroupedShowbackByProjectList
	fmt.Fprintf(os.Stdout, "Response from `ShowbackSummariesAPI.ShowbacksummariesByProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShowbacksummariesByProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 
 **isDeleted** | **bool** |  | 
 **fromDate** | **time.Time** |  | 
 **toDate** | **time.Time** |  | 

### Return type

[**GroupedShowbackByProjectList**](GroupedShowbackByProjectList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowbacksummariesCreate

> ShowbacksummariesCreate(ctx).CreateShowbackSummaryCommand(createShowbackSummaryCommand).Execute()

Create showback summary

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/itera-io/taikungoclient/v2/showbackclient"
)

func main() {
	createShowbackSummaryCommand := *openapiclient.NewCreateShowbackSummaryCommand() // CreateShowbackSummaryCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ShowbackSummariesAPI.ShowbacksummariesCreate(context.Background()).CreateShowbackSummaryCommand(createShowbackSummaryCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShowbackSummariesAPI.ShowbacksummariesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShowbacksummariesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createShowbackSummaryCommand** | [**CreateShowbackSummaryCommand**](CreateShowbackSummaryCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowbacksummariesExportCsv

> CsvExporter ShowbacksummariesExportCsv(ctx).IsEmailEnabled(isEmailEnabled).OrganizationId(organizationId).Execute()

Export Csv

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/itera-io/taikungoclient/v2/showbackclient"
)

func main() {
	isEmailEnabled := true // bool | 
	organizationId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShowbackSummariesAPI.ShowbacksummariesExportCsv(context.Background()).IsEmailEnabled(isEmailEnabled).OrganizationId(organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShowbackSummariesAPI.ShowbacksummariesExportCsv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowbacksummariesExportCsv`: CsvExporter
	fmt.Fprintf(os.Stdout, "Response from `ShowbackSummariesAPI.ShowbacksummariesExportCsv`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShowbacksummariesExportCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isEmailEnabled** | **bool** |  | 
 **organizationId** | **int32** |  | 

### Return type

[**CsvExporter**](CsvExporter.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowbacksummariesGrouped

> []GroupedShowbackSummaryListDto ShowbacksummariesGrouped(ctx).OrganizationId(organizationId).FromDate(fromDate).ToDate(toDate).IsDeleted(isDeleted).Execute()

Grouped showback summary

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/itera-io/taikungoclient/v2/showbackclient"
)

func main() {
	organizationId := int32(56) // int32 |  (optional)
	fromDate := time.Now() // time.Time |  (optional)
	toDate := time.Now() // time.Time |  (optional)
	isDeleted := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShowbackSummariesAPI.ShowbacksummariesGrouped(context.Background()).OrganizationId(organizationId).FromDate(fromDate).ToDate(toDate).IsDeleted(isDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShowbackSummariesAPI.ShowbacksummariesGrouped``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowbacksummariesGrouped`: []GroupedShowbackSummaryListDto
	fmt.Fprintf(os.Stdout, "Response from `ShowbackSummariesAPI.ShowbacksummariesGrouped`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShowbacksummariesGroupedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 
 **fromDate** | **time.Time** |  | 
 **toDate** | **time.Time** |  | 
 **isDeleted** | **bool** |  | 

### Return type

[**[]GroupedShowbackSummaryListDto**](GroupedShowbackSummaryListDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowbacksummariesGroupedList

> GroupedShowbackList ShowbacksummariesGroupedList(ctx).OrganizationId(organizationId).PeriodDuration(periodDuration).IsDeleted(isDeleted).FromDate(fromDate).ToDate(toDate).Execute()

Grouped list of showback summary

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/itera-io/taikungoclient/v2/showbackclient"
)

func main() {
	organizationId := int32(56) // int32 |  (optional)
	periodDuration := "periodDuration_example" // string |  (optional)
	isDeleted := true // bool |  (optional)
	fromDate := time.Now() // time.Time |  (optional)
	toDate := time.Now() // time.Time |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShowbackSummariesAPI.ShowbacksummariesGroupedList(context.Background()).OrganizationId(organizationId).PeriodDuration(periodDuration).IsDeleted(isDeleted).FromDate(fromDate).ToDate(toDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShowbackSummariesAPI.ShowbacksummariesGroupedList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowbacksummariesGroupedList`: GroupedShowbackList
	fmt.Fprintf(os.Stdout, "Response from `ShowbackSummariesAPI.ShowbacksummariesGroupedList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShowbacksummariesGroupedListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 
 **periodDuration** | **string** |  | 
 **isDeleted** | **bool** |  | 
 **fromDate** | **time.Time** |  | 
 **toDate** | **time.Time** |  | 

### Return type

[**GroupedShowbackList**](GroupedShowbackList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

