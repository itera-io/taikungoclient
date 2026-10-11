# DnsCredentialDropdownDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **NullableString** |  | 
**ProviderType** | **NullableString** |  | 
**DomainFilter** | **NullableString** |  | 
**IsDefault** | **bool** |  | 
**IsLocked** | **bool** |  | 

## Methods

### NewDnsCredentialDropdownDto

`func NewDnsCredentialDropdownDto(id int32, name NullableString, providerType NullableString, domainFilter NullableString, isDefault bool, isLocked bool, ) *DnsCredentialDropdownDto`

NewDnsCredentialDropdownDto instantiates a new DnsCredentialDropdownDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsCredentialDropdownDtoWithDefaults

`func NewDnsCredentialDropdownDtoWithDefaults() *DnsCredentialDropdownDto`

NewDnsCredentialDropdownDtoWithDefaults instantiates a new DnsCredentialDropdownDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DnsCredentialDropdownDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DnsCredentialDropdownDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DnsCredentialDropdownDto) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *DnsCredentialDropdownDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DnsCredentialDropdownDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DnsCredentialDropdownDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *DnsCredentialDropdownDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DnsCredentialDropdownDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProviderType

`func (o *DnsCredentialDropdownDto) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *DnsCredentialDropdownDto) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *DnsCredentialDropdownDto) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.


### SetProviderTypeNil

`func (o *DnsCredentialDropdownDto) SetProviderTypeNil(b bool)`

 SetProviderTypeNil sets the value for ProviderType to be an explicit nil

### UnsetProviderType
`func (o *DnsCredentialDropdownDto) UnsetProviderType()`

UnsetProviderType ensures that no value is present for ProviderType, not even an explicit nil
### GetDomainFilter

`func (o *DnsCredentialDropdownDto) GetDomainFilter() string`

GetDomainFilter returns the DomainFilter field if non-nil, zero value otherwise.

### GetDomainFilterOk

`func (o *DnsCredentialDropdownDto) GetDomainFilterOk() (*string, bool)`

GetDomainFilterOk returns a tuple with the DomainFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainFilter

`func (o *DnsCredentialDropdownDto) SetDomainFilter(v string)`

SetDomainFilter sets DomainFilter field to given value.


### SetDomainFilterNil

`func (o *DnsCredentialDropdownDto) SetDomainFilterNil(b bool)`

 SetDomainFilterNil sets the value for DomainFilter to be an explicit nil

### UnsetDomainFilter
`func (o *DnsCredentialDropdownDto) UnsetDomainFilter()`

UnsetDomainFilter ensures that no value is present for DomainFilter, not even an explicit nil
### GetIsDefault

`func (o *DnsCredentialDropdownDto) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *DnsCredentialDropdownDto) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *DnsCredentialDropdownDto) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetIsLocked

`func (o *DnsCredentialDropdownDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *DnsCredentialDropdownDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *DnsCredentialDropdownDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


