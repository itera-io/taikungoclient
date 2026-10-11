# OrganizationDropdownDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **NullableString** |  | 
**IsInfra** | **bool** |  | 
**DiscountRate** | **float64** |  | 
**DomainId** | **int32** |  | 
**DomainName** | **NullableString** |  | 

## Methods

### NewOrganizationDropdownDto

`func NewOrganizationDropdownDto(id int32, name NullableString, isInfra bool, discountRate float64, domainId int32, domainName NullableString, ) *OrganizationDropdownDto`

NewOrganizationDropdownDto instantiates a new OrganizationDropdownDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationDropdownDtoWithDefaults

`func NewOrganizationDropdownDtoWithDefaults() *OrganizationDropdownDto`

NewOrganizationDropdownDtoWithDefaults instantiates a new OrganizationDropdownDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationDropdownDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationDropdownDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationDropdownDto) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *OrganizationDropdownDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationDropdownDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationDropdownDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *OrganizationDropdownDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *OrganizationDropdownDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsInfra

`func (o *OrganizationDropdownDto) GetIsInfra() bool`

GetIsInfra returns the IsInfra field if non-nil, zero value otherwise.

### GetIsInfraOk

`func (o *OrganizationDropdownDto) GetIsInfraOk() (*bool, bool)`

GetIsInfraOk returns a tuple with the IsInfra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInfra

`func (o *OrganizationDropdownDto) SetIsInfra(v bool)`

SetIsInfra sets IsInfra field to given value.


### GetDiscountRate

`func (o *OrganizationDropdownDto) GetDiscountRate() float64`

GetDiscountRate returns the DiscountRate field if non-nil, zero value otherwise.

### GetDiscountRateOk

`func (o *OrganizationDropdownDto) GetDiscountRateOk() (*float64, bool)`

GetDiscountRateOk returns a tuple with the DiscountRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountRate

`func (o *OrganizationDropdownDto) SetDiscountRate(v float64)`

SetDiscountRate sets DiscountRate field to given value.


### GetDomainId

`func (o *OrganizationDropdownDto) GetDomainId() int32`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *OrganizationDropdownDto) GetDomainIdOk() (*int32, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *OrganizationDropdownDto) SetDomainId(v int32)`

SetDomainId sets DomainId field to given value.


### GetDomainName

`func (o *OrganizationDropdownDto) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *OrganizationDropdownDto) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *OrganizationDropdownDto) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.


### SetDomainNameNil

`func (o *OrganizationDropdownDto) SetDomainNameNil(b bool)`

 SetDomainNameNil sets the value for DomainName to be an explicit nil

### UnsetDomainName
`func (o *OrganizationDropdownDto) UnsetDomainName()`

UnsetDomainName ensures that no value is present for DomainName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


