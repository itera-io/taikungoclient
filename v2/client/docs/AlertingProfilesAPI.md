# \AlertingProfilesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AlertingprofilesAssignEmail**](AlertingProfilesAPI.md#AlertingprofilesAssignEmail) | **Put** /api/v1/alertingprofiles/assignemails/{id} | Assign Alerting emails
[**AlertingprofilesAssignWebhooks**](AlertingProfilesAPI.md#AlertingprofilesAssignWebhooks) | **Put** /api/v1/alertingprofiles/assignwebhooks/{id} | Assign Alerting webhooks
[**AlertingprofilesAttach**](AlertingProfilesAPI.md#AlertingprofilesAttach) | **Post** /api/v1/alertingprofiles/attach | Attach alerting profile to project
[**AlertingprofilesCreate**](AlertingProfilesAPI.md#AlertingprofilesCreate) | **Post** /api/v1/alertingprofiles/create | Add Alerting profile
[**AlertingprofilesDelete**](AlertingProfilesAPI.md#AlertingprofilesDelete) | **Delete** /api/v1/alertingprofiles/{id} | Remove Alerting profiles by Id
[**AlertingprofilesDetach**](AlertingProfilesAPI.md#AlertingprofilesDetach) | **Post** /api/v1/alertingprofiles/detach | Detach alerting profile from project
[**AlertingprofilesDropdown**](AlertingProfilesAPI.md#AlertingprofilesDropdown) | **Get** /api/v1/alertingprofiles/list | Retrieve all Alerting profiles for organization
[**AlertingprofilesEdit**](AlertingProfilesAPI.md#AlertingprofilesEdit) | **Post** /api/v1/alertingprofiles/edit | Update Alerting profile
[**AlertingprofilesList**](AlertingProfilesAPI.md#AlertingprofilesList) | **Get** /api/v1/alertingprofiles | Retrieve all Alerting profiles
[**AlertingprofilesLockManager**](AlertingProfilesAPI.md#AlertingprofilesLockManager) | **Post** /api/v1/alertingprofiles/lockmanager | Lock/Unlock Alerting profiles
[**AlertingprofilesVerify**](AlertingProfilesAPI.md#AlertingprofilesVerify) | **Post** /api/v1/alertingprofiles/verifywebhook | Verify webhook endpoint



## AlertingprofilesAssignEmail

> AlertingprofilesAssignEmail(ctx, id).AlertingEmailDto(alertingEmailDto).Execute()

Assign Alerting emails

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
	alertingEmailDto := []openapiclient.AlertingEmailDto{*openapiclient.NewAlertingEmailDto("Email_example")} // []AlertingEmailDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AlertingProfilesAPI.AlertingprofilesAssignEmail(context.Background(), id).AlertingEmailDto(alertingEmailDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertingProfilesAPI.AlertingprofilesAssignEmail``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAlertingprofilesAssignEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alertingEmailDto** | [**[]AlertingEmailDto**](AlertingEmailDto.md) |  | 

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


## AlertingprofilesAssignWebhooks

> AlertingprofilesAssignWebhooks(ctx, id).AlertingWebhookDto(alertingWebhookDto).Execute()

Assign Alerting webhooks

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
	alertingWebhookDto := []openapiclient.AlertingWebhookDto{*openapiclient.NewAlertingWebhookDto()} // []AlertingWebhookDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AlertingProfilesAPI.AlertingprofilesAssignWebhooks(context.Background(), id).AlertingWebhookDto(alertingWebhookDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertingProfilesAPI.AlertingprofilesAssignWebhooks``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAlertingprofilesAssignWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alertingWebhookDto** | [**[]AlertingWebhookDto**](AlertingWebhookDto.md) |  | 

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


## AlertingprofilesAttach

> AlertingprofilesAttach(ctx).AttachDetachAlertingProfileCommand(attachDetachAlertingProfileCommand).Execute()

Attach alerting profile to project

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
	attachDetachAlertingProfileCommand := *openapiclient.NewAttachDetachAlertingProfileCommand(int32(123), NullableInt32(123)) // AttachDetachAlertingProfileCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AlertingProfilesAPI.AlertingprofilesAttach(context.Background()).AttachDetachAlertingProfileCommand(attachDetachAlertingProfileCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertingProfilesAPI.AlertingprofilesAttach``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertingprofilesAttachRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attachDetachAlertingProfileCommand** | [**AttachDetachAlertingProfileCommand**](AttachDetachAlertingProfileCommand.md) |  | 

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


## AlertingprofilesCreate

> ApiResponse AlertingprofilesCreate(ctx).CreateAlertingProfileCommand(createAlertingProfileCommand).Execute()

Add Alerting profile

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
	createAlertingProfileCommand := *openapiclient.NewCreateAlertingProfileCommand() // CreateAlertingProfileCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertingProfilesAPI.AlertingprofilesCreate(context.Background()).CreateAlertingProfileCommand(createAlertingProfileCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertingProfilesAPI.AlertingprofilesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlertingprofilesCreate`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AlertingProfilesAPI.AlertingprofilesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertingprofilesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAlertingProfileCommand** | [**CreateAlertingProfileCommand**](CreateAlertingProfileCommand.md) |  | 

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


## AlertingprofilesDelete

> AlertingprofilesDelete(ctx, id).Execute()

Remove Alerting profiles by Id

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
	r, err := apiClient.AlertingProfilesAPI.AlertingprofilesDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertingProfilesAPI.AlertingprofilesDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAlertingprofilesDeleteRequest struct via the builder pattern


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


## AlertingprofilesDetach

> AlertingprofilesDetach(ctx).AttachDetachAlertingProfileCommand(attachDetachAlertingProfileCommand).Execute()

Detach alerting profile from project

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
	attachDetachAlertingProfileCommand := *openapiclient.NewAttachDetachAlertingProfileCommand(int32(123), NullableInt32(123)) // AttachDetachAlertingProfileCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AlertingProfilesAPI.AlertingprofilesDetach(context.Background()).AttachDetachAlertingProfileCommand(attachDetachAlertingProfileCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertingProfilesAPI.AlertingprofilesDetach``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertingprofilesDetachRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attachDetachAlertingProfileCommand** | [**AttachDetachAlertingProfileCommand**](AttachDetachAlertingProfileCommand.md) |  | 

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


## AlertingprofilesDropdown

> []CommonDropdownDto AlertingprofilesDropdown(ctx).OrganizationId(organizationId).Search(search).Execute()

Retrieve all Alerting profiles for organization

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertingProfilesAPI.AlertingprofilesDropdown(context.Background()).OrganizationId(organizationId).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertingProfilesAPI.AlertingprofilesDropdown``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlertingprofilesDropdown`: []CommonDropdownDto
	fmt.Fprintf(os.Stdout, "Response from `AlertingProfilesAPI.AlertingprofilesDropdown`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertingprofilesDropdownRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 
 **search** | **string** |  | 

### Return type

[**[]CommonDropdownDto**](CommonDropdownDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertingprofilesEdit

> ApiResponse AlertingprofilesEdit(ctx).UpdateAlertingProfileCommand(updateAlertingProfileCommand).Execute()

Update Alerting profile

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
	updateAlertingProfileCommand := *openapiclient.NewUpdateAlertingProfileCommand() // UpdateAlertingProfileCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertingProfilesAPI.AlertingprofilesEdit(context.Background()).UpdateAlertingProfileCommand(updateAlertingProfileCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertingProfilesAPI.AlertingprofilesEdit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlertingprofilesEdit`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AlertingProfilesAPI.AlertingprofilesEdit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertingprofilesEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateAlertingProfileCommand** | [**UpdateAlertingProfileCommand**](UpdateAlertingProfileCommand.md) |  | 

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


## AlertingprofilesList

> AlertingProfilesList AlertingprofilesList(ctx).OrganizationId(organizationId).DomainId(domainId).SortBy(sortBy).SortDirection(sortDirection).Search(search).Id(id).SearchId(searchId).Limit(limit).Offset(offset).Execute()

Retrieve all Alerting profiles

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
	sortBy := "sortBy_example" // string |  (optional)
	sortDirection := "sortDirection_example" // string |  (optional)
	search := "search_example" // string |  (optional)
	id := int32(56) // int32 |  (optional)
	searchId := "searchId_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)
	offset := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertingProfilesAPI.AlertingprofilesList(context.Background()).OrganizationId(organizationId).DomainId(domainId).SortBy(sortBy).SortDirection(sortDirection).Search(search).Id(id).SearchId(searchId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertingProfilesAPI.AlertingprofilesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlertingprofilesList`: AlertingProfilesList
	fmt.Fprintf(os.Stdout, "Response from `AlertingProfilesAPI.AlertingprofilesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertingprofilesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 
 **domainId** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **id** | **int32** |  | 
 **searchId** | **string** |  | 
 **limit** | **int32** |  | [default to 50]
 **offset** | **int32** |  | [default to 0]

### Return type

[**AlertingProfilesList**](AlertingProfilesList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertingprofilesLockManager

> AlertingprofilesLockManager(ctx).AlertingProfilesLockManagerCommand(alertingProfilesLockManagerCommand).Execute()

Lock/Unlock Alerting profiles

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
	alertingProfilesLockManagerCommand := *openapiclient.NewAlertingProfilesLockManagerCommand(int32(123), "Mode_example") // AlertingProfilesLockManagerCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AlertingProfilesAPI.AlertingprofilesLockManager(context.Background()).AlertingProfilesLockManagerCommand(alertingProfilesLockManagerCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertingProfilesAPI.AlertingprofilesLockManager``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertingprofilesLockManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **alertingProfilesLockManagerCommand** | [**AlertingProfilesLockManagerCommand**](AlertingProfilesLockManagerCommand.md) |  | 

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


## AlertingprofilesVerify

> AlertingprofilesVerify(ctx).VerifyWebhookCommand(verifyWebhookCommand).Execute()

Verify webhook endpoint

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
	verifyWebhookCommand := *openapiclient.NewVerifyWebhookCommand("Url_example") // VerifyWebhookCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AlertingProfilesAPI.AlertingprofilesVerify(context.Background()).VerifyWebhookCommand(verifyWebhookCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertingProfilesAPI.AlertingprofilesVerify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertingprofilesVerifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **verifyWebhookCommand** | [**VerifyWebhookCommand**](VerifyWebhookCommand.md) |  | 

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

