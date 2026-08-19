# DomainEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainId** | **int32** |  | 
**DomainName** | **string** |  | 
**Logo** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDomainEntity

`func NewDomainEntity(domainId int32, domainName string, ) *DomainEntity`

NewDomainEntity instantiates a new DomainEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainEntityWithDefaults

`func NewDomainEntityWithDefaults() *DomainEntity`

NewDomainEntityWithDefaults instantiates a new DomainEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainId

`func (o *DomainEntity) GetDomainId() int32`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *DomainEntity) GetDomainIdOk() (*int32, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *DomainEntity) SetDomainId(v int32)`

SetDomainId sets DomainId field to given value.


### GetDomainName

`func (o *DomainEntity) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *DomainEntity) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *DomainEntity) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.


### GetLogo

`func (o *DomainEntity) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *DomainEntity) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *DomainEntity) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *DomainEntity) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### SetLogoNil

`func (o *DomainEntity) SetLogoNil(b bool)`

 SetLogoNil sets the value for Logo to be an explicit nil

### UnsetLogo
`func (o *DomainEntity) UnsetLogo()`

UnsetLogo ensures that no value is present for Logo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


