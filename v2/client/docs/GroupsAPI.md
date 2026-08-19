# \GroupsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsAddOrganizations**](GroupsAPI.md#GroupsAddOrganizations) | **Post** /api/v1/groups/{id}/organizations | Add organizations to group
[**GroupsAddUsers**](GroupsAPI.md#GroupsAddUsers) | **Post** /api/v1/groups/{id}/users | Add users to group
[**GroupsCheckDuplicateEntity**](GroupsAPI.md#GroupsCheckDuplicateEntity) | **Post** /api/v1/groups/check-duplicate-entity | Check duplicate entity
[**GroupsCreate**](GroupsAPI.md#GroupsCreate) | **Post** /api/v1/groups/create | Create group
[**GroupsDelete**](GroupsAPI.md#GroupsDelete) | **Delete** /api/v1/groups/{id} | Delete group
[**GroupsDeleteOrganizations**](GroupsAPI.md#GroupsDeleteOrganizations) | **Post** /api/v1/groups/delete-organizations | Remove organizations from group
[**GroupsDeleteUsers**](GroupsAPI.md#GroupsDeleteUsers) | **Post** /api/v1/groups/delete-users | Remove users from group
[**GroupsList**](GroupsAPI.md#GroupsList) | **Get** /api/v1/groups | Group list for domain
[**GroupsUpdate**](GroupsAPI.md#GroupsUpdate) | **Put** /api/v1/groups/update/{id} | Update group
[**GroupsUpdateGroupOrganization**](GroupsAPI.md#GroupsUpdateGroupOrganization) | **Put** /api/v1/groups/{groupId}/organizations/{organizationId} | Update organization in group



## GroupsAddOrganizations

> GroupsAddOrganizations(ctx, id).CreateGroupOrganizationDto(createGroupOrganizationDto).Execute()

Add organizations to group

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
	createGroupOrganizationDto := []openapiclient.CreateGroupOrganizationDto{*openapiclient.NewCreateGroupOrganizationDto(int32(123), openapiclient.AccessLevelRoles("Member"))} // []CreateGroupOrganizationDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupsAPI.GroupsAddOrganizations(context.Background(), id).CreateGroupOrganizationDto(createGroupOrganizationDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GroupsAddOrganizations``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGroupsAddOrganizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createGroupOrganizationDto** | [**[]CreateGroupOrganizationDto**](CreateGroupOrganizationDto.md) |  | 

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


## GroupsAddUsers

> GroupsAddUsers(ctx, id).CreateGroupUserDto(createGroupUserDto).Execute()

Add users to group

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
	createGroupUserDto := []openapiclient.CreateGroupUserDto{*openapiclient.NewCreateGroupUserDto("Id_example")} // []CreateGroupUserDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupsAPI.GroupsAddUsers(context.Background(), id).CreateGroupUserDto(createGroupUserDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GroupsAddUsers``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGroupsAddUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createGroupUserDto** | [**[]CreateGroupUserDto**](CreateGroupUserDto.md) |  | 

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


## GroupsCheckDuplicateEntity

> GroupsCheckDuplicateEntity(ctx).CheckDuplicateGroupCommand(checkDuplicateGroupCommand).Execute()

Check duplicate entity

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
	checkDuplicateGroupCommand := *openapiclient.NewCheckDuplicateGroupCommand(NullableInt32(123), "Name_example") // CheckDuplicateGroupCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupsAPI.GroupsCheckDuplicateEntity(context.Background()).CheckDuplicateGroupCommand(checkDuplicateGroupCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GroupsCheckDuplicateEntity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGroupsCheckDuplicateEntityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkDuplicateGroupCommand** | [**CheckDuplicateGroupCommand**](CheckDuplicateGroupCommand.md) |  | 

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


## GroupsCreate

> int32 GroupsCreate(ctx).CreateGroupCommand(createGroupCommand).Execute()

Create group

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
	createGroupCommand := *openapiclient.NewCreateGroupCommand("Name_example", int32(123)) // CreateGroupCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GroupsCreate(context.Background()).CreateGroupCommand(createGroupCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GroupsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GroupsCreate`: int32
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GroupsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGroupsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createGroupCommand** | [**CreateGroupCommand**](CreateGroupCommand.md) |  | 

### Return type

**int32**

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsDelete

> GroupsDelete(ctx, id).Execute()

Delete group

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
	r, err := apiClient.GroupsAPI.GroupsDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GroupsDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGroupsDeleteRequest struct via the builder pattern


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


## GroupsDeleteOrganizations

> GroupsDeleteOrganizations(ctx).DeleteOrganizationFromGroupCommand(deleteOrganizationFromGroupCommand).Execute()

Remove organizations from group

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
	deleteOrganizationFromGroupCommand := *openapiclient.NewDeleteOrganizationFromGroupCommand(int32(123), []int32{int32(123)}) // DeleteOrganizationFromGroupCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupsAPI.GroupsDeleteOrganizations(context.Background()).DeleteOrganizationFromGroupCommand(deleteOrganizationFromGroupCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GroupsDeleteOrganizations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGroupsDeleteOrganizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteOrganizationFromGroupCommand** | [**DeleteOrganizationFromGroupCommand**](DeleteOrganizationFromGroupCommand.md) |  | 

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


## GroupsDeleteUsers

> GroupsDeleteUsers(ctx).DeleteUserFromGroupCommand(deleteUserFromGroupCommand).Execute()

Remove users from group

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
	deleteUserFromGroupCommand := *openapiclient.NewDeleteUserFromGroupCommand(int32(123), []string{"Users_example"}) // DeleteUserFromGroupCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupsAPI.GroupsDeleteUsers(context.Background()).DeleteUserFromGroupCommand(deleteUserFromGroupCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GroupsDeleteUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGroupsDeleteUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteUserFromGroupCommand** | [**DeleteUserFromGroupCommand**](DeleteUserFromGroupCommand.md) |  | 

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


## GroupsList

> GroupList GroupsList(ctx).DomainId(domainId).Search(search).Execute()

Group list for domain

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
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GroupsList(context.Background()).DomainId(domainId).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GroupsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GroupsList`: GroupList
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GroupsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGroupsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainId** | **int32** |  | 
 **search** | **string** |  | 

### Return type

[**GroupList**](GroupList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsUpdate

> GroupsUpdate(ctx, id).UpdateGroupDto(updateGroupDto).Execute()

Update group

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
	updateGroupDto := *openapiclient.NewUpdateGroupDto() // UpdateGroupDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupsAPI.GroupsUpdate(context.Background(), id).UpdateGroupDto(updateGroupDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GroupsUpdate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGroupsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateGroupDto** | [**UpdateGroupDto**](UpdateGroupDto.md) |  | 

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


## GroupsUpdateGroupOrganization

> GroupsUpdateGroupOrganization(ctx, groupId, organizationId).UpdateGroupOrganizationDto(updateGroupOrganizationDto).Execute()

Update organization in group

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
	groupId := int32(56) // int32 | 
	organizationId := int32(56) // int32 | 
	updateGroupOrganizationDto := *openapiclient.NewUpdateGroupOrganizationDto("Role_example") // UpdateGroupOrganizationDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupsAPI.GroupsUpdateGroupOrganization(context.Background(), groupId, organizationId).UpdateGroupOrganizationDto(updateGroupOrganizationDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GroupsUpdateGroupOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int32** |  | 
**organizationId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsUpdateGroupOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateGroupOrganizationDto** | [**UpdateGroupOrganizationDto**](UpdateGroupOrganizationDto.md) |  | 

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

