# ProjectAppParamDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **NullableString** |  | 
**Value** | **NullableString** |  | 
**IsEditableWhenInstalling** | **bool** |  | 
**IsEditableAfterInstallation** | **bool** |  | 
**IsMandatory** | **bool** |  | 

## Methods

### NewProjectAppParamDto

`func NewProjectAppParamDto(key NullableString, value NullableString, isEditableWhenInstalling bool, isEditableAfterInstallation bool, isMandatory bool, ) *ProjectAppParamDto`

NewProjectAppParamDto instantiates a new ProjectAppParamDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectAppParamDtoWithDefaults

`func NewProjectAppParamDtoWithDefaults() *ProjectAppParamDto`

NewProjectAppParamDtoWithDefaults instantiates a new ProjectAppParamDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ProjectAppParamDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ProjectAppParamDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ProjectAppParamDto) SetKey(v string)`

SetKey sets Key field to given value.


### SetKeyNil

`func (o *ProjectAppParamDto) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *ProjectAppParamDto) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetValue

`func (o *ProjectAppParamDto) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ProjectAppParamDto) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ProjectAppParamDto) SetValue(v string)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *ProjectAppParamDto) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ProjectAppParamDto) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetIsEditableWhenInstalling

`func (o *ProjectAppParamDto) GetIsEditableWhenInstalling() bool`

GetIsEditableWhenInstalling returns the IsEditableWhenInstalling field if non-nil, zero value otherwise.

### GetIsEditableWhenInstallingOk

`func (o *ProjectAppParamDto) GetIsEditableWhenInstallingOk() (*bool, bool)`

GetIsEditableWhenInstallingOk returns a tuple with the IsEditableWhenInstalling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEditableWhenInstalling

`func (o *ProjectAppParamDto) SetIsEditableWhenInstalling(v bool)`

SetIsEditableWhenInstalling sets IsEditableWhenInstalling field to given value.


### GetIsEditableAfterInstallation

`func (o *ProjectAppParamDto) GetIsEditableAfterInstallation() bool`

GetIsEditableAfterInstallation returns the IsEditableAfterInstallation field if non-nil, zero value otherwise.

### GetIsEditableAfterInstallationOk

`func (o *ProjectAppParamDto) GetIsEditableAfterInstallationOk() (*bool, bool)`

GetIsEditableAfterInstallationOk returns a tuple with the IsEditableAfterInstallation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEditableAfterInstallation

`func (o *ProjectAppParamDto) SetIsEditableAfterInstallation(v bool)`

SetIsEditableAfterInstallation sets IsEditableAfterInstallation field to given value.


### GetIsMandatory

`func (o *ProjectAppParamDto) GetIsMandatory() bool`

GetIsMandatory returns the IsMandatory field if non-nil, zero value otherwise.

### GetIsMandatoryOk

`func (o *ProjectAppParamDto) GetIsMandatoryOk() (*bool, bool)`

GetIsMandatoryOk returns a tuple with the IsMandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMandatory

`func (o *ProjectAppParamDto) SetIsMandatory(v bool)`

SetIsMandatory sets IsMandatory field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


