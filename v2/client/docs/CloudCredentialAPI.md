# \CloudCredentialAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudcredentialsAllFlavors**](CloudCredentialAPI.md#CloudcredentialsAllFlavors) | **Get** /api/v1/cloudcredentials/flavors/{cloudId} | Retrieve all flavors
[**CloudcredentialsDelete**](CloudCredentialAPI.md#CloudcredentialsDelete) | **Delete** /api/v1/cloudcredentials/{cloudId} | Remove cloud credential by cloud Id
[**CloudcredentialsExceeded**](CloudCredentialAPI.md#CloudcredentialsExceeded) | **Get** /api/v1/cloudcredentials/exceeded-quotas | Retrieve cloud credentials exceeded quotas
[**CloudcredentialsGpuFilterMetadata**](CloudCredentialAPI.md#CloudcredentialsGpuFilterMetadata) | **Get** /api/v1/cloudcredentials/flavors/gpu-filter-metadata/{cloudId} | Retrieve GPU filter metadata for a cloud credential
[**CloudcredentialsLockManager**](CloudCredentialAPI.md#CloudcredentialsLockManager) | **Post** /api/v1/cloudcredentials/lockmanager | Lock/Unlock cloud credential
[**CloudcredentialsMakeDefault**](CloudCredentialAPI.md#CloudcredentialsMakeDefault) | **Post** /api/v1/cloudcredentials/makedefault | Make cloud credentials default
[**CloudcredentialsOrgList**](CloudCredentialAPI.md#CloudcredentialsOrgList) | **Get** /api/v1/cloudcredentials | Retrieve a list of cloud credentials by organization Id



## CloudcredentialsAllFlavors

> AllFlavorsList CloudcredentialsAllFlavors(ctx, cloudId).Limit(limit).Offset(offset).StartRam(startRam).EndRam(endRam).StartCpu(startCpu).EndCpu(endCpu).Search(search).FlavorName(flavorName).SortBy(sortBy).SortDirection(sortDirection).HasGpuSupport(hasGpuSupport).StartGpuCount(startGpuCount).EndGpuCount(endGpuCount).GpuModel(gpuModel).Execute()

Retrieve all flavors

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
	limit := int32(56) // int32 |  (optional)
	offset := int32(56) // int32 |  (optional)
	startRam := float64(1.2) // float64 |  (optional)
	endRam := float64(1.2) // float64 |  (optional)
	startCpu := int32(56) // int32 |  (optional)
	endCpu := int32(56) // int32 |  (optional)
	search := "search_example" // string |  (optional)
	flavorName := "flavorName_example" // string |  (optional)
	sortBy := "sortBy_example" // string |  (optional)
	sortDirection := "sortDirection_example" // string |  (optional)
	hasGpuSupport := true // bool |  (optional)
	startGpuCount := int32(56) // int32 |  (optional)
	endGpuCount := int32(56) // int32 |  (optional)
	gpuModel := "gpuModel_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudCredentialAPI.CloudcredentialsAllFlavors(context.Background(), cloudId).Limit(limit).Offset(offset).StartRam(startRam).EndRam(endRam).StartCpu(startCpu).EndCpu(endCpu).Search(search).FlavorName(flavorName).SortBy(sortBy).SortDirection(sortDirection).HasGpuSupport(hasGpuSupport).StartGpuCount(startGpuCount).EndGpuCount(endGpuCount).GpuModel(gpuModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudCredentialAPI.CloudcredentialsAllFlavors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudcredentialsAllFlavors`: AllFlavorsList
	fmt.Fprintf(os.Stdout, "Response from `CloudCredentialAPI.CloudcredentialsAllFlavors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudcredentialsAllFlavorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **startRam** | **float64** |  | 
 **endRam** | **float64** |  | 
 **startCpu** | **int32** |  | 
 **endCpu** | **int32** |  | 
 **search** | **string** |  | 
 **flavorName** | **string** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **hasGpuSupport** | **bool** |  | 
 **startGpuCount** | **int32** |  | 
 **endGpuCount** | **int32** |  | 
 **gpuModel** | **string** |  | 

### Return type

[**AllFlavorsList**](AllFlavorsList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudcredentialsDelete

> CloudcredentialsDelete(ctx, cloudId).Execute()

Remove cloud credential by cloud Id

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CloudCredentialAPI.CloudcredentialsDelete(context.Background(), cloudId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudCredentialAPI.CloudcredentialsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudcredentialsDeleteRequest struct via the builder pattern


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


## CloudcredentialsExceeded

> ExceededQuotaList CloudcredentialsExceeded(ctx).OrganizationId(organizationId).Execute()

Retrieve cloud credentials exceeded quotas

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudCredentialAPI.CloudcredentialsExceeded(context.Background()).OrganizationId(organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudCredentialAPI.CloudcredentialsExceeded``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudcredentialsExceeded`: ExceededQuotaList
	fmt.Fprintf(os.Stdout, "Response from `CloudCredentialAPI.CloudcredentialsExceeded`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudcredentialsExceededRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 

### Return type

[**ExceededQuotaList**](ExceededQuotaList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudcredentialsGpuFilterMetadata

> FlavorGpuFilterMetadataList CloudcredentialsGpuFilterMetadata(ctx, cloudId).Execute()

Retrieve GPU filter metadata for a cloud credential

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudCredentialAPI.CloudcredentialsGpuFilterMetadata(context.Background(), cloudId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudCredentialAPI.CloudcredentialsGpuFilterMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudcredentialsGpuFilterMetadata`: FlavorGpuFilterMetadataList
	fmt.Fprintf(os.Stdout, "Response from `CloudCredentialAPI.CloudcredentialsGpuFilterMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudcredentialsGpuFilterMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FlavorGpuFilterMetadataList**](FlavorGpuFilterMetadataList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudcredentialsLockManager

> CloudcredentialsLockManager(ctx).CloudLockManagerCommand(cloudLockManagerCommand).Execute()

Lock/Unlock cloud credential

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
	cloudLockManagerCommand := *openapiclient.NewCloudLockManagerCommand(int32(123), "Mode_example") // CloudLockManagerCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CloudCredentialAPI.CloudcredentialsLockManager(context.Background()).CloudLockManagerCommand(cloudLockManagerCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudCredentialAPI.CloudcredentialsLockManager``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudcredentialsLockManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudLockManagerCommand** | [**CloudLockManagerCommand**](CloudLockManagerCommand.md) |  | 

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


## CloudcredentialsMakeDefault

> CloudcredentialsMakeDefault(ctx).CredentialMakeDefaultCommand(credentialMakeDefaultCommand).Execute()

Make cloud credentials default

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
	credentialMakeDefaultCommand := *openapiclient.NewCredentialMakeDefaultCommand(int32(123)) // CredentialMakeDefaultCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CloudCredentialAPI.CloudcredentialsMakeDefault(context.Background()).CredentialMakeDefaultCommand(credentialMakeDefaultCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudCredentialAPI.CloudcredentialsMakeDefault``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudcredentialsMakeDefaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **credentialMakeDefaultCommand** | [**CredentialMakeDefaultCommand**](CredentialMakeDefaultCommand.md) |  | 

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


## CloudcredentialsOrgList

> []CloudCredentialsForOrganizationEntity CloudcredentialsOrgList(ctx).DomainId(domainId).OrganizationId(organizationId).Search(search).IsAdmin(isAdmin).IsInfra(isInfra).Id(id).CloudType(cloudType).Execute()

Retrieve a list of cloud credentials by organization Id

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
	organizationId := int32(56) // int32 |  (optional)
	search := "search_example" // string |  (optional)
	isAdmin := true // bool |  (optional)
	isInfra := true // bool |  (optional)
	id := int32(56) // int32 |  (optional)
	cloudType := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudCredentialAPI.CloudcredentialsOrgList(context.Background()).DomainId(domainId).OrganizationId(organizationId).Search(search).IsAdmin(isAdmin).IsInfra(isInfra).Id(id).CloudType(cloudType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudCredentialAPI.CloudcredentialsOrgList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudcredentialsOrgList`: []CloudCredentialsForOrganizationEntity
	fmt.Fprintf(os.Stdout, "Response from `CloudCredentialAPI.CloudcredentialsOrgList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudcredentialsOrgListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainId** | **int32** |  | 
 **organizationId** | **int32** |  | 
 **search** | **string** |  | 
 **isAdmin** | **bool** |  | 
 **isInfra** | **bool** |  | 
 **id** | **int32** |  | 
 **cloudType** | **map[string]interface{}** |  | 

### Return type

[**[]CloudCredentialsForOrganizationEntity**](CloudCredentialsForOrganizationEntity.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

