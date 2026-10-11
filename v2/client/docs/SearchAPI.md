# \SearchAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchAccessProfiles**](SearchAPI.md#SearchAccessProfiles) | **Post** /api/v1/search/access-profiles | Global search for access-profiles
[**SearchBackupCredentials**](SearchAPI.md#SearchBackupCredentials) | **Post** /api/v1/search/backup-credentials | Global search for backup-credentials
[**SearchBillingCredentials**](SearchAPI.md#SearchBillingCredentials) | **Post** /api/v1/search/billing-credentials | Global search for billing-credentials
[**SearchCloudCredentials**](SearchAPI.md#SearchCloudCredentials) | **Post** /api/v1/search/cloud-credentials | Global search for cloud-credentials
[**SearchConfigMaps**](SearchAPI.md#SearchConfigMaps) | **Post** /api/v1/search/config-maps | Global search for config-maps
[**SearchCrds**](SearchAPI.md#SearchCrds) | **Post** /api/v1/search/crds | Global search for crds
[**SearchCronJobs**](SearchAPI.md#SearchCronJobs) | **Post** /api/v1/search/cronjobs | Global search for jobs
[**SearchDaemonSets**](SearchAPI.md#SearchDaemonSets) | **Post** /api/v1/search/daemon-sets | Global search for daemon-sets
[**SearchDeployments**](SearchAPI.md#SearchDeployments) | **Post** /api/v1/search/deployments | Global search for deployments
[**SearchHelmReleases**](SearchAPI.md#SearchHelmReleases) | **Post** /api/v1/search/helm-releases | Global search for helm-releases
[**SearchIngress**](SearchAPI.md#SearchIngress) | **Post** /api/v1/search/ingress | Global search for ingress
[**SearchJobs**](SearchAPI.md#SearchJobs) | **Post** /api/v1/search/jobs | Global search for jobs
[**SearchKubernetesProfiles**](SearchAPI.md#SearchKubernetesProfiles) | **Post** /api/v1/search/kubernetes-profiles | Global search for kubernetes-profiles
[**SearchNamespaces**](SearchAPI.md#SearchNamespaces) | **Post** /api/v1/search/namespaces | Global search for namespaces
[**SearchNetworkPolicies**](SearchAPI.md#SearchNetworkPolicies) | **Post** /api/v1/search/network-policies | Global search for namespaces
[**SearchNodes**](SearchAPI.md#SearchNodes) | **Post** /api/v1/search/nodes | Global search for nodes
[**SearchOrganizations**](SearchAPI.md#SearchOrganizations) | **Post** /api/v1/search/organizations | Global search for organizations
[**SearchPdb**](SearchAPI.md#SearchPdb) | **Post** /api/v1/search/pdb | Global search for pdb
[**SearchPods**](SearchAPI.md#SearchPods) | **Post** /api/v1/search/pods | Global search for pods
[**SearchProjects**](SearchAPI.md#SearchProjects) | **Post** /api/v1/search/projects | Global search for projects
[**SearchPrometheusRules**](SearchAPI.md#SearchPrometheusRules) | **Post** /api/v1/search/prometheus-rules | Global search for prometheus-rules
[**SearchPvcs**](SearchAPI.md#SearchPvcs) | **Post** /api/v1/search/pvcs | Global search for pvcs
[**SearchSecrets**](SearchAPI.md#SearchSecrets) | **Post** /api/v1/search/secrets | Global search for secrets
[**SearchServers**](SearchAPI.md#SearchServers) | **Post** /api/v1/search/servers | Global search for servers
[**SearchServices**](SearchAPI.md#SearchServices) | **Post** /api/v1/search/services | Global search for services
[**SearchStandAloneProfiles**](SearchAPI.md#SearchStandAloneProfiles) | **Post** /api/v1/search/stand-alone-profiles | Global search for stand-alone-profiles
[**SearchStorageClasses**](SearchAPI.md#SearchStorageClasses) | **Post** /api/v1/search/storage-classes | Global search for storage classes.
[**SearchSts**](SearchAPI.md#SearchSts) | **Post** /api/v1/search/sts | Global search for sts
[**SearchUsers**](SearchAPI.md#SearchUsers) | **Post** /api/v1/search/users | Global search for users



## SearchAccessProfiles

> AccessProfilesSearchList SearchAccessProfiles(ctx).AccessProfilesSearchCommand(accessProfilesSearchCommand).Execute()

Global search for access-profiles

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
	accessProfilesSearchCommand := *openapiclient.NewAccessProfilesSearchCommand(NullableInt32(123), NullableInt32(123), "SearchTerm_example") // AccessProfilesSearchCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchAccessProfiles(context.Background()).AccessProfilesSearchCommand(accessProfilesSearchCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchAccessProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchAccessProfiles`: AccessProfilesSearchList
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchAccessProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchAccessProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessProfilesSearchCommand** | [**AccessProfilesSearchCommand**](AccessProfilesSearchCommand.md) |  | 

### Return type

[**AccessProfilesSearchList**](AccessProfilesSearchList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchBackupCredentials

> BackupCredentialsSearchList SearchBackupCredentials(ctx).BackupCredentialsSearchCommand(backupCredentialsSearchCommand).Execute()

Global search for backup-credentials

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
	backupCredentialsSearchCommand := *openapiclient.NewBackupCredentialsSearchCommand(NullableInt32(123), NullableInt32(123), "SearchTerm_example") // BackupCredentialsSearchCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchBackupCredentials(context.Background()).BackupCredentialsSearchCommand(backupCredentialsSearchCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchBackupCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchBackupCredentials`: BackupCredentialsSearchList
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchBackupCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchBackupCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **backupCredentialsSearchCommand** | [**BackupCredentialsSearchCommand**](BackupCredentialsSearchCommand.md) |  | 

### Return type

[**BackupCredentialsSearchList**](BackupCredentialsSearchList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchBillingCredentials

> BillingCredentialsSearchList SearchBillingCredentials(ctx).BillingCredentialsSearchCommand(billingCredentialsSearchCommand).Execute()

Global search for billing-credentials

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
	billingCredentialsSearchCommand := *openapiclient.NewBillingCredentialsSearchCommand(NullableInt32(123), NullableInt32(123), "SearchTerm_example") // BillingCredentialsSearchCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchBillingCredentials(context.Background()).BillingCredentialsSearchCommand(billingCredentialsSearchCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchBillingCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchBillingCredentials`: BillingCredentialsSearchList
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchBillingCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchBillingCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **billingCredentialsSearchCommand** | [**BillingCredentialsSearchCommand**](BillingCredentialsSearchCommand.md) |  | 

### Return type

[**BillingCredentialsSearchList**](BillingCredentialsSearchList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCloudCredentials

> CloudCredentialsSearchList SearchCloudCredentials(ctx).CloudCredentialsSearchCommand(cloudCredentialsSearchCommand).Execute()

Global search for cloud-credentials

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
	cloudCredentialsSearchCommand := *openapiclient.NewCloudCredentialsSearchCommand(NullableInt32(123), NullableInt32(123), "SearchTerm_example") // CloudCredentialsSearchCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchCloudCredentials(context.Background()).CloudCredentialsSearchCommand(cloudCredentialsSearchCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchCloudCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchCloudCredentials`: CloudCredentialsSearchList
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchCloudCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchCloudCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudCredentialsSearchCommand** | [**CloudCredentialsSearchCommand**](CloudCredentialsSearchCommand.md) |  | 

### Return type

[**CloudCredentialsSearchList**](CloudCredentialsSearchList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchConfigMaps

> ConfigMapSearchList SearchConfigMaps(ctx).ConfigMapSearchCommand(configMapSearchCommand).Execute()

Global search for config-maps

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
	configMapSearchCommand := *openapiclient.NewConfigMapSearchCommand(NullableInt32(123), NullableInt32(123), "SearchTerm_example", false) // ConfigMapSearchCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchConfigMaps(context.Background()).ConfigMapSearchCommand(configMapSearchCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchConfigMaps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchConfigMaps`: ConfigMapSearchList
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchConfigMaps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchConfigMapsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configMapSearchCommand** | [**ConfigMapSearchCommand**](ConfigMapSearchCommand.md) |  | 

### Return type

[**ConfigMapSearchList**](ConfigMapSearchList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCrds

> CrdSearchList SearchCrds(ctx).CrdSearchCommand(crdSearchCommand).Execute()

Global search for crds

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
	crdSearchCommand := *openapiclient.NewCrdSearchCommand(NullableInt32(123), NullableInt32(123), "SearchTerm_example", false) // CrdSearchCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchCrds(context.Background()).CrdSearchCommand(crdSearchCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchCrds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchCrds`: CrdSearchList
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchCrds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchCrdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **crdSearchCommand** | [**CrdSearchCommand**](CrdSearchCommand.md) |  | 

### Return type

[**CrdSearchList**](CrdSearchList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCronJobs

> CronjobsSearchList SearchCronJobs(ctx).CronjobsSearchCommand(cronjobsSearchCommand).Execute()

Global search for jobs

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
	cronjobsSearchCommand := *openapiclient.NewCronjobsSearchCommand(NullableInt32(123), NullableInt32(123), "SearchTerm_example", false) // CronjobsSearchCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchCronJobs(context.Background()).CronjobsSearchCommand(cronjobsSearchCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchCronJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchCronJobs`: CronjobsSearchList
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchCronJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchCronJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cronjobsSearchCommand** | [**CronjobsSearchCommand**](CronjobsSearchCommand.md) |  | 

### Return type

[**CronjobsSearchList**](CronjobsSearchList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchDaemonSets

> DaemonSetSearchList SearchDaemonSets(ctx).DaemonSetSearchCommand(daemonSetSearchCommand).Execute()

Global search for daemon-sets

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
	daemonSetSearchCommand := *openapiclient.NewDaemonSetSearchCommand(NullableInt32(123), NullableInt32(123), "SearchTerm_example", false) // DaemonSetSearchCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchDaemonSets(context.Background()).DaemonSetSearchCommand(daemonSetSearchCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchDaemonSets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchDaemonSets`: DaemonSetSearchList
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchDaemonSets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchDaemonSetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **daemonSetSearchCommand** | [**DaemonSetSearchCommand**](DaemonSetSearchCommand.md) |  | 

### Return type

[**DaemonSetSearchList**](DaemonSetSearchList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchDeployments

> DeploymentSearchList SearchDeployments(ctx).DeploymentSearchCommand(deploymentSearchCommand).Execute()

Global search for deployments

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
	deploymentSearchCommand := *openapiclient.NewDeploymentSearchCommand(NullableInt32(123), NullableInt32(123), "SearchTerm_example", false) // DeploymentSearchCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchDeployments(context.Background()).DeploymentSearchCommand(deploymentSearchCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchDeployments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchDeployments`: DeploymentSearchList
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchDeployments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchDeploymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deploymentSearchCommand** | [**DeploymentSearchCommand**](DeploymentSearchCommand.md) |  | 

### Return type

[**DeploymentSearchList**](DeploymentSearchList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchHelmReleases

> HelmReleasesSearchList SearchHelmReleases(ctx).HelmReleasesSearchCommand(helmReleasesSearchCommand).Execute()

Global search for helm-releases

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
	helmReleasesSearchCommand := *openapiclient.NewHelmReleasesSearchCommand(NullableInt32(123), NullableInt32(123), "SearchTerm_example", false) // HelmReleasesSearchCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchHelmReleases(context.Background()).HelmReleasesSearchCommand(helmReleasesSearchCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchHelmReleases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchHelmReleases`: HelmReleasesSearchList
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchHelmReleases`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchHelmReleasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **helmReleasesSearchCommand** | [**HelmReleasesSearchCommand**](HelmReleasesSearchCommand.md) |  | 

### Return type

[**HelmReleasesSearchList**](HelmReleasesSearchList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchIngress

> IngressSearchList SearchIngress(ctx).IngressSearchCommand(ingressSearchCommand).Execute()

Global search for ingress

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
	ingressSearchCommand := *openapiclient.NewIngressSearchCommand(NullableInt32(123), NullableInt32(123), "SearchTerm_example", false) // IngressSearchCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchIngress(context.Background()).IngressSearchCommand(ingressSearchCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchIngress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchIngress`: IngressSearchList
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchIngress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchIngressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ingressSearchCommand** | [**IngressSearchCommand**](IngressSearchCommand.md) |  | 

### Return type

[**IngressSearchList**](IngressSearchList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchJobs

> JobsSearchList SearchJobs(ctx).JobsSearchCommand(jobsSearchCommand).Execute()

Global search for jobs

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
	jobsSearchCommand := *openapiclient.NewJobsSearchCommand(NullableInt32(123), NullableInt32(123), "SearchTerm_example", false) // JobsSearchCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchJobs(context.Background()).JobsSearchCommand(jobsSearchCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchJobs`: JobsSearchList
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobsSearchCommand** | [**JobsSearchCommand**](JobsSearchCommand.md) |  | 

### Return type

[**JobsSearchList**](JobsSearchList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchKubernetesProfiles

> KubernetesProfilesSearchList SearchKubernetesProfiles(ctx).KubernetesProfilesSearchCommand(kubernetesProfilesSearchCommand).Execute()

Global search for kubernetes-profiles

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
	kubernetesProfilesSearchCommand := *openapiclient.NewKubernetesProfilesSearchCommand(NullableInt32(123), NullableInt32(123), "SearchTerm_example") // KubernetesProfilesSearchCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchKubernetesProfiles(context.Background()).KubernetesProfilesSearchCommand(kubernetesProfilesSearchCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchKubernetesProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchKubernetesProfiles`: KubernetesProfilesSearchList
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchKubernetesProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchKubernetesProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kubernetesProfilesSearchCommand** | [**KubernetesProfilesSearchCommand**](KubernetesProfilesSearchCommand.md) |  | 

### Return type

[**KubernetesProfilesSearchList**](KubernetesProfilesSearchList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchNamespaces

> NamespaceSearchList SearchNamespaces(ctx).NamespaceSearchCommand(namespaceSearchCommand).Execute()

Global search for namespaces

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
	namespaceSearchCommand := *openapiclient.NewNamespaceSearchCommand(NullableInt32(123), NullableInt32(123), "SearchTerm_example", false) // NamespaceSearchCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchNamespaces(context.Background()).NamespaceSearchCommand(namespaceSearchCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchNamespaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchNamespaces`: NamespaceSearchList
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchNamespaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchNamespacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespaceSearchCommand** | [**NamespaceSearchCommand**](NamespaceSearchCommand.md) |  | 

### Return type

[**NamespaceSearchList**](NamespaceSearchList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchNetworkPolicies

> NetworkPoliciesSearchList SearchNetworkPolicies(ctx).NetworkPoliciesSearchCommand(networkPoliciesSearchCommand).Execute()

Global search for namespaces

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
	networkPoliciesSearchCommand := *openapiclient.NewNetworkPoliciesSearchCommand(NullableInt32(123), NullableInt32(123), "SearchTerm_example", false) // NetworkPoliciesSearchCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchNetworkPolicies(context.Background()).NetworkPoliciesSearchCommand(networkPoliciesSearchCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchNetworkPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchNetworkPolicies`: NetworkPoliciesSearchList
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchNetworkPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchNetworkPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkPoliciesSearchCommand** | [**NetworkPoliciesSearchCommand**](NetworkPoliciesSearchCommand.md) |  | 

### Return type

[**NetworkPoliciesSearchList**](NetworkPoliciesSearchList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchNodes

> NodesSearchList SearchNodes(ctx).NodesSearchCommand(nodesSearchCommand).Execute()

Global search for nodes

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
	nodesSearchCommand := *openapiclient.NewNodesSearchCommand(NullableInt32(123), NullableInt32(123), "SearchTerm_example", false) // NodesSearchCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchNodes(context.Background()).NodesSearchCommand(nodesSearchCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchNodes`: NodesSearchList
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchNodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nodesSearchCommand** | [**NodesSearchCommand**](NodesSearchCommand.md) |  | 

### Return type

[**NodesSearchList**](NodesSearchList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchOrganizations

> OrganizationSearchList SearchOrganizations(ctx).OrganizationSearchCommand(organizationSearchCommand).Execute()

Global search for organizations

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
	organizationSearchCommand := *openapiclient.NewOrganizationSearchCommand(NullableInt32(123), NullableInt32(123), "SearchTerm_example") // OrganizationSearchCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchOrganizations(context.Background()).OrganizationSearchCommand(organizationSearchCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchOrganizations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchOrganizations`: OrganizationSearchList
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchOrganizations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchOrganizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationSearchCommand** | [**OrganizationSearchCommand**](OrganizationSearchCommand.md) |  | 

### Return type

[**OrganizationSearchList**](OrganizationSearchList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchPdb

> PdbSearchList SearchPdb(ctx).PdbSearchCommand(pdbSearchCommand).Execute()

Global search for pdb

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
	pdbSearchCommand := *openapiclient.NewPdbSearchCommand(NullableInt32(123), NullableInt32(123), "SearchTerm_example", false) // PdbSearchCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchPdb(context.Background()).PdbSearchCommand(pdbSearchCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchPdb``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchPdb`: PdbSearchList
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchPdb`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchPdbRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pdbSearchCommand** | [**PdbSearchCommand**](PdbSearchCommand.md) |  | 

### Return type

[**PdbSearchList**](PdbSearchList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchPods

> PodsSearchList SearchPods(ctx).PodsSearchCommand(podsSearchCommand).Execute()

Global search for pods

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
	podsSearchCommand := *openapiclient.NewPodsSearchCommand(NullableInt32(123), NullableInt32(123), "SearchTerm_example", false) // PodsSearchCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchPods(context.Background()).PodsSearchCommand(podsSearchCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchPods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchPods`: PodsSearchList
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchPods`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchPodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **podsSearchCommand** | [**PodsSearchCommand**](PodsSearchCommand.md) |  | 

### Return type

[**PodsSearchList**](PodsSearchList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchProjects

> ProjectsSearchList SearchProjects(ctx).ProjectsSearchCommand(projectsSearchCommand).Execute()

Global search for projects

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
	projectsSearchCommand := *openapiclient.NewProjectsSearchCommand(NullableInt32(123), NullableInt32(123), "SearchTerm_example") // ProjectsSearchCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchProjects(context.Background()).ProjectsSearchCommand(projectsSearchCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchProjects`: ProjectsSearchList
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectsSearchCommand** | [**ProjectsSearchCommand**](ProjectsSearchCommand.md) |  | 

### Return type

[**ProjectsSearchList**](ProjectsSearchList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchPrometheusRules

> PrometheusRulesSearchList SearchPrometheusRules(ctx).PrometheusRulesSearchCommand(prometheusRulesSearchCommand).Execute()

Global search for prometheus-rules

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
	prometheusRulesSearchCommand := *openapiclient.NewPrometheusRulesSearchCommand(NullableInt32(123), NullableInt32(123), "SearchTerm_example") // PrometheusRulesSearchCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchPrometheusRules(context.Background()).PrometheusRulesSearchCommand(prometheusRulesSearchCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchPrometheusRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchPrometheusRules`: PrometheusRulesSearchList
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchPrometheusRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchPrometheusRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **prometheusRulesSearchCommand** | [**PrometheusRulesSearchCommand**](PrometheusRulesSearchCommand.md) |  | 

### Return type

[**PrometheusRulesSearchList**](PrometheusRulesSearchList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchPvcs

> PvcSearchList SearchPvcs(ctx).PvcSearchCommand(pvcSearchCommand).Execute()

Global search for pvcs

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
	pvcSearchCommand := *openapiclient.NewPvcSearchCommand(NullableInt32(123), NullableInt32(123), "SearchTerm_example", false) // PvcSearchCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchPvcs(context.Background()).PvcSearchCommand(pvcSearchCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchPvcs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchPvcs`: PvcSearchList
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchPvcs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchPvcsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pvcSearchCommand** | [**PvcSearchCommand**](PvcSearchCommand.md) |  | 

### Return type

[**PvcSearchList**](PvcSearchList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSecrets

> SecretSearchList SearchSecrets(ctx).SecretSearchCommand(secretSearchCommand).Execute()

Global search for secrets

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
	secretSearchCommand := *openapiclient.NewSecretSearchCommand(NullableInt32(123), NullableInt32(123), "SearchTerm_example", false) // SecretSearchCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchSecrets(context.Background()).SecretSearchCommand(secretSearchCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchSecrets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSecrets`: SecretSearchList
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchSecrets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **secretSearchCommand** | [**SecretSearchCommand**](SecretSearchCommand.md) |  | 

### Return type

[**SecretSearchList**](SecretSearchList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchServers

> ServersSearchList SearchServers(ctx).ServersSearchCommand(serversSearchCommand).Execute()

Global search for servers

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
	serversSearchCommand := *openapiclient.NewServersSearchCommand(NullableInt32(123), NullableInt32(123), "SearchTerm_example") // ServersSearchCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchServers(context.Background()).ServersSearchCommand(serversSearchCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchServers`: ServersSearchList
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchServers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serversSearchCommand** | [**ServersSearchCommand**](ServersSearchCommand.md) |  | 

### Return type

[**ServersSearchList**](ServersSearchList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchServices

> ServiceSearchList SearchServices(ctx).ServiceSearchCommand(serviceSearchCommand).Execute()

Global search for services

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
	serviceSearchCommand := *openapiclient.NewServiceSearchCommand(NullableInt32(123), NullableInt32(123), "SearchTerm_example", false) // ServiceSearchCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchServices(context.Background()).ServiceSearchCommand(serviceSearchCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchServices`: ServiceSearchList
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceSearchCommand** | [**ServiceSearchCommand**](ServiceSearchCommand.md) |  | 

### Return type

[**ServiceSearchList**](ServiceSearchList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchStandAloneProfiles

> StandAloneProfilesSearchList SearchStandAloneProfiles(ctx).StandAloneProfilesSearchCommand(standAloneProfilesSearchCommand).Execute()

Global search for stand-alone-profiles

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
	standAloneProfilesSearchCommand := *openapiclient.NewStandAloneProfilesSearchCommand(NullableInt32(123), NullableInt32(123), "SearchTerm_example") // StandAloneProfilesSearchCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchStandAloneProfiles(context.Background()).StandAloneProfilesSearchCommand(standAloneProfilesSearchCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchStandAloneProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchStandAloneProfiles`: StandAloneProfilesSearchList
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchStandAloneProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchStandAloneProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **standAloneProfilesSearchCommand** | [**StandAloneProfilesSearchCommand**](StandAloneProfilesSearchCommand.md) |  | 

### Return type

[**StandAloneProfilesSearchList**](StandAloneProfilesSearchList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchStorageClasses

> StorageClassesSearchList SearchStorageClasses(ctx).StorageClassesSearchCommand(storageClassesSearchCommand).Execute()

Global search for storage classes.

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
	storageClassesSearchCommand := *openapiclient.NewStorageClassesSearchCommand(NullableInt32(123), NullableInt32(123), "SearchTerm_example", false) // StorageClassesSearchCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchStorageClasses(context.Background()).StorageClassesSearchCommand(storageClassesSearchCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchStorageClasses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchStorageClasses`: StorageClassesSearchList
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchStorageClasses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchStorageClassesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storageClassesSearchCommand** | [**StorageClassesSearchCommand**](StorageClassesSearchCommand.md) |  | 

### Return type

[**StorageClassesSearchList**](StorageClassesSearchList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSts

> StsSearchList SearchSts(ctx).StsSearchCommand(stsSearchCommand).Execute()

Global search for sts

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
	stsSearchCommand := *openapiclient.NewStsSearchCommand(NullableInt32(123), NullableInt32(123), "SearchTerm_example", false) // StsSearchCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchSts(context.Background()).StsSearchCommand(stsSearchCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchSts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSts`: StsSearchList
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchSts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchStsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stsSearchCommand** | [**StsSearchCommand**](StsSearchCommand.md) |  | 

### Return type

[**StsSearchList**](StsSearchList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchUsers

> UsersSearchList SearchUsers(ctx).UsersSearchCommand(usersSearchCommand).Execute()

Global search for users

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
	usersSearchCommand := *openapiclient.NewUsersSearchCommand(NullableInt32(123), NullableInt32(123), "SearchTerm_example") // UsersSearchCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchUsers(context.Background()).UsersSearchCommand(usersSearchCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchUsers`: UsersSearchList
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **usersSearchCommand** | [**UsersSearchCommand**](UsersSearchCommand.md) |  | 

### Return type

[**UsersSearchList**](UsersSearchList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

