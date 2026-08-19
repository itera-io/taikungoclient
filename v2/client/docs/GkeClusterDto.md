# GkeClusterDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | **NullableString** |  | 
**Name** | **NullableString** |  | 

## Methods

### NewGkeClusterDto

`func NewGkeClusterDto(location NullableString, name NullableString, ) *GkeClusterDto`

NewGkeClusterDto instantiates a new GkeClusterDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGkeClusterDtoWithDefaults

`func NewGkeClusterDtoWithDefaults() *GkeClusterDto`

NewGkeClusterDtoWithDefaults instantiates a new GkeClusterDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *GkeClusterDto) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *GkeClusterDto) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *GkeClusterDto) SetLocation(v string)`

SetLocation sets Location field to given value.


### SetLocationNil

`func (o *GkeClusterDto) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *GkeClusterDto) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetName

`func (o *GkeClusterDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GkeClusterDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GkeClusterDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *GkeClusterDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GkeClusterDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


