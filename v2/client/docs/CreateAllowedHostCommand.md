# CreateAllowedHostCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessProfileId** | **int32** |  | 
**Description** | **NullableString** |  | 
**IpAddress** | **NullableString** |  | 
**MaskBits** | **int32** |  | 

## Methods

### NewCreateAllowedHostCommand

`func NewCreateAllowedHostCommand(accessProfileId int32, description NullableString, ipAddress NullableString, maskBits int32, ) *CreateAllowedHostCommand`

NewCreateAllowedHostCommand instantiates a new CreateAllowedHostCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAllowedHostCommandWithDefaults

`func NewCreateAllowedHostCommandWithDefaults() *CreateAllowedHostCommand`

NewCreateAllowedHostCommandWithDefaults instantiates a new CreateAllowedHostCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessProfileId

`func (o *CreateAllowedHostCommand) GetAccessProfileId() int32`

GetAccessProfileId returns the AccessProfileId field if non-nil, zero value otherwise.

### GetAccessProfileIdOk

`func (o *CreateAllowedHostCommand) GetAccessProfileIdOk() (*int32, bool)`

GetAccessProfileIdOk returns a tuple with the AccessProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessProfileId

`func (o *CreateAllowedHostCommand) SetAccessProfileId(v int32)`

SetAccessProfileId sets AccessProfileId field to given value.


### GetDescription

`func (o *CreateAllowedHostCommand) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateAllowedHostCommand) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateAllowedHostCommand) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *CreateAllowedHostCommand) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateAllowedHostCommand) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIpAddress

`func (o *CreateAllowedHostCommand) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *CreateAllowedHostCommand) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *CreateAllowedHostCommand) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.


### SetIpAddressNil

`func (o *CreateAllowedHostCommand) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *CreateAllowedHostCommand) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetMaskBits

`func (o *CreateAllowedHostCommand) GetMaskBits() int32`

GetMaskBits returns the MaskBits field if non-nil, zero value otherwise.

### GetMaskBitsOk

`func (o *CreateAllowedHostCommand) GetMaskBitsOk() (*int32, bool)`

GetMaskBitsOk returns a tuple with the MaskBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskBits

`func (o *CreateAllowedHostCommand) SetMaskBits(v int32)`

SetMaskBits sets MaskBits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


