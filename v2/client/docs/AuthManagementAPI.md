# \AuthManagementAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthForgotPassword**](AuthManagementAPI.md#AuthForgotPassword) | **Post** /api/v1/auth/forgotpassword | Generate reset password token if you forgot password
[**AuthLogin**](AuthManagementAPI.md#AuthLogin) | **Post** /api/v1/auth/login | Login to API
[**AuthRefresh**](AuthManagementAPI.md#AuthRefresh) | **Post** /api/v1/auth/refresh | Refresh bearer token that generated automatically by API
[**AuthResetPassword**](AuthManagementAPI.md#AuthResetPassword) | **Post** /api/v1/auth/resetpassword | Reset password
[**AuthSsoCallbackPost**](AuthManagementAPI.md#AuthSsoCallbackPost) | **Post** /api/v1/auth/sso/callback | Handle SSO callback from identity provider (POST with body)
[**AuthSsoInit**](AuthManagementAPI.md#AuthSsoInit) | **Post** /api/v1/auth/sso/init | Initiate SSO login flow
[**TwoFactorAuthComplete**](AuthManagementAPI.md#TwoFactorAuthComplete) | **Post** /api/v1/two-factor-auth/complete | Ensure recovery codes saved
[**TwoFactorAuthDisable**](AuthManagementAPI.md#TwoFactorAuthDisable) | **Post** /api/v1/two-factor-auth/disable | Disable 2fa by code
[**TwoFactorAuthLogin**](AuthManagementAPI.md#TwoFactorAuthLogin) | **Post** /api/v1/two-factor-auth/login | Login with TOTP
[**TwoFactorAuthRecovery**](AuthManagementAPI.md#TwoFactorAuthRecovery) | **Post** /api/v1/two-factor-auth/recovery | Login with recovery code
[**TwoFactorAuthSetup**](AuthManagementAPI.md#TwoFactorAuthSetup) | **Post** /api/v1/two-factor-auth/setup | Setup 2fa
[**TwoFactorAuthVerify**](AuthManagementAPI.md#TwoFactorAuthVerify) | **Post** /api/v1/two-factor-auth/verify | Verify 2fa code



## AuthForgotPassword

> AuthForgotPassword(ctx).ForgotPasswordCommand(forgotPasswordCommand).Execute()

Generate reset password token if you forgot password

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
	forgotPasswordCommand := *openapiclient.NewForgotPasswordCommand("Email_example", "DomainName_example") // ForgotPasswordCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthManagementAPI.AuthForgotPassword(context.Background()).ForgotPasswordCommand(forgotPasswordCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthManagementAPI.AuthForgotPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthForgotPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **forgotPasswordCommand** | [**ForgotPasswordCommand**](ForgotPasswordCommand.md) |  | 

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


## AuthLogin

> GetToken AuthLogin(ctx).LoginCommand(loginCommand).Execute()

Login to API

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
	loginCommand := *openapiclient.NewLoginCommand() // LoginCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthManagementAPI.AuthLogin(context.Background()).LoginCommand(loginCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthManagementAPI.AuthLogin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthLogin`: GetToken
	fmt.Fprintf(os.Stdout, "Response from `AuthManagementAPI.AuthLogin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **loginCommand** | [**LoginCommand**](LoginCommand.md) |  | 

### Return type

[**GetToken**](GetToken.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthRefresh

> GetToken AuthRefresh(ctx).RefreshTokenCommand(refreshTokenCommand).Execute()

Refresh bearer token that generated automatically by API

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
	refreshTokenCommand := *openapiclient.NewRefreshTokenCommand("Token_example", "RefreshToken_example") // RefreshTokenCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthManagementAPI.AuthRefresh(context.Background()).RefreshTokenCommand(refreshTokenCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthManagementAPI.AuthRefresh``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthRefresh`: GetToken
	fmt.Fprintf(os.Stdout, "Response from `AuthManagementAPI.AuthRefresh`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthRefreshRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refreshTokenCommand** | [**RefreshTokenCommand**](RefreshTokenCommand.md) |  | 

### Return type

[**GetToken**](GetToken.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthResetPassword

> AuthResetPassword(ctx).ResetPasswordCommand(resetPasswordCommand).Execute()

Reset password

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
	resetPasswordCommand := *openapiclient.NewResetPasswordCommand() // ResetPasswordCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthManagementAPI.AuthResetPassword(context.Background()).ResetPasswordCommand(resetPasswordCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthManagementAPI.AuthResetPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthResetPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resetPasswordCommand** | [**ResetPasswordCommand**](ResetPasswordCommand.md) |  | 

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


## AuthSsoCallbackPost

> GetToken AuthSsoCallbackPost(ctx).SsoCallbackCommand(ssoCallbackCommand).Execute()

Handle SSO callback from identity provider (POST with body)

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
	ssoCallbackCommand := *openapiclient.NewSsoCallbackCommand() // SsoCallbackCommand | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthManagementAPI.AuthSsoCallbackPost(context.Background()).SsoCallbackCommand(ssoCallbackCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthManagementAPI.AuthSsoCallbackPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthSsoCallbackPost`: GetToken
	fmt.Fprintf(os.Stdout, "Response from `AuthManagementAPI.AuthSsoCallbackPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthSsoCallbackPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ssoCallbackCommand** | [**SsoCallbackCommand**](SsoCallbackCommand.md) |  | 

### Return type

[**GetToken**](GetToken.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthSsoInit

> SsoInitResult AuthSsoInit(ctx).SsoInitCommand(ssoInitCommand).Execute()

Initiate SSO login flow



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
	ssoInitCommand := *openapiclient.NewSsoInitCommand() // SsoInitCommand | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthManagementAPI.AuthSsoInit(context.Background()).SsoInitCommand(ssoInitCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthManagementAPI.AuthSsoInit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthSsoInit`: SsoInitResult
	fmt.Fprintf(os.Stdout, "Response from `AuthManagementAPI.AuthSsoInit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthSsoInitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ssoInitCommand** | [**SsoInitCommand**](SsoInitCommand.md) |  | 

### Return type

[**SsoInitResult**](SsoInitResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TwoFactorAuthComplete

> GetToken TwoFactorAuthComplete(ctx).EnsureCodesDownloadedCommand(ensureCodesDownloadedCommand).Execute()

Ensure recovery codes saved

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
	ensureCodesDownloadedCommand := *openapiclient.NewEnsureCodesDownloadedCommand(*openapiclient.NewTwoFactorAuthVerifyResult([]string{"RecoveryCodes_example"})) // EnsureCodesDownloadedCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthManagementAPI.TwoFactorAuthComplete(context.Background()).EnsureCodesDownloadedCommand(ensureCodesDownloadedCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthManagementAPI.TwoFactorAuthComplete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TwoFactorAuthComplete`: GetToken
	fmt.Fprintf(os.Stdout, "Response from `AuthManagementAPI.TwoFactorAuthComplete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTwoFactorAuthCompleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ensureCodesDownloadedCommand** | [**EnsureCodesDownloadedCommand**](EnsureCodesDownloadedCommand.md) |  | 

### Return type

[**GetToken**](GetToken.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TwoFactorAuthDisable

> TwoFactorAuthDisable(ctx).DisableTwoFactorAuthCommand(disableTwoFactorAuthCommand).Execute()

Disable 2fa by code

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
	disableTwoFactorAuthCommand := *openapiclient.NewDisableTwoFactorAuthCommand("Code_example", false) // DisableTwoFactorAuthCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthManagementAPI.TwoFactorAuthDisable(context.Background()).DisableTwoFactorAuthCommand(disableTwoFactorAuthCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthManagementAPI.TwoFactorAuthDisable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTwoFactorAuthDisableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **disableTwoFactorAuthCommand** | [**DisableTwoFactorAuthCommand**](DisableTwoFactorAuthCommand.md) |  | 

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


## TwoFactorAuthLogin

> GetToken TwoFactorAuthLogin(ctx).TwoFactorAuthLoginCommand(twoFactorAuthLoginCommand).Execute()

Login with TOTP

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
	twoFactorAuthLoginCommand := *openapiclient.NewTwoFactorAuthLoginCommand("VerificationCode_example") // TwoFactorAuthLoginCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthManagementAPI.TwoFactorAuthLogin(context.Background()).TwoFactorAuthLoginCommand(twoFactorAuthLoginCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthManagementAPI.TwoFactorAuthLogin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TwoFactorAuthLogin`: GetToken
	fmt.Fprintf(os.Stdout, "Response from `AuthManagementAPI.TwoFactorAuthLogin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTwoFactorAuthLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **twoFactorAuthLoginCommand** | [**TwoFactorAuthLoginCommand**](TwoFactorAuthLoginCommand.md) |  | 

### Return type

[**GetToken**](GetToken.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TwoFactorAuthRecovery

> GetToken TwoFactorAuthRecovery(ctx).TwoFactorAuthRecoveryCommand(twoFactorAuthRecoveryCommand).Execute()

Login with recovery code

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
	twoFactorAuthRecoveryCommand := *openapiclient.NewTwoFactorAuthRecoveryCommand("RecoveryCode_example") // TwoFactorAuthRecoveryCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthManagementAPI.TwoFactorAuthRecovery(context.Background()).TwoFactorAuthRecoveryCommand(twoFactorAuthRecoveryCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthManagementAPI.TwoFactorAuthRecovery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TwoFactorAuthRecovery`: GetToken
	fmt.Fprintf(os.Stdout, "Response from `AuthManagementAPI.TwoFactorAuthRecovery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTwoFactorAuthRecoveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **twoFactorAuthRecoveryCommand** | [**TwoFactorAuthRecoveryCommand**](TwoFactorAuthRecoveryCommand.md) |  | 

### Return type

[**GetToken**](GetToken.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TwoFactorAuthSetup

> TwoFactorAuthSetupResult TwoFactorAuthSetup(ctx).Body(body).Execute()

Setup 2fa

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
	resp, r, err := apiClient.AuthManagementAPI.TwoFactorAuthSetup(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthManagementAPI.TwoFactorAuthSetup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TwoFactorAuthSetup`: TwoFactorAuthSetupResult
	fmt.Fprintf(os.Stdout, "Response from `AuthManagementAPI.TwoFactorAuthSetup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTwoFactorAuthSetupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**TwoFactorAuthSetupResult**](TwoFactorAuthSetupResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TwoFactorAuthVerify

> TwoFactorAuthVerifyResult TwoFactorAuthVerify(ctx).TwoFactorAuthVerifyCommand(twoFactorAuthVerifyCommand).Execute()

Verify 2fa code

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
	twoFactorAuthVerifyCommand := *openapiclient.NewTwoFactorAuthVerifyCommand("VerificationCode_example") // TwoFactorAuthVerifyCommand |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthManagementAPI.TwoFactorAuthVerify(context.Background()).TwoFactorAuthVerifyCommand(twoFactorAuthVerifyCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthManagementAPI.TwoFactorAuthVerify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TwoFactorAuthVerify`: TwoFactorAuthVerifyResult
	fmt.Fprintf(os.Stdout, "Response from `AuthManagementAPI.TwoFactorAuthVerify`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTwoFactorAuthVerifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **twoFactorAuthVerifyCommand** | [**TwoFactorAuthVerifyCommand**](TwoFactorAuthVerifyCommand.md) |  | 

### Return type

[**TwoFactorAuthVerifyResult**](TwoFactorAuthVerifyResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

