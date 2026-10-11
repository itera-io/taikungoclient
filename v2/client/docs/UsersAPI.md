# \UsersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersChangePassword**](UsersAPI.md#UsersChangePassword) | **Post** /api/v1/users/changepassword | Change user password
[**UsersConfirmEmail**](UsersAPI.md#UsersConfirmEmail) | **Post** /api/v1/users/confirmemail | Confirm user email
[**UsersCreate**](UsersAPI.md#UsersCreate) | **Post** /api/v1/users/create | Create user
[**UsersDelete**](UsersAPI.md#UsersDelete) | **Delete** /api/v1/users/{id} | Delete user
[**UsersDeleteMyUser**](UsersAPI.md#UsersDeleteMyUser) | **Post** /api/v1/users/delete | Delete my user
[**UsersDisable**](UsersAPI.md#UsersDisable) | **Post** /api/v1/users/disable | Disable user
[**UsersDowngradeRole**](UsersAPI.md#UsersDowngradeRole) | **Post** /api/v1/users/downgrade | Downgrade user role
[**UsersDropdown**](UsersAPI.md#UsersDropdown) | **Get** /api/v1/users/dropdown | Retrieve user info
[**UsersForceToResetPassword**](UsersAPI.md#UsersForceToResetPassword) | **Post** /api/v1/users/force-to-reset | Force to reset password
[**UsersManageRole**](UsersAPI.md#UsersManageRole) | **Post** /api/v1/users/manage-role | Manage user role
[**UsersPromoteRole**](UsersAPI.md#UsersPromoteRole) | **Post** /api/v1/users/promote-role | Promote user role
[**UsersToggleMaintenanceMode**](UsersAPI.md#UsersToggleMaintenanceMode) | **Post** /api/v1/users/togglemaintenancemode | Toggle maintenance mode
[**UsersToggleNotificationMode**](UsersAPI.md#UsersToggleNotificationMode) | **Post** /api/v1/users/togglenotificationmode | Toggle notification mode
[**UsersUnlock2fa**](UsersAPI.md#UsersUnlock2fa) | **Post** /api/v1/users/unlock-2fa | Unlock 2fa mode
[**UsersUpdateUser**](UsersAPI.md#UsersUpdateUser) | **Post** /api/v1/users/update | Update user
[**UsersUserInfo**](UsersAPI.md#UsersUserInfo) | **Get** /api/v1/users/userinfo | Retrieve user info
[**UsersVerifyEmail**](UsersAPI.md#UsersVerifyEmail) | **Post** /api/v1/users/verifyemail | Verify user email



## UsersChangePassword

> UsersChangePassword(ctx).ChangePasswordCommand(changePasswordCommand).Execute()

Change user password

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
	changePasswordCommand := *openapiclient.NewChangePasswordCommand("Password_example", "NewPassword_example") // ChangePasswordCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersChangePassword(context.Background()).ChangePasswordCommand(changePasswordCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersChangePassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersChangePasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **changePasswordCommand** | [**ChangePasswordCommand**](ChangePasswordCommand.md) |  | 

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


## UsersConfirmEmail

> UsersConfirmEmail(ctx).ConfirmEmailCommand(confirmEmailCommand).Execute()

Confirm user email

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
	confirmEmailCommand := *openapiclient.NewConfirmEmailCommand("NewEmail_example", openapiclient.EmailMode("Confirm")) // ConfirmEmailCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersConfirmEmail(context.Background()).ConfirmEmailCommand(confirmEmailCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersConfirmEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersConfirmEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **confirmEmailCommand** | [**ConfirmEmailCommand**](ConfirmEmailCommand.md) |  | 

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


## UsersCreate

> ApiResponse UsersCreate(ctx).CreateUserCommand(createUserCommand).Execute()

Create user

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
	createUserCommand := *openapiclient.NewCreateUserCommand() // CreateUserCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersCreate(context.Background()).CreateUserCommand(createUserCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersCreate`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUserCommand** | [**CreateUserCommand**](CreateUserCommand.md) |  | 

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


## UsersDelete

> UsersDelete(ctx, id).Execute()

Delete user

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
	r, err := apiClient.UsersAPI.UsersDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUsersDeleteRequest struct via the builder pattern


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


## UsersDeleteMyUser

> UsersDeleteMyUser(ctx).Body(body).Execute()

Delete my user

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
	r, err := apiClient.UsersAPI.UsersDeleteMyUser(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersDeleteMyUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersDeleteMyUserRequest struct via the builder pattern


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


## UsersDisable

> UsersDisable(ctx).DisableUserCommand(disableUserCommand).Execute()

Disable user

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
	disableUserCommand := *openapiclient.NewDisableUserCommand("Id_example", false) // DisableUserCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersDisable(context.Background()).DisableUserCommand(disableUserCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersDisable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersDisableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **disableUserCommand** | [**DisableUserCommand**](DisableUserCommand.md) |  | 

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


## UsersDowngradeRole

> UsersDowngradeRole(ctx).DowngradeUserCommand(downgradeUserCommand).Execute()

Downgrade user role

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
	downgradeUserCommand := *openapiclient.NewDowngradeUserCommand("UserId_example") // DowngradeUserCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersDowngradeRole(context.Background()).DowngradeUserCommand(downgradeUserCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersDowngradeRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersDowngradeRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **downgradeUserCommand** | [**DowngradeUserCommand**](DowngradeUserCommand.md) |  | 

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


## UsersDropdown

> UserByOrgProjectDropdownList UsersDropdown(ctx).OrganizationId(organizationId).ProjectId(projectId).Limit(limit).Offset(offset).Search(search).Execute()

Retrieve user info

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
	projectId := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 500)
	offset := int32(56) // int32 |  (optional) (default to 0)
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersDropdown(context.Background()).OrganizationId(organizationId).ProjectId(projectId).Limit(limit).Offset(offset).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersDropdown``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersDropdown`: UserByOrgProjectDropdownList
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersDropdown`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersDropdownRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 
 **projectId** | **int32** |  | 
 **limit** | **int32** |  | [default to 500]
 **offset** | **int32** |  | [default to 0]
 **search** | **string** |  | 

### Return type

[**UserByOrgProjectDropdownList**](UserByOrgProjectDropdownList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersForceToResetPassword

> UsersForceToResetPassword(ctx).ForceToResetPasswordCommand(forceToResetPasswordCommand).Execute()

Force to reset password

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
	forceToResetPasswordCommand := *openapiclient.NewForceToResetPasswordCommand("Id_example") // ForceToResetPasswordCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersForceToResetPassword(context.Background()).ForceToResetPasswordCommand(forceToResetPasswordCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersForceToResetPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersForceToResetPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **forceToResetPasswordCommand** | [**ForceToResetPasswordCommand**](ForceToResetPasswordCommand.md) |  | 

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


## UsersManageRole

> UsersManageRole(ctx).ManageRoleCommand(manageRoleCommand).Execute()

Manage user role

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
	manageRoleCommand := *openapiclient.NewManageRoleCommand("UserId_example", openapiclient.EGlobalRole("None")) // ManageRoleCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersManageRole(context.Background()).ManageRoleCommand(manageRoleCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersManageRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersManageRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **manageRoleCommand** | [**ManageRoleCommand**](ManageRoleCommand.md) |  | 

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


## UsersPromoteRole

> UsersPromoteRole(ctx).PromoteUserCommand(promoteUserCommand).Execute()

Promote user role

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
	promoteUserCommand := *openapiclient.NewPromoteUserCommand("UserId_example", false) // PromoteUserCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersPromoteRole(context.Background()).PromoteUserCommand(promoteUserCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersPromoteRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersPromoteRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **promoteUserCommand** | [**PromoteUserCommand**](PromoteUserCommand.md) |  | 

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


## UsersToggleMaintenanceMode

> UsersToggleMaintenanceMode(ctx).ToggleMaintenanceModeCommand(toggleMaintenanceModeCommand).Execute()

Toggle maintenance mode

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
	toggleMaintenanceModeCommand := *openapiclient.NewToggleMaintenanceModeCommand("Mode_example") // ToggleMaintenanceModeCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersToggleMaintenanceMode(context.Background()).ToggleMaintenanceModeCommand(toggleMaintenanceModeCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersToggleMaintenanceMode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersToggleMaintenanceModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **toggleMaintenanceModeCommand** | [**ToggleMaintenanceModeCommand**](ToggleMaintenanceModeCommand.md) |  | 

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


## UsersToggleNotificationMode

> UsersToggleNotificationMode(ctx).ToggleNotificationModeCommand(toggleNotificationModeCommand).Execute()

Toggle notification mode

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
	toggleNotificationModeCommand := *openapiclient.NewToggleNotificationModeCommand("Mode_example") // ToggleNotificationModeCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersToggleNotificationMode(context.Background()).ToggleNotificationModeCommand(toggleNotificationModeCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersToggleNotificationMode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersToggleNotificationModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **toggleNotificationModeCommand** | [**ToggleNotificationModeCommand**](ToggleNotificationModeCommand.md) |  | 

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


## UsersUnlock2fa

> UsersUnlock2fa(ctx).Unlock2FaCommand(unlock2FaCommand).Execute()

Unlock 2fa mode

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
	unlock2FaCommand := *openapiclient.NewUnlock2FaCommand("UserId_example") // Unlock2FaCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersUnlock2fa(context.Background()).Unlock2FaCommand(unlock2FaCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersUnlock2fa``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersUnlock2faRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unlock2FaCommand** | [**Unlock2FaCommand**](Unlock2FaCommand.md) |  | 

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


## UsersUpdateUser

> UsersUpdateUser(ctx).UpdateUserCommand(updateUserCommand).Execute()

Update user

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
	updateUserCommand := *openapiclient.NewUpdateUserCommand() // UpdateUserCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersUpdateUser(context.Background()).UpdateUserCommand(updateUserCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersUpdateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateUserCommand** | [**UpdateUserCommand**](UpdateUserCommand.md) |  | 

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


## UsersUserInfo

> UserDetails UsersUserInfo(ctx).Execute()

Retrieve user info

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
	resp, r, err := apiClient.UsersAPI.UsersUserInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersUserInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersUserInfo`: UserDetails
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersUserInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUserInfoRequest struct via the builder pattern


### Return type

[**UserDetails**](UserDetails.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersVerifyEmail

> UsersVerifyEmail(ctx).VerifyEmailCommand(verifyEmailCommand).Execute()

Verify user email

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
	verifyEmailCommand := *openapiclient.NewVerifyEmailCommand("Token_example", openapiclient.EmailMode("Confirm")) // VerifyEmailCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersVerifyEmail(context.Background()).VerifyEmailCommand(verifyEmailCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersVerifyEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersVerifyEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **verifyEmailCommand** | [**VerifyEmailCommand**](VerifyEmailCommand.md) |  | 

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

