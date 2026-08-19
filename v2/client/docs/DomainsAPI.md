# \DomainsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DomainsAddDomainAdmin**](DomainsAPI.md#DomainsAddDomainAdmin) | **Post** /api/v1/domains/add-domain-admin | Add domain admin to domain
[**DomainsCheckDuplicateEntity**](DomainsAPI.md#DomainsCheckDuplicateEntity) | **Post** /api/v1/domains/check-duplicate-entity | Check duplicate entity
[**DomainsCreate**](DomainsAPI.md#DomainsCreate) | **Post** /api/v1/domains/create | Create domain
[**DomainsDelete**](DomainsAPI.md#DomainsDelete) | **Delete** /api/v1/domains/{id} | Delete domain
[**DomainsDetails**](DomainsAPI.md#DomainsDetails) | **Get** /api/v1/domains/{domainId} | Get domain details by domainId
[**DomainsDisable2faManagement**](DomainsAPI.md#DomainsDisable2faManagement) | **Post** /api/v1/domains/disable-2fa-management | Disable 2fa management
[**DomainsDomainGroupDetails**](DomainsAPI.md#DomainsDomainGroupDetails) | **Get** /api/v1/domains/{domainId}/groups/{id} | Get group details by domainId
[**DomainsDomainGroupsDropdown**](DomainsAPI.md#DomainsDomainGroupsDropdown) | **Get** /api/v1/domains/{domainId}/groups/dropdown | Retrieve domain groups dropdown list by cursor-based pagination
[**DomainsDomainGroupsOffsetBasedDropdown**](DomainsAPI.md#DomainsDomainGroupsOffsetBasedDropdown) | **Get** /api/v1/domains/{domainId}/groups/offset-based/dropdown | Retrieve domain groups dropdown list by offset-based pagination
[**DomainsDomainOffsetBasedOrganizationsWithGroup**](DomainsAPI.md#DomainsDomainOffsetBasedOrganizationsWithGroup) | **Get** /api/v1/domains/{domainId}/organizations/offset-based/with/group | Retrieve domain organizations with group list by offset-based pagination
[**DomainsDomainOrganizationDetails**](DomainsAPI.md#DomainsDomainOrganizationDetails) | **Get** /api/v1/domains/{domainId}/organizations/{id} | Get organization details by domainId
[**DomainsDomainOrganizationsAvailable**](DomainsAPI.md#DomainsDomainOrganizationsAvailable) | **Get** /api/v1/domains/{domainId}/organizations/available | Retrieve available organizations for adding to a group
[**DomainsDomainOrganizationsWithGroup**](DomainsAPI.md#DomainsDomainOrganizationsWithGroup) | **Get** /api/v1/domains/{domainId}/organizations/with/group | Retrieve domain organizations with group list by cursor-based pagination
[**DomainsDomainProjectDetails**](DomainsAPI.md#DomainsDomainProjectDetails) | **Get** /api/v1/domains/{domainId}/projects/{id} | Get project details by domainId
[**DomainsDomainProjectsDropdown**](DomainsAPI.md#DomainsDomainProjectsDropdown) | **Get** /api/v1/domains/{domainId}/projects/dropdown | Retrieve domain projects dropdown list by cursor-based pagination
[**DomainsDomainProjectsOffsetBasedDropdown**](DomainsAPI.md#DomainsDomainProjectsOffsetBasedDropdown) | **Get** /api/v1/domains/{domainId}/projects/offset-based/dropdown | Retrieve domain projects dropdown list by offset-based pagination
[**DomainsDomainUserDetails**](DomainsAPI.md#DomainsDomainUserDetails) | **Get** /api/v1/domains/{domainId}/users/{id} | Get user details by domainId
[**DomainsDomainUserDropdown**](DomainsAPI.md#DomainsDomainUserDropdown) | **Get** /api/v1/domains/{domainId}/user/dropdown | Retrieve domain users dropdown list by cursor-based pagination
[**DomainsDomainUserOffsetBasedDropdown**](DomainsAPI.md#DomainsDomainUserOffsetBasedDropdown) | **Get** /api/v1/domains/{domainId}/user/offset-based/dropdown | Retrieve domain users dropdown list by offset-based pagination
[**DomainsDomainUsersAvailable**](DomainsAPI.md#DomainsDomainUsersAvailable) | **Get** /api/v1/domains/{domainId}/users/available | Retrieve available users for adding to a group
[**DomainsEnable2faManagement**](DomainsAPI.md#DomainsEnable2faManagement) | **Post** /api/v1/domains/enable-2fa-management | Enable 2fa management
[**DomainsListDomains**](DomainsAPI.md#DomainsListDomains) | **Get** /api/v1/domains | Retrieve domains list
[**DomainsTransferOwnership**](DomainsAPI.md#DomainsTransferOwnership) | **Post** /api/v1/domains/transfer-ownership | Transfer ownership
[**DomainsUpdate**](DomainsAPI.md#DomainsUpdate) | **Post** /api/v1/domains/update | Update domain



## DomainsAddDomainAdmin

> DomainsAddDomainAdmin(ctx).AddDomainAdminCommand(addDomainAdminCommand).Execute()

Add domain admin to domain

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
	addDomainAdminCommand := *openapiclient.NewAddDomainAdminCommand("UserId_example") // AddDomainAdminCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DomainsAPI.DomainsAddDomainAdmin(context.Background()).AddDomainAdminCommand(addDomainAdminCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DomainsAddDomainAdmin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDomainsAddDomainAdminRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addDomainAdminCommand** | [**AddDomainAdminCommand**](AddDomainAdminCommand.md) |  | 

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


## DomainsCheckDuplicateEntity

> DomainsCheckDuplicateEntity(ctx).CheckDuplicateDomainCommand(checkDuplicateDomainCommand).Execute()

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
	checkDuplicateDomainCommand := *openapiclient.NewCheckDuplicateDomainCommand("Name_example") // CheckDuplicateDomainCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DomainsAPI.DomainsCheckDuplicateEntity(context.Background()).CheckDuplicateDomainCommand(checkDuplicateDomainCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DomainsCheckDuplicateEntity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDomainsCheckDuplicateEntityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkDuplicateDomainCommand** | [**CheckDuplicateDomainCommand**](CheckDuplicateDomainCommand.md) |  | 

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


## DomainsCreate

> int32 DomainsCreate(ctx).CreateDomainCommand(createDomainCommand).Execute()

Create domain

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
	createDomainCommand := *openapiclient.NewCreateDomainCommand("Name_example", "Email_example", false) // CreateDomainCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.DomainsCreate(context.Background()).CreateDomainCommand(createDomainCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DomainsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainsCreate`: int32
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.DomainsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDomainsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDomainCommand** | [**CreateDomainCommand**](CreateDomainCommand.md) |  | 

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


## DomainsDelete

> DomainsDelete(ctx, id).Execute()

Delete domain

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
	r, err := apiClient.DomainsAPI.DomainsDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DomainsDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDomainsDeleteRequest struct via the builder pattern


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


## DomainsDetails

> DomainDetailsDto DomainsDetails(ctx, domainId).Execute()

Get domain details by domainId

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
	domainId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.DomainsDetails(context.Background(), domainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DomainsDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainsDetails`: DomainDetailsDto
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.DomainsDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DomainDetailsDto**](DomainDetailsDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsDisable2faManagement

> DomainsDisable2faManagement(ctx).DisableTwoFaManagementCommand(disableTwoFaManagementCommand).Execute()

Disable 2fa management

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
	disableTwoFaManagementCommand := *openapiclient.NewDisableTwoFaManagementCommand("VerificationCode_example") // DisableTwoFaManagementCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DomainsAPI.DomainsDisable2faManagement(context.Background()).DisableTwoFaManagementCommand(disableTwoFaManagementCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DomainsDisable2faManagement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDomainsDisable2faManagementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **disableTwoFaManagementCommand** | [**DisableTwoFaManagementCommand**](DisableTwoFaManagementCommand.md) |  | 

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


## DomainsDomainGroupDetails

> GroupDetailsDto DomainsDomainGroupDetails(ctx, domainId, id).Execute()

Get group details by domainId

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
	domainId := int32(56) // int32 | 
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.DomainsDomainGroupDetails(context.Background(), domainId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DomainsDomainGroupDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainsDomainGroupDetails`: GroupDetailsDto
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.DomainsDomainGroupDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32** |  | 
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsDomainGroupDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GroupDetailsDto**](GroupDetailsDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsDomainGroupsDropdown

> CommonDropdownDtoCursorPaginatedResponse DomainsDomainGroupsDropdown(ctx, domainId).Limit(limit).CursorId(cursorId).Search(search).Execute()

Retrieve domain groups dropdown list by cursor-based pagination

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
	domainId := int32(56) // int32 | 
	limit := int32(56) // int32 | 
	cursorId := int32(56) // int32 |  (optional)
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.DomainsDomainGroupsDropdown(context.Background(), domainId).Limit(limit).CursorId(cursorId).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DomainsDomainGroupsDropdown``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainsDomainGroupsDropdown`: CommonDropdownDtoCursorPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.DomainsDomainGroupsDropdown`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsDomainGroupsDropdownRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **cursorId** | **int32** |  | 
 **search** | **string** |  | 

### Return type

[**CommonDropdownDtoCursorPaginatedResponse**](CommonDropdownDtoCursorPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsDomainGroupsOffsetBasedDropdown

> GroupOffsetPaginationDropdownList DomainsDomainGroupsOffsetBasedDropdown(ctx, domainId).Offset(offset).Limit(limit).Search(search).SortBy(sortBy).SortDirection(sortDirection).OrganizationId(organizationId).UserId(userId).AccessLevel(accessLevel).Execute()

Retrieve domain groups dropdown list by offset-based pagination

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
	domainId := int32(56) // int32 | 
	offset := int32(56) // int32 |  (optional) (default to 0)
	limit := int32(56) // int32 |  (optional) (default to 100)
	search := "search_example" // string |  (optional)
	sortBy := "sortBy_example" // string |  (optional)
	sortDirection := "sortDirection_example" // string |  (optional)
	organizationId := int32(56) // int32 |  (optional)
	userId := "userId_example" // string |  (optional)
	accessLevel := openapiclient.AccessLevelRoles("Member") // AccessLevelRoles |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.DomainsDomainGroupsOffsetBasedDropdown(context.Background(), domainId).Offset(offset).Limit(limit).Search(search).SortBy(sortBy).SortDirection(sortDirection).OrganizationId(organizationId).UserId(userId).AccessLevel(accessLevel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DomainsDomainGroupsOffsetBasedDropdown``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainsDomainGroupsOffsetBasedDropdown`: GroupOffsetPaginationDropdownList
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.DomainsDomainGroupsOffsetBasedDropdown`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsDomainGroupsOffsetBasedDropdownRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 100]
 **search** | **string** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **organizationId** | **int32** |  | 
 **userId** | **string** |  | 
 **accessLevel** | [**AccessLevelRoles**](AccessLevelRoles.md) |  | 

### Return type

[**GroupOffsetPaginationDropdownList**](GroupOffsetPaginationDropdownList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsDomainOffsetBasedOrganizationsWithGroup

> OrganizationOffsetPaginationWithGroupList DomainsDomainOffsetBasedOrganizationsWithGroup(ctx, domainId).Offset(offset).Limit(limit).Search(search).ProjectId(projectId).GroupId(groupId).UserId(userId).SortDirection(sortDirection).Execute()

Retrieve domain organizations with group list by offset-based pagination

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
	domainId := int32(56) // int32 | 
	offset := int32(56) // int32 |  (optional) (default to 0)
	limit := int32(56) // int32 |  (optional) (default to 100)
	search := "search_example" // string |  (optional)
	projectId := int32(56) // int32 |  (optional)
	groupId := int32(56) // int32 |  (optional)
	userId := "userId_example" // string |  (optional)
	sortDirection := "sortDirection_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.DomainsDomainOffsetBasedOrganizationsWithGroup(context.Background(), domainId).Offset(offset).Limit(limit).Search(search).ProjectId(projectId).GroupId(groupId).UserId(userId).SortDirection(sortDirection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DomainsDomainOffsetBasedOrganizationsWithGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainsDomainOffsetBasedOrganizationsWithGroup`: OrganizationOffsetPaginationWithGroupList
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.DomainsDomainOffsetBasedOrganizationsWithGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsDomainOffsetBasedOrganizationsWithGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 100]
 **search** | **string** |  | 
 **projectId** | **int32** |  | 
 **groupId** | **int32** |  | 
 **userId** | **string** |  | 
 **sortDirection** | **string** |  | 

### Return type

[**OrganizationOffsetPaginationWithGroupList**](OrganizationOffsetPaginationWithGroupList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsDomainOrganizationDetails

> OrganizationDetailsDtoForDomain DomainsDomainOrganizationDetails(ctx, domainId, id).Execute()

Get organization details by domainId

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
	domainId := int32(56) // int32 | 
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.DomainsDomainOrganizationDetails(context.Background(), domainId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DomainsDomainOrganizationDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainsDomainOrganizationDetails`: OrganizationDetailsDtoForDomain
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.DomainsDomainOrganizationDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32** |  | 
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsDomainOrganizationDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrganizationDetailsDtoForDomain**](OrganizationDetailsDtoForDomain.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsDomainOrganizationsAvailable

> []CommonDropdownDto DomainsDomainOrganizationsAvailable(ctx, domainId).GroupId(groupId).Execute()

Retrieve available organizations for adding to a group

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
	domainId := int32(56) // int32 | 
	groupId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.DomainsDomainOrganizationsAvailable(context.Background(), domainId).GroupId(groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DomainsDomainOrganizationsAvailable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainsDomainOrganizationsAvailable`: []CommonDropdownDto
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.DomainsDomainOrganizationsAvailable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsDomainOrganizationsAvailableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupId** | **int32** |  | 

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


## DomainsDomainOrganizationsWithGroup

> OrganizationsWithGroupInfoResultDtoCursorPaginatedResponse DomainsDomainOrganizationsWithGroup(ctx, domainId).GroupId(groupId).Limit(limit).CursorId(cursorId).Search(search).Execute()

Retrieve domain organizations with group list by cursor-based pagination

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
	domainId := int32(56) // int32 | 
	groupId := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 |  (optional)
	cursorId := int32(56) // int32 |  (optional)
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.DomainsDomainOrganizationsWithGroup(context.Background(), domainId).GroupId(groupId).Limit(limit).CursorId(cursorId).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DomainsDomainOrganizationsWithGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainsDomainOrganizationsWithGroup`: OrganizationsWithGroupInfoResultDtoCursorPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.DomainsDomainOrganizationsWithGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsDomainOrganizationsWithGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupId** | **int32** |  | 
 **limit** | **int32** |  | 
 **cursorId** | **int32** |  | 
 **search** | **string** |  | 

### Return type

[**OrganizationsWithGroupInfoResultDtoCursorPaginatedResponse**](OrganizationsWithGroupInfoResultDtoCursorPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsDomainProjectDetails

> ProjectDetailsDto DomainsDomainProjectDetails(ctx, domainId, id).Execute()

Get project details by domainId

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
	domainId := int32(56) // int32 | 
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.DomainsDomainProjectDetails(context.Background(), domainId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DomainsDomainProjectDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainsDomainProjectDetails`: ProjectDetailsDto
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.DomainsDomainProjectDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32** |  | 
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsDomainProjectDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ProjectDetailsDto**](ProjectDetailsDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsDomainProjectsDropdown

> CommonDropdownDtoCursorPaginatedResponse DomainsDomainProjectsDropdown(ctx, domainId).Limit(limit).CursorId(cursorId).Search(search).Execute()

Retrieve domain projects dropdown list by cursor-based pagination

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
	domainId := int32(56) // int32 | 
	limit := int32(56) // int32 | 
	cursorId := int32(56) // int32 |  (optional)
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.DomainsDomainProjectsDropdown(context.Background(), domainId).Limit(limit).CursorId(cursorId).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DomainsDomainProjectsDropdown``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainsDomainProjectsDropdown`: CommonDropdownDtoCursorPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.DomainsDomainProjectsDropdown`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsDomainProjectsDropdownRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **cursorId** | **int32** |  | 
 **search** | **string** |  | 

### Return type

[**CommonDropdownDtoCursorPaginatedResponse**](CommonDropdownDtoCursorPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsDomainProjectsOffsetBasedDropdown

> ProjectOffsetPaginationList DomainsDomainProjectsOffsetBasedDropdown(ctx, domainId).Offset(offset).Limit(limit).Search(search).SortBy(sortBy).SortDirection(sortDirection).Status(status).Health(health).OrganizationId(organizationId).UserId(userId).GroupId(groupId).Execute()

Retrieve domain projects dropdown list by offset-based pagination

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
	domainId := int32(56) // int32 | 
	offset := int32(56) // int32 |  (optional) (default to 0)
	limit := int32(56) // int32 |  (optional) (default to 100)
	search := "search_example" // string |  (optional)
	sortBy := "sortBy_example" // string |  (optional)
	sortDirection := "sortDirection_example" // string |  (optional)
	status := openapiclient.ProjectStatus("Null") // ProjectStatus |  (optional)
	health := openapiclient.ProjectHealth("Null") // ProjectHealth |  (optional)
	organizationId := int32(56) // int32 |  (optional)
	userId := "userId_example" // string |  (optional)
	groupId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.DomainsDomainProjectsOffsetBasedDropdown(context.Background(), domainId).Offset(offset).Limit(limit).Search(search).SortBy(sortBy).SortDirection(sortDirection).Status(status).Health(health).OrganizationId(organizationId).UserId(userId).GroupId(groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DomainsDomainProjectsOffsetBasedDropdown``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainsDomainProjectsOffsetBasedDropdown`: ProjectOffsetPaginationList
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.DomainsDomainProjectsOffsetBasedDropdown`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsDomainProjectsOffsetBasedDropdownRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 100]
 **search** | **string** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **status** | [**ProjectStatus**](ProjectStatus.md) |  | 
 **health** | [**ProjectHealth**](ProjectHealth.md) |  | 
 **organizationId** | **int32** |  | 
 **userId** | **string** |  | 
 **groupId** | **int32** |  | 

### Return type

[**ProjectOffsetPaginationList**](ProjectOffsetPaginationList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsDomainUserDetails

> UserDetailsDto DomainsDomainUserDetails(ctx, domainId, id).Execute()

Get user details by domainId

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
	domainId := int32(56) // int32 | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.DomainsDomainUserDetails(context.Background(), domainId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DomainsDomainUserDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainsDomainUserDetails`: UserDetailsDto
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.DomainsDomainUserDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsDomainUserDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UserDetailsDto**](UserDetailsDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsDomainUserDropdown

> UserBriefDtoCursorTimestampPaginatedResponse DomainsDomainUserDropdown(ctx, domainId).Limit(limit).CursorId(cursorId).Search(search).Execute()

Retrieve domain users dropdown list by cursor-based pagination

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
	domainId := int32(56) // int32 | 
	limit := int32(56) // int32 | 
	cursorId := int64(789) // int64 |  (optional)
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.DomainsDomainUserDropdown(context.Background(), domainId).Limit(limit).CursorId(cursorId).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DomainsDomainUserDropdown``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainsDomainUserDropdown`: UserBriefDtoCursorTimestampPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.DomainsDomainUserDropdown`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsDomainUserDropdownRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **cursorId** | **int64** |  | 
 **search** | **string** |  | 

### Return type

[**UserBriefDtoCursorTimestampPaginatedResponse**](UserBriefDtoCursorTimestampPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsDomainUserOffsetBasedDropdown

> UserOffsetPaginationDropdownList DomainsDomainUserOffsetBasedDropdown(ctx, domainId).FilterRole(filterRole).Limit(limit).Offset(offset).Search(search).SortBy(sortBy).SortDirection(sortDirection).FilterOrganizationId(filterOrganizationId).FilterProjectId(filterProjectId).FilterGroupId(filterGroupId).Execute()

Retrieve domain users dropdown list by offset-based pagination

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
	domainId := int32(56) // int32 | 
	filterRole := openapiclient.EFilteredUserRole("None") // EFilteredUserRole |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)
	offset := int32(56) // int32 |  (optional) (default to 0)
	search := "search_example" // string |  (optional)
	sortBy := "sortBy_example" // string |  (optional)
	sortDirection := "sortDirection_example" // string |  (optional)
	filterOrganizationId := int32(56) // int32 |  (optional)
	filterProjectId := int32(56) // int32 |  (optional)
	filterGroupId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.DomainsDomainUserOffsetBasedDropdown(context.Background(), domainId).FilterRole(filterRole).Limit(limit).Offset(offset).Search(search).SortBy(sortBy).SortDirection(sortDirection).FilterOrganizationId(filterOrganizationId).FilterProjectId(filterProjectId).FilterGroupId(filterGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DomainsDomainUserOffsetBasedDropdown``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainsDomainUserOffsetBasedDropdown`: UserOffsetPaginationDropdownList
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.DomainsDomainUserOffsetBasedDropdown`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsDomainUserOffsetBasedDropdownRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterRole** | [**EFilteredUserRole**](EFilteredUserRole.md) |  | 
 **limit** | **int32** |  | [default to 50]
 **offset** | **int32** |  | [default to 0]
 **search** | **string** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **filterOrganizationId** | **int32** |  | 
 **filterProjectId** | **int32** |  | 
 **filterGroupId** | **int32** |  | 

### Return type

[**UserOffsetPaginationDropdownList**](UserOffsetPaginationDropdownList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsDomainUsersAvailable

> []CommonStringBasedDropdownDto DomainsDomainUsersAvailable(ctx, domainId).GroupId(groupId).Execute()

Retrieve available users for adding to a group

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
	domainId := int32(56) // int32 | 
	groupId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.DomainsDomainUsersAvailable(context.Background(), domainId).GroupId(groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DomainsDomainUsersAvailable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainsDomainUsersAvailable`: []CommonStringBasedDropdownDto
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.DomainsDomainUsersAvailable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsDomainUsersAvailableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupId** | **int32** |  | 

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


## DomainsEnable2faManagement

> DomainsEnable2faManagement(ctx).Body(body).Execute()

Enable 2fa management

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
	r, err := apiClient.DomainsAPI.DomainsEnable2faManagement(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DomainsEnable2faManagement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDomainsEnable2faManagementRequest struct via the builder pattern


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


## DomainsListDomains

> DomainListCursorPaginatedResponse DomainsListDomains(ctx).Limit(limit).CursorId(cursorId).Search(search).Execute()

Retrieve domains list

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
	limit := int32(56) // int32 |  (optional)
	cursorId := int32(56) // int32 |  (optional)
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.DomainsListDomains(context.Background()).Limit(limit).CursorId(cursorId).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DomainsListDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainsListDomains`: DomainListCursorPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.DomainsListDomains`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDomainsListDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **cursorId** | **int32** |  | 
 **search** | **string** |  | 

### Return type

[**DomainListCursorPaginatedResponse**](DomainListCursorPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsTransferOwnership

> DomainsTransferOwnership(ctx).TransferOwnershipCommand(transferOwnershipCommand).Execute()

Transfer ownership

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
	transferOwnershipCommand := *openapiclient.NewTransferOwnershipCommand("UserId_example") // TransferOwnershipCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DomainsAPI.DomainsTransferOwnership(context.Background()).TransferOwnershipCommand(transferOwnershipCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DomainsTransferOwnership``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDomainsTransferOwnershipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transferOwnershipCommand** | [**TransferOwnershipCommand**](TransferOwnershipCommand.md) |  | 

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


## DomainsUpdate

> int32 DomainsUpdate(ctx).UpdateDomainCommand(updateDomainCommand).Execute()

Update domain

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
	updateDomainCommand := *openapiclient.NewUpdateDomainCommand(int32(123), "Name_example", "Email_example") // UpdateDomainCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.DomainsUpdate(context.Background()).UpdateDomainCommand(updateDomainCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DomainsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainsUpdate`: int32
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.DomainsUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDomainsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateDomainCommand** | [**UpdateDomainCommand**](UpdateDomainCommand.md) |  | 

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

