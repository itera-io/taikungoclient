# \CronJobServiceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CronJobScanCertificateCredentialExpiryLegacy**](CronJobServiceAPI.md#CronJobScanCertificateCredentialExpiryLegacy) | **Post** /api/v1/cronjob/scan-certificate-credential-expiry | Legacy alias for scan-certificate-profile-expiry (pre-rename cron path)
[**CronJobScanCertificateProfileExpiry**](CronJobServiceAPI.md#CronJobScanCertificateProfileExpiry) | **Post** /api/v1/cronjob/scan-certificate-profile-expiry | Scan certificate profiles approaching or past expiry and emit alert logs
[**CronjobAutoUpgradeProjects**](CronJobServiceAPI.md#CronjobAutoUpgradeProjects) | **Post** /api/v1/cronjob/auto-upgrade-projects | Upgrade projects that auto-upgrade option enabled
[**CronjobCcfCryptoStatus**](CronJobServiceAPI.md#CronjobCcfCryptoStatus) | **Get** /api/v1/cronjob/ccf-crypto/status | Rollback-safety survey: per-column $enc:v1: row counts and last Transit failure timestamp
[**CronjobCreateKeyPool**](CronJobServiceAPI.md#CronjobCreateKeyPool) | **Post** /api/v1/cronjob/create-key-pool | Create key pool
[**CronjobDeleteExpiredAlerts**](CronJobServiceAPI.md#CronjobDeleteExpiredAlerts) | **Post** /api/v1/cronjob/alerts | Delete expired alerts
[**CronjobDeleteExpiredHistoryLogs**](CronJobServiceAPI.md#CronjobDeleteExpiredHistoryLogs) | **Post** /api/v1/cronjob/history-logs | Delete expired history logs
[**CronjobDeleteExpiredRefreshTokens**](CronJobServiceAPI.md#CronjobDeleteExpiredRefreshTokens) | **Post** /api/v1/cronjob/refresh-tokens | Delete expired refresh tokens
[**CronjobDeleteExpiredServers**](CronJobServiceAPI.md#CronjobDeleteExpiredServers) | **Post** /api/v1/cronjob/delete-expired-servers | Delete expired servers
[**CronjobDeleteImportedBackupLocation**](CronJobServiceAPI.md#CronjobDeleteImportedBackupLocation) | **Post** /api/v1/cronjob/backup-locations | Delete imported backup locations
[**CronjobDeleteKubeConfigs**](CronJobServiceAPI.md#CronjobDeleteKubeConfigs) | **Post** /api/v1/cronjob/delete-kube-configs | Remove deleted user&#39;s kube config
[**CronjobDeleteRemovedSpotInstances**](CronJobServiceAPI.md#CronjobDeleteRemovedSpotInstances) | **Post** /api/v1/cronjob/delete-removed-spot-instances | Delete removed spot instances
[**CronjobDeleteUselessProjectActions**](CronJobServiceAPI.md#CronjobDeleteUselessProjectActions) | **Post** /api/v1/cronjob/project-actions | Delete useless project actions
[**CronjobEmailForProjectExpiration**](CronJobServiceAPI.md#CronjobEmailForProjectExpiration) | **Post** /api/v1/cronjob/project-expiration | Send email to the users about project expiration
[**CronjobFetchArtifactOrganizations**](CronJobServiceAPI.md#CronjobFetchArtifactOrganizations) | **Post** /api/v1/cronjob/fetch-artifact-organizations | Fetch artifact hub organizations
[**CronjobFetchAzureFlavorPrices**](CronJobServiceAPI.md#CronjobFetchAzureFlavorPrices) | **Post** /api/v1/cronjob/fetch-azure-flavor-prices | Fetch azure flavor prices
[**CronjobFetchK8sAlertData**](CronJobServiceAPI.md#CronjobFetchK8sAlertData) | **Post** /api/v1/cronjob/fetch-k8s-alert-data | Fetch k8s alert data
[**CronjobFetchK8sOverviewData**](CronJobServiceAPI.md#CronjobFetchK8sOverviewData) | **Post** /api/v1/cronjob/fetch-k8s-overview-data | Fetch k8s overview data
[**CronjobPurgeExpiredProjects**](CronJobServiceAPI.md#CronjobPurgeExpiredProjects) | **Post** /api/v1/cronjob/purge-expired-projects | Purge expired projects
[**CronjobReconcilePartnerScopes**](CronJobServiceAPI.md#CronjobReconcilePartnerScopes) | **Post** /api/v1/cronjob/reconcile-partner-scopes | Reconcile CCF partner scope holders (moves DEK rows off soft-deleted partner scope holders)
[**CronjobRemindUsersByAlertingProfile**](CronJobServiceAPI.md#CronjobRemindUsersByAlertingProfile) | **Post** /api/v1/cronjob/remind-users-by-alerting-profile | Remind users by alerting profile
[**CronjobRotateTenantKeks**](CronJobServiceAPI.md#CronjobRotateTenantKeks) | **Post** /api/v1/cronjob/rotate-tenant-keks | Rotate tenant KEKs via Transit rewrap
[**CronjobSyncBackupCredentials**](CronJobServiceAPI.md#CronjobSyncBackupCredentials) | **Post** /api/v1/cronjob/sync-backup-credentials | Sync backup credentials
[**CronjobSyncDnsCertConfigs**](CronJobServiceAPI.md#CronjobSyncDnsCertConfigs) | **Post** /api/v1/cronjob/sync-dns-cert-configs | Sync DNS/Cert configs (drift cron — re-fires kubevap addons pipeline when Applied &lt; Desired)
[**CronjobSyncOpaProfiles**](CronJobServiceAPI.md#CronjobSyncOpaProfiles) | **Post** /api/v1/cronjob/sync-opa-profiles | Sync opa profiles
[**CronjobSyncOrganizations**](CronJobServiceAPI.md#CronjobSyncOrganizations) | **Post** /api/v1/cronjob/sync-organizations | Sync organizations
[**CronjobSyncProjectApps**](CronJobServiceAPI.md#CronjobSyncProjectApps) | **Post** /api/v1/cronjob/sync-project-apps | Sync project apps
[**CronjobSyncProjects**](CronJobServiceAPI.md#CronjobSyncProjects) | **Post** /api/v1/cronjob/sync-projects | Sync projects
[**CronjobTektonPipelines**](CronJobServiceAPI.md#CronjobTektonPipelines) | **Post** /api/v1/cronjob/tekton-pipelines | Tekton pipelines
[**CronjobUpdateProjectAppStatus**](CronJobServiceAPI.md#CronjobUpdateProjectAppStatus) | **Post** /api/v1/cronjob/update-project-app-status | Update project app status
[**CronjobUpdateProjectQuotaMessage**](CronJobServiceAPI.md#CronjobUpdateProjectQuotaMessage) | **Post** /api/v1/cronjob/update-project-quota-message | Update project quota message



## CronJobScanCertificateCredentialExpiryLegacy

> CronJobScanCertificateCredentialExpiryLegacy(ctx).Body(body).Execute()

Legacy alias for scan-certificate-profile-expiry (pre-rename cron path)

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
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CronJobServiceAPI.CronJobScanCertificateCredentialExpiryLegacy(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CronJobServiceAPI.CronJobScanCertificateCredentialExpiryLegacy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCronJobScanCertificateCredentialExpiryLegacyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

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


## CronJobScanCertificateProfileExpiry

> CronJobScanCertificateProfileExpiry(ctx).Body(body).Execute()

Scan certificate profiles approaching or past expiry and emit alert logs

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
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CronJobServiceAPI.CronJobScanCertificateProfileExpiry(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CronJobServiceAPI.CronJobScanCertificateProfileExpiry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCronJobScanCertificateProfileExpiryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

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


## CronjobAutoUpgradeProjects

> CronjobAutoUpgradeProjects(ctx).Body(body).Execute()

Upgrade projects that auto-upgrade option enabled

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
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CronJobServiceAPI.CronjobAutoUpgradeProjects(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CronJobServiceAPI.CronjobAutoUpgradeProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCronjobAutoUpgradeProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

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


## CronjobCcfCryptoStatus

> TransitEncryptionStatusDto CronjobCcfCryptoStatus(ctx).Execute()

Rollback-safety survey: per-column $enc:v1: row counts and last Transit failure timestamp

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
	resp, r, err := apiClient.CronJobServiceAPI.CronjobCcfCryptoStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CronJobServiceAPI.CronjobCcfCryptoStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CronjobCcfCryptoStatus`: TransitEncryptionStatusDto
	fmt.Fprintf(os.Stdout, "Response from `CronJobServiceAPI.CronjobCcfCryptoStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCronjobCcfCryptoStatusRequest struct via the builder pattern


### Return type

[**TransitEncryptionStatusDto**](TransitEncryptionStatusDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CronjobCreateKeyPool

> CronjobCreateKeyPool(ctx).Body(body).Execute()

Create key pool

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
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CronJobServiceAPI.CronjobCreateKeyPool(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CronJobServiceAPI.CronjobCreateKeyPool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCronjobCreateKeyPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

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


## CronjobDeleteExpiredAlerts

> CronjobDeleteExpiredAlerts(ctx).Body(body).Execute()

Delete expired alerts

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
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CronJobServiceAPI.CronjobDeleteExpiredAlerts(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CronJobServiceAPI.CronjobDeleteExpiredAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCronjobDeleteExpiredAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

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


## CronjobDeleteExpiredHistoryLogs

> CronjobDeleteExpiredHistoryLogs(ctx).Body(body).Execute()

Delete expired history logs

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
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CronJobServiceAPI.CronjobDeleteExpiredHistoryLogs(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CronJobServiceAPI.CronjobDeleteExpiredHistoryLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCronjobDeleteExpiredHistoryLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

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


## CronjobDeleteExpiredRefreshTokens

> CronjobDeleteExpiredRefreshTokens(ctx).Body(body).Execute()

Delete expired refresh tokens

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
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CronJobServiceAPI.CronjobDeleteExpiredRefreshTokens(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CronJobServiceAPI.CronjobDeleteExpiredRefreshTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCronjobDeleteExpiredRefreshTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

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


## CronjobDeleteExpiredServers

> CronjobDeleteExpiredServers(ctx).Body(body).Execute()

Delete expired servers

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
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CronJobServiceAPI.CronjobDeleteExpiredServers(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CronJobServiceAPI.CronjobDeleteExpiredServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCronjobDeleteExpiredServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

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


## CronjobDeleteImportedBackupLocation

> CronjobDeleteImportedBackupLocation(ctx).Body(body).Execute()

Delete imported backup locations

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
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CronJobServiceAPI.CronjobDeleteImportedBackupLocation(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CronJobServiceAPI.CronjobDeleteImportedBackupLocation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCronjobDeleteImportedBackupLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

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


## CronjobDeleteKubeConfigs

> CronjobDeleteKubeConfigs(ctx).Body(body).Execute()

Remove deleted user's kube config

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
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CronJobServiceAPI.CronjobDeleteKubeConfigs(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CronJobServiceAPI.CronjobDeleteKubeConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCronjobDeleteKubeConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

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


## CronjobDeleteRemovedSpotInstances

> CronjobDeleteRemovedSpotInstances(ctx).Body(body).Execute()

Delete removed spot instances

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
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CronJobServiceAPI.CronjobDeleteRemovedSpotInstances(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CronJobServiceAPI.CronjobDeleteRemovedSpotInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCronjobDeleteRemovedSpotInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

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


## CronjobDeleteUselessProjectActions

> CronjobDeleteUselessProjectActions(ctx).Body(body).Execute()

Delete useless project actions

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
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CronJobServiceAPI.CronjobDeleteUselessProjectActions(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CronJobServiceAPI.CronjobDeleteUselessProjectActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCronjobDeleteUselessProjectActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

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


## CronjobEmailForProjectExpiration

> CronjobEmailForProjectExpiration(ctx).Body(body).Execute()

Send email to the users about project expiration

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
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CronJobServiceAPI.CronjobEmailForProjectExpiration(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CronJobServiceAPI.CronjobEmailForProjectExpiration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCronjobEmailForProjectExpirationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

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


## CronjobFetchArtifactOrganizations

> CronjobFetchArtifactOrganizations(ctx).Body(body).Execute()

Fetch artifact hub organizations

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
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CronJobServiceAPI.CronjobFetchArtifactOrganizations(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CronJobServiceAPI.CronjobFetchArtifactOrganizations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCronjobFetchArtifactOrganizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

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


## CronjobFetchAzureFlavorPrices

> CronjobFetchAzureFlavorPrices(ctx).Body(body).Execute()

Fetch azure flavor prices

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
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CronJobServiceAPI.CronjobFetchAzureFlavorPrices(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CronJobServiceAPI.CronjobFetchAzureFlavorPrices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCronjobFetchAzureFlavorPricesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

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


## CronjobFetchK8sAlertData

> CronjobFetchK8sAlertData(ctx).Body(body).Execute()

Fetch k8s alert data

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
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CronJobServiceAPI.CronjobFetchK8sAlertData(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CronJobServiceAPI.CronjobFetchK8sAlertData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCronjobFetchK8sAlertDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

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


## CronjobFetchK8sOverviewData

> CronjobFetchK8sOverviewData(ctx).Body(body).Execute()

Fetch k8s overview data

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
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CronJobServiceAPI.CronjobFetchK8sOverviewData(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CronJobServiceAPI.CronjobFetchK8sOverviewData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCronjobFetchK8sOverviewDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

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


## CronjobPurgeExpiredProjects

> CronjobPurgeExpiredProjects(ctx).Body(body).Execute()

Purge expired projects

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
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CronJobServiceAPI.CronjobPurgeExpiredProjects(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CronJobServiceAPI.CronjobPurgeExpiredProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCronjobPurgeExpiredProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

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


## CronjobReconcilePartnerScopes

> PartnerScopeReconcileResult CronjobReconcilePartnerScopes(ctx).Execute()

Reconcile CCF partner scope holders (moves DEK rows off soft-deleted partner scope holders)

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
	resp, r, err := apiClient.CronJobServiceAPI.CronjobReconcilePartnerScopes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CronJobServiceAPI.CronjobReconcilePartnerScopes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CronjobReconcilePartnerScopes`: PartnerScopeReconcileResult
	fmt.Fprintf(os.Stdout, "Response from `CronJobServiceAPI.CronjobReconcilePartnerScopes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCronjobReconcilePartnerScopesRequest struct via the builder pattern


### Return type

[**PartnerScopeReconcileResult**](PartnerScopeReconcileResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CronjobRemindUsersByAlertingProfile

> CronjobRemindUsersByAlertingProfile(ctx).RemindUsersByAlertingProfileCommand(remindUsersByAlertingProfileCommand).Execute()

Remind users by alerting profile

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
	remindUsersByAlertingProfileCommand := *openapiclient.NewRemindUsersByAlertingProfileCommand(map[string]interface{}{"key": interface{}(123)}) // RemindUsersByAlertingProfileCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CronJobServiceAPI.CronjobRemindUsersByAlertingProfile(context.Background()).RemindUsersByAlertingProfileCommand(remindUsersByAlertingProfileCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CronJobServiceAPI.CronjobRemindUsersByAlertingProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCronjobRemindUsersByAlertingProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **remindUsersByAlertingProfileCommand** | [**RemindUsersByAlertingProfileCommand**](RemindUsersByAlertingProfileCommand.md) |  | 

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


## CronjobRotateTenantKeks

> CronjobRotateTenantKeks(ctx).RotateTenantKeksCommand(rotateTenantKeksCommand).Execute()

Rotate tenant KEKs via Transit rewrap

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
	rotateTenantKeksCommand := *openapiclient.NewRotateTenantKeksCommand() // RotateTenantKeksCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CronJobServiceAPI.CronjobRotateTenantKeks(context.Background()).RotateTenantKeksCommand(rotateTenantKeksCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CronJobServiceAPI.CronjobRotateTenantKeks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCronjobRotateTenantKeksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rotateTenantKeksCommand** | [**RotateTenantKeksCommand**](RotateTenantKeksCommand.md) |  | 

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


## CronjobSyncBackupCredentials

> CronjobSyncBackupCredentials(ctx).Body(body).Execute()

Sync backup credentials

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
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CronJobServiceAPI.CronjobSyncBackupCredentials(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CronJobServiceAPI.CronjobSyncBackupCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCronjobSyncBackupCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

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


## CronjobSyncDnsCertConfigs

> CronjobSyncDnsCertConfigs(ctx).Body(body).Execute()

Sync DNS/Cert configs (drift cron — re-fires kubevap addons pipeline when Applied < Desired)

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
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CronJobServiceAPI.CronjobSyncDnsCertConfigs(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CronJobServiceAPI.CronjobSyncDnsCertConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCronjobSyncDnsCertConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

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


## CronjobSyncOpaProfiles

> CronjobSyncOpaProfiles(ctx).Body(body).Execute()

Sync opa profiles

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
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CronJobServiceAPI.CronjobSyncOpaProfiles(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CronJobServiceAPI.CronjobSyncOpaProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCronjobSyncOpaProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

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


## CronjobSyncOrganizations

> CronjobSyncOrganizations(ctx).Body(body).Execute()

Sync organizations

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
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CronJobServiceAPI.CronjobSyncOrganizations(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CronJobServiceAPI.CronjobSyncOrganizations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCronjobSyncOrganizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

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


## CronjobSyncProjectApps

> CronjobSyncProjectApps(ctx).SyncProjectAppsCommand(syncProjectAppsCommand).Execute()

Sync project apps

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
	syncProjectAppsCommand := *openapiclient.NewSyncProjectAppsCommand() // SyncProjectAppsCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CronJobServiceAPI.CronjobSyncProjectApps(context.Background()).SyncProjectAppsCommand(syncProjectAppsCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CronJobServiceAPI.CronjobSyncProjectApps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCronjobSyncProjectAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **syncProjectAppsCommand** | [**SyncProjectAppsCommand**](SyncProjectAppsCommand.md) |  | 

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


## CronjobSyncProjects

> CronjobSyncProjects(ctx).Body(body).Execute()

Sync projects

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
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CronJobServiceAPI.CronjobSyncProjects(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CronJobServiceAPI.CronjobSyncProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCronjobSyncProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

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


## CronjobTektonPipelines

> CronjobTektonPipelines(ctx).Body(body).Execute()

Tekton pipelines

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
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CronJobServiceAPI.CronjobTektonPipelines(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CronJobServiceAPI.CronjobTektonPipelines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCronjobTektonPipelinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

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


## CronjobUpdateProjectAppStatus

> CronjobUpdateProjectAppStatus(ctx).Body(body).Execute()

Update project app status

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
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CronJobServiceAPI.CronjobUpdateProjectAppStatus(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CronJobServiceAPI.CronjobUpdateProjectAppStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCronjobUpdateProjectAppStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

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


## CronjobUpdateProjectQuotaMessage

> CronjobUpdateProjectQuotaMessage(ctx).Body(body).Execute()

Update project quota message

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
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CronJobServiceAPI.CronjobUpdateProjectQuotaMessage(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CronJobServiceAPI.CronjobUpdateProjectQuotaMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCronjobUpdateProjectQuotaMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

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

