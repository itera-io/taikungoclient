# \DnsCertAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DnsCertAssignCertificateProfile**](DnsCertAPI.md#DnsCertAssignCertificateProfile) | **Post** /api/v1/dns-cert/assign-certificate-profile | Assign a signing certificate profile to a CA-issuer project
[**DnsCertCertManagerDisable**](DnsCertAPI.md#DnsCertCertManagerDisable) | **Post** /api/v1/dns-cert/cert-manager/disable | Disable cert-manager for a project
[**DnsCertCertManagerEnable**](DnsCertAPI.md#DnsCertCertManagerEnable) | **Post** /api/v1/dns-cert/cert-manager/enable | Enable cert-manager for a project
[**DnsCertDisable**](DnsCertAPI.md#DnsCertDisable) | **Post** /api/v1/dns-cert/disable | Disable DNS/Cert automation
[**DnsCertEnable**](DnsCertAPI.md#DnsCertEnable) | **Post** /api/v1/dns-cert/enable | Enable DNS/Cert automation
[**DnsCertExternalDnsDisable**](DnsCertAPI.md#DnsCertExternalDnsDisable) | **Post** /api/v1/dns-cert/external-dns/disable | Disable external-dns for a project
[**DnsCertExternalDnsEnable**](DnsCertAPI.md#DnsCertExternalDnsEnable) | **Post** /api/v1/dns-cert/external-dns/enable | Enable external-dns for a project
[**DnsCertStatus**](DnsCertAPI.md#DnsCertStatus) | **Get** /api/v1/dns-cert/{projectId} | Get DNS/Cert status for project
[**DnsCertSync**](DnsCertAPI.md#DnsCertSync) | **Post** /api/v1/dns-cert/sync | Sync DNS/Cert to cluster
[**DnsCertTrustManagerDisable**](DnsCertAPI.md#DnsCertTrustManagerDisable) | **Post** /api/v1/dns-cert/trust-manager/disable | Disable trust-manager for a project
[**DnsCertTrustManagerEnable**](DnsCertAPI.md#DnsCertTrustManagerEnable) | **Post** /api/v1/dns-cert/trust-manager/enable | Enable trust-manager for a project
[**DnsCertUnassignCertificateProfile**](DnsCertAPI.md#DnsCertUnassignCertificateProfile) | **Post** /api/v1/dns-cert/unassign-certificate-profile | Remove the signing certificate profile from a CA-issuer project
[**DnsCertValidate**](DnsCertAPI.md#DnsCertValidate) | **Post** /api/v1/dns-cert/validate | Validate DNS/Cert enablement inputs (provider credentials, signing CA reference, issuer rules)



## DnsCertAssignCertificateProfile

> DnsCertAssignCertificateProfile(ctx).AssignCertificateProfileCommand(assignCertificateProfileCommand).Execute()

Assign a signing certificate profile to a CA-issuer project

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
	assignCertificateProfileCommand := *openapiclient.NewAssignCertificateProfileCommand(int32(123), int32(123)) // AssignCertificateProfileCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DnsCertAPI.DnsCertAssignCertificateProfile(context.Background()).AssignCertificateProfileCommand(assignCertificateProfileCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsCertAPI.DnsCertAssignCertificateProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsCertAssignCertificateProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assignCertificateProfileCommand** | [**AssignCertificateProfileCommand**](AssignCertificateProfileCommand.md) |  | 

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


## DnsCertCertManagerDisable

> DnsCertCertManagerDisable(ctx).DisableCertManagerCommand(disableCertManagerCommand).Execute()

Disable cert-manager for a project

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
	disableCertManagerCommand := *openapiclient.NewDisableCertManagerCommand(int32(123)) // DisableCertManagerCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DnsCertAPI.DnsCertCertManagerDisable(context.Background()).DisableCertManagerCommand(disableCertManagerCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsCertAPI.DnsCertCertManagerDisable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsCertCertManagerDisableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **disableCertManagerCommand** | [**DisableCertManagerCommand**](DisableCertManagerCommand.md) |  | 

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


## DnsCertCertManagerEnable

> DnsCertCertManagerEnable(ctx).EnableCertManagerCommand(enableCertManagerCommand).Execute()

Enable cert-manager for a project

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
	enableCertManagerCommand := *openapiclient.NewEnableCertManagerCommand(int32(123), NullableInt32(123), "CertIssuerType_example", "AcmeEmail_example", "Domain_example") // EnableCertManagerCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DnsCertAPI.DnsCertCertManagerEnable(context.Background()).EnableCertManagerCommand(enableCertManagerCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsCertAPI.DnsCertCertManagerEnable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsCertCertManagerEnableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enableCertManagerCommand** | [**EnableCertManagerCommand**](EnableCertManagerCommand.md) |  | 

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


## DnsCertDisable

> DnsCertDisable(ctx).DisableDnsCertCommand(disableDnsCertCommand).Execute()

Disable DNS/Cert automation

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
	disableDnsCertCommand := *openapiclient.NewDisableDnsCertCommand(int32(123)) // DisableDnsCertCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DnsCertAPI.DnsCertDisable(context.Background()).DisableDnsCertCommand(disableDnsCertCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsCertAPI.DnsCertDisable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsCertDisableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **disableDnsCertCommand** | [**DisableDnsCertCommand**](DisableDnsCertCommand.md) |  | 

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


## DnsCertEnable

> DnsCertEnable(ctx).EnableDnsCertCommand(enableDnsCertCommand).Execute()

Enable DNS/Cert automation

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
	enableDnsCertCommand := *openapiclient.NewEnableDnsCertCommand(int32(123), NullableInt32(123), "CertIssuerType_example", "AcmeEmail_example", "Domain_example") // EnableDnsCertCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DnsCertAPI.DnsCertEnable(context.Background()).EnableDnsCertCommand(enableDnsCertCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsCertAPI.DnsCertEnable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsCertEnableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enableDnsCertCommand** | [**EnableDnsCertCommand**](EnableDnsCertCommand.md) |  | 

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


## DnsCertExternalDnsDisable

> DnsCertExternalDnsDisable(ctx).DisableExternalDnsCommand(disableExternalDnsCommand).Execute()

Disable external-dns for a project

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
	disableExternalDnsCommand := *openapiclient.NewDisableExternalDnsCommand(int32(123)) // DisableExternalDnsCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DnsCertAPI.DnsCertExternalDnsDisable(context.Background()).DisableExternalDnsCommand(disableExternalDnsCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsCertAPI.DnsCertExternalDnsDisable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsCertExternalDnsDisableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **disableExternalDnsCommand** | [**DisableExternalDnsCommand**](DisableExternalDnsCommand.md) |  | 

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


## DnsCertExternalDnsEnable

> DnsCertExternalDnsEnable(ctx).EnableExternalDnsCommand(enableExternalDnsCommand).Execute()

Enable external-dns for a project

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
	enableExternalDnsCommand := *openapiclient.NewEnableExternalDnsCommand(int32(123), NullableInt32(123), "Domain_example") // EnableExternalDnsCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DnsCertAPI.DnsCertExternalDnsEnable(context.Background()).EnableExternalDnsCommand(enableExternalDnsCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsCertAPI.DnsCertExternalDnsEnable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsCertExternalDnsEnableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enableExternalDnsCommand** | [**EnableExternalDnsCommand**](EnableExternalDnsCommand.md) |  | 

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


## DnsCertStatus

> GetDnsCertStatusResponse DnsCertStatus(ctx, projectId).Execute()

Get DNS/Cert status for project

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
	resp, r, err := apiClient.DnsCertAPI.DnsCertStatus(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsCertAPI.DnsCertStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DnsCertStatus`: GetDnsCertStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `DnsCertAPI.DnsCertStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsCertStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDnsCertStatusResponse**](GetDnsCertStatusResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsCertSync

> DnsCertSync(ctx).DnsCertSyncCommand(dnsCertSyncCommand).Execute()

Sync DNS/Cert to cluster

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
	dnsCertSyncCommand := *openapiclient.NewDnsCertSyncCommand(int32(123)) // DnsCertSyncCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DnsCertAPI.DnsCertSync(context.Background()).DnsCertSyncCommand(dnsCertSyncCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsCertAPI.DnsCertSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsCertSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dnsCertSyncCommand** | [**DnsCertSyncCommand**](DnsCertSyncCommand.md) |  | 

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


## DnsCertTrustManagerDisable

> DnsCertTrustManagerDisable(ctx).DisableTrustManagerCommand(disableTrustManagerCommand).Execute()

Disable trust-manager for a project

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
	disableTrustManagerCommand := *openapiclient.NewDisableTrustManagerCommand(int32(123)) // DisableTrustManagerCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DnsCertAPI.DnsCertTrustManagerDisable(context.Background()).DisableTrustManagerCommand(disableTrustManagerCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsCertAPI.DnsCertTrustManagerDisable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsCertTrustManagerDisableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **disableTrustManagerCommand** | [**DisableTrustManagerCommand**](DisableTrustManagerCommand.md) |  | 

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


## DnsCertTrustManagerEnable

> DnsCertTrustManagerEnable(ctx).EnableTrustManagerCommand(enableTrustManagerCommand).Execute()

Enable trust-manager for a project

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
	enableTrustManagerCommand := *openapiclient.NewEnableTrustManagerCommand(int32(123)) // EnableTrustManagerCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DnsCertAPI.DnsCertTrustManagerEnable(context.Background()).EnableTrustManagerCommand(enableTrustManagerCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsCertAPI.DnsCertTrustManagerEnable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsCertTrustManagerEnableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enableTrustManagerCommand** | [**EnableTrustManagerCommand**](EnableTrustManagerCommand.md) |  | 

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


## DnsCertUnassignCertificateProfile

> DnsCertUnassignCertificateProfile(ctx).UnassignCertificateProfileCommand(unassignCertificateProfileCommand).Execute()

Remove the signing certificate profile from a CA-issuer project

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
	unassignCertificateProfileCommand := *openapiclient.NewUnassignCertificateProfileCommand(int32(123)) // UnassignCertificateProfileCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DnsCertAPI.DnsCertUnassignCertificateProfile(context.Background()).UnassignCertificateProfileCommand(unassignCertificateProfileCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsCertAPI.DnsCertUnassignCertificateProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsCertUnassignCertificateProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unassignCertificateProfileCommand** | [**UnassignCertificateProfileCommand**](UnassignCertificateProfileCommand.md) |  | 

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


## DnsCertValidate

> DnsCertValidate(ctx).ValidateDnsCertCommand(validateDnsCertCommand).Execute()

Validate DNS/Cert enablement inputs (provider credentials, signing CA reference, issuer rules)

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
	r, err := apiClient.DnsCertAPI.DnsCertValidate(context.Background()).ValidateDnsCertCommand(validateDnsCertCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsCertAPI.DnsCertValidate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsCertValidateRequest struct via the builder pattern


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

