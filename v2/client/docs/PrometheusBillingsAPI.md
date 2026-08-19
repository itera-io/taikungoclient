# \PrometheusBillingsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PrometheusbillingsCreate**](PrometheusBillingsAPI.md#PrometheusbillingsCreate) | **Post** /api/v1/prometheusbillings | Add prometheus billing
[**PrometheusbillingsExportCsv**](PrometheusBillingsAPI.md#PrometheusbillingsExportCsv) | **Get** /api/v1/prometheusbillings/export | Export Csv
[**PrometheusbillingsGroupedList**](PrometheusBillingsAPI.md#PrometheusbillingsGroupedList) | **Post** /api/v1/prometheusbillings/grouped | Retrieve a list of grouped prometheus billing
[**PrometheusbillingsList**](PrometheusBillingsAPI.md#PrometheusbillingsList) | **Get** /api/v1/prometheusbillings | Retrieve all prometheus billing



## PrometheusbillingsCreate

> PrometheusbillingsCreate(ctx).PrometheusBillingCreateCommand(prometheusBillingCreateCommand).Execute()

Add prometheus billing

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
	prometheusBillingCreateCommand := *openapiclient.NewPrometheusBillingCreateCommand(int32(123), int32(123), time.Now(), float64(123)) // PrometheusBillingCreateCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PrometheusBillingsAPI.PrometheusbillingsCreate(context.Background()).PrometheusBillingCreateCommand(prometheusBillingCreateCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrometheusBillingsAPI.PrometheusbillingsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrometheusbillingsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **prometheusBillingCreateCommand** | [**PrometheusBillingCreateCommand**](PrometheusBillingCreateCommand.md) |  | 

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


## PrometheusbillingsExportCsv

> CsvExporter PrometheusbillingsExportCsv(ctx).IsEmailEnabled(isEmailEnabled).OrganizationId(organizationId).StartDate(startDate).EndDate(endDate).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrometheusBillingsAPI.PrometheusbillingsExportCsv(context.Background()).IsEmailEnabled(isEmailEnabled).OrganizationId(organizationId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrometheusBillingsAPI.PrometheusbillingsExportCsv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrometheusbillingsExportCsv`: CsvExporter
	fmt.Fprintf(os.Stdout, "Response from `PrometheusBillingsAPI.PrometheusbillingsExportCsv`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrometheusbillingsExportCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isEmailEnabled** | **bool** |  | 
 **organizationId** | **int32** |  | 
 **startDate** | **time.Time** |  | 
 **endDate** | **time.Time** |  | 

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


## PrometheusbillingsGroupedList

> map[string]interface{} PrometheusbillingsGroupedList(ctx).GroupedPrometheusBillingListQuery(groupedPrometheusBillingListQuery).Execute()

Retrieve a list of grouped prometheus billing

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
	groupedPrometheusBillingListQuery := *openapiclient.NewGroupedPrometheusBillingListQuery(NullableInt32(123), map[string]interface{}{"key": interface{}(123)}) // GroupedPrometheusBillingListQuery |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrometheusBillingsAPI.PrometheusbillingsGroupedList(context.Background()).GroupedPrometheusBillingListQuery(groupedPrometheusBillingListQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrometheusBillingsAPI.PrometheusbillingsGroupedList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrometheusbillingsGroupedList`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PrometheusBillingsAPI.PrometheusbillingsGroupedList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrometheusbillingsGroupedListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupedPrometheusBillingListQuery** | [**GroupedPrometheusBillingListQuery**](GroupedPrometheusBillingListQuery.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrometheusbillingsList

> PrometheusBillingInfo PrometheusbillingsList(ctx).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).StartDate(startDate).EndDate(endDate).OrganizationId(organizationId).Execute()

Retrieve all prometheus billing

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrometheusBillingsAPI.PrometheusbillingsList(context.Background()).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).StartDate(startDate).EndDate(endDate).OrganizationId(organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrometheusBillingsAPI.PrometheusbillingsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrometheusbillingsList`: PrometheusBillingInfo
	fmt.Fprintf(os.Stdout, "Response from `PrometheusBillingsAPI.PrometheusbillingsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrometheusbillingsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **startDate** | **time.Time** |  | 
 **endDate** | **time.Time** |  | 
 **organizationId** | **int32** |  | 

### Return type

[**PrometheusBillingInfo**](PrometheusBillingInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

