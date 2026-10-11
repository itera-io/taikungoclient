# UpdateAirgapGlobalConfigurationCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **NullableString** |  | 
**OciRegistryHost** | **NullableString** |  | 
**OciPlainHttp** | **bool** |  | 
**Username** | **NullableString** |  | 
**Password** | **NullableString** |  | 

## Methods

### NewUpdateAirgapGlobalConfigurationCommand

`func NewUpdateAirgapGlobalConfigurationCommand(id int32, name NullableString, ociRegistryHost NullableString, ociPlainHttp bool, username NullableString, password NullableString, ) *UpdateAirgapGlobalConfigurationCommand`

NewUpdateAirgapGlobalConfigurationCommand instantiates a new UpdateAirgapGlobalConfigurationCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAirgapGlobalConfigurationCommandWithDefaults

`func NewUpdateAirgapGlobalConfigurationCommandWithDefaults() *UpdateAirgapGlobalConfigurationCommand`

NewUpdateAirgapGlobalConfigurationCommandWithDefaults instantiates a new UpdateAirgapGlobalConfigurationCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateAirgapGlobalConfigurationCommand) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateAirgapGlobalConfigurationCommand) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateAirgapGlobalConfigurationCommand) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *UpdateAirgapGlobalConfigurationCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAirgapGlobalConfigurationCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAirgapGlobalConfigurationCommand) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *UpdateAirgapGlobalConfigurationCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateAirgapGlobalConfigurationCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOciRegistryHost

`func (o *UpdateAirgapGlobalConfigurationCommand) GetOciRegistryHost() string`

GetOciRegistryHost returns the OciRegistryHost field if non-nil, zero value otherwise.

### GetOciRegistryHostOk

`func (o *UpdateAirgapGlobalConfigurationCommand) GetOciRegistryHostOk() (*string, bool)`

GetOciRegistryHostOk returns a tuple with the OciRegistryHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOciRegistryHost

`func (o *UpdateAirgapGlobalConfigurationCommand) SetOciRegistryHost(v string)`

SetOciRegistryHost sets OciRegistryHost field to given value.


### SetOciRegistryHostNil

`func (o *UpdateAirgapGlobalConfigurationCommand) SetOciRegistryHostNil(b bool)`

 SetOciRegistryHostNil sets the value for OciRegistryHost to be an explicit nil

### UnsetOciRegistryHost
`func (o *UpdateAirgapGlobalConfigurationCommand) UnsetOciRegistryHost()`

UnsetOciRegistryHost ensures that no value is present for OciRegistryHost, not even an explicit nil
### GetOciPlainHttp

`func (o *UpdateAirgapGlobalConfigurationCommand) GetOciPlainHttp() bool`

GetOciPlainHttp returns the OciPlainHttp field if non-nil, zero value otherwise.

### GetOciPlainHttpOk

`func (o *UpdateAirgapGlobalConfigurationCommand) GetOciPlainHttpOk() (*bool, bool)`

GetOciPlainHttpOk returns a tuple with the OciPlainHttp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOciPlainHttp

`func (o *UpdateAirgapGlobalConfigurationCommand) SetOciPlainHttp(v bool)`

SetOciPlainHttp sets OciPlainHttp field to given value.


### GetUsername

`func (o *UpdateAirgapGlobalConfigurationCommand) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateAirgapGlobalConfigurationCommand) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateAirgapGlobalConfigurationCommand) SetUsername(v string)`

SetUsername sets Username field to given value.


### SetUsernameNil

`func (o *UpdateAirgapGlobalConfigurationCommand) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *UpdateAirgapGlobalConfigurationCommand) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *UpdateAirgapGlobalConfigurationCommand) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateAirgapGlobalConfigurationCommand) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateAirgapGlobalConfigurationCommand) SetPassword(v string)`

SetPassword sets Password field to given value.


### SetPasswordNil

`func (o *UpdateAirgapGlobalConfigurationCommand) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *UpdateAirgapGlobalConfigurationCommand) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


