# CloudSubnetDetailsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubnetId** | **string** |  | 
**SubnetType** | **map[string]interface{}** |  | 
**Name** | Pointer to **string** |  | [optional] [default to ""]
**Zone** | Pointer to **string** |  | [optional] [default to ""]
**Cidr** | Pointer to **string** |  | [optional] [default to ""]
**Ipv6Cidr** | Pointer to **string** |  | [optional] [default to ""]
**AvailableIpCount** | Pointer to **NullableInt32** |  | [optional] 
**NodeCount** | Pointer to **int32** |  | [optional] [default to 0]
**IsDefault** | Pointer to **bool** |  | [optional] [default to false]
**OwnerId** | Pointer to **string** |  | [optional] [default to ""]
**VpcId** | Pointer to **string** |  | [optional] [default to ""]

## Methods

### NewCloudSubnetDetailsDto

`func NewCloudSubnetDetailsDto(subnetId string, subnetType map[string]interface{}, ) *CloudSubnetDetailsDto`

NewCloudSubnetDetailsDto instantiates a new CloudSubnetDetailsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSubnetDetailsDtoWithDefaults

`func NewCloudSubnetDetailsDtoWithDefaults() *CloudSubnetDetailsDto`

NewCloudSubnetDetailsDtoWithDefaults instantiates a new CloudSubnetDetailsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubnetId

`func (o *CloudSubnetDetailsDto) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *CloudSubnetDetailsDto) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *CloudSubnetDetailsDto) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.


### GetSubnetType

`func (o *CloudSubnetDetailsDto) GetSubnetType() map[string]interface{}`

GetSubnetType returns the SubnetType field if non-nil, zero value otherwise.

### GetSubnetTypeOk

`func (o *CloudSubnetDetailsDto) GetSubnetTypeOk() (*map[string]interface{}, bool)`

GetSubnetTypeOk returns a tuple with the SubnetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetType

`func (o *CloudSubnetDetailsDto) SetSubnetType(v map[string]interface{})`

SetSubnetType sets SubnetType field to given value.


### GetName

`func (o *CloudSubnetDetailsDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudSubnetDetailsDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudSubnetDetailsDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudSubnetDetailsDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetZone

`func (o *CloudSubnetDetailsDto) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *CloudSubnetDetailsDto) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *CloudSubnetDetailsDto) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *CloudSubnetDetailsDto) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetCidr

`func (o *CloudSubnetDetailsDto) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *CloudSubnetDetailsDto) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *CloudSubnetDetailsDto) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *CloudSubnetDetailsDto) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetIpv6Cidr

`func (o *CloudSubnetDetailsDto) GetIpv6Cidr() string`

GetIpv6Cidr returns the Ipv6Cidr field if non-nil, zero value otherwise.

### GetIpv6CidrOk

`func (o *CloudSubnetDetailsDto) GetIpv6CidrOk() (*string, bool)`

GetIpv6CidrOk returns a tuple with the Ipv6Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Cidr

`func (o *CloudSubnetDetailsDto) SetIpv6Cidr(v string)`

SetIpv6Cidr sets Ipv6Cidr field to given value.

### HasIpv6Cidr

`func (o *CloudSubnetDetailsDto) HasIpv6Cidr() bool`

HasIpv6Cidr returns a boolean if a field has been set.

### GetAvailableIpCount

`func (o *CloudSubnetDetailsDto) GetAvailableIpCount() int32`

GetAvailableIpCount returns the AvailableIpCount field if non-nil, zero value otherwise.

### GetAvailableIpCountOk

`func (o *CloudSubnetDetailsDto) GetAvailableIpCountOk() (*int32, bool)`

GetAvailableIpCountOk returns a tuple with the AvailableIpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableIpCount

`func (o *CloudSubnetDetailsDto) SetAvailableIpCount(v int32)`

SetAvailableIpCount sets AvailableIpCount field to given value.

### HasAvailableIpCount

`func (o *CloudSubnetDetailsDto) HasAvailableIpCount() bool`

HasAvailableIpCount returns a boolean if a field has been set.

### SetAvailableIpCountNil

`func (o *CloudSubnetDetailsDto) SetAvailableIpCountNil(b bool)`

 SetAvailableIpCountNil sets the value for AvailableIpCount to be an explicit nil

### UnsetAvailableIpCount
`func (o *CloudSubnetDetailsDto) UnsetAvailableIpCount()`

UnsetAvailableIpCount ensures that no value is present for AvailableIpCount, not even an explicit nil
### GetNodeCount

`func (o *CloudSubnetDetailsDto) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *CloudSubnetDetailsDto) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *CloudSubnetDetailsDto) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *CloudSubnetDetailsDto) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetIsDefault

`func (o *CloudSubnetDetailsDto) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *CloudSubnetDetailsDto) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *CloudSubnetDetailsDto) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *CloudSubnetDetailsDto) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetOwnerId

`func (o *CloudSubnetDetailsDto) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *CloudSubnetDetailsDto) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *CloudSubnetDetailsDto) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *CloudSubnetDetailsDto) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetVpcId

`func (o *CloudSubnetDetailsDto) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *CloudSubnetDetailsDto) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *CloudSubnetDetailsDto) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *CloudSubnetDetailsDto) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


