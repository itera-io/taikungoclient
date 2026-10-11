# \CheckerAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckerArtifact**](CheckerAPI.md#CheckerArtifact) | **Post** /api/v1/checker/artifact | Check artifact url
[**CheckerAws**](CheckerAPI.md#CheckerAws) | **Post** /api/v1/checker/aws | Check aws credential
[**CheckerAzure**](CheckerAPI.md#CheckerAzure) | **Post** /api/v1/checker/azure | Check azure credentials
[**CheckerAzureQuota**](CheckerAPI.md#CheckerAzureQuota) | **Post** /api/v1/checker/azure/quota/cpu | Check azure cpu quota
[**CheckerCidr**](CheckerAPI.md#CheckerCidr) | **Post** /api/v1/checker/cidr | Check valid cidr format
[**CheckerCron**](CheckerAPI.md#CheckerCron) | **Post** /api/v1/checker/cron | Check valid cron job format
[**CheckerDns**](CheckerAPI.md#CheckerDns) | **Post** /api/v1/checker/dns | Check valid dns format
[**CheckerGoogle**](CheckerAPI.md#CheckerGoogle) | **Post** /api/v1/checker/google | Check google credentials
[**CheckerHelm**](CheckerAPI.md#CheckerHelm) | **Post** /api/v1/checker/helm | Check helm credential
[**CheckerImportClusterKubeConfig**](CheckerAPI.md#CheckerImportClusterKubeConfig) | **Post** /api/v1/checker/import-cluster-kube-config | Check kube config
[**CheckerKubeConfig**](CheckerAPI.md#CheckerKubeConfig) | **Post** /api/v1/checker/kube-config | Check kube config
[**CheckerNode**](CheckerAPI.md#CheckerNode) | **Post** /api/v1/checker/node | Duplicate server name checker
[**CheckerNtp**](CheckerAPI.md#CheckerNtp) | **Post** /api/v1/checker/ntp | Check valid ntp format
[**CheckerOpenAi**](CheckerAPI.md#CheckerOpenAi) | **Post** /api/v1/checker/openai | Check open-ai token
[**CheckerOpenstack**](CheckerAPI.md#CheckerOpenstack) | **Post** /api/v1/checker/openstack | Check openstack credential
[**CheckerOpenstackTaikunImage**](CheckerAPI.md#CheckerOpenstackTaikunImage) | **Post** /api/v1/checker/openstack-image/{id} | Check openstack taikun image
[**CheckerOpenstackTaikunLbImage**](CheckerAPI.md#CheckerOpenstackTaikunLbImage) | **Post** /api/v1/checker/taikun-lb-image/{id} | Check openstack taikun lb image
[**CheckerOrganization**](CheckerAPI.md#CheckerOrganization) | **Post** /api/v1/checker/organization | Check duplicate org name
[**CheckerProjectApp**](CheckerAPI.md#CheckerProjectApp) | **Post** /api/v1/checker/project-app | Check duplicate project app
[**CheckerPrometheus**](CheckerAPI.md#CheckerPrometheus) | **Post** /api/v1/checker/prometheus | Check prometheus credentials
[**CheckerS3**](CheckerAPI.md#CheckerS3) | **Post** /api/v1/checker/s3 | Check s3 credential
[**CheckerSsh**](CheckerAPI.md#CheckerSsh) | **Post** /api/v1/checker/ssh | Check valid ssh key format
[**CheckerTrustedRegistry**](CheckerAPI.md#CheckerTrustedRegistry) | **Post** /api/v1/checker/trusted-registry | Check trusted registry
[**CheckerUser**](CheckerAPI.md#CheckerUser) | **Post** /api/v1/checker/user | Check duplicate username
[**CheckerYaml**](CheckerAPI.md#CheckerYaml) | **Post** /api/v1/checker/yaml | Check yaml file



## CheckerArtifact

> CheckerArtifact(ctx).ArtifactUrlCheckerCommand(artifactUrlCheckerCommand).Execute()

Check artifact url

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
	artifactUrlCheckerCommand := *openapiclient.NewArtifactUrlCheckerCommand("Url_example", "Name_example", NullableInt32(123)) // ArtifactUrlCheckerCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CheckerAPI.CheckerArtifact(context.Background()).ArtifactUrlCheckerCommand(artifactUrlCheckerCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckerAPI.CheckerArtifact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckerArtifactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **artifactUrlCheckerCommand** | [**ArtifactUrlCheckerCommand**](ArtifactUrlCheckerCommand.md) |  | 

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


## CheckerAws

> CheckerAws(ctx).CheckAwsCommand(checkAwsCommand).Execute()

Check aws credential

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
	checkAwsCommand := *openapiclient.NewCheckAwsCommand("AwsAccessKeyId_example", "AwsSecretAccessKey_example", "RoleArn_example", "ExternalId_example") // CheckAwsCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CheckerAPI.CheckerAws(context.Background()).CheckAwsCommand(checkAwsCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckerAPI.CheckerAws``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckerAwsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkAwsCommand** | [**CheckAwsCommand**](CheckAwsCommand.md) |  | 

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


## CheckerAzure

> CheckerAzure(ctx).CheckAzureCommand(checkAzureCommand).Execute()

Check azure credentials

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
	checkAzureCommand := *openapiclient.NewCheckAzureCommand("AzureClientId_example", "AzureClientSecret_example", "AzureTenantId_example") // CheckAzureCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CheckerAPI.CheckerAzure(context.Background()).CheckAzureCommand(checkAzureCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckerAPI.CheckerAzure``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckerAzureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkAzureCommand** | [**CheckAzureCommand**](CheckAzureCommand.md) |  | 

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


## CheckerAzureQuota

> CheckerAzureQuota(ctx).CheckAzureCpuQuotaCommand(checkAzureCpuQuotaCommand).Execute()

Check azure cpu quota

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
	checkAzureCpuQuotaCommand := *openapiclient.NewCheckAzureCpuQuotaCommand(int32(123)) // CheckAzureCpuQuotaCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CheckerAPI.CheckerAzureQuota(context.Background()).CheckAzureCpuQuotaCommand(checkAzureCpuQuotaCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckerAPI.CheckerAzureQuota``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckerAzureQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkAzureCpuQuotaCommand** | [**CheckAzureCpuQuotaCommand**](CheckAzureCpuQuotaCommand.md) |  | 

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


## CheckerCidr

> CheckerCidr(ctx).CidrCommand(cidrCommand).Execute()

Check valid cidr format

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
	cidrCommand := *openapiclient.NewCidrCommand("Cidr_example") // CidrCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CheckerAPI.CheckerCidr(context.Background()).CidrCommand(cidrCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckerAPI.CheckerCidr``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckerCidrRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cidrCommand** | [**CidrCommand**](CidrCommand.md) |  | 

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


## CheckerCron

> CheckerCron(ctx).CronJobCommand(cronJobCommand).Execute()

Check valid cron job format

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
	cronJobCommand := *openapiclient.NewCronJobCommand("CronPeriod_example") // CronJobCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CheckerAPI.CheckerCron(context.Background()).CronJobCommand(cronJobCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckerAPI.CheckerCron``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckerCronRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cronJobCommand** | [**CronJobCommand**](CronJobCommand.md) |  | 

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


## CheckerDns

> CheckerDns(ctx).DnsCommand(dnsCommand).Execute()

Check valid dns format

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
	dnsCommand := *openapiclient.NewDnsCommand("Address_example") // DnsCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CheckerAPI.CheckerDns(context.Background()).DnsCommand(dnsCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckerAPI.CheckerDns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckerDnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dnsCommand** | [**DnsCommand**](DnsCommand.md) |  | 

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


## CheckerGoogle

> bool CheckerGoogle(ctx).Config(config).Execute()

Check google credentials

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
	resp, r, err := apiClient.CheckerAPI.CheckerGoogle(context.Background()).Config(config).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckerAPI.CheckerGoogle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckerGoogle`: bool
	fmt.Fprintf(os.Stdout, "Response from `CheckerAPI.CheckerGoogle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckerGoogleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **config** | ***os.File** |  | 

### Return type

**bool**

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckerHelm

> CheckerHelm(ctx).HelmCredentialCommand(helmCredentialCommand).Execute()

Check helm credential

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
	helmCredentialCommand := *openapiclient.NewHelmCredentialCommand("Url_example", "Username_example", "Password_example") // HelmCredentialCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CheckerAPI.CheckerHelm(context.Background()).HelmCredentialCommand(helmCredentialCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckerAPI.CheckerHelm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckerHelmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **helmCredentialCommand** | [**HelmCredentialCommand**](HelmCredentialCommand.md) |  | 

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


## CheckerImportClusterKubeConfig

> CheckerImportClusterKubeConfig(ctx).Config(config).CloudId(cloudId).Region(region).ImportClusterName(importClusterName).ResourceGroupName(resourceGroupName).ImportType(importType).Execute()

Check kube config

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
	cloudId := int32(56) // int32 |  (optional)
	region := "region_example" // string |  (optional)
	importClusterName := "importClusterName_example" // string |  (optional)
	resourceGroupName := "resourceGroupName_example" // string |  (optional)
	importType := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CheckerAPI.CheckerImportClusterKubeConfig(context.Background()).Config(config).CloudId(cloudId).Region(region).ImportClusterName(importClusterName).ResourceGroupName(resourceGroupName).ImportType(importType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckerAPI.CheckerImportClusterKubeConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckerImportClusterKubeConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **config** | ***os.File** |  | 
 **cloudId** | **int32** |  | 
 **region** | **string** |  | 
 **importClusterName** | **string** |  | 
 **resourceGroupName** | **string** |  | 
 **importType** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckerKubeConfig

> bool CheckerKubeConfig(ctx).Config(config).Execute()

Check kube config

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
	resp, r, err := apiClient.CheckerAPI.CheckerKubeConfig(context.Background()).Config(config).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckerAPI.CheckerKubeConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckerKubeConfig`: bool
	fmt.Fprintf(os.Stdout, "Response from `CheckerAPI.CheckerKubeConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckerKubeConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **config** | ***os.File** |  | 

### Return type

**bool**

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckerNode

> CheckerNode(ctx).NodeCommand(nodeCommand).Execute()

Duplicate server name checker

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
	nodeCommand := *openapiclient.NewNodeCommand(int32(123), "Name_example") // NodeCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CheckerAPI.CheckerNode(context.Background()).NodeCommand(nodeCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckerAPI.CheckerNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckerNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nodeCommand** | [**NodeCommand**](NodeCommand.md) |  | 

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


## CheckerNtp

> CheckerNtp(ctx).NtpCommand(ntpCommand).Execute()

Check valid ntp format

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
	ntpCommand := *openapiclient.NewNtpCommand("Address_example") // NtpCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CheckerAPI.CheckerNtp(context.Background()).NtpCommand(ntpCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckerAPI.CheckerNtp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckerNtpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ntpCommand** | [**NtpCommand**](NtpCommand.md) |  | 

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


## CheckerOpenAi

> CheckerOpenAi(ctx).OpenAiCheckerCommand(openAiCheckerCommand).Execute()

Check open-ai token

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
	openAiCheckerCommand := *openapiclient.NewOpenAiCheckerCommand("Token_example") // OpenAiCheckerCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CheckerAPI.CheckerOpenAi(context.Background()).OpenAiCheckerCommand(openAiCheckerCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckerAPI.CheckerOpenAi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckerOpenAiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **openAiCheckerCommand** | [**OpenAiCheckerCommand**](OpenAiCheckerCommand.md) |  | 

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


## CheckerOpenstack

> CheckOpenstackAvailabilityResponse CheckerOpenstack(ctx).CheckOpenstackCommand(checkOpenstackCommand).Execute()

Check openstack credential

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
	checkOpenstackCommand := *openapiclient.NewCheckOpenstackCommand() // CheckOpenstackCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CheckerAPI.CheckerOpenstack(context.Background()).CheckOpenstackCommand(checkOpenstackCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckerAPI.CheckerOpenstack``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckerOpenstack`: CheckOpenstackAvailabilityResponse
	fmt.Fprintf(os.Stdout, "Response from `CheckerAPI.CheckerOpenstack`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckerOpenstackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkOpenstackCommand** | [**CheckOpenstackCommand**](CheckOpenstackCommand.md) |  | 

### Return type

[**CheckOpenstackAvailabilityResponse**](CheckOpenstackAvailabilityResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckerOpenstackTaikunImage

> CheckerOpenstackTaikunImage(ctx, id).Execute()

Check openstack taikun image

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
	r, err := apiClient.CheckerAPI.CheckerOpenstackTaikunImage(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckerAPI.CheckerOpenstackTaikunImage``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCheckerOpenstackTaikunImageRequest struct via the builder pattern


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


## CheckerOpenstackTaikunLbImage

> CheckerOpenstackTaikunLbImage(ctx, id).Execute()

Check openstack taikun lb image

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
	r, err := apiClient.CheckerAPI.CheckerOpenstackTaikunLbImage(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckerAPI.CheckerOpenstackTaikunLbImage``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCheckerOpenstackTaikunLbImageRequest struct via the builder pattern


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


## CheckerOrganization

> CheckerOrganization(ctx).OrganizationNameCheckerCommand(organizationNameCheckerCommand).Execute()

Check duplicate org name

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
	organizationNameCheckerCommand := *openapiclient.NewOrganizationNameCheckerCommand(NullableInt32(123), "Name_example", "FullName_example") // OrganizationNameCheckerCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CheckerAPI.CheckerOrganization(context.Background()).OrganizationNameCheckerCommand(organizationNameCheckerCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckerAPI.CheckerOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckerOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationNameCheckerCommand** | [**OrganizationNameCheckerCommand**](OrganizationNameCheckerCommand.md) |  | 

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


## CheckerProjectApp

> CheckerProjectApp(ctx).DuplicateProjectAppCheckerCommand(duplicateProjectAppCheckerCommand).Execute()

Check duplicate project app

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
	duplicateProjectAppCheckerCommand := *openapiclient.NewDuplicateProjectAppCheckerCommand(int32(123), "Name_example", "Namespace_example") // DuplicateProjectAppCheckerCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CheckerAPI.CheckerProjectApp(context.Background()).DuplicateProjectAppCheckerCommand(duplicateProjectAppCheckerCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckerAPI.CheckerProjectApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckerProjectAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **duplicateProjectAppCheckerCommand** | [**DuplicateProjectAppCheckerCommand**](DuplicateProjectAppCheckerCommand.md) |  | 

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


## CheckerPrometheus

> CheckerPrometheus(ctx).CheckPrometheusCommand(checkPrometheusCommand).Execute()

Check prometheus credentials

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
	checkPrometheusCommand := *openapiclient.NewCheckPrometheusCommand("Password_example", "UserName_example", "Url_example") // CheckPrometheusCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CheckerAPI.CheckerPrometheus(context.Background()).CheckPrometheusCommand(checkPrometheusCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckerAPI.CheckerPrometheus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckerPrometheusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkPrometheusCommand** | [**CheckPrometheusCommand**](CheckPrometheusCommand.md) |  | 

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


## CheckerS3

> CheckerS3(ctx).CheckS3Command(checkS3Command).Execute()

Check s3 credential

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
	checkS3Command := *openapiclient.NewCheckS3Command("S3AccessKeyId_example", "S3SecretKey_example", "S3Endpoint_example", "S3Region_example") // CheckS3Command |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CheckerAPI.CheckerS3(context.Background()).CheckS3Command(checkS3Command).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckerAPI.CheckerS3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckerS3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkS3Command** | [**CheckS3Command**](CheckS3Command.md) |  | 

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


## CheckerSsh

> CheckerSsh(ctx).SshKeyCommand(sshKeyCommand).Execute()

Check valid ssh key format

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
	sshKeyCommand := *openapiclient.NewSshKeyCommand("SshPublicKey_example") // SshKeyCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CheckerAPI.CheckerSsh(context.Background()).SshKeyCommand(sshKeyCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckerAPI.CheckerSsh``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckerSshRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sshKeyCommand** | [**SshKeyCommand**](SshKeyCommand.md) |  | 

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


## CheckerTrustedRegistry

> CheckerTrustedRegistry(ctx).CheckTrustedRegistryCommand(checkTrustedRegistryCommand).Execute()

Check trusted registry

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
	checkTrustedRegistryCommand := *openapiclient.NewCheckTrustedRegistryCommand("Registry_example") // CheckTrustedRegistryCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CheckerAPI.CheckerTrustedRegistry(context.Background()).CheckTrustedRegistryCommand(checkTrustedRegistryCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckerAPI.CheckerTrustedRegistry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckerTrustedRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkTrustedRegistryCommand** | [**CheckTrustedRegistryCommand**](CheckTrustedRegistryCommand.md) |  | 

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


## CheckerUser

> CheckerUser(ctx).UserExistCommand(userExistCommand).Execute()

Check duplicate username

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
	userExistCommand := *openapiclient.NewUserExistCommand("Email_example", "UserName_example", NullableInt32(123)) // UserExistCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CheckerAPI.CheckerUser(context.Background()).UserExistCommand(userExistCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckerAPI.CheckerUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckerUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userExistCommand** | [**UserExistCommand**](UserExistCommand.md) |  | 

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


## CheckerYaml

> CheckerYaml(ctx).YamlValidatorCommand(yamlValidatorCommand).Execute()

Check yaml file

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
	yamlValidatorCommand := *openapiclient.NewYamlValidatorCommand("Yaml_example") // YamlValidatorCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CheckerAPI.CheckerYaml(context.Background()).YamlValidatorCommand(yamlValidatorCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckerAPI.CheckerYaml``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckerYamlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **yamlValidatorCommand** | [**YamlValidatorCommand**](YamlValidatorCommand.md) |  | 

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

