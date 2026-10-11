# SshUsersListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **NullableString** |  | 
**SshPublicKey** | **NullableString** |  | 
**AccessProfileName** | **NullableString** |  | 

## Methods

### NewSshUsersListDto

`func NewSshUsersListDto(id int32, name NullableString, sshPublicKey NullableString, accessProfileName NullableString, ) *SshUsersListDto`

NewSshUsersListDto instantiates a new SshUsersListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshUsersListDtoWithDefaults

`func NewSshUsersListDtoWithDefaults() *SshUsersListDto`

NewSshUsersListDtoWithDefaults instantiates a new SshUsersListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SshUsersListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SshUsersListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SshUsersListDto) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *SshUsersListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SshUsersListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SshUsersListDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *SshUsersListDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SshUsersListDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSshPublicKey

`func (o *SshUsersListDto) GetSshPublicKey() string`

GetSshPublicKey returns the SshPublicKey field if non-nil, zero value otherwise.

### GetSshPublicKeyOk

`func (o *SshUsersListDto) GetSshPublicKeyOk() (*string, bool)`

GetSshPublicKeyOk returns a tuple with the SshPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPublicKey

`func (o *SshUsersListDto) SetSshPublicKey(v string)`

SetSshPublicKey sets SshPublicKey field to given value.


### SetSshPublicKeyNil

`func (o *SshUsersListDto) SetSshPublicKeyNil(b bool)`

 SetSshPublicKeyNil sets the value for SshPublicKey to be an explicit nil

### UnsetSshPublicKey
`func (o *SshUsersListDto) UnsetSshPublicKey()`

UnsetSshPublicKey ensures that no value is present for SshPublicKey, not even an explicit nil
### GetAccessProfileName

`func (o *SshUsersListDto) GetAccessProfileName() string`

GetAccessProfileName returns the AccessProfileName field if non-nil, zero value otherwise.

### GetAccessProfileNameOk

`func (o *SshUsersListDto) GetAccessProfileNameOk() (*string, bool)`

GetAccessProfileNameOk returns a tuple with the AccessProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessProfileName

`func (o *SshUsersListDto) SetAccessProfileName(v string)`

SetAccessProfileName sets AccessProfileName field to given value.


### SetAccessProfileNameNil

`func (o *SshUsersListDto) SetAccessProfileNameNil(b bool)`

 SetAccessProfileNameNil sets the value for AccessProfileName to be an explicit nil

### UnsetAccessProfileName
`func (o *SshUsersListDto) UnsetAccessProfileName()`

UnsetAccessProfileName ensures that no value is present for AccessProfileName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


