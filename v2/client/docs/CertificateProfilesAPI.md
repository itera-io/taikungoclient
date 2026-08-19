# \CertificateProfilesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CertificateProfilesCreate**](CertificateProfilesAPI.md#CertificateProfilesCreate) | **Post** /api/v1/certificate-profiles | Create certificate profile (unified issuerMode DTO or legacy kind-based)
[**CertificateProfilesDelete**](CertificateProfilesAPI.md#CertificateProfilesDelete) | **Delete** /api/v1/certificate-profiles/{id} | Delete certificate profile
[**CertificateProfilesDropdown**](CertificateProfilesAPI.md#CertificateProfilesDropdown) | **Get** /api/v1/certificate-profiles | Certificate profile dropdown (use ?kind&#x3D;signing_ca for dns-cert CA picker)
[**CertificateProfilesGetTrustBundle**](CertificateProfilesAPI.md#CertificateProfilesGetTrustBundle) | **Get** /api/v1/certificate-profiles/{id}/trust-bundle | Get trust bundle items (name + PEM) for a certificate profile
[**CertificateProfilesList**](CertificateProfilesAPI.md#CertificateProfilesList) | **Get** /api/v1/certificate-profiles/list | List certificate profiles (optional kind filter: signing_ca, trust_anchor, tls_leaf)
[**CertificateProfilesLockManagement**](CertificateProfilesAPI.md#CertificateProfilesLockManagement) | **Post** /api/v1/certificate-profiles/lockmanager | Lock/unlock certificate profile
[**CertificateProfilesMakeDefault**](CertificateProfilesAPI.md#CertificateProfilesMakeDefault) | **Post** /api/v1/certificate-profiles/makedefault | Make default certificate profile for organization
[**CertificateProfilesParseCertificate**](CertificateProfilesAPI.md#CertificateProfilesParseCertificate) | **Post** /api/v1/certificate-profiles/parse-certificate | Parse a PEM-encoded certificate and return CN and expiry
[**CertificateProfilesUpdate**](CertificateProfilesAPI.md#CertificateProfilesUpdate) | **Put** /api/v1/certificate-profiles | Update certificate profile (name, issuerMode template, trust bundle, and related fields)
[**CertificateProfilesValidate**](CertificateProfilesAPI.md#CertificateProfilesValidate) | **Post** /api/v1/certificate-profiles/validate | Validate certificate profile PEMs or stored profile by certificateProfileId
[**CertificateProfilesValidateTrustBundle**](CertificateProfilesAPI.md#CertificateProfilesValidateTrustBundle) | **Post** /api/v1/certificate-profiles/validate-trust-bundle | Validate a PEM bundle and return the count of parseable certificates



## CertificateProfilesCreate

> ApiResponse CertificateProfilesCreate(ctx).CertificateProfileCreateCommand(certificateProfileCreateCommand).Execute()

Create certificate profile (unified issuerMode DTO or legacy kind-based)

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
	certificateProfileCreateCommand := *openapiclient.NewCertificateProfileCreateCommand("Name_example", NullableInt32(123)) // CertificateProfileCreateCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificateProfilesAPI.CertificateProfilesCreate(context.Background()).CertificateProfileCreateCommand(certificateProfileCreateCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateProfilesAPI.CertificateProfilesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CertificateProfilesCreate`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `CertificateProfilesAPI.CertificateProfilesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateProfilesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **certificateProfileCreateCommand** | [**CertificateProfileCreateCommand**](CertificateProfileCreateCommand.md) |  | 

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


## CertificateProfilesDelete

> CertificateProfilesDelete(ctx, id).Execute()

Delete certificate profile

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
	r, err := apiClient.CertificateProfilesAPI.CertificateProfilesDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateProfilesAPI.CertificateProfilesDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCertificateProfilesDeleteRequest struct via the builder pattern


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


## CertificateProfilesDropdown

> []CertificateProfileDropdownDto CertificateProfilesDropdown(ctx).OrganizationId(organizationId).Search(search).Kind(kind).Execute()

Certificate profile dropdown (use ?kind=signing_ca for dns-cert CA picker)

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
	kind := "kind_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificateProfilesAPI.CertificateProfilesDropdown(context.Background()).OrganizationId(organizationId).Search(search).Kind(kind).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateProfilesAPI.CertificateProfilesDropdown``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CertificateProfilesDropdown`: []CertificateProfileDropdownDto
	fmt.Fprintf(os.Stdout, "Response from `CertificateProfilesAPI.CertificateProfilesDropdown`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateProfilesDropdownRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 
 **search** | **string** |  | 
 **kind** | **string** |  | 

### Return type

[**[]CertificateProfileDropdownDto**](CertificateProfileDropdownDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateProfilesGetTrustBundle

> CertificateProfileTrustBundleResponse CertificateProfilesGetTrustBundle(ctx, id).Execute()

Get trust bundle items (name + PEM) for a certificate profile

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
	resp, r, err := apiClient.CertificateProfilesAPI.CertificateProfilesGetTrustBundle(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateProfilesAPI.CertificateProfilesGetTrustBundle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CertificateProfilesGetTrustBundle`: CertificateProfileTrustBundleResponse
	fmt.Fprintf(os.Stdout, "Response from `CertificateProfilesAPI.CertificateProfilesGetTrustBundle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificateProfilesGetTrustBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CertificateProfileTrustBundleResponse**](CertificateProfileTrustBundleResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateProfilesList

> CertificateProfileListResponse CertificateProfilesList(ctx).OrganizationId(organizationId).DomainId(domainId).Search(search).SearchId(searchId).Id(id).Kind(kind).SortBy(sortBy).SortDirection(sortDirection).Limit(limit).Offset(offset).Execute()

List certificate profiles (optional kind filter: signing_ca, trust_anchor, tls_leaf)

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
	kind := "kind_example" // string |  (optional)
	sortBy := "sortBy_example" // string |  (optional)
	sortDirection := "sortDirection_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)
	offset := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificateProfilesAPI.CertificateProfilesList(context.Background()).OrganizationId(organizationId).DomainId(domainId).Search(search).SearchId(searchId).Id(id).Kind(kind).SortBy(sortBy).SortDirection(sortDirection).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateProfilesAPI.CertificateProfilesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CertificateProfilesList`: CertificateProfileListResponse
	fmt.Fprintf(os.Stdout, "Response from `CertificateProfilesAPI.CertificateProfilesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateProfilesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 
 **domainId** | **int32** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **id** | **int32** |  | 
 **kind** | **string** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **limit** | **int32** |  | [default to 50]
 **offset** | **int32** |  | [default to 0]

### Return type

[**CertificateProfileListResponse**](CertificateProfileListResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateProfilesLockManagement

> CertificateProfilesLockManagement(ctx).CertificateProfileLockCommand(certificateProfileLockCommand).Execute()

Lock/unlock certificate profile

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
	certificateProfileLockCommand := *openapiclient.NewCertificateProfileLockCommand(int32(123), "Mode_example") // CertificateProfileLockCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CertificateProfilesAPI.CertificateProfilesLockManagement(context.Background()).CertificateProfileLockCommand(certificateProfileLockCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateProfilesAPI.CertificateProfilesLockManagement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateProfilesLockManagementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **certificateProfileLockCommand** | [**CertificateProfileLockCommand**](CertificateProfileLockCommand.md) |  | 

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


## CertificateProfilesMakeDefault

> CertificateProfilesMakeDefault(ctx).CertificateProfileMakeDefaultCommand(certificateProfileMakeDefaultCommand).Execute()

Make default certificate profile for organization

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
	certificateProfileMakeDefaultCommand := *openapiclient.NewCertificateProfileMakeDefaultCommand(int32(123)) // CertificateProfileMakeDefaultCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CertificateProfilesAPI.CertificateProfilesMakeDefault(context.Background()).CertificateProfileMakeDefaultCommand(certificateProfileMakeDefaultCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateProfilesAPI.CertificateProfilesMakeDefault``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateProfilesMakeDefaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **certificateProfileMakeDefaultCommand** | [**CertificateProfileMakeDefaultCommand**](CertificateProfileMakeDefaultCommand.md) |  | 

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


## CertificateProfilesParseCertificate

> ParseCertificateResponse CertificateProfilesParseCertificate(ctx).ParseCertificateQuery(parseCertificateQuery).Execute()

Parse a PEM-encoded certificate and return CN and expiry

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
	parseCertificateQuery := *openapiclient.NewParseCertificateQuery("CertPem_example") // ParseCertificateQuery |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificateProfilesAPI.CertificateProfilesParseCertificate(context.Background()).ParseCertificateQuery(parseCertificateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateProfilesAPI.CertificateProfilesParseCertificate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CertificateProfilesParseCertificate`: ParseCertificateResponse
	fmt.Fprintf(os.Stdout, "Response from `CertificateProfilesAPI.CertificateProfilesParseCertificate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateProfilesParseCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **parseCertificateQuery** | [**ParseCertificateQuery**](ParseCertificateQuery.md) |  | 

### Return type

[**ParseCertificateResponse**](ParseCertificateResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateProfilesUpdate

> CertificateProfilesUpdate(ctx).CertificateProfileUpdateCommand(certificateProfileUpdateCommand).Execute()

Update certificate profile (name, issuerMode template, trust bundle, and related fields)

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
	certificateProfileUpdateCommand := *openapiclient.NewCertificateProfileUpdateCommand(int32(123), "Name_example") // CertificateProfileUpdateCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CertificateProfilesAPI.CertificateProfilesUpdate(context.Background()).CertificateProfileUpdateCommand(certificateProfileUpdateCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateProfilesAPI.CertificateProfilesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateProfilesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **certificateProfileUpdateCommand** | [**CertificateProfileUpdateCommand**](CertificateProfileUpdateCommand.md) |  | 

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


## CertificateProfilesValidate

> CertificateProfilesValidate(ctx).CertificateProfileValidateCommand(certificateProfileValidateCommand).Execute()

Validate certificate profile PEMs or stored profile by certificateProfileId

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
	certificateProfileValidateCommand := *openapiclient.NewCertificateProfileValidateCommand(NullableInt32(123)) // CertificateProfileValidateCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CertificateProfilesAPI.CertificateProfilesValidate(context.Background()).CertificateProfileValidateCommand(certificateProfileValidateCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateProfilesAPI.CertificateProfilesValidate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateProfilesValidateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **certificateProfileValidateCommand** | [**CertificateProfileValidateCommand**](CertificateProfileValidateCommand.md) |  | 

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


## CertificateProfilesValidateTrustBundle

> ValidateTrustBundleResponse CertificateProfilesValidateTrustBundle(ctx).ValidateTrustBundleQuery(validateTrustBundleQuery).Execute()

Validate a PEM bundle and return the count of parseable certificates

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
	validateTrustBundleQuery := *openapiclient.NewValidateTrustBundleQuery("PemBundle_example") // ValidateTrustBundleQuery |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificateProfilesAPI.CertificateProfilesValidateTrustBundle(context.Background()).ValidateTrustBundleQuery(validateTrustBundleQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateProfilesAPI.CertificateProfilesValidateTrustBundle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CertificateProfilesValidateTrustBundle`: ValidateTrustBundleResponse
	fmt.Fprintf(os.Stdout, "Response from `CertificateProfilesAPI.CertificateProfilesValidateTrustBundle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateProfilesValidateTrustBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validateTrustBundleQuery** | [**ValidateTrustBundleQuery**](ValidateTrustBundleQuery.md) |  | 

### Return type

[**ValidateTrustBundleResponse**](ValidateTrustBundleResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

