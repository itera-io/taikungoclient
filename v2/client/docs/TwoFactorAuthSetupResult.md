# TwoFactorAuthSetupResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SharedKey** | **NullableString** |  | 
**QrCodeUri** | **NullableString** |  | 

## Methods

### NewTwoFactorAuthSetupResult

`func NewTwoFactorAuthSetupResult(sharedKey NullableString, qrCodeUri NullableString, ) *TwoFactorAuthSetupResult`

NewTwoFactorAuthSetupResult instantiates a new TwoFactorAuthSetupResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTwoFactorAuthSetupResultWithDefaults

`func NewTwoFactorAuthSetupResultWithDefaults() *TwoFactorAuthSetupResult`

NewTwoFactorAuthSetupResultWithDefaults instantiates a new TwoFactorAuthSetupResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSharedKey

`func (o *TwoFactorAuthSetupResult) GetSharedKey() string`

GetSharedKey returns the SharedKey field if non-nil, zero value otherwise.

### GetSharedKeyOk

`func (o *TwoFactorAuthSetupResult) GetSharedKeyOk() (*string, bool)`

GetSharedKeyOk returns a tuple with the SharedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedKey

`func (o *TwoFactorAuthSetupResult) SetSharedKey(v string)`

SetSharedKey sets SharedKey field to given value.


### SetSharedKeyNil

`func (o *TwoFactorAuthSetupResult) SetSharedKeyNil(b bool)`

 SetSharedKeyNil sets the value for SharedKey to be an explicit nil

### UnsetSharedKey
`func (o *TwoFactorAuthSetupResult) UnsetSharedKey()`

UnsetSharedKey ensures that no value is present for SharedKey, not even an explicit nil
### GetQrCodeUri

`func (o *TwoFactorAuthSetupResult) GetQrCodeUri() string`

GetQrCodeUri returns the QrCodeUri field if non-nil, zero value otherwise.

### GetQrCodeUriOk

`func (o *TwoFactorAuthSetupResult) GetQrCodeUriOk() (*string, bool)`

GetQrCodeUriOk returns a tuple with the QrCodeUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrCodeUri

`func (o *TwoFactorAuthSetupResult) SetQrCodeUri(v string)`

SetQrCodeUri sets QrCodeUri field to given value.


### SetQrCodeUriNil

`func (o *TwoFactorAuthSetupResult) SetQrCodeUriNil(b bool)`

 SetQrCodeUriNil sets the value for QrCodeUri to be an explicit nil

### UnsetQrCodeUri
`func (o *TwoFactorAuthSetupResult) UnsetQrCodeUri()`

UnsetQrCodeUri ensures that no value is present for QrCodeUri, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


