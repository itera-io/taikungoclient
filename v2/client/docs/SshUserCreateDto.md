# SshUserCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**SshPublicKey** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSshUserCreateDto

`func NewSshUserCreateDto() *SshUserCreateDto`

NewSshUserCreateDto instantiates a new SshUserCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshUserCreateDtoWithDefaults

`func NewSshUserCreateDtoWithDefaults() *SshUserCreateDto`

NewSshUserCreateDtoWithDefaults instantiates a new SshUserCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SshUserCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SshUserCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SshUserCreateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SshUserCreateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SshUserCreateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SshUserCreateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSshPublicKey

`func (o *SshUserCreateDto) GetSshPublicKey() string`

GetSshPublicKey returns the SshPublicKey field if non-nil, zero value otherwise.

### GetSshPublicKeyOk

`func (o *SshUserCreateDto) GetSshPublicKeyOk() (*string, bool)`

GetSshPublicKeyOk returns a tuple with the SshPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPublicKey

`func (o *SshUserCreateDto) SetSshPublicKey(v string)`

SetSshPublicKey sets SshPublicKey field to given value.

### HasSshPublicKey

`func (o *SshUserCreateDto) HasSshPublicKey() bool`

HasSshPublicKey returns a boolean if a field has been set.

### SetSshPublicKeyNil

`func (o *SshUserCreateDto) SetSshPublicKeyNil(b bool)`

 SetSshPublicKeyNil sets the value for SshPublicKey to be an explicit nil

### UnsetSshPublicKey
`func (o *SshUserCreateDto) UnsetSshPublicKey()`

UnsetSshPublicKey ensures that no value is present for SshPublicKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


