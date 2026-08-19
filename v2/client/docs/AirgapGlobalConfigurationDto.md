# AirgapGlobalConfigurationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**OciRegistryHost** | Pointer to **NullableString** |  | [optional] 
**OciPlainHttp** | Pointer to **bool** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAirgapGlobalConfigurationDto

`func NewAirgapGlobalConfigurationDto() *AirgapGlobalConfigurationDto`

NewAirgapGlobalConfigurationDto instantiates a new AirgapGlobalConfigurationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAirgapGlobalConfigurationDtoWithDefaults

`func NewAirgapGlobalConfigurationDtoWithDefaults() *AirgapGlobalConfigurationDto`

NewAirgapGlobalConfigurationDtoWithDefaults instantiates a new AirgapGlobalConfigurationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AirgapGlobalConfigurationDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AirgapGlobalConfigurationDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AirgapGlobalConfigurationDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AirgapGlobalConfigurationDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AirgapGlobalConfigurationDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AirgapGlobalConfigurationDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AirgapGlobalConfigurationDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AirgapGlobalConfigurationDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AirgapGlobalConfigurationDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AirgapGlobalConfigurationDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOciRegistryHost

`func (o *AirgapGlobalConfigurationDto) GetOciRegistryHost() string`

GetOciRegistryHost returns the OciRegistryHost field if non-nil, zero value otherwise.

### GetOciRegistryHostOk

`func (o *AirgapGlobalConfigurationDto) GetOciRegistryHostOk() (*string, bool)`

GetOciRegistryHostOk returns a tuple with the OciRegistryHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOciRegistryHost

`func (o *AirgapGlobalConfigurationDto) SetOciRegistryHost(v string)`

SetOciRegistryHost sets OciRegistryHost field to given value.

### HasOciRegistryHost

`func (o *AirgapGlobalConfigurationDto) HasOciRegistryHost() bool`

HasOciRegistryHost returns a boolean if a field has been set.

### SetOciRegistryHostNil

`func (o *AirgapGlobalConfigurationDto) SetOciRegistryHostNil(b bool)`

 SetOciRegistryHostNil sets the value for OciRegistryHost to be an explicit nil

### UnsetOciRegistryHost
`func (o *AirgapGlobalConfigurationDto) UnsetOciRegistryHost()`

UnsetOciRegistryHost ensures that no value is present for OciRegistryHost, not even an explicit nil
### GetOciPlainHttp

`func (o *AirgapGlobalConfigurationDto) GetOciPlainHttp() bool`

GetOciPlainHttp returns the OciPlainHttp field if non-nil, zero value otherwise.

### GetOciPlainHttpOk

`func (o *AirgapGlobalConfigurationDto) GetOciPlainHttpOk() (*bool, bool)`

GetOciPlainHttpOk returns a tuple with the OciPlainHttp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOciPlainHttp

`func (o *AirgapGlobalConfigurationDto) SetOciPlainHttp(v bool)`

SetOciPlainHttp sets OciPlainHttp field to given value.

### HasOciPlainHttp

`func (o *AirgapGlobalConfigurationDto) HasOciPlainHttp() bool`

HasOciPlainHttp returns a boolean if a field has been set.

### GetIsActive

`func (o *AirgapGlobalConfigurationDto) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *AirgapGlobalConfigurationDto) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *AirgapGlobalConfigurationDto) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *AirgapGlobalConfigurationDto) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetUsername

`func (o *AirgapGlobalConfigurationDto) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AirgapGlobalConfigurationDto) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AirgapGlobalConfigurationDto) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AirgapGlobalConfigurationDto) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *AirgapGlobalConfigurationDto) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *AirgapGlobalConfigurationDto) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *AirgapGlobalConfigurationDto) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AirgapGlobalConfigurationDto) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AirgapGlobalConfigurationDto) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AirgapGlobalConfigurationDto) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *AirgapGlobalConfigurationDto) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *AirgapGlobalConfigurationDto) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


