# DomainDetailsForUserDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainId** | **int32** |  | 
**Name** | **string** |  | 
**Logo** | **NullableString** |  | 
**Domain** | **NullableString** |  | 

## Methods

### NewDomainDetailsForUserDto

`func NewDomainDetailsForUserDto(domainId int32, name string, logo NullableString, domain NullableString, ) *DomainDetailsForUserDto`

NewDomainDetailsForUserDto instantiates a new DomainDetailsForUserDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainDetailsForUserDtoWithDefaults

`func NewDomainDetailsForUserDtoWithDefaults() *DomainDetailsForUserDto`

NewDomainDetailsForUserDtoWithDefaults instantiates a new DomainDetailsForUserDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainId

`func (o *DomainDetailsForUserDto) GetDomainId() int32`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *DomainDetailsForUserDto) GetDomainIdOk() (*int32, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *DomainDetailsForUserDto) SetDomainId(v int32)`

SetDomainId sets DomainId field to given value.


### GetName

`func (o *DomainDetailsForUserDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DomainDetailsForUserDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DomainDetailsForUserDto) SetName(v string)`

SetName sets Name field to given value.


### GetLogo

`func (o *DomainDetailsForUserDto) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *DomainDetailsForUserDto) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *DomainDetailsForUserDto) SetLogo(v string)`

SetLogo sets Logo field to given value.


### SetLogoNil

`func (o *DomainDetailsForUserDto) SetLogoNil(b bool)`

 SetLogoNil sets the value for Logo to be an explicit nil

### UnsetLogo
`func (o *DomainDetailsForUserDto) UnsetLogo()`

UnsetLogo ensures that no value is present for Logo, not even an explicit nil
### GetDomain

`func (o *DomainDetailsForUserDto) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DomainDetailsForUserDto) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DomainDetailsForUserDto) SetDomain(v string)`

SetDomain sets Domain field to given value.


### SetDomainNil

`func (o *DomainDetailsForUserDto) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *DomainDetailsForUserDto) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


