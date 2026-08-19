# StandaloneProfileListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **NullableString** |  | 
**Revision** | **int32** |  | 
**IsLocked** | **bool** |  | 
**StandAloneProfileSecurityGroups** | [**[]StandaloneProfileSecurityGroupListDto**](StandaloneProfileSecurityGroupListDto.md) |  | 

## Methods

### NewStandaloneProfileListDto

`func NewStandaloneProfileListDto(id int32, name NullableString, revision int32, isLocked bool, standAloneProfileSecurityGroups []StandaloneProfileSecurityGroupListDto, ) *StandaloneProfileListDto`

NewStandaloneProfileListDto instantiates a new StandaloneProfileListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandaloneProfileListDtoWithDefaults

`func NewStandaloneProfileListDtoWithDefaults() *StandaloneProfileListDto`

NewStandaloneProfileListDtoWithDefaults instantiates a new StandaloneProfileListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StandaloneProfileListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StandaloneProfileListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StandaloneProfileListDto) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *StandaloneProfileListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StandaloneProfileListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StandaloneProfileListDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *StandaloneProfileListDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *StandaloneProfileListDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRevision

`func (o *StandaloneProfileListDto) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *StandaloneProfileListDto) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *StandaloneProfileListDto) SetRevision(v int32)`

SetRevision sets Revision field to given value.


### GetIsLocked

`func (o *StandaloneProfileListDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *StandaloneProfileListDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *StandaloneProfileListDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.


### GetStandAloneProfileSecurityGroups

`func (o *StandaloneProfileListDto) GetStandAloneProfileSecurityGroups() []StandaloneProfileSecurityGroupListDto`

GetStandAloneProfileSecurityGroups returns the StandAloneProfileSecurityGroups field if non-nil, zero value otherwise.

### GetStandAloneProfileSecurityGroupsOk

`func (o *StandaloneProfileListDto) GetStandAloneProfileSecurityGroupsOk() (*[]StandaloneProfileSecurityGroupListDto, bool)`

GetStandAloneProfileSecurityGroupsOk returns a tuple with the StandAloneProfileSecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandAloneProfileSecurityGroups

`func (o *StandaloneProfileListDto) SetStandAloneProfileSecurityGroups(v []StandaloneProfileSecurityGroupListDto)`

SetStandAloneProfileSecurityGroups sets StandAloneProfileSecurityGroups field to given value.


### SetStandAloneProfileSecurityGroupsNil

`func (o *StandaloneProfileListDto) SetStandAloneProfileSecurityGroupsNil(b bool)`

 SetStandAloneProfileSecurityGroupsNil sets the value for StandAloneProfileSecurityGroups to be an explicit nil

### UnsetStandAloneProfileSecurityGroups
`func (o *StandaloneProfileListDto) UnsetStandAloneProfileSecurityGroups()`

UnsetStandAloneProfileSecurityGroups ensures that no value is present for StandAloneProfileSecurityGroups, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


