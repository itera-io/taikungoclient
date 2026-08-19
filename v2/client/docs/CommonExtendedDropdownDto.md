# CommonExtendedDropdownDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **NullableString** |  | 
**IsDefault** | **bool** |  | 

## Methods

### NewCommonExtendedDropdownDto

`func NewCommonExtendedDropdownDto(id int32, name NullableString, isDefault bool, ) *CommonExtendedDropdownDto`

NewCommonExtendedDropdownDto instantiates a new CommonExtendedDropdownDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonExtendedDropdownDtoWithDefaults

`func NewCommonExtendedDropdownDtoWithDefaults() *CommonExtendedDropdownDto`

NewCommonExtendedDropdownDtoWithDefaults instantiates a new CommonExtendedDropdownDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommonExtendedDropdownDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommonExtendedDropdownDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommonExtendedDropdownDto) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *CommonExtendedDropdownDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CommonExtendedDropdownDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CommonExtendedDropdownDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CommonExtendedDropdownDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CommonExtendedDropdownDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsDefault

`func (o *CommonExtendedDropdownDto) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *CommonExtendedDropdownDto) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *CommonExtendedDropdownDto) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


