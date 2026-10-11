# CreateSshUserCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** |  | 
**SshPublicKey** | **NullableString** |  | 
**AccessProfileId** | **int32** |  | 

## Methods

### NewCreateSshUserCommand

`func NewCreateSshUserCommand(name NullableString, sshPublicKey NullableString, accessProfileId int32, ) *CreateSshUserCommand`

NewCreateSshUserCommand instantiates a new CreateSshUserCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSshUserCommandWithDefaults

`func NewCreateSshUserCommandWithDefaults() *CreateSshUserCommand`

NewCreateSshUserCommandWithDefaults instantiates a new CreateSshUserCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateSshUserCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSshUserCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSshUserCommand) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CreateSshUserCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateSshUserCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSshPublicKey

`func (o *CreateSshUserCommand) GetSshPublicKey() string`

GetSshPublicKey returns the SshPublicKey field if non-nil, zero value otherwise.

### GetSshPublicKeyOk

`func (o *CreateSshUserCommand) GetSshPublicKeyOk() (*string, bool)`

GetSshPublicKeyOk returns a tuple with the SshPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPublicKey

`func (o *CreateSshUserCommand) SetSshPublicKey(v string)`

SetSshPublicKey sets SshPublicKey field to given value.


### SetSshPublicKeyNil

`func (o *CreateSshUserCommand) SetSshPublicKeyNil(b bool)`

 SetSshPublicKeyNil sets the value for SshPublicKey to be an explicit nil

### UnsetSshPublicKey
`func (o *CreateSshUserCommand) UnsetSshPublicKey()`

UnsetSshPublicKey ensures that no value is present for SshPublicKey, not even an explicit nil
### GetAccessProfileId

`func (o *CreateSshUserCommand) GetAccessProfileId() int32`

GetAccessProfileId returns the AccessProfileId field if non-nil, zero value otherwise.

### GetAccessProfileIdOk

`func (o *CreateSshUserCommand) GetAccessProfileIdOk() (*int32, bool)`

GetAccessProfileIdOk returns a tuple with the AccessProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessProfileId

`func (o *CreateSshUserCommand) SetAccessProfileId(v int32)`

SetAccessProfileId sets AccessProfileId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


