# GroupedProjectDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **NullableString** |  | 
**Status** | [**ProjectStatus**](ProjectStatus.md) |  | 
**Health** | [**ProjectHealth**](ProjectHealth.md) |  | 

## Methods

### NewGroupedProjectDto

`func NewGroupedProjectDto(id int32, name NullableString, status ProjectStatus, health ProjectHealth, ) *GroupedProjectDto`

NewGroupedProjectDto instantiates a new GroupedProjectDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupedProjectDtoWithDefaults

`func NewGroupedProjectDtoWithDefaults() *GroupedProjectDto`

NewGroupedProjectDtoWithDefaults instantiates a new GroupedProjectDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupedProjectDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupedProjectDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupedProjectDto) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *GroupedProjectDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupedProjectDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupedProjectDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *GroupedProjectDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GroupedProjectDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStatus

`func (o *GroupedProjectDto) GetStatus() ProjectStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GroupedProjectDto) GetStatusOk() (*ProjectStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GroupedProjectDto) SetStatus(v ProjectStatus)`

SetStatus sets Status field to given value.


### GetHealth

`func (o *GroupedProjectDto) GetHealth() ProjectHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *GroupedProjectDto) GetHealthOk() (*ProjectHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *GroupedProjectDto) SetHealth(v ProjectHealth)`

SetHealth sets Health field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


