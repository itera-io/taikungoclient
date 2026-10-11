# \GoogleAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GooglecloudBillingAccountList**](GoogleAPI.md#GooglecloudBillingAccountList) | **Post** /api/v1/googlecloud/billing-accounts | Retrieve google billing accounts list
[**GooglecloudCreate**](GoogleAPI.md#GooglecloudCreate) | **Post** /api/v1/googlecloud/create | Create google cloud credential
[**GooglecloudFilestoreList**](GoogleAPI.md#GooglecloudFilestoreList) | **Post** /api/v1/googlecloud/filestore-list | Retrieve list of GCP Filestore instances for a given project and region
[**GooglecloudGkeClusters**](GoogleAPI.md#GooglecloudGkeClusters) | **Post** /api/v1/googlecloud/gke-clusters | List of gke clusters
[**GooglecloudList**](GoogleAPI.md#GooglecloudList) | **Get** /api/v1/googlecloud/list | Retrieve list of google cloud credentials
[**GooglecloudRegionList**](GoogleAPI.md#GooglecloudRegionList) | **Post** /api/v1/googlecloud/regions | Retrieve google region list
[**GooglecloudSubnetworkList**](GoogleAPI.md#GooglecloudSubnetworkList) | **Post** /api/v1/googlecloud/subnetwork-list | Retrieve list of GCP subnetworks for a given project and region
[**GooglecloudVpcList**](GoogleAPI.md#GooglecloudVpcList) | **Post** /api/v1/googlecloud/vpc-list | Retrieve list of GCP VPC networks
[**GooglecloudZoneList**](GoogleAPI.md#GooglecloudZoneList) | **Post** /api/v1/googlecloud/zones | Google zones list



## GooglecloudBillingAccountList

> []CommonStringBasedDropdownDto GooglecloudBillingAccountList(ctx).Config(config).Execute()

Retrieve google billing accounts list

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
	config := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GoogleAPI.GooglecloudBillingAccountList(context.Background()).Config(config).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GoogleAPI.GooglecloudBillingAccountList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GooglecloudBillingAccountList`: []CommonStringBasedDropdownDto
	fmt.Fprintf(os.Stdout, "Response from `GoogleAPI.GooglecloudBillingAccountList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGooglecloudBillingAccountListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **config** | ***os.File** |  | 

### Return type

[**[]CommonStringBasedDropdownDto**](CommonStringBasedDropdownDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GooglecloudCreate

> ApiResponse GooglecloudCreate(ctx).Config(config).Name(name).ImportProject(importProject).FolderId(folderId).BillingAccountId(billingAccountId).AzCount(azCount).Region(region).OrganizationId(organizationId).IpMode(ipMode).VpcMode(vpcMode).ImportedGcpProjectId(importedGcpProjectId).ImportedVpcName(importedVpcName).Subnets(subnets).ImportedFilestoreId(importedFilestoreId).ImportedFilestoreProjectId(importedFilestoreProjectId).Execute()

Create google cloud credential

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
	config := os.NewFile(1234, "some_file") // *os.File |  (optional)
	name := "name_example" // string |  (optional)
	importProject := true // bool |  (optional)
	folderId := "folderId_example" // string |  (optional)
	billingAccountId := "billingAccountId_example" // string |  (optional)
	azCount := int32(56) // int32 |  (optional)
	region := "region_example" // string |  (optional)
	organizationId := int32(56) // int32 |  (optional)
	ipMode := openapiclient.IpMode("Ipv4") // IpMode |  (optional)
	vpcMode := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} |  (optional)
	importedGcpProjectId := "importedGcpProjectId_example" // string |  (optional)
	importedVpcName := "importedVpcName_example" // string |  (optional)
	subnets := []string{"Inner_example"} // []string |  (optional)
	importedFilestoreId := "importedFilestoreId_example" // string |  (optional)
	importedFilestoreProjectId := "importedFilestoreProjectId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GoogleAPI.GooglecloudCreate(context.Background()).Config(config).Name(name).ImportProject(importProject).FolderId(folderId).BillingAccountId(billingAccountId).AzCount(azCount).Region(region).OrganizationId(organizationId).IpMode(ipMode).VpcMode(vpcMode).ImportedGcpProjectId(importedGcpProjectId).ImportedVpcName(importedVpcName).Subnets(subnets).ImportedFilestoreId(importedFilestoreId).ImportedFilestoreProjectId(importedFilestoreProjectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GoogleAPI.GooglecloudCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GooglecloudCreate`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `GoogleAPI.GooglecloudCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGooglecloudCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **config** | ***os.File** |  | 
 **name** | **string** |  | 
 **importProject** | **bool** |  | 
 **folderId** | **string** |  | 
 **billingAccountId** | **string** |  | 
 **azCount** | **int32** |  | 
 **region** | **string** |  | 
 **organizationId** | **int32** |  | 
 **ipMode** | [**IpMode**](IpMode.md) |  | 
 **vpcMode** | **map[string]interface{}** |  | 
 **importedGcpProjectId** | **string** |  | 
 **importedVpcName** | **string** |  | 
 **subnets** | **[]string** |  | 
 **importedFilestoreId** | **string** |  | 
 **importedFilestoreProjectId** | **string** |  | 

### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GooglecloudFilestoreList

> []GcpFilestoreDto GooglecloudFilestoreList(ctx).Config(config).ImportedGcpProjectId(importedGcpProjectId).ImportedFilestoreProjectId(importedFilestoreProjectId).Region(region).CloudId(cloudId).Execute()

Retrieve list of GCP Filestore instances for a given project and region

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
	config := os.NewFile(1234, "some_file") // *os.File |  (optional)
	importedGcpProjectId := "importedGcpProjectId_example" // string |  (optional)
	importedFilestoreProjectId := "importedFilestoreProjectId_example" // string |  (optional)
	region := "region_example" // string |  (optional)
	cloudId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GoogleAPI.GooglecloudFilestoreList(context.Background()).Config(config).ImportedGcpProjectId(importedGcpProjectId).ImportedFilestoreProjectId(importedFilestoreProjectId).Region(region).CloudId(cloudId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GoogleAPI.GooglecloudFilestoreList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GooglecloudFilestoreList`: []GcpFilestoreDto
	fmt.Fprintf(os.Stdout, "Response from `GoogleAPI.GooglecloudFilestoreList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGooglecloudFilestoreListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **config** | ***os.File** |  | 
 **importedGcpProjectId** | **string** |  | 
 **importedFilestoreProjectId** | **string** |  | 
 **region** | **string** |  | 
 **cloudId** | **int32** |  | 

### Return type

[**[]GcpFilestoreDto**](GcpFilestoreDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GooglecloudGkeClusters

> []GkeClusterDto GooglecloudGkeClusters(ctx).GkeClustersListCommand(gkeClustersListCommand).Execute()

List of gke clusters

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
	gkeClustersListCommand := *openapiclient.NewGkeClustersListCommand(int32(123)) // GkeClustersListCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GoogleAPI.GooglecloudGkeClusters(context.Background()).GkeClustersListCommand(gkeClustersListCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GoogleAPI.GooglecloudGkeClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GooglecloudGkeClusters`: []GkeClusterDto
	fmt.Fprintf(os.Stdout, "Response from `GoogleAPI.GooglecloudGkeClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGooglecloudGkeClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gkeClustersListCommand** | [**GkeClustersListCommand**](GkeClustersListCommand.md) |  | 

### Return type

[**[]GkeClusterDto**](GkeClusterDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GooglecloudList

> GoogleCredentialList GooglecloudList(ctx).Limit(limit).Offset(offset).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).DomainId(domainId).Id(id).Execute()

Retrieve list of google cloud credentials

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
	organizationId := int32(56) // int32 |  (optional)
	sortBy := "sortBy_example" // string |  (optional)
	sortDirection := "sortDirection_example" // string |  (optional)
	search := "search_example" // string |  (optional)
	searchId := "searchId_example" // string |  (optional)
	domainId := int32(56) // int32 |  (optional)
	id := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GoogleAPI.GooglecloudList(context.Background()).Limit(limit).Offset(offset).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).DomainId(domainId).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GoogleAPI.GooglecloudList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GooglecloudList`: GoogleCredentialList
	fmt.Fprintf(os.Stdout, "Response from `GoogleAPI.GooglecloudList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGooglecloudListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **organizationId** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **domainId** | **int32** |  | 
 **id** | **int32** |  | 

### Return type

[**GoogleCredentialList**](GoogleCredentialList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GooglecloudRegionList

> []string GooglecloudRegionList(ctx).Config(config).Execute()

Retrieve google region list

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
	config := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GoogleAPI.GooglecloudRegionList(context.Background()).Config(config).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GoogleAPI.GooglecloudRegionList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GooglecloudRegionList`: []string
	fmt.Fprintf(os.Stdout, "Response from `GoogleAPI.GooglecloudRegionList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGooglecloudRegionListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **config** | ***os.File** |  | 

### Return type

**[]string**

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GooglecloudSubnetworkList

> []GcpSubnetworkDto GooglecloudSubnetworkList(ctx).Config(config).ImportedGcpProjectId(importedGcpProjectId).Region(region).NetworkName(networkName).CloudId(cloudId).IpMode(ipMode).Execute()

Retrieve list of GCP subnetworks for a given project and region

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
	config := os.NewFile(1234, "some_file") // *os.File |  (optional)
	importedGcpProjectId := "importedGcpProjectId_example" // string |  (optional)
	region := "region_example" // string |  (optional)
	networkName := "networkName_example" // string |  (optional)
	cloudId := int32(56) // int32 |  (optional)
	ipMode := openapiclient.IpMode("Ipv4") // IpMode |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GoogleAPI.GooglecloudSubnetworkList(context.Background()).Config(config).ImportedGcpProjectId(importedGcpProjectId).Region(region).NetworkName(networkName).CloudId(cloudId).IpMode(ipMode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GoogleAPI.GooglecloudSubnetworkList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GooglecloudSubnetworkList`: []GcpSubnetworkDto
	fmt.Fprintf(os.Stdout, "Response from `GoogleAPI.GooglecloudSubnetworkList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGooglecloudSubnetworkListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **config** | ***os.File** |  | 
 **importedGcpProjectId** | **string** |  | 
 **region** | **string** |  | 
 **networkName** | **string** |  | 
 **cloudId** | **int32** |  | 
 **ipMode** | [**IpMode**](IpMode.md) |  | 

### Return type

[**[]GcpSubnetworkDto**](GcpSubnetworkDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GooglecloudVpcList

> []GcpVpcDto GooglecloudVpcList(ctx).Config(config).ImportedGcpProjectId(importedGcpProjectId).CloudId(cloudId).IpMode(ipMode).Execute()

Retrieve list of GCP VPC networks

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
	config := os.NewFile(1234, "some_file") // *os.File |  (optional)
	importedGcpProjectId := "importedGcpProjectId_example" // string |  (optional)
	cloudId := int32(56) // int32 |  (optional)
	ipMode := openapiclient.IpMode("Ipv4") // IpMode |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GoogleAPI.GooglecloudVpcList(context.Background()).Config(config).ImportedGcpProjectId(importedGcpProjectId).CloudId(cloudId).IpMode(ipMode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GoogleAPI.GooglecloudVpcList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GooglecloudVpcList`: []GcpVpcDto
	fmt.Fprintf(os.Stdout, "Response from `GoogleAPI.GooglecloudVpcList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGooglecloudVpcListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **config** | ***os.File** |  | 
 **importedGcpProjectId** | **string** |  | 
 **cloudId** | **int32** |  | 
 **ipMode** | [**IpMode**](IpMode.md) |  | 

### Return type

[**[]GcpVpcDto**](GcpVpcDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GooglecloudZoneList

> AzResult GooglecloudZoneList(ctx).Config(config).Region(region).CloudId(cloudId).Execute()

Google zones list

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
	config := os.NewFile(1234, "some_file") // *os.File |  (optional)
	region := "region_example" // string |  (optional)
	cloudId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GoogleAPI.GooglecloudZoneList(context.Background()).Config(config).Region(region).CloudId(cloudId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GoogleAPI.GooglecloudZoneList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GooglecloudZoneList`: AzResult
	fmt.Fprintf(os.Stdout, "Response from `GoogleAPI.GooglecloudZoneList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGooglecloudZoneListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **config** | ***os.File** |  | 
 **region** | **string** |  | 
 **cloudId** | **int32** |  | 

### Return type

[**AzResult**](AzResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

