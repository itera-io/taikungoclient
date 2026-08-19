# AwsVpcDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VpcId** | **NullableString** |  | 
**Name** | **NullableString** |  | 
**CidrBlock** | **NullableString** |  | 
**IpMode** | [**IpMode**](IpMode.md) |  | 

## Methods

### NewAwsVpcDto

`func NewAwsVpcDto(vpcId NullableString, name NullableString, cidrBlock NullableString, ipMode IpMode, ) *AwsVpcDto`

NewAwsVpcDto instantiates a new AwsVpcDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsVpcDtoWithDefaults

`func NewAwsVpcDtoWithDefaults() *AwsVpcDto`

NewAwsVpcDtoWithDefaults instantiates a new AwsVpcDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpcId

`func (o *AwsVpcDto) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *AwsVpcDto) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *AwsVpcDto) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.


### SetVpcIdNil

`func (o *AwsVpcDto) SetVpcIdNil(b bool)`

 SetVpcIdNil sets the value for VpcId to be an explicit nil

### UnsetVpcId
`func (o *AwsVpcDto) UnsetVpcId()`

UnsetVpcId ensures that no value is present for VpcId, not even an explicit nil
### GetName

`func (o *AwsVpcDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AwsVpcDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AwsVpcDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *AwsVpcDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AwsVpcDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCidrBlock

`func (o *AwsVpcDto) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *AwsVpcDto) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *AwsVpcDto) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.


### SetCidrBlockNil

`func (o *AwsVpcDto) SetCidrBlockNil(b bool)`

 SetCidrBlockNil sets the value for CidrBlock to be an explicit nil

### UnsetCidrBlock
`func (o *AwsVpcDto) UnsetCidrBlock()`

UnsetCidrBlock ensures that no value is present for CidrBlock, not even an explicit nil
### GetIpMode

`func (o *AwsVpcDto) GetIpMode() IpMode`

GetIpMode returns the IpMode field if non-nil, zero value otherwise.

### GetIpModeOk

`func (o *AwsVpcDto) GetIpModeOk() (*IpMode, bool)`

GetIpModeOk returns a tuple with the IpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpMode

`func (o *AwsVpcDto) SetIpMode(v IpMode)`

SetIpMode sets IpMode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


