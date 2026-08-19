# \OpenstackCloudCredentialAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OpenstackCreate**](OpenstackCloudCredentialAPI.md#OpenstackCreate) | **Post** /api/v1/openstack/create | Add Openstack credentials
[**OpenstackList**](OpenstackCloudCredentialAPI.md#OpenstackList) | **Get** /api/v1/openstack/list | Retrieve list of openstack cloud credentials
[**OpenstackNetworks**](OpenstackCloudCredentialAPI.md#OpenstackNetworks) | **Post** /api/v1/openstack/networks | Openstack network list
[**OpenstackProjects**](OpenstackCloudCredentialAPI.md#OpenstackProjects) | **Post** /api/v1/openstack/projects | Openstack project list
[**OpenstackQuotas**](OpenstackCloudCredentialAPI.md#OpenstackQuotas) | **Post** /api/v1/openstack/quotas | Openstack quota list
[**OpenstackRegionList**](OpenstackCloudCredentialAPI.md#OpenstackRegionList) | **Post** /api/v1/openstack/regions | Retrieve Openstack regions
[**OpenstackStorageTypes**](OpenstackCloudCredentialAPI.md#OpenstackStorageTypes) | **Post** /api/v1/openstack/storage-types | Openstack storage type list
[**OpenstackSubnet**](OpenstackCloudCredentialAPI.md#OpenstackSubnet) | **Post** /api/v1/openstack/subnets | Retrieve Openstack subnets
[**OpenstackUpdate**](OpenstackCloudCredentialAPI.md#OpenstackUpdate) | **Post** /api/v1/openstack/update | Update Openstack credentials
[**OpenstackVolumes**](OpenstackCloudCredentialAPI.md#OpenstackVolumes) | **Post** /api/v1/openstack/volumes | Openstack volume list
[**OpenstackZones**](OpenstackCloudCredentialAPI.md#OpenstackZones) | **Post** /api/v1/openstack/zones | Fetch Openstack zones



## OpenstackCreate

> ApiResponse OpenstackCreate(ctx).CreateOpenstackCloudCommand(createOpenstackCloudCommand).Execute()

Add Openstack credentials

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
	createOpenstackCloudCommand := *openapiclient.NewCreateOpenstackCloudCommand() // CreateOpenstackCloudCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpenstackCloudCredentialAPI.OpenstackCreate(context.Background()).CreateOpenstackCloudCommand(createOpenstackCloudCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenstackCloudCredentialAPI.OpenstackCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpenstackCreate`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `OpenstackCloudCredentialAPI.OpenstackCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpenstackCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createOpenstackCloudCommand** | [**CreateOpenstackCloudCommand**](CreateOpenstackCloudCommand.md) |  | 

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


## OpenstackList

> OpenstackCredentialList OpenstackList(ctx).Limit(limit).Offset(offset).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Id(id).IsInfra(isInfra).DomainId(domainId).Execute()

Retrieve list of openstack cloud credentials

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
	id := int32(56) // int32 |  (optional)
	isInfra := true // bool |  (optional)
	domainId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpenstackCloudCredentialAPI.OpenstackList(context.Background()).Limit(limit).Offset(offset).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Id(id).IsInfra(isInfra).DomainId(domainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenstackCloudCredentialAPI.OpenstackList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpenstackList`: OpenstackCredentialList
	fmt.Fprintf(os.Stdout, "Response from `OpenstackCloudCredentialAPI.OpenstackList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpenstackListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **organizationId** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **id** | **int32** |  | 
 **isInfra** | **bool** |  | 
 **domainId** | **int32** |  | 

### Return type

[**OpenstackCredentialList**](OpenstackCredentialList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpenstackNetworks

> []CommonStringBasedDropdownDto OpenstackNetworks(ctx).OpenStackNetworkListQuery(openStackNetworkListQuery).Execute()

Openstack network list

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
	openStackNetworkListQuery := *openapiclient.NewOpenStackNetworkListQuery("OpenStackUser_example", "OpenStackPassword_example", "OpenStackUrl_example", "OpenStackProjectId_example", "OpenStackDomain_example", "OpenStackRegion_example", false, false) // OpenStackNetworkListQuery |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpenstackCloudCredentialAPI.OpenstackNetworks(context.Background()).OpenStackNetworkListQuery(openStackNetworkListQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenstackCloudCredentialAPI.OpenstackNetworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpenstackNetworks`: []CommonStringBasedDropdownDto
	fmt.Fprintf(os.Stdout, "Response from `OpenstackCloudCredentialAPI.OpenstackNetworks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpenstackNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **openStackNetworkListQuery** | [**OpenStackNetworkListQuery**](OpenStackNetworkListQuery.md) |  | 

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


## OpenstackProjects

> []CommonStringBasedDropdownDto OpenstackProjects(ctx).OpenStackProjectListQuery(openStackProjectListQuery).Execute()

Openstack project list

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
	openStackProjectListQuery := *openapiclient.NewOpenStackProjectListQuery("OpenStackUser_example", "OpenStackPassword_example", "OpenStackUrl_example", "OpenStackDomain_example", false, false) // OpenStackProjectListQuery |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpenstackCloudCredentialAPI.OpenstackProjects(context.Background()).OpenStackProjectListQuery(openStackProjectListQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenstackCloudCredentialAPI.OpenstackProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpenstackProjects`: []CommonStringBasedDropdownDto
	fmt.Fprintf(os.Stdout, "Response from `OpenstackCloudCredentialAPI.OpenstackProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpenstackProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **openStackProjectListQuery** | [**OpenStackProjectListQuery**](OpenStackProjectListQuery.md) |  | 

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


## OpenstackQuotas

> OpenstackQuotaList OpenstackQuotas(ctx).OpenstackQuotasCommand(openstackQuotasCommand).Execute()

Openstack quota list

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
	openstackQuotasCommand := *openapiclient.NewOpenstackQuotasCommand(int32(123)) // OpenstackQuotasCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpenstackCloudCredentialAPI.OpenstackQuotas(context.Background()).OpenstackQuotasCommand(openstackQuotasCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenstackCloudCredentialAPI.OpenstackQuotas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpenstackQuotas`: OpenstackQuotaList
	fmt.Fprintf(os.Stdout, "Response from `OpenstackCloudCredentialAPI.OpenstackQuotas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpenstackQuotasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **openstackQuotasCommand** | [**OpenstackQuotasCommand**](OpenstackQuotasCommand.md) |  | 

### Return type

[**OpenstackQuotaList**](OpenstackQuotaList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpenstackRegionList

> []string OpenstackRegionList(ctx).OpenStackRegionListQuery(openStackRegionListQuery).Execute()

Retrieve Openstack regions

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
	openStackRegionListQuery := *openapiclient.NewOpenStackRegionListQuery("OpenStackUser_example", "OpenStackPassword_example", "OpenStackUrl_example", "OpenStackDomain_example", false, false) // OpenStackRegionListQuery |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpenstackCloudCredentialAPI.OpenstackRegionList(context.Background()).OpenStackRegionListQuery(openStackRegionListQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenstackCloudCredentialAPI.OpenstackRegionList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpenstackRegionList`: []string
	fmt.Fprintf(os.Stdout, "Response from `OpenstackCloudCredentialAPI.OpenstackRegionList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpenstackRegionListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **openStackRegionListQuery** | [**OpenStackRegionListQuery**](OpenStackRegionListQuery.md) |  | 

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


## OpenstackStorageTypes

> []string OpenstackStorageTypes(ctx).OpenstackStorageTypeListQuery(openstackStorageTypeListQuery).Execute()

Openstack storage type list

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
	openstackStorageTypeListQuery := *openapiclient.NewOpenstackStorageTypeListQuery("OpenStackUser_example", "OpenStackPassword_example", "OpenStackUrl_example", "OpenStackDomain_example", "OpenStackRegion_example", false, false, "OpenstackProject_example") // OpenstackStorageTypeListQuery |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpenstackCloudCredentialAPI.OpenstackStorageTypes(context.Background()).OpenstackStorageTypeListQuery(openstackStorageTypeListQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenstackCloudCredentialAPI.OpenstackStorageTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpenstackStorageTypes`: []string
	fmt.Fprintf(os.Stdout, "Response from `OpenstackCloudCredentialAPI.OpenstackStorageTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpenstackStorageTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **openstackStorageTypeListQuery** | [**OpenstackStorageTypeListQuery**](OpenstackStorageTypeListQuery.md) |  | 

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


## OpenstackSubnet

> []Subnet OpenstackSubnet(ctx).OpenstackSubnetListQuery(openstackSubnetListQuery).Execute()

Retrieve Openstack subnets

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
	openstackSubnetListQuery := *openapiclient.NewOpenstackSubnetListQuery("OpenStackUser_example", "OpenStackPassword_example", "OpenStackUrl_example", "OpenStackProject_example", "OpenStackProjectId_example", "OpenStackDomain_example", "OpenStackRegion_example", false) // OpenstackSubnetListQuery |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpenstackCloudCredentialAPI.OpenstackSubnet(context.Background()).OpenstackSubnetListQuery(openstackSubnetListQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenstackCloudCredentialAPI.OpenstackSubnet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpenstackSubnet`: []Subnet
	fmt.Fprintf(os.Stdout, "Response from `OpenstackCloudCredentialAPI.OpenstackSubnet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpenstackSubnetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **openstackSubnetListQuery** | [**OpenstackSubnetListQuery**](OpenstackSubnetListQuery.md) |  | 

### Return type

[**[]Subnet**](Subnet.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpenstackUpdate

> OpenstackUpdate(ctx).UpdateOpenStackCommand(updateOpenStackCommand).Execute()

Update Openstack credentials

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
	updateOpenStackCommand := *openapiclient.NewUpdateOpenStackCommand(int32(123), "Name_example", "OpenStackUser_example", "OpenStackPassword_example") // UpdateOpenStackCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OpenstackCloudCredentialAPI.OpenstackUpdate(context.Background()).UpdateOpenStackCommand(updateOpenStackCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenstackCloudCredentialAPI.OpenstackUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpenstackUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateOpenStackCommand** | [**UpdateOpenStackCommand**](UpdateOpenStackCommand.md) |  | 

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


## OpenstackVolumes

> []string OpenstackVolumes(ctx).OpenstackVolumeTypeListQuery(openstackVolumeTypeListQuery).Execute()

Openstack volume list

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
	openstackVolumeTypeListQuery := *openapiclient.NewOpenstackVolumeTypeListQuery(NullableInt32(123), "OpenStackUser_example", "OpenStackPassword_example", "OpenStackUrl_example", "OpenStackDomain_example", "OpenStackRegion_example", false, false, "OpenstackProject_example") // OpenstackVolumeTypeListQuery |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpenstackCloudCredentialAPI.OpenstackVolumes(context.Background()).OpenstackVolumeTypeListQuery(openstackVolumeTypeListQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenstackCloudCredentialAPI.OpenstackVolumes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpenstackVolumes`: []string
	fmt.Fprintf(os.Stdout, "Response from `OpenstackCloudCredentialAPI.OpenstackVolumes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpenstackVolumesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **openstackVolumeTypeListQuery** | [**OpenstackVolumeTypeListQuery**](OpenstackVolumeTypeListQuery.md) |  | 

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


## OpenstackZones

> []string OpenstackZones(ctx).OpenStackZoneListQuery(openStackZoneListQuery).Execute()

Fetch Openstack zones

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
	openStackZoneListQuery := *openapiclient.NewOpenStackZoneListQuery("OpenStackUser_example", "OpenStackPassword_example", "OpenStackUrl_example", "OpenStackDomain_example", "OpenStackRegion_example", false, false, "OpenstackProject_example") // OpenStackZoneListQuery |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpenstackCloudCredentialAPI.OpenstackZones(context.Background()).OpenStackZoneListQuery(openStackZoneListQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenstackCloudCredentialAPI.OpenstackZones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpenstackZones`: []string
	fmt.Fprintf(os.Stdout, "Response from `OpenstackCloudCredentialAPI.OpenstackZones`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpenstackZonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **openStackZoneListQuery** | [**OpenStackZoneListQuery**](OpenStackZoneListQuery.md) |  | 

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

