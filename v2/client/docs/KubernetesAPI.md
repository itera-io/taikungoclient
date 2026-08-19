# \KubernetesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KubernetesAddK8sAlert**](KubernetesAPI.md#KubernetesAddK8sAlert) | **Post** /api/v1/kubernetes/alert/{projectId} | Add k8s alert
[**KubernetesAddK8sEvents**](KubernetesAPI.md#KubernetesAddK8sEvents) | **Post** /api/v1/kubernetes/event/{projectId} | Add k8s event
[**KubernetesAlertList**](KubernetesAPI.md#KubernetesAlertList) | **Get** /api/v1/kubernetes/{projectId}/alerts | Retrieve a list of alerts for project
[**KubernetesCli**](KubernetesAPI.md#KubernetesCli) | **Post** /api/v1/kubernetes/cli | Execute k8s web socket namespaced pod
[**KubernetesConfigmapList**](KubernetesAPI.md#KubernetesConfigmapList) | **Get** /api/v1/kubernetes/list/{projectId}/configmap | K8s configmap list
[**KubernetesCreateResource**](KubernetesAPI.md#KubernetesCreateResource) | **Post** /api/v1/kubernetes/create-resource | Create kubernetes resource
[**KubernetesCronjobActions**](KubernetesAPI.md#KubernetesCronjobActions) | **Post** /api/v1/kubernetes/cronjob/actions | Cronjob actions
[**KubernetesDaemonsetActions**](KubernetesAPI.md#KubernetesDaemonsetActions) | **Post** /api/v1/kubernetes/daemonset/actions | Daemonset actions
[**KubernetesDaemonsetList**](KubernetesAPI.md#KubernetesDaemonsetList) | **Get** /api/v1/kubernetes/list/{projectId}/daemonset | K8s daemonset list
[**KubernetesDashboardList**](KubernetesAPI.md#KubernetesDashboardList) | **Get** /api/v1/kubernetes/{projectId}/dashboard | Retrieve a list of crd
[**KubernetesDeleteResource**](KubernetesAPI.md#KubernetesDeleteResource) | **Post** /api/v1/kubernetes/delete-resource | Delete kubernetes resource
[**KubernetesDeploymentActions**](KubernetesAPI.md#KubernetesDeploymentActions) | **Post** /api/v1/kubernetes/deployment/actions | Deployment actions
[**KubernetesDeploymentsList**](KubernetesAPI.md#KubernetesDeploymentsList) | **Get** /api/v1/kubernetes/list/{projectId}/deployments | K8s deployment list
[**KubernetesDescribeResource**](KubernetesAPI.md#KubernetesDescribeResource) | **Post** /api/v1/kubernetes/describe-resource | Describe kubernetes resource
[**KubernetesDownload**](KubernetesAPI.md#KubernetesDownload) | **Get** /api/v1/kubernetes/{projectId}/download | Download kube config file
[**KubernetesDownloadManifest**](KubernetesAPI.md#KubernetesDownloadManifest) | **Post** /api/v1/kubernetes/download-manifest | Download manifest
[**KubernetesExport**](KubernetesAPI.md#KubernetesExport) | **Get** /api/v1/kubernetes/export | Export
[**KubernetesForbiddenNamespaces**](KubernetesAPI.md#KubernetesForbiddenNamespaces) | **Get** /api/v1/kubernetes/forbidden-namespaces/{projectId} | Forbidden namespaces for k8s actions
[**KubernetesGetSupportedList**](KubernetesAPI.md#KubernetesGetSupportedList) | **Get** /api/v1/kubernetes/supported/list | Retrieve Taikun supported kubernetes versions
[**KubernetesIngressClasses**](KubernetesAPI.md#KubernetesIngressClasses) | **Post** /api/v1/kubernetes/ingress-classes | List of ingress classes
[**KubernetesIngressList**](KubernetesAPI.md#KubernetesIngressList) | **Get** /api/v1/kubernetes/list/{projectId}/ingress | K8s ingress list
[**KubernetesInteractiveShell**](KubernetesAPI.md#KubernetesInteractiveShell) | **Post** /api/v1/kubernetes/interactive-shell | Produce interactive shell command
[**KubernetesKubeConfig**](KubernetesAPI.md#KubernetesKubeConfig) | **Get** /api/v1/kubernetes/{projectId}/kubeconfig | Retrieve kube config file
[**KubernetesNamespaceList**](KubernetesAPI.md#KubernetesNamespaceList) | **Get** /api/v1/kubernetes/{projectId}/namespaces | Retrieve kube config file
[**KubernetesNetworkPolicyList**](KubernetesAPI.md#KubernetesNetworkPolicyList) | **Get** /api/v1/kubernetes/list/{projectId}/network-policy | K8s network policy list
[**KubernetesNodeActions**](KubernetesAPI.md#KubernetesNodeActions) | **Post** /api/v1/kubernetes/nodes/actions | Node actions
[**KubernetesNodeList**](KubernetesAPI.md#KubernetesNodeList) | **Get** /api/v1/kubernetes/list/{projectId}/nodes | K8s node list
[**KubernetesOverview**](KubernetesAPI.md#KubernetesOverview) | **Get** /api/v1/kubernetes/overview | Overview kubernetes nodes and pods by organization id
[**KubernetesPatchResource**](KubernetesAPI.md#KubernetesPatchResource) | **Post** /api/v1/kubernetes/patch-resource | Patch kubernetes resource
[**KubernetesPodLogs**](KubernetesAPI.md#KubernetesPodLogs) | **Post** /api/v1/kubernetes/podLogs | Retrieve k8s pod logs
[**KubernetesPodsList**](KubernetesAPI.md#KubernetesPodsList) | **Get** /api/v1/kubernetes/{projectId}/pods-list | K8s pods list
[**KubernetesPvcList**](KubernetesAPI.md#KubernetesPvcList) | **Get** /api/v1/kubernetes/list/{projectId}/pvc | K8s pvc list
[**KubernetesQuota**](KubernetesAPI.md#KubernetesQuota) | **Get** /api/v1/kubernetes/{projectId}/quota | K8s quota usage
[**KubernetesRemovealerts**](KubernetesAPI.md#KubernetesRemovealerts) | **Post** /api/v1/kubernetes/removealerts | Remove k8s alerts
[**KubernetesResources**](KubernetesAPI.md#KubernetesResources) | **Get** /api/v1/kubernetes/{projectId}/resources | Retrieve all total count
[**KubernetesSecretList**](KubernetesAPI.md#KubernetesSecretList) | **Get** /api/v1/kubernetes/list/{projectId}/secret | K8s secret list
[**KubernetesServicesList**](KubernetesAPI.md#KubernetesServicesList) | **Get** /api/v1/kubernetes/list/{projectId}/service | K8s service list
[**KubernetesSilenceManager**](KubernetesAPI.md#KubernetesSilenceManager) | **Post** /api/v1/kubernetes/silencemanager | Silence management for k8s alerts
[**KubernetesStsActions**](KubernetesAPI.md#KubernetesStsActions) | **Post** /api/v1/kubernetes/sts/actions | Stateful set actions
[**KubernetesStsList**](KubernetesAPI.md#KubernetesStsList) | **Get** /api/v1/kubernetes/list/{projectId}/sts | K8s sts list
[**KubernetesUpdateAlert**](KubernetesAPI.md#KubernetesUpdateAlert) | **Put** /api/v1/kubernetes/updatealert/{alertId} | Update k8s alert



## KubernetesAddK8sAlert

> KubernetesAddK8sAlert(ctx, projectId).CreateAlertDto(createAlertDto).Execute()

Add k8s alert

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
	createAlertDto := *openapiclient.NewCreateAlertDto() // CreateAlertDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KubernetesAPI.KubernetesAddK8sAlert(context.Background(), projectId).CreateAlertDto(createAlertDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesAddK8sAlert``: %v\n", err)
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

Other parameters are passed through a pointer to a apiKubernetesAddK8sAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createAlertDto** | [**CreateAlertDto**](CreateAlertDto.md) |  | 

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


## KubernetesAddK8sEvents

> KubernetesAddK8sEvents(ctx, projectId).KubernetesEventCreateDto(kubernetesEventCreateDto).Execute()

Add k8s event

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
	kubernetesEventCreateDto := *openapiclient.NewKubernetesEventCreateDto() // KubernetesEventCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KubernetesAPI.KubernetesAddK8sEvents(context.Background(), projectId).KubernetesEventCreateDto(kubernetesEventCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesAddK8sEvents``: %v\n", err)
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

Other parameters are passed through a pointer to a apiKubernetesAddK8sEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kubernetesEventCreateDto** | [**KubernetesEventCreateDto**](KubernetesEventCreateDto.md) |  | 

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


## KubernetesAlertList

> KubernetesAlertList KubernetesAlertList(ctx, projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Type_(type_).StartDate(startDate).EndDate(endDate).Execute()

Retrieve a list of alerts for project

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
	projectId := int32(56) // int32 | 
	limit := int32(56) // int32 |  (optional)
	offset := int32(56) // int32 |  (optional)
	sortBy := "sortBy_example" // string |  (optional)
	sortDirection := "sortDirection_example" // string |  (optional)
	search := "search_example" // string |  (optional)
	type_ := "type__example" // string |  (optional)
	startDate := time.Now() // time.Time |  (optional)
	endDate := time.Now() // time.Time |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesAlertList(context.Background(), projectId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Type_(type_).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesAlertList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesAlertList`: KubernetesAlertList
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesAlertList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesAlertListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **type_** | **string** |  | 
 **startDate** | **time.Time** |  | 
 **endDate** | **time.Time** |  | 

### Return type

[**KubernetesAlertList**](KubernetesAlertList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesCli

> string KubernetesCli(ctx).KubernetesCliCommand(kubernetesCliCommand).Execute()

Execute k8s web socket namespaced pod

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
	kubernetesCliCommand := *openapiclient.NewKubernetesCliCommand(int32(123), NullableInt32(123)) // KubernetesCliCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesCli(context.Background()).KubernetesCliCommand(kubernetesCliCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesCli``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesCli`: string
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesCli`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesCliRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kubernetesCliCommand** | [**KubernetesCliCommand**](KubernetesCliCommand.md) |  | 

### Return type

**string**

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesConfigmapList

> ConfigmapListDtoCursorStringPaginatedResponse KubernetesConfigmapList(ctx, projectId).Limit(limit).Cursor(cursor).SearchTerm(searchTerm).Execute()

K8s configmap list

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
	limit := int32(56) // int32 |  (optional) (default to 50)
	cursor := "cursor_example" // string |  (optional)
	searchTerm := "searchTerm_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesConfigmapList(context.Background(), projectId).Limit(limit).Cursor(cursor).SearchTerm(searchTerm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesConfigmapList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesConfigmapList`: ConfigmapListDtoCursorStringPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesConfigmapList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesConfigmapListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | [default to 50]
 **cursor** | **string** |  | 
 **searchTerm** | **string** |  | 

### Return type

[**ConfigmapListDtoCursorStringPaginatedResponse**](ConfigmapListDtoCursorStringPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesCreateResource

> KubernetesCreateResource(ctx).CreateKubernetesResourceCommand(createKubernetesResourceCommand).Execute()

Create kubernetes resource

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
	createKubernetesResourceCommand := *openapiclient.NewCreateKubernetesResourceCommand(int32(123), "Yaml_example") // CreateKubernetesResourceCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KubernetesAPI.KubernetesCreateResource(context.Background()).CreateKubernetesResourceCommand(createKubernetesResourceCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesCreateResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesCreateResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createKubernetesResourceCommand** | [**CreateKubernetesResourceCommand**](CreateKubernetesResourceCommand.md) |  | 

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


## KubernetesCronjobActions

> KubernetesActionResponse KubernetesCronjobActions(ctx).CronjobActionCommand(cronjobActionCommand).Execute()

Cronjob actions

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
	cronjobActionCommand := *openapiclient.NewCronjobActionCommand(int32(123), []openapiclient.KubernetesCronjobActionRequest{*openapiclient.NewKubernetesCronjobActionRequest("Name_example", "Namespace_example")}, openapiclient.ECronJobAction("Suspend")) // CronjobActionCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesCronjobActions(context.Background()).CronjobActionCommand(cronjobActionCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesCronjobActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesCronjobActions`: KubernetesActionResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesCronjobActions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesCronjobActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cronjobActionCommand** | [**CronjobActionCommand**](CronjobActionCommand.md) |  | 

### Return type

[**KubernetesActionResponse**](KubernetesActionResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDaemonsetActions

> KubernetesDaemonsetActions(ctx).DaemonsetActionCommand(daemonsetActionCommand).Execute()

Daemonset actions

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
	daemonsetActionCommand := *openapiclient.NewDaemonsetActionCommand(int32(123), "Name_example", "Namespace_example", openapiclient.EDaemonSetAction("Restart")) // DaemonsetActionCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KubernetesAPI.KubernetesDaemonsetActions(context.Background()).DaemonsetActionCommand(daemonsetActionCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesDaemonsetActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDaemonsetActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **daemonsetActionCommand** | [**DaemonsetActionCommand**](DaemonsetActionCommand.md) |  | 

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


## KubernetesDaemonsetList

> DaemonsetListDtoCursorStringPaginatedResponse KubernetesDaemonsetList(ctx, projectId).Limit(limit).Cursor(cursor).SearchTerm(searchTerm).Execute()

K8s daemonset list

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
	limit := int32(56) // int32 |  (optional) (default to 50)
	cursor := "cursor_example" // string |  (optional)
	searchTerm := "searchTerm_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesDaemonsetList(context.Background(), projectId).Limit(limit).Cursor(cursor).SearchTerm(searchTerm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesDaemonsetList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesDaemonsetList`: DaemonsetListDtoCursorStringPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesDaemonsetList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDaemonsetListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | [default to 50]
 **cursor** | **string** |  | 
 **searchTerm** | **string** |  | 

### Return type

[**DaemonsetListDtoCursorStringPaginatedResponse**](DaemonsetListDtoCursorStringPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDashboardList

> KubernetesDashboardDto KubernetesDashboardList(ctx, projectId).Execute()

Retrieve a list of crd

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
	resp, r, err := apiClient.KubernetesAPI.KubernetesDashboardList(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesDashboardList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesDashboardList`: KubernetesDashboardDto
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesDashboardList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDashboardListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KubernetesDashboardDto**](KubernetesDashboardDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDeleteResource

> KubernetesActionResponse KubernetesDeleteResource(ctx).DeleteKubernetesResourceCommand(deleteKubernetesResourceCommand).Execute()

Delete kubernetes resource

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
	deleteKubernetesResourceCommand := *openapiclient.NewDeleteKubernetesResourceCommand(int32(123), openapiclient.EKubernetesResource("None"), []openapiclient.KubernetesActionRequest{*openapiclient.NewKubernetesActionRequest("Name_example")}) // DeleteKubernetesResourceCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesDeleteResource(context.Background()).DeleteKubernetesResourceCommand(deleteKubernetesResourceCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesDeleteResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesDeleteResource`: KubernetesActionResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesDeleteResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDeleteResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteKubernetesResourceCommand** | [**DeleteKubernetesResourceCommand**](DeleteKubernetesResourceCommand.md) |  | 

### Return type

[**KubernetesActionResponse**](KubernetesActionResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDeploymentActions

> KubernetesDeploymentActions(ctx).DeploymentActionCommand(deploymentActionCommand).Execute()

Deployment actions

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
	deploymentActionCommand := *openapiclient.NewDeploymentActionCommand(int32(123), "Name_example", "Namespace_example", openapiclient.EDeploymentAction("Pause")) // DeploymentActionCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KubernetesAPI.KubernetesDeploymentActions(context.Background()).DeploymentActionCommand(deploymentActionCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesDeploymentActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDeploymentActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deploymentActionCommand** | [**DeploymentActionCommand**](DeploymentActionCommand.md) |  | 

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


## KubernetesDeploymentsList

> DeploymentListDtoCursorStringPaginatedResponse KubernetesDeploymentsList(ctx, projectId).Limit(limit).Cursor(cursor).SearchTerm(searchTerm).Execute()

K8s deployment list

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
	limit := int32(56) // int32 |  (optional) (default to 50)
	cursor := "cursor_example" // string |  (optional)
	searchTerm := "searchTerm_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesDeploymentsList(context.Background(), projectId).Limit(limit).Cursor(cursor).SearchTerm(searchTerm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesDeploymentsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesDeploymentsList`: DeploymentListDtoCursorStringPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesDeploymentsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDeploymentsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | [default to 50]
 **cursor** | **string** |  | 
 **searchTerm** | **string** |  | 

### Return type

[**DeploymentListDtoCursorStringPaginatedResponse**](DeploymentListDtoCursorStringPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDescribeResource

> string KubernetesDescribeResource(ctx).DescribeKubernetesResourceCommand(describeKubernetesResourceCommand).Execute()

Describe kubernetes resource

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
	describeKubernetesResourceCommand := *openapiclient.NewDescribeKubernetesResourceCommand(int32(123), "Name_example", openapiclient.EKubernetesResource("None")) // DescribeKubernetesResourceCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesDescribeResource(context.Background()).DescribeKubernetesResourceCommand(describeKubernetesResourceCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesDescribeResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesDescribeResource`: string
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesDescribeResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDescribeResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **describeKubernetesResourceCommand** | [**DescribeKubernetesResourceCommand**](DescribeKubernetesResourceCommand.md) |  | 

### Return type

**string**

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDownload

> map[string]interface{} KubernetesDownload(ctx, projectId).Execute()

Download kube config file

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
	resp, r, err := apiClient.KubernetesAPI.KubernetesDownload(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesDownload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesDownload`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesDownload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesDownloadManifest

> CsvExporter KubernetesDownloadManifest(ctx).DownloadKubernetesResourceCommand(downloadKubernetesResourceCommand).Execute()

Download manifest

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
	downloadKubernetesResourceCommand := *openapiclient.NewDownloadKubernetesResourceCommand(int32(123), "Name_example", openapiclient.EKubernetesResource("None")) // DownloadKubernetesResourceCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesDownloadManifest(context.Background()).DownloadKubernetesResourceCommand(downloadKubernetesResourceCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesDownloadManifest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesDownloadManifest`: CsvExporter
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesDownloadManifest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesDownloadManifestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **downloadKubernetesResourceCommand** | [**DownloadKubernetesResourceCommand**](DownloadKubernetesResourceCommand.md) |  | 

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


## KubernetesExport

> CsvExporter KubernetesExport(ctx).ProjectId(projectId).Execute()

Export

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
	resp, r, err := apiClient.KubernetesAPI.KubernetesExport(context.Background()).ProjectId(projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesExport`: CsvExporter
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesExport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **int32** |  | 

### Return type

[**CsvExporter**](CsvExporter.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesForbiddenNamespaces

> []string KubernetesForbiddenNamespaces(ctx, projectId).Execute()

Forbidden namespaces for k8s actions

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
	resp, r, err := apiClient.KubernetesAPI.KubernetesForbiddenNamespaces(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesForbiddenNamespaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesForbiddenNamespaces`: []string
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesForbiddenNamespaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesForbiddenNamespacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesGetSupportedList

> []KubernetesVersionListDto KubernetesGetSupportedList(ctx).Execute()

Retrieve Taikun supported kubernetes versions

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
	resp, r, err := apiClient.KubernetesAPI.KubernetesGetSupportedList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesGetSupportedList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesGetSupportedList`: []KubernetesVersionListDto
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesGetSupportedList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesGetSupportedListRequest struct via the builder pattern


### Return type

[**[]KubernetesVersionListDto**](KubernetesVersionListDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesIngressClasses

> []string KubernetesIngressClasses(ctx).Config(config).CloudId(cloudId).Region(region).ImportClusterName(importClusterName).ResourceGroupName(resourceGroupName).Execute()

List of ingress classes

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesIngressClasses(context.Background()).Config(config).CloudId(cloudId).Region(region).ImportClusterName(importClusterName).ResourceGroupName(resourceGroupName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesIngressClasses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesIngressClasses`: []string
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesIngressClasses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesIngressClassesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **config** | ***os.File** |  | 
 **cloudId** | **int32** |  | 
 **region** | **string** |  | 
 **importClusterName** | **string** |  | 
 **resourceGroupName** | **string** |  | 

### Return type

**[]string**

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesIngressList

> IngressListDtoCursorStringPaginatedResponse KubernetesIngressList(ctx, projectId).Limit(limit).Cursor(cursor).SearchTerm(searchTerm).Execute()

K8s ingress list

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
	limit := int32(56) // int32 |  (optional) (default to 50)
	cursor := "cursor_example" // string |  (optional)
	searchTerm := "searchTerm_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesIngressList(context.Background(), projectId).Limit(limit).Cursor(cursor).SearchTerm(searchTerm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesIngressList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesIngressList`: IngressListDtoCursorStringPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesIngressList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesIngressListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | [default to 50]
 **cursor** | **string** |  | 
 **searchTerm** | **string** |  | 

### Return type

[**IngressListDtoCursorStringPaginatedResponse**](IngressListDtoCursorStringPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesInteractiveShell

> KubernetesInteractiveShellDto KubernetesInteractiveShell(ctx).InteractiveShellSendCommand(interactiveShellSendCommand).Execute()

Produce interactive shell command

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
	interactiveShellSendCommand := *openapiclient.NewInteractiveShellSendCommand(int32(123), "Token_example", NullableInt32(123)) // InteractiveShellSendCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesInteractiveShell(context.Background()).InteractiveShellSendCommand(interactiveShellSendCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesInteractiveShell``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesInteractiveShell`: KubernetesInteractiveShellDto
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesInteractiveShell`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesInteractiveShellRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **interactiveShellSendCommand** | [**InteractiveShellSendCommand**](InteractiveShellSendCommand.md) |  | 

### Return type

[**KubernetesInteractiveShellDto**](KubernetesInteractiveShellDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesKubeConfig

> KubeConfigResponse KubernetesKubeConfig(ctx, projectId).Execute()

Retrieve kube config file

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
	resp, r, err := apiClient.KubernetesAPI.KubernetesKubeConfig(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesKubeConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesKubeConfig`: KubeConfigResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesKubeConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesKubeConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KubeConfigResponse**](KubeConfigResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesNamespaceList

> []string KubernetesNamespaceList(ctx, projectId).Installation(installation).Execute()

Retrieve kube config file

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
	installation := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesNamespaceList(context.Background(), projectId).Installation(installation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesNamespaceList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesNamespaceList`: []string
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesNamespaceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesNamespaceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **installation** | **bool** |  | 

### Return type

**[]string**

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesNetworkPolicyList

> NetworkPolicyListDtoCursorStringPaginatedResponse KubernetesNetworkPolicyList(ctx, projectId).Limit(limit).Cursor(cursor).SearchTerm(searchTerm).Execute()

K8s network policy list

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
	limit := int32(56) // int32 |  (optional) (default to 50)
	cursor := "cursor_example" // string |  (optional)
	searchTerm := "searchTerm_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesNetworkPolicyList(context.Background(), projectId).Limit(limit).Cursor(cursor).SearchTerm(searchTerm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesNetworkPolicyList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesNetworkPolicyList`: NetworkPolicyListDtoCursorStringPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesNetworkPolicyList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesNetworkPolicyListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | [default to 50]
 **cursor** | **string** |  | 
 **searchTerm** | **string** |  | 

### Return type

[**NetworkPolicyListDtoCursorStringPaginatedResponse**](NetworkPolicyListDtoCursorStringPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesNodeActions

> KubernetesNodeActions(ctx).NodeActionCommand(nodeActionCommand).Execute()

Node actions

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
	nodeActionCommand := *openapiclient.NewNodeActionCommand(int32(123), "Name_example", openapiclient.ENodeAction("Cordon")) // NodeActionCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KubernetesAPI.KubernetesNodeActions(context.Background()).NodeActionCommand(nodeActionCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesNodeActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesNodeActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nodeActionCommand** | [**NodeActionCommand**](NodeActionCommand.md) |  | 

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


## KubernetesNodeList

> NodeListDtoCursorStringPaginatedResponse KubernetesNodeList(ctx, projectId).Limit(limit).Cursor(cursor).SearchTerm(searchTerm).Execute()

K8s node list

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
	limit := int32(56) // int32 |  (optional) (default to 50)
	cursor := "cursor_example" // string |  (optional)
	searchTerm := "searchTerm_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesNodeList(context.Background(), projectId).Limit(limit).Cursor(cursor).SearchTerm(searchTerm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesNodeList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesNodeList`: NodeListDtoCursorStringPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesNodeList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesNodeListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | [default to 50]
 **cursor** | **string** |  | 
 **searchTerm** | **string** |  | 

### Return type

[**NodeListDtoCursorStringPaginatedResponse**](NodeListDtoCursorStringPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesOverview

> []KubernetesOverviewDto KubernetesOverview(ctx).OrganizationId(organizationId).Execute()

Overview kubernetes nodes and pods by organization id

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
	organizationId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesOverview(context.Background()).OrganizationId(organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesOverview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesOverview`: []KubernetesOverviewDto
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesOverview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 

### Return type

[**[]KubernetesOverviewDto**](KubernetesOverviewDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesPatchResource

> KubernetesPatchResource(ctx).PatchKubernetesResourceCommand(patchKubernetesResourceCommand).Execute()

Patch kubernetes resource

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
	patchKubernetesResourceCommand := *openapiclient.NewPatchKubernetesResourceCommand(int32(123), "Yaml_example", "Name_example") // PatchKubernetesResourceCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KubernetesAPI.KubernetesPatchResource(context.Background()).PatchKubernetesResourceCommand(patchKubernetesResourceCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesPatchResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesPatchResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchKubernetesResourceCommand** | [**PatchKubernetesResourceCommand**](PatchKubernetesResourceCommand.md) |  | 

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


## KubernetesPodLogs

> string KubernetesPodLogs(ctx).KubernetesPodLogsCommand(kubernetesPodLogsCommand).Execute()

Retrieve k8s pod logs

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
	kubernetesPodLogsCommand := *openapiclient.NewKubernetesPodLogsCommand(int32(123), "Name_example", "Namespace_example", "Container_example") // KubernetesPodLogsCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesPodLogs(context.Background()).KubernetesPodLogsCommand(kubernetesPodLogsCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesPodLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesPodLogs`: string
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesPodLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesPodLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kubernetesPodLogsCommand** | [**KubernetesPodLogsCommand**](KubernetesPodLogsCommand.md) |  | 

### Return type

**string**

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesPodsList

> PodsListDtoCursorStringPaginatedResponse KubernetesPodsList(ctx, projectId).Limit(limit).Cursor(cursor).SearchTerm(searchTerm).Execute()

K8s pods list

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
	limit := int32(56) // int32 |  (optional) (default to 50)
	cursor := "cursor_example" // string |  (optional)
	searchTerm := "searchTerm_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesPodsList(context.Background(), projectId).Limit(limit).Cursor(cursor).SearchTerm(searchTerm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesPodsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesPodsList`: PodsListDtoCursorStringPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesPodsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesPodsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | [default to 50]
 **cursor** | **string** |  | 
 **searchTerm** | **string** |  | 

### Return type

[**PodsListDtoCursorStringPaginatedResponse**](PodsListDtoCursorStringPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesPvcList

> PvcListDtoCursorStringPaginatedResponse KubernetesPvcList(ctx, projectId).Limit(limit).Cursor(cursor).SearchTerm(searchTerm).Execute()

K8s pvc list

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
	limit := int32(56) // int32 |  (optional) (default to 50)
	cursor := "cursor_example" // string |  (optional)
	searchTerm := "searchTerm_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesPvcList(context.Background(), projectId).Limit(limit).Cursor(cursor).SearchTerm(searchTerm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesPvcList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesPvcList`: PvcListDtoCursorStringPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesPvcList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesPvcListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | [default to 50]
 **cursor** | **string** |  | 
 **searchTerm** | **string** |  | 

### Return type

[**PvcListDtoCursorStringPaginatedResponse**](PvcListDtoCursorStringPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesQuota

> KubernetesQuotaListDto KubernetesQuota(ctx, projectId).Execute()

K8s quota usage

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
	resp, r, err := apiClient.KubernetesAPI.KubernetesQuota(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesQuota``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesQuota`: KubernetesQuotaListDto
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesQuota`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KubernetesQuotaListDto**](KubernetesQuotaListDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesRemovealerts

> KubernetesRemovealerts(ctx).DeleteAlertCommand(deleteAlertCommand).Execute()

Remove k8s alerts

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
	deleteAlertCommand := *openapiclient.NewDeleteAlertCommand(int32(123)) // DeleteAlertCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KubernetesAPI.KubernetesRemovealerts(context.Background()).DeleteAlertCommand(deleteAlertCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesRemovealerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesRemovealertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteAlertCommand** | [**DeleteAlertCommand**](DeleteAlertCommand.md) |  | 

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


## KubernetesResources

> KubernetesResourcesDto KubernetesResources(ctx, projectId).Execute()

Retrieve all total count

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
	resp, r, err := apiClient.KubernetesAPI.KubernetesResources(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesResources`: KubernetesResourcesDto
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KubernetesResourcesDto**](KubernetesResourcesDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesSecretList

> SecretListDtoCursorStringPaginatedResponse KubernetesSecretList(ctx, projectId).Limit(limit).Cursor(cursor).SearchTerm(searchTerm).Execute()

K8s secret list

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
	limit := int32(56) // int32 |  (optional) (default to 50)
	cursor := "cursor_example" // string |  (optional)
	searchTerm := "searchTerm_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesSecretList(context.Background(), projectId).Limit(limit).Cursor(cursor).SearchTerm(searchTerm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesSecretList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesSecretList`: SecretListDtoCursorStringPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesSecretList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesSecretListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | [default to 50]
 **cursor** | **string** |  | 
 **searchTerm** | **string** |  | 

### Return type

[**SecretListDtoCursorStringPaginatedResponse**](SecretListDtoCursorStringPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesServicesList

> ServiceListDtoCursorStringPaginatedResponse KubernetesServicesList(ctx, projectId).Limit(limit).Cursor(cursor).SearchTerm(searchTerm).Execute()

K8s service list

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
	limit := int32(56) // int32 |  (optional) (default to 50)
	cursor := "cursor_example" // string |  (optional)
	searchTerm := "searchTerm_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesServicesList(context.Background(), projectId).Limit(limit).Cursor(cursor).SearchTerm(searchTerm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesServicesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesServicesList`: ServiceListDtoCursorStringPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesServicesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesServicesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | [default to 50]
 **cursor** | **string** |  | 
 **searchTerm** | **string** |  | 

### Return type

[**ServiceListDtoCursorStringPaginatedResponse**](ServiceListDtoCursorStringPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesSilenceManager

> KubernetesSilenceManager(ctx).SilenceOperationsCommand(silenceOperationsCommand).Execute()

Silence management for k8s alerts

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
	silenceOperationsCommand := *openapiclient.NewSilenceOperationsCommand(int32(123), "Mode_example", "Reason_example") // SilenceOperationsCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KubernetesAPI.KubernetesSilenceManager(context.Background()).SilenceOperationsCommand(silenceOperationsCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesSilenceManager``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesSilenceManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **silenceOperationsCommand** | [**SilenceOperationsCommand**](SilenceOperationsCommand.md) |  | 

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


## KubernetesStsActions

> KubernetesStsActions(ctx).StsActionCommand(stsActionCommand).Execute()

Stateful set actions

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
	stsActionCommand := *openapiclient.NewStsActionCommand(int32(123), "Name_example", "Namespace_example", openapiclient.EStsAction("Restart")) // StsActionCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KubernetesAPI.KubernetesStsActions(context.Background()).StsActionCommand(stsActionCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesStsActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesStsActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stsActionCommand** | [**StsActionCommand**](StsActionCommand.md) |  | 

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


## KubernetesStsList

> StatefulSetListDtoCursorStringPaginatedResponse KubernetesStsList(ctx, projectId).Limit(limit).Cursor(cursor).SearchTerm(searchTerm).Execute()

K8s sts list

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
	limit := int32(56) // int32 |  (optional) (default to 50)
	cursor := "cursor_example" // string |  (optional)
	searchTerm := "searchTerm_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesStsList(context.Background(), projectId).Limit(limit).Cursor(cursor).SearchTerm(searchTerm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesStsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesStsList`: StatefulSetListDtoCursorStringPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesStsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesStsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | [default to 50]
 **cursor** | **string** |  | 
 **searchTerm** | **string** |  | 

### Return type

[**StatefulSetListDtoCursorStringPaginatedResponse**](StatefulSetListDtoCursorStringPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesUpdateAlert

> KubernetesUpdateAlert(ctx, alertId).UpdateKubernetesAlertDto(updateKubernetesAlertDto).Execute()

Update k8s alert

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
	alertId := int32(56) // int32 | 
	updateKubernetesAlertDto := *openapiclient.NewUpdateKubernetesAlertDto() // UpdateKubernetesAlertDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KubernetesAPI.KubernetesUpdateAlert(context.Background(), alertId).UpdateKubernetesAlertDto(updateKubernetesAlertDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesUpdateAlert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesUpdateAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateKubernetesAlertDto** | [**UpdateKubernetesAlertDto**](UpdateKubernetesAlertDto.md) |  | 

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

