# DnsServerListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Address** | **NullableString** |  | 

## Methods

### NewDnsServerListDto

`func NewDnsServerListDto(id int32, address NullableString, ) *DnsServerListDto`

NewDnsServerListDto instantiates a new DnsServerListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsServerListDtoWithDefaults

`func NewDnsServerListDtoWithDefaults() *DnsServerListDto`

NewDnsServerListDtoWithDefaults instantiates a new DnsServerListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DnsServerListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DnsServerListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DnsServerListDto) SetId(v int32)`

SetId sets Id field to given value.


### GetAddress

`func (o *DnsServerListDto) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DnsServerListDto) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DnsServerListDto) SetAddress(v string)`

SetAddress sets Address field to given value.


### SetAddressNil

`func (o *DnsServerListDto) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *DnsServerListDto) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


