# PackageAutocompleteDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **NullableString** |  | 
**Value** | **NullableString** |  | 
**Description** | **NullableString** |  | 
**Type** | [**ParameterType**](ParameterType.md) |  | 
**IsQuestion** | **bool** |  | 
**Options** | **[]string** |  | 
**IsTaikunLink** | **bool** |  | 

## Methods

### NewPackageAutocompleteDto

`func NewPackageAutocompleteDto(key NullableString, value NullableString, description NullableString, type_ ParameterType, isQuestion bool, options []string, isTaikunLink bool, ) *PackageAutocompleteDto`

NewPackageAutocompleteDto instantiates a new PackageAutocompleteDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageAutocompleteDtoWithDefaults

`func NewPackageAutocompleteDtoWithDefaults() *PackageAutocompleteDto`

NewPackageAutocompleteDtoWithDefaults instantiates a new PackageAutocompleteDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *PackageAutocompleteDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PackageAutocompleteDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PackageAutocompleteDto) SetKey(v string)`

SetKey sets Key field to given value.


### SetKeyNil

`func (o *PackageAutocompleteDto) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *PackageAutocompleteDto) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetValue

`func (o *PackageAutocompleteDto) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PackageAutocompleteDto) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PackageAutocompleteDto) SetValue(v string)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *PackageAutocompleteDto) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *PackageAutocompleteDto) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetDescription

`func (o *PackageAutocompleteDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PackageAutocompleteDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PackageAutocompleteDto) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *PackageAutocompleteDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PackageAutocompleteDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetType

`func (o *PackageAutocompleteDto) GetType() ParameterType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PackageAutocompleteDto) GetTypeOk() (*ParameterType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PackageAutocompleteDto) SetType(v ParameterType)`

SetType sets Type field to given value.


### GetIsQuestion

`func (o *PackageAutocompleteDto) GetIsQuestion() bool`

GetIsQuestion returns the IsQuestion field if non-nil, zero value otherwise.

### GetIsQuestionOk

`func (o *PackageAutocompleteDto) GetIsQuestionOk() (*bool, bool)`

GetIsQuestionOk returns a tuple with the IsQuestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQuestion

`func (o *PackageAutocompleteDto) SetIsQuestion(v bool)`

SetIsQuestion sets IsQuestion field to given value.


### GetOptions

`func (o *PackageAutocompleteDto) GetOptions() []string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *PackageAutocompleteDto) GetOptionsOk() (*[]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *PackageAutocompleteDto) SetOptions(v []string)`

SetOptions sets Options field to given value.


### SetOptionsNil

`func (o *PackageAutocompleteDto) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *PackageAutocompleteDto) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetIsTaikunLink

`func (o *PackageAutocompleteDto) GetIsTaikunLink() bool`

GetIsTaikunLink returns the IsTaikunLink field if non-nil, zero value otherwise.

### GetIsTaikunLinkOk

`func (o *PackageAutocompleteDto) GetIsTaikunLinkOk() (*bool, bool)`

GetIsTaikunLinkOk returns a tuple with the IsTaikunLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTaikunLink

`func (o *PackageAutocompleteDto) SetIsTaikunLink(v bool)`

SetIsTaikunLink sets IsTaikunLink field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


