# TaikunResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSuccess** | Pointer to **bool** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**ProblemDetails** | Pointer to [**ProblemDetails**](ProblemDetails.md) |  | [optional] 
**HttpError** | Pointer to [**ETaikunHttpError**](ETaikunHttpError.md) |  | [optional] 
**RedirectUrl** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTaikunResult

`func NewTaikunResult() *TaikunResult`

NewTaikunResult instantiates a new TaikunResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaikunResultWithDefaults

`func NewTaikunResultWithDefaults() *TaikunResult`

NewTaikunResultWithDefaults instantiates a new TaikunResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSuccess

`func (o *TaikunResult) GetIsSuccess() bool`

GetIsSuccess returns the IsSuccess field if non-nil, zero value otherwise.

### GetIsSuccessOk

`func (o *TaikunResult) GetIsSuccessOk() (*bool, bool)`

GetIsSuccessOk returns a tuple with the IsSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuccess

`func (o *TaikunResult) SetIsSuccess(v bool)`

SetIsSuccess sets IsSuccess field to given value.

### HasIsSuccess

`func (o *TaikunResult) HasIsSuccess() bool`

HasIsSuccess returns a boolean if a field has been set.

### GetErrorMessage

`func (o *TaikunResult) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *TaikunResult) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *TaikunResult) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *TaikunResult) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *TaikunResult) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *TaikunResult) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetProblemDetails

`func (o *TaikunResult) GetProblemDetails() ProblemDetails`

GetProblemDetails returns the ProblemDetails field if non-nil, zero value otherwise.

### GetProblemDetailsOk

`func (o *TaikunResult) GetProblemDetailsOk() (*ProblemDetails, bool)`

GetProblemDetailsOk returns a tuple with the ProblemDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProblemDetails

`func (o *TaikunResult) SetProblemDetails(v ProblemDetails)`

SetProblemDetails sets ProblemDetails field to given value.

### HasProblemDetails

`func (o *TaikunResult) HasProblemDetails() bool`

HasProblemDetails returns a boolean if a field has been set.

### GetHttpError

`func (o *TaikunResult) GetHttpError() ETaikunHttpError`

GetHttpError returns the HttpError field if non-nil, zero value otherwise.

### GetHttpErrorOk

`func (o *TaikunResult) GetHttpErrorOk() (*ETaikunHttpError, bool)`

GetHttpErrorOk returns a tuple with the HttpError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpError

`func (o *TaikunResult) SetHttpError(v ETaikunHttpError)`

SetHttpError sets HttpError field to given value.

### HasHttpError

`func (o *TaikunResult) HasHttpError() bool`

HasHttpError returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *TaikunResult) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *TaikunResult) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *TaikunResult) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *TaikunResult) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### SetRedirectUrlNil

`func (o *TaikunResult) SetRedirectUrlNil(b bool)`

 SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil

### UnsetRedirectUrl
`func (o *TaikunResult) UnsetRedirectUrl()`

UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


