# \CommonAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommonCountries**](CommonAPI.md#CommonCountries) | **Get** /api/v1/common/countries | Retrieve country list
[**CommonEnumValues**](CommonAPI.md#CommonEnumValues) | **Get** /api/v1/common/enumvalues | Retrieve enum values
[**CommonIpRangeCount**](CommonAPI.md#CommonIpRangeCount) | **Post** /api/v1/common/ip-range-count | Retrieve ip address range count
[**CommonIpRangeList**](CommonAPI.md#CommonIpRangeList) | **Post** /api/v1/common/ip-range-list | Retrieve ip address range list
[**CommonSortingElements**](CommonAPI.md#CommonSortingElements) | **Get** /api/v1/common/sorting-elements/{type} | Retrieve sorting values



## CommonCountries

> []CountryListDto CommonCountries(ctx).Search(search).Execute()

Retrieve country list

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
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommonAPI.CommonCountries(context.Background()).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommonAPI.CommonCountries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommonCountries`: []CountryListDto
	fmt.Fprintf(os.Stdout, "Response from `CommonAPI.CommonCountries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommonCountriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** |  | 

### Return type

[**[]CountryListDto**](CountryListDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommonEnumValues

> EnumList CommonEnumValues(ctx).Execute()

Retrieve enum values

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
	resp, r, err := apiClient.CommonAPI.CommonEnumValues(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommonAPI.CommonEnumValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommonEnumValues`: EnumList
	fmt.Fprintf(os.Stdout, "Response from `CommonAPI.CommonEnumValues`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCommonEnumValuesRequest struct via the builder pattern


### Return type

[**EnumList**](EnumList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommonIpRangeCount

> int32 CommonIpRangeCount(ctx).IpAddressRangeCountCommand(ipAddressRangeCountCommand).Execute()

Retrieve ip address range count

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
	ipAddressRangeCountCommand := *openapiclient.NewIpAddressRangeCountCommand("Begin_example", "End_example") // IpAddressRangeCountCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommonAPI.CommonIpRangeCount(context.Background()).IpAddressRangeCountCommand(ipAddressRangeCountCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommonAPI.CommonIpRangeCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommonIpRangeCount`: int32
	fmt.Fprintf(os.Stdout, "Response from `CommonAPI.CommonIpRangeCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommonIpRangeCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ipAddressRangeCountCommand** | [**IpAddressRangeCountCommand**](IpAddressRangeCountCommand.md) |  | 

### Return type

**int32**

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommonIpRangeList

> []string CommonIpRangeList(ctx).IpAddressRangeListCommand(ipAddressRangeListCommand).Execute()

Retrieve ip address range list

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
	ipAddressRangeListCommand := *openapiclient.NewIpAddressRangeListCommand("IpAddress_example", int32(123), "Gateway_example") // IpAddressRangeListCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommonAPI.CommonIpRangeList(context.Background()).IpAddressRangeListCommand(ipAddressRangeListCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommonAPI.CommonIpRangeList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommonIpRangeList`: []string
	fmt.Fprintf(os.Stdout, "Response from `CommonAPI.CommonIpRangeList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommonIpRangeListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ipAddressRangeListCommand** | [**IpAddressRangeListCommand**](IpAddressRangeListCommand.md) |  | 

### Return type

**[]string**

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommonSortingElements

> []string CommonSortingElements(ctx, type_).Execute()

Retrieve sorting values

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
	type_ := "type__example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommonAPI.CommonSortingElements(context.Background(), type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommonAPI.CommonSortingElements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommonSortingElements`: []string
	fmt.Fprintf(os.Stdout, "Response from `CommonAPI.CommonSortingElements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommonSortingElementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

