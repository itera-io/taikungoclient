# CertificateProfileDropdownDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **NullableString** |  | 
**Kind** | Pointer to **NullableString** |  | [optional] 
**IsDefault** | **bool** |  | 
**NotAfter** | Pointer to **NullableTime** |  | [optional] 
**ExpiryStatus** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCertificateProfileDropdownDto

`func NewCertificateProfileDropdownDto(id int32, name NullableString, isDefault bool, ) *CertificateProfileDropdownDto`

NewCertificateProfileDropdownDto instantiates a new CertificateProfileDropdownDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateProfileDropdownDtoWithDefaults

`func NewCertificateProfileDropdownDtoWithDefaults() *CertificateProfileDropdownDto`

NewCertificateProfileDropdownDtoWithDefaults instantiates a new CertificateProfileDropdownDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CertificateProfileDropdownDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificateProfileDropdownDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificateProfileDropdownDto) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *CertificateProfileDropdownDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertificateProfileDropdownDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertificateProfileDropdownDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CertificateProfileDropdownDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CertificateProfileDropdownDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetKind

`func (o *CertificateProfileDropdownDto) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CertificateProfileDropdownDto) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CertificateProfileDropdownDto) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CertificateProfileDropdownDto) HasKind() bool`

HasKind returns a boolean if a field has been set.

### SetKindNil

`func (o *CertificateProfileDropdownDto) SetKindNil(b bool)`

 SetKindNil sets the value for Kind to be an explicit nil

### UnsetKind
`func (o *CertificateProfileDropdownDto) UnsetKind()`

UnsetKind ensures that no value is present for Kind, not even an explicit nil
### GetIsDefault

`func (o *CertificateProfileDropdownDto) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *CertificateProfileDropdownDto) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *CertificateProfileDropdownDto) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetNotAfter

`func (o *CertificateProfileDropdownDto) GetNotAfter() time.Time`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *CertificateProfileDropdownDto) GetNotAfterOk() (*time.Time, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *CertificateProfileDropdownDto) SetNotAfter(v time.Time)`

SetNotAfter sets NotAfter field to given value.

### HasNotAfter

`func (o *CertificateProfileDropdownDto) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.

### SetNotAfterNil

`func (o *CertificateProfileDropdownDto) SetNotAfterNil(b bool)`

 SetNotAfterNil sets the value for NotAfter to be an explicit nil

### UnsetNotAfter
`func (o *CertificateProfileDropdownDto) UnsetNotAfter()`

UnsetNotAfter ensures that no value is present for NotAfter, not even an explicit nil
### GetExpiryStatus

`func (o *CertificateProfileDropdownDto) GetExpiryStatus() string`

GetExpiryStatus returns the ExpiryStatus field if non-nil, zero value otherwise.

### GetExpiryStatusOk

`func (o *CertificateProfileDropdownDto) GetExpiryStatusOk() (*string, bool)`

GetExpiryStatusOk returns a tuple with the ExpiryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryStatus

`func (o *CertificateProfileDropdownDto) SetExpiryStatus(v string)`

SetExpiryStatus sets ExpiryStatus field to given value.

### HasExpiryStatus

`func (o *CertificateProfileDropdownDto) HasExpiryStatus() bool`

HasExpiryStatus returns a boolean if a field has been set.

### SetExpiryStatusNil

`func (o *CertificateProfileDropdownDto) SetExpiryStatusNil(b bool)`

 SetExpiryStatusNil sets the value for ExpiryStatus to be an explicit nil

### UnsetExpiryStatus
`func (o *CertificateProfileDropdownDto) UnsetExpiryStatus()`

UnsetExpiryStatus ensures that no value is present for ExpiryStatus, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


