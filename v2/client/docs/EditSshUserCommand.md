# EditSshUserCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**SshPublicKey** | Pointer to **NullableString** |  | [optional] 
**AccessProfileId** | Pointer to **int32** |  | [optional] 

## Methods

### NewEditSshUserCommand

`func NewEditSshUserCommand() *EditSshUserCommand`

NewEditSshUserCommand instantiates a new EditSshUserCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditSshUserCommandWithDefaults

`func NewEditSshUserCommandWithDefaults() *EditSshUserCommand`

NewEditSshUserCommandWithDefaults instantiates a new EditSshUserCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EditSshUserCommand) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EditSshUserCommand) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EditSshUserCommand) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EditSshUserCommand) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *EditSshUserCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EditSshUserCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EditSshUserCommand) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EditSshUserCommand) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *EditSshUserCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *EditSshUserCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSshPublicKey

`func (o *EditSshUserCommand) GetSshPublicKey() string`

GetSshPublicKey returns the SshPublicKey field if non-nil, zero value otherwise.

### GetSshPublicKeyOk

`func (o *EditSshUserCommand) GetSshPublicKeyOk() (*string, bool)`

GetSshPublicKeyOk returns a tuple with the SshPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPublicKey

`func (o *EditSshUserCommand) SetSshPublicKey(v string)`

SetSshPublicKey sets SshPublicKey field to given value.

### HasSshPublicKey

`func (o *EditSshUserCommand) HasSshPublicKey() bool`

HasSshPublicKey returns a boolean if a field has been set.

### SetSshPublicKeyNil

`func (o *EditSshUserCommand) SetSshPublicKeyNil(b bool)`

 SetSshPublicKeyNil sets the value for SshPublicKey to be an explicit nil

### UnsetSshPublicKey
`func (o *EditSshUserCommand) UnsetSshPublicKey()`

UnsetSshPublicKey ensures that no value is present for SshPublicKey, not even an explicit nil
### GetAccessProfileId

`func (o *EditSshUserCommand) GetAccessProfileId() int32`

GetAccessProfileId returns the AccessProfileId field if non-nil, zero value otherwise.

### GetAccessProfileIdOk

`func (o *EditSshUserCommand) GetAccessProfileIdOk() (*int32, bool)`

GetAccessProfileIdOk returns a tuple with the AccessProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessProfileId

`func (o *EditSshUserCommand) SetAccessProfileId(v int32)`

SetAccessProfileId sets AccessProfileId field to given value.

### HasAccessProfileId

`func (o *EditSshUserCommand) HasAccessProfileId() bool`

HasAccessProfileId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


