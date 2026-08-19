# \InternalAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InternalAlertsList**](InternalAPI.md#InternalAlertsList) | **Get** /api/v1/internal/alerts | Retrieve alerts list
[**InternalExporterDomainsList**](InternalAPI.md#InternalExporterDomainsList) | **Get** /api/v1/internal/prometheus-exporter/domains | Retrieve exporter domains list
[**InternalExporterProjectsList**](InternalAPI.md#InternalExporterProjectsList) | **Get** /api/v1/internal/prometheus-exporter/projects | Retrieve exporter projects list
[**InternalExporterServersList**](InternalAPI.md#InternalExporterServersList) | **Get** /api/v1/internal/prometheus-exporter/servers | Retrieve exporter servers list
[**InternalExporterStandaloneVmsList**](InternalAPI.md#InternalExporterStandaloneVmsList) | **Get** /api/v1/internal/prometheus-exporter/standalone | Retrieve exporter standalone VMs list
[**InternalGetUserInfo**](InternalAPI.md#InternalGetUserInfo) | **Post** /api/v1/internal/get-user-info | Get user info
[**InternalProjectsForBilling**](InternalAPI.md#InternalProjectsForBilling) | **Get** /api/v1/internal/projects-for-billing | Retrieve projects list for internal billing
[**InternalProjectsList**](InternalAPI.md#InternalProjectsList) | **Get** /api/v1/internal/projects | Retrieve projects list



## InternalAlertsList

> []InternalAlertDto InternalAlertsList(ctx).ProjectId(projectId).Limit(limit).LastAlertId(lastAlertId).Execute()

Retrieve alerts list

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
	limit := int32(56) // int32 | 
	lastAlertId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InternalAPI.InternalAlertsList(context.Background()).ProjectId(projectId).Limit(limit).LastAlertId(lastAlertId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternalAPI.InternalAlertsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InternalAlertsList`: []InternalAlertDto
	fmt.Fprintf(os.Stdout, "Response from `InternalAPI.InternalAlertsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInternalAlertsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **int32** |  | 
 **limit** | **int32** |  | 
 **lastAlertId** | **int32** |  | 

### Return type

[**[]InternalAlertDto**](InternalAlertDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InternalExporterDomainsList

> []InternalExporterDomainDto InternalExporterDomainsList(ctx).Limit(limit).LastDomainId(lastDomainId).Execute()

Retrieve exporter domains list

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
	limit := int32(56) // int32 |  (optional) (default to 5000)
	lastDomainId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InternalAPI.InternalExporterDomainsList(context.Background()).Limit(limit).LastDomainId(lastDomainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternalAPI.InternalExporterDomainsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InternalExporterDomainsList`: []InternalExporterDomainDto
	fmt.Fprintf(os.Stdout, "Response from `InternalAPI.InternalExporterDomainsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInternalExporterDomainsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 5000]
 **lastDomainId** | **int32** |  | 

### Return type

[**[]InternalExporterDomainDto**](InternalExporterDomainDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InternalExporterProjectsList

> []InternalExporterProjectDto InternalExporterProjectsList(ctx).Limit(limit).LastProjectId(lastProjectId).Execute()

Retrieve exporter projects list

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
	limit := int32(56) // int32 |  (optional) (default to 5000)
	lastProjectId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InternalAPI.InternalExporterProjectsList(context.Background()).Limit(limit).LastProjectId(lastProjectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternalAPI.InternalExporterProjectsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InternalExporterProjectsList`: []InternalExporterProjectDto
	fmt.Fprintf(os.Stdout, "Response from `InternalAPI.InternalExporterProjectsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInternalExporterProjectsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 5000]
 **lastProjectId** | **int32** |  | 

### Return type

[**[]InternalExporterProjectDto**](InternalExporterProjectDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InternalExporterServersList

> []InternalExporterServerDto InternalExporterServersList(ctx).Limit(limit).LastServerId(lastServerId).Execute()

Retrieve exporter servers list

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
	limit := int32(56) // int32 |  (optional) (default to 5000)
	lastServerId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InternalAPI.InternalExporterServersList(context.Background()).Limit(limit).LastServerId(lastServerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternalAPI.InternalExporterServersList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InternalExporterServersList`: []InternalExporterServerDto
	fmt.Fprintf(os.Stdout, "Response from `InternalAPI.InternalExporterServersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInternalExporterServersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 5000]
 **lastServerId** | **int32** |  | 

### Return type

[**[]InternalExporterServerDto**](InternalExporterServerDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InternalExporterStandaloneVmsList

> []InternalExporterStandaloneVmDto InternalExporterStandaloneVmsList(ctx).Limit(limit).LastStandaloneVmId(lastStandaloneVmId).Execute()

Retrieve exporter standalone VMs list

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
	limit := int32(56) // int32 |  (optional) (default to 5000)
	lastStandaloneVmId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InternalAPI.InternalExporterStandaloneVmsList(context.Background()).Limit(limit).LastStandaloneVmId(lastStandaloneVmId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternalAPI.InternalExporterStandaloneVmsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InternalExporterStandaloneVmsList`: []InternalExporterStandaloneVmDto
	fmt.Fprintf(os.Stdout, "Response from `InternalAPI.InternalExporterStandaloneVmsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInternalExporterStandaloneVmsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 5000]
 **lastStandaloneVmId** | **int32** |  | 

### Return type

[**[]InternalExporterStandaloneVmDto**](InternalExporterStandaloneVmDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InternalGetUserInfo

> UserAuthContext InternalGetUserInfo(ctx).InternalUserInfoListCommand(internalUserInfoListCommand).Execute()

Get user info

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
	internalUserInfoListCommand := *openapiclient.NewInternalUserInfoListCommand("Token_example") // InternalUserInfoListCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InternalAPI.InternalGetUserInfo(context.Background()).InternalUserInfoListCommand(internalUserInfoListCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternalAPI.InternalGetUserInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InternalGetUserInfo`: UserAuthContext
	fmt.Fprintf(os.Stdout, "Response from `InternalAPI.InternalGetUserInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInternalGetUserInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **internalUserInfoListCommand** | [**InternalUserInfoListCommand**](InternalUserInfoListCommand.md) |  | 

### Return type

[**UserAuthContext**](UserAuthContext.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InternalProjectsForBilling

> []ProjectsForBillingDto InternalProjectsForBilling(ctx).Limit(limit).LastProjectId(lastProjectId).Execute()

Retrieve projects list for internal billing

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
	limit := int32(56) // int32 | 
	lastProjectId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InternalAPI.InternalProjectsForBilling(context.Background()).Limit(limit).LastProjectId(lastProjectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternalAPI.InternalProjectsForBilling``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InternalProjectsForBilling`: []ProjectsForBillingDto
	fmt.Fprintf(os.Stdout, "Response from `InternalAPI.InternalProjectsForBilling`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInternalProjectsForBillingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **lastProjectId** | **int32** |  | 

### Return type

[**[]ProjectsForBillingDto**](ProjectsForBillingDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InternalProjectsList

> []InternalProjectListDto InternalProjectsList(ctx).Limit(limit).ProjectId(projectId).LastProjectId(lastProjectId).Execute()

Retrieve projects list

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
	limit := int32(56) // int32 | 
	projectId := int32(56) // int32 |  (optional)
	lastProjectId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InternalAPI.InternalProjectsList(context.Background()).Limit(limit).ProjectId(projectId).LastProjectId(lastProjectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternalAPI.InternalProjectsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InternalProjectsList`: []InternalProjectListDto
	fmt.Fprintf(os.Stdout, "Response from `InternalAPI.InternalProjectsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInternalProjectsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **projectId** | **int32** |  | 
 **lastProjectId** | **int32** |  | 

### Return type

[**[]InternalProjectListDto**](InternalProjectListDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

