# \RobotAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RobotChecker**](RobotAPI.md#RobotChecker) | **Post** /api/v1/robot/checker | Dry run robot user create
[**RobotCreate**](RobotAPI.md#RobotCreate) | **Post** /api/v1/robot/create | Add new robot user
[**RobotDelete**](RobotAPI.md#RobotDelete) | **Delete** /api/v1/robot/{id} | Delete robot user
[**RobotDetails**](RobotAPI.md#RobotDetails) | **Get** /api/v1/robot/details | Retrieve robot user info
[**RobotGetDetails**](RobotAPI.md#RobotGetDetails) | **Get** /api/v1/robot/{id} | Retrieve robot user details by ID
[**RobotList**](RobotAPI.md#RobotList) | **Get** /api/v1/robot/list | Retrieve robot users
[**RobotRegenerate**](RobotAPI.md#RobotRegenerate) | **Post** /api/v1/robot/regenerate | Regenerate new token credentials for robot user
[**RobotScopeList**](RobotAPI.md#RobotScopeList) | **Get** /api/v1/robot/scope-list | Retrieve available scope list flavors
[**RobotStatus**](RobotAPI.md#RobotStatus) | **Post** /api/v1/robot/status | Activate/Deactivate robot user
[**RobotUpdate**](RobotAPI.md#RobotUpdate) | **Put** /api/v1/robot/update | Update robot user&#39;s details
[**RobotUpdateScope**](RobotAPI.md#RobotUpdateScope) | **Put** /api/v1/robot/update-scope | Update robot user&#39;s scope



## RobotChecker

> RobotChecker(ctx).CheckerRobotCommand(checkerRobotCommand).Execute()

Dry run robot user create

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
	checkerRobotCommand := *openapiclient.NewCheckerRobotCommand(NullableInt32(123), NullableInt32(123), time.Now(), "Name_example", "Description_example") // CheckerRobotCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RobotAPI.RobotChecker(context.Background()).CheckerRobotCommand(checkerRobotCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RobotAPI.RobotChecker``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRobotCheckerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkerRobotCommand** | [**CheckerRobotCommand**](CheckerRobotCommand.md) |  | 

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


## RobotCreate

> CreateRobotUserResponse RobotCreate(ctx).CreateRobotUserCommand(createRobotUserCommand).Execute()

Add new robot user

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
	createRobotUserCommand := *openapiclient.NewCreateRobotUserCommand(NullableInt32(123), NullableInt32(123), "Name_example", "Description_example", time.Now(), []string{"Scopes_example"}, []string{"Ips_example"}) // CreateRobotUserCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RobotAPI.RobotCreate(context.Background()).CreateRobotUserCommand(createRobotUserCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RobotAPI.RobotCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RobotCreate`: CreateRobotUserResponse
	fmt.Fprintf(os.Stdout, "Response from `RobotAPI.RobotCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRobotCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRobotUserCommand** | [**CreateRobotUserCommand**](CreateRobotUserCommand.md) |  | 

### Return type

[**CreateRobotUserResponse**](CreateRobotUserResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RobotDelete

> RobotDelete(ctx, id).Execute()

Delete robot user

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RobotAPI.RobotDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RobotAPI.RobotDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRobotDeleteRequest struct via the builder pattern


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


## RobotDetails

> RobotUsersListDto RobotDetails(ctx).Execute()

Retrieve robot user info

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
	resp, r, err := apiClient.RobotAPI.RobotDetails(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RobotAPI.RobotDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RobotDetails`: RobotUsersListDto
	fmt.Fprintf(os.Stdout, "Response from `RobotAPI.RobotDetails`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRobotDetailsRequest struct via the builder pattern


### Return type

[**RobotUsersListDto**](RobotUsersListDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RobotGetDetails

> RobotUsersListDto RobotGetDetails(ctx, id).Execute()

Retrieve robot user details by ID

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RobotAPI.RobotGetDetails(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RobotAPI.RobotGetDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RobotGetDetails`: RobotUsersListDto
	fmt.Fprintf(os.Stdout, "Response from `RobotAPI.RobotGetDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRobotGetDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RobotUsersListDto**](RobotUsersListDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RobotList

> RobotUsersList RobotList(ctx).DomainId(domainId).OrganizationId(organizationId).Search(search).SearchId(searchId).Limit(limit).Offset(offset).Execute()

Retrieve robot users

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
	domainId := int32(56) // int32 |  (optional)
	organizationId := int32(56) // int32 |  (optional)
	search := "search_example" // string |  (optional)
	searchId := "searchId_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)
	offset := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RobotAPI.RobotList(context.Background()).DomainId(domainId).OrganizationId(organizationId).Search(search).SearchId(searchId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RobotAPI.RobotList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RobotList`: RobotUsersList
	fmt.Fprintf(os.Stdout, "Response from `RobotAPI.RobotList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRobotListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainId** | **int32** |  | 
 **organizationId** | **int32** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **limit** | **int32** |  | [default to 50]
 **offset** | **int32** |  | [default to 0]

### Return type

[**RobotUsersList**](RobotUsersList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RobotRegenerate

> CreateRobotUserResponse RobotRegenerate(ctx).RegenerateRobotTokenCommand(regenerateRobotTokenCommand).Execute()

Regenerate new token credentials for robot user

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
	regenerateRobotTokenCommand := *openapiclient.NewRegenerateRobotTokenCommand("Id_example", time.Now()) // RegenerateRobotTokenCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RobotAPI.RobotRegenerate(context.Background()).RegenerateRobotTokenCommand(regenerateRobotTokenCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RobotAPI.RobotRegenerate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RobotRegenerate`: CreateRobotUserResponse
	fmt.Fprintf(os.Stdout, "Response from `RobotAPI.RobotRegenerate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRobotRegenerateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **regenerateRobotTokenCommand** | [**RegenerateRobotTokenCommand**](RegenerateRobotTokenCommand.md) |  | 

### Return type

[**CreateRobotUserResponse**](CreateRobotUserResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RobotScopeList

> []RobotUserScopeData RobotScopeList(ctx).GlobalRobot(globalRobot).Execute()

Retrieve available scope list flavors

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
	globalRobot := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RobotAPI.RobotScopeList(context.Background()).GlobalRobot(globalRobot).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RobotAPI.RobotScopeList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RobotScopeList`: []RobotUserScopeData
	fmt.Fprintf(os.Stdout, "Response from `RobotAPI.RobotScopeList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRobotScopeListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **globalRobot** | **bool** |  | [default to false]

### Return type

[**[]RobotUserScopeData**](RobotUserScopeData.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RobotStatus

> RobotStatus(ctx).RobotUserStatusManagementCommand(robotUserStatusManagementCommand).Execute()

Activate/Deactivate robot user

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
	robotUserStatusManagementCommand := *openapiclient.NewRobotUserStatusManagementCommand("Id_example", "Mode_example") // RobotUserStatusManagementCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RobotAPI.RobotStatus(context.Background()).RobotUserStatusManagementCommand(robotUserStatusManagementCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RobotAPI.RobotStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRobotStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **robotUserStatusManagementCommand** | [**RobotUserStatusManagementCommand**](RobotUserStatusManagementCommand.md) |  | 

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


## RobotUpdate

> RobotUpdate(ctx).EditRobotUserCommand(editRobotUserCommand).Execute()

Update robot user's details

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
	editRobotUserCommand := *openapiclient.NewEditRobotUserCommand("Id_example", "Name_example", "Description_example", time.Now(), []string{"Ips_example"}) // EditRobotUserCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RobotAPI.RobotUpdate(context.Background()).EditRobotUserCommand(editRobotUserCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RobotAPI.RobotUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRobotUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **editRobotUserCommand** | [**EditRobotUserCommand**](EditRobotUserCommand.md) |  | 

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


## RobotUpdateScope

> RobotUpdateScope(ctx).UpdateRobotScopeCommand(updateRobotScopeCommand).Execute()

Update robot user's scope

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
	updateRobotScopeCommand := *openapiclient.NewUpdateRobotScopeCommand("Id_example", []string{"Scopes_example"}) // UpdateRobotScopeCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RobotAPI.RobotUpdateScope(context.Background()).UpdateRobotScopeCommand(updateRobotScopeCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RobotAPI.RobotUpdateScope``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRobotUpdateScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateRobotScopeCommand** | [**UpdateRobotScopeCommand**](UpdateRobotScopeCommand.md) |  | 

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

