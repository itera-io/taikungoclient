# GetToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** |  | 
**RefreshToken** | Pointer to **NullableString** |  | [optional] 
**RefreshTokenExpireTime** | Pointer to **time.Time** |  | [optional] 
**TwoFaEnabled** | Pointer to **NullableBool** |  | [optional] 
**IsForcedToEnabled2Fa** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewGetToken

`func NewGetToken(token string, ) *GetToken`

NewGetToken instantiates a new GetToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTokenWithDefaults

`func NewGetTokenWithDefaults() *GetToken`

NewGetTokenWithDefaults instantiates a new GetToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *GetToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GetToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GetToken) SetToken(v string)`

SetToken sets Token field to given value.


### GetRefreshToken

`func (o *GetToken) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *GetToken) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *GetToken) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *GetToken) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### SetRefreshTokenNil

`func (o *GetToken) SetRefreshTokenNil(b bool)`

 SetRefreshTokenNil sets the value for RefreshToken to be an explicit nil

### UnsetRefreshToken
`func (o *GetToken) UnsetRefreshToken()`

UnsetRefreshToken ensures that no value is present for RefreshToken, not even an explicit nil
### GetRefreshTokenExpireTime

`func (o *GetToken) GetRefreshTokenExpireTime() time.Time`

GetRefreshTokenExpireTime returns the RefreshTokenExpireTime field if non-nil, zero value otherwise.

### GetRefreshTokenExpireTimeOk

`func (o *GetToken) GetRefreshTokenExpireTimeOk() (*time.Time, bool)`

GetRefreshTokenExpireTimeOk returns a tuple with the RefreshTokenExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenExpireTime

`func (o *GetToken) SetRefreshTokenExpireTime(v time.Time)`

SetRefreshTokenExpireTime sets RefreshTokenExpireTime field to given value.

### HasRefreshTokenExpireTime

`func (o *GetToken) HasRefreshTokenExpireTime() bool`

HasRefreshTokenExpireTime returns a boolean if a field has been set.

### GetTwoFaEnabled

`func (o *GetToken) GetTwoFaEnabled() bool`

GetTwoFaEnabled returns the TwoFaEnabled field if non-nil, zero value otherwise.

### GetTwoFaEnabledOk

`func (o *GetToken) GetTwoFaEnabledOk() (*bool, bool)`

GetTwoFaEnabledOk returns a tuple with the TwoFaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFaEnabled

`func (o *GetToken) SetTwoFaEnabled(v bool)`

SetTwoFaEnabled sets TwoFaEnabled field to given value.

### HasTwoFaEnabled

`func (o *GetToken) HasTwoFaEnabled() bool`

HasTwoFaEnabled returns a boolean if a field has been set.

### SetTwoFaEnabledNil

`func (o *GetToken) SetTwoFaEnabledNil(b bool)`

 SetTwoFaEnabledNil sets the value for TwoFaEnabled to be an explicit nil

### UnsetTwoFaEnabled
`func (o *GetToken) UnsetTwoFaEnabled()`

UnsetTwoFaEnabled ensures that no value is present for TwoFaEnabled, not even an explicit nil
### GetIsForcedToEnabled2Fa

`func (o *GetToken) GetIsForcedToEnabled2Fa() bool`

GetIsForcedToEnabled2Fa returns the IsForcedToEnabled2Fa field if non-nil, zero value otherwise.

### GetIsForcedToEnabled2FaOk

`func (o *GetToken) GetIsForcedToEnabled2FaOk() (*bool, bool)`

GetIsForcedToEnabled2FaOk returns a tuple with the IsForcedToEnabled2Fa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForcedToEnabled2Fa

`func (o *GetToken) SetIsForcedToEnabled2Fa(v bool)`

SetIsForcedToEnabled2Fa sets IsForcedToEnabled2Fa field to given value.

### HasIsForcedToEnabled2Fa

`func (o *GetToken) HasIsForcedToEnabled2Fa() bool`

HasIsForcedToEnabled2Fa returns a boolean if a field has been set.

### SetIsForcedToEnabled2FaNil

`func (o *GetToken) SetIsForcedToEnabled2FaNil(b bool)`

 SetIsForcedToEnabled2FaNil sets the value for IsForcedToEnabled2Fa to be an explicit nil

### UnsetIsForcedToEnabled2Fa
`func (o *GetToken) UnsetIsForcedToEnabled2Fa()`

UnsetIsForcedToEnabled2Fa ensures that no value is present for IsForcedToEnabled2Fa, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


