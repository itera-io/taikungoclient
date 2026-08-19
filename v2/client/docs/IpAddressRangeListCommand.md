# IpAddressRangeListCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddress** | **NullableString** |  | 
**NetMask** | **int32** |  | 
**Gateway** | **NullableString** |  | 

## Methods

### NewIpAddressRangeListCommand

`func NewIpAddressRangeListCommand(ipAddress NullableString, netMask int32, gateway NullableString, ) *IpAddressRangeListCommand`

NewIpAddressRangeListCommand instantiates a new IpAddressRangeListCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpAddressRangeListCommandWithDefaults

`func NewIpAddressRangeListCommandWithDefaults() *IpAddressRangeListCommand`

NewIpAddressRangeListCommandWithDefaults instantiates a new IpAddressRangeListCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAddress

`func (o *IpAddressRangeListCommand) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *IpAddressRangeListCommand) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *IpAddressRangeListCommand) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.


### SetIpAddressNil

`func (o *IpAddressRangeListCommand) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *IpAddressRangeListCommand) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetNetMask

`func (o *IpAddressRangeListCommand) GetNetMask() int32`

GetNetMask returns the NetMask field if non-nil, zero value otherwise.

### GetNetMaskOk

`func (o *IpAddressRangeListCommand) GetNetMaskOk() (*int32, bool)`

GetNetMaskOk returns a tuple with the NetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetMask

`func (o *IpAddressRangeListCommand) SetNetMask(v int32)`

SetNetMask sets NetMask field to given value.


### GetGateway

`func (o *IpAddressRangeListCommand) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *IpAddressRangeListCommand) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *IpAddressRangeListCommand) SetGateway(v string)`

SetGateway sets Gateway field to given value.


### SetGatewayNil

`func (o *IpAddressRangeListCommand) SetGatewayNil(b bool)`

 SetGatewayNil sets the value for Gateway to be an explicit nil

### UnsetGateway
`func (o *IpAddressRangeListCommand) UnsetGateway()`

UnsetGateway ensures that no value is present for Gateway, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


