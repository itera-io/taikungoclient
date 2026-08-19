# \DNSCredentialsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DnscredentialsAttachProject**](DNSCredentialsAPI.md#DnscredentialsAttachProject) | **Post** /api/v1/dnscredentials/attach-project | Attach a DNS credential to a project
[**DnscredentialsCreate**](DNSCredentialsAPI.md#DnscredentialsCreate) | **Post** /api/v1/dnscredentials | Add DNS credential
[**DnscredentialsDelete**](DNSCredentialsAPI.md#DnscredentialsDelete) | **Delete** /api/v1/dnscredentials/{id} | Delete DNS credential
[**DnscredentialsDetachProject**](DNSCredentialsAPI.md#DnscredentialsDetachProject) | **Post** /api/v1/dnscredentials/detach-project | Detach the DNS credential from a project
[**DnscredentialsDropdown**](DNSCredentialsAPI.md#DnscredentialsDropdown) | **Get** /api/v1/dnscredentials | Retrieve DNS credentials for organization dropdown
[**DnscredentialsList**](DNSCredentialsAPI.md#DnscredentialsList) | **Get** /api/v1/dnscredentials/list | Retrieve all DNS credentials
[**DnscredentialsLockManagement**](DNSCredentialsAPI.md#DnscredentialsLockManagement) | **Post** /api/v1/dnscredentials/lockmanager | Lock/unlock DNS credential
[**DnscredentialsMakeDefault**](DNSCredentialsAPI.md#DnscredentialsMakeDefault) | **Post** /api/v1/dnscredentials/makedefault | Make default DNS credential
[**DnscredentialsUpdate**](DNSCredentialsAPI.md#DnscredentialsUpdate) | **Put** /api/v1/dnscredentials | Update DNS credential
[**DnscredentialsValidate**](DNSCredentialsAPI.md#DnscredentialsValidate) | **Post** /api/v1/dnscredentials/validate | Validate DNS provider credentials



## DnscredentialsAttachProject

> DnscredentialsAttachProject(ctx).AttachDetachDnsCredentialCommand(attachDetachDnsCredentialCommand).Execute()

Attach a DNS credential to a project

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
	attachDetachDnsCredentialCommand := *openapiclient.NewAttachDetachDnsCredentialCommand(int32(123), NullableInt32(123)) // AttachDetachDnsCredentialCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DNSCredentialsAPI.DnscredentialsAttachProject(context.Background()).AttachDetachDnsCredentialCommand(attachDetachDnsCredentialCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSCredentialsAPI.DnscredentialsAttachProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnscredentialsAttachProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attachDetachDnsCredentialCommand** | [**AttachDetachDnsCredentialCommand**](AttachDetachDnsCredentialCommand.md) |  | 

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


## DnscredentialsCreate

> ApiResponse DnscredentialsCreate(ctx).DnsCredentialCreateCommand(dnsCredentialCreateCommand).Execute()

Add DNS credential

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
	dnsCredentialCreateCommand := *openapiclient.NewDnsCredentialCreateCommand("Name_example", "ProviderType_example", NullableInt32(123), "DomainFilter_example", false, "ApiToken_example", "AccessKeyId_example", "SecretAccessKey_example", "Region_example", "HostedZoneId_example", "TsigKeyName_example", "TsigSecret_example", "TsigAlgorithm_example", "NameServer_example", "AzureTenantId_example", "AzureSubscriptionId_example", "AzureResourceGroup_example", "AzureClientId_example", "AzureClientSecret_example", "GcpProject_example", "GcpServiceAccountJson_example", "AuthMethod_example", "AuthVariant_example", "AwsRoleArn_example", "GcpServiceAccountEmail_example") // DnsCredentialCreateCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DNSCredentialsAPI.DnscredentialsCreate(context.Background()).DnsCredentialCreateCommand(dnsCredentialCreateCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSCredentialsAPI.DnscredentialsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DnscredentialsCreate`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `DNSCredentialsAPI.DnscredentialsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnscredentialsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dnsCredentialCreateCommand** | [**DnsCredentialCreateCommand**](DnsCredentialCreateCommand.md) |  | 

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


## DnscredentialsDelete

> DnscredentialsDelete(ctx, id).Execute()

Delete DNS credential

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
	r, err := apiClient.DNSCredentialsAPI.DnscredentialsDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSCredentialsAPI.DnscredentialsDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDnscredentialsDeleteRequest struct via the builder pattern


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


## DnscredentialsDetachProject

> DnscredentialsDetachProject(ctx).AttachDetachDnsCredentialCommand(attachDetachDnsCredentialCommand).Execute()

Detach the DNS credential from a project

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
	attachDetachDnsCredentialCommand := *openapiclient.NewAttachDetachDnsCredentialCommand(int32(123), NullableInt32(123)) // AttachDetachDnsCredentialCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DNSCredentialsAPI.DnscredentialsDetachProject(context.Background()).AttachDetachDnsCredentialCommand(attachDetachDnsCredentialCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSCredentialsAPI.DnscredentialsDetachProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnscredentialsDetachProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attachDetachDnsCredentialCommand** | [**AttachDetachDnsCredentialCommand**](AttachDetachDnsCredentialCommand.md) |  | 

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


## DnscredentialsDropdown

> []DnsCredentialDropdownDto DnscredentialsDropdown(ctx).OrganizationId(organizationId).Search(search).IncludeLocked(includeLocked).Execute()

Retrieve DNS credentials for organization dropdown

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
	search := "search_example" // string |  (optional)
	includeLocked := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DNSCredentialsAPI.DnscredentialsDropdown(context.Background()).OrganizationId(organizationId).Search(search).IncludeLocked(includeLocked).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSCredentialsAPI.DnscredentialsDropdown``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DnscredentialsDropdown`: []DnsCredentialDropdownDto
	fmt.Fprintf(os.Stdout, "Response from `DNSCredentialsAPI.DnscredentialsDropdown`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnscredentialsDropdownRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 
 **search** | **string** |  | 
 **includeLocked** | **bool** |  | [default to false]

### Return type

[**[]DnsCredentialDropdownDto**](DnsCredentialDropdownDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnscredentialsList

> DnsCredentialListResponse DnscredentialsList(ctx).OrganizationId(organizationId).DomainId(domainId).Search(search).SearchId(searchId).Id(id).SortBy(sortBy).SortDirection(sortDirection).Limit(limit).Offset(offset).Execute()

Retrieve all DNS credentials

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
	domainId := int32(56) // int32 |  (optional)
	search := "search_example" // string |  (optional)
	searchId := "searchId_example" // string |  (optional)
	id := int32(56) // int32 |  (optional)
	sortBy := "sortBy_example" // string |  (optional)
	sortDirection := "sortDirection_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)
	offset := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DNSCredentialsAPI.DnscredentialsList(context.Background()).OrganizationId(organizationId).DomainId(domainId).Search(search).SearchId(searchId).Id(id).SortBy(sortBy).SortDirection(sortDirection).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSCredentialsAPI.DnscredentialsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DnscredentialsList`: DnsCredentialListResponse
	fmt.Fprintf(os.Stdout, "Response from `DNSCredentialsAPI.DnscredentialsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnscredentialsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 
 **domainId** | **int32** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **id** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **limit** | **int32** |  | [default to 50]
 **offset** | **int32** |  | [default to 0]

### Return type

[**DnsCredentialListResponse**](DnsCredentialListResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnscredentialsLockManagement

> DnscredentialsLockManagement(ctx).DnsCredentialLockCommand(dnsCredentialLockCommand).Execute()

Lock/unlock DNS credential

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
	dnsCredentialLockCommand := *openapiclient.NewDnsCredentialLockCommand(int32(123), "Mode_example") // DnsCredentialLockCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DNSCredentialsAPI.DnscredentialsLockManagement(context.Background()).DnsCredentialLockCommand(dnsCredentialLockCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSCredentialsAPI.DnscredentialsLockManagement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnscredentialsLockManagementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dnsCredentialLockCommand** | [**DnsCredentialLockCommand**](DnsCredentialLockCommand.md) |  | 

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


## DnscredentialsMakeDefault

> DnscredentialsMakeDefault(ctx).DnsCredentialMakeDefaultCommand(dnsCredentialMakeDefaultCommand).Execute()

Make default DNS credential

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
	dnsCredentialMakeDefaultCommand := *openapiclient.NewDnsCredentialMakeDefaultCommand(int32(123)) // DnsCredentialMakeDefaultCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DNSCredentialsAPI.DnscredentialsMakeDefault(context.Background()).DnsCredentialMakeDefaultCommand(dnsCredentialMakeDefaultCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSCredentialsAPI.DnscredentialsMakeDefault``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnscredentialsMakeDefaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dnsCredentialMakeDefaultCommand** | [**DnsCredentialMakeDefaultCommand**](DnsCredentialMakeDefaultCommand.md) |  | 

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


## DnscredentialsUpdate

> DnscredentialsUpdate(ctx).DnsCredentialUpdateCommand(dnsCredentialUpdateCommand).Execute()

Update DNS credential

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
	dnsCredentialUpdateCommand := *openapiclient.NewDnsCredentialUpdateCommand(int32(123), "Name_example", "ProviderType_example", "DomainFilter_example", false) // DnsCredentialUpdateCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DNSCredentialsAPI.DnscredentialsUpdate(context.Background()).DnsCredentialUpdateCommand(dnsCredentialUpdateCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSCredentialsAPI.DnscredentialsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnscredentialsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dnsCredentialUpdateCommand** | [**DnsCredentialUpdateCommand**](DnsCredentialUpdateCommand.md) |  | 

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


## DnscredentialsValidate

> DnscredentialsValidate(ctx).ValidateDnsCertCommand(validateDnsCertCommand).Execute()

Validate DNS provider credentials

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
	validateDnsCertCommand := *openapiclient.NewValidateDnsCertCommand("DnsProvider_example", "CertIssuerType_example", "Domain_example", "ApiToken_example", "ZoneId_example", "AccessKeyId_example", "SecretAccessKey_example", "Region_example", "HostedZoneId_example", "TsigKeyName_example", "TsigSecret_example", "TsigAlgorithm_example", "NameServer_example", "AzureTenantId_example", "AzureSubscriptionId_example", "AzureResourceGroup_example", "AzureClientId_example", "AzureClientSecret_example", "GcpProject_example", "GcpServiceAccountJson_example") // ValidateDnsCertCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DNSCredentialsAPI.DnscredentialsValidate(context.Background()).ValidateDnsCertCommand(validateDnsCertCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSCredentialsAPI.DnscredentialsValidate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnscredentialsValidateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validateDnsCertCommand** | [**ValidateDnsCertCommand**](ValidateDnsCertCommand.md) |  | 

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

