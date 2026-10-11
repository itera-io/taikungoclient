# GcpVpcDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** |  | 
**SelfLink** | **NullableString** |  | 
**IpMode** | [**IpMode**](IpMode.md) |  | 

## Methods

### NewGcpVpcDto

`func NewGcpVpcDto(name NullableString, selfLink NullableString, ipMode IpMode, ) *GcpVpcDto`

NewGcpVpcDto instantiates a new GcpVpcDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpVpcDtoWithDefaults

`func NewGcpVpcDtoWithDefaults() *GcpVpcDto`

NewGcpVpcDtoWithDefaults instantiates a new GcpVpcDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GcpVpcDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GcpVpcDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GcpVpcDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *GcpVpcDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GcpVpcDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSelfLink

`func (o *GcpVpcDto) GetSelfLink() string`

GetSelfLink returns the SelfLink field if non-nil, zero value otherwise.

### GetSelfLinkOk

`func (o *GcpVpcDto) GetSelfLinkOk() (*string, bool)`

GetSelfLinkOk returns a tuple with the SelfLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfLink

`func (o *GcpVpcDto) SetSelfLink(v string)`

SetSelfLink sets SelfLink field to given value.


### SetSelfLinkNil

`func (o *GcpVpcDto) SetSelfLinkNil(b bool)`

 SetSelfLinkNil sets the value for SelfLink to be an explicit nil

### UnsetSelfLink
`func (o *GcpVpcDto) UnsetSelfLink()`

UnsetSelfLink ensures that no value is present for SelfLink, not even an explicit nil
### GetIpMode

`func (o *GcpVpcDto) GetIpMode() IpMode`

GetIpMode returns the IpMode field if non-nil, zero value otherwise.

### GetIpModeOk

`func (o *GcpVpcDto) GetIpModeOk() (*IpMode, bool)`

GetIpModeOk returns a tuple with the IpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpMode

`func (o *GcpVpcDto) SetIpMode(v IpMode)`

SetIpMode sets IpMode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


