# \GlobalConfigurationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GlobalConfigurationsCreateAiBackend**](GlobalConfigurationAPI.md#GlobalConfigurationsCreateAiBackend) | **Post** /api/v1/global-configurations/ai-backend/create | Create AI backend
[**GlobalConfigurationsCreateAirgapConfiguration**](GlobalConfigurationAPI.md#GlobalConfigurationsCreateAirgapConfiguration) | **Post** /api/v1/global-configurations/airgap/create | Create airgap global configuration
[**GlobalConfigurationsCreateDnsCertConfiguration**](GlobalConfigurationAPI.md#GlobalConfigurationsCreateDnsCertConfiguration) | **Post** /api/v1/global-configurations/dns-cert/create | Create DNS/Cert global configuration
[**GlobalConfigurationsCreateEmailConfig**](GlobalConfigurationAPI.md#GlobalConfigurationsCreateEmailConfig) | **Post** /api/v1/global-configurations/email-validator-config/create | Create email config
[**GlobalConfigurationsCreateSlackConfiguration**](GlobalConfigurationAPI.md#GlobalConfigurationsCreateSlackConfiguration) | **Post** /api/v1/global-configurations/slack/create | Create slack configuration
[**GlobalConfigurationsDeleteAiBackend**](GlobalConfigurationAPI.md#GlobalConfigurationsDeleteAiBackend) | **Delete** /api/v1/global-configurations/ai-backend/{id} | Remove global ai backend configuration
[**GlobalConfigurationsDeleteAirgapConfiguration**](GlobalConfigurationAPI.md#GlobalConfigurationsDeleteAirgapConfiguration) | **Delete** /api/v1/global-configurations/airgap/{id} | Delete airgap global configuration
[**GlobalConfigurationsDeleteDnsCertConfiguration**](GlobalConfigurationAPI.md#GlobalConfigurationsDeleteDnsCertConfiguration) | **Delete** /api/v1/global-configurations/dns-cert/{id} | Delete DNS/Cert global configuration
[**GlobalConfigurationsDeleteSlackConfiguration**](GlobalConfigurationAPI.md#GlobalConfigurationsDeleteSlackConfiguration) | **Delete** /api/v1/global-configurations/slack/{id} | Remove global slack configuration
[**GlobalConfigurationsListAiBackendConfiguration**](GlobalConfigurationAPI.md#GlobalConfigurationsListAiBackendConfiguration) | **Get** /api/v1/global-configurations/ai-backend/list | Retrieve ai backend configuration list
[**GlobalConfigurationsListAirgapConfiguration**](GlobalConfigurationAPI.md#GlobalConfigurationsListAirgapConfiguration) | **Get** /api/v1/global-configurations/airgap/list | List airgap global configurations
[**GlobalConfigurationsListDnsCertConfiguration**](GlobalConfigurationAPI.md#GlobalConfigurationsListDnsCertConfiguration) | **Get** /api/v1/global-configurations/dns-cert/list | List DNS/Cert global configurations
[**GlobalConfigurationsListSlackConfiguration**](GlobalConfigurationAPI.md#GlobalConfigurationsListSlackConfiguration) | **Get** /api/v1/global-configurations/slack/list | Retrieve slack configuration list
[**GlobalConfigurationsUpdateAirgapConfiguration**](GlobalConfigurationAPI.md#GlobalConfigurationsUpdateAirgapConfiguration) | **Put** /api/v1/global-configurations/airgap/update | Update airgap global configuration
[**GlobalConfigurationsUpdateStatusAiBackend**](GlobalConfigurationAPI.md#GlobalConfigurationsUpdateStatusAiBackend) | **Post** /api/v1/global-configurations/ai-backend/status | Activate/Deactivate global ai-backend configuration
[**GlobalConfigurationsUpdateStatusAirgapConfiguration**](GlobalConfigurationAPI.md#GlobalConfigurationsUpdateStatusAirgapConfiguration) | **Post** /api/v1/global-configurations/airgap/status | Activate/Deactivate airgap global configuration
[**GlobalConfigurationsUpdateStatusDnsCertConfiguration**](GlobalConfigurationAPI.md#GlobalConfigurationsUpdateStatusDnsCertConfiguration) | **Post** /api/v1/global-configurations/dns-cert/status | Activate/Deactivate DNS/Cert global configuration
[**GlobalConfigurationsUpdateStatusSlackConfiguration**](GlobalConfigurationAPI.md#GlobalConfigurationsUpdateStatusSlackConfiguration) | **Post** /api/v1/global-configurations/slack/status | Activate/Deactivate global slack configuration



## GlobalConfigurationsCreateAiBackend

> GlobalConfigurationsCreateAiBackend(ctx).CreateAiBackendsConfigurationCommand(createAiBackendsConfigurationCommand).Execute()

Create AI backend

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
	createAiBackendsConfigurationCommand := *openapiclient.NewCreateAiBackendsConfigurationCommand("Name_example", "Url_example", "Model_example", "ApiKey_example") // CreateAiBackendsConfigurationCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GlobalConfigurationAPI.GlobalConfigurationsCreateAiBackend(context.Background()).CreateAiBackendsConfigurationCommand(createAiBackendsConfigurationCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalConfigurationAPI.GlobalConfigurationsCreateAiBackend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGlobalConfigurationsCreateAiBackendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAiBackendsConfigurationCommand** | [**CreateAiBackendsConfigurationCommand**](CreateAiBackendsConfigurationCommand.md) |  | 

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


## GlobalConfigurationsCreateAirgapConfiguration

> GlobalConfigurationsCreateAirgapConfiguration(ctx).CreateAirgapGlobalConfigurationCommand(createAirgapGlobalConfigurationCommand).Execute()

Create airgap global configuration

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
	createAirgapGlobalConfigurationCommand := *openapiclient.NewCreateAirgapGlobalConfigurationCommand("Name_example", "OciRegistryHost_example", false, "Username_example", "Password_example") // CreateAirgapGlobalConfigurationCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GlobalConfigurationAPI.GlobalConfigurationsCreateAirgapConfiguration(context.Background()).CreateAirgapGlobalConfigurationCommand(createAirgapGlobalConfigurationCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalConfigurationAPI.GlobalConfigurationsCreateAirgapConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGlobalConfigurationsCreateAirgapConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAirgapGlobalConfigurationCommand** | [**CreateAirgapGlobalConfigurationCommand**](CreateAirgapGlobalConfigurationCommand.md) |  | 

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


## GlobalConfigurationsCreateDnsCertConfiguration

> GlobalConfigurationsCreateDnsCertConfiguration(ctx).CreateDnsCertGlobalConfigurationCommand(createDnsCertGlobalConfigurationCommand).Execute()

Create DNS/Cert global configuration

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
	createDnsCertGlobalConfigurationCommand := *openapiclient.NewCreateDnsCertGlobalConfigurationCommand("Name_example", "ExternalDnsChartName_example", "ExternalDnsRepoUrl_example", "ExternalDnsVersion_example", "CertManagerChartName_example", "CertManagerRepoUrl_example", "CertManagerVersion_example", "DefaultCertIssuer_example", "DefaultAcmeServer_example", "SupportedProviders_example") // CreateDnsCertGlobalConfigurationCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GlobalConfigurationAPI.GlobalConfigurationsCreateDnsCertConfiguration(context.Background()).CreateDnsCertGlobalConfigurationCommand(createDnsCertGlobalConfigurationCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalConfigurationAPI.GlobalConfigurationsCreateDnsCertConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGlobalConfigurationsCreateDnsCertConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDnsCertGlobalConfigurationCommand** | [**CreateDnsCertGlobalConfigurationCommand**](CreateDnsCertGlobalConfigurationCommand.md) |  | 

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


## GlobalConfigurationsCreateEmailConfig

> GlobalConfigurationsCreateEmailConfig(ctx).CreateEmailConfigurationCommand(createEmailConfigurationCommand).Execute()

Create email config

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
	createEmailConfigurationCommand := *openapiclient.NewCreateEmailConfigurationCommand("Name_example", "Host_example", "ApiKey_example") // CreateEmailConfigurationCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GlobalConfigurationAPI.GlobalConfigurationsCreateEmailConfig(context.Background()).CreateEmailConfigurationCommand(createEmailConfigurationCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalConfigurationAPI.GlobalConfigurationsCreateEmailConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGlobalConfigurationsCreateEmailConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEmailConfigurationCommand** | [**CreateEmailConfigurationCommand**](CreateEmailConfigurationCommand.md) |  | 

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


## GlobalConfigurationsCreateSlackConfiguration

> GlobalConfigurationsCreateSlackConfiguration(ctx).CreateGlobalSlackConfigurationCommand(createGlobalSlackConfigurationCommand).Execute()

Create slack configuration

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
	createGlobalSlackConfigurationCommand := *openapiclient.NewCreateGlobalSlackConfigurationCommand("Name_example", "Url_example", "UserName_example", "Channel_example") // CreateGlobalSlackConfigurationCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GlobalConfigurationAPI.GlobalConfigurationsCreateSlackConfiguration(context.Background()).CreateGlobalSlackConfigurationCommand(createGlobalSlackConfigurationCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalConfigurationAPI.GlobalConfigurationsCreateSlackConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGlobalConfigurationsCreateSlackConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createGlobalSlackConfigurationCommand** | [**CreateGlobalSlackConfigurationCommand**](CreateGlobalSlackConfigurationCommand.md) |  | 

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


## GlobalConfigurationsDeleteAiBackend

> GlobalConfigurationsDeleteAiBackend(ctx, id).Execute()

Remove global ai backend configuration

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
	r, err := apiClient.GlobalConfigurationAPI.GlobalConfigurationsDeleteAiBackend(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalConfigurationAPI.GlobalConfigurationsDeleteAiBackend``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGlobalConfigurationsDeleteAiBackendRequest struct via the builder pattern


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


## GlobalConfigurationsDeleteAirgapConfiguration

> GlobalConfigurationsDeleteAirgapConfiguration(ctx, id).Execute()

Delete airgap global configuration

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
	r, err := apiClient.GlobalConfigurationAPI.GlobalConfigurationsDeleteAirgapConfiguration(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalConfigurationAPI.GlobalConfigurationsDeleteAirgapConfiguration``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGlobalConfigurationsDeleteAirgapConfigurationRequest struct via the builder pattern


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


## GlobalConfigurationsDeleteDnsCertConfiguration

> GlobalConfigurationsDeleteDnsCertConfiguration(ctx, id).Execute()

Delete DNS/Cert global configuration

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
	r, err := apiClient.GlobalConfigurationAPI.GlobalConfigurationsDeleteDnsCertConfiguration(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalConfigurationAPI.GlobalConfigurationsDeleteDnsCertConfiguration``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGlobalConfigurationsDeleteDnsCertConfigurationRequest struct via the builder pattern


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


## GlobalConfigurationsDeleteSlackConfiguration

> GlobalConfigurationsDeleteSlackConfiguration(ctx, id).Execute()

Remove global slack configuration

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
	r, err := apiClient.GlobalConfigurationAPI.GlobalConfigurationsDeleteSlackConfiguration(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalConfigurationAPI.GlobalConfigurationsDeleteSlackConfiguration``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGlobalConfigurationsDeleteSlackConfigurationRequest struct via the builder pattern


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


## GlobalConfigurationsListAiBackendConfiguration

> GlobalAiBackendConfigurationList GlobalConfigurationsListAiBackendConfiguration(ctx).Limit(limit).Offset(offset).Search(search).Execute()

Retrieve ai backend configuration list

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
	limit := int32(56) // int32 |  (optional) (default to 100)
	offset := int32(56) // int32 |  (optional) (default to 0)
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalConfigurationAPI.GlobalConfigurationsListAiBackendConfiguration(context.Background()).Limit(limit).Offset(offset).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalConfigurationAPI.GlobalConfigurationsListAiBackendConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GlobalConfigurationsListAiBackendConfiguration`: GlobalAiBackendConfigurationList
	fmt.Fprintf(os.Stdout, "Response from `GlobalConfigurationAPI.GlobalConfigurationsListAiBackendConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGlobalConfigurationsListAiBackendConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 100]
 **offset** | **int32** |  | [default to 0]
 **search** | **string** |  | 

### Return type

[**GlobalAiBackendConfigurationList**](GlobalAiBackendConfigurationList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GlobalConfigurationsListAirgapConfiguration

> AirgapGlobalConfigurationList GlobalConfigurationsListAirgapConfiguration(ctx).Limit(limit).Offset(offset).Search(search).Execute()

List airgap global configurations

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
	limit := int32(56) // int32 |  (optional) (default to 100)
	offset := int32(56) // int32 |  (optional) (default to 0)
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalConfigurationAPI.GlobalConfigurationsListAirgapConfiguration(context.Background()).Limit(limit).Offset(offset).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalConfigurationAPI.GlobalConfigurationsListAirgapConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GlobalConfigurationsListAirgapConfiguration`: AirgapGlobalConfigurationList
	fmt.Fprintf(os.Stdout, "Response from `GlobalConfigurationAPI.GlobalConfigurationsListAirgapConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGlobalConfigurationsListAirgapConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 100]
 **offset** | **int32** |  | [default to 0]
 **search** | **string** |  | 

### Return type

[**AirgapGlobalConfigurationList**](AirgapGlobalConfigurationList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GlobalConfigurationsListDnsCertConfiguration

> DnsCertGlobalConfigurationList GlobalConfigurationsListDnsCertConfiguration(ctx).Limit(limit).Offset(offset).Search(search).Execute()

List DNS/Cert global configurations

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
	limit := int32(56) // int32 |  (optional) (default to 100)
	offset := int32(56) // int32 |  (optional) (default to 0)
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalConfigurationAPI.GlobalConfigurationsListDnsCertConfiguration(context.Background()).Limit(limit).Offset(offset).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalConfigurationAPI.GlobalConfigurationsListDnsCertConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GlobalConfigurationsListDnsCertConfiguration`: DnsCertGlobalConfigurationList
	fmt.Fprintf(os.Stdout, "Response from `GlobalConfigurationAPI.GlobalConfigurationsListDnsCertConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGlobalConfigurationsListDnsCertConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 100]
 **offset** | **int32** |  | [default to 0]
 **search** | **string** |  | 

### Return type

[**DnsCertGlobalConfigurationList**](DnsCertGlobalConfigurationList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GlobalConfigurationsListSlackConfiguration

> GlobalSlackConfigurationList GlobalConfigurationsListSlackConfiguration(ctx).Limit(limit).Offset(offset).Search(search).Execute()

Retrieve slack configuration list

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
	limit := int32(56) // int32 |  (optional) (default to 100)
	offset := int32(56) // int32 |  (optional) (default to 0)
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalConfigurationAPI.GlobalConfigurationsListSlackConfiguration(context.Background()).Limit(limit).Offset(offset).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalConfigurationAPI.GlobalConfigurationsListSlackConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GlobalConfigurationsListSlackConfiguration`: GlobalSlackConfigurationList
	fmt.Fprintf(os.Stdout, "Response from `GlobalConfigurationAPI.GlobalConfigurationsListSlackConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGlobalConfigurationsListSlackConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 100]
 **offset** | **int32** |  | [default to 0]
 **search** | **string** |  | 

### Return type

[**GlobalSlackConfigurationList**](GlobalSlackConfigurationList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GlobalConfigurationsUpdateAirgapConfiguration

> GlobalConfigurationsUpdateAirgapConfiguration(ctx).UpdateAirgapGlobalConfigurationCommand(updateAirgapGlobalConfigurationCommand).Execute()

Update airgap global configuration

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
	updateAirgapGlobalConfigurationCommand := *openapiclient.NewUpdateAirgapGlobalConfigurationCommand(int32(123), "Name_example", "OciRegistryHost_example", false, "Username_example", "Password_example") // UpdateAirgapGlobalConfigurationCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GlobalConfigurationAPI.GlobalConfigurationsUpdateAirgapConfiguration(context.Background()).UpdateAirgapGlobalConfigurationCommand(updateAirgapGlobalConfigurationCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalConfigurationAPI.GlobalConfigurationsUpdateAirgapConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGlobalConfigurationsUpdateAirgapConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateAirgapGlobalConfigurationCommand** | [**UpdateAirgapGlobalConfigurationCommand**](UpdateAirgapGlobalConfigurationCommand.md) |  | 

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


## GlobalConfigurationsUpdateStatusAiBackend

> GlobalConfigurationsUpdateStatusAiBackend(ctx).GlobalAiBackendConfigurationStatusManagementCommand(globalAiBackendConfigurationStatusManagementCommand).Execute()

Activate/Deactivate global ai-backend configuration

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
	globalAiBackendConfigurationStatusManagementCommand := *openapiclient.NewGlobalAiBackendConfigurationStatusManagementCommand(int32(123), "Mode_example") // GlobalAiBackendConfigurationStatusManagementCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GlobalConfigurationAPI.GlobalConfigurationsUpdateStatusAiBackend(context.Background()).GlobalAiBackendConfigurationStatusManagementCommand(globalAiBackendConfigurationStatusManagementCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalConfigurationAPI.GlobalConfigurationsUpdateStatusAiBackend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGlobalConfigurationsUpdateStatusAiBackendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **globalAiBackendConfigurationStatusManagementCommand** | [**GlobalAiBackendConfigurationStatusManagementCommand**](GlobalAiBackendConfigurationStatusManagementCommand.md) |  | 

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


## GlobalConfigurationsUpdateStatusAirgapConfiguration

> GlobalConfigurationsUpdateStatusAirgapConfiguration(ctx).AirgapGlobalConfigurationStatusManagementCommand(airgapGlobalConfigurationStatusManagementCommand).Execute()

Activate/Deactivate airgap global configuration

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
	airgapGlobalConfigurationStatusManagementCommand := *openapiclient.NewAirgapGlobalConfigurationStatusManagementCommand(int32(123), "Mode_example") // AirgapGlobalConfigurationStatusManagementCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GlobalConfigurationAPI.GlobalConfigurationsUpdateStatusAirgapConfiguration(context.Background()).AirgapGlobalConfigurationStatusManagementCommand(airgapGlobalConfigurationStatusManagementCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalConfigurationAPI.GlobalConfigurationsUpdateStatusAirgapConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGlobalConfigurationsUpdateStatusAirgapConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **airgapGlobalConfigurationStatusManagementCommand** | [**AirgapGlobalConfigurationStatusManagementCommand**](AirgapGlobalConfigurationStatusManagementCommand.md) |  | 

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


## GlobalConfigurationsUpdateStatusDnsCertConfiguration

> GlobalConfigurationsUpdateStatusDnsCertConfiguration(ctx).DnsCertGlobalConfigurationStatusCommand(dnsCertGlobalConfigurationStatusCommand).Execute()

Activate/Deactivate DNS/Cert global configuration

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
	dnsCertGlobalConfigurationStatusCommand := *openapiclient.NewDnsCertGlobalConfigurationStatusCommand(int32(123), "Mode_example") // DnsCertGlobalConfigurationStatusCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GlobalConfigurationAPI.GlobalConfigurationsUpdateStatusDnsCertConfiguration(context.Background()).DnsCertGlobalConfigurationStatusCommand(dnsCertGlobalConfigurationStatusCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalConfigurationAPI.GlobalConfigurationsUpdateStatusDnsCertConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGlobalConfigurationsUpdateStatusDnsCertConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dnsCertGlobalConfigurationStatusCommand** | [**DnsCertGlobalConfigurationStatusCommand**](DnsCertGlobalConfigurationStatusCommand.md) |  | 

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


## GlobalConfigurationsUpdateStatusSlackConfiguration

> GlobalConfigurationsUpdateStatusSlackConfiguration(ctx).GlobalSlackConfigurationStatusManagementCommand(globalSlackConfigurationStatusManagementCommand).Execute()

Activate/Deactivate global slack configuration

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
	globalSlackConfigurationStatusManagementCommand := *openapiclient.NewGlobalSlackConfigurationStatusManagementCommand(int32(123), "Mode_example") // GlobalSlackConfigurationStatusManagementCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GlobalConfigurationAPI.GlobalConfigurationsUpdateStatusSlackConfiguration(context.Background()).GlobalSlackConfigurationStatusManagementCommand(globalSlackConfigurationStatusManagementCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalConfigurationAPI.GlobalConfigurationsUpdateStatusSlackConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGlobalConfigurationsUpdateStatusSlackConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **globalSlackConfigurationStatusManagementCommand** | [**GlobalSlackConfigurationStatusManagementCommand**](GlobalSlackConfigurationStatusManagementCommand.md) |  | 

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

