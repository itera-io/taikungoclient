# \ProjectDeploymentAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectDeploymentCommit**](ProjectDeploymentAPI.md#ProjectDeploymentCommit) | **Post** /api/v1/project-deployment/commit | Commit
[**ProjectDeploymentCommitVm**](ProjectDeploymentAPI.md#ProjectDeploymentCommitVm) | **Post** /api/v1/project-deployment/commit-vm | Commit
[**ProjectDeploymentCompleted**](ProjectDeploymentAPI.md#ProjectDeploymentCompleted) | **Post** /api/v1/project-deployment/completed | Update project fields
[**ProjectDeploymentDelete**](ProjectDeploymentAPI.md#ProjectDeploymentDelete) | **Post** /api/v1/project-deployment/delete | Delete
[**ProjectDeploymentDeleteVmDisks**](ProjectDeploymentAPI.md#ProjectDeploymentDeleteVmDisks) | **Post** /api/v1/project-deployment/delete-vm-disks | Delete vm disks
[**ProjectDeploymentDeleteVms**](ProjectDeploymentAPI.md#ProjectDeploymentDeleteVms) | **Post** /api/v1/project-deployment/delete-vms | Delete vms
[**ProjectDeploymentDisableAi**](ProjectDeploymentAPI.md#ProjectDeploymentDisableAi) | **Post** /api/v1/project-deployment/disable-ai | Disable ai
[**ProjectDeploymentDisableBackup**](ProjectDeploymentAPI.md#ProjectDeploymentDisableBackup) | **Post** /api/v1/project-deployment/disable-backup | Disable backup
[**ProjectDeploymentDisableMonitoring**](ProjectDeploymentAPI.md#ProjectDeploymentDisableMonitoring) | **Post** /api/v1/project-deployment/disable-monitoring | Disable monitoring
[**ProjectDeploymentDisableOpa**](ProjectDeploymentAPI.md#ProjectDeploymentDisableOpa) | **Post** /api/v1/project-deployment/disable-opa | Disable opa
[**ProjectDeploymentEnableAi**](ProjectDeploymentAPI.md#ProjectDeploymentEnableAi) | **Post** /api/v1/project-deployment/enable-ai | Enable ai
[**ProjectDeploymentEnableBackup**](ProjectDeploymentAPI.md#ProjectDeploymentEnableBackup) | **Post** /api/v1/project-deployment/enable-backup | Enable backup
[**ProjectDeploymentEnableMonitoring**](ProjectDeploymentAPI.md#ProjectDeploymentEnableMonitoring) | **Post** /api/v1/project-deployment/enable-monitoring | Enable monitoring
[**ProjectDeploymentEnableOpa**](ProjectDeploymentAPI.md#ProjectDeploymentEnableOpa) | **Post** /api/v1/project-deployment/enable-opa | Enable opa
[**ProjectDeploymentImportCluster**](ProjectDeploymentAPI.md#ProjectDeploymentImportCluster) | **Post** /api/v1/project-deployment/import-cluster | Import cluster
[**ProjectDeploymentRepair**](ProjectDeploymentAPI.md#ProjectDeploymentRepair) | **Post** /api/v1/project-deployment/repair | Repair
[**ProjectDeploymentRepairVm**](ProjectDeploymentAPI.md#ProjectDeploymentRepairVm) | **Post** /api/v1/project-deployment/repair-vm | Repair Vm
[**ProjectDeploymentUpdate**](ProjectDeploymentAPI.md#ProjectDeploymentUpdate) | **Post** /api/v1/project-deployment/update | Update stage of project
[**ProjectDeploymentUpgrade**](ProjectDeploymentAPI.md#ProjectDeploymentUpgrade) | **Post** /api/v1/project-deployment/upgrade/{projectId} | Upgrade the project&#39;s Kubernetes to the next available version. Project must be READY.



## ProjectDeploymentCommit

> ProjectDeploymentCommit(ctx).ProjectDeploymentCommitCommand(projectDeploymentCommitCommand).Execute()

Commit

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
	projectDeploymentCommitCommand := *openapiclient.NewProjectDeploymentCommitCommand(int32(123)) // ProjectDeploymentCommitCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectDeploymentAPI.ProjectDeploymentCommit(context.Background()).ProjectDeploymentCommitCommand(projectDeploymentCommitCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectDeploymentAPI.ProjectDeploymentCommit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectDeploymentCommitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectDeploymentCommitCommand** | [**ProjectDeploymentCommitCommand**](ProjectDeploymentCommitCommand.md) |  | 

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


## ProjectDeploymentCommitVm

> ProjectDeploymentCommitVm(ctx).DeploymentCommitVmCommand(deploymentCommitVmCommand).Execute()

Commit

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
	deploymentCommitVmCommand := *openapiclient.NewDeploymentCommitVmCommand(int32(123)) // DeploymentCommitVmCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectDeploymentAPI.ProjectDeploymentCommitVm(context.Background()).DeploymentCommitVmCommand(deploymentCommitVmCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectDeploymentAPI.ProjectDeploymentCommitVm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectDeploymentCommitVmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deploymentCommitVmCommand** | [**DeploymentCommitVmCommand**](DeploymentCommitVmCommand.md) |  | 

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


## ProjectDeploymentCompleted

> ProjectDeploymentCompleted(ctx).DeploymentCompletedCommand(deploymentCompletedCommand).Execute()

Update project fields

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
	deploymentCompletedCommand := *openapiclient.NewDeploymentCompletedCommand(int32(123), "Result_example") // DeploymentCompletedCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectDeploymentAPI.ProjectDeploymentCompleted(context.Background()).DeploymentCompletedCommand(deploymentCompletedCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectDeploymentAPI.ProjectDeploymentCompleted``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectDeploymentCompletedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deploymentCompletedCommand** | [**DeploymentCompletedCommand**](DeploymentCompletedCommand.md) |  | 

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


## ProjectDeploymentDelete

> ProjectDeploymentDelete(ctx).ProjectDeploymentDeleteServersCommand(projectDeploymentDeleteServersCommand).Execute()

Delete

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
	projectDeploymentDeleteServersCommand := *openapiclient.NewProjectDeploymentDeleteServersCommand() // ProjectDeploymentDeleteServersCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectDeploymentAPI.ProjectDeploymentDelete(context.Background()).ProjectDeploymentDeleteServersCommand(projectDeploymentDeleteServersCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectDeploymentAPI.ProjectDeploymentDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectDeploymentDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectDeploymentDeleteServersCommand** | [**ProjectDeploymentDeleteServersCommand**](ProjectDeploymentDeleteServersCommand.md) |  | 

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


## ProjectDeploymentDeleteVmDisks

> ProjectDeploymentDeleteVmDisks(ctx).DeleteVmDiskCommand(deleteVmDiskCommand).Execute()

Delete vm disks

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
	deleteVmDiskCommand := *openapiclient.NewDeleteVmDiskCommand(int32(123), []int32{int32(123)}) // DeleteVmDiskCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectDeploymentAPI.ProjectDeploymentDeleteVmDisks(context.Background()).DeleteVmDiskCommand(deleteVmDiskCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectDeploymentAPI.ProjectDeploymentDeleteVmDisks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectDeploymentDeleteVmDisksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteVmDiskCommand** | [**DeleteVmDiskCommand**](DeleteVmDiskCommand.md) |  | 

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


## ProjectDeploymentDeleteVms

> ProjectDeploymentDeleteVms(ctx).ProjectDeploymentDeleteVmsCommand(projectDeploymentDeleteVmsCommand).Execute()

Delete vms

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
	projectDeploymentDeleteVmsCommand := *openapiclient.NewProjectDeploymentDeleteVmsCommand() // ProjectDeploymentDeleteVmsCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectDeploymentAPI.ProjectDeploymentDeleteVms(context.Background()).ProjectDeploymentDeleteVmsCommand(projectDeploymentDeleteVmsCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectDeploymentAPI.ProjectDeploymentDeleteVms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectDeploymentDeleteVmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectDeploymentDeleteVmsCommand** | [**ProjectDeploymentDeleteVmsCommand**](ProjectDeploymentDeleteVmsCommand.md) |  | 

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


## ProjectDeploymentDisableAi

> ProjectDeploymentDisableAi(ctx).DeploymentDisableAiCommand(deploymentDisableAiCommand).Execute()

Disable ai

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
	deploymentDisableAiCommand := *openapiclient.NewDeploymentDisableAiCommand(int32(123)) // DeploymentDisableAiCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectDeploymentAPI.ProjectDeploymentDisableAi(context.Background()).DeploymentDisableAiCommand(deploymentDisableAiCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectDeploymentAPI.ProjectDeploymentDisableAi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectDeploymentDisableAiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deploymentDisableAiCommand** | [**DeploymentDisableAiCommand**](DeploymentDisableAiCommand.md) |  | 

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


## ProjectDeploymentDisableBackup

> ProjectDeploymentDisableBackup(ctx).DeploymentDisableBackupCommand(deploymentDisableBackupCommand).Execute()

Disable backup

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
	deploymentDisableBackupCommand := *openapiclient.NewDeploymentDisableBackupCommand(int32(123)) // DeploymentDisableBackupCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectDeploymentAPI.ProjectDeploymentDisableBackup(context.Background()).DeploymentDisableBackupCommand(deploymentDisableBackupCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectDeploymentAPI.ProjectDeploymentDisableBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectDeploymentDisableBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deploymentDisableBackupCommand** | [**DeploymentDisableBackupCommand**](DeploymentDisableBackupCommand.md) |  | 

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


## ProjectDeploymentDisableMonitoring

> ProjectDeploymentDisableMonitoring(ctx).DeploymentDisableMonitoringCommand(deploymentDisableMonitoringCommand).Execute()

Disable monitoring

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
	deploymentDisableMonitoringCommand := *openapiclient.NewDeploymentDisableMonitoringCommand(int32(123)) // DeploymentDisableMonitoringCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectDeploymentAPI.ProjectDeploymentDisableMonitoring(context.Background()).DeploymentDisableMonitoringCommand(deploymentDisableMonitoringCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectDeploymentAPI.ProjectDeploymentDisableMonitoring``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectDeploymentDisableMonitoringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deploymentDisableMonitoringCommand** | [**DeploymentDisableMonitoringCommand**](DeploymentDisableMonitoringCommand.md) |  | 

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


## ProjectDeploymentDisableOpa

> ProjectDeploymentDisableOpa(ctx).DeploymentDisableOpaCommand(deploymentDisableOpaCommand).Execute()

Disable opa

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
	deploymentDisableOpaCommand := *openapiclient.NewDeploymentDisableOpaCommand(int32(123)) // DeploymentDisableOpaCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectDeploymentAPI.ProjectDeploymentDisableOpa(context.Background()).DeploymentDisableOpaCommand(deploymentDisableOpaCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectDeploymentAPI.ProjectDeploymentDisableOpa``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectDeploymentDisableOpaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deploymentDisableOpaCommand** | [**DeploymentDisableOpaCommand**](DeploymentDisableOpaCommand.md) |  | 

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


## ProjectDeploymentEnableAi

> ProjectDeploymentEnableAi(ctx).DeploymentEnableAiCommand(deploymentEnableAiCommand).Execute()

Enable ai

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
	deploymentEnableAiCommand := *openapiclient.NewDeploymentEnableAiCommand(int32(123), int32(123)) // DeploymentEnableAiCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectDeploymentAPI.ProjectDeploymentEnableAi(context.Background()).DeploymentEnableAiCommand(deploymentEnableAiCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectDeploymentAPI.ProjectDeploymentEnableAi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectDeploymentEnableAiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deploymentEnableAiCommand** | [**DeploymentEnableAiCommand**](DeploymentEnableAiCommand.md) |  | 

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


## ProjectDeploymentEnableBackup

> ProjectDeploymentEnableBackup(ctx).DeploymentEnableBackupCommand(deploymentEnableBackupCommand).Execute()

Enable backup

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
	deploymentEnableBackupCommand := *openapiclient.NewDeploymentEnableBackupCommand(int32(123), int32(123)) // DeploymentEnableBackupCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectDeploymentAPI.ProjectDeploymentEnableBackup(context.Background()).DeploymentEnableBackupCommand(deploymentEnableBackupCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectDeploymentAPI.ProjectDeploymentEnableBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectDeploymentEnableBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deploymentEnableBackupCommand** | [**DeploymentEnableBackupCommand**](DeploymentEnableBackupCommand.md) |  | 

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


## ProjectDeploymentEnableMonitoring

> ProjectDeploymentEnableMonitoring(ctx).DeploymentEnableMonitoringCommand(deploymentEnableMonitoringCommand).Execute()

Enable monitoring

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
	deploymentEnableMonitoringCommand := *openapiclient.NewDeploymentEnableMonitoringCommand(int32(123)) // DeploymentEnableMonitoringCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectDeploymentAPI.ProjectDeploymentEnableMonitoring(context.Background()).DeploymentEnableMonitoringCommand(deploymentEnableMonitoringCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectDeploymentAPI.ProjectDeploymentEnableMonitoring``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectDeploymentEnableMonitoringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deploymentEnableMonitoringCommand** | [**DeploymentEnableMonitoringCommand**](DeploymentEnableMonitoringCommand.md) |  | 

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


## ProjectDeploymentEnableOpa

> ProjectDeploymentEnableOpa(ctx).DeploymentOpaEnableCommand(deploymentOpaEnableCommand).Execute()

Enable opa

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
	deploymentOpaEnableCommand := *openapiclient.NewDeploymentOpaEnableCommand(int32(123), int32(123)) // DeploymentOpaEnableCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectDeploymentAPI.ProjectDeploymentEnableOpa(context.Background()).DeploymentOpaEnableCommand(deploymentOpaEnableCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectDeploymentAPI.ProjectDeploymentEnableOpa``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectDeploymentEnableOpaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deploymentOpaEnableCommand** | [**DeploymentOpaEnableCommand**](DeploymentOpaEnableCommand.md) |  | 

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


## ProjectDeploymentImportCluster

> ProjectDeploymentImportCluster(ctx).Name(name).Config(config).IsTaikunIngressController(isTaikunIngressController).IsExistingIngressController(isExistingIngressController).IngressClass(ingressClass).IngressHost(ingressHost).Continent(continent).ImportType(importType).OrganizationId(organizationId).CloudId(cloudId).ImportClusterName(importClusterName).ResourceGroupName(resourceGroupName).Region(region).Execute()

Import cluster

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
	name := "name_example" // string | 
	config := os.NewFile(1234, "some_file") // *os.File |  (optional)
	isTaikunIngressController := true // bool |  (optional)
	isExistingIngressController := true // bool |  (optional)
	ingressClass := "ingressClass_example" // string |  (optional)
	ingressHost := "ingressHost_example" // string |  (optional)
	continent := "continent_example" // string |  (optional)
	importType := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} |  (optional)
	organizationId := int32(56) // int32 |  (optional)
	cloudId := int32(56) // int32 |  (optional)
	importClusterName := "importClusterName_example" // string |  (optional)
	resourceGroupName := "resourceGroupName_example" // string |  (optional)
	region := "region_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectDeploymentAPI.ProjectDeploymentImportCluster(context.Background()).Name(name).Config(config).IsTaikunIngressController(isTaikunIngressController).IsExistingIngressController(isExistingIngressController).IngressClass(ingressClass).IngressHost(ingressHost).Continent(continent).ImportType(importType).OrganizationId(organizationId).CloudId(cloudId).ImportClusterName(importClusterName).ResourceGroupName(resourceGroupName).Region(region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectDeploymentAPI.ProjectDeploymentImportCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectDeploymentImportClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
 **config** | ***os.File** |  | 
 **isTaikunIngressController** | **bool** |  | 
 **isExistingIngressController** | **bool** |  | 
 **ingressClass** | **string** |  | 
 **ingressHost** | **string** |  | 
 **continent** | **string** |  | 
 **importType** | **map[string]interface{}** |  | 
 **organizationId** | **int32** |  | 
 **cloudId** | **int32** |  | 
 **importClusterName** | **string** |  | 
 **resourceGroupName** | **string** |  | 
 **region** | **string** |  | 

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


## ProjectDeploymentRepair

> ProjectDeploymentRepair(ctx).ProjectDeploymentRepairCommand(projectDeploymentRepairCommand).Execute()

Repair

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
	projectDeploymentRepairCommand := *openapiclient.NewProjectDeploymentRepairCommand(int32(123)) // ProjectDeploymentRepairCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectDeploymentAPI.ProjectDeploymentRepair(context.Background()).ProjectDeploymentRepairCommand(projectDeploymentRepairCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectDeploymentAPI.ProjectDeploymentRepair``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectDeploymentRepairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectDeploymentRepairCommand** | [**ProjectDeploymentRepairCommand**](ProjectDeploymentRepairCommand.md) |  | 

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


## ProjectDeploymentRepairVm

> ProjectDeploymentRepairVm(ctx).ProjectDeploymentRepairVmCommand(projectDeploymentRepairVmCommand).Execute()

Repair Vm

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
	projectDeploymentRepairVmCommand := *openapiclient.NewProjectDeploymentRepairVmCommand(int32(123)) // ProjectDeploymentRepairVmCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectDeploymentAPI.ProjectDeploymentRepairVm(context.Background()).ProjectDeploymentRepairVmCommand(projectDeploymentRepairVmCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectDeploymentAPI.ProjectDeploymentRepairVm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectDeploymentRepairVmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectDeploymentRepairVmCommand** | [**ProjectDeploymentRepairVmCommand**](ProjectDeploymentRepairVmCommand.md) |  | 

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


## ProjectDeploymentUpdate

> ProjectDeploymentUpdate(ctx).UpdateStageCommand(updateStageCommand).Execute()

Update stage of project

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
	updateStageCommand := *openapiclient.NewUpdateStageCommand(int32(123), "Stage_example") // UpdateStageCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectDeploymentAPI.ProjectDeploymentUpdate(context.Background()).UpdateStageCommand(updateStageCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectDeploymentAPI.ProjectDeploymentUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectDeploymentUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateStageCommand** | [**UpdateStageCommand**](UpdateStageCommand.md) |  | 

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


## ProjectDeploymentUpgrade

> ProjectDeploymentUpgrade(ctx, projectId).Execute()

Upgrade the project's Kubernetes to the next available version. Project must be READY.

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
	r, err := apiClient.ProjectDeploymentAPI.ProjectDeploymentUpgrade(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectDeploymentAPI.ProjectDeploymentUpgrade``: %v\n", err)
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

Other parameters are passed through a pointer to a apiProjectDeploymentUpgradeRequest struct via the builder pattern


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

