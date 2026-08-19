# \AWSCloudCredentialAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AwsCreate**](AWSCloudCredentialAPI.md#AwsCreate) | **Post** /api/v1/aws/create | Add Aws credentials
[**AwsEfsFileSystemByVpcList**](AWSCloudCredentialAPI.md#AwsEfsFileSystemByVpcList) | **Post** /api/v1/aws/efs-file-system-by-vpc-list | Retrieve aws efs file system by vpc list
[**AwsEksClusters**](AWSCloudCredentialAPI.md#AwsEksClusters) | **Get** /api/v1/aws/eks-clusters/{cloudId} | Retrieve eks clusters list
[**AwsEksNodeGroups**](AWSCloudCredentialAPI.md#AwsEksNodeGroups) | **Get** /api/v1/aws/eks-node-groups/{projectId} | Retrieve eks node group list
[**AwsGetEksAuthInfo**](AWSCloudCredentialAPI.md#AwsGetEksAuthInfo) | **Post** /api/v1/aws/eks-auth-info | Get EKS auth info
[**AwsList**](AWSCloudCredentialAPI.md#AwsList) | **Get** /api/v1/aws/list | Retrieve list of aws cloud credentials
[**AwsOwners**](AWSCloudCredentialAPI.md#AwsOwners) | **Get** /api/v1/aws/owners | Retrieve aws verified owner list
[**AwsRegionlist**](AWSCloudCredentialAPI.md#AwsRegionlist) | **Post** /api/v1/aws/regions | Retrieve aws regions list
[**AwsSubnetList**](AWSCloudCredentialAPI.md#AwsSubnetList) | **Post** /api/v1/aws/subnet-list | Retrieve aws subnet list
[**AwsUpdate**](AWSCloudCredentialAPI.md#AwsUpdate) | **Post** /api/v1/aws/update | Update AWS credentials
[**AwsValidateOwners**](AWSCloudCredentialAPI.md#AwsValidateOwners) | **Post** /api/v1/aws/validate-owners | Aws validate owners
[**AwsVpcList**](AWSCloudCredentialAPI.md#AwsVpcList) | **Post** /api/v1/aws/vpc-list | Retrieve aws vpc list
[**AwsZones**](AWSCloudCredentialAPI.md#AwsZones) | **Post** /api/v1/aws/zones | Fetch Aws zones



## AwsCreate

> ApiResponse AwsCreate(ctx).CreateAwsCloudCommand(createAwsCloudCommand).Execute()

Add Aws credentials

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
	createAwsCloudCommand := *openapiclient.NewCreateAwsCloudCommand("Name_example", "AwsRegion_example") // CreateAwsCloudCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AWSCloudCredentialAPI.AwsCreate(context.Background()).CreateAwsCloudCommand(createAwsCloudCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AWSCloudCredentialAPI.AwsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AwsCreate`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AWSCloudCredentialAPI.AwsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAwsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAwsCloudCommand** | [**CreateAwsCloudCommand**](CreateAwsCloudCommand.md) |  | 

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


## AwsEfsFileSystemByVpcList

> []EfsFileSystemDto AwsEfsFileSystemByVpcList(ctx).AwsFileSystemsCommand(awsFileSystemsCommand).Execute()

Retrieve aws efs file system by vpc list

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
	awsFileSystemsCommand := *openapiclient.NewAwsFileSystemsCommand("AwsAccessKeyId_example", "AwsSecretAccessKey_example", "RoleArn_example", "ExternalId_example", "AwsRegion_example", "VpcId_example") // AwsFileSystemsCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AWSCloudCredentialAPI.AwsEfsFileSystemByVpcList(context.Background()).AwsFileSystemsCommand(awsFileSystemsCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AWSCloudCredentialAPI.AwsEfsFileSystemByVpcList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AwsEfsFileSystemByVpcList`: []EfsFileSystemDto
	fmt.Fprintf(os.Stdout, "Response from `AWSCloudCredentialAPI.AwsEfsFileSystemByVpcList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAwsEfsFileSystemByVpcListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsFileSystemsCommand** | [**AwsFileSystemsCommand**](AwsFileSystemsCommand.md) |  | 

### Return type

[**[]EfsFileSystemDto**](EfsFileSystemDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AwsEksClusters

> []string AwsEksClusters(ctx, cloudId).Region(region).Execute()

Retrieve eks clusters list

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
	region := "region_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AWSCloudCredentialAPI.AwsEksClusters(context.Background(), cloudId).Region(region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AWSCloudCredentialAPI.AwsEksClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AwsEksClusters`: []string
	fmt.Fprintf(os.Stdout, "Response from `AWSCloudCredentialAPI.AwsEksClusters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAwsEksClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **region** | **string** |  | 

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


## AwsEksNodeGroups

> []AwsEksNodeGroupDto AwsEksNodeGroups(ctx, projectId).Execute()

Retrieve eks node group list

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AWSCloudCredentialAPI.AwsEksNodeGroups(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AWSCloudCredentialAPI.AwsEksNodeGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AwsEksNodeGroups`: []AwsEksNodeGroupDto
	fmt.Fprintf(os.Stdout, "Response from `AWSCloudCredentialAPI.AwsEksNodeGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAwsEksNodeGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AwsEksNodeGroupDto**](AwsEksNodeGroupDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AwsGetEksAuthInfo

> AwsEksClusterDto AwsGetEksAuthInfo(ctx).GetEksClusterAuthCommand(getEksClusterAuthCommand).Execute()

Get EKS auth info

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
	getEksClusterAuthCommand := *openapiclient.NewGetEksClusterAuthCommand(int32(123), "KubeConfigName_example", "ConfigToken_example") // GetEksClusterAuthCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AWSCloudCredentialAPI.AwsGetEksAuthInfo(context.Background()).GetEksClusterAuthCommand(getEksClusterAuthCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AWSCloudCredentialAPI.AwsGetEksAuthInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AwsGetEksAuthInfo`: AwsEksClusterDto
	fmt.Fprintf(os.Stdout, "Response from `AWSCloudCredentialAPI.AwsGetEksAuthInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAwsGetEksAuthInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getEksClusterAuthCommand** | [**GetEksClusterAuthCommand**](GetEksClusterAuthCommand.md) |  | 

### Return type

[**AwsEksClusterDto**](AwsEksClusterDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AwsList

> AwsCredentialList AwsList(ctx).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Id(id).DomainId(domainId).Limit(limit).Offset(offset).Execute()

Retrieve list of aws cloud credentials

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
	resp, r, err := apiClient.AWSCloudCredentialAPI.AwsList(context.Background()).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Id(id).DomainId(domainId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AWSCloudCredentialAPI.AwsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AwsList`: AwsCredentialList
	fmt.Fprintf(os.Stdout, "Response from `AWSCloudCredentialAPI.AwsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAwsListRequest struct via the builder pattern


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

[**AwsCredentialList**](AwsCredentialList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AwsOwners

> []CommonStringBasedDropdownDto AwsOwners(ctx).Execute()

Retrieve aws verified owner list

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
	resp, r, err := apiClient.AWSCloudCredentialAPI.AwsOwners(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AWSCloudCredentialAPI.AwsOwners``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AwsOwners`: []CommonStringBasedDropdownDto
	fmt.Fprintf(os.Stdout, "Response from `AWSCloudCredentialAPI.AwsOwners`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAwsOwnersRequest struct via the builder pattern


### Return type

[**[]CommonStringBasedDropdownDto**](CommonStringBasedDropdownDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AwsRegionlist

> []AwsRegionDto AwsRegionlist(ctx).RegionListCommand(regionListCommand).Execute()

Retrieve aws regions list

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
	regionListCommand := *openapiclient.NewRegionListCommand("AwsAccessKeyId_example", "AwsSecretAccessKey_example", "RoleArn_example", "ExternalId_example", NullableInt32(123)) // RegionListCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AWSCloudCredentialAPI.AwsRegionlist(context.Background()).RegionListCommand(regionListCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AWSCloudCredentialAPI.AwsRegionlist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AwsRegionlist`: []AwsRegionDto
	fmt.Fprintf(os.Stdout, "Response from `AWSCloudCredentialAPI.AwsRegionlist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAwsRegionlistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **regionListCommand** | [**RegionListCommand**](RegionListCommand.md) |  | 

### Return type

[**[]AwsRegionDto**](AwsRegionDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AwsSubnetList

> []AwsSubnetDto AwsSubnetList(ctx).AwsSubnetListCommand(awsSubnetListCommand).Execute()

Retrieve aws subnet list

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
	awsSubnetListCommand := *openapiclient.NewAwsSubnetListCommand("AwsAccessKeyId_example", "AwsSecretAccessKey_example", "RoleArn_example", "ExternalId_example", "AwsRegion_example", "VpcId_example") // AwsSubnetListCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AWSCloudCredentialAPI.AwsSubnetList(context.Background()).AwsSubnetListCommand(awsSubnetListCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AWSCloudCredentialAPI.AwsSubnetList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AwsSubnetList`: []AwsSubnetDto
	fmt.Fprintf(os.Stdout, "Response from `AWSCloudCredentialAPI.AwsSubnetList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAwsSubnetListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsSubnetListCommand** | [**AwsSubnetListCommand**](AwsSubnetListCommand.md) |  | 

### Return type

[**[]AwsSubnetDto**](AwsSubnetDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AwsUpdate

> AwsUpdate(ctx).UpdateAwsCommand(updateAwsCommand).Execute()

Update AWS credentials

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
	updateAwsCommand := *openapiclient.NewUpdateAwsCommand(int32(123), "Name_example", "AwsSecretAccessKey_example", "AwsAccessKeyId_example", "RoleArn_example", "ExternalId_example") // UpdateAwsCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AWSCloudCredentialAPI.AwsUpdate(context.Background()).UpdateAwsCommand(updateAwsCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AWSCloudCredentialAPI.AwsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAwsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateAwsCommand** | [**UpdateAwsCommand**](UpdateAwsCommand.md) |  | 

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


## AwsValidateOwners

> AwsValidateOwners(ctx).AwsValidateOwnerCommand(awsValidateOwnerCommand).Execute()

Aws validate owners

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
	awsValidateOwnerCommand := *openapiclient.NewAwsValidateOwnerCommand(int32(123), []string{"Owners_example"}) // AwsValidateOwnerCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AWSCloudCredentialAPI.AwsValidateOwners(context.Background()).AwsValidateOwnerCommand(awsValidateOwnerCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AWSCloudCredentialAPI.AwsValidateOwners``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAwsValidateOwnersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsValidateOwnerCommand** | [**AwsValidateOwnerCommand**](AwsValidateOwnerCommand.md) |  | 

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


## AwsVpcList

> []AwsVpcDto AwsVpcList(ctx).AwsVpcListCommand(awsVpcListCommand).Execute()

Retrieve aws vpc list

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
	awsVpcListCommand := *openapiclient.NewAwsVpcListCommand("AwsAccessKeyId_example", "AwsSecretAccessKey_example", "RoleArn_example", "ExternalId_example", "AwsRegion_example", openapiclient.IpMode("Ipv4")) // AwsVpcListCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AWSCloudCredentialAPI.AwsVpcList(context.Background()).AwsVpcListCommand(awsVpcListCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AWSCloudCredentialAPI.AwsVpcList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AwsVpcList`: []AwsVpcDto
	fmt.Fprintf(os.Stdout, "Response from `AWSCloudCredentialAPI.AwsVpcList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAwsVpcListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsVpcListCommand** | [**AwsVpcListCommand**](AwsVpcListCommand.md) |  | 

### Return type

[**[]AwsVpcDto**](AwsVpcDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AwsZones

> AzResult AwsZones(ctx).AmazonAvailabilityZonesCommand(amazonAvailabilityZonesCommand).Execute()

Fetch Aws zones

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
	amazonAvailabilityZonesCommand := *openapiclient.NewAmazonAvailabilityZonesCommand("Region_example", "AwsAccessKeyId_example", "AwsSecretAccessKey_example", "RoleArn_example", "ExternalId_example", NullableInt32(123)) // AmazonAvailabilityZonesCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AWSCloudCredentialAPI.AwsZones(context.Background()).AmazonAvailabilityZonesCommand(amazonAvailabilityZonesCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AWSCloudCredentialAPI.AwsZones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AwsZones`: AzResult
	fmt.Fprintf(os.Stdout, "Response from `AWSCloudCredentialAPI.AwsZones`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAwsZonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amazonAvailabilityZonesCommand** | [**AmazonAvailabilityZonesCommand**](AmazonAvailabilityZonesCommand.md) |  | 

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

