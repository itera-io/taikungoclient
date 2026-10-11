# \AzureCloudCredentialAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AzureAksClusters**](AzureCloudCredentialAPI.md#AzureAksClusters) | **Post** /api/v1/azure/aks-clusters | Fetch AKS cluster list
[**AzureCreate**](AzureCloudCredentialAPI.md#AzureCreate) | **Post** /api/v1/azure/create | Add Azure credentials
[**AzureDashboard**](AzureCloudCredentialAPI.md#AzureDashboard) | **Post** /api/v1/azure/quota/list | Fetch Azure quota list
[**AzureFileShareList**](AzureCloudCredentialAPI.md#AzureFileShareList) | **Post** /api/v1/azure/file-share-list | Retrieve Azure File Shares within a resource group
[**AzureList**](AzureCloudCredentialAPI.md#AzureList) | **Get** /api/v1/azure/list | Retrieve list of azure cloud credentials
[**AzureLocations**](AzureCloudCredentialAPI.md#AzureLocations) | **Post** /api/v1/azure/locations | Fetch Azure location list
[**AzureOffers**](AzureCloudCredentialAPI.md#AzureOffers) | **Get** /api/v1/azure/offers/{cloudId}/{publisher} | List Azure offer list by publisher
[**AzurePublishers**](AzureCloudCredentialAPI.md#AzurePublishers) | **Get** /api/v1/azure/publishers/{cloudId} | List Azure publishers list
[**AzureSkus**](AzureCloudCredentialAPI.md#AzureSkus) | **Get** /api/v1/azure/skus/{cloudId}/{publisher}/{offer} | List Azure skus list by publisher and offer
[**AzureSubscriptions**](AzureCloudCredentialAPI.md#AzureSubscriptions) | **Post** /api/v1/azure/subscriptions | Azure subscriptions list
[**AzureUpdate**](AzureCloudCredentialAPI.md#AzureUpdate) | **Post** /api/v1/azure/update | Update Azure credentials
[**AzureVnetList**](AzureCloudCredentialAPI.md#AzureVnetList) | **Post** /api/v1/azure/vnet-list | Retrieve VNets within an Azure resource group
[**AzureVnetResourceGroupList**](AzureCloudCredentialAPI.md#AzureVnetResourceGroupList) | **Post** /api/v1/azure/vnet-resource-group-list | Retrieve Azure resource groups that contain VNets
[**AzureVnetSubnetList**](AzureCloudCredentialAPI.md#AzureVnetSubnetList) | **Post** /api/v1/azure/vnet-subnet-list | Retrieve subnets within an Azure VNet
[**AzureZones**](AzureCloudCredentialAPI.md#AzureZones) | **Post** /api/v1/azure/zones | Fetch Azure zone list



## AzureAksClusters

> []AzureAksClusterDto AzureAksClusters(ctx).AksClusterListCommand(aksClusterListCommand).Execute()

Fetch AKS cluster list

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
	aksClusterListCommand := *openapiclient.NewAksClusterListCommand(int32(123)) // AksClusterListCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AzureCloudCredentialAPI.AzureAksClusters(context.Background()).AksClusterListCommand(aksClusterListCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AzureCloudCredentialAPI.AzureAksClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AzureAksClusters`: []AzureAksClusterDto
	fmt.Fprintf(os.Stdout, "Response from `AzureCloudCredentialAPI.AzureAksClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAzureAksClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aksClusterListCommand** | [**AksClusterListCommand**](AksClusterListCommand.md) |  | 

### Return type

[**[]AzureAksClusterDto**](AzureAksClusterDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AzureCreate

> ApiResponse AzureCreate(ctx).CreateAzureCloudCommand(createAzureCloudCommand).Execute()

Add Azure credentials

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
	createAzureCloudCommand := *openapiclient.NewCreateAzureCloudCommand() // CreateAzureCloudCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AzureCloudCredentialAPI.AzureCreate(context.Background()).CreateAzureCloudCommand(createAzureCloudCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AzureCloudCredentialAPI.AzureCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AzureCreate`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AzureCloudCredentialAPI.AzureCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAzureCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAzureCloudCommand** | [**CreateAzureCloudCommand**](CreateAzureCloudCommand.md) |  | 

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


## AzureDashboard

> []AzureQuotaListRecordDto AzureDashboard(ctx).AzureDashboardCommand(azureDashboardCommand).Execute()

Fetch Azure quota list

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
	azureDashboardCommand := *openapiclient.NewAzureDashboardCommand(int32(123), "FilterBy_example") // AzureDashboardCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AzureCloudCredentialAPI.AzureDashboard(context.Background()).AzureDashboardCommand(azureDashboardCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AzureCloudCredentialAPI.AzureDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AzureDashboard`: []AzureQuotaListRecordDto
	fmt.Fprintf(os.Stdout, "Response from `AzureCloudCredentialAPI.AzureDashboard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAzureDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **azureDashboardCommand** | [**AzureDashboardCommand**](AzureDashboardCommand.md) |  | 

### Return type

[**[]AzureQuotaListRecordDto**](AzureQuotaListRecordDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AzureFileShareList

> []AzureFileShareDto AzureFileShareList(ctx).AzureFileShareListCommand(azureFileShareListCommand).Execute()

Retrieve Azure File Shares within a resource group

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
	azureFileShareListCommand := *openapiclient.NewAzureFileShareListCommand("AzureClientId_example", "AzureClientSecret_example", "AzureTenantId_example", "AzureSubscriptionId_example", "ResourceGroupName_example", NullableInt32(123)) // AzureFileShareListCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AzureCloudCredentialAPI.AzureFileShareList(context.Background()).AzureFileShareListCommand(azureFileShareListCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AzureCloudCredentialAPI.AzureFileShareList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AzureFileShareList`: []AzureFileShareDto
	fmt.Fprintf(os.Stdout, "Response from `AzureCloudCredentialAPI.AzureFileShareList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAzureFileShareListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **azureFileShareListCommand** | [**AzureFileShareListCommand**](AzureFileShareListCommand.md) |  | 

### Return type

[**[]AzureFileShareDto**](AzureFileShareDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AzureList

> AzureCredentialList AzureList(ctx).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Id(id).DomainId(domainId).Limit(limit).Offset(offset).Execute()

Retrieve list of azure cloud credentials

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
	sortBy := "sortBy_example" // string |  (optional)
	sortDirection := "sortDirection_example" // string |  (optional)
	search := "search_example" // string |  (optional)
	searchId := "searchId_example" // string |  (optional)
	id := int32(56) // int32 |  (optional)
	domainId := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)
	offset := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AzureCloudCredentialAPI.AzureList(context.Background()).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Id(id).DomainId(domainId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AzureCloudCredentialAPI.AzureList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AzureList`: AzureCredentialList
	fmt.Fprintf(os.Stdout, "Response from `AzureCloudCredentialAPI.AzureList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAzureListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **id** | **int32** |  | 
 **domainId** | **int32** |  | 
 **limit** | **int32** |  | [default to 50]
 **offset** | **int32** |  | [default to 0]

### Return type

[**AzureCredentialList**](AzureCredentialList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AzureLocations

> []string AzureLocations(ctx).AzureLocationsCommand(azureLocationsCommand).Execute()

Fetch Azure location list

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
	azureLocationsCommand := *openapiclient.NewAzureLocationsCommand("AzureSubscriptionId_example", "AzureClientId_example", "AzureClientSecret_example", "AzureTenantId_example") // AzureLocationsCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AzureCloudCredentialAPI.AzureLocations(context.Background()).AzureLocationsCommand(azureLocationsCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AzureCloudCredentialAPI.AzureLocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AzureLocations`: []string
	fmt.Fprintf(os.Stdout, "Response from `AzureCloudCredentialAPI.AzureLocations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAzureLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **azureLocationsCommand** | [**AzureLocationsCommand**](AzureLocationsCommand.md) |  | 

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


## AzureOffers

> AzureOffersList AzureOffers(ctx, cloudId, publisher).SortBy(sortBy).SortDirection(sortDirection).Search(search).Limit(limit).Offset(offset).Execute()

List Azure offer list by publisher

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
	publisher := "publisher_example" // string | 
	sortBy := "sortBy_example" // string |  (optional)
	sortDirection := "sortDirection_example" // string |  (optional)
	search := "search_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)
	offset := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AzureCloudCredentialAPI.AzureOffers(context.Background(), cloudId, publisher).SortBy(sortBy).SortDirection(sortDirection).Search(search).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AzureCloudCredentialAPI.AzureOffers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AzureOffers`: AzureOffersList
	fmt.Fprintf(os.Stdout, "Response from `AzureCloudCredentialAPI.AzureOffers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int32** |  | 
**publisher** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAzureOffersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **limit** | **int32** |  | [default to 50]
 **offset** | **int32** |  | [default to 0]

### Return type

[**AzureOffersList**](AzureOffersList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AzurePublishers

> AzurePublishersList AzurePublishers(ctx, cloudId).SortBy(sortBy).SortDirection(sortDirection).Search(search).Limit(limit).Offset(offset).Execute()

List Azure publishers list

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
	sortBy := "sortBy_example" // string |  (optional)
	sortDirection := "sortDirection_example" // string |  (optional)
	search := "search_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)
	offset := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AzureCloudCredentialAPI.AzurePublishers(context.Background(), cloudId).SortBy(sortBy).SortDirection(sortDirection).Search(search).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AzureCloudCredentialAPI.AzurePublishers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AzurePublishers`: AzurePublishersList
	fmt.Fprintf(os.Stdout, "Response from `AzureCloudCredentialAPI.AzurePublishers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAzurePublishersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **limit** | **int32** |  | [default to 50]
 **offset** | **int32** |  | [default to 0]

### Return type

[**AzurePublishersList**](AzurePublishersList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AzureSkus

> AzureSkusList AzureSkus(ctx, cloudId, publisher, offer).SortBy(sortBy).SortDirection(sortDirection).Search(search).Limit(limit).Offset(offset).Execute()

List Azure skus list by publisher and offer

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
	publisher := "publisher_example" // string | 
	offer := "offer_example" // string | 
	sortBy := "sortBy_example" // string |  (optional)
	sortDirection := "sortDirection_example" // string |  (optional)
	search := "search_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)
	offset := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AzureCloudCredentialAPI.AzureSkus(context.Background(), cloudId, publisher, offer).SortBy(sortBy).SortDirection(sortDirection).Search(search).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AzureCloudCredentialAPI.AzureSkus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AzureSkus`: AzureSkusList
	fmt.Fprintf(os.Stdout, "Response from `AzureCloudCredentialAPI.AzureSkus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int32** |  | 
**publisher** | **string** |  | 
**offer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAzureSkusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **limit** | **int32** |  | [default to 50]
 **offset** | **int32** |  | [default to 0]

### Return type

[**AzureSkusList**](AzureSkusList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AzureSubscriptions

> []CommonStringBasedDropdownDto AzureSubscriptions(ctx).AzureSubscriptionListCommand(azureSubscriptionListCommand).Execute()

Azure subscriptions list

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
	azureSubscriptionListCommand := *openapiclient.NewAzureSubscriptionListCommand("AzureClientId_example", "AzureClientSecret_example", "AzureTenantId_example") // AzureSubscriptionListCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AzureCloudCredentialAPI.AzureSubscriptions(context.Background()).AzureSubscriptionListCommand(azureSubscriptionListCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AzureCloudCredentialAPI.AzureSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AzureSubscriptions`: []CommonStringBasedDropdownDto
	fmt.Fprintf(os.Stdout, "Response from `AzureCloudCredentialAPI.AzureSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAzureSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **azureSubscriptionListCommand** | [**AzureSubscriptionListCommand**](AzureSubscriptionListCommand.md) |  | 

### Return type

[**[]CommonStringBasedDropdownDto**](CommonStringBasedDropdownDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AzureUpdate

> AzureUpdate(ctx).UpdateAzureCommand(updateAzureCommand).Execute()

Update Azure credentials

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
	updateAzureCommand := *openapiclient.NewUpdateAzureCommand(int32(123), "Name_example", "AzureClientSecret_example", "AzureClientId_example") // UpdateAzureCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AzureCloudCredentialAPI.AzureUpdate(context.Background()).UpdateAzureCommand(updateAzureCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AzureCloudCredentialAPI.AzureUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAzureUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateAzureCommand** | [**UpdateAzureCommand**](UpdateAzureCommand.md) |  | 

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


## AzureVnetList

> []AzureVnetDto AzureVnetList(ctx).AzureVnetListCommand(azureVnetListCommand).Execute()

Retrieve VNets within an Azure resource group

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
	azureVnetListCommand := *openapiclient.NewAzureVnetListCommand("AzureClientId_example", "AzureClientSecret_example", "AzureTenantId_example", "AzureSubscriptionId_example", "ResourceGroupName_example", NullableInt32(123), openapiclient.IpMode("Ipv4")) // AzureVnetListCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AzureCloudCredentialAPI.AzureVnetList(context.Background()).AzureVnetListCommand(azureVnetListCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AzureCloudCredentialAPI.AzureVnetList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AzureVnetList`: []AzureVnetDto
	fmt.Fprintf(os.Stdout, "Response from `AzureCloudCredentialAPI.AzureVnetList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAzureVnetListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **azureVnetListCommand** | [**AzureVnetListCommand**](AzureVnetListCommand.md) |  | 

### Return type

[**[]AzureVnetDto**](AzureVnetDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AzureVnetResourceGroupList

> []string AzureVnetResourceGroupList(ctx).AzureVnetResourceGroupListCommand(azureVnetResourceGroupListCommand).Execute()

Retrieve Azure resource groups that contain VNets

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
	azureVnetResourceGroupListCommand := *openapiclient.NewAzureVnetResourceGroupListCommand("AzureClientId_example", "AzureClientSecret_example", "AzureTenantId_example", "AzureSubscriptionId_example", NullableInt32(123)) // AzureVnetResourceGroupListCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AzureCloudCredentialAPI.AzureVnetResourceGroupList(context.Background()).AzureVnetResourceGroupListCommand(azureVnetResourceGroupListCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AzureCloudCredentialAPI.AzureVnetResourceGroupList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AzureVnetResourceGroupList`: []string
	fmt.Fprintf(os.Stdout, "Response from `AzureCloudCredentialAPI.AzureVnetResourceGroupList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAzureVnetResourceGroupListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **azureVnetResourceGroupListCommand** | [**AzureVnetResourceGroupListCommand**](AzureVnetResourceGroupListCommand.md) |  | 

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


## AzureVnetSubnetList

> []AzureSubnetDto AzureVnetSubnetList(ctx).AzureVnetSubnetListCommand(azureVnetSubnetListCommand).Execute()

Retrieve subnets within an Azure VNet

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
	azureVnetSubnetListCommand := *openapiclient.NewAzureVnetSubnetListCommand("AzureClientId_example", "AzureClientSecret_example", "AzureTenantId_example", "AzureSubscriptionId_example", "ResourceGroupName_example", "VnetName_example", NullableInt32(123), openapiclient.IpMode("Ipv4")) // AzureVnetSubnetListCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AzureCloudCredentialAPI.AzureVnetSubnetList(context.Background()).AzureVnetSubnetListCommand(azureVnetSubnetListCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AzureCloudCredentialAPI.AzureVnetSubnetList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AzureVnetSubnetList`: []AzureSubnetDto
	fmt.Fprintf(os.Stdout, "Response from `AzureCloudCredentialAPI.AzureVnetSubnetList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAzureVnetSubnetListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **azureVnetSubnetListCommand** | [**AzureVnetSubnetListCommand**](AzureVnetSubnetListCommand.md) |  | 

### Return type

[**[]AzureSubnetDto**](AzureSubnetDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AzureZones

> AzResult AzureZones(ctx).AzureZonesCommand(azureZonesCommand).Execute()

Fetch Azure zone list

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
	azureZonesCommand := *openapiclient.NewAzureZonesCommand("AzureSubscriptionId_example", "AzureClientId_example", "AzureClientSecret_example", "AzureTenantId_example", "AzureLocation_example", NullableInt32(123)) // AzureZonesCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AzureCloudCredentialAPI.AzureZones(context.Background()).AzureZonesCommand(azureZonesCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AzureCloudCredentialAPI.AzureZones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AzureZones`: AzResult
	fmt.Fprintf(os.Stdout, "Response from `AzureCloudCredentialAPI.AzureZones`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAzureZonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **azureZonesCommand** | [**AzureZonesCommand**](AzureZonesCommand.md) |  | 

### Return type

[**AzResult**](AzResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

