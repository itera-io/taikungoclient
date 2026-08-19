# GcpSubnetworkDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** |  | 
**SelfLink** | **NullableString** |  | 
**Region** | **NullableString** |  | 
**IpCidrRange** | **NullableString** |  | 
**HasIpv6** | **bool** |  | 
**AvailableIpCount** | **int32** |  | 

## Methods

### NewGcpSubnetworkDto

`func NewGcpSubnetworkDto(name NullableString, selfLink NullableString, region NullableString, ipCidrRange NullableString, hasIpv6 bool, availableIpCount int32, ) *GcpSubnetworkDto`

NewGcpSubnetworkDto instantiates a new GcpSubnetworkDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpSubnetworkDtoWithDefaults

`func NewGcpSubnetworkDtoWithDefaults() *GcpSubnetworkDto`

NewGcpSubnetworkDtoWithDefaults instantiates a new GcpSubnetworkDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GcpSubnetworkDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GcpSubnetworkDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GcpSubnetworkDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *GcpSubnetworkDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GcpSubnetworkDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSelfLink

`func (o *GcpSubnetworkDto) GetSelfLink() string`

GetSelfLink returns the SelfLink field if non-nil, zero value otherwise.

### GetSelfLinkOk

`func (o *GcpSubnetworkDto) GetSelfLinkOk() (*string, bool)`

GetSelfLinkOk returns a tuple with the SelfLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfLink

`func (o *GcpSubnetworkDto) SetSelfLink(v string)`

SetSelfLink sets SelfLink field to given value.


### SetSelfLinkNil

`func (o *GcpSubnetworkDto) SetSelfLinkNil(b bool)`

 SetSelfLinkNil sets the value for SelfLink to be an explicit nil

### UnsetSelfLink
`func (o *GcpSubnetworkDto) UnsetSelfLink()`

UnsetSelfLink ensures that no value is present for SelfLink, not even an explicit nil
### GetRegion

`func (o *GcpSubnetworkDto) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GcpSubnetworkDto) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GcpSubnetworkDto) SetRegion(v string)`

SetRegion sets Region field to given value.


### SetRegionNil

`func (o *GcpSubnetworkDto) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *GcpSubnetworkDto) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetIpCidrRange

`func (o *GcpSubnetworkDto) GetIpCidrRange() string`

GetIpCidrRange returns the IpCidrRange field if non-nil, zero value otherwise.

### GetIpCidrRangeOk

`func (o *GcpSubnetworkDto) GetIpCidrRangeOk() (*string, bool)`

GetIpCidrRangeOk returns a tuple with the IpCidrRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpCidrRange

`func (o *GcpSubnetworkDto) SetIpCidrRange(v string)`

SetIpCidrRange sets IpCidrRange field to given value.


### SetIpCidrRangeNil

`func (o *GcpSubnetworkDto) SetIpCidrRangeNil(b bool)`

 SetIpCidrRangeNil sets the value for IpCidrRange to be an explicit nil

### UnsetIpCidrRange
`func (o *GcpSubnetworkDto) UnsetIpCidrRange()`

UnsetIpCidrRange ensures that no value is present for IpCidrRange, not even an explicit nil
### GetHasIpv6

`func (o *GcpSubnetworkDto) GetHasIpv6() bool`

GetHasIpv6 returns the HasIpv6 field if non-nil, zero value otherwise.

### GetHasIpv6Ok

`func (o *GcpSubnetworkDto) GetHasIpv6Ok() (*bool, bool)`

GetHasIpv6Ok returns a tuple with the HasIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasIpv6

`func (o *GcpSubnetworkDto) SetHasIpv6(v bool)`

SetHasIpv6 sets HasIpv6 field to given value.


### GetAvailableIpCount

`func (o *GcpSubnetworkDto) GetAvailableIpCount() int32`

GetAvailableIpCount returns the AvailableIpCount field if non-nil, zero value otherwise.

### GetAvailableIpCountOk

`func (o *GcpSubnetworkDto) GetAvailableIpCountOk() (*int32, bool)`

GetAvailableIpCountOk returns a tuple with the AvailableIpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableIpCount

`func (o *GcpSubnetworkDto) SetAvailableIpCount(v int32)`

SetAvailableIpCount sets AvailableIpCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


