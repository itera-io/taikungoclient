# AwsSubnetDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubnetId** | **NullableString** |  | 
**Zone** | **NullableString** |  | 
**State** | **NullableString** |  | 
**HasIpv6** | **bool** |  | 
**Cidr** | **NullableString** |  | 
**Ipv6Cidr** | **NullableString** |  | 
**Name** | **NullableString** |  | 
**AvailableIpCount** | **NullableInt32** |  | 
**IsDefault** | **bool** |  | 
**OwnerId** | **NullableString** |  | 
**VpcId** | **NullableString** |  | 

## Methods

### NewAwsSubnetDto

`func NewAwsSubnetDto(subnetId NullableString, zone NullableString, state NullableString, hasIpv6 bool, cidr NullableString, ipv6Cidr NullableString, name NullableString, availableIpCount NullableInt32, isDefault bool, ownerId NullableString, vpcId NullableString, ) *AwsSubnetDto`

NewAwsSubnetDto instantiates a new AwsSubnetDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsSubnetDtoWithDefaults

`func NewAwsSubnetDtoWithDefaults() *AwsSubnetDto`

NewAwsSubnetDtoWithDefaults instantiates a new AwsSubnetDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubnetId

`func (o *AwsSubnetDto) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *AwsSubnetDto) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *AwsSubnetDto) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.


### SetSubnetIdNil

`func (o *AwsSubnetDto) SetSubnetIdNil(b bool)`

 SetSubnetIdNil sets the value for SubnetId to be an explicit nil

### UnsetSubnetId
`func (o *AwsSubnetDto) UnsetSubnetId()`

UnsetSubnetId ensures that no value is present for SubnetId, not even an explicit nil
### GetZone

`func (o *AwsSubnetDto) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *AwsSubnetDto) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *AwsSubnetDto) SetZone(v string)`

SetZone sets Zone field to given value.


### SetZoneNil

`func (o *AwsSubnetDto) SetZoneNil(b bool)`

 SetZoneNil sets the value for Zone to be an explicit nil

### UnsetZone
`func (o *AwsSubnetDto) UnsetZone()`

UnsetZone ensures that no value is present for Zone, not even an explicit nil
### GetState

`func (o *AwsSubnetDto) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AwsSubnetDto) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AwsSubnetDto) SetState(v string)`

SetState sets State field to given value.


### SetStateNil

`func (o *AwsSubnetDto) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *AwsSubnetDto) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetHasIpv6

`func (o *AwsSubnetDto) GetHasIpv6() bool`

GetHasIpv6 returns the HasIpv6 field if non-nil, zero value otherwise.

### GetHasIpv6Ok

`func (o *AwsSubnetDto) GetHasIpv6Ok() (*bool, bool)`

GetHasIpv6Ok returns a tuple with the HasIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasIpv6

`func (o *AwsSubnetDto) SetHasIpv6(v bool)`

SetHasIpv6 sets HasIpv6 field to given value.


### GetCidr

`func (o *AwsSubnetDto) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *AwsSubnetDto) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *AwsSubnetDto) SetCidr(v string)`

SetCidr sets Cidr field to given value.


### SetCidrNil

`func (o *AwsSubnetDto) SetCidrNil(b bool)`

 SetCidrNil sets the value for Cidr to be an explicit nil

### UnsetCidr
`func (o *AwsSubnetDto) UnsetCidr()`

UnsetCidr ensures that no value is present for Cidr, not even an explicit nil
### GetIpv6Cidr

`func (o *AwsSubnetDto) GetIpv6Cidr() string`

GetIpv6Cidr returns the Ipv6Cidr field if non-nil, zero value otherwise.

### GetIpv6CidrOk

`func (o *AwsSubnetDto) GetIpv6CidrOk() (*string, bool)`

GetIpv6CidrOk returns a tuple with the Ipv6Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Cidr

`func (o *AwsSubnetDto) SetIpv6Cidr(v string)`

SetIpv6Cidr sets Ipv6Cidr field to given value.


### SetIpv6CidrNil

`func (o *AwsSubnetDto) SetIpv6CidrNil(b bool)`

 SetIpv6CidrNil sets the value for Ipv6Cidr to be an explicit nil

### UnsetIpv6Cidr
`func (o *AwsSubnetDto) UnsetIpv6Cidr()`

UnsetIpv6Cidr ensures that no value is present for Ipv6Cidr, not even an explicit nil
### GetName

`func (o *AwsSubnetDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AwsSubnetDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AwsSubnetDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *AwsSubnetDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AwsSubnetDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAvailableIpCount

`func (o *AwsSubnetDto) GetAvailableIpCount() int32`

GetAvailableIpCount returns the AvailableIpCount field if non-nil, zero value otherwise.

### GetAvailableIpCountOk

`func (o *AwsSubnetDto) GetAvailableIpCountOk() (*int32, bool)`

GetAvailableIpCountOk returns a tuple with the AvailableIpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableIpCount

`func (o *AwsSubnetDto) SetAvailableIpCount(v int32)`

SetAvailableIpCount sets AvailableIpCount field to given value.


### SetAvailableIpCountNil

`func (o *AwsSubnetDto) SetAvailableIpCountNil(b bool)`

 SetAvailableIpCountNil sets the value for AvailableIpCount to be an explicit nil

### UnsetAvailableIpCount
`func (o *AwsSubnetDto) UnsetAvailableIpCount()`

UnsetAvailableIpCount ensures that no value is present for AvailableIpCount, not even an explicit nil
### GetIsDefault

`func (o *AwsSubnetDto) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *AwsSubnetDto) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *AwsSubnetDto) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetOwnerId

`func (o *AwsSubnetDto) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *AwsSubnetDto) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *AwsSubnetDto) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.


### SetOwnerIdNil

`func (o *AwsSubnetDto) SetOwnerIdNil(b bool)`

 SetOwnerIdNil sets the value for OwnerId to be an explicit nil

### UnsetOwnerId
`func (o *AwsSubnetDto) UnsetOwnerId()`

UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
### GetVpcId

`func (o *AwsSubnetDto) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *AwsSubnetDto) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *AwsSubnetDto) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.


### SetVpcIdNil

`func (o *AwsSubnetDto) SetVpcIdNil(b bool)`

 SetVpcIdNil sets the value for VpcId to be an explicit nil

### UnsetVpcId
`func (o *AwsSubnetDto) UnsetVpcId()`

UnsetVpcId ensures that no value is present for VpcId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


