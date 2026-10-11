# \OrganizationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationsAddPrometheusrules**](OrganizationsAPI.md#OrganizationsAddPrometheusrules) | **Post** /api/v1/organizations/{id}/prometheusrules | Add prometheus rule(s) to organization
[**OrganizationsCreate**](OrganizationsAPI.md#OrganizationsCreate) | **Post** /api/v1/organizations | Add a new organization. Only available for admins.
[**OrganizationsDelete**](OrganizationsAPI.md#OrganizationsDelete) | **Delete** /api/v1/organizations/{id} | Delete the specified organization. Only available for global roles.
[**OrganizationsDeletePrometheusrules**](OrganizationsAPI.md#OrganizationsDeletePrometheusrules) | **Put** /api/v1/organizations/{id}/prometheusrules | Unbind prometheus rule(s) from organization
[**OrganizationsDetails**](OrganizationsAPI.md#OrganizationsDetails) | **Get** /api/v1/organizations/details | Retrieve all data about current organization by Id
[**OrganizationsExportCsv**](OrganizationsAPI.md#OrganizationsExportCsv) | **Get** /api/v1/organizations/export | Export Csv file
[**OrganizationsList**](OrganizationsAPI.md#OrganizationsList) | **Get** /api/v1/organizations | Retrieve all organizations
[**OrganizationsOrganizationList**](OrganizationsAPI.md#OrganizationsOrganizationList) | **Get** /api/v1/organizations/list | Retrieve organizations
[**OrganizationsUpdate**](OrganizationsAPI.md#OrganizationsUpdate) | **Post** /api/v1/organizations/update | Update organization by Id



## OrganizationsAddPrometheusrules

> OrganizationsAddPrometheusrules(ctx, id).AddPrometheusRulesToOrganizationDto(addPrometheusRulesToOrganizationDto).Execute()

Add prometheus rule(s) to organization

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
	addPrometheusRulesToOrganizationDto := []openapiclient.AddPrometheusRulesToOrganizationDto{*openapiclient.NewAddPrometheusRulesToOrganizationDto()} // []AddPrometheusRulesToOrganizationDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.OrganizationsAddPrometheusrules(context.Background(), id).AddPrometheusRulesToOrganizationDto(addPrometheusRulesToOrganizationDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.OrganizationsAddPrometheusrules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsAddPrometheusrulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addPrometheusRulesToOrganizationDto** | [**[]AddPrometheusRulesToOrganizationDto**](AddPrometheusRulesToOrganizationDto.md) |  | 

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


## OrganizationsCreate

> ApiResponse OrganizationsCreate(ctx).OrganizationCreateCommand(organizationCreateCommand).Execute()

Add a new organization. Only available for admins.

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
	organizationCreateCommand := *openapiclient.NewOrganizationCreateCommand() // OrganizationCreateCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.OrganizationsCreate(context.Background()).OrganizationCreateCommand(organizationCreateCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.OrganizationsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationsCreate`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.OrganizationsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationCreateCommand** | [**OrganizationCreateCommand**](OrganizationCreateCommand.md) |  | 

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


## OrganizationsDelete

> OrganizationsDelete(ctx, id).Execute()

Delete the specified organization. Only available for global roles.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.OrganizationsDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.OrganizationsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## OrganizationsDeletePrometheusrules

> OrganizationsDeletePrometheusrules(ctx, id).RequestBody(requestBody).Execute()

Unbind prometheus rule(s) from organization

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
	requestBody := []int32{int32(123)} // []int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.OrganizationsDeletePrometheusrules(context.Background(), id).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.OrganizationsDeletePrometheusrules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsDeletePrometheusrulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]int32** |  | 

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


## OrganizationsDetails

> DashboardChart OrganizationsDetails(ctx).OrganizationId(organizationId).Execute()

Retrieve all data about current organization by Id

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
	organizationId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.OrganizationsDetails(context.Background()).OrganizationId(organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.OrganizationsDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationsDetails`: DashboardChart
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.OrganizationsDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 

### Return type

[**DashboardChart**](DashboardChart.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsExportCsv

> CsvExporter OrganizationsExportCsv(ctx).Execute()

Export Csv file

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
	resp, r, err := apiClient.OrganizationsAPI.OrganizationsExportCsv(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.OrganizationsExportCsv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationsExportCsv`: CsvExporter
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.OrganizationsExportCsv`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsExportCsvRequest struct via the builder pattern


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


## OrganizationsList

> OrganizationsList OrganizationsList(ctx).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Id(id).Execute()

Retrieve all organizations

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
	sortBy := "sortBy_example" // string |  (optional)
	sortDirection := "sortDirection_example" // string |  (optional)
	search := "search_example" // string |  (optional)
	searchId := "searchId_example" // string |  (optional)
	id := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.OrganizationsList(context.Background()).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.OrganizationsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationsList`: OrganizationsList
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.OrganizationsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsListRequest struct via the builder pattern


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

[**OrganizationsList**](OrganizationsList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationList

> []OrganizationDropdownDto OrganizationsOrganizationList(ctx).DomainId(domainId).Search(search).IsInfra(isInfra).PrometheusRuleId(prometheusRuleId).Execute()

Retrieve organizations

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
	domainId := int32(56) // int32 |  (optional)
	search := "search_example" // string |  (optional)
	isInfra := true // bool |  (optional)
	prometheusRuleId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.OrganizationsOrganizationList(context.Background()).DomainId(domainId).Search(search).IsInfra(isInfra).PrometheusRuleId(prometheusRuleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.OrganizationsOrganizationList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationsOrganizationList`: []OrganizationDropdownDto
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.OrganizationsOrganizationList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrganizationListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainId** | **int32** |  | 
 **search** | **string** |  | 
 **isInfra** | **bool** |  | 
 **prometheusRuleId** | **int32** |  | 

### Return type

[**[]OrganizationDropdownDto**](OrganizationDropdownDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsUpdate

> OrganizationsUpdate(ctx).UpdateOrganizationCommand(updateOrganizationCommand).Execute()

Update organization by Id

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
	updateOrganizationCommand := *openapiclient.NewUpdateOrganizationCommand() // UpdateOrganizationCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.OrganizationsUpdate(context.Background()).UpdateOrganizationCommand(updateOrganizationCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.OrganizationsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateOrganizationCommand** | [**UpdateOrganizationCommand**](UpdateOrganizationCommand.md) |  | 

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

