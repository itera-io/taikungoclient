# UpdateGroupDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**ClaimValue** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUpdateGroupDto

`func NewUpdateGroupDto() *UpdateGroupDto`

NewUpdateGroupDto instantiates a new UpdateGroupDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGroupDtoWithDefaults

`func NewUpdateGroupDtoWithDefaults() *UpdateGroupDto`

NewUpdateGroupDtoWithDefaults instantiates a new UpdateGroupDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateGroupDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateGroupDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateGroupDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateGroupDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateGroupDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateGroupDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetClaimValue

`func (o *UpdateGroupDto) GetClaimValue() string`

GetClaimValue returns the ClaimValue field if non-nil, zero value otherwise.

### GetClaimValueOk

`func (o *UpdateGroupDto) GetClaimValueOk() (*string, bool)`

GetClaimValueOk returns a tuple with the ClaimValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimValue

`func (o *UpdateGroupDto) SetClaimValue(v string)`

SetClaimValue sets ClaimValue field to given value.

### HasClaimValue

`func (o *UpdateGroupDto) HasClaimValue() bool`

HasClaimValue returns a boolean if a field has been set.

### SetClaimValueNil

`func (o *UpdateGroupDto) SetClaimValueNil(b bool)`

 SetClaimValueNil sets the value for ClaimValue to be an explicit nil

### UnsetClaimValue
`func (o *UpdateGroupDto) UnsetClaimValue()`

UnsetClaimValue ensures that no value is present for ClaimValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


