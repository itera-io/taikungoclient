# StandAloneProfileForDetailsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **NullableString** |  | 
**PublicKey** | **NullableString** |  | 
**SecurityGroups** | [**[]StandAloneProfileSecurityGroupForDetailsDto**](StandAloneProfileSecurityGroupForDetailsDto.md) |  | 

## Methods

### NewStandAloneProfileForDetailsDto

`func NewStandAloneProfileForDetailsDto(id int32, name NullableString, publicKey NullableString, securityGroups []StandAloneProfileSecurityGroupForDetailsDto, ) *StandAloneProfileForDetailsDto`

NewStandAloneProfileForDetailsDto instantiates a new StandAloneProfileForDetailsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandAloneProfileForDetailsDtoWithDefaults

`func NewStandAloneProfileForDetailsDtoWithDefaults() *StandAloneProfileForDetailsDto`

NewStandAloneProfileForDetailsDtoWithDefaults instantiates a new StandAloneProfileForDetailsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StandAloneProfileForDetailsDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StandAloneProfileForDetailsDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StandAloneProfileForDetailsDto) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *StandAloneProfileForDetailsDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StandAloneProfileForDetailsDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StandAloneProfileForDetailsDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *StandAloneProfileForDetailsDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *StandAloneProfileForDetailsDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPublicKey

`func (o *StandAloneProfileForDetailsDto) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *StandAloneProfileForDetailsDto) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *StandAloneProfileForDetailsDto) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


### SetPublicKeyNil

`func (o *StandAloneProfileForDetailsDto) SetPublicKeyNil(b bool)`

 SetPublicKeyNil sets the value for PublicKey to be an explicit nil

### UnsetPublicKey
`func (o *StandAloneProfileForDetailsDto) UnsetPublicKey()`

UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
### GetSecurityGroups

`func (o *StandAloneProfileForDetailsDto) GetSecurityGroups() []StandAloneProfileSecurityGroupForDetailsDto`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *StandAloneProfileForDetailsDto) GetSecurityGroupsOk() (*[]StandAloneProfileSecurityGroupForDetailsDto, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *StandAloneProfileForDetailsDto) SetSecurityGroups(v []StandAloneProfileSecurityGroupForDetailsDto)`

SetSecurityGroups sets SecurityGroups field to given value.


### SetSecurityGroupsNil

`func (o *StandAloneProfileForDetailsDto) SetSecurityGroupsNil(b bool)`

 SetSecurityGroupsNil sets the value for SecurityGroups to be an explicit nil

### UnsetSecurityGroups
`func (o *StandAloneProfileForDetailsDto) UnsetSecurityGroups()`

UnsetSecurityGroups ensures that no value is present for SecurityGroups, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


