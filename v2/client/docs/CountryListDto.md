# CountryListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** |  | 
**IsEu** | **bool** |  | 

## Methods

### NewCountryListDto

`func NewCountryListDto(name NullableString, isEu bool, ) *CountryListDto`

NewCountryListDto instantiates a new CountryListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountryListDtoWithDefaults

`func NewCountryListDtoWithDefaults() *CountryListDto`

NewCountryListDtoWithDefaults instantiates a new CountryListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CountryListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CountryListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CountryListDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CountryListDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CountryListDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsEu

`func (o *CountryListDto) GetIsEu() bool`

GetIsEu returns the IsEu field if non-nil, zero value otherwise.

### GetIsEuOk

`func (o *CountryListDto) GetIsEuOk() (*bool, bool)`

GetIsEuOk returns a tuple with the IsEu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEu

`func (o *CountryListDto) SetIsEu(v bool)`

SetIsEu sets IsEu field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


