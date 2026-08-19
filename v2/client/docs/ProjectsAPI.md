# \ProjectsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectsAiAnalyzer**](ProjectsAPI.md#ProjectsAiAnalyzer) | **Get** /api/v1/projects/ai-analyze/{projectId} | Analyze cluster by AI model
[**ProjectsAlerts**](ProjectsAPI.md#ProjectsAlerts) | **Post** /api/v1/projects/alerts | Project alerts
[**ProjectsCanAddVcluster**](ProjectsAPI.md#ProjectsCanAddVcluster) | **Get** /api/v1/projects/can-add-vcluster/{projectId} | Visibility of adding vcluster
[**ProjectsCatalogs**](ProjectsAPI.md#ProjectsCatalogs) | **Get** /api/v1/projects/catalogs/{projectId} | Catalogs for project
[**ProjectsChatCompletions**](ProjectsAPI.md#ProjectsChatCompletions) | **Post** /api/v1/projects/chat/completions | AI Chat completions
[**ProjectsClearAirgapOverride**](ProjectsAPI.md#ProjectsClearAirgapOverride) | **Delete** /api/v1/projects/airgap-override/{projectId} | Remove per-project airgap registry override (reverts to global default)
[**ProjectsCreate**](ProjectsAPI.md#ProjectsCreate) | **Post** /api/v1/projects | Create a new project
[**ProjectsDelete**](ProjectsAPI.md#ProjectsDelete) | **Post** /api/v1/projects/delete | Delete the project. The project must be empty (no server) and in READY state
[**ProjectsDescribe**](ProjectsAPI.md#ProjectsDescribe) | **Get** /api/v1/projects/describe/{projectId} | Describe project by Id
[**ProjectsDropdown**](ProjectsAPI.md#ProjectsDropdown) | **Get** /api/v1/projects/list | Retrieve list of projects for dropdown
[**ProjectsEditHealth**](ProjectsAPI.md#ProjectsEditHealth) | **Put** /api/v1/projects/edit/health | Update health status of the project by Id
[**ProjectsEditStatus**](ProjectsAPI.md#ProjectsEditStatus) | **Put** /api/v1/projects/edit/status | Change the project status for the given project. Only available for admin.
[**ProjectsExportLokiLogs**](ProjectsAPI.md#ProjectsExportLokiLogs) | **Post** /api/v1/projects/loki-logs/export | Export loki logs
[**ProjectsExtendLifetime**](ProjectsAPI.md#ProjectsExtendLifetime) | **Post** /api/v1/projects/extend/lifetime | Extend life time of project
[**ProjectsForBilling**](ProjectsAPI.md#ProjectsForBilling) | **Get** /api/v1/projects/forbilling | Retrieve a list of projects for billing
[**ProjectsGetAccessibleUsers**](ProjectsAPI.md#ProjectsGetAccessibleUsers) | **Get** /api/v1/projects/get-accessible-users | Get accessible users for project
[**ProjectsGetAirgapOverride**](ProjectsAPI.md#ProjectsGetAirgapOverride) | **Get** /api/v1/projects/airgap-override/{projectId} | Get per-project airgap registry override
[**ProjectsImportedClusterDetails**](ProjectsAPI.md#ProjectsImportedClusterDetails) | **Get** /api/v1/projects/imported/details/{projectId} | Imported cluster details
[**ProjectsList**](ProjectsAPI.md#ProjectsList) | **Get** /api/v1/projects | Retrieve all projects
[**ProjectsLockManager**](ProjectsAPI.md#ProjectsLockManager) | **Post** /api/v1/projects/lockmanager | Lock/Unlock project
[**ProjectsLokiLogs**](ProjectsAPI.md#ProjectsLokiLogs) | **Post** /api/v1/projects/loki-logs | Retrieve loki logs
[**ProjectsMaintenanceManager**](ProjectsAPI.md#ProjectsMaintenanceManager) | **Post** /api/v1/projects/maintenance-manager | Enable/disable project&#39;s maintenance mode
[**ProjectsMonitoringAlerts**](ProjectsAPI.md#ProjectsMonitoringAlerts) | **Post** /api/v1/projects/monitoringalerts | Monitoring alerts for project
[**ProjectsPrometheusMetrics**](ProjectsAPI.md#ProjectsPrometheusMetrics) | **Post** /api/v1/projects/prometheusmetrics | Prometheus metrics data project
[**ProjectsPrometheusMetricsAutocomplete**](ProjectsAPI.md#ProjectsPrometheusMetricsAutocomplete) | **Post** /api/v1/projects/prometheusmetrics/autocomplete | Prometheus metrics autocomplete values
[**ProjectsSetAirgapOverride**](ProjectsAPI.md#ProjectsSetAirgapOverride) | **Post** /api/v1/projects/airgap-override/set | Set per-project airgap registry override
[**ProjectsToggleFullSpot**](ProjectsAPI.md#ProjectsToggleFullSpot) | **Post** /api/v1/projects/toggle-full-spot | Full spot operations enable/disable
[**ProjectsToggleSpotVms**](ProjectsAPI.md#ProjectsToggleSpotVms) | **Post** /api/v1/projects/toggle-spot-vms | Spot vm(s) operations enable/disable
[**ProjectsToggleSpotWorkers**](ProjectsAPI.md#ProjectsToggleSpotWorkers) | **Post** /api/v1/projects/toggle-spot-workers | Spot worker(s) operations enable/disable
[**ProjectsVisibility**](ProjectsAPI.md#ProjectsVisibility) | **Get** /api/v1/projects/visibility/{projectId} | Visibility of project actions



## ProjectsAiAnalyzer

> AiListDto ProjectsAiAnalyzer(ctx, projectId).Execute()

Analyze cluster by AI model

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
	resp, r, err := apiClient.ProjectsAPI.ProjectsAiAnalyzer(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsAiAnalyzer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsAiAnalyzer`: AiListDto
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectsAiAnalyzer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsAiAnalyzerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AiListDto**](AiListDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsAlerts

> []ProjectDetailsErrorListDto ProjectsAlerts(ctx).ProjectAlertsQuery(projectAlertsQuery).Execute()

Project alerts

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
	projectAlertsQuery := *openapiclient.NewProjectAlertsQuery(int32(123), openapiclient.ProjectType("K8S")) // ProjectAlertsQuery |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ProjectsAlerts(context.Background()).ProjectAlertsQuery(projectAlertsQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsAlerts`: []ProjectDetailsErrorListDto
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectsAlerts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectAlertsQuery** | [**ProjectAlertsQuery**](ProjectAlertsQuery.md) |  | 

### Return type

[**[]ProjectDetailsErrorListDto**](ProjectDetailsErrorListDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsCanAddVcluster

> ProjectCanAddVClusterDto ProjectsCanAddVcluster(ctx, projectId).Execute()

Visibility of adding vcluster

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
	resp, r, err := apiClient.ProjectsAPI.ProjectsCanAddVcluster(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsCanAddVcluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsCanAddVcluster`: ProjectCanAddVClusterDto
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectsCanAddVcluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsCanAddVclusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectCanAddVClusterDto**](ProjectCanAddVClusterDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsCatalogs

> []CatalogsForProjectDto ProjectsCatalogs(ctx, projectId).Execute()

Catalogs for project

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
	resp, r, err := apiClient.ProjectsAPI.ProjectsCatalogs(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsCatalogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsCatalogs`: []CatalogsForProjectDto
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectsCatalogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsCatalogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]CatalogsForProjectDto**](CatalogsForProjectDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsChatCompletions

> map[string]interface{} ProjectsChatCompletions(ctx).ChatCompletionsCommand(chatCompletionsCommand).Execute()

AI Chat completions

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
	chatCompletionsCommand := *openapiclient.NewChatCompletionsCommand(map[string]interface{}{"key": interface{}(123)}) // ChatCompletionsCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ProjectsChatCompletions(context.Background()).ChatCompletionsCommand(chatCompletionsCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsChatCompletions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsChatCompletions`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectsChatCompletions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsChatCompletionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatCompletionsCommand** | [**ChatCompletionsCommand**](ChatCompletionsCommand.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsClearAirgapOverride

> ProjectsClearAirgapOverride(ctx, projectId).Execute()

Remove per-project airgap registry override (reverts to global default)

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
	r, err := apiClient.ProjectsAPI.ProjectsClearAirgapOverride(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsClearAirgapOverride``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsClearAirgapOverrideRequest struct via the builder pattern


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


## ProjectsCreate

> ApiResponse ProjectsCreate(ctx).CreateProjectCommand(createProjectCommand).Execute()

Create a new project

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
	createProjectCommand := *openapiclient.NewCreateProjectCommand() // CreateProjectCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ProjectsCreate(context.Background()).CreateProjectCommand(createProjectCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsCreate`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createProjectCommand** | [**CreateProjectCommand**](CreateProjectCommand.md) |  | 

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


## ProjectsDelete

> ProjectsDelete(ctx).DeleteProjectCommand(deleteProjectCommand).Execute()

Delete the project. The project must be empty (no server) and in READY state

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
	deleteProjectCommand := *openapiclient.NewDeleteProjectCommand(int32(123), false) // DeleteProjectCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.ProjectsDelete(context.Background()).DeleteProjectCommand(deleteProjectCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteProjectCommand** | [**DeleteProjectCommand**](DeleteProjectCommand.md) |  | 

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


## ProjectsDescribe

> ProjectDescribeDto ProjectsDescribe(ctx, projectId).Execute()

Describe project by Id

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
	resp, r, err := apiClient.ProjectsAPI.ProjectsDescribe(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsDescribe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsDescribe`: ProjectDescribeDto
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectsDescribe`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsDescribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectDescribeDto**](ProjectDescribeDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsDropdown

> []CommonDropdownIsBoundDtoForProject ProjectsDropdown(ctx).OrganizationId(organizationId).Search(search).CatalogId(catalogId).Healthy(healthy).UserId(userId).Ready(ready).IsBoundToCatalog(isBoundToCatalog).IsLocked(isLocked).Execute()

Retrieve list of projects for dropdown

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
	catalogId := int32(56) // int32 |  (optional)
	healthy := true // bool |  (optional)
	userId := "userId_example" // string |  (optional)
	ready := true // bool |  (optional)
	isBoundToCatalog := true // bool |  (optional)
	isLocked := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ProjectsDropdown(context.Background()).OrganizationId(organizationId).Search(search).CatalogId(catalogId).Healthy(healthy).UserId(userId).Ready(ready).IsBoundToCatalog(isBoundToCatalog).IsLocked(isLocked).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsDropdown``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsDropdown`: []CommonDropdownIsBoundDtoForProject
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectsDropdown`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsDropdownRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 
 **search** | **string** |  | 
 **catalogId** | **int32** |  | 
 **healthy** | **bool** |  | 
 **userId** | **string** |  | 
 **ready** | **bool** |  | 
 **isBoundToCatalog** | **bool** |  | 
 **isLocked** | **bool** |  | 

### Return type

[**[]CommonDropdownIsBoundDtoForProject**](CommonDropdownIsBoundDtoForProject.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsEditHealth

> ProjectsEditHealth(ctx).UpdateHealthStatusCommand(updateHealthStatusCommand).Execute()

Update health status of the project by Id

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
	updateHealthStatusCommand := *openapiclient.NewUpdateHealthStatusCommand() // UpdateHealthStatusCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.ProjectsEditHealth(context.Background()).UpdateHealthStatusCommand(updateHealthStatusCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsEditHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsEditHealthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateHealthStatusCommand** | [**UpdateHealthStatusCommand**](UpdateHealthStatusCommand.md) |  | 

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


## ProjectsEditStatus

> ProjectsEditStatus(ctx).ResetProjectStatusCommand(resetProjectStatusCommand).Execute()

Change the project status for the given project. Only available for admin.

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
	resetProjectStatusCommand := *openapiclient.NewResetProjectStatusCommand(int32(123), openapiclient.ProjectStatus("Null")) // ResetProjectStatusCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.ProjectsEditStatus(context.Background()).ResetProjectStatusCommand(resetProjectStatusCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsEditStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsEditStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resetProjectStatusCommand** | [**ResetProjectStatusCommand**](ResetProjectStatusCommand.md) |  | 

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


## ProjectsExportLokiLogs

> CsvExporter ProjectsExportLokiLogs(ctx).ExportLokiLogsCommand(exportLokiLogsCommand).Execute()

Export loki logs

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
	exportLokiLogsCommand := *openapiclient.NewExportLokiLogsCommand() // ExportLokiLogsCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ProjectsExportLokiLogs(context.Background()).ExportLokiLogsCommand(exportLokiLogsCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsExportLokiLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsExportLokiLogs`: CsvExporter
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectsExportLokiLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsExportLokiLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **exportLokiLogsCommand** | [**ExportLokiLogsCommand**](ExportLokiLogsCommand.md) |  | 

### Return type

[**CsvExporter**](CsvExporter.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsExtendLifetime

> ProjectsExtendLifetime(ctx).ProjectExtendLifeTimeCommand(projectExtendLifeTimeCommand).Execute()

Extend life time of project

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/itera-io/taikungoclient/v2/client"
)

func main() {
	projectExtendLifeTimeCommand := *openapiclient.NewProjectExtendLifeTimeCommand(int32(123), time.Now(), false) // ProjectExtendLifeTimeCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.ProjectsExtendLifetime(context.Background()).ProjectExtendLifeTimeCommand(projectExtendLifeTimeCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsExtendLifetime``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsExtendLifetimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectExtendLifeTimeCommand** | [**ProjectExtendLifeTimeCommand**](ProjectExtendLifeTimeCommand.md) |  | 

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


## ProjectsForBilling

> []ProjectsForBillingDto ProjectsForBilling(ctx).Execute()

Retrieve a list of projects for billing

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
	resp, r, err := apiClient.ProjectsAPI.ProjectsForBilling(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsForBilling``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsForBilling`: []ProjectsForBillingDto
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectsForBilling`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsForBillingRequest struct via the builder pattern


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


## ProjectsGetAccessibleUsers

> []CommonStringBasedDropdownDto ProjectsGetAccessibleUsers(ctx).ProjectId(projectId).Execute()

Get accessible users for project

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
	resp, r, err := apiClient.ProjectsAPI.ProjectsGetAccessibleUsers(context.Background()).ProjectId(projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsGetAccessibleUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsGetAccessibleUsers`: []CommonStringBasedDropdownDto
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectsGetAccessibleUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsGetAccessibleUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **int32** |  | 

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


## ProjectsGetAirgapOverride

> ProjectAirgapOverrideDto ProjectsGetAirgapOverride(ctx, projectId).Execute()

Get per-project airgap registry override

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
	resp, r, err := apiClient.ProjectsAPI.ProjectsGetAirgapOverride(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsGetAirgapOverride``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsGetAirgapOverride`: ProjectAirgapOverrideDto
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectsGetAirgapOverride`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsGetAirgapOverrideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectAirgapOverrideDto**](ProjectAirgapOverrideDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsImportedClusterDetails

> ImportedClusterDetailsDto ProjectsImportedClusterDetails(ctx, projectId).Execute()

Imported cluster details

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
	resp, r, err := apiClient.ProjectsAPI.ProjectsImportedClusterDetails(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsImportedClusterDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsImportedClusterDetails`: ImportedClusterDetailsDto
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectsImportedClusterDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsImportedClusterDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ImportedClusterDetailsDto**](ImportedClusterDetailsDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsList

> ProjectsList ProjectsList(ctx).Limit(limit).Offset(offset).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).UpdatedAt(updatedAt).SearchId(searchId).Id(id).BackupCredentialId(backupCredentialId).Healthy(healthy).DomainId(domainId).Execute()

Retrieve all projects

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/itera-io/taikungoclient/v2/client"
)

func main() {
	limit := int32(56) // int32 |  (optional)
	offset := int32(56) // int32 |  (optional)
	organizationId := int32(56) // int32 |  (optional)
	sortBy := "sortBy_example" // string |  (optional)
	sortDirection := "sortDirection_example" // string |  (optional)
	search := "search_example" // string |  (optional)
	updatedAt := time.Now() // time.Time |  (optional)
	searchId := "searchId_example" // string |  (optional)
	id := int32(56) // int32 |  (optional)
	backupCredentialId := int32(56) // int32 |  (optional)
	healthy := true // bool |  (optional)
	domainId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ProjectsList(context.Background()).Limit(limit).Offset(offset).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).UpdatedAt(updatedAt).SearchId(searchId).Id(id).BackupCredentialId(backupCredentialId).Healthy(healthy).DomainId(domainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsList`: ProjectsList
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **organizationId** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **updatedAt** | **time.Time** |  | 
 **searchId** | **string** |  | 
 **id** | **int32** |  | 
 **backupCredentialId** | **int32** |  | 
 **healthy** | **bool** |  | 
 **domainId** | **int32** |  | 

### Return type

[**ProjectsList**](ProjectsList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsLockManager

> ProjectsLockManager(ctx).ProjectLockManagerCommand(projectLockManagerCommand).Execute()

Lock/Unlock project

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
	projectLockManagerCommand := *openapiclient.NewProjectLockManagerCommand(int32(123), "Mode_example") // ProjectLockManagerCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.ProjectsLockManager(context.Background()).ProjectLockManagerCommand(projectLockManagerCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsLockManager``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsLockManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectLockManagerCommand** | [**ProjectLockManagerCommand**](ProjectLockManagerCommand.md) |  | 

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


## ProjectsLokiLogs

> []LokiResult ProjectsLokiLogs(ctx).ProjectsLogsCommand(projectsLogsCommand).Execute()

Retrieve loki logs

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
	projectsLogsCommand := *openapiclient.NewProjectsLogsCommand() // ProjectsLogsCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ProjectsLokiLogs(context.Background()).ProjectsLogsCommand(projectsLogsCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsLokiLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsLokiLogs`: []LokiResult
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectsLokiLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsLokiLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectsLogsCommand** | [**ProjectsLogsCommand**](ProjectsLogsCommand.md) |  | 

### Return type

[**[]LokiResult**](LokiResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsMaintenanceManager

> ProjectsMaintenanceManager(ctx).ProjectMaintenanceModeCommand(projectMaintenanceModeCommand).Execute()

Enable/disable project's maintenance mode

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
	projectMaintenanceModeCommand := *openapiclient.NewProjectMaintenanceModeCommand(int32(123), "Mode_example") // ProjectMaintenanceModeCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.ProjectsMaintenanceManager(context.Background()).ProjectMaintenanceModeCommand(projectMaintenanceModeCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsMaintenanceManager``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsMaintenanceManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectMaintenanceModeCommand** | [**ProjectMaintenanceModeCommand**](ProjectMaintenanceModeCommand.md) |  | 

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


## ProjectsMonitoringAlerts

> ProjectMonitoringAlertsDto ProjectsMonitoringAlerts(ctx).ProjectsMonitoringAlertsCommand(projectsMonitoringAlertsCommand).Execute()

Monitoring alerts for project

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
	projectsMonitoringAlertsCommand := *openapiclient.NewProjectsMonitoringAlertsCommand(int32(123)) // ProjectsMonitoringAlertsCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ProjectsMonitoringAlerts(context.Background()).ProjectsMonitoringAlertsCommand(projectsMonitoringAlertsCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsMonitoringAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsMonitoringAlerts`: ProjectMonitoringAlertsDto
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectsMonitoringAlerts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsMonitoringAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectsMonitoringAlertsCommand** | [**ProjectsMonitoringAlertsCommand**](ProjectsMonitoringAlertsCommand.md) |  | 

### Return type

[**ProjectMonitoringAlertsDto**](ProjectMonitoringAlertsDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsPrometheusMetrics

> PrometheusMetricListDto ProjectsPrometheusMetrics(ctx).PrometheusMetricsCommand(prometheusMetricsCommand).Execute()

Prometheus metrics data project

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
	prometheusMetricsCommand := *openapiclient.NewPrometheusMetricsCommand() // PrometheusMetricsCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ProjectsPrometheusMetrics(context.Background()).PrometheusMetricsCommand(prometheusMetricsCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsPrometheusMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsPrometheusMetrics`: PrometheusMetricListDto
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectsPrometheusMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsPrometheusMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **prometheusMetricsCommand** | [**PrometheusMetricsCommand**](PrometheusMetricsCommand.md) |  | 

### Return type

[**PrometheusMetricListDto**](PrometheusMetricListDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsPrometheusMetricsAutocomplete

> []string ProjectsPrometheusMetricsAutocomplete(ctx).PrometheusMetricsAutocompleteCommand(prometheusMetricsAutocompleteCommand).Execute()

Prometheus metrics autocomplete values

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
	prometheusMetricsAutocompleteCommand := *openapiclient.NewPrometheusMetricsAutocompleteCommand(int32(123)) // PrometheusMetricsAutocompleteCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ProjectsPrometheusMetricsAutocomplete(context.Background()).PrometheusMetricsAutocompleteCommand(prometheusMetricsAutocompleteCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsPrometheusMetricsAutocomplete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsPrometheusMetricsAutocomplete`: []string
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectsPrometheusMetricsAutocomplete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsPrometheusMetricsAutocompleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **prometheusMetricsAutocompleteCommand** | [**PrometheusMetricsAutocompleteCommand**](PrometheusMetricsAutocompleteCommand.md) |  | 

### Return type

**[]string**

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsSetAirgapOverride

> ProjectsSetAirgapOverride(ctx).SetProjectAirgapOverrideCommand(setProjectAirgapOverrideCommand).Execute()

Set per-project airgap registry override

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
	setProjectAirgapOverrideCommand := *openapiclient.NewSetProjectAirgapOverrideCommand(int32(123), "OciRegistryHost_example", false) // SetProjectAirgapOverrideCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.ProjectsSetAirgapOverride(context.Background()).SetProjectAirgapOverrideCommand(setProjectAirgapOverrideCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsSetAirgapOverride``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsSetAirgapOverrideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setProjectAirgapOverrideCommand** | [**SetProjectAirgapOverrideCommand**](SetProjectAirgapOverrideCommand.md) |  | 

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


## ProjectsToggleFullSpot

> ProjectsToggleFullSpot(ctx).FullSpotOperationCommand(fullSpotOperationCommand).Execute()

Full spot operations enable/disable

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
	fullSpotOperationCommand := *openapiclient.NewFullSpotOperationCommand(int32(123), "Mode_example") // FullSpotOperationCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.ProjectsToggleFullSpot(context.Background()).FullSpotOperationCommand(fullSpotOperationCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsToggleFullSpot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsToggleFullSpotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fullSpotOperationCommand** | [**FullSpotOperationCommand**](FullSpotOperationCommand.md) |  | 

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


## ProjectsToggleSpotVms

> ProjectsToggleSpotVms(ctx).SpotVmOperationCommand(spotVmOperationCommand).Execute()

Spot vm(s) operations enable/disable

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
	spotVmOperationCommand := *openapiclient.NewSpotVmOperationCommand(int32(123), "Mode_example") // SpotVmOperationCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.ProjectsToggleSpotVms(context.Background()).SpotVmOperationCommand(spotVmOperationCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsToggleSpotVms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsToggleSpotVmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **spotVmOperationCommand** | [**SpotVmOperationCommand**](SpotVmOperationCommand.md) |  | 

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


## ProjectsToggleSpotWorkers

> ProjectsToggleSpotWorkers(ctx).SpotWorkerOperationCommand(spotWorkerOperationCommand).Execute()

Spot worker(s) operations enable/disable

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
	spotWorkerOperationCommand := *openapiclient.NewSpotWorkerOperationCommand(int32(123), "Mode_example") // SpotWorkerOperationCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.ProjectsToggleSpotWorkers(context.Background()).SpotWorkerOperationCommand(spotWorkerOperationCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsToggleSpotWorkers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsToggleSpotWorkersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **spotWorkerOperationCommand** | [**SpotWorkerOperationCommand**](SpotWorkerOperationCommand.md) |  | 

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


## ProjectsVisibility

> ProjectActionVisibilityDto ProjectsVisibility(ctx, projectId).Execute()

Visibility of project actions

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
	resp, r, err := apiClient.ProjectsAPI.ProjectsVisibility(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsVisibility``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsVisibility`: ProjectActionVisibilityDto
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectsVisibility`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsVisibilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectActionVisibilityDto**](ProjectActionVisibilityDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

