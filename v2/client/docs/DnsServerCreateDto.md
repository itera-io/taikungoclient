# DnsServerCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDnsServerCreateDto

`func NewDnsServerCreateDto() *DnsServerCreateDto`

NewDnsServerCreateDto instantiates a new DnsServerCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsServerCreateDtoWithDefaults

`func NewDnsServerCreateDtoWithDefaults() *DnsServerCreateDto`

NewDnsServerCreateDtoWithDefaults instantiates a new DnsServerCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *DnsServerCreateDto) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DnsServerCreateDto) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DnsServerCreateDto) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *DnsServerCreateDto) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *DnsServerCreateDto) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *DnsServerCreateDto) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


