# \KubeConfigRoleAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KubeconfigroleList**](KubeConfigRoleAPI.md#KubeconfigroleList) | **Get** /api/v1/kubeconfigrole | Retrieve list of kube config role



## KubeconfigroleList

> KubeConfigRoleResponse KubeconfigroleList(ctx).Search(search).Execute()

Retrieve list of kube config role

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
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubeConfigRoleAPI.KubeconfigroleList(context.Background()).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubeConfigRoleAPI.KubeconfigroleList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubeconfigroleList`: KubeConfigRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `KubeConfigRoleAPI.KubeconfigroleList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubeconfigroleListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** |  | 

### Return type

[**KubeConfigRoleResponse**](KubeConfigRoleResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

